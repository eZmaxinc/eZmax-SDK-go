# DiscussionGetObjectV2Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjDebugPayload** | [**CommonResponseObjDebugPayload**](CommonResponseObjDebugPayload.md) |  | 
**ObjDebug** | Pointer to [**CommonResponseObjDebug**](CommonResponseObjDebug.md) |  | [optional] 
**MPayload** | [**DiscussionGetObjectV2ResponseMPayload**](DiscussionGetObjectV2ResponseMPayload.md) |  | 

## Methods

### NewDiscussionGetObjectV2Response

`func NewDiscussionGetObjectV2Response(objDebugPayload CommonResponseObjDebugPayload, mPayload DiscussionGetObjectV2ResponseMPayload, ) *DiscussionGetObjectV2Response`

NewDiscussionGetObjectV2Response instantiates a new DiscussionGetObjectV2Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscussionGetObjectV2ResponseWithDefaults

`func NewDiscussionGetObjectV2ResponseWithDefaults() *DiscussionGetObjectV2Response`

NewDiscussionGetObjectV2ResponseWithDefaults instantiates a new DiscussionGetObjectV2Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjDebugPayload

`func (o *DiscussionGetObjectV2Response) GetObjDebugPayload() CommonResponseObjDebugPayload`

GetObjDebugPayload returns the ObjDebugPayload field if non-nil, zero value otherwise.

### GetObjDebugPayloadOk

`func (o *DiscussionGetObjectV2Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayload, bool)`

GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjDebugPayload

`func (o *DiscussionGetObjectV2Response) SetObjDebugPayload(v CommonResponseObjDebugPayload)`

SetObjDebugPayload sets ObjDebugPayload field to given value.


### GetObjDebug

`func (o *DiscussionGetObjectV2Response) GetObjDebug() CommonResponseObjDebug`

GetObjDebug returns the ObjDebug field if non-nil, zero value otherwise.

### GetObjDebugOk

`func (o *DiscussionGetObjectV2Response) GetObjDebugOk() (*CommonResponseObjDebug, bool)`

GetObjDebugOk returns a tuple with the ObjDebug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjDebug

`func (o *DiscussionGetObjectV2Response) SetObjDebug(v CommonResponseObjDebug)`

SetObjDebug sets ObjDebug field to given value.

### HasObjDebug

`func (o *DiscussionGetObjectV2Response) HasObjDebug() bool`

HasObjDebug returns a boolean if a field has been set.

### GetMPayload

`func (o *DiscussionGetObjectV2Response) GetMPayload() DiscussionGetObjectV2ResponseMPayload`

GetMPayload returns the MPayload field if non-nil, zero value otherwise.

### GetMPayloadOk

`func (o *DiscussionGetObjectV2Response) GetMPayloadOk() (*DiscussionGetObjectV2ResponseMPayload, bool)`

GetMPayloadOk returns a tuple with the MPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMPayload

`func (o *DiscussionGetObjectV2Response) SetMPayload(v DiscussionGetObjectV2ResponseMPayload)`

SetMPayload sets MPayload field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


