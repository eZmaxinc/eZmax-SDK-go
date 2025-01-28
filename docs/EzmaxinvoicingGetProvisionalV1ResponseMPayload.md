# EzmaxinvoicingGetProvisionalV1ResponseMPayload

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
**AObjEzmaxinvoicingezsignfolder** | [**[]CustomEzmaxinvoicingEzsignfolderResponse**](CustomEzmaxinvoicingEzsignfolderResponse.md) |  | 
**AObjEzmaxinvoicingezsigndocument** | [**[]CustomEzmaxinvoicingEzsigndocumentResponse**](CustomEzmaxinvoicingEzsigndocumentResponse.md) |  | 

## Methods

### NewEzmaxinvoicingGetProvisionalV1ResponseMPayload

`func NewEzmaxinvoicingGetProvisionalV1ResponseMPayload(fkiEzmaxinvoicingcontractID int32, fkiEzmaxpricingID int32, fkiSystemconfigurationtypeID int32, sSystemconfigurationtypeDescriptionX string, yyyymmEzmaxinvoicing string, iEzmaxinvoicingDays int32, eEzmaxinvoicingPaymenttype FieldEEzmaxinvoicingPaymenttype, dEzmaxinvoicingRebatepaymenttype string, iEzmaxinvoicingContractlength int32, dEzmaxinvoicingRebatecontractlength string, bEzmaxinvoicingRebateEzsignallagents bool, objEzmaxinvoicingcontract EzmaxinvoicingcontractResponseCompound, objEzmaxpricing CustomEzmaxpricingResponse, aObjEzmaxinvoicingsummaryglobal []EzmaxinvoicingsummaryglobalResponseCompound, aObjEzmaxinvoicingsummaryexternal []EzmaxinvoicingsummaryexternalResponseCompound, aObjEzmaxinvoicingsummaryinternal []EzmaxinvoicingsummaryinternalResponseCompound, aObjEzmaxinvoicingagent []EzmaxinvoicingagentResponseCompound, aObjEzmaxinvoicinguser []EzmaxinvoicinguserResponseCompound, aObjEzmaxinvoicingezsignfolder []CustomEzmaxinvoicingEzsignfolderResponse, aObjEzmaxinvoicingezsigndocument []CustomEzmaxinvoicingEzsigndocumentResponse, ) *EzmaxinvoicingGetProvisionalV1ResponseMPayload`

NewEzmaxinvoicingGetProvisionalV1ResponseMPayload instantiates a new EzmaxinvoicingGetProvisionalV1ResponseMPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzmaxinvoicingGetProvisionalV1ResponseMPayloadWithDefaults

`func NewEzmaxinvoicingGetProvisionalV1ResponseMPayloadWithDefaults() *EzmaxinvoicingGetProvisionalV1ResponseMPayload`

NewEzmaxinvoicingGetProvisionalV1ResponseMPayloadWithDefaults instantiates a new EzmaxinvoicingGetProvisionalV1ResponseMPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzmaxinvoicingID

`func (o *EzmaxinvoicingGetProvisionalV1ResponseMPayload) GetPkiEzmaxinvoicingID() int32`

GetPkiEzmaxinvoicingID returns the PkiEzmaxinvoicingID field if non-nil, zero value otherwise.

### GetPkiEzmaxinvoicingIDOk

`func (o *EzmaxinvoicingGetProvisionalV1ResponseMPayload) GetPkiEzmaxinvoicingIDOk() (*int32, bool)`

GetPkiEzmaxinvoicingIDOk returns a tuple with the PkiEzmaxinvoicingID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzmaxinvoicingID

`func (o *EzmaxinvoicingGetProvisionalV1ResponseMPayload) SetPkiEzmaxinvoicingID(v int32)`

SetPkiEzmaxinvoicingID sets PkiEzmaxinvoicingID field to given value.

### HasPkiEzmaxinvoicingID

`func (o *EzmaxinvoicingGetProvisionalV1ResponseMPayload) HasPkiEzmaxinvoicingID() bool`

HasPkiEzmaxinvoicingID returns a boolean if a field has been set.

### GetFkiEzmaxinvoicingcontractID

`func (o *EzmaxinvoicingGetProvisionalV1ResponseMPayload) GetFkiEzmaxinvoicingcontractID() int32`

