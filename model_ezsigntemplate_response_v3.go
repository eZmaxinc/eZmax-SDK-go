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

// checks if the EzsigntemplateResponseV3 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigntemplateResponseV3{}

// EzsigntemplateResponseV3 A Ezsigntemplate Object
type EzsigntemplateResponseV3 struct {
	// The unique ID of the Ezsigntemplate
	PkiEzsigntemplateID int32 `json:"pkiEzsigntemplateID"`
	// The unique ID of the Ezsigntemplatedocument
	FkiEzsigntemplatedocumentID *int32 `json:"fkiEzsigntemplatedocumentID,omitempty"`
	// The unique ID of the Ezsignfoldertype.
	FkiEzsignfoldertypeID *int32 `json:"fkiEzsignfoldertypeID,omitempty"`
	// The unique ID of the Language.  Valid values:  |Value|Description| |-|-| |1|French| |2|English|
	FkiLanguageID int32 `json:"fkiLanguageID"`
	// The unique ID of the Ezdoctemplatedocument
	FkiEzdoctemplatedocumentID *int32 `json:"fkiEzdoctemplatedocumentID,omitempty"`
	// The name of the Ezdoctemplatedocument in the language of the requester
	SEzdoctemplatedocumentNameX *string `json:"sEzdoctemplatedocumentNameX,omitempty" validate:"regexp=^.{0,50}$"`
	// The Name of the Language in the language of the requester
	SLanguageNameX string `json:"sLanguageNameX"`
	// The description of the Ezsigntemplate
	SEzsigntemplateDescription string `json:"sEzsigntemplateDescription" validate:"regexp=^.{0,80}$"`
	// The external description of the Ezsigntemplate
	SEzsigntemplateExternaldescription *string `json:"sEzsigntemplateExternaldescription,omitempty" validate:"regexp=^.{0,75}$"`
	// The comment of the Ezsigntemplate
	TEzsigntemplateComment *string `json:"tEzsigntemplateComment,omitempty"`
	EEzsigntemplateRecognition *FieldEEzsigntemplateRecognition `json:"eEzsigntemplateRecognition,omitempty"`
	// The filename regexp of the Ezsigntemplate.
	SEzsigntemplateFilenameregexp *string `json:"sEzsigntemplateFilenameregexp,omitempty" validate:"regexp=^.{1,50}$"`
	// Whether the Ezsigntemplate can be accessed by admin users only (eUserType=Normal)
	BEzsigntemplateAdminonly bool `json:"bEzsigntemplateAdminonly"`
	// The name of the Ezsignfoldertype in the language of the requester
	SEzsignfoldertypeNameX *string `json:"sEzsignfoldertypeNameX,omitempty"`
	ObjAudit CommonAudit `json:"objAudit"`
	// Whether the Ezsigntemplate if allowed to edit or not
	BEzsigntemplateEditallowed bool `json:"bEzsigntemplateEditallowed"`
	EEzsigntemplateType *FieldEEzsigntemplateType `json:"eEzsigntemplateType,omitempty"`
}

type _EzsigntemplateResponseV3 EzsigntemplateResponseV3

// NewEzsigntemplateResponseV3 instantiates a new EzsigntemplateResponseV3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigntemplateResponseV3(pkiEzsigntemplateID int32, fkiLanguageID int32, sLanguageNameX string, sEzsigntemplateDescription string, bEzsigntemplateAdminonly bool, objAudit CommonAudit, bEzsigntemplateEditallowed bool) *EzsigntemplateResponseV3 {
	this := EzsigntemplateResponseV3{}
	this.PkiEzsigntemplateID = pkiEzsigntemplateID
	this.FkiLanguageID = fkiLanguageID
	this.SLanguageNameX = sLanguageNameX
	this.SEzsigntemplateDescription = sEzsigntemplateDescription
	var eEzsigntemplateRecognition FieldEEzsigntemplateRecognition = NO
	this.EEzsigntemplateRecognition = &eEzsigntemplateRecognition
	this.BEzsigntemplateAdminonly = bEzsigntemplateAdminonly
	this.ObjAudit = objAudit
	this.BEzsigntemplateEditallowed = bEzsigntemplateEditallowed
	return &this
}

