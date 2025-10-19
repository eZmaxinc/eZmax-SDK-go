# EzsigndocumentResponseCompoundV3

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
**IEzsigndocumentFormfieldtotal** | **int32** | The number of total Ezsignformfield that were requested in the Ezsigndocument. | 
**SEzsigndocumentMD5initial** | Pointer to **string** | MD5 Hash of the initial PDF Document before signatures were applied to it. | [optional] 
**TEzsigndocumentDeclinedtosignreason** | Pointer to **string** | A custom text message that will contain the refusal message if the Ezsigndocument is declined to sign | [optional] 
**SEzsigndocumentMD5signed** | Pointer to **string** | MD5 Hash of the final PDF Document after all signatures were applied to it. | [optional] 
**BEzsigndocumentEzsignform** | Pointer to **bool** | If the Ezsigndocument contains an Ezsignform or not | [optional] 
**BEzsigndocumentHassignedsignatures** | Pointer to **bool** | If the Ezsigndocument contains signed signatures (From internal or external sources) | [optional] 
**BEzsigndocumentSendtoged** | Pointer to **bool** | Whether the Ezsigndocument was copied to EDM | [optional] 
**ObjAudit** | Pointer to [**CommonAudit**](CommonAudit.md) |  | [optional] 
**SEzsigndocumentExternalid** | Pointer to **string** | This field can be used to store an External ID from the client&#39;s system.  Anything can be stored in this field, it will never be evaluated by the eZmax system and will be returned AS-IS.  To store multiple values, consider using a JSON formatted structure, a URL encoded string, a CSV or any other custom format.  | [optional] 
**IEzsigndocumentEzsignsignatureattachmenttotal** | **int32** | The number of Ezsigndocumentattachment total | 
**IEzsigndocumentEzsigndiscussiontotal** | **int32** | The total number of Ezsigndiscussions | 
**IEzsigndocumentSteptotal** | **int32** | The total number of steps | 
**IEzsigndocumentStepcurrent** | **int32** | The current step | 
**AObjEzsignfoldersignerassociationstatus** | [**[]CustomEzsignfoldersignerassociationstatusResponseV3**](CustomEzsignfoldersignerassociationstatusResponseV3.md) |  | 
**AObjEzsigndocumentdependency** | Pointer to [**[]EzsigndocumentdependencyResponse**](EzsigndocumentdependencyResponse.md) |  | [optional] 

## Methods

### NewEzsigndocumentResponseCompoundV3

`func NewEzsigndocumentResponseCompoundV3(pkiEzsigndocumentID int32, fkiEzsignfolderID int32, dtEzsigndocumentDuedate string, sEzsigndocumentName string, eEzsigndocumentStep FieldEEzsigndocumentStep, iEzsigndocumentOrder int32, iEzsigndocumentPagetotal int32, iEzsigndocumentSignaturesigned int32, iEzsigndocumentSignaturetotal int32, iEzsigndocumentFormfieldtotal int32, iEzsigndocumentEzsignsignatureattachmenttotal int32, iEzsigndocumentEzsigndiscussiontotal int32, iEzsigndocumentSteptotal int32, iEzsigndocumentStepcurrent int32, aObjEzsignfoldersignerassociationstatus []CustomEzsignfoldersignerassociationstatusResponseV3, ) *EzsigndocumentResponseCompoundV3`

NewEzsigndocumentResponseCompoundV3 instantiates a new EzsigndocumentResponseCompoundV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigndocumentResponseCompoundV3WithDefaults

`func NewEzsigndocumentResponseCompoundV3WithDefaults() *EzsigndocumentResponseCompoundV3`

NewEzsigndocumentResponseCompoundV3WithDefaults instantiates a new EzsigndocumentResponseCompoundV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsigndocumentID

`func (o *EzsigndocumentResponseCompoundV3) GetPkiEzsigndocumentID() int32`

