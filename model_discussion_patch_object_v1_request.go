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

// checks if the DiscussionPatchObjectV1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DiscussionPatchObjectV1Request{}

// DiscussionPatchObjectV1Request Request for PATCH /1/object/discussion/{pkiDiscussionID}
type DiscussionPatchObjectV1Request struct {
	ObjDiscussion DiscussionRequestPatch `json:"objDiscussion"`
}

type _DiscussionPatchObjectV1Request DiscussionPatchObjectV1Request

// NewDiscussionPatchObjectV1Request instantiates a new DiscussionPatchObjectV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiscussionPatchObjectV1Request(objDiscussion DiscussionRequestPatch) *DiscussionPatchObjectV1Request {
	this := DiscussionPatchObjectV1Request{}
	this.ObjDiscussion = objDiscussion
	return &this
}

// NewDiscussionPatchObjectV1RequestWithDefaults instantiates a new DiscussionPatchObjectV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiscussionPatchObjectV1RequestWithDefaults() *DiscussionPatchObjectV1Request {
	this := DiscussionPatchObjectV1Request{}
	return &this
}

// GetObjDiscussion returns the ObjDiscussion field value
func (o *DiscussionPatchObjectV1Request) GetObjDiscussion() DiscussionRequestPatch {
	if o == nil {
		var ret DiscussionRequestPatch
		return ret
	}

	return o.ObjDiscussion
}

// GetObjDiscussionOk returns a tuple with the ObjDiscussion field value
// and a boolean to check if the value has been set.
func (o *DiscussionPatchObjectV1Request) GetObjDiscussionOk() (*DiscussionRequestPatch, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjDiscussion, true
}

// SetObjDiscussion sets field value
func (o *DiscussionPatchObjectV1Request) SetObjDiscussion(v DiscussionRequestPatch) {
	o.ObjDiscussion = v
}

func (o DiscussionPatchObjectV1Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DiscussionPatchObjectV1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objDiscussion"] = o.ObjDiscussion
	return toSerialize, nil
}

func (o *DiscussionPatchObjectV1Request) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"objDiscussion",
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

	varDiscussionPatchObjectV1Request := _DiscussionPatchObjectV1Request{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDiscussionPatchObjectV1Request)

	if err != nil {
		return err
	}

	*o = DiscussionPatchObjectV1Request(varDiscussionPatchObjectV1Request)

	return err
}

type NullableDiscussionPatchObjectV1Request struct {
	value *DiscussionPatchObjectV1Request
	isSet bool
}

func (v NullableDiscussionPatchObjectV1Request) Get() *DiscussionPatchObjectV1Request {
	return v.value
}

func (v *NullableDiscussionPatchObjectV1Request) Set(val *DiscussionPatchObjectV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableDiscussionPatchObjectV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableDiscussionPatchObjectV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiscussionPatchObjectV1Request(val *DiscussionPatchObjectV1Request) *NullableDiscussionPatchObjectV1Request {
	return &NullableDiscussionPatchObjectV1Request{value: val, isSet: true}
}

func (v NullableDiscussionPatchObjectV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiscussionPatchObjectV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


