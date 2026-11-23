# EzmaxpartnerproductstagewebhookResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzmaxpartnerproductstagewebhookID** | Pointer to **int32** | The unique ID of the Ezmaxpartnerproductstagewebhook | [optional] 
**FkiEzmaxpartnerproductstageID** | **int32** | The unique ID of the Ezmaxpartnerproductstage | 
**SEzmaxpartnerproductstageCode** | Pointer to **string** | The code of the sEzmaxpartnerproductstage | [optional] 
**EEzmaxpartnerproductstagewebhookModule** | [**FieldEEzmaxpartnerproductstagewebhookModule**](FieldEEzmaxpartnerproductstagewebhookModule.md) |  | 
**EEzmaxpartnerproductstagewebhookEzmaxpartnerevent** | Pointer to [**FieldEEzmaxpartnerproductstagewebhookEzmaxpartnerevent**](FieldEEzmaxpartnerproductstagewebhookEzmaxpartnerevent.md) |  | [optional] 
**EEzmaxpartnerproductstagewebhookEzsignevent** | Pointer to [**FieldEEzmaxpartnerproductstagewebhookEzsignevent**](FieldEEzmaxpartnerproductstagewebhookEzsignevent.md) |  | [optional] 
**EEzmaxpartnerproductstagewebhookManagementevent** | Pointer to [**FieldEEzmaxpartnerproductstagewebhookManagementevent**](FieldEEzmaxpartnerproductstagewebhookManagementevent.md) |  | [optional] 
**EEzmaxpartnerproductstagewebhookRealestateevent** | Pointer to [**FieldEEzmaxpartnerproductstagewebhookRealestateevent**](FieldEEzmaxpartnerproductstagewebhookRealestateevent.md) |  | [optional] 
**SEzmaxpartnerproductstagewebhookDescription** | Pointer to **string** | The description of the Ezmaxpartnerproductstagewebhook | [optional] 
**SEzmaxpartnerproductstagewebhookUrl** | Pointer to **string** | The url of the Ezmaxpartnerproductstagewebhook | [optional] 
**SEzmaxpartnerproductstagewebhookEmailfailed** | Pointer to **string** | The emailfailed of the Ezmaxpartnerproductstagewebhook | [optional] 

## Methods

### NewEzmaxpartnerproductstagewebhookResponse

`func NewEzmaxpartnerproductstagewebhookResponse(fkiEzmaxpartnerproductstageID int32, eEzmaxpartnerproductstagewebhookModule FieldEEzmaxpartnerproductstagewebhookModule, ) *EzmaxpartnerproductstagewebhookResponse`

NewEzmaxpartnerproductstagewebhookResponse instantiates a new EzmaxpartnerproductstagewebhookResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzmaxpartnerproductstagewebhookResponseWithDefaults

`func NewEzmaxpartnerproductstagewebhookResponseWithDefaults() *EzmaxpartnerproductstagewebhookResponse`

NewEzmaxpartnerproductstagewebhookResponseWithDefaults instantiates a new EzmaxpartnerproductstagewebhookResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzmaxpartnerproductstagewebhookID

`func (o *EzmaxpartnerproductstagewebhookResponse) GetPkiEzmaxpartnerproductstagewebhookID() int32`

GetPkiEzmaxpartnerproductstagewebhookID returns the PkiEzmaxpartnerproductstagewebhookID field if non-nil, zero value otherwise.

### GetPkiEzmaxpartnerproductstagewebhookIDOk

`func (o *EzmaxpartnerproductstagewebhookResponse) GetPkiEzmaxpartnerproductstagewebhookIDOk() (*int32, bool)`

GetPkiEzmaxpartnerproductstagewebhookIDOk returns a tuple with the PkiEzmaxpartnerproductstagewebhookID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzmaxpartnerproductstagewebhookID

`func (o *EzmaxpartnerproductstagewebhookResponse) SetPkiEzmaxpartnerproductstagewebhookID(v int32)`

SetPkiEzmaxpartnerproductstagewebhookID sets PkiEzmaxpartnerproductstagewebhookID field to given value.

### HasPkiEzmaxpartnerproductstagewebhookID

`func (o *EzmaxpartnerproductstagewebhookResponse) HasPkiEzmaxpartnerproductstagewebhookID() bool`

HasPkiEzmaxpartnerproductstagewebhookID returns a boolean if a field has been set.

### GetFkiEzmaxpartnerproductstageID

`func (o *EzmaxpartnerproductstagewebhookResponse) GetFkiEzmaxpartnerproductstageID() int32`

