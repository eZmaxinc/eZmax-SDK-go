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

// checks if the BillingentityinternalCreateObjectV1Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BillingentityinternalCreateObjectV1Response{}

// BillingentityinternalCreateObjectV1Response Response for POST /1/object/billingentityinternal
type BillingentityinternalCreateObjectV1Response struct {
	CommonResponse
	MPayload BillingentityinternalCreateObjectV1ResponseMPayload `json:"mPayload"`
}

type _BillingentityinternalCreateObjectV1Response BillingentityinternalCreateObjectV1Response

// NewBillingentityinternalCreateObjectV1Response instantiates a new BillingentityinternalCreateObjectV1Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingentityinternalCreateObjectV1Response(mPayload BillingentityinternalCreateObjectV1ResponseMPayload, objDebugPayload CommonResponseObjDebugPayload) *BillingentityinternalCreateObjectV1Response {
	this := BillingentityinternalCreateObjectV1Response{}
	this.ObjDebugPayload = objDebugPayload
	this.MPayload = mPayload
	return &this
}

// NewBillingentityinternalCreateObjectV1ResponseWithDefaults instantiates a new BillingentityinternalCreateObjectV1Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingentityinternalCreateObjectV1ResponseWithDefaults() *BillingentityinternalCreateObjectV1Response {
	this := BillingentityinternalCreateObjectV1Response{}
	return &this
}

// GetMPayload returns the MPayload field value
func (o *BillingentityinternalCreateObjectV1Response) GetMPayload() BillingentityinternalCreateObjectV1ResponseMPayload {
	if o == nil {
		var ret BillingentityinternalCreateObjectV1ResponseMPayload
		return ret
	}

	return o.MPayload
}

// GetMPayloadOk returns a tuple with the MPayload field value
// and a boolean to check if the value has been set.
func (o *BillingentityinternalCreateObjectV1Response) GetMPayloadOk() (*BillingentityinternalCreateObjectV1ResponseMPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MPayload, true
}

// SetMPayload sets field value
func (o *BillingentityinternalCreateObjectV1Response) SetMPayload(v BillingentityinternalCreateObjectV1ResponseMPayload) {
	o.MPayload = v
}

func (o BillingentityinternalCreateObjectV1Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BillingentityinternalCreateObjectV1Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mPayload"] = o.MPayload
	return toSerialize, nil
}

func (o *BillingentityinternalCreateObjectV1Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"mPayload",
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

	varBillingentityinternalCreateObjectV1Response := _BillingentityinternalCreateObjectV1Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBillingentityinternalCreateObjectV1Response)

	if err != nil {
		return err
	}

	*o = BillingentityinternalCreateObjectV1Response(varBillingentityinternalCreateObjectV1Response)

	return err
}

type NullableBillingentityinternalCreateObjectV1Response struct {
	value *BillingentityinternalCreateObjectV1Response
	isSet bool
}

func (v NullableBillingentityinternalCreateObjectV1Response) Get() *BillingentityinternalCreateObjectV1Response {
	return v.value
}

func (v *NullableBillingentityinternalCreateObjectV1Response) Set(val *BillingentityinternalCreateObjectV1Response) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingentityinternalCreateObjectV1Response) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingentityinternalCreateObjectV1Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingentityinternalCreateObjectV1Response(val *BillingentityinternalCreateObjectV1Response) *NullableBillingentityinternalCreateObjectV1Response {
	return &NullableBillingentityinternalCreateObjectV1Response{value: val, isSet: true}
}

func (v NullableBillingentityinternalCreateObjectV1Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingentityinternalCreateObjectV1Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


