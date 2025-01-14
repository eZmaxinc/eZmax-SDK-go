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

// checks if the EzsigntemplateelementdependencyResponseCompound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigntemplateelementdependencyResponseCompound{}

// EzsigntemplateelementdependencyResponseCompound An Ezsigntemplateelementdependency Object and children to create a complete structure
type EzsigntemplateelementdependencyResponseCompound struct {
	// The unique ID of the Ezsigntemplateelementdependency
	PkiEzsigntemplateelementdependencyID int32 `json:"pkiEzsigntemplateelementdependencyID"`
	// The unique ID of the Ezsigntemplateformfield
	FkiEzsigntemplateformfieldID *int32 `json:"fkiEzsigntemplateformfieldID,omitempty"`
	// The unique ID of the Ezsigntemplatesignature
	FkiEzsigntemplatesignatureID *int32 `json:"fkiEzsigntemplatesignatureID,omitempty"`
	// The unique ID of the Ezsigntemplateformfield
	FkiEzsigntemplateformfieldIDValidation *int32 `json:"fkiEzsigntemplateformfieldIDValidation,omitempty"`
	// The unique ID of the Ezsigntemplateformfieldgroup
	FkiEzsigntemplateformfieldgroupIDValidation *int32 `json:"fkiEzsigntemplateformfieldgroupIDValidation,omitempty"`
	EEzsigntemplateelementdependencyValidation FieldEEzsigntemplateelementdependencyValidation `json:"eEzsigntemplateelementdependencyValidation"`
	// Whether if it's selected or not when using eEzsigntemplateelementdependencyValidation = Selected
	BEzsigntemplateelementdependencySelected *bool `json:"bEzsigntemplateelementdependencySelected,omitempty"`
	EEzsigntemplateelementdependencyOperator *FieldEEzsigntemplateelementdependencyOperator `json:"eEzsigntemplateelementdependencyOperator,omitempty"`
	// The value of the Ezsignelementdependency
	SEzsigntemplateelementdependencyValue *string `json:"sEzsigntemplateelementdependencyValue,omitempty" validate:"regexp=^.{0,50}$"`
}

type _EzsigntemplateelementdependencyResponseCompound EzsigntemplateelementdependencyResponseCompound

// NewEzsigntemplateelementdependencyResponseCompound instantiates a new EzsigntemplateelementdependencyResponseCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigntemplateelementdependencyResponseCompound(pkiEzsigntemplateelementdependencyID int32, eEzsigntemplateelementdependencyValidation FieldEEzsigntemplateelementdependencyValidation) *EzsigntemplateelementdependencyResponseCompound {
	this := EzsigntemplateelementdependencyResponseCompound{}
	this.PkiEzsigntemplateelementdependencyID = pkiEzsigntemplateelementdependencyID
	this.EEzsigntemplateelementdependencyValidation = eEzsigntemplateelementdependencyValidation
	return &this
}

// NewEzsigntemplateelementdependencyResponseCompoundWithDefaults instantiates a new EzsigntemplateelementdependencyResponseCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigntemplateelementdependencyResponseCompoundWithDefaults() *EzsigntemplateelementdependencyResponseCompound {
	this := EzsigntemplateelementdependencyResponseCompound{}
	return &this
}

// GetPkiEzsigntemplateelementdependencyID returns the PkiEzsigntemplateelementdependencyID field value
func (o *EzsigntemplateelementdependencyResponseCompound) GetPkiEzsigntemplateelementdependencyID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiEzsigntemplateelementdependencyID
}

// GetPkiEzsigntemplateelementdependencyIDOk returns a tuple with the PkiEzsigntemplateelementdependencyID field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateelementdependencyResponseCompound) GetPkiEzsigntemplateelementdependencyIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiEzsigntemplateelementdependencyID, true
}

