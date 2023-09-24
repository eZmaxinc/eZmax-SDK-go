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
)

// checks if the EzsignbulksendCreateEzsignbulksendtransmissionV1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignbulksendCreateEzsignbulksendtransmissionV1Request{}

// EzsignbulksendCreateEzsignbulksendtransmissionV1Request Request for POST /1/object/ezsignbulksend/{pkiEzsignbulksendID}/createEzsignbulksendtransmission
type EzsignbulksendCreateEzsignbulksendtransmissionV1Request struct {
	// The unique ID of the Userlogintype  Valid values:  |Value|Description|Detail| |-|-|-| |1|**Email Only**|The Ezsignsigner will receive a secure link by email| |2|**Email and phone or SMS**|The Ezsignsigner will receive a secure link by email and will need to authenticate using SMS or Phone call. **Additional fee applies**| |3|**Email and secret question**|The Ezsignsigner will receive a secure link by email and will need to authenticate using a predefined question and answer| |4|**In person only**|The Ezsignsigner will only be able to sign \"In-Person\" and there won't be any authentication. No email will be sent for invitation to sign. Make sure you evaluate the risk of signature denial and at minimum, we recommend you use a handwritten signature type| |5|**In person with phone or SMS**|The Ezsignsigner will only be able to sign \"In-Person\" and will need to authenticate using SMS or Phone call. No email will be sent for invitation to sign. **Additional fee applies**|
	FkiUserlogintypeID int32 `json:"fkiUserlogintypeID"`
	// The unique ID of the Ezsigntsarequirement.  Determine if a Time Stamping Authority should add a timestamp on each of the signature. Valid values:  |Value|Description| |-|-| |1|No. TSA Timestamping will requested. This will make all signatures a lot faster since no round-trip to the TSA server will be required. Timestamping will be made using eZsign server's time.| |2|Best effort. Timestamping from a Time Stamping Authority will be requested but is not mandatory. In the very improbable case it cannot be completed, the timestamping will be made using eZsign server's time. **Additional fee applies**| |3|Mandatory. Timestamping from a Time Stamping Authority will be requested and is mandatory. In the very improbable case it cannot be completed, the signature will fail and the user will be asked to retry. **Additional fee applies**|
	FkiEzsigntsarequirementID *int32 `json:"fkiEzsigntsarequirementID,omitempty"`
	// The description of the Ezsignbulksendtransmission
	SEzsignbulksendtransmissionDescription string `json:"sEzsignbulksendtransmissionDescription"`
	// The maximum date and time at which the Ezsigndocument can be signed.
	DtEzsigndocumentDuedate string `json:"dtEzsigndocumentDuedate"`
	EEzsignfolderSendreminderfrequency FieldEEzsignfolderSendreminderfrequency `json:"eEzsignfolderSendreminderfrequency"`
	// A custom text message that will be added to the email sent.
	TExtraMessage string `json:"tExtraMessage"`
	// The Base64 encoded binary content of the CSV file.
	SCsvBase64 string `json:"sCsvBase64"`
}

// NewEzsignbulksendCreateEzsignbulksendtransmissionV1Request instantiates a new EzsignbulksendCreateEzsignbulksendtransmissionV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignbulksendCreateEzsignbulksendtransmissionV1Request(fkiUserlogintypeID int32, sEzsignbulksendtransmissionDescription string, dtEzsigndocumentDuedate string, eEzsignfolderSendreminderfrequency FieldEEzsignfolderSendreminderfrequency, tExtraMessage string, sCsvBase64 string) *EzsignbulksendCreateEzsignbulksendtransmissionV1Request {
	this := EzsignbulksendCreateEzsignbulksendtransmissionV1Request{}
	this.FkiUserlogintypeID = fkiUserlogintypeID
	this.SEzsignbulksendtransmissionDescription = sEzsignbulksendtransmissionDescription
	this.DtEzsigndocumentDuedate = dtEzsigndocumentDuedate
	this.EEzsignfolderSendreminderfrequency = eEzsignfolderSendreminderfrequency
	this.TExtraMessage = tExtraMessage
	this.SCsvBase64 = sCsvBase64
	return &this
}

