// Code generated by go-swagger; DO NOT EDIT.

package denominations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cccfr/gomete/models"
)

// PatchDenominationsIDOKCode is the HTTP code returned for type PatchDenominationsIDOK
const PatchDenominationsIDOKCode int = 200

/*PatchDenominationsIDOK details of updated denomination

swagger:response patchDenominationsIdOK
*/
type PatchDenominationsIDOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Denomination `json:"body,omitempty"`
}

// NewPatchDenominationsIDOK creates PatchDenominationsIDOK with default headers values
func NewPatchDenominationsIDOK() *PatchDenominationsIDOK {

	return &PatchDenominationsIDOK{}
}

// WithPayload adds the payload to the patch denominations Id o k response
func (o *PatchDenominationsIDOK) WithPayload(payload []*models.Denomination) *PatchDenominationsIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch denominations Id o k response
func (o *PatchDenominationsIDOK) SetPayload(payload []*models.Denomination) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchDenominationsIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Denomination, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// PatchDenominationsIDBadRequestCode is the HTTP code returned for type PatchDenominationsIDBadRequest
const PatchDenominationsIDBadRequestCode int = 400

/*PatchDenominationsIDBadRequest bad request

swagger:response patchDenominationsIdBadRequest
*/
type PatchDenominationsIDBadRequest struct {
}

// NewPatchDenominationsIDBadRequest creates PatchDenominationsIDBadRequest with default headers values
func NewPatchDenominationsIDBadRequest() *PatchDenominationsIDBadRequest {

	return &PatchDenominationsIDBadRequest{}
}

// WriteResponse to the client
func (o *PatchDenominationsIDBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// PatchDenominationsIDNotFoundCode is the HTTP code returned for type PatchDenominationsIDNotFound
const PatchDenominationsIDNotFoundCode int = 404

/*PatchDenominationsIDNotFound id not existent

swagger:response patchDenominationsIdNotFound
*/
type PatchDenominationsIDNotFound struct {
}

// NewPatchDenominationsIDNotFound creates PatchDenominationsIDNotFound with default headers values
func NewPatchDenominationsIDNotFound() *PatchDenominationsIDNotFound {

	return &PatchDenominationsIDNotFound{}
}

// WriteResponse to the client
func (o *PatchDenominationsIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// PatchDenominationsIDConflictCode is the HTTP code returned for type PatchDenominationsIDConflict
const PatchDenominationsIDConflictCode int = 409

/*PatchDenominationsIDConflict a denomination with this value already exists

swagger:response patchDenominationsIdConflict
*/
type PatchDenominationsIDConflict struct {
}

// NewPatchDenominationsIDConflict creates PatchDenominationsIDConflict with default headers values
func NewPatchDenominationsIDConflict() *PatchDenominationsIDConflict {

	return &PatchDenominationsIDConflict{}
}

// WriteResponse to the client
func (o *PatchDenominationsIDConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(409)
}

// PatchDenominationsIDInternalServerErrorCode is the HTTP code returned for type PatchDenominationsIDInternalServerError
const PatchDenominationsIDInternalServerErrorCode int = 500

/*PatchDenominationsIDInternalServerError server error

swagger:response patchDenominationsIdInternalServerError
*/
type PatchDenominationsIDInternalServerError struct {
}

// NewPatchDenominationsIDInternalServerError creates PatchDenominationsIDInternalServerError with default headers values
func NewPatchDenominationsIDInternalServerError() *PatchDenominationsIDInternalServerError {

	return &PatchDenominationsIDInternalServerError{}
}

// WriteResponse to the client
func (o *PatchDenominationsIDInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
