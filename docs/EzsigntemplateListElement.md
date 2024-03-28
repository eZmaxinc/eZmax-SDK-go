# EzsigntemplateListElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsigntemplateID** | **int32** | The unique ID of the Ezsigntemplate | 
**FkiEzsignfoldertypeID** | Pointer to **int32** | The unique ID of the Ezsignfoldertype. | [optional] 
**FkiLanguageID** | **int32** | The unique ID of the Language.  Valid values:  |Value|Description| |-|-| |1|French| |2|English| | 
**SEzsigntemplateDescription** | **string** | The description of the Ezsigntemplate | 
**IEzsigntemplatedocumentPagetotal** | Pointer to **int32** | The number of pages in the Ezsigntemplatedocument. | [optional] 
**IEzsigntemplateSignaturetotal** | Pointer to **int32** | The number of total signatures in the Ezsigntemplate. | [optional] 
**IEzsigntemplateFormfieldtotal** | Pointer to **int32** | The number of total form fields in the Ezsigntemplate. | [optional] 
**BEzsigntemplateIncomplete** | **bool** | Indicate the Ezsigntemplate is incomplete and cannot be used | 
**SEzsignfoldertypeNameX** | Pointer to **string** | The name of the Ezsignfoldertype in the language of the requester | [optional] 
**EEzsigntemplateType** | [**FieldEEzsigntemplateType**](FieldEEzsigntemplateType.md) |  | 

## Methods

### NewEzsigntemplateListElement

`func NewEzsigntemplateListElement(pkiEzsigntemplateID int32, fkiLanguageID int32, sEzsigntemplateDescription string, bEzsigntemplateIncomplete bool, eEzsigntemplateType FieldEEzsigntemplateType, ) *EzsigntemplateListElement`

NewEzsigntemplateListElement instantiates a new EzsigntemplateListElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigntemplateListElementWithDefaults

`func NewEzsigntemplateListElementWithDefaults() *EzsigntemplateListElement`

NewEzsigntemplateListElementWithDefaults instantiates a new EzsigntemplateListElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsigntemplateID

`func (o *EzsigntemplateListElement) GetPkiEzsigntemplateID() int32`

GetPkiEzsigntemplateID returns the PkiEzsigntemplateID field if non-nil, zero value otherwise.

### GetPkiEzsigntemplateIDOk

`func (o *EzsigntemplateListElement) GetPkiEzsigntemplateIDOk() (*int32, bool)`

GetPkiEzsigntemplateIDOk returns a tuple with the PkiEzsigntemplateID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsigntemplateID

`func (o *EzsigntemplateListElement) SetPkiEzsigntemplateID(v int32)`

SetPkiEzsigntemplateID sets PkiEzsigntemplateID field to given value.


### GetFkiEzsignfoldertypeID

`func (o *EzsigntemplateListElement) GetFkiEzsignfoldertypeID() int32`

GetFkiEzsignfoldertypeID returns the FkiEzsignfoldertypeID field if non-nil, zero value otherwise.

### GetFkiEzsignfoldertypeIDOk

`func (o *EzsigntemplateListElement) GetFkiEzsignfoldertypeIDOk() (*int32, bool)`

GetFkiEzsignfoldertypeIDOk returns a tuple with the FkiEzsignfoldertypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignfoldertypeID

`func (o *EzsigntemplateListElement) SetFkiEzsignfoldertypeID(v int32)`

SetFkiEzsignfoldertypeID sets FkiEzsignfoldertypeID field to given value.

### HasFkiEzsignfoldertypeID

`func (o *EzsigntemplateListElement) HasFkiEzsignfoldertypeID() bool`

HasFkiEzsignfoldertypeID returns a boolean if a field has been set.

### GetFkiLanguageID

`func (o *EzsigntemplateListElement) GetFkiLanguageID() int32`

GetFkiLanguageID returns the FkiLanguageID field if non-nil, zero value otherwise.

### GetFkiLanguageIDOk

`func (o *EzsigntemplateListElement) GetFkiLanguageIDOk() (*int32, bool)`

GetFkiLanguageIDOk returns a tuple with the FkiLanguageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiLanguageID

`func (o *EzsigntemplateListElement) SetFkiLanguageID(v int32)`

SetFkiLanguageID sets FkiLanguageID field to given value.


### GetSEzsigntemplateDescription

`func (o *EzsigntemplateListElement) GetSEzsigntemplateDescription() string`

GetSEzsigntemplateDescription returns the SEzsigntemplateDescription field if non-nil, zero value otherwise.

### GetSEzsigntemplateDescriptionOk

`func (o *EzsigntemplateListElement) GetSEzsigntemplateDescriptionOk() (*string, bool)`

GetSEzsigntemplateDescriptionOk returns a tuple with the SEzsigntemplateDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplateDescription

`func (o *EzsigntemplateListElement) SetSEzsigntemplateDescription(v string)`

SetSEzsigntemplateDescription sets SEzsigntemplateDescription field to given value.


### GetIEzsigntemplatedocumentPagetotal

`func (o *EzsigntemplateListElement) GetIEzsigntemplatedocumentPagetotal() int32`

GetIEzsigntemplatedocumentPagetotal returns the IEzsigntemplatedocumentPagetotal field if non-nil, zero value otherwise.

### GetIEzsigntemplatedocumentPagetotalOk

