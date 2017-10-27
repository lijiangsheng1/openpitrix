// Code generated by go-swagger; DO NOT EDIT.

package apps

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

// NewDeleteAppsAppIDParams creates a new DeleteAppsAppIDParams object
// with the default values initialized.
func NewDeleteAppsAppIDParams() *DeleteAppsAppIDParams {
	var ()
	return &DeleteAppsAppIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAppsAppIDParamsWithTimeout creates a new DeleteAppsAppIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteAppsAppIDParamsWithTimeout(timeout time.Duration) *DeleteAppsAppIDParams {
	var ()
	return &DeleteAppsAppIDParams{

		timeout: timeout,
	}
}

// NewDeleteAppsAppIDParamsWithContext creates a new DeleteAppsAppIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteAppsAppIDParamsWithContext(ctx context.Context) *DeleteAppsAppIDParams {
	var ()
	return &DeleteAppsAppIDParams{

		Context: ctx,
	}
}

// NewDeleteAppsAppIDParamsWithHTTPClient creates a new DeleteAppsAppIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteAppsAppIDParamsWithHTTPClient(client *http.Client) *DeleteAppsAppIDParams {
	var ()
	return &DeleteAppsAppIDParams{
		HTTPClient: client,
	}
}

/*DeleteAppsAppIDParams contains all the parameters to send to the API endpoint
for the delete apps app ID operation typically these are written to a http.Request
*/
type DeleteAppsAppIDParams struct {

	/*AppID
	  The app's id

	*/
	AppID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete apps app ID params
func (o *DeleteAppsAppIDParams) WithTimeout(timeout time.Duration) *DeleteAppsAppIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete apps app ID params
func (o *DeleteAppsAppIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete apps app ID params
func (o *DeleteAppsAppIDParams) WithContext(ctx context.Context) *DeleteAppsAppIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete apps app ID params
func (o *DeleteAppsAppIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete apps app ID params
func (o *DeleteAppsAppIDParams) WithHTTPClient(client *http.Client) *DeleteAppsAppIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete apps app ID params
func (o *DeleteAppsAppIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppID adds the appID to the delete apps app ID params
func (o *DeleteAppsAppIDParams) WithAppID(appID string) *DeleteAppsAppIDParams {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the delete apps app ID params
func (o *DeleteAppsAppIDParams) SetAppID(appID string) {
	o.AppID = appID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAppsAppIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param appId
	if err := r.SetPathParam("appId", o.AppID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
