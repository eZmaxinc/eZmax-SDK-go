# CustomAttachmentImportIntoEDMRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EAttachmentSource** | **string** | The source of the Attachment | 
**FkiAttachmentID** | Pointer to **int32** | The unique ID of the Attachment. | [optional] 
**FkiInscriptionchecklistID** | Pointer to **int32** | The unique ID of the Inscriptionchecklist | [optional] 
**SAttachmentUrl** | Pointer to **string** | The url of the file to import | [optional] 
**SAttachmentBase64** | Pointer to **string** | The Base64 encoded binary content of the attachment. | [optional] 
**SAttachmentName** | **string** | The name of the Attachment | 
**SAttachmentCategory** | **string** | The attachment category | 
**EAttachmentPrivacy** | [**FieldEAttachmentPrivacy**](FieldEAttachmentPrivacy.md) |  | 
**FkiUserIDSpecific** | Pointer to **int32** | The unique ID of the User | [optional] 
**SAttachmentMD5** | Pointer to **string** | The MD5 of the Attachment | [optional] 
**BAttachmentForceoverwrite** | Pointer to **bool** | Whether we force an overwrite of an existing file | [optional] 
**BAttachmentForcerestore** | Pointer to **bool** | Whether we force a restore of a deleted file | [optional] 

## Methods

### NewCustomAttachmentImportIntoEDMRequest

`func NewCustomAttachmentImportIntoEDMRequest(eAttachmentSource string, sAttachmentName string, sAttachmentCategory string, eAttachmentPrivacy FieldEAttachmentPrivacy, ) *CustomAttachmentImportIntoEDMRequest`

NewCustomAttachmentImportIntoEDMRequest instantiates a new CustomAttachmentImportIntoEDMRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomAttachmentImportIntoEDMRequestWithDefaults

`func NewCustomAttachmentImportIntoEDMRequestWithDefaults() *CustomAttachmentImportIntoEDMRequest`

NewCustomAttachmentImportIntoEDMRequestWithDefaults instantiates a new CustomAttachmentImportIntoEDMRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEAttachmentSource

`func (o *CustomAttachmentImportIntoEDMRequest) GetEAttachmentSource() string`

GetEAttachmentSource returns the EAttachmentSource field if non-nil, zero value otherwise.

### GetEAttachmentSourceOk

`func (o *CustomAttachmentImportIntoEDMRequest) GetEAttachmentSourceOk() (*string, bool)`

GetEAttachmentSourceOk returns a tuple with the EAttachmentSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEAttachmentSource

`func (o *CustomAttachmentImportIntoEDMRequest) SetEAttachmentSource(v string)`

SetEAttachmentSource sets EAttachmentSource field to given value.


### GetFkiAttachmentID

`func (o *CustomAttachmentImportIntoEDMRequest) GetFkiAttachmentID() int32`

GetFkiAttachmentID returns the FkiAttachmentID field if non-nil, zero value otherwise.

### GetFkiAttachmentIDOk

`func (o *CustomAttachmentImportIntoEDMRequest) GetFkiAttachmentIDOk() (*int32, bool)`

GetFkiAttachmentIDOk returns a tuple with the FkiAttachmentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiAttachmentID

`func (o *CustomAttachmentImportIntoEDMRequest) SetFkiAttachmentID(v int32)`

SetFkiAttachmentID sets FkiAttachmentID field to given value.

### HasFkiAttachmentID

`func (o *CustomAttachmentImportIntoEDMRequest) HasFkiAttachmentID() bool`

HasFkiAttachmentID returns a boolean if a field has been set.

### GetFkiInscriptionchecklistID

`func (o *CustomAttachmentImportIntoEDMRequest) GetFkiInscriptionchecklistID() int32`

GetFkiInscriptionchecklistID returns the FkiInscriptionchecklistID field if non-nil, zero value otherwise.

### GetFkiInscriptionchecklistIDOk

`func (o *CustomAttachmentImportIntoEDMRequest) GetFkiInscriptionchecklistIDOk() (*int32, bool)`

