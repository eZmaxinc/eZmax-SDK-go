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

// checks if the EzsignsignergroupmembershipResponseCompound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignsignergroupmembershipResponseCompound{}

// EzsignsignergroupmembershipResponseCompound A Ezsignsignergroupmembership Object
type EzsignsignergroupmembershipResponseCompound struct {
	// The unique ID of the Ezsignsignergroupmembership
	PkiEzsignsignergroupmembershipID int32 `json:"pkiEzsignsignergroupmembershipID"`
	// The unique ID of the Ezsignsignergroup
	FkiEzsignsignergroupID int32 `json:"fkiEzsignsignergroupID"`
	// The unique ID of the Ezsignsigner
	FkiEzsignsignerID *int32 `json:"fkiEzsignsignerID,omitempty"`
	// The unique ID of the User
	FkiUserID *int32 `json:"fkiUserID,omitempty"`
	// The unique ID of the Usergroup
	FkiUsergroupID *int32 `json:"fkiUsergroupID,omitempty"`
}

type _EzsignsignergroupmembershipResponseCompound EzsignsignergroupmembershipResponseCompound

// NewEzsignsignergroupmembershipResponseCompound instantiates a new EzsignsignergroupmembershipResponseCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignsignergroupmembershipResponseCompound(pkiEzsignsignergroupmembershipID int32, fkiEzsignsignergroupID int32) *EzsignsignergroupmembershipResponseCompound {
	this := EzsignsignergroupmembershipResponseCompound{}
	this.PkiEzsignsignergroupmembershipID = pkiEzsignsignergroupmembershipID
	this.FkiEzsignsignergroupID = fkiEzsignsignergroupID
	return &this
}

// NewEzsignsignergroupmembershipResponseCompoundWithDefaults instantiates a new EzsignsignergroupmembershipResponseCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignsignergroupmembershipResponseCompoundWithDefaults() *EzsignsignergroupmembershipResponseCompound {
	this := EzsignsignergroupmembershipResponseCompound{}
	return &this
}

// GetPkiEzsignsignergroupmembershipID returns the PkiEzsignsignergroupmembershipID field value
func (o *EzsignsignergroupmembershipResponseCompound) GetPkiEzsignsignergroupmembershipID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiEzsignsignergroupmembershipID
}

// GetPkiEzsignsignergroupmembershipIDOk returns a tuple with the PkiEzsignsignergroupmembershipID field value
// and a boolean to check if the value has been set.
func (o *EzsignsignergroupmembershipResponseCompound) GetPkiEzsignsignergroupmembershipIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiEzsignsignergroupmembershipID, true
}

// SetPkiEzsignsignergroupmembershipID sets field value
func (o *EzsignsignergroupmembershipResponseCompound) SetPkiEzsignsignergroupmembershipID(v int32) {
	o.PkiEzsignsignergroupmembershipID = v
}

// GetFkiEzsignsignergroupID returns the FkiEzsignsignergroupID field value
func (o *EzsignsignergroupmembershipResponseCompound) GetFkiEzsignsignergroupID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzsignsignergroupID
}

// GetFkiEzsignsignergroupIDOk returns a tuple with the FkiEzsignsignergroupID field value
// and a boolean to check if the value has been set.
func (o *EzsignsignergroupmembershipResponseCompound) GetFkiEzsignsignergroupIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiEzsignsignergroupID, true
}

// SetFkiEzsignsignergroupID sets field value
func (o *EzsignsignergroupmembershipResponseCompound) SetFkiEzsignsignergroupID(v int32) {
	o.FkiEzsignsignergroupID = v
}

// GetFkiEzsignsignerID returns the FkiEzsignsignerID field value if set, zero value otherwise.
func (o *EzsignsignergroupmembershipResponseCompound) GetFkiEzsignsignerID() int32 {
	if o == nil || IsNil(o.FkiEzsignsignerID) {
		var ret int32
		return ret
	}
	return *o.FkiEzsignsignerID
}

// GetFkiEzsignsignerIDOk returns a tuple with the FkiEzsignsignerID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignsignergroupmembershipResponseCompound) GetFkiEzsignsignerIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiEzsignsignerID) {
		return nil, false
	}
	return o.FkiEzsignsignerID, true
}

// HasFkiEzsignsignerID returns a boolean if a field has been set.
func (o *EzsignsignergroupmembershipResponseCompound) HasFkiEzsignsignerID() bool {
	if o != nil && !IsNil(o.FkiEzsignsignerID) {
		return true
	}

	return false
}

