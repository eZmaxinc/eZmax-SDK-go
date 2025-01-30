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

// checks if the VariableexpenseRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VariableexpenseRequest{}

// VariableexpenseRequest A Variableexpense Object
type VariableexpenseRequest struct {
	// The unique ID of the Variableexpense
	PkiVariableexpenseID *int32 `json:"pkiVariableexpenseID,omitempty"`
	// The code of the Variableexpense
	SVariableexpenseCode string `json:"sVariableexpenseCode" validate:"regexp=^.{0,5}$"`
	ObjVariableexpenseDescription MultilingualVariableexpenseDescription `json:"objVariableexpenseDescription"`
	EVariableexpenseTaxable FieldEVariableexpenseTaxable `json:"eVariableexpenseTaxable"`
	// Whether the variableexpense is active or not
	BVariableexpenseIsactive bool `json:"bVariableexpenseIsactive"`
}

type _VariableexpenseRequest VariableexpenseRequest

// NewVariableexpenseRequest instantiates a new VariableexpenseRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVariableexpenseRequest(sVariableexpenseCode string, objVariableexpenseDescription MultilingualVariableexpenseDescription, eVariableexpenseTaxable FieldEVariableexpenseTaxable, bVariableexpenseIsactive bool) *VariableexpenseRequest {
	this := VariableexpenseRequest{}
	this.SVariableexpenseCode = sVariableexpenseCode
	this.ObjVariableexpenseDescription = objVariableexpenseDescription
	this.EVariableexpenseTaxable = eVariableexpenseTaxable
	this.BVariableexpenseIsactive = bVariableexpenseIsactive
	return &this
}

// NewVariableexpenseRequestWithDefaults instantiates a new VariableexpenseRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVariableexpenseRequestWithDefaults() *VariableexpenseRequest {
	this := VariableexpenseRequest{}
	return &this
}

// GetPkiVariableexpenseID returns the PkiVariableexpenseID field value if set, zero value otherwise.
func (o *VariableexpenseRequest) GetPkiVariableexpenseID() int32 {
	if o == nil || IsNil(o.PkiVariableexpenseID) {
		var ret int32
		return ret
	}
	return *o.PkiVariableexpenseID
}

// GetPkiVariableexpenseIDOk returns a tuple with the PkiVariableexpenseID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariableexpenseRequest) GetPkiVariableexpenseIDOk() (*int32, bool) {
	if o == nil || IsNil(o.PkiVariableexpenseID) {
		return nil, false
	}
	return o.PkiVariableexpenseID, true
}

// HasPkiVariableexpenseID returns a boolean if a field has been set.
func (o *VariableexpenseRequest) HasPkiVariableexpenseID() bool {
	if o != nil && !IsNil(o.PkiVariableexpenseID) {
		return true
	}

	return false
}

// SetPkiVariableexpenseID gets a reference to the given int32 and assigns it to the PkiVariableexpenseID field.
func (o *VariableexpenseRequest) SetPkiVariableexpenseID(v int32) {
	o.PkiVariableexpenseID = &v
}

// GetSVariableexpenseCode returns the SVariableexpenseCode field value
func (o *VariableexpenseRequest) GetSVariableexpenseCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SVariableexpenseCode
}

// GetSVariableexpenseCodeOk returns a tuple with the SVariableexpenseCode field value
// and a boolean to check if the value has been set.
func (o *VariableexpenseRequest) GetSVariableexpenseCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SVariableexpenseCode, true
}

// SetSVariableexpenseCode sets field value
func (o *VariableexpenseRequest) SetSVariableexpenseCode(v string) {
	o.SVariableexpenseCode = v
}

// GetObjVariableexpenseDescription returns the ObjVariableexpenseDescription field value
func (o *VariableexpenseRequest) GetObjVariableexpenseDescription() MultilingualVariableexpenseDescription {
	if o == nil {
		var ret MultilingualVariableexpenseDescription
		return ret
	}

	return o.ObjVariableexpenseDescription
}

// GetObjVariableexpenseDescriptionOk returns a tuple with the ObjVariableexpenseDescription field value
// and a boolean to check if the value has been set.
func (o *VariableexpenseRequest) GetObjVariableexpenseDescriptionOk() (*MultilingualVariableexpenseDescription, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjVariableexpenseDescription, true
}

// SetObjVariableexpenseDescription sets field value
func (o *VariableexpenseRequest) SetObjVariableexpenseDescription(v MultilingualVariableexpenseDescription) {
	o.ObjVariableexpenseDescription = v
}

// GetEVariableexpenseTaxable returns the EVariableexpenseTaxable field value
func (o *VariableexpenseRequest) GetEVariableexpenseTaxable() FieldEVariableexpenseTaxable {
	if o == nil {
		var ret FieldEVariableexpenseTaxable
		return ret
	}

	return o.EVariableexpenseTaxable
}

// GetEVariableexpenseTaxableOk returns a tuple with the EVariableexpenseTaxable field value
// and a boolean to check if the value has been set.
func (o *VariableexpenseRequest) GetEVariableexpenseTaxableOk() (*FieldEVariableexpenseTaxable, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EVariableexpenseTaxable, true
}

// SetEVariableexpenseTaxable sets field value
func (o *VariableexpenseRequest) SetEVariableexpenseTaxable(v FieldEVariableexpenseTaxable) {
	o.EVariableexpenseTaxable = v
}

// GetBVariableexpenseIsactive returns the BVariableexpenseIsactive field value
func (o *VariableexpenseRequest) GetBVariableexpenseIsactive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BVariableexpenseIsactive
}

// GetBVariableexpenseIsactiveOk returns a tuple with the BVariableexpenseIsactive field value
// and a boolean to check if the value has been set.
func (o *VariableexpenseRequest) GetBVariableexpenseIsactiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BVariableexpenseIsactive, true
}

// SetBVariableexpenseIsactive sets field value
func (o *VariableexpenseRequest) SetBVariableexpenseIsactive(v bool) {
	o.BVariableexpenseIsactive = v
}

func (o VariableexpenseRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VariableexpenseRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PkiVariableexpenseID) {
		toSerialize["pkiVariableexpenseID"] = o.PkiVariableexpenseID
	}
	toSerialize["sVariableexpenseCode"] = o.SVariableexpenseCode
	toSerialize["objVariableexpenseDescription"] = o.ObjVariableexpenseDescription
	toSerialize["eVariableexpenseTaxable"] = o.EVariableexpenseTaxable
	toSerialize["bVariableexpenseIsactive"] = o.BVariableexpenseIsactive
	return toSerialize, nil
}

func (o *VariableexpenseRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"sVariableexpenseCode",
		"objVariableexpenseDescription",
		"eVariableexpenseTaxable",
		"bVariableexpenseIsactive",
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

	varVariableexpenseRequest := _VariableexpenseRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVariableexpenseRequest)

	if err != nil {
		return err
	}

	*o = VariableexpenseRequest(varVariableexpenseRequest)

	return err
}

type NullableVariableexpenseRequest struct {
	value *VariableexpenseRequest
	isSet bool
}

func (v NullableVariableexpenseRequest) Get() *VariableexpenseRequest {
	return v.value
}

func (v *NullableVariableexpenseRequest) Set(val *VariableexpenseRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableVariableexpenseRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableVariableexpenseRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVariableexpenseRequest(val *VariableexpenseRequest) *NullableVariableexpenseRequest {
	return &NullableVariableexpenseRequest{value: val, isSet: true}
}

func (v NullableVariableexpenseRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVariableexpenseRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


