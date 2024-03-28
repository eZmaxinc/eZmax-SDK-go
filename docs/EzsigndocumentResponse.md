# EzsigndocumentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsigndocumentID** | **int32** | The unique ID of the Ezsigndocument | 
**FkiEzsignfolderID** | **int32** | The unique ID of the Ezsignfolder | 
**FkiEzsignfoldersignerassociationIDDeclinedtosign** | Pointer to **int32** | The unique ID of the Ezsignfoldersignerassociation | [optional] 
**DtEzsigndocumentDuedate** | **string** | The maximum date and time at which the Ezsigndocument can be signed. | 
**DtEzsignformCompleted** | Pointer to **string** | The date and time at which the Ezsignform has been completed. | [optional] 
**FkiLanguageID** | Pointer to **int32** | The unique ID of the Language.  Valid values:  |Value|Description| |-|-| |1|French| |2|English| | [optional] 
**SEzsigndocumentName** | **string** | The name of the document that will be presented to Ezsignfoldersignerassociations | 
**EEzsigndocumentStep** | [**FieldEEzsigndocumentStep**](FieldEEzsigndocumentStep.md) |  | 
**DtEzsigndocumentFirstsend** | Pointer to **string** | The date and time when the Ezsigndocument was first sent. | [optional] 
**DtEzsigndocumentLastsend** | Pointer to **string** | The date and time when the Ezsigndocument was sent the last time. | [optional] 
**IEzsigndocumentOrder** | **int32** | The order in which the Ezsigndocument will be presented to the signatory in the Ezsignfolder. | 
**IEzsigndocumentPagetotal** | **int32** | The number of pages in the Ezsigndocument. | 
**IEzsigndocumentSignaturesigned** | **int32** | The number of signatures that were signed in the document. | 
**IEzsigndocumentSignaturetotal** | **int32** | The number of total signatures that were requested in the Ezsigndocument. | 
**SEzsigndocumentMD5initial** | Pointer to **string** | MD5 Hash of the initial PDF Document before signatures were applied to it. | [optional] 
**TEzsigndocumentDeclinedtosignreason** | Pointer to **string** | A custom text message that will contain the refusal message if the Ezsigndocument is declined to sign | [optional] 
**SEzsigndocumentMD5signed** | Pointer to **string** | MD5 Hash of the final PDF Document after all signatures were applied to it. | [optional] 
**BEzsigndocumentEzsignform** | Pointer to **bool** | If the Ezsigndocument contains an Ezsignform or not | [optional] 
**BEzsigndocumentHassignedsignatures** | Pointer to **bool** | If the Ezsigndocument contains signed signatures (From internal or external sources) | [optional] 
**ObjAudit** | Pointer to [**CommonAudit**](CommonAudit.md) |  | [optional] 
**SEzsigndocumentExternalid** | Pointer to **string** | This field can be used to store an External ID from the client&#39;s system.  Anything can be stored in this field, it will never be evaluated by the eZmax system and will be returned AS-IS.  To store multiple values, consider using a JSON formatted structure, a URL encoded string, a CSV or any other custom format.  | [optional] 
**IEzsigndocumentEzsignsignatureattachmenttotal** | **int32** | The number of Ezsigndocumentattachment total | 
**IEzsigndocumentEzsigndiscussiontotal** | **int32** | The total number of Ezsigndiscussions | 

## Methods

### NewEzsigndocumentResponse

`func NewEzsigndocumentResponse(pkiEzsigndocumentID int32, fkiEzsignfolderID int32, dtEzsigndocumentDuedate string, sEzsigndocumentName string, eEzsigndocumentStep FieldEEzsigndocumentStep, iEzsigndocumentOrder int32, iEzsigndocumentPagetotal int32, iEzsigndocumentSignaturesigned int32, iEzsigndocumentSignaturetotal int32, iEzsigndocumentEzsignsignatureattachmenttotal int32, iEzsigndocumentEzsigndiscussiontotal int32, ) *EzsigndocumentResponse`

NewEzsigndocumentResponse instantiates a new EzsigndocumentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigndocumentResponseWithDefaults

