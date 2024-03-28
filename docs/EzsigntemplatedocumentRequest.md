# EzsigntemplatedocumentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsigntemplatedocumentID** | Pointer to **int32** | The unique ID of the Ezsigntemplatedocument | [optional] 
**FkiEzsigntemplateID** | **int32** | The unique ID of the Ezsigntemplate | 
**FkiEzsigndocumentID** | Pointer to **int32** | The unique ID of the Ezsigndocument | [optional] 
**FkiEzsigntemplatesignerID** | Pointer to **int32** | The unique ID of the Ezsigntemplatesigner | [optional] 
**SEzsigntemplatedocumentName** | **string** | The name of the Ezsigntemplatedocument. | 
**EEzsigntemplatedocumentSource** | **string** | Indicates where to look for the document binary content. | 
**EEzsigntemplatedocumentFormat** | Pointer to **string** | Indicates the format of the template. | [optional] 
**SEzsigntemplatedocumentBase64** | Pointer to **string** | The Base64 encoded binary content of the document.  This field is Required when eEzsigntemplatedocumentSource &#x3D; Base64. | [optional] 
**SEzsigntemplatedocumentUrl** | Pointer to **string** | The url where the document content resides.  This field is Required when eEzsigntemplatedocumentSource &#x3D; Url. | [optional] 
**BEzsigntemplatedocumentForcerepair** | Pointer to **bool** | Try to repair the document or flatten it if it cannot be used for electronic signature. | [optional] 
**EEzsigntemplatedocumentForm** | Pointer to **string** | If the document contains an existing PDF form this property must be set.  **Keep** leaves the form as-is in the document.  **Convert** removes the form and convert all the existing fields to Ezsigntemplateformfieldgroups and assign them to the specified **fkiEzsigntemplatesignerID**  **Discard** removes the form from the document | [optional] 
**SEzsigntemplatedocumentPassword** | Pointer to **string** | If the source template is password protected, the password to open/modify it. | [optional] [default to ""]

## Methods

### NewEzsigntemplatedocumentRequest

`func NewEzsigntemplatedocumentRequest(fkiEzsigntemplateID int32, sEzsigntemplatedocumentName string, eEzsigntemplatedocumentSource string, ) *EzsigntemplatedocumentRequest`

NewEzsigntemplatedocumentRequest instantiates a new EzsigntemplatedocumentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigntemplatedocumentRequestWithDefaults

`func NewEzsigntemplatedocumentRequestWithDefaults() *EzsigntemplatedocumentRequest`

NewEzsigntemplatedocumentRequestWithDefaults instantiates a new EzsigntemplatedocumentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsigntemplatedocumentID

`func (o *EzsigntemplatedocumentRequest) GetPkiEzsigntemplatedocumentID() int32`

GetPkiEzsigntemplatedocumentID returns the PkiEzsigntemplatedocumentID field if non-nil, zero value otherwise.

### GetPkiEzsigntemplatedocumentIDOk

`func (o *EzsigntemplatedocumentRequest) GetPkiEzsigntemplatedocumentIDOk() (*int32, bool)`

GetPkiEzsigntemplatedocumentIDOk returns a tuple with the PkiEzsigntemplatedocumentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsigntemplatedocumentID

`func (o *EzsigntemplatedocumentRequest) SetPkiEzsigntemplatedocumentID(v int32)`

SetPkiEzsigntemplatedocumentID sets PkiEzsigntemplatedocumentID field to given value.

### HasPkiEzsigntemplatedocumentID

`func (o *EzsigntemplatedocumentRequest) HasPkiEzsigntemplatedocumentID() bool`

HasPkiEzsigntemplatedocumentID returns a boolean if a field has been set.

### GetFkiEzsigntemplateID

`func (o *EzsigntemplatedocumentRequest) GetFkiEzsigntemplateID() int32`

GetFkiEzsigntemplateID returns the FkiEzsigntemplateID field if non-nil, zero value otherwise.

### GetFkiEzsigntemplateIDOk

`func (o *EzsigntemplatedocumentRequest) GetFkiEzsigntemplateIDOk() (*int32, bool)`

GetFkiEzsigntemplateIDOk returns a tuple with the FkiEzsigntemplateID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigntemplateID

`func (o *EzsigntemplatedocumentRequest) SetFkiEzsigntemplateID(v int32)`

SetFkiEzsigntemplateID sets FkiEzsigntemplateID field to given value.


### GetFkiEzsigndocumentID

`func (o *EzsigntemplatedocumentRequest) GetFkiEzsigndocumentID() int32`

GetFkiEzsigndocumentID returns the FkiEzsigndocumentID field if non-nil, zero value otherwise.

### GetFkiEzsigndocumentIDOk

`func (o *EzsigntemplatedocumentRequest) GetFkiEzsigndocumentIDOk() (*int32, bool)`

