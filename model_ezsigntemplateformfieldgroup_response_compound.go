/*
eZmax API Definition (Full)

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.2.1
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the EzsigntemplateformfieldgroupResponseCompound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigntemplateformfieldgroupResponseCompound{}

// EzsigntemplateformfieldgroupResponseCompound A Ezsigntemplateformfieldgroup Object and children
type EzsigntemplateformfieldgroupResponseCompound struct {
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
	SEzsigntemplateformfieldgroupRegexp *string `json:"sEzsigntemplateformfieldgroupRegexp,omitempty"`
	EEzsigntemplateformfieldgroupTextvalidation *EnumTextvalidation `json:"eEzsigntemplateformfieldgroupTextvalidation,omitempty"`
	// A tooltip that will be presented to Ezsigntemplatesigner about the Ezsigntemplateformfieldgroup
	TEzsigntemplateformfieldgroupTooltip *string `json:"tEzsigntemplateformfieldgroupTooltip,omitempty"`
	EEzsigntemplateformfieldgroupTooltipposition *FieldEEzsigntemplateformfieldgroupTooltipposition `json:"eEzsigntemplateformfieldgroupTooltipposition,omitempty"`
	AObjEzsigntemplateformfieldgroupsigner []EzsigntemplateformfieldgroupsignerResponseCompound `json:"a_objEzsigntemplateformfieldgroupsigner"`
	AObjDropdownElement []CustomDropdownElementResponseCompound `json:"a_objDropdownElement,omitempty"`
	AObjEzsigntemplateformfield []EzsigntemplateformfieldResponseCompound `json:"a_objEzsigntemplateformfield"`
}

type _EzsigntemplateformfieldgroupResponseCompound EzsigntemplateformfieldgroupResponseCompound

// NewEzsigntemplateformfieldgroupResponseCompound instantiates a new EzsigntemplateformfieldgroupResponseCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigntemplateformfieldgroupResponseCompound(pkiEzsigntemplateformfieldgroupID int32, fkiEzsigntemplatedocumentID int32, eEzsigntemplateformfieldgroupType FieldEEzsigntemplateformfieldgroupType, sEzsigntemplateformfieldgroupLabel string, iEzsigntemplateformfieldgroupStep int32, iEzsigntemplateformfieldgroupFilledmin int32, iEzsigntemplateformfieldgroupFilledmax int32, bEzsigntemplateformfieldgroupReadonly bool, aObjEzsigntemplateformfieldgroupsigner []EzsigntemplateformfieldgroupsignerResponseCompound, aObjEzsigntemplateformfield []EzsigntemplateformfieldResponseCompound) *EzsigntemplateformfieldgroupResponseCompound {
	this := EzsigntemplateformfieldgroupResponseCompound{}
	this.PkiEzsigntemplateformfieldgroupID = pkiEzsigntemplateformfieldgroupID
	this.FkiEzsigntemplatedocumentID = fkiEzsigntemplatedocumentID
	this.EEzsigntemplateformfieldgroupType = eEzsigntemplateformfieldgroupType
	this.SEzsigntemplateformfieldgroupLabel = sEzsigntemplateformfieldgroupLabel
	this.IEzsigntemplateformfieldgroupStep = iEzsigntemplateformfieldgroupStep
	this.IEzsigntemplateformfieldgroupFilledmin = iEzsigntemplateformfieldgroupFilledmin
	this.IEzsigntemplateformfieldgroupFilledmax = iEzsigntemplateformfieldgroupFilledmax
	this.BEzsigntemplateformfieldgroupReadonly = bEzsigntemplateformfieldgroupReadonly
	this.AObjEzsigntemplateformfieldgroupsigner = aObjEzsigntemplateformfieldgroupsigner
	this.AObjEzsigntemplateformfield = aObjEzsigntemplateformfield
	return &this
}

// NewEzsigntemplateformfieldgroupResponseCompoundWithDefaults instantiates a new EzsigntemplateformfieldgroupResponseCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigntemplateformfieldgroupResponseCompoundWithDefaults() *EzsigntemplateformfieldgroupResponseCompound {
	this := EzsigntemplateformfieldgroupResponseCompound{}
	return &this
}

// GetPkiEzsigntemplateformfieldgroupID returns the PkiEzsigntemplateformfieldgroupID field value
func (o *EzsigntemplateformfieldgroupResponseCompound) GetPkiEzsigntemplateformfieldgroupID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiEzsigntemplateformfieldgroupID
}

// GetPkiEzsigntemplateformfieldgroupIDOk returns a tuple with the PkiEzsigntemplateformfieldgroupID field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldgroupResponseCompound) GetPkiEzsigntemplateformfieldgroupIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiEzsigntemplateformfieldgroupID, true
}

// SetPkiEzsigntemplateformfieldgroupID sets field value
func (o *EzsigntemplateformfieldgroupResponseCompound) SetPkiEzsigntemplateformfieldgroupID(v int32) {
	o.PkiEzsigntemplateformfieldgroupID = v
}

// GetFkiEzsigntemplatedocumentID returns the FkiEzsigntemplatedocumentID field value
func (o *EzsigntemplateformfieldgroupResponseCompound) GetFkiEzsigntemplatedocumentID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzsigntemplatedocumentID
}

// GetFkiEzsigntemplatedocumentIDOk returns a tuple with the FkiEzsigntemplatedocumentID field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldgroupResponseCompound) GetFkiEzsigntemplatedocumentIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiEzsigntemplatedocumentID, true
}

// SetFkiEzsigntemplatedocumentID sets field value
func (o *EzsigntemplateformfieldgroupResponseCompound) SetFkiEzsigntemplatedocumentID(v int32) {
	o.FkiEzsigntemplatedocumentID = v
}

// GetEEzsigntemplateformfieldgroupType returns the EEzsigntemplateformfieldgroupType field value
func (o *EzsigntemplateformfieldgroupResponseCompound) GetEEzsigntemplateformfieldgroupType() FieldEEzsigntemplateformfieldgroupType {
	if o == nil {
		var ret FieldEEzsigntemplateformfieldgroupType
		return ret
	}

	return o.EEzsigntemplateformfieldgroupType
}

// GetEEzsigntemplateformfieldgroupTypeOk returns a tuple with the EEzsigntemplateformfieldgroupType field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldgroupResponseCompound) GetEEzsigntemplateformfieldgroupTypeOk() (*FieldEEzsigntemplateformfieldgroupType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EEzsigntemplateformfieldgroupType, true
}

// SetEEzsigntemplateformfieldgroupType sets field value
func (o *EzsigntemplateformfieldgroupResponseCompound) SetEEzsigntemplateformfieldgroupType(v FieldEEzsigntemplateformfieldgroupType) {
	o.EEzsigntemplateformfieldgroupType = v
}

// GetEEzsigntemplateformfieldgroupSignerrequirement returns the EEzsigntemplateformfieldgroupSignerrequirement field value if set, zero value otherwise.
// Deprecated
func (o *EzsigntemplateformfieldgroupResponseCompound) GetEEzsigntemplateformfieldgroupSignerrequirement() FieldEEzsigntemplateformfieldgroupSignerrequirement {
	if o == nil || IsNil(o.EEzsigntemplateformfieldgroupSignerrequirement) {
		var ret FieldEEzsigntemplateformfieldgroupSignerrequirement
		return ret
	}
	return *o.EEzsigntemplateformfieldgroupSignerrequirement
}

// GetEEzsigntemplateformfieldgroupSignerrequirementOk returns a tuple with the EEzsigntemplateformfieldgroupSignerrequirement field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *EzsigntemplateformfieldgroupResponseCompound) GetEEzsigntemplateformfieldgroupSignerrequirementOk() (*FieldEEzsigntemplateformfieldgroupSignerrequirement, bool) {
	if o == nil || IsNil(o.EEzsigntemplateformfieldgroupSignerrequirement) {
		return nil, false
	}
	return o.EEzsigntemplateformfieldgroupSignerrequirement, true
}

// HasEEzsigntemplateformfieldgroupSignerrequirement returns a boolean if a field has been set.
func (o *EzsigntemplateformfieldgroupResponseCompound) HasEEzsigntemplateformfieldgroupSignerrequirement() bool {
	if o != nil && !IsNil(o.EEzsigntemplateformfieldgroupSignerrequirement) {
		return true
	}

	return false
}

// SetEEzsigntemplateformfieldgroupSignerrequirement gets a reference to the given FieldEEzsigntemplateformfieldgroupSignerrequirement and assigns it to the EEzsigntemplateformfieldgroupSignerrequirement field.
// Deprecated
func (o *EzsigntemplateformfieldgroupResponseCompound) SetEEzsigntemplateformfieldgroupSignerrequirement(v FieldEEzsigntemplateformfieldgroupSignerrequirement) {
	o.EEzsigntemplateformfieldgroupSignerrequirement = &v
}

// GetSEzsigntemplateformfieldgroupLabel returns the SEzsigntemplateformfieldgroupLabel field value
func (o *EzsigntemplateformfieldgroupResponseCompound) GetSEzsigntemplateformfieldgroupLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzsigntemplateformfieldgroupLabel
}

// GetSEzsigntemplateformfieldgroupLabelOk returns a tuple with the SEzsigntemplateformfieldgroupLabel field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldgroupResponseCompound) GetSEzsigntemplateformfieldgroupLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SEzsigntemplateformfieldgroupLabel, true
}

// SetSEzsigntemplateformfieldgroupLabel sets field value
func (o *EzsigntemplateformfieldgroupResponseCompound) SetSEzsigntemplateformfieldgroupLabel(v string) {
	o.SEzsigntemplateformfieldgroupLabel = v
}

// GetIEzsigntemplateformfieldgroupStep returns the IEzsigntemplateformfieldgroupStep field value
func (o *EzsigntemplateformfieldgroupResponseCompound) GetIEzsigntemplateformfieldgroupStep() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsigntemplateformfieldgroupStep
}

// GetIEzsigntemplateformfieldgroupStepOk returns a tuple with the IEzsigntemplateformfieldgroupStep field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldgroupResponseCompound) GetIEzsigntemplateformfieldgroupStepOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzsigntemplateformfieldgroupStep, true
}

// SetIEzsigntemplateformfieldgroupStep sets field value
func (o *EzsigntemplateformfieldgroupResponseCompound) SetIEzsigntemplateformfieldgroupStep(v int32) {
	o.IEzsigntemplateformfieldgroupStep = v
}

// GetSEzsigntemplateformfieldgroupDefaultvalue returns the SEzsigntemplateformfieldgroupDefaultvalue field value if set, zero value otherwise.
func (o *EzsigntemplateformfieldgroupResponseCompound) GetSEzsigntemplateformfieldgroupDefaultvalue() string {
	if o == nil || IsNil(o.SEzsigntemplateformfieldgroupDefaultvalue) {
		var ret string
		return ret
	}
	return *o.SEzsigntemplateformfieldgroupDefaultvalue
}

// GetSEzsigntemplateformfieldgroupDefaultvalueOk returns a tuple with the SEzsigntemplateformfieldgroupDefaultvalue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldgroupResponseCompound) GetSEzsigntemplateformfieldgroupDefaultvalueOk() (*string, bool) {
	if o == nil || IsNil(o.SEzsigntemplateformfieldgroupDefaultvalue) {
		return nil, false
	}
	return o.SEzsigntemplateformfieldgroupDefaultvalue, true
}

// HasSEzsigntemplateformfieldgroupDefaultvalue returns a boolean if a field has been set.
func (o *EzsigntemplateformfieldgroupResponseCompound) HasSEzsigntemplateformfieldgroupDefaultvalue() bool {
	if o != nil && !IsNil(o.SEzsigntemplateformfieldgroupDefaultvalue) {
		return true
	}

	return false
}

// SetSEzsigntemplateformfieldgroupDefaultvalue gets a reference to the given string and assigns it to the SEzsigntemplateformfieldgroupDefaultvalue field.
func (o *EzsigntemplateformfieldgroupResponseCompound) SetSEzsigntemplateformfieldgroupDefaultvalue(v string) {
	o.SEzsigntemplateformfieldgroupDefaultvalue = &v
}

// GetIEzsigntemplateformfieldgroupFilledmin returns the IEzsigntemplateformfieldgroupFilledmin field value
func (o *EzsigntemplateformfieldgroupResponseCompound) GetIEzsigntemplateformfieldgroupFilledmin() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsigntemplateformfieldgroupFilledmin
}

// GetIEzsigntemplateformfieldgroupFilledminOk returns a tuple with the IEzsigntemplateformfieldgroupFilledmin field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldgroupResponseCompound) GetIEzsigntemplateformfieldgroupFilledminOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzsigntemplateformfieldgroupFilledmin, true
}

// SetIEzsigntemplateformfieldgroupFilledmin sets field value
func (o *EzsigntemplateformfieldgroupResponseCompound) SetIEzsigntemplateformfieldgroupFilledmin(v int32) {
	o.IEzsigntemplateformfieldgroupFilledmin = v
}

// GetIEzsigntemplateformfieldgroupFilledmax returns the IEzsigntemplateformfieldgroupFilledmax field value
func (o *EzsigntemplateformfieldgroupResponseCompound) GetIEzsigntemplateformfieldgroupFilledmax() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsigntemplateformfieldgroupFilledmax
}

// GetIEzsigntemplateformfieldgroupFilledmaxOk returns a tuple with the IEzsigntemplateformfieldgroupFilledmax field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldgroupResponseCompound) GetIEzsigntemplateformfieldgroupFilledmaxOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzsigntemplateformfieldgroupFilledmax, true
}

// SetIEzsigntemplateformfieldgroupFilledmax sets field value
func (o *EzsigntemplateformfieldgroupResponseCompound) SetIEzsigntemplateformfieldgroupFilledmax(v int32) {
	o.IEzsigntemplateformfieldgroupFilledmax = v
}

// GetBEzsigntemplateformfieldgroupReadonly returns the BEzsigntemplateformfieldgroupReadonly field value
func (o *EzsigntemplateformfieldgroupResponseCompound) GetBEzsigntemplateformfieldgroupReadonly() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BEzsigntemplateformfieldgroupReadonly
}

// GetBEzsigntemplateformfieldgroupReadonlyOk returns a tuple with the BEzsigntemplateformfieldgroupReadonly field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldgroupResponseCompound) GetBEzsigntemplateformfieldgroupReadonlyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BEzsigntemplateformfieldgroupReadonly, true
}

// SetBEzsigntemplateformfieldgroupReadonly sets field value
func (o *EzsigntemplateformfieldgroupResponseCompound) SetBEzsigntemplateformfieldgroupReadonly(v bool) {
	o.BEzsigntemplateformfieldgroupReadonly = v
}

// GetIEzsigntemplateformfieldgroupMaxlength returns the IEzsigntemplateformfieldgroupMaxlength field value if set, zero value otherwise.
func (o *EzsigntemplateformfieldgroupResponseCompound) GetIEzsigntemplateformfieldgroupMaxlength() int32 {
	if o == nil || IsNil(o.IEzsigntemplateformfieldgroupMaxlength) {
		var ret int32
		return ret
	}
	return *o.IEzsigntemplateformfieldgroupMaxlength
}

// GetIEzsigntemplateformfieldgroupMaxlengthOk returns a tuple with the IEzsigntemplateformfieldgroupMaxlength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldgroupResponseCompound) GetIEzsigntemplateformfieldgroupMaxlengthOk() (*int32, bool) {
	if o == nil || IsNil(o.IEzsigntemplateformfieldgroupMaxlength) {
		return nil, false
	}
	return o.IEzsigntemplateformfieldgroupMaxlength, true
}

// HasIEzsigntemplateformfieldgroupMaxlength returns a boolean if a field has been set.
func (o *EzsigntemplateformfieldgroupResponseCompound) HasIEzsigntemplateformfieldgroupMaxlength() bool {
	if o != nil && !IsNil(o.IEzsigntemplateformfieldgroupMaxlength) {
		return true
	}

	return false
}

// SetIEzsigntemplateformfieldgroupMaxlength gets a reference to the given int32 and assigns it to the IEzsigntemplateformfieldgroupMaxlength field.
func (o *EzsigntemplateformfieldgroupResponseCompound) SetIEzsigntemplateformfieldgroupMaxlength(v int32) {
	o.IEzsigntemplateformfieldgroupMaxlength = &v
}

// GetBEzsigntemplateformfieldgroupEncrypted returns the BEzsigntemplateformfieldgroupEncrypted field value if set, zero value otherwise.
func (o *EzsigntemplateformfieldgroupResponseCompound) GetBEzsigntemplateformfieldgroupEncrypted() bool {
	if o == nil || IsNil(o.BEzsigntemplateformfieldgroupEncrypted) {
		var ret bool
		return ret
	}
	return *o.BEzsigntemplateformfieldgroupEncrypted
}

// GetBEzsigntemplateformfieldgroupEncryptedOk returns a tuple with the BEzsigntemplateformfieldgroupEncrypted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldgroupResponseCompound) GetBEzsigntemplateformfieldgroupEncryptedOk() (*bool, bool) {
	if o == nil || IsNil(o.BEzsigntemplateformfieldgroupEncrypted) {
		return nil, false
	}
	return o.BEzsigntemplateformfieldgroupEncrypted, true
}

// HasBEzsigntemplateformfieldgroupEncrypted returns a boolean if a field has been set.
func (o *EzsigntemplateformfieldgroupResponseCompound) HasBEzsigntemplateformfieldgroupEncrypted() bool {
	if o != nil && !IsNil(o.BEzsigntemplateformfieldgroupEncrypted) {
		return true
	}

	return false
}

// SetBEzsigntemplateformfieldgroupEncrypted gets a reference to the given bool and assigns it to the BEzsigntemplateformfieldgroupEncrypted field.
func (o *EzsigntemplateformfieldgroupResponseCompound) SetBEzsigntemplateformfieldgroupEncrypted(v bool) {
	o.BEzsigntemplateformfieldgroupEncrypted = &v
}

// GetSEzsigntemplateformfieldgroupRegexp returns the SEzsigntemplateformfieldgroupRegexp field value if set, zero value otherwise.
func (o *EzsigntemplateformfieldgroupResponseCompound) GetSEzsigntemplateformfieldgroupRegexp() string {
	if o == nil || IsNil(o.SEzsigntemplateformfieldgroupRegexp) {
		var ret string
		return ret
	}
	return *o.SEzsigntemplateformfieldgroupRegexp
}

// GetSEzsigntemplateformfieldgroupRegexpOk returns a tuple with the SEzsigntemplateformfieldgroupRegexp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldgroupResponseCompound) GetSEzsigntemplateformfieldgroupRegexpOk() (*string, bool) {
	if o == nil || IsNil(o.SEzsigntemplateformfieldgroupRegexp) {
		return nil, false
	}
	return o.SEzsigntemplateformfieldgroupRegexp, true
}

// HasSEzsigntemplateformfieldgroupRegexp returns a boolean if a field has been set.
func (o *EzsigntemplateformfieldgroupResponseCompound) HasSEzsigntemplateformfieldgroupRegexp() bool {
	if o != nil && !IsNil(o.SEzsigntemplateformfieldgroupRegexp) {
		return true
	}

	return false
}

// SetSEzsigntemplateformfieldgroupRegexp gets a reference to the given string and assigns it to the SEzsigntemplateformfieldgroupRegexp field.
func (o *EzsigntemplateformfieldgroupResponseCompound) SetSEzsigntemplateformfieldgroupRegexp(v string) {
	o.SEzsigntemplateformfieldgroupRegexp = &v
}

// GetEEzsigntemplateformfieldgroupTextvalidation returns the EEzsigntemplateformfieldgroupTextvalidation field value if set, zero value otherwise.
func (o *EzsigntemplateformfieldgroupResponseCompound) GetEEzsigntemplateformfieldgroupTextvalidation() EnumTextvalidation {
	if o == nil || IsNil(o.EEzsigntemplateformfieldgroupTextvalidation) {
		var ret EnumTextvalidation
		return ret
	}
	return *o.EEzsigntemplateformfieldgroupTextvalidation
}

// GetEEzsigntemplateformfieldgroupTextvalidationOk returns a tuple with the EEzsigntemplateformfieldgroupTextvalidation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldgroupResponseCompound) GetEEzsigntemplateformfieldgroupTextvalidationOk() (*EnumTextvalidation, bool) {
	if o == nil || IsNil(o.EEzsigntemplateformfieldgroupTextvalidation) {
		return nil, false
	}
	return o.EEzsigntemplateformfieldgroupTextvalidation, true
}

// HasEEzsigntemplateformfieldgroupTextvalidation returns a boolean if a field has been set.
func (o *EzsigntemplateformfieldgroupResponseCompound) HasEEzsigntemplateformfieldgroupTextvalidation() bool {
	if o != nil && !IsNil(o.EEzsigntemplateformfieldgroupTextvalidation) {
		return true
	}

	return false
}

// SetEEzsigntemplateformfieldgroupTextvalidation gets a reference to the given EnumTextvalidation and assigns it to the EEzsigntemplateformfieldgroupTextvalidation field.
func (o *EzsigntemplateformfieldgroupResponseCompound) SetEEzsigntemplateformfieldgroupTextvalidation(v EnumTextvalidation) {
	o.EEzsigntemplateformfieldgroupTextvalidation = &v
}

// GetTEzsigntemplateformfieldgroupTooltip returns the TEzsigntemplateformfieldgroupTooltip field value if set, zero value otherwise.
func (o *EzsigntemplateformfieldgroupResponseCompound) GetTEzsigntemplateformfieldgroupTooltip() string {
	if o == nil || IsNil(o.TEzsigntemplateformfieldgroupTooltip) {
		var ret string
		return ret
	}
	return *o.TEzsigntemplateformfieldgroupTooltip
}

// GetTEzsigntemplateformfieldgroupTooltipOk returns a tuple with the TEzsigntemplateformfieldgroupTooltip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldgroupResponseCompound) GetTEzsigntemplateformfieldgroupTooltipOk() (*string, bool) {
	if o == nil || IsNil(o.TEzsigntemplateformfieldgroupTooltip) {
		return nil, false
	}
	return o.TEzsigntemplateformfieldgroupTooltip, true
}

// HasTEzsigntemplateformfieldgroupTooltip returns a boolean if a field has been set.
func (o *EzsigntemplateformfieldgroupResponseCompound) HasTEzsigntemplateformfieldgroupTooltip() bool {
	if o != nil && !IsNil(o.TEzsigntemplateformfieldgroupTooltip) {
		return true
	}

	return false
}

// SetTEzsigntemplateformfieldgroupTooltip gets a reference to the given string and assigns it to the TEzsigntemplateformfieldgroupTooltip field.
func (o *EzsigntemplateformfieldgroupResponseCompound) SetTEzsigntemplateformfieldgroupTooltip(v string) {
	o.TEzsigntemplateformfieldgroupTooltip = &v
}

// GetEEzsigntemplateformfieldgroupTooltipposition returns the EEzsigntemplateformfieldgroupTooltipposition field value if set, zero value otherwise.
func (o *EzsigntemplateformfieldgroupResponseCompound) GetEEzsigntemplateformfieldgroupTooltipposition() FieldEEzsigntemplateformfieldgroupTooltipposition {
	if o == nil || IsNil(o.EEzsigntemplateformfieldgroupTooltipposition) {
		var ret FieldEEzsigntemplateformfieldgroupTooltipposition
		return ret
	}
	return *o.EEzsigntemplateformfieldgroupTooltipposition
}

// GetEEzsigntemplateformfieldgroupTooltippositionOk returns a tuple with the EEzsigntemplateformfieldgroupTooltipposition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldgroupResponseCompound) GetEEzsigntemplateformfieldgroupTooltippositionOk() (*FieldEEzsigntemplateformfieldgroupTooltipposition, bool) {
	if o == nil || IsNil(o.EEzsigntemplateformfieldgroupTooltipposition) {
		return nil, false
	}
	return o.EEzsigntemplateformfieldgroupTooltipposition, true
}

// HasEEzsigntemplateformfieldgroupTooltipposition returns a boolean if a field has been set.
func (o *EzsigntemplateformfieldgroupResponseCompound) HasEEzsigntemplateformfieldgroupTooltipposition() bool {
	if o != nil && !IsNil(o.EEzsigntemplateformfieldgroupTooltipposition) {
		return true
	}

	return false
}

// SetEEzsigntemplateformfieldgroupTooltipposition gets a reference to the given FieldEEzsigntemplateformfieldgroupTooltipposition and assigns it to the EEzsigntemplateformfieldgroupTooltipposition field.
func (o *EzsigntemplateformfieldgroupResponseCompound) SetEEzsigntemplateformfieldgroupTooltipposition(v FieldEEzsigntemplateformfieldgroupTooltipposition) {
	o.EEzsigntemplateformfieldgroupTooltipposition = &v
}

// GetAObjEzsigntemplateformfieldgroupsigner returns the AObjEzsigntemplateformfieldgroupsigner field value
func (o *EzsigntemplateformfieldgroupResponseCompound) GetAObjEzsigntemplateformfieldgroupsigner() []EzsigntemplateformfieldgroupsignerResponseCompound {
	if o == nil {
		var ret []EzsigntemplateformfieldgroupsignerResponseCompound
		return ret
	}

	return o.AObjEzsigntemplateformfieldgroupsigner
}

// GetAObjEzsigntemplateformfieldgroupsignerOk returns a tuple with the AObjEzsigntemplateformfieldgroupsigner field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldgroupResponseCompound) GetAObjEzsigntemplateformfieldgroupsignerOk() ([]EzsigntemplateformfieldgroupsignerResponseCompound, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjEzsigntemplateformfieldgroupsigner, true
}

// SetAObjEzsigntemplateformfieldgroupsigner sets field value
func (o *EzsigntemplateformfieldgroupResponseCompound) SetAObjEzsigntemplateformfieldgroupsigner(v []EzsigntemplateformfieldgroupsignerResponseCompound) {
	o.AObjEzsigntemplateformfieldgroupsigner = v
}

// GetAObjDropdownElement returns the AObjDropdownElement field value if set, zero value otherwise.
func (o *EzsigntemplateformfieldgroupResponseCompound) GetAObjDropdownElement() []CustomDropdownElementResponseCompound {
	if o == nil || IsNil(o.AObjDropdownElement) {
		var ret []CustomDropdownElementResponseCompound
		return ret
	}
	return o.AObjDropdownElement
}

// GetAObjDropdownElementOk returns a tuple with the AObjDropdownElement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldgroupResponseCompound) GetAObjDropdownElementOk() ([]CustomDropdownElementResponseCompound, bool) {
	if o == nil || IsNil(o.AObjDropdownElement) {
		return nil, false
	}
	return o.AObjDropdownElement, true
}

// HasAObjDropdownElement returns a boolean if a field has been set.
func (o *EzsigntemplateformfieldgroupResponseCompound) HasAObjDropdownElement() bool {
	if o != nil && !IsNil(o.AObjDropdownElement) {
		return true
	}

	return false
}

// SetAObjDropdownElement gets a reference to the given []CustomDropdownElementResponseCompound and assigns it to the AObjDropdownElement field.
func (o *EzsigntemplateformfieldgroupResponseCompound) SetAObjDropdownElement(v []CustomDropdownElementResponseCompound) {
	o.AObjDropdownElement = v
}

// GetAObjEzsigntemplateformfield returns the AObjEzsigntemplateformfield field value
func (o *EzsigntemplateformfieldgroupResponseCompound) GetAObjEzsigntemplateformfield() []EzsigntemplateformfieldResponseCompound {
	if o == nil {
		var ret []EzsigntemplateformfieldResponseCompound
		return ret
	}

	return o.AObjEzsigntemplateformfield
}

// GetAObjEzsigntemplateformfieldOk returns a tuple with the AObjEzsigntemplateformfield field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldgroupResponseCompound) GetAObjEzsigntemplateformfieldOk() ([]EzsigntemplateformfieldResponseCompound, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjEzsigntemplateformfield, true
}

// SetAObjEzsigntemplateformfield sets field value
func (o *EzsigntemplateformfieldgroupResponseCompound) SetAObjEzsigntemplateformfield(v []EzsigntemplateformfieldResponseCompound) {
	o.AObjEzsigntemplateformfield = v
}

func (o EzsigntemplateformfieldgroupResponseCompound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigntemplateformfieldgroupResponseCompound) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.EEzsigntemplateformfieldgroupTextvalidation) {
		toSerialize["eEzsigntemplateformfieldgroupTextvalidation"] = o.EEzsigntemplateformfieldgroupTextvalidation
	}
	if !IsNil(o.TEzsigntemplateformfieldgroupTooltip) {
		toSerialize["tEzsigntemplateformfieldgroupTooltip"] = o.TEzsigntemplateformfieldgroupTooltip
	}
	if !IsNil(o.EEzsigntemplateformfieldgroupTooltipposition) {
		toSerialize["eEzsigntemplateformfieldgroupTooltipposition"] = o.EEzsigntemplateformfieldgroupTooltipposition
	}
	toSerialize["a_objEzsigntemplateformfieldgroupsigner"] = o.AObjEzsigntemplateformfieldgroupsigner
	if !IsNil(o.AObjDropdownElement) {
		toSerialize["a_objDropdownElement"] = o.AObjDropdownElement
	}
	toSerialize["a_objEzsigntemplateformfield"] = o.AObjEzsigntemplateformfield
	return toSerialize, nil
}

func (o *EzsigntemplateformfieldgroupResponseCompound) UnmarshalJSON(data []byte) (err error) {
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
		"a_objEzsigntemplateformfieldgroupsigner",
		"a_objEzsigntemplateformfield",
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

	varEzsigntemplateformfieldgroupResponseCompound := _EzsigntemplateformfieldgroupResponseCompound{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsigntemplateformfieldgroupResponseCompound)

	if err != nil {
		return err
	}

	*o = EzsigntemplateformfieldgroupResponseCompound(varEzsigntemplateformfieldgroupResponseCompound)

	return err
}

type NullableEzsigntemplateformfieldgroupResponseCompound struct {
	value *EzsigntemplateformfieldgroupResponseCompound
	isSet bool
}

func (v NullableEzsigntemplateformfieldgroupResponseCompound) Get() *EzsigntemplateformfieldgroupResponseCompound {
	return v.value
}

func (v *NullableEzsigntemplateformfieldgroupResponseCompound) Set(val *EzsigntemplateformfieldgroupResponseCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigntemplateformfieldgroupResponseCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigntemplateformfieldgroupResponseCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigntemplateformfieldgroupResponseCompound(val *EzsigntemplateformfieldgroupResponseCompound) *NullableEzsigntemplateformfieldgroupResponseCompound {
	return &NullableEzsigntemplateformfieldgroupResponseCompound{value: val, isSet: true}
}

func (v NullableEzsigntemplateformfieldgroupResponseCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigntemplateformfieldgroupResponseCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


