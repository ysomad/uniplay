// Code generated by go-swagger; DO NOT EDIT.

package player

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/ysomad/uniplay/internal/gen/swagger2/models"
)

// GetPlayerStatsOKCode is the HTTP code returned for type GetPlayerStatsOK
const GetPlayerStatsOKCode int = 200

/*
GetPlayerStatsOK OK

swagger:response getPlayerStatsOK
*/
type GetPlayerStatsOK struct {

	/*
	  In: Body
	*/
	Payload *models.PlayerStats `json:"body,omitempty"`
}

// NewGetPlayerStatsOK creates GetPlayerStatsOK with default headers values
func NewGetPlayerStatsOK() *GetPlayerStatsOK {

	return &GetPlayerStatsOK{}
}

// WithPayload adds the payload to the get player stats o k response
func (o *GetPlayerStatsOK) WithPayload(payload *models.PlayerStats) *GetPlayerStatsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get player stats o k response
func (o *GetPlayerStatsOK) SetPayload(payload *models.PlayerStats) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPlayerStatsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *GetPlayerStatsOK) GetPlayerStatsResponder() {}

// GetPlayerStatsBadRequestCode is the HTTP code returned for type GetPlayerStatsBadRequest
const GetPlayerStatsBadRequestCode int = 400

/*
GetPlayerStatsBadRequest Bad Request

swagger:response getPlayerStatsBadRequest
*/
type GetPlayerStatsBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetPlayerStatsBadRequest creates GetPlayerStatsBadRequest with default headers values
func NewGetPlayerStatsBadRequest() *GetPlayerStatsBadRequest {

	return &GetPlayerStatsBadRequest{}
}

// WithPayload adds the payload to the get player stats bad request response
func (o *GetPlayerStatsBadRequest) WithPayload(payload *models.Error) *GetPlayerStatsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get player stats bad request response
func (o *GetPlayerStatsBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPlayerStatsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *GetPlayerStatsBadRequest) GetPlayerStatsResponder() {}

// GetPlayerStatsNotFoundCode is the HTTP code returned for type GetPlayerStatsNotFound
const GetPlayerStatsNotFoundCode int = 404

/*
GetPlayerStatsNotFound Not Found

swagger:response getPlayerStatsNotFound
*/
type GetPlayerStatsNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetPlayerStatsNotFound creates GetPlayerStatsNotFound with default headers values
func NewGetPlayerStatsNotFound() *GetPlayerStatsNotFound {

	return &GetPlayerStatsNotFound{}
}

// WithPayload adds the payload to the get player stats not found response
func (o *GetPlayerStatsNotFound) WithPayload(payload *models.Error) *GetPlayerStatsNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get player stats not found response
func (o *GetPlayerStatsNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPlayerStatsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *GetPlayerStatsNotFound) GetPlayerStatsResponder() {}

// GetPlayerStatsInternalServerErrorCode is the HTTP code returned for type GetPlayerStatsInternalServerError
const GetPlayerStatsInternalServerErrorCode int = 500

/*
GetPlayerStatsInternalServerError Internal Server Error

swagger:response getPlayerStatsInternalServerError
*/
type GetPlayerStatsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetPlayerStatsInternalServerError creates GetPlayerStatsInternalServerError with default headers values
func NewGetPlayerStatsInternalServerError() *GetPlayerStatsInternalServerError {

	return &GetPlayerStatsInternalServerError{}
}

// WithPayload adds the payload to the get player stats internal server error response
func (o *GetPlayerStatsInternalServerError) WithPayload(payload *models.Error) *GetPlayerStatsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get player stats internal server error response
func (o *GetPlayerStatsInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPlayerStatsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *GetPlayerStatsInternalServerError) GetPlayerStatsResponder() {}

type GetPlayerStatsNotImplementedResponder struct {
	middleware.Responder
}

func (*GetPlayerStatsNotImplementedResponder) GetPlayerStatsResponder() {}

func GetPlayerStatsNotImplemented() GetPlayerStatsResponder {
	return &GetPlayerStatsNotImplementedResponder{
		middleware.NotImplemented(
			"operation authentication.GetPlayerStats has not yet been implemented",
		),
	}
}

type GetPlayerStatsResponder interface {
	middleware.Responder
	GetPlayerStatsResponder()
}