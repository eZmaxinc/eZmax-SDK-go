# ListGetListpresentationV1Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MPayload** | [**ListGetListpresentationV1ResponseMPayload**](ListGetListpresentationV1ResponseMPayload.md) |  | 
**ObjDebugPayload** | Pointer to [**CommonResponseObjDebugPayload**](CommonResponseObjDebugPayload.md) |  | [optional] 
**ObjDebug** | Pointer to [**CommonResponseObjDebug**](CommonResponseObjDebug.md) |  | [optional] 

## Methods

### NewListGetListpresentationV1Response

`func NewListGetListpresentationV1Response(mPayload ListGetListpresentationV1ResponseMPayload, ) *ListGetListpresentationV1Response`

NewListGetListpresentationV1Response instantiates a new ListGetListpresentationV1Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListGetListpresentationV1ResponseWithDefaults

`func NewListGetListpresentationV1ResponseWithDefaults() *ListGetListpresentationV1Response`

NewListGetListpresentationV1ResponseWithDefaults instantiates a new ListGetListpresentationV1Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMPayload

`func (o *ListGetListpresentationV1Response) GetMPayload() ListGetListpresentationV1ResponseMPayload`

GetMPayload returns the MPayload field if non-nil, zero value otherwise.

### GetMPayloadOk

`func (o *ListGetListpresentationV1Response) GetMPayloadOk() (*ListGetListpresentationV1ResponseMPayload, bool)`

GetMPayloadOk returns a tuple with the MPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMPayload

`func (o *ListGetListpresentationV1Response) SetMPayload(v ListGetListpresentationV1ResponseMPayload)`

SetMPayload sets MPayload field to given value.


### GetObjDebugPayload

`func (o *ListGetListpresentationV1Response) GetObjDebugPayload() CommonResponseObjDebugPayload`

GetObjDebugPayload returns the ObjDebugPayload field if non-nil, zero value otherwise.

### GetObjDebugPayloadOk

`func (o *ListGetListpresentationV1Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayload, bool)`

GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjDebugPayload

`func (o *ListGetListpresentationV1Response) SetObjDebugPayload(v CommonResponseObjDebugPayload)`

SetObjDebugPayload sets ObjDebugPayload field to given value.

### HasObjDebugPayload

`func (o *ListGetListpresentationV1Response) HasObjDebugPayload() bool`

HasObjDebugPayload returns a boolean if a field has been set.

### GetObjDebug

`func (o *ListGetListpresentationV1Response) GetObjDebug() CommonResponseObjDebug`

GetObjDebug returns the ObjDebug field if non-nil, zero value otherwise.

### GetObjDebugOk

`func (o *ListGetListpresentationV1Response) GetObjDebugOk() (*CommonResponseObjDebug, bool)`

GetObjDebugOk returns a tuple with the ObjDebug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjDebug

`func (o *ListGetListpresentationV1Response) SetObjDebug(v CommonResponseObjDebug)`

SetObjDebug sets ObjDebug field to given value.

### HasObjDebug

`func (o *ListGetListpresentationV1Response) HasObjDebug() bool`

HasObjDebug returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


