# EzsignfolderResponseCompoundV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsignfolderID** | **int32** | The unique ID of the Ezsignfolder | 
**FkiEzsignfoldertypeID** | Pointer to **int32** | The unique ID of the Ezsignfoldertype. | [optional] 
**ObjEzsignfoldertype** | Pointer to [**CustomEzsignfoldertypeResponse**](CustomEzsignfoldertypeResponse.md) |  | [optional] 
**FkiTimezoneID** | Pointer to **int32** | The unique ID of the Timezone | [optional] 
**EEzsignfolderCompletion** | [**FieldEEzsignfolderCompletion**](FieldEEzsignfolderCompletion.md) |  | 
**EEzsignfolderDocumentdependency** | Pointer to [**FieldEEzsignfolderDocumentdependency**](FieldEEzsignfolderDocumentdependency.md) |  | [optional] 
**SEzsignfoldertypeNameX** | Pointer to **string** |  | [optional] 
**FkiBillingentityinternalID** | Pointer to **int32** | The unique ID of the Billingentityinternal. | [optional] 
**SBillingentityinternalDescriptionX** | Pointer to **string** | The description of the Billingentityinternal in the language of the requester | [optional] 
**FkiEzsigntsarequirementID** | Pointer to **int32** | The unique ID of the Ezsigntsarequirement.  Determine if a Time Stamping Authority should add a timestamp on each of the signature. Valid values:  |Value|Description| |-|-| |1|No. TSA Timestamping will requested. This will make all signatures a lot faster since no round-trip to the TSA server will be required. Timestamping will be made using eZsign server&#39;s time.| |2|Best effort. Timestamping from a Time Stamping Authority will be requested but is not mandatory. In the very improbable case it cannot be completed, the timestamping will be made using eZsign server&#39;s time. **Additional fee applies**| |3|Mandatory. Timestamping from a Time Stamping Authority will be requested and is mandatory. In the very improbable case it cannot be completed, the signature will fail and the user will be asked to retry. **Additional fee applies**| | [optional] 
**SEzsigntsarequirementDescriptionX** | Pointer to **string** | The description of the Ezsigntsarequirement in the language of the requester | [optional] 
**SEzsignfolderDescription** | **string** | The description of the Ezsignfolder | 
**TEzsignfolderNote** | Pointer to **string** | Note about the Ezsignfolder | [optional] 
**BEzsignfolderIsdisposable** | Pointer to **bool** | If the Ezsigndocument can be disposed | [optional] 
**IEzsignfolderSendreminderfirstdays** | Pointer to **int32** | The number of days before the the first reminder sending | [optional] 
**IEzsignfolderSendreminderotherdays** | Pointer to **int32** | The number of days after the first reminder sending | [optional] 
**DtEzsignfolderDelayedsenddate** | Pointer to **string** | The date and time at which the Ezsignfolder will be sent in the future. | [optional] 
**DtEzsignfolderDuedate** | Pointer to **string** | The maximum date and time at which the Ezsignfolder can be signed. | [optional] 
**DtEzsignfolderSentdate** | Pointer to **string** | The date and time at which the Ezsignfolder was sent the last time. | [optional] 
**DtEzsignfolderScheduledarchive** | Pointer to **string** | The scheduled date and time at which the Ezsignfolder should be archived. | [optional] 
**DtEzsignfolderScheduleddispose** | Pointer to **string** | The scheduled date at which the Ezsignfolder should be Disposed. | [optional] 
**EEzsignfolderStep** | Pointer to [**FieldEEzsignfolderStep**](FieldEEzsignfolderStep.md) |  | [optional] 
**DtEzsignfolderClose** | Pointer to **string** | The date and time at which the Ezsignfolder was closed. Either by applying the last signature or by completing it prematurely. | [optional] 
**DtEzsignfolderArchive** | Pointer to **string** | The date and time at which the Ezsignfolder was archived. | [optional] 
**DtEzsignfolderDispose** | Pointer to **string** | The date and time at which the Ezsignfolder was disposed. | [optional] 
**TEzsignfolderMessage** | Pointer to **string** | A custom text message that will be added to the email sent. | [optional] 
**ObjAudit** | Pointer to [**CommonAudit**](CommonAudit.md) |  | [optional] 
**SEzsignfolderExternalid** | Pointer to **string** | This field can be used to store an External ID from the client&#39;s system.  Anything can be stored in this field, it will never be evaluated by the eZmax system and will be returned AS-IS.  To store multiple values, consider using a JSON formatted structure, a URL encoded string, a CSV or any other custom format.  | [optional] 
**EEzsignfolderAccess** | Pointer to [**ComputedEEzsignfolderAccess**](ComputedEEzsignfolderAccess.md) |  | [optional] 
**ObjTimezone** | Pointer to [**CustomTimezoneWithCodeResponse**](CustomTimezoneWithCodeResponse.md) |  | [optional] 