GetPkiEzsigndocumentID returns the PkiEzsigndocumentID field if non-nil, zero value otherwise.

### GetPkiEzsigndocumentIDOk

`func (o *EzsigndocumentResponseCompoundV3) GetPkiEzsigndocumentIDOk() (*int32, bool)`

GetPkiEzsigndocumentIDOk returns a tuple with the PkiEzsigndocumentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsigndocumentID

`func (o *EzsigndocumentResponseCompoundV3) SetPkiEzsigndocumentID(v int32)`

SetPkiEzsigndocumentID sets PkiEzsigndocumentID field to given value.


### GetFkiEzsignfolderID

`func (o *EzsigndocumentResponseCompoundV3) GetFkiEzsignfolderID() int32`

GetFkiEzsignfolderID returns the FkiEzsignfolderID field if non-nil, zero value otherwise.

### GetFkiEzsignfolderIDOk

`func (o *EzsigndocumentResponseCompoundV3) GetFkiEzsignfolderIDOk() (*int32, bool)`

GetFkiEzsignfolderIDOk returns a tuple with the FkiEzsignfolderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignfolderID

`func (o *EzsigndocumentResponseCompoundV3) SetFkiEzsignfolderID(v int32)`

SetFkiEzsignfolderID sets FkiEzsignfolderID field to given value.


### GetFkiEzsignfoldersignerassociationIDDeclinedtosign

`func (o *EzsigndocumentResponseCompoundV3) GetFkiEzsignfoldersignerassociationIDDeclinedtosign() int32`

GetFkiEzsignfoldersignerassociationIDDeclinedtosign returns the FkiEzsignfoldersignerassociationIDDeclinedtosign field if non-nil, zero value otherwise.

### GetFkiEzsignfoldersignerassociationIDDeclinedtosignOk

`func (o *EzsigndocumentResponseCompoundV3) GetFkiEzsignfoldersignerassociationIDDeclinedtosignOk() (*int32, bool)`

GetFkiEzsignfoldersignerassociationIDDeclinedtosignOk returns a tuple with the FkiEzsignfoldersignerassociationIDDeclinedtosign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignfoldersignerassociationIDDeclinedtosign

`func (o *EzsigndocumentResponseCompoundV3) SetFkiEzsignfoldersignerassociationIDDeclinedtosign(v int32)`

SetFkiEzsignfoldersignerassociationIDDeclinedtosign sets FkiEzsignfoldersignerassociationIDDeclinedtosign field to given value.

### HasFkiEzsignfoldersignerassociationIDDeclinedtosign

`func (o *EzsigndocumentResponseCompoundV3) HasFkiEzsignfoldersignerassociationIDDeclinedtosign() bool`

HasFkiEzsignfoldersignerassociationIDDeclinedtosign returns a boolean if a field has been set.

### GetDtEzsigndocumentDuedate

`func (o *EzsigndocumentResponseCompoundV3) GetDtEzsigndocumentDuedate() string`

GetDtEzsigndocumentDuedate returns the DtEzsigndocumentDuedate field if non-nil, zero value otherwise.

### GetDtEzsigndocumentDuedateOk

`func (o *EzsigndocumentResponseCompoundV3) GetDtEzsigndocumentDuedateOk() (*string, bool)`

GetDtEzsigndocumentDuedateOk returns a tuple with the DtEzsigndocumentDuedate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtEzsigndocumentDuedate

`func (o *EzsigndocumentResponseCompoundV3) SetDtEzsigndocumentDuedate(v string)`

SetDtEzsigndocumentDuedate sets DtEzsigndocumentDuedate field to given value.


### GetDtEzsignformCompleted

`func (o *EzsigndocumentResponseCompoundV3) GetDtEzsignformCompleted() string`

GetDtEzsignformCompleted returns the DtEzsignformCompleted field if non-nil, zero value otherwise.

### GetDtEzsignformCompletedOk

