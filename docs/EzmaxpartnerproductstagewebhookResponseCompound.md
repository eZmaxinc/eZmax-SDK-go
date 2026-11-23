# EzmaxpartnerproductstagewebhookResponseCompound

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

### NewEzmaxpartnerproductstagewebhookResponseCompound

`func NewEzmaxpartnerproductstagewebhookResponseCompound(fkiEzmaxpartnerproductstageID int32, eEzmaxpartnerproductstagewebhookModule FieldEEzmaxpartnerproductstagewebhookModule, ) *EzmaxpartnerproductstagewebhookResponseCompound`

NewEzmaxpartnerproductstagewebhookResponseCompound instantiates a new EzmaxpartnerproductstagewebhookResponseCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzmaxpartnerproductstagewebhookResponseCompoundWithDefaults

`func NewEzmaxpartnerproductstagewebhookResponseCompoundWithDefaults() *EzmaxpartnerproductstagewebhookResponseCompound`

NewEzmaxpartnerproductstagewebhookResponseCompoundWithDefaults instantiates a new EzmaxpartnerproductstagewebhookResponseCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzmaxpartnerproductstagewebhookID

`func (o *EzmaxpartnerproductstagewebhookResponseCompound) GetPkiEzmaxpartnerproductstagewebhookID() int32`

GetPkiEzmaxpartnerproductstagewebhookID returns the PkiEzmaxpartnerproductstagewebhookID field if non-nil, zero value otherwise.

### GetPkiEzmaxpartnerproductstagewebhookIDOk

`func (o *EzmaxpartnerproductstagewebhookResponseCompound) GetPkiEzmaxpartnerproductstagewebhookIDOk() (*int32, bool)`

GetPkiEzmaxpartnerproductstagewebhookIDOk returns a tuple with the PkiEzmaxpartnerproductstagewebhookID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzmaxpartnerproductstagewebhookID

`func (o *EzmaxpartnerproductstagewebhookResponseCompound) SetPkiEzmaxpartnerproductstagewebhookID(v int32)`

SetPkiEzmaxpartnerproductstagewebhookID sets PkiEzmaxpartnerproductstagewebhookID field to given value.

### HasPkiEzmaxpartnerproductstagewebhookID

`func (o *EzmaxpartnerproductstagewebhookResponseCompound) HasPkiEzmaxpartnerproductstagewebhookID() bool`

HasPkiEzmaxpartnerproductstagewebhookID returns a boolean if a field has been set.

### GetFkiEzmaxpartnerproductstageID

`func (o *EzmaxpartnerproductstagewebhookResponseCompound) GetFkiEzmaxpartnerproductstageID() int32`

GetFkiEzmaxpartnerproductstageID returns the FkiEzmaxpartnerproductstageID field if non-nil, zero value otherwise.

### GetFkiEzmaxpartnerproductstageIDOk

`func (o *EzmaxpartnerproductstagewebhookResponseCompound) GetFkiEzmaxpartnerproductstageIDOk() (*int32, bool)`

GetFkiEzmaxpartnerproductstageIDOk returns a tuple with the FkiEzmaxpartnerproductstageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzmaxpartnerproductstageID

`func (o *EzmaxpartnerproductstagewebhookResponseCompound) SetFkiEzmaxpartnerproductstageID(v int32)`

SetFkiEzmaxpartnerproductstageID sets FkiEzmaxpartnerproductstageID field to given value.


### GetSEzmaxpartnerproductstageCode

`func (o *EzmaxpartnerproductstagewebhookResponseCompound) GetSEzmaxpartnerproductstageCode() string`

GetSEzmaxpartnerproductstageCode returns the SEzmaxpartnerproductstageCode field if non-nil, zero value otherwise.

### GetSEzmaxpartnerproductstageCodeOk

`func (o *EzmaxpartnerproductstagewebhookResponseCompound) GetSEzmaxpartnerproductstageCodeOk() (*string, bool)`

GetSEzmaxpartnerproductstageCodeOk returns a tuple with the SEzmaxpartnerproductstageCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzmaxpartnerproductstageCode

`func (o *EzmaxpartnerproductstagewebhookResponseCompound) SetSEzmaxpartnerproductstageCode(v string)`