## Methods

### NewEzsignfolderResponseCompoundV3

`func NewEzsignfolderResponseCompoundV3(pkiEzsignfolderID int32, eEzsignfolderCompletion FieldEEzsignfolderCompletion, sEzsignfolderDescription string, ) *EzsignfolderResponseCompoundV3`

NewEzsignfolderResponseCompoundV3 instantiates a new EzsignfolderResponseCompoundV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsignfolderResponseCompoundV3WithDefaults

`func NewEzsignfolderResponseCompoundV3WithDefaults() *EzsignfolderResponseCompoundV3`

NewEzsignfolderResponseCompoundV3WithDefaults instantiates a new EzsignfolderResponseCompoundV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsignfolderID

`func (o *EzsignfolderResponseCompoundV3) GetPkiEzsignfolderID() int32`

GetPkiEzsignfolderID returns the PkiEzsignfolderID field if non-nil, zero value otherwise.

### GetPkiEzsignfolderIDOk

`func (o *EzsignfolderResponseCompoundV3) GetPkiEzsignfolderIDOk() (*int32, bool)`

GetPkiEzsignfolderIDOk returns a tuple with the PkiEzsignfolderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsignfolderID

`func (o *EzsignfolderResponseCompoundV3) SetPkiEzsignfolderID(v int32)`

SetPkiEzsignfolderID sets PkiEzsignfolderID field to given value.


### GetFkiEzsignfoldertypeID

`func (o *EzsignfolderResponseCompoundV3) GetFkiEzsignfoldertypeID() int32`

GetFkiEzsignfoldertypeID returns the FkiEzsignfoldertypeID field if non-nil, zero value otherwise.

### GetFkiEzsignfoldertypeIDOk

`func (o *EzsignfolderResponseCompoundV3) GetFkiEzsignfoldertypeIDOk() (*int32, bool)`

GetFkiEzsignfoldertypeIDOk returns a tuple with the FkiEzsignfoldertypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignfoldertypeID

`func (o *EzsignfolderResponseCompoundV3) SetFkiEzsignfoldertypeID(v int32)`

SetFkiEzsignfoldertypeID sets FkiEzsignfoldertypeID field to given value.

### HasFkiEzsignfoldertypeID

`func (o *EzsignfolderResponseCompoundV3) HasFkiEzsignfoldertypeID() bool`

HasFkiEzsignfoldertypeID returns a boolean if a field has been set.

### GetObjEzsignfoldertype

`func (o *EzsignfolderResponseCompoundV3) GetObjEzsignfoldertype() CustomEzsignfoldertypeResponse`

GetObjEzsignfoldertype returns the ObjEzsignfoldertype field if non-nil, zero value otherwise.

### GetObjEzsignfoldertypeOk

`func (o *EzsignfolderResponseCompoundV3) GetObjEzsignfoldertypeOk() (*CustomEzsignfoldertypeResponse, bool)`

GetObjEzsignfoldertypeOk returns a tuple with the ObjEzsignfoldertype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjEzsignfoldertype

`func (o *EzsignfolderResponseCompoundV3) SetObjEzsignfoldertype(v CustomEzsignfoldertypeResponse)`

SetObjEzsignfoldertype sets ObjEzsignfoldertype field to given value.

### HasObjEzsignfoldertype

`func (o *EzsignfolderResponseCompoundV3) HasObjEzsignfoldertype() bool`

HasObjEzsignfoldertype returns a boolean if a field has been set.

### GetFkiTimezoneID

`func (o *EzsignfolderResponseCompoundV3) GetFkiTimezoneID() int32`

GetFkiTimezoneID returns the FkiTimezoneID field if non-nil, zero value otherwise.

### GetFkiTimezoneIDOk

`func (o *EzsignfolderResponseCompoundV3) GetFkiTimezoneIDOk() (*int32, bool)`

