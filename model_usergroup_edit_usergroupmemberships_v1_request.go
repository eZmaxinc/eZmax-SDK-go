/*
eZmax API Definition (Full)

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.2.1
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UsergroupEditUsergroupmembershipsV1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UsergroupEditUsergroupmembershipsV1Request{}

// UsergroupEditUsergroupmembershipsV1Request Request for PUT /1/object/usergroup/{pkiUsergroupID}/editUsergroupmemberships
type UsergroupEditUsergroupmembershipsV1Request struct {
	AObjUsergroupmembership []UsergroupmembershipRequestCompound `json:"a_objUsergroupmembership"`
}

type _UsergroupEditUsergroupmembershipsV1Request UsergroupEditUsergroupmembershipsV1Request

// NewUsergroupEditUsergroupmembershipsV1Request instantiates a new UsergroupEditUsergroupmembershipsV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsergroupEditUsergroupmembershipsV1Request(aObjUsergroupmembership []UsergroupmembershipRequestCompound) *UsergroupEditUsergroupmembershipsV1Request {
	this := UsergroupEditUsergroupmembershipsV1Request{}
	this.AObjUsergroupmembership = aObjUsergroupmembership
	return &this
}

// NewUsergroupEditUsergroupmembershipsV1RequestWithDefaults instantiates a new UsergroupEditUsergroupmembershipsV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsergroupEditUsergroupmembershipsV1RequestWithDefaults() *UsergroupEditUsergroupmembershipsV1Request {
	this := UsergroupEditUsergroupmembershipsV1Request{}
	return &this
}

// GetAObjUsergroupmembership returns the AObjUsergroupmembership field value
func (o *UsergroupEditUsergroupmembershipsV1Request) GetAObjUsergroupmembership() []UsergroupmembershipRequestCompound {
	if o == nil {
		var ret []UsergroupmembershipRequestCompound
		return ret
	}

	return o.AObjUsergroupmembership
}

// GetAObjUsergroupmembershipOk returns a tuple with the AObjUsergroupmembership field value
// and a boolean to check if the value has been set.
func (o *UsergroupEditUsergroupmembershipsV1Request) GetAObjUsergroupmembershipOk() ([]UsergroupmembershipRequestCompound, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjUsergroupmembership, true
}

// SetAObjUsergroupmembership sets field value
func (o *UsergroupEditUsergroupmembershipsV1Request) SetAObjUsergroupmembership(v []UsergroupmembershipRequestCompound) {
	o.AObjUsergroupmembership = v
}

func (o UsergroupEditUsergroupmembershipsV1Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsergroupEditUsergroupmembershipsV1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_objUsergroupmembership"] = o.AObjUsergroupmembership
	return toSerialize, nil
}

func (o *UsergroupEditUsergroupmembershipsV1Request) UnmarshalJSON(data []byte) (err error) {
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

	varUsergroupEditUsergroupmembershipsV1Request := _UsergroupEditUsergroupmembershipsV1Request{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUsergroupEditUsergroupmembershipsV1Request)

	if err != nil {
		return err
	}

	*o = UsergroupEditUsergroupmembershipsV1Request(varUsergroupEditUsergroupmembershipsV1Request)

	return err
}

type NullableUsergroupEditUsergroupmembershipsV1Request struct {
	value *UsergroupEditUsergroupmembershipsV1Request
	isSet bool
}

func (v NullableUsergroupEditUsergroupmembershipsV1Request) Get() *UsergroupEditUsergroupmembershipsV1Request {
	return v.value
}

func (v *NullableUsergroupEditUsergroupmembershipsV1Request) Set(val *UsergroupEditUsergroupmembershipsV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableUsergroupEditUsergroupmembershipsV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableUsergroupEditUsergroupmembershipsV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsergroupEditUsergroupmembershipsV1Request(val *UsergroupEditUsergroupmembershipsV1Request) *NullableUsergroupEditUsergroupmembershipsV1Request {
	return &NullableUsergroupEditUsergroupmembershipsV1Request{value: val, isSet: true}
}

func (v NullableUsergroupEditUsergroupmembershipsV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsergroupEditUsergroupmembershipsV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


