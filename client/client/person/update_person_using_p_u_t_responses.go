// Code generated by go-swagger; DO NOT EDIT.

package person

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

// UpdatePersonUsingPUTReader is a Reader for the UpdatePersonUsingPUT structure.
type UpdatePersonUsingPUTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdatePersonUsingPUTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdatePersonUsingPUTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewUpdatePersonUsingPUTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdatePersonUsingPUTBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdatePersonUsingPUTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdatePersonUsingPUTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdatePersonUsingPUTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdatePersonUsingPUTInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /compassintegrationapp/person] updatePersonUsingPUT", response, response.Code())
	}
}

// NewUpdatePersonUsingPUTOK creates a UpdatePersonUsingPUTOK with default headers values
func NewUpdatePersonUsingPUTOK() *UpdatePersonUsingPUTOK {
	return &UpdatePersonUsingPUTOK{}
}

/*
UpdatePersonUsingPUTOK describes a response with status code 200, with default header values.

Success
*/
type UpdatePersonUsingPUTOK struct {
	Payload *models.ResponseWrapperBoolean
}

// IsSuccess returns true when this update person using p u t o k response has a 2xx status code
func (o *UpdatePersonUsingPUTOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update person using p u t o k response has a 3xx status code
func (o *UpdatePersonUsingPUTOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update person using p u t o k response has a 4xx status code
func (o *UpdatePersonUsingPUTOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update person using p u t o k response has a 5xx status code
func (o *UpdatePersonUsingPUTOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update person using p u t o k response a status code equal to that given
func (o *UpdatePersonUsingPUTOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update person using p u t o k response
func (o *UpdatePersonUsingPUTOK) Code() int {
	return 200
}

func (o *UpdatePersonUsingPUTOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /compassintegrationapp/person][%d] updatePersonUsingPUTOK %s", 200, payload)
}

func (o *UpdatePersonUsingPUTOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /compassintegrationapp/person][%d] updatePersonUsingPUTOK %s", 200, payload)
}

func (o *UpdatePersonUsingPUTOK) GetPayload() *models.ResponseWrapperBoolean {
	return o.Payload
}

func (o *UpdatePersonUsingPUTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperBoolean)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdatePersonUsingPUTCreated creates a UpdatePersonUsingPUTCreated with default headers values
func NewUpdatePersonUsingPUTCreated() *UpdatePersonUsingPUTCreated {
	return &UpdatePersonUsingPUTCreated{}
}

/*
UpdatePersonUsingPUTCreated describes a response with status code 201, with default header values.

Created
*/
type UpdatePersonUsingPUTCreated struct {
}

// IsSuccess returns true when this update person using p u t created response has a 2xx status code
func (o *UpdatePersonUsingPUTCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update person using p u t created response has a 3xx status code
func (o *UpdatePersonUsingPUTCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update person using p u t created response has a 4xx status code
func (o *UpdatePersonUsingPUTCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this update person using p u t created response has a 5xx status code
func (o *UpdatePersonUsingPUTCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this update person using p u t created response a status code equal to that given
func (o *UpdatePersonUsingPUTCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the update person using p u t created response
func (o *UpdatePersonUsingPUTCreated) Code() int {
	return 201
}

func (o *UpdatePersonUsingPUTCreated) Error() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/person][%d] updatePersonUsingPUTCreated", 201)
}

func (o *UpdatePersonUsingPUTCreated) String() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/person][%d] updatePersonUsingPUTCreated", 201)
}

func (o *UpdatePersonUsingPUTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdatePersonUsingPUTBadRequest creates a UpdatePersonUsingPUTBadRequest with default headers values
func NewUpdatePersonUsingPUTBadRequest() *UpdatePersonUsingPUTBadRequest {
	return &UpdatePersonUsingPUTBadRequest{}
}

/*
UpdatePersonUsingPUTBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UpdatePersonUsingPUTBadRequest struct {
}

// IsSuccess returns true when this update person using p u t bad request response has a 2xx status code
func (o *UpdatePersonUsingPUTBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update person using p u t bad request response has a 3xx status code
func (o *UpdatePersonUsingPUTBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update person using p u t bad request response has a 4xx status code
func (o *UpdatePersonUsingPUTBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update person using p u t bad request response has a 5xx status code
func (o *UpdatePersonUsingPUTBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update person using p u t bad request response a status code equal to that given
func (o *UpdatePersonUsingPUTBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update person using p u t bad request response
func (o *UpdatePersonUsingPUTBadRequest) Code() int {
	return 400
}

func (o *UpdatePersonUsingPUTBadRequest) Error() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/person][%d] updatePersonUsingPUTBadRequest", 400)
}

func (o *UpdatePersonUsingPUTBadRequest) String() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/person][%d] updatePersonUsingPUTBadRequest", 400)
}

func (o *UpdatePersonUsingPUTBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdatePersonUsingPUTUnauthorized creates a UpdatePersonUsingPUTUnauthorized with default headers values
func NewUpdatePersonUsingPUTUnauthorized() *UpdatePersonUsingPUTUnauthorized {
	return &UpdatePersonUsingPUTUnauthorized{}
}

/*
UpdatePersonUsingPUTUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UpdatePersonUsingPUTUnauthorized struct {
}

// IsSuccess returns true when this update person using p u t unauthorized response has a 2xx status code
func (o *UpdatePersonUsingPUTUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update person using p u t unauthorized response has a 3xx status code
func (o *UpdatePersonUsingPUTUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update person using p u t unauthorized response has a 4xx status code
func (o *UpdatePersonUsingPUTUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update person using p u t unauthorized response has a 5xx status code
func (o *UpdatePersonUsingPUTUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update person using p u t unauthorized response a status code equal to that given
func (o *UpdatePersonUsingPUTUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the update person using p u t unauthorized response
func (o *UpdatePersonUsingPUTUnauthorized) Code() int {
	return 401
}

func (o *UpdatePersonUsingPUTUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/person][%d] updatePersonUsingPUTUnauthorized", 401)
}

func (o *UpdatePersonUsingPUTUnauthorized) String() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/person][%d] updatePersonUsingPUTUnauthorized", 401)
}

func (o *UpdatePersonUsingPUTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdatePersonUsingPUTForbidden creates a UpdatePersonUsingPUTForbidden with default headers values
func NewUpdatePersonUsingPUTForbidden() *UpdatePersonUsingPUTForbidden {
	return &UpdatePersonUsingPUTForbidden{}
}

/*
UpdatePersonUsingPUTForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdatePersonUsingPUTForbidden struct {
}

// IsSuccess returns true when this update person using p u t forbidden response has a 2xx status code
func (o *UpdatePersonUsingPUTForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update person using p u t forbidden response has a 3xx status code
func (o *UpdatePersonUsingPUTForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update person using p u t forbidden response has a 4xx status code
func (o *UpdatePersonUsingPUTForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update person using p u t forbidden response has a 5xx status code
func (o *UpdatePersonUsingPUTForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update person using p u t forbidden response a status code equal to that given
func (o *UpdatePersonUsingPUTForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update person using p u t forbidden response
func (o *UpdatePersonUsingPUTForbidden) Code() int {
	return 403
}

func (o *UpdatePersonUsingPUTForbidden) Error() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/person][%d] updatePersonUsingPUTForbidden", 403)
}

func (o *UpdatePersonUsingPUTForbidden) String() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/person][%d] updatePersonUsingPUTForbidden", 403)
}

func (o *UpdatePersonUsingPUTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdatePersonUsingPUTNotFound creates a UpdatePersonUsingPUTNotFound with default headers values
func NewUpdatePersonUsingPUTNotFound() *UpdatePersonUsingPUTNotFound {
	return &UpdatePersonUsingPUTNotFound{}
}

/*
UpdatePersonUsingPUTNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdatePersonUsingPUTNotFound struct {
}

// IsSuccess returns true when this update person using p u t not found response has a 2xx status code
func (o *UpdatePersonUsingPUTNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update person using p u t not found response has a 3xx status code
func (o *UpdatePersonUsingPUTNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update person using p u t not found response has a 4xx status code
func (o *UpdatePersonUsingPUTNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update person using p u t not found response has a 5xx status code
func (o *UpdatePersonUsingPUTNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update person using p u t not found response a status code equal to that given
func (o *UpdatePersonUsingPUTNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update person using p u t not found response
func (o *UpdatePersonUsingPUTNotFound) Code() int {
	return 404
}

func (o *UpdatePersonUsingPUTNotFound) Error() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/person][%d] updatePersonUsingPUTNotFound", 404)
}

func (o *UpdatePersonUsingPUTNotFound) String() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/person][%d] updatePersonUsingPUTNotFound", 404)
}

func (o *UpdatePersonUsingPUTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdatePersonUsingPUTInternalServerError creates a UpdatePersonUsingPUTInternalServerError with default headers values
func NewUpdatePersonUsingPUTInternalServerError() *UpdatePersonUsingPUTInternalServerError {
	return &UpdatePersonUsingPUTInternalServerError{}
}

/*
UpdatePersonUsingPUTInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type UpdatePersonUsingPUTInternalServerError struct {
}

// IsSuccess returns true when this update person using p u t internal server error response has a 2xx status code
func (o *UpdatePersonUsingPUTInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update person using p u t internal server error response has a 3xx status code
func (o *UpdatePersonUsingPUTInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update person using p u t internal server error response has a 4xx status code
func (o *UpdatePersonUsingPUTInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update person using p u t internal server error response has a 5xx status code
func (o *UpdatePersonUsingPUTInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update person using p u t internal server error response a status code equal to that given
func (o *UpdatePersonUsingPUTInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update person using p u t internal server error response
func (o *UpdatePersonUsingPUTInternalServerError) Code() int {
	return 500
}

func (o *UpdatePersonUsingPUTInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/person][%d] updatePersonUsingPUTInternalServerError", 500)
}

func (o *UpdatePersonUsingPUTInternalServerError) String() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/person][%d] updatePersonUsingPUTInternalServerError", 500)
}

func (o *UpdatePersonUsingPUTInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
