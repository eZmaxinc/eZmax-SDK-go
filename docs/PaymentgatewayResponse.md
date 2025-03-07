# PaymentgatewayResponse

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

### NewPaymentgatewayResponse

`func NewPaymentgatewayResponse(pkiPaymentgatewayID int32, ePaymentgatewayProcessor FieldEPaymentgatewayProcessor, objPaymentgatewayDescription MultilingualPaymentgatewayDescription, ) *PaymentgatewayResponse`

NewPaymentgatewayResponse instantiates a new PaymentgatewayResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentgatewayResponseWithDefaults

`func NewPaymentgatewayResponseWithDefaults() *PaymentgatewayResponse`

NewPaymentgatewayResponseWithDefaults instantiates a new PaymentgatewayResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiPaymentgatewayID

`func (o *PaymentgatewayResponse) GetPkiPaymentgatewayID() int32`

GetPkiPaymentgatewayID returns the PkiPaymentgatewayID field if non-nil, zero value otherwise.

### GetPkiPaymentgatewayIDOk

`func (o *PaymentgatewayResponse) GetPkiPaymentgatewayIDOk() (*int32, bool)`

GetPkiPaymentgatewayIDOk returns a tuple with the PkiPaymentgatewayID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiPaymentgatewayID

`func (o *PaymentgatewayResponse) SetPkiPaymentgatewayID(v int32)`

SetPkiPaymentgatewayID sets PkiPaymentgatewayID field to given value.


### GetFkiCreditcardmerchantID

`func (o *PaymentgatewayResponse) GetFkiCreditcardmerchantID() int32`

GetFkiCreditcardmerchantID returns the FkiCreditcardmerchantID field if non-nil, zero value otherwise.

### GetFkiCreditcardmerchantIDOk

`func (o *PaymentgatewayResponse) GetFkiCreditcardmerchantIDOk() (*int32, bool)`

GetFkiCreditcardmerchantIDOk returns a tuple with the FkiCreditcardmerchantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiCreditcardmerchantID

`func (o *PaymentgatewayResponse) SetFkiCreditcardmerchantID(v int32)`

SetFkiCreditcardmerchantID sets FkiCreditcardmerchantID field to given value.

### HasFkiCreditcardmerchantID

`func (o *PaymentgatewayResponse) HasFkiCreditcardmerchantID() bool`

HasFkiCreditcardmerchantID returns a boolean if a field has been set.

### GetSCreditcardmerchantDescription

`func (o *PaymentgatewayResponse) GetSCreditcardmerchantDescription() string`

GetSCreditcardmerchantDescription returns the SCreditcardmerchantDescription field if non-nil, zero value otherwise.

### GetSCreditcardmerchantDescriptionOk

`func (o *PaymentgatewayResponse) GetSCreditcardmerchantDescriptionOk() (*string, bool)`

GetSCreditcardmerchantDescriptionOk returns a tuple with the SCreditcardmerchantDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCreditcardmerchantDescription

`func (o *PaymentgatewayResponse) SetSCreditcardmerchantDescription(v string)`

SetSCreditcardmerchantDescription sets SCreditcardmerchantDescription field to given value.

### HasSCreditcardmerchantDescription

`func (o *PaymentgatewayResponse) HasSCreditcardmerchantDescription() bool`

HasSCreditcardmerchantDescription returns a boolean if a field has been set.

### GetEPaymentgatewayProcessor

`func (o *PaymentgatewayResponse) GetEPaymentgatewayProcessor() FieldEPaymentgatewayProcessor`

GetEPaymentgatewayProcessor returns the EPaymentgatewayProcessor field if non-nil, zero value otherwise.

### GetEPaymentgatewayProcessorOk

`func (o *PaymentgatewayResponse) GetEPaymentgatewayProcessorOk() (*FieldEPaymentgatewayProcessor, bool)`

GetEPaymentgatewayProcessorOk returns a tuple with the EPaymentgatewayProcessor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPaymentgatewayProcessor

`func (o *PaymentgatewayResponse) SetEPaymentgatewayProcessor(v FieldEPaymentgatewayProcessor)`

SetEPaymentgatewayProcessor sets EPaymentgatewayProcessor field to given value.


### GetObjPaymentgatewayDescription

`func (o *PaymentgatewayResponse) GetObjPaymentgatewayDescription() MultilingualPaymentgatewayDescription`

GetObjPaymentgatewayDescription returns the ObjPaymentgatewayDescription field if non-nil, zero value otherwise.

### GetObjPaymentgatewayDescriptionOk

`func (o *PaymentgatewayResponse) GetObjPaymentgatewayDescriptionOk() (*MultilingualPaymentgatewayDescription, bool)`

GetObjPaymentgatewayDescriptionOk returns a tuple with the ObjPaymentgatewayDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjPaymentgatewayDescription

`func (o *PaymentgatewayResponse) SetObjPaymentgatewayDescription(v MultilingualPaymentgatewayDescription)`

SetObjPaymentgatewayDescription sets ObjPaymentgatewayDescription field to given value.


### GetObjCreditcardmerchant

`func (o *PaymentgatewayResponse) GetObjCreditcardmerchant() CreditcardmerchantResponseCompound`

GetObjCreditcardmerchant returns the ObjCreditcardmerchant field if non-nil, zero value otherwise.

### GetObjCreditcardmerchantOk

`func (o *PaymentgatewayResponse) GetObjCreditcardmerchantOk() (*CreditcardmerchantResponseCompound, bool)`

GetObjCreditcardmerchantOk returns a tuple with the ObjCreditcardmerchant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjCreditcardmerchant

`func (o *PaymentgatewayResponse) SetObjCreditcardmerchant(v CreditcardmerchantResponseCompound)`

SetObjCreditcardmerchant sets ObjCreditcardmerchant field to given value.

### HasObjCreditcardmerchant

`func (o *PaymentgatewayResponse) HasObjCreditcardmerchant() bool`

HasObjCreditcardmerchant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


