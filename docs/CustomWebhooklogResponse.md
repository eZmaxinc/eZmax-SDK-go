# CustomWebhooklogResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DtWebhooklogDate** | **string** | The date and time at which the Webhooklog happened. | 
**TWebhooklogJson** | **string** | The Json containing the Webhook call and return | 

## Methods

### NewCustomWebhooklogResponse

`func NewCustomWebhooklogResponse(dtWebhooklogDate string, tWebhooklogJson string, ) *CustomWebhooklogResponse`

NewCustomWebhooklogResponse instantiates a new CustomWebhooklogResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomWebhooklogResponseWithDefaults

`func NewCustomWebhooklogResponseWithDefaults() *CustomWebhooklogResponse`

NewCustomWebhooklogResponseWithDefaults instantiates a new CustomWebhooklogResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDtWebhooklogDate

`func (o *CustomWebhooklogResponse) GetDtWebhooklogDate() string`

GetDtWebhooklogDate returns the DtWebhooklogDate field if non-nil, zero value otherwise.

### GetDtWebhooklogDateOk

`func (o *CustomWebhooklogResponse) GetDtWebhooklogDateOk() (*string, bool)`

GetDtWebhooklogDateOk returns a tuple with the DtWebhooklogDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtWebhooklogDate

`func (o *CustomWebhooklogResponse) SetDtWebhooklogDate(v string)`

SetDtWebhooklogDate sets DtWebhooklogDate field to given value.


### GetTWebhooklogJson

`func (o *CustomWebhooklogResponse) GetTWebhooklogJson() string`

GetTWebhooklogJson returns the TWebhooklogJson field if non-nil, zero value otherwise.

### GetTWebhooklogJsonOk

`func (o *CustomWebhooklogResponse) GetTWebhooklogJsonOk() (*string, bool)`

GetTWebhooklogJsonOk returns a tuple with the TWebhooklogJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTWebhooklogJson

`func (o *CustomWebhooklogResponse) SetTWebhooklogJson(v string)`

SetTWebhooklogJson sets TWebhooklogJson field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


