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

// NewPolicyLevelOptionsUsingGETParams creates a new PolicyLevelOptionsUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPolicyLevelOptionsUsingGETParams() *PolicyLevelOptionsUsingGETParams {
	return &PolicyLevelOptionsUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPolicyLevelOptionsUsingGETParamsWithTimeout creates a new PolicyLevelOptionsUsingGETParams object
// with the ability to set a timeout on a request.
func NewPolicyLevelOptionsUsingGETParamsWithTimeout(timeout time.Duration) *PolicyLevelOptionsUsingGETParams {
	return &PolicyLevelOptionsUsingGETParams{
		timeout: timeout,
	}
}

// NewPolicyLevelOptionsUsingGETParamsWithContext creates a new PolicyLevelOptionsUsingGETParams object
// with the ability to set a context for a request.
func NewPolicyLevelOptionsUsingGETParamsWithContext(ctx context.Context) *PolicyLevelOptionsUsingGETParams {
	return &PolicyLevelOptionsUsingGETParams{
		Context: ctx,
	}
}

// NewPolicyLevelOptionsUsingGETParamsWithHTTPClient creates a new PolicyLevelOptionsUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewPolicyLevelOptionsUsingGETParamsWithHTTPClient(client *http.Client) *PolicyLevelOptionsUsingGETParams {
	return &PolicyLevelOptionsUsingGETParams{
		HTTPClient: client,
	}
}

/*
PolicyLevelOptionsUsingGETParams contains all the parameters to send to the API endpoint

	for the policy level options using g e t operation.

	Typically these are written to a http.Request.
*/
type PolicyLevelOptionsUsingGETParams struct {

	/* EffectiveDate.

	   effectiveDate
	*/
	EffectiveDate *string

	/* LabelAliasNameSet.

	   labelAliasNameSet
	*/
	LabelAliasNameSet *string

	/* PolicyNumber.

	   policyNumber
	*/
	PolicyNumber *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the policy level options using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PolicyLevelOptionsUsingGETParams) WithDefaults() *PolicyLevelOptionsUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the policy level options using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PolicyLevelOptionsUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the policy level options using g e t params
func (o *PolicyLevelOptionsUsingGETParams) WithTimeout(timeout time.Duration) *PolicyLevelOptionsUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the policy level options using g e t params
func (o *PolicyLevelOptionsUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the policy level options using g e t params
func (o *PolicyLevelOptionsUsingGETParams) WithContext(ctx context.Context) *PolicyLevelOptionsUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the policy level options using g e t params
func (o *PolicyLevelOptionsUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the policy level options using g e t params
func (o *PolicyLevelOptionsUsingGETParams) WithHTTPClient(client *http.Client) *PolicyLevelOptionsUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the policy level options using g e t params
func (o *PolicyLevelOptionsUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEffectiveDate adds the effectiveDate to the policy level options using g e t params
func (o *PolicyLevelOptionsUsingGETParams) WithEffectiveDate(effectiveDate *string) *PolicyLevelOptionsUsingGETParams {
	o.SetEffectiveDate(effectiveDate)
	return o
}

// SetEffectiveDate adds the effectiveDate to the policy level options using g e t params
func (o *PolicyLevelOptionsUsingGETParams) SetEffectiveDate(effectiveDate *string) {
	o.EffectiveDate = effectiveDate
}

// WithLabelAliasNameSet adds the labelAliasNameSet to the policy level options using g e t params
func (o *PolicyLevelOptionsUsingGETParams) WithLabelAliasNameSet(labelAliasNameSet *string) *PolicyLevelOptionsUsingGETParams {
	o.SetLabelAliasNameSet(labelAliasNameSet)
	return o
}

// SetLabelAliasNameSet adds the labelAliasNameSet to the policy level options using g e t params
func (o *PolicyLevelOptionsUsingGETParams) SetLabelAliasNameSet(labelAliasNameSet *string) {
	o.LabelAliasNameSet = labelAliasNameSet
}

// WithPolicyNumber adds the policyNumber to the policy level options using g e t params
func (o *PolicyLevelOptionsUsingGETParams) WithPolicyNumber(policyNumber *string) *PolicyLevelOptionsUsingGETParams {
	o.SetPolicyNumber(policyNumber)
	return o
}

// SetPolicyNumber adds the policyNumber to the policy level options using g e t params
func (o *PolicyLevelOptionsUsingGETParams) SetPolicyNumber(policyNumber *string) {
	o.PolicyNumber = policyNumber
}

// WriteToRequest writes these params to a swagger request
func (o *PolicyLevelOptionsUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.LabelAliasNameSet != nil {

		// query param labelAliasNameSet
		var qrLabelAliasNameSet string

		if o.LabelAliasNameSet != nil {
			qrLabelAliasNameSet = *o.LabelAliasNameSet
		}
		qLabelAliasNameSet := qrLabelAliasNameSet
		if qLabelAliasNameSet != "" {

			if err := r.SetQueryParam("labelAliasNameSet", qLabelAliasNameSet); err != nil {
				return err
			}
		}
	}

	if o.PolicyNumber != nil {

		// query param policyNumber
		var qrPolicyNumber string

		if o.PolicyNumber != nil {
			qrPolicyNumber = *o.PolicyNumber
		}
		qPolicyNumber := qrPolicyNumber
		if qPolicyNumber != "" {

			if err := r.SetQueryParam("policyNumber", qPolicyNumber); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