`func (o *EzsigndocumentResponseCompoundV3) GetDtEzsignformCompletedOk() (*string, bool)`

GetDtEzsignformCompletedOk returns a tuple with the DtEzsignformCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtEzsignformCompleted

`func (o *EzsigndocumentResponseCompoundV3) SetDtEzsignformCompleted(v string)`

SetDtEzsignformCompleted sets DtEzsignformCompleted field to given value.

### HasDtEzsignformCompleted

`func (o *EzsigndocumentResponseCompoundV3) HasDtEzsignformCompleted() bool`

HasDtEzsignformCompleted returns a boolean if a field has been set.

### GetFkiLanguageID

`func (o *EzsigndocumentResponseCompoundV3) GetFkiLanguageID() int32`

GetFkiLanguageID returns the FkiLanguageID field if non-nil, zero value otherwise.

### GetFkiLanguageIDOk

`func (o *EzsigndocumentResponseCompoundV3) GetFkiLanguageIDOk() (*int32, bool)`

GetFkiLanguageIDOk returns a tuple with the FkiLanguageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiLanguageID

`func (o *EzsigndocumentResponseCompoundV3) SetFkiLanguageID(v int32)`

SetFkiLanguageID sets FkiLanguageID field to given value.

### HasFkiLanguageID

`func (o *EzsigndocumentResponseCompoundV3) HasFkiLanguageID() bool`

HasFkiLanguageID returns a boolean if a field has been set.

### GetSEzsigndocumentName

`func (o *EzsigndocumentResponseCompoundV3) GetSEzsigndocumentName() string`

GetSEzsigndocumentName returns the SEzsigndocumentName field if non-nil, zero value otherwise.

### GetSEzsigndocumentNameOk

`func (o *EzsigndocumentResponseCompoundV3) GetSEzsigndocumentNameOk() (*string, bool)`

GetSEzsigndocumentNameOk returns a tuple with the SEzsigndocumentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigndocumentName

`func (o *EzsigndocumentResponseCompoundV3) SetSEzsigndocumentName(v string)`

SetSEzsigndocumentName sets SEzsigndocumentName field to given value.


### GetEEzsigndocumentStep

`func (o *EzsigndocumentResponseCompoundV3) GetEEzsigndocumentStep() FieldEEzsigndocumentStep`

GetEEzsigndocumentStep returns the EEzsigndocumentStep field if non-nil, zero value otherwise.

### GetEEzsigndocumentStepOk

`func (o *EzsigndocumentResponseCompoundV3) GetEEzsigndocumentStepOk() (*FieldEEzsigndocumentStep, bool)`

GetEEzsigndocumentStepOk returns a tuple with the EEzsigndocumentStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigndocumentStep

`func (o *EzsigndocumentResponseCompoundV3) SetEEzsigndocumentStep(v FieldEEzsigndocumentStep)`

SetEEzsigndocumentStep sets EEzsigndocumentStep field to given value.


### GetDtEzsigndocumentFirstsend

`func (o *EzsigndocumentResponseCompoundV3) GetDtEzsigndocumentFirstsend() string`

GetDtEzsigndocumentFirstsend returns the DtEzsigndocumentFirstsend field if non-nil, zero value otherwise.

### GetDtEzsigndocumentFirstsendOk

`func (o *EzsigndocumentResponseCompoundV3) GetDtEzsigndocumentFirstsendOk() (*string, bool)`

GetDtEzsigndocumentFirstsendOk returns a tuple with the DtEzsigndocumentFirstsend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtEzsigndocumentFirstsend

`func (o *EzsigndocumentResponseCompoundV3) SetDtEzsigndocumentFirstsend(v string)`

SetDtEzsigndocumentFirstsend sets DtEzsigndocumentFirstsend field to given value.

### HasDtEzsigndocumentFirstsend

`func (o *EzsigndocumentResponseCompoundV3) HasDtEzsigndocumentFirstsend() bool`

