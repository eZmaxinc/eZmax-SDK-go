# WebhookSendWebhookV1Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EWebhookModule** | [**FieldEWebhookModule**](FieldEWebhookModule.md) |  | 
**EWebhookEzsignevent** | Pointer to [**CustomEWebhookEzsignevent**](CustomEWebhookEzsignevent.md) |  | [optional] 
**EWebhookManagementevent** | Pointer to [**FieldEWebhookManagementevent**](FieldEWebhookManagementevent.md) |  | [optional] 
**FkiEzsignfolderID** | Pointer to **int32** | The unique ID of the Ezsignfolder | [optional] 
**FkiEzsigndocumentID** | Pointer to **int32** | The unique ID of the Ezsigndocument | [optional] 
**FkiEzsignsignerID** | Pointer to **int32** | The unique ID of the Ezsignsigner | [optional] 
**FkiUserID** | Pointer to **int32** | The unique ID of the User | [optional] 
**FkiUserstagedID** | Pointer to **int32** | The unique ID of the Userstaged | [optional] 

## Methods

### NewWebhookSendWebhookV1Request

`func NewWebhookSendWebhookV1Request(eWebhookModule FieldEWebhookModule, ) *WebhookSendWebhookV1Request`

NewWebhookSendWebhookV1Request instantiates a new WebhookSendWebhookV1Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookSendWebhookV1RequestWithDefaults

`func NewWebhookSendWebhookV1RequestWithDefaults() *WebhookSendWebhookV1Request`

NewWebhookSendWebhookV1RequestWithDefaults instantiates a new WebhookSendWebhookV1Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEWebhookModule

`func (o *WebhookSendWebhookV1Request) GetEWebhookModule() FieldEWebhookModule`

GetEWebhookModule returns the EWebhookModule field if non-nil, zero value otherwise.

### GetEWebhookModuleOk

`func (o *WebhookSendWebhookV1Request) GetEWebhookModuleOk() (*FieldEWebhookModule, bool)`

GetEWebhookModuleOk returns a tuple with the EWebhookModule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEWebhookModule

`func (o *WebhookSendWebhookV1Request) SetEWebhookModule(v FieldEWebhookModule)`

SetEWebhookModule sets EWebhookModule field to given value.


### GetEWebhookEzsignevent

`func (o *WebhookSendWebhookV1Request) GetEWebhookEzsignevent() CustomEWebhookEzsignevent`

GetEWebhookEzsignevent returns the EWebhookEzsignevent field if non-nil, zero value otherwise.

### GetEWebhookEzsigneventOk

`func (o *WebhookSendWebhookV1Request) GetEWebhookEzsigneventOk() (*CustomEWebhookEzsignevent, bool)`

GetEWebhookEzsigneventOk returns a tuple with the EWebhookEzsignevent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEWebhookEzsignevent

`func (o *WebhookSendWebhookV1Request) SetEWebhookEzsignevent(v CustomEWebhookEzsignevent)`

SetEWebhookEzsignevent sets EWebhookEzsignevent field to given value.

### HasEWebhookEzsignevent

`func (o *WebhookSendWebhookV1Request) HasEWebhookEzsignevent() bool`

HasEWebhookEzsignevent returns a boolean if a field has been set.

### GetEWebhookManagementevent

`func (o *WebhookSendWebhookV1Request) GetEWebhookManagementevent() FieldEWebhookManagementevent`

GetEWebhookManagementevent returns the EWebhookManagementevent field if non-nil, zero value otherwise.

### GetEWebhookManagementeventOk

`func (o *WebhookSendWebhookV1Request) GetEWebhookManagementeventOk() (*FieldEWebhookManagementevent, bool)`

GetEWebhookManagementeventOk returns a tuple with the EWebhookManagementevent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEWebhookManagementevent

`func (o *WebhookSendWebhookV1Request) SetEWebhookManagementevent(v FieldEWebhookManagementevent)`

SetEWebhookManagementevent sets EWebhookManagementevent field to given value.

### HasEWebhookManagementevent

`func (o *WebhookSendWebhookV1Request) HasEWebhookManagementevent() bool`

HasEWebhookManagementevent returns a boolean if a field has been set.

### GetFkiEzsignfolderID

