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

// checks if the UserstagedResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserstagedResponse{}

// UserstagedResponse A Userstaged Object
type UserstagedResponse struct {
	// The unique ID of the Userstaged
	PkiUserstagedID int32 `json:"pkiUserstagedID"`
	// The unique ID of the Email
	FkiEmailID int32 `json:"fkiEmailID"`
	// The email address.
	SEmailAddress string "json:\"sEmailAddress\" validate:\"regexp=^[\\\\w.%+\\\\-!#$%&'*+\\/=?^`{|}~]+@[a-zA-Z0-9.-]+\\\\.[a-zA-Z]{2,20}$\""
	// The firstname of the Userstaged
	SUserstagedFirstname string `json:"sUserstagedFirstname" validate:"regexp=^.{0,20}$"`
	// The lastname of the Userstaged
	SUserstagedLastname string `json:"sUserstagedLastname" validate:"regexp=^.{0,25}$"`
	// The externalid of the Userstaged
	SUserstagedExternalid string `json:"sUserstagedExternalid" validate:"regexp=^.{1,60}$"`
}

type _UserstagedResponse UserstagedResponse

// NewUserstagedResponse instantiates a new UserstagedResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserstagedResponse(pkiUserstagedID int32, fkiEmailID int32, sEmailAddress string, sUserstagedFirstname string, sUserstagedLastname string, sUserstagedExternalid string) *UserstagedResponse {
	this := UserstagedResponse{}
	this.PkiUserstagedID = pkiUserstagedID
	this.FkiEmailID = fkiEmailID
	this.SEmailAddress = sEmailAddress
	this.SUserstagedFirstname = sUserstagedFirstname
	this.SUserstagedLastname = sUserstagedLastname
	this.SUserstagedExternalid = sUserstagedExternalid
	return &this
}

// NewUserstagedResponseWithDefaults instantiates a new UserstagedResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserstagedResponseWithDefaults() *UserstagedResponse {
	this := UserstagedResponse{}
	return &this
}

// GetPkiUserstagedID returns the PkiUserstagedID field value
func (o *UserstagedResponse) GetPkiUserstagedID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiUserstagedID
}

// GetPkiUserstagedIDOk returns a tuple with the PkiUserstagedID field value
// and a boolean to check if the value has been set.
func (o *UserstagedResponse) GetPkiUserstagedIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiUserstagedID, true
}

// SetPkiUserstagedID sets field value
func (o *UserstagedResponse) SetPkiUserstagedID(v int32) {
	o.PkiUserstagedID = v
}

// GetFkiEmailID returns the FkiEmailID field value
func (o *UserstagedResponse) GetFkiEmailID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEmailID
}

// GetFkiEmailIDOk returns a tuple with the FkiEmailID field value
// and a boolean to check if the value has been set.
func (o *UserstagedResponse) GetFkiEmailIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiEmailID, true
}

// SetFkiEmailID sets field value
func (o *UserstagedResponse) SetFkiEmailID(v int32) {
	o.FkiEmailID = v
}

// GetSEmailAddress returns the SEmailAddress field value
func (o *UserstagedResponse) GetSEmailAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEmailAddress
}

// GetSEmailAddressOk returns a tuple with the SEmailAddress field value
// and a boolean to check if the value has been set.
func (o *UserstagedResponse) GetSEmailAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SEmailAddress, true
}

// SetSEmailAddress sets field value
func (o *UserstagedResponse) SetSEmailAddress(v string) {
	o.SEmailAddress = v
}

// GetSUserstagedFirstname returns the SUserstagedFirstname field value
func (o *UserstagedResponse) GetSUserstagedFirstname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SUserstagedFirstname
}

// GetSUserstagedFirstnameOk returns a tuple with the SUserstagedFirstname field value
// and a boolean to check if the value has been set.
func (o *UserstagedResponse) GetSUserstagedFirstnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SUserstagedFirstname, true
}

// SetSUserstagedFirstname sets field value
func (o *UserstagedResponse) SetSUserstagedFirstname(v string) {
	o.SUserstagedFirstname = v
}

// GetSUserstagedLastname returns the SUserstagedLastname field value
func (o *UserstagedResponse) GetSUserstagedLastname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SUserstagedLastname
}

// GetSUserstagedLastnameOk returns a tuple with the SUserstagedLastname field value
// and a boolean to check if the value has been set.
func (o *UserstagedResponse) GetSUserstagedLastnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SUserstagedLastname, true
}

// SetSUserstagedLastname sets field value
func (o *UserstagedResponse) SetSUserstagedLastname(v string) {
	o.SUserstagedLastname = v
}

// GetSUserstagedExternalid returns the SUserstagedExternalid field value
func (o *UserstagedResponse) GetSUserstagedExternalid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SUserstagedExternalid
}

// GetSUserstagedExternalidOk returns a tuple with the SUserstagedExternalid field value
// and a boolean to check if the value has been set.
func (o *UserstagedResponse) GetSUserstagedExternalidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SUserstagedExternalid, true
}

// SetSUserstagedExternalid sets field value
func (o *UserstagedResponse) SetSUserstagedExternalid(v string) {
	o.SUserstagedExternalid = v
}

func (o UserstagedResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserstagedResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiUserstagedID"] = o.PkiUserstagedID
	toSerialize["fkiEmailID"] = o.FkiEmailID
	toSerialize["sEmailAddress"] = o.SEmailAddress
	toSerialize["sUserstagedFirstname"] = o.SUserstagedFirstname
	toSerialize["sUserstagedLastname"] = o.SUserstagedLastname
	toSerialize["sUserstagedExternalid"] = o.SUserstagedExternalid
	return toSerialize, nil
}

func (o *UserstagedResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiUserstagedID",
		"fkiEmailID",
		"sEmailAddress",
		"sUserstagedFirstname",
		"sUserstagedLastname",
		"sUserstagedExternalid",
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

	varUserstagedResponse := _UserstagedResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserstagedResponse)

	if err != nil {
		return err
	}

	*o = UserstagedResponse(varUserstagedResponse)

	return err
}

type NullableUserstagedResponse struct {
	value *UserstagedResponse
	isSet bool
}

func (v NullableUserstagedResponse) Get() *UserstagedResponse {
	return v.value
}

func (v *NullableUserstagedResponse) Set(val *UserstagedResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserstagedResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserstagedResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserstagedResponse(val *UserstagedResponse) *NullableUserstagedResponse {
	return &NullableUserstagedResponse{value: val, isSet: true}
}

func (v NullableUserstagedResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserstagedResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


