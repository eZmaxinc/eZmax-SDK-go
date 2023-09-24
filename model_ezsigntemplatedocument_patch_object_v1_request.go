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

// checks if the EzsigntemplatedocumentPatchObjectV1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigntemplatedocumentPatchObjectV1Request{}

// EzsigntemplatedocumentPatchObjectV1Request Request for PATCH /1/object/ezsigntemplatedocument/{pkiEzsigntemplatedocumentID}
type EzsigntemplatedocumentPatchObjectV1Request struct {
	ObjEzsigntemplatedocument EzsigntemplatedocumentRequestPatch `json:"objEzsigntemplatedocument"`
}

// NewEzsigntemplatedocumentPatchObjectV1Request instantiates a new EzsigntemplatedocumentPatchObjectV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigntemplatedocumentPatchObjectV1Request(objEzsigntemplatedocument EzsigntemplatedocumentRequestPatch) *EzsigntemplatedocumentPatchObjectV1Request {
	this := EzsigntemplatedocumentPatchObjectV1Request{}
	this.ObjEzsigntemplatedocument = objEzsigntemplatedocument
	return &this
}

// NewEzsigntemplatedocumentPatchObjectV1RequestWithDefaults instantiates a new EzsigntemplatedocumentPatchObjectV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigntemplatedocumentPatchObjectV1RequestWithDefaults() *EzsigntemplatedocumentPatchObjectV1Request {
	this := EzsigntemplatedocumentPatchObjectV1Request{}
	return &this
}

// GetObjEzsigntemplatedocument returns the ObjEzsigntemplatedocument field value
func (o *EzsigntemplatedocumentPatchObjectV1Request) GetObjEzsigntemplatedocument() EzsigntemplatedocumentRequestPatch {
	if o == nil {
		var ret EzsigntemplatedocumentRequestPatch
		return ret
	}

	return o.ObjEzsigntemplatedocument
}

// GetObjEzsigntemplatedocumentOk returns a tuple with the ObjEzsigntemplatedocument field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatedocumentPatchObjectV1Request) GetObjEzsigntemplatedocumentOk() (*EzsigntemplatedocumentRequestPatch, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjEzsigntemplatedocument, true
}

// SetObjEzsigntemplatedocument sets field value
func (o *EzsigntemplatedocumentPatchObjectV1Request) SetObjEzsigntemplatedocument(v EzsigntemplatedocumentRequestPatch) {
	o.ObjEzsigntemplatedocument = v
}

func (o EzsigntemplatedocumentPatchObjectV1Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigntemplatedocumentPatchObjectV1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objEzsigntemplatedocument"] = o.ObjEzsigntemplatedocument
	return toSerialize, nil
}

type NullableEzsigntemplatedocumentPatchObjectV1Request struct {
	value *EzsigntemplatedocumentPatchObjectV1Request
	isSet bool
}

func (v NullableEzsigntemplatedocumentPatchObjectV1Request) Get() *EzsigntemplatedocumentPatchObjectV1Request {
	return v.value
}

func (v *NullableEzsigntemplatedocumentPatchObjectV1Request) Set(val *EzsigntemplatedocumentPatchObjectV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigntemplatedocumentPatchObjectV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigntemplatedocumentPatchObjectV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigntemplatedocumentPatchObjectV1Request(val *EzsigntemplatedocumentPatchObjectV1Request) *NullableEzsigntemplatedocumentPatchObjectV1Request {
	return &NullableEzsigntemplatedocumentPatchObjectV1Request{value: val, isSet: true}
}

func (v NullableEzsigntemplatedocumentPatchObjectV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigntemplatedocumentPatchObjectV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

