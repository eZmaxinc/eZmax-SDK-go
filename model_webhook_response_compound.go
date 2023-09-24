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

// checks if the WebhookResponseCompound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookResponseCompound{}

// WebhookResponseCompound A Webhook Object
type WebhookResponseCompound struct {
	// The unique ID of the Webhook
	PkiWebhookID int32 `json:"pkiWebhookID"`
	// The description of the Webhook
	SWebhookDescription string `json:"sWebhookDescription"`
	// The unique ID of the Ezsignfoldertype.
	FkiEzsignfoldertypeID *int32 `json:"fkiEzsignfoldertypeID,omitempty"`
	// The name of the Ezsignfoldertype in the language of the requester
	SEzsignfoldertypeNameX *string `json:"sEzsignfoldertypeNameX,omitempty"`
	EWebhookModule FieldEWebhookModule `json:"eWebhookModule"`
	EWebhookEzsignevent *FieldEWebhookEzsignevent `json:"eWebhookEzsignevent,omitempty"`
	EWebhookManagementevent *FieldEWebhookManagementevent `json:"eWebhookManagementevent,omitempty"`
	// The URL of the Webhook callback
	SWebhookUrl string `json:"sWebhookUrl"`
	// The email that will receive the Webhook in case all attempts fail
	SWebhookEmailfailed string `json:"sWebhookEmailfailed"`
	// Whether the Webhook is active or not
	BWebhookIsactive *bool `json:"bWebhookIsactive,omitempty"`
	// Wheter the server's SSL certificate should be validated or not. Not recommended to skip for production use
	BWebhookSkipsslvalidation bool `json:"bWebhookSkipsslvalidation"`
	// The concatenated string to describe the Webhook event
	SWebhookEvent string `json:"sWebhookEvent"`
}

// NewWebhookResponseCompound instantiates a new WebhookResponseCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookResponseCompound(pkiWebhookID int32, sWebhookDescription string, eWebhookModule FieldEWebhookModule, sWebhookUrl string, sWebhookEmailfailed string, bWebhookSkipsslvalidation bool, sWebhookEvent string) *WebhookResponseCompound {
	this := WebhookResponseCompound{}
	this.PkiWebhookID = pkiWebhookID
	this.SWebhookDescription = sWebhookDescription
	this.EWebhookModule = eWebhookModule
	this.SWebhookUrl = sWebhookUrl
	this.SWebhookEmailfailed = sWebhookEmailfailed
	this.BWebhookSkipsslvalidation = bWebhookSkipsslvalidation
	this.SWebhookEvent = sWebhookEvent
	return &this
}

// NewWebhookResponseCompoundWithDefaults instantiates a new WebhookResponseCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookResponseCompoundWithDefaults() *WebhookResponseCompound {
	this := WebhookResponseCompound{}
	return &this
}

// GetPkiWebhookID returns the PkiWebhookID field value
func (o *WebhookResponseCompound) GetPkiWebhookID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiWebhookID
}

// GetPkiWebhookIDOk returns a tuple with the PkiWebhookID field value
// and a boolean to check if the value has been set.
func (o *WebhookResponseCompound) GetPkiWebhookIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiWebhookID, true
}

// SetPkiWebhookID sets field value
func (o *WebhookResponseCompound) SetPkiWebhookID(v int32) {
	o.PkiWebhookID = v
}

// GetSWebhookDescription returns the SWebhookDescription field value
func (o *WebhookResponseCompound) GetSWebhookDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SWebhookDescription
}

// GetSWebhookDescriptionOk returns a tuple with the SWebhookDescription field value
// and a boolean to check if the value has been set.
func (o *WebhookResponseCompound) GetSWebhookDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SWebhookDescription, true
}

// SetSWebhookDescription sets field value
func (o *WebhookResponseCompound) SetSWebhookDescription(v string) {
	o.SWebhookDescription = v
}

// GetFkiEzsignfoldertypeID returns the FkiEzsignfoldertypeID field value if set, zero value otherwise.
func (o *WebhookResponseCompound) GetFkiEzsignfoldertypeID() int32 {
	if o == nil || IsNil(o.FkiEzsignfoldertypeID) {
		var ret int32
		return ret
	}
	return *o.FkiEzsignfoldertypeID
}

// GetFkiEzsignfoldertypeIDOk returns a tuple with the FkiEzsignfoldertypeID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookResponseCompound) GetFkiEzsignfoldertypeIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiEzsignfoldertypeID) {
		return nil, false
	}
	return o.FkiEzsignfoldertypeID, true
}

// HasFkiEzsignfoldertypeID returns a boolean if a field has been set.
func (o *WebhookResponseCompound) HasFkiEzsignfoldertypeID() bool {
	if o != nil && !IsNil(o.FkiEzsignfoldertypeID) {
		return true
	}

	return false
}

// SetFkiEzsignfoldertypeID gets a reference to the given int32 and assigns it to the FkiEzsignfoldertypeID field.
func (o *WebhookResponseCompound) SetFkiEzsignfoldertypeID(v int32) {
	o.FkiEzsignfoldertypeID = &v
}

