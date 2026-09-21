# BankaccountListElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiBankaccountID** | **int32** | The unique ID of the Bankaccount | 
**SBankaccountBankname** | **string** | The name of the bank | 
**SBankaccountTransit** | **string** | The transit of the Bankaccount | 
**SBankaccountInstitution** | **string** | The institution of the Bankaccount | 
**SBankaccountAccount** | **string** | The account of the Bankaccount | 
**SBankaccountNumber** | **string** | The account number of the Bankaccount | 
**EBankaccountType** | Pointer to [**FieldEBankaccountType**](FieldEBankaccountType.md) |  | [optional] 
**BBankaccountCommissionprocessing** | Pointer to **bool** | Whether if it&#39;s an commissionprocessing | [optional] 
**BBankaccountIsactive** | **bool** | Whether the Bankaccount is active or not | 
**IGlaccountCode** | **int32** | The Code of the Glaccount | 
**SGlaccountDescriptionX** | **string** | The Description for the Glaccount in the language of the requester | 

## Methods

### NewBankaccountListElement

`func NewBankaccountListElement(pkiBankaccountID int32, sBankaccountBankname string, sBankaccountTransit string, sBankaccountInstitution string, sBankaccountAccount string, sBankaccountNumber string, bBankaccountIsactive bool, iGlaccountCode int32, sGlaccountDescriptionX string, ) *BankaccountListElement`

NewBankaccountListElement instantiates a new BankaccountListElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankaccountListElementWithDefaults

`func NewBankaccountListElementWithDefaults() *BankaccountListElement`

NewBankaccountListElementWithDefaults instantiates a new BankaccountListElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiBankaccountID

`func (o *BankaccountListElement) GetPkiBankaccountID() int32`

GetPkiBankaccountID returns the PkiBankaccountID field if non-nil, zero value otherwise.

### GetPkiBankaccountIDOk

`func (o *BankaccountListElement) GetPkiBankaccountIDOk() (*int32, bool)`

GetPkiBankaccountIDOk returns a tuple with the PkiBankaccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiBankaccountID

`func (o *BankaccountListElement) SetPkiBankaccountID(v int32)`

SetPkiBankaccountID sets PkiBankaccountID field to given value.


### GetSBankaccountBankname

`func (o *BankaccountListElement) GetSBankaccountBankname() string`

GetSBankaccountBankname returns the SBankaccountBankname field if non-nil, zero value otherwise.

### GetSBankaccountBanknameOk

`func (o *BankaccountListElement) GetSBankaccountBanknameOk() (*string, bool)`

GetSBankaccountBanknameOk returns a tuple with the SBankaccountBankname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSBankaccountBankname

`func (o *BankaccountListElement) SetSBankaccountBankname(v string)`

SetSBankaccountBankname sets SBankaccountBankname field to given value.


### GetSBankaccountTransit

`func (o *BankaccountListElement) GetSBankaccountTransit() string`

GetSBankaccountTransit returns the SBankaccountTransit field if non-nil, zero value otherwise.

### GetSBankaccountTransitOk

`func (o *BankaccountListElement) GetSBankaccountTransitOk() (*string, bool)`

GetSBankaccountTransitOk returns a tuple with the SBankaccountTransit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSBankaccountTransit

`func (o *BankaccountListElement) SetSBankaccountTransit(v string)`

SetSBankaccountTransit sets SBankaccountTransit field to given value.


### GetSBankaccountInstitution

`func (o *BankaccountListElement) GetSBankaccountInstitution() string`

GetSBankaccountInstitution returns the SBankaccountInstitution field if non-nil, zero value otherwise.

### GetSBankaccountInstitutionOk

`func (o *BankaccountListElement) GetSBankaccountInstitutionOk() (*string, bool)`

GetSBankaccountInstitutionOk returns a tuple with the SBankaccountInstitution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSBankaccountInstitution

`func (o *BankaccountListElement) SetSBankaccountInstitution(v string)`

SetSBankaccountInstitution sets SBankaccountInstitution field to given value.


### GetSBankaccountAccount

`func (o *BankaccountListElement) GetSBankaccountAccount() string`

GetSBankaccountAccount returns the SBankaccountAccount field if non-nil, zero value otherwise.

### GetSBankaccountAccountOk

`func (o *BankaccountListElement) GetSBankaccountAccountOk() (*string, bool)`

