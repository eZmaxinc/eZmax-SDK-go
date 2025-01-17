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

// checks if the ActivesessionGetCurrentV2Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActivesessionGetCurrentV2Response{}

// ActivesessionGetCurrentV2Response Response for GET /2/object/activesession/getCurrent
type ActivesessionGetCurrentV2Response struct {
	CommonResponse
	MPayload ActivesessionGetCurrentV2ResponseMPayload `json:"mPayload"`
}

type _ActivesessionGetCurrentV2Response ActivesessionGetCurrentV2Response

// NewActivesessionGetCurrentV2Response instantiates a new ActivesessionGetCurrentV2Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActivesessionGetCurrentV2Response(mPayload ActivesessionGetCurrentV2ResponseMPayload, objDebugPayload CommonResponseObjDebugPayload) *ActivesessionGetCurrentV2Response {
	this := ActivesessionGetCurrentV2Response{}
	this.ObjDebugPayload = objDebugPayload
	this.MPayload = mPayload
	return &this
}

// NewActivesessionGetCurrentV2ResponseWithDefaults instantiates a new ActivesessionGetCurrentV2Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActivesessionGetCurrentV2ResponseWithDefaults() *ActivesessionGetCurrentV2Response {
	this := ActivesessionGetCurrentV2Response{}
	return &this
}

// GetMPayload returns the MPayload field value
func (o *ActivesessionGetCurrentV2Response) GetMPayload() ActivesessionGetCurrentV2ResponseMPayload {
	if o == nil {
		var ret ActivesessionGetCurrentV2ResponseMPayload
		return ret
	}

	return o.MPayload
}

// GetMPayloadOk returns a tuple with the MPayload field value
// and a boolean to check if the value has been set.
func (o *ActivesessionGetCurrentV2Response) GetMPayloadOk() (*ActivesessionGetCurrentV2ResponseMPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MPayload, true
}

// SetMPayload sets field value
func (o *ActivesessionGetCurrentV2Response) SetMPayload(v ActivesessionGetCurrentV2ResponseMPayload) {
	o.MPayload = v
}

func (o ActivesessionGetCurrentV2Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActivesessionGetCurrentV2Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mPayload"] = o.MPayload
	return toSerialize, nil
}

func (o *ActivesessionGetCurrentV2Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"mPayload",
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

	varActivesessionGetCurrentV2Response := _ActivesessionGetCurrentV2Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varActivesessionGetCurrentV2Response)

	if err != nil {
		return err
	}

	*o = ActivesessionGetCurrentV2Response(varActivesessionGetCurrentV2Response)

	return err
}

type NullableActivesessionGetCurrentV2Response struct {
	value *ActivesessionGetCurrentV2Response
	isSet bool
}

func (v NullableActivesessionGetCurrentV2Response) Get() *ActivesessionGetCurrentV2Response {
	return v.value
}

func (v *NullableActivesessionGetCurrentV2Response) Set(val *ActivesessionGetCurrentV2Response) {
	v.value = val
	v.isSet = true
}

func (v NullableActivesessionGetCurrentV2Response) IsSet() bool {
	return v.isSet
}

func (v *NullableActivesessionGetCurrentV2Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActivesessionGetCurrentV2Response(val *ActivesessionGetCurrentV2Response) *NullableActivesessionGetCurrentV2Response {
	return &NullableActivesessionGetCurrentV2Response{value: val, isSet: true}
}

func (v NullableActivesessionGetCurrentV2Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActivesessionGetCurrentV2Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


