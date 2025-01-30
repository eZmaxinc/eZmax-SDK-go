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

// checks if the EzsignfolderRequestV3 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignfolderRequestV3{}

// EzsignfolderRequestV3 An Ezsignfolder Object
type EzsignfolderRequestV3 struct {
	// The unique ID of the Ezsignfolder
	PkiEzsignfolderID *int32 `json:"pkiEzsignfolderID,omitempty"`
	// The unique ID of the Ezsignfoldertype.
	FkiEzsignfoldertypeID int32 `json:"fkiEzsignfoldertypeID"`
	// The unique ID of the Timezone
	FkiTimezoneID *int32 `json:"fkiTimezoneID,omitempty"`
	// The unique ID of the Ezsigntsarequirement.  Determine if a Time Stamping Authority should add a timestamp on each of the signature. Valid values:  |Value|Description| |-|-| |1|No. TSA Timestamping will requested. This will make all signatures a lot faster since no round-trip to the TSA server will be required. Timestamping will be made using eZsign server's time.| |2|Best effort. Timestamping from a Time Stamping Authority will be requested but is not mandatory. In the very improbable case it cannot be completed, the timestamping will be made using eZsign server's time. **Additional fee applies**| |3|Mandatory. Timestamping from a Time Stamping Authority will be requested and is mandatory. In the very improbable case it cannot be completed, the signature will fail and the user will be asked to retry. **Additional fee applies**|
	FkiEzsigntsarequirementID *int32 `json:"fkiEzsigntsarequirementID,omitempty"`
	EEzsignfolderDocumentdependency *FieldEEzsignfolderDocumentdependency `json:"eEzsignfolderDocumentdependency,omitempty"`
	// The description of the Ezsignfolder
	SEzsignfolderDescription string `json:"sEzsignfolderDescription" validate:"regexp=^.{0,75}$"`
	// Note about the Ezsignfolder
	TEzsignfolderNote *string `json:"tEzsignfolderNote,omitempty"`
	// A custom text message that will be added to the email sent.
	TEzsignfolderMessage *string `json:"tEzsignfolderMessage,omitempty"`
	// The number of days before the the first reminder sending
	IEzsignfolderSendreminderfirstdays int32 `json:"iEzsignfolderSendreminderfirstdays"`
	// The number of days after the first reminder sending
	IEzsignfolderSendreminderotherdays int32 `json:"iEzsignfolderSendreminderotherdays"`
	// This field can be used to store an External ID from the client's system.  Anything can be stored in this field, it will never be evaluated by the eZmax system and will be returned AS-IS.  To store multiple values, consider using a JSON formatted structure, a URL encoded string, a CSV or any other custom format. 
	SEzsignfolderExternalid *string `json:"sEzsignfolderExternalid,omitempty" validate:"regexp=^.{0,128}$"`
}

type _EzsignfolderRequestV3 EzsignfolderRequestV3

// NewEzsignfolderRequestV3 instantiates a new EzsignfolderRequestV3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignfolderRequestV3(fkiEzsignfoldertypeID int32, sEzsignfolderDescription string, iEzsignfolderSendreminderfirstdays int32, iEzsignfolderSendreminderotherdays int32) *EzsignfolderRequestV3 {
	this := EzsignfolderRequestV3{}
	this.FkiEzsignfoldertypeID = fkiEzsignfoldertypeID
	this.SEzsignfolderDescription = sEzsignfolderDescription
	this.IEzsignfolderSendreminderfirstdays = iEzsignfolderSendreminderfirstdays
	this.IEzsignfolderSendreminderotherdays = iEzsignfolderSendreminderotherdays
	return &this
}

// NewEzsignfolderRequestV3WithDefaults instantiates a new EzsignfolderRequestV3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignfolderRequestV3WithDefaults() *EzsignfolderRequestV3 {
	this := EzsignfolderRequestV3{}
	return &this
}

// GetPkiEzsignfolderID returns the PkiEzsignfolderID field value if set, zero value otherwise.
func (o *EzsignfolderRequestV3) GetPkiEzsignfolderID() int32 {
	if o == nil || IsNil(o.PkiEzsignfolderID) {
		var ret int32
		return ret
	}
	return *o.PkiEzsignfolderID
}

// GetPkiEzsignfolderIDOk returns a tuple with the PkiEzsignfolderID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignfolderRequestV3) GetPkiEzsignfolderIDOk() (*int32, bool) {
	if o == nil || IsNil(o.PkiEzsignfolderID) {
		return nil, false
	}
	return o.PkiEzsignfolderID, true
}

