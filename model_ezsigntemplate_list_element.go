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

// checks if the EzsigntemplateListElement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigntemplateListElement{}

// EzsigntemplateListElement A Ezsigntemplate List Element
type EzsigntemplateListElement struct {
	// The unique ID of the Ezsigntemplate
	PkiEzsigntemplateID int32 `json:"pkiEzsigntemplateID"`
	// The unique ID of the Ezsignfoldertype.
	FkiEzsignfoldertypeID int32 `json:"fkiEzsignfoldertypeID"`
	// The unique ID of the Language.  Valid values:  |Value|Description| |-|-| |1|French| |2|English|
	FkiLanguageID int32 `json:"fkiLanguageID"`
	// The description of the Ezsigntemplate
	SEzsigntemplateDescription string `json:"sEzsigntemplateDescription"`
	// The number of pages in the Ezsigntemplatedocument.
	IEzsigntemplatedocumentPagetotal *int32 `json:"iEzsigntemplatedocumentPagetotal,omitempty"`
	// The number of total signatures in the Ezsigntemplate.
	IEzsigntemplateSignaturetotal *int32 `json:"iEzsigntemplateSignaturetotal,omitempty"`
	// The number of total form fields in the Ezsigntemplate.
	IEzsigntemplateFormfieldtotal *int32 `json:"iEzsigntemplateFormfieldtotal,omitempty"`
	// Indicate the Ezsigntemplate is incomplete and cannot be used
	BEzsigntemplateIncomplete bool `json:"bEzsigntemplateIncomplete"`
	// The name of the Ezsignfoldertype in the language of the requester
	SEzsignfoldertypeNameX string `json:"sEzsignfoldertypeNameX"`
}

// NewEzsigntemplateListElement instantiates a new EzsigntemplateListElement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigntemplateListElement(pkiEzsigntemplateID int32, fkiEzsignfoldertypeID int32, fkiLanguageID int32, sEzsigntemplateDescription string, bEzsigntemplateIncomplete bool, sEzsignfoldertypeNameX string) *EzsigntemplateListElement {
	this := EzsigntemplateListElement{}
	this.PkiEzsigntemplateID = pkiEzsigntemplateID
	this.FkiEzsignfoldertypeID = fkiEzsignfoldertypeID
	this.FkiLanguageID = fkiLanguageID
	this.SEzsigntemplateDescription = sEzsigntemplateDescription
	this.BEzsigntemplateIncomplete = bEzsigntemplateIncomplete
	this.SEzsignfoldertypeNameX = sEzsignfoldertypeNameX
	return &this
}

// NewEzsigntemplateListElementWithDefaults instantiates a new EzsigntemplateListElement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigntemplateListElementWithDefaults() *EzsigntemplateListElement {
	this := EzsigntemplateListElement{}
	return &this
}

// GetPkiEzsigntemplateID returns the PkiEzsigntemplateID field value
func (o *EzsigntemplateListElement) GetPkiEzsigntemplateID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiEzsigntemplateID
}

// GetPkiEzsigntemplateIDOk returns a tuple with the PkiEzsigntemplateID field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateListElement) GetPkiEzsigntemplateIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiEzsigntemplateID, true
}

// SetPkiEzsigntemplateID sets field value
func (o *EzsigntemplateListElement) SetPkiEzsigntemplateID(v int32) {
	o.PkiEzsigntemplateID = v
}

// GetFkiEzsignfoldertypeID returns the FkiEzsignfoldertypeID field value
func (o *EzsigntemplateListElement) GetFkiEzsignfoldertypeID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzsignfoldertypeID
}

// GetFkiEzsignfoldertypeIDOk returns a tuple with the FkiEzsignfoldertypeID field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateListElement) GetFkiEzsignfoldertypeIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiEzsignfoldertypeID, true
}

// SetFkiEzsignfoldertypeID sets field value
func (o *EzsigntemplateListElement) SetFkiEzsignfoldertypeID(v int32) {
	o.FkiEzsignfoldertypeID = v
}

// GetFkiLanguageID returns the FkiLanguageID field value
func (o *EzsigntemplateListElement) GetFkiLanguageID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiLanguageID
}

