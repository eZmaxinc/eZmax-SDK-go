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

// checks if the EzmaxinvoicinguserResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzmaxinvoicinguserResponse{}

// EzmaxinvoicinguserResponse A Ezmaxinvoicinguser Object
type EzmaxinvoicinguserResponse struct {
	// The unique ID of the Ezmaxinvoicinguser
	PkiEzmaxinvoicinguserID *int32 `json:"pkiEzmaxinvoicinguserID,omitempty"`
	// The unique ID of the Ezmaxinvoicing
	FkiEzmaxinvoicingID *int32 `json:"fkiEzmaxinvoicingID,omitempty"`
	// The unique ID of the Billingentityinternal.
	FkiBillingentityinternalID int32 `json:"fkiBillingentityinternalID"`
	// The description of the Billingentityinternal in the language of the requester
	SBillingentityinternalDescriptionX string `json:"sBillingentityinternalDescriptionX"`
	// The unique ID of the User
	FkiUserID int32 `json:"fkiUserID"`
	// The number of ezsign documents
	IEzmaxinvoicinguserEzsigndocument int32 `json:"iEzmaxinvoicinguserEzsigndocument"`
	// Whether there is an eZsign account
	BEzmaxinvoicinguserEzsignaccount bool `json:"bEzmaxinvoicinguserEzsignaccount"`
	// Whether it is billable for eZsign
	BEzmaxinvoicinguserBillableezsign bool `json:"bEzmaxinvoicinguserBillableezsign"`
	EEzmaxinvoicinguserVariationezsign FieldEEzmaxinvoicinguserVariationezsign `json:"eEzmaxinvoicinguserVariationezsign"`
}

// NewEzmaxinvoicinguserResponse instantiates a new EzmaxinvoicinguserResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzmaxinvoicinguserResponse(fkiBillingentityinternalID int32, sBillingentityinternalDescriptionX string, fkiUserID int32, iEzmaxinvoicinguserEzsigndocument int32, bEzmaxinvoicinguserEzsignaccount bool, bEzmaxinvoicinguserBillableezsign bool, eEzmaxinvoicinguserVariationezsign FieldEEzmaxinvoicinguserVariationezsign) *EzmaxinvoicinguserResponse {
	this := EzmaxinvoicinguserResponse{}
	this.FkiBillingentityinternalID = fkiBillingentityinternalID
	this.SBillingentityinternalDescriptionX = sBillingentityinternalDescriptionX
	this.FkiUserID = fkiUserID
	this.IEzmaxinvoicinguserEzsigndocument = iEzmaxinvoicinguserEzsigndocument
	this.BEzmaxinvoicinguserEzsignaccount = bEzmaxinvoicinguserEzsignaccount
	this.BEzmaxinvoicinguserBillableezsign = bEzmaxinvoicinguserBillableezsign
	this.EEzmaxinvoicinguserVariationezsign = eEzmaxinvoicinguserVariationezsign
	return &this
}

// NewEzmaxinvoicinguserResponseWithDefaults instantiates a new EzmaxinvoicinguserResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzmaxinvoicinguserResponseWithDefaults() *EzmaxinvoicinguserResponse {
	this := EzmaxinvoicinguserResponse{}
	return &this
}

// GetPkiEzmaxinvoicinguserID returns the PkiEzmaxinvoicinguserID field value if set, zero value otherwise.
func (o *EzmaxinvoicinguserResponse) GetPkiEzmaxinvoicinguserID() int32 {
	if o == nil || IsNil(o.PkiEzmaxinvoicinguserID) {
		var ret int32
		return ret
	}
	return *o.PkiEzmaxinvoicinguserID
}

// GetPkiEzmaxinvoicinguserIDOk returns a tuple with the PkiEzmaxinvoicinguserID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicinguserResponse) GetPkiEzmaxinvoicinguserIDOk() (*int32, bool) {
	if o == nil || IsNil(o.PkiEzmaxinvoicinguserID) {
		return nil, false
	}
	return o.PkiEzmaxinvoicinguserID, true
}

// HasPkiEzmaxinvoicinguserID returns a boolean if a field has been set.
func (o *EzmaxinvoicinguserResponse) HasPkiEzmaxinvoicinguserID() bool {
	if o != nil && !IsNil(o.PkiEzmaxinvoicinguserID) {
		return true
	}

	return false
}