GetFkiEzsigndocumentIDOk returns a tuple with the FkiEzsigndocumentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigndocumentID

`func (o *EzsigntemplatedocumentRequest) SetFkiEzsigndocumentID(v int32)`

SetFkiEzsigndocumentID sets FkiEzsigndocumentID field to given value.

### HasFkiEzsigndocumentID

`func (o *EzsigntemplatedocumentRequest) HasFkiEzsigndocumentID() bool`

HasFkiEzsigndocumentID returns a boolean if a field has been set.

### GetFkiEzsigntemplatesignerID

`func (o *EzsigntemplatedocumentRequest) GetFkiEzsigntemplatesignerID() int32`

GetFkiEzsigntemplatesignerID returns the FkiEzsigntemplatesignerID field if non-nil, zero value otherwise.

### GetFkiEzsigntemplatesignerIDOk

`func (o *EzsigntemplatedocumentRequest) GetFkiEzsigntemplatesignerIDOk() (*int32, bool)`

GetFkiEzsigntemplatesignerIDOk returns a tuple with the FkiEzsigntemplatesignerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigntemplatesignerID

`func (o *EzsigntemplatedocumentRequest) SetFkiEzsigntemplatesignerID(v int32)`

SetFkiEzsigntemplatesignerID sets FkiEzsigntemplatesignerID field to given value.

### HasFkiEzsigntemplatesignerID

`func (o *EzsigntemplatedocumentRequest) HasFkiEzsigntemplatesignerID() bool`

HasFkiEzsigntemplatesignerID returns a boolean if a field has been set.

### GetSEzsigntemplatedocumentName

`func (o *EzsigntemplatedocumentRequest) GetSEzsigntemplatedocumentName() string`

GetSEzsigntemplatedocumentName returns the SEzsigntemplatedocumentName field if non-nil, zero value otherwise.

### GetSEzsigntemplatedocumentNameOk

`func (o *EzsigntemplatedocumentRequest) GetSEzsigntemplatedocumentNameOk() (*string, bool)`

GetSEzsigntemplatedocumentNameOk returns a tuple with the SEzsigntemplatedocumentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplatedocumentName

`func (o *EzsigntemplatedocumentRequest) SetSEzsigntemplatedocumentName(v string)`

SetSEzsigntemplatedocumentName sets SEzsigntemplatedocumentName field to given value.


### GetEEzsigntemplatedocumentSource

`func (o *EzsigntemplatedocumentRequest) GetEEzsigntemplatedocumentSource() string`

GetEEzsigntemplatedocumentSource returns the EEzsigntemplatedocumentSource field if non-nil, zero value otherwise.

### GetEEzsigntemplatedocumentSourceOk

`func (o *EzsigntemplatedocumentRequest) GetEEzsigntemplatedocumentSourceOk() (*string, bool)`

GetEEzsigntemplatedocumentSourceOk returns a tuple with the EEzsigntemplatedocumentSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplatedocumentSource

`func (o *EzsigntemplatedocumentRequest) SetEEzsigntemplatedocumentSource(v string)`

SetEEzsigntemplatedocumentSource sets EEzsigntemplatedocumentSource field to given value.


### GetEEzsigntemplatedocumentFormat

`func (o *EzsigntemplatedocumentRequest) GetEEzsigntemplatedocumentFormat() string`

GetEEzsigntemplatedocumentFormat returns the EEzsigntemplatedocumentFormat field if non-nil, zero value otherwise.

### GetEEzsigntemplatedocumentFormatOk

`func (o *EzsigntemplatedocumentRequest) GetEEzsigntemplatedocumentFormatOk() (*string, bool)`

GetEEzsigntemplatedocumentFormatOk returns a tuple with the EEzsigntemplatedocumentFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplatedocumentFormat

`func (o *EzsigntemplatedocumentRequest) SetEEzsigntemplatedocumentFormat(v string)`

SetEEzsigntemplatedocumentFormat sets EEzsigntemplatedocumentFormat field to given value.

### HasEEzsigntemplatedocumentFormat

`func (o *EzsigntemplatedocumentRequest) HasEEzsigntemplatedocumentFormat() bool`

HasEEzsigntemplatedocumentFormat returns a boolean if a field has been set.

### GetSEzsigntemplatedocumentBase64

`func (o *EzsigntemplatedocumentRequest) GetSEzsigntemplatedocumentBase64() string`

GetSEzsigntemplatedocumentBase64 returns the SEzsigntemplatedocumentBase64 field if non-nil, zero value otherwise.

### GetSEzsigntemplatedocumentBase64Ok

`func (o *EzsigntemplatedocumentRequest) GetSEzsigntemplatedocumentBase64Ok() (*string, bool)`

GetSEzsigntemplatedocumentBase64Ok returns a tuple with the SEzsigntemplatedocumentBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplatedocumentBase64

