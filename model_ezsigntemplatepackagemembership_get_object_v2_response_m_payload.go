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

// checks if the EzsigntemplatepackagemembershipGetObjectV2ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigntemplatepackagemembershipGetObjectV2ResponseMPayload{}

// EzsigntemplatepackagemembershipGetObjectV2ResponseMPayload Payload for GET /2/object/ezsigntemplatepackagemembership/{pkiEzsigntemplatepackagemembershipID}
type EzsigntemplatepackagemembershipGetObjectV2ResponseMPayload struct {
	ObjEzsigntemplatepackagemembership EzsigntemplatepackagemembershipResponseCompound `json:"objEzsigntemplatepackagemembership"`
}

type _EzsigntemplatepackagemembershipGetObjectV2ResponseMPayload EzsigntemplatepackagemembershipGetObjectV2ResponseMPayload

// NewEzsigntemplatepackagemembershipGetObjectV2ResponseMPayload instantiates a new EzsigntemplatepackagemembershipGetObjectV2ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigntemplatepackagemembershipGetObjectV2ResponseMPayload(objEzsigntemplatepackagemembership EzsigntemplatepackagemembershipResponseCompound) *EzsigntemplatepackagemembershipGetObjectV2ResponseMPayload {
	this := EzsigntemplatepackagemembershipGetObjectV2ResponseMPayload{}
	this.ObjEzsigntemplatepackagemembership = objEzsigntemplatepackagemembership
	return &this
}

// NewEzsigntemplatepackagemembershipGetObjectV2ResponseMPayloadWithDefaults instantiates a new EzsigntemplatepackagemembershipGetObjectV2ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigntemplatepackagemembershipGetObjectV2ResponseMPayloadWithDefaults() *EzsigntemplatepackagemembershipGetObjectV2ResponseMPayload {
	this := EzsigntemplatepackagemembershipGetObjectV2ResponseMPayload{}
	return &this
}

// GetObjEzsigntemplatepackagemembership returns the ObjEzsigntemplatepackagemembership field value
func (o *EzsigntemplatepackagemembershipGetObjectV2ResponseMPayload) GetObjEzsigntemplatepackagemembership() EzsigntemplatepackagemembershipResponseCompound {
	if o == nil {
		var ret EzsigntemplatepackagemembershipResponseCompound
		return ret
	}

	return o.ObjEzsigntemplatepackagemembership
}

// GetObjEzsigntemplatepackagemembershipOk returns a tuple with the ObjEzsigntemplatepackagemembership field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepackagemembershipGetObjectV2ResponseMPayload) GetObjEzsigntemplatepackagemembershipOk() (*EzsigntemplatepackagemembershipResponseCompound, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjEzsigntemplatepackagemembership, true
}

// SetObjEzsigntemplatepackagemembership sets field value
func (o *EzsigntemplatepackagemembershipGetObjectV2ResponseMPayload) SetObjEzsigntemplatepackagemembership(v EzsigntemplatepackagemembershipResponseCompound) {
	o.ObjEzsigntemplatepackagemembership = v
}

func (o EzsigntemplatepackagemembershipGetObjectV2ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigntemplatepackagemembershipGetObjectV2ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objEzsigntemplatepackagemembership"] = o.ObjEzsigntemplatepackagemembership
	return toSerialize, nil
}

func (o *EzsigntemplatepackagemembershipGetObjectV2ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"objEzsigntemplatepackagemembership",
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

	varEzsigntemplatepackagemembershipGetObjectV2ResponseMPayload := _EzsigntemplatepackagemembershipGetObjectV2ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsigntemplatepackagemembershipGetObjectV2ResponseMPayload)

	if err != nil {
		return err
	}

	*o = EzsigntemplatepackagemembershipGetObjectV2ResponseMPayload(varEzsigntemplatepackagemembershipGetObjectV2ResponseMPayload)

	return err
}

type NullableEzsigntemplatepackagemembershipGetObjectV2ResponseMPayload struct {
	value *EzsigntemplatepackagemembershipGetObjectV2ResponseMPayload
	isSet bool
}

func (v NullableEzsigntemplatepackagemembershipGetObjectV2ResponseMPayload) Get() *EzsigntemplatepackagemembershipGetObjectV2ResponseMPayload {
	return v.value
}

func (v *NullableEzsigntemplatepackagemembershipGetObjectV2ResponseMPayload) Set(val *EzsigntemplatepackagemembershipGetObjectV2ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigntemplatepackagemembershipGetObjectV2ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigntemplatepackagemembershipGetObjectV2ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigntemplatepackagemembershipGetObjectV2ResponseMPayload(val *EzsigntemplatepackagemembershipGetObjectV2ResponseMPayload) *NullableEzsigntemplatepackagemembershipGetObjectV2ResponseMPayload {
	return &NullableEzsigntemplatepackagemembershipGetObjectV2ResponseMPayload{value: val, isSet: true}
}

func (v NullableEzsigntemplatepackagemembershipGetObjectV2ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigntemplatepackagemembershipGetObjectV2ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


