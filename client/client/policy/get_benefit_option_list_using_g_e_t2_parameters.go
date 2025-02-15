// Code generated by go-swagger; DO NOT EDIT.

package policy

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

// NewGetBenefitOptionListUsingGET2Params creates a new GetBenefitOptionListUsingGET2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetBenefitOptionListUsingGET2Params() *GetBenefitOptionListUsingGET2Params {
	return &GetBenefitOptionListUsingGET2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetBenefitOptionListUsingGET2ParamsWithTimeout creates a new GetBenefitOptionListUsingGET2Params object
// with the ability to set a timeout on a request.
func NewGetBenefitOptionListUsingGET2ParamsWithTimeout(timeout time.Duration) *GetBenefitOptionListUsingGET2Params {
	return &GetBenefitOptionListUsingGET2Params{
		timeout: timeout,
	}
}

// NewGetBenefitOptionListUsingGET2ParamsWithContext creates a new GetBenefitOptionListUsingGET2Params object
// with the ability to set a context for a request.
func NewGetBenefitOptionListUsingGET2ParamsWithContext(ctx context.Context) *GetBenefitOptionListUsingGET2Params {
	return &GetBenefitOptionListUsingGET2Params{
		Context: ctx,
	}
}

// NewGetBenefitOptionListUsingGET2ParamsWithHTTPClient creates a new GetBenefitOptionListUsingGET2Params object
// with the ability to set a custom HTTPClient for a request.
func NewGetBenefitOptionListUsingGET2ParamsWithHTTPClient(client *http.Client) *GetBenefitOptionListUsingGET2Params {
	return &GetBenefitOptionListUsingGET2Params{
		HTTPClient: client,
	}
}

/*
GetBenefitOptionListUsingGET2Params contains all the parameters to send to the API endpoint

	for the get benefit option list using g e t 2 operation.

	Typically these are written to a http.Request.
*/
type GetBenefitOptionListUsingGET2Params struct {

	/* EffectiveDate.

	   effectiveDate
	*/
	EffectiveDate *string

	/* PolicyNumber.

	   policyNumber
	*/
	PolicyNumber string

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
	VarianceFormat *string

	/* VarianceLevel.

	   varianceLevel
	*/
	VarianceLevel *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get benefit option list using g e t 2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBenefitOptionListUsingGET2Params) WithDefaults() *GetBenefitOptionListUsingGET2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get benefit option list using g e t 2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBenefitOptionListUsingGET2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get benefit option list using g e t 2 params
