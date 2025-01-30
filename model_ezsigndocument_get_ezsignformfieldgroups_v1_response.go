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

// checks if the EzsigndocumentGetEzsignformfieldgroupsV1Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigndocumentGetEzsignformfieldgroupsV1Response{}

// EzsigndocumentGetEzsignformfieldgroupsV1Response Response for GET /1/object/ezsigndocument/{pkiEzsigndocument}/getEzsignformfieldgroups
type EzsigndocumentGetEzsignformfieldgroupsV1Response struct {
	ObjDebugPayload CommonResponseObjDebugPayload `json:"objDebugPayload"`
	ObjDebug *CommonResponseObjDebug `json:"objDebug,omitempty"`
	MPayload EzsigndocumentGetEzsignformfieldgroupsV1ResponseMPayload `json:"mPayload"`
}

type _EzsigndocumentGetEzsignformfieldgroupsV1Response EzsigndocumentGetEzsignformfieldgroupsV1Response

// NewEzsigndocumentGetEzsignformfieldgroupsV1Response instantiates a new EzsigndocumentGetEzsignformfieldgroupsV1Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigndocumentGetEzsignformfieldgroupsV1Response(objDebugPayload CommonResponseObjDebugPayload, mPayload EzsigndocumentGetEzsignformfieldgroupsV1ResponseMPayload) *EzsigndocumentGetEzsignformfieldgroupsV1Response {
	this := EzsigndocumentGetEzsignformfieldgroupsV1Response{}
	this.ObjDebugPayload = objDebugPayload
	this.MPayload = mPayload
	return &this
}

// NewEzsigndocumentGetEzsignformfieldgroupsV1ResponseWithDefaults instantiates a new EzsigndocumentGetEzsignformfieldgroupsV1Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigndocumentGetEzsignformfieldgroupsV1ResponseWithDefaults() *EzsigndocumentGetEzsignformfieldgroupsV1Response {
	this := EzsigndocumentGetEzsignformfieldgroupsV1Response{}
	return &this
}

// GetObjDebugPayload returns the ObjDebugPayload field value
func (o *EzsigndocumentGetEzsignformfieldgroupsV1Response) GetObjDebugPayload() CommonResponseObjDebugPayload {
	if o == nil {
		var ret CommonResponseObjDebugPayload
		return ret
	}

	return o.ObjDebugPayload
}

// GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentGetEzsignformfieldgroupsV1Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjDebugPayload, true
}

// SetObjDebugPayload sets field value
func (o *EzsigndocumentGetEzsignformfieldgroupsV1Response) SetObjDebugPayload(v CommonResponseObjDebugPayload) {
	o.ObjDebugPayload = v
}

// GetObjDebug returns the ObjDebug field value if set, zero value otherwise.
func (o *EzsigndocumentGetEzsignformfieldgroupsV1Response) GetObjDebug() CommonResponseObjDebug {
	if o == nil || IsNil(o.ObjDebug) {
		var ret CommonResponseObjDebug
		return ret
	}
	return *o.ObjDebug
}

// GetObjDebugOk returns a tuple with the ObjDebug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigndocumentGetEzsignformfieldgroupsV1Response) GetObjDebugOk() (*CommonResponseObjDebug, bool) {
	if o == nil || IsNil(o.ObjDebug) {
		return nil, false
	}
	return o.ObjDebug, true
}

// HasObjDebug returns a boolean if a field has been set.
func (o *EzsigndocumentGetEzsignformfieldgroupsV1Response) HasObjDebug() bool {
	if o != nil && !IsNil(o.ObjDebug) {
		return true
	}

	return false
}

// SetObjDebug gets a reference to the given CommonResponseObjDebug and assigns it to the ObjDebug field.
func (o *EzsigndocumentGetEzsignformfieldgroupsV1Response) SetObjDebug(v CommonResponseObjDebug) {
	o.ObjDebug = &v
}

// GetMPayload returns the MPayload field value
func (o *EzsigndocumentGetEzsignformfieldgroupsV1Response) GetMPayload() EzsigndocumentGetEzsignformfieldgroupsV1ResponseMPayload {
	if o == nil {
		var ret EzsigndocumentGetEzsignformfieldgroupsV1ResponseMPayload
		return ret
	}

	return o.MPayload
}

// GetMPayloadOk returns a tuple with the MPayload field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentGetEzsignformfieldgroupsV1Response) GetMPayloadOk() (*EzsigndocumentGetEzsignformfieldgroupsV1ResponseMPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MPayload, true
}

// SetMPayload sets field value
func (o *EzsigndocumentGetEzsignformfieldgroupsV1Response) SetMPayload(v EzsigndocumentGetEzsignformfieldgroupsV1ResponseMPayload) {
	o.MPayload = v
}

func (o EzsigndocumentGetEzsignformfieldgroupsV1Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigndocumentGetEzsignformfieldgroupsV1Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objDebugPayload"] = o.ObjDebugPayload
	if !IsNil(o.ObjDebug) {
		toSerialize["objDebug"] = o.ObjDebug
	}
	toSerialize["mPayload"] = o.MPayload
	return toSerialize, nil
}

func (o *EzsigndocumentGetEzsignformfieldgroupsV1Response) UnmarshalJSON(data []byte) (err error) {
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

	varEzsigndocumentGetEzsignformfieldgroupsV1Response := _EzsigndocumentGetEzsignformfieldgroupsV1Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsigndocumentGetEzsignformfieldgroupsV1Response)

	if err != nil {
		return err
	}

	*o = EzsigndocumentGetEzsignformfieldgroupsV1Response(varEzsigndocumentGetEzsignformfieldgroupsV1Response)

	return err
}

type NullableEzsigndocumentGetEzsignformfieldgroupsV1Response struct {
	value *EzsigndocumentGetEzsignformfieldgroupsV1Response
	isSet bool
}

func (v NullableEzsigndocumentGetEzsignformfieldgroupsV1Response) Get() *EzsigndocumentGetEzsignformfieldgroupsV1Response {
	return v.value
}

func (v *NullableEzsigndocumentGetEzsignformfieldgroupsV1Response) Set(val *EzsigndocumentGetEzsignformfieldgroupsV1Response) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigndocumentGetEzsignformfieldgroupsV1Response) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigndocumentGetEzsignformfieldgroupsV1Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigndocumentGetEzsignformfieldgroupsV1Response(val *EzsigndocumentGetEzsignformfieldgroupsV1Response) *NullableEzsigndocumentGetEzsignformfieldgroupsV1Response {
	return &NullableEzsigndocumentGetEzsignformfieldgroupsV1Response{value: val, isSet: true}
}

func (v NullableEzsigndocumentGetEzsignformfieldgroupsV1Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigndocumentGetEzsignformfieldgroupsV1Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


