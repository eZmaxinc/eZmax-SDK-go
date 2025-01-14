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

// checks if the EzsignfoldertypeGetObjectV4ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignfoldertypeGetObjectV4ResponseMPayload{}

// EzsignfoldertypeGetObjectV4ResponseMPayload Payload for GET /4/object/ezsignfoldertype/{pkiEzsignfoldertypeID}
type EzsignfoldertypeGetObjectV4ResponseMPayload struct {
	ObjEzsignfoldertype EzsignfoldertypeResponseCompoundV4 `json:"objEzsignfoldertype"`
}

type _EzsignfoldertypeGetObjectV4ResponseMPayload EzsignfoldertypeGetObjectV4ResponseMPayload

// NewEzsignfoldertypeGetObjectV4ResponseMPayload instantiates a new EzsignfoldertypeGetObjectV4ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignfoldertypeGetObjectV4ResponseMPayload(objEzsignfoldertype EzsignfoldertypeResponseCompoundV4) *EzsignfoldertypeGetObjectV4ResponseMPayload {
	this := EzsignfoldertypeGetObjectV4ResponseMPayload{}
	this.ObjEzsignfoldertype = objEzsignfoldertype
	return &this
}

// NewEzsignfoldertypeGetObjectV4ResponseMPayloadWithDefaults instantiates a new EzsignfoldertypeGetObjectV4ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignfoldertypeGetObjectV4ResponseMPayloadWithDefaults() *EzsignfoldertypeGetObjectV4ResponseMPayload {
	this := EzsignfoldertypeGetObjectV4ResponseMPayload{}
	return &this
}

// GetObjEzsignfoldertype returns the ObjEzsignfoldertype field value
func (o *EzsignfoldertypeGetObjectV4ResponseMPayload) GetObjEzsignfoldertype() EzsignfoldertypeResponseCompoundV4 {
	if o == nil {
		var ret EzsignfoldertypeResponseCompoundV4
		return ret
	}

	return o.ObjEzsignfoldertype
}

// GetObjEzsignfoldertypeOk returns a tuple with the ObjEzsignfoldertype field value
// and a boolean to check if the value has been set.
func (o *EzsignfoldertypeGetObjectV4ResponseMPayload) GetObjEzsignfoldertypeOk() (*EzsignfoldertypeResponseCompoundV4, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjEzsignfoldertype, true
}

// SetObjEzsignfoldertype sets field value
func (o *EzsignfoldertypeGetObjectV4ResponseMPayload) SetObjEzsignfoldertype(v EzsignfoldertypeResponseCompoundV4) {
	o.ObjEzsignfoldertype = v
}

func (o EzsignfoldertypeGetObjectV4ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignfoldertypeGetObjectV4ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objEzsignfoldertype"] = o.ObjEzsignfoldertype
	return toSerialize, nil
}

func (o *EzsignfoldertypeGetObjectV4ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
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

	varEzsignfoldertypeGetObjectV4ResponseMPayload := _EzsignfoldertypeGetObjectV4ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsignfoldertypeGetObjectV4ResponseMPayload)

	if err != nil {
		return err
	}

	*o = EzsignfoldertypeGetObjectV4ResponseMPayload(varEzsignfoldertypeGetObjectV4ResponseMPayload)

	return err
}

type NullableEzsignfoldertypeGetObjectV4ResponseMPayload struct {
	value *EzsignfoldertypeGetObjectV4ResponseMPayload
	isSet bool
}

func (v NullableEzsignfoldertypeGetObjectV4ResponseMPayload) Get() *EzsignfoldertypeGetObjectV4ResponseMPayload {
	return v.value
}

func (v *NullableEzsignfoldertypeGetObjectV4ResponseMPayload) Set(val *EzsignfoldertypeGetObjectV4ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignfoldertypeGetObjectV4ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignfoldertypeGetObjectV4ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignfoldertypeGetObjectV4ResponseMPayload(val *EzsignfoldertypeGetObjectV4ResponseMPayload) *NullableEzsignfoldertypeGetObjectV4ResponseMPayload {
	return &NullableEzsignfoldertypeGetObjectV4ResponseMPayload{value: val, isSet: true}
}

func (v NullableEzsignfoldertypeGetObjectV4ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignfoldertypeGetObjectV4ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


