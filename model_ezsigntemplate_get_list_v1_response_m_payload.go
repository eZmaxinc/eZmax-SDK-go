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

// checks if the EzsigntemplateGetListV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigntemplateGetListV1ResponseMPayload{}

// EzsigntemplateGetListV1ResponseMPayload Payload for GET /1/object/ezsigntemplate/getList
type EzsigntemplateGetListV1ResponseMPayload struct {
	// The number of rows returned
	IRowReturned int32 `json:"iRowReturned"`
	// The number of rows matching your filters (if any) or the total number of rows
	IRowFiltered int32 `json:"iRowFiltered"`
	AObjEzsigntemplate []EzsigntemplateListElement `json:"a_objEzsigntemplate"`
}

type _EzsigntemplateGetListV1ResponseMPayload EzsigntemplateGetListV1ResponseMPayload

// NewEzsigntemplateGetListV1ResponseMPayload instantiates a new EzsigntemplateGetListV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigntemplateGetListV1ResponseMPayload(iRowReturned int32, iRowFiltered int32, aObjEzsigntemplate []EzsigntemplateListElement) *EzsigntemplateGetListV1ResponseMPayload {
	this := EzsigntemplateGetListV1ResponseMPayload{}
	this.IRowReturned = iRowReturned
	this.IRowFiltered = iRowFiltered
	this.AObjEzsigntemplate = aObjEzsigntemplate
	return &this
}

// NewEzsigntemplateGetListV1ResponseMPayloadWithDefaults instantiates a new EzsigntemplateGetListV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigntemplateGetListV1ResponseMPayloadWithDefaults() *EzsigntemplateGetListV1ResponseMPayload {
	this := EzsigntemplateGetListV1ResponseMPayload{}
	return &this
}

// GetIRowReturned returns the IRowReturned field value
func (o *EzsigntemplateGetListV1ResponseMPayload) GetIRowReturned() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IRowReturned
}

// GetIRowReturnedOk returns a tuple with the IRowReturned field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateGetListV1ResponseMPayload) GetIRowReturnedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IRowReturned, true
}

// SetIRowReturned sets field value
func (o *EzsigntemplateGetListV1ResponseMPayload) SetIRowReturned(v int32) {
	o.IRowReturned = v
}

// GetIRowFiltered returns the IRowFiltered field value
func (o *EzsigntemplateGetListV1ResponseMPayload) GetIRowFiltered() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IRowFiltered
}

// GetIRowFilteredOk returns a tuple with the IRowFiltered field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateGetListV1ResponseMPayload) GetIRowFilteredOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IRowFiltered, true
}

// SetIRowFiltered sets field value
func (o *EzsigntemplateGetListV1ResponseMPayload) SetIRowFiltered(v int32) {
	o.IRowFiltered = v
}

// GetAObjEzsigntemplate returns the AObjEzsigntemplate field value
func (o *EzsigntemplateGetListV1ResponseMPayload) GetAObjEzsigntemplate() []EzsigntemplateListElement {
	if o == nil {
		var ret []EzsigntemplateListElement
		return ret
	}

	return o.AObjEzsigntemplate
}

// GetAObjEzsigntemplateOk returns a tuple with the AObjEzsigntemplate field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateGetListV1ResponseMPayload) GetAObjEzsigntemplateOk() ([]EzsigntemplateListElement, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjEzsigntemplate, true
}

// SetAObjEzsigntemplate sets field value
func (o *EzsigntemplateGetListV1ResponseMPayload) SetAObjEzsigntemplate(v []EzsigntemplateListElement) {
	o.AObjEzsigntemplate = v
}

func (o EzsigntemplateGetListV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigntemplateGetListV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["iRowReturned"] = o.IRowReturned
	toSerialize["iRowFiltered"] = o.IRowFiltered
	toSerialize["a_objEzsigntemplate"] = o.AObjEzsigntemplate
	return toSerialize, nil
}

func (o *EzsigntemplateGetListV1ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"iRowReturned",
		"iRowFiltered",
		"a_objEzsigntemplate",
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

	varEzsigntemplateGetListV1ResponseMPayload := _EzsigntemplateGetListV1ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsigntemplateGetListV1ResponseMPayload)

	if err != nil {
		return err
	}

	*o = EzsigntemplateGetListV1ResponseMPayload(varEzsigntemplateGetListV1ResponseMPayload)

	return err
}

type NullableEzsigntemplateGetListV1ResponseMPayload struct {
	value *EzsigntemplateGetListV1ResponseMPayload
	isSet bool
}

func (v NullableEzsigntemplateGetListV1ResponseMPayload) Get() *EzsigntemplateGetListV1ResponseMPayload {
	return v.value
}

func (v *NullableEzsigntemplateGetListV1ResponseMPayload) Set(val *EzsigntemplateGetListV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigntemplateGetListV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigntemplateGetListV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigntemplateGetListV1ResponseMPayload(val *EzsigntemplateGetListV1ResponseMPayload) *NullableEzsigntemplateGetListV1ResponseMPayload {
	return &NullableEzsigntemplateGetListV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableEzsigntemplateGetListV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigntemplateGetListV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


