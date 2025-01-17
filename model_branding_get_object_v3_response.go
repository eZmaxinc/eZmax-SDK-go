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

// checks if the BrandingGetObjectV3Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BrandingGetObjectV3Response{}

// BrandingGetObjectV3Response Response for GET /3/object/branding/{pkiBrandingID}
type BrandingGetObjectV3Response struct {
	ObjDebugPayload CommonResponseObjDebugPayload `json:"objDebugPayload"`
	ObjDebug *CommonResponseObjDebug `json:"objDebug,omitempty"`
	MPayload BrandingGetObjectV3ResponseMPayload `json:"mPayload"`
}

type _BrandingGetObjectV3Response BrandingGetObjectV3Response

// NewBrandingGetObjectV3Response instantiates a new BrandingGetObjectV3Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBrandingGetObjectV3Response(objDebugPayload CommonResponseObjDebugPayload, mPayload BrandingGetObjectV3ResponseMPayload) *BrandingGetObjectV3Response {
	this := BrandingGetObjectV3Response{}
	this.ObjDebugPayload = objDebugPayload
	this.MPayload = mPayload
	return &this
}

// NewBrandingGetObjectV3ResponseWithDefaults instantiates a new BrandingGetObjectV3Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBrandingGetObjectV3ResponseWithDefaults() *BrandingGetObjectV3Response {
	this := BrandingGetObjectV3Response{}
	return &this
}

// GetObjDebugPayload returns the ObjDebugPayload field value
func (o *BrandingGetObjectV3Response) GetObjDebugPayload() CommonResponseObjDebugPayload {
	if o == nil {
		var ret CommonResponseObjDebugPayload
		return ret
	}

	return o.ObjDebugPayload
}

// GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field value
// and a boolean to check if the value has been set.
func (o *BrandingGetObjectV3Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjDebugPayload, true
}

// SetObjDebugPayload sets field value
func (o *BrandingGetObjectV3Response) SetObjDebugPayload(v CommonResponseObjDebugPayload) {
	o.ObjDebugPayload = v
}

// GetObjDebug returns the ObjDebug field value if set, zero value otherwise.
func (o *BrandingGetObjectV3Response) GetObjDebug() CommonResponseObjDebug {
	if o == nil || IsNil(o.ObjDebug) {
		var ret CommonResponseObjDebug
		return ret
	}
	return *o.ObjDebug
}

// GetObjDebugOk returns a tuple with the ObjDebug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandingGetObjectV3Response) GetObjDebugOk() (*CommonResponseObjDebug, bool) {
	if o == nil || IsNil(o.ObjDebug) {
		return nil, false
	}
	return o.ObjDebug, true
}

// HasObjDebug returns a boolean if a field has been set.
func (o *BrandingGetObjectV3Response) HasObjDebug() bool {
	if o != nil && !IsNil(o.ObjDebug) {
		return true
	}

	return false
}

// SetObjDebug gets a reference to the given CommonResponseObjDebug and assigns it to the ObjDebug field.
func (o *BrandingGetObjectV3Response) SetObjDebug(v CommonResponseObjDebug) {
	o.ObjDebug = &v
}

// GetMPayload returns the MPayload field value
func (o *BrandingGetObjectV3Response) GetMPayload() BrandingGetObjectV3ResponseMPayload {
	if o == nil {
		var ret BrandingGetObjectV3ResponseMPayload
		return ret
	}

	return o.MPayload
}

// GetMPayloadOk returns a tuple with the MPayload field value
// and a boolean to check if the value has been set.
func (o *BrandingGetObjectV3Response) GetMPayloadOk() (*BrandingGetObjectV3ResponseMPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MPayload, true
}

// SetMPayload sets field value
func (o *BrandingGetObjectV3Response) SetMPayload(v BrandingGetObjectV3ResponseMPayload) {
	o.MPayload = v
}

func (o BrandingGetObjectV3Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BrandingGetObjectV3Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objDebugPayload"] = o.ObjDebugPayload
	if !IsNil(o.ObjDebug) {
		toSerialize["objDebug"] = o.ObjDebug
	}
	toSerialize["mPayload"] = o.MPayload
	return toSerialize, nil
}

func (o *BrandingGetObjectV3Response) UnmarshalJSON(data []byte) (err error) {
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

	varBrandingGetObjectV3Response := _BrandingGetObjectV3Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBrandingGetObjectV3Response)

	if err != nil {
		return err
	}

	*o = BrandingGetObjectV3Response(varBrandingGetObjectV3Response)

	return err
}

type NullableBrandingGetObjectV3Response struct {
	value *BrandingGetObjectV3Response
	isSet bool
}

func (v NullableBrandingGetObjectV3Response) Get() *BrandingGetObjectV3Response {
	return v.value
}

func (v *NullableBrandingGetObjectV3Response) Set(val *BrandingGetObjectV3Response) {
	v.value = val
	v.isSet = true
}

func (v NullableBrandingGetObjectV3Response) IsSet() bool {
	return v.isSet
}

func (v *NullableBrandingGetObjectV3Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBrandingGetObjectV3Response(val *BrandingGetObjectV3Response) *NullableBrandingGetObjectV3Response {
	return &NullableBrandingGetObjectV3Response{value: val, isSet: true}
}

func (v NullableBrandingGetObjectV3Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBrandingGetObjectV3Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