// GetFkiLanguageIDOk returns a tuple with the FkiLanguageID field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateListElement) GetFkiLanguageIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiLanguageID, true
}

// SetFkiLanguageID sets field value
func (o *EzsigntemplateListElement) SetFkiLanguageID(v int32) {
	o.FkiLanguageID = v
}

// GetSEzsigntemplateDescription returns the SEzsigntemplateDescription field value
func (o *EzsigntemplateListElement) GetSEzsigntemplateDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzsigntemplateDescription
}

// GetSEzsigntemplateDescriptionOk returns a tuple with the SEzsigntemplateDescription field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateListElement) GetSEzsigntemplateDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SEzsigntemplateDescription, true
}

// SetSEzsigntemplateDescription sets field value
func (o *EzsigntemplateListElement) SetSEzsigntemplateDescription(v string) {
	o.SEzsigntemplateDescription = v
}

// GetIEzsigntemplatedocumentPagetotal returns the IEzsigntemplatedocumentPagetotal field value if set, zero value otherwise.
func (o *EzsigntemplateListElement) GetIEzsigntemplatedocumentPagetotal() int32 {
	if o == nil || IsNil(o.IEzsigntemplatedocumentPagetotal) {
		var ret int32
		return ret
	}
	return *o.IEzsigntemplatedocumentPagetotal
}

// GetIEzsigntemplatedocumentPagetotalOk returns a tuple with the IEzsigntemplatedocumentPagetotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateListElement) GetIEzsigntemplatedocumentPagetotalOk() (*int32, bool) {
	if o == nil || IsNil(o.IEzsigntemplatedocumentPagetotal) {
		return nil, false
	}
	return o.IEzsigntemplatedocumentPagetotal, true
}

// HasIEzsigntemplatedocumentPagetotal returns a boolean if a field has been set.
func (o *EzsigntemplateListElement) HasIEzsigntemplatedocumentPagetotal() bool {
	if o != nil && !IsNil(o.IEzsigntemplatedocumentPagetotal) {
		return true
	}

	return false
}

// SetIEzsigntemplatedocumentPagetotal gets a reference to the given int32 and assigns it to the IEzsigntemplatedocumentPagetotal field.
func (o *EzsigntemplateListElement) SetIEzsigntemplatedocumentPagetotal(v int32) {
	o.IEzsigntemplatedocumentPagetotal = &v
}

// GetIEzsigntemplateSignaturetotal returns the IEzsigntemplateSignaturetotal field value if set, zero value otherwise.
func (o *EzsigntemplateListElement) GetIEzsigntemplateSignaturetotal() int32 {
	if o == nil || IsNil(o.IEzsigntemplateSignaturetotal) {
		var ret int32
		return ret
	}
	return *o.IEzsigntemplateSignaturetotal
}

// GetIEzsigntemplateSignaturetotalOk returns a tuple with the IEzsigntemplateSignaturetotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateListElement) GetIEzsigntemplateSignaturetotalOk() (*int32, bool) {
	if o == nil || IsNil(o.IEzsigntemplateSignaturetotal) {
		return nil, false
	}
	return o.IEzsigntemplateSignaturetotal, true
}

// HasIEzsigntemplateSignaturetotal returns a boolean if a field has been set.
func (o *EzsigntemplateListElement) HasIEzsigntemplateSignaturetotal() bool {
	if o != nil && !IsNil(o.IEzsigntemplateSignaturetotal) {
		return true
	}

	return false
}

// SetIEzsigntemplateSignaturetotal gets a reference to the given int32 and assigns it to the IEzsigntemplateSignaturetotal field.
func (o *EzsigntemplateListElement) SetIEzsigntemplateSignaturetotal(v int32) {
	o.IEzsigntemplateSignaturetotal = &v
}

// GetIEzsigntemplateFormfieldtotal returns the IEzsigntemplateFormfieldtotal field value if set, zero value otherwise.
func (o *EzsigntemplateListElement) GetIEzsigntemplateFormfieldtotal() int32 {
	if o == nil || IsNil(o.IEzsigntemplateFormfieldtotal) {
		var ret int32
		return ret
	}
	return *o.IEzsigntemplateFormfieldtotal
}

