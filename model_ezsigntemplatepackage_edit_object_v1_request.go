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

// checks if the EzsigntemplatepackageEditObjectV1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigntemplatepackageEditObjectV1Request{}

// EzsigntemplatepackageEditObjectV1Request Request for PUT /1/object/ezsigntemplatepackage/{pkiEzsigntemplatepackageID}
type EzsigntemplatepackageEditObjectV1Request struct {
	ObjEzsigntemplatepackage EzsigntemplatepackageRequestCompound `json:"objEzsigntemplatepackage"`
}

// NewEzsigntemplatepackageEditObjectV1Request instantiates a new EzsigntemplatepackageEditObjectV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigntemplatepackageEditObjectV1Request(objEzsigntemplatepackage EzsigntemplatepackageRequestCompound) *EzsigntemplatepackageEditObjectV1Request {
	this := EzsigntemplatepackageEditObjectV1Request{}
	this.ObjEzsigntemplatepackage = objEzsigntemplatepackage
	return &this
}

// NewEzsigntemplatepackageEditObjectV1RequestWithDefaults instantiates a new EzsigntemplatepackageEditObjectV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigntemplatepackageEditObjectV1RequestWithDefaults() *EzsigntemplatepackageEditObjectV1Request {
	this := EzsigntemplatepackageEditObjectV1Request{}
	return &this
}

// GetObjEzsigntemplatepackage returns the ObjEzsigntemplatepackage field value
func (o *EzsigntemplatepackageEditObjectV1Request) GetObjEzsigntemplatepackage() EzsigntemplatepackageRequestCompound {
	if o == nil {
		var ret EzsigntemplatepackageRequestCompound
		return ret
	}

	return o.ObjEzsigntemplatepackage
}

// GetObjEzsigntemplatepackageOk returns a tuple with the ObjEzsigntemplatepackage field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepackageEditObjectV1Request) GetObjEzsigntemplatepackageOk() (*EzsigntemplatepackageRequestCompound, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjEzsigntemplatepackage, true
}

// SetObjEzsigntemplatepackage sets field value
func (o *EzsigntemplatepackageEditObjectV1Request) SetObjEzsigntemplatepackage(v EzsigntemplatepackageRequestCompound) {
	o.ObjEzsigntemplatepackage = v
}

func (o EzsigntemplatepackageEditObjectV1Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigntemplatepackageEditObjectV1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objEzsigntemplatepackage"] = o.ObjEzsigntemplatepackage
	return toSerialize, nil
}

type NullableEzsigntemplatepackageEditObjectV1Request struct {
	value *EzsigntemplatepackageEditObjectV1Request
	isSet bool
}

func (v NullableEzsigntemplatepackageEditObjectV1Request) Get() *EzsigntemplatepackageEditObjectV1Request {
	return v.value
}

func (v *NullableEzsigntemplatepackageEditObjectV1Request) Set(val *EzsigntemplatepackageEditObjectV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigntemplatepackageEditObjectV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigntemplatepackageEditObjectV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigntemplatepackageEditObjectV1Request(val *EzsigntemplatepackageEditObjectV1Request) *NullableEzsigntemplatepackageEditObjectV1Request {
	return &NullableEzsigntemplatepackageEditObjectV1Request{value: val, isSet: true}
}

func (v NullableEzsigntemplatepackageEditObjectV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigntemplatepackageEditObjectV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

