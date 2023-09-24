# EzmaxinvoicingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzmaxinvoicingID** | Pointer to **int32** | The unique ID of the Ezmaxinvoicing | [optional] 
**FkiEzmaxinvoicingcontractID** | **int32** | The unique ID of the Ezmaxinvoicingcontract | 
**FkiEzmaxpricingID** | **int32** | The unique ID of the Ezmaxpricing | 
**FkiSystemconfigurationtypeID** | **int32** | The unique ID of the Systemconfigurationtype | 
**SSystemconfigurationtypeDescriptionX** | **string** | The description of the Systemconfigurationtype in the language of the requester | 
**YyyymmEzmaxinvoicing** | **string** | The YYYYMM period of the Ezmaxinvoicing | 
**IEzmaxinvoicingDays** | **int32** | The number of days invoiced | 
**EEzmaxinvoicingPaymenttype** | [**FieldEEzmaxinvoicingPaymenttype**](FieldEEzmaxinvoicingPaymenttype.md) |  | 
**DEzmaxinvoicingRebatepaymenttype** | **string** | The percentage of rebate depending of the payment type | 
**IEzmaxinvoicingContractlength** | **int32** | The length of the contract in years | 
**DEzmaxinvoicingRebatecontractlength** | **string** | The percentage of rebate depending of the contract length | 
**BEzmaxinvoicingRebateEzsignallagents** | **bool** | Whether the rebate for eZsign is for all agents | 
**ObjAudit** | Pointer to [**CommonAudit**](CommonAudit.md) |  | [optional] 

## Methods

### NewEzmaxinvoicingResponse

`func NewEzmaxinvoicingResponse(fkiEzmaxinvoicingcontractID int32, fkiEzmaxpricingID int32, fkiSystemconfigurationtypeID int32, sSystemconfigurationtypeDescriptionX string, yyyymmEzmaxinvoicing string, iEzmaxinvoicingDays int32, eEzmaxinvoicingPaymenttype FieldEEzmaxinvoicingPaymenttype, dEzmaxinvoicingRebatepaymenttype string, iEzmaxinvoicingContractlength int32, dEzmaxinvoicingRebatecontractlength string, bEzmaxinvoicingRebateEzsignallagents bool, ) *EzmaxinvoicingResponse`

NewEzmaxinvoicingResponse instantiates a new EzmaxinvoicingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzmaxinvoicingResponseWithDefaults

`func NewEzmaxinvoicingResponseWithDefaults() *EzmaxinvoicingResponse`

NewEzmaxinvoicingResponseWithDefaults instantiates a new EzmaxinvoicingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzmaxinvoicingID

`func (o *EzmaxinvoicingResponse) GetPkiEzmaxinvoicingID() int32`

GetPkiEzmaxinvoicingID returns the PkiEzmaxinvoicingID field if non-nil, zero value otherwise.

### GetPkiEzmaxinvoicingIDOk

`func (o *EzmaxinvoicingResponse) GetPkiEzmaxinvoicingIDOk() (*int32, bool)`

GetPkiEzmaxinvoicingIDOk returns a tuple with the PkiEzmaxinvoicingID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzmaxinvoicingID

`func (o *EzmaxinvoicingResponse) SetPkiEzmaxinvoicingID(v int32)`

SetPkiEzmaxinvoicingID sets PkiEzmaxinvoicingID field to given value.

### HasPkiEzmaxinvoicingID

`func (o *EzmaxinvoicingResponse) HasPkiEzmaxinvoicingID() bool`

HasPkiEzmaxinvoicingID returns a boolean if a field has been set.

### GetFkiEzmaxinvoicingcontractID

`func (o *EzmaxinvoicingResponse) GetFkiEzmaxinvoicingcontractID() int32`

GetFkiEzmaxinvoicingcontractID returns the FkiEzmaxinvoicingcontractID field if non-nil, zero value otherwise.

### GetFkiEzmaxinvoicingcontractIDOk

`func (o *EzmaxinvoicingResponse) GetFkiEzmaxinvoicingcontractIDOk() (*int32, bool)`

GetFkiEzmaxinvoicingcontractIDOk returns a tuple with the FkiEzmaxinvoicingcontractID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzmaxinvoicingcontractID

`func (o *EzmaxinvoicingResponse) SetFkiEzmaxinvoicingcontractID(v int32)`

SetFkiEzmaxinvoicingcontractID sets FkiEzmaxinvoicingcontractID field to given value.


### GetFkiEzmaxpricingID

`func (o *EzmaxinvoicingResponse) GetFkiEzmaxpricingID() int32`

