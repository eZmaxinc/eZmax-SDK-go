# EzdoctemplatedocumentResponseCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzdoctemplatedocumentID** | **int32** | The unique ID of the Ezdoctemplatedocument | 
**FkiLanguageID** | **int32** | The unique ID of the Language.  Valid values:  |Value|Description| |-|-| |1|French| |2|English| | 
**FkiEzsignfoldertypeID** | Pointer to **int32** | The unique ID of the Ezsignfoldertype. | [optional] 
**FkiEzdoctemplatetypeID** | **int32** | The unique ID of the Ezdoctemplatetype | 
**FkiEzdoctemplatefieldtypecategoryID** | **int32** | The unique ID of the Ezdoctemplatefieldtypecategory | 
**EEzdoctemplatedocumentPrivacylevel** | Pointer to [**FieldEEzdoctemplatedocumentPrivacylevel**](FieldEEzdoctemplatedocumentPrivacylevel.md) |  | [optional] 
**BEzdoctemplatedocumentIsactive** | **bool** | Whether the ezdoctemplatedocument is active or not | 
**ObjEzdoctemplatedocumentName** | [**MultilingualEzdoctemplatedocumentName**](MultilingualEzdoctemplatedocumentName.md) |  | 
**SEzdoctemplatedocumentNameX** | Pointer to **string** | The name of the Ezdoctemplatedocument in the language of the requester | [optional] 
**SEzsignfoldertypeNameX** | Pointer to **string** | The name of the Ezsignfoldertype in the language of the requester | [optional] 
**SEzdoctemplatefieldtypecategoryDescriptionX** | **string** | The description of the Ezdoctemplatefieldtypecategory in the language of the requester | 
**SEzdoctemplatetypeDescriptionX** | **string** | The description of the Ezdoctemplatetype in the language of the requester | 

## Methods

### NewEzdoctemplatedocumentResponseCompound

`func NewEzdoctemplatedocumentResponseCompound(pkiEzdoctemplatedocumentID int32, fkiLanguageID int32, fkiEzdoctemplatetypeID int32, fkiEzdoctemplatefieldtypecategoryID int32, bEzdoctemplatedocumentIsactive bool, objEzdoctemplatedocumentName MultilingualEzdoctemplatedocumentName, sEzdoctemplatefieldtypecategoryDescriptionX string, sEzdoctemplatetypeDescriptionX string, ) *EzdoctemplatedocumentResponseCompound`

NewEzdoctemplatedocumentResponseCompound instantiates a new EzdoctemplatedocumentResponseCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzdoctemplatedocumentResponseCompoundWithDefaults

`func NewEzdoctemplatedocumentResponseCompoundWithDefaults() *EzdoctemplatedocumentResponseCompound`

NewEzdoctemplatedocumentResponseCompoundWithDefaults instantiates a new EzdoctemplatedocumentResponseCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzdoctemplatedocumentID

`func (o *EzdoctemplatedocumentResponseCompound) GetPkiEzdoctemplatedocumentID() int32`

GetPkiEzdoctemplatedocumentID returns the PkiEzdoctemplatedocumentID field if non-nil, zero value otherwise.

### GetPkiEzdoctemplatedocumentIDOk

`func (o *EzdoctemplatedocumentResponseCompound) GetPkiEzdoctemplatedocumentIDOk() (*int32, bool)`

GetPkiEzdoctemplatedocumentIDOk returns a tuple with the PkiEzdoctemplatedocumentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzdoctemplatedocumentID

`func (o *EzdoctemplatedocumentResponseCompound) SetPkiEzdoctemplatedocumentID(v int32)`

SetPkiEzdoctemplatedocumentID sets PkiEzdoctemplatedocumentID field to given value.


### GetFkiLanguageID

`func (o *EzdoctemplatedocumentResponseCompound) GetFkiLanguageID() int32`

GetFkiLanguageID returns the FkiLanguageID field if non-nil, zero value otherwise.

### GetFkiLanguageIDOk

`func (o *EzdoctemplatedocumentResponseCompound) GetFkiLanguageIDOk() (*int32, bool)`

GetFkiLanguageIDOk returns a tuple with the FkiLanguageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiLanguageID

`func (o *EzdoctemplatedocumentResponseCompound) SetFkiLanguageID(v int32)`

SetFkiLanguageID sets FkiLanguageID field to given value.


### GetFkiEzsignfoldertypeID

`func (o *EzdoctemplatedocumentResponseCompound) GetFkiEzsignfoldertypeID() int32`

GetFkiEzsignfoldertypeID returns the FkiEzsignfoldertypeID field if non-nil, zero value otherwise.

### GetFkiEzsignfoldertypeIDOk

`func (o *EzdoctemplatedocumentResponseCompound) GetFkiEzsignfoldertypeIDOk() (*int32, bool)`

GetFkiEzsignfoldertypeIDOk returns a tuple with the FkiEzsignfoldertypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignfoldertypeID

