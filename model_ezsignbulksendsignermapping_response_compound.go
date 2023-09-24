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

// checks if the EzsignbulksendsignermappingResponseCompound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignbulksendsignermappingResponseCompound{}

// EzsignbulksendsignermappingResponseCompound A Ezsignbulksendsignermapping Object
type EzsignbulksendsignermappingResponseCompound struct {
	// The unique ID of the Ezsignbulksendsignermapping
	PkiEzsignbulksendsignermappingID int32 `json:"pkiEzsignbulksendsignermappingID"`
	// The unique ID of the Ezsignbulksend
	FkiEzsignbulksendID int32 `json:"fkiEzsignbulksendID"`
	// The unique ID of the User
	FkiUserID *int32 `json:"fkiUserID,omitempty"`
	// The description of the Ezsignbulksendsignermapping
	SEzsignbulksendsignermappingDescription string `json:"sEzsignbulksendsignermappingDescription"`
}

// NewEzsignbulksendsignermappingResponseCompound instantiates a new EzsignbulksendsignermappingResponseCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignbulksendsignermappingResponseCompound(pkiEzsignbulksendsignermappingID int32, fkiEzsignbulksendID int32, sEzsignbulksendsignermappingDescription string) *EzsignbulksendsignermappingResponseCompound {
	this := EzsignbulksendsignermappingResponseCompound{}
	this.PkiEzsignbulksendsignermappingID = pkiEzsignbulksendsignermappingID
	this.FkiEzsignbulksendID = fkiEzsignbulksendID
	this.SEzsignbulksendsignermappingDescription = sEzsignbulksendsignermappingDescription
	return &this
}

// NewEzsignbulksendsignermappingResponseCompoundWithDefaults instantiates a new EzsignbulksendsignermappingResponseCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignbulksendsignermappingResponseCompoundWithDefaults() *EzsignbulksendsignermappingResponseCompound {
	this := EzsignbulksendsignermappingResponseCompound{}
	return &this
}

// GetPkiEzsignbulksendsignermappingID returns the PkiEzsignbulksendsignermappingID field value
func (o *EzsignbulksendsignermappingResponseCompound) GetPkiEzsignbulksendsignermappingID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiEzsignbulksendsignermappingID
}

// GetPkiEzsignbulksendsignermappingIDOk returns a tuple with the PkiEzsignbulksendsignermappingID field value
// and a boolean to check if the value has been set.
func (o *EzsignbulksendsignermappingResponseCompound) GetPkiEzsignbulksendsignermappingIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiEzsignbulksendsignermappingID, true
}

// SetPkiEzsignbulksendsignermappingID sets field value
func (o *EzsignbulksendsignermappingResponseCompound) SetPkiEzsignbulksendsignermappingID(v int32) {
	o.PkiEzsignbulksendsignermappingID = v
}

// GetFkiEzsignbulksendID returns the FkiEzsignbulksendID field value
func (o *EzsignbulksendsignermappingResponseCompound) GetFkiEzsignbulksendID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzsignbulksendID
}

// GetFkiEzsignbulksendIDOk returns a tuple with the FkiEzsignbulksendID field value
// and a boolean to check if the value has been set.
func (o *EzsignbulksendsignermappingResponseCompound) GetFkiEzsignbulksendIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiEzsignbulksendID, true
}

// SetFkiEzsignbulksendID sets field value
func (o *EzsignbulksendsignermappingResponseCompound) SetFkiEzsignbulksendID(v int32) {
	o.FkiEzsignbulksendID = v
}

// GetFkiUserID returns the FkiUserID field value if set, zero value otherwise.
func (o *EzsignbulksendsignermappingResponseCompound) GetFkiUserID() int32 {
	if o == nil || IsNil(o.FkiUserID) {
		var ret int32
		return ret
	}
	return *o.FkiUserID
}

// GetFkiUserIDOk returns a tuple with the FkiUserID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignbulksendsignermappingResponseCompound) GetFkiUserIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiUserID) {
		return nil, false
	}
	return o.FkiUserID, true
}

// HasFkiUserID returns a boolean if a field has been set.
func (o *EzsignbulksendsignermappingResponseCompound) HasFkiUserID() bool {
	if o != nil && !IsNil(o.FkiUserID) {
		return true
	}

	return false
}

// SetFkiUserID gets a reference to the given int32 and assigns it to the FkiUserID field.
func (o *EzsignbulksendsignermappingResponseCompound) SetFkiUserID(v int32) {
	o.FkiUserID = &v
}

// GetSEzsignbulksendsignermappingDescription returns the SEzsignbulksendsignermappingDescription field value
func (o *EzsignbulksendsignermappingResponseCompound) GetSEzsignbulksendsignermappingDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzsignbulksendsignermappingDescription
}

// GetSEzsignbulksendsignermappingDescriptionOk returns a tuple with the SEzsignbulksendsignermappingDescription field value
// and a boolean to check if the value has been set.
func (o *EzsignbulksendsignermappingResponseCompound) GetSEzsignbulksendsignermappingDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SEzsignbulksendsignermappingDescription, true
}

// SetSEzsignbulksendsignermappingDescription sets field value
func (o *EzsignbulksendsignermappingResponseCompound) SetSEzsignbulksendsignermappingDescription(v string) {
	o.SEzsignbulksendsignermappingDescription = v
}

func (o EzsignbulksendsignermappingResponseCompound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignbulksendsignermappingResponseCompound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiEzsignbulksendsignermappingID"] = o.PkiEzsignbulksendsignermappingID
	toSerialize["fkiEzsignbulksendID"] = o.FkiEzsignbulksendID
	if !IsNil(o.FkiUserID) {
		toSerialize["fkiUserID"] = o.FkiUserID
	}
	toSerialize["sEzsignbulksendsignermappingDescription"] = o.SEzsignbulksendsignermappingDescription
	return toSerialize, nil
}

type NullableEzsignbulksendsignermappingResponseCompound struct {
	value *EzsignbulksendsignermappingResponseCompound
	isSet bool
}

func (v NullableEzsignbulksendsignermappingResponseCompound) Get() *EzsignbulksendsignermappingResponseCompound {
	return v.value
}

func (v *NullableEzsignbulksendsignermappingResponseCompound) Set(val *EzsignbulksendsignermappingResponseCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignbulksendsignermappingResponseCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignbulksendsignermappingResponseCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignbulksendsignermappingResponseCompound(val *EzsignbulksendsignermappingResponseCompound) *NullableEzsignbulksendsignermappingResponseCompound {
	return &NullableEzsignbulksendsignermappingResponseCompound{value: val, isSet: true}
}

func (v NullableEzsignbulksendsignermappingResponseCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignbulksendsignermappingResponseCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

