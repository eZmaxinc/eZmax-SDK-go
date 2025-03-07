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

// checks if the EzsigntemplateelementdependencyRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigntemplateelementdependencyRequest{}

// EzsigntemplateelementdependencyRequest An Ezsigntemplateelementdependency Object
type EzsigntemplateelementdependencyRequest struct {
	// The unique ID of the Ezsigntemplateelementdependency
	PkiEzsigntemplateelementdependencyID *int32 `json:"pkiEzsigntemplateelementdependencyID,omitempty"`
	// The unique ID of the Ezsigntemplateformfield
	FkiEzsigntemplateformfieldIDValidation *int32 `json:"fkiEzsigntemplateformfieldIDValidation,omitempty"`
	// The unique ID of the Ezsigntemplateformfieldgroup
	FkiEzsigntemplateformfieldgroupIDValidation *int32 `json:"fkiEzsigntemplateformfieldgroupIDValidation,omitempty"`
	// The Label for the Ezsigntemplateformfieldgroup
	SEzsigntemplateelementdependencyEzsigntemplateformfieldgrouplabel *string `json:"sEzsigntemplateelementdependencyEzsigntemplateformfieldgrouplabel,omitempty"`
	// The Label for the Ezsigntemplateformfield
	SEzsigntemplateelementdependencyEzsigntemplateformfieldlabel *string `json:"sEzsigntemplateelementdependencyEzsigntemplateformfieldlabel,omitempty"`
	EEzsigntemplateelementdependencyValidation FieldEEzsigntemplateelementdependencyValidation `json:"eEzsigntemplateelementdependencyValidation"`
	// Whether if it's selected or not when using eEzsigntemplateelementdependencyValidation = Selected
	BEzsigntemplateelementdependencySelected *bool `json:"bEzsigntemplateelementdependencySelected,omitempty"`
	EEzsigntemplateelementdependencyOperator *FieldEEzsigntemplateelementdependencyOperator `json:"eEzsigntemplateelementdependencyOperator,omitempty"`
	// The value of the Ezsignelementdependency
	SEzsigntemplateelementdependencyValue *string `json:"sEzsigntemplateelementdependencyValue,omitempty" validate:"regexp=^.{0,50}$"`
}

type _EzsigntemplateelementdependencyRequest EzsigntemplateelementdependencyRequest

// NewEzsigntemplateelementdependencyRequest instantiates a new EzsigntemplateelementdependencyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigntemplateelementdependencyRequest(eEzsigntemplateelementdependencyValidation FieldEEzsigntemplateelementdependencyValidation) *EzsigntemplateelementdependencyRequest {
	this := EzsigntemplateelementdependencyRequest{}
	this.EEzsigntemplateelementdependencyValidation = eEzsigntemplateelementdependencyValidation
	return &this
}

// NewEzsigntemplateelementdependencyRequestWithDefaults instantiates a new EzsigntemplateelementdependencyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigntemplateelementdependencyRequestWithDefaults() *EzsigntemplateelementdependencyRequest {
	this := EzsigntemplateelementdependencyRequest{}
	return &this
}

// GetPkiEzsigntemplateelementdependencyID returns the PkiEzsigntemplateelementdependencyID field value if set, zero value otherwise.
func (o *EzsigntemplateelementdependencyRequest) GetPkiEzsigntemplateelementdependencyID() int32 {
	if o == nil || IsNil(o.PkiEzsigntemplateelementdependencyID) {
		var ret int32
		return ret
	}
	return *o.PkiEzsigntemplateelementdependencyID
}

// GetPkiEzsigntemplateelementdependencyIDOk returns a tuple with the PkiEzsigntemplateelementdependencyID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateelementdependencyRequest) GetPkiEzsigntemplateelementdependencyIDOk() (*int32, bool) {
	if o == nil || IsNil(o.PkiEzsigntemplateelementdependencyID) {
		return nil, false
	}
	return o.PkiEzsigntemplateelementdependencyID, true
}

