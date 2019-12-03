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

package xgress_proxy

import (
	"github.com/netfoundry/ziti-fabric/fabric/xgress"
	"github.com/netfoundry/ziti-foundation/identity/identity"
	"fmt"
)

func NewFactory(id *identity.TokenId, ctrl xgress.CtrlChannel) xgress.XgressFactory {
	return &factory{id: id, ctrl: ctrl}
}

func (factory *factory) CreateListener(optionsData xgress.XgressOptionsData) (xgress.XgressListener, error) {
	options := xgress.LoadOptions(optionsData)
	service := ""
	if value, found := optionsData["service"]; found {
		service = value.(string)
	} else {
		return nil, fmt.Errorf("missing 'service' configuration option")
	}
	return newListener(factory.id, factory.ctrl, options, service), nil
}

func (factory *factory) CreateDialer(optionsData xgress.XgressOptionsData) (xgress.XgressDialer, error) {
	return nil, fmt.Errorf("not implemented")
}

type factory struct {
	id      *identity.TokenId
	ctrl    xgress.CtrlChannel
	options *xgress.Options
}

