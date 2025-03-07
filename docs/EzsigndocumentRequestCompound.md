# EzsigndocumentRequestCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsigndocumentID** | Pointer to **int32** | The unique ID of the Ezsigndocument | [optional] 
**FkiEzsignfolderID** | **int32** | The unique ID of the Ezsignfolder | 
**FkiEzsigntemplateID** | Pointer to **int32** | The unique ID of the Ezsigntemplate | [optional] 
**FkiEzsignfoldersignerassociationID** | Pointer to **int32** | The unique ID of the Ezsignfoldersignerassociation | [optional] 
**FkiEzsignimportdocumentID** | Pointer to **int32** | The unique ID of the Ezsignimportdocument | [optional] 
**FkiLanguageID** | **int32** | The unique ID of the Language.  Valid values:  |Value|Description| |-|-| |1|French| |2|English| | 
**EEzsigndocumentSource** | **string** | Indicates where to look for the document binary content. | 
**EEzsigndocumentFormat** | Pointer to **string** | Indicates the format of the document. | [optional] 
**SEzsigndocumentBase64** | Pointer to **string** | The Base64 encoded binary content of the document.  This field is Required when eEzsigndocumentSource &#x3D; Base64. | [optional] 
**SEzsigndocumentUrl** | Pointer to **string** | The url where the document content resides.  This field is Required when eEzsigndocumentSource &#x3D; Url. | [optional] 
**BEzsigndocumentForcerepair** | Pointer to **bool** | Try to repair the document or flatten it if it cannot be used for electronic signature.  | [optional] [default to true]
**SEzsigndocumentPassword** | Pointer to **string** | If the source document is password protected, the password to open/modify it. | [optional] 
**EEzsigndocumentForm** | Pointer to **string** | If the document contains an existing PDF form this property must be set.  **Keep** leaves the form as-is in the document.  **Convert** removes the form and convert all the existing fields to Ezsignformfieldgroups and assign them to the specified **fkiEzsignfoldersignerassociationID**  **Discard** removes the form from the document.  **Flatten** prints the form values in the document. | [optional] 
**DtEzsigndocumentDuedate** | **string** | The maximum date and time at which the Ezsigndocument can be signed. | 
**SEzsigndocumentName** | **string** | The name of the document that will be presented to Ezsignfoldersignerassociations | 
**SEzsigndocumentExternalid** | Pointer to **string** | This field can be used to store an External ID from the client&#39;s system.  Anything can be stored in this field, it will never be evaluated by the eZmax system and will be returned AS-IS.  To store multiple values, consider using a JSON formatted structure, a URL encoded string, a CSV or any other custom format.  | [optional] 

## Methods

### NewEzsigndocumentRequestCompound

`func NewEzsigndocumentRequestCompound(fkiEzsignfolderID int32, fkiLanguageID int32, eEzsigndocumentSource string, dtEzsigndocumentDuedate string, sEzsigndocumentName string, ) *EzsigndocumentRequestCompound`

NewEzsigndocumentRequestCompound instantiates a new EzsigndocumentRequestCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigndocumentRequestCompoundWithDefaults

`func NewEzsigndocumentRequestCompoundWithDefaults() *EzsigndocumentRequestCompound`

NewEzsigndocumentRequestCompoundWithDefaults instantiates a new EzsigndocumentRequestCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsigndocumentID

`func (o *EzsigndocumentRequestCompound) GetPkiEzsigndocumentID() int32`

GetPkiEzsigndocumentID returns the PkiEzsigndocumentID field if non-nil, zero value otherwise.

### GetPkiEzsigndocumentIDOk

`func (o *EzsigndocumentRequestCompound) GetPkiEzsigndocumentIDOk() (*int32, bool)`

GetPkiEzsigndocumentIDOk returns a tuple with the PkiEzsigndocumentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsigndocumentID

`func (o *EzsigndocumentRequestCompound) SetPkiEzsigndocumentID(v int32)`

SetPkiEzsigndocumentID sets PkiEzsigndocumentID field to given value.

### HasPkiEzsigndocumentID

`func (o *EzsigndocumentRequestCompound) HasPkiEzsigndocumentID() bool`

HasPkiEzsigndocumentID returns a boolean if a field has been set.

### GetFkiEzsignfolderID

`func (o *EzsigndocumentRequestCompound) GetFkiEzsignfolderID() int32`

