# InvoiceListElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiInvoiceID** | **int32** | The unique ID of the Invoice. | 
**FkiAgentID** | Pointer to **int32** | The unique ID of the Agent. | [optional] 
**FkiBrokerID** | Pointer to **int32** | The unique ID of the Broker. | [optional] 
**FkiCustomerID** | Pointer to **int32** | The unique ID of the Customer. | [optional] 
**FkiPeriodID** | **int32** | The unique ID of the Period | 
**SPeriodYYYYMM** | Pointer to **string** | The YYYYMM of the Period | [optional] 
**BInvoiceIspaid** | **bool** | Whether if it&#39;s an ispaid | 
**DInvoiceTotal** | **string** | The total of the Invoice | 
**DInvoicePaid** | **string** | The paid of the Invoice | 
**DInvoiceBalance** | Pointer to **string** | The balance of the Invoice | [optional] 
**DtInvoiceDate** | **string** | The date of the Invoice | 
**EInvoiceType** | [**FieldEInvoiceType**](FieldEInvoiceType.md) |  | 
**SInvoiceNumber** | Pointer to **string** | The number of Invoice | [optional] 
**SInvoiceRecipient** | Pointer to **string** | The recipient of Invoice | [optional] 

## Methods

### NewInvoiceListElement

`func NewInvoiceListElement(pkiInvoiceID int32, fkiPeriodID int32, bInvoiceIspaid bool, dInvoiceTotal string, dInvoicePaid string, dtInvoiceDate string, eInvoiceType FieldEInvoiceType, ) *InvoiceListElement`

NewInvoiceListElement instantiates a new InvoiceListElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceListElementWithDefaults

`func NewInvoiceListElementWithDefaults() *InvoiceListElement`

NewInvoiceListElementWithDefaults instantiates a new InvoiceListElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiInvoiceID

`func (o *InvoiceListElement) GetPkiInvoiceID() int32`

GetPkiInvoiceID returns the PkiInvoiceID field if non-nil, zero value otherwise.

### GetPkiInvoiceIDOk

`func (o *InvoiceListElement) GetPkiInvoiceIDOk() (*int32, bool)`

GetPkiInvoiceIDOk returns a tuple with the PkiInvoiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiInvoiceID

`func (o *InvoiceListElement) SetPkiInvoiceID(v int32)`

SetPkiInvoiceID sets PkiInvoiceID field to given value.


### GetFkiAgentID

`func (o *InvoiceListElement) GetFkiAgentID() int32`

GetFkiAgentID returns the FkiAgentID field if non-nil, zero value otherwise.

### GetFkiAgentIDOk

`func (o *InvoiceListElement) GetFkiAgentIDOk() (*int32, bool)`

GetFkiAgentIDOk returns a tuple with the FkiAgentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiAgentID

`func (o *InvoiceListElement) SetFkiAgentID(v int32)`

SetFkiAgentID sets FkiAgentID field to given value.

### HasFkiAgentID

`func (o *InvoiceListElement) HasFkiAgentID() bool`

HasFkiAgentID returns a boolean if a field has been set.

### GetFkiBrokerID

`func (o *InvoiceListElement) GetFkiBrokerID() int32`

GetFkiBrokerID returns the FkiBrokerID field if non-nil, zero value otherwise.

### GetFkiBrokerIDOk

`func (o *InvoiceListElement) GetFkiBrokerIDOk() (*int32, bool)`

GetFkiBrokerIDOk returns a tuple with the FkiBrokerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiBrokerID

`func (o *InvoiceListElement) SetFkiBrokerID(v int32)`

SetFkiBrokerID sets FkiBrokerID field to given value.

### HasFkiBrokerID

`func (o *InvoiceListElement) HasFkiBrokerID() bool`

HasFkiBrokerID returns a boolean if a field has been set.

### GetFkiCustomerID

`func (o *InvoiceListElement) GetFkiCustomerID() int32`

GetFkiCustomerID returns the FkiCustomerID field if non-nil, zero value otherwise.

### GetFkiCustomerIDOk

`func (o *InvoiceListElement) GetFkiCustomerIDOk() (*int32, bool)`

GetFkiCustomerIDOk returns a tuple with the FkiCustomerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiCustomerID

`func (o *InvoiceListElement) SetFkiCustomerID(v int32)`

SetFkiCustomerID sets FkiCustomerID field to given value.

### HasFkiCustomerID

`func (o *InvoiceListElement) HasFkiCustomerID() bool`

HasFkiCustomerID returns a boolean if a field has been set.

### GetFkiPeriodID

`func (o *InvoiceListElement) GetFkiPeriodID() int32`

GetFkiPeriodID returns the FkiPeriodID field if non-nil, zero value otherwise.

### GetFkiPeriodIDOk

`func (o *InvoiceListElement) GetFkiPeriodIDOk() (*int32, bool)`

GetFkiPeriodIDOk returns a tuple with the FkiPeriodID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiPeriodID

`func (o *InvoiceListElement) SetFkiPeriodID(v int32)`

SetFkiPeriodID sets FkiPeriodID field to given value.


### GetSPeriodYYYYMM

`func (o *InvoiceListElement) GetSPeriodYYYYMM() string`

GetSPeriodYYYYMM returns the SPeriodYYYYMM field if non-nil, zero value otherwise.

### GetSPeriodYYYYMMOk

`func (o *InvoiceListElement) GetSPeriodYYYYMMOk() (*string, bool)`

GetSPeriodYYYYMMOk returns a tuple with the SPeriodYYYYMM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSPeriodYYYYMM

