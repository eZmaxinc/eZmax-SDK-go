# EzsigndocumentGetObjectV1ResponseMPayload

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

### NewEzsigndocumentGetObjectV1ResponseMPayload

`func NewEzsigndocumentGetObjectV1ResponseMPayload(pkiEzsigndocumentID int32, fkiEzsignfolderID int32, dtEzsigndocumentDuedate string, sEzsigndocumentName string, eEzsigndocumentStep FieldEEzsigndocumentStep, iEzsigndocumentOrder int32, iEzsigndocumentPagetotal int32, iEzsigndocumentSignaturesigned int32, iEzsigndocumentSignaturetotal int32, iEzsigndocumentEzsignsignatureattachmenttotal int32, eEzsigndocumentSteptype ComputedEEzsigndocumentSteptype, iEzsigndocumentStepformtotal int32, iEzsigndocumentStepformcurrent int32, iEzsigndocumentStepsignaturetotal int32, iEzsigndocumentStepsignatureCurrent int32, aObjEzsignfoldersignerassociationstatus []CustomEzsignfoldersignerassociationstatusResponse, ) *EzsigndocumentGetObjectV1ResponseMPayload`

NewEzsigndocumentGetObjectV1ResponseMPayload instantiates a new EzsigndocumentGetObjectV1ResponseMPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigndocumentGetObjectV1ResponseMPayloadWithDefaults

`func NewEzsigndocumentGetObjectV1ResponseMPayloadWithDefaults() *EzsigndocumentGetObjectV1ResponseMPayload`

NewEzsigndocumentGetObjectV1ResponseMPayloadWithDefaults instantiates a new EzsigndocumentGetObjectV1ResponseMPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsigndocumentID

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) GetPkiEzsigndocumentID() int32`

GetPkiEzsigndocumentID returns the PkiEzsigndocumentID field if non-nil, zero value otherwise.

### GetPkiEzsigndocumentIDOk

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) GetPkiEzsigndocumentIDOk() (*int32, bool)`

GetPkiEzsigndocumentIDOk returns a tuple with the PkiEzsigndocumentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsigndocumentID

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) SetPkiEzsigndocumentID(v int32)`

SetPkiEzsigndocumentID sets PkiEzsigndocumentID field to given value.


### GetFkiEzsignfolderID

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) GetFkiEzsignfolderID() int32`

GetFkiEzsignfolderID returns the FkiEzsignfolderID field if non-nil, zero value otherwise.

### GetFkiEzsignfolderIDOk

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) GetFkiEzsignfolderIDOk() (*int32, bool)`

GetFkiEzsignfolderIDOk returns a tuple with the FkiEzsignfolderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignfolderID

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) SetFkiEzsignfolderID(v int32)`

SetFkiEzsignfolderID sets FkiEzsignfolderID field to given value.


### GetFkiEzsignfoldersignerassociationIDDeclinedtosign

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) GetFkiEzsignfoldersignerassociationIDDeclinedtosign() int32`

GetFkiEzsignfoldersignerassociationIDDeclinedtosign returns the FkiEzsignfoldersignerassociationIDDeclinedtosign field if non-nil, zero value otherwise.

### GetFkiEzsignfoldersignerassociationIDDeclinedtosignOk

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) GetFkiEzsignfoldersignerassociationIDDeclinedtosignOk() (*int32, bool)`

GetFkiEzsignfoldersignerassociationIDDeclinedtosignOk returns a tuple with the FkiEzsignfoldersignerassociationIDDeclinedtosign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignfoldersignerassociationIDDeclinedtosign

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) SetFkiEzsignfoldersignerassociationIDDeclinedtosign(v int32)`

SetFkiEzsignfoldersignerassociationIDDeclinedtosign sets FkiEzsignfoldersignerassociationIDDeclinedtosign field to given value.

### HasFkiEzsignfoldersignerassociationIDDeclinedtosign

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) HasFkiEzsignfoldersignerassociationIDDeclinedtosign() bool`

HasFkiEzsignfoldersignerassociationIDDeclinedtosign returns a boolean if a field has been set.

### GetDtEzsigndocumentDuedate

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) GetDtEzsigndocumentDuedate() string`

GetDtEzsigndocumentDuedate returns the DtEzsigndocumentDuedate field if non-nil, zero value otherwise.

### GetDtEzsigndocumentDuedateOk

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) GetDtEzsigndocumentDuedateOk() (*string, bool)`

