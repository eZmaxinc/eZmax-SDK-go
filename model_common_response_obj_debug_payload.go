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
	"bytes"
	"fmt"
)

// checks if the CommonResponseObjDebugPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommonResponseObjDebugPayload{}

// CommonResponseObjDebugPayload This is a debug object containing debugging information on the actual function
type CommonResponseObjDebugPayload struct {
	// The minimum version of the function that can be called
	IVersionMin int32 `json:"iVersionMin"`
	// The maximum version of the function that can be called
	IVersionMax int32 `json:"iVersionMax"`
	// An array of permissions required to access this function.  If the value \"0\" is present in the array, anyone can call this function.  You must have one of the permission to access the function. You don't need to have all of them.
	ARequiredPermission []int32 `json:"a_RequiredPermission"`
	// Wheter the current route is deprecated or not
	BVersionDeprecated bool `json:"bVersionDeprecated"`
}

type _CommonResponseObjDebugPayload CommonResponseObjDebugPayload

// NewCommonResponseObjDebugPayload instantiates a new CommonResponseObjDebugPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonResponseObjDebugPayload(iVersionMin int32, iVersionMax int32, aRequiredPermission []int32, bVersionDeprecated bool) *CommonResponseObjDebugPayload {
	this := CommonResponseObjDebugPayload{}
	this.IVersionMin = iVersionMin
	this.IVersionMax = iVersionMax
	this.ARequiredPermission = aRequiredPermission
	this.BVersionDeprecated = bVersionDeprecated
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
	if o == nil {
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
	if o == nil {
		return nil, false
	}
	return &o.IVersionMax, true
}

// SetIVersionMax sets field value
func (o *CommonResponseObjDebugPayload) SetIVersionMax(v int32) {
	o.IVersionMax = v
}

// GetARequiredPermission returns the ARequiredPermission field value
func (o *CommonResponseObjDebugPayload) GetARequiredPermission() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.ARequiredPermission
}

// GetARequiredPermissionOk returns a tuple with the ARequiredPermission field value
// and a boolean to check if the value has been set.
func (o *CommonResponseObjDebugPayload) GetARequiredPermissionOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ARequiredPermission, true
}

// SetARequiredPermission sets field value
func (o *CommonResponseObjDebugPayload) SetARequiredPermission(v []int32) {
	o.ARequiredPermission = v
}

// GetBVersionDeprecated returns the BVersionDeprecated field value
func (o *CommonResponseObjDebugPayload) GetBVersionDeprecated() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BVersionDeprecated
}

// GetBVersionDeprecatedOk returns a tuple with the BVersionDeprecated field value
// and a boolean to check if the value has been set.
func (o *CommonResponseObjDebugPayload) GetBVersionDeprecatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BVersionDeprecated, true
}

// SetBVersionDeprecated sets field value
func (o *CommonResponseObjDebugPayload) SetBVersionDeprecated(v bool) {
	o.BVersionDeprecated = v
}

func (o CommonResponseObjDebugPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommonResponseObjDebugPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["iVersionMin"] = o.IVersionMin
	toSerialize["iVersionMax"] = o.IVersionMax
	toSerialize["a_RequiredPermission"] = o.ARequiredPermission
	toSerialize["bVersionDeprecated"] = o.BVersionDeprecated
	return toSerialize, nil
}

func (o *CommonResponseObjDebugPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"iVersionMin",
		"iVersionMax",
		"a_RequiredPermission",
		"bVersionDeprecated",
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

	varCommonResponseObjDebugPayload := _CommonResponseObjDebugPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCommonResponseObjDebugPayload)

	if err != nil {
		return err
	}

	*o = CommonResponseObjDebugPayload(varCommonResponseObjDebugPayload)

	return err
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


