// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PortabilityRuleDTO PortabilityRuleDTO
//
// swagger:model PortabilityRuleDTO
type PortabilityRuleDTO struct {

	// coverage change cd
	CoverageChangeCd string `json:"coverageChangeCd,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// effective date string
	EffectiveDateString string `json:"effectiveDateString,omitempty"`

	// event code
	EventCode string `json:"eventCode,omitempty"`

	// tmrl description
	TmrlDescription string `json:"tmrlDescription,omitempty"`

	// type code
	TypeCode string `json:"typeCode,omitempty"`
}

// Validate validates this portability rule d t o
func (m *PortabilityRuleDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this portability rule d t o based on context it is used
func (m *PortabilityRuleDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PortabilityRuleDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortabilityRuleDTO) UnmarshalBinary(b []byte) error {
	var res PortabilityRuleDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
