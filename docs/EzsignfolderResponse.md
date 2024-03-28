# EzsignfolderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsignfolderID** | **int32** | The unique ID of the Ezsignfolder | 
**FkiEzsignfoldertypeID** | Pointer to **int32** | The unique ID of the Ezsignfoldertype. | [optional] 
**ObjEzsignfoldertype** | Pointer to [**CustomEzsignfoldertypeResponse**](CustomEzsignfoldertypeResponse.md) |  | [optional] 
**EEzsignfolderCompletion** | [**FieldEEzsignfolderCompletion**](FieldEEzsignfolderCompletion.md) |  | 
**SEzsignfoldertypeNameX** | Pointer to **string** |  | [optional] 
**FkiBillingentityinternalID** | Pointer to **int32** | The unique ID of the Billingentityinternal. | [optional] 
**SBillingentityinternalDescriptionX** | Pointer to **string** | The description of the Billingentityinternal in the language of the requester | [optional] 
**FkiEzsigntsarequirementID** | Pointer to **int32** | The unique ID of the Ezsigntsarequirement.  Determine if a Time Stamping Authority should add a timestamp on each of the signature. Valid values:  |Value|Description| |-|-| |1|No. TSA Timestamping will requested. This will make all signatures a lot faster since no round-trip to the TSA server will be required. Timestamping will be made using eZsign server&#39;s time.| |2|Best effort. Timestamping from a Time Stamping Authority will be requested but is not mandatory. In the very improbable case it cannot be completed, the timestamping will be made using eZsign server&#39;s time. **Additional fee applies**| |3|Mandatory. Timestamping from a Time Stamping Authority will be requested and is mandatory. In the very improbable case it cannot be completed, the signature will fail and the user will be asked to retry. **Additional fee applies**| | [optional] 
**SEzsigntsarequirementDescriptionX** | Pointer to **string** | The description of the Ezsigntsarequirement in the language of the requester | [optional] 
**SEzsignfolderDescription** | **string** | The description of the Ezsignfolder | 
**TEzsignfolderNote** | Pointer to **string** | Note about the Ezsignfolder | [optional] 
**BEzsignfolderIsdisposable** | Pointer to **bool** | If the Ezsigndocument can be disposed | [optional] 
**EEzsignfolderSendreminderfrequency** | Pointer to [**FieldEEzsignfolderSendreminderfrequency**](FieldEEzsignfolderSendreminderfrequency.md) |  | [optional] 
**DtEzsignfolderDelayedsenddate** | Pointer to **string** | The date and time at which the Ezsignfolder will be sent in the future. | [optional] 
**DtEzsignfolderDuedate** | Pointer to **string** | The maximum date and time at which the Ezsignfolder can be signed. | [optional] 
**DtEzsignfolderSentdate** | Pointer to **string** | The date and time at which the Ezsignfolder was sent the last time. | [optional] 
**DtEzsignfolderScheduledarchive** | Pointer to **string** | The scheduled date and time at which the Ezsignfolder should be archived. | [optional] 
**DtEzsignfolderScheduleddispose** | Pointer to **string** | The scheduled date at which the Ezsignfolder should be Disposed. | [optional] 
**EEzsignfolderStep** | Pointer to [**FieldEEzsignfolderStep**](FieldEEzsignfolderStep.md) |  | [optional] 
**DtEzsignfolderClose** | Pointer to **string** | The date and time at which the Ezsignfolder was closed. Either by applying the last signature or by completing it prematurely. | [optional] 
**TEzsignfolderMessage** | Pointer to **string** | A custom text message that will be added to the email sent. | [optional] 
**ObjAudit** | Pointer to [**CommonAudit**](CommonAudit.md) |  | [optional] 
**SEzsignfolderExternalid** | Pointer to **string** | This field can be used to store an External ID from the client&#39;s system.  Anything can be stored in this field, it will never be evaluated by the eZmax system and will be returned AS-IS.  To store multiple values, consider using a JSON formatted structure, a URL encoded string, a CSV or any other custom format.  | [optional] 

## Methods

### NewEzsignfolderResponse

`func NewEzsignfolderResponse(pkiEzsignfolderID int32, eEzsignfolderCompletion FieldEEzsignfolderCompletion, sEzsignfolderDescription string, ) *EzsignfolderResponse`

