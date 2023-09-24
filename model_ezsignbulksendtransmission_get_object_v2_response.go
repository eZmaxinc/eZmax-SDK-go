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

// checks if the EzsignbulksendtransmissionGetObjectV2Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignbulksendtransmissionGetObjectV2Response{}

// EzsignbulksendtransmissionGetObjectV2Response Response for GET /2/object/ezsignbulksendtransmission/{pkiEzsignbulksendtransmissionID}
type EzsignbulksendtransmissionGetObjectV2Response struct {
	ObjDebugPayload CommonResponseObjDebugPayload `json:"objDebugPayload"`
	ObjDebug *CommonResponseObjDebug `json:"objDebug,omitempty"`
	MPayload EzsignbulksendtransmissionGetObjectV2ResponseMPayload `json:"mPayload"`
}

// NewEzsignbulksendtransmissionGetObjectV2Response instantiates a new EzsignbulksendtransmissionGetObjectV2Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignbulksendtransmissionGetObjectV2Response(objDebugPayload CommonResponseObjDebugPayload, mPayload EzsignbulksendtransmissionGetObjectV2ResponseMPayload) *EzsignbulksendtransmissionGetObjectV2Response {
	this := EzsignbulksendtransmissionGetObjectV2Response{}
	this.ObjDebugPayload = objDebugPayload
	this.MPayload = mPayload
	return &this
}

// NewEzsignbulksendtransmissionGetObjectV2ResponseWithDefaults instantiates a new EzsignbulksendtransmissionGetObjectV2Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignbulksendtransmissionGetObjectV2ResponseWithDefaults() *EzsignbulksendtransmissionGetObjectV2Response {
	this := EzsignbulksendtransmissionGetObjectV2Response{}
	return &this
}

// GetObjDebugPayload returns the ObjDebugPayload field value
func (o *EzsignbulksendtransmissionGetObjectV2Response) GetObjDebugPayload() CommonResponseObjDebugPayload {
	if o == nil {
		var ret CommonResponseObjDebugPayload
		return ret
	}

	return o.ObjDebugPayload
}

// GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field value
// and a boolean to check if the value has been set.
func (o *EzsignbulksendtransmissionGetObjectV2Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjDebugPayload, true
}

// SetObjDebugPayload sets field value
func (o *EzsignbulksendtransmissionGetObjectV2Response) SetObjDebugPayload(v CommonResponseObjDebugPayload) {
	o.ObjDebugPayload = v
}

// GetObjDebug returns the ObjDebug field value if set, zero value otherwise.
func (o *EzsignbulksendtransmissionGetObjectV2Response) GetObjDebug() CommonResponseObjDebug {
	if o == nil || IsNil(o.ObjDebug) {
		var ret CommonResponseObjDebug
		return ret
	}
	return *o.ObjDebug
}

// GetObjDebugOk returns a tuple with the ObjDebug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignbulksendtransmissionGetObjectV2Response) GetObjDebugOk() (*CommonResponseObjDebug, bool) {
	if o == nil || IsNil(o.ObjDebug) {
		return nil, false
	}
	return o.ObjDebug, true
}

// HasObjDebug returns a boolean if a field has been set.
func (o *EzsignbulksendtransmissionGetObjectV2Response) HasObjDebug() bool {
	if o != nil && !IsNil(o.ObjDebug) {
		return true
	}

	return false
}

// SetObjDebug gets a reference to the given CommonResponseObjDebug and assigns it to the ObjDebug field.
func (o *EzsignbulksendtransmissionGetObjectV2Response) SetObjDebug(v CommonResponseObjDebug) {
	o.ObjDebug = &v
}

// GetMPayload returns the MPayload field value
func (o *EzsignbulksendtransmissionGetObjectV2Response) GetMPayload() EzsignbulksendtransmissionGetObjectV2ResponseMPayload {
	if o == nil {
		var ret EzsignbulksendtransmissionGetObjectV2ResponseMPayload
		return ret
	}

	return o.MPayload
}

// GetMPayloadOk returns a tuple with the MPayload field value
// and a boolean to check if the value has been set.
func (o *EzsignbulksendtransmissionGetObjectV2Response) GetMPayloadOk() (*EzsignbulksendtransmissionGetObjectV2ResponseMPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MPayload, true
}

// SetMPayload sets field value
func (o *EzsignbulksendtransmissionGetObjectV2Response) SetMPayload(v EzsignbulksendtransmissionGetObjectV2ResponseMPayload) {
	o.MPayload = v
}

func (o EzsignbulksendtransmissionGetObjectV2Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignbulksendtransmissionGetObjectV2Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objDebugPayload"] = o.ObjDebugPayload
	if !IsNil(o.ObjDebug) {
		toSerialize["objDebug"] = o.ObjDebug
	}
	toSerialize["mPayload"] = o.MPayload
	return toSerialize, nil
}

type NullableEzsignbulksendtransmissionGetObjectV2Response struct {
	value *EzsignbulksendtransmissionGetObjectV2Response
	isSet bool
}

func (v NullableEzsignbulksendtransmissionGetObjectV2Response) Get() *EzsignbulksendtransmissionGetObjectV2Response {
	return v.value
}

func (v *NullableEzsignbulksendtransmissionGetObjectV2Response) Set(val *EzsignbulksendtransmissionGetObjectV2Response) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignbulksendtransmissionGetObjectV2Response) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignbulksendtransmissionGetObjectV2Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignbulksendtransmissionGetObjectV2Response(val *EzsignbulksendtransmissionGetObjectV2Response) *NullableEzsignbulksendtransmissionGetObjectV2Response {
	return &NullableEzsignbulksendtransmissionGetObjectV2Response{value: val, isSet: true}
}

func (v NullableEzsignbulksendtransmissionGetObjectV2Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignbulksendtransmissionGetObjectV2Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


