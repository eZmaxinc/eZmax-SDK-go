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

// checks if the EzsigntemplateCreateObjectV3Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigntemplateCreateObjectV3Response{}

// EzsigntemplateCreateObjectV3Response Response for POST /3/object/ezsigntemplate
type EzsigntemplateCreateObjectV3Response struct {
	ObjDebugPayload CommonResponseObjDebugPayload `json:"objDebugPayload"`
	ObjDebug *CommonResponseObjDebug `json:"objDebug,omitempty"`
	MPayload EzsigntemplateCreateObjectV3ResponseMPayload `json:"mPayload"`
}

type _EzsigntemplateCreateObjectV3Response EzsigntemplateCreateObjectV3Response

// NewEzsigntemplateCreateObjectV3Response instantiates a new EzsigntemplateCreateObjectV3Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigntemplateCreateObjectV3Response(objDebugPayload CommonResponseObjDebugPayload, mPayload EzsigntemplateCreateObjectV3ResponseMPayload) *EzsigntemplateCreateObjectV3Response {
	this := EzsigntemplateCreateObjectV3Response{}
	this.ObjDebugPayload = objDebugPayload
	this.MPayload = mPayload
	return &this
}

// NewEzsigntemplateCreateObjectV3ResponseWithDefaults instantiates a new EzsigntemplateCreateObjectV3Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigntemplateCreateObjectV3ResponseWithDefaults() *EzsigntemplateCreateObjectV3Response {
	this := EzsigntemplateCreateObjectV3Response{}
	return &this
}

// GetObjDebugPayload returns the ObjDebugPayload field value
func (o *EzsigntemplateCreateObjectV3Response) GetObjDebugPayload() CommonResponseObjDebugPayload {
	if o == nil {
		var ret CommonResponseObjDebugPayload
		return ret
	}

	return o.ObjDebugPayload
}

// GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateCreateObjectV3Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjDebugPayload, true
}

// SetObjDebugPayload sets field value
func (o *EzsigntemplateCreateObjectV3Response) SetObjDebugPayload(v CommonResponseObjDebugPayload) {
	o.ObjDebugPayload = v
}

// GetObjDebug returns the ObjDebug field value if set, zero value otherwise.
func (o *EzsigntemplateCreateObjectV3Response) GetObjDebug() CommonResponseObjDebug {
	if o == nil || IsNil(o.ObjDebug) {
		var ret CommonResponseObjDebug
		return ret
	}
	return *o.ObjDebug
}

// GetObjDebugOk returns a tuple with the ObjDebug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateCreateObjectV3Response) GetObjDebugOk() (*CommonResponseObjDebug, bool) {
	if o == nil || IsNil(o.ObjDebug) {
		return nil, false
	}
	return o.ObjDebug, true
}

// HasObjDebug returns a boolean if a field has been set.
func (o *EzsigntemplateCreateObjectV3Response) HasObjDebug() bool {
	if o != nil && !IsNil(o.ObjDebug) {
		return true
	}

	return false
}

// SetObjDebug gets a reference to the given CommonResponseObjDebug and assigns it to the ObjDebug field.
func (o *EzsigntemplateCreateObjectV3Response) SetObjDebug(v CommonResponseObjDebug) {
	o.ObjDebug = &v
}

// GetMPayload returns the MPayload field value
func (o *EzsigntemplateCreateObjectV3Response) GetMPayload() EzsigntemplateCreateObjectV3ResponseMPayload {
	if o == nil {
		var ret EzsigntemplateCreateObjectV3ResponseMPayload
		return ret
	}

	return o.MPayload
}

// GetMPayloadOk returns a tuple with the MPayload field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateCreateObjectV3Response) GetMPayloadOk() (*EzsigntemplateCreateObjectV3ResponseMPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MPayload, true
}

// SetMPayload sets field value
func (o *EzsigntemplateCreateObjectV3Response) SetMPayload(v EzsigntemplateCreateObjectV3ResponseMPayload) {
	o.MPayload = v
}

func (o EzsigntemplateCreateObjectV3Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigntemplateCreateObjectV3Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objDebugPayload"] = o.ObjDebugPayload
	if !IsNil(o.ObjDebug) {
		toSerialize["objDebug"] = o.ObjDebug
	}
	toSerialize["mPayload"] = o.MPayload
	return toSerialize, nil
}

func (o *EzsigntemplateCreateObjectV3Response) UnmarshalJSON(data []byte) (err error) {
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

	varEzsigntemplateCreateObjectV3Response := _EzsigntemplateCreateObjectV3Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsigntemplateCreateObjectV3Response)

	if err != nil {
		return err
	}

	*o = EzsigntemplateCreateObjectV3Response(varEzsigntemplateCreateObjectV3Response)

	return err
}

type NullableEzsigntemplateCreateObjectV3Response struct {
	value *EzsigntemplateCreateObjectV3Response
	isSet bool
}

func (v NullableEzsigntemplateCreateObjectV3Response) Get() *EzsigntemplateCreateObjectV3Response {
	return v.value
}

func (v *NullableEzsigntemplateCreateObjectV3Response) Set(val *EzsigntemplateCreateObjectV3Response) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigntemplateCreateObjectV3Response) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigntemplateCreateObjectV3Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigntemplateCreateObjectV3Response(val *EzsigntemplateCreateObjectV3Response) *NullableEzsigntemplateCreateObjectV3Response {
	return &NullableEzsigntemplateCreateObjectV3Response{value: val, isSet: true}
}

func (v NullableEzsigntemplateCreateObjectV3Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigntemplateCreateObjectV3Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


