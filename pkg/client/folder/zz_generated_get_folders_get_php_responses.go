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

// GetFoldersGetPhpReader is a Reader for the GetFoldersGetPhp structure.
type GetFoldersGetPhpReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFoldersGetPhpReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFoldersGetPhpOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetFoldersGetPhpUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetFoldersGetPhpTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetFoldersGetPhpServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetFoldersGetPhpOK creates a GetFoldersGetPhpOK with default headers values
func NewGetFoldersGetPhpOK() *GetFoldersGetPhpOK {
	return &GetFoldersGetPhpOK{}
}

/*
	GetFoldersGetPhpOK describes a response with status code 200, with default header values.

ok
*/
type GetFoldersGetPhpOK struct {
	Payload []*models.Folder
}

func (o *GetFoldersGetPhpOK) Error() string {
	return fmt.Sprintf("[GET /folders/get.php][%d] getFoldersGetPhpOK  %+v", 200, o.Payload)
}
func (o *GetFoldersGetPhpOK) GetPayload() []*models.Folder {
	return o.Payload
}

func (o *GetFoldersGetPhpOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFoldersGetPhpUnauthorized creates a GetFoldersGetPhpUnauthorized with default headers values
func NewGetFoldersGetPhpUnauthorized() *GetFoldersGetPhpUnauthorized {
	return &GetFoldersGetPhpUnauthorized{}
}

/*
	GetFoldersGetPhpUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetFoldersGetPhpUnauthorized struct {
	Payload *models.Error
}

func (o *GetFoldersGetPhpUnauthorized) Error() string {
	return fmt.Sprintf("[GET /folders/get.php][%d] getFoldersGetPhpUnauthorized  %+v", 401, o.Payload)
}
func (o *GetFoldersGetPhpUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetFoldersGetPhpUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFoldersGetPhpTooManyRequests creates a GetFoldersGetPhpTooManyRequests with default headers values
func NewGetFoldersGetPhpTooManyRequests() *GetFoldersGetPhpTooManyRequests {
	return &GetFoldersGetPhpTooManyRequests{}
}

/*
	GetFoldersGetPhpTooManyRequests describes a response with status code 429, with default header values.

TooManyRequests
*/
type GetFoldersGetPhpTooManyRequests struct {
	Payload *models.Error
}

func (o *GetFoldersGetPhpTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /folders/get.php][%d] getFoldersGetPhpTooManyRequests  %+v", 429, o.Payload)
}
func (o *GetFoldersGetPhpTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetFoldersGetPhpTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFoldersGetPhpServiceUnavailable creates a GetFoldersGetPhpServiceUnavailable with default headers values
func NewGetFoldersGetPhpServiceUnavailable() *GetFoldersGetPhpServiceUnavailable {
	return &GetFoldersGetPhpServiceUnavailable{}
}

/*
	GetFoldersGetPhpServiceUnavailable describes a response with status code 503, with default header values.

ServiceUnavailable
*/
type GetFoldersGetPhpServiceUnavailable struct {
	Payload *models.Error
}

func (o *GetFoldersGetPhpServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /folders/get.php][%d] getFoldersGetPhpServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetFoldersGetPhpServiceUnavailable) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetFoldersGetPhpServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