// NewEzsignbulksendCreateEzsignbulksendtransmissionV1RequestWithDefaults instantiates a new EzsignbulksendCreateEzsignbulksendtransmissionV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignbulksendCreateEzsignbulksendtransmissionV1RequestWithDefaults() *EzsignbulksendCreateEzsignbulksendtransmissionV1Request {
	this := EzsignbulksendCreateEzsignbulksendtransmissionV1Request{}
	return &this
}

// GetFkiUserlogintypeID returns the FkiUserlogintypeID field value
func (o *EzsignbulksendCreateEzsignbulksendtransmissionV1Request) GetFkiUserlogintypeID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiUserlogintypeID
}

// GetFkiUserlogintypeIDOk returns a tuple with the FkiUserlogintypeID field value
// and a boolean to check if the value has been set.
func (o *EzsignbulksendCreateEzsignbulksendtransmissionV1Request) GetFkiUserlogintypeIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiUserlogintypeID, true
}

// SetFkiUserlogintypeID sets field value
func (o *EzsignbulksendCreateEzsignbulksendtransmissionV1Request) SetFkiUserlogintypeID(v int32) {
	o.FkiUserlogintypeID = v
}

// GetFkiEzsigntsarequirementID returns the FkiEzsigntsarequirementID field value if set, zero value otherwise.
func (o *EzsignbulksendCreateEzsignbulksendtransmissionV1Request) GetFkiEzsigntsarequirementID() int32 {
	if o == nil || IsNil(o.FkiEzsigntsarequirementID) {
		var ret int32
		return ret
	}
	return *o.FkiEzsigntsarequirementID
}

// GetFkiEzsigntsarequirementIDOk returns a tuple with the FkiEzsigntsarequirementID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignbulksendCreateEzsignbulksendtransmissionV1Request) GetFkiEzsigntsarequirementIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiEzsigntsarequirementID) {
		return nil, false
	}
	return o.FkiEzsigntsarequirementID, true
}

// HasFkiEzsigntsarequirementID returns a boolean if a field has been set.
func (o *EzsignbulksendCreateEzsignbulksendtransmissionV1Request) HasFkiEzsigntsarequirementID() bool {
	if o != nil && !IsNil(o.FkiEzsigntsarequirementID) {
		return true
	}

	return false
}

// SetFkiEzsigntsarequirementID gets a reference to the given int32 and assigns it to the FkiEzsigntsarequirementID field.
func (o *EzsignbulksendCreateEzsignbulksendtransmissionV1Request) SetFkiEzsigntsarequirementID(v int32) {
	o.FkiEzsigntsarequirementID = &v
}

// GetSEzsignbulksendtransmissionDescription returns the SEzsignbulksendtransmissionDescription field value
func (o *EzsignbulksendCreateEzsignbulksendtransmissionV1Request) GetSEzsignbulksendtransmissionDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzsignbulksendtransmissionDescription
}

// GetSEzsignbulksendtransmissionDescriptionOk returns a tuple with the SEzsignbulksendtransmissionDescription field value
// and a boolean to check if the value has been set.
func (o *EzsignbulksendCreateEzsignbulksendtransmissionV1Request) GetSEzsignbulksendtransmissionDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SEzsignbulksendtransmissionDescription, true
}

// SetSEzsignbulksendtransmissionDescription sets field value
func (o *EzsignbulksendCreateEzsignbulksendtransmissionV1Request) SetSEzsignbulksendtransmissionDescription(v string) {
	o.SEzsignbulksendtransmissionDescription = v
}

// GetDtEzsigndocumentDuedate returns the DtEzsigndocumentDuedate field value
func (o *EzsignbulksendCreateEzsignbulksendtransmissionV1Request) GetDtEzsigndocumentDuedate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DtEzsigndocumentDuedate
}

// GetDtEzsigndocumentDuedateOk returns a tuple with the DtEzsigndocumentDuedate field value
// and a boolean to check if the value has been set.
func (o *EzsignbulksendCreateEzsignbulksendtransmissionV1Request) GetDtEzsigndocumentDuedateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DtEzsigndocumentDuedate, true
}

// SetDtEzsigndocumentDuedate sets field value
func (o *EzsignbulksendCreateEzsignbulksendtransmissionV1Request) SetDtEzsigndocumentDuedate(v string) {
	o.DtEzsigndocumentDuedate = v
}

