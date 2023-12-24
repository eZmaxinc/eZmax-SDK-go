# EzsigndocumentResponseCompound

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
**EEzsigndocumentSteptype** | [**ComputedEEzsigndocumentSteptype**](ComputedEEzsigndocumentSteptype.md) |  | 
**IEzsigndocumentStepformtotal** | **int32** | The total number of steps in the form filling phase | 
**IEzsigndocumentStepformcurrent** | **int32** | The current step in the form filling phase | 
**IEzsigndocumentStepsignaturetotal** | **int32** | The total number of steps in the signature filling phase | 
**IEzsigndocumentStepsignatureCurrent** | **int32** | The current step in the signature phase | 
**AObjEzsignfoldersignerassociationstatus** | [**[]CustomEzsignfoldersignerassociationstatusResponse**](CustomEzsignfoldersignerassociationstatusResponse.md) |  | 

## Methods

### NewEzsigndocumentResponseCompound

`func NewEzsigndocumentResponseCompound(pkiEzsigndocumentID int32, fkiEzsignfolderID int32, dtEzsigndocumentDuedate string, sEzsigndocumentName string, eEzsigndocumentStep FieldEEzsigndocumentStep, iEzsigndocumentOrder int32, iEzsigndocumentPagetotal int32, iEzsigndocumentSignaturesigned int32, iEzsigndocumentSignaturetotal int32, iEzsigndocumentEzsignsignatureattachmenttotal int32, eEzsigndocumentSteptype ComputedEEzsigndocumentSteptype, iEzsigndocumentStepformtotal int32, iEzsigndocumentStepformcurrent int32, iEzsigndocumentStepsignaturetotal int32, iEzsigndocumentStepsignatureCurrent int32, aObjEzsignfoldersignerassociationstatus []CustomEzsignfoldersignerassociationstatusResponse, ) *EzsigndocumentResponseCompound`

NewEzsigndocumentResponseCompound instantiates a new EzsigndocumentResponseCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigndocumentResponseCompoundWithDefaults

`func NewEzsigndocumentResponseCompoundWithDefaults() *EzsigndocumentResponseCompound`

NewEzsigndocumentResponseCompoundWithDefaults instantiates a new EzsigndocumentResponseCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsigndocumentID

`func (o *EzsigndocumentResponseCompound) GetPkiEzsigndocumentID() int32`

GetPkiEzsigndocumentID returns the PkiEzsigndocumentID field if non-nil, zero value otherwise.

### GetPkiEzsigndocumentIDOk

`func (o *EzsigndocumentResponseCompound) GetPkiEzsigndocumentIDOk() (*int32, bool)`

GetPkiEzsigndocumentIDOk returns a tuple with the PkiEzsigndocumentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsigndocumentID

`func (o *EzsigndocumentResponseCompound) SetPkiEzsigndocumentID(v int32)`

SetPkiEzsigndocumentID sets PkiEzsigndocumentID field to given value.


### GetFkiEzsignfolderID

`func (o *EzsigndocumentResponseCompound) GetFkiEzsignfolderID() int32`

GetFkiEzsignfolderID returns the FkiEzsignfolderID field if non-nil, zero value otherwise.

### GetFkiEzsignfolderIDOk

`func (o *EzsigndocumentResponseCompound) GetFkiEzsignfolderIDOk() (*int32, bool)`

GetFkiEzsignfolderIDOk returns a tuple with the FkiEzsignfolderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignfolderID

`func (o *EzsigndocumentResponseCompound) SetFkiEzsignfolderID(v int32)`

SetFkiEzsignfolderID sets FkiEzsignfolderID field to given value.


### GetFkiEzsignfoldersignerassociationIDDeclinedtosign

`func (o *EzsigndocumentResponseCompound) GetFkiEzsignfoldersignerassociationIDDeclinedtosign() int32`

GetFkiEzsignfoldersignerassociationIDDeclinedtosign returns the FkiEzsignfoldersignerassociationIDDeclinedtosign field if non-nil, zero value otherwise.

