// Code generated by go-swagger; DO NOT EDIT.

package error

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// HandleErrorUsingGETReader is a Reader for the HandleErrorUsingGET structure.
type HandleErrorUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HandleErrorUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHandleErrorUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewHandleErrorUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewHandleErrorUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewHandleErrorUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /compassintegrationapp/error] handleErrorUsingGET", response, response.Code())
	}
}

// NewHandleErrorUsingGETOK creates a HandleErrorUsingGETOK with default headers values
func NewHandleErrorUsingGETOK() *HandleErrorUsingGETOK {
	return &HandleErrorUsingGETOK{}
}

/*
HandleErrorUsingGETOK describes a response with status code 200, with default header values.

OK
*/
type HandleErrorUsingGETOK struct {
}

// IsSuccess returns true when this handle error using g e t o k response has a 2xx status code
func (o *HandleErrorUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this handle error using g e t o k response has a 3xx status code
func (o *HandleErrorUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this handle error using g e t o k response has a 4xx status code
func (o *HandleErrorUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this handle error using g e t o k response has a 5xx status code
func (o *HandleErrorUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this handle error using g e t o k response a status code equal to that given
func (o *HandleErrorUsingGETOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the handle error using g e t o k response
func (o *HandleErrorUsingGETOK) Code() int {
	return 200
}

func (o *HandleErrorUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /compassintegrationapp/error][%d] handleErrorUsingGETOK", 200)
}

func (o *HandleErrorUsingGETOK) String() string {
	return fmt.Sprintf("[GET /compassintegrationapp/error][%d] handleErrorUsingGETOK", 200)
}

func (o *HandleErrorUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewHandleErrorUsingGETUnauthorized creates a HandleErrorUsingGETUnauthorized with default headers values
func NewHandleErrorUsingGETUnauthorized() *HandleErrorUsingGETUnauthorized {
	return &HandleErrorUsingGETUnauthorized{}
}

/*
HandleErrorUsingGETUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type HandleErrorUsingGETUnauthorized struct {
}

// IsSuccess returns true when this handle error using g e t unauthorized response has a 2xx status code
func (o *HandleErrorUsingGETUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this handle error using g e t unauthorized response has a 3xx status code
func (o *HandleErrorUsingGETUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this handle error using g e t unauthorized response has a 4xx status code
func (o *HandleErrorUsingGETUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this handle error using g e t unauthorized response has a 5xx status code
func (o *HandleErrorUsingGETUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this handle error using g e t unauthorized response a status code equal to that given
func (o *HandleErrorUsingGETUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the handle error using g e t unauthorized response
func (o *HandleErrorUsingGETUnauthorized) Code() int {
	return 401
}

func (o *HandleErrorUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /compassintegrationapp/error][%d] handleErrorUsingGETUnauthorized", 401)
}

func (o *HandleErrorUsingGETUnauthorized) String() string {
	return fmt.Sprintf("[GET /compassintegrationapp/error][%d] handleErrorUsingGETUnauthorized", 401)
}

func (o *HandleErrorUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewHandleErrorUsingGETForbidden creates a HandleErrorUsingGETForbidden with default headers values
func NewHandleErrorUsingGETForbidden() *HandleErrorUsingGETForbidden {
	return &HandleErrorUsingGETForbidden{}
}

/*
HandleErrorUsingGETForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type HandleErrorUsingGETForbidden struct {
}

// IsSuccess returns true when this handle error using g e t forbidden response has a 2xx status code
func (o *HandleErrorUsingGETForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this handle error using g e t forbidden response has a 3xx status code
func (o *HandleErrorUsingGETForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this handle error using g e t forbidden response has a 4xx status code
func (o *HandleErrorUsingGETForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this handle error using g e t forbidden response has a 5xx status code
func (o *HandleErrorUsingGETForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this handle error using g e t forbidden response a status code equal to that given
func (o *HandleErrorUsingGETForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the handle error using g e t forbidden response
func (o *HandleErrorUsingGETForbidden) Code() int {
	return 403
}

func (o *HandleErrorUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /compassintegrationapp/error][%d] handleErrorUsingGETForbidden", 403)
}

func (o *HandleErrorUsingGETForbidden) String() string {
	return fmt.Sprintf("[GET /compassintegrationapp/error][%d] handleErrorUsingGETForbidden", 403)
}

func (o *HandleErrorUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewHandleErrorUsingGETNotFound creates a HandleErrorUsingGETNotFound with default headers values
func NewHandleErrorUsingGETNotFound() *HandleErrorUsingGETNotFound {
	return &HandleErrorUsingGETNotFound{}
}

/*
HandleErrorUsingGETNotFound describes a response with status code 404, with default header values.

Not Found
*/
type HandleErrorUsingGETNotFound struct {
}

// IsSuccess returns true when this handle error using g e t not found response has a 2xx status code
func (o *HandleErrorUsingGETNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this handle error using g e t not found response has a 3xx status code
func (o *HandleErrorUsingGETNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this handle error using g e t not found response has a 4xx status code
func (o *HandleErrorUsingGETNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this handle error using g e t not found response has a 5xx status code
func (o *HandleErrorUsingGETNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this handle error using g e t not found response a status code equal to that given
func (o *HandleErrorUsingGETNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the handle error using g e t not found response
func (o *HandleErrorUsingGETNotFound) Code() int {
	return 404
}

func (o *HandleErrorUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /compassintegrationapp/error][%d] handleErrorUsingGETNotFound", 404)
}

func (o *HandleErrorUsingGETNotFound) String() string {
	return fmt.Sprintf("[GET /compassintegrationapp/error][%d] handleErrorUsingGETNotFound", 404)
}

func (o *HandleErrorUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
