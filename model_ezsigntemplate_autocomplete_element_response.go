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
	"bytes"
	"fmt"
)

// checks if the EzsigntemplateAutocompleteElementResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigntemplateAutocompleteElementResponse{}

// EzsigntemplateAutocompleteElementResponse A Ezsigntemplate AutocompleteElement Response
type EzsigntemplateAutocompleteElementResponse struct {
	EEzsignfoldertypePrivacylevel FieldEEzsignfoldertypePrivacylevel `json:"eEzsignfoldertypePrivacylevel"`
	// The description of the Ezsigntemplate
	SEzsigntemplateDescription string `json:"sEzsigntemplateDescription"`
	// The unique ID of the Ezsigntemplate
	PkiEzsigntemplateID int32 `json:"pkiEzsigntemplateID"`
	// Whether the Ezsigntemplate is active or not
	BEzsigntemplateIsactive bool `json:"bEzsigntemplateIsactive"`
}

type _EzsigntemplateAutocompleteElementResponse EzsigntemplateAutocompleteElementResponse

// NewEzsigntemplateAutocompleteElementResponse instantiates a new EzsigntemplateAutocompleteElementResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigntemplateAutocompleteElementResponse(eEzsignfoldertypePrivacylevel FieldEEzsignfoldertypePrivacylevel, sEzsigntemplateDescription string, pkiEzsigntemplateID int32, bEzsigntemplateIsactive bool) *EzsigntemplateAutocompleteElementResponse {
	this := EzsigntemplateAutocompleteElementResponse{}
	this.EEzsignfoldertypePrivacylevel = eEzsignfoldertypePrivacylevel
	this.SEzsigntemplateDescription = sEzsigntemplateDescription
	this.PkiEzsigntemplateID = pkiEzsigntemplateID
	this.BEzsigntemplateIsactive = bEzsigntemplateIsactive
	return &this
}

// NewEzsigntemplateAutocompleteElementResponseWithDefaults instantiates a new EzsigntemplateAutocompleteElementResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigntemplateAutocompleteElementResponseWithDefaults() *EzsigntemplateAutocompleteElementResponse {
	this := EzsigntemplateAutocompleteElementResponse{}
	return &this
}

// GetEEzsignfoldertypePrivacylevel returns the EEzsignfoldertypePrivacylevel field value
func (o *EzsigntemplateAutocompleteElementResponse) GetEEzsignfoldertypePrivacylevel() FieldEEzsignfoldertypePrivacylevel {
	if o == nil {
		var ret FieldEEzsignfoldertypePrivacylevel
		return ret
	}

	return o.EEzsignfoldertypePrivacylevel
}

// GetEEzsignfoldertypePrivacylevelOk returns a tuple with the EEzsignfoldertypePrivacylevel field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateAutocompleteElementResponse) GetEEzsignfoldertypePrivacylevelOk() (*FieldEEzsignfoldertypePrivacylevel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EEzsignfoldertypePrivacylevel, true
}

// SetEEzsignfoldertypePrivacylevel sets field value
func (o *EzsigntemplateAutocompleteElementResponse) SetEEzsignfoldertypePrivacylevel(v FieldEEzsignfoldertypePrivacylevel) {
	o.EEzsignfoldertypePrivacylevel = v
}

// GetSEzsigntemplateDescription returns the SEzsigntemplateDescription field value
func (o *EzsigntemplateAutocompleteElementResponse) GetSEzsigntemplateDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzsigntemplateDescription
}

// GetSEzsigntemplateDescriptionOk returns a tuple with the SEzsigntemplateDescription field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateAutocompleteElementResponse) GetSEzsigntemplateDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SEzsigntemplateDescription, true
}

// SetSEzsigntemplateDescription sets field value
func (o *EzsigntemplateAutocompleteElementResponse) SetSEzsigntemplateDescription(v string) {
	o.SEzsigntemplateDescription = v
}

// GetPkiEzsigntemplateID returns the PkiEzsigntemplateID field value
func (o *EzsigntemplateAutocompleteElementResponse) GetPkiEzsigntemplateID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiEzsigntemplateID
}

// GetPkiEzsigntemplateIDOk returns a tuple with the PkiEzsigntemplateID field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateAutocompleteElementResponse) GetPkiEzsigntemplateIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiEzsigntemplateID, true
}

// SetPkiEzsigntemplateID sets field value
func (o *EzsigntemplateAutocompleteElementResponse) SetPkiEzsigntemplateID(v int32) {
	o.PkiEzsigntemplateID = v
}

// GetBEzsigntemplateIsactive returns the BEzsigntemplateIsactive field value
func (o *EzsigntemplateAutocompleteElementResponse) GetBEzsigntemplateIsactive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BEzsigntemplateIsactive
}

// GetBEzsigntemplateIsactiveOk returns a tuple with the BEzsigntemplateIsactive field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateAutocompleteElementResponse) GetBEzsigntemplateIsactiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BEzsigntemplateIsactive, true
}

// SetBEzsigntemplateIsactive sets field value
func (o *EzsigntemplateAutocompleteElementResponse) SetBEzsigntemplateIsactive(v bool) {
	o.BEzsigntemplateIsactive = v
}

func (o EzsigntemplateAutocompleteElementResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigntemplateAutocompleteElementResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["eEzsignfoldertypePrivacylevel"] = o.EEzsignfoldertypePrivacylevel
	toSerialize["sEzsigntemplateDescription"] = o.SEzsigntemplateDescription
	toSerialize["pkiEzsigntemplateID"] = o.PkiEzsigntemplateID
	toSerialize["bEzsigntemplateIsactive"] = o.BEzsigntemplateIsactive
	return toSerialize, nil
}

func (o *EzsigntemplateAutocompleteElementResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"eEzsignfoldertypePrivacylevel",
		"sEzsigntemplateDescription",
		"pkiEzsigntemplateID",
		"bEzsigntemplateIsactive",
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

	varEzsigntemplateAutocompleteElementResponse := _EzsigntemplateAutocompleteElementResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsigntemplateAutocompleteElementResponse)

	if err != nil {
		return err
	}

	*o = EzsigntemplateAutocompleteElementResponse(varEzsigntemplateAutocompleteElementResponse)

	return err
}

type NullableEzsigntemplateAutocompleteElementResponse struct {
	value *EzsigntemplateAutocompleteElementResponse
	isSet bool
}

func (v NullableEzsigntemplateAutocompleteElementResponse) Get() *EzsigntemplateAutocompleteElementResponse {
	return v.value
}

func (v *NullableEzsigntemplateAutocompleteElementResponse) Set(val *EzsigntemplateAutocompleteElementResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigntemplateAutocompleteElementResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigntemplateAutocompleteElementResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigntemplateAutocompleteElementResponse(val *EzsigntemplateAutocompleteElementResponse) *NullableEzsigntemplateAutocompleteElementResponse {
	return &NullableEzsigntemplateAutocompleteElementResponse{value: val, isSet: true}
}

func (v NullableEzsigntemplateAutocompleteElementResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigntemplateAutocompleteElementResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


