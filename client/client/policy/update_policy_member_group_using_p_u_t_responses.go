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

// UpdatePolicyMemberGroupUsingPUTReader is a Reader for the UpdatePolicyMemberGroupUsingPUT structure.
type UpdatePolicyMemberGroupUsingPUTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdatePolicyMemberGroupUsingPUTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdatePolicyMemberGroupUsingPUTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewUpdatePolicyMemberGroupUsingPUTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdatePolicyMemberGroupUsingPUTBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdatePolicyMemberGroupUsingPUTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdatePolicyMemberGroupUsingPUTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdatePolicyMemberGroupUsingPUTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdatePolicyMemberGroupUsingPUTInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /compassintegrationapp/policy/updatePolicyMemberGroup] updatePolicyMemberGroupUsingPUT", response, response.Code())
	}
}

// NewUpdatePolicyMemberGroupUsingPUTOK creates a UpdatePolicyMemberGroupUsingPUTOK with default headers values
func NewUpdatePolicyMemberGroupUsingPUTOK() *UpdatePolicyMemberGroupUsingPUTOK {
	return &UpdatePolicyMemberGroupUsingPUTOK{}
}

/*
UpdatePolicyMemberGroupUsingPUTOK describes a response with status code 200, with default header values.

Success
*/
type UpdatePolicyMemberGroupUsingPUTOK struct {
	Payload *models.ResponseWrapperBoolean
}

