// Code generated by go-swagger; DO NOT EDIT.

package webhooks

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

	models "github.com/mlafeldt/go-launchdarkly/models"
)

// NewPatchWebhookParams creates a new PatchWebhookParams object
// with the default values initialized.
func NewPatchWebhookParams() *PatchWebhookParams {
	var ()
	return &PatchWebhookParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchWebhookParamsWithTimeout creates a new PatchWebhookParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchWebhookParamsWithTimeout(timeout time.Duration) *PatchWebhookParams {
	var ()
	return &PatchWebhookParams{

		timeout: timeout,
	}
}

// NewPatchWebhookParamsWithContext creates a new PatchWebhookParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchWebhookParamsWithContext(ctx context.Context) *PatchWebhookParams {
	var ()
	return &PatchWebhookParams{

		Context: ctx,
	}
}

// NewPatchWebhookParamsWithHTTPClient creates a new PatchWebhookParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchWebhookParamsWithHTTPClient(client *http.Client) *PatchWebhookParams {
	var ()
	return &PatchWebhookParams{
		HTTPClient: client,
	}
}

/*PatchWebhookParams contains all the parameters to send to the API endpoint
for the patch webhook operation typically these are written to a http.Request
*/
type PatchWebhookParams struct {

	/*PatchDelta
	  Requires a JSON Patch representation of the desired changes to the project. 'http://jsonpatch.com/'

	*/
	PatchDelta []*models.PatchOperation
	/*ResourceID
	  The resource ID.

	*/
	ResourceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch webhook params
func (o *PatchWebhookParams) WithTimeout(timeout time.Duration) *PatchWebhookParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch webhook params
func (o *PatchWebhookParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch webhook params
func (o *PatchWebhookParams) WithContext(ctx context.Context) *PatchWebhookParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch webhook params
func (o *PatchWebhookParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch webhook params
func (o *PatchWebhookParams) WithHTTPClient(client *http.Client) *PatchWebhookParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch webhook params
func (o *PatchWebhookParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPatchDelta adds the patchDelta to the patch webhook params
func (o *PatchWebhookParams) WithPatchDelta(patchDelta []*models.PatchOperation) *PatchWebhookParams {
	o.SetPatchDelta(patchDelta)
	return o
}

// SetPatchDelta adds the patchDelta to the patch webhook params
func (o *PatchWebhookParams) SetPatchDelta(patchDelta []*models.PatchOperation) {
	o.PatchDelta = patchDelta
}

// WithResourceID adds the resourceID to the patch webhook params
func (o *PatchWebhookParams) WithResourceID(resourceID string) *PatchWebhookParams {
	o.SetResourceID(resourceID)
	return o
}

// SetResourceID adds the resourceId to the patch webhook params
func (o *PatchWebhookParams) SetResourceID(resourceID string) {
	o.ResourceID = resourceID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchWebhookParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.PatchDelta != nil {
		if err := r.SetBodyParam(o.PatchDelta); err != nil {
			return err
		}
	}

	// path param resourceId
	if err := r.SetPathParam("resourceId", o.ResourceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
