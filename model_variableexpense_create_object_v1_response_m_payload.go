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

// checks if the VariableexpenseCreateObjectV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VariableexpenseCreateObjectV1ResponseMPayload{}

// VariableexpenseCreateObjectV1ResponseMPayload Payload for POST /1/object/variableexpense
type VariableexpenseCreateObjectV1ResponseMPayload struct {
	// An array of unique IDs representing the object that were requested to be created.  They are returned in the same order as the array containing the objects to be created that was sent in the request.
	APkiVariableexpenseID []int32 `json:"a_pkiVariableexpenseID"`
}

// NewVariableexpenseCreateObjectV1ResponseMPayload instantiates a new VariableexpenseCreateObjectV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVariableexpenseCreateObjectV1ResponseMPayload(aPkiVariableexpenseID []int32) *VariableexpenseCreateObjectV1ResponseMPayload {
	this := VariableexpenseCreateObjectV1ResponseMPayload{}
	this.APkiVariableexpenseID = aPkiVariableexpenseID
	return &this
}

// NewVariableexpenseCreateObjectV1ResponseMPayloadWithDefaults instantiates a new VariableexpenseCreateObjectV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVariableexpenseCreateObjectV1ResponseMPayloadWithDefaults() *VariableexpenseCreateObjectV1ResponseMPayload {
	this := VariableexpenseCreateObjectV1ResponseMPayload{}
	return &this
}

// GetAPkiVariableexpenseID returns the APkiVariableexpenseID field value
func (o *VariableexpenseCreateObjectV1ResponseMPayload) GetAPkiVariableexpenseID() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.APkiVariableexpenseID
}

// GetAPkiVariableexpenseIDOk returns a tuple with the APkiVariableexpenseID field value
// and a boolean to check if the value has been set.
func (o *VariableexpenseCreateObjectV1ResponseMPayload) GetAPkiVariableexpenseIDOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.APkiVariableexpenseID, true
}

// SetAPkiVariableexpenseID sets field value
func (o *VariableexpenseCreateObjectV1ResponseMPayload) SetAPkiVariableexpenseID(v []int32) {
	o.APkiVariableexpenseID = v
}

func (o VariableexpenseCreateObjectV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VariableexpenseCreateObjectV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_pkiVariableexpenseID"] = o.APkiVariableexpenseID
	return toSerialize, nil
}

type NullableVariableexpenseCreateObjectV1ResponseMPayload struct {
	value *VariableexpenseCreateObjectV1ResponseMPayload
	isSet bool
}

func (v NullableVariableexpenseCreateObjectV1ResponseMPayload) Get() *VariableexpenseCreateObjectV1ResponseMPayload {
	return v.value
}

func (v *NullableVariableexpenseCreateObjectV1ResponseMPayload) Set(val *VariableexpenseCreateObjectV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableVariableexpenseCreateObjectV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableVariableexpenseCreateObjectV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVariableexpenseCreateObjectV1ResponseMPayload(val *VariableexpenseCreateObjectV1ResponseMPayload) *NullableVariableexpenseCreateObjectV1ResponseMPayload {
	return &NullableVariableexpenseCreateObjectV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableVariableexpenseCreateObjectV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVariableexpenseCreateObjectV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


