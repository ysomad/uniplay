// Code generated by go-swagger; DO NOT EDIT.

package institution

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/ysomad/uniplay/internal/gen/swagger2/models"
)

// GetInstitutionsOKCode is the HTTP code returned for type GetInstitutionsOK
const GetInstitutionsOKCode int = 200

/*
GetInstitutionsOK OK

swagger:response getInstitutionsOK
*/
type GetInstitutionsOK struct {

	/*
	  In: Body
	*/
	Payload models.WeaponList `json:"body,omitempty"`
}

// NewGetInstitutionsOK creates GetInstitutionsOK with default headers values
func NewGetInstitutionsOK() *GetInstitutionsOK {

	return &GetInstitutionsOK{}
}

// WithPayload adds the payload to the get institutions o k response
func (o *GetInstitutionsOK) WithPayload(payload models.WeaponList) *GetInstitutionsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get institutions o k response
func (o *GetInstitutionsOK) SetPayload(payload models.WeaponList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetInstitutionsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

func (o *GetInstitutionsOK) GetInstitutionsResponder() {}

// GetInstitutionsInternalServerErrorCode is the HTTP code returned for type GetInstitutionsInternalServerError
const GetInstitutionsInternalServerErrorCode int = 500

/*
GetInstitutionsInternalServerError Internal Server Error

swagger:response getInstitutionsInternalServerError
*/
type GetInstitutionsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetInstitutionsInternalServerError creates GetInstitutionsInternalServerError with default headers values
func NewGetInstitutionsInternalServerError() *GetInstitutionsInternalServerError {

	return &GetInstitutionsInternalServerError{}
}

// WithPayload adds the payload to the get institutions internal server error response
func (o *GetInstitutionsInternalServerError) WithPayload(payload *models.Error) *GetInstitutionsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get institutions internal server error response
func (o *GetInstitutionsInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetInstitutionsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *GetInstitutionsInternalServerError) GetInstitutionsResponder() {}

type GetInstitutionsNotImplementedResponder struct {
	middleware.Responder
}

func (*GetInstitutionsNotImplementedResponder) GetInstitutionsResponder() {}

func GetInstitutionsNotImplemented() GetInstitutionsResponder {
	return &GetInstitutionsNotImplementedResponder{
		middleware.NotImplemented(
			"operation authentication.GetInstitutions has not yet been implemented",
		),
	}
}

type GetInstitutionsResponder interface {
	middleware.Responder
	GetInstitutionsResponder()
}
