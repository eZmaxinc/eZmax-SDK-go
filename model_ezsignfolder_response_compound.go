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
	"bytes"
	"fmt"
)

// checks if the EzsignfolderResponseCompound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignfolderResponseCompound{}

// EzsignfolderResponseCompound An Ezsignfolder Object and children to create a complete structure
type EzsignfolderResponseCompound struct {
	// The unique ID of the Ezsignfolder
	PkiEzsignfolderID int32 `json:"pkiEzsignfolderID"`
	// The unique ID of the Ezsignfoldertype.
	FkiEzsignfoldertypeID *int32 `json:"fkiEzsignfoldertypeID,omitempty"`
	ObjEzsignfoldertype *CustomEzsignfoldertypeResponse `json:"objEzsignfoldertype,omitempty"`
	EEzsignfolderCompletion FieldEEzsignfolderCompletion `json:"eEzsignfolderCompletion"`
	// Deprecated
	SEzsignfoldertypeNameX *string `json:"sEzsignfoldertypeNameX,omitempty"`
	// The unique ID of the Billingentityinternal.
	FkiBillingentityinternalID *int32 `json:"fkiBillingentityinternalID,omitempty"`
	// The description of the Billingentityinternal in the language of the requester
	SBillingentityinternalDescriptionX *string `json:"sBillingentityinternalDescriptionX,omitempty"`
	// The unique ID of the Ezsigntsarequirement.  Determine if a Time Stamping Authority should add a timestamp on each of the signature. Valid values:  |Value|Description| |-|-| |1|No. TSA Timestamping will requested. This will make all signatures a lot faster since no round-trip to the TSA server will be required. Timestamping will be made using eZsign server's time.| |2|Best effort. Timestamping from a Time Stamping Authority will be requested but is not mandatory. In the very improbable case it cannot be completed, the timestamping will be made using eZsign server's time. **Additional fee applies**| |3|Mandatory. Timestamping from a Time Stamping Authority will be requested and is mandatory. In the very improbable case it cannot be completed, the signature will fail and the user will be asked to retry. **Additional fee applies**|
	FkiEzsigntsarequirementID *int32 `json:"fkiEzsigntsarequirementID,omitempty"`
	// The description of the Ezsigntsarequirement in the language of the requester
	SEzsigntsarequirementDescriptionX *string `json:"sEzsigntsarequirementDescriptionX,omitempty"`
	// The description of the Ezsignfolder
	SEzsignfolderDescription string `json:"sEzsignfolderDescription"`
	// Note about the Ezsignfolder
	TEzsignfolderNote *string `json:"tEzsignfolderNote,omitempty"`
	// If the Ezsigndocument can be disposed
	BEzsignfolderIsdisposable *bool `json:"bEzsignfolderIsdisposable,omitempty"`
	EEzsignfolderSendreminderfrequency *FieldEEzsignfolderSendreminderfrequency `json:"eEzsignfolderSendreminderfrequency,omitempty"`
	// The date and time at which the Ezsignfolder will be sent in the future.
	DtEzsignfolderDelayedsenddate *string `json:"dtEzsignfolderDelayedsenddate,omitempty"`
	// The maximum date and time at which the Ezsignfolder can be signed.
	DtEzsignfolderDuedate *string `json:"dtEzsignfolderDuedate,omitempty"`
	// The date and time at which the Ezsignfolder was sent the last time.
	DtEzsignfolderSentdate *string `json:"dtEzsignfolderSentdate,omitempty"`
	// The scheduled date and time at which the Ezsignfolder should be archived.
	DtEzsignfolderScheduledarchive *string `json:"dtEzsignfolderScheduledarchive,omitempty"`
	// The scheduled date at which the Ezsignfolder should be Disposed.
	DtEzsignfolderScheduleddispose *string `json:"dtEzsignfolderScheduleddispose,omitempty"`
	EEzsignfolderStep *FieldEEzsignfolderStep `json:"eEzsignfolderStep,omitempty"`
	// The date and time at which the Ezsignfolder was closed. Either by applying the last signature or by completing it prematurely.
	DtEzsignfolderClose *string `json:"dtEzsignfolderClose,omitempty"`
	// A custom text message that will be added to the email sent.
	TEzsignfolderMessage *string `json:"tEzsignfolderMessage,omitempty"`
	ObjAudit *CommonAudit `json:"objAudit,omitempty"`
	// This field can be used to store an External ID from the client's system.  Anything can be stored in this field, it will never be evaluated by the eZmax system and will be returned AS-IS.  To store multiple values, consider using a JSON formatted structure, a URL encoded string, a CSV or any other custom format. 
	SEzsignfolderExternalid *string `json:"sEzsignfolderExternalid,omitempty"`
}

type _EzsignfolderResponseCompound EzsignfolderResponseCompound