// HasPkiEzsigntemplateelementdependencyID returns a boolean if a field has been set.
func (o *EzsigntemplateelementdependencyRequest) HasPkiEzsigntemplateelementdependencyID() bool {
	if o != nil && !IsNil(o.PkiEzsigntemplateelementdependencyID) {
		return true
	}

	return false
}

// SetPkiEzsigntemplateelementdependencyID gets a reference to the given int32 and assigns it to the PkiEzsigntemplateelementdependencyID field.
func (o *EzsigntemplateelementdependencyRequest) SetPkiEzsigntemplateelementdependencyID(v int32) {
	o.PkiEzsigntemplateelementdependencyID = &v
}

// GetFkiEzsigntemplateformfieldIDValidation returns the FkiEzsigntemplateformfieldIDValidation field value if set, zero value otherwise.
func (o *EzsigntemplateelementdependencyRequest) GetFkiEzsigntemplateformfieldIDValidation() int32 {
	if o == nil || IsNil(o.FkiEzsigntemplateformfieldIDValidation) {
		var ret int32
		return ret
	}
	return *o.FkiEzsigntemplateformfieldIDValidation
}

// GetFkiEzsigntemplateformfieldIDValidationOk returns a tuple with the FkiEzsigntemplateformfieldIDValidation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateelementdependencyRequest) GetFkiEzsigntemplateformfieldIDValidationOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiEzsigntemplateformfieldIDValidation) {
		return nil, false
	}
	return o.FkiEzsigntemplateformfieldIDValidation, true
}

// HasFkiEzsigntemplateformfieldIDValidation returns a boolean if a field has been set.
func (o *EzsigntemplateelementdependencyRequest) HasFkiEzsigntemplateformfieldIDValidation() bool {
	if o != nil && !IsNil(o.FkiEzsigntemplateformfieldIDValidation) {
		return true
	}

	return false
}

// SetFkiEzsigntemplateformfieldIDValidation gets a reference to the given int32 and assigns it to the FkiEzsigntemplateformfieldIDValidation field.
func (o *EzsigntemplateelementdependencyRequest) SetFkiEzsigntemplateformfieldIDValidation(v int32) {
	o.FkiEzsigntemplateformfieldIDValidation = &v
}

// GetFkiEzsigntemplateformfieldgroupIDValidation returns the FkiEzsigntemplateformfieldgroupIDValidation field value if set, zero value otherwise.
func (o *EzsigntemplateelementdependencyRequest) GetFkiEzsigntemplateformfieldgroupIDValidation() int32 {
	if o == nil || IsNil(o.FkiEzsigntemplateformfieldgroupIDValidation) {
		var ret int32
		return ret
	}
	return *o.FkiEzsigntemplateformfieldgroupIDValidation
}

// GetFkiEzsigntemplateformfieldgroupIDValidationOk returns a tuple with the FkiEzsigntemplateformfieldgroupIDValidation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateelementdependencyRequest) GetFkiEzsigntemplateformfieldgroupIDValidationOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiEzsigntemplateformfieldgroupIDValidation) {
		return nil, false
	}
	return o.FkiEzsigntemplateformfieldgroupIDValidation, true
}

// HasFkiEzsigntemplateformfieldgroupIDValidation returns a boolean if a field has been set.
func (o *EzsigntemplateelementdependencyRequest) HasFkiEzsigntemplateformfieldgroupIDValidation() bool {
	if o != nil && !IsNil(o.FkiEzsigntemplateformfieldgroupIDValidation) {
		return true
	}

	return false
}

// SetFkiEzsigntemplateformfieldgroupIDValidation gets a reference to the given int32 and assigns it to the FkiEzsigntemplateformfieldgroupIDValidation field.
func (o *EzsigntemplateelementdependencyRequest) SetFkiEzsigntemplateformfieldgroupIDValidation(v int32) {
	o.FkiEzsigntemplateformfieldgroupIDValidation = &v
}

