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

// checks if the PermissionResponseCompound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PermissionResponseCompound{}

// PermissionResponseCompound A Permission Object and children to create a complete structure
type PermissionResponseCompound struct {
	// The unique ID of the Permission
	PkiPermissionID int32 `json:"pkiPermissionID"`
	// The unique ID of the User
	FkiUserID *int32 `json:"fkiUserID,omitempty"`
	// The unique ID of the Apikey
	FkiApikeyID *int32 `json:"fkiApikeyID,omitempty"`
	// The unique ID of the Usergroup
	FkiUsergroupID *int32 `json:"fkiUsergroupID,omitempty"`
	// The unique ID of the Company
	FkiCompanyID *int32 `json:"fkiCompanyID,omitempty"`
	// The unique ID of the Modulesection
	FkiModulesectionID int32 `json:"fkiModulesectionID"`
	// The Name of the Company in the language of the requester
	SCompanyNameX *string `json:"sCompanyNameX,omitempty"`
}

type _PermissionResponseCompound PermissionResponseCompound

// NewPermissionResponseCompound instantiates a new PermissionResponseCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPermissionResponseCompound(pkiPermissionID int32, fkiModulesectionID int32) *PermissionResponseCompound {
	this := PermissionResponseCompound{}
	this.PkiPermissionID = pkiPermissionID
	this.FkiModulesectionID = fkiModulesectionID
	return &this
}

// NewPermissionResponseCompoundWithDefaults instantiates a new PermissionResponseCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPermissionResponseCompoundWithDefaults() *PermissionResponseCompound {
	this := PermissionResponseCompound{}
	return &this
}

// GetPkiPermissionID returns the PkiPermissionID field value
func (o *PermissionResponseCompound) GetPkiPermissionID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiPermissionID
}

// GetPkiPermissionIDOk returns a tuple with the PkiPermissionID field value
// and a boolean to check if the value has been set.
func (o *PermissionResponseCompound) GetPkiPermissionIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiPermissionID, true
}

// SetPkiPermissionID sets field value
func (o *PermissionResponseCompound) SetPkiPermissionID(v int32) {
	o.PkiPermissionID = v
}

// GetFkiUserID returns the FkiUserID field value if set, zero value otherwise.
func (o *PermissionResponseCompound) GetFkiUserID() int32 {
	if o == nil || IsNil(o.FkiUserID) {
		var ret int32
		return ret
	}
	return *o.FkiUserID
}

// GetFkiUserIDOk returns a tuple with the FkiUserID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionResponseCompound) GetFkiUserIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiUserID) {
		return nil, false
	}
	return o.FkiUserID, true
}

// HasFkiUserID returns a boolean if a field has been set.
func (o *PermissionResponseCompound) HasFkiUserID() bool {
	if o != nil && !IsNil(o.FkiUserID) {
		return true
	}

	return false
}

// SetFkiUserID gets a reference to the given int32 and assigns it to the FkiUserID field.
func (o *PermissionResponseCompound) SetFkiUserID(v int32) {
	o.FkiUserID = &v
}

// GetFkiApikeyID returns the FkiApikeyID field value if set, zero value otherwise.
func (o *PermissionResponseCompound) GetFkiApikeyID() int32 {
	if o == nil || IsNil(o.FkiApikeyID) {
		var ret int32
		return ret
	}
	return *o.FkiApikeyID
}

// GetFkiApikeyIDOk returns a tuple with the FkiApikeyID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionResponseCompound) GetFkiApikeyIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiApikeyID) {
		return nil, false
	}
	return o.FkiApikeyID, true
}

// HasFkiApikeyID returns a boolean if a field has been set.
func (o *PermissionResponseCompound) HasFkiApikeyID() bool {
	if o != nil && !IsNil(o.FkiApikeyID) {
		return true
	}

	return false
}

// SetFkiApikeyID gets a reference to the given int32 and assigns it to the FkiApikeyID field.
func (o *PermissionResponseCompound) SetFkiApikeyID(v int32) {
	o.FkiApikeyID = &v
}

// GetFkiUsergroupID returns the FkiUsergroupID field value if set, zero value otherwise.
func (o *PermissionResponseCompound) GetFkiUsergroupID() int32 {
	if o == nil || IsNil(o.FkiUsergroupID) {
		var ret int32
		return ret
	}
	return *o.FkiUsergroupID
}

// GetFkiUsergroupIDOk returns a tuple with the FkiUsergroupID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionResponseCompound) GetFkiUsergroupIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiUsergroupID) {
		return nil, false
	}
	return o.FkiUsergroupID, true
}

// HasFkiUsergroupID returns a boolean if a field has been set.
func (o *PermissionResponseCompound) HasFkiUsergroupID() bool {
	if o != nil && !IsNil(o.FkiUsergroupID) {
		return true
	}

	return false
}