GetFkiTimezoneIDOk returns a tuple with the FkiTimezoneID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiTimezoneID

`func (o *EzsignfolderResponseCompoundV3) SetFkiTimezoneID(v int32)`

SetFkiTimezoneID sets FkiTimezoneID field to given value.

### HasFkiTimezoneID

`func (o *EzsignfolderResponseCompoundV3) HasFkiTimezoneID() bool`

HasFkiTimezoneID returns a boolean if a field has been set.

### GetEEzsignfolderCompletion

`func (o *EzsignfolderResponseCompoundV3) GetEEzsignfolderCompletion() FieldEEzsignfolderCompletion`

GetEEzsignfolderCompletion returns the EEzsignfolderCompletion field if non-nil, zero value otherwise.

### GetEEzsignfolderCompletionOk

`func (o *EzsignfolderResponseCompoundV3) GetEEzsignfolderCompletionOk() (*FieldEEzsignfolderCompletion, bool)`

GetEEzsignfolderCompletionOk returns a tuple with the EEzsignfolderCompletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignfolderCompletion

`func (o *EzsignfolderResponseCompoundV3) SetEEzsignfolderCompletion(v FieldEEzsignfolderCompletion)`

SetEEzsignfolderCompletion sets EEzsignfolderCompletion field to given value.


### GetEEzsignfolderDocumentdependency

`func (o *EzsignfolderResponseCompoundV3) GetEEzsignfolderDocumentdependency() FieldEEzsignfolderDocumentdependency`

GetEEzsignfolderDocumentdependency returns the EEzsignfolderDocumentdependency field if non-nil, zero value otherwise.

### GetEEzsignfolderDocumentdependencyOk

`func (o *EzsignfolderResponseCompoundV3) GetEEzsignfolderDocumentdependencyOk() (*FieldEEzsignfolderDocumentdependency, bool)`

GetEEzsignfolderDocumentdependencyOk returns a tuple with the EEzsignfolderDocumentdependency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignfolderDocumentdependency

`func (o *EzsignfolderResponseCompoundV3) SetEEzsignfolderDocumentdependency(v FieldEEzsignfolderDocumentdependency)`

SetEEzsignfolderDocumentdependency sets EEzsignfolderDocumentdependency field to given value.

### HasEEzsignfolderDocumentdependency

`func (o *EzsignfolderResponseCompoundV3) HasEEzsignfolderDocumentdependency() bool`

HasEEzsignfolderDocumentdependency returns a boolean if a field has been set.

### GetSEzsignfoldertypeNameX

`func (o *EzsignfolderResponseCompoundV3) GetSEzsignfoldertypeNameX() string`

GetSEzsignfoldertypeNameX returns the SEzsignfoldertypeNameX field if non-nil, zero value otherwise.

### GetSEzsignfoldertypeNameXOk

`func (o *EzsignfolderResponseCompoundV3) GetSEzsignfoldertypeNameXOk() (*string, bool)`

GetSEzsignfoldertypeNameXOk returns a tuple with the SEzsignfoldertypeNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignfoldertypeNameX

`func (o *EzsignfolderResponseCompoundV3) SetSEzsignfoldertypeNameX(v string)`

SetSEzsignfoldertypeNameX sets SEzsignfoldertypeNameX field to given value.

### HasSEzsignfoldertypeNameX

`func (o *EzsignfolderResponseCompoundV3) HasSEzsignfoldertypeNameX() bool`

HasSEzsignfoldertypeNameX returns a boolean if a field has been set.

### GetFkiBillingentityinternalID

`func (o *EzsignfolderResponseCompoundV3) GetFkiBillingentityinternalID() int32`

GetFkiBillingentityinternalID returns the FkiBillingentityinternalID field if non-nil, zero value otherwise.

### GetFkiBillingentityinternalIDOk

`func (o *EzsignfolderResponseCompoundV3) GetFkiBillingentityinternalIDOk() (*int32, bool)`

GetFkiBillingentityinternalIDOk returns a tuple with the FkiBillingentityinternalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiBillingentityinternalID

`func (o *EzsignfolderResponseCompoundV3) SetFkiBillingentityinternalID(v int32)`

SetFkiBillingentityinternalID sets FkiBillingentityinternalID field to given value.

### HasFkiBillingentityinternalID

`func (o *EzsignfolderResponseCompoundV3) HasFkiBillingentityinternalID() bool`

