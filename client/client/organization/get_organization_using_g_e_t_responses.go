// Code generated by go-swagger; DO NOT EDIT.

package organization

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

// GetOrganizationUsingGETReader is a Reader for the GetOrganizationUsingGET structure.
type GetOrganizationUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganizationUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrganizationUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOrganizationUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetOrganizationUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOrganizationUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOrganizationUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOrganizationUsingGETInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /compassintegrationapp/organization] getOrganizationUsingGET", response, response.Code())
	}
}

// NewGetOrganizationUsingGETOK creates a GetOrganizationUsingGETOK with default headers values
func NewGetOrganizationUsingGETOK() *GetOrganizationUsingGETOK {
	return &GetOrganizationUsingGETOK{}
}

/*
GetOrganizationUsingGETOK describes a response with status code 200, with default header values.

Success
*/
type GetOrganizationUsingGETOK struct {
	Payload *models.ResponseWrapperOrganizationContactDTO
}

// IsSuccess returns true when this get organization using g e t o k response has a 2xx status code
func (o *GetOrganizationUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get organization using g e t o k response has a 3xx status code
func (o *GetOrganizationUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization using g e t o k response has a 4xx status code
func (o *GetOrganizationUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get organization using g e t o k response has a 5xx status code
func (o *GetOrganizationUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization using g e t o k response a status code equal to that given
func (o *GetOrganizationUsingGETOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get organization using g e t o k response
func (o *GetOrganizationUsingGETOK) Code() int {
	return 200
}

func (o *GetOrganizationUsingGETOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /compassintegrationapp/organization][%d] getOrganizationUsingGETOK %s", 200, payload)
}

func (o *GetOrganizationUsingGETOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /compassintegrationapp/organization][%d] getOrganizationUsingGETOK %s", 200, payload)
}

func (o *GetOrganizationUsingGETOK) GetPayload() *models.ResponseWrapperOrganizationContactDTO {
	return o.Payload
}

func (o *GetOrganizationUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperOrganizationContactDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationUsingGETBadRequest creates a GetOrganizationUsingGETBadRequest with default headers values
func NewGetOrganizationUsingGETBadRequest() *GetOrganizationUsingGETBadRequest {
	return &GetOrganizationUsingGETBadRequest{}
}

/*
GetOrganizationUsingGETBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetOrganizationUsingGETBadRequest struct {
}

// IsSuccess returns true when this get organization using g e t bad request response has a 2xx status code
func (o *GetOrganizationUsingGETBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get organization using g e t bad request response has a 3xx status code
func (o *GetOrganizationUsingGETBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization using g e t bad request response has a 4xx status code
func (o *GetOrganizationUsingGETBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get organization using g e t bad request response has a 5xx status code
func (o *GetOrganizationUsingGETBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization using g e t bad request response a status code equal to that given
func (o *GetOrganizationUsingGETBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get organization using g e t bad request response
func (o *GetOrganizationUsingGETBadRequest) Code() int {
	return 400
}

func (o *GetOrganizationUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /compassintegrationapp/organization][%d] getOrganizationUsingGETBadRequest", 400)
}

func (o *GetOrganizationUsingGETBadRequest) String() string {
	return fmt.Sprintf("[GET /compassintegrationapp/organization][%d] getOrganizationUsingGETBadRequest", 400)
}

func (o *GetOrganizationUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetOrganizationUsingGETUnauthorized creates a GetOrganizationUsingGETUnauthorized with default headers values
func NewGetOrganizationUsingGETUnauthorized() *GetOrganizationUsingGETUnauthorized {
	return &GetOrganizationUsingGETUnauthorized{}
}

/*
GetOrganizationUsingGETUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetOrganizationUsingGETUnauthorized struct {
}

// IsSuccess returns true when this get organization using g e t unauthorized response has a 2xx status code
func (o *GetOrganizationUsingGETUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get organization using g e t unauthorized response has a 3xx status code
func (o *GetOrganizationUsingGETUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization using g e t unauthorized response has a 4xx status code
func (o *GetOrganizationUsingGETUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get organization using g e t unauthorized response has a 5xx status code
func (o *GetOrganizationUsingGETUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization using g e t unauthorized response a status code equal to that given
func (o *GetOrganizationUsingGETUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get organization using g e t unauthorized response
func (o *GetOrganizationUsingGETUnauthorized) Code() int {
	return 401
}

func (o *GetOrganizationUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /compassintegrationapp/organization][%d] getOrganizationUsingGETUnauthorized", 401)
}

func (o *GetOrganizationUsingGETUnauthorized) String() string {
	return fmt.Sprintf("[GET /compassintegrationapp/organization][%d] getOrganizationUsingGETUnauthorized", 401)
}

func (o *GetOrganizationUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetOrganizationUsingGETForbidden creates a GetOrganizationUsingGETForbidden with default headers values
func NewGetOrganizationUsingGETForbidden() *GetOrganizationUsingGETForbidden {
	return &GetOrganizationUsingGETForbidden{}
}

/*
GetOrganizationUsingGETForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetOrganizationUsingGETForbidden struct {
}

// IsSuccess returns true when this get organization using g e t forbidden response has a 2xx status code
func (o *GetOrganizationUsingGETForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get organization using g e t forbidden response has a 3xx status code
func (o *GetOrganizationUsingGETForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization using g e t forbidden response has a 4xx status code
func (o *GetOrganizationUsingGETForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get organization using g e t forbidden response has a 5xx status code
func (o *GetOrganizationUsingGETForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization using g e t forbidden response a status code equal to that given
func (o *GetOrganizationUsingGETForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get organization using g e t forbidden response
func (o *GetOrganizationUsingGETForbidden) Code() int {
	return 403
}

func (o *GetOrganizationUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /compassintegrationapp/organization][%d] getOrganizationUsingGETForbidden", 403)
}

func (o *GetOrganizationUsingGETForbidden) String() string {
	return fmt.Sprintf("[GET /compassintegrationapp/organization][%d] getOrganizationUsingGETForbidden", 403)
}

func (o *GetOrganizationUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetOrganizationUsingGETNotFound creates a GetOrganizationUsingGETNotFound with default headers values
func NewGetOrganizationUsingGETNotFound() *GetOrganizationUsingGETNotFound {
	return &GetOrganizationUsingGETNotFound{}
}

/*
GetOrganizationUsingGETNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetOrganizationUsingGETNotFound struct {
}

// IsSuccess returns true when this get organization using g e t not found response has a 2xx status code
func (o *GetOrganizationUsingGETNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get organization using g e t not found response has a 3xx status code
func (o *GetOrganizationUsingGETNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization using g e t not found response has a 4xx status code
func (o *GetOrganizationUsingGETNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get organization using g e t not found response has a 5xx status code
func (o *GetOrganizationUsingGETNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization using g e t not found response a status code equal to that given
func (o *GetOrganizationUsingGETNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get organization using g e t not found response
func (o *GetOrganizationUsingGETNotFound) Code() int {
	return 404
}

func (o *GetOrganizationUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /compassintegrationapp/organization][%d] getOrganizationUsingGETNotFound", 404)
}

func (o *GetOrganizationUsingGETNotFound) String() string {
	return fmt.Sprintf("[GET /compassintegrationapp/organization][%d] getOrganizationUsingGETNotFound", 404)
}

func (o *GetOrganizationUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetOrganizationUsingGETInternalServerError creates a GetOrganizationUsingGETInternalServerError with default headers values
func NewGetOrganizationUsingGETInternalServerError() *GetOrganizationUsingGETInternalServerError {
	return &GetOrganizationUsingGETInternalServerError{}
}

/*
GetOrganizationUsingGETInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type GetOrganizationUsingGETInternalServerError struct {
}

// IsSuccess returns true when this get organization using g e t internal server error response has a 2xx status code
func (o *GetOrganizationUsingGETInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get organization using g e t internal server error response has a 3xx status code
func (o *GetOrganizationUsingGETInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization using g e t internal server error response has a 4xx status code
func (o *GetOrganizationUsingGETInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get organization using g e t internal server error response has a 5xx status code
func (o *GetOrganizationUsingGETInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get organization using g e t internal server error response a status code equal to that given
func (o *GetOrganizationUsingGETInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get organization using g e t internal server error response
func (o *GetOrganizationUsingGETInternalServerError) Code() int {
	return 500
}

func (o *GetOrganizationUsingGETInternalServerError) Error() string {
	return fmt.Sprintf("[GET /compassintegrationapp/organization][%d] getOrganizationUsingGETInternalServerError", 500)
}

func (o *GetOrganizationUsingGETInternalServerError) String() string {
	return fmt.Sprintf("[GET /compassintegrationapp/organization][%d] getOrganizationUsingGETInternalServerError", 500)
}

func (o *GetOrganizationUsingGETInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
