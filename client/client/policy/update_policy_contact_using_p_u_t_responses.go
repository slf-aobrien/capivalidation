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

// UpdatePolicyContactUsingPUTReader is a Reader for the UpdatePolicyContactUsingPUT structure.
type UpdatePolicyContactUsingPUTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdatePolicyContactUsingPUTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdatePolicyContactUsingPUTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewUpdatePolicyContactUsingPUTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdatePolicyContactUsingPUTBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdatePolicyContactUsingPUTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdatePolicyContactUsingPUTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdatePolicyContactUsingPUTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdatePolicyContactUsingPUTInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /compassintegrationapp/policy/policyContact/update] updatePolicyContactUsingPUT", response, response.Code())
	}
}

// NewUpdatePolicyContactUsingPUTOK creates a UpdatePolicyContactUsingPUTOK with default headers values
func NewUpdatePolicyContactUsingPUTOK() *UpdatePolicyContactUsingPUTOK {
	return &UpdatePolicyContactUsingPUTOK{}
}

/*
UpdatePolicyContactUsingPUTOK describes a response with status code 200, with default header values.

Success
*/
type UpdatePolicyContactUsingPUTOK struct {
	Payload *models.ResponseWrapperBoolean
}

// IsSuccess returns true when this update policy contact using p u t o k response has a 2xx status code
func (o *UpdatePolicyContactUsingPUTOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update policy contact using p u t o k response has a 3xx status code
func (o *UpdatePolicyContactUsingPUTOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update policy contact using p u t o k response has a 4xx status code
func (o *UpdatePolicyContactUsingPUTOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update policy contact using p u t o k response has a 5xx status code
func (o *UpdatePolicyContactUsingPUTOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update policy contact using p u t o k response a status code equal to that given
func (o *UpdatePolicyContactUsingPUTOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update policy contact using p u t o k response
func (o *UpdatePolicyContactUsingPUTOK) Code() int {
	return 200
}

func (o *UpdatePolicyContactUsingPUTOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /compassintegrationapp/policy/policyContact/update][%d] updatePolicyContactUsingPUTOK %s", 200, payload)
}

func (o *UpdatePolicyContactUsingPUTOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /compassintegrationapp/policy/policyContact/update][%d] updatePolicyContactUsingPUTOK %s", 200, payload)
}

func (o *UpdatePolicyContactUsingPUTOK) GetPayload() *models.ResponseWrapperBoolean {
	return o.Payload
}

func (o *UpdatePolicyContactUsingPUTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperBoolean)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdatePolicyContactUsingPUTCreated creates a UpdatePolicyContactUsingPUTCreated with default headers values
func NewUpdatePolicyContactUsingPUTCreated() *UpdatePolicyContactUsingPUTCreated {
	return &UpdatePolicyContactUsingPUTCreated{}
}

/*
UpdatePolicyContactUsingPUTCreated describes a response with status code 201, with default header values.

Created
*/
type UpdatePolicyContactUsingPUTCreated struct {
}

// IsSuccess returns true when this update policy contact using p u t created response has a 2xx status code
func (o *UpdatePolicyContactUsingPUTCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update policy contact using p u t created response has a 3xx status code
func (o *UpdatePolicyContactUsingPUTCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update policy contact using p u t created response has a 4xx status code
func (o *UpdatePolicyContactUsingPUTCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this update policy contact using p u t created response has a 5xx status code
func (o *UpdatePolicyContactUsingPUTCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this update policy contact using p u t created response a status code equal to that given
func (o *UpdatePolicyContactUsingPUTCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the update policy contact using p u t created response
func (o *UpdatePolicyContactUsingPUTCreated) Code() int {
	return 201
}

func (o *UpdatePolicyContactUsingPUTCreated) Error() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/policy/policyContact/update][%d] updatePolicyContactUsingPUTCreated", 201)
}

func (o *UpdatePolicyContactUsingPUTCreated) String() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/policy/policyContact/update][%d] updatePolicyContactUsingPUTCreated", 201)
}

func (o *UpdatePolicyContactUsingPUTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdatePolicyContactUsingPUTBadRequest creates a UpdatePolicyContactUsingPUTBadRequest with default headers values
func NewUpdatePolicyContactUsingPUTBadRequest() *UpdatePolicyContactUsingPUTBadRequest {
	return &UpdatePolicyContactUsingPUTBadRequest{}
}

/*
UpdatePolicyContactUsingPUTBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UpdatePolicyContactUsingPUTBadRequest struct {
}

// IsSuccess returns true when this update policy contact using p u t bad request response has a 2xx status code
func (o *UpdatePolicyContactUsingPUTBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update policy contact using p u t bad request response has a 3xx status code
func (o *UpdatePolicyContactUsingPUTBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update policy contact using p u t bad request response has a 4xx status code
func (o *UpdatePolicyContactUsingPUTBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update policy contact using p u t bad request response has a 5xx status code
func (o *UpdatePolicyContactUsingPUTBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update policy contact using p u t bad request response a status code equal to that given
func (o *UpdatePolicyContactUsingPUTBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update policy contact using p u t bad request response
func (o *UpdatePolicyContactUsingPUTBadRequest) Code() int {
	return 400
}

func (o *UpdatePolicyContactUsingPUTBadRequest) Error() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/policy/policyContact/update][%d] updatePolicyContactUsingPUTBadRequest", 400)
}

func (o *UpdatePolicyContactUsingPUTBadRequest) String() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/policy/policyContact/update][%d] updatePolicyContactUsingPUTBadRequest", 400)
}

func (o *UpdatePolicyContactUsingPUTBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdatePolicyContactUsingPUTUnauthorized creates a UpdatePolicyContactUsingPUTUnauthorized with default headers values
func NewUpdatePolicyContactUsingPUTUnauthorized() *UpdatePolicyContactUsingPUTUnauthorized {
	return &UpdatePolicyContactUsingPUTUnauthorized{}
}

/*
UpdatePolicyContactUsingPUTUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UpdatePolicyContactUsingPUTUnauthorized struct {
}

// IsSuccess returns true when this update policy contact using p u t unauthorized response has a 2xx status code
func (o *UpdatePolicyContactUsingPUTUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update policy contact using p u t unauthorized response has a 3xx status code
func (o *UpdatePolicyContactUsingPUTUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update policy contact using p u t unauthorized response has a 4xx status code
func (o *UpdatePolicyContactUsingPUTUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update policy contact using p u t unauthorized response has a 5xx status code
func (o *UpdatePolicyContactUsingPUTUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update policy contact using p u t unauthorized response a status code equal to that given
func (o *UpdatePolicyContactUsingPUTUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the update policy contact using p u t unauthorized response
func (o *UpdatePolicyContactUsingPUTUnauthorized) Code() int {
	return 401
}

func (o *UpdatePolicyContactUsingPUTUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/policy/policyContact/update][%d] updatePolicyContactUsingPUTUnauthorized", 401)
}

func (o *UpdatePolicyContactUsingPUTUnauthorized) String() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/policy/policyContact/update][%d] updatePolicyContactUsingPUTUnauthorized", 401)
}

func (o *UpdatePolicyContactUsingPUTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdatePolicyContactUsingPUTForbidden creates a UpdatePolicyContactUsingPUTForbidden with default headers values
func NewUpdatePolicyContactUsingPUTForbidden() *UpdatePolicyContactUsingPUTForbidden {
	return &UpdatePolicyContactUsingPUTForbidden{}
}

/*
UpdatePolicyContactUsingPUTForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdatePolicyContactUsingPUTForbidden struct {
}

// IsSuccess returns true when this update policy contact using p u t forbidden response has a 2xx status code
func (o *UpdatePolicyContactUsingPUTForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update policy contact using p u t forbidden response has a 3xx status code
func (o *UpdatePolicyContactUsingPUTForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update policy contact using p u t forbidden response has a 4xx status code
func (o *UpdatePolicyContactUsingPUTForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update policy contact using p u t forbidden response has a 5xx status code
func (o *UpdatePolicyContactUsingPUTForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update policy contact using p u t forbidden response a status code equal to that given
func (o *UpdatePolicyContactUsingPUTForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update policy contact using p u t forbidden response
func (o *UpdatePolicyContactUsingPUTForbidden) Code() int {
	return 403
}

func (o *UpdatePolicyContactUsingPUTForbidden) Error() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/policy/policyContact/update][%d] updatePolicyContactUsingPUTForbidden", 403)
}

func (o *UpdatePolicyContactUsingPUTForbidden) String() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/policy/policyContact/update][%d] updatePolicyContactUsingPUTForbidden", 403)
}

func (o *UpdatePolicyContactUsingPUTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdatePolicyContactUsingPUTNotFound creates a UpdatePolicyContactUsingPUTNotFound with default headers values
func NewUpdatePolicyContactUsingPUTNotFound() *UpdatePolicyContactUsingPUTNotFound {
	return &UpdatePolicyContactUsingPUTNotFound{}
}

/*
UpdatePolicyContactUsingPUTNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdatePolicyContactUsingPUTNotFound struct {
}

// IsSuccess returns true when this update policy contact using p u t not found response has a 2xx status code
func (o *UpdatePolicyContactUsingPUTNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update policy contact using p u t not found response has a 3xx status code
func (o *UpdatePolicyContactUsingPUTNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update policy contact using p u t not found response has a 4xx status code
func (o *UpdatePolicyContactUsingPUTNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update policy contact using p u t not found response has a 5xx status code
func (o *UpdatePolicyContactUsingPUTNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update policy contact using p u t not found response a status code equal to that given
func (o *UpdatePolicyContactUsingPUTNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update policy contact using p u t not found response
func (o *UpdatePolicyContactUsingPUTNotFound) Code() int {
	return 404
}

func (o *UpdatePolicyContactUsingPUTNotFound) Error() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/policy/policyContact/update][%d] updatePolicyContactUsingPUTNotFound", 404)
}

func (o *UpdatePolicyContactUsingPUTNotFound) String() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/policy/policyContact/update][%d] updatePolicyContactUsingPUTNotFound", 404)
}

func (o *UpdatePolicyContactUsingPUTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdatePolicyContactUsingPUTInternalServerError creates a UpdatePolicyContactUsingPUTInternalServerError with default headers values
func NewUpdatePolicyContactUsingPUTInternalServerError() *UpdatePolicyContactUsingPUTInternalServerError {
	return &UpdatePolicyContactUsingPUTInternalServerError{}
}

/*
UpdatePolicyContactUsingPUTInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type UpdatePolicyContactUsingPUTInternalServerError struct {
}

// IsSuccess returns true when this update policy contact using p u t internal server error response has a 2xx status code
func (o *UpdatePolicyContactUsingPUTInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update policy contact using p u t internal server error response has a 3xx status code
func (o *UpdatePolicyContactUsingPUTInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update policy contact using p u t internal server error response has a 4xx status code
func (o *UpdatePolicyContactUsingPUTInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update policy contact using p u t internal server error response has a 5xx status code
func (o *UpdatePolicyContactUsingPUTInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update policy contact using p u t internal server error response a status code equal to that given
func (o *UpdatePolicyContactUsingPUTInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update policy contact using p u t internal server error response
func (o *UpdatePolicyContactUsingPUTInternalServerError) Code() int {
	return 500
}

func (o *UpdatePolicyContactUsingPUTInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/policy/policyContact/update][%d] updatePolicyContactUsingPUTInternalServerError", 500)
}

func (o *UpdatePolicyContactUsingPUTInternalServerError) String() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/policy/policyContact/update][%d] updatePolicyContactUsingPUTInternalServerError", 500)
}

func (o *UpdatePolicyContactUsingPUTInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
