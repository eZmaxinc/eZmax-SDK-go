# WebhookGetListV1ResponseMPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IRowReturned** | **int32** | The number of rows returned | 
**IRowFiltered** | **int32** | The number of rows matching your filters (if any) or the total number of rows | 
**AObjWebhook** | [**[]WebhookListElement**](WebhookListElement.md) |  | 

## Methods

### NewWebhookGetListV1ResponseMPayload

`func NewWebhookGetListV1ResponseMPayload(iRowReturned int32, iRowFiltered int32, aObjWebhook []WebhookListElement, ) *WebhookGetListV1ResponseMPayload`

NewWebhookGetListV1ResponseMPayload instantiates a new WebhookGetListV1ResponseMPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookGetListV1ResponseMPayloadWithDefaults

`func NewWebhookGetListV1ResponseMPayloadWithDefaults() *WebhookGetListV1ResponseMPayload`

NewWebhookGetListV1ResponseMPayloadWithDefaults instantiates a new WebhookGetListV1ResponseMPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIRowReturned

`func (o *WebhookGetListV1ResponseMPayload) GetIRowReturned() int32`

GetIRowReturned returns the IRowReturned field if non-nil, zero value otherwise.

### GetIRowReturnedOk

`func (o *WebhookGetListV1ResponseMPayload) GetIRowReturnedOk() (*int32, bool)`

GetIRowReturnedOk returns a tuple with the IRowReturned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIRowReturned

`func (o *WebhookGetListV1ResponseMPayload) SetIRowReturned(v int32)`

SetIRowReturned sets IRowReturned field to given value.


### GetIRowFiltered

`func (o *WebhookGetListV1ResponseMPayload) GetIRowFiltered() int32`

GetIRowFiltered returns the IRowFiltered field if non-nil, zero value otherwise.

### GetIRowFilteredOk

`func (o *WebhookGetListV1ResponseMPayload) GetIRowFilteredOk() (*int32, bool)`

GetIRowFilteredOk returns a tuple with the IRowFiltered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIRowFiltered

`func (o *WebhookGetListV1ResponseMPayload) SetIRowFiltered(v int32)`

SetIRowFiltered sets IRowFiltered field to given value.


### GetAObjWebhook

`func (o *WebhookGetListV1ResponseMPayload) GetAObjWebhook() []WebhookListElement`

GetAObjWebhook returns the AObjWebhook field if non-nil, zero value otherwise.

### GetAObjWebhookOk

`func (o *WebhookGetListV1ResponseMPayload) GetAObjWebhookOk() (*[]WebhookListElement, bool)`

GetAObjWebhookOk returns a tuple with the AObjWebhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjWebhook

`func (o *WebhookGetListV1ResponseMPayload) SetAObjWebhook(v []WebhookListElement)`

SetAObjWebhook sets AObjWebhook field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


