# CreditcardmerchantListElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiCreditcardmerchantID** | **int32** | The unique ID of the Creditcardmerchant | 
**FkiBankaccountID** | **int32** | The unique ID of the Bankaccount | 
**FkiLanguageID** | Pointer to **int32** | The unique ID of the Language.  Valid values:  |Value|Description| |-|-| |1|French| |2|English| | [optional] 
**BCreditcardmerchantDenyvisa** | **bool** | Whether if visa are denied | 
**BCreditcardmerchantDenymastercard** | **bool** | Whether if mastercard are denied | 
**BCreditcardmerchantDenyamex** | **bool** | Whether if amex are denied | 
**BCreditcardmerchantIsactive** | **bool** | Whether the creditcardmerchant is active or not | 
**SCreditcardmerchantDescription** | **string** | The description of the Creditcardmerchant | 
**SCreditcardmerchantStoreid** | **string** | The storeid of the Creditcardmerchant | 

## Methods

### NewCreditcardmerchantListElement

`func NewCreditcardmerchantListElement(pkiCreditcardmerchantID int32, fkiBankaccountID int32, bCreditcardmerchantDenyvisa bool, bCreditcardmerchantDenymastercard bool, bCreditcardmerchantDenyamex bool, bCreditcardmerchantIsactive bool, sCreditcardmerchantDescription string, sCreditcardmerchantStoreid string, ) *CreditcardmerchantListElement`

NewCreditcardmerchantListElement instantiates a new CreditcardmerchantListElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditcardmerchantListElementWithDefaults

`func NewCreditcardmerchantListElementWithDefaults() *CreditcardmerchantListElement`

NewCreditcardmerchantListElementWithDefaults instantiates a new CreditcardmerchantListElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiCreditcardmerchantID

`func (o *CreditcardmerchantListElement) GetPkiCreditcardmerchantID() int32`

GetPkiCreditcardmerchantID returns the PkiCreditcardmerchantID field if non-nil, zero value otherwise.

### GetPkiCreditcardmerchantIDOk

`func (o *CreditcardmerchantListElement) GetPkiCreditcardmerchantIDOk() (*int32, bool)`

GetPkiCreditcardmerchantIDOk returns a tuple with the PkiCreditcardmerchantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiCreditcardmerchantID

`func (o *CreditcardmerchantListElement) SetPkiCreditcardmerchantID(v int32)`

SetPkiCreditcardmerchantID sets PkiCreditcardmerchantID field to given value.


### GetFkiBankaccountID

`func (o *CreditcardmerchantListElement) GetFkiBankaccountID() int32`

GetFkiBankaccountID returns the FkiBankaccountID field if non-nil, zero value otherwise.

### GetFkiBankaccountIDOk

`func (o *CreditcardmerchantListElement) GetFkiBankaccountIDOk() (*int32, bool)`

GetFkiBankaccountIDOk returns a tuple with the FkiBankaccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiBankaccountID

`func (o *CreditcardmerchantListElement) SetFkiBankaccountID(v int32)`

SetFkiBankaccountID sets FkiBankaccountID field to given value.


### GetFkiLanguageID

`func (o *CreditcardmerchantListElement) GetFkiLanguageID() int32`

GetFkiLanguageID returns the FkiLanguageID field if non-nil, zero value otherwise.

### GetFkiLanguageIDOk

`func (o *CreditcardmerchantListElement) GetFkiLanguageIDOk() (*int32, bool)`

GetFkiLanguageIDOk returns a tuple with the FkiLanguageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiLanguageID

`func (o *CreditcardmerchantListElement) SetFkiLanguageID(v int32)`

SetFkiLanguageID sets FkiLanguageID field to given value.

### HasFkiLanguageID

`func (o *CreditcardmerchantListElement) HasFkiLanguageID() bool`

HasFkiLanguageID returns a boolean if a field has been set.

### GetBCreditcardmerchantDenyvisa

`func (o *CreditcardmerchantListElement) GetBCreditcardmerchantDenyvisa() bool`

GetBCreditcardmerchantDenyvisa returns the BCreditcardmerchantDenyvisa field if non-nil, zero value otherwise.

### GetBCreditcardmerchantDenyvisaOk

`func (o *CreditcardmerchantListElement) GetBCreditcardmerchantDenyvisaOk() (*bool, bool)`

