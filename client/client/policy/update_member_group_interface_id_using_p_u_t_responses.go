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

// UpdateMemberGroupInterfaceIDUsingPUTReader is a Reader for the UpdateMemberGroupInterfaceIDUsingPUT structure.
type UpdateMemberGroupInterfaceIDUsingPUTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateMemberGroupInterfaceIDUsingPUTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateMemberGroupInterfaceIDUsingPUTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewUpdateMemberGroupInterfaceIDUsingPUTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateMemberGroupInterfaceIDUsingPUTBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateMemberGroupInterfaceIDUsingPUTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateMemberGroupInterfaceIDUsingPUTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateMemberGroupInterfaceIDUsingPUTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateMemberGroupInterfaceIDUsingPUTInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /compassintegrationapp/policy/memberGroup/interfaceId] updateMemberGroupInterfaceIdUsingPUT", response, response.Code())
	}
}

// NewUpdateMemberGroupInterfaceIDUsingPUTOK creates a UpdateMemberGroupInterfaceIDUsingPUTOK with default headers values
func NewUpdateMemberGroupInterfaceIDUsingPUTOK() *UpdateMemberGroupInterfaceIDUsingPUTOK {
	return &UpdateMemberGroupInterfaceIDUsingPUTOK{}
}

/*
UpdateMemberGroupInterfaceIDUsingPUTOK describes a response with status code 200, with default header values.

Success
*/
type UpdateMemberGroupInterfaceIDUsingPUTOK struct {
	Payload *models.ResponseWrapperBoolean
}

