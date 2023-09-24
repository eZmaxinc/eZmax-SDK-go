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

// checks if the EzsignformfieldgroupResponseCompound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignformfieldgroupResponseCompound{}

// EzsignformfieldgroupResponseCompound An Ezsignformfieldgroup Object and children to create a complete structure
type EzsignformfieldgroupResponseCompound struct {
	// The unique ID of the Ezsignformfieldgroup
	PkiEzsignformfieldgroupID int32 `json:"pkiEzsignformfieldgroupID"`
	// The unique ID of the Ezsigndocument
	FkiEzsigndocumentID int32 `json:"fkiEzsigndocumentID"`
	EEzsignformfieldgroupType FieldEEzsignformfieldgroupType `json:"eEzsignformfieldgroupType"`
	EEzsignformfieldgroupSignerrequirement FieldEEzsignformfieldgroupSignerrequirement `json:"eEzsignformfieldgroupSignerrequirement"`
	// The Label for the Ezsignformfieldgroup
	SEzsignformfieldgroupLabel string `json:"sEzsignformfieldgroupLabel"`
	// The step when the Ezsignsigner will be invited to fill the form fields
	IEzsignformfieldgroupStep int32 `json:"iEzsignformfieldgroupStep"`
	// The default value for the Ezsignformfieldgroup
	SEzsignformfieldgroupDefaultvalue *string `json:"sEzsignformfieldgroupDefaultvalue,omitempty"`
	// The minimum number of Ezsignformfield that must be filled in the Ezsignformfieldgroup
	IEzsignformfieldgroupFilledmin int32 `json:"iEzsignformfieldgroupFilledmin"`
	// The maximum number of Ezsignformfield that must be filled in the Ezsignformfieldgroup
	IEzsignformfieldgroupFilledmax int32 `json:"iEzsignformfieldgroupFilledmax"`
	// Whether the Ezsignformfieldgroup is read only or not.
	BEzsignformfieldgroupReadonly bool `json:"bEzsignformfieldgroupReadonly"`
	// The maximum length for the value in the Ezsignformfieldgroup  This can only be set if eEzsignformfieldgroupType is **Text** or **Textarea**
	IEzsignformfieldgroupMaxlength *int32 `json:"iEzsignformfieldgroupMaxlength,omitempty"`
	// Whether the Ezsignformfieldgroup is encrypted in the database or not. Encrypted values are not displayed on the Ezsigndocument. This can only be set if eEzsignformfieldgroupType is **Text** or **Textarea**
	BEzsignformfieldgroupEncrypted *bool `json:"bEzsignformfieldgroupEncrypted,omitempty"`
	EEzsignformfieldgroupTextvalidation *EnumTextvalidation `json:"eEzsignformfieldgroupTextvalidation,omitempty"`
	// A regular expression to indicate what values are acceptable for the Ezsignformfieldgroup.  This can only be set if eEzsignformfieldgroupType is **Text** or **Textarea**
	SEzsignformfieldgroupRegexp *string `json:"sEzsignformfieldgroupRegexp,omitempty"`
	// A tooltip that will be presented to Ezsignsigner about the Ezsignformfieldgroup
	TEzsignformfieldgroupTooltip *string `json:"tEzsignformfieldgroupTooltip,omitempty"`
	EEzsignformfieldgroupTooltipposition *FieldEEzsignformfieldgroupTooltipposition `json:"eEzsignformfieldgroupTooltipposition,omitempty"`
	AObjEzsignformfield []EzsignformfieldResponseCompound `json:"a_objEzsignformfield"`
	AObjDropdownElement []CustomDropdownElementResponseCompound `json:"a_objDropdownElement,omitempty"`
	AObjEzsignformfieldgroupsigner []EzsignformfieldgroupsignerResponseCompound `json:"a_objEzsignformfieldgroupsigner"`
}