// NewEzsignfolderResponseCompound instantiates a new EzsignfolderResponseCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignfolderResponseCompound(pkiEzsignfolderID int32, eEzsignfolderCompletion FieldEEzsignfolderCompletion, sEzsignfolderDescription string) *EzsignfolderResponseCompound {
	this := EzsignfolderResponseCompound{}
	this.PkiEzsignfolderID = pkiEzsignfolderID
	this.EEzsignfolderCompletion = eEzsignfolderCompletion
	this.SEzsignfolderDescription = sEzsignfolderDescription
	return &this
}

// NewEzsignfolderResponseCompoundWithDefaults instantiates a new EzsignfolderResponseCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignfolderResponseCompoundWithDefaults() *EzsignfolderResponseCompound {
	this := EzsignfolderResponseCompound{}
	return &this
}

// GetPkiEzsignfolderID returns the PkiEzsignfolderID field value
func (o *EzsignfolderResponseCompound) GetPkiEzsignfolderID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiEzsignfolderID
}

// GetPkiEzsignfolderIDOk returns a tuple with the PkiEzsignfolderID field value
// and a boolean to check if the value has been set.
func (o *EzsignfolderResponseCompound) GetPkiEzsignfolderIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiEzsignfolderID, true
}

// SetPkiEzsignfolderID sets field value
func (o *EzsignfolderResponseCompound) SetPkiEzsignfolderID(v int32) {
	o.PkiEzsignfolderID = v
}

// GetFkiEzsignfoldertypeID returns the FkiEzsignfoldertypeID field value if set, zero value otherwise.
func (o *EzsignfolderResponseCompound) GetFkiEzsignfoldertypeID() int32 {
	if o == nil || IsNil(o.FkiEzsignfoldertypeID) {
		var ret int32
		return ret
	}
	return *o.FkiEzsignfoldertypeID
}

// GetFkiEzsignfoldertypeIDOk returns a tuple with the FkiEzsignfoldertypeID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignfolderResponseCompound) GetFkiEzsignfoldertypeIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiEzsignfoldertypeID) {
		return nil, false
	}
	return o.FkiEzsignfoldertypeID, true
}

// HasFkiEzsignfoldertypeID returns a boolean if a field has been set.
func (o *EzsignfolderResponseCompound) HasFkiEzsignfoldertypeID() bool {
	if o != nil && !IsNil(o.FkiEzsignfoldertypeID) {
		return true
	}

	return false
}

// SetFkiEzsignfoldertypeID gets a reference to the given int32 and assigns it to the FkiEzsignfoldertypeID field.
func (o *EzsignfolderResponseCompound) SetFkiEzsignfoldertypeID(v int32) {
	o.FkiEzsignfoldertypeID = &v
}

// GetObjEzsignfoldertype returns the ObjEzsignfoldertype field value if set, zero value otherwise.
func (o *EzsignfolderResponseCompound) GetObjEzsignfoldertype() CustomEzsignfoldertypeResponse {
	if o == nil || IsNil(o.ObjEzsignfoldertype) {
		var ret CustomEzsignfoldertypeResponse
		return ret
	}
	return *o.ObjEzsignfoldertype
}

// GetObjEzsignfoldertypeOk returns a tuple with the ObjEzsignfoldertype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignfolderResponseCompound) GetObjEzsignfoldertypeOk() (*CustomEzsignfoldertypeResponse, bool) {
	if o == nil || IsNil(o.ObjEzsignfoldertype) {
		return nil, false
	}
	return o.ObjEzsignfoldertype, true
}

// HasObjEzsignfoldertype returns a boolean if a field has been set.
func (o *EzsignfolderResponseCompound) HasObjEzsignfoldertype() bool {
	if o != nil && !IsNil(o.ObjEzsignfoldertype) {
		return true
	}

	return false
}

// SetObjEzsignfoldertype gets a reference to the given CustomEzsignfoldertypeResponse and assigns it to the ObjEzsignfoldertype field.
func (o *EzsignfolderResponseCompound) SetObjEzsignfoldertype(v CustomEzsignfoldertypeResponse) {
	o.ObjEzsignfoldertype = &v
}

// GetEEzsignfolderCompletion returns the EEzsignfolderCompletion field value
func (o *EzsignfolderResponseCompound) GetEEzsignfolderCompletion() FieldEEzsignfolderCompletion {
	if o == nil {
		var ret FieldEEzsignfolderCompletion
		return ret
	}

	return o.EEzsignfolderCompletion
}

// GetEEzsignfolderCompletionOk returns a tuple with the EEzsignfolderCompletion field value
// and a boolean to check if the value has been set.
func (o *EzsignfolderResponseCompound) GetEEzsignfolderCompletionOk() (*FieldEEzsignfolderCompletion, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EEzsignfolderCompletion, true
}

// SetEEzsignfolderCompletion sets field value
func (o *EzsignfolderResponseCompound) SetEEzsignfolderCompletion(v FieldEEzsignfolderCompletion) {
	o.EEzsignfolderCompletion = v
}

