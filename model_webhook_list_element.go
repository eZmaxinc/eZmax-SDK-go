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

// checks if the WebhookListElement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookListElement{}

// WebhookListElement A Webhook List Element
type WebhookListElement struct {
	// The unique ID of the Webhook
	PkiWebhookID int32 `json:"pkiWebhookID"`
	// The description of the Webhook
	SWebhookDescription string `json:"sWebhookDescription"`
	// The URL of the Webhook callback
	SWebhookUrl string `json:"sWebhookUrl"`
	// The concatenated string to describe the Webhook event
	SWebhookEvent string `json:"sWebhookEvent"`
	// The email that will receive the Webhook in case all attempts fail
	SWebhookEmailfailed string `json:"sWebhookEmailfailed"`
	EWebhookModule FieldEWebhookModule `json:"eWebhookModule"`
	EWebhookEzsignevent *FieldEWebhookEzsignevent `json:"eWebhookEzsignevent,omitempty"`
	EWebhookManagementevent *FieldEWebhookManagementevent `json:"eWebhookManagementevent,omitempty"`
	// Whether the Webhook is active or not
	BWebhookIsactive bool `json:"bWebhookIsactive"`
	// Whether the requests will be signed or not
	BWebhookIssigned bool `json:"bWebhookIssigned"`
}

type _WebhookListElement WebhookListElement

// NewWebhookListElement instantiates a new WebhookListElement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookListElement(pkiWebhookID int32, sWebhookDescription string, sWebhookUrl string, sWebhookEvent string, sWebhookEmailfailed string, eWebhookModule FieldEWebhookModule, bWebhookIsactive bool, bWebhookIssigned bool) *WebhookListElement {
	this := WebhookListElement{}
	this.PkiWebhookID = pkiWebhookID
	this.SWebhookDescription = sWebhookDescription
	this.SWebhookUrl = sWebhookUrl
	this.SWebhookEvent = sWebhookEvent
	this.SWebhookEmailfailed = sWebhookEmailfailed
	this.EWebhookModule = eWebhookModule
	this.BWebhookIsactive = bWebhookIsactive
	this.BWebhookIssigned = bWebhookIssigned
	return &this
}

// NewWebhookListElementWithDefaults instantiates a new WebhookListElement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookListElementWithDefaults() *WebhookListElement {
	this := WebhookListElement{}
	return &this
}

// GetPkiWebhookID returns the PkiWebhookID field value
func (o *WebhookListElement) GetPkiWebhookID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiWebhookID
}

// GetPkiWebhookIDOk returns a tuple with the PkiWebhookID field value
// and a boolean to check if the value has been set.
func (o *WebhookListElement) GetPkiWebhookIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiWebhookID, true
}

// SetPkiWebhookID sets field value
func (o *WebhookListElement) SetPkiWebhookID(v int32) {
	o.PkiWebhookID = v
}

// GetSWebhookDescription returns the SWebhookDescription field value
func (o *WebhookListElement) GetSWebhookDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SWebhookDescription
}

// GetSWebhookDescriptionOk returns a tuple with the SWebhookDescription field value
// and a boolean to check if the value has been set.
func (o *WebhookListElement) GetSWebhookDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SWebhookDescription, true
}

// SetSWebhookDescription sets field value
func (o *WebhookListElement) SetSWebhookDescription(v string) {
	o.SWebhookDescription = v
}

// GetSWebhookUrl returns the SWebhookUrl field value
func (o *WebhookListElement) GetSWebhookUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SWebhookUrl
}

// GetSWebhookUrlOk returns a tuple with the SWebhookUrl field value
// and a boolean to check if the value has been set.
func (o *WebhookListElement) GetSWebhookUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SWebhookUrl, true
}

// SetSWebhookUrl sets field value
func (o *WebhookListElement) SetSWebhookUrl(v string) {
	o.SWebhookUrl = v
}

// GetSWebhookEvent returns the SWebhookEvent field value
func (o *WebhookListElement) GetSWebhookEvent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SWebhookEvent
}

// GetSWebhookEventOk returns a tuple with the SWebhookEvent field value
// and a boolean to check if the value has been set.
func (o *WebhookListElement) GetSWebhookEventOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SWebhookEvent, true
}

// SetSWebhookEvent sets field value
func (o *WebhookListElement) SetSWebhookEvent(v string) {
	o.SWebhookEvent = v
}

// GetSWebhookEmailfailed returns the SWebhookEmailfailed field value
func (o *WebhookListElement) GetSWebhookEmailfailed() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SWebhookEmailfailed
}

// GetSWebhookEmailfailedOk returns a tuple with the SWebhookEmailfailed field value
// and a boolean to check if the value has been set.
func (o *WebhookListElement) GetSWebhookEmailfailedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SWebhookEmailfailed, true
}

// SetSWebhookEmailfailed sets field value
func (o *WebhookListElement) SetSWebhookEmailfailed(v string) {
	o.SWebhookEmailfailed = v
}

// GetEWebhookModule returns the EWebhookModule field value
func (o *WebhookListElement) GetEWebhookModule() FieldEWebhookModule {
	if o == nil {
		var ret FieldEWebhookModule
		return ret
	}

	return o.EWebhookModule
}

// GetEWebhookModuleOk returns a tuple with the EWebhookModule field value
// and a boolean to check if the value has been set.
func (o *WebhookListElement) GetEWebhookModuleOk() (*FieldEWebhookModule, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EWebhookModule, true
}

// SetEWebhookModule sets field value
func (o *WebhookListElement) SetEWebhookModule(v FieldEWebhookModule) {
	o.EWebhookModule = v
}

