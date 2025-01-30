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

// checks if the UsergroupmembershipCreateObjectV1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UsergroupmembershipCreateObjectV1Request{}

// UsergroupmembershipCreateObjectV1Request Request for POST /1/object/usergroupmembership
type UsergroupmembershipCreateObjectV1Request struct {
	AObjUsergroupmembership []UsergroupmembershipRequestCompound `json:"a_objUsergroupmembership"`
}

type _UsergroupmembershipCreateObjectV1Request UsergroupmembershipCreateObjectV1Request

// NewUsergroupmembershipCreateObjectV1Request instantiates a new UsergroupmembershipCreateObjectV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsergroupmembershipCreateObjectV1Request(aObjUsergroupmembership []UsergroupmembershipRequestCompound) *UsergroupmembershipCreateObjectV1Request {
	this := UsergroupmembershipCreateObjectV1Request{}
	this.AObjUsergroupmembership = aObjUsergroupmembership
	return &this
}

// NewUsergroupmembershipCreateObjectV1RequestWithDefaults instantiates a new UsergroupmembershipCreateObjectV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsergroupmembershipCreateObjectV1RequestWithDefaults() *UsergroupmembershipCreateObjectV1Request {
	this := UsergroupmembershipCreateObjectV1Request{}
	return &this
}

// GetAObjUsergroupmembership returns the AObjUsergroupmembership field value
func (o *UsergroupmembershipCreateObjectV1Request) GetAObjUsergroupmembership() []UsergroupmembershipRequestCompound {
	if o == nil {
		var ret []UsergroupmembershipRequestCompound
		return ret
	}

	return o.AObjUsergroupmembership
}

// GetAObjUsergroupmembershipOk returns a tuple with the AObjUsergroupmembership field value
// and a boolean to check if the value has been set.
func (o *UsergroupmembershipCreateObjectV1Request) GetAObjUsergroupmembershipOk() ([]UsergroupmembershipRequestCompound, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjUsergroupmembership, true
}

// SetAObjUsergroupmembership sets field value
func (o *UsergroupmembershipCreateObjectV1Request) SetAObjUsergroupmembership(v []UsergroupmembershipRequestCompound) {
	o.AObjUsergroupmembership = v
}

func (o UsergroupmembershipCreateObjectV1Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsergroupmembershipCreateObjectV1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_objUsergroupmembership"] = o.AObjUsergroupmembership
	return toSerialize, nil
}

func (o *UsergroupmembershipCreateObjectV1Request) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"a_objUsergroupmembership",
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

	varUsergroupmembershipCreateObjectV1Request := _UsergroupmembershipCreateObjectV1Request{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUsergroupmembershipCreateObjectV1Request)

	if err != nil {
		return err
	}

	*o = UsergroupmembershipCreateObjectV1Request(varUsergroupmembershipCreateObjectV1Request)

	return err
}

type NullableUsergroupmembershipCreateObjectV1Request struct {
	value *UsergroupmembershipCreateObjectV1Request
	isSet bool
}

func (v NullableUsergroupmembershipCreateObjectV1Request) Get() *UsergroupmembershipCreateObjectV1Request {
	return v.value
}

func (v *NullableUsergroupmembershipCreateObjectV1Request) Set(val *UsergroupmembershipCreateObjectV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableUsergroupmembershipCreateObjectV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableUsergroupmembershipCreateObjectV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsergroupmembershipCreateObjectV1Request(val *UsergroupmembershipCreateObjectV1Request) *NullableUsergroupmembershipCreateObjectV1Request {
	return &NullableUsergroupmembershipCreateObjectV1Request{value: val, isSet: true}
}

func (v NullableUsergroupmembershipCreateObjectV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsergroupmembershipCreateObjectV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


