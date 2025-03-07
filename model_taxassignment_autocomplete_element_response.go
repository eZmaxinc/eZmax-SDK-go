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

// checks if the TaxassignmentAutocompleteElementResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaxassignmentAutocompleteElementResponse{}

// TaxassignmentAutocompleteElementResponse A Taxassignment AutocompleteElement Response
type TaxassignmentAutocompleteElementResponse struct {
	// The description of the Taxassignment  in the language of the requester
	STaxassignmentDescriptionX string `json:"sTaxassignmentDescriptionX"`
	// The unique ID of the Taxassignment.  Valid values:  |Value|Description| |-|-| |1|No tax| |2|GST| |3|HST (ON)| |4|HST (NB)| |5|HST (NS)| |6|HST (NL)| |7|HST (PE)| |8|GST + QST (QC)| |9|GST + QST (QC) Non-Recoverable| |10|GST + PST (BC)| |11|GST + PST (SK)| |12|GST + RST (MB)| |13|GST + PST (BC) Non-Recoverable| |14|GST + PST (SK) Non-Recoverable| |15|GST + RST (MB) Non-Recoverable|
	PkiTaxassignmentID int32 `json:"pkiTaxassignmentID"`
	// Whether the Taxassignment is active or not
	BTaxassignmentIsactive bool `json:"bTaxassignmentIsactive"`
}

type _TaxassignmentAutocompleteElementResponse TaxassignmentAutocompleteElementResponse

// NewTaxassignmentAutocompleteElementResponse instantiates a new TaxassignmentAutocompleteElementResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaxassignmentAutocompleteElementResponse(sTaxassignmentDescriptionX string, pkiTaxassignmentID int32, bTaxassignmentIsactive bool) *TaxassignmentAutocompleteElementResponse {
	this := TaxassignmentAutocompleteElementResponse{}
	this.STaxassignmentDescriptionX = sTaxassignmentDescriptionX
	this.PkiTaxassignmentID = pkiTaxassignmentID
	this.BTaxassignmentIsactive = bTaxassignmentIsactive
	return &this
}

// NewTaxassignmentAutocompleteElementResponseWithDefaults instantiates a new TaxassignmentAutocompleteElementResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaxassignmentAutocompleteElementResponseWithDefaults() *TaxassignmentAutocompleteElementResponse {
	this := TaxassignmentAutocompleteElementResponse{}
	return &this
}

// GetSTaxassignmentDescriptionX returns the STaxassignmentDescriptionX field value
func (o *TaxassignmentAutocompleteElementResponse) GetSTaxassignmentDescriptionX() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.STaxassignmentDescriptionX
}

// GetSTaxassignmentDescriptionXOk returns a tuple with the STaxassignmentDescriptionX field value
// and a boolean to check if the value has been set.
func (o *TaxassignmentAutocompleteElementResponse) GetSTaxassignmentDescriptionXOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.STaxassignmentDescriptionX, true
}

// SetSTaxassignmentDescriptionX sets field value
func (o *TaxassignmentAutocompleteElementResponse) SetSTaxassignmentDescriptionX(v string) {
	o.STaxassignmentDescriptionX = v
}

// GetPkiTaxassignmentID returns the PkiTaxassignmentID field value
func (o *TaxassignmentAutocompleteElementResponse) GetPkiTaxassignmentID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiTaxassignmentID
}

// GetPkiTaxassignmentIDOk returns a tuple with the PkiTaxassignmentID field value
// and a boolean to check if the value has been set.
func (o *TaxassignmentAutocompleteElementResponse) GetPkiTaxassignmentIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiTaxassignmentID, true
}

// SetPkiTaxassignmentID sets field value
func (o *TaxassignmentAutocompleteElementResponse) SetPkiTaxassignmentID(v int32) {
	o.PkiTaxassignmentID = v
}

// GetBTaxassignmentIsactive returns the BTaxassignmentIsactive field value
func (o *TaxassignmentAutocompleteElementResponse) GetBTaxassignmentIsactive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BTaxassignmentIsactive
}

// GetBTaxassignmentIsactiveOk returns a tuple with the BTaxassignmentIsactive field value
// and a boolean to check if the value has been set.
func (o *TaxassignmentAutocompleteElementResponse) GetBTaxassignmentIsactiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BTaxassignmentIsactive, true
}

// SetBTaxassignmentIsactive sets field value
func (o *TaxassignmentAutocompleteElementResponse) SetBTaxassignmentIsactive(v bool) {
	o.BTaxassignmentIsactive = v
}

func (o TaxassignmentAutocompleteElementResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaxassignmentAutocompleteElementResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sTaxassignmentDescriptionX"] = o.STaxassignmentDescriptionX
	toSerialize["pkiTaxassignmentID"] = o.PkiTaxassignmentID
	toSerialize["bTaxassignmentIsactive"] = o.BTaxassignmentIsactive
	return toSerialize, nil
}

func (o *TaxassignmentAutocompleteElementResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"sTaxassignmentDescriptionX",
		"pkiTaxassignmentID",
		"bTaxassignmentIsactive",
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

	varTaxassignmentAutocompleteElementResponse := _TaxassignmentAutocompleteElementResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTaxassignmentAutocompleteElementResponse)

	if err != nil {
		return err
	}

	*o = TaxassignmentAutocompleteElementResponse(varTaxassignmentAutocompleteElementResponse)

	return err
}

type NullableTaxassignmentAutocompleteElementResponse struct {
	value *TaxassignmentAutocompleteElementResponse
	isSet bool
}

func (v NullableTaxassignmentAutocompleteElementResponse) Get() *TaxassignmentAutocompleteElementResponse {
	return v.value
}

func (v *NullableTaxassignmentAutocompleteElementResponse) Set(val *TaxassignmentAutocompleteElementResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxassignmentAutocompleteElementResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxassignmentAutocompleteElementResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxassignmentAutocompleteElementResponse(val *TaxassignmentAutocompleteElementResponse) *NullableTaxassignmentAutocompleteElementResponse {
	return &NullableTaxassignmentAutocompleteElementResponse{value: val, isSet: true}
}

func (v NullableTaxassignmentAutocompleteElementResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaxassignmentAutocompleteElementResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