// GetSEzsignfoldertypeNameX returns the SEzsignfoldertypeNameX field value if set, zero value otherwise.
// Deprecated
func (o *EzsignfolderResponseCompound) GetSEzsignfoldertypeNameX() string {
	if o == nil || IsNil(o.SEzsignfoldertypeNameX) {
		var ret string
		return ret
	}
	return *o.SEzsignfoldertypeNameX
}

// GetSEzsignfoldertypeNameXOk returns a tuple with the SEzsignfoldertypeNameX field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *EzsignfolderResponseCompound) GetSEzsignfoldertypeNameXOk() (*string, bool) {
	if o == nil || IsNil(o.SEzsignfoldertypeNameX) {
		return nil, false
	}
	return o.SEzsignfoldertypeNameX, true
}

// HasSEzsignfoldertypeNameX returns a boolean if a field has been set.
func (o *EzsignfolderResponseCompound) HasSEzsignfoldertypeNameX() bool {
	if o != nil && !IsNil(o.SEzsignfoldertypeNameX) {
		return true
	}

	return false
}

// SetSEzsignfoldertypeNameX gets a reference to the given string and assigns it to the SEzsignfoldertypeNameX field.
// Deprecated
func (o *EzsignfolderResponseCompound) SetSEzsignfoldertypeNameX(v string) {
	o.SEzsignfoldertypeNameX = &v
}

// GetFkiBillingentityinternalID returns the FkiBillingentityinternalID field value if set, zero value otherwise.
func (o *EzsignfolderResponseCompound) GetFkiBillingentityinternalID() int32 {
	if o == nil || IsNil(o.FkiBillingentityinternalID) {
		var ret int32
		return ret
	}
	return *o.FkiBillingentityinternalID
}

// GetFkiBillingentityinternalIDOk returns a tuple with the FkiBillingentityinternalID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignfolderResponseCompound) GetFkiBillingentityinternalIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiBillingentityinternalID) {
		return nil, false
	}
	return o.FkiBillingentityinternalID, true
}

// HasFkiBillingentityinternalID returns a boolean if a field has been set.
func (o *EzsignfolderResponseCompound) HasFkiBillingentityinternalID() bool {
	if o != nil && !IsNil(o.FkiBillingentityinternalID) {
		return true
	}

	return false
}

// SetFkiBillingentityinternalID gets a reference to the given int32 and assigns it to the FkiBillingentityinternalID field.
func (o *EzsignfolderResponseCompound) SetFkiBillingentityinternalID(v int32) {
	o.FkiBillingentityinternalID = &v
}

// GetSBillingentityinternalDescriptionX returns the SBillingentityinternalDescriptionX field value if set, zero value otherwise.
func (o *EzsignfolderResponseCompound) GetSBillingentityinternalDescriptionX() string {
	if o == nil || IsNil(o.SBillingentityinternalDescriptionX) {
		var ret string
		return ret
	}
	return *o.SBillingentityinternalDescriptionX
}

// GetSBillingentityinternalDescriptionXOk returns a tuple with the SBillingentityinternalDescriptionX field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignfolderResponseCompound) GetSBillingentityinternalDescriptionXOk() (*string, bool) {
	if o == nil || IsNil(o.SBillingentityinternalDescriptionX) {
		return nil, false
	}
	return o.SBillingentityinternalDescriptionX, true
}

// HasSBillingentityinternalDescriptionX returns a boolean if a field has been set.
func (o *EzsignfolderResponseCompound) HasSBillingentityinternalDescriptionX() bool {
	if o != nil && !IsNil(o.SBillingentityinternalDescriptionX) {
		return true
	}

	return false
}

// SetSBillingentityinternalDescriptionX gets a reference to the given string and assigns it to the SBillingentityinternalDescriptionX field.
func (o *EzsignfolderResponseCompound) SetSBillingentityinternalDescriptionX(v string) {
	o.SBillingentityinternalDescriptionX = &v
}

// GetFkiEzsigntsarequirementID returns the FkiEzsigntsarequirementID field value if set, zero value otherwise.
func (o *EzsignfolderResponseCompound) GetFkiEzsigntsarequirementID() int32 {
	if o == nil || IsNil(o.FkiEzsigntsarequirementID) {
		var ret int32
		return ret
	}
	return *o.FkiEzsigntsarequirementID
}

// GetFkiEzsigntsarequirementIDOk returns a tuple with the FkiEzsigntsarequirementID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignfolderResponseCompound) GetFkiEzsigntsarequirementIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiEzsigntsarequirementID) {
		return nil, false
	}
	return o.FkiEzsigntsarequirementID, true
}

// HasFkiEzsigntsarequirementID returns a boolean if a field has been set.
func (o *EzsignfolderResponseCompound) HasFkiEzsigntsarequirementID() bool {
	if o != nil && !IsNil(o.FkiEzsigntsarequirementID) {
		return true
	}

	return false
}

