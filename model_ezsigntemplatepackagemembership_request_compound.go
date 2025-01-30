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

// checks if the EzsigntemplatepackagemembershipRequestCompound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigntemplatepackagemembershipRequestCompound{}

// EzsigntemplatepackagemembershipRequestCompound A Ezsigntemplatepackagemembership Object and children
type EzsigntemplatepackagemembershipRequestCompound struct {
	// The unique ID of the Ezsigntemplatepackagemembership
	PkiEzsigntemplatepackagemembershipID *int32 `json:"pkiEzsigntemplatepackagemembershipID,omitempty"`
	// The unique ID of the Ezsigntemplatepackage
	FkiEzsigntemplatepackageID int32 `json:"fkiEzsigntemplatepackageID"`
	// The unique ID of the Ezsigntemplate
	FkiEzsigntemplateID int32 `json:"fkiEzsigntemplateID"`
}

type _EzsigntemplatepackagemembershipRequestCompound EzsigntemplatepackagemembershipRequestCompound

// NewEzsigntemplatepackagemembershipRequestCompound instantiates a new EzsigntemplatepackagemembershipRequestCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigntemplatepackagemembershipRequestCompound(fkiEzsigntemplatepackageID int32, fkiEzsigntemplateID int32) *EzsigntemplatepackagemembershipRequestCompound {
	this := EzsigntemplatepackagemembershipRequestCompound{}
	this.FkiEzsigntemplatepackageID = fkiEzsigntemplatepackageID
	this.FkiEzsigntemplateID = fkiEzsigntemplateID
	return &this
}

// NewEzsigntemplatepackagemembershipRequestCompoundWithDefaults instantiates a new EzsigntemplatepackagemembershipRequestCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigntemplatepackagemembershipRequestCompoundWithDefaults() *EzsigntemplatepackagemembershipRequestCompound {
	this := EzsigntemplatepackagemembershipRequestCompound{}
	return &this
}

// GetPkiEzsigntemplatepackagemembershipID returns the PkiEzsigntemplatepackagemembershipID field value if set, zero value otherwise.
func (o *EzsigntemplatepackagemembershipRequestCompound) GetPkiEzsigntemplatepackagemembershipID() int32 {
	if o == nil || IsNil(o.PkiEzsigntemplatepackagemembershipID) {
		var ret int32
		return ret
	}
	return *o.PkiEzsigntemplatepackagemembershipID
}

// GetPkiEzsigntemplatepackagemembershipIDOk returns a tuple with the PkiEzsigntemplatepackagemembershipID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepackagemembershipRequestCompound) GetPkiEzsigntemplatepackagemembershipIDOk() (*int32, bool) {
	if o == nil || IsNil(o.PkiEzsigntemplatepackagemembershipID) {
		return nil, false
	}
	return o.PkiEzsigntemplatepackagemembershipID, true
}

// HasPkiEzsigntemplatepackagemembershipID returns a boolean if a field has been set.
func (o *EzsigntemplatepackagemembershipRequestCompound) HasPkiEzsigntemplatepackagemembershipID() bool {
	if o != nil && !IsNil(o.PkiEzsigntemplatepackagemembershipID) {
		return true
	}

	return false
}

// SetPkiEzsigntemplatepackagemembershipID gets a reference to the given int32 and assigns it to the PkiEzsigntemplatepackagemembershipID field.
func (o *EzsigntemplatepackagemembershipRequestCompound) SetPkiEzsigntemplatepackagemembershipID(v int32) {
	o.PkiEzsigntemplatepackagemembershipID = &v
}

// GetFkiEzsigntemplatepackageID returns the FkiEzsigntemplatepackageID field value
func (o *EzsigntemplatepackagemembershipRequestCompound) GetFkiEzsigntemplatepackageID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzsigntemplatepackageID
}

// GetFkiEzsigntemplatepackageIDOk returns a tuple with the FkiEzsigntemplatepackageID field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepackagemembershipRequestCompound) GetFkiEzsigntemplatepackageIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiEzsigntemplatepackageID, true
}

// SetFkiEzsigntemplatepackageID sets field value
func (o *EzsigntemplatepackagemembershipRequestCompound) SetFkiEzsigntemplatepackageID(v int32) {
	o.FkiEzsigntemplatepackageID = v
}

// GetFkiEzsigntemplateID returns the FkiEzsigntemplateID field value
func (o *EzsigntemplatepackagemembershipRequestCompound) GetFkiEzsigntemplateID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzsigntemplateID
}

// GetFkiEzsigntemplateIDOk returns a tuple with the FkiEzsigntemplateID field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepackagemembershipRequestCompound) GetFkiEzsigntemplateIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiEzsigntemplateID, true
}

// SetFkiEzsigntemplateID sets field value
func (o *EzsigntemplatepackagemembershipRequestCompound) SetFkiEzsigntemplateID(v int32) {
	o.FkiEzsigntemplateID = v
}

func (o EzsigntemplatepackagemembershipRequestCompound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigntemplatepackagemembershipRequestCompound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PkiEzsigntemplatepackagemembershipID) {
		toSerialize["pkiEzsigntemplatepackagemembershipID"] = o.PkiEzsigntemplatepackagemembershipID
	}
	toSerialize["fkiEzsigntemplatepackageID"] = o.FkiEzsigntemplatepackageID
	toSerialize["fkiEzsigntemplateID"] = o.FkiEzsigntemplateID
	return toSerialize, nil
}

func (o *EzsigntemplatepackagemembershipRequestCompound) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fkiEzsigntemplatepackageID",
		"fkiEzsigntemplateID",
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

	varEzsigntemplatepackagemembershipRequestCompound := _EzsigntemplatepackagemembershipRequestCompound{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsigntemplatepackagemembershipRequestCompound)

	if err != nil {
		return err
	}

	*o = EzsigntemplatepackagemembershipRequestCompound(varEzsigntemplatepackagemembershipRequestCompound)

	return err
}

type NullableEzsigntemplatepackagemembershipRequestCompound struct {
	value *EzsigntemplatepackagemembershipRequestCompound
	isSet bool
}

func (v NullableEzsigntemplatepackagemembershipRequestCompound) Get() *EzsigntemplatepackagemembershipRequestCompound {
	return v.value
}

func (v *NullableEzsigntemplatepackagemembershipRequestCompound) Set(val *EzsigntemplatepackagemembershipRequestCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigntemplatepackagemembershipRequestCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigntemplatepackagemembershipRequestCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigntemplatepackagemembershipRequestCompound(val *EzsigntemplatepackagemembershipRequestCompound) *NullableEzsigntemplatepackagemembershipRequestCompound {
	return &NullableEzsigntemplatepackagemembershipRequestCompound{value: val, isSet: true}
}

func (v NullableEzsigntemplatepackagemembershipRequestCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigntemplatepackagemembershipRequestCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


