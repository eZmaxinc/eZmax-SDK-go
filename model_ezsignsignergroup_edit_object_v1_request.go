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

// checks if the EzsignsignergroupEditObjectV1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignsignergroupEditObjectV1Request{}

// EzsignsignergroupEditObjectV1Request Request for PUT /1/object/ezsignsignergroup/{pkiEzsignsignergroupID}
type EzsignsignergroupEditObjectV1Request struct {
	ObjEzsignsignergroup EzsignsignergroupRequestCompound `json:"objEzsignsignergroup"`
}

type _EzsignsignergroupEditObjectV1Request EzsignsignergroupEditObjectV1Request

// NewEzsignsignergroupEditObjectV1Request instantiates a new EzsignsignergroupEditObjectV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignsignergroupEditObjectV1Request(objEzsignsignergroup EzsignsignergroupRequestCompound) *EzsignsignergroupEditObjectV1Request {
	this := EzsignsignergroupEditObjectV1Request{}
	this.ObjEzsignsignergroup = objEzsignsignergroup
	return &this
}

// NewEzsignsignergroupEditObjectV1RequestWithDefaults instantiates a new EzsignsignergroupEditObjectV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignsignergroupEditObjectV1RequestWithDefaults() *EzsignsignergroupEditObjectV1Request {
	this := EzsignsignergroupEditObjectV1Request{}
	return &this
}

// GetObjEzsignsignergroup returns the ObjEzsignsignergroup field value
func (o *EzsignsignergroupEditObjectV1Request) GetObjEzsignsignergroup() EzsignsignergroupRequestCompound {
	if o == nil {
		var ret EzsignsignergroupRequestCompound
		return ret
	}

	return o.ObjEzsignsignergroup
}

// GetObjEzsignsignergroupOk returns a tuple with the ObjEzsignsignergroup field value
// and a boolean to check if the value has been set.
func (o *EzsignsignergroupEditObjectV1Request) GetObjEzsignsignergroupOk() (*EzsignsignergroupRequestCompound, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjEzsignsignergroup, true
}

// SetObjEzsignsignergroup sets field value
func (o *EzsignsignergroupEditObjectV1Request) SetObjEzsignsignergroup(v EzsignsignergroupRequestCompound) {
	o.ObjEzsignsignergroup = v
}

func (o EzsignsignergroupEditObjectV1Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignsignergroupEditObjectV1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objEzsignsignergroup"] = o.ObjEzsignsignergroup
	return toSerialize, nil
}

func (o *EzsignsignergroupEditObjectV1Request) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"objEzsignsignergroup",
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

	varEzsignsignergroupEditObjectV1Request := _EzsignsignergroupEditObjectV1Request{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsignsignergroupEditObjectV1Request)

	if err != nil {
		return err
	}

	*o = EzsignsignergroupEditObjectV1Request(varEzsignsignergroupEditObjectV1Request)

	return err
}

type NullableEzsignsignergroupEditObjectV1Request struct {
	value *EzsignsignergroupEditObjectV1Request
	isSet bool
}

func (v NullableEzsignsignergroupEditObjectV1Request) Get() *EzsignsignergroupEditObjectV1Request {
	return v.value
}

func (v *NullableEzsignsignergroupEditObjectV1Request) Set(val *EzsignsignergroupEditObjectV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignsignergroupEditObjectV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignsignergroupEditObjectV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignsignergroupEditObjectV1Request(val *EzsignsignergroupEditObjectV1Request) *NullableEzsignsignergroupEditObjectV1Request {
	return &NullableEzsignsignergroupEditObjectV1Request{value: val, isSet: true}
}

func (v NullableEzsignsignergroupEditObjectV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignsignergroupEditObjectV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


