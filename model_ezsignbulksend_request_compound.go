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

// checks if the EzsignbulksendRequestCompound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignbulksendRequestCompound{}

// EzsignbulksendRequestCompound A Ezsignbulksend Object and children
type EzsignbulksendRequestCompound struct {
	// The unique ID of the Ezsignbulksend
	PkiEzsignbulksendID *int32 `json:"pkiEzsignbulksendID,omitempty"`
	// The unique ID of the Ezsignfoldertype.
	FkiEzsignfoldertypeID int32 `json:"fkiEzsignfoldertypeID"`
	// The unique ID of the Language.  Valid values:  |Value|Description| |-|-| |1|French| |2|English|
	FkiLanguageID int32 `json:"fkiLanguageID"`
	// The description of the Ezsignbulksend
	SEzsignbulksendDescription string `json:"sEzsignbulksendDescription"`
	// Note about the Ezsignbulksend
	TEzsignbulksendNote string `json:"tEzsignbulksendNote"`
	// Whether the Ezsigntemplatepackage was automatically modified and needs a manual validation
	BEzsignbulksendNeedvalidation bool `json:"bEzsignbulksendNeedvalidation"`
	// Whether the Ezsignbulksend is active or not
	BEzsignbulksendIsactive bool `json:"bEzsignbulksendIsactive"`
}

// NewEzsignbulksendRequestCompound instantiates a new EzsignbulksendRequestCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignbulksendRequestCompound(fkiEzsignfoldertypeID int32, fkiLanguageID int32, sEzsignbulksendDescription string, tEzsignbulksendNote string, bEzsignbulksendNeedvalidation bool, bEzsignbulksendIsactive bool) *EzsignbulksendRequestCompound {
	this := EzsignbulksendRequestCompound{}
	this.FkiEzsignfoldertypeID = fkiEzsignfoldertypeID
	this.FkiLanguageID = fkiLanguageID
	this.SEzsignbulksendDescription = sEzsignbulksendDescription
	this.TEzsignbulksendNote = tEzsignbulksendNote
	this.BEzsignbulksendNeedvalidation = bEzsignbulksendNeedvalidation
	this.BEzsignbulksendIsactive = bEzsignbulksendIsactive
	return &this
}

// NewEzsignbulksendRequestCompoundWithDefaults instantiates a new EzsignbulksendRequestCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignbulksendRequestCompoundWithDefaults() *EzsignbulksendRequestCompound {
	this := EzsignbulksendRequestCompound{}
	return &this
}

// GetPkiEzsignbulksendID returns the PkiEzsignbulksendID field value if set, zero value otherwise.
func (o *EzsignbulksendRequestCompound) GetPkiEzsignbulksendID() int32 {
	if o == nil || IsNil(o.PkiEzsignbulksendID) {
		var ret int32
		return ret
	}
	return *o.PkiEzsignbulksendID
}

// GetPkiEzsignbulksendIDOk returns a tuple with the PkiEzsignbulksendID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignbulksendRequestCompound) GetPkiEzsignbulksendIDOk() (*int32, bool) {
	if o == nil || IsNil(o.PkiEzsignbulksendID) {
		return nil, false
	}
	return o.PkiEzsignbulksendID, true
}

// HasPkiEzsignbulksendID returns a boolean if a field has been set.
func (o *EzsignbulksendRequestCompound) HasPkiEzsignbulksendID() bool {
	if o != nil && !IsNil(o.PkiEzsignbulksendID) {
		return true
	}

	return false
}

// SetPkiEzsignbulksendID gets a reference to the given int32 and assigns it to the PkiEzsignbulksendID field.
func (o *EzsignbulksendRequestCompound) SetPkiEzsignbulksendID(v int32) {
	o.PkiEzsignbulksendID = &v
}

// GetFkiEzsignfoldertypeID returns the FkiEzsignfoldertypeID field value
func (o *EzsignbulksendRequestCompound) GetFkiEzsignfoldertypeID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzsignfoldertypeID
}

// GetFkiEzsignfoldertypeIDOk returns a tuple with the FkiEzsignfoldertypeID field value
// and a boolean to check if the value has been set.
func (o *EzsignbulksendRequestCompound) GetFkiEzsignfoldertypeIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiEzsignfoldertypeID, true
}

// SetFkiEzsignfoldertypeID sets field value
func (o *EzsignbulksendRequestCompound) SetFkiEzsignfoldertypeID(v int32) {
	o.FkiEzsignfoldertypeID = v
}

// GetFkiLanguageID returns the FkiLanguageID field value
func (o *EzsignbulksendRequestCompound) GetFkiLanguageID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiLanguageID
}

