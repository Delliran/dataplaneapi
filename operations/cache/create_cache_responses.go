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

package cache

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v3/models"
)

// CreateCacheCreatedCode is the HTTP code returned for type CreateCacheCreated
const CreateCacheCreatedCode int = 201

/*CreateCacheCreated Cache created

swagger:response createCacheCreated
*/
type CreateCacheCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Cache `json:"body,omitempty"`
}

// NewCreateCacheCreated creates CreateCacheCreated with default headers values
func NewCreateCacheCreated() *CreateCacheCreated {

	return &CreateCacheCreated{}
}

// WithPayload adds the payload to the create cache created response
func (o *CreateCacheCreated) WithPayload(payload *models.Cache) *CreateCacheCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create cache created response
func (o *CreateCacheCreated) SetPayload(payload *models.Cache) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateCacheCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateCacheAcceptedCode is the HTTP code returned for type CreateCacheAccepted
const CreateCacheAcceptedCode int = 202

/*CreateCacheAccepted Configuration change accepted and reload requested

swagger:response createCacheAccepted
*/
type CreateCacheAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`

	/*
	  In: Body
	*/
	Payload *models.Cache `json:"body,omitempty"`
}

// NewCreateCacheAccepted creates CreateCacheAccepted with default headers values
func NewCreateCacheAccepted() *CreateCacheAccepted {

	return &CreateCacheAccepted{}
}

// WithReloadID adds the reloadId to the create cache accepted response
func (o *CreateCacheAccepted) WithReloadID(reloadID string) *CreateCacheAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the create cache accepted response
func (o *CreateCacheAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WithPayload adds the payload to the create cache accepted response
func (o *CreateCacheAccepted) WithPayload(payload *models.Cache) *CreateCacheAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create cache accepted response
func (o *CreateCacheAccepted) SetPayload(payload *models.Cache) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateCacheAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// CreateCacheBadRequestCode is the HTTP code returned for type CreateCacheBadRequest
const CreateCacheBadRequestCode int = 400

/*CreateCacheBadRequest Bad request

swagger:response createCacheBadRequest
*/
type CreateCacheBadRequest struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateCacheBadRequest creates CreateCacheBadRequest with default headers values
func NewCreateCacheBadRequest() *CreateCacheBadRequest {

	return &CreateCacheBadRequest{}
}

// WithConfigurationVersion adds the configurationVersion to the create cache bad request response
func (o *CreateCacheBadRequest) WithConfigurationVersion(configurationVersion string) *CreateCacheBadRequest {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create cache bad request response
func (o *CreateCacheBadRequest) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create cache bad request response
func (o *CreateCacheBadRequest) WithPayload(payload *models.Error) *CreateCacheBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create cache bad request response
func (o *CreateCacheBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateCacheBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// CreateCacheConflictCode is the HTTP code returned for type CreateCacheConflict
const CreateCacheConflictCode int = 409

/*CreateCacheConflict The specified resource already exists

swagger:response createCacheConflict
*/
type CreateCacheConflict struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateCacheConflict creates CreateCacheConflict with default headers values
func NewCreateCacheConflict() *CreateCacheConflict {

	return &CreateCacheConflict{}
}

// WithConfigurationVersion adds the configurationVersion to the create cache conflict response
func (o *CreateCacheConflict) WithConfigurationVersion(configurationVersion string) *CreateCacheConflict {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create cache conflict response
func (o *CreateCacheConflict) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create cache conflict response
func (o *CreateCacheConflict) WithPayload(payload *models.Error) *CreateCacheConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create cache conflict response
func (o *CreateCacheConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateCacheConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := o.ConfigurationVersion
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

/*CreateCacheDefault General Error

swagger:response createCacheDefault
*/
type CreateCacheDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateCacheDefault creates CreateCacheDefault with default headers values
func NewCreateCacheDefault(code int) *CreateCacheDefault {
	if code <= 0 {
		code = 500
	}

	return &CreateCacheDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create cache default response
func (o *CreateCacheDefault) WithStatusCode(code int) *CreateCacheDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create cache default response
func (o *CreateCacheDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the create cache default response
func (o *CreateCacheDefault) WithConfigurationVersion(configurationVersion string) *CreateCacheDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create cache default response
func (o *CreateCacheDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create cache default response
func (o *CreateCacheDefault) WithPayload(payload *models.Error) *CreateCacheDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create cache default response
func (o *CreateCacheDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateCacheDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