// NewEzsignformfieldgroupResponseCompound instantiates a new EzsignformfieldgroupResponseCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignformfieldgroupResponseCompound(pkiEzsignformfieldgroupID int32, fkiEzsigndocumentID int32, eEzsignformfieldgroupType FieldEEzsignformfieldgroupType, eEzsignformfieldgroupSignerrequirement FieldEEzsignformfieldgroupSignerrequirement, sEzsignformfieldgroupLabel string, iEzsignformfieldgroupStep int32, iEzsignformfieldgroupFilledmin int32, iEzsignformfieldgroupFilledmax int32, bEzsignformfieldgroupReadonly bool, aObjEzsignformfield []EzsignformfieldResponseCompound, aObjEzsignformfieldgroupsigner []EzsignformfieldgroupsignerResponseCompound) *EzsignformfieldgroupResponseCompound {
	this := EzsignformfieldgroupResponseCompound{}
	this.PkiEzsignformfieldgroupID = pkiEzsignformfieldgroupID
	this.FkiEzsigndocumentID = fkiEzsigndocumentID
	this.EEzsignformfieldgroupType = eEzsignformfieldgroupType
	this.EEzsignformfieldgroupSignerrequirement = eEzsignformfieldgroupSignerrequirement
	this.SEzsignformfieldgroupLabel = sEzsignformfieldgroupLabel
	this.IEzsignformfieldgroupStep = iEzsignformfieldgroupStep
	this.IEzsignformfieldgroupFilledmin = iEzsignformfieldgroupFilledmin
	this.IEzsignformfieldgroupFilledmax = iEzsignformfieldgroupFilledmax
	this.BEzsignformfieldgroupReadonly = bEzsignformfieldgroupReadonly
	this.AObjEzsignformfield = aObjEzsignformfield
	this.AObjEzsignformfieldgroupsigner = aObjEzsignformfieldgroupsigner
	return &this
}

// NewEzsignformfieldgroupResponseCompoundWithDefaults instantiates a new EzsignformfieldgroupResponseCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignformfieldgroupResponseCompoundWithDefaults() *EzsignformfieldgroupResponseCompound {
	this := EzsignformfieldgroupResponseCompound{}
	return &this
}

// GetPkiEzsignformfieldgroupID returns the PkiEzsignformfieldgroupID field value
func (o *EzsignformfieldgroupResponseCompound) GetPkiEzsignformfieldgroupID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiEzsignformfieldgroupID
}

// GetPkiEzsignformfieldgroupIDOk returns a tuple with the PkiEzsignformfieldgroupID field value
// and a boolean to check if the value has been set.
func (o *EzsignformfieldgroupResponseCompound) GetPkiEzsignformfieldgroupIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiEzsignformfieldgroupID, true
}

// SetPkiEzsignformfieldgroupID sets field value
func (o *EzsignformfieldgroupResponseCompound) SetPkiEzsignformfieldgroupID(v int32) {
	o.PkiEzsignformfieldgroupID = v
}

// GetFkiEzsigndocumentID returns the FkiEzsigndocumentID field value
func (o *EzsignformfieldgroupResponseCompound) GetFkiEzsigndocumentID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzsigndocumentID
}

// GetFkiEzsigndocumentIDOk returns a tuple with the FkiEzsigndocumentID field value
// and a boolean to check if the value has been set.
func (o *EzsignformfieldgroupResponseCompound) GetFkiEzsigndocumentIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiEzsigndocumentID, true
}

// SetFkiEzsigndocumentID sets field value
func (o *EzsignformfieldgroupResponseCompound) SetFkiEzsigndocumentID(v int32) {
	o.FkiEzsigndocumentID = v
}

// GetEEzsignformfieldgroupType returns the EEzsignformfieldgroupType field value
func (o *EzsignformfieldgroupResponseCompound) GetEEzsignformfieldgroupType() FieldEEzsignformfieldgroupType {
	if o == nil {
		var ret FieldEEzsignformfieldgroupType
		return ret
	}

	return o.EEzsignformfieldgroupType
}

