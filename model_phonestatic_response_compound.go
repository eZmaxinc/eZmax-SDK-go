/*
eZmax API Definition (Full)

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.2.0
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
)

// checks if the PhonestaticResponseCompound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PhonestaticResponseCompound{}

// PhonestaticResponseCompound A Phonestatic Object and children to create a complete structure
type PhonestaticResponseCompound struct {
	// The unique ID of the Phone.
	PkiPhonestaticID int32 `json:"pkiPhonestaticID"`
	// A phone number in E.164 Format
	SPhonestaticE164 *string `json:"sPhonestaticE164,omitempty"`
	// The extension of the phone number.
	SPhonestaticExtension *string `json:"sPhonestaticExtension,omitempty"`
}

// NewPhonestaticResponseCompound instantiates a new PhonestaticResponseCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPhonestaticResponseCompound(pkiPhonestaticID int32) *PhonestaticResponseCompound {
	this := PhonestaticResponseCompound{}
	this.PkiPhonestaticID = pkiPhonestaticID
	return &this
}

// NewPhonestaticResponseCompoundWithDefaults instantiates a new PhonestaticResponseCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPhonestaticResponseCompoundWithDefaults() *PhonestaticResponseCompound {
	this := PhonestaticResponseCompound{}
	return &this
}

// GetPkiPhonestaticID returns the PkiPhonestaticID field value
func (o *PhonestaticResponseCompound) GetPkiPhonestaticID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiPhonestaticID
}

// GetPkiPhonestaticIDOk returns a tuple with the PkiPhonestaticID field value
// and a boolean to check if the value has been set.
func (o *PhonestaticResponseCompound) GetPkiPhonestaticIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiPhonestaticID, true
}

// SetPkiPhonestaticID sets field value
func (o *PhonestaticResponseCompound) SetPkiPhonestaticID(v int32) {
	o.PkiPhonestaticID = v
}

// GetSPhonestaticE164 returns the SPhonestaticE164 field value if set, zero value otherwise.
func (o *PhonestaticResponseCompound) GetSPhonestaticE164() string {
	if o == nil || IsNil(o.SPhonestaticE164) {
		var ret string
		return ret
	}
	return *o.SPhonestaticE164
}

// GetSPhonestaticE164Ok returns a tuple with the SPhonestaticE164 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhonestaticResponseCompound) GetSPhonestaticE164Ok() (*string, bool) {
	if o == nil || IsNil(o.SPhonestaticE164) {
		return nil, false
	}
	return o.SPhonestaticE164, true
}

// HasSPhonestaticE164 returns a boolean if a field has been set.
func (o *PhonestaticResponseCompound) HasSPhonestaticE164() bool {
	if o != nil && !IsNil(o.SPhonestaticE164) {
		return true
	}

	return false
}

// SetSPhonestaticE164 gets a reference to the given string and assigns it to the SPhonestaticE164 field.
func (o *PhonestaticResponseCompound) SetSPhonestaticE164(v string) {
	o.SPhonestaticE164 = &v
}

// GetSPhonestaticExtension returns the SPhonestaticExtension field value if set, zero value otherwise.
func (o *PhonestaticResponseCompound) GetSPhonestaticExtension() string {
	if o == nil || IsNil(o.SPhonestaticExtension) {
		var ret string
		return ret
	}
	return *o.SPhonestaticExtension
}

// GetSPhonestaticExtensionOk returns a tuple with the SPhonestaticExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhonestaticResponseCompound) GetSPhonestaticExtensionOk() (*string, bool) {
	if o == nil || IsNil(o.SPhonestaticExtension) {
		return nil, false
	}
	return o.SPhonestaticExtension, true
}

// HasSPhonestaticExtension returns a boolean if a field has been set.
func (o *PhonestaticResponseCompound) HasSPhonestaticExtension() bool {
	if o != nil && !IsNil(o.SPhonestaticExtension) {
		return true
	}

	return false
}

// SetSPhonestaticExtension gets a reference to the given string and assigns it to the SPhonestaticExtension field.
func (o *PhonestaticResponseCompound) SetSPhonestaticExtension(v string) {
	o.SPhonestaticExtension = &v
}

func (o PhonestaticResponseCompound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PhonestaticResponseCompound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiPhonestaticID"] = o.PkiPhonestaticID
	if !IsNil(o.SPhonestaticE164) {
		toSerialize["sPhonestaticE164"] = o.SPhonestaticE164
	}
	if !IsNil(o.SPhonestaticExtension) {
		toSerialize["sPhonestaticExtension"] = o.SPhonestaticExtension
	}
	return toSerialize, nil
}

type NullablePhonestaticResponseCompound struct {
	value *PhonestaticResponseCompound
	isSet bool
}

func (v NullablePhonestaticResponseCompound) Get() *PhonestaticResponseCompound {
	return v.value
}

func (v *NullablePhonestaticResponseCompound) Set(val *PhonestaticResponseCompound) {
	v.value = val
	v.isSet = true
}

func (v NullablePhonestaticResponseCompound) IsSet() bool {
	return v.isSet
}

func (v *NullablePhonestaticResponseCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePhonestaticResponseCompound(val *PhonestaticResponseCompound) *NullablePhonestaticResponseCompound {
	return &NullablePhonestaticResponseCompound{value: val, isSet: true}
}

func (v NullablePhonestaticResponseCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePhonestaticResponseCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


