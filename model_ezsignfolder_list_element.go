/*
eZmax API Definition

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.1.3
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
)

// EzsignfolderListElement An Ezsignfolder List Element
type EzsignfolderListElement struct {
	// The unique ID of the Ezsignfolder
	PkiEzsignfolderID int32 `json:"pkiEzsignfolderID"`
	// The unique ID of the Ezsignfoldertype.
	FkiEzsignfoldertypeID int32 `json:"fkiEzsignfoldertypeID"`
	EEzsignfoldertypePrivacylevel FieldEEzsignfoldertypePrivacylevel `json:"eEzsignfoldertypePrivacylevel"`
	// The name of the Ezsignfoldertype in the language of the requester
	SEzsignfoldertypeNameX string `json:"sEzsignfoldertypeNameX"`
	// The description of the Ezsignfolder
	SEzsignfolderDescription string `json:"sEzsignfolderDescription"`
	EEzsignfolderStep FieldEEzsignfolderStep `json:"eEzsignfolderStep"`
	// The date and time at which the object was created
	DtCreatedDate string `json:"dtCreatedDate"`
	DtEzsignfolderSentdate OneOfstringobject `json:"dtEzsignfolderSentdate"`
	// The date at which no more signature will be accepted on the folder
	DtDueDate OneOfstringobject `json:"dtDueDate"`
	// The total number of Ezsigndocument in the folder
	IEzsigndocument int32 `json:"iEzsigndocument"`
	// The total number of Ezsigndocument in the folder that were saved in the edm system
	IEzsigndocumentEdm int32 `json:"iEzsigndocumentEdm"`
	// The total number of signature blocks in all Ezsigndocuments in the folder
	IEzsignsignature int32 `json:"iEzsignsignature"`
	// The total number of already signed signature blocks in all Ezsigndocuments in the folder
	IEzsignsignatureSigned int32 `json:"iEzsignsignatureSigned"`
}

// NewEzsignfolderListElement instantiates a new EzsignfolderListElement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignfolderListElement(pkiEzsignfolderID int32, fkiEzsignfoldertypeID int32, eEzsignfoldertypePrivacylevel FieldEEzsignfoldertypePrivacylevel, sEzsignfoldertypeNameX string, sEzsignfolderDescription string, eEzsignfolderStep FieldEEzsignfolderStep, dtCreatedDate string, dtEzsignfolderSentdate OneOfstringobject, dtDueDate OneOfstringobject, iEzsigndocument int32, iEzsigndocumentEdm int32, iEzsignsignature int32, iEzsignsignatureSigned int32) *EzsignfolderListElement {
	this := EzsignfolderListElement{}
	this.PkiEzsignfolderID = pkiEzsignfolderID
	this.FkiEzsignfoldertypeID = fkiEzsignfoldertypeID
	this.EEzsignfoldertypePrivacylevel = eEzsignfoldertypePrivacylevel
	this.SEzsignfoldertypeNameX = sEzsignfoldertypeNameX
	this.SEzsignfolderDescription = sEzsignfolderDescription
	this.EEzsignfolderStep = eEzsignfolderStep
	this.DtCreatedDate = dtCreatedDate
	this.DtEzsignfolderSentdate = dtEzsignfolderSentdate
	this.DtDueDate = dtDueDate
	this.IEzsigndocument = iEzsigndocument
	this.IEzsigndocumentEdm = iEzsigndocumentEdm
	this.IEzsignsignature = iEzsignsignature
	this.IEzsignsignatureSigned = iEzsignsignatureSigned
	return &this
}

// NewEzsignfolderListElementWithDefaults instantiates a new EzsignfolderListElement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignfolderListElementWithDefaults() *EzsignfolderListElement {
	this := EzsignfolderListElement{}
	return &this
}

// GetPkiEzsignfolderID returns the PkiEzsignfolderID field value
func (o *EzsignfolderListElement) GetPkiEzsignfolderID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiEzsignfolderID
}

// GetPkiEzsignfolderIDOk returns a tuple with the PkiEzsignfolderID field value
// and a boolean to check if the value has been set.
func (o *EzsignfolderListElement) GetPkiEzsignfolderIDOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PkiEzsignfolderID, true
}

// SetPkiEzsignfolderID sets field value
func (o *EzsignfolderListElement) SetPkiEzsignfolderID(v int32) {
	o.PkiEzsignfolderID = v
}

// GetFkiEzsignfoldertypeID returns the FkiEzsignfoldertypeID field value
func (o *EzsignfolderListElement) GetFkiEzsignfoldertypeID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzsignfoldertypeID
}

// GetFkiEzsignfoldertypeIDOk returns a tuple with the FkiEzsignfoldertypeID field value
// and a boolean to check if the value has been set.
func (o *EzsignfolderListElement) GetFkiEzsignfoldertypeIDOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FkiEzsignfoldertypeID, true
}

// SetFkiEzsignfoldertypeID sets field value
func (o *EzsignfolderListElement) SetFkiEzsignfoldertypeID(v int32) {
	o.FkiEzsignfoldertypeID = v
}

// GetEEzsignfoldertypePrivacylevel returns the EEzsignfoldertypePrivacylevel field value
func (o *EzsignfolderListElement) GetEEzsignfoldertypePrivacylevel() FieldEEzsignfoldertypePrivacylevel {
	if o == nil {
		var ret FieldEEzsignfoldertypePrivacylevel
		return ret
	}

	return o.EEzsignfoldertypePrivacylevel
}

// GetEEzsignfoldertypePrivacylevelOk returns a tuple with the EEzsignfoldertypePrivacylevel field value
// and a boolean to check if the value has been set.
func (o *EzsignfolderListElement) GetEEzsignfoldertypePrivacylevelOk() (*FieldEEzsignfoldertypePrivacylevel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EEzsignfoldertypePrivacylevel, true
}

// SetEEzsignfoldertypePrivacylevel sets field value
func (o *EzsignfolderListElement) SetEEzsignfoldertypePrivacylevel(v FieldEEzsignfoldertypePrivacylevel) {
	o.EEzsignfoldertypePrivacylevel = v
}

// GetSEzsignfoldertypeNameX returns the SEzsignfoldertypeNameX field value
func (o *EzsignfolderListElement) GetSEzsignfoldertypeNameX() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzsignfoldertypeNameX
}

// GetSEzsignfoldertypeNameXOk returns a tuple with the SEzsignfoldertypeNameX field value
// and a boolean to check if the value has been set.
func (o *EzsignfolderListElement) GetSEzsignfoldertypeNameXOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SEzsignfoldertypeNameX, true
}

// SetSEzsignfoldertypeNameX sets field value
func (o *EzsignfolderListElement) SetSEzsignfoldertypeNameX(v string) {
	o.SEzsignfoldertypeNameX = v
}

// GetSEzsignfolderDescription returns the SEzsignfolderDescription field value
func (o *EzsignfolderListElement) GetSEzsignfolderDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzsignfolderDescription
}

// GetSEzsignfolderDescriptionOk returns a tuple with the SEzsignfolderDescription field value
// and a boolean to check if the value has been set.
func (o *EzsignfolderListElement) GetSEzsignfolderDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SEzsignfolderDescription, true
}

// SetSEzsignfolderDescription sets field value
func (o *EzsignfolderListElement) SetSEzsignfolderDescription(v string) {
	o.SEzsignfolderDescription = v
}

// GetEEzsignfolderStep returns the EEzsignfolderStep field value
func (o *EzsignfolderListElement) GetEEzsignfolderStep() FieldEEzsignfolderStep {
	if o == nil {
		var ret FieldEEzsignfolderStep
		return ret
	}

	return o.EEzsignfolderStep
}

// GetEEzsignfolderStepOk returns a tuple with the EEzsignfolderStep field value
// and a boolean to check if the value has been set.
func (o *EzsignfolderListElement) GetEEzsignfolderStepOk() (*FieldEEzsignfolderStep, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EEzsignfolderStep, true
}

// SetEEzsignfolderStep sets field value
func (o *EzsignfolderListElement) SetEEzsignfolderStep(v FieldEEzsignfolderStep) {
	o.EEzsignfolderStep = v
}

// GetDtCreatedDate returns the DtCreatedDate field value
func (o *EzsignfolderListElement) GetDtCreatedDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DtCreatedDate
}

// GetDtCreatedDateOk returns a tuple with the DtCreatedDate field value
// and a boolean to check if the value has been set.
func (o *EzsignfolderListElement) GetDtCreatedDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DtCreatedDate, true
}

// SetDtCreatedDate sets field value
func (o *EzsignfolderListElement) SetDtCreatedDate(v string) {
	o.DtCreatedDate = v
}

// GetDtEzsignfolderSentdate returns the DtEzsignfolderSentdate field value
// If the value is explicit nil, the zero value for OneOfstringobject will be returned
func (o *EzsignfolderListElement) GetDtEzsignfolderSentdate() OneOfstringobject {
	if o == nil {
		var ret OneOfstringobject
		return ret
	}

	return o.DtEzsignfolderSentdate
}

// GetDtEzsignfolderSentdateOk returns a tuple with the DtEzsignfolderSentdate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EzsignfolderListElement) GetDtEzsignfolderSentdateOk() (*OneOfstringobject, bool) {
	if o == nil || o.DtEzsignfolderSentdate == nil {
		return nil, false
	}
	return &o.DtEzsignfolderSentdate, true
}

// SetDtEzsignfolderSentdate sets field value
func (o *EzsignfolderListElement) SetDtEzsignfolderSentdate(v OneOfstringobject) {
	o.DtEzsignfolderSentdate = v
}

// GetDtDueDate returns the DtDueDate field value
// If the value is explicit nil, the zero value for OneOfstringobject will be returned
func (o *EzsignfolderListElement) GetDtDueDate() OneOfstringobject {
	if o == nil {
		var ret OneOfstringobject
		return ret
	}

	return o.DtDueDate
}

// GetDtDueDateOk returns a tuple with the DtDueDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EzsignfolderListElement) GetDtDueDateOk() (*OneOfstringobject, bool) {
	if o == nil || o.DtDueDate == nil {
		return nil, false
	}
	return &o.DtDueDate, true
}

// SetDtDueDate sets field value
func (o *EzsignfolderListElement) SetDtDueDate(v OneOfstringobject) {
	o.DtDueDate = v
}

// GetIEzsigndocument returns the IEzsigndocument field value
func (o *EzsignfolderListElement) GetIEzsigndocument() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsigndocument
}

// GetIEzsigndocumentOk returns a tuple with the IEzsigndocument field value
// and a boolean to check if the value has been set.
func (o *EzsignfolderListElement) GetIEzsigndocumentOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IEzsigndocument, true
}

// SetIEzsigndocument sets field value
func (o *EzsignfolderListElement) SetIEzsigndocument(v int32) {
	o.IEzsigndocument = v
}

// GetIEzsigndocumentEdm returns the IEzsigndocumentEdm field value
func (o *EzsignfolderListElement) GetIEzsigndocumentEdm() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsigndocumentEdm
}

// GetIEzsigndocumentEdmOk returns a tuple with the IEzsigndocumentEdm field value
// and a boolean to check if the value has been set.
func (o *EzsignfolderListElement) GetIEzsigndocumentEdmOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IEzsigndocumentEdm, true
}

// SetIEzsigndocumentEdm sets field value
func (o *EzsignfolderListElement) SetIEzsigndocumentEdm(v int32) {
	o.IEzsigndocumentEdm = v
}

// GetIEzsignsignature returns the IEzsignsignature field value
func (o *EzsignfolderListElement) GetIEzsignsignature() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsignsignature
}

// GetIEzsignsignatureOk returns a tuple with the IEzsignsignature field value
// and a boolean to check if the value has been set.
func (o *EzsignfolderListElement) GetIEzsignsignatureOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IEzsignsignature, true
}

// SetIEzsignsignature sets field value
func (o *EzsignfolderListElement) SetIEzsignsignature(v int32) {
	o.IEzsignsignature = v
}

// GetIEzsignsignatureSigned returns the IEzsignsignatureSigned field value
func (o *EzsignfolderListElement) GetIEzsignsignatureSigned() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsignsignatureSigned
}

// GetIEzsignsignatureSignedOk returns a tuple with the IEzsignsignatureSigned field value
// and a boolean to check if the value has been set.
func (o *EzsignfolderListElement) GetIEzsignsignatureSignedOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IEzsignsignatureSigned, true
}

// SetIEzsignsignatureSigned sets field value
func (o *EzsignfolderListElement) SetIEzsignsignatureSigned(v int32) {
	o.IEzsignsignatureSigned = v
}

func (o EzsignfolderListElement) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pkiEzsignfolderID"] = o.PkiEzsignfolderID
	}
	if true {
		toSerialize["fkiEzsignfoldertypeID"] = o.FkiEzsignfoldertypeID
	}
	if true {
		toSerialize["eEzsignfoldertypePrivacylevel"] = o.EEzsignfoldertypePrivacylevel
	}
	if true {
		toSerialize["sEzsignfoldertypeNameX"] = o.SEzsignfoldertypeNameX
	}
	if true {
		toSerialize["sEzsignfolderDescription"] = o.SEzsignfolderDescription
	}
	if true {
		toSerialize["eEzsignfolderStep"] = o.EEzsignfolderStep
	}
	if true {
		toSerialize["dtCreatedDate"] = o.DtCreatedDate
	}
	if o.DtEzsignfolderSentdate != nil {
		toSerialize["dtEzsignfolderSentdate"] = o.DtEzsignfolderSentdate
	}
	if o.DtDueDate != nil {
		toSerialize["dtDueDate"] = o.DtDueDate
	}
	if true {
		toSerialize["iEzsigndocument"] = o.IEzsigndocument
	}
	if true {
		toSerialize["iEzsigndocumentEdm"] = o.IEzsigndocumentEdm
	}
	if true {
		toSerialize["iEzsignsignature"] = o.IEzsignsignature
	}
	if true {
		toSerialize["iEzsignsignatureSigned"] = o.IEzsignsignatureSigned
	}
	return json.Marshal(toSerialize)
}

type NullableEzsignfolderListElement struct {
	value *EzsignfolderListElement
	isSet bool
}

func (v NullableEzsignfolderListElement) Get() *EzsignfolderListElement {
	return v.value
}

func (v *NullableEzsignfolderListElement) Set(val *EzsignfolderListElement) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignfolderListElement) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignfolderListElement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignfolderListElement(val *EzsignfolderListElement) *NullableEzsignfolderListElement {
	return &NullableEzsignfolderListElement{value: val, isSet: true}
}

func (v NullableEzsignfolderListElement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignfolderListElement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


