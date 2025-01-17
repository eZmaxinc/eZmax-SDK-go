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

// checks if the RejectedoffertopurchaseGetCommunicationsendersV1Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RejectedoffertopurchaseGetCommunicationsendersV1Response{}

// RejectedoffertopurchaseGetCommunicationsendersV1Response Response for GET /1/object/rejectedoffertopurchase/{pkiRejectedoffertopurchaseID}/getCommunicationrecipients
type RejectedoffertopurchaseGetCommunicationsendersV1Response struct {
	CommonResponse
	MPayload RejectedoffertopurchaseGetCommunicationsendersV1ResponseMPayload `json:"mPayload"`
}

type _RejectedoffertopurchaseGetCommunicationsendersV1Response RejectedoffertopurchaseGetCommunicationsendersV1Response

// NewRejectedoffertopurchaseGetCommunicationsendersV1Response instantiates a new RejectedoffertopurchaseGetCommunicationsendersV1Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRejectedoffertopurchaseGetCommunicationsendersV1Response(mPayload RejectedoffertopurchaseGetCommunicationsendersV1ResponseMPayload, objDebugPayload CommonResponseObjDebugPayload) *RejectedoffertopurchaseGetCommunicationsendersV1Response {
	this := RejectedoffertopurchaseGetCommunicationsendersV1Response{}
	this.ObjDebugPayload = objDebugPayload
	this.MPayload = mPayload
	return &this
}

// NewRejectedoffertopurchaseGetCommunicationsendersV1ResponseWithDefaults instantiates a new RejectedoffertopurchaseGetCommunicationsendersV1Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRejectedoffertopurchaseGetCommunicationsendersV1ResponseWithDefaults() *RejectedoffertopurchaseGetCommunicationsendersV1Response {
	this := RejectedoffertopurchaseGetCommunicationsendersV1Response{}
	return &this
}

// GetMPayload returns the MPayload field value
func (o *RejectedoffertopurchaseGetCommunicationsendersV1Response) GetMPayload() RejectedoffertopurchaseGetCommunicationsendersV1ResponseMPayload {
	if o == nil {
		var ret RejectedoffertopurchaseGetCommunicationsendersV1ResponseMPayload
		return ret
	}

	return o.MPayload
}

// GetMPayloadOk returns a tuple with the MPayload field value
// and a boolean to check if the value has been set.
func (o *RejectedoffertopurchaseGetCommunicationsendersV1Response) GetMPayloadOk() (*RejectedoffertopurchaseGetCommunicationsendersV1ResponseMPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MPayload, true
}

// SetMPayload sets field value
func (o *RejectedoffertopurchaseGetCommunicationsendersV1Response) SetMPayload(v RejectedoffertopurchaseGetCommunicationsendersV1ResponseMPayload) {
	o.MPayload = v
}

func (o RejectedoffertopurchaseGetCommunicationsendersV1Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RejectedoffertopurchaseGetCommunicationsendersV1Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mPayload"] = o.MPayload
	return toSerialize, nil
}

func (o *RejectedoffertopurchaseGetCommunicationsendersV1Response) UnmarshalJSON(data []byte) (err error) {
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

	varRejectedoffertopurchaseGetCommunicationsendersV1Response := _RejectedoffertopurchaseGetCommunicationsendersV1Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRejectedoffertopurchaseGetCommunicationsendersV1Response)

	if err != nil {
		return err
	}

	*o = RejectedoffertopurchaseGetCommunicationsendersV1Response(varRejectedoffertopurchaseGetCommunicationsendersV1Response)

	return err
}

type NullableRejectedoffertopurchaseGetCommunicationsendersV1Response struct {
	value *RejectedoffertopurchaseGetCommunicationsendersV1Response
	isSet bool
}

func (v NullableRejectedoffertopurchaseGetCommunicationsendersV1Response) Get() *RejectedoffertopurchaseGetCommunicationsendersV1Response {
	return v.value
}

func (v *NullableRejectedoffertopurchaseGetCommunicationsendersV1Response) Set(val *RejectedoffertopurchaseGetCommunicationsendersV1Response) {
	v.value = val
	v.isSet = true
}

func (v NullableRejectedoffertopurchaseGetCommunicationsendersV1Response) IsSet() bool {
	return v.isSet
}

func (v *NullableRejectedoffertopurchaseGetCommunicationsendersV1Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRejectedoffertopurchaseGetCommunicationsendersV1Response(val *RejectedoffertopurchaseGetCommunicationsendersV1Response) *NullableRejectedoffertopurchaseGetCommunicationsendersV1Response {
	return &NullableRejectedoffertopurchaseGetCommunicationsendersV1Response{value: val, isSet: true}
}

func (v NullableRejectedoffertopurchaseGetCommunicationsendersV1Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRejectedoffertopurchaseGetCommunicationsendersV1Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


