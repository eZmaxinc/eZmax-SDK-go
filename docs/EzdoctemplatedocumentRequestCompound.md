# EzdoctemplatedocumentRequestCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzdoctemplatedocumentID** | Pointer to **int32** | The unique ID of the Ezdoctemplatedocument | [optional] 
**FkiLanguageID** | **int32** | The unique ID of the Language.  Valid values:  |Value|Description| |-|-| |1|French| |2|English| | 
**FkiEzsignfoldertypeID** | Pointer to **int32** | The unique ID of the Ezsignfoldertype. | [optional] 
**FkiEzdoctemplatetypeID** | **int32** | The unique ID of the Ezdoctemplatetype | 
**FkiEzdoctemplatefieldtypecategoryID** | **int32** | The unique ID of the Ezdoctemplatefieldtypecategory | 
**EEzdoctemplatedocumentPrivacylevel** | Pointer to [**FieldEEzdoctemplatedocumentPrivacylevel**](FieldEEzdoctemplatedocumentPrivacylevel.md) |  | [optional] 
**BEzdoctemplatedocumentIsactive** | **bool** | Whether the ezdoctemplatedocument is active or not | 
**ObjEzdoctemplatedocumentName** | [**MultilingualEzdoctemplatedocumentName**](MultilingualEzdoctemplatedocumentName.md) |  | 

## Methods

### NewEzdoctemplatedocumentRequestCompound

`func NewEzdoctemplatedocumentRequestCompound(fkiLanguageID int32, fkiEzdoctemplatetypeID int32, fkiEzdoctemplatefieldtypecategoryID int32, bEzdoctemplatedocumentIsactive bool, objEzdoctemplatedocumentName MultilingualEzdoctemplatedocumentName, ) *EzdoctemplatedocumentRequestCompound`

NewEzdoctemplatedocumentRequestCompound instantiates a new EzdoctemplatedocumentRequestCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzdoctemplatedocumentRequestCompoundWithDefaults

`func NewEzdoctemplatedocumentRequestCompoundWithDefaults() *EzdoctemplatedocumentRequestCompound`

NewEzdoctemplatedocumentRequestCompoundWithDefaults instantiates a new EzdoctemplatedocumentRequestCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzdoctemplatedocumentID

`func (o *EzdoctemplatedocumentRequestCompound) GetPkiEzdoctemplatedocumentID() int32`

GetPkiEzdoctemplatedocumentID returns the PkiEzdoctemplatedocumentID field if non-nil, zero value otherwise.

### GetPkiEzdoctemplatedocumentIDOk

`func (o *EzdoctemplatedocumentRequestCompound) GetPkiEzdoctemplatedocumentIDOk() (*int32, bool)`

GetPkiEzdoctemplatedocumentIDOk returns a tuple with the PkiEzdoctemplatedocumentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzdoctemplatedocumentID

`func (o *EzdoctemplatedocumentRequestCompound) SetPkiEzdoctemplatedocumentID(v int32)`

SetPkiEzdoctemplatedocumentID sets PkiEzdoctemplatedocumentID field to given value.

### HasPkiEzdoctemplatedocumentID

`func (o *EzdoctemplatedocumentRequestCompound) HasPkiEzdoctemplatedocumentID() bool`

HasPkiEzdoctemplatedocumentID returns a boolean if a field has been set.

### GetFkiLanguageID

`func (o *EzdoctemplatedocumentRequestCompound) GetFkiLanguageID() int32`

GetFkiLanguageID returns the FkiLanguageID field if non-nil, zero value otherwise.

### GetFkiLanguageIDOk

`func (o *EzdoctemplatedocumentRequestCompound) GetFkiLanguageIDOk() (*int32, bool)`

GetFkiLanguageIDOk returns a tuple with the FkiLanguageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiLanguageID

`func (o *EzdoctemplatedocumentRequestCompound) SetFkiLanguageID(v int32)`

SetFkiLanguageID sets FkiLanguageID field to given value.


### GetFkiEzsignfoldertypeID

`func (o *EzdoctemplatedocumentRequestCompound) GetFkiEzsignfoldertypeID() int32`

GetFkiEzsignfoldertypeID returns the FkiEzsignfoldertypeID field if non-nil, zero value otherwise.

### GetFkiEzsignfoldertypeIDOk

`func (o *EzdoctemplatedocumentRequestCompound) GetFkiEzsignfoldertypeIDOk() (*int32, bool)`

GetFkiEzsignfoldertypeIDOk returns a tuple with the FkiEzsignfoldertypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignfoldertypeID

`func (o *EzdoctemplatedocumentRequestCompound) SetFkiEzsignfoldertypeID(v int32)`

SetFkiEzsignfoldertypeID sets FkiEzsignfoldertypeID field to given value.

### HasFkiEzsignfoldertypeID

`func (o *EzdoctemplatedocumentRequestCompound) HasFkiEzsignfoldertypeID() bool`

HasFkiEzsignfoldertypeID returns a boolean if a field has been set.

### GetFkiEzdoctemplatetypeID

