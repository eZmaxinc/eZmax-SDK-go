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

// checks if the SubnetResponseCompound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubnetResponseCompound{}

// SubnetResponseCompound A Subnet Object
type SubnetResponseCompound struct {
	// The unique ID of the Subnet
	PkiSubnetID int32 `json:"pkiSubnetID"`
	// The unique ID of the User
	FkiUserID *int32 `json:"fkiUserID,omitempty"`
	// The unique ID of the Apikey
	FkiApikeyID *int32 `json:"fkiApikeyID,omitempty"`
	ObjSubnetDescription MultilingualSubnetDescription `json:"objSubnetDescription"`
	// The network of the Subnet in integer form. For example 8.8.8.0 would be 134744064
	ISubnetNetwork int64 `json:"iSubnetNetwork"`
	// The mask of the Subnet  in integer form. For example 255.255.255.0 would be 4294967040
	ISubnetMask int64 `json:"iSubnetMask"`
}

type _SubnetResponseCompound SubnetResponseCompound

// NewSubnetResponseCompound instantiates a new SubnetResponseCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubnetResponseCompound(pkiSubnetID int32, objSubnetDescription MultilingualSubnetDescription, iSubnetNetwork int64, iSubnetMask int64) *SubnetResponseCompound {
	this := SubnetResponseCompound{}
	this.PkiSubnetID = pkiSubnetID
	this.ObjSubnetDescription = objSubnetDescription
	this.ISubnetNetwork = iSubnetNetwork
	this.ISubnetMask = iSubnetMask
	return &this
}

// NewSubnetResponseCompoundWithDefaults instantiates a new SubnetResponseCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubnetResponseCompoundWithDefaults() *SubnetResponseCompound {
	this := SubnetResponseCompound{}
	return &this
}

// GetPkiSubnetID returns the PkiSubnetID field value
func (o *SubnetResponseCompound) GetPkiSubnetID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiSubnetID
}

// GetPkiSubnetIDOk returns a tuple with the PkiSubnetID field value
// and a boolean to check if the value has been set.
func (o *SubnetResponseCompound) GetPkiSubnetIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiSubnetID, true
}

// SetPkiSubnetID sets field value
func (o *SubnetResponseCompound) SetPkiSubnetID(v int32) {
	o.PkiSubnetID = v
}

// GetFkiUserID returns the FkiUserID field value if set, zero value otherwise.
func (o *SubnetResponseCompound) GetFkiUserID() int32 {
	if o == nil || IsNil(o.FkiUserID) {
		var ret int32
		return ret
	}
	return *o.FkiUserID
}

// GetFkiUserIDOk returns a tuple with the FkiUserID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubnetResponseCompound) GetFkiUserIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiUserID) {
		return nil, false
	}
	return o.FkiUserID, true
}

// HasFkiUserID returns a boolean if a field has been set.
func (o *SubnetResponseCompound) HasFkiUserID() bool {
	if o != nil && !IsNil(o.FkiUserID) {
		return true
	}

	return false
}

// SetFkiUserID gets a reference to the given int32 and assigns it to the FkiUserID field.
func (o *SubnetResponseCompound) SetFkiUserID(v int32) {
	o.FkiUserID = &v
}

// GetFkiApikeyID returns the FkiApikeyID field value if set, zero value otherwise.
func (o *SubnetResponseCompound) GetFkiApikeyID() int32 {
	if o == nil || IsNil(o.FkiApikeyID) {
		var ret int32
		return ret
	}
	return *o.FkiApikeyID
}

// GetFkiApikeyIDOk returns a tuple with the FkiApikeyID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubnetResponseCompound) GetFkiApikeyIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiApikeyID) {
		return nil, false
	}
	return o.FkiApikeyID, true
}

// HasFkiApikeyID returns a boolean if a field has been set.
func (o *SubnetResponseCompound) HasFkiApikeyID() bool {
	if o != nil && !IsNil(o.FkiApikeyID) {
		return true
	}

	return false
}

// SetFkiApikeyID gets a reference to the given int32 and assigns it to the FkiApikeyID field.
func (o *SubnetResponseCompound) SetFkiApikeyID(v int32) {
	o.FkiApikeyID = &v
}

// GetObjSubnetDescription returns the ObjSubnetDescription field value
func (o *SubnetResponseCompound) GetObjSubnetDescription() MultilingualSubnetDescription {
	if o == nil {
		var ret MultilingualSubnetDescription
		return ret
	}

	return o.ObjSubnetDescription
}

// GetObjSubnetDescriptionOk returns a tuple with the ObjSubnetDescription field value
// and a boolean to check if the value has been set.
func (o *SubnetResponseCompound) GetObjSubnetDescriptionOk() (*MultilingualSubnetDescription, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjSubnetDescription, true
}

// SetObjSubnetDescription sets field value
func (o *SubnetResponseCompound) SetObjSubnetDescription(v MultilingualSubnetDescription) {
	o.ObjSubnetDescription = v
}

// GetISubnetNetwork returns the ISubnetNetwork field value
func (o *SubnetResponseCompound) GetISubnetNetwork() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ISubnetNetwork
}

// GetISubnetNetworkOk returns a tuple with the ISubnetNetwork field value
// and a boolean to check if the value has been set.
func (o *SubnetResponseCompound) GetISubnetNetworkOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ISubnetNetwork, true
}

// SetISubnetNetwork sets field value
func (o *SubnetResponseCompound) SetISubnetNetwork(v int64) {
	o.ISubnetNetwork = v
}

// GetISubnetMask returns the ISubnetMask field value
func (o *SubnetResponseCompound) GetISubnetMask() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ISubnetMask
}

// GetISubnetMaskOk returns a tuple with the ISubnetMask field value
// and a boolean to check if the value has been set.
func (o *SubnetResponseCompound) GetISubnetMaskOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ISubnetMask, true
}

// SetISubnetMask sets field value
func (o *SubnetResponseCompound) SetISubnetMask(v int64) {
	o.ISubnetMask = v
}

func (o SubnetResponseCompound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubnetResponseCompound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiSubnetID"] = o.PkiSubnetID
	if !IsNil(o.FkiUserID) {
		toSerialize["fkiUserID"] = o.FkiUserID
	}
	if !IsNil(o.FkiApikeyID) {
		toSerialize["fkiApikeyID"] = o.FkiApikeyID
	}
	toSerialize["objSubnetDescription"] = o.ObjSubnetDescription
	toSerialize["iSubnetNetwork"] = o.ISubnetNetwork
	toSerialize["iSubnetMask"] = o.ISubnetMask
	return toSerialize, nil
}

func (o *SubnetResponseCompound) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiSubnetID",
		"objSubnetDescription",
		"iSubnetNetwork",
		"iSubnetMask",
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

	varSubnetResponseCompound := _SubnetResponseCompound{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSubnetResponseCompound)

	if err != nil {
		return err
	}

	*o = SubnetResponseCompound(varSubnetResponseCompound)

	return err
}

type NullableSubnetResponseCompound struct {
	value *SubnetResponseCompound
	isSet bool
}

func (v NullableSubnetResponseCompound) Get() *SubnetResponseCompound {
	return v.value
}

func (v *NullableSubnetResponseCompound) Set(val *SubnetResponseCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableSubnetResponseCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableSubnetResponseCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubnetResponseCompound(val *SubnetResponseCompound) *NullableSubnetResponseCompound {
	return &NullableSubnetResponseCompound{value: val, isSet: true}
}

func (v NullableSubnetResponseCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubnetResponseCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


