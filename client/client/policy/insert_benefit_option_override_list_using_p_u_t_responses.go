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

// InsertBenefitOptionOverrideListUsingPUTReader is a Reader for the InsertBenefitOptionOverrideListUsingPUT structure.
type InsertBenefitOptionOverrideListUsingPUTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InsertBenefitOptionOverrideListUsingPUTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInsertBenefitOptionOverrideListUsingPUTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewInsertBenefitOptionOverrideListUsingPUTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewInsertBenefitOptionOverrideListUsingPUTBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewInsertBenefitOptionOverrideListUsingPUTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewInsertBenefitOptionOverrideListUsingPUTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewInsertBenefitOptionOverrideListUsingPUTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewInsertBenefitOptionOverrideListUsingPUTInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /compassintegrationapp/policy/benefitOptionOverrideList] insertBenefitOptionOverrideListUsingPUT", response, response.Code())
	}
}

// NewInsertBenefitOptionOverrideListUsingPUTOK creates a InsertBenefitOptionOverrideListUsingPUTOK with default headers values
func NewInsertBenefitOptionOverrideListUsingPUTOK() *InsertBenefitOptionOverrideListUsingPUTOK {
	return &InsertBenefitOptionOverrideListUsingPUTOK{}
}

/*
InsertBenefitOptionOverrideListUsingPUTOK describes a response with status code 200, with default header values.

Success
*/
type InsertBenefitOptionOverrideListUsingPUTOK struct {
	Payload *models.ResponseWrapperBoolean
}

