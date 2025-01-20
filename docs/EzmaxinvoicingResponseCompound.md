# EzmaxinvoicingResponseCompound

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
**ObjEzmaxinvoicingcontract** | [**EzmaxinvoicingcontractResponseCompound**](EzmaxinvoicingcontractResponseCompound.md) |  | 
**ObjEzmaxpricing** | [**CustomEzmaxpricingResponse**](CustomEzmaxpricingResponse.md) |  | 
**AObjEzmaxinvoicingsummaryglobal** | [**[]EzmaxinvoicingsummaryglobalResponseCompound**](EzmaxinvoicingsummaryglobalResponseCompound.md) |  | 
**AObjEzmaxinvoicingsummaryexternal** | [**[]EzmaxinvoicingsummaryexternalResponseCompound**](EzmaxinvoicingsummaryexternalResponseCompound.md) |  | 
**AObjEzmaxinvoicingsummaryinternal** | [**[]EzmaxinvoicingsummaryinternalResponseCompound**](EzmaxinvoicingsummaryinternalResponseCompound.md) |  | 
**AObjEzmaxinvoicingagent** | [**[]EzmaxinvoicingagentResponseCompound**](EzmaxinvoicingagentResponseCompound.md) |  | 
**AObjEzmaxinvoicinguser** | [**[]EzmaxinvoicinguserResponseCompound**](EzmaxinvoicinguserResponseCompound.md) |  | 
**AObjEzmaxinvoicingezsignfolder** | **[]map[string]interface{}** |  | 
**AObjEzmaxinvoicingezsigndocument** | **[]map[string]interface{}** |  | 

## Methods

### NewEzmaxinvoicingResponseCompound

`func NewEzmaxinvoicingResponseCompound(fkiEzmaxinvoicingcontractID int32, fkiEzmaxpricingID int32, fkiSystemconfigurationtypeID int32, sSystemconfigurationtypeDescriptionX string, yyyymmEzmaxinvoicing string, iEzmaxinvoicingDays int32, eEzmaxinvoicingPaymenttype FieldEEzmaxinvoicingPaymenttype, dEzmaxinvoicingRebatepaymenttype string, iEzmaxinvoicingContractlength int32, dEzmaxinvoicingRebatecontractlength string, bEzmaxinvoicingRebateEzsignallagents bool, objEzmaxinvoicingcontract EzmaxinvoicingcontractResponseCompound, objEzmaxpricing CustomEzmaxpricingResponse, aObjEzmaxinvoicingsummaryglobal []EzmaxinvoicingsummaryglobalResponseCompound, aObjEzmaxinvoicingsummaryexternal []EzmaxinvoicingsummaryexternalResponseCompound, aObjEzmaxinvoicingsummaryinternal []EzmaxinvoicingsummaryinternalResponseCompound, aObjEzmaxinvoicingagent []EzmaxinvoicingagentResponseCompound, aObjEzmaxinvoicinguser []EzmaxinvoicinguserResponseCompound, aObjEzmaxinvoicingezsignfolder []CustomEzmaxinvoicingEzsignfolderResponse, aObjEzmaxinvoicingezsigndocument []CustomEzmaxinvoicingEzsigndocumentResponse, ) *EzmaxinvoicingResponseCompound`

NewEzmaxinvoicingResponseCompound instantiates a new EzmaxinvoicingResponseCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzmaxinvoicingResponseCompoundWithDefaults

`func NewEzmaxinvoicingResponseCompoundWithDefaults() *EzmaxinvoicingResponseCompound`

NewEzmaxinvoicingResponseCompoundWithDefaults instantiates a new EzmaxinvoicingResponseCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzmaxinvoicingID

`func (o *EzmaxinvoicingResponseCompound) GetPkiEzmaxinvoicingID() int32`

GetPkiEzmaxinvoicingID returns the PkiEzmaxinvoicingID field if non-nil, zero value otherwise.

### GetPkiEzmaxinvoicingIDOk

`func (o *EzmaxinvoicingResponseCompound) GetPkiEzmaxinvoicingIDOk() (*int32, bool)`