// GetEEzsignformfieldgroupTypeOk returns a tuple with the EEzsignformfieldgroupType field value
// and a boolean to check if the value has been set.
func (o *EzsignformfieldgroupResponseCompound) GetEEzsignformfieldgroupTypeOk() (*FieldEEzsignformfieldgroupType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EEzsignformfieldgroupType, true
}

// SetEEzsignformfieldgroupType sets field value
func (o *EzsignformfieldgroupResponseCompound) SetEEzsignformfieldgroupType(v FieldEEzsignformfieldgroupType) {
	o.EEzsignformfieldgroupType = v
}

// GetEEzsignformfieldgroupSignerrequirement returns the EEzsignformfieldgroupSignerrequirement field value
func (o *EzsignformfieldgroupResponseCompound) GetEEzsignformfieldgroupSignerrequirement() FieldEEzsignformfieldgroupSignerrequirement {
	if o == nil {
		var ret FieldEEzsignformfieldgroupSignerrequirement
		return ret
	}

	return o.EEzsignformfieldgroupSignerrequirement
}

// GetEEzsignformfieldgroupSignerrequirementOk returns a tuple with the EEzsignformfieldgroupSignerrequirement field value
// and a boolean to check if the value has been set.
func (o *EzsignformfieldgroupResponseCompound) GetEEzsignformfieldgroupSignerrequirementOk() (*FieldEEzsignformfieldgroupSignerrequirement, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EEzsignformfieldgroupSignerrequirement, true
}

// SetEEzsignformfieldgroupSignerrequirement sets field value
func (o *EzsignformfieldgroupResponseCompound) SetEEzsignformfieldgroupSignerrequirement(v FieldEEzsignformfieldgroupSignerrequirement) {
	o.EEzsignformfieldgroupSignerrequirement = v
}

// GetSEzsignformfieldgroupLabel returns the SEzsignformfieldgroupLabel field value
func (o *EzsignformfieldgroupResponseCompound) GetSEzsignformfieldgroupLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzsignformfieldgroupLabel
}

// GetSEzsignformfieldgroupLabelOk returns a tuple with the SEzsignformfieldgroupLabel field value
// and a boolean to check if the value has been set.
func (o *EzsignformfieldgroupResponseCompound) GetSEzsignformfieldgroupLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SEzsignformfieldgroupLabel, true
}

// SetSEzsignformfieldgroupLabel sets field value
func (o *EzsignformfieldgroupResponseCompound) SetSEzsignformfieldgroupLabel(v string) {
	o.SEzsignformfieldgroupLabel = v
}

// GetIEzsignformfieldgroupStep returns the IEzsignformfieldgroupStep field value
func (o *EzsignformfieldgroupResponseCompound) GetIEzsignformfieldgroupStep() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsignformfieldgroupStep
}

// GetIEzsignformfieldgroupStepOk returns a tuple with the IEzsignformfieldgroupStep field value
// and a boolean to check if the value has been set.
func (o *EzsignformfieldgroupResponseCompound) GetIEzsignformfieldgroupStepOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzsignformfieldgroupStep, true
}

// SetIEzsignformfieldgroupStep sets field value
func (o *EzsignformfieldgroupResponseCompound) SetIEzsignformfieldgroupStep(v int32) {
	o.IEzsignformfieldgroupStep = v
}

// GetSEzsignformfieldgroupDefaultvalue returns the SEzsignformfieldgroupDefaultvalue field value if set, zero value otherwise.
func (o *EzsignformfieldgroupResponseCompound) GetSEzsignformfieldgroupDefaultvalue() string {
	if o == nil || IsNil(o.SEzsignformfieldgroupDefaultvalue) {
		var ret string
		return ret
	}
	return *o.SEzsignformfieldgroupDefaultvalue
}

