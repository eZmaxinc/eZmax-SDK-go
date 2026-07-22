# WebhookRealestateInscriptionModified

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjWebhook** | [**CustomWebhookResponse**](CustomWebhookResponse.md) |  | 
**AObjAttempt** | [**[]AttemptResponseCompound**](AttemptResponseCompound.md) | An array containing details of previous attempts that were made to deliver the message. The array is empty if it&#39;s the first attempt. | 
**ObjInscription** | [**InscriptionResponse**](InscriptionResponse.md) |  | 

## Methods

### NewWebhookRealestateInscriptionModified

`func NewWebhookRealestateInscriptionModified(objWebhook CustomWebhookResponse, aObjAttempt []AttemptResponseCompound, objInscription InscriptionResponse, ) *WebhookRealestateInscriptionModified`

NewWebhookRealestateInscriptionModified instantiates a new WebhookRealestateInscriptionModified object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookRealestateInscriptionModifiedWithDefaults

`func NewWebhookRealestateInscriptionModifiedWithDefaults() *WebhookRealestateInscriptionModified`

NewWebhookRealestateInscriptionModifiedWithDefaults instantiates a new WebhookRealestateInscriptionModified object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjWebhook

`func (o *WebhookRealestateInscriptionModified) GetObjWebhook() CustomWebhookResponse`

GetObjWebhook returns the ObjWebhook field if non-nil, zero value otherwise.

### GetObjWebhookOk

`func (o *WebhookRealestateInscriptionModified) GetObjWebhookOk() (*CustomWebhookResponse, bool)`

GetObjWebhookOk returns a tuple with the ObjWebhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjWebhook

`func (o *WebhookRealestateInscriptionModified) SetObjWebhook(v CustomWebhookResponse)`

SetObjWebhook sets ObjWebhook field to given value.


### GetAObjAttempt

`func (o *WebhookRealestateInscriptionModified) GetAObjAttempt() []AttemptResponseCompound`

GetAObjAttempt returns the AObjAttempt field if non-nil, zero value otherwise.

### GetAObjAttemptOk

`func (o *WebhookRealestateInscriptionModified) GetAObjAttemptOk() (*[]AttemptResponseCompound, bool)`

GetAObjAttemptOk returns a tuple with the AObjAttempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjAttempt

`func (o *WebhookRealestateInscriptionModified) SetAObjAttempt(v []AttemptResponseCompound)`

SetAObjAttempt sets AObjAttempt field to given value.


### GetObjInscription

`func (o *WebhookRealestateInscriptionModified) GetObjInscription() InscriptionResponse`

GetObjInscription returns the ObjInscription field if non-nil, zero value otherwise.

### GetObjInscriptionOk

`func (o *WebhookRealestateInscriptionModified) GetObjInscriptionOk() (*InscriptionResponse, bool)`

GetObjInscriptionOk returns a tuple with the ObjInscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjInscription

`func (o *WebhookRealestateInscriptionModified) SetObjInscription(v InscriptionResponse)`

SetObjInscription sets ObjInscription field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


