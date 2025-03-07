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

// checks if the EzsignfoldertypeGetObjectV2ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignfoldertypeGetObjectV2ResponseMPayload{}

// EzsignfoldertypeGetObjectV2ResponseMPayload Payload for GET /2/object/ezsignfoldertype/{pkiEzsignfoldertypeID}
type EzsignfoldertypeGetObjectV2ResponseMPayload struct {
	ObjEzsignfoldertype EzsignfoldertypeResponseCompound `json:"objEzsignfoldertype"`
}

type _EzsignfoldertypeGetObjectV2ResponseMPayload EzsignfoldertypeGetObjectV2ResponseMPayload

// NewEzsignfoldertypeGetObjectV2ResponseMPayload instantiates a new EzsignfoldertypeGetObjectV2ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignfoldertypeGetObjectV2ResponseMPayload(objEzsignfoldertype EzsignfoldertypeResponseCompound) *EzsignfoldertypeGetObjectV2ResponseMPayload {
	this := EzsignfoldertypeGetObjectV2ResponseMPayload{}
	this.ObjEzsignfoldertype = objEzsignfoldertype
	return &this
}

// NewEzsignfoldertypeGetObjectV2ResponseMPayloadWithDefaults instantiates a new EzsignfoldertypeGetObjectV2ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignfoldertypeGetObjectV2ResponseMPayloadWithDefaults() *EzsignfoldertypeGetObjectV2ResponseMPayload {
	this := EzsignfoldertypeGetObjectV2ResponseMPayload{}
	return &this
}

// GetObjEzsignfoldertype returns the ObjEzsignfoldertype field value
func (o *EzsignfoldertypeGetObjectV2ResponseMPayload) GetObjEzsignfoldertype() EzsignfoldertypeResponseCompound {
	if o == nil {
		var ret EzsignfoldertypeResponseCompound
		return ret
	}

	return o.ObjEzsignfoldertype
}

// GetObjEzsignfoldertypeOk returns a tuple with the ObjEzsignfoldertype field value
// and a boolean to check if the value has been set.
func (o *EzsignfoldertypeGetObjectV2ResponseMPayload) GetObjEzsignfoldertypeOk() (*EzsignfoldertypeResponseCompound, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjEzsignfoldertype, true
}

// SetObjEzsignfoldertype sets field value
func (o *EzsignfoldertypeGetObjectV2ResponseMPayload) SetObjEzsignfoldertype(v EzsignfoldertypeResponseCompound) {
	o.ObjEzsignfoldertype = v
}

func (o EzsignfoldertypeGetObjectV2ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignfoldertypeGetObjectV2ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objEzsignfoldertype"] = o.ObjEzsignfoldertype
	return toSerialize, nil
}

func (o *EzsignfoldertypeGetObjectV2ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
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

	varEzsignfoldertypeGetObjectV2ResponseMPayload := _EzsignfoldertypeGetObjectV2ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsignfoldertypeGetObjectV2ResponseMPayload)

	if err != nil {
		return err
	}

	*o = EzsignfoldertypeGetObjectV2ResponseMPayload(varEzsignfoldertypeGetObjectV2ResponseMPayload)

	return err
}

type NullableEzsignfoldertypeGetObjectV2ResponseMPayload struct {
	value *EzsignfoldertypeGetObjectV2ResponseMPayload
	isSet bool
}

func (v NullableEzsignfoldertypeGetObjectV2ResponseMPayload) Get() *EzsignfoldertypeGetObjectV2ResponseMPayload {
	return v.value
}

func (v *NullableEzsignfoldertypeGetObjectV2ResponseMPayload) Set(val *EzsignfoldertypeGetObjectV2ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignfoldertypeGetObjectV2ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignfoldertypeGetObjectV2ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignfoldertypeGetObjectV2ResponseMPayload(val *EzsignfoldertypeGetObjectV2ResponseMPayload) *NullableEzsignfoldertypeGetObjectV2ResponseMPayload {
	return &NullableEzsignfoldertypeGetObjectV2ResponseMPayload{value: val, isSet: true}
}

func (v NullableEzsignfoldertypeGetObjectV2ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignfoldertypeGetObjectV2ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