`func (o *EzsigntemplatedocumentRequest) SetSEzsigntemplatedocumentBase64(v string)`

SetSEzsigntemplatedocumentBase64 sets SEzsigntemplatedocumentBase64 field to given value.

### HasSEzsigntemplatedocumentBase64

`func (o *EzsigntemplatedocumentRequest) HasSEzsigntemplatedocumentBase64() bool`

HasSEzsigntemplatedocumentBase64 returns a boolean if a field has been set.

### GetSEzsigntemplatedocumentUrl

`func (o *EzsigntemplatedocumentRequest) GetSEzsigntemplatedocumentUrl() string`

GetSEzsigntemplatedocumentUrl returns the SEzsigntemplatedocumentUrl field if non-nil, zero value otherwise.

### GetSEzsigntemplatedocumentUrlOk

`func (o *EzsigntemplatedocumentRequest) GetSEzsigntemplatedocumentUrlOk() (*string, bool)`

GetSEzsigntemplatedocumentUrlOk returns a tuple with the SEzsigntemplatedocumentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplatedocumentUrl

`func (o *EzsigntemplatedocumentRequest) SetSEzsigntemplatedocumentUrl(v string)`

SetSEzsigntemplatedocumentUrl sets SEzsigntemplatedocumentUrl field to given value.

### HasSEzsigntemplatedocumentUrl

`func (o *EzsigntemplatedocumentRequest) HasSEzsigntemplatedocumentUrl() bool`

HasSEzsigntemplatedocumentUrl returns a boolean if a field has been set.

### GetBEzsigntemplatedocumentForcerepair

`func (o *EzsigntemplatedocumentRequest) GetBEzsigntemplatedocumentForcerepair() bool`

GetBEzsigntemplatedocumentForcerepair returns the BEzsigntemplatedocumentForcerepair field if non-nil, zero value otherwise.

### GetBEzsigntemplatedocumentForcerepairOk

`func (o *EzsigntemplatedocumentRequest) GetBEzsigntemplatedocumentForcerepairOk() (*bool, bool)`

GetBEzsigntemplatedocumentForcerepairOk returns a tuple with the BEzsigntemplatedocumentForcerepair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsigntemplatedocumentForcerepair

`func (o *EzsigntemplatedocumentRequest) SetBEzsigntemplatedocumentForcerepair(v bool)`

SetBEzsigntemplatedocumentForcerepair sets BEzsigntemplatedocumentForcerepair field to given value.

### HasBEzsigntemplatedocumentForcerepair

`func (o *EzsigntemplatedocumentRequest) HasBEzsigntemplatedocumentForcerepair() bool`

HasBEzsigntemplatedocumentForcerepair returns a boolean if a field has been set.

### GetEEzsigntemplatedocumentForm

`func (o *EzsigntemplatedocumentRequest) GetEEzsigntemplatedocumentForm() string`

GetEEzsigntemplatedocumentForm returns the EEzsigntemplatedocumentForm field if non-nil, zero value otherwise.

### GetEEzsigntemplatedocumentFormOk

`func (o *EzsigntemplatedocumentRequest) GetEEzsigntemplatedocumentFormOk() (*string, bool)`

GetEEzsigntemplatedocumentFormOk returns a tuple with the EEzsigntemplatedocumentForm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplatedocumentForm

`func (o *EzsigntemplatedocumentRequest) SetEEzsigntemplatedocumentForm(v string)`

SetEEzsigntemplatedocumentForm sets EEzsigntemplatedocumentForm field to given value.

### HasEEzsigntemplatedocumentForm

`func (o *EzsigntemplatedocumentRequest) HasEEzsigntemplatedocumentForm() bool`

HasEEzsigntemplatedocumentForm returns a boolean if a field has been set.

### GetSEzsigntemplatedocumentPassword

`func (o *EzsigntemplatedocumentRequest) GetSEzsigntemplatedocumentPassword() string`

GetSEzsigntemplatedocumentPassword returns the SEzsigntemplatedocumentPassword field if non-nil, zero value otherwise.

### GetSEzsigntemplatedocumentPasswordOk

`func (o *EzsigntemplatedocumentRequest) GetSEzsigntemplatedocumentPasswordOk() (*string, bool)`

GetSEzsigntemplatedocumentPasswordOk returns a tuple with the SEzsigntemplatedocumentPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplatedocumentPassword

`func (o *EzsigntemplatedocumentRequest) SetSEzsigntemplatedocumentPassword(v string)`

SetSEzsigntemplatedocumentPassword sets SEzsigntemplatedocumentPassword field to given value.

### HasSEzsigntemplatedocumentPassword

`func (o *EzsigntemplatedocumentRequest) HasSEzsigntemplatedocumentPassword() bool`

HasSEzsigntemplatedocumentPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


