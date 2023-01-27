// Code generated by go-swagger; DO NOT EDIT.

package replay

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/ssssargsian/uniplay/internal/gen/swagger2/v1/models"
)

// UploadReplayOKCode is the HTTP code returned for type UploadReplayOK
const UploadReplayOKCode int = 200

/*
UploadReplayOK OK

swagger:response uploadReplayOK
*/
type UploadReplayOK struct {

	/*
	  In: Body
	*/
	Payload *models.UploadReplayResponse `json:"body,omitempty"`
}

// NewUploadReplayOK creates UploadReplayOK with default headers values
func NewUploadReplayOK() *UploadReplayOK {

	return &UploadReplayOK{}
}

// WithPayload adds the payload to the upload replay o k response
func (o *UploadReplayOK) WithPayload(payload *models.UploadReplayResponse) *UploadReplayOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the upload replay o k response
func (o *UploadReplayOK) SetPayload(payload *models.UploadReplayResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UploadReplayOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *UploadReplayOK) UploadReplayResponder() {}

// UploadReplayBadRequestCode is the HTTP code returned for type UploadReplayBadRequest
const UploadReplayBadRequestCode int = 400

/*
UploadReplayBadRequest Bad Request

swagger:response uploadReplayBadRequest
*/
type UploadReplayBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUploadReplayBadRequest creates UploadReplayBadRequest with default headers values
func NewUploadReplayBadRequest() *UploadReplayBadRequest {

	return &UploadReplayBadRequest{}
}

// WithPayload adds the payload to the upload replay bad request response
func (o *UploadReplayBadRequest) WithPayload(payload *models.Error) *UploadReplayBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the upload replay bad request response
func (o *UploadReplayBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UploadReplayBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *UploadReplayBadRequest) UploadReplayResponder() {}

// UploadReplayConflictCode is the HTTP code returned for type UploadReplayConflict
const UploadReplayConflictCode int = 409

/*
UploadReplayConflict Conflict

swagger:response uploadReplayConflict
*/
type UploadReplayConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUploadReplayConflict creates UploadReplayConflict with default headers values
func NewUploadReplayConflict() *UploadReplayConflict {

	return &UploadReplayConflict{}
}

// WithPayload adds the payload to the upload replay conflict response
func (o *UploadReplayConflict) WithPayload(payload *models.Error) *UploadReplayConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the upload replay conflict response
func (o *UploadReplayConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UploadReplayConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *UploadReplayConflict) UploadReplayResponder() {}

// UploadReplayInternalServerErrorCode is the HTTP code returned for type UploadReplayInternalServerError
const UploadReplayInternalServerErrorCode int = 500

/*
UploadReplayInternalServerError Internal Server Error

swagger:response uploadReplayInternalServerError
*/
type UploadReplayInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUploadReplayInternalServerError creates UploadReplayInternalServerError with default headers values
func NewUploadReplayInternalServerError() *UploadReplayInternalServerError {

	return &UploadReplayInternalServerError{}
}

// WithPayload adds the payload to the upload replay internal server error response
func (o *UploadReplayInternalServerError) WithPayload(payload *models.Error) *UploadReplayInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the upload replay internal server error response
func (o *UploadReplayInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UploadReplayInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *UploadReplayInternalServerError) UploadReplayResponder() {}

type UploadReplayNotImplementedResponder struct {
	middleware.Responder
}

func (*UploadReplayNotImplementedResponder) UploadReplayResponder() {}

func UploadReplayNotImplemented() UploadReplayResponder {
	return &UploadReplayNotImplementedResponder{
		middleware.NotImplemented(
			"operation authentication.UploadReplay has not yet been implemented",
		),
	}
}

type UploadReplayResponder interface {
	middleware.Responder
	UploadReplayResponder()
}
