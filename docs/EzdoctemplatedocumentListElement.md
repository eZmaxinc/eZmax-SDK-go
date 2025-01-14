# EzdoctemplatedocumentListElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzdoctemplatedocumentID** | **int32** | The unique ID of the Ezdoctemplatedocument | 
**FkiLanguageID** | **int32** | The unique ID of the Language.  Valid values:  |Value|Description| |-|-| |1|French| |2|English| | 
**FkiEzsignfoldertypeID** | Pointer to **int32** | The unique ID of the Ezsignfoldertype. | [optional] 
**FkiEzdoctemplatetypeID** | **int32** | The unique ID of the Ezdoctemplatetype | 
**FkiEzdoctemplatefieldtypecategoryID** | **int32** | The unique ID of the Ezdoctemplatefieldtypecategory | 
**SEzsignfoldertypeNameX** | Pointer to **string** | The name of the Ezsignfoldertype in the language of the requester | [optional] 
**SEzdoctemplatetypeDescriptionX** | Pointer to **string** | The description of the Ezdoctemplatetype in the language of the requester | [optional] 
**SEzdoctemplatefieldtypecategoryDescriptionX** | Pointer to **string** | The description of the Ezdoctemplatefieldtypecategory in the language of the requester | [optional] 
**EEzdoctemplatedocumentPrivacylevel** | Pointer to [**FieldEEzdoctemplatedocumentPrivacylevel**](FieldEEzdoctemplatedocumentPrivacylevel.md) |  | [optional] 
**BEzdoctemplatedocumentIsactive** | **bool** | Whether the ezdoctemplatedocument is active or not | 
**SEzdoctemplatedocumentNameX** | **string** | The name of the Ezdoctemplatedocument in the language of the requester | 

## Methods

### NewEzdoctemplatedocumentListElement

`func NewEzdoctemplatedocumentListElement(pkiEzdoctemplatedocumentID int32, fkiLanguageID int32, fkiEzdoctemplatetypeID int32, fkiEzdoctemplatefieldtypecategoryID int32, bEzdoctemplatedocumentIsactive bool, sEzdoctemplatedocumentNameX string, ) *EzdoctemplatedocumentListElement`

NewEzdoctemplatedocumentListElement instantiates a new EzdoctemplatedocumentListElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzdoctemplatedocumentListElementWithDefaults

`func NewEzdoctemplatedocumentListElementWithDefaults() *EzdoctemplatedocumentListElement`

NewEzdoctemplatedocumentListElementWithDefaults instantiates a new EzdoctemplatedocumentListElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzdoctemplatedocumentID

`func (o *EzdoctemplatedocumentListElement) GetPkiEzdoctemplatedocumentID() int32`

GetPkiEzdoctemplatedocumentID returns the PkiEzdoctemplatedocumentID field if non-nil, zero value otherwise.

### GetPkiEzdoctemplatedocumentIDOk

`func (o *EzdoctemplatedocumentListElement) GetPkiEzdoctemplatedocumentIDOk() (*int32, bool)`

GetPkiEzdoctemplatedocumentIDOk returns a tuple with the PkiEzdoctemplatedocumentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzdoctemplatedocumentID

`func (o *EzdoctemplatedocumentListElement) SetPkiEzdoctemplatedocumentID(v int32)`

SetPkiEzdoctemplatedocumentID sets PkiEzdoctemplatedocumentID field to given value.


### GetFkiLanguageID

`func (o *EzdoctemplatedocumentListElement) GetFkiLanguageID() int32`

GetFkiLanguageID returns the FkiLanguageID field if non-nil, zero value otherwise.

### GetFkiLanguageIDOk

`func (o *EzdoctemplatedocumentListElement) GetFkiLanguageIDOk() (*int32, bool)`

GetFkiLanguageIDOk returns a tuple with the FkiLanguageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiLanguageID

`func (o *EzdoctemplatedocumentListElement) SetFkiLanguageID(v int32)`

SetFkiLanguageID sets FkiLanguageID field to given value.


### GetFkiEzsignfoldertypeID

`func (o *EzdoctemplatedocumentListElement) GetFkiEzsignfoldertypeID() int32`

GetFkiEzsignfoldertypeID returns the FkiEzsignfoldertypeID field if non-nil, zero value otherwise.

### GetFkiEzsignfoldertypeIDOk

`func (o *EzdoctemplatedocumentListElement) GetFkiEzsignfoldertypeIDOk() (*int32, bool)`

