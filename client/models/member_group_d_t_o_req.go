// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MemberGroupDTOReq MemberGroupDTOReq
//
// swagger:model MemberGroupDTOReq
type MemberGroupDTOReq struct {

	// alias
	Alias string `json:"alias,omitempty"`

	// benefit list
	BenefitList []*BenefitDTOReq `json:"benefitList"`

	// benefit package list
	BenefitPackageList []string `json:"benefitPackageList"`

	// commission package list
	CommissionPackageList []string `json:"commissionPackageList"`

	// description
	Description string `json:"description,omitempty"`

	// effective date
	EffectiveDate string `json:"effectiveDate,omitempty"`

	// expire date
	ExpireDate string `json:"expireDate,omitempty"`

	// fee package list
	FeePackageList []string `json:"feePackageList"`

	// interface Id
	InterfaceID string `json:"interfaceId,omitempty"`

	// key
	Key string `json:"key,omitempty"`

	// member count
	MemberCount string `json:"memberCount,omitempty"`

	// member group class d t o list
	MemberGroupClassDTOList []*MemberGroupClassDTO `json:"memberGroupClassDTOList"`

	// status
	Status string `json:"status,omitempty"`

	// system benefit list
	SystemBenefitList []string `json:"systemBenefitList"`

	// tpa package list
	TpaPackageList []string `json:"tpaPackageList"`

	// transfer rule list
	TransferRuleList []*TransferRuleDTO `json:"transferRuleList"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this member group d t o req
func (m *MemberGroupDTOReq) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBenefitList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBenefitPackageList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCommissionPackageList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFeePackageList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemberGroupClassDTOList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSystemBenefitList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTpaPackageList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTransferRuleList(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MemberGroupDTOReq) validateBenefitList(formats strfmt.Registry) error {
	if swag.IsZero(m.BenefitList) { // not required
		return nil
	}

	for i := 0; i < len(m.BenefitList); i++ {
		if swag.IsZero(m.BenefitList[i]) { // not required
			continue
		}

		if m.BenefitList[i] != nil {
			if err := m.BenefitList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("benefitList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("benefitList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var memberGroupDTOReqBenefitPackageListItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["$10_PREPAID_ADMIN_EXPENSE","$10_SMALL_GRP_LIFE_FEE","$15_ADMIN_FEE_EXPENSE","AA_CHILD_CRITICAL_ILLNESS","AA_CRITICAL_ILLNESS_COMM_RULE_EF","AA_EMPLOYEE_CRITICAL_ILLNESS","AA_SPOUSE_CRITICAL_ILLNESS","ABSENCE_ADA","ABSENCE_ADA_ALLOWANCE_WITHHELD_TPA","ABSENCE_ADA_COMMISSIONS_WITHHELD_TPA","ABSENCE_MANAGEMENT_SERVICES_ABSADA","ABSENCE_MANAGEMENT_SERVICES_STDFMLA","ABSENCE_MANAGEMENT_SERVICES_TPA_ALLOWANCE","ABSENCE_MANAGEMENT_SERVICES_TPA_COMM_WITHHELD","ACCIDENT_COMM_RULE_EF","ACCIDENT_TPA_ALLOWANCE","ACCIDENT_TPA_COMM_WITHHELD","ADDITIONAL_CONTRIBUTORY_LIFE","AGG_SL","ASO_REIMBURSEMENTS_REGULAR_BILLING","BUSINESS_INCOME_BENEFIT","CANCER_COMM_RULE_EF","CANCER_TPA_ALLOWANCE","CANCER_TPA_COMM_WITHHELD","CASH_BENEFIT","CA_PPD_REGULAR_BILLING","CDC_TG_TPA_COMMISSIONS_WITHHELD","CDC_TPA_ALLOWANCE_WITHHELD","CDC_VOL_TPA_COMMISSIONS_WITHHELD","CHILD_ACCIDENT","CHILD_CANCER","CHILD_CRITICAL_ILLNESS","CHILD_DENTAL","CHILD_GAP","CHILD_HOSPITAL_INDEMNITY","CO_PAID_FAMILY_AND_MEDICAL_LEAVE","CO_PFML_ASO_TPA_ALLOWANCE","CO_PFML_FI_TPA_ALLOWANCE","CO_PFML_TPA_COMMISSIONS_WITHHELD","CRITICAL_ILLNESS_TPA_ALLOWANCE","CRITICAL_ILLNESS_TPA_COMM_WITHHELD","CT_PAID_FAMILY_AND_MEDICAL_LEAVE","CT_PFML_ASO_TPA_ALLOWANCE","CT_PFML_FI_TPA_ALLOWANCE","CT_PFML_TPA_COMMISSIONS_WITHHELD","DENTAL_ASO_REGULAR_BILLING","DENTAL_ASO_TPA_ALLOW_WHD","DENTAL_ASO_TPA_COMM_WHD","DENTAL_CHILD","DENTAL_CHILD_ASO","DENTAL_EMPLOYEE","DENTAL_EMPLOYEE_ASO","DENTAL_FAMILY","DENTAL_FAMILY_ASO","DENTAL_SPOUSE","DENTAL_SPOUSE_ASO","DEPENDENT_ADD","DEPENDENT_LIFE","DE_PAID_FAMILY_AND_MEDICAL_LEAVE","DE_PFML_ASO_TPA_ALLOWANCE","DE_PFML_FI_TPA_ALLOWANCE","DE_PFML_TPA_COMMISSIONS_WITHHELD","EE_DENTAL_TPA_ALLOWANCE","EE_LIFE_TPA_ALLOWANCE","EMPLOYEE_ACCIDENT","EMPLOYEE_ADD","EMPLOYEE_ASSISTANCE_PLAN","EMPLOYEE_ASSISTANCE_PLAN_TPA_ALLOWANCE","EMPLOYEE_ASSISTANCE_PLAN_TPA_COMM_WITHHELD","EMPLOYEE_CANCER","EMPLOYEE_CRITICAL_ILLNESS","EMPLOYEE_DENTAL","EMPLOYEE_GAP","EMPLOYEE_HOSPITAL_INDEMNITY","EMPLOYEE_LIFE","ENROLLMENT_FEE","FAMILY_ACCIDENT","FAMILY_CANCER","FAMILY_DENTAL","FAMILY_GAP","FAMILY_HOSPITAL_INDEMNITY","FAMILY_LEAVE_INSURANCE","FLI_TPA_ALLOWANCE_WITHHELD","FLI_TPA_COMMISSION_WITHHELD","GAP_TPA_ALLOWANCE","GAP_TPA_COMM_WITHHELD","HEALTH_NAVIGATOR","HEALTH_NAVIGATOR_BILLING","HEALTH_NAVIGATOR_TPA_ALLOWANCE","HEALTH_NAVIGATOR_TPA_COMM_WITHHELD","HOSPITAL_INDEMNITY_COMM_RULE_EF","HOSPITAL_INDEMNITY_TPA_ALLOW_WITHHELD","HOSPITAL_INDEMNITY_TPA_COMMISSIONS_WITHHELD","IA_CRITICAL_ILLNESS_COMM_RULE_EF","IA_VOLUNTARY_STD","IA_VOL_STD_COMM_RULE_EF","IA_VOL_STD_TPA_ALLOWANCE","IA_VOL_STD_TPA_COMM_WITHHELD","IMPLEMENTATION_FEE","INSURED_CH_BUY_UP_VISION","INSURED_CH_VISION","INSURED_EE_BUY_UP_VISION","INSURED_EE_VISION","INSURED_FF_BUY_UP_VISION","INSURED_FF_VISION","INSURED_SP_BUY_UP_VISION","INSURED_SP_VISION","INSURED_VISION_TPA_ALLOWANCE","LIFE_TPA_COMM_WITHHELD","LIST_BILL_FEE","LONG_TERM_DISABILITY_BENEFIT","LONG_TERM_DISABILITY_PAYROLL","LTD_BUY_UP","LTD_TPA_ALLOWANCE","LTD_TPA_COMM_WITHHELD","MAXWELL_HEALTH","MAXWELL_HEALTH_FEE","MA_PAID_FAMILY_AND_MEDICAL_LEAVE","MA_PFML_ASO_TPA_ALLOWANCE","MA_PFML_FI_TPA_ALLOWANCE","MA_PFML_TPA_COMMISSIONS_WITHHELD","MEMBER_DIRECT_BILLING_IN_ADVANCE","MEMBER_DIRECT_BILLING_IN_ARREARS","MN_PAID_FAMILY_AND_MEDICAL_LEAVE","MN_PFML_ASO_TPA_ALLOWANCE","MN_PFML_FI_TPA_ALLOWANCE","MN_PFML_TPA_COMMISSIONS_WITHHELD","MONTHLY_IN_ARREARS_BILLING","NOTIONAL_BILLING","NY_PAID_FAMILY_LEAVE","NY_PAID_FAMILY_LEAVE_TPA_ALLOWANCE_WITHHELD","NY_PAID_FAMILY_LEAVE_TPA_COMMISSION_WITHHELD","OR_PAID_FAMILY_AND_MEDICAL_LEAVE","OR_PFML_ASO_TPA_ALLOWANCE","OR_PFML_FI_TPA_ALLOWANCE","OR_PFML_TPA_COMMISSIONS_WITHHELD","PREPAID_CHILD_DENTAL","PREPAID_COMMISSION_WITHHELD","PREPAID_EMPLOYEE_DENTAL","PREPAID_FAMILY_DENTAL","PREPAID_SPOUSE_DENTAL","PREPAID_TPA_ALLOWANCE","QUARTERLY_IN_ARREARS_BILLING","RECONCILIATION_BILLING","REGULAR_BILLING","SHORT_TERM_DISABILITY","SIMPLE_BENEFIT_SOLUTIONS","SMALL_GROUP_CERT_FEE","SPEC_SL_E1DEP","SPEC_SL_ECH","SPEC_SL_EE","SPEC_SL_EFAM","SPEC_SL_ESP","SPOUSE_ACCIDENT","SPOUSE_ACCIDENT_DISABILITY","SPOUSE_CANCER","SPOUSE_CRITICAL_ILLNESS","SPOUSE_DENTAL","SPOUSE_GAP","SPOUSE_HOSPITAL_INDEMNITY","STAND_ALONE_CHILD_VOL_ADD","STAND_ALONE_EMPLOYEE_VOL_ADD","STAND_ALONE_FAMILY_VOL_ADD","STAND_ALONE_SPOUSE_VOL_ADD","STAND_ALONE_VOL_ADD_TPA_ALLOWANCE","STAND_ALONE_VOL_ADD_TPA_COMM_WITHHELD","STD_BUY_UP","STD_NY_DBL_TPA_ALLOWANCE","STD_NY_DBL_TPA_COMM_WITHHELD","STD_TPA_ALLOWANCE","STD_TPA_COMM_WITHHELD","STOP_LOSS_AGG_BROKER_BROKER_FEE_WITHHELD","STOP_LOSS_AGG_COMM_WITHHELD_TPA","STOP_LOSS_AGG_TPA_ALLOWANCE_WITHHELD","STOP_LOSS_ANNUAL_BILLING","STOP_LOSS_MONTHLY_BILLING","STOP_LOSS_SPEC_BROKER_BROKER_FEE_WITHHELD","STOP_LOSS_SPEC_COMMISSION_WITHHELD_TPA","STOP_LOSS_SPEC_TPA_ALLOWANCE_WITHHELD","SUN_ADVISOR","SUN_ADVISOR_TPA_ALLOWANCE","SUN_ADVISOR_TPA_COMM_WITHHELD","TG_DEN_TPA_COMM_WITHHELD","TPA_COMMISSIONS_WITHHELD","TRAVEL_ACCIDENT","VISION","VOLUNTARY_CHILD_ADD","VOLUNTARY_CHILD_DENTAL","VOLUNTARY_CHILD_LIFE","VOLUNTARY_DEPENDENT_LIFE","VOLUNTARY_EMPLOYEE_ADD","VOLUNTARY_EMPLOYEE_DENTAL","VOLUNTARY_EMPLOYEE_LIFE","VOLUNTARY_FAMILY_DENTAL","VOLUNTARY_LTD","VOLUNTARY_SPOUSE_ADD","VOLUNTARY_SPOUSE_DENTAL","VOLUNTARY_SPOUSE_LIFE","VOLUNTARY_STD","VOL_DENTAL_TPA_ALLOWANCE","VOL_DEN_TPA_COMM_WITHHELD","VOL_LIFE_TPA_ALLOWANCE","VOL_LIFE_TPA_COMM_WITHHELD","VOL_LTD_TPA_ALLOWANCE","VOL_LTD_TPA_COMM_WITHHELD","VOL_STD_TPA_ALLOWANCE","VOL_STD_TPA_COMM_WITHHELD","WA_PAID_FAMILY_AND_MEDICAL_LEAVE","WA_PFML_TPA_ALLOWANCE","WA_PFML_TPA_COMMISSIONS_WITHHELD"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		memberGroupDTOReqBenefitPackageListItemsEnum = append(memberGroupDTOReqBenefitPackageListItemsEnum, v)
	}
}

func (m *MemberGroupDTOReq) validateBenefitPackageListItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, memberGroupDTOReqBenefitPackageListItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MemberGroupDTOReq) validateBenefitPackageList(formats strfmt.Registry) error {
	if swag.IsZero(m.BenefitPackageList) { // not required
		return nil
	}

	for i := 0; i < len(m.BenefitPackageList); i++ {

		// value enum
		if err := m.validateBenefitPackageListItemsEnum("benefitPackageList"+"."+strconv.Itoa(i), "body", m.BenefitPackageList[i]); err != nil {
			return err
		}

	}

	return nil
}

var memberGroupDTOReqCommissionPackageListItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["$10_PREPAID_ADMIN_EXPENSE","$10_SMALL_GRP_LIFE_FEE","$15_ADMIN_FEE_EXPENSE","AA_CHILD_CRITICAL_ILLNESS","AA_CRITICAL_ILLNESS_COMM_RULE_EF","AA_EMPLOYEE_CRITICAL_ILLNESS","AA_SPOUSE_CRITICAL_ILLNESS","ABSENCE_ADA","ABSENCE_ADA_ALLOWANCE_WITHHELD_TPA","ABSENCE_ADA_COMMISSIONS_WITHHELD_TPA","ABSENCE_MANAGEMENT_SERVICES_ABSADA","ABSENCE_MANAGEMENT_SERVICES_STDFMLA","ABSENCE_MANAGEMENT_SERVICES_TPA_ALLOWANCE","ABSENCE_MANAGEMENT_SERVICES_TPA_COMM_WITHHELD","ACCIDENT_COMM_RULE_EF","ACCIDENT_TPA_ALLOWANCE","ACCIDENT_TPA_COMM_WITHHELD","ADDITIONAL_CONTRIBUTORY_LIFE","AGG_SL","ASO_REIMBURSEMENTS_REGULAR_BILLING","BUSINESS_INCOME_BENEFIT","CANCER_COMM_RULE_EF","CANCER_TPA_ALLOWANCE","CANCER_TPA_COMM_WITHHELD","CASH_BENEFIT","CA_PPD_REGULAR_BILLING","CDC_TG_TPA_COMMISSIONS_WITHHELD","CDC_TPA_ALLOWANCE_WITHHELD","CDC_VOL_TPA_COMMISSIONS_WITHHELD","CHILD_ACCIDENT","CHILD_CANCER","CHILD_CRITICAL_ILLNESS","CHILD_DENTAL","CHILD_GAP","CHILD_HOSPITAL_INDEMNITY","CO_PAID_FAMILY_AND_MEDICAL_LEAVE","CO_PFML_ASO_TPA_ALLOWANCE","CO_PFML_FI_TPA_ALLOWANCE","CO_PFML_TPA_COMMISSIONS_WITHHELD","CRITICAL_ILLNESS_TPA_ALLOWANCE","CRITICAL_ILLNESS_TPA_COMM_WITHHELD","CT_PAID_FAMILY_AND_MEDICAL_LEAVE","CT_PFML_ASO_TPA_ALLOWANCE","CT_PFML_FI_TPA_ALLOWANCE","CT_PFML_TPA_COMMISSIONS_WITHHELD","DENTAL_ASO_REGULAR_BILLING","DENTAL_ASO_TPA_ALLOW_WHD","DENTAL_ASO_TPA_COMM_WHD","DENTAL_CHILD","DENTAL_CHILD_ASO","DENTAL_EMPLOYEE","DENTAL_EMPLOYEE_ASO","DENTAL_FAMILY","DENTAL_FAMILY_ASO","DENTAL_SPOUSE","DENTAL_SPOUSE_ASO","DEPENDENT_ADD","DEPENDENT_LIFE","DE_PAID_FAMILY_AND_MEDICAL_LEAVE","DE_PFML_ASO_TPA_ALLOWANCE","DE_PFML_FI_TPA_ALLOWANCE","DE_PFML_TPA_COMMISSIONS_WITHHELD","EE_DENTAL_TPA_ALLOWANCE","EE_LIFE_TPA_ALLOWANCE","EMPLOYEE_ACCIDENT","EMPLOYEE_ADD","EMPLOYEE_ASSISTANCE_PLAN","EMPLOYEE_ASSISTANCE_PLAN_TPA_ALLOWANCE","EMPLOYEE_ASSISTANCE_PLAN_TPA_COMM_WITHHELD","EMPLOYEE_CANCER","EMPLOYEE_CRITICAL_ILLNESS","EMPLOYEE_DENTAL","EMPLOYEE_GAP","EMPLOYEE_HOSPITAL_INDEMNITY","EMPLOYEE_LIFE","ENROLLMENT_FEE","FAMILY_ACCIDENT","FAMILY_CANCER","FAMILY_DENTAL","FAMILY_GAP","FAMILY_HOSPITAL_INDEMNITY","FAMILY_LEAVE_INSURANCE","FLI_TPA_ALLOWANCE_WITHHELD","FLI_TPA_COMMISSION_WITHHELD","GAP_TPA_ALLOWANCE","GAP_TPA_COMM_WITHHELD","HEALTH_NAVIGATOR","HEALTH_NAVIGATOR_BILLING","HEALTH_NAVIGATOR_TPA_ALLOWANCE","HEALTH_NAVIGATOR_TPA_COMM_WITHHELD","HOSPITAL_INDEMNITY_COMM_RULE_EF","HOSPITAL_INDEMNITY_TPA_ALLOW_WITHHELD","HOSPITAL_INDEMNITY_TPA_COMMISSIONS_WITHHELD","IA_CRITICAL_ILLNESS_COMM_RULE_EF","IA_VOLUNTARY_STD","IA_VOL_STD_COMM_RULE_EF","IA_VOL_STD_TPA_ALLOWANCE","IA_VOL_STD_TPA_COMM_WITHHELD","IMPLEMENTATION_FEE","INSURED_CH_BUY_UP_VISION","INSURED_CH_VISION","INSURED_EE_BUY_UP_VISION","INSURED_EE_VISION","INSURED_FF_BUY_UP_VISION","INSURED_FF_VISION","INSURED_SP_BUY_UP_VISION","INSURED_SP_VISION","INSURED_VISION_TPA_ALLOWANCE","LIFE_TPA_COMM_WITHHELD","LIST_BILL_FEE","LONG_TERM_DISABILITY_BENEFIT","LONG_TERM_DISABILITY_PAYROLL","LTD_BUY_UP","LTD_TPA_ALLOWANCE","LTD_TPA_COMM_WITHHELD","MAXWELL_HEALTH","MAXWELL_HEALTH_FEE","MA_PAID_FAMILY_AND_MEDICAL_LEAVE","MA_PFML_ASO_TPA_ALLOWANCE","MA_PFML_FI_TPA_ALLOWANCE","MA_PFML_TPA_COMMISSIONS_WITHHELD","MEMBER_DIRECT_BILLING_IN_ADVANCE","MEMBER_DIRECT_BILLING_IN_ARREARS","MN_PAID_FAMILY_AND_MEDICAL_LEAVE","MN_PFML_ASO_TPA_ALLOWANCE","MN_PFML_FI_TPA_ALLOWANCE","MN_PFML_TPA_COMMISSIONS_WITHHELD","MONTHLY_IN_ARREARS_BILLING","NOTIONAL_BILLING","NY_PAID_FAMILY_LEAVE","NY_PAID_FAMILY_LEAVE_TPA_ALLOWANCE_WITHHELD","NY_PAID_FAMILY_LEAVE_TPA_COMMISSION_WITHHELD","OR_PAID_FAMILY_AND_MEDICAL_LEAVE","OR_PFML_ASO_TPA_ALLOWANCE","OR_PFML_FI_TPA_ALLOWANCE","OR_PFML_TPA_COMMISSIONS_WITHHELD","PREPAID_CHILD_DENTAL","PREPAID_COMMISSION_WITHHELD","PREPAID_EMPLOYEE_DENTAL","PREPAID_FAMILY_DENTAL","PREPAID_SPOUSE_DENTAL","PREPAID_TPA_ALLOWANCE","QUARTERLY_IN_ARREARS_BILLING","RECONCILIATION_BILLING","REGULAR_BILLING","SHORT_TERM_DISABILITY","SIMPLE_BENEFIT_SOLUTIONS","SMALL_GROUP_CERT_FEE","SPEC_SL_E1DEP","SPEC_SL_ECH","SPEC_SL_EE","SPEC_SL_EFAM","SPEC_SL_ESP","SPOUSE_ACCIDENT","SPOUSE_ACCIDENT_DISABILITY","SPOUSE_CANCER","SPOUSE_CRITICAL_ILLNESS","SPOUSE_DENTAL","SPOUSE_GAP","SPOUSE_HOSPITAL_INDEMNITY","STAND_ALONE_CHILD_VOL_ADD","STAND_ALONE_EMPLOYEE_VOL_ADD","STAND_ALONE_FAMILY_VOL_ADD","STAND_ALONE_SPOUSE_VOL_ADD","STAND_ALONE_VOL_ADD_TPA_ALLOWANCE","STAND_ALONE_VOL_ADD_TPA_COMM_WITHHELD","STD_BUY_UP","STD_NY_DBL_TPA_ALLOWANCE","STD_NY_DBL_TPA_COMM_WITHHELD","STD_TPA_ALLOWANCE","STD_TPA_COMM_WITHHELD","STOP_LOSS_AGG_BROKER_BROKER_FEE_WITHHELD","STOP_LOSS_AGG_COMM_WITHHELD_TPA","STOP_LOSS_AGG_TPA_ALLOWANCE_WITHHELD","STOP_LOSS_ANNUAL_BILLING","STOP_LOSS_MONTHLY_BILLING","STOP_LOSS_SPEC_BROKER_BROKER_FEE_WITHHELD","STOP_LOSS_SPEC_COMMISSION_WITHHELD_TPA","STOP_LOSS_SPEC_TPA_ALLOWANCE_WITHHELD","SUN_ADVISOR","SUN_ADVISOR_TPA_ALLOWANCE","SUN_ADVISOR_TPA_COMM_WITHHELD","TG_DEN_TPA_COMM_WITHHELD","TPA_COMMISSIONS_WITHHELD","TRAVEL_ACCIDENT","VISION","VOLUNTARY_CHILD_ADD","VOLUNTARY_CHILD_DENTAL","VOLUNTARY_CHILD_LIFE","VOLUNTARY_DEPENDENT_LIFE","VOLUNTARY_EMPLOYEE_ADD","VOLUNTARY_EMPLOYEE_DENTAL","VOLUNTARY_EMPLOYEE_LIFE","VOLUNTARY_FAMILY_DENTAL","VOLUNTARY_LTD","VOLUNTARY_SPOUSE_ADD","VOLUNTARY_SPOUSE_DENTAL","VOLUNTARY_SPOUSE_LIFE","VOLUNTARY_STD","VOL_DENTAL_TPA_ALLOWANCE","VOL_DEN_TPA_COMM_WITHHELD","VOL_LIFE_TPA_ALLOWANCE","VOL_LIFE_TPA_COMM_WITHHELD","VOL_LTD_TPA_ALLOWANCE","VOL_LTD_TPA_COMM_WITHHELD","VOL_STD_TPA_ALLOWANCE","VOL_STD_TPA_COMM_WITHHELD","WA_PAID_FAMILY_AND_MEDICAL_LEAVE","WA_PFML_TPA_ALLOWANCE","WA_PFML_TPA_COMMISSIONS_WITHHELD"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		memberGroupDTOReqCommissionPackageListItemsEnum = append(memberGroupDTOReqCommissionPackageListItemsEnum, v)
	}
}

func (m *MemberGroupDTOReq) validateCommissionPackageListItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, memberGroupDTOReqCommissionPackageListItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MemberGroupDTOReq) validateCommissionPackageList(formats strfmt.Registry) error {
	if swag.IsZero(m.CommissionPackageList) { // not required
		return nil
	}

	for i := 0; i < len(m.CommissionPackageList); i++ {

		// value enum
		if err := m.validateCommissionPackageListItemsEnum("commissionPackageList"+"."+strconv.Itoa(i), "body", m.CommissionPackageList[i]); err != nil {
			return err
		}

	}

	return nil
}

var memberGroupDTOReqFeePackageListItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["$10_PREPAID_ADMIN_EXPENSE","$10_SMALL_GRP_LIFE_FEE","$15_ADMIN_FEE_EXPENSE","AA_CHILD_CRITICAL_ILLNESS","AA_CRITICAL_ILLNESS_COMM_RULE_EF","AA_EMPLOYEE_CRITICAL_ILLNESS","AA_SPOUSE_CRITICAL_ILLNESS","ABSENCE_ADA","ABSENCE_ADA_ALLOWANCE_WITHHELD_TPA","ABSENCE_ADA_COMMISSIONS_WITHHELD_TPA","ABSENCE_MANAGEMENT_SERVICES_ABSADA","ABSENCE_MANAGEMENT_SERVICES_STDFMLA","ABSENCE_MANAGEMENT_SERVICES_TPA_ALLOWANCE","ABSENCE_MANAGEMENT_SERVICES_TPA_COMM_WITHHELD","ACCIDENT_COMM_RULE_EF","ACCIDENT_TPA_ALLOWANCE","ACCIDENT_TPA_COMM_WITHHELD","ADDITIONAL_CONTRIBUTORY_LIFE","AGG_SL","ASO_REIMBURSEMENTS_REGULAR_BILLING","BUSINESS_INCOME_BENEFIT","CANCER_COMM_RULE_EF","CANCER_TPA_ALLOWANCE","CANCER_TPA_COMM_WITHHELD","CASH_BENEFIT","CA_PPD_REGULAR_BILLING","CDC_TG_TPA_COMMISSIONS_WITHHELD","CDC_TPA_ALLOWANCE_WITHHELD","CDC_VOL_TPA_COMMISSIONS_WITHHELD","CHILD_ACCIDENT","CHILD_CANCER","CHILD_CRITICAL_ILLNESS","CHILD_DENTAL","CHILD_GAP","CHILD_HOSPITAL_INDEMNITY","CO_PAID_FAMILY_AND_MEDICAL_LEAVE","CO_PFML_ASO_TPA_ALLOWANCE","CO_PFML_FI_TPA_ALLOWANCE","CO_PFML_TPA_COMMISSIONS_WITHHELD","CRITICAL_ILLNESS_TPA_ALLOWANCE","CRITICAL_ILLNESS_TPA_COMM_WITHHELD","CT_PAID_FAMILY_AND_MEDICAL_LEAVE","CT_PFML_ASO_TPA_ALLOWANCE","CT_PFML_FI_TPA_ALLOWANCE","CT_PFML_TPA_COMMISSIONS_WITHHELD","DENTAL_ASO_REGULAR_BILLING","DENTAL_ASO_TPA_ALLOW_WHD","DENTAL_ASO_TPA_COMM_WHD","DENTAL_CHILD","DENTAL_CHILD_ASO","DENTAL_EMPLOYEE","DENTAL_EMPLOYEE_ASO","DENTAL_FAMILY","DENTAL_FAMILY_ASO","DENTAL_SPOUSE","DENTAL_SPOUSE_ASO","DEPENDENT_ADD","DEPENDENT_LIFE","DE_PAID_FAMILY_AND_MEDICAL_LEAVE","DE_PFML_ASO_TPA_ALLOWANCE","DE_PFML_FI_TPA_ALLOWANCE","DE_PFML_TPA_COMMISSIONS_WITHHELD","EE_DENTAL_TPA_ALLOWANCE","EE_LIFE_TPA_ALLOWANCE","EMPLOYEE_ACCIDENT","EMPLOYEE_ADD","EMPLOYEE_ASSISTANCE_PLAN","EMPLOYEE_ASSISTANCE_PLAN_TPA_ALLOWANCE","EMPLOYEE_ASSISTANCE_PLAN_TPA_COMM_WITHHELD","EMPLOYEE_CANCER","EMPLOYEE_CRITICAL_ILLNESS","EMPLOYEE_DENTAL","EMPLOYEE_GAP","EMPLOYEE_HOSPITAL_INDEMNITY","EMPLOYEE_LIFE","ENROLLMENT_FEE","FAMILY_ACCIDENT","FAMILY_CANCER","FAMILY_DENTAL","FAMILY_GAP","FAMILY_HOSPITAL_INDEMNITY","FAMILY_LEAVE_INSURANCE","FLI_TPA_ALLOWANCE_WITHHELD","FLI_TPA_COMMISSION_WITHHELD","GAP_TPA_ALLOWANCE","GAP_TPA_COMM_WITHHELD","HEALTH_NAVIGATOR","HEALTH_NAVIGATOR_BILLING","HEALTH_NAVIGATOR_TPA_ALLOWANCE","HEALTH_NAVIGATOR_TPA_COMM_WITHHELD","HOSPITAL_INDEMNITY_COMM_RULE_EF","HOSPITAL_INDEMNITY_TPA_ALLOW_WITHHELD","HOSPITAL_INDEMNITY_TPA_COMMISSIONS_WITHHELD","IA_CRITICAL_ILLNESS_COMM_RULE_EF","IA_VOLUNTARY_STD","IA_VOL_STD_COMM_RULE_EF","IA_VOL_STD_TPA_ALLOWANCE","IA_VOL_STD_TPA_COMM_WITHHELD","IMPLEMENTATION_FEE","INSURED_CH_BUY_UP_VISION","INSURED_CH_VISION","INSURED_EE_BUY_UP_VISION","INSURED_EE_VISION","INSURED_FF_BUY_UP_VISION","INSURED_FF_VISION","INSURED_SP_BUY_UP_VISION","INSURED_SP_VISION","INSURED_VISION_TPA_ALLOWANCE","LIFE_TPA_COMM_WITHHELD","LIST_BILL_FEE","LONG_TERM_DISABILITY_BENEFIT","LONG_TERM_DISABILITY_PAYROLL","LTD_BUY_UP","LTD_TPA_ALLOWANCE","LTD_TPA_COMM_WITHHELD","MAXWELL_HEALTH","MAXWELL_HEALTH_FEE","MA_PAID_FAMILY_AND_MEDICAL_LEAVE","MA_PFML_ASO_TPA_ALLOWANCE","MA_PFML_FI_TPA_ALLOWANCE","MA_PFML_TPA_COMMISSIONS_WITHHELD","MEMBER_DIRECT_BILLING_IN_ADVANCE","MEMBER_DIRECT_BILLING_IN_ARREARS","MN_PAID_FAMILY_AND_MEDICAL_LEAVE","MN_PFML_ASO_TPA_ALLOWANCE","MN_PFML_FI_TPA_ALLOWANCE","MN_PFML_TPA_COMMISSIONS_WITHHELD","MONTHLY_IN_ARREARS_BILLING","NOTIONAL_BILLING","NY_PAID_FAMILY_LEAVE","NY_PAID_FAMILY_LEAVE_TPA_ALLOWANCE_WITHHELD","NY_PAID_FAMILY_LEAVE_TPA_COMMISSION_WITHHELD","OR_PAID_FAMILY_AND_MEDICAL_LEAVE","OR_PFML_ASO_TPA_ALLOWANCE","OR_PFML_FI_TPA_ALLOWANCE","OR_PFML_TPA_COMMISSIONS_WITHHELD","PREPAID_CHILD_DENTAL","PREPAID_COMMISSION_WITHHELD","PREPAID_EMPLOYEE_DENTAL","PREPAID_FAMILY_DENTAL","PREPAID_SPOUSE_DENTAL","PREPAID_TPA_ALLOWANCE","QUARTERLY_IN_ARREARS_BILLING","RECONCILIATION_BILLING","REGULAR_BILLING","SHORT_TERM_DISABILITY","SIMPLE_BENEFIT_SOLUTIONS","SMALL_GROUP_CERT_FEE","SPEC_SL_E1DEP","SPEC_SL_ECH","SPEC_SL_EE","SPEC_SL_EFAM","SPEC_SL_ESP","SPOUSE_ACCIDENT","SPOUSE_ACCIDENT_DISABILITY","SPOUSE_CANCER","SPOUSE_CRITICAL_ILLNESS","SPOUSE_DENTAL","SPOUSE_GAP","SPOUSE_HOSPITAL_INDEMNITY","STAND_ALONE_CHILD_VOL_ADD","STAND_ALONE_EMPLOYEE_VOL_ADD","STAND_ALONE_FAMILY_VOL_ADD","STAND_ALONE_SPOUSE_VOL_ADD","STAND_ALONE_VOL_ADD_TPA_ALLOWANCE","STAND_ALONE_VOL_ADD_TPA_COMM_WITHHELD","STD_BUY_UP","STD_NY_DBL_TPA_ALLOWANCE","STD_NY_DBL_TPA_COMM_WITHHELD","STD_TPA_ALLOWANCE","STD_TPA_COMM_WITHHELD","STOP_LOSS_AGG_BROKER_BROKER_FEE_WITHHELD","STOP_LOSS_AGG_COMM_WITHHELD_TPA","STOP_LOSS_AGG_TPA_ALLOWANCE_WITHHELD","STOP_LOSS_ANNUAL_BILLING","STOP_LOSS_MONTHLY_BILLING","STOP_LOSS_SPEC_BROKER_BROKER_FEE_WITHHELD","STOP_LOSS_SPEC_COMMISSION_WITHHELD_TPA","STOP_LOSS_SPEC_TPA_ALLOWANCE_WITHHELD","SUN_ADVISOR","SUN_ADVISOR_TPA_ALLOWANCE","SUN_ADVISOR_TPA_COMM_WITHHELD","TG_DEN_TPA_COMM_WITHHELD","TPA_COMMISSIONS_WITHHELD","TRAVEL_ACCIDENT","VISION","VOLUNTARY_CHILD_ADD","VOLUNTARY_CHILD_DENTAL","VOLUNTARY_CHILD_LIFE","VOLUNTARY_DEPENDENT_LIFE","VOLUNTARY_EMPLOYEE_ADD","VOLUNTARY_EMPLOYEE_DENTAL","VOLUNTARY_EMPLOYEE_LIFE","VOLUNTARY_FAMILY_DENTAL","VOLUNTARY_LTD","VOLUNTARY_SPOUSE_ADD","VOLUNTARY_SPOUSE_DENTAL","VOLUNTARY_SPOUSE_LIFE","VOLUNTARY_STD","VOL_DENTAL_TPA_ALLOWANCE","VOL_DEN_TPA_COMM_WITHHELD","VOL_LIFE_TPA_ALLOWANCE","VOL_LIFE_TPA_COMM_WITHHELD","VOL_LTD_TPA_ALLOWANCE","VOL_LTD_TPA_COMM_WITHHELD","VOL_STD_TPA_ALLOWANCE","VOL_STD_TPA_COMM_WITHHELD","WA_PAID_FAMILY_AND_MEDICAL_LEAVE","WA_PFML_TPA_ALLOWANCE","WA_PFML_TPA_COMMISSIONS_WITHHELD"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		memberGroupDTOReqFeePackageListItemsEnum = append(memberGroupDTOReqFeePackageListItemsEnum, v)
	}
}

func (m *MemberGroupDTOReq) validateFeePackageListItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, memberGroupDTOReqFeePackageListItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MemberGroupDTOReq) validateFeePackageList(formats strfmt.Registry) error {
	if swag.IsZero(m.FeePackageList) { // not required
		return nil
	}

	for i := 0; i < len(m.FeePackageList); i++ {

		// value enum
		if err := m.validateFeePackageListItemsEnum("feePackageList"+"."+strconv.Itoa(i), "body", m.FeePackageList[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *MemberGroupDTOReq) validateMemberGroupClassDTOList(formats strfmt.Registry) error {
	if swag.IsZero(m.MemberGroupClassDTOList) { // not required
		return nil
	}

	for i := 0; i < len(m.MemberGroupClassDTOList); i++ {
		if swag.IsZero(m.MemberGroupClassDTOList[i]) { // not required
			continue
		}

		if m.MemberGroupClassDTOList[i] != nil {
			if err := m.MemberGroupClassDTOList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("memberGroupClassDTOList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("memberGroupClassDTOList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var memberGroupDTOReqSystemBenefitListItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AA_CH_CIB","AA_EE_CIB","AA_SP_CIB","ACLI","AGG_SL","AMS","ASO_STD","BIB","CASH","CERT","CH_ACC","CH_CAN","CH_CIB","CH_DEN","CH_DEN_ASO","CH_GAP","CH_HOS_IND","CH_VIS","CO_PFML","CT_PFML","DEN_CH","DEN_EE","DEN_FAM","DEN_SP","DEP_ADD","DE_PFML","DLI","EAP","EE_ACC","EE_ADD","EE_CAN","EE_CIB","EE_DEN","EE_DEN_ASO","EE_GAP","EE_HOS_IND","EE_LIF","EE_VIS","FF_ACC","FF_CAN","FF_DEN","FF_DEN_ASO","FF_GAP","FF_HOS_IND","FF_VIS","FLEX_LTD","FLEX_STD","FLI","FLX_CH_VIS","FLX_EE_VIS","FLX_FF_VIS","FLX_SP_VIS","HEALTH_NAVIGATOR","IA_VSTD","LTD_B","LTD_P","MAXWELL","MA_PFML","MN_PFML","NY_PFL","OR_PFML","PP_CH_DEN","PP_EE_DEN","PP_FF_DEN","PP_SP_DEN","SA_V_CH_ADD","SA_V_EE_ADD","SA_V_FF_ADD","SA_V_SP_ADD","SPEC_SL_E1DEP","SPEC_SL_ECH","SPEC_SL_EE","SPEC_SL_EFAM","SPEC_SL_ESP","SP_ACC","SP_ACC_DIS","SP_CAN","SP_CIB","SP_DEN","SP_DEN_ASO","SP_GAP","SP_HOS_IND","SP_VIS","STD","TRAV_ACC","VISION","V_CH_ADD","V_CH_DEN","V_CH_LIF","V_DLI","V_EE_ADD","V_EE_DEN","V_EE_LIF","V_FF_DEN","V_LTD","V_SP_ADD","V_SP_DEN","V_SP_LIF","V_STD","WA_PFML"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		memberGroupDTOReqSystemBenefitListItemsEnum = append(memberGroupDTOReqSystemBenefitListItemsEnum, v)
	}
}

func (m *MemberGroupDTOReq) validateSystemBenefitListItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, memberGroupDTOReqSystemBenefitListItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MemberGroupDTOReq) validateSystemBenefitList(formats strfmt.Registry) error {
	if swag.IsZero(m.SystemBenefitList) { // not required
		return nil
	}

	for i := 0; i < len(m.SystemBenefitList); i++ {

		// value enum
		if err := m.validateSystemBenefitListItemsEnum("systemBenefitList"+"."+strconv.Itoa(i), "body", m.SystemBenefitList[i]); err != nil {
			return err
		}

	}

	return nil
}

var memberGroupDTOReqTpaPackageListItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["$10_PREPAID_ADMIN_EXPENSE","$10_SMALL_GRP_LIFE_FEE","$15_ADMIN_FEE_EXPENSE","AA_CHILD_CRITICAL_ILLNESS","AA_CRITICAL_ILLNESS_COMM_RULE_EF","AA_EMPLOYEE_CRITICAL_ILLNESS","AA_SPOUSE_CRITICAL_ILLNESS","ABSENCE_ADA","ABSENCE_ADA_ALLOWANCE_WITHHELD_TPA","ABSENCE_ADA_COMMISSIONS_WITHHELD_TPA","ABSENCE_MANAGEMENT_SERVICES_ABSADA","ABSENCE_MANAGEMENT_SERVICES_STDFMLA","ABSENCE_MANAGEMENT_SERVICES_TPA_ALLOWANCE","ABSENCE_MANAGEMENT_SERVICES_TPA_COMM_WITHHELD","ACCIDENT_COMM_RULE_EF","ACCIDENT_TPA_ALLOWANCE","ACCIDENT_TPA_COMM_WITHHELD","ADDITIONAL_CONTRIBUTORY_LIFE","AGG_SL","ASO_REIMBURSEMENTS_REGULAR_BILLING","BUSINESS_INCOME_BENEFIT","CANCER_COMM_RULE_EF","CANCER_TPA_ALLOWANCE","CANCER_TPA_COMM_WITHHELD","CASH_BENEFIT","CA_PPD_REGULAR_BILLING","CDC_TG_TPA_COMMISSIONS_WITHHELD","CDC_TPA_ALLOWANCE_WITHHELD","CDC_VOL_TPA_COMMISSIONS_WITHHELD","CHILD_ACCIDENT","CHILD_CANCER","CHILD_CRITICAL_ILLNESS","CHILD_DENTAL","CHILD_GAP","CHILD_HOSPITAL_INDEMNITY","CO_PAID_FAMILY_AND_MEDICAL_LEAVE","CO_PFML_ASO_TPA_ALLOWANCE","CO_PFML_FI_TPA_ALLOWANCE","CO_PFML_TPA_COMMISSIONS_WITHHELD","CRITICAL_ILLNESS_TPA_ALLOWANCE","CRITICAL_ILLNESS_TPA_COMM_WITHHELD","CT_PAID_FAMILY_AND_MEDICAL_LEAVE","CT_PFML_ASO_TPA_ALLOWANCE","CT_PFML_FI_TPA_ALLOWANCE","CT_PFML_TPA_COMMISSIONS_WITHHELD","DENTAL_ASO_REGULAR_BILLING","DENTAL_ASO_TPA_ALLOW_WHD","DENTAL_ASO_TPA_COMM_WHD","DENTAL_CHILD","DENTAL_CHILD_ASO","DENTAL_EMPLOYEE","DENTAL_EMPLOYEE_ASO","DENTAL_FAMILY","DENTAL_FAMILY_ASO","DENTAL_SPOUSE","DENTAL_SPOUSE_ASO","DEPENDENT_ADD","DEPENDENT_LIFE","DE_PAID_FAMILY_AND_MEDICAL_LEAVE","DE_PFML_ASO_TPA_ALLOWANCE","DE_PFML_FI_TPA_ALLOWANCE","DE_PFML_TPA_COMMISSIONS_WITHHELD","EE_DENTAL_TPA_ALLOWANCE","EE_LIFE_TPA_ALLOWANCE","EMPLOYEE_ACCIDENT","EMPLOYEE_ADD","EMPLOYEE_ASSISTANCE_PLAN","EMPLOYEE_ASSISTANCE_PLAN_TPA_ALLOWANCE","EMPLOYEE_ASSISTANCE_PLAN_TPA_COMM_WITHHELD","EMPLOYEE_CANCER","EMPLOYEE_CRITICAL_ILLNESS","EMPLOYEE_DENTAL","EMPLOYEE_GAP","EMPLOYEE_HOSPITAL_INDEMNITY","EMPLOYEE_LIFE","ENROLLMENT_FEE","FAMILY_ACCIDENT","FAMILY_CANCER","FAMILY_DENTAL","FAMILY_GAP","FAMILY_HOSPITAL_INDEMNITY","FAMILY_LEAVE_INSURANCE","FLI_TPA_ALLOWANCE_WITHHELD","FLI_TPA_COMMISSION_WITHHELD","GAP_TPA_ALLOWANCE","GAP_TPA_COMM_WITHHELD","HEALTH_NAVIGATOR","HEALTH_NAVIGATOR_BILLING","HEALTH_NAVIGATOR_TPA_ALLOWANCE","HEALTH_NAVIGATOR_TPA_COMM_WITHHELD","HOSPITAL_INDEMNITY_COMM_RULE_EF","HOSPITAL_INDEMNITY_TPA_ALLOW_WITHHELD","HOSPITAL_INDEMNITY_TPA_COMMISSIONS_WITHHELD","IA_CRITICAL_ILLNESS_COMM_RULE_EF","IA_VOLUNTARY_STD","IA_VOL_STD_COMM_RULE_EF","IA_VOL_STD_TPA_ALLOWANCE","IA_VOL_STD_TPA_COMM_WITHHELD","IMPLEMENTATION_FEE","INSURED_CH_BUY_UP_VISION","INSURED_CH_VISION","INSURED_EE_BUY_UP_VISION","INSURED_EE_VISION","INSURED_FF_BUY_UP_VISION","INSURED_FF_VISION","INSURED_SP_BUY_UP_VISION","INSURED_SP_VISION","INSURED_VISION_TPA_ALLOWANCE","LIFE_TPA_COMM_WITHHELD","LIST_BILL_FEE","LONG_TERM_DISABILITY_BENEFIT","LONG_TERM_DISABILITY_PAYROLL","LTD_BUY_UP","LTD_TPA_ALLOWANCE","LTD_TPA_COMM_WITHHELD","MAXWELL_HEALTH","MAXWELL_HEALTH_FEE","MA_PAID_FAMILY_AND_MEDICAL_LEAVE","MA_PFML_ASO_TPA_ALLOWANCE","MA_PFML_FI_TPA_ALLOWANCE","MA_PFML_TPA_COMMISSIONS_WITHHELD","MEMBER_DIRECT_BILLING_IN_ADVANCE","MEMBER_DIRECT_BILLING_IN_ARREARS","MN_PAID_FAMILY_AND_MEDICAL_LEAVE","MN_PFML_ASO_TPA_ALLOWANCE","MN_PFML_FI_TPA_ALLOWANCE","MN_PFML_TPA_COMMISSIONS_WITHHELD","MONTHLY_IN_ARREARS_BILLING","NOTIONAL_BILLING","NY_PAID_FAMILY_LEAVE","NY_PAID_FAMILY_LEAVE_TPA_ALLOWANCE_WITHHELD","NY_PAID_FAMILY_LEAVE_TPA_COMMISSION_WITHHELD","OR_PAID_FAMILY_AND_MEDICAL_LEAVE","OR_PFML_ASO_TPA_ALLOWANCE","OR_PFML_FI_TPA_ALLOWANCE","OR_PFML_TPA_COMMISSIONS_WITHHELD","PREPAID_CHILD_DENTAL","PREPAID_COMMISSION_WITHHELD","PREPAID_EMPLOYEE_DENTAL","PREPAID_FAMILY_DENTAL","PREPAID_SPOUSE_DENTAL","PREPAID_TPA_ALLOWANCE","QUARTERLY_IN_ARREARS_BILLING","RECONCILIATION_BILLING","REGULAR_BILLING","SHORT_TERM_DISABILITY","SIMPLE_BENEFIT_SOLUTIONS","SMALL_GROUP_CERT_FEE","SPEC_SL_E1DEP","SPEC_SL_ECH","SPEC_SL_EE","SPEC_SL_EFAM","SPEC_SL_ESP","SPOUSE_ACCIDENT","SPOUSE_ACCIDENT_DISABILITY","SPOUSE_CANCER","SPOUSE_CRITICAL_ILLNESS","SPOUSE_DENTAL","SPOUSE_GAP","SPOUSE_HOSPITAL_INDEMNITY","STAND_ALONE_CHILD_VOL_ADD","STAND_ALONE_EMPLOYEE_VOL_ADD","STAND_ALONE_FAMILY_VOL_ADD","STAND_ALONE_SPOUSE_VOL_ADD","STAND_ALONE_VOL_ADD_TPA_ALLOWANCE","STAND_ALONE_VOL_ADD_TPA_COMM_WITHHELD","STD_BUY_UP","STD_NY_DBL_TPA_ALLOWANCE","STD_NY_DBL_TPA_COMM_WITHHELD","STD_TPA_ALLOWANCE","STD_TPA_COMM_WITHHELD","STOP_LOSS_AGG_BROKER_BROKER_FEE_WITHHELD","STOP_LOSS_AGG_COMM_WITHHELD_TPA","STOP_LOSS_AGG_TPA_ALLOWANCE_WITHHELD","STOP_LOSS_ANNUAL_BILLING","STOP_LOSS_MONTHLY_BILLING","STOP_LOSS_SPEC_BROKER_BROKER_FEE_WITHHELD","STOP_LOSS_SPEC_COMMISSION_WITHHELD_TPA","STOP_LOSS_SPEC_TPA_ALLOWANCE_WITHHELD","SUN_ADVISOR","SUN_ADVISOR_TPA_ALLOWANCE","SUN_ADVISOR_TPA_COMM_WITHHELD","TG_DEN_TPA_COMM_WITHHELD","TPA_COMMISSIONS_WITHHELD","TRAVEL_ACCIDENT","VISION","VOLUNTARY_CHILD_ADD","VOLUNTARY_CHILD_DENTAL","VOLUNTARY_CHILD_LIFE","VOLUNTARY_DEPENDENT_LIFE","VOLUNTARY_EMPLOYEE_ADD","VOLUNTARY_EMPLOYEE_DENTAL","VOLUNTARY_EMPLOYEE_LIFE","VOLUNTARY_FAMILY_DENTAL","VOLUNTARY_LTD","VOLUNTARY_SPOUSE_ADD","VOLUNTARY_SPOUSE_DENTAL","VOLUNTARY_SPOUSE_LIFE","VOLUNTARY_STD","VOL_DENTAL_TPA_ALLOWANCE","VOL_DEN_TPA_COMM_WITHHELD","VOL_LIFE_TPA_ALLOWANCE","VOL_LIFE_TPA_COMM_WITHHELD","VOL_LTD_TPA_ALLOWANCE","VOL_LTD_TPA_COMM_WITHHELD","VOL_STD_TPA_ALLOWANCE","VOL_STD_TPA_COMM_WITHHELD","WA_PAID_FAMILY_AND_MEDICAL_LEAVE","WA_PFML_TPA_ALLOWANCE","WA_PFML_TPA_COMMISSIONS_WITHHELD"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		memberGroupDTOReqTpaPackageListItemsEnum = append(memberGroupDTOReqTpaPackageListItemsEnum, v)
	}
}

func (m *MemberGroupDTOReq) validateTpaPackageListItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, memberGroupDTOReqTpaPackageListItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MemberGroupDTOReq) validateTpaPackageList(formats strfmt.Registry) error {
	if swag.IsZero(m.TpaPackageList) { // not required
		return nil
	}

	for i := 0; i < len(m.TpaPackageList); i++ {

		// value enum
		if err := m.validateTpaPackageListItemsEnum("tpaPackageList"+"."+strconv.Itoa(i), "body", m.TpaPackageList[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *MemberGroupDTOReq) validateTransferRuleList(formats strfmt.Registry) error {
	if swag.IsZero(m.TransferRuleList) { // not required
		return nil
	}

	for i := 0; i < len(m.TransferRuleList); i++ {
		if swag.IsZero(m.TransferRuleList[i]) { // not required
			continue
		}

		if m.TransferRuleList[i] != nil {
			if err := m.TransferRuleList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("transferRuleList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("transferRuleList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this member group d t o req based on the context it is used
func (m *MemberGroupDTOReq) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBenefitList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMemberGroupClassDTOList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTransferRuleList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MemberGroupDTOReq) contextValidateBenefitList(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.BenefitList); i++ {

		if m.BenefitList[i] != nil {

			if swag.IsZero(m.BenefitList[i]) { // not required
				return nil
			}

			if err := m.BenefitList[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("benefitList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("benefitList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MemberGroupDTOReq) contextValidateMemberGroupClassDTOList(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.MemberGroupClassDTOList); i++ {

		if m.MemberGroupClassDTOList[i] != nil {

			if swag.IsZero(m.MemberGroupClassDTOList[i]) { // not required
				return nil
			}

			if err := m.MemberGroupClassDTOList[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("memberGroupClassDTOList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("memberGroupClassDTOList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MemberGroupDTOReq) contextValidateTransferRuleList(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.TransferRuleList); i++ {

		if m.TransferRuleList[i] != nil {

			if swag.IsZero(m.TransferRuleList[i]) { // not required
				return nil
			}

			if err := m.TransferRuleList[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("transferRuleList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("transferRuleList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *MemberGroupDTOReq) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MemberGroupDTOReq) UnmarshalBinary(b []byte) error {
	var res MemberGroupDTOReq
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
