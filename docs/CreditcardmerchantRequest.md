# CreditcardmerchantRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiCreditcardmerchantID** | Pointer to **int32** | The unique ID of the Creditcardmerchant | [optional] 
**FkiBankaccountID** | **int32** | The unique ID of the Bankaccount | 
**FkiLanguageID** | Pointer to **int32** | The unique ID of the Language.  Valid values:  |Value|Description| |-|-| |1|French| |2|English| | [optional] 
**BCreditcardmerchantDenyvisa** | **bool** | Whether if visa are denied | 
**BCreditcardmerchantDenymastercard** | **bool** | Whether if mastercard are denied | 
**BCreditcardmerchantDenyamex** | **bool** | Whether if amex are denied | 
**BCreditcardmerchantIsactive** | **bool** | Whether the creditcardmerchant is active or not | 
**SCreditcardmerchantApitoken** | Pointer to **string** | The apitoken of the Creditcardmerchant | [optional] 
**SCreditcardmerchantDescription** | **string** | The description of the Creditcardmerchant | 
**SCreditcardmerchantStoreid** | **string** | The storeid of the Creditcardmerchant | 

## Methods

### NewCreditcardmerchantRequest

`func NewCreditcardmerchantRequest(fkiBankaccountID int32, bCreditcardmerchantDenyvisa bool, bCreditcardmerchantDenymastercard bool, bCreditcardmerchantDenyamex bool, bCreditcardmerchantIsactive bool, sCreditcardmerchantDescription string, sCreditcardmerchantStoreid string, ) *CreditcardmerchantRequest`

NewCreditcardmerchantRequest instantiates a new CreditcardmerchantRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditcardmerchantRequestWithDefaults

`func NewCreditcardmerchantRequestWithDefaults() *CreditcardmerchantRequest`

NewCreditcardmerchantRequestWithDefaults instantiates a new CreditcardmerchantRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiCreditcardmerchantID

`func (o *CreditcardmerchantRequest) GetPkiCreditcardmerchantID() int32`

GetPkiCreditcardmerchantID returns the PkiCreditcardmerchantID field if non-nil, zero value otherwise.

### GetPkiCreditcardmerchantIDOk

`func (o *CreditcardmerchantRequest) GetPkiCreditcardmerchantIDOk() (*int32, bool)`

GetPkiCreditcardmerchantIDOk returns a tuple with the PkiCreditcardmerchantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiCreditcardmerchantID

`func (o *CreditcardmerchantRequest) SetPkiCreditcardmerchantID(v int32)`

SetPkiCreditcardmerchantID sets PkiCreditcardmerchantID field to given value.

### HasPkiCreditcardmerchantID

`func (o *CreditcardmerchantRequest) HasPkiCreditcardmerchantID() bool`

HasPkiCreditcardmerchantID returns a boolean if a field has been set.

### GetFkiBankaccountID

`func (o *CreditcardmerchantRequest) GetFkiBankaccountID() int32`

GetFkiBankaccountID returns the FkiBankaccountID field if non-nil, zero value otherwise.

### GetFkiBankaccountIDOk

`func (o *CreditcardmerchantRequest) GetFkiBankaccountIDOk() (*int32, bool)`

GetFkiBankaccountIDOk returns a tuple with the FkiBankaccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiBankaccountID

`func (o *CreditcardmerchantRequest) SetFkiBankaccountID(v int32)`

SetFkiBankaccountID sets FkiBankaccountID field to given value.


### GetFkiLanguageID

`func (o *CreditcardmerchantRequest) GetFkiLanguageID() int32`

GetFkiLanguageID returns the FkiLanguageID field if non-nil, zero value otherwise.

### GetFkiLanguageIDOk

`func (o *CreditcardmerchantRequest) GetFkiLanguageIDOk() (*int32, bool)`

GetFkiLanguageIDOk returns a tuple with the FkiLanguageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiLanguageID

`func (o *CreditcardmerchantRequest) SetFkiLanguageID(v int32)`

SetFkiLanguageID sets FkiLanguageID field to given value.

### HasFkiLanguageID

`func (o *CreditcardmerchantRequest) HasFkiLanguageID() bool`

HasFkiLanguageID returns a boolean if a field has been set.

### GetBCreditcardmerchantDenyvisa

`func (o *CreditcardmerchantRequest) GetBCreditcardmerchantDenyvisa() bool`

GetBCreditcardmerchantDenyvisa returns the BCreditcardmerchantDenyvisa field if non-nil, zero value otherwise.

### GetBCreditcardmerchantDenyvisaOk

`func (o *CreditcardmerchantRequest) GetBCreditcardmerchantDenyvisaOk() (*bool, bool)`

GetBCreditcardmerchantDenyvisaOk returns a tuple with the BCreditcardmerchantDenyvisa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBCreditcardmerchantDenyvisa

`func (o *CreditcardmerchantRequest) SetBCreditcardmerchantDenyvisa(v bool)`

