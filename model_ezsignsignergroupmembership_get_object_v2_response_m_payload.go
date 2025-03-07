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

// checks if the EzsignsignergroupmembershipGetObjectV2ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignsignergroupmembershipGetObjectV2ResponseMPayload{}

// EzsignsignergroupmembershipGetObjectV2ResponseMPayload Payload for GET /2/object/ezsignsignergroupmembership/{pkiEzsignsignergroupmembershipID}
type EzsignsignergroupmembershipGetObjectV2ResponseMPayload struct {
	ObjEzsignsignergroupmembership EzsignsignergroupmembershipResponseCompound `json:"objEzsignsignergroupmembership"`
}

type _EzsignsignergroupmembershipGetObjectV2ResponseMPayload EzsignsignergroupmembershipGetObjectV2ResponseMPayload

// NewEzsignsignergroupmembershipGetObjectV2ResponseMPayload instantiates a new EzsignsignergroupmembershipGetObjectV2ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignsignergroupmembershipGetObjectV2ResponseMPayload(objEzsignsignergroupmembership EzsignsignergroupmembershipResponseCompound) *EzsignsignergroupmembershipGetObjectV2ResponseMPayload {
	this := EzsignsignergroupmembershipGetObjectV2ResponseMPayload{}
	this.ObjEzsignsignergroupmembership = objEzsignsignergroupmembership
	return &this
}

// NewEzsignsignergroupmembershipGetObjectV2ResponseMPayloadWithDefaults instantiates a new EzsignsignergroupmembershipGetObjectV2ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignsignergroupmembershipGetObjectV2ResponseMPayloadWithDefaults() *EzsignsignergroupmembershipGetObjectV2ResponseMPayload {
	this := EzsignsignergroupmembershipGetObjectV2ResponseMPayload{}
	return &this
}

// GetObjEzsignsignergroupmembership returns the ObjEzsignsignergroupmembership field value
func (o *EzsignsignergroupmembershipGetObjectV2ResponseMPayload) GetObjEzsignsignergroupmembership() EzsignsignergroupmembershipResponseCompound {
	if o == nil {
		var ret EzsignsignergroupmembershipResponseCompound
		return ret
	}

	return o.ObjEzsignsignergroupmembership
}

// GetObjEzsignsignergroupmembershipOk returns a tuple with the ObjEzsignsignergroupmembership field value
// and a boolean to check if the value has been set.
func (o *EzsignsignergroupmembershipGetObjectV2ResponseMPayload) GetObjEzsignsignergroupmembershipOk() (*EzsignsignergroupmembershipResponseCompound, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjEzsignsignergroupmembership, true
}

// SetObjEzsignsignergroupmembership sets field value
func (o *EzsignsignergroupmembershipGetObjectV2ResponseMPayload) SetObjEzsignsignergroupmembership(v EzsignsignergroupmembershipResponseCompound) {
	o.ObjEzsignsignergroupmembership = v
}

func (o EzsignsignergroupmembershipGetObjectV2ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignsignergroupmembershipGetObjectV2ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objEzsignsignergroupmembership"] = o.ObjEzsignsignergroupmembership
	return toSerialize, nil
}

func (o *EzsignsignergroupmembershipGetObjectV2ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"objEzsignsignergroupmembership",
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

	varEzsignsignergroupmembershipGetObjectV2ResponseMPayload := _EzsignsignergroupmembershipGetObjectV2ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsignsignergroupmembershipGetObjectV2ResponseMPayload)

	if err != nil {
		return err
	}

	*o = EzsignsignergroupmembershipGetObjectV2ResponseMPayload(varEzsignsignergroupmembershipGetObjectV2ResponseMPayload)

	return err
}

type NullableEzsignsignergroupmembershipGetObjectV2ResponseMPayload struct {
	value *EzsignsignergroupmembershipGetObjectV2ResponseMPayload
	isSet bool
}

func (v NullableEzsignsignergroupmembershipGetObjectV2ResponseMPayload) Get() *EzsignsignergroupmembershipGetObjectV2ResponseMPayload {
	return v.value
}

func (v *NullableEzsignsignergroupmembershipGetObjectV2ResponseMPayload) Set(val *EzsignsignergroupmembershipGetObjectV2ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignsignergroupmembershipGetObjectV2ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignsignergroupmembershipGetObjectV2ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignsignergroupmembershipGetObjectV2ResponseMPayload(val *EzsignsignergroupmembershipGetObjectV2ResponseMPayload) *NullableEzsignsignergroupmembershipGetObjectV2ResponseMPayload {
	return &NullableEzsignsignergroupmembershipGetObjectV2ResponseMPayload{value: val, isSet: true}
}

func (v NullableEzsignsignergroupmembershipGetObjectV2ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignsignergroupmembershipGetObjectV2ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


