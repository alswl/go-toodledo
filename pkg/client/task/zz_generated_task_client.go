// Code generated by go-swagger; DO NOT EDIT.

package task

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new task API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for task API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetTasksGetPhp(params *GetTasksGetPhpParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetTasksGetPhpOK, error)

	PostTasksAddPhp(params *PostTasksAddPhpParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostTasksAddPhpOK, error)

	PostTasksDeletePhp(params *PostTasksDeletePhpParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostTasksDeletePhpOK, error)

	PostTasksEditPhp(params *PostTasksEditPhpParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostTasksEditPhpOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetTasksGetPhp get tasks get php API
*/
func (a *Client) GetTasksGetPhp(params *GetTasksGetPhpParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetTasksGetPhpOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTasksGetPhpParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetTasksGetPhp",
		Method:             "GET",
		PathPattern:        "/tasks/get.php",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetTasksGetPhpReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetTasksGetPhpOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetTasksGetPhp: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostTasksAddPhp post tasks add php API
*/
func (a *Client) PostTasksAddPhp(params *PostTasksAddPhpParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostTasksAddPhpOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostTasksAddPhpParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostTasksAddPhp",
		Method:             "POST",
		PathPattern:        "/tasks/add.php",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostTasksAddPhpReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostTasksAddPhpOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostTasksAddPhp: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostTasksDeletePhp post tasks delete php API
*/
func (a *Client) PostTasksDeletePhp(params *PostTasksDeletePhpParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostTasksDeletePhpOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostTasksDeletePhpParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostTasksDeletePhp",
		Method:             "POST",
		PathPattern:        "/tasks/delete.php",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostTasksDeletePhpReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostTasksDeletePhpOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostTasksDeletePhp: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostTasksEditPhp post tasks edit php API
*/
func (a *Client) PostTasksEditPhp(params *PostTasksEditPhpParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostTasksEditPhpOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostTasksEditPhpParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostTasksEditPhp",
		Method:             "POST",
		PathPattern:        "/tasks/edit.php",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostTasksEditPhpReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostTasksEditPhpOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostTasksEditPhp: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
