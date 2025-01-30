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

// checks if the ContactResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContactResponse{}

// ContactResponse A Contact Object
type ContactResponse struct {
	// The unique ID of the Contact
	PkiContactID int32 `json:"pkiContactID"`
	// The unique ID of the Language.  Valid values:  |Value|Description| |-|-| |1|French| |2|English|
	FkiLanguageID int32 `json:"fkiLanguageID"`
	// The unique ID of the Contacttitle.  Valid values:  |Value|Description| |-|-| |1|Ms.| |2|Mr.| |4|(Blank)| |5|Me (For Notaries)|
	FkiContacttitleID int32 `json:"fkiContacttitleID"`
	// The unique ID of the Contactinformations
	FkiContactinformationsID int32 `json:"fkiContactinformationsID"`
	// The Birth Date of the contact
	DtContactBirthdate *string `json:"dtContactBirthdate,omitempty"`
	EContactType FieldEContactType `json:"eContactType"`
	// The First name of the contact
	SContactFirstname string `json:"sContactFirstname" validate:"regexp=^.{1,20}$"`
	// The Last name of the contact
	SContactLastname string `json:"sContactLastname" validate:"regexp=^.{1,25}$"`
	// The Company name of the contact
	SContactCompany *string `json:"sContactCompany,omitempty"`
	// The occupation of the Contact
	SContactOccupation *string `json:"sContactOccupation,omitempty" validate:"regexp=^.{0,50}$"`
	// The note of the Contact
	TContactNote *string `json:"tContactNote,omitempty" validate:"regexp=^.{0,32000}$"`
	// Whether the contact is active or not
	BContactIsactive bool `json:"bContactIsactive"`
	ObjContactinformations ContactinformationsResponseCompound `json:"objContactinformations"`
}

type _ContactResponse ContactResponse

// NewContactResponse instantiates a new ContactResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContactResponse(pkiContactID int32, fkiLanguageID int32, fkiContacttitleID int32, fkiContactinformationsID int32, eContactType FieldEContactType, sContactFirstname string, sContactLastname string, bContactIsactive bool, objContactinformations ContactinformationsResponseCompound) *ContactResponse {
	this := ContactResponse{}
	this.PkiContactID = pkiContactID
	this.FkiLanguageID = fkiLanguageID
	this.FkiContacttitleID = fkiContacttitleID
	this.FkiContactinformationsID = fkiContactinformationsID
	this.EContactType = eContactType
	this.SContactFirstname = sContactFirstname
	this.SContactLastname = sContactLastname
	this.BContactIsactive = bContactIsactive
	this.ObjContactinformations = objContactinformations
	return &this
}

// NewContactResponseWithDefaults instantiates a new ContactResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContactResponseWithDefaults() *ContactResponse {
	this := ContactResponse{}
	return &this
}

// GetPkiContactID returns the PkiContactID field value
func (o *ContactResponse) GetPkiContactID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiContactID
}

// GetPkiContactIDOk returns a tuple with the PkiContactID field value
// and a boolean to check if the value has been set.
func (o *ContactResponse) GetPkiContactIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiContactID, true
}

// SetPkiContactID sets field value
func (o *ContactResponse) SetPkiContactID(v int32) {
	o.PkiContactID = v
}

// GetFkiLanguageID returns the FkiLanguageID field value
func (o *ContactResponse) GetFkiLanguageID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiLanguageID
}

// GetFkiLanguageIDOk returns a tuple with the FkiLanguageID field value
// and a boolean to check if the value has been set.
func (o *ContactResponse) GetFkiLanguageIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiLanguageID, true
}

// SetFkiLanguageID sets field value
func (o *ContactResponse) SetFkiLanguageID(v int32) {
	o.FkiLanguageID = v
}

// GetFkiContacttitleID returns the FkiContacttitleID field value
func (o *ContactResponse) GetFkiContacttitleID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiContacttitleID
}

// GetFkiContacttitleIDOk returns a tuple with the FkiContacttitleID field value
// and a boolean to check if the value has been set.
func (o *ContactResponse) GetFkiContacttitleIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiContacttitleID, true
}

// SetFkiContacttitleID sets field value
func (o *ContactResponse) SetFkiContacttitleID(v int32) {
	o.FkiContacttitleID = v
}

// GetFkiContactinformationsID returns the FkiContactinformationsID field value
func (o *ContactResponse) GetFkiContactinformationsID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiContactinformationsID
}

