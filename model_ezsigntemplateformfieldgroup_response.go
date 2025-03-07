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

// checks if the EzsigntemplateformfieldgroupResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigntemplateformfieldgroupResponse{}

// EzsigntemplateformfieldgroupResponse A Ezsigntemplateformfieldgroup Object
type EzsigntemplateformfieldgroupResponse struct {
	// The unique ID of the Ezsigntemplateformfieldgroup
	PkiEzsigntemplateformfieldgroupID int32 `json:"pkiEzsigntemplateformfieldgroupID"`
	// The unique ID of the Ezsigntemplatedocument
	FkiEzsigntemplatedocumentID int32 `json:"fkiEzsigntemplatedocumentID"`
	EEzsigntemplateformfieldgroupType FieldEEzsigntemplateformfieldgroupType `json:"eEzsigntemplateformfieldgroupType"`
	// Deprecated
	EEzsigntemplateformfieldgroupSignerrequirement *FieldEEzsigntemplateformfieldgroupSignerrequirement `json:"eEzsigntemplateformfieldgroupSignerrequirement,omitempty"`
	// The Label for the Ezsigntemplateformfieldgroup
	SEzsigntemplateformfieldgroupLabel string `json:"sEzsigntemplateformfieldgroupLabel"`
	// The step when the Ezsigntemplatesigner will be invited to fill the form fields
	IEzsigntemplateformfieldgroupStep int32 `json:"iEzsigntemplateformfieldgroupStep"`
	// The default value for the Ezsigntemplateformfieldgroup  You can use the codes below and they will be replaced at signature time.    | Code | Description | Example | | ------------------------- | ------------ | ------------ | | {sUserFirstname} | The first name of the contact | John | | {sUserLastname} | The last name of the contact | Doe | | {sUserJobtitle} | The job title | Sales Representative | | {sEmailAddress} | The email address | email@example.com | | {sPhoneE164} | A phone number in E.164 Format | +15149901516 | | {sPhoneE164Cell} | A phone number in E.164 Format | +15149901516 |
	SEzsigntemplateformfieldgroupDefaultvalue *string `json:"sEzsigntemplateformfieldgroupDefaultvalue,omitempty"`
	// The minimum number of Ezsigntemplateformfield that must be filled in the Ezsigntemplateformfieldgroup
	IEzsigntemplateformfieldgroupFilledmin int32 `json:"iEzsigntemplateformfieldgroupFilledmin"`
	// The maximum number of Ezsigntemplateformfield that must be filled in the Ezsigntemplateformfieldgroup
	IEzsigntemplateformfieldgroupFilledmax int32 `json:"iEzsigntemplateformfieldgroupFilledmax"`
	// Whether the Ezsigntemplateformfieldgroup is read only or not.
	BEzsigntemplateformfieldgroupReadonly bool `json:"bEzsigntemplateformfieldgroupReadonly"`
	// The maximum length for the value in the Ezsigntemplateformfieldgroup  This can only be set if eEzsigntemplateformfieldgroupType is **Text** or **Textarea**
	IEzsigntemplateformfieldgroupMaxlength *int32 `json:"iEzsigntemplateformfieldgroupMaxlength,omitempty"`
	// Whether the Ezsigntemplateformfieldgroup is encrypted in the database or not. Encrypted values are not displayed on the Ezsigndocument. This can only be set if eEzsigntemplateformfieldgroupType is **Text** or **Textarea**
	BEzsigntemplateformfieldgroupEncrypted *bool `json:"bEzsigntemplateformfieldgroupEncrypted,omitempty"`
	// A regular expression to indicate what values are acceptable for the Ezsigntemplateformfieldgroup.  This can only be set if eEzsigntemplateformfieldgroupType is **Text** or **Textarea**
	SEzsigntemplateformfieldgroupRegexp *string `json:"sEzsigntemplateformfieldgroupRegexp,omitempty" validate:"regexp=^\\\\^.*\\\\$$|^$"`
	// Description of validation rule. Show by signatory.
	SEzsigntemplateformfieldgroupTextvalidationcustommessage *string `json:"sEzsigntemplateformfieldgroupTextvalidationcustommessage,omitempty"`
	EEzsigntemplateformfieldgroupTextvalidation *EnumTextvalidation `json:"eEzsigntemplateformfieldgroupTextvalidation,omitempty"`
	// A tooltip that will be presented to Ezsigntemplatesigner about the Ezsigntemplateformfieldgroup
	TEzsigntemplateformfieldgroupTooltip *string `json:"tEzsigntemplateformfieldgroupTooltip,omitempty"`
	EEzsigntemplateformfieldgroupTooltipposition *FieldEEzsigntemplateformfieldgroupTooltipposition `json:"eEzsigntemplateformfieldgroupTooltipposition,omitempty"`
}

