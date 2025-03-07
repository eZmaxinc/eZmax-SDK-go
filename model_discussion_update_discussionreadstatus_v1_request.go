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
)

// checks if the DiscussionUpdateDiscussionreadstatusV1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DiscussionUpdateDiscussionreadstatusV1Request{}

// DiscussionUpdateDiscussionreadstatusV1Request Request for POST /1/object/discussion/{pkiDiscussionID}/updateDiscussionreadstatus
type DiscussionUpdateDiscussionreadstatusV1Request struct {
	// The date of the last discussion message read
	DtDiscussionreadstatusDate *string `json:"dtDiscussionreadstatusDate,omitempty" validate:"regexp=^[0-9]{4}-(0[1-9]|1[0-2])-(0[1-9]|[1-2][0-9]|3[0-1]) ([01]?[0-9]|2[0-3]):([0-5][0-9]):([0-5][0-9])$"`
}

// NewDiscussionUpdateDiscussionreadstatusV1Request instantiates a new DiscussionUpdateDiscussionreadstatusV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiscussionUpdateDiscussionreadstatusV1Request() *DiscussionUpdateDiscussionreadstatusV1Request {
	this := DiscussionUpdateDiscussionreadstatusV1Request{}
	return &this
}

// NewDiscussionUpdateDiscussionreadstatusV1RequestWithDefaults instantiates a new DiscussionUpdateDiscussionreadstatusV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiscussionUpdateDiscussionreadstatusV1RequestWithDefaults() *DiscussionUpdateDiscussionreadstatusV1Request {
	this := DiscussionUpdateDiscussionreadstatusV1Request{}
	return &this
}

// GetDtDiscussionreadstatusDate returns the DtDiscussionreadstatusDate field value if set, zero value otherwise.
func (o *DiscussionUpdateDiscussionreadstatusV1Request) GetDtDiscussionreadstatusDate() string {
	if o == nil || IsNil(o.DtDiscussionreadstatusDate) {
		var ret string
		return ret
	}
	return *o.DtDiscussionreadstatusDate
}

// GetDtDiscussionreadstatusDateOk returns a tuple with the DtDiscussionreadstatusDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscussionUpdateDiscussionreadstatusV1Request) GetDtDiscussionreadstatusDateOk() (*string, bool) {
	if o == nil || IsNil(o.DtDiscussionreadstatusDate) {
		return nil, false
	}
	return o.DtDiscussionreadstatusDate, true
}

// HasDtDiscussionreadstatusDate returns a boolean if a field has been set.
func (o *DiscussionUpdateDiscussionreadstatusV1Request) HasDtDiscussionreadstatusDate() bool {
	if o != nil && !IsNil(o.DtDiscussionreadstatusDate) {
		return true
	}

	return false
}

// SetDtDiscussionreadstatusDate gets a reference to the given string and assigns it to the DtDiscussionreadstatusDate field.
func (o *DiscussionUpdateDiscussionreadstatusV1Request) SetDtDiscussionreadstatusDate(v string) {
	o.DtDiscussionreadstatusDate = &v
}

func (o DiscussionUpdateDiscussionreadstatusV1Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DiscussionUpdateDiscussionreadstatusV1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DtDiscussionreadstatusDate) {
		toSerialize["dtDiscussionreadstatusDate"] = o.DtDiscussionreadstatusDate
	}
	return toSerialize, nil
}

type NullableDiscussionUpdateDiscussionreadstatusV1Request struct {
	value *DiscussionUpdateDiscussionreadstatusV1Request
	isSet bool
}

func (v NullableDiscussionUpdateDiscussionreadstatusV1Request) Get() *DiscussionUpdateDiscussionreadstatusV1Request {
	return v.value
}

func (v *NullableDiscussionUpdateDiscussionreadstatusV1Request) Set(val *DiscussionUpdateDiscussionreadstatusV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableDiscussionUpdateDiscussionreadstatusV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableDiscussionUpdateDiscussionreadstatusV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiscussionUpdateDiscussionreadstatusV1Request(val *DiscussionUpdateDiscussionreadstatusV1Request) *NullableDiscussionUpdateDiscussionreadstatusV1Request {
	return &NullableDiscussionUpdateDiscussionreadstatusV1Request{value: val, isSet: true}
}

func (v NullableDiscussionUpdateDiscussionreadstatusV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiscussionUpdateDiscussionreadstatusV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


