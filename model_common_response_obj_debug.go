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

// checks if the CommonResponseObjDebug type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommonResponseObjDebug{}

// CommonResponseObjDebug This is a generic debug object that is returned by all API requests
type CommonResponseObjDebug struct {
	// The peak memory allocated during the API request execution. Formatted as a human readable string
	SMemoryUsage string `json:"sMemoryUsage"`
	// The total server execution time of the API request execution. Formatted as a human readable string
	SRunTime string `json:"sRunTime"`
	// The number of SQL SELECT queries that were sent to the database server during the API request execution
	ISQLSelects int32 `json:"iSQLSelects"`
	// The number of SQL INSERT/UPDATE/DELETE queries that were sent to the database server during the API request execution
	ISQLQueries int32 `json:"iSQLQueries"`
	// An array of the SQL Queries that were executed during the API request execution
	AObjSQLQuery []CommonResponseObjSQLQuery `json:"a_objSQLQuery"`
}

type _CommonResponseObjDebug CommonResponseObjDebug

// NewCommonResponseObjDebug instantiates a new CommonResponseObjDebug object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonResponseObjDebug(sMemoryUsage string, sRunTime string, iSQLSelects int32, iSQLQueries int32, aObjSQLQuery []CommonResponseObjSQLQuery) *CommonResponseObjDebug {
	this := CommonResponseObjDebug{}
	this.SMemoryUsage = sMemoryUsage
	this.SRunTime = sRunTime
	this.ISQLSelects = iSQLSelects
	this.ISQLQueries = iSQLQueries
	this.AObjSQLQuery = aObjSQLQuery
	return &this
}

// NewCommonResponseObjDebugWithDefaults instantiates a new CommonResponseObjDebug object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonResponseObjDebugWithDefaults() *CommonResponseObjDebug {
	this := CommonResponseObjDebug{}
	return &this
}

// GetSMemoryUsage returns the SMemoryUsage field value
func (o *CommonResponseObjDebug) GetSMemoryUsage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SMemoryUsage
}

// GetSMemoryUsageOk returns a tuple with the SMemoryUsage field value
// and a boolean to check if the value has been set.
func (o *CommonResponseObjDebug) GetSMemoryUsageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SMemoryUsage, true
}

// SetSMemoryUsage sets field value
func (o *CommonResponseObjDebug) SetSMemoryUsage(v string) {
	o.SMemoryUsage = v
}

// GetSRunTime returns the SRunTime field value
func (o *CommonResponseObjDebug) GetSRunTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SRunTime
}

// GetSRunTimeOk returns a tuple with the SRunTime field value
// and a boolean to check if the value has been set.
func (o *CommonResponseObjDebug) GetSRunTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SRunTime, true
}

// SetSRunTime sets field value
func (o *CommonResponseObjDebug) SetSRunTime(v string) {
	o.SRunTime = v
}

// GetISQLSelects returns the ISQLSelects field value
func (o *CommonResponseObjDebug) GetISQLSelects() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ISQLSelects
}

// GetISQLSelectsOk returns a tuple with the ISQLSelects field value
// and a boolean to check if the value has been set.
func (o *CommonResponseObjDebug) GetISQLSelectsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ISQLSelects, true
}

// SetISQLSelects sets field value
func (o *CommonResponseObjDebug) SetISQLSelects(v int32) {
	o.ISQLSelects = v
}

// GetISQLQueries returns the ISQLQueries field value
func (o *CommonResponseObjDebug) GetISQLQueries() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ISQLQueries
}

// GetISQLQueriesOk returns a tuple with the ISQLQueries field value
// and a boolean to check if the value has been set.
func (o *CommonResponseObjDebug) GetISQLQueriesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ISQLQueries, true
}

// SetISQLQueries sets field value
func (o *CommonResponseObjDebug) SetISQLQueries(v int32) {
	o.ISQLQueries = v
}

// GetAObjSQLQuery returns the AObjSQLQuery field value
func (o *CommonResponseObjDebug) GetAObjSQLQuery() []CommonResponseObjSQLQuery {
	if o == nil {
		var ret []CommonResponseObjSQLQuery
		return ret
	}

	return o.AObjSQLQuery
}

// GetAObjSQLQueryOk returns a tuple with the AObjSQLQuery field value
// and a boolean to check if the value has been set.
func (o *CommonResponseObjDebug) GetAObjSQLQueryOk() ([]CommonResponseObjSQLQuery, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjSQLQuery, true
}

// SetAObjSQLQuery sets field value
func (o *CommonResponseObjDebug) SetAObjSQLQuery(v []CommonResponseObjSQLQuery) {
	o.AObjSQLQuery = v
}

func (o CommonResponseObjDebug) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommonResponseObjDebug) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sMemoryUsage"] = o.SMemoryUsage
	toSerialize["sRunTime"] = o.SRunTime
	toSerialize["iSQLSelects"] = o.ISQLSelects
	toSerialize["iSQLQueries"] = o.ISQLQueries
	toSerialize["a_objSQLQuery"] = o.AObjSQLQuery
	return toSerialize, nil
}

func (o *CommonResponseObjDebug) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"sMemoryUsage",
		"sRunTime",
		"iSQLSelects",
		"iSQLQueries",
		"a_objSQLQuery",
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

	varCommonResponseObjDebug := _CommonResponseObjDebug{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCommonResponseObjDebug)

	if err != nil {
		return err
	}

	*o = CommonResponseObjDebug(varCommonResponseObjDebug)

	return err
}

type NullableCommonResponseObjDebug struct {
	value *CommonResponseObjDebug
	isSet bool
}

func (v NullableCommonResponseObjDebug) Get() *CommonResponseObjDebug {
	return v.value
}

func (v *NullableCommonResponseObjDebug) Set(val *CommonResponseObjDebug) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonResponseObjDebug) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonResponseObjDebug) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonResponseObjDebug(val *CommonResponseObjDebug) *NullableCommonResponseObjDebug {
	return &NullableCommonResponseObjDebug{value: val, isSet: true}
}

func (v NullableCommonResponseObjDebug) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonResponseObjDebug) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