// HasPkiEzsignfolderID returns a boolean if a field has been set.
func (o *EzsignfolderRequestV3) HasPkiEzsignfolderID() bool {
	if o != nil && !IsNil(o.PkiEzsignfolderID) {
		return true
	}

	return false
}

// SetPkiEzsignfolderID gets a reference to the given int32 and assigns it to the PkiEzsignfolderID field.
func (o *EzsignfolderRequestV3) SetPkiEzsignfolderID(v int32) {
	o.PkiEzsignfolderID = &v
}

// GetFkiEzsignfoldertypeID returns the FkiEzsignfoldertypeID field value
func (o *EzsignfolderRequestV3) GetFkiEzsignfoldertypeID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzsignfoldertypeID
}

// GetFkiEzsignfoldertypeIDOk returns a tuple with the FkiEzsignfoldertypeID field value
// and a boolean to check if the value has been set.
func (o *EzsignfolderRequestV3) GetFkiEzsignfoldertypeIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiEzsignfoldertypeID, true
}

// SetFkiEzsignfoldertypeID sets field value
func (o *EzsignfolderRequestV3) SetFkiEzsignfoldertypeID(v int32) {
	o.FkiEzsignfoldertypeID = v
}

// GetFkiTimezoneID returns the FkiTimezoneID field value if set, zero value otherwise.
func (o *EzsignfolderRequestV3) GetFkiTimezoneID() int32 {
	if o == nil || IsNil(o.FkiTimezoneID) {
		var ret int32
		return ret
	}
	return *o.FkiTimezoneID
}

// GetFkiTimezoneIDOk returns a tuple with the FkiTimezoneID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignfolderRequestV3) GetFkiTimezoneIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiTimezoneID) {
		return nil, false
	}
	return o.FkiTimezoneID, true
}

// HasFkiTimezoneID returns a boolean if a field has been set.
func (o *EzsignfolderRequestV3) HasFkiTimezoneID() bool {
	if o != nil && !IsNil(o.FkiTimezoneID) {
		return true
	}

	return false
}

// SetFkiTimezoneID gets a reference to the given int32 and assigns it to the FkiTimezoneID field.
func (o *EzsignfolderRequestV3) SetFkiTimezoneID(v int32) {
	o.FkiTimezoneID = &v
}

// GetFkiEzsigntsarequirementID returns the FkiEzsigntsarequirementID field value if set, zero value otherwise.
func (o *EzsignfolderRequestV3) GetFkiEzsigntsarequirementID() int32 {
	if o == nil || IsNil(o.FkiEzsigntsarequirementID) {
		var ret int32
		return ret
	}
	return *o.FkiEzsigntsarequirementID
}

// GetFkiEzsigntsarequirementIDOk returns a tuple with the FkiEzsigntsarequirementID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignfolderRequestV3) GetFkiEzsigntsarequirementIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiEzsigntsarequirementID) {
		return nil, false
	}
	return o.FkiEzsigntsarequirementID, true
}

// HasFkiEzsigntsarequirementID returns a boolean if a field has been set.
func (o *EzsignfolderRequestV3) HasFkiEzsigntsarequirementID() bool {
	if o != nil && !IsNil(o.FkiEzsigntsarequirementID) {
		return true
	}

	return false
}

// SetFkiEzsigntsarequirementID gets a reference to the given int32 and assigns it to the FkiEzsigntsarequirementID field.
func (o *EzsignfolderRequestV3) SetFkiEzsigntsarequirementID(v int32) {
	o.FkiEzsigntsarequirementID = &v
}

// GetEEzsignfolderDocumentdependency returns the EEzsignfolderDocumentdependency field value if set, zero value otherwise.
func (o *EzsignfolderRequestV3) GetEEzsignfolderDocumentdependency() FieldEEzsignfolderDocumentdependency {
	if o == nil || IsNil(o.EEzsignfolderDocumentdependency) {
		var ret FieldEEzsignfolderDocumentdependency
		return ret
	}
	return *o.EEzsignfolderDocumentdependency
}

// GetEEzsignfolderDocumentdependencyOk returns a tuple with the EEzsignfolderDocumentdependency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignfolderRequestV3) GetEEzsignfolderDocumentdependencyOk() (*FieldEEzsignfolderDocumentdependency, bool) {
	if o == nil || IsNil(o.EEzsignfolderDocumentdependency) {
		return nil, false
	}
	return o.EEzsignfolderDocumentdependency, true
}

