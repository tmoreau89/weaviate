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

// ObjectsClassPatchNoContentCode is the HTTP code returned for type ObjectsClassPatchNoContent
const ObjectsClassPatchNoContentCode int = 204

/*
ObjectsClassPatchNoContent Successfully applied. No content provided.

swagger:response objectsClassPatchNoContent
*/
type ObjectsClassPatchNoContent struct {
}

// NewObjectsClassPatchNoContent creates ObjectsClassPatchNoContent with default headers values
func NewObjectsClassPatchNoContent() *ObjectsClassPatchNoContent {

	return &ObjectsClassPatchNoContent{}
}

// WriteResponse to the client
func (o *ObjectsClassPatchNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// ObjectsClassPatchBadRequestCode is the HTTP code returned for type ObjectsClassPatchBadRequest
const ObjectsClassPatchBadRequestCode int = 400

/*
ObjectsClassPatchBadRequest The patch-JSON is malformed.

swagger:response objectsClassPatchBadRequest
*/
type ObjectsClassPatchBadRequest struct {
}

// NewObjectsClassPatchBadRequest creates ObjectsClassPatchBadRequest with default headers values
func NewObjectsClassPatchBadRequest() *ObjectsClassPatchBadRequest {

	return &ObjectsClassPatchBadRequest{}
}

// WriteResponse to the client
func (o *ObjectsClassPatchBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// ObjectsClassPatchUnauthorizedCode is the HTTP code returned for type ObjectsClassPatchUnauthorized
const ObjectsClassPatchUnauthorizedCode int = 401

/*
ObjectsClassPatchUnauthorized Unauthorized or invalid credentials.

swagger:response objectsClassPatchUnauthorized
*/
type ObjectsClassPatchUnauthorized struct {
}

// NewObjectsClassPatchUnauthorized creates ObjectsClassPatchUnauthorized with default headers values
func NewObjectsClassPatchUnauthorized() *ObjectsClassPatchUnauthorized {

	return &ObjectsClassPatchUnauthorized{}
}

// WriteResponse to the client
func (o *ObjectsClassPatchUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// ObjectsClassPatchForbiddenCode is the HTTP code returned for type ObjectsClassPatchForbidden
const ObjectsClassPatchForbiddenCode int = 403

/*
ObjectsClassPatchForbidden Forbidden

swagger:response objectsClassPatchForbidden
*/
type ObjectsClassPatchForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewObjectsClassPatchForbidden creates ObjectsClassPatchForbidden with default headers values
func NewObjectsClassPatchForbidden() *ObjectsClassPatchForbidden {

	return &ObjectsClassPatchForbidden{}
}

// WithPayload adds the payload to the objects class patch forbidden response
func (o *ObjectsClassPatchForbidden) WithPayload(payload *models.ErrorResponse) *ObjectsClassPatchForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the objects class patch forbidden response
func (o *ObjectsClassPatchForbidden) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ObjectsClassPatchForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ObjectsClassPatchNotFoundCode is the HTTP code returned for type ObjectsClassPatchNotFound
const ObjectsClassPatchNotFoundCode int = 404

/*
ObjectsClassPatchNotFound Successful query result but no resource was found.

swagger:response objectsClassPatchNotFound
*/
type ObjectsClassPatchNotFound struct {
}

// NewObjectsClassPatchNotFound creates ObjectsClassPatchNotFound with default headers values
func NewObjectsClassPatchNotFound() *ObjectsClassPatchNotFound {

	return &ObjectsClassPatchNotFound{}
}

// WriteResponse to the client
func (o *ObjectsClassPatchNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// ObjectsClassPatchUnprocessableEntityCode is the HTTP code returned for type ObjectsClassPatchUnprocessableEntity
const ObjectsClassPatchUnprocessableEntityCode int = 422

/*
ObjectsClassPatchUnprocessableEntity The patch-JSON is valid but unprocessable.

swagger:response objectsClassPatchUnprocessableEntity
*/
type ObjectsClassPatchUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewObjectsClassPatchUnprocessableEntity creates ObjectsClassPatchUnprocessableEntity with default headers values
func NewObjectsClassPatchUnprocessableEntity() *ObjectsClassPatchUnprocessableEntity {

	return &ObjectsClassPatchUnprocessableEntity{}
}

// WithPayload adds the payload to the objects class patch unprocessable entity response
func (o *ObjectsClassPatchUnprocessableEntity) WithPayload(payload *models.ErrorResponse) *ObjectsClassPatchUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the objects class patch unprocessable entity response
func (o *ObjectsClassPatchUnprocessableEntity) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ObjectsClassPatchUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ObjectsClassPatchInternalServerErrorCode is the HTTP code returned for type ObjectsClassPatchInternalServerError
const ObjectsClassPatchInternalServerErrorCode int = 500

/*
ObjectsClassPatchInternalServerError An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.

swagger:response objectsClassPatchInternalServerError
*/
type ObjectsClassPatchInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewObjectsClassPatchInternalServerError creates ObjectsClassPatchInternalServerError with default headers values
func NewObjectsClassPatchInternalServerError() *ObjectsClassPatchInternalServerError {

	return &ObjectsClassPatchInternalServerError{}
}

// WithPayload adds the payload to the objects class patch internal server error response
func (o *ObjectsClassPatchInternalServerError) WithPayload(payload *models.ErrorResponse) *ObjectsClassPatchInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the objects class patch internal server error response
func (o *ObjectsClassPatchInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ObjectsClassPatchInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
