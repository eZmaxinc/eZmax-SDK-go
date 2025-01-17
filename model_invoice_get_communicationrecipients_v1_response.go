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

// checks if the InvoiceGetCommunicationrecipientsV1Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvoiceGetCommunicationrecipientsV1Response{}

// InvoiceGetCommunicationrecipientsV1Response Response for GET /1/object/invoice/{pkiInvoiceID}/getCommunicationrecipients
type InvoiceGetCommunicationrecipientsV1Response struct {
	ObjDebugPayload CommonResponseObjDebugPayload `json:"objDebugPayload"`
	ObjDebug *CommonResponseObjDebug `json:"objDebug,omitempty"`
	MPayload InvoiceGetCommunicationrecipientsV1ResponseMPayload `json:"mPayload"`
}

type _InvoiceGetCommunicationrecipientsV1Response InvoiceGetCommunicationrecipientsV1Response

// NewInvoiceGetCommunicationrecipientsV1Response instantiates a new InvoiceGetCommunicationrecipientsV1Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoiceGetCommunicationrecipientsV1Response(objDebugPayload CommonResponseObjDebugPayload, mPayload InvoiceGetCommunicationrecipientsV1ResponseMPayload) *InvoiceGetCommunicationrecipientsV1Response {
	this := InvoiceGetCommunicationrecipientsV1Response{}
	this.ObjDebugPayload = objDebugPayload
	this.MPayload = mPayload
	return &this
}

// NewInvoiceGetCommunicationrecipientsV1ResponseWithDefaults instantiates a new InvoiceGetCommunicationrecipientsV1Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceGetCommunicationrecipientsV1ResponseWithDefaults() *InvoiceGetCommunicationrecipientsV1Response {
	this := InvoiceGetCommunicationrecipientsV1Response{}
	return &this
}

// GetObjDebugPayload returns the ObjDebugPayload field value
func (o *InvoiceGetCommunicationrecipientsV1Response) GetObjDebugPayload() CommonResponseObjDebugPayload {
	if o == nil {
		var ret CommonResponseObjDebugPayload
		return ret
	}

	return o.ObjDebugPayload
}

// GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field value
// and a boolean to check if the value has been set.
func (o *InvoiceGetCommunicationrecipientsV1Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjDebugPayload, true
}

// SetObjDebugPayload sets field value
func (o *InvoiceGetCommunicationrecipientsV1Response) SetObjDebugPayload(v CommonResponseObjDebugPayload) {
	o.ObjDebugPayload = v
}

// GetObjDebug returns the ObjDebug field value if set, zero value otherwise.
func (o *InvoiceGetCommunicationrecipientsV1Response) GetObjDebug() CommonResponseObjDebug {
	if o == nil || IsNil(o.ObjDebug) {
		var ret CommonResponseObjDebug
		return ret
	}
	return *o.ObjDebug
}

// GetObjDebugOk returns a tuple with the ObjDebug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceGetCommunicationrecipientsV1Response) GetObjDebugOk() (*CommonResponseObjDebug, bool) {
	if o == nil || IsNil(o.ObjDebug) {
		return nil, false
	}
	return o.ObjDebug, true
}

// HasObjDebug returns a boolean if a field has been set.
func (o *InvoiceGetCommunicationrecipientsV1Response) HasObjDebug() bool {
	if o != nil && !IsNil(o.ObjDebug) {
		return true
	}

	return false
}

// SetObjDebug gets a reference to the given CommonResponseObjDebug and assigns it to the ObjDebug field.
func (o *InvoiceGetCommunicationrecipientsV1Response) SetObjDebug(v CommonResponseObjDebug) {
	o.ObjDebug = &v
}

// GetMPayload returns the MPayload field value
func (o *InvoiceGetCommunicationrecipientsV1Response) GetMPayload() InvoiceGetCommunicationrecipientsV1ResponseMPayload {
	if o == nil {
		var ret InvoiceGetCommunicationrecipientsV1ResponseMPayload
		return ret
	}

	return o.MPayload
}

// GetMPayloadOk returns a tuple with the MPayload field value
// and a boolean to check if the value has been set.
func (o *InvoiceGetCommunicationrecipientsV1Response) GetMPayloadOk() (*InvoiceGetCommunicationrecipientsV1ResponseMPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MPayload, true
}

// SetMPayload sets field value
func (o *InvoiceGetCommunicationrecipientsV1Response) SetMPayload(v InvoiceGetCommunicationrecipientsV1ResponseMPayload) {
	o.MPayload = v
}

func (o InvoiceGetCommunicationrecipientsV1Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvoiceGetCommunicationrecipientsV1Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objDebugPayload"] = o.ObjDebugPayload
	if !IsNil(o.ObjDebug) {
		toSerialize["objDebug"] = o.ObjDebug
	}
	toSerialize["mPayload"] = o.MPayload
	return toSerialize, nil
}

func (o *InvoiceGetCommunicationrecipientsV1Response) UnmarshalJSON(data []byte) (err error) {
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

	varInvoiceGetCommunicationrecipientsV1Response := _InvoiceGetCommunicationrecipientsV1Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInvoiceGetCommunicationrecipientsV1Response)

	if err != nil {
		return err
	}

	*o = InvoiceGetCommunicationrecipientsV1Response(varInvoiceGetCommunicationrecipientsV1Response)

	return err
}

type NullableInvoiceGetCommunicationrecipientsV1Response struct {
	value *InvoiceGetCommunicationrecipientsV1Response
	isSet bool
}

func (v NullableInvoiceGetCommunicationrecipientsV1Response) Get() *InvoiceGetCommunicationrecipientsV1Response {
	return v.value
}

func (v *NullableInvoiceGetCommunicationrecipientsV1Response) Set(val *InvoiceGetCommunicationrecipientsV1Response) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoiceGetCommunicationrecipientsV1Response) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoiceGetCommunicationrecipientsV1Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoiceGetCommunicationrecipientsV1Response(val *InvoiceGetCommunicationrecipientsV1Response) *NullableInvoiceGetCommunicationrecipientsV1Response {
	return &NullableInvoiceGetCommunicationrecipientsV1Response{value: val, isSet: true}
}

func (v NullableInvoiceGetCommunicationrecipientsV1Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoiceGetCommunicationrecipientsV1Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


