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

// checks if the UsergroupEditPermissionsV1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UsergroupEditPermissionsV1Request{}

// UsergroupEditPermissionsV1Request Request for PUT /1/object/usergroup/{pkiUsergroupID}/editPermissions
type UsergroupEditPermissionsV1Request struct {
	AObjPermission []PermissionRequestCompound `json:"a_objPermission"`
}

type _UsergroupEditPermissionsV1Request UsergroupEditPermissionsV1Request

// NewUsergroupEditPermissionsV1Request instantiates a new UsergroupEditPermissionsV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsergroupEditPermissionsV1Request(aObjPermission []PermissionRequestCompound) *UsergroupEditPermissionsV1Request {
	this := UsergroupEditPermissionsV1Request{}
	this.AObjPermission = aObjPermission
	return &this
}

// NewUsergroupEditPermissionsV1RequestWithDefaults instantiates a new UsergroupEditPermissionsV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsergroupEditPermissionsV1RequestWithDefaults() *UsergroupEditPermissionsV1Request {
	this := UsergroupEditPermissionsV1Request{}
	return &this
}

// GetAObjPermission returns the AObjPermission field value
func (o *UsergroupEditPermissionsV1Request) GetAObjPermission() []PermissionRequestCompound {
	if o == nil {
		var ret []PermissionRequestCompound
		return ret
	}

	return o.AObjPermission
}

// GetAObjPermissionOk returns a tuple with the AObjPermission field value
// and a boolean to check if the value has been set.
func (o *UsergroupEditPermissionsV1Request) GetAObjPermissionOk() ([]PermissionRequestCompound, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjPermission, true
}

// SetAObjPermission sets field value
func (o *UsergroupEditPermissionsV1Request) SetAObjPermission(v []PermissionRequestCompound) {
	o.AObjPermission = v
}

func (o UsergroupEditPermissionsV1Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsergroupEditPermissionsV1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_objPermission"] = o.AObjPermission
	return toSerialize, nil
}

func (o *UsergroupEditPermissionsV1Request) UnmarshalJSON(data []byte) (err error) {
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

	varUsergroupEditPermissionsV1Request := _UsergroupEditPermissionsV1Request{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUsergroupEditPermissionsV1Request)

	if err != nil {
		return err
	}

	*o = UsergroupEditPermissionsV1Request(varUsergroupEditPermissionsV1Request)

	return err
}

type NullableUsergroupEditPermissionsV1Request struct {
	value *UsergroupEditPermissionsV1Request
	isSet bool
}

func (v NullableUsergroupEditPermissionsV1Request) Get() *UsergroupEditPermissionsV1Request {
	return v.value
}

func (v *NullableUsergroupEditPermissionsV1Request) Set(val *UsergroupEditPermissionsV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableUsergroupEditPermissionsV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableUsergroupEditPermissionsV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsergroupEditPermissionsV1Request(val *UsergroupEditPermissionsV1Request) *NullableUsergroupEditPermissionsV1Request {
	return &NullableUsergroupEditPermissionsV1Request{value: val, isSet: true}
}

func (v NullableUsergroupEditPermissionsV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsergroupEditPermissionsV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


