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

// checks if the EzsigntemplatesignerCreateObjectV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigntemplatesignerCreateObjectV1ResponseMPayload{}

// EzsigntemplatesignerCreateObjectV1ResponseMPayload Payload for POST /1/object/ezsigntemplatesigner
type EzsigntemplatesignerCreateObjectV1ResponseMPayload struct {
	// An array of unique IDs representing the object that were requested to be created.  They are returned in the same order as the array containing the objects to be created that was sent in the request.
	APkiEzsigntemplatesignerID []int32 `json:"a_pkiEzsigntemplatesignerID"`
	// Whether the Ezsignbulksend was automatically modified and needs a manual validation
	BEzsigntemplatepackageNeedvalidation bool `json:"bEzsigntemplatepackageNeedvalidation"`
	// Whether the Ezsigntemplatepackage was automatically modified and needs a manual validation
	BEzsignbulksendNeedvalidation bool `json:"bEzsignbulksendNeedvalidation"`
}

// NewEzsigntemplatesignerCreateObjectV1ResponseMPayload instantiates a new EzsigntemplatesignerCreateObjectV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigntemplatesignerCreateObjectV1ResponseMPayload(aPkiEzsigntemplatesignerID []int32, bEzsigntemplatepackageNeedvalidation bool, bEzsignbulksendNeedvalidation bool) *EzsigntemplatesignerCreateObjectV1ResponseMPayload {
	this := EzsigntemplatesignerCreateObjectV1ResponseMPayload{}
	this.APkiEzsigntemplatesignerID = aPkiEzsigntemplatesignerID
	this.BEzsigntemplatepackageNeedvalidation = bEzsigntemplatepackageNeedvalidation
	this.BEzsignbulksendNeedvalidation = bEzsignbulksendNeedvalidation
	return &this
}

// NewEzsigntemplatesignerCreateObjectV1ResponseMPayloadWithDefaults instantiates a new EzsigntemplatesignerCreateObjectV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigntemplatesignerCreateObjectV1ResponseMPayloadWithDefaults() *EzsigntemplatesignerCreateObjectV1ResponseMPayload {
	this := EzsigntemplatesignerCreateObjectV1ResponseMPayload{}
	return &this
}

// GetAPkiEzsigntemplatesignerID returns the APkiEzsigntemplatesignerID field value
func (o *EzsigntemplatesignerCreateObjectV1ResponseMPayload) GetAPkiEzsigntemplatesignerID() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.APkiEzsigntemplatesignerID
}

// GetAPkiEzsigntemplatesignerIDOk returns a tuple with the APkiEzsigntemplatesignerID field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatesignerCreateObjectV1ResponseMPayload) GetAPkiEzsigntemplatesignerIDOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.APkiEzsigntemplatesignerID, true
}

// SetAPkiEzsigntemplatesignerID sets field value
func (o *EzsigntemplatesignerCreateObjectV1ResponseMPayload) SetAPkiEzsigntemplatesignerID(v []int32) {
	o.APkiEzsigntemplatesignerID = v
}

// GetBEzsigntemplatepackageNeedvalidation returns the BEzsigntemplatepackageNeedvalidation field value
func (o *EzsigntemplatesignerCreateObjectV1ResponseMPayload) GetBEzsigntemplatepackageNeedvalidation() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BEzsigntemplatepackageNeedvalidation
}

// GetBEzsigntemplatepackageNeedvalidationOk returns a tuple with the BEzsigntemplatepackageNeedvalidation field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatesignerCreateObjectV1ResponseMPayload) GetBEzsigntemplatepackageNeedvalidationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BEzsigntemplatepackageNeedvalidation, true
}

// SetBEzsigntemplatepackageNeedvalidation sets field value
func (o *EzsigntemplatesignerCreateObjectV1ResponseMPayload) SetBEzsigntemplatepackageNeedvalidation(v bool) {
	o.BEzsigntemplatepackageNeedvalidation = v
}

// GetBEzsignbulksendNeedvalidation returns the BEzsignbulksendNeedvalidation field value
func (o *EzsigntemplatesignerCreateObjectV1ResponseMPayload) GetBEzsignbulksendNeedvalidation() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BEzsignbulksendNeedvalidation
}

// GetBEzsignbulksendNeedvalidationOk returns a tuple with the BEzsignbulksendNeedvalidation field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatesignerCreateObjectV1ResponseMPayload) GetBEzsignbulksendNeedvalidationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BEzsignbulksendNeedvalidation, true
}

// SetBEzsignbulksendNeedvalidation sets field value
func (o *EzsigntemplatesignerCreateObjectV1ResponseMPayload) SetBEzsignbulksendNeedvalidation(v bool) {
	o.BEzsignbulksendNeedvalidation = v
}

func (o EzsigntemplatesignerCreateObjectV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigntemplatesignerCreateObjectV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_pkiEzsigntemplatesignerID"] = o.APkiEzsigntemplatesignerID
	toSerialize["bEzsigntemplatepackageNeedvalidation"] = o.BEzsigntemplatepackageNeedvalidation
	toSerialize["bEzsignbulksendNeedvalidation"] = o.BEzsignbulksendNeedvalidation
	return toSerialize, nil
}

type NullableEzsigntemplatesignerCreateObjectV1ResponseMPayload struct {
	value *EzsigntemplatesignerCreateObjectV1ResponseMPayload
	isSet bool
}

func (v NullableEzsigntemplatesignerCreateObjectV1ResponseMPayload) Get() *EzsigntemplatesignerCreateObjectV1ResponseMPayload {
	return v.value
}

func (v *NullableEzsigntemplatesignerCreateObjectV1ResponseMPayload) Set(val *EzsigntemplatesignerCreateObjectV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigntemplatesignerCreateObjectV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigntemplatesignerCreateObjectV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigntemplatesignerCreateObjectV1ResponseMPayload(val *EzsigntemplatesignerCreateObjectV1ResponseMPayload) *NullableEzsigntemplatesignerCreateObjectV1ResponseMPayload {
	return &NullableEzsigntemplatesignerCreateObjectV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableEzsigntemplatesignerCreateObjectV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigntemplatesignerCreateObjectV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


