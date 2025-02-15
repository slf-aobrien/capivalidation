// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TraditionalSummaryCoverageDTO TraditionalSummaryCoverageDTO
//
// swagger:model TraditionalSummaryCoverageDTO
type TraditionalSummaryCoverageDTO struct {

	// benefit amount
	BenefitAmount float64 `json:"benefitAmount,omitempty"`

	// benefit description
	BenefitDescription string `json:"benefitDescription,omitempty"`

	// benefit short name
	BenefitShortName string `json:"benefitShortName,omitempty"`

	// lives
	Lives int32 `json:"lives,omitempty"`

	// payroll deduction
	PayrollDeduction float64 `json:"payrollDeduction,omitempty"`

	// premium fees
	PremiumFees float64 `json:"premiumFees,omitempty"`
}

// Validate validates this traditional summary coverage d t o
func (m *TraditionalSummaryCoverageDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this traditional summary coverage d t o based on context it is used
func (m *TraditionalSummaryCoverageDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TraditionalSummaryCoverageDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TraditionalSummaryCoverageDTO) UnmarshalBinary(b []byte) error {
	var res TraditionalSummaryCoverageDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
