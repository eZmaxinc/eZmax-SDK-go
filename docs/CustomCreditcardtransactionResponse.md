# CustomCreditcardtransactionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ECreditcardtypeCodename** | Pointer to [**FieldECreditcardtypeCodename**](FieldECreditcardtypeCodename.md) |  | [optional] 
**DCreditcardtransactionAmount** | **string** | The amount of the Creditcardtransaction | 
**SCreditcardtransactionPartiallydecryptednumber** | **string** | The partially decrypted credit card number used in the Creditcardtransaction | 
**SCreditcardtransactionReferencenumber** | **string** | The reference number on the creditcard service for the Creditcardtransaction | 

## Methods

### NewCustomCreditcardtransactionResponse

`func NewCustomCreditcardtransactionResponse(dCreditcardtransactionAmount string, sCreditcardtransactionPartiallydecryptednumber string, sCreditcardtransactionReferencenumber string, ) *CustomCreditcardtransactionResponse`

NewCustomCreditcardtransactionResponse instantiates a new CustomCreditcardtransactionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomCreditcardtransactionResponseWithDefaults

`func NewCustomCreditcardtransactionResponseWithDefaults() *CustomCreditcardtransactionResponse`

NewCustomCreditcardtransactionResponseWithDefaults instantiates a new CustomCreditcardtransactionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetECreditcardtypeCodename

`func (o *CustomCreditcardtransactionResponse) GetECreditcardtypeCodename() FieldECreditcardtypeCodename`

GetECreditcardtypeCodename returns the ECreditcardtypeCodename field if non-nil, zero value otherwise.

### GetECreditcardtypeCodenameOk

`func (o *CustomCreditcardtransactionResponse) GetECreditcardtypeCodenameOk() (*FieldECreditcardtypeCodename, bool)`

GetECreditcardtypeCodenameOk returns a tuple with the ECreditcardtypeCodename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetECreditcardtypeCodename

`func (o *CustomCreditcardtransactionResponse) SetECreditcardtypeCodename(v FieldECreditcardtypeCodename)`

SetECreditcardtypeCodename sets ECreditcardtypeCodename field to given value.

### HasECreditcardtypeCodename

`func (o *CustomCreditcardtransactionResponse) HasECreditcardtypeCodename() bool`

HasECreditcardtypeCodename returns a boolean if a field has been set.

### GetDCreditcardtransactionAmount

`func (o *CustomCreditcardtransactionResponse) GetDCreditcardtransactionAmount() string`

GetDCreditcardtransactionAmount returns the DCreditcardtransactionAmount field if non-nil, zero value otherwise.

### GetDCreditcardtransactionAmountOk

`func (o *CustomCreditcardtransactionResponse) GetDCreditcardtransactionAmountOk() (*string, bool)`

GetDCreditcardtransactionAmountOk returns a tuple with the DCreditcardtransactionAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDCreditcardtransactionAmount

`func (o *CustomCreditcardtransactionResponse) SetDCreditcardtransactionAmount(v string)`

SetDCreditcardtransactionAmount sets DCreditcardtransactionAmount field to given value.


### GetSCreditcardtransactionPartiallydecryptednumber

`func (o *CustomCreditcardtransactionResponse) GetSCreditcardtransactionPartiallydecryptednumber() string`

GetSCreditcardtransactionPartiallydecryptednumber returns the SCreditcardtransactionPartiallydecryptednumber field if non-nil, zero value otherwise.

### GetSCreditcardtransactionPartiallydecryptednumberOk

`func (o *CustomCreditcardtransactionResponse) GetSCreditcardtransactionPartiallydecryptednumberOk() (*string, bool)`

GetSCreditcardtransactionPartiallydecryptednumberOk returns a tuple with the SCreditcardtransactionPartiallydecryptednumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCreditcardtransactionPartiallydecryptednumber

`func (o *CustomCreditcardtransactionResponse) SetSCreditcardtransactionPartiallydecryptednumber(v string)`

SetSCreditcardtransactionPartiallydecryptednumber sets SCreditcardtransactionPartiallydecryptednumber field to given value.


### GetSCreditcardtransactionReferencenumber

`func (o *CustomCreditcardtransactionResponse) GetSCreditcardtransactionReferencenumber() string`

GetSCreditcardtransactionReferencenumber returns the SCreditcardtransactionReferencenumber field if non-nil, zero value otherwise.

### GetSCreditcardtransactionReferencenumberOk

`func (o *CustomCreditcardtransactionResponse) GetSCreditcardtransactionReferencenumberOk() (*string, bool)`

GetSCreditcardtransactionReferencenumberOk returns a tuple with the SCreditcardtransactionReferencenumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCreditcardtransactionReferencenumber

`func (o *CustomCreditcardtransactionResponse) SetSCreditcardtransactionReferencenumber(v string)`

SetSCreditcardtransactionReferencenumber sets SCreditcardtransactionReferencenumber field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


