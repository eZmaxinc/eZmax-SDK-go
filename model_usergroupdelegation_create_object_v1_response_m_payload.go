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

// checks if the UsergroupdelegationCreateObjectV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UsergroupdelegationCreateObjectV1ResponseMPayload{}

// UsergroupdelegationCreateObjectV1ResponseMPayload Payload for POST /1/object/usergroupdelegation
type UsergroupdelegationCreateObjectV1ResponseMPayload struct {
	// An array of unique IDs representing the object that were requested to be created.  They are returned in the same order as the array containing the objects to be created that was sent in the request.
	APkiUsergroupdelegationID []int32 `json:"a_pkiUsergroupdelegationID"`
}

type _UsergroupdelegationCreateObjectV1ResponseMPayload UsergroupdelegationCreateObjectV1ResponseMPayload

// NewUsergroupdelegationCreateObjectV1ResponseMPayload instantiates a new UsergroupdelegationCreateObjectV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsergroupdelegationCreateObjectV1ResponseMPayload(aPkiUsergroupdelegationID []int32) *UsergroupdelegationCreateObjectV1ResponseMPayload {
	this := UsergroupdelegationCreateObjectV1ResponseMPayload{}
	this.APkiUsergroupdelegationID = aPkiUsergroupdelegationID
	return &this
}

// NewUsergroupdelegationCreateObjectV1ResponseMPayloadWithDefaults instantiates a new UsergroupdelegationCreateObjectV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsergroupdelegationCreateObjectV1ResponseMPayloadWithDefaults() *UsergroupdelegationCreateObjectV1ResponseMPayload {
	this := UsergroupdelegationCreateObjectV1ResponseMPayload{}
	return &this
}

// GetAPkiUsergroupdelegationID returns the APkiUsergroupdelegationID field value
func (o *UsergroupdelegationCreateObjectV1ResponseMPayload) GetAPkiUsergroupdelegationID() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.APkiUsergroupdelegationID
}

// GetAPkiUsergroupdelegationIDOk returns a tuple with the APkiUsergroupdelegationID field value
// and a boolean to check if the value has been set.
func (o *UsergroupdelegationCreateObjectV1ResponseMPayload) GetAPkiUsergroupdelegationIDOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.APkiUsergroupdelegationID, true
}

// SetAPkiUsergroupdelegationID sets field value
func (o *UsergroupdelegationCreateObjectV1ResponseMPayload) SetAPkiUsergroupdelegationID(v []int32) {
	o.APkiUsergroupdelegationID = v
}

func (o UsergroupdelegationCreateObjectV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsergroupdelegationCreateObjectV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_pkiUsergroupdelegationID"] = o.APkiUsergroupdelegationID
	return toSerialize, nil
}

func (o *UsergroupdelegationCreateObjectV1ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"a_pkiUsergroupdelegationID",
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

	varUsergroupdelegationCreateObjectV1ResponseMPayload := _UsergroupdelegationCreateObjectV1ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUsergroupdelegationCreateObjectV1ResponseMPayload)

	if err != nil {
		return err
	}

	*o = UsergroupdelegationCreateObjectV1ResponseMPayload(varUsergroupdelegationCreateObjectV1ResponseMPayload)

	return err
}

type NullableUsergroupdelegationCreateObjectV1ResponseMPayload struct {
	value *UsergroupdelegationCreateObjectV1ResponseMPayload
	isSet bool
}

func (v NullableUsergroupdelegationCreateObjectV1ResponseMPayload) Get() *UsergroupdelegationCreateObjectV1ResponseMPayload {
	return v.value
}

func (v *NullableUsergroupdelegationCreateObjectV1ResponseMPayload) Set(val *UsergroupdelegationCreateObjectV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableUsergroupdelegationCreateObjectV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableUsergroupdelegationCreateObjectV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsergroupdelegationCreateObjectV1ResponseMPayload(val *UsergroupdelegationCreateObjectV1ResponseMPayload) *NullableUsergroupdelegationCreateObjectV1ResponseMPayload {
	return &NullableUsergroupdelegationCreateObjectV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableUsergroupdelegationCreateObjectV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsergroupdelegationCreateObjectV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


