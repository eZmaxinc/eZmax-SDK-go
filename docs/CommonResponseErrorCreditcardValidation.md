# CommonResponseErrorCreditcardValidation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SErrorMessage** | **string** | The message giving details about the error | 
**EErrorCode** | [**FieldEErrorCode**](FieldEErrorCode.md) |  | 
**ASErrorMessagedetail** | Pointer to **[]string** | More error message detail | [optional] 
**ObjCreditcardtransactionresponse** | Pointer to [**CustomCreditcardtransactionresponseResponse**](CustomCreditcardtransactionresponseResponse.md) |  | [optional] 

## Methods

### NewCommonResponseErrorCreditcardValidation

`func NewCommonResponseErrorCreditcardValidation(sErrorMessage string, eErrorCode FieldEErrorCode, ) *CommonResponseErrorCreditcardValidation`

NewCommonResponseErrorCreditcardValidation instantiates a new CommonResponseErrorCreditcardValidation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonResponseErrorCreditcardValidationWithDefaults

`func NewCommonResponseErrorCreditcardValidationWithDefaults() *CommonResponseErrorCreditcardValidation`

NewCommonResponseErrorCreditcardValidationWithDefaults instantiates a new CommonResponseErrorCreditcardValidation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSErrorMessage

`func (o *CommonResponseErrorCreditcardValidation) GetSErrorMessage() string`

GetSErrorMessage returns the SErrorMessage field if non-nil, zero value otherwise.

### GetSErrorMessageOk

`func (o *CommonResponseErrorCreditcardValidation) GetSErrorMessageOk() (*string, bool)`

GetSErrorMessageOk returns a tuple with the SErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSErrorMessage

`func (o *CommonResponseErrorCreditcardValidation) SetSErrorMessage(v string)`

SetSErrorMessage sets SErrorMessage field to given value.


### GetEErrorCode

`func (o *CommonResponseErrorCreditcardValidation) GetEErrorCode() FieldEErrorCode`

GetEErrorCode returns the EErrorCode field if non-nil, zero value otherwise.

### GetEErrorCodeOk

`func (o *CommonResponseErrorCreditcardValidation) GetEErrorCodeOk() (*FieldEErrorCode, bool)`

GetEErrorCodeOk returns a tuple with the EErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEErrorCode

`func (o *CommonResponseErrorCreditcardValidation) SetEErrorCode(v FieldEErrorCode)`

SetEErrorCode sets EErrorCode field to given value.


### GetASErrorMessagedetail

`func (o *CommonResponseErrorCreditcardValidation) GetASErrorMessagedetail() []string`

GetASErrorMessagedetail returns the ASErrorMessagedetail field if non-nil, zero value otherwise.

### GetASErrorMessagedetailOk

`func (o *CommonResponseErrorCreditcardValidation) GetASErrorMessagedetailOk() (*[]string, bool)`

GetASErrorMessagedetailOk returns a tuple with the ASErrorMessagedetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetASErrorMessagedetail

`func (o *CommonResponseErrorCreditcardValidation) SetASErrorMessagedetail(v []string)`

SetASErrorMessagedetail sets ASErrorMessagedetail field to given value.

### HasASErrorMessagedetail

`func (o *CommonResponseErrorCreditcardValidation) HasASErrorMessagedetail() bool`

HasASErrorMessagedetail returns a boolean if a field has been set.

### GetObjCreditcardtransactionresponse

`func (o *CommonResponseErrorCreditcardValidation) GetObjCreditcardtransactionresponse() CustomCreditcardtransactionresponseResponse`

GetObjCreditcardtransactionresponse returns the ObjCreditcardtransactionresponse field if non-nil, zero value otherwise.

### GetObjCreditcardtransactionresponseOk

`func (o *CommonResponseErrorCreditcardValidation) GetObjCreditcardtransactionresponseOk() (*CustomCreditcardtransactionresponseResponse, bool)`

GetObjCreditcardtransactionresponseOk returns a tuple with the ObjCreditcardtransactionresponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjCreditcardtransactionresponse

`func (o *CommonResponseErrorCreditcardValidation) SetObjCreditcardtransactionresponse(v CustomCreditcardtransactionresponseResponse)`

SetObjCreditcardtransactionresponse sets ObjCreditcardtransactionresponse field to given value.

### HasObjCreditcardtransactionresponse

`func (o *CommonResponseErrorCreditcardValidation) HasObjCreditcardtransactionresponse() bool`

HasObjCreditcardtransactionresponse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


