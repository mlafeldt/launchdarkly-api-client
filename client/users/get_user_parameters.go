// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetUserParams creates a new GetUserParams object
// with the default values initialized.
func NewGetUserParams() *GetUserParams {
	var ()
	return &GetUserParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUserParamsWithTimeout creates a new GetUserParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUserParamsWithTimeout(timeout time.Duration) *GetUserParams {
	var ()
	return &GetUserParams{

		timeout: timeout,
	}
}

// NewGetUserParamsWithContext creates a new GetUserParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetUserParamsWithContext(ctx context.Context) *GetUserParams {
	var ()
	return &GetUserParams{

		Context: ctx,
	}
}

// NewGetUserParamsWithHTTPClient creates a new GetUserParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetUserParamsWithHTTPClient(client *http.Client) *GetUserParams {
	var ()
	return &GetUserParams{
		HTTPClient: client,
	}
}

/*GetUserParams contains all the parameters to send to the API endpoint
for the get user operation typically these are written to a http.Request
*/
type GetUserParams struct {

	/*EnvironmentKey
	  The environment key, used to tie together flag configuration and users under one environment so they can be managed together.

	*/
	EnvironmentKey string
	/*ProjectKey
	  The project key, used to tie the flags together under one project so they can be managed together.

	*/
	ProjectKey string
	/*UserKey
	  The user's key.

	*/
	UserKey string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get user params
func (o *GetUserParams) WithTimeout(timeout time.Duration) *GetUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get user params
func (o *GetUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get user params
func (o *GetUserParams) WithContext(ctx context.Context) *GetUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get user params
func (o *GetUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get user params
func (o *GetUserParams) WithHTTPClient(client *http.Client) *GetUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get user params
func (o *GetUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnvironmentKey adds the environmentKey to the get user params
func (o *GetUserParams) WithEnvironmentKey(environmentKey string) *GetUserParams {
	o.SetEnvironmentKey(environmentKey)
	return o
}

// SetEnvironmentKey adds the environmentKey to the get user params
func (o *GetUserParams) SetEnvironmentKey(environmentKey string) {
	o.EnvironmentKey = environmentKey
}

// WithProjectKey adds the projectKey to the get user params
func (o *GetUserParams) WithProjectKey(projectKey string) *GetUserParams {
	o.SetProjectKey(projectKey)
	return o
}

// SetProjectKey adds the projectKey to the get user params
func (o *GetUserParams) SetProjectKey(projectKey string) {
	o.ProjectKey = projectKey
}

// WithUserKey adds the userKey to the get user params
func (o *GetUserParams) WithUserKey(userKey string) *GetUserParams {
	o.SetUserKey(userKey)
	return o
}

// SetUserKey adds the userKey to the get user params
func (o *GetUserParams) SetUserKey(userKey string) {
	o.UserKey = userKey
}

// WriteToRequest writes these params to a swagger request
func (o *GetUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param environmentKey
	if err := r.SetPathParam("environmentKey", o.EnvironmentKey); err != nil {
		return err
	}

	// path param projectKey
	if err := r.SetPathParam("projectKey", o.ProjectKey); err != nil {
		return err
	}

	// path param userKey
	if err := r.SetPathParam("userKey", o.UserKey); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
