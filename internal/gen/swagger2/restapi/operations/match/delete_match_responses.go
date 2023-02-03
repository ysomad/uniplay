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

// DeleteMatchNoContentCode is the HTTP code returned for type DeleteMatchNoContent
const DeleteMatchNoContentCode int = 204

/*
DeleteMatchNoContent No Content

swagger:response deleteMatchNoContent
*/
type DeleteMatchNoContent struct {
}

// NewDeleteMatchNoContent creates DeleteMatchNoContent with default headers values
func NewDeleteMatchNoContent() *DeleteMatchNoContent {

	return &DeleteMatchNoContent{}
}

// WriteResponse to the client
func (o *DeleteMatchNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

func (o *DeleteMatchNoContent) DeleteMatchResponder() {}

// DeleteMatchNotFoundCode is the HTTP code returned for type DeleteMatchNotFound
const DeleteMatchNotFoundCode int = 404

/*
DeleteMatchNotFound Not Found

swagger:response deleteMatchNotFound
*/
type DeleteMatchNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteMatchNotFound creates DeleteMatchNotFound with default headers values
func NewDeleteMatchNotFound() *DeleteMatchNotFound {

	return &DeleteMatchNotFound{}
}

// WithPayload adds the payload to the delete match not found response
func (o *DeleteMatchNotFound) WithPayload(payload *models.Error) *DeleteMatchNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete match not found response
func (o *DeleteMatchNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteMatchNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *DeleteMatchNotFound) DeleteMatchResponder() {}

// DeleteMatchInternalServerErrorCode is the HTTP code returned for type DeleteMatchInternalServerError
const DeleteMatchInternalServerErrorCode int = 500

/*
DeleteMatchInternalServerError Internal Server Error

swagger:response deleteMatchInternalServerError
*/
type DeleteMatchInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteMatchInternalServerError creates DeleteMatchInternalServerError with default headers values
func NewDeleteMatchInternalServerError() *DeleteMatchInternalServerError {

	return &DeleteMatchInternalServerError{}
}

// WithPayload adds the payload to the delete match internal server error response
func (o *DeleteMatchInternalServerError) WithPayload(payload *models.Error) *DeleteMatchInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete match internal server error response
func (o *DeleteMatchInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteMatchInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *DeleteMatchInternalServerError) DeleteMatchResponder() {}

type DeleteMatchNotImplementedResponder struct {
	middleware.Responder
}

func (*DeleteMatchNotImplementedResponder) DeleteMatchResponder() {}

func DeleteMatchNotImplemented() DeleteMatchResponder {
	return &DeleteMatchNotImplementedResponder{
		middleware.NotImplemented(
			"operation authentication.DeleteMatch has not yet been implemented",
		),
	}
}

type DeleteMatchResponder interface {
	middleware.Responder
	DeleteMatchResponder()
}
