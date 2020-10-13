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

// PostureCheckProcessCreate posture check process create
//
// swagger:model PostureCheckProcessCreate
type PostureCheckProcessCreate struct {
	descriptionField *string

	nameField *string

	roleAttributesField Attributes

	tagsField Tags

	// process
	// Required: true
	Process *Process `json:"process"`
}

// Description gets the description of this subtype
func (m *PostureCheckProcessCreate) Description() *string {
	return m.descriptionField
}

// SetDescription sets the description of this subtype
func (m *PostureCheckProcessCreate) SetDescription(val *string) {
	m.descriptionField = val
}

// Name gets the name of this subtype
func (m *PostureCheckProcessCreate) Name() *string {
	return m.nameField
}

// SetName sets the name of this subtype
func (m *PostureCheckProcessCreate) SetName(val *string) {
	m.nameField = val
}

// RoleAttributes gets the role attributes of this subtype
func (m *PostureCheckProcessCreate) RoleAttributes() Attributes {
	return m.roleAttributesField
}

// SetRoleAttributes sets the role attributes of this subtype
func (m *PostureCheckProcessCreate) SetRoleAttributes(val Attributes) {
	m.roleAttributesField = val
}

// Tags gets the tags of this subtype
func (m *PostureCheckProcessCreate) Tags() Tags {
	return m.tagsField
}

// SetTags sets the tags of this subtype
func (m *PostureCheckProcessCreate) SetTags(val Tags) {
	m.tagsField = val
}

// TypeID gets the type Id of this subtype
func (m *PostureCheckProcessCreate) TypeID() PostureCheckType {
	return "PROCESS"
}

// SetTypeID sets the type Id of this subtype
func (m *PostureCheckProcessCreate) SetTypeID(val PostureCheckType) {
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *PostureCheckProcessCreate) UnmarshalJSON(raw []byte) error {
	var data struct {

		// process
		// Required: true
		Process *Process `json:"process"`
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

		TypeID PostureCheckType `json:"typeId"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result PostureCheckProcessCreate

	result.descriptionField = base.Description

	result.nameField = base.Name

	result.roleAttributesField = base.RoleAttributes

	result.tagsField = base.Tags

	if base.TypeID != result.TypeID() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid typeId value: %q", base.TypeID)
	}

	result.Process = data.Process

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m PostureCheckProcessCreate) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// process
		// Required: true
		Process *Process `json:"process"`
	}{

		Process: m.Process,
	})
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		Description *string `json:"description"`

		Name *string `json:"name"`

		RoleAttributes Attributes `json:"roleAttributes"`

		Tags Tags `json:"tags"`

		TypeID PostureCheckType `json:"typeId"`
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

// Validate validates this posture check process create
func (m *PostureCheckProcessCreate) Validate(formats strfmt.Registry) error {
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

	if err := m.validateProcess(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostureCheckProcessCreate) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description()); err != nil {
		return err
	}

	return nil
}

func (m *PostureCheckProcessCreate) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name()); err != nil {
		return err
	}

	return nil
}

func (m *PostureCheckProcessCreate) validateRoleAttributes(formats strfmt.Registry) error {

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

func (m *PostureCheckProcessCreate) validateTags(formats strfmt.Registry) error {

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

func (m *PostureCheckProcessCreate) validateProcess(formats strfmt.Registry) error {

	if err := validate.Required("process", "body", m.Process); err != nil {
		return err
	}

	if m.Process != nil {
		if err := m.Process.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("process")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostureCheckProcessCreate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostureCheckProcessCreate) UnmarshalBinary(b []byte) error {
	var res PostureCheckProcessCreate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