GetPkiEzmaxinvoicingIDOk returns a tuple with the PkiEzmaxinvoicingID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzmaxinvoicingID

`func (o *EzmaxinvoicingResponseCompound) SetPkiEzmaxinvoicingID(v int32)`

SetPkiEzmaxinvoicingID sets PkiEzmaxinvoicingID field to given value.

### HasPkiEzmaxinvoicingID

`func (o *EzmaxinvoicingResponseCompound) HasPkiEzmaxinvoicingID() bool`

HasPkiEzmaxinvoicingID returns a boolean if a field has been set.

### GetFkiEzmaxinvoicingcontractID

`func (o *EzmaxinvoicingResponseCompound) GetFkiEzmaxinvoicingcontractID() int32`

GetFkiEzmaxinvoicingcontractID returns the FkiEzmaxinvoicingcontractID field if non-nil, zero value otherwise.

### GetFkiEzmaxinvoicingcontractIDOk

`func (o *EzmaxinvoicingResponseCompound) GetFkiEzmaxinvoicingcontractIDOk() (*int32, bool)`

GetFkiEzmaxinvoicingcontractIDOk returns a tuple with the FkiEzmaxinvoicingcontractID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzmaxinvoicingcontractID

`func (o *EzmaxinvoicingResponseCompound) SetFkiEzmaxinvoicingcontractID(v int32)`

SetFkiEzmaxinvoicingcontractID sets FkiEzmaxinvoicingcontractID field to given value.


### GetFkiEzmaxpricingID

`func (o *EzmaxinvoicingResponseCompound) GetFkiEzmaxpricingID() int32`

GetFkiEzmaxpricingID returns the FkiEzmaxpricingID field if non-nil, zero value otherwise.

### GetFkiEzmaxpricingIDOk

`func (o *EzmaxinvoicingResponseCompound) GetFkiEzmaxpricingIDOk() (*int32, bool)`

GetFkiEzmaxpricingIDOk returns a tuple with the FkiEzmaxpricingID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzmaxpricingID

`func (o *EzmaxinvoicingResponseCompound) SetFkiEzmaxpricingID(v int32)`

SetFkiEzmaxpricingID sets FkiEzmaxpricingID field to given value.


### GetFkiSystemconfigurationtypeID

`func (o *EzmaxinvoicingResponseCompound) GetFkiSystemconfigurationtypeID() int32`

GetFkiSystemconfigurationtypeID returns the FkiSystemconfigurationtypeID field if non-nil, zero value otherwise.

### GetFkiSystemconfigurationtypeIDOk

`func (o *EzmaxinvoicingResponseCompound) GetFkiSystemconfigurationtypeIDOk() (*int32, bool)`

GetFkiSystemconfigurationtypeIDOk returns a tuple with the FkiSystemconfigurationtypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiSystemconfigurationtypeID

`func (o *EzmaxinvoicingResponseCompound) SetFkiSystemconfigurationtypeID(v int32)`

SetFkiSystemconfigurationtypeID sets FkiSystemconfigurationtypeID field to given value.


### GetSSystemconfigurationtypeDescriptionX

`func (o *EzmaxinvoicingResponseCompound) GetSSystemconfigurationtypeDescriptionX() string`

GetSSystemconfigurationtypeDescriptionX returns the SSystemconfigurationtypeDescriptionX field if non-nil, zero value otherwise.

### GetSSystemconfigurationtypeDescriptionXOk

`func (o *EzmaxinvoicingResponseCompound) GetSSystemconfigurationtypeDescriptionXOk() (*string, bool)`

GetSSystemconfigurationtypeDescriptionXOk returns a tuple with the SSystemconfigurationtypeDescriptionX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSSystemconfigurationtypeDescriptionX

`func (o *EzmaxinvoicingResponseCompound) SetSSystemconfigurationtypeDescriptionX(v string)`

SetSSystemconfigurationtypeDescriptionX sets SSystemconfigurationtypeDescriptionX field to given value.


### GetYyyymmEzmaxinvoicing

`func (o *EzmaxinvoicingResponseCompound) GetYyyymmEzmaxinvoicing() string`

