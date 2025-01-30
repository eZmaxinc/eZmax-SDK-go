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

// checks if the ScimServiceProviderConfigPatch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScimServiceProviderConfigPatch{}

// ScimServiceProviderConfigPatch A complex type that specifies PATCH configuration options.
type ScimServiceProviderConfigPatch struct {
	// A Boolean value specifying whether or not the operation is supported.
	Supported bool `json:"supported"`
}

type _ScimServiceProviderConfigPatch ScimServiceProviderConfigPatch

// NewScimServiceProviderConfigPatch instantiates a new ScimServiceProviderConfigPatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScimServiceProviderConfigPatch(supported bool) *ScimServiceProviderConfigPatch {
	this := ScimServiceProviderConfigPatch{}
	this.Supported = supported
	return &this
}

// NewScimServiceProviderConfigPatchWithDefaults instantiates a new ScimServiceProviderConfigPatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScimServiceProviderConfigPatchWithDefaults() *ScimServiceProviderConfigPatch {
	this := ScimServiceProviderConfigPatch{}
	return &this
}

// GetSupported returns the Supported field value
func (o *ScimServiceProviderConfigPatch) GetSupported() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Supported
}

// GetSupportedOk returns a tuple with the Supported field value
// and a boolean to check if the value has been set.
func (o *ScimServiceProviderConfigPatch) GetSupportedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Supported, true
}

// SetSupported sets field value
func (o *ScimServiceProviderConfigPatch) SetSupported(v bool) {
	o.Supported = v
}

func (o ScimServiceProviderConfigPatch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScimServiceProviderConfigPatch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["supported"] = o.Supported
	return toSerialize, nil
}

func (o *ScimServiceProviderConfigPatch) UnmarshalJSON(data []byte) (err error) {
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

	varScimServiceProviderConfigPatch := _ScimServiceProviderConfigPatch{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varScimServiceProviderConfigPatch)

	if err != nil {
		return err
	}

	*o = ScimServiceProviderConfigPatch(varScimServiceProviderConfigPatch)

	return err
}

type NullableScimServiceProviderConfigPatch struct {
	value *ScimServiceProviderConfigPatch
	isSet bool
}

func (v NullableScimServiceProviderConfigPatch) Get() *ScimServiceProviderConfigPatch {
	return v.value
}

func (v *NullableScimServiceProviderConfigPatch) Set(val *ScimServiceProviderConfigPatch) {
	v.value = val
	v.isSet = true
}

func (v NullableScimServiceProviderConfigPatch) IsSet() bool {
	return v.isSet
}

func (v *NullableScimServiceProviderConfigPatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScimServiceProviderConfigPatch(val *ScimServiceProviderConfigPatch) *NullableScimServiceProviderConfigPatch {
	return &NullableScimServiceProviderConfigPatch{value: val, isSet: true}
}

func (v NullableScimServiceProviderConfigPatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScimServiceProviderConfigPatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


