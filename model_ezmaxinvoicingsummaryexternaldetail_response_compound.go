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

// checks if the EzmaxinvoicingsummaryexternaldetailResponseCompound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzmaxinvoicingsummaryexternaldetailResponseCompound{}

// EzmaxinvoicingsummaryexternaldetailResponseCompound A Ezmaxinvoicingsummaryexternaldetail Object
type EzmaxinvoicingsummaryexternaldetailResponseCompound struct {
	// The unique ID of the Ezmaxinvoicingsummaryexternaldetail
	PkiEzmaxinvoicingsummaryexternaldetailID *int32 `json:"pkiEzmaxinvoicingsummaryexternaldetailID,omitempty"`
	// The unique ID of the Ezmaxinvoicingsummaryexternal
	FkiEzmaxinvoicingsummaryexternalID *int32 `json:"fkiEzmaxinvoicingsummaryexternalID,omitempty"`
	// The unique ID of the Ezmaxproduct
	FkiEzmaxproductID int32 `json:"fkiEzmaxproductID"`
	// The description of the Ezmaxproduct in the language of the requester
	SEzmaxproductDescriptionX string `json:"sEzmaxproductDescriptionX"`
	// The count item invoiced for the product
	DEzmaxinvoicingsummaryexternaldetailCountreal string `json:"dEzmaxinvoicingsummaryexternaldetailCountreal"`
	// The subtotal invoiced for the product
	DEzmaxinvoicingsummaryexternaldetailSubtotal string `json:"dEzmaxinvoicingsummaryexternaldetailSubtotal"`
	// The rebate for the product
	DEzmaxinvoicingsummaryexternaldetailRebate string `json:"dEzmaxinvoicingsummaryexternaldetailRebate"`
	// The total invoiced for the product
	DEzmaxinvoicingsummaryexternaldetailTotal string `json:"dEzmaxinvoicingsummaryexternaldetailTotal"`
	// Whether it's an adjustment
	BEzmaxinvoicingsummaryexternaldetailAdjustment bool `json:"bEzmaxinvoicingsummaryexternaldetailAdjustment"`
	// The help message of the Ezmaxproduct in the language of the requester
	TEzmaxproductHelpX string `json:"tEzmaxproductHelpX"`
}

// NewEzmaxinvoicingsummaryexternaldetailResponseCompound instantiates a new EzmaxinvoicingsummaryexternaldetailResponseCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzmaxinvoicingsummaryexternaldetailResponseCompound(fkiEzmaxproductID int32, sEzmaxproductDescriptionX string, dEzmaxinvoicingsummaryexternaldetailCountreal string, dEzmaxinvoicingsummaryexternaldetailSubtotal string, dEzmaxinvoicingsummaryexternaldetailRebate string, dEzmaxinvoicingsummaryexternaldetailTotal string, bEzmaxinvoicingsummaryexternaldetailAdjustment bool, tEzmaxproductHelpX string) *EzmaxinvoicingsummaryexternaldetailResponseCompound {
	this := EzmaxinvoicingsummaryexternaldetailResponseCompound{}
	this.FkiEzmaxproductID = fkiEzmaxproductID
	this.SEzmaxproductDescriptionX = sEzmaxproductDescriptionX
	this.DEzmaxinvoicingsummaryexternaldetailCountreal = dEzmaxinvoicingsummaryexternaldetailCountreal
	this.DEzmaxinvoicingsummaryexternaldetailSubtotal = dEzmaxinvoicingsummaryexternaldetailSubtotal
	this.DEzmaxinvoicingsummaryexternaldetailRebate = dEzmaxinvoicingsummaryexternaldetailRebate
	this.DEzmaxinvoicingsummaryexternaldetailTotal = dEzmaxinvoicingsummaryexternaldetailTotal
	this.BEzmaxinvoicingsummaryexternaldetailAdjustment = bEzmaxinvoicingsummaryexternaldetailAdjustment
	this.TEzmaxproductHelpX = tEzmaxproductHelpX
	return &this
}

// NewEzmaxinvoicingsummaryexternaldetailResponseCompoundWithDefaults instantiates a new EzmaxinvoicingsummaryexternaldetailResponseCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzmaxinvoicingsummaryexternaldetailResponseCompoundWithDefaults() *EzmaxinvoicingsummaryexternaldetailResponseCompound {
	this := EzmaxinvoicingsummaryexternaldetailResponseCompound{}
	return &this
}

