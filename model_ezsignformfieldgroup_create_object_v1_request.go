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

// checks if the EzsignformfieldgroupCreateObjectV1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignformfieldgroupCreateObjectV1Request{}

// EzsignformfieldgroupCreateObjectV1Request Request for POST /1/object/ezsignformfieldgroup
type EzsignformfieldgroupCreateObjectV1Request struct {
	AObjEzsignformfieldgroup []EzsignformfieldgroupRequestCompound `json:"a_objEzsignformfieldgroup"`
}

type _EzsignformfieldgroupCreateObjectV1Request EzsignformfieldgroupCreateObjectV1Request

// NewEzsignformfieldgroupCreateObjectV1Request instantiates a new EzsignformfieldgroupCreateObjectV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignformfieldgroupCreateObjectV1Request(aObjEzsignformfieldgroup []EzsignformfieldgroupRequestCompound) *EzsignformfieldgroupCreateObjectV1Request {
	this := EzsignformfieldgroupCreateObjectV1Request{}
	this.AObjEzsignformfieldgroup = aObjEzsignformfieldgroup
	return &this
}

// NewEzsignformfieldgroupCreateObjectV1RequestWithDefaults instantiates a new EzsignformfieldgroupCreateObjectV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignformfieldgroupCreateObjectV1RequestWithDefaults() *EzsignformfieldgroupCreateObjectV1Request {
	this := EzsignformfieldgroupCreateObjectV1Request{}
	return &this
}

// GetAObjEzsignformfieldgroup returns the AObjEzsignformfieldgroup field value
func (o *EzsignformfieldgroupCreateObjectV1Request) GetAObjEzsignformfieldgroup() []EzsignformfieldgroupRequestCompound {
	if o == nil {
		var ret []EzsignformfieldgroupRequestCompound
		return ret
	}

	return o.AObjEzsignformfieldgroup
}

// GetAObjEzsignformfieldgroupOk returns a tuple with the AObjEzsignformfieldgroup field value
// and a boolean to check if the value has been set.
func (o *EzsignformfieldgroupCreateObjectV1Request) GetAObjEzsignformfieldgroupOk() ([]EzsignformfieldgroupRequestCompound, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjEzsignformfieldgroup, true
}

// SetAObjEzsignformfieldgroup sets field value
func (o *EzsignformfieldgroupCreateObjectV1Request) SetAObjEzsignformfieldgroup(v []EzsignformfieldgroupRequestCompound) {
	o.AObjEzsignformfieldgroup = v
}

func (o EzsignformfieldgroupCreateObjectV1Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignformfieldgroupCreateObjectV1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_objEzsignformfieldgroup"] = o.AObjEzsignformfieldgroup
	return toSerialize, nil
}

func (o *EzsignformfieldgroupCreateObjectV1Request) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"a_objEzsignformfieldgroup",
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

	varEzsignformfieldgroupCreateObjectV1Request := _EzsignformfieldgroupCreateObjectV1Request{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsignformfieldgroupCreateObjectV1Request)

	if err != nil {
		return err
	}

	*o = EzsignformfieldgroupCreateObjectV1Request(varEzsignformfieldgroupCreateObjectV1Request)

	return err
}

type NullableEzsignformfieldgroupCreateObjectV1Request struct {
	value *EzsignformfieldgroupCreateObjectV1Request
	isSet bool
}

func (v NullableEzsignformfieldgroupCreateObjectV1Request) Get() *EzsignformfieldgroupCreateObjectV1Request {
	return v.value
}

func (v *NullableEzsignformfieldgroupCreateObjectV1Request) Set(val *EzsignformfieldgroupCreateObjectV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignformfieldgroupCreateObjectV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignformfieldgroupCreateObjectV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignformfieldgroupCreateObjectV1Request(val *EzsignformfieldgroupCreateObjectV1Request) *NullableEzsignformfieldgroupCreateObjectV1Request {
	return &NullableEzsignformfieldgroupCreateObjectV1Request{value: val, isSet: true}
}

func (v NullableEzsignformfieldgroupCreateObjectV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignformfieldgroupCreateObjectV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


