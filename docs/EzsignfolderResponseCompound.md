# EzsignfolderResponseCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FkiEzsignfoldertypeID** | **int32** | The unique ID of the Ezsignfoldertype.    This value can be queried by the API and is also visible in the admin interface.    There are two types of Ezsignfoldertype. **User** and **Shared**. **User** can only be seen by the user who created the folder or its assistants. Access to **Shared** folders are configurable for access and email delivery. You should typically choose a **Shared** type here. | 
**FkiEzsigntsarequirementID** | **int32** | The unique ID of the Ezsigntsarequirement.  Determine if a Time Stamping Authority should add a timestamp on each of the signature. Valid values:  |Value|Description| |-|-| |1|No. TSA Timestamping will requested. This will make all signatures a lot faster since no round-trip to the TSA server will be required. Timestamping will be made using eZsign server&#39;s time.| |2|Best effort. Timestamping from a Time Stamping Authority will be requested but is not mandatory. In the very improbable case it cannot be completed, the timestamping will be made using eZsign server&#39;s time. **Additional fee applies**| |3|Mandatory. Timestamping from a Time Stamping Authority will be requested and is mandatory. In the very improbable case it cannot be completed, the signature will fail and the user will be asked to retry. **Additional fee applies**| | 
**SEzsignfolderDescription** | **string** | The description of the Ezsign Folder | 
**TEzsignfolderNote** | **string** | Somes extra notes about the eZsign Folder | 
**EEzsignfolderSendreminderfrequency** | [**FieldEEzsignfolderSendreminderfrequency**](FieldEEzsignfolderSendreminderfrequency.md) |  | 
**PkiEzsignfolderID** | **int32** | The unique ID of the Ezsignfolder | 
**DtEzsignfolderSentdate** | **string** | The date and time at which the Ezsign folder was sent the last time. | 
**EEzsignfolderStep** | [**FieldEEzsignfolderStep**](FieldEEzsignfolderStep.md) |  | 
**DtEzsignfolderClose** | **string** | The date and time at which the folder was closed. Either by applying the last signature or by completing it prematurely. | 
**ObjAudit** | [**CommonAudit**](CommonAudit.md) |  | 

## Methods

### NewEzsignfolderResponseCompound

`func NewEzsignfolderResponseCompound(fkiEzsignfoldertypeID int32, fkiEzsigntsarequirementID int32, sEzsignfolderDescription string, tEzsignfolderNote string, eEzsignfolderSendreminderfrequency FieldEEzsignfolderSendreminderfrequency, pkiEzsignfolderID int32, dtEzsignfolderSentdate string, eEzsignfolderStep FieldEEzsignfolderStep, dtEzsignfolderClose string, objAudit CommonAudit, ) *EzsignfolderResponseCompound`

NewEzsignfolderResponseCompound instantiates a new EzsignfolderResponseCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsignfolderResponseCompoundWithDefaults

`func NewEzsignfolderResponseCompoundWithDefaults() *EzsignfolderResponseCompound`

NewEzsignfolderResponseCompoundWithDefaults instantiates a new EzsignfolderResponseCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFkiEzsignfoldertypeID

`func (o *EzsignfolderResponseCompound) GetFkiEzsignfoldertypeID() int32`

GetFkiEzsignfoldertypeID returns the FkiEzsignfoldertypeID field if non-nil, zero value otherwise.

### GetFkiEzsignfoldertypeIDOk

`func (o *EzsignfolderResponseCompound) GetFkiEzsignfoldertypeIDOk() (*int32, bool)`

GetFkiEzsignfoldertypeIDOk returns a tuple with the FkiEzsignfoldertypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignfoldertypeID

`func (o *EzsignfolderResponseCompound) SetFkiEzsignfoldertypeID(v int32)`

SetFkiEzsignfoldertypeID sets FkiEzsignfoldertypeID field to given value.


### GetFkiEzsigntsarequirementID

`func (o *EzsignfolderResponseCompound) GetFkiEzsigntsarequirementID() int32`

GetFkiEzsigntsarequirementID returns the FkiEzsigntsarequirementID field if non-nil, zero value otherwise.

### GetFkiEzsigntsarequirementIDOk

`func (o *EzsignfolderResponseCompound) GetFkiEzsigntsarequirementIDOk() (*int32, bool)`

GetFkiEzsigntsarequirementIDOk returns a tuple with the FkiEzsigntsarequirementID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigntsarequirementID

`func (o *EzsignfolderResponseCompound) SetFkiEzsigntsarequirementID(v int32)`

SetFkiEzsigntsarequirementID sets FkiEzsigntsarequirementID field to given value.


### GetSEzsignfolderDescription

`func (o *EzsignfolderResponseCompound) GetSEzsignfolderDescription() string`

GetSEzsignfolderDescription returns the SEzsignfolderDescription field if non-nil, zero value otherwise.

### GetSEzsignfolderDescriptionOk

`func (o *EzsignfolderResponseCompound) GetSEzsignfolderDescriptionOk() (*string, bool)`

GetSEzsignfolderDescriptionOk returns a tuple with the SEzsignfolderDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignfolderDescription

