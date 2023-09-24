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

// checks if the EzsigntemplateCopyV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigntemplateCopyV1ResponseMPayload{}

// EzsigntemplateCopyV1ResponseMPayload Payload for POST /1/object/ezsigntemplate/{pkiEzsigntemplateID}/copy
type EzsigntemplateCopyV1ResponseMPayload struct {
	// An array of unique IDs representing the object that were requested to be copied.  They are returned in the same order as the array containing the objects to be created that was sent in the request.
	APkiEzsigntemplateID []int32 `json:"a_pkiEzsigntemplateID"`
}

// NewEzsigntemplateCopyV1ResponseMPayload instantiates a new EzsigntemplateCopyV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigntemplateCopyV1ResponseMPayload(aPkiEzsigntemplateID []int32) *EzsigntemplateCopyV1ResponseMPayload {
	this := EzsigntemplateCopyV1ResponseMPayload{}
	this.APkiEzsigntemplateID = aPkiEzsigntemplateID
	return &this
}

// NewEzsigntemplateCopyV1ResponseMPayloadWithDefaults instantiates a new EzsigntemplateCopyV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigntemplateCopyV1ResponseMPayloadWithDefaults() *EzsigntemplateCopyV1ResponseMPayload {
	this := EzsigntemplateCopyV1ResponseMPayload{}
	return &this
}

// GetAPkiEzsigntemplateID returns the APkiEzsigntemplateID field value
func (o *EzsigntemplateCopyV1ResponseMPayload) GetAPkiEzsigntemplateID() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.APkiEzsigntemplateID
}

// GetAPkiEzsigntemplateIDOk returns a tuple with the APkiEzsigntemplateID field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateCopyV1ResponseMPayload) GetAPkiEzsigntemplateIDOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.APkiEzsigntemplateID, true
}

// SetAPkiEzsigntemplateID sets field value
func (o *EzsigntemplateCopyV1ResponseMPayload) SetAPkiEzsigntemplateID(v []int32) {
	o.APkiEzsigntemplateID = v
}

func (o EzsigntemplateCopyV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigntemplateCopyV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_pkiEzsigntemplateID"] = o.APkiEzsigntemplateID
	return toSerialize, nil
}

type NullableEzsigntemplateCopyV1ResponseMPayload struct {
	value *EzsigntemplateCopyV1ResponseMPayload
	isSet bool
}

func (v NullableEzsigntemplateCopyV1ResponseMPayload) Get() *EzsigntemplateCopyV1ResponseMPayload {
	return v.value
}

func (v *NullableEzsigntemplateCopyV1ResponseMPayload) Set(val *EzsigntemplateCopyV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigntemplateCopyV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigntemplateCopyV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigntemplateCopyV1ResponseMPayload(val *EzsigntemplateCopyV1ResponseMPayload) *NullableEzsigntemplateCopyV1ResponseMPayload {
	return &NullableEzsigntemplateCopyV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableEzsigntemplateCopyV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigntemplateCopyV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