GetDtEzsigndocumentDuedateOk returns a tuple with the DtEzsigndocumentDuedate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtEzsigndocumentDuedate

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) SetDtEzsigndocumentDuedate(v string)`

SetDtEzsigndocumentDuedate sets DtEzsigndocumentDuedate field to given value.


### GetDtEzsignformCompleted

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) GetDtEzsignformCompleted() string`

GetDtEzsignformCompleted returns the DtEzsignformCompleted field if non-nil, zero value otherwise.

### GetDtEzsignformCompletedOk

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) GetDtEzsignformCompletedOk() (*string, bool)`

GetDtEzsignformCompletedOk returns a tuple with the DtEzsignformCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtEzsignformCompleted

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) SetDtEzsignformCompleted(v string)`

SetDtEzsignformCompleted sets DtEzsignformCompleted field to given value.

### HasDtEzsignformCompleted

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) HasDtEzsignformCompleted() bool`

HasDtEzsignformCompleted returns a boolean if a field has been set.

### GetFkiLanguageID

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) GetFkiLanguageID() int32`

GetFkiLanguageID returns the FkiLanguageID field if non-nil, zero value otherwise.

### GetFkiLanguageIDOk

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) GetFkiLanguageIDOk() (*int32, bool)`

GetFkiLanguageIDOk returns a tuple with the FkiLanguageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiLanguageID

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) SetFkiLanguageID(v int32)`

SetFkiLanguageID sets FkiLanguageID field to given value.

### HasFkiLanguageID

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) HasFkiLanguageID() bool`

HasFkiLanguageID returns a boolean if a field has been set.

### GetSEzsigndocumentName

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) GetSEzsigndocumentName() string`

GetSEzsigndocumentName returns the SEzsigndocumentName field if non-nil, zero value otherwise.

### GetSEzsigndocumentNameOk

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) GetSEzsigndocumentNameOk() (*string, bool)`

GetSEzsigndocumentNameOk returns a tuple with the SEzsigndocumentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigndocumentName

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) SetSEzsigndocumentName(v string)`

SetSEzsigndocumentName sets SEzsigndocumentName field to given value.


### GetEEzsigndocumentStep

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) GetEEzsigndocumentStep() FieldEEzsigndocumentStep`

GetEEzsigndocumentStep returns the EEzsigndocumentStep field if non-nil, zero value otherwise.

### GetEEzsigndocumentStepOk

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) GetEEzsigndocumentStepOk() (*FieldEEzsigndocumentStep, bool)`

GetEEzsigndocumentStepOk returns a tuple with the EEzsigndocumentStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigndocumentStep

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) SetEEzsigndocumentStep(v FieldEEzsigndocumentStep)`

SetEEzsigndocumentStep sets EEzsigndocumentStep field to given value.


### GetDtEzsigndocumentFirstsend

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) GetDtEzsigndocumentFirstsend() string`

GetDtEzsigndocumentFirstsend returns the DtEzsigndocumentFirstsend field if non-nil, zero value otherwise.

### GetDtEzsigndocumentFirstsendOk

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) GetDtEzsigndocumentFirstsendOk() (*string, bool)`

GetDtEzsigndocumentFirstsendOk returns a tuple with the DtEzsigndocumentFirstsend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtEzsigndocumentFirstsend

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) SetDtEzsigndocumentFirstsend(v string)`

SetDtEzsigndocumentFirstsend sets DtEzsigndocumentFirstsend field to given value.

### HasDtEzsigndocumentFirstsend

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) HasDtEzsigndocumentFirstsend() bool`

HasDtEzsigndocumentFirstsend returns a boolean if a field has been set.

### GetDtEzsigndocumentLastsend

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) GetDtEzsigndocumentLastsend() string`

GetDtEzsigndocumentLastsend returns the DtEzsigndocumentLastsend field if non-nil, zero value otherwise.

### GetDtEzsigndocumentLastsendOk

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) GetDtEzsigndocumentLastsendOk() (*string, bool)`

GetDtEzsigndocumentLastsendOk returns a tuple with the DtEzsigndocumentLastsend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtEzsigndocumentLastsend

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) SetDtEzsigndocumentLastsend(v string)`

SetDtEzsigndocumentLastsend sets DtEzsigndocumentLastsend field to given value.

### HasDtEzsigndocumentLastsend

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) HasDtEzsigndocumentLastsend() bool`

HasDtEzsigndocumentLastsend returns a boolean if a field has been set.

### GetIEzsigndocumentOrder

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) GetIEzsigndocumentOrder() int32`

GetIEzsigndocumentOrder returns the IEzsigndocumentOrder field if non-nil, zero value otherwise.

