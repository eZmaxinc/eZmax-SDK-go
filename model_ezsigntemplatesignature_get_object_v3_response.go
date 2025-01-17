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

// checks if the EzsigntemplatesignatureGetObjectV3Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigntemplatesignatureGetObjectV3Response{}

// EzsigntemplatesignatureGetObjectV3Response Response for GET /3/object/ezsigntemplatesignature/{pkiEzsigntemplatesignatureID}
type EzsigntemplatesignatureGetObjectV3Response struct {
	CommonResponse
	MPayload EzsigntemplatesignatureGetObjectV3ResponseMPayload `json:"mPayload"`
}

type _EzsigntemplatesignatureGetObjectV3Response EzsigntemplatesignatureGetObjectV3Response

// NewEzsigntemplatesignatureGetObjectV3Response instantiates a new EzsigntemplatesignatureGetObjectV3Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigntemplatesignatureGetObjectV3Response(mPayload EzsigntemplatesignatureGetObjectV3ResponseMPayload, objDebugPayload CommonResponseObjDebugPayload) *EzsigntemplatesignatureGetObjectV3Response {
	this := EzsigntemplatesignatureGetObjectV3Response{}
	this.ObjDebugPayload = objDebugPayload
	this.MPayload = mPayload
	return &this
}

// NewEzsigntemplatesignatureGetObjectV3ResponseWithDefaults instantiates a new EzsigntemplatesignatureGetObjectV3Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigntemplatesignatureGetObjectV3ResponseWithDefaults() *EzsigntemplatesignatureGetObjectV3Response {
	this := EzsigntemplatesignatureGetObjectV3Response{}
	return &this
}

// GetMPayload returns the MPayload field value
func (o *EzsigntemplatesignatureGetObjectV3Response) GetMPayload() EzsigntemplatesignatureGetObjectV3ResponseMPayload {
	if o == nil {
		var ret EzsigntemplatesignatureGetObjectV3ResponseMPayload
		return ret
	}

	return o.MPayload
}

// GetMPayloadOk returns a tuple with the MPayload field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatesignatureGetObjectV3Response) GetMPayloadOk() (*EzsigntemplatesignatureGetObjectV3ResponseMPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MPayload, true
}

// SetMPayload sets field value
func (o *EzsigntemplatesignatureGetObjectV3Response) SetMPayload(v EzsigntemplatesignatureGetObjectV3ResponseMPayload) {
	o.MPayload = v
}

func (o EzsigntemplatesignatureGetObjectV3Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigntemplatesignatureGetObjectV3Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mPayload"] = o.MPayload
	return toSerialize, nil
}

func (o *EzsigntemplatesignatureGetObjectV3Response) UnmarshalJSON(data []byte) (err error) {
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

	varEzsigntemplatesignatureGetObjectV3Response := _EzsigntemplatesignatureGetObjectV3Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsigntemplatesignatureGetObjectV3Response)

	if err != nil {
		return err
	}

	*o = EzsigntemplatesignatureGetObjectV3Response(varEzsigntemplatesignatureGetObjectV3Response)

	return err
}

type NullableEzsigntemplatesignatureGetObjectV3Response struct {
	value *EzsigntemplatesignatureGetObjectV3Response
	isSet bool
}

func (v NullableEzsigntemplatesignatureGetObjectV3Response) Get() *EzsigntemplatesignatureGetObjectV3Response {
	return v.value
}

func (v *NullableEzsigntemplatesignatureGetObjectV3Response) Set(val *EzsigntemplatesignatureGetObjectV3Response) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigntemplatesignatureGetObjectV3Response) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigntemplatesignatureGetObjectV3Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigntemplatesignatureGetObjectV3Response(val *EzsigntemplatesignatureGetObjectV3Response) *NullableEzsigntemplatesignatureGetObjectV3Response {
	return &NullableEzsigntemplatesignatureGetObjectV3Response{value: val, isSet: true}
}

func (v NullableEzsigntemplatesignatureGetObjectV3Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigntemplatesignatureGetObjectV3Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