// GetIEzsigntemplateFormfieldtotalOk returns a tuple with the IEzsigntemplateFormfieldtotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateListElement) GetIEzsigntemplateFormfieldtotalOk() (*int32, bool) {
	if o == nil || IsNil(o.IEzsigntemplateFormfieldtotal) {
		return nil, false
	}
	return o.IEzsigntemplateFormfieldtotal, true
}

// HasIEzsigntemplateFormfieldtotal returns a boolean if a field has been set.
func (o *EzsigntemplateListElement) HasIEzsigntemplateFormfieldtotal() bool {
	if o != nil && !IsNil(o.IEzsigntemplateFormfieldtotal) {
		return true
	}

	return false
}

// SetIEzsigntemplateFormfieldtotal gets a reference to the given int32 and assigns it to the IEzsigntemplateFormfieldtotal field.
func (o *EzsigntemplateListElement) SetIEzsigntemplateFormfieldtotal(v int32) {
	o.IEzsigntemplateFormfieldtotal = &v
}

// GetBEzsigntemplateIncomplete returns the BEzsigntemplateIncomplete field value
func (o *EzsigntemplateListElement) GetBEzsigntemplateIncomplete() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BEzsigntemplateIncomplete
}

// GetBEzsigntemplateIncompleteOk returns a tuple with the BEzsigntemplateIncomplete field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateListElement) GetBEzsigntemplateIncompleteOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BEzsigntemplateIncomplete, true
}

// SetBEzsigntemplateIncomplete sets field value
func (o *EzsigntemplateListElement) SetBEzsigntemplateIncomplete(v bool) {
	o.BEzsigntemplateIncomplete = v
}

// GetSEzsignfoldertypeNameX returns the SEzsignfoldertypeNameX field value
func (o *EzsigntemplateListElement) GetSEzsignfoldertypeNameX() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzsignfoldertypeNameX
}

// GetSEzsignfoldertypeNameXOk returns a tuple with the SEzsignfoldertypeNameX field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateListElement) GetSEzsignfoldertypeNameXOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SEzsignfoldertypeNameX, true
}

// SetSEzsignfoldertypeNameX sets field value
func (o *EzsigntemplateListElement) SetSEzsignfoldertypeNameX(v string) {
	o.SEzsignfoldertypeNameX = v
}

func (o EzsigntemplateListElement) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigntemplateListElement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiEzsigntemplateID"] = o.PkiEzsigntemplateID
	toSerialize["fkiEzsignfoldertypeID"] = o.FkiEzsignfoldertypeID
	toSerialize["fkiLanguageID"] = o.FkiLanguageID
	toSerialize["sEzsigntemplateDescription"] = o.SEzsigntemplateDescription
	if !IsNil(o.IEzsigntemplatedocumentPagetotal) {
		toSerialize["iEzsigntemplatedocumentPagetotal"] = o.IEzsigntemplatedocumentPagetotal
	}
	if !IsNil(o.IEzsigntemplateSignaturetotal) {
		toSerialize["iEzsigntemplateSignaturetotal"] = o.IEzsigntemplateSignaturetotal
	}
	if !IsNil(o.IEzsigntemplateFormfieldtotal) {
		toSerialize["iEzsigntemplateFormfieldtotal"] = o.IEzsigntemplateFormfieldtotal
	}
	toSerialize["bEzsigntemplateIncomplete"] = o.BEzsigntemplateIncomplete
	toSerialize["sEzsignfoldertypeNameX"] = o.SEzsignfoldertypeNameX
	return toSerialize, nil
}

type NullableEzsigntemplateListElement struct {
	value *EzsigntemplateListElement
	isSet bool
}

func (v NullableEzsigntemplateListElement) Get() *EzsigntemplateListElement {
	return v.value
}

func (v *NullableEzsigntemplateListElement) Set(val *EzsigntemplateListElement) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigntemplateListElement) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigntemplateListElement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigntemplateListElement(val *EzsigntemplateListElement) *NullableEzsigntemplateListElement {
	return &NullableEzsigntemplateListElement{value: val, isSet: true}
}

func (v NullableEzsigntemplateListElement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigntemplateListElement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

