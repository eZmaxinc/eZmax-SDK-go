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

// checks if the CommonAuditdetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommonAuditdetail{}

// CommonAuditdetail Gives informations about the user that created the object or the last user to have modified it.  If the object was never modified after creation, both Created and Modified informations will be the same. 
type CommonAuditdetail struct {
	// The unique ID of the User
	FkiUserID int32 `json:"fkiUserID"`
	// The unique ID of the Apikey
	FkiApikeyID *int32 `json:"fkiApikeyID,omitempty"`
	// The login name of the User.
	SUserLoginname string `json:"sUserLoginname"`
	// The last name of the user
	SUserLastname string `json:"sUserLastname"`
	// The first name of the user
	SUserFirstname string `json:"sUserFirstname"`
	// The description of the Apikey in the language of the requester
	SApikeyDescriptionX *string `json:"sApikeyDescriptionX,omitempty"`
	// Represent a Date Time. The timezone is the one configured in the User's profile.
	DtAuditdetailDate string `json:"dtAuditdetailDate"`
}

type _CommonAuditdetail CommonAuditdetail

// NewCommonAuditdetail instantiates a new CommonAuditdetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonAuditdetail(fkiUserID int32, sUserLoginname string, sUserLastname string, sUserFirstname string, dtAuditdetailDate string) *CommonAuditdetail {
	this := CommonAuditdetail{}
	this.FkiUserID = fkiUserID
	this.SUserLoginname = sUserLoginname
	this.SUserLastname = sUserLastname
	this.SUserFirstname = sUserFirstname
	this.DtAuditdetailDate = dtAuditdetailDate
	return &this
}

// NewCommonAuditdetailWithDefaults instantiates a new CommonAuditdetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonAuditdetailWithDefaults() *CommonAuditdetail {
	this := CommonAuditdetail{}
	return &this
}

// GetFkiUserID returns the FkiUserID field value
func (o *CommonAuditdetail) GetFkiUserID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiUserID
}

// GetFkiUserIDOk returns a tuple with the FkiUserID field value
// and a boolean to check if the value has been set.
func (o *CommonAuditdetail) GetFkiUserIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiUserID, true
}

// SetFkiUserID sets field value
func (o *CommonAuditdetail) SetFkiUserID(v int32) {
	o.FkiUserID = v
}

// GetFkiApikeyID returns the FkiApikeyID field value if set, zero value otherwise.
func (o *CommonAuditdetail) GetFkiApikeyID() int32 {
	if o == nil || IsNil(o.FkiApikeyID) {
		var ret int32
		return ret
	}
	return *o.FkiApikeyID
}

// GetFkiApikeyIDOk returns a tuple with the FkiApikeyID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonAuditdetail) GetFkiApikeyIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiApikeyID) {
		return nil, false
	}
	return o.FkiApikeyID, true
}

// HasFkiApikeyID returns a boolean if a field has been set.
func (o *CommonAuditdetail) HasFkiApikeyID() bool {
	if o != nil && !IsNil(o.FkiApikeyID) {
		return true
	}

	return false
}

// SetFkiApikeyID gets a reference to the given int32 and assigns it to the FkiApikeyID field.
func (o *CommonAuditdetail) SetFkiApikeyID(v int32) {
	o.FkiApikeyID = &v
}

// GetSUserLoginname returns the SUserLoginname field value
func (o *CommonAuditdetail) GetSUserLoginname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SUserLoginname
}

// GetSUserLoginnameOk returns a tuple with the SUserLoginname field value
// and a boolean to check if the value has been set.
func (o *CommonAuditdetail) GetSUserLoginnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SUserLoginname, true
}

// SetSUserLoginname sets field value
func (o *CommonAuditdetail) SetSUserLoginname(v string) {
	o.SUserLoginname = v
}

// GetSUserLastname returns the SUserLastname field value
func (o *CommonAuditdetail) GetSUserLastname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SUserLastname
}

// GetSUserLastnameOk returns a tuple with the SUserLastname field value
// and a boolean to check if the value has been set.
func (o *CommonAuditdetail) GetSUserLastnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SUserLastname, true
}

