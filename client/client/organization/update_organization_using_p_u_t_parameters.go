// Code generated by go-swagger; DO NOT EDIT.

package organization

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

// NewUpdateOrganizationUsingPUTParams creates a new UpdateOrganizationUsingPUTParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateOrganizationUsingPUTParams() *UpdateOrganizationUsingPUTParams {
	return &UpdateOrganizationUsingPUTParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateOrganizationUsingPUTParamsWithTimeout creates a new UpdateOrganizationUsingPUTParams object
// with the ability to set a timeout on a request.
func NewUpdateOrganizationUsingPUTParamsWithTimeout(timeout time.Duration) *UpdateOrganizationUsingPUTParams {
	return &UpdateOrganizationUsingPUTParams{
		timeout: timeout,
	}
}

// NewUpdateOrganizationUsingPUTParamsWithContext creates a new UpdateOrganizationUsingPUTParams object
// with the ability to set a context for a request.
func NewUpdateOrganizationUsingPUTParamsWithContext(ctx context.Context) *UpdateOrganizationUsingPUTParams {
	return &UpdateOrganizationUsingPUTParams{
		Context: ctx,
	}
}

// NewUpdateOrganizationUsingPUTParamsWithHTTPClient creates a new UpdateOrganizationUsingPUTParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateOrganizationUsingPUTParamsWithHTTPClient(client *http.Client) *UpdateOrganizationUsingPUTParams {
	return &UpdateOrganizationUsingPUTParams{
		HTTPClient: client,
	}
}

/*
UpdateOrganizationUsingPUTParams contains all the parameters to send to the API endpoint

	for the update organization using p u t operation.

	Typically these are written to a http.Request.
*/
type UpdateOrganizationUsingPUTParams struct {

	/* OrganizationDto.

	   organizationDto
	*/
	OrganizationDto *models.OrganizationDTO

	/* RequestApplication.

	   requestApplication
	*/
	RequestApplication *string

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

// WithDefaults hydrates default values in the update organization using p u t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateOrganizationUsingPUTParams) WithDefaults() *UpdateOrganizationUsingPUTParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update organization using p u t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateOrganizationUsingPUTParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update organization using p u t params
func (o *UpdateOrganizationUsingPUTParams) WithTimeout(timeout time.Duration) *UpdateOrganizationUsingPUTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update organization using p u t params
func (o *UpdateOrganizationUsingPUTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update organization using p u t params
func (o *UpdateOrganizationUsingPUTParams) WithContext(ctx context.Context) *UpdateOrganizationUsingPUTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update organization using p u t params
func (o *UpdateOrganizationUsingPUTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update organization using p u t params
func (o *UpdateOrganizationUsingPUTParams) WithHTTPClient(client *http.Client) *UpdateOrganizationUsingPUTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update organization using p u t params
func (o *UpdateOrganizationUsingPUTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationDto adds the organizationDto to the update organization using p u t params
func (o *UpdateOrganizationUsingPUTParams) WithOrganizationDto(organizationDto *models.OrganizationDTO) *UpdateOrganizationUsingPUTParams {
	o.SetOrganizationDto(organizationDto)
	return o
}

// SetOrganizationDto adds the organizationDto to the update organization using p u t params
func (o *UpdateOrganizationUsingPUTParams) SetOrganizationDto(organizationDto *models.OrganizationDTO) {
	o.OrganizationDto = organizationDto
}

// WithRequestApplication adds the requestApplication to the update organization using p u t params
func (o *UpdateOrganizationUsingPUTParams) WithRequestApplication(requestApplication *string) *UpdateOrganizationUsingPUTParams {
	o.SetRequestApplication(requestApplication)
	return o
}

// SetRequestApplication adds the requestApplication to the update organization using p u t params
func (o *UpdateOrganizationUsingPUTParams) SetRequestApplication(requestApplication *string) {
	o.RequestApplication = requestApplication
}

// WithRequestUser adds the requestUser to the update organization using p u t params
func (o *UpdateOrganizationUsingPUTParams) WithRequestUser(requestUser string) *UpdateOrganizationUsingPUTParams {
	o.SetRequestUser(requestUser)
	return o
}

// SetRequestUser adds the requestUser to the update organization using p u t params
func (o *UpdateOrganizationUsingPUTParams) SetRequestUser(requestUser string) {
	o.RequestUser = requestUser
}

// WithVarianceFormat adds the varianceFormat to the update organization using p u t params
func (o *UpdateOrganizationUsingPUTParams) WithVarianceFormat(varianceFormat *string) *UpdateOrganizationUsingPUTParams {
	o.SetVarianceFormat(varianceFormat)
	return o
}

// SetVarianceFormat adds the varianceFormat to the update organization using p u t params
func (o *UpdateOrganizationUsingPUTParams) SetVarianceFormat(varianceFormat *string) {
	o.VarianceFormat = varianceFormat
}

// WithVarianceLevel adds the varianceLevel to the update organization using p u t params
func (o *UpdateOrganizationUsingPUTParams) WithVarianceLevel(varianceLevel *string) *UpdateOrganizationUsingPUTParams {
	o.SetVarianceLevel(varianceLevel)
	return o
}

// SetVarianceLevel adds the varianceLevel to the update organization using p u t params
func (o *UpdateOrganizationUsingPUTParams) SetVarianceLevel(varianceLevel *string) {
	o.VarianceLevel = varianceLevel
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateOrganizationUsingPUTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.OrganizationDto != nil {
		if err := r.SetBodyParam(o.OrganizationDto); err != nil {
			return err
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
