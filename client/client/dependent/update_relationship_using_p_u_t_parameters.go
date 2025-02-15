// Code generated by go-swagger; DO NOT EDIT.

package dependent

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

// NewUpdateRelationshipUsingPUTParams creates a new UpdateRelationshipUsingPUTParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateRelationshipUsingPUTParams() *UpdateRelationshipUsingPUTParams {
	return &UpdateRelationshipUsingPUTParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateRelationshipUsingPUTParamsWithTimeout creates a new UpdateRelationshipUsingPUTParams object
// with the ability to set a timeout on a request.
func NewUpdateRelationshipUsingPUTParamsWithTimeout(timeout time.Duration) *UpdateRelationshipUsingPUTParams {
	return &UpdateRelationshipUsingPUTParams{
		timeout: timeout,
	}
}

// NewUpdateRelationshipUsingPUTParamsWithContext creates a new UpdateRelationshipUsingPUTParams object
// with the ability to set a context for a request.
func NewUpdateRelationshipUsingPUTParamsWithContext(ctx context.Context) *UpdateRelationshipUsingPUTParams {
	return &UpdateRelationshipUsingPUTParams{
		Context: ctx,
	}
}

// NewUpdateRelationshipUsingPUTParamsWithHTTPClient creates a new UpdateRelationshipUsingPUTParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateRelationshipUsingPUTParamsWithHTTPClient(client *http.Client) *UpdateRelationshipUsingPUTParams {
	return &UpdateRelationshipUsingPUTParams{
		HTTPClient: client,
	}
}

