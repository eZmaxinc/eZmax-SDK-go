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

// checks if the DiscussionmessagePatchObjectV1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DiscussionmessagePatchObjectV1Request{}

// DiscussionmessagePatchObjectV1Request Request for PATCH /1/object/discussionmessage/{pkiDiscussionmessageID}
type DiscussionmessagePatchObjectV1Request struct {
	ObjDiscussionmessage DiscussionmessageRequestPatch `json:"objDiscussionmessage"`
}

type _DiscussionmessagePatchObjectV1Request DiscussionmessagePatchObjectV1Request

// NewDiscussionmessagePatchObjectV1Request instantiates a new DiscussionmessagePatchObjectV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiscussionmessagePatchObjectV1Request(objDiscussionmessage DiscussionmessageRequestPatch) *DiscussionmessagePatchObjectV1Request {
	this := DiscussionmessagePatchObjectV1Request{}
	this.ObjDiscussionmessage = objDiscussionmessage
	return &this
}

// NewDiscussionmessagePatchObjectV1RequestWithDefaults instantiates a new DiscussionmessagePatchObjectV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiscussionmessagePatchObjectV1RequestWithDefaults() *DiscussionmessagePatchObjectV1Request {
	this := DiscussionmessagePatchObjectV1Request{}
	return &this
}

// GetObjDiscussionmessage returns the ObjDiscussionmessage field value
func (o *DiscussionmessagePatchObjectV1Request) GetObjDiscussionmessage() DiscussionmessageRequestPatch {
	if o == nil {
		var ret DiscussionmessageRequestPatch
		return ret
	}

	return o.ObjDiscussionmessage
}

// GetObjDiscussionmessageOk returns a tuple with the ObjDiscussionmessage field value
// and a boolean to check if the value has been set.
func (o *DiscussionmessagePatchObjectV1Request) GetObjDiscussionmessageOk() (*DiscussionmessageRequestPatch, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjDiscussionmessage, true
}

// SetObjDiscussionmessage sets field value
func (o *DiscussionmessagePatchObjectV1Request) SetObjDiscussionmessage(v DiscussionmessageRequestPatch) {
	o.ObjDiscussionmessage = v
}

func (o DiscussionmessagePatchObjectV1Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DiscussionmessagePatchObjectV1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objDiscussionmessage"] = o.ObjDiscussionmessage
	return toSerialize, nil
}

func (o *DiscussionmessagePatchObjectV1Request) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"objDiscussionmessage",
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

	varDiscussionmessagePatchObjectV1Request := _DiscussionmessagePatchObjectV1Request{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDiscussionmessagePatchObjectV1Request)

	if err != nil {
		return err
	}

	*o = DiscussionmessagePatchObjectV1Request(varDiscussionmessagePatchObjectV1Request)

	return err
}

type NullableDiscussionmessagePatchObjectV1Request struct {
	value *DiscussionmessagePatchObjectV1Request
	isSet bool
}

func (v NullableDiscussionmessagePatchObjectV1Request) Get() *DiscussionmessagePatchObjectV1Request {
	return v.value
}

func (v *NullableDiscussionmessagePatchObjectV1Request) Set(val *DiscussionmessagePatchObjectV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableDiscussionmessagePatchObjectV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableDiscussionmessagePatchObjectV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiscussionmessagePatchObjectV1Request(val *DiscussionmessagePatchObjectV1Request) *NullableDiscussionmessagePatchObjectV1Request {
	return &NullableDiscussionmessagePatchObjectV1Request{value: val, isSet: true}
}

func (v NullableDiscussionmessagePatchObjectV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiscussionmessagePatchObjectV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