GetFkiEzmaxinvoicingcontractID returns the FkiEzmaxinvoicingcontractID field if non-nil, zero value otherwise.

### GetFkiEzmaxinvoicingcontractIDOk

`func (o *EzmaxinvoicingGetProvisionalV1ResponseMPayload) GetFkiEzmaxinvoicingcontractIDOk() (*int32, bool)`

GetFkiEzmaxinvoicingcontractIDOk returns a tuple with the FkiEzmaxinvoicingcontractID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzmaxinvoicingcontractID

`func (o *EzmaxinvoicingGetProvisionalV1ResponseMPayload) SetFkiEzmaxinvoicingcontractID(v int32)`

SetFkiEzmaxinvoicingcontractID sets FkiEzmaxinvoicingcontractID field to given value.


### GetFkiEzmaxpricingID

`func (o *EzmaxinvoicingGetProvisionalV1ResponseMPayload) GetFkiEzmaxpricingID() int32`

GetFkiEzmaxpricingID returns the FkiEzmaxpricingID field if non-nil, zero value otherwise.

### GetFkiEzmaxpricingIDOk

`func (o *EzmaxinvoicingGetProvisionalV1ResponseMPayload) GetFkiEzmaxpricingIDOk() (*int32, bool)`

GetFkiEzmaxpricingIDOk returns a tuple with the FkiEzmaxpricingID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzmaxpricingID

`func (o *EzmaxinvoicingGetProvisionalV1ResponseMPayload) SetFkiEzmaxpricingID(v int32)`

SetFkiEzmaxpricingID sets FkiEzmaxpricingID field to given value.


### GetFkiSystemconfigurationtypeID

`func (o *EzmaxinvoicingGetProvisionalV1ResponseMPayload) GetFkiSystemconfigurationtypeID() int32`

GetFkiSystemconfigurationtypeID returns the FkiSystemconfigurationtypeID field if non-nil, zero value otherwise.

### GetFkiSystemconfigurationtypeIDOk

`func (o *EzmaxinvoicingGetProvisionalV1ResponseMPayload) GetFkiSystemconfigurationtypeIDOk() (*int32, bool)`

GetFkiSystemconfigurationtypeIDOk returns a tuple with the FkiSystemconfigurationtypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiSystemconfigurationtypeID

`func (o *EzmaxinvoicingGetProvisionalV1ResponseMPayload) SetFkiSystemconfigurationtypeID(v int32)`

SetFkiSystemconfigurationtypeID sets FkiSystemconfigurationtypeID field to given value.


### GetSSystemconfigurationtypeDescriptionX

`func (o *EzmaxinvoicingGetProvisionalV1ResponseMPayload) GetSSystemconfigurationtypeDescriptionX() string`

GetSSystemconfigurationtypeDescriptionX returns the SSystemconfigurationtypeDescriptionX field if non-nil, zero value otherwise.

### GetSSystemconfigurationtypeDescriptionXOk

`func (o *EzmaxinvoicingGetProvisionalV1ResponseMPayload) GetSSystemconfigurationtypeDescriptionXOk() (*string, bool)`

GetSSystemconfigurationtypeDescriptionXOk returns a tuple with the SSystemconfigurationtypeDescriptionX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSSystemconfigurationtypeDescriptionX

`func (o *EzmaxinvoicingGetProvisionalV1ResponseMPayload) SetSSystemconfigurationtypeDescriptionX(v string)`

SetSSystemconfigurationtypeDescriptionX sets SSystemconfigurationtypeDescriptionX field to given value.


### GetYyyymmEzmaxinvoicing

`func (o *EzmaxinvoicingGetProvisionalV1ResponseMPayload) GetYyyymmEzmaxinvoicing() string`

GetYyyymmEzmaxinvoicing returns the YyyymmEzmaxinvoicing field if non-nil, zero value otherwise.

### GetYyyymmEzmaxinvoicingOk

`func (o *EzmaxinvoicingGetProvisionalV1ResponseMPayload) GetYyyymmEzmaxinvoicingOk() (*string, bool)`

GetYyyymmEzmaxinvoicingOk returns a tuple with the YyyymmEzmaxinvoicing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYyyymmEzmaxinvoicing

`func (o *EzmaxinvoicingGetProvisionalV1ResponseMPayload) SetYyyymmEzmaxinvoicing(v string)`

