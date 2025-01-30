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

// checks if the PaymenttermGetListV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymenttermGetListV1ResponseMPayload{}

// PaymenttermGetListV1ResponseMPayload Payload for GET /1/object/paymentterm/getList
type PaymenttermGetListV1ResponseMPayload struct {
	// The number of rows returned
	IRowReturned int32 `json:"iRowReturned"`
	// The number of rows matching your filters (if any) or the total number of rows
	IRowFiltered int32 `json:"iRowFiltered"`
	AObjPaymentterm []PaymenttermListElement `json:"a_objPaymentterm"`
}

type _PaymenttermGetListV1ResponseMPayload PaymenttermGetListV1ResponseMPayload

// NewPaymenttermGetListV1ResponseMPayload instantiates a new PaymenttermGetListV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymenttermGetListV1ResponseMPayload(iRowReturned int32, iRowFiltered int32, aObjPaymentterm []PaymenttermListElement) *PaymenttermGetListV1ResponseMPayload {
	this := PaymenttermGetListV1ResponseMPayload{}
	this.IRowReturned = iRowReturned
	this.IRowFiltered = iRowFiltered
	this.AObjPaymentterm = aObjPaymentterm
	return &this
}

// NewPaymenttermGetListV1ResponseMPayloadWithDefaults instantiates a new PaymenttermGetListV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymenttermGetListV1ResponseMPayloadWithDefaults() *PaymenttermGetListV1ResponseMPayload {
	this := PaymenttermGetListV1ResponseMPayload{}
	return &this
}

// GetIRowReturned returns the IRowReturned field value
func (o *PaymenttermGetListV1ResponseMPayload) GetIRowReturned() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IRowReturned
}

// GetIRowReturnedOk returns a tuple with the IRowReturned field value
// and a boolean to check if the value has been set.
func (o *PaymenttermGetListV1ResponseMPayload) GetIRowReturnedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IRowReturned, true
}

// SetIRowReturned sets field value
func (o *PaymenttermGetListV1ResponseMPayload) SetIRowReturned(v int32) {
	o.IRowReturned = v
}

// GetIRowFiltered returns the IRowFiltered field value
func (o *PaymenttermGetListV1ResponseMPayload) GetIRowFiltered() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IRowFiltered
}

// GetIRowFilteredOk returns a tuple with the IRowFiltered field value
// and a boolean to check if the value has been set.
func (o *PaymenttermGetListV1ResponseMPayload) GetIRowFilteredOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IRowFiltered, true
}

// SetIRowFiltered sets field value
func (o *PaymenttermGetListV1ResponseMPayload) SetIRowFiltered(v int32) {
	o.IRowFiltered = v
}

// GetAObjPaymentterm returns the AObjPaymentterm field value
func (o *PaymenttermGetListV1ResponseMPayload) GetAObjPaymentterm() []PaymenttermListElement {
	if o == nil {
		var ret []PaymenttermListElement
		return ret
	}

	return o.AObjPaymentterm
}

// GetAObjPaymenttermOk returns a tuple with the AObjPaymentterm field value
// and a boolean to check if the value has been set.
func (o *PaymenttermGetListV1ResponseMPayload) GetAObjPaymenttermOk() ([]PaymenttermListElement, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjPaymentterm, true
}

// SetAObjPaymentterm sets field value
func (o *PaymenttermGetListV1ResponseMPayload) SetAObjPaymentterm(v []PaymenttermListElement) {
	o.AObjPaymentterm = v
}

func (o PaymenttermGetListV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymenttermGetListV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["iRowReturned"] = o.IRowReturned
	toSerialize["iRowFiltered"] = o.IRowFiltered
	toSerialize["a_objPaymentterm"] = o.AObjPaymentterm
	return toSerialize, nil
}

func (o *PaymenttermGetListV1ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"iRowReturned",
		"iRowFiltered",
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

	varPaymenttermGetListV1ResponseMPayload := _PaymenttermGetListV1ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPaymenttermGetListV1ResponseMPayload)

	if err != nil {
		return err
	}

	*o = PaymenttermGetListV1ResponseMPayload(varPaymenttermGetListV1ResponseMPayload)

	return err
}

type NullablePaymenttermGetListV1ResponseMPayload struct {
	value *PaymenttermGetListV1ResponseMPayload
	isSet bool
}

func (v NullablePaymenttermGetListV1ResponseMPayload) Get() *PaymenttermGetListV1ResponseMPayload {
	return v.value
}

func (v *NullablePaymenttermGetListV1ResponseMPayload) Set(val *PaymenttermGetListV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymenttermGetListV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymenttermGetListV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymenttermGetListV1ResponseMPayload(val *PaymenttermGetListV1ResponseMPayload) *NullablePaymenttermGetListV1ResponseMPayload {
	return &NullablePaymenttermGetListV1ResponseMPayload{value: val, isSet: true}
}

func (v NullablePaymenttermGetListV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymenttermGetListV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


