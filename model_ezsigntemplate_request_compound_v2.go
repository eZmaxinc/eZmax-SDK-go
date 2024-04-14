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

// checks if the EzsigntemplateRequestCompoundV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigntemplateRequestCompoundV2{}

// EzsigntemplateRequestCompoundV2 A Ezsigntemplate Object and children
type EzsigntemplateRequestCompoundV2 struct {
	// The unique ID of the Ezsigntemplate
	PkiEzsigntemplateID *int32 `json:"pkiEzsigntemplateID,omitempty"`
	// The unique ID of the Ezsignfoldertype.
	FkiEzsignfoldertypeID *int32 `json:"fkiEzsignfoldertypeID,omitempty"`
	// The unique ID of the Language.  Valid values:  |Value|Description| |-|-| |1|French| |2|English|
	FkiLanguageID int32 `json:"fkiLanguageID"`
	// The description of the Ezsigntemplate
	SEzsigntemplateDescription string `json:"sEzsigntemplateDescription"`
	// The filename pattern of the Ezsigntemplate
	SEzsigntemplateFilenamepattern *string `json:"sEzsigntemplateFilenamepattern,omitempty"`
	// Whether the Ezsigntemplate can be accessed by admin users only (eUserType=Normal)
	BEzsigntemplateAdminonly bool `json:"bEzsigntemplateAdminonly"`
	EEzsigntemplateType FieldEEzsigntemplateType `json:"eEzsigntemplateType"`
}

type _EzsigntemplateRequestCompoundV2 EzsigntemplateRequestCompoundV2

// NewEzsigntemplateRequestCompoundV2 instantiates a new EzsigntemplateRequestCompoundV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigntemplateRequestCompoundV2(fkiLanguageID int32, sEzsigntemplateDescription string, bEzsigntemplateAdminonly bool, eEzsigntemplateType FieldEEzsigntemplateType) *EzsigntemplateRequestCompoundV2 {
	this := EzsigntemplateRequestCompoundV2{}
	this.FkiLanguageID = fkiLanguageID
	this.SEzsigntemplateDescription = sEzsigntemplateDescription
	this.BEzsigntemplateAdminonly = bEzsigntemplateAdminonly
	this.EEzsigntemplateType = eEzsigntemplateType
	return &this
}

// NewEzsigntemplateRequestCompoundV2WithDefaults instantiates a new EzsigntemplateRequestCompoundV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigntemplateRequestCompoundV2WithDefaults() *EzsigntemplateRequestCompoundV2 {
	this := EzsigntemplateRequestCompoundV2{}
	return &this
}

// GetPkiEzsigntemplateID returns the PkiEzsigntemplateID field value if set, zero value otherwise.
func (o *EzsigntemplateRequestCompoundV2) GetPkiEzsigntemplateID() int32 {
	if o == nil || IsNil(o.PkiEzsigntemplateID) {
		var ret int32
		return ret
	}
	return *o.PkiEzsigntemplateID
}

// GetPkiEzsigntemplateIDOk returns a tuple with the PkiEzsigntemplateID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateRequestCompoundV2) GetPkiEzsigntemplateIDOk() (*int32, bool) {
	if o == nil || IsNil(o.PkiEzsigntemplateID) {
		return nil, false
	}
	return o.PkiEzsigntemplateID, true
}

// HasPkiEzsigntemplateID returns a boolean if a field has been set.
func (o *EzsigntemplateRequestCompoundV2) HasPkiEzsigntemplateID() bool {
	if o != nil && !IsNil(o.PkiEzsigntemplateID) {
		return true
	}

	return false
}

// SetPkiEzsigntemplateID gets a reference to the given int32 and assigns it to the PkiEzsigntemplateID field.
func (o *EzsigntemplateRequestCompoundV2) SetPkiEzsigntemplateID(v int32) {
	o.PkiEzsigntemplateID = &v
}

// GetFkiEzsignfoldertypeID returns the FkiEzsignfoldertypeID field value if set, zero value otherwise.
func (o *EzsigntemplateRequestCompoundV2) GetFkiEzsignfoldertypeID() int32 {
	if o == nil || IsNil(o.FkiEzsignfoldertypeID) {
		var ret int32
		return ret
	}
	return *o.FkiEzsignfoldertypeID
}

// GetFkiEzsignfoldertypeIDOk returns a tuple with the FkiEzsignfoldertypeID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateRequestCompoundV2) GetFkiEzsignfoldertypeIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiEzsignfoldertypeID) {
		return nil, false
	}
	return o.FkiEzsignfoldertypeID, true
}

// HasFkiEzsignfoldertypeID returns a boolean if a field has been set.
func (o *EzsigntemplateRequestCompoundV2) HasFkiEzsignfoldertypeID() bool {
	if o != nil && !IsNil(o.FkiEzsignfoldertypeID) {
		return true
	}

	return false
}

// SetFkiEzsignfoldertypeID gets a reference to the given int32 and assigns it to the FkiEzsignfoldertypeID field.
func (o *EzsigntemplateRequestCompoundV2) SetFkiEzsignfoldertypeID(v int32) {
	o.FkiEzsignfoldertypeID = &v
}

// GetFkiLanguageID returns the FkiLanguageID field value
func (o *EzsigntemplateRequestCompoundV2) GetFkiLanguageID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiLanguageID
}

// GetFkiLanguageIDOk returns a tuple with the FkiLanguageID field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateRequestCompoundV2) GetFkiLanguageIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiLanguageID, true
}

// SetFkiLanguageID sets field value
func (o *EzsigntemplateRequestCompoundV2) SetFkiLanguageID(v int32) {
	o.FkiLanguageID = v
}