### GetFkiEzsignfoldersignerassociationIDDeclinedtosignOk

`func (o *EzsigndocumentResponseCompound) GetFkiEzsignfoldersignerassociationIDDeclinedtosignOk() (*int32, bool)`

GetFkiEzsignfoldersignerassociationIDDeclinedtosignOk returns a tuple with the FkiEzsignfoldersignerassociationIDDeclinedtosign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignfoldersignerassociationIDDeclinedtosign

`func (o *EzsigndocumentResponseCompound) SetFkiEzsignfoldersignerassociationIDDeclinedtosign(v int32)`

SetFkiEzsignfoldersignerassociationIDDeclinedtosign sets FkiEzsignfoldersignerassociationIDDeclinedtosign field to given value.

### HasFkiEzsignfoldersignerassociationIDDeclinedtosign

`func (o *EzsigndocumentResponseCompound) HasFkiEzsignfoldersignerassociationIDDeclinedtosign() bool`

HasFkiEzsignfoldersignerassociationIDDeclinedtosign returns a boolean if a field has been set.

### GetDtEzsigndocumentDuedate

`func (o *EzsigndocumentResponseCompound) GetDtEzsigndocumentDuedate() string`

GetDtEzsigndocumentDuedate returns the DtEzsigndocumentDuedate field if non-nil, zero value otherwise.

### GetDtEzsigndocumentDuedateOk

`func (o *EzsigndocumentResponseCompound) GetDtEzsigndocumentDuedateOk() (*string, bool)`

GetDtEzsigndocumentDuedateOk returns a tuple with the DtEzsigndocumentDuedate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtEzsigndocumentDuedate

`func (o *EzsigndocumentResponseCompound) SetDtEzsigndocumentDuedate(v string)`

SetDtEzsigndocumentDuedate sets DtEzsigndocumentDuedate field to given value.


### GetDtEzsignformCompleted

`func (o *EzsigndocumentResponseCompound) GetDtEzsignformCompleted() string`

GetDtEzsignformCompleted returns the DtEzsignformCompleted field if non-nil, zero value otherwise.

### GetDtEzsignformCompletedOk

`func (o *EzsigndocumentResponseCompound) GetDtEzsignformCompletedOk() (*string, bool)`

GetDtEzsignformCompletedOk returns a tuple with the DtEzsignformCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtEzsignformCompleted

`func (o *EzsigndocumentResponseCompound) SetDtEzsignformCompleted(v string)`

SetDtEzsignformCompleted sets DtEzsignformCompleted field to given value.

### HasDtEzsignformCompleted

`func (o *EzsigndocumentResponseCompound) HasDtEzsignformCompleted() bool`

HasDtEzsignformCompleted returns a boolean if a field has been set.

### GetFkiLanguageID

`func (o *EzsigndocumentResponseCompound) GetFkiLanguageID() int32`

GetFkiLanguageID returns the FkiLanguageID field if non-nil, zero value otherwise.

### GetFkiLanguageIDOk

`func (o *EzsigndocumentResponseCompound) GetFkiLanguageIDOk() (*int32, bool)`

GetFkiLanguageIDOk returns a tuple with the FkiLanguageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiLanguageID

`func (o *EzsigndocumentResponseCompound) SetFkiLanguageID(v int32)`

SetFkiLanguageID sets FkiLanguageID field to given value.

### HasFkiLanguageID

`func (o *EzsigndocumentResponseCompound) HasFkiLanguageID() bool`

HasFkiLanguageID returns a boolean if a field has been set.

### GetSEzsigndocumentName

`func (o *EzsigndocumentResponseCompound) GetSEzsigndocumentName() string`

GetSEzsigndocumentName returns the SEzsigndocumentName field if non-nil, zero value otherwise.

### GetSEzsigndocumentNameOk

`func (o *EzsigndocumentResponseCompound) GetSEzsigndocumentNameOk() (*string, bool)`

