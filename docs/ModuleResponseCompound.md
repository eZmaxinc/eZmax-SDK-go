# ModuleResponseCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiModuleID** | **int32** | The unique ID of the Module | 
**FkiModulegroupID** | **int32** | The unique ID of the Modulegroup | 
**EModuleInternalname** | **string** | The Internal name of the Module.  This is theoretically an enum field but there are so many possibles values we decided not to list them all. | 
**SModuleNameX** | **string** | The Name of the Module in the language of the requester | 
**BModuleRegistered** | **bool** | Whether the Module is registered or not | 
**BModuleRegisteredapi** | **bool** | Whether the Module is registered or not for api use | 
**AObjModulesection** | Pointer to [**[]ModulesectionResponseCompound**](ModulesectionResponseCompound.md) |  | [optional] 

## Methods

### NewModuleResponseCompound

`func NewModuleResponseCompound(pkiModuleID int32, fkiModulegroupID int32, eModuleInternalname string, sModuleNameX string, bModuleRegistered bool, bModuleRegisteredapi bool, ) *ModuleResponseCompound`

NewModuleResponseCompound instantiates a new ModuleResponseCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModuleResponseCompoundWithDefaults

`func NewModuleResponseCompoundWithDefaults() *ModuleResponseCompound`

NewModuleResponseCompoundWithDefaults instantiates a new ModuleResponseCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiModuleID

`func (o *ModuleResponseCompound) GetPkiModuleID() int32`

GetPkiModuleID returns the PkiModuleID field if non-nil, zero value otherwise.

### GetPkiModuleIDOk

`func (o *ModuleResponseCompound) GetPkiModuleIDOk() (*int32, bool)`

GetPkiModuleIDOk returns a tuple with the PkiModuleID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiModuleID

`func (o *ModuleResponseCompound) SetPkiModuleID(v int32)`

SetPkiModuleID sets PkiModuleID field to given value.


### GetFkiModulegroupID

`func (o *ModuleResponseCompound) GetFkiModulegroupID() int32`

GetFkiModulegroupID returns the FkiModulegroupID field if non-nil, zero value otherwise.

### GetFkiModulegroupIDOk

`func (o *ModuleResponseCompound) GetFkiModulegroupIDOk() (*int32, bool)`

GetFkiModulegroupIDOk returns a tuple with the FkiModulegroupID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiModulegroupID

`func (o *ModuleResponseCompound) SetFkiModulegroupID(v int32)`

SetFkiModulegroupID sets FkiModulegroupID field to given value.


### GetEModuleInternalname

`func (o *ModuleResponseCompound) GetEModuleInternalname() string`

GetEModuleInternalname returns the EModuleInternalname field if non-nil, zero value otherwise.

### GetEModuleInternalnameOk

`func (o *ModuleResponseCompound) GetEModuleInternalnameOk() (*string, bool)`

GetEModuleInternalnameOk returns a tuple with the EModuleInternalname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEModuleInternalname

`func (o *ModuleResponseCompound) SetEModuleInternalname(v string)`

SetEModuleInternalname sets EModuleInternalname field to given value.


### GetSModuleNameX

`func (o *ModuleResponseCompound) GetSModuleNameX() string`

GetSModuleNameX returns the SModuleNameX field if non-nil, zero value otherwise.

### GetSModuleNameXOk

`func (o *ModuleResponseCompound) GetSModuleNameXOk() (*string, bool)`

GetSModuleNameXOk returns a tuple with the SModuleNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSModuleNameX

`func (o *ModuleResponseCompound) SetSModuleNameX(v string)`

SetSModuleNameX sets SModuleNameX field to given value.


### GetBModuleRegistered

`func (o *ModuleResponseCompound) GetBModuleRegistered() bool`

GetBModuleRegistered returns the BModuleRegistered field if non-nil, zero value otherwise.

### GetBModuleRegisteredOk

`func (o *ModuleResponseCompound) GetBModuleRegisteredOk() (*bool, bool)`

GetBModuleRegisteredOk returns a tuple with the BModuleRegistered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBModuleRegistered

`func (o *ModuleResponseCompound) SetBModuleRegistered(v bool)`

SetBModuleRegistered sets BModuleRegistered field to given value.


### GetBModuleRegisteredapi

`func (o *ModuleResponseCompound) GetBModuleRegisteredapi() bool`

GetBModuleRegisteredapi returns the BModuleRegisteredapi field if non-nil, zero value otherwise.

### GetBModuleRegisteredapiOk

`func (o *ModuleResponseCompound) GetBModuleRegisteredapiOk() (*bool, bool)`

GetBModuleRegisteredapiOk returns a tuple with the BModuleRegisteredapi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBModuleRegisteredapi

`func (o *ModuleResponseCompound) SetBModuleRegisteredapi(v bool)`

SetBModuleRegisteredapi sets BModuleRegisteredapi field to given value.


### GetAObjModulesection

`func (o *ModuleResponseCompound) GetAObjModulesection() []ModulesectionResponseCompound`

GetAObjModulesection returns the AObjModulesection field if non-nil, zero value otherwise.

### GetAObjModulesectionOk

`func (o *ModuleResponseCompound) GetAObjModulesectionOk() (*[]ModulesectionResponseCompound, bool)`

GetAObjModulesectionOk returns a tuple with the AObjModulesection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjModulesection

`func (o *ModuleResponseCompound) SetAObjModulesection(v []ModulesectionResponseCompound)`

SetAObjModulesection sets AObjModulesection field to given value.

### HasAObjModulesection

`func (o *ModuleResponseCompound) HasAObjModulesection() bool`

HasAObjModulesection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


