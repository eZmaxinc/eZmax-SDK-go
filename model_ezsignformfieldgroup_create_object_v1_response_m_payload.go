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

// checks if the EzsignformfieldgroupCreateObjectV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignformfieldgroupCreateObjectV1ResponseMPayload{}

// EzsignformfieldgroupCreateObjectV1ResponseMPayload Payload for POST /1/object/ezsignformfieldgroup
type EzsignformfieldgroupCreateObjectV1ResponseMPayload struct {
	// An array of unique IDs representing the object that were requested to be created.  They are returned in the same order as the array containing the objects to be created that was sent in the request.
	APkiEzsignformfieldgroupID []int32 `json:"a_pkiEzsignformfieldgroupID"`
}

// NewEzsignformfieldgroupCreateObjectV1ResponseMPayload instantiates a new EzsignformfieldgroupCreateObjectV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignformfieldgroupCreateObjectV1ResponseMPayload(aPkiEzsignformfieldgroupID []int32) *EzsignformfieldgroupCreateObjectV1ResponseMPayload {
	this := EzsignformfieldgroupCreateObjectV1ResponseMPayload{}
	this.APkiEzsignformfieldgroupID = aPkiEzsignformfieldgroupID
	return &this
}

// NewEzsignformfieldgroupCreateObjectV1ResponseMPayloadWithDefaults instantiates a new EzsignformfieldgroupCreateObjectV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignformfieldgroupCreateObjectV1ResponseMPayloadWithDefaults() *EzsignformfieldgroupCreateObjectV1ResponseMPayload {
	this := EzsignformfieldgroupCreateObjectV1ResponseMPayload{}
	return &this
}

// GetAPkiEzsignformfieldgroupID returns the APkiEzsignformfieldgroupID field value
func (o *EzsignformfieldgroupCreateObjectV1ResponseMPayload) GetAPkiEzsignformfieldgroupID() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.APkiEzsignformfieldgroupID
}

// GetAPkiEzsignformfieldgroupIDOk returns a tuple with the APkiEzsignformfieldgroupID field value
// and a boolean to check if the value has been set.
func (o *EzsignformfieldgroupCreateObjectV1ResponseMPayload) GetAPkiEzsignformfieldgroupIDOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.APkiEzsignformfieldgroupID, true
}

// SetAPkiEzsignformfieldgroupID sets field value
func (o *EzsignformfieldgroupCreateObjectV1ResponseMPayload) SetAPkiEzsignformfieldgroupID(v []int32) {
	o.APkiEzsignformfieldgroupID = v
}

func (o EzsignformfieldgroupCreateObjectV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignformfieldgroupCreateObjectV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_pkiEzsignformfieldgroupID"] = o.APkiEzsignformfieldgroupID
	return toSerialize, nil
}

type NullableEzsignformfieldgroupCreateObjectV1ResponseMPayload struct {
	value *EzsignformfieldgroupCreateObjectV1ResponseMPayload
	isSet bool
}

func (v NullableEzsignformfieldgroupCreateObjectV1ResponseMPayload) Get() *EzsignformfieldgroupCreateObjectV1ResponseMPayload {
	return v.value
}

func (v *NullableEzsignformfieldgroupCreateObjectV1ResponseMPayload) Set(val *EzsignformfieldgroupCreateObjectV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignformfieldgroupCreateObjectV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignformfieldgroupCreateObjectV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignformfieldgroupCreateObjectV1ResponseMPayload(val *EzsignformfieldgroupCreateObjectV1ResponseMPayload) *NullableEzsignformfieldgroupCreateObjectV1ResponseMPayload {
	return &NullableEzsignformfieldgroupCreateObjectV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableEzsignformfieldgroupCreateObjectV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignformfieldgroupCreateObjectV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