HasDtEzsigndocumentFirstsend returns a boolean if a field has been set.

### GetDtEzsigndocumentLastsend

`func (o *EzsigndocumentResponseCompoundV3) GetDtEzsigndocumentLastsend() string`

GetDtEzsigndocumentLastsend returns the DtEzsigndocumentLastsend field if non-nil, zero value otherwise.

### GetDtEzsigndocumentLastsendOk

`func (o *EzsigndocumentResponseCompoundV3) GetDtEzsigndocumentLastsendOk() (*string, bool)`

GetDtEzsigndocumentLastsendOk returns a tuple with the DtEzsigndocumentLastsend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtEzsigndocumentLastsend

`func (o *EzsigndocumentResponseCompoundV3) SetDtEzsigndocumentLastsend(v string)`

SetDtEzsigndocumentLastsend sets DtEzsigndocumentLastsend field to given value.

### HasDtEzsigndocumentLastsend

`func (o *EzsigndocumentResponseCompoundV3) HasDtEzsigndocumentLastsend() bool`

HasDtEzsigndocumentLastsend returns a boolean if a field has been set.

### GetIEzsigndocumentOrder

`func (o *EzsigndocumentResponseCompoundV3) GetIEzsigndocumentOrder() int32`

GetIEzsigndocumentOrder returns the IEzsigndocumentOrder field if non-nil, zero value otherwise.

### GetIEzsigndocumentOrderOk

`func (o *EzsigndocumentResponseCompoundV3) GetIEzsigndocumentOrderOk() (*int32, bool)`

GetIEzsigndocumentOrderOk returns a tuple with the IEzsigndocumentOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigndocumentOrder

`func (o *EzsigndocumentResponseCompoundV3) SetIEzsigndocumentOrder(v int32)`

SetIEzsigndocumentOrder sets IEzsigndocumentOrder field to given value.


### GetIEzsigndocumentPagetotal

`func (o *EzsigndocumentResponseCompoundV3) GetIEzsigndocumentPagetotal() int32`

GetIEzsigndocumentPagetotal returns the IEzsigndocumentPagetotal field if non-nil, zero value otherwise.

### GetIEzsigndocumentPagetotalOk

`func (o *EzsigndocumentResponseCompoundV3) GetIEzsigndocumentPagetotalOk() (*int32, bool)`

GetIEzsigndocumentPagetotalOk returns a tuple with the IEzsigndocumentPagetotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigndocumentPagetotal

`func (o *EzsigndocumentResponseCompoundV3) SetIEzsigndocumentPagetotal(v int32)`

SetIEzsigndocumentPagetotal sets IEzsigndocumentPagetotal field to given value.


### GetIEzsigndocumentSignaturesigned

`func (o *EzsigndocumentResponseCompoundV3) GetIEzsigndocumentSignaturesigned() int32`

GetIEzsigndocumentSignaturesigned returns the IEzsigndocumentSignaturesigned field if non-nil, zero value otherwise.

### GetIEzsigndocumentSignaturesignedOk

`func (o *EzsigndocumentResponseCompoundV3) GetIEzsigndocumentSignaturesignedOk() (*int32, bool)`

GetIEzsigndocumentSignaturesignedOk returns a tuple with the IEzsigndocumentSignaturesigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigndocumentSignaturesigned

`func (o *EzsigndocumentResponseCompoundV3) SetIEzsigndocumentSignaturesigned(v int32)`

SetIEzsigndocumentSignaturesigned sets IEzsigndocumentSignaturesigned field to given value.


### GetIEzsigndocumentSignaturetotal

`func (o *EzsigndocumentResponseCompoundV3) GetIEzsigndocumentSignaturetotal() int32`

GetIEzsigndocumentSignaturetotal returns the IEzsigndocumentSignaturetotal field if non-nil, zero value otherwise.

### GetIEzsigndocumentSignaturetotalOk

