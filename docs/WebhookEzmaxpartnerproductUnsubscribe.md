# WebhookEzmaxpartnerproductUnsubscribe

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjWebhook** | [**CustomWebhookResponse**](CustomWebhookResponse.md) |  | 
**AObjAttempt** | [**[]AttemptResponseCompound**](AttemptResponseCompound.md) | An array containing details of previous attempts that were made to deliver the message. The array is empty if it&#39;s the first attempt. | 
**ObjEzmaxpartnerproduct** | [**CustomEzmaxpartnerproductSubscribe**](CustomEzmaxpartnerproductSubscribe.md) |  | 
**SExternalID** | Pointer to **string** |  | [optional] 

## Methods

### NewWebhookEzmaxpartnerproductUnsubscribe

`func NewWebhookEzmaxpartnerproductUnsubscribe(objWebhook CustomWebhookResponse, aObjAttempt []AttemptResponseCompound, objEzmaxpartnerproduct CustomEzmaxpartnerproductSubscribe, ) *WebhookEzmaxpartnerproductUnsubscribe`

NewWebhookEzmaxpartnerproductUnsubscribe instantiates a new WebhookEzmaxpartnerproductUnsubscribe object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookEzmaxpartnerproductUnsubscribeWithDefaults

`func NewWebhookEzmaxpartnerproductUnsubscribeWithDefaults() *WebhookEzmaxpartnerproductUnsubscribe`

NewWebhookEzmaxpartnerproductUnsubscribeWithDefaults instantiates a new WebhookEzmaxpartnerproductUnsubscribe object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjWebhook

`func (o *WebhookEzmaxpartnerproductUnsubscribe) GetObjWebhook() CustomWebhookResponse`

GetObjWebhook returns the ObjWebhook field if non-nil, zero value otherwise.

### GetObjWebhookOk

`func (o *WebhookEzmaxpartnerproductUnsubscribe) GetObjWebhookOk() (*CustomWebhookResponse, bool)`

GetObjWebhookOk returns a tuple with the ObjWebhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjWebhook

`func (o *WebhookEzmaxpartnerproductUnsubscribe) SetObjWebhook(v CustomWebhookResponse)`

SetObjWebhook sets ObjWebhook field to given value.


### GetAObjAttempt

`func (o *WebhookEzmaxpartnerproductUnsubscribe) GetAObjAttempt() []AttemptResponseCompound`

GetAObjAttempt returns the AObjAttempt field if non-nil, zero value otherwise.

### GetAObjAttemptOk

`func (o *WebhookEzmaxpartnerproductUnsubscribe) GetAObjAttemptOk() (*[]AttemptResponseCompound, bool)`

GetAObjAttemptOk returns a tuple with the AObjAttempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjAttempt

`func (o *WebhookEzmaxpartnerproductUnsubscribe) SetAObjAttempt(v []AttemptResponseCompound)`

SetAObjAttempt sets AObjAttempt field to given value.


### GetObjEzmaxpartnerproduct

`func (o *WebhookEzmaxpartnerproductUnsubscribe) GetObjEzmaxpartnerproduct() CustomEzmaxpartnerproductSubscribe`

GetObjEzmaxpartnerproduct returns the ObjEzmaxpartnerproduct field if non-nil, zero value otherwise.

### GetObjEzmaxpartnerproductOk

`func (o *WebhookEzmaxpartnerproductUnsubscribe) GetObjEzmaxpartnerproductOk() (*CustomEzmaxpartnerproductSubscribe, bool)`

GetObjEzmaxpartnerproductOk returns a tuple with the ObjEzmaxpartnerproduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjEzmaxpartnerproduct

`func (o *WebhookEzmaxpartnerproductUnsubscribe) SetObjEzmaxpartnerproduct(v CustomEzmaxpartnerproductSubscribe)`

SetObjEzmaxpartnerproduct sets ObjEzmaxpartnerproduct field to given value.


### GetSExternalID

`func (o *WebhookEzmaxpartnerproductUnsubscribe) GetSExternalID() string`

GetSExternalID returns the SExternalID field if non-nil, zero value otherwise.

### GetSExternalIDOk

`func (o *WebhookEzmaxpartnerproductUnsubscribe) GetSExternalIDOk() (*string, bool)`

GetSExternalIDOk returns a tuple with the SExternalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSExternalID

`func (o *WebhookEzmaxpartnerproductUnsubscribe) SetSExternalID(v string)`

SetSExternalID sets SExternalID field to given value.

### HasSExternalID

`func (o *WebhookEzmaxpartnerproductUnsubscribe) HasSExternalID() bool`

HasSExternalID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


