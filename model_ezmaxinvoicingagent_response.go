/*
eZmax API Definition (Full)

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.2.2
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the EzmaxinvoicingagentResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzmaxinvoicingagentResponse{}

// EzmaxinvoicingagentResponse A Ezmaxinvoicingagent Object
type EzmaxinvoicingagentResponse struct {
	// The unique ID of the Ezmaxinvoicingagent
	PkiEzmaxinvoicingagentID *int32 `json:"pkiEzmaxinvoicingagentID,omitempty"`
	// The unique ID of the Ezmaxinvoicing
	FkiEzmaxinvoicingID *int32 `json:"fkiEzmaxinvoicingID,omitempty"`
	// The unique ID of the Billingentityinternal.
	FkiBillingentityinternalID int32 `json:"fkiBillingentityinternalID"`
	// The description of the Billingentityinternal in the language of the requester
	SBillingentityinternalDescriptionX string `json:"sBillingentityinternalDescriptionX"`
	// The unique ID of the Agent.
	FkiAgentID *int32 `json:"fkiAgentID,omitempty"`
	// The unique ID of the Broker.
	FkiBrokerID *int32 `json:"fkiBrokerID,omitempty"`
	// The number of sessions
	IEzmaxinvoicingagentSession int32 `json:"iEzmaxinvoicingagentSession"`
	// The number of times this user was cloned
	IEzmaxinvoicingagentCloned int32 `json:"iEzmaxinvoicingagentCloned"`
	// The number of invoices
	IEzmaxinvoicingagentInvoice int32 `json:"iEzmaxinvoicingagentInvoice"`
	// The number of inscriptions
	IEzmaxinvoicingagentInscription int32 `json:"iEzmaxinvoicingagentInscription"`
	// The number of active inscriptions
	IEzmaxinvoicingagentInscriptionactive int32 `json:"iEzmaxinvoicingagentInscriptionactive"`
	// The number of sales
	IEzmaxinvoicingagentSale int32 `json:"iEzmaxinvoicingagentSale"`
	// The number of otherincomes
	IEzmaxinvoicingagentOtherincome int32 `json:"iEzmaxinvoicingagentOtherincome"`
	// The number of commission calculations
	IEzmaxinvoicingagentCommissioncalculation int32 `json:"iEzmaxinvoicingagentCommissioncalculation"`
	// The number of ezsign documents
	IEzmaxinvoicingagentEzsigndocument int32 `json:"iEzmaxinvoicingagentEzsigndocument"`
	// Whether the agent has an eZsign account
	BEzmaxinvoicingagentEzsignaccount bool `json:"bEzmaxinvoicingagentEzsignaccount"`
	// Whether it is billable for eZmax
	BEzmaxinvoicingagentBillableezmax bool `json:"bEzmaxinvoicingagentBillableezmax"`
	EEzmaxinvoicingagentVariationezmax FieldEEzmaxinvoicingagentVariationezmax `json:"eEzmaxinvoicingagentVariationezmax"`
	// Whether it is billable for eZsign
	BEzmaxinvoicingagentBillableezsign bool `json:"bEzmaxinvoicingagentBillableezsign"`
	EEzmaxinvoicingagentVariationezsign FieldEEzmaxinvoicingagentVariationezsign `json:"eEzmaxinvoicingagentVariationezsign"`
}

type _EzmaxinvoicingagentResponse EzmaxinvoicingagentResponse

// NewEzmaxinvoicingagentResponse instantiates a new EzmaxinvoicingagentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzmaxinvoicingagentResponse(fkiBillingentityinternalID int32, sBillingentityinternalDescriptionX string, iEzmaxinvoicingagentSession int32, iEzmaxinvoicingagentCloned int32, iEzmaxinvoicingagentInvoice int32, iEzmaxinvoicingagentInscription int32, iEzmaxinvoicingagentInscriptionactive int32, iEzmaxinvoicingagentSale int32, iEzmaxinvoicingagentOtherincome int32, iEzmaxinvoicingagentCommissioncalculation int32, iEzmaxinvoicingagentEzsigndocument int32, bEzmaxinvoicingagentEzsignaccount bool, bEzmaxinvoicingagentBillableezmax bool, eEzmaxinvoicingagentVariationezmax FieldEEzmaxinvoicingagentVariationezmax, bEzmaxinvoicingagentBillableezsign bool, eEzmaxinvoicingagentVariationezsign FieldEEzmaxinvoicingagentVariationezsign) *EzmaxinvoicingagentResponse {
	this := EzmaxinvoicingagentResponse{}
	this.FkiBillingentityinternalID = fkiBillingentityinternalID
	this.SBillingentityinternalDescriptionX = sBillingentityinternalDescriptionX
	this.IEzmaxinvoicingagentSession = iEzmaxinvoicingagentSession
	this.IEzmaxinvoicingagentCloned = iEzmaxinvoicingagentCloned
	this.IEzmaxinvoicingagentInvoice = iEzmaxinvoicingagentInvoice
	this.IEzmaxinvoicingagentInscription = iEzmaxinvoicingagentInscription
	this.IEzmaxinvoicingagentInscriptionactive = iEzmaxinvoicingagentInscriptionactive
	this.IEzmaxinvoicingagentSale = iEzmaxinvoicingagentSale
	this.IEzmaxinvoicingagentOtherincome = iEzmaxinvoicingagentOtherincome
	this.IEzmaxinvoicingagentCommissioncalculation = iEzmaxinvoicingagentCommissioncalculation
	this.IEzmaxinvoicingagentEzsigndocument = iEzmaxinvoicingagentEzsigndocument
	this.BEzmaxinvoicingagentEzsignaccount = bEzmaxinvoicingagentEzsignaccount
	this.BEzmaxinvoicingagentBillableezmax = bEzmaxinvoicingagentBillableezmax
	this.EEzmaxinvoicingagentVariationezmax = eEzmaxinvoicingagentVariationezmax
	this.BEzmaxinvoicingagentBillableezsign = bEzmaxinvoicingagentBillableezsign
	this.EEzmaxinvoicingagentVariationezsign = eEzmaxinvoicingagentVariationezsign
	return &this
}

// NewEzmaxinvoicingagentResponseWithDefaults instantiates a new EzmaxinvoicingagentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzmaxinvoicingagentResponseWithDefaults() *EzmaxinvoicingagentResponse {
	this := EzmaxinvoicingagentResponse{}
	return &this
}

// GetPkiEzmaxinvoicingagentID returns the PkiEzmaxinvoicingagentID field value if set, zero value otherwise.
func (o *EzmaxinvoicingagentResponse) GetPkiEzmaxinvoicingagentID() int32 {
	if o == nil || IsNil(o.PkiEzmaxinvoicingagentID) {
		var ret int32
		return ret
	}
	return *o.PkiEzmaxinvoicingagentID
}

// GetPkiEzmaxinvoicingagentIDOk returns a tuple with the PkiEzmaxinvoicingagentID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingagentResponse) GetPkiEzmaxinvoicingagentIDOk() (*int32, bool) {
	if o == nil || IsNil(o.PkiEzmaxinvoicingagentID) {
		return nil, false
	}
	return o.PkiEzmaxinvoicingagentID, true
}

// HasPkiEzmaxinvoicingagentID returns a boolean if a field has been set.
func (o *EzmaxinvoicingagentResponse) HasPkiEzmaxinvoicingagentID() bool {
	if o != nil && !IsNil(o.PkiEzmaxinvoicingagentID) {
		return true
	}

	return false
}

// SetPkiEzmaxinvoicingagentID gets a reference to the given int32 and assigns it to the PkiEzmaxinvoicingagentID field.
func (o *EzmaxinvoicingagentResponse) SetPkiEzmaxinvoicingagentID(v int32) {
	o.PkiEzmaxinvoicingagentID = &v
}

// GetFkiEzmaxinvoicingID returns the FkiEzmaxinvoicingID field value if set, zero value otherwise.
func (o *EzmaxinvoicingagentResponse) GetFkiEzmaxinvoicingID() int32 {
	if o == nil || IsNil(o.FkiEzmaxinvoicingID) {
		var ret int32
		return ret
	}
	return *o.FkiEzmaxinvoicingID
}

// GetFkiEzmaxinvoicingIDOk returns a tuple with the FkiEzmaxinvoicingID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingagentResponse) GetFkiEzmaxinvoicingIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiEzmaxinvoicingID) {
		return nil, false
	}
	return o.FkiEzmaxinvoicingID, true
}

// HasFkiEzmaxinvoicingID returns a boolean if a field has been set.
func (o *EzmaxinvoicingagentResponse) HasFkiEzmaxinvoicingID() bool {
	if o != nil && !IsNil(o.FkiEzmaxinvoicingID) {
		return true
	}

	return false
}

// SetFkiEzmaxinvoicingID gets a reference to the given int32 and assigns it to the FkiEzmaxinvoicingID field.
func (o *EzmaxinvoicingagentResponse) SetFkiEzmaxinvoicingID(v int32) {
	o.FkiEzmaxinvoicingID = &v
}

// GetFkiBillingentityinternalID returns the FkiBillingentityinternalID field value
func (o *EzmaxinvoicingagentResponse) GetFkiBillingentityinternalID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiBillingentityinternalID
}

// GetFkiBillingentityinternalIDOk returns a tuple with the FkiBillingentityinternalID field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingagentResponse) GetFkiBillingentityinternalIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiBillingentityinternalID, true
}

// SetFkiBillingentityinternalID sets field value
func (o *EzmaxinvoicingagentResponse) SetFkiBillingentityinternalID(v int32) {
	o.FkiBillingentityinternalID = v
}

// GetSBillingentityinternalDescriptionX returns the SBillingentityinternalDescriptionX field value
func (o *EzmaxinvoicingagentResponse) GetSBillingentityinternalDescriptionX() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SBillingentityinternalDescriptionX
}

// GetSBillingentityinternalDescriptionXOk returns a tuple with the SBillingentityinternalDescriptionX field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingagentResponse) GetSBillingentityinternalDescriptionXOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SBillingentityinternalDescriptionX, true
}

// SetSBillingentityinternalDescriptionX sets field value
func (o *EzmaxinvoicingagentResponse) SetSBillingentityinternalDescriptionX(v string) {
	o.SBillingentityinternalDescriptionX = v
}

// GetFkiAgentID returns the FkiAgentID field value if set, zero value otherwise.
func (o *EzmaxinvoicingagentResponse) GetFkiAgentID() int32 {
	if o == nil || IsNil(o.FkiAgentID) {
		var ret int32
		return ret
	}
	return *o.FkiAgentID
}

// GetFkiAgentIDOk returns a tuple with the FkiAgentID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingagentResponse) GetFkiAgentIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiAgentID) {
		return nil, false
	}
	return o.FkiAgentID, true
}

// HasFkiAgentID returns a boolean if a field has been set.
func (o *EzmaxinvoicingagentResponse) HasFkiAgentID() bool {
	if o != nil && !IsNil(o.FkiAgentID) {
		return true
	}

	return false
}

// SetFkiAgentID gets a reference to the given int32 and assigns it to the FkiAgentID field.
func (o *EzmaxinvoicingagentResponse) SetFkiAgentID(v int32) {
	o.FkiAgentID = &v
}

// GetFkiBrokerID returns the FkiBrokerID field value if set, zero value otherwise.
func (o *EzmaxinvoicingagentResponse) GetFkiBrokerID() int32 {
	if o == nil || IsNil(o.FkiBrokerID) {
		var ret int32
		return ret
	}
	return *o.FkiBrokerID
}

// GetFkiBrokerIDOk returns a tuple with the FkiBrokerID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingagentResponse) GetFkiBrokerIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiBrokerID) {
		return nil, false
	}
	return o.FkiBrokerID, true
}

// HasFkiBrokerID returns a boolean if a field has been set.
func (o *EzmaxinvoicingagentResponse) HasFkiBrokerID() bool {
	if o != nil && !IsNil(o.FkiBrokerID) {
		return true
	}

	return false
}

// SetFkiBrokerID gets a reference to the given int32 and assigns it to the FkiBrokerID field.
func (o *EzmaxinvoicingagentResponse) SetFkiBrokerID(v int32) {
	o.FkiBrokerID = &v
}

// GetIEzmaxinvoicingagentSession returns the IEzmaxinvoicingagentSession field value
func (o *EzmaxinvoicingagentResponse) GetIEzmaxinvoicingagentSession() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzmaxinvoicingagentSession
}

// GetIEzmaxinvoicingagentSessionOk returns a tuple with the IEzmaxinvoicingagentSession field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingagentResponse) GetIEzmaxinvoicingagentSessionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzmaxinvoicingagentSession, true
}

// SetIEzmaxinvoicingagentSession sets field value
func (o *EzmaxinvoicingagentResponse) SetIEzmaxinvoicingagentSession(v int32) {
	o.IEzmaxinvoicingagentSession = v
}

// GetIEzmaxinvoicingagentCloned returns the IEzmaxinvoicingagentCloned field value
func (o *EzmaxinvoicingagentResponse) GetIEzmaxinvoicingagentCloned() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzmaxinvoicingagentCloned
}

// GetIEzmaxinvoicingagentClonedOk returns a tuple with the IEzmaxinvoicingagentCloned field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingagentResponse) GetIEzmaxinvoicingagentClonedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzmaxinvoicingagentCloned, true
}

// SetIEzmaxinvoicingagentCloned sets field value
func (o *EzmaxinvoicingagentResponse) SetIEzmaxinvoicingagentCloned(v int32) {
	o.IEzmaxinvoicingagentCloned = v
}

// GetIEzmaxinvoicingagentInvoice returns the IEzmaxinvoicingagentInvoice field value
func (o *EzmaxinvoicingagentResponse) GetIEzmaxinvoicingagentInvoice() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzmaxinvoicingagentInvoice
}

// GetIEzmaxinvoicingagentInvoiceOk returns a tuple with the IEzmaxinvoicingagentInvoice field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingagentResponse) GetIEzmaxinvoicingagentInvoiceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzmaxinvoicingagentInvoice, true
}

// SetIEzmaxinvoicingagentInvoice sets field value
func (o *EzmaxinvoicingagentResponse) SetIEzmaxinvoicingagentInvoice(v int32) {
	o.IEzmaxinvoicingagentInvoice = v
}

// GetIEzmaxinvoicingagentInscription returns the IEzmaxinvoicingagentInscription field value
func (o *EzmaxinvoicingagentResponse) GetIEzmaxinvoicingagentInscription() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzmaxinvoicingagentInscription
}

// GetIEzmaxinvoicingagentInscriptionOk returns a tuple with the IEzmaxinvoicingagentInscription field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingagentResponse) GetIEzmaxinvoicingagentInscriptionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzmaxinvoicingagentInscription, true
}

// SetIEzmaxinvoicingagentInscription sets field value
func (o *EzmaxinvoicingagentResponse) SetIEzmaxinvoicingagentInscription(v int32) {
	o.IEzmaxinvoicingagentInscription = v
}

// GetIEzmaxinvoicingagentInscriptionactive returns the IEzmaxinvoicingagentInscriptionactive field value
func (o *EzmaxinvoicingagentResponse) GetIEzmaxinvoicingagentInscriptionactive() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzmaxinvoicingagentInscriptionactive
}

// GetIEzmaxinvoicingagentInscriptionactiveOk returns a tuple with the IEzmaxinvoicingagentInscriptionactive field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingagentResponse) GetIEzmaxinvoicingagentInscriptionactiveOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzmaxinvoicingagentInscriptionactive, true
}

// SetIEzmaxinvoicingagentInscriptionactive sets field value
func (o *EzmaxinvoicingagentResponse) SetIEzmaxinvoicingagentInscriptionactive(v int32) {
	o.IEzmaxinvoicingagentInscriptionactive = v
}

// GetIEzmaxinvoicingagentSale returns the IEzmaxinvoicingagentSale field value
func (o *EzmaxinvoicingagentResponse) GetIEzmaxinvoicingagentSale() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzmaxinvoicingagentSale
}

// GetIEzmaxinvoicingagentSaleOk returns a tuple with the IEzmaxinvoicingagentSale field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingagentResponse) GetIEzmaxinvoicingagentSaleOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzmaxinvoicingagentSale, true
}

// SetIEzmaxinvoicingagentSale sets field value
func (o *EzmaxinvoicingagentResponse) SetIEzmaxinvoicingagentSale(v int32) {
	o.IEzmaxinvoicingagentSale = v
}

// GetIEzmaxinvoicingagentOtherincome returns the IEzmaxinvoicingagentOtherincome field value
func (o *EzmaxinvoicingagentResponse) GetIEzmaxinvoicingagentOtherincome() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzmaxinvoicingagentOtherincome
}

// GetIEzmaxinvoicingagentOtherincomeOk returns a tuple with the IEzmaxinvoicingagentOtherincome field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingagentResponse) GetIEzmaxinvoicingagentOtherincomeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzmaxinvoicingagentOtherincome, true
}

// SetIEzmaxinvoicingagentOtherincome sets field value
func (o *EzmaxinvoicingagentResponse) SetIEzmaxinvoicingagentOtherincome(v int32) {
	o.IEzmaxinvoicingagentOtherincome = v
}

// GetIEzmaxinvoicingagentCommissioncalculation returns the IEzmaxinvoicingagentCommissioncalculation field value
func (o *EzmaxinvoicingagentResponse) GetIEzmaxinvoicingagentCommissioncalculation() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzmaxinvoicingagentCommissioncalculation
}

// GetIEzmaxinvoicingagentCommissioncalculationOk returns a tuple with the IEzmaxinvoicingagentCommissioncalculation field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingagentResponse) GetIEzmaxinvoicingagentCommissioncalculationOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzmaxinvoicingagentCommissioncalculation, true
}

// SetIEzmaxinvoicingagentCommissioncalculation sets field value
func (o *EzmaxinvoicingagentResponse) SetIEzmaxinvoicingagentCommissioncalculation(v int32) {
	o.IEzmaxinvoicingagentCommissioncalculation = v
}

// GetIEzmaxinvoicingagentEzsigndocument returns the IEzmaxinvoicingagentEzsigndocument field value
func (o *EzmaxinvoicingagentResponse) GetIEzmaxinvoicingagentEzsigndocument() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzmaxinvoicingagentEzsigndocument
}

// GetIEzmaxinvoicingagentEzsigndocumentOk returns a tuple with the IEzmaxinvoicingagentEzsigndocument field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingagentResponse) GetIEzmaxinvoicingagentEzsigndocumentOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzmaxinvoicingagentEzsigndocument, true
}

// SetIEzmaxinvoicingagentEzsigndocument sets field value
func (o *EzmaxinvoicingagentResponse) SetIEzmaxinvoicingagentEzsigndocument(v int32) {
	o.IEzmaxinvoicingagentEzsigndocument = v
}

// GetBEzmaxinvoicingagentEzsignaccount returns the BEzmaxinvoicingagentEzsignaccount field value
func (o *EzmaxinvoicingagentResponse) GetBEzmaxinvoicingagentEzsignaccount() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BEzmaxinvoicingagentEzsignaccount
}

// GetBEzmaxinvoicingagentEzsignaccountOk returns a tuple with the BEzmaxinvoicingagentEzsignaccount field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingagentResponse) GetBEzmaxinvoicingagentEzsignaccountOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BEzmaxinvoicingagentEzsignaccount, true
}

// SetBEzmaxinvoicingagentEzsignaccount sets field value
func (o *EzmaxinvoicingagentResponse) SetBEzmaxinvoicingagentEzsignaccount(v bool) {
	o.BEzmaxinvoicingagentEzsignaccount = v
}

// GetBEzmaxinvoicingagentBillableezmax returns the BEzmaxinvoicingagentBillableezmax field value
func (o *EzmaxinvoicingagentResponse) GetBEzmaxinvoicingagentBillableezmax() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BEzmaxinvoicingagentBillableezmax
}

// GetBEzmaxinvoicingagentBillableezmaxOk returns a tuple with the BEzmaxinvoicingagentBillableezmax field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingagentResponse) GetBEzmaxinvoicingagentBillableezmaxOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BEzmaxinvoicingagentBillableezmax, true
}

// SetBEzmaxinvoicingagentBillableezmax sets field value
func (o *EzmaxinvoicingagentResponse) SetBEzmaxinvoicingagentBillableezmax(v bool) {
	o.BEzmaxinvoicingagentBillableezmax = v
}

// GetEEzmaxinvoicingagentVariationezmax returns the EEzmaxinvoicingagentVariationezmax field value
func (o *EzmaxinvoicingagentResponse) GetEEzmaxinvoicingagentVariationezmax() FieldEEzmaxinvoicingagentVariationezmax {
	if o == nil {
		var ret FieldEEzmaxinvoicingagentVariationezmax
		return ret
	}

	return o.EEzmaxinvoicingagentVariationezmax
}

// GetEEzmaxinvoicingagentVariationezmaxOk returns a tuple with the EEzmaxinvoicingagentVariationezmax field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingagentResponse) GetEEzmaxinvoicingagentVariationezmaxOk() (*FieldEEzmaxinvoicingagentVariationezmax, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EEzmaxinvoicingagentVariationezmax, true
}

// SetEEzmaxinvoicingagentVariationezmax sets field value
func (o *EzmaxinvoicingagentResponse) SetEEzmaxinvoicingagentVariationezmax(v FieldEEzmaxinvoicingagentVariationezmax) {
	o.EEzmaxinvoicingagentVariationezmax = v
}

// GetBEzmaxinvoicingagentBillableezsign returns the BEzmaxinvoicingagentBillableezsign field value
func (o *EzmaxinvoicingagentResponse) GetBEzmaxinvoicingagentBillableezsign() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BEzmaxinvoicingagentBillableezsign
}

// GetBEzmaxinvoicingagentBillableezsignOk returns a tuple with the BEzmaxinvoicingagentBillableezsign field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingagentResponse) GetBEzmaxinvoicingagentBillableezsignOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BEzmaxinvoicingagentBillableezsign, true
}

// SetBEzmaxinvoicingagentBillableezsign sets field value
func (o *EzmaxinvoicingagentResponse) SetBEzmaxinvoicingagentBillableezsign(v bool) {
	o.BEzmaxinvoicingagentBillableezsign = v
}

// GetEEzmaxinvoicingagentVariationezsign returns the EEzmaxinvoicingagentVariationezsign field value
func (o *EzmaxinvoicingagentResponse) GetEEzmaxinvoicingagentVariationezsign() FieldEEzmaxinvoicingagentVariationezsign {
	if o == nil {
		var ret FieldEEzmaxinvoicingagentVariationezsign
		return ret
	}

	return o.EEzmaxinvoicingagentVariationezsign
}

// GetEEzmaxinvoicingagentVariationezsignOk returns a tuple with the EEzmaxinvoicingagentVariationezsign field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingagentResponse) GetEEzmaxinvoicingagentVariationezsignOk() (*FieldEEzmaxinvoicingagentVariationezsign, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EEzmaxinvoicingagentVariationezsign, true
}

// SetEEzmaxinvoicingagentVariationezsign sets field value
func (o *EzmaxinvoicingagentResponse) SetEEzmaxinvoicingagentVariationezsign(v FieldEEzmaxinvoicingagentVariationezsign) {
	o.EEzmaxinvoicingagentVariationezsign = v
}

func (o EzmaxinvoicingagentResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzmaxinvoicingagentResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PkiEzmaxinvoicingagentID) {
		toSerialize["pkiEzmaxinvoicingagentID"] = o.PkiEzmaxinvoicingagentID
	}
	if !IsNil(o.FkiEzmaxinvoicingID) {
		toSerialize["fkiEzmaxinvoicingID"] = o.FkiEzmaxinvoicingID
	}
	toSerialize["fkiBillingentityinternalID"] = o.FkiBillingentityinternalID
	toSerialize["sBillingentityinternalDescriptionX"] = o.SBillingentityinternalDescriptionX
	if !IsNil(o.FkiAgentID) {
		toSerialize["fkiAgentID"] = o.FkiAgentID
	}
	if !IsNil(o.FkiBrokerID) {
		toSerialize["fkiBrokerID"] = o.FkiBrokerID
	}
	toSerialize["iEzmaxinvoicingagentSession"] = o.IEzmaxinvoicingagentSession
	toSerialize["iEzmaxinvoicingagentCloned"] = o.IEzmaxinvoicingagentCloned
	toSerialize["iEzmaxinvoicingagentInvoice"] = o.IEzmaxinvoicingagentInvoice
	toSerialize["iEzmaxinvoicingagentInscription"] = o.IEzmaxinvoicingagentInscription
	toSerialize["iEzmaxinvoicingagentInscriptionactive"] = o.IEzmaxinvoicingagentInscriptionactive
	toSerialize["iEzmaxinvoicingagentSale"] = o.IEzmaxinvoicingagentSale
	toSerialize["iEzmaxinvoicingagentOtherincome"] = o.IEzmaxinvoicingagentOtherincome
	toSerialize["iEzmaxinvoicingagentCommissioncalculation"] = o.IEzmaxinvoicingagentCommissioncalculation
	toSerialize["iEzmaxinvoicingagentEzsigndocument"] = o.IEzmaxinvoicingagentEzsigndocument
	toSerialize["bEzmaxinvoicingagentEzsignaccount"] = o.BEzmaxinvoicingagentEzsignaccount
	toSerialize["bEzmaxinvoicingagentBillableezmax"] = o.BEzmaxinvoicingagentBillableezmax
	toSerialize["eEzmaxinvoicingagentVariationezmax"] = o.EEzmaxinvoicingagentVariationezmax
	toSerialize["bEzmaxinvoicingagentBillableezsign"] = o.BEzmaxinvoicingagentBillableezsign
	toSerialize["eEzmaxinvoicingagentVariationezsign"] = o.EEzmaxinvoicingagentVariationezsign
	return toSerialize, nil
}

func (o *EzmaxinvoicingagentResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fkiBillingentityinternalID",
		"sBillingentityinternalDescriptionX",
		"iEzmaxinvoicingagentSession",
		"iEzmaxinvoicingagentCloned",
		"iEzmaxinvoicingagentInvoice",
		"iEzmaxinvoicingagentInscription",
		"iEzmaxinvoicingagentInscriptionactive",
		"iEzmaxinvoicingagentSale",
		"iEzmaxinvoicingagentOtherincome",
		"iEzmaxinvoicingagentCommissioncalculation",
		"iEzmaxinvoicingagentEzsigndocument",
		"bEzmaxinvoicingagentEzsignaccount",
		"bEzmaxinvoicingagentBillableezmax",
		"eEzmaxinvoicingagentVariationezmax",
		"bEzmaxinvoicingagentBillableezsign",
		"eEzmaxinvoicingagentVariationezsign",
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

	varEzmaxinvoicingagentResponse := _EzmaxinvoicingagentResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzmaxinvoicingagentResponse)

	if err != nil {
		return err
	}

	*o = EzmaxinvoicingagentResponse(varEzmaxinvoicingagentResponse)

	return err
}

type NullableEzmaxinvoicingagentResponse struct {
	value *EzmaxinvoicingagentResponse
	isSet bool
}

func (v NullableEzmaxinvoicingagentResponse) Get() *EzmaxinvoicingagentResponse {
	return v.value
}

func (v *NullableEzmaxinvoicingagentResponse) Set(val *EzmaxinvoicingagentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEzmaxinvoicingagentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEzmaxinvoicingagentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzmaxinvoicingagentResponse(val *EzmaxinvoicingagentResponse) *NullableEzmaxinvoicingagentResponse {
	return &NullableEzmaxinvoicingagentResponse{value: val, isSet: true}
}

func (v NullableEzmaxinvoicingagentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzmaxinvoicingagentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