// GetPkiEzmaxinvoicingsummaryexternaldetailID returns the PkiEzmaxinvoicingsummaryexternaldetailID field value if set, zero value otherwise.
func (o *EzmaxinvoicingsummaryexternaldetailResponseCompound) GetPkiEzmaxinvoicingsummaryexternaldetailID() int32 {
	if o == nil || IsNil(o.PkiEzmaxinvoicingsummaryexternaldetailID) {
		var ret int32
		return ret
	}
	return *o.PkiEzmaxinvoicingsummaryexternaldetailID
}

// GetPkiEzmaxinvoicingsummaryexternaldetailIDOk returns a tuple with the PkiEzmaxinvoicingsummaryexternaldetailID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingsummaryexternaldetailResponseCompound) GetPkiEzmaxinvoicingsummaryexternaldetailIDOk() (*int32, bool) {
	if o == nil || IsNil(o.PkiEzmaxinvoicingsummaryexternaldetailID) {
		return nil, false
	}
	return o.PkiEzmaxinvoicingsummaryexternaldetailID, true
}

// HasPkiEzmaxinvoicingsummaryexternaldetailID returns a boolean if a field has been set.
func (o *EzmaxinvoicingsummaryexternaldetailResponseCompound) HasPkiEzmaxinvoicingsummaryexternaldetailID() bool {
	if o != nil && !IsNil(o.PkiEzmaxinvoicingsummaryexternaldetailID) {
		return true
	}

	return false
}

// SetPkiEzmaxinvoicingsummaryexternaldetailID gets a reference to the given int32 and assigns it to the PkiEzmaxinvoicingsummaryexternaldetailID field.
func (o *EzmaxinvoicingsummaryexternaldetailResponseCompound) SetPkiEzmaxinvoicingsummaryexternaldetailID(v int32) {
	o.PkiEzmaxinvoicingsummaryexternaldetailID = &v
}

// GetFkiEzmaxinvoicingsummaryexternalID returns the FkiEzmaxinvoicingsummaryexternalID field value if set, zero value otherwise.
func (o *EzmaxinvoicingsummaryexternaldetailResponseCompound) GetFkiEzmaxinvoicingsummaryexternalID() int32 {
	if o == nil || IsNil(o.FkiEzmaxinvoicingsummaryexternalID) {
		var ret int32
		return ret
	}
	return *o.FkiEzmaxinvoicingsummaryexternalID
}

// GetFkiEzmaxinvoicingsummaryexternalIDOk returns a tuple with the FkiEzmaxinvoicingsummaryexternalID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingsummaryexternaldetailResponseCompound) GetFkiEzmaxinvoicingsummaryexternalIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiEzmaxinvoicingsummaryexternalID) {
		return nil, false
	}
	return o.FkiEzmaxinvoicingsummaryexternalID, true
}

// HasFkiEzmaxinvoicingsummaryexternalID returns a boolean if a field has been set.
func (o *EzmaxinvoicingsummaryexternaldetailResponseCompound) HasFkiEzmaxinvoicingsummaryexternalID() bool {
	if o != nil && !IsNil(o.FkiEzmaxinvoicingsummaryexternalID) {
		return true
	}

	return false
}

// SetFkiEzmaxinvoicingsummaryexternalID gets a reference to the given int32 and assigns it to the FkiEzmaxinvoicingsummaryexternalID field.
func (o *EzmaxinvoicingsummaryexternaldetailResponseCompound) SetFkiEzmaxinvoicingsummaryexternalID(v int32) {
	o.FkiEzmaxinvoicingsummaryexternalID = &v
}

// GetFkiEzmaxproductID returns the FkiEzmaxproductID field value
func (o *EzmaxinvoicingsummaryexternaldetailResponseCompound) GetFkiEzmaxproductID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzmaxproductID
}

// GetFkiEzmaxproductIDOk returns a tuple with the FkiEzmaxproductID field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingsummaryexternaldetailResponseCompound) GetFkiEzmaxproductIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiEzmaxproductID, true
}

// SetFkiEzmaxproductID sets field value
func (o *EzmaxinvoicingsummaryexternaldetailResponseCompound) SetFkiEzmaxproductID(v int32) {
	o.FkiEzmaxproductID = v
}

// GetSEzmaxproductDescriptionX returns the SEzmaxproductDescriptionX field value
func (o *EzmaxinvoicingsummaryexternaldetailResponseCompound) GetSEzmaxproductDescriptionX() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzmaxproductDescriptionX
}

// GetSEzmaxproductDescriptionXOk returns a tuple with the SEzmaxproductDescriptionX field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingsummaryexternaldetailResponseCompound) GetSEzmaxproductDescriptionXOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SEzmaxproductDescriptionX, true
}

