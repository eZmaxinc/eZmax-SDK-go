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

// checks if the EzsigntemplateListElement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigntemplateListElement{}

// EzsigntemplateListElement A Ezsigntemplate List Element
type EzsigntemplateListElement struct {
	// The unique ID of the Ezsigntemplate
	PkiEzsigntemplateID int32 `json:"pkiEzsigntemplateID"`
	// The unique ID of the Ezsignfoldertype.
	FkiEzsignfoldertypeID *int32 `json:"fkiEzsignfoldertypeID,omitempty"`
	// The unique ID of the Language.  Valid values:  |Value|Description| |-|-| |1|French| |2|English|
	FkiLanguageID int32 `json:"fkiLanguageID"`
	// The description of the Ezsigntemplate
	SEzsigntemplateDescription string `json:"sEzsigntemplateDescription" validate:"regexp=^.{0,80}$"`
	// The number of pages in the Ezsigntemplatedocument.
	IEzsigntemplatedocumentPagetotal *int32 `json:"iEzsigntemplatedocumentPagetotal,omitempty"`
	// The number of total signatures in the Ezsigntemplate.
	IEzsigntemplateSignaturetotal *int32 `json:"iEzsigntemplateSignaturetotal,omitempty"`
	// The number of total form fields in the Ezsigntemplate.
	IEzsigntemplateFormfieldtotal *int32 `json:"iEzsigntemplateFormfieldtotal,omitempty"`
	// Indicate the Ezsigntemplate is incomplete and cannot be used
	BEzsigntemplateIncomplete bool `json:"bEzsigntemplateIncomplete"`
	// The name of the Ezsignfoldertype in the language of the requester
	SEzsignfoldertypeNameX *string `json:"sEzsignfoldertypeNameX,omitempty"`
	EEzsigntemplateType FieldEEzsigntemplateType `json:"eEzsigntemplateType"`
}

type _EzsigntemplateListElement EzsigntemplateListElement

// NewEzsigntemplateListElement instantiates a new EzsigntemplateListElement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigntemplateListElement(pkiEzsigntemplateID int32, fkiLanguageID int32, sEzsigntemplateDescription string, bEzsigntemplateIncomplete bool, eEzsigntemplateType FieldEEzsigntemplateType) *EzsigntemplateListElement {
	this := EzsigntemplateListElement{}
	this.PkiEzsigntemplateID = pkiEzsigntemplateID
	this.FkiLanguageID = fkiLanguageID
	this.SEzsigntemplateDescription = sEzsigntemplateDescription
	this.BEzsigntemplateIncomplete = bEzsigntemplateIncomplete
	this.EEzsigntemplateType = eEzsigntemplateType
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

// GetFkiEzsignfoldertypeID returns the FkiEzsignfoldertypeID field value if set, zero value otherwise.
func (o *EzsigntemplateListElement) GetFkiEzsignfoldertypeID() int32 {
	if o == nil || IsNil(o.FkiEzsignfoldertypeID) {
		var ret int32
		return ret
	}
	return *o.FkiEzsignfoldertypeID
}

// GetFkiEzsignfoldertypeIDOk returns a tuple with the FkiEzsignfoldertypeID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateListElement) GetFkiEzsignfoldertypeIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiEzsignfoldertypeID) {
		return nil, false
	}
	return o.FkiEzsignfoldertypeID, true
}

// HasFkiEzsignfoldertypeID returns a boolean if a field has been set.
func (o *EzsigntemplateListElement) HasFkiEzsignfoldertypeID() bool {
	if o != nil && !IsNil(o.FkiEzsignfoldertypeID) {
		return true
	}

	return false
}

// SetFkiEzsignfoldertypeID gets a reference to the given int32 and assigns it to the FkiEzsignfoldertypeID field.
func (o *EzsigntemplateListElement) SetFkiEzsignfoldertypeID(v int32) {
	o.FkiEzsignfoldertypeID = &v
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

// GetSEzsignfoldertypeNameX returns the SEzsignfoldertypeNameX field value if set, zero value otherwise.
func (o *EzsigntemplateListElement) GetSEzsignfoldertypeNameX() string {
	if o == nil || IsNil(o.SEzsignfoldertypeNameX) {
		var ret string
		return ret
	}
	return *o.SEzsignfoldertypeNameX
}

// GetSEzsignfoldertypeNameXOk returns a tuple with the SEzsignfoldertypeNameX field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateListElement) GetSEzsignfoldertypeNameXOk() (*string, bool) {
	if o == nil || IsNil(o.SEzsignfoldertypeNameX) {
		return nil, false
	}
	return o.SEzsignfoldertypeNameX, true
}

// HasSEzsignfoldertypeNameX returns a boolean if a field has been set.
func (o *EzsigntemplateListElement) HasSEzsignfoldertypeNameX() bool {
	if o != nil && !IsNil(o.SEzsignfoldertypeNameX) {
		return true
	}

	return false
}

// SetSEzsignfoldertypeNameX gets a reference to the given string and assigns it to the SEzsignfoldertypeNameX field.
func (o *EzsigntemplateListElement) SetSEzsignfoldertypeNameX(v string) {
	o.SEzsignfoldertypeNameX = &v
}

// GetEEzsigntemplateType returns the EEzsigntemplateType field value
func (o *EzsigntemplateListElement) GetEEzsigntemplateType() FieldEEzsigntemplateType {
	if o == nil {
		var ret FieldEEzsigntemplateType
		return ret
	}

	return o.EEzsigntemplateType
}

// GetEEzsigntemplateTypeOk returns a tuple with the EEzsigntemplateType field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateListElement) GetEEzsigntemplateTypeOk() (*FieldEEzsigntemplateType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EEzsigntemplateType, true
}

// SetEEzsigntemplateType sets field value
func (o *EzsigntemplateListElement) SetEEzsigntemplateType(v FieldEEzsigntemplateType) {
	o.EEzsigntemplateType = v
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
	if !IsNil(o.FkiEzsignfoldertypeID) {
		toSerialize["fkiEzsignfoldertypeID"] = o.FkiEzsignfoldertypeID
	}
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
	if !IsNil(o.SEzsignfoldertypeNameX) {
		toSerialize["sEzsignfoldertypeNameX"] = o.SEzsignfoldertypeNameX
	}
	toSerialize["eEzsigntemplateType"] = o.EEzsigntemplateType
	return toSerialize, nil
}

func (o *EzsigntemplateListElement) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiEzsigntemplateID",
		"fkiLanguageID",
		"sEzsigntemplateDescription",
		"bEzsigntemplateIncomplete",
		"eEzsigntemplateType",
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

	varEzsigntemplateListElement := _EzsigntemplateListElement{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsigntemplateListElement)

	if err != nil {
		return err
	}

	*o = EzsigntemplateListElement(varEzsigntemplateListElement)

	return err
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