// SetFkiUsergroupID gets a reference to the given int32 and assigns it to the FkiUsergroupID field.
func (o *PermissionResponseCompound) SetFkiUsergroupID(v int32) {
	o.FkiUsergroupID = &v
}

// GetFkiCompanyID returns the FkiCompanyID field value if set, zero value otherwise.
func (o *PermissionResponseCompound) GetFkiCompanyID() int32 {
	if o == nil || IsNil(o.FkiCompanyID) {
		var ret int32
		return ret
	}
	return *o.FkiCompanyID
}

// GetFkiCompanyIDOk returns a tuple with the FkiCompanyID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionResponseCompound) GetFkiCompanyIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiCompanyID) {
		return nil, false
	}
	return o.FkiCompanyID, true
}

// HasFkiCompanyID returns a boolean if a field has been set.
func (o *PermissionResponseCompound) HasFkiCompanyID() bool {
	if o != nil && !IsNil(o.FkiCompanyID) {
		return true
	}

	return false
}

// SetFkiCompanyID gets a reference to the given int32 and assigns it to the FkiCompanyID field.
func (o *PermissionResponseCompound) SetFkiCompanyID(v int32) {
	o.FkiCompanyID = &v
}

// GetFkiModulesectionID returns the FkiModulesectionID field value
func (o *PermissionResponseCompound) GetFkiModulesectionID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiModulesectionID
}

// GetFkiModulesectionIDOk returns a tuple with the FkiModulesectionID field value
// and a boolean to check if the value has been set.
func (o *PermissionResponseCompound) GetFkiModulesectionIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiModulesectionID, true
}

// SetFkiModulesectionID sets field value
func (o *PermissionResponseCompound) SetFkiModulesectionID(v int32) {
	o.FkiModulesectionID = v
}

// GetSCompanyNameX returns the SCompanyNameX field value if set, zero value otherwise.
func (o *PermissionResponseCompound) GetSCompanyNameX() string {
	if o == nil || IsNil(o.SCompanyNameX) {
		var ret string
		return ret
	}
	return *o.SCompanyNameX
}

// GetSCompanyNameXOk returns a tuple with the SCompanyNameX field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionResponseCompound) GetSCompanyNameXOk() (*string, bool) {
	if o == nil || IsNil(o.SCompanyNameX) {
		return nil, false
	}
	return o.SCompanyNameX, true
}

// HasSCompanyNameX returns a boolean if a field has been set.
func (o *PermissionResponseCompound) HasSCompanyNameX() bool {
	if o != nil && !IsNil(o.SCompanyNameX) {
		return true
	}

	return false
}

// SetSCompanyNameX gets a reference to the given string and assigns it to the SCompanyNameX field.
func (o *PermissionResponseCompound) SetSCompanyNameX(v string) {
	o.SCompanyNameX = &v
}

func (o PermissionResponseCompound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PermissionResponseCompound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiPermissionID"] = o.PkiPermissionID
	if !IsNil(o.FkiUserID) {
		toSerialize["fkiUserID"] = o.FkiUserID
	}
	if !IsNil(o.FkiApikeyID) {
		toSerialize["fkiApikeyID"] = o.FkiApikeyID
	}
	if !IsNil(o.FkiUsergroupID) {
		toSerialize["fkiUsergroupID"] = o.FkiUsergroupID
	}
	if !IsNil(o.FkiCompanyID) {
		toSerialize["fkiCompanyID"] = o.FkiCompanyID
	}
	toSerialize["fkiModulesectionID"] = o.FkiModulesectionID
	if !IsNil(o.SCompanyNameX) {
		toSerialize["sCompanyNameX"] = o.SCompanyNameX
	}
	return toSerialize, nil
}

func (o *PermissionResponseCompound) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiPermissionID",
		"fkiModulesectionID",
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

	varPermissionResponseCompound := _PermissionResponseCompound{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPermissionResponseCompound)

	if err != nil {
		return err
	}

	*o = PermissionResponseCompound(varPermissionResponseCompound)

	return err
}

type NullablePermissionResponseCompound struct {
	value *PermissionResponseCompound
	isSet bool
}

func (v NullablePermissionResponseCompound) Get() *PermissionResponseCompound {
	return v.value
}

func (v *NullablePermissionResponseCompound) Set(val *PermissionResponseCompound) {
	v.value = val
	v.isSet = true
}

func (v NullablePermissionResponseCompound) IsSet() bool {
	return v.isSet
}

func (v *NullablePermissionResponseCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermissionResponseCompound(val *PermissionResponseCompound) *NullablePermissionResponseCompound {
	return &NullablePermissionResponseCompound{value: val, isSet: true}
}

func (v NullablePermissionResponseCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermissionResponseCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