HasFkiBillingentityinternalID returns a boolean if a field has been set.

### GetSBillingentityinternalDescriptionX

`func (o *EzsignfolderResponseCompoundV3) GetSBillingentityinternalDescriptionX() string`

GetSBillingentityinternalDescriptionX returns the SBillingentityinternalDescriptionX field if non-nil, zero value otherwise.

### GetSBillingentityinternalDescriptionXOk

`func (o *EzsignfolderResponseCompoundV3) GetSBillingentityinternalDescriptionXOk() (*string, bool)`

GetSBillingentityinternalDescriptionXOk returns a tuple with the SBillingentityinternalDescriptionX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSBillingentityinternalDescriptionX

`func (o *EzsignfolderResponseCompoundV3) SetSBillingentityinternalDescriptionX(v string)`

SetSBillingentityinternalDescriptionX sets SBillingentityinternalDescriptionX field to given value.

### HasSBillingentityinternalDescriptionX

`func (o *EzsignfolderResponseCompoundV3) HasSBillingentityinternalDescriptionX() bool`

HasSBillingentityinternalDescriptionX returns a boolean if a field has been set.

### GetFkiEzsigntsarequirementID

`func (o *EzsignfolderResponseCompoundV3) GetFkiEzsigntsarequirementID() int32`

GetFkiEzsigntsarequirementID returns the FkiEzsigntsarequirementID field if non-nil, zero value otherwise.

### GetFkiEzsigntsarequirementIDOk

`func (o *EzsignfolderResponseCompoundV3) GetFkiEzsigntsarequirementIDOk() (*int32, bool)`

GetFkiEzsigntsarequirementIDOk returns a tuple with the FkiEzsigntsarequirementID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigntsarequirementID

`func (o *EzsignfolderResponseCompoundV3) SetFkiEzsigntsarequirementID(v int32)`

SetFkiEzsigntsarequirementID sets FkiEzsigntsarequirementID field to given value.

### HasFkiEzsigntsarequirementID

`func (o *EzsignfolderResponseCompoundV3) HasFkiEzsigntsarequirementID() bool`

HasFkiEzsigntsarequirementID returns a boolean if a field has been set.

### GetSEzsigntsarequirementDescriptionX

`func (o *EzsignfolderResponseCompoundV3) GetSEzsigntsarequirementDescriptionX() string`

GetSEzsigntsarequirementDescriptionX returns the SEzsigntsarequirementDescriptionX field if non-nil, zero value otherwise.

### GetSEzsigntsarequirementDescriptionXOk

`func (o *EzsignfolderResponseCompoundV3) GetSEzsigntsarequirementDescriptionXOk() (*string, bool)`

GetSEzsigntsarequirementDescriptionXOk returns a tuple with the SEzsigntsarequirementDescriptionX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntsarequirementDescriptionX

`func (o *EzsignfolderResponseCompoundV3) SetSEzsigntsarequirementDescriptionX(v string)`

SetSEzsigntsarequirementDescriptionX sets SEzsigntsarequirementDescriptionX field to given value.

### HasSEzsigntsarequirementDescriptionX

`func (o *EzsignfolderResponseCompoundV3) HasSEzsigntsarequirementDescriptionX() bool`

HasSEzsigntsarequirementDescriptionX returns a boolean if a field has been set.

### GetSEzsignfolderDescription

`func (o *EzsignfolderResponseCompoundV3) GetSEzsignfolderDescription() string`

GetSEzsignfolderDescription returns the SEzsignfolderDescription field if non-nil, zero value otherwise.

### GetSEzsignfolderDescriptionOk

`func (o *EzsignfolderResponseCompoundV3) GetSEzsignfolderDescriptionOk() (*string, bool)`

GetSEzsignfolderDescriptionOk returns a tuple with the SEzsignfolderDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignfolderDescription

`func (o *EzsignfolderResponseCompoundV3) SetSEzsignfolderDescription(v string)`

SetSEzsignfolderDescription sets SEzsignfolderDescription field to given value.


### GetTEzsignfolderNote

`func (o *EzsignfolderResponseCompoundV3) GetTEzsignfolderNote() string`

GetTEzsignfolderNote returns the TEzsignfolderNote field if non-nil, zero value otherwise.

### GetTEzsignfolderNoteOk

