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

// checks if the EzsignannotationGetObjectV2ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignannotationGetObjectV2ResponseMPayload{}

// EzsignannotationGetObjectV2ResponseMPayload Payload for GET /2/object/ezsignannotation/{pkiEzsignannotationID}
type EzsignannotationGetObjectV2ResponseMPayload struct {
	ObjEzsignannotation EzsignannotationResponseCompound `json:"objEzsignannotation"`
}

type _EzsignannotationGetObjectV2ResponseMPayload EzsignannotationGetObjectV2ResponseMPayload

// NewEzsignannotationGetObjectV2ResponseMPayload instantiates a new EzsignannotationGetObjectV2ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignannotationGetObjectV2ResponseMPayload(objEzsignannotation EzsignannotationResponseCompound) *EzsignannotationGetObjectV2ResponseMPayload {
	this := EzsignannotationGetObjectV2ResponseMPayload{}
	this.ObjEzsignannotation = objEzsignannotation
	return &this
}

// NewEzsignannotationGetObjectV2ResponseMPayloadWithDefaults instantiates a new EzsignannotationGetObjectV2ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignannotationGetObjectV2ResponseMPayloadWithDefaults() *EzsignannotationGetObjectV2ResponseMPayload {
	this := EzsignannotationGetObjectV2ResponseMPayload{}
	return &this
}

// GetObjEzsignannotation returns the ObjEzsignannotation field value
func (o *EzsignannotationGetObjectV2ResponseMPayload) GetObjEzsignannotation() EzsignannotationResponseCompound {
	if o == nil {
		var ret EzsignannotationResponseCompound
		return ret
	}

	return o.ObjEzsignannotation
}

// GetObjEzsignannotationOk returns a tuple with the ObjEzsignannotation field value
// and a boolean to check if the value has been set.
func (o *EzsignannotationGetObjectV2ResponseMPayload) GetObjEzsignannotationOk() (*EzsignannotationResponseCompound, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjEzsignannotation, true
}

// SetObjEzsignannotation sets field value
func (o *EzsignannotationGetObjectV2ResponseMPayload) SetObjEzsignannotation(v EzsignannotationResponseCompound) {
	o.ObjEzsignannotation = v
}

func (o EzsignannotationGetObjectV2ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignannotationGetObjectV2ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objEzsignannotation"] = o.ObjEzsignannotation
	return toSerialize, nil
}

func (o *EzsignannotationGetObjectV2ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"objEzsignannotation",
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

	varEzsignannotationGetObjectV2ResponseMPayload := _EzsignannotationGetObjectV2ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsignannotationGetObjectV2ResponseMPayload)

	if err != nil {
		return err
	}

	*o = EzsignannotationGetObjectV2ResponseMPayload(varEzsignannotationGetObjectV2ResponseMPayload)

	return err
}

type NullableEzsignannotationGetObjectV2ResponseMPayload struct {
	value *EzsignannotationGetObjectV2ResponseMPayload
	isSet bool
}

func (v NullableEzsignannotationGetObjectV2ResponseMPayload) Get() *EzsignannotationGetObjectV2ResponseMPayload {
	return v.value
}

func (v *NullableEzsignannotationGetObjectV2ResponseMPayload) Set(val *EzsignannotationGetObjectV2ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignannotationGetObjectV2ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignannotationGetObjectV2ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignannotationGetObjectV2ResponseMPayload(val *EzsignannotationGetObjectV2ResponseMPayload) *NullableEzsignannotationGetObjectV2ResponseMPayload {
	return &NullableEzsignannotationGetObjectV2ResponseMPayload{value: val, isSet: true}
}

func (v NullableEzsignannotationGetObjectV2ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignannotationGetObjectV2ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