// SetFkiEzsigntsarequirementID gets a reference to the given int32 and assigns it to the FkiEzsigntsarequirementID field.
func (o *EzsignfolderResponseCompound) SetFkiEzsigntsarequirementID(v int32) {
	o.FkiEzsigntsarequirementID = &v
}

// GetSEzsigntsarequirementDescriptionX returns the SEzsigntsarequirementDescriptionX field value if set, zero value otherwise.
func (o *EzsignfolderResponseCompound) GetSEzsigntsarequirementDescriptionX() string {
	if o == nil || IsNil(o.SEzsigntsarequirementDescriptionX) {
		var ret string
		return ret
	}
	return *o.SEzsigntsarequirementDescriptionX
}

// GetSEzsigntsarequirementDescriptionXOk returns a tuple with the SEzsigntsarequirementDescriptionX field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignfolderResponseCompound) GetSEzsigntsarequirementDescriptionXOk() (*string, bool) {
	if o == nil || IsNil(o.SEzsigntsarequirementDescriptionX) {
		return nil, false
	}
	return o.SEzsigntsarequirementDescriptionX, true
}

// HasSEzsigntsarequirementDescriptionX returns a boolean if a field has been set.
func (o *EzsignfolderResponseCompound) HasSEzsigntsarequirementDescriptionX() bool {
	if o != nil && !IsNil(o.SEzsigntsarequirementDescriptionX) {
		return true
	}

	return false
}

// SetSEzsigntsarequirementDescriptionX gets a reference to the given string and assigns it to the SEzsigntsarequirementDescriptionX field.
func (o *EzsignfolderResponseCompound) SetSEzsigntsarequirementDescriptionX(v string) {
	o.SEzsigntsarequirementDescriptionX = &v
}

// GetSEzsignfolderDescription returns the SEzsignfolderDescription field value
func (o *EzsignfolderResponseCompound) GetSEzsignfolderDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzsignfolderDescription
}

// GetSEzsignfolderDescriptionOk returns a tuple with the SEzsignfolderDescription field value
// and a boolean to check if the value has been set.
func (o *EzsignfolderResponseCompound) GetSEzsignfolderDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SEzsignfolderDescription, true
}

// SetSEzsignfolderDescription sets field value
func (o *EzsignfolderResponseCompound) SetSEzsignfolderDescription(v string) {
	o.SEzsignfolderDescription = v
}

// GetTEzsignfolderNote returns the TEzsignfolderNote field value if set, zero value otherwise.
func (o *EzsignfolderResponseCompound) GetTEzsignfolderNote() string {
	if o == nil || IsNil(o.TEzsignfolderNote) {
		var ret string
		return ret
	}
	return *o.TEzsignfolderNote
}

// GetTEzsignfolderNoteOk returns a tuple with the TEzsignfolderNote field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignfolderResponseCompound) GetTEzsignfolderNoteOk() (*string, bool) {
	if o == nil || IsNil(o.TEzsignfolderNote) {
		return nil, false
	}
	return o.TEzsignfolderNote, true
}

// HasTEzsignfolderNote returns a boolean if a field has been set.
func (o *EzsignfolderResponseCompound) HasTEzsignfolderNote() bool {
	if o != nil && !IsNil(o.TEzsignfolderNote) {
		return true
	}

	return false
}

// SetTEzsignfolderNote gets a reference to the given string and assigns it to the TEzsignfolderNote field.
func (o *EzsignfolderResponseCompound) SetTEzsignfolderNote(v string) {
	o.TEzsignfolderNote = &v
}

// GetBEzsignfolderIsdisposable returns the BEzsignfolderIsdisposable field value if set, zero value otherwise.
func (o *EzsignfolderResponseCompound) GetBEzsignfolderIsdisposable() bool {
	if o == nil || IsNil(o.BEzsignfolderIsdisposable) {
		var ret bool
		return ret
	}
	return *o.BEzsignfolderIsdisposable
}

// GetBEzsignfolderIsdisposableOk returns a tuple with the BEzsignfolderIsdisposable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignfolderResponseCompound) GetBEzsignfolderIsdisposableOk() (*bool, bool) {
	if o == nil || IsNil(o.BEzsignfolderIsdisposable) {
		return nil, false
	}
	return o.BEzsignfolderIsdisposable, true
}

// HasBEzsignfolderIsdisposable returns a boolean if a field has been set.
func (o *EzsignfolderResponseCompound) HasBEzsignfolderIsdisposable() bool {
	if o != nil && !IsNil(o.BEzsignfolderIsdisposable) {
		return true
	}

	return false
}

// SetBEzsignfolderIsdisposable gets a reference to the given bool and assigns it to the BEzsignfolderIsdisposable field.
func (o *EzsignfolderResponseCompound) SetBEzsignfolderIsdisposable(v bool) {
	o.BEzsignfolderIsdisposable = &v
}