`func (o *EzsignfolderResponseCompoundV3) GetTEzsignfolderNoteOk() (*string, bool)`

GetTEzsignfolderNoteOk returns a tuple with the TEzsignfolderNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTEzsignfolderNote

`func (o *EzsignfolderResponseCompoundV3) SetTEzsignfolderNote(v string)`

SetTEzsignfolderNote sets TEzsignfolderNote field to given value.

### HasTEzsignfolderNote

`func (o *EzsignfolderResponseCompoundV3) HasTEzsignfolderNote() bool`

HasTEzsignfolderNote returns a boolean if a field has been set.

### GetBEzsignfolderIsdisposable

`func (o *EzsignfolderResponseCompoundV3) GetBEzsignfolderIsdisposable() bool`

GetBEzsignfolderIsdisposable returns the BEzsignfolderIsdisposable field if non-nil, zero value otherwise.

### GetBEzsignfolderIsdisposableOk

`func (o *EzsignfolderResponseCompoundV3) GetBEzsignfolderIsdisposableOk() (*bool, bool)`

GetBEzsignfolderIsdisposableOk returns a tuple with the BEzsignfolderIsdisposable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfolderIsdisposable

`func (o *EzsignfolderResponseCompoundV3) SetBEzsignfolderIsdisposable(v bool)`

SetBEzsignfolderIsdisposable sets BEzsignfolderIsdisposable field to given value.

### HasBEzsignfolderIsdisposable

`func (o *EzsignfolderResponseCompoundV3) HasBEzsignfolderIsdisposable() bool`

HasBEzsignfolderIsdisposable returns a boolean if a field has been set.

### GetIEzsignfolderSendreminderfirstdays

`func (o *EzsignfolderResponseCompoundV3) GetIEzsignfolderSendreminderfirstdays() int32`

GetIEzsignfolderSendreminderfirstdays returns the IEzsignfolderSendreminderfirstdays field if non-nil, zero value otherwise.

### GetIEzsignfolderSendreminderfirstdaysOk

`func (o *EzsignfolderResponseCompoundV3) GetIEzsignfolderSendreminderfirstdaysOk() (*int32, bool)`

GetIEzsignfolderSendreminderfirstdaysOk returns a tuple with the IEzsignfolderSendreminderfirstdays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignfolderSendreminderfirstdays

`func (o *EzsignfolderResponseCompoundV3) SetIEzsignfolderSendreminderfirstdays(v int32)`

SetIEzsignfolderSendreminderfirstdays sets IEzsignfolderSendreminderfirstdays field to given value.

### HasIEzsignfolderSendreminderfirstdays

`func (o *EzsignfolderResponseCompoundV3) HasIEzsignfolderSendreminderfirstdays() bool`

HasIEzsignfolderSendreminderfirstdays returns a boolean if a field has been set.

### GetIEzsignfolderSendreminderotherdays

`func (o *EzsignfolderResponseCompoundV3) GetIEzsignfolderSendreminderotherdays() int32`

GetIEzsignfolderSendreminderotherdays returns the IEzsignfolderSendreminderotherdays field if non-nil, zero value otherwise.

### GetIEzsignfolderSendreminderotherdaysOk

`func (o *EzsignfolderResponseCompoundV3) GetIEzsignfolderSendreminderotherdaysOk() (*int32, bool)`

GetIEzsignfolderSendreminderotherdaysOk returns a tuple with the IEzsignfolderSendreminderotherdays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignfolderSendreminderotherdays

`func (o *EzsignfolderResponseCompoundV3) SetIEzsignfolderSendreminderotherdays(v int32)`

SetIEzsignfolderSendreminderotherdays sets IEzsignfolderSendreminderotherdays field to given value.

### HasIEzsignfolderSendreminderotherdays

`func (o *EzsignfolderResponseCompoundV3) HasIEzsignfolderSendreminderotherdays() bool`

HasIEzsignfolderSendreminderotherdays returns a boolean if a field has been set.

### GetDtEzsignfolderDelayedsenddate

`func (o *EzsignfolderResponseCompoundV3) GetDtEzsignfolderDelayedsenddate() string`

GetDtEzsignfolderDelayedsenddate returns the DtEzsignfolderDelayedsenddate field if non-nil, zero value otherwise.

### GetDtEzsignfolderDelayedsenddateOk

`func (o *EzsignfolderResponseCompoundV3) GetDtEzsignfolderDelayedsenddateOk() (*string, bool)`

