// Code generated by go-swagger; DO NOT EDIT.

package system

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetSystemErrorMessageUsingGETParams creates a new GetSystemErrorMessageUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSystemErrorMessageUsingGETParams() *GetSystemErrorMessageUsingGETParams {
	return &GetSystemErrorMessageUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSystemErrorMessageUsingGETParamsWithTimeout creates a new GetSystemErrorMessageUsingGETParams object
// with the ability to set a timeout on a request.
func NewGetSystemErrorMessageUsingGETParamsWithTimeout(timeout time.Duration) *GetSystemErrorMessageUsingGETParams {
	return &GetSystemErrorMessageUsingGETParams{
		timeout: timeout,
	}
}

// NewGetSystemErrorMessageUsingGETParamsWithContext creates a new GetSystemErrorMessageUsingGETParams object
// with the ability to set a context for a request.
func NewGetSystemErrorMessageUsingGETParamsWithContext(ctx context.Context) *GetSystemErrorMessageUsingGETParams {
	return &GetSystemErrorMessageUsingGETParams{
		Context: ctx,
	}
}

// NewGetSystemErrorMessageUsingGETParamsWithHTTPClient creates a new GetSystemErrorMessageUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSystemErrorMessageUsingGETParamsWithHTTPClient(client *http.Client) *GetSystemErrorMessageUsingGETParams {
	return &GetSystemErrorMessageUsingGETParams{
		HTTPClient: client,
	}
}

/*
GetSystemErrorMessageUsingGETParams contains all the parameters to send to the API endpoint

	for the get system error message using g e t operation.

	Typically these are written to a http.Request.
*/
type GetSystemErrorMessageUsingGETParams struct {

	/* ErrorNumber.

	   errorNumber
	*/
	ErrorNumber *string

	/* ErrorPrefix.

	   errorPrefix
	*/
	ErrorPrefix *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get system error message using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSystemErrorMessageUsingGETParams) WithDefaults() *GetSystemErrorMessageUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get system error message using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSystemErrorMessageUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get system error message using g e t params
func (o *GetSystemErrorMessageUsingGETParams) WithTimeout(timeout time.Duration) *GetSystemErrorMessageUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get system error message using g e t params
func (o *GetSystemErrorMessageUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get system error message using g e t params
func (o *GetSystemErrorMessageUsingGETParams) WithContext(ctx context.Context) *GetSystemErrorMessageUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get system error message using g e t params
func (o *GetSystemErrorMessageUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get system error message using g e t params
func (o *GetSystemErrorMessageUsingGETParams) WithHTTPClient(client *http.Client) *GetSystemErrorMessageUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get system error message using g e t params
func (o *GetSystemErrorMessageUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithErrorNumber adds the errorNumber to the get system error message using g e t params
func (o *GetSystemErrorMessageUsingGETParams) WithErrorNumber(errorNumber *string) *GetSystemErrorMessageUsingGETParams {
	o.SetErrorNumber(errorNumber)
	return o
}

// SetErrorNumber adds the errorNumber to the get system error message using g e t params
func (o *GetSystemErrorMessageUsingGETParams) SetErrorNumber(errorNumber *string) {
	o.ErrorNumber = errorNumber
}

// WithErrorPrefix adds the errorPrefix to the get system error message using g e t params
func (o *GetSystemErrorMessageUsingGETParams) WithErrorPrefix(errorPrefix *string) *GetSystemErrorMessageUsingGETParams {
	o.SetErrorPrefix(errorPrefix)
	return o
}

// SetErrorPrefix adds the errorPrefix to the get system error message using g e t params
func (o *GetSystemErrorMessageUsingGETParams) SetErrorPrefix(errorPrefix *string) {
	o.ErrorPrefix = errorPrefix
}

// WriteToRequest writes these params to a swagger request
func (o *GetSystemErrorMessageUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ErrorNumber != nil {

		// query param errorNumber
		var qrErrorNumber string

		if o.ErrorNumber != nil {
			qrErrorNumber = *o.ErrorNumber
		}
		qErrorNumber := qrErrorNumber
		if qErrorNumber != "" {

			if err := r.SetQueryParam("errorNumber", qErrorNumber); err != nil {
				return err
			}
		}
	}

	if o.ErrorPrefix != nil {

		// query param errorPrefix
		var qrErrorPrefix string

		if o.ErrorPrefix != nil {
			qrErrorPrefix = *o.ErrorPrefix
		}
		qErrorPrefix := qrErrorPrefix
		if qErrorPrefix != "" {

			if err := r.SetQueryParam("errorPrefix", qErrorPrefix); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
