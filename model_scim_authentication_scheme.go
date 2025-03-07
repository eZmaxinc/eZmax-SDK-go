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

// checks if the ScimAuthenticationScheme type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScimAuthenticationScheme{}

// ScimAuthenticationScheme struct for ScimAuthenticationScheme
type ScimAuthenticationScheme struct {
	// A description of the authentication scheme.
	Description string `json:"description"`
	// The common authentication scheme name
	Name string `json:"name"`
	// The authentication scheme.
	Type string `json:"type"`
}

type _ScimAuthenticationScheme ScimAuthenticationScheme

// NewScimAuthenticationScheme instantiates a new ScimAuthenticationScheme object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScimAuthenticationScheme(description string, name string, type_ string) *ScimAuthenticationScheme {
	this := ScimAuthenticationScheme{}
	this.Description = description
	this.Name = name
	this.Type = type_
	return &this
}

// NewScimAuthenticationSchemeWithDefaults instantiates a new ScimAuthenticationScheme object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScimAuthenticationSchemeWithDefaults() *ScimAuthenticationScheme {
	this := ScimAuthenticationScheme{}
	return &this
}

// GetDescription returns the Description field value
func (o *ScimAuthenticationScheme) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ScimAuthenticationScheme) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ScimAuthenticationScheme) SetDescription(v string) {
	o.Description = v
}

// GetName returns the Name field value
func (o *ScimAuthenticationScheme) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ScimAuthenticationScheme) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ScimAuthenticationScheme) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *ScimAuthenticationScheme) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ScimAuthenticationScheme) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ScimAuthenticationScheme) SetType(v string) {
	o.Type = v
}

func (o ScimAuthenticationScheme) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScimAuthenticationScheme) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["description"] = o.Description
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *ScimAuthenticationScheme) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"description",
		"name",
		"type",
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

	varScimAuthenticationScheme := _ScimAuthenticationScheme{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varScimAuthenticationScheme)

	if err != nil {
		return err
	}

	*o = ScimAuthenticationScheme(varScimAuthenticationScheme)

	return err
}

type NullableScimAuthenticationScheme struct {
	value *ScimAuthenticationScheme
	isSet bool
}

func (v NullableScimAuthenticationScheme) Get() *ScimAuthenticationScheme {
	return v.value
}

func (v *NullableScimAuthenticationScheme) Set(val *ScimAuthenticationScheme) {
	v.value = val
	v.isSet = true
}

func (v NullableScimAuthenticationScheme) IsSet() bool {
	return v.isSet
}

func (v *NullableScimAuthenticationScheme) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScimAuthenticationScheme(val *ScimAuthenticationScheme) *NullableScimAuthenticationScheme {
	return &NullableScimAuthenticationScheme{value: val, isSet: true}
}

func (v NullableScimAuthenticationScheme) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScimAuthenticationScheme) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


