# EzsigndocumentRequestCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EEzsigndocumentSource** | **string** | Indicates where to look for the document binary content. | 
**EEzsigndocumentFormat** | **string** | Indicates the format of the document. | 
**SEzsigndocumentBase64** | Pointer to **string** | The Base64 encoded binary content of the document.  This field is Required when eEzsigndocumentSource &#x3D; Base64. | [optional] 
**SEzsigndocumentUrl** | Pointer to **string** | The url where the document content resides.  This field is Required when eEzsigndocumentSource &#x3D; Url. | [optional] 
**BEzsigndocumentForcerepair** | Pointer to **bool** | Try to repair the document or flatten it if it cannot be used for electronic signature.  | [optional] [default to true]
**SEzsigndocumentPassword** | Pointer to **string** | If the source document is password protected, the password to open/modify it. | [optional] [default to ""]
**FkiEzsignfolderID** | **int32** | The unique ID of the Ezsignfolder | 
**DtEzsigndocumentDuedate** | **string** | The maximum date and time at which the document can be signed. | 
**FkiLanguageID** | **int32** | The unique ID of the Language.  Valid values:  |Value|Description| |-|-| |1|French| |2|English| | 
**SEzsigndocumentName** | **string** | The name of the document that will be presented to Ezsignfoldersignerassociations | 

## Methods

### NewEzsigndocumentRequestCompound

`func NewEzsigndocumentRequestCompound(eEzsigndocumentSource string, eEzsigndocumentFormat string, fkiEzsignfolderID int32, dtEzsigndocumentDuedate string, fkiLanguageID int32, sEzsigndocumentName string, ) *EzsigndocumentRequestCompound`

NewEzsigndocumentRequestCompound instantiates a new EzsigndocumentRequestCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigndocumentRequestCompoundWithDefaults

`func NewEzsigndocumentRequestCompoundWithDefaults() *EzsigndocumentRequestCompound`

NewEzsigndocumentRequestCompoundWithDefaults instantiates a new EzsigndocumentRequestCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


