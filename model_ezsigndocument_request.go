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

// checks if the EzsigndocumentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigndocumentRequest{}

// EzsigndocumentRequest An Ezsigndocument Object
type EzsigndocumentRequest struct {
	// The unique ID of the Ezsigndocument
	PkiEzsigndocumentID *int32 `json:"pkiEzsigndocumentID,omitempty"`
	// The unique ID of the Ezsignfolder
	FkiEzsignfolderID int32 `json:"fkiEzsignfolderID"`
	// The unique ID of the Ezsigntemplate
	FkiEzsigntemplateID *int32 `json:"fkiEzsigntemplateID,omitempty"`
	// The unique ID of the Ezsignfoldersignerassociation
	FkiEzsignfoldersignerassociationID *int32 `json:"fkiEzsignfoldersignerassociationID,omitempty"`
	// The unique ID of the Ezsignimportdocument
	FkiEzsignimportdocumentID *int32 `json:"fkiEzsignimportdocumentID,omitempty"`
	// The unique ID of the Language.  Valid values:  |Value|Description| |-|-| |1|French| |2|English|
	FkiLanguageID int32 `json:"fkiLanguageID"`
	// Indicates where to look for the document binary content.
	EEzsigndocumentSource string `json:"eEzsigndocumentSource"`
	// Indicates the format of the document.
	EEzsigndocumentFormat *string `json:"eEzsigndocumentFormat,omitempty"`
	// The Base64 encoded binary content of the document.  This field is Required when eEzsigndocumentSource = Base64.
	SEzsigndocumentBase64 *string `json:"sEzsigndocumentBase64,omitempty"`
	// The url where the document content resides.  This field is Required when eEzsigndocumentSource = Url.
	SEzsigndocumentUrl *string `json:"sEzsigndocumentUrl,omitempty"`
	// Try to repair the document or flatten it if it cannot be used for electronic signature. 
	BEzsigndocumentForcerepair *bool `json:"bEzsigndocumentForcerepair,omitempty"`
	// If the source document is password protected, the password to open/modify it.
	SEzsigndocumentPassword *string `json:"sEzsigndocumentPassword,omitempty"`
	// If the document contains an existing PDF form this property must be set.  **Keep** leaves the form as-is in the document.  **Convert** removes the form and convert all the existing fields to Ezsignformfieldgroups and assign them to the specified **fkiEzsignfoldersignerassociationID**  **Discard** removes the form from the document.  **Flatten** prints the form values in the document.
	EEzsigndocumentForm *string `json:"eEzsigndocumentForm,omitempty"`
	// The maximum date and time at which the Ezsigndocument can be signed.
	DtEzsigndocumentDuedate string `json:"dtEzsigndocumentDuedate"`
	// The name of the document that will be presented to Ezsignfoldersignerassociations
	SEzsigndocumentName string `json:"sEzsigndocumentName"`
	// This field can be used to store an External ID from the client's system.  Anything can be stored in this field, it will never be evaluated by the eZmax system and will be returned AS-IS.  To store multiple values, consider using a JSON formatted structure, a URL encoded string, a CSV or any other custom format. 
	SEzsigndocumentExternalid *string `json:"sEzsigndocumentExternalid,omitempty" validate:"regexp=^.{0,128}$"`
}

type _EzsigndocumentRequest EzsigndocumentRequest

// NewEzsigndocumentRequest instantiates a new EzsigndocumentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigndocumentRequest(fkiEzsignfolderID int32, fkiLanguageID int32, eEzsigndocumentSource string, dtEzsigndocumentDuedate string, sEzsigndocumentName string) *EzsigndocumentRequest {
	this := EzsigndocumentRequest{}
	this.FkiEzsignfolderID = fkiEzsignfolderID
	this.FkiLanguageID = fkiLanguageID
	this.EEzsigndocumentSource = eEzsigndocumentSource
	var bEzsigndocumentForcerepair bool = true
	this.BEzsigndocumentForcerepair = &bEzsigndocumentForcerepair
	this.DtEzsigndocumentDuedate = dtEzsigndocumentDuedate
	this.SEzsigndocumentName = sEzsigndocumentName
	return &this
}

