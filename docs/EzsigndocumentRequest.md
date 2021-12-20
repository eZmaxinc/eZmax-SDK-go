# EzsigndocumentRequest

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
**DtEzsigndocumentDuedate** | **string** | The maximum date and time at which the Ezsigndocument can be signed. | 
**FkiLanguageID** | **int32** | The unique ID of the Language.  Valid values:  |Value|Description| |-|-| |1|French| |2|English| | 
**SEzsigndocumentName** | **string** | The name of the document that will be presented to Ezsignfoldersignerassociations | 

## Methods

### NewEzsigndocumentRequest

`func NewEzsigndocumentRequest(eEzsigndocumentSource string, eEzsigndocumentFormat string, fkiEzsignfolderID int32, dtEzsigndocumentDuedate string, fkiLanguageID int32, sEzsigndocumentName string, ) *EzsigndocumentRequest`

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

### GetSEzsigndocumentUrl

`func (o *EzsigndocumentRequest) GetSEzsigndocumentUrl() string`

GetSEzsigndocumentUrl returns the SEzsigndocumentUrl field if non-nil, zero value otherwise.

### GetSEzsigndocumentUrlOk

`func (o *EzsigndocumentRequest) GetSEzsigndocumentUrlOk() (*string, bool)`

GetSEzsigndocumentUrlOk returns a tuple with the SEzsigndocumentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigndocumentUrl

`func (o *EzsigndocumentRequest) SetSEzsigndocumentUrl(v string)`

SetSEzsigndocumentUrl sets SEzsigndocumentUrl field to given value.

### HasSEzsigndocumentUrl

`func (o *EzsigndocumentRequest) HasSEzsigndocumentUrl() bool`

HasSEzsigndocumentUrl returns a boolean if a field has been set.

### GetBEzsigndocumentForcerepair

`func (o *EzsigndocumentRequest) GetBEzsigndocumentForcerepair() bool`

GetBEzsigndocumentForcerepair returns the BEzsigndocumentForcerepair field if non-nil, zero value otherwise.

### GetBEzsigndocumentForcerepairOk

`func (o *EzsigndocumentRequest) GetBEzsigndocumentForcerepairOk() (*bool, bool)`

GetBEzsigndocumentForcerepairOk returns a tuple with the BEzsigndocumentForcerepair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsigndocumentForcerepair

`func (o *EzsigndocumentRequest) SetBEzsigndocumentForcerepair(v bool)`

SetBEzsigndocumentForcerepair sets BEzsigndocumentForcerepair field to given value.

### HasBEzsigndocumentForcerepair

`func (o *EzsigndocumentRequest) HasBEzsigndocumentForcerepair() bool`

HasBEzsigndocumentForcerepair returns a boolean if a field has been set.

### GetSEzsigndocumentPassword

`func (o *EzsigndocumentRequest) GetSEzsigndocumentPassword() string`

GetSEzsigndocumentPassword returns the SEzsigndocumentPassword field if non-nil, zero value otherwise.

### GetSEzsigndocumentPasswordOk

`func (o *EzsigndocumentRequest) GetSEzsigndocumentPasswordOk() (*string, bool)`

GetSEzsigndocumentPasswordOk returns a tuple with the SEzsigndocumentPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigndocumentPassword

`func (o *EzsigndocumentRequest) SetSEzsigndocumentPassword(v string)`

SetSEzsigndocumentPassword sets SEzsigndocumentPassword field to given value.

### HasSEzsigndocumentPassword

`func (o *EzsigndocumentRequest) HasSEzsigndocumentPassword() bool`

HasSEzsigndocumentPassword returns a boolean if a field has been set.

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


