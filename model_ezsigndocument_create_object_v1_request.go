/*
eZmax API Definition

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.1.0
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
)

// EzsigndocumentCreateObjectV1Request Request for the /1/object/ezsigndocument/createObject API Request
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
	if o == nil || o.ObjEzsigndocument == nil {
		var ret EzsigndocumentRequest
		return ret
	}
	return *o.ObjEzsigndocument
}

// GetObjEzsigndocumentOk returns a tuple with the ObjEzsigndocument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigndocumentCreateObjectV1Request) GetObjEzsigndocumentOk() (*EzsigndocumentRequest, bool) {
	if o == nil || o.ObjEzsigndocument == nil {
		return nil, false
	}
	return o.ObjEzsigndocument, true
}

// HasObjEzsigndocument returns a boolean if a field has been set.
func (o *EzsigndocumentCreateObjectV1Request) HasObjEzsigndocument() bool {
	if o != nil && o.ObjEzsigndocument != nil {
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
	if o == nil || o.ObjEzsigndocumentCompound == nil {
		var ret EzsigndocumentRequestCompound
		return ret
	}
	return *o.ObjEzsigndocumentCompound
}

// GetObjEzsigndocumentCompoundOk returns a tuple with the ObjEzsigndocumentCompound field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigndocumentCreateObjectV1Request) GetObjEzsigndocumentCompoundOk() (*EzsigndocumentRequestCompound, bool) {
	if o == nil || o.ObjEzsigndocumentCompound == nil {
		return nil, false
	}
	return o.ObjEzsigndocumentCompound, true
}

// HasObjEzsigndocumentCompound returns a boolean if a field has been set.
func (o *EzsigndocumentCreateObjectV1Request) HasObjEzsigndocumentCompound() bool {
	if o != nil && o.ObjEzsigndocumentCompound != nil {
		return true
	}

	return false
}

// SetObjEzsigndocumentCompound gets a reference to the given EzsigndocumentRequestCompound and assigns it to the ObjEzsigndocumentCompound field.
func (o *EzsigndocumentCreateObjectV1Request) SetObjEzsigndocumentCompound(v EzsigndocumentRequestCompound) {
	o.ObjEzsigndocumentCompound = &v
}

func (o EzsigndocumentCreateObjectV1Request) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ObjEzsigndocument != nil {
		toSerialize["objEzsigndocument"] = o.ObjEzsigndocument
	}
	if o.ObjEzsigndocumentCompound != nil {
		toSerialize["objEzsigndocumentCompound"] = o.ObjEzsigndocumentCompound
	}
	return json.Marshal(toSerialize)
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