GetFkiEzsignfolderID returns the FkiEzsignfolderID field if non-nil, zero value otherwise.

### GetFkiEzsignfolderIDOk

`func (o *EzsigndocumentRequestCompound) GetFkiEzsignfolderIDOk() (*int32, bool)`

GetFkiEzsignfolderIDOk returns a tuple with the FkiEzsignfolderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignfolderID

`func (o *EzsigndocumentRequestCompound) SetFkiEzsignfolderID(v int32)`

SetFkiEzsignfolderID sets FkiEzsignfolderID field to given value.


### GetFkiEzsigntemplateID

`func (o *EzsigndocumentRequestCompound) GetFkiEzsigntemplateID() int32`

GetFkiEzsigntemplateID returns the FkiEzsigntemplateID field if non-nil, zero value otherwise.

### GetFkiEzsigntemplateIDOk

`func (o *EzsigndocumentRequestCompound) GetFkiEzsigntemplateIDOk() (*int32, bool)`

GetFkiEzsigntemplateIDOk returns a tuple with the FkiEzsigntemplateID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigntemplateID

`func (o *EzsigndocumentRequestCompound) SetFkiEzsigntemplateID(v int32)`

SetFkiEzsigntemplateID sets FkiEzsigntemplateID field to given value.

### HasFkiEzsigntemplateID

`func (o *EzsigndocumentRequestCompound) HasFkiEzsigntemplateID() bool`

HasFkiEzsigntemplateID returns a boolean if a field has been set.

### GetFkiEzsignfoldersignerassociationID

`func (o *EzsigndocumentRequestCompound) GetFkiEzsignfoldersignerassociationID() int32`

GetFkiEzsignfoldersignerassociationID returns the FkiEzsignfoldersignerassociationID field if non-nil, zero value otherwise.

### GetFkiEzsignfoldersignerassociationIDOk

`func (o *EzsigndocumentRequestCompound) GetFkiEzsignfoldersignerassociationIDOk() (*int32, bool)`

GetFkiEzsignfoldersignerassociationIDOk returns a tuple with the FkiEzsignfoldersignerassociationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignfoldersignerassociationID

`func (o *EzsigndocumentRequestCompound) SetFkiEzsignfoldersignerassociationID(v int32)`

SetFkiEzsignfoldersignerassociationID sets FkiEzsignfoldersignerassociationID field to given value.

### HasFkiEzsignfoldersignerassociationID

`func (o *EzsigndocumentRequestCompound) HasFkiEzsignfoldersignerassociationID() bool`

HasFkiEzsignfoldersignerassociationID returns a boolean if a field has been set.

### GetFkiEzsignimportdocumentID

`func (o *EzsigndocumentRequestCompound) GetFkiEzsignimportdocumentID() int32`

GetFkiEzsignimportdocumentID returns the FkiEzsignimportdocumentID field if non-nil, zero value otherwise.

### GetFkiEzsignimportdocumentIDOk

`func (o *EzsigndocumentRequestCompound) GetFkiEzsignimportdocumentIDOk() (*int32, bool)`

GetFkiEzsignimportdocumentIDOk returns a tuple with the FkiEzsignimportdocumentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignimportdocumentID

`func (o *EzsigndocumentRequestCompound) SetFkiEzsignimportdocumentID(v int32)`

SetFkiEzsignimportdocumentID sets FkiEzsignimportdocumentID field to given value.

### HasFkiEzsignimportdocumentID

`func (o *EzsigndocumentRequestCompound) HasFkiEzsignimportdocumentID() bool`

HasFkiEzsignimportdocumentID returns a boolean if a field has been set.

### GetFkiLanguageID

`func (o *EzsigndocumentRequestCompound) GetFkiLanguageID() int32`

GetFkiLanguageID returns the FkiLanguageID field if non-nil, zero value otherwise.

### GetFkiLanguageIDOk

`func (o *EzsigndocumentRequestCompound) GetFkiLanguageIDOk() (*int32, bool)`

GetFkiLanguageIDOk returns a tuple with the FkiLanguageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiLanguageID

`func (o *EzsigndocumentRequestCompound) SetFkiLanguageID(v int32)`

SetFkiLanguageID sets FkiLanguageID field to given value.


### GetEEzsigndocumentSource

