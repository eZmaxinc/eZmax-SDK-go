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

// checks if the EzsignelementdependencyRequestCompound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignelementdependencyRequestCompound{}

// EzsignelementdependencyRequestCompound An Ezsignelementdependency Object and children to create a complete structure
type EzsignelementdependencyRequestCompound struct {
	// The unique ID of the Ezsignelementdependency
	PkiEzsignelementdependencyID *int32 `json:"pkiEzsignelementdependencyID,omitempty"`
	// The unique ID of the Ezsignformfield
	FkiEzsignformfieldIDValidation *int32 `json:"fkiEzsignformfieldIDValidation,omitempty"`
	// The unique ID of the Ezsignformfieldgroup
	FkiEzsignformfieldgroupIDValidation *int32 `json:"fkiEzsignformfieldgroupIDValidation,omitempty"`
	// The Label for the Ezsignformfieldgroup
	SEzsignelementdependencyEzsignformfieldgrouplabel *string `json:"sEzsignelementdependencyEzsignformfieldgrouplabel,omitempty"`
	// The Label for the Ezsignformfield
	SEzsignelementdependencyEzsignformfieldlabel *string `json:"sEzsignelementdependencyEzsignformfieldlabel,omitempty"`
	EEzsignelementdependencyValidation FieldEEzsignelementdependencyValidation `json:"eEzsignelementdependencyValidation"`
	// Whether if it's selected or not when using eEzsignelementdependencyValidation = Selected
	BEzsignelementdependencySelected *bool `json:"bEzsignelementdependencySelected,omitempty"`
	EEzsignelementdependencyOperator *FieldEEzsignelementdependencyOperator `json:"eEzsignelementdependencyOperator,omitempty"`
	// The value of the Ezsignelementdependency
	SEzsignelementdependencyValue *string `json:"sEzsignelementdependencyValue,omitempty"`
}

type _EzsignelementdependencyRequestCompound EzsignelementdependencyRequestCompound

// NewEzsignelementdependencyRequestCompound instantiates a new EzsignelementdependencyRequestCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignelementdependencyRequestCompound(eEzsignelementdependencyValidation FieldEEzsignelementdependencyValidation) *EzsignelementdependencyRequestCompound {
	this := EzsignelementdependencyRequestCompound{}
	this.EEzsignelementdependencyValidation = eEzsignelementdependencyValidation
	return &this
}

// NewEzsignelementdependencyRequestCompoundWithDefaults instantiates a new EzsignelementdependencyRequestCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignelementdependencyRequestCompoundWithDefaults() *EzsignelementdependencyRequestCompound {
	this := EzsignelementdependencyRequestCompound{}
	return &this
}

// GetPkiEzsignelementdependencyID returns the PkiEzsignelementdependencyID field value if set, zero value otherwise.
func (o *EzsignelementdependencyRequestCompound) GetPkiEzsignelementdependencyID() int32 {
	if o == nil || IsNil(o.PkiEzsignelementdependencyID) {
		var ret int32
		return ret
	}
	return *o.PkiEzsignelementdependencyID
}

// GetPkiEzsignelementdependencyIDOk returns a tuple with the PkiEzsignelementdependencyID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignelementdependencyRequestCompound) GetPkiEzsignelementdependencyIDOk() (*int32, bool) {
	if o == nil || IsNil(o.PkiEzsignelementdependencyID) {
		return nil, false
	}
	return o.PkiEzsignelementdependencyID, true
}

// HasPkiEzsignelementdependencyID returns a boolean if a field has been set.
func (o *EzsignelementdependencyRequestCompound) HasPkiEzsignelementdependencyID() bool {
	if o != nil && !IsNil(o.PkiEzsignelementdependencyID) {
		return true
	}

	return false
}

// SetPkiEzsignelementdependencyID gets a reference to the given int32 and assigns it to the PkiEzsignelementdependencyID field.
func (o *EzsignelementdependencyRequestCompound) SetPkiEzsignelementdependencyID(v int32) {
	o.PkiEzsignelementdependencyID = &v
}

// GetFkiEzsignformfieldIDValidation returns the FkiEzsignformfieldIDValidation field value if set, zero value otherwise.
func (o *EzsignelementdependencyRequestCompound) GetFkiEzsignformfieldIDValidation() int32 {
	if o == nil || IsNil(o.FkiEzsignformfieldIDValidation) {
		var ret int32
		return ret
	}
	return *o.FkiEzsignformfieldIDValidation
}

// GetFkiEzsignformfieldIDValidationOk returns a tuple with the FkiEzsignformfieldIDValidation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignelementdependencyRequestCompound) GetFkiEzsignformfieldIDValidationOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiEzsignformfieldIDValidation) {
		return nil, false
	}
	return o.FkiEzsignformfieldIDValidation, true
}

