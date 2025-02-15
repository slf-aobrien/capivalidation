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

// AgencyDTOReq AgencyDTOReq
//
// swagger:model AgencyDTOReq
type AgencyDTOReq struct {

	// accreditation history list
	AccreditationHistoryList []*AccreditationHistoryDTO `json:"accreditationHistoryList"`

	// address
	Address *AddressDTO `json:"address,omitempty"`

	// address list
	AddressList []*AddressDTO `json:"addressList"`

	// agency entity typ code
	AgencyEntityTypCode string `json:"agencyEntityTypCode,omitempty"`

	// agency Id list
	AgencyIDList []*AgencyIDDTO `json:"agencyIdList"`

	// agency indemnity code
	AgencyIndemnityCode string `json:"agencyIndemnityCode,omitempty"`

	// agency national broker code
	AgencyNationalBrokerCode string `json:"agencyNationalBrokerCode,omitempty"`

	// agency number
	AgencyNumber string `json:"agencyNumber,omitempty"`

	// agency stat code
	AgencyStatCode string `json:"agencyStatCode,omitempty"`

	// agency type code
	AgencyTypeCode string `json:"agencyTypeCode,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// discount percent
	DiscountPercent string `json:"discountPercent,omitempty"`

	// email
	Email *EmailDTO `json:"email,omitempty"`

	// email list
	EmailList []*EmailDTO `json:"emailList"`

	// group office number
	GroupOfficeNumber string `json:"groupOfficeNumber,omitempty"`

	// indemnity effective date
	IndemnityEffectiveDate string `json:"indemnityEffectiveDate,omitempty"`

	// organization
	Organization *OrganizationDTO `json:"organization,omitempty"`

	// organization Id
	OrganizationID string `json:"organizationId,omitempty"`

	// person
	Person *PersonDTOReq `json:"person,omitempty"`

	// phone
	Phone *PhoneDTO `json:"phone,omitempty"`

	// phone list
	PhoneList []*PhoneDTO `json:"phoneList"`

	// sib effective date
	SibEffectiveDate string `json:"sibEffectiveDate,omitempty"`

	// sib expiration date
	SibExpirationDate string `json:"sibExpirationDate,omitempty"`
}

// Validate validates this agency d t o req
func (m *AgencyDTOReq) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccreditationHistoryList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAddressList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAgencyIDList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmailList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganization(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePerson(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhone(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhoneList(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AgencyDTOReq) validateAccreditationHistoryList(formats strfmt.Registry) error {
	if swag.IsZero(m.AccreditationHistoryList) { // not required
		return nil
	}

	for i := 0; i < len(m.AccreditationHistoryList); i++ {
		if swag.IsZero(m.AccreditationHistoryList[i]) { // not required
			continue
		}

		if m.AccreditationHistoryList[i] != nil {
			if err := m.AccreditationHistoryList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("accreditationHistoryList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("accreditationHistoryList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AgencyDTOReq) validateAddress(formats strfmt.Registry) error {
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

func (m *AgencyDTOReq) validateAddressList(formats strfmt.Registry) error {
	if swag.IsZero(m.AddressList) { // not required
		return nil
	}

	for i := 0; i < len(m.AddressList); i++ {
		if swag.IsZero(m.AddressList[i]) { // not required
			continue
		}

		if m.AddressList[i] != nil {
			if err := m.AddressList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("addressList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("addressList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AgencyDTOReq) validateAgencyIDList(formats strfmt.Registry) error {
	if swag.IsZero(m.AgencyIDList) { // not required
		return nil
	}

	for i := 0; i < len(m.AgencyIDList); i++ {
		if swag.IsZero(m.AgencyIDList[i]) { // not required
			continue
		}

		if m.AgencyIDList[i] != nil {
			if err := m.AgencyIDList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("agencyIdList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("agencyIdList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AgencyDTOReq) validateEmail(formats strfmt.Registry) error {
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

func (m *AgencyDTOReq) validateEmailList(formats strfmt.Registry) error {
	if swag.IsZero(m.EmailList) { // not required
		return nil
	}

	for i := 0; i < len(m.EmailList); i++ {
		if swag.IsZero(m.EmailList[i]) { // not required
			continue
		}

		if m.EmailList[i] != nil {
			if err := m.EmailList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("emailList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("emailList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AgencyDTOReq) validateOrganization(formats strfmt.Registry) error {
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

func (m *AgencyDTOReq) validatePerson(formats strfmt.Registry) error {
	if swag.IsZero(m.Person) { // not required
		return nil
	}

	if m.Person != nil {
		if err := m.Person.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("person")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("person")
			}
			return err
		}
	}

	return nil
}

func (m *AgencyDTOReq) validatePhone(formats strfmt.Registry) error {
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

func (m *AgencyDTOReq) validatePhoneList(formats strfmt.Registry) error {
	if swag.IsZero(m.PhoneList) { // not required
		return nil
	}

	for i := 0; i < len(m.PhoneList); i++ {
		if swag.IsZero(m.PhoneList[i]) { // not required
			continue
		}

		if m.PhoneList[i] != nil {
			if err := m.PhoneList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("phoneList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("phoneList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this agency d t o req based on the context it is used
func (m *AgencyDTOReq) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAccreditationHistoryList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAddress(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAddressList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAgencyIDList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEmail(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEmailList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOrganization(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePerson(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePhone(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePhoneList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AgencyDTOReq) contextValidateAccreditationHistoryList(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AccreditationHistoryList); i++ {

		if m.AccreditationHistoryList[i] != nil {

			if swag.IsZero(m.AccreditationHistoryList[i]) { // not required
				return nil
			}

			if err := m.AccreditationHistoryList[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("accreditationHistoryList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("accreditationHistoryList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AgencyDTOReq) contextValidateAddress(ctx context.Context, formats strfmt.Registry) error {

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

func (m *AgencyDTOReq) contextValidateAddressList(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AddressList); i++ {

		if m.AddressList[i] != nil {

			if swag.IsZero(m.AddressList[i]) { // not required
				return nil
			}

			if err := m.AddressList[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("addressList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("addressList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AgencyDTOReq) contextValidateAgencyIDList(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AgencyIDList); i++ {

		if m.AgencyIDList[i] != nil {

			if swag.IsZero(m.AgencyIDList[i]) { // not required
				return nil
			}

			if err := m.AgencyIDList[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("agencyIdList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("agencyIdList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AgencyDTOReq) contextValidateEmail(ctx context.Context, formats strfmt.Registry) error {

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

func (m *AgencyDTOReq) contextValidateEmailList(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.EmailList); i++ {

		if m.EmailList[i] != nil {

			if swag.IsZero(m.EmailList[i]) { // not required
				return nil
			}

			if err := m.EmailList[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("emailList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("emailList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AgencyDTOReq) contextValidateOrganization(ctx context.Context, formats strfmt.Registry) error {

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

func (m *AgencyDTOReq) contextValidatePerson(ctx context.Context, formats strfmt.Registry) error {

	if m.Person != nil {

		if swag.IsZero(m.Person) { // not required
			return nil
		}

		if err := m.Person.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("person")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("person")
			}
			return err
		}
	}

	return nil
}

func (m *AgencyDTOReq) contextValidatePhone(ctx context.Context, formats strfmt.Registry) error {

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

func (m *AgencyDTOReq) contextValidatePhoneList(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PhoneList); i++ {

		if m.PhoneList[i] != nil {

			if swag.IsZero(m.PhoneList[i]) { // not required
				return nil
			}

			if err := m.PhoneList[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("phoneList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("phoneList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AgencyDTOReq) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AgencyDTOReq) UnmarshalBinary(b []byte) error {
	var res AgencyDTOReq
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