GetDtEzsignfolderDelayedsenddateOk returns a tuple with the DtEzsignfolderDelayedsenddate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtEzsignfolderDelayedsenddate

`func (o *EzsignfolderResponseCompoundV3) SetDtEzsignfolderDelayedsenddate(v string)`

SetDtEzsignfolderDelayedsenddate sets DtEzsignfolderDelayedsenddate field to given value.

### HasDtEzsignfolderDelayedsenddate

`func (o *EzsignfolderResponseCompoundV3) HasDtEzsignfolderDelayedsenddate() bool`

HasDtEzsignfolderDelayedsenddate returns a boolean if a field has been set.

### GetDtEzsignfolderDuedate

`func (o *EzsignfolderResponseCompoundV3) GetDtEzsignfolderDuedate() string`

GetDtEzsignfolderDuedate returns the DtEzsignfolderDuedate field if non-nil, zero value otherwise.

### GetDtEzsignfolderDuedateOk

`func (o *EzsignfolderResponseCompoundV3) GetDtEzsignfolderDuedateOk() (*string, bool)`

GetDtEzsignfolderDuedateOk returns a tuple with the DtEzsignfolderDuedate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtEzsignfolderDuedate

`func (o *EzsignfolderResponseCompoundV3) SetDtEzsignfolderDuedate(v string)`

SetDtEzsignfolderDuedate sets DtEzsignfolderDuedate field to given value.

### HasDtEzsignfolderDuedate

`func (o *EzsignfolderResponseCompoundV3) HasDtEzsignfolderDuedate() bool`

HasDtEzsignfolderDuedate returns a boolean if a field has been set.

### GetDtEzsignfolderSentdate

`func (o *EzsignfolderResponseCompoundV3) GetDtEzsignfolderSentdate() string`

GetDtEzsignfolderSentdate returns the DtEzsignfolderSentdate field if non-nil, zero value otherwise.

### GetDtEzsignfolderSentdateOk

`func (o *EzsignfolderResponseCompoundV3) GetDtEzsignfolderSentdateOk() (*string, bool)`

GetDtEzsignfolderSentdateOk returns a tuple with the DtEzsignfolderSentdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtEzsignfolderSentdate

`func (o *EzsignfolderResponseCompoundV3) SetDtEzsignfolderSentdate(v string)`

SetDtEzsignfolderSentdate sets DtEzsignfolderSentdate field to given value.

### HasDtEzsignfolderSentdate

`func (o *EzsignfolderResponseCompoundV3) HasDtEzsignfolderSentdate() bool`

HasDtEzsignfolderSentdate returns a boolean if a field has been set.

### GetDtEzsignfolderScheduledarchive

`func (o *EzsignfolderResponseCompoundV3) GetDtEzsignfolderScheduledarchive() string`

GetDtEzsignfolderScheduledarchive returns the DtEzsignfolderScheduledarchive field if non-nil, zero value otherwise.

### GetDtEzsignfolderScheduledarchiveOk

`func (o *EzsignfolderResponseCompoundV3) GetDtEzsignfolderScheduledarchiveOk() (*string, bool)`

GetDtEzsignfolderScheduledarchiveOk returns a tuple with the DtEzsignfolderScheduledarchive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtEzsignfolderScheduledarchive

`func (o *EzsignfolderResponseCompoundV3) SetDtEzsignfolderScheduledarchive(v string)`

SetDtEzsignfolderScheduledarchive sets DtEzsignfolderScheduledarchive field to given value.

### HasDtEzsignfolderScheduledarchive

`func (o *EzsignfolderResponseCompoundV3) HasDtEzsignfolderScheduledarchive() bool`

HasDtEzsignfolderScheduledarchive returns a boolean if a field has been set.

### GetDtEzsignfolderScheduleddispose

`func (o *EzsignfolderResponseCompoundV3) GetDtEzsignfolderScheduleddispose() string`

GetDtEzsignfolderScheduleddispose returns the DtEzsignfolderScheduleddispose field if non-nil, zero value otherwise.

### GetDtEzsignfolderScheduleddisposeOk

`func (o *EzsignfolderResponseCompoundV3) GetDtEzsignfolderScheduleddisposeOk() (*string, bool)`

GetDtEzsignfolderScheduleddisposeOk returns a tuple with the DtEzsignfolderScheduleddispose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtEzsignfolderScheduleddispose

`func (o *EzsignfolderResponseCompoundV3) SetDtEzsignfolderScheduleddispose(v string)`

SetDtEzsignfolderScheduleddispose sets DtEzsignfolderScheduleddispose field to given value.

### HasDtEzsignfolderScheduleddispose

`func (o *EzsignfolderResponseCompoundV3) HasDtEzsignfolderScheduleddispose() bool`

HasDtEzsignfolderScheduleddispose returns a boolean if a field has been set.

### GetEEzsignfolderStep

`func (o *EzsignfolderResponseCompoundV3) GetEEzsignfolderStep() FieldEEzsignfolderStep`

GetEEzsignfolderStep returns the EEzsignfolderStep field if non-nil, zero value otherwise.

### GetEEzsignfolderStepOk

`func (o *EzsignfolderResponseCompoundV3) GetEEzsignfolderStepOk() (*FieldEEzsignfolderStep, bool)`

GetEEzsignfolderStepOk returns a tuple with the EEzsignfolderStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignfolderStep

`func (o *EzsignfolderResponseCompoundV3) SetEEzsignfolderStep(v FieldEEzsignfolderStep)`

SetEEzsignfolderStep sets EEzsignfolderStep field to given value.

### HasEEzsignfolderStep

`func (o *EzsignfolderResponseCompoundV3) HasEEzsignfolderStep() bool`

HasEEzsignfolderStep returns a boolean if a field has been set.

### GetDtEzsignfolderClose

`func (o *EzsignfolderResponseCompoundV3) GetDtEzsignfolderClose() string`

GetDtEzsignfolderClose returns the DtEzsignfolderClose field if non-nil, zero value otherwise.

### GetDtEzsignfolderCloseOk

`func (o *EzsignfolderResponseCompoundV3) GetDtEzsignfolderCloseOk() (*string, bool)`

GetDtEzsignfolderCloseOk returns a tuple with the DtEzsignfolderClose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtEzsignfolderClose

`func (o *EzsignfolderResponseCompoundV3) SetDtEzsignfolderClose(v string)`

SetDtEzsignfolderClose sets DtEzsignfolderClose field to given value.

### HasDtEzsignfolderClose

`func (o *EzsignfolderResponseCompoundV3) HasDtEzsignfolderClose() bool`

HasDtEzsignfolderClose returns a boolean if a field has been set.

### GetDtEzsignfolderArchive

`func (o *EzsignfolderResponseCompoundV3) GetDtEzsignfolderArchive() string`

GetDtEzsignfolderArchive returns the DtEzsignfolderArchive field if non-nil, zero value otherwise.

### GetDtEzsignfolderArchiveOk

`func (o *EzsignfolderResponseCompoundV3) GetDtEzsignfolderArchiveOk() (*string, bool)`

GetDtEzsignfolderArchiveOk returns a tuple with the DtEzsignfolderArchive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtEzsignfolderArchive

`func (o *EzsignfolderResponseCompoundV3) SetDtEzsignfolderArchive(v string)`

SetDtEzsignfolderArchive sets DtEzsignfolderArchive field to given value.

### HasDtEzsignfolderArchive

`func (o *EzsignfolderResponseCompoundV3) HasDtEzsignfolderArchive() bool`

HasDtEzsignfolderArchive returns a boolean if a field has been set.

### GetDtEzsignfolderDispose

`func (o *EzsignfolderResponseCompoundV3) GetDtEzsignfolderDispose() string`

GetDtEzsignfolderDispose returns the DtEzsignfolderDispose field if non-nil, zero value otherwise.

### GetDtEzsignfolderDisposeOk

`func (o *EzsignfolderResponseCompoundV3) GetDtEzsignfolderDisposeOk() (*string, bool)`

GetDtEzsignfolderDisposeOk returns a tuple with the DtEzsignfolderDispose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtEzsignfolderDispose

`func (o *EzsignfolderResponseCompoundV3) SetDtEzsignfolderDispose(v string)`

SetDtEzsignfolderDispose sets DtEzsignfolderDispose field to given value.

### HasDtEzsignfolderDispose

`func (o *EzsignfolderResponseCompoundV3) HasDtEzsignfolderDispose() bool`

HasDtEzsignfolderDispose returns a boolean if a field has been set.

### GetTEzsignfolderMessage