// SetPkiEzsigntemplateelementdependencyID sets field value
func (o *EzsigntemplateelementdependencyResponseCompound) SetPkiEzsigntemplateelementdependencyID(v int32) {
	o.PkiEzsigntemplateelementdependencyID = v
}

// GetFkiEzsigntemplateformfieldID returns the FkiEzsigntemplateformfieldID field value if set, zero value otherwise.
func (o *EzsigntemplateelementdependencyResponseCompound) GetFkiEzsigntemplateformfieldID() int32 {
	if o == nil || IsNil(o.FkiEzsigntemplateformfieldID) {
		var ret int32
		return ret
	}
	return *o.FkiEzsigntemplateformfieldID
}

// GetFkiEzsigntemplateformfieldIDOk returns a tuple with the FkiEzsigntemplateformfieldID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateelementdependencyResponseCompound) GetFkiEzsigntemplateformfieldIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiEzsigntemplateformfieldID) {
		return nil, false
	}
	return o.FkiEzsigntemplateformfieldID, true
}

// HasFkiEzsigntemplateformfieldID returns a boolean if a field has been set.
func (o *EzsigntemplateelementdependencyResponseCompound) HasFkiEzsigntemplateformfieldID() bool {
	if o != nil && !IsNil(o.FkiEzsigntemplateformfieldID) {
		return true
	}

	return false
}

// SetFkiEzsigntemplateformfieldID gets a reference to the given int32 and assigns it to the FkiEzsigntemplateformfieldID field.
func (o *EzsigntemplateelementdependencyResponseCompound) SetFkiEzsigntemplateformfieldID(v int32) {
	o.FkiEzsigntemplateformfieldID = &v
}

// GetFkiEzsigntemplatesignatureID returns the FkiEzsigntemplatesignatureID field value if set, zero value otherwise.
func (o *EzsigntemplateelementdependencyResponseCompound) GetFkiEzsigntemplatesignatureID() int32 {
	if o == nil || IsNil(o.FkiEzsigntemplatesignatureID) {
		var ret int32
		return ret
	}
	return *o.FkiEzsigntemplatesignatureID
}

// GetFkiEzsigntemplatesignatureIDOk returns a tuple with the FkiEzsigntemplatesignatureID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateelementdependencyResponseCompound) GetFkiEzsigntemplatesignatureIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiEzsigntemplatesignatureID) {
		return nil, false
	}
	return o.FkiEzsigntemplatesignatureID, true
}

// HasFkiEzsigntemplatesignatureID returns a boolean if a field has been set.
func (o *EzsigntemplateelementdependencyResponseCompound) HasFkiEzsigntemplatesignatureID() bool {
	if o != nil && !IsNil(o.FkiEzsigntemplatesignatureID) {
		return true
	}

	return false
}

// SetFkiEzsigntemplatesignatureID gets a reference to the given int32 and assigns it to the FkiEzsigntemplatesignatureID field.
func (o *EzsigntemplateelementdependencyResponseCompound) SetFkiEzsigntemplatesignatureID(v int32) {
	o.FkiEzsigntemplatesignatureID = &v
}

// GetFkiEzsigntemplateformfieldIDValidation returns the FkiEzsigntemplateformfieldIDValidation field value if set, zero value otherwise.
func (o *EzsigntemplateelementdependencyResponseCompound) GetFkiEzsigntemplateformfieldIDValidation() int32 {
	if o == nil || IsNil(o.FkiEzsigntemplateformfieldIDValidation) {
		var ret int32
		return ret
	}
	return *o.FkiEzsigntemplateformfieldIDValidation
}

// GetFkiEzsigntemplateformfieldIDValidationOk returns a tuple with the FkiEzsigntemplateformfieldIDValidation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateelementdependencyResponseCompound) GetFkiEzsigntemplateformfieldIDValidationOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiEzsigntemplateformfieldIDValidation) {
		return nil, false
	}
	return o.FkiEzsigntemplateformfieldIDValidation, true
}