// NewEzsigndocumentRequestWithDefaults instantiates a new EzsigndocumentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigndocumentRequestWithDefaults() *EzsigndocumentRequest {
	this := EzsigndocumentRequest{}
	var bEzsigndocumentForcerepair bool = true
	this.BEzsigndocumentForcerepair = &bEzsigndocumentForcerepair
	return &this
}

// GetPkiEzsigndocumentID returns the PkiEzsigndocumentID field value if set, zero value otherwise.
func (o *EzsigndocumentRequest) GetPkiEzsigndocumentID() int32 {
	if o == nil || IsNil(o.PkiEzsigndocumentID) {
		var ret int32
		return ret
	}
	return *o.PkiEzsigndocumentID
}

// GetPkiEzsigndocumentIDOk returns a tuple with the PkiEzsigndocumentID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigndocumentRequest) GetPkiEzsigndocumentIDOk() (*int32, bool) {
	if o == nil || IsNil(o.PkiEzsigndocumentID) {
		return nil, false
	}
	return o.PkiEzsigndocumentID, true
}

// HasPkiEzsigndocumentID returns a boolean if a field has been set.
func (o *EzsigndocumentRequest) HasPkiEzsigndocumentID() bool {
	if o != nil && !IsNil(o.PkiEzsigndocumentID) {
		return true
	}

	return false
}

// SetPkiEzsigndocumentID gets a reference to the given int32 and assigns it to the PkiEzsigndocumentID field.
func (o *EzsigndocumentRequest) SetPkiEzsigndocumentID(v int32) {
	o.PkiEzsigndocumentID = &v
}

// GetFkiEzsignfolderID returns the FkiEzsignfolderID field value
func (o *EzsigndocumentRequest) GetFkiEzsignfolderID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzsignfolderID
}

// GetFkiEzsignfolderIDOk returns a tuple with the FkiEzsignfolderID field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentRequest) GetFkiEzsignfolderIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiEzsignfolderID, true
}

// SetFkiEzsignfolderID sets field value
func (o *EzsigndocumentRequest) SetFkiEzsignfolderID(v int32) {
	o.FkiEzsignfolderID = v
}

// GetFkiEzsigntemplateID returns the FkiEzsigntemplateID field value if set, zero value otherwise.
func (o *EzsigndocumentRequest) GetFkiEzsigntemplateID() int32 {
	if o == nil || IsNil(o.FkiEzsigntemplateID) {
		var ret int32
		return ret
	}
	return *o.FkiEzsigntemplateID
}

// GetFkiEzsigntemplateIDOk returns a tuple with the FkiEzsigntemplateID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigndocumentRequest) GetFkiEzsigntemplateIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiEzsigntemplateID) {
		return nil, false
	}
	return o.FkiEzsigntemplateID, true
}

// HasFkiEzsigntemplateID returns a boolean if a field has been set.
func (o *EzsigndocumentRequest) HasFkiEzsigntemplateID() bool {
	if o != nil && !IsNil(o.FkiEzsigntemplateID) {
		return true
	}

	return false
}

// SetFkiEzsigntemplateID gets a reference to the given int32 and assigns it to the FkiEzsigntemplateID field.
func (o *EzsigndocumentRequest) SetFkiEzsigntemplateID(v int32) {
	o.FkiEzsigntemplateID = &v
}

// GetFkiEzsignfoldersignerassociationID returns the FkiEzsignfoldersignerassociationID field value if set, zero value otherwise.
func (o *EzsigndocumentRequest) GetFkiEzsignfoldersignerassociationID() int32 {
	if o == nil || IsNil(o.FkiEzsignfoldersignerassociationID) {
		var ret int32
		return ret
	}
	return *o.FkiEzsignfoldersignerassociationID
}

// GetFkiEzsignfoldersignerassociationIDOk returns a tuple with the FkiEzsignfoldersignerassociationID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigndocumentRequest) GetFkiEzsignfoldersignerassociationIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiEzsignfoldersignerassociationID) {
		return nil, false
	}
	return o.FkiEzsignfoldersignerassociationID, true
}

