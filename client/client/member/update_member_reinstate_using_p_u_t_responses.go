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

// UpdateMemberReinstateUsingPUTReader is a Reader for the UpdateMemberReinstateUsingPUT structure.
type UpdateMemberReinstateUsingPUTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateMemberReinstateUsingPUTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateMemberReinstateUsingPUTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewUpdateMemberReinstateUsingPUTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateMemberReinstateUsingPUTBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateMemberReinstateUsingPUTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateMemberReinstateUsingPUTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateMemberReinstateUsingPUTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateMemberReinstateUsingPUTInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /compassintegrationapp/member/update/reinstate] updateMemberReinstateUsingPUT", response, response.Code())
	}
}

// NewUpdateMemberReinstateUsingPUTOK creates a UpdateMemberReinstateUsingPUTOK with default headers values
func NewUpdateMemberReinstateUsingPUTOK() *UpdateMemberReinstateUsingPUTOK {
	return &UpdateMemberReinstateUsingPUTOK{}
}

/*
UpdateMemberReinstateUsingPUTOK describes a response with status code 200, with default header values.

Success
*/
type UpdateMemberReinstateUsingPUTOK struct {
	Payload *models.ResponseWrapperBoolean
}

// IsSuccess returns true when this update member reinstate using p u t o k response has a 2xx status code
func (o *UpdateMemberReinstateUsingPUTOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update member reinstate using p u t o k response has a 3xx status code
func (o *UpdateMemberReinstateUsingPUTOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update member reinstate using p u t o k response has a 4xx status code
func (o *UpdateMemberReinstateUsingPUTOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update member reinstate using p u t o k response has a 5xx status code
func (o *UpdateMemberReinstateUsingPUTOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update member reinstate using p u t o k response a status code equal to that given
func (o *UpdateMemberReinstateUsingPUTOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update member reinstate using p u t o k response
func (o *UpdateMemberReinstateUsingPUTOK) Code() int {
	return 200
}

func (o *UpdateMemberReinstateUsingPUTOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /compassintegrationapp/member/update/reinstate][%d] updateMemberReinstateUsingPUTOK %s", 200, payload)
}

func (o *UpdateMemberReinstateUsingPUTOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /compassintegrationapp/member/update/reinstate][%d] updateMemberReinstateUsingPUTOK %s", 200, payload)
}

func (o *UpdateMemberReinstateUsingPUTOK) GetPayload() *models.ResponseWrapperBoolean {
	return o.Payload
}

func (o *UpdateMemberReinstateUsingPUTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperBoolean)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMemberReinstateUsingPUTCreated creates a UpdateMemberReinstateUsingPUTCreated with default headers values
func NewUpdateMemberReinstateUsingPUTCreated() *UpdateMemberReinstateUsingPUTCreated {
	return &UpdateMemberReinstateUsingPUTCreated{}
}

/*
UpdateMemberReinstateUsingPUTCreated describes a response with status code 201, with default header values.

Created
*/
type UpdateMemberReinstateUsingPUTCreated struct {
}

// IsSuccess returns true when this update member reinstate using p u t created response has a 2xx status code
func (o *UpdateMemberReinstateUsingPUTCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update member reinstate using p u t created response has a 3xx status code
func (o *UpdateMemberReinstateUsingPUTCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update member reinstate using p u t created response has a 4xx status code
func (o *UpdateMemberReinstateUsingPUTCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this update member reinstate using p u t created response has a 5xx status code
func (o *UpdateMemberReinstateUsingPUTCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this update member reinstate using p u t created response a status code equal to that given
func (o *UpdateMemberReinstateUsingPUTCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the update member reinstate using p u t created response
func (o *UpdateMemberReinstateUsingPUTCreated) Code() int {
	return 201
}

func (o *UpdateMemberReinstateUsingPUTCreated) Error() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/member/update/reinstate][%d] updateMemberReinstateUsingPUTCreated", 201)
}

func (o *UpdateMemberReinstateUsingPUTCreated) String() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/member/update/reinstate][%d] updateMemberReinstateUsingPUTCreated", 201)
}

func (o *UpdateMemberReinstateUsingPUTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateMemberReinstateUsingPUTBadRequest creates a UpdateMemberReinstateUsingPUTBadRequest with default headers values
func NewUpdateMemberReinstateUsingPUTBadRequest() *UpdateMemberReinstateUsingPUTBadRequest {
	return &UpdateMemberReinstateUsingPUTBadRequest{}
}

/*
UpdateMemberReinstateUsingPUTBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UpdateMemberReinstateUsingPUTBadRequest struct {
}

// IsSuccess returns true when this update member reinstate using p u t bad request response has a 2xx status code
func (o *UpdateMemberReinstateUsingPUTBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update member reinstate using p u t bad request response has a 3xx status code
func (o *UpdateMemberReinstateUsingPUTBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update member reinstate using p u t bad request response has a 4xx status code
func (o *UpdateMemberReinstateUsingPUTBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update member reinstate using p u t bad request response has a 5xx status code
func (o *UpdateMemberReinstateUsingPUTBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update member reinstate using p u t bad request response a status code equal to that given
func (o *UpdateMemberReinstateUsingPUTBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update member reinstate using p u t bad request response
func (o *UpdateMemberReinstateUsingPUTBadRequest) Code() int {
	return 400
}

func (o *UpdateMemberReinstateUsingPUTBadRequest) Error() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/member/update/reinstate][%d] updateMemberReinstateUsingPUTBadRequest", 400)
}

func (o *UpdateMemberReinstateUsingPUTBadRequest) String() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/member/update/reinstate][%d] updateMemberReinstateUsingPUTBadRequest", 400)
}

func (o *UpdateMemberReinstateUsingPUTBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateMemberReinstateUsingPUTUnauthorized creates a UpdateMemberReinstateUsingPUTUnauthorized with default headers values
func NewUpdateMemberReinstateUsingPUTUnauthorized() *UpdateMemberReinstateUsingPUTUnauthorized {
	return &UpdateMemberReinstateUsingPUTUnauthorized{}
}

/*
UpdateMemberReinstateUsingPUTUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UpdateMemberReinstateUsingPUTUnauthorized struct {
}

// IsSuccess returns true when this update member reinstate using p u t unauthorized response has a 2xx status code
func (o *UpdateMemberReinstateUsingPUTUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update member reinstate using p u t unauthorized response has a 3xx status code
func (o *UpdateMemberReinstateUsingPUTUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update member reinstate using p u t unauthorized response has a 4xx status code
func (o *UpdateMemberReinstateUsingPUTUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update member reinstate using p u t unauthorized response has a 5xx status code
func (o *UpdateMemberReinstateUsingPUTUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update member reinstate using p u t unauthorized response a status code equal to that given
func (o *UpdateMemberReinstateUsingPUTUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the update member reinstate using p u t unauthorized response
func (o *UpdateMemberReinstateUsingPUTUnauthorized) Code() int {
	return 401
}

func (o *UpdateMemberReinstateUsingPUTUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/member/update/reinstate][%d] updateMemberReinstateUsingPUTUnauthorized", 401)
}

func (o *UpdateMemberReinstateUsingPUTUnauthorized) String() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/member/update/reinstate][%d] updateMemberReinstateUsingPUTUnauthorized", 401)
}

func (o *UpdateMemberReinstateUsingPUTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateMemberReinstateUsingPUTForbidden creates a UpdateMemberReinstateUsingPUTForbidden with default headers values
func NewUpdateMemberReinstateUsingPUTForbidden() *UpdateMemberReinstateUsingPUTForbidden {
	return &UpdateMemberReinstateUsingPUTForbidden{}
}

/*
UpdateMemberReinstateUsingPUTForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateMemberReinstateUsingPUTForbidden struct {
}

// IsSuccess returns true when this update member reinstate using p u t forbidden response has a 2xx status code
func (o *UpdateMemberReinstateUsingPUTForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update member reinstate using p u t forbidden response has a 3xx status code
func (o *UpdateMemberReinstateUsingPUTForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update member reinstate using p u t forbidden response has a 4xx status code
func (o *UpdateMemberReinstateUsingPUTForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update member reinstate using p u t forbidden response has a 5xx status code
func (o *UpdateMemberReinstateUsingPUTForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update member reinstate using p u t forbidden response a status code equal to that given
func (o *UpdateMemberReinstateUsingPUTForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update member reinstate using p u t forbidden response
func (o *UpdateMemberReinstateUsingPUTForbidden) Code() int {
	return 403
}

func (o *UpdateMemberReinstateUsingPUTForbidden) Error() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/member/update/reinstate][%d] updateMemberReinstateUsingPUTForbidden", 403)
}

func (o *UpdateMemberReinstateUsingPUTForbidden) String() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/member/update/reinstate][%d] updateMemberReinstateUsingPUTForbidden", 403)
}

func (o *UpdateMemberReinstateUsingPUTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateMemberReinstateUsingPUTNotFound creates a UpdateMemberReinstateUsingPUTNotFound with default headers values
func NewUpdateMemberReinstateUsingPUTNotFound() *UpdateMemberReinstateUsingPUTNotFound {
	return &UpdateMemberReinstateUsingPUTNotFound{}
}

/*
UpdateMemberReinstateUsingPUTNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdateMemberReinstateUsingPUTNotFound struct {
}

// IsSuccess returns true when this update member reinstate using p u t not found response has a 2xx status code
func (o *UpdateMemberReinstateUsingPUTNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update member reinstate using p u t not found response has a 3xx status code
func (o *UpdateMemberReinstateUsingPUTNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update member reinstate using p u t not found response has a 4xx status code
func (o *UpdateMemberReinstateUsingPUTNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update member reinstate using p u t not found response has a 5xx status code
func (o *UpdateMemberReinstateUsingPUTNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update member reinstate using p u t not found response a status code equal to that given
func (o *UpdateMemberReinstateUsingPUTNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update member reinstate using p u t not found response
func (o *UpdateMemberReinstateUsingPUTNotFound) Code() int {
	return 404
}

func (o *UpdateMemberReinstateUsingPUTNotFound) Error() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/member/update/reinstate][%d] updateMemberReinstateUsingPUTNotFound", 404)
}

func (o *UpdateMemberReinstateUsingPUTNotFound) String() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/member/update/reinstate][%d] updateMemberReinstateUsingPUTNotFound", 404)
}

func (o *UpdateMemberReinstateUsingPUTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateMemberReinstateUsingPUTInternalServerError creates a UpdateMemberReinstateUsingPUTInternalServerError with default headers values
func NewUpdateMemberReinstateUsingPUTInternalServerError() *UpdateMemberReinstateUsingPUTInternalServerError {
	return &UpdateMemberReinstateUsingPUTInternalServerError{}
}

/*
UpdateMemberReinstateUsingPUTInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type UpdateMemberReinstateUsingPUTInternalServerError struct {
}

// IsSuccess returns true when this update member reinstate using p u t internal server error response has a 2xx status code
func (o *UpdateMemberReinstateUsingPUTInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update member reinstate using p u t internal server error response has a 3xx status code
func (o *UpdateMemberReinstateUsingPUTInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update member reinstate using p u t internal server error response has a 4xx status code
func (o *UpdateMemberReinstateUsingPUTInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update member reinstate using p u t internal server error response has a 5xx status code
func (o *UpdateMemberReinstateUsingPUTInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update member reinstate using p u t internal server error response a status code equal to that given
func (o *UpdateMemberReinstateUsingPUTInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update member reinstate using p u t internal server error response
func (o *UpdateMemberReinstateUsingPUTInternalServerError) Code() int {
	return 500
}

func (o *UpdateMemberReinstateUsingPUTInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/member/update/reinstate][%d] updateMemberReinstateUsingPUTInternalServerError", 500)
}

func (o *UpdateMemberReinstateUsingPUTInternalServerError) String() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/member/update/reinstate][%d] updateMemberReinstateUsingPUTInternalServerError", 500)
}

func (o *UpdateMemberReinstateUsingPUTInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
