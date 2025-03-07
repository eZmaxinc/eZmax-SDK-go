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

// checks if the ApikeyGetListV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApikeyGetListV1ResponseMPayload{}

// ApikeyGetListV1ResponseMPayload Payload for GET /1/object/apikey/getList
type ApikeyGetListV1ResponseMPayload struct {
	// The number of rows returned
	IRowReturned int32 `json:"iRowReturned"`
	// The number of rows matching your filters (if any) or the total number of rows
	IRowFiltered int32 `json:"iRowFiltered"`
	AObjApikey []ApikeyListElement `json:"a_objApikey"`
}

type _ApikeyGetListV1ResponseMPayload ApikeyGetListV1ResponseMPayload

// NewApikeyGetListV1ResponseMPayload instantiates a new ApikeyGetListV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApikeyGetListV1ResponseMPayload(iRowReturned int32, iRowFiltered int32, aObjApikey []ApikeyListElement) *ApikeyGetListV1ResponseMPayload {
	this := ApikeyGetListV1ResponseMPayload{}
	this.IRowReturned = iRowReturned
	this.IRowFiltered = iRowFiltered
	this.AObjApikey = aObjApikey
	return &this
}

// NewApikeyGetListV1ResponseMPayloadWithDefaults instantiates a new ApikeyGetListV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApikeyGetListV1ResponseMPayloadWithDefaults() *ApikeyGetListV1ResponseMPayload {
	this := ApikeyGetListV1ResponseMPayload{}
	return &this
}

// GetIRowReturned returns the IRowReturned field value
func (o *ApikeyGetListV1ResponseMPayload) GetIRowReturned() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IRowReturned
}

// GetIRowReturnedOk returns a tuple with the IRowReturned field value
// and a boolean to check if the value has been set.
func (o *ApikeyGetListV1ResponseMPayload) GetIRowReturnedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IRowReturned, true
}

// SetIRowReturned sets field value
func (o *ApikeyGetListV1ResponseMPayload) SetIRowReturned(v int32) {
	o.IRowReturned = v
}

// GetIRowFiltered returns the IRowFiltered field value
func (o *ApikeyGetListV1ResponseMPayload) GetIRowFiltered() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IRowFiltered
}

// GetIRowFilteredOk returns a tuple with the IRowFiltered field value
// and a boolean to check if the value has been set.
func (o *ApikeyGetListV1ResponseMPayload) GetIRowFilteredOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IRowFiltered, true
}

// SetIRowFiltered sets field value
func (o *ApikeyGetListV1ResponseMPayload) SetIRowFiltered(v int32) {
	o.IRowFiltered = v
}

// GetAObjApikey returns the AObjApikey field value
func (o *ApikeyGetListV1ResponseMPayload) GetAObjApikey() []ApikeyListElement {
	if o == nil {
		var ret []ApikeyListElement
		return ret
	}

	return o.AObjApikey
}

// GetAObjApikeyOk returns a tuple with the AObjApikey field value
// and a boolean to check if the value has been set.
func (o *ApikeyGetListV1ResponseMPayload) GetAObjApikeyOk() ([]ApikeyListElement, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjApikey, true
}

// SetAObjApikey sets field value
func (o *ApikeyGetListV1ResponseMPayload) SetAObjApikey(v []ApikeyListElement) {
	o.AObjApikey = v
}

func (o ApikeyGetListV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApikeyGetListV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["iRowReturned"] = o.IRowReturned
	toSerialize["iRowFiltered"] = o.IRowFiltered
	toSerialize["a_objApikey"] = o.AObjApikey
	return toSerialize, nil
}

func (o *ApikeyGetListV1ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"iRowReturned",
		"iRowFiltered",
		"a_objApikey",
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

	varApikeyGetListV1ResponseMPayload := _ApikeyGetListV1ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApikeyGetListV1ResponseMPayload)

	if err != nil {
		return err
	}

	*o = ApikeyGetListV1ResponseMPayload(varApikeyGetListV1ResponseMPayload)

	return err
}

type NullableApikeyGetListV1ResponseMPayload struct {
	value *ApikeyGetListV1ResponseMPayload
	isSet bool
}

func (v NullableApikeyGetListV1ResponseMPayload) Get() *ApikeyGetListV1ResponseMPayload {
	return v.value
}

func (v *NullableApikeyGetListV1ResponseMPayload) Set(val *ApikeyGetListV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableApikeyGetListV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableApikeyGetListV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApikeyGetListV1ResponseMPayload(val *ApikeyGetListV1ResponseMPayload) *NullableApikeyGetListV1ResponseMPayload {
	return &NullableApikeyGetListV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableApikeyGetListV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApikeyGetListV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


