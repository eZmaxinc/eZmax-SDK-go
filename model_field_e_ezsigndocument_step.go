/*
 * eZmax API Definition
 *
 * This API expose all the functionnalities for the eZmax and eZsign application.  We provide SDKs for customers. They are generated using OpenAPI codegen, we encourage customers to use them as we also provide samples for them.  You can choose to build your own implementation manually or can use any compatible OpenAPI 3.0 generator like Swagger Codegen, OpenAPI codegen or any commercial generators.  If you need helping understanding how to use this API, don't waste too much time looking for it. Contact support-api@ezmax.ca, we're here to help. We are developpers so we know programmers don't like bad documentation. If you don't find what you need in the documentation, let us know, we'll improve it and put you rapidly up on track.
 *
 * API version: 1.0.28
 * Contact: support-api@ezmax.ca
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxinc/eZmax-SDK-go

import (
	"encoding/json"
	"fmt"
)

// FieldEEzsigndocumentStep The signature step of the Ezsigndocument.
type FieldEEzsigndocumentStep string

// List of Field-eEzsigndocumentStep
const (
	UNSENT FieldEEzsigndocumentStep = "Unsent"
	UNSIGNED FieldEEzsigndocumentStep = "Unsigned"
	PARTIALLY_SIGNED FieldEEzsigndocumentStep = "PartiallySigned"
	COMPLETED FieldEEzsigndocumentStep = "Completed"
)

func (v *FieldEEzsigndocumentStep) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FieldEEzsigndocumentStep(value)
	for _, existing := range []FieldEEzsigndocumentStep{ "Unsent", "Unsigned", "PartiallySigned", "Completed",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FieldEEzsigndocumentStep", value)
}

// Ptr returns reference to Field-eEzsigndocumentStep value
func (v FieldEEzsigndocumentStep) Ptr() *FieldEEzsigndocumentStep {
	return &v
}

type NullableFieldEEzsigndocumentStep struct {
	value *FieldEEzsigndocumentStep
	isSet bool
}

func (v NullableFieldEEzsigndocumentStep) Get() *FieldEEzsigndocumentStep {
	return v.value
}

func (v *NullableFieldEEzsigndocumentStep) Set(val *FieldEEzsigndocumentStep) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldEEzsigndocumentStep) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldEEzsigndocumentStep) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldEEzsigndocumentStep(val *FieldEEzsigndocumentStep) *NullableFieldEEzsigndocumentStep {
	return &NullableFieldEEzsigndocumentStep{value: val, isSet: true}
}

func (v NullableFieldEEzsigndocumentStep) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldEEzsigndocumentStep) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

