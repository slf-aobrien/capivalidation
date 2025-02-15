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
)

// NewUpdateMemberLanguageUsingPUTParams creates a new UpdateMemberLanguageUsingPUTParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateMemberLanguageUsingPUTParams() *UpdateMemberLanguageUsingPUTParams {
	return &UpdateMemberLanguageUsingPUTParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateMemberLanguageUsingPUTParamsWithTimeout creates a new UpdateMemberLanguageUsingPUTParams object
// with the ability to set a timeout on a request.
func NewUpdateMemberLanguageUsingPUTParamsWithTimeout(timeout time.Duration) *UpdateMemberLanguageUsingPUTParams {
	return &UpdateMemberLanguageUsingPUTParams{
		timeout: timeout,
	}
}

// NewUpdateMemberLanguageUsingPUTParamsWithContext creates a new UpdateMemberLanguageUsingPUTParams object
// with the ability to set a context for a request.
func NewUpdateMemberLanguageUsingPUTParamsWithContext(ctx context.Context) *UpdateMemberLanguageUsingPUTParams {
	return &UpdateMemberLanguageUsingPUTParams{
		Context: ctx,
	}
}

// NewUpdateMemberLanguageUsingPUTParamsWithHTTPClient creates a new UpdateMemberLanguageUsingPUTParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateMemberLanguageUsingPUTParamsWithHTTPClient(client *http.Client) *UpdateMemberLanguageUsingPUTParams {
	return &UpdateMemberLanguageUsingPUTParams{
		HTTPClient: client,
	}
}

/*
UpdateMemberLanguageUsingPUTParams contains all the parameters to send to the API endpoint

	for the update member language using p u t operation.

	Typically these are written to a http.Request.
*/
type UpdateMemberLanguageUsingPUTParams struct {

	/* CaseMemberKey.

	   caseMemberKey
	*/
	CaseMemberKey *string

	/* EffectiveDate.

	   effectiveDate
	*/
	EffectiveDate *string

	/* Language.

	   language
	*/
	Language string

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

// WithDefaults hydrates default values in the update member language using p u t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateMemberLanguageUsingPUTParams) WithDefaults() *UpdateMemberLanguageUsingPUTParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update member language using p u t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateMemberLanguageUsingPUTParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update member language using p u t params
func (o *UpdateMemberLanguageUsingPUTParams) WithTimeout(timeout time.Duration) *UpdateMemberLanguageUsingPUTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update member language using p u t params
func (o *UpdateMemberLanguageUsingPUTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update member language using p u t params
func (o *UpdateMemberLanguageUsingPUTParams) WithContext(ctx context.Context) *UpdateMemberLanguageUsingPUTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update member language using p u t params
func (o *UpdateMemberLanguageUsingPUTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update member language using p u t params
func (o *UpdateMemberLanguageUsingPUTParams) WithHTTPClient(client *http.Client) *UpdateMemberLanguageUsingPUTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update member language using p u t params
func (o *UpdateMemberLanguageUsingPUTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCaseMemberKey adds the caseMemberKey to the update member language using p u t params
func (o *UpdateMemberLanguageUsingPUTParams) WithCaseMemberKey(caseMemberKey *string) *UpdateMemberLanguageUsingPUTParams {
	o.SetCaseMemberKey(caseMemberKey)
	return o
}

// SetCaseMemberKey adds the caseMemberKey to the update member language using p u t params
func (o *UpdateMemberLanguageUsingPUTParams) SetCaseMemberKey(caseMemberKey *string) {
	o.CaseMemberKey = caseMemberKey
}

// WithEffectiveDate adds the effectiveDate to the update member language using p u t params
func (o *UpdateMemberLanguageUsingPUTParams) WithEffectiveDate(effectiveDate *string) *UpdateMemberLanguageUsingPUTParams {
	o.SetEffectiveDate(effectiveDate)
	return o
}

// SetEffectiveDate adds the effectiveDate to the update member language using p u t params
func (o *UpdateMemberLanguageUsingPUTParams) SetEffectiveDate(effectiveDate *string) {
	o.EffectiveDate = effectiveDate
}

// WithLanguage adds the language to the update member language using p u t params
func (o *UpdateMemberLanguageUsingPUTParams) WithLanguage(language string) *UpdateMemberLanguageUsingPUTParams {
	o.SetLanguage(language)
	return o
}

// SetLanguage adds the language to the update member language using p u t params
func (o *UpdateMemberLanguageUsingPUTParams) SetLanguage(language string) {
	o.Language = language
}

// WithRequestApplication adds the requestApplication to the update member language using p u t params
func (o *UpdateMemberLanguageUsingPUTParams) WithRequestApplication(requestApplication *string) *UpdateMemberLanguageUsingPUTParams {
	o.SetRequestApplication(requestApplication)
	return o
}

// SetRequestApplication adds the requestApplication to the update member language using p u t params
func (o *UpdateMemberLanguageUsingPUTParams) SetRequestApplication(requestApplication *string) {
	o.RequestApplication = requestApplication
}

// WithRequestUser adds the requestUser to the update member language using p u t params
func (o *UpdateMemberLanguageUsingPUTParams) WithRequestUser(requestUser *string) *UpdateMemberLanguageUsingPUTParams {
	o.SetRequestUser(requestUser)
	return o
}

// SetRequestUser adds the requestUser to the update member language using p u t params
func (o *UpdateMemberLanguageUsingPUTParams) SetRequestUser(requestUser *string) {
	o.RequestUser = requestUser
}

// WithVarianceFormat adds the varianceFormat to the update member language using p u t params
func (o *UpdateMemberLanguageUsingPUTParams) WithVarianceFormat(varianceFormat *string) *UpdateMemberLanguageUsingPUTParams {
	o.SetVarianceFormat(varianceFormat)
	return o
}

// SetVarianceFormat adds the varianceFormat to the update member language using p u t params
func (o *UpdateMemberLanguageUsingPUTParams) SetVarianceFormat(varianceFormat *string) {
	o.VarianceFormat = varianceFormat
}

// WithVarianceLevel adds the varianceLevel to the update member language using p u t params
func (o *UpdateMemberLanguageUsingPUTParams) WithVarianceLevel(varianceLevel *string) *UpdateMemberLanguageUsingPUTParams {
	o.SetVarianceLevel(varianceLevel)
	return o
}

// SetVarianceLevel adds the varianceLevel to the update member language using p u t params
func (o *UpdateMemberLanguageUsingPUTParams) SetVarianceLevel(varianceLevel *string) {
	o.VarianceLevel = varianceLevel
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateMemberLanguageUsingPUTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if err := r.SetBodyParam(o.Language); err != nil {
		return err
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
