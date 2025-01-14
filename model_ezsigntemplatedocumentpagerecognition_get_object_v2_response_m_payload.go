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

// checks if the EzsigntemplatedocumentpagerecognitionGetObjectV2ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigntemplatedocumentpagerecognitionGetObjectV2ResponseMPayload{}

// EzsigntemplatedocumentpagerecognitionGetObjectV2ResponseMPayload Payload for GET /2/object/ezsigntemplatedocumentpagerecognition/{pkiEzsigntemplatedocumentpagerecognitionID}
type EzsigntemplatedocumentpagerecognitionGetObjectV2ResponseMPayload struct {
	ObjEzsigntemplatedocumentpagerecognition EzsigntemplatedocumentpagerecognitionResponseCompound `json:"objEzsigntemplatedocumentpagerecognition"`
}

type _EzsigntemplatedocumentpagerecognitionGetObjectV2ResponseMPayload EzsigntemplatedocumentpagerecognitionGetObjectV2ResponseMPayload

// NewEzsigntemplatedocumentpagerecognitionGetObjectV2ResponseMPayload instantiates a new EzsigntemplatedocumentpagerecognitionGetObjectV2ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigntemplatedocumentpagerecognitionGetObjectV2ResponseMPayload(objEzsigntemplatedocumentpagerecognition EzsigntemplatedocumentpagerecognitionResponseCompound) *EzsigntemplatedocumentpagerecognitionGetObjectV2ResponseMPayload {
	this := EzsigntemplatedocumentpagerecognitionGetObjectV2ResponseMPayload{}
	this.ObjEzsigntemplatedocumentpagerecognition = objEzsigntemplatedocumentpagerecognition
	return &this
}

// NewEzsigntemplatedocumentpagerecognitionGetObjectV2ResponseMPayloadWithDefaults instantiates a new EzsigntemplatedocumentpagerecognitionGetObjectV2ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigntemplatedocumentpagerecognitionGetObjectV2ResponseMPayloadWithDefaults() *EzsigntemplatedocumentpagerecognitionGetObjectV2ResponseMPayload {
	this := EzsigntemplatedocumentpagerecognitionGetObjectV2ResponseMPayload{}
	return &this
}

// GetObjEzsigntemplatedocumentpagerecognition returns the ObjEzsigntemplatedocumentpagerecognition field value
func (o *EzsigntemplatedocumentpagerecognitionGetObjectV2ResponseMPayload) GetObjEzsigntemplatedocumentpagerecognition() EzsigntemplatedocumentpagerecognitionResponseCompound {
	if o == nil {
		var ret EzsigntemplatedocumentpagerecognitionResponseCompound
		return ret
	}

	return o.ObjEzsigntemplatedocumentpagerecognition
}

// GetObjEzsigntemplatedocumentpagerecognitionOk returns a tuple with the ObjEzsigntemplatedocumentpagerecognition field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatedocumentpagerecognitionGetObjectV2ResponseMPayload) GetObjEzsigntemplatedocumentpagerecognitionOk() (*EzsigntemplatedocumentpagerecognitionResponseCompound, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjEzsigntemplatedocumentpagerecognition, true
}

// SetObjEzsigntemplatedocumentpagerecognition sets field value
func (o *EzsigntemplatedocumentpagerecognitionGetObjectV2ResponseMPayload) SetObjEzsigntemplatedocumentpagerecognition(v EzsigntemplatedocumentpagerecognitionResponseCompound) {
	o.ObjEzsigntemplatedocumentpagerecognition = v
}

func (o EzsigntemplatedocumentpagerecognitionGetObjectV2ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigntemplatedocumentpagerecognitionGetObjectV2ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objEzsigntemplatedocumentpagerecognition"] = o.ObjEzsigntemplatedocumentpagerecognition
	return toSerialize, nil
}

func (o *EzsigntemplatedocumentpagerecognitionGetObjectV2ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"objEzsigntemplatedocumentpagerecognition",
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

	varEzsigntemplatedocumentpagerecognitionGetObjectV2ResponseMPayload := _EzsigntemplatedocumentpagerecognitionGetObjectV2ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsigntemplatedocumentpagerecognitionGetObjectV2ResponseMPayload)

	if err != nil {
		return err
	}

	*o = EzsigntemplatedocumentpagerecognitionGetObjectV2ResponseMPayload(varEzsigntemplatedocumentpagerecognitionGetObjectV2ResponseMPayload)

	return err
}

type NullableEzsigntemplatedocumentpagerecognitionGetObjectV2ResponseMPayload struct {
	value *EzsigntemplatedocumentpagerecognitionGetObjectV2ResponseMPayload
	isSet bool
}

func (v NullableEzsigntemplatedocumentpagerecognitionGetObjectV2ResponseMPayload) Get() *EzsigntemplatedocumentpagerecognitionGetObjectV2ResponseMPayload {
	return v.value
}

func (v *NullableEzsigntemplatedocumentpagerecognitionGetObjectV2ResponseMPayload) Set(val *EzsigntemplatedocumentpagerecognitionGetObjectV2ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigntemplatedocumentpagerecognitionGetObjectV2ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigntemplatedocumentpagerecognitionGetObjectV2ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigntemplatedocumentpagerecognitionGetObjectV2ResponseMPayload(val *EzsigntemplatedocumentpagerecognitionGetObjectV2ResponseMPayload) *NullableEzsigntemplatedocumentpagerecognitionGetObjectV2ResponseMPayload {
	return &NullableEzsigntemplatedocumentpagerecognitionGetObjectV2ResponseMPayload{value: val, isSet: true}
}

func (v NullableEzsigntemplatedocumentpagerecognitionGetObjectV2ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigntemplatedocumentpagerecognitionGetObjectV2ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


