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

// checks if the PermissionCreateObjectV1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PermissionCreateObjectV1Request{}

// PermissionCreateObjectV1Request Request for POST /1/object/permission
type PermissionCreateObjectV1Request struct {
	AObjPermission []PermissionRequestCompound `json:"a_objPermission"`
}

type _PermissionCreateObjectV1Request PermissionCreateObjectV1Request

// NewPermissionCreateObjectV1Request instantiates a new PermissionCreateObjectV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPermissionCreateObjectV1Request(aObjPermission []PermissionRequestCompound) *PermissionCreateObjectV1Request {
	this := PermissionCreateObjectV1Request{}
	this.AObjPermission = aObjPermission
	return &this
}

// NewPermissionCreateObjectV1RequestWithDefaults instantiates a new PermissionCreateObjectV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPermissionCreateObjectV1RequestWithDefaults() *PermissionCreateObjectV1Request {
	this := PermissionCreateObjectV1Request{}
	return &this
}

// GetAObjPermission returns the AObjPermission field value
func (o *PermissionCreateObjectV1Request) GetAObjPermission() []PermissionRequestCompound {
	if o == nil {
		var ret []PermissionRequestCompound
		return ret
	}

	return o.AObjPermission
}

// GetAObjPermissionOk returns a tuple with the AObjPermission field value
// and a boolean to check if the value has been set.
func (o *PermissionCreateObjectV1Request) GetAObjPermissionOk() ([]PermissionRequestCompound, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjPermission, true
}

// SetAObjPermission sets field value
func (o *PermissionCreateObjectV1Request) SetAObjPermission(v []PermissionRequestCompound) {
	o.AObjPermission = v
}

func (o PermissionCreateObjectV1Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PermissionCreateObjectV1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_objPermission"] = o.AObjPermission
	return toSerialize, nil
}

func (o *PermissionCreateObjectV1Request) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"a_objPermission",
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

	varPermissionCreateObjectV1Request := _PermissionCreateObjectV1Request{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPermissionCreateObjectV1Request)

	if err != nil {
		return err
	}

	*o = PermissionCreateObjectV1Request(varPermissionCreateObjectV1Request)

	return err
}

type NullablePermissionCreateObjectV1Request struct {
	value *PermissionCreateObjectV1Request
	isSet bool
}

func (v NullablePermissionCreateObjectV1Request) Get() *PermissionCreateObjectV1Request {
	return v.value
}

func (v *NullablePermissionCreateObjectV1Request) Set(val *PermissionCreateObjectV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullablePermissionCreateObjectV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullablePermissionCreateObjectV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermissionCreateObjectV1Request(val *PermissionCreateObjectV1Request) *NullablePermissionCreateObjectV1Request {
	return &NullablePermissionCreateObjectV1Request{value: val, isSet: true}
}

func (v NullablePermissionCreateObjectV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermissionCreateObjectV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


