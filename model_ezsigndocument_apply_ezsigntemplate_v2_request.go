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

// checks if the EzsigndocumentApplyEzsigntemplateV2Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigndocumentApplyEzsigntemplateV2Request{}

// EzsigndocumentApplyEzsigntemplateV2Request Request for POST /2/object/ezsigndocument/{pkiEzsigndocumentID}/applyezsigntemplate
type EzsigndocumentApplyEzsigntemplateV2Request struct {
	// The unique ID of the Ezsigntemplate
	FkiEzsigntemplateID int32 `json:"fkiEzsigntemplateID"`
	ASEzsigntemplatesigner []string `json:"a_sEzsigntemplatesigner"`
	APkiEzsignfoldersignerassociationID []int32 `json:"a_pkiEzsignfoldersignerassociationID"`
}

type _EzsigndocumentApplyEzsigntemplateV2Request EzsigndocumentApplyEzsigntemplateV2Request

// NewEzsigndocumentApplyEzsigntemplateV2Request instantiates a new EzsigndocumentApplyEzsigntemplateV2Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigndocumentApplyEzsigntemplateV2Request(fkiEzsigntemplateID int32, aSEzsigntemplatesigner []string, aPkiEzsignfoldersignerassociationID []int32) *EzsigndocumentApplyEzsigntemplateV2Request {
	this := EzsigndocumentApplyEzsigntemplateV2Request{}
	this.FkiEzsigntemplateID = fkiEzsigntemplateID
	this.ASEzsigntemplatesigner = aSEzsigntemplatesigner
	this.APkiEzsignfoldersignerassociationID = aPkiEzsignfoldersignerassociationID
	return &this
}

// NewEzsigndocumentApplyEzsigntemplateV2RequestWithDefaults instantiates a new EzsigndocumentApplyEzsigntemplateV2Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigndocumentApplyEzsigntemplateV2RequestWithDefaults() *EzsigndocumentApplyEzsigntemplateV2Request {
	this := EzsigndocumentApplyEzsigntemplateV2Request{}
	return &this
}

// GetFkiEzsigntemplateID returns the FkiEzsigntemplateID field value
func (o *EzsigndocumentApplyEzsigntemplateV2Request) GetFkiEzsigntemplateID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzsigntemplateID
}

// GetFkiEzsigntemplateIDOk returns a tuple with the FkiEzsigntemplateID field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentApplyEzsigntemplateV2Request) GetFkiEzsigntemplateIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiEzsigntemplateID, true
}

// SetFkiEzsigntemplateID sets field value
func (o *EzsigndocumentApplyEzsigntemplateV2Request) SetFkiEzsigntemplateID(v int32) {
	o.FkiEzsigntemplateID = v
}

// GetASEzsigntemplatesigner returns the ASEzsigntemplatesigner field value
func (o *EzsigndocumentApplyEzsigntemplateV2Request) GetASEzsigntemplatesigner() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ASEzsigntemplatesigner
}

// GetASEzsigntemplatesignerOk returns a tuple with the ASEzsigntemplatesigner field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentApplyEzsigntemplateV2Request) GetASEzsigntemplatesignerOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ASEzsigntemplatesigner, true
}

// SetASEzsigntemplatesigner sets field value
func (o *EzsigndocumentApplyEzsigntemplateV2Request) SetASEzsigntemplatesigner(v []string) {
	o.ASEzsigntemplatesigner = v
}

// GetAPkiEzsignfoldersignerassociationID returns the APkiEzsignfoldersignerassociationID field value
func (o *EzsigndocumentApplyEzsigntemplateV2Request) GetAPkiEzsignfoldersignerassociationID() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.APkiEzsignfoldersignerassociationID
}

// GetAPkiEzsignfoldersignerassociationIDOk returns a tuple with the APkiEzsignfoldersignerassociationID field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentApplyEzsigntemplateV2Request) GetAPkiEzsignfoldersignerassociationIDOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.APkiEzsignfoldersignerassociationID, true
}

// SetAPkiEzsignfoldersignerassociationID sets field value
func (o *EzsigndocumentApplyEzsigntemplateV2Request) SetAPkiEzsignfoldersignerassociationID(v []int32) {
	o.APkiEzsignfoldersignerassociationID = v
}

func (o EzsigndocumentApplyEzsigntemplateV2Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigndocumentApplyEzsigntemplateV2Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fkiEzsigntemplateID"] = o.FkiEzsigntemplateID
	toSerialize["a_sEzsigntemplatesigner"] = o.ASEzsigntemplatesigner
	toSerialize["a_pkiEzsignfoldersignerassociationID"] = o.APkiEzsignfoldersignerassociationID
	return toSerialize, nil
}

func (o *EzsigndocumentApplyEzsigntemplateV2Request) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fkiEzsigntemplateID",
		"a_sEzsigntemplatesigner",
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

	varEzsigndocumentApplyEzsigntemplateV2Request := _EzsigndocumentApplyEzsigntemplateV2Request{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsigndocumentApplyEzsigntemplateV2Request)

	if err != nil {
		return err
	}

	*o = EzsigndocumentApplyEzsigntemplateV2Request(varEzsigndocumentApplyEzsigntemplateV2Request)

	return err
}

type NullableEzsigndocumentApplyEzsigntemplateV2Request struct {
	value *EzsigndocumentApplyEzsigntemplateV2Request
	isSet bool
}

func (v NullableEzsigndocumentApplyEzsigntemplateV2Request) Get() *EzsigndocumentApplyEzsigntemplateV2Request {
	return v.value
}

func (v *NullableEzsigndocumentApplyEzsigntemplateV2Request) Set(val *EzsigndocumentApplyEzsigntemplateV2Request) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigndocumentApplyEzsigntemplateV2Request) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigndocumentApplyEzsigntemplateV2Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigndocumentApplyEzsigntemplateV2Request(val *EzsigndocumentApplyEzsigntemplateV2Request) *NullableEzsigndocumentApplyEzsigntemplateV2Request {
	return &NullableEzsigndocumentApplyEzsigntemplateV2Request{value: val, isSet: true}
}

func (v NullableEzsigndocumentApplyEzsigntemplateV2Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigndocumentApplyEzsigntemplateV2Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


