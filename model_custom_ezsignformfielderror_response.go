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

// checks if the CustomEzsignformfielderrorResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomEzsignformfielderrorResponse{}

// CustomEzsignformfielderrorResponse A Custom Ezsignformfield Object to contain an error list
type CustomEzsignformfielderrorResponse struct {
	// The Label for the Ezsignformfield
	SEzsignformfieldLabel string `json:"sEzsignformfieldLabel"`
	// 
	AObjEzsignformfielderrortest []CustomEzsignformfielderrortestResponse `json:"a_objEzsignformfielderrortest"`
}

// NewCustomEzsignformfielderrorResponse instantiates a new CustomEzsignformfielderrorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomEzsignformfielderrorResponse(sEzsignformfieldLabel string, aObjEzsignformfielderrortest []CustomEzsignformfielderrortestResponse) *CustomEzsignformfielderrorResponse {
	this := CustomEzsignformfielderrorResponse{}
	this.SEzsignformfieldLabel = sEzsignformfieldLabel
	this.AObjEzsignformfielderrortest = aObjEzsignformfielderrortest
	return &this
}

// NewCustomEzsignformfielderrorResponseWithDefaults instantiates a new CustomEzsignformfielderrorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomEzsignformfielderrorResponseWithDefaults() *CustomEzsignformfielderrorResponse {
	this := CustomEzsignformfielderrorResponse{}
	return &this
}

// GetSEzsignformfieldLabel returns the SEzsignformfieldLabel field value
func (o *CustomEzsignformfielderrorResponse) GetSEzsignformfieldLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzsignformfieldLabel
}

// GetSEzsignformfieldLabelOk returns a tuple with the SEzsignformfieldLabel field value
// and a boolean to check if the value has been set.
func (o *CustomEzsignformfielderrorResponse) GetSEzsignformfieldLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SEzsignformfieldLabel, true
}

// SetSEzsignformfieldLabel sets field value
func (o *CustomEzsignformfielderrorResponse) SetSEzsignformfieldLabel(v string) {
	o.SEzsignformfieldLabel = v
}

// GetAObjEzsignformfielderrortest returns the AObjEzsignformfielderrortest field value
func (o *CustomEzsignformfielderrorResponse) GetAObjEzsignformfielderrortest() []CustomEzsignformfielderrortestResponse {
	if o == nil {
		var ret []CustomEzsignformfielderrortestResponse
		return ret
	}

	return o.AObjEzsignformfielderrortest
}

// GetAObjEzsignformfielderrortestOk returns a tuple with the AObjEzsignformfielderrortest field value
// and a boolean to check if the value has been set.
func (o *CustomEzsignformfielderrorResponse) GetAObjEzsignformfielderrortestOk() ([]CustomEzsignformfielderrortestResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjEzsignformfielderrortest, true
}

// SetAObjEzsignformfielderrortest sets field value
func (o *CustomEzsignformfielderrorResponse) SetAObjEzsignformfielderrortest(v []CustomEzsignformfielderrortestResponse) {
	o.AObjEzsignformfielderrortest = v
}

func (o CustomEzsignformfielderrorResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomEzsignformfielderrorResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sEzsignformfieldLabel"] = o.SEzsignformfieldLabel
	toSerialize["a_objEzsignformfielderrortest"] = o.AObjEzsignformfielderrortest
	return toSerialize, nil
}

type NullableCustomEzsignformfielderrorResponse struct {
	value *CustomEzsignformfielderrorResponse
	isSet bool
}

func (v NullableCustomEzsignformfielderrorResponse) Get() *CustomEzsignformfielderrorResponse {
	return v.value
}

func (v *NullableCustomEzsignformfielderrorResponse) Set(val *CustomEzsignformfielderrorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomEzsignformfielderrorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomEzsignformfielderrorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomEzsignformfielderrorResponse(val *CustomEzsignformfielderrorResponse) *NullableCustomEzsignformfielderrorResponse {
	return &NullableCustomEzsignformfielderrorResponse{value: val, isSet: true}
}

func (v NullableCustomEzsignformfielderrorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomEzsignformfielderrorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

