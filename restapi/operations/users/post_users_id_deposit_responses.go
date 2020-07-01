// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PostUsersIDDepositNoContentCode is the HTTP code returned for type PostUsersIDDepositNoContent
const PostUsersIDDepositNoContentCode int = 204

/*PostUsersIDDepositNoContent success

swagger:response postUsersIdDepositNoContent
*/
type PostUsersIDDepositNoContent struct {
}

// NewPostUsersIDDepositNoContent creates PostUsersIDDepositNoContent with default headers values
func NewPostUsersIDDepositNoContent() *PostUsersIDDepositNoContent {

	return &PostUsersIDDepositNoContent{}
}

// WriteResponse to the client
func (o *PostUsersIDDepositNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// PostUsersIDDepositBadRequestCode is the HTTP code returned for type PostUsersIDDepositBadRequest
const PostUsersIDDepositBadRequestCode int = 400

/*PostUsersIDDepositBadRequest Invalid input

swagger:response postUsersIdDepositBadRequest
*/
type PostUsersIDDepositBadRequest struct {
}

// NewPostUsersIDDepositBadRequest creates PostUsersIDDepositBadRequest with default headers values
func NewPostUsersIDDepositBadRequest() *PostUsersIDDepositBadRequest {

	return &PostUsersIDDepositBadRequest{}
}

// WriteResponse to the client
func (o *PostUsersIDDepositBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// PostUsersIDDepositNotFoundCode is the HTTP code returned for type PostUsersIDDepositNotFound
const PostUsersIDDepositNotFoundCode int = 404

/*PostUsersIDDepositNotFound id not existent

swagger:response postUsersIdDepositNotFound
*/
type PostUsersIDDepositNotFound struct {
}

// NewPostUsersIDDepositNotFound creates PostUsersIDDepositNotFound with default headers values
func NewPostUsersIDDepositNotFound() *PostUsersIDDepositNotFound {

	return &PostUsersIDDepositNotFound{}
}

// WriteResponse to the client
func (o *PostUsersIDDepositNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// PostUsersIDDepositInternalServerErrorCode is the HTTP code returned for type PostUsersIDDepositInternalServerError
const PostUsersIDDepositInternalServerErrorCode int = 500

/*PostUsersIDDepositInternalServerError Server error

swagger:response postUsersIdDepositInternalServerError
*/
type PostUsersIDDepositInternalServerError struct {
}

// NewPostUsersIDDepositInternalServerError creates PostUsersIDDepositInternalServerError with default headers values
func NewPostUsersIDDepositInternalServerError() *PostUsersIDDepositInternalServerError {

	return &PostUsersIDDepositInternalServerError{}
}

// WriteResponse to the client
func (o *PostUsersIDDepositInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
