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

// checks if the EzsigntemplatepackageAutocompleteElementResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigntemplatepackageAutocompleteElementResponse{}

// EzsigntemplatepackageAutocompleteElementResponse A Ezsigntemplatepackage AutocompleteElement Response
type EzsigntemplatepackageAutocompleteElementResponse struct {
	EEzsignfoldertypePrivacylevel FieldEEzsignfoldertypePrivacylevel `json:"eEzsignfoldertypePrivacylevel"`
	// The description of the Ezsigntemplatepackage
	SEzsigntemplatepackageDescription string `json:"sEzsigntemplatepackageDescription"`
	// The unique ID of the Ezsigntemplatepackage
	PkiEzsigntemplatepackageID int32 `json:"pkiEzsigntemplatepackageID"`
	// Whether the Ezsigntemplatepackage is active or not
	BEzsigntemplatepackageIsactive bool `json:"bEzsigntemplatepackageIsactive"`
	// Indicates if the element is disabled in the context
	BDisabled bool `json:"bDisabled"`
}

type _EzsigntemplatepackageAutocompleteElementResponse EzsigntemplatepackageAutocompleteElementResponse

// NewEzsigntemplatepackageAutocompleteElementResponse instantiates a new EzsigntemplatepackageAutocompleteElementResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigntemplatepackageAutocompleteElementResponse(eEzsignfoldertypePrivacylevel FieldEEzsignfoldertypePrivacylevel, sEzsigntemplatepackageDescription string, pkiEzsigntemplatepackageID int32, bEzsigntemplatepackageIsactive bool, bDisabled bool) *EzsigntemplatepackageAutocompleteElementResponse {
	this := EzsigntemplatepackageAutocompleteElementResponse{}
	this.EEzsignfoldertypePrivacylevel = eEzsignfoldertypePrivacylevel
	this.SEzsigntemplatepackageDescription = sEzsigntemplatepackageDescription
	this.PkiEzsigntemplatepackageID = pkiEzsigntemplatepackageID
	this.BEzsigntemplatepackageIsactive = bEzsigntemplatepackageIsactive
	this.BDisabled = bDisabled
	return &this
}

// NewEzsigntemplatepackageAutocompleteElementResponseWithDefaults instantiates a new EzsigntemplatepackageAutocompleteElementResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigntemplatepackageAutocompleteElementResponseWithDefaults() *EzsigntemplatepackageAutocompleteElementResponse {
	this := EzsigntemplatepackageAutocompleteElementResponse{}
	return &this
}

// GetEEzsignfoldertypePrivacylevel returns the EEzsignfoldertypePrivacylevel field value
func (o *EzsigntemplatepackageAutocompleteElementResponse) GetEEzsignfoldertypePrivacylevel() FieldEEzsignfoldertypePrivacylevel {
	if o == nil {
		var ret FieldEEzsignfoldertypePrivacylevel
		return ret
	}

	return o.EEzsignfoldertypePrivacylevel
}

// GetEEzsignfoldertypePrivacylevelOk returns a tuple with the EEzsignfoldertypePrivacylevel field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepackageAutocompleteElementResponse) GetEEzsignfoldertypePrivacylevelOk() (*FieldEEzsignfoldertypePrivacylevel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EEzsignfoldertypePrivacylevel, true
}

// SetEEzsignfoldertypePrivacylevel sets field value
func (o *EzsigntemplatepackageAutocompleteElementResponse) SetEEzsignfoldertypePrivacylevel(v FieldEEzsignfoldertypePrivacylevel) {
	o.EEzsignfoldertypePrivacylevel = v
}

// GetSEzsigntemplatepackageDescription returns the SEzsigntemplatepackageDescription field value
func (o *EzsigntemplatepackageAutocompleteElementResponse) GetSEzsigntemplatepackageDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzsigntemplatepackageDescription
}

// GetSEzsigntemplatepackageDescriptionOk returns a tuple with the SEzsigntemplatepackageDescription field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepackageAutocompleteElementResponse) GetSEzsigntemplatepackageDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SEzsigntemplatepackageDescription, true
}

// SetSEzsigntemplatepackageDescription sets field value
func (o *EzsigntemplatepackageAutocompleteElementResponse) SetSEzsigntemplatepackageDescription(v string) {
	o.SEzsigntemplatepackageDescription = v
}