GetFkiEzsignfoldertypeIDOk returns a tuple with the FkiEzsignfoldertypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignfoldertypeID

`func (o *EzdoctemplatedocumentListElement) SetFkiEzsignfoldertypeID(v int32)`

SetFkiEzsignfoldertypeID sets FkiEzsignfoldertypeID field to given value.

### HasFkiEzsignfoldertypeID

`func (o *EzdoctemplatedocumentListElement) HasFkiEzsignfoldertypeID() bool`

HasFkiEzsignfoldertypeID returns a boolean if a field has been set.

### GetFkiEzdoctemplatetypeID

`func (o *EzdoctemplatedocumentListElement) GetFkiEzdoctemplatetypeID() int32`

GetFkiEzdoctemplatetypeID returns the FkiEzdoctemplatetypeID field if non-nil, zero value otherwise.

### GetFkiEzdoctemplatetypeIDOk

`func (o *EzdoctemplatedocumentListElement) GetFkiEzdoctemplatetypeIDOk() (*int32, bool)`

GetFkiEzdoctemplatetypeIDOk returns a tuple with the FkiEzdoctemplatetypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzdoctemplatetypeID

`func (o *EzdoctemplatedocumentListElement) SetFkiEzdoctemplatetypeID(v int32)`

SetFkiEzdoctemplatetypeID sets FkiEzdoctemplatetypeID field to given value.


### GetFkiEzdoctemplatefieldtypecategoryID

`func (o *EzdoctemplatedocumentListElement) GetFkiEzdoctemplatefieldtypecategoryID() int32`

GetFkiEzdoctemplatefieldtypecategoryID returns the FkiEzdoctemplatefieldtypecategoryID field if non-nil, zero value otherwise.

### GetFkiEzdoctemplatefieldtypecategoryIDOk

`func (o *EzdoctemplatedocumentListElement) GetFkiEzdoctemplatefieldtypecategoryIDOk() (*int32, bool)`

GetFkiEzdoctemplatefieldtypecategoryIDOk returns a tuple with the FkiEzdoctemplatefieldtypecategoryID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzdoctemplatefieldtypecategoryID

`func (o *EzdoctemplatedocumentListElement) SetFkiEzdoctemplatefieldtypecategoryID(v int32)`

SetFkiEzdoctemplatefieldtypecategoryID sets FkiEzdoctemplatefieldtypecategoryID field to given value.


### GetSEzsignfoldertypeNameX

`func (o *EzdoctemplatedocumentListElement) GetSEzsignfoldertypeNameX() string`

GetSEzsignfoldertypeNameX returns the SEzsignfoldertypeNameX field if non-nil, zero value otherwise.

### GetSEzsignfoldertypeNameXOk

`func (o *EzdoctemplatedocumentListElement) GetSEzsignfoldertypeNameXOk() (*string, bool)`

GetSEzsignfoldertypeNameXOk returns a tuple with the SEzsignfoldertypeNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignfoldertypeNameX

`func (o *EzdoctemplatedocumentListElement) SetSEzsignfoldertypeNameX(v string)`

SetSEzsignfoldertypeNameX sets SEzsignfoldertypeNameX field to given value.

### HasSEzsignfoldertypeNameX

`func (o *EzdoctemplatedocumentListElement) HasSEzsignfoldertypeNameX() bool`

HasSEzsignfoldertypeNameX returns a boolean if a field has been set.

### GetSEzdoctemplatetypeDescriptionX

`func (o *EzdoctemplatedocumentListElement) GetSEzdoctemplatetypeDescriptionX() string`

GetSEzdoctemplatetypeDescriptionX returns the SEzdoctemplatetypeDescriptionX field if non-nil, zero value otherwise.

### GetSEzdoctemplatetypeDescriptionXOk

`func (o *EzdoctemplatedocumentListElement) GetSEzdoctemplatetypeDescriptionXOk() (*string, bool)`

GetSEzdoctemplatetypeDescriptionXOk returns a tuple with the SEzdoctemplatetypeDescriptionX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzdoctemplatetypeDescriptionX

`func (o *EzdoctemplatedocumentListElement) SetSEzdoctemplatetypeDescriptionX(v string)`

SetSEzdoctemplatetypeDescriptionX sets SEzdoctemplatetypeDescriptionX field to given value.

### HasSEzdoctemplatetypeDescriptionX

`func (o *EzdoctemplatedocumentListElement) HasSEzdoctemplatetypeDescriptionX() bool`

HasSEzdoctemplatetypeDescriptionX returns a boolean if a field has been set.