// SetPkiEzmaxinvoicinguserID gets a reference to the given int32 and assigns it to the PkiEzmaxinvoicinguserID field.
func (o *EzmaxinvoicinguserResponse) SetPkiEzmaxinvoicinguserID(v int32) {
	o.PkiEzmaxinvoicinguserID = &v
}

// GetFkiEzmaxinvoicingID returns the FkiEzmaxinvoicingID field value if set, zero value otherwise.
func (o *EzmaxinvoicinguserResponse) GetFkiEzmaxinvoicingID() int32 {
	if o == nil || IsNil(o.FkiEzmaxinvoicingID) {
		var ret int32
		return ret
	}
	return *o.FkiEzmaxinvoicingID
}

// GetFkiEzmaxinvoicingIDOk returns a tuple with the FkiEzmaxinvoicingID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicinguserResponse) GetFkiEzmaxinvoicingIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiEzmaxinvoicingID) {
		return nil, false
	}
	return o.FkiEzmaxinvoicingID, true
}

// HasFkiEzmaxinvoicingID returns a boolean if a field has been set.
func (o *EzmaxinvoicinguserResponse) HasFkiEzmaxinvoicingID() bool {
	if o != nil && !IsNil(o.FkiEzmaxinvoicingID) {
		return true
	}

	return false
}

// SetFkiEzmaxinvoicingID gets a reference to the given int32 and assigns it to the FkiEzmaxinvoicingID field.
func (o *EzmaxinvoicinguserResponse) SetFkiEzmaxinvoicingID(v int32) {
	o.FkiEzmaxinvoicingID = &v
}

// GetFkiBillingentityinternalID returns the FkiBillingentityinternalID field value
func (o *EzmaxinvoicinguserResponse) GetFkiBillingentityinternalID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiBillingentityinternalID
}

// GetFkiBillingentityinternalIDOk returns a tuple with the FkiBillingentityinternalID field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicinguserResponse) GetFkiBillingentityinternalIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiBillingentityinternalID, true
}

// SetFkiBillingentityinternalID sets field value
func (o *EzmaxinvoicinguserResponse) SetFkiBillingentityinternalID(v int32) {
	o.FkiBillingentityinternalID = v
}

// GetSBillingentityinternalDescriptionX returns the SBillingentityinternalDescriptionX field value
func (o *EzmaxinvoicinguserResponse) GetSBillingentityinternalDescriptionX() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SBillingentityinternalDescriptionX
}

// GetSBillingentityinternalDescriptionXOk returns a tuple with the SBillingentityinternalDescriptionX field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicinguserResponse) GetSBillingentityinternalDescriptionXOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SBillingentityinternalDescriptionX, true
}

// SetSBillingentityinternalDescriptionX sets field value
func (o *EzmaxinvoicinguserResponse) SetSBillingentityinternalDescriptionX(v string) {
	o.SBillingentityinternalDescriptionX = v
}

// GetFkiUserID returns the FkiUserID field value
func (o *EzmaxinvoicinguserResponse) GetFkiUserID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiUserID
}

// GetFkiUserIDOk returns a tuple with the FkiUserID field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicinguserResponse) GetFkiUserIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiUserID, true
}

// SetFkiUserID sets field value
func (o *EzmaxinvoicinguserResponse) SetFkiUserID(v int32) {
	o.FkiUserID = v
}

// GetIEzmaxinvoicinguserEzsigndocument returns the IEzmaxinvoicinguserEzsigndocument field value
func (o *EzmaxinvoicinguserResponse) GetIEzmaxinvoicinguserEzsigndocument() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzmaxinvoicinguserEzsigndocument
}

// GetIEzmaxinvoicinguserEzsigndocumentOk returns a tuple with the IEzmaxinvoicinguserEzsigndocument field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicinguserResponse) GetIEzmaxinvoicinguserEzsigndocumentOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzmaxinvoicinguserEzsigndocument, true
}

// SetIEzmaxinvoicinguserEzsigndocument sets field value
func (o *EzmaxinvoicinguserResponse) SetIEzmaxinvoicinguserEzsigndocument(v int32) {
	o.IEzmaxinvoicinguserEzsigndocument = v
}

// GetBEzmaxinvoicinguserEzsignaccount returns the BEzmaxinvoicinguserEzsignaccount field value
func (o *EzmaxinvoicinguserResponse) GetBEzmaxinvoicinguserEzsignaccount() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BEzmaxinvoicinguserEzsignaccount
}

