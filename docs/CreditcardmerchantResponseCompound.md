# CreditcardmerchantResponseCompound

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

### NewCreditcardmerchantResponseCompound

`func NewCreditcardmerchantResponseCompound(pkiCreditcardmerchantID int32, fkiBankaccountID int32, bCreditcardmerchantDenyvisa bool, bCreditcardmerchantDenymastercard bool, bCreditcardmerchantDenyamex bool, bCreditcardmerchantIsactive bool, sCreditcardmerchantDescription string, sCreditcardmerchantStoreid string, ) *CreditcardmerchantResponseCompound`

NewCreditcardmerchantResponseCompound instantiates a new CreditcardmerchantResponseCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditcardmerchantResponseCompoundWithDefaults

`func NewCreditcardmerchantResponseCompoundWithDefaults() *CreditcardmerchantResponseCompound`

NewCreditcardmerchantResponseCompoundWithDefaults instantiates a new CreditcardmerchantResponseCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiCreditcardmerchantID

`func (o *CreditcardmerchantResponseCompound) GetPkiCreditcardmerchantID() int32`

GetPkiCreditcardmerchantID returns the PkiCreditcardmerchantID field if non-nil, zero value otherwise.

### GetPkiCreditcardmerchantIDOk

`func (o *CreditcardmerchantResponseCompound) GetPkiCreditcardmerchantIDOk() (*int32, bool)`

GetPkiCreditcardmerchantIDOk returns a tuple with the PkiCreditcardmerchantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiCreditcardmerchantID

`func (o *CreditcardmerchantResponseCompound) SetPkiCreditcardmerchantID(v int32)`

SetPkiCreditcardmerchantID sets PkiCreditcardmerchantID field to given value.


### GetFkiBankaccountID

`func (o *CreditcardmerchantResponseCompound) GetFkiBankaccountID() int32`

GetFkiBankaccountID returns the FkiBankaccountID field if non-nil, zero value otherwise.

### GetFkiBankaccountIDOk

`func (o *CreditcardmerchantResponseCompound) GetFkiBankaccountIDOk() (*int32, bool)`

GetFkiBankaccountIDOk returns a tuple with the FkiBankaccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiBankaccountID

`func (o *CreditcardmerchantResponseCompound) SetFkiBankaccountID(v int32)`

SetFkiBankaccountID sets FkiBankaccountID field to given value.


### GetSBankaccountBankname

`func (o *CreditcardmerchantResponseCompound) GetSBankaccountBankname() string`

GetSBankaccountBankname returns the SBankaccountBankname field if non-nil, zero value otherwise.

### GetSBankaccountBanknameOk

`func (o *CreditcardmerchantResponseCompound) GetSBankaccountBanknameOk() (*string, bool)`

GetSBankaccountBanknameOk returns a tuple with the SBankaccountBankname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSBankaccountBankname

`func (o *CreditcardmerchantResponseCompound) SetSBankaccountBankname(v string)`

SetSBankaccountBankname sets SBankaccountBankname field to given value.

### HasSBankaccountBankname

`func (o *CreditcardmerchantResponseCompound) HasSBankaccountBankname() bool`

HasSBankaccountBankname returns a boolean if a field has been set.

### GetFkiLanguageID

`func (o *CreditcardmerchantResponseCompound) GetFkiLanguageID() int32`

GetFkiLanguageID returns the FkiLanguageID field if non-nil, zero value otherwise.

### GetFkiLanguageIDOk

`func (o *CreditcardmerchantResponseCompound) GetFkiLanguageIDOk() (*int32, bool)`

GetFkiLanguageIDOk returns a tuple with the FkiLanguageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiLanguageID

`func (o *CreditcardmerchantResponseCompound) SetFkiLanguageID(v int32)`

SetFkiLanguageID sets FkiLanguageID field to given value.

### HasFkiLanguageID

`func (o *CreditcardmerchantResponseCompound) HasFkiLanguageID() bool`

HasFkiLanguageID returns a boolean if a field has been set.

### GetSLanguageNameX

`func (o *CreditcardmerchantResponseCompound) GetSLanguageNameX() string`

GetSLanguageNameX returns the SLanguageNameX field if non-nil, zero value otherwise.

### GetSLanguageNameXOk

`func (o *CreditcardmerchantResponseCompound) GetSLanguageNameXOk() (*string, bool)`

GetSLanguageNameXOk returns a tuple with the SLanguageNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSLanguageNameX

`func (o *CreditcardmerchantResponseCompound) SetSLanguageNameX(v string)`

SetSLanguageNameX sets SLanguageNameX field to given value.

### HasSLanguageNameX

`func (o *CreditcardmerchantResponseCompound) HasSLanguageNameX() bool`

HasSLanguageNameX returns a boolean if a field has been set.

### GetBCreditcardmerchantDenyvisa