type _EzsigntemplateformfieldgroupResponse EzsigntemplateformfieldgroupResponse

// NewEzsigntemplateformfieldgroupResponse instantiates a new EzsigntemplateformfieldgroupResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigntemplateformfieldgroupResponse(pkiEzsigntemplateformfieldgroupID int32, fkiEzsigntemplatedocumentID int32, eEzsigntemplateformfieldgroupType FieldEEzsigntemplateformfieldgroupType, sEzsigntemplateformfieldgroupLabel string, iEzsigntemplateformfieldgroupStep int32, iEzsigntemplateformfieldgroupFilledmin int32, iEzsigntemplateformfieldgroupFilledmax int32, bEzsigntemplateformfieldgroupReadonly bool) *EzsigntemplateformfieldgroupResponse {
	this := EzsigntemplateformfieldgroupResponse{}
	this.PkiEzsigntemplateformfieldgroupID = pkiEzsigntemplateformfieldgroupID
	this.FkiEzsigntemplatedocumentID = fkiEzsigntemplatedocumentID
	this.EEzsigntemplateformfieldgroupType = eEzsigntemplateformfieldgroupType
	this.SEzsigntemplateformfieldgroupLabel = sEzsigntemplateformfieldgroupLabel
	this.IEzsigntemplateformfieldgroupStep = iEzsigntemplateformfieldgroupStep
	this.IEzsigntemplateformfieldgroupFilledmin = iEzsigntemplateformfieldgroupFilledmin
	this.IEzsigntemplateformfieldgroupFilledmax = iEzsigntemplateformfieldgroupFilledmax
	this.BEzsigntemplateformfieldgroupReadonly = bEzsigntemplateformfieldgroupReadonly
	return &this
}

// NewEzsigntemplateformfieldgroupResponseWithDefaults instantiates a new EzsigntemplateformfieldgroupResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigntemplateformfieldgroupResponseWithDefaults() *EzsigntemplateformfieldgroupResponse {
	this := EzsigntemplateformfieldgroupResponse{}
	return &this
}

// GetPkiEzsigntemplateformfieldgroupID returns the PkiEzsigntemplateformfieldgroupID field value
func (o *EzsigntemplateformfieldgroupResponse) GetPkiEzsigntemplateformfieldgroupID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiEzsigntemplateformfieldgroupID
}

// GetPkiEzsigntemplateformfieldgroupIDOk returns a tuple with the PkiEzsigntemplateformfieldgroupID field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldgroupResponse) GetPkiEzsigntemplateformfieldgroupIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiEzsigntemplateformfieldgroupID, true
}

// SetPkiEzsigntemplateformfieldgroupID sets field value
func (o *EzsigntemplateformfieldgroupResponse) SetPkiEzsigntemplateformfieldgroupID(v int32) {
	o.PkiEzsigntemplateformfieldgroupID = v
}

// GetFkiEzsigntemplatedocumentID returns the FkiEzsigntemplatedocumentID field value
func (o *EzsigntemplateformfieldgroupResponse) GetFkiEzsigntemplatedocumentID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzsigntemplatedocumentID
}

// GetFkiEzsigntemplatedocumentIDOk returns a tuple with the FkiEzsigntemplatedocumentID field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldgroupResponse) GetFkiEzsigntemplatedocumentIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiEzsigntemplatedocumentID, true
}

// SetFkiEzsigntemplatedocumentID sets field value
func (o *EzsigntemplateformfieldgroupResponse) SetFkiEzsigntemplatedocumentID(v int32) {
	o.FkiEzsigntemplatedocumentID = v
}

// GetEEzsigntemplateformfieldgroupType returns the EEzsigntemplateformfieldgroupType field value
func (o *EzsigntemplateformfieldgroupResponse) GetEEzsigntemplateformfieldgroupType() FieldEEzsigntemplateformfieldgroupType {
	if o == nil {
		var ret FieldEEzsigntemplateformfieldgroupType
		return ret
	}

	return o.EEzsigntemplateformfieldgroupType
}

