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

// checks if the EzsigntemplatepublicGetListV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigntemplatepublicGetListV1ResponseMPayload{}

// EzsigntemplatepublicGetListV1ResponseMPayload Payload for GET /1/object/ezsigntemplatepublic/getList
type EzsigntemplatepublicGetListV1ResponseMPayload struct {
	// The number of rows returned
	IRowReturned int32 `json:"iRowReturned"`
	// The number of rows matching your filters (if any) or the total number of rows
	IRowFiltered int32 `json:"iRowFiltered"`
	AObjEzsigntemplatepublic []EzsigntemplatepublicListElement `json:"a_objEzsigntemplatepublic"`
}

type _EzsigntemplatepublicGetListV1ResponseMPayload EzsigntemplatepublicGetListV1ResponseMPayload

// NewEzsigntemplatepublicGetListV1ResponseMPayload instantiates a new EzsigntemplatepublicGetListV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigntemplatepublicGetListV1ResponseMPayload(iRowReturned int32, iRowFiltered int32, aObjEzsigntemplatepublic []EzsigntemplatepublicListElement) *EzsigntemplatepublicGetListV1ResponseMPayload {
	this := EzsigntemplatepublicGetListV1ResponseMPayload{}
	this.IRowReturned = iRowReturned
	this.IRowFiltered = iRowFiltered
	this.AObjEzsigntemplatepublic = aObjEzsigntemplatepublic
	return &this
}

// NewEzsigntemplatepublicGetListV1ResponseMPayloadWithDefaults instantiates a new EzsigntemplatepublicGetListV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigntemplatepublicGetListV1ResponseMPayloadWithDefaults() *EzsigntemplatepublicGetListV1ResponseMPayload {
	this := EzsigntemplatepublicGetListV1ResponseMPayload{}
	return &this
}

// GetIRowReturned returns the IRowReturned field value
func (o *EzsigntemplatepublicGetListV1ResponseMPayload) GetIRowReturned() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IRowReturned
}

// GetIRowReturnedOk returns a tuple with the IRowReturned field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepublicGetListV1ResponseMPayload) GetIRowReturnedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IRowReturned, true
}

// SetIRowReturned sets field value
func (o *EzsigntemplatepublicGetListV1ResponseMPayload) SetIRowReturned(v int32) {
	o.IRowReturned = v
}

// GetIRowFiltered returns the IRowFiltered field value
func (o *EzsigntemplatepublicGetListV1ResponseMPayload) GetIRowFiltered() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IRowFiltered
}

// GetIRowFilteredOk returns a tuple with the IRowFiltered field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepublicGetListV1ResponseMPayload) GetIRowFilteredOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IRowFiltered, true
}

// SetIRowFiltered sets field value
func (o *EzsigntemplatepublicGetListV1ResponseMPayload) SetIRowFiltered(v int32) {
	o.IRowFiltered = v
}

// GetAObjEzsigntemplatepublic returns the AObjEzsigntemplatepublic field value
func (o *EzsigntemplatepublicGetListV1ResponseMPayload) GetAObjEzsigntemplatepublic() []EzsigntemplatepublicListElement {
	if o == nil {
		var ret []EzsigntemplatepublicListElement
		return ret
	}

	return o.AObjEzsigntemplatepublic
}

// GetAObjEzsigntemplatepublicOk returns a tuple with the AObjEzsigntemplatepublic field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepublicGetListV1ResponseMPayload) GetAObjEzsigntemplatepublicOk() ([]EzsigntemplatepublicListElement, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjEzsigntemplatepublic, true
}

// SetAObjEzsigntemplatepublic sets field value
func (o *EzsigntemplatepublicGetListV1ResponseMPayload) SetAObjEzsigntemplatepublic(v []EzsigntemplatepublicListElement) {
	o.AObjEzsigntemplatepublic = v
}

func (o EzsigntemplatepublicGetListV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigntemplatepublicGetListV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["iRowReturned"] = o.IRowReturned
	toSerialize["iRowFiltered"] = o.IRowFiltered
	toSerialize["a_objEzsigntemplatepublic"] = o.AObjEzsigntemplatepublic
	return toSerialize, nil
}

func (o *EzsigntemplatepublicGetListV1ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"iRowReturned",
		"iRowFiltered",
		"a_objEzsigntemplatepublic",
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

	varEzsigntemplatepublicGetListV1ResponseMPayload := _EzsigntemplatepublicGetListV1ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsigntemplatepublicGetListV1ResponseMPayload)

	if err != nil {
		return err
	}

	*o = EzsigntemplatepublicGetListV1ResponseMPayload(varEzsigntemplatepublicGetListV1ResponseMPayload)

	return err
}

type NullableEzsigntemplatepublicGetListV1ResponseMPayload struct {
	value *EzsigntemplatepublicGetListV1ResponseMPayload
	isSet bool
}

func (v NullableEzsigntemplatepublicGetListV1ResponseMPayload) Get() *EzsigntemplatepublicGetListV1ResponseMPayload {
	return v.value
}

func (v *NullableEzsigntemplatepublicGetListV1ResponseMPayload) Set(val *EzsigntemplatepublicGetListV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigntemplatepublicGetListV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigntemplatepublicGetListV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigntemplatepublicGetListV1ResponseMPayload(val *EzsigntemplatepublicGetListV1ResponseMPayload) *NullableEzsigntemplatepublicGetListV1ResponseMPayload {
	return &NullableEzsigntemplatepublicGetListV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableEzsigntemplatepublicGetListV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigntemplatepublicGetListV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


