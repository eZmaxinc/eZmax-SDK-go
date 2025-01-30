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

// checks if the EzsigntemplateformfieldgroupsignerResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigntemplateformfieldgroupsignerResponse{}

// EzsigntemplateformfieldgroupsignerResponse A Ezsigntemplateformfieldgroupsigner Object
type EzsigntemplateformfieldgroupsignerResponse struct {
	// The unique ID of the Ezsigntemplateformfieldgroupsigner
	PkiEzsigntemplateformfieldgroupsignerID int32 `json:"pkiEzsigntemplateformfieldgroupsignerID"`
	// The unique ID of the Ezsigntemplatesigner
	FkiEzsigntemplatesignerID int32 `json:"fkiEzsigntemplatesignerID"`
}

type _EzsigntemplateformfieldgroupsignerResponse EzsigntemplateformfieldgroupsignerResponse

// NewEzsigntemplateformfieldgroupsignerResponse instantiates a new EzsigntemplateformfieldgroupsignerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigntemplateformfieldgroupsignerResponse(pkiEzsigntemplateformfieldgroupsignerID int32, fkiEzsigntemplatesignerID int32) *EzsigntemplateformfieldgroupsignerResponse {
	this := EzsigntemplateformfieldgroupsignerResponse{}
	this.PkiEzsigntemplateformfieldgroupsignerID = pkiEzsigntemplateformfieldgroupsignerID
	this.FkiEzsigntemplatesignerID = fkiEzsigntemplatesignerID
	return &this
}

// NewEzsigntemplateformfieldgroupsignerResponseWithDefaults instantiates a new EzsigntemplateformfieldgroupsignerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigntemplateformfieldgroupsignerResponseWithDefaults() *EzsigntemplateformfieldgroupsignerResponse {
	this := EzsigntemplateformfieldgroupsignerResponse{}
	return &this
}

// GetPkiEzsigntemplateformfieldgroupsignerID returns the PkiEzsigntemplateformfieldgroupsignerID field value
func (o *EzsigntemplateformfieldgroupsignerResponse) GetPkiEzsigntemplateformfieldgroupsignerID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiEzsigntemplateformfieldgroupsignerID
}

// GetPkiEzsigntemplateformfieldgroupsignerIDOk returns a tuple with the PkiEzsigntemplateformfieldgroupsignerID field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldgroupsignerResponse) GetPkiEzsigntemplateformfieldgroupsignerIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiEzsigntemplateformfieldgroupsignerID, true
}

// SetPkiEzsigntemplateformfieldgroupsignerID sets field value
func (o *EzsigntemplateformfieldgroupsignerResponse) SetPkiEzsigntemplateformfieldgroupsignerID(v int32) {
	o.PkiEzsigntemplateformfieldgroupsignerID = v
}

// GetFkiEzsigntemplatesignerID returns the FkiEzsigntemplatesignerID field value
func (o *EzsigntemplateformfieldgroupsignerResponse) GetFkiEzsigntemplatesignerID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzsigntemplatesignerID
}

// GetFkiEzsigntemplatesignerIDOk returns a tuple with the FkiEzsigntemplatesignerID field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldgroupsignerResponse) GetFkiEzsigntemplatesignerIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiEzsigntemplatesignerID, true
}

// SetFkiEzsigntemplatesignerID sets field value
func (o *EzsigntemplateformfieldgroupsignerResponse) SetFkiEzsigntemplatesignerID(v int32) {
	o.FkiEzsigntemplatesignerID = v
}

func (o EzsigntemplateformfieldgroupsignerResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigntemplateformfieldgroupsignerResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiEzsigntemplateformfieldgroupsignerID"] = o.PkiEzsigntemplateformfieldgroupsignerID
	toSerialize["fkiEzsigntemplatesignerID"] = o.FkiEzsigntemplatesignerID
	return toSerialize, nil
}

func (o *EzsigntemplateformfieldgroupsignerResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiEzsigntemplateformfieldgroupsignerID",
		"fkiEzsigntemplatesignerID",
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

	varEzsigntemplateformfieldgroupsignerResponse := _EzsigntemplateformfieldgroupsignerResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsigntemplateformfieldgroupsignerResponse)

	if err != nil {
		return err
	}

	*o = EzsigntemplateformfieldgroupsignerResponse(varEzsigntemplateformfieldgroupsignerResponse)

	return err
}

type NullableEzsigntemplateformfieldgroupsignerResponse struct {
	value *EzsigntemplateformfieldgroupsignerResponse
	isSet bool
}

func (v NullableEzsigntemplateformfieldgroupsignerResponse) Get() *EzsigntemplateformfieldgroupsignerResponse {
	return v.value
}

func (v *NullableEzsigntemplateformfieldgroupsignerResponse) Set(val *EzsigntemplateformfieldgroupsignerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigntemplateformfieldgroupsignerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigntemplateformfieldgroupsignerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigntemplateformfieldgroupsignerResponse(val *EzsigntemplateformfieldgroupsignerResponse) *NullableEzsigntemplateformfieldgroupsignerResponse {
	return &NullableEzsigntemplateformfieldgroupsignerResponse{value: val, isSet: true}
}

func (v NullableEzsigntemplateformfieldgroupsignerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigntemplateformfieldgroupsignerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


