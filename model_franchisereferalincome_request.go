/*
 * eZmax API Definition
 *
 * This API expose all the functionnalities for the eZmax and eZsign applications.
 *
 * API version: 1.0.44
 * Contact: support-api@ezmax.ca
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
)

// FranchisereferalincomeRequest An Franchisereferalincome Object
type FranchisereferalincomeRequest struct {
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

// NewFranchisereferalincomeRequest instantiates a new FranchisereferalincomeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFranchisereferalincomeRequest(fkiFranchisebrokerID int32, fkiFranchisereferalincomeprogramID int32, fkiPeriodID int32, dFranchisereferalincomeLoan string, dFranchisereferalincomeFranchiseamount string, dFranchisereferalincomeFranchisoramount string, dFranchisereferalincomeAgentamount string, dtFranchisereferalincomeDisbursed string, tFranchisereferalincomeComment string, fkiFranchiseofficeID int32, sFranchisereferalincomeRemoteid string) *FranchisereferalincomeRequest {
	this := FranchisereferalincomeRequest{}
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

// NewFranchisereferalincomeRequestWithDefaults instantiates a new FranchisereferalincomeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFranchisereferalincomeRequestWithDefaults() *FranchisereferalincomeRequest {
	this := FranchisereferalincomeRequest{}
	return &this
}

// GetFkiFranchisebrokerID returns the FkiFranchisebrokerID field value
func (o *FranchisereferalincomeRequest) GetFkiFranchisebrokerID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiFranchisebrokerID
}

// GetFkiFranchisebrokerIDOk returns a tuple with the FkiFranchisebrokerID field value
// and a boolean to check if the value has been set.
func (o *FranchisereferalincomeRequest) GetFkiFranchisebrokerIDOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FkiFranchisebrokerID, true
}

// SetFkiFranchisebrokerID sets field value
func (o *FranchisereferalincomeRequest) SetFkiFranchisebrokerID(v int32) {
	o.FkiFranchisebrokerID = v
}

// GetFkiFranchisereferalincomeprogramID returns the FkiFranchisereferalincomeprogramID field value
func (o *FranchisereferalincomeRequest) GetFkiFranchisereferalincomeprogramID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiFranchisereferalincomeprogramID
}

// GetFkiFranchisereferalincomeprogramIDOk returns a tuple with the FkiFranchisereferalincomeprogramID field value
// and a boolean to check if the value has been set.
func (o *FranchisereferalincomeRequest) GetFkiFranchisereferalincomeprogramIDOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FkiFranchisereferalincomeprogramID, true
}

// SetFkiFranchisereferalincomeprogramID sets field value
func (o *FranchisereferalincomeRequest) SetFkiFranchisereferalincomeprogramID(v int32) {
	o.FkiFranchisereferalincomeprogramID = v
}

// GetFkiPeriodID returns the FkiPeriodID field value
func (o *FranchisereferalincomeRequest) GetFkiPeriodID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiPeriodID
}

// GetFkiPeriodIDOk returns a tuple with the FkiPeriodID field value
// and a boolean to check if the value has been set.
func (o *FranchisereferalincomeRequest) GetFkiPeriodIDOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FkiPeriodID, true
}

// SetFkiPeriodID sets field value
func (o *FranchisereferalincomeRequest) SetFkiPeriodID(v int32) {
	o.FkiPeriodID = v
}

// GetDFranchisereferalincomeLoan returns the DFranchisereferalincomeLoan field value
func (o *FranchisereferalincomeRequest) GetDFranchisereferalincomeLoan() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DFranchisereferalincomeLoan
}

// GetDFranchisereferalincomeLoanOk returns a tuple with the DFranchisereferalincomeLoan field value
// and a boolean to check if the value has been set.
func (o *FranchisereferalincomeRequest) GetDFranchisereferalincomeLoanOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DFranchisereferalincomeLoan, true
}

// SetDFranchisereferalincomeLoan sets field value
func (o *FranchisereferalincomeRequest) SetDFranchisereferalincomeLoan(v string) {
	o.DFranchisereferalincomeLoan = v
}

// GetDFranchisereferalincomeFranchiseamount returns the DFranchisereferalincomeFranchiseamount field value
func (o *FranchisereferalincomeRequest) GetDFranchisereferalincomeFranchiseamount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DFranchisereferalincomeFranchiseamount
}

// GetDFranchisereferalincomeFranchiseamountOk returns a tuple with the DFranchisereferalincomeFranchiseamount field value
// and a boolean to check if the value has been set.
func (o *FranchisereferalincomeRequest) GetDFranchisereferalincomeFranchiseamountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DFranchisereferalincomeFranchiseamount, true
}

// SetDFranchisereferalincomeFranchiseamount sets field value
func (o *FranchisereferalincomeRequest) SetDFranchisereferalincomeFranchiseamount(v string) {
	o.DFranchisereferalincomeFranchiseamount = v
}

// GetDFranchisereferalincomeFranchisoramount returns the DFranchisereferalincomeFranchisoramount field value
func (o *FranchisereferalincomeRequest) GetDFranchisereferalincomeFranchisoramount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DFranchisereferalincomeFranchisoramount
}

// GetDFranchisereferalincomeFranchisoramountOk returns a tuple with the DFranchisereferalincomeFranchisoramount field value
// and a boolean to check if the value has been set.
func (o *FranchisereferalincomeRequest) GetDFranchisereferalincomeFranchisoramountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DFranchisereferalincomeFranchisoramount, true
}

// SetDFranchisereferalincomeFranchisoramount sets field value
func (o *FranchisereferalincomeRequest) SetDFranchisereferalincomeFranchisoramount(v string) {
	o.DFranchisereferalincomeFranchisoramount = v
}

// GetDFranchisereferalincomeAgentamount returns the DFranchisereferalincomeAgentamount field value
func (o *FranchisereferalincomeRequest) GetDFranchisereferalincomeAgentamount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DFranchisereferalincomeAgentamount
}

// GetDFranchisereferalincomeAgentamountOk returns a tuple with the DFranchisereferalincomeAgentamount field value
// and a boolean to check if the value has been set.
func (o *FranchisereferalincomeRequest) GetDFranchisereferalincomeAgentamountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DFranchisereferalincomeAgentamount, true
}

// SetDFranchisereferalincomeAgentamount sets field value
func (o *FranchisereferalincomeRequest) SetDFranchisereferalincomeAgentamount(v string) {
	o.DFranchisereferalincomeAgentamount = v
}

// GetDtFranchisereferalincomeDisbursed returns the DtFranchisereferalincomeDisbursed field value
func (o *FranchisereferalincomeRequest) GetDtFranchisereferalincomeDisbursed() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DtFranchisereferalincomeDisbursed
}

// GetDtFranchisereferalincomeDisbursedOk returns a tuple with the DtFranchisereferalincomeDisbursed field value
// and a boolean to check if the value has been set.
func (o *FranchisereferalincomeRequest) GetDtFranchisereferalincomeDisbursedOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DtFranchisereferalincomeDisbursed, true
}

// SetDtFranchisereferalincomeDisbursed sets field value
func (o *FranchisereferalincomeRequest) SetDtFranchisereferalincomeDisbursed(v string) {
	o.DtFranchisereferalincomeDisbursed = v
}

// GetTFranchisereferalincomeComment returns the TFranchisereferalincomeComment field value
func (o *FranchisereferalincomeRequest) GetTFranchisereferalincomeComment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TFranchisereferalincomeComment
}

// GetTFranchisereferalincomeCommentOk returns a tuple with the TFranchisereferalincomeComment field value
// and a boolean to check if the value has been set.
func (o *FranchisereferalincomeRequest) GetTFranchisereferalincomeCommentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TFranchisereferalincomeComment, true
}

// SetTFranchisereferalincomeComment sets field value
func (o *FranchisereferalincomeRequest) SetTFranchisereferalincomeComment(v string) {
	o.TFranchisereferalincomeComment = v
}

// GetFkiFranchiseofficeID returns the FkiFranchiseofficeID field value
func (o *FranchisereferalincomeRequest) GetFkiFranchiseofficeID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiFranchiseofficeID
}

// GetFkiFranchiseofficeIDOk returns a tuple with the FkiFranchiseofficeID field value
// and a boolean to check if the value has been set.
func (o *FranchisereferalincomeRequest) GetFkiFranchiseofficeIDOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FkiFranchiseofficeID, true
}

// SetFkiFranchiseofficeID sets field value
func (o *FranchisereferalincomeRequest) SetFkiFranchiseofficeID(v int32) {
	o.FkiFranchiseofficeID = v
}

// GetSFranchisereferalincomeRemoteid returns the SFranchisereferalincomeRemoteid field value
func (o *FranchisereferalincomeRequest) GetSFranchisereferalincomeRemoteid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SFranchisereferalincomeRemoteid
}

// GetSFranchisereferalincomeRemoteidOk returns a tuple with the SFranchisereferalincomeRemoteid field value
// and a boolean to check if the value has been set.
func (o *FranchisereferalincomeRequest) GetSFranchisereferalincomeRemoteidOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SFranchisereferalincomeRemoteid, true
}

// SetSFranchisereferalincomeRemoteid sets field value
func (o *FranchisereferalincomeRequest) SetSFranchisereferalincomeRemoteid(v string) {
	o.SFranchisereferalincomeRemoteid = v
}

func (o FranchisereferalincomeRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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

type NullableFranchisereferalincomeRequest struct {
	value *FranchisereferalincomeRequest
	isSet bool
}

func (v NullableFranchisereferalincomeRequest) Get() *FranchisereferalincomeRequest {
	return v.value
}

func (v *NullableFranchisereferalincomeRequest) Set(val *FranchisereferalincomeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableFranchisereferalincomeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableFranchisereferalincomeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFranchisereferalincomeRequest(val *FranchisereferalincomeRequest) *NullableFranchisereferalincomeRequest {
	return &NullableFranchisereferalincomeRequest{value: val, isSet: true}
}

func (v NullableFranchisereferalincomeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFranchisereferalincomeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


