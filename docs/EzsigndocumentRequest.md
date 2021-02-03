# EzsigndocumentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EEzsigndocumentSource** | **string** | Indicates where to look for the document binary content. | 
**EEzsigndocumentFormat** | **string** | Indicates the format of the document. | 
**SEzsigndocumentBase64** | Pointer to **string** | The Base64 encoded binary content of the document.  This field is Required when eEzsigndocumentSource &#x3D; Base64. | [optional] 
**FkiEzsignfolderID** | **int32** | A reference to a valid Ezsignfolder.  That value is returned after a successful Ezsignfolder Creation. | 
**DtEzsigndocumentDuedate** | **string** | Represent a Date Time. The timezone is the one configured in the User&#39;s profile. | 
**FkiLanguageID** | **int32** | The unique ID of the Language.  Valid values:  |Value|Description| |-|-| |1|French| |2|English| | 
**SEzsigndocumentFilename** | **string** | The actual file name that will be used when downloading or attaching to an email. | 
**SEzsigndocumentName** | **string** | The name of the document that will be presented to Ezsignfoldersignerassociations | 

## Methods

### NewEzsigndocumentRequest

`func NewEzsigndocumentRequest(eEzsigndocumentSource string, eEzsigndocumentFormat string, fkiEzsignfolderID int32, dtEzsigndocumentDuedate string, fkiLanguageID int32, sEzsigndocumentFilename string, sEzsigndocumentName string, ) *EzsigndocumentRequest`

NewEzsigndocumentRequest instantiates a new EzsigndocumentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigndocumentRequestWithDefaults

`func NewEzsigndocumentRequestWithDefaults() *EzsigndocumentRequest`

NewEzsigndocumentRequestWithDefaults instantiates a new EzsigndocumentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEEzsigndocumentSource

`func (o *EzsigndocumentRequest) GetEEzsigndocumentSource() string`

GetEEzsigndocumentSource returns the EEzsigndocumentSource field if non-nil, zero value otherwise.

### GetEEzsigndocumentSourceOk

`func (o *EzsigndocumentRequest) GetEEzsigndocumentSourceOk() (*string, bool)`

GetEEzsigndocumentSourceOk returns a tuple with the EEzsigndocumentSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigndocumentSource

`func (o *EzsigndocumentRequest) SetEEzsigndocumentSource(v string)`

SetEEzsigndocumentSource sets EEzsigndocumentSource field to given value.


### GetEEzsigndocumentFormat

`func (o *EzsigndocumentRequest) GetEEzsigndocumentFormat() string`

GetEEzsigndocumentFormat returns the EEzsigndocumentFormat field if non-nil, zero value otherwise.

### GetEEzsigndocumentFormatOk

`func (o *EzsigndocumentRequest) GetEEzsigndocumentFormatOk() (*string, bool)`

GetEEzsigndocumentFormatOk returns a tuple with the EEzsigndocumentFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigndocumentFormat

`func (o *EzsigndocumentRequest) SetEEzsigndocumentFormat(v string)`

SetEEzsigndocumentFormat sets EEzsigndocumentFormat field to given value.


### GetSEzsigndocumentBase64

`func (o *EzsigndocumentRequest) GetSEzsigndocumentBase64() string`

GetSEzsigndocumentBase64 returns the SEzsigndocumentBase64 field if non-nil, zero value otherwise.

### GetSEzsigndocumentBase64Ok

`func (o *EzsigndocumentRequest) GetSEzsigndocumentBase64Ok() (*string, bool)`

GetSEzsigndocumentBase64Ok returns a tuple with the SEzsigndocumentBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigndocumentBase64

`func (o *EzsigndocumentRequest) SetSEzsigndocumentBase64(v string)`

SetSEzsigndocumentBase64 sets SEzsigndocumentBase64 field to given value.

### HasSEzsigndocumentBase64

`func (o *EzsigndocumentRequest) HasSEzsigndocumentBase64() bool`

HasSEzsigndocumentBase64 returns a boolean if a field has been set.

### GetFkiEzsignfolderID

`func (o *EzsigndocumentRequest) GetFkiEzsignfolderID() int32`

GetFkiEzsignfolderID returns the FkiEzsignfolderID field if non-nil, zero value otherwise.

### GetFkiEzsignfolderIDOk

`func (o *EzsigndocumentRequest) GetFkiEzsignfolderIDOk() (*int32, bool)`

GetFkiEzsignfolderIDOk returns a tuple with the FkiEzsignfolderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignfolderID

`func (o *EzsigndocumentRequest) SetFkiEzsignfolderID(v int32)`

SetFkiEzsignfolderID sets FkiEzsignfolderID field to given value.


### GetDtEzsigndocumentDuedate

`func (o *EzsigndocumentRequest) GetDtEzsigndocumentDuedate() string`

GetDtEzsigndocumentDuedate returns the DtEzsigndocumentDuedate field if non-nil, zero value otherwise.

### GetDtEzsigndocumentDuedateOk

`func (o *EzsigndocumentRequest) GetDtEzsigndocumentDuedateOk() (*string, bool)`

GetDtEzsigndocumentDuedateOk returns a tuple with the DtEzsigndocumentDuedate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtEzsigndocumentDuedate

`func (o *EzsigndocumentRequest) SetDtEzsigndocumentDuedate(v string)`

SetDtEzsigndocumentDuedate sets DtEzsigndocumentDuedate field to given value.


### GetFkiLanguageID

`func (o *EzsigndocumentRequest) GetFkiLanguageID() int32`

GetFkiLanguageID returns the FkiLanguageID field if non-nil, zero value otherwise.

### GetFkiLanguageIDOk

`func (o *EzsigndocumentRequest) GetFkiLanguageIDOk() (*int32, bool)`

GetFkiLanguageIDOk returns a tuple with the FkiLanguageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiLanguageID

`func (o *EzsigndocumentRequest) SetFkiLanguageID(v int32)`

SetFkiLanguageID sets FkiLanguageID field to given value.


### GetSEzsigndocumentFilename

`func (o *EzsigndocumentRequest) GetSEzsigndocumentFilename() string`

GetSEzsigndocumentFilename returns the SEzsigndocumentFilename field if non-nil, zero value otherwise.

### GetSEzsigndocumentFilenameOk

`func (o *EzsigndocumentRequest) GetSEzsigndocumentFilenameOk() (*string, bool)`

GetSEzsigndocumentFilenameOk returns a tuple with the SEzsigndocumentFilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigndocumentFilename

`func (o *EzsigndocumentRequest) SetSEzsigndocumentFilename(v string)`

SetSEzsigndocumentFilename sets SEzsigndocumentFilename field to given value.


### GetSEzsigndocumentName

`func (o *EzsigndocumentRequest) GetSEzsigndocumentName() string`

GetSEzsigndocumentName returns the SEzsigndocumentName field if non-nil, zero value otherwise.

### GetSEzsigndocumentNameOk

`func (o *EzsigndocumentRequest) GetSEzsigndocumentNameOk() (*string, bool)`

GetSEzsigndocumentNameOk returns a tuple with the SEzsigndocumentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigndocumentName

`func (o *EzsigndocumentRequest) SetSEzsigndocumentName(v string)`

SetSEzsigndocumentName sets SEzsigndocumentName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


