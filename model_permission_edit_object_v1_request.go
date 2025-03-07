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

// checks if the PermissionEditObjectV1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PermissionEditObjectV1Request{}

// PermissionEditObjectV1Request Request for PUT /1/object/permission/{pkiPermissionID}
type PermissionEditObjectV1Request struct {
	ObjPermission PermissionRequestCompound `json:"objPermission"`
}

type _PermissionEditObjectV1Request PermissionEditObjectV1Request

// NewPermissionEditObjectV1Request instantiates a new PermissionEditObjectV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPermissionEditObjectV1Request(objPermission PermissionRequestCompound) *PermissionEditObjectV1Request {
	this := PermissionEditObjectV1Request{}
	this.ObjPermission = objPermission
	return &this
}

// NewPermissionEditObjectV1RequestWithDefaults instantiates a new PermissionEditObjectV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPermissionEditObjectV1RequestWithDefaults() *PermissionEditObjectV1Request {
	this := PermissionEditObjectV1Request{}
	return &this
}

// GetObjPermission returns the ObjPermission field value
func (o *PermissionEditObjectV1Request) GetObjPermission() PermissionRequestCompound {
	if o == nil {
		var ret PermissionRequestCompound
		return ret
	}

	return o.ObjPermission
}

// GetObjPermissionOk returns a tuple with the ObjPermission field value
// and a boolean to check if the value has been set.
func (o *PermissionEditObjectV1Request) GetObjPermissionOk() (*PermissionRequestCompound, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjPermission, true
}

// SetObjPermission sets field value
func (o *PermissionEditObjectV1Request) SetObjPermission(v PermissionRequestCompound) {
	o.ObjPermission = v
}

func (o PermissionEditObjectV1Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PermissionEditObjectV1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objPermission"] = o.ObjPermission
	return toSerialize, nil
}

func (o *PermissionEditObjectV1Request) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"objPermission",
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

	varPermissionEditObjectV1Request := _PermissionEditObjectV1Request{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPermissionEditObjectV1Request)

	if err != nil {
		return err
	}

	*o = PermissionEditObjectV1Request(varPermissionEditObjectV1Request)

	return err
}

type NullablePermissionEditObjectV1Request struct {
	value *PermissionEditObjectV1Request
	isSet bool
}

func (v NullablePermissionEditObjectV1Request) Get() *PermissionEditObjectV1Request {
	return v.value
}

func (v *NullablePermissionEditObjectV1Request) Set(val *PermissionEditObjectV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullablePermissionEditObjectV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullablePermissionEditObjectV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermissionEditObjectV1Request(val *PermissionEditObjectV1Request) *NullablePermissionEditObjectV1Request {
	return &NullablePermissionEditObjectV1Request{value: val, isSet: true}
}

func (v NullablePermissionEditObjectV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermissionEditObjectV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


