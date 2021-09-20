/*
eZmax API Definition

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.0.47
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
)

// AuthenticateAuthenticateV2Response Response for the /2/module/authenticate/authenticate API Request
type AuthenticateAuthenticateV2Response struct {
	MPayload AuthenticateAuthenticateV2ResponseMPayload `json:"mPayload"`
	ObjDebugPayload *CommonResponseObjDebugPayload `json:"objDebugPayload,omitempty"`
	ObjDebug *CommonResponseObjDebug `json:"objDebug,omitempty"`
}

// NewAuthenticateAuthenticateV2Response instantiates a new AuthenticateAuthenticateV2Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticateAuthenticateV2Response(mPayload AuthenticateAuthenticateV2ResponseMPayload) *AuthenticateAuthenticateV2Response {
	this := AuthenticateAuthenticateV2Response{}
	this.MPayload = mPayload
	return &this
}

// NewAuthenticateAuthenticateV2ResponseWithDefaults instantiates a new AuthenticateAuthenticateV2Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticateAuthenticateV2ResponseWithDefaults() *AuthenticateAuthenticateV2Response {
	this := AuthenticateAuthenticateV2Response{}
	return &this
}

// GetMPayload returns the MPayload field value
func (o *AuthenticateAuthenticateV2Response) GetMPayload() AuthenticateAuthenticateV2ResponseMPayload {
	if o == nil {
		var ret AuthenticateAuthenticateV2ResponseMPayload
		return ret
	}

	return o.MPayload
}

// GetMPayloadOk returns a tuple with the MPayload field value
// and a boolean to check if the value has been set.
func (o *AuthenticateAuthenticateV2Response) GetMPayloadOk() (*AuthenticateAuthenticateV2ResponseMPayload, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MPayload, true
}

// SetMPayload sets field value
func (o *AuthenticateAuthenticateV2Response) SetMPayload(v AuthenticateAuthenticateV2ResponseMPayload) {
	o.MPayload = v
}

// GetObjDebugPayload returns the ObjDebugPayload field value if set, zero value otherwise.
func (o *AuthenticateAuthenticateV2Response) GetObjDebugPayload() CommonResponseObjDebugPayload {
	if o == nil || o.ObjDebugPayload == nil {
		var ret CommonResponseObjDebugPayload
		return ret
	}
	return *o.ObjDebugPayload
}

// GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticateAuthenticateV2Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayload, bool) {
	if o == nil || o.ObjDebugPayload == nil {
		return nil, false
	}
	return o.ObjDebugPayload, true
}

// HasObjDebugPayload returns a boolean if a field has been set.
func (o *AuthenticateAuthenticateV2Response) HasObjDebugPayload() bool {
	if o != nil && o.ObjDebugPayload != nil {
		return true
	}

	return false
}

// SetObjDebugPayload gets a reference to the given CommonResponseObjDebugPayload and assigns it to the ObjDebugPayload field.
func (o *AuthenticateAuthenticateV2Response) SetObjDebugPayload(v CommonResponseObjDebugPayload) {
	o.ObjDebugPayload = &v
}

// GetObjDebug returns the ObjDebug field value if set, zero value otherwise.
func (o *AuthenticateAuthenticateV2Response) GetObjDebug() CommonResponseObjDebug {
	if o == nil || o.ObjDebug == nil {
		var ret CommonResponseObjDebug
		return ret
	}
	return *o.ObjDebug
}

// GetObjDebugOk returns a tuple with the ObjDebug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticateAuthenticateV2Response) GetObjDebugOk() (*CommonResponseObjDebug, bool) {
	if o == nil || o.ObjDebug == nil {
		return nil, false
	}
	return o.ObjDebug, true
}

// HasObjDebug returns a boolean if a field has been set.
func (o *AuthenticateAuthenticateV2Response) HasObjDebug() bool {
	if o != nil && o.ObjDebug != nil {
		return true
	}

	return false
}

// SetObjDebug gets a reference to the given CommonResponseObjDebug and assigns it to the ObjDebug field.
func (o *AuthenticateAuthenticateV2Response) SetObjDebug(v CommonResponseObjDebug) {
	o.ObjDebug = &v
}

func (o AuthenticateAuthenticateV2Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["mPayload"] = o.MPayload
	}
	if o.ObjDebugPayload != nil {
		toSerialize["objDebugPayload"] = o.ObjDebugPayload
	}
	if o.ObjDebug != nil {
		toSerialize["objDebug"] = o.ObjDebug
	}
	return json.Marshal(toSerialize)
}

type NullableAuthenticateAuthenticateV2Response struct {
	value *AuthenticateAuthenticateV2Response
	isSet bool
}

func (v NullableAuthenticateAuthenticateV2Response) Get() *AuthenticateAuthenticateV2Response {
	return v.value
}

func (v *NullableAuthenticateAuthenticateV2Response) Set(val *AuthenticateAuthenticateV2Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticateAuthenticateV2Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticateAuthenticateV2Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticateAuthenticateV2Response(val *AuthenticateAuthenticateV2Response) *NullableAuthenticateAuthenticateV2Response {
	return &NullableAuthenticateAuthenticateV2Response{value: val, isSet: true}
}

func (v NullableAuthenticateAuthenticateV2Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticateAuthenticateV2Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


