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

// ReplaceAllHTTPCheckBackendHandlerFunc turns a function with the right signature into a replace all HTTP check backend handler
type ReplaceAllHTTPCheckBackendHandlerFunc func(ReplaceAllHTTPCheckBackendParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn ReplaceAllHTTPCheckBackendHandlerFunc) Handle(params ReplaceAllHTTPCheckBackendParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// ReplaceAllHTTPCheckBackendHandler interface for that can handle valid replace all HTTP check backend params
type ReplaceAllHTTPCheckBackendHandler interface {
	Handle(ReplaceAllHTTPCheckBackendParams, interface{}) middleware.Responder
}

// NewReplaceAllHTTPCheckBackend creates a new http.Handler for the replace all HTTP check backend operation
func NewReplaceAllHTTPCheckBackend(ctx *middleware.Context, handler ReplaceAllHTTPCheckBackendHandler) *ReplaceAllHTTPCheckBackend {
	return &ReplaceAllHTTPCheckBackend{Context: ctx, Handler: handler}
}

/*
	ReplaceAllHTTPCheckBackend swagger:route PUT /services/haproxy/configuration/backends/{parent_name}/http_checks HTTPCheck replaceAllHttpCheckBackend

# Replace an HTTP checks list

Replaces a whole list of HTTP checks with the list given in parameter
*/
type ReplaceAllHTTPCheckBackend struct {
	Context *middleware.Context
	Handler ReplaceAllHTTPCheckBackendHandler
}

func (o *ReplaceAllHTTPCheckBackend) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewReplaceAllHTTPCheckBackendParams()
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