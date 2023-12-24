# AttachmentlogResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FkiAttachmentID** | **int32** | The unique ID of the Attachment. | 
**FkiUserID** | **int32** | The unique ID of the User | 
**DtAttachmentlogDatetime** | **string** | The created date | 
**EAttachmentlogType** | [**FieldEAttachmentlogType**](FieldEAttachmentlogType.md) |  | 
**SAttachmentlogDetail** | Pointer to **string** | The additionnal detail | [optional] 

## Methods

### NewAttachmentlogResponse

`func NewAttachmentlogResponse(fkiAttachmentID int32, fkiUserID int32, dtAttachmentlogDatetime string, eAttachmentlogType FieldEAttachmentlogType, ) *AttachmentlogResponse`

NewAttachmentlogResponse instantiates a new AttachmentlogResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachmentlogResponseWithDefaults

`func NewAttachmentlogResponseWithDefaults() *AttachmentlogResponse`

NewAttachmentlogResponseWithDefaults instantiates a new AttachmentlogResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFkiAttachmentID

`func (o *AttachmentlogResponse) GetFkiAttachmentID() int32`

GetFkiAttachmentID returns the FkiAttachmentID field if non-nil, zero value otherwise.

### GetFkiAttachmentIDOk

`func (o *AttachmentlogResponse) GetFkiAttachmentIDOk() (*int32, bool)`

GetFkiAttachmentIDOk returns a tuple with the FkiAttachmentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiAttachmentID

`func (o *AttachmentlogResponse) SetFkiAttachmentID(v int32)`

SetFkiAttachmentID sets FkiAttachmentID field to given value.


### GetFkiUserID

`func (o *AttachmentlogResponse) GetFkiUserID() int32`

GetFkiUserID returns the FkiUserID field if non-nil, zero value otherwise.

### GetFkiUserIDOk

`func (o *AttachmentlogResponse) GetFkiUserIDOk() (*int32, bool)`

GetFkiUserIDOk returns a tuple with the FkiUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUserID

`func (o *AttachmentlogResponse) SetFkiUserID(v int32)`

SetFkiUserID sets FkiUserID field to given value.


### GetDtAttachmentlogDatetime

`func (o *AttachmentlogResponse) GetDtAttachmentlogDatetime() string`

GetDtAttachmentlogDatetime returns the DtAttachmentlogDatetime field if non-nil, zero value otherwise.

### GetDtAttachmentlogDatetimeOk

`func (o *AttachmentlogResponse) GetDtAttachmentlogDatetimeOk() (*string, bool)`

GetDtAttachmentlogDatetimeOk returns a tuple with the DtAttachmentlogDatetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtAttachmentlogDatetime

`func (o *AttachmentlogResponse) SetDtAttachmentlogDatetime(v string)`

SetDtAttachmentlogDatetime sets DtAttachmentlogDatetime field to given value.


### GetEAttachmentlogType

`func (o *AttachmentlogResponse) GetEAttachmentlogType() FieldEAttachmentlogType`

GetEAttachmentlogType returns the EAttachmentlogType field if non-nil, zero value otherwise.

### GetEAttachmentlogTypeOk

`func (o *AttachmentlogResponse) GetEAttachmentlogTypeOk() (*FieldEAttachmentlogType, bool)`

GetEAttachmentlogTypeOk returns a tuple with the EAttachmentlogType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEAttachmentlogType

`func (o *AttachmentlogResponse) SetEAttachmentlogType(v FieldEAttachmentlogType)`

SetEAttachmentlogType sets EAttachmentlogType field to given value.


### GetSAttachmentlogDetail

`func (o *AttachmentlogResponse) GetSAttachmentlogDetail() string`

GetSAttachmentlogDetail returns the SAttachmentlogDetail field if non-nil, zero value otherwise.

### GetSAttachmentlogDetailOk

`func (o *AttachmentlogResponse) GetSAttachmentlogDetailOk() (*string, bool)`

GetSAttachmentlogDetailOk returns a tuple with the SAttachmentlogDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAttachmentlogDetail

`func (o *AttachmentlogResponse) SetSAttachmentlogDetail(v string)`

SetSAttachmentlogDetail sets SAttachmentlogDetail field to given value.

### HasSAttachmentlogDetail

`func (o *AttachmentlogResponse) HasSAttachmentlogDetail() bool`

HasSAttachmentlogDetail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