`func NewEzsigndocumentResponseWithDefaults() *EzsigndocumentResponse`

NewEzsigndocumentResponseWithDefaults instantiates a new EzsigndocumentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsigndocumentID

`func (o *EzsigndocumentResponse) GetPkiEzsigndocumentID() int32`

GetPkiEzsigndocumentID returns the PkiEzsigndocumentID field if non-nil, zero value otherwise.

### GetPkiEzsigndocumentIDOk

`func (o *EzsigndocumentResponse) GetPkiEzsigndocumentIDOk() (*int32, bool)`

GetPkiEzsigndocumentIDOk returns a tuple with the PkiEzsigndocumentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsigndocumentID

`func (o *EzsigndocumentResponse) SetPkiEzsigndocumentID(v int32)`

SetPkiEzsigndocumentID sets PkiEzsigndocumentID field to given value.


### GetFkiEzsignfolderID

`func (o *EzsigndocumentResponse) GetFkiEzsignfolderID() int32`

GetFkiEzsignfolderID returns the FkiEzsignfolderID field if non-nil, zero value otherwise.

### GetFkiEzsignfolderIDOk

`func (o *EzsigndocumentResponse) GetFkiEzsignfolderIDOk() (*int32, bool)`

GetFkiEzsignfolderIDOk returns a tuple with the FkiEzsignfolderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignfolderID

`func (o *EzsigndocumentResponse) SetFkiEzsignfolderID(v int32)`

SetFkiEzsignfolderID sets FkiEzsignfolderID field to given value.


### GetFkiEzsignfoldersignerassociationIDDeclinedtosign

`func (o *EzsigndocumentResponse) GetFkiEzsignfoldersignerassociationIDDeclinedtosign() int32`

GetFkiEzsignfoldersignerassociationIDDeclinedtosign returns the FkiEzsignfoldersignerassociationIDDeclinedtosign field if non-nil, zero value otherwise.

### GetFkiEzsignfoldersignerassociationIDDeclinedtosignOk

`func (o *EzsigndocumentResponse) GetFkiEzsignfoldersignerassociationIDDeclinedtosignOk() (*int32, bool)`

GetFkiEzsignfoldersignerassociationIDDeclinedtosignOk returns a tuple with the FkiEzsignfoldersignerassociationIDDeclinedtosign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignfoldersignerassociationIDDeclinedtosign

`func (o *EzsigndocumentResponse) SetFkiEzsignfoldersignerassociationIDDeclinedtosign(v int32)`

SetFkiEzsignfoldersignerassociationIDDeclinedtosign sets FkiEzsignfoldersignerassociationIDDeclinedtosign field to given value.

### HasFkiEzsignfoldersignerassociationIDDeclinedtosign

`func (o *EzsigndocumentResponse) HasFkiEzsignfoldersignerassociationIDDeclinedtosign() bool`

HasFkiEzsignfoldersignerassociationIDDeclinedtosign returns a boolean if a field has been set.

### GetDtEzsigndocumentDuedate

`func (o *EzsigndocumentResponse) GetDtEzsigndocumentDuedate() string`

GetDtEzsigndocumentDuedate returns the DtEzsigndocumentDuedate field if non-nil, zero value otherwise.

### GetDtEzsigndocumentDuedateOk

`func (o *EzsigndocumentResponse) GetDtEzsigndocumentDuedateOk() (*string, bool)`

GetDtEzsigndocumentDuedateOk returns a tuple with the DtEzsigndocumentDuedate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtEzsigndocumentDuedate

`func (o *EzsigndocumentResponse) SetDtEzsigndocumentDuedate(v string)`

SetDtEzsigndocumentDuedate sets DtEzsigndocumentDuedate field to given value.


### GetDtEzsignformCompleted

`func (o *EzsigndocumentResponse) GetDtEzsignformCompleted() string`

GetDtEzsignformCompleted returns the DtEzsignformCompleted field if non-nil, zero value otherwise.

### GetDtEzsignformCompletedOk

`func (o *EzsigndocumentResponse) GetDtEzsignformCompletedOk() (*string, bool)`

