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

// checks if the EzsignfoldersignerassociationCreateObjectV2Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignfoldersignerassociationCreateObjectV2Response{}

// EzsignfoldersignerassociationCreateObjectV2Response Response for POST /2/object/ezsignfoldersignerassociation
type EzsignfoldersignerassociationCreateObjectV2Response struct {
	CommonResponse
	MPayload EzsignfoldersignerassociationCreateObjectV2ResponseMPayload `json:"mPayload"`
}

type _EzsignfoldersignerassociationCreateObjectV2Response EzsignfoldersignerassociationCreateObjectV2Response

// NewEzsignfoldersignerassociationCreateObjectV2Response instantiates a new EzsignfoldersignerassociationCreateObjectV2Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignfoldersignerassociationCreateObjectV2Response(mPayload EzsignfoldersignerassociationCreateObjectV2ResponseMPayload, objDebugPayload CommonResponseObjDebugPayload) *EzsignfoldersignerassociationCreateObjectV2Response {
	this := EzsignfoldersignerassociationCreateObjectV2Response{}
	this.ObjDebugPayload = objDebugPayload
	this.MPayload = mPayload
	return &this
}

// NewEzsignfoldersignerassociationCreateObjectV2ResponseWithDefaults instantiates a new EzsignfoldersignerassociationCreateObjectV2Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignfoldersignerassociationCreateObjectV2ResponseWithDefaults() *EzsignfoldersignerassociationCreateObjectV2Response {
	this := EzsignfoldersignerassociationCreateObjectV2Response{}
	return &this
}

// GetMPayload returns the MPayload field value
func (o *EzsignfoldersignerassociationCreateObjectV2Response) GetMPayload() EzsignfoldersignerassociationCreateObjectV2ResponseMPayload {
	if o == nil {
		var ret EzsignfoldersignerassociationCreateObjectV2ResponseMPayload
		return ret
	}

	return o.MPayload
}

// GetMPayloadOk returns a tuple with the MPayload field value
// and a boolean to check if the value has been set.
func (o *EzsignfoldersignerassociationCreateObjectV2Response) GetMPayloadOk() (*EzsignfoldersignerassociationCreateObjectV2ResponseMPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MPayload, true
}

// SetMPayload sets field value
func (o *EzsignfoldersignerassociationCreateObjectV2Response) SetMPayload(v EzsignfoldersignerassociationCreateObjectV2ResponseMPayload) {
	o.MPayload = v
}

func (o EzsignfoldersignerassociationCreateObjectV2Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignfoldersignerassociationCreateObjectV2Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mPayload"] = o.MPayload
	return toSerialize, nil
}

func (o *EzsignfoldersignerassociationCreateObjectV2Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"mPayload",
		"objDebugPayload",
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

	varEzsignfoldersignerassociationCreateObjectV2Response := _EzsignfoldersignerassociationCreateObjectV2Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsignfoldersignerassociationCreateObjectV2Response)

	if err != nil {
		return err
	}

	*o = EzsignfoldersignerassociationCreateObjectV2Response(varEzsignfoldersignerassociationCreateObjectV2Response)

	return err
}

type NullableEzsignfoldersignerassociationCreateObjectV2Response struct {
	value *EzsignfoldersignerassociationCreateObjectV2Response
	isSet bool
}

func (v NullableEzsignfoldersignerassociationCreateObjectV2Response) Get() *EzsignfoldersignerassociationCreateObjectV2Response {
	return v.value
}

func (v *NullableEzsignfoldersignerassociationCreateObjectV2Response) Set(val *EzsignfoldersignerassociationCreateObjectV2Response) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignfoldersignerassociationCreateObjectV2Response) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignfoldersignerassociationCreateObjectV2Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignfoldersignerassociationCreateObjectV2Response(val *EzsignfoldersignerassociationCreateObjectV2Response) *NullableEzsignfoldersignerassociationCreateObjectV2Response {
	return &NullableEzsignfoldersignerassociationCreateObjectV2Response{value: val, isSet: true}
}

func (v NullableEzsignfoldersignerassociationCreateObjectV2Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignfoldersignerassociationCreateObjectV2Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


