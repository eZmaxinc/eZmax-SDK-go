# DepositListElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiDepositID** | **int32** | The unique ID of the Deposit | 
**FkiBankaccountID** | **int32** | The unique ID of the Bankaccount | 
**FkiPeriodID** | **int32** | The unique ID of the Period | 
**FkiUserIDCreated** | **int32** | The unique ID of the User | 
**SDepositNumber** | **string** | The number of the Deposit | 
**DtDepositDate** | **string** | The date of the Deposit | 
**DDepositTotal** | **string** | The total of the Deposit | 
**DtCreatedDate** | **string** | The date and time at which the object was created | 
**BDepositDraft** | **bool** | Whether if it&#39;s an draft | 
**BDepositReconciled** | **bool** | Whether if it&#39;s reconciled | 
**FkiReconciliationID** | Pointer to **int32** | The unique ID of the Reconciliation | [optional] 
**SBankaccountBankname** | **string** | The name of the bank | 
**SPeriodYYYYMM** | **string** | The YYYYMM of the Period | 
**SUserLoginname** | **string** | The login name of the User. | 

## Methods

### NewDepositListElement

`func NewDepositListElement(pkiDepositID int32, fkiBankaccountID int32, fkiPeriodID int32, fkiUserIDCreated int32, sDepositNumber string, dtDepositDate string, dDepositTotal string, dtCreatedDate string, bDepositDraft bool, bDepositReconciled bool, sBankaccountBankname string, sPeriodYYYYMM string, sUserLoginname string, ) *DepositListElement`

NewDepositListElement instantiates a new DepositListElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDepositListElementWithDefaults

`func NewDepositListElementWithDefaults() *DepositListElement`

NewDepositListElementWithDefaults instantiates a new DepositListElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiDepositID

`func (o *DepositListElement) GetPkiDepositID() int32`

GetPkiDepositID returns the PkiDepositID field if non-nil, zero value otherwise.

### GetPkiDepositIDOk

`func (o *DepositListElement) GetPkiDepositIDOk() (*int32, bool)`

GetPkiDepositIDOk returns a tuple with the PkiDepositID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiDepositID

`func (o *DepositListElement) SetPkiDepositID(v int32)`

SetPkiDepositID sets PkiDepositID field to given value.


### GetFkiBankaccountID

`func (o *DepositListElement) GetFkiBankaccountID() int32`

GetFkiBankaccountID returns the FkiBankaccountID field if non-nil, zero value otherwise.

### GetFkiBankaccountIDOk

`func (o *DepositListElement) GetFkiBankaccountIDOk() (*int32, bool)`

GetFkiBankaccountIDOk returns a tuple with the FkiBankaccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiBankaccountID

`func (o *DepositListElement) SetFkiBankaccountID(v int32)`

SetFkiBankaccountID sets FkiBankaccountID field to given value.


### GetFkiPeriodID

`func (o *DepositListElement) GetFkiPeriodID() int32`

GetFkiPeriodID returns the FkiPeriodID field if non-nil, zero value otherwise.

### GetFkiPeriodIDOk

`func (o *DepositListElement) GetFkiPeriodIDOk() (*int32, bool)`

GetFkiPeriodIDOk returns a tuple with the FkiPeriodID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiPeriodID

`func (o *DepositListElement) SetFkiPeriodID(v int32)`

SetFkiPeriodID sets FkiPeriodID field to given value.


### GetFkiUserIDCreated

`func (o *DepositListElement) GetFkiUserIDCreated() int32`

GetFkiUserIDCreated returns the FkiUserIDCreated field if non-nil, zero value otherwise.

### GetFkiUserIDCreatedOk

`func (o *DepositListElement) GetFkiUserIDCreatedOk() (*int32, bool)`

GetFkiUserIDCreatedOk returns a tuple with the FkiUserIDCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUserIDCreated

`func (o *DepositListElement) SetFkiUserIDCreated(v int32)`

SetFkiUserIDCreated sets FkiUserIDCreated field to given value.


### GetSDepositNumber

`func (o *DepositListElement) GetSDepositNumber() string`

GetSDepositNumber returns the SDepositNumber field if non-nil, zero value otherwise.

### GetSDepositNumberOk

`func (o *DepositListElement) GetSDepositNumberOk() (*string, bool)`

GetSDepositNumberOk returns a tuple with the SDepositNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSDepositNumber

`func (o *DepositListElement) SetSDepositNumber(v string)`

SetSDepositNumber sets SDepositNumber field to given value.


### GetDtDepositDate

`func (o *DepositListElement) GetDtDepositDate() string`

GetDtDepositDate returns the DtDepositDate field if non-nil, zero value otherwise.

### GetDtDepositDateOk

`func (o *DepositListElement) GetDtDepositDateOk() (*string, bool)`

