// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_storage_capacity

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
)

// NewPcloudStoragecapacityPoolsGetallParams creates a new PcloudStoragecapacityPoolsGetallParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPcloudStoragecapacityPoolsGetallParams() *PcloudStoragecapacityPoolsGetallParams {
	return &PcloudStoragecapacityPoolsGetallParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudStoragecapacityPoolsGetallParamsWithTimeout creates a new PcloudStoragecapacityPoolsGetallParams object
// with the ability to set a timeout on a request.
func NewPcloudStoragecapacityPoolsGetallParamsWithTimeout(timeout time.Duration) *PcloudStoragecapacityPoolsGetallParams {
	return &PcloudStoragecapacityPoolsGetallParams{
		timeout: timeout,
	}
}

// NewPcloudStoragecapacityPoolsGetallParamsWithContext creates a new PcloudStoragecapacityPoolsGetallParams object
// with the ability to set a context for a request.
func NewPcloudStoragecapacityPoolsGetallParamsWithContext(ctx context.Context) *PcloudStoragecapacityPoolsGetallParams {
	return &PcloudStoragecapacityPoolsGetallParams{
		Context: ctx,
	}
}

// NewPcloudStoragecapacityPoolsGetallParamsWithHTTPClient creates a new PcloudStoragecapacityPoolsGetallParams object
// with the ability to set a custom HTTPClient for a request.
func NewPcloudStoragecapacityPoolsGetallParamsWithHTTPClient(client *http.Client) *PcloudStoragecapacityPoolsGetallParams {
	return &PcloudStoragecapacityPoolsGetallParams{
		HTTPClient: client,
	}
}

/* PcloudStoragecapacityPoolsGetallParams contains all the parameters to send to the API endpoint
   for the pcloud storagecapacity pools getall operation.

   Typically these are written to a http.Request.
*/
type PcloudStoragecapacityPoolsGetallParams struct {

	/* CloudInstanceID.

	   Cloud Instance ID of a PCloud Instance
	*/
	CloudInstanceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pcloud storagecapacity pools getall params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudStoragecapacityPoolsGetallParams) WithDefaults() *PcloudStoragecapacityPoolsGetallParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pcloud storagecapacity pools getall params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudStoragecapacityPoolsGetallParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pcloud storagecapacity pools getall params
func (o *PcloudStoragecapacityPoolsGetallParams) WithTimeout(timeout time.Duration) *PcloudStoragecapacityPoolsGetallParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud storagecapacity pools getall params
func (o *PcloudStoragecapacityPoolsGetallParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud storagecapacity pools getall params
func (o *PcloudStoragecapacityPoolsGetallParams) WithContext(ctx context.Context) *PcloudStoragecapacityPoolsGetallParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud storagecapacity pools getall params
func (o *PcloudStoragecapacityPoolsGetallParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud storagecapacity pools getall params
func (o *PcloudStoragecapacityPoolsGetallParams) WithHTTPClient(client *http.Client) *PcloudStoragecapacityPoolsGetallParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud storagecapacity pools getall params
func (o *PcloudStoragecapacityPoolsGetallParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudInstanceID adds the cloudInstanceID to the pcloud storagecapacity pools getall params
func (o *PcloudStoragecapacityPoolsGetallParams) WithCloudInstanceID(cloudInstanceID string) *PcloudStoragecapacityPoolsGetallParams {
	o.SetCloudInstanceID(cloudInstanceID)
	return o
}

// SetCloudInstanceID adds the cloudInstanceId to the pcloud storagecapacity pools getall params
func (o *PcloudStoragecapacityPoolsGetallParams) SetCloudInstanceID(cloudInstanceID string) {
	o.CloudInstanceID = cloudInstanceID
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudStoragecapacityPoolsGetallParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cloud_instance_id
	if err := r.SetPathParam("cloud_instance_id", o.CloudInstanceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
