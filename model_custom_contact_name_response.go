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
)

// checks if the CustomContactNameResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomContactNameResponse{}

// CustomContactNameResponse A Custom ContactName Object
type CustomContactNameResponse struct {
	// The First name of the contact
	SContactFirstname *string `json:"sContactFirstname,omitempty" validate:"regexp=^.{1,20}$"`
	// The Last name of the contact
	SContactLastname *string `json:"sContactLastname,omitempty" validate:"regexp=^.{1,25}$"`
	// The Company name of the contact
	SContactCompany *string `json:"sContactCompany,omitempty"`
}

// NewCustomContactNameResponse instantiates a new CustomContactNameResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomContactNameResponse() *CustomContactNameResponse {
	this := CustomContactNameResponse{}
	return &this
}

// NewCustomContactNameResponseWithDefaults instantiates a new CustomContactNameResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomContactNameResponseWithDefaults() *CustomContactNameResponse {
	this := CustomContactNameResponse{}
	return &this
}

// GetSContactFirstname returns the SContactFirstname field value if set, zero value otherwise.
func (o *CustomContactNameResponse) GetSContactFirstname() string {
	if o == nil || IsNil(o.SContactFirstname) {
		var ret string
		return ret
	}
	return *o.SContactFirstname
}

// GetSContactFirstnameOk returns a tuple with the SContactFirstname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomContactNameResponse) GetSContactFirstnameOk() (*string, bool) {
	if o == nil || IsNil(o.SContactFirstname) {
		return nil, false
	}
	return o.SContactFirstname, true
}

// HasSContactFirstname returns a boolean if a field has been set.
func (o *CustomContactNameResponse) HasSContactFirstname() bool {
	if o != nil && !IsNil(o.SContactFirstname) {
		return true
	}

	return false
}

// SetSContactFirstname gets a reference to the given string and assigns it to the SContactFirstname field.
func (o *CustomContactNameResponse) SetSContactFirstname(v string) {
	o.SContactFirstname = &v
}

// GetSContactLastname returns the SContactLastname field value if set, zero value otherwise.
func (o *CustomContactNameResponse) GetSContactLastname() string {
	if o == nil || IsNil(o.SContactLastname) {
		var ret string
		return ret
	}
	return *o.SContactLastname
}

// GetSContactLastnameOk returns a tuple with the SContactLastname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomContactNameResponse) GetSContactLastnameOk() (*string, bool) {
	if o == nil || IsNil(o.SContactLastname) {
		return nil, false
	}
	return o.SContactLastname, true
}

// HasSContactLastname returns a boolean if a field has been set.
func (o *CustomContactNameResponse) HasSContactLastname() bool {
	if o != nil && !IsNil(o.SContactLastname) {
		return true
	}

	return false
}

// SetSContactLastname gets a reference to the given string and assigns it to the SContactLastname field.
func (o *CustomContactNameResponse) SetSContactLastname(v string) {
	o.SContactLastname = &v
}

// GetSContactCompany returns the SContactCompany field value if set, zero value otherwise.
func (o *CustomContactNameResponse) GetSContactCompany() string {
	if o == nil || IsNil(o.SContactCompany) {
		var ret string
		return ret
	}
	return *o.SContactCompany
}

// GetSContactCompanyOk returns a tuple with the SContactCompany field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomContactNameResponse) GetSContactCompanyOk() (*string, bool) {
	if o == nil || IsNil(o.SContactCompany) {
		return nil, false
	}
	return o.SContactCompany, true
}

// HasSContactCompany returns a boolean if a field has been set.
func (o *CustomContactNameResponse) HasSContactCompany() bool {
	if o != nil && !IsNil(o.SContactCompany) {
		return true
	}

	return false
}

// SetSContactCompany gets a reference to the given string and assigns it to the SContactCompany field.
func (o *CustomContactNameResponse) SetSContactCompany(v string) {
	o.SContactCompany = &v
}

func (o CustomContactNameResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomContactNameResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SContactFirstname) {
		toSerialize["sContactFirstname"] = o.SContactFirstname
	}
	if !IsNil(o.SContactLastname) {
		toSerialize["sContactLastname"] = o.SContactLastname
	}
	if !IsNil(o.SContactCompany) {
		toSerialize["sContactCompany"] = o.SContactCompany
	}
	return toSerialize, nil
}

type NullableCustomContactNameResponse struct {
	value *CustomContactNameResponse
	isSet bool
}

func (v NullableCustomContactNameResponse) Get() *CustomContactNameResponse {
	return v.value
}

func (v *NullableCustomContactNameResponse) Set(val *CustomContactNameResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomContactNameResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomContactNameResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomContactNameResponse(val *CustomContactNameResponse) *NullableCustomContactNameResponse {
	return &NullableCustomContactNameResponse{value: val, isSet: true}
}

func (v NullableCustomContactNameResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomContactNameResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


