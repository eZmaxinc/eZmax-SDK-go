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

// checks if the VariableexpenseResponseCompound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VariableexpenseResponseCompound{}

// VariableexpenseResponseCompound A Variableexpense Object
type VariableexpenseResponseCompound struct {
	// The unique ID of the Variableexpense
	PkiVariableexpenseID int32 `json:"pkiVariableexpenseID"`
	// The code of the Variableexpense
	SVariableexpenseCode *string `json:"sVariableexpenseCode,omitempty"`
	ObjVariableexpenseDescription MultilingualVariableexpenseDescription `json:"objVariableexpenseDescription"`
	EVariableexpenseTaxable *FieldEVariableexpenseTaxable `json:"eVariableexpenseTaxable,omitempty"`
	// Whether the variableexpense is active or not
	BVariableexpenseIsactive *bool `json:"bVariableexpenseIsactive,omitempty"`
}

type _VariableexpenseResponseCompound VariableexpenseResponseCompound

// NewVariableexpenseResponseCompound instantiates a new VariableexpenseResponseCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVariableexpenseResponseCompound(pkiVariableexpenseID int32, objVariableexpenseDescription MultilingualVariableexpenseDescription) *VariableexpenseResponseCompound {
	this := VariableexpenseResponseCompound{}
	this.PkiVariableexpenseID = pkiVariableexpenseID
	this.ObjVariableexpenseDescription = objVariableexpenseDescription
	return &this
}

// NewVariableexpenseResponseCompoundWithDefaults instantiates a new VariableexpenseResponseCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVariableexpenseResponseCompoundWithDefaults() *VariableexpenseResponseCompound {
	this := VariableexpenseResponseCompound{}
	return &this
}

// GetPkiVariableexpenseID returns the PkiVariableexpenseID field value
func (o *VariableexpenseResponseCompound) GetPkiVariableexpenseID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiVariableexpenseID
}

// GetPkiVariableexpenseIDOk returns a tuple with the PkiVariableexpenseID field value
// and a boolean to check if the value has been set.
func (o *VariableexpenseResponseCompound) GetPkiVariableexpenseIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiVariableexpenseID, true
}

// SetPkiVariableexpenseID sets field value
func (o *VariableexpenseResponseCompound) SetPkiVariableexpenseID(v int32) {
	o.PkiVariableexpenseID = v
}

// GetSVariableexpenseCode returns the SVariableexpenseCode field value if set, zero value otherwise.
func (o *VariableexpenseResponseCompound) GetSVariableexpenseCode() string {
	if o == nil || IsNil(o.SVariableexpenseCode) {
		var ret string
		return ret
	}
	return *o.SVariableexpenseCode
}

// GetSVariableexpenseCodeOk returns a tuple with the SVariableexpenseCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariableexpenseResponseCompound) GetSVariableexpenseCodeOk() (*string, bool) {
	if o == nil || IsNil(o.SVariableexpenseCode) {
		return nil, false
	}
	return o.SVariableexpenseCode, true
}

// HasSVariableexpenseCode returns a boolean if a field has been set.
func (o *VariableexpenseResponseCompound) HasSVariableexpenseCode() bool {
	if o != nil && !IsNil(o.SVariableexpenseCode) {
		return true
	}

	return false
}

// SetSVariableexpenseCode gets a reference to the given string and assigns it to the SVariableexpenseCode field.
func (o *VariableexpenseResponseCompound) SetSVariableexpenseCode(v string) {
	o.SVariableexpenseCode = &v
}

// GetObjVariableexpenseDescription returns the ObjVariableexpenseDescription field value
func (o *VariableexpenseResponseCompound) GetObjVariableexpenseDescription() MultilingualVariableexpenseDescription {
	if o == nil {
		var ret MultilingualVariableexpenseDescription
		return ret
	}

	return o.ObjVariableexpenseDescription
}

// GetObjVariableexpenseDescriptionOk returns a tuple with the ObjVariableexpenseDescription field value
// and a boolean to check if the value has been set.
func (o *VariableexpenseResponseCompound) GetObjVariableexpenseDescriptionOk() (*MultilingualVariableexpenseDescription, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjVariableexpenseDescription, true
}

// SetObjVariableexpenseDescription sets field value
func (o *VariableexpenseResponseCompound) SetObjVariableexpenseDescription(v MultilingualVariableexpenseDescription) {
	o.ObjVariableexpenseDescription = v
}