// GetSEzsignformfieldgroupDefaultvalueOk returns a tuple with the SEzsignformfieldgroupDefaultvalue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignformfieldgroupResponseCompound) GetSEzsignformfieldgroupDefaultvalueOk() (*string, bool) {
	if o == nil || IsNil(o.SEzsignformfieldgroupDefaultvalue) {
		return nil, false
	}
	return o.SEzsignformfieldgroupDefaultvalue, true
}

// HasSEzsignformfieldgroupDefaultvalue returns a boolean if a field has been set.
func (o *EzsignformfieldgroupResponseCompound) HasSEzsignformfieldgroupDefaultvalue() bool {
	if o != nil && !IsNil(o.SEzsignformfieldgroupDefaultvalue) {
		return true
	}

	return false
}

// SetSEzsignformfieldgroupDefaultvalue gets a reference to the given string and assigns it to the SEzsignformfieldgroupDefaultvalue field.
func (o *EzsignformfieldgroupResponseCompound) SetSEzsignformfieldgroupDefaultvalue(v string) {
	o.SEzsignformfieldgroupDefaultvalue = &v
}

// GetIEzsignformfieldgroupFilledmin returns the IEzsignformfieldgroupFilledmin field value
func (o *EzsignformfieldgroupResponseCompound) GetIEzsignformfieldgroupFilledmin() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsignformfieldgroupFilledmin
}

// GetIEzsignformfieldgroupFilledminOk returns a tuple with the IEzsignformfieldgroupFilledmin field value
// and a boolean to check if the value has been set.
func (o *EzsignformfieldgroupResponseCompound) GetIEzsignformfieldgroupFilledminOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzsignformfieldgroupFilledmin, true
}

// SetIEzsignformfieldgroupFilledmin sets field value
func (o *EzsignformfieldgroupResponseCompound) SetIEzsignformfieldgroupFilledmin(v int32) {
	o.IEzsignformfieldgroupFilledmin = v
}

// GetIEzsignformfieldgroupFilledmax returns the IEzsignformfieldgroupFilledmax field value
func (o *EzsignformfieldgroupResponseCompound) GetIEzsignformfieldgroupFilledmax() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsignformfieldgroupFilledmax
}

// GetIEzsignformfieldgroupFilledmaxOk returns a tuple with the IEzsignformfieldgroupFilledmax field value
// and a boolean to check if the value has been set.
func (o *EzsignformfieldgroupResponseCompound) GetIEzsignformfieldgroupFilledmaxOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzsignformfieldgroupFilledmax, true
}

// SetIEzsignformfieldgroupFilledmax sets field value
func (o *EzsignformfieldgroupResponseCompound) SetIEzsignformfieldgroupFilledmax(v int32) {
	o.IEzsignformfieldgroupFilledmax = v
}

// GetBEzsignformfieldgroupReadonly returns the BEzsignformfieldgroupReadonly field value
func (o *EzsignformfieldgroupResponseCompound) GetBEzsignformfieldgroupReadonly() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BEzsignformfieldgroupReadonly
}

// GetBEzsignformfieldgroupReadonlyOk returns a tuple with the BEzsignformfieldgroupReadonly field value
// and a boolean to check if the value has been set.
func (o *EzsignformfieldgroupResponseCompound) GetBEzsignformfieldgroupReadonlyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BEzsignformfieldgroupReadonly, true
}

// SetBEzsignformfieldgroupReadonly sets field value
func (o *EzsignformfieldgroupResponseCompound) SetBEzsignformfieldgroupReadonly(v bool) {
	o.BEzsignformfieldgroupReadonly = v
}

// GetIEzsignformfieldgroupMaxlength returns the IEzsignformfieldgroupMaxlength field value if set, zero value otherwise.
func (o *EzsignformfieldgroupResponseCompound) GetIEzsignformfieldgroupMaxlength() int32 {
	if o == nil || IsNil(o.IEzsignformfieldgroupMaxlength) {
		var ret int32
		return ret
	}
	return *o.IEzsignformfieldgroupMaxlength
}

