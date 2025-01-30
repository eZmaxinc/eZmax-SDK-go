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

// checks if the EzsigndocumentDeclineToSignV1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigndocumentDeclineToSignV1Request{}

// EzsigndocumentDeclineToSignV1Request Request for POST /1/object/ezsigndocument/{pkiEzsigndocumentID}/declineToSign
type EzsigndocumentDeclineToSignV1Request struct {
	// Reason for refusal
	SReason string `json:"sReason" validate:"regexp=^.{0,65535}$"`
}

type _EzsigndocumentDeclineToSignV1Request EzsigndocumentDeclineToSignV1Request

// NewEzsigndocumentDeclineToSignV1Request instantiates a new EzsigndocumentDeclineToSignV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigndocumentDeclineToSignV1Request(sReason string) *EzsigndocumentDeclineToSignV1Request {
	this := EzsigndocumentDeclineToSignV1Request{}
	this.SReason = sReason
	return &this
}

// NewEzsigndocumentDeclineToSignV1RequestWithDefaults instantiates a new EzsigndocumentDeclineToSignV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigndocumentDeclineToSignV1RequestWithDefaults() *EzsigndocumentDeclineToSignV1Request {
	this := EzsigndocumentDeclineToSignV1Request{}
	return &this
}

// GetSReason returns the SReason field value
func (o *EzsigndocumentDeclineToSignV1Request) GetSReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SReason
}

// GetSReasonOk returns a tuple with the SReason field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentDeclineToSignV1Request) GetSReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SReason, true
}

// SetSReason sets field value
func (o *EzsigndocumentDeclineToSignV1Request) SetSReason(v string) {
	o.SReason = v
}

func (o EzsigndocumentDeclineToSignV1Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigndocumentDeclineToSignV1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sReason"] = o.SReason
	return toSerialize, nil
}

func (o *EzsigndocumentDeclineToSignV1Request) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"sReason",
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

	varEzsigndocumentDeclineToSignV1Request := _EzsigndocumentDeclineToSignV1Request{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsigndocumentDeclineToSignV1Request)

	if err != nil {
		return err
	}

	*o = EzsigndocumentDeclineToSignV1Request(varEzsigndocumentDeclineToSignV1Request)

	return err
}

type NullableEzsigndocumentDeclineToSignV1Request struct {
	value *EzsigndocumentDeclineToSignV1Request
	isSet bool
}

func (v NullableEzsigndocumentDeclineToSignV1Request) Get() *EzsigndocumentDeclineToSignV1Request {
	return v.value
}

func (v *NullableEzsigndocumentDeclineToSignV1Request) Set(val *EzsigndocumentDeclineToSignV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigndocumentDeclineToSignV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigndocumentDeclineToSignV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigndocumentDeclineToSignV1Request(val *EzsigndocumentDeclineToSignV1Request) *NullableEzsigndocumentDeclineToSignV1Request {
	return &NullableEzsigndocumentDeclineToSignV1Request{value: val, isSet: true}
}

func (v NullableEzsigndocumentDeclineToSignV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigndocumentDeclineToSignV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


