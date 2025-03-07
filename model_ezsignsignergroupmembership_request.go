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

// checks if the EzsignsignergroupmembershipRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignsignergroupmembershipRequest{}

// EzsignsignergroupmembershipRequest A Ezsignsignergroupmembership Object
type EzsignsignergroupmembershipRequest struct {
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

type _EzsignsignergroupmembershipRequest EzsignsignergroupmembershipRequest

// NewEzsignsignergroupmembershipRequest instantiates a new EzsignsignergroupmembershipRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignsignergroupmembershipRequest(fkiEzsignsignergroupID int32) *EzsignsignergroupmembershipRequest {
	this := EzsignsignergroupmembershipRequest{}
	this.FkiEzsignsignergroupID = fkiEzsignsignergroupID
	return &this
}

// NewEzsignsignergroupmembershipRequestWithDefaults instantiates a new EzsignsignergroupmembershipRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignsignergroupmembershipRequestWithDefaults() *EzsignsignergroupmembershipRequest {
	this := EzsignsignergroupmembershipRequest{}
	return &this
}

// GetPkiEzsignsignergroupmembershipID returns the PkiEzsignsignergroupmembershipID field value if set, zero value otherwise.
func (o *EzsignsignergroupmembershipRequest) GetPkiEzsignsignergroupmembershipID() int32 {
	if o == nil || IsNil(o.PkiEzsignsignergroupmembershipID) {
		var ret int32
		return ret
	}
	return *o.PkiEzsignsignergroupmembershipID
}

// GetPkiEzsignsignergroupmembershipIDOk returns a tuple with the PkiEzsignsignergroupmembershipID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignsignergroupmembershipRequest) GetPkiEzsignsignergroupmembershipIDOk() (*int32, bool) {
	if o == nil || IsNil(o.PkiEzsignsignergroupmembershipID) {
		return nil, false
	}
	return o.PkiEzsignsignergroupmembershipID, true
}

// HasPkiEzsignsignergroupmembershipID returns a boolean if a field has been set.
func (o *EzsignsignergroupmembershipRequest) HasPkiEzsignsignergroupmembershipID() bool {
	if o != nil && !IsNil(o.PkiEzsignsignergroupmembershipID) {
		return true
	}

	return false
}

// SetPkiEzsignsignergroupmembershipID gets a reference to the given int32 and assigns it to the PkiEzsignsignergroupmembershipID field.
func (o *EzsignsignergroupmembershipRequest) SetPkiEzsignsignergroupmembershipID(v int32) {
	o.PkiEzsignsignergroupmembershipID = &v
}

// GetFkiEzsignsignergroupID returns the FkiEzsignsignergroupID field value
func (o *EzsignsignergroupmembershipRequest) GetFkiEzsignsignergroupID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzsignsignergroupID
}

// GetFkiEzsignsignergroupIDOk returns a tuple with the FkiEzsignsignergroupID field value
// and a boolean to check if the value has been set.
func (o *EzsignsignergroupmembershipRequest) GetFkiEzsignsignergroupIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiEzsignsignergroupID, true
}

// SetFkiEzsignsignergroupID sets field value
func (o *EzsignsignergroupmembershipRequest) SetFkiEzsignsignergroupID(v int32) {
	o.FkiEzsignsignergroupID = v
}

// GetFkiEzsignsignerID returns the FkiEzsignsignerID field value if set, zero value otherwise.
func (o *EzsignsignergroupmembershipRequest) GetFkiEzsignsignerID() int32 {
	if o == nil || IsNil(o.FkiEzsignsignerID) {
		var ret int32
		return ret
	}
	return *o.FkiEzsignsignerID
}

// GetFkiEzsignsignerIDOk returns a tuple with the FkiEzsignsignerID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignsignergroupmembershipRequest) GetFkiEzsignsignerIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiEzsignsignerID) {
		return nil, false
	}
	return o.FkiEzsignsignerID, true
}

// HasFkiEzsignsignerID returns a boolean if a field has been set.
func (o *EzsignsignergroupmembershipRequest) HasFkiEzsignsignerID() bool {
	if o != nil && !IsNil(o.FkiEzsignsignerID) {
		return true
	}

	return false
}

// SetFkiEzsignsignerID gets a reference to the given int32 and assigns it to the FkiEzsignsignerID field.
func (o *EzsignsignergroupmembershipRequest) SetFkiEzsignsignerID(v int32) {
	o.FkiEzsignsignerID = &v
}

// GetFkiUserID returns the FkiUserID field value if set, zero value otherwise.
func (o *EzsignsignergroupmembershipRequest) GetFkiUserID() int32 {
	if o == nil || IsNil(o.FkiUserID) {
		var ret int32
		return ret
	}
	return *o.FkiUserID
}

// GetFkiUserIDOk returns a tuple with the FkiUserID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignsignergroupmembershipRequest) GetFkiUserIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiUserID) {
		return nil, false
	}
	return o.FkiUserID, true
}

// HasFkiUserID returns a boolean if a field has been set.
func (o *EzsignsignergroupmembershipRequest) HasFkiUserID() bool {
	if o != nil && !IsNil(o.FkiUserID) {
		return true
	}

	return false
}

// SetFkiUserID gets a reference to the given int32 and assigns it to the FkiUserID field.
func (o *EzsignsignergroupmembershipRequest) SetFkiUserID(v int32) {
	o.FkiUserID = &v
}

// GetFkiUsergroupID returns the FkiUsergroupID field value if set, zero value otherwise.
func (o *EzsignsignergroupmembershipRequest) GetFkiUsergroupID() int32 {
	if o == nil || IsNil(o.FkiUsergroupID) {
		var ret int32
		return ret
	}
	return *o.FkiUsergroupID
}

// GetFkiUsergroupIDOk returns a tuple with the FkiUsergroupID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignsignergroupmembershipRequest) GetFkiUsergroupIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiUsergroupID) {
		return nil, false
	}
	return o.FkiUsergroupID, true
}

// HasFkiUsergroupID returns a boolean if a field has been set.
func (o *EzsignsignergroupmembershipRequest) HasFkiUsergroupID() bool {
	if o != nil && !IsNil(o.FkiUsergroupID) {
		return true
	}

	return false
}

// SetFkiUsergroupID gets a reference to the given int32 and assigns it to the FkiUsergroupID field.
func (o *EzsignsignergroupmembershipRequest) SetFkiUsergroupID(v int32) {
	o.FkiUsergroupID = &v
}

func (o EzsignsignergroupmembershipRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignsignergroupmembershipRequest) ToMap() (map[string]interface{}, error) {
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

func (o *EzsignsignergroupmembershipRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varEzsignsignergroupmembershipRequest := _EzsignsignergroupmembershipRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsignsignergroupmembershipRequest)

	if err != nil {
		return err
	}

	*o = EzsignsignergroupmembershipRequest(varEzsignsignergroupmembershipRequest)

	return err
}

type NullableEzsignsignergroupmembershipRequest struct {
	value *EzsignsignergroupmembershipRequest
	isSet bool
}

func (v NullableEzsignsignergroupmembershipRequest) Get() *EzsignsignergroupmembershipRequest {
	return v.value
}

func (v *NullableEzsignsignergroupmembershipRequest) Set(val *EzsignsignergroupmembershipRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignsignergroupmembershipRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignsignergroupmembershipRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignsignergroupmembershipRequest(val *EzsignsignergroupmembershipRequest) *NullableEzsignsignergroupmembershipRequest {
	return &NullableEzsignsignergroupmembershipRequest{value: val, isSet: true}
}

func (v NullableEzsignsignergroupmembershipRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignsignergroupmembershipRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