// SetFkiEzsignsignerID gets a reference to the given int32 and assigns it to the FkiEzsignsignerID field.
func (o *EzsignsignergroupmembershipResponseCompound) SetFkiEzsignsignerID(v int32) {
	o.FkiEzsignsignerID = &v
}

// GetFkiUserID returns the FkiUserID field value if set, zero value otherwise.
func (o *EzsignsignergroupmembershipResponseCompound) GetFkiUserID() int32 {
	if o == nil || IsNil(o.FkiUserID) {
		var ret int32
		return ret
	}
	return *o.FkiUserID
}

// GetFkiUserIDOk returns a tuple with the FkiUserID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignsignergroupmembershipResponseCompound) GetFkiUserIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiUserID) {
		return nil, false
	}
	return o.FkiUserID, true
}

// HasFkiUserID returns a boolean if a field has been set.
func (o *EzsignsignergroupmembershipResponseCompound) HasFkiUserID() bool {
	if o != nil && !IsNil(o.FkiUserID) {
		return true
	}

	return false
}

// SetFkiUserID gets a reference to the given int32 and assigns it to the FkiUserID field.
func (o *EzsignsignergroupmembershipResponseCompound) SetFkiUserID(v int32) {
	o.FkiUserID = &v
}

// GetFkiUsergroupID returns the FkiUsergroupID field value if set, zero value otherwise.
func (o *EzsignsignergroupmembershipResponseCompound) GetFkiUsergroupID() int32 {
	if o == nil || IsNil(o.FkiUsergroupID) {
		var ret int32
		return ret
	}
	return *o.FkiUsergroupID
}

// GetFkiUsergroupIDOk returns a tuple with the FkiUsergroupID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignsignergroupmembershipResponseCompound) GetFkiUsergroupIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiUsergroupID) {
		return nil, false
	}
	return o.FkiUsergroupID, true
}

// HasFkiUsergroupID returns a boolean if a field has been set.
func (o *EzsignsignergroupmembershipResponseCompound) HasFkiUsergroupID() bool {
	if o != nil && !IsNil(o.FkiUsergroupID) {
		return true
	}

	return false
}

// SetFkiUsergroupID gets a reference to the given int32 and assigns it to the FkiUsergroupID field.
func (o *EzsignsignergroupmembershipResponseCompound) SetFkiUsergroupID(v int32) {
	o.FkiUsergroupID = &v
}

func (o EzsignsignergroupmembershipResponseCompound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignsignergroupmembershipResponseCompound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiEzsignsignergroupmembershipID"] = o.PkiEzsignsignergroupmembershipID
	toSerialize["fkiEzsignsignergroupID"] = o.FkiEzsignsignergroupID
	if !IsNil(o.FkiEzsignsignerID) {
		toSerialize["fkiEzsignsignerID"] = o.FkiEzsignsignerID
	}
	if !IsNil(o.FkiUserID) {
		toSerialize["fkiUserID"] = o.FkiUserID
	}
	if !IsNil(o.FkiUsergroupID) {
		toSerialize["fkiUsergroupID"] = o.FkiUsergroupID
	}
	return toSerialize, nil
}

func (o *EzsignsignergroupmembershipResponseCompound) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiEzsignsignergroupmembershipID",
		"fkiEzsignsignergroupID",
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

	varEzsignsignergroupmembershipResponseCompound := _EzsignsignergroupmembershipResponseCompound{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsignsignergroupmembershipResponseCompound)

	if err != nil {
		return err
	}

	*o = EzsignsignergroupmembershipResponseCompound(varEzsignsignergroupmembershipResponseCompound)

	return err
}

type NullableEzsignsignergroupmembershipResponseCompound struct {
	value *EzsignsignergroupmembershipResponseCompound
	isSet bool
}

func (v NullableEzsignsignergroupmembershipResponseCompound) Get() *EzsignsignergroupmembershipResponseCompound {
	return v.value
}

func (v *NullableEzsignsignergroupmembershipResponseCompound) Set(val *EzsignsignergroupmembershipResponseCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignsignergroupmembershipResponseCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignsignergroupmembershipResponseCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignsignergroupmembershipResponseCompound(val *EzsignsignergroupmembershipResponseCompound) *NullableEzsignsignergroupmembershipResponseCompound {
	return &NullableEzsignsignergroupmembershipResponseCompound{value: val, isSet: true}
}

func (v NullableEzsignsignergroupmembershipResponseCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignsignergroupmembershipResponseCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


