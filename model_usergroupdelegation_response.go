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

// checks if the UsergroupdelegationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UsergroupdelegationResponse{}

// UsergroupdelegationResponse A Usergroupdelegation Object
type UsergroupdelegationResponse struct {
	// The unique ID of the Usergroupdelegation
	PkiUsergroupdelegationID int32 `json:"pkiUsergroupdelegationID"`
	// The unique ID of the Usergroup
	FkiUsergroupID int32 `json:"fkiUsergroupID"`
	// The unique ID of the User
	FkiUserID int32 `json:"fkiUserID"`
	// The first name of the user
	SUserFirstname string `json:"sUserFirstname"`
	// The last name of the user
	SUserLastname string `json:"sUserLastname"`
	// The login name of the User.
	SUserLoginname string `json:"sUserLoginname"`
	// The email address.
	SEmailAddress *string `json:"sEmailAddress,omitempty"`
	// The Name of the Usergroup in the language of the requester
	SUsergroupNameX string `json:"sUsergroupNameX"`
}

type _UsergroupdelegationResponse UsergroupdelegationResponse

// NewUsergroupdelegationResponse instantiates a new UsergroupdelegationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsergroupdelegationResponse(pkiUsergroupdelegationID int32, fkiUsergroupID int32, fkiUserID int32, sUserFirstname string, sUserLastname string, sUserLoginname string, sUsergroupNameX string) *UsergroupdelegationResponse {
	this := UsergroupdelegationResponse{}
	this.PkiUsergroupdelegationID = pkiUsergroupdelegationID
	this.FkiUsergroupID = fkiUsergroupID
	this.FkiUserID = fkiUserID
	this.SUserFirstname = sUserFirstname
	this.SUserLastname = sUserLastname
	this.SUserLoginname = sUserLoginname
	this.SUsergroupNameX = sUsergroupNameX
	return &this
}

// NewUsergroupdelegationResponseWithDefaults instantiates a new UsergroupdelegationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsergroupdelegationResponseWithDefaults() *UsergroupdelegationResponse {
	this := UsergroupdelegationResponse{}
	return &this
}

// GetPkiUsergroupdelegationID returns the PkiUsergroupdelegationID field value
func (o *UsergroupdelegationResponse) GetPkiUsergroupdelegationID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiUsergroupdelegationID
}

// GetPkiUsergroupdelegationIDOk returns a tuple with the PkiUsergroupdelegationID field value
// and a boolean to check if the value has been set.
func (o *UsergroupdelegationResponse) GetPkiUsergroupdelegationIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiUsergroupdelegationID, true
}

// SetPkiUsergroupdelegationID sets field value
func (o *UsergroupdelegationResponse) SetPkiUsergroupdelegationID(v int32) {
	o.PkiUsergroupdelegationID = v
}

// GetFkiUsergroupID returns the FkiUsergroupID field value
func (o *UsergroupdelegationResponse) GetFkiUsergroupID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiUsergroupID
}

// GetFkiUsergroupIDOk returns a tuple with the FkiUsergroupID field value
// and a boolean to check if the value has been set.
func (o *UsergroupdelegationResponse) GetFkiUsergroupIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiUsergroupID, true
}

// SetFkiUsergroupID sets field value
func (o *UsergroupdelegationResponse) SetFkiUsergroupID(v int32) {
	o.FkiUsergroupID = v
}

// GetFkiUserID returns the FkiUserID field value
func (o *UsergroupdelegationResponse) GetFkiUserID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiUserID
}

// GetFkiUserIDOk returns a tuple with the FkiUserID field value
// and a boolean to check if the value has been set.
func (o *UsergroupdelegationResponse) GetFkiUserIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiUserID, true
}

// SetFkiUserID sets field value
func (o *UsergroupdelegationResponse) SetFkiUserID(v int32) {
	o.FkiUserID = v
}

// GetSUserFirstname returns the SUserFirstname field value
func (o *UsergroupdelegationResponse) GetSUserFirstname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SUserFirstname
}

// GetSUserFirstnameOk returns a tuple with the SUserFirstname field value
// and a boolean to check if the value has been set.
func (o *UsergroupdelegationResponse) GetSUserFirstnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SUserFirstname, true
}

// SetSUserFirstname sets field value
func (o *UsergroupdelegationResponse) SetSUserFirstname(v string) {
	o.SUserFirstname = v
}

