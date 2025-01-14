# EzsigntemplatepackageResponseCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsigntemplatepackageID** | **int32** | The unique ID of the Ezsigntemplatepackage | 
**FkiEzsignfoldertypeID** | **int32** | The unique ID of the Ezsignfoldertype. | 
**FkiEzdoctemplatedocumentID** | Pointer to **int32** | The unique ID of the Ezdoctemplatedocument | [optional] 
**FkiLanguageID** | **int32** | The unique ID of the Language.  Valid values:  |Value|Description| |-|-| |1|French| |2|English| | 
**SEzdoctemplatedocumentNameX** | Pointer to **string** | The name of the Ezdoctemplatedocument in the language of the requester | [optional] 
**SLanguageNameX** | **string** | The Name of the Language in the language of the requester | 
**SEzsigntemplatepackageDescription** | **string** | The description of the Ezsigntemplatepackage | 
**BEzsigntemplatepackageAdminonly** | **bool** | Whether the Ezsigntemplatepackage can be accessed by admin users only (eUserType&#x3D;Normal) | 
**BEzsigntemplatepackageNeedvalidation** | **bool** | Whether the Ezsignbulksend was automatically modified and needs a manual validation | 
**BEzsigntemplatepackageIsactive** | **bool** | Whether the Ezsigntemplatepackage is active or not | 
**SEzsignfoldertypeNameX** | **string** | The name of the Ezsignfoldertype in the language of the requester | 
**BEzsigntemplatepackageEditallowed** | **bool** | Whether the Ezsigntemplatepackage if allowed to edit or not | 
**AObjEzsigntemplatepackagesigner** | [**[]EzsigntemplatepackagesignerResponseCompound**](EzsigntemplatepackagesignerResponseCompound.md) |  | 
**AObjEzsigntemplatepackagemembership** | [**[]EzsigntemplatepackagemembershipResponseCompound**](EzsigntemplatepackagemembershipResponseCompound.md) |  | 

## Methods

### NewEzsigntemplatepackageResponseCompound

`func NewEzsigntemplatepackageResponseCompound(pkiEzsigntemplatepackageID int32, fkiEzsignfoldertypeID int32, fkiLanguageID int32, sLanguageNameX string, sEzsigntemplatepackageDescription string, bEzsigntemplatepackageAdminonly bool, bEzsigntemplatepackageNeedvalidation bool, bEzsigntemplatepackageIsactive bool, sEzsignfoldertypeNameX string, bEzsigntemplatepackageEditallowed bool, aObjEzsigntemplatepackagesigner []EzsigntemplatepackagesignerResponseCompound, aObjEzsigntemplatepackagemembership []EzsigntemplatepackagemembershipResponseCompound, ) *EzsigntemplatepackageResponseCompound`

NewEzsigntemplatepackageResponseCompound instantiates a new EzsigntemplatepackageResponseCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigntemplatepackageResponseCompoundWithDefaults

`func NewEzsigntemplatepackageResponseCompoundWithDefaults() *EzsigntemplatepackageResponseCompound`

NewEzsigntemplatepackageResponseCompoundWithDefaults instantiates a new EzsigntemplatepackageResponseCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsigntemplatepackageID

`func (o *EzsigntemplatepackageResponseCompound) GetPkiEzsigntemplatepackageID() int32`

GetPkiEzsigntemplatepackageID returns the PkiEzsigntemplatepackageID field if non-nil, zero value otherwise.

### GetPkiEzsigntemplatepackageIDOk

`func (o *EzsigntemplatepackageResponseCompound) GetPkiEzsigntemplatepackageIDOk() (*int32, bool)`

GetPkiEzsigntemplatepackageIDOk returns a tuple with the PkiEzsigntemplatepackageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsigntemplatepackageID

`func (o *EzsigntemplatepackageResponseCompound) SetPkiEzsigntemplatepackageID(v int32)`

SetPkiEzsigntemplatepackageID sets PkiEzsigntemplatepackageID field to given value.


### GetFkiEzsignfoldertypeID

`func (o *EzsigntemplatepackageResponseCompound) GetFkiEzsignfoldertypeID() int32`

GetFkiEzsignfoldertypeID returns the FkiEzsignfoldertypeID field if non-nil, zero value otherwise.

### GetFkiEzsignfoldertypeIDOk

`func (o *EzsigntemplatepackageResponseCompound) GetFkiEzsignfoldertypeIDOk() (*int32, bool)`

GetFkiEzsignfoldertypeIDOk returns a tuple with the FkiEzsignfoldertypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignfoldertypeID

`func (o *EzsigntemplatepackageResponseCompound) SetFkiEzsignfoldertypeID(v int32)`

SetFkiEzsignfoldertypeID sets FkiEzsignfoldertypeID field to given value.


### GetFkiEzdoctemplatedocumentID

`func (o *EzsigntemplatepackageResponseCompound) GetFkiEzdoctemplatedocumentID() int32`

GetFkiEzdoctemplatedocumentID returns the FkiEzdoctemplatedocumentID field if non-nil, zero value otherwise.

### GetFkiEzdoctemplatedocumentIDOk

`func (o *EzsigntemplatepackageResponseCompound) GetFkiEzdoctemplatedocumentIDOk() (*int32, bool)`

GetFkiEzdoctemplatedocumentIDOk returns a tuple with the FkiEzdoctemplatedocumentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzdoctemplatedocumentID

`func (o *EzsigntemplatepackageResponseCompound) SetFkiEzdoctemplatedocumentID(v int32)`

SetFkiEzdoctemplatedocumentID sets FkiEzdoctemplatedocumentID field to given value.

### HasFkiEzdoctemplatedocumentID

`func (o *EzsigntemplatepackageResponseCompound) HasFkiEzdoctemplatedocumentID() bool`

HasFkiEzdoctemplatedocumentID returns a boolean if a field has been set.

### GetFkiLanguageID

`func (o *EzsigntemplatepackageResponseCompound) GetFkiLanguageID() int32`

GetFkiLanguageID returns the FkiLanguageID field if non-nil, zero value otherwise.

### GetFkiLanguageIDOk

`func (o *EzsigntemplatepackageResponseCompound) GetFkiLanguageIDOk() (*int32, bool)`

GetFkiLanguageIDOk returns a tuple with the FkiLanguageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiLanguageID

`func (o *EzsigntemplatepackageResponseCompound) SetFkiLanguageID(v int32)`

SetFkiLanguageID sets FkiLanguageID field to given value.


### GetSEzdoctemplatedocumentNameX

`func (o *EzsigntemplatepackageResponseCompound) GetSEzdoctemplatedocumentNameX() string`

GetSEzdoctemplatedocumentNameX returns the SEzdoctemplatedocumentNameX field if non-nil, zero value otherwise.

### GetSEzdoctemplatedocumentNameXOk

`func (o *EzsigntemplatepackageResponseCompound) GetSEzdoctemplatedocumentNameXOk() (*string, bool)`

GetSEzdoctemplatedocumentNameXOk returns a tuple with the SEzdoctemplatedocumentNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzdoctemplatedocumentNameX

`func (o *EzsigntemplatepackageResponseCompound) SetSEzdoctemplatedocumentNameX(v string)`

SetSEzdoctemplatedocumentNameX sets SEzdoctemplatedocumentNameX field to given value.

### HasSEzdoctemplatedocumentNameX

`func (o *EzsigntemplatepackageResponseCompound) HasSEzdoctemplatedocumentNameX() bool`

HasSEzdoctemplatedocumentNameX returns a boolean if a field has been set.

### GetSLanguageNameX

`func (o *EzsigntemplatepackageResponseCompound) GetSLanguageNameX() string`

GetSLanguageNameX returns the SLanguageNameX field if non-nil, zero value otherwise.

### GetSLanguageNameXOk

`func (o *EzsigntemplatepackageResponseCompound) GetSLanguageNameXOk() (*string, bool)`

GetSLanguageNameXOk returns a tuple with the SLanguageNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSLanguageNameX

`func (o *EzsigntemplatepackageResponseCompound) SetSLanguageNameX(v string)`

SetSLanguageNameX sets SLanguageNameX field to given value.


### GetSEzsigntemplatepackageDescription

`func (o *EzsigntemplatepackageResponseCompound) GetSEzsigntemplatepackageDescription() string`

GetSEzsigntemplatepackageDescription returns the SEzsigntemplatepackageDescription field if non-nil, zero value otherwise.

### GetSEzsigntemplatepackageDescriptionOk

`func (o *EzsigntemplatepackageResponseCompound) GetSEzsigntemplatepackageDescriptionOk() (*string, bool)`

GetSEzsigntemplatepackageDescriptionOk returns a tuple with the SEzsigntemplatepackageDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplatepackageDescription

`func (o *EzsigntemplatepackageResponseCompound) SetSEzsigntemplatepackageDescription(v string)`

SetSEzsigntemplatepackageDescription sets SEzsigntemplatepackageDescription field to given value.


### GetBEzsigntemplatepackageAdminonly

`func (o *EzsigntemplatepackageResponseCompound) GetBEzsigntemplatepackageAdminonly() bool`

GetBEzsigntemplatepackageAdminonly returns the BEzsigntemplatepackageAdminonly field if non-nil, zero value otherwise.

### GetBEzsigntemplatepackageAdminonlyOk

`func (o *EzsigntemplatepackageResponseCompound) GetBEzsigntemplatepackageAdminonlyOk() (*bool, bool)`

GetBEzsigntemplatepackageAdminonlyOk returns a tuple with the BEzsigntemplatepackageAdminonly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsigntemplatepackageAdminonly

`func (o *EzsigntemplatepackageResponseCompound) SetBEzsigntemplatepackageAdminonly(v bool)`

SetBEzsigntemplatepackageAdminonly sets BEzsigntemplatepackageAdminonly field to given value.


### GetBEzsigntemplatepackageNeedvalidation

`func (o *EzsigntemplatepackageResponseCompound) GetBEzsigntemplatepackageNeedvalidation() bool`

GetBEzsigntemplatepackageNeedvalidation returns the BEzsigntemplatepackageNeedvalidation field if non-nil, zero value otherwise.

### GetBEzsigntemplatepackageNeedvalidationOk

`func (o *EzsigntemplatepackageResponseCompound) GetBEzsigntemplatepackageNeedvalidationOk() (*bool, bool)`

GetBEzsigntemplatepackageNeedvalidationOk returns a tuple with the BEzsigntemplatepackageNeedvalidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsigntemplatepackageNeedvalidation

`func (o *EzsigntemplatepackageResponseCompound) SetBEzsigntemplatepackageNeedvalidation(v bool)`

SetBEzsigntemplatepackageNeedvalidation sets BEzsigntemplatepackageNeedvalidation field to given value.


### GetBEzsigntemplatepackageIsactive

`func (o *EzsigntemplatepackageResponseCompound) GetBEzsigntemplatepackageIsactive() bool`

GetBEzsigntemplatepackageIsactive returns the BEzsigntemplatepackageIsactive field if non-nil, zero value otherwise.

### GetBEzsigntemplatepackageIsactiveOk

`func (o *EzsigntemplatepackageResponseCompound) GetBEzsigntemplatepackageIsactiveOk() (*bool, bool)`

GetBEzsigntemplatepackageIsactiveOk returns a tuple with the BEzsigntemplatepackageIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsigntemplatepackageIsactive

`func (o *EzsigntemplatepackageResponseCompound) SetBEzsigntemplatepackageIsactive(v bool)`

SetBEzsigntemplatepackageIsactive sets BEzsigntemplatepackageIsactive field to given value.


### GetSEzsignfoldertypeNameX

`func (o *EzsigntemplatepackageResponseCompound) GetSEzsignfoldertypeNameX() string`

GetSEzsignfoldertypeNameX returns the SEzsignfoldertypeNameX field if non-nil, zero value otherwise.

### GetSEzsignfoldertypeNameXOk

`func (o *EzsigntemplatepackageResponseCompound) GetSEzsignfoldertypeNameXOk() (*string, bool)`

GetSEzsignfoldertypeNameXOk returns a tuple with the SEzsignfoldertypeNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignfoldertypeNameX

`func (o *EzsigntemplatepackageResponseCompound) SetSEzsignfoldertypeNameX(v string)`

SetSEzsignfoldertypeNameX sets SEzsignfoldertypeNameX field to given value.


### GetBEzsigntemplatepackageEditallowed

`func (o *EzsigntemplatepackageResponseCompound) GetBEzsigntemplatepackageEditallowed() bool`

GetBEzsigntemplatepackageEditallowed returns the BEzsigntemplatepackageEditallowed field if non-nil, zero value otherwise.

### GetBEzsigntemplatepackageEditallowedOk

`func (o *EzsigntemplatepackageResponseCompound) GetBEzsigntemplatepackageEditallowedOk() (*bool, bool)`

GetBEzsigntemplatepackageEditallowedOk returns a tuple with the BEzsigntemplatepackageEditallowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsigntemplatepackageEditallowed

`func (o *EzsigntemplatepackageResponseCompound) SetBEzsigntemplatepackageEditallowed(v bool)`

SetBEzsigntemplatepackageEditallowed sets BEzsigntemplatepackageEditallowed field to given value.


### GetAObjEzsigntemplatepackagesigner

`func (o *EzsigntemplatepackageResponseCompound) GetAObjEzsigntemplatepackagesigner() []EzsigntemplatepackagesignerResponseCompound`

GetAObjEzsigntemplatepackagesigner returns the AObjEzsigntemplatepackagesigner field if non-nil, zero value otherwise.

### GetAObjEzsigntemplatepackagesignerOk

`func (o *EzsigntemplatepackageResponseCompound) GetAObjEzsigntemplatepackagesignerOk() (*[]EzsigntemplatepackagesignerResponseCompound, bool)`

GetAObjEzsigntemplatepackagesignerOk returns a tuple with the AObjEzsigntemplatepackagesigner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEzsigntemplatepackagesigner

`func (o *EzsigntemplatepackageResponseCompound) SetAObjEzsigntemplatepackagesigner(v []EzsigntemplatepackagesignerResponseCompound)`

SetAObjEzsigntemplatepackagesigner sets AObjEzsigntemplatepackagesigner field to given value.


### GetAObjEzsigntemplatepackagemembership

`func (o *EzsigntemplatepackageResponseCompound) GetAObjEzsigntemplatepackagemembership() []EzsigntemplatepackagemembershipResponseCompound`

GetAObjEzsigntemplatepackagemembership returns the AObjEzsigntemplatepackagemembership field if non-nil, zero value otherwise.

### GetAObjEzsigntemplatepackagemembershipOk

`func (o *EzsigntemplatepackageResponseCompound) GetAObjEzsigntemplatepackagemembershipOk() (*[]EzsigntemplatepackagemembershipResponseCompound, bool)`

GetAObjEzsigntemplatepackagemembershipOk returns a tuple with the AObjEzsigntemplatepackagemembership field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEzsigntemplatepackagemembership

`func (o *EzsigntemplatepackageResponseCompound) SetAObjEzsigntemplatepackagemembership(v []EzsigntemplatepackagemembershipResponseCompound)`

SetAObjEzsigntemplatepackagemembership sets AObjEzsigntemplatepackagemembership field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


