// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetUsersStatsHandlerFunc turns a function with the right signature into a get users stats handler
type GetUsersStatsHandlerFunc func(GetUsersStatsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetUsersStatsHandlerFunc) Handle(params GetUsersStatsParams) middleware.Responder {
	return fn(params)
}

// GetUsersStatsHandler interface for that can handle valid get users stats params
type GetUsersStatsHandler interface {
	Handle(GetUsersStatsParams) middleware.Responder
}

// NewGetUsersStats creates a new http.Handler for the get users stats operation
func NewGetUsersStats(ctx *middleware.Context, handler GetUsersStatsHandler) *GetUsersStats {
	return &GetUsersStats{Context: ctx, Handler: handler}
}

/*GetUsersStats swagger:route GET /users/stats/ users getUsersStats

stats about users

*/
type GetUsersStats struct {
	Context *middleware.Context
	Handler GetUsersStatsHandler
}

func (o *GetUsersStats) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetUsersStatsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// GetUsersStatsOKBody get users stats o k body
//
// swagger:model GetUsersStatsOKBody
type GetUsersStatsOKBody struct {

	// active count
	// Required: true
	ActiveCount *int64 `json:"active_count"`

	// as cents
	// Required: true
	BalanceSum *int64 `json:"balance_sum"`

	// user count
	// Required: true
	UserCount *int64 `json:"user_count"`
}

// Validate validates this get users stats o k body
func (o *GetUsersStatsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateActiveCount(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateBalanceSum(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateUserCount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetUsersStatsOKBody) validateActiveCount(formats strfmt.Registry) error {

	if err := validate.Required("getUsersStatsOK"+"."+"active_count", "body", o.ActiveCount); err != nil {
		return err
	}

	return nil
}

func (o *GetUsersStatsOKBody) validateBalanceSum(formats strfmt.Registry) error {

	if err := validate.Required("getUsersStatsOK"+"."+"balance_sum", "body", o.BalanceSum); err != nil {
		return err
	}

	return nil
}

func (o *GetUsersStatsOKBody) validateUserCount(formats strfmt.Registry) error {

	if err := validate.Required("getUsersStatsOK"+"."+"user_count", "body", o.UserCount); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetUsersStatsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetUsersStatsOKBody) UnmarshalBinary(b []byte) error {
	var res GetUsersStatsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
