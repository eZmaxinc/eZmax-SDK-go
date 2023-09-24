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

// checks if the EzsignsignergroupmembershipRequestCompound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignsignergroupmembershipRequestCompound{}

// EzsignsignergroupmembershipRequestCompound A Ezsignsignergroupmembership Object and children
type EzsignsignergroupmembershipRequestCompound struct {
	// The unique ID of the Ezsignsignergroupmembership
	PkiEzsignsignergroupmembershipID *int32 `json:"pkiEzsignsignergroupmembershipID,omitempty"`
	// The unique ID of the Ezsignsignergroup
	FkiEzsignsignergroupID int32 `json:"fkiEzsignsignergroupID"`
	// The unique ID of the Ezsignsigner
	FkiEzsignsignerID *int32 `json:"fkiEzsignsignerID,omitempty"`
	// The unique ID of the User
	FkiUserID *int32 `json:"fkiUserID,omitempty"`
	// The unique ID of the Usergroup
	FkiUsergroupID *int32 `json:"fkiUsergroupID,omitempty"`
}

// NewEzsignsignergroupmembershipRequestCompound instantiates a new EzsignsignergroupmembershipRequestCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignsignergroupmembershipRequestCompound(fkiEzsignsignergroupID int32) *EzsignsignergroupmembershipRequestCompound {
	this := EzsignsignergroupmembershipRequestCompound{}
	this.FkiEzsignsignergroupID = fkiEzsignsignergroupID
	return &this
}

// NewEzsignsignergroupmembershipRequestCompoundWithDefaults instantiates a new EzsignsignergroupmembershipRequestCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignsignergroupmembershipRequestCompoundWithDefaults() *EzsignsignergroupmembershipRequestCompound {
	this := EzsignsignergroupmembershipRequestCompound{}
	return &this
}

// GetPkiEzsignsignergroupmembershipID returns the PkiEzsignsignergroupmembershipID field value if set, zero value otherwise.
func (o *EzsignsignergroupmembershipRequestCompound) GetPkiEzsignsignergroupmembershipID() int32 {
	if o == nil || IsNil(o.PkiEzsignsignergroupmembershipID) {
		var ret int32
		return ret
	}
	return *o.PkiEzsignsignergroupmembershipID
}

// GetPkiEzsignsignergroupmembershipIDOk returns a tuple with the PkiEzsignsignergroupmembershipID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignsignergroupmembershipRequestCompound) GetPkiEzsignsignergroupmembershipIDOk() (*int32, bool) {
	if o == nil || IsNil(o.PkiEzsignsignergroupmembershipID) {
		return nil, false
	}
	return o.PkiEzsignsignergroupmembershipID, true
}

// HasPkiEzsignsignergroupmembershipID returns a boolean if a field has been set.
func (o *EzsignsignergroupmembershipRequestCompound) HasPkiEzsignsignergroupmembershipID() bool {
	if o != nil && !IsNil(o.PkiEzsignsignergroupmembershipID) {
		return true
	}

	return false
}

// SetPkiEzsignsignergroupmembershipID gets a reference to the given int32 and assigns it to the PkiEzsignsignergroupmembershipID field.
func (o *EzsignsignergroupmembershipRequestCompound) SetPkiEzsignsignergroupmembershipID(v int32) {
	o.PkiEzsignsignergroupmembershipID = &v
}

// GetFkiEzsignsignergroupID returns the FkiEzsignsignergroupID field value
func (o *EzsignsignergroupmembershipRequestCompound) GetFkiEzsignsignergroupID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzsignsignergroupID
}

// GetFkiEzsignsignergroupIDOk returns a tuple with the FkiEzsignsignergroupID field value
// and a boolean to check if the value has been set.
func (o *EzsignsignergroupmembershipRequestCompound) GetFkiEzsignsignergroupIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiEzsignsignergroupID, true
}

// SetFkiEzsignsignergroupID sets field value
func (o *EzsignsignergroupmembershipRequestCompound) SetFkiEzsignsignergroupID(v int32) {
	o.FkiEzsignsignergroupID = v
}

// GetFkiEzsignsignerID returns the FkiEzsignsignerID field value if set, zero value otherwise.
func (o *EzsignsignergroupmembershipRequestCompound) GetFkiEzsignsignerID() int32 {
	if o == nil || IsNil(o.FkiEzsignsignerID) {
		var ret int32
		return ret
	}
	return *o.FkiEzsignsignerID
}

