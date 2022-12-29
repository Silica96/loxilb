// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/loxilb-io/loxilb/api/models"
)

// PostConfigParamsNoContentCode is the HTTP code returned for type PostConfigParamsNoContent
const PostConfigParamsNoContentCode int = 204

/*
PostConfigParamsNoContent OK

swagger:response postConfigParamsNoContent
*/
type PostConfigParamsNoContent struct {
}

// NewPostConfigParamsNoContent creates PostConfigParamsNoContent with default headers values
func NewPostConfigParamsNoContent() *PostConfigParamsNoContent {

	return &PostConfigParamsNoContent{}
}

// WriteResponse to the client
func (o *PostConfigParamsNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// PostConfigParamsBadRequestCode is the HTTP code returned for type PostConfigParamsBadRequest
const PostConfigParamsBadRequestCode int = 400

/*
PostConfigParamsBadRequest Malformed arguments for API call

swagger:response postConfigParamsBadRequest
*/
type PostConfigParamsBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostConfigParamsBadRequest creates PostConfigParamsBadRequest with default headers values
func NewPostConfigParamsBadRequest() *PostConfigParamsBadRequest {

	return &PostConfigParamsBadRequest{}
}

// WithPayload adds the payload to the post config params bad request response
func (o *PostConfigParamsBadRequest) WithPayload(payload *models.Error) *PostConfigParamsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post config params bad request response
func (o *PostConfigParamsBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostConfigParamsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostConfigParamsUnauthorizedCode is the HTTP code returned for type PostConfigParamsUnauthorized
const PostConfigParamsUnauthorizedCode int = 401

/*
PostConfigParamsUnauthorized Invalid authentication credentials

swagger:response postConfigParamsUnauthorized
*/
type PostConfigParamsUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostConfigParamsUnauthorized creates PostConfigParamsUnauthorized with default headers values
func NewPostConfigParamsUnauthorized() *PostConfigParamsUnauthorized {

	return &PostConfigParamsUnauthorized{}
}

// WithPayload adds the payload to the post config params unauthorized response
func (o *PostConfigParamsUnauthorized) WithPayload(payload *models.Error) *PostConfigParamsUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post config params unauthorized response
func (o *PostConfigParamsUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostConfigParamsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostConfigParamsForbiddenCode is the HTTP code returned for type PostConfigParamsForbidden
const PostConfigParamsForbiddenCode int = 403

/*
PostConfigParamsForbidden Capacity insufficient

swagger:response postConfigParamsForbidden
*/
type PostConfigParamsForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostConfigParamsForbidden creates PostConfigParamsForbidden with default headers values
func NewPostConfigParamsForbidden() *PostConfigParamsForbidden {

	return &PostConfigParamsForbidden{}
}

// WithPayload adds the payload to the post config params forbidden response
func (o *PostConfigParamsForbidden) WithPayload(payload *models.Error) *PostConfigParamsForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post config params forbidden response
func (o *PostConfigParamsForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostConfigParamsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostConfigParamsNotFoundCode is the HTTP code returned for type PostConfigParamsNotFound
const PostConfigParamsNotFoundCode int = 404

/*
PostConfigParamsNotFound Resource not found

swagger:response postConfigParamsNotFound
*/
type PostConfigParamsNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostConfigParamsNotFound creates PostConfigParamsNotFound with default headers values
func NewPostConfigParamsNotFound() *PostConfigParamsNotFound {

	return &PostConfigParamsNotFound{}
}

// WithPayload adds the payload to the post config params not found response
func (o *PostConfigParamsNotFound) WithPayload(payload *models.Error) *PostConfigParamsNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post config params not found response
func (o *PostConfigParamsNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostConfigParamsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostConfigParamsConflictCode is the HTTP code returned for type PostConfigParamsConflict
const PostConfigParamsConflictCode int = 409

/*
PostConfigParamsConflict Resource Conflict. VLAN already exists OR dependency VRF/VNET not found

swagger:response postConfigParamsConflict
*/
type PostConfigParamsConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostConfigParamsConflict creates PostConfigParamsConflict with default headers values
func NewPostConfigParamsConflict() *PostConfigParamsConflict {

	return &PostConfigParamsConflict{}
}

// WithPayload adds the payload to the post config params conflict response
func (o *PostConfigParamsConflict) WithPayload(payload *models.Error) *PostConfigParamsConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post config params conflict response
func (o *PostConfigParamsConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostConfigParamsConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostConfigParamsInternalServerErrorCode is the HTTP code returned for type PostConfigParamsInternalServerError
const PostConfigParamsInternalServerErrorCode int = 500

/*
PostConfigParamsInternalServerError Internal service error

swagger:response postConfigParamsInternalServerError
*/
type PostConfigParamsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostConfigParamsInternalServerError creates PostConfigParamsInternalServerError with default headers values
func NewPostConfigParamsInternalServerError() *PostConfigParamsInternalServerError {

	return &PostConfigParamsInternalServerError{}
}

// WithPayload adds the payload to the post config params internal server error response
func (o *PostConfigParamsInternalServerError) WithPayload(payload *models.Error) *PostConfigParamsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post config params internal server error response
func (o *PostConfigParamsInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostConfigParamsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostConfigParamsServiceUnavailableCode is the HTTP code returned for type PostConfigParamsServiceUnavailable
const PostConfigParamsServiceUnavailableCode int = 503

/*
PostConfigParamsServiceUnavailable Maintanence mode

swagger:response postConfigParamsServiceUnavailable
*/
type PostConfigParamsServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostConfigParamsServiceUnavailable creates PostConfigParamsServiceUnavailable with default headers values
func NewPostConfigParamsServiceUnavailable() *PostConfigParamsServiceUnavailable {

	return &PostConfigParamsServiceUnavailable{}
}

// WithPayload adds the payload to the post config params service unavailable response
func (o *PostConfigParamsServiceUnavailable) WithPayload(payload *models.Error) *PostConfigParamsServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post config params service unavailable response
func (o *PostConfigParamsServiceUnavailable) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostConfigParamsServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
