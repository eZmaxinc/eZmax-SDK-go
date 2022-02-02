/*
eZmax API Definition

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.1.4
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
)

// ActivesessionResponseCompoundAllOf struct for ActivesessionResponseCompoundAllOf
type ActivesessionResponseCompoundAllOf struct {
	// An array of permissions granted to the user or api key
	APkiPermissionID []int32 `json:"a_pkiPermissionID"`
	ObjUserReal ActivesessionResponseCompoundUser `json:"objUserReal"`
	ObjUserCloned *ActivesessionResponseCompoundUser `json:"objUserCloned,omitempty"`
	ObjApikey *ActivesessionResponseCompoundApikey `json:"objApikey,omitempty"`
	// An Array of Registered modules.  These are the modules that are Licensed to be used by the User or the API Key.
	AEModuleInternalname []string `json:"a_eModuleInternalname"`
}

// NewActivesessionResponseCompoundAllOf instantiates a new ActivesessionResponseCompoundAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActivesessionResponseCompoundAllOf(aPkiPermissionID []int32, objUserReal ActivesessionResponseCompoundUser, aEModuleInternalname []string) *ActivesessionResponseCompoundAllOf {
	this := ActivesessionResponseCompoundAllOf{}
	this.APkiPermissionID = aPkiPermissionID
	this.ObjUserReal = objUserReal
	this.AEModuleInternalname = aEModuleInternalname
	return &this
}

// NewActivesessionResponseCompoundAllOfWithDefaults instantiates a new ActivesessionResponseCompoundAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActivesessionResponseCompoundAllOfWithDefaults() *ActivesessionResponseCompoundAllOf {
	this := ActivesessionResponseCompoundAllOf{}
	return &this
}

// GetAPkiPermissionID returns the APkiPermissionID field value
func (o *ActivesessionResponseCompoundAllOf) GetAPkiPermissionID() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.APkiPermissionID
}

// GetAPkiPermissionIDOk returns a tuple with the APkiPermissionID field value
// and a boolean to check if the value has been set.
func (o *ActivesessionResponseCompoundAllOf) GetAPkiPermissionIDOk() ([]int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.APkiPermissionID, true
}

// SetAPkiPermissionID sets field value
func (o *ActivesessionResponseCompoundAllOf) SetAPkiPermissionID(v []int32) {
	o.APkiPermissionID = v
}

// GetObjUserReal returns the ObjUserReal field value
func (o *ActivesessionResponseCompoundAllOf) GetObjUserReal() ActivesessionResponseCompoundUser {
	if o == nil {
		var ret ActivesessionResponseCompoundUser
		return ret
	}

	return o.ObjUserReal
}

// GetObjUserRealOk returns a tuple with the ObjUserReal field value
// and a boolean to check if the value has been set.
func (o *ActivesessionResponseCompoundAllOf) GetObjUserRealOk() (*ActivesessionResponseCompoundUser, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjUserReal, true
}

// SetObjUserReal sets field value
func (o *ActivesessionResponseCompoundAllOf) SetObjUserReal(v ActivesessionResponseCompoundUser) {
	o.ObjUserReal = v
}

// GetObjUserCloned returns the ObjUserCloned field value if set, zero value otherwise.
func (o *ActivesessionResponseCompoundAllOf) GetObjUserCloned() ActivesessionResponseCompoundUser {
	if o == nil || o.ObjUserCloned == nil {
		var ret ActivesessionResponseCompoundUser
		return ret
	}
	return *o.ObjUserCloned
}

// GetObjUserClonedOk returns a tuple with the ObjUserCloned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivesessionResponseCompoundAllOf) GetObjUserClonedOk() (*ActivesessionResponseCompoundUser, bool) {
	if o == nil || o.ObjUserCloned == nil {
		return nil, false
	}
	return o.ObjUserCloned, true
}

// HasObjUserCloned returns a boolean if a field has been set.
func (o *ActivesessionResponseCompoundAllOf) HasObjUserCloned() bool {
	if o != nil && o.ObjUserCloned != nil {
		return true
	}

	return false
}

// SetObjUserCloned gets a reference to the given ActivesessionResponseCompoundUser and assigns it to the ObjUserCloned field.
func (o *ActivesessionResponseCompoundAllOf) SetObjUserCloned(v ActivesessionResponseCompoundUser) {
	o.ObjUserCloned = &v
}

// GetObjApikey returns the ObjApikey field value if set, zero value otherwise.
func (o *ActivesessionResponseCompoundAllOf) GetObjApikey() ActivesessionResponseCompoundApikey {
	if o == nil || o.ObjApikey == nil {
		var ret ActivesessionResponseCompoundApikey
		return ret
	}
	return *o.ObjApikey
}

// GetObjApikeyOk returns a tuple with the ObjApikey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivesessionResponseCompoundAllOf) GetObjApikeyOk() (*ActivesessionResponseCompoundApikey, bool) {
	if o == nil || o.ObjApikey == nil {
		return nil, false
	}
	return o.ObjApikey, true
}

// HasObjApikey returns a boolean if a field has been set.
func (o *ActivesessionResponseCompoundAllOf) HasObjApikey() bool {
	if o != nil && o.ObjApikey != nil {
		return true
	}

	return false
}

// SetObjApikey gets a reference to the given ActivesessionResponseCompoundApikey and assigns it to the ObjApikey field.
func (o *ActivesessionResponseCompoundAllOf) SetObjApikey(v ActivesessionResponseCompoundApikey) {
	o.ObjApikey = &v
}

// GetAEModuleInternalname returns the AEModuleInternalname field value
func (o *ActivesessionResponseCompoundAllOf) GetAEModuleInternalname() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AEModuleInternalname
}

// GetAEModuleInternalnameOk returns a tuple with the AEModuleInternalname field value
// and a boolean to check if the value has been set.
func (o *ActivesessionResponseCompoundAllOf) GetAEModuleInternalnameOk() ([]string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AEModuleInternalname, true
}

// SetAEModuleInternalname sets field value
func (o *ActivesessionResponseCompoundAllOf) SetAEModuleInternalname(v []string) {
	o.AEModuleInternalname = v
}

func (o ActivesessionResponseCompoundAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["a_pkiPermissionID"] = o.APkiPermissionID
	}
	if true {
		toSerialize["objUserReal"] = o.ObjUserReal
	}
	if o.ObjUserCloned != nil {
		toSerialize["objUserCloned"] = o.ObjUserCloned
	}
	if o.ObjApikey != nil {
		toSerialize["objApikey"] = o.ObjApikey
	}
	if true {
		toSerialize["a_eModuleInternalname"] = o.AEModuleInternalname
	}
	return json.Marshal(toSerialize)
}

type NullableActivesessionResponseCompoundAllOf struct {
	value *ActivesessionResponseCompoundAllOf
	isSet bool
}

func (v NullableActivesessionResponseCompoundAllOf) Get() *ActivesessionResponseCompoundAllOf {
	return v.value
}

func (v *NullableActivesessionResponseCompoundAllOf) Set(val *ActivesessionResponseCompoundAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableActivesessionResponseCompoundAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableActivesessionResponseCompoundAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActivesessionResponseCompoundAllOf(val *ActivesessionResponseCompoundAllOf) *NullableActivesessionResponseCompoundAllOf {
	return &NullableActivesessionResponseCompoundAllOf{value: val, isSet: true}
}

func (v NullableActivesessionResponseCompoundAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActivesessionResponseCompoundAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


