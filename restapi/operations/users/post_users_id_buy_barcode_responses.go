// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PostUsersIDBuyBarcodeNoContentCode is the HTTP code returned for type PostUsersIDBuyBarcodeNoContent
const PostUsersIDBuyBarcodeNoContentCode int = 204

/*PostUsersIDBuyBarcodeNoContent success

swagger:response postUsersIdBuyBarcodeNoContent
*/
type PostUsersIDBuyBarcodeNoContent struct {
}

// NewPostUsersIDBuyBarcodeNoContent creates PostUsersIDBuyBarcodeNoContent with default headers values
func NewPostUsersIDBuyBarcodeNoContent() *PostUsersIDBuyBarcodeNoContent {

	return &PostUsersIDBuyBarcodeNoContent{}
}

// WriteResponse to the client
func (o *PostUsersIDBuyBarcodeNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// PostUsersIDBuyBarcodeBadRequestCode is the HTTP code returned for type PostUsersIDBuyBarcodeBadRequest
const PostUsersIDBuyBarcodeBadRequestCode int = 400

/*PostUsersIDBuyBarcodeBadRequest Invalid input

swagger:response postUsersIdBuyBarcodeBadRequest
*/
type PostUsersIDBuyBarcodeBadRequest struct {
}

// NewPostUsersIDBuyBarcodeBadRequest creates PostUsersIDBuyBarcodeBadRequest with default headers values
func NewPostUsersIDBuyBarcodeBadRequest() *PostUsersIDBuyBarcodeBadRequest {

	return &PostUsersIDBuyBarcodeBadRequest{}
}

// WriteResponse to the client
func (o *PostUsersIDBuyBarcodeBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// PostUsersIDBuyBarcodeNotFoundCode is the HTTP code returned for type PostUsersIDBuyBarcodeNotFound
const PostUsersIDBuyBarcodeNotFoundCode int = 404

/*PostUsersIDBuyBarcodeNotFound id not existent

swagger:response postUsersIdBuyBarcodeNotFound
*/
type PostUsersIDBuyBarcodeNotFound struct {
}

// NewPostUsersIDBuyBarcodeNotFound creates PostUsersIDBuyBarcodeNotFound with default headers values
func NewPostUsersIDBuyBarcodeNotFound() *PostUsersIDBuyBarcodeNotFound {

	return &PostUsersIDBuyBarcodeNotFound{}
}

// WriteResponse to the client
func (o *PostUsersIDBuyBarcodeNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// PostUsersIDBuyBarcodeInternalServerErrorCode is the HTTP code returned for type PostUsersIDBuyBarcodeInternalServerError
const PostUsersIDBuyBarcodeInternalServerErrorCode int = 500

/*PostUsersIDBuyBarcodeInternalServerError Server error

swagger:response postUsersIdBuyBarcodeInternalServerError
*/
type PostUsersIDBuyBarcodeInternalServerError struct {
}

// NewPostUsersIDBuyBarcodeInternalServerError creates PostUsersIDBuyBarcodeInternalServerError with default headers values
func NewPostUsersIDBuyBarcodeInternalServerError() *PostUsersIDBuyBarcodeInternalServerError {

	return &PostUsersIDBuyBarcodeInternalServerError{}
}

// WriteResponse to the client
func (o *PostUsersIDBuyBarcodeInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}