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

// checks if the BrandingCreateObjectV2Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BrandingCreateObjectV2Request{}

// BrandingCreateObjectV2Request Request for POST /2/object/branding
type BrandingCreateObjectV2Request struct {
	AObjBranding []BrandingRequestCompoundV2 `json:"a_objBranding"`
}

type _BrandingCreateObjectV2Request BrandingCreateObjectV2Request

// NewBrandingCreateObjectV2Request instantiates a new BrandingCreateObjectV2Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBrandingCreateObjectV2Request(aObjBranding []BrandingRequestCompoundV2) *BrandingCreateObjectV2Request {
	this := BrandingCreateObjectV2Request{}
	this.AObjBranding = aObjBranding
	return &this
}

// NewBrandingCreateObjectV2RequestWithDefaults instantiates a new BrandingCreateObjectV2Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBrandingCreateObjectV2RequestWithDefaults() *BrandingCreateObjectV2Request {
	this := BrandingCreateObjectV2Request{}
	return &this
}

// GetAObjBranding returns the AObjBranding field value
func (o *BrandingCreateObjectV2Request) GetAObjBranding() []BrandingRequestCompoundV2 {
	if o == nil {
		var ret []BrandingRequestCompoundV2
		return ret
	}

	return o.AObjBranding
}

// GetAObjBrandingOk returns a tuple with the AObjBranding field value
// and a boolean to check if the value has been set.
func (o *BrandingCreateObjectV2Request) GetAObjBrandingOk() ([]BrandingRequestCompoundV2, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjBranding, true
}

// SetAObjBranding sets field value
func (o *BrandingCreateObjectV2Request) SetAObjBranding(v []BrandingRequestCompoundV2) {
	o.AObjBranding = v
}

func (o BrandingCreateObjectV2Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BrandingCreateObjectV2Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_objBranding"] = o.AObjBranding
	return toSerialize, nil
}

func (o *BrandingCreateObjectV2Request) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"a_objBranding",
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

	varBrandingCreateObjectV2Request := _BrandingCreateObjectV2Request{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBrandingCreateObjectV2Request)

	if err != nil {
		return err
	}

	*o = BrandingCreateObjectV2Request(varBrandingCreateObjectV2Request)

	return err
}

type NullableBrandingCreateObjectV2Request struct {
	value *BrandingCreateObjectV2Request
	isSet bool
}

func (v NullableBrandingCreateObjectV2Request) Get() *BrandingCreateObjectV2Request {
	return v.value
}

func (v *NullableBrandingCreateObjectV2Request) Set(val *BrandingCreateObjectV2Request) {
	v.value = val
	v.isSet = true
}

func (v NullableBrandingCreateObjectV2Request) IsSet() bool {
	return v.isSet
}

func (v *NullableBrandingCreateObjectV2Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBrandingCreateObjectV2Request(val *BrandingCreateObjectV2Request) *NullableBrandingCreateObjectV2Request {
	return &NullableBrandingCreateObjectV2Request{value: val, isSet: true}
}

func (v NullableBrandingCreateObjectV2Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBrandingCreateObjectV2Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


