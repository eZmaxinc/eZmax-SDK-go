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

// checks if the EzsignbulksendtransmissionGetObjectV2ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignbulksendtransmissionGetObjectV2ResponseMPayload{}

// EzsignbulksendtransmissionGetObjectV2ResponseMPayload Payload for GET /2/object/ezsignbulksendtransmission/{pkiEzsignbulksendtransmissionID}
type EzsignbulksendtransmissionGetObjectV2ResponseMPayload struct {
	ObjEzsignbulksendtransmission EzsignbulksendtransmissionResponseCompound `json:"objEzsignbulksendtransmission"`
}

// NewEzsignbulksendtransmissionGetObjectV2ResponseMPayload instantiates a new EzsignbulksendtransmissionGetObjectV2ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignbulksendtransmissionGetObjectV2ResponseMPayload(objEzsignbulksendtransmission EzsignbulksendtransmissionResponseCompound) *EzsignbulksendtransmissionGetObjectV2ResponseMPayload {
	this := EzsignbulksendtransmissionGetObjectV2ResponseMPayload{}
	this.ObjEzsignbulksendtransmission = objEzsignbulksendtransmission
	return &this
}

// NewEzsignbulksendtransmissionGetObjectV2ResponseMPayloadWithDefaults instantiates a new EzsignbulksendtransmissionGetObjectV2ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignbulksendtransmissionGetObjectV2ResponseMPayloadWithDefaults() *EzsignbulksendtransmissionGetObjectV2ResponseMPayload {
	this := EzsignbulksendtransmissionGetObjectV2ResponseMPayload{}
	return &this
}

// GetObjEzsignbulksendtransmission returns the ObjEzsignbulksendtransmission field value
func (o *EzsignbulksendtransmissionGetObjectV2ResponseMPayload) GetObjEzsignbulksendtransmission() EzsignbulksendtransmissionResponseCompound {
	if o == nil {
		var ret EzsignbulksendtransmissionResponseCompound
		return ret
	}

	return o.ObjEzsignbulksendtransmission
}

// GetObjEzsignbulksendtransmissionOk returns a tuple with the ObjEzsignbulksendtransmission field value
// and a boolean to check if the value has been set.
func (o *EzsignbulksendtransmissionGetObjectV2ResponseMPayload) GetObjEzsignbulksendtransmissionOk() (*EzsignbulksendtransmissionResponseCompound, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjEzsignbulksendtransmission, true
}

// SetObjEzsignbulksendtransmission sets field value
func (o *EzsignbulksendtransmissionGetObjectV2ResponseMPayload) SetObjEzsignbulksendtransmission(v EzsignbulksendtransmissionResponseCompound) {
	o.ObjEzsignbulksendtransmission = v
}

func (o EzsignbulksendtransmissionGetObjectV2ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignbulksendtransmissionGetObjectV2ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objEzsignbulksendtransmission"] = o.ObjEzsignbulksendtransmission
	return toSerialize, nil
}

type NullableEzsignbulksendtransmissionGetObjectV2ResponseMPayload struct {
	value *EzsignbulksendtransmissionGetObjectV2ResponseMPayload
	isSet bool
}

func (v NullableEzsignbulksendtransmissionGetObjectV2ResponseMPayload) Get() *EzsignbulksendtransmissionGetObjectV2ResponseMPayload {
	return v.value
}

func (v *NullableEzsignbulksendtransmissionGetObjectV2ResponseMPayload) Set(val *EzsignbulksendtransmissionGetObjectV2ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignbulksendtransmissionGetObjectV2ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignbulksendtransmissionGetObjectV2ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignbulksendtransmissionGetObjectV2ResponseMPayload(val *EzsignbulksendtransmissionGetObjectV2ResponseMPayload) *NullableEzsignbulksendtransmissionGetObjectV2ResponseMPayload {
	return &NullableEzsignbulksendtransmissionGetObjectV2ResponseMPayload{value: val, isSet: true}
}

func (v NullableEzsignbulksendtransmissionGetObjectV2ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignbulksendtransmissionGetObjectV2ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


