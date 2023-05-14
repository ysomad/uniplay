// Code generated by go-swagger; DO NOT EDIT.

package team

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/ysomad/uniplay/internal/gen/swagger2/models"
)

// SetTeamCaptainNoContentCode is the HTTP code returned for type SetTeamCaptainNoContent
const SetTeamCaptainNoContentCode int = 204

/*
SetTeamCaptainNoContent No Content

swagger:response setTeamCaptainNoContent
*/
type SetTeamCaptainNoContent struct {
}

// NewSetTeamCaptainNoContent creates SetTeamCaptainNoContent with default headers values
func NewSetTeamCaptainNoContent() *SetTeamCaptainNoContent {

	return &SetTeamCaptainNoContent{}
}

// WriteResponse to the client
func (o *SetTeamCaptainNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

func (o *SetTeamCaptainNoContent) SetTeamCaptainResponder() {}

// SetTeamCaptainInternalServerErrorCode is the HTTP code returned for type SetTeamCaptainInternalServerError
const SetTeamCaptainInternalServerErrorCode int = 500

/*
SetTeamCaptainInternalServerError Internal Server Error

swagger:response setTeamCaptainInternalServerError
*/
type SetTeamCaptainInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewSetTeamCaptainInternalServerError creates SetTeamCaptainInternalServerError with default headers values
func NewSetTeamCaptainInternalServerError() *SetTeamCaptainInternalServerError {

	return &SetTeamCaptainInternalServerError{}
}

// WithPayload adds the payload to the set team captain internal server error response
func (o *SetTeamCaptainInternalServerError) WithPayload(payload *models.Error) *SetTeamCaptainInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the set team captain internal server error response
func (o *SetTeamCaptainInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SetTeamCaptainInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *SetTeamCaptainInternalServerError) SetTeamCaptainResponder() {}

type SetTeamCaptainNotImplementedResponder struct {
	middleware.Responder
}

func (*SetTeamCaptainNotImplementedResponder) SetTeamCaptainResponder() {}

func SetTeamCaptainNotImplemented() SetTeamCaptainResponder {
	return &SetTeamCaptainNotImplementedResponder{
		middleware.NotImplemented(
			"operation authentication.SetTeamCaptain has not yet been implemented",
		),
	}
}

type SetTeamCaptainResponder interface {
	middleware.Responder
	SetTeamCaptainResponder()
}
