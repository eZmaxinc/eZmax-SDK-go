/*
eZmax API Definition (Full)

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.2.2
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the EzsigndocumentGetActionableElementsV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigndocumentGetActionableElementsV1ResponseMPayload{}

// EzsigndocumentGetActionableElementsV1ResponseMPayload Payload for GET /1/object/ezsigndocument/{pkiEzsigndocumentID}/getActionableElements
type EzsigndocumentGetActionableElementsV1ResponseMPayload struct {
	AObjEzsignsignature []EzsignsignatureResponseCompound `json:"a_objEzsignsignature"`
	AObjEzsignformfieldgroup []EzsignformfieldgroupResponseCompound `json:"a_objEzsignformfieldgroup"`
}

type _EzsigndocumentGetActionableElementsV1ResponseMPayload EzsigndocumentGetActionableElementsV1ResponseMPayload

// NewEzsigndocumentGetActionableElementsV1ResponseMPayload instantiates a new EzsigndocumentGetActionableElementsV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigndocumentGetActionableElementsV1ResponseMPayload(aObjEzsignsignature []EzsignsignatureResponseCompound, aObjEzsignformfieldgroup []EzsignformfieldgroupResponseCompound) *EzsigndocumentGetActionableElementsV1ResponseMPayload {
	this := EzsigndocumentGetActionableElementsV1ResponseMPayload{}
	this.AObjEzsignsignature = aObjEzsignsignature
	this.AObjEzsignformfieldgroup = aObjEzsignformfieldgroup
	return &this
}

// NewEzsigndocumentGetActionableElementsV1ResponseMPayloadWithDefaults instantiates a new EzsigndocumentGetActionableElementsV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigndocumentGetActionableElementsV1ResponseMPayloadWithDefaults() *EzsigndocumentGetActionableElementsV1ResponseMPayload {
	this := EzsigndocumentGetActionableElementsV1ResponseMPayload{}
	return &this
}

// GetAObjEzsignsignature returns the AObjEzsignsignature field value
func (o *EzsigndocumentGetActionableElementsV1ResponseMPayload) GetAObjEzsignsignature() []EzsignsignatureResponseCompound {
	if o == nil {
		var ret []EzsignsignatureResponseCompound
		return ret
	}

	return o.AObjEzsignsignature
}

// GetAObjEzsignsignatureOk returns a tuple with the AObjEzsignsignature field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentGetActionableElementsV1ResponseMPayload) GetAObjEzsignsignatureOk() ([]EzsignsignatureResponseCompound, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjEzsignsignature, true
}

// SetAObjEzsignsignature sets field value
func (o *EzsigndocumentGetActionableElementsV1ResponseMPayload) SetAObjEzsignsignature(v []EzsignsignatureResponseCompound) {
	o.AObjEzsignsignature = v
}

// GetAObjEzsignformfieldgroup returns the AObjEzsignformfieldgroup field value
func (o *EzsigndocumentGetActionableElementsV1ResponseMPayload) GetAObjEzsignformfieldgroup() []EzsignformfieldgroupResponseCompound {
	if o == nil {
		var ret []EzsignformfieldgroupResponseCompound
		return ret
	}

	return o.AObjEzsignformfieldgroup
}

// GetAObjEzsignformfieldgroupOk returns a tuple with the AObjEzsignformfieldgroup field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentGetActionableElementsV1ResponseMPayload) GetAObjEzsignformfieldgroupOk() ([]EzsignformfieldgroupResponseCompound, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjEzsignformfieldgroup, true
}

// SetAObjEzsignformfieldgroup sets field value
func (o *EzsigndocumentGetActionableElementsV1ResponseMPayload) SetAObjEzsignformfieldgroup(v []EzsignformfieldgroupResponseCompound) {
	o.AObjEzsignformfieldgroup = v
}

func (o EzsigndocumentGetActionableElementsV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigndocumentGetActionableElementsV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_objEzsignsignature"] = o.AObjEzsignsignature
	toSerialize["a_objEzsignformfieldgroup"] = o.AObjEzsignformfieldgroup
	return toSerialize, nil
}

func (o *EzsigndocumentGetActionableElementsV1ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"a_objEzsignsignature",
		"a_objEzsignformfieldgroup",
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

	varEzsigndocumentGetActionableElementsV1ResponseMPayload := _EzsigndocumentGetActionableElementsV1ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsigndocumentGetActionableElementsV1ResponseMPayload)

	if err != nil {
		return err
	}

	*o = EzsigndocumentGetActionableElementsV1ResponseMPayload(varEzsigndocumentGetActionableElementsV1ResponseMPayload)

	return err
}

type NullableEzsigndocumentGetActionableElementsV1ResponseMPayload struct {
	value *EzsigndocumentGetActionableElementsV1ResponseMPayload
	isSet bool
}

func (v NullableEzsigndocumentGetActionableElementsV1ResponseMPayload) Get() *EzsigndocumentGetActionableElementsV1ResponseMPayload {
	return v.value
}

func (v *NullableEzsigndocumentGetActionableElementsV1ResponseMPayload) Set(val *EzsigndocumentGetActionableElementsV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigndocumentGetActionableElementsV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigndocumentGetActionableElementsV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigndocumentGetActionableElementsV1ResponseMPayload(val *EzsigndocumentGetActionableElementsV1ResponseMPayload) *NullableEzsigndocumentGetActionableElementsV1ResponseMPayload {
	return &NullableEzsigndocumentGetActionableElementsV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableEzsigndocumentGetActionableElementsV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigndocumentGetActionableElementsV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


