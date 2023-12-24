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

// checks if the BillingentityinternalproductRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BillingentityinternalproductRequest{}

// BillingentityinternalproductRequest A Billingentityinternalproduct Object
type BillingentityinternalproductRequest struct {
	// The unique ID of the Billingentityinternalproduct
	PkiBillingentityinternalproductID *int32 `json:"pkiBillingentityinternalproductID,omitempty"`
	// The unique ID of the Ezmaxproduct
	FkiEzmaxproductID int32 `json:"fkiEzmaxproductID"`
	// The unique ID of the Billingentityexternal
	FkiBillingentityexternalID int32 `json:"fkiBillingentityexternalID"`
}

type _BillingentityinternalproductRequest BillingentityinternalproductRequest

// NewBillingentityinternalproductRequest instantiates a new BillingentityinternalproductRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingentityinternalproductRequest(fkiEzmaxproductID int32, fkiBillingentityexternalID int32) *BillingentityinternalproductRequest {
	this := BillingentityinternalproductRequest{}
	this.FkiEzmaxproductID = fkiEzmaxproductID
	this.FkiBillingentityexternalID = fkiBillingentityexternalID
	return &this
}

// NewBillingentityinternalproductRequestWithDefaults instantiates a new BillingentityinternalproductRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingentityinternalproductRequestWithDefaults() *BillingentityinternalproductRequest {
	this := BillingentityinternalproductRequest{}
	return &this
}

// GetPkiBillingentityinternalproductID returns the PkiBillingentityinternalproductID field value if set, zero value otherwise.
func (o *BillingentityinternalproductRequest) GetPkiBillingentityinternalproductID() int32 {
	if o == nil || IsNil(o.PkiBillingentityinternalproductID) {
		var ret int32
		return ret
	}
	return *o.PkiBillingentityinternalproductID
}

// GetPkiBillingentityinternalproductIDOk returns a tuple with the PkiBillingentityinternalproductID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingentityinternalproductRequest) GetPkiBillingentityinternalproductIDOk() (*int32, bool) {
	if o == nil || IsNil(o.PkiBillingentityinternalproductID) {
		return nil, false
	}
	return o.PkiBillingentityinternalproductID, true
}

// HasPkiBillingentityinternalproductID returns a boolean if a field has been set.
func (o *BillingentityinternalproductRequest) HasPkiBillingentityinternalproductID() bool {
	if o != nil && !IsNil(o.PkiBillingentityinternalproductID) {
		return true
	}

	return false
}

// SetPkiBillingentityinternalproductID gets a reference to the given int32 and assigns it to the PkiBillingentityinternalproductID field.
func (o *BillingentityinternalproductRequest) SetPkiBillingentityinternalproductID(v int32) {
	o.PkiBillingentityinternalproductID = &v
}

// GetFkiEzmaxproductID returns the FkiEzmaxproductID field value
func (o *BillingentityinternalproductRequest) GetFkiEzmaxproductID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzmaxproductID
}

// GetFkiEzmaxproductIDOk returns a tuple with the FkiEzmaxproductID field value
// and a boolean to check if the value has been set.
func (o *BillingentityinternalproductRequest) GetFkiEzmaxproductIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiEzmaxproductID, true
}

// SetFkiEzmaxproductID sets field value
func (o *BillingentityinternalproductRequest) SetFkiEzmaxproductID(v int32) {
	o.FkiEzmaxproductID = v
}

// GetFkiBillingentityexternalID returns the FkiBillingentityexternalID field value
func (o *BillingentityinternalproductRequest) GetFkiBillingentityexternalID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiBillingentityexternalID
}

// GetFkiBillingentityexternalIDOk returns a tuple with the FkiBillingentityexternalID field value
// and a boolean to check if the value has been set.
func (o *BillingentityinternalproductRequest) GetFkiBillingentityexternalIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiBillingentityexternalID, true
}

// SetFkiBillingentityexternalID sets field value
func (o *BillingentityinternalproductRequest) SetFkiBillingentityexternalID(v int32) {
	o.FkiBillingentityexternalID = v
}

func (o BillingentityinternalproductRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BillingentityinternalproductRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PkiBillingentityinternalproductID) {
		toSerialize["pkiBillingentityinternalproductID"] = o.PkiBillingentityinternalproductID
	}
	toSerialize["fkiEzmaxproductID"] = o.FkiEzmaxproductID
	toSerialize["fkiBillingentityexternalID"] = o.FkiBillingentityexternalID
	return toSerialize, nil
}

func (o *BillingentityinternalproductRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fkiEzmaxproductID",
		"fkiBillingentityexternalID",
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

	varBillingentityinternalproductRequest := _BillingentityinternalproductRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBillingentityinternalproductRequest)

	if err != nil {
		return err
	}

	*o = BillingentityinternalproductRequest(varBillingentityinternalproductRequest)

	return err
}

type NullableBillingentityinternalproductRequest struct {
	value *BillingentityinternalproductRequest
	isSet bool
}

func (v NullableBillingentityinternalproductRequest) Get() *BillingentityinternalproductRequest {
	return v.value
}

func (v *NullableBillingentityinternalproductRequest) Set(val *BillingentityinternalproductRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingentityinternalproductRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingentityinternalproductRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingentityinternalproductRequest(val *BillingentityinternalproductRequest) *NullableBillingentityinternalproductRequest {
	return &NullableBillingentityinternalproductRequest{value: val, isSet: true}
}

func (v NullableBillingentityinternalproductRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingentityinternalproductRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


