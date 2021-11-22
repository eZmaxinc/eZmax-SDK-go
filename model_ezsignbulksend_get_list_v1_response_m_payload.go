/*
eZmax API Definition

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.1.3
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
)

// EzsignbulksendGetListV1ResponseMPayload Payload for the /1/object/ezsignbulksend/getList API Request
type EzsignbulksendGetListV1ResponseMPayload struct {
	AObjEzsignfolder *[]EzsignbulksendListElement `json:"a_objEzsignfolder,omitempty"`
	// The number of rows returned
	IRowReturned int32 `json:"iRowReturned"`
	// The number of rows matching your filters (if any) or the total number of rows
	IRowFiltered int32 `json:"iRowFiltered"`
}

// NewEzsignbulksendGetListV1ResponseMPayload instantiates a new EzsignbulksendGetListV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignbulksendGetListV1ResponseMPayload(iRowReturned int32, iRowFiltered int32) *EzsignbulksendGetListV1ResponseMPayload {
	this := EzsignbulksendGetListV1ResponseMPayload{}
	this.IRowReturned = iRowReturned
	this.IRowFiltered = iRowFiltered
	return &this
}

// NewEzsignbulksendGetListV1ResponseMPayloadWithDefaults instantiates a new EzsignbulksendGetListV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignbulksendGetListV1ResponseMPayloadWithDefaults() *EzsignbulksendGetListV1ResponseMPayload {
	this := EzsignbulksendGetListV1ResponseMPayload{}
	return &this
}

// GetAObjEzsignfolder returns the AObjEzsignfolder field value if set, zero value otherwise.
func (o *EzsignbulksendGetListV1ResponseMPayload) GetAObjEzsignfolder() []EzsignbulksendListElement {
	if o == nil || o.AObjEzsignfolder == nil {
		var ret []EzsignbulksendListElement
		return ret
	}
	return *o.AObjEzsignfolder
}

// GetAObjEzsignfolderOk returns a tuple with the AObjEzsignfolder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignbulksendGetListV1ResponseMPayload) GetAObjEzsignfolderOk() (*[]EzsignbulksendListElement, bool) {
	if o == nil || o.AObjEzsignfolder == nil {
		return nil, false
	}
	return o.AObjEzsignfolder, true
}

// HasAObjEzsignfolder returns a boolean if a field has been set.
func (o *EzsignbulksendGetListV1ResponseMPayload) HasAObjEzsignfolder() bool {
	if o != nil && o.AObjEzsignfolder != nil {
		return true
	}

	return false
}

// SetAObjEzsignfolder gets a reference to the given []EzsignbulksendListElement and assigns it to the AObjEzsignfolder field.
func (o *EzsignbulksendGetListV1ResponseMPayload) SetAObjEzsignfolder(v []EzsignbulksendListElement) {
	o.AObjEzsignfolder = &v
}

// GetIRowReturned returns the IRowReturned field value
func (o *EzsignbulksendGetListV1ResponseMPayload) GetIRowReturned() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IRowReturned
}

// GetIRowReturnedOk returns a tuple with the IRowReturned field value
// and a boolean to check if the value has been set.
func (o *EzsignbulksendGetListV1ResponseMPayload) GetIRowReturnedOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IRowReturned, true
}

// SetIRowReturned sets field value
func (o *EzsignbulksendGetListV1ResponseMPayload) SetIRowReturned(v int32) {
	o.IRowReturned = v
}

// GetIRowFiltered returns the IRowFiltered field value
func (o *EzsignbulksendGetListV1ResponseMPayload) GetIRowFiltered() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IRowFiltered
}

// GetIRowFilteredOk returns a tuple with the IRowFiltered field value
// and a boolean to check if the value has been set.
func (o *EzsignbulksendGetListV1ResponseMPayload) GetIRowFilteredOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IRowFiltered, true
}

// SetIRowFiltered sets field value
func (o *EzsignbulksendGetListV1ResponseMPayload) SetIRowFiltered(v int32) {
	o.IRowFiltered = v
}

func (o EzsignbulksendGetListV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AObjEzsignfolder != nil {
		toSerialize["a_objEzsignfolder"] = o.AObjEzsignfolder
	}
	if true {
		toSerialize["iRowReturned"] = o.IRowReturned
	}
	if true {
		toSerialize["iRowFiltered"] = o.IRowFiltered
	}
	return json.Marshal(toSerialize)
}

type NullableEzsignbulksendGetListV1ResponseMPayload struct {
	value *EzsignbulksendGetListV1ResponseMPayload
	isSet bool
}

func (v NullableEzsignbulksendGetListV1ResponseMPayload) Get() *EzsignbulksendGetListV1ResponseMPayload {
	return v.value
}

func (v *NullableEzsignbulksendGetListV1ResponseMPayload) Set(val *EzsignbulksendGetListV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignbulksendGetListV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignbulksendGetListV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignbulksendGetListV1ResponseMPayload(val *EzsignbulksendGetListV1ResponseMPayload) *NullableEzsignbulksendGetListV1ResponseMPayload {
	return &NullableEzsignbulksendGetListV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableEzsignbulksendGetListV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignbulksendGetListV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


