# ElectronicfundstransferListElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiElectronicfundstransferID** | **int32** | The unique ID of the Electronicfundstransfer | 
**FkiPaymentmethodID** | **int32** | The unique ID of the Paymentmethod | 
**FkiUserIDCreated** | Pointer to **int32** | The unique ID of the User | [optional] 
**SPaymentmethodDescriptionX** | Pointer to **string** | The description of the Paymentmethod in the language of the requester | [optional] 
**SUserLoginnameCreated** | Pointer to **string** | The login name of the User. | [optional] 
**IElectronicfundstransferNumber** | **int32** | The number of the Electronicfundstransfer | 
**DtElectronicfundstransferDate** | **string** | The date of the Electronicfundstransfer | 
**IElectronicfundstransferTransactioncount** | Pointer to **int32** | The count of transactions in this Electronicfundstransfer. | [optional] 
**DElectronicfundstransferAmountvisible** | Pointer to **string** | The sum of Electronicfundstransfertransaction we can consult. This sum can be different than total amount. | [optional] 

## Methods

### NewElectronicfundstransferListElement

`func NewElectronicfundstransferListElement(pkiElectronicfundstransferID int32, fkiPaymentmethodID int32, iElectronicfundstransferNumber int32, dtElectronicfundstransferDate string, ) *ElectronicfundstransferListElement`

NewElectronicfundstransferListElement instantiates a new ElectronicfundstransferListElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewElectronicfundstransferListElementWithDefaults

`func NewElectronicfundstransferListElementWithDefaults() *ElectronicfundstransferListElement`

NewElectronicfundstransferListElementWithDefaults instantiates a new ElectronicfundstransferListElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiElectronicfundstransferID

`func (o *ElectronicfundstransferListElement) GetPkiElectronicfundstransferID() int32`

GetPkiElectronicfundstransferID returns the PkiElectronicfundstransferID field if non-nil, zero value otherwise.

### GetPkiElectronicfundstransferIDOk

`func (o *ElectronicfundstransferListElement) GetPkiElectronicfundstransferIDOk() (*int32, bool)`

GetPkiElectronicfundstransferIDOk returns a tuple with the PkiElectronicfundstransferID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiElectronicfundstransferID

`func (o *ElectronicfundstransferListElement) SetPkiElectronicfundstransferID(v int32)`

SetPkiElectronicfundstransferID sets PkiElectronicfundstransferID field to given value.


### GetFkiPaymentmethodID

`func (o *ElectronicfundstransferListElement) GetFkiPaymentmethodID() int32`

GetFkiPaymentmethodID returns the FkiPaymentmethodID field if non-nil, zero value otherwise.

### GetFkiPaymentmethodIDOk

`func (o *ElectronicfundstransferListElement) GetFkiPaymentmethodIDOk() (*int32, bool)`

GetFkiPaymentmethodIDOk returns a tuple with the FkiPaymentmethodID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiPaymentmethodID

`func (o *ElectronicfundstransferListElement) SetFkiPaymentmethodID(v int32)`

SetFkiPaymentmethodID sets FkiPaymentmethodID field to given value.


### GetFkiUserIDCreated

`func (o *ElectronicfundstransferListElement) GetFkiUserIDCreated() int32`

GetFkiUserIDCreated returns the FkiUserIDCreated field if non-nil, zero value otherwise.

### GetFkiUserIDCreatedOk

`func (o *ElectronicfundstransferListElement) GetFkiUserIDCreatedOk() (*int32, bool)`

GetFkiUserIDCreatedOk returns a tuple with the FkiUserIDCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUserIDCreated

`func (o *ElectronicfundstransferListElement) SetFkiUserIDCreated(v int32)`

SetFkiUserIDCreated sets FkiUserIDCreated field to given value.

### HasFkiUserIDCreated

`func (o *ElectronicfundstransferListElement) HasFkiUserIDCreated() bool`

HasFkiUserIDCreated returns a boolean if a field has been set.

### GetSPaymentmethodDescriptionX

`func (o *ElectronicfundstransferListElement) GetSPaymentmethodDescriptionX() string`

GetSPaymentmethodDescriptionX returns the SPaymentmethodDescriptionX field if non-nil, zero value otherwise.

### GetSPaymentmethodDescriptionXOk

`func (o *ElectronicfundstransferListElement) GetSPaymentmethodDescriptionXOk() (*string, bool)`

GetSPaymentmethodDescriptionXOk returns a tuple with the SPaymentmethodDescriptionX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSPaymentmethodDescriptionX

`func (o *ElectronicfundstransferListElement) SetSPaymentmethodDescriptionX(v string)`

SetSPaymentmethodDescriptionX sets SPaymentmethodDescriptionX field to given value.