GetBCreditcardmerchantDenyvisaOk returns a tuple with the BCreditcardmerchantDenyvisa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBCreditcardmerchantDenyvisa

`func (o *CreditcardmerchantListElement) SetBCreditcardmerchantDenyvisa(v bool)`

SetBCreditcardmerchantDenyvisa sets BCreditcardmerchantDenyvisa field to given value.


### GetBCreditcardmerchantDenymastercard

`func (o *CreditcardmerchantListElement) GetBCreditcardmerchantDenymastercard() bool`

GetBCreditcardmerchantDenymastercard returns the BCreditcardmerchantDenymastercard field if non-nil, zero value otherwise.

### GetBCreditcardmerchantDenymastercardOk

`func (o *CreditcardmerchantListElement) GetBCreditcardmerchantDenymastercardOk() (*bool, bool)`

GetBCreditcardmerchantDenymastercardOk returns a tuple with the BCreditcardmerchantDenymastercard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBCreditcardmerchantDenymastercard

`func (o *CreditcardmerchantListElement) SetBCreditcardmerchantDenymastercard(v bool)`

SetBCreditcardmerchantDenymastercard sets BCreditcardmerchantDenymastercard field to given value.


### GetBCreditcardmerchantDenyamex

`func (o *CreditcardmerchantListElement) GetBCreditcardmerchantDenyamex() bool`

GetBCreditcardmerchantDenyamex returns the BCreditcardmerchantDenyamex field if non-nil, zero value otherwise.

### GetBCreditcardmerchantDenyamexOk

`func (o *CreditcardmerchantListElement) GetBCreditcardmerchantDenyamexOk() (*bool, bool)`

GetBCreditcardmerchantDenyamexOk returns a tuple with the BCreditcardmerchantDenyamex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBCreditcardmerchantDenyamex

`func (o *CreditcardmerchantListElement) SetBCreditcardmerchantDenyamex(v bool)`

SetBCreditcardmerchantDenyamex sets BCreditcardmerchantDenyamex field to given value.


### GetBCreditcardmerchantIsactive

`func (o *CreditcardmerchantListElement) GetBCreditcardmerchantIsactive() bool`

GetBCreditcardmerchantIsactive returns the BCreditcardmerchantIsactive field if non-nil, zero value otherwise.

### GetBCreditcardmerchantIsactiveOk

`func (o *CreditcardmerchantListElement) GetBCreditcardmerchantIsactiveOk() (*bool, bool)`

GetBCreditcardmerchantIsactiveOk returns a tuple with the BCreditcardmerchantIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBCreditcardmerchantIsactive

`func (o *CreditcardmerchantListElement) SetBCreditcardmerchantIsactive(v bool)`

SetBCreditcardmerchantIsactive sets BCreditcardmerchantIsactive field to given value.


### GetSCreditcardmerchantDescription

`func (o *CreditcardmerchantListElement) GetSCreditcardmerchantDescription() string`

GetSCreditcardmerchantDescription returns the SCreditcardmerchantDescription field if non-nil, zero value otherwise.

### GetSCreditcardmerchantDescriptionOk

`func (o *CreditcardmerchantListElement) GetSCreditcardmerchantDescriptionOk() (*string, bool)`

GetSCreditcardmerchantDescriptionOk returns a tuple with the SCreditcardmerchantDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCreditcardmerchantDescription

`func (o *CreditcardmerchantListElement) SetSCreditcardmerchantDescription(v string)`

SetSCreditcardmerchantDescription sets SCreditcardmerchantDescription field to given value.


### GetSCreditcardmerchantStoreid

`func (o *CreditcardmerchantListElement) GetSCreditcardmerchantStoreid() string`

GetSCreditcardmerchantStoreid returns the SCreditcardmerchantStoreid field if non-nil, zero value otherwise.

### GetSCreditcardmerchantStoreidOk

`func (o *CreditcardmerchantListElement) GetSCreditcardmerchantStoreidOk() (*string, bool)`

GetSCreditcardmerchantStoreidOk returns a tuple with the SCreditcardmerchantStoreid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCreditcardmerchantStoreid

`func (o *CreditcardmerchantListElement) SetSCreditcardmerchantStoreid(v string)`

SetSCreditcardmerchantStoreid sets SCreditcardmerchantStoreid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


