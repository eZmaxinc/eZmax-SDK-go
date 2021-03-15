/*
 * eZmax API Definition
 *
 * This API expose all the functionnalities for the eZmax and eZsign applications.
 *
 * API version: 1.0.35
 * Contact: support-api@ezmax.ca
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxinc/eZmax-SDK-go

import (
	"encoding/json"
)

// CommonResponseObjSQLQuery Definition of objSQLQuery Object
type CommonResponseObjSQLQuery struct {
	// The SQL Query
	SQuery string `json:"sQuery"`
	// Execution time of the SQL Query in seconds
	FDuration float32 `json:"fDuration"`
}

// NewCommonResponseObjSQLQuery instantiates a new CommonResponseObjSQLQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonResponseObjSQLQuery(sQuery string, fDuration float32) *CommonResponseObjSQLQuery {
	this := CommonResponseObjSQLQuery{}
	this.SQuery = sQuery
	this.FDuration = fDuration
	return &this
}

// NewCommonResponseObjSQLQueryWithDefaults instantiates a new CommonResponseObjSQLQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonResponseObjSQLQueryWithDefaults() *CommonResponseObjSQLQuery {
	this := CommonResponseObjSQLQuery{}
	return &this
}

// GetSQuery returns the SQuery field value
func (o *CommonResponseObjSQLQuery) GetSQuery() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SQuery
}

// GetSQueryOk returns a tuple with the SQuery field value
// and a boolean to check if the value has been set.
func (o *CommonResponseObjSQLQuery) GetSQueryOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SQuery, true
}

// SetSQuery sets field value
func (o *CommonResponseObjSQLQuery) SetSQuery(v string) {
	o.SQuery = v
}

// GetFDuration returns the FDuration field value
func (o *CommonResponseObjSQLQuery) GetFDuration() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.FDuration
}

// GetFDurationOk returns a tuple with the FDuration field value
// and a boolean to check if the value has been set.
func (o *CommonResponseObjSQLQuery) GetFDurationOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FDuration, true
}

// SetFDuration sets field value
func (o *CommonResponseObjSQLQuery) SetFDuration(v float32) {
	o.FDuration = v
}

func (o CommonResponseObjSQLQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sQuery"] = o.SQuery
	}
	if true {
		toSerialize["fDuration"] = o.FDuration
	}
	return json.Marshal(toSerialize)
}

type NullableCommonResponseObjSQLQuery struct {
	value *CommonResponseObjSQLQuery
	isSet bool
}

func (v NullableCommonResponseObjSQLQuery) Get() *CommonResponseObjSQLQuery {
	return v.value
}

func (v *NullableCommonResponseObjSQLQuery) Set(val *CommonResponseObjSQLQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonResponseObjSQLQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonResponseObjSQLQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonResponseObjSQLQuery(val *CommonResponseObjSQLQuery) *NullableCommonResponseObjSQLQuery {
	return &NullableCommonResponseObjSQLQuery{value: val, isSet: true}
}

func (v NullableCommonResponseObjSQLQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonResponseObjSQLQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


