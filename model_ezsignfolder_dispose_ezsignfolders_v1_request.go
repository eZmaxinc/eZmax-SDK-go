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

// checks if the EzsignfolderDisposeEzsignfoldersV1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignfolderDisposeEzsignfoldersV1Request{}

// EzsignfolderDisposeEzsignfoldersV1Request Request for POST /1/object/ezsignfolder/disposeEzsignfolders
type EzsignfolderDisposeEzsignfoldersV1Request struct {
	APkiEzsignfolderID []int32 `json:"a_pkiEzsignfolderID"`
}

type _EzsignfolderDisposeEzsignfoldersV1Request EzsignfolderDisposeEzsignfoldersV1Request

// NewEzsignfolderDisposeEzsignfoldersV1Request instantiates a new EzsignfolderDisposeEzsignfoldersV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignfolderDisposeEzsignfoldersV1Request(aPkiEzsignfolderID []int32) *EzsignfolderDisposeEzsignfoldersV1Request {
	this := EzsignfolderDisposeEzsignfoldersV1Request{}
	this.APkiEzsignfolderID = aPkiEzsignfolderID
	return &this
}

// NewEzsignfolderDisposeEzsignfoldersV1RequestWithDefaults instantiates a new EzsignfolderDisposeEzsignfoldersV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignfolderDisposeEzsignfoldersV1RequestWithDefaults() *EzsignfolderDisposeEzsignfoldersV1Request {
	this := EzsignfolderDisposeEzsignfoldersV1Request{}
	return &this
}

// GetAPkiEzsignfolderID returns the APkiEzsignfolderID field value
func (o *EzsignfolderDisposeEzsignfoldersV1Request) GetAPkiEzsignfolderID() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.APkiEzsignfolderID
}

// GetAPkiEzsignfolderIDOk returns a tuple with the APkiEzsignfolderID field value
// and a boolean to check if the value has been set.
func (o *EzsignfolderDisposeEzsignfoldersV1Request) GetAPkiEzsignfolderIDOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.APkiEzsignfolderID, true
}

// SetAPkiEzsignfolderID sets field value
func (o *EzsignfolderDisposeEzsignfoldersV1Request) SetAPkiEzsignfolderID(v []int32) {
	o.APkiEzsignfolderID = v
}

func (o EzsignfolderDisposeEzsignfoldersV1Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignfolderDisposeEzsignfoldersV1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_pkiEzsignfolderID"] = o.APkiEzsignfolderID
	return toSerialize, nil
}

func (o *EzsignfolderDisposeEzsignfoldersV1Request) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"a_pkiEzsignfolderID",
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

	varEzsignfolderDisposeEzsignfoldersV1Request := _EzsignfolderDisposeEzsignfoldersV1Request{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsignfolderDisposeEzsignfoldersV1Request)

	if err != nil {
		return err
	}

	*o = EzsignfolderDisposeEzsignfoldersV1Request(varEzsignfolderDisposeEzsignfoldersV1Request)

	return err
}

type NullableEzsignfolderDisposeEzsignfoldersV1Request struct {
	value *EzsignfolderDisposeEzsignfoldersV1Request
	isSet bool
}

func (v NullableEzsignfolderDisposeEzsignfoldersV1Request) Get() *EzsignfolderDisposeEzsignfoldersV1Request {
	return v.value
}

func (v *NullableEzsignfolderDisposeEzsignfoldersV1Request) Set(val *EzsignfolderDisposeEzsignfoldersV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignfolderDisposeEzsignfoldersV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignfolderDisposeEzsignfoldersV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignfolderDisposeEzsignfoldersV1Request(val *EzsignfolderDisposeEzsignfoldersV1Request) *NullableEzsignfolderDisposeEzsignfoldersV1Request {
	return &NullableEzsignfolderDisposeEzsignfoldersV1Request{value: val, isSet: true}
}

func (v NullableEzsignfolderDisposeEzsignfoldersV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignfolderDisposeEzsignfoldersV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


