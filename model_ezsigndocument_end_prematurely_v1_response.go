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

// checks if the EzsigndocumentEndPrematurelyV1Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigndocumentEndPrematurelyV1Response{}

// EzsigndocumentEndPrematurelyV1Response Response for POST /1/object/ezsigndocument/{pkiEzsigndocument}/endPrematurely
type EzsigndocumentEndPrematurelyV1Response struct {
	ObjDebugPayload CommonResponseObjDebugPayload `json:"objDebugPayload"`
	ObjDebug *CommonResponseObjDebug `json:"objDebug,omitempty"`
}

type _EzsigndocumentEndPrematurelyV1Response EzsigndocumentEndPrematurelyV1Response

// NewEzsigndocumentEndPrematurelyV1Response instantiates a new EzsigndocumentEndPrematurelyV1Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigndocumentEndPrematurelyV1Response(objDebugPayload CommonResponseObjDebugPayload) *EzsigndocumentEndPrematurelyV1Response {
	this := EzsigndocumentEndPrematurelyV1Response{}
	this.ObjDebugPayload = objDebugPayload
	return &this
}

// NewEzsigndocumentEndPrematurelyV1ResponseWithDefaults instantiates a new EzsigndocumentEndPrematurelyV1Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigndocumentEndPrematurelyV1ResponseWithDefaults() *EzsigndocumentEndPrematurelyV1Response {
	this := EzsigndocumentEndPrematurelyV1Response{}
	return &this
}

// GetObjDebugPayload returns the ObjDebugPayload field value
func (o *EzsigndocumentEndPrematurelyV1Response) GetObjDebugPayload() CommonResponseObjDebugPayload {
	if o == nil {
		var ret CommonResponseObjDebugPayload
		return ret
	}

	return o.ObjDebugPayload
}

// GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentEndPrematurelyV1Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjDebugPayload, true
}

// SetObjDebugPayload sets field value
func (o *EzsigndocumentEndPrematurelyV1Response) SetObjDebugPayload(v CommonResponseObjDebugPayload) {
	o.ObjDebugPayload = v
}

// GetObjDebug returns the ObjDebug field value if set, zero value otherwise.
func (o *EzsigndocumentEndPrematurelyV1Response) GetObjDebug() CommonResponseObjDebug {
	if o == nil || IsNil(o.ObjDebug) {
		var ret CommonResponseObjDebug
		return ret
	}
	return *o.ObjDebug
}

// GetObjDebugOk returns a tuple with the ObjDebug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigndocumentEndPrematurelyV1Response) GetObjDebugOk() (*CommonResponseObjDebug, bool) {
	if o == nil || IsNil(o.ObjDebug) {
		return nil, false
	}
	return o.ObjDebug, true
}

// HasObjDebug returns a boolean if a field has been set.
func (o *EzsigndocumentEndPrematurelyV1Response) HasObjDebug() bool {
	if o != nil && !IsNil(o.ObjDebug) {
		return true
	}

	return false
}

// SetObjDebug gets a reference to the given CommonResponseObjDebug and assigns it to the ObjDebug field.
func (o *EzsigndocumentEndPrematurelyV1Response) SetObjDebug(v CommonResponseObjDebug) {
	o.ObjDebug = &v
}

func (o EzsigndocumentEndPrematurelyV1Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigndocumentEndPrematurelyV1Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objDebugPayload"] = o.ObjDebugPayload
	if !IsNil(o.ObjDebug) {
		toSerialize["objDebug"] = o.ObjDebug
	}
	return toSerialize, nil
}

func (o *EzsigndocumentEndPrematurelyV1Response) UnmarshalJSON(data []byte) (err error) {
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

	varEzsigndocumentEndPrematurelyV1Response := _EzsigndocumentEndPrematurelyV1Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsigndocumentEndPrematurelyV1Response)

	if err != nil {
		return err
	}

	*o = EzsigndocumentEndPrematurelyV1Response(varEzsigndocumentEndPrematurelyV1Response)

	return err
}

type NullableEzsigndocumentEndPrematurelyV1Response struct {
	value *EzsigndocumentEndPrematurelyV1Response
	isSet bool
}

func (v NullableEzsigndocumentEndPrematurelyV1Response) Get() *EzsigndocumentEndPrematurelyV1Response {
	return v.value
}

func (v *NullableEzsigndocumentEndPrematurelyV1Response) Set(val *EzsigndocumentEndPrematurelyV1Response) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigndocumentEndPrematurelyV1Response) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigndocumentEndPrematurelyV1Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigndocumentEndPrematurelyV1Response(val *EzsigndocumentEndPrematurelyV1Response) *NullableEzsigndocumentEndPrematurelyV1Response {
	return &NullableEzsigndocumentEndPrematurelyV1Response{value: val, isSet: true}
}

func (v NullableEzsigndocumentEndPrematurelyV1Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigndocumentEndPrematurelyV1Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


