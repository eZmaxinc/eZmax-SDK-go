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

// checks if the BrandingAutocompleteElementResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BrandingAutocompleteElementResponse{}

// BrandingAutocompleteElementResponse Branding AutocompleteElement Response
type BrandingAutocompleteElementResponse struct {
	// The Description of the Branding in the language of the requester
	SBrandingDescriptionX string `json:"sBrandingDescriptionX"`
	// The unique ID of the Branding
	PkiBrandingID int32 `json:"pkiBrandingID"`
	// Whether the Branding is active or not
	BBrandingIsactive bool `json:"bBrandingIsactive"`
}

// NewBrandingAutocompleteElementResponse instantiates a new BrandingAutocompleteElementResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBrandingAutocompleteElementResponse(sBrandingDescriptionX string, pkiBrandingID int32, bBrandingIsactive bool) *BrandingAutocompleteElementResponse {
	this := BrandingAutocompleteElementResponse{}
	this.SBrandingDescriptionX = sBrandingDescriptionX
	this.PkiBrandingID = pkiBrandingID
	this.BBrandingIsactive = bBrandingIsactive
	return &this
}

// NewBrandingAutocompleteElementResponseWithDefaults instantiates a new BrandingAutocompleteElementResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBrandingAutocompleteElementResponseWithDefaults() *BrandingAutocompleteElementResponse {
	this := BrandingAutocompleteElementResponse{}
	return &this
}

// GetSBrandingDescriptionX returns the SBrandingDescriptionX field value
func (o *BrandingAutocompleteElementResponse) GetSBrandingDescriptionX() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SBrandingDescriptionX
}

// GetSBrandingDescriptionXOk returns a tuple with the SBrandingDescriptionX field value
// and a boolean to check if the value has been set.
func (o *BrandingAutocompleteElementResponse) GetSBrandingDescriptionXOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SBrandingDescriptionX, true
}

// SetSBrandingDescriptionX sets field value
func (o *BrandingAutocompleteElementResponse) SetSBrandingDescriptionX(v string) {
	o.SBrandingDescriptionX = v
}

// GetPkiBrandingID returns the PkiBrandingID field value
func (o *BrandingAutocompleteElementResponse) GetPkiBrandingID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiBrandingID
}

// GetPkiBrandingIDOk returns a tuple with the PkiBrandingID field value
// and a boolean to check if the value has been set.
func (o *BrandingAutocompleteElementResponse) GetPkiBrandingIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiBrandingID, true
}

// SetPkiBrandingID sets field value
func (o *BrandingAutocompleteElementResponse) SetPkiBrandingID(v int32) {
	o.PkiBrandingID = v
}

// GetBBrandingIsactive returns the BBrandingIsactive field value
func (o *BrandingAutocompleteElementResponse) GetBBrandingIsactive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BBrandingIsactive
}

// GetBBrandingIsactiveOk returns a tuple with the BBrandingIsactive field value
// and a boolean to check if the value has been set.
func (o *BrandingAutocompleteElementResponse) GetBBrandingIsactiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BBrandingIsactive, true
}

// SetBBrandingIsactive sets field value
func (o *BrandingAutocompleteElementResponse) SetBBrandingIsactive(v bool) {
	o.BBrandingIsactive = v
}

func (o BrandingAutocompleteElementResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BrandingAutocompleteElementResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sBrandingDescriptionX"] = o.SBrandingDescriptionX
	toSerialize["pkiBrandingID"] = o.PkiBrandingID
	toSerialize["bBrandingIsactive"] = o.BBrandingIsactive
	return toSerialize, nil
}

type NullableBrandingAutocompleteElementResponse struct {
	value *BrandingAutocompleteElementResponse
	isSet bool
}

func (v NullableBrandingAutocompleteElementResponse) Get() *BrandingAutocompleteElementResponse {
	return v.value
}

func (v *NullableBrandingAutocompleteElementResponse) Set(val *BrandingAutocompleteElementResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBrandingAutocompleteElementResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBrandingAutocompleteElementResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBrandingAutocompleteElementResponse(val *BrandingAutocompleteElementResponse) *NullableBrandingAutocompleteElementResponse {
	return &NullableBrandingAutocompleteElementResponse{value: val, isSet: true}
}

func (v NullableBrandingAutocompleteElementResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBrandingAutocompleteElementResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