// HasFkiEzsigntemplateformfieldIDValidation returns a boolean if a field has been set.
func (o *EzsigntemplateelementdependencyResponseCompound) HasFkiEzsigntemplateformfieldIDValidation() bool {
	if o != nil && !IsNil(o.FkiEzsigntemplateformfieldIDValidation) {
		return true
	}

	return false
}

// SetFkiEzsigntemplateformfieldIDValidation gets a reference to the given int32 and assigns it to the FkiEzsigntemplateformfieldIDValidation field.
func (o *EzsigntemplateelementdependencyResponseCompound) SetFkiEzsigntemplateformfieldIDValidation(v int32) {
	o.FkiEzsigntemplateformfieldIDValidation = &v
}

// GetFkiEzsigntemplateformfieldgroupIDValidation returns the FkiEzsigntemplateformfieldgroupIDValidation field value if set, zero value otherwise.
func (o *EzsigntemplateelementdependencyResponseCompound) GetFkiEzsigntemplateformfieldgroupIDValidation() int32 {
	if o == nil || IsNil(o.FkiEzsigntemplateformfieldgroupIDValidation) {
		var ret int32
		return ret
	}
	return *o.FkiEzsigntemplateformfieldgroupIDValidation
}

// GetFkiEzsigntemplateformfieldgroupIDValidationOk returns a tuple with the FkiEzsigntemplateformfieldgroupIDValidation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateelementdependencyResponseCompound) GetFkiEzsigntemplateformfieldgroupIDValidationOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiEzsigntemplateformfieldgroupIDValidation) {
		return nil, false
	}
	return o.FkiEzsigntemplateformfieldgroupIDValidation, true
}

// HasFkiEzsigntemplateformfieldgroupIDValidation returns a boolean if a field has been set.
func (o *EzsigntemplateelementdependencyResponseCompound) HasFkiEzsigntemplateformfieldgroupIDValidation() bool {
	if o != nil && !IsNil(o.FkiEzsigntemplateformfieldgroupIDValidation) {
		return true
	}

	return false
}

// SetFkiEzsigntemplateformfieldgroupIDValidation gets a reference to the given int32 and assigns it to the FkiEzsigntemplateformfieldgroupIDValidation field.
func (o *EzsigntemplateelementdependencyResponseCompound) SetFkiEzsigntemplateformfieldgroupIDValidation(v int32) {
	o.FkiEzsigntemplateformfieldgroupIDValidation = &v
}

// GetEEzsigntemplateelementdependencyValidation returns the EEzsigntemplateelementdependencyValidation field value
func (o *EzsigntemplateelementdependencyResponseCompound) GetEEzsigntemplateelementdependencyValidation() FieldEEzsigntemplateelementdependencyValidation {
	if o == nil {
		var ret FieldEEzsigntemplateelementdependencyValidation
		return ret
	}

	return o.EEzsigntemplateelementdependencyValidation
}

// GetEEzsigntemplateelementdependencyValidationOk returns a tuple with the EEzsigntemplateelementdependencyValidation field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateelementdependencyResponseCompound) GetEEzsigntemplateelementdependencyValidationOk() (*FieldEEzsigntemplateelementdependencyValidation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EEzsigntemplateelementdependencyValidation, true
}

// SetEEzsigntemplateelementdependencyValidation sets field value
func (o *EzsigntemplateelementdependencyResponseCompound) SetEEzsigntemplateelementdependencyValidation(v FieldEEzsigntemplateelementdependencyValidation) {
	o.EEzsigntemplateelementdependencyValidation = v
}

// GetBEzsigntemplateelementdependencySelected returns the BEzsigntemplateelementdependencySelected field value if set, zero value otherwise.
func (o *EzsigntemplateelementdependencyResponseCompound) GetBEzsigntemplateelementdependencySelected() bool {
	if o == nil || IsNil(o.BEzsigntemplateelementdependencySelected) {
		var ret bool
		return ret
	}
	return *o.BEzsigntemplateelementdependencySelected
}

