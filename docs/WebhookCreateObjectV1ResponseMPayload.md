# WebhookCreateObjectV1ResponseMPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**APkiWebhookID** | **[]int32** | An array of unique IDs representing the object that were requested to be created.  They are returned in the same order as the array containing the objects to be created that was sent in the request. | 

## Methods

### NewWebhookCreateObjectV1ResponseMPayload

`func NewWebhookCreateObjectV1ResponseMPayload(aPkiWebhookID []int32, ) *WebhookCreateObjectV1ResponseMPayload`

NewWebhookCreateObjectV1ResponseMPayload instantiates a new WebhookCreateObjectV1ResponseMPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookCreateObjectV1ResponseMPayloadWithDefaults

`func NewWebhookCreateObjectV1ResponseMPayloadWithDefaults() *WebhookCreateObjectV1ResponseMPayload`

NewWebhookCreateObjectV1ResponseMPayloadWithDefaults instantiates a new WebhookCreateObjectV1ResponseMPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAPkiWebhookID

`func (o *WebhookCreateObjectV1ResponseMPayload) GetAPkiWebhookID() []int32`

GetAPkiWebhookID returns the APkiWebhookID field if non-nil, zero value otherwise.

### GetAPkiWebhookIDOk

`func (o *WebhookCreateObjectV1ResponseMPayload) GetAPkiWebhookIDOk() (*[]int32, bool)`

GetAPkiWebhookIDOk returns a tuple with the APkiWebhookID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAPkiWebhookID

`func (o *WebhookCreateObjectV1ResponseMPayload) SetAPkiWebhookID(v []int32)`

SetAPkiWebhookID sets APkiWebhookID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


