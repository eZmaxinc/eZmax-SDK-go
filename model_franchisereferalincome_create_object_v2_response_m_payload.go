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

// checks if the FranchisereferalincomeCreateObjectV2ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FranchisereferalincomeCreateObjectV2ResponseMPayload{}

// FranchisereferalincomeCreateObjectV2ResponseMPayload Payload for POST /2/object/franchisereferalincome
type FranchisereferalincomeCreateObjectV2ResponseMPayload struct {
	// An array of unique IDs representing the object that were requested to be created.  They are returned in the same order as the array containing the objects to be created that was sent in the request.
	APkiFranchisereferalincomeID []int32 `json:"a_pkiFranchisereferalincomeID"`
}

type _FranchisereferalincomeCreateObjectV2ResponseMPayload FranchisereferalincomeCreateObjectV2ResponseMPayload

// NewFranchisereferalincomeCreateObjectV2ResponseMPayload instantiates a new FranchisereferalincomeCreateObjectV2ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFranchisereferalincomeCreateObjectV2ResponseMPayload(aPkiFranchisereferalincomeID []int32) *FranchisereferalincomeCreateObjectV2ResponseMPayload {
	this := FranchisereferalincomeCreateObjectV2ResponseMPayload{}
	this.APkiFranchisereferalincomeID = aPkiFranchisereferalincomeID
	return &this
}

// NewFranchisereferalincomeCreateObjectV2ResponseMPayloadWithDefaults instantiates a new FranchisereferalincomeCreateObjectV2ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFranchisereferalincomeCreateObjectV2ResponseMPayloadWithDefaults() *FranchisereferalincomeCreateObjectV2ResponseMPayload {
	this := FranchisereferalincomeCreateObjectV2ResponseMPayload{}
	return &this
}

// GetAPkiFranchisereferalincomeID returns the APkiFranchisereferalincomeID field value
func (o *FranchisereferalincomeCreateObjectV2ResponseMPayload) GetAPkiFranchisereferalincomeID() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.APkiFranchisereferalincomeID
}

// GetAPkiFranchisereferalincomeIDOk returns a tuple with the APkiFranchisereferalincomeID field value
// and a boolean to check if the value has been set.
func (o *FranchisereferalincomeCreateObjectV2ResponseMPayload) GetAPkiFranchisereferalincomeIDOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.APkiFranchisereferalincomeID, true
}

// SetAPkiFranchisereferalincomeID sets field value
func (o *FranchisereferalincomeCreateObjectV2ResponseMPayload) SetAPkiFranchisereferalincomeID(v []int32) {
	o.APkiFranchisereferalincomeID = v
}

func (o FranchisereferalincomeCreateObjectV2ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FranchisereferalincomeCreateObjectV2ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_pkiFranchisereferalincomeID"] = o.APkiFranchisereferalincomeID
	return toSerialize, nil
}

func (o *FranchisereferalincomeCreateObjectV2ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"a_pkiFranchisereferalincomeID",
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

	varFranchisereferalincomeCreateObjectV2ResponseMPayload := _FranchisereferalincomeCreateObjectV2ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFranchisereferalincomeCreateObjectV2ResponseMPayload)

	if err != nil {
		return err
	}

	*o = FranchisereferalincomeCreateObjectV2ResponseMPayload(varFranchisereferalincomeCreateObjectV2ResponseMPayload)

	return err
}

type NullableFranchisereferalincomeCreateObjectV2ResponseMPayload struct {
	value *FranchisereferalincomeCreateObjectV2ResponseMPayload
	isSet bool
}

func (v NullableFranchisereferalincomeCreateObjectV2ResponseMPayload) Get() *FranchisereferalincomeCreateObjectV2ResponseMPayload {
	return v.value
}

func (v *NullableFranchisereferalincomeCreateObjectV2ResponseMPayload) Set(val *FranchisereferalincomeCreateObjectV2ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableFranchisereferalincomeCreateObjectV2ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableFranchisereferalincomeCreateObjectV2ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFranchisereferalincomeCreateObjectV2ResponseMPayload(val *FranchisereferalincomeCreateObjectV2ResponseMPayload) *NullableFranchisereferalincomeCreateObjectV2ResponseMPayload {
	return &NullableFranchisereferalincomeCreateObjectV2ResponseMPayload{value: val, isSet: true}
}

func (v NullableFranchisereferalincomeCreateObjectV2ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFranchisereferalincomeCreateObjectV2ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


