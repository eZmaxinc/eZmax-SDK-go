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

// checks if the WebhookGetListV1Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookGetListV1Response{}

// WebhookGetListV1Response Response for GET /1/object/webhook/getList
type WebhookGetListV1Response struct {
	ObjDebugPayload CommonResponseObjDebugPayloadGetList `json:"objDebugPayload"`
	ObjDebug *CommonResponseObjDebug `json:"objDebug,omitempty"`
	MPayload WebhookGetListV1ResponseMPayload `json:"mPayload"`
}

type _WebhookGetListV1Response WebhookGetListV1Response

// NewWebhookGetListV1Response instantiates a new WebhookGetListV1Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookGetListV1Response(objDebugPayload CommonResponseObjDebugPayloadGetList, mPayload WebhookGetListV1ResponseMPayload) *WebhookGetListV1Response {
	this := WebhookGetListV1Response{}
	this.ObjDebugPayload = objDebugPayload
	this.MPayload = mPayload
	return &this
}

// NewWebhookGetListV1ResponseWithDefaults instantiates a new WebhookGetListV1Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookGetListV1ResponseWithDefaults() *WebhookGetListV1Response {
	this := WebhookGetListV1Response{}
	return &this
}

// GetObjDebugPayload returns the ObjDebugPayload field value
func (o *WebhookGetListV1Response) GetObjDebugPayload() CommonResponseObjDebugPayloadGetList {
	if o == nil {
		var ret CommonResponseObjDebugPayloadGetList
		return ret
	}

	return o.ObjDebugPayload
}

// GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field value
// and a boolean to check if the value has been set.
func (o *WebhookGetListV1Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayloadGetList, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjDebugPayload, true
}

// SetObjDebugPayload sets field value
func (o *WebhookGetListV1Response) SetObjDebugPayload(v CommonResponseObjDebugPayloadGetList) {
	o.ObjDebugPayload = v
}

// GetObjDebug returns the ObjDebug field value if set, zero value otherwise.
func (o *WebhookGetListV1Response) GetObjDebug() CommonResponseObjDebug {
	if o == nil || IsNil(o.ObjDebug) {
		var ret CommonResponseObjDebug
		return ret
	}
	return *o.ObjDebug
}

// GetObjDebugOk returns a tuple with the ObjDebug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookGetListV1Response) GetObjDebugOk() (*CommonResponseObjDebug, bool) {
	if o == nil || IsNil(o.ObjDebug) {
		return nil, false
	}
	return o.ObjDebug, true
}

// HasObjDebug returns a boolean if a field has been set.
func (o *WebhookGetListV1Response) HasObjDebug() bool {
	if o != nil && !IsNil(o.ObjDebug) {
		return true
	}

	return false
}

// SetObjDebug gets a reference to the given CommonResponseObjDebug and assigns it to the ObjDebug field.
func (o *WebhookGetListV1Response) SetObjDebug(v CommonResponseObjDebug) {
	o.ObjDebug = &v
}

// GetMPayload returns the MPayload field value
func (o *WebhookGetListV1Response) GetMPayload() WebhookGetListV1ResponseMPayload {
	if o == nil {
		var ret WebhookGetListV1ResponseMPayload
		return ret
	}

	return o.MPayload
}

// GetMPayloadOk returns a tuple with the MPayload field value
// and a boolean to check if the value has been set.
func (o *WebhookGetListV1Response) GetMPayloadOk() (*WebhookGetListV1ResponseMPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MPayload, true
}

// SetMPayload sets field value
func (o *WebhookGetListV1Response) SetMPayload(v WebhookGetListV1ResponseMPayload) {
	o.MPayload = v
}

func (o WebhookGetListV1Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookGetListV1Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objDebugPayload"] = o.ObjDebugPayload
	if !IsNil(o.ObjDebug) {
		toSerialize["objDebug"] = o.ObjDebug
	}
	toSerialize["mPayload"] = o.MPayload
	return toSerialize, nil
}

func (o *WebhookGetListV1Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"objDebugPayload",
		"mPayload",
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

	varWebhookGetListV1Response := _WebhookGetListV1Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWebhookGetListV1Response)

	if err != nil {
		return err
	}

	*o = WebhookGetListV1Response(varWebhookGetListV1Response)

	return err
}

type NullableWebhookGetListV1Response struct {
	value *WebhookGetListV1Response
	isSet bool
}

func (v NullableWebhookGetListV1Response) Get() *WebhookGetListV1Response {
	return v.value
}

func (v *NullableWebhookGetListV1Response) Set(val *WebhookGetListV1Response) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookGetListV1Response) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookGetListV1Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookGetListV1Response(val *WebhookGetListV1Response) *NullableWebhookGetListV1Response {
	return &NullableWebhookGetListV1Response{value: val, isSet: true}
}

func (v NullableWebhookGetListV1Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookGetListV1Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


