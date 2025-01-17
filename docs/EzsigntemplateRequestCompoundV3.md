# EzsigntemplateRequestCompoundV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsigntemplateID** | Pointer to **int32** | The unique ID of the Ezsigntemplate | [optional] 
**FkiEzsignfoldertypeID** | Pointer to **int32** | The unique ID of the Ezsignfoldertype. | [optional] 
**FkiLanguageID** | **int32** | The unique ID of the Language.  Valid values:  |Value|Description| |-|-| |1|French| |2|English| | 
**FkiEzdoctemplatedocumentID** | Pointer to **int32** | The unique ID of the Ezdoctemplatedocument | [optional] 
**SEzsigntemplateDescription** | **string** | The description of the Ezsigntemplate | 
**SEzsigntemplateExternaldescription** | Pointer to **string** | The external description of the Ezsigntemplate | [optional] 
**TEzsigntemplateComment** | Pointer to **string** | The comment of the Ezsigntemplate | [optional] 
**EEzsigntemplateRecognition** | Pointer to [**FieldEEzsigntemplateRecognition**](FieldEEzsigntemplateRecognition.md) |  | [optional] [default to NO]
**SEzsigntemplateFilenameregexp** | Pointer to **string** | The filename regexp of the Ezsigntemplate. | [optional] 
**BEzsigntemplateAdminonly** | **bool** | Whether the Ezsigntemplate can be accessed by admin users only (eUserType&#x3D;Normal) | 
**EEzsigntemplateType** | [**FieldEEzsigntemplateType**](FieldEEzsigntemplateType.md) |  | 

## Methods

### NewEzsigntemplateRequestCompoundV3

`func NewEzsigntemplateRequestCompoundV3(fkiLanguageID int32, sEzsigntemplateDescription string, bEzsigntemplateAdminonly bool, eEzsigntemplateType FieldEEzsigntemplateType, ) *EzsigntemplateRequestCompoundV3`

NewEzsigntemplateRequestCompoundV3 instantiates a new EzsigntemplateRequestCompoundV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigntemplateRequestCompoundV3WithDefaults

`func NewEzsigntemplateRequestCompoundV3WithDefaults() *EzsigntemplateRequestCompoundV3`

NewEzsigntemplateRequestCompoundV3WithDefaults instantiates a new EzsigntemplateRequestCompoundV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsigntemplateID

`func (o *EzsigntemplateRequestCompoundV3) GetPkiEzsigntemplateID() int32`

GetPkiEzsigntemplateID returns the PkiEzsigntemplateID field if non-nil, zero value otherwise.

### GetPkiEzsigntemplateIDOk

`func (o *EzsigntemplateRequestCompoundV3) GetPkiEzsigntemplateIDOk() (*int32, bool)`

GetPkiEzsigntemplateIDOk returns a tuple with the PkiEzsigntemplateID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsigntemplateID

`func (o *EzsigntemplateRequestCompoundV3) SetPkiEzsigntemplateID(v int32)`

SetPkiEzsigntemplateID sets PkiEzsigntemplateID field to given value.

### HasPkiEzsigntemplateID

`func (o *EzsigntemplateRequestCompoundV3) HasPkiEzsigntemplateID() bool`

HasPkiEzsigntemplateID returns a boolean if a field has been set.

### GetFkiEzsignfoldertypeID

`func (o *EzsigntemplateRequestCompoundV3) GetFkiEzsignfoldertypeID() int32`

GetFkiEzsignfoldertypeID returns the FkiEzsignfoldertypeID field if non-nil, zero value otherwise.

### GetFkiEzsignfoldertypeIDOk

`func (o *EzsigntemplateRequestCompoundV3) GetFkiEzsignfoldertypeIDOk() (*int32, bool)`

GetFkiEzsignfoldertypeIDOk returns a tuple with the FkiEzsignfoldertypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignfoldertypeID

`func (o *EzsigntemplateRequestCompoundV3) SetFkiEzsignfoldertypeID(v int32)`

SetFkiEzsignfoldertypeID sets FkiEzsignfoldertypeID field to given value.

### HasFkiEzsignfoldertypeID

`func (o *EzsigntemplateRequestCompoundV3) HasFkiEzsignfoldertypeID() bool`

HasFkiEzsignfoldertypeID returns a boolean if a field has been set.

### GetFkiLanguageID

`func (o *EzsigntemplateRequestCompoundV3) GetFkiLanguageID() int32`

GetFkiLanguageID returns the FkiLanguageID field if non-nil, zero value otherwise.

### GetFkiLanguageIDOk

`func (o *EzsigntemplateRequestCompoundV3) GetFkiLanguageIDOk() (*int32, bool)`

GetFkiLanguageIDOk returns a tuple with the FkiLanguageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiLanguageID

`func (o *EzsigntemplateRequestCompoundV3) SetFkiLanguageID(v int32)`

SetFkiLanguageID sets FkiLanguageID field to given value.


### GetFkiEzdoctemplatedocumentID

`func (o *EzsigntemplateRequestCompoundV3) GetFkiEzdoctemplatedocumentID() int32`

GetFkiEzdoctemplatedocumentID returns the FkiEzdoctemplatedocumentID field if non-nil, zero value otherwise.

### GetFkiEzdoctemplatedocumentIDOk

`func (o *EzsigntemplateRequestCompoundV3) GetFkiEzdoctemplatedocumentIDOk() (*int32, bool)`

GetFkiEzdoctemplatedocumentIDOk returns a tuple with the FkiEzdoctemplatedocumentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzdoctemplatedocumentID

`func (o *EzsigntemplateRequestCompoundV3) SetFkiEzdoctemplatedocumentID(v int32)`

SetFkiEzdoctemplatedocumentID sets FkiEzdoctemplatedocumentID field to given value.

### HasFkiEzdoctemplatedocumentID

`func (o *EzsigntemplateRequestCompoundV3) HasFkiEzdoctemplatedocumentID() bool`

HasFkiEzdoctemplatedocumentID returns a boolean if a field has been set.

### GetSEzsigntemplateDescription

`func (o *EzsigntemplateRequestCompoundV3) GetSEzsigntemplateDescription() string`

GetSEzsigntemplateDescription returns the SEzsigntemplateDescription field if non-nil, zero value otherwise.

### GetSEzsigntemplateDescriptionOk

`func (o *EzsigntemplateRequestCompoundV3) GetSEzsigntemplateDescriptionOk() (*string, bool)`

GetSEzsigntemplateDescriptionOk returns a tuple with the SEzsigntemplateDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplateDescription

`func (o *EzsigntemplateRequestCompoundV3) SetSEzsigntemplateDescription(v string)`

SetSEzsigntemplateDescription sets SEzsigntemplateDescription field to given value.


### GetSEzsigntemplateExternaldescription

`func (o *EzsigntemplateRequestCompoundV3) GetSEzsigntemplateExternaldescription() string`

GetSEzsigntemplateExternaldescription returns the SEzsigntemplateExternaldescription field if non-nil, zero value otherwise.

### GetSEzsigntemplateExternaldescriptionOk

`func (o *EzsigntemplateRequestCompoundV3) GetSEzsigntemplateExternaldescriptionOk() (*string, bool)`

GetSEzsigntemplateExternaldescriptionOk returns a tuple with the SEzsigntemplateExternaldescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplateExternaldescription

`func (o *EzsigntemplateRequestCompoundV3) SetSEzsigntemplateExternaldescription(v string)`

SetSEzsigntemplateExternaldescription sets SEzsigntemplateExternaldescription field to given value.

### HasSEzsigntemplateExternaldescription

`func (o *EzsigntemplateRequestCompoundV3) HasSEzsigntemplateExternaldescription() bool`

HasSEzsigntemplateExternaldescription returns a boolean if a field has been set.

### GetTEzsigntemplateComment

`func (o *EzsigntemplateRequestCompoundV3) GetTEzsigntemplateComment() string`

GetTEzsigntemplateComment returns the TEzsigntemplateComment field if non-nil, zero value otherwise.

### GetTEzsigntemplateCommentOk

`func (o *EzsigntemplateRequestCompoundV3) GetTEzsigntemplateCommentOk() (*string, bool)`

GetTEzsigntemplateCommentOk returns a tuple with the TEzsigntemplateComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTEzsigntemplateComment

`func (o *EzsigntemplateRequestCompoundV3) SetTEzsigntemplateComment(v string)`

SetTEzsigntemplateComment sets TEzsigntemplateComment field to given value.

### HasTEzsigntemplateComment

`func (o *EzsigntemplateRequestCompoundV3) HasTEzsigntemplateComment() bool`

HasTEzsigntemplateComment returns a boolean if a field has been set.

### GetEEzsigntemplateRecognition

`func (o *EzsigntemplateRequestCompoundV3) GetEEzsigntemplateRecognition() FieldEEzsigntemplateRecognition`

GetEEzsigntemplateRecognition returns the EEzsigntemplateRecognition field if non-nil, zero value otherwise.

### GetEEzsigntemplateRecognitionOk

`func (o *EzsigntemplateRequestCompoundV3) GetEEzsigntemplateRecognitionOk() (*FieldEEzsigntemplateRecognition, bool)`

GetEEzsigntemplateRecognitionOk returns a tuple with the EEzsigntemplateRecognition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplateRecognition

`func (o *EzsigntemplateRequestCompoundV3) SetEEzsigntemplateRecognition(v FieldEEzsigntemplateRecognition)`

SetEEzsigntemplateRecognition sets EEzsigntemplateRecognition field to given value.

### HasEEzsigntemplateRecognition

`func (o *EzsigntemplateRequestCompoundV3) HasEEzsigntemplateRecognition() bool`

HasEEzsigntemplateRecognition returns a boolean if a field has been set.

### GetSEzsigntemplateFilenameregexp

`func (o *EzsigntemplateRequestCompoundV3) GetSEzsigntemplateFilenameregexp() string`

GetSEzsigntemplateFilenameregexp returns the SEzsigntemplateFilenameregexp field if non-nil, zero value otherwise.

### GetSEzsigntemplateFilenameregexpOk

`func (o *EzsigntemplateRequestCompoundV3) GetSEzsigntemplateFilenameregexpOk() (*string, bool)`

GetSEzsigntemplateFilenameregexpOk returns a tuple with the SEzsigntemplateFilenameregexp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplateFilenameregexp

`func (o *EzsigntemplateRequestCompoundV3) SetSEzsigntemplateFilenameregexp(v string)`

SetSEzsigntemplateFilenameregexp sets SEzsigntemplateFilenameregexp field to given value.

### HasSEzsigntemplateFilenameregexp

`func (o *EzsigntemplateRequestCompoundV3) HasSEzsigntemplateFilenameregexp() bool`

HasSEzsigntemplateFilenameregexp returns a boolean if a field has been set.

### GetBEzsigntemplateAdminonly

`func (o *EzsigntemplateRequestCompoundV3) GetBEzsigntemplateAdminonly() bool`

GetBEzsigntemplateAdminonly returns the BEzsigntemplateAdminonly field if non-nil, zero value otherwise.

### GetBEzsigntemplateAdminonlyOk

`func (o *EzsigntemplateRequestCompoundV3) GetBEzsigntemplateAdminonlyOk() (*bool, bool)`

GetBEzsigntemplateAdminonlyOk returns a tuple with the BEzsigntemplateAdminonly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsigntemplateAdminonly

`func (o *EzsigntemplateRequestCompoundV3) SetBEzsigntemplateAdminonly(v bool)`

SetBEzsigntemplateAdminonly sets BEzsigntemplateAdminonly field to given value.


### GetEEzsigntemplateType

`func (o *EzsigntemplateRequestCompoundV3) GetEEzsigntemplateType() FieldEEzsigntemplateType`

GetEEzsigntemplateType returns the EEzsigntemplateType field if non-nil, zero value otherwise.

### GetEEzsigntemplateTypeOk

`func (o *EzsigntemplateRequestCompoundV3) GetEEzsigntemplateTypeOk() (*FieldEEzsigntemplateType, bool)`

GetEEzsigntemplateTypeOk returns a tuple with the EEzsigntemplateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplateType

`func (o *EzsigntemplateRequestCompoundV3) SetEEzsigntemplateType(v FieldEEzsigntemplateType)`

SetEEzsigntemplateType sets EEzsigntemplateType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