GetDtEzsignformCompletedOk returns a tuple with the DtEzsignformCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtEzsignformCompleted

`func (o *EzsigndocumentResponse) SetDtEzsignformCompleted(v string)`

SetDtEzsignformCompleted sets DtEzsignformCompleted field to given value.

### HasDtEzsignformCompleted

`func (o *EzsigndocumentResponse) HasDtEzsignformCompleted() bool`

HasDtEzsignformCompleted returns a boolean if a field has been set.

### GetFkiLanguageID

`func (o *EzsigndocumentResponse) GetFkiLanguageID() int32`

GetFkiLanguageID returns the FkiLanguageID field if non-nil, zero value otherwise.

### GetFkiLanguageIDOk

`func (o *EzsigndocumentResponse) GetFkiLanguageIDOk() (*int32, bool)`

GetFkiLanguageIDOk returns a tuple with the FkiLanguageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiLanguageID

`func (o *EzsigndocumentResponse) SetFkiLanguageID(v int32)`

SetFkiLanguageID sets FkiLanguageID field to given value.

### HasFkiLanguageID

`func (o *EzsigndocumentResponse) HasFkiLanguageID() bool`

HasFkiLanguageID returns a boolean if a field has been set.

### GetSEzsigndocumentName

`func (o *EzsigndocumentResponse) GetSEzsigndocumentName() string`

GetSEzsigndocumentName returns the SEzsigndocumentName field if non-nil, zero value otherwise.

### GetSEzsigndocumentNameOk

`func (o *EzsigndocumentResponse) GetSEzsigndocumentNameOk() (*string, bool)`

GetSEzsigndocumentNameOk returns a tuple with the SEzsigndocumentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigndocumentName

`func (o *EzsigndocumentResponse) SetSEzsigndocumentName(v string)`

SetSEzsigndocumentName sets SEzsigndocumentName field to given value.


### GetEEzsigndocumentStep

`func (o *EzsigndocumentResponse) GetEEzsigndocumentStep() FieldEEzsigndocumentStep`

GetEEzsigndocumentStep returns the EEzsigndocumentStep field if non-nil, zero value otherwise.

### GetEEzsigndocumentStepOk

`func (o *EzsigndocumentResponse) GetEEzsigndocumentStepOk() (*FieldEEzsigndocumentStep, bool)`

GetEEzsigndocumentStepOk returns a tuple with the EEzsigndocumentStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigndocumentStep

`func (o *EzsigndocumentResponse) SetEEzsigndocumentStep(v FieldEEzsigndocumentStep)`

SetEEzsigndocumentStep sets EEzsigndocumentStep field to given value.


### GetDtEzsigndocumentFirstsend

`func (o *EzsigndocumentResponse) GetDtEzsigndocumentFirstsend() string`

GetDtEzsigndocumentFirstsend returns the DtEzsigndocumentFirstsend field if non-nil, zero value otherwise.

### GetDtEzsigndocumentFirstsendOk

`func (o *EzsigndocumentResponse) GetDtEzsigndocumentFirstsendOk() (*string, bool)`

GetDtEzsigndocumentFirstsendOk returns a tuple with the DtEzsigndocumentFirstsend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtEzsigndocumentFirstsend

`func (o *EzsigndocumentResponse) SetDtEzsigndocumentFirstsend(v string)`

SetDtEzsigndocumentFirstsend sets DtEzsigndocumentFirstsend field to given value.

### HasDtEzsigndocumentFirstsend

`func (o *EzsigndocumentResponse) HasDtEzsigndocumentFirstsend() bool`

HasDtEzsigndocumentFirstsend returns a boolean if a field has been set.

### GetDtEzsigndocumentLastsend

`func (o *EzsigndocumentResponse) GetDtEzsigndocumentLastsend() string`

GetDtEzsigndocumentLastsend returns the DtEzsigndocumentLastsend field if non-nil, zero value otherwise.

### GetDtEzsigndocumentLastsendOk

`func (o *EzsigndocumentResponse) GetDtEzsigndocumentLastsendOk() (*string, bool)`

GetDtEzsigndocumentLastsendOk returns a tuple with the DtEzsigndocumentLastsend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtEzsigndocumentLastsend

