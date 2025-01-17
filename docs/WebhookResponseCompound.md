# WebhookResponseCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SWebhookEvent** | Pointer to **string** | The concatenated string to describe the Webhook event | [optional] 
**AObjWebhookheader** | Pointer to [**[]WebhookheaderResponseCompound**](WebhookheaderResponseCompound.md) |  | [optional] 

## Methods

### NewWebhookResponseCompound

`func NewWebhookResponseCompound() *WebhookResponseCompound`

NewWebhookResponseCompound instantiates a new WebhookResponseCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookResponseCompoundWithDefaults

`func NewWebhookResponseCompoundWithDefaults() *WebhookResponseCompound`

NewWebhookResponseCompoundWithDefaults instantiates a new WebhookResponseCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSWebhookEvent

`func (o *WebhookResponseCompound) GetSWebhookEvent() string`

GetSWebhookEvent returns the SWebhookEvent field if non-nil, zero value otherwise.

### GetSWebhookEventOk

`func (o *WebhookResponseCompound) GetSWebhookEventOk() (*string, bool)`

GetSWebhookEventOk returns a tuple with the SWebhookEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSWebhookEvent

`func (o *WebhookResponseCompound) SetSWebhookEvent(v string)`

SetSWebhookEvent sets SWebhookEvent field to given value.

### HasSWebhookEvent

`func (o *WebhookResponseCompound) HasSWebhookEvent() bool`

HasSWebhookEvent returns a boolean if a field has been set.

### GetAObjWebhookheader

`func (o *WebhookResponseCompound) GetAObjWebhookheader() []WebhookheaderResponseCompound`

GetAObjWebhookheader returns the AObjWebhookheader field if non-nil, zero value otherwise.

### GetAObjWebhookheaderOk

`func (o *WebhookResponseCompound) GetAObjWebhookheaderOk() (*[]WebhookheaderResponseCompound, bool)`

GetAObjWebhookheaderOk returns a tuple with the AObjWebhookheader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjWebhookheader

`func (o *WebhookResponseCompound) SetAObjWebhookheader(v []WebhookheaderResponseCompound)`

SetAObjWebhookheader sets AObjWebhookheader field to given value.

### HasAObjWebhookheader

`func (o *WebhookResponseCompound) HasAObjWebhookheader() bool`

HasAObjWebhookheader returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