GetYyyymmEzmaxinvoicing returns the YyyymmEzmaxinvoicing field if non-nil, zero value otherwise.

### GetYyyymmEzmaxinvoicingOk

`func (o *EzmaxinvoicingResponseCompound) GetYyyymmEzmaxinvoicingOk() (*string, bool)`

GetYyyymmEzmaxinvoicingOk returns a tuple with the YyyymmEzmaxinvoicing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYyyymmEzmaxinvoicing

`func (o *EzmaxinvoicingResponseCompound) SetYyyymmEzmaxinvoicing(v string)`

SetYyyymmEzmaxinvoicing sets YyyymmEzmaxinvoicing field to given value.


### GetIEzmaxinvoicingDays

`func (o *EzmaxinvoicingResponseCompound) GetIEzmaxinvoicingDays() int32`

GetIEzmaxinvoicingDays returns the IEzmaxinvoicingDays field if non-nil, zero value otherwise.

### GetIEzmaxinvoicingDaysOk

`func (o *EzmaxinvoicingResponseCompound) GetIEzmaxinvoicingDaysOk() (*int32, bool)`

GetIEzmaxinvoicingDaysOk returns a tuple with the IEzmaxinvoicingDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzmaxinvoicingDays

`func (o *EzmaxinvoicingResponseCompound) SetIEzmaxinvoicingDays(v int32)`

SetIEzmaxinvoicingDays sets IEzmaxinvoicingDays field to given value.


### GetEEzmaxinvoicingPaymenttype

`func (o *EzmaxinvoicingResponseCompound) GetEEzmaxinvoicingPaymenttype() FieldEEzmaxinvoicingPaymenttype`

GetEEzmaxinvoicingPaymenttype returns the EEzmaxinvoicingPaymenttype field if non-nil, zero value otherwise.

### GetEEzmaxinvoicingPaymenttypeOk

`func (o *EzmaxinvoicingResponseCompound) GetEEzmaxinvoicingPaymenttypeOk() (*FieldEEzmaxinvoicingPaymenttype, bool)`

GetEEzmaxinvoicingPaymenttypeOk returns a tuple with the EEzmaxinvoicingPaymenttype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzmaxinvoicingPaymenttype

`func (o *EzmaxinvoicingResponseCompound) SetEEzmaxinvoicingPaymenttype(v FieldEEzmaxinvoicingPaymenttype)`

SetEEzmaxinvoicingPaymenttype sets EEzmaxinvoicingPaymenttype field to given value.


### GetDEzmaxinvoicingRebatepaymenttype

`func (o *EzmaxinvoicingResponseCompound) GetDEzmaxinvoicingRebatepaymenttype() string`

GetDEzmaxinvoicingRebatepaymenttype returns the DEzmaxinvoicingRebatepaymenttype field if non-nil, zero value otherwise.

### GetDEzmaxinvoicingRebatepaymenttypeOk

`func (o *EzmaxinvoicingResponseCompound) GetDEzmaxinvoicingRebatepaymenttypeOk() (*string, bool)`

GetDEzmaxinvoicingRebatepaymenttypeOk returns a tuple with the DEzmaxinvoicingRebatepaymenttype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDEzmaxinvoicingRebatepaymenttype

`func (o *EzmaxinvoicingResponseCompound) SetDEzmaxinvoicingRebatepaymenttype(v string)`

SetDEzmaxinvoicingRebatepaymenttype sets DEzmaxinvoicingRebatepaymenttype field to given value.


### GetIEzmaxinvoicingContractlength

`func (o *EzmaxinvoicingResponseCompound) GetIEzmaxinvoicingContractlength() int32`

GetIEzmaxinvoicingContractlength returns the IEzmaxinvoicingContractlength field if non-nil, zero value otherwise.

### GetIEzmaxinvoicingContractlengthOk

`func (o *EzmaxinvoicingResponseCompound) GetIEzmaxinvoicingContractlengthOk() (*int32, bool)`

GetIEzmaxinvoicingContractlengthOk returns a tuple with the IEzmaxinvoicingContractlength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzmaxinvoicingContractlength

