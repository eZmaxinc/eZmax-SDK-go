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

// checks if the EzsigntemplatepublicCreateEzsignfolderV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigntemplatepublicCreateEzsignfolderV1ResponseMPayload{}

// EzsigntemplatepublicCreateEzsignfolderV1ResponseMPayload Payload for POST /1/object/ezsigntemplatepublic/createEzsignfolder
type EzsigntemplatepublicCreateEzsignfolderV1ResponseMPayload struct {
	// The url to sign the Ezsignfolder created by the Ezsigntemplatepublic. Only used when fkiUserLogintypeID is **No validation** or **Sms only**
	SEzsigntemplatepublicSigningurl *string `json:"sEzsigntemplatepublicSigningurl,omitempty" validate:"regexp=^(https|http):\\/\\/[^\\\\s\\/$.?#].[^\\\\s]*$"`
}

// NewEzsigntemplatepublicCreateEzsignfolderV1ResponseMPayload instantiates a new EzsigntemplatepublicCreateEzsignfolderV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigntemplatepublicCreateEzsignfolderV1ResponseMPayload() *EzsigntemplatepublicCreateEzsignfolderV1ResponseMPayload {
	this := EzsigntemplatepublicCreateEzsignfolderV1ResponseMPayload{}
	return &this
}

// NewEzsigntemplatepublicCreateEzsignfolderV1ResponseMPayloadWithDefaults instantiates a new EzsigntemplatepublicCreateEzsignfolderV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigntemplatepublicCreateEzsignfolderV1ResponseMPayloadWithDefaults() *EzsigntemplatepublicCreateEzsignfolderV1ResponseMPayload {
	this := EzsigntemplatepublicCreateEzsignfolderV1ResponseMPayload{}
	return &this
}

// GetSEzsigntemplatepublicSigningurl returns the SEzsigntemplatepublicSigningurl field value if set, zero value otherwise.
func (o *EzsigntemplatepublicCreateEzsignfolderV1ResponseMPayload) GetSEzsigntemplatepublicSigningurl() string {
	if o == nil || IsNil(o.SEzsigntemplatepublicSigningurl) {
		var ret string
		return ret
	}
	return *o.SEzsigntemplatepublicSigningurl
}

// GetSEzsigntemplatepublicSigningurlOk returns a tuple with the SEzsigntemplatepublicSigningurl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepublicCreateEzsignfolderV1ResponseMPayload) GetSEzsigntemplatepublicSigningurlOk() (*string, bool) {
	if o == nil || IsNil(o.SEzsigntemplatepublicSigningurl) {
		return nil, false
	}
	return o.SEzsigntemplatepublicSigningurl, true
}

// HasSEzsigntemplatepublicSigningurl returns a boolean if a field has been set.
func (o *EzsigntemplatepublicCreateEzsignfolderV1ResponseMPayload) HasSEzsigntemplatepublicSigningurl() bool {
	if o != nil && !IsNil(o.SEzsigntemplatepublicSigningurl) {
		return true
	}

	return false
}

// SetSEzsigntemplatepublicSigningurl gets a reference to the given string and assigns it to the SEzsigntemplatepublicSigningurl field.
func (o *EzsigntemplatepublicCreateEzsignfolderV1ResponseMPayload) SetSEzsigntemplatepublicSigningurl(v string) {
	o.SEzsigntemplatepublicSigningurl = &v
}

func (o EzsigntemplatepublicCreateEzsignfolderV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigntemplatepublicCreateEzsignfolderV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SEzsigntemplatepublicSigningurl) {
		toSerialize["sEzsigntemplatepublicSigningurl"] = o.SEzsigntemplatepublicSigningurl
	}
	return toSerialize, nil
}

type NullableEzsigntemplatepublicCreateEzsignfolderV1ResponseMPayload struct {
	value *EzsigntemplatepublicCreateEzsignfolderV1ResponseMPayload
	isSet bool
}

func (v NullableEzsigntemplatepublicCreateEzsignfolderV1ResponseMPayload) Get() *EzsigntemplatepublicCreateEzsignfolderV1ResponseMPayload {
	return v.value
}

func (v *NullableEzsigntemplatepublicCreateEzsignfolderV1ResponseMPayload) Set(val *EzsigntemplatepublicCreateEzsignfolderV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigntemplatepublicCreateEzsignfolderV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigntemplatepublicCreateEzsignfolderV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigntemplatepublicCreateEzsignfolderV1ResponseMPayload(val *EzsigntemplatepublicCreateEzsignfolderV1ResponseMPayload) *NullableEzsigntemplatepublicCreateEzsignfolderV1ResponseMPayload {
	return &NullableEzsigntemplatepublicCreateEzsignfolderV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableEzsigntemplatepublicCreateEzsignfolderV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigntemplatepublicCreateEzsignfolderV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


