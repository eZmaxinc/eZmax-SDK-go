/*
 * eZmax API Definition
 *
 * This API expose all the functionnalities for the eZmax and eZsign application.  We provide SDKs for customers. They are generated using OpenAPI codegen, we encourage customers to use them as we also provide samples for them.  You can choose to build your own implementation manually or can use any compatible OpenAPI 3.0 generator like Swagger Codegen, OpenAPI codegen or any commercial generators.  If you need helping understanding how to use this API, don't waste too much time looking for it. Contact support-api@ezmax.ca, we're here to help. We are developpers so we know programmers don't like bad documentation. If you don't find what you need in the documentation, let us know, we'll improve it and put you rapidly up on track.
 *
 * API version: 1.0.28
 * Contact: support-api@ezmax.ca
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxinc/eZmax-SDK-go

import (
	"encoding/json"
)

// EzsignfolderResponse An Ezsignfolder Object
type EzsignfolderResponse struct {
	// The unique ID of the Ezsignfoldertype.    This value can be queried by the API and is also visible in the admin interface.    There are two types of Ezsignfoldertype. **User** and **Shared**. **User** can only be seen by the user who created the folder or its assistants. Access to **Shared** folders are configurable for access and email delivery. You should typically choose a **Shared** type here.
	FkiEzsignfoldertypeID int32 `json:"fkiEzsignfoldertypeID"`
	// The unique ID of the Ezsigntsarequirement.  Determine if a Time Stamping Authority should add a timestamp on each of the signature. Valid values:  |Value|Description| |-|-| |1|No. TSA Timestamping will requested. This will make all signatures a lot faster since no round-trip to the TSA server will be required. Timestamping will be made using eZsign server's time.| |2|Best effort. Timestamping from a Time Stamping Authority will be requested but is not mandatory. In the very improbable case it cannot be completed, the timestamping will be made using eZsign server's time. **Additional fee applies**| |3|Mandatory. Timestamping from a Time Stamping Authority will be requested and is mandatory. In the very improbable case it cannot be completed, the signature will fail and the user will be asked to retry. **Additional fee applies**|
	FkiEzsigntsarequirementID int32 `json:"fkiEzsigntsarequirementID"`
	// The description of the Ezsign Folder
	SEzsignfolderDescription string `json:"sEzsignfolderDescription"`
	// Somes extra notes about the eZsign Folder
	TEzsignfolderNote string `json:"tEzsignfolderNote"`
	EEzsignfolderSendreminderfrequency FieldEEzsignfolderSendreminderfrequency `json:"eEzsignfolderSendreminderfrequency"`
	// The unique ID of the Ezsignfolder
	PkiEzsignfolderID int32 `json:"pkiEzsignfolderID"`
	// The date and time at which the Ezsign folder was sent the last time.
	DtEzsignfolderSentdate string `json:"dtEzsignfolderSentdate"`
	EEzsignfolderStep FieldEEzsignfolderStep `json:"eEzsignfolderStep"`
	// The date and time at which the folder was closed. Either by applying the last signature or by completing it prematurely.
	DtEzsignfolderClose string `json:"dtEzsignfolderClose"`
	ObjAudit CommonAudit `json:"objAudit"`
}

// NewEzsignfolderResponse instantiates a new EzsignfolderResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignfolderResponse(fkiEzsignfoldertypeID int32, fkiEzsigntsarequirementID int32, sEzsignfolderDescription string, tEzsignfolderNote string, eEzsignfolderSendreminderfrequency FieldEEzsignfolderSendreminderfrequency, pkiEzsignfolderID int32, dtEzsignfolderSentdate string, eEzsignfolderStep FieldEEzsignfolderStep, dtEzsignfolderClose string, objAudit CommonAudit, ) *EzsignfolderResponse {
	this := EzsignfolderResponse{}
	this.FkiEzsignfoldertypeID = fkiEzsignfoldertypeID
	this.FkiEzsigntsarequirementID = fkiEzsigntsarequirementID
	this.SEzsignfolderDescription = sEzsignfolderDescription
	this.TEzsignfolderNote = tEzsignfolderNote
	this.EEzsignfolderSendreminderfrequency = eEzsignfolderSendreminderfrequency
	this.PkiEzsignfolderID = pkiEzsignfolderID
	this.DtEzsignfolderSentdate = dtEzsignfolderSentdate
	this.EEzsignfolderStep = eEzsignfolderStep
	this.DtEzsignfolderClose = dtEzsignfolderClose
	this.ObjAudit = objAudit
	return &this
}

// NewEzsignfolderResponseWithDefaults instantiates a new EzsignfolderResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignfolderResponseWithDefaults() *EzsignfolderResponse {
	this := EzsignfolderResponse{}
	return &this
}

// GetFkiEzsignfoldertypeID returns the FkiEzsignfoldertypeID field value
func (o *EzsignfolderResponse) GetFkiEzsignfoldertypeID() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.FkiEzsignfoldertypeID
}

// GetFkiEzsignfoldertypeIDOk returns a tuple with the FkiEzsignfoldertypeID field value
// and a boolean to check if the value has been set.
func (o *EzsignfolderResponse) GetFkiEzsignfoldertypeIDOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FkiEzsignfoldertypeID, true
}

// SetFkiEzsignfoldertypeID sets field value
func (o *EzsignfolderResponse) SetFkiEzsignfoldertypeID(v int32) {
	o.FkiEzsignfoldertypeID = v
}

// GetFkiEzsigntsarequirementID returns the FkiEzsigntsarequirementID field value
func (o *EzsignfolderResponse) GetFkiEzsigntsarequirementID() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.FkiEzsigntsarequirementID
}

// GetFkiEzsigntsarequirementIDOk returns a tuple with the FkiEzsigntsarequirementID field value
// and a boolean to check if the value has been set.
func (o *EzsignfolderResponse) GetFkiEzsigntsarequirementIDOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FkiEzsigntsarequirementID, true
}

// SetFkiEzsigntsarequirementID sets field value
func (o *EzsignfolderResponse) SetFkiEzsigntsarequirementID(v int32) {
	o.FkiEzsigntsarequirementID = v
}

// GetSEzsignfolderDescription returns the SEzsignfolderDescription field value
func (o *EzsignfolderResponse) GetSEzsignfolderDescription() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.SEzsignfolderDescription
}

// GetSEzsignfolderDescriptionOk returns a tuple with the SEzsignfolderDescription field value
// and a boolean to check if the value has been set.
func (o *EzsignfolderResponse) GetSEzsignfolderDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SEzsignfolderDescription, true
}

// SetSEzsignfolderDescription sets field value
func (o *EzsignfolderResponse) SetSEzsignfolderDescription(v string) {
	o.SEzsignfolderDescription = v
}

// GetTEzsignfolderNote returns the TEzsignfolderNote field value
func (o *EzsignfolderResponse) GetTEzsignfolderNote() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.TEzsignfolderNote
}

// GetTEzsignfolderNoteOk returns a tuple with the TEzsignfolderNote field value
// and a boolean to check if the value has been set.
func (o *EzsignfolderResponse) GetTEzsignfolderNoteOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TEzsignfolderNote, true
}

// SetTEzsignfolderNote sets field value
func (o *EzsignfolderResponse) SetTEzsignfolderNote(v string) {
	o.TEzsignfolderNote = v
}

// GetEEzsignfolderSendreminderfrequency returns the EEzsignfolderSendreminderfrequency field value
func (o *EzsignfolderResponse) GetEEzsignfolderSendreminderfrequency() FieldEEzsignfolderSendreminderfrequency {
	if o == nil  {
		var ret FieldEEzsignfolderSendreminderfrequency
		return ret
	}

	return o.EEzsignfolderSendreminderfrequency
}

// GetEEzsignfolderSendreminderfrequencyOk returns a tuple with the EEzsignfolderSendreminderfrequency field value
// and a boolean to check if the value has been set.
func (o *EzsignfolderResponse) GetEEzsignfolderSendreminderfrequencyOk() (*FieldEEzsignfolderSendreminderfrequency, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EEzsignfolderSendreminderfrequency, true
}

// SetEEzsignfolderSendreminderfrequency sets field value
func (o *EzsignfolderResponse) SetEEzsignfolderSendreminderfrequency(v FieldEEzsignfolderSendreminderfrequency) {
	o.EEzsignfolderSendreminderfrequency = v
}

// GetPkiEzsignfolderID returns the PkiEzsignfolderID field value
func (o *EzsignfolderResponse) GetPkiEzsignfolderID() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.PkiEzsignfolderID
}

// GetPkiEzsignfolderIDOk returns a tuple with the PkiEzsignfolderID field value
// and a boolean to check if the value has been set.
func (o *EzsignfolderResponse) GetPkiEzsignfolderIDOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PkiEzsignfolderID, true
}

// SetPkiEzsignfolderID sets field value
func (o *EzsignfolderResponse) SetPkiEzsignfolderID(v int32) {
	o.PkiEzsignfolderID = v
}

// GetDtEzsignfolderSentdate returns the DtEzsignfolderSentdate field value
func (o *EzsignfolderResponse) GetDtEzsignfolderSentdate() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.DtEzsignfolderSentdate
}

// GetDtEzsignfolderSentdateOk returns a tuple with the DtEzsignfolderSentdate field value
// and a boolean to check if the value has been set.
func (o *EzsignfolderResponse) GetDtEzsignfolderSentdateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DtEzsignfolderSentdate, true
}

// SetDtEzsignfolderSentdate sets field value
func (o *EzsignfolderResponse) SetDtEzsignfolderSentdate(v string) {
	o.DtEzsignfolderSentdate = v
}

// GetEEzsignfolderStep returns the EEzsignfolderStep field value
func (o *EzsignfolderResponse) GetEEzsignfolderStep() FieldEEzsignfolderStep {
	if o == nil  {
		var ret FieldEEzsignfolderStep
		return ret
	}

	return o.EEzsignfolderStep
}

// GetEEzsignfolderStepOk returns a tuple with the EEzsignfolderStep field value
// and a boolean to check if the value has been set.
func (o *EzsignfolderResponse) GetEEzsignfolderStepOk() (*FieldEEzsignfolderStep, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EEzsignfolderStep, true
}

// SetEEzsignfolderStep sets field value
func (o *EzsignfolderResponse) SetEEzsignfolderStep(v FieldEEzsignfolderStep) {
	o.EEzsignfolderStep = v
}

// GetDtEzsignfolderClose returns the DtEzsignfolderClose field value
func (o *EzsignfolderResponse) GetDtEzsignfolderClose() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.DtEzsignfolderClose
}

// GetDtEzsignfolderCloseOk returns a tuple with the DtEzsignfolderClose field value
// and a boolean to check if the value has been set.
func (o *EzsignfolderResponse) GetDtEzsignfolderCloseOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DtEzsignfolderClose, true
}

// SetDtEzsignfolderClose sets field value
func (o *EzsignfolderResponse) SetDtEzsignfolderClose(v string) {
	o.DtEzsignfolderClose = v
}

// GetObjAudit returns the ObjAudit field value
func (o *EzsignfolderResponse) GetObjAudit() CommonAudit {
	if o == nil  {
		var ret CommonAudit
		return ret
	}

	return o.ObjAudit
}

// GetObjAuditOk returns a tuple with the ObjAudit field value
// and a boolean to check if the value has been set.
func (o *EzsignfolderResponse) GetObjAuditOk() (*CommonAudit, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjAudit, true
}

// SetObjAudit sets field value
func (o *EzsignfolderResponse) SetObjAudit(v CommonAudit) {
	o.ObjAudit = v
}

func (o EzsignfolderResponse) MarshalJSON() ([]byte, error) {
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
	if true {
		toSerialize["pkiEzsignfolderID"] = o.PkiEzsignfolderID
	}
	if true {
		toSerialize["dtEzsignfolderSentdate"] = o.DtEzsignfolderSentdate
	}
	if true {
		toSerialize["eEzsignfolderStep"] = o.EEzsignfolderStep
	}
	if true {
		toSerialize["dtEzsignfolderClose"] = o.DtEzsignfolderClose
	}
	if true {
		toSerialize["objAudit"] = o.ObjAudit
	}
	return json.Marshal(toSerialize)
}

type NullableEzsignfolderResponse struct {
	value *EzsignfolderResponse
	isSet bool
}

func (v NullableEzsignfolderResponse) Get() *EzsignfolderResponse {
	return v.value
}

func (v *NullableEzsignfolderResponse) Set(val *EzsignfolderResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignfolderResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignfolderResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignfolderResponse(val *EzsignfolderResponse) *NullableEzsignfolderResponse {
	return &NullableEzsignfolderResponse{value: val, isSet: true}
}

func (v NullableEzsignfolderResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignfolderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


