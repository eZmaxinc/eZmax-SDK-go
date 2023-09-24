# ModuleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiModuleID** | **int32** | The unique ID of the Module | 
**FkiModulegroupID** | **int32** | The unique ID of the Modulegroup | 
**EModuleInternalname** | **string** | The Internal name of the Module.  This is theoretically an enum field but there are so many possibles values we decided not to list them all. | 
**SModuleNameX** | **string** | The Name of the Module in the language of the requester | 
**BModuleRegistered** | **bool** | Whether the Module is registered or not | 
**BModuleRegisteredapi** | **bool** | Whether the Module is registered or not for api use | 

## Methods

### NewModuleResponse

`func NewModuleResponse(pkiModuleID int32, fkiModulegroupID int32, eModuleInternalname string, sModuleNameX string, bModuleRegistered bool, bModuleRegisteredapi bool, ) *ModuleResponse`

NewModuleResponse instantiates a new ModuleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModuleResponseWithDefaults

`func NewModuleResponseWithDefaults() *ModuleResponse`

NewModuleResponseWithDefaults instantiates a new ModuleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiModuleID

`func (o *ModuleResponse) GetPkiModuleID() int32`

GetPkiModuleID returns the PkiModuleID field if non-nil, zero value otherwise.

### GetPkiModuleIDOk

`func (o *ModuleResponse) GetPkiModuleIDOk() (*int32, bool)`

GetPkiModuleIDOk returns a tuple with the PkiModuleID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiModuleID

`func (o *ModuleResponse) SetPkiModuleID(v int32)`

SetPkiModuleID sets PkiModuleID field to given value.


### GetFkiModulegroupID

`func (o *ModuleResponse) GetFkiModulegroupID() int32`

GetFkiModulegroupID returns the FkiModulegroupID field if non-nil, zero value otherwise.

### GetFkiModulegroupIDOk

`func (o *ModuleResponse) GetFkiModulegroupIDOk() (*int32, bool)`

GetFkiModulegroupIDOk returns a tuple with the FkiModulegroupID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiModulegroupID

`func (o *ModuleResponse) SetFkiModulegroupID(v int32)`

SetFkiModulegroupID sets FkiModulegroupID field to given value.


### GetEModuleInternalname

`func (o *ModuleResponse) GetEModuleInternalname() string`

GetEModuleInternalname returns the EModuleInternalname field if non-nil, zero value otherwise.

### GetEModuleInternalnameOk

`func (o *ModuleResponse) GetEModuleInternalnameOk() (*string, bool)`

GetEModuleInternalnameOk returns a tuple with the EModuleInternalname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEModuleInternalname

`func (o *ModuleResponse) SetEModuleInternalname(v string)`

SetEModuleInternalname sets EModuleInternalname field to given value.


### GetSModuleNameX

`func (o *ModuleResponse) GetSModuleNameX() string`

GetSModuleNameX returns the SModuleNameX field if non-nil, zero value otherwise.

### GetSModuleNameXOk

`func (o *ModuleResponse) GetSModuleNameXOk() (*string, bool)`

GetSModuleNameXOk returns a tuple with the SModuleNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSModuleNameX

`func (o *ModuleResponse) SetSModuleNameX(v string)`

SetSModuleNameX sets SModuleNameX field to given value.


### GetBModuleRegistered

`func (o *ModuleResponse) GetBModuleRegistered() bool`

GetBModuleRegistered returns the BModuleRegistered field if non-nil, zero value otherwise.

### GetBModuleRegisteredOk

`func (o *ModuleResponse) GetBModuleRegisteredOk() (*bool, bool)`

GetBModuleRegisteredOk returns a tuple with the BModuleRegistered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBModuleRegistered

`func (o *ModuleResponse) SetBModuleRegistered(v bool)`

SetBModuleRegistered sets BModuleRegistered field to given value.


### GetBModuleRegisteredapi

`func (o *ModuleResponse) GetBModuleRegisteredapi() bool`

GetBModuleRegisteredapi returns the BModuleRegisteredapi field if non-nil, zero value otherwise.

### GetBModuleRegisteredapiOk

`func (o *ModuleResponse) GetBModuleRegisteredapiOk() (*bool, bool)`

GetBModuleRegisteredapiOk returns a tuple with the BModuleRegisteredapi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBModuleRegisteredapi

`func (o *ModuleResponse) SetBModuleRegisteredapi(v bool)`

SetBModuleRegisteredapi sets BModuleRegisteredapi field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


