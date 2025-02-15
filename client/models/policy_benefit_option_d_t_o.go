// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PolicyBenefitOptionDTO PolicyBenefitOptionDTO
//
// swagger:model PolicyBenefitOptionDTO
type PolicyBenefitOptionDTO struct {

	// benefit description
	BenefitDescription string `json:"benefitDescription,omitempty"`

	// benefit short name
	BenefitShortName string `json:"benefitShortName,omitempty"`

	// code
	Code string `json:"code,omitempty"`

	// decode
	Decode string `json:"decode,omitempty"`

	// effective date
	EffectiveDate string `json:"effectiveDate,omitempty"`

	// expiration date
	ExpirationDate string `json:"expirationDate,omitempty"`

	// field name
	FieldName string `json:"fieldName,omitempty"`

	// label alias name
	LabelAliasName string `json:"labelAliasName,omitempty"`

	// policy number
	PolicyNumber string `json:"policyNumber,omitempty"`

	// replace existing
	ReplaceExisting bool `json:"replaceExisting,omitempty"`
}

// Validate validates this policy benefit option d t o
func (m *PolicyBenefitOptionDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this policy benefit option d t o based on context it is used
func (m *PolicyBenefitOptionDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PolicyBenefitOptionDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PolicyBenefitOptionDTO) UnmarshalBinary(b []byte) error {
	var res PolicyBenefitOptionDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
