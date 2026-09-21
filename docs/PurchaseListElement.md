# PurchaseListElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiPurchaseID** | **int32** | The unique ID of the Purchase | 
**FkiFranchiseofficeID** | Pointer to **int32** | The unique ID of the Franchisereoffice | [optional] 
**FkiSupplierID** | Pointer to **int32** | The unique ID of the Supplier. | [optional] 
**FkiPeriodID** | **int32** | The unique ID of the Period | 
**SPurchaseBillnumber** | **string** | The Billnumber of the Purchase | 
**DtPurchaseDate** | **string** | The Date of the Purchase | 
**DtPurchasePaydate** | **string** | The Paydate of the Purchase | 
**DPurchaseTotal** | **string** | The Total of the Purchase | 
**BPurchaseIsfinanced** | **bool** | Whether if it&#39;s an isfinanced | 
**BPurchaseIspaid** | **bool** | Whether if it&#39;s an ispaid | 
**SFranchiseagenceName** | Pointer to **string** | The name of the Franchiseagence | [optional] 
**IFranchiseofficeCode** | Pointer to **string** | The code of the Franchiseoffice | [optional] 
**SSupplierName** | Pointer to **string** | The name of the Supplier | [optional] 

## Methods

### NewPurchaseListElement

`func NewPurchaseListElement(pkiPurchaseID int32, fkiPeriodID int32, sPurchaseBillnumber string, dtPurchaseDate string, dtPurchasePaydate string, dPurchaseTotal string, bPurchaseIsfinanced bool, bPurchaseIspaid bool, ) *PurchaseListElement`

NewPurchaseListElement instantiates a new PurchaseListElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPurchaseListElementWithDefaults

`func NewPurchaseListElementWithDefaults() *PurchaseListElement`

NewPurchaseListElementWithDefaults instantiates a new PurchaseListElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiPurchaseID

`func (o *PurchaseListElement) GetPkiPurchaseID() int32`

GetPkiPurchaseID returns the PkiPurchaseID field if non-nil, zero value otherwise.

### GetPkiPurchaseIDOk

`func (o *PurchaseListElement) GetPkiPurchaseIDOk() (*int32, bool)`

GetPkiPurchaseIDOk returns a tuple with the PkiPurchaseID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiPurchaseID

`func (o *PurchaseListElement) SetPkiPurchaseID(v int32)`

SetPkiPurchaseID sets PkiPurchaseID field to given value.


### GetFkiFranchiseofficeID

`func (o *PurchaseListElement) GetFkiFranchiseofficeID() int32`

GetFkiFranchiseofficeID returns the FkiFranchiseofficeID field if non-nil, zero value otherwise.

### GetFkiFranchiseofficeIDOk

`func (o *PurchaseListElement) GetFkiFranchiseofficeIDOk() (*int32, bool)`

GetFkiFranchiseofficeIDOk returns a tuple with the FkiFranchiseofficeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiFranchiseofficeID

`func (o *PurchaseListElement) SetFkiFranchiseofficeID(v int32)`

SetFkiFranchiseofficeID sets FkiFranchiseofficeID field to given value.

### HasFkiFranchiseofficeID

`func (o *PurchaseListElement) HasFkiFranchiseofficeID() bool`

HasFkiFranchiseofficeID returns a boolean if a field has been set.

### GetFkiSupplierID

`func (o *PurchaseListElement) GetFkiSupplierID() int32`

GetFkiSupplierID returns the FkiSupplierID field if non-nil, zero value otherwise.

### GetFkiSupplierIDOk

`func (o *PurchaseListElement) GetFkiSupplierIDOk() (*int32, bool)`

GetFkiSupplierIDOk returns a tuple with the FkiSupplierID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiSupplierID

`func (o *PurchaseListElement) SetFkiSupplierID(v int32)`

SetFkiSupplierID sets FkiSupplierID field to given value.

### HasFkiSupplierID

`func (o *PurchaseListElement) HasFkiSupplierID() bool`

HasFkiSupplierID returns a boolean if a field has been set.

### GetFkiPeriodID

`func (o *PurchaseListElement) GetFkiPeriodID() int32`

GetFkiPeriodID returns the FkiPeriodID field if non-nil, zero value otherwise.

### GetFkiPeriodIDOk

`func (o *PurchaseListElement) GetFkiPeriodIDOk() (*int32, bool)`

GetFkiPeriodIDOk returns a tuple with the FkiPeriodID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiPeriodID

`func (o *PurchaseListElement) SetFkiPeriodID(v int32)`

SetFkiPeriodID sets FkiPeriodID field to given value.


### GetSPurchaseBillnumber

`func (o *PurchaseListElement) GetSPurchaseBillnumber() string`

GetSPurchaseBillnumber returns the SPurchaseBillnumber field if non-nil, zero value otherwise.

### GetSPurchaseBillnumberOk

`func (o *PurchaseListElement) GetSPurchaseBillnumberOk() (*string, bool)`

GetSPurchaseBillnumberOk returns a tuple with the SPurchaseBillnumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSPurchaseBillnumber

`func (o *PurchaseListElement) SetSPurchaseBillnumber(v string)`

SetSPurchaseBillnumber sets SPurchaseBillnumber field to given value.


### GetDtPurchaseDate

