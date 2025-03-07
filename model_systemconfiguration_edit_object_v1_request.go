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

// checks if the SystemconfigurationEditObjectV1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SystemconfigurationEditObjectV1Request{}

// SystemconfigurationEditObjectV1Request Request for PUT /1/object/systemconfiguration/{pkiSystemconfigurationID}
type SystemconfigurationEditObjectV1Request struct {
	ObjSystemconfiguration SystemconfigurationRequestCompound `json:"objSystemconfiguration"`
}

type _SystemconfigurationEditObjectV1Request SystemconfigurationEditObjectV1Request

// NewSystemconfigurationEditObjectV1Request instantiates a new SystemconfigurationEditObjectV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemconfigurationEditObjectV1Request(objSystemconfiguration SystemconfigurationRequestCompound) *SystemconfigurationEditObjectV1Request {
	this := SystemconfigurationEditObjectV1Request{}
	this.ObjSystemconfiguration = objSystemconfiguration
	return &this
}

// NewSystemconfigurationEditObjectV1RequestWithDefaults instantiates a new SystemconfigurationEditObjectV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemconfigurationEditObjectV1RequestWithDefaults() *SystemconfigurationEditObjectV1Request {
	this := SystemconfigurationEditObjectV1Request{}
	return &this
}

// GetObjSystemconfiguration returns the ObjSystemconfiguration field value
func (o *SystemconfigurationEditObjectV1Request) GetObjSystemconfiguration() SystemconfigurationRequestCompound {
	if o == nil {
		var ret SystemconfigurationRequestCompound
		return ret
	}

	return o.ObjSystemconfiguration
}

// GetObjSystemconfigurationOk returns a tuple with the ObjSystemconfiguration field value
// and a boolean to check if the value has been set.
func (o *SystemconfigurationEditObjectV1Request) GetObjSystemconfigurationOk() (*SystemconfigurationRequestCompound, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjSystemconfiguration, true
}

// SetObjSystemconfiguration sets field value
func (o *SystemconfigurationEditObjectV1Request) SetObjSystemconfiguration(v SystemconfigurationRequestCompound) {
	o.ObjSystemconfiguration = v
}

func (o SystemconfigurationEditObjectV1Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SystemconfigurationEditObjectV1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objSystemconfiguration"] = o.ObjSystemconfiguration
	return toSerialize, nil
}

func (o *SystemconfigurationEditObjectV1Request) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"objSystemconfiguration",
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

	varSystemconfigurationEditObjectV1Request := _SystemconfigurationEditObjectV1Request{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSystemconfigurationEditObjectV1Request)

	if err != nil {
		return err
	}

	*o = SystemconfigurationEditObjectV1Request(varSystemconfigurationEditObjectV1Request)

	return err
}

type NullableSystemconfigurationEditObjectV1Request struct {
	value *SystemconfigurationEditObjectV1Request
	isSet bool
}

func (v NullableSystemconfigurationEditObjectV1Request) Get() *SystemconfigurationEditObjectV1Request {
	return v.value
}

func (v *NullableSystemconfigurationEditObjectV1Request) Set(val *SystemconfigurationEditObjectV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemconfigurationEditObjectV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemconfigurationEditObjectV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemconfigurationEditObjectV1Request(val *SystemconfigurationEditObjectV1Request) *NullableSystemconfigurationEditObjectV1Request {
	return &NullableSystemconfigurationEditObjectV1Request{value: val, isSet: true}
}

func (v NullableSystemconfigurationEditObjectV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemconfigurationEditObjectV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


