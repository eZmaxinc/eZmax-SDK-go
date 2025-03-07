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
**DtEzsignfolderDelayedsenddate** | Pointer to **string** | The date and time at which the Ezsignfolder will be sent in the future. | [optional] 
**DtEzsignfolderSentdate** | Pointer to **string** | The date and time at which the Ezsignfolder was sent the last time. | [optional] 
**DtEzsignfolderDuedate** | Pointer to **string** | The maximum date and time at which the Ezsignfolder can be signed. | [optional] 
**IEzsigndocument** | **int32** | The total number of Ezsigndocument in the folder | 
**IEzsigndocumentEdm** | **int32** | The total number of Ezsigndocument in the folder that were saved in the edm system | 
**IEzsignsignature** | **int32** | The total number of signature blocks in all Ezsigndocuments in the folder | 
**IEzsignsignatureSigned** | **int32** | The total number of already signed signature blocks in all Ezsigndocuments in the folder | 
**IEzsignformfieldgroup** | **int32** | The total number of Ezsignformfieldgroup in all Ezsigndocuments in the folder | 
**IEzsignformfieldgroupCompleted** | **int32** | The total number of completed Ezsignformfieldgroup in all Ezsigndocuments in the folder | 
**BEzsignformHasdependencies** | Pointer to **bool** | Whether the Ezsignform/Ezsignsignatures has dependencies or not | [optional] 
**DEzsignfolderCompletedpercentage** | **string** | Percentage of Ezsignform/Ezsignsignatures has completed | 
**DEzsignfolderFormcompletedpercentage** | **string** | Percentage of Ezsignform has completed | 
**DEzsignfolderSignaturecompletedpercentage** | **string** | Percentage of Ezsignsignatures has signed | 
**DtEzsignfolderClose** | Pointer to **string** | The date and time at which the Ezsignfolder was closed. Either by applying the last signature or by completing it prematurely. | [optional] 
**DtEzsignfolderArchive** | Pointer to **string** | The date and time at which the Ezsignfolder was archived. | [optional] 
**DtEzsignfolderDispose** | Pointer to **string** | The date and time at which the Ezsignfolder was disposed. | [optional] 
**BEzsignfolderSigner** | Pointer to **bool** | Whether the Ezsignfolder has an Ezsignsignatures that need to be signed or an Ezsignformfieldgroups that need to be filled by the current user | [optional] 

## Methods

### NewEzsignfolderListElement

`func NewEzsignfolderListElement(pkiEzsignfolderID int32, fkiEzsignfoldertypeID int32, eEzsignfoldertypePrivacylevel FieldEEzsignfoldertypePrivacylevel, sEzsignfoldertypeNameX string, sEzsignfolderDescription string, eEzsignfolderStep FieldEEzsignfolderStep, dtCreatedDate string, iEzsigndocument int32, iEzsigndocumentEdm int32, iEzsignsignature int32, iEzsignsignatureSigned int32, iEzsignformfieldgroup int32, iEzsignformfieldgroupCompleted int32, dEzsignfolderCompletedpercentage string, dEzsignfolderFormcompletedpercentage string, dEzsignfolderSignaturecompletedpercentage string, ) *EzsignfolderListElement`

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


### GetDtEzsignfolderDelayedsenddate

`func (o *EzsignfolderListElement) GetDtEzsignfolderDelayedsenddate() string`

GetDtEzsignfolderDelayedsenddate returns the DtEzsignfolderDelayedsenddate field if non-nil, zero value otherwise.

### GetDtEzsignfolderDelayedsenddateOk

`func (o *EzsignfolderListElement) GetDtEzsignfolderDelayedsenddateOk() (*string, bool)`

GetDtEzsignfolderDelayedsenddateOk returns a tuple with the DtEzsignfolderDelayedsenddate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtEzsignfolderDelayedsenddate

`func (o *EzsignfolderListElement) SetDtEzsignfolderDelayedsenddate(v string)`

SetDtEzsignfolderDelayedsenddate sets DtEzsignfolderDelayedsenddate field to given value.

### HasDtEzsignfolderDelayedsenddate

`func (o *EzsignfolderListElement) HasDtEzsignfolderDelayedsenddate() bool`

HasDtEzsignfolderDelayedsenddate returns a boolean if a field has been set.

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

### HasDtEzsignfolderSentdate

`func (o *EzsignfolderListElement) HasDtEzsignfolderSentdate() bool`

HasDtEzsignfolderSentdate returns a boolean if a field has been set.

### GetDtEzsignfolderDuedate

`func (o *EzsignfolderListElement) GetDtEzsignfolderDuedate() string`

