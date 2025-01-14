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

// checks if the ElectronicfundstransferGetCommunicationCountV1Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ElectronicfundstransferGetCommunicationCountV1Response{}

// ElectronicfundstransferGetCommunicationCountV1Response Response for GET /1/object/electronicfundstransfer/{pkiElectronicfundstransferID}/getCommunicationCount
type ElectronicfundstransferGetCommunicationCountV1Response struct {
	ObjDebugPayload CommonResponseObjDebugPayload `json:"objDebugPayload"`
	ObjDebug *CommonResponseObjDebug `json:"objDebug,omitempty"`
	MPayload ElectronicfundstransferGetCommunicationCountV1ResponseMPayload `json:"mPayload"`
}

type _ElectronicfundstransferGetCommunicationCountV1Response ElectronicfundstransferGetCommunicationCountV1Response

// NewElectronicfundstransferGetCommunicationCountV1Response instantiates a new ElectronicfundstransferGetCommunicationCountV1Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewElectronicfundstransferGetCommunicationCountV1Response(objDebugPayload CommonResponseObjDebugPayload, mPayload ElectronicfundstransferGetCommunicationCountV1ResponseMPayload) *ElectronicfundstransferGetCommunicationCountV1Response {
	this := ElectronicfundstransferGetCommunicationCountV1Response{}
	this.ObjDebugPayload = objDebugPayload
	this.MPayload = mPayload
	return &this
}

// NewElectronicfundstransferGetCommunicationCountV1ResponseWithDefaults instantiates a new ElectronicfundstransferGetCommunicationCountV1Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewElectronicfundstransferGetCommunicationCountV1ResponseWithDefaults() *ElectronicfundstransferGetCommunicationCountV1Response {
	this := ElectronicfundstransferGetCommunicationCountV1Response{}
	return &this
}

// GetObjDebugPayload returns the ObjDebugPayload field value
func (o *ElectronicfundstransferGetCommunicationCountV1Response) GetObjDebugPayload() CommonResponseObjDebugPayload {
	if o == nil {
		var ret CommonResponseObjDebugPayload
		return ret
	}

	return o.ObjDebugPayload
}

// GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field value
// and a boolean to check if the value has been set.
func (o *ElectronicfundstransferGetCommunicationCountV1Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjDebugPayload, true
}

// SetObjDebugPayload sets field value
func (o *ElectronicfundstransferGetCommunicationCountV1Response) SetObjDebugPayload(v CommonResponseObjDebugPayload) {
	o.ObjDebugPayload = v
}

// GetObjDebug returns the ObjDebug field value if set, zero value otherwise.
func (o *ElectronicfundstransferGetCommunicationCountV1Response) GetObjDebug() CommonResponseObjDebug {
	if o == nil || IsNil(o.ObjDebug) {
		var ret CommonResponseObjDebug
		return ret
	}
	return *o.ObjDebug
}

// GetObjDebugOk returns a tuple with the ObjDebug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElectronicfundstransferGetCommunicationCountV1Response) GetObjDebugOk() (*CommonResponseObjDebug, bool) {
	if o == nil || IsNil(o.ObjDebug) {
		return nil, false
	}
	return o.ObjDebug, true
}

// HasObjDebug returns a boolean if a field has been set.
func (o *ElectronicfundstransferGetCommunicationCountV1Response) HasObjDebug() bool {
	if o != nil && !IsNil(o.ObjDebug) {
		return true
	}

	return false
}

// SetObjDebug gets a reference to the given CommonResponseObjDebug and assigns it to the ObjDebug field.
func (o *ElectronicfundstransferGetCommunicationCountV1Response) SetObjDebug(v CommonResponseObjDebug) {
	o.ObjDebug = &v
}

// GetMPayload returns the MPayload field value
func (o *ElectronicfundstransferGetCommunicationCountV1Response) GetMPayload() ElectronicfundstransferGetCommunicationCountV1ResponseMPayload {
	if o == nil {
		var ret ElectronicfundstransferGetCommunicationCountV1ResponseMPayload
		return ret
	}

	return o.MPayload
}

// GetMPayloadOk returns a tuple with the MPayload field value
// and a boolean to check if the value has been set.
func (o *ElectronicfundstransferGetCommunicationCountV1Response) GetMPayloadOk() (*ElectronicfundstransferGetCommunicationCountV1ResponseMPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MPayload, true
}

// SetMPayload sets field value
func (o *ElectronicfundstransferGetCommunicationCountV1Response) SetMPayload(v ElectronicfundstransferGetCommunicationCountV1ResponseMPayload) {
	o.MPayload = v
}

func (o ElectronicfundstransferGetCommunicationCountV1Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ElectronicfundstransferGetCommunicationCountV1Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objDebugPayload"] = o.ObjDebugPayload
	if !IsNil(o.ObjDebug) {
		toSerialize["objDebug"] = o.ObjDebug
	}
	toSerialize["mPayload"] = o.MPayload
	return toSerialize, nil
}

func (o *ElectronicfundstransferGetCommunicationCountV1Response) UnmarshalJSON(data []byte) (err error) {
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

	varElectronicfundstransferGetCommunicationCountV1Response := _ElectronicfundstransferGetCommunicationCountV1Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varElectronicfundstransferGetCommunicationCountV1Response)

	if err != nil {
		return err
	}

	*o = ElectronicfundstransferGetCommunicationCountV1Response(varElectronicfundstransferGetCommunicationCountV1Response)

	return err
}

type NullableElectronicfundstransferGetCommunicationCountV1Response struct {
	value *ElectronicfundstransferGetCommunicationCountV1Response
	isSet bool
}

func (v NullableElectronicfundstransferGetCommunicationCountV1Response) Get() *ElectronicfundstransferGetCommunicationCountV1Response {
	return v.value
}

func (v *NullableElectronicfundstransferGetCommunicationCountV1Response) Set(val *ElectronicfundstransferGetCommunicationCountV1Response) {
	v.value = val
	v.isSet = true
}

func (v NullableElectronicfundstransferGetCommunicationCountV1Response) IsSet() bool {
	return v.isSet
}

func (v *NullableElectronicfundstransferGetCommunicationCountV1Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableElectronicfundstransferGetCommunicationCountV1Response(val *ElectronicfundstransferGetCommunicationCountV1Response) *NullableElectronicfundstransferGetCommunicationCountV1Response {
	return &NullableElectronicfundstransferGetCommunicationCountV1Response{value: val, isSet: true}
}

func (v NullableElectronicfundstransferGetCommunicationCountV1Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableElectronicfundstransferGetCommunicationCountV1Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


