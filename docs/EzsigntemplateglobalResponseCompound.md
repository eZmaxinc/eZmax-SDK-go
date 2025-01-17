# EzsigntemplateglobalResponseCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsigntemplateglobalID** | **int32** | The unique ID of the Ezsigntemplateglobal | 
**FkiEzsigntemplateglobaldocumentID** | **int32** | The unique ID of the Ezsigntemplateglobaldocument | 
**FkiModuleID** | **int32** | The unique ID of the Module | 
**SModuleNameX** | Pointer to **string** | The Name of the Module in the language of the requester | [optional] 
**FkiLanguageID** | **int32** | The unique ID of the Language.  Valid values:  |Value|Description| |-|-| |1|French| |2|English| | 
**SLanguageNameX** | **string** | The Name of the Language in the language of the requester | 
**EEzsigntemplateglobalModule** | [**FieldEEzsigntemplateglobalModule**](FieldEEzsigntemplateglobalModule.md) |  | 
**EEzsigntemplateglobalSupplier** | [**FieldEEzsigntemplateglobalSupplier**](FieldEEzsigntemplateglobalSupplier.md) |  | 
**SEzsigntemplateglobalCode** | **string** | The Code of the Ezsigntemplateglobal | 
**SEzsigntemplateglobalDescription** | **string** | The description of the Ezsigntemplate | 
**ObjEzsigntemplateglobaldocument** | Pointer to [**EzsigntemplateglobaldocumentResponse**](EzsigntemplateglobaldocumentResponse.md) |  | [optional] 
**AObjEzsigntemplateglobalsigner** | [**[]EzsigntemplateglobalsignerResponseCompound**](EzsigntemplateglobalsignerResponseCompound.md) |  | 

## Methods

### NewEzsigntemplateglobalResponseCompound

`func NewEzsigntemplateglobalResponseCompound(pkiEzsigntemplateglobalID int32, fkiEzsigntemplateglobaldocumentID int32, fkiModuleID int32, fkiLanguageID int32, sLanguageNameX string, eEzsigntemplateglobalModule FieldEEzsigntemplateglobalModule, eEzsigntemplateglobalSupplier FieldEEzsigntemplateglobalSupplier, sEzsigntemplateglobalCode string, sEzsigntemplateglobalDescription string, aObjEzsigntemplateglobalsigner []EzsigntemplateglobalsignerResponseCompound, ) *EzsigntemplateglobalResponseCompound`

NewEzsigntemplateglobalResponseCompound instantiates a new EzsigntemplateglobalResponseCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigntemplateglobalResponseCompoundWithDefaults

`func NewEzsigntemplateglobalResponseCompoundWithDefaults() *EzsigntemplateglobalResponseCompound`

NewEzsigntemplateglobalResponseCompoundWithDefaults instantiates a new EzsigntemplateglobalResponseCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsigntemplateglobalID

`func (o *EzsigntemplateglobalResponseCompound) GetPkiEzsigntemplateglobalID() int32`

GetPkiEzsigntemplateglobalID returns the PkiEzsigntemplateglobalID field if non-nil, zero value otherwise.

### GetPkiEzsigntemplateglobalIDOk

`func (o *EzsigntemplateglobalResponseCompound) GetPkiEzsigntemplateglobalIDOk() (*int32, bool)`

GetPkiEzsigntemplateglobalIDOk returns a tuple with the PkiEzsigntemplateglobalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsigntemplateglobalID

`func (o *EzsigntemplateglobalResponseCompound) SetPkiEzsigntemplateglobalID(v int32)`

SetPkiEzsigntemplateglobalID sets PkiEzsigntemplateglobalID field to given value.


### GetFkiEzsigntemplateglobaldocumentID

`func (o *EzsigntemplateglobalResponseCompound) GetFkiEzsigntemplateglobaldocumentID() int32`

GetFkiEzsigntemplateglobaldocumentID returns the FkiEzsigntemplateglobaldocumentID field if non-nil, zero value otherwise.

### GetFkiEzsigntemplateglobaldocumentIDOk

`func (o *EzsigntemplateglobalResponseCompound) GetFkiEzsigntemplateglobaldocumentIDOk() (*int32, bool)`

GetFkiEzsigntemplateglobaldocumentIDOk returns a tuple with the FkiEzsigntemplateglobaldocumentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigntemplateglobaldocumentID

`func (o *EzsigntemplateglobalResponseCompound) SetFkiEzsigntemplateglobaldocumentID(v int32)`

SetFkiEzsigntemplateglobaldocumentID sets FkiEzsigntemplateglobaldocumentID field to given value.


### GetFkiModuleID

`func (o *EzsigntemplateglobalResponseCompound) GetFkiModuleID() int32`

GetFkiModuleID returns the FkiModuleID field if non-nil, zero value otherwise.

### GetFkiModuleIDOk

`func (o *EzsigntemplateglobalResponseCompound) GetFkiModuleIDOk() (*int32, bool)`

GetFkiModuleIDOk returns a tuple with the FkiModuleID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiModuleID

`func (o *EzsigntemplateglobalResponseCompound) SetFkiModuleID(v int32)`

SetFkiModuleID sets FkiModuleID field to given value.


### GetSModuleNameX

`func (o *EzsigntemplateglobalResponseCompound) GetSModuleNameX() string`

GetSModuleNameX returns the SModuleNameX field if non-nil, zero value otherwise.

### GetSModuleNameXOk

`func (o *EzsigntemplateglobalResponseCompound) GetSModuleNameXOk() (*string, bool)`

GetSModuleNameXOk returns a tuple with the SModuleNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSModuleNameX

`func (o *EzsigntemplateglobalResponseCompound) SetSModuleNameX(v string)`

SetSModuleNameX sets SModuleNameX field to given value.

### HasSModuleNameX

`func (o *EzsigntemplateglobalResponseCompound) HasSModuleNameX() bool`

HasSModuleNameX returns a boolean if a field has been set.

### GetFkiLanguageID

`func (o *EzsigntemplateglobalResponseCompound) GetFkiLanguageID() int32`

GetFkiLanguageID returns the FkiLanguageID field if non-nil, zero value otherwise.

### GetFkiLanguageIDOk

`func (o *EzsigntemplateglobalResponseCompound) GetFkiLanguageIDOk() (*int32, bool)`

GetFkiLanguageIDOk returns a tuple with the FkiLanguageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiLanguageID

`func (o *EzsigntemplateglobalResponseCompound) SetFkiLanguageID(v int32)`

SetFkiLanguageID sets FkiLanguageID field to given value.


### GetSLanguageNameX

`func (o *EzsigntemplateglobalResponseCompound) GetSLanguageNameX() string`

GetSLanguageNameX returns the SLanguageNameX field if non-nil, zero value otherwise.

### GetSLanguageNameXOk

`func (o *EzsigntemplateglobalResponseCompound) GetSLanguageNameXOk() (*string, bool)`

GetSLanguageNameXOk returns a tuple with the SLanguageNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSLanguageNameX

`func (o *EzsigntemplateglobalResponseCompound) SetSLanguageNameX(v string)`

SetSLanguageNameX sets SLanguageNameX field to given value.


### GetEEzsigntemplateglobalModule

`func (o *EzsigntemplateglobalResponseCompound) GetEEzsigntemplateglobalModule() FieldEEzsigntemplateglobalModule`

GetEEzsigntemplateglobalModule returns the EEzsigntemplateglobalModule field if non-nil, zero value otherwise.

### GetEEzsigntemplateglobalModuleOk

`func (o *EzsigntemplateglobalResponseCompound) GetEEzsigntemplateglobalModuleOk() (*FieldEEzsigntemplateglobalModule, bool)`

GetEEzsigntemplateglobalModuleOk returns a tuple with the EEzsigntemplateglobalModule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplateglobalModule

`func (o *EzsigntemplateglobalResponseCompound) SetEEzsigntemplateglobalModule(v FieldEEzsigntemplateglobalModule)`

SetEEzsigntemplateglobalModule sets EEzsigntemplateglobalModule field to given value.


### GetEEzsigntemplateglobalSupplier

`func (o *EzsigntemplateglobalResponseCompound) GetEEzsigntemplateglobalSupplier() FieldEEzsigntemplateglobalSupplier`

GetEEzsigntemplateglobalSupplier returns the EEzsigntemplateglobalSupplier field if non-nil, zero value otherwise.

### GetEEzsigntemplateglobalSupplierOk

`func (o *EzsigntemplateglobalResponseCompound) GetEEzsigntemplateglobalSupplierOk() (*FieldEEzsigntemplateglobalSupplier, bool)`

