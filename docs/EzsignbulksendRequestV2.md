# EzsignbulksendRequestV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsignbulksendID** | Pointer to **int32** | The unique ID of the Ezsignbulksend | [optional] 
**FkiEzsignfoldertypeID** | **int32** | The unique ID of the Ezsignfoldertype. | 
**FkiLanguageID** | **int32** | The unique ID of the Language.  Valid values:  |Value|Description| |-|-| |1|French| |2|English| | 
**EEzsignbulksendEzsignformfieldorder** | [**FieldEEzsignbulksendEzsignformfieldorder**](FieldEEzsignbulksendEzsignformfieldorder.md) |  | 
**SEzsignbulksendDescription** | **string** | The description of the Ezsignbulksend | 
**TEzsignbulksendNote** | **string** | Note about the Ezsignbulksend | 
**BEzsignbulksendNeedvalidation** | **bool** | Whether the Ezsigntemplatepackage was automatically modified and needs a manual validation | 
**BEzsignbulksendIsactive** | **bool** | Whether the Ezsignbulksend is active or not | 

## Methods

### NewEzsignbulksendRequestV2

`func NewEzsignbulksendRequestV2(fkiEzsignfoldertypeID int32, fkiLanguageID int32, eEzsignbulksendEzsignformfieldorder FieldEEzsignbulksendEzsignformfieldorder, sEzsignbulksendDescription string, tEzsignbulksendNote string, bEzsignbulksendNeedvalidation bool, bEzsignbulksendIsactive bool, ) *EzsignbulksendRequestV2`

NewEzsignbulksendRequestV2 instantiates a new EzsignbulksendRequestV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsignbulksendRequestV2WithDefaults

`func NewEzsignbulksendRequestV2WithDefaults() *EzsignbulksendRequestV2`

NewEzsignbulksendRequestV2WithDefaults instantiates a new EzsignbulksendRequestV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsignbulksendID

`func (o *EzsignbulksendRequestV2) GetPkiEzsignbulksendID() int32`

GetPkiEzsignbulksendID returns the PkiEzsignbulksendID field if non-nil, zero value otherwise.

### GetPkiEzsignbulksendIDOk

`func (o *EzsignbulksendRequestV2) GetPkiEzsignbulksendIDOk() (*int32, bool)`

GetPkiEzsignbulksendIDOk returns a tuple with the PkiEzsignbulksendID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsignbulksendID

`func (o *EzsignbulksendRequestV2) SetPkiEzsignbulksendID(v int32)`

SetPkiEzsignbulksendID sets PkiEzsignbulksendID field to given value.

### HasPkiEzsignbulksendID

`func (o *EzsignbulksendRequestV2) HasPkiEzsignbulksendID() bool`

HasPkiEzsignbulksendID returns a boolean if a field has been set.

### GetFkiEzsignfoldertypeID

`func (o *EzsignbulksendRequestV2) GetFkiEzsignfoldertypeID() int32`

GetFkiEzsignfoldertypeID returns the FkiEzsignfoldertypeID field if non-nil, zero value otherwise.

### GetFkiEzsignfoldertypeIDOk

`func (o *EzsignbulksendRequestV2) GetFkiEzsignfoldertypeIDOk() (*int32, bool)`

GetFkiEzsignfoldertypeIDOk returns a tuple with the FkiEzsignfoldertypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignfoldertypeID

`func (o *EzsignbulksendRequestV2) SetFkiEzsignfoldertypeID(v int32)`

SetFkiEzsignfoldertypeID sets FkiEzsignfoldertypeID field to given value.


### GetFkiLanguageID

`func (o *EzsignbulksendRequestV2) GetFkiLanguageID() int32`

GetFkiLanguageID returns the FkiLanguageID field if non-nil, zero value otherwise.

### GetFkiLanguageIDOk

`func (o *EzsignbulksendRequestV2) GetFkiLanguageIDOk() (*int32, bool)`

GetFkiLanguageIDOk returns a tuple with the FkiLanguageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiLanguageID

`func (o *EzsignbulksendRequestV2) SetFkiLanguageID(v int32)`

