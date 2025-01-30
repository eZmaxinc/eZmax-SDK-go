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

// checks if the EzsignformfieldgroupEditObjectV1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignformfieldgroupEditObjectV1Request{}

// EzsignformfieldgroupEditObjectV1Request Request for PUT /1/object/ezsignformfieldgroup/{pkiEzsignfoldersignerassociationID}
type EzsignformfieldgroupEditObjectV1Request struct {
	ObjEzsignformfieldgroup EzsignformfieldgroupRequestCompound `json:"objEzsignformfieldgroup"`
}

type _EzsignformfieldgroupEditObjectV1Request EzsignformfieldgroupEditObjectV1Request

// NewEzsignformfieldgroupEditObjectV1Request instantiates a new EzsignformfieldgroupEditObjectV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignformfieldgroupEditObjectV1Request(objEzsignformfieldgroup EzsignformfieldgroupRequestCompound) *EzsignformfieldgroupEditObjectV1Request {
	this := EzsignformfieldgroupEditObjectV1Request{}
	this.ObjEzsignformfieldgroup = objEzsignformfieldgroup
	return &this
}

// NewEzsignformfieldgroupEditObjectV1RequestWithDefaults instantiates a new EzsignformfieldgroupEditObjectV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignformfieldgroupEditObjectV1RequestWithDefaults() *EzsignformfieldgroupEditObjectV1Request {
	this := EzsignformfieldgroupEditObjectV1Request{}
	return &this
}

// GetObjEzsignformfieldgroup returns the ObjEzsignformfieldgroup field value
func (o *EzsignformfieldgroupEditObjectV1Request) GetObjEzsignformfieldgroup() EzsignformfieldgroupRequestCompound {
	if o == nil {
		var ret EzsignformfieldgroupRequestCompound
		return ret
	}

	return o.ObjEzsignformfieldgroup
}

// GetObjEzsignformfieldgroupOk returns a tuple with the ObjEzsignformfieldgroup field value
// and a boolean to check if the value has been set.
func (o *EzsignformfieldgroupEditObjectV1Request) GetObjEzsignformfieldgroupOk() (*EzsignformfieldgroupRequestCompound, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjEzsignformfieldgroup, true
}

// SetObjEzsignformfieldgroup sets field value
func (o *EzsignformfieldgroupEditObjectV1Request) SetObjEzsignformfieldgroup(v EzsignformfieldgroupRequestCompound) {
	o.ObjEzsignformfieldgroup = v
}

func (o EzsignformfieldgroupEditObjectV1Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignformfieldgroupEditObjectV1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objEzsignformfieldgroup"] = o.ObjEzsignformfieldgroup
	return toSerialize, nil
}

func (o *EzsignformfieldgroupEditObjectV1Request) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"objEzsignformfieldgroup",
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

	varEzsignformfieldgroupEditObjectV1Request := _EzsignformfieldgroupEditObjectV1Request{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsignformfieldgroupEditObjectV1Request)

	if err != nil {
		return err
	}

	*o = EzsignformfieldgroupEditObjectV1Request(varEzsignformfieldgroupEditObjectV1Request)

	return err
}

type NullableEzsignformfieldgroupEditObjectV1Request struct {
	value *EzsignformfieldgroupEditObjectV1Request
	isSet bool
}

func (v NullableEzsignformfieldgroupEditObjectV1Request) Get() *EzsignformfieldgroupEditObjectV1Request {
	return v.value
}

func (v *NullableEzsignformfieldgroupEditObjectV1Request) Set(val *EzsignformfieldgroupEditObjectV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignformfieldgroupEditObjectV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignformfieldgroupEditObjectV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignformfieldgroupEditObjectV1Request(val *EzsignformfieldgroupEditObjectV1Request) *NullableEzsignformfieldgroupEditObjectV1Request {
	return &NullableEzsignformfieldgroupEditObjectV1Request{value: val, isSet: true}
}

func (v NullableEzsignformfieldgroupEditObjectV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignformfieldgroupEditObjectV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


