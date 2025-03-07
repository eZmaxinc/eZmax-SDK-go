# EzsignimportfolderListElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsignimportfolderID** | **int32** | The unique ID of the Ezsignimportfolder | 
**SEzsignimportfolderName** | **string** | The name of the Ezsignimportfolder | 
**DtCreatedDate** | Pointer to **string** | The date and time at which the object was created | [optional] 
**DtModifiedDate** | Pointer to **string** | The date and time at which the object was last modified | [optional] 
**ITotalEzsignimportdocument** | Pointer to **int32** | The count of Ezsignimportdocument. | [optional] 
**ITotalEzsignimportdocumentNotImported** | Pointer to **int32** | The count of Ezsignimportdocument not imported in an Ezsignfolder. | [optional] 
**EEzsignimportfolderStatus** | Pointer to [**ComputedEEzsignimportfolderStatus**](ComputedEEzsignimportfolderStatus.md) |  | [optional] 

## Methods

### NewEzsignimportfolderListElement

`func NewEzsignimportfolderListElement(pkiEzsignimportfolderID int32, sEzsignimportfolderName string, ) *EzsignimportfolderListElement`

NewEzsignimportfolderListElement instantiates a new EzsignimportfolderListElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsignimportfolderListElementWithDefaults

`func NewEzsignimportfolderListElementWithDefaults() *EzsignimportfolderListElement`

NewEzsignimportfolderListElementWithDefaults instantiates a new EzsignimportfolderListElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsignimportfolderID

`func (o *EzsignimportfolderListElement) GetPkiEzsignimportfolderID() int32`

GetPkiEzsignimportfolderID returns the PkiEzsignimportfolderID field if non-nil, zero value otherwise.

### GetPkiEzsignimportfolderIDOk

`func (o *EzsignimportfolderListElement) GetPkiEzsignimportfolderIDOk() (*int32, bool)`

GetPkiEzsignimportfolderIDOk returns a tuple with the PkiEzsignimportfolderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsignimportfolderID

`func (o *EzsignimportfolderListElement) SetPkiEzsignimportfolderID(v int32)`

SetPkiEzsignimportfolderID sets PkiEzsignimportfolderID field to given value.


### GetSEzsignimportfolderName

`func (o *EzsignimportfolderListElement) GetSEzsignimportfolderName() string`

GetSEzsignimportfolderName returns the SEzsignimportfolderName field if non-nil, zero value otherwise.

### GetSEzsignimportfolderNameOk

`func (o *EzsignimportfolderListElement) GetSEzsignimportfolderNameOk() (*string, bool)`

GetSEzsignimportfolderNameOk returns a tuple with the SEzsignimportfolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignimportfolderName

`func (o *EzsignimportfolderListElement) SetSEzsignimportfolderName(v string)`

SetSEzsignimportfolderName sets SEzsignimportfolderName field to given value.


### GetDtCreatedDate

`func (o *EzsignimportfolderListElement) GetDtCreatedDate() string`

GetDtCreatedDate returns the DtCreatedDate field if non-nil, zero value otherwise.

### GetDtCreatedDateOk

`func (o *EzsignimportfolderListElement) GetDtCreatedDateOk() (*string, bool)`

GetDtCreatedDateOk returns a tuple with the DtCreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtCreatedDate

`func (o *EzsignimportfolderListElement) SetDtCreatedDate(v string)`

SetDtCreatedDate sets DtCreatedDate field to given value.

### HasDtCreatedDate

`func (o *EzsignimportfolderListElement) HasDtCreatedDate() bool`

HasDtCreatedDate returns a boolean if a field has been set.

### GetDtModifiedDate

`func (o *EzsignimportfolderListElement) GetDtModifiedDate() string`

GetDtModifiedDate returns the DtModifiedDate field if non-nil, zero value otherwise.

### GetDtModifiedDateOk

`func (o *EzsignimportfolderListElement) GetDtModifiedDateOk() (*string, bool)`

GetDtModifiedDateOk returns a tuple with the DtModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtModifiedDate

`func (o *EzsignimportfolderListElement) SetDtModifiedDate(v string)`

SetDtModifiedDate sets DtModifiedDate field to given value.

### HasDtModifiedDate

`func (o *EzsignimportfolderListElement) HasDtModifiedDate() bool`

HasDtModifiedDate returns a boolean if a field has been set.

### GetITotalEzsignimportdocument

`func (o *EzsignimportfolderListElement) GetITotalEzsignimportdocument() int32`

GetITotalEzsignimportdocument returns the ITotalEzsignimportdocument field if non-nil, zero value otherwise.

### GetITotalEzsignimportdocumentOk

`func (o *EzsignimportfolderListElement) GetITotalEzsignimportdocumentOk() (*int32, bool)`

GetITotalEzsignimportdocumentOk returns a tuple with the ITotalEzsignimportdocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetITotalEzsignimportdocument

`func (o *EzsignimportfolderListElement) SetITotalEzsignimportdocument(v int32)`

SetITotalEzsignimportdocument sets ITotalEzsignimportdocument field to given value.

### HasITotalEzsignimportdocument

`func (o *EzsignimportfolderListElement) HasITotalEzsignimportdocument() bool`

HasITotalEzsignimportdocument returns a boolean if a field has been set.

### GetITotalEzsignimportdocumentNotImported

`func (o *EzsignimportfolderListElement) GetITotalEzsignimportdocumentNotImported() int32`

GetITotalEzsignimportdocumentNotImported returns the ITotalEzsignimportdocumentNotImported field if non-nil, zero value otherwise.

### GetITotalEzsignimportdocumentNotImportedOk

`func (o *EzsignimportfolderListElement) GetITotalEzsignimportdocumentNotImportedOk() (*int32, bool)`

GetITotalEzsignimportdocumentNotImportedOk returns a tuple with the ITotalEzsignimportdocumentNotImported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetITotalEzsignimportdocumentNotImported

`func (o *EzsignimportfolderListElement) SetITotalEzsignimportdocumentNotImported(v int32)`

SetITotalEzsignimportdocumentNotImported sets ITotalEzsignimportdocumentNotImported field to given value.

### HasITotalEzsignimportdocumentNotImported

`func (o *EzsignimportfolderListElement) HasITotalEzsignimportdocumentNotImported() bool`

HasITotalEzsignimportdocumentNotImported returns a boolean if a field has been set.

### GetEEzsignimportfolderStatus

`func (o *EzsignimportfolderListElement) GetEEzsignimportfolderStatus() ComputedEEzsignimportfolderStatus`

GetEEzsignimportfolderStatus returns the EEzsignimportfolderStatus field if non-nil, zero value otherwise.

### GetEEzsignimportfolderStatusOk

`func (o *EzsignimportfolderListElement) GetEEzsignimportfolderStatusOk() (*ComputedEEzsignimportfolderStatus, bool)`

GetEEzsignimportfolderStatusOk returns a tuple with the EEzsignimportfolderStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignimportfolderStatus

`func (o *EzsignimportfolderListElement) SetEEzsignimportfolderStatus(v ComputedEEzsignimportfolderStatus)`

SetEEzsignimportfolderStatus sets EEzsignimportfolderStatus field to given value.

### HasEEzsignimportfolderStatus

`func (o *EzsignimportfolderListElement) HasEEzsignimportfolderStatus() bool`

HasEEzsignimportfolderStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


