# WebhookEzmaxpartnerproductSubscribe

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjWebhook** | [**CustomWebhookResponse**](CustomWebhookResponse.md) |  | 
**AObjAttempt** | [**[]AttemptResponseCompound**](AttemptResponseCompound.md) | An array containing details of previous attempts that were made to deliver the message. The array is empty if it&#39;s the first attempt. | 
**ObjEzmaxpartnerproduct** | [**CustomEzmaxpartnerproductSubscribe**](CustomEzmaxpartnerproductSubscribe.md) |  | 
**SExternalID** | Pointer to **string** |  | [optional] 
**SApikeyApikey** | Pointer to **string** |  | [optional] 
**SApikeySecret** | Pointer to **string** |  | [optional] 

## Methods

### NewWebhookEzmaxpartnerproductSubscribe

`func NewWebhookEzmaxpartnerproductSubscribe(objWebhook CustomWebhookResponse, aObjAttempt []AttemptResponseCompound, objEzmaxpartnerproduct CustomEzmaxpartnerproductSubscribe, ) *WebhookEzmaxpartnerproductSubscribe`

NewWebhookEzmaxpartnerproductSubscribe instantiates a new WebhookEzmaxpartnerproductSubscribe object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookEzmaxpartnerproductSubscribeWithDefaults

`func NewWebhookEzmaxpartnerproductSubscribeWithDefaults() *WebhookEzmaxpartnerproductSubscribe`

NewWebhookEzmaxpartnerproductSubscribeWithDefaults instantiates a new WebhookEzmaxpartnerproductSubscribe object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjWebhook

`func (o *WebhookEzmaxpartnerproductSubscribe) GetObjWebhook() CustomWebhookResponse`

GetObjWebhook returns the ObjWebhook field if non-nil, zero value otherwise.

### GetObjWebhookOk

`func (o *WebhookEzmaxpartnerproductSubscribe) GetObjWebhookOk() (*CustomWebhookResponse, bool)`

GetObjWebhookOk returns a tuple with the ObjWebhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjWebhook

`func (o *WebhookEzmaxpartnerproductSubscribe) SetObjWebhook(v CustomWebhookResponse)`

SetObjWebhook sets ObjWebhook field to given value.


### GetAObjAttempt

`func (o *WebhookEzmaxpartnerproductSubscribe) GetAObjAttempt() []AttemptResponseCompound`

GetAObjAttempt returns the AObjAttempt field if non-nil, zero value otherwise.

### GetAObjAttemptOk

`func (o *WebhookEzmaxpartnerproductSubscribe) GetAObjAttemptOk() (*[]AttemptResponseCompound, bool)`

GetAObjAttemptOk returns a tuple with the AObjAttempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjAttempt

`func (o *WebhookEzmaxpartnerproductSubscribe) SetAObjAttempt(v []AttemptResponseCompound)`

SetAObjAttempt sets AObjAttempt field to given value.


### GetObjEzmaxpartnerproduct

`func (o *WebhookEzmaxpartnerproductSubscribe) GetObjEzmaxpartnerproduct() CustomEzmaxpartnerproductSubscribe`

GetObjEzmaxpartnerproduct returns the ObjEzmaxpartnerproduct field if non-nil, zero value otherwise.

### GetObjEzmaxpartnerproductOk

`func (o *WebhookEzmaxpartnerproductSubscribe) GetObjEzmaxpartnerproductOk() (*CustomEzmaxpartnerproductSubscribe, bool)`

GetObjEzmaxpartnerproductOk returns a tuple with the ObjEzmaxpartnerproduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjEzmaxpartnerproduct

`func (o *WebhookEzmaxpartnerproductSubscribe) SetObjEzmaxpartnerproduct(v CustomEzmaxpartnerproductSubscribe)`

SetObjEzmaxpartnerproduct sets ObjEzmaxpartnerproduct field to given value.


### GetSExternalID

`func (o *WebhookEzmaxpartnerproductSubscribe) GetSExternalID() string`

GetSExternalID returns the SExternalID field if non-nil, zero value otherwise.

### GetSExternalIDOk

`func (o *WebhookEzmaxpartnerproductSubscribe) GetSExternalIDOk() (*string, bool)`

GetSExternalIDOk returns a tuple with the SExternalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSExternalID

`func (o *WebhookEzmaxpartnerproductSubscribe) SetSExternalID(v string)`

SetSExternalID sets SExternalID field to given value.

### HasSExternalID

`func (o *WebhookEzmaxpartnerproductSubscribe) HasSExternalID() bool`

HasSExternalID returns a boolean if a field has been set.

### GetSApikeyApikey

`func (o *WebhookEzmaxpartnerproductSubscribe) GetSApikeyApikey() string`

GetSApikeyApikey returns the SApikeyApikey field if non-nil, zero value otherwise.

### GetSApikeyApikeyOk

`func (o *WebhookEzmaxpartnerproductSubscribe) GetSApikeyApikeyOk() (*string, bool)`

GetSApikeyApikeyOk returns a tuple with the SApikeyApikey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSApikeyApikey

`func (o *WebhookEzmaxpartnerproductSubscribe) SetSApikeyApikey(v string)`

SetSApikeyApikey sets SApikeyApikey field to given value.

### HasSApikeyApikey

`func (o *WebhookEzmaxpartnerproductSubscribe) HasSApikeyApikey() bool`

HasSApikeyApikey returns a boolean if a field has been set.

### GetSApikeySecret

`func (o *WebhookEzmaxpartnerproductSubscribe) GetSApikeySecret() string`

GetSApikeySecret returns the SApikeySecret field if non-nil, zero value otherwise.

### GetSApikeySecretOk

`func (o *WebhookEzmaxpartnerproductSubscribe) GetSApikeySecretOk() (*string, bool)`

GetSApikeySecretOk returns a tuple with the SApikeySecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSApikeySecret

`func (o *WebhookEzmaxpartnerproductSubscribe) SetSApikeySecret(v string)`

SetSApikeySecret sets SApikeySecret field to given value.

### HasSApikeySecret

`func (o *WebhookEzmaxpartnerproductSubscribe) HasSApikeySecret() bool`

HasSApikeySecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


