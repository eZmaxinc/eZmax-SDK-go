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

// checks if the BillingentityinternalCreateObjectV1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BillingentityinternalCreateObjectV1Request{}

// BillingentityinternalCreateObjectV1Request Request for POST /1/object/billingentityinternal
type BillingentityinternalCreateObjectV1Request struct {
	AObjBillingentityinternal []BillingentityinternalRequestCompound `json:"a_objBillingentityinternal"`
}

// NewBillingentityinternalCreateObjectV1Request instantiates a new BillingentityinternalCreateObjectV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingentityinternalCreateObjectV1Request(aObjBillingentityinternal []BillingentityinternalRequestCompound) *BillingentityinternalCreateObjectV1Request {
	this := BillingentityinternalCreateObjectV1Request{}
	this.AObjBillingentityinternal = aObjBillingentityinternal
	return &this
}

// NewBillingentityinternalCreateObjectV1RequestWithDefaults instantiates a new BillingentityinternalCreateObjectV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingentityinternalCreateObjectV1RequestWithDefaults() *BillingentityinternalCreateObjectV1Request {
	this := BillingentityinternalCreateObjectV1Request{}
	return &this
}

// GetAObjBillingentityinternal returns the AObjBillingentityinternal field value
func (o *BillingentityinternalCreateObjectV1Request) GetAObjBillingentityinternal() []BillingentityinternalRequestCompound {
	if o == nil {
		var ret []BillingentityinternalRequestCompound
		return ret
	}

	return o.AObjBillingentityinternal
}

// GetAObjBillingentityinternalOk returns a tuple with the AObjBillingentityinternal field value
// and a boolean to check if the value has been set.
func (o *BillingentityinternalCreateObjectV1Request) GetAObjBillingentityinternalOk() ([]BillingentityinternalRequestCompound, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjBillingentityinternal, true
}

// SetAObjBillingentityinternal sets field value
func (o *BillingentityinternalCreateObjectV1Request) SetAObjBillingentityinternal(v []BillingentityinternalRequestCompound) {
	o.AObjBillingentityinternal = v
}

func (o BillingentityinternalCreateObjectV1Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BillingentityinternalCreateObjectV1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_objBillingentityinternal"] = o.AObjBillingentityinternal
	return toSerialize, nil
}

type NullableBillingentityinternalCreateObjectV1Request struct {
	value *BillingentityinternalCreateObjectV1Request
	isSet bool
}

func (v NullableBillingentityinternalCreateObjectV1Request) Get() *BillingentityinternalCreateObjectV1Request {
	return v.value
}

func (v *NullableBillingentityinternalCreateObjectV1Request) Set(val *BillingentityinternalCreateObjectV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingentityinternalCreateObjectV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingentityinternalCreateObjectV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingentityinternalCreateObjectV1Request(val *BillingentityinternalCreateObjectV1Request) *NullableBillingentityinternalCreateObjectV1Request {
	return &NullableBillingentityinternalCreateObjectV1Request{value: val, isSet: true}
}

func (v NullableBillingentityinternalCreateObjectV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingentityinternalCreateObjectV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


