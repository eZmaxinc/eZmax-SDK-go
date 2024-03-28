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

// checks if the UsergroupmembershipResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UsergroupmembershipResponse{}

// UsergroupmembershipResponse A Usergroupmembership Object
type UsergroupmembershipResponse struct {
	// The unique ID of the Usergroupmembership
	PkiUsergroupmembershipID int32 `json:"pkiUsergroupmembershipID"`
	// The unique ID of the Usergroup
	FkiUsergroupID int32 `json:"fkiUsergroupID"`
	// The unique ID of the User
	FkiUserID *int32 `json:"fkiUserID,omitempty"`
	// The unique ID of the Usergroupexternal
	FkiUsergroupexternalID *int32 `json:"fkiUsergroupexternalID,omitempty"`
	// The first name of the user
	SUserFirstname *string `json:"sUserFirstname,omitempty"`
	// The last name of the user
	SUserLastname *string `json:"sUserLastname,omitempty"`
	// The login name of the User.
	SUserLoginname *string `json:"sUserLoginname,omitempty"`
	// The email address.
	SEmailAddress *string `json:"sEmailAddress,omitempty"`
	// The Name of the Usergroup in the language of the requester
	SUsergroupNameX string `json:"sUsergroupNameX"`
	// The name of the Usergroupexternal
	SUsergroupexternalName *string `json:"sUsergroupexternalName,omitempty"`
}

type _UsergroupmembershipResponse UsergroupmembershipResponse

// NewUsergroupmembershipResponse instantiates a new UsergroupmembershipResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsergroupmembershipResponse(pkiUsergroupmembershipID int32, fkiUsergroupID int32, sUsergroupNameX string) *UsergroupmembershipResponse {
	this := UsergroupmembershipResponse{}
	this.PkiUsergroupmembershipID = pkiUsergroupmembershipID
	this.FkiUsergroupID = fkiUsergroupID
	this.SUsergroupNameX = sUsergroupNameX
	return &this
}

// NewUsergroupmembershipResponseWithDefaults instantiates a new UsergroupmembershipResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsergroupmembershipResponseWithDefaults() *UsergroupmembershipResponse {
	this := UsergroupmembershipResponse{}
	return &this
}

// GetPkiUsergroupmembershipID returns the PkiUsergroupmembershipID field value
func (o *UsergroupmembershipResponse) GetPkiUsergroupmembershipID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiUsergroupmembershipID
}

// GetPkiUsergroupmembershipIDOk returns a tuple with the PkiUsergroupmembershipID field value
// and a boolean to check if the value has been set.
func (o *UsergroupmembershipResponse) GetPkiUsergroupmembershipIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiUsergroupmembershipID, true
}

// SetPkiUsergroupmembershipID sets field value
func (o *UsergroupmembershipResponse) SetPkiUsergroupmembershipID(v int32) {
	o.PkiUsergroupmembershipID = v
}

// GetFkiUsergroupID returns the FkiUsergroupID field value
func (o *UsergroupmembershipResponse) GetFkiUsergroupID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiUsergroupID
}

// GetFkiUsergroupIDOk returns a tuple with the FkiUsergroupID field value
// and a boolean to check if the value has been set.
func (o *UsergroupmembershipResponse) GetFkiUsergroupIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiUsergroupID, true
}

// SetFkiUsergroupID sets field value
func (o *UsergroupmembershipResponse) SetFkiUsergroupID(v int32) {
	o.FkiUsergroupID = v
}

// GetFkiUserID returns the FkiUserID field value if set, zero value otherwise.
func (o *UsergroupmembershipResponse) GetFkiUserID() int32 {
	if o == nil || IsNil(o.FkiUserID) {
		var ret int32
		return ret
	}
	return *o.FkiUserID
}

// GetFkiUserIDOk returns a tuple with the FkiUserID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsergroupmembershipResponse) GetFkiUserIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiUserID) {
		return nil, false
	}
	return o.FkiUserID, true
}

// HasFkiUserID returns a boolean if a field has been set.
func (o *UsergroupmembershipResponse) HasFkiUserID() bool {
	if o != nil && !IsNil(o.FkiUserID) {
		return true
	}

	return false
}

// SetFkiUserID gets a reference to the given int32 and assigns it to the FkiUserID field.
func (o *UsergroupmembershipResponse) SetFkiUserID(v int32) {
	o.FkiUserID = &v
}

