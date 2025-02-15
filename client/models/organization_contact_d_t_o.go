// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OrganizationContactDTO OrganizationContactDTO
//
// swagger:model OrganizationContactDTO
type OrganizationContactDTO struct {

	// address
	Address *AddressDTO `json:"address,omitempty"`

	// email
	Email *EmailDTO `json:"email,omitempty"`

	// organization
	Organization *OrganizationDTO `json:"organization,omitempty"`

	// phone
	Phone *PhoneDTO `json:"phone,omitempty"`
}

// Validate validates this organization contact d t o
func (m *OrganizationContactDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganization(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhone(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrganizationContactDTO) validateAddress(formats strfmt.Registry) error {
	if swag.IsZero(m.Address) { // not required
		return nil
	}

	if m.Address != nil {
		if err := m.Address.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("address")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("address")
			}
			return err
		}
	}

	return nil
}

func (m *OrganizationContactDTO) validateEmail(formats strfmt.Registry) error {
	if swag.IsZero(m.Email) { // not required
		return nil
	}

	if m.Email != nil {
		if err := m.Email.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("email")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("email")
			}
			return err
		}
	}

	return nil
}

func (m *OrganizationContactDTO) validateOrganization(formats strfmt.Registry) error {
	if swag.IsZero(m.Organization) { // not required
		return nil
	}

	if m.Organization != nil {
		if err := m.Organization.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("organization")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("organization")
			}
			return err
		}
	}

	return nil
}

func (m *OrganizationContactDTO) validatePhone(formats strfmt.Registry) error {
	if swag.IsZero(m.Phone) { // not required
		return nil
	}

	if m.Phone != nil {
		if err := m.Phone.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("phone")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("phone")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this organization contact d t o based on the context it is used
func (m *OrganizationContactDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAddress(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEmail(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOrganization(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePhone(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrganizationContactDTO) contextValidateAddress(ctx context.Context, formats strfmt.Registry) error {

	if m.Address != nil {

		if swag.IsZero(m.Address) { // not required
			return nil
		}

		if err := m.Address.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("address")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("address")
			}
			return err
		}
	}

	return nil
}

func (m *OrganizationContactDTO) contextValidateEmail(ctx context.Context, formats strfmt.Registry) error {

	if m.Email != nil {

		if swag.IsZero(m.Email) { // not required
			return nil
		}

		if err := m.Email.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("email")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("email")
			}
			return err
		}
	}

	return nil
}

func (m *OrganizationContactDTO) contextValidateOrganization(ctx context.Context, formats strfmt.Registry) error {

	if m.Organization != nil {

		if swag.IsZero(m.Organization) { // not required
			return nil
		}

		if err := m.Organization.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("organization")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("organization")
			}
			return err
		}
	}

	return nil
}

func (m *OrganizationContactDTO) contextValidatePhone(ctx context.Context, formats strfmt.Registry) error {

	if m.Phone != nil {

		if swag.IsZero(m.Phone) { // not required
			return nil
		}

		if err := m.Phone.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("phone")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("phone")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OrganizationContactDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrganizationContactDTO) UnmarshalBinary(b []byte) error {
	var res OrganizationContactDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
