//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2019 Weaviate. All rights reserved.
//  LICENSE WEAVIATE OPEN SOURCE: https://www.semi.technology/playbook/playbook/contract-weaviate-OSS.html
//  LICENSE WEAVIATE ENTERPRISE: https://www.semi.technology/playbook/contract-weaviate-enterprise.html
//  CONCEPT: Bob van Luijt (@bobvanluijt)
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/semi-technologies/weaviate/entities/models"
)

// ActionsUpdateReader is a Reader for the ActionsUpdate structure.
type ActionsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ActionsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewActionsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewActionsUpdateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewActionsUpdateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewActionsUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewActionsUpdateUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewActionsUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewActionsUpdateOK creates a ActionsUpdateOK with default headers values
func NewActionsUpdateOK() *ActionsUpdateOK {
	return &ActionsUpdateOK{}
}

/*ActionsUpdateOK handles this case with default header values.

Successfully received.
*/
type ActionsUpdateOK struct {
	Payload *models.Action
}

func (o *ActionsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /actions/{id}][%d] actionsUpdateOK  %+v", 200, o.Payload)
}

func (o *ActionsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Action)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewActionsUpdateUnauthorized creates a ActionsUpdateUnauthorized with default headers values
func NewActionsUpdateUnauthorized() *ActionsUpdateUnauthorized {
	return &ActionsUpdateUnauthorized{}
}

/*ActionsUpdateUnauthorized handles this case with default header values.

Unauthorized or invalid credentials.
*/
type ActionsUpdateUnauthorized struct {
}

func (o *ActionsUpdateUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /actions/{id}][%d] actionsUpdateUnauthorized ", 401)
}

func (o *ActionsUpdateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewActionsUpdateForbidden creates a ActionsUpdateForbidden with default headers values
func NewActionsUpdateForbidden() *ActionsUpdateForbidden {
	return &ActionsUpdateForbidden{}
}

/*ActionsUpdateForbidden handles this case with default header values.

Forbidden
*/
type ActionsUpdateForbidden struct {
	Payload *models.ErrorResponse
}

func (o *ActionsUpdateForbidden) Error() string {
	return fmt.Sprintf("[PUT /actions/{id}][%d] actionsUpdateForbidden  %+v", 403, o.Payload)
}

func (o *ActionsUpdateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewActionsUpdateNotFound creates a ActionsUpdateNotFound with default headers values
func NewActionsUpdateNotFound() *ActionsUpdateNotFound {
	return &ActionsUpdateNotFound{}
}

/*ActionsUpdateNotFound handles this case with default header values.

Successful query result but no resource was found.
*/
type ActionsUpdateNotFound struct {
}

func (o *ActionsUpdateNotFound) Error() string {
	return fmt.Sprintf("[PUT /actions/{id}][%d] actionsUpdateNotFound ", 404)
}

func (o *ActionsUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewActionsUpdateUnprocessableEntity creates a ActionsUpdateUnprocessableEntity with default headers values
func NewActionsUpdateUnprocessableEntity() *ActionsUpdateUnprocessableEntity {
	return &ActionsUpdateUnprocessableEntity{}
}

/*ActionsUpdateUnprocessableEntity handles this case with default header values.

Request body is well-formed (i.e., syntactically correct), but semantically erroneous. Are you sure the class is defined in the configuration file?
*/
type ActionsUpdateUnprocessableEntity struct {
	Payload *models.ErrorResponse
}

func (o *ActionsUpdateUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /actions/{id}][%d] actionsUpdateUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *ActionsUpdateUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewActionsUpdateInternalServerError creates a ActionsUpdateInternalServerError with default headers values
func NewActionsUpdateInternalServerError() *ActionsUpdateInternalServerError {
	return &ActionsUpdateInternalServerError{}
}

/*ActionsUpdateInternalServerError handles this case with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type ActionsUpdateInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *ActionsUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /actions/{id}][%d] actionsUpdateInternalServerError  %+v", 500, o.Payload)
}

func (o *ActionsUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
