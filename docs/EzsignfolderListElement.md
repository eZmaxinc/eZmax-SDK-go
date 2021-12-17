# EzsignfolderListElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsignfolderID** | **int32** | The unique ID of the Ezsignfolder | 
**FkiEzsignfoldertypeID** | **int32** | The unique ID of the Ezsignfoldertype. | 
**EEzsignfoldertypePrivacylevel** | [**FieldEEzsignfoldertypePrivacylevel**](FieldEEzsignfoldertypePrivacylevel.md) |  | 
**SEzsignfoldertypeNameX** | **string** | The name of the Ezsignfoldertype in the language of the requester | 
**SEzsignfolderDescription** | **string** | The description of the Ezsignfolder | 
**EEzsignfolderStep** | [**FieldEEzsignfolderStep**](FieldEEzsignfolderStep.md) |  | 
**DtCreatedDate** | **string** | The date and time at which the object was created | 
**DtEzsignfolderSentdate** | **NullableString** | The date and time at which the Ezsign folder was sent the last time. | 
**DtDueDate** | **NullableString** | Represent a Date Time. The timezone is the one configured in the User&#39;s profile. | 
**IEzsigndocument** | **int32** | The total number of Ezsigndocument in the folder | 
**IEzsigndocumentEdm** | **int32** | The total number of Ezsigndocument in the folder that were saved in the edm system | 
**IEzsignsignature** | **int32** | The total number of signature blocks in all Ezsigndocuments in the folder | 
**IEzsignsignatureSigned** | **int32** | The total number of already signed signature blocks in all Ezsigndocuments in the folder | 

## Methods

### NewEzsignfolderListElement

`func NewEzsignfolderListElement(pkiEzsignfolderID int32, fkiEzsignfoldertypeID int32, eEzsignfoldertypePrivacylevel FieldEEzsignfoldertypePrivacylevel, sEzsignfoldertypeNameX string, sEzsignfolderDescription string, eEzsignfolderStep FieldEEzsignfolderStep, dtCreatedDate string, dtEzsignfolderSentdate NullableString, dtDueDate NullableString, iEzsigndocument int32, iEzsigndocumentEdm int32, iEzsignsignature int32, iEzsignsignatureSigned int32, ) *EzsignfolderListElement`

NewEzsignfolderListElement instantiates a new EzsignfolderListElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsignfolderListElementWithDefaults

`func NewEzsignfolderListElementWithDefaults() *EzsignfolderListElement`

NewEzsignfolderListElementWithDefaults instantiates a new EzsignfolderListElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsignfolderID

`func (o *EzsignfolderListElement) GetPkiEzsignfolderID() int32`

GetPkiEzsignfolderID returns the PkiEzsignfolderID field if non-nil, zero value otherwise.

### GetPkiEzsignfolderIDOk

`func (o *EzsignfolderListElement) GetPkiEzsignfolderIDOk() (*int32, bool)`

GetPkiEzsignfolderIDOk returns a tuple with the PkiEzsignfolderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsignfolderID

`func (o *EzsignfolderListElement) SetPkiEzsignfolderID(v int32)`

SetPkiEzsignfolderID sets PkiEzsignfolderID field to given value.


### GetFkiEzsignfoldertypeID

`func (o *EzsignfolderListElement) GetFkiEzsignfoldertypeID() int32`

GetFkiEzsignfoldertypeID returns the FkiEzsignfoldertypeID field if non-nil, zero value otherwise.

### GetFkiEzsignfoldertypeIDOk

`func (o *EzsignfolderListElement) GetFkiEzsignfoldertypeIDOk() (*int32, bool)`

GetFkiEzsignfoldertypeIDOk returns a tuple with the FkiEzsignfoldertypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignfoldertypeID

`func (o *EzsignfolderListElement) SetFkiEzsignfoldertypeID(v int32)`

SetFkiEzsignfoldertypeID sets FkiEzsignfoldertypeID field to given value.


### GetEEzsignfoldertypePrivacylevel

`func (o *EzsignfolderListElement) GetEEzsignfoldertypePrivacylevel() FieldEEzsignfoldertypePrivacylevel`

