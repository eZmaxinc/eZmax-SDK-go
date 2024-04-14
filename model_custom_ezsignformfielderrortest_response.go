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

// checks if the CustomEzsignformfielderrortestResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomEzsignformfielderrortestResponse{}

// CustomEzsignformfielderrortestResponse A Custom Ezsignformfielderrortest Object to contain the detail of the test error
type CustomEzsignformfielderrortestResponse struct {
	// The name of the test
	SEzsignformfielderrortestName string `json:"sEzsignformfielderrortestName"`
	// The detail why the test failed
	SEzsignformfielderrortestDetail string `json:"sEzsignformfielderrortestDetail"`
}

type _CustomEzsignformfielderrortestResponse CustomEzsignformfielderrortestResponse

// NewCustomEzsignformfielderrortestResponse instantiates a new CustomEzsignformfielderrortestResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomEzsignformfielderrortestResponse(sEzsignformfielderrortestName string, sEzsignformfielderrortestDetail string) *CustomEzsignformfielderrortestResponse {
	this := CustomEzsignformfielderrortestResponse{}
	this.SEzsignformfielderrortestName = sEzsignformfielderrortestName
	this.SEzsignformfielderrortestDetail = sEzsignformfielderrortestDetail
	return &this
}

// NewCustomEzsignformfielderrortestResponseWithDefaults instantiates a new CustomEzsignformfielderrortestResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomEzsignformfielderrortestResponseWithDefaults() *CustomEzsignformfielderrortestResponse {
	this := CustomEzsignformfielderrortestResponse{}
	return &this
}

// GetSEzsignformfielderrortestName returns the SEzsignformfielderrortestName field value
func (o *CustomEzsignformfielderrortestResponse) GetSEzsignformfielderrortestName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzsignformfielderrortestName
}

// GetSEzsignformfielderrortestNameOk returns a tuple with the SEzsignformfielderrortestName field value
// and a boolean to check if the value has been set.
func (o *CustomEzsignformfielderrortestResponse) GetSEzsignformfielderrortestNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SEzsignformfielderrortestName, true
}

// SetSEzsignformfielderrortestName sets field value
func (o *CustomEzsignformfielderrortestResponse) SetSEzsignformfielderrortestName(v string) {
	o.SEzsignformfielderrortestName = v
}

// GetSEzsignformfielderrortestDetail returns the SEzsignformfielderrortestDetail field value
func (o *CustomEzsignformfielderrortestResponse) GetSEzsignformfielderrortestDetail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzsignformfielderrortestDetail
}

// GetSEzsignformfielderrortestDetailOk returns a tuple with the SEzsignformfielderrortestDetail field value
// and a boolean to check if the value has been set.
func (o *CustomEzsignformfielderrortestResponse) GetSEzsignformfielderrortestDetailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SEzsignformfielderrortestDetail, true
}

// SetSEzsignformfielderrortestDetail sets field value
func (o *CustomEzsignformfielderrortestResponse) SetSEzsignformfielderrortestDetail(v string) {
	o.SEzsignformfielderrortestDetail = v
}

func (o CustomEzsignformfielderrortestResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomEzsignformfielderrortestResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sEzsignformfielderrortestName"] = o.SEzsignformfielderrortestName
	toSerialize["sEzsignformfielderrortestDetail"] = o.SEzsignformfielderrortestDetail
	return toSerialize, nil
}

func (o *CustomEzsignformfielderrortestResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"sEzsignformfielderrortestName",
		"sEzsignformfielderrortestDetail",
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

	varCustomEzsignformfielderrortestResponse := _CustomEzsignformfielderrortestResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCustomEzsignformfielderrortestResponse)

	if err != nil {
		return err
	}

	*o = CustomEzsignformfielderrortestResponse(varCustomEzsignformfielderrortestResponse)

	return err
}

type NullableCustomEzsignformfielderrortestResponse struct {
	value *CustomEzsignformfielderrortestResponse
	isSet bool
}

func (v NullableCustomEzsignformfielderrortestResponse) Get() *CustomEzsignformfielderrortestResponse {
	return v.value
}

func (v *NullableCustomEzsignformfielderrortestResponse) Set(val *CustomEzsignformfielderrortestResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomEzsignformfielderrortestResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomEzsignformfielderrortestResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomEzsignformfielderrortestResponse(val *CustomEzsignformfielderrortestResponse) *NullableCustomEzsignformfielderrortestResponse {
	return &NullableCustomEzsignformfielderrortestResponse{value: val, isSet: true}
}

func (v NullableCustomEzsignformfielderrortestResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomEzsignformfielderrortestResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


