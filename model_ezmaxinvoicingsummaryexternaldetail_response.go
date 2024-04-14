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

// checks if the EzmaxinvoicingsummaryexternaldetailResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzmaxinvoicingsummaryexternaldetailResponse{}

// EzmaxinvoicingsummaryexternaldetailResponse A Ezmaxinvoicingsummaryexternaldetail Object
type EzmaxinvoicingsummaryexternaldetailResponse struct {
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

type _EzmaxinvoicingsummaryexternaldetailResponse EzmaxinvoicingsummaryexternaldetailResponse

// NewEzmaxinvoicingsummaryexternaldetailResponse instantiates a new EzmaxinvoicingsummaryexternaldetailResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzmaxinvoicingsummaryexternaldetailResponse(fkiEzmaxproductID int32, sEzmaxproductDescriptionX string, dEzmaxinvoicingsummaryexternaldetailCountreal string, dEzmaxinvoicingsummaryexternaldetailSubtotal string, dEzmaxinvoicingsummaryexternaldetailRebate string, dEzmaxinvoicingsummaryexternaldetailTotal string, bEzmaxinvoicingsummaryexternaldetailAdjustment bool, tEzmaxproductHelpX string) *EzmaxinvoicingsummaryexternaldetailResponse {
	this := EzmaxinvoicingsummaryexternaldetailResponse{}
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

// NewEzmaxinvoicingsummaryexternaldetailResponseWithDefaults instantiates a new EzmaxinvoicingsummaryexternaldetailResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzmaxinvoicingsummaryexternaldetailResponseWithDefaults() *EzmaxinvoicingsummaryexternaldetailResponse {
	this := EzmaxinvoicingsummaryexternaldetailResponse{}
	return &this
}

// GetPkiEzmaxinvoicingsummaryexternaldetailID returns the PkiEzmaxinvoicingsummaryexternaldetailID field value if set, zero value otherwise.
func (o *EzmaxinvoicingsummaryexternaldetailResponse) GetPkiEzmaxinvoicingsummaryexternaldetailID() int32 {
	if o == nil || IsNil(o.PkiEzmaxinvoicingsummaryexternaldetailID) {
		var ret int32
		return ret
	}
	return *o.PkiEzmaxinvoicingsummaryexternaldetailID
}

// GetPkiEzmaxinvoicingsummaryexternaldetailIDOk returns a tuple with the PkiEzmaxinvoicingsummaryexternaldetailID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingsummaryexternaldetailResponse) GetPkiEzmaxinvoicingsummaryexternaldetailIDOk() (*int32, bool) {
	if o == nil || IsNil(o.PkiEzmaxinvoicingsummaryexternaldetailID) {
		return nil, false
	}
	return o.PkiEzmaxinvoicingsummaryexternaldetailID, true
}

// HasPkiEzmaxinvoicingsummaryexternaldetailID returns a boolean if a field has been set.
func (o *EzmaxinvoicingsummaryexternaldetailResponse) HasPkiEzmaxinvoicingsummaryexternaldetailID() bool {
	if o != nil && !IsNil(o.PkiEzmaxinvoicingsummaryexternaldetailID) {
		return true
	}

	return false
}

// SetPkiEzmaxinvoicingsummaryexternaldetailID gets a reference to the given int32 and assigns it to the PkiEzmaxinvoicingsummaryexternaldetailID field.
func (o *EzmaxinvoicingsummaryexternaldetailResponse) SetPkiEzmaxinvoicingsummaryexternaldetailID(v int32) {
	o.PkiEzmaxinvoicingsummaryexternaldetailID = &v
}

// GetFkiEzmaxinvoicingsummaryexternalID returns the FkiEzmaxinvoicingsummaryexternalID field value if set, zero value otherwise.
func (o *EzmaxinvoicingsummaryexternaldetailResponse) GetFkiEzmaxinvoicingsummaryexternalID() int32 {
	if o == nil || IsNil(o.FkiEzmaxinvoicingsummaryexternalID) {
		var ret int32
		return ret
	}
	return *o.FkiEzmaxinvoicingsummaryexternalID
}

// GetFkiEzmaxinvoicingsummaryexternalIDOk returns a tuple with the FkiEzmaxinvoicingsummaryexternalID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingsummaryexternaldetailResponse) GetFkiEzmaxinvoicingsummaryexternalIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiEzmaxinvoicingsummaryexternalID) {
		return nil, false
	}
	return o.FkiEzmaxinvoicingsummaryexternalID, true
}

// HasFkiEzmaxinvoicingsummaryexternalID returns a boolean if a field has been set.
func (o *EzmaxinvoicingsummaryexternaldetailResponse) HasFkiEzmaxinvoicingsummaryexternalID() bool {
	if o != nil && !IsNil(o.FkiEzmaxinvoicingsummaryexternalID) {
		return true
	}

	return false
}

