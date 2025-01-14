# EzsignfolderRequestCompoundV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsignfolderID** | Pointer to **int32** | The unique ID of the Ezsignfolder | [optional] 
**FkiEzsignfoldertypeID** | **int32** | The unique ID of the Ezsignfoldertype. | 
**FkiTimezoneID** | Pointer to **int32** | The unique ID of the Timezone | [optional] 
**FkiEzsigntsarequirementID** | Pointer to **int32** | The unique ID of the Ezsigntsarequirement.  Determine if a Time Stamping Authority should add a timestamp on each of the signature. Valid values:  |Value|Description| |-|-| |1|No. TSA Timestamping will requested. This will make all signatures a lot faster since no round-trip to the TSA server will be required. Timestamping will be made using eZsign server&#39;s time.| |2|Best effort. Timestamping from a Time Stamping Authority will be requested but is not mandatory. In the very improbable case it cannot be completed, the timestamping will be made using eZsign server&#39;s time. **Additional fee applies**| |3|Mandatory. Timestamping from a Time Stamping Authority will be requested and is mandatory. In the very improbable case it cannot be completed, the signature will fail and the user will be asked to retry. **Additional fee applies**| | [optional] 
**EEzsignfolderDocumentdependency** | Pointer to [**FieldEEzsignfolderDocumentdependency**](FieldEEzsignfolderDocumentdependency.md) |  | [optional] 
**SEzsignfolderDescription** | **string** | The description of the Ezsignfolder | 
**TEzsignfolderNote** | Pointer to **string** | Note about the Ezsignfolder | [optional] 
**TEzsignfolderMessage** | Pointer to **string** | A custom text message that will be added to the email sent. | [optional] 
**IEzsignfolderSendreminderfirstdays** | **int32** | The number of days before the the first reminder sending | 
**IEzsignfolderSendreminderotherdays** | **int32** | The number of days after the first reminder sending | 
**SEzsignfolderExternalid** | Pointer to **string** | This field can be used to store an External ID from the client&#39;s system.  Anything can be stored in this field, it will never be evaluated by the eZmax system and will be returned AS-IS.  To store multiple values, consider using a JSON formatted structure, a URL encoded string, a CSV or any other custom format.  | [optional] 

## Methods

### NewEzsignfolderRequestCompoundV3

`func NewEzsignfolderRequestCompoundV3(fkiEzsignfoldertypeID int32, sEzsignfolderDescription string, iEzsignfolderSendreminderfirstdays int32, iEzsignfolderSendreminderotherdays int32, ) *EzsignfolderRequestCompoundV3`

NewEzsignfolderRequestCompoundV3 instantiates a new EzsignfolderRequestCompoundV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsignfolderRequestCompoundV3WithDefaults

`func NewEzsignfolderRequestCompoundV3WithDefaults() *EzsignfolderRequestCompoundV3`

NewEzsignfolderRequestCompoundV3WithDefaults instantiates a new EzsignfolderRequestCompoundV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsignfolderID

`func (o *EzsignfolderRequestCompoundV3) GetPkiEzsignfolderID() int32`

GetPkiEzsignfolderID returns the PkiEzsignfolderID field if non-nil, zero value otherwise.

### GetPkiEzsignfolderIDOk

`func (o *EzsignfolderRequestCompoundV3) GetPkiEzsignfolderIDOk() (*int32, bool)`

GetPkiEzsignfolderIDOk returns a tuple with the PkiEzsignfolderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsignfolderID

`func (o *EzsignfolderRequestCompoundV3) SetPkiEzsignfolderID(v int32)`

SetPkiEzsignfolderID sets PkiEzsignfolderID field to given value.

### HasPkiEzsignfolderID

`func (o *EzsignfolderRequestCompoundV3) HasPkiEzsignfolderID() bool`

HasPkiEzsignfolderID returns a boolean if a field has been set.

### GetFkiEzsignfoldertypeID

`func (o *EzsignfolderRequestCompoundV3) GetFkiEzsignfoldertypeID() int32`

GetFkiEzsignfoldertypeID returns the FkiEzsignfoldertypeID field if non-nil, zero value otherwise.

### GetFkiEzsignfoldertypeIDOk

`func (o *EzsignfolderRequestCompoundV3) GetFkiEzsignfoldertypeIDOk() (*int32, bool)`

