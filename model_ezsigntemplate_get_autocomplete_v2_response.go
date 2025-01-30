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

// checks if the EzsigntemplateGetAutocompleteV2Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigntemplateGetAutocompleteV2Response{}

// EzsigntemplateGetAutocompleteV2Response Response for GET /2/object/ezsigntemplate/getAutocomplete
type EzsigntemplateGetAutocompleteV2Response struct {
	ObjDebugPayload CommonResponseObjDebugPayload `json:"objDebugPayload"`
	ObjDebug *CommonResponseObjDebug `json:"objDebug,omitempty"`
	MPayload EzsigntemplateGetAutocompleteV2ResponseMPayload `json:"mPayload"`
}

type _EzsigntemplateGetAutocompleteV2Response EzsigntemplateGetAutocompleteV2Response

// NewEzsigntemplateGetAutocompleteV2Response instantiates a new EzsigntemplateGetAutocompleteV2Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigntemplateGetAutocompleteV2Response(objDebugPayload CommonResponseObjDebugPayload, mPayload EzsigntemplateGetAutocompleteV2ResponseMPayload) *EzsigntemplateGetAutocompleteV2Response {
	this := EzsigntemplateGetAutocompleteV2Response{}
	this.ObjDebugPayload = objDebugPayload
	this.MPayload = mPayload
	return &this
}

// NewEzsigntemplateGetAutocompleteV2ResponseWithDefaults instantiates a new EzsigntemplateGetAutocompleteV2Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigntemplateGetAutocompleteV2ResponseWithDefaults() *EzsigntemplateGetAutocompleteV2Response {
	this := EzsigntemplateGetAutocompleteV2Response{}
	return &this
}

// GetObjDebugPayload returns the ObjDebugPayload field value
func (o *EzsigntemplateGetAutocompleteV2Response) GetObjDebugPayload() CommonResponseObjDebugPayload {
	if o == nil {
		var ret CommonResponseObjDebugPayload
		return ret
	}

	return o.ObjDebugPayload
}

// GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateGetAutocompleteV2Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjDebugPayload, true
}

// SetObjDebugPayload sets field value
func (o *EzsigntemplateGetAutocompleteV2Response) SetObjDebugPayload(v CommonResponseObjDebugPayload) {
	o.ObjDebugPayload = v
}

// GetObjDebug returns the ObjDebug field value if set, zero value otherwise.
func (o *EzsigntemplateGetAutocompleteV2Response) GetObjDebug() CommonResponseObjDebug {
	if o == nil || IsNil(o.ObjDebug) {
		var ret CommonResponseObjDebug
		return ret
	}
	return *o.ObjDebug
}

// GetObjDebugOk returns a tuple with the ObjDebug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateGetAutocompleteV2Response) GetObjDebugOk() (*CommonResponseObjDebug, bool) {
	if o == nil || IsNil(o.ObjDebug) {
		return nil, false
	}
	return o.ObjDebug, true
}

// HasObjDebug returns a boolean if a field has been set.
func (o *EzsigntemplateGetAutocompleteV2Response) HasObjDebug() bool {
	if o != nil && !IsNil(o.ObjDebug) {
		return true
	}

	return false
}

// SetObjDebug gets a reference to the given CommonResponseObjDebug and assigns it to the ObjDebug field.
func (o *EzsigntemplateGetAutocompleteV2Response) SetObjDebug(v CommonResponseObjDebug) {
	o.ObjDebug = &v
}

// GetMPayload returns the MPayload field value
func (o *EzsigntemplateGetAutocompleteV2Response) GetMPayload() EzsigntemplateGetAutocompleteV2ResponseMPayload {
	if o == nil {
		var ret EzsigntemplateGetAutocompleteV2ResponseMPayload
		return ret
	}

	return o.MPayload
}

// GetMPayloadOk returns a tuple with the MPayload field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateGetAutocompleteV2Response) GetMPayloadOk() (*EzsigntemplateGetAutocompleteV2ResponseMPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MPayload, true
}

// SetMPayload sets field value
func (o *EzsigntemplateGetAutocompleteV2Response) SetMPayload(v EzsigntemplateGetAutocompleteV2ResponseMPayload) {
	o.MPayload = v
}

func (o EzsigntemplateGetAutocompleteV2Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigntemplateGetAutocompleteV2Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objDebugPayload"] = o.ObjDebugPayload
	if !IsNil(o.ObjDebug) {
		toSerialize["objDebug"] = o.ObjDebug
	}
	toSerialize["mPayload"] = o.MPayload
	return toSerialize, nil
}

func (o *EzsigntemplateGetAutocompleteV2Response) UnmarshalJSON(data []byte) (err error) {
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

	varEzsigntemplateGetAutocompleteV2Response := _EzsigntemplateGetAutocompleteV2Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsigntemplateGetAutocompleteV2Response)

	if err != nil {
		return err
	}

	*o = EzsigntemplateGetAutocompleteV2Response(varEzsigntemplateGetAutocompleteV2Response)

	return err
}

type NullableEzsigntemplateGetAutocompleteV2Response struct {
	value *EzsigntemplateGetAutocompleteV2Response
	isSet bool
}

func (v NullableEzsigntemplateGetAutocompleteV2Response) Get() *EzsigntemplateGetAutocompleteV2Response {
	return v.value
}

func (v *NullableEzsigntemplateGetAutocompleteV2Response) Set(val *EzsigntemplateGetAutocompleteV2Response) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigntemplateGetAutocompleteV2Response) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigntemplateGetAutocompleteV2Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigntemplateGetAutocompleteV2Response(val *EzsigntemplateGetAutocompleteV2Response) *NullableEzsigntemplateGetAutocompleteV2Response {
	return &NullableEzsigntemplateGetAutocompleteV2Response{value: val, isSet: true}
}

func (v NullableEzsigntemplateGetAutocompleteV2Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigntemplateGetAutocompleteV2Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


