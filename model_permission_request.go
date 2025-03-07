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

// checks if the PermissionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PermissionRequest{}

// PermissionRequest A Permission Object
type PermissionRequest struct {
	// The unique ID of the Permission
	PkiPermissionID *int32 `json:"pkiPermissionID,omitempty"`
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
}

type _PermissionRequest PermissionRequest

// NewPermissionRequest instantiates a new PermissionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPermissionRequest(fkiModulesectionID int32) *PermissionRequest {
	this := PermissionRequest{}
	this.FkiModulesectionID = fkiModulesectionID
	return &this
}

// NewPermissionRequestWithDefaults instantiates a new PermissionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPermissionRequestWithDefaults() *PermissionRequest {
	this := PermissionRequest{}
	return &this
}

// GetPkiPermissionID returns the PkiPermissionID field value if set, zero value otherwise.
func (o *PermissionRequest) GetPkiPermissionID() int32 {
	if o == nil || IsNil(o.PkiPermissionID) {
		var ret int32
		return ret
	}
	return *o.PkiPermissionID
}

// GetPkiPermissionIDOk returns a tuple with the PkiPermissionID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionRequest) GetPkiPermissionIDOk() (*int32, bool) {
	if o == nil || IsNil(o.PkiPermissionID) {
		return nil, false
	}
	return o.PkiPermissionID, true
}

// HasPkiPermissionID returns a boolean if a field has been set.
func (o *PermissionRequest) HasPkiPermissionID() bool {
	if o != nil && !IsNil(o.PkiPermissionID) {
		return true
	}

	return false
}

// SetPkiPermissionID gets a reference to the given int32 and assigns it to the PkiPermissionID field.
func (o *PermissionRequest) SetPkiPermissionID(v int32) {
	o.PkiPermissionID = &v
}

// GetFkiUserID returns the FkiUserID field value if set, zero value otherwise.
func (o *PermissionRequest) GetFkiUserID() int32 {
	if o == nil || IsNil(o.FkiUserID) {
		var ret int32
		return ret
	}
	return *o.FkiUserID
}

// GetFkiUserIDOk returns a tuple with the FkiUserID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionRequest) GetFkiUserIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiUserID) {
		return nil, false
	}
	return o.FkiUserID, true
}

// HasFkiUserID returns a boolean if a field has been set.
func (o *PermissionRequest) HasFkiUserID() bool {
	if o != nil && !IsNil(o.FkiUserID) {
		return true
	}

	return false
}

// SetFkiUserID gets a reference to the given int32 and assigns it to the FkiUserID field.
func (o *PermissionRequest) SetFkiUserID(v int32) {
	o.FkiUserID = &v
}

// GetFkiApikeyID returns the FkiApikeyID field value if set, zero value otherwise.
func (o *PermissionRequest) GetFkiApikeyID() int32 {
	if o == nil || IsNil(o.FkiApikeyID) {
		var ret int32
		return ret
	}
	return *o.FkiApikeyID
}

// GetFkiApikeyIDOk returns a tuple with the FkiApikeyID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionRequest) GetFkiApikeyIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiApikeyID) {
		return nil, false
	}
	return o.FkiApikeyID, true
}

// HasFkiApikeyID returns a boolean if a field has been set.
func (o *PermissionRequest) HasFkiApikeyID() bool {
	if o != nil && !IsNil(o.FkiApikeyID) {
		return true
	}

	return false
}

// SetFkiApikeyID gets a reference to the given int32 and assigns it to the FkiApikeyID field.
func (o *PermissionRequest) SetFkiApikeyID(v int32) {
	o.FkiApikeyID = &v
}

// GetFkiUsergroupID returns the FkiUsergroupID field value if set, zero value otherwise.
func (o *PermissionRequest) GetFkiUsergroupID() int32 {
	if o == nil || IsNil(o.FkiUsergroupID) {
		var ret int32
		return ret
	}
	return *o.FkiUsergroupID
}

// GetFkiUsergroupIDOk returns a tuple with the FkiUsergroupID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionRequest) GetFkiUsergroupIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiUsergroupID) {
		return nil, false
	}
	return o.FkiUsergroupID, true
}

// HasFkiUsergroupID returns a boolean if a field has been set.
func (o *PermissionRequest) HasFkiUsergroupID() bool {
	if o != nil && !IsNil(o.FkiUsergroupID) {
		return true
	}

	return false
}

// SetFkiUsergroupID gets a reference to the given int32 and assigns it to the FkiUsergroupID field.
func (o *PermissionRequest) SetFkiUsergroupID(v int32) {
	o.FkiUsergroupID = &v
}

// GetFkiCompanyID returns the FkiCompanyID field value if set, zero value otherwise.
func (o *PermissionRequest) GetFkiCompanyID() int32 {
	if o == nil || IsNil(o.FkiCompanyID) {
		var ret int32
		return ret
	}
	return *o.FkiCompanyID
}

// GetFkiCompanyIDOk returns a tuple with the FkiCompanyID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionRequest) GetFkiCompanyIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiCompanyID) {
		return nil, false
	}
	return o.FkiCompanyID, true
}

// HasFkiCompanyID returns a boolean if a field has been set.
func (o *PermissionRequest) HasFkiCompanyID() bool {
	if o != nil && !IsNil(o.FkiCompanyID) {
		return true
	}

	return false
}

// SetFkiCompanyID gets a reference to the given int32 and assigns it to the FkiCompanyID field.
func (o *PermissionRequest) SetFkiCompanyID(v int32) {
	o.FkiCompanyID = &v
}

// GetFkiModulesectionID returns the FkiModulesectionID field value
func (o *PermissionRequest) GetFkiModulesectionID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiModulesectionID
}

// GetFkiModulesectionIDOk returns a tuple with the FkiModulesectionID field value
// and a boolean to check if the value has been set.
func (o *PermissionRequest) GetFkiModulesectionIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiModulesectionID, true
}

// SetFkiModulesectionID sets field value
func (o *PermissionRequest) SetFkiModulesectionID(v int32) {
	o.FkiModulesectionID = v
}

func (o PermissionRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PermissionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PkiPermissionID) {
		toSerialize["pkiPermissionID"] = o.PkiPermissionID
	}
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
	return toSerialize, nil
}

func (o *PermissionRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varPermissionRequest := _PermissionRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPermissionRequest)

	if err != nil {
		return err
	}

	*o = PermissionRequest(varPermissionRequest)

	return err
}

type NullablePermissionRequest struct {
	value *PermissionRequest
	isSet bool
}

func (v NullablePermissionRequest) Get() *PermissionRequest {
	return v.value
}

func (v *NullablePermissionRequest) Set(val *PermissionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePermissionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePermissionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermissionRequest(val *PermissionRequest) *NullablePermissionRequest {
	return &NullablePermissionRequest{value: val, isSet: true}
}

func (v NullablePermissionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermissionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


