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

// checks if the EzsigntemplatedocumentGetWordsPositionsV1Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigntemplatedocumentGetWordsPositionsV1Response{}

// EzsigntemplatedocumentGetWordsPositionsV1Response Response for POST /1/object/ezsigntemplatedocument/{pkiEzsigntemplatedocumentID}/getWordsPositions
type EzsigntemplatedocumentGetWordsPositionsV1Response struct {
	ObjDebugPayload CommonResponseObjDebugPayload `json:"objDebugPayload"`
	ObjDebug *CommonResponseObjDebug `json:"objDebug,omitempty"`
	// Payload for POST /1/object/ezsigntemplatedocument/{pkiEzsigntemplatedocumentID}/getWordsPositions
	MPayload []CustomWordPositionWordResponse `json:"mPayload"`
}

type _EzsigntemplatedocumentGetWordsPositionsV1Response EzsigntemplatedocumentGetWordsPositionsV1Response

// NewEzsigntemplatedocumentGetWordsPositionsV1Response instantiates a new EzsigntemplatedocumentGetWordsPositionsV1Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigntemplatedocumentGetWordsPositionsV1Response(objDebugPayload CommonResponseObjDebugPayload, mPayload []CustomWordPositionWordResponse) *EzsigntemplatedocumentGetWordsPositionsV1Response {
	this := EzsigntemplatedocumentGetWordsPositionsV1Response{}
	this.ObjDebugPayload = objDebugPayload
	this.MPayload = mPayload
	return &this
}

// NewEzsigntemplatedocumentGetWordsPositionsV1ResponseWithDefaults instantiates a new EzsigntemplatedocumentGetWordsPositionsV1Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigntemplatedocumentGetWordsPositionsV1ResponseWithDefaults() *EzsigntemplatedocumentGetWordsPositionsV1Response {
	this := EzsigntemplatedocumentGetWordsPositionsV1Response{}
	return &this
}

// GetObjDebugPayload returns the ObjDebugPayload field value
func (o *EzsigntemplatedocumentGetWordsPositionsV1Response) GetObjDebugPayload() CommonResponseObjDebugPayload {
	if o == nil {
		var ret CommonResponseObjDebugPayload
		return ret
	}

	return o.ObjDebugPayload
}

// GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatedocumentGetWordsPositionsV1Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjDebugPayload, true
}

// SetObjDebugPayload sets field value
func (o *EzsigntemplatedocumentGetWordsPositionsV1Response) SetObjDebugPayload(v CommonResponseObjDebugPayload) {
	o.ObjDebugPayload = v
}

// GetObjDebug returns the ObjDebug field value if set, zero value otherwise.
func (o *EzsigntemplatedocumentGetWordsPositionsV1Response) GetObjDebug() CommonResponseObjDebug {
	if o == nil || IsNil(o.ObjDebug) {
		var ret CommonResponseObjDebug
		return ret
	}
	return *o.ObjDebug
}

// GetObjDebugOk returns a tuple with the ObjDebug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplatedocumentGetWordsPositionsV1Response) GetObjDebugOk() (*CommonResponseObjDebug, bool) {
	if o == nil || IsNil(o.ObjDebug) {
		return nil, false
	}
	return o.ObjDebug, true
}

// HasObjDebug returns a boolean if a field has been set.
func (o *EzsigntemplatedocumentGetWordsPositionsV1Response) HasObjDebug() bool {
	if o != nil && !IsNil(o.ObjDebug) {
		return true
	}

	return false
}

// SetObjDebug gets a reference to the given CommonResponseObjDebug and assigns it to the ObjDebug field.
func (o *EzsigntemplatedocumentGetWordsPositionsV1Response) SetObjDebug(v CommonResponseObjDebug) {
	o.ObjDebug = &v
}

// GetMPayload returns the MPayload field value
func (o *EzsigntemplatedocumentGetWordsPositionsV1Response) GetMPayload() []CustomWordPositionWordResponse {
	if o == nil {
		var ret []CustomWordPositionWordResponse
		return ret
	}

	return o.MPayload
}

// GetMPayloadOk returns a tuple with the MPayload field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatedocumentGetWordsPositionsV1Response) GetMPayloadOk() ([]CustomWordPositionWordResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.MPayload, true
}

// SetMPayload sets field value
func (o *EzsigntemplatedocumentGetWordsPositionsV1Response) SetMPayload(v []CustomWordPositionWordResponse) {
	o.MPayload = v
}

func (o EzsigntemplatedocumentGetWordsPositionsV1Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigntemplatedocumentGetWordsPositionsV1Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objDebugPayload"] = o.ObjDebugPayload
	if !IsNil(o.ObjDebug) {
		toSerialize["objDebug"] = o.ObjDebug
	}
	toSerialize["mPayload"] = o.MPayload
	return toSerialize, nil
}

func (o *EzsigntemplatedocumentGetWordsPositionsV1Response) UnmarshalJSON(data []byte) (err error) {
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

	varEzsigntemplatedocumentGetWordsPositionsV1Response := _EzsigntemplatedocumentGetWordsPositionsV1Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsigntemplatedocumentGetWordsPositionsV1Response)

	if err != nil {
		return err
	}

	*o = EzsigntemplatedocumentGetWordsPositionsV1Response(varEzsigntemplatedocumentGetWordsPositionsV1Response)

	return err
}

type NullableEzsigntemplatedocumentGetWordsPositionsV1Response struct {
	value *EzsigntemplatedocumentGetWordsPositionsV1Response
	isSet bool
}

func (v NullableEzsigntemplatedocumentGetWordsPositionsV1Response) Get() *EzsigntemplatedocumentGetWordsPositionsV1Response {
	return v.value
}

func (v *NullableEzsigntemplatedocumentGetWordsPositionsV1Response) Set(val *EzsigntemplatedocumentGetWordsPositionsV1Response) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigntemplatedocumentGetWordsPositionsV1Response) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigntemplatedocumentGetWordsPositionsV1Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigntemplatedocumentGetWordsPositionsV1Response(val *EzsigntemplatedocumentGetWordsPositionsV1Response) *NullableEzsigntemplatedocumentGetWordsPositionsV1Response {
	return &NullableEzsigntemplatedocumentGetWordsPositionsV1Response{value: val, isSet: true}
}

func (v NullableEzsigntemplatedocumentGetWordsPositionsV1Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigntemplatedocumentGetWordsPositionsV1Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