`func (o *EzsigndocumentResponseCompoundV3) GetIEzsigndocumentSignaturetotalOk() (*int32, bool)`

GetIEzsigndocumentSignaturetotalOk returns a tuple with the IEzsigndocumentSignaturetotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigndocumentSignaturetotal

`func (o *EzsigndocumentResponseCompoundV3) SetIEzsigndocumentSignaturetotal(v int32)`

SetIEzsigndocumentSignaturetotal sets IEzsigndocumentSignaturetotal field to given value.


### GetIEzsigndocumentFormfieldtotal

`func (o *EzsigndocumentResponseCompoundV3) GetIEzsigndocumentFormfieldtotal() int32`

GetIEzsigndocumentFormfieldtotal returns the IEzsigndocumentFormfieldtotal field if non-nil, zero value otherwise.

### GetIEzsigndocumentFormfieldtotalOk

`func (o *EzsigndocumentResponseCompoundV3) GetIEzsigndocumentFormfieldtotalOk() (*int32, bool)`

GetIEzsigndocumentFormfieldtotalOk returns a tuple with the IEzsigndocumentFormfieldtotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigndocumentFormfieldtotal

`func (o *EzsigndocumentResponseCompoundV3) SetIEzsigndocumentFormfieldtotal(v int32)`

SetIEzsigndocumentFormfieldtotal sets IEzsigndocumentFormfieldtotal field to given value.


### GetSEzsigndocumentMD5initial

`func (o *EzsigndocumentResponseCompoundV3) GetSEzsigndocumentMD5initial() string`

GetSEzsigndocumentMD5initial returns the SEzsigndocumentMD5initial field if non-nil, zero value otherwise.

### GetSEzsigndocumentMD5initialOk

`func (o *EzsigndocumentResponseCompoundV3) GetSEzsigndocumentMD5initialOk() (*string, bool)`

GetSEzsigndocumentMD5initialOk returns a tuple with the SEzsigndocumentMD5initial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigndocumentMD5initial

`func (o *EzsigndocumentResponseCompoundV3) SetSEzsigndocumentMD5initial(v string)`

SetSEzsigndocumentMD5initial sets SEzsigndocumentMD5initial field to given value.

### HasSEzsigndocumentMD5initial

`func (o *EzsigndocumentResponseCompoundV3) HasSEzsigndocumentMD5initial() bool`

HasSEzsigndocumentMD5initial returns a boolean if a field has been set.

### GetTEzsigndocumentDeclinedtosignreason

`func (o *EzsigndocumentResponseCompoundV3) GetTEzsigndocumentDeclinedtosignreason() string`

GetTEzsigndocumentDeclinedtosignreason returns the TEzsigndocumentDeclinedtosignreason field if non-nil, zero value otherwise.

### GetTEzsigndocumentDeclinedtosignreasonOk

`func (o *EzsigndocumentResponseCompoundV3) GetTEzsigndocumentDeclinedtosignreasonOk() (*string, bool)`

GetTEzsigndocumentDeclinedtosignreasonOk returns a tuple with the TEzsigndocumentDeclinedtosignreason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTEzsigndocumentDeclinedtosignreason

`func (o *EzsigndocumentResponseCompoundV3) SetTEzsigndocumentDeclinedtosignreason(v string)`

SetTEzsigndocumentDeclinedtosignreason sets TEzsigndocumentDeclinedtosignreason field to given value.

### HasTEzsigndocumentDeclinedtosignreason

`func (o *EzsigndocumentResponseCompoundV3) HasTEzsigndocumentDeclinedtosignreason() bool`

HasTEzsigndocumentDeclinedtosignreason returns a boolean if a field has been set.

### GetSEzsigndocumentMD5signed

`func (o *EzsigndocumentResponseCompoundV3) GetSEzsigndocumentMD5signed() string`

GetSEzsigndocumentMD5signed returns the SEzsigndocumentMD5signed field if non-nil, zero value otherwise.