`func (o *EzsigndocumentResponse) SetDtEzsigndocumentLastsend(v string)`

SetDtEzsigndocumentLastsend sets DtEzsigndocumentLastsend field to given value.

### HasDtEzsigndocumentLastsend

`func (o *EzsigndocumentResponse) HasDtEzsigndocumentLastsend() bool`

HasDtEzsigndocumentLastsend returns a boolean if a field has been set.

### GetIEzsigndocumentOrder

`func (o *EzsigndocumentResponse) GetIEzsigndocumentOrder() int32`

GetIEzsigndocumentOrder returns the IEzsigndocumentOrder field if non-nil, zero value otherwise.

### GetIEzsigndocumentOrderOk

`func (o *EzsigndocumentResponse) GetIEzsigndocumentOrderOk() (*int32, bool)`

GetIEzsigndocumentOrderOk returns a tuple with the IEzsigndocumentOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigndocumentOrder

`func (o *EzsigndocumentResponse) SetIEzsigndocumentOrder(v int32)`

SetIEzsigndocumentOrder sets IEzsigndocumentOrder field to given value.


### GetIEzsigndocumentPagetotal

`func (o *EzsigndocumentResponse) GetIEzsigndocumentPagetotal() int32`

GetIEzsigndocumentPagetotal returns the IEzsigndocumentPagetotal field if non-nil, zero value otherwise.

### GetIEzsigndocumentPagetotalOk

`func (o *EzsigndocumentResponse) GetIEzsigndocumentPagetotalOk() (*int32, bool)`

GetIEzsigndocumentPagetotalOk returns a tuple with the IEzsigndocumentPagetotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigndocumentPagetotal

`func (o *EzsigndocumentResponse) SetIEzsigndocumentPagetotal(v int32)`

SetIEzsigndocumentPagetotal sets IEzsigndocumentPagetotal field to given value.


### GetIEzsigndocumentSignaturesigned

`func (o *EzsigndocumentResponse) GetIEzsigndocumentSignaturesigned() int32`

GetIEzsigndocumentSignaturesigned returns the IEzsigndocumentSignaturesigned field if non-nil, zero value otherwise.

### GetIEzsigndocumentSignaturesignedOk

`func (o *EzsigndocumentResponse) GetIEzsigndocumentSignaturesignedOk() (*int32, bool)`

GetIEzsigndocumentSignaturesignedOk returns a tuple with the IEzsigndocumentSignaturesigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigndocumentSignaturesigned

`func (o *EzsigndocumentResponse) SetIEzsigndocumentSignaturesigned(v int32)`

SetIEzsigndocumentSignaturesigned sets IEzsigndocumentSignaturesigned field to given value.


### GetIEzsigndocumentSignaturetotal

`func (o *EzsigndocumentResponse) GetIEzsigndocumentSignaturetotal() int32`

GetIEzsigndocumentSignaturetotal returns the IEzsigndocumentSignaturetotal field if non-nil, zero value otherwise.

### GetIEzsigndocumentSignaturetotalOk

`func (o *EzsigndocumentResponse) GetIEzsigndocumentSignaturetotalOk() (*int32, bool)`

GetIEzsigndocumentSignaturetotalOk returns a tuple with the IEzsigndocumentSignaturetotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigndocumentSignaturetotal

`func (o *EzsigndocumentResponse) SetIEzsigndocumentSignaturetotal(v int32)`

SetIEzsigndocumentSignaturetotal sets IEzsigndocumentSignaturetotal field to given value.


### GetSEzsigndocumentMD5initial

`func (o *EzsigndocumentResponse) GetSEzsigndocumentMD5initial() string`

GetSEzsigndocumentMD5initial returns the SEzsigndocumentMD5initial field if non-nil, zero value otherwise.

### GetSEzsigndocumentMD5initialOk

`func (o *EzsigndocumentResponse) GetSEzsigndocumentMD5initialOk() (*string, bool)`

GetSEzsigndocumentMD5initialOk returns a tuple with the SEzsigndocumentMD5initial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigndocumentMD5initial