SetFkiLanguageID sets FkiLanguageID field to given value.


### GetEEzsignbulksendEzsignformfieldorder

`func (o *EzsignbulksendRequestV2) GetEEzsignbulksendEzsignformfieldorder() FieldEEzsignbulksendEzsignformfieldorder`

GetEEzsignbulksendEzsignformfieldorder returns the EEzsignbulksendEzsignformfieldorder field if non-nil, zero value otherwise.

### GetEEzsignbulksendEzsignformfieldorderOk

`func (o *EzsignbulksendRequestV2) GetEEzsignbulksendEzsignformfieldorderOk() (*FieldEEzsignbulksendEzsignformfieldorder, bool)`

GetEEzsignbulksendEzsignformfieldorderOk returns a tuple with the EEzsignbulksendEzsignformfieldorder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignbulksendEzsignformfieldorder

`func (o *EzsignbulksendRequestV2) SetEEzsignbulksendEzsignformfieldorder(v FieldEEzsignbulksendEzsignformfieldorder)`

SetEEzsignbulksendEzsignformfieldorder sets EEzsignbulksendEzsignformfieldorder field to given value.


### GetSEzsignbulksendDescription

`func (o *EzsignbulksendRequestV2) GetSEzsignbulksendDescription() string`

GetSEzsignbulksendDescription returns the SEzsignbulksendDescription field if non-nil, zero value otherwise.

### GetSEzsignbulksendDescriptionOk

`func (o *EzsignbulksendRequestV2) GetSEzsignbulksendDescriptionOk() (*string, bool)`

GetSEzsignbulksendDescriptionOk returns a tuple with the SEzsignbulksendDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignbulksendDescription

`func (o *EzsignbulksendRequestV2) SetSEzsignbulksendDescription(v string)`

SetSEzsignbulksendDescription sets SEzsignbulksendDescription field to given value.


### GetTEzsignbulksendNote

`func (o *EzsignbulksendRequestV2) GetTEzsignbulksendNote() string`

GetTEzsignbulksendNote returns the TEzsignbulksendNote field if non-nil, zero value otherwise.

### GetTEzsignbulksendNoteOk

`func (o *EzsignbulksendRequestV2) GetTEzsignbulksendNoteOk() (*string, bool)`

GetTEzsignbulksendNoteOk returns a tuple with the TEzsignbulksendNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTEzsignbulksendNote

`func (o *EzsignbulksendRequestV2) SetTEzsignbulksendNote(v string)`

SetTEzsignbulksendNote sets TEzsignbulksendNote field to given value.


### GetBEzsignbulksendNeedvalidation

`func (o *EzsignbulksendRequestV2) GetBEzsignbulksendNeedvalidation() bool`

GetBEzsignbulksendNeedvalidation returns the BEzsignbulksendNeedvalidation field if non-nil, zero value otherwise.

### GetBEzsignbulksendNeedvalidationOk

`func (o *EzsignbulksendRequestV2) GetBEzsignbulksendNeedvalidationOk() (*bool, bool)`

GetBEzsignbulksendNeedvalidationOk returns a tuple with the BEzsignbulksendNeedvalidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignbulksendNeedvalidation

`func (o *EzsignbulksendRequestV2) SetBEzsignbulksendNeedvalidation(v bool)`

SetBEzsignbulksendNeedvalidation sets BEzsignbulksendNeedvalidation field to given value.


### GetBEzsignbulksendIsactive

`func (o *EzsignbulksendRequestV2) GetBEzsignbulksendIsactive() bool`

GetBEzsignbulksendIsactive returns the BEzsignbulksendIsactive field if non-nil, zero value otherwise.

### GetBEzsignbulksendIsactiveOk

`func (o *EzsignbulksendRequestV2) GetBEzsignbulksendIsactiveOk() (*bool, bool)`

GetBEzsignbulksendIsactiveOk returns a tuple with the BEzsignbulksendIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignbulksendIsactive

`func (o *EzsignbulksendRequestV2) SetBEzsignbulksendIsactive(v bool)`

SetBEzsignbulksendIsactive sets BEzsignbulksendIsactive field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


