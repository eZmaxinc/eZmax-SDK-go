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

// checks if the EzmaxcasePatchObjectV1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzmaxcasePatchObjectV1Request{}

// EzmaxcasePatchObjectV1Request Request for PATCH /1/object/ezmaxcase/{pkiEzmaxcaseID}
type EzmaxcasePatchObjectV1Request struct {
	ObjEzmaxcase EzmaxcaseRequestPatch `json:"objEzmaxcase"`
}

type _EzmaxcasePatchObjectV1Request EzmaxcasePatchObjectV1Request

// NewEzmaxcasePatchObjectV1Request instantiates a new EzmaxcasePatchObjectV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzmaxcasePatchObjectV1Request(objEzmaxcase EzmaxcaseRequestPatch) *EzmaxcasePatchObjectV1Request {
	this := EzmaxcasePatchObjectV1Request{}
	this.ObjEzmaxcase = objEzmaxcase
	return &this
}

// NewEzmaxcasePatchObjectV1RequestWithDefaults instantiates a new EzmaxcasePatchObjectV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzmaxcasePatchObjectV1RequestWithDefaults() *EzmaxcasePatchObjectV1Request {
	this := EzmaxcasePatchObjectV1Request{}
	return &this
}

// GetObjEzmaxcase returns the ObjEzmaxcase field value
func (o *EzmaxcasePatchObjectV1Request) GetObjEzmaxcase() EzmaxcaseRequestPatch {
	if o == nil {
		var ret EzmaxcaseRequestPatch
		return ret
	}

	return o.ObjEzmaxcase
}

// GetObjEzmaxcaseOk returns a tuple with the ObjEzmaxcase field value
// and a boolean to check if the value has been set.
func (o *EzmaxcasePatchObjectV1Request) GetObjEzmaxcaseOk() (*EzmaxcaseRequestPatch, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjEzmaxcase, true
}

// SetObjEzmaxcase sets field value
func (o *EzmaxcasePatchObjectV1Request) SetObjEzmaxcase(v EzmaxcaseRequestPatch) {
	o.ObjEzmaxcase = v
}

func (o EzmaxcasePatchObjectV1Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzmaxcasePatchObjectV1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objEzmaxcase"] = o.ObjEzmaxcase
	return toSerialize, nil
}

func (o *EzmaxcasePatchObjectV1Request) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"objEzmaxcase",
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

	varEzmaxcasePatchObjectV1Request := _EzmaxcasePatchObjectV1Request{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzmaxcasePatchObjectV1Request)

	if err != nil {
		return err
	}

	*o = EzmaxcasePatchObjectV1Request(varEzmaxcasePatchObjectV1Request)

	return err
}

type NullableEzmaxcasePatchObjectV1Request struct {
	value *EzmaxcasePatchObjectV1Request
	isSet bool
}

func (v NullableEzmaxcasePatchObjectV1Request) Get() *EzmaxcasePatchObjectV1Request {
	return v.value
}

func (v *NullableEzmaxcasePatchObjectV1Request) Set(val *EzmaxcasePatchObjectV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableEzmaxcasePatchObjectV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableEzmaxcasePatchObjectV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzmaxcasePatchObjectV1Request(val *EzmaxcasePatchObjectV1Request) *NullableEzmaxcasePatchObjectV1Request {
	return &NullableEzmaxcasePatchObjectV1Request{value: val, isSet: true}
}

func (v NullableEzmaxcasePatchObjectV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzmaxcasePatchObjectV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