GetFkiEzmaxpartnerproductstageID returns the FkiEzmaxpartnerproductstageID field if non-nil, zero value otherwise.

### GetFkiEzmaxpartnerproductstageIDOk

`func (o *EzmaxpartnerproductstagewebhookResponse) GetFkiEzmaxpartnerproductstageIDOk() (*int32, bool)`

GetFkiEzmaxpartnerproductstageIDOk returns a tuple with the FkiEzmaxpartnerproductstageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzmaxpartnerproductstageID

`func (o *EzmaxpartnerproductstagewebhookResponse) SetFkiEzmaxpartnerproductstageID(v int32)`

SetFkiEzmaxpartnerproductstageID sets FkiEzmaxpartnerproductstageID field to given value.


### GetSEzmaxpartnerproductstageCode

`func (o *EzmaxpartnerproductstagewebhookResponse) GetSEzmaxpartnerproductstageCode() string`

GetSEzmaxpartnerproductstageCode returns the SEzmaxpartnerproductstageCode field if non-nil, zero value otherwise.

### GetSEzmaxpartnerproductstageCodeOk

`func (o *EzmaxpartnerproductstagewebhookResponse) GetSEzmaxpartnerproductstageCodeOk() (*string, bool)`

GetSEzmaxpartnerproductstageCodeOk returns a tuple with the SEzmaxpartnerproductstageCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzmaxpartnerproductstageCode

`func (o *EzmaxpartnerproductstagewebhookResponse) SetSEzmaxpartnerproductstageCode(v string)`

SetSEzmaxpartnerproductstageCode sets SEzmaxpartnerproductstageCode field to given value.

### HasSEzmaxpartnerproductstageCode

`func (o *EzmaxpartnerproductstagewebhookResponse) HasSEzmaxpartnerproductstageCode() bool`

HasSEzmaxpartnerproductstageCode returns a boolean if a field has been set.

### GetEEzmaxpartnerproductstagewebhookModule

`func (o *EzmaxpartnerproductstagewebhookResponse) GetEEzmaxpartnerproductstagewebhookModule() FieldEEzmaxpartnerproductstagewebhookModule`

GetEEzmaxpartnerproductstagewebhookModule returns the EEzmaxpartnerproductstagewebhookModule field if non-nil, zero value otherwise.

### GetEEzmaxpartnerproductstagewebhookModuleOk

`func (o *EzmaxpartnerproductstagewebhookResponse) GetEEzmaxpartnerproductstagewebhookModuleOk() (*FieldEEzmaxpartnerproductstagewebhookModule, bool)`

GetEEzmaxpartnerproductstagewebhookModuleOk returns a tuple with the EEzmaxpartnerproductstagewebhookModule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzmaxpartnerproductstagewebhookModule

`func (o *EzmaxpartnerproductstagewebhookResponse) SetEEzmaxpartnerproductstagewebhookModule(v FieldEEzmaxpartnerproductstagewebhookModule)`

SetEEzmaxpartnerproductstagewebhookModule sets EEzmaxpartnerproductstagewebhookModule field to given value.


### GetEEzmaxpartnerproductstagewebhookEzmaxpartnerevent

`func (o *EzmaxpartnerproductstagewebhookResponse) GetEEzmaxpartnerproductstagewebhookEzmaxpartnerevent() FieldEEzmaxpartnerproductstagewebhookEzmaxpartnerevent`

GetEEzmaxpartnerproductstagewebhookEzmaxpartnerevent returns the EEzmaxpartnerproductstagewebhookEzmaxpartnerevent field if non-nil, zero value otherwise.

### GetEEzmaxpartnerproductstagewebhookEzmaxpartnereventOk

`func (o *EzmaxpartnerproductstagewebhookResponse) GetEEzmaxpartnerproductstagewebhookEzmaxpartnereventOk() (*FieldEEzmaxpartnerproductstagewebhookEzmaxpartnerevent, bool)`

GetEEzmaxpartnerproductstagewebhookEzmaxpartnereventOk returns a tuple with the EEzmaxpartnerproductstagewebhookEzmaxpartnerevent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzmaxpartnerproductstagewebhookEzmaxpartnerevent

`func (o *EzmaxpartnerproductstagewebhookResponse) SetEEzmaxpartnerproductstagewebhookEzmaxpartnerevent(v FieldEEzmaxpartnerproductstagewebhookEzmaxpartnerevent)`

SetEEzmaxpartnerproductstagewebhookEzmaxpartnerevent sets EEzmaxpartnerproductstagewebhookEzmaxpartnerevent field to given value.

### HasEEzmaxpartnerproductstagewebhookEzmaxpartnerevent