// GetEEzsigntemplateformfieldgroupTypeOk returns a tuple with the EEzsigntemplateformfieldgroupType field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldgroupResponse) GetEEzsigntemplateformfieldgroupTypeOk() (*FieldEEzsigntemplateformfieldgroupType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EEzsigntemplateformfieldgroupType, true
}

// SetEEzsigntemplateformfieldgroupType sets field value
func (o *EzsigntemplateformfieldgroupResponse) SetEEzsigntemplateformfieldgroupType(v FieldEEzsigntemplateformfieldgroupType) {
	o.EEzsigntemplateformfieldgroupType = v
}

// GetEEzsigntemplateformfieldgroupSignerrequirement returns the EEzsigntemplateformfieldgroupSignerrequirement field value if set, zero value otherwise.
// Deprecated
func (o *EzsigntemplateformfieldgroupResponse) GetEEzsigntemplateformfieldgroupSignerrequirement() FieldEEzsigntemplateformfieldgroupSignerrequirement {
	if o == nil || IsNil(o.EEzsigntemplateformfieldgroupSignerrequirement) {
		var ret FieldEEzsigntemplateformfieldgroupSignerrequirement
		return ret
	}
	return *o.EEzsigntemplateformfieldgroupSignerrequirement
}

// GetEEzsigntemplateformfieldgroupSignerrequirementOk returns a tuple with the EEzsigntemplateformfieldgroupSignerrequirement field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *EzsigntemplateformfieldgroupResponse) GetEEzsigntemplateformfieldgroupSignerrequirementOk() (*FieldEEzsigntemplateformfieldgroupSignerrequirement, bool) {
	if o == nil || IsNil(o.EEzsigntemplateformfieldgroupSignerrequirement) {
		return nil, false
	}
	return o.EEzsigntemplateformfieldgroupSignerrequirement, true
}

// HasEEzsigntemplateformfieldgroupSignerrequirement returns a boolean if a field has been set.
func (o *EzsigntemplateformfieldgroupResponse) HasEEzsigntemplateformfieldgroupSignerrequirement() bool {
	if o != nil && !IsNil(o.EEzsigntemplateformfieldgroupSignerrequirement) {
		return true
	}

	return false
}

// SetEEzsigntemplateformfieldgroupSignerrequirement gets a reference to the given FieldEEzsigntemplateformfieldgroupSignerrequirement and assigns it to the EEzsigntemplateformfieldgroupSignerrequirement field.
// Deprecated
func (o *EzsigntemplateformfieldgroupResponse) SetEEzsigntemplateformfieldgroupSignerrequirement(v FieldEEzsigntemplateformfieldgroupSignerrequirement) {
	o.EEzsigntemplateformfieldgroupSignerrequirement = &v
}

// GetSEzsigntemplateformfieldgroupLabel returns the SEzsigntemplateformfieldgroupLabel field value
func (o *EzsigntemplateformfieldgroupResponse) GetSEzsigntemplateformfieldgroupLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzsigntemplateformfieldgroupLabel
}

// GetSEzsigntemplateformfieldgroupLabelOk returns a tuple with the SEzsigntemplateformfieldgroupLabel field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldgroupResponse) GetSEzsigntemplateformfieldgroupLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SEzsigntemplateformfieldgroupLabel, true
}

// SetSEzsigntemplateformfieldgroupLabel sets field value
func (o *EzsigntemplateformfieldgroupResponse) SetSEzsigntemplateformfieldgroupLabel(v string) {
	o.SEzsigntemplateformfieldgroupLabel = v
}

// GetIEzsigntemplateformfieldgroupStep returns the IEzsigntemplateformfieldgroupStep field value
func (o *EzsigntemplateformfieldgroupResponse) GetIEzsigntemplateformfieldgroupStep() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsigntemplateformfieldgroupStep
}

// GetIEzsigntemplateformfieldgroupStepOk returns a tuple with the IEzsigntemplateformfieldgroupStep field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldgroupResponse) GetIEzsigntemplateformfieldgroupStepOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzsigntemplateformfieldgroupStep, true
}

