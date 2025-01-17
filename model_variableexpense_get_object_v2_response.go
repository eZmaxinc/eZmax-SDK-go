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

// checks if the VariableexpenseGetObjectV2Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VariableexpenseGetObjectV2Response{}

// VariableexpenseGetObjectV2Response Response for GET /2/object/variableexpense/{pkiVariableexpenseID}
type VariableexpenseGetObjectV2Response struct {
	CommonResponse
	MPayload VariableexpenseGetObjectV2ResponseMPayload `json:"mPayload"`
}

type _VariableexpenseGetObjectV2Response VariableexpenseGetObjectV2Response

// NewVariableexpenseGetObjectV2Response instantiates a new VariableexpenseGetObjectV2Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVariableexpenseGetObjectV2Response(mPayload VariableexpenseGetObjectV2ResponseMPayload, objDebugPayload CommonResponseObjDebugPayload) *VariableexpenseGetObjectV2Response {
	this := VariableexpenseGetObjectV2Response{}
	this.ObjDebugPayload = objDebugPayload
	this.MPayload = mPayload
	return &this
}

// NewVariableexpenseGetObjectV2ResponseWithDefaults instantiates a new VariableexpenseGetObjectV2Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVariableexpenseGetObjectV2ResponseWithDefaults() *VariableexpenseGetObjectV2Response {
	this := VariableexpenseGetObjectV2Response{}
	return &this
}

// GetMPayload returns the MPayload field value
func (o *VariableexpenseGetObjectV2Response) GetMPayload() VariableexpenseGetObjectV2ResponseMPayload {
	if o == nil {
		var ret VariableexpenseGetObjectV2ResponseMPayload
		return ret
	}

	return o.MPayload
}

// GetMPayloadOk returns a tuple with the MPayload field value
// and a boolean to check if the value has been set.
func (o *VariableexpenseGetObjectV2Response) GetMPayloadOk() (*VariableexpenseGetObjectV2ResponseMPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MPayload, true
}

// SetMPayload sets field value
func (o *VariableexpenseGetObjectV2Response) SetMPayload(v VariableexpenseGetObjectV2ResponseMPayload) {
	o.MPayload = v
}

func (o VariableexpenseGetObjectV2Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VariableexpenseGetObjectV2Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mPayload"] = o.MPayload
	return toSerialize, nil
}

func (o *VariableexpenseGetObjectV2Response) UnmarshalJSON(data []byte) (err error) {
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

	varVariableexpenseGetObjectV2Response := _VariableexpenseGetObjectV2Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVariableexpenseGetObjectV2Response)

	if err != nil {
		return err
	}

	*o = VariableexpenseGetObjectV2Response(varVariableexpenseGetObjectV2Response)

	return err
}

type NullableVariableexpenseGetObjectV2Response struct {
	value *VariableexpenseGetObjectV2Response
	isSet bool
}

func (v NullableVariableexpenseGetObjectV2Response) Get() *VariableexpenseGetObjectV2Response {
	return v.value
}

func (v *NullableVariableexpenseGetObjectV2Response) Set(val *VariableexpenseGetObjectV2Response) {
	v.value = val
	v.isSet = true
}

func (v NullableVariableexpenseGetObjectV2Response) IsSet() bool {
	return v.isSet
}

func (v *NullableVariableexpenseGetObjectV2Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVariableexpenseGetObjectV2Response(val *VariableexpenseGetObjectV2Response) *NullableVariableexpenseGetObjectV2Response {
	return &NullableVariableexpenseGetObjectV2Response{value: val, isSet: true}
}

func (v NullableVariableexpenseGetObjectV2Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVariableexpenseGetObjectV2Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


