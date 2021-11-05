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

package terminator

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openziti/fabric/rest_model"
)

// ListTerminatorsReader is a Reader for the ListTerminators structure.
type ListTerminatorsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListTerminatorsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListTerminatorsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListTerminatorsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListTerminatorsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListTerminatorsOK creates a ListTerminatorsOK with default headers values
func NewListTerminatorsOK() *ListTerminatorsOK {
	return &ListTerminatorsOK{}
}

/* ListTerminatorsOK describes a response with status code 200, with default header values.

A list of terminators
*/
type ListTerminatorsOK struct {
	Payload *rest_model.ListTerminatorsEnvelope
}

func (o *ListTerminatorsOK) Error() string {
	return fmt.Sprintf("[GET /terminators][%d] listTerminatorsOK  %+v", 200, o.Payload)
}
func (o *ListTerminatorsOK) GetPayload() *rest_model.ListTerminatorsEnvelope {
	return o.Payload
}

func (o *ListTerminatorsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.ListTerminatorsEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListTerminatorsBadRequest creates a ListTerminatorsBadRequest with default headers values
func NewListTerminatorsBadRequest() *ListTerminatorsBadRequest {
	return &ListTerminatorsBadRequest{}
}

/* ListTerminatorsBadRequest describes a response with status code 400, with default header values.

The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information
*/
type ListTerminatorsBadRequest struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *ListTerminatorsBadRequest) Error() string {
	return fmt.Sprintf("[GET /terminators][%d] listTerminatorsBadRequest  %+v", 400, o.Payload)
}
func (o *ListTerminatorsBadRequest) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *ListTerminatorsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListTerminatorsUnauthorized creates a ListTerminatorsUnauthorized with default headers values
func NewListTerminatorsUnauthorized() *ListTerminatorsUnauthorized {
	return &ListTerminatorsUnauthorized{}
}

/* ListTerminatorsUnauthorized describes a response with status code 401, with default header values.

The currently supplied session does not have the correct access rights to request this resource
*/
type ListTerminatorsUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *ListTerminatorsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /terminators][%d] listTerminatorsUnauthorized  %+v", 401, o.Payload)
}
func (o *ListTerminatorsUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *ListTerminatorsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
