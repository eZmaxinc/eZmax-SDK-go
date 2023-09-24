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

// checks if the CommonReportcolumn type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommonReportcolumn{}

// CommonReportcolumn A column in a Reportsection 
type CommonReportcolumn struct {
	ObjReportcellstyleDefault CommonReportcellstyle `json:"objReportcellstyleDefault"`
	// The Reportcolumn width in pixels
	IReportcolumnWidth int32 `json:"iReportcolumnWidth"`
}

// NewCommonReportcolumn instantiates a new CommonReportcolumn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonReportcolumn(objReportcellstyleDefault CommonReportcellstyle, iReportcolumnWidth int32) *CommonReportcolumn {
	this := CommonReportcolumn{}
	this.ObjReportcellstyleDefault = objReportcellstyleDefault
	this.IReportcolumnWidth = iReportcolumnWidth
	return &this
}

// NewCommonReportcolumnWithDefaults instantiates a new CommonReportcolumn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonReportcolumnWithDefaults() *CommonReportcolumn {
	this := CommonReportcolumn{}
	return &this
}

// GetObjReportcellstyleDefault returns the ObjReportcellstyleDefault field value
func (o *CommonReportcolumn) GetObjReportcellstyleDefault() CommonReportcellstyle {
	if o == nil {
		var ret CommonReportcellstyle
		return ret
	}

	return o.ObjReportcellstyleDefault
}

// GetObjReportcellstyleDefaultOk returns a tuple with the ObjReportcellstyleDefault field value
// and a boolean to check if the value has been set.
func (o *CommonReportcolumn) GetObjReportcellstyleDefaultOk() (*CommonReportcellstyle, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjReportcellstyleDefault, true
}

// SetObjReportcellstyleDefault sets field value
func (o *CommonReportcolumn) SetObjReportcellstyleDefault(v CommonReportcellstyle) {
	o.ObjReportcellstyleDefault = v
}

// GetIReportcolumnWidth returns the IReportcolumnWidth field value
func (o *CommonReportcolumn) GetIReportcolumnWidth() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IReportcolumnWidth
}

// GetIReportcolumnWidthOk returns a tuple with the IReportcolumnWidth field value
// and a boolean to check if the value has been set.
func (o *CommonReportcolumn) GetIReportcolumnWidthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IReportcolumnWidth, true
}

// SetIReportcolumnWidth sets field value
func (o *CommonReportcolumn) SetIReportcolumnWidth(v int32) {
	o.IReportcolumnWidth = v
}

func (o CommonReportcolumn) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommonReportcolumn) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objReportcellstyleDefault"] = o.ObjReportcellstyleDefault
	toSerialize["iReportcolumnWidth"] = o.IReportcolumnWidth
	return toSerialize, nil
}

type NullableCommonReportcolumn struct {
	value *CommonReportcolumn
	isSet bool
}

func (v NullableCommonReportcolumn) Get() *CommonReportcolumn {
	return v.value
}

func (v *NullableCommonReportcolumn) Set(val *CommonReportcolumn) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonReportcolumn) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonReportcolumn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonReportcolumn(val *CommonReportcolumn) *NullableCommonReportcolumn {
	return &NullableCommonReportcolumn{value: val, isSet: true}
}

func (v NullableCommonReportcolumn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonReportcolumn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

