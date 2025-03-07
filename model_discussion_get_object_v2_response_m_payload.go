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

// checks if the DiscussionGetObjectV2ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DiscussionGetObjectV2ResponseMPayload{}

// DiscussionGetObjectV2ResponseMPayload Payload for GET /2/object/discussion/{pkiDiscussionID}
type DiscussionGetObjectV2ResponseMPayload struct {
	ObjDiscussion DiscussionResponseCompound `json:"objDiscussion"`
}

type _DiscussionGetObjectV2ResponseMPayload DiscussionGetObjectV2ResponseMPayload

// NewDiscussionGetObjectV2ResponseMPayload instantiates a new DiscussionGetObjectV2ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiscussionGetObjectV2ResponseMPayload(objDiscussion DiscussionResponseCompound) *DiscussionGetObjectV2ResponseMPayload {
	this := DiscussionGetObjectV2ResponseMPayload{}
	this.ObjDiscussion = objDiscussion
	return &this
}

// NewDiscussionGetObjectV2ResponseMPayloadWithDefaults instantiates a new DiscussionGetObjectV2ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiscussionGetObjectV2ResponseMPayloadWithDefaults() *DiscussionGetObjectV2ResponseMPayload {
	this := DiscussionGetObjectV2ResponseMPayload{}
	return &this
}

// GetObjDiscussion returns the ObjDiscussion field value
func (o *DiscussionGetObjectV2ResponseMPayload) GetObjDiscussion() DiscussionResponseCompound {
	if o == nil {
		var ret DiscussionResponseCompound
		return ret
	}

	return o.ObjDiscussion
}

// GetObjDiscussionOk returns a tuple with the ObjDiscussion field value
// and a boolean to check if the value has been set.
func (o *DiscussionGetObjectV2ResponseMPayload) GetObjDiscussionOk() (*DiscussionResponseCompound, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjDiscussion, true
}

// SetObjDiscussion sets field value
func (o *DiscussionGetObjectV2ResponseMPayload) SetObjDiscussion(v DiscussionResponseCompound) {
	o.ObjDiscussion = v
}

func (o DiscussionGetObjectV2ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DiscussionGetObjectV2ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objDiscussion"] = o.ObjDiscussion
	return toSerialize, nil
}

func (o *DiscussionGetObjectV2ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
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

	varDiscussionGetObjectV2ResponseMPayload := _DiscussionGetObjectV2ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDiscussionGetObjectV2ResponseMPayload)

	if err != nil {
		return err
	}

	*o = DiscussionGetObjectV2ResponseMPayload(varDiscussionGetObjectV2ResponseMPayload)

	return err
}

type NullableDiscussionGetObjectV2ResponseMPayload struct {
	value *DiscussionGetObjectV2ResponseMPayload
	isSet bool
}

func (v NullableDiscussionGetObjectV2ResponseMPayload) Get() *DiscussionGetObjectV2ResponseMPayload {
	return v.value
}

func (v *NullableDiscussionGetObjectV2ResponseMPayload) Set(val *DiscussionGetObjectV2ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableDiscussionGetObjectV2ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableDiscussionGetObjectV2ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiscussionGetObjectV2ResponseMPayload(val *DiscussionGetObjectV2ResponseMPayload) *NullableDiscussionGetObjectV2ResponseMPayload {
	return &NullableDiscussionGetObjectV2ResponseMPayload{value: val, isSet: true}
}

func (v NullableDiscussionGetObjectV2ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiscussionGetObjectV2ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