// SetSEzmaxproductDescriptionX sets field value
func (o *EzmaxinvoicingsummaryexternaldetailResponseCompound) SetSEzmaxproductDescriptionX(v string) {
	o.SEzmaxproductDescriptionX = v
}

// GetDEzmaxinvoicingsummaryexternaldetailCountreal returns the DEzmaxinvoicingsummaryexternaldetailCountreal field value
func (o *EzmaxinvoicingsummaryexternaldetailResponseCompound) GetDEzmaxinvoicingsummaryexternaldetailCountreal() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DEzmaxinvoicingsummaryexternaldetailCountreal
}

// GetDEzmaxinvoicingsummaryexternaldetailCountrealOk returns a tuple with the DEzmaxinvoicingsummaryexternaldetailCountreal field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingsummaryexternaldetailResponseCompound) GetDEzmaxinvoicingsummaryexternaldetailCountrealOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DEzmaxinvoicingsummaryexternaldetailCountreal, true
}

// SetDEzmaxinvoicingsummaryexternaldetailCountreal sets field value
func (o *EzmaxinvoicingsummaryexternaldetailResponseCompound) SetDEzmaxinvoicingsummaryexternaldetailCountreal(v string) {
	o.DEzmaxinvoicingsummaryexternaldetailCountreal = v
}

// GetDEzmaxinvoicingsummaryexternaldetailSubtotal returns the DEzmaxinvoicingsummaryexternaldetailSubtotal field value
func (o *EzmaxinvoicingsummaryexternaldetailResponseCompound) GetDEzmaxinvoicingsummaryexternaldetailSubtotal() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DEzmaxinvoicingsummaryexternaldetailSubtotal
}

// GetDEzmaxinvoicingsummaryexternaldetailSubtotalOk returns a tuple with the DEzmaxinvoicingsummaryexternaldetailSubtotal field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingsummaryexternaldetailResponseCompound) GetDEzmaxinvoicingsummaryexternaldetailSubtotalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DEzmaxinvoicingsummaryexternaldetailSubtotal, true
}

// SetDEzmaxinvoicingsummaryexternaldetailSubtotal sets field value
func (o *EzmaxinvoicingsummaryexternaldetailResponseCompound) SetDEzmaxinvoicingsummaryexternaldetailSubtotal(v string) {
	o.DEzmaxinvoicingsummaryexternaldetailSubtotal = v
}

// GetDEzmaxinvoicingsummaryexternaldetailRebate returns the DEzmaxinvoicingsummaryexternaldetailRebate field value
func (o *EzmaxinvoicingsummaryexternaldetailResponseCompound) GetDEzmaxinvoicingsummaryexternaldetailRebate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DEzmaxinvoicingsummaryexternaldetailRebate
}

// GetDEzmaxinvoicingsummaryexternaldetailRebateOk returns a tuple with the DEzmaxinvoicingsummaryexternaldetailRebate field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingsummaryexternaldetailResponseCompound) GetDEzmaxinvoicingsummaryexternaldetailRebateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DEzmaxinvoicingsummaryexternaldetailRebate, true
}

// SetDEzmaxinvoicingsummaryexternaldetailRebate sets field value
func (o *EzmaxinvoicingsummaryexternaldetailResponseCompound) SetDEzmaxinvoicingsummaryexternaldetailRebate(v string) {
	o.DEzmaxinvoicingsummaryexternaldetailRebate = v
}

// GetDEzmaxinvoicingsummaryexternaldetailTotal returns the DEzmaxinvoicingsummaryexternaldetailTotal field value
func (o *EzmaxinvoicingsummaryexternaldetailResponseCompound) GetDEzmaxinvoicingsummaryexternaldetailTotal() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DEzmaxinvoicingsummaryexternaldetailTotal
}

// GetDEzmaxinvoicingsummaryexternaldetailTotalOk returns a tuple with the DEzmaxinvoicingsummaryexternaldetailTotal field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingsummaryexternaldetailResponseCompound) GetDEzmaxinvoicingsummaryexternaldetailTotalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DEzmaxinvoicingsummaryexternaldetailTotal, true
}

// SetDEzmaxinvoicingsummaryexternaldetailTotal sets field value
func (o *EzmaxinvoicingsummaryexternaldetailResponseCompound) SetDEzmaxinvoicingsummaryexternaldetailTotal(v string) {
	o.DEzmaxinvoicingsummaryexternaldetailTotal = v
}

