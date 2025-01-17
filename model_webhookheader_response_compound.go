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

// checks if the WebhookheaderResponseCompound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookheaderResponseCompound{}

// WebhookheaderResponseCompound A Webhookheader Object
type WebhookheaderResponseCompound struct {
	WebhookheaderResponse
}

type _WebhookheaderResponseCompound WebhookheaderResponseCompound

// NewWebhookheaderResponseCompound instantiates a new WebhookheaderResponseCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookheaderResponseCompound(pkiWebhookheaderID int32, fkiWebhookID int32, sWebhookheaderName string, sWebhookheaderValue string) *WebhookheaderResponseCompound {
	this := WebhookheaderResponseCompound{}
	this.PkiWebhookheaderID = pkiWebhookheaderID
	this.FkiWebhookID = fkiWebhookID
	this.SWebhookheaderName = sWebhookheaderName
	this.SWebhookheaderValue = sWebhookheaderValue
	return &this
}

// NewWebhookheaderResponseCompoundWithDefaults instantiates a new WebhookheaderResponseCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookheaderResponseCompoundWithDefaults() *WebhookheaderResponseCompound {
	this := WebhookheaderResponseCompound{}
	return &this
}

func (o WebhookheaderResponseCompound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookheaderResponseCompound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

func (o *WebhookheaderResponseCompound) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiWebhookheaderID",
		"fkiWebhookID",
		"sWebhookheaderName",
		"sWebhookheaderValue",
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

	varWebhookheaderResponseCompound := _WebhookheaderResponseCompound{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWebhookheaderResponseCompound)

	if err != nil {
		return err
	}

	*o = WebhookheaderResponseCompound(varWebhookheaderResponseCompound)

	return err
}

type NullableWebhookheaderResponseCompound struct {
	value *WebhookheaderResponseCompound
	isSet bool
}

func (v NullableWebhookheaderResponseCompound) Get() *WebhookheaderResponseCompound {
	return v.value
}

func (v *NullableWebhookheaderResponseCompound) Set(val *WebhookheaderResponseCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookheaderResponseCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookheaderResponseCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookheaderResponseCompound(val *WebhookheaderResponseCompound) *NullableWebhookheaderResponseCompound {
	return &NullableWebhookheaderResponseCompound{value: val, isSet: true}
}

func (v NullableWebhookheaderResponseCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookheaderResponseCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


