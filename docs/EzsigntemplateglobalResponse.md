# EzsigntemplateglobalResponse

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

## Methods

### NewEzsigntemplateglobalResponse

`func NewEzsigntemplateglobalResponse(pkiEzsigntemplateglobalID int32, fkiEzsigntemplateglobaldocumentID int32, fkiModuleID int32, fkiLanguageID int32, sLanguageNameX string, eEzsigntemplateglobalModule FieldEEzsigntemplateglobalModule, eEzsigntemplateglobalSupplier FieldEEzsigntemplateglobalSupplier, sEzsigntemplateglobalCode string, sEzsigntemplateglobalDescription string, ) *EzsigntemplateglobalResponse`

NewEzsigntemplateglobalResponse instantiates a new EzsigntemplateglobalResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigntemplateglobalResponseWithDefaults

`func NewEzsigntemplateglobalResponseWithDefaults() *EzsigntemplateglobalResponse`

NewEzsigntemplateglobalResponseWithDefaults instantiates a new EzsigntemplateglobalResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsigntemplateglobalID

`func (o *EzsigntemplateglobalResponse) GetPkiEzsigntemplateglobalID() int32`

GetPkiEzsigntemplateglobalID returns the PkiEzsigntemplateglobalID field if non-nil, zero value otherwise.

### GetPkiEzsigntemplateglobalIDOk

`func (o *EzsigntemplateglobalResponse) GetPkiEzsigntemplateglobalIDOk() (*int32, bool)`

GetPkiEzsigntemplateglobalIDOk returns a tuple with the PkiEzsigntemplateglobalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsigntemplateglobalID

`func (o *EzsigntemplateglobalResponse) SetPkiEzsigntemplateglobalID(v int32)`

SetPkiEzsigntemplateglobalID sets PkiEzsigntemplateglobalID field to given value.


### GetFkiEzsigntemplateglobaldocumentID

`func (o *EzsigntemplateglobalResponse) GetFkiEzsigntemplateglobaldocumentID() int32`

GetFkiEzsigntemplateglobaldocumentID returns the FkiEzsigntemplateglobaldocumentID field if non-nil, zero value otherwise.

### GetFkiEzsigntemplateglobaldocumentIDOk

`func (o *EzsigntemplateglobalResponse) GetFkiEzsigntemplateglobaldocumentIDOk() (*int32, bool)`

GetFkiEzsigntemplateglobaldocumentIDOk returns a tuple with the FkiEzsigntemplateglobaldocumentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigntemplateglobaldocumentID

`func (o *EzsigntemplateglobalResponse) SetFkiEzsigntemplateglobaldocumentID(v int32)`

SetFkiEzsigntemplateglobaldocumentID sets FkiEzsigntemplateglobaldocumentID field to given value.


### GetFkiModuleID

`func (o *EzsigntemplateglobalResponse) GetFkiModuleID() int32`

GetFkiModuleID returns the FkiModuleID field if non-nil, zero value otherwise.

### GetFkiModuleIDOk

`func (o *EzsigntemplateglobalResponse) GetFkiModuleIDOk() (*int32, bool)`

GetFkiModuleIDOk returns a tuple with the FkiModuleID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiModuleID

`func (o *EzsigntemplateglobalResponse) SetFkiModuleID(v int32)`

SetFkiModuleID sets FkiModuleID field to given value.


### GetSModuleNameX

`func (o *EzsigntemplateglobalResponse) GetSModuleNameX() string`

GetSModuleNameX returns the SModuleNameX field if non-nil, zero value otherwise.

### GetSModuleNameXOk

`func (o *EzsigntemplateglobalResponse) GetSModuleNameXOk() (*string, bool)`

GetSModuleNameXOk returns a tuple with the SModuleNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSModuleNameX

`func (o *EzsigntemplateglobalResponse) SetSModuleNameX(v string)`

SetSModuleNameX sets SModuleNameX field to given value.

### HasSModuleNameX

`func (o *EzsigntemplateglobalResponse) HasSModuleNameX() bool`

HasSModuleNameX returns a boolean if a field has been set.

### GetFkiLanguageID

`func (o *EzsigntemplateglobalResponse) GetFkiLanguageID() int32`

GetFkiLanguageID returns the FkiLanguageID field if non-nil, zero value otherwise.

### GetFkiLanguageIDOk

