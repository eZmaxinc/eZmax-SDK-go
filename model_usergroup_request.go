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

// checks if the UsergroupRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UsergroupRequest{}

// UsergroupRequest A Usergroup Object
type UsergroupRequest struct {
	// The unique ID of the Usergroup
	PkiUsergroupID *int32 `json:"pkiUsergroupID,omitempty"`
	ObjUsergroupName MultilingualUsergroupName `json:"objUsergroupName"`
}

// NewUsergroupRequest instantiates a new UsergroupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsergroupRequest(objUsergroupName MultilingualUsergroupName) *UsergroupRequest {
	this := UsergroupRequest{}
	this.ObjUsergroupName = objUsergroupName
	return &this
}

// NewUsergroupRequestWithDefaults instantiates a new UsergroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsergroupRequestWithDefaults() *UsergroupRequest {
	this := UsergroupRequest{}
	return &this
}

// GetPkiUsergroupID returns the PkiUsergroupID field value if set, zero value otherwise.
func (o *UsergroupRequest) GetPkiUsergroupID() int32 {
	if o == nil || IsNil(o.PkiUsergroupID) {
		var ret int32
		return ret
	}
	return *o.PkiUsergroupID
}

// GetPkiUsergroupIDOk returns a tuple with the PkiUsergroupID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsergroupRequest) GetPkiUsergroupIDOk() (*int32, bool) {
	if o == nil || IsNil(o.PkiUsergroupID) {
		return nil, false
	}
	return o.PkiUsergroupID, true
}

// HasPkiUsergroupID returns a boolean if a field has been set.
func (o *UsergroupRequest) HasPkiUsergroupID() bool {
	if o != nil && !IsNil(o.PkiUsergroupID) {
		return true
	}

	return false
}

// SetPkiUsergroupID gets a reference to the given int32 and assigns it to the PkiUsergroupID field.
func (o *UsergroupRequest) SetPkiUsergroupID(v int32) {
	o.PkiUsergroupID = &v
}

// GetObjUsergroupName returns the ObjUsergroupName field value
func (o *UsergroupRequest) GetObjUsergroupName() MultilingualUsergroupName {
	if o == nil {
		var ret MultilingualUsergroupName
		return ret
	}

	return o.ObjUsergroupName
}

// GetObjUsergroupNameOk returns a tuple with the ObjUsergroupName field value
// and a boolean to check if the value has been set.
func (o *UsergroupRequest) GetObjUsergroupNameOk() (*MultilingualUsergroupName, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjUsergroupName, true
}

// SetObjUsergroupName sets field value
func (o *UsergroupRequest) SetObjUsergroupName(v MultilingualUsergroupName) {
	o.ObjUsergroupName = v
}

func (o UsergroupRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsergroupRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PkiUsergroupID) {
		toSerialize["pkiUsergroupID"] = o.PkiUsergroupID
	}
	toSerialize["objUsergroupName"] = o.ObjUsergroupName
	return toSerialize, nil
}

type NullableUsergroupRequest struct {
	value *UsergroupRequest
	isSet bool
}

func (v NullableUsergroupRequest) Get() *UsergroupRequest {
	return v.value
}

func (v *NullableUsergroupRequest) Set(val *UsergroupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUsergroupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUsergroupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsergroupRequest(val *UsergroupRequest) *NullableUsergroupRequest {
	return &NullableUsergroupRequest{value: val, isSet: true}
}

func (v NullableUsergroupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsergroupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

