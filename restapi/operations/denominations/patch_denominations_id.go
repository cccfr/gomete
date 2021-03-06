// Code generated by go-swagger; DO NOT EDIT.

package denominations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PatchDenominationsIDHandlerFunc turns a function with the right signature into a patch denominations ID handler
type PatchDenominationsIDHandlerFunc func(PatchDenominationsIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PatchDenominationsIDHandlerFunc) Handle(params PatchDenominationsIDParams) middleware.Responder {
	return fn(params)
}

// PatchDenominationsIDHandler interface for that can handle valid patch denominations ID params
type PatchDenominationsIDHandler interface {
	Handle(PatchDenominationsIDParams) middleware.Responder
}

// NewPatchDenominationsID creates a new http.Handler for the patch denominations ID operation
func NewPatchDenominationsID(ctx *middleware.Context, handler PatchDenominationsIDHandler) *PatchDenominationsID {
	return &PatchDenominationsID{Context: ctx, Handler: handler}
}

/*PatchDenominationsID swagger:route PATCH /denominations/{id}/ denominations patchDenominationsId

update a denomination

*/
type PatchDenominationsID struct {
	Context *middleware.Context
	Handler PatchDenominationsIDHandler
}

func (o *PatchDenominationsID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPatchDenominationsIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PatchDenominationsIDBody patch denominations ID body
//
// swagger:model PatchDenominationsIDBody
type PatchDenominationsIDBody struct {

	// amount
	Amount int64 `json:"amount,omitempty"`

	// image
	Image int64 `json:"image,omitempty"`
}

// Validate validates this patch denominations ID body
func (o *PatchDenominationsIDBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PatchDenominationsIDBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PatchDenominationsIDBody) UnmarshalBinary(b []byte) error {
	var res PatchDenominationsIDBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