// GetSEzsignfoldertypeNameX returns the SEzsignfoldertypeNameX field value if set, zero value otherwise.
func (o *WebhookResponseCompound) GetSEzsignfoldertypeNameX() string {
	if o == nil || IsNil(o.SEzsignfoldertypeNameX) {
		var ret string
		return ret
	}
	return *o.SEzsignfoldertypeNameX
}

// GetSEzsignfoldertypeNameXOk returns a tuple with the SEzsignfoldertypeNameX field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookResponseCompound) GetSEzsignfoldertypeNameXOk() (*string, bool) {
	if o == nil || IsNil(o.SEzsignfoldertypeNameX) {
		return nil, false
	}
	return o.SEzsignfoldertypeNameX, true
}

// HasSEzsignfoldertypeNameX returns a boolean if a field has been set.
func (o *WebhookResponseCompound) HasSEzsignfoldertypeNameX() bool {
	if o != nil && !IsNil(o.SEzsignfoldertypeNameX) {
		return true
	}

	return false
}

// SetSEzsignfoldertypeNameX gets a reference to the given string and assigns it to the SEzsignfoldertypeNameX field.
func (o *WebhookResponseCompound) SetSEzsignfoldertypeNameX(v string) {
	o.SEzsignfoldertypeNameX = &v
}

// GetEWebhookModule returns the EWebhookModule field value
func (o *WebhookResponseCompound) GetEWebhookModule() FieldEWebhookModule {
	if o == nil {
		var ret FieldEWebhookModule
		return ret
	}

	return o.EWebhookModule
}

// GetEWebhookModuleOk returns a tuple with the EWebhookModule field value
// and a boolean to check if the value has been set.
func (o *WebhookResponseCompound) GetEWebhookModuleOk() (*FieldEWebhookModule, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EWebhookModule, true
}

// SetEWebhookModule sets field value
func (o *WebhookResponseCompound) SetEWebhookModule(v FieldEWebhookModule) {
	o.EWebhookModule = v
}

// GetEWebhookEzsignevent returns the EWebhookEzsignevent field value if set, zero value otherwise.
func (o *WebhookResponseCompound) GetEWebhookEzsignevent() FieldEWebhookEzsignevent {
	if o == nil || IsNil(o.EWebhookEzsignevent) {
		var ret FieldEWebhookEzsignevent
		return ret
	}
	return *o.EWebhookEzsignevent
}

// GetEWebhookEzsigneventOk returns a tuple with the EWebhookEzsignevent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookResponseCompound) GetEWebhookEzsigneventOk() (*FieldEWebhookEzsignevent, bool) {
	if o == nil || IsNil(o.EWebhookEzsignevent) {
		return nil, false
	}
	return o.EWebhookEzsignevent, true
}

// HasEWebhookEzsignevent returns a boolean if a field has been set.
func (o *WebhookResponseCompound) HasEWebhookEzsignevent() bool {
	if o != nil && !IsNil(o.EWebhookEzsignevent) {
		return true
	}

	return false
}

// SetEWebhookEzsignevent gets a reference to the given FieldEWebhookEzsignevent and assigns it to the EWebhookEzsignevent field.
func (o *WebhookResponseCompound) SetEWebhookEzsignevent(v FieldEWebhookEzsignevent) {
	o.EWebhookEzsignevent = &v
}

// GetEWebhookManagementevent returns the EWebhookManagementevent field value if set, zero value otherwise.
func (o *WebhookResponseCompound) GetEWebhookManagementevent() FieldEWebhookManagementevent {
	if o == nil || IsNil(o.EWebhookManagementevent) {
		var ret FieldEWebhookManagementevent
		return ret
	}
	return *o.EWebhookManagementevent
}

// GetEWebhookManagementeventOk returns a tuple with the EWebhookManagementevent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookResponseCompound) GetEWebhookManagementeventOk() (*FieldEWebhookManagementevent, bool) {
	if o == nil || IsNil(o.EWebhookManagementevent) {
		return nil, false
	}
	return o.EWebhookManagementevent, true
}

// HasEWebhookManagementevent returns a boolean if a field has been set.
func (o *WebhookResponseCompound) HasEWebhookManagementevent() bool {
	if o != nil && !IsNil(o.EWebhookManagementevent) {
		return true
	}

	return false
}

// SetEWebhookManagementevent gets a reference to the given FieldEWebhookManagementevent and assigns it to the EWebhookManagementevent field.
func (o *WebhookResponseCompound) SetEWebhookManagementevent(v FieldEWebhookManagementevent) {
	o.EWebhookManagementevent = &v
}

// GetSWebhookUrl returns the SWebhookUrl field value
func (o *WebhookResponseCompound) GetSWebhookUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SWebhookUrl
}

// GetSWebhookUrlOk returns a tuple with the SWebhookUrl field value
// and a boolean to check if the value has been set.
func (o *WebhookResponseCompound) GetSWebhookUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SWebhookUrl, true
}