`func (o *EzdoctemplatedocumentResponseCompound) SetFkiEzsignfoldertypeID(v int32)`

SetFkiEzsignfoldertypeID sets FkiEzsignfoldertypeID field to given value.

### HasFkiEzsignfoldertypeID

`func (o *EzdoctemplatedocumentResponseCompound) HasFkiEzsignfoldertypeID() bool`

HasFkiEzsignfoldertypeID returns a boolean if a field has been set.

### GetFkiEzdoctemplatetypeID

`func (o *EzdoctemplatedocumentResponseCompound) GetFkiEzdoctemplatetypeID() int32`

GetFkiEzdoctemplatetypeID returns the FkiEzdoctemplatetypeID field if non-nil, zero value otherwise.

### GetFkiEzdoctemplatetypeIDOk

`func (o *EzdoctemplatedocumentResponseCompound) GetFkiEzdoctemplatetypeIDOk() (*int32, bool)`

GetFkiEzdoctemplatetypeIDOk returns a tuple with the FkiEzdoctemplatetypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzdoctemplatetypeID

`func (o *EzdoctemplatedocumentResponseCompound) SetFkiEzdoctemplatetypeID(v int32)`

SetFkiEzdoctemplatetypeID sets FkiEzdoctemplatetypeID field to given value.


### GetFkiEzdoctemplatefieldtypecategoryID

`func (o *EzdoctemplatedocumentResponseCompound) GetFkiEzdoctemplatefieldtypecategoryID() int32`

GetFkiEzdoctemplatefieldtypecategoryID returns the FkiEzdoctemplatefieldtypecategoryID field if non-nil, zero value otherwise.

### GetFkiEzdoctemplatefieldtypecategoryIDOk

`func (o *EzdoctemplatedocumentResponseCompound) GetFkiEzdoctemplatefieldtypecategoryIDOk() (*int32, bool)`

GetFkiEzdoctemplatefieldtypecategoryIDOk returns a tuple with the FkiEzdoctemplatefieldtypecategoryID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzdoctemplatefieldtypecategoryID

`func (o *EzdoctemplatedocumentResponseCompound) SetFkiEzdoctemplatefieldtypecategoryID(v int32)`

SetFkiEzdoctemplatefieldtypecategoryID sets FkiEzdoctemplatefieldtypecategoryID field to given value.


### GetEEzdoctemplatedocumentPrivacylevel

`func (o *EzdoctemplatedocumentResponseCompound) GetEEzdoctemplatedocumentPrivacylevel() FieldEEzdoctemplatedocumentPrivacylevel`

GetEEzdoctemplatedocumentPrivacylevel returns the EEzdoctemplatedocumentPrivacylevel field if non-nil, zero value otherwise.

### GetEEzdoctemplatedocumentPrivacylevelOk

`func (o *EzdoctemplatedocumentResponseCompound) GetEEzdoctemplatedocumentPrivacylevelOk() (*FieldEEzdoctemplatedocumentPrivacylevel, bool)`

GetEEzdoctemplatedocumentPrivacylevelOk returns a tuple with the EEzdoctemplatedocumentPrivacylevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzdoctemplatedocumentPrivacylevel

`func (o *EzdoctemplatedocumentResponseCompound) SetEEzdoctemplatedocumentPrivacylevel(v FieldEEzdoctemplatedocumentPrivacylevel)`

SetEEzdoctemplatedocumentPrivacylevel sets EEzdoctemplatedocumentPrivacylevel field to given value.

### HasEEzdoctemplatedocumentPrivacylevel

`func (o *EzdoctemplatedocumentResponseCompound) HasEEzdoctemplatedocumentPrivacylevel() bool`

HasEEzdoctemplatedocumentPrivacylevel returns a boolean if a field has been set.

### GetBEzdoctemplatedocumentIsactive

`func (o *EzdoctemplatedocumentResponseCompound) GetBEzdoctemplatedocumentIsactive() bool`

GetBEzdoctemplatedocumentIsactive returns the BEzdoctemplatedocumentIsactive field if non-nil, zero value otherwise.

### GetBEzdoctemplatedocumentIsactiveOk

`func (o *EzdoctemplatedocumentResponseCompound) GetBEzdoctemplatedocumentIsactiveOk() (*bool, bool)`

GetBEzdoctemplatedocumentIsactiveOk returns a tuple with the BEzdoctemplatedocumentIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzdoctemplatedocumentIsactive

`func (o *EzdoctemplatedocumentResponseCompound) SetBEzdoctemplatedocumentIsactive(v bool)`

SetBEzdoctemplatedocumentIsactive sets BEzdoctemplatedocumentIsactive field to given value.


### GetObjEzdoctemplatedocumentName

`func (o *EzdoctemplatedocumentResponseCompound) GetObjEzdoctemplatedocumentName() MultilingualEzdoctemplatedocumentName`

