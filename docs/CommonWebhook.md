# CommonWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjWebhook** | [**WebhookResponse**](webhook-Response.md) |  | 
**AObjAttempt** | [**[]AttemptResponse**](AttemptResponse.md) | An array containing details of previous attempts that were made to deliver the message. The array is empty if it&#39;s the first attempt. | 

## Methods

### NewCommonWebhook

`func NewCommonWebhook(objWebhook WebhookResponse, aObjAttempt []AttemptResponse, ) *CommonWebhook`

NewCommonWebhook instantiates a new CommonWebhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonWebhookWithDefaults

`func NewCommonWebhookWithDefaults() *CommonWebhook`

NewCommonWebhookWithDefaults instantiates a new CommonWebhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjWebhook

`func (o *CommonWebhook) GetObjWebhook() WebhookResponse`

GetObjWebhook returns the ObjWebhook field if non-nil, zero value otherwise.

### GetObjWebhookOk

`func (o *CommonWebhook) GetObjWebhookOk() (*WebhookResponse, bool)`

GetObjWebhookOk returns a tuple with the ObjWebhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjWebhook

`func (o *CommonWebhook) SetObjWebhook(v WebhookResponse)`

SetObjWebhook sets ObjWebhook field to given value.


### GetAObjAttempt

`func (o *CommonWebhook) GetAObjAttempt() []AttemptResponse`

GetAObjAttempt returns the AObjAttempt field if non-nil, zero value otherwise.

### GetAObjAttemptOk

`func (o *CommonWebhook) GetAObjAttemptOk() (*[]AttemptResponse, bool)`

GetAObjAttemptOk returns a tuple with the AObjAttempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjAttempt

`func (o *CommonWebhook) SetAObjAttempt(v []AttemptResponse)`

SetAObjAttempt sets AObjAttempt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