SetYyyymmEzmaxinvoicing sets YyyymmEzmaxinvoicing field to given value.


### GetIEzmaxinvoicingDays

`func (o *EzmaxinvoicingGetProvisionalV1ResponseMPayload) GetIEzmaxinvoicingDays() int32`

GetIEzmaxinvoicingDays returns the IEzmaxinvoicingDays field if non-nil, zero value otherwise.

### GetIEzmaxinvoicingDaysOk

`func (o *EzmaxinvoicingGetProvisionalV1ResponseMPayload) GetIEzmaxinvoicingDaysOk() (*int32, bool)`

GetIEzmaxinvoicingDaysOk returns a tuple with the IEzmaxinvoicingDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzmaxinvoicingDays

`func (o *EzmaxinvoicingGetProvisionalV1ResponseMPayload) SetIEzmaxinvoicingDays(v int32)`

SetIEzmaxinvoicingDays sets IEzmaxinvoicingDays field to given value.


### GetEEzmaxinvoicingPaymenttype

`func (o *EzmaxinvoicingGetProvisionalV1ResponseMPayload) GetEEzmaxinvoicingPaymenttype() FieldEEzmaxinvoicingPaymenttype`

GetEEzmaxinvoicingPaymenttype returns the EEzmaxinvoicingPaymenttype field if non-nil, zero value otherwise.

### GetEEzmaxinvoicingPaymenttypeOk

`func (o *EzmaxinvoicingGetProvisionalV1ResponseMPayload) GetEEzmaxinvoicingPaymenttypeOk() (*FieldEEzmaxinvoicingPaymenttype, bool)`

GetEEzmaxinvoicingPaymenttypeOk returns a tuple with the EEzmaxinvoicingPaymenttype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzmaxinvoicingPaymenttype

`func (o *EzmaxinvoicingGetProvisionalV1ResponseMPayload) SetEEzmaxinvoicingPaymenttype(v FieldEEzmaxinvoicingPaymenttype)`

SetEEzmaxinvoicingPaymenttype sets EEzmaxinvoicingPaymenttype field to given value.


### GetDEzmaxinvoicingRebatepaymenttype

`func (o *EzmaxinvoicingGetProvisionalV1ResponseMPayload) GetDEzmaxinvoicingRebatepaymenttype() string`

GetDEzmaxinvoicingRebatepaymenttype returns the DEzmaxinvoicingRebatepaymenttype field if non-nil, zero value otherwise.

### GetDEzmaxinvoicingRebatepaymenttypeOk

`func (o *EzmaxinvoicingGetProvisionalV1ResponseMPayload) GetDEzmaxinvoicingRebatepaymenttypeOk() (*string, bool)`

GetDEzmaxinvoicingRebatepaymenttypeOk returns a tuple with the DEzmaxinvoicingRebatepaymenttype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDEzmaxinvoicingRebatepaymenttype

`func (o *EzmaxinvoicingGetProvisionalV1ResponseMPayload) SetDEzmaxinvoicingRebatepaymenttype(v string)`

SetDEzmaxinvoicingRebatepaymenttype sets DEzmaxinvoicingRebatepaymenttype field to given value.


### GetIEzmaxinvoicingContractlength

`func (o *EzmaxinvoicingGetProvisionalV1ResponseMPayload) GetIEzmaxinvoicingContractlength() int32`

GetIEzmaxinvoicingContractlength returns the IEzmaxinvoicingContractlength field if non-nil, zero value otherwise.

### GetIEzmaxinvoicingContractlengthOk

`func (o *EzmaxinvoicingGetProvisionalV1ResponseMPayload) GetIEzmaxinvoicingContractlengthOk() (*int32, bool)`

GetIEzmaxinvoicingContractlengthOk returns a tuple with the IEzmaxinvoicingContractlength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzmaxinvoicingContractlength

`func (o *EzmaxinvoicingGetProvisionalV1ResponseMPayload) SetIEzmaxinvoicingContractlength(v int32)`

SetIEzmaxinvoicingContractlength sets IEzmaxinvoicingContractlength field to given value.


### GetDEzmaxinvoicingRebatecontractlength

`func (o *EzmaxinvoicingGetProvisionalV1ResponseMPayload) GetDEzmaxinvoicingRebatecontractlength() string`

