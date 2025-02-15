// Code generated by go-swagger; DO NOT EDIT.

package person

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

// NewUpdateAddressUsingPUT1Params creates a new UpdateAddressUsingPUT1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateAddressUsingPUT1Params() *UpdateAddressUsingPUT1Params {
	return &UpdateAddressUsingPUT1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateAddressUsingPUT1ParamsWithTimeout creates a new UpdateAddressUsingPUT1Params object
// with the ability to set a timeout on a request.
func NewUpdateAddressUsingPUT1ParamsWithTimeout(timeout time.Duration) *UpdateAddressUsingPUT1Params {
	return &UpdateAddressUsingPUT1Params{
		timeout: timeout,
	}
}

// NewUpdateAddressUsingPUT1ParamsWithContext creates a new UpdateAddressUsingPUT1Params object
// with the ability to set a context for a request.
func NewUpdateAddressUsingPUT1ParamsWithContext(ctx context.Context) *UpdateAddressUsingPUT1Params {
	return &UpdateAddressUsingPUT1Params{
		Context: ctx,
	}
}

// NewUpdateAddressUsingPUT1ParamsWithHTTPClient creates a new UpdateAddressUsingPUT1Params object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateAddressUsingPUT1ParamsWithHTTPClient(client *http.Client) *UpdateAddressUsingPUT1Params {
	return &UpdateAddressUsingPUT1Params{
		HTTPClient: client,
	}
}

/*
UpdateAddressUsingPUT1Params contains all the parameters to send to the API endpoint

	for the update address using p u t 1 operation.

	Typically these are written to a http.Request.
*/
type UpdateAddressUsingPUT1Params struct {

	/* ClientID.

	   clientId
	*/
	ClientID *string

	/* ContactDetails.

	   contactDetails
	*/
	ContactDetails *models.ContactDTO

	/* EffectiveDate.

	   effectiveDate
	*/
	EffectiveDate *string

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

// WithDefaults hydrates default values in the update address using p u t 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateAddressUsingPUT1Params) WithDefaults() *UpdateAddressUsingPUT1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update address using p u t 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateAddressUsingPUT1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update address using p u t 1 params
func (o *UpdateAddressUsingPUT1Params) WithTimeout(timeout time.Duration) *UpdateAddressUsingPUT1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update address using p u t 1 params
func (o *UpdateAddressUsingPUT1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update address using p u t 1 params
func (o *UpdateAddressUsingPUT1Params) WithContext(ctx context.Context) *UpdateAddressUsingPUT1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update address using p u t 1 params
func (o *UpdateAddressUsingPUT1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update address using p u t 1 params
func (o *UpdateAddressUsingPUT1Params) WithHTTPClient(client *http.Client) *UpdateAddressUsingPUT1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update address using p u t 1 params
func (o *UpdateAddressUsingPUT1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClientID adds the clientID to the update address using p u t 1 params
func (o *UpdateAddressUsingPUT1Params) WithClientID(clientID *string) *UpdateAddressUsingPUT1Params {
	o.SetClientID(clientID)
	return o
}

// SetClientID adds the clientId to the update address using p u t 1 params
func (o *UpdateAddressUsingPUT1Params) SetClientID(clientID *string) {
	o.ClientID = clientID
}

// WithContactDetails adds the contactDetails to the update address using p u t 1 params
func (o *UpdateAddressUsingPUT1Params) WithContactDetails(contactDetails *models.ContactDTO) *UpdateAddressUsingPUT1Params {
	o.SetContactDetails(contactDetails)
	return o
}

// SetContactDetails adds the contactDetails to the update address using p u t 1 params
func (o *UpdateAddressUsingPUT1Params) SetContactDetails(contactDetails *models.ContactDTO) {
	o.ContactDetails = contactDetails
}

// WithEffectiveDate adds the effectiveDate to the update address using p u t 1 params
func (o *UpdateAddressUsingPUT1Params) WithEffectiveDate(effectiveDate *string) *UpdateAddressUsingPUT1Params {
	o.SetEffectiveDate(effectiveDate)
	return o
}

// SetEffectiveDate adds the effectiveDate to the update address using p u t 1 params
func (o *UpdateAddressUsingPUT1Params) SetEffectiveDate(effectiveDate *string) {
	o.EffectiveDate = effectiveDate
}

// WithRequestApplication adds the requestApplication to the update address using p u t 1 params
func (o *UpdateAddressUsingPUT1Params) WithRequestApplication(requestApplication *string) *UpdateAddressUsingPUT1Params {
	o.SetRequestApplication(requestApplication)
	return o
}

// SetRequestApplication adds the requestApplication to the update address using p u t 1 params
func (o *UpdateAddressUsingPUT1Params) SetRequestApplication(requestApplication *string) {
	o.RequestApplication = requestApplication
}

// WithRequestUser adds the requestUser to the update address using p u t 1 params
func (o *UpdateAddressUsingPUT1Params) WithRequestUser(requestUser *string) *UpdateAddressUsingPUT1Params {
	o.SetRequestUser(requestUser)
	return o
}

// SetRequestUser adds the requestUser to the update address using p u t 1 params
func (o *UpdateAddressUsingPUT1Params) SetRequestUser(requestUser *string) {
	o.RequestUser = requestUser
}

// WithVarianceFormat adds the varianceFormat to the update address using p u t 1 params
func (o *UpdateAddressUsingPUT1Params) WithVarianceFormat(varianceFormat *string) *UpdateAddressUsingPUT1Params {
	o.SetVarianceFormat(varianceFormat)
	return o
}

// SetVarianceFormat adds the varianceFormat to the update address using p u t 1 params
func (o *UpdateAddressUsingPUT1Params) SetVarianceFormat(varianceFormat *string) {
	o.VarianceFormat = varianceFormat
}

// WithVarianceLevel adds the varianceLevel to the update address using p u t 1 params
func (o *UpdateAddressUsingPUT1Params) WithVarianceLevel(varianceLevel *string) *UpdateAddressUsingPUT1Params {
	o.SetVarianceLevel(varianceLevel)
	return o
}

// SetVarianceLevel adds the varianceLevel to the update address using p u t 1 params
func (o *UpdateAddressUsingPUT1Params) SetVarianceLevel(varianceLevel *string) {
	o.VarianceLevel = varianceLevel
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateAddressUsingPUT1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ClientID != nil {

		// query param clientId
		var qrClientID string

		if o.ClientID != nil {
			qrClientID = *o.ClientID
		}
		qClientID := qrClientID
		if qClientID != "" {

			if err := r.SetQueryParam("clientId", qClientID); err != nil {
				return err
			}
		}
	}
	if o.ContactDetails != nil {
		if err := r.SetBodyParam(o.ContactDetails); err != nil {
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