GetDtEzsignfolderDuedate returns the DtEzsignfolderDuedate field if non-nil, zero value otherwise.

### GetDtEzsignfolderDuedateOk

`func (o *EzsignfolderListElement) GetDtEzsignfolderDuedateOk() (*string, bool)`

GetDtEzsignfolderDuedateOk returns a tuple with the DtEzsignfolderDuedate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtEzsignfolderDuedate

`func (o *EzsignfolderListElement) SetDtEzsignfolderDuedate(v string)`

SetDtEzsignfolderDuedate sets DtEzsignfolderDuedate field to given value.

### HasDtEzsignfolderDuedate

`func (o *EzsignfolderListElement) HasDtEzsignfolderDuedate() bool`

HasDtEzsignfolderDuedate returns a boolean if a field has been set.

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


### GetIEzsignformfieldgroup

`func (o *EzsignfolderListElement) GetIEzsignformfieldgroup() int32`

GetIEzsignformfieldgroup returns the IEzsignformfieldgroup field if non-nil, zero value otherwise.

### GetIEzsignformfieldgroupOk

`func (o *EzsignfolderListElement) GetIEzsignformfieldgroupOk() (*int32, bool)`

GetIEzsignformfieldgroupOk returns a tuple with the IEzsignformfieldgroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignformfieldgroup

`func (o *EzsignfolderListElement) SetIEzsignformfieldgroup(v int32)`

SetIEzsignformfieldgroup sets IEzsignformfieldgroup field to given value.


### GetIEzsignformfieldgroupCompleted

`func (o *EzsignfolderListElement) GetIEzsignformfieldgroupCompleted() int32`

GetIEzsignformfieldgroupCompleted returns the IEzsignformfieldgroupCompleted field if non-nil, zero value otherwise.

### GetIEzsignformfieldgroupCompletedOk

`func (o *EzsignfolderListElement) GetIEzsignformfieldgroupCompletedOk() (*int32, bool)`

GetIEzsignformfieldgroupCompletedOk returns a tuple with the IEzsignformfieldgroupCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignformfieldgroupCompleted

`func (o *EzsignfolderListElement) SetIEzsignformfieldgroupCompleted(v int32)`

SetIEzsignformfieldgroupCompleted sets IEzsignformfieldgroupCompleted field to given value.


### GetBEzsignformHasdependencies

`func (o *EzsignfolderListElement) GetBEzsignformHasdependencies() bool`

GetBEzsignformHasdependencies returns the BEzsignformHasdependencies field if non-nil, zero value otherwise.

### GetBEzsignformHasdependenciesOk

`func (o *EzsignfolderListElement) GetBEzsignformHasdependenciesOk() (*bool, bool)`

GetBEzsignformHasdependenciesOk returns a tuple with the BEzsignformHasdependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignformHasdependencies

`func (o *EzsignfolderListElement) SetBEzsignformHasdependencies(v bool)`

SetBEzsignformHasdependencies sets BEzsignformHasdependencies field to given value.

### HasBEzsignformHasdependencies

`func (o *EzsignfolderListElement) HasBEzsignformHasdependencies() bool`

HasBEzsignformHasdependencies returns a boolean if a field has been set.

### GetDEzsignfolderCompletedpercentage

`func (o *EzsignfolderListElement) GetDEzsignfolderCompletedpercentage() string`

GetDEzsignfolderCompletedpercentage returns the DEzsignfolderCompletedpercentage field if non-nil, zero value otherwise.

### GetDEzsignfolderCompletedpercentageOk

`func (o *EzsignfolderListElement) GetDEzsignfolderCompletedpercentageOk() (*string, bool)`

GetDEzsignfolderCompletedpercentageOk returns a tuple with the DEzsignfolderCompletedpercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDEzsignfolderCompletedpercentage

`func (o *EzsignfolderListElement) SetDEzsignfolderCompletedpercentage(v string)`

SetDEzsignfolderCompletedpercentage sets DEzsignfolderCompletedpercentage field to given value.


### GetDEzsignfolderFormcompletedpercentage

`func (o *EzsignfolderListElement) GetDEzsignfolderFormcompletedpercentage() string`

GetDEzsignfolderFormcompletedpercentage returns the DEzsignfolderFormcompletedpercentage field if non-nil, zero value otherwise.

### GetDEzsignfolderFormcompletedpercentageOk

`func (o *EzsignfolderListElement) GetDEzsignfolderFormcompletedpercentageOk() (*string, bool)`

GetDEzsignfolderFormcompletedpercentageOk returns a tuple with the DEzsignfolderFormcompletedpercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDEzsignfolderFormcompletedpercentage

`func (o *EzsignfolderListElement) SetDEzsignfolderFormcompletedpercentage(v string)`

SetDEzsignfolderFormcompletedpercentage sets DEzsignfolderFormcompletedpercentage field to given value.


### GetDEzsignfolderSignaturecompletedpercentage

`func (o *EzsignfolderListElement) GetDEzsignfolderSignaturecompletedpercentage() string`

GetDEzsignfolderSignaturecompletedpercentage returns the DEzsignfolderSignaturecompletedpercentage field if non-nil, zero value otherwise.

### GetDEzsignfolderSignaturecompletedpercentageOk

`func (o *EzsignfolderListElement) GetDEzsignfolderSignaturecompletedpercentageOk() (*string, bool)`

GetDEzsignfolderSignaturecompletedpercentageOk returns a tuple with the DEzsignfolderSignaturecompletedpercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDEzsignfolderSignaturecompletedpercentage

`func (o *EzsignfolderListElement) SetDEzsignfolderSignaturecompletedpercentage(v string)`

SetDEzsignfolderSignaturecompletedpercentage sets DEzsignfolderSignaturecompletedpercentage field to given value.


### GetDtEzsignfolderClose

`func (o *EzsignfolderListElement) GetDtEzsignfolderClose() string`

GetDtEzsignfolderClose returns the DtEzsignfolderClose field if non-nil, zero value otherwise.

### GetDtEzsignfolderCloseOk

`func (o *EzsignfolderListElement) GetDtEzsignfolderCloseOk() (*string, bool)`

GetDtEzsignfolderCloseOk returns a tuple with the DtEzsignfolderClose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtEzsignfolderClose

`func (o *EzsignfolderListElement) SetDtEzsignfolderClose(v string)`

SetDtEzsignfolderClose sets DtEzsignfolderClose field to given value.

### HasDtEzsignfolderClose

`func (o *EzsignfolderListElement) HasDtEzsignfolderClose() bool`

HasDtEzsignfolderClose returns a boolean if a field has been set.

### GetDtEzsignfolderArchive

`func (o *EzsignfolderListElement) GetDtEzsignfolderArchive() string`

GetDtEzsignfolderArchive returns the DtEzsignfolderArchive field if non-nil, zero value otherwise.

### GetDtEzsignfolderArchiveOk

`func (o *EzsignfolderListElement) GetDtEzsignfolderArchiveOk() (*string, bool)`

GetDtEzsignfolderArchiveOk returns a tuple with the DtEzsignfolderArchive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtEzsignfolderArchive

`func (o *EzsignfolderListElement) SetDtEzsignfolderArchive(v string)`

SetDtEzsignfolderArchive sets DtEzsignfolderArchive field to given value.

### HasDtEzsignfolderArchive

`func (o *EzsignfolderListElement) HasDtEzsignfolderArchive() bool`

HasDtEzsignfolderArchive returns a boolean if a field has been set.

### GetDtEzsignfolderDispose

`func (o *EzsignfolderListElement) GetDtEzsignfolderDispose() string`

GetDtEzsignfolderDispose returns the DtEzsignfolderDispose field if non-nil, zero value otherwise.

### GetDtEzsignfolderDisposeOk

`func (o *EzsignfolderListElement) GetDtEzsignfolderDisposeOk() (*string, bool)`

GetDtEzsignfolderDisposeOk returns a tuple with the DtEzsignfolderDispose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtEzsignfolderDispose

`func (o *EzsignfolderListElement) SetDtEzsignfolderDispose(v string)`

SetDtEzsignfolderDispose sets DtEzsignfolderDispose field to given value.

### HasDtEzsignfolderDispose

`func (o *EzsignfolderListElement) HasDtEzsignfolderDispose() bool`

HasDtEzsignfolderDispose returns a boolean if a field has been set.

### GetBEzsignfolderSigner

`func (o *EzsignfolderListElement) GetBEzsignfolderSigner() bool`

GetBEzsignfolderSigner returns the BEzsignfolderSigner field if non-nil, zero value otherwise.

### GetBEzsignfolderSignerOk

`func (o *EzsignfolderListElement) GetBEzsignfolderSignerOk() (*bool, bool)`

GetBEzsignfolderSignerOk returns a tuple with the BEzsignfolderSigner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfolderSigner

`func (o *EzsignfolderListElement) SetBEzsignfolderSigner(v bool)`

SetBEzsignfolderSigner sets BEzsignfolderSigner field to given value.

### HasBEzsignfolderSigner

`func (o *EzsignfolderListElement) HasBEzsignfolderSigner() bool`

HasBEzsignfolderSigner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