// HasFkiEzsignfoldersignerassociationID returns a boolean if a field has been set.
func (o *EzsigndocumentRequest) HasFkiEzsignfoldersignerassociationID() bool {
	if o != nil && !IsNil(o.FkiEzsignfoldersignerassociationID) {
		return true
	}

	return false
}

// SetFkiEzsignfoldersignerassociationID gets a reference to the given int32 and assigns it to the FkiEzsignfoldersignerassociationID field.
func (o *EzsigndocumentRequest) SetFkiEzsignfoldersignerassociationID(v int32) {
	o.FkiEzsignfoldersignerassociationID = &v
}

// GetFkiEzsignimportdocumentID returns the FkiEzsignimportdocumentID field value if set, zero value otherwise.
func (o *EzsigndocumentRequest) GetFkiEzsignimportdocumentID() int32 {
	if o == nil || IsNil(o.FkiEzsignimportdocumentID) {
		var ret int32
		return ret
	}
	return *o.FkiEzsignimportdocumentID
}

// GetFkiEzsignimportdocumentIDOk returns a tuple with the FkiEzsignimportdocumentID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigndocumentRequest) GetFkiEzsignimportdocumentIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiEzsignimportdocumentID) {
		return nil, false
	}
	return o.FkiEzsignimportdocumentID, true
}

// HasFkiEzsignimportdocumentID returns a boolean if a field has been set.
func (o *EzsigndocumentRequest) HasFkiEzsignimportdocumentID() bool {
	if o != nil && !IsNil(o.FkiEzsignimportdocumentID) {
		return true
	}

	return false
}

// SetFkiEzsignimportdocumentID gets a reference to the given int32 and assigns it to the FkiEzsignimportdocumentID field.
func (o *EzsigndocumentRequest) SetFkiEzsignimportdocumentID(v int32) {
	o.FkiEzsignimportdocumentID = &v
}

// GetFkiLanguageID returns the FkiLanguageID field value
func (o *EzsigndocumentRequest) GetFkiLanguageID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiLanguageID
}

// GetFkiLanguageIDOk returns a tuple with the FkiLanguageID field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentRequest) GetFkiLanguageIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiLanguageID, true
}

// SetFkiLanguageID sets field value
func (o *EzsigndocumentRequest) SetFkiLanguageID(v int32) {
	o.FkiLanguageID = v
}

// GetEEzsigndocumentSource returns the EEzsigndocumentSource field value
func (o *EzsigndocumentRequest) GetEEzsigndocumentSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EEzsigndocumentSource
}

// GetEEzsigndocumentSourceOk returns a tuple with the EEzsigndocumentSource field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentRequest) GetEEzsigndocumentSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EEzsigndocumentSource, true
}

// SetEEzsigndocumentSource sets field value
func (o *EzsigndocumentRequest) SetEEzsigndocumentSource(v string) {
	o.EEzsigndocumentSource = v
}

// GetEEzsigndocumentFormat returns the EEzsigndocumentFormat field value if set, zero value otherwise.
func (o *EzsigndocumentRequest) GetEEzsigndocumentFormat() string {
	if o == nil || IsNil(o.EEzsigndocumentFormat) {
		var ret string
		return ret
	}
	return *o.EEzsigndocumentFormat
}

// GetEEzsigndocumentFormatOk returns a tuple with the EEzsigndocumentFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigndocumentRequest) GetEEzsigndocumentFormatOk() (*string, bool) {
	if o == nil || IsNil(o.EEzsigndocumentFormat) {
		return nil, false
	}
	return o.EEzsigndocumentFormat, true
}

// HasEEzsigndocumentFormat returns a boolean if a field has been set.
func (o *EzsigndocumentRequest) HasEEzsigndocumentFormat() bool {
	if o != nil && !IsNil(o.EEzsigndocumentFormat) {
		return true
	}

	return false
}

