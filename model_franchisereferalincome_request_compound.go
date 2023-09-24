/*
eZmax API Definition (Full)

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.2.0
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
)

// checks if the FranchisereferalincomeRequestCompound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FranchisereferalincomeRequestCompound{}

// FranchisereferalincomeRequestCompound A Franchisereferalincome Object and children to create a complete structure
type FranchisereferalincomeRequestCompound struct {
	// The unique ID of the Franchisereferalincome
	PkiFranchisereferalincomeID *int32 `json:"pkiFranchisereferalincomeID,omitempty"`
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
	// Comment about the transaction
	TFranchisereferalincomeComment string `json:"tFranchisereferalincomeComment"`
	// The unique ID of the Franchisereoffice
	FkiFranchiseofficeID int32 `json:"fkiFranchiseofficeID"`
	// 
	SFranchisereferalincomeRemoteid string `json:"sFranchisereferalincomeRemoteid"`
	ObjAddress AddressRequest `json:"objAddress"`
	AObjContact []ContactRequestCompound `json:"a_objContact"`
}

// NewFranchisereferalincomeRequestCompound instantiates a new FranchisereferalincomeRequestCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFranchisereferalincomeRequestCompound(fkiFranchisebrokerID int32, fkiFranchisereferalincomeprogramID int32, fkiPeriodID int32, dFranchisereferalincomeLoan string, dFranchisereferalincomeFranchiseamount string, dFranchisereferalincomeFranchisoramount string, dFranchisereferalincomeAgentamount string, dtFranchisereferalincomeDisbursed string, tFranchisereferalincomeComment string, fkiFranchiseofficeID int32, sFranchisereferalincomeRemoteid string, objAddress AddressRequest, aObjContact []ContactRequestCompound) *FranchisereferalincomeRequestCompound {
	this := FranchisereferalincomeRequestCompound{}
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
	this.ObjAddress = objAddress
	this.AObjContact = aObjContact
	return &this
}

// NewFranchisereferalincomeRequestCompoundWithDefaults instantiates a new FranchisereferalincomeRequestCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFranchisereferalincomeRequestCompoundWithDefaults() *FranchisereferalincomeRequestCompound {
	this := FranchisereferalincomeRequestCompound{}
	return &this
}

// GetPkiFranchisereferalincomeID returns the PkiFranchisereferalincomeID field value if set, zero value otherwise.
func (o *FranchisereferalincomeRequestCompound) GetPkiFranchisereferalincomeID() int32 {
	if o == nil || IsNil(o.PkiFranchisereferalincomeID) {
		var ret int32
		return ret
	}
	return *o.PkiFranchisereferalincomeID
}

// GetPkiFranchisereferalincomeIDOk returns a tuple with the PkiFranchisereferalincomeID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FranchisereferalincomeRequestCompound) GetPkiFranchisereferalincomeIDOk() (*int32, bool) {
	if o == nil || IsNil(o.PkiFranchisereferalincomeID) {
		return nil, false
	}
	return o.PkiFranchisereferalincomeID, true
}

// HasPkiFranchisereferalincomeID returns a boolean if a field has been set.
func (o *FranchisereferalincomeRequestCompound) HasPkiFranchisereferalincomeID() bool {
	if o != nil && !IsNil(o.PkiFranchisereferalincomeID) {
		return true
	}

	return false
}

// SetPkiFranchisereferalincomeID gets a reference to the given int32 and assigns it to the PkiFranchisereferalincomeID field.
func (o *FranchisereferalincomeRequestCompound) SetPkiFranchisereferalincomeID(v int32) {
	o.PkiFranchisereferalincomeID = &v
}

// GetFkiFranchisebrokerID returns the FkiFranchisebrokerID field value
func (o *FranchisereferalincomeRequestCompound) GetFkiFranchisebrokerID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiFranchisebrokerID
}

// GetFkiFranchisebrokerIDOk returns a tuple with the FkiFranchisebrokerID field value
// and a boolean to check if the value has been set.
func (o *FranchisereferalincomeRequestCompound) GetFkiFranchisebrokerIDOk() (*int32, bool) {
	if o == nil {
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
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiFranchisereferalincomeprogramID
}

// GetFkiFranchisereferalincomeprogramIDOk returns a tuple with the FkiFranchisereferalincomeprogramID field value
// and a boolean to check if the value has been set.
func (o *FranchisereferalincomeRequestCompound) GetFkiFranchisereferalincomeprogramIDOk() (*int32, bool) {
	if o == nil {
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
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiPeriodID
}

// GetFkiPeriodIDOk returns a tuple with the FkiPeriodID field value
// and a boolean to check if the value has been set.
func (o *FranchisereferalincomeRequestCompound) GetFkiPeriodIDOk() (*int32, bool) {
	if o == nil {
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
	if o == nil {
		var ret string
		return ret
	}

	return o.DFranchisereferalincomeLoan
}

// GetDFranchisereferalincomeLoanOk returns a tuple with the DFranchisereferalincomeLoan field value
// and a boolean to check if the value has been set.
func (o *FranchisereferalincomeRequestCompound) GetDFranchisereferalincomeLoanOk() (*string, bool) {
	if o == nil {
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
	if o == nil {
		var ret string
		return ret
	}

	return o.DFranchisereferalincomeFranchiseamount
}

// GetDFranchisereferalincomeFranchiseamountOk returns a tuple with the DFranchisereferalincomeFranchiseamount field value
// and a boolean to check if the value has been set.
func (o *FranchisereferalincomeRequestCompound) GetDFranchisereferalincomeFranchiseamountOk() (*string, bool) {
	if o == nil {
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
	if o == nil {
		var ret string
		return ret
	}

	return o.DFranchisereferalincomeFranchisoramount
}

// GetDFranchisereferalincomeFranchisoramountOk returns a tuple with the DFranchisereferalincomeFranchisoramount field value
// and a boolean to check if the value has been set.
func (o *FranchisereferalincomeRequestCompound) GetDFranchisereferalincomeFranchisoramountOk() (*string, bool) {
	if o == nil {
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
	if o == nil {
		var ret string
		return ret
	}

	return o.DFranchisereferalincomeAgentamount
}

// GetDFranchisereferalincomeAgentamountOk returns a tuple with the DFranchisereferalincomeAgentamount field value
// and a boolean to check if the value has been set.
func (o *FranchisereferalincomeRequestCompound) GetDFranchisereferalincomeAgentamountOk() (*string, bool) {
	if o == nil {
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
	if o == nil {
		var ret string
		return ret
	}

	return o.DtFranchisereferalincomeDisbursed
}

// GetDtFranchisereferalincomeDisbursedOk returns a tuple with the DtFranchisereferalincomeDisbursed field value
// and a boolean to check if the value has been set.
func (o *FranchisereferalincomeRequestCompound) GetDtFranchisereferalincomeDisbursedOk() (*string, bool) {
	if o == nil {
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
	if o == nil {
		var ret string
		return ret
	}

	return o.TFranchisereferalincomeComment
}

// GetTFranchisereferalincomeCommentOk returns a tuple with the TFranchisereferalincomeComment field value
// and a boolean to check if the value has been set.
func (o *FranchisereferalincomeRequestCompound) GetTFranchisereferalincomeCommentOk() (*string, bool) {
	if o == nil {
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
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiFranchiseofficeID
}

// GetFkiFranchiseofficeIDOk returns a tuple with the FkiFranchiseofficeID field value
// and a boolean to check if the value has been set.
func (o *FranchisereferalincomeRequestCompound) GetFkiFranchiseofficeIDOk() (*int32, bool) {
	if o == nil {
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
	if o == nil {
		var ret string
		return ret
	}

	return o.SFranchisereferalincomeRemoteid
}

// GetSFranchisereferalincomeRemoteidOk returns a tuple with the SFranchisereferalincomeRemoteid field value
// and a boolean to check if the value has been set.
func (o *FranchisereferalincomeRequestCompound) GetSFranchisereferalincomeRemoteidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SFranchisereferalincomeRemoteid, true
}

// SetSFranchisereferalincomeRemoteid sets field value
func (o *FranchisereferalincomeRequestCompound) SetSFranchisereferalincomeRemoteid(v string) {
	o.SFranchisereferalincomeRemoteid = v
}

// GetObjAddress returns the ObjAddress field value
func (o *FranchisereferalincomeRequestCompound) GetObjAddress() AddressRequest {
	if o == nil {
		var ret AddressRequest
		return ret
	}

	return o.ObjAddress
}

// GetObjAddressOk returns a tuple with the ObjAddress field value
// and a boolean to check if the value has been set.
func (o *FranchisereferalincomeRequestCompound) GetObjAddressOk() (*AddressRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjAddress, true
}

// SetObjAddress sets field value
func (o *FranchisereferalincomeRequestCompound) SetObjAddress(v AddressRequest) {
	o.ObjAddress = v
}

// GetAObjContact returns the AObjContact field value
func (o *FranchisereferalincomeRequestCompound) GetAObjContact() []ContactRequestCompound {
	if o == nil {
		var ret []ContactRequestCompound
		return ret
	}

	return o.AObjContact
}

// GetAObjContactOk returns a tuple with the AObjContact field value
// and a boolean to check if the value has been set.
func (o *FranchisereferalincomeRequestCompound) GetAObjContactOk() ([]ContactRequestCompound, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjContact, true
}

// SetAObjContact sets field value
func (o *FranchisereferalincomeRequestCompound) SetAObjContact(v []ContactRequestCompound) {
	o.AObjContact = v
}

func (o FranchisereferalincomeRequestCompound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FranchisereferalincomeRequestCompound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PkiFranchisereferalincomeID) {
		toSerialize["pkiFranchisereferalincomeID"] = o.PkiFranchisereferalincomeID
	}
	toSerialize["fkiFranchisebrokerID"] = o.FkiFranchisebrokerID
	toSerialize["fkiFranchisereferalincomeprogramID"] = o.FkiFranchisereferalincomeprogramID
	toSerialize["fkiPeriodID"] = o.FkiPeriodID
	toSerialize["dFranchisereferalincomeLoan"] = o.DFranchisereferalincomeLoan
	toSerialize["dFranchisereferalincomeFranchiseamount"] = o.DFranchisereferalincomeFranchiseamount
	toSerialize["dFranchisereferalincomeFranchisoramount"] = o.DFranchisereferalincomeFranchisoramount
	toSerialize["dFranchisereferalincomeAgentamount"] = o.DFranchisereferalincomeAgentamount
	toSerialize["dtFranchisereferalincomeDisbursed"] = o.DtFranchisereferalincomeDisbursed
	toSerialize["tFranchisereferalincomeComment"] = o.TFranchisereferalincomeComment
	toSerialize["fkiFranchiseofficeID"] = o.FkiFranchiseofficeID
	toSerialize["sFranchisereferalincomeRemoteid"] = o.SFranchisereferalincomeRemoteid
	toSerialize["objAddress"] = o.ObjAddress
	toSerialize["a_objContact"] = o.AObjContact
	return toSerialize, nil
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


