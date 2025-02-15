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
	"github.com/go-openapi/swag"
)

// NewUpdateUnderwritingDataCodeUsingPUTParams creates a new UpdateUnderwritingDataCodeUsingPUTParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateUnderwritingDataCodeUsingPUTParams() *UpdateUnderwritingDataCodeUsingPUTParams {
	return &UpdateUnderwritingDataCodeUsingPUTParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateUnderwritingDataCodeUsingPUTParamsWithTimeout creates a new UpdateUnderwritingDataCodeUsingPUTParams object
// with the ability to set a timeout on a request.
func NewUpdateUnderwritingDataCodeUsingPUTParamsWithTimeout(timeout time.Duration) *UpdateUnderwritingDataCodeUsingPUTParams {
	return &UpdateUnderwritingDataCodeUsingPUTParams{
		timeout: timeout,
	}
}

// NewUpdateUnderwritingDataCodeUsingPUTParamsWithContext creates a new UpdateUnderwritingDataCodeUsingPUTParams object
// with the ability to set a context for a request.
func NewUpdateUnderwritingDataCodeUsingPUTParamsWithContext(ctx context.Context) *UpdateUnderwritingDataCodeUsingPUTParams {
	return &UpdateUnderwritingDataCodeUsingPUTParams{
		Context: ctx,
	}
}

// NewUpdateUnderwritingDataCodeUsingPUTParamsWithHTTPClient creates a new UpdateUnderwritingDataCodeUsingPUTParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateUnderwritingDataCodeUsingPUTParamsWithHTTPClient(client *http.Client) *UpdateUnderwritingDataCodeUsingPUTParams {
	return &UpdateUnderwritingDataCodeUsingPUTParams{
		HTTPClient: client,
	}
}