SetBCreditcardmerchantDenyvisa sets BCreditcardmerchantDenyvisa field to given value.


### GetBCreditcardmerchantDenymastercard

`func (o *CreditcardmerchantRequest) GetBCreditcardmerchantDenymastercard() bool`

GetBCreditcardmerchantDenymastercard returns the BCreditcardmerchantDenymastercard field if non-nil, zero value otherwise.

### GetBCreditcardmerchantDenymastercardOk

`func (o *CreditcardmerchantRequest) GetBCreditcardmerchantDenymastercardOk() (*bool, bool)`

GetBCreditcardmerchantDenymastercardOk returns a tuple with the BCreditcardmerchantDenymastercard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBCreditcardmerchantDenymastercard

`func (o *CreditcardmerchantRequest) SetBCreditcardmerchantDenymastercard(v bool)`

SetBCreditcardmerchantDenymastercard sets BCreditcardmerchantDenymastercard field to given value.


### GetBCreditcardmerchantDenyamex

`func (o *CreditcardmerchantRequest) GetBCreditcardmerchantDenyamex() bool`

GetBCreditcardmerchantDenyamex returns the BCreditcardmerchantDenyamex field if non-nil, zero value otherwise.

### GetBCreditcardmerchantDenyamexOk

`func (o *CreditcardmerchantRequest) GetBCreditcardmerchantDenyamexOk() (*bool, bool)`

GetBCreditcardmerchantDenyamexOk returns a tuple with the BCreditcardmerchantDenyamex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBCreditcardmerchantDenyamex

`func (o *CreditcardmerchantRequest) SetBCreditcardmerchantDenyamex(v bool)`

SetBCreditcardmerchantDenyamex sets BCreditcardmerchantDenyamex field to given value.


### GetBCreditcardmerchantIsactive

`func (o *CreditcardmerchantRequest) GetBCreditcardmerchantIsactive() bool`

GetBCreditcardmerchantIsactive returns the BCreditcardmerchantIsactive field if non-nil, zero value otherwise.

### GetBCreditcardmerchantIsactiveOk

`func (o *CreditcardmerchantRequest) GetBCreditcardmerchantIsactiveOk() (*bool, bool)`

GetBCreditcardmerchantIsactiveOk returns a tuple with the BCreditcardmerchantIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBCreditcardmerchantIsactive

`func (o *CreditcardmerchantRequest) SetBCreditcardmerchantIsactive(v bool)`

SetBCreditcardmerchantIsactive sets BCreditcardmerchantIsactive field to given value.


### GetSCreditcardmerchantApitoken

`func (o *CreditcardmerchantRequest) GetSCreditcardmerchantApitoken() string`

GetSCreditcardmerchantApitoken returns the SCreditcardmerchantApitoken field if non-nil, zero value otherwise.

### GetSCreditcardmerchantApitokenOk

`func (o *CreditcardmerchantRequest) GetSCreditcardmerchantApitokenOk() (*string, bool)`

GetSCreditcardmerchantApitokenOk returns a tuple with the SCreditcardmerchantApitoken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCreditcardmerchantApitoken

`func (o *CreditcardmerchantRequest) SetSCreditcardmerchantApitoken(v string)`

SetSCreditcardmerchantApitoken sets SCreditcardmerchantApitoken field to given value.

### HasSCreditcardmerchantApitoken

`func (o *CreditcardmerchantRequest) HasSCreditcardmerchantApitoken() bool`

HasSCreditcardmerchantApitoken returns a boolean if a field has been set.

### GetSCreditcardmerchantDescription

`func (o *CreditcardmerchantRequest) GetSCreditcardmerchantDescription() string`

GetSCreditcardmerchantDescription returns the SCreditcardmerchantDescription field if non-nil, zero value otherwise.

### GetSCreditcardmerchantDescriptionOk

`func (o *CreditcardmerchantRequest) GetSCreditcardmerchantDescriptionOk() (*string, bool)`

GetSCreditcardmerchantDescriptionOk returns a tuple with the SCreditcardmerchantDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCreditcardmerchantDescription

`func (o *CreditcardmerchantRequest) SetSCreditcardmerchantDescription(v string)`

SetSCreditcardmerchantDescription sets SCreditcardmerchantDescription field to given value.


### GetSCreditcardmerchantStoreid

`func (o *CreditcardmerchantRequest) GetSCreditcardmerchantStoreid() string`

GetSCreditcardmerchantStoreid returns the SCreditcardmerchantStoreid field if non-nil, zero value otherwise.

### GetSCreditcardmerchantStoreidOk

`func (o *CreditcardmerchantRequest) GetSCreditcardmerchantStoreidOk() (*string, bool)`

GetSCreditcardmerchantStoreidOk returns a tuple with the SCreditcardmerchantStoreid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCreditcardmerchantStoreid

`func (o *CreditcardmerchantRequest) SetSCreditcardmerchantStoreid(v string)`

SetSCreditcardmerchantStoreid sets SCreditcardmerchantStoreid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


