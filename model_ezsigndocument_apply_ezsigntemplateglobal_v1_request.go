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

// checks if the EzsigndocumentApplyEzsigntemplateglobalV1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigndocumentApplyEzsigntemplateglobalV1Request{}

// EzsigndocumentApplyEzsigntemplateglobalV1Request Request for POST /1/object/ezsigndocument/{pkiEzsigndocumentID}/applyEzsigntemplateglobal
type EzsigndocumentApplyEzsigntemplateglobalV1Request struct {
	// The unique ID of the Ezsigntemplateglobal
	FkiEzsigntemplateglobalID int32 `json:"fkiEzsigntemplateglobalID"`
	ASEzsigntemplateglobalsigner []string `json:"a_sEzsigntemplateglobalsigner"`
	APkiEzsignfoldersignerassociationID []int32 `json:"a_pkiEzsignfoldersignerassociationID"`
}

type _EzsigndocumentApplyEzsigntemplateglobalV1Request EzsigndocumentApplyEzsigntemplateglobalV1Request

// NewEzsigndocumentApplyEzsigntemplateglobalV1Request instantiates a new EzsigndocumentApplyEzsigntemplateglobalV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigndocumentApplyEzsigntemplateglobalV1Request(fkiEzsigntemplateglobalID int32, aSEzsigntemplateglobalsigner []string, aPkiEzsignfoldersignerassociationID []int32) *EzsigndocumentApplyEzsigntemplateglobalV1Request {
	this := EzsigndocumentApplyEzsigntemplateglobalV1Request{}
	this.FkiEzsigntemplateglobalID = fkiEzsigntemplateglobalID
	this.ASEzsigntemplateglobalsigner = aSEzsigntemplateglobalsigner
	this.APkiEzsignfoldersignerassociationID = aPkiEzsignfoldersignerassociationID
	return &this
}

// NewEzsigndocumentApplyEzsigntemplateglobalV1RequestWithDefaults instantiates a new EzsigndocumentApplyEzsigntemplateglobalV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigndocumentApplyEzsigntemplateglobalV1RequestWithDefaults() *EzsigndocumentApplyEzsigntemplateglobalV1Request {
	this := EzsigndocumentApplyEzsigntemplateglobalV1Request{}
	return &this
}

// GetFkiEzsigntemplateglobalID returns the FkiEzsigntemplateglobalID field value
func (o *EzsigndocumentApplyEzsigntemplateglobalV1Request) GetFkiEzsigntemplateglobalID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzsigntemplateglobalID
}

// GetFkiEzsigntemplateglobalIDOk returns a tuple with the FkiEzsigntemplateglobalID field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentApplyEzsigntemplateglobalV1Request) GetFkiEzsigntemplateglobalIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiEzsigntemplateglobalID, true
}

// SetFkiEzsigntemplateglobalID sets field value
func (o *EzsigndocumentApplyEzsigntemplateglobalV1Request) SetFkiEzsigntemplateglobalID(v int32) {
	o.FkiEzsigntemplateglobalID = v
}

// GetASEzsigntemplateglobalsigner returns the ASEzsigntemplateglobalsigner field value
func (o *EzsigndocumentApplyEzsigntemplateglobalV1Request) GetASEzsigntemplateglobalsigner() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ASEzsigntemplateglobalsigner
}

// GetASEzsigntemplateglobalsignerOk returns a tuple with the ASEzsigntemplateglobalsigner field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentApplyEzsigntemplateglobalV1Request) GetASEzsigntemplateglobalsignerOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ASEzsigntemplateglobalsigner, true
}

// SetASEzsigntemplateglobalsigner sets field value
func (o *EzsigndocumentApplyEzsigntemplateglobalV1Request) SetASEzsigntemplateglobalsigner(v []string) {
	o.ASEzsigntemplateglobalsigner = v
}

// GetAPkiEzsignfoldersignerassociationID returns the APkiEzsignfoldersignerassociationID field value
func (o *EzsigndocumentApplyEzsigntemplateglobalV1Request) GetAPkiEzsignfoldersignerassociationID() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.APkiEzsignfoldersignerassociationID
}

// GetAPkiEzsignfoldersignerassociationIDOk returns a tuple with the APkiEzsignfoldersignerassociationID field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentApplyEzsigntemplateglobalV1Request) GetAPkiEzsignfoldersignerassociationIDOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.APkiEzsignfoldersignerassociationID, true
}

// SetAPkiEzsignfoldersignerassociationID sets field value
func (o *EzsigndocumentApplyEzsigntemplateglobalV1Request) SetAPkiEzsignfoldersignerassociationID(v []int32) {
	o.APkiEzsignfoldersignerassociationID = v
}

func (o EzsigndocumentApplyEzsigntemplateglobalV1Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigndocumentApplyEzsigntemplateglobalV1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fkiEzsigntemplateglobalID"] = o.FkiEzsigntemplateglobalID
	toSerialize["a_sEzsigntemplateglobalsigner"] = o.ASEzsigntemplateglobalsigner
	toSerialize["a_pkiEzsignfoldersignerassociationID"] = o.APkiEzsignfoldersignerassociationID
	return toSerialize, nil
}

func (o *EzsigndocumentApplyEzsigntemplateglobalV1Request) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fkiEzsigntemplateglobalID",
		"a_sEzsigntemplateglobalsigner",
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

	varEzsigndocumentApplyEzsigntemplateglobalV1Request := _EzsigndocumentApplyEzsigntemplateglobalV1Request{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsigndocumentApplyEzsigntemplateglobalV1Request)

	if err != nil {
		return err
	}

	*o = EzsigndocumentApplyEzsigntemplateglobalV1Request(varEzsigndocumentApplyEzsigntemplateglobalV1Request)

	return err
}

type NullableEzsigndocumentApplyEzsigntemplateglobalV1Request struct {
	value *EzsigndocumentApplyEzsigntemplateglobalV1Request
	isSet bool
}

func (v NullableEzsigndocumentApplyEzsigntemplateglobalV1Request) Get() *EzsigndocumentApplyEzsigntemplateglobalV1Request {
	return v.value
}

func (v *NullableEzsigndocumentApplyEzsigntemplateglobalV1Request) Set(val *EzsigndocumentApplyEzsigntemplateglobalV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigndocumentApplyEzsigntemplateglobalV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigndocumentApplyEzsigntemplateglobalV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigndocumentApplyEzsigntemplateglobalV1Request(val *EzsigndocumentApplyEzsigntemplateglobalV1Request) *NullableEzsigndocumentApplyEzsigntemplateglobalV1Request {
	return &NullableEzsigndocumentApplyEzsigntemplateglobalV1Request{value: val, isSet: true}
}

func (v NullableEzsigndocumentApplyEzsigntemplateglobalV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigndocumentApplyEzsigntemplateglobalV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