// SetIEzsigntemplateformfieldgroupStep sets field value
func (o *EzsigntemplateformfieldgroupResponse) SetIEzsigntemplateformfieldgroupStep(v int32) {
	o.IEzsigntemplateformfieldgroupStep = v
}

// GetSEzsigntemplateformfieldgroupDefaultvalue returns the SEzsigntemplateformfieldgroupDefaultvalue field value if set, zero value otherwise.
func (o *EzsigntemplateformfieldgroupResponse) GetSEzsigntemplateformfieldgroupDefaultvalue() string {
	if o == nil || IsNil(o.SEzsigntemplateformfieldgroupDefaultvalue) {
		var ret string
		return ret
	}
	return *o.SEzsigntemplateformfieldgroupDefaultvalue
}

// GetSEzsigntemplateformfieldgroupDefaultvalueOk returns a tuple with the SEzsigntemplateformfieldgroupDefaultvalue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldgroupResponse) GetSEzsigntemplateformfieldgroupDefaultvalueOk() (*string, bool) {
	if o == nil || IsNil(o.SEzsigntemplateformfieldgroupDefaultvalue) {
		return nil, false
	}
	return o.SEzsigntemplateformfieldgroupDefaultvalue, true
}

// HasSEzsigntemplateformfieldgroupDefaultvalue returns a boolean if a field has been set.
func (o *EzsigntemplateformfieldgroupResponse) HasSEzsigntemplateformfieldgroupDefaultvalue() bool {
	if o != nil && !IsNil(o.SEzsigntemplateformfieldgroupDefaultvalue) {
		return true
	}

	return false
}

// SetSEzsigntemplateformfieldgroupDefaultvalue gets a reference to the given string and assigns it to the SEzsigntemplateformfieldgroupDefaultvalue field.
func (o *EzsigntemplateformfieldgroupResponse) SetSEzsigntemplateformfieldgroupDefaultvalue(v string) {
	o.SEzsigntemplateformfieldgroupDefaultvalue = &v
}

// GetIEzsigntemplateformfieldgroupFilledmin returns the IEzsigntemplateformfieldgroupFilledmin field value
func (o *EzsigntemplateformfieldgroupResponse) GetIEzsigntemplateformfieldgroupFilledmin() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsigntemplateformfieldgroupFilledmin
}

// GetIEzsigntemplateformfieldgroupFilledminOk returns a tuple with the IEzsigntemplateformfieldgroupFilledmin field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldgroupResponse) GetIEzsigntemplateformfieldgroupFilledminOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzsigntemplateformfieldgroupFilledmin, true
}

// SetIEzsigntemplateformfieldgroupFilledmin sets field value
func (o *EzsigntemplateformfieldgroupResponse) SetIEzsigntemplateformfieldgroupFilledmin(v int32) {
	o.IEzsigntemplateformfieldgroupFilledmin = v
}

// GetIEzsigntemplateformfieldgroupFilledmax returns the IEzsigntemplateformfieldgroupFilledmax field value
func (o *EzsigntemplateformfieldgroupResponse) GetIEzsigntemplateformfieldgroupFilledmax() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsigntemplateformfieldgroupFilledmax
}

// GetIEzsigntemplateformfieldgroupFilledmaxOk returns a tuple with the IEzsigntemplateformfieldgroupFilledmax field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldgroupResponse) GetIEzsigntemplateformfieldgroupFilledmaxOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzsigntemplateformfieldgroupFilledmax, true
}

// SetIEzsigntemplateformfieldgroupFilledmax sets field value
func (o *EzsigntemplateformfieldgroupResponse) SetIEzsigntemplateformfieldgroupFilledmax(v int32) {
	o.IEzsigntemplateformfieldgroupFilledmax = v
}

// GetBEzsigntemplateformfieldgroupReadonly returns the BEzsigntemplateformfieldgroupReadonly field value
func (o *EzsigntemplateformfieldgroupResponse) GetBEzsigntemplateformfieldgroupReadonly() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BEzsigntemplateformfieldgroupReadonly
}

// GetBEzsigntemplateformfieldgroupReadonlyOk returns a tuple with the BEzsigntemplateformfieldgroupReadonly field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldgroupResponse) GetBEzsigntemplateformfieldgroupReadonlyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BEzsigntemplateformfieldgroupReadonly, true
}

// SetBEzsigntemplateformfieldgroupReadonly sets field value
func (o *EzsigntemplateformfieldgroupResponse) SetBEzsigntemplateformfieldgroupReadonly(v bool) {
	o.BEzsigntemplateformfieldgroupReadonly = v
}