// NewEzsigntemplateResponseV3WithDefaults instantiates a new EzsigntemplateResponseV3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigntemplateResponseV3WithDefaults() *EzsigntemplateResponseV3 {
	this := EzsigntemplateResponseV3{}
	var eEzsigntemplateRecognition FieldEEzsigntemplateRecognition = NO
	this.EEzsigntemplateRecognition = &eEzsigntemplateRecognition
	return &this
}

// GetPkiEzsigntemplateID returns the PkiEzsigntemplateID field value
func (o *EzsigntemplateResponseV3) GetPkiEzsigntemplateID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiEzsigntemplateID
}

// GetPkiEzsigntemplateIDOk returns a tuple with the PkiEzsigntemplateID field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateResponseV3) GetPkiEzsigntemplateIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiEzsigntemplateID, true
}

// SetPkiEzsigntemplateID sets field value
func (o *EzsigntemplateResponseV3) SetPkiEzsigntemplateID(v int32) {
	o.PkiEzsigntemplateID = v
}

// GetFkiEzsigntemplatedocumentID returns the FkiEzsigntemplatedocumentID field value if set, zero value otherwise.
func (o *EzsigntemplateResponseV3) GetFkiEzsigntemplatedocumentID() int32 {
	if o == nil || IsNil(o.FkiEzsigntemplatedocumentID) {
		var ret int32
		return ret
	}
	return *o.FkiEzsigntemplatedocumentID
}

// GetFkiEzsigntemplatedocumentIDOk returns a tuple with the FkiEzsigntemplatedocumentID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateResponseV3) GetFkiEzsigntemplatedocumentIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiEzsigntemplatedocumentID) {
		return nil, false
	}
	return o.FkiEzsigntemplatedocumentID, true
}

// HasFkiEzsigntemplatedocumentID returns a boolean if a field has been set.
func (o *EzsigntemplateResponseV3) HasFkiEzsigntemplatedocumentID() bool {
	if o != nil && !IsNil(o.FkiEzsigntemplatedocumentID) {
		return true
	}

	return false
}

// SetFkiEzsigntemplatedocumentID gets a reference to the given int32 and assigns it to the FkiEzsigntemplatedocumentID field.
func (o *EzsigntemplateResponseV3) SetFkiEzsigntemplatedocumentID(v int32) {
	o.FkiEzsigntemplatedocumentID = &v
}

// GetFkiEzsignfoldertypeID returns the FkiEzsignfoldertypeID field value if set, zero value otherwise.
func (o *EzsigntemplateResponseV3) GetFkiEzsignfoldertypeID() int32 {
	if o == nil || IsNil(o.FkiEzsignfoldertypeID) {
		var ret int32
		return ret
	}
	return *o.FkiEzsignfoldertypeID
}

// GetFkiEzsignfoldertypeIDOk returns a tuple with the FkiEzsignfoldertypeID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateResponseV3) GetFkiEzsignfoldertypeIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiEzsignfoldertypeID) {
		return nil, false
	}
	return o.FkiEzsignfoldertypeID, true
}

// HasFkiEzsignfoldertypeID returns a boolean if a field has been set.
func (o *EzsigntemplateResponseV3) HasFkiEzsignfoldertypeID() bool {
	if o != nil && !IsNil(o.FkiEzsignfoldertypeID) {
		return true
	}

	return false
}

// SetFkiEzsignfoldertypeID gets a reference to the given int32 and assigns it to the FkiEzsignfoldertypeID field.
func (o *EzsigntemplateResponseV3) SetFkiEzsignfoldertypeID(v int32) {
	o.FkiEzsignfoldertypeID = &v
}

// GetFkiLanguageID returns the FkiLanguageID field value
func (o *EzsigntemplateResponseV3) GetFkiLanguageID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiLanguageID
}

// GetFkiLanguageIDOk returns a tuple with the FkiLanguageID field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateResponseV3) GetFkiLanguageIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiLanguageID, true
}

// SetFkiLanguageID sets field value
func (o *EzsigntemplateResponseV3) SetFkiLanguageID(v int32) {
	o.FkiLanguageID = v
}

// GetFkiEzdoctemplatedocumentID returns the FkiEzdoctemplatedocumentID field value if set, zero value otherwise.
func (o *EzsigntemplateResponseV3) GetFkiEzdoctemplatedocumentID() int32 {
	if o == nil || IsNil(o.FkiEzdoctemplatedocumentID) {
		var ret int32
		return ret
	}
	return *o.FkiEzdoctemplatedocumentID
}