// IsSuccess returns true when this update member group interface Id using p u t o k response has a 2xx status code
func (o *UpdateMemberGroupInterfaceIDUsingPUTOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update member group interface Id using p u t o k response has a 3xx status code
func (o *UpdateMemberGroupInterfaceIDUsingPUTOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update member group interface Id using p u t o k response has a 4xx status code
func (o *UpdateMemberGroupInterfaceIDUsingPUTOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update member group interface Id using p u t o k response has a 5xx status code
func (o *UpdateMemberGroupInterfaceIDUsingPUTOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update member group interface Id using p u t o k response a status code equal to that given
func (o *UpdateMemberGroupInterfaceIDUsingPUTOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update member group interface Id using p u t o k response
func (o *UpdateMemberGroupInterfaceIDUsingPUTOK) Code() int {
	return 200
}

func (o *UpdateMemberGroupInterfaceIDUsingPUTOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /compassintegrationapp/policy/memberGroup/interfaceId][%d] updateMemberGroupInterfaceIdUsingPUTOK %s", 200, payload)
}

func (o *UpdateMemberGroupInterfaceIDUsingPUTOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /compassintegrationapp/policy/memberGroup/interfaceId][%d] updateMemberGroupInterfaceIdUsingPUTOK %s", 200, payload)
}

func (o *UpdateMemberGroupInterfaceIDUsingPUTOK) GetPayload() *models.ResponseWrapperBoolean {
	return o.Payload
}

func (o *UpdateMemberGroupInterfaceIDUsingPUTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperBoolean)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMemberGroupInterfaceIDUsingPUTCreated creates a UpdateMemberGroupInterfaceIDUsingPUTCreated with default headers values
func NewUpdateMemberGroupInterfaceIDUsingPUTCreated() *UpdateMemberGroupInterfaceIDUsingPUTCreated {
	return &UpdateMemberGroupInterfaceIDUsingPUTCreated{}
}

/*
UpdateMemberGroupInterfaceIDUsingPUTCreated describes a response with status code 201, with default header values.

Created
*/
type UpdateMemberGroupInterfaceIDUsingPUTCreated struct {
}

// IsSuccess returns true when this update member group interface Id using p u t created response has a 2xx status code
func (o *UpdateMemberGroupInterfaceIDUsingPUTCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update member group interface Id using p u t created response has a 3xx status code
func (o *UpdateMemberGroupInterfaceIDUsingPUTCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update member group interface Id using p u t created response has a 4xx status code
func (o *UpdateMemberGroupInterfaceIDUsingPUTCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this update member group interface Id using p u t created response has a 5xx status code
func (o *UpdateMemberGroupInterfaceIDUsingPUTCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this update member group interface Id using p u t created response a status code equal to that given
func (o *UpdateMemberGroupInterfaceIDUsingPUTCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the update member group interface Id using p u t created response
func (o *UpdateMemberGroupInterfaceIDUsingPUTCreated) Code() int {
	return 201
}

func (o *UpdateMemberGroupInterfaceIDUsingPUTCreated) Error() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/policy/memberGroup/interfaceId][%d] updateMemberGroupInterfaceIdUsingPUTCreated", 201)
}

func (o *UpdateMemberGroupInterfaceIDUsingPUTCreated) String() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/policy/memberGroup/interfaceId][%d] updateMemberGroupInterfaceIdUsingPUTCreated", 201)
}

func (o *UpdateMemberGroupInterfaceIDUsingPUTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateMemberGroupInterfaceIDUsingPUTBadRequest creates a UpdateMemberGroupInterfaceIDUsingPUTBadRequest with default headers values
func NewUpdateMemberGroupInterfaceIDUsingPUTBadRequest() *UpdateMemberGroupInterfaceIDUsingPUTBadRequest {
	return &UpdateMemberGroupInterfaceIDUsingPUTBadRequest{}
}

/*
UpdateMemberGroupInterfaceIDUsingPUTBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UpdateMemberGroupInterfaceIDUsingPUTBadRequest struct {
}

// IsSuccess returns true when this update member group interface Id using p u t bad request response has a 2xx status code
func (o *UpdateMemberGroupInterfaceIDUsingPUTBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update member group interface Id using p u t bad request response has a 3xx status code
func (o *UpdateMemberGroupInterfaceIDUsingPUTBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update member group interface Id using p u t bad request response has a 4xx status code
func (o *UpdateMemberGroupInterfaceIDUsingPUTBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update member group interface Id using p u t bad request response has a 5xx status code
func (o *UpdateMemberGroupInterfaceIDUsingPUTBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update member group interface Id using p u t bad request response a status code equal to that given
func (o *UpdateMemberGroupInterfaceIDUsingPUTBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update member group interface Id using p u t bad request response
func (o *UpdateMemberGroupInterfaceIDUsingPUTBadRequest) Code() int {
	return 400
}

func (o *UpdateMemberGroupInterfaceIDUsingPUTBadRequest) Error() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/policy/memberGroup/interfaceId][%d] updateMemberGroupInterfaceIdUsingPUTBadRequest", 400)
}

func (o *UpdateMemberGroupInterfaceIDUsingPUTBadRequest) String() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/policy/memberGroup/interfaceId][%d] updateMemberGroupInterfaceIdUsingPUTBadRequest", 400)
}

func (o *UpdateMemberGroupInterfaceIDUsingPUTBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateMemberGroupInterfaceIDUsingPUTUnauthorized creates a UpdateMemberGroupInterfaceIDUsingPUTUnauthorized with default headers values
func NewUpdateMemberGroupInterfaceIDUsingPUTUnauthorized() *UpdateMemberGroupInterfaceIDUsingPUTUnauthorized {
	return &UpdateMemberGroupInterfaceIDUsingPUTUnauthorized{}
}

/*
UpdateMemberGroupInterfaceIDUsingPUTUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UpdateMemberGroupInterfaceIDUsingPUTUnauthorized struct {
}

// IsSuccess returns true when this update member group interface Id using p u t unauthorized response has a 2xx status code
func (o *UpdateMemberGroupInterfaceIDUsingPUTUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update member group interface Id using p u t unauthorized response has a 3xx status code
func (o *UpdateMemberGroupInterfaceIDUsingPUTUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update member group interface Id using p u t unauthorized response has a 4xx status code
func (o *UpdateMemberGroupInterfaceIDUsingPUTUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update member group interface Id using p u t unauthorized response has a 5xx status code
func (o *UpdateMemberGroupInterfaceIDUsingPUTUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update member group interface Id using p u t unauthorized response a status code equal to that given
func (o *UpdateMemberGroupInterfaceIDUsingPUTUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the update member group interface Id using p u t unauthorized response
func (o *UpdateMemberGroupInterfaceIDUsingPUTUnauthorized) Code() int {
	return 401
}

func (o *UpdateMemberGroupInterfaceIDUsingPUTUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/policy/memberGroup/interfaceId][%d] updateMemberGroupInterfaceIdUsingPUTUnauthorized", 401)
}

func (o *UpdateMemberGroupInterfaceIDUsingPUTUnauthorized) String() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/policy/memberGroup/interfaceId][%d] updateMemberGroupInterfaceIdUsingPUTUnauthorized", 401)
}

func (o *UpdateMemberGroupInterfaceIDUsingPUTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateMemberGroupInterfaceIDUsingPUTForbidden creates a UpdateMemberGroupInterfaceIDUsingPUTForbidden with default headers values
func NewUpdateMemberGroupInterfaceIDUsingPUTForbidden() *UpdateMemberGroupInterfaceIDUsingPUTForbidden {
	return &UpdateMemberGroupInterfaceIDUsingPUTForbidden{}
}

/*
UpdateMemberGroupInterfaceIDUsingPUTForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateMemberGroupInterfaceIDUsingPUTForbidden struct {
}

// IsSuccess returns true when this update member group interface Id using p u t forbidden response has a 2xx status code
func (o *UpdateMemberGroupInterfaceIDUsingPUTForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update member group interface Id using p u t forbidden response has a 3xx status code
func (o *UpdateMemberGroupInterfaceIDUsingPUTForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update member group interface Id using p u t forbidden response has a 4xx status code
func (o *UpdateMemberGroupInterfaceIDUsingPUTForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update member group interface Id using p u t forbidden response has a 5xx status code
func (o *UpdateMemberGroupInterfaceIDUsingPUTForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update member group interface Id using p u t forbidden response a status code equal to that given
func (o *UpdateMemberGroupInterfaceIDUsingPUTForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update member group interface Id using p u t forbidden response
func (o *UpdateMemberGroupInterfaceIDUsingPUTForbidden) Code() int {
	return 403
}

func (o *UpdateMemberGroupInterfaceIDUsingPUTForbidden) Error() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/policy/memberGroup/interfaceId][%d] updateMemberGroupInterfaceIdUsingPUTForbidden", 403)
}

func (o *UpdateMemberGroupInterfaceIDUsingPUTForbidden) String() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/policy/memberGroup/interfaceId][%d] updateMemberGroupInterfaceIdUsingPUTForbidden", 403)
}

func (o *UpdateMemberGroupInterfaceIDUsingPUTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateMemberGroupInterfaceIDUsingPUTNotFound creates a UpdateMemberGroupInterfaceIDUsingPUTNotFound with default headers values
func NewUpdateMemberGroupInterfaceIDUsingPUTNotFound() *UpdateMemberGroupInterfaceIDUsingPUTNotFound {
	return &UpdateMemberGroupInterfaceIDUsingPUTNotFound{}
}

/*
UpdateMemberGroupInterfaceIDUsingPUTNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdateMemberGroupInterfaceIDUsingPUTNotFound struct {
}

// IsSuccess returns true when this update member group interface Id using p u t not found response has a 2xx status code
func (o *UpdateMemberGroupInterfaceIDUsingPUTNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update member group interface Id using p u t not found response has a 3xx status code
func (o *UpdateMemberGroupInterfaceIDUsingPUTNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update member group interface Id using p u t not found response has a 4xx status code
func (o *UpdateMemberGroupInterfaceIDUsingPUTNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update member group interface Id using p u t not found response has a 5xx status code
func (o *UpdateMemberGroupInterfaceIDUsingPUTNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update member group interface Id using p u t not found response a status code equal to that given
func (o *UpdateMemberGroupInterfaceIDUsingPUTNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update member group interface Id using p u t not found response
func (o *UpdateMemberGroupInterfaceIDUsingPUTNotFound) Code() int {
	return 404
}

func (o *UpdateMemberGroupInterfaceIDUsingPUTNotFound) Error() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/policy/memberGroup/interfaceId][%d] updateMemberGroupInterfaceIdUsingPUTNotFound", 404)
}

func (o *UpdateMemberGroupInterfaceIDUsingPUTNotFound) String() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/policy/memberGroup/interfaceId][%d] updateMemberGroupInterfaceIdUsingPUTNotFound", 404)
}

func (o *UpdateMemberGroupInterfaceIDUsingPUTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateMemberGroupInterfaceIDUsingPUTInternalServerError creates a UpdateMemberGroupInterfaceIDUsingPUTInternalServerError with default headers values
func NewUpdateMemberGroupInterfaceIDUsingPUTInternalServerError() *UpdateMemberGroupInterfaceIDUsingPUTInternalServerError {
	return &UpdateMemberGroupInterfaceIDUsingPUTInternalServerError{}
}

/*
UpdateMemberGroupInterfaceIDUsingPUTInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type UpdateMemberGroupInterfaceIDUsingPUTInternalServerError struct {
}

// IsSuccess returns true when this update member group interface Id using p u t internal server error response has a 2xx status code
func (o *UpdateMemberGroupInterfaceIDUsingPUTInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update member group interface Id using p u t internal server error response has a 3xx status code
func (o *UpdateMemberGroupInterfaceIDUsingPUTInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update member group interface Id using p u t internal server error response has a 4xx status code
func (o *UpdateMemberGroupInterfaceIDUsingPUTInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update member group interface Id using p u t internal server error response has a 5xx status code
func (o *UpdateMemberGroupInterfaceIDUsingPUTInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update member group interface Id using p u t internal server error response a status code equal to that given
func (o *UpdateMemberGroupInterfaceIDUsingPUTInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update member group interface Id using p u t internal server error response
func (o *UpdateMemberGroupInterfaceIDUsingPUTInternalServerError) Code() int {
	return 500
}

func (o *UpdateMemberGroupInterfaceIDUsingPUTInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/policy/memberGroup/interfaceId][%d] updateMemberGroupInterfaceIdUsingPUTInternalServerError", 500)
}

func (o *UpdateMemberGroupInterfaceIDUsingPUTInternalServerError) String() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/policy/memberGroup/interfaceId][%d] updateMemberGroupInterfaceIdUsingPUTInternalServerError", 500)
}

func (o *UpdateMemberGroupInterfaceIDUsingPUTInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
