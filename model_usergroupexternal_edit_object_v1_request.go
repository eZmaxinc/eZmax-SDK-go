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

// checks if the UsergroupexternalEditObjectV1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UsergroupexternalEditObjectV1Request{}

// UsergroupexternalEditObjectV1Request Request for PUT /1/object/usergroupexternal/{pkiUsergroupexternalID}
type UsergroupexternalEditObjectV1Request struct {
	ObjUsergroupexternal UsergroupexternalRequestCompound `json:"objUsergroupexternal"`
}

type _UsergroupexternalEditObjectV1Request UsergroupexternalEditObjectV1Request

// NewUsergroupexternalEditObjectV1Request instantiates a new UsergroupexternalEditObjectV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsergroupexternalEditObjectV1Request(objUsergroupexternal UsergroupexternalRequestCompound) *UsergroupexternalEditObjectV1Request {
	this := UsergroupexternalEditObjectV1Request{}
	this.ObjUsergroupexternal = objUsergroupexternal
	return &this
}

// NewUsergroupexternalEditObjectV1RequestWithDefaults instantiates a new UsergroupexternalEditObjectV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsergroupexternalEditObjectV1RequestWithDefaults() *UsergroupexternalEditObjectV1Request {
	this := UsergroupexternalEditObjectV1Request{}
	return &this
}

// GetObjUsergroupexternal returns the ObjUsergroupexternal field value
func (o *UsergroupexternalEditObjectV1Request) GetObjUsergroupexternal() UsergroupexternalRequestCompound {
	if o == nil {
		var ret UsergroupexternalRequestCompound
		return ret
	}

	return o.ObjUsergroupexternal
}

// GetObjUsergroupexternalOk returns a tuple with the ObjUsergroupexternal field value
// and a boolean to check if the value has been set.
func (o *UsergroupexternalEditObjectV1Request) GetObjUsergroupexternalOk() (*UsergroupexternalRequestCompound, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjUsergroupexternal, true
}

// SetObjUsergroupexternal sets field value
func (o *UsergroupexternalEditObjectV1Request) SetObjUsergroupexternal(v UsergroupexternalRequestCompound) {
	o.ObjUsergroupexternal = v
}

func (o UsergroupexternalEditObjectV1Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsergroupexternalEditObjectV1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objUsergroupexternal"] = o.ObjUsergroupexternal
	return toSerialize, nil
}

func (o *UsergroupexternalEditObjectV1Request) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"objUsergroupexternal",
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

	varUsergroupexternalEditObjectV1Request := _UsergroupexternalEditObjectV1Request{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUsergroupexternalEditObjectV1Request)

	if err != nil {
		return err
	}

	*o = UsergroupexternalEditObjectV1Request(varUsergroupexternalEditObjectV1Request)

	return err
}

type NullableUsergroupexternalEditObjectV1Request struct {
	value *UsergroupexternalEditObjectV1Request
	isSet bool
}

func (v NullableUsergroupexternalEditObjectV1Request) Get() *UsergroupexternalEditObjectV1Request {
	return v.value
}

func (v *NullableUsergroupexternalEditObjectV1Request) Set(val *UsergroupexternalEditObjectV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableUsergroupexternalEditObjectV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableUsergroupexternalEditObjectV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsergroupexternalEditObjectV1Request(val *UsergroupexternalEditObjectV1Request) *NullableUsergroupexternalEditObjectV1Request {
	return &NullableUsergroupexternalEditObjectV1Request{value: val, isSet: true}
}

func (v NullableUsergroupexternalEditObjectV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsergroupexternalEditObjectV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


