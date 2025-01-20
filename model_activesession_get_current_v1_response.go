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

// checks if the ActivesessionGetCurrentV1Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActivesessionGetCurrentV1Response{}

// ActivesessionGetCurrentV1Response Response for GET /1/object/activesession/getCurrent
type ActivesessionGetCurrentV1Response struct {
	ObjDebugPayload CommonResponseObjDebugPayload `json:"objDebugPayload"`
	ObjDebug *CommonResponseObjDebug `json:"objDebug,omitempty"`
	MPayload ActivesessionGetCurrentV1ResponseMPayload `json:"mPayload"`
}

type _ActivesessionGetCurrentV1Response ActivesessionGetCurrentV1Response

// NewActivesessionGetCurrentV1Response instantiates a new ActivesessionGetCurrentV1Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActivesessionGetCurrentV1Response(objDebugPayload CommonResponseObjDebugPayload, mPayload ActivesessionGetCurrentV1ResponseMPayload) *ActivesessionGetCurrentV1Response {
	this := ActivesessionGetCurrentV1Response{}
	this.ObjDebugPayload = objDebugPayload
	this.MPayload = mPayload
	return &this
}

// NewActivesessionGetCurrentV1ResponseWithDefaults instantiates a new ActivesessionGetCurrentV1Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActivesessionGetCurrentV1ResponseWithDefaults() *ActivesessionGetCurrentV1Response {
	this := ActivesessionGetCurrentV1Response{}
	return &this
}

// GetObjDebugPayload returns the ObjDebugPayload field value
func (o *ActivesessionGetCurrentV1Response) GetObjDebugPayload() CommonResponseObjDebugPayload {
	if o == nil {
		var ret CommonResponseObjDebugPayload
		return ret
	}

	return o.ObjDebugPayload
}

// GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field value
// and a boolean to check if the value has been set.
func (o *ActivesessionGetCurrentV1Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjDebugPayload, true
}

// SetObjDebugPayload sets field value
func (o *ActivesessionGetCurrentV1Response) SetObjDebugPayload(v CommonResponseObjDebugPayload) {
	o.ObjDebugPayload = v
}

// GetObjDebug returns the ObjDebug field value if set, zero value otherwise.
func (o *ActivesessionGetCurrentV1Response) GetObjDebug() CommonResponseObjDebug {
	if o == nil || IsNil(o.ObjDebug) {
		var ret CommonResponseObjDebug
		return ret
	}
	return *o.ObjDebug
}

// GetObjDebugOk returns a tuple with the ObjDebug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivesessionGetCurrentV1Response) GetObjDebugOk() (*CommonResponseObjDebug, bool) {
	if o == nil || IsNil(o.ObjDebug) {
		return nil, false
	}
	return o.ObjDebug, true
}

// HasObjDebug returns a boolean if a field has been set.
func (o *ActivesessionGetCurrentV1Response) HasObjDebug() bool {
	if o != nil && !IsNil(o.ObjDebug) {
		return true
	}

	return false
}

// SetObjDebug gets a reference to the given CommonResponseObjDebug and assigns it to the ObjDebug field.
func (o *ActivesessionGetCurrentV1Response) SetObjDebug(v CommonResponseObjDebug) {
	o.ObjDebug = &v
}

// GetMPayload returns the MPayload field value
func (o *ActivesessionGetCurrentV1Response) GetMPayload() ActivesessionGetCurrentV1ResponseMPayload {
	if o == nil {
		var ret ActivesessionGetCurrentV1ResponseMPayload
		return ret
	}

	return o.MPayload
}

// GetMPayloadOk returns a tuple with the MPayload field value
// and a boolean to check if the value has been set.
func (o *ActivesessionGetCurrentV1Response) GetMPayloadOk() (*ActivesessionGetCurrentV1ResponseMPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MPayload, true
}

// SetMPayload sets field value
func (o *ActivesessionGetCurrentV1Response) SetMPayload(v ActivesessionGetCurrentV1ResponseMPayload) {
	o.MPayload = v
}

func (o ActivesessionGetCurrentV1Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActivesessionGetCurrentV1Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objDebugPayload"] = o.ObjDebugPayload
	if !IsNil(o.ObjDebug) {
		toSerialize["objDebug"] = o.ObjDebug
	}
	toSerialize["mPayload"] = o.MPayload
	return toSerialize, nil
}

func (o *ActivesessionGetCurrentV1Response) UnmarshalJSON(data []byte) (err error) {
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

	varActivesessionGetCurrentV1Response := _ActivesessionGetCurrentV1Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varActivesessionGetCurrentV1Response)

	if err != nil {
		return err
	}

	*o = ActivesessionGetCurrentV1Response(varActivesessionGetCurrentV1Response)

	return err
}

type NullableActivesessionGetCurrentV1Response struct {
	value *ActivesessionGetCurrentV1Response
	isSet bool
}

func (v NullableActivesessionGetCurrentV1Response) Get() *ActivesessionGetCurrentV1Response {
	return v.value
}

func (v *NullableActivesessionGetCurrentV1Response) Set(val *ActivesessionGetCurrentV1Response) {
	v.value = val
	v.isSet = true
}

func (v NullableActivesessionGetCurrentV1Response) IsSet() bool {
	return v.isSet
}

func (v *NullableActivesessionGetCurrentV1Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActivesessionGetCurrentV1Response(val *ActivesessionGetCurrentV1Response) *NullableActivesessionGetCurrentV1Response {
	return &NullableActivesessionGetCurrentV1Response{value: val, isSet: true}
}

func (v NullableActivesessionGetCurrentV1Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActivesessionGetCurrentV1Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


