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

// NewGetMemberListCensusV2UsingGETParams creates a new GetMemberListCensusV2UsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetMemberListCensusV2UsingGETParams() *GetMemberListCensusV2UsingGETParams {
	return &GetMemberListCensusV2UsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetMemberListCensusV2UsingGETParamsWithTimeout creates a new GetMemberListCensusV2UsingGETParams object
// with the ability to set a timeout on a request.
func NewGetMemberListCensusV2UsingGETParamsWithTimeout(timeout time.Duration) *GetMemberListCensusV2UsingGETParams {
	return &GetMemberListCensusV2UsingGETParams{
		timeout: timeout,
	}
}

// NewGetMemberListCensusV2UsingGETParamsWithContext creates a new GetMemberListCensusV2UsingGETParams object
// with the ability to set a context for a request.
func NewGetMemberListCensusV2UsingGETParamsWithContext(ctx context.Context) *GetMemberListCensusV2UsingGETParams {
	return &GetMemberListCensusV2UsingGETParams{
		Context: ctx,
	}
}

// NewGetMemberListCensusV2UsingGETParamsWithHTTPClient creates a new GetMemberListCensusV2UsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetMemberListCensusV2UsingGETParamsWithHTTPClient(client *http.Client) *GetMemberListCensusV2UsingGETParams {
	return &GetMemberListCensusV2UsingGETParams{
		HTTPClient: client,
	}
}

