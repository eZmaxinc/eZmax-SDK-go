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

// checks if the EzsigntemplatepackagesignermembershipCreateObjectV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigntemplatepackagesignermembershipCreateObjectV1ResponseMPayload{}

// EzsigntemplatepackagesignermembershipCreateObjectV1ResponseMPayload Payload for POST /1/object/ezsigntemplatepackagesignermembership
type EzsigntemplatepackagesignermembershipCreateObjectV1ResponseMPayload struct {
	// An array of unique IDs representing the object that were requested to be created.  They are returned in the same order as the array containing the objects to be created that was sent in the request.
	APkiEzsigntemplatepackagesignermembershipID []int32 `json:"a_pkiEzsigntemplatepackagesignermembershipID"`
}

// NewEzsigntemplatepackagesignermembershipCreateObjectV1ResponseMPayload instantiates a new EzsigntemplatepackagesignermembershipCreateObjectV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigntemplatepackagesignermembershipCreateObjectV1ResponseMPayload(aPkiEzsigntemplatepackagesignermembershipID []int32) *EzsigntemplatepackagesignermembershipCreateObjectV1ResponseMPayload {
	this := EzsigntemplatepackagesignermembershipCreateObjectV1ResponseMPayload{}
	this.APkiEzsigntemplatepackagesignermembershipID = aPkiEzsigntemplatepackagesignermembershipID
	return &this
}

// NewEzsigntemplatepackagesignermembershipCreateObjectV1ResponseMPayloadWithDefaults instantiates a new EzsigntemplatepackagesignermembershipCreateObjectV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigntemplatepackagesignermembershipCreateObjectV1ResponseMPayloadWithDefaults() *EzsigntemplatepackagesignermembershipCreateObjectV1ResponseMPayload {
	this := EzsigntemplatepackagesignermembershipCreateObjectV1ResponseMPayload{}
	return &this
}

// GetAPkiEzsigntemplatepackagesignermembershipID returns the APkiEzsigntemplatepackagesignermembershipID field value
func (o *EzsigntemplatepackagesignermembershipCreateObjectV1ResponseMPayload) GetAPkiEzsigntemplatepackagesignermembershipID() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.APkiEzsigntemplatepackagesignermembershipID
}

// GetAPkiEzsigntemplatepackagesignermembershipIDOk returns a tuple with the APkiEzsigntemplatepackagesignermembershipID field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepackagesignermembershipCreateObjectV1ResponseMPayload) GetAPkiEzsigntemplatepackagesignermembershipIDOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.APkiEzsigntemplatepackagesignermembershipID, true
}

// SetAPkiEzsigntemplatepackagesignermembershipID sets field value
func (o *EzsigntemplatepackagesignermembershipCreateObjectV1ResponseMPayload) SetAPkiEzsigntemplatepackagesignermembershipID(v []int32) {
	o.APkiEzsigntemplatepackagesignermembershipID = v
}

func (o EzsigntemplatepackagesignermembershipCreateObjectV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigntemplatepackagesignermembershipCreateObjectV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_pkiEzsigntemplatepackagesignermembershipID"] = o.APkiEzsigntemplatepackagesignermembershipID
	return toSerialize, nil
}

type NullableEzsigntemplatepackagesignermembershipCreateObjectV1ResponseMPayload struct {
	value *EzsigntemplatepackagesignermembershipCreateObjectV1ResponseMPayload
	isSet bool
}

func (v NullableEzsigntemplatepackagesignermembershipCreateObjectV1ResponseMPayload) Get() *EzsigntemplatepackagesignermembershipCreateObjectV1ResponseMPayload {
	return v.value
}

func (v *NullableEzsigntemplatepackagesignermembershipCreateObjectV1ResponseMPayload) Set(val *EzsigntemplatepackagesignermembershipCreateObjectV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigntemplatepackagesignermembershipCreateObjectV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigntemplatepackagesignermembershipCreateObjectV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigntemplatepackagesignermembershipCreateObjectV1ResponseMPayload(val *EzsigntemplatepackagesignermembershipCreateObjectV1ResponseMPayload) *NullableEzsigntemplatepackagesignermembershipCreateObjectV1ResponseMPayload {
	return &NullableEzsigntemplatepackagesignermembershipCreateObjectV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableEzsigntemplatepackagesignermembershipCreateObjectV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigntemplatepackagesignermembershipCreateObjectV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


