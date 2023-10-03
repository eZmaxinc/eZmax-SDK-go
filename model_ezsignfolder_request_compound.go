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

// checks if the EzsignfolderRequestCompound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignfolderRequestCompound{}

// EzsignfolderRequestCompound An Ezsignfolder Object and children to create a complete structure
type EzsignfolderRequestCompound struct {
	// The unique ID of the Ezsignfolder
	PkiEzsignfolderID *int32 `json:"pkiEzsignfolderID,omitempty"`
	// The unique ID of the Ezsignfoldertype.
	FkiEzsignfoldertypeID int32 `json:"fkiEzsignfoldertypeID"`
	// The unique ID of the Ezsigntsarequirement.  Determine if a Time Stamping Authority should add a timestamp on each of the signature. Valid values:  |Value|Description| |-|-| |1|No. TSA Timestamping will requested. This will make all signatures a lot faster since no round-trip to the TSA server will be required. Timestamping will be made using eZsign server's time.| |2|Best effort. Timestamping from a Time Stamping Authority will be requested but is not mandatory. In the very improbable case it cannot be completed, the timestamping will be made using eZsign server's time. **Additional fee applies**| |3|Mandatory. Timestamping from a Time Stamping Authority will be requested and is mandatory. In the very improbable case it cannot be completed, the signature will fail and the user will be asked to retry. **Additional fee applies**|
	FkiEzsigntsarequirementID *int32 `json:"fkiEzsigntsarequirementID,omitempty"`
	// The description of the Ezsignfolder
	SEzsignfolderDescription string `json:"sEzsignfolderDescription"`
	// Note about the Ezsignfolder
	TEzsignfolderNote *string `json:"tEzsignfolderNote,omitempty"`
	EEzsignfolderSendreminderfrequency FieldEEzsignfolderSendreminderfrequency `json:"eEzsignfolderSendreminderfrequency"`
	// This field can be used to store an External ID from the client's system.  Anything can be stored in this field, it will never be evaluated by the eZmax system and will be returned AS-IS.  To store multiple values, consider using a JSON formatted structure, a URL encoded string, a CSV or any other custom format. 
	SEzsignfolderExternalid *string `json:"sEzsignfolderExternalid,omitempty"`
}

// NewEzsignfolderRequestCompound instantiates a new EzsignfolderRequestCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignfolderRequestCompound(fkiEzsignfoldertypeID int32, sEzsignfolderDescription string, eEzsignfolderSendreminderfrequency FieldEEzsignfolderSendreminderfrequency) *EzsignfolderRequestCompound {
	this := EzsignfolderRequestCompound{}
	this.FkiEzsignfoldertypeID = fkiEzsignfoldertypeID
	this.SEzsignfolderDescription = sEzsignfolderDescription
	this.EEzsignfolderSendreminderfrequency = eEzsignfolderSendreminderfrequency
	return &this
}

// NewEzsignfolderRequestCompoundWithDefaults instantiates a new EzsignfolderRequestCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignfolderRequestCompoundWithDefaults() *EzsignfolderRequestCompound {
	this := EzsignfolderRequestCompound{}
	return &this
}

// GetPkiEzsignfolderID returns the PkiEzsignfolderID field value if set, zero value otherwise.
func (o *EzsignfolderRequestCompound) GetPkiEzsignfolderID() int32 {
	if o == nil || IsNil(o.PkiEzsignfolderID) {
		var ret int32
		return ret
	}
	return *o.PkiEzsignfolderID
}

// GetPkiEzsignfolderIDOk returns a tuple with the PkiEzsignfolderID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignfolderRequestCompound) GetPkiEzsignfolderIDOk() (*int32, bool) {
	if o == nil || IsNil(o.PkiEzsignfolderID) {
		return nil, false
	}
	return o.PkiEzsignfolderID, true
}

// HasPkiEzsignfolderID returns a boolean if a field has been set.
func (o *EzsignfolderRequestCompound) HasPkiEzsignfolderID() bool {
	if o != nil && !IsNil(o.PkiEzsignfolderID) {
		return true
	}

	return false
}

