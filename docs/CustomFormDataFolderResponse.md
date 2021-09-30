# CustomFormDataFolderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsignfolderID** | **int32** | The unique ID of the Ezsignfolder | 
**SEzsignfolderDescription** | **string** | The description of the Ezsign Folder | 
**AObjFormDataDocument** | [**[]CustomFormDataDocumentResponse**](CustomFormDataDocumentResponse.md) |  | 

## Methods

### NewCustomFormDataFolderResponse

`func NewCustomFormDataFolderResponse(pkiEzsignfolderID int32, sEzsignfolderDescription string, aObjFormDataDocument []CustomFormDataDocumentResponse, ) *CustomFormDataFolderResponse`

NewCustomFormDataFolderResponse instantiates a new CustomFormDataFolderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomFormDataFolderResponseWithDefaults

`func NewCustomFormDataFolderResponseWithDefaults() *CustomFormDataFolderResponse`

NewCustomFormDataFolderResponseWithDefaults instantiates a new CustomFormDataFolderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsignfolderID

`func (o *CustomFormDataFolderResponse) GetPkiEzsignfolderID() int32`

GetPkiEzsignfolderID returns the PkiEzsignfolderID field if non-nil, zero value otherwise.

### GetPkiEzsignfolderIDOk

`func (o *CustomFormDataFolderResponse) GetPkiEzsignfolderIDOk() (*int32, bool)`

GetPkiEzsignfolderIDOk returns a tuple with the PkiEzsignfolderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsignfolderID

`func (o *CustomFormDataFolderResponse) SetPkiEzsignfolderID(v int32)`

SetPkiEzsignfolderID sets PkiEzsignfolderID field to given value.


### GetSEzsignfolderDescription

`func (o *CustomFormDataFolderResponse) GetSEzsignfolderDescription() string`

GetSEzsignfolderDescription returns the SEzsignfolderDescription field if non-nil, zero value otherwise.

### GetSEzsignfolderDescriptionOk

`func (o *CustomFormDataFolderResponse) GetSEzsignfolderDescriptionOk() (*string, bool)`

GetSEzsignfolderDescriptionOk returns a tuple with the SEzsignfolderDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignfolderDescription

`func (o *CustomFormDataFolderResponse) SetSEzsignfolderDescription(v string)`

SetSEzsignfolderDescription sets SEzsignfolderDescription field to given value.


### GetAObjFormDataDocument

`func (o *CustomFormDataFolderResponse) GetAObjFormDataDocument() []CustomFormDataDocumentResponse`

GetAObjFormDataDocument returns the AObjFormDataDocument field if non-nil, zero value otherwise.

### GetAObjFormDataDocumentOk

`func (o *CustomFormDataFolderResponse) GetAObjFormDataDocumentOk() (*[]CustomFormDataDocumentResponse, bool)`

GetAObjFormDataDocumentOk returns a tuple with the AObjFormDataDocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjFormDataDocument

`func (o *CustomFormDataFolderResponse) SetAObjFormDataDocument(v []CustomFormDataDocumentResponse)`

SetAObjFormDataDocument sets AObjFormDataDocument field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