NewEzsignfolderResponse instantiates a new EzsignfolderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsignfolderResponseWithDefaults

`func NewEzsignfolderResponseWithDefaults() *EzsignfolderResponse`

NewEzsignfolderResponseWithDefaults instantiates a new EzsignfolderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsignfolderID

`func (o *EzsignfolderResponse) GetPkiEzsignfolderID() int32`

GetPkiEzsignfolderID returns the PkiEzsignfolderID field if non-nil, zero value otherwise.

### GetPkiEzsignfolderIDOk

`func (o *EzsignfolderResponse) GetPkiEzsignfolderIDOk() (*int32, bool)`

GetPkiEzsignfolderIDOk returns a tuple with the PkiEzsignfolderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsignfolderID

`func (o *EzsignfolderResponse) SetPkiEzsignfolderID(v int32)`

SetPkiEzsignfolderID sets PkiEzsignfolderID field to given value.


### GetFkiEzsignfoldertypeID

`func (o *EzsignfolderResponse) GetFkiEzsignfoldertypeID() int32`

GetFkiEzsignfoldertypeID returns the FkiEzsignfoldertypeID field if non-nil, zero value otherwise.

### GetFkiEzsignfoldertypeIDOk

`func (o *EzsignfolderResponse) GetFkiEzsignfoldertypeIDOk() (*int32, bool)`

GetFkiEzsignfoldertypeIDOk returns a tuple with the FkiEzsignfoldertypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignfoldertypeID

`func (o *EzsignfolderResponse) SetFkiEzsignfoldertypeID(v int32)`

SetFkiEzsignfoldertypeID sets FkiEzsignfoldertypeID field to given value.

### HasFkiEzsignfoldertypeID

`func (o *EzsignfolderResponse) HasFkiEzsignfoldertypeID() bool`

HasFkiEzsignfoldertypeID returns a boolean if a field has been set.

### GetObjEzsignfoldertype

`func (o *EzsignfolderResponse) GetObjEzsignfoldertype() CustomEzsignfoldertypeResponse`

GetObjEzsignfoldertype returns the ObjEzsignfoldertype field if non-nil, zero value otherwise.

### GetObjEzsignfoldertypeOk

`func (o *EzsignfolderResponse) GetObjEzsignfoldertypeOk() (*CustomEzsignfoldertypeResponse, bool)`

GetObjEzsignfoldertypeOk returns a tuple with the ObjEzsignfoldertype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjEzsignfoldertype

`func (o *EzsignfolderResponse) SetObjEzsignfoldertype(v CustomEzsignfoldertypeResponse)`

SetObjEzsignfoldertype sets ObjEzsignfoldertype field to given value.

### HasObjEzsignfoldertype

`func (o *EzsignfolderResponse) HasObjEzsignfoldertype() bool`

HasObjEzsignfoldertype returns a boolean if a field has been set.

### GetEEzsignfolderCompletion

`func (o *EzsignfolderResponse) GetEEzsignfolderCompletion() FieldEEzsignfolderCompletion`

GetEEzsignfolderCompletion returns the EEzsignfolderCompletion field if non-nil, zero value otherwise.

### GetEEzsignfolderCompletionOk

`func (o *EzsignfolderResponse) GetEEzsignfolderCompletionOk() (*FieldEEzsignfolderCompletion, bool)`

GetEEzsignfolderCompletionOk returns a tuple with the EEzsignfolderCompletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignfolderCompletion

`func (o *EzsignfolderResponse) SetEEzsignfolderCompletion(v FieldEEzsignfolderCompletion)`

SetEEzsignfolderCompletion sets EEzsignfolderCompletion field to given value.


### GetSEzsignfoldertypeNameX

`func (o *EzsignfolderResponse) GetSEzsignfoldertypeNameX() string`

GetSEzsignfoldertypeNameX returns the SEzsignfoldertypeNameX field if non-nil, zero value otherwise.

### GetSEzsignfoldertypeNameXOk

`func (o *EzsignfolderResponse) GetSEzsignfoldertypeNameXOk() (*string, bool)`

GetSEzsignfoldertypeNameXOk returns a tuple with the SEzsignfoldertypeNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignfoldertypeNameX

`func (o *EzsignfolderResponse) SetSEzsignfoldertypeNameX(v string)`

