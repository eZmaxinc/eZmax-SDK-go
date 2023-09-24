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

// checks if the EzsignfoldersignerassociationEditObjectV1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignfoldersignerassociationEditObjectV1Request{}

// EzsignfoldersignerassociationEditObjectV1Request Request for PUT /1/object/ezsignfoldersignerassociation/{pkiEzsignfoldersignerassociationID}
type EzsignfoldersignerassociationEditObjectV1Request struct {
	ObjEzsignfoldersignerassociation EzsignfoldersignerassociationRequestCompound `json:"objEzsignfoldersignerassociation"`
}

// NewEzsignfoldersignerassociationEditObjectV1Request instantiates a new EzsignfoldersignerassociationEditObjectV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignfoldersignerassociationEditObjectV1Request(objEzsignfoldersignerassociation EzsignfoldersignerassociationRequestCompound) *EzsignfoldersignerassociationEditObjectV1Request {
	this := EzsignfoldersignerassociationEditObjectV1Request{}
	this.ObjEzsignfoldersignerassociation = objEzsignfoldersignerassociation
	return &this
}

// NewEzsignfoldersignerassociationEditObjectV1RequestWithDefaults instantiates a new EzsignfoldersignerassociationEditObjectV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignfoldersignerassociationEditObjectV1RequestWithDefaults() *EzsignfoldersignerassociationEditObjectV1Request {
	this := EzsignfoldersignerassociationEditObjectV1Request{}
	return &this
}

// GetObjEzsignfoldersignerassociation returns the ObjEzsignfoldersignerassociation field value
func (o *EzsignfoldersignerassociationEditObjectV1Request) GetObjEzsignfoldersignerassociation() EzsignfoldersignerassociationRequestCompound {
	if o == nil {
		var ret EzsignfoldersignerassociationRequestCompound
		return ret
	}

	return o.ObjEzsignfoldersignerassociation
}

// GetObjEzsignfoldersignerassociationOk returns a tuple with the ObjEzsignfoldersignerassociation field value
// and a boolean to check if the value has been set.
func (o *EzsignfoldersignerassociationEditObjectV1Request) GetObjEzsignfoldersignerassociationOk() (*EzsignfoldersignerassociationRequestCompound, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjEzsignfoldersignerassociation, true
}

// SetObjEzsignfoldersignerassociation sets field value
func (o *EzsignfoldersignerassociationEditObjectV1Request) SetObjEzsignfoldersignerassociation(v EzsignfoldersignerassociationRequestCompound) {
	o.ObjEzsignfoldersignerassociation = v
}

func (o EzsignfoldersignerassociationEditObjectV1Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignfoldersignerassociationEditObjectV1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objEzsignfoldersignerassociation"] = o.ObjEzsignfoldersignerassociation
	return toSerialize, nil
}

type NullableEzsignfoldersignerassociationEditObjectV1Request struct {
	value *EzsignfoldersignerassociationEditObjectV1Request
	isSet bool
}

func (v NullableEzsignfoldersignerassociationEditObjectV1Request) Get() *EzsignfoldersignerassociationEditObjectV1Request {
	return v.value
}

func (v *NullableEzsignfoldersignerassociationEditObjectV1Request) Set(val *EzsignfoldersignerassociationEditObjectV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignfoldersignerassociationEditObjectV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignfoldersignerassociationEditObjectV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignfoldersignerassociationEditObjectV1Request(val *EzsignfoldersignerassociationEditObjectV1Request) *NullableEzsignfoldersignerassociationEditObjectV1Request {
	return &NullableEzsignfoldersignerassociationEditObjectV1Request{value: val, isSet: true}
}

func (v NullableEzsignfoldersignerassociationEditObjectV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignfoldersignerassociationEditObjectV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

