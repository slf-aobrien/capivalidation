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

	"client/models"
)

// NewInsertAssociatedCompensationEntitiesUsingPOSTParams creates a new InsertAssociatedCompensationEntitiesUsingPOSTParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewInsertAssociatedCompensationEntitiesUsingPOSTParams() *InsertAssociatedCompensationEntitiesUsingPOSTParams {
	return &InsertAssociatedCompensationEntitiesUsingPOSTParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewInsertAssociatedCompensationEntitiesUsingPOSTParamsWithTimeout creates a new InsertAssociatedCompensationEntitiesUsingPOSTParams object
// with the ability to set a timeout on a request.
func NewInsertAssociatedCompensationEntitiesUsingPOSTParamsWithTimeout(timeout time.Duration) *InsertAssociatedCompensationEntitiesUsingPOSTParams {
	return &InsertAssociatedCompensationEntitiesUsingPOSTParams{
		timeout: timeout,
	}
}

// NewInsertAssociatedCompensationEntitiesUsingPOSTParamsWithContext creates a new InsertAssociatedCompensationEntitiesUsingPOSTParams object
// with the ability to set a context for a request.
func NewInsertAssociatedCompensationEntitiesUsingPOSTParamsWithContext(ctx context.Context) *InsertAssociatedCompensationEntitiesUsingPOSTParams {
	return &InsertAssociatedCompensationEntitiesUsingPOSTParams{
		Context: ctx,
	}
}

// NewInsertAssociatedCompensationEntitiesUsingPOSTParamsWithHTTPClient creates a new InsertAssociatedCompensationEntitiesUsingPOSTParams object
// with the ability to set a custom HTTPClient for a request.
func NewInsertAssociatedCompensationEntitiesUsingPOSTParamsWithHTTPClient(client *http.Client) *InsertAssociatedCompensationEntitiesUsingPOSTParams {
	return &InsertAssociatedCompensationEntitiesUsingPOSTParams{
		HTTPClient: client,
	}
}

