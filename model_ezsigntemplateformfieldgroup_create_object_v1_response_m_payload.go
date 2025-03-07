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

// checks if the EzsigntemplateformfieldgroupCreateObjectV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigntemplateformfieldgroupCreateObjectV1ResponseMPayload{}

// EzsigntemplateformfieldgroupCreateObjectV1ResponseMPayload Payload for POST /1/object/ezsigntemplateformfieldgroup
type EzsigntemplateformfieldgroupCreateObjectV1ResponseMPayload struct {
	// An array of unique IDs representing the object that were requested to be created.  They are returned in the same order as the array containing the objects to be created that was sent in the request.
	APkiEzsigntemplateformfieldgroupID []int32 `json:"a_pkiEzsigntemplateformfieldgroupID"`
}

type _EzsigntemplateformfieldgroupCreateObjectV1ResponseMPayload EzsigntemplateformfieldgroupCreateObjectV1ResponseMPayload

// NewEzsigntemplateformfieldgroupCreateObjectV1ResponseMPayload instantiates a new EzsigntemplateformfieldgroupCreateObjectV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigntemplateformfieldgroupCreateObjectV1ResponseMPayload(aPkiEzsigntemplateformfieldgroupID []int32) *EzsigntemplateformfieldgroupCreateObjectV1ResponseMPayload {
	this := EzsigntemplateformfieldgroupCreateObjectV1ResponseMPayload{}
	this.APkiEzsigntemplateformfieldgroupID = aPkiEzsigntemplateformfieldgroupID
	return &this
}

// NewEzsigntemplateformfieldgroupCreateObjectV1ResponseMPayloadWithDefaults instantiates a new EzsigntemplateformfieldgroupCreateObjectV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigntemplateformfieldgroupCreateObjectV1ResponseMPayloadWithDefaults() *EzsigntemplateformfieldgroupCreateObjectV1ResponseMPayload {
	this := EzsigntemplateformfieldgroupCreateObjectV1ResponseMPayload{}
	return &this
}

// GetAPkiEzsigntemplateformfieldgroupID returns the APkiEzsigntemplateformfieldgroupID field value
func (o *EzsigntemplateformfieldgroupCreateObjectV1ResponseMPayload) GetAPkiEzsigntemplateformfieldgroupID() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.APkiEzsigntemplateformfieldgroupID
}

// GetAPkiEzsigntemplateformfieldgroupIDOk returns a tuple with the APkiEzsigntemplateformfieldgroupID field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldgroupCreateObjectV1ResponseMPayload) GetAPkiEzsigntemplateformfieldgroupIDOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.APkiEzsigntemplateformfieldgroupID, true
}

// SetAPkiEzsigntemplateformfieldgroupID sets field value
func (o *EzsigntemplateformfieldgroupCreateObjectV1ResponseMPayload) SetAPkiEzsigntemplateformfieldgroupID(v []int32) {
	o.APkiEzsigntemplateformfieldgroupID = v
}

func (o EzsigntemplateformfieldgroupCreateObjectV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigntemplateformfieldgroupCreateObjectV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_pkiEzsigntemplateformfieldgroupID"] = o.APkiEzsigntemplateformfieldgroupID
	return toSerialize, nil
}

func (o *EzsigntemplateformfieldgroupCreateObjectV1ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"a_pkiEzsigntemplateformfieldgroupID",
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

	varEzsigntemplateformfieldgroupCreateObjectV1ResponseMPayload := _EzsigntemplateformfieldgroupCreateObjectV1ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsigntemplateformfieldgroupCreateObjectV1ResponseMPayload)

	if err != nil {
		return err
	}

	*o = EzsigntemplateformfieldgroupCreateObjectV1ResponseMPayload(varEzsigntemplateformfieldgroupCreateObjectV1ResponseMPayload)

	return err
}

type NullableEzsigntemplateformfieldgroupCreateObjectV1ResponseMPayload struct {
	value *EzsigntemplateformfieldgroupCreateObjectV1ResponseMPayload
	isSet bool
}

func (v NullableEzsigntemplateformfieldgroupCreateObjectV1ResponseMPayload) Get() *EzsigntemplateformfieldgroupCreateObjectV1ResponseMPayload {
	return v.value
}

func (v *NullableEzsigntemplateformfieldgroupCreateObjectV1ResponseMPayload) Set(val *EzsigntemplateformfieldgroupCreateObjectV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigntemplateformfieldgroupCreateObjectV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigntemplateformfieldgroupCreateObjectV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigntemplateformfieldgroupCreateObjectV1ResponseMPayload(val *EzsigntemplateformfieldgroupCreateObjectV1ResponseMPayload) *NullableEzsigntemplateformfieldgroupCreateObjectV1ResponseMPayload {
	return &NullableEzsigntemplateformfieldgroupCreateObjectV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableEzsigntemplateformfieldgroupCreateObjectV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigntemplateformfieldgroupCreateObjectV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


