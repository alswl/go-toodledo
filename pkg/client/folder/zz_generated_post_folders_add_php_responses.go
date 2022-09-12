// Code generated by go-swagger; DO NOT EDIT.

package folder

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/alswl/go-toodledo/pkg/models"
)

// PostFoldersAddPhpReader is a Reader for the PostFoldersAddPhp structure.
type PostFoldersAddPhpReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostFoldersAddPhpReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostFoldersAddPhpOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPostFoldersAddPhpUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostFoldersAddPhpTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostFoldersAddPhpServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostFoldersAddPhpOK creates a PostFoldersAddPhpOK with default headers values
func NewPostFoldersAddPhpOK() *PostFoldersAddPhpOK {
	return &PostFoldersAddPhpOK{}
}

/*
	PostFoldersAddPhpOK describes a response with status code 200, with default header values.

ok
*/
type PostFoldersAddPhpOK struct {
	Payload []*models.Folder
}

func (o *PostFoldersAddPhpOK) Error() string {
	return fmt.Sprintf("[POST /folders/add.php][%d] postFoldersAddPhpOK  %+v", 200, o.Payload)
}
func (o *PostFoldersAddPhpOK) GetPayload() []*models.Folder {
	return o.Payload
}

func (o *PostFoldersAddPhpOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFoldersAddPhpUnauthorized creates a PostFoldersAddPhpUnauthorized with default headers values
func NewPostFoldersAddPhpUnauthorized() *PostFoldersAddPhpUnauthorized {
	return &PostFoldersAddPhpUnauthorized{}
}

/*
	PostFoldersAddPhpUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PostFoldersAddPhpUnauthorized struct {
	Payload *models.Error
}

func (o *PostFoldersAddPhpUnauthorized) Error() string {
	return fmt.Sprintf("[POST /folders/add.php][%d] postFoldersAddPhpUnauthorized  %+v", 401, o.Payload)
}
func (o *PostFoldersAddPhpUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostFoldersAddPhpUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFoldersAddPhpTooManyRequests creates a PostFoldersAddPhpTooManyRequests with default headers values
func NewPostFoldersAddPhpTooManyRequests() *PostFoldersAddPhpTooManyRequests {
	return &PostFoldersAddPhpTooManyRequests{}
}

/*
	PostFoldersAddPhpTooManyRequests describes a response with status code 429, with default header values.

TooManyRequests
*/
type PostFoldersAddPhpTooManyRequests struct {
	Payload *models.Error
}

func (o *PostFoldersAddPhpTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /folders/add.php][%d] postFoldersAddPhpTooManyRequests  %+v", 429, o.Payload)
}
func (o *PostFoldersAddPhpTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostFoldersAddPhpTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFoldersAddPhpServiceUnavailable creates a PostFoldersAddPhpServiceUnavailable with default headers values
func NewPostFoldersAddPhpServiceUnavailable() *PostFoldersAddPhpServiceUnavailable {
	return &PostFoldersAddPhpServiceUnavailable{}
}

/*
	PostFoldersAddPhpServiceUnavailable describes a response with status code 503, with default header values.

ServiceUnavailable
*/
type PostFoldersAddPhpServiceUnavailable struct {
	Payload *models.Error
}

func (o *PostFoldersAddPhpServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /folders/add.php][%d] postFoldersAddPhpServiceUnavailable  %+v", 503, o.Payload)
}
func (o *PostFoldersAddPhpServiceUnavailable) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostFoldersAddPhpServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
