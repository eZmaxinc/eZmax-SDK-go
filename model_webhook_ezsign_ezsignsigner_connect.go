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

// checks if the WebhookEzsignEzsignsignerConnect type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookEzsignEzsignsignerConnect{}

// WebhookEzsignEzsignsignerConnect This is the base Webhook object
type WebhookEzsignEzsignsignerConnect struct {
	ObjWebhook CustomWebhookResponse `json:"objWebhook"`
	// An array containing details of previous attempts that were made to deliver the message. The array is empty if it's the first attempt.
	AObjAttempt []AttemptResponseCompound `json:"a_objAttempt"`
	ObjEzsignfolder *EzsignfolderResponse `json:"objEzsignfolder,omitempty"`
	ObjEzsignfoldersignerassociation EzsignfoldersignerassociationResponseCompound `json:"objEzsignfoldersignerassociation"`
}

// NewWebhookEzsignEzsignsignerConnect instantiates a new WebhookEzsignEzsignsignerConnect object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookEzsignEzsignsignerConnect(objWebhook CustomWebhookResponse, aObjAttempt []AttemptResponseCompound, objEzsignfoldersignerassociation EzsignfoldersignerassociationResponseCompound) *WebhookEzsignEzsignsignerConnect {
	this := WebhookEzsignEzsignsignerConnect{}
	this.ObjWebhook = objWebhook
	this.AObjAttempt = aObjAttempt
	this.ObjEzsignfoldersignerassociation = objEzsignfoldersignerassociation
	return &this
}

// NewWebhookEzsignEzsignsignerConnectWithDefaults instantiates a new WebhookEzsignEzsignsignerConnect object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookEzsignEzsignsignerConnectWithDefaults() *WebhookEzsignEzsignsignerConnect {
	this := WebhookEzsignEzsignsignerConnect{}
	return &this
}

// GetObjWebhook returns the ObjWebhook field value
func (o *WebhookEzsignEzsignsignerConnect) GetObjWebhook() CustomWebhookResponse {
	if o == nil {
		var ret CustomWebhookResponse
		return ret
	}

	return o.ObjWebhook
}

// GetObjWebhookOk returns a tuple with the ObjWebhook field value
// and a boolean to check if the value has been set.
func (o *WebhookEzsignEzsignsignerConnect) GetObjWebhookOk() (*CustomWebhookResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjWebhook, true
}

// SetObjWebhook sets field value
func (o *WebhookEzsignEzsignsignerConnect) SetObjWebhook(v CustomWebhookResponse) {
	o.ObjWebhook = v
}

// GetAObjAttempt returns the AObjAttempt field value
func (o *WebhookEzsignEzsignsignerConnect) GetAObjAttempt() []AttemptResponseCompound {
	if o == nil {
		var ret []AttemptResponseCompound
		return ret
	}

	return o.AObjAttempt
}

// GetAObjAttemptOk returns a tuple with the AObjAttempt field value
// and a boolean to check if the value has been set.
func (o *WebhookEzsignEzsignsignerConnect) GetAObjAttemptOk() ([]AttemptResponseCompound, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjAttempt, true
}

// SetAObjAttempt sets field value
func (o *WebhookEzsignEzsignsignerConnect) SetAObjAttempt(v []AttemptResponseCompound) {
	o.AObjAttempt = v
}

// GetObjEzsignfolder returns the ObjEzsignfolder field value if set, zero value otherwise.
func (o *WebhookEzsignEzsignsignerConnect) GetObjEzsignfolder() EzsignfolderResponse {
	if o == nil || IsNil(o.ObjEzsignfolder) {
		var ret EzsignfolderResponse
		return ret
	}
	return *o.ObjEzsignfolder
}

// GetObjEzsignfolderOk returns a tuple with the ObjEzsignfolder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookEzsignEzsignsignerConnect) GetObjEzsignfolderOk() (*EzsignfolderResponse, bool) {
	if o == nil || IsNil(o.ObjEzsignfolder) {
		return nil, false
	}
	return o.ObjEzsignfolder, true
}

// HasObjEzsignfolder returns a boolean if a field has been set.
func (o *WebhookEzsignEzsignsignerConnect) HasObjEzsignfolder() bool {
	if o != nil && !IsNil(o.ObjEzsignfolder) {
		return true
	}

	return false
}

// SetObjEzsignfolder gets a reference to the given EzsignfolderResponse and assigns it to the ObjEzsignfolder field.
func (o *WebhookEzsignEzsignsignerConnect) SetObjEzsignfolder(v EzsignfolderResponse) {
	o.ObjEzsignfolder = &v
}

// GetObjEzsignfoldersignerassociation returns the ObjEzsignfoldersignerassociation field value
func (o *WebhookEzsignEzsignsignerConnect) GetObjEzsignfoldersignerassociation() EzsignfoldersignerassociationResponseCompound {
	if o == nil {
		var ret EzsignfoldersignerassociationResponseCompound
		return ret
	}

	return o.ObjEzsignfoldersignerassociation
}

// GetObjEzsignfoldersignerassociationOk returns a tuple with the ObjEzsignfoldersignerassociation field value
// and a boolean to check if the value has been set.
func (o *WebhookEzsignEzsignsignerConnect) GetObjEzsignfoldersignerassociationOk() (*EzsignfoldersignerassociationResponseCompound, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjEzsignfoldersignerassociation, true
}

// SetObjEzsignfoldersignerassociation sets field value
func (o *WebhookEzsignEzsignsignerConnect) SetObjEzsignfoldersignerassociation(v EzsignfoldersignerassociationResponseCompound) {
	o.ObjEzsignfoldersignerassociation = v
}

func (o WebhookEzsignEzsignsignerConnect) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookEzsignEzsignsignerConnect) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objWebhook"] = o.ObjWebhook
	toSerialize["a_objAttempt"] = o.AObjAttempt
	if !IsNil(o.ObjEzsignfolder) {
		toSerialize["objEzsignfolder"] = o.ObjEzsignfolder
	}
	toSerialize["objEzsignfoldersignerassociation"] = o.ObjEzsignfoldersignerassociation
	return toSerialize, nil
}

type NullableWebhookEzsignEzsignsignerConnect struct {
	value *WebhookEzsignEzsignsignerConnect
	isSet bool
}

func (v NullableWebhookEzsignEzsignsignerConnect) Get() *WebhookEzsignEzsignsignerConnect {
	return v.value
}

func (v *NullableWebhookEzsignEzsignsignerConnect) Set(val *WebhookEzsignEzsignsignerConnect) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookEzsignEzsignsignerConnect) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookEzsignEzsignsignerConnect) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookEzsignEzsignsignerConnect(val *WebhookEzsignEzsignsignerConnect) *NullableWebhookEzsignEzsignsignerConnect {
	return &NullableWebhookEzsignEzsignsignerConnect{value: val, isSet: true}
}

func (v NullableWebhookEzsignEzsignsignerConnect) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookEzsignEzsignsignerConnect) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


