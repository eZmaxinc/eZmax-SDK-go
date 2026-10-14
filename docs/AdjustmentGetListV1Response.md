# AdjustmentGetListV1Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjDebugPayload** | [**CommonResponseObjDebugPayloadGetList**](CommonResponseObjDebugPayloadGetList.md) |  | 
**ObjDebug** | Pointer to [**CommonResponseObjDebug**](CommonResponseObjDebug.md) |  | [optional] 
**MPayload** | [**AdjustmentGetListV1ResponseMPayload**](AdjustmentGetListV1ResponseMPayload.md) |  | 

## Methods

### NewAdjustmentGetListV1Response

`func NewAdjustmentGetListV1Response(objDebugPayload CommonResponseObjDebugPayloadGetList, mPayload AdjustmentGetListV1ResponseMPayload, ) *AdjustmentGetListV1Response`

NewAdjustmentGetListV1Response instantiates a new AdjustmentGetListV1Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdjustmentGetListV1ResponseWithDefaults

`func NewAdjustmentGetListV1ResponseWithDefaults() *AdjustmentGetListV1Response`

NewAdjustmentGetListV1ResponseWithDefaults instantiates a new AdjustmentGetListV1Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjDebugPayload

`func (o *AdjustmentGetListV1Response) GetObjDebugPayload() CommonResponseObjDebugPayloadGetList`

GetObjDebugPayload returns the ObjDebugPayload field if non-nil, zero value otherwise.

### GetObjDebugPayloadOk

`func (o *AdjustmentGetListV1Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayloadGetList, bool)`

GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjDebugPayload

`func (o *AdjustmentGetListV1Response) SetObjDebugPayload(v CommonResponseObjDebugPayloadGetList)`

SetObjDebugPayload sets ObjDebugPayload field to given value.


### GetObjDebug

`func (o *AdjustmentGetListV1Response) GetObjDebug() CommonResponseObjDebug`

GetObjDebug returns the ObjDebug field if non-nil, zero value otherwise.

### GetObjDebugOk

`func (o *AdjustmentGetListV1Response) GetObjDebugOk() (*CommonResponseObjDebug, bool)`

GetObjDebugOk returns a tuple with the ObjDebug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjDebug

`func (o *AdjustmentGetListV1Response) SetObjDebug(v CommonResponseObjDebug)`

SetObjDebug sets ObjDebug field to given value.

### HasObjDebug

`func (o *AdjustmentGetListV1Response) HasObjDebug() bool`

HasObjDebug returns a boolean if a field has been set.

### GetMPayload

`func (o *AdjustmentGetListV1Response) GetMPayload() AdjustmentGetListV1ResponseMPayload`

GetMPayload returns the MPayload field if non-nil, zero value otherwise.

### GetMPayloadOk

`func (o *AdjustmentGetListV1Response) GetMPayloadOk() (*AdjustmentGetListV1ResponseMPayload, bool)`

GetMPayloadOk returns a tuple with the MPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMPayload

`func (o *AdjustmentGetListV1Response) SetMPayload(v AdjustmentGetListV1ResponseMPayload)`

SetMPayload sets MPayload field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


