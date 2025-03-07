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

// checks if the CompanyAutocompleteElementResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CompanyAutocompleteElementResponse{}

// CompanyAutocompleteElementResponse A Company AutocompleteElement Response
type CompanyAutocompleteElementResponse struct {
	// The unique ID of the Company
	PkiCompanyID int32 `json:"pkiCompanyID"`
	// The Name of the Company in the language of the requester
	SCompanyNameX string `json:"sCompanyNameX"`
	// Whether the Company is active or not
	BCompanyIsactive bool `json:"bCompanyIsactive"`
}

type _CompanyAutocompleteElementResponse CompanyAutocompleteElementResponse

// NewCompanyAutocompleteElementResponse instantiates a new CompanyAutocompleteElementResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompanyAutocompleteElementResponse(pkiCompanyID int32, sCompanyNameX string, bCompanyIsactive bool) *CompanyAutocompleteElementResponse {
	this := CompanyAutocompleteElementResponse{}
	this.PkiCompanyID = pkiCompanyID
	this.SCompanyNameX = sCompanyNameX
	this.BCompanyIsactive = bCompanyIsactive
	return &this
}

// NewCompanyAutocompleteElementResponseWithDefaults instantiates a new CompanyAutocompleteElementResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompanyAutocompleteElementResponseWithDefaults() *CompanyAutocompleteElementResponse {
	this := CompanyAutocompleteElementResponse{}
	return &this
}

// GetPkiCompanyID returns the PkiCompanyID field value
func (o *CompanyAutocompleteElementResponse) GetPkiCompanyID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiCompanyID
}

// GetPkiCompanyIDOk returns a tuple with the PkiCompanyID field value
// and a boolean to check if the value has been set.
func (o *CompanyAutocompleteElementResponse) GetPkiCompanyIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiCompanyID, true
}

// SetPkiCompanyID sets field value
func (o *CompanyAutocompleteElementResponse) SetPkiCompanyID(v int32) {
	o.PkiCompanyID = v
}

// GetSCompanyNameX returns the SCompanyNameX field value
func (o *CompanyAutocompleteElementResponse) GetSCompanyNameX() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SCompanyNameX
}

// GetSCompanyNameXOk returns a tuple with the SCompanyNameX field value
// and a boolean to check if the value has been set.
func (o *CompanyAutocompleteElementResponse) GetSCompanyNameXOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SCompanyNameX, true
}

// SetSCompanyNameX sets field value
func (o *CompanyAutocompleteElementResponse) SetSCompanyNameX(v string) {
	o.SCompanyNameX = v
}

// GetBCompanyIsactive returns the BCompanyIsactive field value
func (o *CompanyAutocompleteElementResponse) GetBCompanyIsactive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BCompanyIsactive
}

// GetBCompanyIsactiveOk returns a tuple with the BCompanyIsactive field value
// and a boolean to check if the value has been set.
func (o *CompanyAutocompleteElementResponse) GetBCompanyIsactiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BCompanyIsactive, true
}

// SetBCompanyIsactive sets field value
func (o *CompanyAutocompleteElementResponse) SetBCompanyIsactive(v bool) {
	o.BCompanyIsactive = v
}

func (o CompanyAutocompleteElementResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CompanyAutocompleteElementResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiCompanyID"] = o.PkiCompanyID
	toSerialize["sCompanyNameX"] = o.SCompanyNameX
	toSerialize["bCompanyIsactive"] = o.BCompanyIsactive
	return toSerialize, nil
}

func (o *CompanyAutocompleteElementResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiCompanyID",
		"sCompanyNameX",
		"bCompanyIsactive",
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

	varCompanyAutocompleteElementResponse := _CompanyAutocompleteElementResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCompanyAutocompleteElementResponse)

	if err != nil {
		return err
	}

	*o = CompanyAutocompleteElementResponse(varCompanyAutocompleteElementResponse)

	return err
}

type NullableCompanyAutocompleteElementResponse struct {
	value *CompanyAutocompleteElementResponse
	isSet bool
}

func (v NullableCompanyAutocompleteElementResponse) Get() *CompanyAutocompleteElementResponse {
	return v.value
}

func (v *NullableCompanyAutocompleteElementResponse) Set(val *CompanyAutocompleteElementResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCompanyAutocompleteElementResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCompanyAutocompleteElementResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompanyAutocompleteElementResponse(val *CompanyAutocompleteElementResponse) *NullableCompanyAutocompleteElementResponse {
	return &NullableCompanyAutocompleteElementResponse{value: val, isSet: true}
}

func (v NullableCompanyAutocompleteElementResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompanyAutocompleteElementResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


