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

// GetTasksDeletedPhpReader is a Reader for the GetTasksDeletedPhp structure.
type GetTasksDeletedPhpReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTasksDeletedPhpReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTasksDeletedPhpOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetTasksDeletedPhpUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetTasksDeletedPhpTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetTasksDeletedPhpServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTasksDeletedPhpOK creates a GetTasksDeletedPhpOK with default headers values
func NewGetTasksDeletedPhpOK() *GetTasksDeletedPhpOK {
	return &GetTasksDeletedPhpOK{}
}

/* GetTasksDeletedPhpOK describes a response with status code 200, with default header values.

ok
*/
type GetTasksDeletedPhpOK struct {
	Payload models.PaginatedAPIResponse
}

func (o *GetTasksDeletedPhpOK) Error() string {
	return fmt.Sprintf("[GET /tasks/deleted.php][%d] getTasksDeletedPhpOK  %+v", 200, o.Payload)
}
func (o *GetTasksDeletedPhpOK) GetPayload() models.PaginatedAPIResponse {
	return o.Payload
}

func (o *GetTasksDeletedPhpOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTasksDeletedPhpUnauthorized creates a GetTasksDeletedPhpUnauthorized with default headers values
func NewGetTasksDeletedPhpUnauthorized() *GetTasksDeletedPhpUnauthorized {
	return &GetTasksDeletedPhpUnauthorized{}
}

/* GetTasksDeletedPhpUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetTasksDeletedPhpUnauthorized struct {
	Payload *models.Error
}

func (o *GetTasksDeletedPhpUnauthorized) Error() string {
	return fmt.Sprintf("[GET /tasks/deleted.php][%d] getTasksDeletedPhpUnauthorized  %+v", 401, o.Payload)
}
func (o *GetTasksDeletedPhpUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetTasksDeletedPhpUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTasksDeletedPhpTooManyRequests creates a GetTasksDeletedPhpTooManyRequests with default headers values
func NewGetTasksDeletedPhpTooManyRequests() *GetTasksDeletedPhpTooManyRequests {
	return &GetTasksDeletedPhpTooManyRequests{}
}

/* GetTasksDeletedPhpTooManyRequests describes a response with status code 429, with default header values.

TooManyRequests
*/
type GetTasksDeletedPhpTooManyRequests struct {
	Payload *models.Error
}

func (o *GetTasksDeletedPhpTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /tasks/deleted.php][%d] getTasksDeletedPhpTooManyRequests  %+v", 429, o.Payload)
}
func (o *GetTasksDeletedPhpTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetTasksDeletedPhpTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTasksDeletedPhpServiceUnavailable creates a GetTasksDeletedPhpServiceUnavailable with default headers values
func NewGetTasksDeletedPhpServiceUnavailable() *GetTasksDeletedPhpServiceUnavailable {
	return &GetTasksDeletedPhpServiceUnavailable{}
}

/* GetTasksDeletedPhpServiceUnavailable describes a response with status code 503, with default header values.

ServiceUnavailable
*/
type GetTasksDeletedPhpServiceUnavailable struct {
	Payload *models.Error
}

func (o *GetTasksDeletedPhpServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /tasks/deleted.php][%d] getTasksDeletedPhpServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetTasksDeletedPhpServiceUnavailable) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetTasksDeletedPhpServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
