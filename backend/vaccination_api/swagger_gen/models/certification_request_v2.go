// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CertificationRequestV2 certification request v2
//
// swagger:model CertificationRequestV2
type CertificationRequestV2 struct {

	// facility
	Facility *CertificationRequestV2Facility `json:"facility,omitempty"`

	// meta
	Meta *CertificationRequestV2Meta `json:"meta,omitempty"`

	// pre enrollment code
	PreEnrollmentCode string `json:"preEnrollmentCode,omitempty"`

	// recipient
	Recipient *CertificationRequestV2Recipient `json:"recipient,omitempty"`

	// vaccination
	Vaccination *CertificationRequestV2Vaccination `json:"vaccination,omitempty"`

	// vaccinator
	Vaccinator *CertificationRequestV2Vaccinator `json:"vaccinator,omitempty"`
}

// Validate validates this certification request v2
func (m *CertificationRequestV2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFacility(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMeta(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecipient(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVaccination(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVaccinator(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CertificationRequestV2) validateFacility(formats strfmt.Registry) error {

	if swag.IsZero(m.Facility) { // not required
		return nil
	}

	if m.Facility != nil {
		if err := m.Facility.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("facility")
			}
			return err
		}
	}

	return nil
}

func (m *CertificationRequestV2) validateMeta(formats strfmt.Registry) error {

	if swag.IsZero(m.Meta) { // not required
		return nil
	}

	if m.Meta != nil {
		if err := m.Meta.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("meta")
			}
			return err
		}
	}

	return nil
}

func (m *CertificationRequestV2) validateRecipient(formats strfmt.Registry) error {

	if swag.IsZero(m.Recipient) { // not required
		return nil
	}

	if m.Recipient != nil {
		if err := m.Recipient.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("recipient")
			}
			return err
		}
	}

	return nil
}

func (m *CertificationRequestV2) validateVaccination(formats strfmt.Registry) error {

	if swag.IsZero(m.Vaccination) { // not required
		return nil
	}

	if m.Vaccination != nil {
		if err := m.Vaccination.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vaccination")
			}
			return err
		}
	}

	return nil
}

