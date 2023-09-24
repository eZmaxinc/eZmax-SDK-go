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

// checks if the ScimServiceProviderConfigEtag type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScimServiceProviderConfigEtag{}

// ScimServiceProviderConfigEtag A complex type that specifies ETag configuration options.
type ScimServiceProviderConfigEtag struct {
	// A Boolean value specifying whether or not the operation is supported.
	Supported bool `json:"supported"`
}

// NewScimServiceProviderConfigEtag instantiates a new ScimServiceProviderConfigEtag object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScimServiceProviderConfigEtag(supported bool) *ScimServiceProviderConfigEtag {
	this := ScimServiceProviderConfigEtag{}
	this.Supported = supported
	return &this
}

// NewScimServiceProviderConfigEtagWithDefaults instantiates a new ScimServiceProviderConfigEtag object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScimServiceProviderConfigEtagWithDefaults() *ScimServiceProviderConfigEtag {
	this := ScimServiceProviderConfigEtag{}
	return &this
}

// GetSupported returns the Supported field value
func (o *ScimServiceProviderConfigEtag) GetSupported() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Supported
}

// GetSupportedOk returns a tuple with the Supported field value
// and a boolean to check if the value has been set.
func (o *ScimServiceProviderConfigEtag) GetSupportedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Supported, true
}

// SetSupported sets field value
func (o *ScimServiceProviderConfigEtag) SetSupported(v bool) {
	o.Supported = v
}

func (o ScimServiceProviderConfigEtag) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScimServiceProviderConfigEtag) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["supported"] = o.Supported
	return toSerialize, nil
}

type NullableScimServiceProviderConfigEtag struct {
	value *ScimServiceProviderConfigEtag
	isSet bool
}

func (v NullableScimServiceProviderConfigEtag) Get() *ScimServiceProviderConfigEtag {
	return v.value
}

func (v *NullableScimServiceProviderConfigEtag) Set(val *ScimServiceProviderConfigEtag) {
	v.value = val
	v.isSet = true
}

func (v NullableScimServiceProviderConfigEtag) IsSet() bool {
	return v.isSet
}

func (v *NullableScimServiceProviderConfigEtag) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScimServiceProviderConfigEtag(val *ScimServiceProviderConfigEtag) *NullableScimServiceProviderConfigEtag {
	return &NullableScimServiceProviderConfigEtag{value: val, isSet: true}
}

func (v NullableScimServiceProviderConfigEtag) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScimServiceProviderConfigEtag) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