SetSEzmaxpartnerproductstageCode sets SEzmaxpartnerproductstageCode field to given value.

### HasSEzmaxpartnerproductstageCode

`func (o *EzmaxpartnerproductstagewebhookResponseCompound) HasSEzmaxpartnerproductstageCode() bool`

HasSEzmaxpartnerproductstageCode returns a boolean if a field has been set.

### GetEEzmaxpartnerproductstagewebhookModule

`func (o *EzmaxpartnerproductstagewebhookResponseCompound) GetEEzmaxpartnerproductstagewebhookModule() FieldEEzmaxpartnerproductstagewebhookModule`

GetEEzmaxpartnerproductstagewebhookModule returns the EEzmaxpartnerproductstagewebhookModule field if non-nil, zero value otherwise.

### GetEEzmaxpartnerproductstagewebhookModuleOk

`func (o *EzmaxpartnerproductstagewebhookResponseCompound) GetEEzmaxpartnerproductstagewebhookModuleOk() (*FieldEEzmaxpartnerproductstagewebhookModule, bool)`

GetEEzmaxpartnerproductstagewebhookModuleOk returns a tuple with the EEzmaxpartnerproductstagewebhookModule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzmaxpartnerproductstagewebhookModule

`func (o *EzmaxpartnerproductstagewebhookResponseCompound) SetEEzmaxpartnerproductstagewebhookModule(v FieldEEzmaxpartnerproductstagewebhookModule)`

SetEEzmaxpartnerproductstagewebhookModule sets EEzmaxpartnerproductstagewebhookModule field to given value.


### GetEEzmaxpartnerproductstagewebhookEzmaxpartnerevent

`func (o *EzmaxpartnerproductstagewebhookResponseCompound) GetEEzmaxpartnerproductstagewebhookEzmaxpartnerevent() FieldEEzmaxpartnerproductstagewebhookEzmaxpartnerevent`

GetEEzmaxpartnerproductstagewebhookEzmaxpartnerevent returns the EEzmaxpartnerproductstagewebhookEzmaxpartnerevent field if non-nil, zero value otherwise.

### GetEEzmaxpartnerproductstagewebhookEzmaxpartnereventOk

`func (o *EzmaxpartnerproductstagewebhookResponseCompound) GetEEzmaxpartnerproductstagewebhookEzmaxpartnereventOk() (*FieldEEzmaxpartnerproductstagewebhookEzmaxpartnerevent, bool)`

GetEEzmaxpartnerproductstagewebhookEzmaxpartnereventOk returns a tuple with the EEzmaxpartnerproductstagewebhookEzmaxpartnerevent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzmaxpartnerproductstagewebhookEzmaxpartnerevent

`func (o *EzmaxpartnerproductstagewebhookResponseCompound) SetEEzmaxpartnerproductstagewebhookEzmaxpartnerevent(v FieldEEzmaxpartnerproductstagewebhookEzmaxpartnerevent)`

SetEEzmaxpartnerproductstagewebhookEzmaxpartnerevent sets EEzmaxpartnerproductstagewebhookEzmaxpartnerevent field to given value.

### HasEEzmaxpartnerproductstagewebhookEzmaxpartnerevent

`func (o *EzmaxpartnerproductstagewebhookResponseCompound) HasEEzmaxpartnerproductstagewebhookEzmaxpartnerevent() bool`

HasEEzmaxpartnerproductstagewebhookEzmaxpartnerevent returns a boolean if a field has been set.

### GetEEzmaxpartnerproductstagewebhookEzsignevent

`func (o *EzmaxpartnerproductstagewebhookResponseCompound) GetEEzmaxpartnerproductstagewebhookEzsignevent() FieldEEzmaxpartnerproductstagewebhookEzsignevent`

GetEEzmaxpartnerproductstagewebhookEzsignevent returns the EEzmaxpartnerproductstagewebhookEzsignevent field if non-nil, zero value otherwise.

### GetEEzmaxpartnerproductstagewebhookEzsigneventOk

`func (o *EzmaxpartnerproductstagewebhookResponseCompound) GetEEzmaxpartnerproductstagewebhookEzsigneventOk() (*FieldEEzmaxpartnerproductstagewebhookEzsignevent, bool)`

