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

// checks if the ApikeyGetObjectV2ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApikeyGetObjectV2ResponseMPayload{}

// ApikeyGetObjectV2ResponseMPayload Payload for GET /2/object/apikey/{pkiApikeyID}
type ApikeyGetObjectV2ResponseMPayload struct {
	ObjApikey ApikeyResponseCompound `json:"objApikey"`
}

type _ApikeyGetObjectV2ResponseMPayload ApikeyGetObjectV2ResponseMPayload

// NewApikeyGetObjectV2ResponseMPayload instantiates a new ApikeyGetObjectV2ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApikeyGetObjectV2ResponseMPayload(objApikey ApikeyResponseCompound) *ApikeyGetObjectV2ResponseMPayload {
	this := ApikeyGetObjectV2ResponseMPayload{}
	this.ObjApikey = objApikey
	return &this
}

// NewApikeyGetObjectV2ResponseMPayloadWithDefaults instantiates a new ApikeyGetObjectV2ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApikeyGetObjectV2ResponseMPayloadWithDefaults() *ApikeyGetObjectV2ResponseMPayload {
	this := ApikeyGetObjectV2ResponseMPayload{}
	return &this
}

// GetObjApikey returns the ObjApikey field value
func (o *ApikeyGetObjectV2ResponseMPayload) GetObjApikey() ApikeyResponseCompound {
	if o == nil {
		var ret ApikeyResponseCompound
		return ret
	}

	return o.ObjApikey
}

// GetObjApikeyOk returns a tuple with the ObjApikey field value
// and a boolean to check if the value has been set.
func (o *ApikeyGetObjectV2ResponseMPayload) GetObjApikeyOk() (*ApikeyResponseCompound, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjApikey, true
}

// SetObjApikey sets field value
func (o *ApikeyGetObjectV2ResponseMPayload) SetObjApikey(v ApikeyResponseCompound) {
	o.ObjApikey = v
}

func (o ApikeyGetObjectV2ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApikeyGetObjectV2ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objApikey"] = o.ObjApikey
	return toSerialize, nil
}

func (o *ApikeyGetObjectV2ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"objApikey",
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

	varApikeyGetObjectV2ResponseMPayload := _ApikeyGetObjectV2ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApikeyGetObjectV2ResponseMPayload)

	if err != nil {
		return err
	}

	*o = ApikeyGetObjectV2ResponseMPayload(varApikeyGetObjectV2ResponseMPayload)

	return err
}

type NullableApikeyGetObjectV2ResponseMPayload struct {
	value *ApikeyGetObjectV2ResponseMPayload
	isSet bool
}

func (v NullableApikeyGetObjectV2ResponseMPayload) Get() *ApikeyGetObjectV2ResponseMPayload {
	return v.value
}

func (v *NullableApikeyGetObjectV2ResponseMPayload) Set(val *ApikeyGetObjectV2ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableApikeyGetObjectV2ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableApikeyGetObjectV2ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApikeyGetObjectV2ResponseMPayload(val *ApikeyGetObjectV2ResponseMPayload) *NullableApikeyGetObjectV2ResponseMPayload {
	return &NullableApikeyGetObjectV2ResponseMPayload{value: val, isSet: true}
}

func (v NullableApikeyGetObjectV2ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApikeyGetObjectV2ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


