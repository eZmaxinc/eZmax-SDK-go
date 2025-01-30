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

// checks if the BuyercontractGetCommunicationCountV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BuyercontractGetCommunicationCountV1ResponseMPayload{}

// BuyercontractGetCommunicationCountV1ResponseMPayload Response for GET /1/object/buyercontract/{pkiBuyercontractID}/getCommunicationCount
type BuyercontractGetCommunicationCountV1ResponseMPayload struct {
	// The count of Communication.
	ICommunicationCount int32 `json:"iCommunicationCount"`
}

type _BuyercontractGetCommunicationCountV1ResponseMPayload BuyercontractGetCommunicationCountV1ResponseMPayload

// NewBuyercontractGetCommunicationCountV1ResponseMPayload instantiates a new BuyercontractGetCommunicationCountV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuyercontractGetCommunicationCountV1ResponseMPayload(iCommunicationCount int32) *BuyercontractGetCommunicationCountV1ResponseMPayload {
	this := BuyercontractGetCommunicationCountV1ResponseMPayload{}
	this.ICommunicationCount = iCommunicationCount
	return &this
}

// NewBuyercontractGetCommunicationCountV1ResponseMPayloadWithDefaults instantiates a new BuyercontractGetCommunicationCountV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBuyercontractGetCommunicationCountV1ResponseMPayloadWithDefaults() *BuyercontractGetCommunicationCountV1ResponseMPayload {
	this := BuyercontractGetCommunicationCountV1ResponseMPayload{}
	return &this
}

// GetICommunicationCount returns the ICommunicationCount field value
func (o *BuyercontractGetCommunicationCountV1ResponseMPayload) GetICommunicationCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ICommunicationCount
}

// GetICommunicationCountOk returns a tuple with the ICommunicationCount field value
// and a boolean to check if the value has been set.
func (o *BuyercontractGetCommunicationCountV1ResponseMPayload) GetICommunicationCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ICommunicationCount, true
}

// SetICommunicationCount sets field value
func (o *BuyercontractGetCommunicationCountV1ResponseMPayload) SetICommunicationCount(v int32) {
	o.ICommunicationCount = v
}

func (o BuyercontractGetCommunicationCountV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BuyercontractGetCommunicationCountV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["iCommunicationCount"] = o.ICommunicationCount
	return toSerialize, nil
}

func (o *BuyercontractGetCommunicationCountV1ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
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

	varBuyercontractGetCommunicationCountV1ResponseMPayload := _BuyercontractGetCommunicationCountV1ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBuyercontractGetCommunicationCountV1ResponseMPayload)

	if err != nil {
		return err
	}

	*o = BuyercontractGetCommunicationCountV1ResponseMPayload(varBuyercontractGetCommunicationCountV1ResponseMPayload)

	return err
}

type NullableBuyercontractGetCommunicationCountV1ResponseMPayload struct {
	value *BuyercontractGetCommunicationCountV1ResponseMPayload
	isSet bool
}

func (v NullableBuyercontractGetCommunicationCountV1ResponseMPayload) Get() *BuyercontractGetCommunicationCountV1ResponseMPayload {
	return v.value
}

func (v *NullableBuyercontractGetCommunicationCountV1ResponseMPayload) Set(val *BuyercontractGetCommunicationCountV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableBuyercontractGetCommunicationCountV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableBuyercontractGetCommunicationCountV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuyercontractGetCommunicationCountV1ResponseMPayload(val *BuyercontractGetCommunicationCountV1ResponseMPayload) *NullableBuyercontractGetCommunicationCountV1ResponseMPayload {
	return &NullableBuyercontractGetCommunicationCountV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableBuyercontractGetCommunicationCountV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuyercontractGetCommunicationCountV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