GetEEzmaxpartnerproductstagewebhookEzsigneventOk returns a tuple with the EEzmaxpartnerproductstagewebhookEzsignevent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzmaxpartnerproductstagewebhookEzsignevent

`func (o *EzmaxpartnerproductstagewebhookResponseCompound) SetEEzmaxpartnerproductstagewebhookEzsignevent(v FieldEEzmaxpartnerproductstagewebhookEzsignevent)`

SetEEzmaxpartnerproductstagewebhookEzsignevent sets EEzmaxpartnerproductstagewebhookEzsignevent field to given value.

### HasEEzmaxpartnerproductstagewebhookEzsignevent

`func (o *EzmaxpartnerproductstagewebhookResponseCompound) HasEEzmaxpartnerproductstagewebhookEzsignevent() bool`

HasEEzmaxpartnerproductstagewebhookEzsignevent returns a boolean if a field has been set.

### GetEEzmaxpartnerproductstagewebhookManagementevent

`func (o *EzmaxpartnerproductstagewebhookResponseCompound) GetEEzmaxpartnerproductstagewebhookManagementevent() FieldEEzmaxpartnerproductstagewebhookManagementevent`

GetEEzmaxpartnerproductstagewebhookManagementevent returns the EEzmaxpartnerproductstagewebhookManagementevent field if non-nil, zero value otherwise.

### GetEEzmaxpartnerproductstagewebhookManagementeventOk

`func (o *EzmaxpartnerproductstagewebhookResponseCompound) GetEEzmaxpartnerproductstagewebhookManagementeventOk() (*FieldEEzmaxpartnerproductstagewebhookManagementevent, bool)`

GetEEzmaxpartnerproductstagewebhookManagementeventOk returns a tuple with the EEzmaxpartnerproductstagewebhookManagementevent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzmaxpartnerproductstagewebhookManagementevent

`func (o *EzmaxpartnerproductstagewebhookResponseCompound) SetEEzmaxpartnerproductstagewebhookManagementevent(v FieldEEzmaxpartnerproductstagewebhookManagementevent)`

SetEEzmaxpartnerproductstagewebhookManagementevent sets EEzmaxpartnerproductstagewebhookManagementevent field to given value.

### HasEEzmaxpartnerproductstagewebhookManagementevent

`func (o *EzmaxpartnerproductstagewebhookResponseCompound) HasEEzmaxpartnerproductstagewebhookManagementevent() bool`

HasEEzmaxpartnerproductstagewebhookManagementevent returns a boolean if a field has been set.

### GetEEzmaxpartnerproductstagewebhookRealestateevent

`func (o *EzmaxpartnerproductstagewebhookResponseCompound) GetEEzmaxpartnerproductstagewebhookRealestateevent() FieldEEzmaxpartnerproductstagewebhookRealestateevent`

GetEEzmaxpartnerproductstagewebhookRealestateevent returns the EEzmaxpartnerproductstagewebhookRealestateevent field if non-nil, zero value otherwise.

### GetEEzmaxpartnerproductstagewebhookRealestateeventOk

`func (o *EzmaxpartnerproductstagewebhookResponseCompound) GetEEzmaxpartnerproductstagewebhookRealestateeventOk() (*FieldEEzmaxpartnerproductstagewebhookRealestateevent, bool)`

GetEEzmaxpartnerproductstagewebhookRealestateeventOk returns a tuple with the EEzmaxpartnerproductstagewebhookRealestateevent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzmaxpartnerproductstagewebhookRealestateevent

`func (o *EzmaxpartnerproductstagewebhookResponseCompound) SetEEzmaxpartnerproductstagewebhookRealestateevent(v FieldEEzmaxpartnerproductstagewebhookRealestateevent)`

SetEEzmaxpartnerproductstagewebhookRealestateevent sets EEzmaxpartnerproductstagewebhookRealestateevent field to given value.

### HasEEzmaxpartnerproductstagewebhookRealestateevent

`func (o *EzmaxpartnerproductstagewebhookResponseCompound) HasEEzmaxpartnerproductstagewebhookRealestateevent() bool`

HasEEzmaxpartnerproductstagewebhookRealestateevent returns a boolean if a field has been set.