GetDEzmaxinvoicingRebatecontractlength returns the DEzmaxinvoicingRebatecontractlength field if non-nil, zero value otherwise.

### GetDEzmaxinvoicingRebatecontractlengthOk

`func (o *EzmaxinvoicingGetProvisionalV1ResponseMPayload) GetDEzmaxinvoicingRebatecontractlengthOk() (*string, bool)`

GetDEzmaxinvoicingRebatecontractlengthOk returns a tuple with the DEzmaxinvoicingRebatecontractlength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDEzmaxinvoicingRebatecontractlength

`func (o *EzmaxinvoicingGetProvisionalV1ResponseMPayload) SetDEzmaxinvoicingRebatecontractlength(v string)`

SetDEzmaxinvoicingRebatecontractlength sets DEzmaxinvoicingRebatecontractlength field to given value.


### GetBEzmaxinvoicingRebateEzsignallagents

`func (o *EzmaxinvoicingGetProvisionalV1ResponseMPayload) GetBEzmaxinvoicingRebateEzsignallagents() bool`

GetBEzmaxinvoicingRebateEzsignallagents returns the BEzmaxinvoicingRebateEzsignallagents field if non-nil, zero value otherwise.

### GetBEzmaxinvoicingRebateEzsignallagentsOk

`func (o *EzmaxinvoicingGetProvisionalV1ResponseMPayload) GetBEzmaxinvoicingRebateEzsignallagentsOk() (*bool, bool)`

GetBEzmaxinvoicingRebateEzsignallagentsOk returns a tuple with the BEzmaxinvoicingRebateEzsignallagents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzmaxinvoicingRebateEzsignallagents

`func (o *EzmaxinvoicingGetProvisionalV1ResponseMPayload) SetBEzmaxinvoicingRebateEzsignallagents(v bool)`

SetBEzmaxinvoicingRebateEzsignallagents sets BEzmaxinvoicingRebateEzsignallagents field to given value.


### GetObjAudit

`func (o *EzmaxinvoicingGetProvisionalV1ResponseMPayload) GetObjAudit() CommonAudit`

GetObjAudit returns the ObjAudit field if non-nil, zero value otherwise.

### GetObjAuditOk

`func (o *EzmaxinvoicingGetProvisionalV1ResponseMPayload) GetObjAuditOk() (*CommonAudit, bool)`

GetObjAuditOk returns a tuple with the ObjAudit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjAudit

`func (o *EzmaxinvoicingGetProvisionalV1ResponseMPayload) SetObjAudit(v CommonAudit)`

SetObjAudit sets ObjAudit field to given value.

### HasObjAudit

`func (o *EzmaxinvoicingGetProvisionalV1ResponseMPayload) HasObjAudit() bool`

HasObjAudit returns a boolean if a field has been set.

### GetObjEzmaxinvoicingcontract

`func (o *EzmaxinvoicingGetProvisionalV1ResponseMPayload) GetObjEzmaxinvoicingcontract() EzmaxinvoicingcontractResponseCompound`

GetObjEzmaxinvoicingcontract returns the ObjEzmaxinvoicingcontract field if non-nil, zero value otherwise.

### GetObjEzmaxinvoicingcontractOk

`func (o *EzmaxinvoicingGetProvisionalV1ResponseMPayload) GetObjEzmaxinvoicingcontractOk() (*EzmaxinvoicingcontractResponseCompound, bool)`

GetObjEzmaxinvoicingcontractOk returns a tuple with the ObjEzmaxinvoicingcontract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjEzmaxinvoicingcontract

`func (o *EzmaxinvoicingGetProvisionalV1ResponseMPayload) SetObjEzmaxinvoicingcontract(v EzmaxinvoicingcontractResponseCompound)`

SetObjEzmaxinvoicingcontract sets ObjEzmaxinvoicingcontract field to given value.


### GetObjEzmaxpricing

`func (o *EzmaxinvoicingGetProvisionalV1ResponseMPayload) GetObjEzmaxpricing() CustomEzmaxpricingResponse`

GetObjEzmaxpricing returns the ObjEzmaxpricing field if non-nil, zero value otherwise.

### GetObjEzmaxpricingOk

`func (o *EzmaxinvoicingGetProvisionalV1ResponseMPayload) GetObjEzmaxpricingOk() (*CustomEzmaxpricingResponse, bool)`

