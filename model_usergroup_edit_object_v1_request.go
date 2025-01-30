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

// checks if the UsergroupEditObjectV1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UsergroupEditObjectV1Request{}

// UsergroupEditObjectV1Request Request for PUT /1/object/usergroup/{pkiUsergroupID}
type UsergroupEditObjectV1Request struct {
	ObjUsergroup UsergroupRequestCompound `json:"objUsergroup"`
}

type _UsergroupEditObjectV1Request UsergroupEditObjectV1Request

// NewUsergroupEditObjectV1Request instantiates a new UsergroupEditObjectV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsergroupEditObjectV1Request(objUsergroup UsergroupRequestCompound) *UsergroupEditObjectV1Request {
	this := UsergroupEditObjectV1Request{}
	this.ObjUsergroup = objUsergroup
	return &this
}

// NewUsergroupEditObjectV1RequestWithDefaults instantiates a new UsergroupEditObjectV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsergroupEditObjectV1RequestWithDefaults() *UsergroupEditObjectV1Request {
	this := UsergroupEditObjectV1Request{}
	return &this
}

// GetObjUsergroup returns the ObjUsergroup field value
func (o *UsergroupEditObjectV1Request) GetObjUsergroup() UsergroupRequestCompound {
	if o == nil {
		var ret UsergroupRequestCompound
		return ret
	}

	return o.ObjUsergroup
}

// GetObjUsergroupOk returns a tuple with the ObjUsergroup field value
// and a boolean to check if the value has been set.
func (o *UsergroupEditObjectV1Request) GetObjUsergroupOk() (*UsergroupRequestCompound, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjUsergroup, true
}

// SetObjUsergroup sets field value
func (o *UsergroupEditObjectV1Request) SetObjUsergroup(v UsergroupRequestCompound) {
	o.ObjUsergroup = v
}

func (o UsergroupEditObjectV1Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsergroupEditObjectV1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objUsergroup"] = o.ObjUsergroup
	return toSerialize, nil
}

func (o *UsergroupEditObjectV1Request) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"objUsergroup",
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

	varUsergroupEditObjectV1Request := _UsergroupEditObjectV1Request{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUsergroupEditObjectV1Request)

	if err != nil {
		return err
	}

	*o = UsergroupEditObjectV1Request(varUsergroupEditObjectV1Request)

	return err
}

type NullableUsergroupEditObjectV1Request struct {
	value *UsergroupEditObjectV1Request
	isSet bool
}

func (v NullableUsergroupEditObjectV1Request) Get() *UsergroupEditObjectV1Request {
	return v.value
}

func (v *NullableUsergroupEditObjectV1Request) Set(val *UsergroupEditObjectV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableUsergroupEditObjectV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableUsergroupEditObjectV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsergroupEditObjectV1Request(val *UsergroupEditObjectV1Request) *NullableUsergroupEditObjectV1Request {
	return &NullableUsergroupEditObjectV1Request{value: val, isSet: true}
}

func (v NullableUsergroupEditObjectV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsergroupEditObjectV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


