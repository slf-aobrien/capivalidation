// Code generated by go-swagger; DO NOT EDIT.

package bill

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

// GetSuspenseAmtForBillsUsingGETReader is a Reader for the GetSuspenseAmtForBillsUsingGET structure.
type GetSuspenseAmtForBillsUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSuspenseAmtForBillsUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSuspenseAmtForBillsUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetSuspenseAmtForBillsUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetSuspenseAmtForBillsUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetSuspenseAmtForBillsUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetSuspenseAmtForBillsUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSuspenseAmtForBillsUsingGETInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /compassintegrationapp/bill/suspenseAmtForBills] getSuspenseAmtForBillsUsingGET", response, response.Code())
	}
}

// NewGetSuspenseAmtForBillsUsingGETOK creates a GetSuspenseAmtForBillsUsingGETOK with default headers values
func NewGetSuspenseAmtForBillsUsingGETOK() *GetSuspenseAmtForBillsUsingGETOK {
	return &GetSuspenseAmtForBillsUsingGETOK{}
}

/*
GetSuspenseAmtForBillsUsingGETOK describes a response with status code 200, with default header values.

Success
*/
type GetSuspenseAmtForBillsUsingGETOK struct {
	Payload *models.ResponseWrapperListBillDTO
}

// IsSuccess returns true when this get suspense amt for bills using g e t o k response has a 2xx status code
func (o *GetSuspenseAmtForBillsUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get suspense amt for bills using g e t o k response has a 3xx status code
func (o *GetSuspenseAmtForBillsUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get suspense amt for bills using g e t o k response has a 4xx status code
func (o *GetSuspenseAmtForBillsUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get suspense amt for bills using g e t o k response has a 5xx status code
func (o *GetSuspenseAmtForBillsUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get suspense amt for bills using g e t o k response a status code equal to that given
func (o *GetSuspenseAmtForBillsUsingGETOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get suspense amt for bills using g e t o k response
func (o *GetSuspenseAmtForBillsUsingGETOK) Code() int {
	return 200
}

func (o *GetSuspenseAmtForBillsUsingGETOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /compassintegrationapp/bill/suspenseAmtForBills][%d] getSuspenseAmtForBillsUsingGETOK %s", 200, payload)
}

func (o *GetSuspenseAmtForBillsUsingGETOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /compassintegrationapp/bill/suspenseAmtForBills][%d] getSuspenseAmtForBillsUsingGETOK %s", 200, payload)
}

func (o *GetSuspenseAmtForBillsUsingGETOK) GetPayload() *models.ResponseWrapperListBillDTO {
	return o.Payload
}

func (o *GetSuspenseAmtForBillsUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperListBillDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSuspenseAmtForBillsUsingGETBadRequest creates a GetSuspenseAmtForBillsUsingGETBadRequest with default headers values
func NewGetSuspenseAmtForBillsUsingGETBadRequest() *GetSuspenseAmtForBillsUsingGETBadRequest {
	return &GetSuspenseAmtForBillsUsingGETBadRequest{}
}

/*
GetSuspenseAmtForBillsUsingGETBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetSuspenseAmtForBillsUsingGETBadRequest struct {
}

// IsSuccess returns true when this get suspense amt for bills using g e t bad request response has a 2xx status code
func (o *GetSuspenseAmtForBillsUsingGETBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get suspense amt for bills using g e t bad request response has a 3xx status code
func (o *GetSuspenseAmtForBillsUsingGETBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get suspense amt for bills using g e t bad request response has a 4xx status code
func (o *GetSuspenseAmtForBillsUsingGETBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get suspense amt for bills using g e t bad request response has a 5xx status code
func (o *GetSuspenseAmtForBillsUsingGETBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get suspense amt for bills using g e t bad request response a status code equal to that given
func (o *GetSuspenseAmtForBillsUsingGETBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get suspense amt for bills using g e t bad request response
func (o *GetSuspenseAmtForBillsUsingGETBadRequest) Code() int {
	return 400
}

func (o *GetSuspenseAmtForBillsUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /compassintegrationapp/bill/suspenseAmtForBills][%d] getSuspenseAmtForBillsUsingGETBadRequest", 400)
}

func (o *GetSuspenseAmtForBillsUsingGETBadRequest) String() string {
	return fmt.Sprintf("[GET /compassintegrationapp/bill/suspenseAmtForBills][%d] getSuspenseAmtForBillsUsingGETBadRequest", 400)
}

func (o *GetSuspenseAmtForBillsUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSuspenseAmtForBillsUsingGETUnauthorized creates a GetSuspenseAmtForBillsUsingGETUnauthorized with default headers values
func NewGetSuspenseAmtForBillsUsingGETUnauthorized() *GetSuspenseAmtForBillsUsingGETUnauthorized {
	return &GetSuspenseAmtForBillsUsingGETUnauthorized{}
}

/*
GetSuspenseAmtForBillsUsingGETUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetSuspenseAmtForBillsUsingGETUnauthorized struct {
}

// IsSuccess returns true when this get suspense amt for bills using g e t unauthorized response has a 2xx status code
func (o *GetSuspenseAmtForBillsUsingGETUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get suspense amt for bills using g e t unauthorized response has a 3xx status code
func (o *GetSuspenseAmtForBillsUsingGETUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get suspense amt for bills using g e t unauthorized response has a 4xx status code
func (o *GetSuspenseAmtForBillsUsingGETUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get suspense amt for bills using g e t unauthorized response has a 5xx status code
func (o *GetSuspenseAmtForBillsUsingGETUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get suspense amt for bills using g e t unauthorized response a status code equal to that given
func (o *GetSuspenseAmtForBillsUsingGETUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get suspense amt for bills using g e t unauthorized response
func (o *GetSuspenseAmtForBillsUsingGETUnauthorized) Code() int {
	return 401
}

func (o *GetSuspenseAmtForBillsUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /compassintegrationapp/bill/suspenseAmtForBills][%d] getSuspenseAmtForBillsUsingGETUnauthorized", 401)
}

func (o *GetSuspenseAmtForBillsUsingGETUnauthorized) String() string {
	return fmt.Sprintf("[GET /compassintegrationapp/bill/suspenseAmtForBills][%d] getSuspenseAmtForBillsUsingGETUnauthorized", 401)
}

func (o *GetSuspenseAmtForBillsUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSuspenseAmtForBillsUsingGETForbidden creates a GetSuspenseAmtForBillsUsingGETForbidden with default headers values
func NewGetSuspenseAmtForBillsUsingGETForbidden() *GetSuspenseAmtForBillsUsingGETForbidden {
	return &GetSuspenseAmtForBillsUsingGETForbidden{}
}

/*
GetSuspenseAmtForBillsUsingGETForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetSuspenseAmtForBillsUsingGETForbidden struct {
}

// IsSuccess returns true when this get suspense amt for bills using g e t forbidden response has a 2xx status code
func (o *GetSuspenseAmtForBillsUsingGETForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get suspense amt for bills using g e t forbidden response has a 3xx status code
func (o *GetSuspenseAmtForBillsUsingGETForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get suspense amt for bills using g e t forbidden response has a 4xx status code
func (o *GetSuspenseAmtForBillsUsingGETForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get suspense amt for bills using g e t forbidden response has a 5xx status code
func (o *GetSuspenseAmtForBillsUsingGETForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get suspense amt for bills using g e t forbidden response a status code equal to that given
func (o *GetSuspenseAmtForBillsUsingGETForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get suspense amt for bills using g e t forbidden response
func (o *GetSuspenseAmtForBillsUsingGETForbidden) Code() int {
	return 403
}

func (o *GetSuspenseAmtForBillsUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /compassintegrationapp/bill/suspenseAmtForBills][%d] getSuspenseAmtForBillsUsingGETForbidden", 403)
}

func (o *GetSuspenseAmtForBillsUsingGETForbidden) String() string {
	return fmt.Sprintf("[GET /compassintegrationapp/bill/suspenseAmtForBills][%d] getSuspenseAmtForBillsUsingGETForbidden", 403)
}

func (o *GetSuspenseAmtForBillsUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSuspenseAmtForBillsUsingGETNotFound creates a GetSuspenseAmtForBillsUsingGETNotFound with default headers values
func NewGetSuspenseAmtForBillsUsingGETNotFound() *GetSuspenseAmtForBillsUsingGETNotFound {
	return &GetSuspenseAmtForBillsUsingGETNotFound{}
}

/*
GetSuspenseAmtForBillsUsingGETNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetSuspenseAmtForBillsUsingGETNotFound struct {
}

// IsSuccess returns true when this get suspense amt for bills using g e t not found response has a 2xx status code
func (o *GetSuspenseAmtForBillsUsingGETNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get suspense amt for bills using g e t not found response has a 3xx status code
func (o *GetSuspenseAmtForBillsUsingGETNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get suspense amt for bills using g e t not found response has a 4xx status code
func (o *GetSuspenseAmtForBillsUsingGETNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get suspense amt for bills using g e t not found response has a 5xx status code
func (o *GetSuspenseAmtForBillsUsingGETNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get suspense amt for bills using g e t not found response a status code equal to that given
func (o *GetSuspenseAmtForBillsUsingGETNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get suspense amt for bills using g e t not found response
func (o *GetSuspenseAmtForBillsUsingGETNotFound) Code() int {
	return 404
}

func (o *GetSuspenseAmtForBillsUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /compassintegrationapp/bill/suspenseAmtForBills][%d] getSuspenseAmtForBillsUsingGETNotFound", 404)
}

func (o *GetSuspenseAmtForBillsUsingGETNotFound) String() string {
	return fmt.Sprintf("[GET /compassintegrationapp/bill/suspenseAmtForBills][%d] getSuspenseAmtForBillsUsingGETNotFound", 404)
}

func (o *GetSuspenseAmtForBillsUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSuspenseAmtForBillsUsingGETInternalServerError creates a GetSuspenseAmtForBillsUsingGETInternalServerError with default headers values
func NewGetSuspenseAmtForBillsUsingGETInternalServerError() *GetSuspenseAmtForBillsUsingGETInternalServerError {
	return &GetSuspenseAmtForBillsUsingGETInternalServerError{}
}

/*
GetSuspenseAmtForBillsUsingGETInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type GetSuspenseAmtForBillsUsingGETInternalServerError struct {
}

// IsSuccess returns true when this get suspense amt for bills using g e t internal server error response has a 2xx status code
func (o *GetSuspenseAmtForBillsUsingGETInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get suspense amt for bills using g e t internal server error response has a 3xx status code
func (o *GetSuspenseAmtForBillsUsingGETInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get suspense amt for bills using g e t internal server error response has a 4xx status code
func (o *GetSuspenseAmtForBillsUsingGETInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get suspense amt for bills using g e t internal server error response has a 5xx status code
func (o *GetSuspenseAmtForBillsUsingGETInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get suspense amt for bills using g e t internal server error response a status code equal to that given
func (o *GetSuspenseAmtForBillsUsingGETInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get suspense amt for bills using g e t internal server error response
func (o *GetSuspenseAmtForBillsUsingGETInternalServerError) Code() int {
	return 500
}

func (o *GetSuspenseAmtForBillsUsingGETInternalServerError) Error() string {
	return fmt.Sprintf("[GET /compassintegrationapp/bill/suspenseAmtForBills][%d] getSuspenseAmtForBillsUsingGETInternalServerError", 500)
}

func (o *GetSuspenseAmtForBillsUsingGETInternalServerError) String() string {
	return fmt.Sprintf("[GET /compassintegrationapp/bill/suspenseAmtForBills][%d] getSuspenseAmtForBillsUsingGETInternalServerError", 500)
}

func (o *GetSuspenseAmtForBillsUsingGETInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
