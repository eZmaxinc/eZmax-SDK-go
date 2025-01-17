# UserstagedGetListV1Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjDebugPayload** | [**CommonResponseObjDebugPayloadGetList**](CommonResponseObjDebugPayloadGetList.md) |  | 
**ObjDebug** | Pointer to [**CommonResponseObjDebug**](CommonResponseObjDebug.md) |  | [optional] 
**MPayload** | [**UserstagedGetListV1ResponseMPayload**](UserstagedGetListV1ResponseMPayload.md) |  | 

## Methods

### NewUserstagedGetListV1Response

`func NewUserstagedGetListV1Response(objDebugPayload CommonResponseObjDebugPayloadGetList, mPayload UserstagedGetListV1ResponseMPayload, ) *UserstagedGetListV1Response`

NewUserstagedGetListV1Response instantiates a new UserstagedGetListV1Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserstagedGetListV1ResponseWithDefaults

`func NewUserstagedGetListV1ResponseWithDefaults() *UserstagedGetListV1Response`

NewUserstagedGetListV1ResponseWithDefaults instantiates a new UserstagedGetListV1Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjDebugPayload

`func (o *UserstagedGetListV1Response) GetObjDebugPayload() CommonResponseObjDebugPayloadGetList`

GetObjDebugPayload returns the ObjDebugPayload field if non-nil, zero value otherwise.

### GetObjDebugPayloadOk

`func (o *UserstagedGetListV1Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayloadGetList, bool)`

GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjDebugPayload

`func (o *UserstagedGetListV1Response) SetObjDebugPayload(v CommonResponseObjDebugPayloadGetList)`

SetObjDebugPayload sets ObjDebugPayload field to given value.


### GetObjDebug

`func (o *UserstagedGetListV1Response) GetObjDebug() CommonResponseObjDebug`

GetObjDebug returns the ObjDebug field if non-nil, zero value otherwise.

### GetObjDebugOk

`func (o *UserstagedGetListV1Response) GetObjDebugOk() (*CommonResponseObjDebug, bool)`

GetObjDebugOk returns a tuple with the ObjDebug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjDebug

`func (o *UserstagedGetListV1Response) SetObjDebug(v CommonResponseObjDebug)`

SetObjDebug sets ObjDebug field to given value.

### HasObjDebug

`func (o *UserstagedGetListV1Response) HasObjDebug() bool`

HasObjDebug returns a boolean if a field has been set.

### GetMPayload

`func (o *UserstagedGetListV1Response) GetMPayload() UserstagedGetListV1ResponseMPayload`

GetMPayload returns the MPayload field if non-nil, zero value otherwise.

### GetMPayloadOk

`func (o *UserstagedGetListV1Response) GetMPayloadOk() (*UserstagedGetListV1ResponseMPayload, bool)`

GetMPayloadOk returns a tuple with the MPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMPayload

`func (o *UserstagedGetListV1Response) SetMPayload(v UserstagedGetListV1ResponseMPayload)`

SetMPayload sets MPayload field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


