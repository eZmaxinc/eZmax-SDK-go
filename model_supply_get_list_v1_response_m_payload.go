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

// checks if the SupplyGetListV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SupplyGetListV1ResponseMPayload{}

// SupplyGetListV1ResponseMPayload Payload for GET /1/object/supply/getList
type SupplyGetListV1ResponseMPayload struct {
	// The number of rows returned
	IRowReturned int32 `json:"iRowReturned"`
	// The number of rows matching your filters (if any) or the total number of rows
	IRowFiltered int32 `json:"iRowFiltered"`
	AObjSupply []SupplyListElement `json:"a_objSupply"`
}

type _SupplyGetListV1ResponseMPayload SupplyGetListV1ResponseMPayload

// NewSupplyGetListV1ResponseMPayload instantiates a new SupplyGetListV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSupplyGetListV1ResponseMPayload(iRowReturned int32, iRowFiltered int32, aObjSupply []SupplyListElement) *SupplyGetListV1ResponseMPayload {
	this := SupplyGetListV1ResponseMPayload{}
	this.IRowReturned = iRowReturned
	this.IRowFiltered = iRowFiltered
	this.AObjSupply = aObjSupply
	return &this
}

// NewSupplyGetListV1ResponseMPayloadWithDefaults instantiates a new SupplyGetListV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSupplyGetListV1ResponseMPayloadWithDefaults() *SupplyGetListV1ResponseMPayload {
	this := SupplyGetListV1ResponseMPayload{}
	return &this
}

// GetIRowReturned returns the IRowReturned field value
func (o *SupplyGetListV1ResponseMPayload) GetIRowReturned() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IRowReturned
}

// GetIRowReturnedOk returns a tuple with the IRowReturned field value
// and a boolean to check if the value has been set.
func (o *SupplyGetListV1ResponseMPayload) GetIRowReturnedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IRowReturned, true
}

// SetIRowReturned sets field value
func (o *SupplyGetListV1ResponseMPayload) SetIRowReturned(v int32) {
	o.IRowReturned = v
}

// GetIRowFiltered returns the IRowFiltered field value
func (o *SupplyGetListV1ResponseMPayload) GetIRowFiltered() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IRowFiltered
}

// GetIRowFilteredOk returns a tuple with the IRowFiltered field value
// and a boolean to check if the value has been set.
func (o *SupplyGetListV1ResponseMPayload) GetIRowFilteredOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IRowFiltered, true
}

// SetIRowFiltered sets field value
func (o *SupplyGetListV1ResponseMPayload) SetIRowFiltered(v int32) {
	o.IRowFiltered = v
}

// GetAObjSupply returns the AObjSupply field value
func (o *SupplyGetListV1ResponseMPayload) GetAObjSupply() []SupplyListElement {
	if o == nil {
		var ret []SupplyListElement
		return ret
	}

	return o.AObjSupply
}

// GetAObjSupplyOk returns a tuple with the AObjSupply field value
// and a boolean to check if the value has been set.
func (o *SupplyGetListV1ResponseMPayload) GetAObjSupplyOk() ([]SupplyListElement, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjSupply, true
}

// SetAObjSupply sets field value
func (o *SupplyGetListV1ResponseMPayload) SetAObjSupply(v []SupplyListElement) {
	o.AObjSupply = v
}

func (o SupplyGetListV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SupplyGetListV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["iRowReturned"] = o.IRowReturned
	toSerialize["iRowFiltered"] = o.IRowFiltered
	toSerialize["a_objSupply"] = o.AObjSupply
	return toSerialize, nil
}

func (o *SupplyGetListV1ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"iRowReturned",
		"iRowFiltered",
		"a_objSupply",
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

	varSupplyGetListV1ResponseMPayload := _SupplyGetListV1ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSupplyGetListV1ResponseMPayload)

	if err != nil {
		return err
	}

	*o = SupplyGetListV1ResponseMPayload(varSupplyGetListV1ResponseMPayload)

	return err
}

type NullableSupplyGetListV1ResponseMPayload struct {
	value *SupplyGetListV1ResponseMPayload
	isSet bool
}

func (v NullableSupplyGetListV1ResponseMPayload) Get() *SupplyGetListV1ResponseMPayload {
	return v.value
}

func (v *NullableSupplyGetListV1ResponseMPayload) Set(val *SupplyGetListV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableSupplyGetListV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableSupplyGetListV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSupplyGetListV1ResponseMPayload(val *SupplyGetListV1ResponseMPayload) *NullableSupplyGetListV1ResponseMPayload {
	return &NullableSupplyGetListV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableSupplyGetListV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSupplyGetListV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


