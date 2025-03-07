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

// checks if the DiscussionmembershipCreateObjectV1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DiscussionmembershipCreateObjectV1Request{}

// DiscussionmembershipCreateObjectV1Request Request for POST /1/object/discussionmembership
type DiscussionmembershipCreateObjectV1Request struct {
	AObjDiscussionmembership []DiscussionmembershipRequestCompound `json:"a_objDiscussionmembership"`
}

type _DiscussionmembershipCreateObjectV1Request DiscussionmembershipCreateObjectV1Request

// NewDiscussionmembershipCreateObjectV1Request instantiates a new DiscussionmembershipCreateObjectV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiscussionmembershipCreateObjectV1Request(aObjDiscussionmembership []DiscussionmembershipRequestCompound) *DiscussionmembershipCreateObjectV1Request {
	this := DiscussionmembershipCreateObjectV1Request{}
	this.AObjDiscussionmembership = aObjDiscussionmembership
	return &this
}

// NewDiscussionmembershipCreateObjectV1RequestWithDefaults instantiates a new DiscussionmembershipCreateObjectV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiscussionmembershipCreateObjectV1RequestWithDefaults() *DiscussionmembershipCreateObjectV1Request {
	this := DiscussionmembershipCreateObjectV1Request{}
	return &this
}

// GetAObjDiscussionmembership returns the AObjDiscussionmembership field value
func (o *DiscussionmembershipCreateObjectV1Request) GetAObjDiscussionmembership() []DiscussionmembershipRequestCompound {
	if o == nil {
		var ret []DiscussionmembershipRequestCompound
		return ret
	}

	return o.AObjDiscussionmembership
}

// GetAObjDiscussionmembershipOk returns a tuple with the AObjDiscussionmembership field value
// and a boolean to check if the value has been set.
func (o *DiscussionmembershipCreateObjectV1Request) GetAObjDiscussionmembershipOk() ([]DiscussionmembershipRequestCompound, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjDiscussionmembership, true
}

// SetAObjDiscussionmembership sets field value
func (o *DiscussionmembershipCreateObjectV1Request) SetAObjDiscussionmembership(v []DiscussionmembershipRequestCompound) {
	o.AObjDiscussionmembership = v
}

func (o DiscussionmembershipCreateObjectV1Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DiscussionmembershipCreateObjectV1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_objDiscussionmembership"] = o.AObjDiscussionmembership
	return toSerialize, nil
}

func (o *DiscussionmembershipCreateObjectV1Request) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"a_objDiscussionmembership",
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

	varDiscussionmembershipCreateObjectV1Request := _DiscussionmembershipCreateObjectV1Request{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDiscussionmembershipCreateObjectV1Request)

	if err != nil {
		return err
	}

	*o = DiscussionmembershipCreateObjectV1Request(varDiscussionmembershipCreateObjectV1Request)

	return err
}

type NullableDiscussionmembershipCreateObjectV1Request struct {
	value *DiscussionmembershipCreateObjectV1Request
	isSet bool
}

func (v NullableDiscussionmembershipCreateObjectV1Request) Get() *DiscussionmembershipCreateObjectV1Request {
	return v.value
}

func (v *NullableDiscussionmembershipCreateObjectV1Request) Set(val *DiscussionmembershipCreateObjectV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableDiscussionmembershipCreateObjectV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableDiscussionmembershipCreateObjectV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiscussionmembershipCreateObjectV1Request(val *DiscussionmembershipCreateObjectV1Request) *NullableDiscussionmembershipCreateObjectV1Request {
	return &NullableDiscussionmembershipCreateObjectV1Request{value: val, isSet: true}
}

func (v NullableDiscussionmembershipCreateObjectV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiscussionmembershipCreateObjectV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


