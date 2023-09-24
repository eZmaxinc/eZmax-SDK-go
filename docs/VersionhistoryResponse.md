# VersionhistoryResponse

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

### NewVersionhistoryResponse

`func NewVersionhistoryResponse(pkiVersionhistoryID int32, objVersionhistoryDetail MultilingualVersionhistoryDetail, dtVersionhistoryDate string, eVersionhistoryType FieldEVersionhistoryType, bVersionhistoryDraft bool, ) *VersionhistoryResponse`

NewVersionhistoryResponse instantiates a new VersionhistoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionhistoryResponseWithDefaults

`func NewVersionhistoryResponseWithDefaults() *VersionhistoryResponse`

NewVersionhistoryResponseWithDefaults instantiates a new VersionhistoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiVersionhistoryID

`func (o *VersionhistoryResponse) GetPkiVersionhistoryID() int32`

GetPkiVersionhistoryID returns the PkiVersionhistoryID field if non-nil, zero value otherwise.

### GetPkiVersionhistoryIDOk

`func (o *VersionhistoryResponse) GetPkiVersionhistoryIDOk() (*int32, bool)`

GetPkiVersionhistoryIDOk returns a tuple with the PkiVersionhistoryID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiVersionhistoryID

`func (o *VersionhistoryResponse) SetPkiVersionhistoryID(v int32)`

SetPkiVersionhistoryID sets PkiVersionhistoryID field to given value.


### GetFkiModuleID

`func (o *VersionhistoryResponse) GetFkiModuleID() int32`

GetFkiModuleID returns the FkiModuleID field if non-nil, zero value otherwise.

### GetFkiModuleIDOk

`func (o *VersionhistoryResponse) GetFkiModuleIDOk() (*int32, bool)`

GetFkiModuleIDOk returns a tuple with the FkiModuleID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiModuleID

`func (o *VersionhistoryResponse) SetFkiModuleID(v int32)`

SetFkiModuleID sets FkiModuleID field to given value.

### HasFkiModuleID

`func (o *VersionhistoryResponse) HasFkiModuleID() bool`

HasFkiModuleID returns a boolean if a field has been set.

### GetFkiModulesectionID

`func (o *VersionhistoryResponse) GetFkiModulesectionID() int32`

GetFkiModulesectionID returns the FkiModulesectionID field if non-nil, zero value otherwise.

### GetFkiModulesectionIDOk

`func (o *VersionhistoryResponse) GetFkiModulesectionIDOk() (*int32, bool)`

GetFkiModulesectionIDOk returns a tuple with the FkiModulesectionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiModulesectionID

`func (o *VersionhistoryResponse) SetFkiModulesectionID(v int32)`

SetFkiModulesectionID sets FkiModulesectionID field to given value.

### HasFkiModulesectionID

`func (o *VersionhistoryResponse) HasFkiModulesectionID() bool`

HasFkiModulesectionID returns a boolean if a field has been set.

### GetSModuleNameX

`func (o *VersionhistoryResponse) GetSModuleNameX() string`

GetSModuleNameX returns the SModuleNameX field if non-nil, zero value otherwise.

### GetSModuleNameXOk

`func (o *VersionhistoryResponse) GetSModuleNameXOk() (*string, bool)`

GetSModuleNameXOk returns a tuple with the SModuleNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSModuleNameX

`func (o *VersionhistoryResponse) SetSModuleNameX(v string)`

SetSModuleNameX sets SModuleNameX field to given value.

### HasSModuleNameX

`func (o *VersionhistoryResponse) HasSModuleNameX() bool`

HasSModuleNameX returns a boolean if a field has been set.

### GetSModulesectionNameX

`func (o *VersionhistoryResponse) GetSModulesectionNameX() string`

GetSModulesectionNameX returns the SModulesectionNameX field if non-nil, zero value otherwise.

### GetSModulesectionNameXOk

`func (o *VersionhistoryResponse) GetSModulesectionNameXOk() (*string, bool)`

GetSModulesectionNameXOk returns a tuple with the SModulesectionNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSModulesectionNameX

`func (o *VersionhistoryResponse) SetSModulesectionNameX(v string)`

SetSModulesectionNameX sets SModulesectionNameX field to given value.

### HasSModulesectionNameX

`func (o *VersionhistoryResponse) HasSModulesectionNameX() bool`

HasSModulesectionNameX returns a boolean if a field has been set.

### GetEVersionhistoryUsertype

`func (o *VersionhistoryResponse) GetEVersionhistoryUsertype() FieldEVersionhistoryUsertype`