GetFkiEzmaxpricingID returns the FkiEzmaxpricingID field if non-nil, zero value otherwise.

### GetFkiEzmaxpricingIDOk

`func (o *EzmaxinvoicingResponse) GetFkiEzmaxpricingIDOk() (*int32, bool)`

GetFkiEzmaxpricingIDOk returns a tuple with the FkiEzmaxpricingID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzmaxpricingID

`func (o *EzmaxinvoicingResponse) SetFkiEzmaxpricingID(v int32)`

SetFkiEzmaxpricingID sets FkiEzmaxpricingID field to given value.


### GetFkiSystemconfigurationtypeID

`func (o *EzmaxinvoicingResponse) GetFkiSystemconfigurationtypeID() int32`

GetFkiSystemconfigurationtypeID returns the FkiSystemconfigurationtypeID field if non-nil, zero value otherwise.

### GetFkiSystemconfigurationtypeIDOk

`func (o *EzmaxinvoicingResponse) GetFkiSystemconfigurationtypeIDOk() (*int32, bool)`

GetFkiSystemconfigurationtypeIDOk returns a tuple with the FkiSystemconfigurationtypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiSystemconfigurationtypeID

`func (o *EzmaxinvoicingResponse) SetFkiSystemconfigurationtypeID(v int32)`

SetFkiSystemconfigurationtypeID sets FkiSystemconfigurationtypeID field to given value.


### GetSSystemconfigurationtypeDescriptionX

`func (o *EzmaxinvoicingResponse) GetSSystemconfigurationtypeDescriptionX() string`

GetSSystemconfigurationtypeDescriptionX returns the SSystemconfigurationtypeDescriptionX field if non-nil, zero value otherwise.

### GetSSystemconfigurationtypeDescriptionXOk

`func (o *EzmaxinvoicingResponse) GetSSystemconfigurationtypeDescriptionXOk() (*string, bool)`

GetSSystemconfigurationtypeDescriptionXOk returns a tuple with the SSystemconfigurationtypeDescriptionX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSSystemconfigurationtypeDescriptionX

`func (o *EzmaxinvoicingResponse) SetSSystemconfigurationtypeDescriptionX(v string)`

SetSSystemconfigurationtypeDescriptionX sets SSystemconfigurationtypeDescriptionX field to given value.


### GetYyyymmEzmaxinvoicing

`func (o *EzmaxinvoicingResponse) GetYyyymmEzmaxinvoicing() string`

GetYyyymmEzmaxinvoicing returns the YyyymmEzmaxinvoicing field if non-nil, zero value otherwise.

### GetYyyymmEzmaxinvoicingOk

`func (o *EzmaxinvoicingResponse) GetYyyymmEzmaxinvoicingOk() (*string, bool)`

GetYyyymmEzmaxinvoicingOk returns a tuple with the YyyymmEzmaxinvoicing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYyyymmEzmaxinvoicing

`func (o *EzmaxinvoicingResponse) SetYyyymmEzmaxinvoicing(v string)`

SetYyyymmEzmaxinvoicing sets YyyymmEzmaxinvoicing field to given value.


### GetIEzmaxinvoicingDays

`func (o *EzmaxinvoicingResponse) GetIEzmaxinvoicingDays() int32`

GetIEzmaxinvoicingDays returns the IEzmaxinvoicingDays field if non-nil, zero value otherwise.

### GetIEzmaxinvoicingDaysOk

`func (o *EzmaxinvoicingResponse) GetIEzmaxinvoicingDaysOk() (*int32, bool)`

GetIEzmaxinvoicingDaysOk returns a tuple with the IEzmaxinvoicingDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzmaxinvoicingDays

`func (o *EzmaxinvoicingResponse) SetIEzmaxinvoicingDays(v int32)`

SetIEzmaxinvoicingDays sets IEzmaxinvoicingDays field to given value.


### GetEEzmaxinvoicingPaymenttype

`func (o *EzmaxinvoicingResponse) GetEEzmaxinvoicingPaymenttype() FieldEEzmaxinvoicingPaymenttype`

GetEEzmaxinvoicingPaymenttype returns the EEzmaxinvoicingPaymenttype field if non-nil, zero value otherwise.

### GetEEzmaxinvoicingPaymenttypeOk

`func (o *EzmaxinvoicingResponse) GetEEzmaxinvoicingPaymenttypeOk() (*FieldEEzmaxinvoicingPaymenttype, bool)`

GetEEzmaxinvoicingPaymenttypeOk returns a tuple with the EEzmaxinvoicingPaymenttype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzmaxinvoicingPaymenttype

`func (o *EzmaxinvoicingResponse) SetEEzmaxinvoicingPaymenttype(v FieldEEzmaxinvoicingPaymenttype)`

SetEEzmaxinvoicingPaymenttype sets EEzmaxinvoicingPaymenttype field to given value.


### GetDEzmaxinvoicingRebatepaymenttype

`func (o *EzmaxinvoicingResponse) GetDEzmaxinvoicingRebatepaymenttype() string`

GetDEzmaxinvoicingRebatepaymenttype returns the DEzmaxinvoicingRebatepaymenttype field if non-nil, zero value otherwise.

### GetDEzmaxinvoicingRebatepaymenttypeOk

`func (o *EzmaxinvoicingResponse) GetDEzmaxinvoicingRebatepaymenttypeOk() (*string, bool)`

GetDEzmaxinvoicingRebatepaymenttypeOk returns a tuple with the DEzmaxinvoicingRebatepaymenttype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDEzmaxinvoicingRebatepaymenttype

`func (o *EzmaxinvoicingResponse) SetDEzmaxinvoicingRebatepaymenttype(v string)`

SetDEzmaxinvoicingRebatepaymenttype sets DEzmaxinvoicingRebatepaymenttype field to given value.


### GetIEzmaxinvoicingContractlength

`func (o *EzmaxinvoicingResponse) GetIEzmaxinvoicingContractlength() int32`

GetIEzmaxinvoicingContractlength returns the IEzmaxinvoicingContractlength field if non-nil, zero value otherwise.

### GetIEzmaxinvoicingContractlengthOk

`func (o *EzmaxinvoicingResponse) GetIEzmaxinvoicingContractlengthOk() (*int32, bool)`

GetIEzmaxinvoicingContractlengthOk returns a tuple with the IEzmaxinvoicingContractlength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzmaxinvoicingContractlength

`func (o *EzmaxinvoicingResponse) SetIEzmaxinvoicingContractlength(v int32)`

SetIEzmaxinvoicingContractlength sets IEzmaxinvoicingContractlength field to given value.


### GetDEzmaxinvoicingRebatecontractlength

`func (o *EzmaxinvoicingResponse) GetDEzmaxinvoicingRebatecontractlength() string`

GetDEzmaxinvoicingRebatecontractlength returns the DEzmaxinvoicingRebatecontractlength field if non-nil, zero value otherwise.

### GetDEzmaxinvoicingRebatecontractlengthOk

`func (o *EzmaxinvoicingResponse) GetDEzmaxinvoicingRebatecontractlengthOk() (*string, bool)`

GetDEzmaxinvoicingRebatecontractlengthOk returns a tuple with the DEzmaxinvoicingRebatecontractlength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDEzmaxinvoicingRebatecontractlength

`func (o *EzmaxinvoicingResponse) SetDEzmaxinvoicingRebatecontractlength(v string)`

SetDEzmaxinvoicingRebatecontractlength sets DEzmaxinvoicingRebatecontractlength field to given value.


### GetBEzmaxinvoicingRebateEzsignallagents

`func (o *EzmaxinvoicingResponse) GetBEzmaxinvoicingRebateEzsignallagents() bool`

GetBEzmaxinvoicingRebateEzsignallagents returns the BEzmaxinvoicingRebateEzsignallagents field if non-nil, zero value otherwise.

### GetBEzmaxinvoicingRebateEzsignallagentsOk

`func (o *EzmaxinvoicingResponse) GetBEzmaxinvoicingRebateEzsignallagentsOk() (*bool, bool)`

GetBEzmaxinvoicingRebateEzsignallagentsOk returns a tuple with the BEzmaxinvoicingRebateEzsignallagents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzmaxinvoicingRebateEzsignallagents

`func (o *EzmaxinvoicingResponse) SetBEzmaxinvoicingRebateEzsignallagents(v bool)`

SetBEzmaxinvoicingRebateEzsignallagents sets BEzmaxinvoicingRebateEzsignallagents field to given value.


### GetObjAudit

`func (o *EzmaxinvoicingResponse) GetObjAudit() CommonAudit`

GetObjAudit returns the ObjAudit field if non-nil, zero value otherwise.

### GetObjAuditOk

`func (o *EzmaxinvoicingResponse) GetObjAuditOk() (*CommonAudit, bool)`

GetObjAuditOk returns a tuple with the ObjAudit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjAudit

`func (o *EzmaxinvoicingResponse) SetObjAudit(v CommonAudit)`

SetObjAudit sets ObjAudit field to given value.

### HasObjAudit

`func (o *EzmaxinvoicingResponse) HasObjAudit() bool`

HasObjAudit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


