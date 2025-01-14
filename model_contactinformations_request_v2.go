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

// checks if the ContactinformationsRequestV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContactinformationsRequestV2{}

// ContactinformationsRequestV2 A Contactinformations Object
type ContactinformationsRequestV2 struct {
	EContactinformationsType FieldEContactinformationsType `json:"eContactinformationsType"`
	// The index in the a_objAddress array (zero based index) representing the Address object that should become the default one.  You can leave the value to 0 if the array is empty.
	IAddressDefault int32 `json:"iAddressDefault"`
	// The index in the a_objPhone array (zero based index) representing the Phone object that should become the default one.  You can leave the value to 0 if the array is empty.
	IPhoneDefault int32 `json:"iPhoneDefault"`
	// The index in the a_objEmail array (zero based index) representing the Email object that should become the default one.  You can leave the value to 0 if the array is empty.
	IEmailDefault int32 `json:"iEmailDefault"`
	// The index in the a_objWebsite array (zero based index) representing the Website object that should become the default one.  You can leave the value to 0 if the array is empty.
	IWebsiteDefault int32 `json:"iWebsiteDefault"`
}

type _ContactinformationsRequestV2 ContactinformationsRequestV2

// NewContactinformationsRequestV2 instantiates a new ContactinformationsRequestV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContactinformationsRequestV2(eContactinformationsType FieldEContactinformationsType, iAddressDefault int32, iPhoneDefault int32, iEmailDefault int32, iWebsiteDefault int32) *ContactinformationsRequestV2 {
	this := ContactinformationsRequestV2{}
	this.EContactinformationsType = eContactinformationsType
	this.IAddressDefault = iAddressDefault
	this.IPhoneDefault = iPhoneDefault
	this.IEmailDefault = iEmailDefault
	this.IWebsiteDefault = iWebsiteDefault
	return &this
}

// NewContactinformationsRequestV2WithDefaults instantiates a new ContactinformationsRequestV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContactinformationsRequestV2WithDefaults() *ContactinformationsRequestV2 {
	this := ContactinformationsRequestV2{}
	return &this
}

// GetEContactinformationsType returns the EContactinformationsType field value
func (o *ContactinformationsRequestV2) GetEContactinformationsType() FieldEContactinformationsType {
	if o == nil {
		var ret FieldEContactinformationsType
		return ret
	}

	return o.EContactinformationsType
}

// GetEContactinformationsTypeOk returns a tuple with the EContactinformationsType field value
// and a boolean to check if the value has been set.
func (o *ContactinformationsRequestV2) GetEContactinformationsTypeOk() (*FieldEContactinformationsType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EContactinformationsType, true
}

// SetEContactinformationsType sets field value
func (o *ContactinformationsRequestV2) SetEContactinformationsType(v FieldEContactinformationsType) {
	o.EContactinformationsType = v
}

// GetIAddressDefault returns the IAddressDefault field value
func (o *ContactinformationsRequestV2) GetIAddressDefault() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IAddressDefault
}

// GetIAddressDefaultOk returns a tuple with the IAddressDefault field value
// and a boolean to check if the value has been set.
func (o *ContactinformationsRequestV2) GetIAddressDefaultOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IAddressDefault, true
}

// SetIAddressDefault sets field value
func (o *ContactinformationsRequestV2) SetIAddressDefault(v int32) {
	o.IAddressDefault = v
}

// GetIPhoneDefault returns the IPhoneDefault field value
func (o *ContactinformationsRequestV2) GetIPhoneDefault() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IPhoneDefault
}

// GetIPhoneDefaultOk returns a tuple with the IPhoneDefault field value
// and a boolean to check if the value has been set.
func (o *ContactinformationsRequestV2) GetIPhoneDefaultOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IPhoneDefault, true
}

// SetIPhoneDefault sets field value
func (o *ContactinformationsRequestV2) SetIPhoneDefault(v int32) {
	o.IPhoneDefault = v
}

// GetIEmailDefault returns the IEmailDefault field value
func (o *ContactinformationsRequestV2) GetIEmailDefault() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEmailDefault
}

// GetIEmailDefaultOk returns a tuple with the IEmailDefault field value
// and a boolean to check if the value has been set.
func (o *ContactinformationsRequestV2) GetIEmailDefaultOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEmailDefault, true
}

// SetIEmailDefault sets field value
func (o *ContactinformationsRequestV2) SetIEmailDefault(v int32) {
	o.IEmailDefault = v
}

// GetIWebsiteDefault returns the IWebsiteDefault field value
func (o *ContactinformationsRequestV2) GetIWebsiteDefault() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IWebsiteDefault
}

// GetIWebsiteDefaultOk returns a tuple with the IWebsiteDefault field value
// and a boolean to check if the value has been set.
func (o *ContactinformationsRequestV2) GetIWebsiteDefaultOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IWebsiteDefault, true
}

// SetIWebsiteDefault sets field value
func (o *ContactinformationsRequestV2) SetIWebsiteDefault(v int32) {
	o.IWebsiteDefault = v
}

func (o ContactinformationsRequestV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContactinformationsRequestV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["eContactinformationsType"] = o.EContactinformationsType
	toSerialize["iAddressDefault"] = o.IAddressDefault
	toSerialize["iPhoneDefault"] = o.IPhoneDefault
	toSerialize["iEmailDefault"] = o.IEmailDefault
	toSerialize["iWebsiteDefault"] = o.IWebsiteDefault
	return toSerialize, nil
}

func (o *ContactinformationsRequestV2) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"eContactinformationsType",
		"iAddressDefault",
		"iPhoneDefault",
		"iEmailDefault",
		"iWebsiteDefault",
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

	varContactinformationsRequestV2 := _ContactinformationsRequestV2{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varContactinformationsRequestV2)

	if err != nil {
		return err
	}

	*o = ContactinformationsRequestV2(varContactinformationsRequestV2)

	return err
}

type NullableContactinformationsRequestV2 struct {
	value *ContactinformationsRequestV2
	isSet bool
}

func (v NullableContactinformationsRequestV2) Get() *ContactinformationsRequestV2 {
	return v.value
}

func (v *NullableContactinformationsRequestV2) Set(val *ContactinformationsRequestV2) {
	v.value = val
	v.isSet = true
}

func (v NullableContactinformationsRequestV2) IsSet() bool {
	return v.isSet
}

func (v *NullableContactinformationsRequestV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContactinformationsRequestV2(val *ContactinformationsRequestV2) *NullableContactinformationsRequestV2 {
	return &NullableContactinformationsRequestV2{value: val, isSet: true}
}

func (v NullableContactinformationsRequestV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContactinformationsRequestV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


