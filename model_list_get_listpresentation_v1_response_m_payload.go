/*
eZmax API Definition

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.1.2
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
)

// ListGetListpresentationV1ResponseMPayload Payload for the GET /1/module/list/listpresentation/{sListName} API Request
type ListGetListpresentationV1ResponseMPayload struct {
	AObjListpresentation []ListpresentationRequest `json:"a_objListpresentation"`
}

// NewListGetListpresentationV1ResponseMPayload instantiates a new ListGetListpresentationV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListGetListpresentationV1ResponseMPayload(aObjListpresentation []ListpresentationRequest) *ListGetListpresentationV1ResponseMPayload {
	this := ListGetListpresentationV1ResponseMPayload{}
	this.AObjListpresentation = aObjListpresentation
	return &this
}

// NewListGetListpresentationV1ResponseMPayloadWithDefaults instantiates a new ListGetListpresentationV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListGetListpresentationV1ResponseMPayloadWithDefaults() *ListGetListpresentationV1ResponseMPayload {
	this := ListGetListpresentationV1ResponseMPayload{}
	return &this
}

// GetAObjListpresentation returns the AObjListpresentation field value
func (o *ListGetListpresentationV1ResponseMPayload) GetAObjListpresentation() []ListpresentationRequest {
	if o == nil {
		var ret []ListpresentationRequest
		return ret
	}

	return o.AObjListpresentation
}

// GetAObjListpresentationOk returns a tuple with the AObjListpresentation field value
// and a boolean to check if the value has been set.
func (o *ListGetListpresentationV1ResponseMPayload) GetAObjListpresentationOk() (*[]ListpresentationRequest, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AObjListpresentation, true
}

// SetAObjListpresentation sets field value
func (o *ListGetListpresentationV1ResponseMPayload) SetAObjListpresentation(v []ListpresentationRequest) {
	o.AObjListpresentation = v
}

func (o ListGetListpresentationV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["a_objListpresentation"] = o.AObjListpresentation
	}
	return json.Marshal(toSerialize)
}

type NullableListGetListpresentationV1ResponseMPayload struct {
	value *ListGetListpresentationV1ResponseMPayload
	isSet bool
}

func (v NullableListGetListpresentationV1ResponseMPayload) Get() *ListGetListpresentationV1ResponseMPayload {
	return v.value
}

func (v *NullableListGetListpresentationV1ResponseMPayload) Set(val *ListGetListpresentationV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableListGetListpresentationV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableListGetListpresentationV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListGetListpresentationV1ResponseMPayload(val *ListGetListpresentationV1ResponseMPayload) *NullableListGetListpresentationV1ResponseMPayload {
	return &NullableListGetListpresentationV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableListGetListpresentationV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListGetListpresentationV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


