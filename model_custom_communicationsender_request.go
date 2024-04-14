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
)

// checks if the CustomCommunicationsenderRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomCommunicationsenderRequest{}

// CustomCommunicationsenderRequest A Communicationsender Object
type CustomCommunicationsenderRequest struct {
	// The unique ID of the Agent.
	FkiAgentID *int32 `json:"fkiAgentID,omitempty"`
	// The unique ID of the Broker.
	FkiBrokerID *int32 `json:"fkiBrokerID,omitempty"`
	// The unique ID of the Mailboxshared
	FkiMailboxsharedID *int32 `json:"fkiMailboxsharedID,omitempty"`
	// The unique ID of the Phonelineshared
	FkiPhonelinesharedID *int32 `json:"fkiPhonelinesharedID,omitempty"`
	// The unique ID of the User
	FkiUserID *int32 `json:"fkiUserID,omitempty"`
}

// NewCustomCommunicationsenderRequest instantiates a new CustomCommunicationsenderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomCommunicationsenderRequest() *CustomCommunicationsenderRequest {
	this := CustomCommunicationsenderRequest{}
	return &this
}

// NewCustomCommunicationsenderRequestWithDefaults instantiates a new CustomCommunicationsenderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomCommunicationsenderRequestWithDefaults() *CustomCommunicationsenderRequest {
	this := CustomCommunicationsenderRequest{}
	return &this
}

// GetFkiAgentID returns the FkiAgentID field value if set, zero value otherwise.
func (o *CustomCommunicationsenderRequest) GetFkiAgentID() int32 {
	if o == nil || IsNil(o.FkiAgentID) {
		var ret int32
		return ret
	}
	return *o.FkiAgentID
}

// GetFkiAgentIDOk returns a tuple with the FkiAgentID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomCommunicationsenderRequest) GetFkiAgentIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiAgentID) {
		return nil, false
	}
	return o.FkiAgentID, true
}

// HasFkiAgentID returns a boolean if a field has been set.
func (o *CustomCommunicationsenderRequest) HasFkiAgentID() bool {
	if o != nil && !IsNil(o.FkiAgentID) {
		return true
	}

	return false
}

// SetFkiAgentID gets a reference to the given int32 and assigns it to the FkiAgentID field.
func (o *CustomCommunicationsenderRequest) SetFkiAgentID(v int32) {
	o.FkiAgentID = &v
}

// GetFkiBrokerID returns the FkiBrokerID field value if set, zero value otherwise.
func (o *CustomCommunicationsenderRequest) GetFkiBrokerID() int32 {
	if o == nil || IsNil(o.FkiBrokerID) {
		var ret int32
		return ret
	}
	return *o.FkiBrokerID
}

// GetFkiBrokerIDOk returns a tuple with the FkiBrokerID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomCommunicationsenderRequest) GetFkiBrokerIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiBrokerID) {
		return nil, false
	}
	return o.FkiBrokerID, true
}

// HasFkiBrokerID returns a boolean if a field has been set.
func (o *CustomCommunicationsenderRequest) HasFkiBrokerID() bool {
	if o != nil && !IsNil(o.FkiBrokerID) {
		return true
	}

	return false
}

// SetFkiBrokerID gets a reference to the given int32 and assigns it to the FkiBrokerID field.
func (o *CustomCommunicationsenderRequest) SetFkiBrokerID(v int32) {
	o.FkiBrokerID = &v
}

// GetFkiMailboxsharedID returns the FkiMailboxsharedID field value if set, zero value otherwise.
func (o *CustomCommunicationsenderRequest) GetFkiMailboxsharedID() int32 {
	if o == nil || IsNil(o.FkiMailboxsharedID) {
		var ret int32
		return ret
	}
	return *o.FkiMailboxsharedID
}

// GetFkiMailboxsharedIDOk returns a tuple with the FkiMailboxsharedID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomCommunicationsenderRequest) GetFkiMailboxsharedIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiMailboxsharedID) {
		return nil, false
	}
	return o.FkiMailboxsharedID, true
}