// SetEEzsigndocumentFormat gets a reference to the given string and assigns it to the EEzsigndocumentFormat field.
func (o *EzsigndocumentRequest) SetEEzsigndocumentFormat(v string) {
	o.EEzsigndocumentFormat = &v
}

// GetSEzsigndocumentBase64 returns the SEzsigndocumentBase64 field value if set, zero value otherwise.
func (o *EzsigndocumentRequest) GetSEzsigndocumentBase64() string {
	if o == nil || IsNil(o.SEzsigndocumentBase64) {
		var ret string
		return ret
	}
	return *o.SEzsigndocumentBase64
}

// GetSEzsigndocumentBase64Ok returns a tuple with the SEzsigndocumentBase64 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigndocumentRequest) GetSEzsigndocumentBase64Ok() (*string, bool) {
	if o == nil || IsNil(o.SEzsigndocumentBase64) {
		return nil, false
	}
	return o.SEzsigndocumentBase64, true
}

// HasSEzsigndocumentBase64 returns a boolean if a field has been set.
func (o *EzsigndocumentRequest) HasSEzsigndocumentBase64() bool {
	if o != nil && !IsNil(o.SEzsigndocumentBase64) {
		return true
	}

	return false
}

// SetSEzsigndocumentBase64 gets a reference to the given string and assigns it to the SEzsigndocumentBase64 field.
func (o *EzsigndocumentRequest) SetSEzsigndocumentBase64(v string) {
	o.SEzsigndocumentBase64 = &v
}

// GetSEzsigndocumentUrl returns the SEzsigndocumentUrl field value if set, zero value otherwise.
func (o *EzsigndocumentRequest) GetSEzsigndocumentUrl() string {
	if o == nil || IsNil(o.SEzsigndocumentUrl) {
		var ret string
		return ret
	}
	return *o.SEzsigndocumentUrl
}

// GetSEzsigndocumentUrlOk returns a tuple with the SEzsigndocumentUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigndocumentRequest) GetSEzsigndocumentUrlOk() (*string, bool) {
	if o == nil || IsNil(o.SEzsigndocumentUrl) {
		return nil, false
	}
	return o.SEzsigndocumentUrl, true
}

// HasSEzsigndocumentUrl returns a boolean if a field has been set.
func (o *EzsigndocumentRequest) HasSEzsigndocumentUrl() bool {
	if o != nil && !IsNil(o.SEzsigndocumentUrl) {
		return true
	}

	return false
}

// SetSEzsigndocumentUrl gets a reference to the given string and assigns it to the SEzsigndocumentUrl field.
func (o *EzsigndocumentRequest) SetSEzsigndocumentUrl(v string) {
	o.SEzsigndocumentUrl = &v
}

// GetBEzsigndocumentForcerepair returns the BEzsigndocumentForcerepair field value if set, zero value otherwise.
func (o *EzsigndocumentRequest) GetBEzsigndocumentForcerepair() bool {
	if o == nil || IsNil(o.BEzsigndocumentForcerepair) {
		var ret bool
		return ret
	}
	return *o.BEzsigndocumentForcerepair
}

// GetBEzsigndocumentForcerepairOk returns a tuple with the BEzsigndocumentForcerepair field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigndocumentRequest) GetBEzsigndocumentForcerepairOk() (*bool, bool) {
	if o == nil || IsNil(o.BEzsigndocumentForcerepair) {
		return nil, false
	}
	return o.BEzsigndocumentForcerepair, true
}

// HasBEzsigndocumentForcerepair returns a boolean if a field has been set.
func (o *EzsigndocumentRequest) HasBEzsigndocumentForcerepair() bool {
	if o != nil && !IsNil(o.BEzsigndocumentForcerepair) {
		return true
	}

	return false
}

// SetBEzsigndocumentForcerepair gets a reference to the given bool and assigns it to the BEzsigndocumentForcerepair field.
func (o *EzsigndocumentRequest) SetBEzsigndocumentForcerepair(v bool) {
	o.BEzsigndocumentForcerepair = &v
}

