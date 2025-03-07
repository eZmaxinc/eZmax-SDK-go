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

// checks if the CreditcardmerchantGetListV1Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreditcardmerchantGetListV1Response{}

// CreditcardmerchantGetListV1Response Response for GET /1/object/creditcardmerchant/getList
type CreditcardmerchantGetListV1Response struct {
	ObjDebugPayload CommonResponseObjDebugPayloadGetList `json:"objDebugPayload"`
	ObjDebug *CommonResponseObjDebug `json:"objDebug,omitempty"`
	MPayload CreditcardmerchantGetListV1ResponseMPayload `json:"mPayload"`
}

type _CreditcardmerchantGetListV1Response CreditcardmerchantGetListV1Response

// NewCreditcardmerchantGetListV1Response instantiates a new CreditcardmerchantGetListV1Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditcardmerchantGetListV1Response(objDebugPayload CommonResponseObjDebugPayloadGetList, mPayload CreditcardmerchantGetListV1ResponseMPayload) *CreditcardmerchantGetListV1Response {
	this := CreditcardmerchantGetListV1Response{}
	this.ObjDebugPayload = objDebugPayload
	this.MPayload = mPayload
	return &this
}

// NewCreditcardmerchantGetListV1ResponseWithDefaults instantiates a new CreditcardmerchantGetListV1Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditcardmerchantGetListV1ResponseWithDefaults() *CreditcardmerchantGetListV1Response {
	this := CreditcardmerchantGetListV1Response{}
	return &this
}

// GetObjDebugPayload returns the ObjDebugPayload field value
func (o *CreditcardmerchantGetListV1Response) GetObjDebugPayload() CommonResponseObjDebugPayloadGetList {
	if o == nil {
		var ret CommonResponseObjDebugPayloadGetList
		return ret
	}

	return o.ObjDebugPayload
}

// GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field value
// and a boolean to check if the value has been set.
func (o *CreditcardmerchantGetListV1Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayloadGetList, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjDebugPayload, true
}

// SetObjDebugPayload sets field value
func (o *CreditcardmerchantGetListV1Response) SetObjDebugPayload(v CommonResponseObjDebugPayloadGetList) {
	o.ObjDebugPayload = v
}

// GetObjDebug returns the ObjDebug field value if set, zero value otherwise.
func (o *CreditcardmerchantGetListV1Response) GetObjDebug() CommonResponseObjDebug {
	if o == nil || IsNil(o.ObjDebug) {
		var ret CommonResponseObjDebug
		return ret
	}
	return *o.ObjDebug
}

// GetObjDebugOk returns a tuple with the ObjDebug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditcardmerchantGetListV1Response) GetObjDebugOk() (*CommonResponseObjDebug, bool) {
	if o == nil || IsNil(o.ObjDebug) {
		return nil, false
	}
	return o.ObjDebug, true
}

// HasObjDebug returns a boolean if a field has been set.
func (o *CreditcardmerchantGetListV1Response) HasObjDebug() bool {
	if o != nil && !IsNil(o.ObjDebug) {
		return true
	}

	return false
}

// SetObjDebug gets a reference to the given CommonResponseObjDebug and assigns it to the ObjDebug field.
func (o *CreditcardmerchantGetListV1Response) SetObjDebug(v CommonResponseObjDebug) {
	o.ObjDebug = &v
}

// GetMPayload returns the MPayload field value
func (o *CreditcardmerchantGetListV1Response) GetMPayload() CreditcardmerchantGetListV1ResponseMPayload {
	if o == nil {
		var ret CreditcardmerchantGetListV1ResponseMPayload
		return ret
	}

	return o.MPayload
}

// GetMPayloadOk returns a tuple with the MPayload field value
// and a boolean to check if the value has been set.
func (o *CreditcardmerchantGetListV1Response) GetMPayloadOk() (*CreditcardmerchantGetListV1ResponseMPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MPayload, true
}

// SetMPayload sets field value
func (o *CreditcardmerchantGetListV1Response) SetMPayload(v CreditcardmerchantGetListV1ResponseMPayload) {
	o.MPayload = v
}

func (o CreditcardmerchantGetListV1Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreditcardmerchantGetListV1Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objDebugPayload"] = o.ObjDebugPayload
	if !IsNil(o.ObjDebug) {
		toSerialize["objDebug"] = o.ObjDebug
	}
	toSerialize["mPayload"] = o.MPayload
	return toSerialize, nil
}

func (o *CreditcardmerchantGetListV1Response) UnmarshalJSON(data []byte) (err error) {
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

	varCreditcardmerchantGetListV1Response := _CreditcardmerchantGetListV1Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreditcardmerchantGetListV1Response)

	if err != nil {
		return err
	}

	*o = CreditcardmerchantGetListV1Response(varCreditcardmerchantGetListV1Response)

	return err
}

type NullableCreditcardmerchantGetListV1Response struct {
	value *CreditcardmerchantGetListV1Response
	isSet bool
}

func (v NullableCreditcardmerchantGetListV1Response) Get() *CreditcardmerchantGetListV1Response {
	return v.value
}

func (v *NullableCreditcardmerchantGetListV1Response) Set(val *CreditcardmerchantGetListV1Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditcardmerchantGetListV1Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditcardmerchantGetListV1Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditcardmerchantGetListV1Response(val *CreditcardmerchantGetListV1Response) *NullableCreditcardmerchantGetListV1Response {
	return &NullableCreditcardmerchantGetListV1Response{value: val, isSet: true}
}

func (v NullableCreditcardmerchantGetListV1Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditcardmerchantGetListV1Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


