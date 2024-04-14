/*
eZmax API Definition (Full)

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.2.1
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the EzmaxinvoicingResponseCompound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzmaxinvoicingResponseCompound{}

// EzmaxinvoicingResponseCompound A Ezmaxinvoicing Object
type EzmaxinvoicingResponseCompound struct {
	// The unique ID of the Ezmaxinvoicing
	PkiEzmaxinvoicingID *int32 `json:"pkiEzmaxinvoicingID,omitempty"`
	// The unique ID of the Ezmaxinvoicingcontract
	FkiEzmaxinvoicingcontractID int32 `json:"fkiEzmaxinvoicingcontractID"`
	// The unique ID of the Ezmaxpricing
	FkiEzmaxpricingID int32 `json:"fkiEzmaxpricingID"`
	// The unique ID of the Systemconfigurationtype
	FkiSystemconfigurationtypeID int32 `json:"fkiSystemconfigurationtypeID"`
	// The description of the Systemconfigurationtype in the language of the requester
	SSystemconfigurationtypeDescriptionX string `json:"sSystemconfigurationtypeDescriptionX"`
	// The YYYYMM period of the Ezmaxinvoicing
	YyyymmEzmaxinvoicing string `json:"yyyymmEzmaxinvoicing"`
	// The number of days invoiced
	IEzmaxinvoicingDays int32 `json:"iEzmaxinvoicingDays"`
	EEzmaxinvoicingPaymenttype FieldEEzmaxinvoicingPaymenttype `json:"eEzmaxinvoicingPaymenttype"`
	// The percentage of rebate depending of the payment type
	DEzmaxinvoicingRebatepaymenttype string `json:"dEzmaxinvoicingRebatepaymenttype"`
	// The length of the contract in years
	IEzmaxinvoicingContractlength int32 `json:"iEzmaxinvoicingContractlength"`
	// The percentage of rebate depending of the contract length
	DEzmaxinvoicingRebatecontractlength string `json:"dEzmaxinvoicingRebatecontractlength"`
	// Whether the rebate for eZsign is for all agents
	BEzmaxinvoicingRebateEzsignallagents bool `json:"bEzmaxinvoicingRebateEzsignallagents"`
	ObjAudit *CommonAudit `json:"objAudit,omitempty"`
	ObjEzmaxinvoicingcontract EzmaxinvoicingcontractResponseCompound `json:"objEzmaxinvoicingcontract"`
	ObjEzmaxpricing CustomEzmaxpricingResponse `json:"objEzmaxpricing"`
	AObjEzmaxinvoicingsummaryglobal []EzmaxinvoicingsummaryglobalResponseCompound `json:"a_objEzmaxinvoicingsummaryglobal"`
	AObjEzmaxinvoicingsummaryexternal []EzmaxinvoicingsummaryexternalResponseCompound `json:"a_objEzmaxinvoicingsummaryexternal"`
	AObjEzmaxinvoicingsummaryinternal []EzmaxinvoicingsummaryinternalResponseCompound `json:"a_objEzmaxinvoicingsummaryinternal"`
	AObjEzmaxinvoicingagent []EzmaxinvoicingagentResponseCompound `json:"a_objEzmaxinvoicingagent"`
	AObjEzmaxinvoicinguser []EzmaxinvoicinguserResponseCompound `json:"a_objEzmaxinvoicinguser"`
	AObjEzmaxinvoicingezsignfolder []CustomEzmaxinvoicingEzsignfolderResponse `json:"a_objEzmaxinvoicingezsignfolder"`
	AObjEzmaxinvoicingezsigndocument []CustomEzmaxinvoicingEzsigndocumentResponse `json:"a_objEzmaxinvoicingezsigndocument"`
}

type _EzmaxinvoicingResponseCompound EzmaxinvoicingResponseCompound

// NewEzmaxinvoicingResponseCompound instantiates a new EzmaxinvoicingResponseCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzmaxinvoicingResponseCompound(fkiEzmaxinvoicingcontractID int32, fkiEzmaxpricingID int32, fkiSystemconfigurationtypeID int32, sSystemconfigurationtypeDescriptionX string, yyyymmEzmaxinvoicing string, iEzmaxinvoicingDays int32, eEzmaxinvoicingPaymenttype FieldEEzmaxinvoicingPaymenttype, dEzmaxinvoicingRebatepaymenttype string, iEzmaxinvoicingContractlength int32, dEzmaxinvoicingRebatecontractlength string, bEzmaxinvoicingRebateEzsignallagents bool, objEzmaxinvoicingcontract EzmaxinvoicingcontractResponseCompound, objEzmaxpricing CustomEzmaxpricingResponse, aObjEzmaxinvoicingsummaryglobal []EzmaxinvoicingsummaryglobalResponseCompound, aObjEzmaxinvoicingsummaryexternal []EzmaxinvoicingsummaryexternalResponseCompound, aObjEzmaxinvoicingsummaryinternal []EzmaxinvoicingsummaryinternalResponseCompound, aObjEzmaxinvoicingagent []EzmaxinvoicingagentResponseCompound, aObjEzmaxinvoicinguser []EzmaxinvoicinguserResponseCompound, aObjEzmaxinvoicingezsignfolder []CustomEzmaxinvoicingEzsignfolderResponse, aObjEzmaxinvoicingezsigndocument []CustomEzmaxinvoicingEzsigndocumentResponse) *EzmaxinvoicingResponseCompound {
	this := EzmaxinvoicingResponseCompound{}
	this.FkiEzmaxinvoicingcontractID = fkiEzmaxinvoicingcontractID
	this.FkiEzmaxpricingID = fkiEzmaxpricingID
	this.FkiSystemconfigurationtypeID = fkiSystemconfigurationtypeID
	this.SSystemconfigurationtypeDescriptionX = sSystemconfigurationtypeDescriptionX
	this.YyyymmEzmaxinvoicing = yyyymmEzmaxinvoicing
	this.IEzmaxinvoicingDays = iEzmaxinvoicingDays
	this.EEzmaxinvoicingPaymenttype = eEzmaxinvoicingPaymenttype
	this.DEzmaxinvoicingRebatepaymenttype = dEzmaxinvoicingRebatepaymenttype
	this.IEzmaxinvoicingContractlength = iEzmaxinvoicingContractlength
	this.DEzmaxinvoicingRebatecontractlength = dEzmaxinvoicingRebatecontractlength
	this.BEzmaxinvoicingRebateEzsignallagents = bEzmaxinvoicingRebateEzsignallagents
	this.ObjEzmaxinvoicingcontract = objEzmaxinvoicingcontract
	this.ObjEzmaxpricing = objEzmaxpricing
	this.AObjEzmaxinvoicingsummaryglobal = aObjEzmaxinvoicingsummaryglobal
	this.AObjEzmaxinvoicingsummaryexternal = aObjEzmaxinvoicingsummaryexternal
	this.AObjEzmaxinvoicingsummaryinternal = aObjEzmaxinvoicingsummaryinternal
	this.AObjEzmaxinvoicingagent = aObjEzmaxinvoicingagent
	this.AObjEzmaxinvoicinguser = aObjEzmaxinvoicinguser
	this.AObjEzmaxinvoicingezsignfolder = aObjEzmaxinvoicingezsignfolder
	this.AObjEzmaxinvoicingezsigndocument = aObjEzmaxinvoicingezsigndocument
	return &this
}

// NewEzmaxinvoicingResponseCompoundWithDefaults instantiates a new EzmaxinvoicingResponseCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzmaxinvoicingResponseCompoundWithDefaults() *EzmaxinvoicingResponseCompound {
	this := EzmaxinvoicingResponseCompound{}
	return &this
}

// GetPkiEzmaxinvoicingID returns the PkiEzmaxinvoicingID field value if set, zero value otherwise.
func (o *EzmaxinvoicingResponseCompound) GetPkiEzmaxinvoicingID() int32 {
	if o == nil || IsNil(o.PkiEzmaxinvoicingID) {
		var ret int32
		return ret
	}
	return *o.PkiEzmaxinvoicingID
}

// GetPkiEzmaxinvoicingIDOk returns a tuple with the PkiEzmaxinvoicingID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingResponseCompound) GetPkiEzmaxinvoicingIDOk() (*int32, bool) {
	if o == nil || IsNil(o.PkiEzmaxinvoicingID) {
		return nil, false
	}
	return o.PkiEzmaxinvoicingID, true
}

// HasPkiEzmaxinvoicingID returns a boolean if a field has been set.
func (o *EzmaxinvoicingResponseCompound) HasPkiEzmaxinvoicingID() bool {
	if o != nil && !IsNil(o.PkiEzmaxinvoicingID) {
		return true
	}

	return false
}

// SetPkiEzmaxinvoicingID gets a reference to the given int32 and assigns it to the PkiEzmaxinvoicingID field.
func (o *EzmaxinvoicingResponseCompound) SetPkiEzmaxinvoicingID(v int32) {
	o.PkiEzmaxinvoicingID = &v
}

// GetFkiEzmaxinvoicingcontractID returns the FkiEzmaxinvoicingcontractID field value
func (o *EzmaxinvoicingResponseCompound) GetFkiEzmaxinvoicingcontractID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzmaxinvoicingcontractID
}

// GetFkiEzmaxinvoicingcontractIDOk returns a tuple with the FkiEzmaxinvoicingcontractID field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingResponseCompound) GetFkiEzmaxinvoicingcontractIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiEzmaxinvoicingcontractID, true
}

// SetFkiEzmaxinvoicingcontractID sets field value
func (o *EzmaxinvoicingResponseCompound) SetFkiEzmaxinvoicingcontractID(v int32) {
	o.FkiEzmaxinvoicingcontractID = v
}

// GetFkiEzmaxpricingID returns the FkiEzmaxpricingID field value
func (o *EzmaxinvoicingResponseCompound) GetFkiEzmaxpricingID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzmaxpricingID
}

// GetFkiEzmaxpricingIDOk returns a tuple with the FkiEzmaxpricingID field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingResponseCompound) GetFkiEzmaxpricingIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiEzmaxpricingID, true
}

// SetFkiEzmaxpricingID sets field value
func (o *EzmaxinvoicingResponseCompound) SetFkiEzmaxpricingID(v int32) {
	o.FkiEzmaxpricingID = v
}

// GetFkiSystemconfigurationtypeID returns the FkiSystemconfigurationtypeID field value
func (o *EzmaxinvoicingResponseCompound) GetFkiSystemconfigurationtypeID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiSystemconfigurationtypeID
}

// GetFkiSystemconfigurationtypeIDOk returns a tuple with the FkiSystemconfigurationtypeID field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingResponseCompound) GetFkiSystemconfigurationtypeIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiSystemconfigurationtypeID, true
}

// SetFkiSystemconfigurationtypeID sets field value
func (o *EzmaxinvoicingResponseCompound) SetFkiSystemconfigurationtypeID(v int32) {
	o.FkiSystemconfigurationtypeID = v
}

// GetSSystemconfigurationtypeDescriptionX returns the SSystemconfigurationtypeDescriptionX field value
func (o *EzmaxinvoicingResponseCompound) GetSSystemconfigurationtypeDescriptionX() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SSystemconfigurationtypeDescriptionX
}

// GetSSystemconfigurationtypeDescriptionXOk returns a tuple with the SSystemconfigurationtypeDescriptionX field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingResponseCompound) GetSSystemconfigurationtypeDescriptionXOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SSystemconfigurationtypeDescriptionX, true
}

// SetSSystemconfigurationtypeDescriptionX sets field value
func (o *EzmaxinvoicingResponseCompound) SetSSystemconfigurationtypeDescriptionX(v string) {
	o.SSystemconfigurationtypeDescriptionX = v
}

// GetYyyymmEzmaxinvoicing returns the YyyymmEzmaxinvoicing field value
func (o *EzmaxinvoicingResponseCompound) GetYyyymmEzmaxinvoicing() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.YyyymmEzmaxinvoicing
}

// GetYyyymmEzmaxinvoicingOk returns a tuple with the YyyymmEzmaxinvoicing field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingResponseCompound) GetYyyymmEzmaxinvoicingOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.YyyymmEzmaxinvoicing, true
}

// SetYyyymmEzmaxinvoicing sets field value
func (o *EzmaxinvoicingResponseCompound) SetYyyymmEzmaxinvoicing(v string) {
	o.YyyymmEzmaxinvoicing = v
}

// GetIEzmaxinvoicingDays returns the IEzmaxinvoicingDays field value
func (o *EzmaxinvoicingResponseCompound) GetIEzmaxinvoicingDays() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzmaxinvoicingDays
}

// GetIEzmaxinvoicingDaysOk returns a tuple with the IEzmaxinvoicingDays field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingResponseCompound) GetIEzmaxinvoicingDaysOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzmaxinvoicingDays, true
}

// SetIEzmaxinvoicingDays sets field value
func (o *EzmaxinvoicingResponseCompound) SetIEzmaxinvoicingDays(v int32) {
	o.IEzmaxinvoicingDays = v
}

// GetEEzmaxinvoicingPaymenttype returns the EEzmaxinvoicingPaymenttype field value
func (o *EzmaxinvoicingResponseCompound) GetEEzmaxinvoicingPaymenttype() FieldEEzmaxinvoicingPaymenttype {
	if o == nil {
		var ret FieldEEzmaxinvoicingPaymenttype
		return ret
	}

	return o.EEzmaxinvoicingPaymenttype
}

// GetEEzmaxinvoicingPaymenttypeOk returns a tuple with the EEzmaxinvoicingPaymenttype field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingResponseCompound) GetEEzmaxinvoicingPaymenttypeOk() (*FieldEEzmaxinvoicingPaymenttype, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EEzmaxinvoicingPaymenttype, true
}

// SetEEzmaxinvoicingPaymenttype sets field value
func (o *EzmaxinvoicingResponseCompound) SetEEzmaxinvoicingPaymenttype(v FieldEEzmaxinvoicingPaymenttype) {
	o.EEzmaxinvoicingPaymenttype = v
}

// GetDEzmaxinvoicingRebatepaymenttype returns the DEzmaxinvoicingRebatepaymenttype field value
func (o *EzmaxinvoicingResponseCompound) GetDEzmaxinvoicingRebatepaymenttype() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DEzmaxinvoicingRebatepaymenttype
}

// GetDEzmaxinvoicingRebatepaymenttypeOk returns a tuple with the DEzmaxinvoicingRebatepaymenttype field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingResponseCompound) GetDEzmaxinvoicingRebatepaymenttypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DEzmaxinvoicingRebatepaymenttype, true
}

// SetDEzmaxinvoicingRebatepaymenttype sets field value
func (o *EzmaxinvoicingResponseCompound) SetDEzmaxinvoicingRebatepaymenttype(v string) {
	o.DEzmaxinvoicingRebatepaymenttype = v
}

// GetIEzmaxinvoicingContractlength returns the IEzmaxinvoicingContractlength field value
func (o *EzmaxinvoicingResponseCompound) GetIEzmaxinvoicingContractlength() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzmaxinvoicingContractlength
}

// GetIEzmaxinvoicingContractlengthOk returns a tuple with the IEzmaxinvoicingContractlength field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingResponseCompound) GetIEzmaxinvoicingContractlengthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzmaxinvoicingContractlength, true
}

// SetIEzmaxinvoicingContractlength sets field value
func (o *EzmaxinvoicingResponseCompound) SetIEzmaxinvoicingContractlength(v int32) {
	o.IEzmaxinvoicingContractlength = v
}

// GetDEzmaxinvoicingRebatecontractlength returns the DEzmaxinvoicingRebatecontractlength field value
func (o *EzmaxinvoicingResponseCompound) GetDEzmaxinvoicingRebatecontractlength() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DEzmaxinvoicingRebatecontractlength
}

// GetDEzmaxinvoicingRebatecontractlengthOk returns a tuple with the DEzmaxinvoicingRebatecontractlength field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingResponseCompound) GetDEzmaxinvoicingRebatecontractlengthOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DEzmaxinvoicingRebatecontractlength, true
}

// SetDEzmaxinvoicingRebatecontractlength sets field value
func (o *EzmaxinvoicingResponseCompound) SetDEzmaxinvoicingRebatecontractlength(v string) {
	o.DEzmaxinvoicingRebatecontractlength = v
}

// GetBEzmaxinvoicingRebateEzsignallagents returns the BEzmaxinvoicingRebateEzsignallagents field value
func (o *EzmaxinvoicingResponseCompound) GetBEzmaxinvoicingRebateEzsignallagents() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BEzmaxinvoicingRebateEzsignallagents
}

// GetBEzmaxinvoicingRebateEzsignallagentsOk returns a tuple with the BEzmaxinvoicingRebateEzsignallagents field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingResponseCompound) GetBEzmaxinvoicingRebateEzsignallagentsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BEzmaxinvoicingRebateEzsignallagents, true
}

// SetBEzmaxinvoicingRebateEzsignallagents sets field value
func (o *EzmaxinvoicingResponseCompound) SetBEzmaxinvoicingRebateEzsignallagents(v bool) {
	o.BEzmaxinvoicingRebateEzsignallagents = v
}

// GetObjAudit returns the ObjAudit field value if set, zero value otherwise.
func (o *EzmaxinvoicingResponseCompound) GetObjAudit() CommonAudit {
	if o == nil || IsNil(o.ObjAudit) {
		var ret CommonAudit
		return ret
	}
	return *o.ObjAudit
}

// GetObjAuditOk returns a tuple with the ObjAudit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingResponseCompound) GetObjAuditOk() (*CommonAudit, bool) {
	if o == nil || IsNil(o.ObjAudit) {
		return nil, false
	}
	return o.ObjAudit, true
}

// HasObjAudit returns a boolean if a field has been set.
func (o *EzmaxinvoicingResponseCompound) HasObjAudit() bool {
	if o != nil && !IsNil(o.ObjAudit) {
		return true
	}

	return false
}

// SetObjAudit gets a reference to the given CommonAudit and assigns it to the ObjAudit field.
func (o *EzmaxinvoicingResponseCompound) SetObjAudit(v CommonAudit) {
	o.ObjAudit = &v
}

// GetObjEzmaxinvoicingcontract returns the ObjEzmaxinvoicingcontract field value
func (o *EzmaxinvoicingResponseCompound) GetObjEzmaxinvoicingcontract() EzmaxinvoicingcontractResponseCompound {
	if o == nil {
		var ret EzmaxinvoicingcontractResponseCompound
		return ret
	}

	return o.ObjEzmaxinvoicingcontract
}

// GetObjEzmaxinvoicingcontractOk returns a tuple with the ObjEzmaxinvoicingcontract field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingResponseCompound) GetObjEzmaxinvoicingcontractOk() (*EzmaxinvoicingcontractResponseCompound, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjEzmaxinvoicingcontract, true
}

// SetObjEzmaxinvoicingcontract sets field value
func (o *EzmaxinvoicingResponseCompound) SetObjEzmaxinvoicingcontract(v EzmaxinvoicingcontractResponseCompound) {
	o.ObjEzmaxinvoicingcontract = v
}

// GetObjEzmaxpricing returns the ObjEzmaxpricing field value
func (o *EzmaxinvoicingResponseCompound) GetObjEzmaxpricing() CustomEzmaxpricingResponse {
	if o == nil {
		var ret CustomEzmaxpricingResponse
		return ret
	}

	return o.ObjEzmaxpricing
}

// GetObjEzmaxpricingOk returns a tuple with the ObjEzmaxpricing field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingResponseCompound) GetObjEzmaxpricingOk() (*CustomEzmaxpricingResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjEzmaxpricing, true
}

// SetObjEzmaxpricing sets field value
func (o *EzmaxinvoicingResponseCompound) SetObjEzmaxpricing(v CustomEzmaxpricingResponse) {
	o.ObjEzmaxpricing = v
}

// GetAObjEzmaxinvoicingsummaryglobal returns the AObjEzmaxinvoicingsummaryglobal field value
func (o *EzmaxinvoicingResponseCompound) GetAObjEzmaxinvoicingsummaryglobal() []EzmaxinvoicingsummaryglobalResponseCompound {
	if o == nil {
		var ret []EzmaxinvoicingsummaryglobalResponseCompound
		return ret
	}

	return o.AObjEzmaxinvoicingsummaryglobal
}

// GetAObjEzmaxinvoicingsummaryglobalOk returns a tuple with the AObjEzmaxinvoicingsummaryglobal field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingResponseCompound) GetAObjEzmaxinvoicingsummaryglobalOk() ([]EzmaxinvoicingsummaryglobalResponseCompound, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjEzmaxinvoicingsummaryglobal, true
}

// SetAObjEzmaxinvoicingsummaryglobal sets field value
func (o *EzmaxinvoicingResponseCompound) SetAObjEzmaxinvoicingsummaryglobal(v []EzmaxinvoicingsummaryglobalResponseCompound) {
	o.AObjEzmaxinvoicingsummaryglobal = v
}

// GetAObjEzmaxinvoicingsummaryexternal returns the AObjEzmaxinvoicingsummaryexternal field value
func (o *EzmaxinvoicingResponseCompound) GetAObjEzmaxinvoicingsummaryexternal() []EzmaxinvoicingsummaryexternalResponseCompound {
	if o == nil {
		var ret []EzmaxinvoicingsummaryexternalResponseCompound
		return ret
	}

	return o.AObjEzmaxinvoicingsummaryexternal
}

// GetAObjEzmaxinvoicingsummaryexternalOk returns a tuple with the AObjEzmaxinvoicingsummaryexternal field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingResponseCompound) GetAObjEzmaxinvoicingsummaryexternalOk() ([]EzmaxinvoicingsummaryexternalResponseCompound, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjEzmaxinvoicingsummaryexternal, true
}

// SetAObjEzmaxinvoicingsummaryexternal sets field value
func (o *EzmaxinvoicingResponseCompound) SetAObjEzmaxinvoicingsummaryexternal(v []EzmaxinvoicingsummaryexternalResponseCompound) {
	o.AObjEzmaxinvoicingsummaryexternal = v
}

// GetAObjEzmaxinvoicingsummaryinternal returns the AObjEzmaxinvoicingsummaryinternal field value
func (o *EzmaxinvoicingResponseCompound) GetAObjEzmaxinvoicingsummaryinternal() []EzmaxinvoicingsummaryinternalResponseCompound {
	if o == nil {
		var ret []EzmaxinvoicingsummaryinternalResponseCompound
		return ret
	}

	return o.AObjEzmaxinvoicingsummaryinternal
}

// GetAObjEzmaxinvoicingsummaryinternalOk returns a tuple with the AObjEzmaxinvoicingsummaryinternal field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingResponseCompound) GetAObjEzmaxinvoicingsummaryinternalOk() ([]EzmaxinvoicingsummaryinternalResponseCompound, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjEzmaxinvoicingsummaryinternal, true
}

// SetAObjEzmaxinvoicingsummaryinternal sets field value
func (o *EzmaxinvoicingResponseCompound) SetAObjEzmaxinvoicingsummaryinternal(v []EzmaxinvoicingsummaryinternalResponseCompound) {
	o.AObjEzmaxinvoicingsummaryinternal = v
}

// GetAObjEzmaxinvoicingagent returns the AObjEzmaxinvoicingagent field value
func (o *EzmaxinvoicingResponseCompound) GetAObjEzmaxinvoicingagent() []EzmaxinvoicingagentResponseCompound {
	if o == nil {
		var ret []EzmaxinvoicingagentResponseCompound
		return ret
	}

	return o.AObjEzmaxinvoicingagent
}

// GetAObjEzmaxinvoicingagentOk returns a tuple with the AObjEzmaxinvoicingagent field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingResponseCompound) GetAObjEzmaxinvoicingagentOk() ([]EzmaxinvoicingagentResponseCompound, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjEzmaxinvoicingagent, true
}

// SetAObjEzmaxinvoicingagent sets field value
func (o *EzmaxinvoicingResponseCompound) SetAObjEzmaxinvoicingagent(v []EzmaxinvoicingagentResponseCompound) {
	o.AObjEzmaxinvoicingagent = v
}

// GetAObjEzmaxinvoicinguser returns the AObjEzmaxinvoicinguser field value
func (o *EzmaxinvoicingResponseCompound) GetAObjEzmaxinvoicinguser() []EzmaxinvoicinguserResponseCompound {
	if o == nil {
		var ret []EzmaxinvoicinguserResponseCompound
		return ret
	}

	return o.AObjEzmaxinvoicinguser
}

// GetAObjEzmaxinvoicinguserOk returns a tuple with the AObjEzmaxinvoicinguser field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingResponseCompound) GetAObjEzmaxinvoicinguserOk() ([]EzmaxinvoicinguserResponseCompound, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjEzmaxinvoicinguser, true
}

// SetAObjEzmaxinvoicinguser sets field value
func (o *EzmaxinvoicingResponseCompound) SetAObjEzmaxinvoicinguser(v []EzmaxinvoicinguserResponseCompound) {
	o.AObjEzmaxinvoicinguser = v
}

// GetAObjEzmaxinvoicingezsignfolder returns the AObjEzmaxinvoicingezsignfolder field value
func (o *EzmaxinvoicingResponseCompound) GetAObjEzmaxinvoicingezsignfolder() []CustomEzmaxinvoicingEzsignfolderResponse {
	if o == nil {
		var ret []CustomEzmaxinvoicingEzsignfolderResponse
		return ret
	}

	return o.AObjEzmaxinvoicingezsignfolder
}

// GetAObjEzmaxinvoicingezsignfolderOk returns a tuple with the AObjEzmaxinvoicingezsignfolder field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingResponseCompound) GetAObjEzmaxinvoicingezsignfolderOk() ([]CustomEzmaxinvoicingEzsignfolderResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjEzmaxinvoicingezsignfolder, true
}

// SetAObjEzmaxinvoicingezsignfolder sets field value
func (o *EzmaxinvoicingResponseCompound) SetAObjEzmaxinvoicingezsignfolder(v []CustomEzmaxinvoicingEzsignfolderResponse) {
	o.AObjEzmaxinvoicingezsignfolder = v
}

// GetAObjEzmaxinvoicingezsigndocument returns the AObjEzmaxinvoicingezsigndocument field value
func (o *EzmaxinvoicingResponseCompound) GetAObjEzmaxinvoicingezsigndocument() []CustomEzmaxinvoicingEzsigndocumentResponse {
	if o == nil {
		var ret []CustomEzmaxinvoicingEzsigndocumentResponse
		return ret
	}

	return o.AObjEzmaxinvoicingezsigndocument
}

// GetAObjEzmaxinvoicingezsigndocumentOk returns a tuple with the AObjEzmaxinvoicingezsigndocument field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingResponseCompound) GetAObjEzmaxinvoicingezsigndocumentOk() ([]CustomEzmaxinvoicingEzsigndocumentResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjEzmaxinvoicingezsigndocument, true
}

// SetAObjEzmaxinvoicingezsigndocument sets field value
func (o *EzmaxinvoicingResponseCompound) SetAObjEzmaxinvoicingezsigndocument(v []CustomEzmaxinvoicingEzsigndocumentResponse) {
	o.AObjEzmaxinvoicingezsigndocument = v
}

func (o EzmaxinvoicingResponseCompound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzmaxinvoicingResponseCompound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PkiEzmaxinvoicingID) {
		toSerialize["pkiEzmaxinvoicingID"] = o.PkiEzmaxinvoicingID
	}
	toSerialize["fkiEzmaxinvoicingcontractID"] = o.FkiEzmaxinvoicingcontractID
	toSerialize["fkiEzmaxpricingID"] = o.FkiEzmaxpricingID
	toSerialize["fkiSystemconfigurationtypeID"] = o.FkiSystemconfigurationtypeID
	toSerialize["sSystemconfigurationtypeDescriptionX"] = o.SSystemconfigurationtypeDescriptionX
	toSerialize["yyyymmEzmaxinvoicing"] = o.YyyymmEzmaxinvoicing
	toSerialize["iEzmaxinvoicingDays"] = o.IEzmaxinvoicingDays
	toSerialize["eEzmaxinvoicingPaymenttype"] = o.EEzmaxinvoicingPaymenttype
	toSerialize["dEzmaxinvoicingRebatepaymenttype"] = o.DEzmaxinvoicingRebatepaymenttype
	toSerialize["iEzmaxinvoicingContractlength"] = o.IEzmaxinvoicingContractlength
	toSerialize["dEzmaxinvoicingRebatecontractlength"] = o.DEzmaxinvoicingRebatecontractlength
	toSerialize["bEzmaxinvoicingRebateEzsignallagents"] = o.BEzmaxinvoicingRebateEzsignallagents
	if !IsNil(o.ObjAudit) {
		toSerialize["objAudit"] = o.ObjAudit
	}
	toSerialize["objEzmaxinvoicingcontract"] = o.ObjEzmaxinvoicingcontract
	toSerialize["objEzmaxpricing"] = o.ObjEzmaxpricing
	toSerialize["a_objEzmaxinvoicingsummaryglobal"] = o.AObjEzmaxinvoicingsummaryglobal
	toSerialize["a_objEzmaxinvoicingsummaryexternal"] = o.AObjEzmaxinvoicingsummaryexternal
	toSerialize["a_objEzmaxinvoicingsummaryinternal"] = o.AObjEzmaxinvoicingsummaryinternal
	toSerialize["a_objEzmaxinvoicingagent"] = o.AObjEzmaxinvoicingagent
	toSerialize["a_objEzmaxinvoicinguser"] = o.AObjEzmaxinvoicinguser
	toSerialize["a_objEzmaxinvoicingezsignfolder"] = o.AObjEzmaxinvoicingezsignfolder
	toSerialize["a_objEzmaxinvoicingezsigndocument"] = o.AObjEzmaxinvoicingezsigndocument
	return toSerialize, nil
}

func (o *EzmaxinvoicingResponseCompound) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fkiEzmaxinvoicingcontractID",
		"fkiEzmaxpricingID",
		"fkiSystemconfigurationtypeID",
		"sSystemconfigurationtypeDescriptionX",
		"yyyymmEzmaxinvoicing",
		"iEzmaxinvoicingDays",
		"eEzmaxinvoicingPaymenttype",
		"dEzmaxinvoicingRebatepaymenttype",
		"iEzmaxinvoicingContractlength",
		"dEzmaxinvoicingRebatecontractlength",
		"bEzmaxinvoicingRebateEzsignallagents",
		"objEzmaxinvoicingcontract",
		"objEzmaxpricing",
		"a_objEzmaxinvoicingsummaryglobal",
		"a_objEzmaxinvoicingsummaryexternal",
		"a_objEzmaxinvoicingsummaryinternal",
		"a_objEzmaxinvoicingagent",
		"a_objEzmaxinvoicinguser",
		"a_objEzmaxinvoicingezsignfolder",
		"a_objEzmaxinvoicingezsigndocument",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varEzmaxinvoicingResponseCompound := _EzmaxinvoicingResponseCompound{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzmaxinvoicingResponseCompound)

	if err != nil {
		return err
	}

	*o = EzmaxinvoicingResponseCompound(varEzmaxinvoicingResponseCompound)

	return err
}

type NullableEzmaxinvoicingResponseCompound struct {
	value *EzmaxinvoicingResponseCompound
	isSet bool
}

func (v NullableEzmaxinvoicingResponseCompound) Get() *EzmaxinvoicingResponseCompound {
	return v.value
}

func (v *NullableEzmaxinvoicingResponseCompound) Set(val *EzmaxinvoicingResponseCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableEzmaxinvoicingResponseCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableEzmaxinvoicingResponseCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzmaxinvoicingResponseCompound(val *EzmaxinvoicingResponseCompound) *NullableEzmaxinvoicingResponseCompound {
	return &NullableEzmaxinvoicingResponseCompound{value: val, isSet: true}
}

func (v NullableEzmaxinvoicingResponseCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzmaxinvoicingResponseCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


