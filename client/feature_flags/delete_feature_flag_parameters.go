// Code generated by go-swagger; DO NOT EDIT.

package feature_flags

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

// NewDeleteFeatureFlagParams creates a new DeleteFeatureFlagParams object
// with the default values initialized.
func NewDeleteFeatureFlagParams() *DeleteFeatureFlagParams {
	var ()
	return &DeleteFeatureFlagParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteFeatureFlagParamsWithTimeout creates a new DeleteFeatureFlagParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteFeatureFlagParamsWithTimeout(timeout time.Duration) *DeleteFeatureFlagParams {
	var ()
	return &DeleteFeatureFlagParams{

		timeout: timeout,
	}
}

// NewDeleteFeatureFlagParamsWithContext creates a new DeleteFeatureFlagParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteFeatureFlagParamsWithContext(ctx context.Context) *DeleteFeatureFlagParams {
	var ()
	return &DeleteFeatureFlagParams{

		Context: ctx,
	}
}

// NewDeleteFeatureFlagParamsWithHTTPClient creates a new DeleteFeatureFlagParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteFeatureFlagParamsWithHTTPClient(client *http.Client) *DeleteFeatureFlagParams {
	var ()
	return &DeleteFeatureFlagParams{
		HTTPClient: client,
	}
}

/*DeleteFeatureFlagParams contains all the parameters to send to the API endpoint
for the delete feature flag operation typically these are written to a http.Request
*/
type DeleteFeatureFlagParams struct {

	/*FeatureFlagKey
	  The feature flag's key. The key identifies the flag in your code.

	*/
	FeatureFlagKey string
	/*ProjectKey
	  The project key, used to tie the flags together under one project so they can be managed together.

	*/
	ProjectKey string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete feature flag params
func (o *DeleteFeatureFlagParams) WithTimeout(timeout time.Duration) *DeleteFeatureFlagParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete feature flag params
func (o *DeleteFeatureFlagParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete feature flag params
func (o *DeleteFeatureFlagParams) WithContext(ctx context.Context) *DeleteFeatureFlagParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete feature flag params
func (o *DeleteFeatureFlagParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete feature flag params
func (o *DeleteFeatureFlagParams) WithHTTPClient(client *http.Client) *DeleteFeatureFlagParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete feature flag params
func (o *DeleteFeatureFlagParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFeatureFlagKey adds the featureFlagKey to the delete feature flag params
func (o *DeleteFeatureFlagParams) WithFeatureFlagKey(featureFlagKey string) *DeleteFeatureFlagParams {
	o.SetFeatureFlagKey(featureFlagKey)
	return o
}

// SetFeatureFlagKey adds the featureFlagKey to the delete feature flag params
func (o *DeleteFeatureFlagParams) SetFeatureFlagKey(featureFlagKey string) {
	o.FeatureFlagKey = featureFlagKey
}

// WithProjectKey adds the projectKey to the delete feature flag params
func (o *DeleteFeatureFlagParams) WithProjectKey(projectKey string) *DeleteFeatureFlagParams {
	o.SetProjectKey(projectKey)
	return o
}

// SetProjectKey adds the projectKey to the delete feature flag params
func (o *DeleteFeatureFlagParams) SetProjectKey(projectKey string) {
	o.ProjectKey = projectKey
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteFeatureFlagParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param featureFlagKey
	if err := r.SetPathParam("featureFlagKey", o.FeatureFlagKey); err != nil {
		return err
	}

	// path param projectKey
	if err := r.SetPathParam("projectKey", o.ProjectKey); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