// SetFkiEzmaxinvoicingsummaryexternalID gets a reference to the given int32 and assigns it to the FkiEzmaxinvoicingsummaryexternalID field.
func (o *EzmaxinvoicingsummaryexternaldetailResponse) SetFkiEzmaxinvoicingsummaryexternalID(v int32) {
	o.FkiEzmaxinvoicingsummaryexternalID = &v
}

// GetFkiEzmaxproductID returns the FkiEzmaxproductID field value
func (o *EzmaxinvoicingsummaryexternaldetailResponse) GetFkiEzmaxproductID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzmaxproductID
}

// GetFkiEzmaxproductIDOk returns a tuple with the FkiEzmaxproductID field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingsummaryexternaldetailResponse) GetFkiEzmaxproductIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiEzmaxproductID, true
}

// SetFkiEzmaxproductID sets field value
func (o *EzmaxinvoicingsummaryexternaldetailResponse) SetFkiEzmaxproductID(v int32) {
	o.FkiEzmaxproductID = v
}

// GetSEzmaxproductDescriptionX returns the SEzmaxproductDescriptionX field value
func (o *EzmaxinvoicingsummaryexternaldetailResponse) GetSEzmaxproductDescriptionX() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzmaxproductDescriptionX
}

// GetSEzmaxproductDescriptionXOk returns a tuple with the SEzmaxproductDescriptionX field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingsummaryexternaldetailResponse) GetSEzmaxproductDescriptionXOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SEzmaxproductDescriptionX, true
}

// SetSEzmaxproductDescriptionX sets field value
func (o *EzmaxinvoicingsummaryexternaldetailResponse) SetSEzmaxproductDescriptionX(v string) {
	o.SEzmaxproductDescriptionX = v
}

// GetDEzmaxinvoicingsummaryexternaldetailCountreal returns the DEzmaxinvoicingsummaryexternaldetailCountreal field value
func (o *EzmaxinvoicingsummaryexternaldetailResponse) GetDEzmaxinvoicingsummaryexternaldetailCountreal() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DEzmaxinvoicingsummaryexternaldetailCountreal
}

// GetDEzmaxinvoicingsummaryexternaldetailCountrealOk returns a tuple with the DEzmaxinvoicingsummaryexternaldetailCountreal field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingsummaryexternaldetailResponse) GetDEzmaxinvoicingsummaryexternaldetailCountrealOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DEzmaxinvoicingsummaryexternaldetailCountreal, true
}

// SetDEzmaxinvoicingsummaryexternaldetailCountreal sets field value
func (o *EzmaxinvoicingsummaryexternaldetailResponse) SetDEzmaxinvoicingsummaryexternaldetailCountreal(v string) {
	o.DEzmaxinvoicingsummaryexternaldetailCountreal = v
}

// GetDEzmaxinvoicingsummaryexternaldetailSubtotal returns the DEzmaxinvoicingsummaryexternaldetailSubtotal field value
func (o *EzmaxinvoicingsummaryexternaldetailResponse) GetDEzmaxinvoicingsummaryexternaldetailSubtotal() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DEzmaxinvoicingsummaryexternaldetailSubtotal
}

// GetDEzmaxinvoicingsummaryexternaldetailSubtotalOk returns a tuple with the DEzmaxinvoicingsummaryexternaldetailSubtotal field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingsummaryexternaldetailResponse) GetDEzmaxinvoicingsummaryexternaldetailSubtotalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DEzmaxinvoicingsummaryexternaldetailSubtotal, true
}

// SetDEzmaxinvoicingsummaryexternaldetailSubtotal sets field value
func (o *EzmaxinvoicingsummaryexternaldetailResponse) SetDEzmaxinvoicingsummaryexternaldetailSubtotal(v string) {
	o.DEzmaxinvoicingsummaryexternaldetailSubtotal = v
}

// GetDEzmaxinvoicingsummaryexternaldetailRebate returns the DEzmaxinvoicingsummaryexternaldetailRebate field value
func (o *EzmaxinvoicingsummaryexternaldetailResponse) GetDEzmaxinvoicingsummaryexternaldetailRebate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DEzmaxinvoicingsummaryexternaldetailRebate
}

// GetDEzmaxinvoicingsummaryexternaldetailRebateOk returns a tuple with the DEzmaxinvoicingsummaryexternaldetailRebate field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingsummaryexternaldetailResponse) GetDEzmaxinvoicingsummaryexternaldetailRebateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DEzmaxinvoicingsummaryexternaldetailRebate, true
}

// SetDEzmaxinvoicingsummaryexternaldetailRebate sets field value
func (o *EzmaxinvoicingsummaryexternaldetailResponse) SetDEzmaxinvoicingsummaryexternaldetailRebate(v string) {
	o.DEzmaxinvoicingsummaryexternaldetailRebate = v
}

