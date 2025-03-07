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

// checks if the UsergroupCreateObjectV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UsergroupCreateObjectV1ResponseMPayload{}

// UsergroupCreateObjectV1ResponseMPayload Payload for POST /1/object/usergroup
type UsergroupCreateObjectV1ResponseMPayload struct {
	// An array of unique IDs representing the object that were requested to be created.  They are returned in the same order as the array containing the objects to be created that was sent in the request.
	APkiUsergroupID []int32 `json:"a_pkiUsergroupID"`
}

type _UsergroupCreateObjectV1ResponseMPayload UsergroupCreateObjectV1ResponseMPayload

// NewUsergroupCreateObjectV1ResponseMPayload instantiates a new UsergroupCreateObjectV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsergroupCreateObjectV1ResponseMPayload(aPkiUsergroupID []int32) *UsergroupCreateObjectV1ResponseMPayload {
	this := UsergroupCreateObjectV1ResponseMPayload{}
	this.APkiUsergroupID = aPkiUsergroupID
	return &this
}

// NewUsergroupCreateObjectV1ResponseMPayloadWithDefaults instantiates a new UsergroupCreateObjectV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsergroupCreateObjectV1ResponseMPayloadWithDefaults() *UsergroupCreateObjectV1ResponseMPayload {
	this := UsergroupCreateObjectV1ResponseMPayload{}
	return &this
}

// GetAPkiUsergroupID returns the APkiUsergroupID field value
func (o *UsergroupCreateObjectV1ResponseMPayload) GetAPkiUsergroupID() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.APkiUsergroupID
}

// GetAPkiUsergroupIDOk returns a tuple with the APkiUsergroupID field value
// and a boolean to check if the value has been set.
func (o *UsergroupCreateObjectV1ResponseMPayload) GetAPkiUsergroupIDOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.APkiUsergroupID, true
}

// SetAPkiUsergroupID sets field value
func (o *UsergroupCreateObjectV1ResponseMPayload) SetAPkiUsergroupID(v []int32) {
	o.APkiUsergroupID = v
}

func (o UsergroupCreateObjectV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsergroupCreateObjectV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_pkiUsergroupID"] = o.APkiUsergroupID
	return toSerialize, nil
}

func (o *UsergroupCreateObjectV1ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"a_pkiUsergroupID",
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

	varUsergroupCreateObjectV1ResponseMPayload := _UsergroupCreateObjectV1ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUsergroupCreateObjectV1ResponseMPayload)

	if err != nil {
		return err
	}

	*o = UsergroupCreateObjectV1ResponseMPayload(varUsergroupCreateObjectV1ResponseMPayload)

	return err
}

type NullableUsergroupCreateObjectV1ResponseMPayload struct {
	value *UsergroupCreateObjectV1ResponseMPayload
	isSet bool
}

func (v NullableUsergroupCreateObjectV1ResponseMPayload) Get() *UsergroupCreateObjectV1ResponseMPayload {
	return v.value
}

func (v *NullableUsergroupCreateObjectV1ResponseMPayload) Set(val *UsergroupCreateObjectV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableUsergroupCreateObjectV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableUsergroupCreateObjectV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsergroupCreateObjectV1ResponseMPayload(val *UsergroupCreateObjectV1ResponseMPayload) *NullableUsergroupCreateObjectV1ResponseMPayload {
	return &NullableUsergroupCreateObjectV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableUsergroupCreateObjectV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsergroupCreateObjectV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