### GetSEzsigndocumentMD5signedOk

`func (o *EzsigndocumentResponseCompoundV3) GetSEzsigndocumentMD5signedOk() (*string, bool)`

GetSEzsigndocumentMD5signedOk returns a tuple with the SEzsigndocumentMD5signed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigndocumentMD5signed

`func (o *EzsigndocumentResponseCompoundV3) SetSEzsigndocumentMD5signed(v string)`

SetSEzsigndocumentMD5signed sets SEzsigndocumentMD5signed field to given value.

### HasSEzsigndocumentMD5signed

`func (o *EzsigndocumentResponseCompoundV3) HasSEzsigndocumentMD5signed() bool`

HasSEzsigndocumentMD5signed returns a boolean if a field has been set.

### GetBEzsigndocumentEzsignform

`func (o *EzsigndocumentResponseCompoundV3) GetBEzsigndocumentEzsignform() bool`

GetBEzsigndocumentEzsignform returns the BEzsigndocumentEzsignform field if non-nil, zero value otherwise.

### GetBEzsigndocumentEzsignformOk

`func (o *EzsigndocumentResponseCompoundV3) GetBEzsigndocumentEzsignformOk() (*bool, bool)`

GetBEzsigndocumentEzsignformOk returns a tuple with the BEzsigndocumentEzsignform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsigndocumentEzsignform

`func (o *EzsigndocumentResponseCompoundV3) SetBEzsigndocumentEzsignform(v bool)`

SetBEzsigndocumentEzsignform sets BEzsigndocumentEzsignform field to given value.

### HasBEzsigndocumentEzsignform

`func (o *EzsigndocumentResponseCompoundV3) HasBEzsigndocumentEzsignform() bool`

HasBEzsigndocumentEzsignform returns a boolean if a field has been set.

### GetBEzsigndocumentHassignedsignatures

`func (o *EzsigndocumentResponseCompoundV3) GetBEzsigndocumentHassignedsignatures() bool`

GetBEzsigndocumentHassignedsignatures returns the BEzsigndocumentHassignedsignatures field if non-nil, zero value otherwise.

### GetBEzsigndocumentHassignedsignaturesOk

`func (o *EzsigndocumentResponseCompoundV3) GetBEzsigndocumentHassignedsignaturesOk() (*bool, bool)`

GetBEzsigndocumentHassignedsignaturesOk returns a tuple with the BEzsigndocumentHassignedsignatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsigndocumentHassignedsignatures

`func (o *EzsigndocumentResponseCompoundV3) SetBEzsigndocumentHassignedsignatures(v bool)`

SetBEzsigndocumentHassignedsignatures sets BEzsigndocumentHassignedsignatures field to given value.

### HasBEzsigndocumentHassignedsignatures

`func (o *EzsigndocumentResponseCompoundV3) HasBEzsigndocumentHassignedsignatures() bool`

HasBEzsigndocumentHassignedsignatures returns a boolean if a field has been set.

### GetBEzsigndocumentSendtoged

`func (o *EzsigndocumentResponseCompoundV3) GetBEzsigndocumentSendtoged() bool`

GetBEzsigndocumentSendtoged returns the BEzsigndocumentSendtoged field if non-nil, zero value otherwise.

### GetBEzsigndocumentSendtogedOk

`func (o *EzsigndocumentResponseCompoundV3) GetBEzsigndocumentSendtogedOk() (*bool, bool)`

GetBEzsigndocumentSendtogedOk returns a tuple with the BEzsigndocumentSendtoged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsigndocumentSendtoged

`func (o *EzsigndocumentResponseCompoundV3) SetBEzsigndocumentSendtoged(v bool)`

SetBEzsigndocumentSendtoged sets BEzsigndocumentSendtoged field to given value.

### HasBEzsigndocumentSendtoged

`func (o *EzsigndocumentResponseCompoundV3) HasBEzsigndocumentSendtoged() bool`