// GetIEzsignformfieldgroupMaxlengthOk returns a tuple with the IEzsignformfieldgroupMaxlength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignformfieldgroupResponseCompound) GetIEzsignformfieldgroupMaxlengthOk() (*int32, bool) {
	if o == nil || IsNil(o.IEzsignformfieldgroupMaxlength) {
		return nil, false
	}
	return o.IEzsignformfieldgroupMaxlength, true
}

// HasIEzsignformfieldgroupMaxlength returns a boolean if a field has been set.
func (o *EzsignformfieldgroupResponseCompound) HasIEzsignformfieldgroupMaxlength() bool {
	if o != nil && !IsNil(o.IEzsignformfieldgroupMaxlength) {
		return true
	}

	return false
}

// SetIEzsignformfieldgroupMaxlength gets a reference to the given int32 and assigns it to the IEzsignformfieldgroupMaxlength field.
func (o *EzsignformfieldgroupResponseCompound) SetIEzsignformfieldgroupMaxlength(v int32) {
	o.IEzsignformfieldgroupMaxlength = &v
}

// GetBEzsignformfieldgroupEncrypted returns the BEzsignformfieldgroupEncrypted field value if set, zero value otherwise.
func (o *EzsignformfieldgroupResponseCompound) GetBEzsignformfieldgroupEncrypted() bool {
	if o == nil || IsNil(o.BEzsignformfieldgroupEncrypted) {
		var ret bool
		return ret
	}
	return *o.BEzsignformfieldgroupEncrypted
}

// GetBEzsignformfieldgroupEncryptedOk returns a tuple with the BEzsignformfieldgroupEncrypted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignformfieldgroupResponseCompound) GetBEzsignformfieldgroupEncryptedOk() (*bool, bool) {
	if o == nil || IsNil(o.BEzsignformfieldgroupEncrypted) {
		return nil, false
	}
	return o.BEzsignformfieldgroupEncrypted, true
}

// HasBEzsignformfieldgroupEncrypted returns a boolean if a field has been set.
func (o *EzsignformfieldgroupResponseCompound) HasBEzsignformfieldgroupEncrypted() bool {
	if o != nil && !IsNil(o.BEzsignformfieldgroupEncrypted) {
		return true
	}

	return false
}

// SetBEzsignformfieldgroupEncrypted gets a reference to the given bool and assigns it to the BEzsignformfieldgroupEncrypted field.
func (o *EzsignformfieldgroupResponseCompound) SetBEzsignformfieldgroupEncrypted(v bool) {
	o.BEzsignformfieldgroupEncrypted = &v
}

// GetEEzsignformfieldgroupTextvalidation returns the EEzsignformfieldgroupTextvalidation field value if set, zero value otherwise.
func (o *EzsignformfieldgroupResponseCompound) GetEEzsignformfieldgroupTextvalidation() EnumTextvalidation {
	if o == nil || IsNil(o.EEzsignformfieldgroupTextvalidation) {
		var ret EnumTextvalidation
		return ret
	}
	return *o.EEzsignformfieldgroupTextvalidation
}

// GetEEzsignformfieldgroupTextvalidationOk returns a tuple with the EEzsignformfieldgroupTextvalidation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignformfieldgroupResponseCompound) GetEEzsignformfieldgroupTextvalidationOk() (*EnumTextvalidation, bool) {
	if o == nil || IsNil(o.EEzsignformfieldgroupTextvalidation) {
		return nil, false
	}
	return o.EEzsignformfieldgroupTextvalidation, true
}

// HasEEzsignformfieldgroupTextvalidation returns a boolean if a field has been set.
func (o *EzsignformfieldgroupResponseCompound) HasEEzsignformfieldgroupTextvalidation() bool {
	if o != nil && !IsNil(o.EEzsignformfieldgroupTextvalidation) {
		return true
	}

	return false
}

// SetEEzsignformfieldgroupTextvalidation gets a reference to the given EnumTextvalidation and assigns it to the EEzsignformfieldgroupTextvalidation field.
func (o *EzsignformfieldgroupResponseCompound) SetEEzsignformfieldgroupTextvalidation(v EnumTextvalidation) {
	o.EEzsignformfieldgroupTextvalidation = &v
}

