// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_v_p_n_policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudIkepoliciesGetallReader is a Reader for the PcloudIkepoliciesGetall structure.
type PcloudIkepoliciesGetallReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudIkepoliciesGetallReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudIkepoliciesGetallOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudIkepoliciesGetallBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudIkepoliciesGetallUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudIkepoliciesGetallForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudIkepoliciesGetallNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudIkepoliciesGetallInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPcloudIkepoliciesGetallOK creates a PcloudIkepoliciesGetallOK with default headers values
func NewPcloudIkepoliciesGetallOK() *PcloudIkepoliciesGetallOK {
	return &PcloudIkepoliciesGetallOK{}
}

/* PcloudIkepoliciesGetallOK describes a response with status code 200, with default header values.

OK
*/
type PcloudIkepoliciesGetallOK struct {
	Payload *models.IKEPolicies
}

func (o *PcloudIkepoliciesGetallOK) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/ike-policies][%d] pcloudIkepoliciesGetallOK  %+v", 200, o.Payload)
}
func (o *PcloudIkepoliciesGetallOK) GetPayload() *models.IKEPolicies {
	return o.Payload
}

func (o *PcloudIkepoliciesGetallOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IKEPolicies)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudIkepoliciesGetallBadRequest creates a PcloudIkepoliciesGetallBadRequest with default headers values
func NewPcloudIkepoliciesGetallBadRequest() *PcloudIkepoliciesGetallBadRequest {
	return &PcloudIkepoliciesGetallBadRequest{}
}

/* PcloudIkepoliciesGetallBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudIkepoliciesGetallBadRequest struct {
	Payload *models.Error
}

func (o *PcloudIkepoliciesGetallBadRequest) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/ike-policies][%d] pcloudIkepoliciesGetallBadRequest  %+v", 400, o.Payload)
}
func (o *PcloudIkepoliciesGetallBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudIkepoliciesGetallBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudIkepoliciesGetallUnauthorized creates a PcloudIkepoliciesGetallUnauthorized with default headers values
func NewPcloudIkepoliciesGetallUnauthorized() *PcloudIkepoliciesGetallUnauthorized {
	return &PcloudIkepoliciesGetallUnauthorized{}
}

/* PcloudIkepoliciesGetallUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudIkepoliciesGetallUnauthorized struct {
	Payload *models.Error
}

func (o *PcloudIkepoliciesGetallUnauthorized) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/ike-policies][%d] pcloudIkepoliciesGetallUnauthorized  %+v", 401, o.Payload)
}
func (o *PcloudIkepoliciesGetallUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudIkepoliciesGetallUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudIkepoliciesGetallForbidden creates a PcloudIkepoliciesGetallForbidden with default headers values
func NewPcloudIkepoliciesGetallForbidden() *PcloudIkepoliciesGetallForbidden {
	return &PcloudIkepoliciesGetallForbidden{}
}

/* PcloudIkepoliciesGetallForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudIkepoliciesGetallForbidden struct {
	Payload *models.Error
}

func (o *PcloudIkepoliciesGetallForbidden) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/ike-policies][%d] pcloudIkepoliciesGetallForbidden  %+v", 403, o.Payload)
}
func (o *PcloudIkepoliciesGetallForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudIkepoliciesGetallForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudIkepoliciesGetallNotFound creates a PcloudIkepoliciesGetallNotFound with default headers values
func NewPcloudIkepoliciesGetallNotFound() *PcloudIkepoliciesGetallNotFound {
	return &PcloudIkepoliciesGetallNotFound{}
}

/* PcloudIkepoliciesGetallNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudIkepoliciesGetallNotFound struct {
	Payload *models.Error
}

func (o *PcloudIkepoliciesGetallNotFound) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/ike-policies][%d] pcloudIkepoliciesGetallNotFound  %+v", 404, o.Payload)
}
func (o *PcloudIkepoliciesGetallNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudIkepoliciesGetallNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudIkepoliciesGetallInternalServerError creates a PcloudIkepoliciesGetallInternalServerError with default headers values
func NewPcloudIkepoliciesGetallInternalServerError() *PcloudIkepoliciesGetallInternalServerError {
	return &PcloudIkepoliciesGetallInternalServerError{}
}

/* PcloudIkepoliciesGetallInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudIkepoliciesGetallInternalServerError struct {
	Payload *models.Error
}

func (o *PcloudIkepoliciesGetallInternalServerError) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/ike-policies][%d] pcloudIkepoliciesGetallInternalServerError  %+v", 500, o.Payload)
}
func (o *PcloudIkepoliciesGetallInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudIkepoliciesGetallInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
