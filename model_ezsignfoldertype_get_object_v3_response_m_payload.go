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

// checks if the EzsignfoldertypeGetObjectV3ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignfoldertypeGetObjectV3ResponseMPayload{}

// EzsignfoldertypeGetObjectV3ResponseMPayload Payload for GET /3/object/ezsignfoldertype/{pkiEzsignfoldertypeID}
type EzsignfoldertypeGetObjectV3ResponseMPayload struct {
	ObjEzsignfoldertype EzsignfoldertypeResponseCompoundV3 `json:"objEzsignfoldertype"`
}

type _EzsignfoldertypeGetObjectV3ResponseMPayload EzsignfoldertypeGetObjectV3ResponseMPayload

// NewEzsignfoldertypeGetObjectV3ResponseMPayload instantiates a new EzsignfoldertypeGetObjectV3ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignfoldertypeGetObjectV3ResponseMPayload(objEzsignfoldertype EzsignfoldertypeResponseCompoundV3) *EzsignfoldertypeGetObjectV3ResponseMPayload {
	this := EzsignfoldertypeGetObjectV3ResponseMPayload{}
	this.ObjEzsignfoldertype = objEzsignfoldertype
	return &this
}

// NewEzsignfoldertypeGetObjectV3ResponseMPayloadWithDefaults instantiates a new EzsignfoldertypeGetObjectV3ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignfoldertypeGetObjectV3ResponseMPayloadWithDefaults() *EzsignfoldertypeGetObjectV3ResponseMPayload {
	this := EzsignfoldertypeGetObjectV3ResponseMPayload{}
	return &this
}

// GetObjEzsignfoldertype returns the ObjEzsignfoldertype field value
func (o *EzsignfoldertypeGetObjectV3ResponseMPayload) GetObjEzsignfoldertype() EzsignfoldertypeResponseCompoundV3 {
	if o == nil {
		var ret EzsignfoldertypeResponseCompoundV3
		return ret
	}

	return o.ObjEzsignfoldertype
}

// GetObjEzsignfoldertypeOk returns a tuple with the ObjEzsignfoldertype field value
// and a boolean to check if the value has been set.
func (o *EzsignfoldertypeGetObjectV3ResponseMPayload) GetObjEzsignfoldertypeOk() (*EzsignfoldertypeResponseCompoundV3, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjEzsignfoldertype, true
}

// SetObjEzsignfoldertype sets field value
func (o *EzsignfoldertypeGetObjectV3ResponseMPayload) SetObjEzsignfoldertype(v EzsignfoldertypeResponseCompoundV3) {
	o.ObjEzsignfoldertype = v
}

func (o EzsignfoldertypeGetObjectV3ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignfoldertypeGetObjectV3ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objEzsignfoldertype"] = o.ObjEzsignfoldertype
	return toSerialize, nil
}

func (o *EzsignfoldertypeGetObjectV3ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"objEzsignfoldertype",
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

	varEzsignfoldertypeGetObjectV3ResponseMPayload := _EzsignfoldertypeGetObjectV3ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsignfoldertypeGetObjectV3ResponseMPayload)

	if err != nil {
		return err
	}

	*o = EzsignfoldertypeGetObjectV3ResponseMPayload(varEzsignfoldertypeGetObjectV3ResponseMPayload)

	return err
}

type NullableEzsignfoldertypeGetObjectV3ResponseMPayload struct {
	value *EzsignfoldertypeGetObjectV3ResponseMPayload
	isSet bool
}

func (v NullableEzsignfoldertypeGetObjectV3ResponseMPayload) Get() *EzsignfoldertypeGetObjectV3ResponseMPayload {
	return v.value
}

func (v *NullableEzsignfoldertypeGetObjectV3ResponseMPayload) Set(val *EzsignfoldertypeGetObjectV3ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignfoldertypeGetObjectV3ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignfoldertypeGetObjectV3ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignfoldertypeGetObjectV3ResponseMPayload(val *EzsignfoldertypeGetObjectV3ResponseMPayload) *NullableEzsignfoldertypeGetObjectV3ResponseMPayload {
	return &NullableEzsignfoldertypeGetObjectV3ResponseMPayload{value: val, isSet: true}
}

func (v NullableEzsignfoldertypeGetObjectV3ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignfoldertypeGetObjectV3ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


