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

// checks if the EzsignfoldersignerassociationPatchObjectV1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignfoldersignerassociationPatchObjectV1Request{}

// EzsignfoldersignerassociationPatchObjectV1Request Request for PATCH /1/object/ezsignfoldersignerassociation/{pkiEzsignfoldersignerassociationID}
type EzsignfoldersignerassociationPatchObjectV1Request struct {
	ObjEzsignfoldersignerassociation EzsignfoldersignerassociationRequestPatch `json:"objEzsignfoldersignerassociation"`
}

// NewEzsignfoldersignerassociationPatchObjectV1Request instantiates a new EzsignfoldersignerassociationPatchObjectV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignfoldersignerassociationPatchObjectV1Request(objEzsignfoldersignerassociation EzsignfoldersignerassociationRequestPatch) *EzsignfoldersignerassociationPatchObjectV1Request {
	this := EzsignfoldersignerassociationPatchObjectV1Request{}
	this.ObjEzsignfoldersignerassociation = objEzsignfoldersignerassociation
	return &this
}

// NewEzsignfoldersignerassociationPatchObjectV1RequestWithDefaults instantiates a new EzsignfoldersignerassociationPatchObjectV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignfoldersignerassociationPatchObjectV1RequestWithDefaults() *EzsignfoldersignerassociationPatchObjectV1Request {
	this := EzsignfoldersignerassociationPatchObjectV1Request{}
	return &this
}

// GetObjEzsignfoldersignerassociation returns the ObjEzsignfoldersignerassociation field value
func (o *EzsignfoldersignerassociationPatchObjectV1Request) GetObjEzsignfoldersignerassociation() EzsignfoldersignerassociationRequestPatch {
	if o == nil {
		var ret EzsignfoldersignerassociationRequestPatch
		return ret
	}

	return o.ObjEzsignfoldersignerassociation
}

// GetObjEzsignfoldersignerassociationOk returns a tuple with the ObjEzsignfoldersignerassociation field value
// and a boolean to check if the value has been set.
func (o *EzsignfoldersignerassociationPatchObjectV1Request) GetObjEzsignfoldersignerassociationOk() (*EzsignfoldersignerassociationRequestPatch, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjEzsignfoldersignerassociation, true
}

// SetObjEzsignfoldersignerassociation sets field value
func (o *EzsignfoldersignerassociationPatchObjectV1Request) SetObjEzsignfoldersignerassociation(v EzsignfoldersignerassociationRequestPatch) {
	o.ObjEzsignfoldersignerassociation = v
}

func (o EzsignfoldersignerassociationPatchObjectV1Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignfoldersignerassociationPatchObjectV1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objEzsignfoldersignerassociation"] = o.ObjEzsignfoldersignerassociation
	return toSerialize, nil
}

type NullableEzsignfoldersignerassociationPatchObjectV1Request struct {
	value *EzsignfoldersignerassociationPatchObjectV1Request
	isSet bool
}

func (v NullableEzsignfoldersignerassociationPatchObjectV1Request) Get() *EzsignfoldersignerassociationPatchObjectV1Request {
	return v.value
}

func (v *NullableEzsignfoldersignerassociationPatchObjectV1Request) Set(val *EzsignfoldersignerassociationPatchObjectV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignfoldersignerassociationPatchObjectV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignfoldersignerassociationPatchObjectV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignfoldersignerassociationPatchObjectV1Request(val *EzsignfoldersignerassociationPatchObjectV1Request) *NullableEzsignfoldersignerassociationPatchObjectV1Request {
	return &NullableEzsignfoldersignerassociationPatchObjectV1Request{value: val, isSet: true}
}

func (v NullableEzsignfoldersignerassociationPatchObjectV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignfoldersignerassociationPatchObjectV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

