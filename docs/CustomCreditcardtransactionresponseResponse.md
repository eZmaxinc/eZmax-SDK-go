# CustomCreditcardtransactionresponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SCreditcardtransactionISOcode** | **string** | The ISO code | 
**SCreditcardtransactionResponsecode** | **string** | The response code | 
**SCreditcardtransactionResponseterminalmessage** | **string** | The terminal response message | 
**ECreditcardtransactionAvsresult** | Pointer to [**FieldECreditcardtransactionAvsresult**](FieldECreditcardtransactionAvsresult.md) |  | [optional] 
**ECreditcardtransactionCvdresult** | Pointer to [**FieldECreditcardtransactionCvdresult**](FieldECreditcardtransactionCvdresult.md) |  | [optional] 

## Methods

### NewCustomCreditcardtransactionresponseResponse

`func NewCustomCreditcardtransactionresponseResponse(sCreditcardtransactionISOcode string, sCreditcardtransactionResponsecode string, sCreditcardtransactionResponseterminalmessage string, ) *CustomCreditcardtransactionresponseResponse`

NewCustomCreditcardtransactionresponseResponse instantiates a new CustomCreditcardtransactionresponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomCreditcardtransactionresponseResponseWithDefaults

`func NewCustomCreditcardtransactionresponseResponseWithDefaults() *CustomCreditcardtransactionresponseResponse`

NewCustomCreditcardtransactionresponseResponseWithDefaults instantiates a new CustomCreditcardtransactionresponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSCreditcardtransactionISOcode

`func (o *CustomCreditcardtransactionresponseResponse) GetSCreditcardtransactionISOcode() string`

GetSCreditcardtransactionISOcode returns the SCreditcardtransactionISOcode field if non-nil, zero value otherwise.

### GetSCreditcardtransactionISOcodeOk

`func (o *CustomCreditcardtransactionresponseResponse) GetSCreditcardtransactionISOcodeOk() (*string, bool)`

GetSCreditcardtransactionISOcodeOk returns a tuple with the SCreditcardtransactionISOcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCreditcardtransactionISOcode

`func (o *CustomCreditcardtransactionresponseResponse) SetSCreditcardtransactionISOcode(v string)`

SetSCreditcardtransactionISOcode sets SCreditcardtransactionISOcode field to given value.


### GetSCreditcardtransactionResponsecode

`func (o *CustomCreditcardtransactionresponseResponse) GetSCreditcardtransactionResponsecode() string`

GetSCreditcardtransactionResponsecode returns the SCreditcardtransactionResponsecode field if non-nil, zero value otherwise.

### GetSCreditcardtransactionResponsecodeOk

`func (o *CustomCreditcardtransactionresponseResponse) GetSCreditcardtransactionResponsecodeOk() (*string, bool)`

GetSCreditcardtransactionResponsecodeOk returns a tuple with the SCreditcardtransactionResponsecode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCreditcardtransactionResponsecode

`func (o *CustomCreditcardtransactionresponseResponse) SetSCreditcardtransactionResponsecode(v string)`

SetSCreditcardtransactionResponsecode sets SCreditcardtransactionResponsecode field to given value.


### GetSCreditcardtransactionResponseterminalmessage

`func (o *CustomCreditcardtransactionresponseResponse) GetSCreditcardtransactionResponseterminalmessage() string`

GetSCreditcardtransactionResponseterminalmessage returns the SCreditcardtransactionResponseterminalmessage field if non-nil, zero value otherwise.

### GetSCreditcardtransactionResponseterminalmessageOk

`func (o *CustomCreditcardtransactionresponseResponse) GetSCreditcardtransactionResponseterminalmessageOk() (*string, bool)`

GetSCreditcardtransactionResponseterminalmessageOk returns a tuple with the SCreditcardtransactionResponseterminalmessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCreditcardtransactionResponseterminalmessage

`func (o *CustomCreditcardtransactionresponseResponse) SetSCreditcardtransactionResponseterminalmessage(v string)`

SetSCreditcardtransactionResponseterminalmessage sets SCreditcardtransactionResponseterminalmessage field to given value.


### GetECreditcardtransactionAvsresult

`func (o *CustomCreditcardtransactionresponseResponse) GetECreditcardtransactionAvsresult() FieldECreditcardtransactionAvsresult`

GetECreditcardtransactionAvsresult returns the ECreditcardtransactionAvsresult field if non-nil, zero value otherwise.

### GetECreditcardtransactionAvsresultOk

`func (o *CustomCreditcardtransactionresponseResponse) GetECreditcardtransactionAvsresultOk() (*FieldECreditcardtransactionAvsresult, bool)`

GetECreditcardtransactionAvsresultOk returns a tuple with the ECreditcardtransactionAvsresult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetECreditcardtransactionAvsresult

`func (o *CustomCreditcardtransactionresponseResponse) SetECreditcardtransactionAvsresult(v FieldECreditcardtransactionAvsresult)`

SetECreditcardtransactionAvsresult sets ECreditcardtransactionAvsresult field to given value.

### HasECreditcardtransactionAvsresult

`func (o *CustomCreditcardtransactionresponseResponse) HasECreditcardtransactionAvsresult() bool`

HasECreditcardtransactionAvsresult returns a boolean if a field has been set.

### GetECreditcardtransactionCvdresult

`func (o *CustomCreditcardtransactionresponseResponse) GetECreditcardtransactionCvdresult() FieldECreditcardtransactionCvdresult`

GetECreditcardtransactionCvdresult returns the ECreditcardtransactionCvdresult field if non-nil, zero value otherwise.

### GetECreditcardtransactionCvdresultOk

`func (o *CustomCreditcardtransactionresponseResponse) GetECreditcardtransactionCvdresultOk() (*FieldECreditcardtransactionCvdresult, bool)`

GetECreditcardtransactionCvdresultOk returns a tuple with the ECreditcardtransactionCvdresult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetECreditcardtransactionCvdresult

`func (o *CustomCreditcardtransactionresponseResponse) SetECreditcardtransactionCvdresult(v FieldECreditcardtransactionCvdresult)`

SetECreditcardtransactionCvdresult sets ECreditcardtransactionCvdresult field to given value.

### HasECreditcardtransactionCvdresult

`func (o *CustomCreditcardtransactionresponseResponse) HasECreditcardtransactionCvdresult() bool`

HasECreditcardtransactionCvdresult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