// IsSuccess returns true when this update policy member group using p u t o k response has a 2xx status code
func (o *UpdatePolicyMemberGroupUsingPUTOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update policy member group using p u t o k response has a 3xx status code
func (o *UpdatePolicyMemberGroupUsingPUTOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update policy member group using p u t o k response has a 4xx status code
func (o *UpdatePolicyMemberGroupUsingPUTOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update policy member group using p u t o k response has a 5xx status code
func (o *UpdatePolicyMemberGroupUsingPUTOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update policy member group using p u t o k response a status code equal to that given
func (o *UpdatePolicyMemberGroupUsingPUTOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update policy member group using p u t o k response
func (o *UpdatePolicyMemberGroupUsingPUTOK) Code() int {
	return 200
}

func (o *UpdatePolicyMemberGroupUsingPUTOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /compassintegrationapp/policy/updatePolicyMemberGroup][%d] updatePolicyMemberGroupUsingPUTOK %s", 200, payload)
}

func (o *UpdatePolicyMemberGroupUsingPUTOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /compassintegrationapp/policy/updatePolicyMemberGroup][%d] updatePolicyMemberGroupUsingPUTOK %s", 200, payload)
}

func (o *UpdatePolicyMemberGroupUsingPUTOK) GetPayload() *models.ResponseWrapperBoolean {
	return o.Payload
}

func (o *UpdatePolicyMemberGroupUsingPUTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperBoolean)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdatePolicyMemberGroupUsingPUTCreated creates a UpdatePolicyMemberGroupUsingPUTCreated with default headers values
func NewUpdatePolicyMemberGroupUsingPUTCreated() *UpdatePolicyMemberGroupUsingPUTCreated {
	return &UpdatePolicyMemberGroupUsingPUTCreated{}
}

/*
UpdatePolicyMemberGroupUsingPUTCreated describes a response with status code 201, with default header values.

Created
*/
type UpdatePolicyMemberGroupUsingPUTCreated struct {
}

// IsSuccess returns true when this update policy member group using p u t created response has a 2xx status code
func (o *UpdatePolicyMemberGroupUsingPUTCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update policy member group using p u t created response has a 3xx status code
func (o *UpdatePolicyMemberGroupUsingPUTCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update policy member group using p u t created response has a 4xx status code
func (o *UpdatePolicyMemberGroupUsingPUTCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this update policy member group using p u t created response has a 5xx status code
func (o *UpdatePolicyMemberGroupUsingPUTCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this update policy member group using p u t created response a status code equal to that given
func (o *UpdatePolicyMemberGroupUsingPUTCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the update policy member group using p u t created response
func (o *UpdatePolicyMemberGroupUsingPUTCreated) Code() int {
	return 201
}

func (o *UpdatePolicyMemberGroupUsingPUTCreated) Error() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/policy/updatePolicyMemberGroup][%d] updatePolicyMemberGroupUsingPUTCreated", 201)
}

func (o *UpdatePolicyMemberGroupUsingPUTCreated) String() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/policy/updatePolicyMemberGroup][%d] updatePolicyMemberGroupUsingPUTCreated", 201)
}

func (o *UpdatePolicyMemberGroupUsingPUTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdatePolicyMemberGroupUsingPUTBadRequest creates a UpdatePolicyMemberGroupUsingPUTBadRequest with default headers values
func NewUpdatePolicyMemberGroupUsingPUTBadRequest() *UpdatePolicyMemberGroupUsingPUTBadRequest {
	return &UpdatePolicyMemberGroupUsingPUTBadRequest{}
}

/*
UpdatePolicyMemberGroupUsingPUTBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UpdatePolicyMemberGroupUsingPUTBadRequest struct {
}

// IsSuccess returns true when this update policy member group using p u t bad request response has a 2xx status code
func (o *UpdatePolicyMemberGroupUsingPUTBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update policy member group using p u t bad request response has a 3xx status code
func (o *UpdatePolicyMemberGroupUsingPUTBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update policy member group using p u t bad request response has a 4xx status code
func (o *UpdatePolicyMemberGroupUsingPUTBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update policy member group using p u t bad request response has a 5xx status code
func (o *UpdatePolicyMemberGroupUsingPUTBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update policy member group using p u t bad request response a status code equal to that given
func (o *UpdatePolicyMemberGroupUsingPUTBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update policy member group using p u t bad request response
func (o *UpdatePolicyMemberGroupUsingPUTBadRequest) Code() int {
	return 400
}

func (o *UpdatePolicyMemberGroupUsingPUTBadRequest) Error() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/policy/updatePolicyMemberGroup][%d] updatePolicyMemberGroupUsingPUTBadRequest", 400)
}

func (o *UpdatePolicyMemberGroupUsingPUTBadRequest) String() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/policy/updatePolicyMemberGroup][%d] updatePolicyMemberGroupUsingPUTBadRequest", 400)
}

func (o *UpdatePolicyMemberGroupUsingPUTBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdatePolicyMemberGroupUsingPUTUnauthorized creates a UpdatePolicyMemberGroupUsingPUTUnauthorized with default headers values
func NewUpdatePolicyMemberGroupUsingPUTUnauthorized() *UpdatePolicyMemberGroupUsingPUTUnauthorized {
	return &UpdatePolicyMemberGroupUsingPUTUnauthorized{}
}

/*
UpdatePolicyMemberGroupUsingPUTUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UpdatePolicyMemberGroupUsingPUTUnauthorized struct {
}

// IsSuccess returns true when this update policy member group using p u t unauthorized response has a 2xx status code
func (o *UpdatePolicyMemberGroupUsingPUTUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update policy member group using p u t unauthorized response has a 3xx status code
func (o *UpdatePolicyMemberGroupUsingPUTUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update policy member group using p u t unauthorized response has a 4xx status code
func (o *UpdatePolicyMemberGroupUsingPUTUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update policy member group using p u t unauthorized response has a 5xx status code
func (o *UpdatePolicyMemberGroupUsingPUTUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update policy member group using p u t unauthorized response a status code equal to that given
func (o *UpdatePolicyMemberGroupUsingPUTUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the update policy member group using p u t unauthorized response
func (o *UpdatePolicyMemberGroupUsingPUTUnauthorized) Code() int {
	return 401
}

func (o *UpdatePolicyMemberGroupUsingPUTUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/policy/updatePolicyMemberGroup][%d] updatePolicyMemberGroupUsingPUTUnauthorized", 401)
}

func (o *UpdatePolicyMemberGroupUsingPUTUnauthorized) String() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/policy/updatePolicyMemberGroup][%d] updatePolicyMemberGroupUsingPUTUnauthorized", 401)
}

func (o *UpdatePolicyMemberGroupUsingPUTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdatePolicyMemberGroupUsingPUTForbidden creates a UpdatePolicyMemberGroupUsingPUTForbidden with default headers values
func NewUpdatePolicyMemberGroupUsingPUTForbidden() *UpdatePolicyMemberGroupUsingPUTForbidden {
	return &UpdatePolicyMemberGroupUsingPUTForbidden{}
}

/*
UpdatePolicyMemberGroupUsingPUTForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdatePolicyMemberGroupUsingPUTForbidden struct {
}

// IsSuccess returns true when this update policy member group using p u t forbidden response has a 2xx status code
func (o *UpdatePolicyMemberGroupUsingPUTForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update policy member group using p u t forbidden response has a 3xx status code
func (o *UpdatePolicyMemberGroupUsingPUTForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update policy member group using p u t forbidden response has a 4xx status code
func (o *UpdatePolicyMemberGroupUsingPUTForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update policy member group using p u t forbidden response has a 5xx status code
func (o *UpdatePolicyMemberGroupUsingPUTForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update policy member group using p u t forbidden response a status code equal to that given
func (o *UpdatePolicyMemberGroupUsingPUTForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update policy member group using p u t forbidden response
func (o *UpdatePolicyMemberGroupUsingPUTForbidden) Code() int {
	return 403
}

func (o *UpdatePolicyMemberGroupUsingPUTForbidden) Error() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/policy/updatePolicyMemberGroup][%d] updatePolicyMemberGroupUsingPUTForbidden", 403)
}

func (o *UpdatePolicyMemberGroupUsingPUTForbidden) String() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/policy/updatePolicyMemberGroup][%d] updatePolicyMemberGroupUsingPUTForbidden", 403)
}

func (o *UpdatePolicyMemberGroupUsingPUTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdatePolicyMemberGroupUsingPUTNotFound creates a UpdatePolicyMemberGroupUsingPUTNotFound with default headers values
func NewUpdatePolicyMemberGroupUsingPUTNotFound() *UpdatePolicyMemberGroupUsingPUTNotFound {
	return &UpdatePolicyMemberGroupUsingPUTNotFound{}
}

/*
UpdatePolicyMemberGroupUsingPUTNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdatePolicyMemberGroupUsingPUTNotFound struct {
}

// IsSuccess returns true when this update policy member group using p u t not found response has a 2xx status code
func (o *UpdatePolicyMemberGroupUsingPUTNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update policy member group using p u t not found response has a 3xx status code
func (o *UpdatePolicyMemberGroupUsingPUTNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update policy member group using p u t not found response has a 4xx status code
func (o *UpdatePolicyMemberGroupUsingPUTNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update policy member group using p u t not found response has a 5xx status code
func (o *UpdatePolicyMemberGroupUsingPUTNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update policy member group using p u t not found response a status code equal to that given
func (o *UpdatePolicyMemberGroupUsingPUTNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update policy member group using p u t not found response
func (o *UpdatePolicyMemberGroupUsingPUTNotFound) Code() int {
	return 404
}

func (o *UpdatePolicyMemberGroupUsingPUTNotFound) Error() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/policy/updatePolicyMemberGroup][%d] updatePolicyMemberGroupUsingPUTNotFound", 404)
}

func (o *UpdatePolicyMemberGroupUsingPUTNotFound) String() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/policy/updatePolicyMemberGroup][%d] updatePolicyMemberGroupUsingPUTNotFound", 404)
}

func (o *UpdatePolicyMemberGroupUsingPUTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdatePolicyMemberGroupUsingPUTInternalServerError creates a UpdatePolicyMemberGroupUsingPUTInternalServerError with default headers values
func NewUpdatePolicyMemberGroupUsingPUTInternalServerError() *UpdatePolicyMemberGroupUsingPUTInternalServerError {
	return &UpdatePolicyMemberGroupUsingPUTInternalServerError{}
}

/*
UpdatePolicyMemberGroupUsingPUTInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type UpdatePolicyMemberGroupUsingPUTInternalServerError struct {
}

// IsSuccess returns true when this update policy member group using p u t internal server error response has a 2xx status code
func (o *UpdatePolicyMemberGroupUsingPUTInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update policy member group using p u t internal server error response has a 3xx status code
func (o *UpdatePolicyMemberGroupUsingPUTInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update policy member group using p u t internal server error response has a 4xx status code
func (o *UpdatePolicyMemberGroupUsingPUTInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update policy member group using p u t internal server error response has a 5xx status code
func (o *UpdatePolicyMemberGroupUsingPUTInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update policy member group using p u t internal server error response a status code equal to that given
func (o *UpdatePolicyMemberGroupUsingPUTInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update policy member group using p u t internal server error response
func (o *UpdatePolicyMemberGroupUsingPUTInternalServerError) Code() int {
	return 500
}

func (o *UpdatePolicyMemberGroupUsingPUTInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/policy/updatePolicyMemberGroup][%d] updatePolicyMemberGroupUsingPUTInternalServerError", 500)
}

func (o *UpdatePolicyMemberGroupUsingPUTInternalServerError) String() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/policy/updatePolicyMemberGroup][%d] updatePolicyMemberGroupUsingPUTInternalServerError", 500)
}

func (o *UpdatePolicyMemberGroupUsingPUTInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
