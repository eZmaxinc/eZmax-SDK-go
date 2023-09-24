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

// checks if the ActivesessionGetListV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActivesessionGetListV1ResponseMPayload{}

// ActivesessionGetListV1ResponseMPayload Payload for GET /1/object/activesession/getList
type ActivesessionGetListV1ResponseMPayload struct {
	// The number of rows returned
	IRowReturned int32 `json:"iRowReturned"`
	// The number of rows matching your filters (if any) or the total number of rows
	IRowFiltered int32 `json:"iRowFiltered"`
	AObjActivesession []ActivesessionListElement `json:"a_objActivesession"`
}

// NewActivesessionGetListV1ResponseMPayload instantiates a new ActivesessionGetListV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActivesessionGetListV1ResponseMPayload(iRowReturned int32, iRowFiltered int32, aObjActivesession []ActivesessionListElement) *ActivesessionGetListV1ResponseMPayload {
	this := ActivesessionGetListV1ResponseMPayload{}
	this.IRowReturned = iRowReturned
	this.IRowFiltered = iRowFiltered
	this.AObjActivesession = aObjActivesession
	return &this
}

// NewActivesessionGetListV1ResponseMPayloadWithDefaults instantiates a new ActivesessionGetListV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActivesessionGetListV1ResponseMPayloadWithDefaults() *ActivesessionGetListV1ResponseMPayload {
	this := ActivesessionGetListV1ResponseMPayload{}
	return &this
}

// GetIRowReturned returns the IRowReturned field value
func (o *ActivesessionGetListV1ResponseMPayload) GetIRowReturned() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IRowReturned
}

// GetIRowReturnedOk returns a tuple with the IRowReturned field value
// and a boolean to check if the value has been set.
func (o *ActivesessionGetListV1ResponseMPayload) GetIRowReturnedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IRowReturned, true
}

// SetIRowReturned sets field value
func (o *ActivesessionGetListV1ResponseMPayload) SetIRowReturned(v int32) {
	o.IRowReturned = v
}

// GetIRowFiltered returns the IRowFiltered field value
func (o *ActivesessionGetListV1ResponseMPayload) GetIRowFiltered() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IRowFiltered
}

// GetIRowFilteredOk returns a tuple with the IRowFiltered field value
// and a boolean to check if the value has been set.
func (o *ActivesessionGetListV1ResponseMPayload) GetIRowFilteredOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IRowFiltered, true
}

// SetIRowFiltered sets field value
func (o *ActivesessionGetListV1ResponseMPayload) SetIRowFiltered(v int32) {
	o.IRowFiltered = v
}

// GetAObjActivesession returns the AObjActivesession field value
func (o *ActivesessionGetListV1ResponseMPayload) GetAObjActivesession() []ActivesessionListElement {
	if o == nil {
		var ret []ActivesessionListElement
		return ret
	}

	return o.AObjActivesession
}

// GetAObjActivesessionOk returns a tuple with the AObjActivesession field value
// and a boolean to check if the value has been set.
func (o *ActivesessionGetListV1ResponseMPayload) GetAObjActivesessionOk() ([]ActivesessionListElement, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjActivesession, true
}

// SetAObjActivesession sets field value
func (o *ActivesessionGetListV1ResponseMPayload) SetAObjActivesession(v []ActivesessionListElement) {
	o.AObjActivesession = v
}

func (o ActivesessionGetListV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActivesessionGetListV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["iRowReturned"] = o.IRowReturned
	toSerialize["iRowFiltered"] = o.IRowFiltered
	toSerialize["a_objActivesession"] = o.AObjActivesession
	return toSerialize, nil
}

type NullableActivesessionGetListV1ResponseMPayload struct {
	value *ActivesessionGetListV1ResponseMPayload
	isSet bool
}

func (v NullableActivesessionGetListV1ResponseMPayload) Get() *ActivesessionGetListV1ResponseMPayload {
	return v.value
}

func (v *NullableActivesessionGetListV1ResponseMPayload) Set(val *ActivesessionGetListV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableActivesessionGetListV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableActivesessionGetListV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActivesessionGetListV1ResponseMPayload(val *ActivesessionGetListV1ResponseMPayload) *NullableActivesessionGetListV1ResponseMPayload {
	return &NullableActivesessionGetListV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableActivesessionGetListV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActivesessionGetListV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