// GetEWebhookEzsignevent returns the EWebhookEzsignevent field value if set, zero value otherwise.
func (o *WebhookListElement) GetEWebhookEzsignevent() FieldEWebhookEzsignevent {
	if o == nil || IsNil(o.EWebhookEzsignevent) {
		var ret FieldEWebhookEzsignevent
		return ret
	}
	return *o.EWebhookEzsignevent
}

// GetEWebhookEzsigneventOk returns a tuple with the EWebhookEzsignevent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookListElement) GetEWebhookEzsigneventOk() (*FieldEWebhookEzsignevent, bool) {
	if o == nil || IsNil(o.EWebhookEzsignevent) {
		return nil, false
	}
	return o.EWebhookEzsignevent, true
}

// HasEWebhookEzsignevent returns a boolean if a field has been set.
func (o *WebhookListElement) HasEWebhookEzsignevent() bool {
	if o != nil && !IsNil(o.EWebhookEzsignevent) {
		return true
	}

	return false
}

// SetEWebhookEzsignevent gets a reference to the given FieldEWebhookEzsignevent and assigns it to the EWebhookEzsignevent field.
func (o *WebhookListElement) SetEWebhookEzsignevent(v FieldEWebhookEzsignevent) {
	o.EWebhookEzsignevent = &v
}

// GetEWebhookManagementevent returns the EWebhookManagementevent field value if set, zero value otherwise.
func (o *WebhookListElement) GetEWebhookManagementevent() FieldEWebhookManagementevent {
	if o == nil || IsNil(o.EWebhookManagementevent) {
		var ret FieldEWebhookManagementevent
		return ret
	}
	return *o.EWebhookManagementevent
}

// GetEWebhookManagementeventOk returns a tuple with the EWebhookManagementevent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookListElement) GetEWebhookManagementeventOk() (*FieldEWebhookManagementevent, bool) {
	if o == nil || IsNil(o.EWebhookManagementevent) {
		return nil, false
	}
	return o.EWebhookManagementevent, true
}

// HasEWebhookManagementevent returns a boolean if a field has been set.
func (o *WebhookListElement) HasEWebhookManagementevent() bool {
	if o != nil && !IsNil(o.EWebhookManagementevent) {
		return true
	}

	return false
}

// SetEWebhookManagementevent gets a reference to the given FieldEWebhookManagementevent and assigns it to the EWebhookManagementevent field.
func (o *WebhookListElement) SetEWebhookManagementevent(v FieldEWebhookManagementevent) {
	o.EWebhookManagementevent = &v
}

// GetBWebhookIsactive returns the BWebhookIsactive field value
func (o *WebhookListElement) GetBWebhookIsactive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BWebhookIsactive
}

// GetBWebhookIsactiveOk returns a tuple with the BWebhookIsactive field value
// and a boolean to check if the value has been set.
func (o *WebhookListElement) GetBWebhookIsactiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BWebhookIsactive, true
}

// SetBWebhookIsactive sets field value
func (o *WebhookListElement) SetBWebhookIsactive(v bool) {
	o.BWebhookIsactive = v
}

// GetBWebhookIssigned returns the BWebhookIssigned field value
func (o *WebhookListElement) GetBWebhookIssigned() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BWebhookIssigned
}

// GetBWebhookIssignedOk returns a tuple with the BWebhookIssigned field value
// and a boolean to check if the value has been set.
func (o *WebhookListElement) GetBWebhookIssignedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BWebhookIssigned, true
}

// SetBWebhookIssigned sets field value
func (o *WebhookListElement) SetBWebhookIssigned(v bool) {
	o.BWebhookIssigned = v
}

func (o WebhookListElement) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookListElement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiWebhookID"] = o.PkiWebhookID
	toSerialize["sWebhookDescription"] = o.SWebhookDescription
	toSerialize["sWebhookUrl"] = o.SWebhookUrl
	toSerialize["sWebhookEvent"] = o.SWebhookEvent
	toSerialize["sWebhookEmailfailed"] = o.SWebhookEmailfailed
	toSerialize["eWebhookModule"] = o.EWebhookModule
	if !IsNil(o.EWebhookEzsignevent) {
		toSerialize["eWebhookEzsignevent"] = o.EWebhookEzsignevent
	}
	if !IsNil(o.EWebhookManagementevent) {
		toSerialize["eWebhookManagementevent"] = o.EWebhookManagementevent
	}
	toSerialize["bWebhookIsactive"] = o.BWebhookIsactive
	toSerialize["bWebhookIssigned"] = o.BWebhookIssigned
	return toSerialize, nil
}

func (o *WebhookListElement) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiWebhookID",
		"sWebhookDescription",
		"sWebhookUrl",
		"sWebhookEvent",
		"sWebhookEmailfailed",
		"eWebhookModule",
		"bWebhookIsactive",
		"bWebhookIssigned",
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

	varWebhookListElement := _WebhookListElement{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWebhookListElement)

	if err != nil {
		return err
	}

	*o = WebhookListElement(varWebhookListElement)

	return err
}

type NullableWebhookListElement struct {
	value *WebhookListElement
	isSet bool
}

func (v NullableWebhookListElement) Get() *WebhookListElement {
	return v.value
}

func (v *NullableWebhookListElement) Set(val *WebhookListElement) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookListElement) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookListElement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookListElement(val *WebhookListElement) *NullableWebhookListElement {
	return &NullableWebhookListElement{value: val, isSet: true}
}

func (v NullableWebhookListElement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookListElement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