GetObjEzmaxpricingOk returns a tuple with the ObjEzmaxpricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjEzmaxpricing

`func (o *EzmaxinvoicingGetProvisionalV1ResponseMPayload) SetObjEzmaxpricing(v CustomEzmaxpricingResponse)`

SetObjEzmaxpricing sets ObjEzmaxpricing field to given value.


### GetAObjEzmaxinvoicingsummaryglobal

`func (o *EzmaxinvoicingGetProvisionalV1ResponseMPayload) GetAObjEzmaxinvoicingsummaryglobal() []EzmaxinvoicingsummaryglobalResponseCompound`

GetAObjEzmaxinvoicingsummaryglobal returns the AObjEzmaxinvoicingsummaryglobal field if non-nil, zero value otherwise.

### GetAObjEzmaxinvoicingsummaryglobalOk

`func (o *EzmaxinvoicingGetProvisionalV1ResponseMPayload) GetAObjEzmaxinvoicingsummaryglobalOk() (*[]EzmaxinvoicingsummaryglobalResponseCompound, bool)`

GetAObjEzmaxinvoicingsummaryglobalOk returns a tuple with the AObjEzmaxinvoicingsummaryglobal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEzmaxinvoicingsummaryglobal

`func (o *EzmaxinvoicingGetProvisionalV1ResponseMPayload) SetAObjEzmaxinvoicingsummaryglobal(v []EzmaxinvoicingsummaryglobalResponseCompound)`

SetAObjEzmaxinvoicingsummaryglobal sets AObjEzmaxinvoicingsummaryglobal field to given value.


### GetAObjEzmaxinvoicingsummaryexternal

`func (o *EzmaxinvoicingGetProvisionalV1ResponseMPayload) GetAObjEzmaxinvoicingsummaryexternal() []EzmaxinvoicingsummaryexternalResponseCompound`

GetAObjEzmaxinvoicingsummaryexternal returns the AObjEzmaxinvoicingsummaryexternal field if non-nil, zero value otherwise.

### GetAObjEzmaxinvoicingsummaryexternalOk

`func (o *EzmaxinvoicingGetProvisionalV1ResponseMPayload) GetAObjEzmaxinvoicingsummaryexternalOk() (*[]EzmaxinvoicingsummaryexternalResponseCompound, bool)`

GetAObjEzmaxinvoicingsummaryexternalOk returns a tuple with the AObjEzmaxinvoicingsummaryexternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEzmaxinvoicingsummaryexternal

`func (o *EzmaxinvoicingGetProvisionalV1ResponseMPayload) SetAObjEzmaxinvoicingsummaryexternal(v []EzmaxinvoicingsummaryexternalResponseCompound)`

SetAObjEzmaxinvoicingsummaryexternal sets AObjEzmaxinvoicingsummaryexternal field to given value.


### GetAObjEzmaxinvoicingsummaryinternal

`func (o *EzmaxinvoicingGetProvisionalV1ResponseMPayload) GetAObjEzmaxinvoicingsummaryinternal() []EzmaxinvoicingsummaryinternalResponseCompound`

GetAObjEzmaxinvoicingsummaryinternal returns the AObjEzmaxinvoicingsummaryinternal field if non-nil, zero value otherwise.

### GetAObjEzmaxinvoicingsummaryinternalOk

`func (o *EzmaxinvoicingGetProvisionalV1ResponseMPayload) GetAObjEzmaxinvoicingsummaryinternalOk() (*[]EzmaxinvoicingsummaryinternalResponseCompound, bool)`

GetAObjEzmaxinvoicingsummaryinternalOk returns a tuple with the AObjEzmaxinvoicingsummaryinternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEzmaxinvoicingsummaryinternal

`func (o *EzmaxinvoicingGetProvisionalV1ResponseMPayload) SetAObjEzmaxinvoicingsummaryinternal(v []EzmaxinvoicingsummaryinternalResponseCompound)`

SetAObjEzmaxinvoicingsummaryinternal sets AObjEzmaxinvoicingsummaryinternal field to given value.


### GetAObjEzmaxinvoicingagent

`func (o *EzmaxinvoicingGetProvisionalV1ResponseMPayload) GetAObjEzmaxinvoicingagent() []EzmaxinvoicingagentResponseCompound`