// HasEEzsignfolderDocumentdependency returns a boolean if a field has been set.
func (o *EzsignfolderRequestV3) HasEEzsignfolderDocumentdependency() bool {
	if o != nil && !IsNil(o.EEzsignfolderDocumentdependency) {
		return true
	}

	return false
}

// SetEEzsignfolderDocumentdependency gets a reference to the given FieldEEzsignfolderDocumentdependency and assigns it to the EEzsignfolderDocumentdependency field.
func (o *EzsignfolderRequestV3) SetEEzsignfolderDocumentdependency(v FieldEEzsignfolderDocumentdependency) {
	o.EEzsignfolderDocumentdependency = &v
}

// GetSEzsignfolderDescription returns the SEzsignfolderDescription field value
func (o *EzsignfolderRequestV3) GetSEzsignfolderDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzsignfolderDescription
}

// GetSEzsignfolderDescriptionOk returns a tuple with the SEzsignfolderDescription field value
// and a boolean to check if the value has been set.
func (o *EzsignfolderRequestV3) GetSEzsignfolderDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SEzsignfolderDescription, true
}

// SetSEzsignfolderDescription sets field value
func (o *EzsignfolderRequestV3) SetSEzsignfolderDescription(v string) {
	o.SEzsignfolderDescription = v
}

// GetTEzsignfolderNote returns the TEzsignfolderNote field value if set, zero value otherwise.
func (o *EzsignfolderRequestV3) GetTEzsignfolderNote() string {
	if o == nil || IsNil(o.TEzsignfolderNote) {
		var ret string
		return ret
	}
	return *o.TEzsignfolderNote
}

// GetTEzsignfolderNoteOk returns a tuple with the TEzsignfolderNote field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignfolderRequestV3) GetTEzsignfolderNoteOk() (*string, bool) {
	if o == nil || IsNil(o.TEzsignfolderNote) {
		return nil, false
	}
	return o.TEzsignfolderNote, true
}

// HasTEzsignfolderNote returns a boolean if a field has been set.
func (o *EzsignfolderRequestV3) HasTEzsignfolderNote() bool {
	if o != nil && !IsNil(o.TEzsignfolderNote) {
		return true
	}

	return false
}

// SetTEzsignfolderNote gets a reference to the given string and assigns it to the TEzsignfolderNote field.
func (o *EzsignfolderRequestV3) SetTEzsignfolderNote(v string) {
	o.TEzsignfolderNote = &v
}

// GetTEzsignfolderMessage returns the TEzsignfolderMessage field value if set, zero value otherwise.
func (o *EzsignfolderRequestV3) GetTEzsignfolderMessage() string {
	if o == nil || IsNil(o.TEzsignfolderMessage) {
		var ret string
		return ret
	}
	return *o.TEzsignfolderMessage
}

// GetTEzsignfolderMessageOk returns a tuple with the TEzsignfolderMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignfolderRequestV3) GetTEzsignfolderMessageOk() (*string, bool) {
	if o == nil || IsNil(o.TEzsignfolderMessage) {
		return nil, false
	}
	return o.TEzsignfolderMessage, true
}

// HasTEzsignfolderMessage returns a boolean if a field has been set.
func (o *EzsignfolderRequestV3) HasTEzsignfolderMessage() bool {
	if o != nil && !IsNil(o.TEzsignfolderMessage) {
		return true
	}

	return false
}

// SetTEzsignfolderMessage gets a reference to the given string and assigns it to the TEzsignfolderMessage field.
func (o *EzsignfolderRequestV3) SetTEzsignfolderMessage(v string) {
	o.TEzsignfolderMessage = &v
}

// GetIEzsignfolderSendreminderfirstdays returns the IEzsignfolderSendreminderfirstdays field value
func (o *EzsignfolderRequestV3) GetIEzsignfolderSendreminderfirstdays() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsignfolderSendreminderfirstdays
}

// GetIEzsignfolderSendreminderfirstdaysOk returns a tuple with the IEzsignfolderSendreminderfirstdays field value
// and a boolean to check if the value has been set.
func (o *EzsignfolderRequestV3) GetIEzsignfolderSendreminderfirstdaysOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzsignfolderSendreminderfirstdays, true
}

// SetIEzsignfolderSendreminderfirstdays sets field value
func (o *EzsignfolderRequestV3) SetIEzsignfolderSendreminderfirstdays(v int32) {
	o.IEzsignfolderSendreminderfirstdays = v
}

// GetIEzsignfolderSendreminderotherdays returns the IEzsignfolderSendreminderotherdays field value
func (o *EzsignfolderRequestV3) GetIEzsignfolderSendreminderotherdays() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsignfolderSendreminderotherdays
}