// IsSuccess returns true when this insert benefit option override list using p u t o k response has a 2xx status code
func (o *InsertBenefitOptionOverrideListUsingPUTOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this insert benefit option override list using p u t o k response has a 3xx status code
func (o *InsertBenefitOptionOverrideListUsingPUTOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this insert benefit option override list using p u t o k response has a 4xx status code
func (o *InsertBenefitOptionOverrideListUsingPUTOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this insert benefit option override list using p u t o k response has a 5xx status code
func (o *InsertBenefitOptionOverrideListUsingPUTOK) IsServerError() bool {
	return false
}

// IsCode returns true when this insert benefit option override list using p u t o k response a status code equal to that given
func (o *InsertBenefitOptionOverrideListUsingPUTOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the insert benefit option override list using p u t o k response
func (o *InsertBenefitOptionOverrideListUsingPUTOK) Code() int {
	return 200
}

func (o *InsertBenefitOptionOverrideListUsingPUTOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /compassintegrationapp/policy/benefitOptionOverrideList][%d] insertBenefitOptionOverrideListUsingPUTOK %s", 200, payload)
}

func (o *InsertBenefitOptionOverrideListUsingPUTOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /compassintegrationapp/policy/benefitOptionOverrideList][%d] insertBenefitOptionOverrideListUsingPUTOK %s", 200, payload)
}

func (o *InsertBenefitOptionOverrideListUsingPUTOK) GetPayload() *models.ResponseWrapperBoolean {
	return o.Payload
}

func (o *InsertBenefitOptionOverrideListUsingPUTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperBoolean)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInsertBenefitOptionOverrideListUsingPUTCreated creates a InsertBenefitOptionOverrideListUsingPUTCreated with default headers values
func NewInsertBenefitOptionOverrideListUsingPUTCreated() *InsertBenefitOptionOverrideListUsingPUTCreated {
	return &InsertBenefitOptionOverrideListUsingPUTCreated{}
}

/*
InsertBenefitOptionOverrideListUsingPUTCreated describes a response with status code 201, with default header values.

Created
*/
type InsertBenefitOptionOverrideListUsingPUTCreated struct {
}

// IsSuccess returns true when this insert benefit option override list using p u t created response has a 2xx status code
func (o *InsertBenefitOptionOverrideListUsingPUTCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this insert benefit option override list using p u t created response has a 3xx status code
func (o *InsertBenefitOptionOverrideListUsingPUTCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this insert benefit option override list using p u t created response has a 4xx status code
func (o *InsertBenefitOptionOverrideListUsingPUTCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this insert benefit option override list using p u t created response has a 5xx status code
func (o *InsertBenefitOptionOverrideListUsingPUTCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this insert benefit option override list using p u t created response a status code equal to that given
func (o *InsertBenefitOptionOverrideListUsingPUTCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the insert benefit option override list using p u t created response
func (o *InsertBenefitOptionOverrideListUsingPUTCreated) Code() int {
	return 201
}

func (o *InsertBenefitOptionOverrideListUsingPUTCreated) Error() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/policy/benefitOptionOverrideList][%d] insertBenefitOptionOverrideListUsingPUTCreated", 201)
}

func (o *InsertBenefitOptionOverrideListUsingPUTCreated) String() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/policy/benefitOptionOverrideList][%d] insertBenefitOptionOverrideListUsingPUTCreated", 201)
}

func (o *InsertBenefitOptionOverrideListUsingPUTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInsertBenefitOptionOverrideListUsingPUTBadRequest creates a InsertBenefitOptionOverrideListUsingPUTBadRequest with default headers values
func NewInsertBenefitOptionOverrideListUsingPUTBadRequest() *InsertBenefitOptionOverrideListUsingPUTBadRequest {
	return &InsertBenefitOptionOverrideListUsingPUTBadRequest{}
}

/*
InsertBenefitOptionOverrideListUsingPUTBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type InsertBenefitOptionOverrideListUsingPUTBadRequest struct {
}

// IsSuccess returns true when this insert benefit option override list using p u t bad request response has a 2xx status code
func (o *InsertBenefitOptionOverrideListUsingPUTBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this insert benefit option override list using p u t bad request response has a 3xx status code
func (o *InsertBenefitOptionOverrideListUsingPUTBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this insert benefit option override list using p u t bad request response has a 4xx status code
func (o *InsertBenefitOptionOverrideListUsingPUTBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this insert benefit option override list using p u t bad request response has a 5xx status code
func (o *InsertBenefitOptionOverrideListUsingPUTBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this insert benefit option override list using p u t bad request response a status code equal to that given
func (o *InsertBenefitOptionOverrideListUsingPUTBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the insert benefit option override list using p u t bad request response
func (o *InsertBenefitOptionOverrideListUsingPUTBadRequest) Code() int {
	return 400
}

func (o *InsertBenefitOptionOverrideListUsingPUTBadRequest) Error() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/policy/benefitOptionOverrideList][%d] insertBenefitOptionOverrideListUsingPUTBadRequest", 400)
}

func (o *InsertBenefitOptionOverrideListUsingPUTBadRequest) String() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/policy/benefitOptionOverrideList][%d] insertBenefitOptionOverrideListUsingPUTBadRequest", 400)
}

func (o *InsertBenefitOptionOverrideListUsingPUTBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInsertBenefitOptionOverrideListUsingPUTUnauthorized creates a InsertBenefitOptionOverrideListUsingPUTUnauthorized with default headers values
func NewInsertBenefitOptionOverrideListUsingPUTUnauthorized() *InsertBenefitOptionOverrideListUsingPUTUnauthorized {
	return &InsertBenefitOptionOverrideListUsingPUTUnauthorized{}
}

/*
InsertBenefitOptionOverrideListUsingPUTUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type InsertBenefitOptionOverrideListUsingPUTUnauthorized struct {
}

// IsSuccess returns true when this insert benefit option override list using p u t unauthorized response has a 2xx status code
func (o *InsertBenefitOptionOverrideListUsingPUTUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this insert benefit option override list using p u t unauthorized response has a 3xx status code
func (o *InsertBenefitOptionOverrideListUsingPUTUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this insert benefit option override list using p u t unauthorized response has a 4xx status code
func (o *InsertBenefitOptionOverrideListUsingPUTUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this insert benefit option override list using p u t unauthorized response has a 5xx status code
func (o *InsertBenefitOptionOverrideListUsingPUTUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this insert benefit option override list using p u t unauthorized response a status code equal to that given
func (o *InsertBenefitOptionOverrideListUsingPUTUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the insert benefit option override list using p u t unauthorized response
func (o *InsertBenefitOptionOverrideListUsingPUTUnauthorized) Code() int {
	return 401
}

func (o *InsertBenefitOptionOverrideListUsingPUTUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/policy/benefitOptionOverrideList][%d] insertBenefitOptionOverrideListUsingPUTUnauthorized", 401)
}

func (o *InsertBenefitOptionOverrideListUsingPUTUnauthorized) String() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/policy/benefitOptionOverrideList][%d] insertBenefitOptionOverrideListUsingPUTUnauthorized", 401)
}

func (o *InsertBenefitOptionOverrideListUsingPUTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInsertBenefitOptionOverrideListUsingPUTForbidden creates a InsertBenefitOptionOverrideListUsingPUTForbidden with default headers values
func NewInsertBenefitOptionOverrideListUsingPUTForbidden() *InsertBenefitOptionOverrideListUsingPUTForbidden {
	return &InsertBenefitOptionOverrideListUsingPUTForbidden{}
}

/*
InsertBenefitOptionOverrideListUsingPUTForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type InsertBenefitOptionOverrideListUsingPUTForbidden struct {
}

// IsSuccess returns true when this insert benefit option override list using p u t forbidden response has a 2xx status code
func (o *InsertBenefitOptionOverrideListUsingPUTForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this insert benefit option override list using p u t forbidden response has a 3xx status code
func (o *InsertBenefitOptionOverrideListUsingPUTForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this insert benefit option override list using p u t forbidden response has a 4xx status code
func (o *InsertBenefitOptionOverrideListUsingPUTForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this insert benefit option override list using p u t forbidden response has a 5xx status code
func (o *InsertBenefitOptionOverrideListUsingPUTForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this insert benefit option override list using p u t forbidden response a status code equal to that given
func (o *InsertBenefitOptionOverrideListUsingPUTForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the insert benefit option override list using p u t forbidden response
func (o *InsertBenefitOptionOverrideListUsingPUTForbidden) Code() int {
	return 403
}

func (o *InsertBenefitOptionOverrideListUsingPUTForbidden) Error() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/policy/benefitOptionOverrideList][%d] insertBenefitOptionOverrideListUsingPUTForbidden", 403)
}

func (o *InsertBenefitOptionOverrideListUsingPUTForbidden) String() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/policy/benefitOptionOverrideList][%d] insertBenefitOptionOverrideListUsingPUTForbidden", 403)
}

func (o *InsertBenefitOptionOverrideListUsingPUTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInsertBenefitOptionOverrideListUsingPUTNotFound creates a InsertBenefitOptionOverrideListUsingPUTNotFound with default headers values
func NewInsertBenefitOptionOverrideListUsingPUTNotFound() *InsertBenefitOptionOverrideListUsingPUTNotFound {
	return &InsertBenefitOptionOverrideListUsingPUTNotFound{}
}

/*
InsertBenefitOptionOverrideListUsingPUTNotFound describes a response with status code 404, with default header values.

Not Found
*/
type InsertBenefitOptionOverrideListUsingPUTNotFound struct {
}

// IsSuccess returns true when this insert benefit option override list using p u t not found response has a 2xx status code
func (o *InsertBenefitOptionOverrideListUsingPUTNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this insert benefit option override list using p u t not found response has a 3xx status code
func (o *InsertBenefitOptionOverrideListUsingPUTNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this insert benefit option override list using p u t not found response has a 4xx status code
func (o *InsertBenefitOptionOverrideListUsingPUTNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this insert benefit option override list using p u t not found response has a 5xx status code
func (o *InsertBenefitOptionOverrideListUsingPUTNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this insert benefit option override list using p u t not found response a status code equal to that given
func (o *InsertBenefitOptionOverrideListUsingPUTNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the insert benefit option override list using p u t not found response
func (o *InsertBenefitOptionOverrideListUsingPUTNotFound) Code() int {
	return 404
}

func (o *InsertBenefitOptionOverrideListUsingPUTNotFound) Error() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/policy/benefitOptionOverrideList][%d] insertBenefitOptionOverrideListUsingPUTNotFound", 404)
}

func (o *InsertBenefitOptionOverrideListUsingPUTNotFound) String() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/policy/benefitOptionOverrideList][%d] insertBenefitOptionOverrideListUsingPUTNotFound", 404)
}

func (o *InsertBenefitOptionOverrideListUsingPUTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInsertBenefitOptionOverrideListUsingPUTInternalServerError creates a InsertBenefitOptionOverrideListUsingPUTInternalServerError with default headers values
func NewInsertBenefitOptionOverrideListUsingPUTInternalServerError() *InsertBenefitOptionOverrideListUsingPUTInternalServerError {
	return &InsertBenefitOptionOverrideListUsingPUTInternalServerError{}
}

/*
InsertBenefitOptionOverrideListUsingPUTInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type InsertBenefitOptionOverrideListUsingPUTInternalServerError struct {
}

// IsSuccess returns true when this insert benefit option override list using p u t internal server error response has a 2xx status code
func (o *InsertBenefitOptionOverrideListUsingPUTInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this insert benefit option override list using p u t internal server error response has a 3xx status code
func (o *InsertBenefitOptionOverrideListUsingPUTInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this insert benefit option override list using p u t internal server error response has a 4xx status code
func (o *InsertBenefitOptionOverrideListUsingPUTInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this insert benefit option override list using p u t internal server error response has a 5xx status code
func (o *InsertBenefitOptionOverrideListUsingPUTInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this insert benefit option override list using p u t internal server error response a status code equal to that given
func (o *InsertBenefitOptionOverrideListUsingPUTInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the insert benefit option override list using p u t internal server error response
func (o *InsertBenefitOptionOverrideListUsingPUTInternalServerError) Code() int {
	return 500
}

func (o *InsertBenefitOptionOverrideListUsingPUTInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/policy/benefitOptionOverrideList][%d] insertBenefitOptionOverrideListUsingPUTInternalServerError", 500)
}

func (o *InsertBenefitOptionOverrideListUsingPUTInternalServerError) String() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/policy/benefitOptionOverrideList][%d] insertBenefitOptionOverrideListUsingPUTInternalServerError", 500)
}

func (o *InsertBenefitOptionOverrideListUsingPUTInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
