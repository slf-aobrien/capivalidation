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
	"github.com/go-openapi/validate"
)

// ResponseWrapperListBillGroupRuleDTO ResponseWrapper«List«BillGroupRuleDTO»»
//
// swagger:model ResponseWrapper«List«BillGroupRuleDTO»»
type ResponseWrapperListBillGroupRuleDTO struct {

	// fail
	Fail bool `json:"fail,omitempty"`

	// http response code
	HTTPResponseCode int32 `json:"httpResponseCode,omitempty"`

	// payload
	Payload []*BillGroupRuleDTORes `json:"payload"`

	// response code
	ResponseCode int32 `json:"responseCode,omitempty"`

	// response message
	ResponseMessage string `json:"responseMessage,omitempty"`

	// success
	Success bool `json:"success,omitempty"`

	// total elements
	TotalElements int64 `json:"totalElements,omitempty"`

	// total pages
	TotalPages int64 `json:"totalPages,omitempty"`

	// uuid
	UUID string `json:"uuid,omitempty"`

	// variance detail map
	VarianceDetailMap map[string][]VarianceDetailDTO `json:"varianceDetailMap,omitempty"`

	// variance map
	VarianceMap map[string][]VarianceDTO `json:"varianceMap,omitempty"`

	// warn
	Warn bool `json:"warn,omitempty"`
}

// Validate validates this response wrapper list bill group rule d t o
func (m *ResponseWrapperListBillGroupRuleDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePayload(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVarianceDetailMap(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVarianceMap(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResponseWrapperListBillGroupRuleDTO) validatePayload(formats strfmt.Registry) error {
	if swag.IsZero(m.Payload) { // not required
		return nil
	}

	for i := 0; i < len(m.Payload); i++ {
		if swag.IsZero(m.Payload[i]) { // not required
			continue
		}

		if m.Payload[i] != nil {
			if err := m.Payload[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("payload" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("payload" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ResponseWrapperListBillGroupRuleDTO) validateVarianceDetailMap(formats strfmt.Registry) error {
	if swag.IsZero(m.VarianceDetailMap) { // not required
		return nil
	}

	for k := range m.VarianceDetailMap {

		if err := validate.Required("varianceDetailMap"+"."+k, "body", m.VarianceDetailMap[k]); err != nil {
			return err
		}

		if err := validate.UniqueItems("varianceDetailMap"+"."+k, "body", m.VarianceDetailMap[k]); err != nil {
			return err
		}

		for i := 0; i < len(m.VarianceDetailMap[k]); i++ {

			if err := m.VarianceDetailMap[k][i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("varianceDetailMap" + "." + k + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("varianceDetailMap" + "." + k + "." + strconv.Itoa(i))
				}
				return err
			}

		}

	}

	return nil
}

func (m *ResponseWrapperListBillGroupRuleDTO) validateVarianceMap(formats strfmt.Registry) error {
	if swag.IsZero(m.VarianceMap) { // not required
		return nil
	}

	for k := range m.VarianceMap {

		if err := validate.Required("varianceMap"+"."+k, "body", m.VarianceMap[k]); err != nil {
			return err
		}

		if err := validate.UniqueItems("varianceMap"+"."+k, "body", m.VarianceMap[k]); err != nil {
			return err
		}

		for i := 0; i < len(m.VarianceMap[k]); i++ {

			if err := m.VarianceMap[k][i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("varianceMap" + "." + k + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("varianceMap" + "." + k + "." + strconv.Itoa(i))
				}
				return err
			}

		}

	}

	return nil
}

// ContextValidate validate this response wrapper list bill group rule d t o based on the context it is used
func (m *ResponseWrapperListBillGroupRuleDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePayload(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVarianceDetailMap(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVarianceMap(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResponseWrapperListBillGroupRuleDTO) contextValidatePayload(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Payload); i++ {

		if m.Payload[i] != nil {

			if swag.IsZero(m.Payload[i]) { // not required
				return nil
			}

			if err := m.Payload[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("payload" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("payload" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ResponseWrapperListBillGroupRuleDTO) contextValidateVarianceDetailMap(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.VarianceDetailMap {

		for i := 0; i < len(m.VarianceDetailMap[k]); i++ {

			if swag.IsZero(m.VarianceDetailMap[k][i]) { // not required
				return nil
			}

			if err := m.VarianceDetailMap[k][i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("varianceDetailMap" + "." + k + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("varianceDetailMap" + "." + k + "." + strconv.Itoa(i))
				}
				return err
			}

		}

	}

	return nil
}

func (m *ResponseWrapperListBillGroupRuleDTO) contextValidateVarianceMap(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.VarianceMap {

		for i := 0; i < len(m.VarianceMap[k]); i++ {

			if swag.IsZero(m.VarianceMap[k][i]) { // not required
				return nil
			}

			if err := m.VarianceMap[k][i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("varianceMap" + "." + k + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("varianceMap" + "." + k + "." + strconv.Itoa(i))
				}
				return err
			}

		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ResponseWrapperListBillGroupRuleDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResponseWrapperListBillGroupRuleDTO) UnmarshalBinary(b []byte) error {
	var res ResponseWrapperListBillGroupRuleDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
