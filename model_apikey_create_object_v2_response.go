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

// checks if the ApikeyCreateObjectV2Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApikeyCreateObjectV2Response{}

// ApikeyCreateObjectV2Response Response for POST /2/object/apikey
type ApikeyCreateObjectV2Response struct {
	ObjDebugPayload CommonResponseObjDebugPayload `json:"objDebugPayload"`
	ObjDebug *CommonResponseObjDebug `json:"objDebug,omitempty"`
	MPayload ApikeyCreateObjectV2ResponseMPayload `json:"mPayload"`
}

type _ApikeyCreateObjectV2Response ApikeyCreateObjectV2Response

// NewApikeyCreateObjectV2Response instantiates a new ApikeyCreateObjectV2Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApikeyCreateObjectV2Response(objDebugPayload CommonResponseObjDebugPayload, mPayload ApikeyCreateObjectV2ResponseMPayload) *ApikeyCreateObjectV2Response {
	this := ApikeyCreateObjectV2Response{}
	this.ObjDebugPayload = objDebugPayload
	this.MPayload = mPayload
	return &this
}

// NewApikeyCreateObjectV2ResponseWithDefaults instantiates a new ApikeyCreateObjectV2Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApikeyCreateObjectV2ResponseWithDefaults() *ApikeyCreateObjectV2Response {
	this := ApikeyCreateObjectV2Response{}
	return &this
}

// GetObjDebugPayload returns the ObjDebugPayload field value
func (o *ApikeyCreateObjectV2Response) GetObjDebugPayload() CommonResponseObjDebugPayload {
	if o == nil {
		var ret CommonResponseObjDebugPayload
		return ret
	}

	return o.ObjDebugPayload
}

// GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field value
// and a boolean to check if the value has been set.
func (o *ApikeyCreateObjectV2Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjDebugPayload, true
}

// SetObjDebugPayload sets field value
func (o *ApikeyCreateObjectV2Response) SetObjDebugPayload(v CommonResponseObjDebugPayload) {
	o.ObjDebugPayload = v
}

// GetObjDebug returns the ObjDebug field value if set, zero value otherwise.
func (o *ApikeyCreateObjectV2Response) GetObjDebug() CommonResponseObjDebug {
	if o == nil || IsNil(o.ObjDebug) {
		var ret CommonResponseObjDebug
		return ret
	}
	return *o.ObjDebug
}

// GetObjDebugOk returns a tuple with the ObjDebug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApikeyCreateObjectV2Response) GetObjDebugOk() (*CommonResponseObjDebug, bool) {
	if o == nil || IsNil(o.ObjDebug) {
		return nil, false
	}
	return o.ObjDebug, true
}

// HasObjDebug returns a boolean if a field has been set.
func (o *ApikeyCreateObjectV2Response) HasObjDebug() bool {
	if o != nil && !IsNil(o.ObjDebug) {
		return true
	}

	return false
}

// SetObjDebug gets a reference to the given CommonResponseObjDebug and assigns it to the ObjDebug field.
func (o *ApikeyCreateObjectV2Response) SetObjDebug(v CommonResponseObjDebug) {
	o.ObjDebug = &v
}

// GetMPayload returns the MPayload field value
func (o *ApikeyCreateObjectV2Response) GetMPayload() ApikeyCreateObjectV2ResponseMPayload {
	if o == nil {
		var ret ApikeyCreateObjectV2ResponseMPayload
		return ret
	}

	return o.MPayload
}

// GetMPayloadOk returns a tuple with the MPayload field value
// and a boolean to check if the value has been set.
func (o *ApikeyCreateObjectV2Response) GetMPayloadOk() (*ApikeyCreateObjectV2ResponseMPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MPayload, true
}

// SetMPayload sets field value
func (o *ApikeyCreateObjectV2Response) SetMPayload(v ApikeyCreateObjectV2ResponseMPayload) {
	o.MPayload = v
}

func (o ApikeyCreateObjectV2Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApikeyCreateObjectV2Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objDebugPayload"] = o.ObjDebugPayload
	if !IsNil(o.ObjDebug) {
		toSerialize["objDebug"] = o.ObjDebug
	}
	toSerialize["mPayload"] = o.MPayload
	return toSerialize, nil
}

func (o *ApikeyCreateObjectV2Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"objDebugPayload",
		"mPayload",
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

	varApikeyCreateObjectV2Response := _ApikeyCreateObjectV2Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApikeyCreateObjectV2Response)

	if err != nil {
		return err
	}

	*o = ApikeyCreateObjectV2Response(varApikeyCreateObjectV2Response)

	return err
}

type NullableApikeyCreateObjectV2Response struct {
	value *ApikeyCreateObjectV2Response
	isSet bool
}

func (v NullableApikeyCreateObjectV2Response) Get() *ApikeyCreateObjectV2Response {
	return v.value
}

func (v *NullableApikeyCreateObjectV2Response) Set(val *ApikeyCreateObjectV2Response) {
	v.value = val
	v.isSet = true
}

func (v NullableApikeyCreateObjectV2Response) IsSet() bool {
	return v.isSet
}

func (v *NullableApikeyCreateObjectV2Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApikeyCreateObjectV2Response(val *ApikeyCreateObjectV2Response) *NullableApikeyCreateObjectV2Response {
	return &NullableApikeyCreateObjectV2Response{value: val, isSet: true}
}

func (v NullableApikeyCreateObjectV2Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApikeyCreateObjectV2Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


