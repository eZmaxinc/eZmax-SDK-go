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

// checks if the CreditcardclientCreateObjectV1Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreditcardclientCreateObjectV1Response{}

// CreditcardclientCreateObjectV1Response Response for POST /1/object/creditcardclient
type CreditcardclientCreateObjectV1Response struct {
	ObjDebugPayload CommonResponseObjDebugPayload `json:"objDebugPayload"`
	ObjDebug *CommonResponseObjDebug `json:"objDebug,omitempty"`
	MPayload CreditcardclientCreateObjectV1ResponseMPayload `json:"mPayload"`
}

type _CreditcardclientCreateObjectV1Response CreditcardclientCreateObjectV1Response

// NewCreditcardclientCreateObjectV1Response instantiates a new CreditcardclientCreateObjectV1Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditcardclientCreateObjectV1Response(objDebugPayload CommonResponseObjDebugPayload, mPayload CreditcardclientCreateObjectV1ResponseMPayload) *CreditcardclientCreateObjectV1Response {
	this := CreditcardclientCreateObjectV1Response{}
	this.ObjDebugPayload = objDebugPayload
	this.MPayload = mPayload
	return &this
}

// NewCreditcardclientCreateObjectV1ResponseWithDefaults instantiates a new CreditcardclientCreateObjectV1Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditcardclientCreateObjectV1ResponseWithDefaults() *CreditcardclientCreateObjectV1Response {
	this := CreditcardclientCreateObjectV1Response{}
	return &this
}

// GetObjDebugPayload returns the ObjDebugPayload field value
func (o *CreditcardclientCreateObjectV1Response) GetObjDebugPayload() CommonResponseObjDebugPayload {
	if o == nil {
		var ret CommonResponseObjDebugPayload
		return ret
	}

	return o.ObjDebugPayload
}

// GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field value
// and a boolean to check if the value has been set.
func (o *CreditcardclientCreateObjectV1Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjDebugPayload, true
}

// SetObjDebugPayload sets field value
func (o *CreditcardclientCreateObjectV1Response) SetObjDebugPayload(v CommonResponseObjDebugPayload) {
	o.ObjDebugPayload = v
}

// GetObjDebug returns the ObjDebug field value if set, zero value otherwise.
func (o *CreditcardclientCreateObjectV1Response) GetObjDebug() CommonResponseObjDebug {
	if o == nil || IsNil(o.ObjDebug) {
		var ret CommonResponseObjDebug
		return ret
	}
	return *o.ObjDebug
}

// GetObjDebugOk returns a tuple with the ObjDebug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditcardclientCreateObjectV1Response) GetObjDebugOk() (*CommonResponseObjDebug, bool) {
	if o == nil || IsNil(o.ObjDebug) {
		return nil, false
	}
	return o.ObjDebug, true
}

// HasObjDebug returns a boolean if a field has been set.
func (o *CreditcardclientCreateObjectV1Response) HasObjDebug() bool {
	if o != nil && !IsNil(o.ObjDebug) {
		return true
	}

	return false
}

// SetObjDebug gets a reference to the given CommonResponseObjDebug and assigns it to the ObjDebug field.
func (o *CreditcardclientCreateObjectV1Response) SetObjDebug(v CommonResponseObjDebug) {
	o.ObjDebug = &v
}

// GetMPayload returns the MPayload field value
func (o *CreditcardclientCreateObjectV1Response) GetMPayload() CreditcardclientCreateObjectV1ResponseMPayload {
	if o == nil {
		var ret CreditcardclientCreateObjectV1ResponseMPayload
		return ret
	}

	return o.MPayload
}

// GetMPayloadOk returns a tuple with the MPayload field value
// and a boolean to check if the value has been set.
func (o *CreditcardclientCreateObjectV1Response) GetMPayloadOk() (*CreditcardclientCreateObjectV1ResponseMPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MPayload, true
}

// SetMPayload sets field value
func (o *CreditcardclientCreateObjectV1Response) SetMPayload(v CreditcardclientCreateObjectV1ResponseMPayload) {
	o.MPayload = v
}

func (o CreditcardclientCreateObjectV1Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreditcardclientCreateObjectV1Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objDebugPayload"] = o.ObjDebugPayload
	if !IsNil(o.ObjDebug) {
		toSerialize["objDebug"] = o.ObjDebug
	}
	toSerialize["mPayload"] = o.MPayload
	return toSerialize, nil
}

func (o *CreditcardclientCreateObjectV1Response) UnmarshalJSON(data []byte) (err error) {
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

	varCreditcardclientCreateObjectV1Response := _CreditcardclientCreateObjectV1Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreditcardclientCreateObjectV1Response)

	if err != nil {
		return err
	}

	*o = CreditcardclientCreateObjectV1Response(varCreditcardclientCreateObjectV1Response)

	return err
}

type NullableCreditcardclientCreateObjectV1Response struct {
	value *CreditcardclientCreateObjectV1Response
	isSet bool
}

func (v NullableCreditcardclientCreateObjectV1Response) Get() *CreditcardclientCreateObjectV1Response {
	return v.value
}

func (v *NullableCreditcardclientCreateObjectV1Response) Set(val *CreditcardclientCreateObjectV1Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditcardclientCreateObjectV1Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditcardclientCreateObjectV1Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditcardclientCreateObjectV1Response(val *CreditcardclientCreateObjectV1Response) *NullableCreditcardclientCreateObjectV1Response {
	return &NullableCreditcardclientCreateObjectV1Response{value: val, isSet: true}
}

func (v NullableCreditcardclientCreateObjectV1Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditcardclientCreateObjectV1Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


