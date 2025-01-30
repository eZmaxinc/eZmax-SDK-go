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

// checks if the EzsignannotationCreateObjectV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignannotationCreateObjectV1ResponseMPayload{}

// EzsignannotationCreateObjectV1ResponseMPayload Payload for POST /1/object/ezsignannotation
type EzsignannotationCreateObjectV1ResponseMPayload struct {
	// An array of unique IDs representing the object that were requested to be created.  They are returned in the same order as the array containing the objects to be created that was sent in the request.
	APkiEzsignannotationID []int32 `json:"a_pkiEzsignannotationID"`
}

type _EzsignannotationCreateObjectV1ResponseMPayload EzsignannotationCreateObjectV1ResponseMPayload

// NewEzsignannotationCreateObjectV1ResponseMPayload instantiates a new EzsignannotationCreateObjectV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignannotationCreateObjectV1ResponseMPayload(aPkiEzsignannotationID []int32) *EzsignannotationCreateObjectV1ResponseMPayload {
	this := EzsignannotationCreateObjectV1ResponseMPayload{}
	this.APkiEzsignannotationID = aPkiEzsignannotationID
	return &this
}

// NewEzsignannotationCreateObjectV1ResponseMPayloadWithDefaults instantiates a new EzsignannotationCreateObjectV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignannotationCreateObjectV1ResponseMPayloadWithDefaults() *EzsignannotationCreateObjectV1ResponseMPayload {
	this := EzsignannotationCreateObjectV1ResponseMPayload{}
	return &this
}

// GetAPkiEzsignannotationID returns the APkiEzsignannotationID field value
func (o *EzsignannotationCreateObjectV1ResponseMPayload) GetAPkiEzsignannotationID() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.APkiEzsignannotationID
}

// GetAPkiEzsignannotationIDOk returns a tuple with the APkiEzsignannotationID field value
// and a boolean to check if the value has been set.
func (o *EzsignannotationCreateObjectV1ResponseMPayload) GetAPkiEzsignannotationIDOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.APkiEzsignannotationID, true
}

// SetAPkiEzsignannotationID sets field value
func (o *EzsignannotationCreateObjectV1ResponseMPayload) SetAPkiEzsignannotationID(v []int32) {
	o.APkiEzsignannotationID = v
}

func (o EzsignannotationCreateObjectV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignannotationCreateObjectV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_pkiEzsignannotationID"] = o.APkiEzsignannotationID
	return toSerialize, nil
}

func (o *EzsignannotationCreateObjectV1ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"a_pkiEzsignannotationID",
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

	varEzsignannotationCreateObjectV1ResponseMPayload := _EzsignannotationCreateObjectV1ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsignannotationCreateObjectV1ResponseMPayload)

	if err != nil {
		return err
	}

	*o = EzsignannotationCreateObjectV1ResponseMPayload(varEzsignannotationCreateObjectV1ResponseMPayload)

	return err
}

type NullableEzsignannotationCreateObjectV1ResponseMPayload struct {
	value *EzsignannotationCreateObjectV1ResponseMPayload
	isSet bool
}

func (v NullableEzsignannotationCreateObjectV1ResponseMPayload) Get() *EzsignannotationCreateObjectV1ResponseMPayload {
	return v.value
}

func (v *NullableEzsignannotationCreateObjectV1ResponseMPayload) Set(val *EzsignannotationCreateObjectV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignannotationCreateObjectV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignannotationCreateObjectV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignannotationCreateObjectV1ResponseMPayload(val *EzsignannotationCreateObjectV1ResponseMPayload) *NullableEzsignannotationCreateObjectV1ResponseMPayload {
	return &NullableEzsignannotationCreateObjectV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableEzsignannotationCreateObjectV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignannotationCreateObjectV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


