/*
eZmax API Definition (Full)

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.2.1
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CustomUserNameResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomUserNameResponse{}

// CustomUserNameResponse A User name Object
type CustomUserNameResponse struct {
	// The last name of the user
	SUserLastname string `json:"sUserLastname"`
	// The first name of the user
	SUserFirstname string `json:"sUserFirstname"`
}

type _CustomUserNameResponse CustomUserNameResponse

// NewCustomUserNameResponse instantiates a new CustomUserNameResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomUserNameResponse(sUserLastname string, sUserFirstname string) *CustomUserNameResponse {
	this := CustomUserNameResponse{}
	this.SUserLastname = sUserLastname
	this.SUserFirstname = sUserFirstname
	return &this
}

// NewCustomUserNameResponseWithDefaults instantiates a new CustomUserNameResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomUserNameResponseWithDefaults() *CustomUserNameResponse {
	this := CustomUserNameResponse{}
	return &this
}

// GetSUserLastname returns the SUserLastname field value
func (o *CustomUserNameResponse) GetSUserLastname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SUserLastname
}

// GetSUserLastnameOk returns a tuple with the SUserLastname field value
// and a boolean to check if the value has been set.
func (o *CustomUserNameResponse) GetSUserLastnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SUserLastname, true
}

// SetSUserLastname sets field value
func (o *CustomUserNameResponse) SetSUserLastname(v string) {
	o.SUserLastname = v
}

// GetSUserFirstname returns the SUserFirstname field value
func (o *CustomUserNameResponse) GetSUserFirstname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SUserFirstname
}

// GetSUserFirstnameOk returns a tuple with the SUserFirstname field value
// and a boolean to check if the value has been set.
func (o *CustomUserNameResponse) GetSUserFirstnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SUserFirstname, true
}

// SetSUserFirstname sets field value
func (o *CustomUserNameResponse) SetSUserFirstname(v string) {
	o.SUserFirstname = v
}

func (o CustomUserNameResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomUserNameResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sUserLastname"] = o.SUserLastname
	toSerialize["sUserFirstname"] = o.SUserFirstname
	return toSerialize, nil
}

func (o *CustomUserNameResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"sUserLastname",
		"sUserFirstname",
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

	varCustomUserNameResponse := _CustomUserNameResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCustomUserNameResponse)

	if err != nil {
		return err
	}

	*o = CustomUserNameResponse(varCustomUserNameResponse)

	return err
}

type NullableCustomUserNameResponse struct {
	value *CustomUserNameResponse
	isSet bool
}

func (v NullableCustomUserNameResponse) Get() *CustomUserNameResponse {
	return v.value
}

func (v *NullableCustomUserNameResponse) Set(val *CustomUserNameResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomUserNameResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomUserNameResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomUserNameResponse(val *CustomUserNameResponse) *NullableCustomUserNameResponse {
	return &NullableCustomUserNameResponse{value: val, isSet: true}
}

func (v NullableCustomUserNameResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomUserNameResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


