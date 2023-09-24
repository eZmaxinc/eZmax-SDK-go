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
)

// checks if the PaymenttermCreateObjectV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymenttermCreateObjectV1ResponseMPayload{}

// PaymenttermCreateObjectV1ResponseMPayload Payload for POST /1/object/paymentterm
type PaymenttermCreateObjectV1ResponseMPayload struct {
	// An array of unique IDs representing the object that were requested to be created.  They are returned in the same order as the array containing the objects to be created that was sent in the request.
	APkiPaymenttermID []int32 `json:"a_pkiPaymenttermID"`
}

// NewPaymenttermCreateObjectV1ResponseMPayload instantiates a new PaymenttermCreateObjectV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymenttermCreateObjectV1ResponseMPayload(aPkiPaymenttermID []int32) *PaymenttermCreateObjectV1ResponseMPayload {
	this := PaymenttermCreateObjectV1ResponseMPayload{}
	this.APkiPaymenttermID = aPkiPaymenttermID
	return &this
}

// NewPaymenttermCreateObjectV1ResponseMPayloadWithDefaults instantiates a new PaymenttermCreateObjectV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymenttermCreateObjectV1ResponseMPayloadWithDefaults() *PaymenttermCreateObjectV1ResponseMPayload {
	this := PaymenttermCreateObjectV1ResponseMPayload{}
	return &this
}

// GetAPkiPaymenttermID returns the APkiPaymenttermID field value
func (o *PaymenttermCreateObjectV1ResponseMPayload) GetAPkiPaymenttermID() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.APkiPaymenttermID
}

// GetAPkiPaymenttermIDOk returns a tuple with the APkiPaymenttermID field value
// and a boolean to check if the value has been set.
func (o *PaymenttermCreateObjectV1ResponseMPayload) GetAPkiPaymenttermIDOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.APkiPaymenttermID, true
}

// SetAPkiPaymenttermID sets field value
func (o *PaymenttermCreateObjectV1ResponseMPayload) SetAPkiPaymenttermID(v []int32) {
	o.APkiPaymenttermID = v
}

func (o PaymenttermCreateObjectV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymenttermCreateObjectV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_pkiPaymenttermID"] = o.APkiPaymenttermID
	return toSerialize, nil
}

type NullablePaymenttermCreateObjectV1ResponseMPayload struct {
	value *PaymenttermCreateObjectV1ResponseMPayload
	isSet bool
}

func (v NullablePaymenttermCreateObjectV1ResponseMPayload) Get() *PaymenttermCreateObjectV1ResponseMPayload {
	return v.value
}

func (v *NullablePaymenttermCreateObjectV1ResponseMPayload) Set(val *PaymenttermCreateObjectV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymenttermCreateObjectV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymenttermCreateObjectV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymenttermCreateObjectV1ResponseMPayload(val *PaymenttermCreateObjectV1ResponseMPayload) *NullablePaymenttermCreateObjectV1ResponseMPayload {
	return &NullablePaymenttermCreateObjectV1ResponseMPayload{value: val, isSet: true}
}

func (v NullablePaymenttermCreateObjectV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymenttermCreateObjectV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

