// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PostUsersIDBuyNoContentCode is the HTTP code returned for type PostUsersIDBuyNoContent
const PostUsersIDBuyNoContentCode int = 204

/*PostUsersIDBuyNoContent success

swagger:response postUsersIdBuyNoContent
*/
type PostUsersIDBuyNoContent struct {
}

// NewPostUsersIDBuyNoContent creates PostUsersIDBuyNoContent with default headers values
func NewPostUsersIDBuyNoContent() *PostUsersIDBuyNoContent {

	return &PostUsersIDBuyNoContent{}
}

// WriteResponse to the client
func (o *PostUsersIDBuyNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// PostUsersIDBuyBadRequestCode is the HTTP code returned for type PostUsersIDBuyBadRequest
const PostUsersIDBuyBadRequestCode int = 400

/*PostUsersIDBuyBadRequest Invalid input

swagger:response postUsersIdBuyBadRequest
*/
type PostUsersIDBuyBadRequest struct {
}

// NewPostUsersIDBuyBadRequest creates PostUsersIDBuyBadRequest with default headers values
func NewPostUsersIDBuyBadRequest() *PostUsersIDBuyBadRequest {

	return &PostUsersIDBuyBadRequest{}
}

// WriteResponse to the client
func (o *PostUsersIDBuyBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// PostUsersIDBuyNotFoundCode is the HTTP code returned for type PostUsersIDBuyNotFound
const PostUsersIDBuyNotFoundCode int = 404

/*PostUsersIDBuyNotFound id not existent

swagger:response postUsersIdBuyNotFound
*/
type PostUsersIDBuyNotFound struct {
}

// NewPostUsersIDBuyNotFound creates PostUsersIDBuyNotFound with default headers values
func NewPostUsersIDBuyNotFound() *PostUsersIDBuyNotFound {

	return &PostUsersIDBuyNotFound{}
}

// WriteResponse to the client
func (o *PostUsersIDBuyNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// PostUsersIDBuyInternalServerErrorCode is the HTTP code returned for type PostUsersIDBuyInternalServerError
const PostUsersIDBuyInternalServerErrorCode int = 500

/*PostUsersIDBuyInternalServerError Server error

swagger:response postUsersIdBuyInternalServerError
*/
type PostUsersIDBuyInternalServerError struct {
}

// NewPostUsersIDBuyInternalServerError creates PostUsersIDBuyInternalServerError with default headers values
func NewPostUsersIDBuyInternalServerError() *PostUsersIDBuyInternalServerError {

	return &PostUsersIDBuyInternalServerError{}
}

// WriteResponse to the client
func (o *PostUsersIDBuyInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
