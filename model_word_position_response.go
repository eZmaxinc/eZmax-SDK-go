/*
eZmax API Definition

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.0.48
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
)

// WordPositionResponse A Word Position Object
type WordPositionResponse struct {
	// The page where the word occurence was found
	IPage *int32 `json:"iPage,omitempty"`
	// The X coordinate (Horizontal) where the Word occurence was found.  Coordinate is calculated at 100dpi (dot per inch).
	IX *int32 `json:"iX,omitempty"`
	// The Y coordinate (Vertical) where the Word occurence was found.  Coordinate is calculated at 100dpi (dot per inch).
	IY *int32 `json:"iY,omitempty"`
}

// NewWordPositionResponse instantiates a new WordPositionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWordPositionResponse() *WordPositionResponse {
	this := WordPositionResponse{}
	return &this
}

// NewWordPositionResponseWithDefaults instantiates a new WordPositionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWordPositionResponseWithDefaults() *WordPositionResponse {
	this := WordPositionResponse{}
	return &this
}

// GetIPage returns the IPage field value if set, zero value otherwise.
func (o *WordPositionResponse) GetIPage() int32 {
	if o == nil || o.IPage == nil {
		var ret int32
		return ret
	}
	return *o.IPage
}

// GetIPageOk returns a tuple with the IPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WordPositionResponse) GetIPageOk() (*int32, bool) {
	if o == nil || o.IPage == nil {
		return nil, false
	}
	return o.IPage, true
}

// HasIPage returns a boolean if a field has been set.
func (o *WordPositionResponse) HasIPage() bool {
	if o != nil && o.IPage != nil {
		return true
	}

	return false
}

// SetIPage gets a reference to the given int32 and assigns it to the IPage field.
func (o *WordPositionResponse) SetIPage(v int32) {
	o.IPage = &v
}

// GetIX returns the IX field value if set, zero value otherwise.
func (o *WordPositionResponse) GetIX() int32 {
	if o == nil || o.IX == nil {
		var ret int32
		return ret
	}
	return *o.IX
}

// GetIXOk returns a tuple with the IX field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WordPositionResponse) GetIXOk() (*int32, bool) {
	if o == nil || o.IX == nil {
		return nil, false
	}
	return o.IX, true
}

// HasIX returns a boolean if a field has been set.
func (o *WordPositionResponse) HasIX() bool {
	if o != nil && o.IX != nil {
		return true
	}

	return false
}

// SetIX gets a reference to the given int32 and assigns it to the IX field.
func (o *WordPositionResponse) SetIX(v int32) {
	o.IX = &v
}

// GetIY returns the IY field value if set, zero value otherwise.
func (o *WordPositionResponse) GetIY() int32 {
	if o == nil || o.IY == nil {
		var ret int32
		return ret
	}
	return *o.IY
}

// GetIYOk returns a tuple with the IY field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WordPositionResponse) GetIYOk() (*int32, bool) {
	if o == nil || o.IY == nil {
		return nil, false
	}
	return o.IY, true
}

// HasIY returns a boolean if a field has been set.
func (o *WordPositionResponse) HasIY() bool {
	if o != nil && o.IY != nil {
		return true
	}

	return false
}

// SetIY gets a reference to the given int32 and assigns it to the IY field.
func (o *WordPositionResponse) SetIY(v int32) {
	o.IY = &v
}

func (o WordPositionResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IPage != nil {
		toSerialize["iPage"] = o.IPage
	}
	if o.IX != nil {
		toSerialize["iX"] = o.IX
	}
	if o.IY != nil {
		toSerialize["iY"] = o.IY
	}
	return json.Marshal(toSerialize)
}

type NullableWordPositionResponse struct {
	value *WordPositionResponse
	isSet bool
}

func (v NullableWordPositionResponse) Get() *WordPositionResponse {
	return v.value
}

func (v *NullableWordPositionResponse) Set(val *WordPositionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWordPositionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWordPositionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWordPositionResponse(val *WordPositionResponse) *NullableWordPositionResponse {
	return &NullableWordPositionResponse{value: val, isSet: true}
}

func (v NullableWordPositionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWordPositionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

