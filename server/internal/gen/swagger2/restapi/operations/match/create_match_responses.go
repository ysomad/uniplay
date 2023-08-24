// Code generated by go-swagger; DO NOT EDIT.

package match

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/ysomad/uniplay/internal/gen/swagger2/models"
)

// CreateMatchOKCode is the HTTP code returned for type CreateMatchOK
const CreateMatchOKCode int = 200

/*
CreateMatchOK OK

swagger:response createMatchOK
*/
type CreateMatchOK struct {

	/*
	  In: Body
	*/
	Payload *models.CreateMatchResponse `json:"body,omitempty"`
}

// NewCreateMatchOK creates CreateMatchOK with default headers values
func NewCreateMatchOK() *CreateMatchOK {

	return &CreateMatchOK{}
}

// WithPayload adds the payload to the create match o k response
func (o *CreateMatchOK) WithPayload(payload *models.CreateMatchResponse) *CreateMatchOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create match o k response
func (o *CreateMatchOK) SetPayload(payload *models.CreateMatchResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateMatchOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *CreateMatchOK) CreateMatchResponder() {}

// CreateMatchBadRequestCode is the HTTP code returned for type CreateMatchBadRequest
const CreateMatchBadRequestCode int = 400

/*
CreateMatchBadRequest Bad Request

swagger:response createMatchBadRequest
*/
type CreateMatchBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateMatchBadRequest creates CreateMatchBadRequest with default headers values
func NewCreateMatchBadRequest() *CreateMatchBadRequest {

	return &CreateMatchBadRequest{}
}

// WithPayload adds the payload to the create match bad request response
func (o *CreateMatchBadRequest) WithPayload(payload *models.Error) *CreateMatchBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create match bad request response
func (o *CreateMatchBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateMatchBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *CreateMatchBadRequest) CreateMatchResponder() {}

// CreateMatchConflictCode is the HTTP code returned for type CreateMatchConflict
const CreateMatchConflictCode int = 409

/*
CreateMatchConflict Conflict

swagger:response createMatchConflict
*/
type CreateMatchConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateMatchConflict creates CreateMatchConflict with default headers values
func NewCreateMatchConflict() *CreateMatchConflict {

	return &CreateMatchConflict{}
}

// WithPayload adds the payload to the create match conflict response
func (o *CreateMatchConflict) WithPayload(payload *models.Error) *CreateMatchConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create match conflict response
func (o *CreateMatchConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateMatchConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *CreateMatchConflict) CreateMatchResponder() {}

// CreateMatchInternalServerErrorCode is the HTTP code returned for type CreateMatchInternalServerError
const CreateMatchInternalServerErrorCode int = 500

/*
CreateMatchInternalServerError Internal Server Error

swagger:response createMatchInternalServerError
*/
type CreateMatchInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateMatchInternalServerError creates CreateMatchInternalServerError with default headers values
func NewCreateMatchInternalServerError() *CreateMatchInternalServerError {

	return &CreateMatchInternalServerError{}
}

// WithPayload adds the payload to the create match internal server error response
func (o *CreateMatchInternalServerError) WithPayload(payload *models.Error) *CreateMatchInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create match internal server error response
func (o *CreateMatchInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateMatchInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *CreateMatchInternalServerError) CreateMatchResponder() {}

type CreateMatchNotImplementedResponder struct {
	middleware.Responder
}

func (*CreateMatchNotImplementedResponder) CreateMatchResponder() {}

func CreateMatchNotImplemented() CreateMatchResponder {
	return &CreateMatchNotImplementedResponder{
		middleware.NotImplemented(
			"operation authentication.CreateMatch has not yet been implemented",
		),
	}
}

type CreateMatchResponder interface {
	middleware.Responder
	CreateMatchResponder()
}