// Code generated by go-swagger; DO NOT EDIT.

package barcodes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetBarcodesIDHandlerFunc turns a function with the right signature into a get barcodes ID handler
type GetBarcodesIDHandlerFunc func(GetBarcodesIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetBarcodesIDHandlerFunc) Handle(params GetBarcodesIDParams) middleware.Responder {
	return fn(params)
}

// GetBarcodesIDHandler interface for that can handle valid get barcodes ID params
type GetBarcodesIDHandler interface {
	Handle(GetBarcodesIDParams) middleware.Responder
}

// NewGetBarcodesID creates a new http.Handler for the get barcodes ID operation
func NewGetBarcodesID(ctx *middleware.Context, handler GetBarcodesIDHandler) *GetBarcodesID {
	return &GetBarcodesID{Context: ctx, Handler: handler}
}

/*GetBarcodesID swagger:route GET /barcodes/{id}/ barcodes getBarcodesId

get details for a given barcode

*/
type GetBarcodesID struct {
	Context *middleware.Context
	Handler GetBarcodesIDHandler
}

func (o *GetBarcodesID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetBarcodesIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
