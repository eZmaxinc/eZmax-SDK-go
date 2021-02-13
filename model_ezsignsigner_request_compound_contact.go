/*
 * eZmax API Definition
 *
 * This API expose all the functionnalities for the eZmax and eZsign application.
 *
 * API version: 1.0.30
 * Contact: support-api@ezmax.ca
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxinc/eZmax-SDK-go

import (
	"encoding/json"
)

// EzsignsignerRequestCompoundContact A Ezsignsigner->Contact Object and children to create a complete structure
type EzsignsignerRequestCompoundContact struct {
	// The first name of the Contact
	SContactFirstname string `json:"sContactFirstname"`
	// The last name of the Contact
	SContactLastname string `json:"sContactLastname"`
	// The unique ID of the Language.  Valid values:  |Value|Description| |-|-| |1|French| |2|English|
	FkiLanguageID int32 `json:"fkiLanguageID"`
	// The email address of the contact. Must be filled if email authentification was requested
	SEmailAddress *string `json:"sEmailAddress,omitempty"`
	// The Phone number of the contact. Use format \"5149901516\" for North American Numbers (Without \"1\" for long distance code) you would dial like this: 1-514-990-1516. Use format \"498945233886\" for international numbers (Without \"011\") you would dial like this: +49 89 452 33 88-6. In this example \"49\" is the country code of Germany.
	SPhoneNumber *string `json:"sPhoneNumber,omitempty"`
	// The Cell Phone number of the contact. Use format \"5149901516\" for North American Numbers (Without \"1\" for long distance code) you would dial like this: 1-514-990-1516. Use format \"498945233886\" for international numbers (Without \"011\") you would dial like this: +49 89 452 33 88-6. In this example \"49\" is the country code of Germany.
	SPhoneNumberCell *string `json:"sPhoneNumberCell,omitempty"`
}

// NewEzsignsignerRequestCompoundContact instantiates a new EzsignsignerRequestCompoundContact object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignsignerRequestCompoundContact(sContactFirstname string, sContactLastname string, fkiLanguageID int32, ) *EzsignsignerRequestCompoundContact {
	this := EzsignsignerRequestCompoundContact{}
	this.SContactFirstname = sContactFirstname
	this.SContactLastname = sContactLastname
	this.FkiLanguageID = fkiLanguageID
	return &this
}

// NewEzsignsignerRequestCompoundContactWithDefaults instantiates a new EzsignsignerRequestCompoundContact object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignsignerRequestCompoundContactWithDefaults() *EzsignsignerRequestCompoundContact {
	this := EzsignsignerRequestCompoundContact{}
	return &this
}

// GetSContactFirstname returns the SContactFirstname field value
func (o *EzsignsignerRequestCompoundContact) GetSContactFirstname() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.SContactFirstname
}

// GetSContactFirstnameOk returns a tuple with the SContactFirstname field value
// and a boolean to check if the value has been set.
func (o *EzsignsignerRequestCompoundContact) GetSContactFirstnameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SContactFirstname, true
}

// SetSContactFirstname sets field value
func (o *EzsignsignerRequestCompoundContact) SetSContactFirstname(v string) {
	o.SContactFirstname = v
}

// GetSContactLastname returns the SContactLastname field value
func (o *EzsignsignerRequestCompoundContact) GetSContactLastname() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.SContactLastname
}

// GetSContactLastnameOk returns a tuple with the SContactLastname field value
// and a boolean to check if the value has been set.
func (o *EzsignsignerRequestCompoundContact) GetSContactLastnameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SContactLastname, true
}

// SetSContactLastname sets field value
func (o *EzsignsignerRequestCompoundContact) SetSContactLastname(v string) {
	o.SContactLastname = v
}

// GetFkiLanguageID returns the FkiLanguageID field value
func (o *EzsignsignerRequestCompoundContact) GetFkiLanguageID() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.FkiLanguageID
}

// GetFkiLanguageIDOk returns a tuple with the FkiLanguageID field value
// and a boolean to check if the value has been set.
func (o *EzsignsignerRequestCompoundContact) GetFkiLanguageIDOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FkiLanguageID, true
}

// SetFkiLanguageID sets field value
func (o *EzsignsignerRequestCompoundContact) SetFkiLanguageID(v int32) {
	o.FkiLanguageID = v
}

// GetSEmailAddress returns the SEmailAddress field value if set, zero value otherwise.
func (o *EzsignsignerRequestCompoundContact) GetSEmailAddress() string {
	if o == nil || o.SEmailAddress == nil {
		var ret string
		return ret
	}
	return *o.SEmailAddress
}

// GetSEmailAddressOk returns a tuple with the SEmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignsignerRequestCompoundContact) GetSEmailAddressOk() (*string, bool) {
	if o == nil || o.SEmailAddress == nil {
		return nil, false
	}
	return o.SEmailAddress, true
}

// HasSEmailAddress returns a boolean if a field has been set.
func (o *EzsignsignerRequestCompoundContact) HasSEmailAddress() bool {
	if o != nil && o.SEmailAddress != nil {
		return true
	}

	return false
}

// SetSEmailAddress gets a reference to the given string and assigns it to the SEmailAddress field.
func (o *EzsignsignerRequestCompoundContact) SetSEmailAddress(v string) {
	o.SEmailAddress = &v
}

// GetSPhoneNumber returns the SPhoneNumber field value if set, zero value otherwise.
func (o *EzsignsignerRequestCompoundContact) GetSPhoneNumber() string {
	if o == nil || o.SPhoneNumber == nil {
		var ret string
		return ret
	}
	return *o.SPhoneNumber
}

// GetSPhoneNumberOk returns a tuple with the SPhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignsignerRequestCompoundContact) GetSPhoneNumberOk() (*string, bool) {
	if o == nil || o.SPhoneNumber == nil {
		return nil, false
	}
	return o.SPhoneNumber, true
}

// HasSPhoneNumber returns a boolean if a field has been set.
func (o *EzsignsignerRequestCompoundContact) HasSPhoneNumber() bool {
	if o != nil && o.SPhoneNumber != nil {
		return true
	}

	return false
}

// SetSPhoneNumber gets a reference to the given string and assigns it to the SPhoneNumber field.
func (o *EzsignsignerRequestCompoundContact) SetSPhoneNumber(v string) {
	o.SPhoneNumber = &v
}

// GetSPhoneNumberCell returns the SPhoneNumberCell field value if set, zero value otherwise.
func (o *EzsignsignerRequestCompoundContact) GetSPhoneNumberCell() string {
	if o == nil || o.SPhoneNumberCell == nil {
		var ret string
		return ret
	}
	return *o.SPhoneNumberCell
}

// GetSPhoneNumberCellOk returns a tuple with the SPhoneNumberCell field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignsignerRequestCompoundContact) GetSPhoneNumberCellOk() (*string, bool) {
	if o == nil || o.SPhoneNumberCell == nil {
		return nil, false
	}
	return o.SPhoneNumberCell, true
}

// HasSPhoneNumberCell returns a boolean if a field has been set.
func (o *EzsignsignerRequestCompoundContact) HasSPhoneNumberCell() bool {
	if o != nil && o.SPhoneNumberCell != nil {
		return true
	}

	return false
}

// SetSPhoneNumberCell gets a reference to the given string and assigns it to the SPhoneNumberCell field.
func (o *EzsignsignerRequestCompoundContact) SetSPhoneNumberCell(v string) {
	o.SPhoneNumberCell = &v
}

func (o EzsignsignerRequestCompoundContact) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sContactFirstname"] = o.SContactFirstname
	}
	if true {
		toSerialize["sContactLastname"] = o.SContactLastname
	}
	if true {
		toSerialize["fkiLanguageID"] = o.FkiLanguageID
	}
	if o.SEmailAddress != nil {
		toSerialize["sEmailAddress"] = o.SEmailAddress
	}
	if o.SPhoneNumber != nil {
		toSerialize["sPhoneNumber"] = o.SPhoneNumber
	}
	if o.SPhoneNumberCell != nil {
		toSerialize["sPhoneNumberCell"] = o.SPhoneNumberCell
	}
	return json.Marshal(toSerialize)
}

type NullableEzsignsignerRequestCompoundContact struct {
	value *EzsignsignerRequestCompoundContact
	isSet bool
}

func (v NullableEzsignsignerRequestCompoundContact) Get() *EzsignsignerRequestCompoundContact {
	return v.value
}

func (v *NullableEzsignsignerRequestCompoundContact) Set(val *EzsignsignerRequestCompoundContact) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignsignerRequestCompoundContact) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignsignerRequestCompoundContact) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignsignerRequestCompoundContact(val *EzsignsignerRequestCompoundContact) *NullableEzsignsignerRequestCompoundContact {
	return &NullableEzsignsignerRequestCompoundContact{value: val, isSet: true}
}

func (v NullableEzsignsignerRequestCompoundContact) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignsignerRequestCompoundContact) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


