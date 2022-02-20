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

// EzsigntemplatepackageListElement An Ezsigntemplatepackage List Element
type EzsigntemplatepackageListElement struct {
	// The unique ID of the Ezsigntemplatepackage
	PkiEzsigntemplatepackageID int32 `json:"pkiEzsigntemplatepackageID"`
	// The unique ID of the Department.
	FkiDepartmentID NullableInt32 `json:"fkiDepartmentID"`
	// The unique ID of the Team
	FkiTeamID NullableInt32 `json:"fkiTeamID"`
	// The unique ID of the Ezsignfoldertype.
	FkiEzsignfoldertypeID NullableInt32 `json:"fkiEzsignfoldertypeID"`
	// The unique ID of the Language.  Valid values:  |Value|Description| |-|-| |1|French| |2|English|
	FkiLanguageID int32 `json:"fkiLanguageID"`
	EEzsigntemplatepackageType FieldEEzsigntemplatepackageType `json:"eEzsigntemplatepackageType"`
	// The description of the Ezsigntemplatepackage
	SEzsigntemplatepackageDescription string `json:"sEzsigntemplatepackageDescription"`
	// Whether the Ezsigntemplatepackage is active or not
	BEzsigntemplatepackageIsactive bool `json:"bEzsigntemplatepackageIsactive"`
	// The total number of Ezsigntemplatepackagemembership in the Ezsigntemplatepackage
	IEzsigntemplatepackagemembership int32 `json:"iEzsigntemplatepackagemembership"`
}

// NewEzsigntemplatepackageListElement instantiates a new EzsigntemplatepackageListElement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigntemplatepackageListElement(pkiEzsigntemplatepackageID int32, fkiDepartmentID NullableInt32, fkiTeamID NullableInt32, fkiEzsignfoldertypeID NullableInt32, fkiLanguageID int32, eEzsigntemplatepackageType FieldEEzsigntemplatepackageType, sEzsigntemplatepackageDescription string, bEzsigntemplatepackageIsactive bool, iEzsigntemplatepackagemembership int32) *EzsigntemplatepackageListElement {
	this := EzsigntemplatepackageListElement{}
	this.PkiEzsigntemplatepackageID = pkiEzsigntemplatepackageID
	this.FkiDepartmentID = fkiDepartmentID
	this.FkiTeamID = fkiTeamID
	this.FkiEzsignfoldertypeID = fkiEzsignfoldertypeID
	this.FkiLanguageID = fkiLanguageID
	this.EEzsigntemplatepackageType = eEzsigntemplatepackageType
	this.SEzsigntemplatepackageDescription = sEzsigntemplatepackageDescription
	this.BEzsigntemplatepackageIsactive = bEzsigntemplatepackageIsactive
	this.IEzsigntemplatepackagemembership = iEzsigntemplatepackagemembership
	return &this
}

// NewEzsigntemplatepackageListElementWithDefaults instantiates a new EzsigntemplatepackageListElement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigntemplatepackageListElementWithDefaults() *EzsigntemplatepackageListElement {
	this := EzsigntemplatepackageListElement{}
	return &this
}

// GetPkiEzsigntemplatepackageID returns the PkiEzsigntemplatepackageID field value
func (o *EzsigntemplatepackageListElement) GetPkiEzsigntemplatepackageID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiEzsigntemplatepackageID
}

// GetPkiEzsigntemplatepackageIDOk returns a tuple with the PkiEzsigntemplatepackageID field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepackageListElement) GetPkiEzsigntemplatepackageIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiEzsigntemplatepackageID, true
}

// SetPkiEzsigntemplatepackageID sets field value
func (o *EzsigntemplatepackageListElement) SetPkiEzsigntemplatepackageID(v int32) {
	o.PkiEzsigntemplatepackageID = v
}

// GetFkiDepartmentID returns the FkiDepartmentID field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *EzsigntemplatepackageListElement) GetFkiDepartmentID() int32 {
	if o == nil || o.FkiDepartmentID.Get() == nil {
		var ret int32
		return ret
	}

	return *o.FkiDepartmentID.Get()
}

// GetFkiDepartmentIDOk returns a tuple with the FkiDepartmentID field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EzsigntemplatepackageListElement) GetFkiDepartmentIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.FkiDepartmentID.Get(), o.FkiDepartmentID.IsSet()
}

// SetFkiDepartmentID sets field value
func (o *EzsigntemplatepackageListElement) SetFkiDepartmentID(v int32) {
	o.FkiDepartmentID.Set(&v)
}

// GetFkiTeamID returns the FkiTeamID field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *EzsigntemplatepackageListElement) GetFkiTeamID() int32 {
	if o == nil || o.FkiTeamID.Get() == nil {
		var ret int32
		return ret
	}

	return *o.FkiTeamID.Get()
}

// GetFkiTeamIDOk returns a tuple with the FkiTeamID field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EzsigntemplatepackageListElement) GetFkiTeamIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.FkiTeamID.Get(), o.FkiTeamID.IsSet()
}

// SetFkiTeamID sets field value
func (o *EzsigntemplatepackageListElement) SetFkiTeamID(v int32) {
	o.FkiTeamID.Set(&v)
}

// GetFkiEzsignfoldertypeID returns the FkiEzsignfoldertypeID field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *EzsigntemplatepackageListElement) GetFkiEzsignfoldertypeID() int32 {
	if o == nil || o.FkiEzsignfoldertypeID.Get() == nil {
		var ret int32
		return ret
	}

	return *o.FkiEzsignfoldertypeID.Get()
}

