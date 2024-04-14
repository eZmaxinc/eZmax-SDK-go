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

// checks if the EzsignfoldersignerassociationCreateObjectV2ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignfoldersignerassociationCreateObjectV2ResponseMPayload{}

// EzsignfoldersignerassociationCreateObjectV2ResponseMPayload Payload for POST /2/object/ezsignfoldersignerassociation
type EzsignfoldersignerassociationCreateObjectV2ResponseMPayload struct {
	// An array of unique IDs representing the object that were requested to be created.  They are returned in the same order as the array containing the objects to be created that was sent in the request.
	APkiEzsignfoldersignerassociationID []int32 `json:"a_pkiEzsignfoldersignerassociationID"`
}

type _EzsignfoldersignerassociationCreateObjectV2ResponseMPayload EzsignfoldersignerassociationCreateObjectV2ResponseMPayload

// NewEzsignfoldersignerassociationCreateObjectV2ResponseMPayload instantiates a new EzsignfoldersignerassociationCreateObjectV2ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignfoldersignerassociationCreateObjectV2ResponseMPayload(aPkiEzsignfoldersignerassociationID []int32) *EzsignfoldersignerassociationCreateObjectV2ResponseMPayload {
	this := EzsignfoldersignerassociationCreateObjectV2ResponseMPayload{}
	this.APkiEzsignfoldersignerassociationID = aPkiEzsignfoldersignerassociationID
	return &this
}

// NewEzsignfoldersignerassociationCreateObjectV2ResponseMPayloadWithDefaults instantiates a new EzsignfoldersignerassociationCreateObjectV2ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignfoldersignerassociationCreateObjectV2ResponseMPayloadWithDefaults() *EzsignfoldersignerassociationCreateObjectV2ResponseMPayload {
	this := EzsignfoldersignerassociationCreateObjectV2ResponseMPayload{}
	return &this
}

// GetAPkiEzsignfoldersignerassociationID returns the APkiEzsignfoldersignerassociationID field value
func (o *EzsignfoldersignerassociationCreateObjectV2ResponseMPayload) GetAPkiEzsignfoldersignerassociationID() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.APkiEzsignfoldersignerassociationID
}

// GetAPkiEzsignfoldersignerassociationIDOk returns a tuple with the APkiEzsignfoldersignerassociationID field value
// and a boolean to check if the value has been set.
func (o *EzsignfoldersignerassociationCreateObjectV2ResponseMPayload) GetAPkiEzsignfoldersignerassociationIDOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.APkiEzsignfoldersignerassociationID, true
}

// SetAPkiEzsignfoldersignerassociationID sets field value
func (o *EzsignfoldersignerassociationCreateObjectV2ResponseMPayload) SetAPkiEzsignfoldersignerassociationID(v []int32) {
	o.APkiEzsignfoldersignerassociationID = v
}

func (o EzsignfoldersignerassociationCreateObjectV2ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignfoldersignerassociationCreateObjectV2ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_pkiEzsignfoldersignerassociationID"] = o.APkiEzsignfoldersignerassociationID
	return toSerialize, nil
}

func (o *EzsignfoldersignerassociationCreateObjectV2ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
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

	varEzsignfoldersignerassociationCreateObjectV2ResponseMPayload := _EzsignfoldersignerassociationCreateObjectV2ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsignfoldersignerassociationCreateObjectV2ResponseMPayload)

	if err != nil {
		return err
	}

	*o = EzsignfoldersignerassociationCreateObjectV2ResponseMPayload(varEzsignfoldersignerassociationCreateObjectV2ResponseMPayload)

	return err
}

type NullableEzsignfoldersignerassociationCreateObjectV2ResponseMPayload struct {
	value *EzsignfoldersignerassociationCreateObjectV2ResponseMPayload
	isSet bool
}

func (v NullableEzsignfoldersignerassociationCreateObjectV2ResponseMPayload) Get() *EzsignfoldersignerassociationCreateObjectV2ResponseMPayload {
	return v.value
}

func (v *NullableEzsignfoldersignerassociationCreateObjectV2ResponseMPayload) Set(val *EzsignfoldersignerassociationCreateObjectV2ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignfoldersignerassociationCreateObjectV2ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignfoldersignerassociationCreateObjectV2ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignfoldersignerassociationCreateObjectV2ResponseMPayload(val *EzsignfoldersignerassociationCreateObjectV2ResponseMPayload) *NullableEzsignfoldersignerassociationCreateObjectV2ResponseMPayload {
	return &NullableEzsignfoldersignerassociationCreateObjectV2ResponseMPayload{value: val, isSet: true}
}

func (v NullableEzsignfoldersignerassociationCreateObjectV2ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignfoldersignerassociationCreateObjectV2ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


