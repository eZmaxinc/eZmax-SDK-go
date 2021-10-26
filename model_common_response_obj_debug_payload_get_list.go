/*
eZmax API Definition

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.1.2
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
)

// CommonResponseObjDebugPayloadGetList This is a debug object containing debugging information on the actual function
type CommonResponseObjDebugPayloadGetList struct {
	// The minimum version of the function that can be called
	IVersionMin int32 `json:"iVersionMin"`
	// The maximum version of the function that can be called
	IVersionMax int32 `json:"iVersionMax"`
	// An array of permissions required to access this function.  If the value \"0\" is present in the array, anyone can call this function.  You must have one of the permission to access the function. You don't need to have all of them.
	ARequiredPermission []int32 `json:"a_RequiredPermission"`
	AFilter CommonResponseFilter `json:"a_Filter"`
	// List of available values for *eOrderBy*
	AOrderBy map[string]string `json:"a_OrderBy"`
}

// NewCommonResponseObjDebugPayloadGetList instantiates a new CommonResponseObjDebugPayloadGetList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonResponseObjDebugPayloadGetList(iVersionMin int32, iVersionMax int32, aRequiredPermission []int32, aFilter CommonResponseFilter, aOrderBy map[string]string) *CommonResponseObjDebugPayloadGetList {
	this := CommonResponseObjDebugPayloadGetList{}
	this.IVersionMin = iVersionMin
	this.IVersionMax = iVersionMax
	this.ARequiredPermission = aRequiredPermission
	this.AFilter = aFilter
	this.AOrderBy = aOrderBy
	return &this
}

// NewCommonResponseObjDebugPayloadGetListWithDefaults instantiates a new CommonResponseObjDebugPayloadGetList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonResponseObjDebugPayloadGetListWithDefaults() *CommonResponseObjDebugPayloadGetList {
	this := CommonResponseObjDebugPayloadGetList{}
	return &this
}

// GetIVersionMin returns the IVersionMin field value
func (o *CommonResponseObjDebugPayloadGetList) GetIVersionMin() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IVersionMin
}

// GetIVersionMinOk returns a tuple with the IVersionMin field value
// and a boolean to check if the value has been set.
func (o *CommonResponseObjDebugPayloadGetList) GetIVersionMinOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IVersionMin, true
}

// SetIVersionMin sets field value
func (o *CommonResponseObjDebugPayloadGetList) SetIVersionMin(v int32) {
	o.IVersionMin = v
}

// GetIVersionMax returns the IVersionMax field value
func (o *CommonResponseObjDebugPayloadGetList) GetIVersionMax() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IVersionMax
}

// GetIVersionMaxOk returns a tuple with the IVersionMax field value
// and a boolean to check if the value has been set.
func (o *CommonResponseObjDebugPayloadGetList) GetIVersionMaxOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IVersionMax, true
}

// SetIVersionMax sets field value
func (o *CommonResponseObjDebugPayloadGetList) SetIVersionMax(v int32) {
	o.IVersionMax = v
}

// GetARequiredPermission returns the ARequiredPermission field value
func (o *CommonResponseObjDebugPayloadGetList) GetARequiredPermission() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.ARequiredPermission
}

// GetARequiredPermissionOk returns a tuple with the ARequiredPermission field value
// and a boolean to check if the value has been set.
func (o *CommonResponseObjDebugPayloadGetList) GetARequiredPermissionOk() (*[]int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ARequiredPermission, true
}

// SetARequiredPermission sets field value
func (o *CommonResponseObjDebugPayloadGetList) SetARequiredPermission(v []int32) {
	o.ARequiredPermission = v
}

// GetAFilter returns the AFilter field value
func (o *CommonResponseObjDebugPayloadGetList) GetAFilter() CommonResponseFilter {
	if o == nil {
		var ret CommonResponseFilter
		return ret
	}

	return o.AFilter
}

// GetAFilterOk returns a tuple with the AFilter field value
// and a boolean to check if the value has been set.
func (o *CommonResponseObjDebugPayloadGetList) GetAFilterOk() (*CommonResponseFilter, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AFilter, true
}

// SetAFilter sets field value
func (o *CommonResponseObjDebugPayloadGetList) SetAFilter(v CommonResponseFilter) {
	o.AFilter = v
}

// GetAOrderBy returns the AOrderBy field value
func (o *CommonResponseObjDebugPayloadGetList) GetAOrderBy() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.AOrderBy
}

// GetAOrderByOk returns a tuple with the AOrderBy field value
// and a boolean to check if the value has been set.
func (o *CommonResponseObjDebugPayloadGetList) GetAOrderByOk() (*map[string]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AOrderBy, true
}

// SetAOrderBy sets field value
func (o *CommonResponseObjDebugPayloadGetList) SetAOrderBy(v map[string]string) {
	o.AOrderBy = v
}

func (o CommonResponseObjDebugPayloadGetList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["iVersionMin"] = o.IVersionMin
	}
	if true {
		toSerialize["iVersionMax"] = o.IVersionMax
	}
	if true {
		toSerialize["a_RequiredPermission"] = o.ARequiredPermission
	}
	if true {
		toSerialize["a_Filter"] = o.AFilter
	}
	if true {
		toSerialize["a_OrderBy"] = o.AOrderBy
	}
	return json.Marshal(toSerialize)
}

type NullableCommonResponseObjDebugPayloadGetList struct {
	value *CommonResponseObjDebugPayloadGetList
	isSet bool
}

func (v NullableCommonResponseObjDebugPayloadGetList) Get() *CommonResponseObjDebugPayloadGetList {
	return v.value
}

func (v *NullableCommonResponseObjDebugPayloadGetList) Set(val *CommonResponseObjDebugPayloadGetList) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonResponseObjDebugPayloadGetList) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonResponseObjDebugPayloadGetList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonResponseObjDebugPayloadGetList(val *CommonResponseObjDebugPayloadGetList) *NullableCommonResponseObjDebugPayloadGetList {
	return &NullableCommonResponseObjDebugPayloadGetList{value: val, isSet: true}
}

func (v NullableCommonResponseObjDebugPayloadGetList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonResponseObjDebugPayloadGetList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


