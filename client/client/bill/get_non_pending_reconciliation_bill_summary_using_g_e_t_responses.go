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

// GetNonPendingReconciliationBillSummaryUsingGETReader is a Reader for the GetNonPendingReconciliationBillSummaryUsingGET structure.
type GetNonPendingReconciliationBillSummaryUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNonPendingReconciliationBillSummaryUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNonPendingReconciliationBillSummaryUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetNonPendingReconciliationBillSummaryUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetNonPendingReconciliationBillSummaryUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetNonPendingReconciliationBillSummaryUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetNonPendingReconciliationBillSummaryUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetNonPendingReconciliationBillSummaryUsingGETInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /compassintegrationapp/bill/nonPendingReconciliationSummary] getNonPendingReconciliationBillSummaryUsingGET", response, response.Code())
	}
}

// NewGetNonPendingReconciliationBillSummaryUsingGETOK creates a GetNonPendingReconciliationBillSummaryUsingGETOK with default headers values
func NewGetNonPendingReconciliationBillSummaryUsingGETOK() *GetNonPendingReconciliationBillSummaryUsingGETOK {
	return &GetNonPendingReconciliationBillSummaryUsingGETOK{}
}

/*
GetNonPendingReconciliationBillSummaryUsingGETOK describes a response with status code 200, with default header values.

Success
*/
type GetNonPendingReconciliationBillSummaryUsingGETOK struct {
	Payload *models.ResponseWrapperListBillSummaryDTO
}

