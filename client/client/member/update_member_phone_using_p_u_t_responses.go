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

// UpdateMemberPhoneUsingPUTReader is a Reader for the UpdateMemberPhoneUsingPUT structure.
type UpdateMemberPhoneUsingPUTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateMemberPhoneUsingPUTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateMemberPhoneUsingPUTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewUpdateMemberPhoneUsingPUTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateMemberPhoneUsingPUTBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateMemberPhoneUsingPUTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateMemberPhoneUsingPUTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateMemberPhoneUsingPUTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateMemberPhoneUsingPUTInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /compassintegrationapp/member/update/phone] updateMemberPhoneUsingPUT", response, response.Code())
	}
}

// NewUpdateMemberPhoneUsingPUTOK creates a UpdateMemberPhoneUsingPUTOK with default headers values
func NewUpdateMemberPhoneUsingPUTOK() *UpdateMemberPhoneUsingPUTOK {
	return &UpdateMemberPhoneUsingPUTOK{}
}

/*
UpdateMemberPhoneUsingPUTOK describes a response with status code 200, with default header values.

Success
*/
type UpdateMemberPhoneUsingPUTOK struct {
	Payload *models.ResponseWrapperBoolean
}

// IsSuccess returns true when this update member phone using p u t o k response has a 2xx status code
func (o *UpdateMemberPhoneUsingPUTOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update member phone using p u t o k response has a 3xx status code
func (o *UpdateMemberPhoneUsingPUTOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update member phone using p u t o k response has a 4xx status code
func (o *UpdateMemberPhoneUsingPUTOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update member phone using p u t o k response has a 5xx status code
func (o *UpdateMemberPhoneUsingPUTOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update member phone using p u t o k response a status code equal to that given
func (o *UpdateMemberPhoneUsingPUTOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update member phone using p u t o k response
func (o *UpdateMemberPhoneUsingPUTOK) Code() int {
	return 200
}

func (o *UpdateMemberPhoneUsingPUTOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /compassintegrationapp/member/update/phone][%d] updateMemberPhoneUsingPUTOK %s", 200, payload)
}

func (o *UpdateMemberPhoneUsingPUTOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /compassintegrationapp/member/update/phone][%d] updateMemberPhoneUsingPUTOK %s", 200, payload)
}

func (o *UpdateMemberPhoneUsingPUTOK) GetPayload() *models.ResponseWrapperBoolean {
	return o.Payload
}

func (o *UpdateMemberPhoneUsingPUTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperBoolean)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMemberPhoneUsingPUTCreated creates a UpdateMemberPhoneUsingPUTCreated with default headers values
func NewUpdateMemberPhoneUsingPUTCreated() *UpdateMemberPhoneUsingPUTCreated {
	return &UpdateMemberPhoneUsingPUTCreated{}
}

/*
UpdateMemberPhoneUsingPUTCreated describes a response with status code 201, with default header values.

Created
*/
type UpdateMemberPhoneUsingPUTCreated struct {
}

// IsSuccess returns true when this update member phone using p u t created response has a 2xx status code
func (o *UpdateMemberPhoneUsingPUTCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update member phone using p u t created response has a 3xx status code
func (o *UpdateMemberPhoneUsingPUTCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update member phone using p u t created response has a 4xx status code
func (o *UpdateMemberPhoneUsingPUTCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this update member phone using p u t created response has a 5xx status code
func (o *UpdateMemberPhoneUsingPUTCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this update member phone using p u t created response a status code equal to that given
func (o *UpdateMemberPhoneUsingPUTCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the update member phone using p u t created response
func (o *UpdateMemberPhoneUsingPUTCreated) Code() int {
	return 201
}

func (o *UpdateMemberPhoneUsingPUTCreated) Error() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/member/update/phone][%d] updateMemberPhoneUsingPUTCreated", 201)
}

func (o *UpdateMemberPhoneUsingPUTCreated) String() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/member/update/phone][%d] updateMemberPhoneUsingPUTCreated", 201)
}

func (o *UpdateMemberPhoneUsingPUTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateMemberPhoneUsingPUTBadRequest creates a UpdateMemberPhoneUsingPUTBadRequest with default headers values
func NewUpdateMemberPhoneUsingPUTBadRequest() *UpdateMemberPhoneUsingPUTBadRequest {
	return &UpdateMemberPhoneUsingPUTBadRequest{}
}

/*
UpdateMemberPhoneUsingPUTBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UpdateMemberPhoneUsingPUTBadRequest struct {
}

// IsSuccess returns true when this update member phone using p u t bad request response has a 2xx status code
func (o *UpdateMemberPhoneUsingPUTBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update member phone using p u t bad request response has a 3xx status code
func (o *UpdateMemberPhoneUsingPUTBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update member phone using p u t bad request response has a 4xx status code
func (o *UpdateMemberPhoneUsingPUTBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update member phone using p u t bad request response has a 5xx status code
func (o *UpdateMemberPhoneUsingPUTBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update member phone using p u t bad request response a status code equal to that given
func (o *UpdateMemberPhoneUsingPUTBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update member phone using p u t bad request response
func (o *UpdateMemberPhoneUsingPUTBadRequest) Code() int {
	return 400
}

func (o *UpdateMemberPhoneUsingPUTBadRequest) Error() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/member/update/phone][%d] updateMemberPhoneUsingPUTBadRequest", 400)
}

func (o *UpdateMemberPhoneUsingPUTBadRequest) String() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/member/update/phone][%d] updateMemberPhoneUsingPUTBadRequest", 400)
}

func (o *UpdateMemberPhoneUsingPUTBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateMemberPhoneUsingPUTUnauthorized creates a UpdateMemberPhoneUsingPUTUnauthorized with default headers values
func NewUpdateMemberPhoneUsingPUTUnauthorized() *UpdateMemberPhoneUsingPUTUnauthorized {
	return &UpdateMemberPhoneUsingPUTUnauthorized{}
}

/*
UpdateMemberPhoneUsingPUTUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UpdateMemberPhoneUsingPUTUnauthorized struct {
}

// IsSuccess returns true when this update member phone using p u t unauthorized response has a 2xx status code
func (o *UpdateMemberPhoneUsingPUTUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update member phone using p u t unauthorized response has a 3xx status code
func (o *UpdateMemberPhoneUsingPUTUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update member phone using p u t unauthorized response has a 4xx status code
func (o *UpdateMemberPhoneUsingPUTUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update member phone using p u t unauthorized response has a 5xx status code
func (o *UpdateMemberPhoneUsingPUTUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update member phone using p u t unauthorized response a status code equal to that given
func (o *UpdateMemberPhoneUsingPUTUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the update member phone using p u t unauthorized response
func (o *UpdateMemberPhoneUsingPUTUnauthorized) Code() int {
	return 401
}

func (o *UpdateMemberPhoneUsingPUTUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/member/update/phone][%d] updateMemberPhoneUsingPUTUnauthorized", 401)
}

func (o *UpdateMemberPhoneUsingPUTUnauthorized) String() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/member/update/phone][%d] updateMemberPhoneUsingPUTUnauthorized", 401)
}

func (o *UpdateMemberPhoneUsingPUTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateMemberPhoneUsingPUTForbidden creates a UpdateMemberPhoneUsingPUTForbidden with default headers values
func NewUpdateMemberPhoneUsingPUTForbidden() *UpdateMemberPhoneUsingPUTForbidden {
	return &UpdateMemberPhoneUsingPUTForbidden{}
}

/*
UpdateMemberPhoneUsingPUTForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateMemberPhoneUsingPUTForbidden struct {
}

// IsSuccess returns true when this update member phone using p u t forbidden response has a 2xx status code
func (o *UpdateMemberPhoneUsingPUTForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update member phone using p u t forbidden response has a 3xx status code
func (o *UpdateMemberPhoneUsingPUTForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update member phone using p u t forbidden response has a 4xx status code
func (o *UpdateMemberPhoneUsingPUTForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update member phone using p u t forbidden response has a 5xx status code
func (o *UpdateMemberPhoneUsingPUTForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update member phone using p u t forbidden response a status code equal to that given
func (o *UpdateMemberPhoneUsingPUTForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update member phone using p u t forbidden response
func (o *UpdateMemberPhoneUsingPUTForbidden) Code() int {
	return 403
}

func (o *UpdateMemberPhoneUsingPUTForbidden) Error() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/member/update/phone][%d] updateMemberPhoneUsingPUTForbidden", 403)
}

func (o *UpdateMemberPhoneUsingPUTForbidden) String() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/member/update/phone][%d] updateMemberPhoneUsingPUTForbidden", 403)
}

func (o *UpdateMemberPhoneUsingPUTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateMemberPhoneUsingPUTNotFound creates a UpdateMemberPhoneUsingPUTNotFound with default headers values
func NewUpdateMemberPhoneUsingPUTNotFound() *UpdateMemberPhoneUsingPUTNotFound {
	return &UpdateMemberPhoneUsingPUTNotFound{}
}

/*
UpdateMemberPhoneUsingPUTNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdateMemberPhoneUsingPUTNotFound struct {
}

// IsSuccess returns true when this update member phone using p u t not found response has a 2xx status code
func (o *UpdateMemberPhoneUsingPUTNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update member phone using p u t not found response has a 3xx status code
func (o *UpdateMemberPhoneUsingPUTNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update member phone using p u t not found response has a 4xx status code
func (o *UpdateMemberPhoneUsingPUTNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update member phone using p u t not found response has a 5xx status code
func (o *UpdateMemberPhoneUsingPUTNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update member phone using p u t not found response a status code equal to that given
func (o *UpdateMemberPhoneUsingPUTNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update member phone using p u t not found response
func (o *UpdateMemberPhoneUsingPUTNotFound) Code() int {
	return 404
}

func (o *UpdateMemberPhoneUsingPUTNotFound) Error() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/member/update/phone][%d] updateMemberPhoneUsingPUTNotFound", 404)
}

func (o *UpdateMemberPhoneUsingPUTNotFound) String() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/member/update/phone][%d] updateMemberPhoneUsingPUTNotFound", 404)
}

func (o *UpdateMemberPhoneUsingPUTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateMemberPhoneUsingPUTInternalServerError creates a UpdateMemberPhoneUsingPUTInternalServerError with default headers values
func NewUpdateMemberPhoneUsingPUTInternalServerError() *UpdateMemberPhoneUsingPUTInternalServerError {
	return &UpdateMemberPhoneUsingPUTInternalServerError{}
}

/*
UpdateMemberPhoneUsingPUTInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type UpdateMemberPhoneUsingPUTInternalServerError struct {
}

// IsSuccess returns true when this update member phone using p u t internal server error response has a 2xx status code
func (o *UpdateMemberPhoneUsingPUTInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update member phone using p u t internal server error response has a 3xx status code
func (o *UpdateMemberPhoneUsingPUTInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update member phone using p u t internal server error response has a 4xx status code
func (o *UpdateMemberPhoneUsingPUTInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update member phone using p u t internal server error response has a 5xx status code
func (o *UpdateMemberPhoneUsingPUTInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update member phone using p u t internal server error response a status code equal to that given
func (o *UpdateMemberPhoneUsingPUTInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update member phone using p u t internal server error response
func (o *UpdateMemberPhoneUsingPUTInternalServerError) Code() int {
	return 500
}

func (o *UpdateMemberPhoneUsingPUTInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/member/update/phone][%d] updateMemberPhoneUsingPUTInternalServerError", 500)
}

func (o *UpdateMemberPhoneUsingPUTInternalServerError) String() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/member/update/phone][%d] updateMemberPhoneUsingPUTInternalServerError", 500)
}

func (o *UpdateMemberPhoneUsingPUTInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
