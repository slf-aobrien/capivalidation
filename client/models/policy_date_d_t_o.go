// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PolicyDateDTO PolicyDateDTO
//
// swagger:model PolicyDateDTO
type PolicyDateDTO struct {

	// date
	Date string `json:"date,omitempty"`

	// date type
	// Enum: ["ACTUARIAL_VALUATION_DATE","ADD_A_COVERAGE_DATE","ANNIVERSARY_DATE","ANNIVERSARY_RELEVANT_DATE","ANNUAL_REPORTING_DATE","BALANCE_DATE","BEGINNING_COST_BASIS_DATE","BENEFIT_RATE_INDEX_DATE","CEASE_OVERRIDE_PRICING_UV1","CELL_RISK_CALCULATION_DATE","CHECK_RECEIPT_DATE","COMMISSION_END_DATE","CONVERSION_DATE","CONVERSION_DATE_SKIP_PPPR_PRIOR_TO_THIS_DATE","CONVERSION_RISK_CALCULATION_DATE","DAILY_PRICING_DATE","DISTRIBUTIONS_PROCESS_DATE","EOI_MEMBER_COUNT_CONVERSION","INITIAL_BENEFIT_CALCULATION_DATE","LAST_ACTUARIAL_VALUATION_DATE","LAST_AWARD_YEAR_END","LAST_BILL_DUE_DATE","LAST_DEATH_IN_SERVICE_UNIT_RATE_REVISION_DATE","LAST_DECLARED_BONUS_DATE","LAST_FUNDING_REVIEW_PROCESSING_DATE","LAST_INDEMNITY_WRAP_UP","LAST_INTEREST_DECLARATION_PRIOR_TO_TRANSFER_CONVERSION_ONLY","LAST_INTEREST_DECLARATION_PROCESSING_DATE","LAST_INTERIM_DIV_INCOME_BONUS_DT","LAST_INTERIM_INCOME_BONUS","LAST_INTERIM_INTEREST_INCOME_BONUS_DATE","LAST_INTERIM_RENT_INCOME_BONUS_DT","LAST_RECOSTING_PROCESSING_DATE","LAST_REMITTANCE_COVERED_DATE","LAST_RISK_COSTING_DATE","LAST_SALARY_UPDATE_DATE","LAST_SINGLE_UNIT_PRICE_ADJUSTMENT","LAST_STATEMENT_DATE","LAST_TAX_YEAR_END_DATE","NEXT_ANNIVERSARY_DATE","NM_LAST_ANNUAL_PREMIUM_RISK_CALC","ONE_APRIL_1992","ONE_JANUARY_1999","ONE_MAY_1996","OPTION_DATE","ORIGINAL_OPTION_DATE","OVERRIDE_PREMIUM_BILLING","PAID_TO_DATE","PBA_TO_DPBA_CHANGE_DATE","POLICY_CEASE","POST_INTEREST_INITIAL_DATE","PRIOR_INTEREST_DECLARATION_PROCESSING_DATE","PRIOR_YEARS_EXPENSE_BILL_DUE_DATE","PRIOR_YEARS_INTEREST_DECLARATION_DATE","PRIOR_YEARS_RECOSTING_DATE","PSO_APPROVAL_DATE","RATE_CHANGE_DATE","RATE_TRANCHE_HISTORY_DATE","REALLOCATION_ANDT","REALLOCATION_BASE_DATE","REALLOCATION_BCHG","REALLOCATION_BEOP","REALLOCATION_BNTR","REALLOCATION_CSST","REALLOCATION_FCOR","REALLOCATION_FDRD","REALLOCATION_FRCR","REALLOCATION_GLRD","REALLOCATION_GRDT","REALLOCATION_PMTD","REALLOCATION_PMTR","REALLOCATION_POLICY_MAINTENANCE","REALLOCATION_RENEWAL","REALLOCATION_RERL","REALLOCATION_RKBS","REALLOCATION_RKPS","REALLOCATION_RKRL","REALLOCATION_RRED","REALLOCATION_RTSC","REALLOCATION_SCPR","REALLOCATION_SNDT","REALLOCATION_UNDC","REALLOCATION_UWOT","REASSURANCE_CONVERSION_DATE","REPORTING_BALANCE_DATE","RESERVED_FOR_CLIENT","RESERVED_FOR_REALLOCATION","RFT_CALCULATION_DATE","RISK_BENEFIT_EXPIRY_DATE","RISK_GUARANTEE_DATE","SCHEME_COMMENCEMENT_DATE","SECOND_SHORT_YEAR_SCHEME_DATE","SHADOW_CONVERSION_DATE","SHORT_1ST_BILL_PERIOD_END_DATE","SPECIAL_CONVERSION_DATE","SPECIAL_MLC_DATE","SUM_OF_TRANS_BEGIN_DATE_INCLUSIVE","SUM_OF_TRANS_END_DATE_INCLUSIVE","SWITCH_TO_DUAL","TABLE_CONVERSION_DATE","THIRTY_APR_96","THIRTY_ONE_DECEMBER_1988","TRANSACTION_DETAIL_JOURNAL_LOGS_AUDIT_DATE","TRANSFER_DATE_FOR_CONVERSION_FUNDS_ONLY","USE_FATUM_AGE_APPROACH_ON_OR_AFTER_THIS_DATE","VALID_TARGET_RETIREMENT_DATE_USER_DEFINED","WRAPUP_DATE"]
	DateType string `json:"dateType,omitempty"`

	// effective date
	EffectiveDate string `json:"effectiveDate,omitempty"`

	// policy number
	PolicyNumber string `json:"policyNumber,omitempty"`
}

