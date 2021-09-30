# CustomFormDataDocumentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsigndocumentID** | **int32** | The unique ID of the Ezsigndocument | 
**FkiEzsignfolderID** | **int32** | The unique ID of the Ezsignfolder | 
**SEzsigndocumentName** | **string** | The name of the document that will be presented to Ezsignfoldersignerassociations | 
**DtModifiedDate** | **string** | The date and time at which the object was last modified | 
**AObjFormDataSigner** | [**[]CustomFormDataSignerResponse**](CustomFormDataSignerResponse.md) |  | 

## Methods

### NewCustomFormDataDocumentResponse

`func NewCustomFormDataDocumentResponse(pkiEzsigndocumentID int32, fkiEzsignfolderID int32, sEzsigndocumentName string, dtModifiedDate string, aObjFormDataSigner []CustomFormDataSignerResponse, ) *CustomFormDataDocumentResponse`

NewCustomFormDataDocumentResponse instantiates a new CustomFormDataDocumentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomFormDataDocumentResponseWithDefaults

`func NewCustomFormDataDocumentResponseWithDefaults() *CustomFormDataDocumentResponse`

NewCustomFormDataDocumentResponseWithDefaults instantiates a new CustomFormDataDocumentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsigndocumentID

`func (o *CustomFormDataDocumentResponse) GetPkiEzsigndocumentID() int32`

GetPkiEzsigndocumentID returns the PkiEzsigndocumentID field if non-nil, zero value otherwise.

### GetPkiEzsigndocumentIDOk

`func (o *CustomFormDataDocumentResponse) GetPkiEzsigndocumentIDOk() (*int32, bool)`

GetPkiEzsigndocumentIDOk returns a tuple with the PkiEzsigndocumentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsigndocumentID

`func (o *CustomFormDataDocumentResponse) SetPkiEzsigndocumentID(v int32)`

SetPkiEzsigndocumentID sets PkiEzsigndocumentID field to given value.


### GetFkiEzsignfolderID

`func (o *CustomFormDataDocumentResponse) GetFkiEzsignfolderID() int32`

GetFkiEzsignfolderID returns the FkiEzsignfolderID field if non-nil, zero value otherwise.

### GetFkiEzsignfolderIDOk

`func (o *CustomFormDataDocumentResponse) GetFkiEzsignfolderIDOk() (*int32, bool)`

GetFkiEzsignfolderIDOk returns a tuple with the FkiEzsignfolderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignfolderID

`func (o *CustomFormDataDocumentResponse) SetFkiEzsignfolderID(v int32)`

SetFkiEzsignfolderID sets FkiEzsignfolderID field to given value.


### GetSEzsigndocumentName

`func (o *CustomFormDataDocumentResponse) GetSEzsigndocumentName() string`

GetSEzsigndocumentName returns the SEzsigndocumentName field if non-nil, zero value otherwise.

### GetSEzsigndocumentNameOk

`func (o *CustomFormDataDocumentResponse) GetSEzsigndocumentNameOk() (*string, bool)`

GetSEzsigndocumentNameOk returns a tuple with the SEzsigndocumentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigndocumentName

`func (o *CustomFormDataDocumentResponse) SetSEzsigndocumentName(v string)`

SetSEzsigndocumentName sets SEzsigndocumentName field to given value.


### GetDtModifiedDate

`func (o *CustomFormDataDocumentResponse) GetDtModifiedDate() string`

GetDtModifiedDate returns the DtModifiedDate field if non-nil, zero value otherwise.

### GetDtModifiedDateOk

`func (o *CustomFormDataDocumentResponse) GetDtModifiedDateOk() (*string, bool)`

GetDtModifiedDateOk returns a tuple with the DtModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtModifiedDate

`func (o *CustomFormDataDocumentResponse) SetDtModifiedDate(v string)`

SetDtModifiedDate sets DtModifiedDate field to given value.


### GetAObjFormDataSigner

`func (o *CustomFormDataDocumentResponse) GetAObjFormDataSigner() []CustomFormDataSignerResponse`

GetAObjFormDataSigner returns the AObjFormDataSigner field if non-nil, zero value otherwise.

### GetAObjFormDataSignerOk

`func (o *CustomFormDataDocumentResponse) GetAObjFormDataSignerOk() (*[]CustomFormDataSignerResponse, bool)`

GetAObjFormDataSignerOk returns a tuple with the AObjFormDataSigner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjFormDataSigner

`func (o *CustomFormDataDocumentResponse) SetAObjFormDataSigner(v []CustomFormDataSignerResponse)`

SetAObjFormDataSigner sets AObjFormDataSigner field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


