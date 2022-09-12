// Code generated by go-swagger; DO NOT EDIT.

package context

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	ccontext "context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/alswl/go-toodledo/pkg/models"
)

// PostContextsDeletePhpReader is a Reader for the PostContextsDeletePhp structure.
type PostContextsDeletePhpReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostContextsDeletePhpReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostContextsDeletePhpOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPostContextsDeletePhpUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostContextsDeletePhpTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostContextsDeletePhpServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostContextsDeletePhpOK creates a PostContextsDeletePhpOK with default headers values
func NewPostContextsDeletePhpOK() *PostContextsDeletePhpOK {
	return &PostContextsDeletePhpOK{}
}

/*
	PostContextsDeletePhpOK describes a response with status code 200, with default header values.

ok
*/
type PostContextsDeletePhpOK struct {
	Payload *PostContextsDeletePhpOKBody
}

func (o *PostContextsDeletePhpOK) Error() string {
	return fmt.Sprintf("[POST /contexts/delete.php][%d] postContextsDeletePhpOK  %+v", 200, o.Payload)
}
func (o *PostContextsDeletePhpOK) GetPayload() *PostContextsDeletePhpOKBody {
	return o.Payload
}

func (o *PostContextsDeletePhpOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostContextsDeletePhpOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostContextsDeletePhpUnauthorized creates a PostContextsDeletePhpUnauthorized with default headers values
func NewPostContextsDeletePhpUnauthorized() *PostContextsDeletePhpUnauthorized {
	return &PostContextsDeletePhpUnauthorized{}
}

/*
	PostContextsDeletePhpUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PostContextsDeletePhpUnauthorized struct {
	Payload *models.Error
}

func (o *PostContextsDeletePhpUnauthorized) Error() string {
	return fmt.Sprintf("[POST /contexts/delete.php][%d] postContextsDeletePhpUnauthorized  %+v", 401, o.Payload)
}
func (o *PostContextsDeletePhpUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostContextsDeletePhpUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostContextsDeletePhpTooManyRequests creates a PostContextsDeletePhpTooManyRequests with default headers values
func NewPostContextsDeletePhpTooManyRequests() *PostContextsDeletePhpTooManyRequests {
	return &PostContextsDeletePhpTooManyRequests{}
}

/*
	PostContextsDeletePhpTooManyRequests describes a response with status code 429, with default header values.

TooManyRequests
*/
type PostContextsDeletePhpTooManyRequests struct {
	Payload *models.Error
}

func (o *PostContextsDeletePhpTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /contexts/delete.php][%d] postContextsDeletePhpTooManyRequests  %+v", 429, o.Payload)
}
func (o *PostContextsDeletePhpTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostContextsDeletePhpTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostContextsDeletePhpServiceUnavailable creates a PostContextsDeletePhpServiceUnavailable with default headers values
func NewPostContextsDeletePhpServiceUnavailable() *PostContextsDeletePhpServiceUnavailable {
	return &PostContextsDeletePhpServiceUnavailable{}
}

/*
	PostContextsDeletePhpServiceUnavailable describes a response with status code 503, with default header values.

ServiceUnavailable
*/
type PostContextsDeletePhpServiceUnavailable struct {
	Payload *models.Error
}

func (o *PostContextsDeletePhpServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /contexts/delete.php][%d] postContextsDeletePhpServiceUnavailable  %+v", 503, o.Payload)
}
func (o *PostContextsDeletePhpServiceUnavailable) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostContextsDeletePhpServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
PostContextsDeletePhpOKBody post contexts delete php o k body
swagger:model PostContextsDeletePhpOKBody
*/
type PostContextsDeletePhpOKBody struct {

	// deleted
	Deleted int64 `json:"deleted,omitempty"`
}

// Validate validates this post contexts delete php o k body
func (o *PostContextsDeletePhpOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post contexts delete php o k body based on context it is used
func (o *PostContextsDeletePhpOKBody) ContextValidate(ctx ccontext.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostContextsDeletePhpOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostContextsDeletePhpOKBody) UnmarshalBinary(b []byte) error {
	var res PostContextsDeletePhpOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