### GetSEzmaxpartnerproductstagewebhookDescription

`func (o *EzmaxpartnerproductstagewebhookResponseCompound) GetSEzmaxpartnerproductstagewebhookDescription() string`

GetSEzmaxpartnerproductstagewebhookDescription returns the SEzmaxpartnerproductstagewebhookDescription field if non-nil, zero value otherwise.

### GetSEzmaxpartnerproductstagewebhookDescriptionOk

`func (o *EzmaxpartnerproductstagewebhookResponseCompound) GetSEzmaxpartnerproductstagewebhookDescriptionOk() (*string, bool)`

GetSEzmaxpartnerproductstagewebhookDescriptionOk returns a tuple with the SEzmaxpartnerproductstagewebhookDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzmaxpartnerproductstagewebhookDescription

`func (o *EzmaxpartnerproductstagewebhookResponseCompound) SetSEzmaxpartnerproductstagewebhookDescription(v string)`

SetSEzmaxpartnerproductstagewebhookDescription sets SEzmaxpartnerproductstagewebhookDescription field to given value.

### HasSEzmaxpartnerproductstagewebhookDescription

`func (o *EzmaxpartnerproductstagewebhookResponseCompound) HasSEzmaxpartnerproductstagewebhookDescription() bool`

HasSEzmaxpartnerproductstagewebhookDescription returns a boolean if a field has been set.

### GetSEzmaxpartnerproductstagewebhookUrl

`func (o *EzmaxpartnerproductstagewebhookResponseCompound) GetSEzmaxpartnerproductstagewebhookUrl() string`

GetSEzmaxpartnerproductstagewebhookUrl returns the SEzmaxpartnerproductstagewebhookUrl field if non-nil, zero value otherwise.

### GetSEzmaxpartnerproductstagewebhookUrlOk

`func (o *EzmaxpartnerproductstagewebhookResponseCompound) GetSEzmaxpartnerproductstagewebhookUrlOk() (*string, bool)`

GetSEzmaxpartnerproductstagewebhookUrlOk returns a tuple with the SEzmaxpartnerproductstagewebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzmaxpartnerproductstagewebhookUrl

`func (o *EzmaxpartnerproductstagewebhookResponseCompound) SetSEzmaxpartnerproductstagewebhookUrl(v string)`

SetSEzmaxpartnerproductstagewebhookUrl sets SEzmaxpartnerproductstagewebhookUrl field to given value.

### HasSEzmaxpartnerproductstagewebhookUrl

`func (o *EzmaxpartnerproductstagewebhookResponseCompound) HasSEzmaxpartnerproductstagewebhookUrl() bool`

HasSEzmaxpartnerproductstagewebhookUrl returns a boolean if a field has been set.

### GetSEzmaxpartnerproductstagewebhookEmailfailed

`func (o *EzmaxpartnerproductstagewebhookResponseCompound) GetSEzmaxpartnerproductstagewebhookEmailfailed() string`

GetSEzmaxpartnerproductstagewebhookEmailfailed returns the SEzmaxpartnerproductstagewebhookEmailfailed field if non-nil, zero value otherwise.

### GetSEzmaxpartnerproductstagewebhookEmailfailedOk

`func (o *EzmaxpartnerproductstagewebhookResponseCompound) GetSEzmaxpartnerproductstagewebhookEmailfailedOk() (*string, bool)`

GetSEzmaxpartnerproductstagewebhookEmailfailedOk returns a tuple with the SEzmaxpartnerproductstagewebhookEmailfailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzmaxpartnerproductstagewebhookEmailfailed

`func (o *EzmaxpartnerproductstagewebhookResponseCompound) SetSEzmaxpartnerproductstagewebhookEmailfailed(v string)`

SetSEzmaxpartnerproductstagewebhookEmailfailed sets SEzmaxpartnerproductstagewebhookEmailfailed field to given value.

### HasSEzmaxpartnerproductstagewebhookEmailfailed

`func (o *EzmaxpartnerproductstagewebhookResponseCompound) HasSEzmaxpartnerproductstagewebhookEmailfailed() bool`

HasSEzmaxpartnerproductstagewebhookEmailfailed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


