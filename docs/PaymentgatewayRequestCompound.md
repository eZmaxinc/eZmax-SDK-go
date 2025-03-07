# PaymentgatewayRequestCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiPaymentgatewayID** | Pointer to **int32** | The unique ID of the Paymentgateway | [optional] 
**EPaymentgatewayProcessor** | [**FieldEPaymentgatewayProcessor**](FieldEPaymentgatewayProcessor.md) |  | 
**ObjPaymentgatewayDescription** | [**MultilingualPaymentgatewayDescription**](MultilingualPaymentgatewayDescription.md) |  | 
**ObjCreditcardmerchant** | Pointer to [**CreditcardmerchantRequestCompound**](CreditcardmerchantRequestCompound.md) |  | [optional] 

## Methods

### NewPaymentgatewayRequestCompound

`func NewPaymentgatewayRequestCompound(ePaymentgatewayProcessor FieldEPaymentgatewayProcessor, objPaymentgatewayDescription MultilingualPaymentgatewayDescription, ) *PaymentgatewayRequestCompound`

NewPaymentgatewayRequestCompound instantiates a new PaymentgatewayRequestCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentgatewayRequestCompoundWithDefaults

`func NewPaymentgatewayRequestCompoundWithDefaults() *PaymentgatewayRequestCompound`

NewPaymentgatewayRequestCompoundWithDefaults instantiates a new PaymentgatewayRequestCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiPaymentgatewayID

`func (o *PaymentgatewayRequestCompound) GetPkiPaymentgatewayID() int32`

GetPkiPaymentgatewayID returns the PkiPaymentgatewayID field if non-nil, zero value otherwise.

### GetPkiPaymentgatewayIDOk

`func (o *PaymentgatewayRequestCompound) GetPkiPaymentgatewayIDOk() (*int32, bool)`

GetPkiPaymentgatewayIDOk returns a tuple with the PkiPaymentgatewayID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiPaymentgatewayID

`func (o *PaymentgatewayRequestCompound) SetPkiPaymentgatewayID(v int32)`

SetPkiPaymentgatewayID sets PkiPaymentgatewayID field to given value.

### HasPkiPaymentgatewayID

`func (o *PaymentgatewayRequestCompound) HasPkiPaymentgatewayID() bool`

HasPkiPaymentgatewayID returns a boolean if a field has been set.

### GetEPaymentgatewayProcessor

`func (o *PaymentgatewayRequestCompound) GetEPaymentgatewayProcessor() FieldEPaymentgatewayProcessor`

GetEPaymentgatewayProcessor returns the EPaymentgatewayProcessor field if non-nil, zero value otherwise.

### GetEPaymentgatewayProcessorOk

`func (o *PaymentgatewayRequestCompound) GetEPaymentgatewayProcessorOk() (*FieldEPaymentgatewayProcessor, bool)`

GetEPaymentgatewayProcessorOk returns a tuple with the EPaymentgatewayProcessor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPaymentgatewayProcessor

`func (o *PaymentgatewayRequestCompound) SetEPaymentgatewayProcessor(v FieldEPaymentgatewayProcessor)`

SetEPaymentgatewayProcessor sets EPaymentgatewayProcessor field to given value.


### GetObjPaymentgatewayDescription

`func (o *PaymentgatewayRequestCompound) GetObjPaymentgatewayDescription() MultilingualPaymentgatewayDescription`

GetObjPaymentgatewayDescription returns the ObjPaymentgatewayDescription field if non-nil, zero value otherwise.

### GetObjPaymentgatewayDescriptionOk

`func (o *PaymentgatewayRequestCompound) GetObjPaymentgatewayDescriptionOk() (*MultilingualPaymentgatewayDescription, bool)`

GetObjPaymentgatewayDescriptionOk returns a tuple with the ObjPaymentgatewayDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjPaymentgatewayDescription

`func (o *PaymentgatewayRequestCompound) SetObjPaymentgatewayDescription(v MultilingualPaymentgatewayDescription)`

SetObjPaymentgatewayDescription sets ObjPaymentgatewayDescription field to given value.


### GetObjCreditcardmerchant

`func (o *PaymentgatewayRequestCompound) GetObjCreditcardmerchant() CreditcardmerchantRequestCompound`

GetObjCreditcardmerchant returns the ObjCreditcardmerchant field if non-nil, zero value otherwise.

### GetObjCreditcardmerchantOk

`func (o *PaymentgatewayRequestCompound) GetObjCreditcardmerchantOk() (*CreditcardmerchantRequestCompound, bool)`

GetObjCreditcardmerchantOk returns a tuple with the ObjCreditcardmerchant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjCreditcardmerchant

`func (o *PaymentgatewayRequestCompound) SetObjCreditcardmerchant(v CreditcardmerchantRequestCompound)`

SetObjCreditcardmerchant sets ObjCreditcardmerchant field to given value.

### HasObjCreditcardmerchant

`func (o *PaymentgatewayRequestCompound) HasObjCreditcardmerchant() bool`

HasObjCreditcardmerchant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


