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

// checks if the EzdoctemplatedocumentCreateObjectV1Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzdoctemplatedocumentCreateObjectV1Response{}

// EzdoctemplatedocumentCreateObjectV1Response Response for POST /1/object/ezdoctemplatedocument
type EzdoctemplatedocumentCreateObjectV1Response struct {
	CommonResponse
	MPayload EzdoctemplatedocumentCreateObjectV1ResponseMPayload `json:"mPayload"`
}

type _EzdoctemplatedocumentCreateObjectV1Response EzdoctemplatedocumentCreateObjectV1Response

// NewEzdoctemplatedocumentCreateObjectV1Response instantiates a new EzdoctemplatedocumentCreateObjectV1Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzdoctemplatedocumentCreateObjectV1Response(mPayload EzdoctemplatedocumentCreateObjectV1ResponseMPayload, objDebugPayload CommonResponseObjDebugPayload) *EzdoctemplatedocumentCreateObjectV1Response {
	this := EzdoctemplatedocumentCreateObjectV1Response{}
	this.ObjDebugPayload = objDebugPayload
	this.MPayload = mPayload
	return &this
}

// NewEzdoctemplatedocumentCreateObjectV1ResponseWithDefaults instantiates a new EzdoctemplatedocumentCreateObjectV1Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzdoctemplatedocumentCreateObjectV1ResponseWithDefaults() *EzdoctemplatedocumentCreateObjectV1Response {
	this := EzdoctemplatedocumentCreateObjectV1Response{}
	return &this
}

// GetMPayload returns the MPayload field value
func (o *EzdoctemplatedocumentCreateObjectV1Response) GetMPayload() EzdoctemplatedocumentCreateObjectV1ResponseMPayload {
	if o == nil {
		var ret EzdoctemplatedocumentCreateObjectV1ResponseMPayload
		return ret
	}

	return o.MPayload
}

// GetMPayloadOk returns a tuple with the MPayload field value
// and a boolean to check if the value has been set.
func (o *EzdoctemplatedocumentCreateObjectV1Response) GetMPayloadOk() (*EzdoctemplatedocumentCreateObjectV1ResponseMPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MPayload, true
}

// SetMPayload sets field value
func (o *EzdoctemplatedocumentCreateObjectV1Response) SetMPayload(v EzdoctemplatedocumentCreateObjectV1ResponseMPayload) {
	o.MPayload = v
}

func (o EzdoctemplatedocumentCreateObjectV1Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzdoctemplatedocumentCreateObjectV1Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mPayload"] = o.MPayload
	return toSerialize, nil
}

func (o *EzdoctemplatedocumentCreateObjectV1Response) UnmarshalJSON(data []byte) (err error) {
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

	varEzdoctemplatedocumentCreateObjectV1Response := _EzdoctemplatedocumentCreateObjectV1Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzdoctemplatedocumentCreateObjectV1Response)

	if err != nil {
		return err
	}

	*o = EzdoctemplatedocumentCreateObjectV1Response(varEzdoctemplatedocumentCreateObjectV1Response)

	return err
}

type NullableEzdoctemplatedocumentCreateObjectV1Response struct {
	value *EzdoctemplatedocumentCreateObjectV1Response
	isSet bool
}

func (v NullableEzdoctemplatedocumentCreateObjectV1Response) Get() *EzdoctemplatedocumentCreateObjectV1Response {
	return v.value
}

func (v *NullableEzdoctemplatedocumentCreateObjectV1Response) Set(val *EzdoctemplatedocumentCreateObjectV1Response) {
	v.value = val
	v.isSet = true
}

func (v NullableEzdoctemplatedocumentCreateObjectV1Response) IsSet() bool {
	return v.isSet
}

func (v *NullableEzdoctemplatedocumentCreateObjectV1Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzdoctemplatedocumentCreateObjectV1Response(val *EzdoctemplatedocumentCreateObjectV1Response) *NullableEzdoctemplatedocumentCreateObjectV1Response {
	return &NullableEzdoctemplatedocumentCreateObjectV1Response{value: val, isSet: true}
}

func (v NullableEzdoctemplatedocumentCreateObjectV1Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzdoctemplatedocumentCreateObjectV1Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


