// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetUsersBarcodeBarcodeHandlerFunc turns a function with the right signature into a get users barcode barcode handler
type GetUsersBarcodeBarcodeHandlerFunc func(GetUsersBarcodeBarcodeParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetUsersBarcodeBarcodeHandlerFunc) Handle(params GetUsersBarcodeBarcodeParams) middleware.Responder {
	return fn(params)
}

// GetUsersBarcodeBarcodeHandler interface for that can handle valid get users barcode barcode params
type GetUsersBarcodeBarcodeHandler interface {
	Handle(GetUsersBarcodeBarcodeParams) middleware.Responder
}

// NewGetUsersBarcodeBarcode creates a new http.Handler for the get users barcode barcode operation
func NewGetUsersBarcodeBarcode(ctx *middleware.Context, handler GetUsersBarcodeBarcodeHandler) *GetUsersBarcodeBarcode {
	return &GetUsersBarcodeBarcode{Context: ctx, Handler: handler}
}

/*GetUsersBarcodeBarcode swagger:route GET /users/barcode/{barcode}/ users getUsersBarcodeBarcode

get users details by associated barcode

*/
type GetUsersBarcodeBarcode struct {
	Context *middleware.Context
	Handler GetUsersBarcodeBarcodeHandler
}

func (o *GetUsersBarcodeBarcode) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetUsersBarcodeBarcodeParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
