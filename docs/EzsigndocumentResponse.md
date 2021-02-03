# EzsigndocumentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FkiEzsignfolderID** | **int32** | The unique ID of the Ezsignfolder | 
**DtEzsigndocumentDuedate** | **string** | The maximum date and time at which the document can be signed. | 
**FkiLanguageID** | **int32** | The unique ID of the Language.  Valid values:  |Value|Description| |-|-| |1|French| |2|English| | 
**SEzsigndocumentFilename** | **string** | The actual file name that will be used when downloading or attaching to an email. | 
**SEzsigndocumentName** | **string** | The name of the document that will be presented to Ezsignfoldersignerassociations | 
**PkiEzsigndocumentID** | **int32** | The unique ID of the Ezsigntemplate | 
**EEzsigndocumentStep** | [**FieldEEzsigndocumentStep**](Field-eEzsigndocumentStep.md) |  | 
**DtEzsigndocumentFirstsend** | **string** | The date and time when the Ezsigndocument was first sent. | 
**DtEzsigndocumentLastsend** | **string** | The date and time when the Ezsigndocument was sent the last time. | 
**IEzsigndocumentOrder** | **int32** | The order in which the Ezsigndocument will be presented to the signatory in the Ezsignfolder. | 
**IEzsigndocumentPagetotal** | **int32** | The number of pages in the Ezsigndocument. | 
**IEzsigndocumentSignaturesigned** | **int32** | The number of signatures that were signed in the document. | 
**IEzsigndocumentSignaturetotal** | **int32** | The number of total signatures that were requested in the Ezsigndocument. | 
**SEzsigndocumentMD5initial** | **string** | MD5 Hash of the initial PDF Document before signatures were applied to it. | 
**SEzsigndocumentMD5signed** | **string** | MD5 Hash of the final PDF Document after all signatures were applied to it. | 
**ObjAudit** | [**CommonAudit**](Common-Audit.md) |  | 

## Methods

### NewEzsigndocumentResponse

`func NewEzsigndocumentResponse(fkiEzsignfolderID int32, dtEzsigndocumentDuedate string, fkiLanguageID int32, sEzsigndocumentFilename string, sEzsigndocumentName string, pkiEzsigndocumentID int32, eEzsigndocumentStep FieldEEzsigndocumentStep, dtEzsigndocumentFirstsend string, dtEzsigndocumentLastsend string, iEzsigndocumentOrder int32, iEzsigndocumentPagetotal int32, iEzsigndocumentSignaturesigned int32, iEzsigndocumentSignaturetotal int32, sEzsigndocumentMD5initial string, sEzsigndocumentMD5signed string, objAudit CommonAudit, ) *EzsigndocumentResponse`

NewEzsigndocumentResponse instantiates a new EzsigndocumentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigndocumentResponseWithDefaults

`func NewEzsigndocumentResponseWithDefaults() *EzsigndocumentResponse`

NewEzsigndocumentResponseWithDefaults instantiates a new EzsigndocumentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


### GetSEzsigndocumentFilename

`func (o *EzsigndocumentResponse) GetSEzsigndocumentFilename() string`

GetSEzsigndocumentFilename returns the SEzsigndocumentFilename field if non-nil, zero value otherwise.

### GetSEzsigndocumentFilenameOk

`func (o *EzsigndocumentResponse) GetSEzsigndocumentFilenameOk() (*string, bool)`

GetSEzsigndocumentFilenameOk returns a tuple with the SEzsigndocumentFilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigndocumentFilename

`func (o *EzsigndocumentResponse) SetSEzsigndocumentFilename(v string)`

SetSEzsigndocumentFilename sets SEzsigndocumentFilename field to given value.


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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


