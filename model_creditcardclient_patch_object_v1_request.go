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

// checks if the CreditcardclientPatchObjectV1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreditcardclientPatchObjectV1Request{}

// CreditcardclientPatchObjectV1Request Request for PATCH /1/object/creditcardclient/{pkiCreditcardclientID}
type CreditcardclientPatchObjectV1Request struct {
	ObjCreditcardclient CreditcardclientRequestPatch `json:"objCreditcardclient"`
}

type _CreditcardclientPatchObjectV1Request CreditcardclientPatchObjectV1Request

// NewCreditcardclientPatchObjectV1Request instantiates a new CreditcardclientPatchObjectV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditcardclientPatchObjectV1Request(objCreditcardclient CreditcardclientRequestPatch) *CreditcardclientPatchObjectV1Request {
	this := CreditcardclientPatchObjectV1Request{}
	this.ObjCreditcardclient = objCreditcardclient
	return &this
}

// NewCreditcardclientPatchObjectV1RequestWithDefaults instantiates a new CreditcardclientPatchObjectV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditcardclientPatchObjectV1RequestWithDefaults() *CreditcardclientPatchObjectV1Request {
	this := CreditcardclientPatchObjectV1Request{}
	return &this
}

// GetObjCreditcardclient returns the ObjCreditcardclient field value
func (o *CreditcardclientPatchObjectV1Request) GetObjCreditcardclient() CreditcardclientRequestPatch {
	if o == nil {
		var ret CreditcardclientRequestPatch
		return ret
	}

	return o.ObjCreditcardclient
}

// GetObjCreditcardclientOk returns a tuple with the ObjCreditcardclient field value
// and a boolean to check if the value has been set.
func (o *CreditcardclientPatchObjectV1Request) GetObjCreditcardclientOk() (*CreditcardclientRequestPatch, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjCreditcardclient, true
}

// SetObjCreditcardclient sets field value
func (o *CreditcardclientPatchObjectV1Request) SetObjCreditcardclient(v CreditcardclientRequestPatch) {
	o.ObjCreditcardclient = v
}

func (o CreditcardclientPatchObjectV1Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreditcardclientPatchObjectV1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objCreditcardclient"] = o.ObjCreditcardclient
	return toSerialize, nil
}

func (o *CreditcardclientPatchObjectV1Request) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"objCreditcardclient",
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

	varCreditcardclientPatchObjectV1Request := _CreditcardclientPatchObjectV1Request{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreditcardclientPatchObjectV1Request)

	if err != nil {
		return err
	}

	*o = CreditcardclientPatchObjectV1Request(varCreditcardclientPatchObjectV1Request)

	return err
}

type NullableCreditcardclientPatchObjectV1Request struct {
	value *CreditcardclientPatchObjectV1Request
	isSet bool
}

func (v NullableCreditcardclientPatchObjectV1Request) Get() *CreditcardclientPatchObjectV1Request {
	return v.value
}

func (v *NullableCreditcardclientPatchObjectV1Request) Set(val *CreditcardclientPatchObjectV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditcardclientPatchObjectV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditcardclientPatchObjectV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditcardclientPatchObjectV1Request(val *CreditcardclientPatchObjectV1Request) *NullableCreditcardclientPatchObjectV1Request {
	return &NullableCreditcardclientPatchObjectV1Request{value: val, isSet: true}
}

func (v NullableCreditcardclientPatchObjectV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditcardclientPatchObjectV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


