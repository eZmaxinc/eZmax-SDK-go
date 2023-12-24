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

// checks if the EzsignsignergroupRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignsignergroupRequest{}

// EzsignsignergroupRequest A Ezsignsignergroup Object
type EzsignsignergroupRequest struct {
	// The unique ID of the Ezsignsignergroup
	PkiEzsignsignergroupID *int32 `json:"pkiEzsignsignergroupID,omitempty"`
	// The unique ID of the Ezsignfolder
	FkiEzsignfolderID int32 `json:"fkiEzsignfolderID"`
	ObjEzsignsignergroupDescription MultilingualEzsignsignergroupDescription `json:"objEzsignsignergroupDescription"`
}

type _EzsignsignergroupRequest EzsignsignergroupRequest

// NewEzsignsignergroupRequest instantiates a new EzsignsignergroupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignsignergroupRequest(fkiEzsignfolderID int32, objEzsignsignergroupDescription MultilingualEzsignsignergroupDescription) *EzsignsignergroupRequest {
	this := EzsignsignergroupRequest{}
	this.FkiEzsignfolderID = fkiEzsignfolderID
	this.ObjEzsignsignergroupDescription = objEzsignsignergroupDescription
	return &this
}

// NewEzsignsignergroupRequestWithDefaults instantiates a new EzsignsignergroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignsignergroupRequestWithDefaults() *EzsignsignergroupRequest {
	this := EzsignsignergroupRequest{}
	return &this
}

// GetPkiEzsignsignergroupID returns the PkiEzsignsignergroupID field value if set, zero value otherwise.
func (o *EzsignsignergroupRequest) GetPkiEzsignsignergroupID() int32 {
	if o == nil || IsNil(o.PkiEzsignsignergroupID) {
		var ret int32
		return ret
	}
	return *o.PkiEzsignsignergroupID
}

// GetPkiEzsignsignergroupIDOk returns a tuple with the PkiEzsignsignergroupID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignsignergroupRequest) GetPkiEzsignsignergroupIDOk() (*int32, bool) {
	if o == nil || IsNil(o.PkiEzsignsignergroupID) {
		return nil, false
	}
	return o.PkiEzsignsignergroupID, true
}

// HasPkiEzsignsignergroupID returns a boolean if a field has been set.
func (o *EzsignsignergroupRequest) HasPkiEzsignsignergroupID() bool {
	if o != nil && !IsNil(o.PkiEzsignsignergroupID) {
		return true
	}

	return false
}

// SetPkiEzsignsignergroupID gets a reference to the given int32 and assigns it to the PkiEzsignsignergroupID field.
func (o *EzsignsignergroupRequest) SetPkiEzsignsignergroupID(v int32) {
	o.PkiEzsignsignergroupID = &v
}

// GetFkiEzsignfolderID returns the FkiEzsignfolderID field value
func (o *EzsignsignergroupRequest) GetFkiEzsignfolderID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzsignfolderID
}

// GetFkiEzsignfolderIDOk returns a tuple with the FkiEzsignfolderID field value
// and a boolean to check if the value has been set.
func (o *EzsignsignergroupRequest) GetFkiEzsignfolderIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiEzsignfolderID, true
}

// SetFkiEzsignfolderID sets field value
func (o *EzsignsignergroupRequest) SetFkiEzsignfolderID(v int32) {
	o.FkiEzsignfolderID = v
}

// GetObjEzsignsignergroupDescription returns the ObjEzsignsignergroupDescription field value
func (o *EzsignsignergroupRequest) GetObjEzsignsignergroupDescription() MultilingualEzsignsignergroupDescription {
	if o == nil {
		var ret MultilingualEzsignsignergroupDescription
		return ret
	}

	return o.ObjEzsignsignergroupDescription
}

// GetObjEzsignsignergroupDescriptionOk returns a tuple with the ObjEzsignsignergroupDescription field value
// and a boolean to check if the value has been set.
func (o *EzsignsignergroupRequest) GetObjEzsignsignergroupDescriptionOk() (*MultilingualEzsignsignergroupDescription, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjEzsignsignergroupDescription, true
}

// SetObjEzsignsignergroupDescription sets field value
func (o *EzsignsignergroupRequest) SetObjEzsignsignergroupDescription(v MultilingualEzsignsignergroupDescription) {
	o.ObjEzsignsignergroupDescription = v
}

func (o EzsignsignergroupRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignsignergroupRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PkiEzsignsignergroupID) {
		toSerialize["pkiEzsignsignergroupID"] = o.PkiEzsignsignergroupID
	}
	toSerialize["fkiEzsignfolderID"] = o.FkiEzsignfolderID
	toSerialize["objEzsignsignergroupDescription"] = o.ObjEzsignsignergroupDescription
	return toSerialize, nil
}

func (o *EzsignsignergroupRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fkiEzsignfolderID",
		"objEzsignsignergroupDescription",
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

	varEzsignsignergroupRequest := _EzsignsignergroupRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsignsignergroupRequest)

	if err != nil {
		return err
	}

	*o = EzsignsignergroupRequest(varEzsignsignergroupRequest)

	return err
}

type NullableEzsignsignergroupRequest struct {
	value *EzsignsignergroupRequest
	isSet bool
}

func (v NullableEzsignsignergroupRequest) Get() *EzsignsignergroupRequest {
	return v.value
}

func (v *NullableEzsignsignergroupRequest) Set(val *EzsignsignergroupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignsignergroupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignsignergroupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignsignergroupRequest(val *EzsignsignergroupRequest) *NullableEzsignsignergroupRequest {
	return &NullableEzsignsignergroupRequest{value: val, isSet: true}
}

func (v NullableEzsignsignergroupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignsignergroupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


