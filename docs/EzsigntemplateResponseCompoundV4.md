# EzsigntemplateResponseCompoundV4

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsigntemplateID** | **int32** | The unique ID of the Ezsigntemplate | 
**FkiEzsigntemplatedocumentID** | Pointer to **int32** | The unique ID of the Ezsigntemplatedocument | [optional] 
**FkiEzsignfoldertypeID** | Pointer to **int32** | The unique ID of the Ezsignfoldertype. | [optional] 
**ObjEzsignfoldertype** | Pointer to [**CustomEzsignfoldertypeTemplateResponse**](CustomEzsignfoldertypeTemplateResponse.md) |  | [optional] 
**FkiLanguageID** | **int32** | The unique ID of the Language.  Valid values:  |Value|Description| |-|-| |1|French| |2|English| | 
**FkiEzdoctemplatedocumentID** | Pointer to **int32** | The unique ID of the Ezdoctemplatedocument | [optional] 
**SEzdoctemplatedocumentNameX** | Pointer to **string** | The name of the Ezdoctemplatedocument in the language of the requester | [optional] 
**SLanguageNameX** | **string** | The Name of the Language in the language of the requester | 
**SEzsigntemplateDescription** | **string** | The description of the Ezsigntemplate | 
**SEzsigntemplateExternaldescription** | Pointer to **string** | The external description of the Ezsigntemplate | [optional] 
**TEzsigntemplateComment** | Pointer to **string** | The comment of the Ezsigntemplate | [optional] 
**EEzsigntemplateRecognition** | Pointer to [**FieldEEzsigntemplateRecognition**](FieldEEzsigntemplateRecognition.md) |  | [optional] [default to NO]
**SEzsigntemplateFilenameregexp** | Pointer to **string** | The filename regexp of the Ezsigntemplate. | [optional] 
**BEzsigntemplateAdminonly** | **bool** | Whether the Ezsigntemplate can be accessed by admin users only (eUserType&#x3D;Normal) | 
**SEzsignfoldertypeNameX** | Pointer to **string** | The name of the Ezsignfoldertype in the language of the requester | [optional] 
**ObjAudit** | [**CommonAudit**](CommonAudit.md) |  | 
**BEzsigntemplateEditallowed** | **bool** | Whether the Ezsigntemplate if allowed to edit or not | 
**EEzsigntemplateType** | Pointer to [**FieldEEzsigntemplateType**](FieldEEzsigntemplateType.md) |  | [optional] 
**ObjEzsigntemplatedocument** | Pointer to [**EzsigntemplatedocumentResponse**](EzsigntemplatedocumentResponse.md) |  | [optional] 
**AObjEzsigntemplatesigner** | [**[]EzsigntemplatesignerResponseCompoundV3**](EzsigntemplatesignerResponseCompoundV3.md) |  | 
**AObjEzsigntemplateannotation** | Pointer to [**[]EzsigntemplateannotationResponseCompound**](EzsigntemplateannotationResponseCompound.md) |  | [optional] 

## Methods

### NewEzsigntemplateResponseCompoundV4

`func NewEzsigntemplateResponseCompoundV4(pkiEzsigntemplateID int32, fkiLanguageID int32, sLanguageNameX string, sEzsigntemplateDescription string, bEzsigntemplateAdminonly bool, objAudit CommonAudit, bEzsigntemplateEditallowed bool, aObjEzsigntemplatesigner []EzsigntemplatesignerResponseCompoundV3, ) *EzsigntemplateResponseCompoundV4`

NewEzsigntemplateResponseCompoundV4 instantiates a new EzsigntemplateResponseCompoundV4 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigntemplateResponseCompoundV4WithDefaults

`func NewEzsigntemplateResponseCompoundV4WithDefaults() *EzsigntemplateResponseCompoundV4`

NewEzsigntemplateResponseCompoundV4WithDefaults instantiates a new EzsigntemplateResponseCompoundV4 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsigntemplateID

`func (o *EzsigntemplateResponseCompoundV4) GetPkiEzsigntemplateID() int32`

GetPkiEzsigntemplateID returns the PkiEzsigntemplateID field if non-nil, zero value otherwise.

### GetPkiEzsigntemplateIDOk

`func (o *EzsigntemplateResponseCompoundV4) GetPkiEzsigntemplateIDOk() (*int32, bool)`

GetPkiEzsigntemplateIDOk returns a tuple with the PkiEzsigntemplateID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsigntemplateID

`func (o *EzsigntemplateResponseCompoundV4) SetPkiEzsigntemplateID(v int32)`

SetPkiEzsigntemplateID sets PkiEzsigntemplateID field to given value.


### GetFkiEzsigntemplatedocumentID

`func (o *EzsigntemplateResponseCompoundV4) GetFkiEzsigntemplatedocumentID() int32`

GetFkiEzsigntemplatedocumentID returns the FkiEzsigntemplatedocumentID field if non-nil, zero value otherwise.

### GetFkiEzsigntemplatedocumentIDOk

`func (o *EzsigntemplateResponseCompoundV4) GetFkiEzsigntemplatedocumentIDOk() (*int32, bool)`

GetFkiEzsigntemplatedocumentIDOk returns a tuple with the FkiEzsigntemplatedocumentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigntemplatedocumentID

`func (o *EzsigntemplateResponseCompoundV4) SetFkiEzsigntemplatedocumentID(v int32)`

SetFkiEzsigntemplatedocumentID sets FkiEzsigntemplatedocumentID field to given value.

### HasFkiEzsigntemplatedocumentID

`func (o *EzsigntemplateResponseCompoundV4) HasFkiEzsigntemplatedocumentID() bool`

HasFkiEzsigntemplatedocumentID returns a boolean if a field has been set.

### GetFkiEzsignfoldertypeID

`func (o *EzsigntemplateResponseCompoundV4) GetFkiEzsignfoldertypeID() int32`

GetFkiEzsignfoldertypeID returns the FkiEzsignfoldertypeID field if non-nil, zero value otherwise.

### GetFkiEzsignfoldertypeIDOk

`func (o *EzsigntemplateResponseCompoundV4) GetFkiEzsignfoldertypeIDOk() (*int32, bool)`

GetFkiEzsignfoldertypeIDOk returns a tuple with the FkiEzsignfoldertypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignfoldertypeID

`func (o *EzsigntemplateResponseCompoundV4) SetFkiEzsignfoldertypeID(v int32)`

SetFkiEzsignfoldertypeID sets FkiEzsignfoldertypeID field to given value.

### HasFkiEzsignfoldertypeID

`func (o *EzsigntemplateResponseCompoundV4) HasFkiEzsignfoldertypeID() bool`

HasFkiEzsignfoldertypeID returns a boolean if a field has been set.

### GetObjEzsignfoldertype

`func (o *EzsigntemplateResponseCompoundV4) GetObjEzsignfoldertype() CustomEzsignfoldertypeTemplateResponse`

GetObjEzsignfoldertype returns the ObjEzsignfoldertype field if non-nil, zero value otherwise.

### GetObjEzsignfoldertypeOk

`func (o *EzsigntemplateResponseCompoundV4) GetObjEzsignfoldertypeOk() (*CustomEzsignfoldertypeTemplateResponse, bool)`

GetObjEzsignfoldertypeOk returns a tuple with the ObjEzsignfoldertype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjEzsignfoldertype

`func (o *EzsigntemplateResponseCompoundV4) SetObjEzsignfoldertype(v CustomEzsignfoldertypeTemplateResponse)`

SetObjEzsignfoldertype sets ObjEzsignfoldertype field to given value.

### HasObjEzsignfoldertype

`func (o *EzsigntemplateResponseCompoundV4) HasObjEzsignfoldertype() bool`

HasObjEzsignfoldertype returns a boolean if a field has been set.

### GetFkiLanguageID

`func (o *EzsigntemplateResponseCompoundV4) GetFkiLanguageID() int32`

GetFkiLanguageID returns the FkiLanguageID field if non-nil, zero value otherwise.

### GetFkiLanguageIDOk

`func (o *EzsigntemplateResponseCompoundV4) GetFkiLanguageIDOk() (*int32, bool)`

GetFkiLanguageIDOk returns a tuple with the FkiLanguageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiLanguageID

`func (o *EzsigntemplateResponseCompoundV4) SetFkiLanguageID(v int32)`

SetFkiLanguageID sets FkiLanguageID field to given value.


### GetFkiEzdoctemplatedocumentID

`func (o *EzsigntemplateResponseCompoundV4) GetFkiEzdoctemplatedocumentID() int32`

GetFkiEzdoctemplatedocumentID returns the FkiEzdoctemplatedocumentID field if non-nil, zero value otherwise.

### GetFkiEzdoctemplatedocumentIDOk

`func (o *EzsigntemplateResponseCompoundV4) GetFkiEzdoctemplatedocumentIDOk() (*int32, bool)`

GetFkiEzdoctemplatedocumentIDOk returns a tuple with the FkiEzdoctemplatedocumentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzdoctemplatedocumentID

`func (o *EzsigntemplateResponseCompoundV4) SetFkiEzdoctemplatedocumentID(v int32)`

SetFkiEzdoctemplatedocumentID sets FkiEzdoctemplatedocumentID field to given value.

### HasFkiEzdoctemplatedocumentID

`func (o *EzsigntemplateResponseCompoundV4) HasFkiEzdoctemplatedocumentID() bool`

HasFkiEzdoctemplatedocumentID returns a boolean if a field has been set.

### GetSEzdoctemplatedocumentNameX

`func (o *EzsigntemplateResponseCompoundV4) GetSEzdoctemplatedocumentNameX() string`

GetSEzdoctemplatedocumentNameX returns the SEzdoctemplatedocumentNameX field if non-nil, zero value otherwise.

### GetSEzdoctemplatedocumentNameXOk

`func (o *EzsigntemplateResponseCompoundV4) GetSEzdoctemplatedocumentNameXOk() (*string, bool)`

GetSEzdoctemplatedocumentNameXOk returns a tuple with the SEzdoctemplatedocumentNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzdoctemplatedocumentNameX

`func (o *EzsigntemplateResponseCompoundV4) SetSEzdoctemplatedocumentNameX(v string)`

SetSEzdoctemplatedocumentNameX sets SEzdoctemplatedocumentNameX field to given value.

### HasSEzdoctemplatedocumentNameX

`func (o *EzsigntemplateResponseCompoundV4) HasSEzdoctemplatedocumentNameX() bool`

HasSEzdoctemplatedocumentNameX returns a boolean if a field has been set.

### GetSLanguageNameX

`func (o *EzsigntemplateResponseCompoundV4) GetSLanguageNameX() string`

GetSLanguageNameX returns the SLanguageNameX field if non-nil, zero value otherwise.

### GetSLanguageNameXOk

`func (o *EzsigntemplateResponseCompoundV4) GetSLanguageNameXOk() (*string, bool)`

GetSLanguageNameXOk returns a tuple with the SLanguageNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSLanguageNameX

`func (o *EzsigntemplateResponseCompoundV4) SetSLanguageNameX(v string)`

SetSLanguageNameX sets SLanguageNameX field to given value.


### GetSEzsigntemplateDescription

`func (o *EzsigntemplateResponseCompoundV4) GetSEzsigntemplateDescription() string`

GetSEzsigntemplateDescription returns the SEzsigntemplateDescription field if non-nil, zero value otherwise.

### GetSEzsigntemplateDescriptionOk

`func (o *EzsigntemplateResponseCompoundV4) GetSEzsigntemplateDescriptionOk() (*string, bool)`

GetSEzsigntemplateDescriptionOk returns a tuple with the SEzsigntemplateDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplateDescription

`func (o *EzsigntemplateResponseCompoundV4) SetSEzsigntemplateDescription(v string)`

SetSEzsigntemplateDescription sets SEzsigntemplateDescription field to given value.


### GetSEzsigntemplateExternaldescription

`func (o *EzsigntemplateResponseCompoundV4) GetSEzsigntemplateExternaldescription() string`

GetSEzsigntemplateExternaldescription returns the SEzsigntemplateExternaldescription field if non-nil, zero value otherwise.

### GetSEzsigntemplateExternaldescriptionOk

`func (o *EzsigntemplateResponseCompoundV4) GetSEzsigntemplateExternaldescriptionOk() (*string, bool)`

GetSEzsigntemplateExternaldescriptionOk returns a tuple with the SEzsigntemplateExternaldescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplateExternaldescription

`func (o *EzsigntemplateResponseCompoundV4) SetSEzsigntemplateExternaldescription(v string)`

SetSEzsigntemplateExternaldescription sets SEzsigntemplateExternaldescription field to given value.

### HasSEzsigntemplateExternaldescription

`func (o *EzsigntemplateResponseCompoundV4) HasSEzsigntemplateExternaldescription() bool`

HasSEzsigntemplateExternaldescription returns a boolean if a field has been set.

### GetTEzsigntemplateComment

`func (o *EzsigntemplateResponseCompoundV4) GetTEzsigntemplateComment() string`

GetTEzsigntemplateComment returns the TEzsigntemplateComment field if non-nil, zero value otherwise.

### GetTEzsigntemplateCommentOk

`func (o *EzsigntemplateResponseCompoundV4) GetTEzsigntemplateCommentOk() (*string, bool)`

GetTEzsigntemplateCommentOk returns a tuple with the TEzsigntemplateComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTEzsigntemplateComment

`func (o *EzsigntemplateResponseCompoundV4) SetTEzsigntemplateComment(v string)`

SetTEzsigntemplateComment sets TEzsigntemplateComment field to given value.

### HasTEzsigntemplateComment

`func (o *EzsigntemplateResponseCompoundV4) HasTEzsigntemplateComment() bool`

HasTEzsigntemplateComment returns a boolean if a field has been set.

### GetEEzsigntemplateRecognition

`func (o *EzsigntemplateResponseCompoundV4) GetEEzsigntemplateRecognition() FieldEEzsigntemplateRecognition`

GetEEzsigntemplateRecognition returns the EEzsigntemplateRecognition field if non-nil, zero value otherwise.

### GetEEzsigntemplateRecognitionOk

`func (o *EzsigntemplateResponseCompoundV4) GetEEzsigntemplateRecognitionOk() (*FieldEEzsigntemplateRecognition, bool)`

GetEEzsigntemplateRecognitionOk returns a tuple with the EEzsigntemplateRecognition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplateRecognition

`func (o *EzsigntemplateResponseCompoundV4) SetEEzsigntemplateRecognition(v FieldEEzsigntemplateRecognition)`

SetEEzsigntemplateRecognition sets EEzsigntemplateRecognition field to given value.

### HasEEzsigntemplateRecognition

`func (o *EzsigntemplateResponseCompoundV4) HasEEzsigntemplateRecognition() bool`

HasEEzsigntemplateRecognition returns a boolean if a field has been set.

### GetSEzsigntemplateFilenameregexp

`func (o *EzsigntemplateResponseCompoundV4) GetSEzsigntemplateFilenameregexp() string`

GetSEzsigntemplateFilenameregexp returns the SEzsigntemplateFilenameregexp field if non-nil, zero value otherwise.

### GetSEzsigntemplateFilenameregexpOk

`func (o *EzsigntemplateResponseCompoundV4) GetSEzsigntemplateFilenameregexpOk() (*string, bool)`

GetSEzsigntemplateFilenameregexpOk returns a tuple with the SEzsigntemplateFilenameregexp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplateFilenameregexp

`func (o *EzsigntemplateResponseCompoundV4) SetSEzsigntemplateFilenameregexp(v string)`

SetSEzsigntemplateFilenameregexp sets SEzsigntemplateFilenameregexp field to given value.

### HasSEzsigntemplateFilenameregexp

`func (o *EzsigntemplateResponseCompoundV4) HasSEzsigntemplateFilenameregexp() bool`

HasSEzsigntemplateFilenameregexp returns a boolean if a field has been set.

### GetBEzsigntemplateAdminonly

`func (o *EzsigntemplateResponseCompoundV4) GetBEzsigntemplateAdminonly() bool`

GetBEzsigntemplateAdminonly returns the BEzsigntemplateAdminonly field if non-nil, zero value otherwise.

### GetBEzsigntemplateAdminonlyOk

`func (o *EzsigntemplateResponseCompoundV4) GetBEzsigntemplateAdminonlyOk() (*bool, bool)`

GetBEzsigntemplateAdminonlyOk returns a tuple with the BEzsigntemplateAdminonly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsigntemplateAdminonly

`func (o *EzsigntemplateResponseCompoundV4) SetBEzsigntemplateAdminonly(v bool)`

SetBEzsigntemplateAdminonly sets BEzsigntemplateAdminonly field to given value.


### GetSEzsignfoldertypeNameX

`func (o *EzsigntemplateResponseCompoundV4) GetSEzsignfoldertypeNameX() string`

GetSEzsignfoldertypeNameX returns the SEzsignfoldertypeNameX field if non-nil, zero value otherwise.

### GetSEzsignfoldertypeNameXOk

`func (o *EzsigntemplateResponseCompoundV4) GetSEzsignfoldertypeNameXOk() (*string, bool)`

GetSEzsignfoldertypeNameXOk returns a tuple with the SEzsignfoldertypeNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignfoldertypeNameX

`func (o *EzsigntemplateResponseCompoundV4) SetSEzsignfoldertypeNameX(v string)`

SetSEzsignfoldertypeNameX sets SEzsignfoldertypeNameX field to given value.

### HasSEzsignfoldertypeNameX

`func (o *EzsigntemplateResponseCompoundV4) HasSEzsignfoldertypeNameX() bool`

HasSEzsignfoldertypeNameX returns a boolean if a field has been set.

### GetObjAudit

`func (o *EzsigntemplateResponseCompoundV4) GetObjAudit() CommonAudit`

GetObjAudit returns the ObjAudit field if non-nil, zero value otherwise.

### GetObjAuditOk

`func (o *EzsigntemplateResponseCompoundV4) GetObjAuditOk() (*CommonAudit, bool)`

GetObjAuditOk returns a tuple with the ObjAudit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjAudit

`func (o *EzsigntemplateResponseCompoundV4) SetObjAudit(v CommonAudit)`

SetObjAudit sets ObjAudit field to given value.


### GetBEzsigntemplateEditallowed

`func (o *EzsigntemplateResponseCompoundV4) GetBEzsigntemplateEditallowed() bool`

GetBEzsigntemplateEditallowed returns the BEzsigntemplateEditallowed field if non-nil, zero value otherwise.

### GetBEzsigntemplateEditallowedOk

`func (o *EzsigntemplateResponseCompoundV4) GetBEzsigntemplateEditallowedOk() (*bool, bool)`

GetBEzsigntemplateEditallowedOk returns a tuple with the BEzsigntemplateEditallowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsigntemplateEditallowed

`func (o *EzsigntemplateResponseCompoundV4) SetBEzsigntemplateEditallowed(v bool)`

SetBEzsigntemplateEditallowed sets BEzsigntemplateEditallowed field to given value.


### GetEEzsigntemplateType

`func (o *EzsigntemplateResponseCompoundV4) GetEEzsigntemplateType() FieldEEzsigntemplateType`

GetEEzsigntemplateType returns the EEzsigntemplateType field if non-nil, zero value otherwise.

### GetEEzsigntemplateTypeOk

`func (o *EzsigntemplateResponseCompoundV4) GetEEzsigntemplateTypeOk() (*FieldEEzsigntemplateType, bool)`

GetEEzsigntemplateTypeOk returns a tuple with the EEzsigntemplateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplateType

`func (o *EzsigntemplateResponseCompoundV4) SetEEzsigntemplateType(v FieldEEzsigntemplateType)`

SetEEzsigntemplateType sets EEzsigntemplateType field to given value.

### HasEEzsigntemplateType

`func (o *EzsigntemplateResponseCompoundV4) HasEEzsigntemplateType() bool`

HasEEzsigntemplateType returns a boolean if a field has been set.

### GetObjEzsigntemplatedocument

`func (o *EzsigntemplateResponseCompoundV4) GetObjEzsigntemplatedocument() EzsigntemplatedocumentResponse`

GetObjEzsigntemplatedocument returns the ObjEzsigntemplatedocument field if non-nil, zero value otherwise.

### GetObjEzsigntemplatedocumentOk

`func (o *EzsigntemplateResponseCompoundV4) GetObjEzsigntemplatedocumentOk() (*EzsigntemplatedocumentResponse, bool)`

GetObjEzsigntemplatedocumentOk returns a tuple with the ObjEzsigntemplatedocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjEzsigntemplatedocument

`func (o *EzsigntemplateResponseCompoundV4) SetObjEzsigntemplatedocument(v EzsigntemplatedocumentResponse)`

SetObjEzsigntemplatedocument sets ObjEzsigntemplatedocument field to given value.

### HasObjEzsigntemplatedocument

`func (o *EzsigntemplateResponseCompoundV4) HasObjEzsigntemplatedocument() bool`

HasObjEzsigntemplatedocument returns a boolean if a field has been set.

### GetAObjEzsigntemplatesigner

`func (o *EzsigntemplateResponseCompoundV4) GetAObjEzsigntemplatesigner() []EzsigntemplatesignerResponseCompoundV3`

GetAObjEzsigntemplatesigner returns the AObjEzsigntemplatesigner field if non-nil, zero value otherwise.

### GetAObjEzsigntemplatesignerOk

`func (o *EzsigntemplateResponseCompoundV4) GetAObjEzsigntemplatesignerOk() (*[]EzsigntemplatesignerResponseCompoundV3, bool)`

GetAObjEzsigntemplatesignerOk returns a tuple with the AObjEzsigntemplatesigner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEzsigntemplatesigner

`func (o *EzsigntemplateResponseCompoundV4) SetAObjEzsigntemplatesigner(v []EzsigntemplatesignerResponseCompoundV3)`

SetAObjEzsigntemplatesigner sets AObjEzsigntemplatesigner field to given value.


### GetAObjEzsigntemplateannotation

`func (o *EzsigntemplateResponseCompoundV4) GetAObjEzsigntemplateannotation() []EzsigntemplateannotationResponseCompound`

GetAObjEzsigntemplateannotation returns the AObjEzsigntemplateannotation field if non-nil, zero value otherwise.

### GetAObjEzsigntemplateannotationOk

`func (o *EzsigntemplateResponseCompoundV4) GetAObjEzsigntemplateannotationOk() (*[]EzsigntemplateannotationResponseCompound, bool)`

GetAObjEzsigntemplateannotationOk returns a tuple with the AObjEzsigntemplateannotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEzsigntemplateannotation

`func (o *EzsigntemplateResponseCompoundV4) SetAObjEzsigntemplateannotation(v []EzsigntemplateannotationResponseCompound)`

SetAObjEzsigntemplateannotation sets AObjEzsigntemplateannotation field to given value.

### HasAObjEzsigntemplateannotation

`func (o *EzsigntemplateResponseCompoundV4) HasAObjEzsigntemplateannotation() bool`

HasAObjEzsigntemplateannotation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


