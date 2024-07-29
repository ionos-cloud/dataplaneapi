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

package http_error_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v6/models"
)

// ReplaceHTTPErrorRuleDefaultsOKCode is the HTTP code returned for type ReplaceHTTPErrorRuleDefaultsOK
const ReplaceHTTPErrorRuleDefaultsOKCode int = 200

/*
ReplaceHTTPErrorRuleDefaultsOK HTTP Error Rule replaced

swagger:response replaceHttpErrorRuleDefaultsOK
*/
type ReplaceHTTPErrorRuleDefaultsOK struct {

	/*
	  In: Body
	*/
	Payload *models.HTTPErrorRule `json:"body,omitempty"`
}

// NewReplaceHTTPErrorRuleDefaultsOK creates ReplaceHTTPErrorRuleDefaultsOK with default headers values
func NewReplaceHTTPErrorRuleDefaultsOK() *ReplaceHTTPErrorRuleDefaultsOK {

	return &ReplaceHTTPErrorRuleDefaultsOK{}
}

// WithPayload adds the payload to the replace Http error rule defaults o k response
func (o *ReplaceHTTPErrorRuleDefaultsOK) WithPayload(payload *models.HTTPErrorRule) *ReplaceHTTPErrorRuleDefaultsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace Http error rule defaults o k response
func (o *ReplaceHTTPErrorRuleDefaultsOK) SetPayload(payload *models.HTTPErrorRule) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceHTTPErrorRuleDefaultsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceHTTPErrorRuleDefaultsAcceptedCode is the HTTP code returned for type ReplaceHTTPErrorRuleDefaultsAccepted
const ReplaceHTTPErrorRuleDefaultsAcceptedCode int = 202

