# EzsignbulksendResponseV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsignbulksendID** | **int32** | The unique ID of the Ezsignbulksend | 
**FkiEzsignfoldertypeID** | **int32** | The unique ID of the Ezsignfoldertype. | 
**FkiLanguageID** | **int32** | The unique ID of the Language.  Valid values:  |Value|Description| |-|-| |1|French| |2|English| | 
**SLanguageNameX** | **string** | The Name of the Language in the language of the requester | 
**EEzsignbulksendEzsignformfieldorder** | [**FieldEEzsignbulksendEzsignformfieldorder**](FieldEEzsignbulksendEzsignformfieldorder.md) |  | 
**EEzsignfoldertypePrivacylevel** | [**FieldEEzsignfoldertypePrivacylevel**](FieldEEzsignfoldertypePrivacylevel.md) |  | 
**SEzsignfoldertypeNameX** | **string** | The name of the Ezsignfoldertype in the language of the requester | 
**SEzsignbulksendDescription** | **string** | The description of the Ezsignbulksend | 
**TEzsignbulksendNote** | **string** | Note about the Ezsignbulksend | 
**BEzsignbulksendNeedvalidation** | **bool** | Whether the Ezsigntemplatepackage was automatically modified and needs a manual validation | 
**BEzsignbulksendIsactive** | **bool** | Whether the Ezsignbulksend is active or not | 
**ObjAudit** | [**CommonAudit**](CommonAudit.md) |  | 

## Methods

### NewEzsignbulksendResponseV3

`func NewEzsignbulksendResponseV3(pkiEzsignbulksendID int32, fkiEzsignfoldertypeID int32, fkiLanguageID int32, sLanguageNameX string, eEzsignbulksendEzsignformfieldorder FieldEEzsignbulksendEzsignformfieldorder, eEzsignfoldertypePrivacylevel FieldEEzsignfoldertypePrivacylevel, sEzsignfoldertypeNameX string, sEzsignbulksendDescription string, tEzsignbulksendNote string, bEzsignbulksendNeedvalidation bool, bEzsignbulksendIsactive bool, objAudit CommonAudit, ) *EzsignbulksendResponseV3`

NewEzsignbulksendResponseV3 instantiates a new EzsignbulksendResponseV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsignbulksendResponseV3WithDefaults

`func NewEzsignbulksendResponseV3WithDefaults() *EzsignbulksendResponseV3`

NewEzsignbulksendResponseV3WithDefaults instantiates a new EzsignbulksendResponseV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsignbulksendID

`func (o *EzsignbulksendResponseV3) GetPkiEzsignbulksendID() int32`

GetPkiEzsignbulksendID returns the PkiEzsignbulksendID field if non-nil, zero value otherwise.

### GetPkiEzsignbulksendIDOk

`func (o *EzsignbulksendResponseV3) GetPkiEzsignbulksendIDOk() (*int32, bool)`

GetPkiEzsignbulksendIDOk returns a tuple with the PkiEzsignbulksendID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsignbulksendID

`func (o *EzsignbulksendResponseV3) SetPkiEzsignbulksendID(v int32)`

SetPkiEzsignbulksendID sets PkiEzsignbulksendID field to given value.


### GetFkiEzsignfoldertypeID

`func (o *EzsignbulksendResponseV3) GetFkiEzsignfoldertypeID() int32`

GetFkiEzsignfoldertypeID returns the FkiEzsignfoldertypeID field if non-nil, zero value otherwise.

### GetFkiEzsignfoldertypeIDOk

`func (o *EzsignbulksendResponseV3) GetFkiEzsignfoldertypeIDOk() (*int32, bool)`

GetFkiEzsignfoldertypeIDOk returns a tuple with the FkiEzsignfoldertypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignfoldertypeID

`func (o *EzsignbulksendResponseV3) SetFkiEzsignfoldertypeID(v int32)`

SetFkiEzsignfoldertypeID sets FkiEzsignfoldertypeID field to given value.


### GetFkiLanguageID

`func (o *EzsignbulksendResponseV3) GetFkiLanguageID() int32`

GetFkiLanguageID returns the FkiLanguageID field if non-nil, zero value otherwise.

### GetFkiLanguageIDOk

`func (o *EzsignbulksendResponseV3) GetFkiLanguageIDOk() (*int32, bool)`

GetFkiLanguageIDOk returns a tuple with the FkiLanguageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiLanguageID

`func (o *EzsignbulksendResponseV3) SetFkiLanguageID(v int32)`

SetFkiLanguageID sets FkiLanguageID field to given value.


### GetSLanguageNameX

`func (o *EzsignbulksendResponseV3) GetSLanguageNameX() string`

GetSLanguageNameX returns the SLanguageNameX field if non-nil, zero value otherwise.

### GetSLanguageNameXOk

`func (o *EzsignbulksendResponseV3) GetSLanguageNameXOk() (*string, bool)`

GetSLanguageNameXOk returns a tuple with the SLanguageNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSLanguageNameX

`func (o *EzsignbulksendResponseV3) SetSLanguageNameX(v string)`

SetSLanguageNameX sets SLanguageNameX field to given value.


### GetEEzsignbulksendEzsignformfieldorder

`func (o *EzsignbulksendResponseV3) GetEEzsignbulksendEzsignformfieldorder() FieldEEzsignbulksendEzsignformfieldorder`

GetEEzsignbulksendEzsignformfieldorder returns the EEzsignbulksendEzsignformfieldorder field if non-nil, zero value otherwise.

### GetEEzsignbulksendEzsignformfieldorderOk

`func (o *EzsignbulksendResponseV3) GetEEzsignbulksendEzsignformfieldorderOk() (*FieldEEzsignbulksendEzsignformfieldorder, bool)`

