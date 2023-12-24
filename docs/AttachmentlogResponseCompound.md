# AttachmentlogResponseCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FkiAttachmentID** | **int32** | The unique ID of the Attachment. | 
**FkiUserID** | **int32** | The unique ID of the User | 
**DtAttachmentlogDatetime** | **string** | The created date | 
**EAttachmentlogType** | [**FieldEAttachmentlogType**](FieldEAttachmentlogType.md) |  | 
**SAttachmentlogDetail** | Pointer to **string** | The additionnal detail | [optional] 

## Methods

### NewAttachmentlogResponseCompound

`func NewAttachmentlogResponseCompound(fkiAttachmentID int32, fkiUserID int32, dtAttachmentlogDatetime string, eAttachmentlogType FieldEAttachmentlogType, ) *AttachmentlogResponseCompound`

NewAttachmentlogResponseCompound instantiates a new AttachmentlogResponseCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachmentlogResponseCompoundWithDefaults

`func NewAttachmentlogResponseCompoundWithDefaults() *AttachmentlogResponseCompound`

NewAttachmentlogResponseCompoundWithDefaults instantiates a new AttachmentlogResponseCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFkiAttachmentID

`func (o *AttachmentlogResponseCompound) GetFkiAttachmentID() int32`

GetFkiAttachmentID returns the FkiAttachmentID field if non-nil, zero value otherwise.

### GetFkiAttachmentIDOk

`func (o *AttachmentlogResponseCompound) GetFkiAttachmentIDOk() (*int32, bool)`

GetFkiAttachmentIDOk returns a tuple with the FkiAttachmentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiAttachmentID

`func (o *AttachmentlogResponseCompound) SetFkiAttachmentID(v int32)`

SetFkiAttachmentID sets FkiAttachmentID field to given value.


### GetFkiUserID

`func (o *AttachmentlogResponseCompound) GetFkiUserID() int32`

GetFkiUserID returns the FkiUserID field if non-nil, zero value otherwise.

### GetFkiUserIDOk

`func (o *AttachmentlogResponseCompound) GetFkiUserIDOk() (*int32, bool)`

GetFkiUserIDOk returns a tuple with the FkiUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUserID

`func (o *AttachmentlogResponseCompound) SetFkiUserID(v int32)`

SetFkiUserID sets FkiUserID field to given value.


### GetDtAttachmentlogDatetime

`func (o *AttachmentlogResponseCompound) GetDtAttachmentlogDatetime() string`

GetDtAttachmentlogDatetime returns the DtAttachmentlogDatetime field if non-nil, zero value otherwise.

### GetDtAttachmentlogDatetimeOk

`func (o *AttachmentlogResponseCompound) GetDtAttachmentlogDatetimeOk() (*string, bool)`

GetDtAttachmentlogDatetimeOk returns a tuple with the DtAttachmentlogDatetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtAttachmentlogDatetime

`func (o *AttachmentlogResponseCompound) SetDtAttachmentlogDatetime(v string)`

SetDtAttachmentlogDatetime sets DtAttachmentlogDatetime field to given value.


### GetEAttachmentlogType

`func (o *AttachmentlogResponseCompound) GetEAttachmentlogType() FieldEAttachmentlogType`

GetEAttachmentlogType returns the EAttachmentlogType field if non-nil, zero value otherwise.

### GetEAttachmentlogTypeOk

`func (o *AttachmentlogResponseCompound) GetEAttachmentlogTypeOk() (*FieldEAttachmentlogType, bool)`

GetEAttachmentlogTypeOk returns a tuple with the EAttachmentlogType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEAttachmentlogType

`func (o *AttachmentlogResponseCompound) SetEAttachmentlogType(v FieldEAttachmentlogType)`

SetEAttachmentlogType sets EAttachmentlogType field to given value.


### GetSAttachmentlogDetail

`func (o *AttachmentlogResponseCompound) GetSAttachmentlogDetail() string`

GetSAttachmentlogDetail returns the SAttachmentlogDetail field if non-nil, zero value otherwise.

### GetSAttachmentlogDetailOk

`func (o *AttachmentlogResponseCompound) GetSAttachmentlogDetailOk() (*string, bool)`

GetSAttachmentlogDetailOk returns a tuple with the SAttachmentlogDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAttachmentlogDetail

`func (o *AttachmentlogResponseCompound) SetSAttachmentlogDetail(v string)`

SetSAttachmentlogDetail sets SAttachmentlogDetail field to given value.

### HasSAttachmentlogDetail

`func (o *AttachmentlogResponseCompound) HasSAttachmentlogDetail() bool`

HasSAttachmentlogDetail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


