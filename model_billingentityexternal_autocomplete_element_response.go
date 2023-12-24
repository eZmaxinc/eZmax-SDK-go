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

// checks if the BillingentityexternalAutocompleteElementResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BillingentityexternalAutocompleteElementResponse{}

// BillingentityexternalAutocompleteElementResponse A Billingentityexternal AutocompleteElement Response
type BillingentityexternalAutocompleteElementResponse struct {
	// The unique ID of the Billingentityexternal
	PkiBillingentityexternalID int32 `json:"pkiBillingentityexternalID"`
	// The description of the Billingentityexternal
	SBillingentityexternalDescription string `json:"sBillingentityexternalDescription"`
	// Whether the Billingentityexternal is active or not
	BBillingentityexternalIsactive bool `json:"bBillingentityexternalIsactive"`
}

type _BillingentityexternalAutocompleteElementResponse BillingentityexternalAutocompleteElementResponse

// NewBillingentityexternalAutocompleteElementResponse instantiates a new BillingentityexternalAutocompleteElementResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingentityexternalAutocompleteElementResponse(pkiBillingentityexternalID int32, sBillingentityexternalDescription string, bBillingentityexternalIsactive bool) *BillingentityexternalAutocompleteElementResponse {
	this := BillingentityexternalAutocompleteElementResponse{}
	this.PkiBillingentityexternalID = pkiBillingentityexternalID
	this.SBillingentityexternalDescription = sBillingentityexternalDescription
	this.BBillingentityexternalIsactive = bBillingentityexternalIsactive
	return &this
}

// NewBillingentityexternalAutocompleteElementResponseWithDefaults instantiates a new BillingentityexternalAutocompleteElementResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingentityexternalAutocompleteElementResponseWithDefaults() *BillingentityexternalAutocompleteElementResponse {
	this := BillingentityexternalAutocompleteElementResponse{}
	return &this
}

// GetPkiBillingentityexternalID returns the PkiBillingentityexternalID field value
func (o *BillingentityexternalAutocompleteElementResponse) GetPkiBillingentityexternalID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiBillingentityexternalID
}

// GetPkiBillingentityexternalIDOk returns a tuple with the PkiBillingentityexternalID field value
// and a boolean to check if the value has been set.
func (o *BillingentityexternalAutocompleteElementResponse) GetPkiBillingentityexternalIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiBillingentityexternalID, true
}

// SetPkiBillingentityexternalID sets field value
func (o *BillingentityexternalAutocompleteElementResponse) SetPkiBillingentityexternalID(v int32) {
	o.PkiBillingentityexternalID = v
}

// GetSBillingentityexternalDescription returns the SBillingentityexternalDescription field value
func (o *BillingentityexternalAutocompleteElementResponse) GetSBillingentityexternalDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SBillingentityexternalDescription
}

// GetSBillingentityexternalDescriptionOk returns a tuple with the SBillingentityexternalDescription field value
// and a boolean to check if the value has been set.
func (o *BillingentityexternalAutocompleteElementResponse) GetSBillingentityexternalDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SBillingentityexternalDescription, true
}

// SetSBillingentityexternalDescription sets field value
func (o *BillingentityexternalAutocompleteElementResponse) SetSBillingentityexternalDescription(v string) {
	o.SBillingentityexternalDescription = v
}

// GetBBillingentityexternalIsactive returns the BBillingentityexternalIsactive field value
func (o *BillingentityexternalAutocompleteElementResponse) GetBBillingentityexternalIsactive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BBillingentityexternalIsactive
}

// GetBBillingentityexternalIsactiveOk returns a tuple with the BBillingentityexternalIsactive field value
// and a boolean to check if the value has been set.
func (o *BillingentityexternalAutocompleteElementResponse) GetBBillingentityexternalIsactiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BBillingentityexternalIsactive, true
}

// SetBBillingentityexternalIsactive sets field value
func (o *BillingentityexternalAutocompleteElementResponse) SetBBillingentityexternalIsactive(v bool) {
	o.BBillingentityexternalIsactive = v
}

func (o BillingentityexternalAutocompleteElementResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BillingentityexternalAutocompleteElementResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiBillingentityexternalID"] = o.PkiBillingentityexternalID
	toSerialize["sBillingentityexternalDescription"] = o.SBillingentityexternalDescription
	toSerialize["bBillingentityexternalIsactive"] = o.BBillingentityexternalIsactive
	return toSerialize, nil
}

func (o *BillingentityexternalAutocompleteElementResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiBillingentityexternalID",
		"sBillingentityexternalDescription",
		"bBillingentityexternalIsactive",
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

	varBillingentityexternalAutocompleteElementResponse := _BillingentityexternalAutocompleteElementResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBillingentityexternalAutocompleteElementResponse)

	if err != nil {
		return err
	}

	*o = BillingentityexternalAutocompleteElementResponse(varBillingentityexternalAutocompleteElementResponse)

	return err
}

type NullableBillingentityexternalAutocompleteElementResponse struct {
	value *BillingentityexternalAutocompleteElementResponse
	isSet bool
}

func (v NullableBillingentityexternalAutocompleteElementResponse) Get() *BillingentityexternalAutocompleteElementResponse {
	return v.value
}

func (v *NullableBillingentityexternalAutocompleteElementResponse) Set(val *BillingentityexternalAutocompleteElementResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingentityexternalAutocompleteElementResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingentityexternalAutocompleteElementResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingentityexternalAutocompleteElementResponse(val *BillingentityexternalAutocompleteElementResponse) *NullableBillingentityexternalAutocompleteElementResponse {
	return &NullableBillingentityexternalAutocompleteElementResponse{value: val, isSet: true}
}

func (v NullableBillingentityexternalAutocompleteElementResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingentityexternalAutocompleteElementResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