/*
UpdateRelationshipUsingPUTParams contains all the parameters to send to the API endpoint

	for the update relationship using p u t operation.

	Typically these are written to a http.Request.
*/
type UpdateRelationshipUsingPUTParams struct {

	/* CaseMemberKey.

	   caseMemberKey
	*/
	CaseMemberKey *string

	/* DependentClientID.

	   dependentClientId
	*/
	DependentClientID *string

	/* EffectiveDate.

	   effectiveDate
	*/
	EffectiveDate *string

	/* PerformReallocation.

	   performReallocation
	*/
	PerformReallocation *bool

	/* Relationship.

	   relationship
	*/
	Relationship *string

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

// WithDefaults hydrates default values in the update relationship using p u t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateRelationshipUsingPUTParams) WithDefaults() *UpdateRelationshipUsingPUTParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update relationship using p u t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateRelationshipUsingPUTParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update relationship using p u t params
func (o *UpdateRelationshipUsingPUTParams) WithTimeout(timeout time.Duration) *UpdateRelationshipUsingPUTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update relationship using p u t params
func (o *UpdateRelationshipUsingPUTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update relationship using p u t params
func (o *UpdateRelationshipUsingPUTParams) WithContext(ctx context.Context) *UpdateRelationshipUsingPUTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update relationship using p u t params
func (o *UpdateRelationshipUsingPUTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update relationship using p u t params
func (o *UpdateRelationshipUsingPUTParams) WithHTTPClient(client *http.Client) *UpdateRelationshipUsingPUTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update relationship using p u t params
func (o *UpdateRelationshipUsingPUTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCaseMemberKey adds the caseMemberKey to the update relationship using p u t params
func (o *UpdateRelationshipUsingPUTParams) WithCaseMemberKey(caseMemberKey *string) *UpdateRelationshipUsingPUTParams {
	o.SetCaseMemberKey(caseMemberKey)
	return o
}

// SetCaseMemberKey adds the caseMemberKey to the update relationship using p u t params
func (o *UpdateRelationshipUsingPUTParams) SetCaseMemberKey(caseMemberKey *string) {
	o.CaseMemberKey = caseMemberKey
}

// WithDependentClientID adds the dependentClientID to the update relationship using p u t params
func (o *UpdateRelationshipUsingPUTParams) WithDependentClientID(dependentClientID *string) *UpdateRelationshipUsingPUTParams {
	o.SetDependentClientID(dependentClientID)
	return o
}

// SetDependentClientID adds the dependentClientId to the update relationship using p u t params
func (o *UpdateRelationshipUsingPUTParams) SetDependentClientID(dependentClientID *string) {
	o.DependentClientID = dependentClientID
}

// WithEffectiveDate adds the effectiveDate to the update relationship using p u t params
func (o *UpdateRelationshipUsingPUTParams) WithEffectiveDate(effectiveDate *string) *UpdateRelationshipUsingPUTParams {
	o.SetEffectiveDate(effectiveDate)
	return o
}

// SetEffectiveDate adds the effectiveDate to the update relationship using p u t params
func (o *UpdateRelationshipUsingPUTParams) SetEffectiveDate(effectiveDate *string) {
	o.EffectiveDate = effectiveDate
}

// WithPerformReallocation adds the performReallocation to the update relationship using p u t params
func (o *UpdateRelationshipUsingPUTParams) WithPerformReallocation(performReallocation *bool) *UpdateRelationshipUsingPUTParams {
	o.SetPerformReallocation(performReallocation)
	return o
}

// SetPerformReallocation adds the performReallocation to the update relationship using p u t params
func (o *UpdateRelationshipUsingPUTParams) SetPerformReallocation(performReallocation *bool) {
	o.PerformReallocation = performReallocation
}

// WithRelationship adds the relationship to the update relationship using p u t params
func (o *UpdateRelationshipUsingPUTParams) WithRelationship(relationship *string) *UpdateRelationshipUsingPUTParams {
	o.SetRelationship(relationship)
	return o
}

// SetRelationship adds the relationship to the update relationship using p u t params
func (o *UpdateRelationshipUsingPUTParams) SetRelationship(relationship *string) {
	o.Relationship = relationship
}

// WithRequestApplication adds the requestApplication to the update relationship using p u t params
func (o *UpdateRelationshipUsingPUTParams) WithRequestApplication(requestApplication *string) *UpdateRelationshipUsingPUTParams {
	o.SetRequestApplication(requestApplication)
	return o
}

// SetRequestApplication adds the requestApplication to the update relationship using p u t params
func (o *UpdateRelationshipUsingPUTParams) SetRequestApplication(requestApplication *string) {
	o.RequestApplication = requestApplication
}

// WithRequestUser adds the requestUser to the update relationship using p u t params
func (o *UpdateRelationshipUsingPUTParams) WithRequestUser(requestUser *string) *UpdateRelationshipUsingPUTParams {
	o.SetRequestUser(requestUser)
	return o
}

// SetRequestUser adds the requestUser to the update relationship using p u t params
func (o *UpdateRelationshipUsingPUTParams) SetRequestUser(requestUser *string) {
	o.RequestUser = requestUser
}

// WithVarianceFormat adds the varianceFormat to the update relationship using p u t params
func (o *UpdateRelationshipUsingPUTParams) WithVarianceFormat(varianceFormat *string) *UpdateRelationshipUsingPUTParams {
	o.SetVarianceFormat(varianceFormat)
	return o
}

// SetVarianceFormat adds the varianceFormat to the update relationship using p u t params
func (o *UpdateRelationshipUsingPUTParams) SetVarianceFormat(varianceFormat *string) {
	o.VarianceFormat = varianceFormat
}

// WithVarianceLevel adds the varianceLevel to the update relationship using p u t params
func (o *UpdateRelationshipUsingPUTParams) WithVarianceLevel(varianceLevel *string) *UpdateRelationshipUsingPUTParams {
	o.SetVarianceLevel(varianceLevel)
	return o
}

// SetVarianceLevel adds the varianceLevel to the update relationship using p u t params
func (o *UpdateRelationshipUsingPUTParams) SetVarianceLevel(varianceLevel *string) {
	o.VarianceLevel = varianceLevel
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateRelationshipUsingPUTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if o.DependentClientID != nil {

		// query param dependentClientId
		var qrDependentClientID string

		if o.DependentClientID != nil {
			qrDependentClientID = *o.DependentClientID
		}
		qDependentClientID := qrDependentClientID
		if qDependentClientID != "" {

			if err := r.SetQueryParam("dependentClientId", qDependentClientID); err != nil {
				return err
			}
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

	if o.Relationship != nil {

		// query param relationship
		var qrRelationship string

		if o.Relationship != nil {
			qrRelationship = *o.Relationship
		}
		qRelationship := qrRelationship
		if qRelationship != "" {

			if err := r.SetQueryParam("relationship", qRelationship); err != nil {
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