GetAObjEzmaxinvoicingagent returns the AObjEzmaxinvoicingagent field if non-nil, zero value otherwise.

### GetAObjEzmaxinvoicingagentOk

`func (o *EzmaxinvoicingGetProvisionalV1ResponseMPayload) GetAObjEzmaxinvoicingagentOk() (*[]EzmaxinvoicingagentResponseCompound, bool)`

GetAObjEzmaxinvoicingagentOk returns a tuple with the AObjEzmaxinvoicingagent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEzmaxinvoicingagent

`func (o *EzmaxinvoicingGetProvisionalV1ResponseMPayload) SetAObjEzmaxinvoicingagent(v []EzmaxinvoicingagentResponseCompound)`

SetAObjEzmaxinvoicingagent sets AObjEzmaxinvoicingagent field to given value.


### GetAObjEzmaxinvoicinguser

`func (o *EzmaxinvoicingGetProvisionalV1ResponseMPayload) GetAObjEzmaxinvoicinguser() []EzmaxinvoicinguserResponseCompound`

GetAObjEzmaxinvoicinguser returns the AObjEzmaxinvoicinguser field if non-nil, zero value otherwise.

### GetAObjEzmaxinvoicinguserOk

`func (o *EzmaxinvoicingGetProvisionalV1ResponseMPayload) GetAObjEzmaxinvoicinguserOk() (*[]EzmaxinvoicinguserResponseCompound, bool)`

GetAObjEzmaxinvoicinguserOk returns a tuple with the AObjEzmaxinvoicinguser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEzmaxinvoicinguser

`func (o *EzmaxinvoicingGetProvisionalV1ResponseMPayload) SetAObjEzmaxinvoicinguser(v []EzmaxinvoicinguserResponseCompound)`

SetAObjEzmaxinvoicinguser sets AObjEzmaxinvoicinguser field to given value.


### GetAObjEzmaxinvoicingezsignfolder

`func (o *EzmaxinvoicingGetProvisionalV1ResponseMPayload) GetAObjEzmaxinvoicingezsignfolder() []CustomEzmaxinvoicingEzsignfolderResponse`

GetAObjEzmaxinvoicingezsignfolder returns the AObjEzmaxinvoicingezsignfolder field if non-nil, zero value otherwise.

### GetAObjEzmaxinvoicingezsignfolderOk

`func (o *EzmaxinvoicingGetProvisionalV1ResponseMPayload) GetAObjEzmaxinvoicingezsignfolderOk() (*[]CustomEzmaxinvoicingEzsignfolderResponse, bool)`

GetAObjEzmaxinvoicingezsignfolderOk returns a tuple with the AObjEzmaxinvoicingezsignfolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEzmaxinvoicingezsignfolder

`func (o *EzmaxinvoicingGetProvisionalV1ResponseMPayload) SetAObjEzmaxinvoicingezsignfolder(v []CustomEzmaxinvoicingEzsignfolderResponse)`

SetAObjEzmaxinvoicingezsignfolder sets AObjEzmaxinvoicingezsignfolder field to given value.


### GetAObjEzmaxinvoicingezsigndocument

`func (o *EzmaxinvoicingGetProvisionalV1ResponseMPayload) GetAObjEzmaxinvoicingezsigndocument() []CustomEzmaxinvoicingEzsigndocumentResponse`

GetAObjEzmaxinvoicingezsigndocument returns the AObjEzmaxinvoicingezsigndocument field if non-nil, zero value otherwise.

### GetAObjEzmaxinvoicingezsigndocumentOk

`func (o *EzmaxinvoicingGetProvisionalV1ResponseMPayload) GetAObjEzmaxinvoicingezsigndocumentOk() (*[]CustomEzmaxinvoicingEzsigndocumentResponse, bool)`

GetAObjEzmaxinvoicingezsigndocumentOk returns a tuple with the AObjEzmaxinvoicingezsigndocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEzmaxinvoicingezsigndocument

`func (o *EzmaxinvoicingGetProvisionalV1ResponseMPayload) SetAObjEzmaxinvoicingezsigndocument(v []CustomEzmaxinvoicingEzsigndocumentResponse)`

SetAObjEzmaxinvoicingezsigndocument sets AObjEzmaxinvoicingezsigndocument field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


