# WebhookUserUserCreated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjUser** | [**UserResponse**](user-Response.md) |  | 
**ObjWebhook** | [**WebhookResponse**](webhook-Response.md) |  | 
**AObjAttempt** | [**[]AttemptResponse**](AttemptResponse.md) | An array containing details of previous attempts that were made to deliver the message. The array is empty if it&#39;s the first attempt. | 

## Methods

### NewWebhookUserUserCreated

`func NewWebhookUserUserCreated(objUser UserResponse, objWebhook WebhookResponse, aObjAttempt []AttemptResponse, ) *WebhookUserUserCreated`

NewWebhookUserUserCreated instantiates a new WebhookUserUserCreated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookUserUserCreatedWithDefaults

`func NewWebhookUserUserCreatedWithDefaults() *WebhookUserUserCreated`

NewWebhookUserUserCreatedWithDefaults instantiates a new WebhookUserUserCreated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjUser

`func (o *WebhookUserUserCreated) GetObjUser() UserResponse`

GetObjUser returns the ObjUser field if non-nil, zero value otherwise.

### GetObjUserOk

`func (o *WebhookUserUserCreated) GetObjUserOk() (*UserResponse, bool)`

GetObjUserOk returns a tuple with the ObjUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjUser

`func (o *WebhookUserUserCreated) SetObjUser(v UserResponse)`

SetObjUser sets ObjUser field to given value.


### GetObjWebhook

`func (o *WebhookUserUserCreated) GetObjWebhook() WebhookResponse`

GetObjWebhook returns the ObjWebhook field if non-nil, zero value otherwise.

### GetObjWebhookOk

`func (o *WebhookUserUserCreated) GetObjWebhookOk() (*WebhookResponse, bool)`

GetObjWebhookOk returns a tuple with the ObjWebhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjWebhook

`func (o *WebhookUserUserCreated) SetObjWebhook(v WebhookResponse)`

SetObjWebhook sets ObjWebhook field to given value.


### GetAObjAttempt

`func (o *WebhookUserUserCreated) GetAObjAttempt() []AttemptResponse`

GetAObjAttempt returns the AObjAttempt field if non-nil, zero value otherwise.

### GetAObjAttemptOk

`func (o *WebhookUserUserCreated) GetAObjAttemptOk() (*[]AttemptResponse, bool)`

GetAObjAttemptOk returns a tuple with the AObjAttempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjAttempt

`func (o *WebhookUserUserCreated) SetAObjAttempt(v []AttemptResponse)`

SetAObjAttempt sets AObjAttempt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


