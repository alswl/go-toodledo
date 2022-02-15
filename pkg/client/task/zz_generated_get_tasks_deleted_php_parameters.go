// Code generated by go-swagger; DO NOT EDIT.

package task

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
	"github.com/go-openapi/swag"
)

// NewGetTasksDeletedPhpParams creates a new GetTasksDeletedPhpParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetTasksDeletedPhpParams() *GetTasksDeletedPhpParams {
	return &GetTasksDeletedPhpParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetTasksDeletedPhpParamsWithTimeout creates a new GetTasksDeletedPhpParams object
// with the ability to set a timeout on a request.
func NewGetTasksDeletedPhpParamsWithTimeout(timeout time.Duration) *GetTasksDeletedPhpParams {
	return &GetTasksDeletedPhpParams{
		timeout: timeout,
	}
}

// NewGetTasksDeletedPhpParamsWithContext creates a new GetTasksDeletedPhpParams object
// with the ability to set a context for a request.
func NewGetTasksDeletedPhpParamsWithContext(ctx ccontext.Context) *GetTasksDeletedPhpParams {
	return &GetTasksDeletedPhpParams{
		Context: ctx,
	}
}

// NewGetTasksDeletedPhpParamsWithHTTPClient creates a new GetTasksDeletedPhpParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetTasksDeletedPhpParamsWithHTTPClient(client *http.Client) *GetTasksDeletedPhpParams {
	return &GetTasksDeletedPhpParams{
		HTTPClient: client,
	}
}

/* GetTasksDeletedPhpParams contains all the parameters to send to the API endpoint
   for the get tasks deleted php operation.

   Typically these are written to a http.Request.
*/
type GetTasksDeletedPhpParams struct {

	/* After.

	   A GMT unix timestamp. Used to find tasks with a modified date and time before this date and time.
	*/
	After *int64

	timeout    time.Duration
	Context    ccontext.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get tasks deleted php params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTasksDeletedPhpParams) WithDefaults() *GetTasksDeletedPhpParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get tasks deleted php params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTasksDeletedPhpParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get tasks deleted php params
func (o *GetTasksDeletedPhpParams) WithTimeout(timeout time.Duration) *GetTasksDeletedPhpParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get tasks deleted php params
func (o *GetTasksDeletedPhpParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get tasks deleted php params
func (o *GetTasksDeletedPhpParams) WithContext(ctx ccontext.Context) *GetTasksDeletedPhpParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get tasks deleted php params
func (o *GetTasksDeletedPhpParams) SetContext(ctx ccontext.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get tasks deleted php params
func (o *GetTasksDeletedPhpParams) WithHTTPClient(client *http.Client) *GetTasksDeletedPhpParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get tasks deleted php params
func (o *GetTasksDeletedPhpParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAfter adds the after to the get tasks deleted php params
func (o *GetTasksDeletedPhpParams) WithAfter(after *int64) *GetTasksDeletedPhpParams {
	o.SetAfter(after)
	return o
}

// SetAfter adds the after to the get tasks deleted php params
func (o *GetTasksDeletedPhpParams) SetAfter(after *int64) {
	o.After = after
}

// WriteToRequest writes these params to a swagger request
func (o *GetTasksDeletedPhpParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.After != nil {

		// query param after
		var qrAfter int64

		if o.After != nil {
			qrAfter = *o.After
		}
		qAfter := swag.FormatInt64(qrAfter)
		if qAfter != "" {

			if err := r.SetQueryParam("after", qAfter); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
