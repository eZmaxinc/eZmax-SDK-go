/*
 * eZmax API Definition
 *
 * This API expose all the functionnalities for the eZmax and eZsign application.  We provide SDKs for customers. They are generated using OpenAPI codegen, we encourage customers to use them as we also provide samples for them.  You can choose to build your own implementation manually or can use any compatible OpenAPI 3.0 generator like Swagger Codegen, OpenAPI codegen or any commercial generators.  If you need helping understanding how to use this API, don't waste too much time looking for it. Contact support-api@ezmax.ca, we're here to help. We are developpers so we know programmers don't like bad documentation. If you don't find what you need in the documentation, let us know, we'll improve it and put you rapidly up on track.
 *
 * API version: 1.0.27
 * Contact: support-api@ezmax.ca
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxinc/eZmax-SDK-go

import (
	"encoding/json"
)

// FranchisereferalincomeRequestCompound A Franchisereferalincome Object and children to create a complete structure
type FranchisereferalincomeRequestCompound struct {
	ObjAddress *AddressRequest `json:"objAddress,omitempty"`
	AObjContact []ContactRequestCompound `json:"a_objContact"`
	// The unique ID of the Franchisebroker
	FkiFranchisebrokerID int32 `json:"fkiFranchisebrokerID"`
	// The unique ID of the Franchisereferalincomeprogram
	FkiFranchisereferalincomeprogramID int32 `json:"fkiFranchisereferalincomeprogramID"`
	// The unique ID of the Period
	FkiPeriodID int32 `json:"fkiPeriodID"`
	// The loan amount
	DFranchisereferalincomeLoan string `json:"dFranchisereferalincomeLoan"`
	// The amount that will be given to the franchise
	DFranchisereferalincomeFranchiseamount string `json:"dFranchisereferalincomeFranchiseamount"`
	// The amount that will be kept by the franchisor
	DFranchisereferalincomeFranchisoramount string `json:"dFranchisereferalincomeFranchisoramount"`
	// The amount that will be given to the agent
	DFranchisereferalincomeAgentamount string `json:"dFranchisereferalincomeAgentamount"`
	// The date the amounts were disbursed
	DtFranchisereferalincomeDisbursed string `json:"dtFranchisereferalincomeDisbursed"`
	// A comment about the transaction
	TFranchisereferalincomeComment string `json:"tFranchisereferalincomeComment"`
	// The unique ID of the Franchisereoffice
	FkiFranchiseofficeID int32 `json:"fkiFranchiseofficeID"`
	SFranchisereferalincomeRemoteid string `json:"sFranchisereferalincomeRemoteid"`
}

// NewFranchisereferalincomeRequestCompound instantiates a new FranchisereferalincomeRequestCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFranchisereferalincomeRequestCompound(aObjContact []ContactRequestCompound, fkiFranchisebrokerID int32, fkiFranchisereferalincomeprogramID int32, fkiPeriodID int32, dFranchisereferalincomeLoan string, dFranchisereferalincomeFranchiseamount string, dFranchisereferalincomeFranchisoramount string, dFranchisereferalincomeAgentamount string, dtFranchisereferalincomeDisbursed string, tFranchisereferalincomeComment string, fkiFranchiseofficeID int32, sFranchisereferalincomeRemoteid string, ) *FranchisereferalincomeRequestCompound {
	this := FranchisereferalincomeRequestCompound{}
	this.AObjContact = aObjContact
	this.FkiFranchisebrokerID = fkiFranchisebrokerID
	this.FkiFranchisereferalincomeprogramID = fkiFranchisereferalincomeprogramID
	this.FkiPeriodID = fkiPeriodID
	this.DFranchisereferalincomeLoan = dFranchisereferalincomeLoan
	this.DFranchisereferalincomeFranchiseamount = dFranchisereferalincomeFranchiseamount
	this.DFranchisereferalincomeFranchisoramount = dFranchisereferalincomeFranchisoramount
	this.DFranchisereferalincomeAgentamount = dFranchisereferalincomeAgentamount
	this.DtFranchisereferalincomeDisbursed = dtFranchisereferalincomeDisbursed
	this.TFranchisereferalincomeComment = tFranchisereferalincomeComment
	this.FkiFranchiseofficeID = fkiFranchiseofficeID
	this.SFranchisereferalincomeRemoteid = sFranchisereferalincomeRemoteid
	return &this
}

// NewFranchisereferalincomeRequestCompoundWithDefaults instantiates a new FranchisereferalincomeRequestCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFranchisereferalincomeRequestCompoundWithDefaults() *FranchisereferalincomeRequestCompound {
	this := FranchisereferalincomeRequestCompound{}
	return &this
}

// GetObjAddress returns the ObjAddress field value if set, zero value otherwise.
func (o *FranchisereferalincomeRequestCompound) GetObjAddress() AddressRequest {
	if o == nil || o.ObjAddress == nil {
		var ret AddressRequest
		return ret
	}
	return *o.ObjAddress
}

// GetObjAddressOk returns a tuple with the ObjAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FranchisereferalincomeRequestCompound) GetObjAddressOk() (*AddressRequest, bool) {
	if o == nil || o.ObjAddress == nil {
		return nil, false
	}
	return o.ObjAddress, true
}

// HasObjAddress returns a boolean if a field has been set.
func (o *FranchisereferalincomeRequestCompound) HasObjAddress() bool {
	if o != nil && o.ObjAddress != nil {
		return true
	}

	return false
}

// SetObjAddress gets a reference to the given AddressRequest and assigns it to the ObjAddress field.
func (o *FranchisereferalincomeRequestCompound) SetObjAddress(v AddressRequest) {
	o.ObjAddress = &v
}

// GetAObjContact returns the AObjContact field value
func (o *FranchisereferalincomeRequestCompound) GetAObjContact() []ContactRequestCompound {
	if o == nil  {
		var ret []ContactRequestCompound
		return ret
	}

	return o.AObjContact
}

// GetAObjContactOk returns a tuple with the AObjContact field value
// and a boolean to check if the value has been set.
func (o *FranchisereferalincomeRequestCompound) GetAObjContactOk() (*[]ContactRequestCompound, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AObjContact, true
}

// SetAObjContact sets field value
func (o *FranchisereferalincomeRequestCompound) SetAObjContact(v []ContactRequestCompound) {
	o.AObjContact = v
}

// GetFkiFranchisebrokerID returns the FkiFranchisebrokerID field value
func (o *FranchisereferalincomeRequestCompound) GetFkiFranchisebrokerID() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.FkiFranchisebrokerID
}

// GetFkiFranchisebrokerIDOk returns a tuple with the FkiFranchisebrokerID field value
// and a boolean to check if the value has been set.
func (o *FranchisereferalincomeRequestCompound) GetFkiFranchisebrokerIDOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FkiFranchisebrokerID, true
}

// SetFkiFranchisebrokerID sets field value
func (o *FranchisereferalincomeRequestCompound) SetFkiFranchisebrokerID(v int32) {
	o.FkiFranchisebrokerID = v
}

// GetFkiFranchisereferalincomeprogramID returns the FkiFranchisereferalincomeprogramID field value
func (o *FranchisereferalincomeRequestCompound) GetFkiFranchisereferalincomeprogramID() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.FkiFranchisereferalincomeprogramID
}

// GetFkiFranchisereferalincomeprogramIDOk returns a tuple with the FkiFranchisereferalincomeprogramID field value
// and a boolean to check if the value has been set.
func (o *FranchisereferalincomeRequestCompound) GetFkiFranchisereferalincomeprogramIDOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FkiFranchisereferalincomeprogramID, true
}

// SetFkiFranchisereferalincomeprogramID sets field value
func (o *FranchisereferalincomeRequestCompound) SetFkiFranchisereferalincomeprogramID(v int32) {
	o.FkiFranchisereferalincomeprogramID = v
}

// GetFkiPeriodID returns the FkiPeriodID field value
func (o *FranchisereferalincomeRequestCompound) GetFkiPeriodID() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.FkiPeriodID
}

// GetFkiPeriodIDOk returns a tuple with the FkiPeriodID field value
// and a boolean to check if the value has been set.
func (o *FranchisereferalincomeRequestCompound) GetFkiPeriodIDOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FkiPeriodID, true
}

// SetFkiPeriodID sets field value
func (o *FranchisereferalincomeRequestCompound) SetFkiPeriodID(v int32) {
	o.FkiPeriodID = v
}

// GetDFranchisereferalincomeLoan returns the DFranchisereferalincomeLoan field value
func (o *FranchisereferalincomeRequestCompound) GetDFranchisereferalincomeLoan() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.DFranchisereferalincomeLoan
}

// GetDFranchisereferalincomeLoanOk returns a tuple with the DFranchisereferalincomeLoan field value
// and a boolean to check if the value has been set.
func (o *FranchisereferalincomeRequestCompound) GetDFranchisereferalincomeLoanOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DFranchisereferalincomeLoan, true
}

// SetDFranchisereferalincomeLoan sets field value
func (o *FranchisereferalincomeRequestCompound) SetDFranchisereferalincomeLoan(v string) {
	o.DFranchisereferalincomeLoan = v
}

// GetDFranchisereferalincomeFranchiseamount returns the DFranchisereferalincomeFranchiseamount field value
func (o *FranchisereferalincomeRequestCompound) GetDFranchisereferalincomeFranchiseamount() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.DFranchisereferalincomeFranchiseamount
}

// GetDFranchisereferalincomeFranchiseamountOk returns a tuple with the DFranchisereferalincomeFranchiseamount field value
// and a boolean to check if the value has been set.
func (o *FranchisereferalincomeRequestCompound) GetDFranchisereferalincomeFranchiseamountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DFranchisereferalincomeFranchiseamount, true
}

// SetDFranchisereferalincomeFranchiseamount sets field value
func (o *FranchisereferalincomeRequestCompound) SetDFranchisereferalincomeFranchiseamount(v string) {
	o.DFranchisereferalincomeFranchiseamount = v
}

// GetDFranchisereferalincomeFranchisoramount returns the DFranchisereferalincomeFranchisoramount field value
func (o *FranchisereferalincomeRequestCompound) GetDFranchisereferalincomeFranchisoramount() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.DFranchisereferalincomeFranchisoramount
}

// GetDFranchisereferalincomeFranchisoramountOk returns a tuple with the DFranchisereferalincomeFranchisoramount field value
// and a boolean to check if the value has been set.
func (o *FranchisereferalincomeRequestCompound) GetDFranchisereferalincomeFranchisoramountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DFranchisereferalincomeFranchisoramount, true
}

// SetDFranchisereferalincomeFranchisoramount sets field value
func (o *FranchisereferalincomeRequestCompound) SetDFranchisereferalincomeFranchisoramount(v string) {
	o.DFranchisereferalincomeFranchisoramount = v
}

// GetDFranchisereferalincomeAgentamount returns the DFranchisereferalincomeAgentamount field value
func (o *FranchisereferalincomeRequestCompound) GetDFranchisereferalincomeAgentamount() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.DFranchisereferalincomeAgentamount
}

// GetDFranchisereferalincomeAgentamountOk returns a tuple with the DFranchisereferalincomeAgentamount field value
// and a boolean to check if the value has been set.
func (o *FranchisereferalincomeRequestCompound) GetDFranchisereferalincomeAgentamountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DFranchisereferalincomeAgentamount, true
}

// SetDFranchisereferalincomeAgentamount sets field value
func (o *FranchisereferalincomeRequestCompound) SetDFranchisereferalincomeAgentamount(v string) {
	o.DFranchisereferalincomeAgentamount = v
}

// GetDtFranchisereferalincomeDisbursed returns the DtFranchisereferalincomeDisbursed field value
func (o *FranchisereferalincomeRequestCompound) GetDtFranchisereferalincomeDisbursed() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.DtFranchisereferalincomeDisbursed
}

// GetDtFranchisereferalincomeDisbursedOk returns a tuple with the DtFranchisereferalincomeDisbursed field value
// and a boolean to check if the value has been set.
func (o *FranchisereferalincomeRequestCompound) GetDtFranchisereferalincomeDisbursedOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DtFranchisereferalincomeDisbursed, true
}

// SetDtFranchisereferalincomeDisbursed sets field value
func (o *FranchisereferalincomeRequestCompound) SetDtFranchisereferalincomeDisbursed(v string) {
	o.DtFranchisereferalincomeDisbursed = v
}

// GetTFranchisereferalincomeComment returns the TFranchisereferalincomeComment field value
func (o *FranchisereferalincomeRequestCompound) GetTFranchisereferalincomeComment() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.TFranchisereferalincomeComment
}

// GetTFranchisereferalincomeCommentOk returns a tuple with the TFranchisereferalincomeComment field value
// and a boolean to check if the value has been set.
func (o *FranchisereferalincomeRequestCompound) GetTFranchisereferalincomeCommentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TFranchisereferalincomeComment, true
}

// SetTFranchisereferalincomeComment sets field value
func (o *FranchisereferalincomeRequestCompound) SetTFranchisereferalincomeComment(v string) {
	o.TFranchisereferalincomeComment = v
}

// GetFkiFranchiseofficeID returns the FkiFranchiseofficeID field value
func (o *FranchisereferalincomeRequestCompound) GetFkiFranchiseofficeID() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.FkiFranchiseofficeID
}

// GetFkiFranchiseofficeIDOk returns a tuple with the FkiFranchiseofficeID field value
// and a boolean to check if the value has been set.
func (o *FranchisereferalincomeRequestCompound) GetFkiFranchiseofficeIDOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FkiFranchiseofficeID, true
}

// SetFkiFranchiseofficeID sets field value
func (o *FranchisereferalincomeRequestCompound) SetFkiFranchiseofficeID(v int32) {
	o.FkiFranchiseofficeID = v
}

// GetSFranchisereferalincomeRemoteid returns the SFranchisereferalincomeRemoteid field value
func (o *FranchisereferalincomeRequestCompound) GetSFranchisereferalincomeRemoteid() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.SFranchisereferalincomeRemoteid
}

// GetSFranchisereferalincomeRemoteidOk returns a tuple with the SFranchisereferalincomeRemoteid field value
// and a boolean to check if the value has been set.
func (o *FranchisereferalincomeRequestCompound) GetSFranchisereferalincomeRemoteidOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SFranchisereferalincomeRemoteid, true
}

// SetSFranchisereferalincomeRemoteid sets field value
func (o *FranchisereferalincomeRequestCompound) SetSFranchisereferalincomeRemoteid(v string) {
	o.SFranchisereferalincomeRemoteid = v
}

func (o FranchisereferalincomeRequestCompound) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ObjAddress != nil {
		toSerialize["objAddress"] = o.ObjAddress
	}
	if true {
		toSerialize["a_objContact"] = o.AObjContact
	}
	if true {
		toSerialize["fkiFranchisebrokerID"] = o.FkiFranchisebrokerID
	}
	if true {
		toSerialize["fkiFranchisereferalincomeprogramID"] = o.FkiFranchisereferalincomeprogramID
	}
	if true {
		toSerialize["fkiPeriodID"] = o.FkiPeriodID
	}
	if true {
		toSerialize["dFranchisereferalincomeLoan"] = o.DFranchisereferalincomeLoan
	}
	if true {
		toSerialize["dFranchisereferalincomeFranchiseamount"] = o.DFranchisereferalincomeFranchiseamount
	}
	if true {
		toSerialize["dFranchisereferalincomeFranchisoramount"] = o.DFranchisereferalincomeFranchisoramount
	}
	if true {
		toSerialize["dFranchisereferalincomeAgentamount"] = o.DFranchisereferalincomeAgentamount
	}
	if true {
		toSerialize["dtFranchisereferalincomeDisbursed"] = o.DtFranchisereferalincomeDisbursed
	}
	if true {
		toSerialize["tFranchisereferalincomeComment"] = o.TFranchisereferalincomeComment
	}
	if true {
		toSerialize["fkiFranchiseofficeID"] = o.FkiFranchiseofficeID
	}
	if true {
		toSerialize["sFranchisereferalincomeRemoteid"] = o.SFranchisereferalincomeRemoteid
	}
	return json.Marshal(toSerialize)
}

type NullableFranchisereferalincomeRequestCompound struct {
	value *FranchisereferalincomeRequestCompound
	isSet bool
}

func (v NullableFranchisereferalincomeRequestCompound) Get() *FranchisereferalincomeRequestCompound {
	return v.value
}

func (v *NullableFranchisereferalincomeRequestCompound) Set(val *FranchisereferalincomeRequestCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableFranchisereferalincomeRequestCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableFranchisereferalincomeRequestCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFranchisereferalincomeRequestCompound(val *FranchisereferalincomeRequestCompound) *NullableFranchisereferalincomeRequestCompound {
	return &NullableFranchisereferalincomeRequestCompound{value: val, isSet: true}
}

func (v NullableFranchisereferalincomeRequestCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFranchisereferalincomeRequestCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


