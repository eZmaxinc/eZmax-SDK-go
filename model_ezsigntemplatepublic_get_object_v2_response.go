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

// checks if the EzsigntemplatepublicGetObjectV2Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigntemplatepublicGetObjectV2Response{}

// EzsigntemplatepublicGetObjectV2Response Response for GET /2/object/ezsigntemplatepublic/{pkiEzsigntemplatepublicID}
type EzsigntemplatepublicGetObjectV2Response struct {
	CommonResponse
	MPayload EzsigntemplatepublicGetObjectV2ResponseMPayload `json:"mPayload"`
}

type _EzsigntemplatepublicGetObjectV2Response EzsigntemplatepublicGetObjectV2Response

// NewEzsigntemplatepublicGetObjectV2Response instantiates a new EzsigntemplatepublicGetObjectV2Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigntemplatepublicGetObjectV2Response(mPayload EzsigntemplatepublicGetObjectV2ResponseMPayload, objDebugPayload CommonResponseObjDebugPayload) *EzsigntemplatepublicGetObjectV2Response {
	this := EzsigntemplatepublicGetObjectV2Response{}
	this.ObjDebugPayload = objDebugPayload
	this.MPayload = mPayload
	return &this
}

// NewEzsigntemplatepublicGetObjectV2ResponseWithDefaults instantiates a new EzsigntemplatepublicGetObjectV2Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigntemplatepublicGetObjectV2ResponseWithDefaults() *EzsigntemplatepublicGetObjectV2Response {
	this := EzsigntemplatepublicGetObjectV2Response{}
	return &this
}

// GetMPayload returns the MPayload field value
func (o *EzsigntemplatepublicGetObjectV2Response) GetMPayload() EzsigntemplatepublicGetObjectV2ResponseMPayload {
	if o == nil {
		var ret EzsigntemplatepublicGetObjectV2ResponseMPayload
		return ret
	}

	return o.MPayload
}

// GetMPayloadOk returns a tuple with the MPayload field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepublicGetObjectV2Response) GetMPayloadOk() (*EzsigntemplatepublicGetObjectV2ResponseMPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MPayload, true
}

// SetMPayload sets field value
func (o *EzsigntemplatepublicGetObjectV2Response) SetMPayload(v EzsigntemplatepublicGetObjectV2ResponseMPayload) {
	o.MPayload = v
}

func (o EzsigntemplatepublicGetObjectV2Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigntemplatepublicGetObjectV2Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mPayload"] = o.MPayload
	return toSerialize, nil
}

func (o *EzsigntemplatepublicGetObjectV2Response) UnmarshalJSON(data []byte) (err error) {
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

	varEzsigntemplatepublicGetObjectV2Response := _EzsigntemplatepublicGetObjectV2Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsigntemplatepublicGetObjectV2Response)

	if err != nil {
		return err
	}

	*o = EzsigntemplatepublicGetObjectV2Response(varEzsigntemplatepublicGetObjectV2Response)

	return err
}

type NullableEzsigntemplatepublicGetObjectV2Response struct {
	value *EzsigntemplatepublicGetObjectV2Response
	isSet bool
}

func (v NullableEzsigntemplatepublicGetObjectV2Response) Get() *EzsigntemplatepublicGetObjectV2Response {
	return v.value
}

func (v *NullableEzsigntemplatepublicGetObjectV2Response) Set(val *EzsigntemplatepublicGetObjectV2Response) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigntemplatepublicGetObjectV2Response) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigntemplatepublicGetObjectV2Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigntemplatepublicGetObjectV2Response(val *EzsigntemplatepublicGetObjectV2Response) *NullableEzsigntemplatepublicGetObjectV2Response {
	return &NullableEzsigntemplatepublicGetObjectV2Response{value: val, isSet: true}
}

func (v NullableEzsigntemplatepublicGetObjectV2Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigntemplatepublicGetObjectV2Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


