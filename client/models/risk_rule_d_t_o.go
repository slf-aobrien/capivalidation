// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RiskRuleDTO RiskRuleDTO
//
// swagger:model RiskRuleDTO
type RiskRuleDTO struct {

	// dependent data code
	DependentDataCode string `json:"dependentDataCode,omitempty"`

	// risk benefit tranch list
	RiskBenefitTranchList []*RiskBenefitTranchDTO `json:"riskBenefitTranchList"`

	// risk enhanced eligibility rule list
	RiskEnhancedEligibilityRuleList []*RiskEnhancedEligibilityRuleDTO `json:"riskEnhancedEligibilityRuleList"`

	// risk premium tranch list
	RiskPremiumTranchList []*RiskPremiumTranchDTO `json:"riskPremiumTranchList"`

	// risk rule detail list
	RiskRuleDetailList []*RiskRuleDetailDTO `json:"riskRuleDetailList"`

	// rkrl ben dtrm code
	RkrlBenDtrmCode string `json:"rkrlBenDtrmCode,omitempty"`

	// rkrl prrt code
	RkrlPrrtCode string `json:"rkrlPrrtCode,omitempty"`

	// system benefit
	SystemBenefit string `json:"systemBenefit,omitempty"`

	// underwriting data code
	UnderwritingDataCode string `json:"underwritingDataCode,omitempty"`
}

// Validate validates this risk rule d t o
func (m *RiskRuleDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRiskBenefitTranchList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRiskEnhancedEligibilityRuleList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRiskPremiumTranchList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRiskRuleDetailList(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RiskRuleDTO) validateRiskBenefitTranchList(formats strfmt.Registry) error {
	if swag.IsZero(m.RiskBenefitTranchList) { // not required
		return nil
	}

	for i := 0; i < len(m.RiskBenefitTranchList); i++ {
		if swag.IsZero(m.RiskBenefitTranchList[i]) { // not required
			continue
		}

		if m.RiskBenefitTranchList[i] != nil {
			if err := m.RiskBenefitTranchList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("riskBenefitTranchList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("riskBenefitTranchList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RiskRuleDTO) validateRiskEnhancedEligibilityRuleList(formats strfmt.Registry) error {
	if swag.IsZero(m.RiskEnhancedEligibilityRuleList) { // not required
		return nil
	}

	for i := 0; i < len(m.RiskEnhancedEligibilityRuleList); i++ {
		if swag.IsZero(m.RiskEnhancedEligibilityRuleList[i]) { // not required
			continue
		}

		if m.RiskEnhancedEligibilityRuleList[i] != nil {
			if err := m.RiskEnhancedEligibilityRuleList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("riskEnhancedEligibilityRuleList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("riskEnhancedEligibilityRuleList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RiskRuleDTO) validateRiskPremiumTranchList(formats strfmt.Registry) error {
	if swag.IsZero(m.RiskPremiumTranchList) { // not required
		return nil
	}

	for i := 0; i < len(m.RiskPremiumTranchList); i++ {
		if swag.IsZero(m.RiskPremiumTranchList[i]) { // not required
			continue
		}

		if m.RiskPremiumTranchList[i] != nil {
			if err := m.RiskPremiumTranchList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("riskPremiumTranchList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("riskPremiumTranchList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RiskRuleDTO) validateRiskRuleDetailList(formats strfmt.Registry) error {
	if swag.IsZero(m.RiskRuleDetailList) { // not required
		return nil
	}

	for i := 0; i < len(m.RiskRuleDetailList); i++ {
		if swag.IsZero(m.RiskRuleDetailList[i]) { // not required
			continue
		}

		if m.RiskRuleDetailList[i] != nil {
			if err := m.RiskRuleDetailList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("riskRuleDetailList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("riskRuleDetailList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this risk rule d t o based on the context it is used
func (m *RiskRuleDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRiskBenefitTranchList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRiskEnhancedEligibilityRuleList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRiskPremiumTranchList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRiskRuleDetailList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RiskRuleDTO) contextValidateRiskBenefitTranchList(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.RiskBenefitTranchList); i++ {

		if m.RiskBenefitTranchList[i] != nil {

			if swag.IsZero(m.RiskBenefitTranchList[i]) { // not required
				return nil
			}

			if err := m.RiskBenefitTranchList[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("riskBenefitTranchList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("riskBenefitTranchList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RiskRuleDTO) contextValidateRiskEnhancedEligibilityRuleList(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.RiskEnhancedEligibilityRuleList); i++ {

		if m.RiskEnhancedEligibilityRuleList[i] != nil {

			if swag.IsZero(m.RiskEnhancedEligibilityRuleList[i]) { // not required
				return nil
			}

			if err := m.RiskEnhancedEligibilityRuleList[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("riskEnhancedEligibilityRuleList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("riskEnhancedEligibilityRuleList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RiskRuleDTO) contextValidateRiskPremiumTranchList(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.RiskPremiumTranchList); i++ {

		if m.RiskPremiumTranchList[i] != nil {

			if swag.IsZero(m.RiskPremiumTranchList[i]) { // not required
				return nil
			}

			if err := m.RiskPremiumTranchList[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("riskPremiumTranchList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("riskPremiumTranchList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RiskRuleDTO) contextValidateRiskRuleDetailList(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.RiskRuleDetailList); i++ {

		if m.RiskRuleDetailList[i] != nil {

			if swag.IsZero(m.RiskRuleDetailList[i]) { // not required
				return nil
			}

			if err := m.RiskRuleDetailList[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("riskRuleDetailList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("riskRuleDetailList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RiskRuleDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RiskRuleDTO) UnmarshalBinary(b []byte) error {
	var res RiskRuleDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
