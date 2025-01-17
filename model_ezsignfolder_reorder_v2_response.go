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

// checks if the EzsignfolderReorderV2Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignfolderReorderV2Response{}

// EzsignfolderReorderV2Response Response for POST /2/object/ezsignfolder/{pkiEzsignfolderID}/reorder
type EzsignfolderReorderV2Response struct {
	ObjDebugPayload CommonResponseObjDebugPayload `json:"objDebugPayload"`
	ObjDebug *CommonResponseObjDebug `json:"objDebug,omitempty"`
}

type _EzsignfolderReorderV2Response EzsignfolderReorderV2Response

// NewEzsignfolderReorderV2Response instantiates a new EzsignfolderReorderV2Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignfolderReorderV2Response(objDebugPayload CommonResponseObjDebugPayload) *EzsignfolderReorderV2Response {
	this := EzsignfolderReorderV2Response{}
	this.ObjDebugPayload = objDebugPayload
	return &this
}

// NewEzsignfolderReorderV2ResponseWithDefaults instantiates a new EzsignfolderReorderV2Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignfolderReorderV2ResponseWithDefaults() *EzsignfolderReorderV2Response {
	this := EzsignfolderReorderV2Response{}
	return &this
}

// GetObjDebugPayload returns the ObjDebugPayload field value
func (o *EzsignfolderReorderV2Response) GetObjDebugPayload() CommonResponseObjDebugPayload {
	if o == nil {
		var ret CommonResponseObjDebugPayload
		return ret
	}

	return o.ObjDebugPayload
}

// GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field value
// and a boolean to check if the value has been set.
func (o *EzsignfolderReorderV2Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjDebugPayload, true
}

// SetObjDebugPayload sets field value
func (o *EzsignfolderReorderV2Response) SetObjDebugPayload(v CommonResponseObjDebugPayload) {
	o.ObjDebugPayload = v
}

// GetObjDebug returns the ObjDebug field value if set, zero value otherwise.
func (o *EzsignfolderReorderV2Response) GetObjDebug() CommonResponseObjDebug {
	if o == nil || IsNil(o.ObjDebug) {
		var ret CommonResponseObjDebug
		return ret
	}
	return *o.ObjDebug
}

// GetObjDebugOk returns a tuple with the ObjDebug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignfolderReorderV2Response) GetObjDebugOk() (*CommonResponseObjDebug, bool) {
	if o == nil || IsNil(o.ObjDebug) {
		return nil, false
	}
	return o.ObjDebug, true
}

// HasObjDebug returns a boolean if a field has been set.
func (o *EzsignfolderReorderV2Response) HasObjDebug() bool {
	if o != nil && !IsNil(o.ObjDebug) {
		return true
	}

	return false
}

// SetObjDebug gets a reference to the given CommonResponseObjDebug and assigns it to the ObjDebug field.
func (o *EzsignfolderReorderV2Response) SetObjDebug(v CommonResponseObjDebug) {
	o.ObjDebug = &v
}

func (o EzsignfolderReorderV2Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignfolderReorderV2Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objDebugPayload"] = o.ObjDebugPayload
	if !IsNil(o.ObjDebug) {
		toSerialize["objDebug"] = o.ObjDebug
	}
	return toSerialize, nil
}

func (o *EzsignfolderReorderV2Response) UnmarshalJSON(data []byte) (err error) {
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

	varEzsignfolderReorderV2Response := _EzsignfolderReorderV2Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsignfolderReorderV2Response)

	if err != nil {
		return err
	}

	*o = EzsignfolderReorderV2Response(varEzsignfolderReorderV2Response)

	return err
}

type NullableEzsignfolderReorderV2Response struct {
	value *EzsignfolderReorderV2Response
	isSet bool
}

func (v NullableEzsignfolderReorderV2Response) Get() *EzsignfolderReorderV2Response {
	return v.value
}

func (v *NullableEzsignfolderReorderV2Response) Set(val *EzsignfolderReorderV2Response) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignfolderReorderV2Response) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignfolderReorderV2Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignfolderReorderV2Response(val *EzsignfolderReorderV2Response) *NullableEzsignfolderReorderV2Response {
	return &NullableEzsignfolderReorderV2Response{value: val, isSet: true}
}

func (v NullableEzsignfolderReorderV2Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignfolderReorderV2Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


