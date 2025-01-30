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

// checks if the EzsignfolderImportEzsigntemplatepackageV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignfolderImportEzsigntemplatepackageV1ResponseMPayload{}

// EzsignfolderImportEzsigntemplatepackageV1ResponseMPayload Payload for POST /1/object/ezsignfolder/{pkiEzsignfolderID}/importEzsigntemplatepackage
type EzsignfolderImportEzsigntemplatepackageV1ResponseMPayload struct {
	AObjEzsigndocument []EzsigndocumentResponseCompound `json:"a_objEzsigndocument"`
}

type _EzsignfolderImportEzsigntemplatepackageV1ResponseMPayload EzsignfolderImportEzsigntemplatepackageV1ResponseMPayload

// NewEzsignfolderImportEzsigntemplatepackageV1ResponseMPayload instantiates a new EzsignfolderImportEzsigntemplatepackageV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignfolderImportEzsigntemplatepackageV1ResponseMPayload(aObjEzsigndocument []EzsigndocumentResponseCompound) *EzsignfolderImportEzsigntemplatepackageV1ResponseMPayload {
	this := EzsignfolderImportEzsigntemplatepackageV1ResponseMPayload{}
	this.AObjEzsigndocument = aObjEzsigndocument
	return &this
}

// NewEzsignfolderImportEzsigntemplatepackageV1ResponseMPayloadWithDefaults instantiates a new EzsignfolderImportEzsigntemplatepackageV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignfolderImportEzsigntemplatepackageV1ResponseMPayloadWithDefaults() *EzsignfolderImportEzsigntemplatepackageV1ResponseMPayload {
	this := EzsignfolderImportEzsigntemplatepackageV1ResponseMPayload{}
	return &this
}

// GetAObjEzsigndocument returns the AObjEzsigndocument field value
func (o *EzsignfolderImportEzsigntemplatepackageV1ResponseMPayload) GetAObjEzsigndocument() []EzsigndocumentResponseCompound {
	if o == nil {
		var ret []EzsigndocumentResponseCompound
		return ret
	}

	return o.AObjEzsigndocument
}

// GetAObjEzsigndocumentOk returns a tuple with the AObjEzsigndocument field value
// and a boolean to check if the value has been set.
func (o *EzsignfolderImportEzsigntemplatepackageV1ResponseMPayload) GetAObjEzsigndocumentOk() ([]EzsigndocumentResponseCompound, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjEzsigndocument, true
}

// SetAObjEzsigndocument sets field value
func (o *EzsignfolderImportEzsigntemplatepackageV1ResponseMPayload) SetAObjEzsigndocument(v []EzsigndocumentResponseCompound) {
	o.AObjEzsigndocument = v
}

func (o EzsignfolderImportEzsigntemplatepackageV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignfolderImportEzsigntemplatepackageV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_objEzsigndocument"] = o.AObjEzsigndocument
	return toSerialize, nil
}

func (o *EzsignfolderImportEzsigntemplatepackageV1ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"a_objEzsigndocument",
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

	varEzsignfolderImportEzsigntemplatepackageV1ResponseMPayload := _EzsignfolderImportEzsigntemplatepackageV1ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsignfolderImportEzsigntemplatepackageV1ResponseMPayload)

	if err != nil {
		return err
	}

	*o = EzsignfolderImportEzsigntemplatepackageV1ResponseMPayload(varEzsignfolderImportEzsigntemplatepackageV1ResponseMPayload)

	return err
}

type NullableEzsignfolderImportEzsigntemplatepackageV1ResponseMPayload struct {
	value *EzsignfolderImportEzsigntemplatepackageV1ResponseMPayload
	isSet bool
}

func (v NullableEzsignfolderImportEzsigntemplatepackageV1ResponseMPayload) Get() *EzsignfolderImportEzsigntemplatepackageV1ResponseMPayload {
	return v.value
}

func (v *NullableEzsignfolderImportEzsigntemplatepackageV1ResponseMPayload) Set(val *EzsignfolderImportEzsigntemplatepackageV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignfolderImportEzsigntemplatepackageV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignfolderImportEzsigntemplatepackageV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignfolderImportEzsigntemplatepackageV1ResponseMPayload(val *EzsignfolderImportEzsigntemplatepackageV1ResponseMPayload) *NullableEzsignfolderImportEzsigntemplatepackageV1ResponseMPayload {
	return &NullableEzsignfolderImportEzsigntemplatepackageV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableEzsignfolderImportEzsigntemplatepackageV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignfolderImportEzsigntemplatepackageV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


