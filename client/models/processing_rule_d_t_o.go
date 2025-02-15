// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ProcessingRuleDTO ProcessingRuleDTO
//
// swagger:model ProcessingRuleDTO
type ProcessingRuleDTO struct {

	// effective date
	EffectiveDate string `json:"effectiveDate,omitempty"`

	// expiry date
	ExpiryDate string `json:"expiryDate,omitempty"`

	// process method
	ProcessMethod string `json:"processMethod,omitempty"`

	// process option
	ProcessOption bool `json:"processOption,omitempty"`
}

// Validate validates this processing rule d t o
func (m *ProcessingRuleDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this processing rule d t o based on context it is used
func (m *ProcessingRuleDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ProcessingRuleDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProcessingRuleDTO) UnmarshalBinary(b []byte) error {
	var res ProcessingRuleDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
