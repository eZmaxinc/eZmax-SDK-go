# EzsigntemplatepackageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsigntemplatepackageID** | Pointer to **int32** | The unique ID of the Ezsigntemplatepackage | [optional] 
**FkiEzsignfoldertypeID** | **int32** | The unique ID of the Ezsignfoldertype. | 
**FkiEzdoctemplatedocumentID** | Pointer to **int32** | The unique ID of the Ezdoctemplatedocument | [optional] 
**FkiLanguageID** | **int32** | The unique ID of the Language.  Valid values:  |Value|Description| |-|-| |1|French| |2|English| | 
**SEzsigntemplatepackageDescription** | **string** | The description of the Ezsigntemplatepackage | 
**BEzsigntemplatepackageAdminonly** | **bool** | Whether the Ezsigntemplatepackage can be accessed by admin users only (eUserType&#x3D;Normal) | 
**BEzsigntemplatepackageIsactive** | **bool** | Whether the Ezsigntemplatepackage is active or not | 

## Methods

### NewEzsigntemplatepackageRequest

`func NewEzsigntemplatepackageRequest(fkiEzsignfoldertypeID int32, fkiLanguageID int32, sEzsigntemplatepackageDescription string, bEzsigntemplatepackageAdminonly bool, bEzsigntemplatepackageIsactive bool, ) *EzsigntemplatepackageRequest`

NewEzsigntemplatepackageRequest instantiates a new EzsigntemplatepackageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigntemplatepackageRequestWithDefaults

`func NewEzsigntemplatepackageRequestWithDefaults() *EzsigntemplatepackageRequest`

NewEzsigntemplatepackageRequestWithDefaults instantiates a new EzsigntemplatepackageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsigntemplatepackageID

`func (o *EzsigntemplatepackageRequest) GetPkiEzsigntemplatepackageID() int32`

GetPkiEzsigntemplatepackageID returns the PkiEzsigntemplatepackageID field if non-nil, zero value otherwise.

### GetPkiEzsigntemplatepackageIDOk

`func (o *EzsigntemplatepackageRequest) GetPkiEzsigntemplatepackageIDOk() (*int32, bool)`

GetPkiEzsigntemplatepackageIDOk returns a tuple with the PkiEzsigntemplatepackageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsigntemplatepackageID

`func (o *EzsigntemplatepackageRequest) SetPkiEzsigntemplatepackageID(v int32)`

SetPkiEzsigntemplatepackageID sets PkiEzsigntemplatepackageID field to given value.

### HasPkiEzsigntemplatepackageID

`func (o *EzsigntemplatepackageRequest) HasPkiEzsigntemplatepackageID() bool`

HasPkiEzsigntemplatepackageID returns a boolean if a field has been set.

### GetFkiEzsignfoldertypeID

`func (o *EzsigntemplatepackageRequest) GetFkiEzsignfoldertypeID() int32`

GetFkiEzsignfoldertypeID returns the FkiEzsignfoldertypeID field if non-nil, zero value otherwise.

### GetFkiEzsignfoldertypeIDOk

`func (o *EzsigntemplatepackageRequest) GetFkiEzsignfoldertypeIDOk() (*int32, bool)`

GetFkiEzsignfoldertypeIDOk returns a tuple with the FkiEzsignfoldertypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignfoldertypeID

`func (o *EzsigntemplatepackageRequest) SetFkiEzsignfoldertypeID(v int32)`

SetFkiEzsignfoldertypeID sets FkiEzsignfoldertypeID field to given value.


### GetFkiEzdoctemplatedocumentID

`func (o *EzsigntemplatepackageRequest) GetFkiEzdoctemplatedocumentID() int32`

GetFkiEzdoctemplatedocumentID returns the FkiEzdoctemplatedocumentID field if non-nil, zero value otherwise.

### GetFkiEzdoctemplatedocumentIDOk

`func (o *EzsigntemplatepackageRequest) GetFkiEzdoctemplatedocumentIDOk() (*int32, bool)`

GetFkiEzdoctemplatedocumentIDOk returns a tuple with the FkiEzdoctemplatedocumentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzdoctemplatedocumentID

`func (o *EzsigntemplatepackageRequest) SetFkiEzdoctemplatedocumentID(v int32)`

SetFkiEzdoctemplatedocumentID sets FkiEzdoctemplatedocumentID field to given value.

### HasFkiEzdoctemplatedocumentID

`func (o *EzsigntemplatepackageRequest) HasFkiEzdoctemplatedocumentID() bool`

HasFkiEzdoctemplatedocumentID returns a boolean if a field has been set.

### GetFkiLanguageID

`func (o *EzsigntemplatepackageRequest) GetFkiLanguageID() int32`

GetFkiLanguageID returns the FkiLanguageID field if non-nil, zero value otherwise.

### GetFkiLanguageIDOk

`func (o *EzsigntemplatepackageRequest) GetFkiLanguageIDOk() (*int32, bool)`

GetFkiLanguageIDOk returns a tuple with the FkiLanguageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiLanguageID

`func (o *EzsigntemplatepackageRequest) SetFkiLanguageID(v int32)`

SetFkiLanguageID sets FkiLanguageID field to given value.


### GetSEzsigntemplatepackageDescription

`func (o *EzsigntemplatepackageRequest) GetSEzsigntemplatepackageDescription() string`

GetSEzsigntemplatepackageDescription returns the SEzsigntemplatepackageDescription field if non-nil, zero value otherwise.

### GetSEzsigntemplatepackageDescriptionOk

`func (o *EzsigntemplatepackageRequest) GetSEzsigntemplatepackageDescriptionOk() (*string, bool)`

GetSEzsigntemplatepackageDescriptionOk returns a tuple with the SEzsigntemplatepackageDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplatepackageDescription

`func (o *EzsigntemplatepackageRequest) SetSEzsigntemplatepackageDescription(v string)`

SetSEzsigntemplatepackageDescription sets SEzsigntemplatepackageDescription field to given value.


### GetBEzsigntemplatepackageAdminonly

`func (o *EzsigntemplatepackageRequest) GetBEzsigntemplatepackageAdminonly() bool`

GetBEzsigntemplatepackageAdminonly returns the BEzsigntemplatepackageAdminonly field if non-nil, zero value otherwise.

### GetBEzsigntemplatepackageAdminonlyOk

`func (o *EzsigntemplatepackageRequest) GetBEzsigntemplatepackageAdminonlyOk() (*bool, bool)`

GetBEzsigntemplatepackageAdminonlyOk returns a tuple with the BEzsigntemplatepackageAdminonly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsigntemplatepackageAdminonly

`func (o *EzsigntemplatepackageRequest) SetBEzsigntemplatepackageAdminonly(v bool)`

SetBEzsigntemplatepackageAdminonly sets BEzsigntemplatepackageAdminonly field to given value.


### GetBEzsigntemplatepackageIsactive

`func (o *EzsigntemplatepackageRequest) GetBEzsigntemplatepackageIsactive() bool`

GetBEzsigntemplatepackageIsactive returns the BEzsigntemplatepackageIsactive field if non-nil, zero value otherwise.

### GetBEzsigntemplatepackageIsactiveOk

`func (o *EzsigntemplatepackageRequest) GetBEzsigntemplatepackageIsactiveOk() (*bool, bool)`

GetBEzsigntemplatepackageIsactiveOk returns a tuple with the BEzsigntemplatepackageIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsigntemplatepackageIsactive

`func (o *EzsigntemplatepackageRequest) SetBEzsigntemplatepackageIsactive(v bool)`

SetBEzsigntemplatepackageIsactive sets BEzsigntemplatepackageIsactive field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


