// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Server server
//
// swagger:model Server
type Server struct {

	// currency
	// Required: true
	Currency *string `json:"currency"`

	// If true the currency symbol is show before the sum
	// Required: true
	CurrencyBefore bool `json:"currency_before"`

	// if None, the currency uses no subdevision and therefore no decimal seperator is used. E.g. in Sweden
	// Required: true
	DecimalSeperator *string `json:"decimal_seperator"`

	// defaults
	Defaults *ServerDefaults `json:"defaults,omitempty"`

	// unit of energy e.g. kcal or kj
	// Required: true
	Energy *string `json:"energy"`

	// global credit limit in cent (subdevision of currency); to disable set to false
	GlobalCreditLimit int64 `json:"global_credit_limit,omitempty"`

	// version
	// Required: true
	Version *string `json:"version"`
}

// Validate validates this server
func (m *Server) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCurrency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrencyBefore(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDecimalSeperator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDefaults(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnergy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Server) validateCurrency(formats strfmt.Registry) error {

	if err := validate.Required("currency", "body", m.Currency); err != nil {
		return err
	}

	return nil
}

func (m *Server) validateCurrencyBefore(formats strfmt.Registry) error {

	if err := validate.Required("currency_before", "body", bool(m.CurrencyBefore)); err != nil {
		return err
	}

	return nil
}

func (m *Server) validateDecimalSeperator(formats strfmt.Registry) error {

	if err := validate.Required("decimal_seperator", "body", m.DecimalSeperator); err != nil {
		return err
	}

	return nil
}

func (m *Server) validateDefaults(formats strfmt.Registry) error {

	if swag.IsZero(m.Defaults) { // not required
		return nil
	}

	if m.Defaults != nil {
		if err := m.Defaults.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("defaults")
			}
			return err
		}
	}

	return nil
}

func (m *Server) validateEnergy(formats strfmt.Registry) error {

	if err := validate.Required("energy", "body", m.Energy); err != nil {
		return err
	}

	return nil
}

func (m *Server) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Server) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Server) UnmarshalBinary(b []byte) error {
	var res Server
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ServerDefaults defaults for new products
//
// swagger:model ServerDefaults
type ServerDefaults struct {

	// default weather a product is active upon creation or not
	Active bool `json:"active,omitempty"`

	// default volume percent of alcohol (with two decimal places)
	Alcohol int64 `json:"alcohol,omitempty"`

	// default caffeine contents in mg per 100ml/g
	Caffeine int64 `json:"caffeine,omitempty"`

	// default energy per 100g / 100ml without decimal places
	Energy int64 `json:"energy,omitempty"`

	// default package size
	PackageSize string `json:"package_size,omitempty"`

	// default price in cent
	Price int64 `json:"price,omitempty"`

	// default sugar amount per 100g / 100ml with one decimal place
	Sugar int64 `json:"sugar,omitempty"`
}

// Validate validates this server defaults
func (m *ServerDefaults) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServerDefaults) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServerDefaults) UnmarshalBinary(b []byte) error {
	var res ServerDefaults
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
