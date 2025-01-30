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

// checks if the EzsigntsarequirementAutocompleteElementResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigntsarequirementAutocompleteElementResponse{}

// EzsigntsarequirementAutocompleteElementResponse A Ezsigntsarequirement AutocompleteElement Response
type EzsigntsarequirementAutocompleteElementResponse struct {
	// The description of the Ezsigntsarequirement in the language of the requester
	SEzsigntsarequirementDescriptionX string `json:"sEzsigntsarequirementDescriptionX"`
	// The unique ID of the Ezsigntsarequirement.  Determine if a Time Stamping Authority should add a timestamp on each of the signature. Valid values:  |Value|Description| |-|-| |1|No. TSA Timestamping will requested. This will make all signatures a lot faster since no round-trip to the TSA server will be required. Timestamping will be made using eZsign server's time.| |2|Best effort. Timestamping from a Time Stamping Authority will be requested but is not mandatory. In the very improbable case it cannot be completed, the timestamping will be made using eZsign server's time. **Additional fee applies**| |3|Mandatory. Timestamping from a Time Stamping Authority will be requested and is mandatory. In the very improbable case it cannot be completed, the signature will fail and the user will be asked to retry. **Additional fee applies**|
	PkiEzsigntsarequirementID int32 `json:"pkiEzsigntsarequirementID"`
	// Whether the Ezsigntsarequirement is active or not
	BEzsigntsarequirementIsactive bool `json:"bEzsigntsarequirementIsactive"`
	// Indicates if the element is disabled in the context
	BDisabled bool `json:"bDisabled"`
}

type _EzsigntsarequirementAutocompleteElementResponse EzsigntsarequirementAutocompleteElementResponse

// NewEzsigntsarequirementAutocompleteElementResponse instantiates a new EzsigntsarequirementAutocompleteElementResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigntsarequirementAutocompleteElementResponse(sEzsigntsarequirementDescriptionX string, pkiEzsigntsarequirementID int32, bEzsigntsarequirementIsactive bool, bDisabled bool) *EzsigntsarequirementAutocompleteElementResponse {
	this := EzsigntsarequirementAutocompleteElementResponse{}
	this.SEzsigntsarequirementDescriptionX = sEzsigntsarequirementDescriptionX
	this.PkiEzsigntsarequirementID = pkiEzsigntsarequirementID
	this.BEzsigntsarequirementIsactive = bEzsigntsarequirementIsactive
	this.BDisabled = bDisabled
	return &this
}

// NewEzsigntsarequirementAutocompleteElementResponseWithDefaults instantiates a new EzsigntsarequirementAutocompleteElementResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigntsarequirementAutocompleteElementResponseWithDefaults() *EzsigntsarequirementAutocompleteElementResponse {
	this := EzsigntsarequirementAutocompleteElementResponse{}
	return &this
}

// GetSEzsigntsarequirementDescriptionX returns the SEzsigntsarequirementDescriptionX field value
func (o *EzsigntsarequirementAutocompleteElementResponse) GetSEzsigntsarequirementDescriptionX() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzsigntsarequirementDescriptionX
}

// GetSEzsigntsarequirementDescriptionXOk returns a tuple with the SEzsigntsarequirementDescriptionX field value
// and a boolean to check if the value has been set.
func (o *EzsigntsarequirementAutocompleteElementResponse) GetSEzsigntsarequirementDescriptionXOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SEzsigntsarequirementDescriptionX, true
}

// SetSEzsigntsarequirementDescriptionX sets field value
func (o *EzsigntsarequirementAutocompleteElementResponse) SetSEzsigntsarequirementDescriptionX(v string) {
	o.SEzsigntsarequirementDescriptionX = v
}

// GetPkiEzsigntsarequirementID returns the PkiEzsigntsarequirementID field value
func (o *EzsigntsarequirementAutocompleteElementResponse) GetPkiEzsigntsarequirementID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiEzsigntsarequirementID
}

