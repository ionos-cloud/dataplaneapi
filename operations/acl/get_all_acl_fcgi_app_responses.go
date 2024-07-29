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

package acl

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v6/models"
)

// GetAllACLFCGIAppOKCode is the HTTP code returned for type GetAllACLFCGIAppOK
const GetAllACLFCGIAppOKCode int = 200

/*
GetAllACLFCGIAppOK Successful operation

swagger:response getAllAclFcgiAppOK
*/
type GetAllACLFCGIAppOK struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload models.Acls `json:"body,omitempty"`
}

// NewGetAllACLFCGIAppOK creates GetAllACLFCGIAppOK with default headers values
func NewGetAllACLFCGIAppOK() *GetAllACLFCGIAppOK {

	return &GetAllACLFCGIAppOK{}
}

// WithConfigurationVersion adds the configurationVersion to the get all Acl Fcgi app o k response
func (o *GetAllACLFCGIAppOK) WithConfigurationVersion(configurationVersion string) *GetAllACLFCGIAppOK {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get all Acl Fcgi app o k response
func (o *GetAllACLFCGIAppOK) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get all Acl Fcgi app o k response
func (o *GetAllACLFCGIAppOK) WithPayload(payload models.Acls) *GetAllACLFCGIAppOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get all Acl Fcgi app o k response
func (o *GetAllACLFCGIAppOK) SetPayload(payload models.Acls) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAllACLFCGIAppOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := o.ConfigurationVersion
	if configurationVersion != "" {
		rw.Header().Set("Configuration-Version", configurationVersion)
	}

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.Acls{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*
GetAllACLFCGIAppDefault General Error

swagger:response getAllAclFcgiAppDefault
*/
type GetAllACLFCGIAppDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAllACLFCGIAppDefault creates GetAllACLFCGIAppDefault with default headers values
func NewGetAllACLFCGIAppDefault(code int) *GetAllACLFCGIAppDefault {
	if code <= 0 {
		code = 500
	}

	return &GetAllACLFCGIAppDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get all Acl FCGI app default response
func (o *GetAllACLFCGIAppDefault) WithStatusCode(code int) *GetAllACLFCGIAppDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get all Acl FCGI app default response
func (o *GetAllACLFCGIAppDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the get all Acl FCGI app default response
func (o *GetAllACLFCGIAppDefault) WithConfigurationVersion(configurationVersion string) *GetAllACLFCGIAppDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get all Acl FCGI app default response
func (o *GetAllACLFCGIAppDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get all Acl FCGI app default response
func (o *GetAllACLFCGIAppDefault) WithPayload(payload *models.Error) *GetAllACLFCGIAppDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get all Acl FCGI app default response
func (o *GetAllACLFCGIAppDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAllACLFCGIAppDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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