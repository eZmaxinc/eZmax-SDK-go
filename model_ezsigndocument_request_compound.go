/*
eZmax API Definition

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.1.1
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
)

// EzsigndocumentRequestCompound An Ezsigndocument Object and children to create a complete structure
type EzsigndocumentRequestCompound struct {
	// Indicates where to look for the document binary content.
	EEzsigndocumentSource string `json:"eEzsigndocumentSource"`
	// Indicates the format of the document.
	EEzsigndocumentFormat string `json:"eEzsigndocumentFormat"`
	// The Base64 encoded binary content of the document.  This field is Required when eEzsigndocumentSource = Base64.
	SEzsigndocumentBase64 *string `json:"sEzsigndocumentBase64,omitempty"`
	// If the source document is password protected, the password to open/modify it.
	SEzsigndocumentPassword *string `json:"sEzsigndocumentPassword,omitempty"`
	// The unique ID of the Ezsignfolder
	FkiEzsignfolderID int32 `json:"fkiEzsignfolderID"`
	// The maximum date and time at which the document can be signed.
	DtEzsigndocumentDuedate string `json:"dtEzsigndocumentDuedate"`
	// The unique ID of the Language.  Valid values:  |Value|Description| |-|-| |1|French| |2|English|
	FkiLanguageID int32 `json:"fkiLanguageID"`
	// The name of the document that will be presented to Ezsignfoldersignerassociations
	SEzsigndocumentName string `json:"sEzsigndocumentName"`
}

// NewEzsigndocumentRequestCompound instantiates a new EzsigndocumentRequestCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigndocumentRequestCompound(eEzsigndocumentSource string, eEzsigndocumentFormat string, fkiEzsignfolderID int32, dtEzsigndocumentDuedate string, fkiLanguageID int32, sEzsigndocumentName string) *EzsigndocumentRequestCompound {
	this := EzsigndocumentRequestCompound{}
	this.EEzsigndocumentSource = eEzsigndocumentSource
	this.EEzsigndocumentFormat = eEzsigndocumentFormat
	var sEzsigndocumentPassword string = ""
	this.SEzsigndocumentPassword = &sEzsigndocumentPassword
	this.FkiEzsignfolderID = fkiEzsignfolderID
	this.DtEzsigndocumentDuedate = dtEzsigndocumentDuedate
	this.FkiLanguageID = fkiLanguageID
	this.SEzsigndocumentName = sEzsigndocumentName
	return &this
}

// NewEzsigndocumentRequestCompoundWithDefaults instantiates a new EzsigndocumentRequestCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigndocumentRequestCompoundWithDefaults() *EzsigndocumentRequestCompound {
	this := EzsigndocumentRequestCompound{}
	var sEzsigndocumentPassword string = ""
	this.SEzsigndocumentPassword = &sEzsigndocumentPassword
	return &this
}

// GetEEzsigndocumentSource returns the EEzsigndocumentSource field value
func (o *EzsigndocumentRequestCompound) GetEEzsigndocumentSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EEzsigndocumentSource
}

// GetEEzsigndocumentSourceOk returns a tuple with the EEzsigndocumentSource field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentRequestCompound) GetEEzsigndocumentSourceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EEzsigndocumentSource, true
}

// SetEEzsigndocumentSource sets field value
func (o *EzsigndocumentRequestCompound) SetEEzsigndocumentSource(v string) {
	o.EEzsigndocumentSource = v
}

// GetEEzsigndocumentFormat returns the EEzsigndocumentFormat field value
func (o *EzsigndocumentRequestCompound) GetEEzsigndocumentFormat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EEzsigndocumentFormat
}

// GetEEzsigndocumentFormatOk returns a tuple with the EEzsigndocumentFormat field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentRequestCompound) GetEEzsigndocumentFormatOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EEzsigndocumentFormat, true
}

// SetEEzsigndocumentFormat sets field value
func (o *EzsigndocumentRequestCompound) SetEEzsigndocumentFormat(v string) {
	o.EEzsigndocumentFormat = v
}

// GetSEzsigndocumentBase64 returns the SEzsigndocumentBase64 field value if set, zero value otherwise.
func (o *EzsigndocumentRequestCompound) GetSEzsigndocumentBase64() string {
	if o == nil || o.SEzsigndocumentBase64 == nil {
		var ret string
		return ret
	}
	return *o.SEzsigndocumentBase64
}

// GetSEzsigndocumentBase64Ok returns a tuple with the SEzsigndocumentBase64 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigndocumentRequestCompound) GetSEzsigndocumentBase64Ok() (*string, bool) {
	if o == nil || o.SEzsigndocumentBase64 == nil {
		return nil, false
	}
	return o.SEzsigndocumentBase64, true
}

// HasSEzsigndocumentBase64 returns a boolean if a field has been set.
func (o *EzsigndocumentRequestCompound) HasSEzsigndocumentBase64() bool {
	if o != nil && o.SEzsigndocumentBase64 != nil {
		return true
	}

	return false
}

// SetSEzsigndocumentBase64 gets a reference to the given string and assigns it to the SEzsigndocumentBase64 field.
func (o *EzsigndocumentRequestCompound) SetSEzsigndocumentBase64(v string) {
	o.SEzsigndocumentBase64 = &v
}

// GetSEzsigndocumentPassword returns the SEzsigndocumentPassword field value if set, zero value otherwise.
func (o *EzsigndocumentRequestCompound) GetSEzsigndocumentPassword() string {
	if o == nil || o.SEzsigndocumentPassword == nil {
		var ret string
		return ret
	}
	return *o.SEzsigndocumentPassword
}

// GetSEzsigndocumentPasswordOk returns a tuple with the SEzsigndocumentPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigndocumentRequestCompound) GetSEzsigndocumentPasswordOk() (*string, bool) {
	if o == nil || o.SEzsigndocumentPassword == nil {
		return nil, false
	}
	return o.SEzsigndocumentPassword, true
}

// HasSEzsigndocumentPassword returns a boolean if a field has been set.
func (o *EzsigndocumentRequestCompound) HasSEzsigndocumentPassword() bool {
	if o != nil && o.SEzsigndocumentPassword != nil {
		return true
	}

	return false
}

// SetSEzsigndocumentPassword gets a reference to the given string and assigns it to the SEzsigndocumentPassword field.
func (o *EzsigndocumentRequestCompound) SetSEzsigndocumentPassword(v string) {
	o.SEzsigndocumentPassword = &v
}

// GetFkiEzsignfolderID returns the FkiEzsignfolderID field value
func (o *EzsigndocumentRequestCompound) GetFkiEzsignfolderID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzsignfolderID
}

// GetFkiEzsignfolderIDOk returns a tuple with the FkiEzsignfolderID field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentRequestCompound) GetFkiEzsignfolderIDOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FkiEzsignfolderID, true
}

// SetFkiEzsignfolderID sets field value
func (o *EzsigndocumentRequestCompound) SetFkiEzsignfolderID(v int32) {
	o.FkiEzsignfolderID = v
}

// GetDtEzsigndocumentDuedate returns the DtEzsigndocumentDuedate field value
func (o *EzsigndocumentRequestCompound) GetDtEzsigndocumentDuedate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DtEzsigndocumentDuedate
}

// GetDtEzsigndocumentDuedateOk returns a tuple with the DtEzsigndocumentDuedate field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentRequestCompound) GetDtEzsigndocumentDuedateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DtEzsigndocumentDuedate, true
}

// SetDtEzsigndocumentDuedate sets field value
func (o *EzsigndocumentRequestCompound) SetDtEzsigndocumentDuedate(v string) {
	o.DtEzsigndocumentDuedate = v
}

// GetFkiLanguageID returns the FkiLanguageID field value
func (o *EzsigndocumentRequestCompound) GetFkiLanguageID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiLanguageID
}

// GetFkiLanguageIDOk returns a tuple with the FkiLanguageID field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentRequestCompound) GetFkiLanguageIDOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FkiLanguageID, true
}

// SetFkiLanguageID sets field value
func (o *EzsigndocumentRequestCompound) SetFkiLanguageID(v int32) {
	o.FkiLanguageID = v
}

// GetSEzsigndocumentName returns the SEzsigndocumentName field value
func (o *EzsigndocumentRequestCompound) GetSEzsigndocumentName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzsigndocumentName
}

// GetSEzsigndocumentNameOk returns a tuple with the SEzsigndocumentName field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentRequestCompound) GetSEzsigndocumentNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SEzsigndocumentName, true
}

// SetSEzsigndocumentName sets field value
func (o *EzsigndocumentRequestCompound) SetSEzsigndocumentName(v string) {
	o.SEzsigndocumentName = v
}

func (o EzsigndocumentRequestCompound) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["eEzsigndocumentSource"] = o.EEzsigndocumentSource
	}
	if true {
		toSerialize["eEzsigndocumentFormat"] = o.EEzsigndocumentFormat
	}
	if o.SEzsigndocumentBase64 != nil {
		toSerialize["sEzsigndocumentBase64"] = o.SEzsigndocumentBase64
	}
	if o.SEzsigndocumentPassword != nil {
		toSerialize["sEzsigndocumentPassword"] = o.SEzsigndocumentPassword
	}
	if true {
		toSerialize["fkiEzsignfolderID"] = o.FkiEzsignfolderID
	}
	if true {
		toSerialize["dtEzsigndocumentDuedate"] = o.DtEzsigndocumentDuedate
	}
	if true {
		toSerialize["fkiLanguageID"] = o.FkiLanguageID
	}
	if true {
		toSerialize["sEzsigndocumentName"] = o.SEzsigndocumentName
	}
	return json.Marshal(toSerialize)
}

type NullableEzsigndocumentRequestCompound struct {
	value *EzsigndocumentRequestCompound
	isSet bool
}

func (v NullableEzsigndocumentRequestCompound) Get() *EzsigndocumentRequestCompound {
	return v.value
}

func (v *NullableEzsigndocumentRequestCompound) Set(val *EzsigndocumentRequestCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigndocumentRequestCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigndocumentRequestCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigndocumentRequestCompound(val *EzsigndocumentRequestCompound) *NullableEzsigndocumentRequestCompound {
	return &NullableEzsigndocumentRequestCompound{value: val, isSet: true}
}

func (v NullableEzsigndocumentRequestCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigndocumentRequestCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


