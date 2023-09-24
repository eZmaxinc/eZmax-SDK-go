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
)

// checks if the EzsigntemplatepackageGetAutocompleteV2Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigntemplatepackageGetAutocompleteV2Response{}

// EzsigntemplatepackageGetAutocompleteV2Response Response for GET /2/object/ezsigntemplatepackage/getAutocomplete
type EzsigntemplatepackageGetAutocompleteV2Response struct {
	ObjDebugPayload CommonResponseObjDebugPayload `json:"objDebugPayload"`
	ObjDebug *CommonResponseObjDebug `json:"objDebug,omitempty"`
	MPayload EzsigntemplatepackageGetAutocompleteV2ResponseMPayload `json:"mPayload"`
}

// NewEzsigntemplatepackageGetAutocompleteV2Response instantiates a new EzsigntemplatepackageGetAutocompleteV2Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigntemplatepackageGetAutocompleteV2Response(objDebugPayload CommonResponseObjDebugPayload, mPayload EzsigntemplatepackageGetAutocompleteV2ResponseMPayload) *EzsigntemplatepackageGetAutocompleteV2Response {
	this := EzsigntemplatepackageGetAutocompleteV2Response{}
	this.ObjDebugPayload = objDebugPayload
	this.MPayload = mPayload
	return &this
}

// NewEzsigntemplatepackageGetAutocompleteV2ResponseWithDefaults instantiates a new EzsigntemplatepackageGetAutocompleteV2Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigntemplatepackageGetAutocompleteV2ResponseWithDefaults() *EzsigntemplatepackageGetAutocompleteV2Response {
	this := EzsigntemplatepackageGetAutocompleteV2Response{}
	return &this
}

// GetObjDebugPayload returns the ObjDebugPayload field value
func (o *EzsigntemplatepackageGetAutocompleteV2Response) GetObjDebugPayload() CommonResponseObjDebugPayload {
	if o == nil {
		var ret CommonResponseObjDebugPayload
		return ret
	}

	return o.ObjDebugPayload
}

// GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepackageGetAutocompleteV2Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjDebugPayload, true
}

// SetObjDebugPayload sets field value
func (o *EzsigntemplatepackageGetAutocompleteV2Response) SetObjDebugPayload(v CommonResponseObjDebugPayload) {
	o.ObjDebugPayload = v
}

// GetObjDebug returns the ObjDebug field value if set, zero value otherwise.
func (o *EzsigntemplatepackageGetAutocompleteV2Response) GetObjDebug() CommonResponseObjDebug {
	if o == nil || IsNil(o.ObjDebug) {
		var ret CommonResponseObjDebug
		return ret
	}
	return *o.ObjDebug
}

// GetObjDebugOk returns a tuple with the ObjDebug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepackageGetAutocompleteV2Response) GetObjDebugOk() (*CommonResponseObjDebug, bool) {
	if o == nil || IsNil(o.ObjDebug) {
		return nil, false
	}
	return o.ObjDebug, true
}

// HasObjDebug returns a boolean if a field has been set.
func (o *EzsigntemplatepackageGetAutocompleteV2Response) HasObjDebug() bool {
	if o != nil && !IsNil(o.ObjDebug) {
		return true
	}

	return false
}

// SetObjDebug gets a reference to the given CommonResponseObjDebug and assigns it to the ObjDebug field.
func (o *EzsigntemplatepackageGetAutocompleteV2Response) SetObjDebug(v CommonResponseObjDebug) {
	o.ObjDebug = &v
}

// GetMPayload returns the MPayload field value
func (o *EzsigntemplatepackageGetAutocompleteV2Response) GetMPayload() EzsigntemplatepackageGetAutocompleteV2ResponseMPayload {
	if o == nil {
		var ret EzsigntemplatepackageGetAutocompleteV2ResponseMPayload
		return ret
	}

	return o.MPayload
}

// GetMPayloadOk returns a tuple with the MPayload field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepackageGetAutocompleteV2Response) GetMPayloadOk() (*EzsigntemplatepackageGetAutocompleteV2ResponseMPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MPayload, true
}

// SetMPayload sets field value
func (o *EzsigntemplatepackageGetAutocompleteV2Response) SetMPayload(v EzsigntemplatepackageGetAutocompleteV2ResponseMPayload) {
	o.MPayload = v
}

func (o EzsigntemplatepackageGetAutocompleteV2Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigntemplatepackageGetAutocompleteV2Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objDebugPayload"] = o.ObjDebugPayload
	if !IsNil(o.ObjDebug) {
		toSerialize["objDebug"] = o.ObjDebug
	}
	toSerialize["mPayload"] = o.MPayload
	return toSerialize, nil
}

type NullableEzsigntemplatepackageGetAutocompleteV2Response struct {
	value *EzsigntemplatepackageGetAutocompleteV2Response
	isSet bool
}

func (v NullableEzsigntemplatepackageGetAutocompleteV2Response) Get() *EzsigntemplatepackageGetAutocompleteV2Response {
	return v.value
}

func (v *NullableEzsigntemplatepackageGetAutocompleteV2Response) Set(val *EzsigntemplatepackageGetAutocompleteV2Response) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigntemplatepackageGetAutocompleteV2Response) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigntemplatepackageGetAutocompleteV2Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigntemplatepackageGetAutocompleteV2Response(val *EzsigntemplatepackageGetAutocompleteV2Response) *NullableEzsigntemplatepackageGetAutocompleteV2Response {
	return &NullableEzsigntemplatepackageGetAutocompleteV2Response{value: val, isSet: true}
}

func (v NullableEzsigntemplatepackageGetAutocompleteV2Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigntemplatepackageGetAutocompleteV2Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