// GetBEzmaxinvoicingsummaryexternaldetailAdjustment returns the BEzmaxinvoicingsummaryexternaldetailAdjustment field value
func (o *EzmaxinvoicingsummaryexternaldetailResponseCompound) GetBEzmaxinvoicingsummaryexternaldetailAdjustment() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BEzmaxinvoicingsummaryexternaldetailAdjustment
}

// GetBEzmaxinvoicingsummaryexternaldetailAdjustmentOk returns a tuple with the BEzmaxinvoicingsummaryexternaldetailAdjustment field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingsummaryexternaldetailResponseCompound) GetBEzmaxinvoicingsummaryexternaldetailAdjustmentOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BEzmaxinvoicingsummaryexternaldetailAdjustment, true
}

// SetBEzmaxinvoicingsummaryexternaldetailAdjustment sets field value
func (o *EzmaxinvoicingsummaryexternaldetailResponseCompound) SetBEzmaxinvoicingsummaryexternaldetailAdjustment(v bool) {
	o.BEzmaxinvoicingsummaryexternaldetailAdjustment = v
}

// GetTEzmaxproductHelpX returns the TEzmaxproductHelpX field value
func (o *EzmaxinvoicingsummaryexternaldetailResponseCompound) GetTEzmaxproductHelpX() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TEzmaxproductHelpX
}

// GetTEzmaxproductHelpXOk returns a tuple with the TEzmaxproductHelpX field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingsummaryexternaldetailResponseCompound) GetTEzmaxproductHelpXOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TEzmaxproductHelpX, true
}

// SetTEzmaxproductHelpX sets field value
func (o *EzmaxinvoicingsummaryexternaldetailResponseCompound) SetTEzmaxproductHelpX(v string) {
	o.TEzmaxproductHelpX = v
}

func (o EzmaxinvoicingsummaryexternaldetailResponseCompound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzmaxinvoicingsummaryexternaldetailResponseCompound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PkiEzmaxinvoicingsummaryexternaldetailID) {
		toSerialize["pkiEzmaxinvoicingsummaryexternaldetailID"] = o.PkiEzmaxinvoicingsummaryexternaldetailID
	}
	if !IsNil(o.FkiEzmaxinvoicingsummaryexternalID) {
		toSerialize["fkiEzmaxinvoicingsummaryexternalID"] = o.FkiEzmaxinvoicingsummaryexternalID
	}
	toSerialize["fkiEzmaxproductID"] = o.FkiEzmaxproductID
	toSerialize["sEzmaxproductDescriptionX"] = o.SEzmaxproductDescriptionX
	toSerialize["dEzmaxinvoicingsummaryexternaldetailCountreal"] = o.DEzmaxinvoicingsummaryexternaldetailCountreal
	toSerialize["dEzmaxinvoicingsummaryexternaldetailSubtotal"] = o.DEzmaxinvoicingsummaryexternaldetailSubtotal
	toSerialize["dEzmaxinvoicingsummaryexternaldetailRebate"] = o.DEzmaxinvoicingsummaryexternaldetailRebate
	toSerialize["dEzmaxinvoicingsummaryexternaldetailTotal"] = o.DEzmaxinvoicingsummaryexternaldetailTotal
	toSerialize["bEzmaxinvoicingsummaryexternaldetailAdjustment"] = o.BEzmaxinvoicingsummaryexternaldetailAdjustment
	toSerialize["tEzmaxproductHelpX"] = o.TEzmaxproductHelpX
	return toSerialize, nil
}

type NullableEzmaxinvoicingsummaryexternaldetailResponseCompound struct {
	value *EzmaxinvoicingsummaryexternaldetailResponseCompound
	isSet bool
}

func (v NullableEzmaxinvoicingsummaryexternaldetailResponseCompound) Get() *EzmaxinvoicingsummaryexternaldetailResponseCompound {
	return v.value
}

func (v *NullableEzmaxinvoicingsummaryexternaldetailResponseCompound) Set(val *EzmaxinvoicingsummaryexternaldetailResponseCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableEzmaxinvoicingsummaryexternaldetailResponseCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableEzmaxinvoicingsummaryexternaldetailResponseCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzmaxinvoicingsummaryexternaldetailResponseCompound(val *EzmaxinvoicingsummaryexternaldetailResponseCompound) *NullableEzmaxinvoicingsummaryexternaldetailResponseCompound {
	return &NullableEzmaxinvoicingsummaryexternaldetailResponseCompound{value: val, isSet: true}
}

func (v NullableEzmaxinvoicingsummaryexternaldetailResponseCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzmaxinvoicingsummaryexternaldetailResponseCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