GetFkiEzsignfoldertypeIDOk returns a tuple with the FkiEzsignfoldertypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignfoldertypeID

`func (o *EzsignfolderRequestCompoundV3) SetFkiEzsignfoldertypeID(v int32)`

SetFkiEzsignfoldertypeID sets FkiEzsignfoldertypeID field to given value.


### GetFkiTimezoneID

`func (o *EzsignfolderRequestCompoundV3) GetFkiTimezoneID() int32`

GetFkiTimezoneID returns the FkiTimezoneID field if non-nil, zero value otherwise.

### GetFkiTimezoneIDOk

`func (o *EzsignfolderRequestCompoundV3) GetFkiTimezoneIDOk() (*int32, bool)`

GetFkiTimezoneIDOk returns a tuple with the FkiTimezoneID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiTimezoneID

`func (o *EzsignfolderRequestCompoundV3) SetFkiTimezoneID(v int32)`

SetFkiTimezoneID sets FkiTimezoneID field to given value.

### HasFkiTimezoneID

`func (o *EzsignfolderRequestCompoundV3) HasFkiTimezoneID() bool`

HasFkiTimezoneID returns a boolean if a field has been set.

### GetFkiEzsigntsarequirementID

`func (o *EzsignfolderRequestCompoundV3) GetFkiEzsigntsarequirementID() int32`

GetFkiEzsigntsarequirementID returns the FkiEzsigntsarequirementID field if non-nil, zero value otherwise.

### GetFkiEzsigntsarequirementIDOk

`func (o *EzsignfolderRequestCompoundV3) GetFkiEzsigntsarequirementIDOk() (*int32, bool)`

GetFkiEzsigntsarequirementIDOk returns a tuple with the FkiEzsigntsarequirementID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigntsarequirementID

`func (o *EzsignfolderRequestCompoundV3) SetFkiEzsigntsarequirementID(v int32)`

SetFkiEzsigntsarequirementID sets FkiEzsigntsarequirementID field to given value.

### HasFkiEzsigntsarequirementID

`func (o *EzsignfolderRequestCompoundV3) HasFkiEzsigntsarequirementID() bool`

HasFkiEzsigntsarequirementID returns a boolean if a field has been set.

### GetEEzsignfolderDocumentdependency

`func (o *EzsignfolderRequestCompoundV3) GetEEzsignfolderDocumentdependency() FieldEEzsignfolderDocumentdependency`

GetEEzsignfolderDocumentdependency returns the EEzsignfolderDocumentdependency field if non-nil, zero value otherwise.

### GetEEzsignfolderDocumentdependencyOk

`func (o *EzsignfolderRequestCompoundV3) GetEEzsignfolderDocumentdependencyOk() (*FieldEEzsignfolderDocumentdependency, bool)`

GetEEzsignfolderDocumentdependencyOk returns a tuple with the EEzsignfolderDocumentdependency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignfolderDocumentdependency

`func (o *EzsignfolderRequestCompoundV3) SetEEzsignfolderDocumentdependency(v FieldEEzsignfolderDocumentdependency)`

SetEEzsignfolderDocumentdependency sets EEzsignfolderDocumentdependency field to given value.

### HasEEzsignfolderDocumentdependency

`func (o *EzsignfolderRequestCompoundV3) HasEEzsignfolderDocumentdependency() bool`

HasEEzsignfolderDocumentdependency returns a boolean if a field has been set.

### GetSEzsignfolderDescription

`func (o *EzsignfolderRequestCompoundV3) GetSEzsignfolderDescription() string`

GetSEzsignfolderDescription returns the SEzsignfolderDescription field if non-nil, zero value otherwise.

### GetSEzsignfolderDescriptionOk

`func (o *EzsignfolderRequestCompoundV3) GetSEzsignfolderDescriptionOk() (*string, bool)`

GetSEzsignfolderDescriptionOk returns a tuple with the SEzsignfolderDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignfolderDescription

`func (o *EzsignfolderRequestCompoundV3) SetSEzsignfolderDescription(v string)`

SetSEzsignfolderDescription sets SEzsignfolderDescription field to given value.


### GetTEzsignfolderNote

`func (o *EzsignfolderRequestCompoundV3) GetTEzsignfolderNote() string`

GetTEzsignfolderNote returns the TEzsignfolderNote field if non-nil, zero value otherwise.

