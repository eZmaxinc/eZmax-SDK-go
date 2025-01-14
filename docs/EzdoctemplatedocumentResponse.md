# EzdoctemplatedocumentResponse

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

### NewEzdoctemplatedocumentResponse

`func NewEzdoctemplatedocumentResponse(pkiEzdoctemplatedocumentID int32, fkiLanguageID int32, fkiEzdoctemplatetypeID int32, fkiEzdoctemplatefieldtypecategoryID int32, bEzdoctemplatedocumentIsactive bool, objEzdoctemplatedocumentName MultilingualEzdoctemplatedocumentName, sEzdoctemplatefieldtypecategoryDescriptionX string, sEzdoctemplatetypeDescriptionX string, ) *EzdoctemplatedocumentResponse`

NewEzdoctemplatedocumentResponse instantiates a new EzdoctemplatedocumentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzdoctemplatedocumentResponseWithDefaults

`func NewEzdoctemplatedocumentResponseWithDefaults() *EzdoctemplatedocumentResponse`

NewEzdoctemplatedocumentResponseWithDefaults instantiates a new EzdoctemplatedocumentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzdoctemplatedocumentID

`func (o *EzdoctemplatedocumentResponse) GetPkiEzdoctemplatedocumentID() int32`

GetPkiEzdoctemplatedocumentID returns the PkiEzdoctemplatedocumentID field if non-nil, zero value otherwise.

### GetPkiEzdoctemplatedocumentIDOk

`func (o *EzdoctemplatedocumentResponse) GetPkiEzdoctemplatedocumentIDOk() (*int32, bool)`

GetPkiEzdoctemplatedocumentIDOk returns a tuple with the PkiEzdoctemplatedocumentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzdoctemplatedocumentID

`func (o *EzdoctemplatedocumentResponse) SetPkiEzdoctemplatedocumentID(v int32)`

SetPkiEzdoctemplatedocumentID sets PkiEzdoctemplatedocumentID field to given value.


### GetFkiLanguageID

`func (o *EzdoctemplatedocumentResponse) GetFkiLanguageID() int32`

GetFkiLanguageID returns the FkiLanguageID field if non-nil, zero value otherwise.

### GetFkiLanguageIDOk

`func (o *EzdoctemplatedocumentResponse) GetFkiLanguageIDOk() (*int32, bool)`

GetFkiLanguageIDOk returns a tuple with the FkiLanguageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiLanguageID

`func (o *EzdoctemplatedocumentResponse) SetFkiLanguageID(v int32)`

SetFkiLanguageID sets FkiLanguageID field to given value.


### GetFkiEzsignfoldertypeID

`func (o *EzdoctemplatedocumentResponse) GetFkiEzsignfoldertypeID() int32`

GetFkiEzsignfoldertypeID returns the FkiEzsignfoldertypeID field if non-nil, zero value otherwise.

### GetFkiEzsignfoldertypeIDOk

`func (o *EzdoctemplatedocumentResponse) GetFkiEzsignfoldertypeIDOk() (*int32, bool)`

GetFkiEzsignfoldertypeIDOk returns a tuple with the FkiEzsignfoldertypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignfoldertypeID

`func (o *EzdoctemplatedocumentResponse) SetFkiEzsignfoldertypeID(v int32)`

SetFkiEzsignfoldertypeID sets FkiEzsignfoldertypeID field to given value.

### HasFkiEzsignfoldertypeID

`func (o *EzdoctemplatedocumentResponse) HasFkiEzsignfoldertypeID() bool`

HasFkiEzsignfoldertypeID returns a boolean if a field has been set.

### GetFkiEzdoctemplatetypeID

`func (o *EzdoctemplatedocumentResponse) GetFkiEzdoctemplatetypeID() int32`

GetFkiEzdoctemplatetypeID returns the FkiEzdoctemplatetypeID field if non-nil, zero value otherwise.

### GetFkiEzdoctemplatetypeIDOk

`func (o *EzdoctemplatedocumentResponse) GetFkiEzdoctemplatetypeIDOk() (*int32, bool)`

GetFkiEzdoctemplatetypeIDOk returns a tuple with the FkiEzdoctemplatetypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzdoctemplatetypeID

`func (o *EzdoctemplatedocumentResponse) SetFkiEzdoctemplatetypeID(v int32)`

SetFkiEzdoctemplatetypeID sets FkiEzdoctemplatetypeID field to given value.


### GetFkiEzdoctemplatefieldtypecategoryID

`func (o *EzdoctemplatedocumentResponse) GetFkiEzdoctemplatefieldtypecategoryID() int32`

GetFkiEzdoctemplatefieldtypecategoryID returns the FkiEzdoctemplatefieldtypecategoryID field if non-nil, zero value otherwise.

### GetFkiEzdoctemplatefieldtypecategoryIDOk

`func (o *EzdoctemplatedocumentResponse) GetFkiEzdoctemplatefieldtypecategoryIDOk() (*int32, bool)`

GetFkiEzdoctemplatefieldtypecategoryIDOk returns a tuple with the FkiEzdoctemplatefieldtypecategoryID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzdoctemplatefieldtypecategoryID

`func (o *EzdoctemplatedocumentResponse) SetFkiEzdoctemplatefieldtypecategoryID(v int32)`

SetFkiEzdoctemplatefieldtypecategoryID sets FkiEzdoctemplatefieldtypecategoryID field to given value.


### GetEEzdoctemplatedocumentPrivacylevel

`func (o *EzdoctemplatedocumentResponse) GetEEzdoctemplatedocumentPrivacylevel() FieldEEzdoctemplatedocumentPrivacylevel`

GetEEzdoctemplatedocumentPrivacylevel returns the EEzdoctemplatedocumentPrivacylevel field if non-nil, zero value otherwise.

### GetEEzdoctemplatedocumentPrivacylevelOk

`func (o *EzdoctemplatedocumentResponse) GetEEzdoctemplatedocumentPrivacylevelOk() (*FieldEEzdoctemplatedocumentPrivacylevel, bool)`

GetEEzdoctemplatedocumentPrivacylevelOk returns a tuple with the EEzdoctemplatedocumentPrivacylevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzdoctemplatedocumentPrivacylevel

`func (o *EzdoctemplatedocumentResponse) SetEEzdoctemplatedocumentPrivacylevel(v FieldEEzdoctemplatedocumentPrivacylevel)`

SetEEzdoctemplatedocumentPrivacylevel sets EEzdoctemplatedocumentPrivacylevel field to given value.

### HasEEzdoctemplatedocumentPrivacylevel

`func (o *EzdoctemplatedocumentResponse) HasEEzdoctemplatedocumentPrivacylevel() bool`

HasEEzdoctemplatedocumentPrivacylevel returns a boolean if a field has been set.

### GetBEzdoctemplatedocumentIsactive

`func (o *EzdoctemplatedocumentResponse) GetBEzdoctemplatedocumentIsactive() bool`

GetBEzdoctemplatedocumentIsactive returns the BEzdoctemplatedocumentIsactive field if non-nil, zero value otherwise.

### GetBEzdoctemplatedocumentIsactiveOk

`func (o *EzdoctemplatedocumentResponse) GetBEzdoctemplatedocumentIsactiveOk() (*bool, bool)`

GetBEzdoctemplatedocumentIsactiveOk returns a tuple with the BEzdoctemplatedocumentIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzdoctemplatedocumentIsactive

`func (o *EzdoctemplatedocumentResponse) SetBEzdoctemplatedocumentIsactive(v bool)`

SetBEzdoctemplatedocumentIsactive sets BEzdoctemplatedocumentIsactive field to given value.


### GetObjEzdoctemplatedocumentName

`func (o *EzdoctemplatedocumentResponse) GetObjEzdoctemplatedocumentName() MultilingualEzdoctemplatedocumentName`

GetObjEzdoctemplatedocumentName returns the ObjEzdoctemplatedocumentName field if non-nil, zero value otherwise.

### GetObjEzdoctemplatedocumentNameOk

`func (o *EzdoctemplatedocumentResponse) GetObjEzdoctemplatedocumentNameOk() (*MultilingualEzdoctemplatedocumentName, bool)`

GetObjEzdoctemplatedocumentNameOk returns a tuple with the ObjEzdoctemplatedocumentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjEzdoctemplatedocumentName

`func (o *EzdoctemplatedocumentResponse) SetObjEzdoctemplatedocumentName(v MultilingualEzdoctemplatedocumentName)`