GetEEzsigntemplateglobalSupplierOk returns a tuple with the EEzsigntemplateglobalSupplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplateglobalSupplier

`func (o *EzsigntemplateglobalResponseCompound) SetEEzsigntemplateglobalSupplier(v FieldEEzsigntemplateglobalSupplier)`

SetEEzsigntemplateglobalSupplier sets EEzsigntemplateglobalSupplier field to given value.


### GetSEzsigntemplateglobalCode

`func (o *EzsigntemplateglobalResponseCompound) GetSEzsigntemplateglobalCode() string`

GetSEzsigntemplateglobalCode returns the SEzsigntemplateglobalCode field if non-nil, zero value otherwise.

### GetSEzsigntemplateglobalCodeOk

`func (o *EzsigntemplateglobalResponseCompound) GetSEzsigntemplateglobalCodeOk() (*string, bool)`

GetSEzsigntemplateglobalCodeOk returns a tuple with the SEzsigntemplateglobalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplateglobalCode

`func (o *EzsigntemplateglobalResponseCompound) SetSEzsigntemplateglobalCode(v string)`

SetSEzsigntemplateglobalCode sets SEzsigntemplateglobalCode field to given value.


### GetSEzsigntemplateglobalDescription

`func (o *EzsigntemplateglobalResponseCompound) GetSEzsigntemplateglobalDescription() string`

GetSEzsigntemplateglobalDescription returns the SEzsigntemplateglobalDescription field if non-nil, zero value otherwise.

### GetSEzsigntemplateglobalDescriptionOk

`func (o *EzsigntemplateglobalResponseCompound) GetSEzsigntemplateglobalDescriptionOk() (*string, bool)`

GetSEzsigntemplateglobalDescriptionOk returns a tuple with the SEzsigntemplateglobalDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplateglobalDescription

`func (o *EzsigntemplateglobalResponseCompound) SetSEzsigntemplateglobalDescription(v string)`

SetSEzsigntemplateglobalDescription sets SEzsigntemplateglobalDescription field to given value.


### GetObjEzsigntemplateglobaldocument

`func (o *EzsigntemplateglobalResponseCompound) GetObjEzsigntemplateglobaldocument() EzsigntemplateglobaldocumentResponse`

GetObjEzsigntemplateglobaldocument returns the ObjEzsigntemplateglobaldocument field if non-nil, zero value otherwise.

### GetObjEzsigntemplateglobaldocumentOk

`func (o *EzsigntemplateglobalResponseCompound) GetObjEzsigntemplateglobaldocumentOk() (*EzsigntemplateglobaldocumentResponse, bool)`

GetObjEzsigntemplateglobaldocumentOk returns a tuple with the ObjEzsigntemplateglobaldocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjEzsigntemplateglobaldocument

`func (o *EzsigntemplateglobalResponseCompound) SetObjEzsigntemplateglobaldocument(v EzsigntemplateglobaldocumentResponse)`

SetObjEzsigntemplateglobaldocument sets ObjEzsigntemplateglobaldocument field to given value.

### HasObjEzsigntemplateglobaldocument

`func (o *EzsigntemplateglobalResponseCompound) HasObjEzsigntemplateglobaldocument() bool`

HasObjEzsigntemplateglobaldocument returns a boolean if a field has been set.

### GetAObjEzsigntemplateglobalsigner

`func (o *EzsigntemplateglobalResponseCompound) GetAObjEzsigntemplateglobalsigner() []EzsigntemplateglobalsignerResponseCompound`

GetAObjEzsigntemplateglobalsigner returns the AObjEzsigntemplateglobalsigner field if non-nil, zero value otherwise.

### GetAObjEzsigntemplateglobalsignerOk

`func (o *EzsigntemplateglobalResponseCompound) GetAObjEzsigntemplateglobalsignerOk() (*[]EzsigntemplateglobalsignerResponseCompound, bool)`

GetAObjEzsigntemplateglobalsignerOk returns a tuple with the AObjEzsigntemplateglobalsigner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEzsigntemplateglobalsigner

`func (o *EzsigntemplateglobalResponseCompound) SetAObjEzsigntemplateglobalsigner(v []EzsigntemplateglobalsignerResponseCompound)`

SetAObjEzsigntemplateglobalsigner sets AObjEzsigntemplateglobalsigner field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


