# ModulegroupResponseCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiModulegroupID** | **int32** | The unique ID of the Modulegroup | 
**SModulegroupNameX** | **string** | The name of the Modulegroup in the language of the requester | 
**AObjModule** | Pointer to [**[]ModuleResponseCompound**](ModuleResponseCompound.md) |  | [optional] 

## Methods

### NewModulegroupResponseCompound

`func NewModulegroupResponseCompound(pkiModulegroupID int32, sModulegroupNameX string, ) *ModulegroupResponseCompound`

NewModulegroupResponseCompound instantiates a new ModulegroupResponseCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModulegroupResponseCompoundWithDefaults

`func NewModulegroupResponseCompoundWithDefaults() *ModulegroupResponseCompound`

NewModulegroupResponseCompoundWithDefaults instantiates a new ModulegroupResponseCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiModulegroupID

`func (o *ModulegroupResponseCompound) GetPkiModulegroupID() int32`

GetPkiModulegroupID returns the PkiModulegroupID field if non-nil, zero value otherwise.

### GetPkiModulegroupIDOk

`func (o *ModulegroupResponseCompound) GetPkiModulegroupIDOk() (*int32, bool)`

GetPkiModulegroupIDOk returns a tuple with the PkiModulegroupID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiModulegroupID

`func (o *ModulegroupResponseCompound) SetPkiModulegroupID(v int32)`

SetPkiModulegroupID sets PkiModulegroupID field to given value.


### GetSModulegroupNameX

`func (o *ModulegroupResponseCompound) GetSModulegroupNameX() string`

GetSModulegroupNameX returns the SModulegroupNameX field if non-nil, zero value otherwise.

### GetSModulegroupNameXOk

`func (o *ModulegroupResponseCompound) GetSModulegroupNameXOk() (*string, bool)`

GetSModulegroupNameXOk returns a tuple with the SModulegroupNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSModulegroupNameX

`func (o *ModulegroupResponseCompound) SetSModulegroupNameX(v string)`

SetSModulegroupNameX sets SModulegroupNameX field to given value.


### GetAObjModule

`func (o *ModulegroupResponseCompound) GetAObjModule() []ModuleResponseCompound`

GetAObjModule returns the AObjModule field if non-nil, zero value otherwise.

### GetAObjModuleOk

`func (o *ModulegroupResponseCompound) GetAObjModuleOk() (*[]ModuleResponseCompound, bool)`

GetAObjModuleOk returns a tuple with the AObjModule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjModule

`func (o *ModulegroupResponseCompound) SetAObjModule(v []ModuleResponseCompound)`

SetAObjModule sets AObjModule field to given value.

### HasAObjModule

`func (o *ModulegroupResponseCompound) HasAObjModule() bool`

HasAObjModule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


