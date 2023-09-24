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
)

// checks if the EzsignsignergroupGetObjectV2Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignsignergroupGetObjectV2Response{}

// EzsignsignergroupGetObjectV2Response Response for GET /2/object/ezsignsignergroup/{pkiEzsignsignergroupID}
type EzsignsignergroupGetObjectV2Response struct {
	ObjDebugPayload CommonResponseObjDebugPayload `json:"objDebugPayload"`
	ObjDebug *CommonResponseObjDebug `json:"objDebug,omitempty"`
	MPayload EzsignsignergroupGetObjectV2ResponseMPayload `json:"mPayload"`
}

// NewEzsignsignergroupGetObjectV2Response instantiates a new EzsignsignergroupGetObjectV2Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignsignergroupGetObjectV2Response(objDebugPayload CommonResponseObjDebugPayload, mPayload EzsignsignergroupGetObjectV2ResponseMPayload) *EzsignsignergroupGetObjectV2Response {
	this := EzsignsignergroupGetObjectV2Response{}
	this.ObjDebugPayload = objDebugPayload
	this.MPayload = mPayload
	return &this
}

// NewEzsignsignergroupGetObjectV2ResponseWithDefaults instantiates a new EzsignsignergroupGetObjectV2Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignsignergroupGetObjectV2ResponseWithDefaults() *EzsignsignergroupGetObjectV2Response {
	this := EzsignsignergroupGetObjectV2Response{}
	return &this
}

// GetObjDebugPayload returns the ObjDebugPayload field value
func (o *EzsignsignergroupGetObjectV2Response) GetObjDebugPayload() CommonResponseObjDebugPayload {
	if o == nil {
		var ret CommonResponseObjDebugPayload
		return ret
	}

	return o.ObjDebugPayload
}

// GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field value
// and a boolean to check if the value has been set.
func (o *EzsignsignergroupGetObjectV2Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjDebugPayload, true
}

// SetObjDebugPayload sets field value
func (o *EzsignsignergroupGetObjectV2Response) SetObjDebugPayload(v CommonResponseObjDebugPayload) {
	o.ObjDebugPayload = v
}

// GetObjDebug returns the ObjDebug field value if set, zero value otherwise.
func (o *EzsignsignergroupGetObjectV2Response) GetObjDebug() CommonResponseObjDebug {
	if o == nil || IsNil(o.ObjDebug) {
		var ret CommonResponseObjDebug
		return ret
	}
	return *o.ObjDebug
}

// GetObjDebugOk returns a tuple with the ObjDebug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignsignergroupGetObjectV2Response) GetObjDebugOk() (*CommonResponseObjDebug, bool) {
	if o == nil || IsNil(o.ObjDebug) {
		return nil, false
	}
	return o.ObjDebug, true
}

// HasObjDebug returns a boolean if a field has been set.
func (o *EzsignsignergroupGetObjectV2Response) HasObjDebug() bool {
	if o != nil && !IsNil(o.ObjDebug) {
		return true
	}

	return false
}

// SetObjDebug gets a reference to the given CommonResponseObjDebug and assigns it to the ObjDebug field.
func (o *EzsignsignergroupGetObjectV2Response) SetObjDebug(v CommonResponseObjDebug) {
	o.ObjDebug = &v
}

// GetMPayload returns the MPayload field value
func (o *EzsignsignergroupGetObjectV2Response) GetMPayload() EzsignsignergroupGetObjectV2ResponseMPayload {
	if o == nil {
		var ret EzsignsignergroupGetObjectV2ResponseMPayload
		return ret
	}

	return o.MPayload
}

// GetMPayloadOk returns a tuple with the MPayload field value
// and a boolean to check if the value has been set.
func (o *EzsignsignergroupGetObjectV2Response) GetMPayloadOk() (*EzsignsignergroupGetObjectV2ResponseMPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MPayload, true
}

// SetMPayload sets field value
func (o *EzsignsignergroupGetObjectV2Response) SetMPayload(v EzsignsignergroupGetObjectV2ResponseMPayload) {
	o.MPayload = v
}

func (o EzsignsignergroupGetObjectV2Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignsignergroupGetObjectV2Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objDebugPayload"] = o.ObjDebugPayload
	if !IsNil(o.ObjDebug) {
		toSerialize["objDebug"] = o.ObjDebug
	}
	toSerialize["mPayload"] = o.MPayload
	return toSerialize, nil
}

type NullableEzsignsignergroupGetObjectV2Response struct {
	value *EzsignsignergroupGetObjectV2Response
	isSet bool
}

func (v NullableEzsignsignergroupGetObjectV2Response) Get() *EzsignsignergroupGetObjectV2Response {
	return v.value
}

func (v *NullableEzsignsignergroupGetObjectV2Response) Set(val *EzsignsignergroupGetObjectV2Response) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignsignergroupGetObjectV2Response) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignsignergroupGetObjectV2Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignsignergroupGetObjectV2Response(val *EzsignsignergroupGetObjectV2Response) *NullableEzsignsignergroupGetObjectV2Response {
	return &NullableEzsignsignergroupGetObjectV2Response{value: val, isSet: true}
}

func (v NullableEzsignsignergroupGetObjectV2Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignsignergroupGetObjectV2Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


