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

// checks if the EzsigntemplateglobalAutocompleteElementResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigntemplateglobalAutocompleteElementResponse{}

// EzsigntemplateglobalAutocompleteElementResponse A Ezsigntemplate AutocompleteElement Response
type EzsigntemplateglobalAutocompleteElementResponse struct {
	// The unique ID of the Ezsigntemplateglobal
	PkiEzsigntemplateglobalID int32 `json:"pkiEzsigntemplateglobalID"`
	// The description of the Ezsigntemplate
	SEzsigntemplateglobalDescription string `json:"sEzsigntemplateglobalDescription"`
	// Whether the Ezsigntemplate is active or not
	BEzsigntemplateglobalIsactive bool `json:"bEzsigntemplateglobalIsactive"`
}

type _EzsigntemplateglobalAutocompleteElementResponse EzsigntemplateglobalAutocompleteElementResponse

// NewEzsigntemplateglobalAutocompleteElementResponse instantiates a new EzsigntemplateglobalAutocompleteElementResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigntemplateglobalAutocompleteElementResponse(pkiEzsigntemplateglobalID int32, sEzsigntemplateglobalDescription string, bEzsigntemplateglobalIsactive bool) *EzsigntemplateglobalAutocompleteElementResponse {
	this := EzsigntemplateglobalAutocompleteElementResponse{}
	this.PkiEzsigntemplateglobalID = pkiEzsigntemplateglobalID
	this.SEzsigntemplateglobalDescription = sEzsigntemplateglobalDescription
	this.BEzsigntemplateglobalIsactive = bEzsigntemplateglobalIsactive
	return &this
}

// NewEzsigntemplateglobalAutocompleteElementResponseWithDefaults instantiates a new EzsigntemplateglobalAutocompleteElementResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigntemplateglobalAutocompleteElementResponseWithDefaults() *EzsigntemplateglobalAutocompleteElementResponse {
	this := EzsigntemplateglobalAutocompleteElementResponse{}
	return &this
}

// GetPkiEzsigntemplateglobalID returns the PkiEzsigntemplateglobalID field value
func (o *EzsigntemplateglobalAutocompleteElementResponse) GetPkiEzsigntemplateglobalID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiEzsigntemplateglobalID
}

// GetPkiEzsigntemplateglobalIDOk returns a tuple with the PkiEzsigntemplateglobalID field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateglobalAutocompleteElementResponse) GetPkiEzsigntemplateglobalIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiEzsigntemplateglobalID, true
}

// SetPkiEzsigntemplateglobalID sets field value
func (o *EzsigntemplateglobalAutocompleteElementResponse) SetPkiEzsigntemplateglobalID(v int32) {
	o.PkiEzsigntemplateglobalID = v
}

// GetSEzsigntemplateglobalDescription returns the SEzsigntemplateglobalDescription field value
func (o *EzsigntemplateglobalAutocompleteElementResponse) GetSEzsigntemplateglobalDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzsigntemplateglobalDescription
}

// GetSEzsigntemplateglobalDescriptionOk returns a tuple with the SEzsigntemplateglobalDescription field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateglobalAutocompleteElementResponse) GetSEzsigntemplateglobalDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SEzsigntemplateglobalDescription, true
}

// SetSEzsigntemplateglobalDescription sets field value
func (o *EzsigntemplateglobalAutocompleteElementResponse) SetSEzsigntemplateglobalDescription(v string) {
	o.SEzsigntemplateglobalDescription = v
}

// GetBEzsigntemplateglobalIsactive returns the BEzsigntemplateglobalIsactive field value
func (o *EzsigntemplateglobalAutocompleteElementResponse) GetBEzsigntemplateglobalIsactive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BEzsigntemplateglobalIsactive
}

// GetBEzsigntemplateglobalIsactiveOk returns a tuple with the BEzsigntemplateglobalIsactive field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateglobalAutocompleteElementResponse) GetBEzsigntemplateglobalIsactiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BEzsigntemplateglobalIsactive, true
}

// SetBEzsigntemplateglobalIsactive sets field value
func (o *EzsigntemplateglobalAutocompleteElementResponse) SetBEzsigntemplateglobalIsactive(v bool) {
	o.BEzsigntemplateglobalIsactive = v
}

func (o EzsigntemplateglobalAutocompleteElementResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigntemplateglobalAutocompleteElementResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiEzsigntemplateglobalID"] = o.PkiEzsigntemplateglobalID
	toSerialize["sEzsigntemplateglobalDescription"] = o.SEzsigntemplateglobalDescription
	toSerialize["bEzsigntemplateglobalIsactive"] = o.BEzsigntemplateglobalIsactive
	return toSerialize, nil
}

func (o *EzsigntemplateglobalAutocompleteElementResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiEzsigntemplateglobalID",
		"sEzsigntemplateglobalDescription",
		"bEzsigntemplateglobalIsactive",
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

	varEzsigntemplateglobalAutocompleteElementResponse := _EzsigntemplateglobalAutocompleteElementResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsigntemplateglobalAutocompleteElementResponse)

	if err != nil {
		return err
	}

	*o = EzsigntemplateglobalAutocompleteElementResponse(varEzsigntemplateglobalAutocompleteElementResponse)

	return err
}

type NullableEzsigntemplateglobalAutocompleteElementResponse struct {
	value *EzsigntemplateglobalAutocompleteElementResponse
	isSet bool
}

func (v NullableEzsigntemplateglobalAutocompleteElementResponse) Get() *EzsigntemplateglobalAutocompleteElementResponse {
	return v.value
}

func (v *NullableEzsigntemplateglobalAutocompleteElementResponse) Set(val *EzsigntemplateglobalAutocompleteElementResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigntemplateglobalAutocompleteElementResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigntemplateglobalAutocompleteElementResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigntemplateglobalAutocompleteElementResponse(val *EzsigntemplateglobalAutocompleteElementResponse) *NullableEzsigntemplateglobalAutocompleteElementResponse {
	return &NullableEzsigntemplateglobalAutocompleteElementResponse{value: val, isSet: true}
}

func (v NullableEzsigntemplateglobalAutocompleteElementResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigntemplateglobalAutocompleteElementResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