// GetBEzmaxinvoicinguserEzsignaccountOk returns a tuple with the BEzmaxinvoicinguserEzsignaccount field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicinguserResponse) GetBEzmaxinvoicinguserEzsignaccountOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BEzmaxinvoicinguserEzsignaccount, true
}

// SetBEzmaxinvoicinguserEzsignaccount sets field value
func (o *EzmaxinvoicinguserResponse) SetBEzmaxinvoicinguserEzsignaccount(v bool) {
	o.BEzmaxinvoicinguserEzsignaccount = v
}

// GetBEzmaxinvoicinguserBillableezsign returns the BEzmaxinvoicinguserBillableezsign field value
func (o *EzmaxinvoicinguserResponse) GetBEzmaxinvoicinguserBillableezsign() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BEzmaxinvoicinguserBillableezsign
}

// GetBEzmaxinvoicinguserBillableezsignOk returns a tuple with the BEzmaxinvoicinguserBillableezsign field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicinguserResponse) GetBEzmaxinvoicinguserBillableezsignOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BEzmaxinvoicinguserBillableezsign, true
}

// SetBEzmaxinvoicinguserBillableezsign sets field value
func (o *EzmaxinvoicinguserResponse) SetBEzmaxinvoicinguserBillableezsign(v bool) {
	o.BEzmaxinvoicinguserBillableezsign = v
}

// GetEEzmaxinvoicinguserVariationezsign returns the EEzmaxinvoicinguserVariationezsign field value
func (o *EzmaxinvoicinguserResponse) GetEEzmaxinvoicinguserVariationezsign() FieldEEzmaxinvoicinguserVariationezsign {
	if o == nil {
		var ret FieldEEzmaxinvoicinguserVariationezsign
		return ret
	}

	return o.EEzmaxinvoicinguserVariationezsign
}

// GetEEzmaxinvoicinguserVariationezsignOk returns a tuple with the EEzmaxinvoicinguserVariationezsign field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicinguserResponse) GetEEzmaxinvoicinguserVariationezsignOk() (*FieldEEzmaxinvoicinguserVariationezsign, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EEzmaxinvoicinguserVariationezsign, true
}

// SetEEzmaxinvoicinguserVariationezsign sets field value
func (o *EzmaxinvoicinguserResponse) SetEEzmaxinvoicinguserVariationezsign(v FieldEEzmaxinvoicinguserVariationezsign) {
	o.EEzmaxinvoicinguserVariationezsign = v
}

func (o EzmaxinvoicinguserResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzmaxinvoicinguserResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PkiEzmaxinvoicinguserID) {
		toSerialize["pkiEzmaxinvoicinguserID"] = o.PkiEzmaxinvoicinguserID
	}
	if !IsNil(o.FkiEzmaxinvoicingID) {
		toSerialize["fkiEzmaxinvoicingID"] = o.FkiEzmaxinvoicingID
	}
	toSerialize["fkiBillingentityinternalID"] = o.FkiBillingentityinternalID
	toSerialize["sBillingentityinternalDescriptionX"] = o.SBillingentityinternalDescriptionX
	toSerialize["fkiUserID"] = o.FkiUserID
	toSerialize["iEzmaxinvoicinguserEzsigndocument"] = o.IEzmaxinvoicinguserEzsigndocument
	toSerialize["bEzmaxinvoicinguserEzsignaccount"] = o.BEzmaxinvoicinguserEzsignaccount
	toSerialize["bEzmaxinvoicinguserBillableezsign"] = o.BEzmaxinvoicinguserBillableezsign
	toSerialize["eEzmaxinvoicinguserVariationezsign"] = o.EEzmaxinvoicinguserVariationezsign
	return toSerialize, nil
}

type NullableEzmaxinvoicinguserResponse struct {
	value *EzmaxinvoicinguserResponse
	isSet bool
}

func (v NullableEzmaxinvoicinguserResponse) Get() *EzmaxinvoicinguserResponse {
	return v.value
}

func (v *NullableEzmaxinvoicinguserResponse) Set(val *EzmaxinvoicinguserResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEzmaxinvoicinguserResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEzmaxinvoicinguserResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzmaxinvoicinguserResponse(val *EzmaxinvoicinguserResponse) *NullableEzmaxinvoicinguserResponse {
	return &NullableEzmaxinvoicinguserResponse{value: val, isSet: true}
}

func (v NullableEzmaxinvoicinguserResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzmaxinvoicinguserResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