// IsSuccess returns true when this get non pending reconciliation bill summary using g e t o k response has a 2xx status code
func (o *GetNonPendingReconciliationBillSummaryUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get non pending reconciliation bill summary using g e t o k response has a 3xx status code
func (o *GetNonPendingReconciliationBillSummaryUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get non pending reconciliation bill summary using g e t o k response has a 4xx status code
func (o *GetNonPendingReconciliationBillSummaryUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get non pending reconciliation bill summary using g e t o k response has a 5xx status code
func (o *GetNonPendingReconciliationBillSummaryUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get non pending reconciliation bill summary using g e t o k response a status code equal to that given
func (o *GetNonPendingReconciliationBillSummaryUsingGETOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get non pending reconciliation bill summary using g e t o k response
func (o *GetNonPendingReconciliationBillSummaryUsingGETOK) Code() int {
	return 200
}

func (o *GetNonPendingReconciliationBillSummaryUsingGETOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /compassintegrationapp/bill/nonPendingReconciliationSummary][%d] getNonPendingReconciliationBillSummaryUsingGETOK %s", 200, payload)
}

func (o *GetNonPendingReconciliationBillSummaryUsingGETOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /compassintegrationapp/bill/nonPendingReconciliationSummary][%d] getNonPendingReconciliationBillSummaryUsingGETOK %s", 200, payload)
}

func (o *GetNonPendingReconciliationBillSummaryUsingGETOK) GetPayload() *models.ResponseWrapperListBillSummaryDTO {
	return o.Payload
}

func (o *GetNonPendingReconciliationBillSummaryUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperListBillSummaryDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNonPendingReconciliationBillSummaryUsingGETBadRequest creates a GetNonPendingReconciliationBillSummaryUsingGETBadRequest with default headers values
func NewGetNonPendingReconciliationBillSummaryUsingGETBadRequest() *GetNonPendingReconciliationBillSummaryUsingGETBadRequest {
	return &GetNonPendingReconciliationBillSummaryUsingGETBadRequest{}
}

/*
GetNonPendingReconciliationBillSummaryUsingGETBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetNonPendingReconciliationBillSummaryUsingGETBadRequest struct {
}

// IsSuccess returns true when this get non pending reconciliation bill summary using g e t bad request response has a 2xx status code
func (o *GetNonPendingReconciliationBillSummaryUsingGETBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get non pending reconciliation bill summary using g e t bad request response has a 3xx status code
func (o *GetNonPendingReconciliationBillSummaryUsingGETBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get non pending reconciliation bill summary using g e t bad request response has a 4xx status code
func (o *GetNonPendingReconciliationBillSummaryUsingGETBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get non pending reconciliation bill summary using g e t bad request response has a 5xx status code
func (o *GetNonPendingReconciliationBillSummaryUsingGETBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get non pending reconciliation bill summary using g e t bad request response a status code equal to that given
func (o *GetNonPendingReconciliationBillSummaryUsingGETBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get non pending reconciliation bill summary using g e t bad request response
func (o *GetNonPendingReconciliationBillSummaryUsingGETBadRequest) Code() int {
	return 400
}

func (o *GetNonPendingReconciliationBillSummaryUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /compassintegrationapp/bill/nonPendingReconciliationSummary][%d] getNonPendingReconciliationBillSummaryUsingGETBadRequest", 400)
}

func (o *GetNonPendingReconciliationBillSummaryUsingGETBadRequest) String() string {
	return fmt.Sprintf("[GET /compassintegrationapp/bill/nonPendingReconciliationSummary][%d] getNonPendingReconciliationBillSummaryUsingGETBadRequest", 400)
}

func (o *GetNonPendingReconciliationBillSummaryUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetNonPendingReconciliationBillSummaryUsingGETUnauthorized creates a GetNonPendingReconciliationBillSummaryUsingGETUnauthorized with default headers values
func NewGetNonPendingReconciliationBillSummaryUsingGETUnauthorized() *GetNonPendingReconciliationBillSummaryUsingGETUnauthorized {
	return &GetNonPendingReconciliationBillSummaryUsingGETUnauthorized{}
}

/*
GetNonPendingReconciliationBillSummaryUsingGETUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetNonPendingReconciliationBillSummaryUsingGETUnauthorized struct {
}

// IsSuccess returns true when this get non pending reconciliation bill summary using g e t unauthorized response has a 2xx status code
func (o *GetNonPendingReconciliationBillSummaryUsingGETUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get non pending reconciliation bill summary using g e t unauthorized response has a 3xx status code
func (o *GetNonPendingReconciliationBillSummaryUsingGETUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get non pending reconciliation bill summary using g e t unauthorized response has a 4xx status code
func (o *GetNonPendingReconciliationBillSummaryUsingGETUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get non pending reconciliation bill summary using g e t unauthorized response has a 5xx status code
func (o *GetNonPendingReconciliationBillSummaryUsingGETUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get non pending reconciliation bill summary using g e t unauthorized response a status code equal to that given
func (o *GetNonPendingReconciliationBillSummaryUsingGETUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get non pending reconciliation bill summary using g e t unauthorized response
func (o *GetNonPendingReconciliationBillSummaryUsingGETUnauthorized) Code() int {
	return 401
}

func (o *GetNonPendingReconciliationBillSummaryUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /compassintegrationapp/bill/nonPendingReconciliationSummary][%d] getNonPendingReconciliationBillSummaryUsingGETUnauthorized", 401)
}

func (o *GetNonPendingReconciliationBillSummaryUsingGETUnauthorized) String() string {
	return fmt.Sprintf("[GET /compassintegrationapp/bill/nonPendingReconciliationSummary][%d] getNonPendingReconciliationBillSummaryUsingGETUnauthorized", 401)
}

func (o *GetNonPendingReconciliationBillSummaryUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetNonPendingReconciliationBillSummaryUsingGETForbidden creates a GetNonPendingReconciliationBillSummaryUsingGETForbidden with default headers values
func NewGetNonPendingReconciliationBillSummaryUsingGETForbidden() *GetNonPendingReconciliationBillSummaryUsingGETForbidden {
	return &GetNonPendingReconciliationBillSummaryUsingGETForbidden{}
}

/*
GetNonPendingReconciliationBillSummaryUsingGETForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetNonPendingReconciliationBillSummaryUsingGETForbidden struct {
}

// IsSuccess returns true when this get non pending reconciliation bill summary using g e t forbidden response has a 2xx status code
func (o *GetNonPendingReconciliationBillSummaryUsingGETForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get non pending reconciliation bill summary using g e t forbidden response has a 3xx status code
func (o *GetNonPendingReconciliationBillSummaryUsingGETForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get non pending reconciliation bill summary using g e t forbidden response has a 4xx status code
func (o *GetNonPendingReconciliationBillSummaryUsingGETForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get non pending reconciliation bill summary using g e t forbidden response has a 5xx status code
func (o *GetNonPendingReconciliationBillSummaryUsingGETForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get non pending reconciliation bill summary using g e t forbidden response a status code equal to that given
func (o *GetNonPendingReconciliationBillSummaryUsingGETForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get non pending reconciliation bill summary using g e t forbidden response
func (o *GetNonPendingReconciliationBillSummaryUsingGETForbidden) Code() int {
	return 403
}

func (o *GetNonPendingReconciliationBillSummaryUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /compassintegrationapp/bill/nonPendingReconciliationSummary][%d] getNonPendingReconciliationBillSummaryUsingGETForbidden", 403)
}

func (o *GetNonPendingReconciliationBillSummaryUsingGETForbidden) String() string {
	return fmt.Sprintf("[GET /compassintegrationapp/bill/nonPendingReconciliationSummary][%d] getNonPendingReconciliationBillSummaryUsingGETForbidden", 403)
}

func (o *GetNonPendingReconciliationBillSummaryUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetNonPendingReconciliationBillSummaryUsingGETNotFound creates a GetNonPendingReconciliationBillSummaryUsingGETNotFound with default headers values
func NewGetNonPendingReconciliationBillSummaryUsingGETNotFound() *GetNonPendingReconciliationBillSummaryUsingGETNotFound {
	return &GetNonPendingReconciliationBillSummaryUsingGETNotFound{}
}

/*
GetNonPendingReconciliationBillSummaryUsingGETNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetNonPendingReconciliationBillSummaryUsingGETNotFound struct {
}

// IsSuccess returns true when this get non pending reconciliation bill summary using g e t not found response has a 2xx status code
func (o *GetNonPendingReconciliationBillSummaryUsingGETNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get non pending reconciliation bill summary using g e t not found response has a 3xx status code
func (o *GetNonPendingReconciliationBillSummaryUsingGETNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get non pending reconciliation bill summary using g e t not found response has a 4xx status code
func (o *GetNonPendingReconciliationBillSummaryUsingGETNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get non pending reconciliation bill summary using g e t not found response has a 5xx status code
func (o *GetNonPendingReconciliationBillSummaryUsingGETNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get non pending reconciliation bill summary using g e t not found response a status code equal to that given
func (o *GetNonPendingReconciliationBillSummaryUsingGETNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get non pending reconciliation bill summary using g e t not found response
func (o *GetNonPendingReconciliationBillSummaryUsingGETNotFound) Code() int {
	return 404
}

func (o *GetNonPendingReconciliationBillSummaryUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /compassintegrationapp/bill/nonPendingReconciliationSummary][%d] getNonPendingReconciliationBillSummaryUsingGETNotFound", 404)
}

func (o *GetNonPendingReconciliationBillSummaryUsingGETNotFound) String() string {
	return fmt.Sprintf("[GET /compassintegrationapp/bill/nonPendingReconciliationSummary][%d] getNonPendingReconciliationBillSummaryUsingGETNotFound", 404)
}

func (o *GetNonPendingReconciliationBillSummaryUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetNonPendingReconciliationBillSummaryUsingGETInternalServerError creates a GetNonPendingReconciliationBillSummaryUsingGETInternalServerError with default headers values
func NewGetNonPendingReconciliationBillSummaryUsingGETInternalServerError() *GetNonPendingReconciliationBillSummaryUsingGETInternalServerError {
	return &GetNonPendingReconciliationBillSummaryUsingGETInternalServerError{}
}

/*
GetNonPendingReconciliationBillSummaryUsingGETInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type GetNonPendingReconciliationBillSummaryUsingGETInternalServerError struct {
}

// IsSuccess returns true when this get non pending reconciliation bill summary using g e t internal server error response has a 2xx status code
func (o *GetNonPendingReconciliationBillSummaryUsingGETInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get non pending reconciliation bill summary using g e t internal server error response has a 3xx status code
func (o *GetNonPendingReconciliationBillSummaryUsingGETInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get non pending reconciliation bill summary using g e t internal server error response has a 4xx status code
func (o *GetNonPendingReconciliationBillSummaryUsingGETInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get non pending reconciliation bill summary using g e t internal server error response has a 5xx status code
func (o *GetNonPendingReconciliationBillSummaryUsingGETInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get non pending reconciliation bill summary using g e t internal server error response a status code equal to that given
func (o *GetNonPendingReconciliationBillSummaryUsingGETInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get non pending reconciliation bill summary using g e t internal server error response
func (o *GetNonPendingReconciliationBillSummaryUsingGETInternalServerError) Code() int {
	return 500
}

func (o *GetNonPendingReconciliationBillSummaryUsingGETInternalServerError) Error() string {
	return fmt.Sprintf("[GET /compassintegrationapp/bill/nonPendingReconciliationSummary][%d] getNonPendingReconciliationBillSummaryUsingGETInternalServerError", 500)
}

func (o *GetNonPendingReconciliationBillSummaryUsingGETInternalServerError) String() string {
	return fmt.Sprintf("[GET /compassintegrationapp/bill/nonPendingReconciliationSummary][%d] getNonPendingReconciliationBillSummaryUsingGETInternalServerError", 500)
}

func (o *GetNonPendingReconciliationBillSummaryUsingGETInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