// GetBEzsigntemplateelementdependencySelectedOk returns a tuple with the BEzsigntemplateelementdependencySelected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateelementdependencyResponseCompound) GetBEzsigntemplateelementdependencySelectedOk() (*bool, bool) {
	if o == nil || IsNil(o.BEzsigntemplateelementdependencySelected) {
		return nil, false
	}
	return o.BEzsigntemplateelementdependencySelected, true
}

// HasBEzsigntemplateelementdependencySelected returns a boolean if a field has been set.
func (o *EzsigntemplateelementdependencyResponseCompound) HasBEzsigntemplateelementdependencySelected() bool {
	if o != nil && !IsNil(o.BEzsigntemplateelementdependencySelected) {
		return true
	}

	return false
}

// SetBEzsigntemplateelementdependencySelected gets a reference to the given bool and assigns it to the BEzsigntemplateelementdependencySelected field.
func (o *EzsigntemplateelementdependencyResponseCompound) SetBEzsigntemplateelementdependencySelected(v bool) {
	o.BEzsigntemplateelementdependencySelected = &v
}

// GetEEzsigntemplateelementdependencyOperator returns the EEzsigntemplateelementdependencyOperator field value if set, zero value otherwise.
func (o *EzsigntemplateelementdependencyResponseCompound) GetEEzsigntemplateelementdependencyOperator() FieldEEzsigntemplateelementdependencyOperator {
	if o == nil || IsNil(o.EEzsigntemplateelementdependencyOperator) {
		var ret FieldEEzsigntemplateelementdependencyOperator
		return ret
	}
	return *o.EEzsigntemplateelementdependencyOperator
}

// GetEEzsigntemplateelementdependencyOperatorOk returns a tuple with the EEzsigntemplateelementdependencyOperator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateelementdependencyResponseCompound) GetEEzsigntemplateelementdependencyOperatorOk() (*FieldEEzsigntemplateelementdependencyOperator, bool) {
	if o == nil || IsNil(o.EEzsigntemplateelementdependencyOperator) {
		return nil, false
	}
	return o.EEzsigntemplateelementdependencyOperator, true
}

// HasEEzsigntemplateelementdependencyOperator returns a boolean if a field has been set.
func (o *EzsigntemplateelementdependencyResponseCompound) HasEEzsigntemplateelementdependencyOperator() bool {
	if o != nil && !IsNil(o.EEzsigntemplateelementdependencyOperator) {
		return true
	}

	return false
}

// SetEEzsigntemplateelementdependencyOperator gets a reference to the given FieldEEzsigntemplateelementdependencyOperator and assigns it to the EEzsigntemplateelementdependencyOperator field.
func (o *EzsigntemplateelementdependencyResponseCompound) SetEEzsigntemplateelementdependencyOperator(v FieldEEzsigntemplateelementdependencyOperator) {
	o.EEzsigntemplateelementdependencyOperator = &v
}

// GetSEzsigntemplateelementdependencyValue returns the SEzsigntemplateelementdependencyValue field value if set, zero value otherwise.
func (o *EzsigntemplateelementdependencyResponseCompound) GetSEzsigntemplateelementdependencyValue() string {
	if o == nil || IsNil(o.SEzsigntemplateelementdependencyValue) {
		var ret string
		return ret
	}
	return *o.SEzsigntemplateelementdependencyValue
}

// GetSEzsigntemplateelementdependencyValueOk returns a tuple with the SEzsigntemplateelementdependencyValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateelementdependencyResponseCompound) GetSEzsigntemplateelementdependencyValueOk() (*string, bool) {
	if o == nil || IsNil(o.SEzsigntemplateelementdependencyValue) {
		return nil, false
	}
	return o.SEzsigntemplateelementdependencyValue, true
}

