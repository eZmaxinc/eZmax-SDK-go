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

// checks if the EzsigntemplateformfieldgroupsignerRequestCompound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigntemplateformfieldgroupsignerRequestCompound{}

// EzsigntemplateformfieldgroupsignerRequestCompound An Ezsigntemplateformfieldgroupsigner Object and children to create a complete structure
type EzsigntemplateformfieldgroupsignerRequestCompound struct {
	// The unique ID of the Ezsigntemplateformfieldgroupsigner
	PkiEzsigntemplateformfieldgroupsignerID *int32 `json:"pkiEzsigntemplateformfieldgroupsignerID,omitempty"`
	// The unique ID of the Ezsigntemplatesigner
	FkiEzsigntemplatesignerID int32 `json:"fkiEzsigntemplatesignerID"`
}

type _EzsigntemplateformfieldgroupsignerRequestCompound EzsigntemplateformfieldgroupsignerRequestCompound

// NewEzsigntemplateformfieldgroupsignerRequestCompound instantiates a new EzsigntemplateformfieldgroupsignerRequestCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigntemplateformfieldgroupsignerRequestCompound(fkiEzsigntemplatesignerID int32) *EzsigntemplateformfieldgroupsignerRequestCompound {
	this := EzsigntemplateformfieldgroupsignerRequestCompound{}
	this.FkiEzsigntemplatesignerID = fkiEzsigntemplatesignerID
	return &this
}

// NewEzsigntemplateformfieldgroupsignerRequestCompoundWithDefaults instantiates a new EzsigntemplateformfieldgroupsignerRequestCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigntemplateformfieldgroupsignerRequestCompoundWithDefaults() *EzsigntemplateformfieldgroupsignerRequestCompound {
	this := EzsigntemplateformfieldgroupsignerRequestCompound{}
	return &this
}

// GetPkiEzsigntemplateformfieldgroupsignerID returns the PkiEzsigntemplateformfieldgroupsignerID field value if set, zero value otherwise.
func (o *EzsigntemplateformfieldgroupsignerRequestCompound) GetPkiEzsigntemplateformfieldgroupsignerID() int32 {
	if o == nil || IsNil(o.PkiEzsigntemplateformfieldgroupsignerID) {
		var ret int32
		return ret
	}
	return *o.PkiEzsigntemplateformfieldgroupsignerID
}

// GetPkiEzsigntemplateformfieldgroupsignerIDOk returns a tuple with the PkiEzsigntemplateformfieldgroupsignerID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldgroupsignerRequestCompound) GetPkiEzsigntemplateformfieldgroupsignerIDOk() (*int32, bool) {
	if o == nil || IsNil(o.PkiEzsigntemplateformfieldgroupsignerID) {
		return nil, false
	}
	return o.PkiEzsigntemplateformfieldgroupsignerID, true
}

// HasPkiEzsigntemplateformfieldgroupsignerID returns a boolean if a field has been set.
func (o *EzsigntemplateformfieldgroupsignerRequestCompound) HasPkiEzsigntemplateformfieldgroupsignerID() bool {
	if o != nil && !IsNil(o.PkiEzsigntemplateformfieldgroupsignerID) {
		return true
	}

	return false
}

// SetPkiEzsigntemplateformfieldgroupsignerID gets a reference to the given int32 and assigns it to the PkiEzsigntemplateformfieldgroupsignerID field.
func (o *EzsigntemplateformfieldgroupsignerRequestCompound) SetPkiEzsigntemplateformfieldgroupsignerID(v int32) {
	o.PkiEzsigntemplateformfieldgroupsignerID = &v
}

// GetFkiEzsigntemplatesignerID returns the FkiEzsigntemplatesignerID field value
func (o *EzsigntemplateformfieldgroupsignerRequestCompound) GetFkiEzsigntemplatesignerID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzsigntemplatesignerID
}

// GetFkiEzsigntemplatesignerIDOk returns a tuple with the FkiEzsigntemplatesignerID field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldgroupsignerRequestCompound) GetFkiEzsigntemplatesignerIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiEzsigntemplatesignerID, true
}

// SetFkiEzsigntemplatesignerID sets field value
func (o *EzsigntemplateformfieldgroupsignerRequestCompound) SetFkiEzsigntemplatesignerID(v int32) {
	o.FkiEzsigntemplatesignerID = v
}

func (o EzsigntemplateformfieldgroupsignerRequestCompound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigntemplateformfieldgroupsignerRequestCompound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PkiEzsigntemplateformfieldgroupsignerID) {
		toSerialize["pkiEzsigntemplateformfieldgroupsignerID"] = o.PkiEzsigntemplateformfieldgroupsignerID
	}
	toSerialize["fkiEzsigntemplatesignerID"] = o.FkiEzsigntemplatesignerID
	return toSerialize, nil
}

func (o *EzsigntemplateformfieldgroupsignerRequestCompound) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fkiEzsigntemplatesignerID",
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

	varEzsigntemplateformfieldgroupsignerRequestCompound := _EzsigntemplateformfieldgroupsignerRequestCompound{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsigntemplateformfieldgroupsignerRequestCompound)

	if err != nil {
		return err
	}

	*o = EzsigntemplateformfieldgroupsignerRequestCompound(varEzsigntemplateformfieldgroupsignerRequestCompound)

	return err
}

type NullableEzsigntemplateformfieldgroupsignerRequestCompound struct {
	value *EzsigntemplateformfieldgroupsignerRequestCompound
	isSet bool
}

func (v NullableEzsigntemplateformfieldgroupsignerRequestCompound) Get() *EzsigntemplateformfieldgroupsignerRequestCompound {
	return v.value
}

func (v *NullableEzsigntemplateformfieldgroupsignerRequestCompound) Set(val *EzsigntemplateformfieldgroupsignerRequestCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigntemplateformfieldgroupsignerRequestCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigntemplateformfieldgroupsignerRequestCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigntemplateformfieldgroupsignerRequestCompound(val *EzsigntemplateformfieldgroupsignerRequestCompound) *NullableEzsigntemplateformfieldgroupsignerRequestCompound {
	return &NullableEzsigntemplateformfieldgroupsignerRequestCompound{value: val, isSet: true}
}

func (v NullableEzsigntemplateformfieldgroupsignerRequestCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigntemplateformfieldgroupsignerRequestCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


