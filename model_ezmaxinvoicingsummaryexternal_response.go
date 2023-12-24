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

// checks if the EzmaxinvoicingsummaryexternalResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzmaxinvoicingsummaryexternalResponse{}

// EzmaxinvoicingsummaryexternalResponse A Ezmaxinvoicingsummaryexternal Object
type EzmaxinvoicingsummaryexternalResponse struct {
	// The unique ID of the Ezmaxinvoicingsummaryexternal
	PkiEzmaxinvoicingsummaryexternalID *int32 `json:"pkiEzmaxinvoicingsummaryexternalID,omitempty"`
	// The unique ID of the Ezmaxinvoicing
	FkiEzmaxinvoicingID *int32 `json:"fkiEzmaxinvoicingID,omitempty"`
	// The unique ID of the Billingentityexternal
	FkiBillingentityexternalID int32 `json:"fkiBillingentityexternalID"`
	// The description of the Billingentityexternal
	SBillingentityexternalDescription string `json:"sBillingentityexternalDescription"`
	// The description of the Ezmaxinvoicingsummaryexternal
	SEzmaxinvoicingsummaryexternalDescription string `json:"sEzmaxinvoicingsummaryexternalDescription"`
}

type _EzmaxinvoicingsummaryexternalResponse EzmaxinvoicingsummaryexternalResponse

// NewEzmaxinvoicingsummaryexternalResponse instantiates a new EzmaxinvoicingsummaryexternalResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzmaxinvoicingsummaryexternalResponse(fkiBillingentityexternalID int32, sBillingentityexternalDescription string, sEzmaxinvoicingsummaryexternalDescription string) *EzmaxinvoicingsummaryexternalResponse {
	this := EzmaxinvoicingsummaryexternalResponse{}
	this.FkiBillingentityexternalID = fkiBillingentityexternalID
	this.SBillingentityexternalDescription = sBillingentityexternalDescription
	this.SEzmaxinvoicingsummaryexternalDescription = sEzmaxinvoicingsummaryexternalDescription
	return &this
}

// NewEzmaxinvoicingsummaryexternalResponseWithDefaults instantiates a new EzmaxinvoicingsummaryexternalResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzmaxinvoicingsummaryexternalResponseWithDefaults() *EzmaxinvoicingsummaryexternalResponse {
	this := EzmaxinvoicingsummaryexternalResponse{}
	return &this
}

// GetPkiEzmaxinvoicingsummaryexternalID returns the PkiEzmaxinvoicingsummaryexternalID field value if set, zero value otherwise.
func (o *EzmaxinvoicingsummaryexternalResponse) GetPkiEzmaxinvoicingsummaryexternalID() int32 {
	if o == nil || IsNil(o.PkiEzmaxinvoicingsummaryexternalID) {
		var ret int32
		return ret
	}
	return *o.PkiEzmaxinvoicingsummaryexternalID
}

// GetPkiEzmaxinvoicingsummaryexternalIDOk returns a tuple with the PkiEzmaxinvoicingsummaryexternalID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingsummaryexternalResponse) GetPkiEzmaxinvoicingsummaryexternalIDOk() (*int32, bool) {
	if o == nil || IsNil(o.PkiEzmaxinvoicingsummaryexternalID) {
		return nil, false
	}
	return o.PkiEzmaxinvoicingsummaryexternalID, true
}

// HasPkiEzmaxinvoicingsummaryexternalID returns a boolean if a field has been set.
func (o *EzmaxinvoicingsummaryexternalResponse) HasPkiEzmaxinvoicingsummaryexternalID() bool {
	if o != nil && !IsNil(o.PkiEzmaxinvoicingsummaryexternalID) {
		return true
	}

	return false
}

// SetPkiEzmaxinvoicingsummaryexternalID gets a reference to the given int32 and assigns it to the PkiEzmaxinvoicingsummaryexternalID field.
func (o *EzmaxinvoicingsummaryexternalResponse) SetPkiEzmaxinvoicingsummaryexternalID(v int32) {
	o.PkiEzmaxinvoicingsummaryexternalID = &v
}

// GetFkiEzmaxinvoicingID returns the FkiEzmaxinvoicingID field value if set, zero value otherwise.
func (o *EzmaxinvoicingsummaryexternalResponse) GetFkiEzmaxinvoicingID() int32 {
	if o == nil || IsNil(o.FkiEzmaxinvoicingID) {
		var ret int32
		return ret
	}
	return *o.FkiEzmaxinvoicingID
}

// GetFkiEzmaxinvoicingIDOk returns a tuple with the FkiEzmaxinvoicingID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingsummaryexternalResponse) GetFkiEzmaxinvoicingIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiEzmaxinvoicingID) {
		return nil, false
	}
	return o.FkiEzmaxinvoicingID, true
}

