# WebhookListElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiWebhookID** | **int32** | The unique ID of the Webhook | 
**SWebhookDescription** | **string** | The description of the Webhook | 
**SWebhookUrl** | **string** | The URL of the Webhook callback | 
**SWebhookEvent** | **string** | The concatenated string to describe the Webhook event | 
**SWebhookEmailfailed** | **string** | The email that will receive the Webhook in case all attempts fail | 
**EWebhookModule** | [**FieldEWebhookModule**](FieldEWebhookModule.md) |  | 
**EWebhookEzsignevent** | Pointer to [**FieldEWebhookEzsignevent**](FieldEWebhookEzsignevent.md) |  | [optional] 
**EWebhookManagementevent** | Pointer to [**FieldEWebhookManagementevent**](FieldEWebhookManagementevent.md) |  | [optional] 
**BWebhookIsactive** | **bool** | Whether the Webhook is active or not | 
**BWebhookIssigned** | **bool** | Whether the requests will be signed or not | 

## Methods

### NewWebhookListElement

`func NewWebhookListElement(pkiWebhookID int32, sWebhookDescription string, sWebhookUrl string, sWebhookEvent string, sWebhookEmailfailed string, eWebhookModule FieldEWebhookModule, bWebhookIsactive bool, bWebhookIssigned bool, ) *WebhookListElement`

NewWebhookListElement instantiates a new WebhookListElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookListElementWithDefaults

`func NewWebhookListElementWithDefaults() *WebhookListElement`

NewWebhookListElementWithDefaults instantiates a new WebhookListElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiWebhookID

`func (o *WebhookListElement) GetPkiWebhookID() int32`

GetPkiWebhookID returns the PkiWebhookID field if non-nil, zero value otherwise.

### GetPkiWebhookIDOk

`func (o *WebhookListElement) GetPkiWebhookIDOk() (*int32, bool)`

GetPkiWebhookIDOk returns a tuple with the PkiWebhookID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiWebhookID

`func (o *WebhookListElement) SetPkiWebhookID(v int32)`

SetPkiWebhookID sets PkiWebhookID field to given value.


### GetSWebhookDescription

`func (o *WebhookListElement) GetSWebhookDescription() string`

GetSWebhookDescription returns the SWebhookDescription field if non-nil, zero value otherwise.

### GetSWebhookDescriptionOk

`func (o *WebhookListElement) GetSWebhookDescriptionOk() (*string, bool)`

GetSWebhookDescriptionOk returns a tuple with the SWebhookDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSWebhookDescription

`func (o *WebhookListElement) SetSWebhookDescription(v string)`

SetSWebhookDescription sets SWebhookDescription field to given value.


### GetSWebhookUrl

`func (o *WebhookListElement) GetSWebhookUrl() string`

GetSWebhookUrl returns the SWebhookUrl field if non-nil, zero value otherwise.

### GetSWebhookUrlOk

`func (o *WebhookListElement) GetSWebhookUrlOk() (*string, bool)`

GetSWebhookUrlOk returns a tuple with the SWebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSWebhookUrl

`func (o *WebhookListElement) SetSWebhookUrl(v string)`

SetSWebhookUrl sets SWebhookUrl field to given value.


### GetSWebhookEvent

`func (o *WebhookListElement) GetSWebhookEvent() string`

GetSWebhookEvent returns the SWebhookEvent field if non-nil, zero value otherwise.

### GetSWebhookEventOk

`func (o *WebhookListElement) GetSWebhookEventOk() (*string, bool)`

GetSWebhookEventOk returns a tuple with the SWebhookEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSWebhookEvent

`func (o *WebhookListElement) SetSWebhookEvent(v string)`

SetSWebhookEvent sets SWebhookEvent field to given value.


### GetSWebhookEmailfailed

`func (o *WebhookListElement) GetSWebhookEmailfailed() string`

GetSWebhookEmailfailed returns the SWebhookEmailfailed field if non-nil, zero value otherwise.

