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

// checks if the EzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1Response{}

// EzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1Response Response for GET /1/object/ezsignbulksendtransmission/{pkiEzsignbulksendtransmissionID}/getEzsignsignaturesAutomatic
type EzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1Response struct {
	CommonResponse
	MPayload EzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1ResponseMPayload `json:"mPayload"`
}

type _EzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1Response EzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1Response

// NewEzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1Response instantiates a new EzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1Response(mPayload EzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1ResponseMPayload, objDebugPayload CommonResponseObjDebugPayload) *EzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1Response {
	this := EzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1Response{}
	this.ObjDebugPayload = objDebugPayload
	this.MPayload = mPayload
	return &this
}

// NewEzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1ResponseWithDefaults instantiates a new EzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1ResponseWithDefaults() *EzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1Response {
	this := EzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1Response{}
	return &this
}

// GetMPayload returns the MPayload field value
func (o *EzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1Response) GetMPayload() EzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1ResponseMPayload {
	if o == nil {
		var ret EzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1ResponseMPayload
		return ret
	}

	return o.MPayload
}

// GetMPayloadOk returns a tuple with the MPayload field value
// and a boolean to check if the value has been set.
func (o *EzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1Response) GetMPayloadOk() (*EzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1ResponseMPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MPayload, true
}

// SetMPayload sets field value
func (o *EzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1Response) SetMPayload(v EzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1ResponseMPayload) {
	o.MPayload = v
}

func (o EzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mPayload"] = o.MPayload
	return toSerialize, nil
}

func (o *EzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"mPayload",
		"objDebugPayload",
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

	varEzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1Response := _EzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1Response)

	if err != nil {
		return err
	}

	*o = EzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1Response(varEzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1Response)

	return err
}

type NullableEzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1Response struct {
	value *EzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1Response
	isSet bool
}

func (v NullableEzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1Response) Get() *EzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1Response {
	return v.value
}

func (v *NullableEzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1Response) Set(val *EzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1Response) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1Response) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1Response(val *EzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1Response) *NullableEzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1Response {
	return &NullableEzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1Response{value: val, isSet: true}
}

func (v NullableEzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


