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
	"bytes"
	"fmt"
)

// checks if the WebhookheaderRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookheaderRequest{}

// WebhookheaderRequest A webhookheader object
type WebhookheaderRequest struct {
	// The unique ID of the Webhookheader
	PkiWebhookheaderID *int32 `json:"pkiWebhookheaderID,omitempty"`
	// The Name of the Webhookheader
	SWebhookheaderName string `json:"sWebhookheaderName"`
	// The Value of the Webhookheader
	SWebhookheaderValue string `json:"sWebhookheaderValue"`
}

type _WebhookheaderRequest WebhookheaderRequest

// NewWebhookheaderRequest instantiates a new WebhookheaderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookheaderRequest(sWebhookheaderName string, sWebhookheaderValue string) *WebhookheaderRequest {
	this := WebhookheaderRequest{}
	this.SWebhookheaderName = sWebhookheaderName
	this.SWebhookheaderValue = sWebhookheaderValue
	return &this
}

// NewWebhookheaderRequestWithDefaults instantiates a new WebhookheaderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookheaderRequestWithDefaults() *WebhookheaderRequest {
	this := WebhookheaderRequest{}
	return &this
}

// GetPkiWebhookheaderID returns the PkiWebhookheaderID field value if set, zero value otherwise.
func (o *WebhookheaderRequest) GetPkiWebhookheaderID() int32 {
	if o == nil || IsNil(o.PkiWebhookheaderID) {
		var ret int32
		return ret
	}
	return *o.PkiWebhookheaderID
}

// GetPkiWebhookheaderIDOk returns a tuple with the PkiWebhookheaderID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookheaderRequest) GetPkiWebhookheaderIDOk() (*int32, bool) {
	if o == nil || IsNil(o.PkiWebhookheaderID) {
		return nil, false
	}
	return o.PkiWebhookheaderID, true
}

// HasPkiWebhookheaderID returns a boolean if a field has been set.
func (o *WebhookheaderRequest) HasPkiWebhookheaderID() bool {
	if o != nil && !IsNil(o.PkiWebhookheaderID) {
		return true
	}

	return false
}

// SetPkiWebhookheaderID gets a reference to the given int32 and assigns it to the PkiWebhookheaderID field.
func (o *WebhookheaderRequest) SetPkiWebhookheaderID(v int32) {
	o.PkiWebhookheaderID = &v
}

// GetSWebhookheaderName returns the SWebhookheaderName field value
func (o *WebhookheaderRequest) GetSWebhookheaderName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SWebhookheaderName
}

// GetSWebhookheaderNameOk returns a tuple with the SWebhookheaderName field value
// and a boolean to check if the value has been set.
func (o *WebhookheaderRequest) GetSWebhookheaderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SWebhookheaderName, true
}

// SetSWebhookheaderName sets field value
func (o *WebhookheaderRequest) SetSWebhookheaderName(v string) {
	o.SWebhookheaderName = v
}

// GetSWebhookheaderValue returns the SWebhookheaderValue field value
func (o *WebhookheaderRequest) GetSWebhookheaderValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SWebhookheaderValue
}

// GetSWebhookheaderValueOk returns a tuple with the SWebhookheaderValue field value
// and a boolean to check if the value has been set.
func (o *WebhookheaderRequest) GetSWebhookheaderValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SWebhookheaderValue, true
}

// SetSWebhookheaderValue sets field value
func (o *WebhookheaderRequest) SetSWebhookheaderValue(v string) {
	o.SWebhookheaderValue = v
}

func (o WebhookheaderRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookheaderRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PkiWebhookheaderID) {
		toSerialize["pkiWebhookheaderID"] = o.PkiWebhookheaderID
	}
	toSerialize["sWebhookheaderName"] = o.SWebhookheaderName
	toSerialize["sWebhookheaderValue"] = o.SWebhookheaderValue
	return toSerialize, nil
}

func (o *WebhookheaderRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varWebhookheaderRequest := _WebhookheaderRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWebhookheaderRequest)

	if err != nil {
		return err
	}

	*o = WebhookheaderRequest(varWebhookheaderRequest)

	return err
}

type NullableWebhookheaderRequest struct {
	value *WebhookheaderRequest
	isSet bool
}

func (v NullableWebhookheaderRequest) Get() *WebhookheaderRequest {
	return v.value
}

func (v *NullableWebhookheaderRequest) Set(val *WebhookheaderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookheaderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookheaderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookheaderRequest(val *WebhookheaderRequest) *NullableWebhookheaderRequest {
	return &NullableWebhookheaderRequest{value: val, isSet: true}
}

func (v NullableWebhookheaderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookheaderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