// HasFkiEzsignformfieldIDValidation returns a boolean if a field has been set.
func (o *EzsignelementdependencyRequestCompound) HasFkiEzsignformfieldIDValidation() bool {
	if o != nil && !IsNil(o.FkiEzsignformfieldIDValidation) {
		return true
	}

	return false
}

// SetFkiEzsignformfieldIDValidation gets a reference to the given int32 and assigns it to the FkiEzsignformfieldIDValidation field.
func (o *EzsignelementdependencyRequestCompound) SetFkiEzsignformfieldIDValidation(v int32) {
	o.FkiEzsignformfieldIDValidation = &v
}

// GetFkiEzsignformfieldgroupIDValidation returns the FkiEzsignformfieldgroupIDValidation field value if set, zero value otherwise.
func (o *EzsignelementdependencyRequestCompound) GetFkiEzsignformfieldgroupIDValidation() int32 {
	if o == nil || IsNil(o.FkiEzsignformfieldgroupIDValidation) {
		var ret int32
		return ret
	}
	return *o.FkiEzsignformfieldgroupIDValidation
}

// GetFkiEzsignformfieldgroupIDValidationOk returns a tuple with the FkiEzsignformfieldgroupIDValidation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignelementdependencyRequestCompound) GetFkiEzsignformfieldgroupIDValidationOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiEzsignformfieldgroupIDValidation) {
		return nil, false
	}
	return o.FkiEzsignformfieldgroupIDValidation, true
}

// HasFkiEzsignformfieldgroupIDValidation returns a boolean if a field has been set.
func (o *EzsignelementdependencyRequestCompound) HasFkiEzsignformfieldgroupIDValidation() bool {
	if o != nil && !IsNil(o.FkiEzsignformfieldgroupIDValidation) {
		return true
	}

	return false
}

// SetFkiEzsignformfieldgroupIDValidation gets a reference to the given int32 and assigns it to the FkiEzsignformfieldgroupIDValidation field.
func (o *EzsignelementdependencyRequestCompound) SetFkiEzsignformfieldgroupIDValidation(v int32) {
	o.FkiEzsignformfieldgroupIDValidation = &v
}

// GetSEzsignelementdependencyEzsignformfieldgrouplabel returns the SEzsignelementdependencyEzsignformfieldgrouplabel field value if set, zero value otherwise.
func (o *EzsignelementdependencyRequestCompound) GetSEzsignelementdependencyEzsignformfieldgrouplabel() string {
	if o == nil || IsNil(o.SEzsignelementdependencyEzsignformfieldgrouplabel) {
		var ret string
		return ret
	}
	return *o.SEzsignelementdependencyEzsignformfieldgrouplabel
}

// GetSEzsignelementdependencyEzsignformfieldgrouplabelOk returns a tuple with the SEzsignelementdependencyEzsignformfieldgrouplabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignelementdependencyRequestCompound) GetSEzsignelementdependencyEzsignformfieldgrouplabelOk() (*string, bool) {
	if o == nil || IsNil(o.SEzsignelementdependencyEzsignformfieldgrouplabel) {
		return nil, false
	}
	return o.SEzsignelementdependencyEzsignformfieldgrouplabel, true
}

// HasSEzsignelementdependencyEzsignformfieldgrouplabel returns a boolean if a field has been set.
func (o *EzsignelementdependencyRequestCompound) HasSEzsignelementdependencyEzsignformfieldgrouplabel() bool {
	if o != nil && !IsNil(o.SEzsignelementdependencyEzsignformfieldgrouplabel) {
		return true
	}

	return false
}

// SetSEzsignelementdependencyEzsignformfieldgrouplabel gets a reference to the given string and assigns it to the SEzsignelementdependencyEzsignformfieldgrouplabel field.
func (o *EzsignelementdependencyRequestCompound) SetSEzsignelementdependencyEzsignformfieldgrouplabel(v string) {
	o.SEzsignelementdependencyEzsignformfieldgrouplabel = &v
}

// GetSEzsignelementdependencyEzsignformfieldlabel returns the SEzsignelementdependencyEzsignformfieldlabel field value if set, zero value otherwise.
func (o *EzsignelementdependencyRequestCompound) GetSEzsignelementdependencyEzsignformfieldlabel() string {
	if o == nil || IsNil(o.SEzsignelementdependencyEzsignformfieldlabel) {
		var ret string
		return ret
	}
	return *o.SEzsignelementdependencyEzsignformfieldlabel
}

