# AgentGetCommunicationrecipientsV1Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjDebugPayload** | [**CommonResponseObjDebugPayload**](CommonResponseObjDebugPayload.md) |  | 
**ObjDebug** | Pointer to [**CommonResponseObjDebug**](CommonResponseObjDebug.md) |  | [optional] 
**MPayload** | [**AgentGetCommunicationrecipientsV1ResponseMPayload**](AgentGetCommunicationrecipientsV1ResponseMPayload.md) |  | 

## Methods

### NewAgentGetCommunicationrecipientsV1Response

`func NewAgentGetCommunicationrecipientsV1Response(objDebugPayload CommonResponseObjDebugPayload, mPayload AgentGetCommunicationrecipientsV1ResponseMPayload, ) *AgentGetCommunicationrecipientsV1Response`

NewAgentGetCommunicationrecipientsV1Response instantiates a new AgentGetCommunicationrecipientsV1Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentGetCommunicationrecipientsV1ResponseWithDefaults

`func NewAgentGetCommunicationrecipientsV1ResponseWithDefaults() *AgentGetCommunicationrecipientsV1Response`

NewAgentGetCommunicationrecipientsV1ResponseWithDefaults instantiates a new AgentGetCommunicationrecipientsV1Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjDebugPayload

`func (o *AgentGetCommunicationrecipientsV1Response) GetObjDebugPayload() CommonResponseObjDebugPayload`

GetObjDebugPayload returns the ObjDebugPayload field if non-nil, zero value otherwise.

### GetObjDebugPayloadOk

`func (o *AgentGetCommunicationrecipientsV1Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayload, bool)`

GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjDebugPayload

`func (o *AgentGetCommunicationrecipientsV1Response) SetObjDebugPayload(v CommonResponseObjDebugPayload)`

SetObjDebugPayload sets ObjDebugPayload field to given value.


### GetObjDebug

`func (o *AgentGetCommunicationrecipientsV1Response) GetObjDebug() CommonResponseObjDebug`

GetObjDebug returns the ObjDebug field if non-nil, zero value otherwise.

### GetObjDebugOk

`func (o *AgentGetCommunicationrecipientsV1Response) GetObjDebugOk() (*CommonResponseObjDebug, bool)`

GetObjDebugOk returns a tuple with the ObjDebug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjDebug

`func (o *AgentGetCommunicationrecipientsV1Response) SetObjDebug(v CommonResponseObjDebug)`

SetObjDebug sets ObjDebug field to given value.

### HasObjDebug

`func (o *AgentGetCommunicationrecipientsV1Response) HasObjDebug() bool`

HasObjDebug returns a boolean if a field has been set.

### GetMPayload

`func (o *AgentGetCommunicationrecipientsV1Response) GetMPayload() AgentGetCommunicationrecipientsV1ResponseMPayload`

GetMPayload returns the MPayload field if non-nil, zero value otherwise.

### GetMPayloadOk

`func (o *AgentGetCommunicationrecipientsV1Response) GetMPayloadOk() (*AgentGetCommunicationrecipientsV1ResponseMPayload, bool)`

GetMPayloadOk returns a tuple with the MPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMPayload

`func (o *AgentGetCommunicationrecipientsV1Response) SetMPayload(v AgentGetCommunicationrecipientsV1ResponseMPayload)`

SetMPayload sets MPayload field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


