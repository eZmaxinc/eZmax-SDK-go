# DeposittransitchequeListElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiDeposittransitchequeID** | **int32** | The unique ID of the Deposittransitcheque | 
**FkiBankaccountID** | **int32** | The unique ID of the Bankaccount | 
**FkiDepositID** | Pointer to **int32** | The unique ID of the Deposit | [optional] 
**SDeposittransitchequeNumber** | **string** | The number of the Deposittransitcheque | 
**DtDeposittransitchequeDate** | **string** | The date of the Deposittransitcheque | 
**DDeposittransitchequeAmount** | **string** | The amount of the Deposittransitcheque | 
**SBankaccountBankname** | **string** | The name of the bank | 
**SPeriodYYYYMM** | **string** | The YYYYMM of the Period | 
**SCustomerName** | **string** | The name of the Customer | 

## Methods

### NewDeposittransitchequeListElement

`func NewDeposittransitchequeListElement(pkiDeposittransitchequeID int32, fkiBankaccountID int32, sDeposittransitchequeNumber string, dtDeposittransitchequeDate string, dDeposittransitchequeAmount string, sBankaccountBankname string, sPeriodYYYYMM string, sCustomerName string, ) *DeposittransitchequeListElement`

NewDeposittransitchequeListElement instantiates a new DeposittransitchequeListElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeposittransitchequeListElementWithDefaults

`func NewDeposittransitchequeListElementWithDefaults() *DeposittransitchequeListElement`

NewDeposittransitchequeListElementWithDefaults instantiates a new DeposittransitchequeListElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiDeposittransitchequeID

`func (o *DeposittransitchequeListElement) GetPkiDeposittransitchequeID() int32`

GetPkiDeposittransitchequeID returns the PkiDeposittransitchequeID field if non-nil, zero value otherwise.

### GetPkiDeposittransitchequeIDOk

`func (o *DeposittransitchequeListElement) GetPkiDeposittransitchequeIDOk() (*int32, bool)`

GetPkiDeposittransitchequeIDOk returns a tuple with the PkiDeposittransitchequeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiDeposittransitchequeID

`func (o *DeposittransitchequeListElement) SetPkiDeposittransitchequeID(v int32)`

SetPkiDeposittransitchequeID sets PkiDeposittransitchequeID field to given value.


### GetFkiBankaccountID

`func (o *DeposittransitchequeListElement) GetFkiBankaccountID() int32`

GetFkiBankaccountID returns the FkiBankaccountID field if non-nil, zero value otherwise.

### GetFkiBankaccountIDOk

`func (o *DeposittransitchequeListElement) GetFkiBankaccountIDOk() (*int32, bool)`

GetFkiBankaccountIDOk returns a tuple with the FkiBankaccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiBankaccountID

`func (o *DeposittransitchequeListElement) SetFkiBankaccountID(v int32)`

SetFkiBankaccountID sets FkiBankaccountID field to given value.


### GetFkiDepositID

`func (o *DeposittransitchequeListElement) GetFkiDepositID() int32`

GetFkiDepositID returns the FkiDepositID field if non-nil, zero value otherwise.

### GetFkiDepositIDOk

`func (o *DeposittransitchequeListElement) GetFkiDepositIDOk() (*int32, bool)`

GetFkiDepositIDOk returns a tuple with the FkiDepositID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiDepositID

`func (o *DeposittransitchequeListElement) SetFkiDepositID(v int32)`

SetFkiDepositID sets FkiDepositID field to given value.

### HasFkiDepositID

`func (o *DeposittransitchequeListElement) HasFkiDepositID() bool`

HasFkiDepositID returns a boolean if a field has been set.

### GetSDeposittransitchequeNumber

`func (o *DeposittransitchequeListElement) GetSDeposittransitchequeNumber() string`

GetSDeposittransitchequeNumber returns the SDeposittransitchequeNumber field if non-nil, zero value otherwise.

### GetSDeposittransitchequeNumberOk

`func (o *DeposittransitchequeListElement) GetSDeposittransitchequeNumberOk() (*string, bool)`

GetSDeposittransitchequeNumberOk returns a tuple with the SDeposittransitchequeNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSDeposittransitchequeNumber

`func (o *DeposittransitchequeListElement) SetSDeposittransitchequeNumber(v string)`

SetSDeposittransitchequeNumber sets SDeposittransitchequeNumber field to given value.


### GetDtDeposittransitchequeDate

`func (o *DeposittransitchequeListElement) GetDtDeposittransitchequeDate() string`

GetDtDeposittransitchequeDate returns the DtDeposittransitchequeDate field if non-nil, zero value otherwise.

### GetDtDeposittransitchequeDateOk

`func (o *DeposittransitchequeListElement) GetDtDeposittransitchequeDateOk() (*string, bool)`

GetDtDeposittransitchequeDateOk returns a tuple with the DtDeposittransitchequeDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtDeposittransitchequeDate

`func (o *DeposittransitchequeListElement) SetDtDeposittransitchequeDate(v string)`

SetDtDeposittransitchequeDate sets DtDeposittransitchequeDate field to given value.


### GetDDeposittransitchequeAmount

`func (o *DeposittransitchequeListElement) GetDDeposittransitchequeAmount() string`

GetDDeposittransitchequeAmount returns the DDeposittransitchequeAmount field if non-nil, zero value otherwise.

### GetDDeposittransitchequeAmountOk

`func (o *DeposittransitchequeListElement) GetDDeposittransitchequeAmountOk() (*string, bool)`

GetDDeposittransitchequeAmountOk returns a tuple with the DDeposittransitchequeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDDeposittransitchequeAmount

`func (o *DeposittransitchequeListElement) SetDDeposittransitchequeAmount(v string)`

SetDDeposittransitchequeAmount sets DDeposittransitchequeAmount field to given value.


### GetSBankaccountBankname

`func (o *DeposittransitchequeListElement) GetSBankaccountBankname() string`

GetSBankaccountBankname returns the SBankaccountBankname field if non-nil, zero value otherwise.

### GetSBankaccountBanknameOk

`func (o *DeposittransitchequeListElement) GetSBankaccountBanknameOk() (*string, bool)`

GetSBankaccountBanknameOk returns a tuple with the SBankaccountBankname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSBankaccountBankname

`func (o *DeposittransitchequeListElement) SetSBankaccountBankname(v string)`

SetSBankaccountBankname sets SBankaccountBankname field to given value.


### GetSPeriodYYYYMM

`func (o *DeposittransitchequeListElement) GetSPeriodYYYYMM() string`

GetSPeriodYYYYMM returns the SPeriodYYYYMM field if non-nil, zero value otherwise.

### GetSPeriodYYYYMMOk

`func (o *DeposittransitchequeListElement) GetSPeriodYYYYMMOk() (*string, bool)`

GetSPeriodYYYYMMOk returns a tuple with the SPeriodYYYYMM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSPeriodYYYYMM

`func (o *DeposittransitchequeListElement) SetSPeriodYYYYMM(v string)`

SetSPeriodYYYYMM sets SPeriodYYYYMM field to given value.


### GetSCustomerName

`func (o *DeposittransitchequeListElement) GetSCustomerName() string`

GetSCustomerName returns the SCustomerName field if non-nil, zero value otherwise.

### GetSCustomerNameOk

`func (o *DeposittransitchequeListElement) GetSCustomerNameOk() (*string, bool)`

GetSCustomerNameOk returns a tuple with the SCustomerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCustomerName

`func (o *DeposittransitchequeListElement) SetSCustomerName(v string)`

SetSCustomerName sets SCustomerName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


