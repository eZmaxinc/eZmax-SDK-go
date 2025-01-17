# WebhookGetObjectV2Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjDebugPayload** | [**CommonResponseObjDebugPayload**](CommonResponseObjDebugPayload.md) |  | 
**ObjDebug** | Pointer to [**CommonResponseObjDebug**](CommonResponseObjDebug.md) |  | [optional] 
**MPayload** | [**WebhookGetObjectV2ResponseMPayload**](WebhookGetObjectV2ResponseMPayload.md) |  | 

## Methods

### NewWebhookGetObjectV2Response

`func NewWebhookGetObjectV2Response(objDebugPayload CommonResponseObjDebugPayload, mPayload WebhookGetObjectV2ResponseMPayload, ) *WebhookGetObjectV2Response`

NewWebhookGetObjectV2Response instantiates a new WebhookGetObjectV2Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookGetObjectV2ResponseWithDefaults

`func NewWebhookGetObjectV2ResponseWithDefaults() *WebhookGetObjectV2Response`

NewWebhookGetObjectV2ResponseWithDefaults instantiates a new WebhookGetObjectV2Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjDebugPayload

`func (o *WebhookGetObjectV2Response) GetObjDebugPayload() CommonResponseObjDebugPayload`

GetObjDebugPayload returns the ObjDebugPayload field if non-nil, zero value otherwise.

### GetObjDebugPayloadOk

`func (o *WebhookGetObjectV2Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayload, bool)`

GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjDebugPayload

`func (o *WebhookGetObjectV2Response) SetObjDebugPayload(v CommonResponseObjDebugPayload)`

SetObjDebugPayload sets ObjDebugPayload field to given value.


### GetObjDebug

`func (o *WebhookGetObjectV2Response) GetObjDebug() CommonResponseObjDebug`

GetObjDebug returns the ObjDebug field if non-nil, zero value otherwise.

### GetObjDebugOk

`func (o *WebhookGetObjectV2Response) GetObjDebugOk() (*CommonResponseObjDebug, bool)`

GetObjDebugOk returns a tuple with the ObjDebug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjDebug

`func (o *WebhookGetObjectV2Response) SetObjDebug(v CommonResponseObjDebug)`

SetObjDebug sets ObjDebug field to given value.

### HasObjDebug

`func (o *WebhookGetObjectV2Response) HasObjDebug() bool`

HasObjDebug returns a boolean if a field has been set.

### GetMPayload

`func (o *WebhookGetObjectV2Response) GetMPayload() WebhookGetObjectV2ResponseMPayload`

GetMPayload returns the MPayload field if non-nil, zero value otherwise.

### GetMPayloadOk

`func (o *WebhookGetObjectV2Response) GetMPayloadOk() (*WebhookGetObjectV2ResponseMPayload, bool)`

GetMPayloadOk returns a tuple with the MPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMPayload

`func (o *WebhookGetObjectV2Response) SetMPayload(v WebhookGetObjectV2ResponseMPayload)`

SetMPayload sets MPayload field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


