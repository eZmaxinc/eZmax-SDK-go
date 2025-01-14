# CommonResponseErrorSTemporaryFileUrl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SErrorMessage** | **string** | The message giving details about the error | 
**EErrorCode** | [**FieldEErrorCode**](FieldEErrorCode.md) |  | 
**ASErrorMessagedetail** | Pointer to **[]string** | More error message detail | [optional] 
**STemporaryFileUrl** | Pointer to **string** | The Temporary File Url of the document that was uploaded. That url can be reused instead of uploading the file again. | [optional] 

## Methods

### NewCommonResponseErrorSTemporaryFileUrl

`func NewCommonResponseErrorSTemporaryFileUrl(sErrorMessage string, eErrorCode FieldEErrorCode, ) *CommonResponseErrorSTemporaryFileUrl`

NewCommonResponseErrorSTemporaryFileUrl instantiates a new CommonResponseErrorSTemporaryFileUrl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonResponseErrorSTemporaryFileUrlWithDefaults

`func NewCommonResponseErrorSTemporaryFileUrlWithDefaults() *CommonResponseErrorSTemporaryFileUrl`

NewCommonResponseErrorSTemporaryFileUrlWithDefaults instantiates a new CommonResponseErrorSTemporaryFileUrl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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

`func (o *CommonResponseErrorSTemporaryFileUrl) GetEErrorCode() FieldEErrorCode`

GetEErrorCode returns the EErrorCode field if non-nil, zero value otherwise.

### GetEErrorCodeOk

`func (o *CommonResponseErrorSTemporaryFileUrl) GetEErrorCodeOk() (*FieldEErrorCode, bool)`

GetEErrorCodeOk returns a tuple with the EErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEErrorCode

`func (o *CommonResponseErrorSTemporaryFileUrl) SetEErrorCode(v FieldEErrorCode)`

SetEErrorCode sets EErrorCode field to given value.


### GetASErrorMessagedetail

`func (o *CommonResponseErrorSTemporaryFileUrl) GetASErrorMessagedetail() []string`

GetASErrorMessagedetail returns the ASErrorMessagedetail field if non-nil, zero value otherwise.

### GetASErrorMessagedetailOk

`func (o *CommonResponseErrorSTemporaryFileUrl) GetASErrorMessagedetailOk() (*[]string, bool)`

GetASErrorMessagedetailOk returns a tuple with the ASErrorMessagedetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetASErrorMessagedetail

`func (o *CommonResponseErrorSTemporaryFileUrl) SetASErrorMessagedetail(v []string)`

SetASErrorMessagedetail sets ASErrorMessagedetail field to given value.

### HasASErrorMessagedetail

`func (o *CommonResponseErrorSTemporaryFileUrl) HasASErrorMessagedetail() bool`

HasASErrorMessagedetail returns a boolean if a field has been set.

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