// GetSEzsignelementdependencyEzsignformfieldlabelOk returns a tuple with the SEzsignelementdependencyEzsignformfieldlabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignelementdependencyRequestCompound) GetSEzsignelementdependencyEzsignformfieldlabelOk() (*string, bool) {
	if o == nil || IsNil(o.SEzsignelementdependencyEzsignformfieldlabel) {
		return nil, false
	}
	return o.SEzsignelementdependencyEzsignformfieldlabel, true
}

// HasSEzsignelementdependencyEzsignformfieldlabel returns a boolean if a field has been set.
func (o *EzsignelementdependencyRequestCompound) HasSEzsignelementdependencyEzsignformfieldlabel() bool {
	if o != nil && !IsNil(o.SEzsignelementdependencyEzsignformfieldlabel) {
		return true
	}

	return false
}

// SetSEzsignelementdependencyEzsignformfieldlabel gets a reference to the given string and assigns it to the SEzsignelementdependencyEzsignformfieldlabel field.
func (o *EzsignelementdependencyRequestCompound) SetSEzsignelementdependencyEzsignformfieldlabel(v string) {
	o.SEzsignelementdependencyEzsignformfieldlabel = &v
}

// GetEEzsignelementdependencyValidation returns the EEzsignelementdependencyValidation field value
func (o *EzsignelementdependencyRequestCompound) GetEEzsignelementdependencyValidation() FieldEEzsignelementdependencyValidation {
	if o == nil {
		var ret FieldEEzsignelementdependencyValidation
		return ret
	}

	return o.EEzsignelementdependencyValidation
}

// GetEEzsignelementdependencyValidationOk returns a tuple with the EEzsignelementdependencyValidation field value
// and a boolean to check if the value has been set.
func (o *EzsignelementdependencyRequestCompound) GetEEzsignelementdependencyValidationOk() (*FieldEEzsignelementdependencyValidation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EEzsignelementdependencyValidation, true
}

// SetEEzsignelementdependencyValidation sets field value
func (o *EzsignelementdependencyRequestCompound) SetEEzsignelementdependencyValidation(v FieldEEzsignelementdependencyValidation) {
	o.EEzsignelementdependencyValidation = v
}

// GetBEzsignelementdependencySelected returns the BEzsignelementdependencySelected field value if set, zero value otherwise.
func (o *EzsignelementdependencyRequestCompound) GetBEzsignelementdependencySelected() bool {
	if o == nil || IsNil(o.BEzsignelementdependencySelected) {
		var ret bool
		return ret
	}
	return *o.BEzsignelementdependencySelected
}

// GetBEzsignelementdependencySelectedOk returns a tuple with the BEzsignelementdependencySelected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignelementdependencyRequestCompound) GetBEzsignelementdependencySelectedOk() (*bool, bool) {
	if o == nil || IsNil(o.BEzsignelementdependencySelected) {
		return nil, false
	}
	return o.BEzsignelementdependencySelected, true
}

// HasBEzsignelementdependencySelected returns a boolean if a field has been set.
func (o *EzsignelementdependencyRequestCompound) HasBEzsignelementdependencySelected() bool {
	if o != nil && !IsNil(o.BEzsignelementdependencySelected) {
		return true
	}

	return false
}

// SetBEzsignelementdependencySelected gets a reference to the given bool and assigns it to the BEzsignelementdependencySelected field.
func (o *EzsignelementdependencyRequestCompound) SetBEzsignelementdependencySelected(v bool) {
	o.BEzsignelementdependencySelected = &v
}

// GetEEzsignelementdependencyOperator returns the EEzsignelementdependencyOperator field value if set, zero value otherwise.
func (o *EzsignelementdependencyRequestCompound) GetEEzsignelementdependencyOperator() FieldEEzsignelementdependencyOperator {
	if o == nil || IsNil(o.EEzsignelementdependencyOperator) {
		var ret FieldEEzsignelementdependencyOperator
		return ret
	}
	return *o.EEzsignelementdependencyOperator
}

// GetEEzsignelementdependencyOperatorOk returns a tuple with the EEzsignelementdependencyOperator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignelementdependencyRequestCompound) GetEEzsignelementdependencyOperatorOk() (*FieldEEzsignelementdependencyOperator, bool) {
	if o == nil || IsNil(o.EEzsignelementdependencyOperator) {
		return nil, false
	}
	return o.EEzsignelementdependencyOperator, true
}

// HasEEzsignelementdependencyOperator returns a boolean if a field has been set.
func (o *EzsignelementdependencyRequestCompound) HasEEzsignelementdependencyOperator() bool {
	if o != nil && !IsNil(o.EEzsignelementdependencyOperator) {
		return true
	}

	return false
}

