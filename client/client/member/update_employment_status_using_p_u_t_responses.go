// Code generated by go-swagger; DO NOT EDIT.

package member

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

// UpdateEmploymentStatusUsingPUTReader is a Reader for the UpdateEmploymentStatusUsingPUT structure.
type UpdateEmploymentStatusUsingPUTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateEmploymentStatusUsingPUTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateEmploymentStatusUsingPUTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewUpdateEmploymentStatusUsingPUTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateEmploymentStatusUsingPUTBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateEmploymentStatusUsingPUTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateEmploymentStatusUsingPUTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateEmploymentStatusUsingPUTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateEmploymentStatusUsingPUTInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /compassintegrationapp/member/update/employmentStatus] updateEmploymentStatusUsingPUT", response, response.Code())
	}
}

// NewUpdateEmploymentStatusUsingPUTOK creates a UpdateEmploymentStatusUsingPUTOK with default headers values
func NewUpdateEmploymentStatusUsingPUTOK() *UpdateEmploymentStatusUsingPUTOK {
	return &UpdateEmploymentStatusUsingPUTOK{}
}

/*
UpdateEmploymentStatusUsingPUTOK describes a response with status code 200, with default header values.

Success
*/
type UpdateEmploymentStatusUsingPUTOK struct {
	Payload *models.ResponseWrapperBoolean
}