// GetEEzsignfolderSendreminderfrequency returns the EEzsignfolderSendreminderfrequency field value if set, zero value otherwise.
func (o *EzsignfolderResponseCompound) GetEEzsignfolderSendreminderfrequency() FieldEEzsignfolderSendreminderfrequency {
	if o == nil || IsNil(o.EEzsignfolderSendreminderfrequency) {
		var ret FieldEEzsignfolderSendreminderfrequency
		return ret
	}
	return *o.EEzsignfolderSendreminderfrequency
}

// GetEEzsignfolderSendreminderfrequencyOk returns a tuple with the EEzsignfolderSendreminderfrequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignfolderResponseCompound) GetEEzsignfolderSendreminderfrequencyOk() (*FieldEEzsignfolderSendreminderfrequency, bool) {
	if o == nil || IsNil(o.EEzsignfolderSendreminderfrequency) {
		return nil, false
	}
	return o.EEzsignfolderSendreminderfrequency, true
}

// HasEEzsignfolderSendreminderfrequency returns a boolean if a field has been set.
func (o *EzsignfolderResponseCompound) HasEEzsignfolderSendreminderfrequency() bool {
	if o != nil && !IsNil(o.EEzsignfolderSendreminderfrequency) {
		return true
	}

	return false
}

// SetEEzsignfolderSendreminderfrequency gets a reference to the given FieldEEzsignfolderSendreminderfrequency and assigns it to the EEzsignfolderSendreminderfrequency field.
func (o *EzsignfolderResponseCompound) SetEEzsignfolderSendreminderfrequency(v FieldEEzsignfolderSendreminderfrequency) {
	o.EEzsignfolderSendreminderfrequency = &v
}

// GetDtEzsignfolderDelayedsenddate returns the DtEzsignfolderDelayedsenddate field value if set, zero value otherwise.
func (o *EzsignfolderResponseCompound) GetDtEzsignfolderDelayedsenddate() string {
	if o == nil || IsNil(o.DtEzsignfolderDelayedsenddate) {
		var ret string
		return ret
	}
	return *o.DtEzsignfolderDelayedsenddate
}

// GetDtEzsignfolderDelayedsenddateOk returns a tuple with the DtEzsignfolderDelayedsenddate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignfolderResponseCompound) GetDtEzsignfolderDelayedsenddateOk() (*string, bool) {
	if o == nil || IsNil(o.DtEzsignfolderDelayedsenddate) {
		return nil, false
	}
	return o.DtEzsignfolderDelayedsenddate, true
}

// HasDtEzsignfolderDelayedsenddate returns a boolean if a field has been set.
func (o *EzsignfolderResponseCompound) HasDtEzsignfolderDelayedsenddate() bool {
	if o != nil && !IsNil(o.DtEzsignfolderDelayedsenddate) {
		return true
	}

	return false
}

// SetDtEzsignfolderDelayedsenddate gets a reference to the given string and assigns it to the DtEzsignfolderDelayedsenddate field.
func (o *EzsignfolderResponseCompound) SetDtEzsignfolderDelayedsenddate(v string) {
	o.DtEzsignfolderDelayedsenddate = &v
}

// GetDtEzsignfolderDuedate returns the DtEzsignfolderDuedate field value if set, zero value otherwise.
func (o *EzsignfolderResponseCompound) GetDtEzsignfolderDuedate() string {
	if o == nil || IsNil(o.DtEzsignfolderDuedate) {
		var ret string
		return ret
	}
	return *o.DtEzsignfolderDuedate
}

// GetDtEzsignfolderDuedateOk returns a tuple with the DtEzsignfolderDuedate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignfolderResponseCompound) GetDtEzsignfolderDuedateOk() (*string, bool) {
	if o == nil || IsNil(o.DtEzsignfolderDuedate) {
		return nil, false
	}
	return o.DtEzsignfolderDuedate, true
}

// HasDtEzsignfolderDuedate returns a boolean if a field has been set.
func (o *EzsignfolderResponseCompound) HasDtEzsignfolderDuedate() bool {
	if o != nil && !IsNil(o.DtEzsignfolderDuedate) {
		return true
	}

	return false
}

// SetDtEzsignfolderDuedate gets a reference to the given string and assigns it to the DtEzsignfolderDuedate field.
func (o *EzsignfolderResponseCompound) SetDtEzsignfolderDuedate(v string) {
	o.DtEzsignfolderDuedate = &v
}

// GetDtEzsignfolderSentdate returns the DtEzsignfolderSentdate field value if set, zero value otherwise.
func (o *EzsignfolderResponseCompound) GetDtEzsignfolderSentdate() string {
	if o == nil || IsNil(o.DtEzsignfolderSentdate) {
		var ret string
		return ret
	}
	return *o.DtEzsignfolderSentdate
}

