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

// checks if the EzsignbulksendGetObjectV2Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignbulksendGetObjectV2Response{}

// EzsignbulksendGetObjectV2Response Response for GET /2/object/ezsignbulksend/{pkiEzsignbulksendID}
type EzsignbulksendGetObjectV2Response struct {
	ObjDebugPayload CommonResponseObjDebugPayload `json:"objDebugPayload"`
	ObjDebug *CommonResponseObjDebug `json:"objDebug,omitempty"`
	MPayload EzsignbulksendGetObjectV2ResponseMPayload `json:"mPayload"`
}

type _EzsignbulksendGetObjectV2Response EzsignbulksendGetObjectV2Response

// NewEzsignbulksendGetObjectV2Response instantiates a new EzsignbulksendGetObjectV2Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignbulksendGetObjectV2Response(objDebugPayload CommonResponseObjDebugPayload, mPayload EzsignbulksendGetObjectV2ResponseMPayload) *EzsignbulksendGetObjectV2Response {
	this := EzsignbulksendGetObjectV2Response{}
	this.ObjDebugPayload = objDebugPayload
	this.MPayload = mPayload
	return &this
}

// NewEzsignbulksendGetObjectV2ResponseWithDefaults instantiates a new EzsignbulksendGetObjectV2Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignbulksendGetObjectV2ResponseWithDefaults() *EzsignbulksendGetObjectV2Response {
	this := EzsignbulksendGetObjectV2Response{}
	return &this
}

// GetObjDebugPayload returns the ObjDebugPayload field value
func (o *EzsignbulksendGetObjectV2Response) GetObjDebugPayload() CommonResponseObjDebugPayload {
	if o == nil {
		var ret CommonResponseObjDebugPayload
		return ret
	}

	return o.ObjDebugPayload
}

// GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field value
// and a boolean to check if the value has been set.
func (o *EzsignbulksendGetObjectV2Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjDebugPayload, true
}

// SetObjDebugPayload sets field value
func (o *EzsignbulksendGetObjectV2Response) SetObjDebugPayload(v CommonResponseObjDebugPayload) {
	o.ObjDebugPayload = v
}

// GetObjDebug returns the ObjDebug field value if set, zero value otherwise.
func (o *EzsignbulksendGetObjectV2Response) GetObjDebug() CommonResponseObjDebug {
	if o == nil || IsNil(o.ObjDebug) {
		var ret CommonResponseObjDebug
		return ret
	}
	return *o.ObjDebug
}

// GetObjDebugOk returns a tuple with the ObjDebug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignbulksendGetObjectV2Response) GetObjDebugOk() (*CommonResponseObjDebug, bool) {
	if o == nil || IsNil(o.ObjDebug) {
		return nil, false
	}
	return o.ObjDebug, true
}

// HasObjDebug returns a boolean if a field has been set.
func (o *EzsignbulksendGetObjectV2Response) HasObjDebug() bool {
	if o != nil && !IsNil(o.ObjDebug) {
		return true
	}

	return false
}

// SetObjDebug gets a reference to the given CommonResponseObjDebug and assigns it to the ObjDebug field.
func (o *EzsignbulksendGetObjectV2Response) SetObjDebug(v CommonResponseObjDebug) {
	o.ObjDebug = &v
}

// GetMPayload returns the MPayload field value
func (o *EzsignbulksendGetObjectV2Response) GetMPayload() EzsignbulksendGetObjectV2ResponseMPayload {
	if o == nil {
		var ret EzsignbulksendGetObjectV2ResponseMPayload
		return ret
	}

	return o.MPayload
}

// GetMPayloadOk returns a tuple with the MPayload field value
// and a boolean to check if the value has been set.
func (o *EzsignbulksendGetObjectV2Response) GetMPayloadOk() (*EzsignbulksendGetObjectV2ResponseMPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MPayload, true
}

// SetMPayload sets field value
func (o *EzsignbulksendGetObjectV2Response) SetMPayload(v EzsignbulksendGetObjectV2ResponseMPayload) {
	o.MPayload = v
}

func (o EzsignbulksendGetObjectV2Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignbulksendGetObjectV2Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objDebugPayload"] = o.ObjDebugPayload
	if !IsNil(o.ObjDebug) {
		toSerialize["objDebug"] = o.ObjDebug
	}
	toSerialize["mPayload"] = o.MPayload
	return toSerialize, nil
}

func (o *EzsignbulksendGetObjectV2Response) UnmarshalJSON(data []byte) (err error) {
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

	varEzsignbulksendGetObjectV2Response := _EzsignbulksendGetObjectV2Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsignbulksendGetObjectV2Response)

	if err != nil {
		return err
	}

	*o = EzsignbulksendGetObjectV2Response(varEzsignbulksendGetObjectV2Response)

	return err
}

type NullableEzsignbulksendGetObjectV2Response struct {
	value *EzsignbulksendGetObjectV2Response
	isSet bool
}

func (v NullableEzsignbulksendGetObjectV2Response) Get() *EzsignbulksendGetObjectV2Response {
	return v.value
}

func (v *NullableEzsignbulksendGetObjectV2Response) Set(val *EzsignbulksendGetObjectV2Response) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignbulksendGetObjectV2Response) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignbulksendGetObjectV2Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignbulksendGetObjectV2Response(val *EzsignbulksendGetObjectV2Response) *NullableEzsignbulksendGetObjectV2Response {
	return &NullableEzsignbulksendGetObjectV2Response{value: val, isSet: true}
}

func (v NullableEzsignbulksendGetObjectV2Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignbulksendGetObjectV2Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


