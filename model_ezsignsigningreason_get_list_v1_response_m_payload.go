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

// checks if the EzsignsigningreasonGetListV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignsigningreasonGetListV1ResponseMPayload{}

// EzsignsigningreasonGetListV1ResponseMPayload Payload for GET /1/object/ezsignsigningreason/getList
type EzsignsigningreasonGetListV1ResponseMPayload struct {
	// The number of rows returned
	IRowReturned int32 `json:"iRowReturned"`
	// The number of rows matching your filters (if any) or the total number of rows
	IRowFiltered int32 `json:"iRowFiltered"`
	AObjEzsignsigningreason []EzsignsigningreasonListElement `json:"a_objEzsignsigningreason"`
}

type _EzsignsigningreasonGetListV1ResponseMPayload EzsignsigningreasonGetListV1ResponseMPayload

// NewEzsignsigningreasonGetListV1ResponseMPayload instantiates a new EzsignsigningreasonGetListV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignsigningreasonGetListV1ResponseMPayload(iRowReturned int32, iRowFiltered int32, aObjEzsignsigningreason []EzsignsigningreasonListElement) *EzsignsigningreasonGetListV1ResponseMPayload {
	this := EzsignsigningreasonGetListV1ResponseMPayload{}
	this.IRowReturned = iRowReturned
	this.IRowFiltered = iRowFiltered
	this.AObjEzsignsigningreason = aObjEzsignsigningreason
	return &this
}

// NewEzsignsigningreasonGetListV1ResponseMPayloadWithDefaults instantiates a new EzsignsigningreasonGetListV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignsigningreasonGetListV1ResponseMPayloadWithDefaults() *EzsignsigningreasonGetListV1ResponseMPayload {
	this := EzsignsigningreasonGetListV1ResponseMPayload{}
	return &this
}

// GetIRowReturned returns the IRowReturned field value
func (o *EzsignsigningreasonGetListV1ResponseMPayload) GetIRowReturned() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IRowReturned
}

// GetIRowReturnedOk returns a tuple with the IRowReturned field value
// and a boolean to check if the value has been set.
func (o *EzsignsigningreasonGetListV1ResponseMPayload) GetIRowReturnedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IRowReturned, true
}

// SetIRowReturned sets field value
func (o *EzsignsigningreasonGetListV1ResponseMPayload) SetIRowReturned(v int32) {
	o.IRowReturned = v
}

// GetIRowFiltered returns the IRowFiltered field value
func (o *EzsignsigningreasonGetListV1ResponseMPayload) GetIRowFiltered() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IRowFiltered
}

// GetIRowFilteredOk returns a tuple with the IRowFiltered field value
// and a boolean to check if the value has been set.
func (o *EzsignsigningreasonGetListV1ResponseMPayload) GetIRowFilteredOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IRowFiltered, true
}

// SetIRowFiltered sets field value
func (o *EzsignsigningreasonGetListV1ResponseMPayload) SetIRowFiltered(v int32) {
	o.IRowFiltered = v
}

// GetAObjEzsignsigningreason returns the AObjEzsignsigningreason field value
func (o *EzsignsigningreasonGetListV1ResponseMPayload) GetAObjEzsignsigningreason() []EzsignsigningreasonListElement {
	if o == nil {
		var ret []EzsignsigningreasonListElement
		return ret
	}

	return o.AObjEzsignsigningreason
}

// GetAObjEzsignsigningreasonOk returns a tuple with the AObjEzsignsigningreason field value
// and a boolean to check if the value has been set.
func (o *EzsignsigningreasonGetListV1ResponseMPayload) GetAObjEzsignsigningreasonOk() ([]EzsignsigningreasonListElement, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjEzsignsigningreason, true
}

// SetAObjEzsignsigningreason sets field value
func (o *EzsignsigningreasonGetListV1ResponseMPayload) SetAObjEzsignsigningreason(v []EzsignsigningreasonListElement) {
	o.AObjEzsignsigningreason = v
}

func (o EzsignsigningreasonGetListV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignsigningreasonGetListV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["iRowReturned"] = o.IRowReturned
	toSerialize["iRowFiltered"] = o.IRowFiltered
	toSerialize["a_objEzsignsigningreason"] = o.AObjEzsignsigningreason
	return toSerialize, nil
}

func (o *EzsignsigningreasonGetListV1ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"iRowReturned",
		"iRowFiltered",
		"a_objEzsignsigningreason",
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

	varEzsignsigningreasonGetListV1ResponseMPayload := _EzsignsigningreasonGetListV1ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsignsigningreasonGetListV1ResponseMPayload)

	if err != nil {
		return err
	}

	*o = EzsignsigningreasonGetListV1ResponseMPayload(varEzsignsigningreasonGetListV1ResponseMPayload)

	return err
}

type NullableEzsignsigningreasonGetListV1ResponseMPayload struct {
	value *EzsignsigningreasonGetListV1ResponseMPayload
	isSet bool
}

func (v NullableEzsignsigningreasonGetListV1ResponseMPayload) Get() *EzsignsigningreasonGetListV1ResponseMPayload {
	return v.value
}

func (v *NullableEzsignsigningreasonGetListV1ResponseMPayload) Set(val *EzsignsigningreasonGetListV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignsigningreasonGetListV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignsigningreasonGetListV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignsigningreasonGetListV1ResponseMPayload(val *EzsignsigningreasonGetListV1ResponseMPayload) *NullableEzsignsigningreasonGetListV1ResponseMPayload {
	return &NullableEzsignsigningreasonGetListV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableEzsignsigningreasonGetListV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignsigningreasonGetListV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


