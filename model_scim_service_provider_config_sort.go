/*
eZmax API Definition (Full)

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.2.2
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ScimServiceProviderConfigSort type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScimServiceProviderConfigSort{}

// ScimServiceProviderConfigSort A complex type that specifies Sort configuration options.
type ScimServiceProviderConfigSort struct {
	// A Boolean value specifying whether or not sorting is supported.
	Supported bool `json:"supported"`
}

type _ScimServiceProviderConfigSort ScimServiceProviderConfigSort

// NewScimServiceProviderConfigSort instantiates a new ScimServiceProviderConfigSort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScimServiceProviderConfigSort(supported bool) *ScimServiceProviderConfigSort {
	this := ScimServiceProviderConfigSort{}
	this.Supported = supported
	return &this
}

// NewScimServiceProviderConfigSortWithDefaults instantiates a new ScimServiceProviderConfigSort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScimServiceProviderConfigSortWithDefaults() *ScimServiceProviderConfigSort {
	this := ScimServiceProviderConfigSort{}
	return &this
}

// GetSupported returns the Supported field value
func (o *ScimServiceProviderConfigSort) GetSupported() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Supported
}

// GetSupportedOk returns a tuple with the Supported field value
// and a boolean to check if the value has been set.
func (o *ScimServiceProviderConfigSort) GetSupportedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Supported, true
}

// SetSupported sets field value
func (o *ScimServiceProviderConfigSort) SetSupported(v bool) {
	o.Supported = v
}

func (o ScimServiceProviderConfigSort) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScimServiceProviderConfigSort) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["supported"] = o.Supported
	return toSerialize, nil
}

func (o *ScimServiceProviderConfigSort) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"supported",
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

	varScimServiceProviderConfigSort := _ScimServiceProviderConfigSort{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varScimServiceProviderConfigSort)

	if err != nil {
		return err
	}

	*o = ScimServiceProviderConfigSort(varScimServiceProviderConfigSort)

	return err
}

type NullableScimServiceProviderConfigSort struct {
	value *ScimServiceProviderConfigSort
	isSet bool
}

func (v NullableScimServiceProviderConfigSort) Get() *ScimServiceProviderConfigSort {
	return v.value
}

func (v *NullableScimServiceProviderConfigSort) Set(val *ScimServiceProviderConfigSort) {
	v.value = val
	v.isSet = true
}

func (v NullableScimServiceProviderConfigSort) IsSet() bool {
	return v.isSet
}

func (v *NullableScimServiceProviderConfigSort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScimServiceProviderConfigSort(val *ScimServiceProviderConfigSort) *NullableScimServiceProviderConfigSort {
	return &NullableScimServiceProviderConfigSort{value: val, isSet: true}
}

func (v NullableScimServiceProviderConfigSort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScimServiceProviderConfigSort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


