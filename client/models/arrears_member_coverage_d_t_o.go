// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ArrearsMemberCoverageDTO ArrearsMemberCoverageDTO
//
// swagger:model ArrearsMemberCoverageDTO
type ArrearsMemberCoverageDTO struct {

	// age band
	AgeBand string `json:"ageBand,omitempty"`

	// benefit amount
	BenefitAmount float64 `json:"benefitAmount,omitempty"`

	// benefit short name
	BenefitShortName string `json:"benefitShortName,omitempty"`

	// benefit total
	BenefitTotal float64 `json:"benefitTotal,omitempty"`

	// employee contrib
	EmployeeContrib float64 `json:"employeeContrib,omitempty"`

	// employer contrib
	EmployerContrib float64 `json:"employerContrib,omitempty"`

	// member record type
	MemberRecordType string `json:"memberRecordType,omitempty"`

	// past due amount
	PastDueAmount float64 `json:"pastDueAmount,omitempty"`

	// payroll period
	PayrollPeriod string `json:"payrollPeriod,omitempty"`

	// rollup indicator
	RollupIndicator string `json:"rollupIndicator,omitempty"`
}

// Validate validates this arrears member coverage d t o
func (m *ArrearsMemberCoverageDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this arrears member coverage d t o based on context it is used
func (m *ArrearsMemberCoverageDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ArrearsMemberCoverageDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ArrearsMemberCoverageDTO) UnmarshalBinary(b []byte) error {
	var res ArrearsMemberCoverageDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
