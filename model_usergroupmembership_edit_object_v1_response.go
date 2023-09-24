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

// checks if the UsergroupmembershipEditObjectV1Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UsergroupmembershipEditObjectV1Response{}

// UsergroupmembershipEditObjectV1Response Response for PUT /1/object/usergroupmembership/{pkiUsergroupmembershipID}
type UsergroupmembershipEditObjectV1Response struct {
	ObjDebugPayload CommonResponseObjDebugPayload `json:"objDebugPayload"`
	ObjDebug *CommonResponseObjDebug `json:"objDebug,omitempty"`
}

// NewUsergroupmembershipEditObjectV1Response instantiates a new UsergroupmembershipEditObjectV1Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsergroupmembershipEditObjectV1Response(objDebugPayload CommonResponseObjDebugPayload) *UsergroupmembershipEditObjectV1Response {
	this := UsergroupmembershipEditObjectV1Response{}
	this.ObjDebugPayload = objDebugPayload
	return &this
}

// NewUsergroupmembershipEditObjectV1ResponseWithDefaults instantiates a new UsergroupmembershipEditObjectV1Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsergroupmembershipEditObjectV1ResponseWithDefaults() *UsergroupmembershipEditObjectV1Response {
	this := UsergroupmembershipEditObjectV1Response{}
	return &this
}

// GetObjDebugPayload returns the ObjDebugPayload field value
func (o *UsergroupmembershipEditObjectV1Response) GetObjDebugPayload() CommonResponseObjDebugPayload {
	if o == nil {
		var ret CommonResponseObjDebugPayload
		return ret
	}

	return o.ObjDebugPayload
}

// GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field value
// and a boolean to check if the value has been set.
func (o *UsergroupmembershipEditObjectV1Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjDebugPayload, true
}

// SetObjDebugPayload sets field value
func (o *UsergroupmembershipEditObjectV1Response) SetObjDebugPayload(v CommonResponseObjDebugPayload) {
	o.ObjDebugPayload = v
}

// GetObjDebug returns the ObjDebug field value if set, zero value otherwise.
func (o *UsergroupmembershipEditObjectV1Response) GetObjDebug() CommonResponseObjDebug {
	if o == nil || IsNil(o.ObjDebug) {
		var ret CommonResponseObjDebug
		return ret
	}
	return *o.ObjDebug
}

// GetObjDebugOk returns a tuple with the ObjDebug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsergroupmembershipEditObjectV1Response) GetObjDebugOk() (*CommonResponseObjDebug, bool) {
	if o == nil || IsNil(o.ObjDebug) {
		return nil, false
	}
	return o.ObjDebug, true
}

// HasObjDebug returns a boolean if a field has been set.
func (o *UsergroupmembershipEditObjectV1Response) HasObjDebug() bool {
	if o != nil && !IsNil(o.ObjDebug) {
		return true
	}

	return false
}

// SetObjDebug gets a reference to the given CommonResponseObjDebug and assigns it to the ObjDebug field.
func (o *UsergroupmembershipEditObjectV1Response) SetObjDebug(v CommonResponseObjDebug) {
	o.ObjDebug = &v
}

func (o UsergroupmembershipEditObjectV1Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsergroupmembershipEditObjectV1Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objDebugPayload"] = o.ObjDebugPayload
	if !IsNil(o.ObjDebug) {
		toSerialize["objDebug"] = o.ObjDebug
	}
	return toSerialize, nil
}

type NullableUsergroupmembershipEditObjectV1Response struct {
	value *UsergroupmembershipEditObjectV1Response
	isSet bool
}

func (v NullableUsergroupmembershipEditObjectV1Response) Get() *UsergroupmembershipEditObjectV1Response {
	return v.value
}

func (v *NullableUsergroupmembershipEditObjectV1Response) Set(val *UsergroupmembershipEditObjectV1Response) {
	v.value = val
	v.isSet = true
}

func (v NullableUsergroupmembershipEditObjectV1Response) IsSet() bool {
	return v.isSet
}

func (v *NullableUsergroupmembershipEditObjectV1Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsergroupmembershipEditObjectV1Response(val *UsergroupmembershipEditObjectV1Response) *NullableUsergroupmembershipEditObjectV1Response {
	return &NullableUsergroupmembershipEditObjectV1Response{value: val, isSet: true}
}

func (v NullableUsergroupmembershipEditObjectV1Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsergroupmembershipEditObjectV1Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


