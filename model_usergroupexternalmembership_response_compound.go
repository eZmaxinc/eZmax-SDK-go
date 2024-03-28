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

// checks if the UsergroupexternalmembershipResponseCompound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UsergroupexternalmembershipResponseCompound{}

// UsergroupexternalmembershipResponseCompound A Usergroupexternalmembership Object
type UsergroupexternalmembershipResponseCompound struct {
	// The unique ID of the Usergroupexternalmembership
	PkiUsergroupexternalmembershipID int32 `json:"pkiUsergroupexternalmembershipID"`
	// The unique ID of the Usergroupexternal
	FkiUsergroupexternalID int32 `json:"fkiUsergroupexternalID"`
	// The unique ID of the User
	FkiUserID int32 `json:"fkiUserID"`
	// The first name of the user
	SUserFirstname string `json:"sUserFirstname"`
	// The last name of the user
	SUserLastname string `json:"sUserLastname"`
	// The login name of the User.
	SUserLoginname string `json:"sUserLoginname"`
	// The email address.
	SEmailAddress string `json:"sEmailAddress"`
	// The name of the Usergroupexternal
	SUsergroupexternalName string `json:"sUsergroupexternalName"`
}

type _UsergroupexternalmembershipResponseCompound UsergroupexternalmembershipResponseCompound

// NewUsergroupexternalmembershipResponseCompound instantiates a new UsergroupexternalmembershipResponseCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsergroupexternalmembershipResponseCompound(pkiUsergroupexternalmembershipID int32, fkiUsergroupexternalID int32, fkiUserID int32, sUserFirstname string, sUserLastname string, sUserLoginname string, sEmailAddress string, sUsergroupexternalName string) *UsergroupexternalmembershipResponseCompound {
	this := UsergroupexternalmembershipResponseCompound{}
	this.PkiUsergroupexternalmembershipID = pkiUsergroupexternalmembershipID
	this.FkiUsergroupexternalID = fkiUsergroupexternalID
	this.FkiUserID = fkiUserID
	this.SUserFirstname = sUserFirstname
	this.SUserLastname = sUserLastname
	this.SUserLoginname = sUserLoginname
	this.SEmailAddress = sEmailAddress
	this.SUsergroupexternalName = sUsergroupexternalName
	return &this
}

// NewUsergroupexternalmembershipResponseCompoundWithDefaults instantiates a new UsergroupexternalmembershipResponseCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsergroupexternalmembershipResponseCompoundWithDefaults() *UsergroupexternalmembershipResponseCompound {
	this := UsergroupexternalmembershipResponseCompound{}
	return &this
}

// GetPkiUsergroupexternalmembershipID returns the PkiUsergroupexternalmembershipID field value
func (o *UsergroupexternalmembershipResponseCompound) GetPkiUsergroupexternalmembershipID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiUsergroupexternalmembershipID
}

// GetPkiUsergroupexternalmembershipIDOk returns a tuple with the PkiUsergroupexternalmembershipID field value
// and a boolean to check if the value has been set.
func (o *UsergroupexternalmembershipResponseCompound) GetPkiUsergroupexternalmembershipIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiUsergroupexternalmembershipID, true
}

// SetPkiUsergroupexternalmembershipID sets field value
func (o *UsergroupexternalmembershipResponseCompound) SetPkiUsergroupexternalmembershipID(v int32) {
	o.PkiUsergroupexternalmembershipID = v
}

// GetFkiUsergroupexternalID returns the FkiUsergroupexternalID field value
func (o *UsergroupexternalmembershipResponseCompound) GetFkiUsergroupexternalID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiUsergroupexternalID
}

// GetFkiUsergroupexternalIDOk returns a tuple with the FkiUsergroupexternalID field value
// and a boolean to check if the value has been set.
func (o *UsergroupexternalmembershipResponseCompound) GetFkiUsergroupexternalIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiUsergroupexternalID, true
}

// SetFkiUsergroupexternalID sets field value
func (o *UsergroupexternalmembershipResponseCompound) SetFkiUsergroupexternalID(v int32) {
	o.FkiUsergroupexternalID = v
}

// GetFkiUserID returns the FkiUserID field value
func (o *UsergroupexternalmembershipResponseCompound) GetFkiUserID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiUserID
}

// GetFkiUserIDOk returns a tuple with the FkiUserID field value
// and a boolean to check if the value has been set.
func (o *UsergroupexternalmembershipResponseCompound) GetFkiUserIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiUserID, true
}

// SetFkiUserID sets field value
func (o *UsergroupexternalmembershipResponseCompound) SetFkiUserID(v int32) {
	o.FkiUserID = v
}

// GetSUserFirstname returns the SUserFirstname field value
func (o *UsergroupexternalmembershipResponseCompound) GetSUserFirstname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SUserFirstname
}

