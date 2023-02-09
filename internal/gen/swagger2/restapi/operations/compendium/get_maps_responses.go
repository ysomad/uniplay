// Code generated by go-swagger; DO NOT EDIT.

package compendium

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/ysomad/uniplay/internal/gen/swagger2/models"
)

// GetMapsOKCode is the HTTP code returned for type GetMapsOK
const GetMapsOKCode int = 200

/*
GetMapsOK OK

swagger:response getMapsOK
*/
type GetMapsOK struct {

	/*
	  In: Body
	*/
	Payload models.MapList `json:"body,omitempty"`
}

// NewGetMapsOK creates GetMapsOK with default headers values
func NewGetMapsOK() *GetMapsOK {

	return &GetMapsOK{}
}

// WithPayload adds the payload to the get maps o k response
func (o *GetMapsOK) WithPayload(payload models.MapList) *GetMapsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get maps o k response
func (o *GetMapsOK) SetPayload(payload models.MapList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMapsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.MapList{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

func (o *GetMapsOK) GetMapsResponder() {}

// GetMapsInternalServerErrorCode is the HTTP code returned for type GetMapsInternalServerError
const GetMapsInternalServerErrorCode int = 500

/*
GetMapsInternalServerError Internal Server Error

swagger:response getMapsInternalServerError
*/
type GetMapsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetMapsInternalServerError creates GetMapsInternalServerError with default headers values
func NewGetMapsInternalServerError() *GetMapsInternalServerError {

	return &GetMapsInternalServerError{}
}

// WithPayload adds the payload to the get maps internal server error response
func (o *GetMapsInternalServerError) WithPayload(payload *models.Error) *GetMapsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get maps internal server error response
func (o *GetMapsInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMapsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *GetMapsInternalServerError) GetMapsResponder() {}

type GetMapsNotImplementedResponder struct {
	middleware.Responder
}

func (*GetMapsNotImplementedResponder) GetMapsResponder() {}

func GetMapsNotImplemented() GetMapsResponder {
	return &GetMapsNotImplementedResponder{
		middleware.NotImplemented(
			"operation authentication.GetMaps has not yet been implemented",
		),
	}
}

type GetMapsResponder interface {
	middleware.Responder
	GetMapsResponder()
}
