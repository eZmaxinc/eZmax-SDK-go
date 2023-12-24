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

// checks if the EzsigntemplateGetListV1Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigntemplateGetListV1Response{}

// EzsigntemplateGetListV1Response Response for GET /1/object/ezsigntemplate/getList
type EzsigntemplateGetListV1Response struct {
	ObjDebugPayload CommonResponseObjDebugPayloadGetList `json:"objDebugPayload"`
	ObjDebug *CommonResponseObjDebug `json:"objDebug,omitempty"`
	MPayload EzsigntemplateGetListV1ResponseMPayload `json:"mPayload"`
}

type _EzsigntemplateGetListV1Response EzsigntemplateGetListV1Response

// NewEzsigntemplateGetListV1Response instantiates a new EzsigntemplateGetListV1Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigntemplateGetListV1Response(objDebugPayload CommonResponseObjDebugPayloadGetList, mPayload EzsigntemplateGetListV1ResponseMPayload) *EzsigntemplateGetListV1Response {
	this := EzsigntemplateGetListV1Response{}
	this.ObjDebugPayload = objDebugPayload
	this.MPayload = mPayload
	return &this
}

// NewEzsigntemplateGetListV1ResponseWithDefaults instantiates a new EzsigntemplateGetListV1Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigntemplateGetListV1ResponseWithDefaults() *EzsigntemplateGetListV1Response {
	this := EzsigntemplateGetListV1Response{}
	return &this
}

// GetObjDebugPayload returns the ObjDebugPayload field value
func (o *EzsigntemplateGetListV1Response) GetObjDebugPayload() CommonResponseObjDebugPayloadGetList {
	if o == nil {
		var ret CommonResponseObjDebugPayloadGetList
		return ret
	}

	return o.ObjDebugPayload
}

// GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateGetListV1Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayloadGetList, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjDebugPayload, true
}

// SetObjDebugPayload sets field value
func (o *EzsigntemplateGetListV1Response) SetObjDebugPayload(v CommonResponseObjDebugPayloadGetList) {
	o.ObjDebugPayload = v
}

// GetObjDebug returns the ObjDebug field value if set, zero value otherwise.
func (o *EzsigntemplateGetListV1Response) GetObjDebug() CommonResponseObjDebug {
	if o == nil || IsNil(o.ObjDebug) {
		var ret CommonResponseObjDebug
		return ret
	}
	return *o.ObjDebug
}

// GetObjDebugOk returns a tuple with the ObjDebug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateGetListV1Response) GetObjDebugOk() (*CommonResponseObjDebug, bool) {
	if o == nil || IsNil(o.ObjDebug) {
		return nil, false
	}
	return o.ObjDebug, true
}

// HasObjDebug returns a boolean if a field has been set.
func (o *EzsigntemplateGetListV1Response) HasObjDebug() bool {
	if o != nil && !IsNil(o.ObjDebug) {
		return true
	}

	return false
}

// SetObjDebug gets a reference to the given CommonResponseObjDebug and assigns it to the ObjDebug field.
func (o *EzsigntemplateGetListV1Response) SetObjDebug(v CommonResponseObjDebug) {
	o.ObjDebug = &v
}

// GetMPayload returns the MPayload field value
func (o *EzsigntemplateGetListV1Response) GetMPayload() EzsigntemplateGetListV1ResponseMPayload {
	if o == nil {
		var ret EzsigntemplateGetListV1ResponseMPayload
		return ret
	}

	return o.MPayload
}

// GetMPayloadOk returns a tuple with the MPayload field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateGetListV1Response) GetMPayloadOk() (*EzsigntemplateGetListV1ResponseMPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MPayload, true
}

// SetMPayload sets field value
func (o *EzsigntemplateGetListV1Response) SetMPayload(v EzsigntemplateGetListV1ResponseMPayload) {
	o.MPayload = v
}

func (o EzsigntemplateGetListV1Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigntemplateGetListV1Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objDebugPayload"] = o.ObjDebugPayload
	if !IsNil(o.ObjDebug) {
		toSerialize["objDebug"] = o.ObjDebug
	}
	toSerialize["mPayload"] = o.MPayload
	return toSerialize, nil
}

func (o *EzsigntemplateGetListV1Response) UnmarshalJSON(data []byte) (err error) {
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

	varEzsigntemplateGetListV1Response := _EzsigntemplateGetListV1Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsigntemplateGetListV1Response)

	if err != nil {
		return err
	}

	*o = EzsigntemplateGetListV1Response(varEzsigntemplateGetListV1Response)

	return err
}

type NullableEzsigntemplateGetListV1Response struct {
	value *EzsigntemplateGetListV1Response
	isSet bool
}

func (v NullableEzsigntemplateGetListV1Response) Get() *EzsigntemplateGetListV1Response {
	return v.value
}

func (v *NullableEzsigntemplateGetListV1Response) Set(val *EzsigntemplateGetListV1Response) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigntemplateGetListV1Response) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigntemplateGetListV1Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigntemplateGetListV1Response(val *EzsigntemplateGetListV1Response) *NullableEzsigntemplateGetListV1Response {
	return &NullableEzsigntemplateGetListV1Response{value: val, isSet: true}
}

func (v NullableEzsigntemplateGetListV1Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigntemplateGetListV1Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


