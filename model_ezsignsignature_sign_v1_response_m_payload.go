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

// checks if the EzsignsignatureSignV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignsignatureSignV1ResponseMPayload{}

// EzsignsignatureSignV1ResponseMPayload Response for POST /1/object/ezsignsignature/{pkiEzsignsignatureID}/sign
type EzsignsignatureSignV1ResponseMPayload struct {
	// The date the Ezsignsignature was signed in folder's timezone
	DtEzsignsignatureDateInFolderTimezone string `json:"dtEzsignsignatureDateInFolderTimezone" validate:"regexp=^[0-9]{4}-(0[1-9]|1[0-2])-(0[1-9]|[1-2][0-9]|3[0-1]) ([01]?[0-9]|2[0-3]):([0-5][0-9]):([0-5][0-9])$"`
	ObjTimezone *CustomTimezoneWithCodeResponse `json:"objTimezone,omitempty"`
}

type _EzsignsignatureSignV1ResponseMPayload EzsignsignatureSignV1ResponseMPayload

// NewEzsignsignatureSignV1ResponseMPayload instantiates a new EzsignsignatureSignV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignsignatureSignV1ResponseMPayload(dtEzsignsignatureDateInFolderTimezone string) *EzsignsignatureSignV1ResponseMPayload {
	this := EzsignsignatureSignV1ResponseMPayload{}
	this.DtEzsignsignatureDateInFolderTimezone = dtEzsignsignatureDateInFolderTimezone
	return &this
}

// NewEzsignsignatureSignV1ResponseMPayloadWithDefaults instantiates a new EzsignsignatureSignV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignsignatureSignV1ResponseMPayloadWithDefaults() *EzsignsignatureSignV1ResponseMPayload {
	this := EzsignsignatureSignV1ResponseMPayload{}
	return &this
}

// GetDtEzsignsignatureDateInFolderTimezone returns the DtEzsignsignatureDateInFolderTimezone field value
func (o *EzsignsignatureSignV1ResponseMPayload) GetDtEzsignsignatureDateInFolderTimezone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DtEzsignsignatureDateInFolderTimezone
}

// GetDtEzsignsignatureDateInFolderTimezoneOk returns a tuple with the DtEzsignsignatureDateInFolderTimezone field value
// and a boolean to check if the value has been set.
func (o *EzsignsignatureSignV1ResponseMPayload) GetDtEzsignsignatureDateInFolderTimezoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DtEzsignsignatureDateInFolderTimezone, true
}

// SetDtEzsignsignatureDateInFolderTimezone sets field value
func (o *EzsignsignatureSignV1ResponseMPayload) SetDtEzsignsignatureDateInFolderTimezone(v string) {
	o.DtEzsignsignatureDateInFolderTimezone = v
}

// GetObjTimezone returns the ObjTimezone field value if set, zero value otherwise.
func (o *EzsignsignatureSignV1ResponseMPayload) GetObjTimezone() CustomTimezoneWithCodeResponse {
	if o == nil || IsNil(o.ObjTimezone) {
		var ret CustomTimezoneWithCodeResponse
		return ret
	}
	return *o.ObjTimezone
}

// GetObjTimezoneOk returns a tuple with the ObjTimezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignsignatureSignV1ResponseMPayload) GetObjTimezoneOk() (*CustomTimezoneWithCodeResponse, bool) {
	if o == nil || IsNil(o.ObjTimezone) {
		return nil, false
	}
	return o.ObjTimezone, true
}

// HasObjTimezone returns a boolean if a field has been set.
func (o *EzsignsignatureSignV1ResponseMPayload) HasObjTimezone() bool {
	if o != nil && !IsNil(o.ObjTimezone) {
		return true
	}

	return false
}

// SetObjTimezone gets a reference to the given CustomTimezoneWithCodeResponse and assigns it to the ObjTimezone field.
func (o *EzsignsignatureSignV1ResponseMPayload) SetObjTimezone(v CustomTimezoneWithCodeResponse) {
	o.ObjTimezone = &v
}

func (o EzsignsignatureSignV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignsignatureSignV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dtEzsignsignatureDateInFolderTimezone"] = o.DtEzsignsignatureDateInFolderTimezone
	if !IsNil(o.ObjTimezone) {
		toSerialize["objTimezone"] = o.ObjTimezone
	}
	return toSerialize, nil
}

func (o *EzsignsignatureSignV1ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"dtEzsignsignatureDateInFolderTimezone",
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

	varEzsignsignatureSignV1ResponseMPayload := _EzsignsignatureSignV1ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsignsignatureSignV1ResponseMPayload)

	if err != nil {
		return err
	}

	*o = EzsignsignatureSignV1ResponseMPayload(varEzsignsignatureSignV1ResponseMPayload)

	return err
}

type NullableEzsignsignatureSignV1ResponseMPayload struct {
	value *EzsignsignatureSignV1ResponseMPayload
	isSet bool
}

func (v NullableEzsignsignatureSignV1ResponseMPayload) Get() *EzsignsignatureSignV1ResponseMPayload {
	return v.value
}

func (v *NullableEzsignsignatureSignV1ResponseMPayload) Set(val *EzsignsignatureSignV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignsignatureSignV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignsignatureSignV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignsignatureSignV1ResponseMPayload(val *EzsignsignatureSignV1ResponseMPayload) *NullableEzsignsignatureSignV1ResponseMPayload {
	return &NullableEzsignsignatureSignV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableEzsignsignatureSignV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignsignatureSignV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