GetEEzsignfoldertypePrivacylevel returns the EEzsignfoldertypePrivacylevel field if non-nil, zero value otherwise.

### GetEEzsignfoldertypePrivacylevelOk

`func (o *EzsignfolderListElement) GetEEzsignfoldertypePrivacylevelOk() (*FieldEEzsignfoldertypePrivacylevel, bool)`

GetEEzsignfoldertypePrivacylevelOk returns a tuple with the EEzsignfoldertypePrivacylevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignfoldertypePrivacylevel

`func (o *EzsignfolderListElement) SetEEzsignfoldertypePrivacylevel(v FieldEEzsignfoldertypePrivacylevel)`

SetEEzsignfoldertypePrivacylevel sets EEzsignfoldertypePrivacylevel field to given value.


### GetSEzsignfoldertypeNameX

`func (o *EzsignfolderListElement) GetSEzsignfoldertypeNameX() string`

GetSEzsignfoldertypeNameX returns the SEzsignfoldertypeNameX field if non-nil, zero value otherwise.

### GetSEzsignfoldertypeNameXOk

`func (o *EzsignfolderListElement) GetSEzsignfoldertypeNameXOk() (*string, bool)`

GetSEzsignfoldertypeNameXOk returns a tuple with the SEzsignfoldertypeNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignfoldertypeNameX

`func (o *EzsignfolderListElement) SetSEzsignfoldertypeNameX(v string)`

SetSEzsignfoldertypeNameX sets SEzsignfoldertypeNameX field to given value.


### GetSEzsignfolderDescription

`func (o *EzsignfolderListElement) GetSEzsignfolderDescription() string`

GetSEzsignfolderDescription returns the SEzsignfolderDescription field if non-nil, zero value otherwise.

### GetSEzsignfolderDescriptionOk

`func (o *EzsignfolderListElement) GetSEzsignfolderDescriptionOk() (*string, bool)`

GetSEzsignfolderDescriptionOk returns a tuple with the SEzsignfolderDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignfolderDescription

`func (o *EzsignfolderListElement) SetSEzsignfolderDescription(v string)`

SetSEzsignfolderDescription sets SEzsignfolderDescription field to given value.


### GetEEzsignfolderStep

`func (o *EzsignfolderListElement) GetEEzsignfolderStep() FieldEEzsignfolderStep`

GetEEzsignfolderStep returns the EEzsignfolderStep field if non-nil, zero value otherwise.

### GetEEzsignfolderStepOk

`func (o *EzsignfolderListElement) GetEEzsignfolderStepOk() (*FieldEEzsignfolderStep, bool)`

GetEEzsignfolderStepOk returns a tuple with the EEzsignfolderStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignfolderStep

`func (o *EzsignfolderListElement) SetEEzsignfolderStep(v FieldEEzsignfolderStep)`

SetEEzsignfolderStep sets EEzsignfolderStep field to given value.


### GetDtCreatedDate

`func (o *EzsignfolderListElement) GetDtCreatedDate() string`

GetDtCreatedDate returns the DtCreatedDate field if non-nil, zero value otherwise.

### GetDtCreatedDateOk

`func (o *EzsignfolderListElement) GetDtCreatedDateOk() (*string, bool)`

GetDtCreatedDateOk returns a tuple with the DtCreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtCreatedDate

`func (o *EzsignfolderListElement) SetDtCreatedDate(v string)`

SetDtCreatedDate sets DtCreatedDate field to given value.


### GetDtEzsignfolderSentdate

`func (o *EzsignfolderListElement) GetDtEzsignfolderSentdate() string`

GetDtEzsignfolderSentdate returns the DtEzsignfolderSentdate field if non-nil, zero value otherwise.

### GetDtEzsignfolderSentdateOk

`func (o *EzsignfolderListElement) GetDtEzsignfolderSentdateOk() (*string, bool)`

GetDtEzsignfolderSentdateOk returns a tuple with the DtEzsignfolderSentdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtEzsignfolderSentdate

`func (o *EzsignfolderListElement) SetDtEzsignfolderSentdate(v string)`

SetDtEzsignfolderSentdate sets DtEzsignfolderSentdate field to given value.


### SetDtEzsignfolderSentdateNil

`func (o *EzsignfolderListElement) SetDtEzsignfolderSentdateNil(b bool)`

 SetDtEzsignfolderSentdateNil sets the value for DtEzsignfolderSentdate to be an explicit nil

### UnsetDtEzsignfolderSentdate
`func (o *EzsignfolderListElement) UnsetDtEzsignfolderSentdate()`

UnsetDtEzsignfolderSentdate ensures that no value is present for DtEzsignfolderSentdate, not even an explicit nil
### GetDtDueDate

`func (o *EzsignfolderListElement) GetDtDueDate() string`

GetDtDueDate returns the DtDueDate field if non-nil, zero value otherwise.

### GetDtDueDateOk

`func (o *EzsignfolderListElement) GetDtDueDateOk() (*string, bool)`

GetDtDueDateOk returns a tuple with the DtDueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtDueDate

`func (o *EzsignfolderListElement) SetDtDueDate(v string)`

SetDtDueDate sets DtDueDate field to given value.


### SetDtDueDateNil

`func (o *EzsignfolderListElement) SetDtDueDateNil(b bool)`

 SetDtDueDateNil sets the value for DtDueDate to be an explicit nil

### UnsetDtDueDate
`func (o *EzsignfolderListElement) UnsetDtDueDate()`

UnsetDtDueDate ensures that no value is present for DtDueDate, not even an explicit nil
### GetIEzsigndocument

`func (o *EzsignfolderListElement) GetIEzsigndocument() int32`

GetIEzsigndocument returns the IEzsigndocument field if non-nil, zero value otherwise.

### GetIEzsigndocumentOk

`func (o *EzsignfolderListElement) GetIEzsigndocumentOk() (*int32, bool)`

GetIEzsigndocumentOk returns a tuple with the IEzsigndocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigndocument

`func (o *EzsignfolderListElement) SetIEzsigndocument(v int32)`

SetIEzsigndocument sets IEzsigndocument field to given value.


### GetIEzsigndocumentEdm

`func (o *EzsignfolderListElement) GetIEzsigndocumentEdm() int32`

GetIEzsigndocumentEdm returns the IEzsigndocumentEdm field if non-nil, zero value otherwise.

### GetIEzsigndocumentEdmOk

`func (o *EzsignfolderListElement) GetIEzsigndocumentEdmOk() (*int32, bool)`

GetIEzsigndocumentEdmOk returns a tuple with the IEzsigndocumentEdm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigndocumentEdm

`func (o *EzsignfolderListElement) SetIEzsigndocumentEdm(v int32)`

SetIEzsigndocumentEdm sets IEzsigndocumentEdm field to given value.


### GetIEzsignsignature

`func (o *EzsignfolderListElement) GetIEzsignsignature() int32`

GetIEzsignsignature returns the IEzsignsignature field if non-nil, zero value otherwise.

### GetIEzsignsignatureOk

`func (o *EzsignfolderListElement) GetIEzsignsignatureOk() (*int32, bool)`

GetIEzsignsignatureOk returns a tuple with the IEzsignsignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignsignature

`func (o *EzsignfolderListElement) SetIEzsignsignature(v int32)`

SetIEzsignsignature sets IEzsignsignature field to given value.


### GetIEzsignsignatureSigned

`func (o *EzsignfolderListElement) GetIEzsignsignatureSigned() int32`

GetIEzsignsignatureSigned returns the IEzsignsignatureSigned field if non-nil, zero value otherwise.

### GetIEzsignsignatureSignedOk

`func (o *EzsignfolderListElement) GetIEzsignsignatureSignedOk() (*int32, bool)`

GetIEzsignsignatureSignedOk returns a tuple with the IEzsignsignatureSigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignsignatureSigned

`func (o *EzsignfolderListElement) SetIEzsignsignatureSigned(v int32)`

SetIEzsignsignatureSigned sets IEzsignsignatureSigned field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