// GetFkiLanguageIDOk returns a tuple with the FkiLanguageID field value
// and a boolean to check if the value has been set.
func (o *EzsignbulksendRequestCompound) GetFkiLanguageIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiLanguageID, true
}

// SetFkiLanguageID sets field value
func (o *EzsignbulksendRequestCompound) SetFkiLanguageID(v int32) {
	o.FkiLanguageID = v
}

// GetSEzsignbulksendDescription returns the SEzsignbulksendDescription field value
func (o *EzsignbulksendRequestCompound) GetSEzsignbulksendDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzsignbulksendDescription
}

// GetSEzsignbulksendDescriptionOk returns a tuple with the SEzsignbulksendDescription field value
// and a boolean to check if the value has been set.
func (o *EzsignbulksendRequestCompound) GetSEzsignbulksendDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SEzsignbulksendDescription, true
}

// SetSEzsignbulksendDescription sets field value
func (o *EzsignbulksendRequestCompound) SetSEzsignbulksendDescription(v string) {
	o.SEzsignbulksendDescription = v
}

// GetTEzsignbulksendNote returns the TEzsignbulksendNote field value
func (o *EzsignbulksendRequestCompound) GetTEzsignbulksendNote() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TEzsignbulksendNote
}

// GetTEzsignbulksendNoteOk returns a tuple with the TEzsignbulksendNote field value
// and a boolean to check if the value has been set.
func (o *EzsignbulksendRequestCompound) GetTEzsignbulksendNoteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TEzsignbulksendNote, true
}

// SetTEzsignbulksendNote sets field value
func (o *EzsignbulksendRequestCompound) SetTEzsignbulksendNote(v string) {
	o.TEzsignbulksendNote = v
}

// GetBEzsignbulksendNeedvalidation returns the BEzsignbulksendNeedvalidation field value
func (o *EzsignbulksendRequestCompound) GetBEzsignbulksendNeedvalidation() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BEzsignbulksendNeedvalidation
}

// GetBEzsignbulksendNeedvalidationOk returns a tuple with the BEzsignbulksendNeedvalidation field value
// and a boolean to check if the value has been set.
func (o *EzsignbulksendRequestCompound) GetBEzsignbulksendNeedvalidationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BEzsignbulksendNeedvalidation, true
}

// SetBEzsignbulksendNeedvalidation sets field value
func (o *EzsignbulksendRequestCompound) SetBEzsignbulksendNeedvalidation(v bool) {
	o.BEzsignbulksendNeedvalidation = v
}

// GetBEzsignbulksendIsactive returns the BEzsignbulksendIsactive field value
func (o *EzsignbulksendRequestCompound) GetBEzsignbulksendIsactive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BEzsignbulksendIsactive
}

// GetBEzsignbulksendIsactiveOk returns a tuple with the BEzsignbulksendIsactive field value
// and a boolean to check if the value has been set.
func (o *EzsignbulksendRequestCompound) GetBEzsignbulksendIsactiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BEzsignbulksendIsactive, true
}

// SetBEzsignbulksendIsactive sets field value
func (o *EzsignbulksendRequestCompound) SetBEzsignbulksendIsactive(v bool) {
	o.BEzsignbulksendIsactive = v
}

func (o EzsignbulksendRequestCompound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignbulksendRequestCompound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PkiEzsignbulksendID) {
		toSerialize["pkiEzsignbulksendID"] = o.PkiEzsignbulksendID
	}
	toSerialize["fkiEzsignfoldertypeID"] = o.FkiEzsignfoldertypeID
	toSerialize["fkiLanguageID"] = o.FkiLanguageID
	toSerialize["sEzsignbulksendDescription"] = o.SEzsignbulksendDescription
	toSerialize["tEzsignbulksendNote"] = o.TEzsignbulksendNote
	toSerialize["bEzsignbulksendNeedvalidation"] = o.BEzsignbulksendNeedvalidation
	toSerialize["bEzsignbulksendIsactive"] = o.BEzsignbulksendIsactive
	return toSerialize, nil
}

type NullableEzsignbulksendRequestCompound struct {
	value *EzsignbulksendRequestCompound
	isSet bool
}

func (v NullableEzsignbulksendRequestCompound) Get() *EzsignbulksendRequestCompound {
	return v.value
}

func (v *NullableEzsignbulksendRequestCompound) Set(val *EzsignbulksendRequestCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignbulksendRequestCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignbulksendRequestCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignbulksendRequestCompound(val *EzsignbulksendRequestCompound) *NullableEzsignbulksendRequestCompound {
	return &NullableEzsignbulksendRequestCompound{value: val, isSet: true}
}

func (v NullableEzsignbulksendRequestCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignbulksendRequestCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