func (m *CertificationRequestV2) validateVaccinator(formats strfmt.Registry) error {

	if swag.IsZero(m.Vaccinator) { // not required
		return nil
	}

	if m.Vaccinator != nil {
		if err := m.Vaccinator.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vaccinator")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CertificationRequestV2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CertificationRequestV2) UnmarshalBinary(b []byte) error {
	var res CertificationRequestV2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CertificationRequestV2Facility certification request v2 facility
//
// swagger:model CertificationRequestV2Facility
type CertificationRequestV2Facility struct {

	// address
	Address *CertificationRequestV2FacilityAddress `json:"address,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// operator Id
	OperatorID string `json:"operatorId,omitempty"`

	// operator name
	OperatorName string `json:"operatorName,omitempty"`
}

// Validate validates this certification request v2 facility
func (m *CertificationRequestV2Facility) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CertificationRequestV2Facility) validateAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.Address) { // not required
		return nil
	}

	if m.Address != nil {
		if err := m.Address.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("facility" + "." + "address")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CertificationRequestV2Facility) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CertificationRequestV2Facility) UnmarshalBinary(b []byte) error {
	var res CertificationRequestV2Facility
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CertificationRequestV2FacilityAddress certification request v2 facility address
//
// swagger:model CertificationRequestV2FacilityAddress
type CertificationRequestV2FacilityAddress struct {

	// address line1
	AddressLine1 string `json:"addressLine1,omitempty"`

	// address line2
	AddressLine2 string `json:"addressLine2,omitempty"`

	// district
	District string `json:"district,omitempty"`

	// pincode
	Pincode int64 `json:"pincode,omitempty"`

	// state
	State string `json:"state,omitempty"`
}

// Validate validates this certification request v2 facility address
func (m *CertificationRequestV2FacilityAddress) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CertificationRequestV2FacilityAddress) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CertificationRequestV2FacilityAddress) UnmarshalBinary(b []byte) error {
	var res CertificationRequestV2FacilityAddress
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CertificationRequestV2Meta certification request v2 meta
//
// swagger:model CertificationRequestV2Meta
type CertificationRequestV2Meta struct {

	// CG - central govt, SG - State govt, PR - Private Facility
	// Enum: [CG SG PR]
	FacilityType string `json:"facilityType,omitempty"`

	// G for Govt, S for Self, V for Voucher, D for DBT, A for AyushmanBharat, I for Other Insurance
	// Enum: [G S V D A I]
	PaymentType string `json:"paymentType,omitempty"`

	// F for frontline officers, C for comorbidity category, R for regular others
	// Enum: [F C R]
	RegistrationCategory string `json:"registrationCategory,omitempty"`

	// Recipient data capture mode DE for Data Entry, SQ for Signed QR, RQ for Regular QR, PR for Pre Registration
	// Enum: [DE SQ RQ PR]
	RegistrationDataMode string `json:"registrationDataMode,omitempty"`

	// Vaccination session duration in minutes
	SessionDurationInMinutes int64 `json:"sessionDurationInMinutes,omitempty"`

	// upload timestamp
	// Format: date-time
	UploadTimestamp strfmt.DateTime `json:"uploadTimestamp,omitempty"`

	// vaccination app
	VaccinationApp *CertificationRequestV2MetaVaccinationApp `json:"vaccinationApp,omitempty"`

	// ID verification number of attempts (ex number of attempts done for Aadhaar)
	VerificationAttempts int64 `json:"verificationAttempts,omitempty"`

	// ID verification duration (duration in seconds for ID verification)
	VerificationDurationInSeconds int64 `json:"verificationDurationInSeconds,omitempty"`

	// Time between verification and vaccination (in minutes)
	WaitForVaccinationInMinutes int64 `json:"waitForVaccinationInMinutes,omitempty"`
}

// Validate validates this certification request v2 meta
func (m *CertificationRequestV2Meta) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFacilityType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaymentType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegistrationCategory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegistrationDataMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUploadTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVaccinationApp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var certificationRequestV2MetaTypeFacilityTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CG","SG","PR"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		certificationRequestV2MetaTypeFacilityTypePropEnum = append(certificationRequestV2MetaTypeFacilityTypePropEnum, v)
	}
}

const (

	// CertificationRequestV2MetaFacilityTypeCG captures enum value "CG"
	CertificationRequestV2MetaFacilityTypeCG string = "CG"

	// CertificationRequestV2MetaFacilityTypeSG captures enum value "SG"
	CertificationRequestV2MetaFacilityTypeSG string = "SG"

	// CertificationRequestV2MetaFacilityTypePR captures enum value "PR"
	CertificationRequestV2MetaFacilityTypePR string = "PR"
)

// prop value enum
func (m *CertificationRequestV2Meta) validateFacilityTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, certificationRequestV2MetaTypeFacilityTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CertificationRequestV2Meta) validateFacilityType(formats strfmt.Registry) error {

	if swag.IsZero(m.FacilityType) { // not required
		return nil
	}

	// value enum
	if err := m.validateFacilityTypeEnum("meta"+"."+"facilityType", "body", m.FacilityType); err != nil {
		return err
	}

	return nil
}

var certificationRequestV2MetaTypePaymentTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["G","S","V","D","A","I"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		certificationRequestV2MetaTypePaymentTypePropEnum = append(certificationRequestV2MetaTypePaymentTypePropEnum, v)
	}
}

const (

	// CertificationRequestV2MetaPaymentTypeG captures enum value "G"
	CertificationRequestV2MetaPaymentTypeG string = "G"

	// CertificationRequestV2MetaPaymentTypeS captures enum value "S"
	CertificationRequestV2MetaPaymentTypeS string = "S"

	// CertificationRequestV2MetaPaymentTypeV captures enum value "V"
	CertificationRequestV2MetaPaymentTypeV string = "V"

	// CertificationRequestV2MetaPaymentTypeD captures enum value "D"
	CertificationRequestV2MetaPaymentTypeD string = "D"

	// CertificationRequestV2MetaPaymentTypeA captures enum value "A"
	CertificationRequestV2MetaPaymentTypeA string = "A"

	// CertificationRequestV2MetaPaymentTypeI captures enum value "I"
	CertificationRequestV2MetaPaymentTypeI string = "I"
)

// prop value enum
func (m *CertificationRequestV2Meta) validatePaymentTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, certificationRequestV2MetaTypePaymentTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CertificationRequestV2Meta) validatePaymentType(formats strfmt.Registry) error {

	if swag.IsZero(m.PaymentType) { // not required
		return nil
	}

	// value enum
	if err := m.validatePaymentTypeEnum("meta"+"."+"paymentType", "body", m.PaymentType); err != nil {
		return err
	}

	return nil
}

var certificationRequestV2MetaTypeRegistrationCategoryPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["F","C","R"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		certificationRequestV2MetaTypeRegistrationCategoryPropEnum = append(certificationRequestV2MetaTypeRegistrationCategoryPropEnum, v)
	}
}

const (

	// CertificationRequestV2MetaRegistrationCategoryF captures enum value "F"
	CertificationRequestV2MetaRegistrationCategoryF string = "F"

	// CertificationRequestV2MetaRegistrationCategoryC captures enum value "C"
	CertificationRequestV2MetaRegistrationCategoryC string = "C"

	// CertificationRequestV2MetaRegistrationCategoryR captures enum value "R"
	CertificationRequestV2MetaRegistrationCategoryR string = "R"
)

// prop value enum
func (m *CertificationRequestV2Meta) validateRegistrationCategoryEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, certificationRequestV2MetaTypeRegistrationCategoryPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CertificationRequestV2Meta) validateRegistrationCategory(formats strfmt.Registry) error {

	if swag.IsZero(m.RegistrationCategory) { // not required
		return nil
	}

	// value enum
	if err := m.validateRegistrationCategoryEnum("meta"+"."+"registrationCategory", "body", m.RegistrationCategory); err != nil {
		return err
	}

	return nil
}

var certificationRequestV2MetaTypeRegistrationDataModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DE","SQ","RQ","PR"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		certificationRequestV2MetaTypeRegistrationDataModePropEnum = append(certificationRequestV2MetaTypeRegistrationDataModePropEnum, v)
	}
}

const (

	// CertificationRequestV2MetaRegistrationDataModeDE captures enum value "DE"
	CertificationRequestV2MetaRegistrationDataModeDE string = "DE"

	// CertificationRequestV2MetaRegistrationDataModeSQ captures enum value "SQ"
	CertificationRequestV2MetaRegistrationDataModeSQ string = "SQ"

	// CertificationRequestV2MetaRegistrationDataModeRQ captures enum value "RQ"
	CertificationRequestV2MetaRegistrationDataModeRQ string = "RQ"

	// CertificationRequestV2MetaRegistrationDataModePR captures enum value "PR"
	CertificationRequestV2MetaRegistrationDataModePR string = "PR"
)

// prop value enum
func (m *CertificationRequestV2Meta) validateRegistrationDataModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, certificationRequestV2MetaTypeRegistrationDataModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CertificationRequestV2Meta) validateRegistrationDataMode(formats strfmt.Registry) error {

	if swag.IsZero(m.RegistrationDataMode) { // not required
		return nil
	}

	// value enum
	if err := m.validateRegistrationDataModeEnum("meta"+"."+"registrationDataMode", "body", m.RegistrationDataMode); err != nil {
		return err
	}

	return nil
}

func (m *CertificationRequestV2Meta) validateUploadTimestamp(formats strfmt.Registry) error {

	if swag.IsZero(m.UploadTimestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("meta"+"."+"uploadTimestamp", "body", "date-time", m.UploadTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CertificationRequestV2Meta) validateVaccinationApp(formats strfmt.Registry) error {

	if swag.IsZero(m.VaccinationApp) { // not required
		return nil
	}

	if m.VaccinationApp != nil {
		if err := m.VaccinationApp.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("meta" + "." + "vaccinationApp")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CertificationRequestV2Meta) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CertificationRequestV2Meta) UnmarshalBinary(b []byte) error {
	var res CertificationRequestV2Meta
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CertificationRequestV2MetaVaccinationApp certification request v2 meta vaccination app
//
// swagger:model CertificationRequestV2MetaVaccinationApp
type CertificationRequestV2MetaVaccinationApp struct {

	// app mode
	// Enum: [Online Offline]
	AppMode string `json:"appMode,omitempty"`

	// W - Wifi, M - mobile data, L - LAN/WAN
	// Enum: [W M L]
	ConnectionType string `json:"connectionType,omitempty"`

	// Type of device, D - Desktop, M - Mobile, T - Tablet
	// Enum: [D M T]
	Device string `json:"device,omitempty"`

	// Type of Operating system on the device, W for Windows, A for Android, L for Linux, M for Mac, I for ios
	// Enum: [W A L M I]
	DeviceOS string `json:"deviceOS,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// o s version
	OSVersion string `json:"oSVersion,omitempty"`

	// type
	// Enum: [P M D]
	Type string `json:"type,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this certification request v2 meta vaccination app
func (m *CertificationRequestV2MetaVaccinationApp) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConnectionType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDevice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeviceOS(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var certificationRequestV2MetaVaccinationAppTypeAppModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Online","Offline"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		certificationRequestV2MetaVaccinationAppTypeAppModePropEnum = append(certificationRequestV2MetaVaccinationAppTypeAppModePropEnum, v)
	}
}

const (

	// CertificationRequestV2MetaVaccinationAppAppModeOnline captures enum value "Online"
	CertificationRequestV2MetaVaccinationAppAppModeOnline string = "Online"

	// CertificationRequestV2MetaVaccinationAppAppModeOffline captures enum value "Offline"
	CertificationRequestV2MetaVaccinationAppAppModeOffline string = "Offline"
)

// prop value enum
func (m *CertificationRequestV2MetaVaccinationApp) validateAppModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, certificationRequestV2MetaVaccinationAppTypeAppModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CertificationRequestV2MetaVaccinationApp) validateAppMode(formats strfmt.Registry) error {

	if swag.IsZero(m.AppMode) { // not required
		return nil
	}

	// value enum
	if err := m.validateAppModeEnum("meta"+"."+"vaccinationApp"+"."+"appMode", "body", m.AppMode); err != nil {
		return err
	}

	return nil
}

var certificationRequestV2MetaVaccinationAppTypeConnectionTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["W","M","L"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		certificationRequestV2MetaVaccinationAppTypeConnectionTypePropEnum = append(certificationRequestV2MetaVaccinationAppTypeConnectionTypePropEnum, v)
	}
}

const (

	// CertificationRequestV2MetaVaccinationAppConnectionTypeW captures enum value "W"
	CertificationRequestV2MetaVaccinationAppConnectionTypeW string = "W"

	// CertificationRequestV2MetaVaccinationAppConnectionTypeM captures enum value "M"
	CertificationRequestV2MetaVaccinationAppConnectionTypeM string = "M"

	// CertificationRequestV2MetaVaccinationAppConnectionTypeL captures enum value "L"
	CertificationRequestV2MetaVaccinationAppConnectionTypeL string = "L"
)

// prop value enum
func (m *CertificationRequestV2MetaVaccinationApp) validateConnectionTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, certificationRequestV2MetaVaccinationAppTypeConnectionTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CertificationRequestV2MetaVaccinationApp) validateConnectionType(formats strfmt.Registry) error {

	if swag.IsZero(m.ConnectionType) { // not required
		return nil
	}

	// value enum
	if err := m.validateConnectionTypeEnum("meta"+"."+"vaccinationApp"+"."+"connectionType", "body", m.ConnectionType); err != nil {
		return err
	}

	return nil
}

var certificationRequestV2MetaVaccinationAppTypeDevicePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["D","M","T"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		certificationRequestV2MetaVaccinationAppTypeDevicePropEnum = append(certificationRequestV2MetaVaccinationAppTypeDevicePropEnum, v)
	}
}

const (

	// CertificationRequestV2MetaVaccinationAppDeviceD captures enum value "D"
	CertificationRequestV2MetaVaccinationAppDeviceD string = "D"

	// CertificationRequestV2MetaVaccinationAppDeviceM captures enum value "M"
	CertificationRequestV2MetaVaccinationAppDeviceM string = "M"

	// CertificationRequestV2MetaVaccinationAppDeviceT captures enum value "T"
	CertificationRequestV2MetaVaccinationAppDeviceT string = "T"
)

// prop value enum
func (m *CertificationRequestV2MetaVaccinationApp) validateDeviceEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, certificationRequestV2MetaVaccinationAppTypeDevicePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CertificationRequestV2MetaVaccinationApp) validateDevice(formats strfmt.Registry) error {

	if swag.IsZero(m.Device) { // not required
		return nil
	}

	// value enum
	if err := m.validateDeviceEnum("meta"+"."+"vaccinationApp"+"."+"device", "body", m.Device); err != nil {
		return err
	}

	return nil
}

var certificationRequestV2MetaVaccinationAppTypeDeviceOSPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["W","A","L","M","I"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		certificationRequestV2MetaVaccinationAppTypeDeviceOSPropEnum = append(certificationRequestV2MetaVaccinationAppTypeDeviceOSPropEnum, v)
	}
}

const (

	// CertificationRequestV2MetaVaccinationAppDeviceOSW captures enum value "W"
	CertificationRequestV2MetaVaccinationAppDeviceOSW string = "W"

	// CertificationRequestV2MetaVaccinationAppDeviceOSA captures enum value "A"
	CertificationRequestV2MetaVaccinationAppDeviceOSA string = "A"

	// CertificationRequestV2MetaVaccinationAppDeviceOSL captures enum value "L"
	CertificationRequestV2MetaVaccinationAppDeviceOSL string = "L"

	// CertificationRequestV2MetaVaccinationAppDeviceOSM captures enum value "M"
	CertificationRequestV2MetaVaccinationAppDeviceOSM string = "M"

	// CertificationRequestV2MetaVaccinationAppDeviceOSI captures enum value "I"
	CertificationRequestV2MetaVaccinationAppDeviceOSI string = "I"
)

// prop value enum
func (m *CertificationRequestV2MetaVaccinationApp) validateDeviceOSEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, certificationRequestV2MetaVaccinationAppTypeDeviceOSPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CertificationRequestV2MetaVaccinationApp) validateDeviceOS(formats strfmt.Registry) error {

	if swag.IsZero(m.DeviceOS) { // not required
		return nil
	}

	// value enum
	if err := m.validateDeviceOSEnum("meta"+"."+"vaccinationApp"+"."+"deviceOS", "body", m.DeviceOS); err != nil {
		return err
	}

	return nil
}

var certificationRequestV2MetaVaccinationAppTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["P","M","D"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		certificationRequestV2MetaVaccinationAppTypeTypePropEnum = append(certificationRequestV2MetaVaccinationAppTypeTypePropEnum, v)
	}
}

const (

	// CertificationRequestV2MetaVaccinationAppTypeP captures enum value "P"
	CertificationRequestV2MetaVaccinationAppTypeP string = "P"

	// CertificationRequestV2MetaVaccinationAppTypeM captures enum value "M"
	CertificationRequestV2MetaVaccinationAppTypeM string = "M"

	// CertificationRequestV2MetaVaccinationAppTypeD captures enum value "D"
	CertificationRequestV2MetaVaccinationAppTypeD string = "D"
)

// prop value enum
func (m *CertificationRequestV2MetaVaccinationApp) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, certificationRequestV2MetaVaccinationAppTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CertificationRequestV2MetaVaccinationApp) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("meta"+"."+"vaccinationApp"+"."+"type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CertificationRequestV2MetaVaccinationApp) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CertificationRequestV2MetaVaccinationApp) UnmarshalBinary(b []byte) error {
	var res CertificationRequestV2MetaVaccinationApp
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CertificationRequestV2Recipient certification request v2 recipient
//
// swagger:model CertificationRequestV2Recipient
type CertificationRequestV2Recipient struct {

	// address
	Address *CertificationRequestV2RecipientAddress `json:"address,omitempty"`

	// age
	Age string `json:"age,omitempty"`

	// contact
	Contact []string `json:"contact"`

	// dob
	// Format: date
	Dob strfmt.Date `json:"dob,omitempty"`

	// gender
	Gender string `json:"gender,omitempty"`

	// identity
	Identity string `json:"identity,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// nationality
	Nationality string `json:"nationality,omitempty"`
}

// Validate validates this certification request v2 recipient
func (m *CertificationRequestV2Recipient) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDob(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CertificationRequestV2Recipient) validateAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.Address) { // not required
		return nil
	}

	if m.Address != nil {
		if err := m.Address.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("recipient" + "." + "address")
			}
			return err
		}
	}

	return nil
}

