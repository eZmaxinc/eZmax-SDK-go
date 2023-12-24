/*
eZmax API Definition (Full)

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.2.0
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the EzsignsignergroupResponseCompound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignsignergroupResponseCompound{}

// EzsignsignergroupResponseCompound An Ezsignsignergroup Object
type EzsignsignergroupResponseCompound struct {
	// The unique ID of the Ezsignsignergroup
	PkiEzsignsignergroupID int32 `json:"pkiEzsignsignergroupID"`
	ObjEzsignsignergroupDescription MultilingualEzsignsignergroupDescription `json:"objEzsignsignergroupDescription"`
	// The Description of the Ezsignsignergroup in the language of the requester
	SEzsignsignergroupDescriptionX *string `json:"sEzsignsignergroupDescriptionX,omitempty"`
}

type _EzsignsignergroupResponseCompound EzsignsignergroupResponseCompound

// NewEzsignsignergroupResponseCompound instantiates a new EzsignsignergroupResponseCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignsignergroupResponseCompound(pkiEzsignsignergroupID int32, objEzsignsignergroupDescription MultilingualEzsignsignergroupDescription) *EzsignsignergroupResponseCompound {
	this := EzsignsignergroupResponseCompound{}
	this.PkiEzsignsignergroupID = pkiEzsignsignergroupID
	this.ObjEzsignsignergroupDescription = objEzsignsignergroupDescription
	return &this
}

// NewEzsignsignergroupResponseCompoundWithDefaults instantiates a new EzsignsignergroupResponseCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignsignergroupResponseCompoundWithDefaults() *EzsignsignergroupResponseCompound {
	this := EzsignsignergroupResponseCompound{}
	return &this
}

// GetPkiEzsignsignergroupID returns the PkiEzsignsignergroupID field value
func (o *EzsignsignergroupResponseCompound) GetPkiEzsignsignergroupID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiEzsignsignergroupID
}

// GetPkiEzsignsignergroupIDOk returns a tuple with the PkiEzsignsignergroupID field value
// and a boolean to check if the value has been set.
func (o *EzsignsignergroupResponseCompound) GetPkiEzsignsignergroupIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiEzsignsignergroupID, true
}

// SetPkiEzsignsignergroupID sets field value
func (o *EzsignsignergroupResponseCompound) SetPkiEzsignsignergroupID(v int32) {
	o.PkiEzsignsignergroupID = v
}

// GetObjEzsignsignergroupDescription returns the ObjEzsignsignergroupDescription field value
func (o *EzsignsignergroupResponseCompound) GetObjEzsignsignergroupDescription() MultilingualEzsignsignergroupDescription {
	if o == nil {
		var ret MultilingualEzsignsignergroupDescription
		return ret
	}

	return o.ObjEzsignsignergroupDescription
}

// GetObjEzsignsignergroupDescriptionOk returns a tuple with the ObjEzsignsignergroupDescription field value
// and a boolean to check if the value has been set.
func (o *EzsignsignergroupResponseCompound) GetObjEzsignsignergroupDescriptionOk() (*MultilingualEzsignsignergroupDescription, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjEzsignsignergroupDescription, true
}

// SetObjEzsignsignergroupDescription sets field value
func (o *EzsignsignergroupResponseCompound) SetObjEzsignsignergroupDescription(v MultilingualEzsignsignergroupDescription) {
	o.ObjEzsignsignergroupDescription = v
}

// GetSEzsignsignergroupDescriptionX returns the SEzsignsignergroupDescriptionX field value if set, zero value otherwise.
func (o *EzsignsignergroupResponseCompound) GetSEzsignsignergroupDescriptionX() string {
	if o == nil || IsNil(o.SEzsignsignergroupDescriptionX) {
		var ret string
		return ret
	}
	return *o.SEzsignsignergroupDescriptionX
}

// GetSEzsignsignergroupDescriptionXOk returns a tuple with the SEzsignsignergroupDescriptionX field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignsignergroupResponseCompound) GetSEzsignsignergroupDescriptionXOk() (*string, bool) {
	if o == nil || IsNil(o.SEzsignsignergroupDescriptionX) {
		return nil, false
	}
	return o.SEzsignsignergroupDescriptionX, true
}

// HasSEzsignsignergroupDescriptionX returns a boolean if a field has been set.
func (o *EzsignsignergroupResponseCompound) HasSEzsignsignergroupDescriptionX() bool {
	if o != nil && !IsNil(o.SEzsignsignergroupDescriptionX) {
		return true
	}

	return false
}

// SetSEzsignsignergroupDescriptionX gets a reference to the given string and assigns it to the SEzsignsignergroupDescriptionX field.
func (o *EzsignsignergroupResponseCompound) SetSEzsignsignergroupDescriptionX(v string) {
	o.SEzsignsignergroupDescriptionX = &v
}

func (o EzsignsignergroupResponseCompound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignsignergroupResponseCompound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiEzsignsignergroupID"] = o.PkiEzsignsignergroupID
	toSerialize["objEzsignsignergroupDescription"] = o.ObjEzsignsignergroupDescription
	if !IsNil(o.SEzsignsignergroupDescriptionX) {
		toSerialize["sEzsignsignergroupDescriptionX"] = o.SEzsignsignergroupDescriptionX
	}
	return toSerialize, nil
}

func (o *EzsignsignergroupResponseCompound) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiEzsignsignergroupID",
		"objEzsignsignergroupDescription",
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

	varEzsignsignergroupResponseCompound := _EzsignsignergroupResponseCompound{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsignsignergroupResponseCompound)

	if err != nil {
		return err
	}

	*o = EzsignsignergroupResponseCompound(varEzsignsignergroupResponseCompound)

	return err
}

type NullableEzsignsignergroupResponseCompound struct {
	value *EzsignsignergroupResponseCompound
	isSet bool
}

func (v NullableEzsignsignergroupResponseCompound) Get() *EzsignsignergroupResponseCompound {
	return v.value
}

func (v *NullableEzsignsignergroupResponseCompound) Set(val *EzsignsignergroupResponseCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignsignergroupResponseCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignsignergroupResponseCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignsignergroupResponseCompound(val *EzsignsignergroupResponseCompound) *NullableEzsignsignergroupResponseCompound {
	return &NullableEzsignsignergroupResponseCompound{value: val, isSet: true}
}

func (v NullableEzsignsignergroupResponseCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignsignergroupResponseCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


