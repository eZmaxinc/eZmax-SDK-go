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

// checks if the CorsGetObjectV2Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CorsGetObjectV2Response{}

// CorsGetObjectV2Response Response for GET /2/object/cors/{pkiCorsID}
type CorsGetObjectV2Response struct {
	CommonResponse
	MPayload CorsGetObjectV2ResponseMPayload `json:"mPayload"`
}

type _CorsGetObjectV2Response CorsGetObjectV2Response

// NewCorsGetObjectV2Response instantiates a new CorsGetObjectV2Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCorsGetObjectV2Response(mPayload CorsGetObjectV2ResponseMPayload, objDebugPayload CommonResponseObjDebugPayload) *CorsGetObjectV2Response {
	this := CorsGetObjectV2Response{}
	this.ObjDebugPayload = objDebugPayload
	this.MPayload = mPayload
	return &this
}

// NewCorsGetObjectV2ResponseWithDefaults instantiates a new CorsGetObjectV2Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCorsGetObjectV2ResponseWithDefaults() *CorsGetObjectV2Response {
	this := CorsGetObjectV2Response{}
	return &this
}

// GetMPayload returns the MPayload field value
func (o *CorsGetObjectV2Response) GetMPayload() CorsGetObjectV2ResponseMPayload {
	if o == nil {
		var ret CorsGetObjectV2ResponseMPayload
		return ret
	}

	return o.MPayload
}

// GetMPayloadOk returns a tuple with the MPayload field value
// and a boolean to check if the value has been set.
func (o *CorsGetObjectV2Response) GetMPayloadOk() (*CorsGetObjectV2ResponseMPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MPayload, true
}

// SetMPayload sets field value
func (o *CorsGetObjectV2Response) SetMPayload(v CorsGetObjectV2ResponseMPayload) {
	o.MPayload = v
}

func (o CorsGetObjectV2Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CorsGetObjectV2Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mPayload"] = o.MPayload
	return toSerialize, nil
}

func (o *CorsGetObjectV2Response) UnmarshalJSON(data []byte) (err error) {
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

	varCorsGetObjectV2Response := _CorsGetObjectV2Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCorsGetObjectV2Response)

	if err != nil {
		return err
	}

	*o = CorsGetObjectV2Response(varCorsGetObjectV2Response)

	return err
}

type NullableCorsGetObjectV2Response struct {
	value *CorsGetObjectV2Response
	isSet bool
}

func (v NullableCorsGetObjectV2Response) Get() *CorsGetObjectV2Response {
	return v.value
}

func (v *NullableCorsGetObjectV2Response) Set(val *CorsGetObjectV2Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCorsGetObjectV2Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCorsGetObjectV2Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCorsGetObjectV2Response(val *CorsGetObjectV2Response) *NullableCorsGetObjectV2Response {
	return &NullableCorsGetObjectV2Response{value: val, isSet: true}
}

func (v NullableCorsGetObjectV2Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCorsGetObjectV2Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


