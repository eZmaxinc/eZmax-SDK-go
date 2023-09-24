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

// checks if the EzsignformfieldgroupsignerRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignformfieldgroupsignerRequest{}

// EzsignformfieldgroupsignerRequest A Ezsignformfieldgroupsigner Object
type EzsignformfieldgroupsignerRequest struct {
	// The unique ID of the Ezsignformfieldgroupsigner
	PkiEzsignformfieldgroupsignerID *int32 `json:"pkiEzsignformfieldgroupsignerID,omitempty"`
	// The unique ID of the Ezsignfoldersignerassociation
	FkiEzsignfoldersignerassociationID int32 `json:"fkiEzsignfoldersignerassociationID"`
}

// NewEzsignformfieldgroupsignerRequest instantiates a new EzsignformfieldgroupsignerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignformfieldgroupsignerRequest(fkiEzsignfoldersignerassociationID int32) *EzsignformfieldgroupsignerRequest {
	this := EzsignformfieldgroupsignerRequest{}
	this.FkiEzsignfoldersignerassociationID = fkiEzsignfoldersignerassociationID
	return &this
}

// NewEzsignformfieldgroupsignerRequestWithDefaults instantiates a new EzsignformfieldgroupsignerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignformfieldgroupsignerRequestWithDefaults() *EzsignformfieldgroupsignerRequest {
	this := EzsignformfieldgroupsignerRequest{}
	return &this
}

// GetPkiEzsignformfieldgroupsignerID returns the PkiEzsignformfieldgroupsignerID field value if set, zero value otherwise.
func (o *EzsignformfieldgroupsignerRequest) GetPkiEzsignformfieldgroupsignerID() int32 {
	if o == nil || IsNil(o.PkiEzsignformfieldgroupsignerID) {
		var ret int32
		return ret
	}
	return *o.PkiEzsignformfieldgroupsignerID
}

// GetPkiEzsignformfieldgroupsignerIDOk returns a tuple with the PkiEzsignformfieldgroupsignerID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignformfieldgroupsignerRequest) GetPkiEzsignformfieldgroupsignerIDOk() (*int32, bool) {
	if o == nil || IsNil(o.PkiEzsignformfieldgroupsignerID) {
		return nil, false
	}
	return o.PkiEzsignformfieldgroupsignerID, true
}

// HasPkiEzsignformfieldgroupsignerID returns a boolean if a field has been set.
func (o *EzsignformfieldgroupsignerRequest) HasPkiEzsignformfieldgroupsignerID() bool {
	if o != nil && !IsNil(o.PkiEzsignformfieldgroupsignerID) {
		return true
	}

	return false
}

// SetPkiEzsignformfieldgroupsignerID gets a reference to the given int32 and assigns it to the PkiEzsignformfieldgroupsignerID field.
func (o *EzsignformfieldgroupsignerRequest) SetPkiEzsignformfieldgroupsignerID(v int32) {
	o.PkiEzsignformfieldgroupsignerID = &v
}

// GetFkiEzsignfoldersignerassociationID returns the FkiEzsignfoldersignerassociationID field value
func (o *EzsignformfieldgroupsignerRequest) GetFkiEzsignfoldersignerassociationID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzsignfoldersignerassociationID
}

// GetFkiEzsignfoldersignerassociationIDOk returns a tuple with the FkiEzsignfoldersignerassociationID field value
// and a boolean to check if the value has been set.
func (o *EzsignformfieldgroupsignerRequest) GetFkiEzsignfoldersignerassociationIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiEzsignfoldersignerassociationID, true
}

// SetFkiEzsignfoldersignerassociationID sets field value
func (o *EzsignformfieldgroupsignerRequest) SetFkiEzsignfoldersignerassociationID(v int32) {
	o.FkiEzsignfoldersignerassociationID = v
}

func (o EzsignformfieldgroupsignerRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignformfieldgroupsignerRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PkiEzsignformfieldgroupsignerID) {
		toSerialize["pkiEzsignformfieldgroupsignerID"] = o.PkiEzsignformfieldgroupsignerID
	}
	toSerialize["fkiEzsignfoldersignerassociationID"] = o.FkiEzsignfoldersignerassociationID
	return toSerialize, nil
}

type NullableEzsignformfieldgroupsignerRequest struct {
	value *EzsignformfieldgroupsignerRequest
	isSet bool
}

func (v NullableEzsignformfieldgroupsignerRequest) Get() *EzsignformfieldgroupsignerRequest {
	return v.value
}

func (v *NullableEzsignformfieldgroupsignerRequest) Set(val *EzsignformfieldgroupsignerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignformfieldgroupsignerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignformfieldgroupsignerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignformfieldgroupsignerRequest(val *EzsignformfieldgroupsignerRequest) *NullableEzsignformfieldgroupsignerRequest {
	return &NullableEzsignformfieldgroupsignerRequest{value: val, isSet: true}
}

func (v NullableEzsignformfieldgroupsignerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignformfieldgroupsignerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

