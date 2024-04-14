/*
eZmax API Definition (Full)

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.2.1
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the PhoneResponseCompound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PhoneResponseCompound{}

// PhoneResponseCompound A Phone Object and children to create a complete structure
type PhoneResponseCompound struct {
	// The unique ID of the Phone.
	PkiPhoneID int32 `json:"pkiPhoneID"`
	// The unique ID of the Phonetype.  Valid values:  |Value|Description| |-|-| |1|Office| |2|Home| |3|Mobile| |4|Fax| |5|Pager| |6|Toll Free|
	FkiPhonetypeID int32 `json:"fkiPhonetypeID"`
	// Deprecated
	EPhoneType *FieldEPhoneType `json:"ePhoneType,omitempty"`
	// A phone number in E.164 Format
	SPhoneE164 *string `json:"sPhoneE164,omitempty"`
	// The extension of the phone number.  The extension is the \"123\" section in this sample phone number: (514) 990-1516 x123.  It can also be used with international phone numbers
	SPhoneExtension *string `json:"sPhoneExtension,omitempty"`
	// Indicate the phone number is an international phone number.
	BPhoneInternational *bool `json:"bPhoneInternational,omitempty"`
}

type _PhoneResponseCompound PhoneResponseCompound

// NewPhoneResponseCompound instantiates a new PhoneResponseCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPhoneResponseCompound(pkiPhoneID int32, fkiPhonetypeID int32) *PhoneResponseCompound {
	this := PhoneResponseCompound{}
	this.PkiPhoneID = pkiPhoneID
	this.FkiPhonetypeID = fkiPhonetypeID
	return &this
}

// NewPhoneResponseCompoundWithDefaults instantiates a new PhoneResponseCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPhoneResponseCompoundWithDefaults() *PhoneResponseCompound {
	this := PhoneResponseCompound{}
	return &this
}

// GetPkiPhoneID returns the PkiPhoneID field value
func (o *PhoneResponseCompound) GetPkiPhoneID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiPhoneID
}

// GetPkiPhoneIDOk returns a tuple with the PkiPhoneID field value
// and a boolean to check if the value has been set.
func (o *PhoneResponseCompound) GetPkiPhoneIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiPhoneID, true
}

// SetPkiPhoneID sets field value
func (o *PhoneResponseCompound) SetPkiPhoneID(v int32) {
	o.PkiPhoneID = v
}

// GetFkiPhonetypeID returns the FkiPhonetypeID field value
func (o *PhoneResponseCompound) GetFkiPhonetypeID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiPhonetypeID
}

// GetFkiPhonetypeIDOk returns a tuple with the FkiPhonetypeID field value
// and a boolean to check if the value has been set.
func (o *PhoneResponseCompound) GetFkiPhonetypeIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiPhonetypeID, true
}

// SetFkiPhonetypeID sets field value
func (o *PhoneResponseCompound) SetFkiPhonetypeID(v int32) {
	o.FkiPhonetypeID = v
}

// GetEPhoneType returns the EPhoneType field value if set, zero value otherwise.
// Deprecated
func (o *PhoneResponseCompound) GetEPhoneType() FieldEPhoneType {
	if o == nil || IsNil(o.EPhoneType) {
		var ret FieldEPhoneType
		return ret
	}
	return *o.EPhoneType
}

// GetEPhoneTypeOk returns a tuple with the EPhoneType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *PhoneResponseCompound) GetEPhoneTypeOk() (*FieldEPhoneType, bool) {
	if o == nil || IsNil(o.EPhoneType) {
		return nil, false
	}
	return o.EPhoneType, true
}

// HasEPhoneType returns a boolean if a field has been set.
func (o *PhoneResponseCompound) HasEPhoneType() bool {
	if o != nil && !IsNil(o.EPhoneType) {
		return true
	}

	return false
}

// SetEPhoneType gets a reference to the given FieldEPhoneType and assigns it to the EPhoneType field.
// Deprecated
func (o *PhoneResponseCompound) SetEPhoneType(v FieldEPhoneType) {
	o.EPhoneType = &v
}

// GetSPhoneE164 returns the SPhoneE164 field value if set, zero value otherwise.
func (o *PhoneResponseCompound) GetSPhoneE164() string {
	if o == nil || IsNil(o.SPhoneE164) {
		var ret string
		return ret
	}
	return *o.SPhoneE164
}

// GetSPhoneE164Ok returns a tuple with the SPhoneE164 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhoneResponseCompound) GetSPhoneE164Ok() (*string, bool) {
	if o == nil || IsNil(o.SPhoneE164) {
		return nil, false
	}
	return o.SPhoneE164, true
}

// HasSPhoneE164 returns a boolean if a field has been set.
func (o *PhoneResponseCompound) HasSPhoneE164() bool {
	if o != nil && !IsNil(o.SPhoneE164) {
		return true
	}

	return false
}

// SetSPhoneE164 gets a reference to the given string and assigns it to the SPhoneE164 field.
func (o *PhoneResponseCompound) SetSPhoneE164(v string) {
	o.SPhoneE164 = &v
}

// GetSPhoneExtension returns the SPhoneExtension field value if set, zero value otherwise.
func (o *PhoneResponseCompound) GetSPhoneExtension() string {
	if o == nil || IsNil(o.SPhoneExtension) {
		var ret string
		return ret
	}
	return *o.SPhoneExtension
}

// GetSPhoneExtensionOk returns a tuple with the SPhoneExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhoneResponseCompound) GetSPhoneExtensionOk() (*string, bool) {
	if o == nil || IsNil(o.SPhoneExtension) {
		return nil, false
	}
	return o.SPhoneExtension, true
}

// HasSPhoneExtension returns a boolean if a field has been set.
func (o *PhoneResponseCompound) HasSPhoneExtension() bool {
	if o != nil && !IsNil(o.SPhoneExtension) {
		return true
	}

	return false
}

// SetSPhoneExtension gets a reference to the given string and assigns it to the SPhoneExtension field.
func (o *PhoneResponseCompound) SetSPhoneExtension(v string) {
	o.SPhoneExtension = &v
}

// GetBPhoneInternational returns the BPhoneInternational field value if set, zero value otherwise.
func (o *PhoneResponseCompound) GetBPhoneInternational() bool {
	if o == nil || IsNil(o.BPhoneInternational) {
		var ret bool
		return ret
	}
	return *o.BPhoneInternational
}

// GetBPhoneInternationalOk returns a tuple with the BPhoneInternational field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhoneResponseCompound) GetBPhoneInternationalOk() (*bool, bool) {
	if o == nil || IsNil(o.BPhoneInternational) {
		return nil, false
	}
	return o.BPhoneInternational, true
}

// HasBPhoneInternational returns a boolean if a field has been set.
func (o *PhoneResponseCompound) HasBPhoneInternational() bool {
	if o != nil && !IsNil(o.BPhoneInternational) {
		return true
	}

	return false
}

// SetBPhoneInternational gets a reference to the given bool and assigns it to the BPhoneInternational field.
func (o *PhoneResponseCompound) SetBPhoneInternational(v bool) {
	o.BPhoneInternational = &v
}

func (o PhoneResponseCompound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PhoneResponseCompound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiPhoneID"] = o.PkiPhoneID
	toSerialize["fkiPhonetypeID"] = o.FkiPhonetypeID
	if !IsNil(o.EPhoneType) {
		toSerialize["ePhoneType"] = o.EPhoneType
	}
	if !IsNil(o.SPhoneE164) {
		toSerialize["sPhoneE164"] = o.SPhoneE164
	}
	if !IsNil(o.SPhoneExtension) {
		toSerialize["sPhoneExtension"] = o.SPhoneExtension
	}
	if !IsNil(o.BPhoneInternational) {
		toSerialize["bPhoneInternational"] = o.BPhoneInternational
	}
	return toSerialize, nil
}

func (o *PhoneResponseCompound) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiPhoneID",
		"fkiPhonetypeID",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varPhoneResponseCompound := _PhoneResponseCompound{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPhoneResponseCompound)

	if err != nil {
		return err
	}

	*o = PhoneResponseCompound(varPhoneResponseCompound)

	return err
}

type NullablePhoneResponseCompound struct {
	value *PhoneResponseCompound
	isSet bool
}

func (v NullablePhoneResponseCompound) Get() *PhoneResponseCompound {
	return v.value
}

func (v *NullablePhoneResponseCompound) Set(val *PhoneResponseCompound) {
	v.value = val
	v.isSet = true
}

func (v NullablePhoneResponseCompound) IsSet() bool {
	return v.isSet
}

func (v *NullablePhoneResponseCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePhoneResponseCompound(val *PhoneResponseCompound) *NullablePhoneResponseCompound {
	return &NullablePhoneResponseCompound{value: val, isSet: true}
}

func (v NullablePhoneResponseCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePhoneResponseCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


