// Code generated by go-swagger; DO NOT EDIT.

package member

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
	"github.com/go-openapi/swag"
)

// NewUpdateBenefitTerminateUsingPUTParams creates a new UpdateBenefitTerminateUsingPUTParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateBenefitTerminateUsingPUTParams() *UpdateBenefitTerminateUsingPUTParams {
	return &UpdateBenefitTerminateUsingPUTParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateBenefitTerminateUsingPUTParamsWithTimeout creates a new UpdateBenefitTerminateUsingPUTParams object
// with the ability to set a timeout on a request.
func NewUpdateBenefitTerminateUsingPUTParamsWithTimeout(timeout time.Duration) *UpdateBenefitTerminateUsingPUTParams {
	return &UpdateBenefitTerminateUsingPUTParams{
		timeout: timeout,
	}
}

// NewUpdateBenefitTerminateUsingPUTParamsWithContext creates a new UpdateBenefitTerminateUsingPUTParams object
// with the ability to set a context for a request.
func NewUpdateBenefitTerminateUsingPUTParamsWithContext(ctx context.Context) *UpdateBenefitTerminateUsingPUTParams {
	return &UpdateBenefitTerminateUsingPUTParams{
		Context: ctx,
	}
}

// NewUpdateBenefitTerminateUsingPUTParamsWithHTTPClient creates a new UpdateBenefitTerminateUsingPUTParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateBenefitTerminateUsingPUTParamsWithHTTPClient(client *http.Client) *UpdateBenefitTerminateUsingPUTParams {
	return &UpdateBenefitTerminateUsingPUTParams{
		HTTPClient: client,
	}
}

