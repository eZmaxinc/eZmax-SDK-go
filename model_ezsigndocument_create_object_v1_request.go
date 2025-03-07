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
)

// checks if the EzsigndocumentCreateObjectV1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigndocumentCreateObjectV1Request{}

// EzsigndocumentCreateObjectV1Request Request for POST /1/object/ezsigndocument
type EzsigndocumentCreateObjectV1Request struct {
	ObjEzsigndocument *EzsigndocumentRequest `json:"objEzsigndocument,omitempty"`
	ObjEzsigndocumentCompound *EzsigndocumentRequestCompound `json:"objEzsigndocumentCompound,omitempty"`
}

// NewEzsigndocumentCreateObjectV1Request instantiates a new EzsigndocumentCreateObjectV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigndocumentCreateObjectV1Request() *EzsigndocumentCreateObjectV1Request {
	this := EzsigndocumentCreateObjectV1Request{}
	return &this
}

// NewEzsigndocumentCreateObjectV1RequestWithDefaults instantiates a new EzsigndocumentCreateObjectV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigndocumentCreateObjectV1RequestWithDefaults() *EzsigndocumentCreateObjectV1Request {
	this := EzsigndocumentCreateObjectV1Request{}
	return &this
}

// GetObjEzsigndocument returns the ObjEzsigndocument field value if set, zero value otherwise.
func (o *EzsigndocumentCreateObjectV1Request) GetObjEzsigndocument() EzsigndocumentRequest {
	if o == nil || IsNil(o.ObjEzsigndocument) {
		var ret EzsigndocumentRequest
		return ret
	}
	return *o.ObjEzsigndocument
}

// GetObjEzsigndocumentOk returns a tuple with the ObjEzsigndocument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigndocumentCreateObjectV1Request) GetObjEzsigndocumentOk() (*EzsigndocumentRequest, bool) {
	if o == nil || IsNil(o.ObjEzsigndocument) {
		return nil, false
	}
	return o.ObjEzsigndocument, true
}

// HasObjEzsigndocument returns a boolean if a field has been set.
func (o *EzsigndocumentCreateObjectV1Request) HasObjEzsigndocument() bool {
	if o != nil && !IsNil(o.ObjEzsigndocument) {
		return true
	}

	return false
}

// SetObjEzsigndocument gets a reference to the given EzsigndocumentRequest and assigns it to the ObjEzsigndocument field.
func (o *EzsigndocumentCreateObjectV1Request) SetObjEzsigndocument(v EzsigndocumentRequest) {
	o.ObjEzsigndocument = &v
}

// GetObjEzsigndocumentCompound returns the ObjEzsigndocumentCompound field value if set, zero value otherwise.
func (o *EzsigndocumentCreateObjectV1Request) GetObjEzsigndocumentCompound() EzsigndocumentRequestCompound {
	if o == nil || IsNil(o.ObjEzsigndocumentCompound) {
		var ret EzsigndocumentRequestCompound
		return ret
	}
	return *o.ObjEzsigndocumentCompound
}

// GetObjEzsigndocumentCompoundOk returns a tuple with the ObjEzsigndocumentCompound field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigndocumentCreateObjectV1Request) GetObjEzsigndocumentCompoundOk() (*EzsigndocumentRequestCompound, bool) {
	if o == nil || IsNil(o.ObjEzsigndocumentCompound) {
		return nil, false
	}
	return o.ObjEzsigndocumentCompound, true
}

// HasObjEzsigndocumentCompound returns a boolean if a field has been set.
func (o *EzsigndocumentCreateObjectV1Request) HasObjEzsigndocumentCompound() bool {
	if o != nil && !IsNil(o.ObjEzsigndocumentCompound) {
		return true
	}

	return false
}

// SetObjEzsigndocumentCompound gets a reference to the given EzsigndocumentRequestCompound and assigns it to the ObjEzsigndocumentCompound field.
func (o *EzsigndocumentCreateObjectV1Request) SetObjEzsigndocumentCompound(v EzsigndocumentRequestCompound) {
	o.ObjEzsigndocumentCompound = &v
}

func (o EzsigndocumentCreateObjectV1Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigndocumentCreateObjectV1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ObjEzsigndocument) {
		toSerialize["objEzsigndocument"] = o.ObjEzsigndocument
	}
	if !IsNil(o.ObjEzsigndocumentCompound) {
		toSerialize["objEzsigndocumentCompound"] = o.ObjEzsigndocumentCompound
	}
	return toSerialize, nil
}

type NullableEzsigndocumentCreateObjectV1Request struct {
	value *EzsigndocumentCreateObjectV1Request
	isSet bool
}

func (v NullableEzsigndocumentCreateObjectV1Request) Get() *EzsigndocumentCreateObjectV1Request {
	return v.value
}

func (v *NullableEzsigndocumentCreateObjectV1Request) Set(val *EzsigndocumentCreateObjectV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigndocumentCreateObjectV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigndocumentCreateObjectV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigndocumentCreateObjectV1Request(val *EzsigndocumentCreateObjectV1Request) *NullableEzsigndocumentCreateObjectV1Request {
	return &NullableEzsigndocumentCreateObjectV1Request{value: val, isSet: true}
}

func (v NullableEzsigndocumentCreateObjectV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigndocumentCreateObjectV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


