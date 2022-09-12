// Code generated by go-swagger; DO NOT EDIT.

package account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	ccontext "context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetAccountGetPhpParams creates a new GetAccountGetPhpParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAccountGetPhpParams() *GetAccountGetPhpParams {
	return &GetAccountGetPhpParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAccountGetPhpParamsWithTimeout creates a new GetAccountGetPhpParams object
// with the ability to set a timeout on a request.
func NewGetAccountGetPhpParamsWithTimeout(timeout time.Duration) *GetAccountGetPhpParams {
	return &GetAccountGetPhpParams{
		timeout: timeout,
	}
}

// NewGetAccountGetPhpParamsWithContext creates a new GetAccountGetPhpParams object
// with the ability to set a context for a request.
func NewGetAccountGetPhpParamsWithContext(ctx ccontext.Context) *GetAccountGetPhpParams {
	return &GetAccountGetPhpParams{
		Context: ctx,
	}
}

// NewGetAccountGetPhpParamsWithHTTPClient creates a new GetAccountGetPhpParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAccountGetPhpParamsWithHTTPClient(client *http.Client) *GetAccountGetPhpParams {
	return &GetAccountGetPhpParams{
		HTTPClient: client,
	}
}

/*
GetAccountGetPhpParams contains all the parameters to send to the API endpoint

	for the get account get php operation.

	Typically these are written to a http.Request.
*/
type GetAccountGetPhpParams struct {
	timeout    time.Duration
	Context    ccontext.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get account get php params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAccountGetPhpParams) WithDefaults() *GetAccountGetPhpParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get account get php params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAccountGetPhpParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get account get php params
func (o *GetAccountGetPhpParams) WithTimeout(timeout time.Duration) *GetAccountGetPhpParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get account get php params
func (o *GetAccountGetPhpParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get account get php params
func (o *GetAccountGetPhpParams) WithContext(ctx ccontext.Context) *GetAccountGetPhpParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get account get php params
func (o *GetAccountGetPhpParams) SetContext(ctx ccontext.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get account get php params
func (o *GetAccountGetPhpParams) WithHTTPClient(client *http.Client) *GetAccountGetPhpParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get account get php params
func (o *GetAccountGetPhpParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAccountGetPhpParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
