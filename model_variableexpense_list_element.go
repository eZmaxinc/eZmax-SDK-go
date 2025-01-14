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

// checks if the VariableexpenseListElement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VariableexpenseListElement{}

// VariableexpenseListElement A Variableexpense List Element
type VariableexpenseListElement struct {
	// The unique ID of the Variableexpense
	PkiVariableexpenseID int32 `json:"pkiVariableexpenseID"`
	// The code of the Variableexpense
	SVariableexpenseCode *string `json:"sVariableexpenseCode,omitempty" validate:"regexp=^.{0,5}$"`
	// The description of the Variableexpense in the language of the requester
	SVariableexpenseDescriptionX *string `json:"sVariableexpenseDescriptionX,omitempty" validate:"regexp=^.{0,40}$"`
	EVariableexpenseTaxable *FieldEVariableexpenseTaxable `json:"eVariableexpenseTaxable,omitempty"`
	// Whether the variableexpense is active or not
	BVariableexpenseIsactive *bool `json:"bVariableexpenseIsactive,omitempty"`
}

type _VariableexpenseListElement VariableexpenseListElement

// NewVariableexpenseListElement instantiates a new VariableexpenseListElement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVariableexpenseListElement(pkiVariableexpenseID int32) *VariableexpenseListElement {
	this := VariableexpenseListElement{}
	this.PkiVariableexpenseID = pkiVariableexpenseID
	return &this
}

// NewVariableexpenseListElementWithDefaults instantiates a new VariableexpenseListElement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVariableexpenseListElementWithDefaults() *VariableexpenseListElement {
	this := VariableexpenseListElement{}
	return &this
}

// GetPkiVariableexpenseID returns the PkiVariableexpenseID field value
func (o *VariableexpenseListElement) GetPkiVariableexpenseID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiVariableexpenseID
}

// GetPkiVariableexpenseIDOk returns a tuple with the PkiVariableexpenseID field value
// and a boolean to check if the value has been set.
func (o *VariableexpenseListElement) GetPkiVariableexpenseIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiVariableexpenseID, true
}

// SetPkiVariableexpenseID sets field value
func (o *VariableexpenseListElement) SetPkiVariableexpenseID(v int32) {
	o.PkiVariableexpenseID = v
}

// GetSVariableexpenseCode returns the SVariableexpenseCode field value if set, zero value otherwise.
func (o *VariableexpenseListElement) GetSVariableexpenseCode() string {
	if o == nil || IsNil(o.SVariableexpenseCode) {
		var ret string
		return ret
	}
	return *o.SVariableexpenseCode
}

// GetSVariableexpenseCodeOk returns a tuple with the SVariableexpenseCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariableexpenseListElement) GetSVariableexpenseCodeOk() (*string, bool) {
	if o == nil || IsNil(o.SVariableexpenseCode) {
		return nil, false
	}
	return o.SVariableexpenseCode, true
}

// HasSVariableexpenseCode returns a boolean if a field has been set.
func (o *VariableexpenseListElement) HasSVariableexpenseCode() bool {
	if o != nil && !IsNil(o.SVariableexpenseCode) {
		return true
	}

	return false
}

// SetSVariableexpenseCode gets a reference to the given string and assigns it to the SVariableexpenseCode field.
func (o *VariableexpenseListElement) SetSVariableexpenseCode(v string) {
	o.SVariableexpenseCode = &v
}

// GetSVariableexpenseDescriptionX returns the SVariableexpenseDescriptionX field value if set, zero value otherwise.
func (o *VariableexpenseListElement) GetSVariableexpenseDescriptionX() string {
	if o == nil || IsNil(o.SVariableexpenseDescriptionX) {
		var ret string
		return ret
	}
	return *o.SVariableexpenseDescriptionX
}

// GetSVariableexpenseDescriptionXOk returns a tuple with the SVariableexpenseDescriptionX field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariableexpenseListElement) GetSVariableexpenseDescriptionXOk() (*string, bool) {
	if o == nil || IsNil(o.SVariableexpenseDescriptionX) {
		return nil, false
	}
	return o.SVariableexpenseDescriptionX, true
}

// HasSVariableexpenseDescriptionX returns a boolean if a field has been set.
func (o *VariableexpenseListElement) HasSVariableexpenseDescriptionX() bool {
	if o != nil && !IsNil(o.SVariableexpenseDescriptionX) {
		return true
	}

	return false
}

// SetSVariableexpenseDescriptionX gets a reference to the given string and assigns it to the SVariableexpenseDescriptionX field.
func (o *VariableexpenseListElement) SetSVariableexpenseDescriptionX(v string) {
	o.SVariableexpenseDescriptionX = &v
}

