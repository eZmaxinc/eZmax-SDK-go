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

// checks if the EzsignbulksendCreateObjectV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignbulksendCreateObjectV1ResponseMPayload{}

// EzsignbulksendCreateObjectV1ResponseMPayload Payload for POST /1/object/ezsignbulksend
type EzsignbulksendCreateObjectV1ResponseMPayload struct {
	// An array of unique IDs representing the object that were requested to be created.  They are returned in the same order as the array containing the objects to be created that was sent in the request.
	APkiEzsignbulksendID []int32 `json:"a_pkiEzsignbulksendID"`
}

type _EzsignbulksendCreateObjectV1ResponseMPayload EzsignbulksendCreateObjectV1ResponseMPayload

// NewEzsignbulksendCreateObjectV1ResponseMPayload instantiates a new EzsignbulksendCreateObjectV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignbulksendCreateObjectV1ResponseMPayload(aPkiEzsignbulksendID []int32) *EzsignbulksendCreateObjectV1ResponseMPayload {
	this := EzsignbulksendCreateObjectV1ResponseMPayload{}
	this.APkiEzsignbulksendID = aPkiEzsignbulksendID
	return &this
}

// NewEzsignbulksendCreateObjectV1ResponseMPayloadWithDefaults instantiates a new EzsignbulksendCreateObjectV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignbulksendCreateObjectV1ResponseMPayloadWithDefaults() *EzsignbulksendCreateObjectV1ResponseMPayload {
	this := EzsignbulksendCreateObjectV1ResponseMPayload{}
	return &this
}

// GetAPkiEzsignbulksendID returns the APkiEzsignbulksendID field value
func (o *EzsignbulksendCreateObjectV1ResponseMPayload) GetAPkiEzsignbulksendID() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.APkiEzsignbulksendID
}

// GetAPkiEzsignbulksendIDOk returns a tuple with the APkiEzsignbulksendID field value
// and a boolean to check if the value has been set.
func (o *EzsignbulksendCreateObjectV1ResponseMPayload) GetAPkiEzsignbulksendIDOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.APkiEzsignbulksendID, true
}

// SetAPkiEzsignbulksendID sets field value
func (o *EzsignbulksendCreateObjectV1ResponseMPayload) SetAPkiEzsignbulksendID(v []int32) {
	o.APkiEzsignbulksendID = v
}

func (o EzsignbulksendCreateObjectV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignbulksendCreateObjectV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_pkiEzsignbulksendID"] = o.APkiEzsignbulksendID
	return toSerialize, nil
}

func (o *EzsignbulksendCreateObjectV1ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"a_pkiEzsignbulksendID",
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

	varEzsignbulksendCreateObjectV1ResponseMPayload := _EzsignbulksendCreateObjectV1ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsignbulksendCreateObjectV1ResponseMPayload)

	if err != nil {
		return err
	}

	*o = EzsignbulksendCreateObjectV1ResponseMPayload(varEzsignbulksendCreateObjectV1ResponseMPayload)

	return err
}

type NullableEzsignbulksendCreateObjectV1ResponseMPayload struct {
	value *EzsignbulksendCreateObjectV1ResponseMPayload
	isSet bool
}

func (v NullableEzsignbulksendCreateObjectV1ResponseMPayload) Get() *EzsignbulksendCreateObjectV1ResponseMPayload {
	return v.value
}

func (v *NullableEzsignbulksendCreateObjectV1ResponseMPayload) Set(val *EzsignbulksendCreateObjectV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignbulksendCreateObjectV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignbulksendCreateObjectV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignbulksendCreateObjectV1ResponseMPayload(val *EzsignbulksendCreateObjectV1ResponseMPayload) *NullableEzsignbulksendCreateObjectV1ResponseMPayload {
	return &NullableEzsignbulksendCreateObjectV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableEzsignbulksendCreateObjectV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignbulksendCreateObjectV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