// SetPkiEzsignfolderID gets a reference to the given int32 and assigns it to the PkiEzsignfolderID field.
func (o *EzsignfolderRequestCompound) SetPkiEzsignfolderID(v int32) {
	o.PkiEzsignfolderID = &v
}

// GetFkiEzsignfoldertypeID returns the FkiEzsignfoldertypeID field value
func (o *EzsignfolderRequestCompound) GetFkiEzsignfoldertypeID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzsignfoldertypeID
}

// GetFkiEzsignfoldertypeIDOk returns a tuple with the FkiEzsignfoldertypeID field value
// and a boolean to check if the value has been set.
func (o *EzsignfolderRequestCompound) GetFkiEzsignfoldertypeIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiEzsignfoldertypeID, true
}

// SetFkiEzsignfoldertypeID sets field value
func (o *EzsignfolderRequestCompound) SetFkiEzsignfoldertypeID(v int32) {
	o.FkiEzsignfoldertypeID = v
}

// GetFkiEzsigntsarequirementID returns the FkiEzsigntsarequirementID field value if set, zero value otherwise.
func (o *EzsignfolderRequestCompound) GetFkiEzsigntsarequirementID() int32 {
	if o == nil || IsNil(o.FkiEzsigntsarequirementID) {
		var ret int32
		return ret
	}
	return *o.FkiEzsigntsarequirementID
}

// GetFkiEzsigntsarequirementIDOk returns a tuple with the FkiEzsigntsarequirementID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignfolderRequestCompound) GetFkiEzsigntsarequirementIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiEzsigntsarequirementID) {
		return nil, false
	}
	return o.FkiEzsigntsarequirementID, true
}

// HasFkiEzsigntsarequirementID returns a boolean if a field has been set.
func (o *EzsignfolderRequestCompound) HasFkiEzsigntsarequirementID() bool {
	if o != nil && !IsNil(o.FkiEzsigntsarequirementID) {
		return true
	}

	return false
}

// SetFkiEzsigntsarequirementID gets a reference to the given int32 and assigns it to the FkiEzsigntsarequirementID field.
func (o *EzsignfolderRequestCompound) SetFkiEzsigntsarequirementID(v int32) {
	o.FkiEzsigntsarequirementID = &v
}

// GetSEzsignfolderDescription returns the SEzsignfolderDescription field value
func (o *EzsignfolderRequestCompound) GetSEzsignfolderDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzsignfolderDescription
}

// GetSEzsignfolderDescriptionOk returns a tuple with the SEzsignfolderDescription field value
// and a boolean to check if the value has been set.
func (o *EzsignfolderRequestCompound) GetSEzsignfolderDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SEzsignfolderDescription, true
}

// SetSEzsignfolderDescription sets field value
func (o *EzsignfolderRequestCompound) SetSEzsignfolderDescription(v string) {
	o.SEzsignfolderDescription = v
}

// GetTEzsignfolderNote returns the TEzsignfolderNote field value if set, zero value otherwise.
func (o *EzsignfolderRequestCompound) GetTEzsignfolderNote() string {
	if o == nil || IsNil(o.TEzsignfolderNote) {
		var ret string
		return ret
	}
	return *o.TEzsignfolderNote
}

// GetTEzsignfolderNoteOk returns a tuple with the TEzsignfolderNote field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignfolderRequestCompound) GetTEzsignfolderNoteOk() (*string, bool) {
	if o == nil || IsNil(o.TEzsignfolderNote) {
		return nil, false
	}
	return o.TEzsignfolderNote, true
}

// HasTEzsignfolderNote returns a boolean if a field has been set.
func (o *EzsignfolderRequestCompound) HasTEzsignfolderNote() bool {
	if o != nil && !IsNil(o.TEzsignfolderNote) {
		return true
	}

	return false
}

// SetTEzsignfolderNote gets a reference to the given string and assigns it to the TEzsignfolderNote field.
func (o *EzsignfolderRequestCompound) SetTEzsignfolderNote(v string) {
	o.TEzsignfolderNote = &v
}

