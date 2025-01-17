# VersionhistoryResponseCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiVersionhistoryID** | **int32** | The unique ID of the Versionhistory | 
**FkiModuleID** | Pointer to **int32** | The unique ID of the Module | [optional] 
**FkiModulesectionID** | Pointer to **int32** | The unique ID of the Modulesection | [optional] 
**SModuleNameX** | Pointer to **string** | The Name of the Module in the language of the requester | [optional] 
**SModulesectionNameX** | Pointer to **string** | The Name of the Modulesection in the language of the requester | [optional] 
**EVersionhistoryUsertype** | Pointer to [**FieldEVersionhistoryUsertype**](FieldEVersionhistoryUsertype.md) |  | [optional] 
**ObjVersionhistoryDetail** | [**MultilingualVersionhistoryDetail**](MultilingualVersionhistoryDetail.md) |  | 
**DtVersionhistoryDate** | **string** | The date  at which the Versionhistory was published or should be published | 
**DtVersionhistoryDateend** | Pointer to **string** | The date  at which the Versionhistory will no longer be visible | [optional] 
**EVersionhistoryType** | [**FieldEVersionhistoryType**](FieldEVersionhistoryType.md) |  | 
**BVersionhistoryDraft** | **bool** | Whether the Versionhistory is published or still a draft | 

## Methods

### NewVersionhistoryResponseCompound

`func NewVersionhistoryResponseCompound(pkiVersionhistoryID int32, objVersionhistoryDetail MultilingualVersionhistoryDetail, dtVersionhistoryDate string, eVersionhistoryType FieldEVersionhistoryType, bVersionhistoryDraft bool, ) *VersionhistoryResponseCompound`

NewVersionhistoryResponseCompound instantiates a new VersionhistoryResponseCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionhistoryResponseCompoundWithDefaults

`func NewVersionhistoryResponseCompoundWithDefaults() *VersionhistoryResponseCompound`

NewVersionhistoryResponseCompoundWithDefaults instantiates a new VersionhistoryResponseCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiVersionhistoryID

`func (o *VersionhistoryResponseCompound) GetPkiVersionhistoryID() int32`

GetPkiVersionhistoryID returns the PkiVersionhistoryID field if non-nil, zero value otherwise.

### GetPkiVersionhistoryIDOk

`func (o *VersionhistoryResponseCompound) GetPkiVersionhistoryIDOk() (*int32, bool)`

GetPkiVersionhistoryIDOk returns a tuple with the PkiVersionhistoryID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiVersionhistoryID

`func (o *VersionhistoryResponseCompound) SetPkiVersionhistoryID(v int32)`

SetPkiVersionhistoryID sets PkiVersionhistoryID field to given value.


### GetFkiModuleID

`func (o *VersionhistoryResponseCompound) GetFkiModuleID() int32`

GetFkiModuleID returns the FkiModuleID field if non-nil, zero value otherwise.

### GetFkiModuleIDOk

`func (o *VersionhistoryResponseCompound) GetFkiModuleIDOk() (*int32, bool)`

GetFkiModuleIDOk returns a tuple with the FkiModuleID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiModuleID

`func (o *VersionhistoryResponseCompound) SetFkiModuleID(v int32)`

SetFkiModuleID sets FkiModuleID field to given value.

### HasFkiModuleID

`func (o *VersionhistoryResponseCompound) HasFkiModuleID() bool`

HasFkiModuleID returns a boolean if a field has been set.

### GetFkiModulesectionID

`func (o *VersionhistoryResponseCompound) GetFkiModulesectionID() int32`

GetFkiModulesectionID returns the FkiModulesectionID field if non-nil, zero value otherwise.

### GetFkiModulesectionIDOk

`func (o *VersionhistoryResponseCompound) GetFkiModulesectionIDOk() (*int32, bool)`

GetFkiModulesectionIDOk returns a tuple with the FkiModulesectionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiModulesectionID

`func (o *VersionhistoryResponseCompound) SetFkiModulesectionID(v int32)`

SetFkiModulesectionID sets FkiModulesectionID field to given value.

### HasFkiModulesectionID

`func (o *VersionhistoryResponseCompound) HasFkiModulesectionID() bool`

HasFkiModulesectionID returns a boolean if a field has been set.

### GetSModuleNameX

`func (o *VersionhistoryResponseCompound) GetSModuleNameX() string`

GetSModuleNameX returns the SModuleNameX field if non-nil, zero value otherwise.

### GetSModuleNameXOk

`func (o *VersionhistoryResponseCompound) GetSModuleNameXOk() (*string, bool)`

GetSModuleNameXOk returns a tuple with the SModuleNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSModuleNameX

`func (o *VersionhistoryResponseCompound) SetSModuleNameX(v string)`

SetSModuleNameX sets SModuleNameX field to given value.

### HasSModuleNameX

`func (o *VersionhistoryResponseCompound) HasSModuleNameX() bool`

HasSModuleNameX returns a boolean if a field has been set.

### GetSModulesectionNameX

`func (o *VersionhistoryResponseCompound) GetSModulesectionNameX() string`

GetSModulesectionNameX returns the SModulesectionNameX field if non-nil, zero value otherwise.

### GetSModulesectionNameXOk

`func (o *VersionhistoryResponseCompound) GetSModulesectionNameXOk() (*string, bool)`

