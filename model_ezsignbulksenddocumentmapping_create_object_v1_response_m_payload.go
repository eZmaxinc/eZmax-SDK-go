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

// checks if the EzsignbulksenddocumentmappingCreateObjectV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignbulksenddocumentmappingCreateObjectV1ResponseMPayload{}

// EzsignbulksenddocumentmappingCreateObjectV1ResponseMPayload Payload for POST /1/object/ezsignbulksenddocumentmapping
type EzsignbulksenddocumentmappingCreateObjectV1ResponseMPayload struct {
	// An array of unique IDs representing the object that were requested to be created.  They are returned in the same order as the array containing the objects to be created that was sent in the request.
	APkiEzsignbulksenddocumentmappingID []int32 `json:"a_pkiEzsignbulksenddocumentmappingID"`
}

// NewEzsignbulksenddocumentmappingCreateObjectV1ResponseMPayload instantiates a new EzsignbulksenddocumentmappingCreateObjectV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignbulksenddocumentmappingCreateObjectV1ResponseMPayload(aPkiEzsignbulksenddocumentmappingID []int32) *EzsignbulksenddocumentmappingCreateObjectV1ResponseMPayload {
	this := EzsignbulksenddocumentmappingCreateObjectV1ResponseMPayload{}
	this.APkiEzsignbulksenddocumentmappingID = aPkiEzsignbulksenddocumentmappingID
	return &this
}

// NewEzsignbulksenddocumentmappingCreateObjectV1ResponseMPayloadWithDefaults instantiates a new EzsignbulksenddocumentmappingCreateObjectV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignbulksenddocumentmappingCreateObjectV1ResponseMPayloadWithDefaults() *EzsignbulksenddocumentmappingCreateObjectV1ResponseMPayload {
	this := EzsignbulksenddocumentmappingCreateObjectV1ResponseMPayload{}
	return &this
}

// GetAPkiEzsignbulksenddocumentmappingID returns the APkiEzsignbulksenddocumentmappingID field value
func (o *EzsignbulksenddocumentmappingCreateObjectV1ResponseMPayload) GetAPkiEzsignbulksenddocumentmappingID() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.APkiEzsignbulksenddocumentmappingID
}

// GetAPkiEzsignbulksenddocumentmappingIDOk returns a tuple with the APkiEzsignbulksenddocumentmappingID field value
// and a boolean to check if the value has been set.
func (o *EzsignbulksenddocumentmappingCreateObjectV1ResponseMPayload) GetAPkiEzsignbulksenddocumentmappingIDOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.APkiEzsignbulksenddocumentmappingID, true
}

// SetAPkiEzsignbulksenddocumentmappingID sets field value
func (o *EzsignbulksenddocumentmappingCreateObjectV1ResponseMPayload) SetAPkiEzsignbulksenddocumentmappingID(v []int32) {
	o.APkiEzsignbulksenddocumentmappingID = v
}

func (o EzsignbulksenddocumentmappingCreateObjectV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignbulksenddocumentmappingCreateObjectV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_pkiEzsignbulksenddocumentmappingID"] = o.APkiEzsignbulksenddocumentmappingID
	return toSerialize, nil
}

type NullableEzsignbulksenddocumentmappingCreateObjectV1ResponseMPayload struct {
	value *EzsignbulksenddocumentmappingCreateObjectV1ResponseMPayload
	isSet bool
}

func (v NullableEzsignbulksenddocumentmappingCreateObjectV1ResponseMPayload) Get() *EzsignbulksenddocumentmappingCreateObjectV1ResponseMPayload {
	return v.value
}

func (v *NullableEzsignbulksenddocumentmappingCreateObjectV1ResponseMPayload) Set(val *EzsignbulksenddocumentmappingCreateObjectV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignbulksenddocumentmappingCreateObjectV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignbulksenddocumentmappingCreateObjectV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignbulksenddocumentmappingCreateObjectV1ResponseMPayload(val *EzsignbulksenddocumentmappingCreateObjectV1ResponseMPayload) *NullableEzsignbulksenddocumentmappingCreateObjectV1ResponseMPayload {
	return &NullableEzsignbulksenddocumentmappingCreateObjectV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableEzsignbulksenddocumentmappingCreateObjectV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignbulksenddocumentmappingCreateObjectV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