/*
UpdateUnderwritingDataCodeUsingPUTParams contains all the parameters to send to the API endpoint

	for the update underwriting data code using p u t operation.

	Typically these are written to a http.Request.
*/
type UpdateUnderwritingDataCodeUsingPUTParams struct {

	/* BenDtrmCd.

	   benDtrmCd
	*/
	BenDtrmCd *string

	/* BenefitDescription.

	   benefitDescription
	*/
	BenefitDescription string

	/* DependentData.

	   dependentData
	*/
	DependentData *string

	/* EffectiveDate.

	   effectiveDate
	*/
	EffectiveDate *string

	/* IsOverridden.

	   isOverridden
	*/
	IsOverridden *bool

	/* MemberGroupDescription.

	   memberGroupDescription
	*/
	MemberGroupDescription *string

	/* PolicyNumber.

	   policyNumber
	*/
	PolicyNumber string

	/* PrrtCd.

	   prrtCd
	*/
	PrrtCd *string

	/* RequestApplication.

	   requestApplication
	*/
	RequestApplication string

	/* RequestUser.

	   requestUser
	*/
	RequestUser string

	/* UnderwritingData.

	   underwritingData
	*/
	UnderwritingData *string

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

// WithDefaults hydrates default values in the update underwriting data code using p u t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateUnderwritingDataCodeUsingPUTParams) WithDefaults() *UpdateUnderwritingDataCodeUsingPUTParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update underwriting data code using p u t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateUnderwritingDataCodeUsingPUTParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update underwriting data code using p u t params
func (o *UpdateUnderwritingDataCodeUsingPUTParams) WithTimeout(timeout time.Duration) *UpdateUnderwritingDataCodeUsingPUTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update underwriting data code using p u t params
func (o *UpdateUnderwritingDataCodeUsingPUTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update underwriting data code using p u t params
func (o *UpdateUnderwritingDataCodeUsingPUTParams) WithContext(ctx context.Context) *UpdateUnderwritingDataCodeUsingPUTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update underwriting data code using p u t params
func (o *UpdateUnderwritingDataCodeUsingPUTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update underwriting data code using p u t params
func (o *UpdateUnderwritingDataCodeUsingPUTParams) WithHTTPClient(client *http.Client) *UpdateUnderwritingDataCodeUsingPUTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update underwriting data code using p u t params
func (o *UpdateUnderwritingDataCodeUsingPUTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBenDtrmCd adds the benDtrmCd to the update underwriting data code using p u t params
func (o *UpdateUnderwritingDataCodeUsingPUTParams) WithBenDtrmCd(benDtrmCd *string) *UpdateUnderwritingDataCodeUsingPUTParams {
	o.SetBenDtrmCd(benDtrmCd)
	return o
}

// SetBenDtrmCd adds the benDtrmCd to the update underwriting data code using p u t params
func (o *UpdateUnderwritingDataCodeUsingPUTParams) SetBenDtrmCd(benDtrmCd *string) {
	o.BenDtrmCd = benDtrmCd
}

// WithBenefitDescription adds the benefitDescription to the update underwriting data code using p u t params
func (o *UpdateUnderwritingDataCodeUsingPUTParams) WithBenefitDescription(benefitDescription string) *UpdateUnderwritingDataCodeUsingPUTParams {
	o.SetBenefitDescription(benefitDescription)
	return o
}

// SetBenefitDescription adds the benefitDescription to the update underwriting data code using p u t params
func (o *UpdateUnderwritingDataCodeUsingPUTParams) SetBenefitDescription(benefitDescription string) {
	o.BenefitDescription = benefitDescription
}

// WithDependentData adds the dependentData to the update underwriting data code using p u t params
func (o *UpdateUnderwritingDataCodeUsingPUTParams) WithDependentData(dependentData *string) *UpdateUnderwritingDataCodeUsingPUTParams {
	o.SetDependentData(dependentData)
	return o
}

// SetDependentData adds the dependentData to the update underwriting data code using p u t params
func (o *UpdateUnderwritingDataCodeUsingPUTParams) SetDependentData(dependentData *string) {
	o.DependentData = dependentData
}

// WithEffectiveDate adds the effectiveDate to the update underwriting data code using p u t params
func (o *UpdateUnderwritingDataCodeUsingPUTParams) WithEffectiveDate(effectiveDate *string) *UpdateUnderwritingDataCodeUsingPUTParams {
	o.SetEffectiveDate(effectiveDate)
	return o
}

// SetEffectiveDate adds the effectiveDate to the update underwriting data code using p u t params
func (o *UpdateUnderwritingDataCodeUsingPUTParams) SetEffectiveDate(effectiveDate *string) {
	o.EffectiveDate = effectiveDate
}

// WithIsOverridden adds the isOverridden to the update underwriting data code using p u t params
func (o *UpdateUnderwritingDataCodeUsingPUTParams) WithIsOverridden(isOverridden *bool) *UpdateUnderwritingDataCodeUsingPUTParams {
	o.SetIsOverridden(isOverridden)
	return o
}

// SetIsOverridden adds the isOverridden to the update underwriting data code using p u t params
func (o *UpdateUnderwritingDataCodeUsingPUTParams) SetIsOverridden(isOverridden *bool) {
	o.IsOverridden = isOverridden
}

// WithMemberGroupDescription adds the memberGroupDescription to the update underwriting data code using p u t params
func (o *UpdateUnderwritingDataCodeUsingPUTParams) WithMemberGroupDescription(memberGroupDescription *string) *UpdateUnderwritingDataCodeUsingPUTParams {
	o.SetMemberGroupDescription(memberGroupDescription)
	return o
}

// SetMemberGroupDescription adds the memberGroupDescription to the update underwriting data code using p u t params
func (o *UpdateUnderwritingDataCodeUsingPUTParams) SetMemberGroupDescription(memberGroupDescription *string) {
	o.MemberGroupDescription = memberGroupDescription
}

// WithPolicyNumber adds the policyNumber to the update underwriting data code using p u t params
func (o *UpdateUnderwritingDataCodeUsingPUTParams) WithPolicyNumber(policyNumber string) *UpdateUnderwritingDataCodeUsingPUTParams {
	o.SetPolicyNumber(policyNumber)
	return o
}

// SetPolicyNumber adds the policyNumber to the update underwriting data code using p u t params
func (o *UpdateUnderwritingDataCodeUsingPUTParams) SetPolicyNumber(policyNumber string) {
	o.PolicyNumber = policyNumber
}

// WithPrrtCd adds the prrtCd to the update underwriting data code using p u t params
func (o *UpdateUnderwritingDataCodeUsingPUTParams) WithPrrtCd(prrtCd *string) *UpdateUnderwritingDataCodeUsingPUTParams {
	o.SetPrrtCd(prrtCd)
	return o
}

// SetPrrtCd adds the prrtCd to the update underwriting data code using p u t params
func (o *UpdateUnderwritingDataCodeUsingPUTParams) SetPrrtCd(prrtCd *string) {
	o.PrrtCd = prrtCd
}

// WithRequestApplication adds the requestApplication to the update underwriting data code using p u t params
func (o *UpdateUnderwritingDataCodeUsingPUTParams) WithRequestApplication(requestApplication string) *UpdateUnderwritingDataCodeUsingPUTParams {
	o.SetRequestApplication(requestApplication)
	return o
}

// SetRequestApplication adds the requestApplication to the update underwriting data code using p u t params
func (o *UpdateUnderwritingDataCodeUsingPUTParams) SetRequestApplication(requestApplication string) {
	o.RequestApplication = requestApplication
}

// WithRequestUser adds the requestUser to the update underwriting data code using p u t params
func (o *UpdateUnderwritingDataCodeUsingPUTParams) WithRequestUser(requestUser string) *UpdateUnderwritingDataCodeUsingPUTParams {
	o.SetRequestUser(requestUser)
	return o
}

// SetRequestUser adds the requestUser to the update underwriting data code using p u t params
func (o *UpdateUnderwritingDataCodeUsingPUTParams) SetRequestUser(requestUser string) {
	o.RequestUser = requestUser
}

// WithUnderwritingData adds the underwritingData to the update underwriting data code using p u t params
func (o *UpdateUnderwritingDataCodeUsingPUTParams) WithUnderwritingData(underwritingData *string) *UpdateUnderwritingDataCodeUsingPUTParams {
	o.SetUnderwritingData(underwritingData)
	return o
}

// SetUnderwritingData adds the underwritingData to the update underwriting data code using p u t params
func (o *UpdateUnderwritingDataCodeUsingPUTParams) SetUnderwritingData(underwritingData *string) {
	o.UnderwritingData = underwritingData
}

// WithVarianceFormat adds the varianceFormat to the update underwriting data code using p u t params
func (o *UpdateUnderwritingDataCodeUsingPUTParams) WithVarianceFormat(varianceFormat *string) *UpdateUnderwritingDataCodeUsingPUTParams {
	o.SetVarianceFormat(varianceFormat)
	return o
}

// SetVarianceFormat adds the varianceFormat to the update underwriting data code using p u t params
func (o *UpdateUnderwritingDataCodeUsingPUTParams) SetVarianceFormat(varianceFormat *string) {
	o.VarianceFormat = varianceFormat
}

// WithVarianceLevel adds the varianceLevel to the update underwriting data code using p u t params
func (o *UpdateUnderwritingDataCodeUsingPUTParams) WithVarianceLevel(varianceLevel *string) *UpdateUnderwritingDataCodeUsingPUTParams {
	o.SetVarianceLevel(varianceLevel)
	return o
}

// SetVarianceLevel adds the varianceLevel to the update underwriting data code using p u t params
func (o *UpdateUnderwritingDataCodeUsingPUTParams) SetVarianceLevel(varianceLevel *string) {
	o.VarianceLevel = varianceLevel
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateUnderwritingDataCodeUsingPUTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.BenDtrmCd != nil {

		// query param benDtrmCd
		var qrBenDtrmCd string

		if o.BenDtrmCd != nil {
			qrBenDtrmCd = *o.BenDtrmCd
		}
		qBenDtrmCd := qrBenDtrmCd
		if qBenDtrmCd != "" {

			if err := r.SetQueryParam("benDtrmCd", qBenDtrmCd); err != nil {
				return err
			}
		}
	}

	// query param benefitDescription
	qrBenefitDescription := o.BenefitDescription
	qBenefitDescription := qrBenefitDescription
	if qBenefitDescription != "" {

		if err := r.SetQueryParam("benefitDescription", qBenefitDescription); err != nil {
			return err
		}
	}

	if o.DependentData != nil {

		// query param dependentData
		var qrDependentData string

		if o.DependentData != nil {
			qrDependentData = *o.DependentData
		}
		qDependentData := qrDependentData
		if qDependentData != "" {

			if err := r.SetQueryParam("dependentData", qDependentData); err != nil {
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

	if o.IsOverridden != nil {

		// query param isOverridden
		var qrIsOverridden bool

		if o.IsOverridden != nil {
			qrIsOverridden = *o.IsOverridden
		}
		qIsOverridden := swag.FormatBool(qrIsOverridden)
		if qIsOverridden != "" {

			if err := r.SetQueryParam("isOverridden", qIsOverridden); err != nil {
				return err
			}
		}
	}

	if o.MemberGroupDescription != nil {

		// query param memberGroupDescription
		var qrMemberGroupDescription string

		if o.MemberGroupDescription != nil {
			qrMemberGroupDescription = *o.MemberGroupDescription
		}
		qMemberGroupDescription := qrMemberGroupDescription
		if qMemberGroupDescription != "" {

			if err := r.SetQueryParam("memberGroupDescription", qMemberGroupDescription); err != nil {
				return err
			}
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

	if o.PrrtCd != nil {

		// query param prrtCd
		var qrPrrtCd string

		if o.PrrtCd != nil {
			qrPrrtCd = *o.PrrtCd
		}
		qPrrtCd := qrPrrtCd
		if qPrrtCd != "" {

			if err := r.SetQueryParam("prrtCd", qPrrtCd); err != nil {
				return err
			}
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

	if o.UnderwritingData != nil {

		// query param underwritingData
		var qrUnderwritingData string

		if o.UnderwritingData != nil {
			qrUnderwritingData = *o.UnderwritingData
		}
		qUnderwritingData := qrUnderwritingData
		if qUnderwritingData != "" {

			if err := r.SetQueryParam("underwritingData", qUnderwritingData); err != nil {
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
