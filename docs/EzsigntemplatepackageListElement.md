# EzsigntemplatepackageListElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsigntemplatepackageID** | **int32** | The unique ID of the Ezsigntemplatepackage | 
**FkiEzsignfoldertypeID** | **int32** | The unique ID of the Ezsignfoldertype. | 
**FkiLanguageID** | **int32** | The unique ID of the Language.  Valid values:  |Value|Description| |-|-| |1|French| |2|English| | 
**SEzsigntemplatepackageDescription** | **string** | The description of the Ezsigntemplatepackage | 
**BEzsigntemplatepackageNeedvalidation** | **bool** | Whether the Ezsignbulksend was automatically modified and needs a manual validation | 
**IEzsigntemplatepackagemembership** | **int32** | The total number of Ezsigntemplatepackagemembership in the Ezsigntemplatepackage | 
**SEzsignfoldertypeNameX** | **string** | The name of the Ezsignfoldertype in the language of the requester | 
**BEzsigntemplatepackageIsactive** | **bool** | Whether the Ezsigntemplatepackage is active or not | 

## Methods

### NewEzsigntemplatepackageListElement

`func NewEzsigntemplatepackageListElement(pkiEzsigntemplatepackageID int32, fkiEzsignfoldertypeID int32, fkiLanguageID int32, sEzsigntemplatepackageDescription string, bEzsigntemplatepackageNeedvalidation bool, iEzsigntemplatepackagemembership int32, sEzsignfoldertypeNameX string, bEzsigntemplatepackageIsactive bool, ) *EzsigntemplatepackageListElement`

NewEzsigntemplatepackageListElement instantiates a new EzsigntemplatepackageListElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigntemplatepackageListElementWithDefaults

`func NewEzsigntemplatepackageListElementWithDefaults() *EzsigntemplatepackageListElement`

NewEzsigntemplatepackageListElementWithDefaults instantiates a new EzsigntemplatepackageListElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsigntemplatepackageID

`func (o *EzsigntemplatepackageListElement) GetPkiEzsigntemplatepackageID() int32`

GetPkiEzsigntemplatepackageID returns the PkiEzsigntemplatepackageID field if non-nil, zero value otherwise.

### GetPkiEzsigntemplatepackageIDOk

`func (o *EzsigntemplatepackageListElement) GetPkiEzsigntemplatepackageIDOk() (*int32, bool)`

GetPkiEzsigntemplatepackageIDOk returns a tuple with the PkiEzsigntemplatepackageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsigntemplatepackageID

`func (o *EzsigntemplatepackageListElement) SetPkiEzsigntemplatepackageID(v int32)`

SetPkiEzsigntemplatepackageID sets PkiEzsigntemplatepackageID field to given value.


### GetFkiEzsignfoldertypeID

`func (o *EzsigntemplatepackageListElement) GetFkiEzsignfoldertypeID() int32`

GetFkiEzsignfoldertypeID returns the FkiEzsignfoldertypeID field if non-nil, zero value otherwise.

### GetFkiEzsignfoldertypeIDOk

`func (o *EzsigntemplatepackageListElement) GetFkiEzsignfoldertypeIDOk() (*int32, bool)`

GetFkiEzsignfoldertypeIDOk returns a tuple with the FkiEzsignfoldertypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignfoldertypeID

`func (o *EzsigntemplatepackageListElement) SetFkiEzsignfoldertypeID(v int32)`

SetFkiEzsignfoldertypeID sets FkiEzsignfoldertypeID field to given value.


### GetFkiLanguageID

`func (o *EzsigntemplatepackageListElement) GetFkiLanguageID() int32`

GetFkiLanguageID returns the FkiLanguageID field if non-nil, zero value otherwise.

### GetFkiLanguageIDOk

`func (o *EzsigntemplatepackageListElement) GetFkiLanguageIDOk() (*int32, bool)`

GetFkiLanguageIDOk returns a tuple with the FkiLanguageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiLanguageID

`func (o *EzsigntemplatepackageListElement) SetFkiLanguageID(v int32)`

SetFkiLanguageID sets FkiLanguageID field to given value.


