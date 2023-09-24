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

// checks if the BillingentityinternalCreateObjectV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BillingentityinternalCreateObjectV1ResponseMPayload{}

// BillingentityinternalCreateObjectV1ResponseMPayload Payload for POST /1/object/billingentityinternal
type BillingentityinternalCreateObjectV1ResponseMPayload struct {
	// An array of unique IDs representing the object that were requested to be created.  They are returned in the same order as the array containing the objects to be created that was sent in the request.
	APkiBillingentityinternalID []int32 `json:"a_pkiBillingentityinternalID"`
}

// NewBillingentityinternalCreateObjectV1ResponseMPayload instantiates a new BillingentityinternalCreateObjectV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingentityinternalCreateObjectV1ResponseMPayload(aPkiBillingentityinternalID []int32) *BillingentityinternalCreateObjectV1ResponseMPayload {
	this := BillingentityinternalCreateObjectV1ResponseMPayload{}
	this.APkiBillingentityinternalID = aPkiBillingentityinternalID
	return &this
}

// NewBillingentityinternalCreateObjectV1ResponseMPayloadWithDefaults instantiates a new BillingentityinternalCreateObjectV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingentityinternalCreateObjectV1ResponseMPayloadWithDefaults() *BillingentityinternalCreateObjectV1ResponseMPayload {
	this := BillingentityinternalCreateObjectV1ResponseMPayload{}
	return &this
}

// GetAPkiBillingentityinternalID returns the APkiBillingentityinternalID field value
func (o *BillingentityinternalCreateObjectV1ResponseMPayload) GetAPkiBillingentityinternalID() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.APkiBillingentityinternalID
}

// GetAPkiBillingentityinternalIDOk returns a tuple with the APkiBillingentityinternalID field value
// and a boolean to check if the value has been set.
func (o *BillingentityinternalCreateObjectV1ResponseMPayload) GetAPkiBillingentityinternalIDOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.APkiBillingentityinternalID, true
}

// SetAPkiBillingentityinternalID sets field value
func (o *BillingentityinternalCreateObjectV1ResponseMPayload) SetAPkiBillingentityinternalID(v []int32) {
	o.APkiBillingentityinternalID = v
}

func (o BillingentityinternalCreateObjectV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BillingentityinternalCreateObjectV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_pkiBillingentityinternalID"] = o.APkiBillingentityinternalID
	return toSerialize, nil
}

type NullableBillingentityinternalCreateObjectV1ResponseMPayload struct {
	value *BillingentityinternalCreateObjectV1ResponseMPayload
	isSet bool
}

func (v NullableBillingentityinternalCreateObjectV1ResponseMPayload) Get() *BillingentityinternalCreateObjectV1ResponseMPayload {
	return v.value
}

func (v *NullableBillingentityinternalCreateObjectV1ResponseMPayload) Set(val *BillingentityinternalCreateObjectV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingentityinternalCreateObjectV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingentityinternalCreateObjectV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingentityinternalCreateObjectV1ResponseMPayload(val *BillingentityinternalCreateObjectV1ResponseMPayload) *NullableBillingentityinternalCreateObjectV1ResponseMPayload {
	return &NullableBillingentityinternalCreateObjectV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableBillingentityinternalCreateObjectV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingentityinternalCreateObjectV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

