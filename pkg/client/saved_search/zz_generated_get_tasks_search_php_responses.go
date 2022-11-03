// Code generated by go-swagger; DO NOT EDIT.

package saved_search

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/alswl/go-toodledo/pkg/models"
)

// GetTasksSearchPhpReader is a Reader for the GetTasksSearchPhp structure.
type GetTasksSearchPhpReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTasksSearchPhpReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTasksSearchPhpOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetTasksSearchPhpUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetTasksSearchPhpTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetTasksSearchPhpServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTasksSearchPhpOK creates a GetTasksSearchPhpOK with default headers values
func NewGetTasksSearchPhpOK() *GetTasksSearchPhpOK {
	return &GetTasksSearchPhpOK{}
}

/*
GetTasksSearchPhpOK describes a response with status code 200, with default header values.

ok
*/
type GetTasksSearchPhpOK struct {
	Payload map[string]models.SavedSearch
}

func (o *GetTasksSearchPhpOK) Error() string {
	return fmt.Sprintf("[GET /tasks/search.php][%d] getTasksSearchPhpOK  %+v", 200, o.Payload)
}
func (o *GetTasksSearchPhpOK) GetPayload() map[string]models.SavedSearch {
	return o.Payload
}

func (o *GetTasksSearchPhpOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTasksSearchPhpUnauthorized creates a GetTasksSearchPhpUnauthorized with default headers values
func NewGetTasksSearchPhpUnauthorized() *GetTasksSearchPhpUnauthorized {
	return &GetTasksSearchPhpUnauthorized{}
}

/*
GetTasksSearchPhpUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetTasksSearchPhpUnauthorized struct {
	Payload *models.Error
}

func (o *GetTasksSearchPhpUnauthorized) Error() string {
	return fmt.Sprintf("[GET /tasks/search.php][%d] getTasksSearchPhpUnauthorized  %+v", 401, o.Payload)
}
func (o *GetTasksSearchPhpUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetTasksSearchPhpUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTasksSearchPhpTooManyRequests creates a GetTasksSearchPhpTooManyRequests with default headers values
func NewGetTasksSearchPhpTooManyRequests() *GetTasksSearchPhpTooManyRequests {
	return &GetTasksSearchPhpTooManyRequests{}
}

/*
GetTasksSearchPhpTooManyRequests describes a response with status code 429, with default header values.

TooManyRequests
*/
type GetTasksSearchPhpTooManyRequests struct {
	Payload *models.Error
}

func (o *GetTasksSearchPhpTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /tasks/search.php][%d] getTasksSearchPhpTooManyRequests  %+v", 429, o.Payload)
}
func (o *GetTasksSearchPhpTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetTasksSearchPhpTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTasksSearchPhpServiceUnavailable creates a GetTasksSearchPhpServiceUnavailable with default headers values
func NewGetTasksSearchPhpServiceUnavailable() *GetTasksSearchPhpServiceUnavailable {
	return &GetTasksSearchPhpServiceUnavailable{}
}

/*
GetTasksSearchPhpServiceUnavailable describes a response with status code 503, with default header values.

ServiceUnavailable
*/
type GetTasksSearchPhpServiceUnavailable struct {
	Payload *models.Error
}

func (o *GetTasksSearchPhpServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /tasks/search.php][%d] getTasksSearchPhpServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetTasksSearchPhpServiceUnavailable) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetTasksSearchPhpServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