// SetEEzsignelementdependencyOperator gets a reference to the given FieldEEzsignelementdependencyOperator and assigns it to the EEzsignelementdependencyOperator field.
func (o *EzsignelementdependencyRequestCompound) SetEEzsignelementdependencyOperator(v FieldEEzsignelementdependencyOperator) {
	o.EEzsignelementdependencyOperator = &v
}

// GetSEzsignelementdependencyValue returns the SEzsignelementdependencyValue field value if set, zero value otherwise.
func (o *EzsignelementdependencyRequestCompound) GetSEzsignelementdependencyValue() string {
	if o == nil || IsNil(o.SEzsignelementdependencyValue) {
		var ret string
		return ret
	}
	return *o.SEzsignelementdependencyValue
}

// GetSEzsignelementdependencyValueOk returns a tuple with the SEzsignelementdependencyValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignelementdependencyRequestCompound) GetSEzsignelementdependencyValueOk() (*string, bool) {
	if o == nil || IsNil(o.SEzsignelementdependencyValue) {
		return nil, false
	}
	return o.SEzsignelementdependencyValue, true
}

// HasSEzsignelementdependencyValue returns a boolean if a field has been set.
func (o *EzsignelementdependencyRequestCompound) HasSEzsignelementdependencyValue() bool {
	if o != nil && !IsNil(o.SEzsignelementdependencyValue) {
		return true
	}

	return false
}

// SetSEzsignelementdependencyValue gets a reference to the given string and assigns it to the SEzsignelementdependencyValue field.
func (o *EzsignelementdependencyRequestCompound) SetSEzsignelementdependencyValue(v string) {
	o.SEzsignelementdependencyValue = &v
}

func (o EzsignelementdependencyRequestCompound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignelementdependencyRequestCompound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PkiEzsignelementdependencyID) {
		toSerialize["pkiEzsignelementdependencyID"] = o.PkiEzsignelementdependencyID
	}
	if !IsNil(o.FkiEzsignformfieldIDValidation) {
		toSerialize["fkiEzsignformfieldIDValidation"] = o.FkiEzsignformfieldIDValidation
	}
	if !IsNil(o.FkiEzsignformfieldgroupIDValidation) {
		toSerialize["fkiEzsignformfieldgroupIDValidation"] = o.FkiEzsignformfieldgroupIDValidation
	}
	if !IsNil(o.SEzsignelementdependencyEzsignformfieldgrouplabel) {
		toSerialize["sEzsignelementdependencyEzsignformfieldgrouplabel"] = o.SEzsignelementdependencyEzsignformfieldgrouplabel
	}
	if !IsNil(o.SEzsignelementdependencyEzsignformfieldlabel) {
		toSerialize["sEzsignelementdependencyEzsignformfieldlabel"] = o.SEzsignelementdependencyEzsignformfieldlabel
	}
	toSerialize["eEzsignelementdependencyValidation"] = o.EEzsignelementdependencyValidation
	if !IsNil(o.BEzsignelementdependencySelected) {
		toSerialize["bEzsignelementdependencySelected"] = o.BEzsignelementdependencySelected
	}
	if !IsNil(o.EEzsignelementdependencyOperator) {
		toSerialize["eEzsignelementdependencyOperator"] = o.EEzsignelementdependencyOperator
	}
	if !IsNil(o.SEzsignelementdependencyValue) {
		toSerialize["sEzsignelementdependencyValue"] = o.SEzsignelementdependencyValue
	}
	return toSerialize, nil
}

func (o *EzsignelementdependencyRequestCompound) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"eEzsignelementdependencyValidation",
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

	varEzsignelementdependencyRequestCompound := _EzsignelementdependencyRequestCompound{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsignelementdependencyRequestCompound)

	if err != nil {
		return err
	}

	*o = EzsignelementdependencyRequestCompound(varEzsignelementdependencyRequestCompound)

	return err
}

type NullableEzsignelementdependencyRequestCompound struct {
	value *EzsignelementdependencyRequestCompound
	isSet bool
}

func (v NullableEzsignelementdependencyRequestCompound) Get() *EzsignelementdependencyRequestCompound {
	return v.value
}

func (v *NullableEzsignelementdependencyRequestCompound) Set(val *EzsignelementdependencyRequestCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignelementdependencyRequestCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignelementdependencyRequestCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignelementdependencyRequestCompound(val *EzsignelementdependencyRequestCompound) *NullableEzsignelementdependencyRequestCompound {
	return &NullableEzsignelementdependencyRequestCompound{value: val, isSet: true}
}

func (v NullableEzsignelementdependencyRequestCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignelementdependencyRequestCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