### GetSWebhookEmailfailedOk

`func (o *WebhookListElement) GetSWebhookEmailfailedOk() (*string, bool)`

GetSWebhookEmailfailedOk returns a tuple with the SWebhookEmailfailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSWebhookEmailfailed

`func (o *WebhookListElement) SetSWebhookEmailfailed(v string)`

SetSWebhookEmailfailed sets SWebhookEmailfailed field to given value.


### GetEWebhookModule

`func (o *WebhookListElement) GetEWebhookModule() FieldEWebhookModule`

GetEWebhookModule returns the EWebhookModule field if non-nil, zero value otherwise.

### GetEWebhookModuleOk

`func (o *WebhookListElement) GetEWebhookModuleOk() (*FieldEWebhookModule, bool)`

GetEWebhookModuleOk returns a tuple with the EWebhookModule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEWebhookModule

`func (o *WebhookListElement) SetEWebhookModule(v FieldEWebhookModule)`

SetEWebhookModule sets EWebhookModule field to given value.


### GetEWebhookEzsignevent

`func (o *WebhookListElement) GetEWebhookEzsignevent() FieldEWebhookEzsignevent`

GetEWebhookEzsignevent returns the EWebhookEzsignevent field if non-nil, zero value otherwise.

### GetEWebhookEzsigneventOk

`func (o *WebhookListElement) GetEWebhookEzsigneventOk() (*FieldEWebhookEzsignevent, bool)`

GetEWebhookEzsigneventOk returns a tuple with the EWebhookEzsignevent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEWebhookEzsignevent

`func (o *WebhookListElement) SetEWebhookEzsignevent(v FieldEWebhookEzsignevent)`

SetEWebhookEzsignevent sets EWebhookEzsignevent field to given value.

### HasEWebhookEzsignevent

`func (o *WebhookListElement) HasEWebhookEzsignevent() bool`

HasEWebhookEzsignevent returns a boolean if a field has been set.

### GetEWebhookManagementevent

`func (o *WebhookListElement) GetEWebhookManagementevent() FieldEWebhookManagementevent`

GetEWebhookManagementevent returns the EWebhookManagementevent field if non-nil, zero value otherwise.

### GetEWebhookManagementeventOk

`func (o *WebhookListElement) GetEWebhookManagementeventOk() (*FieldEWebhookManagementevent, bool)`

GetEWebhookManagementeventOk returns a tuple with the EWebhookManagementevent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEWebhookManagementevent

`func (o *WebhookListElement) SetEWebhookManagementevent(v FieldEWebhookManagementevent)`

SetEWebhookManagementevent sets EWebhookManagementevent field to given value.

### HasEWebhookManagementevent

`func (o *WebhookListElement) HasEWebhookManagementevent() bool`

HasEWebhookManagementevent returns a boolean if a field has been set.

### GetBWebhookIsactive

`func (o *WebhookListElement) GetBWebhookIsactive() bool`

GetBWebhookIsactive returns the BWebhookIsactive field if non-nil, zero value otherwise.

### GetBWebhookIsactiveOk

`func (o *WebhookListElement) GetBWebhookIsactiveOk() (*bool, bool)`

GetBWebhookIsactiveOk returns a tuple with the BWebhookIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBWebhookIsactive

`func (o *WebhookListElement) SetBWebhookIsactive(v bool)`

SetBWebhookIsactive sets BWebhookIsactive field to given value.


### GetBWebhookIssigned

`func (o *WebhookListElement) GetBWebhookIssigned() bool`

GetBWebhookIssigned returns the BWebhookIssigned field if non-nil, zero value otherwise.

### GetBWebhookIssignedOk

`func (o *WebhookListElement) GetBWebhookIssignedOk() (*bool, bool)`

GetBWebhookIssignedOk returns a tuple with the BWebhookIssigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBWebhookIssigned

`func (o *WebhookListElement) SetBWebhookIssigned(v bool)`

SetBWebhookIssigned sets BWebhookIssigned field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


