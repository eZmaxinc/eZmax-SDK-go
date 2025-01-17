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

// checks if the TranqcontractGetCommunicationListV1Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TranqcontractGetCommunicationListV1Response{}

// TranqcontractGetCommunicationListV1Response Response for GET /1/object/tranqcontract/{pkiTranqcontractID}/getCommunicationList
type TranqcontractGetCommunicationListV1Response struct {
	CommonResponseGetList
	MPayload TranqcontractGetCommunicationListV1ResponseMPayload `json:"mPayload"`
}

type _TranqcontractGetCommunicationListV1Response TranqcontractGetCommunicationListV1Response

// NewTranqcontractGetCommunicationListV1Response instantiates a new TranqcontractGetCommunicationListV1Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTranqcontractGetCommunicationListV1Response(mPayload TranqcontractGetCommunicationListV1ResponseMPayload, objDebugPayload CommonResponseObjDebugPayloadGetList) *TranqcontractGetCommunicationListV1Response {
	this := TranqcontractGetCommunicationListV1Response{}
	this.ObjDebugPayload = objDebugPayload
	this.MPayload = mPayload
	return &this
}

// NewTranqcontractGetCommunicationListV1ResponseWithDefaults instantiates a new TranqcontractGetCommunicationListV1Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTranqcontractGetCommunicationListV1ResponseWithDefaults() *TranqcontractGetCommunicationListV1Response {
	this := TranqcontractGetCommunicationListV1Response{}
	return &this
}

// GetMPayload returns the MPayload field value
func (o *TranqcontractGetCommunicationListV1Response) GetMPayload() TranqcontractGetCommunicationListV1ResponseMPayload {
	if o == nil {
		var ret TranqcontractGetCommunicationListV1ResponseMPayload
		return ret
	}

	return o.MPayload
}

// GetMPayloadOk returns a tuple with the MPayload field value
// and a boolean to check if the value has been set.
func (o *TranqcontractGetCommunicationListV1Response) GetMPayloadOk() (*TranqcontractGetCommunicationListV1ResponseMPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MPayload, true
}

// SetMPayload sets field value
func (o *TranqcontractGetCommunicationListV1Response) SetMPayload(v TranqcontractGetCommunicationListV1ResponseMPayload) {
	o.MPayload = v
}

func (o TranqcontractGetCommunicationListV1Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TranqcontractGetCommunicationListV1Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mPayload"] = o.MPayload
	return toSerialize, nil
}

func (o *TranqcontractGetCommunicationListV1Response) UnmarshalJSON(data []byte) (err error) {
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

	varTranqcontractGetCommunicationListV1Response := _TranqcontractGetCommunicationListV1Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTranqcontractGetCommunicationListV1Response)

	if err != nil {
		return err
	}

	*o = TranqcontractGetCommunicationListV1Response(varTranqcontractGetCommunicationListV1Response)

	return err
}

type NullableTranqcontractGetCommunicationListV1Response struct {
	value *TranqcontractGetCommunicationListV1Response
	isSet bool
}

func (v NullableTranqcontractGetCommunicationListV1Response) Get() *TranqcontractGetCommunicationListV1Response {
	return v.value
}

func (v *NullableTranqcontractGetCommunicationListV1Response) Set(val *TranqcontractGetCommunicationListV1Response) {
	v.value = val
	v.isSet = true
}

func (v NullableTranqcontractGetCommunicationListV1Response) IsSet() bool {
	return v.isSet
}

func (v *NullableTranqcontractGetCommunicationListV1Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTranqcontractGetCommunicationListV1Response(val *TranqcontractGetCommunicationListV1Response) *NullableTranqcontractGetCommunicationListV1Response {
	return &NullableTranqcontractGetCommunicationListV1Response{value: val, isSet: true}
}

func (v NullableTranqcontractGetCommunicationListV1Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTranqcontractGetCommunicationListV1Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