### GetTEzsignfolderNoteOk

`func (o *EzsignfolderRequestCompoundV3) GetTEzsignfolderNoteOk() (*string, bool)`

GetTEzsignfolderNoteOk returns a tuple with the TEzsignfolderNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTEzsignfolderNote

`func (o *EzsignfolderRequestCompoundV3) SetTEzsignfolderNote(v string)`

SetTEzsignfolderNote sets TEzsignfolderNote field to given value.

### HasTEzsignfolderNote

`func (o *EzsignfolderRequestCompoundV3) HasTEzsignfolderNote() bool`

HasTEzsignfolderNote returns a boolean if a field has been set.

### GetTEzsignfolderMessage

`func (o *EzsignfolderRequestCompoundV3) GetTEzsignfolderMessage() string`

GetTEzsignfolderMessage returns the TEzsignfolderMessage field if non-nil, zero value otherwise.

### GetTEzsignfolderMessageOk

`func (o *EzsignfolderRequestCompoundV3) GetTEzsignfolderMessageOk() (*string, bool)`

GetTEzsignfolderMessageOk returns a tuple with the TEzsignfolderMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTEzsignfolderMessage

`func (o *EzsignfolderRequestCompoundV3) SetTEzsignfolderMessage(v string)`

SetTEzsignfolderMessage sets TEzsignfolderMessage field to given value.

### HasTEzsignfolderMessage

`func (o *EzsignfolderRequestCompoundV3) HasTEzsignfolderMessage() bool`

HasTEzsignfolderMessage returns a boolean if a field has been set.

### GetIEzsignfolderSendreminderfirstdays

`func (o *EzsignfolderRequestCompoundV3) GetIEzsignfolderSendreminderfirstdays() int32`

GetIEzsignfolderSendreminderfirstdays returns the IEzsignfolderSendreminderfirstdays field if non-nil, zero value otherwise.

### GetIEzsignfolderSendreminderfirstdaysOk

`func (o *EzsignfolderRequestCompoundV3) GetIEzsignfolderSendreminderfirstdaysOk() (*int32, bool)`

GetIEzsignfolderSendreminderfirstdaysOk returns a tuple with the IEzsignfolderSendreminderfirstdays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignfolderSendreminderfirstdays

`func (o *EzsignfolderRequestCompoundV3) SetIEzsignfolderSendreminderfirstdays(v int32)`

SetIEzsignfolderSendreminderfirstdays sets IEzsignfolderSendreminderfirstdays field to given value.


### GetIEzsignfolderSendreminderotherdays

`func (o *EzsignfolderRequestCompoundV3) GetIEzsignfolderSendreminderotherdays() int32`

GetIEzsignfolderSendreminderotherdays returns the IEzsignfolderSendreminderotherdays field if non-nil, zero value otherwise.

### GetIEzsignfolderSendreminderotherdaysOk

`func (o *EzsignfolderRequestCompoundV3) GetIEzsignfolderSendreminderotherdaysOk() (*int32, bool)`

GetIEzsignfolderSendreminderotherdaysOk returns a tuple with the IEzsignfolderSendreminderotherdays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignfolderSendreminderotherdays

`func (o *EzsignfolderRequestCompoundV3) SetIEzsignfolderSendreminderotherdays(v int32)`

SetIEzsignfolderSendreminderotherdays sets IEzsignfolderSendreminderotherdays field to given value.


### GetSEzsignfolderExternalid

`func (o *EzsignfolderRequestCompoundV3) GetSEzsignfolderExternalid() string`

GetSEzsignfolderExternalid returns the SEzsignfolderExternalid field if non-nil, zero value otherwise.

### GetSEzsignfolderExternalidOk

`func (o *EzsignfolderRequestCompoundV3) GetSEzsignfolderExternalidOk() (*string, bool)`

GetSEzsignfolderExternalidOk returns a tuple with the SEzsignfolderExternalid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignfolderExternalid

`func (o *EzsignfolderRequestCompoundV3) SetSEzsignfolderExternalid(v string)`

SetSEzsignfolderExternalid sets SEzsignfolderExternalid field to given value.

### HasSEzsignfolderExternalid

`func (o *EzsignfolderRequestCompoundV3) HasSEzsignfolderExternalid() bool`

HasSEzsignfolderExternalid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