`func (o *EzmaxinvoicingResponseCompound) SetIEzmaxinvoicingContractlength(v int32)`

SetIEzmaxinvoicingContractlength sets IEzmaxinvoicingContractlength field to given value.


### GetDEzmaxinvoicingRebatecontractlength

`func (o *EzmaxinvoicingResponseCompound) GetDEzmaxinvoicingRebatecontractlength() string`

GetDEzmaxinvoicingRebatecontractlength returns the DEzmaxinvoicingRebatecontractlength field if non-nil, zero value otherwise.

### GetDEzmaxinvoicingRebatecontractlengthOk

`func (o *EzmaxinvoicingResponseCompound) GetDEzmaxinvoicingRebatecontractlengthOk() (*string, bool)`

GetDEzmaxinvoicingRebatecontractlengthOk returns a tuple with the DEzmaxinvoicingRebatecontractlength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDEzmaxinvoicingRebatecontractlength

`func (o *EzmaxinvoicingResponseCompound) SetDEzmaxinvoicingRebatecontractlength(v string)`

SetDEzmaxinvoicingRebatecontractlength sets DEzmaxinvoicingRebatecontractlength field to given value.


### GetBEzmaxinvoicingRebateEzsignallagents

`func (o *EzmaxinvoicingResponseCompound) GetBEzmaxinvoicingRebateEzsignallagents() bool`

GetBEzmaxinvoicingRebateEzsignallagents returns the BEzmaxinvoicingRebateEzsignallagents field if non-nil, zero value otherwise.

### GetBEzmaxinvoicingRebateEzsignallagentsOk

`func (o *EzmaxinvoicingResponseCompound) GetBEzmaxinvoicingRebateEzsignallagentsOk() (*bool, bool)`

GetBEzmaxinvoicingRebateEzsignallagentsOk returns a tuple with the BEzmaxinvoicingRebateEzsignallagents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzmaxinvoicingRebateEzsignallagents

`func (o *EzmaxinvoicingResponseCompound) SetBEzmaxinvoicingRebateEzsignallagents(v bool)`

SetBEzmaxinvoicingRebateEzsignallagents sets BEzmaxinvoicingRebateEzsignallagents field to given value.


### GetObjAudit

`func (o *EzmaxinvoicingResponseCompound) GetObjAudit() CommonAudit`

GetObjAudit returns the ObjAudit field if non-nil, zero value otherwise.

### GetObjAuditOk

`func (o *EzmaxinvoicingResponseCompound) GetObjAuditOk() (*CommonAudit, bool)`

GetObjAuditOk returns a tuple with the ObjAudit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjAudit

`func (o *EzmaxinvoicingResponseCompound) SetObjAudit(v CommonAudit)`

SetObjAudit sets ObjAudit field to given value.

### HasObjAudit

`func (o *EzmaxinvoicingResponseCompound) HasObjAudit() bool`

HasObjAudit returns a boolean if a field has been set.

### GetObjEzmaxinvoicingcontract

`func (o *EzmaxinvoicingResponseCompound) GetObjEzmaxinvoicingcontract() EzmaxinvoicingcontractResponseCompound`

GetObjEzmaxinvoicingcontract returns the ObjEzmaxinvoicingcontract field if non-nil, zero value otherwise.

### GetObjEzmaxinvoicingcontractOk

`func (o *EzmaxinvoicingResponseCompound) GetObjEzmaxinvoicingcontractOk() (*EzmaxinvoicingcontractResponseCompound, bool)`

GetObjEzmaxinvoicingcontractOk returns a tuple with the ObjEzmaxinvoicingcontract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjEzmaxinvoicingcontract

`func (o *EzmaxinvoicingResponseCompound) SetObjEzmaxinvoicingcontract(v EzmaxinvoicingcontractResponseCompound)`

SetObjEzmaxinvoicingcontract sets ObjEzmaxinvoicingcontract field to given value.


### GetObjEzmaxpricing

`func (o *EzmaxinvoicingResponseCompound) GetObjEzmaxpricing() CustomEzmaxpricingResponse`

