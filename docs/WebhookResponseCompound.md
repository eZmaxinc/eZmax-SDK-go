# WebhookResponseCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiWebhookID** | **int32** | The unique ID of the Webhook | 
**SWebhookDescription** | **string** | The description of the Webhook | 
**FkiEzsignfoldertypeID** | Pointer to **int32** | The unique ID of the Ezsignfoldertype. | [optional] 
**SEzsignfoldertypeNameX** | Pointer to **string** | The name of the Ezsignfoldertype in the language of the requester | [optional] 
**EWebhookModule** | [**FieldEWebhookModule**](FieldEWebhookModule.md) |  | 
**EWebhookEzsignevent** | Pointer to [**FieldEWebhookEzsignevent**](FieldEWebhookEzsignevent.md) |  | [optional] 
**EWebhookManagementevent** | Pointer to [**FieldEWebhookManagementevent**](FieldEWebhookManagementevent.md) |  | [optional] 
**SWebhookUrl** | **string** | The URL of the Webhook callback | 
**SWebhookEmailfailed** | **string** | The email that will receive the Webhook in case all attempts fail | 
**SWebhookApikey** | Pointer to **string** | The Apikey for the Webhook.  This will be hidden if we are not creating or regenerating the Apikey. | [optional] 
**SWebhookSecret** | Pointer to **string** | The Secret for the Webhook.  This will be hidden if we are not creating or regenerating the Apikey. | [optional] 
**BWebhookIsactive** | **bool** | Whether the Webhook is active or not | 
**BWebhookIssigned** | **bool** | Whether the requests will be signed or not | 
**BWebhookSkipsslvalidation** | **bool** | Wheter the server&#39;s SSL certificate should be validated or not. Not recommended to skip for production use | 
**SWebhookEvent** | Pointer to **string** | The concatenated string to describe the Webhook event | [optional] 

## Methods

### NewWebhookResponseCompound

`func NewWebhookResponseCompound(pkiWebhookID int32, sWebhookDescription string, eWebhookModule FieldEWebhookModule, sWebhookUrl string, sWebhookEmailfailed string, bWebhookIsactive bool, bWebhookIssigned bool, bWebhookSkipsslvalidation bool, ) *WebhookResponseCompound`

NewWebhookResponseCompound instantiates a new WebhookResponseCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookResponseCompoundWithDefaults

`func NewWebhookResponseCompoundWithDefaults() *WebhookResponseCompound`

NewWebhookResponseCompoundWithDefaults instantiates a new WebhookResponseCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiWebhookID

`func (o *WebhookResponseCompound) GetPkiWebhookID() int32`

GetPkiWebhookID returns the PkiWebhookID field if non-nil, zero value otherwise.

### GetPkiWebhookIDOk

`func (o *WebhookResponseCompound) GetPkiWebhookIDOk() (*int32, bool)`

GetPkiWebhookIDOk returns a tuple with the PkiWebhookID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiWebhookID

`func (o *WebhookResponseCompound) SetPkiWebhookID(v int32)`

SetPkiWebhookID sets PkiWebhookID field to given value.


### GetSWebhookDescription

`func (o *WebhookResponseCompound) GetSWebhookDescription() string`

GetSWebhookDescription returns the SWebhookDescription field if non-nil, zero value otherwise.

### GetSWebhookDescriptionOk

`func (o *WebhookResponseCompound) GetSWebhookDescriptionOk() (*string, bool)`

GetSWebhookDescriptionOk returns a tuple with the SWebhookDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSWebhookDescription

`func (o *WebhookResponseCompound) SetSWebhookDescription(v string)`

SetSWebhookDescription sets SWebhookDescription field to given value.


### GetFkiEzsignfoldertypeID

`func (o *WebhookResponseCompound) GetFkiEzsignfoldertypeID() int32`

GetFkiEzsignfoldertypeID returns the FkiEzsignfoldertypeID field if non-nil, zero value otherwise.

### GetFkiEzsignfoldertypeIDOk

`func (o *WebhookResponseCompound) GetFkiEzsignfoldertypeIDOk() (*int32, bool)`

GetFkiEzsignfoldertypeIDOk returns a tuple with the FkiEzsignfoldertypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignfoldertypeID

`func (o *WebhookResponseCompound) SetFkiEzsignfoldertypeID(v int32)`

SetFkiEzsignfoldertypeID sets FkiEzsignfoldertypeID field to given value.

### HasFkiEzsignfoldertypeID

`func (o *WebhookResponseCompound) HasFkiEzsignfoldertypeID() bool`

HasFkiEzsignfoldertypeID returns a boolean if a field has been set.

### GetSEzsignfoldertypeNameX

`func (o *WebhookResponseCompound) GetSEzsignfoldertypeNameX() string`

GetSEzsignfoldertypeNameX returns the SEzsignfoldertypeNameX field if non-nil, zero value otherwise.

### GetSEzsignfoldertypeNameXOk

`func (o *WebhookResponseCompound) GetSEzsignfoldertypeNameXOk() (*string, bool)`

GetSEzsignfoldertypeNameXOk returns a tuple with the SEzsignfoldertypeNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignfoldertypeNameX

`func (o *WebhookResponseCompound) SetSEzsignfoldertypeNameX(v string)`

SetSEzsignfoldertypeNameX sets SEzsignfoldertypeNameX field to given value.

### HasSEzsignfoldertypeNameX

`func (o *WebhookResponseCompound) HasSEzsignfoldertypeNameX() bool`

HasSEzsignfoldertypeNameX returns a boolean if a field has been set.

### GetEWebhookModule

`func (o *WebhookResponseCompound) GetEWebhookModule() FieldEWebhookModule`

GetEWebhookModule returns the EWebhookModule field if non-nil, zero value otherwise.

### GetEWebhookModuleOk

`func (o *WebhookResponseCompound) GetEWebhookModuleOk() (*FieldEWebhookModule, bool)`

GetEWebhookModuleOk returns a tuple with the EWebhookModule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEWebhookModule

`func (o *WebhookResponseCompound) SetEWebhookModule(v FieldEWebhookModule)`

SetEWebhookModule sets EWebhookModule field to given value.


### GetEWebhookEzsignevent

`func (o *WebhookResponseCompound) GetEWebhookEzsignevent() FieldEWebhookEzsignevent`

GetEWebhookEzsignevent returns the EWebhookEzsignevent field if non-nil, zero value otherwise.

### GetEWebhookEzsigneventOk

`func (o *WebhookResponseCompound) GetEWebhookEzsigneventOk() (*FieldEWebhookEzsignevent, bool)`

GetEWebhookEzsigneventOk returns a tuple with the EWebhookEzsignevent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEWebhookEzsignevent

`func (o *WebhookResponseCompound) SetEWebhookEzsignevent(v FieldEWebhookEzsignevent)`

SetEWebhookEzsignevent sets EWebhookEzsignevent field to given value.

### HasEWebhookEzsignevent

`func (o *WebhookResponseCompound) HasEWebhookEzsignevent() bool`

HasEWebhookEzsignevent returns a boolean if a field has been set.

### GetEWebhookManagementevent

`func (o *WebhookResponseCompound) GetEWebhookManagementevent() FieldEWebhookManagementevent`

GetEWebhookManagementevent returns the EWebhookManagementevent field if non-nil, zero value otherwise.

### GetEWebhookManagementeventOk

`func (o *WebhookResponseCompound) GetEWebhookManagementeventOk() (*FieldEWebhookManagementevent, bool)`

GetEWebhookManagementeventOk returns a tuple with the EWebhookManagementevent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEWebhookManagementevent

`func (o *WebhookResponseCompound) SetEWebhookManagementevent(v FieldEWebhookManagementevent)`

SetEWebhookManagementevent sets EWebhookManagementevent field to given value.

### HasEWebhookManagementevent

`func (o *WebhookResponseCompound) HasEWebhookManagementevent() bool`

HasEWebhookManagementevent returns a boolean if a field has been set.

### GetSWebhookUrl

`func (o *WebhookResponseCompound) GetSWebhookUrl() string`

GetSWebhookUrl returns the SWebhookUrl field if non-nil, zero value otherwise.

### GetSWebhookUrlOk

`func (o *WebhookResponseCompound) GetSWebhookUrlOk() (*string, bool)`

GetSWebhookUrlOk returns a tuple with the SWebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSWebhookUrl

`func (o *WebhookResponseCompound) SetSWebhookUrl(v string)`

SetSWebhookUrl sets SWebhookUrl field to given value.


### GetSWebhookEmailfailed

`func (o *WebhookResponseCompound) GetSWebhookEmailfailed() string`

GetSWebhookEmailfailed returns the SWebhookEmailfailed field if non-nil, zero value otherwise.

### GetSWebhookEmailfailedOk

`func (o *WebhookResponseCompound) GetSWebhookEmailfailedOk() (*string, bool)`

GetSWebhookEmailfailedOk returns a tuple with the SWebhookEmailfailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSWebhookEmailfailed

`func (o *WebhookResponseCompound) SetSWebhookEmailfailed(v string)`

SetSWebhookEmailfailed sets SWebhookEmailfailed field to given value.


### GetSWebhookApikey

`func (o *WebhookResponseCompound) GetSWebhookApikey() string`

GetSWebhookApikey returns the SWebhookApikey field if non-nil, zero value otherwise.

### GetSWebhookApikeyOk

`func (o *WebhookResponseCompound) GetSWebhookApikeyOk() (*string, bool)`

GetSWebhookApikeyOk returns a tuple with the SWebhookApikey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSWebhookApikey

`func (o *WebhookResponseCompound) SetSWebhookApikey(v string)`

SetSWebhookApikey sets SWebhookApikey field to given value.

### HasSWebhookApikey

`func (o *WebhookResponseCompound) HasSWebhookApikey() bool`

HasSWebhookApikey returns a boolean if a field has been set.

### GetSWebhookSecret

`func (o *WebhookResponseCompound) GetSWebhookSecret() string`

GetSWebhookSecret returns the SWebhookSecret field if non-nil, zero value otherwise.

### GetSWebhookSecretOk

`func (o *WebhookResponseCompound) GetSWebhookSecretOk() (*string, bool)`

GetSWebhookSecretOk returns a tuple with the SWebhookSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSWebhookSecret

`func (o *WebhookResponseCompound) SetSWebhookSecret(v string)`

SetSWebhookSecret sets SWebhookSecret field to given value.

### HasSWebhookSecret

`func (o *WebhookResponseCompound) HasSWebhookSecret() bool`

HasSWebhookSecret returns a boolean if a field has been set.

### GetBWebhookIsactive

`func (o *WebhookResponseCompound) GetBWebhookIsactive() bool`

GetBWebhookIsactive returns the BWebhookIsactive field if non-nil, zero value otherwise.

### GetBWebhookIsactiveOk

`func (o *WebhookResponseCompound) GetBWebhookIsactiveOk() (*bool, bool)`

GetBWebhookIsactiveOk returns a tuple with the BWebhookIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBWebhookIsactive

`func (o *WebhookResponseCompound) SetBWebhookIsactive(v bool)`

SetBWebhookIsactive sets BWebhookIsactive field to given value.


### GetBWebhookIssigned

`func (o *WebhookResponseCompound) GetBWebhookIssigned() bool`

GetBWebhookIssigned returns the BWebhookIssigned field if non-nil, zero value otherwise.

### GetBWebhookIssignedOk

`func (o *WebhookResponseCompound) GetBWebhookIssignedOk() (*bool, bool)`

GetBWebhookIssignedOk returns a tuple with the BWebhookIssigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBWebhookIssigned

`func (o *WebhookResponseCompound) SetBWebhookIssigned(v bool)`

SetBWebhookIssigned sets BWebhookIssigned field to given value.


### GetBWebhookSkipsslvalidation

`func (o *WebhookResponseCompound) GetBWebhookSkipsslvalidation() bool`

GetBWebhookSkipsslvalidation returns the BWebhookSkipsslvalidation field if non-nil, zero value otherwise.

### GetBWebhookSkipsslvalidationOk

`func (o *WebhookResponseCompound) GetBWebhookSkipsslvalidationOk() (*bool, bool)`

GetBWebhookSkipsslvalidationOk returns a tuple with the BWebhookSkipsslvalidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBWebhookSkipsslvalidation

`func (o *WebhookResponseCompound) SetBWebhookSkipsslvalidation(v bool)`

SetBWebhookSkipsslvalidation sets BWebhookSkipsslvalidation field to given value.


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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