// Validate validates this policy date d t o
func (m *PolicyDateDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var policyDateDTOTypeDateTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ACTUARIAL_VALUATION_DATE","ADD_A_COVERAGE_DATE","ANNIVERSARY_DATE","ANNIVERSARY_RELEVANT_DATE","ANNUAL_REPORTING_DATE","BALANCE_DATE","BEGINNING_COST_BASIS_DATE","BENEFIT_RATE_INDEX_DATE","CEASE_OVERRIDE_PRICING_UV1","CELL_RISK_CALCULATION_DATE","CHECK_RECEIPT_DATE","COMMISSION_END_DATE","CONVERSION_DATE","CONVERSION_DATE_SKIP_PPPR_PRIOR_TO_THIS_DATE","CONVERSION_RISK_CALCULATION_DATE","DAILY_PRICING_DATE","DISTRIBUTIONS_PROCESS_DATE","EOI_MEMBER_COUNT_CONVERSION","INITIAL_BENEFIT_CALCULATION_DATE","LAST_ACTUARIAL_VALUATION_DATE","LAST_AWARD_YEAR_END","LAST_BILL_DUE_DATE","LAST_DEATH_IN_SERVICE_UNIT_RATE_REVISION_DATE","LAST_DECLARED_BONUS_DATE","LAST_FUNDING_REVIEW_PROCESSING_DATE","LAST_INDEMNITY_WRAP_UP","LAST_INTEREST_DECLARATION_PRIOR_TO_TRANSFER_CONVERSION_ONLY","LAST_INTEREST_DECLARATION_PROCESSING_DATE","LAST_INTERIM_DIV_INCOME_BONUS_DT","LAST_INTERIM_INCOME_BONUS","LAST_INTERIM_INTEREST_INCOME_BONUS_DATE","LAST_INTERIM_RENT_INCOME_BONUS_DT","LAST_RECOSTING_PROCESSING_DATE","LAST_REMITTANCE_COVERED_DATE","LAST_RISK_COSTING_DATE","LAST_SALARY_UPDATE_DATE","LAST_SINGLE_UNIT_PRICE_ADJUSTMENT","LAST_STATEMENT_DATE","LAST_TAX_YEAR_END_DATE","NEXT_ANNIVERSARY_DATE","NM_LAST_ANNUAL_PREMIUM_RISK_CALC","ONE_APRIL_1992","ONE_JANUARY_1999","ONE_MAY_1996","OPTION_DATE","ORIGINAL_OPTION_DATE","OVERRIDE_PREMIUM_BILLING","PAID_TO_DATE","PBA_TO_DPBA_CHANGE_DATE","POLICY_CEASE","POST_INTEREST_INITIAL_DATE","PRIOR_INTEREST_DECLARATION_PROCESSING_DATE","PRIOR_YEARS_EXPENSE_BILL_DUE_DATE","PRIOR_YEARS_INTEREST_DECLARATION_DATE","PRIOR_YEARS_RECOSTING_DATE","PSO_APPROVAL_DATE","RATE_CHANGE_DATE","RATE_TRANCHE_HISTORY_DATE","REALLOCATION_ANDT","REALLOCATION_BASE_DATE","REALLOCATION_BCHG","REALLOCATION_BEOP","REALLOCATION_BNTR","REALLOCATION_CSST","REALLOCATION_FCOR","REALLOCATION_FDRD","REALLOCATION_FRCR","REALLOCATION_GLRD","REALLOCATION_GRDT","REALLOCATION_PMTD","REALLOCATION_PMTR","REALLOCATION_POLICY_MAINTENANCE","REALLOCATION_RENEWAL","REALLOCATION_RERL","REALLOCATION_RKBS","REALLOCATION_RKPS","REALLOCATION_RKRL","REALLOCATION_RRED","REALLOCATION_RTSC","REALLOCATION_SCPR","REALLOCATION_SNDT","REALLOCATION_UNDC","REALLOCATION_UWOT","REASSURANCE_CONVERSION_DATE","REPORTING_BALANCE_DATE","RESERVED_FOR_CLIENT","RESERVED_FOR_REALLOCATION","RFT_CALCULATION_DATE","RISK_BENEFIT_EXPIRY_DATE","RISK_GUARANTEE_DATE","SCHEME_COMMENCEMENT_DATE","SECOND_SHORT_YEAR_SCHEME_DATE","SHADOW_CONVERSION_DATE","SHORT_1ST_BILL_PERIOD_END_DATE","SPECIAL_CONVERSION_DATE","SPECIAL_MLC_DATE","SUM_OF_TRANS_BEGIN_DATE_INCLUSIVE","SUM_OF_TRANS_END_DATE_INCLUSIVE","SWITCH_TO_DUAL","TABLE_CONVERSION_DATE","THIRTY_APR_96","THIRTY_ONE_DECEMBER_1988","TRANSACTION_DETAIL_JOURNAL_LOGS_AUDIT_DATE","TRANSFER_DATE_FOR_CONVERSION_FUNDS_ONLY","USE_FATUM_AGE_APPROACH_ON_OR_AFTER_THIS_DATE","VALID_TARGET_RETIREMENT_DATE_USER_DEFINED","WRAPUP_DATE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		policyDateDTOTypeDateTypePropEnum = append(policyDateDTOTypeDateTypePropEnum, v)
	}
}

