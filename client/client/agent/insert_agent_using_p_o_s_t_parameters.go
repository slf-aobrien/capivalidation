// Code generated by go-swagger; DO NOT EDIT.

package agent

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

	"client/models"
)

// NewInsertAgentUsingPOSTParams creates a new InsertAgentUsingPOSTParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewInsertAgentUsingPOSTParams() *InsertAgentUsingPOSTParams {
	return &InsertAgentUsingPOSTParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewInsertAgentUsingPOSTParamsWithTimeout creates a new InsertAgentUsingPOSTParams object
// with the ability to set a timeout on a request.
func NewInsertAgentUsingPOSTParamsWithTimeout(timeout time.Duration) *InsertAgentUsingPOSTParams {
	return &InsertAgentUsingPOSTParams{
		timeout: timeout,
	}
}

// NewInsertAgentUsingPOSTParamsWithContext creates a new InsertAgentUsingPOSTParams object
// with the ability to set a context for a request.
func NewInsertAgentUsingPOSTParamsWithContext(ctx context.Context) *InsertAgentUsingPOSTParams {
	return &InsertAgentUsingPOSTParams{
		Context: ctx,
	}
}

// NewInsertAgentUsingPOSTParamsWithHTTPClient creates a new InsertAgentUsingPOSTParams object
// with the ability to set a custom HTTPClient for a request.
func NewInsertAgentUsingPOSTParamsWithHTTPClient(client *http.Client) *InsertAgentUsingPOSTParams {
	return &InsertAgentUsingPOSTParams{
		HTTPClient: client,
	}
}

/*
InsertAgentUsingPOSTParams contains all the parameters to send to the API endpoint

	for the insert agent using p o s t operation.

	Typically these are written to a http.Request.
*/
type InsertAgentUsingPOSTParams struct {

	/* Agent.

	   agent
	*/
	Agent *models.AgentDTOReq

	/* EffectiveDate.

	   effectiveDate
	*/
	EffectiveDate string

	/* RequestApplication.

	   requestApplication
	*/
	RequestApplication string

	/* RequestUser.

	   requestUser
	*/
	RequestUser string

	/* VarianceFormat.

	   varianceFormat
	*/
	VarianceFormat string

	/* VarianceLevel.

	   varianceLevel
	*/
	VarianceLevel string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the insert agent using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InsertAgentUsingPOSTParams) WithDefaults() *InsertAgentUsingPOSTParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the insert agent using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InsertAgentUsingPOSTParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the insert agent using p o s t params
func (o *InsertAgentUsingPOSTParams) WithTimeout(timeout time.Duration) *InsertAgentUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the insert agent using p o s t params
func (o *InsertAgentUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the insert agent using p o s t params
func (o *InsertAgentUsingPOSTParams) WithContext(ctx context.Context) *InsertAgentUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the insert agent using p o s t params
func (o *InsertAgentUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the insert agent using p o s t params
func (o *InsertAgentUsingPOSTParams) WithHTTPClient(client *http.Client) *InsertAgentUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the insert agent using p o s t params
func (o *InsertAgentUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAgent adds the agent to the insert agent using p o s t params
func (o *InsertAgentUsingPOSTParams) WithAgent(agent *models.AgentDTOReq) *InsertAgentUsingPOSTParams {
	o.SetAgent(agent)
	return o
}

// SetAgent adds the agent to the insert agent using p o s t params
func (o *InsertAgentUsingPOSTParams) SetAgent(agent *models.AgentDTOReq) {
	o.Agent = agent
}

// WithEffectiveDate adds the effectiveDate to the insert agent using p o s t params
func (o *InsertAgentUsingPOSTParams) WithEffectiveDate(effectiveDate string) *InsertAgentUsingPOSTParams {
	o.SetEffectiveDate(effectiveDate)
	return o
}

// SetEffectiveDate adds the effectiveDate to the insert agent using p o s t params
func (o *InsertAgentUsingPOSTParams) SetEffectiveDate(effectiveDate string) {
	o.EffectiveDate = effectiveDate
}

// WithRequestApplication adds the requestApplication to the insert agent using p o s t params
func (o *InsertAgentUsingPOSTParams) WithRequestApplication(requestApplication string) *InsertAgentUsingPOSTParams {
	o.SetRequestApplication(requestApplication)
	return o
}

// SetRequestApplication adds the requestApplication to the insert agent using p o s t params
func (o *InsertAgentUsingPOSTParams) SetRequestApplication(requestApplication string) {
	o.RequestApplication = requestApplication
}

// WithRequestUser adds the requestUser to the insert agent using p o s t params
func (o *InsertAgentUsingPOSTParams) WithRequestUser(requestUser string) *InsertAgentUsingPOSTParams {
	o.SetRequestUser(requestUser)
	return o
}

// SetRequestUser adds the requestUser to the insert agent using p o s t params
func (o *InsertAgentUsingPOSTParams) SetRequestUser(requestUser string) {
	o.RequestUser = requestUser
}

// WithVarianceFormat adds the varianceFormat to the insert agent using p o s t params
func (o *InsertAgentUsingPOSTParams) WithVarianceFormat(varianceFormat string) *InsertAgentUsingPOSTParams {
	o.SetVarianceFormat(varianceFormat)
	return o
}

// SetVarianceFormat adds the varianceFormat to the insert agent using p o s t params
func (o *InsertAgentUsingPOSTParams) SetVarianceFormat(varianceFormat string) {
	o.VarianceFormat = varianceFormat
}

// WithVarianceLevel adds the varianceLevel to the insert agent using p o s t params
func (o *InsertAgentUsingPOSTParams) WithVarianceLevel(varianceLevel string) *InsertAgentUsingPOSTParams {
	o.SetVarianceLevel(varianceLevel)
	return o
}

// SetVarianceLevel adds the varianceLevel to the insert agent using p o s t params
func (o *InsertAgentUsingPOSTParams) SetVarianceLevel(varianceLevel string) {
	o.VarianceLevel = varianceLevel
}

// WriteToRequest writes these params to a swagger request
func (o *InsertAgentUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Agent != nil {
		if err := r.SetBodyParam(o.Agent); err != nil {
			return err
		}
	}

	// query param effectiveDate
	qrEffectiveDate := o.EffectiveDate
	qEffectiveDate := qrEffectiveDate
	if qEffectiveDate != "" {

		if err := r.SetQueryParam("effectiveDate", qEffectiveDate); err != nil {
			return err
		}
	}

	// query param requestApplication
	qrRequestApplication := o.RequestApplication
	qRequestApplication := qrRequestApplication
	if qRequestApplication != "" {

		if err := r.SetQueryParam("requestApplication", qRequestApplication); err != nil {
			return err
		}
	}

	// query param requestUser
	qrRequestUser := o.RequestUser
	qRequestUser := qrRequestUser
	if qRequestUser != "" {

		if err := r.SetQueryParam("requestUser", qRequestUser); err != nil {
			return err
		}
	}

	// query param varianceFormat
	qrVarianceFormat := o.VarianceFormat
	qVarianceFormat := qrVarianceFormat
	if qVarianceFormat != "" {

		if err := r.SetQueryParam("varianceFormat", qVarianceFormat); err != nil {
			return err
		}
	}

	// query param varianceLevel
	qrVarianceLevel := o.VarianceLevel
	qVarianceLevel := qrVarianceLevel
	if qVarianceLevel != "" {

		if err := r.SetQueryParam("varianceLevel", qVarianceLevel); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