// GetFkiContactinformationsIDOk returns a tuple with the FkiContactinformationsID field value
// and a boolean to check if the value has been set.
func (o *ContactResponse) GetFkiContactinformationsIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiContactinformationsID, true
}

// SetFkiContactinformationsID sets field value
func (o *ContactResponse) SetFkiContactinformationsID(v int32) {
	o.FkiContactinformationsID = v
}

// GetDtContactBirthdate returns the DtContactBirthdate field value if set, zero value otherwise.
func (o *ContactResponse) GetDtContactBirthdate() string {
	if o == nil || IsNil(o.DtContactBirthdate) {
		var ret string
		return ret
	}
	return *o.DtContactBirthdate
}

// GetDtContactBirthdateOk returns a tuple with the DtContactBirthdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactResponse) GetDtContactBirthdateOk() (*string, bool) {
	if o == nil || IsNil(o.DtContactBirthdate) {
		return nil, false
	}
	return o.DtContactBirthdate, true
}

// HasDtContactBirthdate returns a boolean if a field has been set.
func (o *ContactResponse) HasDtContactBirthdate() bool {
	if o != nil && !IsNil(o.DtContactBirthdate) {
		return true
	}

	return false
}

// SetDtContactBirthdate gets a reference to the given string and assigns it to the DtContactBirthdate field.
func (o *ContactResponse) SetDtContactBirthdate(v string) {
	o.DtContactBirthdate = &v
}

// GetEContactType returns the EContactType field value
func (o *ContactResponse) GetEContactType() FieldEContactType {
	if o == nil {
		var ret FieldEContactType
		return ret
	}

	return o.EContactType
}

// GetEContactTypeOk returns a tuple with the EContactType field value
// and a boolean to check if the value has been set.
func (o *ContactResponse) GetEContactTypeOk() (*FieldEContactType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EContactType, true
}

// SetEContactType sets field value
func (o *ContactResponse) SetEContactType(v FieldEContactType) {
	o.EContactType = v
}

// GetSContactFirstname returns the SContactFirstname field value
func (o *ContactResponse) GetSContactFirstname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SContactFirstname
}

// GetSContactFirstnameOk returns a tuple with the SContactFirstname field value
// and a boolean to check if the value has been set.
func (o *ContactResponse) GetSContactFirstnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SContactFirstname, true
}

// SetSContactFirstname sets field value
func (o *ContactResponse) SetSContactFirstname(v string) {
	o.SContactFirstname = v
}

// GetSContactLastname returns the SContactLastname field value
func (o *ContactResponse) GetSContactLastname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SContactLastname
}

// GetSContactLastnameOk returns a tuple with the SContactLastname field value
// and a boolean to check if the value has been set.
func (o *ContactResponse) GetSContactLastnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SContactLastname, true
}

// SetSContactLastname sets field value
func (o *ContactResponse) SetSContactLastname(v string) {
	o.SContactLastname = v
}

// GetSContactCompany returns the SContactCompany field value if set, zero value otherwise.
func (o *ContactResponse) GetSContactCompany() string {
	if o == nil || IsNil(o.SContactCompany) {
		var ret string
		return ret
	}
	return *o.SContactCompany
}

// GetSContactCompanyOk returns a tuple with the SContactCompany field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactResponse) GetSContactCompanyOk() (*string, bool) {
	if o == nil || IsNil(o.SContactCompany) {
		return nil, false
	}
	return o.SContactCompany, true
}

// HasSContactCompany returns a boolean if a field has been set.
func (o *ContactResponse) HasSContactCompany() bool {
	if o != nil && !IsNil(o.SContactCompany) {
		return true
	}

	return false
}

// SetSContactCompany gets a reference to the given string and assigns it to the SContactCompany field.
func (o *ContactResponse) SetSContactCompany(v string) {
	o.SContactCompany = &v
}

// GetSContactOccupation returns the SContactOccupation field value if set, zero value otherwise.
func (o *ContactResponse) GetSContactOccupation() string {
	if o == nil || IsNil(o.SContactOccupation) {
		var ret string
		return ret
	}
	return *o.SContactOccupation
}

// GetSContactOccupationOk returns a tuple with the SContactOccupation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactResponse) GetSContactOccupationOk() (*string, bool) {
	if o == nil || IsNil(o.SContactOccupation) {
		return nil, false
	}
	return o.SContactOccupation, true
}