### GetSEzsigntemplatepackageDescription

`func (o *EzsigntemplatepackageListElement) GetSEzsigntemplatepackageDescription() string`

GetSEzsigntemplatepackageDescription returns the SEzsigntemplatepackageDescription field if non-nil, zero value otherwise.

### GetSEzsigntemplatepackageDescriptionOk

`func (o *EzsigntemplatepackageListElement) GetSEzsigntemplatepackageDescriptionOk() (*string, bool)`

GetSEzsigntemplatepackageDescriptionOk returns a tuple with the SEzsigntemplatepackageDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplatepackageDescription

`func (o *EzsigntemplatepackageListElement) SetSEzsigntemplatepackageDescription(v string)`

SetSEzsigntemplatepackageDescription sets SEzsigntemplatepackageDescription field to given value.


### GetBEzsigntemplatepackageNeedvalidation

`func (o *EzsigntemplatepackageListElement) GetBEzsigntemplatepackageNeedvalidation() bool`

GetBEzsigntemplatepackageNeedvalidation returns the BEzsigntemplatepackageNeedvalidation field if non-nil, zero value otherwise.

### GetBEzsigntemplatepackageNeedvalidationOk

`func (o *EzsigntemplatepackageListElement) GetBEzsigntemplatepackageNeedvalidationOk() (*bool, bool)`

GetBEzsigntemplatepackageNeedvalidationOk returns a tuple with the BEzsigntemplatepackageNeedvalidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsigntemplatepackageNeedvalidation

`func (o *EzsigntemplatepackageListElement) SetBEzsigntemplatepackageNeedvalidation(v bool)`

SetBEzsigntemplatepackageNeedvalidation sets BEzsigntemplatepackageNeedvalidation field to given value.


### GetIEzsigntemplatepackagemembership

`func (o *EzsigntemplatepackageListElement) GetIEzsigntemplatepackagemembership() int32`

GetIEzsigntemplatepackagemembership returns the IEzsigntemplatepackagemembership field if non-nil, zero value otherwise.

### GetIEzsigntemplatepackagemembershipOk

`func (o *EzsigntemplatepackageListElement) GetIEzsigntemplatepackagemembershipOk() (*int32, bool)`

GetIEzsigntemplatepackagemembershipOk returns a tuple with the IEzsigntemplatepackagemembership field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplatepackagemembership

`func (o *EzsigntemplatepackageListElement) SetIEzsigntemplatepackagemembership(v int32)`

SetIEzsigntemplatepackagemembership sets IEzsigntemplatepackagemembership field to given value.


### GetSEzsignfoldertypeNameX

`func (o *EzsigntemplatepackageListElement) GetSEzsignfoldertypeNameX() string`

GetSEzsignfoldertypeNameX returns the SEzsignfoldertypeNameX field if non-nil, zero value otherwise.

### GetSEzsignfoldertypeNameXOk

`func (o *EzsigntemplatepackageListElement) GetSEzsignfoldertypeNameXOk() (*string, bool)`

GetSEzsignfoldertypeNameXOk returns a tuple with the SEzsignfoldertypeNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignfoldertypeNameX

`func (o *EzsigntemplatepackageListElement) SetSEzsignfoldertypeNameX(v string)`

SetSEzsignfoldertypeNameX sets SEzsignfoldertypeNameX field to given value.


### GetBEzsigntemplatepackageIsactive

`func (o *EzsigntemplatepackageListElement) GetBEzsigntemplatepackageIsactive() bool`

GetBEzsigntemplatepackageIsactive returns the BEzsigntemplatepackageIsactive field if non-nil, zero value otherwise.

### GetBEzsigntemplatepackageIsactiveOk

`func (o *EzsigntemplatepackageListElement) GetBEzsigntemplatepackageIsactiveOk() (*bool, bool)`

GetBEzsigntemplatepackageIsactiveOk returns a tuple with the BEzsigntemplatepackageIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsigntemplatepackageIsactive

`func (o *EzsigntemplatepackageListElement) SetBEzsigntemplatepackageIsactive(v bool)`

SetBEzsigntemplatepackageIsactive sets BEzsigntemplatepackageIsactive field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


