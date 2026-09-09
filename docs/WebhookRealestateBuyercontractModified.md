# WebhookRealestateBuyercontractModified

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjWebhook** | [**CustomWebhookResponse**](CustomWebhookResponse.md) |  | 
**AObjAttempt** | [**[]AttemptResponseCompound**](AttemptResponseCompound.md) | An array containing details of previous attempts that were made to deliver the message. The array is empty if it&#39;s the first attempt. | 
**ObjBuyercontract** | [**BuyercontractResponse**](BuyercontractResponse.md) |  | 

## Methods

### NewWebhookRealestateBuyercontractModified

`func NewWebhookRealestateBuyercontractModified(objWebhook CustomWebhookResponse, aObjAttempt []AttemptResponseCompound, objBuyercontract BuyercontractResponse, ) *WebhookRealestateBuyercontractModified`

NewWebhookRealestateBuyercontractModified instantiates a new WebhookRealestateBuyercontractModified object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookRealestateBuyercontractModifiedWithDefaults

`func NewWebhookRealestateBuyercontractModifiedWithDefaults() *WebhookRealestateBuyercontractModified`

NewWebhookRealestateBuyercontractModifiedWithDefaults instantiates a new WebhookRealestateBuyercontractModified object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjWebhook

`func (o *WebhookRealestateBuyercontractModified) GetObjWebhook() CustomWebhookResponse`

GetObjWebhook returns the ObjWebhook field if non-nil, zero value otherwise.

### GetObjWebhookOk

`func (o *WebhookRealestateBuyercontractModified) GetObjWebhookOk() (*CustomWebhookResponse, bool)`

GetObjWebhookOk returns a tuple with the ObjWebhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjWebhook

`func (o *WebhookRealestateBuyercontractModified) SetObjWebhook(v CustomWebhookResponse)`

SetObjWebhook sets ObjWebhook field to given value.


### GetAObjAttempt

`func (o *WebhookRealestateBuyercontractModified) GetAObjAttempt() []AttemptResponseCompound`

GetAObjAttempt returns the AObjAttempt field if non-nil, zero value otherwise.

### GetAObjAttemptOk

`func (o *WebhookRealestateBuyercontractModified) GetAObjAttemptOk() (*[]AttemptResponseCompound, bool)`

GetAObjAttemptOk returns a tuple with the AObjAttempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjAttempt

`func (o *WebhookRealestateBuyercontractModified) SetAObjAttempt(v []AttemptResponseCompound)`

SetAObjAttempt sets AObjAttempt field to given value.


### GetObjBuyercontract

`func (o *WebhookRealestateBuyercontractModified) GetObjBuyercontract() BuyercontractResponse`

GetObjBuyercontract returns the ObjBuyercontract field if non-nil, zero value otherwise.

### GetObjBuyercontractOk

`func (o *WebhookRealestateBuyercontractModified) GetObjBuyercontractOk() (*BuyercontractResponse, bool)`

GetObjBuyercontractOk returns a tuple with the ObjBuyercontract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjBuyercontract

`func (o *WebhookRealestateBuyercontractModified) SetObjBuyercontract(v BuyercontractResponse)`

SetObjBuyercontract sets ObjBuyercontract field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


