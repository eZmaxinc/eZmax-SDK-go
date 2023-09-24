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

// checks if the UsergroupGetUsergroupmembershipsV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UsergroupGetUsergroupmembershipsV1ResponseMPayload{}

// UsergroupGetUsergroupmembershipsV1ResponseMPayload Response for GET /1/object/usergroup/{pkiUsergroupID}/getUsergroupmemberships
type UsergroupGetUsergroupmembershipsV1ResponseMPayload struct {
	AObjUsergroupmembership []UsergroupmembershipResponseCompound `json:"a_objUsergroupmembership"`
}

// NewUsergroupGetUsergroupmembershipsV1ResponseMPayload instantiates a new UsergroupGetUsergroupmembershipsV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsergroupGetUsergroupmembershipsV1ResponseMPayload(aObjUsergroupmembership []UsergroupmembershipResponseCompound) *UsergroupGetUsergroupmembershipsV1ResponseMPayload {
	this := UsergroupGetUsergroupmembershipsV1ResponseMPayload{}
	this.AObjUsergroupmembership = aObjUsergroupmembership
	return &this
}

// NewUsergroupGetUsergroupmembershipsV1ResponseMPayloadWithDefaults instantiates a new UsergroupGetUsergroupmembershipsV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsergroupGetUsergroupmembershipsV1ResponseMPayloadWithDefaults() *UsergroupGetUsergroupmembershipsV1ResponseMPayload {
	this := UsergroupGetUsergroupmembershipsV1ResponseMPayload{}
	return &this
}

// GetAObjUsergroupmembership returns the AObjUsergroupmembership field value
func (o *UsergroupGetUsergroupmembershipsV1ResponseMPayload) GetAObjUsergroupmembership() []UsergroupmembershipResponseCompound {
	if o == nil {
		var ret []UsergroupmembershipResponseCompound
		return ret
	}

	return o.AObjUsergroupmembership
}

// GetAObjUsergroupmembershipOk returns a tuple with the AObjUsergroupmembership field value
// and a boolean to check if the value has been set.
func (o *UsergroupGetUsergroupmembershipsV1ResponseMPayload) GetAObjUsergroupmembershipOk() ([]UsergroupmembershipResponseCompound, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjUsergroupmembership, true
}

// SetAObjUsergroupmembership sets field value
func (o *UsergroupGetUsergroupmembershipsV1ResponseMPayload) SetAObjUsergroupmembership(v []UsergroupmembershipResponseCompound) {
	o.AObjUsergroupmembership = v
}

func (o UsergroupGetUsergroupmembershipsV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsergroupGetUsergroupmembershipsV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_objUsergroupmembership"] = o.AObjUsergroupmembership
	return toSerialize, nil
}

type NullableUsergroupGetUsergroupmembershipsV1ResponseMPayload struct {
	value *UsergroupGetUsergroupmembershipsV1ResponseMPayload
	isSet bool
}

func (v NullableUsergroupGetUsergroupmembershipsV1ResponseMPayload) Get() *UsergroupGetUsergroupmembershipsV1ResponseMPayload {
	return v.value
}

func (v *NullableUsergroupGetUsergroupmembershipsV1ResponseMPayload) Set(val *UsergroupGetUsergroupmembershipsV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableUsergroupGetUsergroupmembershipsV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableUsergroupGetUsergroupmembershipsV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsergroupGetUsergroupmembershipsV1ResponseMPayload(val *UsergroupGetUsergroupmembershipsV1ResponseMPayload) *NullableUsergroupGetUsergroupmembershipsV1ResponseMPayload {
	return &NullableUsergroupGetUsergroupmembershipsV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableUsergroupGetUsergroupmembershipsV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsergroupGetUsergroupmembershipsV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