`func (o *EzsigntemplateListElement) GetIEzsigntemplatedocumentPagetotalOk() (*int32, bool)`

GetIEzsigntemplatedocumentPagetotalOk returns a tuple with the IEzsigntemplatedocumentPagetotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplatedocumentPagetotal

`func (o *EzsigntemplateListElement) SetIEzsigntemplatedocumentPagetotal(v int32)`

SetIEzsigntemplatedocumentPagetotal sets IEzsigntemplatedocumentPagetotal field to given value.

### HasIEzsigntemplatedocumentPagetotal

`func (o *EzsigntemplateListElement) HasIEzsigntemplatedocumentPagetotal() bool`

HasIEzsigntemplatedocumentPagetotal returns a boolean if a field has been set.

### GetIEzsigntemplateSignaturetotal

`func (o *EzsigntemplateListElement) GetIEzsigntemplateSignaturetotal() int32`

GetIEzsigntemplateSignaturetotal returns the IEzsigntemplateSignaturetotal field if non-nil, zero value otherwise.

### GetIEzsigntemplateSignaturetotalOk

`func (o *EzsigntemplateListElement) GetIEzsigntemplateSignaturetotalOk() (*int32, bool)`

GetIEzsigntemplateSignaturetotalOk returns a tuple with the IEzsigntemplateSignaturetotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplateSignaturetotal

`func (o *EzsigntemplateListElement) SetIEzsigntemplateSignaturetotal(v int32)`

SetIEzsigntemplateSignaturetotal sets IEzsigntemplateSignaturetotal field to given value.

### HasIEzsigntemplateSignaturetotal

`func (o *EzsigntemplateListElement) HasIEzsigntemplateSignaturetotal() bool`

HasIEzsigntemplateSignaturetotal returns a boolean if a field has been set.

### GetIEzsigntemplateFormfieldtotal

`func (o *EzsigntemplateListElement) GetIEzsigntemplateFormfieldtotal() int32`

GetIEzsigntemplateFormfieldtotal returns the IEzsigntemplateFormfieldtotal field if non-nil, zero value otherwise.

### GetIEzsigntemplateFormfieldtotalOk

`func (o *EzsigntemplateListElement) GetIEzsigntemplateFormfieldtotalOk() (*int32, bool)`

GetIEzsigntemplateFormfieldtotalOk returns a tuple with the IEzsigntemplateFormfieldtotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplateFormfieldtotal

`func (o *EzsigntemplateListElement) SetIEzsigntemplateFormfieldtotal(v int32)`

SetIEzsigntemplateFormfieldtotal sets IEzsigntemplateFormfieldtotal field to given value.

### HasIEzsigntemplateFormfieldtotal

`func (o *EzsigntemplateListElement) HasIEzsigntemplateFormfieldtotal() bool`

HasIEzsigntemplateFormfieldtotal returns a boolean if a field has been set.

### GetBEzsigntemplateIncomplete

`func (o *EzsigntemplateListElement) GetBEzsigntemplateIncomplete() bool`

GetBEzsigntemplateIncomplete returns the BEzsigntemplateIncomplete field if non-nil, zero value otherwise.

### GetBEzsigntemplateIncompleteOk

`func (o *EzsigntemplateListElement) GetBEzsigntemplateIncompleteOk() (*bool, bool)`

GetBEzsigntemplateIncompleteOk returns a tuple with the BEzsigntemplateIncomplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsigntemplateIncomplete

`func (o *EzsigntemplateListElement) SetBEzsigntemplateIncomplete(v bool)`

SetBEzsigntemplateIncomplete sets BEzsigntemplateIncomplete field to given value.


### GetSEzsignfoldertypeNameX

`func (o *EzsigntemplateListElement) GetSEzsignfoldertypeNameX() string`

GetSEzsignfoldertypeNameX returns the SEzsignfoldertypeNameX field if non-nil, zero value otherwise.

### GetSEzsignfoldertypeNameXOk

`func (o *EzsigntemplateListElement) GetSEzsignfoldertypeNameXOk() (*string, bool)`

GetSEzsignfoldertypeNameXOk returns a tuple with the SEzsignfoldertypeNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignfoldertypeNameX

`func (o *EzsigntemplateListElement) SetSEzsignfoldertypeNameX(v string)`

SetSEzsignfoldertypeNameX sets SEzsignfoldertypeNameX field to given value.

### HasSEzsignfoldertypeNameX

`func (o *EzsigntemplateListElement) HasSEzsignfoldertypeNameX() bool`

HasSEzsignfoldertypeNameX returns a boolean if a field has been set.

### GetEEzsigntemplateType

`func (o *EzsigntemplateListElement) GetEEzsigntemplateType() FieldEEzsigntemplateType`

GetEEzsigntemplateType returns the EEzsigntemplateType field if non-nil, zero value otherwise.

### GetEEzsigntemplateTypeOk

`func (o *EzsigntemplateListElement) GetEEzsigntemplateTypeOk() (*FieldEEzsigntemplateType, bool)`

GetEEzsigntemplateTypeOk returns a tuple with the EEzsigntemplateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplateType

`func (o *EzsigntemplateListElement) SetEEzsigntemplateType(v FieldEEzsigntemplateType)`

SetEEzsigntemplateType sets EEzsigntemplateType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


