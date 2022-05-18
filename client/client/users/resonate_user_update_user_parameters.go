// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/resonatecoop/user-api/client/models"
)

// NewResonateUserUpdateUserParams creates a new ResonateUserUpdateUserParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewResonateUserUpdateUserParams() *ResonateUserUpdateUserParams {
	return &ResonateUserUpdateUserParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewResonateUserUpdateUserParamsWithTimeout creates a new ResonateUserUpdateUserParams object
// with the ability to set a timeout on a request.
func NewResonateUserUpdateUserParamsWithTimeout(timeout time.Duration) *ResonateUserUpdateUserParams {
	return &ResonateUserUpdateUserParams{
		timeout: timeout,
	}
}

// NewResonateUserUpdateUserParamsWithContext creates a new ResonateUserUpdateUserParams object
// with the ability to set a context for a request.
func NewResonateUserUpdateUserParamsWithContext(ctx context.Context) *ResonateUserUpdateUserParams {
	return &ResonateUserUpdateUserParams{
		Context: ctx,
	}
}

// NewResonateUserUpdateUserParamsWithHTTPClient creates a new ResonateUserUpdateUserParams object
// with the ability to set a custom HTTPClient for a request.
func NewResonateUserUpdateUserParamsWithHTTPClient(client *http.Client) *ResonateUserUpdateUserParams {
	return &ResonateUserUpdateUserParams{
		HTTPClient: client,
	}
}

/* ResonateUserUpdateUserParams contains all the parameters to send to the API endpoint
   for the resonate user update user operation.

   Typically these are written to a http.Request.
*/
type ResonateUserUpdateUserParams struct {

	// Body.
	Body *models.UserUserUpdateRequest

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the resonate user update user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ResonateUserUpdateUserParams) WithDefaults() *ResonateUserUpdateUserParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the resonate user update user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ResonateUserUpdateUserParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the resonate user update user params
func (o *ResonateUserUpdateUserParams) WithTimeout(timeout time.Duration) *ResonateUserUpdateUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the resonate user update user params
func (o *ResonateUserUpdateUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the resonate user update user params
func (o *ResonateUserUpdateUserParams) WithContext(ctx context.Context) *ResonateUserUpdateUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the resonate user update user params
func (o *ResonateUserUpdateUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the resonate user update user params
func (o *ResonateUserUpdateUserParams) WithHTTPClient(client *http.Client) *ResonateUserUpdateUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the resonate user update user params
func (o *ResonateUserUpdateUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the resonate user update user params
func (o *ResonateUserUpdateUserParams) WithBody(body *models.UserUserUpdateRequest) *ResonateUserUpdateUserParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the resonate user update user params
func (o *ResonateUserUpdateUserParams) SetBody(body *models.UserUserUpdateRequest) {
	o.Body = body
}

// WithID adds the id to the resonate user update user params
func (o *ResonateUserUpdateUserParams) WithID(id string) *ResonateUserUpdateUserParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the resonate user update user params
func (o *ResonateUserUpdateUserParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ResonateUserUpdateUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
