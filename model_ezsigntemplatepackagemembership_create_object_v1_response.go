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

// checks if the EzsigntemplatepackagemembershipCreateObjectV1Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigntemplatepackagemembershipCreateObjectV1Response{}

// EzsigntemplatepackagemembershipCreateObjectV1Response Response for POST /1/object/ezsigntemplatepackagemembership
type EzsigntemplatepackagemembershipCreateObjectV1Response struct {
	ObjDebugPayload CommonResponseObjDebugPayload `json:"objDebugPayload"`
	ObjDebug *CommonResponseObjDebug `json:"objDebug,omitempty"`
	MPayload EzsigntemplatepackagemembershipCreateObjectV1ResponseMPayload `json:"mPayload"`
}

type _EzsigntemplatepackagemembershipCreateObjectV1Response EzsigntemplatepackagemembershipCreateObjectV1Response

// NewEzsigntemplatepackagemembershipCreateObjectV1Response instantiates a new EzsigntemplatepackagemembershipCreateObjectV1Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigntemplatepackagemembershipCreateObjectV1Response(objDebugPayload CommonResponseObjDebugPayload, mPayload EzsigntemplatepackagemembershipCreateObjectV1ResponseMPayload) *EzsigntemplatepackagemembershipCreateObjectV1Response {
	this := EzsigntemplatepackagemembershipCreateObjectV1Response{}
	this.ObjDebugPayload = objDebugPayload
	this.MPayload = mPayload
	return &this
}

// NewEzsigntemplatepackagemembershipCreateObjectV1ResponseWithDefaults instantiates a new EzsigntemplatepackagemembershipCreateObjectV1Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigntemplatepackagemembershipCreateObjectV1ResponseWithDefaults() *EzsigntemplatepackagemembershipCreateObjectV1Response {
	this := EzsigntemplatepackagemembershipCreateObjectV1Response{}
	return &this
}

// GetObjDebugPayload returns the ObjDebugPayload field value
func (o *EzsigntemplatepackagemembershipCreateObjectV1Response) GetObjDebugPayload() CommonResponseObjDebugPayload {
	if o == nil {
		var ret CommonResponseObjDebugPayload
		return ret
	}

	return o.ObjDebugPayload
}

// GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepackagemembershipCreateObjectV1Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjDebugPayload, true
}

// SetObjDebugPayload sets field value
func (o *EzsigntemplatepackagemembershipCreateObjectV1Response) SetObjDebugPayload(v CommonResponseObjDebugPayload) {
	o.ObjDebugPayload = v
}

// GetObjDebug returns the ObjDebug field value if set, zero value otherwise.
func (o *EzsigntemplatepackagemembershipCreateObjectV1Response) GetObjDebug() CommonResponseObjDebug {
	if o == nil || IsNil(o.ObjDebug) {
		var ret CommonResponseObjDebug
		return ret
	}
	return *o.ObjDebug
}

// GetObjDebugOk returns a tuple with the ObjDebug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepackagemembershipCreateObjectV1Response) GetObjDebugOk() (*CommonResponseObjDebug, bool) {
	if o == nil || IsNil(o.ObjDebug) {
		return nil, false
	}
	return o.ObjDebug, true
}

// HasObjDebug returns a boolean if a field has been set.
func (o *EzsigntemplatepackagemembershipCreateObjectV1Response) HasObjDebug() bool {
	if o != nil && !IsNil(o.ObjDebug) {
		return true
	}

	return false
}

// SetObjDebug gets a reference to the given CommonResponseObjDebug and assigns it to the ObjDebug field.
func (o *EzsigntemplatepackagemembershipCreateObjectV1Response) SetObjDebug(v CommonResponseObjDebug) {
	o.ObjDebug = &v
}

// GetMPayload returns the MPayload field value
func (o *EzsigntemplatepackagemembershipCreateObjectV1Response) GetMPayload() EzsigntemplatepackagemembershipCreateObjectV1ResponseMPayload {
	if o == nil {
		var ret EzsigntemplatepackagemembershipCreateObjectV1ResponseMPayload
		return ret
	}

	return o.MPayload
}

// GetMPayloadOk returns a tuple with the MPayload field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepackagemembershipCreateObjectV1Response) GetMPayloadOk() (*EzsigntemplatepackagemembershipCreateObjectV1ResponseMPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MPayload, true
}

// SetMPayload sets field value
func (o *EzsigntemplatepackagemembershipCreateObjectV1Response) SetMPayload(v EzsigntemplatepackagemembershipCreateObjectV1ResponseMPayload) {
	o.MPayload = v
}

func (o EzsigntemplatepackagemembershipCreateObjectV1Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigntemplatepackagemembershipCreateObjectV1Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objDebugPayload"] = o.ObjDebugPayload
	if !IsNil(o.ObjDebug) {
		toSerialize["objDebug"] = o.ObjDebug
	}
	toSerialize["mPayload"] = o.MPayload
	return toSerialize, nil
}

func (o *EzsigntemplatepackagemembershipCreateObjectV1Response) UnmarshalJSON(data []byte) (err error) {
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

	varEzsigntemplatepackagemembershipCreateObjectV1Response := _EzsigntemplatepackagemembershipCreateObjectV1Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsigntemplatepackagemembershipCreateObjectV1Response)

	if err != nil {
		return err
	}

	*o = EzsigntemplatepackagemembershipCreateObjectV1Response(varEzsigntemplatepackagemembershipCreateObjectV1Response)

	return err
}

type NullableEzsigntemplatepackagemembershipCreateObjectV1Response struct {
	value *EzsigntemplatepackagemembershipCreateObjectV1Response
	isSet bool
}

func (v NullableEzsigntemplatepackagemembershipCreateObjectV1Response) Get() *EzsigntemplatepackagemembershipCreateObjectV1Response {
	return v.value
}

func (v *NullableEzsigntemplatepackagemembershipCreateObjectV1Response) Set(val *EzsigntemplatepackagemembershipCreateObjectV1Response) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigntemplatepackagemembershipCreateObjectV1Response) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigntemplatepackagemembershipCreateObjectV1Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigntemplatepackagemembershipCreateObjectV1Response(val *EzsigntemplatepackagemembershipCreateObjectV1Response) *NullableEzsigntemplatepackagemembershipCreateObjectV1Response {
	return &NullableEzsigntemplatepackagemembershipCreateObjectV1Response{value: val, isSet: true}
}

func (v NullableEzsigntemplatepackagemembershipCreateObjectV1Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigntemplatepackagemembershipCreateObjectV1Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


