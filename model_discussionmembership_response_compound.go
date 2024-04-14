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

// checks if the DiscussionmembershipResponseCompound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DiscussionmembershipResponseCompound{}

// DiscussionmembershipResponseCompound A Discussionmembership Object and children
type DiscussionmembershipResponseCompound struct {
	// The unique ID of the Discussionmembership
	PkiDiscussionmembershipID int32 `json:"pkiDiscussionmembershipID"`
	// The unique ID of the Discussion
	FkiDiscussionID int32 `json:"fkiDiscussionID"`
	// The unique ID of the User
	FkiUserID *int32 `json:"fkiUserID,omitempty"`
	// The unique ID of the Usergroup
	FkiUsergroupID *int32 `json:"fkiUsergroupID,omitempty"`
	// The unique ID of the Modulesection
	FkiModulesectionID *int32 `json:"fkiModulesectionID,omitempty"`
	// The Description containing the detail of who the Discussionmembership refers to
	SDiscussionmembershipDescription string `json:"sDiscussionmembershipDescription"`
	// The joined date of the Discussionmembership
	DtDiscussionmembershipJoined string `json:"dtDiscussionmembershipJoined"`
}

type _DiscussionmembershipResponseCompound DiscussionmembershipResponseCompound

// NewDiscussionmembershipResponseCompound instantiates a new DiscussionmembershipResponseCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiscussionmembershipResponseCompound(pkiDiscussionmembershipID int32, fkiDiscussionID int32, sDiscussionmembershipDescription string, dtDiscussionmembershipJoined string) *DiscussionmembershipResponseCompound {
	this := DiscussionmembershipResponseCompound{}
	this.PkiDiscussionmembershipID = pkiDiscussionmembershipID
	this.FkiDiscussionID = fkiDiscussionID
	this.SDiscussionmembershipDescription = sDiscussionmembershipDescription
	this.DtDiscussionmembershipJoined = dtDiscussionmembershipJoined
	return &this
}

// NewDiscussionmembershipResponseCompoundWithDefaults instantiates a new DiscussionmembershipResponseCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiscussionmembershipResponseCompoundWithDefaults() *DiscussionmembershipResponseCompound {
	this := DiscussionmembershipResponseCompound{}
	return &this
}

// GetPkiDiscussionmembershipID returns the PkiDiscussionmembershipID field value
func (o *DiscussionmembershipResponseCompound) GetPkiDiscussionmembershipID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiDiscussionmembershipID
}

// GetPkiDiscussionmembershipIDOk returns a tuple with the PkiDiscussionmembershipID field value
// and a boolean to check if the value has been set.
func (o *DiscussionmembershipResponseCompound) GetPkiDiscussionmembershipIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiDiscussionmembershipID, true
}

// SetPkiDiscussionmembershipID sets field value
func (o *DiscussionmembershipResponseCompound) SetPkiDiscussionmembershipID(v int32) {
	o.PkiDiscussionmembershipID = v
}

// GetFkiDiscussionID returns the FkiDiscussionID field value
func (o *DiscussionmembershipResponseCompound) GetFkiDiscussionID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiDiscussionID
}

// GetFkiDiscussionIDOk returns a tuple with the FkiDiscussionID field value
// and a boolean to check if the value has been set.
func (o *DiscussionmembershipResponseCompound) GetFkiDiscussionIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiDiscussionID, true
}

// SetFkiDiscussionID sets field value
func (o *DiscussionmembershipResponseCompound) SetFkiDiscussionID(v int32) {
	o.FkiDiscussionID = v
}

// GetFkiUserID returns the FkiUserID field value if set, zero value otherwise.
func (o *DiscussionmembershipResponseCompound) GetFkiUserID() int32 {
	if o == nil || IsNil(o.FkiUserID) {
		var ret int32
		return ret
	}
	return *o.FkiUserID
}

// GetFkiUserIDOk returns a tuple with the FkiUserID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscussionmembershipResponseCompound) GetFkiUserIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiUserID) {
		return nil, false
	}
	return o.FkiUserID, true
}

// HasFkiUserID returns a boolean if a field has been set.
func (o *DiscussionmembershipResponseCompound) HasFkiUserID() bool {
	if o != nil && !IsNil(o.FkiUserID) {
		return true
	}

	return false
}

// SetFkiUserID gets a reference to the given int32 and assigns it to the FkiUserID field.
func (o *DiscussionmembershipResponseCompound) SetFkiUserID(v int32) {
	o.FkiUserID = &v
}

// GetFkiUsergroupID returns the FkiUsergroupID field value if set, zero value otherwise.
func (o *DiscussionmembershipResponseCompound) GetFkiUsergroupID() int32 {
	if o == nil || IsNil(o.FkiUsergroupID) {
		var ret int32
		return ret
	}
	return *o.FkiUsergroupID
}

// GetFkiUsergroupIDOk returns a tuple with the FkiUsergroupID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscussionmembershipResponseCompound) GetFkiUsergroupIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiUsergroupID) {
		return nil, false
	}
	return o.FkiUsergroupID, true
}