// GetSEzsigntemplateelementdependencyEzsigntemplateformfieldgrouplabel returns the SEzsigntemplateelementdependencyEzsigntemplateformfieldgrouplabel field value if set, zero value otherwise.
func (o *EzsigntemplateelementdependencyRequest) GetSEzsigntemplateelementdependencyEzsigntemplateformfieldgrouplabel() string {
	if o == nil || IsNil(o.SEzsigntemplateelementdependencyEzsigntemplateformfieldgrouplabel) {
		var ret string
		return ret
	}
	return *o.SEzsigntemplateelementdependencyEzsigntemplateformfieldgrouplabel
}

// GetSEzsigntemplateelementdependencyEzsigntemplateformfieldgrouplabelOk returns a tuple with the SEzsigntemplateelementdependencyEzsigntemplateformfieldgrouplabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateelementdependencyRequest) GetSEzsigntemplateelementdependencyEzsigntemplateformfieldgrouplabelOk() (*string, bool) {
	if o == nil || IsNil(o.SEzsigntemplateelementdependencyEzsigntemplateformfieldgrouplabel) {
		return nil, false
	}
	return o.SEzsigntemplateelementdependencyEzsigntemplateformfieldgrouplabel, true
}

// HasSEzsigntemplateelementdependencyEzsigntemplateformfieldgrouplabel returns a boolean if a field has been set.
func (o *EzsigntemplateelementdependencyRequest) HasSEzsigntemplateelementdependencyEzsigntemplateformfieldgrouplabel() bool {
	if o != nil && !IsNil(o.SEzsigntemplateelementdependencyEzsigntemplateformfieldgrouplabel) {
		return true
	}

	return false
}

// SetSEzsigntemplateelementdependencyEzsigntemplateformfieldgrouplabel gets a reference to the given string and assigns it to the SEzsigntemplateelementdependencyEzsigntemplateformfieldgrouplabel field.
func (o *EzsigntemplateelementdependencyRequest) SetSEzsigntemplateelementdependencyEzsigntemplateformfieldgrouplabel(v string) {
	o.SEzsigntemplateelementdependencyEzsigntemplateformfieldgrouplabel = &v
}

// GetSEzsigntemplateelementdependencyEzsigntemplateformfieldlabel returns the SEzsigntemplateelementdependencyEzsigntemplateformfieldlabel field value if set, zero value otherwise.
func (o *EzsigntemplateelementdependencyRequest) GetSEzsigntemplateelementdependencyEzsigntemplateformfieldlabel() string {
	if o == nil || IsNil(o.SEzsigntemplateelementdependencyEzsigntemplateformfieldlabel) {
		var ret string
		return ret
	}
	return *o.SEzsigntemplateelementdependencyEzsigntemplateformfieldlabel
}

// GetSEzsigntemplateelementdependencyEzsigntemplateformfieldlabelOk returns a tuple with the SEzsigntemplateelementdependencyEzsigntemplateformfieldlabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateelementdependencyRequest) GetSEzsigntemplateelementdependencyEzsigntemplateformfieldlabelOk() (*string, bool) {
	if o == nil || IsNil(o.SEzsigntemplateelementdependencyEzsigntemplateformfieldlabel) {
		return nil, false
	}
	return o.SEzsigntemplateelementdependencyEzsigntemplateformfieldlabel, true
}

// HasSEzsigntemplateelementdependencyEzsigntemplateformfieldlabel returns a boolean if a field has been set.
func (o *EzsigntemplateelementdependencyRequest) HasSEzsigntemplateelementdependencyEzsigntemplateformfieldlabel() bool {
	if o != nil && !IsNil(o.SEzsigntemplateelementdependencyEzsigntemplateformfieldlabel) {
		return true
	}

	return false
}

// SetSEzsigntemplateelementdependencyEzsigntemplateformfieldlabel gets a reference to the given string and assigns it to the SEzsigntemplateelementdependencyEzsigntemplateformfieldlabel field.
func (o *EzsigntemplateelementdependencyRequest) SetSEzsigntemplateelementdependencyEzsigntemplateformfieldlabel(v string) {
	o.SEzsigntemplateelementdependencyEzsigntemplateformfieldlabel = &v
}

