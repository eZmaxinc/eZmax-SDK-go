# EzsignfolderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsignfolderID** | Pointer to **int32** | The unique ID of the Ezsignfolder | [optional] 
**FkiEzsignfoldertypeID** | **int32** | The unique ID of the Ezsignfoldertype. | 
**FkiTimezoneID** | Pointer to **int32** | The unique ID of the Timezone | [optional] 
**FkiEzsigntsarequirementID** | Pointer to **int32** | The unique ID of the Ezsigntsarequirement.  Determine if a Time Stamping Authority should add a timestamp on each of the signature. Valid values:  |Value|Description| |-|-| |1|No. TSA Timestamping will requested. This will make all signatures a lot faster since no round-trip to the TSA server will be required. Timestamping will be made using eZsign server&#39;s time.| |2|Best effort. Timestamping from a Time Stamping Authority will be requested but is not mandatory. In the very improbable case it cannot be completed, the timestamping will be made using eZsign server&#39;s time. **Additional fee applies**| |3|Mandatory. Timestamping from a Time Stamping Authority will be requested and is mandatory. In the very improbable case it cannot be completed, the signature will fail and the user will be asked to retry. **Additional fee applies**| | [optional] 
**SEzsignfolderDescription** | **string** | The description of the Ezsignfolder | 
**TEzsignfolderNote** | Pointer to **string** | Note about the Ezsignfolder | [optional] 
**EEzsignfolderSendreminderfrequency** | [**FieldEEzsignfolderSendreminderfrequency**](FieldEEzsignfolderSendreminderfrequency.md) |  | 
**SEzsignfolderExternalid** | Pointer to **string** | This field can be used to store an External ID from the client&#39;s system.  Anything can be stored in this field, it will never be evaluated by the eZmax system and will be returned AS-IS.  To store multiple values, consider using a JSON formatted structure, a URL encoded string, a CSV or any other custom format.  | [optional] 

## Methods

### NewEzsignfolderRequest

`func NewEzsignfolderRequest(fkiEzsignfoldertypeID int32, sEzsignfolderDescription string, eEzsignfolderSendreminderfrequency FieldEEzsignfolderSendreminderfrequency, ) *EzsignfolderRequest`

NewEzsignfolderRequest instantiates a new EzsignfolderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsignfolderRequestWithDefaults

`func NewEzsignfolderRequestWithDefaults() *EzsignfolderRequest`

NewEzsignfolderRequestWithDefaults instantiates a new EzsignfolderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsignfolderID

`func (o *EzsignfolderRequest) GetPkiEzsignfolderID() int32`

GetPkiEzsignfolderID returns the PkiEzsignfolderID field if non-nil, zero value otherwise.

### GetPkiEzsignfolderIDOk

`func (o *EzsignfolderRequest) GetPkiEzsignfolderIDOk() (*int32, bool)`

GetPkiEzsignfolderIDOk returns a tuple with the PkiEzsignfolderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsignfolderID

`func (o *EzsignfolderRequest) SetPkiEzsignfolderID(v int32)`

SetPkiEzsignfolderID sets PkiEzsignfolderID field to given value.

### HasPkiEzsignfolderID

`func (o *EzsignfolderRequest) HasPkiEzsignfolderID() bool`

HasPkiEzsignfolderID returns a boolean if a field has been set.

### GetFkiEzsignfoldertypeID

`func (o *EzsignfolderRequest) GetFkiEzsignfoldertypeID() int32`

GetFkiEzsignfoldertypeID returns the FkiEzsignfoldertypeID field if non-nil, zero value otherwise.

### GetFkiEzsignfoldertypeIDOk

`func (o *EzsignfolderRequest) GetFkiEzsignfoldertypeIDOk() (*int32, bool)`

GetFkiEzsignfoldertypeIDOk returns a tuple with the FkiEzsignfoldertypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignfoldertypeID

`func (o *EzsignfolderRequest) SetFkiEzsignfoldertypeID(v int32)`

SetFkiEzsignfoldertypeID sets FkiEzsignfoldertypeID field to given value.


### GetFkiTimezoneID

`func (o *EzsignfolderRequest) GetFkiTimezoneID() int32`

GetFkiTimezoneID returns the FkiTimezoneID field if non-nil, zero value otherwise.

### GetFkiTimezoneIDOk

`func (o *EzsignfolderRequest) GetFkiTimezoneIDOk() (*int32, bool)`

GetFkiTimezoneIDOk returns a tuple with the FkiTimezoneID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiTimezoneID

`func (o *EzsignfolderRequest) SetFkiTimezoneID(v int32)`

