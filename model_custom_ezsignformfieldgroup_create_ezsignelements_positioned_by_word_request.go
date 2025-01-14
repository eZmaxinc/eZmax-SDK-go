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

// checks if the CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest{}

// CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest An Ezsignformfieldgroup Object in the context of a createEzsignelementsPositionedByWord path
type CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest struct {
	// The unique ID of the Ezsignformfieldgroup
	PkiEzsignformfieldgroupID *int32 `json:"pkiEzsignformfieldgroupID,omitempty"`
	// The unique ID of the Ezsigndocument
	FkiEzsigndocumentID int32 `json:"fkiEzsigndocumentID"`
	EEzsignformfieldgroupType FieldEEzsignformfieldgroupType `json:"eEzsignformfieldgroupType"`
	// Deprecated
	EEzsignformfieldgroupSignerrequirement *FieldEEzsignformfieldgroupSignerrequirement `json:"eEzsignformfieldgroupSignerrequirement,omitempty"`
	// The Label for the Ezsignformfieldgroup
	SEzsignformfieldgroupLabel string `json:"sEzsignformfieldgroupLabel"`
	// The step when the Ezsignsigner will be invited to fill the form fields
	IEzsignformfieldgroupStep int32 `json:"iEzsignformfieldgroupStep"`
	// The default value for the Ezsignformfieldgroup  You can use the codes below and they will be replaced at signature time.    | Code | Description | Example | | ------------------------- | ------------ | ------------ | | {sUserFirstname} | The first name of the contact | John | | {sUserLastname} | The last name of the contact | Doe | | {sUserJobtitle} | The job title | Sales Representative | | {sCompany} | Company name | eZmax Solutions Inc. | | {sEmailAddress} | The email address | email@example.com | | {sPhoneE164} | A phone number in E.164 Format | +15149901516 | | {sPhoneE164Cell} | A phone number in E.164 Format | +15149901516 |
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
	// A regular expression to indicate what values are acceptable for the Ezsignformfieldgroup.  This can only be set if eEzsignformfieldgroupType is **Text** or **Textarea**
	SEzsignformfieldgroupRegexp *string `json:"sEzsignformfieldgroupRegexp,omitempty" validate:"regexp=^\\\\^.*\\\\$$|^$"`
	// Description of validation rule. Show by signatory.
	SEzsignformfieldgroupTextvalidationcustommessage *string `json:"sEzsignformfieldgroupTextvalidationcustommessage,omitempty"`
	// A tooltip that will be presented to Ezsignsigner about the Ezsignformfieldgroup
	TEzsignformfieldgroupTooltip *string `json:"tEzsignformfieldgroupTooltip,omitempty"`
	EEzsignformfieldgroupTooltipposition *FieldEEzsignformfieldgroupTooltipposition `json:"eEzsignformfieldgroupTooltipposition,omitempty"`
	EEzsignformfieldgroupTextvalidation *EnumTextvalidation `json:"eEzsignformfieldgroupTextvalidation,omitempty"`
	AObjEzsignformfieldgroupsigner []EzsignformfieldgroupsignerRequestCompound `json:"a_objEzsignformfieldgroupsigner"`
	AObjDropdownElement []CustomDropdownElementRequestCompound `json:"a_objDropdownElement,omitempty"`
	AObjEzsignformfield []EzsignformfieldRequestCompound `json:"a_objEzsignformfield"`
	ObjCreateezsignelementspositionedbyword CustomCreateEzsignelementsPositionedByWordRequest `json:"objCreateezsignelementspositionedbyword"`
}

type _CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest

// NewCustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest instantiates a new CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest(fkiEzsigndocumentID int32, eEzsignformfieldgroupType FieldEEzsignformfieldgroupType, sEzsignformfieldgroupLabel string, iEzsignformfieldgroupStep int32, iEzsignformfieldgroupFilledmin int32, iEzsignformfieldgroupFilledmax int32, bEzsignformfieldgroupReadonly bool, aObjEzsignformfieldgroupsigner []EzsignformfieldgroupsignerRequestCompound, aObjEzsignformfield []EzsignformfieldRequestCompound, objCreateezsignelementspositionedbyword CustomCreateEzsignelementsPositionedByWordRequest) *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest {
	this := CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest{}
	this.FkiEzsigndocumentID = fkiEzsigndocumentID
	this.EEzsignformfieldgroupType = eEzsignformfieldgroupType
	this.SEzsignformfieldgroupLabel = sEzsignformfieldgroupLabel
	this.IEzsignformfieldgroupStep = iEzsignformfieldgroupStep
	this.IEzsignformfieldgroupFilledmin = iEzsignformfieldgroupFilledmin
	this.IEzsignformfieldgroupFilledmax = iEzsignformfieldgroupFilledmax
	this.BEzsignformfieldgroupReadonly = bEzsignformfieldgroupReadonly
	this.AObjEzsignformfieldgroupsigner = aObjEzsignformfieldgroupsigner
	this.AObjEzsignformfield = aObjEzsignformfield
	this.ObjCreateezsignelementspositionedbyword = objCreateezsignelementspositionedbyword
	return &this
}

// NewCustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequestWithDefaults instantiates a new CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequestWithDefaults() *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest {
	this := CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest{}
	return &this
}

// GetPkiEzsignformfieldgroupID returns the PkiEzsignformfieldgroupID field value if set, zero value otherwise.
func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) GetPkiEzsignformfieldgroupID() int32 {
	if o == nil || IsNil(o.PkiEzsignformfieldgroupID) {
		var ret int32
		return ret
	}
	return *o.PkiEzsignformfieldgroupID
}

// GetPkiEzsignformfieldgroupIDOk returns a tuple with the PkiEzsignformfieldgroupID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) GetPkiEzsignformfieldgroupIDOk() (*int32, bool) {
	if o == nil || IsNil(o.PkiEzsignformfieldgroupID) {
		return nil, false
	}
	return o.PkiEzsignformfieldgroupID, true
}

// HasPkiEzsignformfieldgroupID returns a boolean if a field has been set.
func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) HasPkiEzsignformfieldgroupID() bool {
	if o != nil && !IsNil(o.PkiEzsignformfieldgroupID) {
		return true
	}

	return false
}

// SetPkiEzsignformfieldgroupID gets a reference to the given int32 and assigns it to the PkiEzsignformfieldgroupID field.
func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) SetPkiEzsignformfieldgroupID(v int32) {
	o.PkiEzsignformfieldgroupID = &v
}

// GetFkiEzsigndocumentID returns the FkiEzsigndocumentID field value
func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) GetFkiEzsigndocumentID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzsigndocumentID
}

// GetFkiEzsigndocumentIDOk returns a tuple with the FkiEzsigndocumentID field value
// and a boolean to check if the value has been set.
func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) GetFkiEzsigndocumentIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiEzsigndocumentID, true
}

// SetFkiEzsigndocumentID sets field value
func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) SetFkiEzsigndocumentID(v int32) {
	o.FkiEzsigndocumentID = v
}

// GetEEzsignformfieldgroupType returns the EEzsignformfieldgroupType field value
func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) GetEEzsignformfieldgroupType() FieldEEzsignformfieldgroupType {
	if o == nil {
		var ret FieldEEzsignformfieldgroupType
		return ret
	}

	return o.EEzsignformfieldgroupType
}

// GetEEzsignformfieldgroupTypeOk returns a tuple with the EEzsignformfieldgroupType field value
// and a boolean to check if the value has been set.
func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) GetEEzsignformfieldgroupTypeOk() (*FieldEEzsignformfieldgroupType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EEzsignformfieldgroupType, true
}