`func (o *PurchaseListElement) GetDtPurchaseDate() string`

GetDtPurchaseDate returns the DtPurchaseDate field if non-nil, zero value otherwise.

### GetDtPurchaseDateOk

`func (o *PurchaseListElement) GetDtPurchaseDateOk() (*string, bool)`

GetDtPurchaseDateOk returns a tuple with the DtPurchaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtPurchaseDate

`func (o *PurchaseListElement) SetDtPurchaseDate(v string)`

SetDtPurchaseDate sets DtPurchaseDate field to given value.


### GetDtPurchasePaydate

`func (o *PurchaseListElement) GetDtPurchasePaydate() string`

GetDtPurchasePaydate returns the DtPurchasePaydate field if non-nil, zero value otherwise.

### GetDtPurchasePaydateOk

`func (o *PurchaseListElement) GetDtPurchasePaydateOk() (*string, bool)`

GetDtPurchasePaydateOk returns a tuple with the DtPurchasePaydate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtPurchasePaydate

`func (o *PurchaseListElement) SetDtPurchasePaydate(v string)`

SetDtPurchasePaydate sets DtPurchasePaydate field to given value.


### GetDPurchaseTotal

`func (o *PurchaseListElement) GetDPurchaseTotal() string`

GetDPurchaseTotal returns the DPurchaseTotal field if non-nil, zero value otherwise.

### GetDPurchaseTotalOk

`func (o *PurchaseListElement) GetDPurchaseTotalOk() (*string, bool)`

GetDPurchaseTotalOk returns a tuple with the DPurchaseTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDPurchaseTotal

`func (o *PurchaseListElement) SetDPurchaseTotal(v string)`

SetDPurchaseTotal sets DPurchaseTotal field to given value.


### GetBPurchaseIsfinanced

`func (o *PurchaseListElement) GetBPurchaseIsfinanced() bool`

GetBPurchaseIsfinanced returns the BPurchaseIsfinanced field if non-nil, zero value otherwise.

### GetBPurchaseIsfinancedOk

`func (o *PurchaseListElement) GetBPurchaseIsfinancedOk() (*bool, bool)`

GetBPurchaseIsfinancedOk returns a tuple with the BPurchaseIsfinanced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBPurchaseIsfinanced

`func (o *PurchaseListElement) SetBPurchaseIsfinanced(v bool)`

SetBPurchaseIsfinanced sets BPurchaseIsfinanced field to given value.


### GetBPurchaseIspaid

`func (o *PurchaseListElement) GetBPurchaseIspaid() bool`

GetBPurchaseIspaid returns the BPurchaseIspaid field if non-nil, zero value otherwise.

### GetBPurchaseIspaidOk

`func (o *PurchaseListElement) GetBPurchaseIspaidOk() (*bool, bool)`

GetBPurchaseIspaidOk returns a tuple with the BPurchaseIspaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBPurchaseIspaid

`func (o *PurchaseListElement) SetBPurchaseIspaid(v bool)`

SetBPurchaseIspaid sets BPurchaseIspaid field to given value.


### GetSFranchiseagenceName

`func (o *PurchaseListElement) GetSFranchiseagenceName() string`

GetSFranchiseagenceName returns the SFranchiseagenceName field if non-nil, zero value otherwise.

### GetSFranchiseagenceNameOk

`func (o *PurchaseListElement) GetSFranchiseagenceNameOk() (*string, bool)`

GetSFranchiseagenceNameOk returns a tuple with the SFranchiseagenceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSFranchiseagenceName

`func (o *PurchaseListElement) SetSFranchiseagenceName(v string)`

SetSFranchiseagenceName sets SFranchiseagenceName field to given value.

### HasSFranchiseagenceName

`func (o *PurchaseListElement) HasSFranchiseagenceName() bool`

HasSFranchiseagenceName returns a boolean if a field has been set.

### GetIFranchiseofficeCode

`func (o *PurchaseListElement) GetIFranchiseofficeCode() string`

GetIFranchiseofficeCode returns the IFranchiseofficeCode field if non-nil, zero value otherwise.

### GetIFranchiseofficeCodeOk

`func (o *PurchaseListElement) GetIFranchiseofficeCodeOk() (*string, bool)`

GetIFranchiseofficeCodeOk returns a tuple with the IFranchiseofficeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIFranchiseofficeCode

`func (o *PurchaseListElement) SetIFranchiseofficeCode(v string)`

SetIFranchiseofficeCode sets IFranchiseofficeCode field to given value.

### HasIFranchiseofficeCode

`func (o *PurchaseListElement) HasIFranchiseofficeCode() bool`

HasIFranchiseofficeCode returns a boolean if a field has been set.

### GetSSupplierName

`func (o *PurchaseListElement) GetSSupplierName() string`

GetSSupplierName returns the SSupplierName field if non-nil, zero value otherwise.

### GetSSupplierNameOk

`func (o *PurchaseListElement) GetSSupplierNameOk() (*string, bool)`

GetSSupplierNameOk returns a tuple with the SSupplierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSSupplierName

`func (o *PurchaseListElement) SetSSupplierName(v string)`

SetSSupplierName sets SSupplierName field to given value.

### HasSSupplierName

`func (o *PurchaseListElement) HasSSupplierName() bool`

HasSSupplierName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


