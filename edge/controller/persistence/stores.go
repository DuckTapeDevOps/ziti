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

package persistence

import (
	"github.com/netfoundry/ziti-foundation/storage/boltz"
	"github.com/netfoundry/ziti-foundation/util/errorz"
	"go.etcd.io/bbolt"
)

type Stores struct {
	DbProvider DbProvider

	ApiSession    ApiSessionStore
	Appwan        AppwanStore
	Ca            CaStore
	Cluster       ClusterStore
	EdgeRouter    EdgeRouterStore
	EdgeService   EdgeServiceStore
	EventLog      EventLogStore
	GeoRegion     GeoRegionStore
	Identity      IdentityStore
	IdentityType  IdentityTypeStore
	Session       SessionStore
	Enrollment    EnrollmentStore
	Authenticator AuthenticatorStore

	storeList []Store
}

func (stores *Stores) getStoreForEntity(entity boltz.BaseEntity) boltz.CrudStore {
	for _, store := range stores.storeList {
		if store.GetEntityType() == entity.GetEntityType() {
			return store
		}
	}
	return nil
}

type stores struct {
	DbProvider DbProvider

	apiSession    *apiSessionStoreImpl
	appwan        *appwanStoreImpl
	ca            *caStoreImpl
	cluster       *clusterStoreImpl
	edgeRouter    *edgeRouterStoreImpl
	edgeService   *edgeServiceStoreImpl
	eventLog      *eventLogStoreImpl
	geoRegion     *geoRegionStoreImpl
	identity      *identityStoreImpl
	identityType  *IdentityTypeStoreImpl
	session       *sessionStoreImpl
	enrollment    *enrollmentStoreImpl
	authenticator *authenticatorStoreImpl
}

func NewBoltStores(dbProvider DbProvider) (*Stores, error) {
	errorHolder := &errorz.ErrorHolderImpl{}

	internalStores := &stores{
		DbProvider: dbProvider,
	}

	internalStores.apiSession = newApiSessionStore(internalStores)
	internalStores.appwan = newAppwanStore(internalStores)
	internalStores.ca = newCaStore(internalStores)
	internalStores.cluster = newClusterStore(internalStores)
	internalStores.edgeRouter = newEdgeRouterStore(internalStores)
	internalStores.edgeService = newEdgeServiceStore(internalStores, dbProvider.GetServiceStore())
	internalStores.eventLog = newEventLogStore(internalStores)
	internalStores.geoRegion = newGeoRegionStore(internalStores)
	internalStores.identity = newIdentityStore(internalStores)
	internalStores.identityType = newIdentityTypeStore(internalStores)
	internalStores.enrollment = newEnrollmentStore(internalStores)
	internalStores.authenticator = newAuthenticatorStore(internalStores)
	internalStores.session = newSessionStore(internalStores)

	externalStores := &Stores{
		DbProvider: dbProvider,

		ApiSession:    internalStores.apiSession,
		Appwan:        internalStores.appwan,
		Ca:            internalStores.ca,
		Cluster:       internalStores.cluster,
		EdgeRouter:    internalStores.edgeRouter,
		EdgeService:   internalStores.edgeService,
		EventLog:      internalStores.eventLog,
		GeoRegion:     internalStores.geoRegion,
		Identity:      internalStores.identity,
		IdentityType:  internalStores.identityType,
		Session:       internalStores.session,
		Authenticator: internalStores.authenticator,
		Enrollment:    internalStores.enrollment,
	}

	externalStores.storeList = []Store{
		internalStores.apiSession,
		internalStores.appwan,
		internalStores.ca,
		internalStores.cluster,
		internalStores.edgeRouter,
		internalStores.edgeService,
		internalStores.geoRegion,
		internalStores.identity,
		internalStores.identityType,
		internalStores.session,
		internalStores.authenticator,
		internalStores.enrollment,
	}

	err := dbProvider.GetDb().Update(func(tx *bbolt.Tx) error {
		for _, store := range externalStores.storeList {
			store.initializeLocal()
		}
		for _, store := range externalStores.storeList {
			store.initializeLinked()
		}
		for _, store := range externalStores.storeList {
			store.initializeIndexes(tx, errorHolder)
		}
		return nil
	})

	errorHolder.SetError(err)
	if errorHolder.HasError() {
		return nil, errorHolder.GetError()
	}
	return externalStores, nil
}

type QueryContext struct {
	Stores *Stores
	Tx     *bbolt.Tx
}