// GetFkiEzsignfoldertypeIDOk returns a tuple with the FkiEzsignfoldertypeID field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EzsigntemplatepackageListElement) GetFkiEzsignfoldertypeIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.FkiEzsignfoldertypeID.Get(), o.FkiEzsignfoldertypeID.IsSet()
}

// SetFkiEzsignfoldertypeID sets field value
func (o *EzsigntemplatepackageListElement) SetFkiEzsignfoldertypeID(v int32) {
	o.FkiEzsignfoldertypeID.Set(&v)
}

// GetFkiLanguageID returns the FkiLanguageID field value
func (o *EzsigntemplatepackageListElement) GetFkiLanguageID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiLanguageID
}

// GetFkiLanguageIDOk returns a tuple with the FkiLanguageID field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepackageListElement) GetFkiLanguageIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiLanguageID, true
}

// SetFkiLanguageID sets field value
func (o *EzsigntemplatepackageListElement) SetFkiLanguageID(v int32) {
	o.FkiLanguageID = v
}

// GetEEzsigntemplatepackageType returns the EEzsigntemplatepackageType field value
func (o *EzsigntemplatepackageListElement) GetEEzsigntemplatepackageType() FieldEEzsigntemplatepackageType {
	if o == nil {
		var ret FieldEEzsigntemplatepackageType
		return ret
	}

	return o.EEzsigntemplatepackageType
}

// GetEEzsigntemplatepackageTypeOk returns a tuple with the EEzsigntemplatepackageType field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepackageListElement) GetEEzsigntemplatepackageTypeOk() (*FieldEEzsigntemplatepackageType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EEzsigntemplatepackageType, true
}

// SetEEzsigntemplatepackageType sets field value
func (o *EzsigntemplatepackageListElement) SetEEzsigntemplatepackageType(v FieldEEzsigntemplatepackageType) {
	o.EEzsigntemplatepackageType = v
}

// GetSEzsigntemplatepackageDescription returns the SEzsigntemplatepackageDescription field value
func (o *EzsigntemplatepackageListElement) GetSEzsigntemplatepackageDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzsigntemplatepackageDescription
}

// GetSEzsigntemplatepackageDescriptionOk returns a tuple with the SEzsigntemplatepackageDescription field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepackageListElement) GetSEzsigntemplatepackageDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SEzsigntemplatepackageDescription, true
}

// SetSEzsigntemplatepackageDescription sets field value
func (o *EzsigntemplatepackageListElement) SetSEzsigntemplatepackageDescription(v string) {
	o.SEzsigntemplatepackageDescription = v
}

// GetBEzsigntemplatepackageIsactive returns the BEzsigntemplatepackageIsactive field value
func (o *EzsigntemplatepackageListElement) GetBEzsigntemplatepackageIsactive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BEzsigntemplatepackageIsactive
}

// GetBEzsigntemplatepackageIsactiveOk returns a tuple with the BEzsigntemplatepackageIsactive field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepackageListElement) GetBEzsigntemplatepackageIsactiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BEzsigntemplatepackageIsactive, true
}

// SetBEzsigntemplatepackageIsactive sets field value
func (o *EzsigntemplatepackageListElement) SetBEzsigntemplatepackageIsactive(v bool) {
	o.BEzsigntemplatepackageIsactive = v
}

// GetIEzsigntemplatepackagemembership returns the IEzsigntemplatepackagemembership field value
func (o *EzsigntemplatepackageListElement) GetIEzsigntemplatepackagemembership() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsigntemplatepackagemembership
}

// GetIEzsigntemplatepackagemembershipOk returns a tuple with the IEzsigntemplatepackagemembership field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepackageListElement) GetIEzsigntemplatepackagemembershipOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzsigntemplatepackagemembership, true
}

// SetIEzsigntemplatepackagemembership sets field value
func (o *EzsigntemplatepackageListElement) SetIEzsigntemplatepackagemembership(v int32) {
	o.IEzsigntemplatepackagemembership = v
}

func (o EzsigntemplatepackageListElement) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pkiEzsigntemplatepackageID"] = o.PkiEzsigntemplatepackageID
	}
	if true {
		toSerialize["fkiDepartmentID"] = o.FkiDepartmentID.Get()
	}
	if true {
		toSerialize["fkiTeamID"] = o.FkiTeamID.Get()
	}
	if true {
		toSerialize["fkiEzsignfoldertypeID"] = o.FkiEzsignfoldertypeID.Get()
	}
	if true {
		toSerialize["fkiLanguageID"] = o.FkiLanguageID
	}
	if true {
		toSerialize["eEzsigntemplatepackageType"] = o.EEzsigntemplatepackageType
	}
	if true {
		toSerialize["sEzsigntemplatepackageDescription"] = o.SEzsigntemplatepackageDescription
	}
	if true {
		toSerialize["bEzsigntemplatepackageIsactive"] = o.BEzsigntemplatepackageIsactive
	}
	if true {
		toSerialize["iEzsigntemplatepackagemembership"] = o.IEzsigntemplatepackagemembership
	}
	return json.Marshal(toSerialize)
}

type NullableEzsigntemplatepackageListElement struct {
	value *EzsigntemplatepackageListElement
	isSet bool
}

func (v NullableEzsigntemplatepackageListElement) Get() *EzsigntemplatepackageListElement {
	return v.value
}

func (v *NullableEzsigntemplatepackageListElement) Set(val *EzsigntemplatepackageListElement) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigntemplatepackageListElement) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigntemplatepackageListElement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigntemplatepackageListElement(val *EzsigntemplatepackageListElement) *NullableEzsigntemplatepackageListElement {
	return &NullableEzsigntemplatepackageListElement{value: val, isSet: true}
}

func (v NullableEzsigntemplatepackageListElement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigntemplatepackageListElement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


