// Code generated by go-swagger; DO NOT EDIT.

//
// Copyright NetFoundry, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// __          __              _
// \ \        / /             (_)
//  \ \  /\  / /_ _ _ __ _ __  _ _ __   __ _
//   \ \/  \/ / _` | '__| '_ \| | '_ \ / _` |
//    \  /\  / (_| | |  | | | | | | | | (_| | : This file is generated, do not edit it.
//     \/  \/ \__,_|_|  |_| |_|_|_| |_|\__, |
//                                      __/ |
//                                     |___/

package rest_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostureCheckDomainUpdate posture check domain update
//
// swagger:model PostureCheckDomainUpdate
type PostureCheckDomainUpdate struct {
	descriptionField *string

	nameField *string

	roleAttributesField Attributes

	tagsField Tags

	// domains
	// Required: true
	// Min Items: 1
	Domains []string `json:"domains"`
}

// Description gets the description of this subtype
func (m *PostureCheckDomainUpdate) Description() *string {
	return m.descriptionField
}

// SetDescription sets the description of this subtype
func (m *PostureCheckDomainUpdate) SetDescription(val *string) {
	m.descriptionField = val
}

// Name gets the name of this subtype
func (m *PostureCheckDomainUpdate) Name() *string {
	return m.nameField
}

// SetName sets the name of this subtype
func (m *PostureCheckDomainUpdate) SetName(val *string) {
	m.nameField = val
}

// RoleAttributes gets the role attributes of this subtype
func (m *PostureCheckDomainUpdate) RoleAttributes() Attributes {
	return m.roleAttributesField
}

// SetRoleAttributes sets the role attributes of this subtype
func (m *PostureCheckDomainUpdate) SetRoleAttributes(val Attributes) {
	m.roleAttributesField = val
}

// Tags gets the tags of this subtype
func (m *PostureCheckDomainUpdate) Tags() Tags {
	return m.tagsField
}

// SetTags sets the tags of this subtype
func (m *PostureCheckDomainUpdate) SetTags(val Tags) {
	m.tagsField = val
}

// TypeID gets the type Id of this subtype
func (m *PostureCheckDomainUpdate) TypeID() PostureCheckType {
	return "DOMAIN"
}

// SetTypeID sets the type Id of this subtype
func (m *PostureCheckDomainUpdate) SetTypeID(val PostureCheckType) {
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *PostureCheckDomainUpdate) UnmarshalJSON(raw []byte) error {
	var data struct {

		// domains
		// Required: true
		// Min Items: 1
		Domains []string `json:"domains"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		Description *string `json:"description"`

		Name *string `json:"name"`

		RoleAttributes Attributes `json:"roleAttributes"`

		Tags Tags `json:"tags"`

		TypeID PostureCheckType `json:"typeId,omitempty"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result PostureCheckDomainUpdate

	result.descriptionField = base.Description

	result.nameField = base.Name

	result.roleAttributesField = base.RoleAttributes

	result.tagsField = base.Tags

	if base.TypeID != result.TypeID() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid typeId value: %q", base.TypeID)
	}

	result.Domains = data.Domains

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m PostureCheckDomainUpdate) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// domains
		// Required: true
		// Min Items: 1
		Domains []string `json:"domains"`
	}{

		Domains: m.Domains,
	})
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		Description *string `json:"description"`

		Name *string `json:"name"`

		RoleAttributes Attributes `json:"roleAttributes"`

		Tags Tags `json:"tags"`

		TypeID PostureCheckType `json:"typeId,omitempty"`
	}{

		Description: m.Description(),

		Name: m.Name(),

		RoleAttributes: m.RoleAttributes(),

		Tags: m.Tags(),

		TypeID: m.TypeID(),
	})
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this posture check domain update
func (m *PostureCheckDomainUpdate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoleAttributes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDomains(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostureCheckDomainUpdate) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description()); err != nil {
		return err
	}

	return nil
}

func (m *PostureCheckDomainUpdate) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name()); err != nil {
		return err
	}

	return nil
}

func (m *PostureCheckDomainUpdate) validateRoleAttributes(formats strfmt.Registry) error {

	if swag.IsZero(m.RoleAttributes()) { // not required
		return nil
	}

	if err := m.RoleAttributes().Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("roleAttributes")
		}
		return err
	}

	return nil
}

func (m *PostureCheckDomainUpdate) validateTags(formats strfmt.Registry) error {

	if swag.IsZero(m.Tags()) { // not required
		return nil
	}

	if err := m.Tags().Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("tags")
		}
		return err
	}

	return nil
}

func (m *PostureCheckDomainUpdate) validateDomains(formats strfmt.Registry) error {

	if err := validate.Required("domains", "body", m.Domains); err != nil {
		return err
	}

	iDomainsSize := int64(len(m.Domains))

	if err := validate.MinItems("domains", "body", iDomainsSize, 1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostureCheckDomainUpdate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostureCheckDomainUpdate) UnmarshalBinary(b []byte) error {
	var res PostureCheckDomainUpdate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
