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

// checks if the EzsignannotationGetObjectV2Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignannotationGetObjectV2Response{}

// EzsignannotationGetObjectV2Response Response for GET /2/object/ezsignannotation/{pkiEzsignannotationID}
type EzsignannotationGetObjectV2Response struct {
	CommonResponse
	MPayload EzsignannotationGetObjectV2ResponseMPayload `json:"mPayload"`
}

type _EzsignannotationGetObjectV2Response EzsignannotationGetObjectV2Response

// NewEzsignannotationGetObjectV2Response instantiates a new EzsignannotationGetObjectV2Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignannotationGetObjectV2Response(mPayload EzsignannotationGetObjectV2ResponseMPayload, objDebugPayload CommonResponseObjDebugPayload) *EzsignannotationGetObjectV2Response {
	this := EzsignannotationGetObjectV2Response{}
	this.ObjDebugPayload = objDebugPayload
	this.MPayload = mPayload
	return &this
}

// NewEzsignannotationGetObjectV2ResponseWithDefaults instantiates a new EzsignannotationGetObjectV2Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignannotationGetObjectV2ResponseWithDefaults() *EzsignannotationGetObjectV2Response {
	this := EzsignannotationGetObjectV2Response{}
	return &this
}

// GetMPayload returns the MPayload field value
func (o *EzsignannotationGetObjectV2Response) GetMPayload() EzsignannotationGetObjectV2ResponseMPayload {
	if o == nil {
		var ret EzsignannotationGetObjectV2ResponseMPayload
		return ret
	}

	return o.MPayload
}

// GetMPayloadOk returns a tuple with the MPayload field value
// and a boolean to check if the value has been set.
func (o *EzsignannotationGetObjectV2Response) GetMPayloadOk() (*EzsignannotationGetObjectV2ResponseMPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MPayload, true
}

// SetMPayload sets field value
func (o *EzsignannotationGetObjectV2Response) SetMPayload(v EzsignannotationGetObjectV2ResponseMPayload) {
	o.MPayload = v
}

func (o EzsignannotationGetObjectV2Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignannotationGetObjectV2Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mPayload"] = o.MPayload
	return toSerialize, nil
}

func (o *EzsignannotationGetObjectV2Response) UnmarshalJSON(data []byte) (err error) {
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

	varEzsignannotationGetObjectV2Response := _EzsignannotationGetObjectV2Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsignannotationGetObjectV2Response)

	if err != nil {
		return err
	}

	*o = EzsignannotationGetObjectV2Response(varEzsignannotationGetObjectV2Response)

	return err
}

type NullableEzsignannotationGetObjectV2Response struct {
	value *EzsignannotationGetObjectV2Response
	isSet bool
}

func (v NullableEzsignannotationGetObjectV2Response) Get() *EzsignannotationGetObjectV2Response {
	return v.value
}

func (v *NullableEzsignannotationGetObjectV2Response) Set(val *EzsignannotationGetObjectV2Response) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignannotationGetObjectV2Response) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignannotationGetObjectV2Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignannotationGetObjectV2Response(val *EzsignannotationGetObjectV2Response) *NullableEzsignannotationGetObjectV2Response {
	return &NullableEzsignannotationGetObjectV2Response{value: val, isSet: true}
}

func (v NullableEzsignannotationGetObjectV2Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignannotationGetObjectV2Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