GetSBankaccountAccountOk returns a tuple with the SBankaccountAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSBankaccountAccount

`func (o *BankaccountListElement) SetSBankaccountAccount(v string)`

SetSBankaccountAccount sets SBankaccountAccount field to given value.


### GetSBankaccountNumber

`func (o *BankaccountListElement) GetSBankaccountNumber() string`

GetSBankaccountNumber returns the SBankaccountNumber field if non-nil, zero value otherwise.

### GetSBankaccountNumberOk

`func (o *BankaccountListElement) GetSBankaccountNumberOk() (*string, bool)`

GetSBankaccountNumberOk returns a tuple with the SBankaccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSBankaccountNumber

`func (o *BankaccountListElement) SetSBankaccountNumber(v string)`

SetSBankaccountNumber sets SBankaccountNumber field to given value.


### GetEBankaccountType

`func (o *BankaccountListElement) GetEBankaccountType() FieldEBankaccountType`

GetEBankaccountType returns the EBankaccountType field if non-nil, zero value otherwise.

### GetEBankaccountTypeOk

`func (o *BankaccountListElement) GetEBankaccountTypeOk() (*FieldEBankaccountType, bool)`

GetEBankaccountTypeOk returns a tuple with the EBankaccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEBankaccountType

`func (o *BankaccountListElement) SetEBankaccountType(v FieldEBankaccountType)`

SetEBankaccountType sets EBankaccountType field to given value.

### HasEBankaccountType

`func (o *BankaccountListElement) HasEBankaccountType() bool`

HasEBankaccountType returns a boolean if a field has been set.

### GetBBankaccountCommissionprocessing

`func (o *BankaccountListElement) GetBBankaccountCommissionprocessing() bool`

GetBBankaccountCommissionprocessing returns the BBankaccountCommissionprocessing field if non-nil, zero value otherwise.

### GetBBankaccountCommissionprocessingOk

`func (o *BankaccountListElement) GetBBankaccountCommissionprocessingOk() (*bool, bool)`

GetBBankaccountCommissionprocessingOk returns a tuple with the BBankaccountCommissionprocessing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBBankaccountCommissionprocessing

`func (o *BankaccountListElement) SetBBankaccountCommissionprocessing(v bool)`

SetBBankaccountCommissionprocessing sets BBankaccountCommissionprocessing field to given value.

### HasBBankaccountCommissionprocessing

`func (o *BankaccountListElement) HasBBankaccountCommissionprocessing() bool`

HasBBankaccountCommissionprocessing returns a boolean if a field has been set.

### GetBBankaccountIsactive

`func (o *BankaccountListElement) GetBBankaccountIsactive() bool`

GetBBankaccountIsactive returns the BBankaccountIsactive field if non-nil, zero value otherwise.

### GetBBankaccountIsactiveOk

`func (o *BankaccountListElement) GetBBankaccountIsactiveOk() (*bool, bool)`

GetBBankaccountIsactiveOk returns a tuple with the BBankaccountIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBBankaccountIsactive

`func (o *BankaccountListElement) SetBBankaccountIsactive(v bool)`

SetBBankaccountIsactive sets BBankaccountIsactive field to given value.


### GetIGlaccountCode

`func (o *BankaccountListElement) GetIGlaccountCode() int32`

GetIGlaccountCode returns the IGlaccountCode field if non-nil, zero value otherwise.

### GetIGlaccountCodeOk

`func (o *BankaccountListElement) GetIGlaccountCodeOk() (*int32, bool)`

GetIGlaccountCodeOk returns a tuple with the IGlaccountCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIGlaccountCode

`func (o *BankaccountListElement) SetIGlaccountCode(v int32)`

SetIGlaccountCode sets IGlaccountCode field to given value.


### GetSGlaccountDescriptionX

`func (o *BankaccountListElement) GetSGlaccountDescriptionX() string`

GetSGlaccountDescriptionX returns the SGlaccountDescriptionX field if non-nil, zero value otherwise.

### GetSGlaccountDescriptionXOk

`func (o *BankaccountListElement) GetSGlaccountDescriptionXOk() (*string, bool)`

GetSGlaccountDescriptionXOk returns a tuple with the SGlaccountDescriptionX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSGlaccountDescriptionX

`func (o *BankaccountListElement) SetSGlaccountDescriptionX(v string)`

SetSGlaccountDescriptionX sets SGlaccountDescriptionX field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