GetObjEzdoctemplatedocumentName returns the ObjEzdoctemplatedocumentName field if non-nil, zero value otherwise.

### GetObjEzdoctemplatedocumentNameOk

`func (o *EzdoctemplatedocumentResponseCompound) GetObjEzdoctemplatedocumentNameOk() (*MultilingualEzdoctemplatedocumentName, bool)`

GetObjEzdoctemplatedocumentNameOk returns a tuple with the ObjEzdoctemplatedocumentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjEzdoctemplatedocumentName

`func (o *EzdoctemplatedocumentResponseCompound) SetObjEzdoctemplatedocumentName(v MultilingualEzdoctemplatedocumentName)`

SetObjEzdoctemplatedocumentName sets ObjEzdoctemplatedocumentName field to given value.


### GetSEzdoctemplatedocumentNameX

`func (o *EzdoctemplatedocumentResponseCompound) GetSEzdoctemplatedocumentNameX() string`

GetSEzdoctemplatedocumentNameX returns the SEzdoctemplatedocumentNameX field if non-nil, zero value otherwise.

### GetSEzdoctemplatedocumentNameXOk

`func (o *EzdoctemplatedocumentResponseCompound) GetSEzdoctemplatedocumentNameXOk() (*string, bool)`

GetSEzdoctemplatedocumentNameXOk returns a tuple with the SEzdoctemplatedocumentNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzdoctemplatedocumentNameX

`func (o *EzdoctemplatedocumentResponseCompound) SetSEzdoctemplatedocumentNameX(v string)`

SetSEzdoctemplatedocumentNameX sets SEzdoctemplatedocumentNameX field to given value.

### HasSEzdoctemplatedocumentNameX

`func (o *EzdoctemplatedocumentResponseCompound) HasSEzdoctemplatedocumentNameX() bool`

HasSEzdoctemplatedocumentNameX returns a boolean if a field has been set.

### GetSEzsignfoldertypeNameX

`func (o *EzdoctemplatedocumentResponseCompound) GetSEzsignfoldertypeNameX() string`

GetSEzsignfoldertypeNameX returns the SEzsignfoldertypeNameX field if non-nil, zero value otherwise.

### GetSEzsignfoldertypeNameXOk

`func (o *EzdoctemplatedocumentResponseCompound) GetSEzsignfoldertypeNameXOk() (*string, bool)`

GetSEzsignfoldertypeNameXOk returns a tuple with the SEzsignfoldertypeNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignfoldertypeNameX

`func (o *EzdoctemplatedocumentResponseCompound) SetSEzsignfoldertypeNameX(v string)`

SetSEzsignfoldertypeNameX sets SEzsignfoldertypeNameX field to given value.

### HasSEzsignfoldertypeNameX

`func (o *EzdoctemplatedocumentResponseCompound) HasSEzsignfoldertypeNameX() bool`

HasSEzsignfoldertypeNameX returns a boolean if a field has been set.

### GetSEzdoctemplatefieldtypecategoryDescriptionX

`func (o *EzdoctemplatedocumentResponseCompound) GetSEzdoctemplatefieldtypecategoryDescriptionX() string`

GetSEzdoctemplatefieldtypecategoryDescriptionX returns the SEzdoctemplatefieldtypecategoryDescriptionX field if non-nil, zero value otherwise.

### GetSEzdoctemplatefieldtypecategoryDescriptionXOk

`func (o *EzdoctemplatedocumentResponseCompound) GetSEzdoctemplatefieldtypecategoryDescriptionXOk() (*string, bool)`

GetSEzdoctemplatefieldtypecategoryDescriptionXOk returns a tuple with the SEzdoctemplatefieldtypecategoryDescriptionX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzdoctemplatefieldtypecategoryDescriptionX

`func (o *EzdoctemplatedocumentResponseCompound) SetSEzdoctemplatefieldtypecategoryDescriptionX(v string)`

SetSEzdoctemplatefieldtypecategoryDescriptionX sets SEzdoctemplatefieldtypecategoryDescriptionX field to given value.


### GetSEzdoctemplatetypeDescriptionX

`func (o *EzdoctemplatedocumentResponseCompound) GetSEzdoctemplatetypeDescriptionX() string`

GetSEzdoctemplatetypeDescriptionX returns the SEzdoctemplatetypeDescriptionX field if non-nil, zero value otherwise.

### GetSEzdoctemplatetypeDescriptionXOk

`func (o *EzdoctemplatedocumentResponseCompound) GetSEzdoctemplatetypeDescriptionXOk() (*string, bool)`

GetSEzdoctemplatetypeDescriptionXOk returns a tuple with the SEzdoctemplatetypeDescriptionX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzdoctemplatetypeDescriptionX

`func (o *EzdoctemplatedocumentResponseCompound) SetSEzdoctemplatetypeDescriptionX(v string)`

SetSEzdoctemplatetypeDescriptionX sets SEzdoctemplatetypeDescriptionX field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


