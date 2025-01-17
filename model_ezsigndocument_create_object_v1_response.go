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

// checks if the EzsigndocumentCreateObjectV1Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigndocumentCreateObjectV1Response{}

// EzsigndocumentCreateObjectV1Response Response for POST /1/object/ezsigndocument
type EzsigndocumentCreateObjectV1Response struct {
	ObjDebugPayload CommonResponseObjDebugPayload `json:"objDebugPayload"`
	ObjDebug *CommonResponseObjDebug `json:"objDebug,omitempty"`
	MPayload EzsigndocumentCreateObjectV1ResponseMPayload `json:"mPayload"`
}

type _EzsigndocumentCreateObjectV1Response EzsigndocumentCreateObjectV1Response

// NewEzsigndocumentCreateObjectV1Response instantiates a new EzsigndocumentCreateObjectV1Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigndocumentCreateObjectV1Response(objDebugPayload CommonResponseObjDebugPayload, mPayload EzsigndocumentCreateObjectV1ResponseMPayload) *EzsigndocumentCreateObjectV1Response {
	this := EzsigndocumentCreateObjectV1Response{}
	this.ObjDebugPayload = objDebugPayload
	this.MPayload = mPayload
	return &this
}

// NewEzsigndocumentCreateObjectV1ResponseWithDefaults instantiates a new EzsigndocumentCreateObjectV1Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigndocumentCreateObjectV1ResponseWithDefaults() *EzsigndocumentCreateObjectV1Response {
	this := EzsigndocumentCreateObjectV1Response{}
	return &this
}

// GetObjDebugPayload returns the ObjDebugPayload field value
func (o *EzsigndocumentCreateObjectV1Response) GetObjDebugPayload() CommonResponseObjDebugPayload {
	if o == nil {
		var ret CommonResponseObjDebugPayload
		return ret
	}

	return o.ObjDebugPayload
}

// GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentCreateObjectV1Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjDebugPayload, true
}

// SetObjDebugPayload sets field value
func (o *EzsigndocumentCreateObjectV1Response) SetObjDebugPayload(v CommonResponseObjDebugPayload) {
	o.ObjDebugPayload = v
}

// GetObjDebug returns the ObjDebug field value if set, zero value otherwise.
func (o *EzsigndocumentCreateObjectV1Response) GetObjDebug() CommonResponseObjDebug {
	if o == nil || IsNil(o.ObjDebug) {
		var ret CommonResponseObjDebug
		return ret
	}
	return *o.ObjDebug
}

// GetObjDebugOk returns a tuple with the ObjDebug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigndocumentCreateObjectV1Response) GetObjDebugOk() (*CommonResponseObjDebug, bool) {
	if o == nil || IsNil(o.ObjDebug) {
		return nil, false
	}
	return o.ObjDebug, true
}

// HasObjDebug returns a boolean if a field has been set.
func (o *EzsigndocumentCreateObjectV1Response) HasObjDebug() bool {
	if o != nil && !IsNil(o.ObjDebug) {
		return true
	}

	return false
}

// SetObjDebug gets a reference to the given CommonResponseObjDebug and assigns it to the ObjDebug field.
func (o *EzsigndocumentCreateObjectV1Response) SetObjDebug(v CommonResponseObjDebug) {
	o.ObjDebug = &v
}

// GetMPayload returns the MPayload field value
func (o *EzsigndocumentCreateObjectV1Response) GetMPayload() EzsigndocumentCreateObjectV1ResponseMPayload {
	if o == nil {
		var ret EzsigndocumentCreateObjectV1ResponseMPayload
		return ret
	}

	return o.MPayload
}

// GetMPayloadOk returns a tuple with the MPayload field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentCreateObjectV1Response) GetMPayloadOk() (*EzsigndocumentCreateObjectV1ResponseMPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MPayload, true
}

// SetMPayload sets field value
func (o *EzsigndocumentCreateObjectV1Response) SetMPayload(v EzsigndocumentCreateObjectV1ResponseMPayload) {
	o.MPayload = v
}

func (o EzsigndocumentCreateObjectV1Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigndocumentCreateObjectV1Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objDebugPayload"] = o.ObjDebugPayload
	if !IsNil(o.ObjDebug) {
		toSerialize["objDebug"] = o.ObjDebug
	}
	toSerialize["mPayload"] = o.MPayload
	return toSerialize, nil
}

func (o *EzsigndocumentCreateObjectV1Response) UnmarshalJSON(data []byte) (err error) {
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

	varEzsigndocumentCreateObjectV1Response := _EzsigndocumentCreateObjectV1Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsigndocumentCreateObjectV1Response)

	if err != nil {
		return err
	}

	*o = EzsigndocumentCreateObjectV1Response(varEzsigndocumentCreateObjectV1Response)

	return err
}

type NullableEzsigndocumentCreateObjectV1Response struct {
	value *EzsigndocumentCreateObjectV1Response
	isSet bool
}

func (v NullableEzsigndocumentCreateObjectV1Response) Get() *EzsigndocumentCreateObjectV1Response {
	return v.value
}

func (v *NullableEzsigndocumentCreateObjectV1Response) Set(val *EzsigndocumentCreateObjectV1Response) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigndocumentCreateObjectV1Response) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigndocumentCreateObjectV1Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigndocumentCreateObjectV1Response(val *EzsigndocumentCreateObjectV1Response) *NullableEzsigndocumentCreateObjectV1Response {
	return &NullableEzsigndocumentCreateObjectV1Response{value: val, isSet: true}
}

func (v NullableEzsigndocumentCreateObjectV1Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigndocumentCreateObjectV1Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


