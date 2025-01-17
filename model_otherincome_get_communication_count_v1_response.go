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

// checks if the OtherincomeGetCommunicationCountV1Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OtherincomeGetCommunicationCountV1Response{}

// OtherincomeGetCommunicationCountV1Response Response for GET /1/object/otherincome/{pkiOtherincomeID}/getCommunicationCount
type OtherincomeGetCommunicationCountV1Response struct {
	CommonResponse
	MPayload OtherincomeGetCommunicationCountV1ResponseMPayload `json:"mPayload"`
}

type _OtherincomeGetCommunicationCountV1Response OtherincomeGetCommunicationCountV1Response

// NewOtherincomeGetCommunicationCountV1Response instantiates a new OtherincomeGetCommunicationCountV1Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOtherincomeGetCommunicationCountV1Response(mPayload OtherincomeGetCommunicationCountV1ResponseMPayload, objDebugPayload CommonResponseObjDebugPayload) *OtherincomeGetCommunicationCountV1Response {
	this := OtherincomeGetCommunicationCountV1Response{}
	this.ObjDebugPayload = objDebugPayload
	this.MPayload = mPayload
	return &this
}

// NewOtherincomeGetCommunicationCountV1ResponseWithDefaults instantiates a new OtherincomeGetCommunicationCountV1Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOtherincomeGetCommunicationCountV1ResponseWithDefaults() *OtherincomeGetCommunicationCountV1Response {
	this := OtherincomeGetCommunicationCountV1Response{}
	return &this
}

// GetMPayload returns the MPayload field value
func (o *OtherincomeGetCommunicationCountV1Response) GetMPayload() OtherincomeGetCommunicationCountV1ResponseMPayload {
	if o == nil {
		var ret OtherincomeGetCommunicationCountV1ResponseMPayload
		return ret
	}

	return o.MPayload
}

// GetMPayloadOk returns a tuple with the MPayload field value
// and a boolean to check if the value has been set.
func (o *OtherincomeGetCommunicationCountV1Response) GetMPayloadOk() (*OtherincomeGetCommunicationCountV1ResponseMPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MPayload, true
}

// SetMPayload sets field value
func (o *OtherincomeGetCommunicationCountV1Response) SetMPayload(v OtherincomeGetCommunicationCountV1ResponseMPayload) {
	o.MPayload = v
}

func (o OtherincomeGetCommunicationCountV1Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OtherincomeGetCommunicationCountV1Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mPayload"] = o.MPayload
	return toSerialize, nil
}

func (o *OtherincomeGetCommunicationCountV1Response) UnmarshalJSON(data []byte) (err error) {
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

	varOtherincomeGetCommunicationCountV1Response := _OtherincomeGetCommunicationCountV1Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOtherincomeGetCommunicationCountV1Response)

	if err != nil {
		return err
	}

	*o = OtherincomeGetCommunicationCountV1Response(varOtherincomeGetCommunicationCountV1Response)

	return err
}

type NullableOtherincomeGetCommunicationCountV1Response struct {
	value *OtherincomeGetCommunicationCountV1Response
	isSet bool
}

func (v NullableOtherincomeGetCommunicationCountV1Response) Get() *OtherincomeGetCommunicationCountV1Response {
	return v.value
}

func (v *NullableOtherincomeGetCommunicationCountV1Response) Set(val *OtherincomeGetCommunicationCountV1Response) {
	v.value = val
	v.isSet = true
}

func (v NullableOtherincomeGetCommunicationCountV1Response) IsSet() bool {
	return v.isSet
}

func (v *NullableOtherincomeGetCommunicationCountV1Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOtherincomeGetCommunicationCountV1Response(val *OtherincomeGetCommunicationCountV1Response) *NullableOtherincomeGetCommunicationCountV1Response {
	return &NullableOtherincomeGetCommunicationCountV1Response{value: val, isSet: true}
}

func (v NullableOtherincomeGetCommunicationCountV1Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOtherincomeGetCommunicationCountV1Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


