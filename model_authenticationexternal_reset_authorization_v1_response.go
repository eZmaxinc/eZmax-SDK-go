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

// checks if the AuthenticationexternalResetAuthorizationV1Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthenticationexternalResetAuthorizationV1Response{}

// AuthenticationexternalResetAuthorizationV1Response Response for POST /1/object/authenticationexternal/{pkiAuthenticationexternalID}/resetAuthorization
type AuthenticationexternalResetAuthorizationV1Response struct {
	ObjDebugPayload CommonResponseObjDebugPayload `json:"objDebugPayload"`
	ObjDebug *CommonResponseObjDebug `json:"objDebug,omitempty"`
}

type _AuthenticationexternalResetAuthorizationV1Response AuthenticationexternalResetAuthorizationV1Response

// NewAuthenticationexternalResetAuthorizationV1Response instantiates a new AuthenticationexternalResetAuthorizationV1Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticationexternalResetAuthorizationV1Response(objDebugPayload CommonResponseObjDebugPayload) *AuthenticationexternalResetAuthorizationV1Response {
	this := AuthenticationexternalResetAuthorizationV1Response{}
	this.ObjDebugPayload = objDebugPayload
	return &this
}

// NewAuthenticationexternalResetAuthorizationV1ResponseWithDefaults instantiates a new AuthenticationexternalResetAuthorizationV1Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticationexternalResetAuthorizationV1ResponseWithDefaults() *AuthenticationexternalResetAuthorizationV1Response {
	this := AuthenticationexternalResetAuthorizationV1Response{}
	return &this
}

// GetObjDebugPayload returns the ObjDebugPayload field value
func (o *AuthenticationexternalResetAuthorizationV1Response) GetObjDebugPayload() CommonResponseObjDebugPayload {
	if o == nil {
		var ret CommonResponseObjDebugPayload
		return ret
	}

	return o.ObjDebugPayload
}

// GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field value
// and a boolean to check if the value has been set.
func (o *AuthenticationexternalResetAuthorizationV1Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjDebugPayload, true
}

// SetObjDebugPayload sets field value
func (o *AuthenticationexternalResetAuthorizationV1Response) SetObjDebugPayload(v CommonResponseObjDebugPayload) {
	o.ObjDebugPayload = v
}

// GetObjDebug returns the ObjDebug field value if set, zero value otherwise.
func (o *AuthenticationexternalResetAuthorizationV1Response) GetObjDebug() CommonResponseObjDebug {
	if o == nil || IsNil(o.ObjDebug) {
		var ret CommonResponseObjDebug
		return ret
	}
	return *o.ObjDebug
}

// GetObjDebugOk returns a tuple with the ObjDebug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationexternalResetAuthorizationV1Response) GetObjDebugOk() (*CommonResponseObjDebug, bool) {
	if o == nil || IsNil(o.ObjDebug) {
		return nil, false
	}
	return o.ObjDebug, true
}

// HasObjDebug returns a boolean if a field has been set.
func (o *AuthenticationexternalResetAuthorizationV1Response) HasObjDebug() bool {
	if o != nil && !IsNil(o.ObjDebug) {
		return true
	}

	return false
}

// SetObjDebug gets a reference to the given CommonResponseObjDebug and assigns it to the ObjDebug field.
func (o *AuthenticationexternalResetAuthorizationV1Response) SetObjDebug(v CommonResponseObjDebug) {
	o.ObjDebug = &v
}

func (o AuthenticationexternalResetAuthorizationV1Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthenticationexternalResetAuthorizationV1Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objDebugPayload"] = o.ObjDebugPayload
	if !IsNil(o.ObjDebug) {
		toSerialize["objDebug"] = o.ObjDebug
	}
	return toSerialize, nil
}

func (o *AuthenticationexternalResetAuthorizationV1Response) UnmarshalJSON(data []byte) (err error) {
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

	varAuthenticationexternalResetAuthorizationV1Response := _AuthenticationexternalResetAuthorizationV1Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAuthenticationexternalResetAuthorizationV1Response)

	if err != nil {
		return err
	}

	*o = AuthenticationexternalResetAuthorizationV1Response(varAuthenticationexternalResetAuthorizationV1Response)

	return err
}

type NullableAuthenticationexternalResetAuthorizationV1Response struct {
	value *AuthenticationexternalResetAuthorizationV1Response
	isSet bool
}

func (v NullableAuthenticationexternalResetAuthorizationV1Response) Get() *AuthenticationexternalResetAuthorizationV1Response {
	return v.value
}

func (v *NullableAuthenticationexternalResetAuthorizationV1Response) Set(val *AuthenticationexternalResetAuthorizationV1Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticationexternalResetAuthorizationV1Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticationexternalResetAuthorizationV1Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticationexternalResetAuthorizationV1Response(val *AuthenticationexternalResetAuthorizationV1Response) *NullableAuthenticationexternalResetAuthorizationV1Response {
	return &NullableAuthenticationexternalResetAuthorizationV1Response{value: val, isSet: true}
}

func (v NullableAuthenticationexternalResetAuthorizationV1Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticationexternalResetAuthorizationV1Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