GetSEzsigndocumentNameOk returns a tuple with the SEzsigndocumentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigndocumentName

`func (o *EzsigndocumentResponseCompound) SetSEzsigndocumentName(v string)`

SetSEzsigndocumentName sets SEzsigndocumentName field to given value.


### GetEEzsigndocumentStep

`func (o *EzsigndocumentResponseCompound) GetEEzsigndocumentStep() FieldEEzsigndocumentStep`

GetEEzsigndocumentStep returns the EEzsigndocumentStep field if non-nil, zero value otherwise.

### GetEEzsigndocumentStepOk

`func (o *EzsigndocumentResponseCompound) GetEEzsigndocumentStepOk() (*FieldEEzsigndocumentStep, bool)`

GetEEzsigndocumentStepOk returns a tuple with the EEzsigndocumentStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigndocumentStep

`func (o *EzsigndocumentResponseCompound) SetEEzsigndocumentStep(v FieldEEzsigndocumentStep)`

SetEEzsigndocumentStep sets EEzsigndocumentStep field to given value.


### GetDtEzsigndocumentFirstsend

`func (o *EzsigndocumentResponseCompound) GetDtEzsigndocumentFirstsend() string`

GetDtEzsigndocumentFirstsend returns the DtEzsigndocumentFirstsend field if non-nil, zero value otherwise.

### GetDtEzsigndocumentFirstsendOk

`func (o *EzsigndocumentResponseCompound) GetDtEzsigndocumentFirstsendOk() (*string, bool)`

GetDtEzsigndocumentFirstsendOk returns a tuple with the DtEzsigndocumentFirstsend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtEzsigndocumentFirstsend

`func (o *EzsigndocumentResponseCompound) SetDtEzsigndocumentFirstsend(v string)`

SetDtEzsigndocumentFirstsend sets DtEzsigndocumentFirstsend field to given value.

### HasDtEzsigndocumentFirstsend

`func (o *EzsigndocumentResponseCompound) HasDtEzsigndocumentFirstsend() bool`

HasDtEzsigndocumentFirstsend returns a boolean if a field has been set.

### GetDtEzsigndocumentLastsend

`func (o *EzsigndocumentResponseCompound) GetDtEzsigndocumentLastsend() string`

GetDtEzsigndocumentLastsend returns the DtEzsigndocumentLastsend field if non-nil, zero value otherwise.

### GetDtEzsigndocumentLastsendOk

`func (o *EzsigndocumentResponseCompound) GetDtEzsigndocumentLastsendOk() (*string, bool)`

GetDtEzsigndocumentLastsendOk returns a tuple with the DtEzsigndocumentLastsend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtEzsigndocumentLastsend

`func (o *EzsigndocumentResponseCompound) SetDtEzsigndocumentLastsend(v string)`

SetDtEzsigndocumentLastsend sets DtEzsigndocumentLastsend field to given value.

### HasDtEzsigndocumentLastsend

`func (o *EzsigndocumentResponseCompound) HasDtEzsigndocumentLastsend() bool`

HasDtEzsigndocumentLastsend returns a boolean if a field has been set.

### GetIEzsigndocumentOrder

`func (o *EzsigndocumentResponseCompound) GetIEzsigndocumentOrder() int32`

GetIEzsigndocumentOrder returns the IEzsigndocumentOrder field if non-nil, zero value otherwise.

### GetIEzsigndocumentOrderOk

`func (o *EzsigndocumentResponseCompound) GetIEzsigndocumentOrderOk() (*int32, bool)`

GetIEzsigndocumentOrderOk returns a tuple with the IEzsigndocumentOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigndocumentOrder

`func (o *EzsigndocumentResponseCompound) SetIEzsigndocumentOrder(v int32)`

SetIEzsigndocumentOrder sets IEzsigndocumentOrder field to given value.


### GetIEzsigndocumentPagetotal

`func (o *EzsigndocumentResponseCompound) GetIEzsigndocumentPagetotal() int32`

