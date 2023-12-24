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
	"bytes"
	"fmt"
)

// checks if the EzsigntemplatesignerEditObjectV1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigntemplatesignerEditObjectV1Request{}

// EzsigntemplatesignerEditObjectV1Request Request for PUT /1/object/ezsigntemplatesigner/{pkiEzsigntemplatesignerID}
type EzsigntemplatesignerEditObjectV1Request struct {
	ObjEzsigntemplatesigner EzsigntemplatesignerRequestCompound `json:"objEzsigntemplatesigner"`
}

type _EzsigntemplatesignerEditObjectV1Request EzsigntemplatesignerEditObjectV1Request

// NewEzsigntemplatesignerEditObjectV1Request instantiates a new EzsigntemplatesignerEditObjectV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigntemplatesignerEditObjectV1Request(objEzsigntemplatesigner EzsigntemplatesignerRequestCompound) *EzsigntemplatesignerEditObjectV1Request {
	this := EzsigntemplatesignerEditObjectV1Request{}
	this.ObjEzsigntemplatesigner = objEzsigntemplatesigner
	return &this
}

// NewEzsigntemplatesignerEditObjectV1RequestWithDefaults instantiates a new EzsigntemplatesignerEditObjectV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigntemplatesignerEditObjectV1RequestWithDefaults() *EzsigntemplatesignerEditObjectV1Request {
	this := EzsigntemplatesignerEditObjectV1Request{}
	return &this
}

// GetObjEzsigntemplatesigner returns the ObjEzsigntemplatesigner field value
func (o *EzsigntemplatesignerEditObjectV1Request) GetObjEzsigntemplatesigner() EzsigntemplatesignerRequestCompound {
	if o == nil {
		var ret EzsigntemplatesignerRequestCompound
		return ret
	}

	return o.ObjEzsigntemplatesigner
}

// GetObjEzsigntemplatesignerOk returns a tuple with the ObjEzsigntemplatesigner field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatesignerEditObjectV1Request) GetObjEzsigntemplatesignerOk() (*EzsigntemplatesignerRequestCompound, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjEzsigntemplatesigner, true
}

// SetObjEzsigntemplatesigner sets field value
func (o *EzsigntemplatesignerEditObjectV1Request) SetObjEzsigntemplatesigner(v EzsigntemplatesignerRequestCompound) {
	o.ObjEzsigntemplatesigner = v
}

func (o EzsigntemplatesignerEditObjectV1Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigntemplatesignerEditObjectV1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objEzsigntemplatesigner"] = o.ObjEzsigntemplatesigner
	return toSerialize, nil
}

func (o *EzsigntemplatesignerEditObjectV1Request) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"objEzsigntemplatesigner",
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

	varEzsigntemplatesignerEditObjectV1Request := _EzsigntemplatesignerEditObjectV1Request{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsigntemplatesignerEditObjectV1Request)

	if err != nil {
		return err
	}

	*o = EzsigntemplatesignerEditObjectV1Request(varEzsigntemplatesignerEditObjectV1Request)

	return err
}

type NullableEzsigntemplatesignerEditObjectV1Request struct {
	value *EzsigntemplatesignerEditObjectV1Request
	isSet bool
}

func (v NullableEzsigntemplatesignerEditObjectV1Request) Get() *EzsigntemplatesignerEditObjectV1Request {
	return v.value
}

func (v *NullableEzsigntemplatesignerEditObjectV1Request) Set(val *EzsigntemplatesignerEditObjectV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigntemplatesignerEditObjectV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigntemplatesignerEditObjectV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigntemplatesignerEditObjectV1Request(val *EzsigntemplatesignerEditObjectV1Request) *NullableEzsigntemplatesignerEditObjectV1Request {
	return &NullableEzsigntemplatesignerEditObjectV1Request{value: val, isSet: true}
}

func (v NullableEzsigntemplatesignerEditObjectV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigntemplatesignerEditObjectV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


