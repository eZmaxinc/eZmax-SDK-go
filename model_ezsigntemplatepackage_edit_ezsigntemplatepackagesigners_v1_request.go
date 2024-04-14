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

// checks if the EzsigntemplatepackageEditEzsigntemplatepackagesignersV1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigntemplatepackageEditEzsigntemplatepackagesignersV1Request{}

// EzsigntemplatepackageEditEzsigntemplatepackagesignersV1Request Request for PUT /1/object/ezsigntemplatepackage/{pkiEzsigntemplatepackageID}/editEzsigntemplatepackagesigners
type EzsigntemplatepackageEditEzsigntemplatepackagesignersV1Request struct {
	AObjEzsigntemplatepackagesigner []EzsigntemplatepackagesignerRequestCompound `json:"a_objEzsigntemplatepackagesigner"`
}

type _EzsigntemplatepackageEditEzsigntemplatepackagesignersV1Request EzsigntemplatepackageEditEzsigntemplatepackagesignersV1Request

// NewEzsigntemplatepackageEditEzsigntemplatepackagesignersV1Request instantiates a new EzsigntemplatepackageEditEzsigntemplatepackagesignersV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigntemplatepackageEditEzsigntemplatepackagesignersV1Request(aObjEzsigntemplatepackagesigner []EzsigntemplatepackagesignerRequestCompound) *EzsigntemplatepackageEditEzsigntemplatepackagesignersV1Request {
	this := EzsigntemplatepackageEditEzsigntemplatepackagesignersV1Request{}
	this.AObjEzsigntemplatepackagesigner = aObjEzsigntemplatepackagesigner
	return &this
}

// NewEzsigntemplatepackageEditEzsigntemplatepackagesignersV1RequestWithDefaults instantiates a new EzsigntemplatepackageEditEzsigntemplatepackagesignersV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigntemplatepackageEditEzsigntemplatepackagesignersV1RequestWithDefaults() *EzsigntemplatepackageEditEzsigntemplatepackagesignersV1Request {
	this := EzsigntemplatepackageEditEzsigntemplatepackagesignersV1Request{}
	return &this
}

// GetAObjEzsigntemplatepackagesigner returns the AObjEzsigntemplatepackagesigner field value
func (o *EzsigntemplatepackageEditEzsigntemplatepackagesignersV1Request) GetAObjEzsigntemplatepackagesigner() []EzsigntemplatepackagesignerRequestCompound {
	if o == nil {
		var ret []EzsigntemplatepackagesignerRequestCompound
		return ret
	}

	return o.AObjEzsigntemplatepackagesigner
}

// GetAObjEzsigntemplatepackagesignerOk returns a tuple with the AObjEzsigntemplatepackagesigner field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepackageEditEzsigntemplatepackagesignersV1Request) GetAObjEzsigntemplatepackagesignerOk() ([]EzsigntemplatepackagesignerRequestCompound, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjEzsigntemplatepackagesigner, true
}

// SetAObjEzsigntemplatepackagesigner sets field value
func (o *EzsigntemplatepackageEditEzsigntemplatepackagesignersV1Request) SetAObjEzsigntemplatepackagesigner(v []EzsigntemplatepackagesignerRequestCompound) {
	o.AObjEzsigntemplatepackagesigner = v
}

func (o EzsigntemplatepackageEditEzsigntemplatepackagesignersV1Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigntemplatepackageEditEzsigntemplatepackagesignersV1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_objEzsigntemplatepackagesigner"] = o.AObjEzsigntemplatepackagesigner
	return toSerialize, nil
}

func (o *EzsigntemplatepackageEditEzsigntemplatepackagesignersV1Request) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"a_objEzsigntemplatepackagesigner",
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

	varEzsigntemplatepackageEditEzsigntemplatepackagesignersV1Request := _EzsigntemplatepackageEditEzsigntemplatepackagesignersV1Request{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsigntemplatepackageEditEzsigntemplatepackagesignersV1Request)

	if err != nil {
		return err
	}

	*o = EzsigntemplatepackageEditEzsigntemplatepackagesignersV1Request(varEzsigntemplatepackageEditEzsigntemplatepackagesignersV1Request)

	return err
}

type NullableEzsigntemplatepackageEditEzsigntemplatepackagesignersV1Request struct {
	value *EzsigntemplatepackageEditEzsigntemplatepackagesignersV1Request
	isSet bool
}

func (v NullableEzsigntemplatepackageEditEzsigntemplatepackagesignersV1Request) Get() *EzsigntemplatepackageEditEzsigntemplatepackagesignersV1Request {
	return v.value
}

func (v *NullableEzsigntemplatepackageEditEzsigntemplatepackagesignersV1Request) Set(val *EzsigntemplatepackageEditEzsigntemplatepackagesignersV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigntemplatepackageEditEzsigntemplatepackagesignersV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigntemplatepackageEditEzsigntemplatepackagesignersV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigntemplatepackageEditEzsigntemplatepackagesignersV1Request(val *EzsigntemplatepackageEditEzsigntemplatepackagesignersV1Request) *NullableEzsigntemplatepackageEditEzsigntemplatepackagesignersV1Request {
	return &NullableEzsigntemplatepackageEditEzsigntemplatepackagesignersV1Request{value: val, isSet: true}
}

func (v NullableEzsigntemplatepackageEditEzsigntemplatepackagesignersV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigntemplatepackageEditEzsigntemplatepackagesignersV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


