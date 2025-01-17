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

// checks if the EzsignsignatureCreateObjectV2Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignsignatureCreateObjectV2Response{}

// EzsignsignatureCreateObjectV2Response Response for POST /2/object/ezsignsignature
type EzsignsignatureCreateObjectV2Response struct {
	CommonResponse
	MPayload EzsignsignatureCreateObjectV2ResponseMPayload `json:"mPayload"`
}

type _EzsignsignatureCreateObjectV2Response EzsignsignatureCreateObjectV2Response

// NewEzsignsignatureCreateObjectV2Response instantiates a new EzsignsignatureCreateObjectV2Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignsignatureCreateObjectV2Response(mPayload EzsignsignatureCreateObjectV2ResponseMPayload, objDebugPayload CommonResponseObjDebugPayload) *EzsignsignatureCreateObjectV2Response {
	this := EzsignsignatureCreateObjectV2Response{}
	this.ObjDebugPayload = objDebugPayload
	this.MPayload = mPayload
	return &this
}

// NewEzsignsignatureCreateObjectV2ResponseWithDefaults instantiates a new EzsignsignatureCreateObjectV2Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignsignatureCreateObjectV2ResponseWithDefaults() *EzsignsignatureCreateObjectV2Response {
	this := EzsignsignatureCreateObjectV2Response{}
	return &this
}

// GetMPayload returns the MPayload field value
func (o *EzsignsignatureCreateObjectV2Response) GetMPayload() EzsignsignatureCreateObjectV2ResponseMPayload {
	if o == nil {
		var ret EzsignsignatureCreateObjectV2ResponseMPayload
		return ret
	}

	return o.MPayload
}

// GetMPayloadOk returns a tuple with the MPayload field value
// and a boolean to check if the value has been set.
func (o *EzsignsignatureCreateObjectV2Response) GetMPayloadOk() (*EzsignsignatureCreateObjectV2ResponseMPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MPayload, true
}

// SetMPayload sets field value
func (o *EzsignsignatureCreateObjectV2Response) SetMPayload(v EzsignsignatureCreateObjectV2ResponseMPayload) {
	o.MPayload = v
}

func (o EzsignsignatureCreateObjectV2Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignsignatureCreateObjectV2Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mPayload"] = o.MPayload
	return toSerialize, nil
}

func (o *EzsignsignatureCreateObjectV2Response) UnmarshalJSON(data []byte) (err error) {
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

	varEzsignsignatureCreateObjectV2Response := _EzsignsignatureCreateObjectV2Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsignsignatureCreateObjectV2Response)

	if err != nil {
		return err
	}

	*o = EzsignsignatureCreateObjectV2Response(varEzsignsignatureCreateObjectV2Response)

	return err
}

type NullableEzsignsignatureCreateObjectV2Response struct {
	value *EzsignsignatureCreateObjectV2Response
	isSet bool
}

func (v NullableEzsignsignatureCreateObjectV2Response) Get() *EzsignsignatureCreateObjectV2Response {
	return v.value
}

func (v *NullableEzsignsignatureCreateObjectV2Response) Set(val *EzsignsignatureCreateObjectV2Response) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignsignatureCreateObjectV2Response) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignsignatureCreateObjectV2Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignsignatureCreateObjectV2Response(val *EzsignsignatureCreateObjectV2Response) *NullableEzsignsignatureCreateObjectV2Response {
	return &NullableEzsignsignatureCreateObjectV2Response{value: val, isSet: true}
}

func (v NullableEzsignsignatureCreateObjectV2Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignsignatureCreateObjectV2Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


