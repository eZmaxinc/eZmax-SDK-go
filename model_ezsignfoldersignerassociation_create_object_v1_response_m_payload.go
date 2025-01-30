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

// checks if the EzsignfoldersignerassociationCreateObjectV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignfoldersignerassociationCreateObjectV1ResponseMPayload{}

// EzsignfoldersignerassociationCreateObjectV1ResponseMPayload Payload for POST /1/object/ezsignfoldersignerassociation
type EzsignfoldersignerassociationCreateObjectV1ResponseMPayload struct {
	// An array of unique IDs representing the object that were requested to be created.  They are returned in the same order as the array containing the objects to be created that was sent in the request.
	APkiEzsignfoldersignerassociationID []int32 `json:"a_pkiEzsignfoldersignerassociationID"`
}

type _EzsignfoldersignerassociationCreateObjectV1ResponseMPayload EzsignfoldersignerassociationCreateObjectV1ResponseMPayload

// NewEzsignfoldersignerassociationCreateObjectV1ResponseMPayload instantiates a new EzsignfoldersignerassociationCreateObjectV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignfoldersignerassociationCreateObjectV1ResponseMPayload(aPkiEzsignfoldersignerassociationID []int32) *EzsignfoldersignerassociationCreateObjectV1ResponseMPayload {
	this := EzsignfoldersignerassociationCreateObjectV1ResponseMPayload{}
	this.APkiEzsignfoldersignerassociationID = aPkiEzsignfoldersignerassociationID
	return &this
}

// NewEzsignfoldersignerassociationCreateObjectV1ResponseMPayloadWithDefaults instantiates a new EzsignfoldersignerassociationCreateObjectV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignfoldersignerassociationCreateObjectV1ResponseMPayloadWithDefaults() *EzsignfoldersignerassociationCreateObjectV1ResponseMPayload {
	this := EzsignfoldersignerassociationCreateObjectV1ResponseMPayload{}
	return &this
}

// GetAPkiEzsignfoldersignerassociationID returns the APkiEzsignfoldersignerassociationID field value
func (o *EzsignfoldersignerassociationCreateObjectV1ResponseMPayload) GetAPkiEzsignfoldersignerassociationID() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.APkiEzsignfoldersignerassociationID
}

// GetAPkiEzsignfoldersignerassociationIDOk returns a tuple with the APkiEzsignfoldersignerassociationID field value
// and a boolean to check if the value has been set.
func (o *EzsignfoldersignerassociationCreateObjectV1ResponseMPayload) GetAPkiEzsignfoldersignerassociationIDOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.APkiEzsignfoldersignerassociationID, true
}

// SetAPkiEzsignfoldersignerassociationID sets field value
func (o *EzsignfoldersignerassociationCreateObjectV1ResponseMPayload) SetAPkiEzsignfoldersignerassociationID(v []int32) {
	o.APkiEzsignfoldersignerassociationID = v
}

func (o EzsignfoldersignerassociationCreateObjectV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignfoldersignerassociationCreateObjectV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_pkiEzsignfoldersignerassociationID"] = o.APkiEzsignfoldersignerassociationID
	return toSerialize, nil
}

func (o *EzsignfoldersignerassociationCreateObjectV1ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"a_pkiEzsignfoldersignerassociationID",
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

	varEzsignfoldersignerassociationCreateObjectV1ResponseMPayload := _EzsignfoldersignerassociationCreateObjectV1ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsignfoldersignerassociationCreateObjectV1ResponseMPayload)

	if err != nil {
		return err
	}

	*o = EzsignfoldersignerassociationCreateObjectV1ResponseMPayload(varEzsignfoldersignerassociationCreateObjectV1ResponseMPayload)

	return err
}

type NullableEzsignfoldersignerassociationCreateObjectV1ResponseMPayload struct {
	value *EzsignfoldersignerassociationCreateObjectV1ResponseMPayload
	isSet bool
}

func (v NullableEzsignfoldersignerassociationCreateObjectV1ResponseMPayload) Get() *EzsignfoldersignerassociationCreateObjectV1ResponseMPayload {
	return v.value
}

func (v *NullableEzsignfoldersignerassociationCreateObjectV1ResponseMPayload) Set(val *EzsignfoldersignerassociationCreateObjectV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignfoldersignerassociationCreateObjectV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignfoldersignerassociationCreateObjectV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignfoldersignerassociationCreateObjectV1ResponseMPayload(val *EzsignfoldersignerassociationCreateObjectV1ResponseMPayload) *NullableEzsignfoldersignerassociationCreateObjectV1ResponseMPayload {
	return &NullableEzsignfoldersignerassociationCreateObjectV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableEzsignfoldersignerassociationCreateObjectV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignfoldersignerassociationCreateObjectV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


