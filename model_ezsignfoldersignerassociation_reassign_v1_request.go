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

// checks if the EzsignfoldersignerassociationReassignV1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignfoldersignerassociationReassignV1Request{}

// EzsignfoldersignerassociationReassignV1Request Request for POST /1/object/ezsignfoldersignerassociation/{pkiEzsignfoldersignerassociationID}/reassign
type EzsignfoldersignerassociationReassignV1Request struct {
	// The unique ID of the Ezsignfoldersignerassociation
	FkiEzsignfoldersignerassociationID *int32 `json:"fkiEzsignfoldersignerassociationID,omitempty"`
}

// NewEzsignfoldersignerassociationReassignV1Request instantiates a new EzsignfoldersignerassociationReassignV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignfoldersignerassociationReassignV1Request() *EzsignfoldersignerassociationReassignV1Request {
	this := EzsignfoldersignerassociationReassignV1Request{}
	return &this
}

// NewEzsignfoldersignerassociationReassignV1RequestWithDefaults instantiates a new EzsignfoldersignerassociationReassignV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignfoldersignerassociationReassignV1RequestWithDefaults() *EzsignfoldersignerassociationReassignV1Request {
	this := EzsignfoldersignerassociationReassignV1Request{}
	return &this
}

// GetFkiEzsignfoldersignerassociationID returns the FkiEzsignfoldersignerassociationID field value if set, zero value otherwise.
func (o *EzsignfoldersignerassociationReassignV1Request) GetFkiEzsignfoldersignerassociationID() int32 {
	if o == nil || IsNil(o.FkiEzsignfoldersignerassociationID) {
		var ret int32
		return ret
	}
	return *o.FkiEzsignfoldersignerassociationID
}

// GetFkiEzsignfoldersignerassociationIDOk returns a tuple with the FkiEzsignfoldersignerassociationID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignfoldersignerassociationReassignV1Request) GetFkiEzsignfoldersignerassociationIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiEzsignfoldersignerassociationID) {
		return nil, false
	}
	return o.FkiEzsignfoldersignerassociationID, true
}

// HasFkiEzsignfoldersignerassociationID returns a boolean if a field has been set.
func (o *EzsignfoldersignerassociationReassignV1Request) HasFkiEzsignfoldersignerassociationID() bool {
	if o != nil && !IsNil(o.FkiEzsignfoldersignerassociationID) {
		return true
	}

	return false
}

// SetFkiEzsignfoldersignerassociationID gets a reference to the given int32 and assigns it to the FkiEzsignfoldersignerassociationID field.
func (o *EzsignfoldersignerassociationReassignV1Request) SetFkiEzsignfoldersignerassociationID(v int32) {
	o.FkiEzsignfoldersignerassociationID = &v
}

func (o EzsignfoldersignerassociationReassignV1Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignfoldersignerassociationReassignV1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FkiEzsignfoldersignerassociationID) {
		toSerialize["fkiEzsignfoldersignerassociationID"] = o.FkiEzsignfoldersignerassociationID
	}
	return toSerialize, nil
}

type NullableEzsignfoldersignerassociationReassignV1Request struct {
	value *EzsignfoldersignerassociationReassignV1Request
	isSet bool
}

func (v NullableEzsignfoldersignerassociationReassignV1Request) Get() *EzsignfoldersignerassociationReassignV1Request {
	return v.value
}

func (v *NullableEzsignfoldersignerassociationReassignV1Request) Set(val *EzsignfoldersignerassociationReassignV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignfoldersignerassociationReassignV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignfoldersignerassociationReassignV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignfoldersignerassociationReassignV1Request(val *EzsignfoldersignerassociationReassignV1Request) *NullableEzsignfoldersignerassociationReassignV1Request {
	return &NullableEzsignfoldersignerassociationReassignV1Request{value: val, isSet: true}
}

func (v NullableEzsignfoldersignerassociationReassignV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignfoldersignerassociationReassignV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


