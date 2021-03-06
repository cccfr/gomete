// Code generated by go-swagger; DO NOT EDIT.

package image

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cccfr/gomete/models"
)

// GetImagesIDOKCode is the HTTP code returned for type GetImagesIDOK
const GetImagesIDOKCode int = 200

/*GetImagesIDOK image metadata

swagger:response getImagesIdOK
*/
type GetImagesIDOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Image `json:"body,omitempty"`
}

// NewGetImagesIDOK creates GetImagesIDOK with default headers values
func NewGetImagesIDOK() *GetImagesIDOK {

	return &GetImagesIDOK{}
}

// WithPayload adds the payload to the get images Id o k response
func (o *GetImagesIDOK) WithPayload(payload []*models.Image) *GetImagesIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get images Id o k response
func (o *GetImagesIDOK) SetPayload(payload []*models.Image) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetImagesIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Image, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