// SetSUserLastname sets field value
func (o *CommonAuditdetail) SetSUserLastname(v string) {
	o.SUserLastname = v
}

// GetSUserFirstname returns the SUserFirstname field value
func (o *CommonAuditdetail) GetSUserFirstname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SUserFirstname
}

// GetSUserFirstnameOk returns a tuple with the SUserFirstname field value
// and a boolean to check if the value has been set.
func (o *CommonAuditdetail) GetSUserFirstnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SUserFirstname, true
}

// SetSUserFirstname sets field value
func (o *CommonAuditdetail) SetSUserFirstname(v string) {
	o.SUserFirstname = v
}

// GetSApikeyDescriptionX returns the SApikeyDescriptionX field value if set, zero value otherwise.
func (o *CommonAuditdetail) GetSApikeyDescriptionX() string {
	if o == nil || IsNil(o.SApikeyDescriptionX) {
		var ret string
		return ret
	}
	return *o.SApikeyDescriptionX
}

// GetSApikeyDescriptionXOk returns a tuple with the SApikeyDescriptionX field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonAuditdetail) GetSApikeyDescriptionXOk() (*string, bool) {
	if o == nil || IsNil(o.SApikeyDescriptionX) {
		return nil, false
	}
	return o.SApikeyDescriptionX, true
}

// HasSApikeyDescriptionX returns a boolean if a field has been set.
func (o *CommonAuditdetail) HasSApikeyDescriptionX() bool {
	if o != nil && !IsNil(o.SApikeyDescriptionX) {
		return true
	}

	return false
}

// SetSApikeyDescriptionX gets a reference to the given string and assigns it to the SApikeyDescriptionX field.
func (o *CommonAuditdetail) SetSApikeyDescriptionX(v string) {
	o.SApikeyDescriptionX = &v
}

// GetDtAuditdetailDate returns the DtAuditdetailDate field value
func (o *CommonAuditdetail) GetDtAuditdetailDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DtAuditdetailDate
}

// GetDtAuditdetailDateOk returns a tuple with the DtAuditdetailDate field value
// and a boolean to check if the value has been set.
func (o *CommonAuditdetail) GetDtAuditdetailDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DtAuditdetailDate, true
}

// SetDtAuditdetailDate sets field value
func (o *CommonAuditdetail) SetDtAuditdetailDate(v string) {
	o.DtAuditdetailDate = v
}

func (o CommonAuditdetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommonAuditdetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fkiUserID"] = o.FkiUserID
	if !IsNil(o.FkiApikeyID) {
		toSerialize["fkiApikeyID"] = o.FkiApikeyID
	}
	toSerialize["sUserLoginname"] = o.SUserLoginname
	toSerialize["sUserLastname"] = o.SUserLastname
	toSerialize["sUserFirstname"] = o.SUserFirstname
	if !IsNil(o.SApikeyDescriptionX) {
		toSerialize["sApikeyDescriptionX"] = o.SApikeyDescriptionX
	}
	toSerialize["dtAuditdetailDate"] = o.DtAuditdetailDate
	return toSerialize, nil
}

func (o *CommonAuditdetail) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fkiUserID",
		"sUserLoginname",
		"sUserLastname",
		"sUserFirstname",
		"dtAuditdetailDate",
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

	varCommonAuditdetail := _CommonAuditdetail{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCommonAuditdetail)

	if err != nil {
		return err
	}

	*o = CommonAuditdetail(varCommonAuditdetail)

	return err
}

type NullableCommonAuditdetail struct {
	value *CommonAuditdetail
	isSet bool
}

func (v NullableCommonAuditdetail) Get() *CommonAuditdetail {
	return v.value
}

func (v *NullableCommonAuditdetail) Set(val *CommonAuditdetail) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonAuditdetail) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonAuditdetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonAuditdetail(val *CommonAuditdetail) *NullableCommonAuditdetail {
	return &NullableCommonAuditdetail{value: val, isSet: true}
}

func (v NullableCommonAuditdetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonAuditdetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


