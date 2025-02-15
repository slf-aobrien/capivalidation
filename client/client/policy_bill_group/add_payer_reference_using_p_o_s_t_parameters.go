// Code generated by go-swagger; DO NOT EDIT.

package policy_bill_group

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

// NewAddPayerReferenceUsingPOSTParams creates a new AddPayerReferenceUsingPOSTParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddPayerReferenceUsingPOSTParams() *AddPayerReferenceUsingPOSTParams {
	return &AddPayerReferenceUsingPOSTParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddPayerReferenceUsingPOSTParamsWithTimeout creates a new AddPayerReferenceUsingPOSTParams object
// with the ability to set a timeout on a request.
func NewAddPayerReferenceUsingPOSTParamsWithTimeout(timeout time.Duration) *AddPayerReferenceUsingPOSTParams {
	return &AddPayerReferenceUsingPOSTParams{
		timeout: timeout,
	}
}

// NewAddPayerReferenceUsingPOSTParamsWithContext creates a new AddPayerReferenceUsingPOSTParams object
// with the ability to set a context for a request.
func NewAddPayerReferenceUsingPOSTParamsWithContext(ctx context.Context) *AddPayerReferenceUsingPOSTParams {
	return &AddPayerReferenceUsingPOSTParams{
		Context: ctx,
	}
}

// NewAddPayerReferenceUsingPOSTParamsWithHTTPClient creates a new AddPayerReferenceUsingPOSTParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddPayerReferenceUsingPOSTParamsWithHTTPClient(client *http.Client) *AddPayerReferenceUsingPOSTParams {
	return &AddPayerReferenceUsingPOSTParams{
		HTTPClient: client,
	}
}

/*
AddPayerReferenceUsingPOSTParams contains all the parameters to send to the API endpoint

	for the add payer reference using p o s t operation.

	Typically these are written to a http.Request.
*/
type AddPayerReferenceUsingPOSTParams struct {

	/* BillGroupDescription.

	   Existing Bill Group Description for which payer needs to be added.
	*/
	BillGroupDescription string

	/* PolicyEffectiveDate.

	   Date should be in MM/DD/YYYY format.
	*/
	PolicyEffectiveDate string

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

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add payer reference using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddPayerReferenceUsingPOSTParams) WithDefaults() *AddPayerReferenceUsingPOSTParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add payer reference using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddPayerReferenceUsingPOSTParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add payer reference using p o s t params
func (o *AddPayerReferenceUsingPOSTParams) WithTimeout(timeout time.Duration) *AddPayerReferenceUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add payer reference using p o s t params
func (o *AddPayerReferenceUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add payer reference using p o s t params
func (o *AddPayerReferenceUsingPOSTParams) WithContext(ctx context.Context) *AddPayerReferenceUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add payer reference using p o s t params
func (o *AddPayerReferenceUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add payer reference using p o s t params
func (o *AddPayerReferenceUsingPOSTParams) WithHTTPClient(client *http.Client) *AddPayerReferenceUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add payer reference using p o s t params
func (o *AddPayerReferenceUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBillGroupDescription adds the billGroupDescription to the add payer reference using p o s t params
func (o *AddPayerReferenceUsingPOSTParams) WithBillGroupDescription(billGroupDescription string) *AddPayerReferenceUsingPOSTParams {
	o.SetBillGroupDescription(billGroupDescription)
	return o
}

// SetBillGroupDescription adds the billGroupDescription to the add payer reference using p o s t params
func (o *AddPayerReferenceUsingPOSTParams) SetBillGroupDescription(billGroupDescription string) {
	o.BillGroupDescription = billGroupDescription
}

// WithPolicyEffectiveDate adds the policyEffectiveDate to the add payer reference using p o s t params
func (o *AddPayerReferenceUsingPOSTParams) WithPolicyEffectiveDate(policyEffectiveDate string) *AddPayerReferenceUsingPOSTParams {
	o.SetPolicyEffectiveDate(policyEffectiveDate)
	return o
}

// SetPolicyEffectiveDate adds the policyEffectiveDate to the add payer reference using p o s t params
func (o *AddPayerReferenceUsingPOSTParams) SetPolicyEffectiveDate(policyEffectiveDate string) {
	o.PolicyEffectiveDate = policyEffectiveDate
}

// WithPolicyNumber adds the policyNumber to the add payer reference using p o s t params
func (o *AddPayerReferenceUsingPOSTParams) WithPolicyNumber(policyNumber string) *AddPayerReferenceUsingPOSTParams {
	o.SetPolicyNumber(policyNumber)
	return o
}

// SetPolicyNumber adds the policyNumber to the add payer reference using p o s t params
func (o *AddPayerReferenceUsingPOSTParams) SetPolicyNumber(policyNumber string) {
	o.PolicyNumber = policyNumber
}

// WithRequestApplication adds the requestApplication to the add payer reference using p o s t params
func (o *AddPayerReferenceUsingPOSTParams) WithRequestApplication(requestApplication string) *AddPayerReferenceUsingPOSTParams {
	o.SetRequestApplication(requestApplication)
	return o
}

// SetRequestApplication adds the requestApplication to the add payer reference using p o s t params
func (o *AddPayerReferenceUsingPOSTParams) SetRequestApplication(requestApplication string) {
	o.RequestApplication = requestApplication
}

// WithRequestUser adds the requestUser to the add payer reference using p o s t params
func (o *AddPayerReferenceUsingPOSTParams) WithRequestUser(requestUser string) *AddPayerReferenceUsingPOSTParams {
	o.SetRequestUser(requestUser)
	return o
}

// SetRequestUser adds the requestUser to the add payer reference using p o s t params
func (o *AddPayerReferenceUsingPOSTParams) SetRequestUser(requestUser string) {
	o.RequestUser = requestUser
}

// WriteToRequest writes these params to a swagger request
func (o *AddPayerReferenceUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param billGroupDescription
	qrBillGroupDescription := o.BillGroupDescription
	qBillGroupDescription := qrBillGroupDescription
	if qBillGroupDescription != "" {

		if err := r.SetQueryParam("billGroupDescription", qBillGroupDescription); err != nil {
			return err
		}
	}

	// query param policyEffectiveDate
	qrPolicyEffectiveDate := o.PolicyEffectiveDate
	qPolicyEffectiveDate := qrPolicyEffectiveDate
	if qPolicyEffectiveDate != "" {

		if err := r.SetQueryParam("policyEffectiveDate", qPolicyEffectiveDate); err != nil {
			return err
		}
	}

	// query param policyNumber
	qrPolicyNumber := o.PolicyNumber
	qPolicyNumber := qrPolicyNumber
	if qPolicyNumber != "" {

		if err := r.SetQueryParam("policyNumber", qPolicyNumber); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
