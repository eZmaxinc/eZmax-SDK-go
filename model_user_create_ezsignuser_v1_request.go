/*
eZmax API Definition

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.1.3
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
)

// UserCreateEzsignuserV1Request Request for the /1/module/user/createEzsignuser API Request
type UserCreateEzsignuserV1Request struct {
	// The unique ID of the Language.  Valid values:  |Value|Description| |-|-| |1|French| |2|English|
	FkiLanguageID int32 `json:"fkiLanguageID"`
	// The First name of the user
	SUserFirstname string `json:"sUserFirstname"`
	// The Last name of the user
	SUserLastname string `json:"sUserLastname"`
	// The email address.
	SEmailAddress string `json:"sEmailAddress"`
	// The region of the phone number. (For a North America Number only)  The region is the \"514\" section in this sample phone number: (514) 990-1516 x123
	SPhoneRegion string `json:"sPhoneRegion"`
	// The exchange of the phone number. (For a North America Number only)  The exchange is the \"990\" section in this sample phone number: (514) 990-1516 x123
	SPhoneExchange string `json:"sPhoneExchange"`
	// The number of the phone number. (For a North America Number only)  The number is the \"1516\" section in this sample phone number: (514) 990-1516 x123
	SPhoneNumber string `json:"sPhoneNumber"`
	// The extension of the phone number.  The extension is the \"123\" section in this sample phone number: (514) 990-1516 x123.  It can also be used with international phone numbers
	SPhoneExtension *string `json:"sPhoneExtension,omitempty"`
}

// NewUserCreateEzsignuserV1Request instantiates a new UserCreateEzsignuserV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserCreateEzsignuserV1Request(fkiLanguageID int32, sUserFirstname string, sUserLastname string, sEmailAddress string, sPhoneRegion string, sPhoneExchange string, sPhoneNumber string) *UserCreateEzsignuserV1Request {
	this := UserCreateEzsignuserV1Request{}
	this.FkiLanguageID = fkiLanguageID
	this.SUserFirstname = sUserFirstname
	this.SUserLastname = sUserLastname
	this.SEmailAddress = sEmailAddress
	this.SPhoneRegion = sPhoneRegion
	this.SPhoneExchange = sPhoneExchange
	this.SPhoneNumber = sPhoneNumber
	return &this
}

// NewUserCreateEzsignuserV1RequestWithDefaults instantiates a new UserCreateEzsignuserV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserCreateEzsignuserV1RequestWithDefaults() *UserCreateEzsignuserV1Request {
	this := UserCreateEzsignuserV1Request{}
	return &this
}

// GetFkiLanguageID returns the FkiLanguageID field value
func (o *UserCreateEzsignuserV1Request) GetFkiLanguageID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiLanguageID
}

// GetFkiLanguageIDOk returns a tuple with the FkiLanguageID field value
// and a boolean to check if the value has been set.
func (o *UserCreateEzsignuserV1Request) GetFkiLanguageIDOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FkiLanguageID, true
}

// SetFkiLanguageID sets field value
func (o *UserCreateEzsignuserV1Request) SetFkiLanguageID(v int32) {
	o.FkiLanguageID = v
}

// GetSUserFirstname returns the SUserFirstname field value
func (o *UserCreateEzsignuserV1Request) GetSUserFirstname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SUserFirstname
}

// GetSUserFirstnameOk returns a tuple with the SUserFirstname field value
// and a boolean to check if the value has been set.
func (o *UserCreateEzsignuserV1Request) GetSUserFirstnameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SUserFirstname, true
}

// SetSUserFirstname sets field value
func (o *UserCreateEzsignuserV1Request) SetSUserFirstname(v string) {
	o.SUserFirstname = v
}

// GetSUserLastname returns the SUserLastname field value
func (o *UserCreateEzsignuserV1Request) GetSUserLastname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SUserLastname
}

// GetSUserLastnameOk returns a tuple with the SUserLastname field value
// and a boolean to check if the value has been set.
func (o *UserCreateEzsignuserV1Request) GetSUserLastnameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SUserLastname, true
}

// SetSUserLastname sets field value
func (o *UserCreateEzsignuserV1Request) SetSUserLastname(v string) {
	o.SUserLastname = v
}

// GetSEmailAddress returns the SEmailAddress field value
func (o *UserCreateEzsignuserV1Request) GetSEmailAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEmailAddress
}

// GetSEmailAddressOk returns a tuple with the SEmailAddress field value
// and a boolean to check if the value has been set.
func (o *UserCreateEzsignuserV1Request) GetSEmailAddressOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SEmailAddress, true
}

// SetSEmailAddress sets field value
func (o *UserCreateEzsignuserV1Request) SetSEmailAddress(v string) {
	o.SEmailAddress = v
}

// GetSPhoneRegion returns the SPhoneRegion field value
func (o *UserCreateEzsignuserV1Request) GetSPhoneRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SPhoneRegion
}

// GetSPhoneRegionOk returns a tuple with the SPhoneRegion field value
// and a boolean to check if the value has been set.
func (o *UserCreateEzsignuserV1Request) GetSPhoneRegionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SPhoneRegion, true
}

// SetSPhoneRegion sets field value
func (o *UserCreateEzsignuserV1Request) SetSPhoneRegion(v string) {
	o.SPhoneRegion = v
}

// GetSPhoneExchange returns the SPhoneExchange field value
func (o *UserCreateEzsignuserV1Request) GetSPhoneExchange() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SPhoneExchange
}

// GetSPhoneExchangeOk returns a tuple with the SPhoneExchange field value
// and a boolean to check if the value has been set.
func (o *UserCreateEzsignuserV1Request) GetSPhoneExchangeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SPhoneExchange, true
}

// SetSPhoneExchange sets field value
func (o *UserCreateEzsignuserV1Request) SetSPhoneExchange(v string) {
	o.SPhoneExchange = v
}

// GetSPhoneNumber returns the SPhoneNumber field value
func (o *UserCreateEzsignuserV1Request) GetSPhoneNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SPhoneNumber
}

// GetSPhoneNumberOk returns a tuple with the SPhoneNumber field value
// and a boolean to check if the value has been set.
func (o *UserCreateEzsignuserV1Request) GetSPhoneNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SPhoneNumber, true
}

// SetSPhoneNumber sets field value
func (o *UserCreateEzsignuserV1Request) SetSPhoneNumber(v string) {
	o.SPhoneNumber = v
}

// GetSPhoneExtension returns the SPhoneExtension field value if set, zero value otherwise.
func (o *UserCreateEzsignuserV1Request) GetSPhoneExtension() string {
	if o == nil || o.SPhoneExtension == nil {
		var ret string
		return ret
	}
	return *o.SPhoneExtension
}

// GetSPhoneExtensionOk returns a tuple with the SPhoneExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCreateEzsignuserV1Request) GetSPhoneExtensionOk() (*string, bool) {
	if o == nil || o.SPhoneExtension == nil {
		return nil, false
	}
	return o.SPhoneExtension, true
}

// HasSPhoneExtension returns a boolean if a field has been set.
func (o *UserCreateEzsignuserV1Request) HasSPhoneExtension() bool {
	if o != nil && o.SPhoneExtension != nil {
		return true
	}

	return false
}

// SetSPhoneExtension gets a reference to the given string and assigns it to the SPhoneExtension field.
func (o *UserCreateEzsignuserV1Request) SetSPhoneExtension(v string) {
	o.SPhoneExtension = &v
}

func (o UserCreateEzsignuserV1Request) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["fkiLanguageID"] = o.FkiLanguageID
	}
	if true {
		toSerialize["sUserFirstname"] = o.SUserFirstname
	}
	if true {
		toSerialize["sUserLastname"] = o.SUserLastname
	}
	if true {
		toSerialize["sEmailAddress"] = o.SEmailAddress
	}
	if true {
		toSerialize["sPhoneRegion"] = o.SPhoneRegion
	}
	if true {
		toSerialize["sPhoneExchange"] = o.SPhoneExchange
	}
	if true {
		toSerialize["sPhoneNumber"] = o.SPhoneNumber
	}
	if o.SPhoneExtension != nil {
		toSerialize["sPhoneExtension"] = o.SPhoneExtension
	}
	return json.Marshal(toSerialize)
}

type NullableUserCreateEzsignuserV1Request struct {
	value *UserCreateEzsignuserV1Request
	isSet bool
}

func (v NullableUserCreateEzsignuserV1Request) Get() *UserCreateEzsignuserV1Request {
	return v.value
}

func (v *NullableUserCreateEzsignuserV1Request) Set(val *UserCreateEzsignuserV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableUserCreateEzsignuserV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableUserCreateEzsignuserV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserCreateEzsignuserV1Request(val *UserCreateEzsignuserV1Request) *NullableUserCreateEzsignuserV1Request {
	return &NullableUserCreateEzsignuserV1Request{value: val, isSet: true}
}

func (v NullableUserCreateEzsignuserV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserCreateEzsignuserV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


