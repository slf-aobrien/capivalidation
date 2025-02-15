// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RateScaleOptionDTO RateScaleOptionDTO
//
// swagger:model RateScaleOptionDTO
type RateScaleOptionDTO struct {

	// basis
	Basis string `json:"basis,omitempty"`

	// domain name
	DomainName string `json:"domainName,omitempty"`

	// label alias name
	LabelAliasName string `json:"labelAliasName,omitempty"`

	// lower amount
	LowerAmount string `json:"lowerAmount,omitempty"`

	// reference code
	ReferenceCode string `json:"referenceCode,omitempty"`

	// service definition
	ServiceDefinition string `json:"serviceDefinition,omitempty"`

	// upper amount
	UpperAmount string `json:"upperAmount,omitempty"`
}

// Validate validates this rate scale option d t o
func (m *RateScaleOptionDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this rate scale option d t o based on context it is used
func (m *RateScaleOptionDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RateScaleOptionDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RateScaleOptionDTO) UnmarshalBinary(b []byte) error {
	var res RateScaleOptionDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
