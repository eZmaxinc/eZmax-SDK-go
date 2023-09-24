# CommonFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SFileName** | **string** | The name of the file | 
**SFileUrl** | Pointer to **string** | The URL used to reach the File | [optional] 
**SFileBase64** | Pointer to **string** | The Base64 encoded binary content of the File | [optional] 
**EFileSource** | **string** | The source of the File | 

## Methods

### NewCommonFile

`func NewCommonFile(sFileName string, eFileSource string, ) *CommonFile`

NewCommonFile instantiates a new CommonFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonFileWithDefaults

`func NewCommonFileWithDefaults() *CommonFile`

NewCommonFileWithDefaults instantiates a new CommonFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSFileName

`func (o *CommonFile) GetSFileName() string`

GetSFileName returns the SFileName field if non-nil, zero value otherwise.

### GetSFileNameOk

`func (o *CommonFile) GetSFileNameOk() (*string, bool)`

GetSFileNameOk returns a tuple with the SFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSFileName

`func (o *CommonFile) SetSFileName(v string)`

SetSFileName sets SFileName field to given value.


### GetSFileUrl

`func (o *CommonFile) GetSFileUrl() string`

GetSFileUrl returns the SFileUrl field if non-nil, zero value otherwise.

### GetSFileUrlOk

`func (o *CommonFile) GetSFileUrlOk() (*string, bool)`

GetSFileUrlOk returns a tuple with the SFileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSFileUrl

`func (o *CommonFile) SetSFileUrl(v string)`

SetSFileUrl sets SFileUrl field to given value.

### HasSFileUrl

`func (o *CommonFile) HasSFileUrl() bool`

HasSFileUrl returns a boolean if a field has been set.

### GetSFileBase64

`func (o *CommonFile) GetSFileBase64() string`

GetSFileBase64 returns the SFileBase64 field if non-nil, zero value otherwise.

### GetSFileBase64Ok

`func (o *CommonFile) GetSFileBase64Ok() (*string, bool)`

GetSFileBase64Ok returns a tuple with the SFileBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSFileBase64

`func (o *CommonFile) SetSFileBase64(v string)`

SetSFileBase64 sets SFileBase64 field to given value.

### HasSFileBase64

`func (o *CommonFile) HasSFileBase64() bool`

HasSFileBase64 returns a boolean if a field has been set.

### GetEFileSource

`func (o *CommonFile) GetEFileSource() string`

GetEFileSource returns the EFileSource field if non-nil, zero value otherwise.

### GetEFileSourceOk

`func (o *CommonFile) GetEFileSourceOk() (*string, bool)`

GetEFileSourceOk returns a tuple with the EFileSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEFileSource

`func (o *CommonFile) SetEFileSource(v string)`

SetEFileSource sets EFileSource field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