// GetSEzsignformfieldgroupRegexp returns the SEzsignformfieldgroupRegexp field value if set, zero value otherwise.
func (o *EzsignformfieldgroupResponseCompound) GetSEzsignformfieldgroupRegexp() string {
	if o == nil || IsNil(o.SEzsignformfieldgroupRegexp) {
		var ret string
		return ret
	}
	return *o.SEzsignformfieldgroupRegexp
}

// GetSEzsignformfieldgroupRegexpOk returns a tuple with the SEzsignformfieldgroupRegexp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignformfieldgroupResponseCompound) GetSEzsignformfieldgroupRegexpOk() (*string, bool) {
	if o == nil || IsNil(o.SEzsignformfieldgroupRegexp) {
		return nil, false
	}
	return o.SEzsignformfieldgroupRegexp, true
}

// HasSEzsignformfieldgroupRegexp returns a boolean if a field has been set.
func (o *EzsignformfieldgroupResponseCompound) HasSEzsignformfieldgroupRegexp() bool {
	if o != nil && !IsNil(o.SEzsignformfieldgroupRegexp) {
		return true
	}

	return false
}

// SetSEzsignformfieldgroupRegexp gets a reference to the given string and assigns it to the SEzsignformfieldgroupRegexp field.
func (o *EzsignformfieldgroupResponseCompound) SetSEzsignformfieldgroupRegexp(v string) {
	o.SEzsignformfieldgroupRegexp = &v
}

// GetTEzsignformfieldgroupTooltip returns the TEzsignformfieldgroupTooltip field value if set, zero value otherwise.
func (o *EzsignformfieldgroupResponseCompound) GetTEzsignformfieldgroupTooltip() string {
	if o == nil || IsNil(o.TEzsignformfieldgroupTooltip) {
		var ret string
		return ret
	}
	return *o.TEzsignformfieldgroupTooltip
}

// GetTEzsignformfieldgroupTooltipOk returns a tuple with the TEzsignformfieldgroupTooltip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignformfieldgroupResponseCompound) GetTEzsignformfieldgroupTooltipOk() (*string, bool) {
	if o == nil || IsNil(o.TEzsignformfieldgroupTooltip) {
		return nil, false
	}
	return o.TEzsignformfieldgroupTooltip, true
}

// HasTEzsignformfieldgroupTooltip returns a boolean if a field has been set.
func (o *EzsignformfieldgroupResponseCompound) HasTEzsignformfieldgroupTooltip() bool {
	if o != nil && !IsNil(o.TEzsignformfieldgroupTooltip) {
		return true
	}

	return false
}

// SetTEzsignformfieldgroupTooltip gets a reference to the given string and assigns it to the TEzsignformfieldgroupTooltip field.
func (o *EzsignformfieldgroupResponseCompound) SetTEzsignformfieldgroupTooltip(v string) {
	o.TEzsignformfieldgroupTooltip = &v
}

// GetEEzsignformfieldgroupTooltipposition returns the EEzsignformfieldgroupTooltipposition field value if set, zero value otherwise.
func (o *EzsignformfieldgroupResponseCompound) GetEEzsignformfieldgroupTooltipposition() FieldEEzsignformfieldgroupTooltipposition {
	if o == nil || IsNil(o.EEzsignformfieldgroupTooltipposition) {
		var ret FieldEEzsignformfieldgroupTooltipposition
		return ret
	}
	return *o.EEzsignformfieldgroupTooltipposition
}

// GetEEzsignformfieldgroupTooltippositionOk returns a tuple with the EEzsignformfieldgroupTooltipposition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignformfieldgroupResponseCompound) GetEEzsignformfieldgroupTooltippositionOk() (*FieldEEzsignformfieldgroupTooltipposition, bool) {
	if o == nil || IsNil(o.EEzsignformfieldgroupTooltipposition) {
		return nil, false
	}
	return o.EEzsignformfieldgroupTooltipposition, true
}

