# PaymentgatewayRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiPaymentgatewayID** | Pointer to **int32** | The unique ID of the Paymentgateway | [optional] 
**EPaymentgatewayProcessor** | [**FieldEPaymentgatewayProcessor**](FieldEPaymentgatewayProcessor.md) |  | 
**ObjPaymentgatewayDescription** | [**MultilingualPaymentgatewayDescription**](MultilingualPaymentgatewayDescription.md) |  | 
**ObjCreditcardmerchant** | Pointer to [**CreditcardmerchantRequestCompound**](CreditcardmerchantRequestCompound.md) |  | [optional] 

## Methods

### NewPaymentgatewayRequest

`func NewPaymentgatewayRequest(ePaymentgatewayProcessor FieldEPaymentgatewayProcessor, objPaymentgatewayDescription MultilingualPaymentgatewayDescription, ) *PaymentgatewayRequest`

NewPaymentgatewayRequest instantiates a new PaymentgatewayRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentgatewayRequestWithDefaults

`func NewPaymentgatewayRequestWithDefaults() *PaymentgatewayRequest`

NewPaymentgatewayRequestWithDefaults instantiates a new PaymentgatewayRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiPaymentgatewayID

`func (o *PaymentgatewayRequest) GetPkiPaymentgatewayID() int32`

GetPkiPaymentgatewayID returns the PkiPaymentgatewayID field if non-nil, zero value otherwise.

### GetPkiPaymentgatewayIDOk

`func (o *PaymentgatewayRequest) GetPkiPaymentgatewayIDOk() (*int32, bool)`

GetPkiPaymentgatewayIDOk returns a tuple with the PkiPaymentgatewayID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiPaymentgatewayID

`func (o *PaymentgatewayRequest) SetPkiPaymentgatewayID(v int32)`

SetPkiPaymentgatewayID sets PkiPaymentgatewayID field to given value.

### HasPkiPaymentgatewayID

`func (o *PaymentgatewayRequest) HasPkiPaymentgatewayID() bool`

HasPkiPaymentgatewayID returns a boolean if a field has been set.

### GetEPaymentgatewayProcessor

`func (o *PaymentgatewayRequest) GetEPaymentgatewayProcessor() FieldEPaymentgatewayProcessor`

GetEPaymentgatewayProcessor returns the EPaymentgatewayProcessor field if non-nil, zero value otherwise.

### GetEPaymentgatewayProcessorOk

`func (o *PaymentgatewayRequest) GetEPaymentgatewayProcessorOk() (*FieldEPaymentgatewayProcessor, bool)`

GetEPaymentgatewayProcessorOk returns a tuple with the EPaymentgatewayProcessor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPaymentgatewayProcessor

`func (o *PaymentgatewayRequest) SetEPaymentgatewayProcessor(v FieldEPaymentgatewayProcessor)`

SetEPaymentgatewayProcessor sets EPaymentgatewayProcessor field to given value.


### GetObjPaymentgatewayDescription

`func (o *PaymentgatewayRequest) GetObjPaymentgatewayDescription() MultilingualPaymentgatewayDescription`

GetObjPaymentgatewayDescription returns the ObjPaymentgatewayDescription field if non-nil, zero value otherwise.

### GetObjPaymentgatewayDescriptionOk

`func (o *PaymentgatewayRequest) GetObjPaymentgatewayDescriptionOk() (*MultilingualPaymentgatewayDescription, bool)`

GetObjPaymentgatewayDescriptionOk returns a tuple with the ObjPaymentgatewayDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjPaymentgatewayDescription

`func (o *PaymentgatewayRequest) SetObjPaymentgatewayDescription(v MultilingualPaymentgatewayDescription)`

SetObjPaymentgatewayDescription sets ObjPaymentgatewayDescription field to given value.


### GetObjCreditcardmerchant

`func (o *PaymentgatewayRequest) GetObjCreditcardmerchant() CreditcardmerchantRequestCompound`

GetObjCreditcardmerchant returns the ObjCreditcardmerchant field if non-nil, zero value otherwise.

### GetObjCreditcardmerchantOk

`func (o *PaymentgatewayRequest) GetObjCreditcardmerchantOk() (*CreditcardmerchantRequestCompound, bool)`

GetObjCreditcardmerchantOk returns a tuple with the ObjCreditcardmerchant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjCreditcardmerchant

`func (o *PaymentgatewayRequest) SetObjCreditcardmerchant(v CreditcardmerchantRequestCompound)`

SetObjCreditcardmerchant sets ObjCreditcardmerchant field to given value.

### HasObjCreditcardmerchant

`func (o *PaymentgatewayRequest) HasObjCreditcardmerchant() bool`

HasObjCreditcardmerchant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


