# EzsigntemplatepackageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsigntemplatepackageID** | **int32** | The unique ID of the Ezsigntemplatepackage | 
**FkiEzsignfoldertypeID** | **int32** | The unique ID of the Ezsignfoldertype. | 
**FkiLanguageID** | **int32** | The unique ID of the Language.  Valid values:  |Value|Description| |-|-| |1|French| |2|English| | 
**SLanguageNameX** | **string** | The Name of the Language in the language of the requester | 
**SEzsigntemplatepackageDescription** | **string** | The description of the Ezsigntemplatepackage | 
**BEzsigntemplatepackageAdminonly** | **bool** | Whether the Ezsigntemplatepackage can be accessed by admin users only (eUserType&#x3D;Normal) | 
**BEzsigntemplatepackageNeedvalidation** | **bool** | Whether the Ezsignbulksend was automatically modified and needs a manual validation | 
**BEzsigntemplatepackageIsactive** | **bool** | Whether the Ezsigntemplatepackage is active or not | 
**SEzsignfoldertypeNameX** | **string** | The name of the Ezsignfoldertype in the language of the requester | 

## Methods

### NewEzsigntemplatepackageResponse

`func NewEzsigntemplatepackageResponse(pkiEzsigntemplatepackageID int32, fkiEzsignfoldertypeID int32, fkiLanguageID int32, sLanguageNameX string, sEzsigntemplatepackageDescription string, bEzsigntemplatepackageAdminonly bool, bEzsigntemplatepackageNeedvalidation bool, bEzsigntemplatepackageIsactive bool, sEzsignfoldertypeNameX string, ) *EzsigntemplatepackageResponse`

NewEzsigntemplatepackageResponse instantiates a new EzsigntemplatepackageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigntemplatepackageResponseWithDefaults

`func NewEzsigntemplatepackageResponseWithDefaults() *EzsigntemplatepackageResponse`

NewEzsigntemplatepackageResponseWithDefaults instantiates a new EzsigntemplatepackageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsigntemplatepackageID

`func (o *EzsigntemplatepackageResponse) GetPkiEzsigntemplatepackageID() int32`

GetPkiEzsigntemplatepackageID returns the PkiEzsigntemplatepackageID field if non-nil, zero value otherwise.

### GetPkiEzsigntemplatepackageIDOk

`func (o *EzsigntemplatepackageResponse) GetPkiEzsigntemplatepackageIDOk() (*int32, bool)`

GetPkiEzsigntemplatepackageIDOk returns a tuple with the PkiEzsigntemplatepackageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsigntemplatepackageID

`func (o *EzsigntemplatepackageResponse) SetPkiEzsigntemplatepackageID(v int32)`

SetPkiEzsigntemplatepackageID sets PkiEzsigntemplatepackageID field to given value.


### GetFkiEzsignfoldertypeID

`func (o *EzsigntemplatepackageResponse) GetFkiEzsignfoldertypeID() int32`

GetFkiEzsignfoldertypeID returns the FkiEzsignfoldertypeID field if non-nil, zero value otherwise.

### GetFkiEzsignfoldertypeIDOk

`func (o *EzsigntemplatepackageResponse) GetFkiEzsignfoldertypeIDOk() (*int32, bool)`

GetFkiEzsignfoldertypeIDOk returns a tuple with the FkiEzsignfoldertypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignfoldertypeID

`func (o *EzsigntemplatepackageResponse) SetFkiEzsignfoldertypeID(v int32)`

SetFkiEzsignfoldertypeID sets FkiEzsignfoldertypeID field to given value.


### GetFkiLanguageID

`func (o *EzsigntemplatepackageResponse) GetFkiLanguageID() int32`

GetFkiLanguageID returns the FkiLanguageID field if non-nil, zero value otherwise.

### GetFkiLanguageIDOk

`func (o *EzsigntemplatepackageResponse) GetFkiLanguageIDOk() (*int32, bool)`

GetFkiLanguageIDOk returns a tuple with the FkiLanguageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiLanguageID

`func (o *EzsigntemplatepackageResponse) SetFkiLanguageID(v int32)`

SetFkiLanguageID sets FkiLanguageID field to given value.


### GetSLanguageNameX

`func (o *EzsigntemplatepackageResponse) GetSLanguageNameX() string`

GetSLanguageNameX returns the SLanguageNameX field if non-nil, zero value otherwise.

### GetSLanguageNameXOk

`func (o *EzsigntemplatepackageResponse) GetSLanguageNameXOk() (*string, bool)`

GetSLanguageNameXOk returns a tuple with the SLanguageNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSLanguageNameX

`func (o *EzsigntemplatepackageResponse) SetSLanguageNameX(v string)`

SetSLanguageNameX sets SLanguageNameX field to given value.


### GetSEzsigntemplatepackageDescription

`func (o *EzsigntemplatepackageResponse) GetSEzsigntemplatepackageDescription() string`

GetSEzsigntemplatepackageDescription returns the SEzsigntemplatepackageDescription field if non-nil, zero value otherwise.

### GetSEzsigntemplatepackageDescriptionOk

`func (o *EzsigntemplatepackageResponse) GetSEzsigntemplatepackageDescriptionOk() (*string, bool)`

GetSEzsigntemplatepackageDescriptionOk returns a tuple with the SEzsigntemplatepackageDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplatepackageDescription

`func (o *EzsigntemplatepackageResponse) SetSEzsigntemplatepackageDescription(v string)`

SetSEzsigntemplatepackageDescription sets SEzsigntemplatepackageDescription field to given value.


### GetBEzsigntemplatepackageAdminonly

`func (o *EzsigntemplatepackageResponse) GetBEzsigntemplatepackageAdminonly() bool`

GetBEzsigntemplatepackageAdminonly returns the BEzsigntemplatepackageAdminonly field if non-nil, zero value otherwise.

### GetBEzsigntemplatepackageAdminonlyOk

`func (o *EzsigntemplatepackageResponse) GetBEzsigntemplatepackageAdminonlyOk() (*bool, bool)`

GetBEzsigntemplatepackageAdminonlyOk returns a tuple with the BEzsigntemplatepackageAdminonly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsigntemplatepackageAdminonly

`func (o *EzsigntemplatepackageResponse) SetBEzsigntemplatepackageAdminonly(v bool)`

SetBEzsigntemplatepackageAdminonly sets BEzsigntemplatepackageAdminonly field to given value.


### GetBEzsigntemplatepackageNeedvalidation

`func (o *EzsigntemplatepackageResponse) GetBEzsigntemplatepackageNeedvalidation() bool`

GetBEzsigntemplatepackageNeedvalidation returns the BEzsigntemplatepackageNeedvalidation field if non-nil, zero value otherwise.

### GetBEzsigntemplatepackageNeedvalidationOk

`func (o *EzsigntemplatepackageResponse) GetBEzsigntemplatepackageNeedvalidationOk() (*bool, bool)`

GetBEzsigntemplatepackageNeedvalidationOk returns a tuple with the BEzsigntemplatepackageNeedvalidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsigntemplatepackageNeedvalidation

`func (o *EzsigntemplatepackageResponse) SetBEzsigntemplatepackageNeedvalidation(v bool)`

SetBEzsigntemplatepackageNeedvalidation sets BEzsigntemplatepackageNeedvalidation field to given value.


### GetBEzsigntemplatepackageIsactive

`func (o *EzsigntemplatepackageResponse) GetBEzsigntemplatepackageIsactive() bool`

GetBEzsigntemplatepackageIsactive returns the BEzsigntemplatepackageIsactive field if non-nil, zero value otherwise.

### GetBEzsigntemplatepackageIsactiveOk

`func (o *EzsigntemplatepackageResponse) GetBEzsigntemplatepackageIsactiveOk() (*bool, bool)`

GetBEzsigntemplatepackageIsactiveOk returns a tuple with the BEzsigntemplatepackageIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsigntemplatepackageIsactive

`func (o *EzsigntemplatepackageResponse) SetBEzsigntemplatepackageIsactive(v bool)`

SetBEzsigntemplatepackageIsactive sets BEzsigntemplatepackageIsactive field to given value.


### GetSEzsignfoldertypeNameX

`func (o *EzsigntemplatepackageResponse) GetSEzsignfoldertypeNameX() string`

GetSEzsignfoldertypeNameX returns the SEzsignfoldertypeNameX field if non-nil, zero value otherwise.

### GetSEzsignfoldertypeNameXOk

`func (o *EzsigntemplatepackageResponse) GetSEzsignfoldertypeNameXOk() (*string, bool)`

GetSEzsignfoldertypeNameXOk returns a tuple with the SEzsignfoldertypeNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignfoldertypeNameX

`func (o *EzsigntemplatepackageResponse) SetSEzsignfoldertypeNameX(v string)`

SetSEzsignfoldertypeNameX sets SEzsignfoldertypeNameX field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