`func (o *CreditcardmerchantResponseCompound) GetBCreditcardmerchantDenyvisa() bool`

GetBCreditcardmerchantDenyvisa returns the BCreditcardmerchantDenyvisa field if non-nil, zero value otherwise.

### GetBCreditcardmerchantDenyvisaOk

`func (o *CreditcardmerchantResponseCompound) GetBCreditcardmerchantDenyvisaOk() (*bool, bool)`

GetBCreditcardmerchantDenyvisaOk returns a tuple with the BCreditcardmerchantDenyvisa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBCreditcardmerchantDenyvisa

`func (o *CreditcardmerchantResponseCompound) SetBCreditcardmerchantDenyvisa(v bool)`

SetBCreditcardmerchantDenyvisa sets BCreditcardmerchantDenyvisa field to given value.


### GetBCreditcardmerchantDenymastercard

`func (o *CreditcardmerchantResponseCompound) GetBCreditcardmerchantDenymastercard() bool`

GetBCreditcardmerchantDenymastercard returns the BCreditcardmerchantDenymastercard field if non-nil, zero value otherwise.

### GetBCreditcardmerchantDenymastercardOk

`func (o *CreditcardmerchantResponseCompound) GetBCreditcardmerchantDenymastercardOk() (*bool, bool)`

GetBCreditcardmerchantDenymastercardOk returns a tuple with the BCreditcardmerchantDenymastercard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBCreditcardmerchantDenymastercard

`func (o *CreditcardmerchantResponseCompound) SetBCreditcardmerchantDenymastercard(v bool)`

SetBCreditcardmerchantDenymastercard sets BCreditcardmerchantDenymastercard field to given value.


### GetBCreditcardmerchantDenyamex

`func (o *CreditcardmerchantResponseCompound) GetBCreditcardmerchantDenyamex() bool`

GetBCreditcardmerchantDenyamex returns the BCreditcardmerchantDenyamex field if non-nil, zero value otherwise.

### GetBCreditcardmerchantDenyamexOk

`func (o *CreditcardmerchantResponseCompound) GetBCreditcardmerchantDenyamexOk() (*bool, bool)`

GetBCreditcardmerchantDenyamexOk returns a tuple with the BCreditcardmerchantDenyamex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBCreditcardmerchantDenyamex

`func (o *CreditcardmerchantResponseCompound) SetBCreditcardmerchantDenyamex(v bool)`

SetBCreditcardmerchantDenyamex sets BCreditcardmerchantDenyamex field to given value.


### GetBCreditcardmerchantIsactive

`func (o *CreditcardmerchantResponseCompound) GetBCreditcardmerchantIsactive() bool`

GetBCreditcardmerchantIsactive returns the BCreditcardmerchantIsactive field if non-nil, zero value otherwise.

### GetBCreditcardmerchantIsactiveOk

`func (o *CreditcardmerchantResponseCompound) GetBCreditcardmerchantIsactiveOk() (*bool, bool)`

GetBCreditcardmerchantIsactiveOk returns a tuple with the BCreditcardmerchantIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBCreditcardmerchantIsactive

`func (o *CreditcardmerchantResponseCompound) SetBCreditcardmerchantIsactive(v bool)`

SetBCreditcardmerchantIsactive sets BCreditcardmerchantIsactive field to given value.


### GetSCreditcardmerchantDescription

`func (o *CreditcardmerchantResponseCompound) GetSCreditcardmerchantDescription() string`

GetSCreditcardmerchantDescription returns the SCreditcardmerchantDescription field if non-nil, zero value otherwise.

### GetSCreditcardmerchantDescriptionOk

`func (o *CreditcardmerchantResponseCompound) GetSCreditcardmerchantDescriptionOk() (*string, bool)`

GetSCreditcardmerchantDescriptionOk returns a tuple with the SCreditcardmerchantDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCreditcardmerchantDescription

`func (o *CreditcardmerchantResponseCompound) SetSCreditcardmerchantDescription(v string)`

SetSCreditcardmerchantDescription sets SCreditcardmerchantDescription field to given value.


### GetSCreditcardmerchantStoreid

`func (o *CreditcardmerchantResponseCompound) GetSCreditcardmerchantStoreid() string`

GetSCreditcardmerchantStoreid returns the SCreditcardmerchantStoreid field if non-nil, zero value otherwise.

### GetSCreditcardmerchantStoreidOk

`func (o *CreditcardmerchantResponseCompound) GetSCreditcardmerchantStoreidOk() (*string, bool)`

GetSCreditcardmerchantStoreidOk returns a tuple with the SCreditcardmerchantStoreid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCreditcardmerchantStoreid

`func (o *CreditcardmerchantResponseCompound) SetSCreditcardmerchantStoreid(v string)`

SetSCreditcardmerchantStoreid sets SCreditcardmerchantStoreid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


