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

// GetTeamListOKCode is the HTTP code returned for type GetTeamListOK
const GetTeamListOKCode int = 200

/*
GetTeamListOK OK

swagger:response getTeamListOK
*/
type GetTeamListOK struct {

	/*
	  In: Body
	*/
	Payload models.WeaponList `json:"body,omitempty"`
}

// NewGetTeamListOK creates GetTeamListOK with default headers values
func NewGetTeamListOK() *GetTeamListOK {

	return &GetTeamListOK{}
}

// WithPayload adds the payload to the get team list o k response
func (o *GetTeamListOK) WithPayload(payload models.WeaponList) *GetTeamListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get team list o k response
func (o *GetTeamListOK) SetPayload(payload models.WeaponList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTeamListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.WeaponList{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

func (o *GetTeamListOK) GetTeamListResponder() {}

// GetTeamListInternalServerErrorCode is the HTTP code returned for type GetTeamListInternalServerError
const GetTeamListInternalServerErrorCode int = 500

/*
GetTeamListInternalServerError Internal Server Error

swagger:response getTeamListInternalServerError
*/
type GetTeamListInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetTeamListInternalServerError creates GetTeamListInternalServerError with default headers values
func NewGetTeamListInternalServerError() *GetTeamListInternalServerError {

	return &GetTeamListInternalServerError{}
}

// WithPayload adds the payload to the get team list internal server error response
func (o *GetTeamListInternalServerError) WithPayload(payload *models.Error) *GetTeamListInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get team list internal server error response
func (o *GetTeamListInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTeamListInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *GetTeamListInternalServerError) GetTeamListResponder() {}

type GetTeamListNotImplementedResponder struct {
	middleware.Responder
}

func (*GetTeamListNotImplementedResponder) GetTeamListResponder() {}

func GetTeamListNotImplemented() GetTeamListResponder {
	return &GetTeamListNotImplementedResponder{
		middleware.NotImplemented(
			"operation authentication.GetTeamList has not yet been implemented",
		),
	}
}

type GetTeamListResponder interface {
	middleware.Responder
	GetTeamListResponder()
}
