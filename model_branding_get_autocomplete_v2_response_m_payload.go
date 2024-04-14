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

// checks if the BrandingGetAutocompleteV2ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BrandingGetAutocompleteV2ResponseMPayload{}

// BrandingGetAutocompleteV2ResponseMPayload Payload for POST /2/object/branding/getAutocomplete
type BrandingGetAutocompleteV2ResponseMPayload struct {
	// An array of Branding object containing the description, ID and active status about the element.
	AObjBranding []BrandingAutocompleteElementResponse `json:"a_objBranding"`
}

type _BrandingGetAutocompleteV2ResponseMPayload BrandingGetAutocompleteV2ResponseMPayload

// NewBrandingGetAutocompleteV2ResponseMPayload instantiates a new BrandingGetAutocompleteV2ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBrandingGetAutocompleteV2ResponseMPayload(aObjBranding []BrandingAutocompleteElementResponse) *BrandingGetAutocompleteV2ResponseMPayload {
	this := BrandingGetAutocompleteV2ResponseMPayload{}
	this.AObjBranding = aObjBranding
	return &this
}

// NewBrandingGetAutocompleteV2ResponseMPayloadWithDefaults instantiates a new BrandingGetAutocompleteV2ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBrandingGetAutocompleteV2ResponseMPayloadWithDefaults() *BrandingGetAutocompleteV2ResponseMPayload {
	this := BrandingGetAutocompleteV2ResponseMPayload{}
	return &this
}

// GetAObjBranding returns the AObjBranding field value
func (o *BrandingGetAutocompleteV2ResponseMPayload) GetAObjBranding() []BrandingAutocompleteElementResponse {
	if o == nil {
		var ret []BrandingAutocompleteElementResponse
		return ret
	}

	return o.AObjBranding
}

// GetAObjBrandingOk returns a tuple with the AObjBranding field value
// and a boolean to check if the value has been set.
func (o *BrandingGetAutocompleteV2ResponseMPayload) GetAObjBrandingOk() ([]BrandingAutocompleteElementResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjBranding, true
}

// SetAObjBranding sets field value
func (o *BrandingGetAutocompleteV2ResponseMPayload) SetAObjBranding(v []BrandingAutocompleteElementResponse) {
	o.AObjBranding = v
}

func (o BrandingGetAutocompleteV2ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BrandingGetAutocompleteV2ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_objBranding"] = o.AObjBranding
	return toSerialize, nil
}

func (o *BrandingGetAutocompleteV2ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
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

	varBrandingGetAutocompleteV2ResponseMPayload := _BrandingGetAutocompleteV2ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBrandingGetAutocompleteV2ResponseMPayload)

	if err != nil {
		return err
	}

	*o = BrandingGetAutocompleteV2ResponseMPayload(varBrandingGetAutocompleteV2ResponseMPayload)

	return err
}

type NullableBrandingGetAutocompleteV2ResponseMPayload struct {
	value *BrandingGetAutocompleteV2ResponseMPayload
	isSet bool
}

func (v NullableBrandingGetAutocompleteV2ResponseMPayload) Get() *BrandingGetAutocompleteV2ResponseMPayload {
	return v.value
}

func (v *NullableBrandingGetAutocompleteV2ResponseMPayload) Set(val *BrandingGetAutocompleteV2ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableBrandingGetAutocompleteV2ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableBrandingGetAutocompleteV2ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBrandingGetAutocompleteV2ResponseMPayload(val *BrandingGetAutocompleteV2ResponseMPayload) *NullableBrandingGetAutocompleteV2ResponseMPayload {
	return &NullableBrandingGetAutocompleteV2ResponseMPayload{value: val, isSet: true}
}

func (v NullableBrandingGetAutocompleteV2ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBrandingGetAutocompleteV2ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