// GetEEzsigntemplateelementdependencyValidation returns the EEzsigntemplateelementdependencyValidation field value
func (o *EzsigntemplateelementdependencyRequest) GetEEzsigntemplateelementdependencyValidation() FieldEEzsigntemplateelementdependencyValidation {
	if o == nil {
		var ret FieldEEzsigntemplateelementdependencyValidation
		return ret
	}

	return o.EEzsigntemplateelementdependencyValidation
}

// GetEEzsigntemplateelementdependencyValidationOk returns a tuple with the EEzsigntemplateelementdependencyValidation field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateelementdependencyRequest) GetEEzsigntemplateelementdependencyValidationOk() (*FieldEEzsigntemplateelementdependencyValidation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EEzsigntemplateelementdependencyValidation, true
}

// SetEEzsigntemplateelementdependencyValidation sets field value
func (o *EzsigntemplateelementdependencyRequest) SetEEzsigntemplateelementdependencyValidation(v FieldEEzsigntemplateelementdependencyValidation) {
	o.EEzsigntemplateelementdependencyValidation = v
}

// GetBEzsigntemplateelementdependencySelected returns the BEzsigntemplateelementdependencySelected field value if set, zero value otherwise.
func (o *EzsigntemplateelementdependencyRequest) GetBEzsigntemplateelementdependencySelected() bool {
	if o == nil || IsNil(o.BEzsigntemplateelementdependencySelected) {
		var ret bool
		return ret
	}
	return *o.BEzsigntemplateelementdependencySelected
}

// GetBEzsigntemplateelementdependencySelectedOk returns a tuple with the BEzsigntemplateelementdependencySelected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateelementdependencyRequest) GetBEzsigntemplateelementdependencySelectedOk() (*bool, bool) {
	if o == nil || IsNil(o.BEzsigntemplateelementdependencySelected) {
		return nil, false
	}
	return o.BEzsigntemplateelementdependencySelected, true
}

// HasBEzsigntemplateelementdependencySelected returns a boolean if a field has been set.
func (o *EzsigntemplateelementdependencyRequest) HasBEzsigntemplateelementdependencySelected() bool {
	if o != nil && !IsNil(o.BEzsigntemplateelementdependencySelected) {
		return true
	}

	return false
}

// SetBEzsigntemplateelementdependencySelected gets a reference to the given bool and assigns it to the BEzsigntemplateelementdependencySelected field.
func (o *EzsigntemplateelementdependencyRequest) SetBEzsigntemplateelementdependencySelected(v bool) {
	o.BEzsigntemplateelementdependencySelected = &v
}

// GetEEzsigntemplateelementdependencyOperator returns the EEzsigntemplateelementdependencyOperator field value if set, zero value otherwise.
func (o *EzsigntemplateelementdependencyRequest) GetEEzsigntemplateelementdependencyOperator() FieldEEzsigntemplateelementdependencyOperator {
	if o == nil || IsNil(o.EEzsigntemplateelementdependencyOperator) {
		var ret FieldEEzsigntemplateelementdependencyOperator
		return ret
	}
	return *o.EEzsigntemplateelementdependencyOperator
}

// GetEEzsigntemplateelementdependencyOperatorOk returns a tuple with the EEzsigntemplateelementdependencyOperator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateelementdependencyRequest) GetEEzsigntemplateelementdependencyOperatorOk() (*FieldEEzsigntemplateelementdependencyOperator, bool) {
	if o == nil || IsNil(o.EEzsigntemplateelementdependencyOperator) {
		return nil, false
	}
	return o.EEzsigntemplateelementdependencyOperator, true
}

// HasEEzsigntemplateelementdependencyOperator returns a boolean if a field has been set.
func (o *EzsigntemplateelementdependencyRequest) HasEEzsigntemplateelementdependencyOperator() bool {
	if o != nil && !IsNil(o.EEzsigntemplateelementdependencyOperator) {
		return true
	}

	return false
}

// SetEEzsigntemplateelementdependencyOperator gets a reference to the given FieldEEzsigntemplateelementdependencyOperator and assigns it to the EEzsigntemplateelementdependencyOperator field.
func (o *EzsigntemplateelementdependencyRequest) SetEEzsigntemplateelementdependencyOperator(v FieldEEzsigntemplateelementdependencyOperator) {
	o.EEzsigntemplateelementdependencyOperator = &v
}