`func (o *EzsignfolderResponseCompound) SetSEzsignfolderDescription(v string)`

SetSEzsignfolderDescription sets SEzsignfolderDescription field to given value.


### GetTEzsignfolderNote

`func (o *EzsignfolderResponseCompound) GetTEzsignfolderNote() string`

GetTEzsignfolderNote returns the TEzsignfolderNote field if non-nil, zero value otherwise.

### GetTEzsignfolderNoteOk

`func (o *EzsignfolderResponseCompound) GetTEzsignfolderNoteOk() (*string, bool)`

GetTEzsignfolderNoteOk returns a tuple with the TEzsignfolderNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTEzsignfolderNote

`func (o *EzsignfolderResponseCompound) SetTEzsignfolderNote(v string)`

SetTEzsignfolderNote sets TEzsignfolderNote field to given value.


### GetEEzsignfolderSendreminderfrequency

`func (o *EzsignfolderResponseCompound) GetEEzsignfolderSendreminderfrequency() FieldEEzsignfolderSendreminderfrequency`

GetEEzsignfolderSendreminderfrequency returns the EEzsignfolderSendreminderfrequency field if non-nil, zero value otherwise.

### GetEEzsignfolderSendreminderfrequencyOk

`func (o *EzsignfolderResponseCompound) GetEEzsignfolderSendreminderfrequencyOk() (*FieldEEzsignfolderSendreminderfrequency, bool)`

GetEEzsignfolderSendreminderfrequencyOk returns a tuple with the EEzsignfolderSendreminderfrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignfolderSendreminderfrequency

`func (o *EzsignfolderResponseCompound) SetEEzsignfolderSendreminderfrequency(v FieldEEzsignfolderSendreminderfrequency)`

SetEEzsignfolderSendreminderfrequency sets EEzsignfolderSendreminderfrequency field to given value.


### GetPkiEzsignfolderID

`func (o *EzsignfolderResponseCompound) GetPkiEzsignfolderID() int32`

GetPkiEzsignfolderID returns the PkiEzsignfolderID field if non-nil, zero value otherwise.

### GetPkiEzsignfolderIDOk

`func (o *EzsignfolderResponseCompound) GetPkiEzsignfolderIDOk() (*int32, bool)`

GetPkiEzsignfolderIDOk returns a tuple with the PkiEzsignfolderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsignfolderID

`func (o *EzsignfolderResponseCompound) SetPkiEzsignfolderID(v int32)`

SetPkiEzsignfolderID sets PkiEzsignfolderID field to given value.


### GetDtEzsignfolderSentdate

`func (o *EzsignfolderResponseCompound) GetDtEzsignfolderSentdate() string`

GetDtEzsignfolderSentdate returns the DtEzsignfolderSentdate field if non-nil, zero value otherwise.

### GetDtEzsignfolderSentdateOk

`func (o *EzsignfolderResponseCompound) GetDtEzsignfolderSentdateOk() (*string, bool)`

GetDtEzsignfolderSentdateOk returns a tuple with the DtEzsignfolderSentdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtEzsignfolderSentdate

`func (o *EzsignfolderResponseCompound) SetDtEzsignfolderSentdate(v string)`

SetDtEzsignfolderSentdate sets DtEzsignfolderSentdate field to given value.


### GetEEzsignfolderStep

`func (o *EzsignfolderResponseCompound) GetEEzsignfolderStep() FieldEEzsignfolderStep`

GetEEzsignfolderStep returns the EEzsignfolderStep field if non-nil, zero value otherwise.

### GetEEzsignfolderStepOk

`func (o *EzsignfolderResponseCompound) GetEEzsignfolderStepOk() (*FieldEEzsignfolderStep, bool)`

GetEEzsignfolderStepOk returns a tuple with the EEzsignfolderStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignfolderStep

`func (o *EzsignfolderResponseCompound) SetEEzsignfolderStep(v FieldEEzsignfolderStep)`

SetEEzsignfolderStep sets EEzsignfolderStep field to given value.


### GetDtEzsignfolderClose

`func (o *EzsignfolderResponseCompound) GetDtEzsignfolderClose() string`

GetDtEzsignfolderClose returns the DtEzsignfolderClose field if non-nil, zero value otherwise.

### GetDtEzsignfolderCloseOk

`func (o *EzsignfolderResponseCompound) GetDtEzsignfolderCloseOk() (*string, bool)`

GetDtEzsignfolderCloseOk returns a tuple with the DtEzsignfolderClose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtEzsignfolderClose

`func (o *EzsignfolderResponseCompound) SetDtEzsignfolderClose(v string)`

SetDtEzsignfolderClose sets DtEzsignfolderClose field to given value.


### GetObjAudit

`func (o *EzsignfolderResponseCompound) GetObjAudit() CommonAudit`

GetObjAudit returns the ObjAudit field if non-nil, zero value otherwise.

### GetObjAuditOk

`func (o *EzsignfolderResponseCompound) GetObjAuditOk() (*CommonAudit, bool)`

GetObjAuditOk returns a tuple with the ObjAudit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjAudit

`func (o *EzsignfolderResponseCompound) SetObjAudit(v CommonAudit)`

SetObjAudit sets ObjAudit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