SetSEzsignfoldertypeNameX sets SEzsignfoldertypeNameX field to given value.

### HasSEzsignfoldertypeNameX

`func (o *EzsignfolderResponse) HasSEzsignfoldertypeNameX() bool`

HasSEzsignfoldertypeNameX returns a boolean if a field has been set.

### GetFkiBillingentityinternalID

`func (o *EzsignfolderResponse) GetFkiBillingentityinternalID() int32`

GetFkiBillingentityinternalID returns the FkiBillingentityinternalID field if non-nil, zero value otherwise.

### GetFkiBillingentityinternalIDOk

`func (o *EzsignfolderResponse) GetFkiBillingentityinternalIDOk() (*int32, bool)`

GetFkiBillingentityinternalIDOk returns a tuple with the FkiBillingentityinternalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiBillingentityinternalID

`func (o *EzsignfolderResponse) SetFkiBillingentityinternalID(v int32)`

SetFkiBillingentityinternalID sets FkiBillingentityinternalID field to given value.

### HasFkiBillingentityinternalID

`func (o *EzsignfolderResponse) HasFkiBillingentityinternalID() bool`

HasFkiBillingentityinternalID returns a boolean if a field has been set.

### GetSBillingentityinternalDescriptionX

`func (o *EzsignfolderResponse) GetSBillingentityinternalDescriptionX() string`

GetSBillingentityinternalDescriptionX returns the SBillingentityinternalDescriptionX field if non-nil, zero value otherwise.

### GetSBillingentityinternalDescriptionXOk

`func (o *EzsignfolderResponse) GetSBillingentityinternalDescriptionXOk() (*string, bool)`

GetSBillingentityinternalDescriptionXOk returns a tuple with the SBillingentityinternalDescriptionX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSBillingentityinternalDescriptionX

`func (o *EzsignfolderResponse) SetSBillingentityinternalDescriptionX(v string)`

SetSBillingentityinternalDescriptionX sets SBillingentityinternalDescriptionX field to given value.

### HasSBillingentityinternalDescriptionX

`func (o *EzsignfolderResponse) HasSBillingentityinternalDescriptionX() bool`

HasSBillingentityinternalDescriptionX returns a boolean if a field has been set.

### GetFkiEzsigntsarequirementID

`func (o *EzsignfolderResponse) GetFkiEzsigntsarequirementID() int32`

GetFkiEzsigntsarequirementID returns the FkiEzsigntsarequirementID field if non-nil, zero value otherwise.

### GetFkiEzsigntsarequirementIDOk

`func (o *EzsignfolderResponse) GetFkiEzsigntsarequirementIDOk() (*int32, bool)`

GetFkiEzsigntsarequirementIDOk returns a tuple with the FkiEzsigntsarequirementID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigntsarequirementID

`func (o *EzsignfolderResponse) SetFkiEzsigntsarequirementID(v int32)`

SetFkiEzsigntsarequirementID sets FkiEzsigntsarequirementID field to given value.

### HasFkiEzsigntsarequirementID

`func (o *EzsignfolderResponse) HasFkiEzsigntsarequirementID() bool`

HasFkiEzsigntsarequirementID returns a boolean if a field has been set.

### GetSEzsigntsarequirementDescriptionX

`func (o *EzsignfolderResponse) GetSEzsigntsarequirementDescriptionX() string`

GetSEzsigntsarequirementDescriptionX returns the SEzsigntsarequirementDescriptionX field if non-nil, zero value otherwise.

### GetSEzsigntsarequirementDescriptionXOk

`func (o *EzsignfolderResponse) GetSEzsigntsarequirementDescriptionXOk() (*string, bool)`

GetSEzsigntsarequirementDescriptionXOk returns a tuple with the SEzsigntsarequirementDescriptionX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntsarequirementDescriptionX

`func (o *EzsignfolderResponse) SetSEzsigntsarequirementDescriptionX(v string)`

SetSEzsigntsarequirementDescriptionX sets SEzsigntsarequirementDescriptionX field to given value.

### HasSEzsigntsarequirementDescriptionX

`func (o *EzsignfolderResponse) HasSEzsigntsarequirementDescriptionX() bool`

HasSEzsigntsarequirementDescriptionX returns a boolean if a field has been set.

