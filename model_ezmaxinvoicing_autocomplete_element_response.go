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

// checks if the EzmaxinvoicingAutocompleteElementResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzmaxinvoicingAutocompleteElementResponse{}

// EzmaxinvoicingAutocompleteElementResponse A Ezmaxinvoicing AutocompleteElement Response
type EzmaxinvoicingAutocompleteElementResponse struct {
	// The YYYYMM period of the Ezmaxinvoicing
	YyyymmEzmaxinvoicing string `json:"yyyymmEzmaxinvoicing"`
	// The unique ID of the Ezmaxinvoicing
	PkiEzmaxinvoicingID int32 `json:"pkiEzmaxinvoicingID"`
	// Whether the Ezmaxinvoicing is active or not
	BEzmaxinvoicingIsactive bool `json:"bEzmaxinvoicingIsactive"`
}

type _EzmaxinvoicingAutocompleteElementResponse EzmaxinvoicingAutocompleteElementResponse

// NewEzmaxinvoicingAutocompleteElementResponse instantiates a new EzmaxinvoicingAutocompleteElementResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzmaxinvoicingAutocompleteElementResponse(yyyymmEzmaxinvoicing string, pkiEzmaxinvoicingID int32, bEzmaxinvoicingIsactive bool) *EzmaxinvoicingAutocompleteElementResponse {
	this := EzmaxinvoicingAutocompleteElementResponse{}
	this.YyyymmEzmaxinvoicing = yyyymmEzmaxinvoicing
	this.PkiEzmaxinvoicingID = pkiEzmaxinvoicingID
	this.BEzmaxinvoicingIsactive = bEzmaxinvoicingIsactive
	return &this
}

// NewEzmaxinvoicingAutocompleteElementResponseWithDefaults instantiates a new EzmaxinvoicingAutocompleteElementResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzmaxinvoicingAutocompleteElementResponseWithDefaults() *EzmaxinvoicingAutocompleteElementResponse {
	this := EzmaxinvoicingAutocompleteElementResponse{}
	return &this
}

// GetYyyymmEzmaxinvoicing returns the YyyymmEzmaxinvoicing field value
func (o *EzmaxinvoicingAutocompleteElementResponse) GetYyyymmEzmaxinvoicing() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.YyyymmEzmaxinvoicing
}

// GetYyyymmEzmaxinvoicingOk returns a tuple with the YyyymmEzmaxinvoicing field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingAutocompleteElementResponse) GetYyyymmEzmaxinvoicingOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.YyyymmEzmaxinvoicing, true
}

// SetYyyymmEzmaxinvoicing sets field value
func (o *EzmaxinvoicingAutocompleteElementResponse) SetYyyymmEzmaxinvoicing(v string) {
	o.YyyymmEzmaxinvoicing = v
}

// GetPkiEzmaxinvoicingID returns the PkiEzmaxinvoicingID field value
func (o *EzmaxinvoicingAutocompleteElementResponse) GetPkiEzmaxinvoicingID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiEzmaxinvoicingID
}

// GetPkiEzmaxinvoicingIDOk returns a tuple with the PkiEzmaxinvoicingID field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingAutocompleteElementResponse) GetPkiEzmaxinvoicingIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiEzmaxinvoicingID, true
}

// SetPkiEzmaxinvoicingID sets field value
func (o *EzmaxinvoicingAutocompleteElementResponse) SetPkiEzmaxinvoicingID(v int32) {
	o.PkiEzmaxinvoicingID = v
}

// GetBEzmaxinvoicingIsactive returns the BEzmaxinvoicingIsactive field value
func (o *EzmaxinvoicingAutocompleteElementResponse) GetBEzmaxinvoicingIsactive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BEzmaxinvoicingIsactive
}

// GetBEzmaxinvoicingIsactiveOk returns a tuple with the BEzmaxinvoicingIsactive field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingAutocompleteElementResponse) GetBEzmaxinvoicingIsactiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BEzmaxinvoicingIsactive, true
}

// SetBEzmaxinvoicingIsactive sets field value
func (o *EzmaxinvoicingAutocompleteElementResponse) SetBEzmaxinvoicingIsactive(v bool) {
	o.BEzmaxinvoicingIsactive = v
}

func (o EzmaxinvoicingAutocompleteElementResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzmaxinvoicingAutocompleteElementResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["yyyymmEzmaxinvoicing"] = o.YyyymmEzmaxinvoicing
	toSerialize["pkiEzmaxinvoicingID"] = o.PkiEzmaxinvoicingID
	toSerialize["bEzmaxinvoicingIsactive"] = o.BEzmaxinvoicingIsactive
	return toSerialize, nil
}

func (o *EzmaxinvoicingAutocompleteElementResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"yyyymmEzmaxinvoicing",
		"pkiEzmaxinvoicingID",
		"bEzmaxinvoicingIsactive",
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

	varEzmaxinvoicingAutocompleteElementResponse := _EzmaxinvoicingAutocompleteElementResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzmaxinvoicingAutocompleteElementResponse)

	if err != nil {
		return err
	}

	*o = EzmaxinvoicingAutocompleteElementResponse(varEzmaxinvoicingAutocompleteElementResponse)

	return err
}

type NullableEzmaxinvoicingAutocompleteElementResponse struct {
	value *EzmaxinvoicingAutocompleteElementResponse
	isSet bool
}

func (v NullableEzmaxinvoicingAutocompleteElementResponse) Get() *EzmaxinvoicingAutocompleteElementResponse {
	return v.value
}

func (v *NullableEzmaxinvoicingAutocompleteElementResponse) Set(val *EzmaxinvoicingAutocompleteElementResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEzmaxinvoicingAutocompleteElementResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEzmaxinvoicingAutocompleteElementResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzmaxinvoicingAutocompleteElementResponse(val *EzmaxinvoicingAutocompleteElementResponse) *NullableEzmaxinvoicingAutocompleteElementResponse {
	return &NullableEzmaxinvoicingAutocompleteElementResponse{value: val, isSet: true}
}

func (v NullableEzmaxinvoicingAutocompleteElementResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzmaxinvoicingAutocompleteElementResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


