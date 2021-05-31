/*
 * eZmax API Definition
 *
 * This API expose all the functionnalities for the eZmax and eZsign applications.
 *
 * API version: 1.0.44
 * Contact: support-api@ezmax.ca
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
)

// CommonResponseObjDebugPayload This is a debug object containing debugging information on the actual function
type CommonResponseObjDebugPayload struct {
	// The minimum version of the function that can be called
	IVersionMin int32 `json:"iVersionMin"`
	// The maximum version of the function that can be called
	IVersionMax int32 `json:"iVersionMax"`
	// An array of permissions required to access this function.  If the value \"0\" is present in the array, anyone can call this function.  You must have one of the permission to access the function. You don't need to have all of them.
	ARequiredPermissions []int32 `json:"a_RequiredPermissions"`
}

// NewCommonResponseObjDebugPayload instantiates a new CommonResponseObjDebugPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonResponseObjDebugPayload(iVersionMin int32, iVersionMax int32, aRequiredPermissions []int32) *CommonResponseObjDebugPayload {
	this := CommonResponseObjDebugPayload{}
	this.IVersionMin = iVersionMin
	this.IVersionMax = iVersionMax
	this.ARequiredPermissions = aRequiredPermissions
	return &this
}

// NewCommonResponseObjDebugPayloadWithDefaults instantiates a new CommonResponseObjDebugPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonResponseObjDebugPayloadWithDefaults() *CommonResponseObjDebugPayload {
	this := CommonResponseObjDebugPayload{}
	return &this
}

// GetIVersionMin returns the IVersionMin field value
func (o *CommonResponseObjDebugPayload) GetIVersionMin() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IVersionMin
}

// GetIVersionMinOk returns a tuple with the IVersionMin field value
// and a boolean to check if the value has been set.
func (o *CommonResponseObjDebugPayload) GetIVersionMinOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IVersionMin, true
}

// SetIVersionMin sets field value
func (o *CommonResponseObjDebugPayload) SetIVersionMin(v int32) {
	o.IVersionMin = v
}

// GetIVersionMax returns the IVersionMax field value
func (o *CommonResponseObjDebugPayload) GetIVersionMax() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IVersionMax
}

// GetIVersionMaxOk returns a tuple with the IVersionMax field value
// and a boolean to check if the value has been set.
func (o *CommonResponseObjDebugPayload) GetIVersionMaxOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IVersionMax, true
}

// SetIVersionMax sets field value
func (o *CommonResponseObjDebugPayload) SetIVersionMax(v int32) {
	o.IVersionMax = v
}

// GetARequiredPermissions returns the ARequiredPermissions field value
func (o *CommonResponseObjDebugPayload) GetARequiredPermissions() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.ARequiredPermissions
}

// GetARequiredPermissionsOk returns a tuple with the ARequiredPermissions field value
// and a boolean to check if the value has been set.
func (o *CommonResponseObjDebugPayload) GetARequiredPermissionsOk() (*[]int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ARequiredPermissions, true
}

// SetARequiredPermissions sets field value
func (o *CommonResponseObjDebugPayload) SetARequiredPermissions(v []int32) {
	o.ARequiredPermissions = v
}

func (o CommonResponseObjDebugPayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["iVersionMin"] = o.IVersionMin
	}
	if true {
		toSerialize["iVersionMax"] = o.IVersionMax
	}
	if true {
		toSerialize["a_RequiredPermissions"] = o.ARequiredPermissions
	}
	return json.Marshal(toSerialize)
}

type NullableCommonResponseObjDebugPayload struct {
	value *CommonResponseObjDebugPayload
	isSet bool
}

func (v NullableCommonResponseObjDebugPayload) Get() *CommonResponseObjDebugPayload {
	return v.value
}

func (v *NullableCommonResponseObjDebugPayload) Set(val *CommonResponseObjDebugPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonResponseObjDebugPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonResponseObjDebugPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonResponseObjDebugPayload(val *CommonResponseObjDebugPayload) *NullableCommonResponseObjDebugPayload {
	return &NullableCommonResponseObjDebugPayload{value: val, isSet: true}
}

func (v NullableCommonResponseObjDebugPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonResponseObjDebugPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


