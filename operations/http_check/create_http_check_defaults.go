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

package http_check

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CreateHTTPCheckDefaultsHandlerFunc turns a function with the right signature into a create HTTP check defaults handler
type CreateHTTPCheckDefaultsHandlerFunc func(CreateHTTPCheckDefaultsParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateHTTPCheckDefaultsHandlerFunc) Handle(params CreateHTTPCheckDefaultsParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// CreateHTTPCheckDefaultsHandler interface for that can handle valid create HTTP check defaults params
type CreateHTTPCheckDefaultsHandler interface {
	Handle(CreateHTTPCheckDefaultsParams, interface{}) middleware.Responder
}

// NewCreateHTTPCheckDefaults creates a new http.Handler for the create HTTP check defaults operation
func NewCreateHTTPCheckDefaults(ctx *middleware.Context, handler CreateHTTPCheckDefaultsHandler) *CreateHTTPCheckDefaults {
	return &CreateHTTPCheckDefaults{Context: ctx, Handler: handler}
}

/*
	CreateHTTPCheckDefaults swagger:route POST /services/haproxy/configuration/defaults/{parent_name}/http_checks/{index} HTTPCheck createHttpCheckDefaults

# Add a new HTTP check

Adds a new HTTP check of the specified type in the specified parent.
*/
type CreateHTTPCheckDefaults struct {
	Context *middleware.Context
	Handler CreateHTTPCheckDefaultsHandler
}

func (o *CreateHTTPCheckDefaults) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewCreateHTTPCheckDefaultsParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc.(interface{}) // this is really a interface{}, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}