// GetSEzsigntemplateDescription returns the SEzsigntemplateDescription field value
func (o *EzsigntemplateRequestCompoundV2) GetSEzsigntemplateDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzsigntemplateDescription
}

// GetSEzsigntemplateDescriptionOk returns a tuple with the SEzsigntemplateDescription field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateRequestCompoundV2) GetSEzsigntemplateDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SEzsigntemplateDescription, true
}

// SetSEzsigntemplateDescription sets field value
func (o *EzsigntemplateRequestCompoundV2) SetSEzsigntemplateDescription(v string) {
	o.SEzsigntemplateDescription = v
}

// GetSEzsigntemplateFilenamepattern returns the SEzsigntemplateFilenamepattern field value if set, zero value otherwise.
func (o *EzsigntemplateRequestCompoundV2) GetSEzsigntemplateFilenamepattern() string {
	if o == nil || IsNil(o.SEzsigntemplateFilenamepattern) {
		var ret string
		return ret
	}
	return *o.SEzsigntemplateFilenamepattern
}

// GetSEzsigntemplateFilenamepatternOk returns a tuple with the SEzsigntemplateFilenamepattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateRequestCompoundV2) GetSEzsigntemplateFilenamepatternOk() (*string, bool) {
	if o == nil || IsNil(o.SEzsigntemplateFilenamepattern) {
		return nil, false
	}
	return o.SEzsigntemplateFilenamepattern, true
}

// HasSEzsigntemplateFilenamepattern returns a boolean if a field has been set.
func (o *EzsigntemplateRequestCompoundV2) HasSEzsigntemplateFilenamepattern() bool {
	if o != nil && !IsNil(o.SEzsigntemplateFilenamepattern) {
		return true
	}

	return false
}

// SetSEzsigntemplateFilenamepattern gets a reference to the given string and assigns it to the SEzsigntemplateFilenamepattern field.
func (o *EzsigntemplateRequestCompoundV2) SetSEzsigntemplateFilenamepattern(v string) {
	o.SEzsigntemplateFilenamepattern = &v
}

// GetBEzsigntemplateAdminonly returns the BEzsigntemplateAdminonly field value
func (o *EzsigntemplateRequestCompoundV2) GetBEzsigntemplateAdminonly() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BEzsigntemplateAdminonly
}

// GetBEzsigntemplateAdminonlyOk returns a tuple with the BEzsigntemplateAdminonly field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateRequestCompoundV2) GetBEzsigntemplateAdminonlyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BEzsigntemplateAdminonly, true
}

// SetBEzsigntemplateAdminonly sets field value
func (o *EzsigntemplateRequestCompoundV2) SetBEzsigntemplateAdminonly(v bool) {
	o.BEzsigntemplateAdminonly = v
}

// GetEEzsigntemplateType returns the EEzsigntemplateType field value
func (o *EzsigntemplateRequestCompoundV2) GetEEzsigntemplateType() FieldEEzsigntemplateType {
	if o == nil {
		var ret FieldEEzsigntemplateType
		return ret
	}

	return o.EEzsigntemplateType
}

// GetEEzsigntemplateTypeOk returns a tuple with the EEzsigntemplateType field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateRequestCompoundV2) GetEEzsigntemplateTypeOk() (*FieldEEzsigntemplateType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EEzsigntemplateType, true
}

// SetEEzsigntemplateType sets field value
func (o *EzsigntemplateRequestCompoundV2) SetEEzsigntemplateType(v FieldEEzsigntemplateType) {
	o.EEzsigntemplateType = v
}

func (o EzsigntemplateRequestCompoundV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigntemplateRequestCompoundV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PkiEzsigntemplateID) {
		toSerialize["pkiEzsigntemplateID"] = o.PkiEzsigntemplateID
	}
	if !IsNil(o.FkiEzsignfoldertypeID) {
		toSerialize["fkiEzsignfoldertypeID"] = o.FkiEzsignfoldertypeID
	}
	toSerialize["fkiLanguageID"] = o.FkiLanguageID
	toSerialize["sEzsigntemplateDescription"] = o.SEzsigntemplateDescription
	if !IsNil(o.SEzsigntemplateFilenamepattern) {
		toSerialize["sEzsigntemplateFilenamepattern"] = o.SEzsigntemplateFilenamepattern
	}
	toSerialize["bEzsigntemplateAdminonly"] = o.BEzsigntemplateAdminonly
	toSerialize["eEzsigntemplateType"] = o.EEzsigntemplateType
	return toSerialize, nil
}

func (o *EzsigntemplateRequestCompoundV2) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fkiLanguageID",
		"sEzsigntemplateDescription",
		"bEzsigntemplateAdminonly",
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

	varEzsigntemplateRequestCompoundV2 := _EzsigntemplateRequestCompoundV2{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsigntemplateRequestCompoundV2)

	if err != nil {
		return err
	}

	*o = EzsigntemplateRequestCompoundV2(varEzsigntemplateRequestCompoundV2)

	return err
}

type NullableEzsigntemplateRequestCompoundV2 struct {
	value *EzsigntemplateRequestCompoundV2
	isSet bool
}

func (v NullableEzsigntemplateRequestCompoundV2) Get() *EzsigntemplateRequestCompoundV2 {
	return v.value
}

func (v *NullableEzsigntemplateRequestCompoundV2) Set(val *EzsigntemplateRequestCompoundV2) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigntemplateRequestCompoundV2) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigntemplateRequestCompoundV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigntemplateRequestCompoundV2(val *EzsigntemplateRequestCompoundV2) *NullableEzsigntemplateRequestCompoundV2 {
	return &NullableEzsigntemplateRequestCompoundV2{value: val, isSet: true}
}

func (v NullableEzsigntemplateRequestCompoundV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigntemplateRequestCompoundV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


