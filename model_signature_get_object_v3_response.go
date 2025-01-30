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

// checks if the SignatureGetObjectV3Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SignatureGetObjectV3Response{}

// SignatureGetObjectV3Response Response for GET /3/object/signature/{pkiSignatureID}
type SignatureGetObjectV3Response struct {
	ObjDebugPayload CommonResponseObjDebugPayload `json:"objDebugPayload"`
	ObjDebug *CommonResponseObjDebug `json:"objDebug,omitempty"`
	MPayload SignatureGetObjectV3ResponseMPayload `json:"mPayload"`
}

type _SignatureGetObjectV3Response SignatureGetObjectV3Response

// NewSignatureGetObjectV3Response instantiates a new SignatureGetObjectV3Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignatureGetObjectV3Response(objDebugPayload CommonResponseObjDebugPayload, mPayload SignatureGetObjectV3ResponseMPayload) *SignatureGetObjectV3Response {
	this := SignatureGetObjectV3Response{}
	this.ObjDebugPayload = objDebugPayload
	this.MPayload = mPayload
	return &this
}

// NewSignatureGetObjectV3ResponseWithDefaults instantiates a new SignatureGetObjectV3Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignatureGetObjectV3ResponseWithDefaults() *SignatureGetObjectV3Response {
	this := SignatureGetObjectV3Response{}
	return &this
}

// GetObjDebugPayload returns the ObjDebugPayload field value
func (o *SignatureGetObjectV3Response) GetObjDebugPayload() CommonResponseObjDebugPayload {
	if o == nil {
		var ret CommonResponseObjDebugPayload
		return ret
	}

	return o.ObjDebugPayload
}

// GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field value
// and a boolean to check if the value has been set.
func (o *SignatureGetObjectV3Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjDebugPayload, true
}

// SetObjDebugPayload sets field value
func (o *SignatureGetObjectV3Response) SetObjDebugPayload(v CommonResponseObjDebugPayload) {
	o.ObjDebugPayload = v
}

// GetObjDebug returns the ObjDebug field value if set, zero value otherwise.
func (o *SignatureGetObjectV3Response) GetObjDebug() CommonResponseObjDebug {
	if o == nil || IsNil(o.ObjDebug) {
		var ret CommonResponseObjDebug
		return ret
	}
	return *o.ObjDebug
}

// GetObjDebugOk returns a tuple with the ObjDebug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignatureGetObjectV3Response) GetObjDebugOk() (*CommonResponseObjDebug, bool) {
	if o == nil || IsNil(o.ObjDebug) {
		return nil, false
	}
	return o.ObjDebug, true
}

// HasObjDebug returns a boolean if a field has been set.
func (o *SignatureGetObjectV3Response) HasObjDebug() bool {
	if o != nil && !IsNil(o.ObjDebug) {
		return true
	}

	return false
}

// SetObjDebug gets a reference to the given CommonResponseObjDebug and assigns it to the ObjDebug field.
func (o *SignatureGetObjectV3Response) SetObjDebug(v CommonResponseObjDebug) {
	o.ObjDebug = &v
}

// GetMPayload returns the MPayload field value
func (o *SignatureGetObjectV3Response) GetMPayload() SignatureGetObjectV3ResponseMPayload {
	if o == nil {
		var ret SignatureGetObjectV3ResponseMPayload
		return ret
	}

	return o.MPayload
}

// GetMPayloadOk returns a tuple with the MPayload field value
// and a boolean to check if the value has been set.
func (o *SignatureGetObjectV3Response) GetMPayloadOk() (*SignatureGetObjectV3ResponseMPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MPayload, true
}

// SetMPayload sets field value
func (o *SignatureGetObjectV3Response) SetMPayload(v SignatureGetObjectV3ResponseMPayload) {
	o.MPayload = v
}

func (o SignatureGetObjectV3Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SignatureGetObjectV3Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objDebugPayload"] = o.ObjDebugPayload
	if !IsNil(o.ObjDebug) {
		toSerialize["objDebug"] = o.ObjDebug
	}
	toSerialize["mPayload"] = o.MPayload
	return toSerialize, nil
}

func (o *SignatureGetObjectV3Response) UnmarshalJSON(data []byte) (err error) {
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

	varSignatureGetObjectV3Response := _SignatureGetObjectV3Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSignatureGetObjectV3Response)

	if err != nil {
		return err
	}

	*o = SignatureGetObjectV3Response(varSignatureGetObjectV3Response)

	return err
}

type NullableSignatureGetObjectV3Response struct {
	value *SignatureGetObjectV3Response
	isSet bool
}

func (v NullableSignatureGetObjectV3Response) Get() *SignatureGetObjectV3Response {
	return v.value
}

func (v *NullableSignatureGetObjectV3Response) Set(val *SignatureGetObjectV3Response) {
	v.value = val
	v.isSet = true
}

func (v NullableSignatureGetObjectV3Response) IsSet() bool {
	return v.isSet
}

func (v *NullableSignatureGetObjectV3Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignatureGetObjectV3Response(val *SignatureGetObjectV3Response) *NullableSignatureGetObjectV3Response {
	return &NullableSignatureGetObjectV3Response{value: val, isSet: true}
}

func (v NullableSignatureGetObjectV3Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignatureGetObjectV3Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