GetFkiInscriptionchecklistIDOk returns a tuple with the FkiInscriptionchecklistID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiInscriptionchecklistID

`func (o *CustomAttachmentImportIntoEDMRequest) SetFkiInscriptionchecklistID(v int32)`

SetFkiInscriptionchecklistID sets FkiInscriptionchecklistID field to given value.

### HasFkiInscriptionchecklistID

`func (o *CustomAttachmentImportIntoEDMRequest) HasFkiInscriptionchecklistID() bool`

HasFkiInscriptionchecklistID returns a boolean if a field has been set.

### GetSAttachmentUrl

`func (o *CustomAttachmentImportIntoEDMRequest) GetSAttachmentUrl() string`

GetSAttachmentUrl returns the SAttachmentUrl field if non-nil, zero value otherwise.

### GetSAttachmentUrlOk

`func (o *CustomAttachmentImportIntoEDMRequest) GetSAttachmentUrlOk() (*string, bool)`

GetSAttachmentUrlOk returns a tuple with the SAttachmentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAttachmentUrl

`func (o *CustomAttachmentImportIntoEDMRequest) SetSAttachmentUrl(v string)`

SetSAttachmentUrl sets SAttachmentUrl field to given value.

### HasSAttachmentUrl

`func (o *CustomAttachmentImportIntoEDMRequest) HasSAttachmentUrl() bool`

HasSAttachmentUrl returns a boolean if a field has been set.

### GetSAttachmentBase64

`func (o *CustomAttachmentImportIntoEDMRequest) GetSAttachmentBase64() string`

GetSAttachmentBase64 returns the SAttachmentBase64 field if non-nil, zero value otherwise.

### GetSAttachmentBase64Ok

`func (o *CustomAttachmentImportIntoEDMRequest) GetSAttachmentBase64Ok() (*string, bool)`

GetSAttachmentBase64Ok returns a tuple with the SAttachmentBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAttachmentBase64

`func (o *CustomAttachmentImportIntoEDMRequest) SetSAttachmentBase64(v string)`

SetSAttachmentBase64 sets SAttachmentBase64 field to given value.

### HasSAttachmentBase64

`func (o *CustomAttachmentImportIntoEDMRequest) HasSAttachmentBase64() bool`

HasSAttachmentBase64 returns a boolean if a field has been set.

### GetSAttachmentName

`func (o *CustomAttachmentImportIntoEDMRequest) GetSAttachmentName() string`

GetSAttachmentName returns the SAttachmentName field if non-nil, zero value otherwise.

### GetSAttachmentNameOk

`func (o *CustomAttachmentImportIntoEDMRequest) GetSAttachmentNameOk() (*string, bool)`

GetSAttachmentNameOk returns a tuple with the SAttachmentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAttachmentName

`func (o *CustomAttachmentImportIntoEDMRequest) SetSAttachmentName(v string)`

SetSAttachmentName sets SAttachmentName field to given value.


### GetSAttachmentCategory

`func (o *CustomAttachmentImportIntoEDMRequest) GetSAttachmentCategory() string`

GetSAttachmentCategory returns the SAttachmentCategory field if non-nil, zero value otherwise.

### GetSAttachmentCategoryOk

`func (o *CustomAttachmentImportIntoEDMRequest) GetSAttachmentCategoryOk() (*string, bool)`

GetSAttachmentCategoryOk returns a tuple with the SAttachmentCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAttachmentCategory

`func (o *CustomAttachmentImportIntoEDMRequest) SetSAttachmentCategory(v string)`

SetSAttachmentCategory sets SAttachmentCategory field to given value.


### GetEAttachmentPrivacy

`func (o *CustomAttachmentImportIntoEDMRequest) GetEAttachmentPrivacy() FieldEAttachmentPrivacy`

GetEAttachmentPrivacy returns the EAttachmentPrivacy field if non-nil, zero value otherwise.

### GetEAttachmentPrivacyOk

`func (o *CustomAttachmentImportIntoEDMRequest) GetEAttachmentPrivacyOk() (*FieldEAttachmentPrivacy, bool)`

