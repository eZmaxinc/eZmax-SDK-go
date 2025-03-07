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

// checks if the BillingentityinternalGetObjectV2ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BillingentityinternalGetObjectV2ResponseMPayload{}

// BillingentityinternalGetObjectV2ResponseMPayload Payload for GET /2/object/billingentityinternal/{pkiBillingentityinternalID}
type BillingentityinternalGetObjectV2ResponseMPayload struct {
	ObjBillingentityinternal BillingentityinternalResponseCompound `json:"objBillingentityinternal"`
}

type _BillingentityinternalGetObjectV2ResponseMPayload BillingentityinternalGetObjectV2ResponseMPayload

// NewBillingentityinternalGetObjectV2ResponseMPayload instantiates a new BillingentityinternalGetObjectV2ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingentityinternalGetObjectV2ResponseMPayload(objBillingentityinternal BillingentityinternalResponseCompound) *BillingentityinternalGetObjectV2ResponseMPayload {
	this := BillingentityinternalGetObjectV2ResponseMPayload{}
	this.ObjBillingentityinternal = objBillingentityinternal
	return &this
}

// NewBillingentityinternalGetObjectV2ResponseMPayloadWithDefaults instantiates a new BillingentityinternalGetObjectV2ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingentityinternalGetObjectV2ResponseMPayloadWithDefaults() *BillingentityinternalGetObjectV2ResponseMPayload {
	this := BillingentityinternalGetObjectV2ResponseMPayload{}
	return &this
}

// GetObjBillingentityinternal returns the ObjBillingentityinternal field value
func (o *BillingentityinternalGetObjectV2ResponseMPayload) GetObjBillingentityinternal() BillingentityinternalResponseCompound {
	if o == nil {
		var ret BillingentityinternalResponseCompound
		return ret
	}

	return o.ObjBillingentityinternal
}

// GetObjBillingentityinternalOk returns a tuple with the ObjBillingentityinternal field value
// and a boolean to check if the value has been set.
func (o *BillingentityinternalGetObjectV2ResponseMPayload) GetObjBillingentityinternalOk() (*BillingentityinternalResponseCompound, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjBillingentityinternal, true
}

// SetObjBillingentityinternal sets field value
func (o *BillingentityinternalGetObjectV2ResponseMPayload) SetObjBillingentityinternal(v BillingentityinternalResponseCompound) {
	o.ObjBillingentityinternal = v
}

func (o BillingentityinternalGetObjectV2ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BillingentityinternalGetObjectV2ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objBillingentityinternal"] = o.ObjBillingentityinternal
	return toSerialize, nil
}

func (o *BillingentityinternalGetObjectV2ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"objBillingentityinternal",
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

	varBillingentityinternalGetObjectV2ResponseMPayload := _BillingentityinternalGetObjectV2ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBillingentityinternalGetObjectV2ResponseMPayload)

	if err != nil {
		return err
	}

	*o = BillingentityinternalGetObjectV2ResponseMPayload(varBillingentityinternalGetObjectV2ResponseMPayload)

	return err
}

type NullableBillingentityinternalGetObjectV2ResponseMPayload struct {
	value *BillingentityinternalGetObjectV2ResponseMPayload
	isSet bool
}

func (v NullableBillingentityinternalGetObjectV2ResponseMPayload) Get() *BillingentityinternalGetObjectV2ResponseMPayload {
	return v.value
}

func (v *NullableBillingentityinternalGetObjectV2ResponseMPayload) Set(val *BillingentityinternalGetObjectV2ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingentityinternalGetObjectV2ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingentityinternalGetObjectV2ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingentityinternalGetObjectV2ResponseMPayload(val *BillingentityinternalGetObjectV2ResponseMPayload) *NullableBillingentityinternalGetObjectV2ResponseMPayload {
	return &NullableBillingentityinternalGetObjectV2ResponseMPayload{value: val, isSet: true}
}

func (v NullableBillingentityinternalGetObjectV2ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingentityinternalGetObjectV2ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