// GetSEzsigndocumentPassword returns the SEzsigndocumentPassword field value if set, zero value otherwise.
func (o *EzsigndocumentRequest) GetSEzsigndocumentPassword() string {
	if o == nil || IsNil(o.SEzsigndocumentPassword) {
		var ret string
		return ret
	}
	return *o.SEzsigndocumentPassword
}

// GetSEzsigndocumentPasswordOk returns a tuple with the SEzsigndocumentPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigndocumentRequest) GetSEzsigndocumentPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.SEzsigndocumentPassword) {
		return nil, false
	}
	return o.SEzsigndocumentPassword, true
}

// HasSEzsigndocumentPassword returns a boolean if a field has been set.
func (o *EzsigndocumentRequest) HasSEzsigndocumentPassword() bool {
	if o != nil && !IsNil(o.SEzsigndocumentPassword) {
		return true
	}

	return false
}

// SetSEzsigndocumentPassword gets a reference to the given string and assigns it to the SEzsigndocumentPassword field.
func (o *EzsigndocumentRequest) SetSEzsigndocumentPassword(v string) {
	o.SEzsigndocumentPassword = &v
}

// GetEEzsigndocumentForm returns the EEzsigndocumentForm field value if set, zero value otherwise.
func (o *EzsigndocumentRequest) GetEEzsigndocumentForm() string {
	if o == nil || IsNil(o.EEzsigndocumentForm) {
		var ret string
		return ret
	}
	return *o.EEzsigndocumentForm
}

// GetEEzsigndocumentFormOk returns a tuple with the EEzsigndocumentForm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigndocumentRequest) GetEEzsigndocumentFormOk() (*string, bool) {
	if o == nil || IsNil(o.EEzsigndocumentForm) {
		return nil, false
	}
	return o.EEzsigndocumentForm, true
}

// HasEEzsigndocumentForm returns a boolean if a field has been set.
func (o *EzsigndocumentRequest) HasEEzsigndocumentForm() bool {
	if o != nil && !IsNil(o.EEzsigndocumentForm) {
		return true
	}

	return false
}

// SetEEzsigndocumentForm gets a reference to the given string and assigns it to the EEzsigndocumentForm field.
func (o *EzsigndocumentRequest) SetEEzsigndocumentForm(v string) {
	o.EEzsigndocumentForm = &v
}

// GetDtEzsigndocumentDuedate returns the DtEzsigndocumentDuedate field value
func (o *EzsigndocumentRequest) GetDtEzsigndocumentDuedate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DtEzsigndocumentDuedate
}

// GetDtEzsigndocumentDuedateOk returns a tuple with the DtEzsigndocumentDuedate field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentRequest) GetDtEzsigndocumentDuedateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DtEzsigndocumentDuedate, true
}

// SetDtEzsigndocumentDuedate sets field value
func (o *EzsigndocumentRequest) SetDtEzsigndocumentDuedate(v string) {
	o.DtEzsigndocumentDuedate = v
}

// GetSEzsigndocumentName returns the SEzsigndocumentName field value
func (o *EzsigndocumentRequest) GetSEzsigndocumentName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzsigndocumentName
}

// GetSEzsigndocumentNameOk returns a tuple with the SEzsigndocumentName field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentRequest) GetSEzsigndocumentNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SEzsigndocumentName, true
}

// SetSEzsigndocumentName sets field value
func (o *EzsigndocumentRequest) SetSEzsigndocumentName(v string) {
	o.SEzsigndocumentName = v
}

// GetSEzsigndocumentExternalid returns the SEzsigndocumentExternalid field value if set, zero value otherwise.
func (o *EzsigndocumentRequest) GetSEzsigndocumentExternalid() string {
	if o == nil || IsNil(o.SEzsigndocumentExternalid) {
		var ret string
		return ret
	}
	return *o.SEzsigndocumentExternalid
}

// GetSEzsigndocumentExternalidOk returns a tuple with the SEzsigndocumentExternalid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigndocumentRequest) GetSEzsigndocumentExternalidOk() (*string, bool) {
	if o == nil || IsNil(o.SEzsigndocumentExternalid) {
		return nil, false
	}
	return o.SEzsigndocumentExternalid, true
}

