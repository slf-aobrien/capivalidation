// Code generated by go-swagger; DO NOT EDIT.

package agency

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

// GetAgenciesUsingGETReader is a Reader for the GetAgenciesUsingGET structure.
type GetAgenciesUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAgenciesUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAgenciesUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAgenciesUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAgenciesUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAgenciesUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAgenciesUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAgenciesUsingGETInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /compassintegrationapp/agency/list] getAgenciesUsingGET", response, response.Code())
	}
}

// NewGetAgenciesUsingGETOK creates a GetAgenciesUsingGETOK with default headers values
func NewGetAgenciesUsingGETOK() *GetAgenciesUsingGETOK {
	return &GetAgenciesUsingGETOK{}
}

/*
GetAgenciesUsingGETOK describes a response with status code 200, with default header values.

Success
*/
type GetAgenciesUsingGETOK struct {
	Payload *models.ResponseWrapperListAgencyDTO
}

// IsSuccess returns true when this get agencies using g e t o k response has a 2xx status code
func (o *GetAgenciesUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get agencies using g e t o k response has a 3xx status code
func (o *GetAgenciesUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get agencies using g e t o k response has a 4xx status code
func (o *GetAgenciesUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get agencies using g e t o k response has a 5xx status code
func (o *GetAgenciesUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get agencies using g e t o k response a status code equal to that given
func (o *GetAgenciesUsingGETOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get agencies using g e t o k response
func (o *GetAgenciesUsingGETOK) Code() int {
	return 200
}

func (o *GetAgenciesUsingGETOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /compassintegrationapp/agency/list][%d] getAgenciesUsingGETOK %s", 200, payload)
}

func (o *GetAgenciesUsingGETOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /compassintegrationapp/agency/list][%d] getAgenciesUsingGETOK %s", 200, payload)
}

func (o *GetAgenciesUsingGETOK) GetPayload() *models.ResponseWrapperListAgencyDTO {
	return o.Payload
}

func (o *GetAgenciesUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperListAgencyDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAgenciesUsingGETBadRequest creates a GetAgenciesUsingGETBadRequest with default headers values
func NewGetAgenciesUsingGETBadRequest() *GetAgenciesUsingGETBadRequest {
	return &GetAgenciesUsingGETBadRequest{}
}

/*
GetAgenciesUsingGETBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetAgenciesUsingGETBadRequest struct {
}

// IsSuccess returns true when this get agencies using g e t bad request response has a 2xx status code
func (o *GetAgenciesUsingGETBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get agencies using g e t bad request response has a 3xx status code
func (o *GetAgenciesUsingGETBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get agencies using g e t bad request response has a 4xx status code
func (o *GetAgenciesUsingGETBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get agencies using g e t bad request response has a 5xx status code
func (o *GetAgenciesUsingGETBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get agencies using g e t bad request response a status code equal to that given
func (o *GetAgenciesUsingGETBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get agencies using g e t bad request response
func (o *GetAgenciesUsingGETBadRequest) Code() int {
	return 400
}

func (o *GetAgenciesUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /compassintegrationapp/agency/list][%d] getAgenciesUsingGETBadRequest", 400)
}

func (o *GetAgenciesUsingGETBadRequest) String() string {
	return fmt.Sprintf("[GET /compassintegrationapp/agency/list][%d] getAgenciesUsingGETBadRequest", 400)
}

func (o *GetAgenciesUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAgenciesUsingGETUnauthorized creates a GetAgenciesUsingGETUnauthorized with default headers values
func NewGetAgenciesUsingGETUnauthorized() *GetAgenciesUsingGETUnauthorized {
	return &GetAgenciesUsingGETUnauthorized{}
}

/*
GetAgenciesUsingGETUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetAgenciesUsingGETUnauthorized struct {
}

// IsSuccess returns true when this get agencies using g e t unauthorized response has a 2xx status code
func (o *GetAgenciesUsingGETUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get agencies using g e t unauthorized response has a 3xx status code
func (o *GetAgenciesUsingGETUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get agencies using g e t unauthorized response has a 4xx status code
func (o *GetAgenciesUsingGETUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get agencies using g e t unauthorized response has a 5xx status code
func (o *GetAgenciesUsingGETUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get agencies using g e t unauthorized response a status code equal to that given
func (o *GetAgenciesUsingGETUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get agencies using g e t unauthorized response
func (o *GetAgenciesUsingGETUnauthorized) Code() int {
	return 401
}

func (o *GetAgenciesUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /compassintegrationapp/agency/list][%d] getAgenciesUsingGETUnauthorized", 401)
}

func (o *GetAgenciesUsingGETUnauthorized) String() string {
	return fmt.Sprintf("[GET /compassintegrationapp/agency/list][%d] getAgenciesUsingGETUnauthorized", 401)
}

func (o *GetAgenciesUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAgenciesUsingGETForbidden creates a GetAgenciesUsingGETForbidden with default headers values
func NewGetAgenciesUsingGETForbidden() *GetAgenciesUsingGETForbidden {
	return &GetAgenciesUsingGETForbidden{}
}

/*
GetAgenciesUsingGETForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAgenciesUsingGETForbidden struct {
}

// IsSuccess returns true when this get agencies using g e t forbidden response has a 2xx status code
func (o *GetAgenciesUsingGETForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get agencies using g e t forbidden response has a 3xx status code
func (o *GetAgenciesUsingGETForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get agencies using g e t forbidden response has a 4xx status code
func (o *GetAgenciesUsingGETForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get agencies using g e t forbidden response has a 5xx status code
func (o *GetAgenciesUsingGETForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get agencies using g e t forbidden response a status code equal to that given
func (o *GetAgenciesUsingGETForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get agencies using g e t forbidden response
func (o *GetAgenciesUsingGETForbidden) Code() int {
	return 403
}

func (o *GetAgenciesUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /compassintegrationapp/agency/list][%d] getAgenciesUsingGETForbidden", 403)
}

func (o *GetAgenciesUsingGETForbidden) String() string {
	return fmt.Sprintf("[GET /compassintegrationapp/agency/list][%d] getAgenciesUsingGETForbidden", 403)
}

func (o *GetAgenciesUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAgenciesUsingGETNotFound creates a GetAgenciesUsingGETNotFound with default headers values
func NewGetAgenciesUsingGETNotFound() *GetAgenciesUsingGETNotFound {
	return &GetAgenciesUsingGETNotFound{}
}

/*
GetAgenciesUsingGETNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetAgenciesUsingGETNotFound struct {
}

// IsSuccess returns true when this get agencies using g e t not found response has a 2xx status code
func (o *GetAgenciesUsingGETNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get agencies using g e t not found response has a 3xx status code
func (o *GetAgenciesUsingGETNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get agencies using g e t not found response has a 4xx status code
func (o *GetAgenciesUsingGETNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get agencies using g e t not found response has a 5xx status code
func (o *GetAgenciesUsingGETNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get agencies using g e t not found response a status code equal to that given
func (o *GetAgenciesUsingGETNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get agencies using g e t not found response
func (o *GetAgenciesUsingGETNotFound) Code() int {
	return 404
}

func (o *GetAgenciesUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /compassintegrationapp/agency/list][%d] getAgenciesUsingGETNotFound", 404)
}

func (o *GetAgenciesUsingGETNotFound) String() string {
	return fmt.Sprintf("[GET /compassintegrationapp/agency/list][%d] getAgenciesUsingGETNotFound", 404)
}

func (o *GetAgenciesUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAgenciesUsingGETInternalServerError creates a GetAgenciesUsingGETInternalServerError with default headers values
func NewGetAgenciesUsingGETInternalServerError() *GetAgenciesUsingGETInternalServerError {
	return &GetAgenciesUsingGETInternalServerError{}
}

/*
GetAgenciesUsingGETInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type GetAgenciesUsingGETInternalServerError struct {
}

// IsSuccess returns true when this get agencies using g e t internal server error response has a 2xx status code
func (o *GetAgenciesUsingGETInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get agencies using g e t internal server error response has a 3xx status code
func (o *GetAgenciesUsingGETInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get agencies using g e t internal server error response has a 4xx status code
func (o *GetAgenciesUsingGETInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get agencies using g e t internal server error response has a 5xx status code
func (o *GetAgenciesUsingGETInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get agencies using g e t internal server error response a status code equal to that given
func (o *GetAgenciesUsingGETInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get agencies using g e t internal server error response
func (o *GetAgenciesUsingGETInternalServerError) Code() int {
	return 500
}

func (o *GetAgenciesUsingGETInternalServerError) Error() string {
	return fmt.Sprintf("[GET /compassintegrationapp/agency/list][%d] getAgenciesUsingGETInternalServerError", 500)
}

func (o *GetAgenciesUsingGETInternalServerError) String() string {
	return fmt.Sprintf("[GET /compassintegrationapp/agency/list][%d] getAgenciesUsingGETInternalServerError", 500)
}

func (o *GetAgenciesUsingGETInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
