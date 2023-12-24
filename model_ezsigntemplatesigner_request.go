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

// checks if the EzsigntemplatesignerRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigntemplatesignerRequest{}

// EzsigntemplatesignerRequest A Ezsigntemplatesigner Object
type EzsigntemplatesignerRequest struct {
	// The unique ID of the Ezsigntemplatesigner
	PkiEzsigntemplatesignerID *int32 `json:"pkiEzsigntemplatesignerID,omitempty"`
	// The unique ID of the Ezsigntemplate
	FkiEzsigntemplateID int32 `json:"fkiEzsigntemplateID"`
	// The description of the Ezsigntemplatesigner
	SEzsigntemplatesignerDescription string `json:"sEzsigntemplatesignerDescription"`
}

type _EzsigntemplatesignerRequest EzsigntemplatesignerRequest

// NewEzsigntemplatesignerRequest instantiates a new EzsigntemplatesignerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigntemplatesignerRequest(fkiEzsigntemplateID int32, sEzsigntemplatesignerDescription string) *EzsigntemplatesignerRequest {
	this := EzsigntemplatesignerRequest{}
	this.FkiEzsigntemplateID = fkiEzsigntemplateID
	this.SEzsigntemplatesignerDescription = sEzsigntemplatesignerDescription
	return &this
}

// NewEzsigntemplatesignerRequestWithDefaults instantiates a new EzsigntemplatesignerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigntemplatesignerRequestWithDefaults() *EzsigntemplatesignerRequest {
	this := EzsigntemplatesignerRequest{}
	return &this
}

// GetPkiEzsigntemplatesignerID returns the PkiEzsigntemplatesignerID field value if set, zero value otherwise.
func (o *EzsigntemplatesignerRequest) GetPkiEzsigntemplatesignerID() int32 {
	if o == nil || IsNil(o.PkiEzsigntemplatesignerID) {
		var ret int32
		return ret
	}
	return *o.PkiEzsigntemplatesignerID
}

// GetPkiEzsigntemplatesignerIDOk returns a tuple with the PkiEzsigntemplatesignerID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplatesignerRequest) GetPkiEzsigntemplatesignerIDOk() (*int32, bool) {
	if o == nil || IsNil(o.PkiEzsigntemplatesignerID) {
		return nil, false
	}
	return o.PkiEzsigntemplatesignerID, true
}

// HasPkiEzsigntemplatesignerID returns a boolean if a field has been set.
func (o *EzsigntemplatesignerRequest) HasPkiEzsigntemplatesignerID() bool {
	if o != nil && !IsNil(o.PkiEzsigntemplatesignerID) {
		return true
	}

	return false
}

// SetPkiEzsigntemplatesignerID gets a reference to the given int32 and assigns it to the PkiEzsigntemplatesignerID field.
func (o *EzsigntemplatesignerRequest) SetPkiEzsigntemplatesignerID(v int32) {
	o.PkiEzsigntemplatesignerID = &v
}

// GetFkiEzsigntemplateID returns the FkiEzsigntemplateID field value
func (o *EzsigntemplatesignerRequest) GetFkiEzsigntemplateID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzsigntemplateID
}

// GetFkiEzsigntemplateIDOk returns a tuple with the FkiEzsigntemplateID field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatesignerRequest) GetFkiEzsigntemplateIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiEzsigntemplateID, true
}

// SetFkiEzsigntemplateID sets field value
func (o *EzsigntemplatesignerRequest) SetFkiEzsigntemplateID(v int32) {
	o.FkiEzsigntemplateID = v
}

// GetSEzsigntemplatesignerDescription returns the SEzsigntemplatesignerDescription field value
func (o *EzsigntemplatesignerRequest) GetSEzsigntemplatesignerDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzsigntemplatesignerDescription
}

// GetSEzsigntemplatesignerDescriptionOk returns a tuple with the SEzsigntemplatesignerDescription field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatesignerRequest) GetSEzsigntemplatesignerDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SEzsigntemplatesignerDescription, true
}

// SetSEzsigntemplatesignerDescription sets field value
func (o *EzsigntemplatesignerRequest) SetSEzsigntemplatesignerDescription(v string) {
	o.SEzsigntemplatesignerDescription = v
}

func (o EzsigntemplatesignerRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigntemplatesignerRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PkiEzsigntemplatesignerID) {
		toSerialize["pkiEzsigntemplatesignerID"] = o.PkiEzsigntemplatesignerID
	}
	toSerialize["fkiEzsigntemplateID"] = o.FkiEzsigntemplateID
	toSerialize["sEzsigntemplatesignerDescription"] = o.SEzsigntemplatesignerDescription
	return toSerialize, nil
}

func (o *EzsigntemplatesignerRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fkiEzsigntemplateID",
		"sEzsigntemplatesignerDescription",
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

	varEzsigntemplatesignerRequest := _EzsigntemplatesignerRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsigntemplatesignerRequest)

	if err != nil {
		return err
	}

	*o = EzsigntemplatesignerRequest(varEzsigntemplatesignerRequest)

	return err
}

type NullableEzsigntemplatesignerRequest struct {
	value *EzsigntemplatesignerRequest
	isSet bool
}

func (v NullableEzsigntemplatesignerRequest) Get() *EzsigntemplatesignerRequest {
	return v.value
}

func (v *NullableEzsigntemplatesignerRequest) Set(val *EzsigntemplatesignerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigntemplatesignerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigntemplatesignerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigntemplatesignerRequest(val *EzsigntemplatesignerRequest) *NullableEzsigntemplatesignerRequest {
	return &NullableEzsigntemplatesignerRequest{value: val, isSet: true}
}

func (v NullableEzsigntemplatesignerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigntemplatesignerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