// GetIEzsignfolderSendreminderotherdaysOk returns a tuple with the IEzsignfolderSendreminderotherdays field value
// and a boolean to check if the value has been set.
func (o *EzsignfolderRequestV3) GetIEzsignfolderSendreminderotherdaysOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzsignfolderSendreminderotherdays, true
}

// SetIEzsignfolderSendreminderotherdays sets field value
func (o *EzsignfolderRequestV3) SetIEzsignfolderSendreminderotherdays(v int32) {
	o.IEzsignfolderSendreminderotherdays = v
}

// GetSEzsignfolderExternalid returns the SEzsignfolderExternalid field value if set, zero value otherwise.
func (o *EzsignfolderRequestV3) GetSEzsignfolderExternalid() string {
	if o == nil || IsNil(o.SEzsignfolderExternalid) {
		var ret string
		return ret
	}
	return *o.SEzsignfolderExternalid
}

// GetSEzsignfolderExternalidOk returns a tuple with the SEzsignfolderExternalid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignfolderRequestV3) GetSEzsignfolderExternalidOk() (*string, bool) {
	if o == nil || IsNil(o.SEzsignfolderExternalid) {
		return nil, false
	}
	return o.SEzsignfolderExternalid, true
}

// HasSEzsignfolderExternalid returns a boolean if a field has been set.
func (o *EzsignfolderRequestV3) HasSEzsignfolderExternalid() bool {
	if o != nil && !IsNil(o.SEzsignfolderExternalid) {
		return true
	}

	return false
}

// SetSEzsignfolderExternalid gets a reference to the given string and assigns it to the SEzsignfolderExternalid field.
func (o *EzsignfolderRequestV3) SetSEzsignfolderExternalid(v string) {
	o.SEzsignfolderExternalid = &v
}

func (o EzsignfolderRequestV3) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignfolderRequestV3) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PkiEzsignfolderID) {
		toSerialize["pkiEzsignfolderID"] = o.PkiEzsignfolderID
	}
	toSerialize["fkiEzsignfoldertypeID"] = o.FkiEzsignfoldertypeID
	if !IsNil(o.FkiTimezoneID) {
		toSerialize["fkiTimezoneID"] = o.FkiTimezoneID
	}
	if !IsNil(o.FkiEzsigntsarequirementID) {
		toSerialize["fkiEzsigntsarequirementID"] = o.FkiEzsigntsarequirementID
	}
	if !IsNil(o.EEzsignfolderDocumentdependency) {
		toSerialize["eEzsignfolderDocumentdependency"] = o.EEzsignfolderDocumentdependency
	}
	toSerialize["sEzsignfolderDescription"] = o.SEzsignfolderDescription
	if !IsNil(o.TEzsignfolderNote) {
		toSerialize["tEzsignfolderNote"] = o.TEzsignfolderNote
	}
	if !IsNil(o.TEzsignfolderMessage) {
		toSerialize["tEzsignfolderMessage"] = o.TEzsignfolderMessage
	}
	toSerialize["iEzsignfolderSendreminderfirstdays"] = o.IEzsignfolderSendreminderfirstdays
	toSerialize["iEzsignfolderSendreminderotherdays"] = o.IEzsignfolderSendreminderotherdays
	if !IsNil(o.SEzsignfolderExternalid) {
		toSerialize["sEzsignfolderExternalid"] = o.SEzsignfolderExternalid
	}
	return toSerialize, nil
}

func (o *EzsignfolderRequestV3) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fkiEzsignfoldertypeID",
		"sEzsignfolderDescription",
		"iEzsignfolderSendreminderfirstdays",
		"iEzsignfolderSendreminderotherdays",
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

	varEzsignfolderRequestV3 := _EzsignfolderRequestV3{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsignfolderRequestV3)

	if err != nil {
		return err
	}

	*o = EzsignfolderRequestV3(varEzsignfolderRequestV3)

	return err
}

type NullableEzsignfolderRequestV3 struct {
	value *EzsignfolderRequestV3
	isSet bool
}

func (v NullableEzsignfolderRequestV3) Get() *EzsignfolderRequestV3 {
	return v.value
}

func (v *NullableEzsignfolderRequestV3) Set(val *EzsignfolderRequestV3) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignfolderRequestV3) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignfolderRequestV3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignfolderRequestV3(val *EzsignfolderRequestV3) *NullableEzsignfolderRequestV3 {
	return &NullableEzsignfolderRequestV3{value: val, isSet: true}
}

func (v NullableEzsignfolderRequestV3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignfolderRequestV3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


