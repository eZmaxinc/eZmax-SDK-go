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

// checks if the EzsigntemplatesignatureEditObjectV2Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigntemplatesignatureEditObjectV2Request{}

// EzsigntemplatesignatureEditObjectV2Request Request for PUT /2/object/ezsigntemplatesignature/{pkiEzsigntemplatesignatureID}
type EzsigntemplatesignatureEditObjectV2Request struct {
	ObjEzsigntemplatesignature EzsigntemplatesignatureRequestCompoundV2 `json:"objEzsigntemplatesignature"`
}

type _EzsigntemplatesignatureEditObjectV2Request EzsigntemplatesignatureEditObjectV2Request

// NewEzsigntemplatesignatureEditObjectV2Request instantiates a new EzsigntemplatesignatureEditObjectV2Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigntemplatesignatureEditObjectV2Request(objEzsigntemplatesignature EzsigntemplatesignatureRequestCompoundV2) *EzsigntemplatesignatureEditObjectV2Request {
	this := EzsigntemplatesignatureEditObjectV2Request{}
	this.ObjEzsigntemplatesignature = objEzsigntemplatesignature
	return &this
}

// NewEzsigntemplatesignatureEditObjectV2RequestWithDefaults instantiates a new EzsigntemplatesignatureEditObjectV2Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigntemplatesignatureEditObjectV2RequestWithDefaults() *EzsigntemplatesignatureEditObjectV2Request {
	this := EzsigntemplatesignatureEditObjectV2Request{}
	return &this
}

// GetObjEzsigntemplatesignature returns the ObjEzsigntemplatesignature field value
func (o *EzsigntemplatesignatureEditObjectV2Request) GetObjEzsigntemplatesignature() EzsigntemplatesignatureRequestCompoundV2 {
	if o == nil {
		var ret EzsigntemplatesignatureRequestCompoundV2
		return ret
	}

	return o.ObjEzsigntemplatesignature
}

// GetObjEzsigntemplatesignatureOk returns a tuple with the ObjEzsigntemplatesignature field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatesignatureEditObjectV2Request) GetObjEzsigntemplatesignatureOk() (*EzsigntemplatesignatureRequestCompoundV2, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjEzsigntemplatesignature, true
}

// SetObjEzsigntemplatesignature sets field value
func (o *EzsigntemplatesignatureEditObjectV2Request) SetObjEzsigntemplatesignature(v EzsigntemplatesignatureRequestCompoundV2) {
	o.ObjEzsigntemplatesignature = v
}

func (o EzsigntemplatesignatureEditObjectV2Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigntemplatesignatureEditObjectV2Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objEzsigntemplatesignature"] = o.ObjEzsigntemplatesignature
	return toSerialize, nil
}

func (o *EzsigntemplatesignatureEditObjectV2Request) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"objEzsigntemplatesignature",
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

	varEzsigntemplatesignatureEditObjectV2Request := _EzsigntemplatesignatureEditObjectV2Request{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsigntemplatesignatureEditObjectV2Request)

	if err != nil {
		return err
	}

	*o = EzsigntemplatesignatureEditObjectV2Request(varEzsigntemplatesignatureEditObjectV2Request)

	return err
}

type NullableEzsigntemplatesignatureEditObjectV2Request struct {
	value *EzsigntemplatesignatureEditObjectV2Request
	isSet bool
}

func (v NullableEzsigntemplatesignatureEditObjectV2Request) Get() *EzsigntemplatesignatureEditObjectV2Request {
	return v.value
}

func (v *NullableEzsigntemplatesignatureEditObjectV2Request) Set(val *EzsigntemplatesignatureEditObjectV2Request) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigntemplatesignatureEditObjectV2Request) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigntemplatesignatureEditObjectV2Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigntemplatesignatureEditObjectV2Request(val *EzsigntemplatesignatureEditObjectV2Request) *NullableEzsigntemplatesignatureEditObjectV2Request {
	return &NullableEzsigntemplatesignatureEditObjectV2Request{value: val, isSet: true}
}

func (v NullableEzsigntemplatesignatureEditObjectV2Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigntemplatesignatureEditObjectV2Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


