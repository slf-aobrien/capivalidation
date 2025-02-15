// Code generated by go-swagger; DO NOT EDIT.

package policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"client/models"
)

// AddPolicyBenefitStopLossUsingPOSTReader is a Reader for the AddPolicyBenefitStopLossUsingPOST structure.
type AddPolicyBenefitStopLossUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddPolicyBenefitStopLossUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddPolicyBenefitStopLossUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewAddPolicyBenefitStopLossUsingPOSTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddPolicyBenefitStopLossUsingPOSTBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAddPolicyBenefitStopLossUsingPOSTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAddPolicyBenefitStopLossUsingPOSTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAddPolicyBenefitStopLossUsingPOSTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAddPolicyBenefitStopLossUsingPOSTInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /compassintegrationapp/policy/policyBenefit_StopLoss] addPolicyBenefitStopLossUsingPOST", response, response.Code())
	}
}

// NewAddPolicyBenefitStopLossUsingPOSTOK creates a AddPolicyBenefitStopLossUsingPOSTOK with default headers values
func NewAddPolicyBenefitStopLossUsingPOSTOK() *AddPolicyBenefitStopLossUsingPOSTOK {
	return &AddPolicyBenefitStopLossUsingPOSTOK{}
}

/*
AddPolicyBenefitStopLossUsingPOSTOK describes a response with status code 200, with default header values.

Success
*/
type AddPolicyBenefitStopLossUsingPOSTOK struct {
	Payload *models.ResponseWrapperBoolean
}

