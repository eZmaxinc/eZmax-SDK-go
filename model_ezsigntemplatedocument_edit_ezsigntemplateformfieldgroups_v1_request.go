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
	"bytes"
	"fmt"
)

// checks if the EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Request{}

// EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Request Request for PUT /1/object/ezsigntemplatedocument/{pkiEzsigntemplatedocumentID}/editEzsigntemplateformfieldgroups
type EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Request struct {
	AObjEzsigntemplateformfieldgroup []EzsigntemplateformfieldgroupRequestCompound `json:"a_objEzsigntemplateformfieldgroup"`
}

type _EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Request EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Request

// NewEzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Request instantiates a new EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Request(aObjEzsigntemplateformfieldgroup []EzsigntemplateformfieldgroupRequestCompound) *EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Request {
	this := EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Request{}
	this.AObjEzsigntemplateformfieldgroup = aObjEzsigntemplateformfieldgroup
	return &this
}

// NewEzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1RequestWithDefaults instantiates a new EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1RequestWithDefaults() *EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Request {
	this := EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Request{}
	return &this
}

// GetAObjEzsigntemplateformfieldgroup returns the AObjEzsigntemplateformfieldgroup field value
func (o *EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Request) GetAObjEzsigntemplateformfieldgroup() []EzsigntemplateformfieldgroupRequestCompound {
	if o == nil {
		var ret []EzsigntemplateformfieldgroupRequestCompound
		return ret
	}

	return o.AObjEzsigntemplateformfieldgroup
}

// GetAObjEzsigntemplateformfieldgroupOk returns a tuple with the AObjEzsigntemplateformfieldgroup field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Request) GetAObjEzsigntemplateformfieldgroupOk() ([]EzsigntemplateformfieldgroupRequestCompound, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjEzsigntemplateformfieldgroup, true
}

// SetAObjEzsigntemplateformfieldgroup sets field value
func (o *EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Request) SetAObjEzsigntemplateformfieldgroup(v []EzsigntemplateformfieldgroupRequestCompound) {
	o.AObjEzsigntemplateformfieldgroup = v
}

func (o EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_objEzsigntemplateformfieldgroup"] = o.AObjEzsigntemplateformfieldgroup
	return toSerialize, nil
}

func (o *EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Request) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"a_objEzsigntemplateformfieldgroup",
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

	varEzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Request := _EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Request{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Request)

	if err != nil {
		return err
	}

	*o = EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Request(varEzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Request)

	return err
}

type NullableEzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Request struct {
	value *EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Request
	isSet bool
}

func (v NullableEzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Request) Get() *EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Request {
	return v.value
}

func (v *NullableEzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Request) Set(val *EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Request(val *EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Request) *NullableEzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Request {
	return &NullableEzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Request{value: val, isSet: true}
}

func (v NullableEzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


