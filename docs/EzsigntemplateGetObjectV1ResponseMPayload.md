# EzsigntemplateGetObjectV1ResponseMPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsigntemplateID** | **int32** | The unique ID of the Ezsigntemplate | 
**FkiEzsigntemplatedocumentID** | Pointer to **int32** | The unique ID of the Ezsigntemplatedocument | [optional] 
**FkiEzsignfoldertypeID** | **int32** | The unique ID of the Ezsignfoldertype. | 
**FkiLanguageID** | **int32** | The unique ID of the Language.  Valid values:  |Value|Description| |-|-| |1|French| |2|English| | 
**SLanguageNameX** | **string** | The Name of the Language in the language of the requester | 
**SEzsigntemplateDescription** | **string** | The description of the Ezsigntemplate | 
**BEzsigntemplateAdminonly** | **bool** | Whether the Ezsigntemplate can be accessed by admin users only (eUserType&#x3D;Normal) | 
**SEzsignfoldertypeNameX** | **string** | The name of the Ezsignfoldertype in the language of the requester | 
**ObjAudit** | [**CommonAudit**](CommonAudit.md) |  | 
**ObjEzsigntemplatedocument** | Pointer to [**EzsigntemplatedocumentResponse**](EzsigntemplatedocumentResponse.md) |  | [optional] 
**AObjEzsigntemplatesigner** | [**[]EzsigntemplatesignerResponseCompound**](EzsigntemplatesignerResponseCompound.md) |  | 

## Methods

### NewEzsigntemplateGetObjectV1ResponseMPayload

`func NewEzsigntemplateGetObjectV1ResponseMPayload(pkiEzsigntemplateID int32, fkiEzsignfoldertypeID int32, fkiLanguageID int32, sLanguageNameX string, sEzsigntemplateDescription string, bEzsigntemplateAdminonly bool, sEzsignfoldertypeNameX string, objAudit CommonAudit, aObjEzsigntemplatesigner []EzsigntemplatesignerResponseCompound, ) *EzsigntemplateGetObjectV1ResponseMPayload`

NewEzsigntemplateGetObjectV1ResponseMPayload instantiates a new EzsigntemplateGetObjectV1ResponseMPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigntemplateGetObjectV1ResponseMPayloadWithDefaults

`func NewEzsigntemplateGetObjectV1ResponseMPayloadWithDefaults() *EzsigntemplateGetObjectV1ResponseMPayload`

NewEzsigntemplateGetObjectV1ResponseMPayloadWithDefaults instantiates a new EzsigntemplateGetObjectV1ResponseMPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsigntemplateID

`func (o *EzsigntemplateGetObjectV1ResponseMPayload) GetPkiEzsigntemplateID() int32`

GetPkiEzsigntemplateID returns the PkiEzsigntemplateID field if non-nil, zero value otherwise.

### GetPkiEzsigntemplateIDOk

`func (o *EzsigntemplateGetObjectV1ResponseMPayload) GetPkiEzsigntemplateIDOk() (*int32, bool)`

GetPkiEzsigntemplateIDOk returns a tuple with the PkiEzsigntemplateID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsigntemplateID

`func (o *EzsigntemplateGetObjectV1ResponseMPayload) SetPkiEzsigntemplateID(v int32)`

SetPkiEzsigntemplateID sets PkiEzsigntemplateID field to given value.


### GetFkiEzsigntemplatedocumentID

`func (o *EzsigntemplateGetObjectV1ResponseMPayload) GetFkiEzsigntemplatedocumentID() int32`

GetFkiEzsigntemplatedocumentID returns the FkiEzsigntemplatedocumentID field if non-nil, zero value otherwise.

### GetFkiEzsigntemplatedocumentIDOk

`func (o *EzsigntemplateGetObjectV1ResponseMPayload) GetFkiEzsigntemplatedocumentIDOk() (*int32, bool)`

GetFkiEzsigntemplatedocumentIDOk returns a tuple with the FkiEzsigntemplatedocumentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigntemplatedocumentID

`func (o *EzsigntemplateGetObjectV1ResponseMPayload) SetFkiEzsigntemplatedocumentID(v int32)`

SetFkiEzsigntemplatedocumentID sets FkiEzsigntemplatedocumentID field to given value.

### HasFkiEzsigntemplatedocumentID

`func (o *EzsigntemplateGetObjectV1ResponseMPayload) HasFkiEzsigntemplatedocumentID() bool`

HasFkiEzsigntemplatedocumentID returns a boolean if a field has been set.

### GetFkiEzsignfoldertypeID

`func (o *EzsigntemplateGetObjectV1ResponseMPayload) GetFkiEzsignfoldertypeID() int32`

GetFkiEzsignfoldertypeID returns the FkiEzsignfoldertypeID field if non-nil, zero value otherwise.

### GetFkiEzsignfoldertypeIDOk

`func (o *EzsigntemplateGetObjectV1ResponseMPayload) GetFkiEzsignfoldertypeIDOk() (*int32, bool)`

GetFkiEzsignfoldertypeIDOk returns a tuple with the FkiEzsignfoldertypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignfoldertypeID

`func (o *EzsigntemplateGetObjectV1ResponseMPayload) SetFkiEzsignfoldertypeID(v int32)`

SetFkiEzsignfoldertypeID sets FkiEzsignfoldertypeID field to given value.


### GetFkiLanguageID

`func (o *EzsigntemplateGetObjectV1ResponseMPayload) GetFkiLanguageID() int32`

GetFkiLanguageID returns the FkiLanguageID field if non-nil, zero value otherwise.

### GetFkiLanguageIDOk

`func (o *EzsigntemplateGetObjectV1ResponseMPayload) GetFkiLanguageIDOk() (*int32, bool)`

GetFkiLanguageIDOk returns a tuple with the FkiLanguageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiLanguageID

`func (o *EzsigntemplateGetObjectV1ResponseMPayload) SetFkiLanguageID(v int32)`

SetFkiLanguageID sets FkiLanguageID field to given value.


### GetSLanguageNameX

`func (o *EzsigntemplateGetObjectV1ResponseMPayload) GetSLanguageNameX() string`

GetSLanguageNameX returns the SLanguageNameX field if non-nil, zero value otherwise.

### GetSLanguageNameXOk

`func (o *EzsigntemplateGetObjectV1ResponseMPayload) GetSLanguageNameXOk() (*string, bool)`

GetSLanguageNameXOk returns a tuple with the SLanguageNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSLanguageNameX

`func (o *EzsigntemplateGetObjectV1ResponseMPayload) SetSLanguageNameX(v string)`

SetSLanguageNameX sets SLanguageNameX field to given value.


### GetSEzsigntemplateDescription

`func (o *EzsigntemplateGetObjectV1ResponseMPayload) GetSEzsigntemplateDescription() string`

GetSEzsigntemplateDescription returns the SEzsigntemplateDescription field if non-nil, zero value otherwise.

### GetSEzsigntemplateDescriptionOk

`func (o *EzsigntemplateGetObjectV1ResponseMPayload) GetSEzsigntemplateDescriptionOk() (*string, bool)`

GetSEzsigntemplateDescriptionOk returns a tuple with the SEzsigntemplateDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplateDescription

`func (o *EzsigntemplateGetObjectV1ResponseMPayload) SetSEzsigntemplateDescription(v string)`

SetSEzsigntemplateDescription sets SEzsigntemplateDescription field to given value.


### GetBEzsigntemplateAdminonly

`func (o *EzsigntemplateGetObjectV1ResponseMPayload) GetBEzsigntemplateAdminonly() bool`

GetBEzsigntemplateAdminonly returns the BEzsigntemplateAdminonly field if non-nil, zero value otherwise.

### GetBEzsigntemplateAdminonlyOk

`func (o *EzsigntemplateGetObjectV1ResponseMPayload) GetBEzsigntemplateAdminonlyOk() (*bool, bool)`

GetBEzsigntemplateAdminonlyOk returns a tuple with the BEzsigntemplateAdminonly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsigntemplateAdminonly

`func (o *EzsigntemplateGetObjectV1ResponseMPayload) SetBEzsigntemplateAdminonly(v bool)`

SetBEzsigntemplateAdminonly sets BEzsigntemplateAdminonly field to given value.


### GetSEzsignfoldertypeNameX

`func (o *EzsigntemplateGetObjectV1ResponseMPayload) GetSEzsignfoldertypeNameX() string`

GetSEzsignfoldertypeNameX returns the SEzsignfoldertypeNameX field if non-nil, zero value otherwise.

### GetSEzsignfoldertypeNameXOk

`func (o *EzsigntemplateGetObjectV1ResponseMPayload) GetSEzsignfoldertypeNameXOk() (*string, bool)`

GetSEzsignfoldertypeNameXOk returns a tuple with the SEzsignfoldertypeNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignfoldertypeNameX

`func (o *EzsigntemplateGetObjectV1ResponseMPayload) SetSEzsignfoldertypeNameX(v string)`

SetSEzsignfoldertypeNameX sets SEzsignfoldertypeNameX field to given value.


### GetObjAudit

`func (o *EzsigntemplateGetObjectV1ResponseMPayload) GetObjAudit() CommonAudit`

GetObjAudit returns the ObjAudit field if non-nil, zero value otherwise.

### GetObjAuditOk

`func (o *EzsigntemplateGetObjectV1ResponseMPayload) GetObjAuditOk() (*CommonAudit, bool)`

GetObjAuditOk returns a tuple with the ObjAudit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjAudit

`func (o *EzsigntemplateGetObjectV1ResponseMPayload) SetObjAudit(v CommonAudit)`

SetObjAudit sets ObjAudit field to given value.


### GetObjEzsigntemplatedocument

`func (o *EzsigntemplateGetObjectV1ResponseMPayload) GetObjEzsigntemplatedocument() EzsigntemplatedocumentResponse`

GetObjEzsigntemplatedocument returns the ObjEzsigntemplatedocument field if non-nil, zero value otherwise.

### GetObjEzsigntemplatedocumentOk

`func (o *EzsigntemplateGetObjectV1ResponseMPayload) GetObjEzsigntemplatedocumentOk() (*EzsigntemplatedocumentResponse, bool)`

GetObjEzsigntemplatedocumentOk returns a tuple with the ObjEzsigntemplatedocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjEzsigntemplatedocument

`func (o *EzsigntemplateGetObjectV1ResponseMPayload) SetObjEzsigntemplatedocument(v EzsigntemplatedocumentResponse)`

SetObjEzsigntemplatedocument sets ObjEzsigntemplatedocument field to given value.

### HasObjEzsigntemplatedocument

`func (o *EzsigntemplateGetObjectV1ResponseMPayload) HasObjEzsigntemplatedocument() bool`

HasObjEzsigntemplatedocument returns a boolean if a field has been set.

### GetAObjEzsigntemplatesigner

`func (o *EzsigntemplateGetObjectV1ResponseMPayload) GetAObjEzsigntemplatesigner() []EzsigntemplatesignerResponseCompound`

GetAObjEzsigntemplatesigner returns the AObjEzsigntemplatesigner field if non-nil, zero value otherwise.

### GetAObjEzsigntemplatesignerOk

`func (o *EzsigntemplateGetObjectV1ResponseMPayload) GetAObjEzsigntemplatesignerOk() (*[]EzsigntemplatesignerResponseCompound, bool)`

GetAObjEzsigntemplatesignerOk returns a tuple with the AObjEzsigntemplatesigner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEzsigntemplatesigner

`func (o *EzsigntemplateGetObjectV1ResponseMPayload) SetAObjEzsigntemplatesigner(v []EzsigntemplatesignerResponseCompound)`

SetAObjEzsigntemplatesigner sets AObjEzsigntemplatesigner field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

