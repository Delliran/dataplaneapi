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

package tcp_check

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ReplaceTCPCheckHandlerFunc turns a function with the right signature into a replace TCP check handler
type ReplaceTCPCheckHandlerFunc func(ReplaceTCPCheckParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn ReplaceTCPCheckHandlerFunc) Handle(params ReplaceTCPCheckParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// ReplaceTCPCheckHandler interface for that can handle valid replace TCP check params
type ReplaceTCPCheckHandler interface {
	Handle(ReplaceTCPCheckParams, interface{}) middleware.Responder
}

// NewReplaceTCPCheck creates a new http.Handler for the replace TCP check operation
func NewReplaceTCPCheck(ctx *middleware.Context, handler ReplaceTCPCheckHandler) *ReplaceTCPCheck {
	return &ReplaceTCPCheck{Context: ctx, Handler: handler}
}

/*ReplaceTCPCheck swagger:route PUT /services/haproxy/configuration/tcp_checks/{index} TCPCheck replaceTcpCheck

Replace a TCP check

Replaces a TCP Check configuration by it's index in the specified parent.

*/
type ReplaceTCPCheck struct {
	Context *middleware.Context
	Handler ReplaceTCPCheckHandler
}

func (o *ReplaceTCPCheck) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReplaceTCPCheckParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
