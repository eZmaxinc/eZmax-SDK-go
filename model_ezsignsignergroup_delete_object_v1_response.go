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

// checks if the EzsignsignergroupDeleteObjectV1Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignsignergroupDeleteObjectV1Response{}

// EzsignsignergroupDeleteObjectV1Response Response for DELETE /1/object/ezsignsignergroup/{pkiEzsignsignergroupID}
type EzsignsignergroupDeleteObjectV1Response struct {
	ObjDebugPayload CommonResponseObjDebugPayload `json:"objDebugPayload"`
	ObjDebug *CommonResponseObjDebug `json:"objDebug,omitempty"`
}

type _EzsignsignergroupDeleteObjectV1Response EzsignsignergroupDeleteObjectV1Response

// NewEzsignsignergroupDeleteObjectV1Response instantiates a new EzsignsignergroupDeleteObjectV1Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignsignergroupDeleteObjectV1Response(objDebugPayload CommonResponseObjDebugPayload) *EzsignsignergroupDeleteObjectV1Response {
	this := EzsignsignergroupDeleteObjectV1Response{}
	this.ObjDebugPayload = objDebugPayload
	return &this
}

// NewEzsignsignergroupDeleteObjectV1ResponseWithDefaults instantiates a new EzsignsignergroupDeleteObjectV1Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignsignergroupDeleteObjectV1ResponseWithDefaults() *EzsignsignergroupDeleteObjectV1Response {
	this := EzsignsignergroupDeleteObjectV1Response{}
	return &this
}

// GetObjDebugPayload returns the ObjDebugPayload field value
func (o *EzsignsignergroupDeleteObjectV1Response) GetObjDebugPayload() CommonResponseObjDebugPayload {
	if o == nil {
		var ret CommonResponseObjDebugPayload
		return ret
	}

	return o.ObjDebugPayload
}

// GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field value
// and a boolean to check if the value has been set.
func (o *EzsignsignergroupDeleteObjectV1Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjDebugPayload, true
}

// SetObjDebugPayload sets field value
func (o *EzsignsignergroupDeleteObjectV1Response) SetObjDebugPayload(v CommonResponseObjDebugPayload) {
	o.ObjDebugPayload = v
}

// GetObjDebug returns the ObjDebug field value if set, zero value otherwise.
func (o *EzsignsignergroupDeleteObjectV1Response) GetObjDebug() CommonResponseObjDebug {
	if o == nil || IsNil(o.ObjDebug) {
		var ret CommonResponseObjDebug
		return ret
	}
	return *o.ObjDebug
}

// GetObjDebugOk returns a tuple with the ObjDebug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignsignergroupDeleteObjectV1Response) GetObjDebugOk() (*CommonResponseObjDebug, bool) {
	if o == nil || IsNil(o.ObjDebug) {
		return nil, false
	}
	return o.ObjDebug, true
}

// HasObjDebug returns a boolean if a field has been set.
func (o *EzsignsignergroupDeleteObjectV1Response) HasObjDebug() bool {
	if o != nil && !IsNil(o.ObjDebug) {
		return true
	}

	return false
}

// SetObjDebug gets a reference to the given CommonResponseObjDebug and assigns it to the ObjDebug field.
func (o *EzsignsignergroupDeleteObjectV1Response) SetObjDebug(v CommonResponseObjDebug) {
	o.ObjDebug = &v
}

func (o EzsignsignergroupDeleteObjectV1Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignsignergroupDeleteObjectV1Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objDebugPayload"] = o.ObjDebugPayload
	if !IsNil(o.ObjDebug) {
		toSerialize["objDebug"] = o.ObjDebug
	}
	return toSerialize, nil
}

func (o *EzsignsignergroupDeleteObjectV1Response) UnmarshalJSON(data []byte) (err error) {
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

	varEzsignsignergroupDeleteObjectV1Response := _EzsignsignergroupDeleteObjectV1Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsignsignergroupDeleteObjectV1Response)

	if err != nil {
		return err
	}

	*o = EzsignsignergroupDeleteObjectV1Response(varEzsignsignergroupDeleteObjectV1Response)

	return err
}

type NullableEzsignsignergroupDeleteObjectV1Response struct {
	value *EzsignsignergroupDeleteObjectV1Response
	isSet bool
}

func (v NullableEzsignsignergroupDeleteObjectV1Response) Get() *EzsignsignergroupDeleteObjectV1Response {
	return v.value
}

func (v *NullableEzsignsignergroupDeleteObjectV1Response) Set(val *EzsignsignergroupDeleteObjectV1Response) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignsignergroupDeleteObjectV1Response) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignsignergroupDeleteObjectV1Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignsignergroupDeleteObjectV1Response(val *EzsignsignergroupDeleteObjectV1Response) *NullableEzsignsignergroupDeleteObjectV1Response {
	return &NullableEzsignsignergroupDeleteObjectV1Response{value: val, isSet: true}
}

func (v NullableEzsignsignergroupDeleteObjectV1Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignsignergroupDeleteObjectV1Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