// GetDtEzsignfolderSentdateOk returns a tuple with the DtEzsignfolderSentdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignfolderResponseCompound) GetDtEzsignfolderSentdateOk() (*string, bool) {
	if o == nil || IsNil(o.DtEzsignfolderSentdate) {
		return nil, false
	}
	return o.DtEzsignfolderSentdate, true
}

// HasDtEzsignfolderSentdate returns a boolean if a field has been set.
func (o *EzsignfolderResponseCompound) HasDtEzsignfolderSentdate() bool {
	if o != nil && !IsNil(o.DtEzsignfolderSentdate) {
		return true
	}

	return false
}

// SetDtEzsignfolderSentdate gets a reference to the given string and assigns it to the DtEzsignfolderSentdate field.
func (o *EzsignfolderResponseCompound) SetDtEzsignfolderSentdate(v string) {
	o.DtEzsignfolderSentdate = &v
}

// GetDtEzsignfolderScheduledarchive returns the DtEzsignfolderScheduledarchive field value if set, zero value otherwise.
func (o *EzsignfolderResponseCompound) GetDtEzsignfolderScheduledarchive() string {
	if o == nil || IsNil(o.DtEzsignfolderScheduledarchive) {
		var ret string
		return ret
	}
	return *o.DtEzsignfolderScheduledarchive
}

// GetDtEzsignfolderScheduledarchiveOk returns a tuple with the DtEzsignfolderScheduledarchive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignfolderResponseCompound) GetDtEzsignfolderScheduledarchiveOk() (*string, bool) {
	if o == nil || IsNil(o.DtEzsignfolderScheduledarchive) {
		return nil, false
	}
	return o.DtEzsignfolderScheduledarchive, true
}

// HasDtEzsignfolderScheduledarchive returns a boolean if a field has been set.
func (o *EzsignfolderResponseCompound) HasDtEzsignfolderScheduledarchive() bool {
	if o != nil && !IsNil(o.DtEzsignfolderScheduledarchive) {
		return true
	}

	return false
}

// SetDtEzsignfolderScheduledarchive gets a reference to the given string and assigns it to the DtEzsignfolderScheduledarchive field.
func (o *EzsignfolderResponseCompound) SetDtEzsignfolderScheduledarchive(v string) {
	o.DtEzsignfolderScheduledarchive = &v
}

// GetDtEzsignfolderScheduleddispose returns the DtEzsignfolderScheduleddispose field value if set, zero value otherwise.
func (o *EzsignfolderResponseCompound) GetDtEzsignfolderScheduleddispose() string {
	if o == nil || IsNil(o.DtEzsignfolderScheduleddispose) {
		var ret string
		return ret
	}
	return *o.DtEzsignfolderScheduleddispose
}

// GetDtEzsignfolderScheduleddisposeOk returns a tuple with the DtEzsignfolderScheduleddispose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignfolderResponseCompound) GetDtEzsignfolderScheduleddisposeOk() (*string, bool) {
	if o == nil || IsNil(o.DtEzsignfolderScheduleddispose) {
		return nil, false
	}
	return o.DtEzsignfolderScheduleddispose, true
}

// HasDtEzsignfolderScheduleddispose returns a boolean if a field has been set.
func (o *EzsignfolderResponseCompound) HasDtEzsignfolderScheduleddispose() bool {
	if o != nil && !IsNil(o.DtEzsignfolderScheduleddispose) {
		return true
	}

	return false
}

// SetDtEzsignfolderScheduleddispose gets a reference to the given string and assigns it to the DtEzsignfolderScheduleddispose field.
func (o *EzsignfolderResponseCompound) SetDtEzsignfolderScheduleddispose(v string) {
	o.DtEzsignfolderScheduleddispose = &v
}

// GetEEzsignfolderStep returns the EEzsignfolderStep field value if set, zero value otherwise.
func (o *EzsignfolderResponseCompound) GetEEzsignfolderStep() FieldEEzsignfolderStep {
	if o == nil || IsNil(o.EEzsignfolderStep) {
		var ret FieldEEzsignfolderStep
		return ret
	}
	return *o.EEzsignfolderStep
}

// GetEEzsignfolderStepOk returns a tuple with the EEzsignfolderStep field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignfolderResponseCompound) GetEEzsignfolderStepOk() (*FieldEEzsignfolderStep, bool) {
	if o == nil || IsNil(o.EEzsignfolderStep) {
		return nil, false
	}
	return o.EEzsignfolderStep, true
}

// HasEEzsignfolderStep returns a boolean if a field has been set.
func (o *EzsignfolderResponseCompound) HasEEzsignfolderStep() bool {
	if o != nil && !IsNil(o.EEzsignfolderStep) {
		return true
	}

	return false
}

