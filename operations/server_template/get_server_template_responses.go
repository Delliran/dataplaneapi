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

package server_template

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	"github.com/haproxytech/client-native/v3/models"
)

// GetServerTemplateOKCode is the HTTP code returned for type GetServerTemplateOK
const GetServerTemplateOKCode int = 200

/*GetServerTemplateOK Successful operation

swagger:response getServerTemplateOK
*/
type GetServerTemplateOK struct {
	/*Configuration file version

	 */
	ConfigurationVersion int64 `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *GetServerTemplateOKBody `json:"body,omitempty"`
}

// NewGetServerTemplateOK creates GetServerTemplateOK with default headers values
func NewGetServerTemplateOK() *GetServerTemplateOK {

	return &GetServerTemplateOK{}
}

// WithConfigurationVersion adds the configurationVersion to the get server template o k response
func (o *GetServerTemplateOK) WithConfigurationVersion(configurationVersion int64) *GetServerTemplateOK {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get server template o k response
func (o *GetServerTemplateOK) SetConfigurationVersion(configurationVersion int64) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get server template o k response
func (o *GetServerTemplateOK) WithPayload(payload *GetServerTemplateOKBody) *GetServerTemplateOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get server template o k response
func (o *GetServerTemplateOK) SetPayload(payload *GetServerTemplateOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetServerTemplateOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := swag.FormatInt64(o.ConfigurationVersion)
	if configurationVersion != "" {
		rw.Header().Set("Configuration-Version", configurationVersion)
	}

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetServerTemplateNotFoundCode is the HTTP code returned for type GetServerTemplateNotFound
const GetServerTemplateNotFoundCode int = 404

/*GetServerTemplateNotFound The specified resource was not found

swagger:response getServerTemplateNotFound
*/
type GetServerTemplateNotFound struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetServerTemplateNotFound creates GetServerTemplateNotFound with default headers values
func NewGetServerTemplateNotFound() *GetServerTemplateNotFound {

	return &GetServerTemplateNotFound{}
}

// WithConfigurationVersion adds the configurationVersion to the get server template not found response
func (o *GetServerTemplateNotFound) WithConfigurationVersion(configurationVersion string) *GetServerTemplateNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get server template not found response
func (o *GetServerTemplateNotFound) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get server template not found response
func (o *GetServerTemplateNotFound) WithPayload(payload *models.Error) *GetServerTemplateNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get server template not found response
func (o *GetServerTemplateNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetServerTemplateNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

/*GetServerTemplateDefault General Error

swagger:response getServerTemplateDefault
*/
type GetServerTemplateDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetServerTemplateDefault creates GetServerTemplateDefault with default headers values
func NewGetServerTemplateDefault(code int) *GetServerTemplateDefault {
	if code <= 0 {
		code = 500
	}

	return &GetServerTemplateDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get server template default response
func (o *GetServerTemplateDefault) WithStatusCode(code int) *GetServerTemplateDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get server template default response
func (o *GetServerTemplateDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the get server template default response
func (o *GetServerTemplateDefault) WithConfigurationVersion(configurationVersion string) *GetServerTemplateDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get server template default response
func (o *GetServerTemplateDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get server template default response
func (o *GetServerTemplateDefault) WithPayload(payload *models.Error) *GetServerTemplateDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get server template default response
func (o *GetServerTemplateDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetServerTemplateDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
