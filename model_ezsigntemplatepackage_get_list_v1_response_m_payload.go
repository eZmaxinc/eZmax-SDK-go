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

// checks if the EzsigntemplatepackageGetListV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigntemplatepackageGetListV1ResponseMPayload{}

// EzsigntemplatepackageGetListV1ResponseMPayload Payload for GET /1/object/ezsigntemplatepackage/getList
type EzsigntemplatepackageGetListV1ResponseMPayload struct {
	CommonGetListV1ResponseMPayload
	AObjEzsigntemplatepackage []EzsigntemplatepackageListElement `json:"a_objEzsigntemplatepackage"`
}

type _EzsigntemplatepackageGetListV1ResponseMPayload EzsigntemplatepackageGetListV1ResponseMPayload

// NewEzsigntemplatepackageGetListV1ResponseMPayload instantiates a new EzsigntemplatepackageGetListV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigntemplatepackageGetListV1ResponseMPayload(aObjEzsigntemplatepackage []EzsigntemplatepackageListElement, iRowReturned int32, iRowFiltered int32) *EzsigntemplatepackageGetListV1ResponseMPayload {
	this := EzsigntemplatepackageGetListV1ResponseMPayload{}
	this.IRowReturned = iRowReturned
	this.IRowFiltered = iRowFiltered
	this.AObjEzsigntemplatepackage = aObjEzsigntemplatepackage
	return &this
}

// NewEzsigntemplatepackageGetListV1ResponseMPayloadWithDefaults instantiates a new EzsigntemplatepackageGetListV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigntemplatepackageGetListV1ResponseMPayloadWithDefaults() *EzsigntemplatepackageGetListV1ResponseMPayload {
	this := EzsigntemplatepackageGetListV1ResponseMPayload{}
	return &this
}

// GetAObjEzsigntemplatepackage returns the AObjEzsigntemplatepackage field value
func (o *EzsigntemplatepackageGetListV1ResponseMPayload) GetAObjEzsigntemplatepackage() []EzsigntemplatepackageListElement {
	if o == nil {
		var ret []EzsigntemplatepackageListElement
		return ret
	}

	return o.AObjEzsigntemplatepackage
}

// GetAObjEzsigntemplatepackageOk returns a tuple with the AObjEzsigntemplatepackage field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepackageGetListV1ResponseMPayload) GetAObjEzsigntemplatepackageOk() ([]EzsigntemplatepackageListElement, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjEzsigntemplatepackage, true
}

// SetAObjEzsigntemplatepackage sets field value
func (o *EzsigntemplatepackageGetListV1ResponseMPayload) SetAObjEzsigntemplatepackage(v []EzsigntemplatepackageListElement) {
	o.AObjEzsigntemplatepackage = v
}

func (o EzsigntemplatepackageGetListV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigntemplatepackageGetListV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_objEzsigntemplatepackage"] = o.AObjEzsigntemplatepackage
	return toSerialize, nil
}

func (o *EzsigntemplatepackageGetListV1ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"a_objEzsigntemplatepackage",
		"iRowReturned",
		"iRowFiltered",
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

	varEzsigntemplatepackageGetListV1ResponseMPayload := _EzsigntemplatepackageGetListV1ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsigntemplatepackageGetListV1ResponseMPayload)

	if err != nil {
		return err
	}

	*o = EzsigntemplatepackageGetListV1ResponseMPayload(varEzsigntemplatepackageGetListV1ResponseMPayload)

	return err
}

type NullableEzsigntemplatepackageGetListV1ResponseMPayload struct {
	value *EzsigntemplatepackageGetListV1ResponseMPayload
	isSet bool
}

func (v NullableEzsigntemplatepackageGetListV1ResponseMPayload) Get() *EzsigntemplatepackageGetListV1ResponseMPayload {
	return v.value
}

func (v *NullableEzsigntemplatepackageGetListV1ResponseMPayload) Set(val *EzsigntemplatepackageGetListV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigntemplatepackageGetListV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigntemplatepackageGetListV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigntemplatepackageGetListV1ResponseMPayload(val *EzsigntemplatepackageGetListV1ResponseMPayload) *NullableEzsigntemplatepackageGetListV1ResponseMPayload {
	return &NullableEzsigntemplatepackageGetListV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableEzsigntemplatepackageGetListV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigntemplatepackageGetListV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


