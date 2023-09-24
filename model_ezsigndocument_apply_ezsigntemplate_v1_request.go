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

// checks if the EzsigndocumentApplyEzsigntemplateV1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigndocumentApplyEzsigntemplateV1Request{}

// EzsigndocumentApplyEzsigntemplateV1Request Request for POST /1/object/ezsigndocument/{pkiEzsigndocumentID}/applyezsigntemplate
type EzsigndocumentApplyEzsigntemplateV1Request struct {
	// The unique ID of the Ezsigntemplate
	FkiEzsigntemplateID int32 `json:"fkiEzsigntemplateID"`
	ASEzsigntemplatesigner []string `json:"a_sEzsigntemplatesigner"`
	APkiEzsignfoldersignerassociationID []int32 `json:"a_pkiEzsignfoldersignerassociationID"`
}

// NewEzsigndocumentApplyEzsigntemplateV1Request instantiates a new EzsigndocumentApplyEzsigntemplateV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigndocumentApplyEzsigntemplateV1Request(fkiEzsigntemplateID int32, aSEzsigntemplatesigner []string, aPkiEzsignfoldersignerassociationID []int32) *EzsigndocumentApplyEzsigntemplateV1Request {
	this := EzsigndocumentApplyEzsigntemplateV1Request{}
	this.FkiEzsigntemplateID = fkiEzsigntemplateID
	this.ASEzsigntemplatesigner = aSEzsigntemplatesigner
	this.APkiEzsignfoldersignerassociationID = aPkiEzsignfoldersignerassociationID
	return &this
}

// NewEzsigndocumentApplyEzsigntemplateV1RequestWithDefaults instantiates a new EzsigndocumentApplyEzsigntemplateV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigndocumentApplyEzsigntemplateV1RequestWithDefaults() *EzsigndocumentApplyEzsigntemplateV1Request {
	this := EzsigndocumentApplyEzsigntemplateV1Request{}
	return &this
}

// GetFkiEzsigntemplateID returns the FkiEzsigntemplateID field value
func (o *EzsigndocumentApplyEzsigntemplateV1Request) GetFkiEzsigntemplateID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzsigntemplateID
}

// GetFkiEzsigntemplateIDOk returns a tuple with the FkiEzsigntemplateID field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentApplyEzsigntemplateV1Request) GetFkiEzsigntemplateIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiEzsigntemplateID, true
}

// SetFkiEzsigntemplateID sets field value
func (o *EzsigndocumentApplyEzsigntemplateV1Request) SetFkiEzsigntemplateID(v int32) {
	o.FkiEzsigntemplateID = v
}

// GetASEzsigntemplatesigner returns the ASEzsigntemplatesigner field value
func (o *EzsigndocumentApplyEzsigntemplateV1Request) GetASEzsigntemplatesigner() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ASEzsigntemplatesigner
}

// GetASEzsigntemplatesignerOk returns a tuple with the ASEzsigntemplatesigner field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentApplyEzsigntemplateV1Request) GetASEzsigntemplatesignerOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ASEzsigntemplatesigner, true
}

// SetASEzsigntemplatesigner sets field value
func (o *EzsigndocumentApplyEzsigntemplateV1Request) SetASEzsigntemplatesigner(v []string) {
	o.ASEzsigntemplatesigner = v
}

// GetAPkiEzsignfoldersignerassociationID returns the APkiEzsignfoldersignerassociationID field value
func (o *EzsigndocumentApplyEzsigntemplateV1Request) GetAPkiEzsignfoldersignerassociationID() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.APkiEzsignfoldersignerassociationID
}

// GetAPkiEzsignfoldersignerassociationIDOk returns a tuple with the APkiEzsignfoldersignerassociationID field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentApplyEzsigntemplateV1Request) GetAPkiEzsignfoldersignerassociationIDOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.APkiEzsignfoldersignerassociationID, true
}

// SetAPkiEzsignfoldersignerassociationID sets field value
func (o *EzsigndocumentApplyEzsigntemplateV1Request) SetAPkiEzsignfoldersignerassociationID(v []int32) {
	o.APkiEzsignfoldersignerassociationID = v
}

func (o EzsigndocumentApplyEzsigntemplateV1Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigndocumentApplyEzsigntemplateV1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fkiEzsigntemplateID"] = o.FkiEzsigntemplateID
	toSerialize["a_sEzsigntemplatesigner"] = o.ASEzsigntemplatesigner
	toSerialize["a_pkiEzsignfoldersignerassociationID"] = o.APkiEzsignfoldersignerassociationID
	return toSerialize, nil
}

type NullableEzsigndocumentApplyEzsigntemplateV1Request struct {
	value *EzsigndocumentApplyEzsigntemplateV1Request
	isSet bool
}

func (v NullableEzsigndocumentApplyEzsigntemplateV1Request) Get() *EzsigndocumentApplyEzsigntemplateV1Request {
	return v.value
}

func (v *NullableEzsigndocumentApplyEzsigntemplateV1Request) Set(val *EzsigndocumentApplyEzsigntemplateV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigndocumentApplyEzsigntemplateV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigndocumentApplyEzsigntemplateV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigndocumentApplyEzsigntemplateV1Request(val *EzsigndocumentApplyEzsigntemplateV1Request) *NullableEzsigndocumentApplyEzsigntemplateV1Request {
	return &NullableEzsigndocumentApplyEzsigntemplateV1Request{value: val, isSet: true}
}

func (v NullableEzsigndocumentApplyEzsigntemplateV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigndocumentApplyEzsigntemplateV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