/*
InsertAssociatedCompensationEntitiesUsingPOSTParams contains all the parameters to send to the API endpoint

	for the insert associated compensation entities using p o s t operation.

	Typically these are written to a http.Request.
*/
type InsertAssociatedCompensationEntitiesUsingPOSTParams struct {

	/* CompensationEntityList.

	   compensationEntityList
	*/
	CompensationEntityList []*models.CompensationEntityDTO

	/* EffectiveDate.

	   effectiveDate
	*/
	EffectiveDate *string

	/* PolicyNumber.

	   policyNumber
	*/
	PolicyNumber *string

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

// WithDefaults hydrates default values in the insert associated compensation entities using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InsertAssociatedCompensationEntitiesUsingPOSTParams) WithDefaults() *InsertAssociatedCompensationEntitiesUsingPOSTParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the insert associated compensation entities using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InsertAssociatedCompensationEntitiesUsingPOSTParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the insert associated compensation entities using p o s t params
func (o *InsertAssociatedCompensationEntitiesUsingPOSTParams) WithTimeout(timeout time.Duration) *InsertAssociatedCompensationEntitiesUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the insert associated compensation entities using p o s t params
func (o *InsertAssociatedCompensationEntitiesUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the insert associated compensation entities using p o s t params
func (o *InsertAssociatedCompensationEntitiesUsingPOSTParams) WithContext(ctx context.Context) *InsertAssociatedCompensationEntitiesUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the insert associated compensation entities using p o s t params
func (o *InsertAssociatedCompensationEntitiesUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the insert associated compensation entities using p o s t params
func (o *InsertAssociatedCompensationEntitiesUsingPOSTParams) WithHTTPClient(client *http.Client) *InsertAssociatedCompensationEntitiesUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the insert associated compensation entities using p o s t params
func (o *InsertAssociatedCompensationEntitiesUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCompensationEntityList adds the compensationEntityList to the insert associated compensation entities using p o s t params
func (o *InsertAssociatedCompensationEntitiesUsingPOSTParams) WithCompensationEntityList(compensationEntityList []*models.CompensationEntityDTO) *InsertAssociatedCompensationEntitiesUsingPOSTParams {
	o.SetCompensationEntityList(compensationEntityList)
	return o
}

// SetCompensationEntityList adds the compensationEntityList to the insert associated compensation entities using p o s t params
func (o *InsertAssociatedCompensationEntitiesUsingPOSTParams) SetCompensationEntityList(compensationEntityList []*models.CompensationEntityDTO) {
	o.CompensationEntityList = compensationEntityList
}

// WithEffectiveDate adds the effectiveDate to the insert associated compensation entities using p o s t params
func (o *InsertAssociatedCompensationEntitiesUsingPOSTParams) WithEffectiveDate(effectiveDate *string) *InsertAssociatedCompensationEntitiesUsingPOSTParams {
	o.SetEffectiveDate(effectiveDate)
	return o
}

// SetEffectiveDate adds the effectiveDate to the insert associated compensation entities using p o s t params
func (o *InsertAssociatedCompensationEntitiesUsingPOSTParams) SetEffectiveDate(effectiveDate *string) {
	o.EffectiveDate = effectiveDate
}

// WithPolicyNumber adds the policyNumber to the insert associated compensation entities using p o s t params
func (o *InsertAssociatedCompensationEntitiesUsingPOSTParams) WithPolicyNumber(policyNumber *string) *InsertAssociatedCompensationEntitiesUsingPOSTParams {
	o.SetPolicyNumber(policyNumber)
	return o
}

// SetPolicyNumber adds the policyNumber to the insert associated compensation entities using p o s t params
func (o *InsertAssociatedCompensationEntitiesUsingPOSTParams) SetPolicyNumber(policyNumber *string) {
	o.PolicyNumber = policyNumber
}

// WithRequestApplication adds the requestApplication to the insert associated compensation entities using p o s t params
func (o *InsertAssociatedCompensationEntitiesUsingPOSTParams) WithRequestApplication(requestApplication *string) *InsertAssociatedCompensationEntitiesUsingPOSTParams {
	o.SetRequestApplication(requestApplication)
	return o
}

// SetRequestApplication adds the requestApplication to the insert associated compensation entities using p o s t params
func (o *InsertAssociatedCompensationEntitiesUsingPOSTParams) SetRequestApplication(requestApplication *string) {
	o.RequestApplication = requestApplication
}

// WithRequestUser adds the requestUser to the insert associated compensation entities using p o s t params
func (o *InsertAssociatedCompensationEntitiesUsingPOSTParams) WithRequestUser(requestUser *string) *InsertAssociatedCompensationEntitiesUsingPOSTParams {
	o.SetRequestUser(requestUser)
	return o
}

// SetRequestUser adds the requestUser to the insert associated compensation entities using p o s t params
func (o *InsertAssociatedCompensationEntitiesUsingPOSTParams) SetRequestUser(requestUser *string) {
	o.RequestUser = requestUser
}

// WithVarianceFormat adds the varianceFormat to the insert associated compensation entities using p o s t params
func (o *InsertAssociatedCompensationEntitiesUsingPOSTParams) WithVarianceFormat(varianceFormat *string) *InsertAssociatedCompensationEntitiesUsingPOSTParams {
	o.SetVarianceFormat(varianceFormat)
	return o
}

// SetVarianceFormat adds the varianceFormat to the insert associated compensation entities using p o s t params
func (o *InsertAssociatedCompensationEntitiesUsingPOSTParams) SetVarianceFormat(varianceFormat *string) {
	o.VarianceFormat = varianceFormat
}

// WithVarianceLevel adds the varianceLevel to the insert associated compensation entities using p o s t params
func (o *InsertAssociatedCompensationEntitiesUsingPOSTParams) WithVarianceLevel(varianceLevel *string) *InsertAssociatedCompensationEntitiesUsingPOSTParams {
	o.SetVarianceLevel(varianceLevel)
	return o
}

// SetVarianceLevel adds the varianceLevel to the insert associated compensation entities using p o s t params
func (o *InsertAssociatedCompensationEntitiesUsingPOSTParams) SetVarianceLevel(varianceLevel *string) {
	o.VarianceLevel = varianceLevel
}

// WriteToRequest writes these params to a swagger request
func (o *InsertAssociatedCompensationEntitiesUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.CompensationEntityList != nil {
		if err := r.SetBodyParam(o.CompensationEntityList); err != nil {
			return err
		}
	}

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
