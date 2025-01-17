# CustomWebhookResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PksCustomerCode** | **string** | The customer code assigned to your account | 
**BWebhookTest** | **bool** | Wheter the webhook received is a manual test or a real event | 
**EWebhookEmittype** | Pointer to **string** | Wheter the webhook received is a manual test or a real event | [optional] 

## Methods

### NewCustomWebhookResponse

`func NewCustomWebhookResponse(pksCustomerCode string, bWebhookTest bool, ) *CustomWebhookResponse`

NewCustomWebhookResponse instantiates a new CustomWebhookResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomWebhookResponseWithDefaults

`func NewCustomWebhookResponseWithDefaults() *CustomWebhookResponse`

NewCustomWebhookResponseWithDefaults instantiates a new CustomWebhookResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPksCustomerCode

`func (o *CustomWebhookResponse) GetPksCustomerCode() string`

GetPksCustomerCode returns the PksCustomerCode field if non-nil, zero value otherwise.

### GetPksCustomerCodeOk

`func (o *CustomWebhookResponse) GetPksCustomerCodeOk() (*string, bool)`

GetPksCustomerCodeOk returns a tuple with the PksCustomerCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPksCustomerCode

`func (o *CustomWebhookResponse) SetPksCustomerCode(v string)`

SetPksCustomerCode sets PksCustomerCode field to given value.


### GetBWebhookTest

`func (o *CustomWebhookResponse) GetBWebhookTest() bool`

GetBWebhookTest returns the BWebhookTest field if non-nil, zero value otherwise.

### GetBWebhookTestOk

`func (o *CustomWebhookResponse) GetBWebhookTestOk() (*bool, bool)`

GetBWebhookTestOk returns a tuple with the BWebhookTest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBWebhookTest

`func (o *CustomWebhookResponse) SetBWebhookTest(v bool)`

SetBWebhookTest sets BWebhookTest field to given value.


### GetEWebhookEmittype

`func (o *CustomWebhookResponse) GetEWebhookEmittype() string`

GetEWebhookEmittype returns the EWebhookEmittype field if non-nil, zero value otherwise.

### GetEWebhookEmittypeOk

`func (o *CustomWebhookResponse) GetEWebhookEmittypeOk() (*string, bool)`

GetEWebhookEmittypeOk returns a tuple with the EWebhookEmittype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEWebhookEmittype

`func (o *CustomWebhookResponse) SetEWebhookEmittype(v string)`

SetEWebhookEmittype sets EWebhookEmittype field to given value.

### HasEWebhookEmittype

`func (o *CustomWebhookResponse) HasEWebhookEmittype() bool`

HasEWebhookEmittype returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


