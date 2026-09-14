# DepositGetCommunicationListV1Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjDebugPayload** | [**CommonResponseObjDebugPayloadGetList**](CommonResponseObjDebugPayloadGetList.md) |  | 
**ObjDebug** | Pointer to [**CommonResponseObjDebug**](CommonResponseObjDebug.md) |  | [optional] 
**MPayload** | [**DepositGetCommunicationListV1ResponseMPayload**](DepositGetCommunicationListV1ResponseMPayload.md) |  | 

## Methods

### NewDepositGetCommunicationListV1Response

`func NewDepositGetCommunicationListV1Response(objDebugPayload CommonResponseObjDebugPayloadGetList, mPayload DepositGetCommunicationListV1ResponseMPayload, ) *DepositGetCommunicationListV1Response`

NewDepositGetCommunicationListV1Response instantiates a new DepositGetCommunicationListV1Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDepositGetCommunicationListV1ResponseWithDefaults

`func NewDepositGetCommunicationListV1ResponseWithDefaults() *DepositGetCommunicationListV1Response`

NewDepositGetCommunicationListV1ResponseWithDefaults instantiates a new DepositGetCommunicationListV1Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjDebugPayload

`func (o *DepositGetCommunicationListV1Response) GetObjDebugPayload() CommonResponseObjDebugPayloadGetList`

GetObjDebugPayload returns the ObjDebugPayload field if non-nil, zero value otherwise.

### GetObjDebugPayloadOk

`func (o *DepositGetCommunicationListV1Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayloadGetList, bool)`

GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjDebugPayload

`func (o *DepositGetCommunicationListV1Response) SetObjDebugPayload(v CommonResponseObjDebugPayloadGetList)`

SetObjDebugPayload sets ObjDebugPayload field to given value.


### GetObjDebug

`func (o *DepositGetCommunicationListV1Response) GetObjDebug() CommonResponseObjDebug`

GetObjDebug returns the ObjDebug field if non-nil, zero value otherwise.

### GetObjDebugOk

`func (o *DepositGetCommunicationListV1Response) GetObjDebugOk() (*CommonResponseObjDebug, bool)`

GetObjDebugOk returns a tuple with the ObjDebug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjDebug

`func (o *DepositGetCommunicationListV1Response) SetObjDebug(v CommonResponseObjDebug)`

SetObjDebug sets ObjDebug field to given value.

### HasObjDebug

`func (o *DepositGetCommunicationListV1Response) HasObjDebug() bool`

HasObjDebug returns a boolean if a field has been set.

### GetMPayload

`func (o *DepositGetCommunicationListV1Response) GetMPayload() DepositGetCommunicationListV1ResponseMPayload`

GetMPayload returns the MPayload field if non-nil, zero value otherwise.

### GetMPayloadOk

`func (o *DepositGetCommunicationListV1Response) GetMPayloadOk() (*DepositGetCommunicationListV1ResponseMPayload, bool)`

GetMPayloadOk returns a tuple with the MPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMPayload

`func (o *DepositGetCommunicationListV1Response) SetMPayload(v DepositGetCommunicationListV1ResponseMPayload)`

SetMPayload sets MPayload field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


