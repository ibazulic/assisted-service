// Code generated by go-swagger; DO NOT EDIT.

package operators

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openshift/assisted-service/models"
)

// V2ReportMonitoredOperatorStatusOKCode is the HTTP code returned for type V2ReportMonitoredOperatorStatusOK
const V2ReportMonitoredOperatorStatusOKCode int = 200

/*V2ReportMonitoredOperatorStatusOK Success.

swagger:response v2ReportMonitoredOperatorStatusOK
*/
type V2ReportMonitoredOperatorStatusOK struct {
}

// NewV2ReportMonitoredOperatorStatusOK creates V2ReportMonitoredOperatorStatusOK with default headers values
func NewV2ReportMonitoredOperatorStatusOK() *V2ReportMonitoredOperatorStatusOK {

	return &V2ReportMonitoredOperatorStatusOK{}
}

// WriteResponse to the client
func (o *V2ReportMonitoredOperatorStatusOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// V2ReportMonitoredOperatorStatusBadRequestCode is the HTTP code returned for type V2ReportMonitoredOperatorStatusBadRequest
const V2ReportMonitoredOperatorStatusBadRequestCode int = 400

/*V2ReportMonitoredOperatorStatusBadRequest Bad Request

swagger:response v2ReportMonitoredOperatorStatusBadRequest
*/
type V2ReportMonitoredOperatorStatusBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewV2ReportMonitoredOperatorStatusBadRequest creates V2ReportMonitoredOperatorStatusBadRequest with default headers values
func NewV2ReportMonitoredOperatorStatusBadRequest() *V2ReportMonitoredOperatorStatusBadRequest {

	return &V2ReportMonitoredOperatorStatusBadRequest{}
}

// WithPayload adds the payload to the v2 report monitored operator status bad request response
func (o *V2ReportMonitoredOperatorStatusBadRequest) WithPayload(payload *models.Error) *V2ReportMonitoredOperatorStatusBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 report monitored operator status bad request response
func (o *V2ReportMonitoredOperatorStatusBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2ReportMonitoredOperatorStatusBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2ReportMonitoredOperatorStatusUnauthorizedCode is the HTTP code returned for type V2ReportMonitoredOperatorStatusUnauthorized
const V2ReportMonitoredOperatorStatusUnauthorizedCode int = 401

/*V2ReportMonitoredOperatorStatusUnauthorized Unauthorized.

swagger:response v2ReportMonitoredOperatorStatusUnauthorized
*/
type V2ReportMonitoredOperatorStatusUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewV2ReportMonitoredOperatorStatusUnauthorized creates V2ReportMonitoredOperatorStatusUnauthorized with default headers values
func NewV2ReportMonitoredOperatorStatusUnauthorized() *V2ReportMonitoredOperatorStatusUnauthorized {

	return &V2ReportMonitoredOperatorStatusUnauthorized{}
}

// WithPayload adds the payload to the v2 report monitored operator status unauthorized response
func (o *V2ReportMonitoredOperatorStatusUnauthorized) WithPayload(payload *models.InfraError) *V2ReportMonitoredOperatorStatusUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 report monitored operator status unauthorized response
func (o *V2ReportMonitoredOperatorStatusUnauthorized) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2ReportMonitoredOperatorStatusUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2ReportMonitoredOperatorStatusForbiddenCode is the HTTP code returned for type V2ReportMonitoredOperatorStatusForbidden
const V2ReportMonitoredOperatorStatusForbiddenCode int = 403

/*V2ReportMonitoredOperatorStatusForbidden Forbidden.

swagger:response v2ReportMonitoredOperatorStatusForbidden
*/
type V2ReportMonitoredOperatorStatusForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewV2ReportMonitoredOperatorStatusForbidden creates V2ReportMonitoredOperatorStatusForbidden with default headers values
func NewV2ReportMonitoredOperatorStatusForbidden() *V2ReportMonitoredOperatorStatusForbidden {

	return &V2ReportMonitoredOperatorStatusForbidden{}
}

// WithPayload adds the payload to the v2 report monitored operator status forbidden response
func (o *V2ReportMonitoredOperatorStatusForbidden) WithPayload(payload *models.InfraError) *V2ReportMonitoredOperatorStatusForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 report monitored operator status forbidden response
func (o *V2ReportMonitoredOperatorStatusForbidden) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2ReportMonitoredOperatorStatusForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2ReportMonitoredOperatorStatusNotFoundCode is the HTTP code returned for type V2ReportMonitoredOperatorStatusNotFound
const V2ReportMonitoredOperatorStatusNotFoundCode int = 404

/*V2ReportMonitoredOperatorStatusNotFound Error.

swagger:response v2ReportMonitoredOperatorStatusNotFound
*/
type V2ReportMonitoredOperatorStatusNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewV2ReportMonitoredOperatorStatusNotFound creates V2ReportMonitoredOperatorStatusNotFound with default headers values
func NewV2ReportMonitoredOperatorStatusNotFound() *V2ReportMonitoredOperatorStatusNotFound {

	return &V2ReportMonitoredOperatorStatusNotFound{}
}

// WithPayload adds the payload to the v2 report monitored operator status not found response
func (o *V2ReportMonitoredOperatorStatusNotFound) WithPayload(payload *models.Error) *V2ReportMonitoredOperatorStatusNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 report monitored operator status not found response
func (o *V2ReportMonitoredOperatorStatusNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2ReportMonitoredOperatorStatusNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2ReportMonitoredOperatorStatusMethodNotAllowedCode is the HTTP code returned for type V2ReportMonitoredOperatorStatusMethodNotAllowed
const V2ReportMonitoredOperatorStatusMethodNotAllowedCode int = 405

/*V2ReportMonitoredOperatorStatusMethodNotAllowed Method Not Allowed.

swagger:response v2ReportMonitoredOperatorStatusMethodNotAllowed
*/
type V2ReportMonitoredOperatorStatusMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewV2ReportMonitoredOperatorStatusMethodNotAllowed creates V2ReportMonitoredOperatorStatusMethodNotAllowed with default headers values
func NewV2ReportMonitoredOperatorStatusMethodNotAllowed() *V2ReportMonitoredOperatorStatusMethodNotAllowed {

	return &V2ReportMonitoredOperatorStatusMethodNotAllowed{}
}

// WithPayload adds the payload to the v2 report monitored operator status method not allowed response
func (o *V2ReportMonitoredOperatorStatusMethodNotAllowed) WithPayload(payload *models.Error) *V2ReportMonitoredOperatorStatusMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 report monitored operator status method not allowed response
func (o *V2ReportMonitoredOperatorStatusMethodNotAllowed) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2ReportMonitoredOperatorStatusMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2ReportMonitoredOperatorStatusConflictCode is the HTTP code returned for type V2ReportMonitoredOperatorStatusConflict
const V2ReportMonitoredOperatorStatusConflictCode int = 409

/*V2ReportMonitoredOperatorStatusConflict Error.

swagger:response v2ReportMonitoredOperatorStatusConflict
*/
type V2ReportMonitoredOperatorStatusConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewV2ReportMonitoredOperatorStatusConflict creates V2ReportMonitoredOperatorStatusConflict with default headers values
func NewV2ReportMonitoredOperatorStatusConflict() *V2ReportMonitoredOperatorStatusConflict {

	return &V2ReportMonitoredOperatorStatusConflict{}
}

// WithPayload adds the payload to the v2 report monitored operator status conflict response
func (o *V2ReportMonitoredOperatorStatusConflict) WithPayload(payload *models.Error) *V2ReportMonitoredOperatorStatusConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 report monitored operator status conflict response
func (o *V2ReportMonitoredOperatorStatusConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2ReportMonitoredOperatorStatusConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2ReportMonitoredOperatorStatusInternalServerErrorCode is the HTTP code returned for type V2ReportMonitoredOperatorStatusInternalServerError
const V2ReportMonitoredOperatorStatusInternalServerErrorCode int = 500

/*V2ReportMonitoredOperatorStatusInternalServerError Error.

swagger:response v2ReportMonitoredOperatorStatusInternalServerError
*/
type V2ReportMonitoredOperatorStatusInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewV2ReportMonitoredOperatorStatusInternalServerError creates V2ReportMonitoredOperatorStatusInternalServerError with default headers values
func NewV2ReportMonitoredOperatorStatusInternalServerError() *V2ReportMonitoredOperatorStatusInternalServerError {

	return &V2ReportMonitoredOperatorStatusInternalServerError{}
}

// WithPayload adds the payload to the v2 report monitored operator status internal server error response
func (o *V2ReportMonitoredOperatorStatusInternalServerError) WithPayload(payload *models.Error) *V2ReportMonitoredOperatorStatusInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 report monitored operator status internal server error response
func (o *V2ReportMonitoredOperatorStatusInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2ReportMonitoredOperatorStatusInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2ReportMonitoredOperatorStatusServiceUnavailableCode is the HTTP code returned for type V2ReportMonitoredOperatorStatusServiceUnavailable
const V2ReportMonitoredOperatorStatusServiceUnavailableCode int = 503

/*V2ReportMonitoredOperatorStatusServiceUnavailable Unavailable.

swagger:response v2ReportMonitoredOperatorStatusServiceUnavailable
*/
type V2ReportMonitoredOperatorStatusServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewV2ReportMonitoredOperatorStatusServiceUnavailable creates V2ReportMonitoredOperatorStatusServiceUnavailable with default headers values
func NewV2ReportMonitoredOperatorStatusServiceUnavailable() *V2ReportMonitoredOperatorStatusServiceUnavailable {

	return &V2ReportMonitoredOperatorStatusServiceUnavailable{}
}

// WithPayload adds the payload to the v2 report monitored operator status service unavailable response
func (o *V2ReportMonitoredOperatorStatusServiceUnavailable) WithPayload(payload *models.Error) *V2ReportMonitoredOperatorStatusServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 report monitored operator status service unavailable response
func (o *V2ReportMonitoredOperatorStatusServiceUnavailable) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2ReportMonitoredOperatorStatusServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
