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

// checks if the EzsignfoldersignerassociationCreateObjectV1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignfoldersignerassociationCreateObjectV1Request{}

// EzsignfoldersignerassociationCreateObjectV1Request Request for POST /1/object/ezsignfoldersignerassociation
type EzsignfoldersignerassociationCreateObjectV1Request struct {
	ObjEzsignfoldersignerassociation *EzsignfoldersignerassociationRequest `json:"objEzsignfoldersignerassociation,omitempty"`
	ObjEzsignfoldersignerassociationCompound *EzsignfoldersignerassociationRequestCompound `json:"objEzsignfoldersignerassociationCompound,omitempty"`
}

// NewEzsignfoldersignerassociationCreateObjectV1Request instantiates a new EzsignfoldersignerassociationCreateObjectV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignfoldersignerassociationCreateObjectV1Request() *EzsignfoldersignerassociationCreateObjectV1Request {
	this := EzsignfoldersignerassociationCreateObjectV1Request{}
	return &this
}

// NewEzsignfoldersignerassociationCreateObjectV1RequestWithDefaults instantiates a new EzsignfoldersignerassociationCreateObjectV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignfoldersignerassociationCreateObjectV1RequestWithDefaults() *EzsignfoldersignerassociationCreateObjectV1Request {
	this := EzsignfoldersignerassociationCreateObjectV1Request{}
	return &this
}

// GetObjEzsignfoldersignerassociation returns the ObjEzsignfoldersignerassociation field value if set, zero value otherwise.
func (o *EzsignfoldersignerassociationCreateObjectV1Request) GetObjEzsignfoldersignerassociation() EzsignfoldersignerassociationRequest {
	if o == nil || IsNil(o.ObjEzsignfoldersignerassociation) {
		var ret EzsignfoldersignerassociationRequest
		return ret
	}
	return *o.ObjEzsignfoldersignerassociation
}

// GetObjEzsignfoldersignerassociationOk returns a tuple with the ObjEzsignfoldersignerassociation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignfoldersignerassociationCreateObjectV1Request) GetObjEzsignfoldersignerassociationOk() (*EzsignfoldersignerassociationRequest, bool) {
	if o == nil || IsNil(o.ObjEzsignfoldersignerassociation) {
		return nil, false
	}
	return o.ObjEzsignfoldersignerassociation, true
}

// HasObjEzsignfoldersignerassociation returns a boolean if a field has been set.
func (o *EzsignfoldersignerassociationCreateObjectV1Request) HasObjEzsignfoldersignerassociation() bool {
	if o != nil && !IsNil(o.ObjEzsignfoldersignerassociation) {
		return true
	}

	return false
}

// SetObjEzsignfoldersignerassociation gets a reference to the given EzsignfoldersignerassociationRequest and assigns it to the ObjEzsignfoldersignerassociation field.
func (o *EzsignfoldersignerassociationCreateObjectV1Request) SetObjEzsignfoldersignerassociation(v EzsignfoldersignerassociationRequest) {
	o.ObjEzsignfoldersignerassociation = &v
}

// GetObjEzsignfoldersignerassociationCompound returns the ObjEzsignfoldersignerassociationCompound field value if set, zero value otherwise.
func (o *EzsignfoldersignerassociationCreateObjectV1Request) GetObjEzsignfoldersignerassociationCompound() EzsignfoldersignerassociationRequestCompound {
	if o == nil || IsNil(o.ObjEzsignfoldersignerassociationCompound) {
		var ret EzsignfoldersignerassociationRequestCompound
		return ret
	}
	return *o.ObjEzsignfoldersignerassociationCompound
}

// GetObjEzsignfoldersignerassociationCompoundOk returns a tuple with the ObjEzsignfoldersignerassociationCompound field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignfoldersignerassociationCreateObjectV1Request) GetObjEzsignfoldersignerassociationCompoundOk() (*EzsignfoldersignerassociationRequestCompound, bool) {
	if o == nil || IsNil(o.ObjEzsignfoldersignerassociationCompound) {
		return nil, false
	}
	return o.ObjEzsignfoldersignerassociationCompound, true
}

// HasObjEzsignfoldersignerassociationCompound returns a boolean if a field has been set.
func (o *EzsignfoldersignerassociationCreateObjectV1Request) HasObjEzsignfoldersignerassociationCompound() bool {
	if o != nil && !IsNil(o.ObjEzsignfoldersignerassociationCompound) {
		return true
	}

	return false
}

// SetObjEzsignfoldersignerassociationCompound gets a reference to the given EzsignfoldersignerassociationRequestCompound and assigns it to the ObjEzsignfoldersignerassociationCompound field.
func (o *EzsignfoldersignerassociationCreateObjectV1Request) SetObjEzsignfoldersignerassociationCompound(v EzsignfoldersignerassociationRequestCompound) {
	o.ObjEzsignfoldersignerassociationCompound = &v
}

func (o EzsignfoldersignerassociationCreateObjectV1Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignfoldersignerassociationCreateObjectV1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ObjEzsignfoldersignerassociation) {
		toSerialize["objEzsignfoldersignerassociation"] = o.ObjEzsignfoldersignerassociation
	}
	if !IsNil(o.ObjEzsignfoldersignerassociationCompound) {
		toSerialize["objEzsignfoldersignerassociationCompound"] = o.ObjEzsignfoldersignerassociationCompound
	}
	return toSerialize, nil
}

type NullableEzsignfoldersignerassociationCreateObjectV1Request struct {
	value *EzsignfoldersignerassociationCreateObjectV1Request
	isSet bool
}

func (v NullableEzsignfoldersignerassociationCreateObjectV1Request) Get() *EzsignfoldersignerassociationCreateObjectV1Request {
	return v.value
}

func (v *NullableEzsignfoldersignerassociationCreateObjectV1Request) Set(val *EzsignfoldersignerassociationCreateObjectV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignfoldersignerassociationCreateObjectV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignfoldersignerassociationCreateObjectV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignfoldersignerassociationCreateObjectV1Request(val *EzsignfoldersignerassociationCreateObjectV1Request) *NullableEzsignfoldersignerassociationCreateObjectV1Request {
	return &NullableEzsignfoldersignerassociationCreateObjectV1Request{value: val, isSet: true}
}

func (v NullableEzsignfoldersignerassociationCreateObjectV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignfoldersignerassociationCreateObjectV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