`func (o *EzsignfolderResponseCompoundV3) GetTEzsignfolderMessage() string`

GetTEzsignfolderMessage returns the TEzsignfolderMessage field if non-nil, zero value otherwise.

### GetTEzsignfolderMessageOk

`func (o *EzsignfolderResponseCompoundV3) GetTEzsignfolderMessageOk() (*string, bool)`

GetTEzsignfolderMessageOk returns a tuple with the TEzsignfolderMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTEzsignfolderMessage

`func (o *EzsignfolderResponseCompoundV3) SetTEzsignfolderMessage(v string)`

SetTEzsignfolderMessage sets TEzsignfolderMessage field to given value.

### HasTEzsignfolderMessage

`func (o *EzsignfolderResponseCompoundV3) HasTEzsignfolderMessage() bool`

HasTEzsignfolderMessage returns a boolean if a field has been set.

### GetObjAudit

`func (o *EzsignfolderResponseCompoundV3) GetObjAudit() CommonAudit`

GetObjAudit returns the ObjAudit field if non-nil, zero value otherwise.

### GetObjAuditOk

`func (o *EzsignfolderResponseCompoundV3) GetObjAuditOk() (*CommonAudit, bool)`

GetObjAuditOk returns a tuple with the ObjAudit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjAudit

`func (o *EzsignfolderResponseCompoundV3) SetObjAudit(v CommonAudit)`

SetObjAudit sets ObjAudit field to given value.

### HasObjAudit

`func (o *EzsignfolderResponseCompoundV3) HasObjAudit() bool`

HasObjAudit returns a boolean if a field has been set.

### GetSEzsignfolderExternalid

`func (o *EzsignfolderResponseCompoundV3) GetSEzsignfolderExternalid() string`

GetSEzsignfolderExternalid returns the SEzsignfolderExternalid field if non-nil, zero value otherwise.

### GetSEzsignfolderExternalidOk

`func (o *EzsignfolderResponseCompoundV3) GetSEzsignfolderExternalidOk() (*string, bool)`

GetSEzsignfolderExternalidOk returns a tuple with the SEzsignfolderExternalid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignfolderExternalid

`func (o *EzsignfolderResponseCompoundV3) SetSEzsignfolderExternalid(v string)`

SetSEzsignfolderExternalid sets SEzsignfolderExternalid field to given value.

### HasSEzsignfolderExternalid

`func (o *EzsignfolderResponseCompoundV3) HasSEzsignfolderExternalid() bool`

HasSEzsignfolderExternalid returns a boolean if a field has been set.

### GetEEzsignfolderAccess

`func (o *EzsignfolderResponseCompoundV3) GetEEzsignfolderAccess() ComputedEEzsignfolderAccess`

GetEEzsignfolderAccess returns the EEzsignfolderAccess field if non-nil, zero value otherwise.

### GetEEzsignfolderAccessOk

`func (o *EzsignfolderResponseCompoundV3) GetEEzsignfolderAccessOk() (*ComputedEEzsignfolderAccess, bool)`

GetEEzsignfolderAccessOk returns a tuple with the EEzsignfolderAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignfolderAccess

`func (o *EzsignfolderResponseCompoundV3) SetEEzsignfolderAccess(v ComputedEEzsignfolderAccess)`

SetEEzsignfolderAccess sets EEzsignfolderAccess field to given value.

### HasEEzsignfolderAccess

`func (o *EzsignfolderResponseCompoundV3) HasEEzsignfolderAccess() bool`

HasEEzsignfolderAccess returns a boolean if a field has been set.

### GetObjTimezone

`func (o *EzsignfolderResponseCompoundV3) GetObjTimezone() CustomTimezoneWithCodeResponse`

GetObjTimezone returns the ObjTimezone field if non-nil, zero value otherwise.

### GetObjTimezoneOk

`func (o *EzsignfolderResponseCompoundV3) GetObjTimezoneOk() (*CustomTimezoneWithCodeResponse, bool)`

GetObjTimezoneOk returns a tuple with the ObjTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjTimezone

`func (o *EzsignfolderResponseCompoundV3) SetObjTimezone(v CustomTimezoneWithCodeResponse)`

SetObjTimezone sets ObjTimezone field to given value.

### HasObjTimezone

`func (o *EzsignfolderResponseCompoundV3) HasObjTimezone() bool`

HasObjTimezone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


