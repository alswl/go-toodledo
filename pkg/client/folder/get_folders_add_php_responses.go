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

// GetFoldersAddPhpReader is a Reader for the GetFoldersAddPhp structure.
type GetFoldersAddPhpReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFoldersAddPhpReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFoldersAddPhpOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetFoldersAddPhpUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetFoldersAddPhpTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetFoldersAddPhpServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetFoldersAddPhpOK creates a GetFoldersAddPhpOK with default headers values
func NewGetFoldersAddPhpOK() *GetFoldersAddPhpOK {
	return &GetFoldersAddPhpOK{}
}

/* GetFoldersAddPhpOK describes a response with status code 200, with default header values.

ok
*/
type GetFoldersAddPhpOK struct {
	Payload []*models.Folder
}

func (o *GetFoldersAddPhpOK) Error() string {
	return fmt.Sprintf("[GET /folders/add.php][%d] getFoldersAddPhpOK  %+v", 200, o.Payload)
}
func (o *GetFoldersAddPhpOK) GetPayload() []*models.Folder {
	return o.Payload
}

func (o *GetFoldersAddPhpOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFoldersAddPhpUnauthorized creates a GetFoldersAddPhpUnauthorized with default headers values
func NewGetFoldersAddPhpUnauthorized() *GetFoldersAddPhpUnauthorized {
	return &GetFoldersAddPhpUnauthorized{}
}

/* GetFoldersAddPhpUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetFoldersAddPhpUnauthorized struct {
	Payload *models.Error
}

func (o *GetFoldersAddPhpUnauthorized) Error() string {
	return fmt.Sprintf("[GET /folders/add.php][%d] getFoldersAddPhpUnauthorized  %+v", 401, o.Payload)
}
func (o *GetFoldersAddPhpUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetFoldersAddPhpUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFoldersAddPhpTooManyRequests creates a GetFoldersAddPhpTooManyRequests with default headers values
func NewGetFoldersAddPhpTooManyRequests() *GetFoldersAddPhpTooManyRequests {
	return &GetFoldersAddPhpTooManyRequests{}
}

/* GetFoldersAddPhpTooManyRequests describes a response with status code 429, with default header values.

TooManyRequests
*/
type GetFoldersAddPhpTooManyRequests struct {
	Payload *models.Error
}

func (o *GetFoldersAddPhpTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /folders/add.php][%d] getFoldersAddPhpTooManyRequests  %+v", 429, o.Payload)
}
func (o *GetFoldersAddPhpTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetFoldersAddPhpTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFoldersAddPhpServiceUnavailable creates a GetFoldersAddPhpServiceUnavailable with default headers values
func NewGetFoldersAddPhpServiceUnavailable() *GetFoldersAddPhpServiceUnavailable {
	return &GetFoldersAddPhpServiceUnavailable{}
}

/* GetFoldersAddPhpServiceUnavailable describes a response with status code 503, with default header values.

ServiceUnavailable
*/
type GetFoldersAddPhpServiceUnavailable struct {
	Payload *models.Error
}

func (o *GetFoldersAddPhpServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /folders/add.php][%d] getFoldersAddPhpServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetFoldersAddPhpServiceUnavailable) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetFoldersAddPhpServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