// GetEVariableexpenseTaxable returns the EVariableexpenseTaxable field value if set, zero value otherwise.
func (o *VariableexpenseResponseCompound) GetEVariableexpenseTaxable() FieldEVariableexpenseTaxable {
	if o == nil || IsNil(o.EVariableexpenseTaxable) {
		var ret FieldEVariableexpenseTaxable
		return ret
	}
	return *o.EVariableexpenseTaxable
}

// GetEVariableexpenseTaxableOk returns a tuple with the EVariableexpenseTaxable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariableexpenseResponseCompound) GetEVariableexpenseTaxableOk() (*FieldEVariableexpenseTaxable, bool) {
	if o == nil || IsNil(o.EVariableexpenseTaxable) {
		return nil, false
	}
	return o.EVariableexpenseTaxable, true
}

// HasEVariableexpenseTaxable returns a boolean if a field has been set.
func (o *VariableexpenseResponseCompound) HasEVariableexpenseTaxable() bool {
	if o != nil && !IsNil(o.EVariableexpenseTaxable) {
		return true
	}

	return false
}

// SetEVariableexpenseTaxable gets a reference to the given FieldEVariableexpenseTaxable and assigns it to the EVariableexpenseTaxable field.
func (o *VariableexpenseResponseCompound) SetEVariableexpenseTaxable(v FieldEVariableexpenseTaxable) {
	o.EVariableexpenseTaxable = &v
}

// GetBVariableexpenseIsactive returns the BVariableexpenseIsactive field value if set, zero value otherwise.
func (o *VariableexpenseResponseCompound) GetBVariableexpenseIsactive() bool {
	if o == nil || IsNil(o.BVariableexpenseIsactive) {
		var ret bool
		return ret
	}
	return *o.BVariableexpenseIsactive
}

// GetBVariableexpenseIsactiveOk returns a tuple with the BVariableexpenseIsactive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariableexpenseResponseCompound) GetBVariableexpenseIsactiveOk() (*bool, bool) {
	if o == nil || IsNil(o.BVariableexpenseIsactive) {
		return nil, false
	}
	return o.BVariableexpenseIsactive, true
}

// HasBVariableexpenseIsactive returns a boolean if a field has been set.
func (o *VariableexpenseResponseCompound) HasBVariableexpenseIsactive() bool {
	if o != nil && !IsNil(o.BVariableexpenseIsactive) {
		return true
	}

	return false
}

// SetBVariableexpenseIsactive gets a reference to the given bool and assigns it to the BVariableexpenseIsactive field.
func (o *VariableexpenseResponseCompound) SetBVariableexpenseIsactive(v bool) {
	o.BVariableexpenseIsactive = &v
}

func (o VariableexpenseResponseCompound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VariableexpenseResponseCompound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiVariableexpenseID"] = o.PkiVariableexpenseID
	if !IsNil(o.SVariableexpenseCode) {
		toSerialize["sVariableexpenseCode"] = o.SVariableexpenseCode
	}
	toSerialize["objVariableexpenseDescription"] = o.ObjVariableexpenseDescription
	if !IsNil(o.EVariableexpenseTaxable) {
		toSerialize["eVariableexpenseTaxable"] = o.EVariableexpenseTaxable
	}
	if !IsNil(o.BVariableexpenseIsactive) {
		toSerialize["bVariableexpenseIsactive"] = o.BVariableexpenseIsactive
	}
	return toSerialize, nil
}

func (o *VariableexpenseResponseCompound) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiVariableexpenseID",
		"objVariableexpenseDescription",
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

	varVariableexpenseResponseCompound := _VariableexpenseResponseCompound{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVariableexpenseResponseCompound)

	if err != nil {
		return err
	}

	*o = VariableexpenseResponseCompound(varVariableexpenseResponseCompound)

	return err
}

type NullableVariableexpenseResponseCompound struct {
	value *VariableexpenseResponseCompound
	isSet bool
}

func (v NullableVariableexpenseResponseCompound) Get() *VariableexpenseResponseCompound {
	return v.value
}

func (v *NullableVariableexpenseResponseCompound) Set(val *VariableexpenseResponseCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableVariableexpenseResponseCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableVariableexpenseResponseCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVariableexpenseResponseCompound(val *VariableexpenseResponseCompound) *NullableVariableexpenseResponseCompound {
	return &NullableVariableexpenseResponseCompound{value: val, isSet: true}
}

func (v NullableVariableexpenseResponseCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVariableexpenseResponseCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


