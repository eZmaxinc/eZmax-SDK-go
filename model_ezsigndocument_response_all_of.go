/*
 * eZmax API Definition
 *
 * This API expose all the functionnalities for the eZmax and eZsign application.  We provide SDKs for customers. They are generated using OpenAPI codegen, we encourage customers to use them as we also provide samples for them.  You can choose to build your own implementation manually or can use any compatible OpenAPI 3.0 generator like Swagger Codegen, OpenAPI codegen or any commercial generators.  If you need helping understanding how to use this API, don't waste too much time looking for it. Contact support-api@ezmax.ca, we're here to help. We are developpers so we know programmers don't like bad documentation. If you don't find what you need in the documentation, let us know, we'll improve it and put you rapidly up on track.
 *
 * API version: 1.0.27
 * Contact: support-api@ezmax.ca
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxinc/eZmax-SDK-go

import (
	"encoding/json"
)

// EzsigndocumentResponseAllOf struct for EzsigndocumentResponseAllOf
type EzsigndocumentResponseAllOf struct {
	// The unique ID of the Ezsignfolder
	FkiEzsignfolderID int32 `json:"fkiEzsignfolderID"`
	// The maximum date and time at which the document can be signed.
	DtEzsigndocumentDuedate string `json:"dtEzsigndocumentDuedate"`
	// The unique ID of the Language.  Valid values:  |Value|Description| |-|-| |1|French| |2|English|
	FkiLanguageID int32 `json:"fkiLanguageID"`
	// The actual file name that will be used when downloading or attaching to an email.
	SEzsigndocumentFilename string `json:"sEzsigndocumentFilename"`
	// The name of the document that will be presented to Ezsignfoldersignerassociations
	SEzsigndocumentName string `json:"sEzsigndocumentName"`
	// The unique ID of the Ezsigntemplate
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

// NewEzsigndocumentResponseAllOf instantiates a new EzsigndocumentResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigndocumentResponseAllOf(fkiEzsignfolderID int32, dtEzsigndocumentDuedate string, fkiLanguageID int32, sEzsigndocumentFilename string, sEzsigndocumentName string, pkiEzsigndocumentID int32, eEzsigndocumentStep FieldEEzsigndocumentStep, dtEzsigndocumentFirstsend string, dtEzsigndocumentLastsend string, iEzsigndocumentOrder int32, iEzsigndocumentPagetotal int32, iEzsigndocumentSignaturesigned int32, iEzsigndocumentSignaturetotal int32, sEzsigndocumentMD5initial string, sEzsigndocumentMD5signed string, objAudit CommonAudit, ) *EzsigndocumentResponseAllOf {
	this := EzsigndocumentResponseAllOf{}
	this.FkiEzsignfolderID = fkiEzsignfolderID
	this.DtEzsigndocumentDuedate = dtEzsigndocumentDuedate
	this.FkiLanguageID = fkiLanguageID
	this.SEzsigndocumentFilename = sEzsigndocumentFilename
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

// NewEzsigndocumentResponseAllOfWithDefaults instantiates a new EzsigndocumentResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigndocumentResponseAllOfWithDefaults() *EzsigndocumentResponseAllOf {
	this := EzsigndocumentResponseAllOf{}
	return &this
}

// GetFkiEzsignfolderID returns the FkiEzsignfolderID field value
func (o *EzsigndocumentResponseAllOf) GetFkiEzsignfolderID() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.FkiEzsignfolderID
}

// GetFkiEzsignfolderIDOk returns a tuple with the FkiEzsignfolderID field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentResponseAllOf) GetFkiEzsignfolderIDOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FkiEzsignfolderID, true
}

// SetFkiEzsignfolderID sets field value
func (o *EzsigndocumentResponseAllOf) SetFkiEzsignfolderID(v int32) {
	o.FkiEzsignfolderID = v
}

// GetDtEzsigndocumentDuedate returns the DtEzsigndocumentDuedate field value
func (o *EzsigndocumentResponseAllOf) GetDtEzsigndocumentDuedate() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.DtEzsigndocumentDuedate
}

// GetDtEzsigndocumentDuedateOk returns a tuple with the DtEzsigndocumentDuedate field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentResponseAllOf) GetDtEzsigndocumentDuedateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DtEzsigndocumentDuedate, true
}

// SetDtEzsigndocumentDuedate sets field value
func (o *EzsigndocumentResponseAllOf) SetDtEzsigndocumentDuedate(v string) {
	o.DtEzsigndocumentDuedate = v
}

// GetFkiLanguageID returns the FkiLanguageID field value
func (o *EzsigndocumentResponseAllOf) GetFkiLanguageID() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.FkiLanguageID
}

// GetFkiLanguageIDOk returns a tuple with the FkiLanguageID field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentResponseAllOf) GetFkiLanguageIDOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FkiLanguageID, true
}

// SetFkiLanguageID sets field value
func (o *EzsigndocumentResponseAllOf) SetFkiLanguageID(v int32) {
	o.FkiLanguageID = v
}

// GetSEzsigndocumentFilename returns the SEzsigndocumentFilename field value
func (o *EzsigndocumentResponseAllOf) GetSEzsigndocumentFilename() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.SEzsigndocumentFilename
}

// GetSEzsigndocumentFilenameOk returns a tuple with the SEzsigndocumentFilename field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentResponseAllOf) GetSEzsigndocumentFilenameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SEzsigndocumentFilename, true
}

// SetSEzsigndocumentFilename sets field value
func (o *EzsigndocumentResponseAllOf) SetSEzsigndocumentFilename(v string) {
	o.SEzsigndocumentFilename = v
}

// GetSEzsigndocumentName returns the SEzsigndocumentName field value
func (o *EzsigndocumentResponseAllOf) GetSEzsigndocumentName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.SEzsigndocumentName
}

// GetSEzsigndocumentNameOk returns a tuple with the SEzsigndocumentName field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentResponseAllOf) GetSEzsigndocumentNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SEzsigndocumentName, true
}

// SetSEzsigndocumentName sets field value
func (o *EzsigndocumentResponseAllOf) SetSEzsigndocumentName(v string) {
	o.SEzsigndocumentName = v
}

// GetPkiEzsigndocumentID returns the PkiEzsigndocumentID field value
func (o *EzsigndocumentResponseAllOf) GetPkiEzsigndocumentID() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.PkiEzsigndocumentID
}

// GetPkiEzsigndocumentIDOk returns a tuple with the PkiEzsigndocumentID field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentResponseAllOf) GetPkiEzsigndocumentIDOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PkiEzsigndocumentID, true
}

// SetPkiEzsigndocumentID sets field value
func (o *EzsigndocumentResponseAllOf) SetPkiEzsigndocumentID(v int32) {
	o.PkiEzsigndocumentID = v
}

// GetEEzsigndocumentStep returns the EEzsigndocumentStep field value
func (o *EzsigndocumentResponseAllOf) GetEEzsigndocumentStep() FieldEEzsigndocumentStep {
	if o == nil  {
		var ret FieldEEzsigndocumentStep
		return ret
	}

	return o.EEzsigndocumentStep
}

// GetEEzsigndocumentStepOk returns a tuple with the EEzsigndocumentStep field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentResponseAllOf) GetEEzsigndocumentStepOk() (*FieldEEzsigndocumentStep, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EEzsigndocumentStep, true
}

// SetEEzsigndocumentStep sets field value
func (o *EzsigndocumentResponseAllOf) SetEEzsigndocumentStep(v FieldEEzsigndocumentStep) {
	o.EEzsigndocumentStep = v
}

// GetDtEzsigndocumentFirstsend returns the DtEzsigndocumentFirstsend field value
func (o *EzsigndocumentResponseAllOf) GetDtEzsigndocumentFirstsend() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.DtEzsigndocumentFirstsend
}

// GetDtEzsigndocumentFirstsendOk returns a tuple with the DtEzsigndocumentFirstsend field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentResponseAllOf) GetDtEzsigndocumentFirstsendOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DtEzsigndocumentFirstsend, true
}

// SetDtEzsigndocumentFirstsend sets field value
func (o *EzsigndocumentResponseAllOf) SetDtEzsigndocumentFirstsend(v string) {
	o.DtEzsigndocumentFirstsend = v
}

// GetDtEzsigndocumentLastsend returns the DtEzsigndocumentLastsend field value
func (o *EzsigndocumentResponseAllOf) GetDtEzsigndocumentLastsend() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.DtEzsigndocumentLastsend
}

// GetDtEzsigndocumentLastsendOk returns a tuple with the DtEzsigndocumentLastsend field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentResponseAllOf) GetDtEzsigndocumentLastsendOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DtEzsigndocumentLastsend, true
}

// SetDtEzsigndocumentLastsend sets field value
func (o *EzsigndocumentResponseAllOf) SetDtEzsigndocumentLastsend(v string) {
	o.DtEzsigndocumentLastsend = v
}

// GetIEzsigndocumentOrder returns the IEzsigndocumentOrder field value
func (o *EzsigndocumentResponseAllOf) GetIEzsigndocumentOrder() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.IEzsigndocumentOrder
}

// GetIEzsigndocumentOrderOk returns a tuple with the IEzsigndocumentOrder field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentResponseAllOf) GetIEzsigndocumentOrderOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IEzsigndocumentOrder, true
}

// SetIEzsigndocumentOrder sets field value
func (o *EzsigndocumentResponseAllOf) SetIEzsigndocumentOrder(v int32) {
	o.IEzsigndocumentOrder = v
}

// GetIEzsigndocumentPagetotal returns the IEzsigndocumentPagetotal field value
func (o *EzsigndocumentResponseAllOf) GetIEzsigndocumentPagetotal() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.IEzsigndocumentPagetotal
}

// GetIEzsigndocumentPagetotalOk returns a tuple with the IEzsigndocumentPagetotal field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentResponseAllOf) GetIEzsigndocumentPagetotalOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IEzsigndocumentPagetotal, true
}

// SetIEzsigndocumentPagetotal sets field value
func (o *EzsigndocumentResponseAllOf) SetIEzsigndocumentPagetotal(v int32) {
	o.IEzsigndocumentPagetotal = v
}

// GetIEzsigndocumentSignaturesigned returns the IEzsigndocumentSignaturesigned field value
func (o *EzsigndocumentResponseAllOf) GetIEzsigndocumentSignaturesigned() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.IEzsigndocumentSignaturesigned
}

// GetIEzsigndocumentSignaturesignedOk returns a tuple with the IEzsigndocumentSignaturesigned field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentResponseAllOf) GetIEzsigndocumentSignaturesignedOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IEzsigndocumentSignaturesigned, true
}

// SetIEzsigndocumentSignaturesigned sets field value
func (o *EzsigndocumentResponseAllOf) SetIEzsigndocumentSignaturesigned(v int32) {
	o.IEzsigndocumentSignaturesigned = v
}

// GetIEzsigndocumentSignaturetotal returns the IEzsigndocumentSignaturetotal field value
func (o *EzsigndocumentResponseAllOf) GetIEzsigndocumentSignaturetotal() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.IEzsigndocumentSignaturetotal
}

// GetIEzsigndocumentSignaturetotalOk returns a tuple with the IEzsigndocumentSignaturetotal field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentResponseAllOf) GetIEzsigndocumentSignaturetotalOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IEzsigndocumentSignaturetotal, true
}

// SetIEzsigndocumentSignaturetotal sets field value
func (o *EzsigndocumentResponseAllOf) SetIEzsigndocumentSignaturetotal(v int32) {
	o.IEzsigndocumentSignaturetotal = v
}

// GetSEzsigndocumentMD5initial returns the SEzsigndocumentMD5initial field value
func (o *EzsigndocumentResponseAllOf) GetSEzsigndocumentMD5initial() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.SEzsigndocumentMD5initial
}

// GetSEzsigndocumentMD5initialOk returns a tuple with the SEzsigndocumentMD5initial field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentResponseAllOf) GetSEzsigndocumentMD5initialOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SEzsigndocumentMD5initial, true
}

// SetSEzsigndocumentMD5initial sets field value
func (o *EzsigndocumentResponseAllOf) SetSEzsigndocumentMD5initial(v string) {
	o.SEzsigndocumentMD5initial = v
}

// GetSEzsigndocumentMD5signed returns the SEzsigndocumentMD5signed field value
func (o *EzsigndocumentResponseAllOf) GetSEzsigndocumentMD5signed() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.SEzsigndocumentMD5signed
}

// GetSEzsigndocumentMD5signedOk returns a tuple with the SEzsigndocumentMD5signed field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentResponseAllOf) GetSEzsigndocumentMD5signedOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SEzsigndocumentMD5signed, true
}

// SetSEzsigndocumentMD5signed sets field value
func (o *EzsigndocumentResponseAllOf) SetSEzsigndocumentMD5signed(v string) {
	o.SEzsigndocumentMD5signed = v
}

// GetObjAudit returns the ObjAudit field value
func (o *EzsigndocumentResponseAllOf) GetObjAudit() CommonAudit {
	if o == nil  {
		var ret CommonAudit
		return ret
	}

	return o.ObjAudit
}

// GetObjAuditOk returns a tuple with the ObjAudit field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentResponseAllOf) GetObjAuditOk() (*CommonAudit, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjAudit, true
}

// SetObjAudit sets field value
func (o *EzsigndocumentResponseAllOf) SetObjAudit(v CommonAudit) {
	o.ObjAudit = v
}

func (o EzsigndocumentResponseAllOf) MarshalJSON() ([]byte, error) {
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
		toSerialize["sEzsigndocumentFilename"] = o.SEzsigndocumentFilename
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

type NullableEzsigndocumentResponseAllOf struct {
	value *EzsigndocumentResponseAllOf
	isSet bool
}

func (v NullableEzsigndocumentResponseAllOf) Get() *EzsigndocumentResponseAllOf {
	return v.value
}

func (v *NullableEzsigndocumentResponseAllOf) Set(val *EzsigndocumentResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigndocumentResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigndocumentResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigndocumentResponseAllOf(val *EzsigndocumentResponseAllOf) *NullableEzsigndocumentResponseAllOf {
	return &NullableEzsigndocumentResponseAllOf{value: val, isSet: true}
}

func (v NullableEzsigndocumentResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigndocumentResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


