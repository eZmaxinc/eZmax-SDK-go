/*
eZmax API Definition (Full)

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.2.0
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the EzsigntemplateCreateObjectV2ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigntemplateCreateObjectV2ResponseMPayload{}

// EzsigntemplateCreateObjectV2ResponseMPayload Payload for POST /2/object/ezsigntemplate
type EzsigntemplateCreateObjectV2ResponseMPayload struct {
	// An array of unique IDs representing the object that were requested to be created.  They are returned in the same order as the array containing the objects to be created that was sent in the request.
	APkiEzsigntemplateID []int32 `json:"a_pkiEzsigntemplateID"`
}

type _EzsigntemplateCreateObjectV2ResponseMPayload EzsigntemplateCreateObjectV2ResponseMPayload

// NewEzsigntemplateCreateObjectV2ResponseMPayload instantiates a new EzsigntemplateCreateObjectV2ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigntemplateCreateObjectV2ResponseMPayload(aPkiEzsigntemplateID []int32) *EzsigntemplateCreateObjectV2ResponseMPayload {
	this := EzsigntemplateCreateObjectV2ResponseMPayload{}
	this.APkiEzsigntemplateID = aPkiEzsigntemplateID
	return &this
}

// NewEzsigntemplateCreateObjectV2ResponseMPayloadWithDefaults instantiates a new EzsigntemplateCreateObjectV2ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigntemplateCreateObjectV2ResponseMPayloadWithDefaults() *EzsigntemplateCreateObjectV2ResponseMPayload {
	this := EzsigntemplateCreateObjectV2ResponseMPayload{}
	return &this
}

// GetAPkiEzsigntemplateID returns the APkiEzsigntemplateID field value
func (o *EzsigntemplateCreateObjectV2ResponseMPayload) GetAPkiEzsigntemplateID() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.APkiEzsigntemplateID
}

// GetAPkiEzsigntemplateIDOk returns a tuple with the APkiEzsigntemplateID field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateCreateObjectV2ResponseMPayload) GetAPkiEzsigntemplateIDOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.APkiEzsigntemplateID, true
}

// SetAPkiEzsigntemplateID sets field value
func (o *EzsigntemplateCreateObjectV2ResponseMPayload) SetAPkiEzsigntemplateID(v []int32) {
	o.APkiEzsigntemplateID = v
}

func (o EzsigntemplateCreateObjectV2ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigntemplateCreateObjectV2ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_pkiEzsigntemplateID"] = o.APkiEzsigntemplateID
	return toSerialize, nil
}

func (o *EzsigntemplateCreateObjectV2ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"a_pkiEzsigntemplateID",
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

	varEzsigntemplateCreateObjectV2ResponseMPayload := _EzsigntemplateCreateObjectV2ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsigntemplateCreateObjectV2ResponseMPayload)

	if err != nil {
		return err
	}

	*o = EzsigntemplateCreateObjectV2ResponseMPayload(varEzsigntemplateCreateObjectV2ResponseMPayload)

	return err
}

type NullableEzsigntemplateCreateObjectV2ResponseMPayload struct {
	value *EzsigntemplateCreateObjectV2ResponseMPayload
	isSet bool
}

func (v NullableEzsigntemplateCreateObjectV2ResponseMPayload) Get() *EzsigntemplateCreateObjectV2ResponseMPayload {
	return v.value
}

func (v *NullableEzsigntemplateCreateObjectV2ResponseMPayload) Set(val *EzsigntemplateCreateObjectV2ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigntemplateCreateObjectV2ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigntemplateCreateObjectV2ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigntemplateCreateObjectV2ResponseMPayload(val *EzsigntemplateCreateObjectV2ResponseMPayload) *NullableEzsigntemplateCreateObjectV2ResponseMPayload {
	return &NullableEzsigntemplateCreateObjectV2ResponseMPayload{value: val, isSet: true}
}

func (v NullableEzsigntemplateCreateObjectV2ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigntemplateCreateObjectV2ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


