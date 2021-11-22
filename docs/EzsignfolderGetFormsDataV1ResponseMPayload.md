# EzsignfolderGetFormsDataV1ResponseMPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsignfolderID** | **int32** | The unique ID of the Ezsignfolder | 
**SEzsignfolderDescription** | **string** | The description of the Ezsignfolder | 
**AObjFormDataDocument** | [**[]CustomFormDataDocumentResponse**](CustomFormDataDocumentResponse.md) |  | 

## Methods

### NewEzsignfolderGetFormsDataV1ResponseMPayload

`func NewEzsignfolderGetFormsDataV1ResponseMPayload(pkiEzsignfolderID int32, sEzsignfolderDescription string, aObjFormDataDocument []CustomFormDataDocumentResponse, ) *EzsignfolderGetFormsDataV1ResponseMPayload`

NewEzsignfolderGetFormsDataV1ResponseMPayload instantiates a new EzsignfolderGetFormsDataV1ResponseMPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsignfolderGetFormsDataV1ResponseMPayloadWithDefaults

`func NewEzsignfolderGetFormsDataV1ResponseMPayloadWithDefaults() *EzsignfolderGetFormsDataV1ResponseMPayload`

NewEzsignfolderGetFormsDataV1ResponseMPayloadWithDefaults instantiates a new EzsignfolderGetFormsDataV1ResponseMPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsignfolderID

`func (o *EzsignfolderGetFormsDataV1ResponseMPayload) GetPkiEzsignfolderID() int32`

GetPkiEzsignfolderID returns the PkiEzsignfolderID field if non-nil, zero value otherwise.

### GetPkiEzsignfolderIDOk

`func (o *EzsignfolderGetFormsDataV1ResponseMPayload) GetPkiEzsignfolderIDOk() (*int32, bool)`

GetPkiEzsignfolderIDOk returns a tuple with the PkiEzsignfolderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsignfolderID

`func (o *EzsignfolderGetFormsDataV1ResponseMPayload) SetPkiEzsignfolderID(v int32)`

SetPkiEzsignfolderID sets PkiEzsignfolderID field to given value.


### GetSEzsignfolderDescription

`func (o *EzsignfolderGetFormsDataV1ResponseMPayload) GetSEzsignfolderDescription() string`

GetSEzsignfolderDescription returns the SEzsignfolderDescription field if non-nil, zero value otherwise.

### GetSEzsignfolderDescriptionOk

`func (o *EzsignfolderGetFormsDataV1ResponseMPayload) GetSEzsignfolderDescriptionOk() (*string, bool)`

GetSEzsignfolderDescriptionOk returns a tuple with the SEzsignfolderDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignfolderDescription

`func (o *EzsignfolderGetFormsDataV1ResponseMPayload) SetSEzsignfolderDescription(v string)`

SetSEzsignfolderDescription sets SEzsignfolderDescription field to given value.


### GetAObjFormDataDocument

`func (o *EzsignfolderGetFormsDataV1ResponseMPayload) GetAObjFormDataDocument() []CustomFormDataDocumentResponse`

GetAObjFormDataDocument returns the AObjFormDataDocument field if non-nil, zero value otherwise.

### GetAObjFormDataDocumentOk

`func (o *EzsignfolderGetFormsDataV1ResponseMPayload) GetAObjFormDataDocumentOk() (*[]CustomFormDataDocumentResponse, bool)`

GetAObjFormDataDocumentOk returns a tuple with the AObjFormDataDocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjFormDataDocument

`func (o *EzsignfolderGetFormsDataV1ResponseMPayload) SetAObjFormDataDocument(v []CustomFormDataDocumentResponse)`

SetAObjFormDataDocument sets AObjFormDataDocument field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


