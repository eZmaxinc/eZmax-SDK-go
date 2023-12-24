/*
eZmax API Definition (Full)

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.2.0
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the EzsignsigningreasonGetListV1Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignsigningreasonGetListV1Response{}

// EzsignsigningreasonGetListV1Response Response for GET /1/object/ezsignsigningreason/getList
type EzsignsigningreasonGetListV1Response struct {
	ObjDebugPayload CommonResponseObjDebugPayloadGetList `json:"objDebugPayload"`
	ObjDebug *CommonResponseObjDebug `json:"objDebug,omitempty"`
	MPayload EzsignsigningreasonGetListV1ResponseMPayload `json:"mPayload"`
}

type _EzsignsigningreasonGetListV1Response EzsignsigningreasonGetListV1Response

// NewEzsignsigningreasonGetListV1Response instantiates a new EzsignsigningreasonGetListV1Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignsigningreasonGetListV1Response(objDebugPayload CommonResponseObjDebugPayloadGetList, mPayload EzsignsigningreasonGetListV1ResponseMPayload) *EzsignsigningreasonGetListV1Response {
	this := EzsignsigningreasonGetListV1Response{}
	this.ObjDebugPayload = objDebugPayload
	this.MPayload = mPayload
	return &this
}

// NewEzsignsigningreasonGetListV1ResponseWithDefaults instantiates a new EzsignsigningreasonGetListV1Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignsigningreasonGetListV1ResponseWithDefaults() *EzsignsigningreasonGetListV1Response {
	this := EzsignsigningreasonGetListV1Response{}
	return &this
}

// GetObjDebugPayload returns the ObjDebugPayload field value
func (o *EzsignsigningreasonGetListV1Response) GetObjDebugPayload() CommonResponseObjDebugPayloadGetList {
	if o == nil {
		var ret CommonResponseObjDebugPayloadGetList
		return ret
	}

	return o.ObjDebugPayload
}

// GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field value
// and a boolean to check if the value has been set.
func (o *EzsignsigningreasonGetListV1Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayloadGetList, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjDebugPayload, true
}

// SetObjDebugPayload sets field value
func (o *EzsignsigningreasonGetListV1Response) SetObjDebugPayload(v CommonResponseObjDebugPayloadGetList) {
	o.ObjDebugPayload = v
}

// GetObjDebug returns the ObjDebug field value if set, zero value otherwise.
func (o *EzsignsigningreasonGetListV1Response) GetObjDebug() CommonResponseObjDebug {
	if o == nil || IsNil(o.ObjDebug) {
		var ret CommonResponseObjDebug
		return ret
	}
	return *o.ObjDebug
}

// GetObjDebugOk returns a tuple with the ObjDebug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignsigningreasonGetListV1Response) GetObjDebugOk() (*CommonResponseObjDebug, bool) {
	if o == nil || IsNil(o.ObjDebug) {
		return nil, false
	}
	return o.ObjDebug, true
}

// HasObjDebug returns a boolean if a field has been set.
func (o *EzsignsigningreasonGetListV1Response) HasObjDebug() bool {
	if o != nil && !IsNil(o.ObjDebug) {
		return true
	}

	return false
}

// SetObjDebug gets a reference to the given CommonResponseObjDebug and assigns it to the ObjDebug field.
func (o *EzsignsigningreasonGetListV1Response) SetObjDebug(v CommonResponseObjDebug) {
	o.ObjDebug = &v
}

// GetMPayload returns the MPayload field value
func (o *EzsignsigningreasonGetListV1Response) GetMPayload() EzsignsigningreasonGetListV1ResponseMPayload {
	if o == nil {
		var ret EzsignsigningreasonGetListV1ResponseMPayload
		return ret
	}

	return o.MPayload
}

// GetMPayloadOk returns a tuple with the MPayload field value
// and a boolean to check if the value has been set.
func (o *EzsignsigningreasonGetListV1Response) GetMPayloadOk() (*EzsignsigningreasonGetListV1ResponseMPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MPayload, true
}

// SetMPayload sets field value
func (o *EzsignsigningreasonGetListV1Response) SetMPayload(v EzsignsigningreasonGetListV1ResponseMPayload) {
	o.MPayload = v
}

func (o EzsignsigningreasonGetListV1Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignsigningreasonGetListV1Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objDebugPayload"] = o.ObjDebugPayload
	if !IsNil(o.ObjDebug) {
		toSerialize["objDebug"] = o.ObjDebug
	}
	toSerialize["mPayload"] = o.MPayload
	return toSerialize, nil
}

func (o *EzsignsigningreasonGetListV1Response) UnmarshalJSON(data []byte) (err error) {
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

	varEzsignsigningreasonGetListV1Response := _EzsignsigningreasonGetListV1Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsignsigningreasonGetListV1Response)

	if err != nil {
		return err
	}

	*o = EzsignsigningreasonGetListV1Response(varEzsignsigningreasonGetListV1Response)

	return err
}

type NullableEzsignsigningreasonGetListV1Response struct {
	value *EzsignsigningreasonGetListV1Response
	isSet bool
}

func (v NullableEzsignsigningreasonGetListV1Response) Get() *EzsignsigningreasonGetListV1Response {
	return v.value
}

func (v *NullableEzsignsigningreasonGetListV1Response) Set(val *EzsignsigningreasonGetListV1Response) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignsigningreasonGetListV1Response) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignsigningreasonGetListV1Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignsigningreasonGetListV1Response(val *EzsignsigningreasonGetListV1Response) *NullableEzsignsigningreasonGetListV1Response {
	return &NullableEzsignsigningreasonGetListV1Response{value: val, isSet: true}
}

func (v NullableEzsignsigningreasonGetListV1Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignsigningreasonGetListV1Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


