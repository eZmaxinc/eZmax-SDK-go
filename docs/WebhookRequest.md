# WebhookRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiWebhookID** | Pointer to **int32** | The unique ID of the Webhook | [optional] 
**FkiEzsignfoldertypeID** | Pointer to **int32** | The unique ID of the Ezsignfoldertype. | [optional] 
**SWebhookDescription** | **string** | The description of the Webhook | 
**EWebhookModule** | [**FieldEWebhookModule**](FieldEWebhookModule.md) |  | 
**EWebhookEzsignevent** | Pointer to [**FieldEWebhookEzsignevent**](FieldEWebhookEzsignevent.md) |  | [optional] 
**EWebhookManagementevent** | Pointer to [**FieldEWebhookManagementevent**](FieldEWebhookManagementevent.md) |  | [optional] 
**SWebhookUrl** | **string** | The URL of the Webhook callback | 
**SWebhookEmailfailed** | **string** | The email that will receive the Webhook in case all attempts fail | 
**BWebhookIsactive** | **bool** | Whether the Webhook is active or not | 
**BWebhookSkipsslvalidation** | **bool** | Wheter the server&#39;s SSL certificate should be validated or not. Not recommended to skip for production use | 

## Methods

### NewWebhookRequest

`func NewWebhookRequest(sWebhookDescription string, eWebhookModule FieldEWebhookModule, sWebhookUrl string, sWebhookEmailfailed string, bWebhookIsactive bool, bWebhookSkipsslvalidation bool, ) *WebhookRequest`

NewWebhookRequest instantiates a new WebhookRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookRequestWithDefaults

`func NewWebhookRequestWithDefaults() *WebhookRequest`

NewWebhookRequestWithDefaults instantiates a new WebhookRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiWebhookID

`func (o *WebhookRequest) GetPkiWebhookID() int32`

GetPkiWebhookID returns the PkiWebhookID field if non-nil, zero value otherwise.

### GetPkiWebhookIDOk

`func (o *WebhookRequest) GetPkiWebhookIDOk() (*int32, bool)`

GetPkiWebhookIDOk returns a tuple with the PkiWebhookID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiWebhookID

`func (o *WebhookRequest) SetPkiWebhookID(v int32)`

SetPkiWebhookID sets PkiWebhookID field to given value.

### HasPkiWebhookID

`func (o *WebhookRequest) HasPkiWebhookID() bool`

HasPkiWebhookID returns a boolean if a field has been set.

### GetFkiEzsignfoldertypeID

`func (o *WebhookRequest) GetFkiEzsignfoldertypeID() int32`

GetFkiEzsignfoldertypeID returns the FkiEzsignfoldertypeID field if non-nil, zero value otherwise.

### GetFkiEzsignfoldertypeIDOk

`func (o *WebhookRequest) GetFkiEzsignfoldertypeIDOk() (*int32, bool)`

GetFkiEzsignfoldertypeIDOk returns a tuple with the FkiEzsignfoldertypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignfoldertypeID

`func (o *WebhookRequest) SetFkiEzsignfoldertypeID(v int32)`

SetFkiEzsignfoldertypeID sets FkiEzsignfoldertypeID field to given value.

### HasFkiEzsignfoldertypeID

`func (o *WebhookRequest) HasFkiEzsignfoldertypeID() bool`

HasFkiEzsignfoldertypeID returns a boolean if a field has been set.

### GetSWebhookDescription

`func (o *WebhookRequest) GetSWebhookDescription() string`

GetSWebhookDescription returns the SWebhookDescription field if non-nil, zero value otherwise.

### GetSWebhookDescriptionOk

`func (o *WebhookRequest) GetSWebhookDescriptionOk() (*string, bool)`

GetSWebhookDescriptionOk returns a tuple with the SWebhookDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSWebhookDescription

`func (o *WebhookRequest) SetSWebhookDescription(v string)`

SetSWebhookDescription sets SWebhookDescription field to given value.


### GetEWebhookModule

`func (o *WebhookRequest) GetEWebhookModule() FieldEWebhookModule`

GetEWebhookModule returns the EWebhookModule field if non-nil, zero value otherwise.

### GetEWebhookModuleOk

`func (o *WebhookRequest) GetEWebhookModuleOk() (*FieldEWebhookModule, bool)`

GetEWebhookModuleOk returns a tuple with the EWebhookModule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEWebhookModule

`func (o *WebhookRequest) SetEWebhookModule(v FieldEWebhookModule)`

SetEWebhookModule sets EWebhookModule field to given value.


### GetEWebhookEzsignevent

`func (o *WebhookRequest) GetEWebhookEzsignevent() FieldEWebhookEzsignevent`

GetEWebhookEzsignevent returns the EWebhookEzsignevent field if non-nil, zero value otherwise.

### GetEWebhookEzsigneventOk

`func (o *WebhookRequest) GetEWebhookEzsigneventOk() (*FieldEWebhookEzsignevent, bool)`

GetEWebhookEzsigneventOk returns a tuple with the EWebhookEzsignevent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEWebhookEzsignevent

`func (o *WebhookRequest) SetEWebhookEzsignevent(v FieldEWebhookEzsignevent)`

SetEWebhookEzsignevent sets EWebhookEzsignevent field to given value.

### HasEWebhookEzsignevent

`func (o *WebhookRequest) HasEWebhookEzsignevent() bool`

HasEWebhookEzsignevent returns a boolean if a field has been set.

### GetEWebhookManagementevent

`func (o *WebhookRequest) GetEWebhookManagementevent() FieldEWebhookManagementevent`

GetEWebhookManagementevent returns the EWebhookManagementevent field if non-nil, zero value otherwise.

### GetEWebhookManagementeventOk

`func (o *WebhookRequest) GetEWebhookManagementeventOk() (*FieldEWebhookManagementevent, bool)`

GetEWebhookManagementeventOk returns a tuple with the EWebhookManagementevent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEWebhookManagementevent

`func (o *WebhookRequest) SetEWebhookManagementevent(v FieldEWebhookManagementevent)`

SetEWebhookManagementevent sets EWebhookManagementevent field to given value.

### HasEWebhookManagementevent

`func (o *WebhookRequest) HasEWebhookManagementevent() bool`

HasEWebhookManagementevent returns a boolean if a field has been set.

### GetSWebhookUrl

`func (o *WebhookRequest) GetSWebhookUrl() string`

GetSWebhookUrl returns the SWebhookUrl field if non-nil, zero value otherwise.

### GetSWebhookUrlOk

`func (o *WebhookRequest) GetSWebhookUrlOk() (*string, bool)`

GetSWebhookUrlOk returns a tuple with the SWebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSWebhookUrl

`func (o *WebhookRequest) SetSWebhookUrl(v string)`

SetSWebhookUrl sets SWebhookUrl field to given value.


### GetSWebhookEmailfailed

`func (o *WebhookRequest) GetSWebhookEmailfailed() string`

GetSWebhookEmailfailed returns the SWebhookEmailfailed field if non-nil, zero value otherwise.

### GetSWebhookEmailfailedOk

`func (o *WebhookRequest) GetSWebhookEmailfailedOk() (*string, bool)`

GetSWebhookEmailfailedOk returns a tuple with the SWebhookEmailfailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSWebhookEmailfailed

`func (o *WebhookRequest) SetSWebhookEmailfailed(v string)`

SetSWebhookEmailfailed sets SWebhookEmailfailed field to given value.


### GetBWebhookIsactive

`func (o *WebhookRequest) GetBWebhookIsactive() bool`

GetBWebhookIsactive returns the BWebhookIsactive field if non-nil, zero value otherwise.

### GetBWebhookIsactiveOk

`func (o *WebhookRequest) GetBWebhookIsactiveOk() (*bool, bool)`

GetBWebhookIsactiveOk returns a tuple with the BWebhookIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBWebhookIsactive

`func (o *WebhookRequest) SetBWebhookIsactive(v bool)`

SetBWebhookIsactive sets BWebhookIsactive field to given value.


### GetBWebhookSkipsslvalidation

`func (o *WebhookRequest) GetBWebhookSkipsslvalidation() bool`

GetBWebhookSkipsslvalidation returns the BWebhookSkipsslvalidation field if non-nil, zero value otherwise.

### GetBWebhookSkipsslvalidationOk

`func (o *WebhookRequest) GetBWebhookSkipsslvalidationOk() (*bool, bool)`

GetBWebhookSkipsslvalidationOk returns a tuple with the BWebhookSkipsslvalidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBWebhookSkipsslvalidation

`func (o *WebhookRequest) SetBWebhookSkipsslvalidation(v bool)`

SetBWebhookSkipsslvalidation sets BWebhookSkipsslvalidation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