// GetFkiEzsignsignerIDOk returns a tuple with the FkiEzsignsignerID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignsignergroupmembershipRequestCompound) GetFkiEzsignsignerIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiEzsignsignerID) {
		return nil, false
	}
	return o.FkiEzsignsignerID, true
}

// HasFkiEzsignsignerID returns a boolean if a field has been set.
func (o *EzsignsignergroupmembershipRequestCompound) HasFkiEzsignsignerID() bool {
	if o != nil && !IsNil(o.FkiEzsignsignerID) {
		return true
	}

	return false
}

// SetFkiEzsignsignerID gets a reference to the given int32 and assigns it to the FkiEzsignsignerID field.
func (o *EzsignsignergroupmembershipRequestCompound) SetFkiEzsignsignerID(v int32) {
	o.FkiEzsignsignerID = &v
}

// GetFkiUserID returns the FkiUserID field value if set, zero value otherwise.
func (o *EzsignsignergroupmembershipRequestCompound) GetFkiUserID() int32 {
	if o == nil || IsNil(o.FkiUserID) {
		var ret int32
		return ret
	}
	return *o.FkiUserID
}

// GetFkiUserIDOk returns a tuple with the FkiUserID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignsignergroupmembershipRequestCompound) GetFkiUserIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiUserID) {
		return nil, false
	}
	return o.FkiUserID, true
}

// HasFkiUserID returns a boolean if a field has been set.
func (o *EzsignsignergroupmembershipRequestCompound) HasFkiUserID() bool {
	if o != nil && !IsNil(o.FkiUserID) {
		return true
	}

	return false
}

// SetFkiUserID gets a reference to the given int32 and assigns it to the FkiUserID field.
func (o *EzsignsignergroupmembershipRequestCompound) SetFkiUserID(v int32) {
	o.FkiUserID = &v
}

// GetFkiUsergroupID returns the FkiUsergroupID field value if set, zero value otherwise.
func (o *EzsignsignergroupmembershipRequestCompound) GetFkiUsergroupID() int32 {
	if o == nil || IsNil(o.FkiUsergroupID) {
		var ret int32
		return ret
	}
	return *o.FkiUsergroupID
}

// GetFkiUsergroupIDOk returns a tuple with the FkiUsergroupID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignsignergroupmembershipRequestCompound) GetFkiUsergroupIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiUsergroupID) {
		return nil, false
	}
	return o.FkiUsergroupID, true
}

// HasFkiUsergroupID returns a boolean if a field has been set.
func (o *EzsignsignergroupmembershipRequestCompound) HasFkiUsergroupID() bool {
	if o != nil && !IsNil(o.FkiUsergroupID) {
		return true
	}

	return false
}

// SetFkiUsergroupID gets a reference to the given int32 and assigns it to the FkiUsergroupID field.
func (o *EzsignsignergroupmembershipRequestCompound) SetFkiUsergroupID(v int32) {
	o.FkiUsergroupID = &v
}

func (o EzsignsignergroupmembershipRequestCompound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignsignergroupmembershipRequestCompound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PkiEzsignsignergroupmembershipID) {
		toSerialize["pkiEzsignsignergroupmembershipID"] = o.PkiEzsignsignergroupmembershipID
	}
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

type NullableEzsignsignergroupmembershipRequestCompound struct {
	value *EzsignsignergroupmembershipRequestCompound
	isSet bool
}

func (v NullableEzsignsignergroupmembershipRequestCompound) Get() *EzsignsignergroupmembershipRequestCompound {
	return v.value
}

func (v *NullableEzsignsignergroupmembershipRequestCompound) Set(val *EzsignsignergroupmembershipRequestCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignsignergroupmembershipRequestCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignsignergroupmembershipRequestCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignsignergroupmembershipRequestCompound(val *EzsignsignergroupmembershipRequestCompound) *NullableEzsignsignergroupmembershipRequestCompound {
	return &NullableEzsignsignergroupmembershipRequestCompound{value: val, isSet: true}
}

func (v NullableEzsignsignergroupmembershipRequestCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignsignergroupmembershipRequestCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