// GetSEzsigntemplateelementdependencyValue returns the SEzsigntemplateelementdependencyValue field value if set, zero value otherwise.
func (o *EzsigntemplateelementdependencyRequest) GetSEzsigntemplateelementdependencyValue() string {
	if o == nil || IsNil(o.SEzsigntemplateelementdependencyValue) {
		var ret string
		return ret
	}
	return *o.SEzsigntemplateelementdependencyValue
}

// GetSEzsigntemplateelementdependencyValueOk returns a tuple with the SEzsigntemplateelementdependencyValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateelementdependencyRequest) GetSEzsigntemplateelementdependencyValueOk() (*string, bool) {
	if o == nil || IsNil(o.SEzsigntemplateelementdependencyValue) {
		return nil, false
	}
	return o.SEzsigntemplateelementdependencyValue, true
}

// HasSEzsigntemplateelementdependencyValue returns a boolean if a field has been set.
func (o *EzsigntemplateelementdependencyRequest) HasSEzsigntemplateelementdependencyValue() bool {
	if o != nil && !IsNil(o.SEzsigntemplateelementdependencyValue) {
		return true
	}

	return false
}

// SetSEzsigntemplateelementdependencyValue gets a reference to the given string and assigns it to the SEzsigntemplateelementdependencyValue field.
func (o *EzsigntemplateelementdependencyRequest) SetSEzsigntemplateelementdependencyValue(v string) {
	o.SEzsigntemplateelementdependencyValue = &v
}

func (o EzsigntemplateelementdependencyRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigntemplateelementdependencyRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PkiEzsigntemplateelementdependencyID) {
		toSerialize["pkiEzsigntemplateelementdependencyID"] = o.PkiEzsigntemplateelementdependencyID
	}
	if !IsNil(o.FkiEzsigntemplateformfieldIDValidation) {
		toSerialize["fkiEzsigntemplateformfieldIDValidation"] = o.FkiEzsigntemplateformfieldIDValidation
	}
	if !IsNil(o.FkiEzsigntemplateformfieldgroupIDValidation) {
		toSerialize["fkiEzsigntemplateformfieldgroupIDValidation"] = o.FkiEzsigntemplateformfieldgroupIDValidation
	}
	if !IsNil(o.SEzsigntemplateelementdependencyEzsigntemplateformfieldgrouplabel) {
		toSerialize["sEzsigntemplateelementdependencyEzsigntemplateformfieldgrouplabel"] = o.SEzsigntemplateelementdependencyEzsigntemplateformfieldgrouplabel
	}
	if !IsNil(o.SEzsigntemplateelementdependencyEzsigntemplateformfieldlabel) {
		toSerialize["sEzsigntemplateelementdependencyEzsigntemplateformfieldlabel"] = o.SEzsigntemplateelementdependencyEzsigntemplateformfieldlabel
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

func (o *EzsigntemplateelementdependencyRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varEzsigntemplateelementdependencyRequest := _EzsigntemplateelementdependencyRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsigntemplateelementdependencyRequest)

	if err != nil {
		return err
	}

	*o = EzsigntemplateelementdependencyRequest(varEzsigntemplateelementdependencyRequest)

	return err
}

type NullableEzsigntemplateelementdependencyRequest struct {
	value *EzsigntemplateelementdependencyRequest
	isSet bool
}

func (v NullableEzsigntemplateelementdependencyRequest) Get() *EzsigntemplateelementdependencyRequest {
	return v.value
}

func (v *NullableEzsigntemplateelementdependencyRequest) Set(val *EzsigntemplateelementdependencyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigntemplateelementdependencyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigntemplateelementdependencyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigntemplateelementdependencyRequest(val *EzsigntemplateelementdependencyRequest) *NullableEzsigntemplateelementdependencyRequest {
	return &NullableEzsigntemplateelementdependencyRequest{value: val, isSet: true}
}

func (v NullableEzsigntemplateelementdependencyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigntemplateelementdependencyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


