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

// checks if the CreditcardclientPatchObjectV1Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreditcardclientPatchObjectV1Response{}

// CreditcardclientPatchObjectV1Response Response for PATCH /1/object/creditcardclient/{pkiCreditcardclientID}
type CreditcardclientPatchObjectV1Response struct {
	CommonResponse
}

type _CreditcardclientPatchObjectV1Response CreditcardclientPatchObjectV1Response

// NewCreditcardclientPatchObjectV1Response instantiates a new CreditcardclientPatchObjectV1Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditcardclientPatchObjectV1Response(objDebugPayload CommonResponseObjDebugPayload) *CreditcardclientPatchObjectV1Response {
	this := CreditcardclientPatchObjectV1Response{}
	this.ObjDebugPayload = objDebugPayload
	return &this
}

// NewCreditcardclientPatchObjectV1ResponseWithDefaults instantiates a new CreditcardclientPatchObjectV1Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditcardclientPatchObjectV1ResponseWithDefaults() *CreditcardclientPatchObjectV1Response {
	this := CreditcardclientPatchObjectV1Response{}
	return &this
}

func (o CreditcardclientPatchObjectV1Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreditcardclientPatchObjectV1Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

func (o *CreditcardclientPatchObjectV1Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"objDebugPayload",
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

	varCreditcardclientPatchObjectV1Response := _CreditcardclientPatchObjectV1Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreditcardclientPatchObjectV1Response)

	if err != nil {
		return err
	}

	*o = CreditcardclientPatchObjectV1Response(varCreditcardclientPatchObjectV1Response)

	return err
}

type NullableCreditcardclientPatchObjectV1Response struct {
	value *CreditcardclientPatchObjectV1Response
	isSet bool
}

func (v NullableCreditcardclientPatchObjectV1Response) Get() *CreditcardclientPatchObjectV1Response {
	return v.value
}

func (v *NullableCreditcardclientPatchObjectV1Response) Set(val *CreditcardclientPatchObjectV1Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditcardclientPatchObjectV1Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditcardclientPatchObjectV1Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditcardclientPatchObjectV1Response(val *CreditcardclientPatchObjectV1Response) *NullableCreditcardclientPatchObjectV1Response {
	return &NullableCreditcardclientPatchObjectV1Response{value: val, isSet: true}
}

func (v NullableCreditcardclientPatchObjectV1Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditcardclientPatchObjectV1Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


