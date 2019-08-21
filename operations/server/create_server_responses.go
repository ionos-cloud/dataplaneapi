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

package server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	"github.com/haproxytech/models"
)

// CreateServerCreatedCode is the HTTP code returned for type CreateServerCreated
const CreateServerCreatedCode int = 201

/*CreateServerCreated Server created

swagger:response createServerCreated
*/
type CreateServerCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Server `json:"body,omitempty"`
}

// NewCreateServerCreated creates CreateServerCreated with default headers values
func NewCreateServerCreated() *CreateServerCreated {

	return &CreateServerCreated{}
}

// WithPayload adds the payload to the create server created response
func (o *CreateServerCreated) WithPayload(payload *models.Server) *CreateServerCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create server created response
func (o *CreateServerCreated) SetPayload(payload *models.Server) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateServerCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateServerAcceptedCode is the HTTP code returned for type CreateServerAccepted
const CreateServerAcceptedCode int = 202

/*CreateServerAccepted Configuration change accepted and reload requested

swagger:response createServerAccepted
*/
type CreateServerAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`

	/*
	  In: Body
	*/
	Payload *models.Server `json:"body,omitempty"`
}

// NewCreateServerAccepted creates CreateServerAccepted with default headers values
func NewCreateServerAccepted() *CreateServerAccepted {

	return &CreateServerAccepted{}
}

// WithReloadID adds the reloadId to the create server accepted response
func (o *CreateServerAccepted) WithReloadID(reloadID string) *CreateServerAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the create server accepted response
func (o *CreateServerAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WithPayload adds the payload to the create server accepted response
func (o *CreateServerAccepted) WithPayload(payload *models.Server) *CreateServerAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create server accepted response
func (o *CreateServerAccepted) SetPayload(payload *models.Server) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateServerAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// CreateServerBadRequestCode is the HTTP code returned for type CreateServerBadRequest
const CreateServerBadRequestCode int = 400

/*CreateServerBadRequest Bad request

swagger:response createServerBadRequest
*/
type CreateServerBadRequest struct {
	/*Configuration file version

	 */
	ConfigurationVersion int64 `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateServerBadRequest creates CreateServerBadRequest with default headers values
func NewCreateServerBadRequest() *CreateServerBadRequest {

	return &CreateServerBadRequest{}
}

// WithConfigurationVersion adds the configurationVersion to the create server bad request response
func (o *CreateServerBadRequest) WithConfigurationVersion(configurationVersion int64) *CreateServerBadRequest {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create server bad request response
func (o *CreateServerBadRequest) SetConfigurationVersion(configurationVersion int64) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create server bad request response
func (o *CreateServerBadRequest) WithPayload(payload *models.Error) *CreateServerBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create server bad request response
func (o *CreateServerBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateServerBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// CreateServerConflictCode is the HTTP code returned for type CreateServerConflict
const CreateServerConflictCode int = 409

/*CreateServerConflict The specified resource already exists

swagger:response createServerConflict
*/
type CreateServerConflict struct {
	/*Configuration file version

	 */
	ConfigurationVersion int64 `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateServerConflict creates CreateServerConflict with default headers values
func NewCreateServerConflict() *CreateServerConflict {

	return &CreateServerConflict{}
}

// WithConfigurationVersion adds the configurationVersion to the create server conflict response
func (o *CreateServerConflict) WithConfigurationVersion(configurationVersion int64) *CreateServerConflict {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create server conflict response
func (o *CreateServerConflict) SetConfigurationVersion(configurationVersion int64) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create server conflict response
func (o *CreateServerConflict) WithPayload(payload *models.Error) *CreateServerConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create server conflict response
func (o *CreateServerConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateServerConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

/*CreateServerDefault General Error

swagger:response createServerDefault
*/
type CreateServerDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion int64 `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateServerDefault creates CreateServerDefault with default headers values
func NewCreateServerDefault(code int) *CreateServerDefault {
	if code <= 0 {
		code = 500
	}

	return &CreateServerDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create server default response
func (o *CreateServerDefault) WithStatusCode(code int) *CreateServerDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create server default response
func (o *CreateServerDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the create server default response
func (o *CreateServerDefault) WithConfigurationVersion(configurationVersion int64) *CreateServerDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create server default response
func (o *CreateServerDefault) SetConfigurationVersion(configurationVersion int64) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create server default response
func (o *CreateServerDefault) WithPayload(payload *models.Error) *CreateServerDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create server default response
func (o *CreateServerDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateServerDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
