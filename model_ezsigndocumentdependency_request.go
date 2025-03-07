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

// checks if the EzsigndocumentdependencyRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigndocumentdependencyRequest{}

// EzsigndocumentdependencyRequest An Ezsigndocumentdependency Object
type EzsigndocumentdependencyRequest struct {
	// The unique ID of the Ezsigndocumentdependency
	PkiEzsigndocumentdependencyID *int32 `json:"pkiEzsigndocumentdependencyID,omitempty"`
	// The unique ID of the Ezsigndocument
	FkiEzsigndocumentIDdependency int32 `json:"fkiEzsigndocumentIDdependency"`
}

type _EzsigndocumentdependencyRequest EzsigndocumentdependencyRequest

// NewEzsigndocumentdependencyRequest instantiates a new EzsigndocumentdependencyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigndocumentdependencyRequest(fkiEzsigndocumentIDdependency int32) *EzsigndocumentdependencyRequest {
	this := EzsigndocumentdependencyRequest{}
	this.FkiEzsigndocumentIDdependency = fkiEzsigndocumentIDdependency
	return &this
}

// NewEzsigndocumentdependencyRequestWithDefaults instantiates a new EzsigndocumentdependencyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigndocumentdependencyRequestWithDefaults() *EzsigndocumentdependencyRequest {
	this := EzsigndocumentdependencyRequest{}
	return &this
}

// GetPkiEzsigndocumentdependencyID returns the PkiEzsigndocumentdependencyID field value if set, zero value otherwise.
func (o *EzsigndocumentdependencyRequest) GetPkiEzsigndocumentdependencyID() int32 {
	if o == nil || IsNil(o.PkiEzsigndocumentdependencyID) {
		var ret int32
		return ret
	}
	return *o.PkiEzsigndocumentdependencyID
}

// GetPkiEzsigndocumentdependencyIDOk returns a tuple with the PkiEzsigndocumentdependencyID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigndocumentdependencyRequest) GetPkiEzsigndocumentdependencyIDOk() (*int32, bool) {
	if o == nil || IsNil(o.PkiEzsigndocumentdependencyID) {
		return nil, false
	}
	return o.PkiEzsigndocumentdependencyID, true
}

// HasPkiEzsigndocumentdependencyID returns a boolean if a field has been set.
func (o *EzsigndocumentdependencyRequest) HasPkiEzsigndocumentdependencyID() bool {
	if o != nil && !IsNil(o.PkiEzsigndocumentdependencyID) {
		return true
	}

	return false
}

// SetPkiEzsigndocumentdependencyID gets a reference to the given int32 and assigns it to the PkiEzsigndocumentdependencyID field.
func (o *EzsigndocumentdependencyRequest) SetPkiEzsigndocumentdependencyID(v int32) {
	o.PkiEzsigndocumentdependencyID = &v
}

// GetFkiEzsigndocumentIDdependency returns the FkiEzsigndocumentIDdependency field value
func (o *EzsigndocumentdependencyRequest) GetFkiEzsigndocumentIDdependency() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzsigndocumentIDdependency
}

// GetFkiEzsigndocumentIDdependencyOk returns a tuple with the FkiEzsigndocumentIDdependency field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentdependencyRequest) GetFkiEzsigndocumentIDdependencyOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiEzsigndocumentIDdependency, true
}

// SetFkiEzsigndocumentIDdependency sets field value
func (o *EzsigndocumentdependencyRequest) SetFkiEzsigndocumentIDdependency(v int32) {
	o.FkiEzsigndocumentIDdependency = v
}

func (o EzsigndocumentdependencyRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigndocumentdependencyRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PkiEzsigndocumentdependencyID) {
		toSerialize["pkiEzsigndocumentdependencyID"] = o.PkiEzsigndocumentdependencyID
	}
	toSerialize["fkiEzsigndocumentIDdependency"] = o.FkiEzsigndocumentIDdependency
	return toSerialize, nil
}

func (o *EzsigndocumentdependencyRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fkiEzsigndocumentIDdependency",
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

	varEzsigndocumentdependencyRequest := _EzsigndocumentdependencyRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsigndocumentdependencyRequest)

	if err != nil {
		return err
	}

	*o = EzsigndocumentdependencyRequest(varEzsigndocumentdependencyRequest)

	return err
}

type NullableEzsigndocumentdependencyRequest struct {
	value *EzsigndocumentdependencyRequest
	isSet bool
}

func (v NullableEzsigndocumentdependencyRequest) Get() *EzsigndocumentdependencyRequest {
	return v.value
}

func (v *NullableEzsigndocumentdependencyRequest) Set(val *EzsigndocumentdependencyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigndocumentdependencyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigndocumentdependencyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigndocumentdependencyRequest(val *EzsigndocumentdependencyRequest) *NullableEzsigndocumentdependencyRequest {
	return &NullableEzsigndocumentdependencyRequest{value: val, isSet: true}
}

func (v NullableEzsigndocumentdependencyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigndocumentdependencyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


