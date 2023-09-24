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

// checks if the VariableexpenseAutocompleteElementResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VariableexpenseAutocompleteElementResponse{}

// VariableexpenseAutocompleteElementResponse A Variableexpense AutocompleteElement Response
type VariableexpenseAutocompleteElementResponse struct {
	// The description of the Variableexpense in the language of the requester
	SVariableexpenseDescriptionX string `json:"sVariableexpenseDescriptionX"`
	// The unique ID of the Variableexpense
	PkiVariableexpenseID int32 `json:"pkiVariableexpenseID"`
	// Whether the variableexpense is active or not
	BVariableexpenseIsactive bool `json:"bVariableexpenseIsactive"`
}

// NewVariableexpenseAutocompleteElementResponse instantiates a new VariableexpenseAutocompleteElementResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVariableexpenseAutocompleteElementResponse(sVariableexpenseDescriptionX string, pkiVariableexpenseID int32, bVariableexpenseIsactive bool) *VariableexpenseAutocompleteElementResponse {
	this := VariableexpenseAutocompleteElementResponse{}
	this.SVariableexpenseDescriptionX = sVariableexpenseDescriptionX
	this.PkiVariableexpenseID = pkiVariableexpenseID
	this.BVariableexpenseIsactive = bVariableexpenseIsactive
	return &this
}

// NewVariableexpenseAutocompleteElementResponseWithDefaults instantiates a new VariableexpenseAutocompleteElementResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVariableexpenseAutocompleteElementResponseWithDefaults() *VariableexpenseAutocompleteElementResponse {
	this := VariableexpenseAutocompleteElementResponse{}
	return &this
}

// GetSVariableexpenseDescriptionX returns the SVariableexpenseDescriptionX field value
func (o *VariableexpenseAutocompleteElementResponse) GetSVariableexpenseDescriptionX() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SVariableexpenseDescriptionX
}

// GetSVariableexpenseDescriptionXOk returns a tuple with the SVariableexpenseDescriptionX field value
// and a boolean to check if the value has been set.
func (o *VariableexpenseAutocompleteElementResponse) GetSVariableexpenseDescriptionXOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SVariableexpenseDescriptionX, true
}

// SetSVariableexpenseDescriptionX sets field value
func (o *VariableexpenseAutocompleteElementResponse) SetSVariableexpenseDescriptionX(v string) {
	o.SVariableexpenseDescriptionX = v
}

// GetPkiVariableexpenseID returns the PkiVariableexpenseID field value
func (o *VariableexpenseAutocompleteElementResponse) GetPkiVariableexpenseID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiVariableexpenseID
}

// GetPkiVariableexpenseIDOk returns a tuple with the PkiVariableexpenseID field value
// and a boolean to check if the value has been set.
func (o *VariableexpenseAutocompleteElementResponse) GetPkiVariableexpenseIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiVariableexpenseID, true
}

// SetPkiVariableexpenseID sets field value
func (o *VariableexpenseAutocompleteElementResponse) SetPkiVariableexpenseID(v int32) {
	o.PkiVariableexpenseID = v
}

// GetBVariableexpenseIsactive returns the BVariableexpenseIsactive field value
func (o *VariableexpenseAutocompleteElementResponse) GetBVariableexpenseIsactive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BVariableexpenseIsactive
}

// GetBVariableexpenseIsactiveOk returns a tuple with the BVariableexpenseIsactive field value
// and a boolean to check if the value has been set.
func (o *VariableexpenseAutocompleteElementResponse) GetBVariableexpenseIsactiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BVariableexpenseIsactive, true
}

// SetBVariableexpenseIsactive sets field value
func (o *VariableexpenseAutocompleteElementResponse) SetBVariableexpenseIsactive(v bool) {
	o.BVariableexpenseIsactive = v
}

func (o VariableexpenseAutocompleteElementResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VariableexpenseAutocompleteElementResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sVariableexpenseDescriptionX"] = o.SVariableexpenseDescriptionX
	toSerialize["pkiVariableexpenseID"] = o.PkiVariableexpenseID
	toSerialize["bVariableexpenseIsactive"] = o.BVariableexpenseIsactive
	return toSerialize, nil
}

type NullableVariableexpenseAutocompleteElementResponse struct {
	value *VariableexpenseAutocompleteElementResponse
	isSet bool
}

func (v NullableVariableexpenseAutocompleteElementResponse) Get() *VariableexpenseAutocompleteElementResponse {
	return v.value
}

func (v *NullableVariableexpenseAutocompleteElementResponse) Set(val *VariableexpenseAutocompleteElementResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableVariableexpenseAutocompleteElementResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableVariableexpenseAutocompleteElementResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVariableexpenseAutocompleteElementResponse(val *VariableexpenseAutocompleteElementResponse) *NullableVariableexpenseAutocompleteElementResponse {
	return &NullableVariableexpenseAutocompleteElementResponse{value: val, isSet: true}
}

func (v NullableVariableexpenseAutocompleteElementResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVariableexpenseAutocompleteElementResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