// SetEEzsignformfieldgroupType sets field value
func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) SetEEzsignformfieldgroupType(v FieldEEzsignformfieldgroupType) {
	o.EEzsignformfieldgroupType = v
}

// GetEEzsignformfieldgroupSignerrequirement returns the EEzsignformfieldgroupSignerrequirement field value if set, zero value otherwise.
// Deprecated
func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) GetEEzsignformfieldgroupSignerrequirement() FieldEEzsignformfieldgroupSignerrequirement {
	if o == nil || IsNil(o.EEzsignformfieldgroupSignerrequirement) {
		var ret FieldEEzsignformfieldgroupSignerrequirement
		return ret
	}
	return *o.EEzsignformfieldgroupSignerrequirement
}

// GetEEzsignformfieldgroupSignerrequirementOk returns a tuple with the EEzsignformfieldgroupSignerrequirement field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) GetEEzsignformfieldgroupSignerrequirementOk() (*FieldEEzsignformfieldgroupSignerrequirement, bool) {
	if o == nil || IsNil(o.EEzsignformfieldgroupSignerrequirement) {
		return nil, false
	}
	return o.EEzsignformfieldgroupSignerrequirement, true
}

// HasEEzsignformfieldgroupSignerrequirement returns a boolean if a field has been set.
func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) HasEEzsignformfieldgroupSignerrequirement() bool {
	if o != nil && !IsNil(o.EEzsignformfieldgroupSignerrequirement) {
		return true
	}

	return false
}

// SetEEzsignformfieldgroupSignerrequirement gets a reference to the given FieldEEzsignformfieldgroupSignerrequirement and assigns it to the EEzsignformfieldgroupSignerrequirement field.
// Deprecated
func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) SetEEzsignformfieldgroupSignerrequirement(v FieldEEzsignformfieldgroupSignerrequirement) {
	o.EEzsignformfieldgroupSignerrequirement = &v
}

// GetSEzsignformfieldgroupLabel returns the SEzsignformfieldgroupLabel field value
func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) GetSEzsignformfieldgroupLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzsignformfieldgroupLabel
}

// GetSEzsignformfieldgroupLabelOk returns a tuple with the SEzsignformfieldgroupLabel field value
// and a boolean to check if the value has been set.
func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) GetSEzsignformfieldgroupLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SEzsignformfieldgroupLabel, true
}

// SetSEzsignformfieldgroupLabel sets field value
func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) SetSEzsignformfieldgroupLabel(v string) {
	o.SEzsignformfieldgroupLabel = v
}

// GetIEzsignformfieldgroupStep returns the IEzsignformfieldgroupStep field value
func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) GetIEzsignformfieldgroupStep() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsignformfieldgroupStep
}

// GetIEzsignformfieldgroupStepOk returns a tuple with the IEzsignformfieldgroupStep field value
// and a boolean to check if the value has been set.
func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) GetIEzsignformfieldgroupStepOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzsignformfieldgroupStep, true
}

// SetIEzsignformfieldgroupStep sets field value
func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) SetIEzsignformfieldgroupStep(v int32) {
	o.IEzsignformfieldgroupStep = v
}

// GetSEzsignformfieldgroupDefaultvalue returns the SEzsignformfieldgroupDefaultvalue field value if set, zero value otherwise.
func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) GetSEzsignformfieldgroupDefaultvalue() string {
	if o == nil || IsNil(o.SEzsignformfieldgroupDefaultvalue) {
		var ret string
		return ret
	}
	return *o.SEzsignformfieldgroupDefaultvalue
}

// GetSEzsignformfieldgroupDefaultvalueOk returns a tuple with the SEzsignformfieldgroupDefaultvalue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) GetSEzsignformfieldgroupDefaultvalueOk() (*string, bool) {
	if o == nil || IsNil(o.SEzsignformfieldgroupDefaultvalue) {
		return nil, false
	}
	return o.SEzsignformfieldgroupDefaultvalue, true
}

