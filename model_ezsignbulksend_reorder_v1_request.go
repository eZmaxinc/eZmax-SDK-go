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

// checks if the EzsignbulksendReorderV1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignbulksendReorderV1Request{}

// EzsignbulksendReorderV1Request Request for POST /1/object/ezsignbulksend/{pkiEzsignbulksendID}/reorder
type EzsignbulksendReorderV1Request struct {
	APkiEzsignbulksenddocumentmappingID []int32 `json:"a_pkiEzsignbulksenddocumentmappingID"`
}

// NewEzsignbulksendReorderV1Request instantiates a new EzsignbulksendReorderV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignbulksendReorderV1Request(aPkiEzsignbulksenddocumentmappingID []int32) *EzsignbulksendReorderV1Request {
	this := EzsignbulksendReorderV1Request{}
	this.APkiEzsignbulksenddocumentmappingID = aPkiEzsignbulksenddocumentmappingID
	return &this
}

// NewEzsignbulksendReorderV1RequestWithDefaults instantiates a new EzsignbulksendReorderV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignbulksendReorderV1RequestWithDefaults() *EzsignbulksendReorderV1Request {
	this := EzsignbulksendReorderV1Request{}
	return &this
}

// GetAPkiEzsignbulksenddocumentmappingID returns the APkiEzsignbulksenddocumentmappingID field value
func (o *EzsignbulksendReorderV1Request) GetAPkiEzsignbulksenddocumentmappingID() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.APkiEzsignbulksenddocumentmappingID
}

// GetAPkiEzsignbulksenddocumentmappingIDOk returns a tuple with the APkiEzsignbulksenddocumentmappingID field value
// and a boolean to check if the value has been set.
func (o *EzsignbulksendReorderV1Request) GetAPkiEzsignbulksenddocumentmappingIDOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.APkiEzsignbulksenddocumentmappingID, true
}

// SetAPkiEzsignbulksenddocumentmappingID sets field value
func (o *EzsignbulksendReorderV1Request) SetAPkiEzsignbulksenddocumentmappingID(v []int32) {
	o.APkiEzsignbulksenddocumentmappingID = v
}

func (o EzsignbulksendReorderV1Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignbulksendReorderV1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_pkiEzsignbulksenddocumentmappingID"] = o.APkiEzsignbulksenddocumentmappingID
	return toSerialize, nil
}

type NullableEzsignbulksendReorderV1Request struct {
	value *EzsignbulksendReorderV1Request
	isSet bool
}

func (v NullableEzsignbulksendReorderV1Request) Get() *EzsignbulksendReorderV1Request {
	return v.value
}

func (v *NullableEzsignbulksendReorderV1Request) Set(val *EzsignbulksendReorderV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignbulksendReorderV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignbulksendReorderV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignbulksendReorderV1Request(val *EzsignbulksendReorderV1Request) *NullableEzsignbulksendReorderV1Request {
	return &NullableEzsignbulksendReorderV1Request{value: val, isSet: true}
}

func (v NullableEzsignbulksendReorderV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignbulksendReorderV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