// GetFkiUsergroupexternalID returns the FkiUsergroupexternalID field value if set, zero value otherwise.
func (o *UsergroupmembershipResponse) GetFkiUsergroupexternalID() int32 {
	if o == nil || IsNil(o.FkiUsergroupexternalID) {
		var ret int32
		return ret
	}
	return *o.FkiUsergroupexternalID
}

// GetFkiUsergroupexternalIDOk returns a tuple with the FkiUsergroupexternalID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsergroupmembershipResponse) GetFkiUsergroupexternalIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiUsergroupexternalID) {
		return nil, false
	}
	return o.FkiUsergroupexternalID, true
}

// HasFkiUsergroupexternalID returns a boolean if a field has been set.
func (o *UsergroupmembershipResponse) HasFkiUsergroupexternalID() bool {
	if o != nil && !IsNil(o.FkiUsergroupexternalID) {
		return true
	}

	return false
}

// SetFkiUsergroupexternalID gets a reference to the given int32 and assigns it to the FkiUsergroupexternalID field.
func (o *UsergroupmembershipResponse) SetFkiUsergroupexternalID(v int32) {
	o.FkiUsergroupexternalID = &v
}

// GetSUserFirstname returns the SUserFirstname field value if set, zero value otherwise.
func (o *UsergroupmembershipResponse) GetSUserFirstname() string {
	if o == nil || IsNil(o.SUserFirstname) {
		var ret string
		return ret
	}
	return *o.SUserFirstname
}

// GetSUserFirstnameOk returns a tuple with the SUserFirstname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsergroupmembershipResponse) GetSUserFirstnameOk() (*string, bool) {
	if o == nil || IsNil(o.SUserFirstname) {
		return nil, false
	}
	return o.SUserFirstname, true
}

// HasSUserFirstname returns a boolean if a field has been set.
func (o *UsergroupmembershipResponse) HasSUserFirstname() bool {
	if o != nil && !IsNil(o.SUserFirstname) {
		return true
	}

	return false
}

// SetSUserFirstname gets a reference to the given string and assigns it to the SUserFirstname field.
func (o *UsergroupmembershipResponse) SetSUserFirstname(v string) {
	o.SUserFirstname = &v
}

// GetSUserLastname returns the SUserLastname field value if set, zero value otherwise.
func (o *UsergroupmembershipResponse) GetSUserLastname() string {
	if o == nil || IsNil(o.SUserLastname) {
		var ret string
		return ret
	}
	return *o.SUserLastname
}

// GetSUserLastnameOk returns a tuple with the SUserLastname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsergroupmembershipResponse) GetSUserLastnameOk() (*string, bool) {
	if o == nil || IsNil(o.SUserLastname) {
		return nil, false
	}
	return o.SUserLastname, true
}

// HasSUserLastname returns a boolean if a field has been set.
func (o *UsergroupmembershipResponse) HasSUserLastname() bool {
	if o != nil && !IsNil(o.SUserLastname) {
		return true
	}

	return false
}

// SetSUserLastname gets a reference to the given string and assigns it to the SUserLastname field.
func (o *UsergroupmembershipResponse) SetSUserLastname(v string) {
	o.SUserLastname = &v
}

// GetSUserLoginname returns the SUserLoginname field value if set, zero value otherwise.
func (o *UsergroupmembershipResponse) GetSUserLoginname() string {
	if o == nil || IsNil(o.SUserLoginname) {
		var ret string
		return ret
	}
	return *o.SUserLoginname
}

// GetSUserLoginnameOk returns a tuple with the SUserLoginname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsergroupmembershipResponse) GetSUserLoginnameOk() (*string, bool) {
	if o == nil || IsNil(o.SUserLoginname) {
		return nil, false
	}
	return o.SUserLoginname, true
}

// HasSUserLoginname returns a boolean if a field has been set.
func (o *UsergroupmembershipResponse) HasSUserLoginname() bool {
	if o != nil && !IsNil(o.SUserLoginname) {
		return true
	}

	return false
}

// SetSUserLoginname gets a reference to the given string and assigns it to the SUserLoginname field.
func (o *UsergroupmembershipResponse) SetSUserLoginname(v string) {
	o.SUserLoginname = &v
}

// GetSEmailAddress returns the SEmailAddress field value if set, zero value otherwise.
func (o *UsergroupmembershipResponse) GetSEmailAddress() string {
	if o == nil || IsNil(o.SEmailAddress) {
		var ret string
		return ret
	}
	return *o.SEmailAddress
}

// GetSEmailAddressOk returns a tuple with the SEmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsergroupmembershipResponse) GetSEmailAddressOk() (*string, bool) {
	if o == nil || IsNil(o.SEmailAddress) {
		return nil, false
	}
	return o.SEmailAddress, true
}