// HasSEzsignformfieldgroupDefaultvalue returns a boolean if a field has been set.
func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) HasSEzsignformfieldgroupDefaultvalue() bool {
	if o != nil && !IsNil(o.SEzsignformfieldgroupDefaultvalue) {
		return true
	}

	return false
}

// SetSEzsignformfieldgroupDefaultvalue gets a reference to the given string and assigns it to the SEzsignformfieldgroupDefaultvalue field.
func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) SetSEzsignformfieldgroupDefaultvalue(v string) {
	o.SEzsignformfieldgroupDefaultvalue = &v
}

// GetIEzsignformfieldgroupFilledmin returns the IEzsignformfieldgroupFilledmin field value
func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) GetIEzsignformfieldgroupFilledmin() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsignformfieldgroupFilledmin
}

// GetIEzsignformfieldgroupFilledminOk returns a tuple with the IEzsignformfieldgroupFilledmin field value
// and a boolean to check if the value has been set.
func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) GetIEzsignformfieldgroupFilledminOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzsignformfieldgroupFilledmin, true
}

// SetIEzsignformfieldgroupFilledmin sets field value
func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) SetIEzsignformfieldgroupFilledmin(v int32) {
	o.IEzsignformfieldgroupFilledmin = v
}

// GetIEzsignformfieldgroupFilledmax returns the IEzsignformfieldgroupFilledmax field value
func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) GetIEzsignformfieldgroupFilledmax() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsignformfieldgroupFilledmax
}

// GetIEzsignformfieldgroupFilledmaxOk returns a tuple with the IEzsignformfieldgroupFilledmax field value
// and a boolean to check if the value has been set.
func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) GetIEzsignformfieldgroupFilledmaxOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzsignformfieldgroupFilledmax, true
}

// SetIEzsignformfieldgroupFilledmax sets field value
func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) SetIEzsignformfieldgroupFilledmax(v int32) {
	o.IEzsignformfieldgroupFilledmax = v
}

// GetBEzsignformfieldgroupReadonly returns the BEzsignformfieldgroupReadonly field value
func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) GetBEzsignformfieldgroupReadonly() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BEzsignformfieldgroupReadonly
}

// GetBEzsignformfieldgroupReadonlyOk returns a tuple with the BEzsignformfieldgroupReadonly field value
// and a boolean to check if the value has been set.
func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) GetBEzsignformfieldgroupReadonlyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BEzsignformfieldgroupReadonly, true
}

// SetBEzsignformfieldgroupReadonly sets field value
func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) SetBEzsignformfieldgroupReadonly(v bool) {
	o.BEzsignformfieldgroupReadonly = v
}

// GetIEzsignformfieldgroupMaxlength returns the IEzsignformfieldgroupMaxlength field value if set, zero value otherwise.
func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) GetIEzsignformfieldgroupMaxlength() int32 {
	if o == nil || IsNil(o.IEzsignformfieldgroupMaxlength) {
		var ret int32
		return ret
	}
	return *o.IEzsignformfieldgroupMaxlength
}

// GetIEzsignformfieldgroupMaxlengthOk returns a tuple with the IEzsignformfieldgroupMaxlength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) GetIEzsignformfieldgroupMaxlengthOk() (*int32, bool) {
	if o == nil || IsNil(o.IEzsignformfieldgroupMaxlength) {
		return nil, false
	}
	return o.IEzsignformfieldgroupMaxlength, true
}

// HasIEzsignformfieldgroupMaxlength returns a boolean if a field has been set.
func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) HasIEzsignformfieldgroupMaxlength() bool {
	if o != nil && !IsNil(o.IEzsignformfieldgroupMaxlength) {
		return true
	}

	return false
}

// SetIEzsignformfieldgroupMaxlength gets a reference to the given int32 and assigns it to the IEzsignformfieldgroupMaxlength field.
func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) SetIEzsignformfieldgroupMaxlength(v int32) {
	o.IEzsignformfieldgroupMaxlength = &v
}