// GetEVariableexpenseTaxable returns the EVariableexpenseTaxable field value if set, zero value otherwise.
func (o *VariableexpenseListElement) GetEVariableexpenseTaxable() FieldEVariableexpenseTaxable {
	if o == nil || IsNil(o.EVariableexpenseTaxable) {
		var ret FieldEVariableexpenseTaxable
		return ret
	}
	return *o.EVariableexpenseTaxable
}

// GetEVariableexpenseTaxableOk returns a tuple with the EVariableexpenseTaxable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariableexpenseListElement) GetEVariableexpenseTaxableOk() (*FieldEVariableexpenseTaxable, bool) {
	if o == nil || IsNil(o.EVariableexpenseTaxable) {
		return nil, false
	}
	return o.EVariableexpenseTaxable, true
}

// HasEVariableexpenseTaxable returns a boolean if a field has been set.
func (o *VariableexpenseListElement) HasEVariableexpenseTaxable() bool {
	if o != nil && !IsNil(o.EVariableexpenseTaxable) {
		return true
	}

	return false
}

// SetEVariableexpenseTaxable gets a reference to the given FieldEVariableexpenseTaxable and assigns it to the EVariableexpenseTaxable field.
func (o *VariableexpenseListElement) SetEVariableexpenseTaxable(v FieldEVariableexpenseTaxable) {
	o.EVariableexpenseTaxable = &v
}

// GetBVariableexpenseIsactive returns the BVariableexpenseIsactive field value if set, zero value otherwise.
func (o *VariableexpenseListElement) GetBVariableexpenseIsactive() bool {
	if o == nil || IsNil(o.BVariableexpenseIsactive) {
		var ret bool
		return ret
	}
	return *o.BVariableexpenseIsactive
}

// GetBVariableexpenseIsactiveOk returns a tuple with the BVariableexpenseIsactive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariableexpenseListElement) GetBVariableexpenseIsactiveOk() (*bool, bool) {
	if o == nil || IsNil(o.BVariableexpenseIsactive) {
		return nil, false
	}
	return o.BVariableexpenseIsactive, true
}

// HasBVariableexpenseIsactive returns a boolean if a field has been set.
func (o *VariableexpenseListElement) HasBVariableexpenseIsactive() bool {
	if o != nil && !IsNil(o.BVariableexpenseIsactive) {
		return true
	}

	return false
}

// SetBVariableexpenseIsactive gets a reference to the given bool and assigns it to the BVariableexpenseIsactive field.
func (o *VariableexpenseListElement) SetBVariableexpenseIsactive(v bool) {
	o.BVariableexpenseIsactive = &v
}

func (o VariableexpenseListElement) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VariableexpenseListElement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiVariableexpenseID"] = o.PkiVariableexpenseID
	if !IsNil(o.SVariableexpenseCode) {
		toSerialize["sVariableexpenseCode"] = o.SVariableexpenseCode
	}
	if !IsNil(o.SVariableexpenseDescriptionX) {
		toSerialize["sVariableexpenseDescriptionX"] = o.SVariableexpenseDescriptionX
	}
	if !IsNil(o.EVariableexpenseTaxable) {
		toSerialize["eVariableexpenseTaxable"] = o.EVariableexpenseTaxable
	}
	if !IsNil(o.BVariableexpenseIsactive) {
		toSerialize["bVariableexpenseIsactive"] = o.BVariableexpenseIsactive
	}
	return toSerialize, nil
}

func (o *VariableexpenseListElement) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiVariableexpenseID",
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

	varVariableexpenseListElement := _VariableexpenseListElement{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVariableexpenseListElement)

	if err != nil {
		return err
	}

	*o = VariableexpenseListElement(varVariableexpenseListElement)

	return err
}

type NullableVariableexpenseListElement struct {
	value *VariableexpenseListElement
	isSet bool
}

func (v NullableVariableexpenseListElement) Get() *VariableexpenseListElement {
	return v.value
}

func (v *NullableVariableexpenseListElement) Set(val *VariableexpenseListElement) {
	v.value = val
	v.isSet = true
}

func (v NullableVariableexpenseListElement) IsSet() bool {
	return v.isSet
}

func (v *NullableVariableexpenseListElement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVariableexpenseListElement(val *VariableexpenseListElement) *NullableVariableexpenseListElement {
	return &NullableVariableexpenseListElement{value: val, isSet: true}
}

func (v NullableVariableexpenseListElement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVariableexpenseListElement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