// GetIEzsigntemplateformfieldgroupMaxlength returns the IEzsigntemplateformfieldgroupMaxlength field value if set, zero value otherwise.
func (o *EzsigntemplateformfieldgroupResponse) GetIEzsigntemplateformfieldgroupMaxlength() int32 {
	if o == nil || IsNil(o.IEzsigntemplateformfieldgroupMaxlength) {
		var ret int32
		return ret
	}
	return *o.IEzsigntemplateformfieldgroupMaxlength
}

// GetIEzsigntemplateformfieldgroupMaxlengthOk returns a tuple with the IEzsigntemplateformfieldgroupMaxlength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldgroupResponse) GetIEzsigntemplateformfieldgroupMaxlengthOk() (*int32, bool) {
	if o == nil || IsNil(o.IEzsigntemplateformfieldgroupMaxlength) {
		return nil, false
	}
	return o.IEzsigntemplateformfieldgroupMaxlength, true
}

// HasIEzsigntemplateformfieldgroupMaxlength returns a boolean if a field has been set.
func (o *EzsigntemplateformfieldgroupResponse) HasIEzsigntemplateformfieldgroupMaxlength() bool {
	if o != nil && !IsNil(o.IEzsigntemplateformfieldgroupMaxlength) {
		return true
	}

	return false
}

// SetIEzsigntemplateformfieldgroupMaxlength gets a reference to the given int32 and assigns it to the IEzsigntemplateformfieldgroupMaxlength field.
func (o *EzsigntemplateformfieldgroupResponse) SetIEzsigntemplateformfieldgroupMaxlength(v int32) {
	o.IEzsigntemplateformfieldgroupMaxlength = &v
}

// GetBEzsigntemplateformfieldgroupEncrypted returns the BEzsigntemplateformfieldgroupEncrypted field value if set, zero value otherwise.
func (o *EzsigntemplateformfieldgroupResponse) GetBEzsigntemplateformfieldgroupEncrypted() bool {
	if o == nil || IsNil(o.BEzsigntemplateformfieldgroupEncrypted) {
		var ret bool
		return ret
	}
	return *o.BEzsigntemplateformfieldgroupEncrypted
}

// GetBEzsigntemplateformfieldgroupEncryptedOk returns a tuple with the BEzsigntemplateformfieldgroupEncrypted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldgroupResponse) GetBEzsigntemplateformfieldgroupEncryptedOk() (*bool, bool) {
	if o == nil || IsNil(o.BEzsigntemplateformfieldgroupEncrypted) {
		return nil, false
	}
	return o.BEzsigntemplateformfieldgroupEncrypted, true
}

// HasBEzsigntemplateformfieldgroupEncrypted returns a boolean if a field has been set.
func (o *EzsigntemplateformfieldgroupResponse) HasBEzsigntemplateformfieldgroupEncrypted() bool {
	if o != nil && !IsNil(o.BEzsigntemplateformfieldgroupEncrypted) {
		return true
	}

	return false
}

// SetBEzsigntemplateformfieldgroupEncrypted gets a reference to the given bool and assigns it to the BEzsigntemplateformfieldgroupEncrypted field.
func (o *EzsigntemplateformfieldgroupResponse) SetBEzsigntemplateformfieldgroupEncrypted(v bool) {
	o.BEzsigntemplateformfieldgroupEncrypted = &v
}

// GetSEzsigntemplateformfieldgroupRegexp returns the SEzsigntemplateformfieldgroupRegexp field value if set, zero value otherwise.
func (o *EzsigntemplateformfieldgroupResponse) GetSEzsigntemplateformfieldgroupRegexp() string {
	if o == nil || IsNil(o.SEzsigntemplateformfieldgroupRegexp) {
		var ret string
		return ret
	}
	return *o.SEzsigntemplateformfieldgroupRegexp
}

// GetSEzsigntemplateformfieldgroupRegexpOk returns a tuple with the SEzsigntemplateformfieldgroupRegexp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldgroupResponse) GetSEzsigntemplateformfieldgroupRegexpOk() (*string, bool) {
	if o == nil || IsNil(o.SEzsigntemplateformfieldgroupRegexp) {
		return nil, false
	}
	return o.SEzsigntemplateformfieldgroupRegexp, true
}

