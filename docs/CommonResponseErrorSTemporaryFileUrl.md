# CommonResponseErrorSTemporaryFileUrl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**STemporaryFileUrl** | Pointer to **string** | The Temporary File Url of the document that was uploaded. That url can be reused instead of uploading the file again. | [optional] 
**SErrorMessage** | **string** | More detail about the error | 
**EErrorCode** | Pointer to **string** | The error code. See documentation for valid values | [optional] 

## Methods

### NewCommonResponseErrorSTemporaryFileUrl

`func NewCommonResponseErrorSTemporaryFileUrl(sErrorMessage string, ) *CommonResponseErrorSTemporaryFileUrl`

NewCommonResponseErrorSTemporaryFileUrl instantiates a new CommonResponseErrorSTemporaryFileUrl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonResponseErrorSTemporaryFileUrlWithDefaults

`func NewCommonResponseErrorSTemporaryFileUrlWithDefaults() *CommonResponseErrorSTemporaryFileUrl`

NewCommonResponseErrorSTemporaryFileUrlWithDefaults instantiates a new CommonResponseErrorSTemporaryFileUrl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSTemporaryFileUrl

`func (o *CommonResponseErrorSTemporaryFileUrl) GetSTemporaryFileUrl() string`

GetSTemporaryFileUrl returns the STemporaryFileUrl field if non-nil, zero value otherwise.

### GetSTemporaryFileUrlOk

`func (o *CommonResponseErrorSTemporaryFileUrl) GetSTemporaryFileUrlOk() (*string, bool)`

GetSTemporaryFileUrlOk returns a tuple with the STemporaryFileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSTemporaryFileUrl

`func (o *CommonResponseErrorSTemporaryFileUrl) SetSTemporaryFileUrl(v string)`

SetSTemporaryFileUrl sets STemporaryFileUrl field to given value.

### HasSTemporaryFileUrl

`func (o *CommonResponseErrorSTemporaryFileUrl) HasSTemporaryFileUrl() bool`

HasSTemporaryFileUrl returns a boolean if a field has been set.

### GetSErrorMessage

`func (o *CommonResponseErrorSTemporaryFileUrl) GetSErrorMessage() string`

GetSErrorMessage returns the SErrorMessage field if non-nil, zero value otherwise.

### GetSErrorMessageOk

`func (o *CommonResponseErrorSTemporaryFileUrl) GetSErrorMessageOk() (*string, bool)`

GetSErrorMessageOk returns a tuple with the SErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSErrorMessage

`func (o *CommonResponseErrorSTemporaryFileUrl) SetSErrorMessage(v string)`

SetSErrorMessage sets SErrorMessage field to given value.


### GetEErrorCode

`func (o *CommonResponseErrorSTemporaryFileUrl) GetEErrorCode() string`

GetEErrorCode returns the EErrorCode field if non-nil, zero value otherwise.

### GetEErrorCodeOk

`func (o *CommonResponseErrorSTemporaryFileUrl) GetEErrorCodeOk() (*string, bool)`

GetEErrorCodeOk returns a tuple with the EErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEErrorCode

`func (o *CommonResponseErrorSTemporaryFileUrl) SetEErrorCode(v string)`

SetEErrorCode sets EErrorCode field to given value.

### HasEErrorCode

`func (o *CommonResponseErrorSTemporaryFileUrl) HasEErrorCode() bool`

HasEErrorCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