SetFkiTimezoneID sets FkiTimezoneID field to given value.

### HasFkiTimezoneID

`func (o *EzsignfolderRequest) HasFkiTimezoneID() bool`

HasFkiTimezoneID returns a boolean if a field has been set.

### GetFkiEzsigntsarequirementID

`func (o *EzsignfolderRequest) GetFkiEzsigntsarequirementID() int32`

GetFkiEzsigntsarequirementID returns the FkiEzsigntsarequirementID field if non-nil, zero value otherwise.

### GetFkiEzsigntsarequirementIDOk

`func (o *EzsignfolderRequest) GetFkiEzsigntsarequirementIDOk() (*int32, bool)`

GetFkiEzsigntsarequirementIDOk returns a tuple with the FkiEzsigntsarequirementID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigntsarequirementID

`func (o *EzsignfolderRequest) SetFkiEzsigntsarequirementID(v int32)`

SetFkiEzsigntsarequirementID sets FkiEzsigntsarequirementID field to given value.

### HasFkiEzsigntsarequirementID

`func (o *EzsignfolderRequest) HasFkiEzsigntsarequirementID() bool`

HasFkiEzsigntsarequirementID returns a boolean if a field has been set.

### GetSEzsignfolderDescription

`func (o *EzsignfolderRequest) GetSEzsignfolderDescription() string`

GetSEzsignfolderDescription returns the SEzsignfolderDescription field if non-nil, zero value otherwise.

### GetSEzsignfolderDescriptionOk

`func (o *EzsignfolderRequest) GetSEzsignfolderDescriptionOk() (*string, bool)`

GetSEzsignfolderDescriptionOk returns a tuple with the SEzsignfolderDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignfolderDescription

`func (o *EzsignfolderRequest) SetSEzsignfolderDescription(v string)`

SetSEzsignfolderDescription sets SEzsignfolderDescription field to given value.


### GetTEzsignfolderNote

`func (o *EzsignfolderRequest) GetTEzsignfolderNote() string`

GetTEzsignfolderNote returns the TEzsignfolderNote field if non-nil, zero value otherwise.

### GetTEzsignfolderNoteOk

`func (o *EzsignfolderRequest) GetTEzsignfolderNoteOk() (*string, bool)`

GetTEzsignfolderNoteOk returns a tuple with the TEzsignfolderNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTEzsignfolderNote

`func (o *EzsignfolderRequest) SetTEzsignfolderNote(v string)`

SetTEzsignfolderNote sets TEzsignfolderNote field to given value.

### HasTEzsignfolderNote

`func (o *EzsignfolderRequest) HasTEzsignfolderNote() bool`

HasTEzsignfolderNote returns a boolean if a field has been set.

### GetEEzsignfolderSendreminderfrequency

`func (o *EzsignfolderRequest) GetEEzsignfolderSendreminderfrequency() FieldEEzsignfolderSendreminderfrequency`

GetEEzsignfolderSendreminderfrequency returns the EEzsignfolderSendreminderfrequency field if non-nil, zero value otherwise.

### GetEEzsignfolderSendreminderfrequencyOk

`func (o *EzsignfolderRequest) GetEEzsignfolderSendreminderfrequencyOk() (*FieldEEzsignfolderSendreminderfrequency, bool)`

GetEEzsignfolderSendreminderfrequencyOk returns a tuple with the EEzsignfolderSendreminderfrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignfolderSendreminderfrequency

`func (o *EzsignfolderRequest) SetEEzsignfolderSendreminderfrequency(v FieldEEzsignfolderSendreminderfrequency)`

SetEEzsignfolderSendreminderfrequency sets EEzsignfolderSendreminderfrequency field to given value.


### GetSEzsignfolderExternalid

`func (o *EzsignfolderRequest) GetSEzsignfolderExternalid() string`

GetSEzsignfolderExternalid returns the SEzsignfolderExternalid field if non-nil, zero value otherwise.

### GetSEzsignfolderExternalidOk

`func (o *EzsignfolderRequest) GetSEzsignfolderExternalidOk() (*string, bool)`

GetSEzsignfolderExternalidOk returns a tuple with the SEzsignfolderExternalid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignfolderExternalid

`func (o *EzsignfolderRequest) SetSEzsignfolderExternalid(v string)`

SetSEzsignfolderExternalid sets SEzsignfolderExternalid field to given value.

### HasSEzsignfolderExternalid

`func (o *EzsignfolderRequest) HasSEzsignfolderExternalid() bool`

HasSEzsignfolderExternalid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