// HasFkiMailboxsharedID returns a boolean if a field has been set.
func (o *CustomCommunicationsenderRequest) HasFkiMailboxsharedID() bool {
	if o != nil && !IsNil(o.FkiMailboxsharedID) {
		return true
	}

	return false
}

// SetFkiMailboxsharedID gets a reference to the given int32 and assigns it to the FkiMailboxsharedID field.
func (o *CustomCommunicationsenderRequest) SetFkiMailboxsharedID(v int32) {
	o.FkiMailboxsharedID = &v
}

// GetFkiPhonelinesharedID returns the FkiPhonelinesharedID field value if set, zero value otherwise.
func (o *CustomCommunicationsenderRequest) GetFkiPhonelinesharedID() int32 {
	if o == nil || IsNil(o.FkiPhonelinesharedID) {
		var ret int32
		return ret
	}
	return *o.FkiPhonelinesharedID
}

// GetFkiPhonelinesharedIDOk returns a tuple with the FkiPhonelinesharedID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomCommunicationsenderRequest) GetFkiPhonelinesharedIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiPhonelinesharedID) {
		return nil, false
	}
	return o.FkiPhonelinesharedID, true
}

// HasFkiPhonelinesharedID returns a boolean if a field has been set.
func (o *CustomCommunicationsenderRequest) HasFkiPhonelinesharedID() bool {
	if o != nil && !IsNil(o.FkiPhonelinesharedID) {
		return true
	}

	return false
}

// SetFkiPhonelinesharedID gets a reference to the given int32 and assigns it to the FkiPhonelinesharedID field.
func (o *CustomCommunicationsenderRequest) SetFkiPhonelinesharedID(v int32) {
	o.FkiPhonelinesharedID = &v
}

// GetFkiUserID returns the FkiUserID field value if set, zero value otherwise.
func (o *CustomCommunicationsenderRequest) GetFkiUserID() int32 {
	if o == nil || IsNil(o.FkiUserID) {
		var ret int32
		return ret
	}
	return *o.FkiUserID
}

// GetFkiUserIDOk returns a tuple with the FkiUserID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomCommunicationsenderRequest) GetFkiUserIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiUserID) {
		return nil, false
	}
	return o.FkiUserID, true
}

// HasFkiUserID returns a boolean if a field has been set.
func (o *CustomCommunicationsenderRequest) HasFkiUserID() bool {
	if o != nil && !IsNil(o.FkiUserID) {
		return true
	}

	return false
}

// SetFkiUserID gets a reference to the given int32 and assigns it to the FkiUserID field.
func (o *CustomCommunicationsenderRequest) SetFkiUserID(v int32) {
	o.FkiUserID = &v
}

func (o CustomCommunicationsenderRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomCommunicationsenderRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FkiAgentID) {
		toSerialize["fkiAgentID"] = o.FkiAgentID
	}
	if !IsNil(o.FkiBrokerID) {
		toSerialize["fkiBrokerID"] = o.FkiBrokerID
	}
	if !IsNil(o.FkiMailboxsharedID) {
		toSerialize["fkiMailboxsharedID"] = o.FkiMailboxsharedID
	}
	if !IsNil(o.FkiPhonelinesharedID) {
		toSerialize["fkiPhonelinesharedID"] = o.FkiPhonelinesharedID
	}
	if !IsNil(o.FkiUserID) {
		toSerialize["fkiUserID"] = o.FkiUserID
	}
	return toSerialize, nil
}

type NullableCustomCommunicationsenderRequest struct {
	value *CustomCommunicationsenderRequest
	isSet bool
}

func (v NullableCustomCommunicationsenderRequest) Get() *CustomCommunicationsenderRequest {
	return v.value
}

func (v *NullableCustomCommunicationsenderRequest) Set(val *CustomCommunicationsenderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomCommunicationsenderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomCommunicationsenderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomCommunicationsenderRequest(val *CustomCommunicationsenderRequest) *NullableCustomCommunicationsenderRequest {
	return &NullableCustomCommunicationsenderRequest{value: val, isSet: true}
}

func (v NullableCustomCommunicationsenderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomCommunicationsenderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