### GetIEzsigndocumentOrderOk

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) GetIEzsigndocumentOrderOk() (*int32, bool)`

GetIEzsigndocumentOrderOk returns a tuple with the IEzsigndocumentOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigndocumentOrder

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) SetIEzsigndocumentOrder(v int32)`

SetIEzsigndocumentOrder sets IEzsigndocumentOrder field to given value.


### GetIEzsigndocumentPagetotal

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) GetIEzsigndocumentPagetotal() int32`

GetIEzsigndocumentPagetotal returns the IEzsigndocumentPagetotal field if non-nil, zero value otherwise.

### GetIEzsigndocumentPagetotalOk

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) GetIEzsigndocumentPagetotalOk() (*int32, bool)`

GetIEzsigndocumentPagetotalOk returns a tuple with the IEzsigndocumentPagetotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigndocumentPagetotal

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) SetIEzsigndocumentPagetotal(v int32)`

SetIEzsigndocumentPagetotal sets IEzsigndocumentPagetotal field to given value.


### GetIEzsigndocumentSignaturesigned

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) GetIEzsigndocumentSignaturesigned() int32`

GetIEzsigndocumentSignaturesigned returns the IEzsigndocumentSignaturesigned field if non-nil, zero value otherwise.

### GetIEzsigndocumentSignaturesignedOk

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) GetIEzsigndocumentSignaturesignedOk() (*int32, bool)`

GetIEzsigndocumentSignaturesignedOk returns a tuple with the IEzsigndocumentSignaturesigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigndocumentSignaturesigned

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) SetIEzsigndocumentSignaturesigned(v int32)`

SetIEzsigndocumentSignaturesigned sets IEzsigndocumentSignaturesigned field to given value.


### GetIEzsigndocumentSignaturetotal

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) GetIEzsigndocumentSignaturetotal() int32`

GetIEzsigndocumentSignaturetotal returns the IEzsigndocumentSignaturetotal field if non-nil, zero value otherwise.

### GetIEzsigndocumentSignaturetotalOk

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) GetIEzsigndocumentSignaturetotalOk() (*int32, bool)`

GetIEzsigndocumentSignaturetotalOk returns a tuple with the IEzsigndocumentSignaturetotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigndocumentSignaturetotal

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) SetIEzsigndocumentSignaturetotal(v int32)`

SetIEzsigndocumentSignaturetotal sets IEzsigndocumentSignaturetotal field to given value.


### GetSEzsigndocumentMD5initial

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) GetSEzsigndocumentMD5initial() string`

GetSEzsigndocumentMD5initial returns the SEzsigndocumentMD5initial field if non-nil, zero value otherwise.

### GetSEzsigndocumentMD5initialOk

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) GetSEzsigndocumentMD5initialOk() (*string, bool)`

GetSEzsigndocumentMD5initialOk returns a tuple with the SEzsigndocumentMD5initial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigndocumentMD5initial

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) SetSEzsigndocumentMD5initial(v string)`

SetSEzsigndocumentMD5initial sets SEzsigndocumentMD5initial field to given value.

### HasSEzsigndocumentMD5initial

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) HasSEzsigndocumentMD5initial() bool`

HasSEzsigndocumentMD5initial returns a boolean if a field has been set.

### GetTEzsigndocumentDeclinedtosignreason

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) GetTEzsigndocumentDeclinedtosignreason() string`

GetTEzsigndocumentDeclinedtosignreason returns the TEzsigndocumentDeclinedtosignreason field if non-nil, zero value otherwise.

### GetTEzsigndocumentDeclinedtosignreasonOk

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) GetTEzsigndocumentDeclinedtosignreasonOk() (*string, bool)`

GetTEzsigndocumentDeclinedtosignreasonOk returns a tuple with the TEzsigndocumentDeclinedtosignreason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTEzsigndocumentDeclinedtosignreason

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) SetTEzsigndocumentDeclinedtosignreason(v string)`

SetTEzsigndocumentDeclinedtosignreason sets TEzsigndocumentDeclinedtosignreason field to given value.

### HasTEzsigndocumentDeclinedtosignreason

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) HasTEzsigndocumentDeclinedtosignreason() bool`

HasTEzsigndocumentDeclinedtosignreason returns a boolean if a field has been set.

### GetSEzsigndocumentMD5signed

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) GetSEzsigndocumentMD5signed() string`

GetSEzsigndocumentMD5signed returns the SEzsigndocumentMD5signed field if non-nil, zero value otherwise.