// GetBEzsignformfieldgroupEncrypted returns the BEzsignformfieldgroupEncrypted field value if set, zero value otherwise.
func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) GetBEzsignformfieldgroupEncrypted() bool {
	if o == nil || IsNil(o.BEzsignformfieldgroupEncrypted) {
		var ret bool
		return ret
	}
	return *o.BEzsignformfieldgroupEncrypted
}

// GetBEzsignformfieldgroupEncryptedOk returns a tuple with the BEzsignformfieldgroupEncrypted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) GetBEzsignformfieldgroupEncryptedOk() (*bool, bool) {
	if o == nil || IsNil(o.BEzsignformfieldgroupEncrypted) {
		return nil, false
	}
	return o.BEzsignformfieldgroupEncrypted, true
}

// HasBEzsignformfieldgroupEncrypted returns a boolean if a field has been set.
func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) HasBEzsignformfieldgroupEncrypted() bool {
	if o != nil && !IsNil(o.BEzsignformfieldgroupEncrypted) {
		return true
	}

	return false
}

// SetBEzsignformfieldgroupEncrypted gets a reference to the given bool and assigns it to the BEzsignformfieldgroupEncrypted field.
func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) SetBEzsignformfieldgroupEncrypted(v bool) {
	o.BEzsignformfieldgroupEncrypted = &v
}

// GetSEzsignformfieldgroupRegexp returns the SEzsignformfieldgroupRegexp field value if set, zero value otherwise.
func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) GetSEzsignformfieldgroupRegexp() string {
	if o == nil || IsNil(o.SEzsignformfieldgroupRegexp) {
		var ret string
		return ret
	}
	return *o.SEzsignformfieldgroupRegexp
}

// GetSEzsignformfieldgroupRegexpOk returns a tuple with the SEzsignformfieldgroupRegexp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) GetSEzsignformfieldgroupRegexpOk() (*string, bool) {
	if o == nil || IsNil(o.SEzsignformfieldgroupRegexp) {
		return nil, false
	}
	return o.SEzsignformfieldgroupRegexp, true
}

// HasSEzsignformfieldgroupRegexp returns a boolean if a field has been set.
func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) HasSEzsignformfieldgroupRegexp() bool {
	if o != nil && !IsNil(o.SEzsignformfieldgroupRegexp) {
		return true
	}

	return false
}

// SetSEzsignformfieldgroupRegexp gets a reference to the given string and assigns it to the SEzsignformfieldgroupRegexp field.
func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) SetSEzsignformfieldgroupRegexp(v string) {
	o.SEzsignformfieldgroupRegexp = &v
}

// GetSEzsignformfieldgroupTextvalidationcustommessage returns the SEzsignformfieldgroupTextvalidationcustommessage field value if set, zero value otherwise.
func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) GetSEzsignformfieldgroupTextvalidationcustommessage() string {
	if o == nil || IsNil(o.SEzsignformfieldgroupTextvalidationcustommessage) {
		var ret string
		return ret
	}
	return *o.SEzsignformfieldgroupTextvalidationcustommessage
}

// GetSEzsignformfieldgroupTextvalidationcustommessageOk returns a tuple with the SEzsignformfieldgroupTextvalidationcustommessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) GetSEzsignformfieldgroupTextvalidationcustommessageOk() (*string, bool) {
	if o == nil || IsNil(o.SEzsignformfieldgroupTextvalidationcustommessage) {
		return nil, false
	}
	return o.SEzsignformfieldgroupTextvalidationcustommessage, true
}

// HasSEzsignformfieldgroupTextvalidationcustommessage returns a boolean if a field has been set.
func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) HasSEzsignformfieldgroupTextvalidationcustommessage() bool {
	if o != nil && !IsNil(o.SEzsignformfieldgroupTextvalidationcustommessage) {
		return true
	}

	return false
}