SetObjEzdoctemplatedocumentName sets ObjEzdoctemplatedocumentName field to given value.


### GetSEzdoctemplatedocumentNameX

`func (o *EzdoctemplatedocumentResponse) GetSEzdoctemplatedocumentNameX() string`

GetSEzdoctemplatedocumentNameX returns the SEzdoctemplatedocumentNameX field if non-nil, zero value otherwise.

### GetSEzdoctemplatedocumentNameXOk

`func (o *EzdoctemplatedocumentResponse) GetSEzdoctemplatedocumentNameXOk() (*string, bool)`

GetSEzdoctemplatedocumentNameXOk returns a tuple with the SEzdoctemplatedocumentNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzdoctemplatedocumentNameX

`func (o *EzdoctemplatedocumentResponse) SetSEzdoctemplatedocumentNameX(v string)`

SetSEzdoctemplatedocumentNameX sets SEzdoctemplatedocumentNameX field to given value.

### HasSEzdoctemplatedocumentNameX

`func (o *EzdoctemplatedocumentResponse) HasSEzdoctemplatedocumentNameX() bool`

HasSEzdoctemplatedocumentNameX returns a boolean if a field has been set.

### GetSEzsignfoldertypeNameX

`func (o *EzdoctemplatedocumentResponse) GetSEzsignfoldertypeNameX() string`

GetSEzsignfoldertypeNameX returns the SEzsignfoldertypeNameX field if non-nil, zero value otherwise.

### GetSEzsignfoldertypeNameXOk

`func (o *EzdoctemplatedocumentResponse) GetSEzsignfoldertypeNameXOk() (*string, bool)`

GetSEzsignfoldertypeNameXOk returns a tuple with the SEzsignfoldertypeNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignfoldertypeNameX

`func (o *EzdoctemplatedocumentResponse) SetSEzsignfoldertypeNameX(v string)`

SetSEzsignfoldertypeNameX sets SEzsignfoldertypeNameX field to given value.

### HasSEzsignfoldertypeNameX

`func (o *EzdoctemplatedocumentResponse) HasSEzsignfoldertypeNameX() bool`

HasSEzsignfoldertypeNameX returns a boolean if a field has been set.

### GetSEzdoctemplatefieldtypecategoryDescriptionX

`func (o *EzdoctemplatedocumentResponse) GetSEzdoctemplatefieldtypecategoryDescriptionX() string`

GetSEzdoctemplatefieldtypecategoryDescriptionX returns the SEzdoctemplatefieldtypecategoryDescriptionX field if non-nil, zero value otherwise.

### GetSEzdoctemplatefieldtypecategoryDescriptionXOk

`func (o *EzdoctemplatedocumentResponse) GetSEzdoctemplatefieldtypecategoryDescriptionXOk() (*string, bool)`

GetSEzdoctemplatefieldtypecategoryDescriptionXOk returns a tuple with the SEzdoctemplatefieldtypecategoryDescriptionX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzdoctemplatefieldtypecategoryDescriptionX

`func (o *EzdoctemplatedocumentResponse) SetSEzdoctemplatefieldtypecategoryDescriptionX(v string)`

SetSEzdoctemplatefieldtypecategoryDescriptionX sets SEzdoctemplatefieldtypecategoryDescriptionX field to given value.


### GetSEzdoctemplatetypeDescriptionX

`func (o *EzdoctemplatedocumentResponse) GetSEzdoctemplatetypeDescriptionX() string`

GetSEzdoctemplatetypeDescriptionX returns the SEzdoctemplatetypeDescriptionX field if non-nil, zero value otherwise.

### GetSEzdoctemplatetypeDescriptionXOk

`func (o *EzdoctemplatedocumentResponse) GetSEzdoctemplatetypeDescriptionXOk() (*string, bool)`

GetSEzdoctemplatetypeDescriptionXOk returns a tuple with the SEzdoctemplatetypeDescriptionX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzdoctemplatetypeDescriptionX

`func (o *EzdoctemplatedocumentResponse) SetSEzdoctemplatetypeDescriptionX(v string)`

SetSEzdoctemplatetypeDescriptionX sets SEzdoctemplatetypeDescriptionX field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


