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

package posture_checks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openziti/edge/rest_model"
)

// CreatePostureResponseBulkOKCode is the HTTP code returned for type CreatePostureResponseBulkOK
const CreatePostureResponseBulkOKCode int = 200

/*CreatePostureResponseBulkOK Contains a list of services that have had their timers altered

swagger:response createPostureResponseBulkOK
*/
type CreatePostureResponseBulkOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.PostureResponseEnvelope `json:"body,omitempty"`
}

// NewCreatePostureResponseBulkOK creates CreatePostureResponseBulkOK with default headers values
func NewCreatePostureResponseBulkOK() *CreatePostureResponseBulkOK {

	return &CreatePostureResponseBulkOK{}
}

// WithPayload adds the payload to the create posture response bulk o k response
func (o *CreatePostureResponseBulkOK) WithPayload(payload *rest_model.PostureResponseEnvelope) *CreatePostureResponseBulkOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create posture response bulk o k response
func (o *CreatePostureResponseBulkOK) SetPayload(payload *rest_model.PostureResponseEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreatePostureResponseBulkOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreatePostureResponseBulkBadRequestCode is the HTTP code returned for type CreatePostureResponseBulkBadRequest
const CreatePostureResponseBulkBadRequestCode int = 400

/*CreatePostureResponseBulkBadRequest The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information

swagger:response createPostureResponseBulkBadRequest
*/
type CreatePostureResponseBulkBadRequest struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewCreatePostureResponseBulkBadRequest creates CreatePostureResponseBulkBadRequest with default headers values
func NewCreatePostureResponseBulkBadRequest() *CreatePostureResponseBulkBadRequest {

	return &CreatePostureResponseBulkBadRequest{}
}

// WithPayload adds the payload to the create posture response bulk bad request response
func (o *CreatePostureResponseBulkBadRequest) WithPayload(payload *rest_model.APIErrorEnvelope) *CreatePostureResponseBulkBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create posture response bulk bad request response
func (o *CreatePostureResponseBulkBadRequest) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreatePostureResponseBulkBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreatePostureResponseBulkUnauthorizedCode is the HTTP code returned for type CreatePostureResponseBulkUnauthorized
const CreatePostureResponseBulkUnauthorizedCode int = 401

/*CreatePostureResponseBulkUnauthorized The currently supplied session does not have the correct access rights to request this resource

swagger:response createPostureResponseBulkUnauthorized
*/
type CreatePostureResponseBulkUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewCreatePostureResponseBulkUnauthorized creates CreatePostureResponseBulkUnauthorized with default headers values
func NewCreatePostureResponseBulkUnauthorized() *CreatePostureResponseBulkUnauthorized {

	return &CreatePostureResponseBulkUnauthorized{}
}

// WithPayload adds the payload to the create posture response bulk unauthorized response
func (o *CreatePostureResponseBulkUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *CreatePostureResponseBulkUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create posture response bulk unauthorized response
func (o *CreatePostureResponseBulkUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreatePostureResponseBulkUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