// IsSuccess returns true when this update employment status using p u t o k response has a 2xx status code
func (o *UpdateEmploymentStatusUsingPUTOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update employment status using p u t o k response has a 3xx status code
func (o *UpdateEmploymentStatusUsingPUTOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update employment status using p u t o k response has a 4xx status code
func (o *UpdateEmploymentStatusUsingPUTOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update employment status using p u t o k response has a 5xx status code
func (o *UpdateEmploymentStatusUsingPUTOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update employment status using p u t o k response a status code equal to that given
func (o *UpdateEmploymentStatusUsingPUTOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update employment status using p u t o k response
func (o *UpdateEmploymentStatusUsingPUTOK) Code() int {
	return 200
}

func (o *UpdateEmploymentStatusUsingPUTOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /compassintegrationapp/member/update/employmentStatus][%d] updateEmploymentStatusUsingPUTOK %s", 200, payload)
}

func (o *UpdateEmploymentStatusUsingPUTOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /compassintegrationapp/member/update/employmentStatus][%d] updateEmploymentStatusUsingPUTOK %s", 200, payload)
}

func (o *UpdateEmploymentStatusUsingPUTOK) GetPayload() *models.ResponseWrapperBoolean {
	return o.Payload
}

func (o *UpdateEmploymentStatusUsingPUTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperBoolean)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateEmploymentStatusUsingPUTCreated creates a UpdateEmploymentStatusUsingPUTCreated with default headers values
func NewUpdateEmploymentStatusUsingPUTCreated() *UpdateEmploymentStatusUsingPUTCreated {
	return &UpdateEmploymentStatusUsingPUTCreated{}
}

/*
UpdateEmploymentStatusUsingPUTCreated describes a response with status code 201, with default header values.

Created
*/
type UpdateEmploymentStatusUsingPUTCreated struct {
}

// IsSuccess returns true when this update employment status using p u t created response has a 2xx status code
func (o *UpdateEmploymentStatusUsingPUTCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update employment status using p u t created response has a 3xx status code
func (o *UpdateEmploymentStatusUsingPUTCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update employment status using p u t created response has a 4xx status code
func (o *UpdateEmploymentStatusUsingPUTCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this update employment status using p u t created response has a 5xx status code
func (o *UpdateEmploymentStatusUsingPUTCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this update employment status using p u t created response a status code equal to that given
func (o *UpdateEmploymentStatusUsingPUTCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the update employment status using p u t created response
func (o *UpdateEmploymentStatusUsingPUTCreated) Code() int {
	return 201
}

func (o *UpdateEmploymentStatusUsingPUTCreated) Error() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/member/update/employmentStatus][%d] updateEmploymentStatusUsingPUTCreated", 201)
}

func (o *UpdateEmploymentStatusUsingPUTCreated) String() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/member/update/employmentStatus][%d] updateEmploymentStatusUsingPUTCreated", 201)
}

func (o *UpdateEmploymentStatusUsingPUTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateEmploymentStatusUsingPUTBadRequest creates a UpdateEmploymentStatusUsingPUTBadRequest with default headers values
func NewUpdateEmploymentStatusUsingPUTBadRequest() *UpdateEmploymentStatusUsingPUTBadRequest {
	return &UpdateEmploymentStatusUsingPUTBadRequest{}
}

/*
UpdateEmploymentStatusUsingPUTBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UpdateEmploymentStatusUsingPUTBadRequest struct {
}

// IsSuccess returns true when this update employment status using p u t bad request response has a 2xx status code
func (o *UpdateEmploymentStatusUsingPUTBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update employment status using p u t bad request response has a 3xx status code
func (o *UpdateEmploymentStatusUsingPUTBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update employment status using p u t bad request response has a 4xx status code
func (o *UpdateEmploymentStatusUsingPUTBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update employment status using p u t bad request response has a 5xx status code
func (o *UpdateEmploymentStatusUsingPUTBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update employment status using p u t bad request response a status code equal to that given
func (o *UpdateEmploymentStatusUsingPUTBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update employment status using p u t bad request response
func (o *UpdateEmploymentStatusUsingPUTBadRequest) Code() int {
	return 400
}

func (o *UpdateEmploymentStatusUsingPUTBadRequest) Error() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/member/update/employmentStatus][%d] updateEmploymentStatusUsingPUTBadRequest", 400)
}

func (o *UpdateEmploymentStatusUsingPUTBadRequest) String() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/member/update/employmentStatus][%d] updateEmploymentStatusUsingPUTBadRequest", 400)
}

func (o *UpdateEmploymentStatusUsingPUTBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateEmploymentStatusUsingPUTUnauthorized creates a UpdateEmploymentStatusUsingPUTUnauthorized with default headers values
func NewUpdateEmploymentStatusUsingPUTUnauthorized() *UpdateEmploymentStatusUsingPUTUnauthorized {
	return &UpdateEmploymentStatusUsingPUTUnauthorized{}
}

/*
UpdateEmploymentStatusUsingPUTUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UpdateEmploymentStatusUsingPUTUnauthorized struct {
}

// IsSuccess returns true when this update employment status using p u t unauthorized response has a 2xx status code
func (o *UpdateEmploymentStatusUsingPUTUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update employment status using p u t unauthorized response has a 3xx status code
func (o *UpdateEmploymentStatusUsingPUTUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update employment status using p u t unauthorized response has a 4xx status code
func (o *UpdateEmploymentStatusUsingPUTUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update employment status using p u t unauthorized response has a 5xx status code
func (o *UpdateEmploymentStatusUsingPUTUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update employment status using p u t unauthorized response a status code equal to that given
func (o *UpdateEmploymentStatusUsingPUTUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the update employment status using p u t unauthorized response
func (o *UpdateEmploymentStatusUsingPUTUnauthorized) Code() int {
	return 401
}

func (o *UpdateEmploymentStatusUsingPUTUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/member/update/employmentStatus][%d] updateEmploymentStatusUsingPUTUnauthorized", 401)
}

func (o *UpdateEmploymentStatusUsingPUTUnauthorized) String() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/member/update/employmentStatus][%d] updateEmploymentStatusUsingPUTUnauthorized", 401)
}

func (o *UpdateEmploymentStatusUsingPUTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateEmploymentStatusUsingPUTForbidden creates a UpdateEmploymentStatusUsingPUTForbidden with default headers values
func NewUpdateEmploymentStatusUsingPUTForbidden() *UpdateEmploymentStatusUsingPUTForbidden {
	return &UpdateEmploymentStatusUsingPUTForbidden{}
}

/*
UpdateEmploymentStatusUsingPUTForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateEmploymentStatusUsingPUTForbidden struct {
}

// IsSuccess returns true when this update employment status using p u t forbidden response has a 2xx status code
func (o *UpdateEmploymentStatusUsingPUTForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update employment status using p u t forbidden response has a 3xx status code
func (o *UpdateEmploymentStatusUsingPUTForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update employment status using p u t forbidden response has a 4xx status code
func (o *UpdateEmploymentStatusUsingPUTForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update employment status using p u t forbidden response has a 5xx status code
func (o *UpdateEmploymentStatusUsingPUTForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update employment status using p u t forbidden response a status code equal to that given
func (o *UpdateEmploymentStatusUsingPUTForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update employment status using p u t forbidden response
func (o *UpdateEmploymentStatusUsingPUTForbidden) Code() int {
	return 403
}

func (o *UpdateEmploymentStatusUsingPUTForbidden) Error() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/member/update/employmentStatus][%d] updateEmploymentStatusUsingPUTForbidden", 403)
}

func (o *UpdateEmploymentStatusUsingPUTForbidden) String() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/member/update/employmentStatus][%d] updateEmploymentStatusUsingPUTForbidden", 403)
}

func (o *UpdateEmploymentStatusUsingPUTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateEmploymentStatusUsingPUTNotFound creates a UpdateEmploymentStatusUsingPUTNotFound with default headers values
func NewUpdateEmploymentStatusUsingPUTNotFound() *UpdateEmploymentStatusUsingPUTNotFound {
	return &UpdateEmploymentStatusUsingPUTNotFound{}
}

/*
UpdateEmploymentStatusUsingPUTNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdateEmploymentStatusUsingPUTNotFound struct {
}

// IsSuccess returns true when this update employment status using p u t not found response has a 2xx status code
func (o *UpdateEmploymentStatusUsingPUTNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update employment status using p u t not found response has a 3xx status code
func (o *UpdateEmploymentStatusUsingPUTNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update employment status using p u t not found response has a 4xx status code
func (o *UpdateEmploymentStatusUsingPUTNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update employment status using p u t not found response has a 5xx status code
func (o *UpdateEmploymentStatusUsingPUTNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update employment status using p u t not found response a status code equal to that given
func (o *UpdateEmploymentStatusUsingPUTNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update employment status using p u t not found response
func (o *UpdateEmploymentStatusUsingPUTNotFound) Code() int {
	return 404
}

func (o *UpdateEmploymentStatusUsingPUTNotFound) Error() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/member/update/employmentStatus][%d] updateEmploymentStatusUsingPUTNotFound", 404)
}

func (o *UpdateEmploymentStatusUsingPUTNotFound) String() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/member/update/employmentStatus][%d] updateEmploymentStatusUsingPUTNotFound", 404)
}

func (o *UpdateEmploymentStatusUsingPUTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateEmploymentStatusUsingPUTInternalServerError creates a UpdateEmploymentStatusUsingPUTInternalServerError with default headers values
func NewUpdateEmploymentStatusUsingPUTInternalServerError() *UpdateEmploymentStatusUsingPUTInternalServerError {
	return &UpdateEmploymentStatusUsingPUTInternalServerError{}
}

/*
UpdateEmploymentStatusUsingPUTInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type UpdateEmploymentStatusUsingPUTInternalServerError struct {
}

// IsSuccess returns true when this update employment status using p u t internal server error response has a 2xx status code
func (o *UpdateEmploymentStatusUsingPUTInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update employment status using p u t internal server error response has a 3xx status code
func (o *UpdateEmploymentStatusUsingPUTInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update employment status using p u t internal server error response has a 4xx status code
func (o *UpdateEmploymentStatusUsingPUTInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update employment status using p u t internal server error response has a 5xx status code
func (o *UpdateEmploymentStatusUsingPUTInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update employment status using p u t internal server error response a status code equal to that given
func (o *UpdateEmploymentStatusUsingPUTInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update employment status using p u t internal server error response
func (o *UpdateEmploymentStatusUsingPUTInternalServerError) Code() int {
	return 500
}

func (o *UpdateEmploymentStatusUsingPUTInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/member/update/employmentStatus][%d] updateEmploymentStatusUsingPUTInternalServerError", 500)
}

func (o *UpdateEmploymentStatusUsingPUTInternalServerError) String() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/member/update/employmentStatus][%d] updateEmploymentStatusUsingPUTInternalServerError", 500)
}

func (o *UpdateEmploymentStatusUsingPUTInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
