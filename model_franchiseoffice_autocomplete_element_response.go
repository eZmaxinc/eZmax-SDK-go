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

// checks if the FranchiseofficeAutocompleteElementResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FranchiseofficeAutocompleteElementResponse{}

// FranchiseofficeAutocompleteElementResponse A Franchiseoffice AutocompleteElement Response
type FranchiseofficeAutocompleteElementResponse struct {
	// The description of the Franchiseoffice in the language of the requester
	SFranchiseofficeDescription string `json:"sFranchiseofficeDescription"`
	// The unique ID of the Franchisereoffice
	PkiFranchiseofficeID int32 `json:"pkiFranchiseofficeID"`
	// Whether the Franchiseoffice is active or not
	BFranchiseofficeIsactive bool `json:"bFranchiseofficeIsactive"`
}

type _FranchiseofficeAutocompleteElementResponse FranchiseofficeAutocompleteElementResponse

// NewFranchiseofficeAutocompleteElementResponse instantiates a new FranchiseofficeAutocompleteElementResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFranchiseofficeAutocompleteElementResponse(sFranchiseofficeDescription string, pkiFranchiseofficeID int32, bFranchiseofficeIsactive bool) *FranchiseofficeAutocompleteElementResponse {
	this := FranchiseofficeAutocompleteElementResponse{}
	this.SFranchiseofficeDescription = sFranchiseofficeDescription
	this.PkiFranchiseofficeID = pkiFranchiseofficeID
	this.BFranchiseofficeIsactive = bFranchiseofficeIsactive
	return &this
}

// NewFranchiseofficeAutocompleteElementResponseWithDefaults instantiates a new FranchiseofficeAutocompleteElementResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFranchiseofficeAutocompleteElementResponseWithDefaults() *FranchiseofficeAutocompleteElementResponse {
	this := FranchiseofficeAutocompleteElementResponse{}
	return &this
}

// GetSFranchiseofficeDescription returns the SFranchiseofficeDescription field value
func (o *FranchiseofficeAutocompleteElementResponse) GetSFranchiseofficeDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SFranchiseofficeDescription
}

// GetSFranchiseofficeDescriptionOk returns a tuple with the SFranchiseofficeDescription field value
// and a boolean to check if the value has been set.
func (o *FranchiseofficeAutocompleteElementResponse) GetSFranchiseofficeDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SFranchiseofficeDescription, true
}

// SetSFranchiseofficeDescription sets field value
func (o *FranchiseofficeAutocompleteElementResponse) SetSFranchiseofficeDescription(v string) {
	o.SFranchiseofficeDescription = v
}

// GetPkiFranchiseofficeID returns the PkiFranchiseofficeID field value
func (o *FranchiseofficeAutocompleteElementResponse) GetPkiFranchiseofficeID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiFranchiseofficeID
}

// GetPkiFranchiseofficeIDOk returns a tuple with the PkiFranchiseofficeID field value
// and a boolean to check if the value has been set.
func (o *FranchiseofficeAutocompleteElementResponse) GetPkiFranchiseofficeIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiFranchiseofficeID, true
}

// SetPkiFranchiseofficeID sets field value
func (o *FranchiseofficeAutocompleteElementResponse) SetPkiFranchiseofficeID(v int32) {
	o.PkiFranchiseofficeID = v
}

// GetBFranchiseofficeIsactive returns the BFranchiseofficeIsactive field value
func (o *FranchiseofficeAutocompleteElementResponse) GetBFranchiseofficeIsactive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BFranchiseofficeIsactive
}

// GetBFranchiseofficeIsactiveOk returns a tuple with the BFranchiseofficeIsactive field value
// and a boolean to check if the value has been set.
func (o *FranchiseofficeAutocompleteElementResponse) GetBFranchiseofficeIsactiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BFranchiseofficeIsactive, true
}

// SetBFranchiseofficeIsactive sets field value
func (o *FranchiseofficeAutocompleteElementResponse) SetBFranchiseofficeIsactive(v bool) {
	o.BFranchiseofficeIsactive = v
}

func (o FranchiseofficeAutocompleteElementResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FranchiseofficeAutocompleteElementResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sFranchiseofficeDescription"] = o.SFranchiseofficeDescription
	toSerialize["pkiFranchiseofficeID"] = o.PkiFranchiseofficeID
	toSerialize["bFranchiseofficeIsactive"] = o.BFranchiseofficeIsactive
	return toSerialize, nil
}

func (o *FranchiseofficeAutocompleteElementResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"sFranchiseofficeDescription",
		"pkiFranchiseofficeID",
		"bFranchiseofficeIsactive",
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

	varFranchiseofficeAutocompleteElementResponse := _FranchiseofficeAutocompleteElementResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFranchiseofficeAutocompleteElementResponse)

	if err != nil {
		return err
	}

	*o = FranchiseofficeAutocompleteElementResponse(varFranchiseofficeAutocompleteElementResponse)

	return err
}

type NullableFranchiseofficeAutocompleteElementResponse struct {
	value *FranchiseofficeAutocompleteElementResponse
	isSet bool
}

func (v NullableFranchiseofficeAutocompleteElementResponse) Get() *FranchiseofficeAutocompleteElementResponse {
	return v.value
}

func (v *NullableFranchiseofficeAutocompleteElementResponse) Set(val *FranchiseofficeAutocompleteElementResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFranchiseofficeAutocompleteElementResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFranchiseofficeAutocompleteElementResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFranchiseofficeAutocompleteElementResponse(val *FranchiseofficeAutocompleteElementResponse) *NullableFranchiseofficeAutocompleteElementResponse {
	return &NullableFranchiseofficeAutocompleteElementResponse{value: val, isSet: true}
}

func (v NullableFranchiseofficeAutocompleteElementResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFranchiseofficeAutocompleteElementResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