// SetSEzsignformfieldgroupTextvalidationcustommessage gets a reference to the given string and assigns it to the SEzsignformfieldgroupTextvalidationcustommessage field.
func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) SetSEzsignformfieldgroupTextvalidationcustommessage(v string) {
	o.SEzsignformfieldgroupTextvalidationcustommessage = &v
}

// GetTEzsignformfieldgroupTooltip returns the TEzsignformfieldgroupTooltip field value if set, zero value otherwise.
func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) GetTEzsignformfieldgroupTooltip() string {
	if o == nil || IsNil(o.TEzsignformfieldgroupTooltip) {
		var ret string
		return ret
	}
	return *o.TEzsignformfieldgroupTooltip
}

// GetTEzsignformfieldgroupTooltipOk returns a tuple with the TEzsignformfieldgroupTooltip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) GetTEzsignformfieldgroupTooltipOk() (*string, bool) {
	if o == nil || IsNil(o.TEzsignformfieldgroupTooltip) {
		return nil, false
	}
	return o.TEzsignformfieldgroupTooltip, true
}

// HasTEzsignformfieldgroupTooltip returns a boolean if a field has been set.
func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) HasTEzsignformfieldgroupTooltip() bool {
	if o != nil && !IsNil(o.TEzsignformfieldgroupTooltip) {
		return true
	}

	return false
}

// SetTEzsignformfieldgroupTooltip gets a reference to the given string and assigns it to the TEzsignformfieldgroupTooltip field.
func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) SetTEzsignformfieldgroupTooltip(v string) {
	o.TEzsignformfieldgroupTooltip = &v
}

// GetEEzsignformfieldgroupTooltipposition returns the EEzsignformfieldgroupTooltipposition field value if set, zero value otherwise.
func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) GetEEzsignformfieldgroupTooltipposition() FieldEEzsignformfieldgroupTooltipposition {
	if o == nil || IsNil(o.EEzsignformfieldgroupTooltipposition) {
		var ret FieldEEzsignformfieldgroupTooltipposition
		return ret
	}
	return *o.EEzsignformfieldgroupTooltipposition
}

// GetEEzsignformfieldgroupTooltippositionOk returns a tuple with the EEzsignformfieldgroupTooltipposition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) GetEEzsignformfieldgroupTooltippositionOk() (*FieldEEzsignformfieldgroupTooltipposition, bool) {
	if o == nil || IsNil(o.EEzsignformfieldgroupTooltipposition) {
		return nil, false
	}
	return o.EEzsignformfieldgroupTooltipposition, true
}

// HasEEzsignformfieldgroupTooltipposition returns a boolean if a field has been set.
func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) HasEEzsignformfieldgroupTooltipposition() bool {
	if o != nil && !IsNil(o.EEzsignformfieldgroupTooltipposition) {
		return true
	}

	return false
}

// SetEEzsignformfieldgroupTooltipposition gets a reference to the given FieldEEzsignformfieldgroupTooltipposition and assigns it to the EEzsignformfieldgroupTooltipposition field.
func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) SetEEzsignformfieldgroupTooltipposition(v FieldEEzsignformfieldgroupTooltipposition) {
	o.EEzsignformfieldgroupTooltipposition = &v
}

// GetEEzsignformfieldgroupTextvalidation returns the EEzsignformfieldgroupTextvalidation field value if set, zero value otherwise.
func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) GetEEzsignformfieldgroupTextvalidation() EnumTextvalidation {
	if o == nil || IsNil(o.EEzsignformfieldgroupTextvalidation) {
		var ret EnumTextvalidation
		return ret
	}
	return *o.EEzsignformfieldgroupTextvalidation
}

// GetEEzsignformfieldgroupTextvalidationOk returns a tuple with the EEzsignformfieldgroupTextvalidation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) GetEEzsignformfieldgroupTextvalidationOk() (*EnumTextvalidation, bool) {
	if o == nil || IsNil(o.EEzsignformfieldgroupTextvalidation) {
		return nil, false
	}
	return o.EEzsignformfieldgroupTextvalidation, true
}

