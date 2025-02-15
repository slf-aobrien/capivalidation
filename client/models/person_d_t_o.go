// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PersonDTO PersonDTO
//
// swagger:model PersonDTO
type PersonDTO struct {

	// audit Id
	AuditID string `json:"auditId,omitempty"`

	// birth date
	BirthDate string `json:"birthDate,omitempty"`

	// birth date as date
	// Format: date-time
	BirthDateAsDate strfmt.DateTime `json:"birthDateAsDate,omitempty"`

	// client Id
	ClientID string `json:"clientId,omitempty"`

	// first name
	FirstName string `json:"firstName,omitempty"`

	// gender
	Gender string `json:"gender,omitempty"`

	// height
	Height string `json:"height,omitempty"`

	// last name
	LastName string `json:"lastName,omitempty"`

	// marital status
	MaritalStatus string `json:"maritalStatus,omitempty"`

	// middle name
	MiddleName string `json:"middleName,omitempty"`

	// name prefix
	NamePrefix string `json:"namePrefix,omitempty"`

	// name suffix
	NameSuffix string `json:"nameSuffix,omitempty"`

	// passport number
	PassportNumber string `json:"passportNumber,omitempty"`

	// preferred language
	PreferredLanguage string `json:"preferredLanguage,omitempty"`

	// social security number
	SocialSecurityNumber string `json:"socialSecurityNumber,omitempty"`

	// weight
	Weight string `json:"weight,omitempty"`
}

// Validate validates this person d t o
func (m *PersonDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBirthDateAsDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PersonDTO) validateBirthDateAsDate(formats strfmt.Registry) error {
	if swag.IsZero(m.BirthDateAsDate) { // not required
		return nil
	}

	if err := validate.FormatOf("birthDateAsDate", "body", "date-time", m.BirthDateAsDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this person d t o based on context it is used
func (m *PersonDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PersonDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PersonDTO) UnmarshalBinary(b []byte) error {
	var res PersonDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