`func (o *EzsigndocumentResponse) SetSEzsigndocumentMD5initial(v string)`

SetSEzsigndocumentMD5initial sets SEzsigndocumentMD5initial field to given value.

### HasSEzsigndocumentMD5initial

`func (o *EzsigndocumentResponse) HasSEzsigndocumentMD5initial() bool`

HasSEzsigndocumentMD5initial returns a boolean if a field has been set.

### GetTEzsigndocumentDeclinedtosignreason

`func (o *EzsigndocumentResponse) GetTEzsigndocumentDeclinedtosignreason() string`

GetTEzsigndocumentDeclinedtosignreason returns the TEzsigndocumentDeclinedtosignreason field if non-nil, zero value otherwise.

### GetTEzsigndocumentDeclinedtosignreasonOk

`func (o *EzsigndocumentResponse) GetTEzsigndocumentDeclinedtosignreasonOk() (*string, bool)`

GetTEzsigndocumentDeclinedtosignreasonOk returns a tuple with the TEzsigndocumentDeclinedtosignreason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTEzsigndocumentDeclinedtosignreason

`func (o *EzsigndocumentResponse) SetTEzsigndocumentDeclinedtosignreason(v string)`

SetTEzsigndocumentDeclinedtosignreason sets TEzsigndocumentDeclinedtosignreason field to given value.

### HasTEzsigndocumentDeclinedtosignreason

`func (o *EzsigndocumentResponse) HasTEzsigndocumentDeclinedtosignreason() bool`

HasTEzsigndocumentDeclinedtosignreason returns a boolean if a field has been set.

### GetSEzsigndocumentMD5signed

`func (o *EzsigndocumentResponse) GetSEzsigndocumentMD5signed() string`

GetSEzsigndocumentMD5signed returns the SEzsigndocumentMD5signed field if non-nil, zero value otherwise.

### GetSEzsigndocumentMD5signedOk

`func (o *EzsigndocumentResponse) GetSEzsigndocumentMD5signedOk() (*string, bool)`

GetSEzsigndocumentMD5signedOk returns a tuple with the SEzsigndocumentMD5signed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigndocumentMD5signed

`func (o *EzsigndocumentResponse) SetSEzsigndocumentMD5signed(v string)`

SetSEzsigndocumentMD5signed sets SEzsigndocumentMD5signed field to given value.

### HasSEzsigndocumentMD5signed

`func (o *EzsigndocumentResponse) HasSEzsigndocumentMD5signed() bool`

HasSEzsigndocumentMD5signed returns a boolean if a field has been set.

### GetBEzsigndocumentEzsignform

`func (o *EzsigndocumentResponse) GetBEzsigndocumentEzsignform() bool`

GetBEzsigndocumentEzsignform returns the BEzsigndocumentEzsignform field if non-nil, zero value otherwise.

### GetBEzsigndocumentEzsignformOk

`func (o *EzsigndocumentResponse) GetBEzsigndocumentEzsignformOk() (*bool, bool)`

GetBEzsigndocumentEzsignformOk returns a tuple with the BEzsigndocumentEzsignform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsigndocumentEzsignform

`func (o *EzsigndocumentResponse) SetBEzsigndocumentEzsignform(v bool)`

SetBEzsigndocumentEzsignform sets BEzsigndocumentEzsignform field to given value.

### HasBEzsigndocumentEzsignform

`func (o *EzsigndocumentResponse) HasBEzsigndocumentEzsignform() bool`

HasBEzsigndocumentEzsignform returns a boolean if a field has been set.

### GetBEzsigndocumentHassignedsignatures

`func (o *EzsigndocumentResponse) GetBEzsigndocumentHassignedsignatures() bool`

GetBEzsigndocumentHassignedsignatures returns the BEzsigndocumentHassignedsignatures field if non-nil, zero value otherwise.

### GetBEzsigndocumentHassignedsignaturesOk

`func (o *EzsigndocumentResponse) GetBEzsigndocumentHassignedsignaturesOk() (*bool, bool)`

GetBEzsigndocumentHassignedsignaturesOk returns a tuple with the BEzsigndocumentHassignedsignatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsigndocumentHassignedsignatures

