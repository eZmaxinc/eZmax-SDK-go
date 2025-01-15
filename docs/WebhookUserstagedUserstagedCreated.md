# WebhookUserstagedUserstagedCreated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjWebhook** | [**CustomWebhookResponse**](CustomWebhookResponse.md) |  | 
**AObjAttempt** | [**[]AttemptResponse**](AttemptResponse.md) | An array containing details of previous attempts that were made to deliver the message. The array is empty if it&#39;s the first attempt. | 
**ObjUserstaged** | [**UserstagedResponse**](UserstagedResponse.md) | A Userstaged Object | 

## Methods

### NewWebhookUserstagedUserstagedCreated

`func NewWebhookUserstagedUserstagedCreated(objWebhook CustomWebhookResponse, aObjAttempt []AttemptResponseCompound, objUserstaged UserstagedResponse, ) *WebhookUserstagedUserstagedCreated`

NewWebhookUserstagedUserstagedCreated instantiates a new WebhookUserstagedUserstagedCreated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookUserstagedUserstagedCreatedWithDefaults

`func NewWebhookUserstagedUserstagedCreatedWithDefaults() *WebhookUserstagedUserstagedCreated`

NewWebhookUserstagedUserstagedCreatedWithDefaults instantiates a new WebhookUserstagedUserstagedCreated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjWebhook

`func (o *WebhookUserstagedUserstagedCreated) GetObjWebhook() CustomWebhookResponse`

GetObjWebhook returns the ObjWebhook field if non-nil, zero value otherwise.

### GetObjWebhookOk

`func (o *WebhookUserstagedUserstagedCreated) GetObjWebhookOk() (*CustomWebhookResponse, bool)`

GetObjWebhookOk returns a tuple with the ObjWebhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjWebhook

`func (o *WebhookUserstagedUserstagedCreated) SetObjWebhook(v CustomWebhookResponse)`

SetObjWebhook sets ObjWebhook field to given value.


### GetAObjAttempt

`func (o *WebhookUserstagedUserstagedCreated) GetAObjAttempt() []AttemptResponseCompound`

GetAObjAttempt returns the AObjAttempt field if non-nil, zero value otherwise.

### GetAObjAttemptOk

`func (o *WebhookUserstagedUserstagedCreated) GetAObjAttemptOk() (*[]AttemptResponseCompound, bool)`

GetAObjAttemptOk returns a tuple with the AObjAttempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjAttempt

`func (o *WebhookUserstagedUserstagedCreated) SetAObjAttempt(v []AttemptResponseCompound)`

SetAObjAttempt sets AObjAttempt field to given value.


### GetObjUserstaged

`func (o *WebhookUserstagedUserstagedCreated) GetObjUserstaged() UserstagedResponse`

GetObjUserstaged returns the ObjUserstaged field if non-nil, zero value otherwise.

### GetObjUserstagedOk

`func (o *WebhookUserstagedUserstagedCreated) GetObjUserstagedOk() (*UserstagedResponse, bool)`

GetObjUserstagedOk returns a tuple with the ObjUserstaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjUserstaged

`func (o *WebhookUserstagedUserstagedCreated) SetObjUserstaged(v UserstagedResponse)`

SetObjUserstaged sets ObjUserstaged field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


