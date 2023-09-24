# CommonResponseWarning

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SWarningMessage** | **string** | More detail about the warning | 
**EWarningCode** | **string** | The warning code. See documentation for valid values | 

## Methods

### NewCommonResponseWarning

`func NewCommonResponseWarning(sWarningMessage string, eWarningCode string, ) *CommonResponseWarning`

NewCommonResponseWarning instantiates a new CommonResponseWarning object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonResponseWarningWithDefaults

`func NewCommonResponseWarningWithDefaults() *CommonResponseWarning`

NewCommonResponseWarningWithDefaults instantiates a new CommonResponseWarning object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSWarningMessage

`func (o *CommonResponseWarning) GetSWarningMessage() string`

GetSWarningMessage returns the SWarningMessage field if non-nil, zero value otherwise.

### GetSWarningMessageOk

`func (o *CommonResponseWarning) GetSWarningMessageOk() (*string, bool)`

GetSWarningMessageOk returns a tuple with the SWarningMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSWarningMessage

`func (o *CommonResponseWarning) SetSWarningMessage(v string)`

SetSWarningMessage sets SWarningMessage field to given value.


### GetEWarningCode

`func (o *CommonResponseWarning) GetEWarningCode() string`

GetEWarningCode returns the EWarningCode field if non-nil, zero value otherwise.

### GetEWarningCodeOk

`func (o *CommonResponseWarning) GetEWarningCodeOk() (*string, bool)`

GetEWarningCodeOk returns a tuple with the EWarningCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEWarningCode

`func (o *CommonResponseWarning) SetEWarningCode(v string)`

SetEWarningCode sets EWarningCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