/*
UpdateBenefitTerminateUsingPUTParams contains all the parameters to send to the API endpoint

	for the update benefit terminate using p u t operation.

	Typically these are written to a http.Request.
*/
type UpdateBenefitTerminateUsingPUTParams struct {

	/* BenefitElectionsToTerminate.

	   benefitElectionsToTerminate
	*/
	BenefitElectionsToTerminate map[string]string

	/* CaseMemberKey.

	   caseMemberKey
	*/
	CaseMemberKey *string

	/* PerformReallocation.

	   performReallocation
	*/
	PerformReallocation *bool

	/* RequestApplication.

	   requestApplication
	*/
	RequestApplication *string

	/* RequestUser.

	   requestUser
	*/
	RequestUser *string

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

// WithDefaults hydrates default values in the update benefit terminate using p u t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateBenefitTerminateUsingPUTParams) WithDefaults() *UpdateBenefitTerminateUsingPUTParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update benefit terminate using p u t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateBenefitTerminateUsingPUTParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update benefit terminate using p u t params
func (o *UpdateBenefitTerminateUsingPUTParams) WithTimeout(timeout time.Duration) *UpdateBenefitTerminateUsingPUTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update benefit terminate using p u t params
func (o *UpdateBenefitTerminateUsingPUTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update benefit terminate using p u t params
func (o *UpdateBenefitTerminateUsingPUTParams) WithContext(ctx context.Context) *UpdateBenefitTerminateUsingPUTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update benefit terminate using p u t params
func (o *UpdateBenefitTerminateUsingPUTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update benefit terminate using p u t params
func (o *UpdateBenefitTerminateUsingPUTParams) WithHTTPClient(client *http.Client) *UpdateBenefitTerminateUsingPUTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update benefit terminate using p u t params
func (o *UpdateBenefitTerminateUsingPUTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBenefitElectionsToTerminate adds the benefitElectionsToTerminate to the update benefit terminate using p u t params
func (o *UpdateBenefitTerminateUsingPUTParams) WithBenefitElectionsToTerminate(benefitElectionsToTerminate map[string]string) *UpdateBenefitTerminateUsingPUTParams {
	o.SetBenefitElectionsToTerminate(benefitElectionsToTerminate)
	return o
}

// SetBenefitElectionsToTerminate adds the benefitElectionsToTerminate to the update benefit terminate using p u t params
func (o *UpdateBenefitTerminateUsingPUTParams) SetBenefitElectionsToTerminate(benefitElectionsToTerminate map[string]string) {
	o.BenefitElectionsToTerminate = benefitElectionsToTerminate
}

// WithCaseMemberKey adds the caseMemberKey to the update benefit terminate using p u t params
func (o *UpdateBenefitTerminateUsingPUTParams) WithCaseMemberKey(caseMemberKey *string) *UpdateBenefitTerminateUsingPUTParams {
	o.SetCaseMemberKey(caseMemberKey)
	return o
}

// SetCaseMemberKey adds the caseMemberKey to the update benefit terminate using p u t params
func (o *UpdateBenefitTerminateUsingPUTParams) SetCaseMemberKey(caseMemberKey *string) {
	o.CaseMemberKey = caseMemberKey
}

// WithPerformReallocation adds the performReallocation to the update benefit terminate using p u t params
func (o *UpdateBenefitTerminateUsingPUTParams) WithPerformReallocation(performReallocation *bool) *UpdateBenefitTerminateUsingPUTParams {
	o.SetPerformReallocation(performReallocation)
	return o
}

// SetPerformReallocation adds the performReallocation to the update benefit terminate using p u t params
func (o *UpdateBenefitTerminateUsingPUTParams) SetPerformReallocation(performReallocation *bool) {
	o.PerformReallocation = performReallocation
}

// WithRequestApplication adds the requestApplication to the update benefit terminate using p u t params
func (o *UpdateBenefitTerminateUsingPUTParams) WithRequestApplication(requestApplication *string) *UpdateBenefitTerminateUsingPUTParams {
	o.SetRequestApplication(requestApplication)
	return o
}

// SetRequestApplication adds the requestApplication to the update benefit terminate using p u t params
func (o *UpdateBenefitTerminateUsingPUTParams) SetRequestApplication(requestApplication *string) {
	o.RequestApplication = requestApplication
}

// WithRequestUser adds the requestUser to the update benefit terminate using p u t params
func (o *UpdateBenefitTerminateUsingPUTParams) WithRequestUser(requestUser *string) *UpdateBenefitTerminateUsingPUTParams {
	o.SetRequestUser(requestUser)
	return o
}

// SetRequestUser adds the requestUser to the update benefit terminate using p u t params
func (o *UpdateBenefitTerminateUsingPUTParams) SetRequestUser(requestUser *string) {
	o.RequestUser = requestUser
}

// WithVarianceFormat adds the varianceFormat to the update benefit terminate using p u t params
func (o *UpdateBenefitTerminateUsingPUTParams) WithVarianceFormat(varianceFormat *string) *UpdateBenefitTerminateUsingPUTParams {
	o.SetVarianceFormat(varianceFormat)
	return o
}

// SetVarianceFormat adds the varianceFormat to the update benefit terminate using p u t params
func (o *UpdateBenefitTerminateUsingPUTParams) SetVarianceFormat(varianceFormat *string) {
	o.VarianceFormat = varianceFormat
}

// WithVarianceLevel adds the varianceLevel to the update benefit terminate using p u t params
func (o *UpdateBenefitTerminateUsingPUTParams) WithVarianceLevel(varianceLevel *string) *UpdateBenefitTerminateUsingPUTParams {
	o.SetVarianceLevel(varianceLevel)
	return o
}

// SetVarianceLevel adds the varianceLevel to the update benefit terminate using p u t params
func (o *UpdateBenefitTerminateUsingPUTParams) SetVarianceLevel(varianceLevel *string) {
	o.VarianceLevel = varianceLevel
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateBenefitTerminateUsingPUTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.BenefitElectionsToTerminate != nil {
		if err := r.SetBodyParam(o.BenefitElectionsToTerminate); err != nil {
			return err
		}
	}

	if o.CaseMemberKey != nil {

		// query param caseMemberKey
		var qrCaseMemberKey string

		if o.CaseMemberKey != nil {
			qrCaseMemberKey = *o.CaseMemberKey
		}
		qCaseMemberKey := qrCaseMemberKey
		if qCaseMemberKey != "" {

			if err := r.SetQueryParam("caseMemberKey", qCaseMemberKey); err != nil {
				return err
			}
		}
	}

	if o.PerformReallocation != nil {

		// query param performReallocation
		var qrPerformReallocation bool

		if o.PerformReallocation != nil {
			qrPerformReallocation = *o.PerformReallocation
		}
		qPerformReallocation := swag.FormatBool(qrPerformReallocation)
		if qPerformReallocation != "" {

			if err := r.SetQueryParam("performReallocation", qPerformReallocation); err != nil {
				return err
			}
		}
	}

	if o.RequestApplication != nil {

		// query param requestApplication
		var qrRequestApplication string

		if o.RequestApplication != nil {
			qrRequestApplication = *o.RequestApplication
		}
		qRequestApplication := qrRequestApplication
		if qRequestApplication != "" {

			if err := r.SetQueryParam("requestApplication", qRequestApplication); err != nil {
				return err
			}
		}
	}

	if o.RequestUser != nil {

		// query param requestUser
		var qrRequestUser string

		if o.RequestUser != nil {
			qrRequestUser = *o.RequestUser
		}
		qRequestUser := qrRequestUser
		if qRequestUser != "" {

			if err := r.SetQueryParam("requestUser", qRequestUser); err != nil {
				return err
			}
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
