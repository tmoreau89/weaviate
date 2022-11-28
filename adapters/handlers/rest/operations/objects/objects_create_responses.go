//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2022 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package objects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/semi-technologies/weaviate/models"
)

// ObjectsCreateOKCode is the HTTP code returned for type ObjectsCreateOK
const ObjectsCreateOKCode int = 200

/*
ObjectsCreateOK Object created.

swagger:response objectsCreateOK
*/
type ObjectsCreateOK struct {

	/*
	  In: Body
	*/
	Payload *models.Object `json:"body,omitempty"`
}

// NewObjectsCreateOK creates ObjectsCreateOK with default headers values
func NewObjectsCreateOK() *ObjectsCreateOK {

	return &ObjectsCreateOK{}
}

// WithPayload adds the payload to the objects create o k response
func (o *ObjectsCreateOK) WithPayload(payload *models.Object) *ObjectsCreateOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the objects create o k response
func (o *ObjectsCreateOK) SetPayload(payload *models.Object) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ObjectsCreateOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ObjectsCreateUnauthorizedCode is the HTTP code returned for type ObjectsCreateUnauthorized
const ObjectsCreateUnauthorizedCode int = 401

/*
ObjectsCreateUnauthorized Unauthorized or invalid credentials.

swagger:response objectsCreateUnauthorized
*/
type ObjectsCreateUnauthorized struct {
}

// NewObjectsCreateUnauthorized creates ObjectsCreateUnauthorized with default headers values
func NewObjectsCreateUnauthorized() *ObjectsCreateUnauthorized {

	return &ObjectsCreateUnauthorized{}
}

// WriteResponse to the client
func (o *ObjectsCreateUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// ObjectsCreateForbiddenCode is the HTTP code returned for type ObjectsCreateForbidden
const ObjectsCreateForbiddenCode int = 403

/*
ObjectsCreateForbidden Forbidden

swagger:response objectsCreateForbidden
*/
type ObjectsCreateForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewObjectsCreateForbidden creates ObjectsCreateForbidden with default headers values
func NewObjectsCreateForbidden() *ObjectsCreateForbidden {

	return &ObjectsCreateForbidden{}
}

// WithPayload adds the payload to the objects create forbidden response
func (o *ObjectsCreateForbidden) WithPayload(payload *models.ErrorResponse) *ObjectsCreateForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the objects create forbidden response
func (o *ObjectsCreateForbidden) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ObjectsCreateForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ObjectsCreateUnprocessableEntityCode is the HTTP code returned for type ObjectsCreateUnprocessableEntity
const ObjectsCreateUnprocessableEntityCode int = 422

/*
ObjectsCreateUnprocessableEntity Request body is well-formed (i.e., syntactically correct), but semantically erroneous. Are you sure the class is defined in the configuration file?

swagger:response objectsCreateUnprocessableEntity
*/
type ObjectsCreateUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewObjectsCreateUnprocessableEntity creates ObjectsCreateUnprocessableEntity with default headers values
func NewObjectsCreateUnprocessableEntity() *ObjectsCreateUnprocessableEntity {

	return &ObjectsCreateUnprocessableEntity{}
}

// WithPayload adds the payload to the objects create unprocessable entity response
func (o *ObjectsCreateUnprocessableEntity) WithPayload(payload *models.ErrorResponse) *ObjectsCreateUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the objects create unprocessable entity response
func (o *ObjectsCreateUnprocessableEntity) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ObjectsCreateUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ObjectsCreateInternalServerErrorCode is the HTTP code returned for type ObjectsCreateInternalServerError
const ObjectsCreateInternalServerErrorCode int = 500

/*
ObjectsCreateInternalServerError An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.

swagger:response objectsCreateInternalServerError
*/
type ObjectsCreateInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewObjectsCreateInternalServerError creates ObjectsCreateInternalServerError with default headers values
func NewObjectsCreateInternalServerError() *ObjectsCreateInternalServerError {

	return &ObjectsCreateInternalServerError{}
}

// WithPayload adds the payload to the objects create internal server error response
func (o *ObjectsCreateInternalServerError) WithPayload(payload *models.ErrorResponse) *ObjectsCreateInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the objects create internal server error response
func (o *ObjectsCreateInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ObjectsCreateInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
