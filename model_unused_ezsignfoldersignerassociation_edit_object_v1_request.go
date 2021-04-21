/*
 * eZmax API Definition
 *
 * This API expose all the functionnalities for the eZmax and eZsign applications.
 *
 * API version: 1.0.41
 * Contact: support-api@ezmax.ca
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
)

// UNUSEDEzsignfoldersignerassociationEditObjectV1Request Request for the /1/object/ezsignfoldersignerassociation/editObject API Request
type UNUSEDEzsignfoldersignerassociationEditObjectV1Request struct {
	ObjEzsignfoldersignerassociation *EzsignfoldersignerassociationRequest `json:"objEzsignfoldersignerassociation,omitempty"`
}

// NewUNUSEDEzsignfoldersignerassociationEditObjectV1Request instantiates a new UNUSEDEzsignfoldersignerassociationEditObjectV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUNUSEDEzsignfoldersignerassociationEditObjectV1Request() *UNUSEDEzsignfoldersignerassociationEditObjectV1Request {
	this := UNUSEDEzsignfoldersignerassociationEditObjectV1Request{}
	return &this
}

// NewUNUSEDEzsignfoldersignerassociationEditObjectV1RequestWithDefaults instantiates a new UNUSEDEzsignfoldersignerassociationEditObjectV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUNUSEDEzsignfoldersignerassociationEditObjectV1RequestWithDefaults() *UNUSEDEzsignfoldersignerassociationEditObjectV1Request {
	this := UNUSEDEzsignfoldersignerassociationEditObjectV1Request{}
	return &this
}

// GetObjEzsignfoldersignerassociation returns the ObjEzsignfoldersignerassociation field value if set, zero value otherwise.
func (o *UNUSEDEzsignfoldersignerassociationEditObjectV1Request) GetObjEzsignfoldersignerassociation() EzsignfoldersignerassociationRequest {
	if o == nil || o.ObjEzsignfoldersignerassociation == nil {
		var ret EzsignfoldersignerassociationRequest
		return ret
	}
	return *o.ObjEzsignfoldersignerassociation
}

// GetObjEzsignfoldersignerassociationOk returns a tuple with the ObjEzsignfoldersignerassociation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UNUSEDEzsignfoldersignerassociationEditObjectV1Request) GetObjEzsignfoldersignerassociationOk() (*EzsignfoldersignerassociationRequest, bool) {
	if o == nil || o.ObjEzsignfoldersignerassociation == nil {
		return nil, false
	}
	return o.ObjEzsignfoldersignerassociation, true
}

// HasObjEzsignfoldersignerassociation returns a boolean if a field has been set.
func (o *UNUSEDEzsignfoldersignerassociationEditObjectV1Request) HasObjEzsignfoldersignerassociation() bool {
	if o != nil && o.ObjEzsignfoldersignerassociation != nil {
		return true
	}

	return false
}

// SetObjEzsignfoldersignerassociation gets a reference to the given EzsignfoldersignerassociationRequest and assigns it to the ObjEzsignfoldersignerassociation field.
func (o *UNUSEDEzsignfoldersignerassociationEditObjectV1Request) SetObjEzsignfoldersignerassociation(v EzsignfoldersignerassociationRequest) {
	o.ObjEzsignfoldersignerassociation = &v
}

func (o UNUSEDEzsignfoldersignerassociationEditObjectV1Request) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ObjEzsignfoldersignerassociation != nil {
		toSerialize["objEzsignfoldersignerassociation"] = o.ObjEzsignfoldersignerassociation
	}
	return json.Marshal(toSerialize)
}

type NullableUNUSEDEzsignfoldersignerassociationEditObjectV1Request struct {
	value *UNUSEDEzsignfoldersignerassociationEditObjectV1Request
	isSet bool
}

func (v NullableUNUSEDEzsignfoldersignerassociationEditObjectV1Request) Get() *UNUSEDEzsignfoldersignerassociationEditObjectV1Request {
	return v.value
}

func (v *NullableUNUSEDEzsignfoldersignerassociationEditObjectV1Request) Set(val *UNUSEDEzsignfoldersignerassociationEditObjectV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableUNUSEDEzsignfoldersignerassociationEditObjectV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableUNUSEDEzsignfoldersignerassociationEditObjectV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUNUSEDEzsignfoldersignerassociationEditObjectV1Request(val *UNUSEDEzsignfoldersignerassociationEditObjectV1Request) *NullableUNUSEDEzsignfoldersignerassociationEditObjectV1Request {
	return &NullableUNUSEDEzsignfoldersignerassociationEditObjectV1Request{value: val, isSet: true}
}

func (v NullableUNUSEDEzsignfoldersignerassociationEditObjectV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUNUSEDEzsignfoldersignerassociationEditObjectV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