GetDtDepositDateOk returns a tuple with the DtDepositDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtDepositDate

`func (o *DepositListElement) SetDtDepositDate(v string)`

SetDtDepositDate sets DtDepositDate field to given value.


### GetDDepositTotal

`func (o *DepositListElement) GetDDepositTotal() string`

GetDDepositTotal returns the DDepositTotal field if non-nil, zero value otherwise.

### GetDDepositTotalOk

`func (o *DepositListElement) GetDDepositTotalOk() (*string, bool)`

GetDDepositTotalOk returns a tuple with the DDepositTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDDepositTotal

`func (o *DepositListElement) SetDDepositTotal(v string)`

SetDDepositTotal sets DDepositTotal field to given value.


### GetDtCreatedDate

`func (o *DepositListElement) GetDtCreatedDate() string`

GetDtCreatedDate returns the DtCreatedDate field if non-nil, zero value otherwise.

### GetDtCreatedDateOk

`func (o *DepositListElement) GetDtCreatedDateOk() (*string, bool)`

GetDtCreatedDateOk returns a tuple with the DtCreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtCreatedDate

`func (o *DepositListElement) SetDtCreatedDate(v string)`

SetDtCreatedDate sets DtCreatedDate field to given value.


### GetBDepositDraft

`func (o *DepositListElement) GetBDepositDraft() bool`

GetBDepositDraft returns the BDepositDraft field if non-nil, zero value otherwise.

### GetBDepositDraftOk

`func (o *DepositListElement) GetBDepositDraftOk() (*bool, bool)`

GetBDepositDraftOk returns a tuple with the BDepositDraft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBDepositDraft

`func (o *DepositListElement) SetBDepositDraft(v bool)`

SetBDepositDraft sets BDepositDraft field to given value.


### GetBDepositReconciled

`func (o *DepositListElement) GetBDepositReconciled() bool`

GetBDepositReconciled returns the BDepositReconciled field if non-nil, zero value otherwise.

### GetBDepositReconciledOk

`func (o *DepositListElement) GetBDepositReconciledOk() (*bool, bool)`

GetBDepositReconciledOk returns a tuple with the BDepositReconciled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBDepositReconciled

`func (o *DepositListElement) SetBDepositReconciled(v bool)`

SetBDepositReconciled sets BDepositReconciled field to given value.


### GetFkiReconciliationID

`func (o *DepositListElement) GetFkiReconciliationID() int32`

GetFkiReconciliationID returns the FkiReconciliationID field if non-nil, zero value otherwise.

### GetFkiReconciliationIDOk

`func (o *DepositListElement) GetFkiReconciliationIDOk() (*int32, bool)`

GetFkiReconciliationIDOk returns a tuple with the FkiReconciliationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiReconciliationID

`func (o *DepositListElement) SetFkiReconciliationID(v int32)`

SetFkiReconciliationID sets FkiReconciliationID field to given value.

### HasFkiReconciliationID

`func (o *DepositListElement) HasFkiReconciliationID() bool`

HasFkiReconciliationID returns a boolean if a field has been set.

### GetSBankaccountBankname

`func (o *DepositListElement) GetSBankaccountBankname() string`

GetSBankaccountBankname returns the SBankaccountBankname field if non-nil, zero value otherwise.

### GetSBankaccountBanknameOk

`func (o *DepositListElement) GetSBankaccountBanknameOk() (*string, bool)`

GetSBankaccountBanknameOk returns a tuple with the SBankaccountBankname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSBankaccountBankname

`func (o *DepositListElement) SetSBankaccountBankname(v string)`

SetSBankaccountBankname sets SBankaccountBankname field to given value.


### GetSPeriodYYYYMM

`func (o *DepositListElement) GetSPeriodYYYYMM() string`

GetSPeriodYYYYMM returns the SPeriodYYYYMM field if non-nil, zero value otherwise.

### GetSPeriodYYYYMMOk

`func (o *DepositListElement) GetSPeriodYYYYMMOk() (*string, bool)`

GetSPeriodYYYYMMOk returns a tuple with the SPeriodYYYYMM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSPeriodYYYYMM

`func (o *DepositListElement) SetSPeriodYYYYMM(v string)`

SetSPeriodYYYYMM sets SPeriodYYYYMM field to given value.


### GetSUserLoginname

`func (o *DepositListElement) GetSUserLoginname() string`

GetSUserLoginname returns the SUserLoginname field if non-nil, zero value otherwise.

### GetSUserLoginnameOk

`func (o *DepositListElement) GetSUserLoginnameOk() (*string, bool)`

GetSUserLoginnameOk returns a tuple with the SUserLoginname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUserLoginname

`func (o *DepositListElement) SetSUserLoginname(v string)`

SetSUserLoginname sets SUserLoginname field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


