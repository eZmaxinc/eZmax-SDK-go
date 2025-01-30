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

// checks if the DiscussionCreateObjectV1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DiscussionCreateObjectV1Request{}

// DiscussionCreateObjectV1Request Request for POST /1/object/discussion
type DiscussionCreateObjectV1Request struct {
	AObjDiscussion []DiscussionRequestCompound `json:"a_objDiscussion"`
}

type _DiscussionCreateObjectV1Request DiscussionCreateObjectV1Request

// NewDiscussionCreateObjectV1Request instantiates a new DiscussionCreateObjectV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiscussionCreateObjectV1Request(aObjDiscussion []DiscussionRequestCompound) *DiscussionCreateObjectV1Request {
	this := DiscussionCreateObjectV1Request{}
	this.AObjDiscussion = aObjDiscussion
	return &this
}

// NewDiscussionCreateObjectV1RequestWithDefaults instantiates a new DiscussionCreateObjectV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiscussionCreateObjectV1RequestWithDefaults() *DiscussionCreateObjectV1Request {
	this := DiscussionCreateObjectV1Request{}
	return &this
}

// GetAObjDiscussion returns the AObjDiscussion field value
func (o *DiscussionCreateObjectV1Request) GetAObjDiscussion() []DiscussionRequestCompound {
	if o == nil {
		var ret []DiscussionRequestCompound
		return ret
	}

	return o.AObjDiscussion
}

// GetAObjDiscussionOk returns a tuple with the AObjDiscussion field value
// and a boolean to check if the value has been set.
func (o *DiscussionCreateObjectV1Request) GetAObjDiscussionOk() ([]DiscussionRequestCompound, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjDiscussion, true
}

// SetAObjDiscussion sets field value
func (o *DiscussionCreateObjectV1Request) SetAObjDiscussion(v []DiscussionRequestCompound) {
	o.AObjDiscussion = v
}

func (o DiscussionCreateObjectV1Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DiscussionCreateObjectV1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_objDiscussion"] = o.AObjDiscussion
	return toSerialize, nil
}

func (o *DiscussionCreateObjectV1Request) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"a_objDiscussion",
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

	varDiscussionCreateObjectV1Request := _DiscussionCreateObjectV1Request{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDiscussionCreateObjectV1Request)

	if err != nil {
		return err
	}

	*o = DiscussionCreateObjectV1Request(varDiscussionCreateObjectV1Request)

	return err
}

type NullableDiscussionCreateObjectV1Request struct {
	value *DiscussionCreateObjectV1Request
	isSet bool
}

func (v NullableDiscussionCreateObjectV1Request) Get() *DiscussionCreateObjectV1Request {
	return v.value
}

func (v *NullableDiscussionCreateObjectV1Request) Set(val *DiscussionCreateObjectV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableDiscussionCreateObjectV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableDiscussionCreateObjectV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiscussionCreateObjectV1Request(val *DiscussionCreateObjectV1Request) *NullableDiscussionCreateObjectV1Request {
	return &NullableDiscussionCreateObjectV1Request{value: val, isSet: true}
}

func (v NullableDiscussionCreateObjectV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiscussionCreateObjectV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


