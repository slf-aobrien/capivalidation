// Code generated by go-swagger; DO NOT EDIT.

package cache

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

// GetCacheUsingGET1Reader is a Reader for the GetCacheUsingGET1 structure.
type GetCacheUsingGET1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCacheUsingGET1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCacheUsingGET1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetCacheUsingGET1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetCacheUsingGET1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCacheUsingGET1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetCacheUsingGET1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCacheUsingGET1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /compassintegrationapp/cache/{cacheName}] getCacheUsingGET_1", response, response.Code())
	}
}

// NewGetCacheUsingGET1OK creates a GetCacheUsingGET1OK with default headers values
func NewGetCacheUsingGET1OK() *GetCacheUsingGET1OK {
	return &GetCacheUsingGET1OK{}
}

/*
GetCacheUsingGET1OK describes a response with status code 200, with default header values.

Success
*/
type GetCacheUsingGET1OK struct {
	Payload *models.ResponseWrapperListCache
}

// IsSuccess returns true when this get cache using g e t1 o k response has a 2xx status code
func (o *GetCacheUsingGET1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get cache using g e t1 o k response has a 3xx status code
func (o *GetCacheUsingGET1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cache using g e t1 o k response has a 4xx status code
func (o *GetCacheUsingGET1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get cache using g e t1 o k response has a 5xx status code
func (o *GetCacheUsingGET1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get cache using g e t1 o k response a status code equal to that given
func (o *GetCacheUsingGET1OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get cache using g e t1 o k response
func (o *GetCacheUsingGET1OK) Code() int {
	return 200
}

func (o *GetCacheUsingGET1OK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /compassintegrationapp/cache/{cacheName}][%d] getCacheUsingGET1OK %s", 200, payload)
}

func (o *GetCacheUsingGET1OK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /compassintegrationapp/cache/{cacheName}][%d] getCacheUsingGET1OK %s", 200, payload)
}

func (o *GetCacheUsingGET1OK) GetPayload() *models.ResponseWrapperListCache {
	return o.Payload
}

func (o *GetCacheUsingGET1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperListCache)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCacheUsingGET1BadRequest creates a GetCacheUsingGET1BadRequest with default headers values
func NewGetCacheUsingGET1BadRequest() *GetCacheUsingGET1BadRequest {
	return &GetCacheUsingGET1BadRequest{}
}

/*
GetCacheUsingGET1BadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetCacheUsingGET1BadRequest struct {
}

// IsSuccess returns true when this get cache using g e t1 bad request response has a 2xx status code
func (o *GetCacheUsingGET1BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get cache using g e t1 bad request response has a 3xx status code
func (o *GetCacheUsingGET1BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cache using g e t1 bad request response has a 4xx status code
func (o *GetCacheUsingGET1BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get cache using g e t1 bad request response has a 5xx status code
func (o *GetCacheUsingGET1BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get cache using g e t1 bad request response a status code equal to that given
func (o *GetCacheUsingGET1BadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get cache using g e t1 bad request response
func (o *GetCacheUsingGET1BadRequest) Code() int {
	return 400
}

func (o *GetCacheUsingGET1BadRequest) Error() string {
	return fmt.Sprintf("[GET /compassintegrationapp/cache/{cacheName}][%d] getCacheUsingGET1BadRequest", 400)
}

func (o *GetCacheUsingGET1BadRequest) String() string {
	return fmt.Sprintf("[GET /compassintegrationapp/cache/{cacheName}][%d] getCacheUsingGET1BadRequest", 400)
}

func (o *GetCacheUsingGET1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCacheUsingGET1Unauthorized creates a GetCacheUsingGET1Unauthorized with default headers values
func NewGetCacheUsingGET1Unauthorized() *GetCacheUsingGET1Unauthorized {
	return &GetCacheUsingGET1Unauthorized{}
}

/*
GetCacheUsingGET1Unauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetCacheUsingGET1Unauthorized struct {
}

// IsSuccess returns true when this get cache using g e t1 unauthorized response has a 2xx status code
func (o *GetCacheUsingGET1Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get cache using g e t1 unauthorized response has a 3xx status code
func (o *GetCacheUsingGET1Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cache using g e t1 unauthorized response has a 4xx status code
func (o *GetCacheUsingGET1Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get cache using g e t1 unauthorized response has a 5xx status code
func (o *GetCacheUsingGET1Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get cache using g e t1 unauthorized response a status code equal to that given
func (o *GetCacheUsingGET1Unauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get cache using g e t1 unauthorized response
func (o *GetCacheUsingGET1Unauthorized) Code() int {
	return 401
}

func (o *GetCacheUsingGET1Unauthorized) Error() string {
	return fmt.Sprintf("[GET /compassintegrationapp/cache/{cacheName}][%d] getCacheUsingGET1Unauthorized", 401)
}

func (o *GetCacheUsingGET1Unauthorized) String() string {
	return fmt.Sprintf("[GET /compassintegrationapp/cache/{cacheName}][%d] getCacheUsingGET1Unauthorized", 401)
}

func (o *GetCacheUsingGET1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCacheUsingGET1Forbidden creates a GetCacheUsingGET1Forbidden with default headers values
func NewGetCacheUsingGET1Forbidden() *GetCacheUsingGET1Forbidden {
	return &GetCacheUsingGET1Forbidden{}
}

/*
GetCacheUsingGET1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetCacheUsingGET1Forbidden struct {
}

// IsSuccess returns true when this get cache using g e t1 forbidden response has a 2xx status code
func (o *GetCacheUsingGET1Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get cache using g e t1 forbidden response has a 3xx status code
func (o *GetCacheUsingGET1Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cache using g e t1 forbidden response has a 4xx status code
func (o *GetCacheUsingGET1Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get cache using g e t1 forbidden response has a 5xx status code
func (o *GetCacheUsingGET1Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get cache using g e t1 forbidden response a status code equal to that given
func (o *GetCacheUsingGET1Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get cache using g e t1 forbidden response
func (o *GetCacheUsingGET1Forbidden) Code() int {
	return 403
}

func (o *GetCacheUsingGET1Forbidden) Error() string {
	return fmt.Sprintf("[GET /compassintegrationapp/cache/{cacheName}][%d] getCacheUsingGET1Forbidden", 403)
}

func (o *GetCacheUsingGET1Forbidden) String() string {
	return fmt.Sprintf("[GET /compassintegrationapp/cache/{cacheName}][%d] getCacheUsingGET1Forbidden", 403)
}

func (o *GetCacheUsingGET1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCacheUsingGET1NotFound creates a GetCacheUsingGET1NotFound with default headers values
func NewGetCacheUsingGET1NotFound() *GetCacheUsingGET1NotFound {
	return &GetCacheUsingGET1NotFound{}
}

/*
GetCacheUsingGET1NotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetCacheUsingGET1NotFound struct {
}

// IsSuccess returns true when this get cache using g e t1 not found response has a 2xx status code
func (o *GetCacheUsingGET1NotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get cache using g e t1 not found response has a 3xx status code
func (o *GetCacheUsingGET1NotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cache using g e t1 not found response has a 4xx status code
func (o *GetCacheUsingGET1NotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get cache using g e t1 not found response has a 5xx status code
func (o *GetCacheUsingGET1NotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get cache using g e t1 not found response a status code equal to that given
func (o *GetCacheUsingGET1NotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get cache using g e t1 not found response
func (o *GetCacheUsingGET1NotFound) Code() int {
	return 404
}

func (o *GetCacheUsingGET1NotFound) Error() string {
	return fmt.Sprintf("[GET /compassintegrationapp/cache/{cacheName}][%d] getCacheUsingGET1NotFound", 404)
}

func (o *GetCacheUsingGET1NotFound) String() string {
	return fmt.Sprintf("[GET /compassintegrationapp/cache/{cacheName}][%d] getCacheUsingGET1NotFound", 404)
}

func (o *GetCacheUsingGET1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCacheUsingGET1InternalServerError creates a GetCacheUsingGET1InternalServerError with default headers values
func NewGetCacheUsingGET1InternalServerError() *GetCacheUsingGET1InternalServerError {
	return &GetCacheUsingGET1InternalServerError{}
}

/*
GetCacheUsingGET1InternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type GetCacheUsingGET1InternalServerError struct {
}

// IsSuccess returns true when this get cache using g e t1 internal server error response has a 2xx status code
func (o *GetCacheUsingGET1InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get cache using g e t1 internal server error response has a 3xx status code
func (o *GetCacheUsingGET1InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cache using g e t1 internal server error response has a 4xx status code
func (o *GetCacheUsingGET1InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get cache using g e t1 internal server error response has a 5xx status code
func (o *GetCacheUsingGET1InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get cache using g e t1 internal server error response a status code equal to that given
func (o *GetCacheUsingGET1InternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get cache using g e t1 internal server error response
func (o *GetCacheUsingGET1InternalServerError) Code() int {
	return 500
}

func (o *GetCacheUsingGET1InternalServerError) Error() string {
	return fmt.Sprintf("[GET /compassintegrationapp/cache/{cacheName}][%d] getCacheUsingGET1InternalServerError", 500)
}

func (o *GetCacheUsingGET1InternalServerError) String() string {
	return fmt.Sprintf("[GET /compassintegrationapp/cache/{cacheName}][%d] getCacheUsingGET1InternalServerError", 500)
}

func (o *GetCacheUsingGET1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