func (m *CertificationRequestV2Recipient) validateDob(formats strfmt.Registry) error {

	if swag.IsZero(m.Dob) { // not required
		return nil
	}

	if err := validate.FormatOf("recipient"+"."+"dob", "body", "date", m.Dob.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CertificationRequestV2Recipient) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CertificationRequestV2Recipient) UnmarshalBinary(b []byte) error {
	var res CertificationRequestV2Recipient
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CertificationRequestV2RecipientAddress certification request v2 recipient address
//
// swagger:model CertificationRequestV2RecipientAddress
type CertificationRequestV2RecipientAddress struct {

	// address line1
	AddressLine1 string `json:"addressLine1,omitempty"`

	// address line2
	AddressLine2 string `json:"addressLine2,omitempty"`

	// district
	District string `json:"district,omitempty"`

	// pincode
	Pincode int64 `json:"pincode,omitempty"`

	// state
	State string `json:"state,omitempty"`
}

// Validate validates this certification request v2 recipient address
func (m *CertificationRequestV2RecipientAddress) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CertificationRequestV2RecipientAddress) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CertificationRequestV2RecipientAddress) UnmarshalBinary(b []byte) error {
	var res CertificationRequestV2RecipientAddress
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CertificationRequestV2Vaccination certification request v2 vaccination
//
// swagger:model CertificationRequestV2Vaccination
type CertificationRequestV2Vaccination struct {

	// batch
	Batch string `json:"batch,omitempty"`

	// date
	// Format: date-time
	Date strfmt.DateTime `json:"date,omitempty"`

	// Dose number for example 1 for first dose of 2 doses
	Dose float64 `json:"dose,omitempty"`

	// effective start
	// Format: date
	EffectiveStart strfmt.Date `json:"effectiveStart,omitempty"`

	// effective until
	// Format: date
	EffectiveUntil strfmt.Date `json:"effectiveUntil,omitempty"`

	// manufacturer
	Manufacturer string `json:"manufacturer,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// Total number of doses required for this vaccination.
	TotalDoses float64 `json:"totalDoses,omitempty"`
}

// Validate validates this certification request v2 vaccination
func (m *CertificationRequestV2Vaccination) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEffectiveStart(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEffectiveUntil(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CertificationRequestV2Vaccination) validateDate(formats strfmt.Registry) error {

	if swag.IsZero(m.Date) { // not required
		return nil
	}

	if err := validate.FormatOf("vaccination"+"."+"date", "body", "date-time", m.Date.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CertificationRequestV2Vaccination) validateEffectiveStart(formats strfmt.Registry) error {

	if swag.IsZero(m.EffectiveStart) { // not required
		return nil
	}

	if err := validate.FormatOf("vaccination"+"."+"effectiveStart", "body", "date", m.EffectiveStart.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CertificationRequestV2Vaccination) validateEffectiveUntil(formats strfmt.Registry) error {

	if swag.IsZero(m.EffectiveUntil) { // not required
		return nil
	}

	if err := validate.FormatOf("vaccination"+"."+"effectiveUntil", "body", "date", m.EffectiveUntil.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CertificationRequestV2Vaccination) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CertificationRequestV2Vaccination) UnmarshalBinary(b []byte) error {
	var res CertificationRequestV2Vaccination
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CertificationRequestV2Vaccinator certification request v2 vaccinator
//
// swagger:model CertificationRequestV2Vaccinator
type CertificationRequestV2Vaccinator struct {

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this certification request v2 vaccinator
func (m *CertificationRequestV2Vaccinator) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CertificationRequestV2Vaccinator) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CertificationRequestV2Vaccinator) UnmarshalBinary(b []byte) error {
	var res CertificationRequestV2Vaccinator
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
