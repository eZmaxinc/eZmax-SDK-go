# WebhookResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiWebhookID** | **int32** | The Webhook ID. This value is visible in the admin interface. | 
**EWebhookModule** | **string** | The Module generating the Event. | 
**EWebhookEzsignevent** | Pointer to **string** | This Ezsign Event. This property will be set only if the Module is \&quot;Ezsign\&quot;. | [optional] 
**PksCustomerCode** | **string** | The Customer Code in which the event was generated | 
**SWebhookUrl** | **string** | The url being called | 
**SWebhookEmailfailed** | **string** | The email that will receive the webhook in case all attempts fail. | 
**EWebhookManagementevent** | Pointer to **string** | This Management Event. This property will be set only if the Module is \&quot;Management\&quot;. | [optional] 

## Methods

### NewWebhookResponse

`func NewWebhookResponse(pkiWebhookID int32, eWebhookModule string, pksCustomerCode string, sWebhookUrl string, sWebhookEmailfailed string, ) *WebhookResponse`

NewWebhookResponse instantiates a new WebhookResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookResponseWithDefaults

`func NewWebhookResponseWithDefaults() *WebhookResponse`

NewWebhookResponseWithDefaults instantiates a new WebhookResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiWebhookID

`func (o *WebhookResponse) GetPkiWebhookID() int32`

GetPkiWebhookID returns the PkiWebhookID field if non-nil, zero value otherwise.

### GetPkiWebhookIDOk

`func (o *WebhookResponse) GetPkiWebhookIDOk() (*int32, bool)`

GetPkiWebhookIDOk returns a tuple with the PkiWebhookID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiWebhookID

`func (o *WebhookResponse) SetPkiWebhookID(v int32)`

SetPkiWebhookID sets PkiWebhookID field to given value.


### GetEWebhookModule

`func (o *WebhookResponse) GetEWebhookModule() string`

GetEWebhookModule returns the EWebhookModule field if non-nil, zero value otherwise.

### GetEWebhookModuleOk

`func (o *WebhookResponse) GetEWebhookModuleOk() (*string, bool)`

GetEWebhookModuleOk returns a tuple with the EWebhookModule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEWebhookModule

`func (o *WebhookResponse) SetEWebhookModule(v string)`

SetEWebhookModule sets EWebhookModule field to given value.


### GetEWebhookEzsignevent

`func (o *WebhookResponse) GetEWebhookEzsignevent() string`

GetEWebhookEzsignevent returns the EWebhookEzsignevent field if non-nil, zero value otherwise.

### GetEWebhookEzsigneventOk

`func (o *WebhookResponse) GetEWebhookEzsigneventOk() (*string, bool)`

GetEWebhookEzsigneventOk returns a tuple with the EWebhookEzsignevent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEWebhookEzsignevent

`func (o *WebhookResponse) SetEWebhookEzsignevent(v string)`

SetEWebhookEzsignevent sets EWebhookEzsignevent field to given value.

### HasEWebhookEzsignevent

`func (o *WebhookResponse) HasEWebhookEzsignevent() bool`

HasEWebhookEzsignevent returns a boolean if a field has been set.

### GetPksCustomerCode

`func (o *WebhookResponse) GetPksCustomerCode() string`

GetPksCustomerCode returns the PksCustomerCode field if non-nil, zero value otherwise.

### GetPksCustomerCodeOk

`func (o *WebhookResponse) GetPksCustomerCodeOk() (*string, bool)`

GetPksCustomerCodeOk returns a tuple with the PksCustomerCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPksCustomerCode

`func (o *WebhookResponse) SetPksCustomerCode(v string)`

SetPksCustomerCode sets PksCustomerCode field to given value.


### GetSWebhookUrl

`func (o *WebhookResponse) GetSWebhookUrl() string`

GetSWebhookUrl returns the SWebhookUrl field if non-nil, zero value otherwise.

### GetSWebhookUrlOk

`func (o *WebhookResponse) GetSWebhookUrlOk() (*string, bool)`

GetSWebhookUrlOk returns a tuple with the SWebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSWebhookUrl

`func (o *WebhookResponse) SetSWebhookUrl(v string)`

SetSWebhookUrl sets SWebhookUrl field to given value.


### GetSWebhookEmailfailed

`func (o *WebhookResponse) GetSWebhookEmailfailed() string`

GetSWebhookEmailfailed returns the SWebhookEmailfailed field if non-nil, zero value otherwise.

### GetSWebhookEmailfailedOk

`func (o *WebhookResponse) GetSWebhookEmailfailedOk() (*string, bool)`

GetSWebhookEmailfailedOk returns a tuple with the SWebhookEmailfailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSWebhookEmailfailed

`func (o *WebhookResponse) SetSWebhookEmailfailed(v string)`

SetSWebhookEmailfailed sets SWebhookEmailfailed field to given value.


### GetEWebhookManagementevent

`func (o *WebhookResponse) GetEWebhookManagementevent() string`

GetEWebhookManagementevent returns the EWebhookManagementevent field if non-nil, zero value otherwise.

### GetEWebhookManagementeventOk

`func (o *WebhookResponse) GetEWebhookManagementeventOk() (*string, bool)`

GetEWebhookManagementeventOk returns a tuple with the EWebhookManagementevent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEWebhookManagementevent

`func (o *WebhookResponse) SetEWebhookManagementevent(v string)`

SetEWebhookManagementevent sets EWebhookManagementevent field to given value.

### HasEWebhookManagementevent

`func (o *WebhookResponse) HasEWebhookManagementevent() bool`

HasEWebhookManagementevent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