`func (o *EzsigndocumentRequestCompound) GetEEzsigndocumentSource() string`

GetEEzsigndocumentSource returns the EEzsigndocumentSource field if non-nil, zero value otherwise.

### GetEEzsigndocumentSourceOk

`func (o *EzsigndocumentRequestCompound) GetEEzsigndocumentSourceOk() (*string, bool)`

GetEEzsigndocumentSourceOk returns a tuple with the EEzsigndocumentSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigndocumentSource

`func (o *EzsigndocumentRequestCompound) SetEEzsigndocumentSource(v string)`

SetEEzsigndocumentSource sets EEzsigndocumentSource field to given value.


### GetEEzsigndocumentFormat

`func (o *EzsigndocumentRequestCompound) GetEEzsigndocumentFormat() string`

GetEEzsigndocumentFormat returns the EEzsigndocumentFormat field if non-nil, zero value otherwise.

### GetEEzsigndocumentFormatOk

`func (o *EzsigndocumentRequestCompound) GetEEzsigndocumentFormatOk() (*string, bool)`

GetEEzsigndocumentFormatOk returns a tuple with the EEzsigndocumentFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigndocumentFormat

`func (o *EzsigndocumentRequestCompound) SetEEzsigndocumentFormat(v string)`

SetEEzsigndocumentFormat sets EEzsigndocumentFormat field to given value.

### HasEEzsigndocumentFormat

`func (o *EzsigndocumentRequestCompound) HasEEzsigndocumentFormat() bool`

HasEEzsigndocumentFormat returns a boolean if a field has been set.

### GetSEzsigndocumentBase64

`func (o *EzsigndocumentRequestCompound) GetSEzsigndocumentBase64() string`

GetSEzsigndocumentBase64 returns the SEzsigndocumentBase64 field if non-nil, zero value otherwise.

### GetSEzsigndocumentBase64Ok

`func (o *EzsigndocumentRequestCompound) GetSEzsigndocumentBase64Ok() (*string, bool)`

GetSEzsigndocumentBase64Ok returns a tuple with the SEzsigndocumentBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigndocumentBase64

`func (o *EzsigndocumentRequestCompound) SetSEzsigndocumentBase64(v string)`

SetSEzsigndocumentBase64 sets SEzsigndocumentBase64 field to given value.

### HasSEzsigndocumentBase64

`func (o *EzsigndocumentRequestCompound) HasSEzsigndocumentBase64() bool`

HasSEzsigndocumentBase64 returns a boolean if a field has been set.

### GetSEzsigndocumentUrl

`func (o *EzsigndocumentRequestCompound) GetSEzsigndocumentUrl() string`

GetSEzsigndocumentUrl returns the SEzsigndocumentUrl field if non-nil, zero value otherwise.

### GetSEzsigndocumentUrlOk

`func (o *EzsigndocumentRequestCompound) GetSEzsigndocumentUrlOk() (*string, bool)`

GetSEzsigndocumentUrlOk returns a tuple with the SEzsigndocumentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigndocumentUrl

`func (o *EzsigndocumentRequestCompound) SetSEzsigndocumentUrl(v string)`

SetSEzsigndocumentUrl sets SEzsigndocumentUrl field to given value.

### HasSEzsigndocumentUrl

`func (o *EzsigndocumentRequestCompound) HasSEzsigndocumentUrl() bool`

HasSEzsigndocumentUrl returns a boolean if a field has been set.

### GetBEzsigndocumentForcerepair

`func (o *EzsigndocumentRequestCompound) GetBEzsigndocumentForcerepair() bool`

GetBEzsigndocumentForcerepair returns the BEzsigndocumentForcerepair field if non-nil, zero value otherwise.

### GetBEzsigndocumentForcerepairOk

`func (o *EzsigndocumentRequestCompound) GetBEzsigndocumentForcerepairOk() (*bool, bool)`

GetBEzsigndocumentForcerepairOk returns a tuple with the BEzsigndocumentForcerepair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsigndocumentForcerepair

`func (o *EzsigndocumentRequestCompound) SetBEzsigndocumentForcerepair(v bool)`

SetBEzsigndocumentForcerepair sets BEzsigndocumentForcerepair field to given value.

### HasBEzsigndocumentForcerepair

