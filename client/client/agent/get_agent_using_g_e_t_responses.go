// Code generated by go-swagger; DO NOT EDIT.

package agent

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

// GetAgentUsingGETReader is a Reader for the GetAgentUsingGET structure.
type GetAgentUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAgentUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAgentUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAgentUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAgentUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAgentUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAgentUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAgentUsingGETInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /compassintegrationapp/agent] getAgentUsingGET", response, response.Code())
	}
}

// NewGetAgentUsingGETOK creates a GetAgentUsingGETOK with default headers values
func NewGetAgentUsingGETOK() *GetAgentUsingGETOK {
	return &GetAgentUsingGETOK{}
}

/*
GetAgentUsingGETOK describes a response with status code 200, with default header values.

Success
*/
type GetAgentUsingGETOK struct {
	Payload *models.ResponseWrapperListAgentDTO
}

// IsSuccess returns true when this get agent using g e t o k response has a 2xx status code
func (o *GetAgentUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get agent using g e t o k response has a 3xx status code
func (o *GetAgentUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get agent using g e t o k response has a 4xx status code
func (o *GetAgentUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get agent using g e t o k response has a 5xx status code
func (o *GetAgentUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get agent using g e t o k response a status code equal to that given
func (o *GetAgentUsingGETOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get agent using g e t o k response
func (o *GetAgentUsingGETOK) Code() int {
	return 200
}

func (o *GetAgentUsingGETOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /compassintegrationapp/agent][%d] getAgentUsingGETOK %s", 200, payload)
}

func (o *GetAgentUsingGETOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /compassintegrationapp/agent][%d] getAgentUsingGETOK %s", 200, payload)
}

func (o *GetAgentUsingGETOK) GetPayload() *models.ResponseWrapperListAgentDTO {
	return o.Payload
}

func (o *GetAgentUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperListAgentDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAgentUsingGETBadRequest creates a GetAgentUsingGETBadRequest with default headers values
func NewGetAgentUsingGETBadRequest() *GetAgentUsingGETBadRequest {
	return &GetAgentUsingGETBadRequest{}
}

/*
GetAgentUsingGETBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetAgentUsingGETBadRequest struct {
}

// IsSuccess returns true when this get agent using g e t bad request response has a 2xx status code
func (o *GetAgentUsingGETBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get agent using g e t bad request response has a 3xx status code
func (o *GetAgentUsingGETBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get agent using g e t bad request response has a 4xx status code
func (o *GetAgentUsingGETBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get agent using g e t bad request response has a 5xx status code
func (o *GetAgentUsingGETBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get agent using g e t bad request response a status code equal to that given
func (o *GetAgentUsingGETBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get agent using g e t bad request response
func (o *GetAgentUsingGETBadRequest) Code() int {
	return 400
}

func (o *GetAgentUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /compassintegrationapp/agent][%d] getAgentUsingGETBadRequest", 400)
}

func (o *GetAgentUsingGETBadRequest) String() string {
	return fmt.Sprintf("[GET /compassintegrationapp/agent][%d] getAgentUsingGETBadRequest", 400)
}

func (o *GetAgentUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAgentUsingGETUnauthorized creates a GetAgentUsingGETUnauthorized with default headers values
func NewGetAgentUsingGETUnauthorized() *GetAgentUsingGETUnauthorized {
	return &GetAgentUsingGETUnauthorized{}
}

/*
GetAgentUsingGETUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetAgentUsingGETUnauthorized struct {
}

// IsSuccess returns true when this get agent using g e t unauthorized response has a 2xx status code
func (o *GetAgentUsingGETUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get agent using g e t unauthorized response has a 3xx status code
func (o *GetAgentUsingGETUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get agent using g e t unauthorized response has a 4xx status code
func (o *GetAgentUsingGETUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get agent using g e t unauthorized response has a 5xx status code
func (o *GetAgentUsingGETUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get agent using g e t unauthorized response a status code equal to that given
func (o *GetAgentUsingGETUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get agent using g e t unauthorized response
func (o *GetAgentUsingGETUnauthorized) Code() int {
	return 401
}

func (o *GetAgentUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /compassintegrationapp/agent][%d] getAgentUsingGETUnauthorized", 401)
}

func (o *GetAgentUsingGETUnauthorized) String() string {
	return fmt.Sprintf("[GET /compassintegrationapp/agent][%d] getAgentUsingGETUnauthorized", 401)
}

func (o *GetAgentUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAgentUsingGETForbidden creates a GetAgentUsingGETForbidden with default headers values
func NewGetAgentUsingGETForbidden() *GetAgentUsingGETForbidden {
	return &GetAgentUsingGETForbidden{}
}

/*
GetAgentUsingGETForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAgentUsingGETForbidden struct {
}

// IsSuccess returns true when this get agent using g e t forbidden response has a 2xx status code
func (o *GetAgentUsingGETForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get agent using g e t forbidden response has a 3xx status code
func (o *GetAgentUsingGETForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get agent using g e t forbidden response has a 4xx status code
func (o *GetAgentUsingGETForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get agent using g e t forbidden response has a 5xx status code
func (o *GetAgentUsingGETForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get agent using g e t forbidden response a status code equal to that given
func (o *GetAgentUsingGETForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get agent using g e t forbidden response
func (o *GetAgentUsingGETForbidden) Code() int {
	return 403
}

func (o *GetAgentUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /compassintegrationapp/agent][%d] getAgentUsingGETForbidden", 403)
}

func (o *GetAgentUsingGETForbidden) String() string {
	return fmt.Sprintf("[GET /compassintegrationapp/agent][%d] getAgentUsingGETForbidden", 403)
}

func (o *GetAgentUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAgentUsingGETNotFound creates a GetAgentUsingGETNotFound with default headers values
func NewGetAgentUsingGETNotFound() *GetAgentUsingGETNotFound {
	return &GetAgentUsingGETNotFound{}
}

/*
GetAgentUsingGETNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetAgentUsingGETNotFound struct {
}

// IsSuccess returns true when this get agent using g e t not found response has a 2xx status code
func (o *GetAgentUsingGETNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get agent using g e t not found response has a 3xx status code
func (o *GetAgentUsingGETNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get agent using g e t not found response has a 4xx status code
func (o *GetAgentUsingGETNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get agent using g e t not found response has a 5xx status code
func (o *GetAgentUsingGETNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get agent using g e t not found response a status code equal to that given
func (o *GetAgentUsingGETNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get agent using g e t not found response
func (o *GetAgentUsingGETNotFound) Code() int {
	return 404
}

func (o *GetAgentUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /compassintegrationapp/agent][%d] getAgentUsingGETNotFound", 404)
}

func (o *GetAgentUsingGETNotFound) String() string {
	return fmt.Sprintf("[GET /compassintegrationapp/agent][%d] getAgentUsingGETNotFound", 404)
}

func (o *GetAgentUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAgentUsingGETInternalServerError creates a GetAgentUsingGETInternalServerError with default headers values
func NewGetAgentUsingGETInternalServerError() *GetAgentUsingGETInternalServerError {
	return &GetAgentUsingGETInternalServerError{}
}

/*
GetAgentUsingGETInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type GetAgentUsingGETInternalServerError struct {
}

// IsSuccess returns true when this get agent using g e t internal server error response has a 2xx status code
func (o *GetAgentUsingGETInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get agent using g e t internal server error response has a 3xx status code
func (o *GetAgentUsingGETInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get agent using g e t internal server error response has a 4xx status code
func (o *GetAgentUsingGETInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get agent using g e t internal server error response has a 5xx status code
func (o *GetAgentUsingGETInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get agent using g e t internal server error response a status code equal to that given
func (o *GetAgentUsingGETInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get agent using g e t internal server error response
func (o *GetAgentUsingGETInternalServerError) Code() int {
	return 500
}

func (o *GetAgentUsingGETInternalServerError) Error() string {
	return fmt.Sprintf("[GET /compassintegrationapp/agent][%d] getAgentUsingGETInternalServerError", 500)
}

func (o *GetAgentUsingGETInternalServerError) String() string {
	return fmt.Sprintf("[GET /compassintegrationapp/agent][%d] getAgentUsingGETInternalServerError", 500)
}

func (o *GetAgentUsingGETInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
