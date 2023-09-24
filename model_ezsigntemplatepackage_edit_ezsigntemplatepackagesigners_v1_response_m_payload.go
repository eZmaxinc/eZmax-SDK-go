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

// checks if the EzsigntemplatepackageEditEzsigntemplatepackagesignersV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigntemplatepackageEditEzsigntemplatepackagesignersV1ResponseMPayload{}

// EzsigntemplatepackageEditEzsigntemplatepackagesignersV1ResponseMPayload Payload for PUT /1/object/ezsigntemplatepackage/{pkiEzsigntemplatepackageID}/editEzsigntemplatepackagesigners
type EzsigntemplatepackageEditEzsigntemplatepackagesignersV1ResponseMPayload struct {
	APkiEzsigntemplatepackagesignerID []int32 `json:"a_pkiEzsigntemplatepackagesignerID"`
}

// NewEzsigntemplatepackageEditEzsigntemplatepackagesignersV1ResponseMPayload instantiates a new EzsigntemplatepackageEditEzsigntemplatepackagesignersV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigntemplatepackageEditEzsigntemplatepackagesignersV1ResponseMPayload(aPkiEzsigntemplatepackagesignerID []int32) *EzsigntemplatepackageEditEzsigntemplatepackagesignersV1ResponseMPayload {
	this := EzsigntemplatepackageEditEzsigntemplatepackagesignersV1ResponseMPayload{}
	this.APkiEzsigntemplatepackagesignerID = aPkiEzsigntemplatepackagesignerID
	return &this
}

// NewEzsigntemplatepackageEditEzsigntemplatepackagesignersV1ResponseMPayloadWithDefaults instantiates a new EzsigntemplatepackageEditEzsigntemplatepackagesignersV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigntemplatepackageEditEzsigntemplatepackagesignersV1ResponseMPayloadWithDefaults() *EzsigntemplatepackageEditEzsigntemplatepackagesignersV1ResponseMPayload {
	this := EzsigntemplatepackageEditEzsigntemplatepackagesignersV1ResponseMPayload{}
	return &this
}

// GetAPkiEzsigntemplatepackagesignerID returns the APkiEzsigntemplatepackagesignerID field value
func (o *EzsigntemplatepackageEditEzsigntemplatepackagesignersV1ResponseMPayload) GetAPkiEzsigntemplatepackagesignerID() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.APkiEzsigntemplatepackagesignerID
}

// GetAPkiEzsigntemplatepackagesignerIDOk returns a tuple with the APkiEzsigntemplatepackagesignerID field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepackageEditEzsigntemplatepackagesignersV1ResponseMPayload) GetAPkiEzsigntemplatepackagesignerIDOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.APkiEzsigntemplatepackagesignerID, true
}

// SetAPkiEzsigntemplatepackagesignerID sets field value
func (o *EzsigntemplatepackageEditEzsigntemplatepackagesignersV1ResponseMPayload) SetAPkiEzsigntemplatepackagesignerID(v []int32) {
	o.APkiEzsigntemplatepackagesignerID = v
}

func (o EzsigntemplatepackageEditEzsigntemplatepackagesignersV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigntemplatepackageEditEzsigntemplatepackagesignersV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_pkiEzsigntemplatepackagesignerID"] = o.APkiEzsigntemplatepackagesignerID
	return toSerialize, nil
}

type NullableEzsigntemplatepackageEditEzsigntemplatepackagesignersV1ResponseMPayload struct {
	value *EzsigntemplatepackageEditEzsigntemplatepackagesignersV1ResponseMPayload
	isSet bool
}

func (v NullableEzsigntemplatepackageEditEzsigntemplatepackagesignersV1ResponseMPayload) Get() *EzsigntemplatepackageEditEzsigntemplatepackagesignersV1ResponseMPayload {
	return v.value
}

func (v *NullableEzsigntemplatepackageEditEzsigntemplatepackagesignersV1ResponseMPayload) Set(val *EzsigntemplatepackageEditEzsigntemplatepackagesignersV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigntemplatepackageEditEzsigntemplatepackagesignersV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigntemplatepackageEditEzsigntemplatepackagesignersV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigntemplatepackageEditEzsigntemplatepackagesignersV1ResponseMPayload(val *EzsigntemplatepackageEditEzsigntemplatepackagesignersV1ResponseMPayload) *NullableEzsigntemplatepackageEditEzsigntemplatepackagesignersV1ResponseMPayload {
	return &NullableEzsigntemplatepackageEditEzsigntemplatepackagesignersV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableEzsigntemplatepackageEditEzsigntemplatepackagesignersV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigntemplatepackageEditEzsigntemplatepackagesignersV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


