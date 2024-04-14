/*
eZmax API Definition (Full)

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.2.1
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the EzsigndocumentCreateEzsignelementsPositionedByWordV1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigndocumentCreateEzsignelementsPositionedByWordV1Request{}

// EzsigndocumentCreateEzsignelementsPositionedByWordV1Request Request for POST /1/object/ezsigndocument/{pkiEzsigndocumentID}/createEzsignelementsPositionedByWord
type EzsigndocumentCreateEzsignelementsPositionedByWordV1Request struct {
	AObjEzsignformfieldgroup []CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest `json:"a_objEzsignformfieldgroup"`
	AObjEzsignsignature []CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest `json:"a_objEzsignsignature"`
}

type _EzsigndocumentCreateEzsignelementsPositionedByWordV1Request EzsigndocumentCreateEzsignelementsPositionedByWordV1Request

// NewEzsigndocumentCreateEzsignelementsPositionedByWordV1Request instantiates a new EzsigndocumentCreateEzsignelementsPositionedByWordV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigndocumentCreateEzsignelementsPositionedByWordV1Request(aObjEzsignformfieldgroup []CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest, aObjEzsignsignature []CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) *EzsigndocumentCreateEzsignelementsPositionedByWordV1Request {
	this := EzsigndocumentCreateEzsignelementsPositionedByWordV1Request{}
	this.AObjEzsignformfieldgroup = aObjEzsignformfieldgroup
	this.AObjEzsignsignature = aObjEzsignsignature
	return &this
}

// NewEzsigndocumentCreateEzsignelementsPositionedByWordV1RequestWithDefaults instantiates a new EzsigndocumentCreateEzsignelementsPositionedByWordV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigndocumentCreateEzsignelementsPositionedByWordV1RequestWithDefaults() *EzsigndocumentCreateEzsignelementsPositionedByWordV1Request {
	this := EzsigndocumentCreateEzsignelementsPositionedByWordV1Request{}
	return &this
}

// GetAObjEzsignformfieldgroup returns the AObjEzsignformfieldgroup field value
func (o *EzsigndocumentCreateEzsignelementsPositionedByWordV1Request) GetAObjEzsignformfieldgroup() []CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest {
	if o == nil {
		var ret []CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest
		return ret
	}

	return o.AObjEzsignformfieldgroup
}

// GetAObjEzsignformfieldgroupOk returns a tuple with the AObjEzsignformfieldgroup field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentCreateEzsignelementsPositionedByWordV1Request) GetAObjEzsignformfieldgroupOk() ([]CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjEzsignformfieldgroup, true
}

// SetAObjEzsignformfieldgroup sets field value
func (o *EzsigndocumentCreateEzsignelementsPositionedByWordV1Request) SetAObjEzsignformfieldgroup(v []CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) {
	o.AObjEzsignformfieldgroup = v
}

// GetAObjEzsignsignature returns the AObjEzsignsignature field value
func (o *EzsigndocumentCreateEzsignelementsPositionedByWordV1Request) GetAObjEzsignsignature() []CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest {
	if o == nil {
		var ret []CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest
		return ret
	}

	return o.AObjEzsignsignature
}

// GetAObjEzsignsignatureOk returns a tuple with the AObjEzsignsignature field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentCreateEzsignelementsPositionedByWordV1Request) GetAObjEzsignsignatureOk() ([]CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjEzsignsignature, true
}

// SetAObjEzsignsignature sets field value
func (o *EzsigndocumentCreateEzsignelementsPositionedByWordV1Request) SetAObjEzsignsignature(v []CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) {
	o.AObjEzsignsignature = v
}

func (o EzsigndocumentCreateEzsignelementsPositionedByWordV1Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigndocumentCreateEzsignelementsPositionedByWordV1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_objEzsignformfieldgroup"] = o.AObjEzsignformfieldgroup
	toSerialize["a_objEzsignsignature"] = o.AObjEzsignsignature
	return toSerialize, nil
}

func (o *EzsigndocumentCreateEzsignelementsPositionedByWordV1Request) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"a_objEzsignformfieldgroup",
		"a_objEzsignsignature",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varEzsigndocumentCreateEzsignelementsPositionedByWordV1Request := _EzsigndocumentCreateEzsignelementsPositionedByWordV1Request{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsigndocumentCreateEzsignelementsPositionedByWordV1Request)

	if err != nil {
		return err
	}

	*o = EzsigndocumentCreateEzsignelementsPositionedByWordV1Request(varEzsigndocumentCreateEzsignelementsPositionedByWordV1Request)

	return err
}

type NullableEzsigndocumentCreateEzsignelementsPositionedByWordV1Request struct {
	value *EzsigndocumentCreateEzsignelementsPositionedByWordV1Request
	isSet bool
}

func (v NullableEzsigndocumentCreateEzsignelementsPositionedByWordV1Request) Get() *EzsigndocumentCreateEzsignelementsPositionedByWordV1Request {
	return v.value
}

func (v *NullableEzsigndocumentCreateEzsignelementsPositionedByWordV1Request) Set(val *EzsigndocumentCreateEzsignelementsPositionedByWordV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigndocumentCreateEzsignelementsPositionedByWordV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigndocumentCreateEzsignelementsPositionedByWordV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigndocumentCreateEzsignelementsPositionedByWordV1Request(val *EzsigndocumentCreateEzsignelementsPositionedByWordV1Request) *NullableEzsigndocumentCreateEzsignelementsPositionedByWordV1Request {
	return &NullableEzsigndocumentCreateEzsignelementsPositionedByWordV1Request{value: val, isSet: true}
}

func (v NullableEzsigndocumentCreateEzsignelementsPositionedByWordV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigndocumentCreateEzsignelementsPositionedByWordV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