/*
GetMemberListCensusV2UsingGETParams contains all the parameters to send to the API endpoint

	for the get member list census v2 using g e t operation.

	Typically these are written to a http.Request.
*/
type GetMemberListCensusV2UsingGETParams struct {

	/* EffectiveDate.

	   effectiveDate
	*/
	EffectiveDate *string

	/* MemberStatusCodes.

	   memberStatusCodes
	*/
	MemberStatusCodes *string

	/* PageEnd.

	   pageEnd
	*/
	PageEnd *string

	/* PageStart.

	   pageStart
	*/
	PageStart *string

	/* PolicyNumber.

	   policyNumber
	*/
	PolicyNumber *string

	/* RelationshipCodes.

	   relationshipCodes
	*/
	RelationshipCodes *string

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

// WithDefaults hydrates default values in the get member list census v2 using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMemberListCensusV2UsingGETParams) WithDefaults() *GetMemberListCensusV2UsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get member list census v2 using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMemberListCensusV2UsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get member list census v2 using g e t params
func (o *GetMemberListCensusV2UsingGETParams) WithTimeout(timeout time.Duration) *GetMemberListCensusV2UsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get member list census v2 using g e t params
func (o *GetMemberListCensusV2UsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get member list census v2 using g e t params
func (o *GetMemberListCensusV2UsingGETParams) WithContext(ctx context.Context) *GetMemberListCensusV2UsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get member list census v2 using g e t params
func (o *GetMemberListCensusV2UsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get member list census v2 using g e t params
func (o *GetMemberListCensusV2UsingGETParams) WithHTTPClient(client *http.Client) *GetMemberListCensusV2UsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get member list census v2 using g e t params
func (o *GetMemberListCensusV2UsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEffectiveDate adds the effectiveDate to the get member list census v2 using g e t params
func (o *GetMemberListCensusV2UsingGETParams) WithEffectiveDate(effectiveDate *string) *GetMemberListCensusV2UsingGETParams {
	o.SetEffectiveDate(effectiveDate)
	return o
}

// SetEffectiveDate adds the effectiveDate to the get member list census v2 using g e t params
func (o *GetMemberListCensusV2UsingGETParams) SetEffectiveDate(effectiveDate *string) {
	o.EffectiveDate = effectiveDate
}

// WithMemberStatusCodes adds the memberStatusCodes to the get member list census v2 using g e t params
func (o *GetMemberListCensusV2UsingGETParams) WithMemberStatusCodes(memberStatusCodes *string) *GetMemberListCensusV2UsingGETParams {
	o.SetMemberStatusCodes(memberStatusCodes)
	return o
}

// SetMemberStatusCodes adds the memberStatusCodes to the get member list census v2 using g e t params
func (o *GetMemberListCensusV2UsingGETParams) SetMemberStatusCodes(memberStatusCodes *string) {
	o.MemberStatusCodes = memberStatusCodes
}

// WithPageEnd adds the pageEnd to the get member list census v2 using g e t params
func (o *GetMemberListCensusV2UsingGETParams) WithPageEnd(pageEnd *string) *GetMemberListCensusV2UsingGETParams {
	o.SetPageEnd(pageEnd)
	return o
}

// SetPageEnd adds the pageEnd to the get member list census v2 using g e t params
func (o *GetMemberListCensusV2UsingGETParams) SetPageEnd(pageEnd *string) {
	o.PageEnd = pageEnd
}

// WithPageStart adds the pageStart to the get member list census v2 using g e t params
func (o *GetMemberListCensusV2UsingGETParams) WithPageStart(pageStart *string) *GetMemberListCensusV2UsingGETParams {
	o.SetPageStart(pageStart)
	return o
}

// SetPageStart adds the pageStart to the get member list census v2 using g e t params
func (o *GetMemberListCensusV2UsingGETParams) SetPageStart(pageStart *string) {
	o.PageStart = pageStart
}

// WithPolicyNumber adds the policyNumber to the get member list census v2 using g e t params
func (o *GetMemberListCensusV2UsingGETParams) WithPolicyNumber(policyNumber *string) *GetMemberListCensusV2UsingGETParams {
	o.SetPolicyNumber(policyNumber)
	return o
}

// SetPolicyNumber adds the policyNumber to the get member list census v2 using g e t params
func (o *GetMemberListCensusV2UsingGETParams) SetPolicyNumber(policyNumber *string) {
	o.PolicyNumber = policyNumber
}

// WithRelationshipCodes adds the relationshipCodes to the get member list census v2 using g e t params
func (o *GetMemberListCensusV2UsingGETParams) WithRelationshipCodes(relationshipCodes *string) *GetMemberListCensusV2UsingGETParams {
	o.SetRelationshipCodes(relationshipCodes)
	return o
}

// SetRelationshipCodes adds the relationshipCodes to the get member list census v2 using g e t params
func (o *GetMemberListCensusV2UsingGETParams) SetRelationshipCodes(relationshipCodes *string) {
	o.RelationshipCodes = relationshipCodes
}

// WithRequestApplication adds the requestApplication to the get member list census v2 using g e t params
func (o *GetMemberListCensusV2UsingGETParams) WithRequestApplication(requestApplication *string) *GetMemberListCensusV2UsingGETParams {
	o.SetRequestApplication(requestApplication)
	return o
}

// SetRequestApplication adds the requestApplication to the get member list census v2 using g e t params
func (o *GetMemberListCensusV2UsingGETParams) SetRequestApplication(requestApplication *string) {
	o.RequestApplication = requestApplication
}

// WithRequestUser adds the requestUser to the get member list census v2 using g e t params
func (o *GetMemberListCensusV2UsingGETParams) WithRequestUser(requestUser *string) *GetMemberListCensusV2UsingGETParams {
	o.SetRequestUser(requestUser)
	return o
}

// SetRequestUser adds the requestUser to the get member list census v2 using g e t params
func (o *GetMemberListCensusV2UsingGETParams) SetRequestUser(requestUser *string) {
	o.RequestUser = requestUser
}

// WithVarianceFormat adds the varianceFormat to the get member list census v2 using g e t params
func (o *GetMemberListCensusV2UsingGETParams) WithVarianceFormat(varianceFormat *string) *GetMemberListCensusV2UsingGETParams {
	o.SetVarianceFormat(varianceFormat)
	return o
}

// SetVarianceFormat adds the varianceFormat to the get member list census v2 using g e t params
func (o *GetMemberListCensusV2UsingGETParams) SetVarianceFormat(varianceFormat *string) {
	o.VarianceFormat = varianceFormat
}

// WithVarianceLevel adds the varianceLevel to the get member list census v2 using g e t params
func (o *GetMemberListCensusV2UsingGETParams) WithVarianceLevel(varianceLevel *string) *GetMemberListCensusV2UsingGETParams {
	o.SetVarianceLevel(varianceLevel)
	return o
}

// SetVarianceLevel adds the varianceLevel to the get member list census v2 using g e t params
func (o *GetMemberListCensusV2UsingGETParams) SetVarianceLevel(varianceLevel *string) {
	o.VarianceLevel = varianceLevel
}

// WriteToRequest writes these params to a swagger request
func (o *GetMemberListCensusV2UsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.MemberStatusCodes != nil {

		// query param memberStatusCodes
		var qrMemberStatusCodes string

		if o.MemberStatusCodes != nil {
			qrMemberStatusCodes = *o.MemberStatusCodes
		}
		qMemberStatusCodes := qrMemberStatusCodes
		if qMemberStatusCodes != "" {

			if err := r.SetQueryParam("memberStatusCodes", qMemberStatusCodes); err != nil {
				return err
			}
		}
	}

	if o.PageEnd != nil {

		// query param pageEnd
		var qrPageEnd string

		if o.PageEnd != nil {
			qrPageEnd = *o.PageEnd
		}
		qPageEnd := qrPageEnd
		if qPageEnd != "" {

			if err := r.SetQueryParam("pageEnd", qPageEnd); err != nil {
				return err
			}
		}
	}

	if o.PageStart != nil {

		// query param pageStart
		var qrPageStart string

		if o.PageStart != nil {
			qrPageStart = *o.PageStart
		}
		qPageStart := qrPageStart
		if qPageStart != "" {

			if err := r.SetQueryParam("pageStart", qPageStart); err != nil {
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

	if o.RelationshipCodes != nil {

		// query param relationshipCodes
		var qrRelationshipCodes string

		if o.RelationshipCodes != nil {
			qrRelationshipCodes = *o.RelationshipCodes
		}
		qRelationshipCodes := qrRelationshipCodes
		if qRelationshipCodes != "" {

			if err := r.SetQueryParam("relationshipCodes", qRelationshipCodes); err != nil {
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
