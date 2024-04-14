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

// checks if the PaymenttermCreateObjectV1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymenttermCreateObjectV1Request{}

// PaymenttermCreateObjectV1Request Request for POST /1/object/paymentterm
type PaymenttermCreateObjectV1Request struct {
	AObjPaymentterm []PaymenttermRequestCompound `json:"a_objPaymentterm"`
}

type _PaymenttermCreateObjectV1Request PaymenttermCreateObjectV1Request

// NewPaymenttermCreateObjectV1Request instantiates a new PaymenttermCreateObjectV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymenttermCreateObjectV1Request(aObjPaymentterm []PaymenttermRequestCompound) *PaymenttermCreateObjectV1Request {
	this := PaymenttermCreateObjectV1Request{}
	this.AObjPaymentterm = aObjPaymentterm
	return &this
}

// NewPaymenttermCreateObjectV1RequestWithDefaults instantiates a new PaymenttermCreateObjectV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymenttermCreateObjectV1RequestWithDefaults() *PaymenttermCreateObjectV1Request {
	this := PaymenttermCreateObjectV1Request{}
	return &this
}

// GetAObjPaymentterm returns the AObjPaymentterm field value
func (o *PaymenttermCreateObjectV1Request) GetAObjPaymentterm() []PaymenttermRequestCompound {
	if o == nil {
		var ret []PaymenttermRequestCompound
		return ret
	}

	return o.AObjPaymentterm
}

// GetAObjPaymenttermOk returns a tuple with the AObjPaymentterm field value
// and a boolean to check if the value has been set.
func (o *PaymenttermCreateObjectV1Request) GetAObjPaymenttermOk() ([]PaymenttermRequestCompound, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjPaymentterm, true
}

// SetAObjPaymentterm sets field value
func (o *PaymenttermCreateObjectV1Request) SetAObjPaymentterm(v []PaymenttermRequestCompound) {
	o.AObjPaymentterm = v
}

func (o PaymenttermCreateObjectV1Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymenttermCreateObjectV1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_objPaymentterm"] = o.AObjPaymentterm
	return toSerialize, nil
}

func (o *PaymenttermCreateObjectV1Request) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"a_objPaymentterm",
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

	varPaymenttermCreateObjectV1Request := _PaymenttermCreateObjectV1Request{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPaymenttermCreateObjectV1Request)

	if err != nil {
		return err
	}

	*o = PaymenttermCreateObjectV1Request(varPaymenttermCreateObjectV1Request)

	return err
}

type NullablePaymenttermCreateObjectV1Request struct {
	value *PaymenttermCreateObjectV1Request
	isSet bool
}

func (v NullablePaymenttermCreateObjectV1Request) Get() *PaymenttermCreateObjectV1Request {
	return v.value
}

func (v *NullablePaymenttermCreateObjectV1Request) Set(val *PaymenttermCreateObjectV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymenttermCreateObjectV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymenttermCreateObjectV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymenttermCreateObjectV1Request(val *PaymenttermCreateObjectV1Request) *NullablePaymenttermCreateObjectV1Request {
	return &NullablePaymenttermCreateObjectV1Request{value: val, isSet: true}
}

func (v NullablePaymenttermCreateObjectV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymenttermCreateObjectV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