GetIEzsigndocumentPagetotal returns the IEzsigndocumentPagetotal field if non-nil, zero value otherwise.

### GetIEzsigndocumentPagetotalOk

`func (o *EzsigndocumentResponseCompound) GetIEzsigndocumentPagetotalOk() (*int32, bool)`

GetIEzsigndocumentPagetotalOk returns a tuple with the IEzsigndocumentPagetotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigndocumentPagetotal

`func (o *EzsigndocumentResponseCompound) SetIEzsigndocumentPagetotal(v int32)`

SetIEzsigndocumentPagetotal sets IEzsigndocumentPagetotal field to given value.


### GetIEzsigndocumentSignaturesigned

`func (o *EzsigndocumentResponseCompound) GetIEzsigndocumentSignaturesigned() int32`

GetIEzsigndocumentSignaturesigned returns the IEzsigndocumentSignaturesigned field if non-nil, zero value otherwise.

### GetIEzsigndocumentSignaturesignedOk

`func (o *EzsigndocumentResponseCompound) GetIEzsigndocumentSignaturesignedOk() (*int32, bool)`

GetIEzsigndocumentSignaturesignedOk returns a tuple with the IEzsigndocumentSignaturesigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigndocumentSignaturesigned

`func (o *EzsigndocumentResponseCompound) SetIEzsigndocumentSignaturesigned(v int32)`

SetIEzsigndocumentSignaturesigned sets IEzsigndocumentSignaturesigned field to given value.


### GetIEzsigndocumentSignaturetotal

`func (o *EzsigndocumentResponseCompound) GetIEzsigndocumentSignaturetotal() int32`

GetIEzsigndocumentSignaturetotal returns the IEzsigndocumentSignaturetotal field if non-nil, zero value otherwise.

### GetIEzsigndocumentSignaturetotalOk

`func (o *EzsigndocumentResponseCompound) GetIEzsigndocumentSignaturetotalOk() (*int32, bool)`

GetIEzsigndocumentSignaturetotalOk returns a tuple with the IEzsigndocumentSignaturetotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigndocumentSignaturetotal

`func (o *EzsigndocumentResponseCompound) SetIEzsigndocumentSignaturetotal(v int32)`

SetIEzsigndocumentSignaturetotal sets IEzsigndocumentSignaturetotal field to given value.


### GetSEzsigndocumentMD5initial

`func (o *EzsigndocumentResponseCompound) GetSEzsigndocumentMD5initial() string`

GetSEzsigndocumentMD5initial returns the SEzsigndocumentMD5initial field if non-nil, zero value otherwise.

### GetSEzsigndocumentMD5initialOk

`func (o *EzsigndocumentResponseCompound) GetSEzsigndocumentMD5initialOk() (*string, bool)`

GetSEzsigndocumentMD5initialOk returns a tuple with the SEzsigndocumentMD5initial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigndocumentMD5initial

`func (o *EzsigndocumentResponseCompound) SetSEzsigndocumentMD5initial(v string)`

SetSEzsigndocumentMD5initial sets SEzsigndocumentMD5initial field to given value.

### HasSEzsigndocumentMD5initial

`func (o *EzsigndocumentResponseCompound) HasSEzsigndocumentMD5initial() bool`

HasSEzsigndocumentMD5initial returns a boolean if a field has been set.

### GetTEzsigndocumentDeclinedtosignreason

`func (o *EzsigndocumentResponseCompound) GetTEzsigndocumentDeclinedtosignreason() string`

GetTEzsigndocumentDeclinedtosignreason returns the TEzsigndocumentDeclinedtosignreason field if non-nil, zero value otherwise.

### GetTEzsigndocumentDeclinedtosignreasonOk

`func (o *EzsigndocumentResponseCompound) GetTEzsigndocumentDeclinedtosignreasonOk() (*string, bool)`

GetTEzsigndocumentDeclinedtosignreasonOk returns a tuple with the TEzsigndocumentDeclinedtosignreason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTEzsigndocumentDeclinedtosignreason

`func (o *EzsigndocumentResponseCompound) SetTEzsigndocumentDeclinedtosignreason(v string)`

SetTEzsigndocumentDeclinedtosignreason sets TEzsigndocumentDeclinedtosignreason field to given value.

### HasTEzsigndocumentDeclinedtosignreason

`func (o *EzsigndocumentResponseCompound) HasTEzsigndocumentDeclinedtosignreason() bool`

HasTEzsigndocumentDeclinedtosignreason returns a boolean if a field has been set.

### GetSEzsigndocumentMD5signed

`func (o *EzsigndocumentResponseCompound) GetSEzsigndocumentMD5signed() string`

GetSEzsigndocumentMD5signed returns the SEzsigndocumentMD5signed field if non-nil, zero value otherwise.

### GetSEzsigndocumentMD5signedOk

`func (o *EzsigndocumentResponseCompound) GetSEzsigndocumentMD5signedOk() (*string, bool)`

GetSEzsigndocumentMD5signedOk returns a tuple with the SEzsigndocumentMD5signed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigndocumentMD5signed

`func (o *EzsigndocumentResponseCompound) SetSEzsigndocumentMD5signed(v string)`

SetSEzsigndocumentMD5signed sets SEzsigndocumentMD5signed field to given value.

### HasSEzsigndocumentMD5signed

`func (o *EzsigndocumentResponseCompound) HasSEzsigndocumentMD5signed() bool`

HasSEzsigndocumentMD5signed returns a boolean if a field has been set.

### GetBEzsigndocumentEzsignform

`func (o *EzsigndocumentResponseCompound) GetBEzsigndocumentEzsignform() bool`

GetBEzsigndocumentEzsignform returns the BEzsigndocumentEzsignform field if non-nil, zero value otherwise.

### GetBEzsigndocumentEzsignformOk

`func (o *EzsigndocumentResponseCompound) GetBEzsigndocumentEzsignformOk() (*bool, bool)`

GetBEzsigndocumentEzsignformOk returns a tuple with the BEzsigndocumentEzsignform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsigndocumentEzsignform

`func (o *EzsigndocumentResponseCompound) SetBEzsigndocumentEzsignform(v bool)`

SetBEzsigndocumentEzsignform sets BEzsigndocumentEzsignform field to given value.

### HasBEzsigndocumentEzsignform

`func (o *EzsigndocumentResponseCompound) HasBEzsigndocumentEzsignform() bool`

HasBEzsigndocumentEzsignform returns a boolean if a field has been set.

### GetBEzsigndocumentHassignedsignatures

`func (o *EzsigndocumentResponseCompound) GetBEzsigndocumentHassignedsignatures() bool`

GetBEzsigndocumentHassignedsignatures returns the BEzsigndocumentHassignedsignatures field if non-nil, zero value otherwise.

### GetBEzsigndocumentHassignedsignaturesOk

`func (o *EzsigndocumentResponseCompound) GetBEzsigndocumentHassignedsignaturesOk() (*bool, bool)`

GetBEzsigndocumentHassignedsignaturesOk returns a tuple with the BEzsigndocumentHassignedsignatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsigndocumentHassignedsignatures

`func (o *EzsigndocumentResponseCompound) SetBEzsigndocumentHassignedsignatures(v bool)`

SetBEzsigndocumentHassignedsignatures sets BEzsigndocumentHassignedsignatures field to given value.

### HasBEzsigndocumentHassignedsignatures

`func (o *EzsigndocumentResponseCompound) HasBEzsigndocumentHassignedsignatures() bool`

HasBEzsigndocumentHassignedsignatures returns a boolean if a field has been set.

### GetObjAudit

`func (o *EzsigndocumentResponseCompound) GetObjAudit() CommonAudit`

GetObjAudit returns the ObjAudit field if non-nil, zero value otherwise.

### GetObjAuditOk

