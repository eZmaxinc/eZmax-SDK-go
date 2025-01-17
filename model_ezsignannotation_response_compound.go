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

// checks if the EzsignannotationResponseCompound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignannotationResponseCompound{}

// EzsignannotationResponseCompound A Ezsignannotation Object
type EzsignannotationResponseCompound struct {
	EzsignannotationResponse
	ObjTextstylestatic *TextstylestaticResponseCompound `json:"objTextstylestatic,omitempty"`
}

type _EzsignannotationResponseCompound EzsignannotationResponseCompound

// NewEzsignannotationResponseCompound instantiates a new EzsignannotationResponseCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignannotationResponseCompound(pkiEzsignannotationID int32, fkiEzsigndocumentID int32, eEzsignannotationType FieldEEzsignannotationType, iEzsignannotationX int32, iEzsignannotationY int32, iEzsignpagePagenumber int32) *EzsignannotationResponseCompound {
	this := EzsignannotationResponseCompound{}
	this.PkiEzsignannotationID = pkiEzsignannotationID
	this.FkiEzsigndocumentID = fkiEzsigndocumentID
	this.EEzsignannotationType = eEzsignannotationType
	this.IEzsignannotationX = iEzsignannotationX
	this.IEzsignannotationY = iEzsignannotationY
	this.IEzsignpagePagenumber = iEzsignpagePagenumber
	return &this
}

// NewEzsignannotationResponseCompoundWithDefaults instantiates a new EzsignannotationResponseCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignannotationResponseCompoundWithDefaults() *EzsignannotationResponseCompound {
	this := EzsignannotationResponseCompound{}
	return &this
}

// GetObjTextstylestatic returns the ObjTextstylestatic field value if set, zero value otherwise.
func (o *EzsignannotationResponseCompound) GetObjTextstylestatic() TextstylestaticResponseCompound {
	if o == nil || IsNil(o.ObjTextstylestatic) {
		var ret TextstylestaticResponseCompound
		return ret
	}
	return *o.ObjTextstylestatic
}

// GetObjTextstylestaticOk returns a tuple with the ObjTextstylestatic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignannotationResponseCompound) GetObjTextstylestaticOk() (*TextstylestaticResponseCompound, bool) {
	if o == nil || IsNil(o.ObjTextstylestatic) {
		return nil, false
	}
	return o.ObjTextstylestatic, true
}

// HasObjTextstylestatic returns a boolean if a field has been set.
func (o *EzsignannotationResponseCompound) HasObjTextstylestatic() bool {
	if o != nil && !IsNil(o.ObjTextstylestatic) {
		return true
	}

	return false
}

// SetObjTextstylestatic gets a reference to the given TextstylestaticResponseCompound and assigns it to the ObjTextstylestatic field.
func (o *EzsignannotationResponseCompound) SetObjTextstylestatic(v TextstylestaticResponseCompound) {
	o.ObjTextstylestatic = &v
}

func (o EzsignannotationResponseCompound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignannotationResponseCompound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ObjTextstylestatic) {
		toSerialize["objTextstylestatic"] = o.ObjTextstylestatic
	}
	return toSerialize, nil
}

func (o *EzsignannotationResponseCompound) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiEzsignannotationID",
		"fkiEzsigndocumentID",
		"eEzsignannotationType",
		"iEzsignannotationX",
		"iEzsignannotationY",
		"iEzsignpagePagenumber",
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

	varEzsignannotationResponseCompound := _EzsignannotationResponseCompound{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsignannotationResponseCompound)

	if err != nil {
		return err
	}

	*o = EzsignannotationResponseCompound(varEzsignannotationResponseCompound)

	return err
}

type NullableEzsignannotationResponseCompound struct {
	value *EzsignannotationResponseCompound
	isSet bool
}

func (v NullableEzsignannotationResponseCompound) Get() *EzsignannotationResponseCompound {
	return v.value
}

func (v *NullableEzsignannotationResponseCompound) Set(val *EzsignannotationResponseCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignannotationResponseCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignannotationResponseCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignannotationResponseCompound(val *EzsignannotationResponseCompound) *NullableEzsignannotationResponseCompound {
	return &NullableEzsignannotationResponseCompound{value: val, isSet: true}
}

func (v NullableEzsignannotationResponseCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignannotationResponseCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


