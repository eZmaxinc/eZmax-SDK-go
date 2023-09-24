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

// checks if the CustomEzmaxpricingResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomEzmaxpricingResponse{}

// CustomEzmaxpricingResponse A Custom Ezmaxpricing Object
type CustomEzmaxpricingResponse struct {
	// The unique ID of the Ezmaxpricing
	PkiEzmaxpricingID int32 `json:"pkiEzmaxpricingID"`
	// The rebate offered when eZsign is taken for all agents
	DEzmaxpricingRebateezsignallagents string `json:"dEzmaxpricingRebateezsignallagents"`
	// The start date of the Ezmaxpricing
	DtEzmaxpricingStart string `json:"dtEzmaxpricingStart"`
	// The end date of the Ezmaxpricing
	DtEzmaxpricingEnd *string `json:"dtEzmaxpricingEnd,omitempty"`
}

// NewCustomEzmaxpricingResponse instantiates a new CustomEzmaxpricingResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomEzmaxpricingResponse(pkiEzmaxpricingID int32, dEzmaxpricingRebateezsignallagents string, dtEzmaxpricingStart string) *CustomEzmaxpricingResponse {
	this := CustomEzmaxpricingResponse{}
	this.PkiEzmaxpricingID = pkiEzmaxpricingID
	this.DEzmaxpricingRebateezsignallagents = dEzmaxpricingRebateezsignallagents
	this.DtEzmaxpricingStart = dtEzmaxpricingStart
	return &this
}

// NewCustomEzmaxpricingResponseWithDefaults instantiates a new CustomEzmaxpricingResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomEzmaxpricingResponseWithDefaults() *CustomEzmaxpricingResponse {
	this := CustomEzmaxpricingResponse{}
	return &this
}

// GetPkiEzmaxpricingID returns the PkiEzmaxpricingID field value
func (o *CustomEzmaxpricingResponse) GetPkiEzmaxpricingID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiEzmaxpricingID
}

// GetPkiEzmaxpricingIDOk returns a tuple with the PkiEzmaxpricingID field value
// and a boolean to check if the value has been set.
func (o *CustomEzmaxpricingResponse) GetPkiEzmaxpricingIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiEzmaxpricingID, true
}

// SetPkiEzmaxpricingID sets field value
func (o *CustomEzmaxpricingResponse) SetPkiEzmaxpricingID(v int32) {
	o.PkiEzmaxpricingID = v
}

// GetDEzmaxpricingRebateezsignallagents returns the DEzmaxpricingRebateezsignallagents field value
func (o *CustomEzmaxpricingResponse) GetDEzmaxpricingRebateezsignallagents() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DEzmaxpricingRebateezsignallagents
}

// GetDEzmaxpricingRebateezsignallagentsOk returns a tuple with the DEzmaxpricingRebateezsignallagents field value
// and a boolean to check if the value has been set.
func (o *CustomEzmaxpricingResponse) GetDEzmaxpricingRebateezsignallagentsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DEzmaxpricingRebateezsignallagents, true
}

// SetDEzmaxpricingRebateezsignallagents sets field value
func (o *CustomEzmaxpricingResponse) SetDEzmaxpricingRebateezsignallagents(v string) {
	o.DEzmaxpricingRebateezsignallagents = v
}

// GetDtEzmaxpricingStart returns the DtEzmaxpricingStart field value
func (o *CustomEzmaxpricingResponse) GetDtEzmaxpricingStart() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DtEzmaxpricingStart
}

// GetDtEzmaxpricingStartOk returns a tuple with the DtEzmaxpricingStart field value
// and a boolean to check if the value has been set.
func (o *CustomEzmaxpricingResponse) GetDtEzmaxpricingStartOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DtEzmaxpricingStart, true
}

// SetDtEzmaxpricingStart sets field value
func (o *CustomEzmaxpricingResponse) SetDtEzmaxpricingStart(v string) {
	o.DtEzmaxpricingStart = v
}

// GetDtEzmaxpricingEnd returns the DtEzmaxpricingEnd field value if set, zero value otherwise.
func (o *CustomEzmaxpricingResponse) GetDtEzmaxpricingEnd() string {
	if o == nil || IsNil(o.DtEzmaxpricingEnd) {
		var ret string
		return ret
	}
	return *o.DtEzmaxpricingEnd
}

// GetDtEzmaxpricingEndOk returns a tuple with the DtEzmaxpricingEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomEzmaxpricingResponse) GetDtEzmaxpricingEndOk() (*string, bool) {
	if o == nil || IsNil(o.DtEzmaxpricingEnd) {
		return nil, false
	}
	return o.DtEzmaxpricingEnd, true
}

// HasDtEzmaxpricingEnd returns a boolean if a field has been set.
func (o *CustomEzmaxpricingResponse) HasDtEzmaxpricingEnd() bool {
	if o != nil && !IsNil(o.DtEzmaxpricingEnd) {
		return true
	}

	return false
}

// SetDtEzmaxpricingEnd gets a reference to the given string and assigns it to the DtEzmaxpricingEnd field.
func (o *CustomEzmaxpricingResponse) SetDtEzmaxpricingEnd(v string) {
	o.DtEzmaxpricingEnd = &v
}

func (o CustomEzmaxpricingResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomEzmaxpricingResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiEzmaxpricingID"] = o.PkiEzmaxpricingID
	toSerialize["dEzmaxpricingRebateezsignallagents"] = o.DEzmaxpricingRebateezsignallagents
	toSerialize["dtEzmaxpricingStart"] = o.DtEzmaxpricingStart
	if !IsNil(o.DtEzmaxpricingEnd) {
		toSerialize["dtEzmaxpricingEnd"] = o.DtEzmaxpricingEnd
	}
	return toSerialize, nil
}

type NullableCustomEzmaxpricingResponse struct {
	value *CustomEzmaxpricingResponse
	isSet bool
}

func (v NullableCustomEzmaxpricingResponse) Get() *CustomEzmaxpricingResponse {
	return v.value
}

func (v *NullableCustomEzmaxpricingResponse) Set(val *CustomEzmaxpricingResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomEzmaxpricingResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomEzmaxpricingResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomEzmaxpricingResponse(val *CustomEzmaxpricingResponse) *NullableCustomEzmaxpricingResponse {
	return &NullableCustomEzmaxpricingResponse{value: val, isSet: true}
}

func (v NullableCustomEzmaxpricingResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomEzmaxpricingResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

