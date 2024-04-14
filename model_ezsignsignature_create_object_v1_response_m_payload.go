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

// checks if the EzsignsignatureCreateObjectV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignsignatureCreateObjectV1ResponseMPayload{}

// EzsignsignatureCreateObjectV1ResponseMPayload Payload for POST /1/object/ezsignsignature
type EzsignsignatureCreateObjectV1ResponseMPayload struct {
	// An array of unique IDs representing the object that were requested to be created.  They are returned in the same order as the array containing the objects to be created that was sent in the request.
	APkiEzsignsignatureID []int32 `json:"a_pkiEzsignsignatureID"`
}

type _EzsignsignatureCreateObjectV1ResponseMPayload EzsignsignatureCreateObjectV1ResponseMPayload

// NewEzsignsignatureCreateObjectV1ResponseMPayload instantiates a new EzsignsignatureCreateObjectV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignsignatureCreateObjectV1ResponseMPayload(aPkiEzsignsignatureID []int32) *EzsignsignatureCreateObjectV1ResponseMPayload {
	this := EzsignsignatureCreateObjectV1ResponseMPayload{}
	this.APkiEzsignsignatureID = aPkiEzsignsignatureID
	return &this
}

// NewEzsignsignatureCreateObjectV1ResponseMPayloadWithDefaults instantiates a new EzsignsignatureCreateObjectV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignsignatureCreateObjectV1ResponseMPayloadWithDefaults() *EzsignsignatureCreateObjectV1ResponseMPayload {
	this := EzsignsignatureCreateObjectV1ResponseMPayload{}
	return &this
}

// GetAPkiEzsignsignatureID returns the APkiEzsignsignatureID field value
func (o *EzsignsignatureCreateObjectV1ResponseMPayload) GetAPkiEzsignsignatureID() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.APkiEzsignsignatureID
}

// GetAPkiEzsignsignatureIDOk returns a tuple with the APkiEzsignsignatureID field value
// and a boolean to check if the value has been set.
func (o *EzsignsignatureCreateObjectV1ResponseMPayload) GetAPkiEzsignsignatureIDOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.APkiEzsignsignatureID, true
}

// SetAPkiEzsignsignatureID sets field value
func (o *EzsignsignatureCreateObjectV1ResponseMPayload) SetAPkiEzsignsignatureID(v []int32) {
	o.APkiEzsignsignatureID = v
}

func (o EzsignsignatureCreateObjectV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignsignatureCreateObjectV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_pkiEzsignsignatureID"] = o.APkiEzsignsignatureID
	return toSerialize, nil
}

func (o *EzsignsignatureCreateObjectV1ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"a_pkiEzsignsignatureID",
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

	varEzsignsignatureCreateObjectV1ResponseMPayload := _EzsignsignatureCreateObjectV1ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsignsignatureCreateObjectV1ResponseMPayload)

	if err != nil {
		return err
	}

	*o = EzsignsignatureCreateObjectV1ResponseMPayload(varEzsignsignatureCreateObjectV1ResponseMPayload)

	return err
}

type NullableEzsignsignatureCreateObjectV1ResponseMPayload struct {
	value *EzsignsignatureCreateObjectV1ResponseMPayload
	isSet bool
}

func (v NullableEzsignsignatureCreateObjectV1ResponseMPayload) Get() *EzsignsignatureCreateObjectV1ResponseMPayload {
	return v.value
}

func (v *NullableEzsignsignatureCreateObjectV1ResponseMPayload) Set(val *EzsignsignatureCreateObjectV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignsignatureCreateObjectV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignsignatureCreateObjectV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignsignatureCreateObjectV1ResponseMPayload(val *EzsignsignatureCreateObjectV1ResponseMPayload) *NullableEzsignsignatureCreateObjectV1ResponseMPayload {
	return &NullableEzsignsignatureCreateObjectV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableEzsignsignatureCreateObjectV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignsignatureCreateObjectV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


