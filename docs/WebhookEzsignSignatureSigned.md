# WebhookEzsignSignatureSigned

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjWebhook** | [**CustomWebhookResponse**](CustomWebhookResponse.md) |  | 
**AObjAttempt** | [**[]AttemptResponse**](AttemptResponse.md) | An array containing details of previous attempts that were made to deliver the message. The array is empty if it&#39;s the first attempt. | 
**ObjEzsignsignature** | [**EzsignsignatureResponse**](EzsignsignatureResponse.md) |  | 

## Methods

### NewWebhookEzsignSignatureSigned

`func NewWebhookEzsignSignatureSigned(objWebhook CustomWebhookResponse, aObjAttempt []AttemptResponseCompound, objEzsignsignature EzsignsignatureResponse, ) *WebhookEzsignSignatureSigned`

NewWebhookEzsignSignatureSigned instantiates a new WebhookEzsignSignatureSigned object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookEzsignSignatureSignedWithDefaults

`func NewWebhookEzsignSignatureSignedWithDefaults() *WebhookEzsignSignatureSigned`

NewWebhookEzsignSignatureSignedWithDefaults instantiates a new WebhookEzsignSignatureSigned object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjWebhook

`func (o *WebhookEzsignSignatureSigned) GetObjWebhook() CustomWebhookResponse`

GetObjWebhook returns the ObjWebhook field if non-nil, zero value otherwise.

### GetObjWebhookOk

`func (o *WebhookEzsignSignatureSigned) GetObjWebhookOk() (*CustomWebhookResponse, bool)`

GetObjWebhookOk returns a tuple with the ObjWebhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjWebhook

`func (o *WebhookEzsignSignatureSigned) SetObjWebhook(v CustomWebhookResponse)`

SetObjWebhook sets ObjWebhook field to given value.


### GetAObjAttempt

`func (o *WebhookEzsignSignatureSigned) GetAObjAttempt() []AttemptResponseCompound`

GetAObjAttempt returns the AObjAttempt field if non-nil, zero value otherwise.

### GetAObjAttemptOk

`func (o *WebhookEzsignSignatureSigned) GetAObjAttemptOk() (*[]AttemptResponseCompound, bool)`

GetAObjAttemptOk returns a tuple with the AObjAttempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjAttempt

`func (o *WebhookEzsignSignatureSigned) SetAObjAttempt(v []AttemptResponseCompound)`

SetAObjAttempt sets AObjAttempt field to given value.


### GetObjEzsignsignature

`func (o *WebhookEzsignSignatureSigned) GetObjEzsignsignature() EzsignsignatureResponse`

GetObjEzsignsignature returns the ObjEzsignsignature field if non-nil, zero value otherwise.

### GetObjEzsignsignatureOk

`func (o *WebhookEzsignSignatureSigned) GetObjEzsignsignatureOk() (*EzsignsignatureResponse, bool)`

GetObjEzsignsignatureOk returns a tuple with the ObjEzsignsignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjEzsignsignature

`func (o *WebhookEzsignSignatureSigned) SetObjEzsignsignature(v EzsignsignatureResponse)`

SetObjEzsignsignature sets ObjEzsignsignature field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