// SetEEzsignfolderStep gets a reference to the given FieldEEzsignfolderStep and assigns it to the EEzsignfolderStep field.
func (o *EzsignfolderResponseCompound) SetEEzsignfolderStep(v FieldEEzsignfolderStep) {
	o.EEzsignfolderStep = &v
}

// GetDtEzsignfolderClose returns the DtEzsignfolderClose field value if set, zero value otherwise.
func (o *EzsignfolderResponseCompound) GetDtEzsignfolderClose() string {
	if o == nil || IsNil(o.DtEzsignfolderClose) {
		var ret string
		return ret
	}
	return *o.DtEzsignfolderClose
}

// GetDtEzsignfolderCloseOk returns a tuple with the DtEzsignfolderClose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignfolderResponseCompound) GetDtEzsignfolderCloseOk() (*string, bool) {
	if o == nil || IsNil(o.DtEzsignfolderClose) {
		return nil, false
	}
	return o.DtEzsignfolderClose, true
}

// HasDtEzsignfolderClose returns a boolean if a field has been set.
func (o *EzsignfolderResponseCompound) HasDtEzsignfolderClose() bool {
	if o != nil && !IsNil(o.DtEzsignfolderClose) {
		return true
	}

	return false
}

// SetDtEzsignfolderClose gets a reference to the given string and assigns it to the DtEzsignfolderClose field.
func (o *EzsignfolderResponseCompound) SetDtEzsignfolderClose(v string) {
	o.DtEzsignfolderClose = &v
}

// GetTEzsignfolderMessage returns the TEzsignfolderMessage field value if set, zero value otherwise.
func (o *EzsignfolderResponseCompound) GetTEzsignfolderMessage() string {
	if o == nil || IsNil(o.TEzsignfolderMessage) {
		var ret string
		return ret
	}
	return *o.TEzsignfolderMessage
}

// GetTEzsignfolderMessageOk returns a tuple with the TEzsignfolderMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignfolderResponseCompound) GetTEzsignfolderMessageOk() (*string, bool) {
	if o == nil || IsNil(o.TEzsignfolderMessage) {
		return nil, false
	}
	return o.TEzsignfolderMessage, true
}

// HasTEzsignfolderMessage returns a boolean if a field has been set.
func (o *EzsignfolderResponseCompound) HasTEzsignfolderMessage() bool {
	if o != nil && !IsNil(o.TEzsignfolderMessage) {
		return true
	}

	return false
}

// SetTEzsignfolderMessage gets a reference to the given string and assigns it to the TEzsignfolderMessage field.
func (o *EzsignfolderResponseCompound) SetTEzsignfolderMessage(v string) {
	o.TEzsignfolderMessage = &v
}

// GetObjAudit returns the ObjAudit field value if set, zero value otherwise.
func (o *EzsignfolderResponseCompound) GetObjAudit() CommonAudit {
	if o == nil || IsNil(o.ObjAudit) {
		var ret CommonAudit
		return ret
	}
	return *o.ObjAudit
}

// GetObjAuditOk returns a tuple with the ObjAudit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignfolderResponseCompound) GetObjAuditOk() (*CommonAudit, bool) {
	if o == nil || IsNil(o.ObjAudit) {
		return nil, false
	}
	return o.ObjAudit, true
}

// HasObjAudit returns a boolean if a field has been set.
func (o *EzsignfolderResponseCompound) HasObjAudit() bool {
	if o != nil && !IsNil(o.ObjAudit) {
		return true
	}

	return false
}

// SetObjAudit gets a reference to the given CommonAudit and assigns it to the ObjAudit field.
func (o *EzsignfolderResponseCompound) SetObjAudit(v CommonAudit) {
	o.ObjAudit = &v
}

// GetSEzsignfolderExternalid returns the SEzsignfolderExternalid field value if set, zero value otherwise.
func (o *EzsignfolderResponseCompound) GetSEzsignfolderExternalid() string {
	if o == nil || IsNil(o.SEzsignfolderExternalid) {
		var ret string
		return ret
	}
	return *o.SEzsignfolderExternalid
}

// GetSEzsignfolderExternalidOk returns a tuple with the SEzsignfolderExternalid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignfolderResponseCompound) GetSEzsignfolderExternalidOk() (*string, bool) {
	if o == nil || IsNil(o.SEzsignfolderExternalid) {
		return nil, false
	}
	return o.SEzsignfolderExternalid, true
}

// HasSEzsignfolderExternalid returns a boolean if a field has been set.
func (o *EzsignfolderResponseCompound) HasSEzsignfolderExternalid() bool {
	if o != nil && !IsNil(o.SEzsignfolderExternalid) {
		return true
	}

	return false
}

// SetSEzsignfolderExternalid gets a reference to the given string and assigns it to the SEzsignfolderExternalid field.
func (o *EzsignfolderResponseCompound) SetSEzsignfolderExternalid(v string) {
	o.SEzsignfolderExternalid = &v
}