`func (o *EzsigntemplateglobalResponse) GetFkiLanguageIDOk() (*int32, bool)`

GetFkiLanguageIDOk returns a tuple with the FkiLanguageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiLanguageID

`func (o *EzsigntemplateglobalResponse) SetFkiLanguageID(v int32)`

SetFkiLanguageID sets FkiLanguageID field to given value.


### GetSLanguageNameX

`func (o *EzsigntemplateglobalResponse) GetSLanguageNameX() string`

GetSLanguageNameX returns the SLanguageNameX field if non-nil, zero value otherwise.

### GetSLanguageNameXOk

`func (o *EzsigntemplateglobalResponse) GetSLanguageNameXOk() (*string, bool)`

GetSLanguageNameXOk returns a tuple with the SLanguageNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSLanguageNameX

`func (o *EzsigntemplateglobalResponse) SetSLanguageNameX(v string)`

SetSLanguageNameX sets SLanguageNameX field to given value.


### GetEEzsigntemplateglobalModule

`func (o *EzsigntemplateglobalResponse) GetEEzsigntemplateglobalModule() FieldEEzsigntemplateglobalModule`

GetEEzsigntemplateglobalModule returns the EEzsigntemplateglobalModule field if non-nil, zero value otherwise.

### GetEEzsigntemplateglobalModuleOk

`func (o *EzsigntemplateglobalResponse) GetEEzsigntemplateglobalModuleOk() (*FieldEEzsigntemplateglobalModule, bool)`

GetEEzsigntemplateglobalModuleOk returns a tuple with the EEzsigntemplateglobalModule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplateglobalModule

`func (o *EzsigntemplateglobalResponse) SetEEzsigntemplateglobalModule(v FieldEEzsigntemplateglobalModule)`

SetEEzsigntemplateglobalModule sets EEzsigntemplateglobalModule field to given value.


### GetEEzsigntemplateglobalSupplier

`func (o *EzsigntemplateglobalResponse) GetEEzsigntemplateglobalSupplier() FieldEEzsigntemplateglobalSupplier`

GetEEzsigntemplateglobalSupplier returns the EEzsigntemplateglobalSupplier field if non-nil, zero value otherwise.

### GetEEzsigntemplateglobalSupplierOk

`func (o *EzsigntemplateglobalResponse) GetEEzsigntemplateglobalSupplierOk() (*FieldEEzsigntemplateglobalSupplier, bool)`

GetEEzsigntemplateglobalSupplierOk returns a tuple with the EEzsigntemplateglobalSupplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplateglobalSupplier

`func (o *EzsigntemplateglobalResponse) SetEEzsigntemplateglobalSupplier(v FieldEEzsigntemplateglobalSupplier)`

SetEEzsigntemplateglobalSupplier sets EEzsigntemplateglobalSupplier field to given value.


### GetSEzsigntemplateglobalCode

`func (o *EzsigntemplateglobalResponse) GetSEzsigntemplateglobalCode() string`

GetSEzsigntemplateglobalCode returns the SEzsigntemplateglobalCode field if non-nil, zero value otherwise.

### GetSEzsigntemplateglobalCodeOk

`func (o *EzsigntemplateglobalResponse) GetSEzsigntemplateglobalCodeOk() (*string, bool)`

GetSEzsigntemplateglobalCodeOk returns a tuple with the SEzsigntemplateglobalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplateglobalCode

`func (o *EzsigntemplateglobalResponse) SetSEzsigntemplateglobalCode(v string)`

SetSEzsigntemplateglobalCode sets SEzsigntemplateglobalCode field to given value.


### GetSEzsigntemplateglobalDescription

`func (o *EzsigntemplateglobalResponse) GetSEzsigntemplateglobalDescription() string`

GetSEzsigntemplateglobalDescription returns the SEzsigntemplateglobalDescription field if non-nil, zero value otherwise.

### GetSEzsigntemplateglobalDescriptionOk

`func (o *EzsigntemplateglobalResponse) GetSEzsigntemplateglobalDescriptionOk() (*string, bool)`

GetSEzsigntemplateglobalDescriptionOk returns a tuple with the SEzsigntemplateglobalDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplateglobalDescription

`func (o *EzsigntemplateglobalResponse) SetSEzsigntemplateglobalDescription(v string)`

SetSEzsigntemplateglobalDescription sets SEzsigntemplateglobalDescription field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


