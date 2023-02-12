// Code generated by go-swagger; DO NOT EDIT.

package task

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/alswl/go-toodledo/pkg/models"
)

// GetTasksGetPhpReader is a Reader for the GetTasksGetPhp structure.
type GetTasksGetPhpReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTasksGetPhpReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTasksGetPhpOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetTasksGetPhpUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetTasksGetPhpTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetTasksGetPhpServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTasksGetPhpOK creates a GetTasksGetPhpOK with default headers values
func NewGetTasksGetPhpOK() *GetTasksGetPhpOK {
	return &GetTasksGetPhpOK{}
}

/*
GetTasksGetPhpOK describes a response with status code 200, with default header values.

ok
*/
type GetTasksGetPhpOK struct {
	Payload models.PaginatedAPIResponse
}

func (o *GetTasksGetPhpOK) Error() string {
	return fmt.Sprintf("[GET /tasks/get.php][%d] getTasksGetPhpOK  %+v", 200, o.Payload)
}
func (o *GetTasksGetPhpOK) GetPayload() models.PaginatedAPIResponse {
	return o.Payload
}

func (o *GetTasksGetPhpOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTasksGetPhpUnauthorized creates a GetTasksGetPhpUnauthorized with default headers values
func NewGetTasksGetPhpUnauthorized() *GetTasksGetPhpUnauthorized {
	return &GetTasksGetPhpUnauthorized{}
}

/*
GetTasksGetPhpUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetTasksGetPhpUnauthorized struct {
	Payload *models.Error
}

func (o *GetTasksGetPhpUnauthorized) Error() string {
	return fmt.Sprintf("[GET /tasks/get.php][%d] getTasksGetPhpUnauthorized  %+v", 401, o.Payload)
}
func (o *GetTasksGetPhpUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetTasksGetPhpUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTasksGetPhpTooManyRequests creates a GetTasksGetPhpTooManyRequests with default headers values
func NewGetTasksGetPhpTooManyRequests() *GetTasksGetPhpTooManyRequests {
	return &GetTasksGetPhpTooManyRequests{}
}

/*
GetTasksGetPhpTooManyRequests describes a response with status code 429, with default header values.

TooManyRequests
*/
type GetTasksGetPhpTooManyRequests struct {
	Payload *models.Error
}

func (o *GetTasksGetPhpTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /tasks/get.php][%d] getTasksGetPhpTooManyRequests  %+v", 429, o.Payload)
}
func (o *GetTasksGetPhpTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetTasksGetPhpTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTasksGetPhpServiceUnavailable creates a GetTasksGetPhpServiceUnavailable with default headers values
func NewGetTasksGetPhpServiceUnavailable() *GetTasksGetPhpServiceUnavailable {
	return &GetTasksGetPhpServiceUnavailable{}
}

/*
GetTasksGetPhpServiceUnavailable describes a response with status code 503, with default header values.

ServiceUnavailable
*/
type GetTasksGetPhpServiceUnavailable struct {
	Payload *models.Error
}

func (o *GetTasksGetPhpServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /tasks/get.php][%d] getTasksGetPhpServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetTasksGetPhpServiceUnavailable) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetTasksGetPhpServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}