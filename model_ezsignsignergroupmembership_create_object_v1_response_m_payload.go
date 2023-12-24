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

// checks if the EzsignsignergroupmembershipCreateObjectV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignsignergroupmembershipCreateObjectV1ResponseMPayload{}

// EzsignsignergroupmembershipCreateObjectV1ResponseMPayload Payload for POST /1/object/ezsignsignergroupmembership
type EzsignsignergroupmembershipCreateObjectV1ResponseMPayload struct {
	// An array of unique IDs representing the object that were requested to be created.  They are returned in the same order as the array containing the objects to be created that was sent in the request.
	APkiEzsignsignergroupmembershipID []int32 `json:"a_pkiEzsignsignergroupmembershipID"`
}

type _EzsignsignergroupmembershipCreateObjectV1ResponseMPayload EzsignsignergroupmembershipCreateObjectV1ResponseMPayload

// NewEzsignsignergroupmembershipCreateObjectV1ResponseMPayload instantiates a new EzsignsignergroupmembershipCreateObjectV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignsignergroupmembershipCreateObjectV1ResponseMPayload(aPkiEzsignsignergroupmembershipID []int32) *EzsignsignergroupmembershipCreateObjectV1ResponseMPayload {
	this := EzsignsignergroupmembershipCreateObjectV1ResponseMPayload{}
	this.APkiEzsignsignergroupmembershipID = aPkiEzsignsignergroupmembershipID
	return &this
}

// NewEzsignsignergroupmembershipCreateObjectV1ResponseMPayloadWithDefaults instantiates a new EzsignsignergroupmembershipCreateObjectV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignsignergroupmembershipCreateObjectV1ResponseMPayloadWithDefaults() *EzsignsignergroupmembershipCreateObjectV1ResponseMPayload {
	this := EzsignsignergroupmembershipCreateObjectV1ResponseMPayload{}
	return &this
}

// GetAPkiEzsignsignergroupmembershipID returns the APkiEzsignsignergroupmembershipID field value
func (o *EzsignsignergroupmembershipCreateObjectV1ResponseMPayload) GetAPkiEzsignsignergroupmembershipID() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.APkiEzsignsignergroupmembershipID
}

// GetAPkiEzsignsignergroupmembershipIDOk returns a tuple with the APkiEzsignsignergroupmembershipID field value
// and a boolean to check if the value has been set.
func (o *EzsignsignergroupmembershipCreateObjectV1ResponseMPayload) GetAPkiEzsignsignergroupmembershipIDOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.APkiEzsignsignergroupmembershipID, true
}

// SetAPkiEzsignsignergroupmembershipID sets field value
func (o *EzsignsignergroupmembershipCreateObjectV1ResponseMPayload) SetAPkiEzsignsignergroupmembershipID(v []int32) {
	o.APkiEzsignsignergroupmembershipID = v
}

func (o EzsignsignergroupmembershipCreateObjectV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignsignergroupmembershipCreateObjectV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_pkiEzsignsignergroupmembershipID"] = o.APkiEzsignsignergroupmembershipID
	return toSerialize, nil
}

func (o *EzsignsignergroupmembershipCreateObjectV1ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"a_pkiEzsignsignergroupmembershipID",
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

	varEzsignsignergroupmembershipCreateObjectV1ResponseMPayload := _EzsignsignergroupmembershipCreateObjectV1ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsignsignergroupmembershipCreateObjectV1ResponseMPayload)

	if err != nil {
		return err
	}

	*o = EzsignsignergroupmembershipCreateObjectV1ResponseMPayload(varEzsignsignergroupmembershipCreateObjectV1ResponseMPayload)

	return err
}

type NullableEzsignsignergroupmembershipCreateObjectV1ResponseMPayload struct {
	value *EzsignsignergroupmembershipCreateObjectV1ResponseMPayload
	isSet bool
}

func (v NullableEzsignsignergroupmembershipCreateObjectV1ResponseMPayload) Get() *EzsignsignergroupmembershipCreateObjectV1ResponseMPayload {
	return v.value
}

func (v *NullableEzsignsignergroupmembershipCreateObjectV1ResponseMPayload) Set(val *EzsignsignergroupmembershipCreateObjectV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignsignergroupmembershipCreateObjectV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignsignergroupmembershipCreateObjectV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignsignergroupmembershipCreateObjectV1ResponseMPayload(val *EzsignsignergroupmembershipCreateObjectV1ResponseMPayload) *NullableEzsignsignergroupmembershipCreateObjectV1ResponseMPayload {
	return &NullableEzsignsignergroupmembershipCreateObjectV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableEzsignsignergroupmembershipCreateObjectV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignsignergroupmembershipCreateObjectV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


