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

// checks if the CorsEditObjectV1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CorsEditObjectV1Request{}

// CorsEditObjectV1Request Request for PUT /1/object/cors/{pkiCorsID}
type CorsEditObjectV1Request struct {
	ObjCors CorsRequestCompound `json:"objCors"`
}

type _CorsEditObjectV1Request CorsEditObjectV1Request

// NewCorsEditObjectV1Request instantiates a new CorsEditObjectV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCorsEditObjectV1Request(objCors CorsRequestCompound) *CorsEditObjectV1Request {
	this := CorsEditObjectV1Request{}
	this.ObjCors = objCors
	return &this
}

// NewCorsEditObjectV1RequestWithDefaults instantiates a new CorsEditObjectV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCorsEditObjectV1RequestWithDefaults() *CorsEditObjectV1Request {
	this := CorsEditObjectV1Request{}
	return &this
}

// GetObjCors returns the ObjCors field value
func (o *CorsEditObjectV1Request) GetObjCors() CorsRequestCompound {
	if o == nil {
		var ret CorsRequestCompound
		return ret
	}

	return o.ObjCors
}

// GetObjCorsOk returns a tuple with the ObjCors field value
// and a boolean to check if the value has been set.
func (o *CorsEditObjectV1Request) GetObjCorsOk() (*CorsRequestCompound, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjCors, true
}

// SetObjCors sets field value
func (o *CorsEditObjectV1Request) SetObjCors(v CorsRequestCompound) {
	o.ObjCors = v
}

func (o CorsEditObjectV1Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CorsEditObjectV1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objCors"] = o.ObjCors
	return toSerialize, nil
}

func (o *CorsEditObjectV1Request) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"objCors",
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

	varCorsEditObjectV1Request := _CorsEditObjectV1Request{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCorsEditObjectV1Request)

	if err != nil {
		return err
	}

	*o = CorsEditObjectV1Request(varCorsEditObjectV1Request)

	return err
}

type NullableCorsEditObjectV1Request struct {
	value *CorsEditObjectV1Request
	isSet bool
}

func (v NullableCorsEditObjectV1Request) Get() *CorsEditObjectV1Request {
	return v.value
}

func (v *NullableCorsEditObjectV1Request) Set(val *CorsEditObjectV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableCorsEditObjectV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableCorsEditObjectV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCorsEditObjectV1Request(val *CorsEditObjectV1Request) *NullableCorsEditObjectV1Request {
	return &NullableCorsEditObjectV1Request{value: val, isSet: true}
}

func (v NullableCorsEditObjectV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCorsEditObjectV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


