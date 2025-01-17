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

// checks if the EzsigndocumentCreateEzsignelementsPositionedByWordV1Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigndocumentCreateEzsignelementsPositionedByWordV1Response{}

// EzsigndocumentCreateEzsignelementsPositionedByWordV1Response Response for POST /1/object/ezsigndocument/{pkiEzsigndocumentID}/createEzsignelementsPositionedByWord
type EzsigndocumentCreateEzsignelementsPositionedByWordV1Response struct {
	CommonResponse
	MPayload EzsigndocumentCreateEzsignelementsPositionedByWordV1ResponseMPayload `json:"mPayload"`
}

type _EzsigndocumentCreateEzsignelementsPositionedByWordV1Response EzsigndocumentCreateEzsignelementsPositionedByWordV1Response

// NewEzsigndocumentCreateEzsignelementsPositionedByWordV1Response instantiates a new EzsigndocumentCreateEzsignelementsPositionedByWordV1Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigndocumentCreateEzsignelementsPositionedByWordV1Response(mPayload EzsigndocumentCreateEzsignelementsPositionedByWordV1ResponseMPayload, objDebugPayload CommonResponseObjDebugPayload) *EzsigndocumentCreateEzsignelementsPositionedByWordV1Response {
	this := EzsigndocumentCreateEzsignelementsPositionedByWordV1Response{}
	this.ObjDebugPayload = objDebugPayload
	this.MPayload = mPayload
	return &this
}

// NewEzsigndocumentCreateEzsignelementsPositionedByWordV1ResponseWithDefaults instantiates a new EzsigndocumentCreateEzsignelementsPositionedByWordV1Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigndocumentCreateEzsignelementsPositionedByWordV1ResponseWithDefaults() *EzsigndocumentCreateEzsignelementsPositionedByWordV1Response {
	this := EzsigndocumentCreateEzsignelementsPositionedByWordV1Response{}
	return &this
}

// GetMPayload returns the MPayload field value
func (o *EzsigndocumentCreateEzsignelementsPositionedByWordV1Response) GetMPayload() EzsigndocumentCreateEzsignelementsPositionedByWordV1ResponseMPayload {
	if o == nil {
		var ret EzsigndocumentCreateEzsignelementsPositionedByWordV1ResponseMPayload
		return ret
	}

	return o.MPayload
}

// GetMPayloadOk returns a tuple with the MPayload field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentCreateEzsignelementsPositionedByWordV1Response) GetMPayloadOk() (*EzsigndocumentCreateEzsignelementsPositionedByWordV1ResponseMPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MPayload, true
}

// SetMPayload sets field value
func (o *EzsigndocumentCreateEzsignelementsPositionedByWordV1Response) SetMPayload(v EzsigndocumentCreateEzsignelementsPositionedByWordV1ResponseMPayload) {
	o.MPayload = v
}

func (o EzsigndocumentCreateEzsignelementsPositionedByWordV1Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigndocumentCreateEzsignelementsPositionedByWordV1Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mPayload"] = o.MPayload
	return toSerialize, nil
}

func (o *EzsigndocumentCreateEzsignelementsPositionedByWordV1Response) UnmarshalJSON(data []byte) (err error) {
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

	varEzsigndocumentCreateEzsignelementsPositionedByWordV1Response := _EzsigndocumentCreateEzsignelementsPositionedByWordV1Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsigndocumentCreateEzsignelementsPositionedByWordV1Response)

	if err != nil {
		return err
	}

	*o = EzsigndocumentCreateEzsignelementsPositionedByWordV1Response(varEzsigndocumentCreateEzsignelementsPositionedByWordV1Response)

	return err
}

type NullableEzsigndocumentCreateEzsignelementsPositionedByWordV1Response struct {
	value *EzsigndocumentCreateEzsignelementsPositionedByWordV1Response
	isSet bool
}

func (v NullableEzsigndocumentCreateEzsignelementsPositionedByWordV1Response) Get() *EzsigndocumentCreateEzsignelementsPositionedByWordV1Response {
	return v.value
}

func (v *NullableEzsigndocumentCreateEzsignelementsPositionedByWordV1Response) Set(val *EzsigndocumentCreateEzsignelementsPositionedByWordV1Response) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigndocumentCreateEzsignelementsPositionedByWordV1Response) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigndocumentCreateEzsignelementsPositionedByWordV1Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigndocumentCreateEzsignelementsPositionedByWordV1Response(val *EzsigndocumentCreateEzsignelementsPositionedByWordV1Response) *NullableEzsigndocumentCreateEzsignelementsPositionedByWordV1Response {
	return &NullableEzsigndocumentCreateEzsignelementsPositionedByWordV1Response{value: val, isSet: true}
}

func (v NullableEzsigndocumentCreateEzsignelementsPositionedByWordV1Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigndocumentCreateEzsignelementsPositionedByWordV1Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