// GetEEzsignfolderSendreminderfrequency returns the EEzsignfolderSendreminderfrequency field value
func (o *EzsignbulksendCreateEzsignbulksendtransmissionV1Request) GetEEzsignfolderSendreminderfrequency() FieldEEzsignfolderSendreminderfrequency {
	if o == nil {
		var ret FieldEEzsignfolderSendreminderfrequency
		return ret
	}

	return o.EEzsignfolderSendreminderfrequency
}

// GetEEzsignfolderSendreminderfrequencyOk returns a tuple with the EEzsignfolderSendreminderfrequency field value
// and a boolean to check if the value has been set.
func (o *EzsignbulksendCreateEzsignbulksendtransmissionV1Request) GetEEzsignfolderSendreminderfrequencyOk() (*FieldEEzsignfolderSendreminderfrequency, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EEzsignfolderSendreminderfrequency, true
}

// SetEEzsignfolderSendreminderfrequency sets field value
func (o *EzsignbulksendCreateEzsignbulksendtransmissionV1Request) SetEEzsignfolderSendreminderfrequency(v FieldEEzsignfolderSendreminderfrequency) {
	o.EEzsignfolderSendreminderfrequency = v
}

// GetTExtraMessage returns the TExtraMessage field value
func (o *EzsignbulksendCreateEzsignbulksendtransmissionV1Request) GetTExtraMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TExtraMessage
}

// GetTExtraMessageOk returns a tuple with the TExtraMessage field value
// and a boolean to check if the value has been set.
func (o *EzsignbulksendCreateEzsignbulksendtransmissionV1Request) GetTExtraMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TExtraMessage, true
}

// SetTExtraMessage sets field value
func (o *EzsignbulksendCreateEzsignbulksendtransmissionV1Request) SetTExtraMessage(v string) {
	o.TExtraMessage = v
}

// GetSCsvBase64 returns the SCsvBase64 field value
func (o *EzsignbulksendCreateEzsignbulksendtransmissionV1Request) GetSCsvBase64() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SCsvBase64
}

// GetSCsvBase64Ok returns a tuple with the SCsvBase64 field value
// and a boolean to check if the value has been set.
func (o *EzsignbulksendCreateEzsignbulksendtransmissionV1Request) GetSCsvBase64Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SCsvBase64, true
}

// SetSCsvBase64 sets field value
func (o *EzsignbulksendCreateEzsignbulksendtransmissionV1Request) SetSCsvBase64(v string) {
	o.SCsvBase64 = v
}

func (o EzsignbulksendCreateEzsignbulksendtransmissionV1Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignbulksendCreateEzsignbulksendtransmissionV1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fkiUserlogintypeID"] = o.FkiUserlogintypeID
	if !IsNil(o.FkiEzsigntsarequirementID) {
		toSerialize["fkiEzsigntsarequirementID"] = o.FkiEzsigntsarequirementID
	}
	toSerialize["sEzsignbulksendtransmissionDescription"] = o.SEzsignbulksendtransmissionDescription
	toSerialize["dtEzsigndocumentDuedate"] = o.DtEzsigndocumentDuedate
	toSerialize["eEzsignfolderSendreminderfrequency"] = o.EEzsignfolderSendreminderfrequency
	toSerialize["tExtraMessage"] = o.TExtraMessage
	toSerialize["sCsvBase64"] = o.SCsvBase64
	return toSerialize, nil
}

type NullableEzsignbulksendCreateEzsignbulksendtransmissionV1Request struct {
	value *EzsignbulksendCreateEzsignbulksendtransmissionV1Request
	isSet bool
}

func (v NullableEzsignbulksendCreateEzsignbulksendtransmissionV1Request) Get() *EzsignbulksendCreateEzsignbulksendtransmissionV1Request {
	return v.value
}

func (v *NullableEzsignbulksendCreateEzsignbulksendtransmissionV1Request) Set(val *EzsignbulksendCreateEzsignbulksendtransmissionV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignbulksendCreateEzsignbulksendtransmissionV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignbulksendCreateEzsignbulksendtransmissionV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignbulksendCreateEzsignbulksendtransmissionV1Request(val *EzsignbulksendCreateEzsignbulksendtransmissionV1Request) *NullableEzsignbulksendCreateEzsignbulksendtransmissionV1Request {
	return &NullableEzsignbulksendCreateEzsignbulksendtransmissionV1Request{value: val, isSet: true}
}

func (v NullableEzsignbulksendCreateEzsignbulksendtransmissionV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignbulksendCreateEzsignbulksendtransmissionV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