// GetPkiEzsigntsarequirementIDOk returns a tuple with the PkiEzsigntsarequirementID field value
// and a boolean to check if the value has been set.
func (o *EzsigntsarequirementAutocompleteElementResponse) GetPkiEzsigntsarequirementIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiEzsigntsarequirementID, true
}

// SetPkiEzsigntsarequirementID sets field value
func (o *EzsigntsarequirementAutocompleteElementResponse) SetPkiEzsigntsarequirementID(v int32) {
	o.PkiEzsigntsarequirementID = v
}

// GetBEzsigntsarequirementIsactive returns the BEzsigntsarequirementIsactive field value
func (o *EzsigntsarequirementAutocompleteElementResponse) GetBEzsigntsarequirementIsactive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BEzsigntsarequirementIsactive
}

// GetBEzsigntsarequirementIsactiveOk returns a tuple with the BEzsigntsarequirementIsactive field value
// and a boolean to check if the value has been set.
func (o *EzsigntsarequirementAutocompleteElementResponse) GetBEzsigntsarequirementIsactiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BEzsigntsarequirementIsactive, true
}

// SetBEzsigntsarequirementIsactive sets field value
func (o *EzsigntsarequirementAutocompleteElementResponse) SetBEzsigntsarequirementIsactive(v bool) {
	o.BEzsigntsarequirementIsactive = v
}

// GetBDisabled returns the BDisabled field value
func (o *EzsigntsarequirementAutocompleteElementResponse) GetBDisabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BDisabled
}

// GetBDisabledOk returns a tuple with the BDisabled field value
// and a boolean to check if the value has been set.
func (o *EzsigntsarequirementAutocompleteElementResponse) GetBDisabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BDisabled, true
}

// SetBDisabled sets field value
func (o *EzsigntsarequirementAutocompleteElementResponse) SetBDisabled(v bool) {
	o.BDisabled = v
}

func (o EzsigntsarequirementAutocompleteElementResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigntsarequirementAutocompleteElementResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sEzsigntsarequirementDescriptionX"] = o.SEzsigntsarequirementDescriptionX
	toSerialize["pkiEzsigntsarequirementID"] = o.PkiEzsigntsarequirementID
	toSerialize["bEzsigntsarequirementIsactive"] = o.BEzsigntsarequirementIsactive
	toSerialize["bDisabled"] = o.BDisabled
	return toSerialize, nil
}

func (o *EzsigntsarequirementAutocompleteElementResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"sEzsigntsarequirementDescriptionX",
		"pkiEzsigntsarequirementID",
		"bEzsigntsarequirementIsactive",
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

	varEzsigntsarequirementAutocompleteElementResponse := _EzsigntsarequirementAutocompleteElementResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsigntsarequirementAutocompleteElementResponse)

	if err != nil {
		return err
	}

	*o = EzsigntsarequirementAutocompleteElementResponse(varEzsigntsarequirementAutocompleteElementResponse)

	return err
}

type NullableEzsigntsarequirementAutocompleteElementResponse struct {
	value *EzsigntsarequirementAutocompleteElementResponse
	isSet bool
}

func (v NullableEzsigntsarequirementAutocompleteElementResponse) Get() *EzsigntsarequirementAutocompleteElementResponse {
	return v.value
}

func (v *NullableEzsigntsarequirementAutocompleteElementResponse) Set(val *EzsigntsarequirementAutocompleteElementResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigntsarequirementAutocompleteElementResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigntsarequirementAutocompleteElementResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigntsarequirementAutocompleteElementResponse(val *EzsigntsarequirementAutocompleteElementResponse) *NullableEzsigntsarequirementAutocompleteElementResponse {
	return &NullableEzsigntsarequirementAutocompleteElementResponse{value: val, isSet: true}
}

func (v NullableEzsigntsarequirementAutocompleteElementResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigntsarequirementAutocompleteElementResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


