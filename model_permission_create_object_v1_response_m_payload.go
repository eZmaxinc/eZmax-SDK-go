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

// checks if the PermissionCreateObjectV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PermissionCreateObjectV1ResponseMPayload{}

// PermissionCreateObjectV1ResponseMPayload Payload for POST /1/object/permission
type PermissionCreateObjectV1ResponseMPayload struct {
	// An array of unique IDs representing the object that were requested to be created.  They are returned in the same order as the array containing the objects to be created that was sent in the request.
	APkiPermissionID []int32 `json:"a_pkiPermissionID"`
}

type _PermissionCreateObjectV1ResponseMPayload PermissionCreateObjectV1ResponseMPayload

// NewPermissionCreateObjectV1ResponseMPayload instantiates a new PermissionCreateObjectV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPermissionCreateObjectV1ResponseMPayload(aPkiPermissionID []int32) *PermissionCreateObjectV1ResponseMPayload {
	this := PermissionCreateObjectV1ResponseMPayload{}
	this.APkiPermissionID = aPkiPermissionID
	return &this
}

// NewPermissionCreateObjectV1ResponseMPayloadWithDefaults instantiates a new PermissionCreateObjectV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPermissionCreateObjectV1ResponseMPayloadWithDefaults() *PermissionCreateObjectV1ResponseMPayload {
	this := PermissionCreateObjectV1ResponseMPayload{}
	return &this
}

// GetAPkiPermissionID returns the APkiPermissionID field value
func (o *PermissionCreateObjectV1ResponseMPayload) GetAPkiPermissionID() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.APkiPermissionID
}

// GetAPkiPermissionIDOk returns a tuple with the APkiPermissionID field value
// and a boolean to check if the value has been set.
func (o *PermissionCreateObjectV1ResponseMPayload) GetAPkiPermissionIDOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.APkiPermissionID, true
}

// SetAPkiPermissionID sets field value
func (o *PermissionCreateObjectV1ResponseMPayload) SetAPkiPermissionID(v []int32) {
	o.APkiPermissionID = v
}

func (o PermissionCreateObjectV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PermissionCreateObjectV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_pkiPermissionID"] = o.APkiPermissionID
	return toSerialize, nil
}

func (o *PermissionCreateObjectV1ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"a_pkiPermissionID",
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

	varPermissionCreateObjectV1ResponseMPayload := _PermissionCreateObjectV1ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPermissionCreateObjectV1ResponseMPayload)

	if err != nil {
		return err
	}

	*o = PermissionCreateObjectV1ResponseMPayload(varPermissionCreateObjectV1ResponseMPayload)

	return err
}

type NullablePermissionCreateObjectV1ResponseMPayload struct {
	value *PermissionCreateObjectV1ResponseMPayload
	isSet bool
}

func (v NullablePermissionCreateObjectV1ResponseMPayload) Get() *PermissionCreateObjectV1ResponseMPayload {
	return v.value
}

func (v *NullablePermissionCreateObjectV1ResponseMPayload) Set(val *PermissionCreateObjectV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullablePermissionCreateObjectV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullablePermissionCreateObjectV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermissionCreateObjectV1ResponseMPayload(val *PermissionCreateObjectV1ResponseMPayload) *NullablePermissionCreateObjectV1ResponseMPayload {
	return &NullablePermissionCreateObjectV1ResponseMPayload{value: val, isSet: true}
}

func (v NullablePermissionCreateObjectV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermissionCreateObjectV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


