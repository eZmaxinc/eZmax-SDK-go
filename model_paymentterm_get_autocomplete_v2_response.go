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

// checks if the PaymenttermGetAutocompleteV2Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymenttermGetAutocompleteV2Response{}

// PaymenttermGetAutocompleteV2Response Response for GET /2/object/paymentterm/getAutocomplete
type PaymenttermGetAutocompleteV2Response struct {
	ObjDebugPayload CommonResponseObjDebugPayload `json:"objDebugPayload"`
	ObjDebug *CommonResponseObjDebug `json:"objDebug,omitempty"`
	MPayload PaymenttermGetAutocompleteV2ResponseMPayload `json:"mPayload"`
}

type _PaymenttermGetAutocompleteV2Response PaymenttermGetAutocompleteV2Response

// NewPaymenttermGetAutocompleteV2Response instantiates a new PaymenttermGetAutocompleteV2Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymenttermGetAutocompleteV2Response(objDebugPayload CommonResponseObjDebugPayload, mPayload PaymenttermGetAutocompleteV2ResponseMPayload) *PaymenttermGetAutocompleteV2Response {
	this := PaymenttermGetAutocompleteV2Response{}
	this.ObjDebugPayload = objDebugPayload
	this.MPayload = mPayload
	return &this
}

// NewPaymenttermGetAutocompleteV2ResponseWithDefaults instantiates a new PaymenttermGetAutocompleteV2Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymenttermGetAutocompleteV2ResponseWithDefaults() *PaymenttermGetAutocompleteV2Response {
	this := PaymenttermGetAutocompleteV2Response{}
	return &this
}

// GetObjDebugPayload returns the ObjDebugPayload field value
func (o *PaymenttermGetAutocompleteV2Response) GetObjDebugPayload() CommonResponseObjDebugPayload {
	if o == nil {
		var ret CommonResponseObjDebugPayload
		return ret
	}

	return o.ObjDebugPayload
}

// GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field value
// and a boolean to check if the value has been set.
func (o *PaymenttermGetAutocompleteV2Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjDebugPayload, true
}

// SetObjDebugPayload sets field value
func (o *PaymenttermGetAutocompleteV2Response) SetObjDebugPayload(v CommonResponseObjDebugPayload) {
	o.ObjDebugPayload = v
}

// GetObjDebug returns the ObjDebug field value if set, zero value otherwise.
func (o *PaymenttermGetAutocompleteV2Response) GetObjDebug() CommonResponseObjDebug {
	if o == nil || IsNil(o.ObjDebug) {
		var ret CommonResponseObjDebug
		return ret
	}
	return *o.ObjDebug
}

// GetObjDebugOk returns a tuple with the ObjDebug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymenttermGetAutocompleteV2Response) GetObjDebugOk() (*CommonResponseObjDebug, bool) {
	if o == nil || IsNil(o.ObjDebug) {
		return nil, false
	}
	return o.ObjDebug, true
}

// HasObjDebug returns a boolean if a field has been set.
func (o *PaymenttermGetAutocompleteV2Response) HasObjDebug() bool {
	if o != nil && !IsNil(o.ObjDebug) {
		return true
	}

	return false
}

// SetObjDebug gets a reference to the given CommonResponseObjDebug and assigns it to the ObjDebug field.
func (o *PaymenttermGetAutocompleteV2Response) SetObjDebug(v CommonResponseObjDebug) {
	o.ObjDebug = &v
}

// GetMPayload returns the MPayload field value
func (o *PaymenttermGetAutocompleteV2Response) GetMPayload() PaymenttermGetAutocompleteV2ResponseMPayload {
	if o == nil {
		var ret PaymenttermGetAutocompleteV2ResponseMPayload
		return ret
	}

	return o.MPayload
}

// GetMPayloadOk returns a tuple with the MPayload field value
// and a boolean to check if the value has been set.
func (o *PaymenttermGetAutocompleteV2Response) GetMPayloadOk() (*PaymenttermGetAutocompleteV2ResponseMPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MPayload, true
}

// SetMPayload sets field value
func (o *PaymenttermGetAutocompleteV2Response) SetMPayload(v PaymenttermGetAutocompleteV2ResponseMPayload) {
	o.MPayload = v
}

func (o PaymenttermGetAutocompleteV2Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymenttermGetAutocompleteV2Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objDebugPayload"] = o.ObjDebugPayload
	if !IsNil(o.ObjDebug) {
		toSerialize["objDebug"] = o.ObjDebug
	}
	toSerialize["mPayload"] = o.MPayload
	return toSerialize, nil
}

func (o *PaymenttermGetAutocompleteV2Response) UnmarshalJSON(data []byte) (err error) {
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

	varPaymenttermGetAutocompleteV2Response := _PaymenttermGetAutocompleteV2Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPaymenttermGetAutocompleteV2Response)

	if err != nil {
		return err
	}

	*o = PaymenttermGetAutocompleteV2Response(varPaymenttermGetAutocompleteV2Response)

	return err
}

type NullablePaymenttermGetAutocompleteV2Response struct {
	value *PaymenttermGetAutocompleteV2Response
	isSet bool
}

func (v NullablePaymenttermGetAutocompleteV2Response) Get() *PaymenttermGetAutocompleteV2Response {
	return v.value
}

func (v *NullablePaymenttermGetAutocompleteV2Response) Set(val *PaymenttermGetAutocompleteV2Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymenttermGetAutocompleteV2Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymenttermGetAutocompleteV2Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymenttermGetAutocompleteV2Response(val *PaymenttermGetAutocompleteV2Response) *NullablePaymenttermGetAutocompleteV2Response {
	return &NullablePaymenttermGetAutocompleteV2Response{value: val, isSet: true}
}

func (v NullablePaymenttermGetAutocompleteV2Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymenttermGetAutocompleteV2Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


