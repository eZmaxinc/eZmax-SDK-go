# EzsigndocumentGetFormDataV1ResponseMPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsigndocumentID** | **int32** | The unique ID of the Ezsigndocument | 
**FkiEzsignfolderID** | **int32** | The unique ID of the Ezsignfolder | 
**SEzsigndocumentName** | **string** | The name of the document that will be presented to Ezsignfoldersignerassociations | 
**DtModifiedDate** | **string** | The date and time at which the object was last modified | 
**AObjFormDataSigner** | [**[]CustomFormDataSignerResponse**](CustomFormDataSignerResponse.md) |  | 

## Methods

### NewEzsigndocumentGetFormDataV1ResponseMPayload

`func NewEzsigndocumentGetFormDataV1ResponseMPayload(pkiEzsigndocumentID int32, fkiEzsignfolderID int32, sEzsigndocumentName string, dtModifiedDate string, aObjFormDataSigner []CustomFormDataSignerResponse, ) *EzsigndocumentGetFormDataV1ResponseMPayload`

NewEzsigndocumentGetFormDataV1ResponseMPayload instantiates a new EzsigndocumentGetFormDataV1ResponseMPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigndocumentGetFormDataV1ResponseMPayloadWithDefaults

`func NewEzsigndocumentGetFormDataV1ResponseMPayloadWithDefaults() *EzsigndocumentGetFormDataV1ResponseMPayload`

NewEzsigndocumentGetFormDataV1ResponseMPayloadWithDefaults instantiates a new EzsigndocumentGetFormDataV1ResponseMPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsigndocumentID

`func (o *EzsigndocumentGetFormDataV1ResponseMPayload) GetPkiEzsigndocumentID() int32`

GetPkiEzsigndocumentID returns the PkiEzsigndocumentID field if non-nil, zero value otherwise.

### GetPkiEzsigndocumentIDOk

`func (o *EzsigndocumentGetFormDataV1ResponseMPayload) GetPkiEzsigndocumentIDOk() (*int32, bool)`

GetPkiEzsigndocumentIDOk returns a tuple with the PkiEzsigndocumentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsigndocumentID

`func (o *EzsigndocumentGetFormDataV1ResponseMPayload) SetPkiEzsigndocumentID(v int32)`

SetPkiEzsigndocumentID sets PkiEzsigndocumentID field to given value.


### GetFkiEzsignfolderID

`func (o *EzsigndocumentGetFormDataV1ResponseMPayload) GetFkiEzsignfolderID() int32`

GetFkiEzsignfolderID returns the FkiEzsignfolderID field if non-nil, zero value otherwise.

### GetFkiEzsignfolderIDOk

`func (o *EzsigndocumentGetFormDataV1ResponseMPayload) GetFkiEzsignfolderIDOk() (*int32, bool)`

GetFkiEzsignfolderIDOk returns a tuple with the FkiEzsignfolderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignfolderID

`func (o *EzsigndocumentGetFormDataV1ResponseMPayload) SetFkiEzsignfolderID(v int32)`

SetFkiEzsignfolderID sets FkiEzsignfolderID field to given value.


### GetSEzsigndocumentName

`func (o *EzsigndocumentGetFormDataV1ResponseMPayload) GetSEzsigndocumentName() string`

GetSEzsigndocumentName returns the SEzsigndocumentName field if non-nil, zero value otherwise.

### GetSEzsigndocumentNameOk

`func (o *EzsigndocumentGetFormDataV1ResponseMPayload) GetSEzsigndocumentNameOk() (*string, bool)`

GetSEzsigndocumentNameOk returns a tuple with the SEzsigndocumentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigndocumentName

`func (o *EzsigndocumentGetFormDataV1ResponseMPayload) SetSEzsigndocumentName(v string)`

SetSEzsigndocumentName sets SEzsigndocumentName field to given value.


### GetDtModifiedDate

`func (o *EzsigndocumentGetFormDataV1ResponseMPayload) GetDtModifiedDate() string`

GetDtModifiedDate returns the DtModifiedDate field if non-nil, zero value otherwise.

### GetDtModifiedDateOk

`func (o *EzsigndocumentGetFormDataV1ResponseMPayload) GetDtModifiedDateOk() (*string, bool)`

GetDtModifiedDateOk returns a tuple with the DtModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtModifiedDate

`func (o *EzsigndocumentGetFormDataV1ResponseMPayload) SetDtModifiedDate(v string)`

SetDtModifiedDate sets DtModifiedDate field to given value.


### GetAObjFormDataSigner

`func (o *EzsigndocumentGetFormDataV1ResponseMPayload) GetAObjFormDataSigner() []CustomFormDataSignerResponse`

GetAObjFormDataSigner returns the AObjFormDataSigner field if non-nil, zero value otherwise.

### GetAObjFormDataSignerOk

`func (o *EzsigndocumentGetFormDataV1ResponseMPayload) GetAObjFormDataSignerOk() (*[]CustomFormDataSignerResponse, bool)`

GetAObjFormDataSignerOk returns a tuple with the AObjFormDataSigner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjFormDataSigner

`func (o *EzsigndocumentGetFormDataV1ResponseMPayload) SetAObjFormDataSigner(v []CustomFormDataSignerResponse)`

SetAObjFormDataSigner sets AObjFormDataSigner field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


