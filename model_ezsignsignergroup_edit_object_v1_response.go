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

// checks if the EzsignsignergroupEditObjectV1Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignsignergroupEditObjectV1Response{}

// EzsignsignergroupEditObjectV1Response Response for PUT /1/object/ezsignsignergroup/{pkiEzsignsignergroupID}
type EzsignsignergroupEditObjectV1Response struct {
	ObjDebugPayload CommonResponseObjDebugPayload `json:"objDebugPayload"`
	ObjDebug *CommonResponseObjDebug `json:"objDebug,omitempty"`
}

type _EzsignsignergroupEditObjectV1Response EzsignsignergroupEditObjectV1Response

// NewEzsignsignergroupEditObjectV1Response instantiates a new EzsignsignergroupEditObjectV1Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignsignergroupEditObjectV1Response(objDebugPayload CommonResponseObjDebugPayload) *EzsignsignergroupEditObjectV1Response {
	this := EzsignsignergroupEditObjectV1Response{}
	this.ObjDebugPayload = objDebugPayload
	return &this
}

// NewEzsignsignergroupEditObjectV1ResponseWithDefaults instantiates a new EzsignsignergroupEditObjectV1Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignsignergroupEditObjectV1ResponseWithDefaults() *EzsignsignergroupEditObjectV1Response {
	this := EzsignsignergroupEditObjectV1Response{}
	return &this
}

// GetObjDebugPayload returns the ObjDebugPayload field value
func (o *EzsignsignergroupEditObjectV1Response) GetObjDebugPayload() CommonResponseObjDebugPayload {
	if o == nil {
		var ret CommonResponseObjDebugPayload
		return ret
	}

	return o.ObjDebugPayload
}

// GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field value
// and a boolean to check if the value has been set.
func (o *EzsignsignergroupEditObjectV1Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjDebugPayload, true
}

// SetObjDebugPayload sets field value
func (o *EzsignsignergroupEditObjectV1Response) SetObjDebugPayload(v CommonResponseObjDebugPayload) {
	o.ObjDebugPayload = v
}

// GetObjDebug returns the ObjDebug field value if set, zero value otherwise.
func (o *EzsignsignergroupEditObjectV1Response) GetObjDebug() CommonResponseObjDebug {
	if o == nil || IsNil(o.ObjDebug) {
		var ret CommonResponseObjDebug
		return ret
	}
	return *o.ObjDebug
}

// GetObjDebugOk returns a tuple with the ObjDebug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignsignergroupEditObjectV1Response) GetObjDebugOk() (*CommonResponseObjDebug, bool) {
	if o == nil || IsNil(o.ObjDebug) {
		return nil, false
	}
	return o.ObjDebug, true
}

// HasObjDebug returns a boolean if a field has been set.
func (o *EzsignsignergroupEditObjectV1Response) HasObjDebug() bool {
	if o != nil && !IsNil(o.ObjDebug) {
		return true
	}

	return false
}

// SetObjDebug gets a reference to the given CommonResponseObjDebug and assigns it to the ObjDebug field.
func (o *EzsignsignergroupEditObjectV1Response) SetObjDebug(v CommonResponseObjDebug) {
	o.ObjDebug = &v
}

func (o EzsignsignergroupEditObjectV1Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignsignergroupEditObjectV1Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objDebugPayload"] = o.ObjDebugPayload
	if !IsNil(o.ObjDebug) {
		toSerialize["objDebug"] = o.ObjDebug
	}
	return toSerialize, nil
}

func (o *EzsignsignergroupEditObjectV1Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varEzsignsignergroupEditObjectV1Response := _EzsignsignergroupEditObjectV1Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsignsignergroupEditObjectV1Response)

	if err != nil {
		return err
	}

	*o = EzsignsignergroupEditObjectV1Response(varEzsignsignergroupEditObjectV1Response)

	return err
}

type NullableEzsignsignergroupEditObjectV1Response struct {
	value *EzsignsignergroupEditObjectV1Response
	isSet bool
}

func (v NullableEzsignsignergroupEditObjectV1Response) Get() *EzsignsignergroupEditObjectV1Response {
	return v.value
}

func (v *NullableEzsignsignergroupEditObjectV1Response) Set(val *EzsignsignergroupEditObjectV1Response) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignsignergroupEditObjectV1Response) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignsignergroupEditObjectV1Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignsignergroupEditObjectV1Response(val *EzsignsignergroupEditObjectV1Response) *NullableEzsignsignergroupEditObjectV1Response {
	return &NullableEzsignsignergroupEditObjectV1Response{value: val, isSet: true}
}

func (v NullableEzsignsignergroupEditObjectV1Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignsignergroupEditObjectV1Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


