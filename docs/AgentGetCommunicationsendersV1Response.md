# AgentGetCommunicationsendersV1Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjDebugPayload** | [**CommonResponseObjDebugPayload**](CommonResponseObjDebugPayload.md) |  | 
**ObjDebug** | Pointer to [**CommonResponseObjDebug**](CommonResponseObjDebug.md) |  | [optional] 
**MPayload** | [**AgentGetCommunicationsendersV1ResponseMPayload**](AgentGetCommunicationsendersV1ResponseMPayload.md) |  | 

## Methods

### NewAgentGetCommunicationsendersV1Response

`func NewAgentGetCommunicationsendersV1Response(objDebugPayload CommonResponseObjDebugPayload, mPayload AgentGetCommunicationsendersV1ResponseMPayload, ) *AgentGetCommunicationsendersV1Response`

NewAgentGetCommunicationsendersV1Response instantiates a new AgentGetCommunicationsendersV1Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentGetCommunicationsendersV1ResponseWithDefaults

`func NewAgentGetCommunicationsendersV1ResponseWithDefaults() *AgentGetCommunicationsendersV1Response`

NewAgentGetCommunicationsendersV1ResponseWithDefaults instantiates a new AgentGetCommunicationsendersV1Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjDebugPayload

`func (o *AgentGetCommunicationsendersV1Response) GetObjDebugPayload() CommonResponseObjDebugPayload`

GetObjDebugPayload returns the ObjDebugPayload field if non-nil, zero value otherwise.

### GetObjDebugPayloadOk

`func (o *AgentGetCommunicationsendersV1Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayload, bool)`

GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjDebugPayload

`func (o *AgentGetCommunicationsendersV1Response) SetObjDebugPayload(v CommonResponseObjDebugPayload)`

SetObjDebugPayload sets ObjDebugPayload field to given value.


### GetObjDebug

`func (o *AgentGetCommunicationsendersV1Response) GetObjDebug() CommonResponseObjDebug`

GetObjDebug returns the ObjDebug field if non-nil, zero value otherwise.

### GetObjDebugOk

`func (o *AgentGetCommunicationsendersV1Response) GetObjDebugOk() (*CommonResponseObjDebug, bool)`

GetObjDebugOk returns a tuple with the ObjDebug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjDebug

`func (o *AgentGetCommunicationsendersV1Response) SetObjDebug(v CommonResponseObjDebug)`

SetObjDebug sets ObjDebug field to given value.

### HasObjDebug

`func (o *AgentGetCommunicationsendersV1Response) HasObjDebug() bool`

HasObjDebug returns a boolean if a field has been set.

### GetMPayload

`func (o *AgentGetCommunicationsendersV1Response) GetMPayload() AgentGetCommunicationsendersV1ResponseMPayload`

GetMPayload returns the MPayload field if non-nil, zero value otherwise.

### GetMPayloadOk

`func (o *AgentGetCommunicationsendersV1Response) GetMPayloadOk() (*AgentGetCommunicationsendersV1ResponseMPayload, bool)`

GetMPayloadOk returns a tuple with the MPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMPayload

`func (o *AgentGetCommunicationsendersV1Response) SetMPayload(v AgentGetCommunicationsendersV1ResponseMPayload)`

SetMPayload sets MPayload field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


