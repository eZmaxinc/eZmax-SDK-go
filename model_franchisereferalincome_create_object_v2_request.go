/*
eZmax API Definition (Full)

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.2.2
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the FranchisereferalincomeCreateObjectV2Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FranchisereferalincomeCreateObjectV2Request{}

// FranchisereferalincomeCreateObjectV2Request Request for POST /2/object/franchisereferalincome
type FranchisereferalincomeCreateObjectV2Request struct {
	AObjFranchisereferalincome []FranchisereferalincomeRequestCompound `json:"a_objFranchisereferalincome"`
}

type _FranchisereferalincomeCreateObjectV2Request FranchisereferalincomeCreateObjectV2Request

// NewFranchisereferalincomeCreateObjectV2Request instantiates a new FranchisereferalincomeCreateObjectV2Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFranchisereferalincomeCreateObjectV2Request(aObjFranchisereferalincome []FranchisereferalincomeRequestCompound) *FranchisereferalincomeCreateObjectV2Request {
	this := FranchisereferalincomeCreateObjectV2Request{}
	this.AObjFranchisereferalincome = aObjFranchisereferalincome
	return &this
}

// NewFranchisereferalincomeCreateObjectV2RequestWithDefaults instantiates a new FranchisereferalincomeCreateObjectV2Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFranchisereferalincomeCreateObjectV2RequestWithDefaults() *FranchisereferalincomeCreateObjectV2Request {
	this := FranchisereferalincomeCreateObjectV2Request{}
	return &this
}

// GetAObjFranchisereferalincome returns the AObjFranchisereferalincome field value
func (o *FranchisereferalincomeCreateObjectV2Request) GetAObjFranchisereferalincome() []FranchisereferalincomeRequestCompound {
	if o == nil {
		var ret []FranchisereferalincomeRequestCompound
		return ret
	}

	return o.AObjFranchisereferalincome
}

// GetAObjFranchisereferalincomeOk returns a tuple with the AObjFranchisereferalincome field value
// and a boolean to check if the value has been set.
func (o *FranchisereferalincomeCreateObjectV2Request) GetAObjFranchisereferalincomeOk() ([]FranchisereferalincomeRequestCompound, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjFranchisereferalincome, true
}

// SetAObjFranchisereferalincome sets field value
func (o *FranchisereferalincomeCreateObjectV2Request) SetAObjFranchisereferalincome(v []FranchisereferalincomeRequestCompound) {
	o.AObjFranchisereferalincome = v
}

func (o FranchisereferalincomeCreateObjectV2Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FranchisereferalincomeCreateObjectV2Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_objFranchisereferalincome"] = o.AObjFranchisereferalincome
	return toSerialize, nil
}

func (o *FranchisereferalincomeCreateObjectV2Request) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"a_objFranchisereferalincome",
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

	varFranchisereferalincomeCreateObjectV2Request := _FranchisereferalincomeCreateObjectV2Request{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFranchisereferalincomeCreateObjectV2Request)

	if err != nil {
		return err
	}

	*o = FranchisereferalincomeCreateObjectV2Request(varFranchisereferalincomeCreateObjectV2Request)

	return err
}

type NullableFranchisereferalincomeCreateObjectV2Request struct {
	value *FranchisereferalincomeCreateObjectV2Request
	isSet bool
}

func (v NullableFranchisereferalincomeCreateObjectV2Request) Get() *FranchisereferalincomeCreateObjectV2Request {
	return v.value
}

func (v *NullableFranchisereferalincomeCreateObjectV2Request) Set(val *FranchisereferalincomeCreateObjectV2Request) {
	v.value = val
	v.isSet = true
}

func (v NullableFranchisereferalincomeCreateObjectV2Request) IsSet() bool {
	return v.isSet
}

func (v *NullableFranchisereferalincomeCreateObjectV2Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFranchisereferalincomeCreateObjectV2Request(val *FranchisereferalincomeCreateObjectV2Request) *NullableFranchisereferalincomeCreateObjectV2Request {
	return &NullableFranchisereferalincomeCreateObjectV2Request{value: val, isSet: true}
}

func (v NullableFranchisereferalincomeCreateObjectV2Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFranchisereferalincomeCreateObjectV2Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


