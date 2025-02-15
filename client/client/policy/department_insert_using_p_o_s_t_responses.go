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

// DepartmentInsertUsingPOSTReader is a Reader for the DepartmentInsertUsingPOST structure.
type DepartmentInsertUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DepartmentInsertUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDepartmentInsertUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewDepartmentInsertUsingPOSTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDepartmentInsertUsingPOSTBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDepartmentInsertUsingPOSTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDepartmentInsertUsingPOSTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDepartmentInsertUsingPOSTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDepartmentInsertUsingPOSTInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /compassintegrationapp/policy/departmentInsert] departmentInsertUsingPOST", response, response.Code())
	}
}

// NewDepartmentInsertUsingPOSTOK creates a DepartmentInsertUsingPOSTOK with default headers values
func NewDepartmentInsertUsingPOSTOK() *DepartmentInsertUsingPOSTOK {
	return &DepartmentInsertUsingPOSTOK{}
}

/*
DepartmentInsertUsingPOSTOK describes a response with status code 200, with default header values.

Success
*/
type DepartmentInsertUsingPOSTOK struct {
	Payload *models.ResponseWrapperDepartmentDTO
}

// IsSuccess returns true when this department insert using p o s t o k response has a 2xx status code
func (o *DepartmentInsertUsingPOSTOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this department insert using p o s t o k response has a 3xx status code
func (o *DepartmentInsertUsingPOSTOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this department insert using p o s t o k response has a 4xx status code
func (o *DepartmentInsertUsingPOSTOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this department insert using p o s t o k response has a 5xx status code
func (o *DepartmentInsertUsingPOSTOK) IsServerError() bool {
	return false
}

// IsCode returns true when this department insert using p o s t o k response a status code equal to that given
func (o *DepartmentInsertUsingPOSTOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the department insert using p o s t o k response
func (o *DepartmentInsertUsingPOSTOK) Code() int {
	return 200
}

func (o *DepartmentInsertUsingPOSTOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /compassintegrationapp/policy/departmentInsert][%d] departmentInsertUsingPOSTOK %s", 200, payload)
}

func (o *DepartmentInsertUsingPOSTOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /compassintegrationapp/policy/departmentInsert][%d] departmentInsertUsingPOSTOK %s", 200, payload)
}

func (o *DepartmentInsertUsingPOSTOK) GetPayload() *models.ResponseWrapperDepartmentDTO {
	return o.Payload
}

func (o *DepartmentInsertUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperDepartmentDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDepartmentInsertUsingPOSTCreated creates a DepartmentInsertUsingPOSTCreated with default headers values
func NewDepartmentInsertUsingPOSTCreated() *DepartmentInsertUsingPOSTCreated {
	return &DepartmentInsertUsingPOSTCreated{}
}

/*
DepartmentInsertUsingPOSTCreated describes a response with status code 201, with default header values.

Created
*/
type DepartmentInsertUsingPOSTCreated struct {
}

// IsSuccess returns true when this department insert using p o s t created response has a 2xx status code
func (o *DepartmentInsertUsingPOSTCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this department insert using p o s t created response has a 3xx status code
func (o *DepartmentInsertUsingPOSTCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this department insert using p o s t created response has a 4xx status code
func (o *DepartmentInsertUsingPOSTCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this department insert using p o s t created response has a 5xx status code
func (o *DepartmentInsertUsingPOSTCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this department insert using p o s t created response a status code equal to that given
func (o *DepartmentInsertUsingPOSTCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the department insert using p o s t created response
func (o *DepartmentInsertUsingPOSTCreated) Code() int {
	return 201
}

func (o *DepartmentInsertUsingPOSTCreated) Error() string {
	return fmt.Sprintf("[POST /compassintegrationapp/policy/departmentInsert][%d] departmentInsertUsingPOSTCreated", 201)
}

func (o *DepartmentInsertUsingPOSTCreated) String() string {
	return fmt.Sprintf("[POST /compassintegrationapp/policy/departmentInsert][%d] departmentInsertUsingPOSTCreated", 201)
}

func (o *DepartmentInsertUsingPOSTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDepartmentInsertUsingPOSTBadRequest creates a DepartmentInsertUsingPOSTBadRequest with default headers values
func NewDepartmentInsertUsingPOSTBadRequest() *DepartmentInsertUsingPOSTBadRequest {
	return &DepartmentInsertUsingPOSTBadRequest{}
}

/*
DepartmentInsertUsingPOSTBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DepartmentInsertUsingPOSTBadRequest struct {
}

// IsSuccess returns true when this department insert using p o s t bad request response has a 2xx status code
func (o *DepartmentInsertUsingPOSTBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this department insert using p o s t bad request response has a 3xx status code
func (o *DepartmentInsertUsingPOSTBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this department insert using p o s t bad request response has a 4xx status code
func (o *DepartmentInsertUsingPOSTBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this department insert using p o s t bad request response has a 5xx status code
func (o *DepartmentInsertUsingPOSTBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this department insert using p o s t bad request response a status code equal to that given
func (o *DepartmentInsertUsingPOSTBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the department insert using p o s t bad request response
func (o *DepartmentInsertUsingPOSTBadRequest) Code() int {
	return 400
}

func (o *DepartmentInsertUsingPOSTBadRequest) Error() string {
	return fmt.Sprintf("[POST /compassintegrationapp/policy/departmentInsert][%d] departmentInsertUsingPOSTBadRequest", 400)
}

func (o *DepartmentInsertUsingPOSTBadRequest) String() string {
	return fmt.Sprintf("[POST /compassintegrationapp/policy/departmentInsert][%d] departmentInsertUsingPOSTBadRequest", 400)
}

func (o *DepartmentInsertUsingPOSTBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDepartmentInsertUsingPOSTUnauthorized creates a DepartmentInsertUsingPOSTUnauthorized with default headers values
func NewDepartmentInsertUsingPOSTUnauthorized() *DepartmentInsertUsingPOSTUnauthorized {
	return &DepartmentInsertUsingPOSTUnauthorized{}
}

/*
DepartmentInsertUsingPOSTUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DepartmentInsertUsingPOSTUnauthorized struct {
}

// IsSuccess returns true when this department insert using p o s t unauthorized response has a 2xx status code
func (o *DepartmentInsertUsingPOSTUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this department insert using p o s t unauthorized response has a 3xx status code
func (o *DepartmentInsertUsingPOSTUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this department insert using p o s t unauthorized response has a 4xx status code
func (o *DepartmentInsertUsingPOSTUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this department insert using p o s t unauthorized response has a 5xx status code
func (o *DepartmentInsertUsingPOSTUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this department insert using p o s t unauthorized response a status code equal to that given
func (o *DepartmentInsertUsingPOSTUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the department insert using p o s t unauthorized response
func (o *DepartmentInsertUsingPOSTUnauthorized) Code() int {
	return 401
}

func (o *DepartmentInsertUsingPOSTUnauthorized) Error() string {
	return fmt.Sprintf("[POST /compassintegrationapp/policy/departmentInsert][%d] departmentInsertUsingPOSTUnauthorized", 401)
}

func (o *DepartmentInsertUsingPOSTUnauthorized) String() string {
	return fmt.Sprintf("[POST /compassintegrationapp/policy/departmentInsert][%d] departmentInsertUsingPOSTUnauthorized", 401)
}

func (o *DepartmentInsertUsingPOSTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDepartmentInsertUsingPOSTForbidden creates a DepartmentInsertUsingPOSTForbidden with default headers values
func NewDepartmentInsertUsingPOSTForbidden() *DepartmentInsertUsingPOSTForbidden {
	return &DepartmentInsertUsingPOSTForbidden{}
}

/*
DepartmentInsertUsingPOSTForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DepartmentInsertUsingPOSTForbidden struct {
}

// IsSuccess returns true when this department insert using p o s t forbidden response has a 2xx status code
func (o *DepartmentInsertUsingPOSTForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this department insert using p o s t forbidden response has a 3xx status code
func (o *DepartmentInsertUsingPOSTForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this department insert using p o s t forbidden response has a 4xx status code
func (o *DepartmentInsertUsingPOSTForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this department insert using p o s t forbidden response has a 5xx status code
func (o *DepartmentInsertUsingPOSTForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this department insert using p o s t forbidden response a status code equal to that given
func (o *DepartmentInsertUsingPOSTForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the department insert using p o s t forbidden response
func (o *DepartmentInsertUsingPOSTForbidden) Code() int {
	return 403
}

func (o *DepartmentInsertUsingPOSTForbidden) Error() string {
	return fmt.Sprintf("[POST /compassintegrationapp/policy/departmentInsert][%d] departmentInsertUsingPOSTForbidden", 403)
}

func (o *DepartmentInsertUsingPOSTForbidden) String() string {
	return fmt.Sprintf("[POST /compassintegrationapp/policy/departmentInsert][%d] departmentInsertUsingPOSTForbidden", 403)
}

func (o *DepartmentInsertUsingPOSTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDepartmentInsertUsingPOSTNotFound creates a DepartmentInsertUsingPOSTNotFound with default headers values
func NewDepartmentInsertUsingPOSTNotFound() *DepartmentInsertUsingPOSTNotFound {
	return &DepartmentInsertUsingPOSTNotFound{}
}

/*
DepartmentInsertUsingPOSTNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DepartmentInsertUsingPOSTNotFound struct {
}

// IsSuccess returns true when this department insert using p o s t not found response has a 2xx status code
func (o *DepartmentInsertUsingPOSTNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this department insert using p o s t not found response has a 3xx status code
func (o *DepartmentInsertUsingPOSTNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this department insert using p o s t not found response has a 4xx status code
func (o *DepartmentInsertUsingPOSTNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this department insert using p o s t not found response has a 5xx status code
func (o *DepartmentInsertUsingPOSTNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this department insert using p o s t not found response a status code equal to that given
func (o *DepartmentInsertUsingPOSTNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the department insert using p o s t not found response
func (o *DepartmentInsertUsingPOSTNotFound) Code() int {
	return 404
}

func (o *DepartmentInsertUsingPOSTNotFound) Error() string {
	return fmt.Sprintf("[POST /compassintegrationapp/policy/departmentInsert][%d] departmentInsertUsingPOSTNotFound", 404)
}

func (o *DepartmentInsertUsingPOSTNotFound) String() string {
	return fmt.Sprintf("[POST /compassintegrationapp/policy/departmentInsert][%d] departmentInsertUsingPOSTNotFound", 404)
}

func (o *DepartmentInsertUsingPOSTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDepartmentInsertUsingPOSTInternalServerError creates a DepartmentInsertUsingPOSTInternalServerError with default headers values
func NewDepartmentInsertUsingPOSTInternalServerError() *DepartmentInsertUsingPOSTInternalServerError {
	return &DepartmentInsertUsingPOSTInternalServerError{}
}

/*
DepartmentInsertUsingPOSTInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type DepartmentInsertUsingPOSTInternalServerError struct {
}

// IsSuccess returns true when this department insert using p o s t internal server error response has a 2xx status code
func (o *DepartmentInsertUsingPOSTInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this department insert using p o s t internal server error response has a 3xx status code
func (o *DepartmentInsertUsingPOSTInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this department insert using p o s t internal server error response has a 4xx status code
func (o *DepartmentInsertUsingPOSTInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this department insert using p o s t internal server error response has a 5xx status code
func (o *DepartmentInsertUsingPOSTInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this department insert using p o s t internal server error response a status code equal to that given
func (o *DepartmentInsertUsingPOSTInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the department insert using p o s t internal server error response
func (o *DepartmentInsertUsingPOSTInternalServerError) Code() int {
	return 500
}

func (o *DepartmentInsertUsingPOSTInternalServerError) Error() string {
	return fmt.Sprintf("[POST /compassintegrationapp/policy/departmentInsert][%d] departmentInsertUsingPOSTInternalServerError", 500)
}

func (o *DepartmentInsertUsingPOSTInternalServerError) String() string {
	return fmt.Sprintf("[POST /compassintegrationapp/policy/departmentInsert][%d] departmentInsertUsingPOSTInternalServerError", 500)
}

func (o *DepartmentInsertUsingPOSTInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
