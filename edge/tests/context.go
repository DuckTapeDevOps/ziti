// +build apitests

/*
	Copyright 2019 Netfoundry, Inc.

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

	https://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/

package tests

import (
	cryptoTls "crypto/tls"
	"fmt"
	"gopkg.in/resty.v1"
	"net"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"testing"
	"time"

	"github.com/netfoundry/ziti-foundation/common/constants"
	"github.com/netfoundry/ziti-edge/edge/controller/server"
	"github.com/netfoundry/ziti-fabric/fabric/controller"
	"github.com/Jeffail/gabs"
	"github.com/google/uuid"
	"github.com/michaelquigley/pfxlog"
	"github.com/netfoundry/ziti-foundation/transport"
	"github.com/netfoundry/ziti-foundation/transport/quic"
	"github.com/netfoundry/ziti-foundation/transport/tcp"
	"github.com/netfoundry/ziti-foundation/transport/tls"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func init() {
	pfxlog.Global(logrus.DebugLevel)
	pfxlog.SetPrefix("bitbucket.org/netfoundry/")
	logrus.SetFormatter(pfxlog.NewFormatterStartingToday())

	transport.AddAddressParser(quic.AddressParser{})
	transport.AddAddressParser(tls.AddressParser{})
	transport.AddAddressParser(tcp.AddressParser{})
}

type TestContext struct {
	ApiHost            string
	AdminPassword      string
	AdminUsername      string
	fabricController   *controller.Controller
	EdgeController     *server.Controller
	adminSessionId     string
	req                *require.Assertions
	client             *resty.Client
	enabledJsonLogging bool
	clusterId          string
}

var defaultTestContext = &TestContext{}

func NewTestContext(t *testing.T) *TestContext {
	return &TestContext{
		ApiHost:       "127.0.0.1:1280",
		AdminUsername: "admin", // make this uuid.New().String() once we're off of PG
		AdminPassword: "admin", // make this uuid.New().String() once we're off of PG
		req:           require.New(t),
	}
}

func GetTestContext() *TestContext {
	return defaultTestContext
}

func (ctx *TestContext) Transport() *http.Transport {
	return &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		TLSClientConfig: &cryptoTls.Config{
			InsecureSkipVerify: true,
		},
	}
}

func (ctx *TestContext) HttpClient(transport *http.Transport) *http.Client {
	jar, err := cookiejar.New(&cookiejar.Options{})
	ctx.req.NoError(err)

	return &http.Client{
		Transport:     transport,
		CheckRedirect: nil,
		Jar:           jar,
		Timeout:       20 * time.Second,
	}
}

func (ctx *TestContext) Client(httpClient *http.Client) *resty.Client {
	client := resty.NewWithClient(httpClient)
	return client
}

func (ctx *TestContext) NewClient() *resty.Client {
	return ctx.Client(ctx.HttpClient(ctx.Transport()))
}

func (ctx *TestContext) DefaultClient() *resty.Client {
	if ctx.client == nil {
		ctx.client = ctx.Client(ctx.HttpClient(ctx.Transport()))
		ctx.client.SetHostURL("https://" + ctx.ApiHost)
	}
	return ctx.client
}

func (ctx *TestContext) startServer() {
	log := pfxlog.Logger()
	_ = os.Mkdir("testdata", os.FileMode(0755))
	_ = os.Remove("testdata/ctrl.db")

	log.Info("loading config")
	config, err := controller.LoadConfig("ats-ctrl.yml")
	ctx.req.NoError(err)

	log.Info("creating fabric controller")
	ctx.fabricController, err = controller.NewController(config)
	ctx.req.NoError(err)

	log.Info("creating edge controller")
	ctx.EdgeController, err = server.NewController(config)
	ctx.req.NoError(err)

	ctx.EdgeController.SetHostController(ctx.fabricController)

	ctx.EdgeController.Initialize()

	err = ctx.EdgeController.AppEnv.Handlers.Identity.HandleInitializeDefaultAdmin(ctx.AdminUsername, ctx.AdminPassword, uuid.New().String())
	if err != nil {
		log.WithError(err).Warn("error during initialize admin")
	}

	// Note we're not starting the fabric controller. Shouldn't need any of it for testing the edge API
	ctx.EdgeController.Run()
	go func() {
		err = ctx.fabricController.Run()
		ctx.req.NoError(err)
	}()
	err = ctx.waitForPort(time.Minute * 5)
	ctx.req.NoError(err)
}

func (ctx *TestContext) waitForPort(duration time.Duration) error {
	now := time.Now()
	endTime := now.Add(duration)
	maxWait := duration
	for {
		conn, err := net.DialTimeout("tcp", ctx.ApiHost, maxWait)
		if err == nil {
			_ = conn.Close()
			return nil
		}
		now = time.Now()
		if !now.Before(endTime) {
			return err
		}
		maxWait = endTime.Sub(now)
		time.Sleep(10 * time.Millisecond)
	}
}

func (ctx *TestContext) requireAdminLogin() {
	ctx.adminSessionId = ctx.requireLogin(ctx.AdminUsername, ctx.AdminPassword)
}

func (ctx *TestContext) requireLogin(username, password string) string {
	sessionId, err := ctx.login(username, password)
	ctx.req.NoError(err)
	return sessionId
}

func (ctx *TestContext) login(username, password string) (string, error) {
	client := ctx.DefaultClient()

	body := gabs.New()
	if _, err := body.SetP(username, "username"); err != nil {
		return "", errors.WithStack(err)
	}
	if _, err := body.SetP(password, "password"); err != nil {
		return "", errors.WithStack(err)
	}

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body.String()).
		Post("/authenticate?method=password")

	if err != nil {
		return "", errors.WithStack(err)
	}

	if resp.StatusCode() != http.StatusOK {
		return "", errors.Errorf("expected status code %d got %d", http.StatusOK, resp.StatusCode())
	}

	sessionId := resp.Header().Get("zt-session")
	if sessionId == "" {
		return "", errors.New("expected header zt-session to not be empty")
	}
	return sessionId, nil
}

func (ctx *TestContext) teardown() {
	pfxlog.Logger().Info("tearing down test context")
	ctx.EdgeController.Shutdown()
	ctx.fabricController.Shutdown()
}

func (ctx *TestContext) requireCreateCluster(name string) string {
	entityJson := ctx.newNamedEntityJson(name).String()
	httpCode, body := ctx.createEntityOfType("clusters", entityJson)
	ctx.req.Equal(http.StatusCreated, httpCode)
	return ctx.getEntityId(body)
}

func (ctx *TestContext) requireCreateIdentity(name string, password string, isAdmin bool) string {
	entityData := gabs.New()
	ctx.setJsonValue(entityData, name, "name")
	ctx.setJsonValue(entityData, "User", "type")
	ctx.setJsonValue(entityData, isAdmin, "isAdmin")

	enrollments := map[string]interface{}{
		"updb": name,
	}
	ctx.setJsonValue(entityData, enrollments, "enrollment")

	entityJson := entityData.String()
	httpCode, body := ctx.createEntityOfType("identities", entityJson)
	ctx.req.Equal(http.StatusCreated, httpCode)
	id := ctx.getEntityId(body)
	ctx.completeEnrollment(id, password)
	return id
}

func (ctx *TestContext) requireCreateEntity(entity testEntity) string {
	httpStatus, body := ctx.createEntity(entity)
	ctx.req.Equal(http.StatusCreated, httpStatus)
	return ctx.getEntityId(body)
}

func (ctx *TestContext) requireDeleteEntity(entity testEntity) {
	httpStatus, _ := ctx.deleteEntityOfType(entity.getEntityType(), entity.getId())
	ctx.req.Equal(http.StatusOK, httpStatus)
}

func (ctx *TestContext) createEntity(entity testEntity) (int, []byte) {
	return ctx.createEntityOfType(entity.getEntityType(), entity.toJson(true, ctx))
}

func (ctx *TestContext) createEntityOfType(entityType string, body string) (int, []byte) {
	client := ctx.DefaultClient()
	resp, err := client.
		R().
		SetHeader("Content-Type", "application/json").
		SetHeader(constants.ZitiSession, ctx.adminSessionId).
		SetBody(body).
		Post("/" + entityType)

	ctx.req.NoError(err)
	ctx.logJson(resp.Body())
	return resp.StatusCode(), resp.Body()
}

func (ctx *TestContext) deleteEntityOfType(entityType string, id string) (int, []byte) {
	client := ctx.DefaultClient()
	resp, err := client.
		R().
		SetHeader("Content-Type", "application/json").
		SetHeader(constants.ZitiSession, ctx.adminSessionId).
		Delete("/" + entityType + "/" + id)

	ctx.req.NoError(err)
	ctx.logJson(resp.Body())
	return resp.StatusCode(), resp.Body()
}

func (ctx *TestContext) requireUpdateEntity(entity testEntity) {
	httpStatus, _ := ctx.updateEntity(entity)
	ctx.req.Equal(http.StatusOK, httpStatus)
}

func (ctx *TestContext) updateEntity(entity testEntity) (int, []byte) {
	return ctx.updateEntityOfType(entity.getId(), entity.getEntityType(), entity.toJson(false, ctx))
}

func (ctx *TestContext) updateEntityOfType(id string, entityType string, body string) (int, []byte) {
	client := ctx.DefaultClient()
	urlPath := fmt.Sprintf("/%v/%v", entityType, id)
	pfxlog.Logger().Infof("url path: %v", urlPath)
	resp, err := client.
		R().
		SetHeader("Content-Type", "application/json").
		SetHeader(constants.ZitiSession, ctx.adminSessionId).
		SetBody(body).
		Put(urlPath)

	ctx.req.NoError(err)
	ctx.logJson(resp.Body())
	return resp.StatusCode(), resp.Body()
}

func (ctx *TestContext) completeEnrollment(identityId string, password string) {
	result := ctx.requireQuery(ctx.adminSessionId, fmt.Sprintf("identities/%v", identityId))
	path := result.Search(path("data.enrollment.updb.token")...)
	ctx.req.NotNil(path)
	str, ok := path.Data().(string)
	ctx.req.True(ok)

	enrollBody := gabs.New()
	ctx.setJsonValue(enrollBody, password, "password")

	resp, err := ctx.DefaultClient().
		R().
		SetHeader("Content-Type", "application/json").
		SetHeader(constants.ZitiSession, ctx.adminSessionId).
		SetBody(enrollBody.String()).
		Post("enroll?token=" + str)
	ctx.req.NoError(err)
	ctx.logJson(resp.Body())
	ctx.req.Equal(http.StatusOK, resp.StatusCode())
}

func (ctx *TestContext) requireQuery(token, url string) *gabs.Container {
	httpStatus, body := ctx.query(token, url)
	ctx.logJson(body)
	ctx.req.Equal(http.StatusOK, httpStatus)
	return ctx.parseJson(body)
}

func (ctx *TestContext) query(token, url string) (int, []byte) {
	client := ctx.DefaultClient()
	pfxlog.Logger().Infof("using session id: %v", token)
	resp, err := client.
		R().
		SetHeader("Content-Type", "application/json").
		SetHeader(constants.ZitiSession, token).
		Get("/" + url)
	ctx.req.NoError(err)
	return resp.StatusCode(), resp.Body()
}

func (ctx *TestContext) requireAddAssociation(url string, ids ...string) {
	httpStatus, _ := ctx.addAssociation(url, ids...)
	ctx.req.Equal(http.StatusOK, httpStatus)
}

func (ctx *TestContext) addAssociation(url string, ids ...string) (int, []byte) {
	return ctx.updateAssociation(http.MethodPut, url, ids...)
}

func (ctx *TestContext) requireRemoveAssociation(url string, ids ...string) {
	httpStatus, _ := ctx.removeAssociation(url, ids...)
	ctx.req.Equal(http.StatusOK, httpStatus)
}

func (ctx *TestContext) removeAssociation(url string, ids ...string) (int, []byte) {
	return ctx.updateAssociation(http.MethodDelete, url, ids...)
}

func (ctx *TestContext) updateAssociation(method, url string, ids ...string) (int, []byte) {
	client := ctx.DefaultClient()
	resp, err := client.
		R().
		SetHeader("Content-Type", "application/json").
		SetHeader(constants.ZitiSession, ctx.adminSessionId).
		SetBody(ctx.idsJson(ids...).String()).
		Execute(method, "/"+url)
	ctx.req.NoError(err)
	ctx.logJson(resp.Body())
	return resp.StatusCode(), resp.Body()
}

func (ctx *TestContext) isServiceVisibleToUser(info *userInfo, serviceId string) bool {
	query := url.QueryEscape(fmt.Sprintf(`id = "%v"`, serviceId))
	result := ctx.requireQuery(info.sessionId, "services?filter="+query)
	data := ctx.requirePath(result, "data")
	return nil != ctx.childWith(data, "id", serviceId)
}

func (ctx *TestContext) newTestService() *testService {
	if ctx.clusterId == "" {
		ctx.clusterId = ctx.requireCreateCluster(uuid.New().String())
	}

	return &testService{
		name:            uuid.New().String(),
		dnsHostname:     uuid.New().String(),
		dnsPort:         0,
		egressRouter:    uuid.New().String(),
		endpointAddress: uuid.New().String(),
		clusterIds:      []string{ctx.clusterId},
		hostIds:         nil,
		tags:            nil,
	}
}

func (ctx *TestContext) requireCreateNewService() *testService {
	service := ctx.newTestService()
	service.id = ctx.requireCreateEntity(service)
	return service
}

type userInfo struct {
	username   string
	password   string
	identityId string
	sessionId  string
}

func (ctx *TestContext) createUserAndLogin(isAdmin bool) *userInfo {
	result := &userInfo{
		username: uuid.New().String(),
		password: uuid.New().String(),
	}
	result.identityId = ctx.requireCreateIdentity(result.username, result.password, isAdmin)
	result.sessionId = ctx.requireLogin(result.username, result.password)
	return result
}

func (ctx *TestContext) validateEntityWithQuery(entity testEntity) *gabs.Container {
	return ctx.validateEntityWithQueryAndSession(ctx.adminSessionId, entity)
}

func (ctx *TestContext) getEntityDates(jsonEntity *gabs.Container) (time.Time, time.Time) {
	createdAtStr := jsonEntity.S("createdAt").Data().(string)
	updatedAtStr := jsonEntity.S("updatedAt").Data().(string)

	ctx.req.NotNil(createdAtStr)
	ctx.req.NotNil(updatedAtStr)

	createdAt, err := time.Parse(time.RFC3339, createdAtStr)
	ctx.req.NoError(err)
	updatedAt, err := time.Parse(time.RFC3339, updatedAtStr)
	ctx.req.NoError(err)
	return createdAt, updatedAt
}

func (ctx *TestContext) validateDateFieldsForCreate(start time.Time, jsonEntity *gabs.Container) time.Time {
	now := time.Now()
	createdAt, updatedAt := ctx.getEntityDates(jsonEntity)
	ctx.req.Equal(createdAt, updatedAt)

	ctx.req.True(start.Before(createdAt) || start.Equal(createdAt))
	ctx.req.True(now.After(createdAt) || now.Equal(createdAt))

	return createdAt
}

func (ctx *TestContext) validateDateFieldsForUpdate(start time.Time, origCreatedAt time.Time, jsonEntity *gabs.Container) time.Time {
	now := time.Now()
	createdAt, updatedAt := ctx.getEntityDates(jsonEntity)
	ctx.req.Equal(origCreatedAt, createdAt)

	ctx.req.True(createdAt.Before(updatedAt))
	ctx.req.True(start.Before(updatedAt) || start.Equal(updatedAt))
	ctx.req.True(now.After(updatedAt) || now.Equal(updatedAt))

	return createdAt
}

func (ctx *TestContext) validateEntityWithQueryAndSession(sessionId string, entity testEntity) *gabs.Container {
	query := url.QueryEscape(fmt.Sprintf(`id = "%v"`, entity.getId()))
	result := ctx.requireQuery(sessionId, entity.getEntityType()+"?filter="+query)
	data := ctx.requirePath(result, "data")
	jsonEntity := ctx.requireChildWith(data, "id", entity.getId())
	return ctx.validateEntity(entity, jsonEntity)
}

func (ctx *TestContext) validateEntityWithLookup(entity testEntity) *gabs.Container {
	return ctx.validateEntityWithLookupAndSession(ctx.adminSessionId, entity)
}

func (ctx *TestContext) validateEntityWithLookupAndSession(sessionId string, entity testEntity) *gabs.Container {
	result := ctx.requireQuery(sessionId, entity.getEntityType()+"/"+entity.getId())
	jsonEntity := ctx.requirePath(result, "data")
	return ctx.validateEntity(entity, jsonEntity)
}

func (ctx *TestContext) validateEntity(entity testEntity, jsonEntity *gabs.Container) *gabs.Container {
	entity.validate(ctx, jsonEntity)
	return jsonEntity
}

func toIntfSlice(in []string) []interface{} {
	var result []interface{}
	for _, i := range in {
		result = append(result, i)
	}
	return result
}