// HasSEzsigntemplateelementdependencyValue returns a boolean if a field has been set.
func (o *EzsigntemplateelementdependencyResponseCompound) HasSEzsigntemplateelementdependencyValue() bool {
	if o != nil && !IsNil(o.SEzsigntemplateelementdependencyValue) {
		return true
	}

	return false
}

// SetSEzsigntemplateelementdependencyValue gets a reference to the given string and assigns it to the SEzsigntemplateelementdependencyValue field.
func (o *EzsigntemplateelementdependencyResponseCompound) SetSEzsigntemplateelementdependencyValue(v string) {
	o.SEzsigntemplateelementdependencyValue = &v
}

func (o EzsigntemplateelementdependencyResponseCompound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigntemplateelementdependencyResponseCompound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiEzsigntemplateelementdependencyID"] = o.PkiEzsigntemplateelementdependencyID
	if !IsNil(o.FkiEzsigntemplateformfieldID) {
		toSerialize["fkiEzsigntemplateformfieldID"] = o.FkiEzsigntemplateformfieldID
	}
	if !IsNil(o.FkiEzsigntemplatesignatureID) {
		toSerialize["fkiEzsigntemplatesignatureID"] = o.FkiEzsigntemplatesignatureID
	}
	if !IsNil(o.FkiEzsigntemplateformfieldIDValidation) {
		toSerialize["fkiEzsigntemplateformfieldIDValidation"] = o.FkiEzsigntemplateformfieldIDValidation
	}
	if !IsNil(o.FkiEzsigntemplateformfieldgroupIDValidation) {
		toSerialize["fkiEzsigntemplateformfieldgroupIDValidation"] = o.FkiEzsigntemplateformfieldgroupIDValidation
	}
	toSerialize["eEzsigntemplateelementdependencyValidation"] = o.EEzsigntemplateelementdependencyValidation
	if !IsNil(o.BEzsigntemplateelementdependencySelected) {
		toSerialize["bEzsigntemplateelementdependencySelected"] = o.BEzsigntemplateelementdependencySelected
	}
	if !IsNil(o.EEzsigntemplateelementdependencyOperator) {
		toSerialize["eEzsigntemplateelementdependencyOperator"] = o.EEzsigntemplateelementdependencyOperator
	}
	if !IsNil(o.SEzsigntemplateelementdependencyValue) {
		toSerialize["sEzsigntemplateelementdependencyValue"] = o.SEzsigntemplateelementdependencyValue
	}
	return toSerialize, nil
}

func (o *EzsigntemplateelementdependencyResponseCompound) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiEzsigntemplateelementdependencyID",
		"eEzsigntemplateelementdependencyValidation",
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

	varEzsigntemplateelementdependencyResponseCompound := _EzsigntemplateelementdependencyResponseCompound{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsigntemplateelementdependencyResponseCompound)

	if err != nil {
		return err
	}

	*o = EzsigntemplateelementdependencyResponseCompound(varEzsigntemplateelementdependencyResponseCompound)

	return err
}

type NullableEzsigntemplateelementdependencyResponseCompound struct {
	value *EzsigntemplateelementdependencyResponseCompound
	isSet bool
}

func (v NullableEzsigntemplateelementdependencyResponseCompound) Get() *EzsigntemplateelementdependencyResponseCompound {
	return v.value
}

func (v *NullableEzsigntemplateelementdependencyResponseCompound) Set(val *EzsigntemplateelementdependencyResponseCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigntemplateelementdependencyResponseCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigntemplateelementdependencyResponseCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigntemplateelementdependencyResponseCompound(val *EzsigntemplateelementdependencyResponseCompound) *NullableEzsigntemplateelementdependencyResponseCompound {
	return &NullableEzsigntemplateelementdependencyResponseCompound{value: val, isSet: true}
}

func (v NullableEzsigntemplateelementdependencyResponseCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigntemplateelementdependencyResponseCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


