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

// checks if the EzsignsigningreasonAutocompleteElementResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignsigningreasonAutocompleteElementResponse{}

// EzsignsigningreasonAutocompleteElementResponse A Ezsignsigningreason AutocompleteElement Response
type EzsignsigningreasonAutocompleteElementResponse struct {
	// The unique ID of the Ezsignsigningreason
	PkiEzsignsigningreasonID int32 `json:"pkiEzsignsigningreasonID"`
	// The description of the Ezsignsigningreason in the language of the requester
	SEzsignsigningreasonDescriptionX string `json:"sEzsignsigningreasonDescriptionX" validate:"regexp=^.{0,50}$"`
	// Whether the ezsignsigningreason is active or not
	BEzsignsigningreasonIsactive bool `json:"bEzsignsigningreasonIsactive"`
}

type _EzsignsigningreasonAutocompleteElementResponse EzsignsigningreasonAutocompleteElementResponse

// NewEzsignsigningreasonAutocompleteElementResponse instantiates a new EzsignsigningreasonAutocompleteElementResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignsigningreasonAutocompleteElementResponse(pkiEzsignsigningreasonID int32, sEzsignsigningreasonDescriptionX string, bEzsignsigningreasonIsactive bool) *EzsignsigningreasonAutocompleteElementResponse {
	this := EzsignsigningreasonAutocompleteElementResponse{}
	this.PkiEzsignsigningreasonID = pkiEzsignsigningreasonID
	this.SEzsignsigningreasonDescriptionX = sEzsignsigningreasonDescriptionX
	this.BEzsignsigningreasonIsactive = bEzsignsigningreasonIsactive
	return &this
}

// NewEzsignsigningreasonAutocompleteElementResponseWithDefaults instantiates a new EzsignsigningreasonAutocompleteElementResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignsigningreasonAutocompleteElementResponseWithDefaults() *EzsignsigningreasonAutocompleteElementResponse {
	this := EzsignsigningreasonAutocompleteElementResponse{}
	return &this
}

// GetPkiEzsignsigningreasonID returns the PkiEzsignsigningreasonID field value
func (o *EzsignsigningreasonAutocompleteElementResponse) GetPkiEzsignsigningreasonID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiEzsignsigningreasonID
}

// GetPkiEzsignsigningreasonIDOk returns a tuple with the PkiEzsignsigningreasonID field value
// and a boolean to check if the value has been set.
func (o *EzsignsigningreasonAutocompleteElementResponse) GetPkiEzsignsigningreasonIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiEzsignsigningreasonID, true
}

// SetPkiEzsignsigningreasonID sets field value
func (o *EzsignsigningreasonAutocompleteElementResponse) SetPkiEzsignsigningreasonID(v int32) {
	o.PkiEzsignsigningreasonID = v
}

// GetSEzsignsigningreasonDescriptionX returns the SEzsignsigningreasonDescriptionX field value
func (o *EzsignsigningreasonAutocompleteElementResponse) GetSEzsignsigningreasonDescriptionX() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzsignsigningreasonDescriptionX
}

// GetSEzsignsigningreasonDescriptionXOk returns a tuple with the SEzsignsigningreasonDescriptionX field value
// and a boolean to check if the value has been set.
func (o *EzsignsigningreasonAutocompleteElementResponse) GetSEzsignsigningreasonDescriptionXOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SEzsignsigningreasonDescriptionX, true
}

// SetSEzsignsigningreasonDescriptionX sets field value
func (o *EzsignsigningreasonAutocompleteElementResponse) SetSEzsignsigningreasonDescriptionX(v string) {
	o.SEzsignsigningreasonDescriptionX = v
}

// GetBEzsignsigningreasonIsactive returns the BEzsignsigningreasonIsactive field value
func (o *EzsignsigningreasonAutocompleteElementResponse) GetBEzsignsigningreasonIsactive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BEzsignsigningreasonIsactive
}

// GetBEzsignsigningreasonIsactiveOk returns a tuple with the BEzsignsigningreasonIsactive field value
// and a boolean to check if the value has been set.
func (o *EzsignsigningreasonAutocompleteElementResponse) GetBEzsignsigningreasonIsactiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BEzsignsigningreasonIsactive, true
}

// SetBEzsignsigningreasonIsactive sets field value
func (o *EzsignsigningreasonAutocompleteElementResponse) SetBEzsignsigningreasonIsactive(v bool) {
	o.BEzsignsigningreasonIsactive = v
}

func (o EzsignsigningreasonAutocompleteElementResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignsigningreasonAutocompleteElementResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiEzsignsigningreasonID"] = o.PkiEzsignsigningreasonID
	toSerialize["sEzsignsigningreasonDescriptionX"] = o.SEzsignsigningreasonDescriptionX
	toSerialize["bEzsignsigningreasonIsactive"] = o.BEzsignsigningreasonIsactive
	return toSerialize, nil
}

func (o *EzsignsigningreasonAutocompleteElementResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiEzsignsigningreasonID",
		"sEzsignsigningreasonDescriptionX",
		"bEzsignsigningreasonIsactive",
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

	varEzsignsigningreasonAutocompleteElementResponse := _EzsignsigningreasonAutocompleteElementResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsignsigningreasonAutocompleteElementResponse)

	if err != nil {
		return err
	}

	*o = EzsignsigningreasonAutocompleteElementResponse(varEzsignsigningreasonAutocompleteElementResponse)

	return err
}

type NullableEzsignsigningreasonAutocompleteElementResponse struct {
	value *EzsignsigningreasonAutocompleteElementResponse
	isSet bool
}

func (v NullableEzsignsigningreasonAutocompleteElementResponse) Get() *EzsignsigningreasonAutocompleteElementResponse {
	return v.value
}

func (v *NullableEzsignsigningreasonAutocompleteElementResponse) Set(val *EzsignsigningreasonAutocompleteElementResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignsigningreasonAutocompleteElementResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignsigningreasonAutocompleteElementResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignsigningreasonAutocompleteElementResponse(val *EzsignsigningreasonAutocompleteElementResponse) *NullableEzsignsigningreasonAutocompleteElementResponse {
	return &NullableEzsignsigningreasonAutocompleteElementResponse{value: val, isSet: true}
}

func (v NullableEzsignsigningreasonAutocompleteElementResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignsigningreasonAutocompleteElementResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


