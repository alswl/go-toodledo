// Code generated by go-swagger; DO NOT EDIT.

package goal

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

// NewPostGoalsEditPhpParams creates a new PostGoalsEditPhpParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostGoalsEditPhpParams() *PostGoalsEditPhpParams {
	return &PostGoalsEditPhpParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostGoalsEditPhpParamsWithTimeout creates a new PostGoalsEditPhpParams object
// with the ability to set a timeout on a request.
func NewPostGoalsEditPhpParamsWithTimeout(timeout time.Duration) *PostGoalsEditPhpParams {
	return &PostGoalsEditPhpParams{
		timeout: timeout,
	}
}

// NewPostGoalsEditPhpParamsWithContext creates a new PostGoalsEditPhpParams object
// with the ability to set a context for a request.
func NewPostGoalsEditPhpParamsWithContext(ctx ccontext.Context) *PostGoalsEditPhpParams {
	return &PostGoalsEditPhpParams{
		Context: ctx,
	}
}

// NewPostGoalsEditPhpParamsWithHTTPClient creates a new PostGoalsEditPhpParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostGoalsEditPhpParamsWithHTTPClient(client *http.Client) *PostGoalsEditPhpParams {
	return &PostGoalsEditPhpParams{
		HTTPClient: client,
	}
}

/* PostGoalsEditPhpParams contains all the parameters to send to the API endpoint
   for the post goals edit php operation.

   Typically these are written to a http.Request.
*/
type PostGoalsEditPhpParams struct {

	// Archived.
	Archived *int64

	// Contributes.
	Contributes *int64

	// ID.
	ID string

	// Level.
	Level *int64

	// Name.
	Name *string

	// Note.
	Note *string

	// Private.
	Private *int64

	timeout    time.Duration
	Context    ccontext.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post goals edit php params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostGoalsEditPhpParams) WithDefaults() *PostGoalsEditPhpParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post goals edit php params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostGoalsEditPhpParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post goals edit php params
func (o *PostGoalsEditPhpParams) WithTimeout(timeout time.Duration) *PostGoalsEditPhpParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post goals edit php params
func (o *PostGoalsEditPhpParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post goals edit php params
func (o *PostGoalsEditPhpParams) WithContext(ctx ccontext.Context) *PostGoalsEditPhpParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post goals edit php params
func (o *PostGoalsEditPhpParams) SetContext(ctx ccontext.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post goals edit php params
func (o *PostGoalsEditPhpParams) WithHTTPClient(client *http.Client) *PostGoalsEditPhpParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post goals edit php params
func (o *PostGoalsEditPhpParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithArchived adds the archived to the post goals edit php params
func (o *PostGoalsEditPhpParams) WithArchived(archived *int64) *PostGoalsEditPhpParams {
	o.SetArchived(archived)
	return o
}

// SetArchived adds the archived to the post goals edit php params
func (o *PostGoalsEditPhpParams) SetArchived(archived *int64) {
	o.Archived = archived
}

// WithContributes adds the contributes to the post goals edit php params
func (o *PostGoalsEditPhpParams) WithContributes(contributes *int64) *PostGoalsEditPhpParams {
	o.SetContributes(contributes)
	return o
}

// SetContributes adds the contributes to the post goals edit php params
func (o *PostGoalsEditPhpParams) SetContributes(contributes *int64) {
	o.Contributes = contributes
}

// WithID adds the id to the post goals edit php params
func (o *PostGoalsEditPhpParams) WithID(id string) *PostGoalsEditPhpParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post goals edit php params
func (o *PostGoalsEditPhpParams) SetID(id string) {
	o.ID = id
}

// WithLevel adds the level to the post goals edit php params
func (o *PostGoalsEditPhpParams) WithLevel(level *int64) *PostGoalsEditPhpParams {
	o.SetLevel(level)
	return o
}

// SetLevel adds the level to the post goals edit php params
func (o *PostGoalsEditPhpParams) SetLevel(level *int64) {
	o.Level = level
}

// WithName adds the name to the post goals edit php params
func (o *PostGoalsEditPhpParams) WithName(name *string) *PostGoalsEditPhpParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the post goals edit php params
func (o *PostGoalsEditPhpParams) SetName(name *string) {
	o.Name = name
}

// WithNote adds the note to the post goals edit php params
func (o *PostGoalsEditPhpParams) WithNote(note *string) *PostGoalsEditPhpParams {
	o.SetNote(note)
	return o
}

// SetNote adds the note to the post goals edit php params
func (o *PostGoalsEditPhpParams) SetNote(note *string) {
	o.Note = note
}

// WithPrivate adds the private to the post goals edit php params
func (o *PostGoalsEditPhpParams) WithPrivate(private *int64) *PostGoalsEditPhpParams {
	o.SetPrivate(private)
	return o
}

// SetPrivate adds the private to the post goals edit php params
func (o *PostGoalsEditPhpParams) SetPrivate(private *int64) {
	o.Private = private
}

// WriteToRequest writes these params to a swagger request
func (o *PostGoalsEditPhpParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Archived != nil {

		// query param archived
		var qrArchived int64

		if o.Archived != nil {
			qrArchived = *o.Archived
		}
		qArchived := swag.FormatInt64(qrArchived)
		if qArchived != "" {

			if err := r.SetQueryParam("archived", qArchived); err != nil {
				return err
			}
		}
	}

	if o.Contributes != nil {

		// query param contributes
		var qrContributes int64

		if o.Contributes != nil {
			qrContributes = *o.Contributes
		}
		qContributes := swag.FormatInt64(qrContributes)
		if qContributes != "" {

			if err := r.SetQueryParam("contributes", qContributes); err != nil {
				return err
			}
		}
	}

	// query param id
	qrID := o.ID
	qID := qrID
	if qID != "" {

		if err := r.SetQueryParam("id", qID); err != nil {
			return err
		}
	}

	if o.Level != nil {

		// query param level
		var qrLevel int64

		if o.Level != nil {
			qrLevel = *o.Level
		}
		qLevel := swag.FormatInt64(qrLevel)
		if qLevel != "" {

			if err := r.SetQueryParam("level", qLevel); err != nil {
				return err
			}
		}
	}

	if o.Name != nil {

		// query param name
		var qrName string

		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}
	}

	if o.Note != nil {

		// query param note
		var qrNote string

		if o.Note != nil {
			qrNote = *o.Note
		}
		qNote := qrNote
		if qNote != "" {

			if err := r.SetQueryParam("note", qNote); err != nil {
				return err
			}
		}
	}

	if o.Private != nil {

		// query param private
		var qrPrivate int64

		if o.Private != nil {
			qrPrivate = *o.Private
		}
		qPrivate := swag.FormatInt64(qrPrivate)
		if qPrivate != "" {

			if err := r.SetQueryParam("private", qPrivate); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}