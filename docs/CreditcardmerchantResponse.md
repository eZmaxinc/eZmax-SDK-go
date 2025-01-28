# CreditcardmerchantResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiCreditcardmerchantID** | **int32** | The unique ID of the Creditcardmerchant | 
**FkiBankaccountID** | **int32** | The unique ID of the Bankaccount | 
**SBankaccountBankname** | Pointer to **string** | The name of the bank | [optional] 
**FkiLanguageID** | Pointer to **int32** | The unique ID of the Language.  Valid values:  |Value|Description| |-|-| |1|French| |2|English| | [optional] 
**SLanguageNameX** | Pointer to **string** | The Name of the Language in the language of the requester | [optional] 
**BCreditcardmerchantDenyvisa** | **bool** | Whether if visa are denied | 
**BCreditcardmerchantDenymastercard** | **bool** | Whether if mastercard are denied | 
**BCreditcardmerchantDenyamex** | **bool** | Whether if amex are denied | 
**BCreditcardmerchantIsactive** | **bool** | Whether the creditcardmerchant is active or not | 
**SCreditcardmerchantDescription** | **string** | The description of the Creditcardmerchant | 
**SCreditcardmerchantStoreid** | **string** | The storeid of the Creditcardmerchant | 

## Methods

### NewCreditcardmerchantResponse

`func NewCreditcardmerchantResponse(pkiCreditcardmerchantID int32, fkiBankaccountID int32, bCreditcardmerchantDenyvisa bool, bCreditcardmerchantDenymastercard bool, bCreditcardmerchantDenyamex bool, bCreditcardmerchantIsactive bool, sCreditcardmerchantDescription string, sCreditcardmerchantStoreid string, ) *CreditcardmerchantResponse`

NewCreditcardmerchantResponse instantiates a new CreditcardmerchantResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditcardmerchantResponseWithDefaults

`func NewCreditcardmerchantResponseWithDefaults() *CreditcardmerchantResponse`

NewCreditcardmerchantResponseWithDefaults instantiates a new CreditcardmerchantResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiCreditcardmerchantID

`func (o *CreditcardmerchantResponse) GetPkiCreditcardmerchantID() int32`

GetPkiCreditcardmerchantID returns the PkiCreditcardmerchantID field if non-nil, zero value otherwise.

### GetPkiCreditcardmerchantIDOk

`func (o *CreditcardmerchantResponse) GetPkiCreditcardmerchantIDOk() (*int32, bool)`

GetPkiCreditcardmerchantIDOk returns a tuple with the PkiCreditcardmerchantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiCreditcardmerchantID

`func (o *CreditcardmerchantResponse) SetPkiCreditcardmerchantID(v int32)`

SetPkiCreditcardmerchantID sets PkiCreditcardmerchantID field to given value.


### GetFkiBankaccountID

`func (o *CreditcardmerchantResponse) GetFkiBankaccountID() int32`

GetFkiBankaccountID returns the FkiBankaccountID field if non-nil, zero value otherwise.

### GetFkiBankaccountIDOk

`func (o *CreditcardmerchantResponse) GetFkiBankaccountIDOk() (*int32, bool)`

GetFkiBankaccountIDOk returns a tuple with the FkiBankaccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiBankaccountID

`func (o *CreditcardmerchantResponse) SetFkiBankaccountID(v int32)`

SetFkiBankaccountID sets FkiBankaccountID field to given value.


### GetSBankaccountBankname

`func (o *CreditcardmerchantResponse) GetSBankaccountBankname() string`

GetSBankaccountBankname returns the SBankaccountBankname field if non-nil, zero value otherwise.

### GetSBankaccountBanknameOk

`func (o *CreditcardmerchantResponse) GetSBankaccountBanknameOk() (*string, bool)`

GetSBankaccountBanknameOk returns a tuple with the SBankaccountBankname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSBankaccountBankname

`func (o *CreditcardmerchantResponse) SetSBankaccountBankname(v string)`

SetSBankaccountBankname sets SBankaccountBankname field to given value.

### HasSBankaccountBankname

`func (o *CreditcardmerchantResponse) HasSBankaccountBankname() bool`

HasSBankaccountBankname returns a boolean if a field has been set.

### GetFkiLanguageID

`func (o *CreditcardmerchantResponse) GetFkiLanguageID() int32`

GetFkiLanguageID returns the FkiLanguageID field if non-nil, zero value otherwise.

### GetFkiLanguageIDOk

`func (o *CreditcardmerchantResponse) GetFkiLanguageIDOk() (*int32, bool)`

GetFkiLanguageIDOk returns a tuple with the FkiLanguageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiLanguageID

`func (o *CreditcardmerchantResponse) SetFkiLanguageID(v int32)`

SetFkiLanguageID sets FkiLanguageID field to given value.

### HasFkiLanguageID

`func (o *CreditcardmerchantResponse) HasFkiLanguageID() bool`

HasFkiLanguageID returns a boolean if a field has been set.

### GetSLanguageNameX

`func (o *CreditcardmerchantResponse) GetSLanguageNameX() string`

GetSLanguageNameX returns the SLanguageNameX field if non-nil, zero value otherwise.

### GetSLanguageNameXOk

