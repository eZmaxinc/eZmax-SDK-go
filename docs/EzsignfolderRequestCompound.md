# EzsignfolderRequestCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AEzsignfoldersignerassociation** | [**[]EzsignfoldersignerassociationRequest**](EzsignfoldersignerassociationRequest.md) | An array of signers that will be invited to sign the Ezsigndocuments | 
**FkiEzsignfoldertypeID** | **int32** | The unique ID of the Ezsignfoldertype. | 
**FkiEzsigntsarequirementID** | **int32** | The unique ID of the Ezsigntsarequirement.  Determine if a Time Stamping Authority should add a timestamp on each of the signature. Valid values:  |Value|Description| |-|-| |1|No. TSA Timestamping will requested. This will make all signatures a lot faster since no round-trip to the TSA server will be required. Timestamping will be made using eZsign server&#39;s time.| |2|Best effort. Timestamping from a Time Stamping Authority will be requested but is not mandatory. In the very improbable case it cannot be completed, the timestamping will be made using eZsign server&#39;s time. **Additional fee applies**| |3|Mandatory. Timestamping from a Time Stamping Authority will be requested and is mandatory. In the very improbable case it cannot be completed, the signature will fail and the user will be asked to retry. **Additional fee applies**| | 
**SEzsignfolderDescription** | **string** | The description of the Ezsign Folder | 
**TEzsignfolderNote** | **string** | Somes extra notes about the eZsign Folder | 
**EEzsignfolderSendreminderfrequency** | [**FieldEEzsignfolderSendreminderfrequency**](FieldEEzsignfolderSendreminderfrequency.md) |  | 

## Methods

### NewEzsignfolderRequestCompound

`func NewEzsignfolderRequestCompound(aEzsignfoldersignerassociation []EzsignfoldersignerassociationRequest, fkiEzsignfoldertypeID int32, fkiEzsigntsarequirementID int32, sEzsignfolderDescription string, tEzsignfolderNote string, eEzsignfolderSendreminderfrequency FieldEEzsignfolderSendreminderfrequency, ) *EzsignfolderRequestCompound`

NewEzsignfolderRequestCompound instantiates a new EzsignfolderRequestCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsignfolderRequestCompoundWithDefaults

`func NewEzsignfolderRequestCompoundWithDefaults() *EzsignfolderRequestCompound`

NewEzsignfolderRequestCompoundWithDefaults instantiates a new EzsignfolderRequestCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAEzsignfoldersignerassociation

`func (o *EzsignfolderRequestCompound) GetAEzsignfoldersignerassociation() []EzsignfoldersignerassociationRequest`

GetAEzsignfoldersignerassociation returns the AEzsignfoldersignerassociation field if non-nil, zero value otherwise.

### GetAEzsignfoldersignerassociationOk

`func (o *EzsignfolderRequestCompound) GetAEzsignfoldersignerassociationOk() (*[]EzsignfoldersignerassociationRequest, bool)`

GetAEzsignfoldersignerassociationOk returns a tuple with the AEzsignfoldersignerassociation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAEzsignfoldersignerassociation

`func (o *EzsignfolderRequestCompound) SetAEzsignfoldersignerassociation(v []EzsignfoldersignerassociationRequest)`

SetAEzsignfoldersignerassociation sets AEzsignfoldersignerassociation field to given value.


### GetFkiEzsignfoldertypeID

`func (o *EzsignfolderRequestCompound) GetFkiEzsignfoldertypeID() int32`

GetFkiEzsignfoldertypeID returns the FkiEzsignfoldertypeID field if non-nil, zero value otherwise.

### GetFkiEzsignfoldertypeIDOk

`func (o *EzsignfolderRequestCompound) GetFkiEzsignfoldertypeIDOk() (*int32, bool)`

GetFkiEzsignfoldertypeIDOk returns a tuple with the FkiEzsignfoldertypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignfoldertypeID

`func (o *EzsignfolderRequestCompound) SetFkiEzsignfoldertypeID(v int32)`

SetFkiEzsignfoldertypeID sets FkiEzsignfoldertypeID field to given value.


### GetFkiEzsigntsarequirementID

`func (o *EzsignfolderRequestCompound) GetFkiEzsigntsarequirementID() int32`

GetFkiEzsigntsarequirementID returns the FkiEzsigntsarequirementID field if non-nil, zero value otherwise.

### GetFkiEzsigntsarequirementIDOk

`func (o *EzsignfolderRequestCompound) GetFkiEzsigntsarequirementIDOk() (*int32, bool)`

GetFkiEzsigntsarequirementIDOk returns a tuple with the FkiEzsigntsarequirementID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigntsarequirementID

`func (o *EzsignfolderRequestCompound) SetFkiEzsigntsarequirementID(v int32)`

SetFkiEzsigntsarequirementID sets FkiEzsigntsarequirementID field to given value.


### GetSEzsignfolderDescription

`func (o *EzsignfolderRequestCompound) GetSEzsignfolderDescription() string`

GetSEzsignfolderDescription returns the SEzsignfolderDescription field if non-nil, zero value otherwise.

### GetSEzsignfolderDescriptionOk

`func (o *EzsignfolderRequestCompound) GetSEzsignfolderDescriptionOk() (*string, bool)`

GetSEzsignfolderDescriptionOk returns a tuple with the SEzsignfolderDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignfolderDescription

`func (o *EzsignfolderRequestCompound) SetSEzsignfolderDescription(v string)`

SetSEzsignfolderDescription sets SEzsignfolderDescription field to given value.


### GetTEzsignfolderNote

`func (o *EzsignfolderRequestCompound) GetTEzsignfolderNote() string`

GetTEzsignfolderNote returns the TEzsignfolderNote field if non-nil, zero value otherwise.

### GetTEzsignfolderNoteOk

`func (o *EzsignfolderRequestCompound) GetTEzsignfolderNoteOk() (*string, bool)`

GetTEzsignfolderNoteOk returns a tuple with the TEzsignfolderNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTEzsignfolderNote

`func (o *EzsignfolderRequestCompound) SetTEzsignfolderNote(v string)`

SetTEzsignfolderNote sets TEzsignfolderNote field to given value.


### GetEEzsignfolderSendreminderfrequency

`func (o *EzsignfolderRequestCompound) GetEEzsignfolderSendreminderfrequency() FieldEEzsignfolderSendreminderfrequency`

GetEEzsignfolderSendreminderfrequency returns the EEzsignfolderSendreminderfrequency field if non-nil, zero value otherwise.

### GetEEzsignfolderSendreminderfrequencyOk

`func (o *EzsignfolderRequestCompound) GetEEzsignfolderSendreminderfrequencyOk() (*FieldEEzsignfolderSendreminderfrequency, bool)`

GetEEzsignfolderSendreminderfrequencyOk returns a tuple with the EEzsignfolderSendreminderfrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignfolderSendreminderfrequency

`func (o *EzsignfolderRequestCompound) SetEEzsignfolderSendreminderfrequency(v FieldEEzsignfolderSendreminderfrequency)`

SetEEzsignfolderSendreminderfrequency sets EEzsignfolderSendreminderfrequency field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