`func (o *EzsigndocumentResponse) SetBEzsigndocumentHassignedsignatures(v bool)`

SetBEzsigndocumentHassignedsignatures sets BEzsigndocumentHassignedsignatures field to given value.

### HasBEzsigndocumentHassignedsignatures

`func (o *EzsigndocumentResponse) HasBEzsigndocumentHassignedsignatures() bool`

HasBEzsigndocumentHassignedsignatures returns a boolean if a field has been set.

### GetObjAudit

`func (o *EzsigndocumentResponse) GetObjAudit() CommonAudit`

GetObjAudit returns the ObjAudit field if non-nil, zero value otherwise.

### GetObjAuditOk

`func (o *EzsigndocumentResponse) GetObjAuditOk() (*CommonAudit, bool)`

GetObjAuditOk returns a tuple with the ObjAudit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjAudit

`func (o *EzsigndocumentResponse) SetObjAudit(v CommonAudit)`

SetObjAudit sets ObjAudit field to given value.

### HasObjAudit

`func (o *EzsigndocumentResponse) HasObjAudit() bool`

HasObjAudit returns a boolean if a field has been set.

### GetSEzsigndocumentExternalid

`func (o *EzsigndocumentResponse) GetSEzsigndocumentExternalid() string`

GetSEzsigndocumentExternalid returns the SEzsigndocumentExternalid field if non-nil, zero value otherwise.

### GetSEzsigndocumentExternalidOk

`func (o *EzsigndocumentResponse) GetSEzsigndocumentExternalidOk() (*string, bool)`

GetSEzsigndocumentExternalidOk returns a tuple with the SEzsigndocumentExternalid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigndocumentExternalid

`func (o *EzsigndocumentResponse) SetSEzsigndocumentExternalid(v string)`

SetSEzsigndocumentExternalid sets SEzsigndocumentExternalid field to given value.

### HasSEzsigndocumentExternalid

`func (o *EzsigndocumentResponse) HasSEzsigndocumentExternalid() bool`

HasSEzsigndocumentExternalid returns a boolean if a field has been set.

### GetIEzsigndocumentEzsignsignatureattachmenttotal

`func (o *EzsigndocumentResponse) GetIEzsigndocumentEzsignsignatureattachmenttotal() int32`

GetIEzsigndocumentEzsignsignatureattachmenttotal returns the IEzsigndocumentEzsignsignatureattachmenttotal field if non-nil, zero value otherwise.

### GetIEzsigndocumentEzsignsignatureattachmenttotalOk

`func (o *EzsigndocumentResponse) GetIEzsigndocumentEzsignsignatureattachmenttotalOk() (*int32, bool)`

GetIEzsigndocumentEzsignsignatureattachmenttotalOk returns a tuple with the IEzsigndocumentEzsignsignatureattachmenttotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigndocumentEzsignsignatureattachmenttotal

`func (o *EzsigndocumentResponse) SetIEzsigndocumentEzsignsignatureattachmenttotal(v int32)`

SetIEzsigndocumentEzsignsignatureattachmenttotal sets IEzsigndocumentEzsignsignatureattachmenttotal field to given value.


### GetIEzsigndocumentEzsigndiscussiontotal

`func (o *EzsigndocumentResponse) GetIEzsigndocumentEzsigndiscussiontotal() int32`

GetIEzsigndocumentEzsigndiscussiontotal returns the IEzsigndocumentEzsigndiscussiontotal field if non-nil, zero value otherwise.

### GetIEzsigndocumentEzsigndiscussiontotalOk

`func (o *EzsigndocumentResponse) GetIEzsigndocumentEzsigndiscussiontotalOk() (*int32, bool)`

GetIEzsigndocumentEzsigndiscussiontotalOk returns a tuple with the IEzsigndocumentEzsigndiscussiontotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigndocumentEzsigndiscussiontotal

`func (o *EzsigndocumentResponse) SetIEzsigndocumentEzsigndiscussiontotal(v int32)`

SetIEzsigndocumentEzsigndiscussiontotal sets IEzsigndocumentEzsigndiscussiontotal field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


