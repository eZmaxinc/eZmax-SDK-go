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

// checks if the ScimGroupMember type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScimGroupMember{}

// ScimGroupMember struct for ScimGroupMember
type ScimGroupMember struct {
	Value *string `json:"value,omitempty"`
	Display *string `json:"display,omitempty"`
	Type *string `json:"type,omitempty"`
	Ref *string `json:"$ref,omitempty"`
}

// NewScimGroupMember instantiates a new ScimGroupMember object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScimGroupMember() *ScimGroupMember {
	this := ScimGroupMember{}
	return &this
}

// NewScimGroupMemberWithDefaults instantiates a new ScimGroupMember object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScimGroupMemberWithDefaults() *ScimGroupMember {
	this := ScimGroupMember{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ScimGroupMember) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimGroupMember) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ScimGroupMember) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ScimGroupMember) SetValue(v string) {
	o.Value = &v
}

// GetDisplay returns the Display field value if set, zero value otherwise.
func (o *ScimGroupMember) GetDisplay() string {
	if o == nil || IsNil(o.Display) {
		var ret string
		return ret
	}
	return *o.Display
}

// GetDisplayOk returns a tuple with the Display field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimGroupMember) GetDisplayOk() (*string, bool) {
	if o == nil || IsNil(o.Display) {
		return nil, false
	}
	return o.Display, true
}

// HasDisplay returns a boolean if a field has been set.
func (o *ScimGroupMember) HasDisplay() bool {
	if o != nil && !IsNil(o.Display) {
		return true
	}

	return false
}

// SetDisplay gets a reference to the given string and assigns it to the Display field.
func (o *ScimGroupMember) SetDisplay(v string) {
	o.Display = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ScimGroupMember) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimGroupMember) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ScimGroupMember) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ScimGroupMember) SetType(v string) {
	o.Type = &v
}

// GetRef returns the Ref field value if set, zero value otherwise.
func (o *ScimGroupMember) GetRef() string {
	if o == nil || IsNil(o.Ref) {
		var ret string
		return ret
	}
	return *o.Ref
}

// GetRefOk returns a tuple with the Ref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimGroupMember) GetRefOk() (*string, bool) {
	if o == nil || IsNil(o.Ref) {
		return nil, false
	}
	return o.Ref, true
}

// HasRef returns a boolean if a field has been set.
func (o *ScimGroupMember) HasRef() bool {
	if o != nil && !IsNil(o.Ref) {
		return true
	}

	return false
}

// SetRef gets a reference to the given string and assigns it to the Ref field.
func (o *ScimGroupMember) SetRef(v string) {
	o.Ref = &v
}

func (o ScimGroupMember) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScimGroupMember) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Display) {
		toSerialize["display"] = o.Display
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Ref) {
		toSerialize["$ref"] = o.Ref
	}
	return toSerialize, nil
}

type NullableScimGroupMember struct {
	value *ScimGroupMember
	isSet bool
}

func (v NullableScimGroupMember) Get() *ScimGroupMember {
	return v.value
}

func (v *NullableScimGroupMember) Set(val *ScimGroupMember) {
	v.value = val
	v.isSet = true
}

func (v NullableScimGroupMember) IsSet() bool {
	return v.isSet
}

func (v *NullableScimGroupMember) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScimGroupMember(val *ScimGroupMember) *NullableScimGroupMember {
	return &NullableScimGroupMember{value: val, isSet: true}
}

func (v NullableScimGroupMember) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScimGroupMember) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