### GetSEzsignfolderDescription

`func (o *EzsignfolderResponse) GetSEzsignfolderDescription() string`

GetSEzsignfolderDescription returns the SEzsignfolderDescription field if non-nil, zero value otherwise.

### GetSEzsignfolderDescriptionOk

`func (o *EzsignfolderResponse) GetSEzsignfolderDescriptionOk() (*string, bool)`

GetSEzsignfolderDescriptionOk returns a tuple with the SEzsignfolderDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignfolderDescription

`func (o *EzsignfolderResponse) SetSEzsignfolderDescription(v string)`

SetSEzsignfolderDescription sets SEzsignfolderDescription field to given value.


### GetTEzsignfolderNote

`func (o *EzsignfolderResponse) GetTEzsignfolderNote() string`

GetTEzsignfolderNote returns the TEzsignfolderNote field if non-nil, zero value otherwise.

### GetTEzsignfolderNoteOk

`func (o *EzsignfolderResponse) GetTEzsignfolderNoteOk() (*string, bool)`

GetTEzsignfolderNoteOk returns a tuple with the TEzsignfolderNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTEzsignfolderNote

`func (o *EzsignfolderResponse) SetTEzsignfolderNote(v string)`

SetTEzsignfolderNote sets TEzsignfolderNote field to given value.

### HasTEzsignfolderNote

`func (o *EzsignfolderResponse) HasTEzsignfolderNote() bool`

HasTEzsignfolderNote returns a boolean if a field has been set.

### GetBEzsignfolderIsdisposable

`func (o *EzsignfolderResponse) GetBEzsignfolderIsdisposable() bool`

GetBEzsignfolderIsdisposable returns the BEzsignfolderIsdisposable field if non-nil, zero value otherwise.

### GetBEzsignfolderIsdisposableOk

`func (o *EzsignfolderResponse) GetBEzsignfolderIsdisposableOk() (*bool, bool)`

GetBEzsignfolderIsdisposableOk returns a tuple with the BEzsignfolderIsdisposable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfolderIsdisposable

`func (o *EzsignfolderResponse) SetBEzsignfolderIsdisposable(v bool)`

SetBEzsignfolderIsdisposable sets BEzsignfolderIsdisposable field to given value.

### HasBEzsignfolderIsdisposable

`func (o *EzsignfolderResponse) HasBEzsignfolderIsdisposable() bool`

HasBEzsignfolderIsdisposable returns a boolean if a field has been set.

### GetEEzsignfolderSendreminderfrequency

`func (o *EzsignfolderResponse) GetEEzsignfolderSendreminderfrequency() FieldEEzsignfolderSendreminderfrequency`

GetEEzsignfolderSendreminderfrequency returns the EEzsignfolderSendreminderfrequency field if non-nil, zero value otherwise.

### GetEEzsignfolderSendreminderfrequencyOk

`func (o *EzsignfolderResponse) GetEEzsignfolderSendreminderfrequencyOk() (*FieldEEzsignfolderSendreminderfrequency, bool)`

GetEEzsignfolderSendreminderfrequencyOk returns a tuple with the EEzsignfolderSendreminderfrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignfolderSendreminderfrequency

`func (o *EzsignfolderResponse) SetEEzsignfolderSendreminderfrequency(v FieldEEzsignfolderSendreminderfrequency)`

SetEEzsignfolderSendreminderfrequency sets EEzsignfolderSendreminderfrequency field to given value.

### HasEEzsignfolderSendreminderfrequency

`func (o *EzsignfolderResponse) HasEEzsignfolderSendreminderfrequency() bool`

HasEEzsignfolderSendreminderfrequency returns a boolean if a field has been set.

### GetDtEzsignfolderDelayedsenddate

`func (o *EzsignfolderResponse) GetDtEzsignfolderDelayedsenddate() string`

GetDtEzsignfolderDelayedsenddate returns the DtEzsignfolderDelayedsenddate field if non-nil, zero value otherwise.

### GetDtEzsignfolderDelayedsenddateOk

`func (o *EzsignfolderResponse) GetDtEzsignfolderDelayedsenddateOk() (*string, bool)`

GetDtEzsignfolderDelayedsenddateOk returns a tuple with the DtEzsignfolderDelayedsenddate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtEzsignfolderDelayedsenddate

`func (o *EzsignfolderResponse) SetDtEzsignfolderDelayedsenddate(v string)`

SetDtEzsignfolderDelayedsenddate sets DtEzsignfolderDelayedsenddate field to given value.

### HasDtEzsignfolderDelayedsenddate

`func (o *EzsignfolderResponse) HasDtEzsignfolderDelayedsenddate() bool`

HasDtEzsignfolderDelayedsenddate returns a boolean if a field has been set.

### GetDtEzsignfolderDuedate

`func (o *EzsignfolderResponse) GetDtEzsignfolderDuedate() string`

GetDtEzsignfolderDuedate returns the DtEzsignfolderDuedate field if non-nil, zero value otherwise.

### GetDtEzsignfolderDuedateOk

`func (o *EzsignfolderResponse) GetDtEzsignfolderDuedateOk() (*string, bool)`

GetDtEzsignfolderDuedateOk returns a tuple with the DtEzsignfolderDuedate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtEzsignfolderDuedate

`func (o *EzsignfolderResponse) SetDtEzsignfolderDuedate(v string)`

SetDtEzsignfolderDuedate sets DtEzsignfolderDuedate field to given value.

### HasDtEzsignfolderDuedate

`func (o *EzsignfolderResponse) HasDtEzsignfolderDuedate() bool`

HasDtEzsignfolderDuedate returns a boolean if a field has been set.

### GetDtEzsignfolderSentdate

`func (o *EzsignfolderResponse) GetDtEzsignfolderSentdate() string`

GetDtEzsignfolderSentdate returns the DtEzsignfolderSentdate field if non-nil, zero value otherwise.

### GetDtEzsignfolderSentdateOk

`func (o *EzsignfolderResponse) GetDtEzsignfolderSentdateOk() (*string, bool)`

GetDtEzsignfolderSentdateOk returns a tuple with the DtEzsignfolderSentdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtEzsignfolderSentdate

`func (o *EzsignfolderResponse) SetDtEzsignfolderSentdate(v string)`

SetDtEzsignfolderSentdate sets DtEzsignfolderSentdate field to given value.

### HasDtEzsignfolderSentdate

`func (o *EzsignfolderResponse) HasDtEzsignfolderSentdate() bool`

HasDtEzsignfolderSentdate returns a boolean if a field has been set.

### GetDtEzsignfolderScheduledarchive

`func (o *EzsignfolderResponse) GetDtEzsignfolderScheduledarchive() string`

GetDtEzsignfolderScheduledarchive returns the DtEzsignfolderScheduledarchive field if non-nil, zero value otherwise.

### GetDtEzsignfolderScheduledarchiveOk

`func (o *EzsignfolderResponse) GetDtEzsignfolderScheduledarchiveOk() (*string, bool)`

GetDtEzsignfolderScheduledarchiveOk returns a tuple with the DtEzsignfolderScheduledarchive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtEzsignfolderScheduledarchive

`func (o *EzsignfolderResponse) SetDtEzsignfolderScheduledarchive(v string)`

SetDtEzsignfolderScheduledarchive sets DtEzsignfolderScheduledarchive field to given value.

### HasDtEzsignfolderScheduledarchive

`func (o *EzsignfolderResponse) HasDtEzsignfolderScheduledarchive() bool`

HasDtEzsignfolderScheduledarchive returns a boolean if a field has been set.

### GetDtEzsignfolderScheduleddispose

`func (o *EzsignfolderResponse) GetDtEzsignfolderScheduleddispose() string`

GetDtEzsignfolderScheduleddispose returns the DtEzsignfolderScheduleddispose field if non-nil, zero value otherwise.

### GetDtEzsignfolderScheduleddisposeOk

`func (o *EzsignfolderResponse) GetDtEzsignfolderScheduleddisposeOk() (*string, bool)`

GetDtEzsignfolderScheduleddisposeOk returns a tuple with the DtEzsignfolderScheduleddispose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtEzsignfolderScheduleddispose

`func (o *EzsignfolderResponse) SetDtEzsignfolderScheduleddispose(v string)`

SetDtEzsignfolderScheduleddispose sets DtEzsignfolderScheduleddispose field to given value.

### HasDtEzsignfolderScheduleddispose

