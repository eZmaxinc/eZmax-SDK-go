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

// checks if the EzsigntemplatepackagesignerRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigntemplatepackagesignerRequest{}

// EzsigntemplatepackagesignerRequest A Ezsigntemplatepackagesigner Object
type EzsigntemplatepackagesignerRequest struct {
	// The unique ID of the Ezsigntemplatepackagesigner
	PkiEzsigntemplatepackagesignerID *int32 `json:"pkiEzsigntemplatepackagesignerID,omitempty"`
	// The unique ID of the Ezsigntemplatepackage
	FkiEzsigntemplatepackageID int32 `json:"fkiEzsigntemplatepackageID"`
	// The description of the Ezsigntemplatepackagesigner
	SEzsigntemplatepackagesignerDescription string `json:"sEzsigntemplatepackagesignerDescription"`
}

type _EzsigntemplatepackagesignerRequest EzsigntemplatepackagesignerRequest

// NewEzsigntemplatepackagesignerRequest instantiates a new EzsigntemplatepackagesignerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigntemplatepackagesignerRequest(fkiEzsigntemplatepackageID int32, sEzsigntemplatepackagesignerDescription string) *EzsigntemplatepackagesignerRequest {
	this := EzsigntemplatepackagesignerRequest{}
	this.FkiEzsigntemplatepackageID = fkiEzsigntemplatepackageID
	this.SEzsigntemplatepackagesignerDescription = sEzsigntemplatepackagesignerDescription
	return &this
}

// NewEzsigntemplatepackagesignerRequestWithDefaults instantiates a new EzsigntemplatepackagesignerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigntemplatepackagesignerRequestWithDefaults() *EzsigntemplatepackagesignerRequest {
	this := EzsigntemplatepackagesignerRequest{}
	return &this
}

// GetPkiEzsigntemplatepackagesignerID returns the PkiEzsigntemplatepackagesignerID field value if set, zero value otherwise.
func (o *EzsigntemplatepackagesignerRequest) GetPkiEzsigntemplatepackagesignerID() int32 {
	if o == nil || IsNil(o.PkiEzsigntemplatepackagesignerID) {
		var ret int32
		return ret
	}
	return *o.PkiEzsigntemplatepackagesignerID
}

// GetPkiEzsigntemplatepackagesignerIDOk returns a tuple with the PkiEzsigntemplatepackagesignerID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepackagesignerRequest) GetPkiEzsigntemplatepackagesignerIDOk() (*int32, bool) {
	if o == nil || IsNil(o.PkiEzsigntemplatepackagesignerID) {
		return nil, false
	}
	return o.PkiEzsigntemplatepackagesignerID, true
}

// HasPkiEzsigntemplatepackagesignerID returns a boolean if a field has been set.
func (o *EzsigntemplatepackagesignerRequest) HasPkiEzsigntemplatepackagesignerID() bool {
	if o != nil && !IsNil(o.PkiEzsigntemplatepackagesignerID) {
		return true
	}

	return false
}

// SetPkiEzsigntemplatepackagesignerID gets a reference to the given int32 and assigns it to the PkiEzsigntemplatepackagesignerID field.
func (o *EzsigntemplatepackagesignerRequest) SetPkiEzsigntemplatepackagesignerID(v int32) {
	o.PkiEzsigntemplatepackagesignerID = &v
}

// GetFkiEzsigntemplatepackageID returns the FkiEzsigntemplatepackageID field value
func (o *EzsigntemplatepackagesignerRequest) GetFkiEzsigntemplatepackageID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzsigntemplatepackageID
}

// GetFkiEzsigntemplatepackageIDOk returns a tuple with the FkiEzsigntemplatepackageID field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepackagesignerRequest) GetFkiEzsigntemplatepackageIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiEzsigntemplatepackageID, true
}

// SetFkiEzsigntemplatepackageID sets field value
func (o *EzsigntemplatepackagesignerRequest) SetFkiEzsigntemplatepackageID(v int32) {
	o.FkiEzsigntemplatepackageID = v
}

// GetSEzsigntemplatepackagesignerDescription returns the SEzsigntemplatepackagesignerDescription field value
func (o *EzsigntemplatepackagesignerRequest) GetSEzsigntemplatepackagesignerDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzsigntemplatepackagesignerDescription
}

// GetSEzsigntemplatepackagesignerDescriptionOk returns a tuple with the SEzsigntemplatepackagesignerDescription field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepackagesignerRequest) GetSEzsigntemplatepackagesignerDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SEzsigntemplatepackagesignerDescription, true
}

// SetSEzsigntemplatepackagesignerDescription sets field value
func (o *EzsigntemplatepackagesignerRequest) SetSEzsigntemplatepackagesignerDescription(v string) {
	o.SEzsigntemplatepackagesignerDescription = v
}

func (o EzsigntemplatepackagesignerRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigntemplatepackagesignerRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PkiEzsigntemplatepackagesignerID) {
		toSerialize["pkiEzsigntemplatepackagesignerID"] = o.PkiEzsigntemplatepackagesignerID
	}
	toSerialize["fkiEzsigntemplatepackageID"] = o.FkiEzsigntemplatepackageID
	toSerialize["sEzsigntemplatepackagesignerDescription"] = o.SEzsigntemplatepackagesignerDescription
	return toSerialize, nil
}

func (o *EzsigntemplatepackagesignerRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fkiEzsigntemplatepackageID",
		"sEzsigntemplatepackagesignerDescription",
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

	varEzsigntemplatepackagesignerRequest := _EzsigntemplatepackagesignerRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsigntemplatepackagesignerRequest)

	if err != nil {
		return err
	}

	*o = EzsigntemplatepackagesignerRequest(varEzsigntemplatepackagesignerRequest)

	return err
}

type NullableEzsigntemplatepackagesignerRequest struct {
	value *EzsigntemplatepackagesignerRequest
	isSet bool
}

func (v NullableEzsigntemplatepackagesignerRequest) Get() *EzsigntemplatepackagesignerRequest {
	return v.value
}

func (v *NullableEzsigntemplatepackagesignerRequest) Set(val *EzsigntemplatepackagesignerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigntemplatepackagesignerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigntemplatepackagesignerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigntemplatepackagesignerRequest(val *EzsigntemplatepackagesignerRequest) *NullableEzsigntemplatepackagesignerRequest {
	return &NullableEzsigntemplatepackagesignerRequest{value: val, isSet: true}
}

func (v NullableEzsigntemplatepackagesignerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigntemplatepackagesignerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


