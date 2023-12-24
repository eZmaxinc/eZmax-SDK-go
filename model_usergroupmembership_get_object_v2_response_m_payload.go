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

// checks if the UsergroupmembershipGetObjectV2ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UsergroupmembershipGetObjectV2ResponseMPayload{}

// UsergroupmembershipGetObjectV2ResponseMPayload Payload for GET /2/object/usergroupmembership/{pkiUsergroupmembershipID}
type UsergroupmembershipGetObjectV2ResponseMPayload struct {
	ObjUsergroupmembership UsergroupmembershipResponseCompound `json:"objUsergroupmembership"`
}

type _UsergroupmembershipGetObjectV2ResponseMPayload UsergroupmembershipGetObjectV2ResponseMPayload

// NewUsergroupmembershipGetObjectV2ResponseMPayload instantiates a new UsergroupmembershipGetObjectV2ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsergroupmembershipGetObjectV2ResponseMPayload(objUsergroupmembership UsergroupmembershipResponseCompound) *UsergroupmembershipGetObjectV2ResponseMPayload {
	this := UsergroupmembershipGetObjectV2ResponseMPayload{}
	this.ObjUsergroupmembership = objUsergroupmembership
	return &this
}

// NewUsergroupmembershipGetObjectV2ResponseMPayloadWithDefaults instantiates a new UsergroupmembershipGetObjectV2ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsergroupmembershipGetObjectV2ResponseMPayloadWithDefaults() *UsergroupmembershipGetObjectV2ResponseMPayload {
	this := UsergroupmembershipGetObjectV2ResponseMPayload{}
	return &this
}

// GetObjUsergroupmembership returns the ObjUsergroupmembership field value
func (o *UsergroupmembershipGetObjectV2ResponseMPayload) GetObjUsergroupmembership() UsergroupmembershipResponseCompound {
	if o == nil {
		var ret UsergroupmembershipResponseCompound
		return ret
	}

	return o.ObjUsergroupmembership
}

// GetObjUsergroupmembershipOk returns a tuple with the ObjUsergroupmembership field value
// and a boolean to check if the value has been set.
func (o *UsergroupmembershipGetObjectV2ResponseMPayload) GetObjUsergroupmembershipOk() (*UsergroupmembershipResponseCompound, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjUsergroupmembership, true
}

// SetObjUsergroupmembership sets field value
func (o *UsergroupmembershipGetObjectV2ResponseMPayload) SetObjUsergroupmembership(v UsergroupmembershipResponseCompound) {
	o.ObjUsergroupmembership = v
}

func (o UsergroupmembershipGetObjectV2ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsergroupmembershipGetObjectV2ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objUsergroupmembership"] = o.ObjUsergroupmembership
	return toSerialize, nil
}

func (o *UsergroupmembershipGetObjectV2ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"objUsergroupmembership",
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

	varUsergroupmembershipGetObjectV2ResponseMPayload := _UsergroupmembershipGetObjectV2ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUsergroupmembershipGetObjectV2ResponseMPayload)

	if err != nil {
		return err
	}

	*o = UsergroupmembershipGetObjectV2ResponseMPayload(varUsergroupmembershipGetObjectV2ResponseMPayload)

	return err
}

type NullableUsergroupmembershipGetObjectV2ResponseMPayload struct {
	value *UsergroupmembershipGetObjectV2ResponseMPayload
	isSet bool
}

func (v NullableUsergroupmembershipGetObjectV2ResponseMPayload) Get() *UsergroupmembershipGetObjectV2ResponseMPayload {
	return v.value
}

func (v *NullableUsergroupmembershipGetObjectV2ResponseMPayload) Set(val *UsergroupmembershipGetObjectV2ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableUsergroupmembershipGetObjectV2ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableUsergroupmembershipGetObjectV2ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsergroupmembershipGetObjectV2ResponseMPayload(val *UsergroupmembershipGetObjectV2ResponseMPayload) *NullableUsergroupmembershipGetObjectV2ResponseMPayload {
	return &NullableUsergroupmembershipGetObjectV2ResponseMPayload{value: val, isSet: true}
}

func (v NullableUsergroupmembershipGetObjectV2ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsergroupmembershipGetObjectV2ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