// GetFkiEzdoctemplatedocumentIDOk returns a tuple with the FkiEzdoctemplatedocumentID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateResponseV3) GetFkiEzdoctemplatedocumentIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiEzdoctemplatedocumentID) {
		return nil, false
	}
	return o.FkiEzdoctemplatedocumentID, true
}

// HasFkiEzdoctemplatedocumentID returns a boolean if a field has been set.
func (o *EzsigntemplateResponseV3) HasFkiEzdoctemplatedocumentID() bool {
	if o != nil && !IsNil(o.FkiEzdoctemplatedocumentID) {
		return true
	}

	return false
}

// SetFkiEzdoctemplatedocumentID gets a reference to the given int32 and assigns it to the FkiEzdoctemplatedocumentID field.
func (o *EzsigntemplateResponseV3) SetFkiEzdoctemplatedocumentID(v int32) {
	o.FkiEzdoctemplatedocumentID = &v
}

// GetSEzdoctemplatedocumentNameX returns the SEzdoctemplatedocumentNameX field value if set, zero value otherwise.
func (o *EzsigntemplateResponseV3) GetSEzdoctemplatedocumentNameX() string {
	if o == nil || IsNil(o.SEzdoctemplatedocumentNameX) {
		var ret string
		return ret
	}
	return *o.SEzdoctemplatedocumentNameX
}

// GetSEzdoctemplatedocumentNameXOk returns a tuple with the SEzdoctemplatedocumentNameX field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateResponseV3) GetSEzdoctemplatedocumentNameXOk() (*string, bool) {
	if o == nil || IsNil(o.SEzdoctemplatedocumentNameX) {
		return nil, false
	}
	return o.SEzdoctemplatedocumentNameX, true
}

// HasSEzdoctemplatedocumentNameX returns a boolean if a field has been set.
func (o *EzsigntemplateResponseV3) HasSEzdoctemplatedocumentNameX() bool {
	if o != nil && !IsNil(o.SEzdoctemplatedocumentNameX) {
		return true
	}

	return false
}

// SetSEzdoctemplatedocumentNameX gets a reference to the given string and assigns it to the SEzdoctemplatedocumentNameX field.
func (o *EzsigntemplateResponseV3) SetSEzdoctemplatedocumentNameX(v string) {
	o.SEzdoctemplatedocumentNameX = &v
}

// GetSLanguageNameX returns the SLanguageNameX field value
func (o *EzsigntemplateResponseV3) GetSLanguageNameX() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SLanguageNameX
}

// GetSLanguageNameXOk returns a tuple with the SLanguageNameX field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateResponseV3) GetSLanguageNameXOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SLanguageNameX, true
}

// SetSLanguageNameX sets field value
func (o *EzsigntemplateResponseV3) SetSLanguageNameX(v string) {
	o.SLanguageNameX = v
}

// GetSEzsigntemplateDescription returns the SEzsigntemplateDescription field value
func (o *EzsigntemplateResponseV3) GetSEzsigntemplateDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzsigntemplateDescription
}

// GetSEzsigntemplateDescriptionOk returns a tuple with the SEzsigntemplateDescription field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateResponseV3) GetSEzsigntemplateDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SEzsigntemplateDescription, true
}

// SetSEzsigntemplateDescription sets field value
func (o *EzsigntemplateResponseV3) SetSEzsigntemplateDescription(v string) {
	o.SEzsigntemplateDescription = v
}

// GetSEzsigntemplateExternaldescription returns the SEzsigntemplateExternaldescription field value if set, zero value otherwise.
func (o *EzsigntemplateResponseV3) GetSEzsigntemplateExternaldescription() string {
	if o == nil || IsNil(o.SEzsigntemplateExternaldescription) {
		var ret string
		return ret
	}
	return *o.SEzsigntemplateExternaldescription
}

// GetSEzsigntemplateExternaldescriptionOk returns a tuple with the SEzsigntemplateExternaldescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateResponseV3) GetSEzsigntemplateExternaldescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.SEzsigntemplateExternaldescription) {
		return nil, false
	}
	return o.SEzsigntemplateExternaldescription, true
}

