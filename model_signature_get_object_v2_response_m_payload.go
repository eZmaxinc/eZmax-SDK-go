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

// checks if the SignatureGetObjectV2ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SignatureGetObjectV2ResponseMPayload{}

// SignatureGetObjectV2ResponseMPayload Payload for GET /2/object/signature/{pkiSignatureID}
type SignatureGetObjectV2ResponseMPayload struct {
	ObjSignature SignatureResponseCompound `json:"objSignature"`
}

type _SignatureGetObjectV2ResponseMPayload SignatureGetObjectV2ResponseMPayload

// NewSignatureGetObjectV2ResponseMPayload instantiates a new SignatureGetObjectV2ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignatureGetObjectV2ResponseMPayload(objSignature SignatureResponseCompound) *SignatureGetObjectV2ResponseMPayload {
	this := SignatureGetObjectV2ResponseMPayload{}
	this.ObjSignature = objSignature
	return &this
}

// NewSignatureGetObjectV2ResponseMPayloadWithDefaults instantiates a new SignatureGetObjectV2ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignatureGetObjectV2ResponseMPayloadWithDefaults() *SignatureGetObjectV2ResponseMPayload {
	this := SignatureGetObjectV2ResponseMPayload{}
	return &this
}

// GetObjSignature returns the ObjSignature field value
func (o *SignatureGetObjectV2ResponseMPayload) GetObjSignature() SignatureResponseCompound {
	if o == nil {
		var ret SignatureResponseCompound
		return ret
	}

	return o.ObjSignature
}

// GetObjSignatureOk returns a tuple with the ObjSignature field value
// and a boolean to check if the value has been set.
func (o *SignatureGetObjectV2ResponseMPayload) GetObjSignatureOk() (*SignatureResponseCompound, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjSignature, true
}

// SetObjSignature sets field value
func (o *SignatureGetObjectV2ResponseMPayload) SetObjSignature(v SignatureResponseCompound) {
	o.ObjSignature = v
}

func (o SignatureGetObjectV2ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SignatureGetObjectV2ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objSignature"] = o.ObjSignature
	return toSerialize, nil
}

func (o *SignatureGetObjectV2ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"objSignature",
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

	varSignatureGetObjectV2ResponseMPayload := _SignatureGetObjectV2ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSignatureGetObjectV2ResponseMPayload)

	if err != nil {
		return err
	}

	*o = SignatureGetObjectV2ResponseMPayload(varSignatureGetObjectV2ResponseMPayload)

	return err
}

type NullableSignatureGetObjectV2ResponseMPayload struct {
	value *SignatureGetObjectV2ResponseMPayload
	isSet bool
}

func (v NullableSignatureGetObjectV2ResponseMPayload) Get() *SignatureGetObjectV2ResponseMPayload {
	return v.value
}

func (v *NullableSignatureGetObjectV2ResponseMPayload) Set(val *SignatureGetObjectV2ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableSignatureGetObjectV2ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableSignatureGetObjectV2ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignatureGetObjectV2ResponseMPayload(val *SignatureGetObjectV2ResponseMPayload) *NullableSignatureGetObjectV2ResponseMPayload {
	return &NullableSignatureGetObjectV2ResponseMPayload{value: val, isSet: true}
}

func (v NullableSignatureGetObjectV2ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignatureGetObjectV2ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


