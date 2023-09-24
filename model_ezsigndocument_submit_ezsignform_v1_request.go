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

// checks if the EzsigndocumentSubmitEzsignformV1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigndocumentSubmitEzsignformV1Request{}

// EzsigndocumentSubmitEzsignformV1Request Request for POST /1/object/ezsigndocument/{pkiEzsigndocumentID}/submitEzsignform
type EzsigndocumentSubmitEzsignformV1Request struct {
	// Whether the Ezsignform submitted is a draft or not.
	BEzsignformIsdraft bool `json:"bEzsignformIsdraft"`
	AObjEzsignformfieldgroup []CustomEzsignformfieldgroupRequest `json:"a_objEzsignformfieldgroup"`
}

// NewEzsigndocumentSubmitEzsignformV1Request instantiates a new EzsigndocumentSubmitEzsignformV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigndocumentSubmitEzsignformV1Request(bEzsignformIsdraft bool, aObjEzsignformfieldgroup []CustomEzsignformfieldgroupRequest) *EzsigndocumentSubmitEzsignformV1Request {
	this := EzsigndocumentSubmitEzsignformV1Request{}
	this.BEzsignformIsdraft = bEzsignformIsdraft
	this.AObjEzsignformfieldgroup = aObjEzsignformfieldgroup
	return &this
}

// NewEzsigndocumentSubmitEzsignformV1RequestWithDefaults instantiates a new EzsigndocumentSubmitEzsignformV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigndocumentSubmitEzsignformV1RequestWithDefaults() *EzsigndocumentSubmitEzsignformV1Request {
	this := EzsigndocumentSubmitEzsignformV1Request{}
	return &this
}

// GetBEzsignformIsdraft returns the BEzsignformIsdraft field value
func (o *EzsigndocumentSubmitEzsignformV1Request) GetBEzsignformIsdraft() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BEzsignformIsdraft
}

// GetBEzsignformIsdraftOk returns a tuple with the BEzsignformIsdraft field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentSubmitEzsignformV1Request) GetBEzsignformIsdraftOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BEzsignformIsdraft, true
}

// SetBEzsignformIsdraft sets field value
func (o *EzsigndocumentSubmitEzsignformV1Request) SetBEzsignformIsdraft(v bool) {
	o.BEzsignformIsdraft = v
}

// GetAObjEzsignformfieldgroup returns the AObjEzsignformfieldgroup field value
func (o *EzsigndocumentSubmitEzsignformV1Request) GetAObjEzsignformfieldgroup() []CustomEzsignformfieldgroupRequest {
	if o == nil {
		var ret []CustomEzsignformfieldgroupRequest
		return ret
	}

	return o.AObjEzsignformfieldgroup
}

// GetAObjEzsignformfieldgroupOk returns a tuple with the AObjEzsignformfieldgroup field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentSubmitEzsignformV1Request) GetAObjEzsignformfieldgroupOk() ([]CustomEzsignformfieldgroupRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjEzsignformfieldgroup, true
}

// SetAObjEzsignformfieldgroup sets field value
func (o *EzsigndocumentSubmitEzsignformV1Request) SetAObjEzsignformfieldgroup(v []CustomEzsignformfieldgroupRequest) {
	o.AObjEzsignformfieldgroup = v
}

func (o EzsigndocumentSubmitEzsignformV1Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigndocumentSubmitEzsignformV1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["bEzsignformIsdraft"] = o.BEzsignformIsdraft
	toSerialize["a_objEzsignformfieldgroup"] = o.AObjEzsignformfieldgroup
	return toSerialize, nil
}

type NullableEzsigndocumentSubmitEzsignformV1Request struct {
	value *EzsigndocumentSubmitEzsignformV1Request
	isSet bool
}

func (v NullableEzsigndocumentSubmitEzsignformV1Request) Get() *EzsigndocumentSubmitEzsignformV1Request {
	return v.value
}

func (v *NullableEzsigndocumentSubmitEzsignformV1Request) Set(val *EzsigndocumentSubmitEzsignformV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigndocumentSubmitEzsignformV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigndocumentSubmitEzsignformV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigndocumentSubmitEzsignformV1Request(val *EzsigndocumentSubmitEzsignformV1Request) *NullableEzsigndocumentSubmitEzsignformV1Request {
	return &NullableEzsigndocumentSubmitEzsignformV1Request{value: val, isSet: true}
}

func (v NullableEzsigndocumentSubmitEzsignformV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigndocumentSubmitEzsignformV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

