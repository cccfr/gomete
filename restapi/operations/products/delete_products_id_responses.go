// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// DeleteProductsIDOKCode is the HTTP code returned for type DeleteProductsIDOK
const DeleteProductsIDOKCode int = 200

/*DeleteProductsIDOK product deleted

swagger:response deleteProductsIdOK
*/
type DeleteProductsIDOK struct {
}

// NewDeleteProductsIDOK creates DeleteProductsIDOK with default headers values
func NewDeleteProductsIDOK() *DeleteProductsIDOK {

	return &DeleteProductsIDOK{}
}

// WriteResponse to the client
func (o *DeleteProductsIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// DeleteProductsIDNotFoundCode is the HTTP code returned for type DeleteProductsIDNotFound
const DeleteProductsIDNotFoundCode int = 404

/*DeleteProductsIDNotFound id not existent

swagger:response deleteProductsIdNotFound
*/
type DeleteProductsIDNotFound struct {
}

// NewDeleteProductsIDNotFound creates DeleteProductsIDNotFound with default headers values
func NewDeleteProductsIDNotFound() *DeleteProductsIDNotFound {

	return &DeleteProductsIDNotFound{}
}

// WriteResponse to the client
func (o *DeleteProductsIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// DeleteProductsIDInternalServerErrorCode is the HTTP code returned for type DeleteProductsIDInternalServerError
const DeleteProductsIDInternalServerErrorCode int = 500

/*DeleteProductsIDInternalServerError server error

swagger:response deleteProductsIdInternalServerError
*/
type DeleteProductsIDInternalServerError struct {
}

// NewDeleteProductsIDInternalServerError creates DeleteProductsIDInternalServerError with default headers values
func NewDeleteProductsIDInternalServerError() *DeleteProductsIDInternalServerError {

	return &DeleteProductsIDInternalServerError{}
}

// WriteResponse to the client
func (o *DeleteProductsIDInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}