// HasEEzsignformfieldgroupTooltipposition returns a boolean if a field has been set.
func (o *EzsignformfieldgroupResponseCompound) HasEEzsignformfieldgroupTooltipposition() bool {
	if o != nil && !IsNil(o.EEzsignformfieldgroupTooltipposition) {
		return true
	}

	return false
}

// SetEEzsignformfieldgroupTooltipposition gets a reference to the given FieldEEzsignformfieldgroupTooltipposition and assigns it to the EEzsignformfieldgroupTooltipposition field.
func (o *EzsignformfieldgroupResponseCompound) SetEEzsignformfieldgroupTooltipposition(v FieldEEzsignformfieldgroupTooltipposition) {
	o.EEzsignformfieldgroupTooltipposition = &v
}

// GetAObjEzsignformfield returns the AObjEzsignformfield field value
func (o *EzsignformfieldgroupResponseCompound) GetAObjEzsignformfield() []EzsignformfieldResponseCompound {
	if o == nil {
		var ret []EzsignformfieldResponseCompound
		return ret
	}

	return o.AObjEzsignformfield
}

// GetAObjEzsignformfieldOk returns a tuple with the AObjEzsignformfield field value
// and a boolean to check if the value has been set.
func (o *EzsignformfieldgroupResponseCompound) GetAObjEzsignformfieldOk() ([]EzsignformfieldResponseCompound, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjEzsignformfield, true
}

// SetAObjEzsignformfield sets field value
func (o *EzsignformfieldgroupResponseCompound) SetAObjEzsignformfield(v []EzsignformfieldResponseCompound) {
	o.AObjEzsignformfield = v
}

// GetAObjDropdownElement returns the AObjDropdownElement field value if set, zero value otherwise.
func (o *EzsignformfieldgroupResponseCompound) GetAObjDropdownElement() []CustomDropdownElementResponseCompound {
	if o == nil || IsNil(o.AObjDropdownElement) {
		var ret []CustomDropdownElementResponseCompound
		return ret
	}
	return o.AObjDropdownElement
}

// GetAObjDropdownElementOk returns a tuple with the AObjDropdownElement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignformfieldgroupResponseCompound) GetAObjDropdownElementOk() ([]CustomDropdownElementResponseCompound, bool) {
	if o == nil || IsNil(o.AObjDropdownElement) {
		return nil, false
	}
	return o.AObjDropdownElement, true
}

// HasAObjDropdownElement returns a boolean if a field has been set.
func (o *EzsignformfieldgroupResponseCompound) HasAObjDropdownElement() bool {
	if o != nil && !IsNil(o.AObjDropdownElement) {
		return true
	}

	return false
}

// SetAObjDropdownElement gets a reference to the given []CustomDropdownElementResponseCompound and assigns it to the AObjDropdownElement field.
func (o *EzsignformfieldgroupResponseCompound) SetAObjDropdownElement(v []CustomDropdownElementResponseCompound) {
	o.AObjDropdownElement = v
}

// GetAObjEzsignformfieldgroupsigner returns the AObjEzsignformfieldgroupsigner field value
func (o *EzsignformfieldgroupResponseCompound) GetAObjEzsignformfieldgroupsigner() []EzsignformfieldgroupsignerResponseCompound {
	if o == nil {
		var ret []EzsignformfieldgroupsignerResponseCompound
		return ret
	}

	return o.AObjEzsignformfieldgroupsigner
}

// GetAObjEzsignformfieldgroupsignerOk returns a tuple with the AObjEzsignformfieldgroupsigner field value
// and a boolean to check if the value has been set.
func (o *EzsignformfieldgroupResponseCompound) GetAObjEzsignformfieldgroupsignerOk() ([]EzsignformfieldgroupsignerResponseCompound, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjEzsignformfieldgroupsigner, true
}