// HasSEzsigntemplateExternaldescription returns a boolean if a field has been set.
func (o *EzsigntemplateResponseV3) HasSEzsigntemplateExternaldescription() bool {
	if o != nil && !IsNil(o.SEzsigntemplateExternaldescription) {
		return true
	}

	return false
}

// SetSEzsigntemplateExternaldescription gets a reference to the given string and assigns it to the SEzsigntemplateExternaldescription field.
func (o *EzsigntemplateResponseV3) SetSEzsigntemplateExternaldescription(v string) {
	o.SEzsigntemplateExternaldescription = &v
}

// GetTEzsigntemplateComment returns the TEzsigntemplateComment field value if set, zero value otherwise.
func (o *EzsigntemplateResponseV3) GetTEzsigntemplateComment() string {
	if o == nil || IsNil(o.TEzsigntemplateComment) {
		var ret string
		return ret
	}
	return *o.TEzsigntemplateComment
}

// GetTEzsigntemplateCommentOk returns a tuple with the TEzsigntemplateComment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateResponseV3) GetTEzsigntemplateCommentOk() (*string, bool) {
	if o == nil || IsNil(o.TEzsigntemplateComment) {
		return nil, false
	}
	return o.TEzsigntemplateComment, true
}

// HasTEzsigntemplateComment returns a boolean if a field has been set.
func (o *EzsigntemplateResponseV3) HasTEzsigntemplateComment() bool {
	if o != nil && !IsNil(o.TEzsigntemplateComment) {
		return true
	}

	return false
}

// SetTEzsigntemplateComment gets a reference to the given string and assigns it to the TEzsigntemplateComment field.
func (o *EzsigntemplateResponseV3) SetTEzsigntemplateComment(v string) {
	o.TEzsigntemplateComment = &v
}

// GetEEzsigntemplateRecognition returns the EEzsigntemplateRecognition field value if set, zero value otherwise.
func (o *EzsigntemplateResponseV3) GetEEzsigntemplateRecognition() FieldEEzsigntemplateRecognition {
	if o == nil || IsNil(o.EEzsigntemplateRecognition) {
		var ret FieldEEzsigntemplateRecognition
		return ret
	}
	return *o.EEzsigntemplateRecognition
}

// GetEEzsigntemplateRecognitionOk returns a tuple with the EEzsigntemplateRecognition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateResponseV3) GetEEzsigntemplateRecognitionOk() (*FieldEEzsigntemplateRecognition, bool) {
	if o == nil || IsNil(o.EEzsigntemplateRecognition) {
		return nil, false
	}
	return o.EEzsigntemplateRecognition, true
}

// HasEEzsigntemplateRecognition returns a boolean if a field has been set.
func (o *EzsigntemplateResponseV3) HasEEzsigntemplateRecognition() bool {
	if o != nil && !IsNil(o.EEzsigntemplateRecognition) {
		return true
	}

	return false
}

// SetEEzsigntemplateRecognition gets a reference to the given FieldEEzsigntemplateRecognition and assigns it to the EEzsigntemplateRecognition field.
func (o *EzsigntemplateResponseV3) SetEEzsigntemplateRecognition(v FieldEEzsigntemplateRecognition) {
	o.EEzsigntemplateRecognition = &v
}

// GetSEzsigntemplateFilenameregexp returns the SEzsigntemplateFilenameregexp field value if set, zero value otherwise.
func (o *EzsigntemplateResponseV3) GetSEzsigntemplateFilenameregexp() string {
	if o == nil || IsNil(o.SEzsigntemplateFilenameregexp) {
		var ret string
		return ret
	}
	return *o.SEzsigntemplateFilenameregexp
}

// GetSEzsigntemplateFilenameregexpOk returns a tuple with the SEzsigntemplateFilenameregexp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateResponseV3) GetSEzsigntemplateFilenameregexpOk() (*string, bool) {
	if o == nil || IsNil(o.SEzsigntemplateFilenameregexp) {
		return nil, false
	}
	return o.SEzsigntemplateFilenameregexp, true
}

// HasSEzsigntemplateFilenameregexp returns a boolean if a field has been set.
func (o *EzsigntemplateResponseV3) HasSEzsigntemplateFilenameregexp() bool {
	if o != nil && !IsNil(o.SEzsigntemplateFilenameregexp) {
		return true
	}

	return false
}