const (

	// PolicyDateDTODateTypeACTUARIALVALUATIONDATE captures enum value "ACTUARIAL_VALUATION_DATE"
	PolicyDateDTODateTypeACTUARIALVALUATIONDATE string = "ACTUARIAL_VALUATION_DATE"

	// PolicyDateDTODateTypeADDACOVERAGEDATE captures enum value "ADD_A_COVERAGE_DATE"
	PolicyDateDTODateTypeADDACOVERAGEDATE string = "ADD_A_COVERAGE_DATE"

	// PolicyDateDTODateTypeANNIVERSARYDATE captures enum value "ANNIVERSARY_DATE"
	PolicyDateDTODateTypeANNIVERSARYDATE string = "ANNIVERSARY_DATE"

	// PolicyDateDTODateTypeANNIVERSARYRELEVANTDATE captures enum value "ANNIVERSARY_RELEVANT_DATE"
	PolicyDateDTODateTypeANNIVERSARYRELEVANTDATE string = "ANNIVERSARY_RELEVANT_DATE"

	// PolicyDateDTODateTypeANNUALREPORTINGDATE captures enum value "ANNUAL_REPORTING_DATE"
	PolicyDateDTODateTypeANNUALREPORTINGDATE string = "ANNUAL_REPORTING_DATE"

	// PolicyDateDTODateTypeBALANCEDATE captures enum value "BALANCE_DATE"
	PolicyDateDTODateTypeBALANCEDATE string = "BALANCE_DATE"

	// PolicyDateDTODateTypeBEGINNINGCOSTBASISDATE captures enum value "BEGINNING_COST_BASIS_DATE"
	PolicyDateDTODateTypeBEGINNINGCOSTBASISDATE string = "BEGINNING_COST_BASIS_DATE"

	// PolicyDateDTODateTypeBENEFITRATEINDEXDATE captures enum value "BENEFIT_RATE_INDEX_DATE"
	PolicyDateDTODateTypeBENEFITRATEINDEXDATE string = "BENEFIT_RATE_INDEX_DATE"

	// PolicyDateDTODateTypeCEASEOVERRIDEPRICINGUV1 captures enum value "CEASE_OVERRIDE_PRICING_UV1"
	PolicyDateDTODateTypeCEASEOVERRIDEPRICINGUV1 string = "CEASE_OVERRIDE_PRICING_UV1"

	// PolicyDateDTODateTypeCELLRISKCALCULATIONDATE captures enum value "CELL_RISK_CALCULATION_DATE"
	PolicyDateDTODateTypeCELLRISKCALCULATIONDATE string = "CELL_RISK_CALCULATION_DATE"

	// PolicyDateDTODateTypeCHECKRECEIPTDATE captures enum value "CHECK_RECEIPT_DATE"
	PolicyDateDTODateTypeCHECKRECEIPTDATE string = "CHECK_RECEIPT_DATE"

	// PolicyDateDTODateTypeCOMMISSIONENDDATE captures enum value "COMMISSION_END_DATE"
	PolicyDateDTODateTypeCOMMISSIONENDDATE string = "COMMISSION_END_DATE"

	// PolicyDateDTODateTypeCONVERSIONDATE captures enum value "CONVERSION_DATE"
	PolicyDateDTODateTypeCONVERSIONDATE string = "CONVERSION_DATE"

	// PolicyDateDTODateTypeCONVERSIONDATESKIPPPPRPRIORTOTHISDATE captures enum value "CONVERSION_DATE_SKIP_PPPR_PRIOR_TO_THIS_DATE"
	PolicyDateDTODateTypeCONVERSIONDATESKIPPPPRPRIORTOTHISDATE string = "CONVERSION_DATE_SKIP_PPPR_PRIOR_TO_THIS_DATE"

	// PolicyDateDTODateTypeCONVERSIONRISKCALCULATIONDATE captures enum value "CONVERSION_RISK_CALCULATION_DATE"
	PolicyDateDTODateTypeCONVERSIONRISKCALCULATIONDATE string = "CONVERSION_RISK_CALCULATION_DATE"

	// PolicyDateDTODateTypeDAILYPRICINGDATE captures enum value "DAILY_PRICING_DATE"
	PolicyDateDTODateTypeDAILYPRICINGDATE string = "DAILY_PRICING_DATE"

	// PolicyDateDTODateTypeDISTRIBUTIONSPROCESSDATE captures enum value "DISTRIBUTIONS_PROCESS_DATE"
	PolicyDateDTODateTypeDISTRIBUTIONSPROCESSDATE string = "DISTRIBUTIONS_PROCESS_DATE"

	// PolicyDateDTODateTypeEOIMEMBERCOUNTCONVERSION captures enum value "EOI_MEMBER_COUNT_CONVERSION"
	PolicyDateDTODateTypeEOIMEMBERCOUNTCONVERSION string = "EOI_MEMBER_COUNT_CONVERSION"

	// PolicyDateDTODateTypeINITIALBENEFITCALCULATIONDATE captures enum value "INITIAL_BENEFIT_CALCULATION_DATE"
	PolicyDateDTODateTypeINITIALBENEFITCALCULATIONDATE string = "INITIAL_BENEFIT_CALCULATION_DATE"

	// PolicyDateDTODateTypeLASTACTUARIALVALUATIONDATE captures enum value "LAST_ACTUARIAL_VALUATION_DATE"
	PolicyDateDTODateTypeLASTACTUARIALVALUATIONDATE string = "LAST_ACTUARIAL_VALUATION_DATE"

	// PolicyDateDTODateTypeLASTAWARDYEAREND captures enum value "LAST_AWARD_YEAR_END"
	PolicyDateDTODateTypeLASTAWARDYEAREND string = "LAST_AWARD_YEAR_END"

	// PolicyDateDTODateTypeLASTBILLDUEDATE captures enum value "LAST_BILL_DUE_DATE"
	PolicyDateDTODateTypeLASTBILLDUEDATE string = "LAST_BILL_DUE_DATE"

	// PolicyDateDTODateTypeLASTDEATHINSERVICEUNITRATEREVISIONDATE captures enum value "LAST_DEATH_IN_SERVICE_UNIT_RATE_REVISION_DATE"
	PolicyDateDTODateTypeLASTDEATHINSERVICEUNITRATEREVISIONDATE string = "LAST_DEATH_IN_SERVICE_UNIT_RATE_REVISION_DATE"

	// PolicyDateDTODateTypeLASTDECLAREDBONUSDATE captures enum value "LAST_DECLARED_BONUS_DATE"
	PolicyDateDTODateTypeLASTDECLAREDBONUSDATE string = "LAST_DECLARED_BONUS_DATE"

	// PolicyDateDTODateTypeLASTFUNDINGREVIEWPROCESSINGDATE captures enum value "LAST_FUNDING_REVIEW_PROCESSING_DATE"
	PolicyDateDTODateTypeLASTFUNDINGREVIEWPROCESSINGDATE string = "LAST_FUNDING_REVIEW_PROCESSING_DATE"

	// PolicyDateDTODateTypeLASTINDEMNITYWRAPUP captures enum value "LAST_INDEMNITY_WRAP_UP"
	PolicyDateDTODateTypeLASTINDEMNITYWRAPUP string = "LAST_INDEMNITY_WRAP_UP"

	// PolicyDateDTODateTypeLASTINTERESTDECLARATIONPRIORTOTRANSFERCONVERSIONONLY captures enum value "LAST_INTEREST_DECLARATION_PRIOR_TO_TRANSFER_CONVERSION_ONLY"
	PolicyDateDTODateTypeLASTINTERESTDECLARATIONPRIORTOTRANSFERCONVERSIONONLY string = "LAST_INTEREST_DECLARATION_PRIOR_TO_TRANSFER_CONVERSION_ONLY"

	// PolicyDateDTODateTypeLASTINTERESTDECLARATIONPROCESSINGDATE captures enum value "LAST_INTEREST_DECLARATION_PROCESSING_DATE"
	PolicyDateDTODateTypeLASTINTERESTDECLARATIONPROCESSINGDATE string = "LAST_INTEREST_DECLARATION_PROCESSING_DATE"

	// PolicyDateDTODateTypeLASTINTERIMDIVINCOMEBONUSDT captures enum value "LAST_INTERIM_DIV_INCOME_BONUS_DT"
	PolicyDateDTODateTypeLASTINTERIMDIVINCOMEBONUSDT string = "LAST_INTERIM_DIV_INCOME_BONUS_DT"

	// PolicyDateDTODateTypeLASTINTERIMINCOMEBONUS captures enum value "LAST_INTERIM_INCOME_BONUS"
	PolicyDateDTODateTypeLASTINTERIMINCOMEBONUS string = "LAST_INTERIM_INCOME_BONUS"

	// PolicyDateDTODateTypeLASTINTERIMINTERESTINCOMEBONUSDATE captures enum value "LAST_INTERIM_INTEREST_INCOME_BONUS_DATE"
	PolicyDateDTODateTypeLASTINTERIMINTERESTINCOMEBONUSDATE string = "LAST_INTERIM_INTEREST_INCOME_BONUS_DATE"

	// PolicyDateDTODateTypeLASTINTERIMRENTINCOMEBONUSDT captures enum value "LAST_INTERIM_RENT_INCOME_BONUS_DT"
	PolicyDateDTODateTypeLASTINTERIMRENTINCOMEBONUSDT string = "LAST_INTERIM_RENT_INCOME_BONUS_DT"

	// PolicyDateDTODateTypeLASTRECOSTINGPROCESSINGDATE captures enum value "LAST_RECOSTING_PROCESSING_DATE"
	PolicyDateDTODateTypeLASTRECOSTINGPROCESSINGDATE string = "LAST_RECOSTING_PROCESSING_DATE"

	// PolicyDateDTODateTypeLASTREMITTANCECOVEREDDATE captures enum value "LAST_REMITTANCE_COVERED_DATE"
	PolicyDateDTODateTypeLASTREMITTANCECOVEREDDATE string = "LAST_REMITTANCE_COVERED_DATE"

	// PolicyDateDTODateTypeLASTRISKCOSTINGDATE captures enum value "LAST_RISK_COSTING_DATE"
	PolicyDateDTODateTypeLASTRISKCOSTINGDATE string = "LAST_RISK_COSTING_DATE"

	// PolicyDateDTODateTypeLASTSALARYUPDATEDATE captures enum value "LAST_SALARY_UPDATE_DATE"
	PolicyDateDTODateTypeLASTSALARYUPDATEDATE string = "LAST_SALARY_UPDATE_DATE"

	// PolicyDateDTODateTypeLASTSINGLEUNITPRICEADJUSTMENT captures enum value "LAST_SINGLE_UNIT_PRICE_ADJUSTMENT"
	PolicyDateDTODateTypeLASTSINGLEUNITPRICEADJUSTMENT string = "LAST_SINGLE_UNIT_PRICE_ADJUSTMENT"

	// PolicyDateDTODateTypeLASTSTATEMENTDATE captures enum value "LAST_STATEMENT_DATE"
	PolicyDateDTODateTypeLASTSTATEMENTDATE string = "LAST_STATEMENT_DATE"

	// PolicyDateDTODateTypeLASTTAXYEARENDDATE captures enum value "LAST_TAX_YEAR_END_DATE"
	PolicyDateDTODateTypeLASTTAXYEARENDDATE string = "LAST_TAX_YEAR_END_DATE"

	// PolicyDateDTODateTypeNEXTANNIVERSARYDATE captures enum value "NEXT_ANNIVERSARY_DATE"
	PolicyDateDTODateTypeNEXTANNIVERSARYDATE string = "NEXT_ANNIVERSARY_DATE"

	// PolicyDateDTODateTypeNMLASTANNUALPREMIUMRISKCALC captures enum value "NM_LAST_ANNUAL_PREMIUM_RISK_CALC"
	PolicyDateDTODateTypeNMLASTANNUALPREMIUMRISKCALC string = "NM_LAST_ANNUAL_PREMIUM_RISK_CALC"

	// PolicyDateDTODateTypeONEAPRIL1992 captures enum value "ONE_APRIL_1992"
	PolicyDateDTODateTypeONEAPRIL1992 string = "ONE_APRIL_1992"

	// PolicyDateDTODateTypeONEJANUARY1999 captures enum value "ONE_JANUARY_1999"
	PolicyDateDTODateTypeONEJANUARY1999 string = "ONE_JANUARY_1999"

	// PolicyDateDTODateTypeONEMAY1996 captures enum value "ONE_MAY_1996"
	PolicyDateDTODateTypeONEMAY1996 string = "ONE_MAY_1996"

	// PolicyDateDTODateTypeOPTIONDATE captures enum value "OPTION_DATE"
	PolicyDateDTODateTypeOPTIONDATE string = "OPTION_DATE"

	// PolicyDateDTODateTypeORIGINALOPTIONDATE captures enum value "ORIGINAL_OPTION_DATE"
	PolicyDateDTODateTypeORIGINALOPTIONDATE string = "ORIGINAL_OPTION_DATE"

	// PolicyDateDTODateTypeOVERRIDEPREMIUMBILLING captures enum value "OVERRIDE_PREMIUM_BILLING"
	PolicyDateDTODateTypeOVERRIDEPREMIUMBILLING string = "OVERRIDE_PREMIUM_BILLING"

	// PolicyDateDTODateTypePAIDTODATE captures enum value "PAID_TO_DATE"
	PolicyDateDTODateTypePAIDTODATE string = "PAID_TO_DATE"

	// PolicyDateDTODateTypePBATODPBACHANGEDATE captures enum value "PBA_TO_DPBA_CHANGE_DATE"
	PolicyDateDTODateTypePBATODPBACHANGEDATE string = "PBA_TO_DPBA_CHANGE_DATE"

	// PolicyDateDTODateTypePOLICYCEASE captures enum value "POLICY_CEASE"
	PolicyDateDTODateTypePOLICYCEASE string = "POLICY_CEASE"

	// PolicyDateDTODateTypePOSTINTERESTINITIALDATE captures enum value "POST_INTEREST_INITIAL_DATE"
	PolicyDateDTODateTypePOSTINTERESTINITIALDATE string = "POST_INTEREST_INITIAL_DATE"

	// PolicyDateDTODateTypePRIORINTERESTDECLARATIONPROCESSINGDATE captures enum value "PRIOR_INTEREST_DECLARATION_PROCESSING_DATE"
	PolicyDateDTODateTypePRIORINTERESTDECLARATIONPROCESSINGDATE string = "PRIOR_INTEREST_DECLARATION_PROCESSING_DATE"

	// PolicyDateDTODateTypePRIORYEARSEXPENSEBILLDUEDATE captures enum value "PRIOR_YEARS_EXPENSE_BILL_DUE_DATE"
	PolicyDateDTODateTypePRIORYEARSEXPENSEBILLDUEDATE string = "PRIOR_YEARS_EXPENSE_BILL_DUE_DATE"

	// PolicyDateDTODateTypePRIORYEARSINTERESTDECLARATIONDATE captures enum value "PRIOR_YEARS_INTEREST_DECLARATION_DATE"
	PolicyDateDTODateTypePRIORYEARSINTERESTDECLARATIONDATE string = "PRIOR_YEARS_INTEREST_DECLARATION_DATE"

	// PolicyDateDTODateTypePRIORYEARSRECOSTINGDATE captures enum value "PRIOR_YEARS_RECOSTING_DATE"
	PolicyDateDTODateTypePRIORYEARSRECOSTINGDATE string = "PRIOR_YEARS_RECOSTING_DATE"

	// PolicyDateDTODateTypePSOAPPROVALDATE captures enum value "PSO_APPROVAL_DATE"
	PolicyDateDTODateTypePSOAPPROVALDATE string = "PSO_APPROVAL_DATE"

	// PolicyDateDTODateTypeRATECHANGEDATE captures enum value "RATE_CHANGE_DATE"
	PolicyDateDTODateTypeRATECHANGEDATE string = "RATE_CHANGE_DATE"

	// PolicyDateDTODateTypeRATETRANCHEHISTORYDATE captures enum value "RATE_TRANCHE_HISTORY_DATE"
	PolicyDateDTODateTypeRATETRANCHEHISTORYDATE string = "RATE_TRANCHE_HISTORY_DATE"

	// PolicyDateDTODateTypeREALLOCATIONANDT captures enum value "REALLOCATION_ANDT"
	PolicyDateDTODateTypeREALLOCATIONANDT string = "REALLOCATION_ANDT"

	// PolicyDateDTODateTypeREALLOCATIONBASEDATE captures enum value "REALLOCATION_BASE_DATE"
	PolicyDateDTODateTypeREALLOCATIONBASEDATE string = "REALLOCATION_BASE_DATE"

	// PolicyDateDTODateTypeREALLOCATIONBCHG captures enum value "REALLOCATION_BCHG"
	PolicyDateDTODateTypeREALLOCATIONBCHG string = "REALLOCATION_BCHG"

	// PolicyDateDTODateTypeREALLOCATIONBEOP captures enum value "REALLOCATION_BEOP"
	PolicyDateDTODateTypeREALLOCATIONBEOP string = "REALLOCATION_BEOP"

	// PolicyDateDTODateTypeREALLOCATIONBNTR captures enum value "REALLOCATION_BNTR"
	PolicyDateDTODateTypeREALLOCATIONBNTR string = "REALLOCATION_BNTR"

	// PolicyDateDTODateTypeREALLOCATIONCSST captures enum value "REALLOCATION_CSST"
	PolicyDateDTODateTypeREALLOCATIONCSST string = "REALLOCATION_CSST"

	// PolicyDateDTODateTypeREALLOCATIONFCOR captures enum value "REALLOCATION_FCOR"
	PolicyDateDTODateTypeREALLOCATIONFCOR string = "REALLOCATION_FCOR"

	// PolicyDateDTODateTypeREALLOCATIONFDRD captures enum value "REALLOCATION_FDRD"
	PolicyDateDTODateTypeREALLOCATIONFDRD string = "REALLOCATION_FDRD"

	// PolicyDateDTODateTypeREALLOCATIONFRCR captures enum value "REALLOCATION_FRCR"
	PolicyDateDTODateTypeREALLOCATIONFRCR string = "REALLOCATION_FRCR"

	// PolicyDateDTODateTypeREALLOCATIONGLRD captures enum value "REALLOCATION_GLRD"
	PolicyDateDTODateTypeREALLOCATIONGLRD string = "REALLOCATION_GLRD"

	// PolicyDateDTODateTypeREALLOCATIONGRDT captures enum value "REALLOCATION_GRDT"
	PolicyDateDTODateTypeREALLOCATIONGRDT string = "REALLOCATION_GRDT"

	// PolicyDateDTODateTypeREALLOCATIONPMTD captures enum value "REALLOCATION_PMTD"
	PolicyDateDTODateTypeREALLOCATIONPMTD string = "REALLOCATION_PMTD"

	// PolicyDateDTODateTypeREALLOCATIONPMTR captures enum value "REALLOCATION_PMTR"
	PolicyDateDTODateTypeREALLOCATIONPMTR string = "REALLOCATION_PMTR"

	// PolicyDateDTODateTypeREALLOCATIONPOLICYMAINTENANCE captures enum value "REALLOCATION_POLICY_MAINTENANCE"
	PolicyDateDTODateTypeREALLOCATIONPOLICYMAINTENANCE string = "REALLOCATION_POLICY_MAINTENANCE"

	// PolicyDateDTODateTypeREALLOCATIONRENEWAL captures enum value "REALLOCATION_RENEWAL"
	PolicyDateDTODateTypeREALLOCATIONRENEWAL string = "REALLOCATION_RENEWAL"

	// PolicyDateDTODateTypeREALLOCATIONRERL captures enum value "REALLOCATION_RERL"
	PolicyDateDTODateTypeREALLOCATIONRERL string = "REALLOCATION_RERL"

	// PolicyDateDTODateTypeREALLOCATIONRKBS captures enum value "REALLOCATION_RKBS"
	PolicyDateDTODateTypeREALLOCATIONRKBS string = "REALLOCATION_RKBS"

	// PolicyDateDTODateTypeREALLOCATIONRKPS captures enum value "REALLOCATION_RKPS"
	PolicyDateDTODateTypeREALLOCATIONRKPS string = "REALLOCATION_RKPS"

	// PolicyDateDTODateTypeREALLOCATIONRKRL captures enum value "REALLOCATION_RKRL"
	PolicyDateDTODateTypeREALLOCATIONRKRL string = "REALLOCATION_RKRL"

	// PolicyDateDTODateTypeREALLOCATIONRRED captures enum value "REALLOCATION_RRED"
	PolicyDateDTODateTypeREALLOCATIONRRED string = "REALLOCATION_RRED"

	// PolicyDateDTODateTypeREALLOCATIONRTSC captures enum value "REALLOCATION_RTSC"
	PolicyDateDTODateTypeREALLOCATIONRTSC string = "REALLOCATION_RTSC"

	// PolicyDateDTODateTypeREALLOCATIONSCPR captures enum value "REALLOCATION_SCPR"
	PolicyDateDTODateTypeREALLOCATIONSCPR string = "REALLOCATION_SCPR"

	// PolicyDateDTODateTypeREALLOCATIONSNDT captures enum value "REALLOCATION_SNDT"
	PolicyDateDTODateTypeREALLOCATIONSNDT string = "REALLOCATION_SNDT"

	// PolicyDateDTODateTypeREALLOCATIONUNDC captures enum value "REALLOCATION_UNDC"
	PolicyDateDTODateTypeREALLOCATIONUNDC string = "REALLOCATION_UNDC"

	// PolicyDateDTODateTypeREALLOCATIONUWOT captures enum value "REALLOCATION_UWOT"
	PolicyDateDTODateTypeREALLOCATIONUWOT string = "REALLOCATION_UWOT"

	// PolicyDateDTODateTypeREASSURANCECONVERSIONDATE captures enum value "REASSURANCE_CONVERSION_DATE"
	PolicyDateDTODateTypeREASSURANCECONVERSIONDATE string = "REASSURANCE_CONVERSION_DATE"

	// PolicyDateDTODateTypeREPORTINGBALANCEDATE captures enum value "REPORTING_BALANCE_DATE"
	PolicyDateDTODateTypeREPORTINGBALANCEDATE string = "REPORTING_BALANCE_DATE"

	// PolicyDateDTODateTypeRESERVEDFORCLIENT captures enum value "RESERVED_FOR_CLIENT"
	PolicyDateDTODateTypeRESERVEDFORCLIENT string = "RESERVED_FOR_CLIENT"

	// PolicyDateDTODateTypeRESERVEDFORREALLOCATION captures enum value "RESERVED_FOR_REALLOCATION"
	PolicyDateDTODateTypeRESERVEDFORREALLOCATION string = "RESERVED_FOR_REALLOCATION"

	// PolicyDateDTODateTypeRFTCALCULATIONDATE captures enum value "RFT_CALCULATION_DATE"
	PolicyDateDTODateTypeRFTCALCULATIONDATE string = "RFT_CALCULATION_DATE"

	// PolicyDateDTODateTypeRISKBENEFITEXPIRYDATE captures enum value "RISK_BENEFIT_EXPIRY_DATE"
	PolicyDateDTODateTypeRISKBENEFITEXPIRYDATE string = "RISK_BENEFIT_EXPIRY_DATE"

	// PolicyDateDTODateTypeRISKGUARANTEEDATE captures enum value "RISK_GUARANTEE_DATE"
	PolicyDateDTODateTypeRISKGUARANTEEDATE string = "RISK_GUARANTEE_DATE"

	// PolicyDateDTODateTypeSCHEMECOMMENCEMENTDATE captures enum value "SCHEME_COMMENCEMENT_DATE"
	PolicyDateDTODateTypeSCHEMECOMMENCEMENTDATE string = "SCHEME_COMMENCEMENT_DATE"

	// PolicyDateDTODateTypeSECONDSHORTYEARSCHEMEDATE captures enum value "SECOND_SHORT_YEAR_SCHEME_DATE"
	PolicyDateDTODateTypeSECONDSHORTYEARSCHEMEDATE string = "SECOND_SHORT_YEAR_SCHEME_DATE"

	// PolicyDateDTODateTypeSHADOWCONVERSIONDATE captures enum value "SHADOW_CONVERSION_DATE"
	PolicyDateDTODateTypeSHADOWCONVERSIONDATE string = "SHADOW_CONVERSION_DATE"

	// PolicyDateDTODateTypeSHORT1STBILLPERIODENDDATE captures enum value "SHORT_1ST_BILL_PERIOD_END_DATE"
	PolicyDateDTODateTypeSHORT1STBILLPERIODENDDATE string = "SHORT_1ST_BILL_PERIOD_END_DATE"

	// PolicyDateDTODateTypeSPECIALCONVERSIONDATE captures enum value "SPECIAL_CONVERSION_DATE"
	PolicyDateDTODateTypeSPECIALCONVERSIONDATE string = "SPECIAL_CONVERSION_DATE"

	// PolicyDateDTODateTypeSPECIALMLCDATE captures enum value "SPECIAL_MLC_DATE"
	PolicyDateDTODateTypeSPECIALMLCDATE string = "SPECIAL_MLC_DATE"

	// PolicyDateDTODateTypeSUMOFTRANSBEGINDATEINCLUSIVE captures enum value "SUM_OF_TRANS_BEGIN_DATE_INCLUSIVE"
	PolicyDateDTODateTypeSUMOFTRANSBEGINDATEINCLUSIVE string = "SUM_OF_TRANS_BEGIN_DATE_INCLUSIVE"

	// PolicyDateDTODateTypeSUMOFTRANSENDDATEINCLUSIVE captures enum value "SUM_OF_TRANS_END_DATE_INCLUSIVE"
	PolicyDateDTODateTypeSUMOFTRANSENDDATEINCLUSIVE string = "SUM_OF_TRANS_END_DATE_INCLUSIVE"

	// PolicyDateDTODateTypeSWITCHTODUAL captures enum value "SWITCH_TO_DUAL"
	PolicyDateDTODateTypeSWITCHTODUAL string = "SWITCH_TO_DUAL"

	// PolicyDateDTODateTypeTABLECONVERSIONDATE captures enum value "TABLE_CONVERSION_DATE"
	PolicyDateDTODateTypeTABLECONVERSIONDATE string = "TABLE_CONVERSION_DATE"

	// PolicyDateDTODateTypeTHIRTYAPR96 captures enum value "THIRTY_APR_96"
	PolicyDateDTODateTypeTHIRTYAPR96 string = "THIRTY_APR_96"

	// PolicyDateDTODateTypeTHIRTYONEDECEMBER1988 captures enum value "THIRTY_ONE_DECEMBER_1988"
	PolicyDateDTODateTypeTHIRTYONEDECEMBER1988 string = "THIRTY_ONE_DECEMBER_1988"

	// PolicyDateDTODateTypeTRANSACTIONDETAILJOURNALLOGSAUDITDATE captures enum value "TRANSACTION_DETAIL_JOURNAL_LOGS_AUDIT_DATE"
	PolicyDateDTODateTypeTRANSACTIONDETAILJOURNALLOGSAUDITDATE string = "TRANSACTION_DETAIL_JOURNAL_LOGS_AUDIT_DATE"

	// PolicyDateDTODateTypeTRANSFERDATEFORCONVERSIONFUNDSONLY captures enum value "TRANSFER_DATE_FOR_CONVERSION_FUNDS_ONLY"
	PolicyDateDTODateTypeTRANSFERDATEFORCONVERSIONFUNDSONLY string = "TRANSFER_DATE_FOR_CONVERSION_FUNDS_ONLY"

	// PolicyDateDTODateTypeUSEFATUMAGEAPPROACHONORAFTERTHISDATE captures enum value "USE_FATUM_AGE_APPROACH_ON_OR_AFTER_THIS_DATE"
	PolicyDateDTODateTypeUSEFATUMAGEAPPROACHONORAFTERTHISDATE string = "USE_FATUM_AGE_APPROACH_ON_OR_AFTER_THIS_DATE"

	// PolicyDateDTODateTypeVALIDTARGETRETIREMENTDATEUSERDEFINED captures enum value "VALID_TARGET_RETIREMENT_DATE_USER_DEFINED"
	PolicyDateDTODateTypeVALIDTARGETRETIREMENTDATEUSERDEFINED string = "VALID_TARGET_RETIREMENT_DATE_USER_DEFINED"

	// PolicyDateDTODateTypeWRAPUPDATE captures enum value "WRAPUP_DATE"
	PolicyDateDTODateTypeWRAPUPDATE string = "WRAPUP_DATE"
)

// prop value enum
func (m *PolicyDateDTO) validateDateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, policyDateDTOTypeDateTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PolicyDateDTO) validateDateType(formats strfmt.Registry) error {
	if swag.IsZero(m.DateType) { // not required
		return nil
	}

	// value enum
	if err := m.validateDateTypeEnum("dateType", "body", m.DateType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this policy date d t o based on context it is used
func (m *PolicyDateDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PolicyDateDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PolicyDateDTO) UnmarshalBinary(b []byte) error {
	var res PolicyDateDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
