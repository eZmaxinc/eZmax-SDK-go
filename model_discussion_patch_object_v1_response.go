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

// checks if the DiscussionPatchObjectV1Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DiscussionPatchObjectV1Response{}

// DiscussionPatchObjectV1Response Response for PATCH /1/object/discussion/{pkiDiscussionID}
type DiscussionPatchObjectV1Response struct {
	CommonResponse
}

type _DiscussionPatchObjectV1Response DiscussionPatchObjectV1Response

// NewDiscussionPatchObjectV1Response instantiates a new DiscussionPatchObjectV1Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiscussionPatchObjectV1Response(objDebugPayload CommonResponseObjDebugPayload) *DiscussionPatchObjectV1Response {
	this := DiscussionPatchObjectV1Response{}
	this.ObjDebugPayload = objDebugPayload
	return &this
}

// NewDiscussionPatchObjectV1ResponseWithDefaults instantiates a new DiscussionPatchObjectV1Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiscussionPatchObjectV1ResponseWithDefaults() *DiscussionPatchObjectV1Response {
	this := DiscussionPatchObjectV1Response{}
	return &this
}

func (o DiscussionPatchObjectV1Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DiscussionPatchObjectV1Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

func (o *DiscussionPatchObjectV1Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"objDebugPayload",
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

	varDiscussionPatchObjectV1Response := _DiscussionPatchObjectV1Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDiscussionPatchObjectV1Response)

	if err != nil {
		return err
	}

	*o = DiscussionPatchObjectV1Response(varDiscussionPatchObjectV1Response)

	return err
}

type NullableDiscussionPatchObjectV1Response struct {
	value *DiscussionPatchObjectV1Response
	isSet bool
}

func (v NullableDiscussionPatchObjectV1Response) Get() *DiscussionPatchObjectV1Response {
	return v.value
}

func (v *NullableDiscussionPatchObjectV1Response) Set(val *DiscussionPatchObjectV1Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDiscussionPatchObjectV1Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDiscussionPatchObjectV1Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiscussionPatchObjectV1Response(val *DiscussionPatchObjectV1Response) *NullableDiscussionPatchObjectV1Response {
	return &NullableDiscussionPatchObjectV1Response{value: val, isSet: true}
}

func (v NullableDiscussionPatchObjectV1Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiscussionPatchObjectV1Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


