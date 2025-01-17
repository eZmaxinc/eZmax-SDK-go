# CommonResponseErrorTooManyRequests

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SErrorMessage** | **string** | The message giving details about the error | 
**EErrorCode** | [**FieldEErrorCode**](FieldEErrorCode.md) |  | 
**ASErrorMessagedetail** | Pointer to **[]string** | More error message detail | [optional] 

## Methods

### NewCommonResponseErrorTooManyRequests

`func NewCommonResponseErrorTooManyRequests(sErrorMessage string, eErrorCode FieldEErrorCode, ) *CommonResponseErrorTooManyRequests`

NewCommonResponseErrorTooManyRequests instantiates a new CommonResponseErrorTooManyRequests object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonResponseErrorTooManyRequestsWithDefaults

`func NewCommonResponseErrorTooManyRequestsWithDefaults() *CommonResponseErrorTooManyRequests`

NewCommonResponseErrorTooManyRequestsWithDefaults instantiates a new CommonResponseErrorTooManyRequests object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSErrorMessage

`func (o *CommonResponseErrorTooManyRequests) GetSErrorMessage() string`

GetSErrorMessage returns the SErrorMessage field if non-nil, zero value otherwise.

### GetSErrorMessageOk

`func (o *CommonResponseErrorTooManyRequests) GetSErrorMessageOk() (*string, bool)`

GetSErrorMessageOk returns a tuple with the SErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSErrorMessage

`func (o *CommonResponseErrorTooManyRequests) SetSErrorMessage(v string)`

SetSErrorMessage sets SErrorMessage field to given value.


### GetEErrorCode

`func (o *CommonResponseErrorTooManyRequests) GetEErrorCode() FieldEErrorCode`

GetEErrorCode returns the EErrorCode field if non-nil, zero value otherwise.

### GetEErrorCodeOk

`func (o *CommonResponseErrorTooManyRequests) GetEErrorCodeOk() (*FieldEErrorCode, bool)`

GetEErrorCodeOk returns a tuple with the EErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEErrorCode

`func (o *CommonResponseErrorTooManyRequests) SetEErrorCode(v FieldEErrorCode)`

SetEErrorCode sets EErrorCode field to given value.


### GetASErrorMessagedetail

`func (o *CommonResponseErrorTooManyRequests) GetASErrorMessagedetail() []string`

GetASErrorMessagedetail returns the ASErrorMessagedetail field if non-nil, zero value otherwise.

### GetASErrorMessagedetailOk

`func (o *CommonResponseErrorTooManyRequests) GetASErrorMessagedetailOk() (*[]string, bool)`

GetASErrorMessagedetailOk returns a tuple with the ASErrorMessagedetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetASErrorMessagedetail

`func (o *CommonResponseErrorTooManyRequests) SetASErrorMessagedetail(v []string)`

SetASErrorMessagedetail sets ASErrorMessagedetail field to given value.

### HasASErrorMessagedetail

`func (o *CommonResponseErrorTooManyRequests) HasASErrorMessagedetail() bool`

HasASErrorMessagedetail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