// HasSEzsigntemplateformfieldgroupRegexp returns a boolean if a field has been set.
func (o *EzsigntemplateformfieldgroupResponse) HasSEzsigntemplateformfieldgroupRegexp() bool {
	if o != nil && !IsNil(o.SEzsigntemplateformfieldgroupRegexp) {
		return true
	}

	return false
}

// SetSEzsigntemplateformfieldgroupRegexp gets a reference to the given string and assigns it to the SEzsigntemplateformfieldgroupRegexp field.
func (o *EzsigntemplateformfieldgroupResponse) SetSEzsigntemplateformfieldgroupRegexp(v string) {
	o.SEzsigntemplateformfieldgroupRegexp = &v
}

// GetSEzsigntemplateformfieldgroupTextvalidationcustommessage returns the SEzsigntemplateformfieldgroupTextvalidationcustommessage field value if set, zero value otherwise.
func (o *EzsigntemplateformfieldgroupResponse) GetSEzsigntemplateformfieldgroupTextvalidationcustommessage() string {
	if o == nil || IsNil(o.SEzsigntemplateformfieldgroupTextvalidationcustommessage) {
		var ret string
		return ret
	}
	return *o.SEzsigntemplateformfieldgroupTextvalidationcustommessage
}

// GetSEzsigntemplateformfieldgroupTextvalidationcustommessageOk returns a tuple with the SEzsigntemplateformfieldgroupTextvalidationcustommessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldgroupResponse) GetSEzsigntemplateformfieldgroupTextvalidationcustommessageOk() (*string, bool) {
	if o == nil || IsNil(o.SEzsigntemplateformfieldgroupTextvalidationcustommessage) {
		return nil, false
	}
	return o.SEzsigntemplateformfieldgroupTextvalidationcustommessage, true
}

// HasSEzsigntemplateformfieldgroupTextvalidationcustommessage returns a boolean if a field has been set.
func (o *EzsigntemplateformfieldgroupResponse) HasSEzsigntemplateformfieldgroupTextvalidationcustommessage() bool {
	if o != nil && !IsNil(o.SEzsigntemplateformfieldgroupTextvalidationcustommessage) {
		return true
	}

	return false
}

// SetSEzsigntemplateformfieldgroupTextvalidationcustommessage gets a reference to the given string and assigns it to the SEzsigntemplateformfieldgroupTextvalidationcustommessage field.
func (o *EzsigntemplateformfieldgroupResponse) SetSEzsigntemplateformfieldgroupTextvalidationcustommessage(v string) {
	o.SEzsigntemplateformfieldgroupTextvalidationcustommessage = &v
}

// GetEEzsigntemplateformfieldgroupTextvalidation returns the EEzsigntemplateformfieldgroupTextvalidation field value if set, zero value otherwise.
func (o *EzsigntemplateformfieldgroupResponse) GetEEzsigntemplateformfieldgroupTextvalidation() EnumTextvalidation {
	if o == nil || IsNil(o.EEzsigntemplateformfieldgroupTextvalidation) {
		var ret EnumTextvalidation
		return ret
	}
	return *o.EEzsigntemplateformfieldgroupTextvalidation
}

// GetEEzsigntemplateformfieldgroupTextvalidationOk returns a tuple with the EEzsigntemplateformfieldgroupTextvalidation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldgroupResponse) GetEEzsigntemplateformfieldgroupTextvalidationOk() (*EnumTextvalidation, bool) {
	if o == nil || IsNil(o.EEzsigntemplateformfieldgroupTextvalidation) {
		return nil, false
	}
	return o.EEzsigntemplateformfieldgroupTextvalidation, true
}

// HasEEzsigntemplateformfieldgroupTextvalidation returns a boolean if a field has been set.
func (o *EzsigntemplateformfieldgroupResponse) HasEEzsigntemplateformfieldgroupTextvalidation() bool {
	if o != nil && !IsNil(o.EEzsigntemplateformfieldgroupTextvalidation) {
		return true
	}

	return false
}

// SetEEzsigntemplateformfieldgroupTextvalidation gets a reference to the given EnumTextvalidation and assigns it to the EEzsigntemplateformfieldgroupTextvalidation field.
func (o *EzsigntemplateformfieldgroupResponse) SetEEzsigntemplateformfieldgroupTextvalidation(v EnumTextvalidation) {
	o.EEzsigntemplateformfieldgroupTextvalidation = &v
}

