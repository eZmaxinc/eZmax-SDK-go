# EzsigntemplateResponseCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsigntemplateID** | **int32** | The unique ID of the Ezsigntemplate | 
**FkiEzsigntemplatedocumentID** | Pointer to **int32** | The unique ID of the Ezsigntemplatedocument | [optional] 
**FkiEzsignfoldertypeID** | Pointer to **int32** | The unique ID of the Ezsignfoldertype. | [optional] 
**FkiLanguageID** | **int32** | The unique ID of the Language.  Valid values:  |Value|Description| |-|-| |1|French| |2|English| | 
**FkiEzdoctemplatedocumentID** | Pointer to **int32** | The unique ID of the Ezdoctemplatedocument | [optional] 
**SLanguageNameX** | **string** | The Name of the Language in the language of the requester | 
**SEzsigntemplateDescription** | **string** | The description of the Ezsigntemplate | 
**SEzsigntemplateExternaldescription** | Pointer to **string** | The external description of the Ezsigntemplate | [optional] 
**TEzsigntemplateComment** | Pointer to **string** | The comment of the Ezsigntemplate | [optional] 
**SEzsigntemplateFilenamepattern** | Pointer to **string** | The filename pattern of the Ezsigntemplate | [optional] 
**BEzsigntemplateAdminonly** | **bool** | Whether the Ezsigntemplate can be accessed by admin users only (eUserType&#x3D;Normal) | 
**SEzsignfoldertypeNameX** | Pointer to **string** | The name of the Ezsignfoldertype in the language of the requester | [optional] 
**ObjAudit** | [**CommonAudit**](CommonAudit.md) |  | 
**BEzsigntemplateEditallowed** | **bool** | Whether the Ezsigntemplate if allowed to edit or not | 
**EEzsigntemplateType** | Pointer to [**FieldEEzsigntemplateType**](FieldEEzsigntemplateType.md) |  | [optional] 
**ObjEzsigntemplatedocument** | Pointer to [**EzsigntemplatedocumentResponse**](EzsigntemplatedocumentResponse.md) |  | [optional] 
**AObjEzsigntemplatesigner** | [**[]EzsigntemplatesignerResponseCompound**](EzsigntemplatesignerResponseCompound.md) |  | 

## Methods

### NewEzsigntemplateResponseCompound

`func NewEzsigntemplateResponseCompound(pkiEzsigntemplateID int32, fkiLanguageID int32, sLanguageNameX string, sEzsigntemplateDescription string, bEzsigntemplateAdminonly bool, objAudit CommonAudit, bEzsigntemplateEditallowed bool, aObjEzsigntemplatesigner []EzsigntemplatesignerResponseCompound, ) *EzsigntemplateResponseCompound`

NewEzsigntemplateResponseCompound instantiates a new EzsigntemplateResponseCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigntemplateResponseCompoundWithDefaults

`func NewEzsigntemplateResponseCompoundWithDefaults() *EzsigntemplateResponseCompound`

NewEzsigntemplateResponseCompoundWithDefaults instantiates a new EzsigntemplateResponseCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsigntemplateID

`func (o *EzsigntemplateResponseCompound) GetPkiEzsigntemplateID() int32`

GetPkiEzsigntemplateID returns the PkiEzsigntemplateID field if non-nil, zero value otherwise.

### GetPkiEzsigntemplateIDOk

`func (o *EzsigntemplateResponseCompound) GetPkiEzsigntemplateIDOk() (*int32, bool)`

GetPkiEzsigntemplateIDOk returns a tuple with the PkiEzsigntemplateID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsigntemplateID

`func (o *EzsigntemplateResponseCompound) SetPkiEzsigntemplateID(v int32)`

SetPkiEzsigntemplateID sets PkiEzsigntemplateID field to given value.


### GetFkiEzsigntemplatedocumentID

`func (o *EzsigntemplateResponseCompound) GetFkiEzsigntemplatedocumentID() int32`

GetFkiEzsigntemplatedocumentID returns the FkiEzsigntemplatedocumentID field if non-nil, zero value otherwise.

### GetFkiEzsigntemplatedocumentIDOk

`func (o *EzsigntemplateResponseCompound) GetFkiEzsigntemplatedocumentIDOk() (*int32, bool)`

GetFkiEzsigntemplatedocumentIDOk returns a tuple with the FkiEzsigntemplatedocumentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigntemplatedocumentID

`func (o *EzsigntemplateResponseCompound) SetFkiEzsigntemplatedocumentID(v int32)`

SetFkiEzsigntemplatedocumentID sets FkiEzsigntemplatedocumentID field to given value.

### HasFkiEzsigntemplatedocumentID

`func (o *EzsigntemplateResponseCompound) HasFkiEzsigntemplatedocumentID() bool`

HasFkiEzsigntemplatedocumentID returns a boolean if a field has been set.

### GetFkiEzsignfoldertypeID

`func (o *EzsigntemplateResponseCompound) GetFkiEzsignfoldertypeID() int32`

GetFkiEzsignfoldertypeID returns the FkiEzsignfoldertypeID field if non-nil, zero value otherwise.

### GetFkiEzsignfoldertypeIDOk

`func (o *EzsigntemplateResponseCompound) GetFkiEzsignfoldertypeIDOk() (*int32, bool)`

GetFkiEzsignfoldertypeIDOk returns a tuple with the FkiEzsignfoldertypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignfoldertypeID

`func (o *EzsigntemplateResponseCompound) SetFkiEzsignfoldertypeID(v int32)`

SetFkiEzsignfoldertypeID sets FkiEzsignfoldertypeID field to given value.

### HasFkiEzsignfoldertypeID

`func (o *EzsigntemplateResponseCompound) HasFkiEzsignfoldertypeID() bool`

HasFkiEzsignfoldertypeID returns a boolean if a field has been set.

### GetFkiLanguageID

`func (o *EzsigntemplateResponseCompound) GetFkiLanguageID() int32`

GetFkiLanguageID returns the FkiLanguageID field if non-nil, zero value otherwise.

### GetFkiLanguageIDOk

`func (o *EzsigntemplateResponseCompound) GetFkiLanguageIDOk() (*int32, bool)`

GetFkiLanguageIDOk returns a tuple with the FkiLanguageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiLanguageID

`func (o *EzsigntemplateResponseCompound) SetFkiLanguageID(v int32)`

SetFkiLanguageID sets FkiLanguageID field to given value.


### GetFkiEzdoctemplatedocumentID

`func (o *EzsigntemplateResponseCompound) GetFkiEzdoctemplatedocumentID() int32`

GetFkiEzdoctemplatedocumentID returns the FkiEzdoctemplatedocumentID field if non-nil, zero value otherwise.

### GetFkiEzdoctemplatedocumentIDOk

`func (o *EzsigntemplateResponseCompound) GetFkiEzdoctemplatedocumentIDOk() (*int32, bool)`

GetFkiEzdoctemplatedocumentIDOk returns a tuple with the FkiEzdoctemplatedocumentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzdoctemplatedocumentID

`func (o *EzsigntemplateResponseCompound) SetFkiEzdoctemplatedocumentID(v int32)`

SetFkiEzdoctemplatedocumentID sets FkiEzdoctemplatedocumentID field to given value.

### HasFkiEzdoctemplatedocumentID

`func (o *EzsigntemplateResponseCompound) HasFkiEzdoctemplatedocumentID() bool`

HasFkiEzdoctemplatedocumentID returns a boolean if a field has been set.

### GetSLanguageNameX

`func (o *EzsigntemplateResponseCompound) GetSLanguageNameX() string`

GetSLanguageNameX returns the SLanguageNameX field if non-nil, zero value otherwise.

### GetSLanguageNameXOk

`func (o *EzsigntemplateResponseCompound) GetSLanguageNameXOk() (*string, bool)`

GetSLanguageNameXOk returns a tuple with the SLanguageNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSLanguageNameX

`func (o *EzsigntemplateResponseCompound) SetSLanguageNameX(v string)`

SetSLanguageNameX sets SLanguageNameX field to given value.


### GetSEzsigntemplateDescription

`func (o *EzsigntemplateResponseCompound) GetSEzsigntemplateDescription() string`

GetSEzsigntemplateDescription returns the SEzsigntemplateDescription field if non-nil, zero value otherwise.

### GetSEzsigntemplateDescriptionOk

`func (o *EzsigntemplateResponseCompound) GetSEzsigntemplateDescriptionOk() (*string, bool)`

GetSEzsigntemplateDescriptionOk returns a tuple with the SEzsigntemplateDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplateDescription

`func (o *EzsigntemplateResponseCompound) SetSEzsigntemplateDescription(v string)`

SetSEzsigntemplateDescription sets SEzsigntemplateDescription field to given value.


### GetSEzsigntemplateExternaldescription

`func (o *EzsigntemplateResponseCompound) GetSEzsigntemplateExternaldescription() string`

GetSEzsigntemplateExternaldescription returns the SEzsigntemplateExternaldescription field if non-nil, zero value otherwise.

### GetSEzsigntemplateExternaldescriptionOk

`func (o *EzsigntemplateResponseCompound) GetSEzsigntemplateExternaldescriptionOk() (*string, bool)`

GetSEzsigntemplateExternaldescriptionOk returns a tuple with the SEzsigntemplateExternaldescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplateExternaldescription

`func (o *EzsigntemplateResponseCompound) SetSEzsigntemplateExternaldescription(v string)`

SetSEzsigntemplateExternaldescription sets SEzsigntemplateExternaldescription field to given value.

### HasSEzsigntemplateExternaldescription

`func (o *EzsigntemplateResponseCompound) HasSEzsigntemplateExternaldescription() bool`

HasSEzsigntemplateExternaldescription returns a boolean if a field has been set.

### GetTEzsigntemplateComment

`func (o *EzsigntemplateResponseCompound) GetTEzsigntemplateComment() string`

GetTEzsigntemplateComment returns the TEzsigntemplateComment field if non-nil, zero value otherwise.

### GetTEzsigntemplateCommentOk

`func (o *EzsigntemplateResponseCompound) GetTEzsigntemplateCommentOk() (*string, bool)`

GetTEzsigntemplateCommentOk returns a tuple with the TEzsigntemplateComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTEzsigntemplateComment

`func (o *EzsigntemplateResponseCompound) SetTEzsigntemplateComment(v string)`

SetTEzsigntemplateComment sets TEzsigntemplateComment field to given value.

### HasTEzsigntemplateComment

`func (o *EzsigntemplateResponseCompound) HasTEzsigntemplateComment() bool`

HasTEzsigntemplateComment returns a boolean if a field has been set.

### GetSEzsigntemplateFilenamepattern

`func (o *EzsigntemplateResponseCompound) GetSEzsigntemplateFilenamepattern() string`

GetSEzsigntemplateFilenamepattern returns the SEzsigntemplateFilenamepattern field if non-nil, zero value otherwise.

### GetSEzsigntemplateFilenamepatternOk

`func (o *EzsigntemplateResponseCompound) GetSEzsigntemplateFilenamepatternOk() (*string, bool)`

GetSEzsigntemplateFilenamepatternOk returns a tuple with the SEzsigntemplateFilenamepattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplateFilenamepattern

`func (o *EzsigntemplateResponseCompound) SetSEzsigntemplateFilenamepattern(v string)`

SetSEzsigntemplateFilenamepattern sets SEzsigntemplateFilenamepattern field to given value.

### HasSEzsigntemplateFilenamepattern

`func (o *EzsigntemplateResponseCompound) HasSEzsigntemplateFilenamepattern() bool`

HasSEzsigntemplateFilenamepattern returns a boolean if a field has been set.

### GetBEzsigntemplateAdminonly

`func (o *EzsigntemplateResponseCompound) GetBEzsigntemplateAdminonly() bool`

GetBEzsigntemplateAdminonly returns the BEzsigntemplateAdminonly field if non-nil, zero value otherwise.

### GetBEzsigntemplateAdminonlyOk

`func (o *EzsigntemplateResponseCompound) GetBEzsigntemplateAdminonlyOk() (*bool, bool)`

GetBEzsigntemplateAdminonlyOk returns a tuple with the BEzsigntemplateAdminonly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsigntemplateAdminonly

`func (o *EzsigntemplateResponseCompound) SetBEzsigntemplateAdminonly(v bool)`

SetBEzsigntemplateAdminonly sets BEzsigntemplateAdminonly field to given value.


### GetSEzsignfoldertypeNameX

`func (o *EzsigntemplateResponseCompound) GetSEzsignfoldertypeNameX() string`

GetSEzsignfoldertypeNameX returns the SEzsignfoldertypeNameX field if non-nil, zero value otherwise.

### GetSEzsignfoldertypeNameXOk

`func (o *EzsigntemplateResponseCompound) GetSEzsignfoldertypeNameXOk() (*string, bool)`

GetSEzsignfoldertypeNameXOk returns a tuple with the SEzsignfoldertypeNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignfoldertypeNameX

`func (o *EzsigntemplateResponseCompound) SetSEzsignfoldertypeNameX(v string)`

SetSEzsignfoldertypeNameX sets SEzsignfoldertypeNameX field to given value.

### HasSEzsignfoldertypeNameX

`func (o *EzsigntemplateResponseCompound) HasSEzsignfoldertypeNameX() bool`

HasSEzsignfoldertypeNameX returns a boolean if a field has been set.

### GetObjAudit

`func (o *EzsigntemplateResponseCompound) GetObjAudit() CommonAudit`

GetObjAudit returns the ObjAudit field if non-nil, zero value otherwise.

### GetObjAuditOk

`func (o *EzsigntemplateResponseCompound) GetObjAuditOk() (*CommonAudit, bool)`

GetObjAuditOk returns a tuple with the ObjAudit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjAudit

`func (o *EzsigntemplateResponseCompound) SetObjAudit(v CommonAudit)`

SetObjAudit sets ObjAudit field to given value.


### GetBEzsigntemplateEditallowed

`func (o *EzsigntemplateResponseCompound) GetBEzsigntemplateEditallowed() bool`

GetBEzsigntemplateEditallowed returns the BEzsigntemplateEditallowed field if non-nil, zero value otherwise.

### GetBEzsigntemplateEditallowedOk

`func (o *EzsigntemplateResponseCompound) GetBEzsigntemplateEditallowedOk() (*bool, bool)`

GetBEzsigntemplateEditallowedOk returns a tuple with the BEzsigntemplateEditallowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsigntemplateEditallowed

`func (o *EzsigntemplateResponseCompound) SetBEzsigntemplateEditallowed(v bool)`

SetBEzsigntemplateEditallowed sets BEzsigntemplateEditallowed field to given value.


### GetEEzsigntemplateType

`func (o *EzsigntemplateResponseCompound) GetEEzsigntemplateType() FieldEEzsigntemplateType`

GetEEzsigntemplateType returns the EEzsigntemplateType field if non-nil, zero value otherwise.

### GetEEzsigntemplateTypeOk

`func (o *EzsigntemplateResponseCompound) GetEEzsigntemplateTypeOk() (*FieldEEzsigntemplateType, bool)`

GetEEzsigntemplateTypeOk returns a tuple with the EEzsigntemplateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplateType

`func (o *EzsigntemplateResponseCompound) SetEEzsigntemplateType(v FieldEEzsigntemplateType)`

SetEEzsigntemplateType sets EEzsigntemplateType field to given value.

### HasEEzsigntemplateType

`func (o *EzsigntemplateResponseCompound) HasEEzsigntemplateType() bool`

HasEEzsigntemplateType returns a boolean if a field has been set.

### GetObjEzsigntemplatedocument

`func (o *EzsigntemplateResponseCompound) GetObjEzsigntemplatedocument() EzsigntemplatedocumentResponse`

GetObjEzsigntemplatedocument returns the ObjEzsigntemplatedocument field if non-nil, zero value otherwise.

### GetObjEzsigntemplatedocumentOk

`func (o *EzsigntemplateResponseCompound) GetObjEzsigntemplatedocumentOk() (*EzsigntemplatedocumentResponse, bool)`

GetObjEzsigntemplatedocumentOk returns a tuple with the ObjEzsigntemplatedocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjEzsigntemplatedocument

`func (o *EzsigntemplateResponseCompound) SetObjEzsigntemplatedocument(v EzsigntemplatedocumentResponse)`

SetObjEzsigntemplatedocument sets ObjEzsigntemplatedocument field to given value.

### HasObjEzsigntemplatedocument

`func (o *EzsigntemplateResponseCompound) HasObjEzsigntemplatedocument() bool`

HasObjEzsigntemplatedocument returns a boolean if a field has been set.

### GetAObjEzsigntemplatesigner

`func (o *EzsigntemplateResponseCompound) GetAObjEzsigntemplatesigner() []EzsigntemplatesignerResponseCompound`

GetAObjEzsigntemplatesigner returns the AObjEzsigntemplatesigner field if non-nil, zero value otherwise.

### GetAObjEzsigntemplatesignerOk

`func (o *EzsigntemplateResponseCompound) GetAObjEzsigntemplatesignerOk() (*[]EzsigntemplatesignerResponseCompound, bool)`

GetAObjEzsigntemplatesignerOk returns a tuple with the AObjEzsigntemplatesigner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEzsigntemplatesigner

`func (o *EzsigntemplateResponseCompound) SetAObjEzsigntemplatesigner(v []EzsigntemplatesignerResponseCompound)`

SetAObjEzsigntemplatesigner sets AObjEzsigntemplatesigner field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