HasBEzsigndocumentSendtoged returns a boolean if a field has been set.

### GetObjAudit

`func (o *EzsigndocumentResponseCompoundV3) GetObjAudit() CommonAudit`

GetObjAudit returns the ObjAudit field if non-nil, zero value otherwise.

### GetObjAuditOk

`func (o *EzsigndocumentResponseCompoundV3) GetObjAuditOk() (*CommonAudit, bool)`

GetObjAuditOk returns a tuple with the ObjAudit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjAudit

`func (o *EzsigndocumentResponseCompoundV3) SetObjAudit(v CommonAudit)`

SetObjAudit sets ObjAudit field to given value.

### HasObjAudit

`func (o *EzsigndocumentResponseCompoundV3) HasObjAudit() bool`

HasObjAudit returns a boolean if a field has been set.

### GetSEzsigndocumentExternalid

`func (o *EzsigndocumentResponseCompoundV3) GetSEzsigndocumentExternalid() string`

GetSEzsigndocumentExternalid returns the SEzsigndocumentExternalid field if non-nil, zero value otherwise.

### GetSEzsigndocumentExternalidOk

`func (o *EzsigndocumentResponseCompoundV3) GetSEzsigndocumentExternalidOk() (*string, bool)`

GetSEzsigndocumentExternalidOk returns a tuple with the SEzsigndocumentExternalid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigndocumentExternalid

`func (o *EzsigndocumentResponseCompoundV3) SetSEzsigndocumentExternalid(v string)`

SetSEzsigndocumentExternalid sets SEzsigndocumentExternalid field to given value.

### HasSEzsigndocumentExternalid

`func (o *EzsigndocumentResponseCompoundV3) HasSEzsigndocumentExternalid() bool`

HasSEzsigndocumentExternalid returns a boolean if a field has been set.

### GetIEzsigndocumentEzsignsignatureattachmenttotal

`func (o *EzsigndocumentResponseCompoundV3) GetIEzsigndocumentEzsignsignatureattachmenttotal() int32`

GetIEzsigndocumentEzsignsignatureattachmenttotal returns the IEzsigndocumentEzsignsignatureattachmenttotal field if non-nil, zero value otherwise.

### GetIEzsigndocumentEzsignsignatureattachmenttotalOk

`func (o *EzsigndocumentResponseCompoundV3) GetIEzsigndocumentEzsignsignatureattachmenttotalOk() (*int32, bool)`

GetIEzsigndocumentEzsignsignatureattachmenttotalOk returns a tuple with the IEzsigndocumentEzsignsignatureattachmenttotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigndocumentEzsignsignatureattachmenttotal

`func (o *EzsigndocumentResponseCompoundV3) SetIEzsigndocumentEzsignsignatureattachmenttotal(v int32)`

SetIEzsigndocumentEzsignsignatureattachmenttotal sets IEzsigndocumentEzsignsignatureattachmenttotal field to given value.


### GetIEzsigndocumentEzsigndiscussiontotal

`func (o *EzsigndocumentResponseCompoundV3) GetIEzsigndocumentEzsigndiscussiontotal() int32`

GetIEzsigndocumentEzsigndiscussiontotal returns the IEzsigndocumentEzsigndiscussiontotal field if non-nil, zero value otherwise.

### GetIEzsigndocumentEzsigndiscussiontotalOk

`func (o *EzsigndocumentResponseCompoundV3) GetIEzsigndocumentEzsigndiscussiontotalOk() (*int32, bool)`

GetIEzsigndocumentEzsigndiscussiontotalOk returns a tuple with the IEzsigndocumentEzsigndiscussiontotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigndocumentEzsigndiscussiontotal

`func (o *EzsigndocumentResponseCompoundV3) SetIEzsigndocumentEzsigndiscussiontotal(v int32)`

SetIEzsigndocumentEzsigndiscussiontotal sets IEzsigndocumentEzsigndiscussiontotal field to given value.


