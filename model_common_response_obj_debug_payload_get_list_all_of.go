/*
eZmax API Definition

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.1.4
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
)

// CommonResponseObjDebugPayloadGetListAllOf struct for CommonResponseObjDebugPayloadGetListAllOf
type CommonResponseObjDebugPayloadGetListAllOf struct {
	AFilter CommonResponseFilter `json:"a_Filter"`
	// List of available values for *eOrderBy*
	AOrderBy map[string]string `json:"a_OrderBy"`
}

// NewCommonResponseObjDebugPayloadGetListAllOf instantiates a new CommonResponseObjDebugPayloadGetListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonResponseObjDebugPayloadGetListAllOf(aFilter CommonResponseFilter, aOrderBy map[string]string) *CommonResponseObjDebugPayloadGetListAllOf {
	this := CommonResponseObjDebugPayloadGetListAllOf{}
	this.AFilter = aFilter
	this.AOrderBy = aOrderBy
	return &this
}

// NewCommonResponseObjDebugPayloadGetListAllOfWithDefaults instantiates a new CommonResponseObjDebugPayloadGetListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonResponseObjDebugPayloadGetListAllOfWithDefaults() *CommonResponseObjDebugPayloadGetListAllOf {
	this := CommonResponseObjDebugPayloadGetListAllOf{}
	return &this
}

// GetAFilter returns the AFilter field value
func (o *CommonResponseObjDebugPayloadGetListAllOf) GetAFilter() CommonResponseFilter {
	if o == nil {
		var ret CommonResponseFilter
		return ret
	}

	return o.AFilter
}

// GetAFilterOk returns a tuple with the AFilter field value
// and a boolean to check if the value has been set.
func (o *CommonResponseObjDebugPayloadGetListAllOf) GetAFilterOk() (*CommonResponseFilter, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AFilter, true
}

// SetAFilter sets field value
func (o *CommonResponseObjDebugPayloadGetListAllOf) SetAFilter(v CommonResponseFilter) {
	o.AFilter = v
}

// GetAOrderBy returns the AOrderBy field value
func (o *CommonResponseObjDebugPayloadGetListAllOf) GetAOrderBy() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.AOrderBy
}

// GetAOrderByOk returns a tuple with the AOrderBy field value
// and a boolean to check if the value has been set.
func (o *CommonResponseObjDebugPayloadGetListAllOf) GetAOrderByOk() (*map[string]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AOrderBy, true
}

// SetAOrderBy sets field value
func (o *CommonResponseObjDebugPayloadGetListAllOf) SetAOrderBy(v map[string]string) {
	o.AOrderBy = v
}

func (o CommonResponseObjDebugPayloadGetListAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["a_Filter"] = o.AFilter
	}
	if true {
		toSerialize["a_OrderBy"] = o.AOrderBy
	}
	return json.Marshal(toSerialize)
}

type NullableCommonResponseObjDebugPayloadGetListAllOf struct {
	value *CommonResponseObjDebugPayloadGetListAllOf
	isSet bool
}

func (v NullableCommonResponseObjDebugPayloadGetListAllOf) Get() *CommonResponseObjDebugPayloadGetListAllOf {
	return v.value
}

func (v *NullableCommonResponseObjDebugPayloadGetListAllOf) Set(val *CommonResponseObjDebugPayloadGetListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonResponseObjDebugPayloadGetListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonResponseObjDebugPayloadGetListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonResponseObjDebugPayloadGetListAllOf(val *CommonResponseObjDebugPayloadGetListAllOf) *NullableCommonResponseObjDebugPayloadGetListAllOf {
	return &NullableCommonResponseObjDebugPayloadGetListAllOf{value: val, isSet: true}
}

func (v NullableCommonResponseObjDebugPayloadGetListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonResponseObjDebugPayloadGetListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


