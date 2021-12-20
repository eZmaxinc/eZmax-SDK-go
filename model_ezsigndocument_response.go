/*
eZmax API Definition

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.1.4
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
)

// EzsigndocumentResponse An Ezsigndocument Object
type EzsigndocumentResponse struct {
	// The unique ID of the Ezsignfolder
	FkiEzsignfolderID int32 `json:"fkiEzsignfolderID"`
	// The maximum date and time at which the Ezsigndocument can be signed.
	DtEzsigndocumentDuedate string `json:"dtEzsigndocumentDuedate"`
	// The unique ID of the Language.  Valid values:  |Value|Description| |-|-| |1|French| |2|English|
	FkiLanguageID int32 `json:"fkiLanguageID"`
	// The name of the document that will be presented to Ezsignfoldersignerassociations
	SEzsigndocumentName string `json:"sEzsigndocumentName"`
	// The unique ID of the Ezsigndocument
	PkiEzsigndocumentID int32 `json:"pkiEzsigndocumentID"`
	EEzsigndocumentStep FieldEEzsigndocumentStep `json:"eEzsigndocumentStep"`
	// The date and time when the Ezsigndocument was first sent.
	DtEzsigndocumentFirstsend string `json:"dtEzsigndocumentFirstsend"`
	// The date and time when the Ezsigndocument was sent the last time.
	DtEzsigndocumentLastsend string `json:"dtEzsigndocumentLastsend"`
	// The order in which the Ezsigndocument will be presented to the signatory in the Ezsignfolder.
	IEzsigndocumentOrder int32 `json:"iEzsigndocumentOrder"`
	// The number of pages in the Ezsigndocument.
	IEzsigndocumentPagetotal int32 `json:"iEzsigndocumentPagetotal"`
	// The number of signatures that were signed in the document.
	IEzsigndocumentSignaturesigned int32 `json:"iEzsigndocumentSignaturesigned"`
	// The number of total signatures that were requested in the Ezsigndocument.
	IEzsigndocumentSignaturetotal int32 `json:"iEzsigndocumentSignaturetotal"`
	// MD5 Hash of the initial PDF Document before signatures were applied to it.
	SEzsigndocumentMD5initial string `json:"sEzsigndocumentMD5initial"`
	// MD5 Hash of the final PDF Document after all signatures were applied to it.
	SEzsigndocumentMD5signed string `json:"sEzsigndocumentMD5signed"`
	ObjAudit CommonAudit `json:"objAudit"`
}

// NewEzsigndocumentResponse instantiates a new EzsigndocumentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigndocumentResponse(fkiEzsignfolderID int32, dtEzsigndocumentDuedate string, fkiLanguageID int32, sEzsigndocumentName string, pkiEzsigndocumentID int32, eEzsigndocumentStep FieldEEzsigndocumentStep, dtEzsigndocumentFirstsend string, dtEzsigndocumentLastsend string, iEzsigndocumentOrder int32, iEzsigndocumentPagetotal int32, iEzsigndocumentSignaturesigned int32, iEzsigndocumentSignaturetotal int32, sEzsigndocumentMD5initial string, sEzsigndocumentMD5signed string, objAudit CommonAudit) *EzsigndocumentResponse {
	this := EzsigndocumentResponse{}
	this.FkiEzsignfolderID = fkiEzsignfolderID
	this.DtEzsigndocumentDuedate = dtEzsigndocumentDuedate
	this.FkiLanguageID = fkiLanguageID
	this.SEzsigndocumentName = sEzsigndocumentName
	this.PkiEzsigndocumentID = pkiEzsigndocumentID
	this.EEzsigndocumentStep = eEzsigndocumentStep
	this.DtEzsigndocumentFirstsend = dtEzsigndocumentFirstsend
	this.DtEzsigndocumentLastsend = dtEzsigndocumentLastsend
	this.IEzsigndocumentOrder = iEzsigndocumentOrder
	this.IEzsigndocumentPagetotal = iEzsigndocumentPagetotal
	this.IEzsigndocumentSignaturesigned = iEzsigndocumentSignaturesigned
	this.IEzsigndocumentSignaturetotal = iEzsigndocumentSignaturetotal
	this.SEzsigndocumentMD5initial = sEzsigndocumentMD5initial
	this.SEzsigndocumentMD5signed = sEzsigndocumentMD5signed
	this.ObjAudit = objAudit
	return &this
}

// NewEzsigndocumentResponseWithDefaults instantiates a new EzsigndocumentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigndocumentResponseWithDefaults() *EzsigndocumentResponse {
	this := EzsigndocumentResponse{}
	return &this
}

// GetFkiEzsignfolderID returns the FkiEzsignfolderID field value
func (o *EzsigndocumentResponse) GetFkiEzsignfolderID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzsignfolderID
}

// GetFkiEzsignfolderIDOk returns a tuple with the FkiEzsignfolderID field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentResponse) GetFkiEzsignfolderIDOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FkiEzsignfolderID, true
}

// SetFkiEzsignfolderID sets field value
func (o *EzsigndocumentResponse) SetFkiEzsignfolderID(v int32) {
	o.FkiEzsignfolderID = v
}

// GetDtEzsigndocumentDuedate returns the DtEzsigndocumentDuedate field value
func (o *EzsigndocumentResponse) GetDtEzsigndocumentDuedate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DtEzsigndocumentDuedate
}

// GetDtEzsigndocumentDuedateOk returns a tuple with the DtEzsigndocumentDuedate field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentResponse) GetDtEzsigndocumentDuedateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DtEzsigndocumentDuedate, true
}

// SetDtEzsigndocumentDuedate sets field value
func (o *EzsigndocumentResponse) SetDtEzsigndocumentDuedate(v string) {
	o.DtEzsigndocumentDuedate = v
}

// GetFkiLanguageID returns the FkiLanguageID field value
func (o *EzsigndocumentResponse) GetFkiLanguageID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiLanguageID
}

// GetFkiLanguageIDOk returns a tuple with the FkiLanguageID field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentResponse) GetFkiLanguageIDOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FkiLanguageID, true
}

// SetFkiLanguageID sets field value
func (o *EzsigndocumentResponse) SetFkiLanguageID(v int32) {
	o.FkiLanguageID = v
}

// GetSEzsigndocumentName returns the SEzsigndocumentName field value
func (o *EzsigndocumentResponse) GetSEzsigndocumentName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzsigndocumentName
}

// GetSEzsigndocumentNameOk returns a tuple with the SEzsigndocumentName field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentResponse) GetSEzsigndocumentNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SEzsigndocumentName, true
}

// SetSEzsigndocumentName sets field value
func (o *EzsigndocumentResponse) SetSEzsigndocumentName(v string) {
	o.SEzsigndocumentName = v
}

// GetPkiEzsigndocumentID returns the PkiEzsigndocumentID field value
func (o *EzsigndocumentResponse) GetPkiEzsigndocumentID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiEzsigndocumentID
}

// GetPkiEzsigndocumentIDOk returns a tuple with the PkiEzsigndocumentID field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentResponse) GetPkiEzsigndocumentIDOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PkiEzsigndocumentID, true
}

// SetPkiEzsigndocumentID sets field value
func (o *EzsigndocumentResponse) SetPkiEzsigndocumentID(v int32) {
	o.PkiEzsigndocumentID = v
}

// GetEEzsigndocumentStep returns the EEzsigndocumentStep field value
func (o *EzsigndocumentResponse) GetEEzsigndocumentStep() FieldEEzsigndocumentStep {
	if o == nil {
		var ret FieldEEzsigndocumentStep
		return ret
	}

	return o.EEzsigndocumentStep
}

// GetEEzsigndocumentStepOk returns a tuple with the EEzsigndocumentStep field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentResponse) GetEEzsigndocumentStepOk() (*FieldEEzsigndocumentStep, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EEzsigndocumentStep, true
}

// SetEEzsigndocumentStep sets field value
func (o *EzsigndocumentResponse) SetEEzsigndocumentStep(v FieldEEzsigndocumentStep) {
	o.EEzsigndocumentStep = v
}

// GetDtEzsigndocumentFirstsend returns the DtEzsigndocumentFirstsend field value
func (o *EzsigndocumentResponse) GetDtEzsigndocumentFirstsend() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DtEzsigndocumentFirstsend
}

// GetDtEzsigndocumentFirstsendOk returns a tuple with the DtEzsigndocumentFirstsend field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentResponse) GetDtEzsigndocumentFirstsendOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DtEzsigndocumentFirstsend, true
}

// SetDtEzsigndocumentFirstsend sets field value
func (o *EzsigndocumentResponse) SetDtEzsigndocumentFirstsend(v string) {
	o.DtEzsigndocumentFirstsend = v
}

// GetDtEzsigndocumentLastsend returns the DtEzsigndocumentLastsend field value
func (o *EzsigndocumentResponse) GetDtEzsigndocumentLastsend() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DtEzsigndocumentLastsend
}

// GetDtEzsigndocumentLastsendOk returns a tuple with the DtEzsigndocumentLastsend field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentResponse) GetDtEzsigndocumentLastsendOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DtEzsigndocumentLastsend, true
}

// SetDtEzsigndocumentLastsend sets field value
func (o *EzsigndocumentResponse) SetDtEzsigndocumentLastsend(v string) {
	o.DtEzsigndocumentLastsend = v
}

// GetIEzsigndocumentOrder returns the IEzsigndocumentOrder field value
func (o *EzsigndocumentResponse) GetIEzsigndocumentOrder() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsigndocumentOrder
}

// GetIEzsigndocumentOrderOk returns a tuple with the IEzsigndocumentOrder field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentResponse) GetIEzsigndocumentOrderOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IEzsigndocumentOrder, true
}

// SetIEzsigndocumentOrder sets field value
func (o *EzsigndocumentResponse) SetIEzsigndocumentOrder(v int32) {
	o.IEzsigndocumentOrder = v
}

// GetIEzsigndocumentPagetotal returns the IEzsigndocumentPagetotal field value
func (o *EzsigndocumentResponse) GetIEzsigndocumentPagetotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsigndocumentPagetotal
}

// GetIEzsigndocumentPagetotalOk returns a tuple with the IEzsigndocumentPagetotal field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentResponse) GetIEzsigndocumentPagetotalOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IEzsigndocumentPagetotal, true
}

// SetIEzsigndocumentPagetotal sets field value
func (o *EzsigndocumentResponse) SetIEzsigndocumentPagetotal(v int32) {
	o.IEzsigndocumentPagetotal = v
}

// GetIEzsigndocumentSignaturesigned returns the IEzsigndocumentSignaturesigned field value
func (o *EzsigndocumentResponse) GetIEzsigndocumentSignaturesigned() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsigndocumentSignaturesigned
}

// GetIEzsigndocumentSignaturesignedOk returns a tuple with the IEzsigndocumentSignaturesigned field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentResponse) GetIEzsigndocumentSignaturesignedOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IEzsigndocumentSignaturesigned, true
}

// SetIEzsigndocumentSignaturesigned sets field value
func (o *EzsigndocumentResponse) SetIEzsigndocumentSignaturesigned(v int32) {
	o.IEzsigndocumentSignaturesigned = v
}

// GetIEzsigndocumentSignaturetotal returns the IEzsigndocumentSignaturetotal field value
func (o *EzsigndocumentResponse) GetIEzsigndocumentSignaturetotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsigndocumentSignaturetotal
}

// GetIEzsigndocumentSignaturetotalOk returns a tuple with the IEzsigndocumentSignaturetotal field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentResponse) GetIEzsigndocumentSignaturetotalOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IEzsigndocumentSignaturetotal, true
}

// SetIEzsigndocumentSignaturetotal sets field value
func (o *EzsigndocumentResponse) SetIEzsigndocumentSignaturetotal(v int32) {
	o.IEzsigndocumentSignaturetotal = v
}

// GetSEzsigndocumentMD5initial returns the SEzsigndocumentMD5initial field value
func (o *EzsigndocumentResponse) GetSEzsigndocumentMD5initial() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzsigndocumentMD5initial
}

// GetSEzsigndocumentMD5initialOk returns a tuple with the SEzsigndocumentMD5initial field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentResponse) GetSEzsigndocumentMD5initialOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SEzsigndocumentMD5initial, true
}

// SetSEzsigndocumentMD5initial sets field value
func (o *EzsigndocumentResponse) SetSEzsigndocumentMD5initial(v string) {
	o.SEzsigndocumentMD5initial = v
}

// GetSEzsigndocumentMD5signed returns the SEzsigndocumentMD5signed field value
func (o *EzsigndocumentResponse) GetSEzsigndocumentMD5signed() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzsigndocumentMD5signed
}

// GetSEzsigndocumentMD5signedOk returns a tuple with the SEzsigndocumentMD5signed field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentResponse) GetSEzsigndocumentMD5signedOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SEzsigndocumentMD5signed, true
}

// SetSEzsigndocumentMD5signed sets field value
func (o *EzsigndocumentResponse) SetSEzsigndocumentMD5signed(v string) {
	o.SEzsigndocumentMD5signed = v
}

// GetObjAudit returns the ObjAudit field value
func (o *EzsigndocumentResponse) GetObjAudit() CommonAudit {
	if o == nil {
		var ret CommonAudit
		return ret
	}

	return o.ObjAudit
}

// GetObjAuditOk returns a tuple with the ObjAudit field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentResponse) GetObjAuditOk() (*CommonAudit, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjAudit, true
}

// SetObjAudit sets field value
func (o *EzsigndocumentResponse) SetObjAudit(v CommonAudit) {
	o.ObjAudit = v
}

func (o EzsigndocumentResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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
	if true {
		toSerialize["pkiEzsigndocumentID"] = o.PkiEzsigndocumentID
	}
	if true {
		toSerialize["eEzsigndocumentStep"] = o.EEzsigndocumentStep
	}
	if true {
		toSerialize["dtEzsigndocumentFirstsend"] = o.DtEzsigndocumentFirstsend
	}
	if true {
		toSerialize["dtEzsigndocumentLastsend"] = o.DtEzsigndocumentLastsend
	}
	if true {
		toSerialize["iEzsigndocumentOrder"] = o.IEzsigndocumentOrder
	}
	if true {
		toSerialize["iEzsigndocumentPagetotal"] = o.IEzsigndocumentPagetotal
	}
	if true {
		toSerialize["iEzsigndocumentSignaturesigned"] = o.IEzsigndocumentSignaturesigned
	}
	if true {
		toSerialize["iEzsigndocumentSignaturetotal"] = o.IEzsigndocumentSignaturetotal
	}
	if true {
		toSerialize["sEzsigndocumentMD5initial"] = o.SEzsigndocumentMD5initial
	}
	if true {
		toSerialize["sEzsigndocumentMD5signed"] = o.SEzsigndocumentMD5signed
	}
	if true {
		toSerialize["objAudit"] = o.ObjAudit
	}
	return json.Marshal(toSerialize)
}

type NullableEzsigndocumentResponse struct {
	value *EzsigndocumentResponse
	isSet bool
}

func (v NullableEzsigndocumentResponse) Get() *EzsigndocumentResponse {
	return v.value
}

func (v *NullableEzsigndocumentResponse) Set(val *EzsigndocumentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigndocumentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigndocumentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigndocumentResponse(val *EzsigndocumentResponse) *NullableEzsigndocumentResponse {
	return &NullableEzsigndocumentResponse{value: val, isSet: true}
}

func (v NullableEzsigndocumentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigndocumentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


