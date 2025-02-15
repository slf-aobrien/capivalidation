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

// NewFindPersonUsingGETParams creates a new FindPersonUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindPersonUsingGETParams() *FindPersonUsingGETParams {
	return &FindPersonUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindPersonUsingGETParamsWithTimeout creates a new FindPersonUsingGETParams object
// with the ability to set a timeout on a request.
func NewFindPersonUsingGETParamsWithTimeout(timeout time.Duration) *FindPersonUsingGETParams {
	return &FindPersonUsingGETParams{
		timeout: timeout,
	}
}

// NewFindPersonUsingGETParamsWithContext creates a new FindPersonUsingGETParams object
// with the ability to set a context for a request.
func NewFindPersonUsingGETParamsWithContext(ctx context.Context) *FindPersonUsingGETParams {
	return &FindPersonUsingGETParams{
		Context: ctx,
	}
}

// NewFindPersonUsingGETParamsWithHTTPClient creates a new FindPersonUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindPersonUsingGETParamsWithHTTPClient(client *http.Client) *FindPersonUsingGETParams {
	return &FindPersonUsingGETParams{
		HTTPClient: client,
	}
}

/*
FindPersonUsingGETParams contains all the parameters to send to the API endpoint

	for the find person using g e t operation.

	Typically these are written to a http.Request.
*/
type FindPersonUsingGETParams struct {

	/* BirthDate.

	   birthDate
	*/
	BirthDate *string

	/* ClientID.

	   clientId
	*/
	ClientID *string

	/* FirstName.

	   firstName
	*/
	FirstName *string

	/* LastName.

	   lastName
	*/
	LastName *string

	/* MiddleInitial.

	   middleInitial
	*/
	MiddleInitial *string

	/* RequestApplication.

	   requestApplication
	*/
	RequestApplication *string

	/* RequestUser.

	   requestUser
	*/
	RequestUser *string

	/* Ssn.

	   ssn
	*/
	Ssn *string

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

// WithDefaults hydrates default values in the find person using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindPersonUsingGETParams) WithDefaults() *FindPersonUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find person using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindPersonUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find person using g e t params
func (o *FindPersonUsingGETParams) WithTimeout(timeout time.Duration) *FindPersonUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find person using g e t params
func (o *FindPersonUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find person using g e t params
func (o *FindPersonUsingGETParams) WithContext(ctx context.Context) *FindPersonUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find person using g e t params
func (o *FindPersonUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find person using g e t params
func (o *FindPersonUsingGETParams) WithHTTPClient(client *http.Client) *FindPersonUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find person using g e t params
func (o *FindPersonUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBirthDate adds the birthDate to the find person using g e t params
func (o *FindPersonUsingGETParams) WithBirthDate(birthDate *string) *FindPersonUsingGETParams {
	o.SetBirthDate(birthDate)
	return o
}

// SetBirthDate adds the birthDate to the find person using g e t params
func (o *FindPersonUsingGETParams) SetBirthDate(birthDate *string) {
	o.BirthDate = birthDate
}

// WithClientID adds the clientID to the find person using g e t params
func (o *FindPersonUsingGETParams) WithClientID(clientID *string) *FindPersonUsingGETParams {
	o.SetClientID(clientID)
	return o
}

// SetClientID adds the clientId to the find person using g e t params
func (o *FindPersonUsingGETParams) SetClientID(clientID *string) {
	o.ClientID = clientID
}

// WithFirstName adds the firstName to the find person using g e t params
func (o *FindPersonUsingGETParams) WithFirstName(firstName *string) *FindPersonUsingGETParams {
	o.SetFirstName(firstName)
	return o
}

// SetFirstName adds the firstName to the find person using g e t params
func (o *FindPersonUsingGETParams) SetFirstName(firstName *string) {
	o.FirstName = firstName
}

// WithLastName adds the lastName to the find person using g e t params
func (o *FindPersonUsingGETParams) WithLastName(lastName *string) *FindPersonUsingGETParams {
	o.SetLastName(lastName)
	return o
}

// SetLastName adds the lastName to the find person using g e t params
func (o *FindPersonUsingGETParams) SetLastName(lastName *string) {
	o.LastName = lastName
}

// WithMiddleInitial adds the middleInitial to the find person using g e t params
func (o *FindPersonUsingGETParams) WithMiddleInitial(middleInitial *string) *FindPersonUsingGETParams {
	o.SetMiddleInitial(middleInitial)
	return o
}

// SetMiddleInitial adds the middleInitial to the find person using g e t params
func (o *FindPersonUsingGETParams) SetMiddleInitial(middleInitial *string) {
	o.MiddleInitial = middleInitial
}

// WithRequestApplication adds the requestApplication to the find person using g e t params
func (o *FindPersonUsingGETParams) WithRequestApplication(requestApplication *string) *FindPersonUsingGETParams {
	o.SetRequestApplication(requestApplication)
	return o
}

// SetRequestApplication adds the requestApplication to the find person using g e t params
func (o *FindPersonUsingGETParams) SetRequestApplication(requestApplication *string) {
	o.RequestApplication = requestApplication
}

// WithRequestUser adds the requestUser to the find person using g e t params
func (o *FindPersonUsingGETParams) WithRequestUser(requestUser *string) *FindPersonUsingGETParams {
	o.SetRequestUser(requestUser)
	return o
}

// SetRequestUser adds the requestUser to the find person using g e t params
func (o *FindPersonUsingGETParams) SetRequestUser(requestUser *string) {
	o.RequestUser = requestUser
}

// WithSsn adds the ssn to the find person using g e t params
func (o *FindPersonUsingGETParams) WithSsn(ssn *string) *FindPersonUsingGETParams {
	o.SetSsn(ssn)
	return o
}

// SetSsn adds the ssn to the find person using g e t params
func (o *FindPersonUsingGETParams) SetSsn(ssn *string) {
	o.Ssn = ssn
}

// WithVarianceFormat adds the varianceFormat to the find person using g e t params
func (o *FindPersonUsingGETParams) WithVarianceFormat(varianceFormat *string) *FindPersonUsingGETParams {
	o.SetVarianceFormat(varianceFormat)
	return o
}

// SetVarianceFormat adds the varianceFormat to the find person using g e t params
func (o *FindPersonUsingGETParams) SetVarianceFormat(varianceFormat *string) {
	o.VarianceFormat = varianceFormat
}

// WithVarianceLevel adds the varianceLevel to the find person using g e t params
func (o *FindPersonUsingGETParams) WithVarianceLevel(varianceLevel *string) *FindPersonUsingGETParams {
	o.SetVarianceLevel(varianceLevel)
	return o
}

// SetVarianceLevel adds the varianceLevel to the find person using g e t params
func (o *FindPersonUsingGETParams) SetVarianceLevel(varianceLevel *string) {
	o.VarianceLevel = varianceLevel
}

// WriteToRequest writes these params to a swagger request
func (o *FindPersonUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.BirthDate != nil {

		// query param birthDate
		var qrBirthDate string

		if o.BirthDate != nil {
			qrBirthDate = *o.BirthDate
		}
		qBirthDate := qrBirthDate
		if qBirthDate != "" {

			if err := r.SetQueryParam("birthDate", qBirthDate); err != nil {
				return err
			}
		}
	}

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

	if o.FirstName != nil {

		// query param firstName
		var qrFirstName string

		if o.FirstName != nil {
			qrFirstName = *o.FirstName
		}
		qFirstName := qrFirstName
		if qFirstName != "" {

			if err := r.SetQueryParam("firstName", qFirstName); err != nil {
				return err
			}
		}
	}

	if o.LastName != nil {

		// query param lastName
		var qrLastName string

		if o.LastName != nil {
			qrLastName = *o.LastName
		}
		qLastName := qrLastName
		if qLastName != "" {

			if err := r.SetQueryParam("lastName", qLastName); err != nil {
				return err
			}
		}
	}

	if o.MiddleInitial != nil {

		// query param middleInitial
		var qrMiddleInitial string

		if o.MiddleInitial != nil {
			qrMiddleInitial = *o.MiddleInitial
		}
		qMiddleInitial := qrMiddleInitial
		if qMiddleInitial != "" {

			if err := r.SetQueryParam("middleInitial", qMiddleInitial); err != nil {
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

	if o.Ssn != nil {

		// query param ssn
		var qrSsn string

		if o.Ssn != nil {
			qrSsn = *o.Ssn
		}
		qSsn := qrSsn
		if qSsn != "" {

			if err := r.SetQueryParam("ssn", qSsn); err != nil {
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