// HasSEzsigndocumentExternalid returns a boolean if a field has been set.
func (o *EzsigndocumentRequest) HasSEzsigndocumentExternalid() bool {
	if o != nil && !IsNil(o.SEzsigndocumentExternalid) {
		return true
	}

	return false
}

// SetSEzsigndocumentExternalid gets a reference to the given string and assigns it to the SEzsigndocumentExternalid field.
func (o *EzsigndocumentRequest) SetSEzsigndocumentExternalid(v string) {
	o.SEzsigndocumentExternalid = &v
}

func (o EzsigndocumentRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigndocumentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PkiEzsigndocumentID) {
		toSerialize["pkiEzsigndocumentID"] = o.PkiEzsigndocumentID
	}
	toSerialize["fkiEzsignfolderID"] = o.FkiEzsignfolderID
	if !IsNil(o.FkiEzsigntemplateID) {
		toSerialize["fkiEzsigntemplateID"] = o.FkiEzsigntemplateID
	}
	if !IsNil(o.FkiEzsignfoldersignerassociationID) {
		toSerialize["fkiEzsignfoldersignerassociationID"] = o.FkiEzsignfoldersignerassociationID
	}
	if !IsNil(o.FkiEzsignimportdocumentID) {
		toSerialize["fkiEzsignimportdocumentID"] = o.FkiEzsignimportdocumentID
	}
	toSerialize["fkiLanguageID"] = o.FkiLanguageID
	toSerialize["eEzsigndocumentSource"] = o.EEzsigndocumentSource
	if !IsNil(o.EEzsigndocumentFormat) {
		toSerialize["eEzsigndocumentFormat"] = o.EEzsigndocumentFormat
	}
	if !IsNil(o.SEzsigndocumentBase64) {
		toSerialize["sEzsigndocumentBase64"] = o.SEzsigndocumentBase64
	}
	if !IsNil(o.SEzsigndocumentUrl) {
		toSerialize["sEzsigndocumentUrl"] = o.SEzsigndocumentUrl
	}
	if !IsNil(o.BEzsigndocumentForcerepair) {
		toSerialize["bEzsigndocumentForcerepair"] = o.BEzsigndocumentForcerepair
	}
	if !IsNil(o.SEzsigndocumentPassword) {
		toSerialize["sEzsigndocumentPassword"] = o.SEzsigndocumentPassword
	}
	if !IsNil(o.EEzsigndocumentForm) {
		toSerialize["eEzsigndocumentForm"] = o.EEzsigndocumentForm
	}
	toSerialize["dtEzsigndocumentDuedate"] = o.DtEzsigndocumentDuedate
	toSerialize["sEzsigndocumentName"] = o.SEzsigndocumentName
	if !IsNil(o.SEzsigndocumentExternalid) {
		toSerialize["sEzsigndocumentExternalid"] = o.SEzsigndocumentExternalid
	}
	return toSerialize, nil
}

func (o *EzsigndocumentRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fkiEzsignfolderID",
		"fkiLanguageID",
		"eEzsigndocumentSource",
		"dtEzsigndocumentDuedate",
		"sEzsigndocumentName",
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

	varEzsigndocumentRequest := _EzsigndocumentRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsigndocumentRequest)

	if err != nil {
		return err
	}

	*o = EzsigndocumentRequest(varEzsigndocumentRequest)

	return err
}

type NullableEzsigndocumentRequest struct {
	value *EzsigndocumentRequest
	isSet bool
}

func (v NullableEzsigndocumentRequest) Get() *EzsigndocumentRequest {
	return v.value
}

func (v *NullableEzsigndocumentRequest) Set(val *EzsigndocumentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigndocumentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigndocumentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigndocumentRequest(val *EzsigndocumentRequest) *NullableEzsigndocumentRequest {
	return &NullableEzsigndocumentRequest{value: val, isSet: true}
}

func (v NullableEzsigndocumentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigndocumentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


