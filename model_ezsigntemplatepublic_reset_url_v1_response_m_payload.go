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

// checks if the EzsigntemplatepublicResetUrlV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigntemplatepublicResetUrlV1ResponseMPayload{}

// EzsigntemplatepublicResetUrlV1ResponseMPayload Payload for POST /1/object/ezsigntemplatepublic/{pkiEzsigntemplatepublicID}/resetUrl
type EzsigntemplatepublicResetUrlV1ResponseMPayload struct {
	// The url of the Ezsigntemplatepublic  You can add these value as query parameters to prefill the corresponding role  |Parameter|Description| |-|-| |sEzsigntemplatesignerDescription|The role to fill| |sContactFirstname|The contact firstname| |sContactLastname|The contact lastname| |sEmailAddress|The contact email| |sPhoneE164|The contact phone number| |sPhoneE164Cell|The contact cell phone number|
	SEzsigntemplatepublicUrl string `json:"sEzsigntemplatepublicUrl" validate:"regexp=^(https|http):\\/\\/[^\\\\s\\/$.?#].[^\\\\s]*$"`
}

type _EzsigntemplatepublicResetUrlV1ResponseMPayload EzsigntemplatepublicResetUrlV1ResponseMPayload

// NewEzsigntemplatepublicResetUrlV1ResponseMPayload instantiates a new EzsigntemplatepublicResetUrlV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigntemplatepublicResetUrlV1ResponseMPayload(sEzsigntemplatepublicUrl string) *EzsigntemplatepublicResetUrlV1ResponseMPayload {
	this := EzsigntemplatepublicResetUrlV1ResponseMPayload{}
	this.SEzsigntemplatepublicUrl = sEzsigntemplatepublicUrl
	return &this
}

// NewEzsigntemplatepublicResetUrlV1ResponseMPayloadWithDefaults instantiates a new EzsigntemplatepublicResetUrlV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigntemplatepublicResetUrlV1ResponseMPayloadWithDefaults() *EzsigntemplatepublicResetUrlV1ResponseMPayload {
	this := EzsigntemplatepublicResetUrlV1ResponseMPayload{}
	return &this
}

// GetSEzsigntemplatepublicUrl returns the SEzsigntemplatepublicUrl field value
func (o *EzsigntemplatepublicResetUrlV1ResponseMPayload) GetSEzsigntemplatepublicUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzsigntemplatepublicUrl
}

// GetSEzsigntemplatepublicUrlOk returns a tuple with the SEzsigntemplatepublicUrl field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepublicResetUrlV1ResponseMPayload) GetSEzsigntemplatepublicUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SEzsigntemplatepublicUrl, true
}

// SetSEzsigntemplatepublicUrl sets field value
func (o *EzsigntemplatepublicResetUrlV1ResponseMPayload) SetSEzsigntemplatepublicUrl(v string) {
	o.SEzsigntemplatepublicUrl = v
}

func (o EzsigntemplatepublicResetUrlV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigntemplatepublicResetUrlV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sEzsigntemplatepublicUrl"] = o.SEzsigntemplatepublicUrl
	return toSerialize, nil
}

func (o *EzsigntemplatepublicResetUrlV1ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"sEzsigntemplatepublicUrl",
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

	varEzsigntemplatepublicResetUrlV1ResponseMPayload := _EzsigntemplatepublicResetUrlV1ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsigntemplatepublicResetUrlV1ResponseMPayload)

	if err != nil {
		return err
	}

	*o = EzsigntemplatepublicResetUrlV1ResponseMPayload(varEzsigntemplatepublicResetUrlV1ResponseMPayload)

	return err
}

type NullableEzsigntemplatepublicResetUrlV1ResponseMPayload struct {
	value *EzsigntemplatepublicResetUrlV1ResponseMPayload
	isSet bool
}

func (v NullableEzsigntemplatepublicResetUrlV1ResponseMPayload) Get() *EzsigntemplatepublicResetUrlV1ResponseMPayload {
	return v.value
}

func (v *NullableEzsigntemplatepublicResetUrlV1ResponseMPayload) Set(val *EzsigntemplatepublicResetUrlV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigntemplatepublicResetUrlV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigntemplatepublicResetUrlV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigntemplatepublicResetUrlV1ResponseMPayload(val *EzsigntemplatepublicResetUrlV1ResponseMPayload) *NullableEzsigntemplatepublicResetUrlV1ResponseMPayload {
	return &NullableEzsigntemplatepublicResetUrlV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableEzsigntemplatepublicResetUrlV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigntemplatepublicResetUrlV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


