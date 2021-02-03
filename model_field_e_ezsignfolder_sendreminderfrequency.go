/*
 * eZmax API Definition
 *
 * This API expose all the functionnalities for the eZmax and eZsign application.  We provide SDKs for customers. They are generated using OpenAPI codegen, we encourage customers to use them as we also provide samples for them.  You can choose to build your own implementation manually or can use any compatible OpenAPI 3.0 generator like Swagger Codegen, OpenAPI codegen or any commercial generators.  If you need helping understanding how to use this API, don't waste too much time looking for it. Contact support-api@ezmax.ca, we're here to help. We are developpers so we know programmers don't like bad documentation. If you don't find what you need in the documentation, let us know, we'll improve it and put you rapidly up on track.
 *
 * API version: 1.0.27
 * Contact: support-api@ezmax.ca
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxinc/eZmax-SDK-go

import (
	"encoding/json"
	"fmt"
)

// FieldEEzsignfolderSendreminderfrequency Frequency at which reminders will be sent to signers that haven't signed the documents
type FieldEEzsignfolderSendreminderfrequency string

// List of Field-eEzsignfolderSendreminderfrequency
const (
	NONE FieldEEzsignfolderSendreminderfrequency = "None"
	DAILY FieldEEzsignfolderSendreminderfrequency = "Daily"
	WEEKLY FieldEEzsignfolderSendreminderfrequency = "Weekly"
)

func (v *FieldEEzsignfolderSendreminderfrequency) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FieldEEzsignfolderSendreminderfrequency(value)
	for _, existing := range []FieldEEzsignfolderSendreminderfrequency{ "None", "Daily", "Weekly",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FieldEEzsignfolderSendreminderfrequency", value)
}

// Ptr returns reference to Field-eEzsignfolderSendreminderfrequency value
func (v FieldEEzsignfolderSendreminderfrequency) Ptr() *FieldEEzsignfolderSendreminderfrequency {
	return &v
}

type NullableFieldEEzsignfolderSendreminderfrequency struct {
	value *FieldEEzsignfolderSendreminderfrequency
	isSet bool
}

func (v NullableFieldEEzsignfolderSendreminderfrequency) Get() *FieldEEzsignfolderSendreminderfrequency {
	return v.value
}

func (v *NullableFieldEEzsignfolderSendreminderfrequency) Set(val *FieldEEzsignfolderSendreminderfrequency) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldEEzsignfolderSendreminderfrequency) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldEEzsignfolderSendreminderfrequency) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldEEzsignfolderSendreminderfrequency(val *FieldEEzsignfolderSendreminderfrequency) *NullableFieldEEzsignfolderSendreminderfrequency {
	return &NullableFieldEEzsignfolderSendreminderfrequency{value: val, isSet: true}
}

func (v NullableFieldEEzsignfolderSendreminderfrequency) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldEEzsignfolderSendreminderfrequency) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