// HasEEzsignformfieldgroupTextvalidation returns a boolean if a field has been set.
func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) HasEEzsignformfieldgroupTextvalidation() bool {
	if o != nil && !IsNil(o.EEzsignformfieldgroupTextvalidation) {
		return true
	}

	return false
}

// SetEEzsignformfieldgroupTextvalidation gets a reference to the given EnumTextvalidation and assigns it to the EEzsignformfieldgroupTextvalidation field.
func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) SetEEzsignformfieldgroupTextvalidation(v EnumTextvalidation) {
	o.EEzsignformfieldgroupTextvalidation = &v
}

// GetAObjEzsignformfieldgroupsigner returns the AObjEzsignformfieldgroupsigner field value
func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) GetAObjEzsignformfieldgroupsigner() []EzsignformfieldgroupsignerRequestCompound {
	if o == nil {
		var ret []EzsignformfieldgroupsignerRequestCompound
		return ret
	}

	return o.AObjEzsignformfieldgroupsigner
}

// GetAObjEzsignformfieldgroupsignerOk returns a tuple with the AObjEzsignformfieldgroupsigner field value
// and a boolean to check if the value has been set.
func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) GetAObjEzsignformfieldgroupsignerOk() ([]EzsignformfieldgroupsignerRequestCompound, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjEzsignformfieldgroupsigner, true
}

// SetAObjEzsignformfieldgroupsigner sets field value
func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) SetAObjEzsignformfieldgroupsigner(v []EzsignformfieldgroupsignerRequestCompound) {
	o.AObjEzsignformfieldgroupsigner = v
}

// GetAObjDropdownElement returns the AObjDropdownElement field value if set, zero value otherwise.
func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) GetAObjDropdownElement() []CustomDropdownElementRequestCompound {
	if o == nil || IsNil(o.AObjDropdownElement) {
		var ret []CustomDropdownElementRequestCompound
		return ret
	}
	return o.AObjDropdownElement
}

// GetAObjDropdownElementOk returns a tuple with the AObjDropdownElement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) GetAObjDropdownElementOk() ([]CustomDropdownElementRequestCompound, bool) {
	if o == nil || IsNil(o.AObjDropdownElement) {
		return nil, false
	}
	return o.AObjDropdownElement, true
}

// HasAObjDropdownElement returns a boolean if a field has been set.
func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) HasAObjDropdownElement() bool {
	if o != nil && !IsNil(o.AObjDropdownElement) {
		return true
	}

	return false
}

// SetAObjDropdownElement gets a reference to the given []CustomDropdownElementRequestCompound and assigns it to the AObjDropdownElement field.
func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) SetAObjDropdownElement(v []CustomDropdownElementRequestCompound) {
	o.AObjDropdownElement = v
}

// GetAObjEzsignformfield returns the AObjEzsignformfield field value
func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) GetAObjEzsignformfield() []EzsignformfieldRequestCompound {
	if o == nil {
		var ret []EzsignformfieldRequestCompound
		return ret
	}

	return o.AObjEzsignformfield
}

// GetAObjEzsignformfieldOk returns a tuple with the AObjEzsignformfield field value
// and a boolean to check if the value has been set.
func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) GetAObjEzsignformfieldOk() ([]EzsignformfieldRequestCompound, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjEzsignformfield, true
}

// SetAObjEzsignformfield sets field value
func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) SetAObjEzsignformfield(v []EzsignformfieldRequestCompound) {
	o.AObjEzsignformfield = v
}

// GetObjCreateezsignelementspositionedbyword returns the ObjCreateezsignelementspositionedbyword field value
func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) GetObjCreateezsignelementspositionedbyword() CustomCreateEzsignelementsPositionedByWordRequest {
	if o == nil {
		var ret CustomCreateEzsignelementsPositionedByWordRequest
		return ret
	}

	return o.ObjCreateezsignelementspositionedbyword
}

