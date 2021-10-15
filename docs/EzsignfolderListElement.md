# EzsignfolderListElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsignfolderID** | **int32** | The unique ID of the Ezsignfolder | 
**FkiEzsignfoldertypeID** | **int32** | The unique ID of the Ezsignfoldertype. | 
**SEzsignfoldertypeNameX** | **string** | The name of the Ezsignfoldertype in the language of the requester | 
**SEzsignfolderDescription** | **string** | The description of the Ezsign Folder | 
**EEzsignfolderStep** | [**FieldEEzsignfolderStep**](FieldEEzsignfolderStep.md) |  | 
**DtCreatedDate** | **string** | The date and time at which the object was created | 
**DtEzsignfolderSentdate** | [**OneOfstringnull**](oneOf&lt;string,null&gt;.md) |  | 
**DtDueDate** | [**OneOfstringnull**](oneOf&lt;string,null&gt;.md) | The date at which no more signature will be accepted on the folder | 
**ITotalDocument** | **int32** | The total number of Ezsigndocument in the folder | 
**ITotalDocumentEdm** | **int32** | The total number of Ezsigndocument in the folder that were saved in the edm system | 
**ITotalSignature** | **int32** | The total number of signature blocks in all Ezsigndocuments in the folder | 
**ITotalSignatureSigned** | **int32** | The total number of already signed signature blocks in all Ezsigndocuments in the folder | 

## Methods

### NewEzsignfolderListElement

`func NewEzsignfolderListElement(pkiEzsignfolderID int32, fkiEzsignfoldertypeID int32, sEzsignfoldertypeNameX string, sEzsignfolderDescription string, eEzsignfolderStep FieldEEzsignfolderStep, dtCreatedDate string, dtEzsignfolderSentdate OneOfstringnull, dtDueDate OneOfstringnull, iTotalDocument int32, iTotalDocumentEdm int32, iTotalSignature int32, iTotalSignatureSigned int32, ) *EzsignfolderListElement`

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

`func (o *EzsignfolderListElement) GetDtEzsignfolderSentdate() OneOfstringnull`

GetDtEzsignfolderSentdate returns the DtEzsignfolderSentdate field if non-nil, zero value otherwise.

### GetDtEzsignfolderSentdateOk

`func (o *EzsignfolderListElement) GetDtEzsignfolderSentdateOk() (*OneOfstringnull, bool)`

GetDtEzsignfolderSentdateOk returns a tuple with the DtEzsignfolderSentdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtEzsignfolderSentdate

`func (o *EzsignfolderListElement) SetDtEzsignfolderSentdate(v OneOfstringnull)`

SetDtEzsignfolderSentdate sets DtEzsignfolderSentdate field to given value.


### SetDtEzsignfolderSentdateNil

`func (o *EzsignfolderListElement) SetDtEzsignfolderSentdateNil(b bool)`

 SetDtEzsignfolderSentdateNil sets the value for DtEzsignfolderSentdate to be an explicit nil

### UnsetDtEzsignfolderSentdate
`func (o *EzsignfolderListElement) UnsetDtEzsignfolderSentdate()`

UnsetDtEzsignfolderSentdate ensures that no value is present for DtEzsignfolderSentdate, not even an explicit nil
### GetDtDueDate

`func (o *EzsignfolderListElement) GetDtDueDate() OneOfstringnull`

GetDtDueDate returns the DtDueDate field if non-nil, zero value otherwise.

### GetDtDueDateOk

`func (o *EzsignfolderListElement) GetDtDueDateOk() (*OneOfstringnull, bool)`

GetDtDueDateOk returns a tuple with the DtDueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtDueDate

`func (o *EzsignfolderListElement) SetDtDueDate(v OneOfstringnull)`

SetDtDueDate sets DtDueDate field to given value.


### SetDtDueDateNil

`func (o *EzsignfolderListElement) SetDtDueDateNil(b bool)`

 SetDtDueDateNil sets the value for DtDueDate to be an explicit nil

### UnsetDtDueDate
`func (o *EzsignfolderListElement) UnsetDtDueDate()`

UnsetDtDueDate ensures that no value is present for DtDueDate, not even an explicit nil
### GetITotalDocument

`func (o *EzsignfolderListElement) GetITotalDocument() int32`

GetITotalDocument returns the ITotalDocument field if non-nil, zero value otherwise.

### GetITotalDocumentOk

`func (o *EzsignfolderListElement) GetITotalDocumentOk() (*int32, bool)`

GetITotalDocumentOk returns a tuple with the ITotalDocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetITotalDocument

`func (o *EzsignfolderListElement) SetITotalDocument(v int32)`

SetITotalDocument sets ITotalDocument field to given value.


### GetITotalDocumentEdm

`func (o *EzsignfolderListElement) GetITotalDocumentEdm() int32`

GetITotalDocumentEdm returns the ITotalDocumentEdm field if non-nil, zero value otherwise.

### GetITotalDocumentEdmOk

`func (o *EzsignfolderListElement) GetITotalDocumentEdmOk() (*int32, bool)`

GetITotalDocumentEdmOk returns a tuple with the ITotalDocumentEdm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetITotalDocumentEdm

`func (o *EzsignfolderListElement) SetITotalDocumentEdm(v int32)`

SetITotalDocumentEdm sets ITotalDocumentEdm field to given value.


### GetITotalSignature

`func (o *EzsignfolderListElement) GetITotalSignature() int32`

GetITotalSignature returns the ITotalSignature field if non-nil, zero value otherwise.

### GetITotalSignatureOk

`func (o *EzsignfolderListElement) GetITotalSignatureOk() (*int32, bool)`

GetITotalSignatureOk returns a tuple with the ITotalSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetITotalSignature

`func (o *EzsignfolderListElement) SetITotalSignature(v int32)`

SetITotalSignature sets ITotalSignature field to given value.


### GetITotalSignatureSigned

`func (o *EzsignfolderListElement) GetITotalSignatureSigned() int32`

GetITotalSignatureSigned returns the ITotalSignatureSigned field if non-nil, zero value otherwise.

### GetITotalSignatureSignedOk

`func (o *EzsignfolderListElement) GetITotalSignatureSignedOk() (*int32, bool)`

GetITotalSignatureSignedOk returns a tuple with the ITotalSignatureSigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetITotalSignatureSigned

`func (o *EzsignfolderListElement) SetITotalSignatureSigned(v int32)`

SetITotalSignatureSigned sets ITotalSignatureSigned field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


