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

package http_response_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	"github.com/haproxytech/models"
)

// DeleteHTTPResponseRuleAcceptedCode is the HTTP code returned for type DeleteHTTPResponseRuleAccepted
const DeleteHTTPResponseRuleAcceptedCode int = 202

/*DeleteHTTPResponseRuleAccepted Configuration change accepted and reload requested

swagger:response deleteHttpResponseRuleAccepted
*/
type DeleteHTTPResponseRuleAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`
}

// NewDeleteHTTPResponseRuleAccepted creates DeleteHTTPResponseRuleAccepted with default headers values
func NewDeleteHTTPResponseRuleAccepted() *DeleteHTTPResponseRuleAccepted {

	return &DeleteHTTPResponseRuleAccepted{}
}

// WithReloadID adds the reloadId to the delete Http response rule accepted response
func (o *DeleteHTTPResponseRuleAccepted) WithReloadID(reloadID string) *DeleteHTTPResponseRuleAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the delete Http response rule accepted response
func (o *DeleteHTTPResponseRuleAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WriteResponse to the client
func (o *DeleteHTTPResponseRuleAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Reload-ID

	reloadID := o.ReloadID
	if reloadID != "" {
		rw.Header().Set("Reload-ID", reloadID)
	}

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(202)
}

// DeleteHTTPResponseRuleNoContentCode is the HTTP code returned for type DeleteHTTPResponseRuleNoContent
const DeleteHTTPResponseRuleNoContentCode int = 204

/*DeleteHTTPResponseRuleNoContent HTTP Response Rule deleted

swagger:response deleteHttpResponseRuleNoContent
*/
type DeleteHTTPResponseRuleNoContent struct {
}

// NewDeleteHTTPResponseRuleNoContent creates DeleteHTTPResponseRuleNoContent with default headers values
func NewDeleteHTTPResponseRuleNoContent() *DeleteHTTPResponseRuleNoContent {

	return &DeleteHTTPResponseRuleNoContent{}
}

// WriteResponse to the client
func (o *DeleteHTTPResponseRuleNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteHTTPResponseRuleNotFoundCode is the HTTP code returned for type DeleteHTTPResponseRuleNotFound
const DeleteHTTPResponseRuleNotFoundCode int = 404

/*DeleteHTTPResponseRuleNotFound The specified resource was not found

swagger:response deleteHttpResponseRuleNotFound
*/
type DeleteHTTPResponseRuleNotFound struct {
	/*Configuration file version

	 */
	ConfigurationVersion int64 `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteHTTPResponseRuleNotFound creates DeleteHTTPResponseRuleNotFound with default headers values
func NewDeleteHTTPResponseRuleNotFound() *DeleteHTTPResponseRuleNotFound {

	return &DeleteHTTPResponseRuleNotFound{}
}

// WithConfigurationVersion adds the configurationVersion to the delete Http response rule not found response
func (o *DeleteHTTPResponseRuleNotFound) WithConfigurationVersion(configurationVersion int64) *DeleteHTTPResponseRuleNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the delete Http response rule not found response
func (o *DeleteHTTPResponseRuleNotFound) SetConfigurationVersion(configurationVersion int64) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the delete Http response rule not found response
func (o *DeleteHTTPResponseRuleNotFound) WithPayload(payload *models.Error) *DeleteHTTPResponseRuleNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete Http response rule not found response
func (o *DeleteHTTPResponseRuleNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteHTTPResponseRuleNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

/*DeleteHTTPResponseRuleDefault General Error

swagger:response deleteHttpResponseRuleDefault
*/
type DeleteHTTPResponseRuleDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion int64 `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteHTTPResponseRuleDefault creates DeleteHTTPResponseRuleDefault with default headers values
func NewDeleteHTTPResponseRuleDefault(code int) *DeleteHTTPResponseRuleDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteHTTPResponseRuleDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete HTTP response rule default response
func (o *DeleteHTTPResponseRuleDefault) WithStatusCode(code int) *DeleteHTTPResponseRuleDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete HTTP response rule default response
func (o *DeleteHTTPResponseRuleDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the delete HTTP response rule default response
func (o *DeleteHTTPResponseRuleDefault) WithConfigurationVersion(configurationVersion int64) *DeleteHTTPResponseRuleDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the delete HTTP response rule default response
func (o *DeleteHTTPResponseRuleDefault) SetConfigurationVersion(configurationVersion int64) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the delete HTTP response rule default response
func (o *DeleteHTTPResponseRuleDefault) WithPayload(payload *models.Error) *DeleteHTTPResponseRuleDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete HTTP response rule default response
func (o *DeleteHTTPResponseRuleDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteHTTPResponseRuleDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