GetEVersionhistoryUsertype returns the EVersionhistoryUsertype field if non-nil, zero value otherwise.

### GetEVersionhistoryUsertypeOk

`func (o *VersionhistoryResponse) GetEVersionhistoryUsertypeOk() (*FieldEVersionhistoryUsertype, bool)`

GetEVersionhistoryUsertypeOk returns a tuple with the EVersionhistoryUsertype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEVersionhistoryUsertype

`func (o *VersionhistoryResponse) SetEVersionhistoryUsertype(v FieldEVersionhistoryUsertype)`

SetEVersionhistoryUsertype sets EVersionhistoryUsertype field to given value.

### HasEVersionhistoryUsertype

`func (o *VersionhistoryResponse) HasEVersionhistoryUsertype() bool`

HasEVersionhistoryUsertype returns a boolean if a field has been set.

### GetObjVersionhistoryDetail

`func (o *VersionhistoryResponse) GetObjVersionhistoryDetail() MultilingualVersionhistoryDetail`

GetObjVersionhistoryDetail returns the ObjVersionhistoryDetail field if non-nil, zero value otherwise.

### GetObjVersionhistoryDetailOk

`func (o *VersionhistoryResponse) GetObjVersionhistoryDetailOk() (*MultilingualVersionhistoryDetail, bool)`

GetObjVersionhistoryDetailOk returns a tuple with the ObjVersionhistoryDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjVersionhistoryDetail

`func (o *VersionhistoryResponse) SetObjVersionhistoryDetail(v MultilingualVersionhistoryDetail)`

SetObjVersionhistoryDetail sets ObjVersionhistoryDetail field to given value.


### GetDtVersionhistoryDate

`func (o *VersionhistoryResponse) GetDtVersionhistoryDate() string`

GetDtVersionhistoryDate returns the DtVersionhistoryDate field if non-nil, zero value otherwise.

### GetDtVersionhistoryDateOk

`func (o *VersionhistoryResponse) GetDtVersionhistoryDateOk() (*string, bool)`

GetDtVersionhistoryDateOk returns a tuple with the DtVersionhistoryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtVersionhistoryDate

`func (o *VersionhistoryResponse) SetDtVersionhistoryDate(v string)`

SetDtVersionhistoryDate sets DtVersionhistoryDate field to given value.


### GetDtVersionhistoryDateend

`func (o *VersionhistoryResponse) GetDtVersionhistoryDateend() string`

GetDtVersionhistoryDateend returns the DtVersionhistoryDateend field if non-nil, zero value otherwise.

### GetDtVersionhistoryDateendOk

`func (o *VersionhistoryResponse) GetDtVersionhistoryDateendOk() (*string, bool)`

GetDtVersionhistoryDateendOk returns a tuple with the DtVersionhistoryDateend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtVersionhistoryDateend

`func (o *VersionhistoryResponse) SetDtVersionhistoryDateend(v string)`

SetDtVersionhistoryDateend sets DtVersionhistoryDateend field to given value.

### HasDtVersionhistoryDateend

`func (o *VersionhistoryResponse) HasDtVersionhistoryDateend() bool`

HasDtVersionhistoryDateend returns a boolean if a field has been set.

### GetEVersionhistoryType

`func (o *VersionhistoryResponse) GetEVersionhistoryType() FieldEVersionhistoryType`

GetEVersionhistoryType returns the EVersionhistoryType field if non-nil, zero value otherwise.

### GetEVersionhistoryTypeOk

`func (o *VersionhistoryResponse) GetEVersionhistoryTypeOk() (*FieldEVersionhistoryType, bool)`

GetEVersionhistoryTypeOk returns a tuple with the EVersionhistoryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEVersionhistoryType

`func (o *VersionhistoryResponse) SetEVersionhistoryType(v FieldEVersionhistoryType)`

SetEVersionhistoryType sets EVersionhistoryType field to given value.


### GetBVersionhistoryDraft

`func (o *VersionhistoryResponse) GetBVersionhistoryDraft() bool`

GetBVersionhistoryDraft returns the BVersionhistoryDraft field if non-nil, zero value otherwise.

### GetBVersionhistoryDraftOk

`func (o *VersionhistoryResponse) GetBVersionhistoryDraftOk() (*bool, bool)`

GetBVersionhistoryDraftOk returns a tuple with the BVersionhistoryDraft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBVersionhistoryDraft

`func (o *VersionhistoryResponse) SetBVersionhistoryDraft(v bool)`

SetBVersionhistoryDraft sets BVersionhistoryDraft field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


