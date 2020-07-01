// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cccfr/gomete/models"
)

// PostProductsCreatedCode is the HTTP code returned for type PostProductsCreated
const PostProductsCreatedCode int = 201

/*PostProductsCreated created Product

swagger:response postProductsCreated
*/
type PostProductsCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Product `json:"body,omitempty"`
}

// NewPostProductsCreated creates PostProductsCreated with default headers values
func NewPostProductsCreated() *PostProductsCreated {

	return &PostProductsCreated{}
}

// WithPayload adds the payload to the post products created response
func (o *PostProductsCreated) WithPayload(payload *models.Product) *PostProductsCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post products created response
func (o *PostProductsCreated) SetPayload(payload *models.Product) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostProductsCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostProductsBadRequestCode is the HTTP code returned for type PostProductsBadRequest
const PostProductsBadRequestCode int = 400

/*PostProductsBadRequest invalid input

swagger:response postProductsBadRequest
*/
type PostProductsBadRequest struct {
}

// NewPostProductsBadRequest creates PostProductsBadRequest with default headers values
func NewPostProductsBadRequest() *PostProductsBadRequest {

	return &PostProductsBadRequest{}
}

// WriteResponse to the client
func (o *PostProductsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// PostProductsConflictCode is the HTTP code returned for type PostProductsConflict
const PostProductsConflictCode int = 409

/*PostProductsConflict a product with the same name already exists

swagger:response postProductsConflict
*/
type PostProductsConflict struct {
}

// NewPostProductsConflict creates PostProductsConflict with default headers values
func NewPostProductsConflict() *PostProductsConflict {

	return &PostProductsConflict{}
}

// WriteResponse to the client
func (o *PostProductsConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(409)
}

// PostProductsInternalServerErrorCode is the HTTP code returned for type PostProductsInternalServerError
const PostProductsInternalServerErrorCode int = 500

/*PostProductsInternalServerError server error

swagger:response postProductsInternalServerError
*/
type PostProductsInternalServerError struct {
}

// NewPostProductsInternalServerError creates PostProductsInternalServerError with default headers values
func NewPostProductsInternalServerError() *PostProductsInternalServerError {

	return &PostProductsInternalServerError{}
}

// WriteResponse to the client
func (o *PostProductsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