// GetTEzsigntemplateformfieldgroupTooltip returns the TEzsigntemplateformfieldgroupTooltip field value if set, zero value otherwise.
func (o *EzsigntemplateformfieldgroupResponse) GetTEzsigntemplateformfieldgroupTooltip() string {
	if o == nil || IsNil(o.TEzsigntemplateformfieldgroupTooltip) {
		var ret string
		return ret
	}
	return *o.TEzsigntemplateformfieldgroupTooltip
}

// GetTEzsigntemplateformfieldgroupTooltipOk returns a tuple with the TEzsigntemplateformfieldgroupTooltip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldgroupResponse) GetTEzsigntemplateformfieldgroupTooltipOk() (*string, bool) {
	if o == nil || IsNil(o.TEzsigntemplateformfieldgroupTooltip) {
		return nil, false
	}
	return o.TEzsigntemplateformfieldgroupTooltip, true
}

// HasTEzsigntemplateformfieldgroupTooltip returns a boolean if a field has been set.
func (o *EzsigntemplateformfieldgroupResponse) HasTEzsigntemplateformfieldgroupTooltip() bool {
	if o != nil && !IsNil(o.TEzsigntemplateformfieldgroupTooltip) {
		return true
	}

	return false
}

// SetTEzsigntemplateformfieldgroupTooltip gets a reference to the given string and assigns it to the TEzsigntemplateformfieldgroupTooltip field.
func (o *EzsigntemplateformfieldgroupResponse) SetTEzsigntemplateformfieldgroupTooltip(v string) {
	o.TEzsigntemplateformfieldgroupTooltip = &v
}

// GetEEzsigntemplateformfieldgroupTooltipposition returns the EEzsigntemplateformfieldgroupTooltipposition field value if set, zero value otherwise.
func (o *EzsigntemplateformfieldgroupResponse) GetEEzsigntemplateformfieldgroupTooltipposition() FieldEEzsigntemplateformfieldgroupTooltipposition {
	if o == nil || IsNil(o.EEzsigntemplateformfieldgroupTooltipposition) {
		var ret FieldEEzsigntemplateformfieldgroupTooltipposition
		return ret
	}
	return *o.EEzsigntemplateformfieldgroupTooltipposition
}

// GetEEzsigntemplateformfieldgroupTooltippositionOk returns a tuple with the EEzsigntemplateformfieldgroupTooltipposition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldgroupResponse) GetEEzsigntemplateformfieldgroupTooltippositionOk() (*FieldEEzsigntemplateformfieldgroupTooltipposition, bool) {
	if o == nil || IsNil(o.EEzsigntemplateformfieldgroupTooltipposition) {
		return nil, false
	}
	return o.EEzsigntemplateformfieldgroupTooltipposition, true
}

// HasEEzsigntemplateformfieldgroupTooltipposition returns a boolean if a field has been set.
func (o *EzsigntemplateformfieldgroupResponse) HasEEzsigntemplateformfieldgroupTooltipposition() bool {
	if o != nil && !IsNil(o.EEzsigntemplateformfieldgroupTooltipposition) {
		return true
	}

	return false
}

// SetEEzsigntemplateformfieldgroupTooltipposition gets a reference to the given FieldEEzsigntemplateformfieldgroupTooltipposition and assigns it to the EEzsigntemplateformfieldgroupTooltipposition field.
func (o *EzsigntemplateformfieldgroupResponse) SetEEzsigntemplateformfieldgroupTooltipposition(v FieldEEzsigntemplateformfieldgroupTooltipposition) {
	o.EEzsigntemplateformfieldgroupTooltipposition = &v
}

func (o EzsigntemplateformfieldgroupResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigntemplateformfieldgroupResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiEzsigntemplateformfieldgroupID"] = o.PkiEzsigntemplateformfieldgroupID
	toSerialize["fkiEzsigntemplatedocumentID"] = o.FkiEzsigntemplatedocumentID
	toSerialize["eEzsigntemplateformfieldgroupType"] = o.EEzsigntemplateformfieldgroupType
	if !IsNil(o.EEzsigntemplateformfieldgroupSignerrequirement) {
		toSerialize["eEzsigntemplateformfieldgroupSignerrequirement"] = o.EEzsigntemplateformfieldgroupSignerrequirement
	}
	toSerialize["sEzsigntemplateformfieldgroupLabel"] = o.SEzsigntemplateformfieldgroupLabel
	toSerialize["iEzsigntemplateformfieldgroupStep"] = o.IEzsigntemplateformfieldgroupStep
	if !IsNil(o.SEzsigntemplateformfieldgroupDefaultvalue) {
		toSerialize["sEzsigntemplateformfieldgroupDefaultvalue"] = o.SEzsigntemplateformfieldgroupDefaultvalue
	}
	toSerialize["iEzsigntemplateformfieldgroupFilledmin"] = o.IEzsigntemplateformfieldgroupFilledmin
	toSerialize["iEzsigntemplateformfieldgroupFilledmax"] = o.IEzsigntemplateformfieldgroupFilledmax
	toSerialize["bEzsigntemplateformfieldgroupReadonly"] = o.BEzsigntemplateformfieldgroupReadonly
	if !IsNil(o.IEzsigntemplateformfieldgroupMaxlength) {
		toSerialize["iEzsigntemplateformfieldgroupMaxlength"] = o.IEzsigntemplateformfieldgroupMaxlength
	}
	if !IsNil(o.BEzsigntemplateformfieldgroupEncrypted) {
		toSerialize["bEzsigntemplateformfieldgroupEncrypted"] = o.BEzsigntemplateformfieldgroupEncrypted
	}
	if !IsNil(o.SEzsigntemplateformfieldgroupRegexp) {
		toSerialize["sEzsigntemplateformfieldgroupRegexp"] = o.SEzsigntemplateformfieldgroupRegexp
	}
	if !IsNil(o.SEzsigntemplateformfieldgroupTextvalidationcustommessage) {
		toSerialize["sEzsigntemplateformfieldgroupTextvalidationcustommessage"] = o.SEzsigntemplateformfieldgroupTextvalidationcustommessage
	}
	if !IsNil(o.EEzsigntemplateformfieldgroupTextvalidation) {
		toSerialize["eEzsigntemplateformfieldgroupTextvalidation"] = o.EEzsigntemplateformfieldgroupTextvalidation
	}
	if !IsNil(o.TEzsigntemplateformfieldgroupTooltip) {
		toSerialize["tEzsigntemplateformfieldgroupTooltip"] = o.TEzsigntemplateformfieldgroupTooltip
	}
	if !IsNil(o.EEzsigntemplateformfieldgroupTooltipposition) {
		toSerialize["eEzsigntemplateformfieldgroupTooltipposition"] = o.EEzsigntemplateformfieldgroupTooltipposition
	}
	return toSerialize, nil
}

func (o *EzsigntemplateformfieldgroupResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiEzsigntemplateformfieldgroupID",
		"fkiEzsigntemplatedocumentID",
		"eEzsigntemplateformfieldgroupType",
		"sEzsigntemplateformfieldgroupLabel",
		"iEzsigntemplateformfieldgroupStep",
		"iEzsigntemplateformfieldgroupFilledmin",
		"iEzsigntemplateformfieldgroupFilledmax",
		"bEzsigntemplateformfieldgroupReadonly",
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

	varEzsigntemplateformfieldgroupResponse := _EzsigntemplateformfieldgroupResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsigntemplateformfieldgroupResponse)

	if err != nil {
		return err
	}

	*o = EzsigntemplateformfieldgroupResponse(varEzsigntemplateformfieldgroupResponse)

	return err
}

type NullableEzsigntemplateformfieldgroupResponse struct {
	value *EzsigntemplateformfieldgroupResponse
	isSet bool
}

func (v NullableEzsigntemplateformfieldgroupResponse) Get() *EzsigntemplateformfieldgroupResponse {
	return v.value
}

func (v *NullableEzsigntemplateformfieldgroupResponse) Set(val *EzsigntemplateformfieldgroupResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigntemplateformfieldgroupResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigntemplateformfieldgroupResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigntemplateformfieldgroupResponse(val *EzsigntemplateformfieldgroupResponse) *NullableEzsigntemplateformfieldgroupResponse {
	return &NullableEzsigntemplateformfieldgroupResponse{value: val, isSet: true}
}

func (v NullableEzsigntemplateformfieldgroupResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigntemplateformfieldgroupResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


