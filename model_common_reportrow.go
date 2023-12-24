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
	"bytes"
	"fmt"
)

// checks if the CommonReportrow type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommonReportrow{}

// CommonReportrow A row in a Reportsubsectionpart 
type CommonReportrow struct {
	AObjReportcell []CommonReportcell `json:"a_objReportcell"`
	// The reportrow height in pixels
	IReportrowHeight int32 `json:"iReportrowHeight"`
}

type _CommonReportrow CommonReportrow

// NewCommonReportrow instantiates a new CommonReportrow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonReportrow(aObjReportcell []CommonReportcell, iReportrowHeight int32) *CommonReportrow {
	this := CommonReportrow{}
	this.AObjReportcell = aObjReportcell
	this.IReportrowHeight = iReportrowHeight
	return &this
}

// NewCommonReportrowWithDefaults instantiates a new CommonReportrow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonReportrowWithDefaults() *CommonReportrow {
	this := CommonReportrow{}
	return &this
}

// GetAObjReportcell returns the AObjReportcell field value
func (o *CommonReportrow) GetAObjReportcell() []CommonReportcell {
	if o == nil {
		var ret []CommonReportcell
		return ret
	}

	return o.AObjReportcell
}

// GetAObjReportcellOk returns a tuple with the AObjReportcell field value
// and a boolean to check if the value has been set.
func (o *CommonReportrow) GetAObjReportcellOk() ([]CommonReportcell, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjReportcell, true
}

// SetAObjReportcell sets field value
func (o *CommonReportrow) SetAObjReportcell(v []CommonReportcell) {
	o.AObjReportcell = v
}

// GetIReportrowHeight returns the IReportrowHeight field value
func (o *CommonReportrow) GetIReportrowHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IReportrowHeight
}

// GetIReportrowHeightOk returns a tuple with the IReportrowHeight field value
// and a boolean to check if the value has been set.
func (o *CommonReportrow) GetIReportrowHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IReportrowHeight, true
}

// SetIReportrowHeight sets field value
func (o *CommonReportrow) SetIReportrowHeight(v int32) {
	o.IReportrowHeight = v
}

func (o CommonReportrow) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommonReportrow) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_objReportcell"] = o.AObjReportcell
	toSerialize["iReportrowHeight"] = o.IReportrowHeight
	return toSerialize, nil
}

func (o *CommonReportrow) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"a_objReportcell",
		"iReportrowHeight",
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

	varCommonReportrow := _CommonReportrow{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCommonReportrow)

	if err != nil {
		return err
	}

	*o = CommonReportrow(varCommonReportrow)

	return err
}

type NullableCommonReportrow struct {
	value *CommonReportrow
	isSet bool
}

func (v NullableCommonReportrow) Get() *CommonReportrow {
	return v.value
}

func (v *NullableCommonReportrow) Set(val *CommonReportrow) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonReportrow) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonReportrow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonReportrow(val *CommonReportrow) *NullableCommonReportrow {
	return &NullableCommonReportrow{value: val, isSet: true}
}

func (v NullableCommonReportrow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonReportrow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