// GetObjCreateezsignelementspositionedbywordOk returns a tuple with the ObjCreateezsignelementspositionedbyword field value
// and a boolean to check if the value has been set.
func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) GetObjCreateezsignelementspositionedbywordOk() (*CustomCreateEzsignelementsPositionedByWordRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjCreateezsignelementspositionedbyword, true
}

// SetObjCreateezsignelementspositionedbyword sets field value
func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) SetObjCreateezsignelementspositionedbyword(v CustomCreateEzsignelementsPositionedByWordRequest) {
	o.ObjCreateezsignelementspositionedbyword = v
}

func (o CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PkiEzsignformfieldgroupID) {
		toSerialize["pkiEzsignformfieldgroupID"] = o.PkiEzsignformfieldgroupID
	}
	toSerialize["fkiEzsigndocumentID"] = o.FkiEzsigndocumentID
	toSerialize["eEzsignformfieldgroupType"] = o.EEzsignformfieldgroupType
	if !IsNil(o.EEzsignformfieldgroupSignerrequirement) {
		toSerialize["eEzsignformfieldgroupSignerrequirement"] = o.EEzsignformfieldgroupSignerrequirement
	}
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
	if !IsNil(o.SEzsignformfieldgroupRegexp) {
		toSerialize["sEzsignformfieldgroupRegexp"] = o.SEzsignformfieldgroupRegexp
	}
	if !IsNil(o.SEzsignformfieldgroupTextvalidationcustommessage) {
		toSerialize["sEzsignformfieldgroupTextvalidationcustommessage"] = o.SEzsignformfieldgroupTextvalidationcustommessage
	}
	if !IsNil(o.TEzsignformfieldgroupTooltip) {
		toSerialize["tEzsignformfieldgroupTooltip"] = o.TEzsignformfieldgroupTooltip
	}
	if !IsNil(o.EEzsignformfieldgroupTooltipposition) {
		toSerialize["eEzsignformfieldgroupTooltipposition"] = o.EEzsignformfieldgroupTooltipposition
	}
	if !IsNil(o.EEzsignformfieldgroupTextvalidation) {
		toSerialize["eEzsignformfieldgroupTextvalidation"] = o.EEzsignformfieldgroupTextvalidation
	}
	toSerialize["a_objEzsignformfieldgroupsigner"] = o.AObjEzsignformfieldgroupsigner
	if !IsNil(o.AObjDropdownElement) {
		toSerialize["a_objDropdownElement"] = o.AObjDropdownElement
	}
	toSerialize["a_objEzsignformfield"] = o.AObjEzsignformfield
	toSerialize["objCreateezsignelementspositionedbyword"] = o.ObjCreateezsignelementspositionedbyword
	return toSerialize, nil
}

func (o *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fkiEzsigndocumentID",
		"eEzsignformfieldgroupType",
		"sEzsignformfieldgroupLabel",
		"iEzsignformfieldgroupStep",
		"iEzsignformfieldgroupFilledmin",
		"iEzsignformfieldgroupFilledmax",
		"bEzsignformfieldgroupReadonly",
		"a_objEzsignformfieldgroupsigner",
		"a_objEzsignformfield",
		"objCreateezsignelementspositionedbyword",
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

	varCustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest := _CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest)

	if err != nil {
		return err
	}

	*o = CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest(varCustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest)

	return err
}

type NullableCustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest struct {
	value *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest
	isSet bool
}

func (v NullableCustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) Get() *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest {
	return v.value
}

func (v *NullableCustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) Set(val *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest(val *CustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) *NullableCustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest {
	return &NullableCustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest{value: val, isSet: true}
}

func (v NullableCustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomEzsignformfieldgroupCreateEzsignelementsPositionedByWordRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