`func (o *EzsigndocumentResponseCompound) GetObjAuditOk() (*CommonAudit, bool)`

GetObjAuditOk returns a tuple with the ObjAudit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjAudit

`func (o *EzsigndocumentResponseCompound) SetObjAudit(v CommonAudit)`

SetObjAudit sets ObjAudit field to given value.

### HasObjAudit

`func (o *EzsigndocumentResponseCompound) HasObjAudit() bool`

HasObjAudit returns a boolean if a field has been set.

### GetSEzsigndocumentExternalid

`func (o *EzsigndocumentResponseCompound) GetSEzsigndocumentExternalid() string`

GetSEzsigndocumentExternalid returns the SEzsigndocumentExternalid field if non-nil, zero value otherwise.

### GetSEzsigndocumentExternalidOk

`func (o *EzsigndocumentResponseCompound) GetSEzsigndocumentExternalidOk() (*string, bool)`

GetSEzsigndocumentExternalidOk returns a tuple with the SEzsigndocumentExternalid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigndocumentExternalid

`func (o *EzsigndocumentResponseCompound) SetSEzsigndocumentExternalid(v string)`

SetSEzsigndocumentExternalid sets SEzsigndocumentExternalid field to given value.

### HasSEzsigndocumentExternalid

`func (o *EzsigndocumentResponseCompound) HasSEzsigndocumentExternalid() bool`

HasSEzsigndocumentExternalid returns a boolean if a field has been set.

### GetIEzsigndocumentEzsignsignatureattachmenttotal

`func (o *EzsigndocumentResponseCompound) GetIEzsigndocumentEzsignsignatureattachmenttotal() int32`

GetIEzsigndocumentEzsignsignatureattachmenttotal returns the IEzsigndocumentEzsignsignatureattachmenttotal field if non-nil, zero value otherwise.

### GetIEzsigndocumentEzsignsignatureattachmenttotalOk

`func (o *EzsigndocumentResponseCompound) GetIEzsigndocumentEzsignsignatureattachmenttotalOk() (*int32, bool)`

GetIEzsigndocumentEzsignsignatureattachmenttotalOk returns a tuple with the IEzsigndocumentEzsignsignatureattachmenttotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigndocumentEzsignsignatureattachmenttotal

`func (o *EzsigndocumentResponseCompound) SetIEzsigndocumentEzsignsignatureattachmenttotal(v int32)`

SetIEzsigndocumentEzsignsignatureattachmenttotal sets IEzsigndocumentEzsignsignatureattachmenttotal field to given value.


### GetEEzsigndocumentSteptype

`func (o *EzsigndocumentResponseCompound) GetEEzsigndocumentSteptype() ComputedEEzsigndocumentSteptype`

GetEEzsigndocumentSteptype returns the EEzsigndocumentSteptype field if non-nil, zero value otherwise.

### GetEEzsigndocumentSteptypeOk

`func (o *EzsigndocumentResponseCompound) GetEEzsigndocumentSteptypeOk() (*ComputedEEzsigndocumentSteptype, bool)`

GetEEzsigndocumentSteptypeOk returns a tuple with the EEzsigndocumentSteptype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigndocumentSteptype

`func (o *EzsigndocumentResponseCompound) SetEEzsigndocumentSteptype(v ComputedEEzsigndocumentSteptype)`

SetEEzsigndocumentSteptype sets EEzsigndocumentSteptype field to given value.


### GetIEzsigndocumentStepformtotal

`func (o *EzsigndocumentResponseCompound) GetIEzsigndocumentStepformtotal() int32`

GetIEzsigndocumentStepformtotal returns the IEzsigndocumentStepformtotal field if non-nil, zero value otherwise.

### GetIEzsigndocumentStepformtotalOk

`func (o *EzsigndocumentResponseCompound) GetIEzsigndocumentStepformtotalOk() (*int32, bool)`

GetIEzsigndocumentStepformtotalOk returns a tuple with the IEzsigndocumentStepformtotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigndocumentStepformtotal

`func (o *EzsigndocumentResponseCompound) SetIEzsigndocumentStepformtotal(v int32)`

SetIEzsigndocumentStepformtotal sets IEzsigndocumentStepformtotal field to given value.


### GetIEzsigndocumentStepformcurrent

`func (o *EzsigndocumentResponseCompound) GetIEzsigndocumentStepformcurrent() int32`

GetIEzsigndocumentStepformcurrent returns the IEzsigndocumentStepformcurrent field if non-nil, zero value otherwise.

### GetIEzsigndocumentStepformcurrentOk

`func (o *EzsigndocumentResponseCompound) GetIEzsigndocumentStepformcurrentOk() (*int32, bool)`

GetIEzsigndocumentStepformcurrentOk returns a tuple with the IEzsigndocumentStepformcurrent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigndocumentStepformcurrent

`func (o *EzsigndocumentResponseCompound) SetIEzsigndocumentStepformcurrent(v int32)`

SetIEzsigndocumentStepformcurrent sets IEzsigndocumentStepformcurrent field to given value.


### GetIEzsigndocumentStepsignaturetotal

`func (o *EzsigndocumentResponseCompound) GetIEzsigndocumentStepsignaturetotal() int32`

GetIEzsigndocumentStepsignaturetotal returns the IEzsigndocumentStepsignaturetotal field if non-nil, zero value otherwise.

### GetIEzsigndocumentStepsignaturetotalOk

`func (o *EzsigndocumentResponseCompound) GetIEzsigndocumentStepsignaturetotalOk() (*int32, bool)`

GetIEzsigndocumentStepsignaturetotalOk returns a tuple with the IEzsigndocumentStepsignaturetotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigndocumentStepsignaturetotal

`func (o *EzsigndocumentResponseCompound) SetIEzsigndocumentStepsignaturetotal(v int32)`

SetIEzsigndocumentStepsignaturetotal sets IEzsigndocumentStepsignaturetotal field to given value.


### GetIEzsigndocumentStepsignatureCurrent

`func (o *EzsigndocumentResponseCompound) GetIEzsigndocumentStepsignatureCurrent() int32`

GetIEzsigndocumentStepsignatureCurrent returns the IEzsigndocumentStepsignatureCurrent field if non-nil, zero value otherwise.

### GetIEzsigndocumentStepsignatureCurrentOk

`func (o *EzsigndocumentResponseCompound) GetIEzsigndocumentStepsignatureCurrentOk() (*int32, bool)`

GetIEzsigndocumentStepsignatureCurrentOk returns a tuple with the IEzsigndocumentStepsignatureCurrent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigndocumentStepsignatureCurrent

`func (o *EzsigndocumentResponseCompound) SetIEzsigndocumentStepsignatureCurrent(v int32)`

SetIEzsigndocumentStepsignatureCurrent sets IEzsigndocumentStepsignatureCurrent field to given value.


### GetAObjEzsignfoldersignerassociationstatus

`func (o *EzsigndocumentResponseCompound) GetAObjEzsignfoldersignerassociationstatus() []CustomEzsignfoldersignerassociationstatusResponse`

GetAObjEzsignfoldersignerassociationstatus returns the AObjEzsignfoldersignerassociationstatus field if non-nil, zero value otherwise.

### GetAObjEzsignfoldersignerassociationstatusOk

`func (o *EzsigndocumentResponseCompound) GetAObjEzsignfoldersignerassociationstatusOk() (*[]CustomEzsignfoldersignerassociationstatusResponse, bool)`

GetAObjEzsignfoldersignerassociationstatusOk returns a tuple with the AObjEzsignfoldersignerassociationstatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEzsignfoldersignerassociationstatus

`func (o *EzsigndocumentResponseCompound) SetAObjEzsignfoldersignerassociationstatus(v []CustomEzsignfoldersignerassociationstatusResponse)`

SetAObjEzsignfoldersignerassociationstatus sets AObjEzsignfoldersignerassociationstatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