// HasSEmailAddress returns a boolean if a field has been set.
func (o *UsergroupmembershipResponse) HasSEmailAddress() bool {
	if o != nil && !IsNil(o.SEmailAddress) {
		return true
	}

	return false
}

// SetSEmailAddress gets a reference to the given string and assigns it to the SEmailAddress field.
func (o *UsergroupmembershipResponse) SetSEmailAddress(v string) {
	o.SEmailAddress = &v
}

// GetSUsergroupNameX returns the SUsergroupNameX field value
func (o *UsergroupmembershipResponse) GetSUsergroupNameX() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SUsergroupNameX
}

// GetSUsergroupNameXOk returns a tuple with the SUsergroupNameX field value
// and a boolean to check if the value has been set.
func (o *UsergroupmembershipResponse) GetSUsergroupNameXOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SUsergroupNameX, true
}

// SetSUsergroupNameX sets field value
func (o *UsergroupmembershipResponse) SetSUsergroupNameX(v string) {
	o.SUsergroupNameX = v
}

// GetSUsergroupexternalName returns the SUsergroupexternalName field value if set, zero value otherwise.
func (o *UsergroupmembershipResponse) GetSUsergroupexternalName() string {
	if o == nil || IsNil(o.SUsergroupexternalName) {
		var ret string
		return ret
	}
	return *o.SUsergroupexternalName
}

// GetSUsergroupexternalNameOk returns a tuple with the SUsergroupexternalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsergroupmembershipResponse) GetSUsergroupexternalNameOk() (*string, bool) {
	if o == nil || IsNil(o.SUsergroupexternalName) {
		return nil, false
	}
	return o.SUsergroupexternalName, true
}

// HasSUsergroupexternalName returns a boolean if a field has been set.
func (o *UsergroupmembershipResponse) HasSUsergroupexternalName() bool {
	if o != nil && !IsNil(o.SUsergroupexternalName) {
		return true
	}

	return false
}

// SetSUsergroupexternalName gets a reference to the given string and assigns it to the SUsergroupexternalName field.
func (o *UsergroupmembershipResponse) SetSUsergroupexternalName(v string) {
	o.SUsergroupexternalName = &v
}

func (o UsergroupmembershipResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsergroupmembershipResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiUsergroupmembershipID"] = o.PkiUsergroupmembershipID
	toSerialize["fkiUsergroupID"] = o.FkiUsergroupID
	if !IsNil(o.FkiUserID) {
		toSerialize["fkiUserID"] = o.FkiUserID
	}
	if !IsNil(o.FkiUsergroupexternalID) {
		toSerialize["fkiUsergroupexternalID"] = o.FkiUsergroupexternalID
	}
	if !IsNil(o.SUserFirstname) {
		toSerialize["sUserFirstname"] = o.SUserFirstname
	}
	if !IsNil(o.SUserLastname) {
		toSerialize["sUserLastname"] = o.SUserLastname
	}
	if !IsNil(o.SUserLoginname) {
		toSerialize["sUserLoginname"] = o.SUserLoginname
	}
	if !IsNil(o.SEmailAddress) {
		toSerialize["sEmailAddress"] = o.SEmailAddress
	}
	toSerialize["sUsergroupNameX"] = o.SUsergroupNameX
	if !IsNil(o.SUsergroupexternalName) {
		toSerialize["sUsergroupexternalName"] = o.SUsergroupexternalName
	}
	return toSerialize, nil
}

func (o *UsergroupmembershipResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiUsergroupmembershipID",
		"fkiUsergroupID",
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

	varUsergroupmembershipResponse := _UsergroupmembershipResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUsergroupmembershipResponse)

	if err != nil {
		return err
	}

	*o = UsergroupmembershipResponse(varUsergroupmembershipResponse)

	return err
}

type NullableUsergroupmembershipResponse struct {
	value *UsergroupmembershipResponse
	isSet bool
}

func (v NullableUsergroupmembershipResponse) Get() *UsergroupmembershipResponse {
	return v.value
}

func (v *NullableUsergroupmembershipResponse) Set(val *UsergroupmembershipResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUsergroupmembershipResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUsergroupmembershipResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsergroupmembershipResponse(val *UsergroupmembershipResponse) *NullableUsergroupmembershipResponse {
	return &NullableUsergroupmembershipResponse{value: val, isSet: true}
}

func (v NullableUsergroupmembershipResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsergroupmembershipResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