`func (o *WebhookSendWebhookV1Request) GetFkiEzsignfolderID() int32`

GetFkiEzsignfolderID returns the FkiEzsignfolderID field if non-nil, zero value otherwise.

### GetFkiEzsignfolderIDOk

`func (o *WebhookSendWebhookV1Request) GetFkiEzsignfolderIDOk() (*int32, bool)`

GetFkiEzsignfolderIDOk returns a tuple with the FkiEzsignfolderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignfolderID

`func (o *WebhookSendWebhookV1Request) SetFkiEzsignfolderID(v int32)`

SetFkiEzsignfolderID sets FkiEzsignfolderID field to given value.

### HasFkiEzsignfolderID

`func (o *WebhookSendWebhookV1Request) HasFkiEzsignfolderID() bool`

HasFkiEzsignfolderID returns a boolean if a field has been set.

### GetFkiEzsigndocumentID

`func (o *WebhookSendWebhookV1Request) GetFkiEzsigndocumentID() int32`

GetFkiEzsigndocumentID returns the FkiEzsigndocumentID field if non-nil, zero value otherwise.

### GetFkiEzsigndocumentIDOk

`func (o *WebhookSendWebhookV1Request) GetFkiEzsigndocumentIDOk() (*int32, bool)`

GetFkiEzsigndocumentIDOk returns a tuple with the FkiEzsigndocumentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigndocumentID

`func (o *WebhookSendWebhookV1Request) SetFkiEzsigndocumentID(v int32)`

SetFkiEzsigndocumentID sets FkiEzsigndocumentID field to given value.

### HasFkiEzsigndocumentID

`func (o *WebhookSendWebhookV1Request) HasFkiEzsigndocumentID() bool`

HasFkiEzsigndocumentID returns a boolean if a field has been set.

### GetFkiEzsignsignerID

`func (o *WebhookSendWebhookV1Request) GetFkiEzsignsignerID() int32`

GetFkiEzsignsignerID returns the FkiEzsignsignerID field if non-nil, zero value otherwise.

### GetFkiEzsignsignerIDOk

`func (o *WebhookSendWebhookV1Request) GetFkiEzsignsignerIDOk() (*int32, bool)`

GetFkiEzsignsignerIDOk returns a tuple with the FkiEzsignsignerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignsignerID

`func (o *WebhookSendWebhookV1Request) SetFkiEzsignsignerID(v int32)`

SetFkiEzsignsignerID sets FkiEzsignsignerID field to given value.

### HasFkiEzsignsignerID

`func (o *WebhookSendWebhookV1Request) HasFkiEzsignsignerID() bool`

HasFkiEzsignsignerID returns a boolean if a field has been set.

### GetFkiUserID

`func (o *WebhookSendWebhookV1Request) GetFkiUserID() int32`

GetFkiUserID returns the FkiUserID field if non-nil, zero value otherwise.

### GetFkiUserIDOk

`func (o *WebhookSendWebhookV1Request) GetFkiUserIDOk() (*int32, bool)`

GetFkiUserIDOk returns a tuple with the FkiUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUserID

`func (o *WebhookSendWebhookV1Request) SetFkiUserID(v int32)`

SetFkiUserID sets FkiUserID field to given value.

### HasFkiUserID

`func (o *WebhookSendWebhookV1Request) HasFkiUserID() bool`

HasFkiUserID returns a boolean if a field has been set.

### GetFkiUserstagedID

`func (o *WebhookSendWebhookV1Request) GetFkiUserstagedID() int32`

GetFkiUserstagedID returns the FkiUserstagedID field if non-nil, zero value otherwise.

### GetFkiUserstagedIDOk

`func (o *WebhookSendWebhookV1Request) GetFkiUserstagedIDOk() (*int32, bool)`

GetFkiUserstagedIDOk returns a tuple with the FkiUserstagedID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUserstagedID

`func (o *WebhookSendWebhookV1Request) SetFkiUserstagedID(v int32)`

SetFkiUserstagedID sets FkiUserstagedID field to given value.

### HasFkiUserstagedID

`func (o *WebhookSendWebhookV1Request) HasFkiUserstagedID() bool`

HasFkiUserstagedID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