GetObjEzmaxpricing returns the ObjEzmaxpricing field if non-nil, zero value otherwise.

### GetObjEzmaxpricingOk

`func (o *EzmaxinvoicingResponseCompound) GetObjEzmaxpricingOk() (*CustomEzmaxpricingResponse, bool)`

GetObjEzmaxpricingOk returns a tuple with the ObjEzmaxpricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjEzmaxpricing

`func (o *EzmaxinvoicingResponseCompound) SetObjEzmaxpricing(v CustomEzmaxpricingResponse)`

SetObjEzmaxpricing sets ObjEzmaxpricing field to given value.


### GetAObjEzmaxinvoicingsummaryglobal

`func (o *EzmaxinvoicingResponseCompound) GetAObjEzmaxinvoicingsummaryglobal() []EzmaxinvoicingsummaryglobalResponseCompound`

GetAObjEzmaxinvoicingsummaryglobal returns the AObjEzmaxinvoicingsummaryglobal field if non-nil, zero value otherwise.

### GetAObjEzmaxinvoicingsummaryglobalOk

`func (o *EzmaxinvoicingResponseCompound) GetAObjEzmaxinvoicingsummaryglobalOk() (*[]EzmaxinvoicingsummaryglobalResponseCompound, bool)`

GetAObjEzmaxinvoicingsummaryglobalOk returns a tuple with the AObjEzmaxinvoicingsummaryglobal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEzmaxinvoicingsummaryglobal

`func (o *EzmaxinvoicingResponseCompound) SetAObjEzmaxinvoicingsummaryglobal(v []EzmaxinvoicingsummaryglobalResponseCompound)`

SetAObjEzmaxinvoicingsummaryglobal sets AObjEzmaxinvoicingsummaryglobal field to given value.


### GetAObjEzmaxinvoicingsummaryexternal

`func (o *EzmaxinvoicingResponseCompound) GetAObjEzmaxinvoicingsummaryexternal() []EzmaxinvoicingsummaryexternalResponseCompound`

GetAObjEzmaxinvoicingsummaryexternal returns the AObjEzmaxinvoicingsummaryexternal field if non-nil, zero value otherwise.

### GetAObjEzmaxinvoicingsummaryexternalOk

`func (o *EzmaxinvoicingResponseCompound) GetAObjEzmaxinvoicingsummaryexternalOk() (*[]EzmaxinvoicingsummaryexternalResponseCompound, bool)`

GetAObjEzmaxinvoicingsummaryexternalOk returns a tuple with the AObjEzmaxinvoicingsummaryexternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEzmaxinvoicingsummaryexternal

`func (o *EzmaxinvoicingResponseCompound) SetAObjEzmaxinvoicingsummaryexternal(v []EzmaxinvoicingsummaryexternalResponseCompound)`

SetAObjEzmaxinvoicingsummaryexternal sets AObjEzmaxinvoicingsummaryexternal field to given value.


### GetAObjEzmaxinvoicingsummaryinternal

`func (o *EzmaxinvoicingResponseCompound) GetAObjEzmaxinvoicingsummaryinternal() []EzmaxinvoicingsummaryinternalResponseCompound`

GetAObjEzmaxinvoicingsummaryinternal returns the AObjEzmaxinvoicingsummaryinternal field if non-nil, zero value otherwise.

### GetAObjEzmaxinvoicingsummaryinternalOk

`func (o *EzmaxinvoicingResponseCompound) GetAObjEzmaxinvoicingsummaryinternalOk() (*[]EzmaxinvoicingsummaryinternalResponseCompound, bool)`

GetAObjEzmaxinvoicingsummaryinternalOk returns a tuple with the AObjEzmaxinvoicingsummaryinternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEzmaxinvoicingsummaryinternal

`func (o *EzmaxinvoicingResponseCompound) SetAObjEzmaxinvoicingsummaryinternal(v []EzmaxinvoicingsummaryinternalResponseCompound)`

SetAObjEzmaxinvoicingsummaryinternal sets AObjEzmaxinvoicingsummaryinternal field to given value.


### GetAObjEzmaxinvoicingagent

