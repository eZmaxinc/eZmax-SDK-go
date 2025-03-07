# PaymentgatewayResponseCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiPaymentgatewayID** | **int32** | The unique ID of the Paymentgateway | 
**FkiCreditcardmerchantID** | Pointer to **int32** | The unique ID of the Creditcardmerchant | [optional] 
**SCreditcardmerchantDescription** | Pointer to **string** | The description of the Creditcardmerchant | [optional] 
**EPaymentgatewayProcessor** | [**FieldEPaymentgatewayProcessor**](FieldEPaymentgatewayProcessor.md) |  | 
**ObjPaymentgatewayDescription** | [**MultilingualPaymentgatewayDescription**](MultilingualPaymentgatewayDescription.md) |  | 
**ObjCreditcardmerchant** | Pointer to [**CreditcardmerchantResponseCompound**](CreditcardmerchantResponseCompound.md) |  | [optional] 

## Methods

### NewPaymentgatewayResponseCompound

`func NewPaymentgatewayResponseCompound(pkiPaymentgatewayID int32, ePaymentgatewayProcessor FieldEPaymentgatewayProcessor, objPaymentgatewayDescription MultilingualPaymentgatewayDescription, ) *PaymentgatewayResponseCompound`

NewPaymentgatewayResponseCompound instantiates a new PaymentgatewayResponseCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentgatewayResponseCompoundWithDefaults

`func NewPaymentgatewayResponseCompoundWithDefaults() *PaymentgatewayResponseCompound`

NewPaymentgatewayResponseCompoundWithDefaults instantiates a new PaymentgatewayResponseCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiPaymentgatewayID

`func (o *PaymentgatewayResponseCompound) GetPkiPaymentgatewayID() int32`

GetPkiPaymentgatewayID returns the PkiPaymentgatewayID field if non-nil, zero value otherwise.

### GetPkiPaymentgatewayIDOk

`func (o *PaymentgatewayResponseCompound) GetPkiPaymentgatewayIDOk() (*int32, bool)`

GetPkiPaymentgatewayIDOk returns a tuple with the PkiPaymentgatewayID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiPaymentgatewayID

`func (o *PaymentgatewayResponseCompound) SetPkiPaymentgatewayID(v int32)`

SetPkiPaymentgatewayID sets PkiPaymentgatewayID field to given value.


### GetFkiCreditcardmerchantID

`func (o *PaymentgatewayResponseCompound) GetFkiCreditcardmerchantID() int32`

GetFkiCreditcardmerchantID returns the FkiCreditcardmerchantID field if non-nil, zero value otherwise.

### GetFkiCreditcardmerchantIDOk

`func (o *PaymentgatewayResponseCompound) GetFkiCreditcardmerchantIDOk() (*int32, bool)`

GetFkiCreditcardmerchantIDOk returns a tuple with the FkiCreditcardmerchantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiCreditcardmerchantID

`func (o *PaymentgatewayResponseCompound) SetFkiCreditcardmerchantID(v int32)`

SetFkiCreditcardmerchantID sets FkiCreditcardmerchantID field to given value.

### HasFkiCreditcardmerchantID

`func (o *PaymentgatewayResponseCompound) HasFkiCreditcardmerchantID() bool`

HasFkiCreditcardmerchantID returns a boolean if a field has been set.

### GetSCreditcardmerchantDescription

`func (o *PaymentgatewayResponseCompound) GetSCreditcardmerchantDescription() string`

GetSCreditcardmerchantDescription returns the SCreditcardmerchantDescription field if non-nil, zero value otherwise.

### GetSCreditcardmerchantDescriptionOk

`func (o *PaymentgatewayResponseCompound) GetSCreditcardmerchantDescriptionOk() (*string, bool)`

GetSCreditcardmerchantDescriptionOk returns a tuple with the SCreditcardmerchantDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCreditcardmerchantDescription

`func (o *PaymentgatewayResponseCompound) SetSCreditcardmerchantDescription(v string)`

SetSCreditcardmerchantDescription sets SCreditcardmerchantDescription field to given value.

### HasSCreditcardmerchantDescription

`func (o *PaymentgatewayResponseCompound) HasSCreditcardmerchantDescription() bool`

HasSCreditcardmerchantDescription returns a boolean if a field has been set.

### GetEPaymentgatewayProcessor

`func (o *PaymentgatewayResponseCompound) GetEPaymentgatewayProcessor() FieldEPaymentgatewayProcessor`

GetEPaymentgatewayProcessor returns the EPaymentgatewayProcessor field if non-nil, zero value otherwise.

### GetEPaymentgatewayProcessorOk

`func (o *PaymentgatewayResponseCompound) GetEPaymentgatewayProcessorOk() (*FieldEPaymentgatewayProcessor, bool)`

GetEPaymentgatewayProcessorOk returns a tuple with the EPaymentgatewayProcessor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPaymentgatewayProcessor

`func (o *PaymentgatewayResponseCompound) SetEPaymentgatewayProcessor(v FieldEPaymentgatewayProcessor)`

SetEPaymentgatewayProcessor sets EPaymentgatewayProcessor field to given value.


### GetObjPaymentgatewayDescription

`func (o *PaymentgatewayResponseCompound) GetObjPaymentgatewayDescription() MultilingualPaymentgatewayDescription`

GetObjPaymentgatewayDescription returns the ObjPaymentgatewayDescription field if non-nil, zero value otherwise.

### GetObjPaymentgatewayDescriptionOk

`func (o *PaymentgatewayResponseCompound) GetObjPaymentgatewayDescriptionOk() (*MultilingualPaymentgatewayDescription, bool)`

GetObjPaymentgatewayDescriptionOk returns a tuple with the ObjPaymentgatewayDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjPaymentgatewayDescription

`func (o *PaymentgatewayResponseCompound) SetObjPaymentgatewayDescription(v MultilingualPaymentgatewayDescription)`

SetObjPaymentgatewayDescription sets ObjPaymentgatewayDescription field to given value.


### GetObjCreditcardmerchant

`func (o *PaymentgatewayResponseCompound) GetObjCreditcardmerchant() CreditcardmerchantResponseCompound`

GetObjCreditcardmerchant returns the ObjCreditcardmerchant field if non-nil, zero value otherwise.

### GetObjCreditcardmerchantOk

`func (o *PaymentgatewayResponseCompound) GetObjCreditcardmerchantOk() (*CreditcardmerchantResponseCompound, bool)`

GetObjCreditcardmerchantOk returns a tuple with the ObjCreditcardmerchant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjCreditcardmerchant

`func (o *PaymentgatewayResponseCompound) SetObjCreditcardmerchant(v CreditcardmerchantResponseCompound)`

SetObjCreditcardmerchant sets ObjCreditcardmerchant field to given value.

### HasObjCreditcardmerchant

`func (o *PaymentgatewayResponseCompound) HasObjCreditcardmerchant() bool`

HasObjCreditcardmerchant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


