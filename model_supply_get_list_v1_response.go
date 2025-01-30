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

// checks if the SupplyGetListV1Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SupplyGetListV1Response{}

// SupplyGetListV1Response Response for GET /1/object/supply/getList
type SupplyGetListV1Response struct {
	ObjDebugPayload CommonResponseObjDebugPayloadGetList `json:"objDebugPayload"`
	ObjDebug *CommonResponseObjDebug `json:"objDebug,omitempty"`
	MPayload SupplyGetListV1ResponseMPayload `json:"mPayload"`
}

type _SupplyGetListV1Response SupplyGetListV1Response

// NewSupplyGetListV1Response instantiates a new SupplyGetListV1Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSupplyGetListV1Response(objDebugPayload CommonResponseObjDebugPayloadGetList, mPayload SupplyGetListV1ResponseMPayload) *SupplyGetListV1Response {
	this := SupplyGetListV1Response{}
	this.ObjDebugPayload = objDebugPayload
	this.MPayload = mPayload
	return &this
}

// NewSupplyGetListV1ResponseWithDefaults instantiates a new SupplyGetListV1Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSupplyGetListV1ResponseWithDefaults() *SupplyGetListV1Response {
	this := SupplyGetListV1Response{}
	return &this
}

// GetObjDebugPayload returns the ObjDebugPayload field value
func (o *SupplyGetListV1Response) GetObjDebugPayload() CommonResponseObjDebugPayloadGetList {
	if o == nil {
		var ret CommonResponseObjDebugPayloadGetList
		return ret
	}

	return o.ObjDebugPayload
}

// GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field value
// and a boolean to check if the value has been set.
func (o *SupplyGetListV1Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayloadGetList, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjDebugPayload, true
}

// SetObjDebugPayload sets field value
func (o *SupplyGetListV1Response) SetObjDebugPayload(v CommonResponseObjDebugPayloadGetList) {
	o.ObjDebugPayload = v
}

// GetObjDebug returns the ObjDebug field value if set, zero value otherwise.
func (o *SupplyGetListV1Response) GetObjDebug() CommonResponseObjDebug {
	if o == nil || IsNil(o.ObjDebug) {
		var ret CommonResponseObjDebug
		return ret
	}
	return *o.ObjDebug
}

// GetObjDebugOk returns a tuple with the ObjDebug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupplyGetListV1Response) GetObjDebugOk() (*CommonResponseObjDebug, bool) {
	if o == nil || IsNil(o.ObjDebug) {
		return nil, false
	}
	return o.ObjDebug, true
}

// HasObjDebug returns a boolean if a field has been set.
func (o *SupplyGetListV1Response) HasObjDebug() bool {
	if o != nil && !IsNil(o.ObjDebug) {
		return true
	}

	return false
}

// SetObjDebug gets a reference to the given CommonResponseObjDebug and assigns it to the ObjDebug field.
func (o *SupplyGetListV1Response) SetObjDebug(v CommonResponseObjDebug) {
	o.ObjDebug = &v
}

// GetMPayload returns the MPayload field value
func (o *SupplyGetListV1Response) GetMPayload() SupplyGetListV1ResponseMPayload {
	if o == nil {
		var ret SupplyGetListV1ResponseMPayload
		return ret
	}

	return o.MPayload
}

// GetMPayloadOk returns a tuple with the MPayload field value
// and a boolean to check if the value has been set.
func (o *SupplyGetListV1Response) GetMPayloadOk() (*SupplyGetListV1ResponseMPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MPayload, true
}

// SetMPayload sets field value
func (o *SupplyGetListV1Response) SetMPayload(v SupplyGetListV1ResponseMPayload) {
	o.MPayload = v
}

func (o SupplyGetListV1Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SupplyGetListV1Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objDebugPayload"] = o.ObjDebugPayload
	if !IsNil(o.ObjDebug) {
		toSerialize["objDebug"] = o.ObjDebug
	}
	toSerialize["mPayload"] = o.MPayload
	return toSerialize, nil
}

func (o *SupplyGetListV1Response) UnmarshalJSON(data []byte) (err error) {
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

	varSupplyGetListV1Response := _SupplyGetListV1Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSupplyGetListV1Response)

	if err != nil {
		return err
	}

	*o = SupplyGetListV1Response(varSupplyGetListV1Response)

	return err
}

type NullableSupplyGetListV1Response struct {
	value *SupplyGetListV1Response
	isSet bool
}

func (v NullableSupplyGetListV1Response) Get() *SupplyGetListV1Response {
	return v.value
}

func (v *NullableSupplyGetListV1Response) Set(val *SupplyGetListV1Response) {
	v.value = val
	v.isSet = true
}

func (v NullableSupplyGetListV1Response) IsSet() bool {
	return v.isSet
}

func (v *NullableSupplyGetListV1Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSupplyGetListV1Response(val *SupplyGetListV1Response) *NullableSupplyGetListV1Response {
	return &NullableSupplyGetListV1Response{value: val, isSet: true}
}

func (v NullableSupplyGetListV1Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSupplyGetListV1Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