`func (o *EzsigndocumentRequestCompound) HasBEzsigndocumentForcerepair() bool`

HasBEzsigndocumentForcerepair returns a boolean if a field has been set.

### GetSEzsigndocumentPassword

`func (o *EzsigndocumentRequestCompound) GetSEzsigndocumentPassword() string`

GetSEzsigndocumentPassword returns the SEzsigndocumentPassword field if non-nil, zero value otherwise.

### GetSEzsigndocumentPasswordOk

`func (o *EzsigndocumentRequestCompound) GetSEzsigndocumentPasswordOk() (*string, bool)`

GetSEzsigndocumentPasswordOk returns a tuple with the SEzsigndocumentPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigndocumentPassword

`func (o *EzsigndocumentRequestCompound) SetSEzsigndocumentPassword(v string)`

SetSEzsigndocumentPassword sets SEzsigndocumentPassword field to given value.

### HasSEzsigndocumentPassword

`func (o *EzsigndocumentRequestCompound) HasSEzsigndocumentPassword() bool`

HasSEzsigndocumentPassword returns a boolean if a field has been set.

### GetEEzsigndocumentForm

`func (o *EzsigndocumentRequestCompound) GetEEzsigndocumentForm() string`

GetEEzsigndocumentForm returns the EEzsigndocumentForm field if non-nil, zero value otherwise.

### GetEEzsigndocumentFormOk

`func (o *EzsigndocumentRequestCompound) GetEEzsigndocumentFormOk() (*string, bool)`

GetEEzsigndocumentFormOk returns a tuple with the EEzsigndocumentForm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigndocumentForm

`func (o *EzsigndocumentRequestCompound) SetEEzsigndocumentForm(v string)`

SetEEzsigndocumentForm sets EEzsigndocumentForm field to given value.

### HasEEzsigndocumentForm

`func (o *EzsigndocumentRequestCompound) HasEEzsigndocumentForm() bool`

HasEEzsigndocumentForm returns a boolean if a field has been set.

### GetDtEzsigndocumentDuedate

`func (o *EzsigndocumentRequestCompound) GetDtEzsigndocumentDuedate() string`

GetDtEzsigndocumentDuedate returns the DtEzsigndocumentDuedate field if non-nil, zero value otherwise.

### GetDtEzsigndocumentDuedateOk

`func (o *EzsigndocumentRequestCompound) GetDtEzsigndocumentDuedateOk() (*string, bool)`

GetDtEzsigndocumentDuedateOk returns a tuple with the DtEzsigndocumentDuedate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtEzsigndocumentDuedate

`func (o *EzsigndocumentRequestCompound) SetDtEzsigndocumentDuedate(v string)`

SetDtEzsigndocumentDuedate sets DtEzsigndocumentDuedate field to given value.


### GetSEzsigndocumentName

`func (o *EzsigndocumentRequestCompound) GetSEzsigndocumentName() string`

GetSEzsigndocumentName returns the SEzsigndocumentName field if non-nil, zero value otherwise.

### GetSEzsigndocumentNameOk

`func (o *EzsigndocumentRequestCompound) GetSEzsigndocumentNameOk() (*string, bool)`

GetSEzsigndocumentNameOk returns a tuple with the SEzsigndocumentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigndocumentName

`func (o *EzsigndocumentRequestCompound) SetSEzsigndocumentName(v string)`

SetSEzsigndocumentName sets SEzsigndocumentName field to given value.


### GetSEzsigndocumentExternalid

`func (o *EzsigndocumentRequestCompound) GetSEzsigndocumentExternalid() string`

GetSEzsigndocumentExternalid returns the SEzsigndocumentExternalid field if non-nil, zero value otherwise.

### GetSEzsigndocumentExternalidOk

`func (o *EzsigndocumentRequestCompound) GetSEzsigndocumentExternalidOk() (*string, bool)`

GetSEzsigndocumentExternalidOk returns a tuple with the SEzsigndocumentExternalid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigndocumentExternalid

`func (o *EzsigndocumentRequestCompound) SetSEzsigndocumentExternalid(v string)`

SetSEzsigndocumentExternalid sets SEzsigndocumentExternalid field to given value.

### HasSEzsigndocumentExternalid

`func (o *EzsigndocumentRequestCompound) HasSEzsigndocumentExternalid() bool`

HasSEzsigndocumentExternalid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


