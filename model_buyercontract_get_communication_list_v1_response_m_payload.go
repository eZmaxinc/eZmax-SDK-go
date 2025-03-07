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

// checks if the BuyercontractGetCommunicationListV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BuyercontractGetCommunicationListV1ResponseMPayload{}

// BuyercontractGetCommunicationListV1ResponseMPayload Response for GET /1/object/buyercontract/{pkiBuyercontractID}/getCommunicationList
type BuyercontractGetCommunicationListV1ResponseMPayload struct {
	AObjCommunication []CustomCommunicationListElementResponse `json:"a_objCommunication"`
}

type _BuyercontractGetCommunicationListV1ResponseMPayload BuyercontractGetCommunicationListV1ResponseMPayload

// NewBuyercontractGetCommunicationListV1ResponseMPayload instantiates a new BuyercontractGetCommunicationListV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuyercontractGetCommunicationListV1ResponseMPayload(aObjCommunication []CustomCommunicationListElementResponse) *BuyercontractGetCommunicationListV1ResponseMPayload {
	this := BuyercontractGetCommunicationListV1ResponseMPayload{}
	this.AObjCommunication = aObjCommunication
	return &this
}

// NewBuyercontractGetCommunicationListV1ResponseMPayloadWithDefaults instantiates a new BuyercontractGetCommunicationListV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBuyercontractGetCommunicationListV1ResponseMPayloadWithDefaults() *BuyercontractGetCommunicationListV1ResponseMPayload {
	this := BuyercontractGetCommunicationListV1ResponseMPayload{}
	return &this
}

// GetAObjCommunication returns the AObjCommunication field value
func (o *BuyercontractGetCommunicationListV1ResponseMPayload) GetAObjCommunication() []CustomCommunicationListElementResponse {
	if o == nil {
		var ret []CustomCommunicationListElementResponse
		return ret
	}

	return o.AObjCommunication
}

// GetAObjCommunicationOk returns a tuple with the AObjCommunication field value
// and a boolean to check if the value has been set.
func (o *BuyercontractGetCommunicationListV1ResponseMPayload) GetAObjCommunicationOk() ([]CustomCommunicationListElementResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjCommunication, true
}

// SetAObjCommunication sets field value
func (o *BuyercontractGetCommunicationListV1ResponseMPayload) SetAObjCommunication(v []CustomCommunicationListElementResponse) {
	o.AObjCommunication = v
}

func (o BuyercontractGetCommunicationListV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BuyercontractGetCommunicationListV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_objCommunication"] = o.AObjCommunication
	return toSerialize, nil
}

func (o *BuyercontractGetCommunicationListV1ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"a_objCommunication",
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

	varBuyercontractGetCommunicationListV1ResponseMPayload := _BuyercontractGetCommunicationListV1ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBuyercontractGetCommunicationListV1ResponseMPayload)

	if err != nil {
		return err
	}

	*o = BuyercontractGetCommunicationListV1ResponseMPayload(varBuyercontractGetCommunicationListV1ResponseMPayload)

	return err
}

type NullableBuyercontractGetCommunicationListV1ResponseMPayload struct {
	value *BuyercontractGetCommunicationListV1ResponseMPayload
	isSet bool
}

func (v NullableBuyercontractGetCommunicationListV1ResponseMPayload) Get() *BuyercontractGetCommunicationListV1ResponseMPayload {
	return v.value
}

func (v *NullableBuyercontractGetCommunicationListV1ResponseMPayload) Set(val *BuyercontractGetCommunicationListV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableBuyercontractGetCommunicationListV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableBuyercontractGetCommunicationListV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuyercontractGetCommunicationListV1ResponseMPayload(val *BuyercontractGetCommunicationListV1ResponseMPayload) *NullableBuyercontractGetCommunicationListV1ResponseMPayload {
	return &NullableBuyercontractGetCommunicationListV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableBuyercontractGetCommunicationListV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuyercontractGetCommunicationListV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


