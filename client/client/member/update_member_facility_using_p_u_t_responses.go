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

// UpdateMemberFacilityUsingPUTReader is a Reader for the UpdateMemberFacilityUsingPUT structure.
type UpdateMemberFacilityUsingPUTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateMemberFacilityUsingPUTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateMemberFacilityUsingPUTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewUpdateMemberFacilityUsingPUTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateMemberFacilityUsingPUTBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateMemberFacilityUsingPUTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateMemberFacilityUsingPUTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateMemberFacilityUsingPUTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateMemberFacilityUsingPUTInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /compassintegrationapp/member/update/facility] updateMemberFacilityUsingPUT", response, response.Code())
	}
}

// NewUpdateMemberFacilityUsingPUTOK creates a UpdateMemberFacilityUsingPUTOK with default headers values
func NewUpdateMemberFacilityUsingPUTOK() *UpdateMemberFacilityUsingPUTOK {
	return &UpdateMemberFacilityUsingPUTOK{}
}

/*
UpdateMemberFacilityUsingPUTOK describes a response with status code 200, with default header values.

Success
*/
type UpdateMemberFacilityUsingPUTOK struct {
	Payload *models.ResponseWrapperBoolean
}

// IsSuccess returns true when this update member facility using p u t o k response has a 2xx status code
func (o *UpdateMemberFacilityUsingPUTOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update member facility using p u t o k response has a 3xx status code
func (o *UpdateMemberFacilityUsingPUTOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update member facility using p u t o k response has a 4xx status code
func (o *UpdateMemberFacilityUsingPUTOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update member facility using p u t o k response has a 5xx status code
func (o *UpdateMemberFacilityUsingPUTOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update member facility using p u t o k response a status code equal to that given
func (o *UpdateMemberFacilityUsingPUTOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update member facility using p u t o k response
func (o *UpdateMemberFacilityUsingPUTOK) Code() int {
	return 200
}

func (o *UpdateMemberFacilityUsingPUTOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /compassintegrationapp/member/update/facility][%d] updateMemberFacilityUsingPUTOK %s", 200, payload)
}

func (o *UpdateMemberFacilityUsingPUTOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /compassintegrationapp/member/update/facility][%d] updateMemberFacilityUsingPUTOK %s", 200, payload)
}

func (o *UpdateMemberFacilityUsingPUTOK) GetPayload() *models.ResponseWrapperBoolean {
	return o.Payload
}

func (o *UpdateMemberFacilityUsingPUTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperBoolean)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMemberFacilityUsingPUTCreated creates a UpdateMemberFacilityUsingPUTCreated with default headers values
func NewUpdateMemberFacilityUsingPUTCreated() *UpdateMemberFacilityUsingPUTCreated {
	return &UpdateMemberFacilityUsingPUTCreated{}
}

/*
UpdateMemberFacilityUsingPUTCreated describes a response with status code 201, with default header values.

Created
*/
type UpdateMemberFacilityUsingPUTCreated struct {
}

// IsSuccess returns true when this update member facility using p u t created response has a 2xx status code
func (o *UpdateMemberFacilityUsingPUTCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update member facility using p u t created response has a 3xx status code
func (o *UpdateMemberFacilityUsingPUTCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update member facility using p u t created response has a 4xx status code
func (o *UpdateMemberFacilityUsingPUTCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this update member facility using p u t created response has a 5xx status code
func (o *UpdateMemberFacilityUsingPUTCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this update member facility using p u t created response a status code equal to that given
func (o *UpdateMemberFacilityUsingPUTCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the update member facility using p u t created response
func (o *UpdateMemberFacilityUsingPUTCreated) Code() int {
	return 201
}

func (o *UpdateMemberFacilityUsingPUTCreated) Error() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/member/update/facility][%d] updateMemberFacilityUsingPUTCreated", 201)
}

func (o *UpdateMemberFacilityUsingPUTCreated) String() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/member/update/facility][%d] updateMemberFacilityUsingPUTCreated", 201)
}

func (o *UpdateMemberFacilityUsingPUTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateMemberFacilityUsingPUTBadRequest creates a UpdateMemberFacilityUsingPUTBadRequest with default headers values
func NewUpdateMemberFacilityUsingPUTBadRequest() *UpdateMemberFacilityUsingPUTBadRequest {
	return &UpdateMemberFacilityUsingPUTBadRequest{}
}

/*
UpdateMemberFacilityUsingPUTBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UpdateMemberFacilityUsingPUTBadRequest struct {
}

// IsSuccess returns true when this update member facility using p u t bad request response has a 2xx status code
func (o *UpdateMemberFacilityUsingPUTBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update member facility using p u t bad request response has a 3xx status code
func (o *UpdateMemberFacilityUsingPUTBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update member facility using p u t bad request response has a 4xx status code
func (o *UpdateMemberFacilityUsingPUTBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update member facility using p u t bad request response has a 5xx status code
func (o *UpdateMemberFacilityUsingPUTBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update member facility using p u t bad request response a status code equal to that given
func (o *UpdateMemberFacilityUsingPUTBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update member facility using p u t bad request response
func (o *UpdateMemberFacilityUsingPUTBadRequest) Code() int {
	return 400
}

func (o *UpdateMemberFacilityUsingPUTBadRequest) Error() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/member/update/facility][%d] updateMemberFacilityUsingPUTBadRequest", 400)
}

func (o *UpdateMemberFacilityUsingPUTBadRequest) String() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/member/update/facility][%d] updateMemberFacilityUsingPUTBadRequest", 400)
}

func (o *UpdateMemberFacilityUsingPUTBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateMemberFacilityUsingPUTUnauthorized creates a UpdateMemberFacilityUsingPUTUnauthorized with default headers values
func NewUpdateMemberFacilityUsingPUTUnauthorized() *UpdateMemberFacilityUsingPUTUnauthorized {
	return &UpdateMemberFacilityUsingPUTUnauthorized{}
}

/*
UpdateMemberFacilityUsingPUTUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UpdateMemberFacilityUsingPUTUnauthorized struct {
}

// IsSuccess returns true when this update member facility using p u t unauthorized response has a 2xx status code
func (o *UpdateMemberFacilityUsingPUTUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update member facility using p u t unauthorized response has a 3xx status code
func (o *UpdateMemberFacilityUsingPUTUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update member facility using p u t unauthorized response has a 4xx status code
func (o *UpdateMemberFacilityUsingPUTUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update member facility using p u t unauthorized response has a 5xx status code
func (o *UpdateMemberFacilityUsingPUTUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update member facility using p u t unauthorized response a status code equal to that given
func (o *UpdateMemberFacilityUsingPUTUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the update member facility using p u t unauthorized response
func (o *UpdateMemberFacilityUsingPUTUnauthorized) Code() int {
	return 401
}

func (o *UpdateMemberFacilityUsingPUTUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/member/update/facility][%d] updateMemberFacilityUsingPUTUnauthorized", 401)
}

func (o *UpdateMemberFacilityUsingPUTUnauthorized) String() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/member/update/facility][%d] updateMemberFacilityUsingPUTUnauthorized", 401)
}

func (o *UpdateMemberFacilityUsingPUTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateMemberFacilityUsingPUTForbidden creates a UpdateMemberFacilityUsingPUTForbidden with default headers values
func NewUpdateMemberFacilityUsingPUTForbidden() *UpdateMemberFacilityUsingPUTForbidden {
	return &UpdateMemberFacilityUsingPUTForbidden{}
}

/*
UpdateMemberFacilityUsingPUTForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateMemberFacilityUsingPUTForbidden struct {
}

// IsSuccess returns true when this update member facility using p u t forbidden response has a 2xx status code
func (o *UpdateMemberFacilityUsingPUTForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update member facility using p u t forbidden response has a 3xx status code
func (o *UpdateMemberFacilityUsingPUTForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update member facility using p u t forbidden response has a 4xx status code
func (o *UpdateMemberFacilityUsingPUTForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update member facility using p u t forbidden response has a 5xx status code
func (o *UpdateMemberFacilityUsingPUTForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update member facility using p u t forbidden response a status code equal to that given
func (o *UpdateMemberFacilityUsingPUTForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update member facility using p u t forbidden response
func (o *UpdateMemberFacilityUsingPUTForbidden) Code() int {
	return 403
}

func (o *UpdateMemberFacilityUsingPUTForbidden) Error() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/member/update/facility][%d] updateMemberFacilityUsingPUTForbidden", 403)
}

func (o *UpdateMemberFacilityUsingPUTForbidden) String() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/member/update/facility][%d] updateMemberFacilityUsingPUTForbidden", 403)
}

func (o *UpdateMemberFacilityUsingPUTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateMemberFacilityUsingPUTNotFound creates a UpdateMemberFacilityUsingPUTNotFound with default headers values
func NewUpdateMemberFacilityUsingPUTNotFound() *UpdateMemberFacilityUsingPUTNotFound {
	return &UpdateMemberFacilityUsingPUTNotFound{}
}

/*
UpdateMemberFacilityUsingPUTNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdateMemberFacilityUsingPUTNotFound struct {
}

// IsSuccess returns true when this update member facility using p u t not found response has a 2xx status code
func (o *UpdateMemberFacilityUsingPUTNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update member facility using p u t not found response has a 3xx status code
func (o *UpdateMemberFacilityUsingPUTNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update member facility using p u t not found response has a 4xx status code
func (o *UpdateMemberFacilityUsingPUTNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update member facility using p u t not found response has a 5xx status code
func (o *UpdateMemberFacilityUsingPUTNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update member facility using p u t not found response a status code equal to that given
func (o *UpdateMemberFacilityUsingPUTNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update member facility using p u t not found response
func (o *UpdateMemberFacilityUsingPUTNotFound) Code() int {
	return 404
}

func (o *UpdateMemberFacilityUsingPUTNotFound) Error() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/member/update/facility][%d] updateMemberFacilityUsingPUTNotFound", 404)
}

func (o *UpdateMemberFacilityUsingPUTNotFound) String() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/member/update/facility][%d] updateMemberFacilityUsingPUTNotFound", 404)
}

func (o *UpdateMemberFacilityUsingPUTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateMemberFacilityUsingPUTInternalServerError creates a UpdateMemberFacilityUsingPUTInternalServerError with default headers values
func NewUpdateMemberFacilityUsingPUTInternalServerError() *UpdateMemberFacilityUsingPUTInternalServerError {
	return &UpdateMemberFacilityUsingPUTInternalServerError{}
}

/*
UpdateMemberFacilityUsingPUTInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type UpdateMemberFacilityUsingPUTInternalServerError struct {
}

// IsSuccess returns true when this update member facility using p u t internal server error response has a 2xx status code
func (o *UpdateMemberFacilityUsingPUTInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update member facility using p u t internal server error response has a 3xx status code
func (o *UpdateMemberFacilityUsingPUTInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update member facility using p u t internal server error response has a 4xx status code
func (o *UpdateMemberFacilityUsingPUTInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update member facility using p u t internal server error response has a 5xx status code
func (o *UpdateMemberFacilityUsingPUTInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update member facility using p u t internal server error response a status code equal to that given
func (o *UpdateMemberFacilityUsingPUTInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update member facility using p u t internal server error response
func (o *UpdateMemberFacilityUsingPUTInternalServerError) Code() int {
	return 500
}

func (o *UpdateMemberFacilityUsingPUTInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/member/update/facility][%d] updateMemberFacilityUsingPUTInternalServerError", 500)
}

func (o *UpdateMemberFacilityUsingPUTInternalServerError) String() string {
	return fmt.Sprintf("[PUT /compassintegrationapp/member/update/facility][%d] updateMemberFacilityUsingPUTInternalServerError", 500)
}

func (o *UpdateMemberFacilityUsingPUTInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