// HasFkiEzmaxinvoicingID returns a boolean if a field has been set.
func (o *EzmaxinvoicingsummaryexternalResponse) HasFkiEzmaxinvoicingID() bool {
	if o != nil && !IsNil(o.FkiEzmaxinvoicingID) {
		return true
	}

	return false
}

// SetFkiEzmaxinvoicingID gets a reference to the given int32 and assigns it to the FkiEzmaxinvoicingID field.
func (o *EzmaxinvoicingsummaryexternalResponse) SetFkiEzmaxinvoicingID(v int32) {
	o.FkiEzmaxinvoicingID = &v
}

// GetFkiBillingentityexternalID returns the FkiBillingentityexternalID field value
func (o *EzmaxinvoicingsummaryexternalResponse) GetFkiBillingentityexternalID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiBillingentityexternalID
}

// GetFkiBillingentityexternalIDOk returns a tuple with the FkiBillingentityexternalID field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingsummaryexternalResponse) GetFkiBillingentityexternalIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiBillingentityexternalID, true
}

// SetFkiBillingentityexternalID sets field value
func (o *EzmaxinvoicingsummaryexternalResponse) SetFkiBillingentityexternalID(v int32) {
	o.FkiBillingentityexternalID = v
}

// GetSBillingentityexternalDescription returns the SBillingentityexternalDescription field value
func (o *EzmaxinvoicingsummaryexternalResponse) GetSBillingentityexternalDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SBillingentityexternalDescription
}

// GetSBillingentityexternalDescriptionOk returns a tuple with the SBillingentityexternalDescription field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingsummaryexternalResponse) GetSBillingentityexternalDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SBillingentityexternalDescription, true
}

// SetSBillingentityexternalDescription sets field value
func (o *EzmaxinvoicingsummaryexternalResponse) SetSBillingentityexternalDescription(v string) {
	o.SBillingentityexternalDescription = v
}

// GetSEzmaxinvoicingsummaryexternalDescription returns the SEzmaxinvoicingsummaryexternalDescription field value
func (o *EzmaxinvoicingsummaryexternalResponse) GetSEzmaxinvoicingsummaryexternalDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzmaxinvoicingsummaryexternalDescription
}

// GetSEzmaxinvoicingsummaryexternalDescriptionOk returns a tuple with the SEzmaxinvoicingsummaryexternalDescription field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingsummaryexternalResponse) GetSEzmaxinvoicingsummaryexternalDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SEzmaxinvoicingsummaryexternalDescription, true
}

// SetSEzmaxinvoicingsummaryexternalDescription sets field value
func (o *EzmaxinvoicingsummaryexternalResponse) SetSEzmaxinvoicingsummaryexternalDescription(v string) {
	o.SEzmaxinvoicingsummaryexternalDescription = v
}

func (o EzmaxinvoicingsummaryexternalResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzmaxinvoicingsummaryexternalResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PkiEzmaxinvoicingsummaryexternalID) {
		toSerialize["pkiEzmaxinvoicingsummaryexternalID"] = o.PkiEzmaxinvoicingsummaryexternalID
	}
	if !IsNil(o.FkiEzmaxinvoicingID) {
		toSerialize["fkiEzmaxinvoicingID"] = o.FkiEzmaxinvoicingID
	}
	toSerialize["fkiBillingentityexternalID"] = o.FkiBillingentityexternalID
	toSerialize["sBillingentityexternalDescription"] = o.SBillingentityexternalDescription
	toSerialize["sEzmaxinvoicingsummaryexternalDescription"] = o.SEzmaxinvoicingsummaryexternalDescription
	return toSerialize, nil
}

func (o *EzmaxinvoicingsummaryexternalResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fkiBillingentityexternalID",
		"sBillingentityexternalDescription",
		"sEzmaxinvoicingsummaryexternalDescription",
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

	varEzmaxinvoicingsummaryexternalResponse := _EzmaxinvoicingsummaryexternalResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzmaxinvoicingsummaryexternalResponse)

	if err != nil {
		return err
	}

	*o = EzmaxinvoicingsummaryexternalResponse(varEzmaxinvoicingsummaryexternalResponse)

	return err
}

type NullableEzmaxinvoicingsummaryexternalResponse struct {
	value *EzmaxinvoicingsummaryexternalResponse
	isSet bool
}

func (v NullableEzmaxinvoicingsummaryexternalResponse) Get() *EzmaxinvoicingsummaryexternalResponse {
	return v.value
}

func (v *NullableEzmaxinvoicingsummaryexternalResponse) Set(val *EzmaxinvoicingsummaryexternalResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEzmaxinvoicingsummaryexternalResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEzmaxinvoicingsummaryexternalResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzmaxinvoicingsummaryexternalResponse(val *EzmaxinvoicingsummaryexternalResponse) *NullableEzmaxinvoicingsummaryexternalResponse {
	return &NullableEzmaxinvoicingsummaryexternalResponse{value: val, isSet: true}
}

func (v NullableEzmaxinvoicingsummaryexternalResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzmaxinvoicingsummaryexternalResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


