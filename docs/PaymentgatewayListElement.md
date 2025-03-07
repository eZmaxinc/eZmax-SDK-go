# PaymentgatewayListElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiPaymentgatewayID** | **int32** | The unique ID of the Paymentgateway | 
**FkiCreditcardmerchantID** | **int32** | The unique ID of the Creditcardmerchant | 
**EPaymentgatewayProcessor** | [**FieldEPaymentgatewayProcessor**](FieldEPaymentgatewayProcessor.md) |  | 
**SPaymentgatewayDescriptionX** | **string** | The description of the Paymentgateway in the language of the requester | 

## Methods

### NewPaymentgatewayListElement

`func NewPaymentgatewayListElement(pkiPaymentgatewayID int32, fkiCreditcardmerchantID int32, ePaymentgatewayProcessor FieldEPaymentgatewayProcessor, sPaymentgatewayDescriptionX string, ) *PaymentgatewayListElement`

NewPaymentgatewayListElement instantiates a new PaymentgatewayListElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentgatewayListElementWithDefaults

`func NewPaymentgatewayListElementWithDefaults() *PaymentgatewayListElement`

NewPaymentgatewayListElementWithDefaults instantiates a new PaymentgatewayListElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiPaymentgatewayID

`func (o *PaymentgatewayListElement) GetPkiPaymentgatewayID() int32`

GetPkiPaymentgatewayID returns the PkiPaymentgatewayID field if non-nil, zero value otherwise.

### GetPkiPaymentgatewayIDOk

`func (o *PaymentgatewayListElement) GetPkiPaymentgatewayIDOk() (*int32, bool)`

GetPkiPaymentgatewayIDOk returns a tuple with the PkiPaymentgatewayID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiPaymentgatewayID

`func (o *PaymentgatewayListElement) SetPkiPaymentgatewayID(v int32)`

SetPkiPaymentgatewayID sets PkiPaymentgatewayID field to given value.


### GetFkiCreditcardmerchantID

`func (o *PaymentgatewayListElement) GetFkiCreditcardmerchantID() int32`

GetFkiCreditcardmerchantID returns the FkiCreditcardmerchantID field if non-nil, zero value otherwise.

### GetFkiCreditcardmerchantIDOk

`func (o *PaymentgatewayListElement) GetFkiCreditcardmerchantIDOk() (*int32, bool)`

GetFkiCreditcardmerchantIDOk returns a tuple with the FkiCreditcardmerchantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiCreditcardmerchantID

`func (o *PaymentgatewayListElement) SetFkiCreditcardmerchantID(v int32)`

SetFkiCreditcardmerchantID sets FkiCreditcardmerchantID field to given value.


### GetEPaymentgatewayProcessor

`func (o *PaymentgatewayListElement) GetEPaymentgatewayProcessor() FieldEPaymentgatewayProcessor`

GetEPaymentgatewayProcessor returns the EPaymentgatewayProcessor field if non-nil, zero value otherwise.

### GetEPaymentgatewayProcessorOk

`func (o *PaymentgatewayListElement) GetEPaymentgatewayProcessorOk() (*FieldEPaymentgatewayProcessor, bool)`

GetEPaymentgatewayProcessorOk returns a tuple with the EPaymentgatewayProcessor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPaymentgatewayProcessor

`func (o *PaymentgatewayListElement) SetEPaymentgatewayProcessor(v FieldEPaymentgatewayProcessor)`

SetEPaymentgatewayProcessor sets EPaymentgatewayProcessor field to given value.


### GetSPaymentgatewayDescriptionX

`func (o *PaymentgatewayListElement) GetSPaymentgatewayDescriptionX() string`

GetSPaymentgatewayDescriptionX returns the SPaymentgatewayDescriptionX field if non-nil, zero value otherwise.

### GetSPaymentgatewayDescriptionXOk

`func (o *PaymentgatewayListElement) GetSPaymentgatewayDescriptionXOk() (*string, bool)`

GetSPaymentgatewayDescriptionXOk returns a tuple with the SPaymentgatewayDescriptionX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSPaymentgatewayDescriptionX

`func (o *PaymentgatewayListElement) SetSPaymentgatewayDescriptionX(v string)`

SetSPaymentgatewayDescriptionX sets SPaymentgatewayDescriptionX field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