// GetSUserLastname returns the SUserLastname field value
func (o *UsergroupdelegationResponse) GetSUserLastname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SUserLastname
}

// GetSUserLastnameOk returns a tuple with the SUserLastname field value
// and a boolean to check if the value has been set.
func (o *UsergroupdelegationResponse) GetSUserLastnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SUserLastname, true
}

// SetSUserLastname sets field value
func (o *UsergroupdelegationResponse) SetSUserLastname(v string) {
	o.SUserLastname = v
}

// GetSUserLoginname returns the SUserLoginname field value
func (o *UsergroupdelegationResponse) GetSUserLoginname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SUserLoginname
}

// GetSUserLoginnameOk returns a tuple with the SUserLoginname field value
// and a boolean to check if the value has been set.
func (o *UsergroupdelegationResponse) GetSUserLoginnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SUserLoginname, true
}

// SetSUserLoginname sets field value
func (o *UsergroupdelegationResponse) SetSUserLoginname(v string) {
	o.SUserLoginname = v
}

// GetSEmailAddress returns the SEmailAddress field value if set, zero value otherwise.
func (o *UsergroupdelegationResponse) GetSEmailAddress() string {
	if o == nil || IsNil(o.SEmailAddress) {
		var ret string
		return ret
	}
	return *o.SEmailAddress
}

// GetSEmailAddressOk returns a tuple with the SEmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsergroupdelegationResponse) GetSEmailAddressOk() (*string, bool) {
	if o == nil || IsNil(o.SEmailAddress) {
		return nil, false
	}
	return o.SEmailAddress, true
}

// HasSEmailAddress returns a boolean if a field has been set.
func (o *UsergroupdelegationResponse) HasSEmailAddress() bool {
	if o != nil && !IsNil(o.SEmailAddress) {
		return true
	}

	return false
}

// SetSEmailAddress gets a reference to the given string and assigns it to the SEmailAddress field.
func (o *UsergroupdelegationResponse) SetSEmailAddress(v string) {
	o.SEmailAddress = &v
}

// GetSUsergroupNameX returns the SUsergroupNameX field value
func (o *UsergroupdelegationResponse) GetSUsergroupNameX() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SUsergroupNameX
}

// GetSUsergroupNameXOk returns a tuple with the SUsergroupNameX field value
// and a boolean to check if the value has been set.
func (o *UsergroupdelegationResponse) GetSUsergroupNameXOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SUsergroupNameX, true
}

// SetSUsergroupNameX sets field value
func (o *UsergroupdelegationResponse) SetSUsergroupNameX(v string) {
	o.SUsergroupNameX = v
}

func (o UsergroupdelegationResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsergroupdelegationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiUsergroupdelegationID"] = o.PkiUsergroupdelegationID
	toSerialize["fkiUsergroupID"] = o.FkiUsergroupID
	toSerialize["fkiUserID"] = o.FkiUserID
	toSerialize["sUserFirstname"] = o.SUserFirstname
	toSerialize["sUserLastname"] = o.SUserLastname
	toSerialize["sUserLoginname"] = o.SUserLoginname
	if !IsNil(o.SEmailAddress) {
		toSerialize["sEmailAddress"] = o.SEmailAddress
	}
	toSerialize["sUsergroupNameX"] = o.SUsergroupNameX
	return toSerialize, nil
}

func (o *UsergroupdelegationResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiUsergroupdelegationID",
		"fkiUsergroupID",
		"fkiUserID",
		"sUserFirstname",
		"sUserLastname",
		"sUserLoginname",
		"sUsergroupNameX",
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

	varUsergroupdelegationResponse := _UsergroupdelegationResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUsergroupdelegationResponse)

	if err != nil {
		return err
	}

	*o = UsergroupdelegationResponse(varUsergroupdelegationResponse)

	return err
}

type NullableUsergroupdelegationResponse struct {
	value *UsergroupdelegationResponse
	isSet bool
}

func (v NullableUsergroupdelegationResponse) Get() *UsergroupdelegationResponse {
	return v.value
}

func (v *NullableUsergroupdelegationResponse) Set(val *UsergroupdelegationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUsergroupdelegationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUsergroupdelegationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsergroupdelegationResponse(val *UsergroupdelegationResponse) *NullableUsergroupdelegationResponse {
	return &NullableUsergroupdelegationResponse{value: val, isSet: true}
}

func (v NullableUsergroupdelegationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsergroupdelegationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


