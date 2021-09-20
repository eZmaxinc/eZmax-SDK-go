/*
eZmax API Definition

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.0.47
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
)

// EzsignfolderRequestCompound An Ezsignfolder Object and children to create a complete structure
type EzsignfolderRequestCompound struct {
	// An array of signers that will be invited to sign the Ezsigndocuments
	AEzsignfoldersignerassociation []EzsignfoldersignerassociationRequest `json:"a_Ezsignfoldersignerassociation"`
	// The unique ID of the Ezsignfoldertype.    This value can be queried by the API and is also visible in the admin interface.    There are two types of Ezsignfoldertype. **User** and **Shared**. **User** can only be seen by the user who created the folder or its assistants. Access to **Shared** folders are configurable for access and email delivery. You should typically choose a **Shared** type here.
	FkiEzsignfoldertypeID int32 `json:"fkiEzsignfoldertypeID"`
	// The unique ID of the Ezsigntsarequirement.  Determine if a Time Stamping Authority should add a timestamp on each of the signature. Valid values:  |Value|Description| |-|-| |1|No. TSA Timestamping will requested. This will make all signatures a lot faster since no round-trip to the TSA server will be required. Timestamping will be made using eZsign server's time.| |2|Best effort. Timestamping from a Time Stamping Authority will be requested but is not mandatory. In the very improbable case it cannot be completed, the timestamping will be made using eZsign server's time. **Additional fee applies**| |3|Mandatory. Timestamping from a Time Stamping Authority will be requested and is mandatory. In the very improbable case it cannot be completed, the signature will fail and the user will be asked to retry. **Additional fee applies**|
	FkiEzsigntsarequirementID int32 `json:"fkiEzsigntsarequirementID"`
	// The description of the Ezsign Folder
	SEzsignfolderDescription string `json:"sEzsignfolderDescription"`
	// Somes extra notes about the eZsign Folder
	TEzsignfolderNote string `json:"tEzsignfolderNote"`
	EEzsignfolderSendreminderfrequency FieldEEzsignfolderSendreminderfrequency `json:"eEzsignfolderSendreminderfrequency"`
}

// NewEzsignfolderRequestCompound instantiates a new EzsignfolderRequestCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignfolderRequestCompound(aEzsignfoldersignerassociation []EzsignfoldersignerassociationRequest, fkiEzsignfoldertypeID int32, fkiEzsigntsarequirementID int32, sEzsignfolderDescription string, tEzsignfolderNote string, eEzsignfolderSendreminderfrequency FieldEEzsignfolderSendreminderfrequency) *EzsignfolderRequestCompound {
	this := EzsignfolderRequestCompound{}
	this.AEzsignfoldersignerassociation = aEzsignfoldersignerassociation
	this.FkiEzsignfoldertypeID = fkiEzsignfoldertypeID
	this.FkiEzsigntsarequirementID = fkiEzsigntsarequirementID
	this.SEzsignfolderDescription = sEzsignfolderDescription
	this.TEzsignfolderNote = tEzsignfolderNote
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

// GetAEzsignfoldersignerassociation returns the AEzsignfoldersignerassociation field value
func (o *EzsignfolderRequestCompound) GetAEzsignfoldersignerassociation() []EzsignfoldersignerassociationRequest {
	if o == nil {
		var ret []EzsignfoldersignerassociationRequest
		return ret
	}

	return o.AEzsignfoldersignerassociation
}

// GetAEzsignfoldersignerassociationOk returns a tuple with the AEzsignfoldersignerassociation field value
// and a boolean to check if the value has been set.
func (o *EzsignfolderRequestCompound) GetAEzsignfoldersignerassociationOk() (*[]EzsignfoldersignerassociationRequest, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AEzsignfoldersignerassociation, true
}

// SetAEzsignfoldersignerassociation sets field value
func (o *EzsignfolderRequestCompound) SetAEzsignfoldersignerassociation(v []EzsignfoldersignerassociationRequest) {
	o.AEzsignfoldersignerassociation = v
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
	if o == nil  {
		return nil, false
	}
	return &o.FkiEzsignfoldertypeID, true
}

// SetFkiEzsignfoldertypeID sets field value
func (o *EzsignfolderRequestCompound) SetFkiEzsignfoldertypeID(v int32) {
	o.FkiEzsignfoldertypeID = v
}

// GetFkiEzsigntsarequirementID returns the FkiEzsigntsarequirementID field value
func (o *EzsignfolderRequestCompound) GetFkiEzsigntsarequirementID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzsigntsarequirementID
}

// GetFkiEzsigntsarequirementIDOk returns a tuple with the FkiEzsigntsarequirementID field value
// and a boolean to check if the value has been set.
func (o *EzsignfolderRequestCompound) GetFkiEzsigntsarequirementIDOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FkiEzsigntsarequirementID, true
}

// SetFkiEzsigntsarequirementID sets field value
func (o *EzsignfolderRequestCompound) SetFkiEzsigntsarequirementID(v int32) {
	o.FkiEzsigntsarequirementID = v
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
	if o == nil  {
		return nil, false
	}
	return &o.SEzsignfolderDescription, true
}

// SetSEzsignfolderDescription sets field value
func (o *EzsignfolderRequestCompound) SetSEzsignfolderDescription(v string) {
	o.SEzsignfolderDescription = v
}

// GetTEzsignfolderNote returns the TEzsignfolderNote field value
func (o *EzsignfolderRequestCompound) GetTEzsignfolderNote() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TEzsignfolderNote
}

// GetTEzsignfolderNoteOk returns a tuple with the TEzsignfolderNote field value
// and a boolean to check if the value has been set.
func (o *EzsignfolderRequestCompound) GetTEzsignfolderNoteOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TEzsignfolderNote, true
}

// SetTEzsignfolderNote sets field value
func (o *EzsignfolderRequestCompound) SetTEzsignfolderNote(v string) {
	o.TEzsignfolderNote = v
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
	if o == nil  {
		return nil, false
	}
	return &o.EEzsignfolderSendreminderfrequency, true
}

// SetEEzsignfolderSendreminderfrequency sets field value
func (o *EzsignfolderRequestCompound) SetEEzsignfolderSendreminderfrequency(v FieldEEzsignfolderSendreminderfrequency) {
	o.EEzsignfolderSendreminderfrequency = v
}

func (o EzsignfolderRequestCompound) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["a_Ezsignfoldersignerassociation"] = o.AEzsignfoldersignerassociation
	}
	if true {
		toSerialize["fkiEzsignfoldertypeID"] = o.FkiEzsignfoldertypeID
	}
	if true {
		toSerialize["fkiEzsigntsarequirementID"] = o.FkiEzsigntsarequirementID
	}
	if true {
		toSerialize["sEzsignfolderDescription"] = o.SEzsignfolderDescription
	}
	if true {
		toSerialize["tEzsignfolderNote"] = o.TEzsignfolderNote
	}
	if true {
		toSerialize["eEzsignfolderSendreminderfrequency"] = o.EEzsignfolderSendreminderfrequency
	}
	return json.Marshal(toSerialize)
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


