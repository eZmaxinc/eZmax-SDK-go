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

// checks if the EzsignsigningreasonRequestCompound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignsigningreasonRequestCompound{}

// EzsignsigningreasonRequestCompound A Ezsignsigningreason Object and children
type EzsignsigningreasonRequestCompound struct {
	// The unique ID of the Ezsignsigningreason
	PkiEzsignsigningreasonID *int32 `json:"pkiEzsignsigningreasonID,omitempty"`
	ObjEzsignsigningreasonDescription MultilingualEzsignsigningreasonDescription `json:"objEzsignsigningreasonDescription"`
	// Whether the ezsignsigningreason is active or not
	BEzsignsigningreasonIsactive bool `json:"bEzsignsigningreasonIsactive"`
}

type _EzsignsigningreasonRequestCompound EzsignsigningreasonRequestCompound

// NewEzsignsigningreasonRequestCompound instantiates a new EzsignsigningreasonRequestCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignsigningreasonRequestCompound(objEzsignsigningreasonDescription MultilingualEzsignsigningreasonDescription, bEzsignsigningreasonIsactive bool) *EzsignsigningreasonRequestCompound {
	this := EzsignsigningreasonRequestCompound{}
	this.ObjEzsignsigningreasonDescription = objEzsignsigningreasonDescription
	this.BEzsignsigningreasonIsactive = bEzsignsigningreasonIsactive
	return &this
}

// NewEzsignsigningreasonRequestCompoundWithDefaults instantiates a new EzsignsigningreasonRequestCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignsigningreasonRequestCompoundWithDefaults() *EzsignsigningreasonRequestCompound {
	this := EzsignsigningreasonRequestCompound{}
	return &this
}

// GetPkiEzsignsigningreasonID returns the PkiEzsignsigningreasonID field value if set, zero value otherwise.
func (o *EzsignsigningreasonRequestCompound) GetPkiEzsignsigningreasonID() int32 {
	if o == nil || IsNil(o.PkiEzsignsigningreasonID) {
		var ret int32
		return ret
	}
	return *o.PkiEzsignsigningreasonID
}

// GetPkiEzsignsigningreasonIDOk returns a tuple with the PkiEzsignsigningreasonID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignsigningreasonRequestCompound) GetPkiEzsignsigningreasonIDOk() (*int32, bool) {
	if o == nil || IsNil(o.PkiEzsignsigningreasonID) {
		return nil, false
	}
	return o.PkiEzsignsigningreasonID, true
}

// HasPkiEzsignsigningreasonID returns a boolean if a field has been set.
func (o *EzsignsigningreasonRequestCompound) HasPkiEzsignsigningreasonID() bool {
	if o != nil && !IsNil(o.PkiEzsignsigningreasonID) {
		return true
	}

	return false
}

// SetPkiEzsignsigningreasonID gets a reference to the given int32 and assigns it to the PkiEzsignsigningreasonID field.
func (o *EzsignsigningreasonRequestCompound) SetPkiEzsignsigningreasonID(v int32) {
	o.PkiEzsignsigningreasonID = &v
}

// GetObjEzsignsigningreasonDescription returns the ObjEzsignsigningreasonDescription field value
func (o *EzsignsigningreasonRequestCompound) GetObjEzsignsigningreasonDescription() MultilingualEzsignsigningreasonDescription {
	if o == nil {
		var ret MultilingualEzsignsigningreasonDescription
		return ret
	}

	return o.ObjEzsignsigningreasonDescription
}

// GetObjEzsignsigningreasonDescriptionOk returns a tuple with the ObjEzsignsigningreasonDescription field value
// and a boolean to check if the value has been set.
func (o *EzsignsigningreasonRequestCompound) GetObjEzsignsigningreasonDescriptionOk() (*MultilingualEzsignsigningreasonDescription, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjEzsignsigningreasonDescription, true
}

// SetObjEzsignsigningreasonDescription sets field value
func (o *EzsignsigningreasonRequestCompound) SetObjEzsignsigningreasonDescription(v MultilingualEzsignsigningreasonDescription) {
	o.ObjEzsignsigningreasonDescription = v
}

// GetBEzsignsigningreasonIsactive returns the BEzsignsigningreasonIsactive field value
func (o *EzsignsigningreasonRequestCompound) GetBEzsignsigningreasonIsactive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BEzsignsigningreasonIsactive
}

// GetBEzsignsigningreasonIsactiveOk returns a tuple with the BEzsignsigningreasonIsactive field value
// and a boolean to check if the value has been set.
func (o *EzsignsigningreasonRequestCompound) GetBEzsignsigningreasonIsactiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BEzsignsigningreasonIsactive, true
}

// SetBEzsignsigningreasonIsactive sets field value
func (o *EzsignsigningreasonRequestCompound) SetBEzsignsigningreasonIsactive(v bool) {
	o.BEzsignsigningreasonIsactive = v
}

func (o EzsignsigningreasonRequestCompound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignsigningreasonRequestCompound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PkiEzsignsigningreasonID) {
		toSerialize["pkiEzsignsigningreasonID"] = o.PkiEzsignsigningreasonID
	}
	toSerialize["objEzsignsigningreasonDescription"] = o.ObjEzsignsigningreasonDescription
	toSerialize["bEzsignsigningreasonIsactive"] = o.BEzsignsigningreasonIsactive
	return toSerialize, nil
}

func (o *EzsignsigningreasonRequestCompound) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"objEzsignsigningreasonDescription",
		"bEzsignsigningreasonIsactive",
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

	varEzsignsigningreasonRequestCompound := _EzsignsigningreasonRequestCompound{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsignsigningreasonRequestCompound)

	if err != nil {
		return err
	}

	*o = EzsignsigningreasonRequestCompound(varEzsignsigningreasonRequestCompound)

	return err
}

type NullableEzsignsigningreasonRequestCompound struct {
	value *EzsignsigningreasonRequestCompound
	isSet bool
}

func (v NullableEzsignsigningreasonRequestCompound) Get() *EzsignsigningreasonRequestCompound {
	return v.value
}

func (v *NullableEzsignsigningreasonRequestCompound) Set(val *EzsignsigningreasonRequestCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignsigningreasonRequestCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignsigningreasonRequestCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignsigningreasonRequestCompound(val *EzsignsigningreasonRequestCompound) *NullableEzsignsigningreasonRequestCompound {
	return &NullableEzsignsigningreasonRequestCompound{value: val, isSet: true}
}

func (v NullableEzsignsigningreasonRequestCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignsigningreasonRequestCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


