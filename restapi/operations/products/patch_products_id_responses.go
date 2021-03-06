// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cccfr/gomete/models"
)

// PatchProductsIDOKCode is the HTTP code returned for type PatchProductsIDOK
const PatchProductsIDOKCode int = 200

/*PatchProductsIDOK edited product

swagger:response patchProductsIdOK
*/
type PatchProductsIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.Product `json:"body,omitempty"`
}

// NewPatchProductsIDOK creates PatchProductsIDOK with default headers values
func NewPatchProductsIDOK() *PatchProductsIDOK {

	return &PatchProductsIDOK{}
}

// WithPayload adds the payload to the patch products Id o k response
func (o *PatchProductsIDOK) WithPayload(payload *models.Product) *PatchProductsIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch products Id o k response
func (o *PatchProductsIDOK) SetPayload(payload *models.Product) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchProductsIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchProductsIDBadRequestCode is the HTTP code returned for type PatchProductsIDBadRequest
const PatchProductsIDBadRequestCode int = 400

/*PatchProductsIDBadRequest invalid input

swagger:response patchProductsIdBadRequest
*/
type PatchProductsIDBadRequest struct {
}

// NewPatchProductsIDBadRequest creates PatchProductsIDBadRequest with default headers values
func NewPatchProductsIDBadRequest() *PatchProductsIDBadRequest {

	return &PatchProductsIDBadRequest{}
}

// WriteResponse to the client
func (o *PatchProductsIDBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// PatchProductsIDNotFoundCode is the HTTP code returned for type PatchProductsIDNotFound
const PatchProductsIDNotFoundCode int = 404

/*PatchProductsIDNotFound id not existent

swagger:response patchProductsIdNotFound
*/
type PatchProductsIDNotFound struct {
}

// NewPatchProductsIDNotFound creates PatchProductsIDNotFound with default headers values
func NewPatchProductsIDNotFound() *PatchProductsIDNotFound {

	return &PatchProductsIDNotFound{}
}

// WriteResponse to the client
func (o *PatchProductsIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// PatchProductsIDInternalServerErrorCode is the HTTP code returned for type PatchProductsIDInternalServerError
const PatchProductsIDInternalServerErrorCode int = 500

/*PatchProductsIDInternalServerError server error

swagger:response patchProductsIdInternalServerError
*/
type PatchProductsIDInternalServerError struct {
}

// NewPatchProductsIDInternalServerError creates PatchProductsIDInternalServerError with default headers values
func NewPatchProductsIDInternalServerError() *PatchProductsIDInternalServerError {

	return &PatchProductsIDInternalServerError{}
}

// WriteResponse to the client
func (o *PatchProductsIDInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
