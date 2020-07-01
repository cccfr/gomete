// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// DeleteUsersIDOKCode is the HTTP code returned for type DeleteUsersIDOK
const DeleteUsersIDOKCode int = 200

/*DeleteUsersIDOK success

swagger:response deleteUsersIdOK
*/
type DeleteUsersIDOK struct {
}

// NewDeleteUsersIDOK creates DeleteUsersIDOK with default headers values
func NewDeleteUsersIDOK() *DeleteUsersIDOK {

	return &DeleteUsersIDOK{}
}

// WriteResponse to the client
func (o *DeleteUsersIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// DeleteUsersIDNotFoundCode is the HTTP code returned for type DeleteUsersIDNotFound
const DeleteUsersIDNotFoundCode int = 404

/*DeleteUsersIDNotFound id not existent

swagger:response deleteUsersIdNotFound
*/
type DeleteUsersIDNotFound struct {
}

// NewDeleteUsersIDNotFound creates DeleteUsersIDNotFound with default headers values
func NewDeleteUsersIDNotFound() *DeleteUsersIDNotFound {

	return &DeleteUsersIDNotFound{}
}

// WriteResponse to the client
func (o *DeleteUsersIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// DeleteUsersIDInternalServerErrorCode is the HTTP code returned for type DeleteUsersIDInternalServerError
const DeleteUsersIDInternalServerErrorCode int = 500

/*DeleteUsersIDInternalServerError Server error

swagger:response deleteUsersIdInternalServerError
*/
type DeleteUsersIDInternalServerError struct {
}

// NewDeleteUsersIDInternalServerError creates DeleteUsersIDInternalServerError with default headers values
func NewDeleteUsersIDInternalServerError() *DeleteUsersIDInternalServerError {

	return &DeleteUsersIDInternalServerError{}
}

// WriteResponse to the client
func (o *DeleteUsersIDInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
