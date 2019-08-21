// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2019 HAProxy Technologies
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package discovery

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	"github.com/haproxytech/models"
)

// GetConfigurationEndpointsOKCode is the HTTP code returned for type GetConfigurationEndpointsOK
const GetConfigurationEndpointsOKCode int = 200

/*GetConfigurationEndpointsOK Success

swagger:response getConfigurationEndpointsOK
*/
type GetConfigurationEndpointsOK struct {

	/*
	  In: Body
	*/
	Payload models.Endpoints `json:"body,omitempty"`
}

// NewGetConfigurationEndpointsOK creates GetConfigurationEndpointsOK with default headers values
func NewGetConfigurationEndpointsOK() *GetConfigurationEndpointsOK {

	return &GetConfigurationEndpointsOK{}
}

// WithPayload adds the payload to the get configuration endpoints o k response
func (o *GetConfigurationEndpointsOK) WithPayload(payload models.Endpoints) *GetConfigurationEndpointsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get configuration endpoints o k response
func (o *GetConfigurationEndpointsOK) SetPayload(payload models.Endpoints) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConfigurationEndpointsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.Endpoints{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*GetConfigurationEndpointsDefault General Error

swagger:response getConfigurationEndpointsDefault
*/
type GetConfigurationEndpointsDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion int64 `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetConfigurationEndpointsDefault creates GetConfigurationEndpointsDefault with default headers values
func NewGetConfigurationEndpointsDefault(code int) *GetConfigurationEndpointsDefault {
	if code <= 0 {
		code = 500
	}

	return &GetConfigurationEndpointsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get configuration endpoints default response
func (o *GetConfigurationEndpointsDefault) WithStatusCode(code int) *GetConfigurationEndpointsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get configuration endpoints default response
func (o *GetConfigurationEndpointsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the get configuration endpoints default response
func (o *GetConfigurationEndpointsDefault) WithConfigurationVersion(configurationVersion int64) *GetConfigurationEndpointsDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get configuration endpoints default response
func (o *GetConfigurationEndpointsDefault) SetConfigurationVersion(configurationVersion int64) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get configuration endpoints default response
func (o *GetConfigurationEndpointsDefault) WithPayload(payload *models.Error) *GetConfigurationEndpointsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get configuration endpoints default response
func (o *GetConfigurationEndpointsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConfigurationEndpointsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := swag.FormatInt64(o.ConfigurationVersion)
	if configurationVersion != "" {
		rw.Header().Set("Configuration-Version", configurationVersion)
	}

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
