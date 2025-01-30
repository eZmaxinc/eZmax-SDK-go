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

// checks if the ContactRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContactRequest{}

// ContactRequest A Contact Object
type ContactRequest struct {
	// The unique ID of the Contacttitle.  Valid values:  |Value|Description| |-|-| |1|Ms.| |2|Mr.| |4|(Blank)| |5|Me (For Notaries)|
	FkiContacttitleID int32 `json:"fkiContacttitleID"`
	// The unique ID of the Language.  Valid values:  |Value|Description| |-|-| |1|French| |2|English|
	FkiLanguageID int32 `json:"fkiLanguageID"`
	// The First name of the contact
	SContactFirstname string `json:"sContactFirstname" validate:"regexp=^.{1,20}$"`
	// The Last name of the contact
	SContactLastname string `json:"sContactLastname" validate:"regexp=^.{1,25}$"`
	// The Company name of the contact
	SContactCompany string `json:"sContactCompany"`
	// The Birth Date of the contact
	DtContactBirthdate *string `json:"dtContactBirthdate,omitempty"`
}

type _ContactRequest ContactRequest

// NewContactRequest instantiates a new ContactRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContactRequest(fkiContacttitleID int32, fkiLanguageID int32, sContactFirstname string, sContactLastname string, sContactCompany string) *ContactRequest {
	this := ContactRequest{}
	this.FkiContacttitleID = fkiContacttitleID
	this.FkiLanguageID = fkiLanguageID
	this.SContactFirstname = sContactFirstname
	this.SContactLastname = sContactLastname
	this.SContactCompany = sContactCompany
	return &this
}

// NewContactRequestWithDefaults instantiates a new ContactRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContactRequestWithDefaults() *ContactRequest {
	this := ContactRequest{}
	return &this
}

// GetFkiContacttitleID returns the FkiContacttitleID field value
func (o *ContactRequest) GetFkiContacttitleID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiContacttitleID
}

// GetFkiContacttitleIDOk returns a tuple with the FkiContacttitleID field value
// and a boolean to check if the value has been set.
func (o *ContactRequest) GetFkiContacttitleIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiContacttitleID, true
}

// SetFkiContacttitleID sets field value
func (o *ContactRequest) SetFkiContacttitleID(v int32) {
	o.FkiContacttitleID = v
}

// GetFkiLanguageID returns the FkiLanguageID field value
func (o *ContactRequest) GetFkiLanguageID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiLanguageID
}

// GetFkiLanguageIDOk returns a tuple with the FkiLanguageID field value
// and a boolean to check if the value has been set.
func (o *ContactRequest) GetFkiLanguageIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiLanguageID, true
}

// SetFkiLanguageID sets field value
func (o *ContactRequest) SetFkiLanguageID(v int32) {
	o.FkiLanguageID = v
}

// GetSContactFirstname returns the SContactFirstname field value
func (o *ContactRequest) GetSContactFirstname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SContactFirstname
}

// GetSContactFirstnameOk returns a tuple with the SContactFirstname field value
// and a boolean to check if the value has been set.
func (o *ContactRequest) GetSContactFirstnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SContactFirstname, true
}

// SetSContactFirstname sets field value
func (o *ContactRequest) SetSContactFirstname(v string) {
	o.SContactFirstname = v
}

// GetSContactLastname returns the SContactLastname field value
func (o *ContactRequest) GetSContactLastname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SContactLastname
}

// GetSContactLastnameOk returns a tuple with the SContactLastname field value
// and a boolean to check if the value has been set.
func (o *ContactRequest) GetSContactLastnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SContactLastname, true
}

// SetSContactLastname sets field value
func (o *ContactRequest) SetSContactLastname(v string) {
	o.SContactLastname = v
}

// GetSContactCompany returns the SContactCompany field value
func (o *ContactRequest) GetSContactCompany() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SContactCompany
}

// GetSContactCompanyOk returns a tuple with the SContactCompany field value
// and a boolean to check if the value has been set.
func (o *ContactRequest) GetSContactCompanyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SContactCompany, true
}

// SetSContactCompany sets field value
func (o *ContactRequest) SetSContactCompany(v string) {
	o.SContactCompany = v
}

// GetDtContactBirthdate returns the DtContactBirthdate field value if set, zero value otherwise.
func (o *ContactRequest) GetDtContactBirthdate() string {
	if o == nil || IsNil(o.DtContactBirthdate) {
		var ret string
		return ret
	}
	return *o.DtContactBirthdate
}

// GetDtContactBirthdateOk returns a tuple with the DtContactBirthdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactRequest) GetDtContactBirthdateOk() (*string, bool) {
	if o == nil || IsNil(o.DtContactBirthdate) {
		return nil, false
	}
	return o.DtContactBirthdate, true
}

// HasDtContactBirthdate returns a boolean if a field has been set.
func (o *ContactRequest) HasDtContactBirthdate() bool {
	if o != nil && !IsNil(o.DtContactBirthdate) {
		return true
	}

	return false
}

// SetDtContactBirthdate gets a reference to the given string and assigns it to the DtContactBirthdate field.
func (o *ContactRequest) SetDtContactBirthdate(v string) {
	o.DtContactBirthdate = &v
}

func (o ContactRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContactRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fkiContacttitleID"] = o.FkiContacttitleID
	toSerialize["fkiLanguageID"] = o.FkiLanguageID
	toSerialize["sContactFirstname"] = o.SContactFirstname
	toSerialize["sContactLastname"] = o.SContactLastname
	toSerialize["sContactCompany"] = o.SContactCompany
	if !IsNil(o.DtContactBirthdate) {
		toSerialize["dtContactBirthdate"] = o.DtContactBirthdate
	}
	return toSerialize, nil
}

func (o *ContactRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fkiContacttitleID",
		"fkiLanguageID",
		"sContactFirstname",
		"sContactLastname",
		"sContactCompany",
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

	varContactRequest := _ContactRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varContactRequest)

	if err != nil {
		return err
	}

	*o = ContactRequest(varContactRequest)

	return err
}

type NullableContactRequest struct {
	value *ContactRequest
	isSet bool
}

func (v NullableContactRequest) Get() *ContactRequest {
	return v.value
}

func (v *NullableContactRequest) Set(val *ContactRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableContactRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableContactRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContactRequest(val *ContactRequest) *NullableContactRequest {
	return &NullableContactRequest{value: val, isSet: true}
}

func (v NullableContactRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContactRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


