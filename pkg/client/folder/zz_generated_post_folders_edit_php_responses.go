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

// PostFoldersEditPhpReader is a Reader for the PostFoldersEditPhp structure.
type PostFoldersEditPhpReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostFoldersEditPhpReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostFoldersEditPhpOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPostFoldersEditPhpUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostFoldersEditPhpTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostFoldersEditPhpServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostFoldersEditPhpOK creates a PostFoldersEditPhpOK with default headers values
func NewPostFoldersEditPhpOK() *PostFoldersEditPhpOK {
	return &PostFoldersEditPhpOK{}
}

/* PostFoldersEditPhpOK describes a response with status code 200, with default header values.

ok
*/
type PostFoldersEditPhpOK struct {
	Payload []*models.Folder
}

// Error ...
func (o *PostFoldersEditPhpOK) Error() string {
	return fmt.Sprintf("[POST /folders/edit.php][%d] postFoldersEditPhpOK  %+v", 200, o.Payload)
}

// GetPayload ...
func (o *PostFoldersEditPhpOK) GetPayload() []*models.Folder {
	return o.Payload
}

func (o *PostFoldersEditPhpOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFoldersEditPhpUnauthorized creates a PostFoldersEditPhpUnauthorized with default headers values
func NewPostFoldersEditPhpUnauthorized() *PostFoldersEditPhpUnauthorized {
	return &PostFoldersEditPhpUnauthorized{}
}

/* PostFoldersEditPhpUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PostFoldersEditPhpUnauthorized struct {
	Payload *models.Error
}

// Error ...
func (o *PostFoldersEditPhpUnauthorized) Error() string {
	return fmt.Sprintf("[POST /folders/edit.php][%d] postFoldersEditPhpUnauthorized  %+v", 401, o.Payload)
}

// GetPayload ...
func (o *PostFoldersEditPhpUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostFoldersEditPhpUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFoldersEditPhpTooManyRequests creates a PostFoldersEditPhpTooManyRequests with default headers values
func NewPostFoldersEditPhpTooManyRequests() *PostFoldersEditPhpTooManyRequests {
	return &PostFoldersEditPhpTooManyRequests{}
}

/* PostFoldersEditPhpTooManyRequests describes a response with status code 429, with default header values.

TooManyRequests
*/
type PostFoldersEditPhpTooManyRequests struct {
	Payload *models.Error
}

// Error ...
func (o *PostFoldersEditPhpTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /folders/edit.php][%d] postFoldersEditPhpTooManyRequests  %+v", 429, o.Payload)
}

// GetPayload ...
func (o *PostFoldersEditPhpTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostFoldersEditPhpTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFoldersEditPhpServiceUnavailable creates a PostFoldersEditPhpServiceUnavailable with default headers values
func NewPostFoldersEditPhpServiceUnavailable() *PostFoldersEditPhpServiceUnavailable {
	return &PostFoldersEditPhpServiceUnavailable{}
}

/* PostFoldersEditPhpServiceUnavailable describes a response with status code 503, with default header values.

ServiceUnavailable
*/
type PostFoldersEditPhpServiceUnavailable struct {
	Payload *models.Error
}

// Error ...
func (o *PostFoldersEditPhpServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /folders/edit.php][%d] postFoldersEditPhpServiceUnavailable  %+v", 503, o.Payload)
}

// GetPayload ...
func (o *PostFoldersEditPhpServiceUnavailable) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostFoldersEditPhpServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