// SetSEzsigntemplateFilenameregexp gets a reference to the given string and assigns it to the SEzsigntemplateFilenameregexp field.
func (o *EzsigntemplateResponseV3) SetSEzsigntemplateFilenameregexp(v string) {
	o.SEzsigntemplateFilenameregexp = &v
}

// GetBEzsigntemplateAdminonly returns the BEzsigntemplateAdminonly field value
func (o *EzsigntemplateResponseV3) GetBEzsigntemplateAdminonly() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BEzsigntemplateAdminonly
}

// GetBEzsigntemplateAdminonlyOk returns a tuple with the BEzsigntemplateAdminonly field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateResponseV3) GetBEzsigntemplateAdminonlyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BEzsigntemplateAdminonly, true
}

// SetBEzsigntemplateAdminonly sets field value
func (o *EzsigntemplateResponseV3) SetBEzsigntemplateAdminonly(v bool) {
	o.BEzsigntemplateAdminonly = v
}

// GetSEzsignfoldertypeNameX returns the SEzsignfoldertypeNameX field value if set, zero value otherwise.
func (o *EzsigntemplateResponseV3) GetSEzsignfoldertypeNameX() string {
	if o == nil || IsNil(o.SEzsignfoldertypeNameX) {
		var ret string
		return ret
	}
	return *o.SEzsignfoldertypeNameX
}

// GetSEzsignfoldertypeNameXOk returns a tuple with the SEzsignfoldertypeNameX field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateResponseV3) GetSEzsignfoldertypeNameXOk() (*string, bool) {
	if o == nil || IsNil(o.SEzsignfoldertypeNameX) {
		return nil, false
	}
	return o.SEzsignfoldertypeNameX, true
}

// HasSEzsignfoldertypeNameX returns a boolean if a field has been set.
func (o *EzsigntemplateResponseV3) HasSEzsignfoldertypeNameX() bool {
	if o != nil && !IsNil(o.SEzsignfoldertypeNameX) {
		return true
	}

	return false
}

// SetSEzsignfoldertypeNameX gets a reference to the given string and assigns it to the SEzsignfoldertypeNameX field.
func (o *EzsigntemplateResponseV3) SetSEzsignfoldertypeNameX(v string) {
	o.SEzsignfoldertypeNameX = &v
}

// GetObjAudit returns the ObjAudit field value
func (o *EzsigntemplateResponseV3) GetObjAudit() CommonAudit {
	if o == nil {
		var ret CommonAudit
		return ret
	}

	return o.ObjAudit
}

// GetObjAuditOk returns a tuple with the ObjAudit field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateResponseV3) GetObjAuditOk() (*CommonAudit, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjAudit, true
}

// SetObjAudit sets field value
func (o *EzsigntemplateResponseV3) SetObjAudit(v CommonAudit) {
	o.ObjAudit = v
}

// GetBEzsigntemplateEditallowed returns the BEzsigntemplateEditallowed field value
func (o *EzsigntemplateResponseV3) GetBEzsigntemplateEditallowed() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BEzsigntemplateEditallowed
}

// GetBEzsigntemplateEditallowedOk returns a tuple with the BEzsigntemplateEditallowed field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateResponseV3) GetBEzsigntemplateEditallowedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BEzsigntemplateEditallowed, true
}

// SetBEzsigntemplateEditallowed sets field value
func (o *EzsigntemplateResponseV3) SetBEzsigntemplateEditallowed(v bool) {
	o.BEzsigntemplateEditallowed = v
}

// GetEEzsigntemplateType returns the EEzsigntemplateType field value if set, zero value otherwise.
func (o *EzsigntemplateResponseV3) GetEEzsigntemplateType() FieldEEzsigntemplateType {
	if o == nil || IsNil(o.EEzsigntemplateType) {
		var ret FieldEEzsigntemplateType
		return ret
	}
	return *o.EEzsigntemplateType
}

// GetEEzsigntemplateTypeOk returns a tuple with the EEzsigntemplateType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateResponseV3) GetEEzsigntemplateTypeOk() (*FieldEEzsigntemplateType, bool) {
	if o == nil || IsNil(o.EEzsigntemplateType) {
		return nil, false
	}
	return o.EEzsigntemplateType, true
}