func (o *GetBenefitOptionListUsingGET2Params) WithTimeout(timeout time.Duration) *GetBenefitOptionListUsingGET2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get benefit option list using g e t 2 params
func (o *GetBenefitOptionListUsingGET2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get benefit option list using g e t 2 params
func (o *GetBenefitOptionListUsingGET2Params) WithContext(ctx context.Context) *GetBenefitOptionListUsingGET2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get benefit option list using g e t 2 params
func (o *GetBenefitOptionListUsingGET2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get benefit option list using g e t 2 params
func (o *GetBenefitOptionListUsingGET2Params) WithHTTPClient(client *http.Client) *GetBenefitOptionListUsingGET2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get benefit option list using g e t 2 params
func (o *GetBenefitOptionListUsingGET2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEffectiveDate adds the effectiveDate to the get benefit option list using g e t 2 params
func (o *GetBenefitOptionListUsingGET2Params) WithEffectiveDate(effectiveDate *string) *GetBenefitOptionListUsingGET2Params {
	o.SetEffectiveDate(effectiveDate)
	return o
}

// SetEffectiveDate adds the effectiveDate to the get benefit option list using g e t 2 params
func (o *GetBenefitOptionListUsingGET2Params) SetEffectiveDate(effectiveDate *string) {
	o.EffectiveDate = effectiveDate
}

// WithPolicyNumber adds the policyNumber to the get benefit option list using g e t 2 params
func (o *GetBenefitOptionListUsingGET2Params) WithPolicyNumber(policyNumber string) *GetBenefitOptionListUsingGET2Params {
	o.SetPolicyNumber(policyNumber)
	return o
}

// SetPolicyNumber adds the policyNumber to the get benefit option list using g e t 2 params
func (o *GetBenefitOptionListUsingGET2Params) SetPolicyNumber(policyNumber string) {
	o.PolicyNumber = policyNumber
}

// WithRequestApplication adds the requestApplication to the get benefit option list using g e t 2 params
func (o *GetBenefitOptionListUsingGET2Params) WithRequestApplication(requestApplication string) *GetBenefitOptionListUsingGET2Params {
	o.SetRequestApplication(requestApplication)
	return o
}

// SetRequestApplication adds the requestApplication to the get benefit option list using g e t 2 params
func (o *GetBenefitOptionListUsingGET2Params) SetRequestApplication(requestApplication string) {
	o.RequestApplication = requestApplication
}

// WithRequestUser adds the requestUser to the get benefit option list using g e t 2 params
func (o *GetBenefitOptionListUsingGET2Params) WithRequestUser(requestUser string) *GetBenefitOptionListUsingGET2Params {
	o.SetRequestUser(requestUser)
	return o
}

// SetRequestUser adds the requestUser to the get benefit option list using g e t 2 params
func (o *GetBenefitOptionListUsingGET2Params) SetRequestUser(requestUser string) {
	o.RequestUser = requestUser
}

// WithVarianceFormat adds the varianceFormat to the get benefit option list using g e t 2 params
func (o *GetBenefitOptionListUsingGET2Params) WithVarianceFormat(varianceFormat *string) *GetBenefitOptionListUsingGET2Params {
	o.SetVarianceFormat(varianceFormat)
	return o
}

// SetVarianceFormat adds the varianceFormat to the get benefit option list using g e t 2 params
func (o *GetBenefitOptionListUsingGET2Params) SetVarianceFormat(varianceFormat *string) {
	o.VarianceFormat = varianceFormat
}

// WithVarianceLevel adds the varianceLevel to the get benefit option list using g e t 2 params
func (o *GetBenefitOptionListUsingGET2Params) WithVarianceLevel(varianceLevel *string) *GetBenefitOptionListUsingGET2Params {
	o.SetVarianceLevel(varianceLevel)
	return o
}

// SetVarianceLevel adds the varianceLevel to the get benefit option list using g e t 2 params
func (o *GetBenefitOptionListUsingGET2Params) SetVarianceLevel(varianceLevel *string) {
	o.VarianceLevel = varianceLevel
}

// WriteToRequest writes these params to a swagger request
func (o *GetBenefitOptionListUsingGET2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.EffectiveDate != nil {

		// query param effectiveDate
		var qrEffectiveDate string

		if o.EffectiveDate != nil {
			qrEffectiveDate = *o.EffectiveDate
		}
		qEffectiveDate := qrEffectiveDate
		if qEffectiveDate != "" {

			if err := r.SetQueryParam("effectiveDate", qEffectiveDate); err != nil {
				return err
			}
		}
	}

	// path param policyNumber
	if err := r.SetPathParam("policyNumber", o.PolicyNumber); err != nil {
		return err
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

	if o.VarianceFormat != nil {

		// query param varianceFormat
		var qrVarianceFormat string

		if o.VarianceFormat != nil {
			qrVarianceFormat = *o.VarianceFormat
		}
		qVarianceFormat := qrVarianceFormat
		if qVarianceFormat != "" {

			if err := r.SetQueryParam("varianceFormat", qVarianceFormat); err != nil {
				return err
			}
		}
	}

	if o.VarianceLevel != nil {

		// query param varianceLevel
		var qrVarianceLevel string

		if o.VarianceLevel != nil {
			qrVarianceLevel = *o.VarianceLevel
		}
		qVarianceLevel := qrVarianceLevel
		if qVarianceLevel != "" {

			if err := r.SetQueryParam("varianceLevel", qVarianceLevel); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
