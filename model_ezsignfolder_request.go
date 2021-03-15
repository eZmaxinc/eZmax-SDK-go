/*
 * eZmax API Definition
 *
 * This API expose all the functionnalities for the eZmax and eZsign applications.
 *
 * API version: 1.0.35
 * Contact: support-api@ezmax.ca
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxinc/eZmax-SDK-go

import (
	"encoding/json"
)

// EzsignfolderRequest An Ezsignfolder Object
type EzsignfolderRequest struct {
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

// NewEzsignfolderRequest instantiates a new EzsignfolderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignfolderRequest(fkiEzsignfoldertypeID int32, fkiEzsigntsarequirementID int32, sEzsignfolderDescription string, tEzsignfolderNote string, eEzsignfolderSendreminderfrequency FieldEEzsignfolderSendreminderfrequency) *EzsignfolderRequest {
	this := EzsignfolderRequest{}
	this.FkiEzsignfoldertypeID = fkiEzsignfoldertypeID
	this.FkiEzsigntsarequirementID = fkiEzsigntsarequirementID
	this.SEzsignfolderDescription = sEzsignfolderDescription
	this.TEzsignfolderNote = tEzsignfolderNote
	this.EEzsignfolderSendreminderfrequency = eEzsignfolderSendreminderfrequency
	return &this
}

// NewEzsignfolderRequestWithDefaults instantiates a new EzsignfolderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignfolderRequestWithDefaults() *EzsignfolderRequest {
	this := EzsignfolderRequest{}
	return &this
}

// GetFkiEzsignfoldertypeID returns the FkiEzsignfoldertypeID field value
func (o *EzsignfolderRequest) GetFkiEzsignfoldertypeID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzsignfoldertypeID
}

// GetFkiEzsignfoldertypeIDOk returns a tuple with the FkiEzsignfoldertypeID field value
// and a boolean to check if the value has been set.
func (o *EzsignfolderRequest) GetFkiEzsignfoldertypeIDOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FkiEzsignfoldertypeID, true
}

// SetFkiEzsignfoldertypeID sets field value
func (o *EzsignfolderRequest) SetFkiEzsignfoldertypeID(v int32) {
	o.FkiEzsignfoldertypeID = v
}

// GetFkiEzsigntsarequirementID returns the FkiEzsigntsarequirementID field value
func (o *EzsignfolderRequest) GetFkiEzsigntsarequirementID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzsigntsarequirementID
}

// GetFkiEzsigntsarequirementIDOk returns a tuple with the FkiEzsigntsarequirementID field value
// and a boolean to check if the value has been set.
func (o *EzsignfolderRequest) GetFkiEzsigntsarequirementIDOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FkiEzsigntsarequirementID, true
}

// SetFkiEzsigntsarequirementID sets field value
func (o *EzsignfolderRequest) SetFkiEzsigntsarequirementID(v int32) {
	o.FkiEzsigntsarequirementID = v
}

// GetSEzsignfolderDescription returns the SEzsignfolderDescription field value
func (o *EzsignfolderRequest) GetSEzsignfolderDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzsignfolderDescription
}

// GetSEzsignfolderDescriptionOk returns a tuple with the SEzsignfolderDescription field value
// and a boolean to check if the value has been set.
func (o *EzsignfolderRequest) GetSEzsignfolderDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SEzsignfolderDescription, true
}

// SetSEzsignfolderDescription sets field value
func (o *EzsignfolderRequest) SetSEzsignfolderDescription(v string) {
	o.SEzsignfolderDescription = v
}

// GetTEzsignfolderNote returns the TEzsignfolderNote field value
func (o *EzsignfolderRequest) GetTEzsignfolderNote() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TEzsignfolderNote
}

// GetTEzsignfolderNoteOk returns a tuple with the TEzsignfolderNote field value
// and a boolean to check if the value has been set.
func (o *EzsignfolderRequest) GetTEzsignfolderNoteOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TEzsignfolderNote, true
}

// SetTEzsignfolderNote sets field value
func (o *EzsignfolderRequest) SetTEzsignfolderNote(v string) {
	o.TEzsignfolderNote = v
}

// GetEEzsignfolderSendreminderfrequency returns the EEzsignfolderSendreminderfrequency field value
func (o *EzsignfolderRequest) GetEEzsignfolderSendreminderfrequency() FieldEEzsignfolderSendreminderfrequency {
	if o == nil {
		var ret FieldEEzsignfolderSendreminderfrequency
		return ret
	}

	return o.EEzsignfolderSendreminderfrequency
}

// GetEEzsignfolderSendreminderfrequencyOk returns a tuple with the EEzsignfolderSendreminderfrequency field value
// and a boolean to check if the value has been set.
func (o *EzsignfolderRequest) GetEEzsignfolderSendreminderfrequencyOk() (*FieldEEzsignfolderSendreminderfrequency, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EEzsignfolderSendreminderfrequency, true
}

// SetEEzsignfolderSendreminderfrequency sets field value
func (o *EzsignfolderRequest) SetEEzsignfolderSendreminderfrequency(v FieldEEzsignfolderSendreminderfrequency) {
	o.EEzsignfolderSendreminderfrequency = v
}

func (o EzsignfolderRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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

type NullableEzsignfolderRequest struct {
	value *EzsignfolderRequest
	isSet bool
}

func (v NullableEzsignfolderRequest) Get() *EzsignfolderRequest {
	return v.value
}

func (v *NullableEzsignfolderRequest) Set(val *EzsignfolderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignfolderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignfolderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignfolderRequest(val *EzsignfolderRequest) *NullableEzsignfolderRequest {
	return &NullableEzsignfolderRequest{value: val, isSet: true}
}

func (v NullableEzsignfolderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignfolderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


