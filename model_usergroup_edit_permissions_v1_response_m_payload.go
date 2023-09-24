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

// checks if the UsergroupEditPermissionsV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UsergroupEditPermissionsV1ResponseMPayload{}

// UsergroupEditPermissionsV1ResponseMPayload Payload for PUT /1/object/usergroup/{pkiUsergroupID}/editPermissions
type UsergroupEditPermissionsV1ResponseMPayload struct {
	APkiPermissionID []int32 `json:"a_pkiPermissionID"`
}

// NewUsergroupEditPermissionsV1ResponseMPayload instantiates a new UsergroupEditPermissionsV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsergroupEditPermissionsV1ResponseMPayload(aPkiPermissionID []int32) *UsergroupEditPermissionsV1ResponseMPayload {
	this := UsergroupEditPermissionsV1ResponseMPayload{}
	this.APkiPermissionID = aPkiPermissionID
	return &this
}

// NewUsergroupEditPermissionsV1ResponseMPayloadWithDefaults instantiates a new UsergroupEditPermissionsV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsergroupEditPermissionsV1ResponseMPayloadWithDefaults() *UsergroupEditPermissionsV1ResponseMPayload {
	this := UsergroupEditPermissionsV1ResponseMPayload{}
	return &this
}

// GetAPkiPermissionID returns the APkiPermissionID field value
func (o *UsergroupEditPermissionsV1ResponseMPayload) GetAPkiPermissionID() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.APkiPermissionID
}

// GetAPkiPermissionIDOk returns a tuple with the APkiPermissionID field value
// and a boolean to check if the value has been set.
func (o *UsergroupEditPermissionsV1ResponseMPayload) GetAPkiPermissionIDOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.APkiPermissionID, true
}

// SetAPkiPermissionID sets field value
func (o *UsergroupEditPermissionsV1ResponseMPayload) SetAPkiPermissionID(v []int32) {
	o.APkiPermissionID = v
}

func (o UsergroupEditPermissionsV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsergroupEditPermissionsV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_pkiPermissionID"] = o.APkiPermissionID
	return toSerialize, nil
}

type NullableUsergroupEditPermissionsV1ResponseMPayload struct {
	value *UsergroupEditPermissionsV1ResponseMPayload
	isSet bool
}

func (v NullableUsergroupEditPermissionsV1ResponseMPayload) Get() *UsergroupEditPermissionsV1ResponseMPayload {
	return v.value
}

func (v *NullableUsergroupEditPermissionsV1ResponseMPayload) Set(val *UsergroupEditPermissionsV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableUsergroupEditPermissionsV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableUsergroupEditPermissionsV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsergroupEditPermissionsV1ResponseMPayload(val *UsergroupEditPermissionsV1ResponseMPayload) *NullableUsergroupEditPermissionsV1ResponseMPayload {
	return &NullableUsergroupEditPermissionsV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableUsergroupEditPermissionsV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsergroupEditPermissionsV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


