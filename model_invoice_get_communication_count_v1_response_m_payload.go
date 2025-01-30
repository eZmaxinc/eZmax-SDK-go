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

// checks if the InvoiceGetCommunicationCountV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvoiceGetCommunicationCountV1ResponseMPayload{}

// InvoiceGetCommunicationCountV1ResponseMPayload Response for GET /1/object/invoice/{pkiInvoiceID}/getCommunicationCount
type InvoiceGetCommunicationCountV1ResponseMPayload struct {
	// The count of Communication.
	ICommunicationCount int32 `json:"iCommunicationCount"`
}

type _InvoiceGetCommunicationCountV1ResponseMPayload InvoiceGetCommunicationCountV1ResponseMPayload

// NewInvoiceGetCommunicationCountV1ResponseMPayload instantiates a new InvoiceGetCommunicationCountV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoiceGetCommunicationCountV1ResponseMPayload(iCommunicationCount int32) *InvoiceGetCommunicationCountV1ResponseMPayload {
	this := InvoiceGetCommunicationCountV1ResponseMPayload{}
	this.ICommunicationCount = iCommunicationCount
	return &this
}

// NewInvoiceGetCommunicationCountV1ResponseMPayloadWithDefaults instantiates a new InvoiceGetCommunicationCountV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceGetCommunicationCountV1ResponseMPayloadWithDefaults() *InvoiceGetCommunicationCountV1ResponseMPayload {
	this := InvoiceGetCommunicationCountV1ResponseMPayload{}
	return &this
}

// GetICommunicationCount returns the ICommunicationCount field value
func (o *InvoiceGetCommunicationCountV1ResponseMPayload) GetICommunicationCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ICommunicationCount
}

// GetICommunicationCountOk returns a tuple with the ICommunicationCount field value
// and a boolean to check if the value has been set.
func (o *InvoiceGetCommunicationCountV1ResponseMPayload) GetICommunicationCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ICommunicationCount, true
}

// SetICommunicationCount sets field value
func (o *InvoiceGetCommunicationCountV1ResponseMPayload) SetICommunicationCount(v int32) {
	o.ICommunicationCount = v
}

func (o InvoiceGetCommunicationCountV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvoiceGetCommunicationCountV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["iCommunicationCount"] = o.ICommunicationCount
	return toSerialize, nil
}

func (o *InvoiceGetCommunicationCountV1ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"iCommunicationCount",
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

	varInvoiceGetCommunicationCountV1ResponseMPayload := _InvoiceGetCommunicationCountV1ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInvoiceGetCommunicationCountV1ResponseMPayload)

	if err != nil {
		return err
	}

	*o = InvoiceGetCommunicationCountV1ResponseMPayload(varInvoiceGetCommunicationCountV1ResponseMPayload)

	return err
}

type NullableInvoiceGetCommunicationCountV1ResponseMPayload struct {
	value *InvoiceGetCommunicationCountV1ResponseMPayload
	isSet bool
}

func (v NullableInvoiceGetCommunicationCountV1ResponseMPayload) Get() *InvoiceGetCommunicationCountV1ResponseMPayload {
	return v.value
}

func (v *NullableInvoiceGetCommunicationCountV1ResponseMPayload) Set(val *InvoiceGetCommunicationCountV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoiceGetCommunicationCountV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoiceGetCommunicationCountV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoiceGetCommunicationCountV1ResponseMPayload(val *InvoiceGetCommunicationCountV1ResponseMPayload) *NullableInvoiceGetCommunicationCountV1ResponseMPayload {
	return &NullableInvoiceGetCommunicationCountV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableInvoiceGetCommunicationCountV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoiceGetCommunicationCountV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


