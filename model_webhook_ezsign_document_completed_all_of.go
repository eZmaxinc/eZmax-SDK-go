/*
 * eZmax API Definition
 *
 * This API expose all the functionnalities for the eZmax and eZsign application.
 *
 * API version: 1.0.30
 * Contact: support-api@ezmax.ca
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxinc/eZmax-SDK-go

import (
	"encoding/json"
)

// WebhookEzsignDocumentCompletedAllOf struct for WebhookEzsignDocumentCompletedAllOf
type WebhookEzsignDocumentCompletedAllOf struct {
	ObjEzsigndocument EzsigndocumentResponse `json:"objEzsigndocument"`
}

// NewWebhookEzsignDocumentCompletedAllOf instantiates a new WebhookEzsignDocumentCompletedAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookEzsignDocumentCompletedAllOf(objEzsigndocument EzsigndocumentResponse, ) *WebhookEzsignDocumentCompletedAllOf {
	this := WebhookEzsignDocumentCompletedAllOf{}
	this.ObjEzsigndocument = objEzsigndocument
	return &this
}

// NewWebhookEzsignDocumentCompletedAllOfWithDefaults instantiates a new WebhookEzsignDocumentCompletedAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookEzsignDocumentCompletedAllOfWithDefaults() *WebhookEzsignDocumentCompletedAllOf {
	this := WebhookEzsignDocumentCompletedAllOf{}
	return &this
}

// GetObjEzsigndocument returns the ObjEzsigndocument field value
func (o *WebhookEzsignDocumentCompletedAllOf) GetObjEzsigndocument() EzsigndocumentResponse {
	if o == nil  {
		var ret EzsigndocumentResponse
		return ret
	}

	return o.ObjEzsigndocument
}

// GetObjEzsigndocumentOk returns a tuple with the ObjEzsigndocument field value
// and a boolean to check if the value has been set.
func (o *WebhookEzsignDocumentCompletedAllOf) GetObjEzsigndocumentOk() (*EzsigndocumentResponse, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjEzsigndocument, true
}

// SetObjEzsigndocument sets field value
func (o *WebhookEzsignDocumentCompletedAllOf) SetObjEzsigndocument(v EzsigndocumentResponse) {
	o.ObjEzsigndocument = v
}

func (o WebhookEzsignDocumentCompletedAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["objEzsigndocument"] = o.ObjEzsigndocument
	}
	return json.Marshal(toSerialize)
}

type NullableWebhookEzsignDocumentCompletedAllOf struct {
	value *WebhookEzsignDocumentCompletedAllOf
	isSet bool
}

func (v NullableWebhookEzsignDocumentCompletedAllOf) Get() *WebhookEzsignDocumentCompletedAllOf {
	return v.value
}

func (v *NullableWebhookEzsignDocumentCompletedAllOf) Set(val *WebhookEzsignDocumentCompletedAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookEzsignDocumentCompletedAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookEzsignDocumentCompletedAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookEzsignDocumentCompletedAllOf(val *WebhookEzsignDocumentCompletedAllOf) *NullableWebhookEzsignDocumentCompletedAllOf {
	return &NullableWebhookEzsignDocumentCompletedAllOf{value: val, isSet: true}
}

func (v NullableWebhookEzsignDocumentCompletedAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookEzsignDocumentCompletedAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


