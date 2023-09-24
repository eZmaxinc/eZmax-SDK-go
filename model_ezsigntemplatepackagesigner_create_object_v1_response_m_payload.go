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

// checks if the EzsigntemplatepackagesignerCreateObjectV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigntemplatepackagesignerCreateObjectV1ResponseMPayload{}

// EzsigntemplatepackagesignerCreateObjectV1ResponseMPayload Payload for POST /1/object/ezsigntemplatepackagesigner
type EzsigntemplatepackagesignerCreateObjectV1ResponseMPayload struct {
	// An array of unique IDs representing the object that were requested to be created.  They are returned in the same order as the array containing the objects to be created that was sent in the request.
	APkiEzsigntemplatepackagesignerID []int32 `json:"a_pkiEzsigntemplatepackagesignerID"`
}

// NewEzsigntemplatepackagesignerCreateObjectV1ResponseMPayload instantiates a new EzsigntemplatepackagesignerCreateObjectV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigntemplatepackagesignerCreateObjectV1ResponseMPayload(aPkiEzsigntemplatepackagesignerID []int32) *EzsigntemplatepackagesignerCreateObjectV1ResponseMPayload {
	this := EzsigntemplatepackagesignerCreateObjectV1ResponseMPayload{}
	this.APkiEzsigntemplatepackagesignerID = aPkiEzsigntemplatepackagesignerID
	return &this
}

// NewEzsigntemplatepackagesignerCreateObjectV1ResponseMPayloadWithDefaults instantiates a new EzsigntemplatepackagesignerCreateObjectV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigntemplatepackagesignerCreateObjectV1ResponseMPayloadWithDefaults() *EzsigntemplatepackagesignerCreateObjectV1ResponseMPayload {
	this := EzsigntemplatepackagesignerCreateObjectV1ResponseMPayload{}
	return &this
}

// GetAPkiEzsigntemplatepackagesignerID returns the APkiEzsigntemplatepackagesignerID field value
func (o *EzsigntemplatepackagesignerCreateObjectV1ResponseMPayload) GetAPkiEzsigntemplatepackagesignerID() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.APkiEzsigntemplatepackagesignerID
}

// GetAPkiEzsigntemplatepackagesignerIDOk returns a tuple with the APkiEzsigntemplatepackagesignerID field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepackagesignerCreateObjectV1ResponseMPayload) GetAPkiEzsigntemplatepackagesignerIDOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.APkiEzsigntemplatepackagesignerID, true
}

// SetAPkiEzsigntemplatepackagesignerID sets field value
func (o *EzsigntemplatepackagesignerCreateObjectV1ResponseMPayload) SetAPkiEzsigntemplatepackagesignerID(v []int32) {
	o.APkiEzsigntemplatepackagesignerID = v
}

func (o EzsigntemplatepackagesignerCreateObjectV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigntemplatepackagesignerCreateObjectV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_pkiEzsigntemplatepackagesignerID"] = o.APkiEzsigntemplatepackagesignerID
	return toSerialize, nil
}

type NullableEzsigntemplatepackagesignerCreateObjectV1ResponseMPayload struct {
	value *EzsigntemplatepackagesignerCreateObjectV1ResponseMPayload
	isSet bool
}

func (v NullableEzsigntemplatepackagesignerCreateObjectV1ResponseMPayload) Get() *EzsigntemplatepackagesignerCreateObjectV1ResponseMPayload {
	return v.value
}

func (v *NullableEzsigntemplatepackagesignerCreateObjectV1ResponseMPayload) Set(val *EzsigntemplatepackagesignerCreateObjectV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigntemplatepackagesignerCreateObjectV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigntemplatepackagesignerCreateObjectV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigntemplatepackagesignerCreateObjectV1ResponseMPayload(val *EzsigntemplatepackagesignerCreateObjectV1ResponseMPayload) *NullableEzsigntemplatepackagesignerCreateObjectV1ResponseMPayload {
	return &NullableEzsigntemplatepackagesignerCreateObjectV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableEzsigntemplatepackagesignerCreateObjectV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigntemplatepackagesignerCreateObjectV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


