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

// checks if the SystemconfigurationGetObjectV2ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SystemconfigurationGetObjectV2ResponseMPayload{}

// SystemconfigurationGetObjectV2ResponseMPayload Payload for GET /2/object/systemconfiguration/{pkiSystemconfigurationID}
type SystemconfigurationGetObjectV2ResponseMPayload struct {
	ObjSystemconfiguration SystemconfigurationResponseCompound `json:"objSystemconfiguration"`
}

type _SystemconfigurationGetObjectV2ResponseMPayload SystemconfigurationGetObjectV2ResponseMPayload

// NewSystemconfigurationGetObjectV2ResponseMPayload instantiates a new SystemconfigurationGetObjectV2ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemconfigurationGetObjectV2ResponseMPayload(objSystemconfiguration SystemconfigurationResponseCompound) *SystemconfigurationGetObjectV2ResponseMPayload {
	this := SystemconfigurationGetObjectV2ResponseMPayload{}
	this.ObjSystemconfiguration = objSystemconfiguration
	return &this
}

// NewSystemconfigurationGetObjectV2ResponseMPayloadWithDefaults instantiates a new SystemconfigurationGetObjectV2ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemconfigurationGetObjectV2ResponseMPayloadWithDefaults() *SystemconfigurationGetObjectV2ResponseMPayload {
	this := SystemconfigurationGetObjectV2ResponseMPayload{}
	return &this
}

// GetObjSystemconfiguration returns the ObjSystemconfiguration field value
func (o *SystemconfigurationGetObjectV2ResponseMPayload) GetObjSystemconfiguration() SystemconfigurationResponseCompound {
	if o == nil {
		var ret SystemconfigurationResponseCompound
		return ret
	}

	return o.ObjSystemconfiguration
}

// GetObjSystemconfigurationOk returns a tuple with the ObjSystemconfiguration field value
// and a boolean to check if the value has been set.
func (o *SystemconfigurationGetObjectV2ResponseMPayload) GetObjSystemconfigurationOk() (*SystemconfigurationResponseCompound, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjSystemconfiguration, true
}

// SetObjSystemconfiguration sets field value
func (o *SystemconfigurationGetObjectV2ResponseMPayload) SetObjSystemconfiguration(v SystemconfigurationResponseCompound) {
	o.ObjSystemconfiguration = v
}

func (o SystemconfigurationGetObjectV2ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SystemconfigurationGetObjectV2ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objSystemconfiguration"] = o.ObjSystemconfiguration
	return toSerialize, nil
}

func (o *SystemconfigurationGetObjectV2ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"objSystemconfiguration",
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

	varSystemconfigurationGetObjectV2ResponseMPayload := _SystemconfigurationGetObjectV2ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSystemconfigurationGetObjectV2ResponseMPayload)

	if err != nil {
		return err
	}

	*o = SystemconfigurationGetObjectV2ResponseMPayload(varSystemconfigurationGetObjectV2ResponseMPayload)

	return err
}

type NullableSystemconfigurationGetObjectV2ResponseMPayload struct {
	value *SystemconfigurationGetObjectV2ResponseMPayload
	isSet bool
}

func (v NullableSystemconfigurationGetObjectV2ResponseMPayload) Get() *SystemconfigurationGetObjectV2ResponseMPayload {
	return v.value
}

func (v *NullableSystemconfigurationGetObjectV2ResponseMPayload) Set(val *SystemconfigurationGetObjectV2ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemconfigurationGetObjectV2ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemconfigurationGetObjectV2ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemconfigurationGetObjectV2ResponseMPayload(val *SystemconfigurationGetObjectV2ResponseMPayload) *NullableSystemconfigurationGetObjectV2ResponseMPayload {
	return &NullableSystemconfigurationGetObjectV2ResponseMPayload{value: val, isSet: true}
}

func (v NullableSystemconfigurationGetObjectV2ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemconfigurationGetObjectV2ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