GetEAttachmentPrivacyOk returns a tuple with the EAttachmentPrivacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEAttachmentPrivacy

`func (o *CustomAttachmentImportIntoEDMRequest) SetEAttachmentPrivacy(v FieldEAttachmentPrivacy)`

SetEAttachmentPrivacy sets EAttachmentPrivacy field to given value.


### GetFkiUserIDSpecific

`func (o *CustomAttachmentImportIntoEDMRequest) GetFkiUserIDSpecific() int32`

GetFkiUserIDSpecific returns the FkiUserIDSpecific field if non-nil, zero value otherwise.

### GetFkiUserIDSpecificOk

`func (o *CustomAttachmentImportIntoEDMRequest) GetFkiUserIDSpecificOk() (*int32, bool)`

GetFkiUserIDSpecificOk returns a tuple with the FkiUserIDSpecific field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUserIDSpecific

`func (o *CustomAttachmentImportIntoEDMRequest) SetFkiUserIDSpecific(v int32)`

SetFkiUserIDSpecific sets FkiUserIDSpecific field to given value.

### HasFkiUserIDSpecific

`func (o *CustomAttachmentImportIntoEDMRequest) HasFkiUserIDSpecific() bool`

HasFkiUserIDSpecific returns a boolean if a field has been set.

### GetSAttachmentMD5

`func (o *CustomAttachmentImportIntoEDMRequest) GetSAttachmentMD5() string`

GetSAttachmentMD5 returns the SAttachmentMD5 field if non-nil, zero value otherwise.

### GetSAttachmentMD5Ok

`func (o *CustomAttachmentImportIntoEDMRequest) GetSAttachmentMD5Ok() (*string, bool)`

GetSAttachmentMD5Ok returns a tuple with the SAttachmentMD5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAttachmentMD5

`func (o *CustomAttachmentImportIntoEDMRequest) SetSAttachmentMD5(v string)`

SetSAttachmentMD5 sets SAttachmentMD5 field to given value.

### HasSAttachmentMD5

`func (o *CustomAttachmentImportIntoEDMRequest) HasSAttachmentMD5() bool`

HasSAttachmentMD5 returns a boolean if a field has been set.

### GetBAttachmentForceoverwrite

`func (o *CustomAttachmentImportIntoEDMRequest) GetBAttachmentForceoverwrite() bool`

GetBAttachmentForceoverwrite returns the BAttachmentForceoverwrite field if non-nil, zero value otherwise.

### GetBAttachmentForceoverwriteOk

`func (o *CustomAttachmentImportIntoEDMRequest) GetBAttachmentForceoverwriteOk() (*bool, bool)`

GetBAttachmentForceoverwriteOk returns a tuple with the BAttachmentForceoverwrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBAttachmentForceoverwrite

`func (o *CustomAttachmentImportIntoEDMRequest) SetBAttachmentForceoverwrite(v bool)`

SetBAttachmentForceoverwrite sets BAttachmentForceoverwrite field to given value.

### HasBAttachmentForceoverwrite

`func (o *CustomAttachmentImportIntoEDMRequest) HasBAttachmentForceoverwrite() bool`

HasBAttachmentForceoverwrite returns a boolean if a field has been set.

### GetBAttachmentForcerestore

`func (o *CustomAttachmentImportIntoEDMRequest) GetBAttachmentForcerestore() bool`

GetBAttachmentForcerestore returns the BAttachmentForcerestore field if non-nil, zero value otherwise.

### GetBAttachmentForcerestoreOk

`func (o *CustomAttachmentImportIntoEDMRequest) GetBAttachmentForcerestoreOk() (*bool, bool)`

GetBAttachmentForcerestoreOk returns a tuple with the BAttachmentForcerestore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBAttachmentForcerestore

`func (o *CustomAttachmentImportIntoEDMRequest) SetBAttachmentForcerestore(v bool)`

SetBAttachmentForcerestore sets BAttachmentForcerestore field to given value.

### HasBAttachmentForcerestore

`func (o *CustomAttachmentImportIntoEDMRequest) HasBAttachmentForcerestore() bool`

HasBAttachmentForcerestore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