`func (o *CreditcardmerchantResponse) GetSLanguageNameXOk() (*string, bool)`

GetSLanguageNameXOk returns a tuple with the SLanguageNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSLanguageNameX

`func (o *CreditcardmerchantResponse) SetSLanguageNameX(v string)`

SetSLanguageNameX sets SLanguageNameX field to given value.

### HasSLanguageNameX

`func (o *CreditcardmerchantResponse) HasSLanguageNameX() bool`

HasSLanguageNameX returns a boolean if a field has been set.

### GetBCreditcardmerchantDenyvisa

`func (o *CreditcardmerchantResponse) GetBCreditcardmerchantDenyvisa() bool`

GetBCreditcardmerchantDenyvisa returns the BCreditcardmerchantDenyvisa field if non-nil, zero value otherwise.

### GetBCreditcardmerchantDenyvisaOk

`func (o *CreditcardmerchantResponse) GetBCreditcardmerchantDenyvisaOk() (*bool, bool)`

GetBCreditcardmerchantDenyvisaOk returns a tuple with the BCreditcardmerchantDenyvisa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBCreditcardmerchantDenyvisa

`func (o *CreditcardmerchantResponse) SetBCreditcardmerchantDenyvisa(v bool)`

SetBCreditcardmerchantDenyvisa sets BCreditcardmerchantDenyvisa field to given value.


### GetBCreditcardmerchantDenymastercard

`func (o *CreditcardmerchantResponse) GetBCreditcardmerchantDenymastercard() bool`

GetBCreditcardmerchantDenymastercard returns the BCreditcardmerchantDenymastercard field if non-nil, zero value otherwise.

### GetBCreditcardmerchantDenymastercardOk

`func (o *CreditcardmerchantResponse) GetBCreditcardmerchantDenymastercardOk() (*bool, bool)`

GetBCreditcardmerchantDenymastercardOk returns a tuple with the BCreditcardmerchantDenymastercard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBCreditcardmerchantDenymastercard

`func (o *CreditcardmerchantResponse) SetBCreditcardmerchantDenymastercard(v bool)`

SetBCreditcardmerchantDenymastercard sets BCreditcardmerchantDenymastercard field to given value.


### GetBCreditcardmerchantDenyamex

`func (o *CreditcardmerchantResponse) GetBCreditcardmerchantDenyamex() bool`

GetBCreditcardmerchantDenyamex returns the BCreditcardmerchantDenyamex field if non-nil, zero value otherwise.

### GetBCreditcardmerchantDenyamexOk

`func (o *CreditcardmerchantResponse) GetBCreditcardmerchantDenyamexOk() (*bool, bool)`

GetBCreditcardmerchantDenyamexOk returns a tuple with the BCreditcardmerchantDenyamex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBCreditcardmerchantDenyamex

`func (o *CreditcardmerchantResponse) SetBCreditcardmerchantDenyamex(v bool)`

SetBCreditcardmerchantDenyamex sets BCreditcardmerchantDenyamex field to given value.


### GetBCreditcardmerchantIsactive

`func (o *CreditcardmerchantResponse) GetBCreditcardmerchantIsactive() bool`

GetBCreditcardmerchantIsactive returns the BCreditcardmerchantIsactive field if non-nil, zero value otherwise.

### GetBCreditcardmerchantIsactiveOk

`func (o *CreditcardmerchantResponse) GetBCreditcardmerchantIsactiveOk() (*bool, bool)`

GetBCreditcardmerchantIsactiveOk returns a tuple with the BCreditcardmerchantIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBCreditcardmerchantIsactive

`func (o *CreditcardmerchantResponse) SetBCreditcardmerchantIsactive(v bool)`

SetBCreditcardmerchantIsactive sets BCreditcardmerchantIsactive field to given value.


### GetSCreditcardmerchantDescription

`func (o *CreditcardmerchantResponse) GetSCreditcardmerchantDescription() string`

GetSCreditcardmerchantDescription returns the SCreditcardmerchantDescription field if non-nil, zero value otherwise.

### GetSCreditcardmerchantDescriptionOk

`func (o *CreditcardmerchantResponse) GetSCreditcardmerchantDescriptionOk() (*string, bool)`

GetSCreditcardmerchantDescriptionOk returns a tuple with the SCreditcardmerchantDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCreditcardmerchantDescription

`func (o *CreditcardmerchantResponse) SetSCreditcardmerchantDescription(v string)`

SetSCreditcardmerchantDescription sets SCreditcardmerchantDescription field to given value.


### GetSCreditcardmerchantStoreid

`func (o *CreditcardmerchantResponse) GetSCreditcardmerchantStoreid() string`

GetSCreditcardmerchantStoreid returns the SCreditcardmerchantStoreid field if non-nil, zero value otherwise.

### GetSCreditcardmerchantStoreidOk

`func (o *CreditcardmerchantResponse) GetSCreditcardmerchantStoreidOk() (*string, bool)`

GetSCreditcardmerchantStoreidOk returns a tuple with the SCreditcardmerchantStoreid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCreditcardmerchantStoreid

`func (o *CreditcardmerchantResponse) SetSCreditcardmerchantStoreid(v string)`

SetSCreditcardmerchantStoreid sets SCreditcardmerchantStoreid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


