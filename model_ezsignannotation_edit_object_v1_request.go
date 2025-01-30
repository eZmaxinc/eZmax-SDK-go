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

// checks if the EzsignannotationEditObjectV1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignannotationEditObjectV1Request{}

// EzsignannotationEditObjectV1Request Request for PUT /1/object/ezsignannotation/{pkiEzsignannotationID}
type EzsignannotationEditObjectV1Request struct {
	ObjEzsignannotation EzsignannotationRequestCompound `json:"objEzsignannotation"`
}

type _EzsignannotationEditObjectV1Request EzsignannotationEditObjectV1Request

// NewEzsignannotationEditObjectV1Request instantiates a new EzsignannotationEditObjectV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignannotationEditObjectV1Request(objEzsignannotation EzsignannotationRequestCompound) *EzsignannotationEditObjectV1Request {
	this := EzsignannotationEditObjectV1Request{}
	this.ObjEzsignannotation = objEzsignannotation
	return &this
}

// NewEzsignannotationEditObjectV1RequestWithDefaults instantiates a new EzsignannotationEditObjectV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignannotationEditObjectV1RequestWithDefaults() *EzsignannotationEditObjectV1Request {
	this := EzsignannotationEditObjectV1Request{}
	return &this
}

// GetObjEzsignannotation returns the ObjEzsignannotation field value
func (o *EzsignannotationEditObjectV1Request) GetObjEzsignannotation() EzsignannotationRequestCompound {
	if o == nil {
		var ret EzsignannotationRequestCompound
		return ret
	}

	return o.ObjEzsignannotation
}

// GetObjEzsignannotationOk returns a tuple with the ObjEzsignannotation field value
// and a boolean to check if the value has been set.
func (o *EzsignannotationEditObjectV1Request) GetObjEzsignannotationOk() (*EzsignannotationRequestCompound, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjEzsignannotation, true
}

// SetObjEzsignannotation sets field value
func (o *EzsignannotationEditObjectV1Request) SetObjEzsignannotation(v EzsignannotationRequestCompound) {
	o.ObjEzsignannotation = v
}

func (o EzsignannotationEditObjectV1Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignannotationEditObjectV1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objEzsignannotation"] = o.ObjEzsignannotation
	return toSerialize, nil
}

func (o *EzsignannotationEditObjectV1Request) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"objEzsignannotation",
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

	varEzsignannotationEditObjectV1Request := _EzsignannotationEditObjectV1Request{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsignannotationEditObjectV1Request)

	if err != nil {
		return err
	}

	*o = EzsignannotationEditObjectV1Request(varEzsignannotationEditObjectV1Request)

	return err
}

type NullableEzsignannotationEditObjectV1Request struct {
	value *EzsignannotationEditObjectV1Request
	isSet bool
}

func (v NullableEzsignannotationEditObjectV1Request) Get() *EzsignannotationEditObjectV1Request {
	return v.value
}

func (v *NullableEzsignannotationEditObjectV1Request) Set(val *EzsignannotationEditObjectV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignannotationEditObjectV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignannotationEditObjectV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignannotationEditObjectV1Request(val *EzsignannotationEditObjectV1Request) *NullableEzsignannotationEditObjectV1Request {
	return &NullableEzsignannotationEditObjectV1Request{value: val, isSet: true}
}

func (v NullableEzsignannotationEditObjectV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignannotationEditObjectV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