// HasEEzsigntemplateType returns a boolean if a field has been set.
func (o *EzsigntemplateResponseV3) HasEEzsigntemplateType() bool {
	if o != nil && !IsNil(o.EEzsigntemplateType) {
		return true
	}

	return false
}

// SetEEzsigntemplateType gets a reference to the given FieldEEzsigntemplateType and assigns it to the EEzsigntemplateType field.
func (o *EzsigntemplateResponseV3) SetEEzsigntemplateType(v FieldEEzsigntemplateType) {
	o.EEzsigntemplateType = &v
}

func (o EzsigntemplateResponseV3) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigntemplateResponseV3) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiEzsigntemplateID"] = o.PkiEzsigntemplateID
	if !IsNil(o.FkiEzsigntemplatedocumentID) {
		toSerialize["fkiEzsigntemplatedocumentID"] = o.FkiEzsigntemplatedocumentID
	}
	if !IsNil(o.FkiEzsignfoldertypeID) {
		toSerialize["fkiEzsignfoldertypeID"] = o.FkiEzsignfoldertypeID
	}
	toSerialize["fkiLanguageID"] = o.FkiLanguageID
	if !IsNil(o.FkiEzdoctemplatedocumentID) {
		toSerialize["fkiEzdoctemplatedocumentID"] = o.FkiEzdoctemplatedocumentID
	}
	if !IsNil(o.SEzdoctemplatedocumentNameX) {
		toSerialize["sEzdoctemplatedocumentNameX"] = o.SEzdoctemplatedocumentNameX
	}
	toSerialize["sLanguageNameX"] = o.SLanguageNameX
	toSerialize["sEzsigntemplateDescription"] = o.SEzsigntemplateDescription
	if !IsNil(o.SEzsigntemplateExternaldescription) {
		toSerialize["sEzsigntemplateExternaldescription"] = o.SEzsigntemplateExternaldescription
	}
	if !IsNil(o.TEzsigntemplateComment) {
		toSerialize["tEzsigntemplateComment"] = o.TEzsigntemplateComment
	}
	if !IsNil(o.EEzsigntemplateRecognition) {
		toSerialize["eEzsigntemplateRecognition"] = o.EEzsigntemplateRecognition
	}
	if !IsNil(o.SEzsigntemplateFilenameregexp) {
		toSerialize["sEzsigntemplateFilenameregexp"] = o.SEzsigntemplateFilenameregexp
	}
	toSerialize["bEzsigntemplateAdminonly"] = o.BEzsigntemplateAdminonly
	if !IsNil(o.SEzsignfoldertypeNameX) {
		toSerialize["sEzsignfoldertypeNameX"] = o.SEzsignfoldertypeNameX
	}
	toSerialize["objAudit"] = o.ObjAudit
	toSerialize["bEzsigntemplateEditallowed"] = o.BEzsigntemplateEditallowed
	if !IsNil(o.EEzsigntemplateType) {
		toSerialize["eEzsigntemplateType"] = o.EEzsigntemplateType
	}
	return toSerialize, nil
}

func (o *EzsigntemplateResponseV3) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiEzsigntemplateID",
		"fkiLanguageID",
		"sLanguageNameX",
		"sEzsigntemplateDescription",
		"bEzsigntemplateAdminonly",
		"objAudit",
		"bEzsigntemplateEditallowed",
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

	varEzsigntemplateResponseV3 := _EzsigntemplateResponseV3{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsigntemplateResponseV3)

	if err != nil {
		return err
	}

	*o = EzsigntemplateResponseV3(varEzsigntemplateResponseV3)

	return err
}

type NullableEzsigntemplateResponseV3 struct {
	value *EzsigntemplateResponseV3
	isSet bool
}

func (v NullableEzsigntemplateResponseV3) Get() *EzsigntemplateResponseV3 {
	return v.value
}

func (v *NullableEzsigntemplateResponseV3) Set(val *EzsigntemplateResponseV3) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigntemplateResponseV3) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigntemplateResponseV3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigntemplateResponseV3(val *EzsigntemplateResponseV3) *NullableEzsigntemplateResponseV3 {
	return &NullableEzsigntemplateResponseV3{value: val, isSet: true}
}

func (v NullableEzsigntemplateResponseV3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigntemplateResponseV3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