`func (o *EzmaxpartnerproductstagewebhookResponse) HasEEzmaxpartnerproductstagewebhookEzmaxpartnerevent() bool`

HasEEzmaxpartnerproductstagewebhookEzmaxpartnerevent returns a boolean if a field has been set.

### GetEEzmaxpartnerproductstagewebhookEzsignevent

`func (o *EzmaxpartnerproductstagewebhookResponse) GetEEzmaxpartnerproductstagewebhookEzsignevent() FieldEEzmaxpartnerproductstagewebhookEzsignevent`

GetEEzmaxpartnerproductstagewebhookEzsignevent returns the EEzmaxpartnerproductstagewebhookEzsignevent field if non-nil, zero value otherwise.

### GetEEzmaxpartnerproductstagewebhookEzsigneventOk

`func (o *EzmaxpartnerproductstagewebhookResponse) GetEEzmaxpartnerproductstagewebhookEzsigneventOk() (*FieldEEzmaxpartnerproductstagewebhookEzsignevent, bool)`

GetEEzmaxpartnerproductstagewebhookEzsigneventOk returns a tuple with the EEzmaxpartnerproductstagewebhookEzsignevent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzmaxpartnerproductstagewebhookEzsignevent

`func (o *EzmaxpartnerproductstagewebhookResponse) SetEEzmaxpartnerproductstagewebhookEzsignevent(v FieldEEzmaxpartnerproductstagewebhookEzsignevent)`

SetEEzmaxpartnerproductstagewebhookEzsignevent sets EEzmaxpartnerproductstagewebhookEzsignevent field to given value.

### HasEEzmaxpartnerproductstagewebhookEzsignevent

`func (o *EzmaxpartnerproductstagewebhookResponse) HasEEzmaxpartnerproductstagewebhookEzsignevent() bool`

HasEEzmaxpartnerproductstagewebhookEzsignevent returns a boolean if a field has been set.

### GetEEzmaxpartnerproductstagewebhookManagementevent

`func (o *EzmaxpartnerproductstagewebhookResponse) GetEEzmaxpartnerproductstagewebhookManagementevent() FieldEEzmaxpartnerproductstagewebhookManagementevent`

GetEEzmaxpartnerproductstagewebhookManagementevent returns the EEzmaxpartnerproductstagewebhookManagementevent field if non-nil, zero value otherwise.

### GetEEzmaxpartnerproductstagewebhookManagementeventOk

`func (o *EzmaxpartnerproductstagewebhookResponse) GetEEzmaxpartnerproductstagewebhookManagementeventOk() (*FieldEEzmaxpartnerproductstagewebhookManagementevent, bool)`

GetEEzmaxpartnerproductstagewebhookManagementeventOk returns a tuple with the EEzmaxpartnerproductstagewebhookManagementevent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzmaxpartnerproductstagewebhookManagementevent

`func (o *EzmaxpartnerproductstagewebhookResponse) SetEEzmaxpartnerproductstagewebhookManagementevent(v FieldEEzmaxpartnerproductstagewebhookManagementevent)`

SetEEzmaxpartnerproductstagewebhookManagementevent sets EEzmaxpartnerproductstagewebhookManagementevent field to given value.

### HasEEzmaxpartnerproductstagewebhookManagementevent

`func (o *EzmaxpartnerproductstagewebhookResponse) HasEEzmaxpartnerproductstagewebhookManagementevent() bool`

HasEEzmaxpartnerproductstagewebhookManagementevent returns a boolean if a field has been set.

### GetEEzmaxpartnerproductstagewebhookRealestateevent

`func (o *EzmaxpartnerproductstagewebhookResponse) GetEEzmaxpartnerproductstagewebhookRealestateevent() FieldEEzmaxpartnerproductstagewebhookRealestateevent`

GetEEzmaxpartnerproductstagewebhookRealestateevent returns the EEzmaxpartnerproductstagewebhookRealestateevent field if non-nil, zero value otherwise.

### GetEEzmaxpartnerproductstagewebhookRealestateeventOk

`func (o *EzmaxpartnerproductstagewebhookResponse) GetEEzmaxpartnerproductstagewebhookRealestateeventOk() (*FieldEEzmaxpartnerproductstagewebhookRealestateevent, bool)`

GetEEzmaxpartnerproductstagewebhookRealestateeventOk returns a tuple with the EEzmaxpartnerproductstagewebhookRealestateevent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzmaxpartnerproductstagewebhookRealestateevent

`func (o *EzmaxpartnerproductstagewebhookResponse) SetEEzmaxpartnerproductstagewebhookRealestateevent(v FieldEEzmaxpartnerproductstagewebhookRealestateevent)`