GetSModulesectionNameXOk returns a tuple with the SModulesectionNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSModulesectionNameX

`func (o *VersionhistoryResponseCompound) SetSModulesectionNameX(v string)`

SetSModulesectionNameX sets SModulesectionNameX field to given value.

### HasSModulesectionNameX

`func (o *VersionhistoryResponseCompound) HasSModulesectionNameX() bool`

HasSModulesectionNameX returns a boolean if a field has been set.

### GetEVersionhistoryUsertype

`func (o *VersionhistoryResponseCompound) GetEVersionhistoryUsertype() FieldEVersionhistoryUsertype`

GetEVersionhistoryUsertype returns the EVersionhistoryUsertype field if non-nil, zero value otherwise.

### GetEVersionhistoryUsertypeOk

`func (o *VersionhistoryResponseCompound) GetEVersionhistoryUsertypeOk() (*FieldEVersionhistoryUsertype, bool)`

GetEVersionhistoryUsertypeOk returns a tuple with the EVersionhistoryUsertype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEVersionhistoryUsertype

`func (o *VersionhistoryResponseCompound) SetEVersionhistoryUsertype(v FieldEVersionhistoryUsertype)`

SetEVersionhistoryUsertype sets EVersionhistoryUsertype field to given value.

### HasEVersionhistoryUsertype

`func (o *VersionhistoryResponseCompound) HasEVersionhistoryUsertype() bool`

HasEVersionhistoryUsertype returns a boolean if a field has been set.

### GetObjVersionhistoryDetail

`func (o *VersionhistoryResponseCompound) GetObjVersionhistoryDetail() MultilingualVersionhistoryDetail`

GetObjVersionhistoryDetail returns the ObjVersionhistoryDetail field if non-nil, zero value otherwise.

### GetObjVersionhistoryDetailOk

`func (o *VersionhistoryResponseCompound) GetObjVersionhistoryDetailOk() (*MultilingualVersionhistoryDetail, bool)`

GetObjVersionhistoryDetailOk returns a tuple with the ObjVersionhistoryDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjVersionhistoryDetail

`func (o *VersionhistoryResponseCompound) SetObjVersionhistoryDetail(v MultilingualVersionhistoryDetail)`

SetObjVersionhistoryDetail sets ObjVersionhistoryDetail field to given value.


### GetDtVersionhistoryDate

`func (o *VersionhistoryResponseCompound) GetDtVersionhistoryDate() string`

GetDtVersionhistoryDate returns the DtVersionhistoryDate field if non-nil, zero value otherwise.

### GetDtVersionhistoryDateOk

`func (o *VersionhistoryResponseCompound) GetDtVersionhistoryDateOk() (*string, bool)`

GetDtVersionhistoryDateOk returns a tuple with the DtVersionhistoryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtVersionhistoryDate

`func (o *VersionhistoryResponseCompound) SetDtVersionhistoryDate(v string)`

SetDtVersionhistoryDate sets DtVersionhistoryDate field to given value.


### GetDtVersionhistoryDateend

`func (o *VersionhistoryResponseCompound) GetDtVersionhistoryDateend() string`

GetDtVersionhistoryDateend returns the DtVersionhistoryDateend field if non-nil, zero value otherwise.

### GetDtVersionhistoryDateendOk

`func (o *VersionhistoryResponseCompound) GetDtVersionhistoryDateendOk() (*string, bool)`

GetDtVersionhistoryDateendOk returns a tuple with the DtVersionhistoryDateend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtVersionhistoryDateend

`func (o *VersionhistoryResponseCompound) SetDtVersionhistoryDateend(v string)`

SetDtVersionhistoryDateend sets DtVersionhistoryDateend field to given value.

### HasDtVersionhistoryDateend

`func (o *VersionhistoryResponseCompound) HasDtVersionhistoryDateend() bool`

HasDtVersionhistoryDateend returns a boolean if a field has been set.

### GetEVersionhistoryType

`func (o *VersionhistoryResponseCompound) GetEVersionhistoryType() FieldEVersionhistoryType`

GetEVersionhistoryType returns the EVersionhistoryType field if non-nil, zero value otherwise.

### GetEVersionhistoryTypeOk

`func (o *VersionhistoryResponseCompound) GetEVersionhistoryTypeOk() (*FieldEVersionhistoryType, bool)`

GetEVersionhistoryTypeOk returns a tuple with the EVersionhistoryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEVersionhistoryType

`func (o *VersionhistoryResponseCompound) SetEVersionhistoryType(v FieldEVersionhistoryType)`

SetEVersionhistoryType sets EVersionhistoryType field to given value.


### GetBVersionhistoryDraft

`func (o *VersionhistoryResponseCompound) GetBVersionhistoryDraft() bool`

GetBVersionhistoryDraft returns the BVersionhistoryDraft field if non-nil, zero value otherwise.

### GetBVersionhistoryDraftOk

`func (o *VersionhistoryResponseCompound) GetBVersionhistoryDraftOk() (*bool, bool)`

GetBVersionhistoryDraftOk returns a tuple with the BVersionhistoryDraft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBVersionhistoryDraft

`func (o *VersionhistoryResponseCompound) SetBVersionhistoryDraft(v bool)`

SetBVersionhistoryDraft sets BVersionhistoryDraft field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


