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

// checks if the EzsignsignergroupCreateObjectV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignsignergroupCreateObjectV1ResponseMPayload{}

// EzsignsignergroupCreateObjectV1ResponseMPayload Payload for POST /1/object/ezsignsignergroup
type EzsignsignergroupCreateObjectV1ResponseMPayload struct {
	// An array of unique IDs representing the object that were requested to be created.  They are returned in the same order as the array containing the objects to be created that was sent in the request.
	APkiEzsignsignergroupID []int32 `json:"a_pkiEzsignsignergroupID"`
}

// NewEzsignsignergroupCreateObjectV1ResponseMPayload instantiates a new EzsignsignergroupCreateObjectV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignsignergroupCreateObjectV1ResponseMPayload(aPkiEzsignsignergroupID []int32) *EzsignsignergroupCreateObjectV1ResponseMPayload {
	this := EzsignsignergroupCreateObjectV1ResponseMPayload{}
	this.APkiEzsignsignergroupID = aPkiEzsignsignergroupID
	return &this
}

// NewEzsignsignergroupCreateObjectV1ResponseMPayloadWithDefaults instantiates a new EzsignsignergroupCreateObjectV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignsignergroupCreateObjectV1ResponseMPayloadWithDefaults() *EzsignsignergroupCreateObjectV1ResponseMPayload {
	this := EzsignsignergroupCreateObjectV1ResponseMPayload{}
	return &this
}

// GetAPkiEzsignsignergroupID returns the APkiEzsignsignergroupID field value
func (o *EzsignsignergroupCreateObjectV1ResponseMPayload) GetAPkiEzsignsignergroupID() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.APkiEzsignsignergroupID
}

// GetAPkiEzsignsignergroupIDOk returns a tuple with the APkiEzsignsignergroupID field value
// and a boolean to check if the value has been set.
func (o *EzsignsignergroupCreateObjectV1ResponseMPayload) GetAPkiEzsignsignergroupIDOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.APkiEzsignsignergroupID, true
}

// SetAPkiEzsignsignergroupID sets field value
func (o *EzsignsignergroupCreateObjectV1ResponseMPayload) SetAPkiEzsignsignergroupID(v []int32) {
	o.APkiEzsignsignergroupID = v
}

func (o EzsignsignergroupCreateObjectV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignsignergroupCreateObjectV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_pkiEzsignsignergroupID"] = o.APkiEzsignsignergroupID
	return toSerialize, nil
}

type NullableEzsignsignergroupCreateObjectV1ResponseMPayload struct {
	value *EzsignsignergroupCreateObjectV1ResponseMPayload
	isSet bool
}

func (v NullableEzsignsignergroupCreateObjectV1ResponseMPayload) Get() *EzsignsignergroupCreateObjectV1ResponseMPayload {
	return v.value
}

func (v *NullableEzsignsignergroupCreateObjectV1ResponseMPayload) Set(val *EzsignsignergroupCreateObjectV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignsignergroupCreateObjectV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignsignergroupCreateObjectV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignsignergroupCreateObjectV1ResponseMPayload(val *EzsignsignergroupCreateObjectV1ResponseMPayload) *NullableEzsignsignergroupCreateObjectV1ResponseMPayload {
	return &NullableEzsignsignergroupCreateObjectV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableEzsignsignergroupCreateObjectV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignsignergroupCreateObjectV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

