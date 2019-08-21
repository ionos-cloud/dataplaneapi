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

package tcp_request_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	"github.com/haproxytech/models"
)

// ReplaceTCPRequestRuleOKCode is the HTTP code returned for type ReplaceTCPRequestRuleOK
const ReplaceTCPRequestRuleOKCode int = 200

/*ReplaceTCPRequestRuleOK TCP Request Rule replaced

swagger:response replaceTcpRequestRuleOK
*/
type ReplaceTCPRequestRuleOK struct {

	/*
	  In: Body
	*/
	Payload *models.TCPRequestRule `json:"body,omitempty"`
}

// NewReplaceTCPRequestRuleOK creates ReplaceTCPRequestRuleOK with default headers values
func NewReplaceTCPRequestRuleOK() *ReplaceTCPRequestRuleOK {

	return &ReplaceTCPRequestRuleOK{}
}

// WithPayload adds the payload to the replace Tcp request rule o k response
func (o *ReplaceTCPRequestRuleOK) WithPayload(payload *models.TCPRequestRule) *ReplaceTCPRequestRuleOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace Tcp request rule o k response
func (o *ReplaceTCPRequestRuleOK) SetPayload(payload *models.TCPRequestRule) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceTCPRequestRuleOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceTCPRequestRuleAcceptedCode is the HTTP code returned for type ReplaceTCPRequestRuleAccepted
const ReplaceTCPRequestRuleAcceptedCode int = 202

/*ReplaceTCPRequestRuleAccepted Configuration change accepted and reload requested

swagger:response replaceTcpRequestRuleAccepted
*/
type ReplaceTCPRequestRuleAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`

	/*
	  In: Body
	*/
	Payload *models.TCPRequestRule `json:"body,omitempty"`
}

// NewReplaceTCPRequestRuleAccepted creates ReplaceTCPRequestRuleAccepted with default headers values
func NewReplaceTCPRequestRuleAccepted() *ReplaceTCPRequestRuleAccepted {

	return &ReplaceTCPRequestRuleAccepted{}
}

// WithReloadID adds the reloadId to the replace Tcp request rule accepted response
func (o *ReplaceTCPRequestRuleAccepted) WithReloadID(reloadID string) *ReplaceTCPRequestRuleAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the replace Tcp request rule accepted response
func (o *ReplaceTCPRequestRuleAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WithPayload adds the payload to the replace Tcp request rule accepted response
func (o *ReplaceTCPRequestRuleAccepted) WithPayload(payload *models.TCPRequestRule) *ReplaceTCPRequestRuleAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace Tcp request rule accepted response
func (o *ReplaceTCPRequestRuleAccepted) SetPayload(payload *models.TCPRequestRule) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceTCPRequestRuleAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// ReplaceTCPRequestRuleBadRequestCode is the HTTP code returned for type ReplaceTCPRequestRuleBadRequest
const ReplaceTCPRequestRuleBadRequestCode int = 400

/*ReplaceTCPRequestRuleBadRequest Bad request

swagger:response replaceTcpRequestRuleBadRequest
*/
type ReplaceTCPRequestRuleBadRequest struct {
	/*Configuration file version

	 */
	ConfigurationVersion int64 `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceTCPRequestRuleBadRequest creates ReplaceTCPRequestRuleBadRequest with default headers values
func NewReplaceTCPRequestRuleBadRequest() *ReplaceTCPRequestRuleBadRequest {

	return &ReplaceTCPRequestRuleBadRequest{}
}

// WithConfigurationVersion adds the configurationVersion to the replace Tcp request rule bad request response
func (o *ReplaceTCPRequestRuleBadRequest) WithConfigurationVersion(configurationVersion int64) *ReplaceTCPRequestRuleBadRequest {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the replace Tcp request rule bad request response
func (o *ReplaceTCPRequestRuleBadRequest) SetConfigurationVersion(configurationVersion int64) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the replace Tcp request rule bad request response
func (o *ReplaceTCPRequestRuleBadRequest) WithPayload(payload *models.Error) *ReplaceTCPRequestRuleBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace Tcp request rule bad request response
func (o *ReplaceTCPRequestRuleBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceTCPRequestRuleBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// ReplaceTCPRequestRuleNotFoundCode is the HTTP code returned for type ReplaceTCPRequestRuleNotFound
const ReplaceTCPRequestRuleNotFoundCode int = 404

/*ReplaceTCPRequestRuleNotFound The specified resource was not found

swagger:response replaceTcpRequestRuleNotFound
*/
type ReplaceTCPRequestRuleNotFound struct {
	/*Configuration file version

	 */
	ConfigurationVersion int64 `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceTCPRequestRuleNotFound creates ReplaceTCPRequestRuleNotFound with default headers values
func NewReplaceTCPRequestRuleNotFound() *ReplaceTCPRequestRuleNotFound {

	return &ReplaceTCPRequestRuleNotFound{}
}

// WithConfigurationVersion adds the configurationVersion to the replace Tcp request rule not found response
func (o *ReplaceTCPRequestRuleNotFound) WithConfigurationVersion(configurationVersion int64) *ReplaceTCPRequestRuleNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the replace Tcp request rule not found response
func (o *ReplaceTCPRequestRuleNotFound) SetConfigurationVersion(configurationVersion int64) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the replace Tcp request rule not found response
func (o *ReplaceTCPRequestRuleNotFound) WithPayload(payload *models.Error) *ReplaceTCPRequestRuleNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace Tcp request rule not found response
func (o *ReplaceTCPRequestRuleNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceTCPRequestRuleNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

/*ReplaceTCPRequestRuleDefault General Error

swagger:response replaceTcpRequestRuleDefault
*/
type ReplaceTCPRequestRuleDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion int64 `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceTCPRequestRuleDefault creates ReplaceTCPRequestRuleDefault with default headers values
func NewReplaceTCPRequestRuleDefault(code int) *ReplaceTCPRequestRuleDefault {
	if code <= 0 {
		code = 500
	}

	return &ReplaceTCPRequestRuleDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the replace TCP request rule default response
func (o *ReplaceTCPRequestRuleDefault) WithStatusCode(code int) *ReplaceTCPRequestRuleDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the replace TCP request rule default response
func (o *ReplaceTCPRequestRuleDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the replace TCP request rule default response
func (o *ReplaceTCPRequestRuleDefault) WithConfigurationVersion(configurationVersion int64) *ReplaceTCPRequestRuleDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the replace TCP request rule default response
func (o *ReplaceTCPRequestRuleDefault) SetConfigurationVersion(configurationVersion int64) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the replace TCP request rule default response
func (o *ReplaceTCPRequestRuleDefault) WithPayload(payload *models.Error) *ReplaceTCPRequestRuleDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace TCP request rule default response
func (o *ReplaceTCPRequestRuleDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceTCPRequestRuleDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