`func (o *EzdoctemplatedocumentRequestCompound) GetFkiEzdoctemplatetypeID() int32`

GetFkiEzdoctemplatetypeID returns the FkiEzdoctemplatetypeID field if non-nil, zero value otherwise.

### GetFkiEzdoctemplatetypeIDOk

`func (o *EzdoctemplatedocumentRequestCompound) GetFkiEzdoctemplatetypeIDOk() (*int32, bool)`

GetFkiEzdoctemplatetypeIDOk returns a tuple with the FkiEzdoctemplatetypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzdoctemplatetypeID

`func (o *EzdoctemplatedocumentRequestCompound) SetFkiEzdoctemplatetypeID(v int32)`

SetFkiEzdoctemplatetypeID sets FkiEzdoctemplatetypeID field to given value.


### GetFkiEzdoctemplatefieldtypecategoryID

`func (o *EzdoctemplatedocumentRequestCompound) GetFkiEzdoctemplatefieldtypecategoryID() int32`

GetFkiEzdoctemplatefieldtypecategoryID returns the FkiEzdoctemplatefieldtypecategoryID field if non-nil, zero value otherwise.

### GetFkiEzdoctemplatefieldtypecategoryIDOk

`func (o *EzdoctemplatedocumentRequestCompound) GetFkiEzdoctemplatefieldtypecategoryIDOk() (*int32, bool)`

GetFkiEzdoctemplatefieldtypecategoryIDOk returns a tuple with the FkiEzdoctemplatefieldtypecategoryID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzdoctemplatefieldtypecategoryID

`func (o *EzdoctemplatedocumentRequestCompound) SetFkiEzdoctemplatefieldtypecategoryID(v int32)`

SetFkiEzdoctemplatefieldtypecategoryID sets FkiEzdoctemplatefieldtypecategoryID field to given value.


### GetEEzdoctemplatedocumentPrivacylevel

`func (o *EzdoctemplatedocumentRequestCompound) GetEEzdoctemplatedocumentPrivacylevel() FieldEEzdoctemplatedocumentPrivacylevel`

GetEEzdoctemplatedocumentPrivacylevel returns the EEzdoctemplatedocumentPrivacylevel field if non-nil, zero value otherwise.

### GetEEzdoctemplatedocumentPrivacylevelOk

`func (o *EzdoctemplatedocumentRequestCompound) GetEEzdoctemplatedocumentPrivacylevelOk() (*FieldEEzdoctemplatedocumentPrivacylevel, bool)`

GetEEzdoctemplatedocumentPrivacylevelOk returns a tuple with the EEzdoctemplatedocumentPrivacylevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzdoctemplatedocumentPrivacylevel

`func (o *EzdoctemplatedocumentRequestCompound) SetEEzdoctemplatedocumentPrivacylevel(v FieldEEzdoctemplatedocumentPrivacylevel)`

SetEEzdoctemplatedocumentPrivacylevel sets EEzdoctemplatedocumentPrivacylevel field to given value.

### HasEEzdoctemplatedocumentPrivacylevel

`func (o *EzdoctemplatedocumentRequestCompound) HasEEzdoctemplatedocumentPrivacylevel() bool`

HasEEzdoctemplatedocumentPrivacylevel returns a boolean if a field has been set.

### GetBEzdoctemplatedocumentIsactive

`func (o *EzdoctemplatedocumentRequestCompound) GetBEzdoctemplatedocumentIsactive() bool`

GetBEzdoctemplatedocumentIsactive returns the BEzdoctemplatedocumentIsactive field if non-nil, zero value otherwise.

### GetBEzdoctemplatedocumentIsactiveOk

`func (o *EzdoctemplatedocumentRequestCompound) GetBEzdoctemplatedocumentIsactiveOk() (*bool, bool)`

GetBEzdoctemplatedocumentIsactiveOk returns a tuple with the BEzdoctemplatedocumentIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzdoctemplatedocumentIsactive

`func (o *EzdoctemplatedocumentRequestCompound) SetBEzdoctemplatedocumentIsactive(v bool)`

SetBEzdoctemplatedocumentIsactive sets BEzdoctemplatedocumentIsactive field to given value.


### GetObjEzdoctemplatedocumentName

`func (o *EzdoctemplatedocumentRequestCompound) GetObjEzdoctemplatedocumentName() MultilingualEzdoctemplatedocumentName`

GetObjEzdoctemplatedocumentName returns the ObjEzdoctemplatedocumentName field if non-nil, zero value otherwise.

### GetObjEzdoctemplatedocumentNameOk

`func (o *EzdoctemplatedocumentRequestCompound) GetObjEzdoctemplatedocumentNameOk() (*MultilingualEzdoctemplatedocumentName, bool)`

GetObjEzdoctemplatedocumentNameOk returns a tuple with the ObjEzdoctemplatedocumentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjEzdoctemplatedocumentName

`func (o *EzdoctemplatedocumentRequestCompound) SetObjEzdoctemplatedocumentName(v MultilingualEzdoctemplatedocumentName)`

SetObjEzdoctemplatedocumentName sets ObjEzdoctemplatedocumentName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


