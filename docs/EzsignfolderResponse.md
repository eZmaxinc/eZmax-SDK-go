# EzsignfolderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsignfolderID** | **int32** | The unique ID of the Ezsignfolder | 
**FkiEzsignfoldertypeID** | **int32** | The unique ID of the Ezsignfoldertype. | 
**SEzsignfoldertypeNameX** | **string** | The name of the Ezsignfoldertype in the language of the requester | 
**FkiBillingentityinternalID** | **int32** | The unique ID of the Billingentityinternal. | 
**SBillingentityinternalDescriptionX** | **string** | The description of the Billingentityinternal in the language of the requester | 
**FkiEzsigntsarequirementID** | **int32** | The unique ID of the Ezsigntsarequirement.  Determine if a Time Stamping Authority should add a timestamp on each of the signature. Valid values:  |Value|Description| |-|-| |1|No. TSA Timestamping will requested. This will make all signatures a lot faster since no round-trip to the TSA server will be required. Timestamping will be made using eZsign server&#39;s time.| |2|Best effort. Timestamping from a Time Stamping Authority will be requested but is not mandatory. In the very improbable case it cannot be completed, the timestamping will be made using eZsign server&#39;s time. **Additional fee applies**| |3|Mandatory. Timestamping from a Time Stamping Authority will be requested and is mandatory. In the very improbable case it cannot be completed, the signature will fail and the user will be asked to retry. **Additional fee applies**| | 
**SEzsigntsarequirementDescriptionX** | **string** | The description of the Ezsigntsarequirement in the language of the requester | 
**SEzsignfolderDescription** | **string** | The description of the Ezsignfolder | 
**TEzsignfolderNote** | **string** | Somes extra notes about the eZsign Folder | 
**EEzsignfolderSendreminderfrequency** | [**FieldEEzsignfolderSendreminderfrequency**](FieldEEzsignfolderSendreminderfrequency.md) |  | 
**DtEzsignfolderDuedate** | **string** | The maximum date and time at which the Ezsignfolder can be signed. | 
**DtEzsignfolderSentdate** | **NullableString** | The date and time at which the Ezsign folder was sent the last time. | 
**DtEzsignfolderScheduledarchive** | **string** | The scheduled date and time at which the Ezsignfolder should be archived. | 
**DtEzsignfolderScheduleddestruction** | **string** | The scheduled date and time at which the Ezsignfolder should be Destroyed. | 
**EEzsignfolderStep** | [**FieldEEzsignfolderStep**](FieldEEzsignfolderStep.md) |  | 
**DtEzsignfolderClose** | **string** | The date and time at which the folder was closed. Either by applying the last signature or by completing it prematurely. | 
**ObjAudit** | [**CommonAudit**](CommonAudit.md) |  | 

## Methods

### NewEzsignfolderResponse

`func NewEzsignfolderResponse(pkiEzsignfolderID int32, fkiEzsignfoldertypeID int32, sEzsignfoldertypeNameX string, fkiBillingentityinternalID int32, sBillingentityinternalDescriptionX string, fkiEzsigntsarequirementID int32, sEzsigntsarequirementDescriptionX string, sEzsignfolderDescription string, tEzsignfolderNote string, eEzsignfolderSendreminderfrequency FieldEEzsignfolderSendreminderfrequency, dtEzsignfolderDuedate string, dtEzsignfolderSentdate NullableString, dtEzsignfolderScheduledarchive string, dtEzsignfolderScheduleddestruction string, eEzsignfolderStep FieldEEzsignfolderStep, dtEzsignfolderClose string, objAudit CommonAudit, ) *EzsignfolderResponse`

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


### SetDtEzsignfolderSentdateNil

`func (o *EzsignfolderResponse) SetDtEzsignfolderSentdateNil(b bool)`

 SetDtEzsignfolderSentdateNil sets the value for DtEzsignfolderSentdate to be an explicit nil

### UnsetDtEzsignfolderSentdate
`func (o *EzsignfolderResponse) UnsetDtEzsignfolderSentdate()`

UnsetDtEzsignfolderSentdate ensures that no value is present for DtEzsignfolderSentdate, not even an explicit nil
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


### GetDtEzsignfolderScheduleddestruction

`func (o *EzsignfolderResponse) GetDtEzsignfolderScheduleddestruction() string`

GetDtEzsignfolderScheduleddestruction returns the DtEzsignfolderScheduleddestruction field if non-nil, zero value otherwise.

### GetDtEzsignfolderScheduleddestructionOk

`func (o *EzsignfolderResponse) GetDtEzsignfolderScheduleddestructionOk() (*string, bool)`

GetDtEzsignfolderScheduleddestructionOk returns a tuple with the DtEzsignfolderScheduleddestruction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtEzsignfolderScheduleddestruction

`func (o *EzsignfolderResponse) SetDtEzsignfolderScheduleddestruction(v string)`

SetDtEzsignfolderScheduleddestruction sets DtEzsignfolderScheduleddestruction field to given value.


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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