// SetAObjEzsignformfieldgroupsigner sets field value
func (o *EzsignformfieldgroupResponseCompound) SetAObjEzsignformfieldgroupsigner(v []EzsignformfieldgroupsignerResponseCompound) {
	o.AObjEzsignformfieldgroupsigner = v
}

func (o EzsignformfieldgroupResponseCompound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignformfieldgroupResponseCompound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiEzsignformfieldgroupID"] = o.PkiEzsignformfieldgroupID
	toSerialize["fkiEzsigndocumentID"] = o.FkiEzsigndocumentID
	toSerialize["eEzsignformfieldgroupType"] = o.EEzsignformfieldgroupType
	toSerialize["eEzsignformfieldgroupSignerrequirement"] = o.EEzsignformfieldgroupSignerrequirement
	toSerialize["sEzsignformfieldgroupLabel"] = o.SEzsignformfieldgroupLabel
	toSerialize["iEzsignformfieldgroupStep"] = o.IEzsignformfieldgroupStep
	if !IsNil(o.SEzsignformfieldgroupDefaultvalue) {
		toSerialize["sEzsignformfieldgroupDefaultvalue"] = o.SEzsignformfieldgroupDefaultvalue
	}
	toSerialize["iEzsignformfieldgroupFilledmin"] = o.IEzsignformfieldgroupFilledmin
	toSerialize["iEzsignformfieldgroupFilledmax"] = o.IEzsignformfieldgroupFilledmax
	toSerialize["bEzsignformfieldgroupReadonly"] = o.BEzsignformfieldgroupReadonly
	if !IsNil(o.IEzsignformfieldgroupMaxlength) {
		toSerialize["iEzsignformfieldgroupMaxlength"] = o.IEzsignformfieldgroupMaxlength
	}
	if !IsNil(o.BEzsignformfieldgroupEncrypted) {
		toSerialize["bEzsignformfieldgroupEncrypted"] = o.BEzsignformfieldgroupEncrypted
	}
	if !IsNil(o.EEzsignformfieldgroupTextvalidation) {
		toSerialize["eEzsignformfieldgroupTextvalidation"] = o.EEzsignformfieldgroupTextvalidation
	}
	if !IsNil(o.SEzsignformfieldgroupRegexp) {
		toSerialize["sEzsignformfieldgroupRegexp"] = o.SEzsignformfieldgroupRegexp
	}
	if !IsNil(o.TEzsignformfieldgroupTooltip) {
		toSerialize["tEzsignformfieldgroupTooltip"] = o.TEzsignformfieldgroupTooltip
	}
	if !IsNil(o.EEzsignformfieldgroupTooltipposition) {
		toSerialize["eEzsignformfieldgroupTooltipposition"] = o.EEzsignformfieldgroupTooltipposition
	}
	toSerialize["a_objEzsignformfield"] = o.AObjEzsignformfield
	if !IsNil(o.AObjDropdownElement) {
		toSerialize["a_objDropdownElement"] = o.AObjDropdownElement
	}
	toSerialize["a_objEzsignformfieldgroupsigner"] = o.AObjEzsignformfieldgroupsigner
	return toSerialize, nil
}

type NullableEzsignformfieldgroupResponseCompound struct {
	value *EzsignformfieldgroupResponseCompound
	isSet bool
}

func (v NullableEzsignformfieldgroupResponseCompound) Get() *EzsignformfieldgroupResponseCompound {
	return v.value
}

func (v *NullableEzsignformfieldgroupResponseCompound) Set(val *EzsignformfieldgroupResponseCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignformfieldgroupResponseCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignformfieldgroupResponseCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignformfieldgroupResponseCompound(val *EzsignformfieldgroupResponseCompound) *NullableEzsignformfieldgroupResponseCompound {
	return &NullableEzsignformfieldgroupResponseCompound{value: val, isSet: true}
}

func (v NullableEzsignformfieldgroupResponseCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignformfieldgroupResponseCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