`func (o *EzsignfolderResponse) HasDtEzsignfolderScheduleddispose() bool`

HasDtEzsignfolderScheduleddispose returns a boolean if a field has been set.

### GetEEzsignfolderStep

`func (o *EzsignfolderResponse) GetEEzsignfolderStep() FieldEEzsignfolderStep`

GetEEzsignfolderStep returns the EEzsignfolderStep field if non-nil, zero value otherwise.

### GetEEzsignfolderStepOk

`func (o *EzsignfolderResponse) GetEEzsignfolderStepOk() (*FieldEEzsignfolderStep, bool)`

GetEEzsignfolderStepOk returns a tuple with the EEzsignfolderStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignfolderStep

`func (o *EzsignfolderResponse) SetEEzsignfolderStep(v FieldEEzsignfolderStep)`

SetEEzsignfolderStep sets EEzsignfolderStep field to given value.

### HasEEzsignfolderStep

`func (o *EzsignfolderResponse) HasEEzsignfolderStep() bool`

HasEEzsignfolderStep returns a boolean if a field has been set.

### GetDtEzsignfolderClose

`func (o *EzsignfolderResponse) GetDtEzsignfolderClose() string`

GetDtEzsignfolderClose returns the DtEzsignfolderClose field if non-nil, zero value otherwise.

### GetDtEzsignfolderCloseOk

`func (o *EzsignfolderResponse) GetDtEzsignfolderCloseOk() (*string, bool)`

GetDtEzsignfolderCloseOk returns a tuple with the DtEzsignfolderClose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtEzsignfolderClose

`func (o *EzsignfolderResponse) SetDtEzsignfolderClose(v string)`

SetDtEzsignfolderClose sets DtEzsignfolderClose field to given value.

### HasDtEzsignfolderClose

`func (o *EzsignfolderResponse) HasDtEzsignfolderClose() bool`

HasDtEzsignfolderClose returns a boolean if a field has been set.

### GetTEzsignfolderMessage

`func (o *EzsignfolderResponse) GetTEzsignfolderMessage() string`

GetTEzsignfolderMessage returns the TEzsignfolderMessage field if non-nil, zero value otherwise.

### GetTEzsignfolderMessageOk

`func (o *EzsignfolderResponse) GetTEzsignfolderMessageOk() (*string, bool)`

GetTEzsignfolderMessageOk returns a tuple with the TEzsignfolderMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTEzsignfolderMessage

`func (o *EzsignfolderResponse) SetTEzsignfolderMessage(v string)`

SetTEzsignfolderMessage sets TEzsignfolderMessage field to given value.

### HasTEzsignfolderMessage

`func (o *EzsignfolderResponse) HasTEzsignfolderMessage() bool`

HasTEzsignfolderMessage returns a boolean if a field has been set.

### GetObjAudit

`func (o *EzsignfolderResponse) GetObjAudit() CommonAudit`

GetObjAudit returns the ObjAudit field if non-nil, zero value otherwise.

### GetObjAuditOk

`func (o *EzsignfolderResponse) GetObjAuditOk() (*CommonAudit, bool)`

GetObjAuditOk returns a tuple with the ObjAudit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjAudit

`func (o *EzsignfolderResponse) SetObjAudit(v CommonAudit)`

SetObjAudit sets ObjAudit field to given value.

### HasObjAudit

`func (o *EzsignfolderResponse) HasObjAudit() bool`

HasObjAudit returns a boolean if a field has been set.

### GetSEzsignfolderExternalid

`func (o *EzsignfolderResponse) GetSEzsignfolderExternalid() string`

GetSEzsignfolderExternalid returns the SEzsignfolderExternalid field if non-nil, zero value otherwise.

### GetSEzsignfolderExternalidOk

`func (o *EzsignfolderResponse) GetSEzsignfolderExternalidOk() (*string, bool)`

GetSEzsignfolderExternalidOk returns a tuple with the SEzsignfolderExternalid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignfolderExternalid

`func (o *EzsignfolderResponse) SetSEzsignfolderExternalid(v string)`

SetSEzsignfolderExternalid sets SEzsignfolderExternalid field to given value.

### HasSEzsignfolderExternalid

`func (o *EzsignfolderResponse) HasSEzsignfolderExternalid() bool`

HasSEzsignfolderExternalid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