// HasFkiUsergroupID returns a boolean if a field has been set.
func (o *DiscussionmembershipResponseCompound) HasFkiUsergroupID() bool {
	if o != nil && !IsNil(o.FkiUsergroupID) {
		return true
	}

	return false
}

// SetFkiUsergroupID gets a reference to the given int32 and assigns it to the FkiUsergroupID field.
func (o *DiscussionmembershipResponseCompound) SetFkiUsergroupID(v int32) {
	o.FkiUsergroupID = &v
}

// GetFkiModulesectionID returns the FkiModulesectionID field value if set, zero value otherwise.
func (o *DiscussionmembershipResponseCompound) GetFkiModulesectionID() int32 {
	if o == nil || IsNil(o.FkiModulesectionID) {
		var ret int32
		return ret
	}
	return *o.FkiModulesectionID
}

// GetFkiModulesectionIDOk returns a tuple with the FkiModulesectionID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscussionmembershipResponseCompound) GetFkiModulesectionIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiModulesectionID) {
		return nil, false
	}
	return o.FkiModulesectionID, true
}

// HasFkiModulesectionID returns a boolean if a field has been set.
func (o *DiscussionmembershipResponseCompound) HasFkiModulesectionID() bool {
	if o != nil && !IsNil(o.FkiModulesectionID) {
		return true
	}

	return false
}

// SetFkiModulesectionID gets a reference to the given int32 and assigns it to the FkiModulesectionID field.
func (o *DiscussionmembershipResponseCompound) SetFkiModulesectionID(v int32) {
	o.FkiModulesectionID = &v
}

// GetSDiscussionmembershipDescription returns the SDiscussionmembershipDescription field value
func (o *DiscussionmembershipResponseCompound) GetSDiscussionmembershipDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SDiscussionmembershipDescription
}

// GetSDiscussionmembershipDescriptionOk returns a tuple with the SDiscussionmembershipDescription field value
// and a boolean to check if the value has been set.
func (o *DiscussionmembershipResponseCompound) GetSDiscussionmembershipDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SDiscussionmembershipDescription, true
}

// SetSDiscussionmembershipDescription sets field value
func (o *DiscussionmembershipResponseCompound) SetSDiscussionmembershipDescription(v string) {
	o.SDiscussionmembershipDescription = v
}

// GetDtDiscussionmembershipJoined returns the DtDiscussionmembershipJoined field value
func (o *DiscussionmembershipResponseCompound) GetDtDiscussionmembershipJoined() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DtDiscussionmembershipJoined
}

// GetDtDiscussionmembershipJoinedOk returns a tuple with the DtDiscussionmembershipJoined field value
// and a boolean to check if the value has been set.
func (o *DiscussionmembershipResponseCompound) GetDtDiscussionmembershipJoinedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DtDiscussionmembershipJoined, true
}

// SetDtDiscussionmembershipJoined sets field value
func (o *DiscussionmembershipResponseCompound) SetDtDiscussionmembershipJoined(v string) {
	o.DtDiscussionmembershipJoined = v
}

func (o DiscussionmembershipResponseCompound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DiscussionmembershipResponseCompound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiDiscussionmembershipID"] = o.PkiDiscussionmembershipID
	toSerialize["fkiDiscussionID"] = o.FkiDiscussionID
	if !IsNil(o.FkiUserID) {
		toSerialize["fkiUserID"] = o.FkiUserID
	}
	if !IsNil(o.FkiUsergroupID) {
		toSerialize["fkiUsergroupID"] = o.FkiUsergroupID
	}
	if !IsNil(o.FkiModulesectionID) {
		toSerialize["fkiModulesectionID"] = o.FkiModulesectionID
	}
	toSerialize["sDiscussionmembershipDescription"] = o.SDiscussionmembershipDescription
	toSerialize["dtDiscussionmembershipJoined"] = o.DtDiscussionmembershipJoined
	return toSerialize, nil
}

func (o *DiscussionmembershipResponseCompound) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiDiscussionmembershipID",
		"fkiDiscussionID",
		"sDiscussionmembershipDescription",
		"dtDiscussionmembershipJoined",
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

	varDiscussionmembershipResponseCompound := _DiscussionmembershipResponseCompound{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDiscussionmembershipResponseCompound)

	if err != nil {
		return err
	}

	*o = DiscussionmembershipResponseCompound(varDiscussionmembershipResponseCompound)

	return err
}

type NullableDiscussionmembershipResponseCompound struct {
	value *DiscussionmembershipResponseCompound
	isSet bool
}

func (v NullableDiscussionmembershipResponseCompound) Get() *DiscussionmembershipResponseCompound {
	return v.value
}

func (v *NullableDiscussionmembershipResponseCompound) Set(val *DiscussionmembershipResponseCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableDiscussionmembershipResponseCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableDiscussionmembershipResponseCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiscussionmembershipResponseCompound(val *DiscussionmembershipResponseCompound) *NullableDiscussionmembershipResponseCompound {
	return &NullableDiscussionmembershipResponseCompound{value: val, isSet: true}
}

func (v NullableDiscussionmembershipResponseCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiscussionmembershipResponseCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


