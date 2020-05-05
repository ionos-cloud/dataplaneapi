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

package nameserver

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	"github.com/haproxytech/models/v2"
)

// CreateNameserverCreatedCode is the HTTP code returned for type CreateNameserverCreated
const CreateNameserverCreatedCode int = 201

/*CreateNameserverCreated Nameserver created

swagger:response createNameserverCreated
*/
type CreateNameserverCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Nameserver `json:"body,omitempty"`
}

// NewCreateNameserverCreated creates CreateNameserverCreated with default headers values
func NewCreateNameserverCreated() *CreateNameserverCreated {

	return &CreateNameserverCreated{}
}

// WithPayload adds the payload to the create nameserver created response
func (o *CreateNameserverCreated) WithPayload(payload *models.Nameserver) *CreateNameserverCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create nameserver created response
func (o *CreateNameserverCreated) SetPayload(payload *models.Nameserver) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateNameserverCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateNameserverAcceptedCode is the HTTP code returned for type CreateNameserverAccepted
const CreateNameserverAcceptedCode int = 202

/*CreateNameserverAccepted Configuration change accepted and reload requested

swagger:response createNameserverAccepted
*/
type CreateNameserverAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`

	/*
	  In: Body
	*/
	Payload *models.Nameserver `json:"body,omitempty"`
}

// NewCreateNameserverAccepted creates CreateNameserverAccepted with default headers values
func NewCreateNameserverAccepted() *CreateNameserverAccepted {

	return &CreateNameserverAccepted{}
}

// WithReloadID adds the reloadId to the create nameserver accepted response
func (o *CreateNameserverAccepted) WithReloadID(reloadID string) *CreateNameserverAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the create nameserver accepted response
func (o *CreateNameserverAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WithPayload adds the payload to the create nameserver accepted response
func (o *CreateNameserverAccepted) WithPayload(payload *models.Nameserver) *CreateNameserverAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create nameserver accepted response
func (o *CreateNameserverAccepted) SetPayload(payload *models.Nameserver) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateNameserverAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Reload-ID

	reloadID := o.ReloadID
	if reloadID != "" {
		rw.Header().Set("Reload-ID", reloadID)
	}

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateNameserverBadRequestCode is the HTTP code returned for type CreateNameserverBadRequest
const CreateNameserverBadRequestCode int = 400

/*CreateNameserverBadRequest Bad request

swagger:response createNameserverBadRequest
*/
type CreateNameserverBadRequest struct {
	/*Configuration file version

	  Default: 0
	*/
	ConfigurationVersion int64 `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateNameserverBadRequest creates CreateNameserverBadRequest with default headers values
func NewCreateNameserverBadRequest() *CreateNameserverBadRequest {

	var (
		// initialize headers with default values

		configurationVersionDefault = int64(0)
	)

	return &CreateNameserverBadRequest{

		ConfigurationVersion: configurationVersionDefault,
	}
}

// WithConfigurationVersion adds the configurationVersion to the create nameserver bad request response
func (o *CreateNameserverBadRequest) WithConfigurationVersion(configurationVersion int64) *CreateNameserverBadRequest {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create nameserver bad request response
func (o *CreateNameserverBadRequest) SetConfigurationVersion(configurationVersion int64) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create nameserver bad request response
func (o *CreateNameserverBadRequest) WithPayload(payload *models.Error) *CreateNameserverBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create nameserver bad request response
func (o *CreateNameserverBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateNameserverBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := swag.FormatInt64(o.ConfigurationVersion)
	if configurationVersion != "" {
		rw.Header().Set("Configuration-Version", configurationVersion)
	}

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateNameserverConflictCode is the HTTP code returned for type CreateNameserverConflict
const CreateNameserverConflictCode int = 409

/*CreateNameserverConflict The specified resource already exists

swagger:response createNameserverConflict
*/
type CreateNameserverConflict struct {
	/*Configuration file version

	  Default: 0
	*/
	ConfigurationVersion int64 `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateNameserverConflict creates CreateNameserverConflict with default headers values
func NewCreateNameserverConflict() *CreateNameserverConflict {

	var (
		// initialize headers with default values

		configurationVersionDefault = int64(0)
	)

	return &CreateNameserverConflict{

		ConfigurationVersion: configurationVersionDefault,
	}
}

// WithConfigurationVersion adds the configurationVersion to the create nameserver conflict response
func (o *CreateNameserverConflict) WithConfigurationVersion(configurationVersion int64) *CreateNameserverConflict {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create nameserver conflict response
func (o *CreateNameserverConflict) SetConfigurationVersion(configurationVersion int64) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create nameserver conflict response
func (o *CreateNameserverConflict) WithPayload(payload *models.Error) *CreateNameserverConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create nameserver conflict response
func (o *CreateNameserverConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateNameserverConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := swag.FormatInt64(o.ConfigurationVersion)
	if configurationVersion != "" {
		rw.Header().Set("Configuration-Version", configurationVersion)
	}

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*CreateNameserverDefault General Error

swagger:response createNameserverDefault
*/
type CreateNameserverDefault struct {
	_statusCode int
	/*Configuration file version

	  Default: 0
	*/
	ConfigurationVersion int64 `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateNameserverDefault creates CreateNameserverDefault with default headers values
func NewCreateNameserverDefault(code int) *CreateNameserverDefault {
	if code <= 0 {
		code = 500
	}

	var (
		// initialize headers with default values

		configurationVersionDefault = int64(0)
	)

	return &CreateNameserverDefault{
		_statusCode: code,

		ConfigurationVersion: configurationVersionDefault,
	}
}

// WithStatusCode adds the status to the create nameserver default response
func (o *CreateNameserverDefault) WithStatusCode(code int) *CreateNameserverDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create nameserver default response
func (o *CreateNameserverDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the create nameserver default response
func (o *CreateNameserverDefault) WithConfigurationVersion(configurationVersion int64) *CreateNameserverDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create nameserver default response
func (o *CreateNameserverDefault) SetConfigurationVersion(configurationVersion int64) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create nameserver default response
func (o *CreateNameserverDefault) WithPayload(payload *models.Error) *CreateNameserverDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create nameserver default response
func (o *CreateNameserverDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateNameserverDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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