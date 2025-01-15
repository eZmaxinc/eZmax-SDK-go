# ModulesectionResponseCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiModulesectionID** | **int32** | The unique ID of the Modulesection | 
**FkiModuleID** | **int32** | The unique ID of the Module | 
**SModulesectionInternalname** | **string** | The Internal name of the Module section. | 
**SModulesectionNameX** | **string** | The Name of the Modulesection in the language of the requester | 
**AObjPermission** | Pointer to [**[]PermissionResponse**](PermissionResponse.md) |  | [optional] 

## Methods

### NewModulesectionResponseCompound

`func NewModulesectionResponseCompound(pkiModulesectionID int32, fkiModuleID int32, sModulesectionInternalname string, sModulesectionNameX string, ) *ModulesectionResponseCompound`

NewModulesectionResponseCompound instantiates a new ModulesectionResponseCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModulesectionResponseCompoundWithDefaults

`func NewModulesectionResponseCompoundWithDefaults() *ModulesectionResponseCompound`

NewModulesectionResponseCompoundWithDefaults instantiates a new ModulesectionResponseCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiModulesectionID

`func (o *ModulesectionResponseCompound) GetPkiModulesectionID() int32`

GetPkiModulesectionID returns the PkiModulesectionID field if non-nil, zero value otherwise.

### GetPkiModulesectionIDOk

`func (o *ModulesectionResponseCompound) GetPkiModulesectionIDOk() (*int32, bool)`

GetPkiModulesectionIDOk returns a tuple with the PkiModulesectionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiModulesectionID

`func (o *ModulesectionResponseCompound) SetPkiModulesectionID(v int32)`

SetPkiModulesectionID sets PkiModulesectionID field to given value.


### GetFkiModuleID

`func (o *ModulesectionResponseCompound) GetFkiModuleID() int32`

GetFkiModuleID returns the FkiModuleID field if non-nil, zero value otherwise.

### GetFkiModuleIDOk

`func (o *ModulesectionResponseCompound) GetFkiModuleIDOk() (*int32, bool)`

GetFkiModuleIDOk returns a tuple with the FkiModuleID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiModuleID

`func (o *ModulesectionResponseCompound) SetFkiModuleID(v int32)`

SetFkiModuleID sets FkiModuleID field to given value.


### GetSModulesectionInternalname

`func (o *ModulesectionResponseCompound) GetSModulesectionInternalname() string`

GetSModulesectionInternalname returns the SModulesectionInternalname field if non-nil, zero value otherwise.

### GetSModulesectionInternalnameOk

`func (o *ModulesectionResponseCompound) GetSModulesectionInternalnameOk() (*string, bool)`

GetSModulesectionInternalnameOk returns a tuple with the SModulesectionInternalname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSModulesectionInternalname

`func (o *ModulesectionResponseCompound) SetSModulesectionInternalname(v string)`

SetSModulesectionInternalname sets SModulesectionInternalname field to given value.


### GetSModulesectionNameX

`func (o *ModulesectionResponseCompound) GetSModulesectionNameX() string`

GetSModulesectionNameX returns the SModulesectionNameX field if non-nil, zero value otherwise.

### GetSModulesectionNameXOk

`func (o *ModulesectionResponseCompound) GetSModulesectionNameXOk() (*string, bool)`

GetSModulesectionNameXOk returns a tuple with the SModulesectionNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSModulesectionNameX

`func (o *ModulesectionResponseCompound) SetSModulesectionNameX(v string)`

SetSModulesectionNameX sets SModulesectionNameX field to given value.


### GetAObjPermission

`func (o *ModulesectionResponseCompound) GetAObjPermission() []PermissionResponseCompound`

GetAObjPermission returns the AObjPermission field if non-nil, zero value otherwise.

### GetAObjPermissionOk

`func (o *ModulesectionResponseCompound) GetAObjPermissionOk() (*[]PermissionResponseCompound, bool)`

GetAObjPermissionOk returns a tuple with the AObjPermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjPermission

`func (o *ModulesectionResponseCompound) SetAObjPermission(v []PermissionResponseCompound)`

SetAObjPermission sets AObjPermission field to given value.

### HasAObjPermission

`func (o *ModulesectionResponseCompound) HasAObjPermission() bool`

HasAObjPermission returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