func (o EzsignfolderResponseCompound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignfolderResponseCompound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiEzsignfolderID"] = o.PkiEzsignfolderID
	if !IsNil(o.FkiEzsignfoldertypeID) {
		toSerialize["fkiEzsignfoldertypeID"] = o.FkiEzsignfoldertypeID
	}
	if !IsNil(o.ObjEzsignfoldertype) {
		toSerialize["objEzsignfoldertype"] = o.ObjEzsignfoldertype
	}
	toSerialize["eEzsignfolderCompletion"] = o.EEzsignfolderCompletion
	if !IsNil(o.SEzsignfoldertypeNameX) {
		toSerialize["sEzsignfoldertypeNameX"] = o.SEzsignfoldertypeNameX
	}
	if !IsNil(o.FkiBillingentityinternalID) {
		toSerialize["fkiBillingentityinternalID"] = o.FkiBillingentityinternalID
	}
	if !IsNil(o.SBillingentityinternalDescriptionX) {
		toSerialize["sBillingentityinternalDescriptionX"] = o.SBillingentityinternalDescriptionX
	}
	if !IsNil(o.FkiEzsigntsarequirementID) {
		toSerialize["fkiEzsigntsarequirementID"] = o.FkiEzsigntsarequirementID
	}
	if !IsNil(o.SEzsigntsarequirementDescriptionX) {
		toSerialize["sEzsigntsarequirementDescriptionX"] = o.SEzsigntsarequirementDescriptionX
	}
	toSerialize["sEzsignfolderDescription"] = o.SEzsignfolderDescription
	if !IsNil(o.TEzsignfolderNote) {
		toSerialize["tEzsignfolderNote"] = o.TEzsignfolderNote
	}
	if !IsNil(o.BEzsignfolderIsdisposable) {
		toSerialize["bEzsignfolderIsdisposable"] = o.BEzsignfolderIsdisposable
	}
	if !IsNil(o.EEzsignfolderSendreminderfrequency) {
		toSerialize["eEzsignfolderSendreminderfrequency"] = o.EEzsignfolderSendreminderfrequency
	}
	if !IsNil(o.DtEzsignfolderDelayedsenddate) {
		toSerialize["dtEzsignfolderDelayedsenddate"] = o.DtEzsignfolderDelayedsenddate
	}
	if !IsNil(o.DtEzsignfolderDuedate) {
		toSerialize["dtEzsignfolderDuedate"] = o.DtEzsignfolderDuedate
	}
	if !IsNil(o.DtEzsignfolderSentdate) {
		toSerialize["dtEzsignfolderSentdate"] = o.DtEzsignfolderSentdate
	}
	if !IsNil(o.DtEzsignfolderScheduledarchive) {
		toSerialize["dtEzsignfolderScheduledarchive"] = o.DtEzsignfolderScheduledarchive
	}
	if !IsNil(o.DtEzsignfolderScheduleddispose) {
		toSerialize["dtEzsignfolderScheduleddispose"] = o.DtEzsignfolderScheduleddispose
	}
	if !IsNil(o.EEzsignfolderStep) {
		toSerialize["eEzsignfolderStep"] = o.EEzsignfolderStep
	}
	if !IsNil(o.DtEzsignfolderClose) {
		toSerialize["dtEzsignfolderClose"] = o.DtEzsignfolderClose
	}
	if !IsNil(o.TEzsignfolderMessage) {
		toSerialize["tEzsignfolderMessage"] = o.TEzsignfolderMessage
	}
	if !IsNil(o.ObjAudit) {
		toSerialize["objAudit"] = o.ObjAudit
	}
	if !IsNil(o.SEzsignfolderExternalid) {
		toSerialize["sEzsignfolderExternalid"] = o.SEzsignfolderExternalid
	}
	return toSerialize, nil
}

func (o *EzsignfolderResponseCompound) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiEzsignfolderID",
		"eEzsignfolderCompletion",
		"sEzsignfolderDescription",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varEzsignfolderResponseCompound := _EzsignfolderResponseCompound{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsignfolderResponseCompound)

	if err != nil {
		return err
	}

	*o = EzsignfolderResponseCompound(varEzsignfolderResponseCompound)

	return err
}

type NullableEzsignfolderResponseCompound struct {
	value *EzsignfolderResponseCompound
	isSet bool
}

func (v NullableEzsignfolderResponseCompound) Get() *EzsignfolderResponseCompound {
	return v.value
}

func (v *NullableEzsignfolderResponseCompound) Set(val *EzsignfolderResponseCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignfolderResponseCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignfolderResponseCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignfolderResponseCompound(val *EzsignfolderResponseCompound) *NullableEzsignfolderResponseCompound {
	return &NullableEzsignfolderResponseCompound{value: val, isSet: true}
}

func (v NullableEzsignfolderResponseCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignfolderResponseCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


