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

// checks if the EzmaxinvoicingGetProvisionalV1Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzmaxinvoicingGetProvisionalV1Response{}

// EzmaxinvoicingGetProvisionalV1Response Response for GET /1/object/ezmaxinvoicing/getProvisional
type EzmaxinvoicingGetProvisionalV1Response struct {
	ObjDebugPayload CommonResponseObjDebugPayload `json:"objDebugPayload"`
	ObjDebug *CommonResponseObjDebug `json:"objDebug,omitempty"`
	MPayload EzmaxinvoicingGetProvisionalV1ResponseMPayload `json:"mPayload"`
}

type _EzmaxinvoicingGetProvisionalV1Response EzmaxinvoicingGetProvisionalV1Response

// NewEzmaxinvoicingGetProvisionalV1Response instantiates a new EzmaxinvoicingGetProvisionalV1Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzmaxinvoicingGetProvisionalV1Response(objDebugPayload CommonResponseObjDebugPayload, mPayload EzmaxinvoicingGetProvisionalV1ResponseMPayload) *EzmaxinvoicingGetProvisionalV1Response {
	this := EzmaxinvoicingGetProvisionalV1Response{}
	this.ObjDebugPayload = objDebugPayload
	this.MPayload = mPayload
	return &this
}

// NewEzmaxinvoicingGetProvisionalV1ResponseWithDefaults instantiates a new EzmaxinvoicingGetProvisionalV1Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzmaxinvoicingGetProvisionalV1ResponseWithDefaults() *EzmaxinvoicingGetProvisionalV1Response {
	this := EzmaxinvoicingGetProvisionalV1Response{}
	return &this
}

// GetObjDebugPayload returns the ObjDebugPayload field value
func (o *EzmaxinvoicingGetProvisionalV1Response) GetObjDebugPayload() CommonResponseObjDebugPayload {
	if o == nil {
		var ret CommonResponseObjDebugPayload
		return ret
	}

	return o.ObjDebugPayload
}

// GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingGetProvisionalV1Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjDebugPayload, true
}

// SetObjDebugPayload sets field value
func (o *EzmaxinvoicingGetProvisionalV1Response) SetObjDebugPayload(v CommonResponseObjDebugPayload) {
	o.ObjDebugPayload = v
}

// GetObjDebug returns the ObjDebug field value if set, zero value otherwise.
func (o *EzmaxinvoicingGetProvisionalV1Response) GetObjDebug() CommonResponseObjDebug {
	if o == nil || IsNil(o.ObjDebug) {
		var ret CommonResponseObjDebug
		return ret
	}
	return *o.ObjDebug
}

// GetObjDebugOk returns a tuple with the ObjDebug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingGetProvisionalV1Response) GetObjDebugOk() (*CommonResponseObjDebug, bool) {
	if o == nil || IsNil(o.ObjDebug) {
		return nil, false
	}
	return o.ObjDebug, true
}

// HasObjDebug returns a boolean if a field has been set.
func (o *EzmaxinvoicingGetProvisionalV1Response) HasObjDebug() bool {
	if o != nil && !IsNil(o.ObjDebug) {
		return true
	}

	return false
}

// SetObjDebug gets a reference to the given CommonResponseObjDebug and assigns it to the ObjDebug field.
func (o *EzmaxinvoicingGetProvisionalV1Response) SetObjDebug(v CommonResponseObjDebug) {
	o.ObjDebug = &v
}

// GetMPayload returns the MPayload field value
func (o *EzmaxinvoicingGetProvisionalV1Response) GetMPayload() EzmaxinvoicingGetProvisionalV1ResponseMPayload {
	if o == nil {
		var ret EzmaxinvoicingGetProvisionalV1ResponseMPayload
		return ret
	}

	return o.MPayload
}

// GetMPayloadOk returns a tuple with the MPayload field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingGetProvisionalV1Response) GetMPayloadOk() (*EzmaxinvoicingGetProvisionalV1ResponseMPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MPayload, true
}

// SetMPayload sets field value
func (o *EzmaxinvoicingGetProvisionalV1Response) SetMPayload(v EzmaxinvoicingGetProvisionalV1ResponseMPayload) {
	o.MPayload = v
}

func (o EzmaxinvoicingGetProvisionalV1Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzmaxinvoicingGetProvisionalV1Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objDebugPayload"] = o.ObjDebugPayload
	if !IsNil(o.ObjDebug) {
		toSerialize["objDebug"] = o.ObjDebug
	}
	toSerialize["mPayload"] = o.MPayload
	return toSerialize, nil
}

func (o *EzmaxinvoicingGetProvisionalV1Response) UnmarshalJSON(data []byte) (err error) {
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

	varEzmaxinvoicingGetProvisionalV1Response := _EzmaxinvoicingGetProvisionalV1Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzmaxinvoicingGetProvisionalV1Response)

	if err != nil {
		return err
	}

	*o = EzmaxinvoicingGetProvisionalV1Response(varEzmaxinvoicingGetProvisionalV1Response)

	return err
}

type NullableEzmaxinvoicingGetProvisionalV1Response struct {
	value *EzmaxinvoicingGetProvisionalV1Response
	isSet bool
}

func (v NullableEzmaxinvoicingGetProvisionalV1Response) Get() *EzmaxinvoicingGetProvisionalV1Response {
	return v.value
}

func (v *NullableEzmaxinvoicingGetProvisionalV1Response) Set(val *EzmaxinvoicingGetProvisionalV1Response) {
	v.value = val
	v.isSet = true
}

func (v NullableEzmaxinvoicingGetProvisionalV1Response) IsSet() bool {
	return v.isSet
}

func (v *NullableEzmaxinvoicingGetProvisionalV1Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzmaxinvoicingGetProvisionalV1Response(val *EzmaxinvoicingGetProvisionalV1Response) *NullableEzmaxinvoicingGetProvisionalV1Response {
	return &NullableEzmaxinvoicingGetProvisionalV1Response{value: val, isSet: true}
}

func (v NullableEzmaxinvoicingGetProvisionalV1Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzmaxinvoicingGetProvisionalV1Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


