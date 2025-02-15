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

// GetEDXPolicyAgentInfoUsingGETReader is a Reader for the GetEDXPolicyAgentInfoUsingGET structure.
type GetEDXPolicyAgentInfoUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEDXPolicyAgentInfoUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEDXPolicyAgentInfoUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetEDXPolicyAgentInfoUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetEDXPolicyAgentInfoUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetEDXPolicyAgentInfoUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetEDXPolicyAgentInfoUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetEDXPolicyAgentInfoUsingGETInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /compassintegrationapp/policy/edxPolicyAgent] getEDXPolicyAgentInfoUsingGET", response, response.Code())
	}
}

// NewGetEDXPolicyAgentInfoUsingGETOK creates a GetEDXPolicyAgentInfoUsingGETOK with default headers values
func NewGetEDXPolicyAgentInfoUsingGETOK() *GetEDXPolicyAgentInfoUsingGETOK {
	return &GetEDXPolicyAgentInfoUsingGETOK{}
}

/*
GetEDXPolicyAgentInfoUsingGETOK describes a response with status code 200, with default header values.

Success
*/
type GetEDXPolicyAgentInfoUsingGETOK struct {
	Payload *models.ResponseWrapperPolicyDTO
}

// IsSuccess returns true when this get e d x policy agent info using g e t o k response has a 2xx status code
func (o *GetEDXPolicyAgentInfoUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get e d x policy agent info using g e t o k response has a 3xx status code
func (o *GetEDXPolicyAgentInfoUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get e d x policy agent info using g e t o k response has a 4xx status code
func (o *GetEDXPolicyAgentInfoUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get e d x policy agent info using g e t o k response has a 5xx status code
func (o *GetEDXPolicyAgentInfoUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get e d x policy agent info using g e t o k response a status code equal to that given
func (o *GetEDXPolicyAgentInfoUsingGETOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get e d x policy agent info using g e t o k response
func (o *GetEDXPolicyAgentInfoUsingGETOK) Code() int {
	return 200
}

func (o *GetEDXPolicyAgentInfoUsingGETOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /compassintegrationapp/policy/edxPolicyAgent][%d] getEDXPolicyAgentInfoUsingGETOK %s", 200, payload)
}

func (o *GetEDXPolicyAgentInfoUsingGETOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /compassintegrationapp/policy/edxPolicyAgent][%d] getEDXPolicyAgentInfoUsingGETOK %s", 200, payload)
}

func (o *GetEDXPolicyAgentInfoUsingGETOK) GetPayload() *models.ResponseWrapperPolicyDTO {
	return o.Payload
}

func (o *GetEDXPolicyAgentInfoUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperPolicyDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEDXPolicyAgentInfoUsingGETBadRequest creates a GetEDXPolicyAgentInfoUsingGETBadRequest with default headers values
func NewGetEDXPolicyAgentInfoUsingGETBadRequest() *GetEDXPolicyAgentInfoUsingGETBadRequest {
	return &GetEDXPolicyAgentInfoUsingGETBadRequest{}
}

/*
GetEDXPolicyAgentInfoUsingGETBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetEDXPolicyAgentInfoUsingGETBadRequest struct {
}

// IsSuccess returns true when this get e d x policy agent info using g e t bad request response has a 2xx status code
func (o *GetEDXPolicyAgentInfoUsingGETBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get e d x policy agent info using g e t bad request response has a 3xx status code
func (o *GetEDXPolicyAgentInfoUsingGETBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get e d x policy agent info using g e t bad request response has a 4xx status code
func (o *GetEDXPolicyAgentInfoUsingGETBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get e d x policy agent info using g e t bad request response has a 5xx status code
func (o *GetEDXPolicyAgentInfoUsingGETBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get e d x policy agent info using g e t bad request response a status code equal to that given
func (o *GetEDXPolicyAgentInfoUsingGETBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get e d x policy agent info using g e t bad request response
func (o *GetEDXPolicyAgentInfoUsingGETBadRequest) Code() int {
	return 400
}

func (o *GetEDXPolicyAgentInfoUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /compassintegrationapp/policy/edxPolicyAgent][%d] getEDXPolicyAgentInfoUsingGETBadRequest", 400)
}

func (o *GetEDXPolicyAgentInfoUsingGETBadRequest) String() string {
	return fmt.Sprintf("[GET /compassintegrationapp/policy/edxPolicyAgent][%d] getEDXPolicyAgentInfoUsingGETBadRequest", 400)
}

func (o *GetEDXPolicyAgentInfoUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetEDXPolicyAgentInfoUsingGETUnauthorized creates a GetEDXPolicyAgentInfoUsingGETUnauthorized with default headers values
func NewGetEDXPolicyAgentInfoUsingGETUnauthorized() *GetEDXPolicyAgentInfoUsingGETUnauthorized {
	return &GetEDXPolicyAgentInfoUsingGETUnauthorized{}
}

/*
GetEDXPolicyAgentInfoUsingGETUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetEDXPolicyAgentInfoUsingGETUnauthorized struct {
}

// IsSuccess returns true when this get e d x policy agent info using g e t unauthorized response has a 2xx status code
func (o *GetEDXPolicyAgentInfoUsingGETUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get e d x policy agent info using g e t unauthorized response has a 3xx status code
func (o *GetEDXPolicyAgentInfoUsingGETUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get e d x policy agent info using g e t unauthorized response has a 4xx status code
func (o *GetEDXPolicyAgentInfoUsingGETUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get e d x policy agent info using g e t unauthorized response has a 5xx status code
func (o *GetEDXPolicyAgentInfoUsingGETUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get e d x policy agent info using g e t unauthorized response a status code equal to that given
func (o *GetEDXPolicyAgentInfoUsingGETUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get e d x policy agent info using g e t unauthorized response
func (o *GetEDXPolicyAgentInfoUsingGETUnauthorized) Code() int {
	return 401
}

func (o *GetEDXPolicyAgentInfoUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /compassintegrationapp/policy/edxPolicyAgent][%d] getEDXPolicyAgentInfoUsingGETUnauthorized", 401)
}

func (o *GetEDXPolicyAgentInfoUsingGETUnauthorized) String() string {
	return fmt.Sprintf("[GET /compassintegrationapp/policy/edxPolicyAgent][%d] getEDXPolicyAgentInfoUsingGETUnauthorized", 401)
}

func (o *GetEDXPolicyAgentInfoUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetEDXPolicyAgentInfoUsingGETForbidden creates a GetEDXPolicyAgentInfoUsingGETForbidden with default headers values
func NewGetEDXPolicyAgentInfoUsingGETForbidden() *GetEDXPolicyAgentInfoUsingGETForbidden {
	return &GetEDXPolicyAgentInfoUsingGETForbidden{}
}

/*
GetEDXPolicyAgentInfoUsingGETForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetEDXPolicyAgentInfoUsingGETForbidden struct {
}

// IsSuccess returns true when this get e d x policy agent info using g e t forbidden response has a 2xx status code
func (o *GetEDXPolicyAgentInfoUsingGETForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get e d x policy agent info using g e t forbidden response has a 3xx status code
func (o *GetEDXPolicyAgentInfoUsingGETForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get e d x policy agent info using g e t forbidden response has a 4xx status code
func (o *GetEDXPolicyAgentInfoUsingGETForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get e d x policy agent info using g e t forbidden response has a 5xx status code
func (o *GetEDXPolicyAgentInfoUsingGETForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get e d x policy agent info using g e t forbidden response a status code equal to that given
func (o *GetEDXPolicyAgentInfoUsingGETForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get e d x policy agent info using g e t forbidden response
func (o *GetEDXPolicyAgentInfoUsingGETForbidden) Code() int {
	return 403
}

func (o *GetEDXPolicyAgentInfoUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /compassintegrationapp/policy/edxPolicyAgent][%d] getEDXPolicyAgentInfoUsingGETForbidden", 403)
}

func (o *GetEDXPolicyAgentInfoUsingGETForbidden) String() string {
	return fmt.Sprintf("[GET /compassintegrationapp/policy/edxPolicyAgent][%d] getEDXPolicyAgentInfoUsingGETForbidden", 403)
}

func (o *GetEDXPolicyAgentInfoUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetEDXPolicyAgentInfoUsingGETNotFound creates a GetEDXPolicyAgentInfoUsingGETNotFound with default headers values
func NewGetEDXPolicyAgentInfoUsingGETNotFound() *GetEDXPolicyAgentInfoUsingGETNotFound {
	return &GetEDXPolicyAgentInfoUsingGETNotFound{}
}

/*
GetEDXPolicyAgentInfoUsingGETNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetEDXPolicyAgentInfoUsingGETNotFound struct {
}

// IsSuccess returns true when this get e d x policy agent info using g e t not found response has a 2xx status code
func (o *GetEDXPolicyAgentInfoUsingGETNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get e d x policy agent info using g e t not found response has a 3xx status code
func (o *GetEDXPolicyAgentInfoUsingGETNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get e d x policy agent info using g e t not found response has a 4xx status code
func (o *GetEDXPolicyAgentInfoUsingGETNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get e d x policy agent info using g e t not found response has a 5xx status code
func (o *GetEDXPolicyAgentInfoUsingGETNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get e d x policy agent info using g e t not found response a status code equal to that given
func (o *GetEDXPolicyAgentInfoUsingGETNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get e d x policy agent info using g e t not found response
func (o *GetEDXPolicyAgentInfoUsingGETNotFound) Code() int {
	return 404
}

func (o *GetEDXPolicyAgentInfoUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /compassintegrationapp/policy/edxPolicyAgent][%d] getEDXPolicyAgentInfoUsingGETNotFound", 404)
}

func (o *GetEDXPolicyAgentInfoUsingGETNotFound) String() string {
	return fmt.Sprintf("[GET /compassintegrationapp/policy/edxPolicyAgent][%d] getEDXPolicyAgentInfoUsingGETNotFound", 404)
}

func (o *GetEDXPolicyAgentInfoUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetEDXPolicyAgentInfoUsingGETInternalServerError creates a GetEDXPolicyAgentInfoUsingGETInternalServerError with default headers values
func NewGetEDXPolicyAgentInfoUsingGETInternalServerError() *GetEDXPolicyAgentInfoUsingGETInternalServerError {
	return &GetEDXPolicyAgentInfoUsingGETInternalServerError{}
}

/*
GetEDXPolicyAgentInfoUsingGETInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type GetEDXPolicyAgentInfoUsingGETInternalServerError struct {
}

// IsSuccess returns true when this get e d x policy agent info using g e t internal server error response has a 2xx status code
func (o *GetEDXPolicyAgentInfoUsingGETInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get e d x policy agent info using g e t internal server error response has a 3xx status code
func (o *GetEDXPolicyAgentInfoUsingGETInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get e d x policy agent info using g e t internal server error response has a 4xx status code
func (o *GetEDXPolicyAgentInfoUsingGETInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get e d x policy agent info using g e t internal server error response has a 5xx status code
func (o *GetEDXPolicyAgentInfoUsingGETInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get e d x policy agent info using g e t internal server error response a status code equal to that given
func (o *GetEDXPolicyAgentInfoUsingGETInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get e d x policy agent info using g e t internal server error response
func (o *GetEDXPolicyAgentInfoUsingGETInternalServerError) Code() int {
	return 500
}

func (o *GetEDXPolicyAgentInfoUsingGETInternalServerError) Error() string {
	return fmt.Sprintf("[GET /compassintegrationapp/policy/edxPolicyAgent][%d] getEDXPolicyAgentInfoUsingGETInternalServerError", 500)
}

func (o *GetEDXPolicyAgentInfoUsingGETInternalServerError) String() string {
	return fmt.Sprintf("[GET /compassintegrationapp/policy/edxPolicyAgent][%d] getEDXPolicyAgentInfoUsingGETInternalServerError", 500)
}

func (o *GetEDXPolicyAgentInfoUsingGETInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
