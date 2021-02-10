/*
 * eZmax API Definition
 *
 * This API expose all the functionnalities for the eZmax and eZsign application.  We provide SDKs for customers. They are generated using OpenAPI codegen, we encourage customers to use them as we also provide samples for them.  You can choose to build your own implementation manually or can use any compatible OpenAPI 3.0 generator like Swagger Codegen, OpenAPI codegen or any commercial generators.  If you need helping understanding how to use this API, don't waste too much time looking for it. Contact support-api@ezmax.ca, we're here to help. We are developpers so we know programmers don't like bad documentation. If you don't find what you need in the documentation, let us know, we'll improve it and put you rapidly up on track.
 *
 * API version: 1.0.29
 * Contact: support-api@ezmax.ca
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxinc/eZmax-SDK-go

import (
	"encoding/json"
	"fmt"
)

// FieldEEzsignfolderStep The signature step of the Ezsignfolder.
type FieldEEzsignfolderStep string

// List of Field-eEzsignfolderStep
const (
	UNSENT FieldEEzsignfolderStep = "Unsent"
	SENT FieldEEzsignfolderStep = "Sent"
	PARTIALLY_SIGNED FieldEEzsignfolderStep = "PartiallySigned"
	EXPIRED FieldEEzsignfolderStep = "Expired"
	SIGNED FieldEEzsignfolderStep = "Signed"
	COMPLETED FieldEEzsignfolderStep = "Completed"
	ARCHIVED FieldEEzsignfolderStep = "Archived"
)

func (v *FieldEEzsignfolderStep) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FieldEEzsignfolderStep(value)
	for _, existing := range []FieldEEzsignfolderStep{ "Unsent", "Sent", "PartiallySigned", "Expired", "Signed", "Completed", "Archived",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FieldEEzsignfolderStep", value)
}

// Ptr returns reference to Field-eEzsignfolderStep value
func (v FieldEEzsignfolderStep) Ptr() *FieldEEzsignfolderStep {
	return &v
}

type NullableFieldEEzsignfolderStep struct {
	value *FieldEEzsignfolderStep
	isSet bool
}

func (v NullableFieldEEzsignfolderStep) Get() *FieldEEzsignfolderStep {
	return v.value
}

func (v *NullableFieldEEzsignfolderStep) Set(val *FieldEEzsignfolderStep) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldEEzsignfolderStep) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldEEzsignfolderStep) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldEEzsignfolderStep(val *FieldEEzsignfolderStep) *NullableFieldEEzsignfolderStep {
	return &NullableFieldEEzsignfolderStep{value: val, isSet: true}
}

func (v NullableFieldEEzsignfolderStep) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldEEzsignfolderStep) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