// IsSuccess returns true when this add policy benefit stop loss using p o s t o k response has a 2xx status code
func (o *AddPolicyBenefitStopLossUsingPOSTOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this add policy benefit stop loss using p o s t o k response has a 3xx status code
func (o *AddPolicyBenefitStopLossUsingPOSTOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add policy benefit stop loss using p o s t o k response has a 4xx status code
func (o *AddPolicyBenefitStopLossUsingPOSTOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this add policy benefit stop loss using p o s t o k response has a 5xx status code
func (o *AddPolicyBenefitStopLossUsingPOSTOK) IsServerError() bool {
	return false
}

// IsCode returns true when this add policy benefit stop loss using p o s t o k response a status code equal to that given
func (o *AddPolicyBenefitStopLossUsingPOSTOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the add policy benefit stop loss using p o s t o k response
func (o *AddPolicyBenefitStopLossUsingPOSTOK) Code() int {
	return 200
}

func (o *AddPolicyBenefitStopLossUsingPOSTOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /compassintegrationapp/policy/policyBenefit_StopLoss][%d] addPolicyBenefitStopLossUsingPOSTOK %s", 200, payload)
}

func (o *AddPolicyBenefitStopLossUsingPOSTOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /compassintegrationapp/policy/policyBenefit_StopLoss][%d] addPolicyBenefitStopLossUsingPOSTOK %s", 200, payload)
}

func (o *AddPolicyBenefitStopLossUsingPOSTOK) GetPayload() *models.ResponseWrapperBoolean {
	return o.Payload
}

func (o *AddPolicyBenefitStopLossUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperBoolean)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddPolicyBenefitStopLossUsingPOSTCreated creates a AddPolicyBenefitStopLossUsingPOSTCreated with default headers values
func NewAddPolicyBenefitStopLossUsingPOSTCreated() *AddPolicyBenefitStopLossUsingPOSTCreated {
	return &AddPolicyBenefitStopLossUsingPOSTCreated{}
}

/*
AddPolicyBenefitStopLossUsingPOSTCreated describes a response with status code 201, with default header values.

Created
*/
type AddPolicyBenefitStopLossUsingPOSTCreated struct {
}

// IsSuccess returns true when this add policy benefit stop loss using p o s t created response has a 2xx status code
func (o *AddPolicyBenefitStopLossUsingPOSTCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this add policy benefit stop loss using p o s t created response has a 3xx status code
func (o *AddPolicyBenefitStopLossUsingPOSTCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add policy benefit stop loss using p o s t created response has a 4xx status code
func (o *AddPolicyBenefitStopLossUsingPOSTCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this add policy benefit stop loss using p o s t created response has a 5xx status code
func (o *AddPolicyBenefitStopLossUsingPOSTCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this add policy benefit stop loss using p o s t created response a status code equal to that given
func (o *AddPolicyBenefitStopLossUsingPOSTCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the add policy benefit stop loss using p o s t created response
func (o *AddPolicyBenefitStopLossUsingPOSTCreated) Code() int {
	return 201
}

func (o *AddPolicyBenefitStopLossUsingPOSTCreated) Error() string {
	return fmt.Sprintf("[POST /compassintegrationapp/policy/policyBenefit_StopLoss][%d] addPolicyBenefitStopLossUsingPOSTCreated", 201)
}

func (o *AddPolicyBenefitStopLossUsingPOSTCreated) String() string {
	return fmt.Sprintf("[POST /compassintegrationapp/policy/policyBenefit_StopLoss][%d] addPolicyBenefitStopLossUsingPOSTCreated", 201)
}

func (o *AddPolicyBenefitStopLossUsingPOSTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddPolicyBenefitStopLossUsingPOSTBadRequest creates a AddPolicyBenefitStopLossUsingPOSTBadRequest with default headers values
func NewAddPolicyBenefitStopLossUsingPOSTBadRequest() *AddPolicyBenefitStopLossUsingPOSTBadRequest {
	return &AddPolicyBenefitStopLossUsingPOSTBadRequest{}
}

/*
AddPolicyBenefitStopLossUsingPOSTBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AddPolicyBenefitStopLossUsingPOSTBadRequest struct {
}

// IsSuccess returns true when this add policy benefit stop loss using p o s t bad request response has a 2xx status code
func (o *AddPolicyBenefitStopLossUsingPOSTBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add policy benefit stop loss using p o s t bad request response has a 3xx status code
func (o *AddPolicyBenefitStopLossUsingPOSTBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add policy benefit stop loss using p o s t bad request response has a 4xx status code
func (o *AddPolicyBenefitStopLossUsingPOSTBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this add policy benefit stop loss using p o s t bad request response has a 5xx status code
func (o *AddPolicyBenefitStopLossUsingPOSTBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this add policy benefit stop loss using p o s t bad request response a status code equal to that given
func (o *AddPolicyBenefitStopLossUsingPOSTBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the add policy benefit stop loss using p o s t bad request response
func (o *AddPolicyBenefitStopLossUsingPOSTBadRequest) Code() int {
	return 400
}

func (o *AddPolicyBenefitStopLossUsingPOSTBadRequest) Error() string {
	return fmt.Sprintf("[POST /compassintegrationapp/policy/policyBenefit_StopLoss][%d] addPolicyBenefitStopLossUsingPOSTBadRequest", 400)
}

func (o *AddPolicyBenefitStopLossUsingPOSTBadRequest) String() string {
	return fmt.Sprintf("[POST /compassintegrationapp/policy/policyBenefit_StopLoss][%d] addPolicyBenefitStopLossUsingPOSTBadRequest", 400)
}

func (o *AddPolicyBenefitStopLossUsingPOSTBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddPolicyBenefitStopLossUsingPOSTUnauthorized creates a AddPolicyBenefitStopLossUsingPOSTUnauthorized with default headers values
func NewAddPolicyBenefitStopLossUsingPOSTUnauthorized() *AddPolicyBenefitStopLossUsingPOSTUnauthorized {
	return &AddPolicyBenefitStopLossUsingPOSTUnauthorized{}
}

/*
AddPolicyBenefitStopLossUsingPOSTUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type AddPolicyBenefitStopLossUsingPOSTUnauthorized struct {
}

// IsSuccess returns true when this add policy benefit stop loss using p o s t unauthorized response has a 2xx status code
func (o *AddPolicyBenefitStopLossUsingPOSTUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add policy benefit stop loss using p o s t unauthorized response has a 3xx status code
func (o *AddPolicyBenefitStopLossUsingPOSTUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add policy benefit stop loss using p o s t unauthorized response has a 4xx status code
func (o *AddPolicyBenefitStopLossUsingPOSTUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this add policy benefit stop loss using p o s t unauthorized response has a 5xx status code
func (o *AddPolicyBenefitStopLossUsingPOSTUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this add policy benefit stop loss using p o s t unauthorized response a status code equal to that given
func (o *AddPolicyBenefitStopLossUsingPOSTUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the add policy benefit stop loss using p o s t unauthorized response
func (o *AddPolicyBenefitStopLossUsingPOSTUnauthorized) Code() int {
	return 401
}

func (o *AddPolicyBenefitStopLossUsingPOSTUnauthorized) Error() string {
	return fmt.Sprintf("[POST /compassintegrationapp/policy/policyBenefit_StopLoss][%d] addPolicyBenefitStopLossUsingPOSTUnauthorized", 401)
}

func (o *AddPolicyBenefitStopLossUsingPOSTUnauthorized) String() string {
	return fmt.Sprintf("[POST /compassintegrationapp/policy/policyBenefit_StopLoss][%d] addPolicyBenefitStopLossUsingPOSTUnauthorized", 401)
}

func (o *AddPolicyBenefitStopLossUsingPOSTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddPolicyBenefitStopLossUsingPOSTForbidden creates a AddPolicyBenefitStopLossUsingPOSTForbidden with default headers values
func NewAddPolicyBenefitStopLossUsingPOSTForbidden() *AddPolicyBenefitStopLossUsingPOSTForbidden {
	return &AddPolicyBenefitStopLossUsingPOSTForbidden{}
}

/*
AddPolicyBenefitStopLossUsingPOSTForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AddPolicyBenefitStopLossUsingPOSTForbidden struct {
}

// IsSuccess returns true when this add policy benefit stop loss using p o s t forbidden response has a 2xx status code
func (o *AddPolicyBenefitStopLossUsingPOSTForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add policy benefit stop loss using p o s t forbidden response has a 3xx status code
func (o *AddPolicyBenefitStopLossUsingPOSTForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add policy benefit stop loss using p o s t forbidden response has a 4xx status code
func (o *AddPolicyBenefitStopLossUsingPOSTForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this add policy benefit stop loss using p o s t forbidden response has a 5xx status code
func (o *AddPolicyBenefitStopLossUsingPOSTForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this add policy benefit stop loss using p o s t forbidden response a status code equal to that given
func (o *AddPolicyBenefitStopLossUsingPOSTForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the add policy benefit stop loss using p o s t forbidden response
func (o *AddPolicyBenefitStopLossUsingPOSTForbidden) Code() int {
	return 403
}

func (o *AddPolicyBenefitStopLossUsingPOSTForbidden) Error() string {
	return fmt.Sprintf("[POST /compassintegrationapp/policy/policyBenefit_StopLoss][%d] addPolicyBenefitStopLossUsingPOSTForbidden", 403)
}

func (o *AddPolicyBenefitStopLossUsingPOSTForbidden) String() string {
	return fmt.Sprintf("[POST /compassintegrationapp/policy/policyBenefit_StopLoss][%d] addPolicyBenefitStopLossUsingPOSTForbidden", 403)
}

func (o *AddPolicyBenefitStopLossUsingPOSTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddPolicyBenefitStopLossUsingPOSTNotFound creates a AddPolicyBenefitStopLossUsingPOSTNotFound with default headers values
func NewAddPolicyBenefitStopLossUsingPOSTNotFound() *AddPolicyBenefitStopLossUsingPOSTNotFound {
	return &AddPolicyBenefitStopLossUsingPOSTNotFound{}
}

/*
AddPolicyBenefitStopLossUsingPOSTNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AddPolicyBenefitStopLossUsingPOSTNotFound struct {
}

// IsSuccess returns true when this add policy benefit stop loss using p o s t not found response has a 2xx status code
func (o *AddPolicyBenefitStopLossUsingPOSTNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add policy benefit stop loss using p o s t not found response has a 3xx status code
func (o *AddPolicyBenefitStopLossUsingPOSTNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add policy benefit stop loss using p o s t not found response has a 4xx status code
func (o *AddPolicyBenefitStopLossUsingPOSTNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this add policy benefit stop loss using p o s t not found response has a 5xx status code
func (o *AddPolicyBenefitStopLossUsingPOSTNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this add policy benefit stop loss using p o s t not found response a status code equal to that given
func (o *AddPolicyBenefitStopLossUsingPOSTNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the add policy benefit stop loss using p o s t not found response
func (o *AddPolicyBenefitStopLossUsingPOSTNotFound) Code() int {
	return 404
}

func (o *AddPolicyBenefitStopLossUsingPOSTNotFound) Error() string {
	return fmt.Sprintf("[POST /compassintegrationapp/policy/policyBenefit_StopLoss][%d] addPolicyBenefitStopLossUsingPOSTNotFound", 404)
}

func (o *AddPolicyBenefitStopLossUsingPOSTNotFound) String() string {
	return fmt.Sprintf("[POST /compassintegrationapp/policy/policyBenefit_StopLoss][%d] addPolicyBenefitStopLossUsingPOSTNotFound", 404)
}

func (o *AddPolicyBenefitStopLossUsingPOSTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddPolicyBenefitStopLossUsingPOSTInternalServerError creates a AddPolicyBenefitStopLossUsingPOSTInternalServerError with default headers values
func NewAddPolicyBenefitStopLossUsingPOSTInternalServerError() *AddPolicyBenefitStopLossUsingPOSTInternalServerError {
	return &AddPolicyBenefitStopLossUsingPOSTInternalServerError{}
}

/*
AddPolicyBenefitStopLossUsingPOSTInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type AddPolicyBenefitStopLossUsingPOSTInternalServerError struct {
}

// IsSuccess returns true when this add policy benefit stop loss using p o s t internal server error response has a 2xx status code
func (o *AddPolicyBenefitStopLossUsingPOSTInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add policy benefit stop loss using p o s t internal server error response has a 3xx status code
func (o *AddPolicyBenefitStopLossUsingPOSTInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add policy benefit stop loss using p o s t internal server error response has a 4xx status code
func (o *AddPolicyBenefitStopLossUsingPOSTInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this add policy benefit stop loss using p o s t internal server error response has a 5xx status code
func (o *AddPolicyBenefitStopLossUsingPOSTInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this add policy benefit stop loss using p o s t internal server error response a status code equal to that given
func (o *AddPolicyBenefitStopLossUsingPOSTInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the add policy benefit stop loss using p o s t internal server error response
func (o *AddPolicyBenefitStopLossUsingPOSTInternalServerError) Code() int {
	return 500
}

func (o *AddPolicyBenefitStopLossUsingPOSTInternalServerError) Error() string {
	return fmt.Sprintf("[POST /compassintegrationapp/policy/policyBenefit_StopLoss][%d] addPolicyBenefitStopLossUsingPOSTInternalServerError", 500)
}

func (o *AddPolicyBenefitStopLossUsingPOSTInternalServerError) String() string {
	return fmt.Sprintf("[POST /compassintegrationapp/policy/policyBenefit_StopLoss][%d] addPolicyBenefitStopLossUsingPOSTInternalServerError", 500)
}

func (o *AddPolicyBenefitStopLossUsingPOSTInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
