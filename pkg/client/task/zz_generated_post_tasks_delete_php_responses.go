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

// PostTasksDeletePhpReader is a Reader for the PostTasksDeletePhp structure.
type PostTasksDeletePhpReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostTasksDeletePhpReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostTasksDeletePhpOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPostTasksDeletePhpUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostTasksDeletePhpTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostTasksDeletePhpServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostTasksDeletePhpOK creates a PostTasksDeletePhpOK with default headers values
func NewPostTasksDeletePhpOK() *PostTasksDeletePhpOK {
	return &PostTasksDeletePhpOK{}
}

/* PostTasksDeletePhpOK describes a response with status code 200, with default header values.

ok
*/
type PostTasksDeletePhpOK struct {
	Payload []*models.TaskDeleteItem
}

func (o *PostTasksDeletePhpOK) Error() string {
	return fmt.Sprintf("[POST /tasks/delete.php][%d] postTasksDeletePhpOK  %+v", 200, o.Payload)
}
func (o *PostTasksDeletePhpOK) GetPayload() []*models.TaskDeleteItem {
	return o.Payload
}

func (o *PostTasksDeletePhpOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTasksDeletePhpUnauthorized creates a PostTasksDeletePhpUnauthorized with default headers values
func NewPostTasksDeletePhpUnauthorized() *PostTasksDeletePhpUnauthorized {
	return &PostTasksDeletePhpUnauthorized{}
}

/* PostTasksDeletePhpUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PostTasksDeletePhpUnauthorized struct {
	Payload *models.Error
}

func (o *PostTasksDeletePhpUnauthorized) Error() string {
	return fmt.Sprintf("[POST /tasks/delete.php][%d] postTasksDeletePhpUnauthorized  %+v", 401, o.Payload)
}
func (o *PostTasksDeletePhpUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostTasksDeletePhpUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTasksDeletePhpTooManyRequests creates a PostTasksDeletePhpTooManyRequests with default headers values
func NewPostTasksDeletePhpTooManyRequests() *PostTasksDeletePhpTooManyRequests {
	return &PostTasksDeletePhpTooManyRequests{}
}

/* PostTasksDeletePhpTooManyRequests describes a response with status code 429, with default header values.

TooManyRequests
*/
type PostTasksDeletePhpTooManyRequests struct {
	Payload *models.Error
}

func (o *PostTasksDeletePhpTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /tasks/delete.php][%d] postTasksDeletePhpTooManyRequests  %+v", 429, o.Payload)
}
func (o *PostTasksDeletePhpTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostTasksDeletePhpTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTasksDeletePhpServiceUnavailable creates a PostTasksDeletePhpServiceUnavailable with default headers values
func NewPostTasksDeletePhpServiceUnavailable() *PostTasksDeletePhpServiceUnavailable {
	return &PostTasksDeletePhpServiceUnavailable{}
}

/* PostTasksDeletePhpServiceUnavailable describes a response with status code 503, with default header values.

ServiceUnavailable
*/
type PostTasksDeletePhpServiceUnavailable struct {
	Payload *models.Error
}

func (o *PostTasksDeletePhpServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /tasks/delete.php][%d] postTasksDeletePhpServiceUnavailable  %+v", 503, o.Payload)
}
func (o *PostTasksDeletePhpServiceUnavailable) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostTasksDeletePhpServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
