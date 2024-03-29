// Code generated by go-swagger; DO NOT EDIT.

package goal

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/alswl/go-toodledo/pkg/models"
)

// PostGoalsEditPhpReader is a Reader for the PostGoalsEditPhp structure.
type PostGoalsEditPhpReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostGoalsEditPhpReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostGoalsEditPhpOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPostGoalsEditPhpUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostGoalsEditPhpTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostGoalsEditPhpServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostGoalsEditPhpOK creates a PostGoalsEditPhpOK with default headers values
func NewPostGoalsEditPhpOK() *PostGoalsEditPhpOK {
	return &PostGoalsEditPhpOK{}
}

/*
PostGoalsEditPhpOK describes a response with status code 200, with default header values.

ok
*/
type PostGoalsEditPhpOK struct {
	Payload []*models.Goal
}

func (o *PostGoalsEditPhpOK) Error() string {
	return fmt.Sprintf("[POST /goals/edit.php][%d] postGoalsEditPhpOK  %+v", 200, o.Payload)
}
func (o *PostGoalsEditPhpOK) GetPayload() []*models.Goal {
	return o.Payload
}

func (o *PostGoalsEditPhpOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGoalsEditPhpUnauthorized creates a PostGoalsEditPhpUnauthorized with default headers values
func NewPostGoalsEditPhpUnauthorized() *PostGoalsEditPhpUnauthorized {
	return &PostGoalsEditPhpUnauthorized{}
}

/*
PostGoalsEditPhpUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PostGoalsEditPhpUnauthorized struct {
	Payload *models.Error
}

func (o *PostGoalsEditPhpUnauthorized) Error() string {
	return fmt.Sprintf("[POST /goals/edit.php][%d] postGoalsEditPhpUnauthorized  %+v", 401, o.Payload)
}
func (o *PostGoalsEditPhpUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostGoalsEditPhpUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGoalsEditPhpTooManyRequests creates a PostGoalsEditPhpTooManyRequests with default headers values
func NewPostGoalsEditPhpTooManyRequests() *PostGoalsEditPhpTooManyRequests {
	return &PostGoalsEditPhpTooManyRequests{}
}

/*
PostGoalsEditPhpTooManyRequests describes a response with status code 429, with default header values.

TooManyRequests
*/
type PostGoalsEditPhpTooManyRequests struct {
	Payload *models.Error
}

func (o *PostGoalsEditPhpTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /goals/edit.php][%d] postGoalsEditPhpTooManyRequests  %+v", 429, o.Payload)
}
func (o *PostGoalsEditPhpTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostGoalsEditPhpTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGoalsEditPhpServiceUnavailable creates a PostGoalsEditPhpServiceUnavailable with default headers values
func NewPostGoalsEditPhpServiceUnavailable() *PostGoalsEditPhpServiceUnavailable {
	return &PostGoalsEditPhpServiceUnavailable{}
}

/*
PostGoalsEditPhpServiceUnavailable describes a response with status code 503, with default header values.

ServiceUnavailable
*/
type PostGoalsEditPhpServiceUnavailable struct {
	Payload *models.Error
}

func (o *PostGoalsEditPhpServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /goals/edit.php][%d] postGoalsEditPhpServiceUnavailable  %+v", 503, o.Payload)
}
func (o *PostGoalsEditPhpServiceUnavailable) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostGoalsEditPhpServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