### GetSEzsigndocumentMD5signedOk

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) GetSEzsigndocumentMD5signedOk() (*string, bool)`

GetSEzsigndocumentMD5signedOk returns a tuple with the SEzsigndocumentMD5signed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigndocumentMD5signed

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) SetSEzsigndocumentMD5signed(v string)`

SetSEzsigndocumentMD5signed sets SEzsigndocumentMD5signed field to given value.

### HasSEzsigndocumentMD5signed

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) HasSEzsigndocumentMD5signed() bool`

HasSEzsigndocumentMD5signed returns a boolean if a field has been set.

### GetBEzsigndocumentEzsignform

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) GetBEzsigndocumentEzsignform() bool`

GetBEzsigndocumentEzsignform returns the BEzsigndocumentEzsignform field if non-nil, zero value otherwise.

### GetBEzsigndocumentEzsignformOk

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) GetBEzsigndocumentEzsignformOk() (*bool, bool)`

GetBEzsigndocumentEzsignformOk returns a tuple with the BEzsigndocumentEzsignform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsigndocumentEzsignform

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) SetBEzsigndocumentEzsignform(v bool)`

SetBEzsigndocumentEzsignform sets BEzsigndocumentEzsignform field to given value.

### HasBEzsigndocumentEzsignform

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) HasBEzsigndocumentEzsignform() bool`

HasBEzsigndocumentEzsignform returns a boolean if a field has been set.

### GetBEzsigndocumentHassignedsignatures

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) GetBEzsigndocumentHassignedsignatures() bool`

GetBEzsigndocumentHassignedsignatures returns the BEzsigndocumentHassignedsignatures field if non-nil, zero value otherwise.

### GetBEzsigndocumentHassignedsignaturesOk

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) GetBEzsigndocumentHassignedsignaturesOk() (*bool, bool)`

GetBEzsigndocumentHassignedsignaturesOk returns a tuple with the BEzsigndocumentHassignedsignatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsigndocumentHassignedsignatures

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) SetBEzsigndocumentHassignedsignatures(v bool)`

SetBEzsigndocumentHassignedsignatures sets BEzsigndocumentHassignedsignatures field to given value.

### HasBEzsigndocumentHassignedsignatures

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) HasBEzsigndocumentHassignedsignatures() bool`

HasBEzsigndocumentHassignedsignatures returns a boolean if a field has been set.

### GetObjAudit

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) GetObjAudit() CommonAudit`

GetObjAudit returns the ObjAudit field if non-nil, zero value otherwise.

### GetObjAuditOk

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) GetObjAuditOk() (*CommonAudit, bool)`

GetObjAuditOk returns a tuple with the ObjAudit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjAudit

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) SetObjAudit(v CommonAudit)`

SetObjAudit sets ObjAudit field to given value.

### HasObjAudit

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) HasObjAudit() bool`

HasObjAudit returns a boolean if a field has been set.

### GetSEzsigndocumentExternalid

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) GetSEzsigndocumentExternalid() string`

GetSEzsigndocumentExternalid returns the SEzsigndocumentExternalid field if non-nil, zero value otherwise.

### GetSEzsigndocumentExternalidOk

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) GetSEzsigndocumentExternalidOk() (*string, bool)`

GetSEzsigndocumentExternalidOk returns a tuple with the SEzsigndocumentExternalid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigndocumentExternalid

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) SetSEzsigndocumentExternalid(v string)`

SetSEzsigndocumentExternalid sets SEzsigndocumentExternalid field to given value.

### HasSEzsigndocumentExternalid

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) HasSEzsigndocumentExternalid() bool`

HasSEzsigndocumentExternalid returns a boolean if a field has been set.

### GetIEzsigndocumentEzsignsignatureattachmenttotal

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) GetIEzsigndocumentEzsignsignatureattachmenttotal() int32`

GetIEzsigndocumentEzsignsignatureattachmenttotal returns the IEzsigndocumentEzsignsignatureattachmenttotal field if non-nil, zero value otherwise.

### GetIEzsigndocumentEzsignsignatureattachmenttotalOk

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) GetIEzsigndocumentEzsignsignatureattachmenttotalOk() (*int32, bool)`

GetIEzsigndocumentEzsignsignatureattachmenttotalOk returns a tuple with the IEzsigndocumentEzsignsignatureattachmenttotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigndocumentEzsignsignatureattachmenttotal

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) SetIEzsigndocumentEzsignsignatureattachmenttotal(v int32)`

SetIEzsigndocumentEzsignsignatureattachmenttotal sets IEzsigndocumentEzsignsignatureattachmenttotal field to given value.


