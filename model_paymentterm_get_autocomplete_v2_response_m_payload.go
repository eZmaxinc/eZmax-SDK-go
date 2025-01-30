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

// checks if the PaymenttermGetAutocompleteV2ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymenttermGetAutocompleteV2ResponseMPayload{}

// PaymenttermGetAutocompleteV2ResponseMPayload Payload for POST /2/object/paymentterm/getAutocomplete
type PaymenttermGetAutocompleteV2ResponseMPayload struct {
	// An array of Paymentterm autocomplete element response.
	AObjPaymentterm []PaymenttermAutocompleteElementResponse `json:"a_objPaymentterm"`
}

type _PaymenttermGetAutocompleteV2ResponseMPayload PaymenttermGetAutocompleteV2ResponseMPayload

// NewPaymenttermGetAutocompleteV2ResponseMPayload instantiates a new PaymenttermGetAutocompleteV2ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymenttermGetAutocompleteV2ResponseMPayload(aObjPaymentterm []PaymenttermAutocompleteElementResponse) *PaymenttermGetAutocompleteV2ResponseMPayload {
	this := PaymenttermGetAutocompleteV2ResponseMPayload{}
	this.AObjPaymentterm = aObjPaymentterm
	return &this
}

// NewPaymenttermGetAutocompleteV2ResponseMPayloadWithDefaults instantiates a new PaymenttermGetAutocompleteV2ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymenttermGetAutocompleteV2ResponseMPayloadWithDefaults() *PaymenttermGetAutocompleteV2ResponseMPayload {
	this := PaymenttermGetAutocompleteV2ResponseMPayload{}
	return &this
}

// GetAObjPaymentterm returns the AObjPaymentterm field value
func (o *PaymenttermGetAutocompleteV2ResponseMPayload) GetAObjPaymentterm() []PaymenttermAutocompleteElementResponse {
	if o == nil {
		var ret []PaymenttermAutocompleteElementResponse
		return ret
	}

	return o.AObjPaymentterm
}

// GetAObjPaymenttermOk returns a tuple with the AObjPaymentterm field value
// and a boolean to check if the value has been set.
func (o *PaymenttermGetAutocompleteV2ResponseMPayload) GetAObjPaymenttermOk() ([]PaymenttermAutocompleteElementResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjPaymentterm, true
}

// SetAObjPaymentterm sets field value
func (o *PaymenttermGetAutocompleteV2ResponseMPayload) SetAObjPaymentterm(v []PaymenttermAutocompleteElementResponse) {
	o.AObjPaymentterm = v
}

func (o PaymenttermGetAutocompleteV2ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymenttermGetAutocompleteV2ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_objPaymentterm"] = o.AObjPaymentterm
	return toSerialize, nil
}

func (o *PaymenttermGetAutocompleteV2ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
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

	varPaymenttermGetAutocompleteV2ResponseMPayload := _PaymenttermGetAutocompleteV2ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPaymenttermGetAutocompleteV2ResponseMPayload)

	if err != nil {
		return err
	}

	*o = PaymenttermGetAutocompleteV2ResponseMPayload(varPaymenttermGetAutocompleteV2ResponseMPayload)

	return err
}

type NullablePaymenttermGetAutocompleteV2ResponseMPayload struct {
	value *PaymenttermGetAutocompleteV2ResponseMPayload
	isSet bool
}

func (v NullablePaymenttermGetAutocompleteV2ResponseMPayload) Get() *PaymenttermGetAutocompleteV2ResponseMPayload {
	return v.value
}

func (v *NullablePaymenttermGetAutocompleteV2ResponseMPayload) Set(val *PaymenttermGetAutocompleteV2ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymenttermGetAutocompleteV2ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymenttermGetAutocompleteV2ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymenttermGetAutocompleteV2ResponseMPayload(val *PaymenttermGetAutocompleteV2ResponseMPayload) *NullablePaymenttermGetAutocompleteV2ResponseMPayload {
	return &NullablePaymenttermGetAutocompleteV2ResponseMPayload{value: val, isSet: true}
}

func (v NullablePaymenttermGetAutocompleteV2ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymenttermGetAutocompleteV2ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