// GetEEzsignfolderSendreminderfrequency returns the EEzsignfolderSendreminderfrequency field value
func (o *EzsignfolderRequestCompound) GetEEzsignfolderSendreminderfrequency() FieldEEzsignfolderSendreminderfrequency {
	if o == nil {
		var ret FieldEEzsignfolderSendreminderfrequency
		return ret
	}

	return o.EEzsignfolderSendreminderfrequency
}

// GetEEzsignfolderSendreminderfrequencyOk returns a tuple with the EEzsignfolderSendreminderfrequency field value
// and a boolean to check if the value has been set.
func (o *EzsignfolderRequestCompound) GetEEzsignfolderSendreminderfrequencyOk() (*FieldEEzsignfolderSendreminderfrequency, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EEzsignfolderSendreminderfrequency, true
}

// SetEEzsignfolderSendreminderfrequency sets field value
func (o *EzsignfolderRequestCompound) SetEEzsignfolderSendreminderfrequency(v FieldEEzsignfolderSendreminderfrequency) {
	o.EEzsignfolderSendreminderfrequency = v
}

// GetSEzsignfolderExternalid returns the SEzsignfolderExternalid field value if set, zero value otherwise.
func (o *EzsignfolderRequestCompound) GetSEzsignfolderExternalid() string {
	if o == nil || IsNil(o.SEzsignfolderExternalid) {
		var ret string
		return ret
	}
	return *o.SEzsignfolderExternalid
}

// GetSEzsignfolderExternalidOk returns a tuple with the SEzsignfolderExternalid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignfolderRequestCompound) GetSEzsignfolderExternalidOk() (*string, bool) {
	if o == nil || IsNil(o.SEzsignfolderExternalid) {
		return nil, false
	}
	return o.SEzsignfolderExternalid, true
}

// HasSEzsignfolderExternalid returns a boolean if a field has been set.
func (o *EzsignfolderRequestCompound) HasSEzsignfolderExternalid() bool {
	if o != nil && !IsNil(o.SEzsignfolderExternalid) {
		return true
	}

	return false
}

// SetSEzsignfolderExternalid gets a reference to the given string and assigns it to the SEzsignfolderExternalid field.
func (o *EzsignfolderRequestCompound) SetSEzsignfolderExternalid(v string) {
	o.SEzsignfolderExternalid = &v
}

func (o EzsignfolderRequestCompound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignfolderRequestCompound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PkiEzsignfolderID) {
		toSerialize["pkiEzsignfolderID"] = o.PkiEzsignfolderID
	}
	toSerialize["fkiEzsignfoldertypeID"] = o.FkiEzsignfoldertypeID
	if !IsNil(o.FkiEzsigntsarequirementID) {
		toSerialize["fkiEzsigntsarequirementID"] = o.FkiEzsigntsarequirementID
	}
	toSerialize["sEzsignfolderDescription"] = o.SEzsignfolderDescription
	if !IsNil(o.TEzsignfolderNote) {
		toSerialize["tEzsignfolderNote"] = o.TEzsignfolderNote
	}
	toSerialize["eEzsignfolderSendreminderfrequency"] = o.EEzsignfolderSendreminderfrequency
	if !IsNil(o.SEzsignfolderExternalid) {
		toSerialize["sEzsignfolderExternalid"] = o.SEzsignfolderExternalid
	}
	return toSerialize, nil
}

type NullableEzsignfolderRequestCompound struct {
	value *EzsignfolderRequestCompound
	isSet bool
}

func (v NullableEzsignfolderRequestCompound) Get() *EzsignfolderRequestCompound {
	return v.value
}

func (v *NullableEzsignfolderRequestCompound) Set(val *EzsignfolderRequestCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignfolderRequestCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignfolderRequestCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignfolderRequestCompound(val *EzsignfolderRequestCompound) *NullableEzsignfolderRequestCompound {
	return &NullableEzsignfolderRequestCompound{value: val, isSet: true}
}

func (v NullableEzsignfolderRequestCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignfolderRequestCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


