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

package transactions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	"github.com/haproxytech/models"
)

// StartTransactionCreatedCode is the HTTP code returned for type StartTransactionCreated
const StartTransactionCreatedCode int = 201

/*StartTransactionCreated Transaction started

swagger:response startTransactionCreated
*/
type StartTransactionCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Transaction `json:"body,omitempty"`
}

// NewStartTransactionCreated creates StartTransactionCreated with default headers values
func NewStartTransactionCreated() *StartTransactionCreated {

	return &StartTransactionCreated{}
}

// WithPayload adds the payload to the start transaction created response
func (o *StartTransactionCreated) WithPayload(payload *models.Transaction) *StartTransactionCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the start transaction created response
func (o *StartTransactionCreated) SetPayload(payload *models.Transaction) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *StartTransactionCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*StartTransactionDefault General Error

swagger:response startTransactionDefault
*/
type StartTransactionDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion int64 `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewStartTransactionDefault creates StartTransactionDefault with default headers values
func NewStartTransactionDefault(code int) *StartTransactionDefault {
	if code <= 0 {
		code = 500
	}

	return &StartTransactionDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the start transaction default response
func (o *StartTransactionDefault) WithStatusCode(code int) *StartTransactionDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the start transaction default response
func (o *StartTransactionDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the start transaction default response
func (o *StartTransactionDefault) WithConfigurationVersion(configurationVersion int64) *StartTransactionDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the start transaction default response
func (o *StartTransactionDefault) SetConfigurationVersion(configurationVersion int64) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the start transaction default response
func (o *StartTransactionDefault) WithPayload(payload *models.Error) *StartTransactionDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the start transaction default response
func (o *StartTransactionDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *StartTransactionDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
