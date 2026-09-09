# DepositImportIntoEDMV1Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjDebugPayload** | [**CommonResponseObjDebugPayload**](CommonResponseObjDebugPayload.md) |  | 
**ObjDebug** | Pointer to [**CommonResponseObjDebug**](CommonResponseObjDebug.md) |  | [optional] 
**MPayload** | [**DepositImportIntoEDMV1ResponseMPayload**](DepositImportIntoEDMV1ResponseMPayload.md) |  | 

## Methods

### NewDepositImportIntoEDMV1Response

`func NewDepositImportIntoEDMV1Response(objDebugPayload CommonResponseObjDebugPayload, mPayload DepositImportIntoEDMV1ResponseMPayload, ) *DepositImportIntoEDMV1Response`

NewDepositImportIntoEDMV1Response instantiates a new DepositImportIntoEDMV1Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDepositImportIntoEDMV1ResponseWithDefaults

`func NewDepositImportIntoEDMV1ResponseWithDefaults() *DepositImportIntoEDMV1Response`

NewDepositImportIntoEDMV1ResponseWithDefaults instantiates a new DepositImportIntoEDMV1Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjDebugPayload

`func (o *DepositImportIntoEDMV1Response) GetObjDebugPayload() CommonResponseObjDebugPayload`

GetObjDebugPayload returns the ObjDebugPayload field if non-nil, zero value otherwise.

### GetObjDebugPayloadOk

`func (o *DepositImportIntoEDMV1Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayload, bool)`

GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjDebugPayload

`func (o *DepositImportIntoEDMV1Response) SetObjDebugPayload(v CommonResponseObjDebugPayload)`

SetObjDebugPayload sets ObjDebugPayload field to given value.


### GetObjDebug

`func (o *DepositImportIntoEDMV1Response) GetObjDebug() CommonResponseObjDebug`

GetObjDebug returns the ObjDebug field if non-nil, zero value otherwise.

### GetObjDebugOk

`func (o *DepositImportIntoEDMV1Response) GetObjDebugOk() (*CommonResponseObjDebug, bool)`

GetObjDebugOk returns a tuple with the ObjDebug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjDebug

`func (o *DepositImportIntoEDMV1Response) SetObjDebug(v CommonResponseObjDebug)`

SetObjDebug sets ObjDebug field to given value.

### HasObjDebug

`func (o *DepositImportIntoEDMV1Response) HasObjDebug() bool`

HasObjDebug returns a boolean if a field has been set.

### GetMPayload

`func (o *DepositImportIntoEDMV1Response) GetMPayload() DepositImportIntoEDMV1ResponseMPayload`

GetMPayload returns the MPayload field if non-nil, zero value otherwise.

### GetMPayloadOk

`func (o *DepositImportIntoEDMV1Response) GetMPayloadOk() (*DepositImportIntoEDMV1ResponseMPayload, bool)`

GetMPayloadOk returns a tuple with the MPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMPayload

`func (o *DepositImportIntoEDMV1Response) SetMPayload(v DepositImportIntoEDMV1ResponseMPayload)`

SetMPayload sets MPayload field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


