/*
	Copyright NetFoundry, Inc.

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

package helpers

import (
	"fmt"
	"github.com/openziti/ziti/ziti/cmd/ziti/constants"
	"github.com/pkg/errors"
	"os"
	"strings"
)

const (
	PathSeparator = "/"

	ZitiHomeVarName = "ZITI_HOME"

	ZitiCtrlNameVarName = "ZITI_CONTROLLER_NAME"

	ZitiEdgeRouterHostnameVarName = "ZITI_EDGE_ROUTER_HOSTNAME"

	ZitiEdgeRouterPortVarName = "ZITI_EDGE_ROUTER_PORT"

	ZitiEdgeCtrlIdentityCertVarName = "ZITI_EDGE_CTRL_IDENTITY_CERT"

	ZitiEdgeCtrlIdentityServerCertVarName = "ZITI_EDGE_CTRL_IDENTITY_SERVER_CERT"

	ZitiEdgeCtrlIdentityKeyVarName = "ZITI_EDGE_CTRL_IDENTITY_KEY"

	ZitiEdgeCtrlIdentityCAVarName = "ZITI_EDGE_CTRL_IDENTITY_CA"

	ZitiCtrlIdentityCertVarName = "ZITI_CTRL_IDENTITY_CERT"

	ZitiCtrlIdentityServerCertVarName = "ZITI_CTRL_IDENTITY_SERVER_CERT"

	ZitiCtrlIdentityKeyVarName = "ZITI_CTRL_IDENTITY_KEY"

	ZitiCtrlIdentityCAVarName = "ZITI_CTRL_IDENTITY_CA"

	ZitiSigningCertVarName = "ZITI_SIGNING_CERT"

	ZitiSigningKeyVarName = "ZITI_SIGNING_KEY"

	ZitiCtrlListenerHostPortVarName = "ZITI_CTRL_LISTENER_HOST_PORT"

	ZitiCtrlMgmtListenerHostPortVarName = "ZITI_CTRL_MGMT_HOST_PORT"

	ZitiEdgeCtrlListenerHostPortVarName = "ZITI_CTRL_EDGE_LISTENER_HOST_PORT"

	ZitiEdgeCtrlAdvertisedHostPortVarName = "ZITI_EDGE_CTRL_ADVERTISED_HOST_PORT"
)

func HomeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	h := os.Getenv("USERPROFILE") // windows
	if h == "" {
		h = "."
	}
	return strings.ReplaceAll(h, "\\", PathSeparator)
}

func WorkingDir() (string, error) {
	wd, err := os.Getwd()
	if wd == "" || err != nil {
		return "", err
	}

	return strings.ReplaceAll(wd, "\\", PathSeparator), nil
}

func GetZitiHome() (string, error) {

	// Get path from env variable
	retVal := os.Getenv(ZitiHomeVarName)

	if retVal == "" {
		// If not set, create a default path of the current working directory
		workingDir, err := WorkingDir()
		if err != nil {
			return "", err
		}

		err = os.Setenv(ZitiHomeVarName, workingDir)
		if err != nil {
			return "", err
		}

		retVal = os.Getenv(ZitiHomeVarName)
	}

	return strings.ReplaceAll(retVal, "\\", PathSeparator), nil
}

func GetZitiIdentityCert() (string, error) {
	workingDir, err := WorkingDir()
	if err != nil {
		return "", err
	}

	controllerName, err := GetZitiCtrlName()
	if err != nil {
		err := errors.Wrap(err, "Unable to get "+ZitiCtrlNameVarName)
		if err != nil {
			return "", err
		}
	}

	return getOrSetEnvVar(ZitiCtrlIdentityCertVarName, fmt.Sprintf("%s/%s-client.cert", workingDir, controllerName))
}

func GetZitiIdentityServerCert() (string, error) {
	workingDir, err := WorkingDir()
	if err != nil {
		return "", err
	}

	controllerName, err := GetZitiCtrlName()
	if err != nil {
		err := errors.Wrap(err, "Unable to get "+ZitiCtrlNameVarName)
		if err != nil {
			return "", err
		}
	}

	return getOrSetEnvVar(ZitiCtrlIdentityServerCertVarName, fmt.Sprintf("%s/%s-server.pem", workingDir, controllerName))
}

func GetZitiIdentityKey() (string, error) {
	workingDir, err := WorkingDir()
	if err != nil {
		return "", err
	}

	controllerName, err := GetZitiCtrlName()
	if err != nil {
		err := errors.Wrap(err, "Unable to get "+ZitiCtrlNameVarName)
		if err != nil {
			return "", err
		}
	}

	return getOrSetEnvVar(ZitiCtrlIdentityKeyVarName, fmt.Sprintf("%s/%s-server.key", workingDir, controllerName))
}

func GetZitiIdentityCA() (string, error) {
	workingDir, err := WorkingDir()
	if err != nil {
		return "", err
	}

	controllerName, err := GetZitiCtrlName()
	if err != nil {
		err := errors.Wrap(err, "Unable to get "+ZitiCtrlNameVarName)
		if err != nil {
			return "", err
		}
	}

	return getOrSetEnvVar(ZitiCtrlIdentityCAVarName, fmt.Sprintf("%s/%s-cas.pem", workingDir, controllerName))
}

func GetZitiEdgeIdentityCert() (string, error) {
	workingDir, err := WorkingDir()
	if err != nil {
		return "", err
	}

	controllerName, err := GetZitiCtrlName()
	if err != nil {
		err := errors.Wrap(err, "Unable to get "+ZitiCtrlNameVarName)
		if err != nil {
			return "", err
		}
	}

	return getOrSetEnvVar(ZitiEdgeCtrlIdentityCertVarName, fmt.Sprintf("%s/%s-client.cert", workingDir, controllerName))
}

func GetZitiEdgeIdentityServerCert() (string, error) {
	workingDir, err := WorkingDir()
	if err != nil {
		return "", err
	}

	controllerName, err := GetZitiCtrlName()
	if err != nil {
		err := errors.Wrap(err, "Unable to get "+ZitiCtrlNameVarName)
		if err != nil {
			return "", err
		}
	}

	return getOrSetEnvVar(ZitiEdgeCtrlIdentityServerCertVarName, fmt.Sprintf("%s/%s-server.pem", workingDir, controllerName))
}

func GetZitiEdgeIdentityKey() (string, error) {
	workingDir, err := WorkingDir()
	if err != nil {
		return "", err
	}

	controllerName, err := GetZitiCtrlName()
	if err != nil {
		err := errors.Wrap(err, "Unable to get "+ZitiCtrlNameVarName)
		if err != nil {
			return "", err
		}
	}

	return getOrSetEnvVar(ZitiEdgeCtrlIdentityKeyVarName, fmt.Sprintf("%s/%s-server.key", workingDir, controllerName))
}

func GetZitiEdgeIdentityCA() (string, error) {
	workingDir, err := WorkingDir()
	if err != nil {
		return "", err
	}

	controllerName, err := GetZitiCtrlName()
	if err != nil {
		err := errors.Wrap(err, "Unable to get "+ZitiCtrlNameVarName)
		if err != nil {
			return "", err
		}
	}

	return getOrSetEnvVar(ZitiEdgeCtrlIdentityCAVarName, fmt.Sprintf("%s/%s-cas.pem", workingDir, controllerName))
}

func GetZitiCtrlListenerHostPort() (string, error) {
	return getOrSetEnvVar(ZitiCtrlListenerHostPortVarName, constants.DefaultZitiControllerListenerHostPort)
}

func GetZitiCtrlMgmtListenerHostPort() (string, error) {
	return getOrSetEnvVar(ZitiCtrlMgmtListenerHostPortVarName, constants.DefaultZitiMgmtControllerListenerHostPort)
}

func GetZitiCtrlName() (string, error) {
	return getOrSetEnvVar(ZitiCtrlNameVarName, constants.DefaultZitiControllerName)
}

func GetZitiEdgeRouterHostname() (string, error) {
	hostname, err := os.Hostname()
	if err != nil {
		err := errors.Wrap(err, "Unable to get hostname")
		if err != nil {
			return "", err
		}
	}
	return getOrSetEnvVar(ZitiEdgeRouterHostnameVarName, hostname)
}

func GetZitiEdgeRouterPort() (string, error) {
	return getOrSetEnvVar(ZitiEdgeRouterPortVarName, constants.DefaultZitiEdgeRouterPort)
}

func GetZitiSigningCert() (string, error) {
	workingDir, err := WorkingDir()
	if err != nil {
		return "", err
	}

	return getOrSetEnvVar(ZitiSigningCertVarName, fmt.Sprintf("%s/signingCert.cert", workingDir))
}

func GetZitiSigningKey() (string, error) {
	workingDir, err := WorkingDir()
	if err != nil {
		return "", err
	}

	return getOrSetEnvVar(ZitiSigningKeyVarName, fmt.Sprintf("%s/signingKey.key", workingDir))
}

func GetZitiEdgeCtrlListenerHostPort() (string, error) {
	return getOrSetEnvVar(ZitiEdgeCtrlListenerHostPortVarName, constants.DefaultZitiEdgeListenerHostPort)
}

func GetZitiEdgeCtrlAdvertisedHostPort() (string, error) {
	edgeCtrlListenerHostPort, err := GetZitiEdgeCtrlListenerHostPort()
	if err != nil {
		err := errors.Wrap(err, "Unable to get "+ZitiEdgeCtrlListenerHostPortVarName)
		if err != nil {
			return "", err
		}
	}

	return getOrSetEnvVar(ZitiEdgeCtrlAdvertisedHostPortVarName, edgeCtrlListenerHostPort)
}

func getOrSetEnvVar(envVarName string, defaultValue string) (string, error) {

	// Get path from env variable
	retVal := os.Getenv(envVarName)
	if retVal != "" {
		return retVal, nil
	}

	err := os.Setenv(envVarName, defaultValue)
	if err != nil {
		return "", err
	}

	retVal = os.Getenv(envVarName)

	return retVal, nil
}
