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

// checks if the DiscussionCreateObjectV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DiscussionCreateObjectV1ResponseMPayload{}

// DiscussionCreateObjectV1ResponseMPayload Payload for POST /1/object/discussion
type DiscussionCreateObjectV1ResponseMPayload struct {
	// An array of unique IDs representing the object that were requested to be created.  They are returned in the same order as the array containing the objects to be created that was sent in the request.
	APkiDiscussionID []int32 `json:"a_pkiDiscussionID"`
}

type _DiscussionCreateObjectV1ResponseMPayload DiscussionCreateObjectV1ResponseMPayload

// NewDiscussionCreateObjectV1ResponseMPayload instantiates a new DiscussionCreateObjectV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiscussionCreateObjectV1ResponseMPayload(aPkiDiscussionID []int32) *DiscussionCreateObjectV1ResponseMPayload {
	this := DiscussionCreateObjectV1ResponseMPayload{}
	this.APkiDiscussionID = aPkiDiscussionID
	return &this
}

// NewDiscussionCreateObjectV1ResponseMPayloadWithDefaults instantiates a new DiscussionCreateObjectV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiscussionCreateObjectV1ResponseMPayloadWithDefaults() *DiscussionCreateObjectV1ResponseMPayload {
	this := DiscussionCreateObjectV1ResponseMPayload{}
	return &this
}

// GetAPkiDiscussionID returns the APkiDiscussionID field value
func (o *DiscussionCreateObjectV1ResponseMPayload) GetAPkiDiscussionID() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.APkiDiscussionID
}

// GetAPkiDiscussionIDOk returns a tuple with the APkiDiscussionID field value
// and a boolean to check if the value has been set.
func (o *DiscussionCreateObjectV1ResponseMPayload) GetAPkiDiscussionIDOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.APkiDiscussionID, true
}

// SetAPkiDiscussionID sets field value
func (o *DiscussionCreateObjectV1ResponseMPayload) SetAPkiDiscussionID(v []int32) {
	o.APkiDiscussionID = v
}

func (o DiscussionCreateObjectV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DiscussionCreateObjectV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_pkiDiscussionID"] = o.APkiDiscussionID
	return toSerialize, nil
}

func (o *DiscussionCreateObjectV1ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"a_pkiDiscussionID",
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

	varDiscussionCreateObjectV1ResponseMPayload := _DiscussionCreateObjectV1ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDiscussionCreateObjectV1ResponseMPayload)

	if err != nil {
		return err
	}

	*o = DiscussionCreateObjectV1ResponseMPayload(varDiscussionCreateObjectV1ResponseMPayload)

	return err
}

type NullableDiscussionCreateObjectV1ResponseMPayload struct {
	value *DiscussionCreateObjectV1ResponseMPayload
	isSet bool
}

func (v NullableDiscussionCreateObjectV1ResponseMPayload) Get() *DiscussionCreateObjectV1ResponseMPayload {
	return v.value
}

func (v *NullableDiscussionCreateObjectV1ResponseMPayload) Set(val *DiscussionCreateObjectV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableDiscussionCreateObjectV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableDiscussionCreateObjectV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiscussionCreateObjectV1ResponseMPayload(val *DiscussionCreateObjectV1ResponseMPayload) *NullableDiscussionCreateObjectV1ResponseMPayload {
	return &NullableDiscussionCreateObjectV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableDiscussionCreateObjectV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiscussionCreateObjectV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


