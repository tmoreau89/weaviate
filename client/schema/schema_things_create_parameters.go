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

package schema

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/semi-technologies/weaviate/entities/models"
)

// NewSchemaThingsCreateParams creates a new SchemaThingsCreateParams object
// with the default values initialized.
func NewSchemaThingsCreateParams() *SchemaThingsCreateParams {
	var ()
	return &SchemaThingsCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSchemaThingsCreateParamsWithTimeout creates a new SchemaThingsCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSchemaThingsCreateParamsWithTimeout(timeout time.Duration) *SchemaThingsCreateParams {
	var ()
	return &SchemaThingsCreateParams{

		timeout: timeout,
	}
}

// NewSchemaThingsCreateParamsWithContext creates a new SchemaThingsCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewSchemaThingsCreateParamsWithContext(ctx context.Context) *SchemaThingsCreateParams {
	var ()
	return &SchemaThingsCreateParams{

		Context: ctx,
	}
}

// NewSchemaThingsCreateParamsWithHTTPClient creates a new SchemaThingsCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSchemaThingsCreateParamsWithHTTPClient(client *http.Client) *SchemaThingsCreateParams {
	var ()
	return &SchemaThingsCreateParams{
		HTTPClient: client,
	}
}

/*SchemaThingsCreateParams contains all the parameters to send to the API endpoint
for the schema things create operation typically these are written to a http.Request
*/
type SchemaThingsCreateParams struct {

	/*ThingClass*/
	ThingClass *models.Class

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the schema things create params
func (o *SchemaThingsCreateParams) WithTimeout(timeout time.Duration) *SchemaThingsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the schema things create params
func (o *SchemaThingsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the schema things create params
func (o *SchemaThingsCreateParams) WithContext(ctx context.Context) *SchemaThingsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the schema things create params
func (o *SchemaThingsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the schema things create params
func (o *SchemaThingsCreateParams) WithHTTPClient(client *http.Client) *SchemaThingsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the schema things create params
func (o *SchemaThingsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithThingClass adds the thingClass to the schema things create params
func (o *SchemaThingsCreateParams) WithThingClass(thingClass *models.Class) *SchemaThingsCreateParams {
	o.SetThingClass(thingClass)
	return o
}

// SetThingClass adds the thingClass to the schema things create params
func (o *SchemaThingsCreateParams) SetThingClass(thingClass *models.Class) {
	o.ThingClass = thingClass
}

// WriteToRequest writes these params to a swagger request
func (o *SchemaThingsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ThingClass != nil {
		if err := r.SetBodyParam(o.ThingClass); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
