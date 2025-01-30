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

// checks if the EzsignsignergroupGetObjectV2ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignsignergroupGetObjectV2ResponseMPayload{}

// EzsignsignergroupGetObjectV2ResponseMPayload Payload for GET /2/object/ezsignsignergroup/{pkiEzsignsignergroupID}
type EzsignsignergroupGetObjectV2ResponseMPayload struct {
	ObjEzsignsignergroup EzsignsignergroupResponseCompound `json:"objEzsignsignergroup"`
}

type _EzsignsignergroupGetObjectV2ResponseMPayload EzsignsignergroupGetObjectV2ResponseMPayload

// NewEzsignsignergroupGetObjectV2ResponseMPayload instantiates a new EzsignsignergroupGetObjectV2ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignsignergroupGetObjectV2ResponseMPayload(objEzsignsignergroup EzsignsignergroupResponseCompound) *EzsignsignergroupGetObjectV2ResponseMPayload {
	this := EzsignsignergroupGetObjectV2ResponseMPayload{}
	this.ObjEzsignsignergroup = objEzsignsignergroup
	return &this
}

// NewEzsignsignergroupGetObjectV2ResponseMPayloadWithDefaults instantiates a new EzsignsignergroupGetObjectV2ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignsignergroupGetObjectV2ResponseMPayloadWithDefaults() *EzsignsignergroupGetObjectV2ResponseMPayload {
	this := EzsignsignergroupGetObjectV2ResponseMPayload{}
	return &this
}

// GetObjEzsignsignergroup returns the ObjEzsignsignergroup field value
func (o *EzsignsignergroupGetObjectV2ResponseMPayload) GetObjEzsignsignergroup() EzsignsignergroupResponseCompound {
	if o == nil {
		var ret EzsignsignergroupResponseCompound
		return ret
	}

	return o.ObjEzsignsignergroup
}

// GetObjEzsignsignergroupOk returns a tuple with the ObjEzsignsignergroup field value
// and a boolean to check if the value has been set.
func (o *EzsignsignergroupGetObjectV2ResponseMPayload) GetObjEzsignsignergroupOk() (*EzsignsignergroupResponseCompound, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjEzsignsignergroup, true
}

// SetObjEzsignsignergroup sets field value
func (o *EzsignsignergroupGetObjectV2ResponseMPayload) SetObjEzsignsignergroup(v EzsignsignergroupResponseCompound) {
	o.ObjEzsignsignergroup = v
}

func (o EzsignsignergroupGetObjectV2ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignsignergroupGetObjectV2ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objEzsignsignergroup"] = o.ObjEzsignsignergroup
	return toSerialize, nil
}

func (o *EzsignsignergroupGetObjectV2ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"objEzsignsignergroup",
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

	varEzsignsignergroupGetObjectV2ResponseMPayload := _EzsignsignergroupGetObjectV2ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsignsignergroupGetObjectV2ResponseMPayload)

	if err != nil {
		return err
	}

	*o = EzsignsignergroupGetObjectV2ResponseMPayload(varEzsignsignergroupGetObjectV2ResponseMPayload)

	return err
}

type NullableEzsignsignergroupGetObjectV2ResponseMPayload struct {
	value *EzsignsignergroupGetObjectV2ResponseMPayload
	isSet bool
}

func (v NullableEzsignsignergroupGetObjectV2ResponseMPayload) Get() *EzsignsignergroupGetObjectV2ResponseMPayload {
	return v.value
}

func (v *NullableEzsignsignergroupGetObjectV2ResponseMPayload) Set(val *EzsignsignergroupGetObjectV2ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignsignergroupGetObjectV2ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignsignergroupGetObjectV2ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignsignergroupGetObjectV2ResponseMPayload(val *EzsignsignergroupGetObjectV2ResponseMPayload) *NullableEzsignsignergroupGetObjectV2ResponseMPayload {
	return &NullableEzsignsignergroupGetObjectV2ResponseMPayload{value: val, isSet: true}
}

func (v NullableEzsignsignergroupGetObjectV2ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignsignergroupGetObjectV2ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