### GetSEzdoctemplatefieldtypecategoryDescriptionX

`func (o *EzdoctemplatedocumentListElement) GetSEzdoctemplatefieldtypecategoryDescriptionX() string`

GetSEzdoctemplatefieldtypecategoryDescriptionX returns the SEzdoctemplatefieldtypecategoryDescriptionX field if non-nil, zero value otherwise.

### GetSEzdoctemplatefieldtypecategoryDescriptionXOk

`func (o *EzdoctemplatedocumentListElement) GetSEzdoctemplatefieldtypecategoryDescriptionXOk() (*string, bool)`

GetSEzdoctemplatefieldtypecategoryDescriptionXOk returns a tuple with the SEzdoctemplatefieldtypecategoryDescriptionX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzdoctemplatefieldtypecategoryDescriptionX

`func (o *EzdoctemplatedocumentListElement) SetSEzdoctemplatefieldtypecategoryDescriptionX(v string)`

SetSEzdoctemplatefieldtypecategoryDescriptionX sets SEzdoctemplatefieldtypecategoryDescriptionX field to given value.

### HasSEzdoctemplatefieldtypecategoryDescriptionX

`func (o *EzdoctemplatedocumentListElement) HasSEzdoctemplatefieldtypecategoryDescriptionX() bool`

HasSEzdoctemplatefieldtypecategoryDescriptionX returns a boolean if a field has been set.

### GetEEzdoctemplatedocumentPrivacylevel

`func (o *EzdoctemplatedocumentListElement) GetEEzdoctemplatedocumentPrivacylevel() FieldEEzdoctemplatedocumentPrivacylevel`

GetEEzdoctemplatedocumentPrivacylevel returns the EEzdoctemplatedocumentPrivacylevel field if non-nil, zero value otherwise.

### GetEEzdoctemplatedocumentPrivacylevelOk

`func (o *EzdoctemplatedocumentListElement) GetEEzdoctemplatedocumentPrivacylevelOk() (*FieldEEzdoctemplatedocumentPrivacylevel, bool)`

GetEEzdoctemplatedocumentPrivacylevelOk returns a tuple with the EEzdoctemplatedocumentPrivacylevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzdoctemplatedocumentPrivacylevel

`func (o *EzdoctemplatedocumentListElement) SetEEzdoctemplatedocumentPrivacylevel(v FieldEEzdoctemplatedocumentPrivacylevel)`

SetEEzdoctemplatedocumentPrivacylevel sets EEzdoctemplatedocumentPrivacylevel field to given value.

### HasEEzdoctemplatedocumentPrivacylevel

`func (o *EzdoctemplatedocumentListElement) HasEEzdoctemplatedocumentPrivacylevel() bool`

HasEEzdoctemplatedocumentPrivacylevel returns a boolean if a field has been set.

### GetBEzdoctemplatedocumentIsactive

`func (o *EzdoctemplatedocumentListElement) GetBEzdoctemplatedocumentIsactive() bool`

GetBEzdoctemplatedocumentIsactive returns the BEzdoctemplatedocumentIsactive field if non-nil, zero value otherwise.

### GetBEzdoctemplatedocumentIsactiveOk

`func (o *EzdoctemplatedocumentListElement) GetBEzdoctemplatedocumentIsactiveOk() (*bool, bool)`

GetBEzdoctemplatedocumentIsactiveOk returns a tuple with the BEzdoctemplatedocumentIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzdoctemplatedocumentIsactive

`func (o *EzdoctemplatedocumentListElement) SetBEzdoctemplatedocumentIsactive(v bool)`

SetBEzdoctemplatedocumentIsactive sets BEzdoctemplatedocumentIsactive field to given value.


### GetSEzdoctemplatedocumentNameX

`func (o *EzdoctemplatedocumentListElement) GetSEzdoctemplatedocumentNameX() string`

GetSEzdoctemplatedocumentNameX returns the SEzdoctemplatedocumentNameX field if non-nil, zero value otherwise.

### GetSEzdoctemplatedocumentNameXOk

`func (o *EzdoctemplatedocumentListElement) GetSEzdoctemplatedocumentNameXOk() (*string, bool)`

GetSEzdoctemplatedocumentNameXOk returns a tuple with the SEzdoctemplatedocumentNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzdoctemplatedocumentNameX

`func (o *EzdoctemplatedocumentListElement) SetSEzdoctemplatedocumentNameX(v string)`

SetSEzdoctemplatedocumentNameX sets SEzdoctemplatedocumentNameX field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