GetEEzsignbulksendEzsignformfieldorderOk returns a tuple with the EEzsignbulksendEzsignformfieldorder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignbulksendEzsignformfieldorder

`func (o *EzsignbulksendResponseV3) SetEEzsignbulksendEzsignformfieldorder(v FieldEEzsignbulksendEzsignformfieldorder)`

SetEEzsignbulksendEzsignformfieldorder sets EEzsignbulksendEzsignformfieldorder field to given value.


### GetEEzsignfoldertypePrivacylevel

`func (o *EzsignbulksendResponseV3) GetEEzsignfoldertypePrivacylevel() FieldEEzsignfoldertypePrivacylevel`

GetEEzsignfoldertypePrivacylevel returns the EEzsignfoldertypePrivacylevel field if non-nil, zero value otherwise.

### GetEEzsignfoldertypePrivacylevelOk

`func (o *EzsignbulksendResponseV3) GetEEzsignfoldertypePrivacylevelOk() (*FieldEEzsignfoldertypePrivacylevel, bool)`

GetEEzsignfoldertypePrivacylevelOk returns a tuple with the EEzsignfoldertypePrivacylevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignfoldertypePrivacylevel

`func (o *EzsignbulksendResponseV3) SetEEzsignfoldertypePrivacylevel(v FieldEEzsignfoldertypePrivacylevel)`

SetEEzsignfoldertypePrivacylevel sets EEzsignfoldertypePrivacylevel field to given value.


### GetSEzsignfoldertypeNameX

`func (o *EzsignbulksendResponseV3) GetSEzsignfoldertypeNameX() string`

GetSEzsignfoldertypeNameX returns the SEzsignfoldertypeNameX field if non-nil, zero value otherwise.

### GetSEzsignfoldertypeNameXOk

`func (o *EzsignbulksendResponseV3) GetSEzsignfoldertypeNameXOk() (*string, bool)`

GetSEzsignfoldertypeNameXOk returns a tuple with the SEzsignfoldertypeNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignfoldertypeNameX

`func (o *EzsignbulksendResponseV3) SetSEzsignfoldertypeNameX(v string)`

SetSEzsignfoldertypeNameX sets SEzsignfoldertypeNameX field to given value.


### GetSEzsignbulksendDescription

`func (o *EzsignbulksendResponseV3) GetSEzsignbulksendDescription() string`

GetSEzsignbulksendDescription returns the SEzsignbulksendDescription field if non-nil, zero value otherwise.

### GetSEzsignbulksendDescriptionOk

`func (o *EzsignbulksendResponseV3) GetSEzsignbulksendDescriptionOk() (*string, bool)`

GetSEzsignbulksendDescriptionOk returns a tuple with the SEzsignbulksendDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignbulksendDescription

`func (o *EzsignbulksendResponseV3) SetSEzsignbulksendDescription(v string)`

SetSEzsignbulksendDescription sets SEzsignbulksendDescription field to given value.


### GetTEzsignbulksendNote

`func (o *EzsignbulksendResponseV3) GetTEzsignbulksendNote() string`

GetTEzsignbulksendNote returns the TEzsignbulksendNote field if non-nil, zero value otherwise.

### GetTEzsignbulksendNoteOk

`func (o *EzsignbulksendResponseV3) GetTEzsignbulksendNoteOk() (*string, bool)`

GetTEzsignbulksendNoteOk returns a tuple with the TEzsignbulksendNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTEzsignbulksendNote

`func (o *EzsignbulksendResponseV3) SetTEzsignbulksendNote(v string)`

SetTEzsignbulksendNote sets TEzsignbulksendNote field to given value.


### GetBEzsignbulksendNeedvalidation

`func (o *EzsignbulksendResponseV3) GetBEzsignbulksendNeedvalidation() bool`

GetBEzsignbulksendNeedvalidation returns the BEzsignbulksendNeedvalidation field if non-nil, zero value otherwise.

### GetBEzsignbulksendNeedvalidationOk

`func (o *EzsignbulksendResponseV3) GetBEzsignbulksendNeedvalidationOk() (*bool, bool)`

GetBEzsignbulksendNeedvalidationOk returns a tuple with the BEzsignbulksendNeedvalidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignbulksendNeedvalidation

`func (o *EzsignbulksendResponseV3) SetBEzsignbulksendNeedvalidation(v bool)`

SetBEzsignbulksendNeedvalidation sets BEzsignbulksendNeedvalidation field to given value.


### GetBEzsignbulksendIsactive

`func (o *EzsignbulksendResponseV3) GetBEzsignbulksendIsactive() bool`

GetBEzsignbulksendIsactive returns the BEzsignbulksendIsactive field if non-nil, zero value otherwise.

### GetBEzsignbulksendIsactiveOk

`func (o *EzsignbulksendResponseV3) GetBEzsignbulksendIsactiveOk() (*bool, bool)`

GetBEzsignbulksendIsactiveOk returns a tuple with the BEzsignbulksendIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignbulksendIsactive

`func (o *EzsignbulksendResponseV3) SetBEzsignbulksendIsactive(v bool)`

SetBEzsignbulksendIsactive sets BEzsignbulksendIsactive field to given value.


### GetObjAudit

`func (o *EzsignbulksendResponseV3) GetObjAudit() CommonAudit`

GetObjAudit returns the ObjAudit field if non-nil, zero value otherwise.

### GetObjAuditOk

`func (o *EzsignbulksendResponseV3) GetObjAuditOk() (*CommonAudit, bool)`

GetObjAuditOk returns a tuple with the ObjAudit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjAudit

`func (o *EzsignbulksendResponseV3) SetObjAudit(v CommonAudit)`

SetObjAudit sets ObjAudit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


