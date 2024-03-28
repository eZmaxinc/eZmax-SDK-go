# WebhookheaderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiWebhookheaderID** | Pointer to **int32** | The unique ID of the Webhookheader | [optional] 
**SWebhookheaderName** | **string** | The Name of the Webhookheader | 
**SWebhookheaderValue** | **string** | The Value of the Webhookheader | 

## Methods

### NewWebhookheaderRequest

`func NewWebhookheaderRequest(sWebhookheaderName string, sWebhookheaderValue string, ) *WebhookheaderRequest`

NewWebhookheaderRequest instantiates a new WebhookheaderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookheaderRequestWithDefaults

`func NewWebhookheaderRequestWithDefaults() *WebhookheaderRequest`

NewWebhookheaderRequestWithDefaults instantiates a new WebhookheaderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiWebhookheaderID

`func (o *WebhookheaderRequest) GetPkiWebhookheaderID() int32`

GetPkiWebhookheaderID returns the PkiWebhookheaderID field if non-nil, zero value otherwise.

### GetPkiWebhookheaderIDOk

`func (o *WebhookheaderRequest) GetPkiWebhookheaderIDOk() (*int32, bool)`

GetPkiWebhookheaderIDOk returns a tuple with the PkiWebhookheaderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiWebhookheaderID

`func (o *WebhookheaderRequest) SetPkiWebhookheaderID(v int32)`

SetPkiWebhookheaderID sets PkiWebhookheaderID field to given value.

### HasPkiWebhookheaderID

`func (o *WebhookheaderRequest) HasPkiWebhookheaderID() bool`

HasPkiWebhookheaderID returns a boolean if a field has been set.

### GetSWebhookheaderName

`func (o *WebhookheaderRequest) GetSWebhookheaderName() string`

GetSWebhookheaderName returns the SWebhookheaderName field if non-nil, zero value otherwise.

### GetSWebhookheaderNameOk

`func (o *WebhookheaderRequest) GetSWebhookheaderNameOk() (*string, bool)`

GetSWebhookheaderNameOk returns a tuple with the SWebhookheaderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSWebhookheaderName

`func (o *WebhookheaderRequest) SetSWebhookheaderName(v string)`

SetSWebhookheaderName sets SWebhookheaderName field to given value.


### GetSWebhookheaderValue

`func (o *WebhookheaderRequest) GetSWebhookheaderValue() string`

GetSWebhookheaderValue returns the SWebhookheaderValue field if non-nil, zero value otherwise.

### GetSWebhookheaderValueOk

`func (o *WebhookheaderRequest) GetSWebhookheaderValueOk() (*string, bool)`

GetSWebhookheaderValueOk returns a tuple with the SWebhookheaderValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSWebhookheaderValue

`func (o *WebhookheaderRequest) SetSWebhookheaderValue(v string)`

SetSWebhookheaderValue sets SWebhookheaderValue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