`func (o *EzmaxinvoicingResponseCompound) GetAObjEzmaxinvoicingagent() []EzmaxinvoicingagentResponseCompound`

GetAObjEzmaxinvoicingagent returns the AObjEzmaxinvoicingagent field if non-nil, zero value otherwise.

### GetAObjEzmaxinvoicingagentOk

`func (o *EzmaxinvoicingResponseCompound) GetAObjEzmaxinvoicingagentOk() (*[]EzmaxinvoicingagentResponseCompound, bool)`

GetAObjEzmaxinvoicingagentOk returns a tuple with the AObjEzmaxinvoicingagent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEzmaxinvoicingagent

`func (o *EzmaxinvoicingResponseCompound) SetAObjEzmaxinvoicingagent(v []EzmaxinvoicingagentResponseCompound)`

SetAObjEzmaxinvoicingagent sets AObjEzmaxinvoicingagent field to given value.


### GetAObjEzmaxinvoicinguser

`func (o *EzmaxinvoicingResponseCompound) GetAObjEzmaxinvoicinguser() []EzmaxinvoicinguserResponseCompound`

GetAObjEzmaxinvoicinguser returns the AObjEzmaxinvoicinguser field if non-nil, zero value otherwise.

### GetAObjEzmaxinvoicinguserOk

`func (o *EzmaxinvoicingResponseCompound) GetAObjEzmaxinvoicinguserOk() (*[]EzmaxinvoicinguserResponseCompound, bool)`

GetAObjEzmaxinvoicinguserOk returns a tuple with the AObjEzmaxinvoicinguser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEzmaxinvoicinguser

`func (o *EzmaxinvoicingResponseCompound) SetAObjEzmaxinvoicinguser(v []EzmaxinvoicinguserResponseCompound)`

SetAObjEzmaxinvoicinguser sets AObjEzmaxinvoicinguser field to given value.


### GetAObjEzmaxinvoicingezsignfolder

`func (o *EzmaxinvoicingResponseCompound) GetAObjEzmaxinvoicingezsignfolder() []CustomEzmaxinvoicingEzsignfolderResponse`

GetAObjEzmaxinvoicingezsignfolder returns the AObjEzmaxinvoicingezsignfolder field if non-nil, zero value otherwise.

### GetAObjEzmaxinvoicingezsignfolderOk

`func (o *EzmaxinvoicingResponseCompound) GetAObjEzmaxinvoicingezsignfolderOk() (*[]CustomEzmaxinvoicingEzsignfolderResponse, bool)`

GetAObjEzmaxinvoicingezsignfolderOk returns a tuple with the AObjEzmaxinvoicingezsignfolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEzmaxinvoicingezsignfolder

`func (o *EzmaxinvoicingResponseCompound) SetAObjEzmaxinvoicingezsignfolder(v []CustomEzmaxinvoicingEzsignfolderResponse)`

SetAObjEzmaxinvoicingezsignfolder sets AObjEzmaxinvoicingezsignfolder field to given value.


### GetAObjEzmaxinvoicingezsigndocument

`func (o *EzmaxinvoicingResponseCompound) GetAObjEzmaxinvoicingezsigndocument() []CustomEzmaxinvoicingEzsigndocumentResponse`

GetAObjEzmaxinvoicingezsigndocument returns the AObjEzmaxinvoicingezsigndocument field if non-nil, zero value otherwise.

### GetAObjEzmaxinvoicingezsigndocumentOk

`func (o *EzmaxinvoicingResponseCompound) GetAObjEzmaxinvoicingezsigndocumentOk() (*[]CustomEzmaxinvoicingEzsigndocumentResponse, bool)`

GetAObjEzmaxinvoicingezsigndocumentOk returns a tuple with the AObjEzmaxinvoicingezsigndocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEzmaxinvoicingezsigndocument

`func (o *EzmaxinvoicingResponseCompound) SetAObjEzmaxinvoicingezsigndocument(v []CustomEzmaxinvoicingEzsigndocumentResponse)`

SetAObjEzmaxinvoicingezsigndocument sets AObjEzmaxinvoicingezsigndocument field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


