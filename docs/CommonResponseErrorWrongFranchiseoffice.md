# CommonResponseErrorWrongFranchiseoffice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SErrorMessage** | **string** | The message giving details about the error | 
**EErrorCode** | [**FieldEErrorCode**](FieldEErrorCode.md) |  | 
**ASErrorMessagedetail** | Pointer to **[]string** | More error message detail | [optional] 
**FkiFranchiseagenceID** | **int32** | The unique ID of the Franchiseagence | 
**SFranchiseagenceName** | **string** | The name of the Franchiseagence | 
**FkiFranchiseofficeID** | **int32** | The unique ID of the Franchisereoffice | 
**IFranchiseofficeCode** | **string** | The code of the Franchiseoffice | 

## Methods

### NewCommonResponseErrorWrongFranchiseoffice

`func NewCommonResponseErrorWrongFranchiseoffice(sErrorMessage string, eErrorCode FieldEErrorCode, fkiFranchiseagenceID int32, sFranchiseagenceName string, fkiFranchiseofficeID int32, iFranchiseofficeCode string, ) *CommonResponseErrorWrongFranchiseoffice`

NewCommonResponseErrorWrongFranchiseoffice instantiates a new CommonResponseErrorWrongFranchiseoffice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonResponseErrorWrongFranchiseofficeWithDefaults

`func NewCommonResponseErrorWrongFranchiseofficeWithDefaults() *CommonResponseErrorWrongFranchiseoffice`

NewCommonResponseErrorWrongFranchiseofficeWithDefaults instantiates a new CommonResponseErrorWrongFranchiseoffice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSErrorMessage

`func (o *CommonResponseErrorWrongFranchiseoffice) GetSErrorMessage() string`

GetSErrorMessage returns the SErrorMessage field if non-nil, zero value otherwise.

### GetSErrorMessageOk

`func (o *CommonResponseErrorWrongFranchiseoffice) GetSErrorMessageOk() (*string, bool)`

GetSErrorMessageOk returns a tuple with the SErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSErrorMessage

`func (o *CommonResponseErrorWrongFranchiseoffice) SetSErrorMessage(v string)`

SetSErrorMessage sets SErrorMessage field to given value.


### GetEErrorCode

`func (o *CommonResponseErrorWrongFranchiseoffice) GetEErrorCode() FieldEErrorCode`

GetEErrorCode returns the EErrorCode field if non-nil, zero value otherwise.

### GetEErrorCodeOk

`func (o *CommonResponseErrorWrongFranchiseoffice) GetEErrorCodeOk() (*FieldEErrorCode, bool)`

GetEErrorCodeOk returns a tuple with the EErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEErrorCode

`func (o *CommonResponseErrorWrongFranchiseoffice) SetEErrorCode(v FieldEErrorCode)`

SetEErrorCode sets EErrorCode field to given value.


### GetASErrorMessagedetail

`func (o *CommonResponseErrorWrongFranchiseoffice) GetASErrorMessagedetail() []string`

GetASErrorMessagedetail returns the ASErrorMessagedetail field if non-nil, zero value otherwise.

### GetASErrorMessagedetailOk

`func (o *CommonResponseErrorWrongFranchiseoffice) GetASErrorMessagedetailOk() (*[]string, bool)`

GetASErrorMessagedetailOk returns a tuple with the ASErrorMessagedetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetASErrorMessagedetail

`func (o *CommonResponseErrorWrongFranchiseoffice) SetASErrorMessagedetail(v []string)`

SetASErrorMessagedetail sets ASErrorMessagedetail field to given value.

### HasASErrorMessagedetail

`func (o *CommonResponseErrorWrongFranchiseoffice) HasASErrorMessagedetail() bool`

HasASErrorMessagedetail returns a boolean if a field has been set.

### GetFkiFranchiseagenceID

`func (o *CommonResponseErrorWrongFranchiseoffice) GetFkiFranchiseagenceID() int32`

GetFkiFranchiseagenceID returns the FkiFranchiseagenceID field if non-nil, zero value otherwise.

### GetFkiFranchiseagenceIDOk

`func (o *CommonResponseErrorWrongFranchiseoffice) GetFkiFranchiseagenceIDOk() (*int32, bool)`

GetFkiFranchiseagenceIDOk returns a tuple with the FkiFranchiseagenceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiFranchiseagenceID

`func (o *CommonResponseErrorWrongFranchiseoffice) SetFkiFranchiseagenceID(v int32)`

SetFkiFranchiseagenceID sets FkiFranchiseagenceID field to given value.


### GetSFranchiseagenceName

`func (o *CommonResponseErrorWrongFranchiseoffice) GetSFranchiseagenceName() string`

GetSFranchiseagenceName returns the SFranchiseagenceName field if non-nil, zero value otherwise.

### GetSFranchiseagenceNameOk

`func (o *CommonResponseErrorWrongFranchiseoffice) GetSFranchiseagenceNameOk() (*string, bool)`

GetSFranchiseagenceNameOk returns a tuple with the SFranchiseagenceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSFranchiseagenceName

`func (o *CommonResponseErrorWrongFranchiseoffice) SetSFranchiseagenceName(v string)`

SetSFranchiseagenceName sets SFranchiseagenceName field to given value.


### GetFkiFranchiseofficeID

`func (o *CommonResponseErrorWrongFranchiseoffice) GetFkiFranchiseofficeID() int32`

GetFkiFranchiseofficeID returns the FkiFranchiseofficeID field if non-nil, zero value otherwise.

### GetFkiFranchiseofficeIDOk

`func (o *CommonResponseErrorWrongFranchiseoffice) GetFkiFranchiseofficeIDOk() (*int32, bool)`

GetFkiFranchiseofficeIDOk returns a tuple with the FkiFranchiseofficeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiFranchiseofficeID

`func (o *CommonResponseErrorWrongFranchiseoffice) SetFkiFranchiseofficeID(v int32)`

SetFkiFranchiseofficeID sets FkiFranchiseofficeID field to given value.


### GetIFranchiseofficeCode

`func (o *CommonResponseErrorWrongFranchiseoffice) GetIFranchiseofficeCode() string`

GetIFranchiseofficeCode returns the IFranchiseofficeCode field if non-nil, zero value otherwise.

### GetIFranchiseofficeCodeOk

`func (o *CommonResponseErrorWrongFranchiseoffice) GetIFranchiseofficeCodeOk() (*string, bool)`

GetIFranchiseofficeCodeOk returns a tuple with the IFranchiseofficeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIFranchiseofficeCode

`func (o *CommonResponseErrorWrongFranchiseoffice) SetIFranchiseofficeCode(v string)`

SetIFranchiseofficeCode sets IFranchiseofficeCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


