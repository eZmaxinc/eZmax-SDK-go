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

// checks if the SupplyCreateObjectV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SupplyCreateObjectV1ResponseMPayload{}

// SupplyCreateObjectV1ResponseMPayload Payload for POST /1/object/supply
type SupplyCreateObjectV1ResponseMPayload struct {
	// An array of unique IDs representing the object that were requested to be created.  They are returned in the same order as the array containing the objects to be created that was sent in the request.
	APkiSupplyID []int32 `json:"a_pkiSupplyID"`
}

type _SupplyCreateObjectV1ResponseMPayload SupplyCreateObjectV1ResponseMPayload

// NewSupplyCreateObjectV1ResponseMPayload instantiates a new SupplyCreateObjectV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSupplyCreateObjectV1ResponseMPayload(aPkiSupplyID []int32) *SupplyCreateObjectV1ResponseMPayload {
	this := SupplyCreateObjectV1ResponseMPayload{}
	this.APkiSupplyID = aPkiSupplyID
	return &this
}

// NewSupplyCreateObjectV1ResponseMPayloadWithDefaults instantiates a new SupplyCreateObjectV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSupplyCreateObjectV1ResponseMPayloadWithDefaults() *SupplyCreateObjectV1ResponseMPayload {
	this := SupplyCreateObjectV1ResponseMPayload{}
	return &this
}

// GetAPkiSupplyID returns the APkiSupplyID field value
func (o *SupplyCreateObjectV1ResponseMPayload) GetAPkiSupplyID() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.APkiSupplyID
}

// GetAPkiSupplyIDOk returns a tuple with the APkiSupplyID field value
// and a boolean to check if the value has been set.
func (o *SupplyCreateObjectV1ResponseMPayload) GetAPkiSupplyIDOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.APkiSupplyID, true
}

// SetAPkiSupplyID sets field value
func (o *SupplyCreateObjectV1ResponseMPayload) SetAPkiSupplyID(v []int32) {
	o.APkiSupplyID = v
}

func (o SupplyCreateObjectV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SupplyCreateObjectV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_pkiSupplyID"] = o.APkiSupplyID
	return toSerialize, nil
}

func (o *SupplyCreateObjectV1ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"a_pkiSupplyID",
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

	varSupplyCreateObjectV1ResponseMPayload := _SupplyCreateObjectV1ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSupplyCreateObjectV1ResponseMPayload)

	if err != nil {
		return err
	}

	*o = SupplyCreateObjectV1ResponseMPayload(varSupplyCreateObjectV1ResponseMPayload)

	return err
}

type NullableSupplyCreateObjectV1ResponseMPayload struct {
	value *SupplyCreateObjectV1ResponseMPayload
	isSet bool
}

func (v NullableSupplyCreateObjectV1ResponseMPayload) Get() *SupplyCreateObjectV1ResponseMPayload {
	return v.value
}

func (v *NullableSupplyCreateObjectV1ResponseMPayload) Set(val *SupplyCreateObjectV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableSupplyCreateObjectV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableSupplyCreateObjectV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSupplyCreateObjectV1ResponseMPayload(val *SupplyCreateObjectV1ResponseMPayload) *NullableSupplyCreateObjectV1ResponseMPayload {
	return &NullableSupplyCreateObjectV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableSupplyCreateObjectV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSupplyCreateObjectV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