// GetSUserFirstnameOk returns a tuple with the SUserFirstname field value
// and a boolean to check if the value has been set.
func (o *UsergroupexternalmembershipResponseCompound) GetSUserFirstnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SUserFirstname, true
}

// SetSUserFirstname sets field value
func (o *UsergroupexternalmembershipResponseCompound) SetSUserFirstname(v string) {
	o.SUserFirstname = v
}

// GetSUserLastname returns the SUserLastname field value
func (o *UsergroupexternalmembershipResponseCompound) GetSUserLastname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SUserLastname
}

// GetSUserLastnameOk returns a tuple with the SUserLastname field value
// and a boolean to check if the value has been set.
func (o *UsergroupexternalmembershipResponseCompound) GetSUserLastnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SUserLastname, true
}

// SetSUserLastname sets field value
func (o *UsergroupexternalmembershipResponseCompound) SetSUserLastname(v string) {
	o.SUserLastname = v
}

// GetSUserLoginname returns the SUserLoginname field value
func (o *UsergroupexternalmembershipResponseCompound) GetSUserLoginname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SUserLoginname
}

// GetSUserLoginnameOk returns a tuple with the SUserLoginname field value
// and a boolean to check if the value has been set.
func (o *UsergroupexternalmembershipResponseCompound) GetSUserLoginnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SUserLoginname, true
}

// SetSUserLoginname sets field value
func (o *UsergroupexternalmembershipResponseCompound) SetSUserLoginname(v string) {
	o.SUserLoginname = v
}

// GetSEmailAddress returns the SEmailAddress field value
func (o *UsergroupexternalmembershipResponseCompound) GetSEmailAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEmailAddress
}

// GetSEmailAddressOk returns a tuple with the SEmailAddress field value
// and a boolean to check if the value has been set.
func (o *UsergroupexternalmembershipResponseCompound) GetSEmailAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SEmailAddress, true
}

// SetSEmailAddress sets field value
func (o *UsergroupexternalmembershipResponseCompound) SetSEmailAddress(v string) {
	o.SEmailAddress = v
}

// GetSUsergroupexternalName returns the SUsergroupexternalName field value
func (o *UsergroupexternalmembershipResponseCompound) GetSUsergroupexternalName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SUsergroupexternalName
}

// GetSUsergroupexternalNameOk returns a tuple with the SUsergroupexternalName field value
// and a boolean to check if the value has been set.
func (o *UsergroupexternalmembershipResponseCompound) GetSUsergroupexternalNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SUsergroupexternalName, true
}

// SetSUsergroupexternalName sets field value
func (o *UsergroupexternalmembershipResponseCompound) SetSUsergroupexternalName(v string) {
	o.SUsergroupexternalName = v
}

func (o UsergroupexternalmembershipResponseCompound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsergroupexternalmembershipResponseCompound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiUsergroupexternalmembershipID"] = o.PkiUsergroupexternalmembershipID
	toSerialize["fkiUsergroupexternalID"] = o.FkiUsergroupexternalID
	toSerialize["fkiUserID"] = o.FkiUserID
	toSerialize["sUserFirstname"] = o.SUserFirstname
	toSerialize["sUserLastname"] = o.SUserLastname
	toSerialize["sUserLoginname"] = o.SUserLoginname
	toSerialize["sEmailAddress"] = o.SEmailAddress
	toSerialize["sUsergroupexternalName"] = o.SUsergroupexternalName
	return toSerialize, nil
}

func (o *UsergroupexternalmembershipResponseCompound) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiUsergroupexternalmembershipID",
		"fkiUsergroupexternalID",
		"fkiUserID",
		"sUserFirstname",
		"sUserLastname",
		"sUserLoginname",
		"sEmailAddress",
		"sUsergroupexternalName",
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

	varUsergroupexternalmembershipResponseCompound := _UsergroupexternalmembershipResponseCompound{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUsergroupexternalmembershipResponseCompound)

	if err != nil {
		return err
	}

	*o = UsergroupexternalmembershipResponseCompound(varUsergroupexternalmembershipResponseCompound)

	return err
}

type NullableUsergroupexternalmembershipResponseCompound struct {
	value *UsergroupexternalmembershipResponseCompound
	isSet bool
}

func (v NullableUsergroupexternalmembershipResponseCompound) Get() *UsergroupexternalmembershipResponseCompound {
	return v.value
}

func (v *NullableUsergroupexternalmembershipResponseCompound) Set(val *UsergroupexternalmembershipResponseCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableUsergroupexternalmembershipResponseCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableUsergroupexternalmembershipResponseCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsergroupexternalmembershipResponseCompound(val *UsergroupexternalmembershipResponseCompound) *NullableUsergroupexternalmembershipResponseCompound {
	return &NullableUsergroupexternalmembershipResponseCompound{value: val, isSet: true}
}

func (v NullableUsergroupexternalmembershipResponseCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsergroupexternalmembershipResponseCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

