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

// checks if the EzsigntemplatepublicEditObjectV1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigntemplatepublicEditObjectV1Request{}

// EzsigntemplatepublicEditObjectV1Request Request for PUT /1/object/ezsigntemplatepublic/{pkiEzsigntemplatepublicID}
type EzsigntemplatepublicEditObjectV1Request struct {
	ObjEzsigntemplatepublic EzsigntemplatepublicRequestCompound `json:"objEzsigntemplatepublic"`
}

type _EzsigntemplatepublicEditObjectV1Request EzsigntemplatepublicEditObjectV1Request

// NewEzsigntemplatepublicEditObjectV1Request instantiates a new EzsigntemplatepublicEditObjectV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigntemplatepublicEditObjectV1Request(objEzsigntemplatepublic EzsigntemplatepublicRequestCompound) *EzsigntemplatepublicEditObjectV1Request {
	this := EzsigntemplatepublicEditObjectV1Request{}
	this.ObjEzsigntemplatepublic = objEzsigntemplatepublic
	return &this
}

// NewEzsigntemplatepublicEditObjectV1RequestWithDefaults instantiates a new EzsigntemplatepublicEditObjectV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigntemplatepublicEditObjectV1RequestWithDefaults() *EzsigntemplatepublicEditObjectV1Request {
	this := EzsigntemplatepublicEditObjectV1Request{}
	return &this
}

// GetObjEzsigntemplatepublic returns the ObjEzsigntemplatepublic field value
func (o *EzsigntemplatepublicEditObjectV1Request) GetObjEzsigntemplatepublic() EzsigntemplatepublicRequestCompound {
	if o == nil {
		var ret EzsigntemplatepublicRequestCompound
		return ret
	}

	return o.ObjEzsigntemplatepublic
}

// GetObjEzsigntemplatepublicOk returns a tuple with the ObjEzsigntemplatepublic field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepublicEditObjectV1Request) GetObjEzsigntemplatepublicOk() (*EzsigntemplatepublicRequestCompound, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjEzsigntemplatepublic, true
}

// SetObjEzsigntemplatepublic sets field value
func (o *EzsigntemplatepublicEditObjectV1Request) SetObjEzsigntemplatepublic(v EzsigntemplatepublicRequestCompound) {
	o.ObjEzsigntemplatepublic = v
}

func (o EzsigntemplatepublicEditObjectV1Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigntemplatepublicEditObjectV1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objEzsigntemplatepublic"] = o.ObjEzsigntemplatepublic
	return toSerialize, nil
}

func (o *EzsigntemplatepublicEditObjectV1Request) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"objEzsigntemplatepublic",
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

	varEzsigntemplatepublicEditObjectV1Request := _EzsigntemplatepublicEditObjectV1Request{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsigntemplatepublicEditObjectV1Request)

	if err != nil {
		return err
	}

	*o = EzsigntemplatepublicEditObjectV1Request(varEzsigntemplatepublicEditObjectV1Request)

	return err
}

type NullableEzsigntemplatepublicEditObjectV1Request struct {
	value *EzsigntemplatepublicEditObjectV1Request
	isSet bool
}

func (v NullableEzsigntemplatepublicEditObjectV1Request) Get() *EzsigntemplatepublicEditObjectV1Request {
	return v.value
}

func (v *NullableEzsigntemplatepublicEditObjectV1Request) Set(val *EzsigntemplatepublicEditObjectV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigntemplatepublicEditObjectV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigntemplatepublicEditObjectV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigntemplatepublicEditObjectV1Request(val *EzsigntemplatepublicEditObjectV1Request) *NullableEzsigntemplatepublicEditObjectV1Request {
	return &NullableEzsigntemplatepublicEditObjectV1Request{value: val, isSet: true}
}

func (v NullableEzsigntemplatepublicEditObjectV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigntemplatepublicEditObjectV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