// HasSContactOccupation returns a boolean if a field has been set.
func (o *ContactResponse) HasSContactOccupation() bool {
	if o != nil && !IsNil(o.SContactOccupation) {
		return true
	}

	return false
}

// SetSContactOccupation gets a reference to the given string and assigns it to the SContactOccupation field.
func (o *ContactResponse) SetSContactOccupation(v string) {
	o.SContactOccupation = &v
}

// GetTContactNote returns the TContactNote field value if set, zero value otherwise.
func (o *ContactResponse) GetTContactNote() string {
	if o == nil || IsNil(o.TContactNote) {
		var ret string
		return ret
	}
	return *o.TContactNote
}

// GetTContactNoteOk returns a tuple with the TContactNote field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactResponse) GetTContactNoteOk() (*string, bool) {
	if o == nil || IsNil(o.TContactNote) {
		return nil, false
	}
	return o.TContactNote, true
}

// HasTContactNote returns a boolean if a field has been set.
func (o *ContactResponse) HasTContactNote() bool {
	if o != nil && !IsNil(o.TContactNote) {
		return true
	}

	return false
}

// SetTContactNote gets a reference to the given string and assigns it to the TContactNote field.
func (o *ContactResponse) SetTContactNote(v string) {
	o.TContactNote = &v
}

// GetBContactIsactive returns the BContactIsactive field value
func (o *ContactResponse) GetBContactIsactive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BContactIsactive
}

// GetBContactIsactiveOk returns a tuple with the BContactIsactive field value
// and a boolean to check if the value has been set.
func (o *ContactResponse) GetBContactIsactiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BContactIsactive, true
}

// SetBContactIsactive sets field value
func (o *ContactResponse) SetBContactIsactive(v bool) {
	o.BContactIsactive = v
}

// GetObjContactinformations returns the ObjContactinformations field value
func (o *ContactResponse) GetObjContactinformations() ContactinformationsResponseCompound {
	if o == nil {
		var ret ContactinformationsResponseCompound
		return ret
	}

	return o.ObjContactinformations
}

// GetObjContactinformationsOk returns a tuple with the ObjContactinformations field value
// and a boolean to check if the value has been set.
func (o *ContactResponse) GetObjContactinformationsOk() (*ContactinformationsResponseCompound, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjContactinformations, true
}

// SetObjContactinformations sets field value
func (o *ContactResponse) SetObjContactinformations(v ContactinformationsResponseCompound) {
	o.ObjContactinformations = v
}

func (o ContactResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContactResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiContactID"] = o.PkiContactID
	toSerialize["fkiLanguageID"] = o.FkiLanguageID
	toSerialize["fkiContacttitleID"] = o.FkiContacttitleID
	toSerialize["fkiContactinformationsID"] = o.FkiContactinformationsID
	if !IsNil(o.DtContactBirthdate) {
		toSerialize["dtContactBirthdate"] = o.DtContactBirthdate
	}
	toSerialize["eContactType"] = o.EContactType
	toSerialize["sContactFirstname"] = o.SContactFirstname
	toSerialize["sContactLastname"] = o.SContactLastname
	if !IsNil(o.SContactCompany) {
		toSerialize["sContactCompany"] = o.SContactCompany
	}
	if !IsNil(o.SContactOccupation) {
		toSerialize["sContactOccupation"] = o.SContactOccupation
	}
	if !IsNil(o.TContactNote) {
		toSerialize["tContactNote"] = o.TContactNote
	}
	toSerialize["bContactIsactive"] = o.BContactIsactive
	toSerialize["objContactinformations"] = o.ObjContactinformations
	return toSerialize, nil
}

func (o *ContactResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiContactID",
		"fkiLanguageID",
		"fkiContacttitleID",
		"fkiContactinformationsID",
		"eContactType",
		"sContactFirstname",
		"sContactLastname",
		"bContactIsactive",
		"objContactinformations",
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

	varContactResponse := _ContactResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varContactResponse)

	if err != nil {
		return err
	}

	*o = ContactResponse(varContactResponse)

	return err
}

type NullableContactResponse struct {
	value *ContactResponse
	isSet bool
}

func (v NullableContactResponse) Get() *ContactResponse {
	return v.value
}

func (v *NullableContactResponse) Set(val *ContactResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableContactResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableContactResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContactResponse(val *ContactResponse) *NullableContactResponse {
	return &NullableContactResponse{value: val, isSet: true}
}

func (v NullableContactResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContactResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


