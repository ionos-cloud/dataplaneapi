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

package backend

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	"github.com/haproxytech/models"
)

// GetBackendOKCode is the HTTP code returned for type GetBackendOK
const GetBackendOKCode int = 200

/*GetBackendOK Successful operation

swagger:response getBackendOK
*/
type GetBackendOK struct {
	/*Configuration file version

	 */
	ConfigurationVersion int64 `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *GetBackendOKBody `json:"body,omitempty"`
}

// NewGetBackendOK creates GetBackendOK with default headers values
func NewGetBackendOK() *GetBackendOK {

	return &GetBackendOK{}
}

// WithConfigurationVersion adds the configurationVersion to the get backend o k response
func (o *GetBackendOK) WithConfigurationVersion(configurationVersion int64) *GetBackendOK {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get backend o k response
func (o *GetBackendOK) SetConfigurationVersion(configurationVersion int64) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get backend o k response
func (o *GetBackendOK) WithPayload(payload *GetBackendOKBody) *GetBackendOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get backend o k response
func (o *GetBackendOK) SetPayload(payload *GetBackendOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBackendOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := swag.FormatInt64(o.ConfigurationVersion)
	if configurationVersion != "" {
		rw.Header().Set("Configuration-Version", configurationVersion)
	}

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetBackendNotFoundCode is the HTTP code returned for type GetBackendNotFound
const GetBackendNotFoundCode int = 404

/*GetBackendNotFound The specified resource was not found

swagger:response getBackendNotFound
*/
type GetBackendNotFound struct {
	/*Configuration file version

	 */
	ConfigurationVersion int64 `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetBackendNotFound creates GetBackendNotFound with default headers values
func NewGetBackendNotFound() *GetBackendNotFound {

	return &GetBackendNotFound{}
}

// WithConfigurationVersion adds the configurationVersion to the get backend not found response
func (o *GetBackendNotFound) WithConfigurationVersion(configurationVersion int64) *GetBackendNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get backend not found response
func (o *GetBackendNotFound) SetConfigurationVersion(configurationVersion int64) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get backend not found response
func (o *GetBackendNotFound) WithPayload(payload *models.Error) *GetBackendNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get backend not found response
func (o *GetBackendNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBackendNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := swag.FormatInt64(o.ConfigurationVersion)
	if configurationVersion != "" {
		rw.Header().Set("Configuration-Version", configurationVersion)
	}

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetBackendDefault General Error

swagger:response getBackendDefault
*/
type GetBackendDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion int64 `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetBackendDefault creates GetBackendDefault with default headers values
func NewGetBackendDefault(code int) *GetBackendDefault {
	if code <= 0 {
		code = 500
	}

	return &GetBackendDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get backend default response
func (o *GetBackendDefault) WithStatusCode(code int) *GetBackendDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get backend default response
func (o *GetBackendDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the get backend default response
func (o *GetBackendDefault) WithConfigurationVersion(configurationVersion int64) *GetBackendDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get backend default response
func (o *GetBackendDefault) SetConfigurationVersion(configurationVersion int64) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get backend default response
func (o *GetBackendDefault) WithPayload(payload *models.Error) *GetBackendDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get backend default response
func (o *GetBackendDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBackendDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