SetEEzmaxpartnerproductstagewebhookRealestateevent sets EEzmaxpartnerproductstagewebhookRealestateevent field to given value.

### HasEEzmaxpartnerproductstagewebhookRealestateevent

`func (o *EzmaxpartnerproductstagewebhookResponse) HasEEzmaxpartnerproductstagewebhookRealestateevent() bool`

HasEEzmaxpartnerproductstagewebhookRealestateevent returns a boolean if a field has been set.

### GetSEzmaxpartnerproductstagewebhookDescription

`func (o *EzmaxpartnerproductstagewebhookResponse) GetSEzmaxpartnerproductstagewebhookDescription() string`

GetSEzmaxpartnerproductstagewebhookDescription returns the SEzmaxpartnerproductstagewebhookDescription field if non-nil, zero value otherwise.

### GetSEzmaxpartnerproductstagewebhookDescriptionOk

`func (o *EzmaxpartnerproductstagewebhookResponse) GetSEzmaxpartnerproductstagewebhookDescriptionOk() (*string, bool)`

GetSEzmaxpartnerproductstagewebhookDescriptionOk returns a tuple with the SEzmaxpartnerproductstagewebhookDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzmaxpartnerproductstagewebhookDescription

`func (o *EzmaxpartnerproductstagewebhookResponse) SetSEzmaxpartnerproductstagewebhookDescription(v string)`

SetSEzmaxpartnerproductstagewebhookDescription sets SEzmaxpartnerproductstagewebhookDescription field to given value.

### HasSEzmaxpartnerproductstagewebhookDescription

`func (o *EzmaxpartnerproductstagewebhookResponse) HasSEzmaxpartnerproductstagewebhookDescription() bool`

HasSEzmaxpartnerproductstagewebhookDescription returns a boolean if a field has been set.

### GetSEzmaxpartnerproductstagewebhookUrl

`func (o *EzmaxpartnerproductstagewebhookResponse) GetSEzmaxpartnerproductstagewebhookUrl() string`

GetSEzmaxpartnerproductstagewebhookUrl returns the SEzmaxpartnerproductstagewebhookUrl field if non-nil, zero value otherwise.

### GetSEzmaxpartnerproductstagewebhookUrlOk

`func (o *EzmaxpartnerproductstagewebhookResponse) GetSEzmaxpartnerproductstagewebhookUrlOk() (*string, bool)`

GetSEzmaxpartnerproductstagewebhookUrlOk returns a tuple with the SEzmaxpartnerproductstagewebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzmaxpartnerproductstagewebhookUrl

`func (o *EzmaxpartnerproductstagewebhookResponse) SetSEzmaxpartnerproductstagewebhookUrl(v string)`

SetSEzmaxpartnerproductstagewebhookUrl sets SEzmaxpartnerproductstagewebhookUrl field to given value.

### HasSEzmaxpartnerproductstagewebhookUrl

`func (o *EzmaxpartnerproductstagewebhookResponse) HasSEzmaxpartnerproductstagewebhookUrl() bool`

HasSEzmaxpartnerproductstagewebhookUrl returns a boolean if a field has been set.

### GetSEzmaxpartnerproductstagewebhookEmailfailed

`func (o *EzmaxpartnerproductstagewebhookResponse) GetSEzmaxpartnerproductstagewebhookEmailfailed() string`

GetSEzmaxpartnerproductstagewebhookEmailfailed returns the SEzmaxpartnerproductstagewebhookEmailfailed field if non-nil, zero value otherwise.

### GetSEzmaxpartnerproductstagewebhookEmailfailedOk

`func (o *EzmaxpartnerproductstagewebhookResponse) GetSEzmaxpartnerproductstagewebhookEmailfailedOk() (*string, bool)`

GetSEzmaxpartnerproductstagewebhookEmailfailedOk returns a tuple with the SEzmaxpartnerproductstagewebhookEmailfailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzmaxpartnerproductstagewebhookEmailfailed

`func (o *EzmaxpartnerproductstagewebhookResponse) SetSEzmaxpartnerproductstagewebhookEmailfailed(v string)`

SetSEzmaxpartnerproductstagewebhookEmailfailed sets SEzmaxpartnerproductstagewebhookEmailfailed field to given value.

### HasSEzmaxpartnerproductstagewebhookEmailfailed

`func (o *EzmaxpartnerproductstagewebhookResponse) HasSEzmaxpartnerproductstagewebhookEmailfailed() bool`

HasSEzmaxpartnerproductstagewebhookEmailfailed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