/*
ReplaceHTTPErrorRuleDefaultsAccepted Configuration change accepted and reload requested

swagger:response replaceHttpErrorRuleDefaultsAccepted
*/
type ReplaceHTTPErrorRuleDefaultsAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`

	/*
	  In: Body
	*/
	Payload *models.HTTPErrorRule `json:"body,omitempty"`
}

// NewReplaceHTTPErrorRuleDefaultsAccepted creates ReplaceHTTPErrorRuleDefaultsAccepted with default headers values
func NewReplaceHTTPErrorRuleDefaultsAccepted() *ReplaceHTTPErrorRuleDefaultsAccepted {

	return &ReplaceHTTPErrorRuleDefaultsAccepted{}
}

// WithReloadID adds the reloadId to the replace Http error rule defaults accepted response
func (o *ReplaceHTTPErrorRuleDefaultsAccepted) WithReloadID(reloadID string) *ReplaceHTTPErrorRuleDefaultsAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the replace Http error rule defaults accepted response
func (o *ReplaceHTTPErrorRuleDefaultsAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WithPayload adds the payload to the replace Http error rule defaults accepted response
func (o *ReplaceHTTPErrorRuleDefaultsAccepted) WithPayload(payload *models.HTTPErrorRule) *ReplaceHTTPErrorRuleDefaultsAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace Http error rule defaults accepted response
func (o *ReplaceHTTPErrorRuleDefaultsAccepted) SetPayload(payload *models.HTTPErrorRule) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceHTTPErrorRuleDefaultsAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// ReplaceHTTPErrorRuleDefaultsBadRequestCode is the HTTP code returned for type ReplaceHTTPErrorRuleDefaultsBadRequest
const ReplaceHTTPErrorRuleDefaultsBadRequestCode int = 400

/*
ReplaceHTTPErrorRuleDefaultsBadRequest Bad request

swagger:response replaceHttpErrorRuleDefaultsBadRequest
*/
type ReplaceHTTPErrorRuleDefaultsBadRequest struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceHTTPErrorRuleDefaultsBadRequest creates ReplaceHTTPErrorRuleDefaultsBadRequest with default headers values
func NewReplaceHTTPErrorRuleDefaultsBadRequest() *ReplaceHTTPErrorRuleDefaultsBadRequest {

	return &ReplaceHTTPErrorRuleDefaultsBadRequest{}
}

// WithConfigurationVersion adds the configurationVersion to the replace Http error rule defaults bad request response
func (o *ReplaceHTTPErrorRuleDefaultsBadRequest) WithConfigurationVersion(configurationVersion string) *ReplaceHTTPErrorRuleDefaultsBadRequest {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the replace Http error rule defaults bad request response
func (o *ReplaceHTTPErrorRuleDefaultsBadRequest) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the replace Http error rule defaults bad request response
func (o *ReplaceHTTPErrorRuleDefaultsBadRequest) WithPayload(payload *models.Error) *ReplaceHTTPErrorRuleDefaultsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace Http error rule defaults bad request response
func (o *ReplaceHTTPErrorRuleDefaultsBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceHTTPErrorRuleDefaultsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := o.ConfigurationVersion
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

// ReplaceHTTPErrorRuleDefaultsNotFoundCode is the HTTP code returned for type ReplaceHTTPErrorRuleDefaultsNotFound
const ReplaceHTTPErrorRuleDefaultsNotFoundCode int = 404

/*
ReplaceHTTPErrorRuleDefaultsNotFound The specified resource was not found

swagger:response replaceHttpErrorRuleDefaultsNotFound
*/
type ReplaceHTTPErrorRuleDefaultsNotFound struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceHTTPErrorRuleDefaultsNotFound creates ReplaceHTTPErrorRuleDefaultsNotFound with default headers values
func NewReplaceHTTPErrorRuleDefaultsNotFound() *ReplaceHTTPErrorRuleDefaultsNotFound {

	return &ReplaceHTTPErrorRuleDefaultsNotFound{}
}

// WithConfigurationVersion adds the configurationVersion to the replace Http error rule defaults not found response
func (o *ReplaceHTTPErrorRuleDefaultsNotFound) WithConfigurationVersion(configurationVersion string) *ReplaceHTTPErrorRuleDefaultsNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the replace Http error rule defaults not found response
func (o *ReplaceHTTPErrorRuleDefaultsNotFound) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the replace Http error rule defaults not found response
func (o *ReplaceHTTPErrorRuleDefaultsNotFound) WithPayload(payload *models.Error) *ReplaceHTTPErrorRuleDefaultsNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace Http error rule defaults not found response
func (o *ReplaceHTTPErrorRuleDefaultsNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceHTTPErrorRuleDefaultsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := o.ConfigurationVersion
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

/*
ReplaceHTTPErrorRuleDefaultsDefault General Error

swagger:response replaceHttpErrorRuleDefaultsDefault
*/
type ReplaceHTTPErrorRuleDefaultsDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceHTTPErrorRuleDefaultsDefault creates ReplaceHTTPErrorRuleDefaultsDefault with default headers values
func NewReplaceHTTPErrorRuleDefaultsDefault(code int) *ReplaceHTTPErrorRuleDefaultsDefault {
	if code <= 0 {
		code = 500
	}

	return &ReplaceHTTPErrorRuleDefaultsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the replace HTTP error rule defaults default response
func (o *ReplaceHTTPErrorRuleDefaultsDefault) WithStatusCode(code int) *ReplaceHTTPErrorRuleDefaultsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the replace HTTP error rule defaults default response
func (o *ReplaceHTTPErrorRuleDefaultsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the replace HTTP error rule defaults default response
func (o *ReplaceHTTPErrorRuleDefaultsDefault) WithConfigurationVersion(configurationVersion string) *ReplaceHTTPErrorRuleDefaultsDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the replace HTTP error rule defaults default response
func (o *ReplaceHTTPErrorRuleDefaultsDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the replace HTTP error rule defaults default response
func (o *ReplaceHTTPErrorRuleDefaultsDefault) WithPayload(payload *models.Error) *ReplaceHTTPErrorRuleDefaultsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace HTTP error rule defaults default response
func (o *ReplaceHTTPErrorRuleDefaultsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceHTTPErrorRuleDefaultsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := o.ConfigurationVersion
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