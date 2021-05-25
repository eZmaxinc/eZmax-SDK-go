# CommonResponseError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SErrorMessage** | **string** | More detail about the error | 
**EErrorCode** | Pointer to **string** | The error code. See documentation for valid values | [optional] 

## Methods

### NewCommonResponseError

`func NewCommonResponseError(sErrorMessage string, ) *CommonResponseError`

NewCommonResponseError instantiates a new CommonResponseError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonResponseErrorWithDefaults

`func NewCommonResponseErrorWithDefaults() *CommonResponseError`

NewCommonResponseErrorWithDefaults instantiates a new CommonResponseError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSErrorMessage

`func (o *CommonResponseError) GetSErrorMessage() string`

GetSErrorMessage returns the SErrorMessage field if non-nil, zero value otherwise.

### GetSErrorMessageOk

`func (o *CommonResponseError) GetSErrorMessageOk() (*string, bool)`

GetSErrorMessageOk returns a tuple with the SErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSErrorMessage

`func (o *CommonResponseError) SetSErrorMessage(v string)`

SetSErrorMessage sets SErrorMessage field to given value.


### GetEErrorCode

`func (o *CommonResponseError) GetEErrorCode() string`

GetEErrorCode returns the EErrorCode field if non-nil, zero value otherwise.

### GetEErrorCodeOk

`func (o *CommonResponseError) GetEErrorCodeOk() (*string, bool)`

GetEErrorCodeOk returns a tuple with the EErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEErrorCode

`func (o *CommonResponseError) SetEErrorCode(v string)`

SetEErrorCode sets EErrorCode field to given value.

### HasEErrorCode

`func (o *CommonResponseError) HasEErrorCode() bool`

HasEErrorCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


