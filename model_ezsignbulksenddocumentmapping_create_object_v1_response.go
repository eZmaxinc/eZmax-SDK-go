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

// checks if the EzsignbulksenddocumentmappingCreateObjectV1Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignbulksenddocumentmappingCreateObjectV1Response{}

// EzsignbulksenddocumentmappingCreateObjectV1Response Response for POST /1/object/ezsignbulksenddocumentmapping
type EzsignbulksenddocumentmappingCreateObjectV1Response struct {
	CommonResponse
	MPayload EzsignbulksenddocumentmappingCreateObjectV1ResponseMPayload `json:"mPayload"`
}

type _EzsignbulksenddocumentmappingCreateObjectV1Response EzsignbulksenddocumentmappingCreateObjectV1Response

// NewEzsignbulksenddocumentmappingCreateObjectV1Response instantiates a new EzsignbulksenddocumentmappingCreateObjectV1Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignbulksenddocumentmappingCreateObjectV1Response(mPayload EzsignbulksenddocumentmappingCreateObjectV1ResponseMPayload, objDebugPayload CommonResponseObjDebugPayload) *EzsignbulksenddocumentmappingCreateObjectV1Response {
	this := EzsignbulksenddocumentmappingCreateObjectV1Response{}
	this.ObjDebugPayload = objDebugPayload
	this.MPayload = mPayload
	return &this
}

// NewEzsignbulksenddocumentmappingCreateObjectV1ResponseWithDefaults instantiates a new EzsignbulksenddocumentmappingCreateObjectV1Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignbulksenddocumentmappingCreateObjectV1ResponseWithDefaults() *EzsignbulksenddocumentmappingCreateObjectV1Response {
	this := EzsignbulksenddocumentmappingCreateObjectV1Response{}
	return &this
}

// GetMPayload returns the MPayload field value
func (o *EzsignbulksenddocumentmappingCreateObjectV1Response) GetMPayload() EzsignbulksenddocumentmappingCreateObjectV1ResponseMPayload {
	if o == nil {
		var ret EzsignbulksenddocumentmappingCreateObjectV1ResponseMPayload
		return ret
	}

	return o.MPayload
}

// GetMPayloadOk returns a tuple with the MPayload field value
// and a boolean to check if the value has been set.
func (o *EzsignbulksenddocumentmappingCreateObjectV1Response) GetMPayloadOk() (*EzsignbulksenddocumentmappingCreateObjectV1ResponseMPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MPayload, true
}

// SetMPayload sets field value
func (o *EzsignbulksenddocumentmappingCreateObjectV1Response) SetMPayload(v EzsignbulksenddocumentmappingCreateObjectV1ResponseMPayload) {
	o.MPayload = v
}

func (o EzsignbulksenddocumentmappingCreateObjectV1Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignbulksenddocumentmappingCreateObjectV1Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mPayload"] = o.MPayload
	return toSerialize, nil
}

func (o *EzsignbulksenddocumentmappingCreateObjectV1Response) UnmarshalJSON(data []byte) (err error) {
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

	varEzsignbulksenddocumentmappingCreateObjectV1Response := _EzsignbulksenddocumentmappingCreateObjectV1Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsignbulksenddocumentmappingCreateObjectV1Response)

	if err != nil {
		return err
	}

	*o = EzsignbulksenddocumentmappingCreateObjectV1Response(varEzsignbulksenddocumentmappingCreateObjectV1Response)

	return err
}

type NullableEzsignbulksenddocumentmappingCreateObjectV1Response struct {
	value *EzsignbulksenddocumentmappingCreateObjectV1Response
	isSet bool
}

func (v NullableEzsignbulksenddocumentmappingCreateObjectV1Response) Get() *EzsignbulksenddocumentmappingCreateObjectV1Response {
	return v.value
}

func (v *NullableEzsignbulksenddocumentmappingCreateObjectV1Response) Set(val *EzsignbulksenddocumentmappingCreateObjectV1Response) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignbulksenddocumentmappingCreateObjectV1Response) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignbulksenddocumentmappingCreateObjectV1Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignbulksenddocumentmappingCreateObjectV1Response(val *EzsignbulksenddocumentmappingCreateObjectV1Response) *NullableEzsignbulksenddocumentmappingCreateObjectV1Response {
	return &NullableEzsignbulksenddocumentmappingCreateObjectV1Response{value: val, isSet: true}
}

func (v NullableEzsignbulksenddocumentmappingCreateObjectV1Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignbulksenddocumentmappingCreateObjectV1Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