### GetIEzsigndocumentSteptotal

`func (o *EzsigndocumentResponseCompoundV3) GetIEzsigndocumentSteptotal() int32`

GetIEzsigndocumentSteptotal returns the IEzsigndocumentSteptotal field if non-nil, zero value otherwise.

### GetIEzsigndocumentSteptotalOk

`func (o *EzsigndocumentResponseCompoundV3) GetIEzsigndocumentSteptotalOk() (*int32, bool)`

GetIEzsigndocumentSteptotalOk returns a tuple with the IEzsigndocumentSteptotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigndocumentSteptotal

`func (o *EzsigndocumentResponseCompoundV3) SetIEzsigndocumentSteptotal(v int32)`

SetIEzsigndocumentSteptotal sets IEzsigndocumentSteptotal field to given value.


### GetIEzsigndocumentStepcurrent

`func (o *EzsigndocumentResponseCompoundV3) GetIEzsigndocumentStepcurrent() int32`

GetIEzsigndocumentStepcurrent returns the IEzsigndocumentStepcurrent field if non-nil, zero value otherwise.

### GetIEzsigndocumentStepcurrentOk

`func (o *EzsigndocumentResponseCompoundV3) GetIEzsigndocumentStepcurrentOk() (*int32, bool)`

GetIEzsigndocumentStepcurrentOk returns a tuple with the IEzsigndocumentStepcurrent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigndocumentStepcurrent

`func (o *EzsigndocumentResponseCompoundV3) SetIEzsigndocumentStepcurrent(v int32)`

SetIEzsigndocumentStepcurrent sets IEzsigndocumentStepcurrent field to given value.


### GetAObjEzsignfoldersignerassociationstatus

`func (o *EzsigndocumentResponseCompoundV3) GetAObjEzsignfoldersignerassociationstatus() []CustomEzsignfoldersignerassociationstatusResponseV3`

GetAObjEzsignfoldersignerassociationstatus returns the AObjEzsignfoldersignerassociationstatus field if non-nil, zero value otherwise.

### GetAObjEzsignfoldersignerassociationstatusOk

`func (o *EzsigndocumentResponseCompoundV3) GetAObjEzsignfoldersignerassociationstatusOk() (*[]CustomEzsignfoldersignerassociationstatusResponseV3, bool)`

GetAObjEzsignfoldersignerassociationstatusOk returns a tuple with the AObjEzsignfoldersignerassociationstatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEzsignfoldersignerassociationstatus

`func (o *EzsigndocumentResponseCompoundV3) SetAObjEzsignfoldersignerassociationstatus(v []CustomEzsignfoldersignerassociationstatusResponseV3)`

SetAObjEzsignfoldersignerassociationstatus sets AObjEzsignfoldersignerassociationstatus field to given value.


### GetAObjEzsigndocumentdependency

`func (o *EzsigndocumentResponseCompoundV3) GetAObjEzsigndocumentdependency() []EzsigndocumentdependencyResponse`

GetAObjEzsigndocumentdependency returns the AObjEzsigndocumentdependency field if non-nil, zero value otherwise.

### GetAObjEzsigndocumentdependencyOk

`func (o *EzsigndocumentResponseCompoundV3) GetAObjEzsigndocumentdependencyOk() (*[]EzsigndocumentdependencyResponse, bool)`

GetAObjEzsigndocumentdependencyOk returns a tuple with the AObjEzsigndocumentdependency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEzsigndocumentdependency

`func (o *EzsigndocumentResponseCompoundV3) SetAObjEzsigndocumentdependency(v []EzsigndocumentdependencyResponse)`

SetAObjEzsigndocumentdependency sets AObjEzsigndocumentdependency field to given value.

### HasAObjEzsigndocumentdependency

`func (o *EzsigndocumentResponseCompoundV3) HasAObjEzsigndocumentdependency() bool`

HasAObjEzsigndocumentdependency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