// SetSWebhookUrl sets field value
func (o *WebhookResponseCompound) SetSWebhookUrl(v string) {
	o.SWebhookUrl = v
}

// GetSWebhookEmailfailed returns the SWebhookEmailfailed field value
func (o *WebhookResponseCompound) GetSWebhookEmailfailed() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SWebhookEmailfailed
}

// GetSWebhookEmailfailedOk returns a tuple with the SWebhookEmailfailed field value
// and a boolean to check if the value has been set.
func (o *WebhookResponseCompound) GetSWebhookEmailfailedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SWebhookEmailfailed, true
}

// SetSWebhookEmailfailed sets field value
func (o *WebhookResponseCompound) SetSWebhookEmailfailed(v string) {
	o.SWebhookEmailfailed = v
}

// GetBWebhookIsactive returns the BWebhookIsactive field value if set, zero value otherwise.
func (o *WebhookResponseCompound) GetBWebhookIsactive() bool {
	if o == nil || IsNil(o.BWebhookIsactive) {
		var ret bool
		return ret
	}
	return *o.BWebhookIsactive
}

// GetBWebhookIsactiveOk returns a tuple with the BWebhookIsactive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookResponseCompound) GetBWebhookIsactiveOk() (*bool, bool) {
	if o == nil || IsNil(o.BWebhookIsactive) {
		return nil, false
	}
	return o.BWebhookIsactive, true
}

// HasBWebhookIsactive returns a boolean if a field has been set.
func (o *WebhookResponseCompound) HasBWebhookIsactive() bool {
	if o != nil && !IsNil(o.BWebhookIsactive) {
		return true
	}

	return false
}

// SetBWebhookIsactive gets a reference to the given bool and assigns it to the BWebhookIsactive field.
func (o *WebhookResponseCompound) SetBWebhookIsactive(v bool) {
	o.BWebhookIsactive = &v
}

// GetBWebhookSkipsslvalidation returns the BWebhookSkipsslvalidation field value
func (o *WebhookResponseCompound) GetBWebhookSkipsslvalidation() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BWebhookSkipsslvalidation
}

// GetBWebhookSkipsslvalidationOk returns a tuple with the BWebhookSkipsslvalidation field value
// and a boolean to check if the value has been set.
func (o *WebhookResponseCompound) GetBWebhookSkipsslvalidationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BWebhookSkipsslvalidation, true
}

// SetBWebhookSkipsslvalidation sets field value
func (o *WebhookResponseCompound) SetBWebhookSkipsslvalidation(v bool) {
	o.BWebhookSkipsslvalidation = v
}

// GetSWebhookEvent returns the SWebhookEvent field value
func (o *WebhookResponseCompound) GetSWebhookEvent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SWebhookEvent
}

// GetSWebhookEventOk returns a tuple with the SWebhookEvent field value
// and a boolean to check if the value has been set.
func (o *WebhookResponseCompound) GetSWebhookEventOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SWebhookEvent, true
}

// SetSWebhookEvent sets field value
func (o *WebhookResponseCompound) SetSWebhookEvent(v string) {
	o.SWebhookEvent = v
}

func (o WebhookResponseCompound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookResponseCompound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiWebhookID"] = o.PkiWebhookID
	toSerialize["sWebhookDescription"] = o.SWebhookDescription
	if !IsNil(o.FkiEzsignfoldertypeID) {
		toSerialize["fkiEzsignfoldertypeID"] = o.FkiEzsignfoldertypeID
	}
	if !IsNil(o.SEzsignfoldertypeNameX) {
		toSerialize["sEzsignfoldertypeNameX"] = o.SEzsignfoldertypeNameX
	}
	toSerialize["eWebhookModule"] = o.EWebhookModule
	if !IsNil(o.EWebhookEzsignevent) {
		toSerialize["eWebhookEzsignevent"] = o.EWebhookEzsignevent
	}
	if !IsNil(o.EWebhookManagementevent) {
		toSerialize["eWebhookManagementevent"] = o.EWebhookManagementevent
	}
	toSerialize["sWebhookUrl"] = o.SWebhookUrl
	toSerialize["sWebhookEmailfailed"] = o.SWebhookEmailfailed
	if !IsNil(o.BWebhookIsactive) {
		toSerialize["bWebhookIsactive"] = o.BWebhookIsactive
	}
	toSerialize["bWebhookSkipsslvalidation"] = o.BWebhookSkipsslvalidation
	toSerialize["sWebhookEvent"] = o.SWebhookEvent
	return toSerialize, nil
}

type NullableWebhookResponseCompound struct {
	value *WebhookResponseCompound
	isSet bool
}

func (v NullableWebhookResponseCompound) Get() *WebhookResponseCompound {
	return v.value
}

func (v *NullableWebhookResponseCompound) Set(val *WebhookResponseCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookResponseCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookResponseCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookResponseCompound(val *WebhookResponseCompound) *NullableWebhookResponseCompound {
	return &NullableWebhookResponseCompound{value: val, isSet: true}
}

func (v NullableWebhookResponseCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookResponseCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

