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

// checks if the EzsigndocumentCreateElementV3Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigndocumentCreateElementV3Response{}

// EzsigndocumentCreateElementV3Response A Ezsigndocument createObject Response
type EzsigndocumentCreateElementV3Response struct {
	// The unique ID of the Ezsigndocument
	PkiEzsigndocumentID int32 `json:"pkiEzsigndocumentID"`
	// An array of possibly matching template.
	AObjMatchingtemplate []EzsigndocumentMatchingtemplateV3Response `json:"a_objMatchingtemplate"`
}

type _EzsigndocumentCreateElementV3Response EzsigndocumentCreateElementV3Response

// NewEzsigndocumentCreateElementV3Response instantiates a new EzsigndocumentCreateElementV3Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigndocumentCreateElementV3Response(pkiEzsigndocumentID int32, aObjMatchingtemplate []EzsigndocumentMatchingtemplateV3Response) *EzsigndocumentCreateElementV3Response {
	this := EzsigndocumentCreateElementV3Response{}
	this.PkiEzsigndocumentID = pkiEzsigndocumentID
	this.AObjMatchingtemplate = aObjMatchingtemplate
	return &this
}

// NewEzsigndocumentCreateElementV3ResponseWithDefaults instantiates a new EzsigndocumentCreateElementV3Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigndocumentCreateElementV3ResponseWithDefaults() *EzsigndocumentCreateElementV3Response {
	this := EzsigndocumentCreateElementV3Response{}
	return &this
}

// GetPkiEzsigndocumentID returns the PkiEzsigndocumentID field value
func (o *EzsigndocumentCreateElementV3Response) GetPkiEzsigndocumentID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiEzsigndocumentID
}

// GetPkiEzsigndocumentIDOk returns a tuple with the PkiEzsigndocumentID field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentCreateElementV3Response) GetPkiEzsigndocumentIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiEzsigndocumentID, true
}

// SetPkiEzsigndocumentID sets field value
func (o *EzsigndocumentCreateElementV3Response) SetPkiEzsigndocumentID(v int32) {
	o.PkiEzsigndocumentID = v
}

// GetAObjMatchingtemplate returns the AObjMatchingtemplate field value
func (o *EzsigndocumentCreateElementV3Response) GetAObjMatchingtemplate() []EzsigndocumentMatchingtemplateV3Response {
	if o == nil {
		var ret []EzsigndocumentMatchingtemplateV3Response
		return ret
	}

	return o.AObjMatchingtemplate
}

// GetAObjMatchingtemplateOk returns a tuple with the AObjMatchingtemplate field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentCreateElementV3Response) GetAObjMatchingtemplateOk() ([]EzsigndocumentMatchingtemplateV3Response, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjMatchingtemplate, true
}

// SetAObjMatchingtemplate sets field value
func (o *EzsigndocumentCreateElementV3Response) SetAObjMatchingtemplate(v []EzsigndocumentMatchingtemplateV3Response) {
	o.AObjMatchingtemplate = v
}

func (o EzsigndocumentCreateElementV3Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigndocumentCreateElementV3Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiEzsigndocumentID"] = o.PkiEzsigndocumentID
	toSerialize["a_objMatchingtemplate"] = o.AObjMatchingtemplate
	return toSerialize, nil
}

func (o *EzsigndocumentCreateElementV3Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiEzsigndocumentID",
		"a_objMatchingtemplate",
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

	varEzsigndocumentCreateElementV3Response := _EzsigndocumentCreateElementV3Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsigndocumentCreateElementV3Response)

	if err != nil {
		return err
	}

	*o = EzsigndocumentCreateElementV3Response(varEzsigndocumentCreateElementV3Response)

	return err
}

type NullableEzsigndocumentCreateElementV3Response struct {
	value *EzsigndocumentCreateElementV3Response
	isSet bool
}

func (v NullableEzsigndocumentCreateElementV3Response) Get() *EzsigndocumentCreateElementV3Response {
	return v.value
}

func (v *NullableEzsigndocumentCreateElementV3Response) Set(val *EzsigndocumentCreateElementV3Response) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigndocumentCreateElementV3Response) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigndocumentCreateElementV3Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigndocumentCreateElementV3Response(val *EzsigndocumentCreateElementV3Response) *NullableEzsigndocumentCreateElementV3Response {
	return &NullableEzsigndocumentCreateElementV3Response{value: val, isSet: true}
}

func (v NullableEzsigndocumentCreateElementV3Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigndocumentCreateElementV3Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