`func (o *InvoiceListElement) SetSPeriodYYYYMM(v string)`

SetSPeriodYYYYMM sets SPeriodYYYYMM field to given value.

### HasSPeriodYYYYMM

`func (o *InvoiceListElement) HasSPeriodYYYYMM() bool`

HasSPeriodYYYYMM returns a boolean if a field has been set.

### GetBInvoiceIspaid

`func (o *InvoiceListElement) GetBInvoiceIspaid() bool`

GetBInvoiceIspaid returns the BInvoiceIspaid field if non-nil, zero value otherwise.

### GetBInvoiceIspaidOk

`func (o *InvoiceListElement) GetBInvoiceIspaidOk() (*bool, bool)`

GetBInvoiceIspaidOk returns a tuple with the BInvoiceIspaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBInvoiceIspaid

`func (o *InvoiceListElement) SetBInvoiceIspaid(v bool)`

SetBInvoiceIspaid sets BInvoiceIspaid field to given value.


### GetDInvoiceTotal

`func (o *InvoiceListElement) GetDInvoiceTotal() string`

GetDInvoiceTotal returns the DInvoiceTotal field if non-nil, zero value otherwise.

### GetDInvoiceTotalOk

`func (o *InvoiceListElement) GetDInvoiceTotalOk() (*string, bool)`

GetDInvoiceTotalOk returns a tuple with the DInvoiceTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDInvoiceTotal

`func (o *InvoiceListElement) SetDInvoiceTotal(v string)`

SetDInvoiceTotal sets DInvoiceTotal field to given value.


### GetDInvoicePaid

`func (o *InvoiceListElement) GetDInvoicePaid() string`

GetDInvoicePaid returns the DInvoicePaid field if non-nil, zero value otherwise.

### GetDInvoicePaidOk

`func (o *InvoiceListElement) GetDInvoicePaidOk() (*string, bool)`

GetDInvoicePaidOk returns a tuple with the DInvoicePaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDInvoicePaid

`func (o *InvoiceListElement) SetDInvoicePaid(v string)`

SetDInvoicePaid sets DInvoicePaid field to given value.


### GetDInvoiceBalance

`func (o *InvoiceListElement) GetDInvoiceBalance() string`

GetDInvoiceBalance returns the DInvoiceBalance field if non-nil, zero value otherwise.

### GetDInvoiceBalanceOk

`func (o *InvoiceListElement) GetDInvoiceBalanceOk() (*string, bool)`

GetDInvoiceBalanceOk returns a tuple with the DInvoiceBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDInvoiceBalance

`func (o *InvoiceListElement) SetDInvoiceBalance(v string)`

SetDInvoiceBalance sets DInvoiceBalance field to given value.

### HasDInvoiceBalance

`func (o *InvoiceListElement) HasDInvoiceBalance() bool`

HasDInvoiceBalance returns a boolean if a field has been set.

### GetDtInvoiceDate

`func (o *InvoiceListElement) GetDtInvoiceDate() string`

GetDtInvoiceDate returns the DtInvoiceDate field if non-nil, zero value otherwise.

### GetDtInvoiceDateOk

`func (o *InvoiceListElement) GetDtInvoiceDateOk() (*string, bool)`

GetDtInvoiceDateOk returns a tuple with the DtInvoiceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtInvoiceDate

`func (o *InvoiceListElement) SetDtInvoiceDate(v string)`

SetDtInvoiceDate sets DtInvoiceDate field to given value.


### GetEInvoiceType

`func (o *InvoiceListElement) GetEInvoiceType() FieldEInvoiceType`

GetEInvoiceType returns the EInvoiceType field if non-nil, zero value otherwise.

### GetEInvoiceTypeOk

`func (o *InvoiceListElement) GetEInvoiceTypeOk() (*FieldEInvoiceType, bool)`

GetEInvoiceTypeOk returns a tuple with the EInvoiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEInvoiceType

`func (o *InvoiceListElement) SetEInvoiceType(v FieldEInvoiceType)`

SetEInvoiceType sets EInvoiceType field to given value.


### GetSInvoiceNumber

`func (o *InvoiceListElement) GetSInvoiceNumber() string`

GetSInvoiceNumber returns the SInvoiceNumber field if non-nil, zero value otherwise.

### GetSInvoiceNumberOk

`func (o *InvoiceListElement) GetSInvoiceNumberOk() (*string, bool)`

GetSInvoiceNumberOk returns a tuple with the SInvoiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSInvoiceNumber

`func (o *InvoiceListElement) SetSInvoiceNumber(v string)`

SetSInvoiceNumber sets SInvoiceNumber field to given value.

### HasSInvoiceNumber

`func (o *InvoiceListElement) HasSInvoiceNumber() bool`

HasSInvoiceNumber returns a boolean if a field has been set.

### GetSInvoiceRecipient

`func (o *InvoiceListElement) GetSInvoiceRecipient() string`

GetSInvoiceRecipient returns the SInvoiceRecipient field if non-nil, zero value otherwise.

### GetSInvoiceRecipientOk

`func (o *InvoiceListElement) GetSInvoiceRecipientOk() (*string, bool)`

GetSInvoiceRecipientOk returns a tuple with the SInvoiceRecipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSInvoiceRecipient

`func (o *InvoiceListElement) SetSInvoiceRecipient(v string)`

SetSInvoiceRecipient sets SInvoiceRecipient field to given value.

### HasSInvoiceRecipient

`func (o *InvoiceListElement) HasSInvoiceRecipient() bool`

HasSInvoiceRecipient returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