### HasSPaymentmethodDescriptionX

`func (o *ElectronicfundstransferListElement) HasSPaymentmethodDescriptionX() bool`

HasSPaymentmethodDescriptionX returns a boolean if a field has been set.

### GetSUserLoginnameCreated

`func (o *ElectronicfundstransferListElement) GetSUserLoginnameCreated() string`

GetSUserLoginnameCreated returns the SUserLoginnameCreated field if non-nil, zero value otherwise.

### GetSUserLoginnameCreatedOk

`func (o *ElectronicfundstransferListElement) GetSUserLoginnameCreatedOk() (*string, bool)`

GetSUserLoginnameCreatedOk returns a tuple with the SUserLoginnameCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUserLoginnameCreated

`func (o *ElectronicfundstransferListElement) SetSUserLoginnameCreated(v string)`

SetSUserLoginnameCreated sets SUserLoginnameCreated field to given value.

### HasSUserLoginnameCreated

`func (o *ElectronicfundstransferListElement) HasSUserLoginnameCreated() bool`

HasSUserLoginnameCreated returns a boolean if a field has been set.

### GetIElectronicfundstransferNumber

`func (o *ElectronicfundstransferListElement) GetIElectronicfundstransferNumber() int32`

GetIElectronicfundstransferNumber returns the IElectronicfundstransferNumber field if non-nil, zero value otherwise.

### GetIElectronicfundstransferNumberOk

`func (o *ElectronicfundstransferListElement) GetIElectronicfundstransferNumberOk() (*int32, bool)`

GetIElectronicfundstransferNumberOk returns a tuple with the IElectronicfundstransferNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIElectronicfundstransferNumber

`func (o *ElectronicfundstransferListElement) SetIElectronicfundstransferNumber(v int32)`

SetIElectronicfundstransferNumber sets IElectronicfundstransferNumber field to given value.


### GetDtElectronicfundstransferDate

`func (o *ElectronicfundstransferListElement) GetDtElectronicfundstransferDate() string`

GetDtElectronicfundstransferDate returns the DtElectronicfundstransferDate field if non-nil, zero value otherwise.

### GetDtElectronicfundstransferDateOk

`func (o *ElectronicfundstransferListElement) GetDtElectronicfundstransferDateOk() (*string, bool)`

GetDtElectronicfundstransferDateOk returns a tuple with the DtElectronicfundstransferDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtElectronicfundstransferDate

`func (o *ElectronicfundstransferListElement) SetDtElectronicfundstransferDate(v string)`

SetDtElectronicfundstransferDate sets DtElectronicfundstransferDate field to given value.


### GetIElectronicfundstransferTransactioncount

`func (o *ElectronicfundstransferListElement) GetIElectronicfundstransferTransactioncount() int32`

GetIElectronicfundstransferTransactioncount returns the IElectronicfundstransferTransactioncount field if non-nil, zero value otherwise.

### GetIElectronicfundstransferTransactioncountOk

`func (o *ElectronicfundstransferListElement) GetIElectronicfundstransferTransactioncountOk() (*int32, bool)`

GetIElectronicfundstransferTransactioncountOk returns a tuple with the IElectronicfundstransferTransactioncount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIElectronicfundstransferTransactioncount

`func (o *ElectronicfundstransferListElement) SetIElectronicfundstransferTransactioncount(v int32)`

SetIElectronicfundstransferTransactioncount sets IElectronicfundstransferTransactioncount field to given value.

### HasIElectronicfundstransferTransactioncount

`func (o *ElectronicfundstransferListElement) HasIElectronicfundstransferTransactioncount() bool`

HasIElectronicfundstransferTransactioncount returns a boolean if a field has been set.

### GetDElectronicfundstransferAmountvisible

`func (o *ElectronicfundstransferListElement) GetDElectronicfundstransferAmountvisible() string`

GetDElectronicfundstransferAmountvisible returns the DElectronicfundstransferAmountvisible field if non-nil, zero value otherwise.

### GetDElectronicfundstransferAmountvisibleOk

`func (o *ElectronicfundstransferListElement) GetDElectronicfundstransferAmountvisibleOk() (*string, bool)`

GetDElectronicfundstransferAmountvisibleOk returns a tuple with the DElectronicfundstransferAmountvisible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDElectronicfundstransferAmountvisible

`func (o *ElectronicfundstransferListElement) SetDElectronicfundstransferAmountvisible(v string)`

SetDElectronicfundstransferAmountvisible sets DElectronicfundstransferAmountvisible field to given value.

### HasDElectronicfundstransferAmountvisible

`func (o *ElectronicfundstransferListElement) HasDElectronicfundstransferAmountvisible() bool`

HasDElectronicfundstransferAmountvisible returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


