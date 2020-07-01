// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cccfr/gomete/models"
)

// GetUsersBarcodeBarcodeOKCode is the HTTP code returned for type GetUsersBarcodeBarcodeOK
const GetUsersBarcodeBarcodeOKCode int = 200

/*GetUsersBarcodeBarcodeOK user's details

swagger:response getUsersBarcodeBarcodeOK
*/
type GetUsersBarcodeBarcodeOK struct {

	/*
	  In: Body
	*/
	Payload *models.User `json:"body,omitempty"`
}

// NewGetUsersBarcodeBarcodeOK creates GetUsersBarcodeBarcodeOK with default headers values
func NewGetUsersBarcodeBarcodeOK() *GetUsersBarcodeBarcodeOK {

	return &GetUsersBarcodeBarcodeOK{}
}

// WithPayload adds the payload to the get users barcode barcode o k response
func (o *GetUsersBarcodeBarcodeOK) WithPayload(payload *models.User) *GetUsersBarcodeBarcodeOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get users barcode barcode o k response
func (o *GetUsersBarcodeBarcodeOK) SetPayload(payload *models.User) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUsersBarcodeBarcodeOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetUsersBarcodeBarcodeBadRequestCode is the HTTP code returned for type GetUsersBarcodeBarcodeBadRequest
const GetUsersBarcodeBarcodeBadRequestCode int = 400

/*GetUsersBarcodeBarcodeBadRequest no user associated to the given barcode

swagger:response getUsersBarcodeBarcodeBadRequest
*/
type GetUsersBarcodeBarcodeBadRequest struct {
}

// NewGetUsersBarcodeBarcodeBadRequest creates GetUsersBarcodeBarcodeBadRequest with default headers values
func NewGetUsersBarcodeBarcodeBadRequest() *GetUsersBarcodeBarcodeBadRequest {

	return &GetUsersBarcodeBarcodeBadRequest{}
}

// WriteResponse to the client
func (o *GetUsersBarcodeBarcodeBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// GetUsersBarcodeBarcodeNotFoundCode is the HTTP code returned for type GetUsersBarcodeBarcodeNotFound
const GetUsersBarcodeBarcodeNotFoundCode int = 404

/*GetUsersBarcodeBarcodeNotFound barcode not found

swagger:response getUsersBarcodeBarcodeNotFound
*/
type GetUsersBarcodeBarcodeNotFound struct {
}

// NewGetUsersBarcodeBarcodeNotFound creates GetUsersBarcodeBarcodeNotFound with default headers values
func NewGetUsersBarcodeBarcodeNotFound() *GetUsersBarcodeBarcodeNotFound {

	return &GetUsersBarcodeBarcodeNotFound{}
}

// WriteResponse to the client
func (o *GetUsersBarcodeBarcodeNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}