// GetPkiEzsigntemplatepackageID returns the PkiEzsigntemplatepackageID field value
func (o *EzsigntemplatepackageAutocompleteElementResponse) GetPkiEzsigntemplatepackageID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiEzsigntemplatepackageID
}

// GetPkiEzsigntemplatepackageIDOk returns a tuple with the PkiEzsigntemplatepackageID field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepackageAutocompleteElementResponse) GetPkiEzsigntemplatepackageIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiEzsigntemplatepackageID, true
}

// SetPkiEzsigntemplatepackageID sets field value
func (o *EzsigntemplatepackageAutocompleteElementResponse) SetPkiEzsigntemplatepackageID(v int32) {
	o.PkiEzsigntemplatepackageID = v
}

// GetBEzsigntemplatepackageIsactive returns the BEzsigntemplatepackageIsactive field value
func (o *EzsigntemplatepackageAutocompleteElementResponse) GetBEzsigntemplatepackageIsactive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BEzsigntemplatepackageIsactive
}

// GetBEzsigntemplatepackageIsactiveOk returns a tuple with the BEzsigntemplatepackageIsactive field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepackageAutocompleteElementResponse) GetBEzsigntemplatepackageIsactiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BEzsigntemplatepackageIsactive, true
}

// SetBEzsigntemplatepackageIsactive sets field value
func (o *EzsigntemplatepackageAutocompleteElementResponse) SetBEzsigntemplatepackageIsactive(v bool) {
	o.BEzsigntemplatepackageIsactive = v
}

// GetBDisabled returns the BDisabled field value
func (o *EzsigntemplatepackageAutocompleteElementResponse) GetBDisabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BDisabled
}

// GetBDisabledOk returns a tuple with the BDisabled field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepackageAutocompleteElementResponse) GetBDisabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BDisabled, true
}

// SetBDisabled sets field value
func (o *EzsigntemplatepackageAutocompleteElementResponse) SetBDisabled(v bool) {
	o.BDisabled = v
}

func (o EzsigntemplatepackageAutocompleteElementResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigntemplatepackageAutocompleteElementResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["eEzsignfoldertypePrivacylevel"] = o.EEzsignfoldertypePrivacylevel
	toSerialize["sEzsigntemplatepackageDescription"] = o.SEzsigntemplatepackageDescription
	toSerialize["pkiEzsigntemplatepackageID"] = o.PkiEzsigntemplatepackageID
	toSerialize["bEzsigntemplatepackageIsactive"] = o.BEzsigntemplatepackageIsactive
	toSerialize["bDisabled"] = o.BDisabled
	return toSerialize, nil
}

func (o *EzsigntemplatepackageAutocompleteElementResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"eEzsignfoldertypePrivacylevel",
		"sEzsigntemplatepackageDescription",
		"pkiEzsigntemplatepackageID",
		"bEzsigntemplatepackageIsactive",
		"bDisabled",
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

	varEzsigntemplatepackageAutocompleteElementResponse := _EzsigntemplatepackageAutocompleteElementResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsigntemplatepackageAutocompleteElementResponse)

	if err != nil {
		return err
	}

	*o = EzsigntemplatepackageAutocompleteElementResponse(varEzsigntemplatepackageAutocompleteElementResponse)

	return err
}

type NullableEzsigntemplatepackageAutocompleteElementResponse struct {
	value *EzsigntemplatepackageAutocompleteElementResponse
	isSet bool
}

func (v NullableEzsigntemplatepackageAutocompleteElementResponse) Get() *EzsigntemplatepackageAutocompleteElementResponse {
	return v.value
}

func (v *NullableEzsigntemplatepackageAutocompleteElementResponse) Set(val *EzsigntemplatepackageAutocompleteElementResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigntemplatepackageAutocompleteElementResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigntemplatepackageAutocompleteElementResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigntemplatepackageAutocompleteElementResponse(val *EzsigntemplatepackageAutocompleteElementResponse) *NullableEzsigntemplatepackageAutocompleteElementResponse {
	return &NullableEzsigntemplatepackageAutocompleteElementResponse{value: val, isSet: true}
}

func (v NullableEzsigntemplatepackageAutocompleteElementResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigntemplatepackageAutocompleteElementResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