### GetEEzsigndocumentSteptype

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) GetEEzsigndocumentSteptype() ComputedEEzsigndocumentSteptype`

GetEEzsigndocumentSteptype returns the EEzsigndocumentSteptype field if non-nil, zero value otherwise.

### GetEEzsigndocumentSteptypeOk

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) GetEEzsigndocumentSteptypeOk() (*ComputedEEzsigndocumentSteptype, bool)`

GetEEzsigndocumentSteptypeOk returns a tuple with the EEzsigndocumentSteptype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigndocumentSteptype

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) SetEEzsigndocumentSteptype(v ComputedEEzsigndocumentSteptype)`

SetEEzsigndocumentSteptype sets EEzsigndocumentSteptype field to given value.


### GetIEzsigndocumentStepformtotal

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) GetIEzsigndocumentStepformtotal() int32`

GetIEzsigndocumentStepformtotal returns the IEzsigndocumentStepformtotal field if non-nil, zero value otherwise.

### GetIEzsigndocumentStepformtotalOk

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) GetIEzsigndocumentStepformtotalOk() (*int32, bool)`

GetIEzsigndocumentStepformtotalOk returns a tuple with the IEzsigndocumentStepformtotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigndocumentStepformtotal

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) SetIEzsigndocumentStepformtotal(v int32)`

SetIEzsigndocumentStepformtotal sets IEzsigndocumentStepformtotal field to given value.


### GetIEzsigndocumentStepformcurrent

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) GetIEzsigndocumentStepformcurrent() int32`

GetIEzsigndocumentStepformcurrent returns the IEzsigndocumentStepformcurrent field if non-nil, zero value otherwise.

### GetIEzsigndocumentStepformcurrentOk

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) GetIEzsigndocumentStepformcurrentOk() (*int32, bool)`

GetIEzsigndocumentStepformcurrentOk returns a tuple with the IEzsigndocumentStepformcurrent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigndocumentStepformcurrent

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) SetIEzsigndocumentStepformcurrent(v int32)`

SetIEzsigndocumentStepformcurrent sets IEzsigndocumentStepformcurrent field to given value.


### GetIEzsigndocumentStepsignaturetotal

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) GetIEzsigndocumentStepsignaturetotal() int32`

GetIEzsigndocumentStepsignaturetotal returns the IEzsigndocumentStepsignaturetotal field if non-nil, zero value otherwise.

### GetIEzsigndocumentStepsignaturetotalOk

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) GetIEzsigndocumentStepsignaturetotalOk() (*int32, bool)`

GetIEzsigndocumentStepsignaturetotalOk returns a tuple with the IEzsigndocumentStepsignaturetotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigndocumentStepsignaturetotal

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) SetIEzsigndocumentStepsignaturetotal(v int32)`

SetIEzsigndocumentStepsignaturetotal sets IEzsigndocumentStepsignaturetotal field to given value.


### GetIEzsigndocumentStepsignatureCurrent

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) GetIEzsigndocumentStepsignatureCurrent() int32`

GetIEzsigndocumentStepsignatureCurrent returns the IEzsigndocumentStepsignatureCurrent field if non-nil, zero value otherwise.

### GetIEzsigndocumentStepsignatureCurrentOk

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) GetIEzsigndocumentStepsignatureCurrentOk() (*int32, bool)`

GetIEzsigndocumentStepsignatureCurrentOk returns a tuple with the IEzsigndocumentStepsignatureCurrent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigndocumentStepsignatureCurrent

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) SetIEzsigndocumentStepsignatureCurrent(v int32)`

SetIEzsigndocumentStepsignatureCurrent sets IEzsigndocumentStepsignatureCurrent field to given value.


### GetAObjEzsignfoldersignerassociationstatus

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) GetAObjEzsignfoldersignerassociationstatus() []CustomEzsignfoldersignerassociationstatusResponse`

GetAObjEzsignfoldersignerassociationstatus returns the AObjEzsignfoldersignerassociationstatus field if non-nil, zero value otherwise.

### GetAObjEzsignfoldersignerassociationstatusOk

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) GetAObjEzsignfoldersignerassociationstatusOk() (*[]CustomEzsignfoldersignerassociationstatusResponse, bool)`

GetAObjEzsignfoldersignerassociationstatusOk returns a tuple with the AObjEzsignfoldersignerassociationstatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEzsignfoldersignerassociationstatus

`func (o *EzsigndocumentGetObjectV1ResponseMPayload) SetAObjEzsignfoldersignerassociationstatus(v []CustomEzsignfoldersignerassociationstatusResponse)`

SetAObjEzsignfoldersignerassociationstatus sets AObjEzsignfoldersignerassociationstatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


