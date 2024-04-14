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

// checks if the UsergroupmembershipRequestCompound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UsergroupmembershipRequestCompound{}

// UsergroupmembershipRequestCompound A Usergroupmembership Object and children
type UsergroupmembershipRequestCompound struct {
	// The unique ID of the Usergroupmembership
	PkiUsergroupmembershipID *int32 `json:"pkiUsergroupmembershipID,omitempty"`
	// The unique ID of the Usergroup
	FkiUsergroupID int32 `json:"fkiUsergroupID"`
	// The unique ID of the User
	FkiUserID *int32 `json:"fkiUserID,omitempty"`
	// The unique ID of the Usergroupexternal
	FkiUsergroupexternalID *int32 `json:"fkiUsergroupexternalID,omitempty"`
}

type _UsergroupmembershipRequestCompound UsergroupmembershipRequestCompound

// NewUsergroupmembershipRequestCompound instantiates a new UsergroupmembershipRequestCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsergroupmembershipRequestCompound(fkiUsergroupID int32) *UsergroupmembershipRequestCompound {
	this := UsergroupmembershipRequestCompound{}
	this.FkiUsergroupID = fkiUsergroupID
	return &this
}

// NewUsergroupmembershipRequestCompoundWithDefaults instantiates a new UsergroupmembershipRequestCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsergroupmembershipRequestCompoundWithDefaults() *UsergroupmembershipRequestCompound {
	this := UsergroupmembershipRequestCompound{}
	return &this
}

// GetPkiUsergroupmembershipID returns the PkiUsergroupmembershipID field value if set, zero value otherwise.
func (o *UsergroupmembershipRequestCompound) GetPkiUsergroupmembershipID() int32 {
	if o == nil || IsNil(o.PkiUsergroupmembershipID) {
		var ret int32
		return ret
	}
	return *o.PkiUsergroupmembershipID
}

// GetPkiUsergroupmembershipIDOk returns a tuple with the PkiUsergroupmembershipID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsergroupmembershipRequestCompound) GetPkiUsergroupmembershipIDOk() (*int32, bool) {
	if o == nil || IsNil(o.PkiUsergroupmembershipID) {
		return nil, false
	}
	return o.PkiUsergroupmembershipID, true
}

// HasPkiUsergroupmembershipID returns a boolean if a field has been set.
func (o *UsergroupmembershipRequestCompound) HasPkiUsergroupmembershipID() bool {
	if o != nil && !IsNil(o.PkiUsergroupmembershipID) {
		return true
	}

	return false
}

// SetPkiUsergroupmembershipID gets a reference to the given int32 and assigns it to the PkiUsergroupmembershipID field.
func (o *UsergroupmembershipRequestCompound) SetPkiUsergroupmembershipID(v int32) {
	o.PkiUsergroupmembershipID = &v
}

// GetFkiUsergroupID returns the FkiUsergroupID field value
func (o *UsergroupmembershipRequestCompound) GetFkiUsergroupID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiUsergroupID
}

// GetFkiUsergroupIDOk returns a tuple with the FkiUsergroupID field value
// and a boolean to check if the value has been set.
func (o *UsergroupmembershipRequestCompound) GetFkiUsergroupIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiUsergroupID, true
}

// SetFkiUsergroupID sets field value
func (o *UsergroupmembershipRequestCompound) SetFkiUsergroupID(v int32) {
	o.FkiUsergroupID = v
}

// GetFkiUserID returns the FkiUserID field value if set, zero value otherwise.
func (o *UsergroupmembershipRequestCompound) GetFkiUserID() int32 {
	if o == nil || IsNil(o.FkiUserID) {
		var ret int32
		return ret
	}
	return *o.FkiUserID
}

// GetFkiUserIDOk returns a tuple with the FkiUserID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsergroupmembershipRequestCompound) GetFkiUserIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiUserID) {
		return nil, false
	}
	return o.FkiUserID, true
}

// HasFkiUserID returns a boolean if a field has been set.
func (o *UsergroupmembershipRequestCompound) HasFkiUserID() bool {
	if o != nil && !IsNil(o.FkiUserID) {
		return true
	}

	return false
}

// SetFkiUserID gets a reference to the given int32 and assigns it to the FkiUserID field.
func (o *UsergroupmembershipRequestCompound) SetFkiUserID(v int32) {
	o.FkiUserID = &v
}

// GetFkiUsergroupexternalID returns the FkiUsergroupexternalID field value if set, zero value otherwise.
func (o *UsergroupmembershipRequestCompound) GetFkiUsergroupexternalID() int32 {
	if o == nil || IsNil(o.FkiUsergroupexternalID) {
		var ret int32
		return ret
	}
	return *o.FkiUsergroupexternalID
}

// GetFkiUsergroupexternalIDOk returns a tuple with the FkiUsergroupexternalID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsergroupmembershipRequestCompound) GetFkiUsergroupexternalIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiUsergroupexternalID) {
		return nil, false
	}
	return o.FkiUsergroupexternalID, true
}

// HasFkiUsergroupexternalID returns a boolean if a field has been set.
func (o *UsergroupmembershipRequestCompound) HasFkiUsergroupexternalID() bool {
	if o != nil && !IsNil(o.FkiUsergroupexternalID) {
		return true
	}

	return false
}

// SetFkiUsergroupexternalID gets a reference to the given int32 and assigns it to the FkiUsergroupexternalID field.
func (o *UsergroupmembershipRequestCompound) SetFkiUsergroupexternalID(v int32) {
	o.FkiUsergroupexternalID = &v
}

func (o UsergroupmembershipRequestCompound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsergroupmembershipRequestCompound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PkiUsergroupmembershipID) {
		toSerialize["pkiUsergroupmembershipID"] = o.PkiUsergroupmembershipID
	}
	toSerialize["fkiUsergroupID"] = o.FkiUsergroupID
	if !IsNil(o.FkiUserID) {
		toSerialize["fkiUserID"] = o.FkiUserID
	}
	if !IsNil(o.FkiUsergroupexternalID) {
		toSerialize["fkiUsergroupexternalID"] = o.FkiUsergroupexternalID
	}
	return toSerialize, nil
}

func (o *UsergroupmembershipRequestCompound) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fkiUsergroupID",
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

	varUsergroupmembershipRequestCompound := _UsergroupmembershipRequestCompound{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUsergroupmembershipRequestCompound)

	if err != nil {
		return err
	}

	*o = UsergroupmembershipRequestCompound(varUsergroupmembershipRequestCompound)

	return err
}

type NullableUsergroupmembershipRequestCompound struct {
	value *UsergroupmembershipRequestCompound
	isSet bool
}

func (v NullableUsergroupmembershipRequestCompound) Get() *UsergroupmembershipRequestCompound {
	return v.value
}

func (v *NullableUsergroupmembershipRequestCompound) Set(val *UsergroupmembershipRequestCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableUsergroupmembershipRequestCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableUsergroupmembershipRequestCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsergroupmembershipRequestCompound(val *UsergroupmembershipRequestCompound) *NullableUsergroupmembershipRequestCompound {
	return &NullableUsergroupmembershipRequestCompound{value: val, isSet: true}
}

func (v NullableUsergroupmembershipRequestCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsergroupmembershipRequestCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


