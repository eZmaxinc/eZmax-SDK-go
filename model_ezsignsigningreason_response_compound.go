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

// checks if the EzsignsigningreasonResponseCompound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignsigningreasonResponseCompound{}

// EzsignsigningreasonResponseCompound A Ezsignsigningreason Object
type EzsignsigningreasonResponseCompound struct {
	// The unique ID of the Ezsignsigningreason
	PkiEzsignsigningreasonID int32 `json:"pkiEzsignsigningreasonID"`
	ObjEzsignsigningreasonDescription MultilingualEzsignsigningreasonDescription `json:"objEzsignsigningreasonDescription"`
	// Whether the ezsignsigningreason is active or not
	BEzsignsigningreasonIsactive bool `json:"bEzsignsigningreasonIsactive"`
}

type _EzsignsigningreasonResponseCompound EzsignsigningreasonResponseCompound

// NewEzsignsigningreasonResponseCompound instantiates a new EzsignsigningreasonResponseCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignsigningreasonResponseCompound(pkiEzsignsigningreasonID int32, objEzsignsigningreasonDescription MultilingualEzsignsigningreasonDescription, bEzsignsigningreasonIsactive bool) *EzsignsigningreasonResponseCompound {
	this := EzsignsigningreasonResponseCompound{}
	this.PkiEzsignsigningreasonID = pkiEzsignsigningreasonID
	this.ObjEzsignsigningreasonDescription = objEzsignsigningreasonDescription
	this.BEzsignsigningreasonIsactive = bEzsignsigningreasonIsactive
	return &this
}

// NewEzsignsigningreasonResponseCompoundWithDefaults instantiates a new EzsignsigningreasonResponseCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignsigningreasonResponseCompoundWithDefaults() *EzsignsigningreasonResponseCompound {
	this := EzsignsigningreasonResponseCompound{}
	return &this
}

// GetPkiEzsignsigningreasonID returns the PkiEzsignsigningreasonID field value
func (o *EzsignsigningreasonResponseCompound) GetPkiEzsignsigningreasonID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiEzsignsigningreasonID
}

// GetPkiEzsignsigningreasonIDOk returns a tuple with the PkiEzsignsigningreasonID field value
// and a boolean to check if the value has been set.
func (o *EzsignsigningreasonResponseCompound) GetPkiEzsignsigningreasonIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiEzsignsigningreasonID, true
}

// SetPkiEzsignsigningreasonID sets field value
func (o *EzsignsigningreasonResponseCompound) SetPkiEzsignsigningreasonID(v int32) {
	o.PkiEzsignsigningreasonID = v
}

// GetObjEzsignsigningreasonDescription returns the ObjEzsignsigningreasonDescription field value
func (o *EzsignsigningreasonResponseCompound) GetObjEzsignsigningreasonDescription() MultilingualEzsignsigningreasonDescription {
	if o == nil {
		var ret MultilingualEzsignsigningreasonDescription
		return ret
	}

	return o.ObjEzsignsigningreasonDescription
}

// GetObjEzsignsigningreasonDescriptionOk returns a tuple with the ObjEzsignsigningreasonDescription field value
// and a boolean to check if the value has been set.
func (o *EzsignsigningreasonResponseCompound) GetObjEzsignsigningreasonDescriptionOk() (*MultilingualEzsignsigningreasonDescription, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjEzsignsigningreasonDescription, true
}

// SetObjEzsignsigningreasonDescription sets field value
func (o *EzsignsigningreasonResponseCompound) SetObjEzsignsigningreasonDescription(v MultilingualEzsignsigningreasonDescription) {
	o.ObjEzsignsigningreasonDescription = v
}

// GetBEzsignsigningreasonIsactive returns the BEzsignsigningreasonIsactive field value
func (o *EzsignsigningreasonResponseCompound) GetBEzsignsigningreasonIsactive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BEzsignsigningreasonIsactive
}

// GetBEzsignsigningreasonIsactiveOk returns a tuple with the BEzsignsigningreasonIsactive field value
// and a boolean to check if the value has been set.
func (o *EzsignsigningreasonResponseCompound) GetBEzsignsigningreasonIsactiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BEzsignsigningreasonIsactive, true
}

// SetBEzsignsigningreasonIsactive sets field value
func (o *EzsignsigningreasonResponseCompound) SetBEzsignsigningreasonIsactive(v bool) {
	o.BEzsignsigningreasonIsactive = v
}

func (o EzsignsigningreasonResponseCompound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignsigningreasonResponseCompound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiEzsignsigningreasonID"] = o.PkiEzsignsigningreasonID
	toSerialize["objEzsignsigningreasonDescription"] = o.ObjEzsignsigningreasonDescription
	toSerialize["bEzsignsigningreasonIsactive"] = o.BEzsignsigningreasonIsactive
	return toSerialize, nil
}

func (o *EzsignsigningreasonResponseCompound) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiEzsignsigningreasonID",
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

	varEzsignsigningreasonResponseCompound := _EzsignsigningreasonResponseCompound{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsignsigningreasonResponseCompound)

	if err != nil {
		return err
	}

	*o = EzsignsigningreasonResponseCompound(varEzsignsigningreasonResponseCompound)

	return err
}

type NullableEzsignsigningreasonResponseCompound struct {
	value *EzsignsigningreasonResponseCompound
	isSet bool
}

func (v NullableEzsignsigningreasonResponseCompound) Get() *EzsignsigningreasonResponseCompound {
	return v.value
}

func (v *NullableEzsignsigningreasonResponseCompound) Set(val *EzsignsigningreasonResponseCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignsigningreasonResponseCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignsigningreasonResponseCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignsigningreasonResponseCompound(val *EzsignsigningreasonResponseCompound) *NullableEzsignsigningreasonResponseCompound {
	return &NullableEzsignsigningreasonResponseCompound{value: val, isSet: true}
}

func (v NullableEzsignsigningreasonResponseCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignsigningreasonResponseCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