// GetDEzmaxinvoicingsummaryexternaldetailTotal returns the DEzmaxinvoicingsummaryexternaldetailTotal field value
func (o *EzmaxinvoicingsummaryexternaldetailResponse) GetDEzmaxinvoicingsummaryexternaldetailTotal() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DEzmaxinvoicingsummaryexternaldetailTotal
}

// GetDEzmaxinvoicingsummaryexternaldetailTotalOk returns a tuple with the DEzmaxinvoicingsummaryexternaldetailTotal field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingsummaryexternaldetailResponse) GetDEzmaxinvoicingsummaryexternaldetailTotalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DEzmaxinvoicingsummaryexternaldetailTotal, true
}

// SetDEzmaxinvoicingsummaryexternaldetailTotal sets field value
func (o *EzmaxinvoicingsummaryexternaldetailResponse) SetDEzmaxinvoicingsummaryexternaldetailTotal(v string) {
	o.DEzmaxinvoicingsummaryexternaldetailTotal = v
}

// GetBEzmaxinvoicingsummaryexternaldetailAdjustment returns the BEzmaxinvoicingsummaryexternaldetailAdjustment field value
func (o *EzmaxinvoicingsummaryexternaldetailResponse) GetBEzmaxinvoicingsummaryexternaldetailAdjustment() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BEzmaxinvoicingsummaryexternaldetailAdjustment
}

// GetBEzmaxinvoicingsummaryexternaldetailAdjustmentOk returns a tuple with the BEzmaxinvoicingsummaryexternaldetailAdjustment field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingsummaryexternaldetailResponse) GetBEzmaxinvoicingsummaryexternaldetailAdjustmentOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BEzmaxinvoicingsummaryexternaldetailAdjustment, true
}

// SetBEzmaxinvoicingsummaryexternaldetailAdjustment sets field value
func (o *EzmaxinvoicingsummaryexternaldetailResponse) SetBEzmaxinvoicingsummaryexternaldetailAdjustment(v bool) {
	o.BEzmaxinvoicingsummaryexternaldetailAdjustment = v
}

// GetTEzmaxproductHelpX returns the TEzmaxproductHelpX field value
func (o *EzmaxinvoicingsummaryexternaldetailResponse) GetTEzmaxproductHelpX() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TEzmaxproductHelpX
}

// GetTEzmaxproductHelpXOk returns a tuple with the TEzmaxproductHelpX field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingsummaryexternaldetailResponse) GetTEzmaxproductHelpXOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TEzmaxproductHelpX, true
}

// SetTEzmaxproductHelpX sets field value
func (o *EzmaxinvoicingsummaryexternaldetailResponse) SetTEzmaxproductHelpX(v string) {
	o.TEzmaxproductHelpX = v
}

func (o EzmaxinvoicingsummaryexternaldetailResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzmaxinvoicingsummaryexternaldetailResponse) ToMap() (map[string]interface{}, error) {
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

func (o *EzmaxinvoicingsummaryexternaldetailResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fkiEzmaxproductID",
		"sEzmaxproductDescriptionX",
		"dEzmaxinvoicingsummaryexternaldetailCountreal",
		"dEzmaxinvoicingsummaryexternaldetailSubtotal",
		"dEzmaxinvoicingsummaryexternaldetailRebate",
		"dEzmaxinvoicingsummaryexternaldetailTotal",
		"bEzmaxinvoicingsummaryexternaldetailAdjustment",
		"tEzmaxproductHelpX",
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

	varEzmaxinvoicingsummaryexternaldetailResponse := _EzmaxinvoicingsummaryexternaldetailResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzmaxinvoicingsummaryexternaldetailResponse)

	if err != nil {
		return err
	}

	*o = EzmaxinvoicingsummaryexternaldetailResponse(varEzmaxinvoicingsummaryexternaldetailResponse)

	return err
}

type NullableEzmaxinvoicingsummaryexternaldetailResponse struct {
	value *EzmaxinvoicingsummaryexternaldetailResponse
	isSet bool
}

func (v NullableEzmaxinvoicingsummaryexternaldetailResponse) Get() *EzmaxinvoicingsummaryexternaldetailResponse {
	return v.value
}

func (v *NullableEzmaxinvoicingsummaryexternaldetailResponse) Set(val *EzmaxinvoicingsummaryexternaldetailResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEzmaxinvoicingsummaryexternaldetailResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEzmaxinvoicingsummaryexternaldetailResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzmaxinvoicingsummaryexternaldetailResponse(val *EzmaxinvoicingsummaryexternaldetailResponse) *NullableEzmaxinvoicingsummaryexternaldetailResponse {
	return &NullableEzmaxinvoicingsummaryexternaldetailResponse{value: val, isSet: true}
}

func (v NullableEzmaxinvoicingsummaryexternaldetailResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzmaxinvoicingsummaryexternaldetailResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


