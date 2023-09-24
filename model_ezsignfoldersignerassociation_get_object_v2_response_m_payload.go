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
)

// checks if the EzsignfoldersignerassociationGetObjectV2ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignfoldersignerassociationGetObjectV2ResponseMPayload{}

// EzsignfoldersignerassociationGetObjectV2ResponseMPayload Payload for GET /2/object/ezsignfoldersignerassociation/{pkiEzsignfoldersignerassociationID}
type EzsignfoldersignerassociationGetObjectV2ResponseMPayload struct {
	ObjEzsignfoldersignerassociation EzsignfoldersignerassociationResponseCompound `json:"objEzsignfoldersignerassociation"`
}

// NewEzsignfoldersignerassociationGetObjectV2ResponseMPayload instantiates a new EzsignfoldersignerassociationGetObjectV2ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignfoldersignerassociationGetObjectV2ResponseMPayload(objEzsignfoldersignerassociation EzsignfoldersignerassociationResponseCompound) *EzsignfoldersignerassociationGetObjectV2ResponseMPayload {
	this := EzsignfoldersignerassociationGetObjectV2ResponseMPayload{}
	this.ObjEzsignfoldersignerassociation = objEzsignfoldersignerassociation
	return &this
}

// NewEzsignfoldersignerassociationGetObjectV2ResponseMPayloadWithDefaults instantiates a new EzsignfoldersignerassociationGetObjectV2ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignfoldersignerassociationGetObjectV2ResponseMPayloadWithDefaults() *EzsignfoldersignerassociationGetObjectV2ResponseMPayload {
	this := EzsignfoldersignerassociationGetObjectV2ResponseMPayload{}
	return &this
}

// GetObjEzsignfoldersignerassociation returns the ObjEzsignfoldersignerassociation field value
func (o *EzsignfoldersignerassociationGetObjectV2ResponseMPayload) GetObjEzsignfoldersignerassociation() EzsignfoldersignerassociationResponseCompound {
	if o == nil {
		var ret EzsignfoldersignerassociationResponseCompound
		return ret
	}

	return o.ObjEzsignfoldersignerassociation
}

// GetObjEzsignfoldersignerassociationOk returns a tuple with the ObjEzsignfoldersignerassociation field value
// and a boolean to check if the value has been set.
func (o *EzsignfoldersignerassociationGetObjectV2ResponseMPayload) GetObjEzsignfoldersignerassociationOk() (*EzsignfoldersignerassociationResponseCompound, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjEzsignfoldersignerassociation, true
}

// SetObjEzsignfoldersignerassociation sets field value
func (o *EzsignfoldersignerassociationGetObjectV2ResponseMPayload) SetObjEzsignfoldersignerassociation(v EzsignfoldersignerassociationResponseCompound) {
	o.ObjEzsignfoldersignerassociation = v
}

func (o EzsignfoldersignerassociationGetObjectV2ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignfoldersignerassociationGetObjectV2ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objEzsignfoldersignerassociation"] = o.ObjEzsignfoldersignerassociation
	return toSerialize, nil
}

type NullableEzsignfoldersignerassociationGetObjectV2ResponseMPayload struct {
	value *EzsignfoldersignerassociationGetObjectV2ResponseMPayload
	isSet bool
}

func (v NullableEzsignfoldersignerassociationGetObjectV2ResponseMPayload) Get() *EzsignfoldersignerassociationGetObjectV2ResponseMPayload {
	return v.value
}

func (v *NullableEzsignfoldersignerassociationGetObjectV2ResponseMPayload) Set(val *EzsignfoldersignerassociationGetObjectV2ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignfoldersignerassociationGetObjectV2ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignfoldersignerassociationGetObjectV2ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignfoldersignerassociationGetObjectV2ResponseMPayload(val *EzsignfoldersignerassociationGetObjectV2ResponseMPayload) *NullableEzsignfoldersignerassociationGetObjectV2ResponseMPayload {
	return &NullableEzsignfoldersignerassociationGetObjectV2ResponseMPayload{value: val, isSet: true}
}

func (v NullableEzsignfoldersignerassociationGetObjectV2ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignfoldersignerassociationGetObjectV2ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

