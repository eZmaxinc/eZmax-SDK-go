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
	"bytes"
	"fmt"
)

// checks if the EzsigntemplateRequestCompound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigntemplateRequestCompound{}

// EzsigntemplateRequestCompound A Ezsigntemplate Object and children
type EzsigntemplateRequestCompound struct {
	// The unique ID of the Ezsigntemplate
	PkiEzsigntemplateID *int32 `json:"pkiEzsigntemplateID,omitempty"`
	// The unique ID of the Ezsignfoldertype.
	FkiEzsignfoldertypeID int32 `json:"fkiEzsignfoldertypeID"`
	// The unique ID of the Language.  Valid values:  |Value|Description| |-|-| |1|French| |2|English|
	FkiLanguageID int32 `json:"fkiLanguageID"`
	// The description of the Ezsigntemplate
	SEzsigntemplateDescription string `json:"sEzsigntemplateDescription"`
	// The filename pattern of the Ezsigntemplate
	SEzsigntemplateFilenamepattern *string `json:"sEzsigntemplateFilenamepattern,omitempty"`
	// Whether the Ezsigntemplate can be accessed by admin users only (eUserType=Normal)
	BEzsigntemplateAdminonly bool `json:"bEzsigntemplateAdminonly"`
}

type _EzsigntemplateRequestCompound EzsigntemplateRequestCompound

// NewEzsigntemplateRequestCompound instantiates a new EzsigntemplateRequestCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigntemplateRequestCompound(fkiEzsignfoldertypeID int32, fkiLanguageID int32, sEzsigntemplateDescription string, bEzsigntemplateAdminonly bool) *EzsigntemplateRequestCompound {
	this := EzsigntemplateRequestCompound{}
	this.FkiEzsignfoldertypeID = fkiEzsignfoldertypeID
	this.FkiLanguageID = fkiLanguageID
	this.SEzsigntemplateDescription = sEzsigntemplateDescription
	this.BEzsigntemplateAdminonly = bEzsigntemplateAdminonly
	return &this
}

// NewEzsigntemplateRequestCompoundWithDefaults instantiates a new EzsigntemplateRequestCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigntemplateRequestCompoundWithDefaults() *EzsigntemplateRequestCompound {
	this := EzsigntemplateRequestCompound{}
	return &this
}

// GetPkiEzsigntemplateID returns the PkiEzsigntemplateID field value if set, zero value otherwise.
func (o *EzsigntemplateRequestCompound) GetPkiEzsigntemplateID() int32 {
	if o == nil || IsNil(o.PkiEzsigntemplateID) {
		var ret int32
		return ret
	}
	return *o.PkiEzsigntemplateID
}

// GetPkiEzsigntemplateIDOk returns a tuple with the PkiEzsigntemplateID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateRequestCompound) GetPkiEzsigntemplateIDOk() (*int32, bool) {
	if o == nil || IsNil(o.PkiEzsigntemplateID) {
		return nil, false
	}
	return o.PkiEzsigntemplateID, true
}

// HasPkiEzsigntemplateID returns a boolean if a field has been set.
func (o *EzsigntemplateRequestCompound) HasPkiEzsigntemplateID() bool {
	if o != nil && !IsNil(o.PkiEzsigntemplateID) {
		return true
	}

	return false
}

// SetPkiEzsigntemplateID gets a reference to the given int32 and assigns it to the PkiEzsigntemplateID field.
func (o *EzsigntemplateRequestCompound) SetPkiEzsigntemplateID(v int32) {
	o.PkiEzsigntemplateID = &v
}

// GetFkiEzsignfoldertypeID returns the FkiEzsignfoldertypeID field value
func (o *EzsigntemplateRequestCompound) GetFkiEzsignfoldertypeID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzsignfoldertypeID
}

// GetFkiEzsignfoldertypeIDOk returns a tuple with the FkiEzsignfoldertypeID field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateRequestCompound) GetFkiEzsignfoldertypeIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiEzsignfoldertypeID, true
}

// SetFkiEzsignfoldertypeID sets field value
func (o *EzsigntemplateRequestCompound) SetFkiEzsignfoldertypeID(v int32) {
	o.FkiEzsignfoldertypeID = v
}

// GetFkiLanguageID returns the FkiLanguageID field value
func (o *EzsigntemplateRequestCompound) GetFkiLanguageID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiLanguageID
}

// GetFkiLanguageIDOk returns a tuple with the FkiLanguageID field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateRequestCompound) GetFkiLanguageIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiLanguageID, true
}

// SetFkiLanguageID sets field value
func (o *EzsigntemplateRequestCompound) SetFkiLanguageID(v int32) {
	o.FkiLanguageID = v
}

// GetSEzsigntemplateDescription returns the SEzsigntemplateDescription field value
func (o *EzsigntemplateRequestCompound) GetSEzsigntemplateDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzsigntemplateDescription
}

// GetSEzsigntemplateDescriptionOk returns a tuple with the SEzsigntemplateDescription field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateRequestCompound) GetSEzsigntemplateDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SEzsigntemplateDescription, true
}

// SetSEzsigntemplateDescription sets field value
func (o *EzsigntemplateRequestCompound) SetSEzsigntemplateDescription(v string) {
	o.SEzsigntemplateDescription = v
}

// GetSEzsigntemplateFilenamepattern returns the SEzsigntemplateFilenamepattern field value if set, zero value otherwise.
func (o *EzsigntemplateRequestCompound) GetSEzsigntemplateFilenamepattern() string {
	if o == nil || IsNil(o.SEzsigntemplateFilenamepattern) {
		var ret string
		return ret
	}
	return *o.SEzsigntemplateFilenamepattern
}

// GetSEzsigntemplateFilenamepatternOk returns a tuple with the SEzsigntemplateFilenamepattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateRequestCompound) GetSEzsigntemplateFilenamepatternOk() (*string, bool) {
	if o == nil || IsNil(o.SEzsigntemplateFilenamepattern) {
		return nil, false
	}
	return o.SEzsigntemplateFilenamepattern, true
}

// HasSEzsigntemplateFilenamepattern returns a boolean if a field has been set.
func (o *EzsigntemplateRequestCompound) HasSEzsigntemplateFilenamepattern() bool {
	if o != nil && !IsNil(o.SEzsigntemplateFilenamepattern) {
		return true
	}

	return false
}

// SetSEzsigntemplateFilenamepattern gets a reference to the given string and assigns it to the SEzsigntemplateFilenamepattern field.
func (o *EzsigntemplateRequestCompound) SetSEzsigntemplateFilenamepattern(v string) {
	o.SEzsigntemplateFilenamepattern = &v
}

// GetBEzsigntemplateAdminonly returns the BEzsigntemplateAdminonly field value
func (o *EzsigntemplateRequestCompound) GetBEzsigntemplateAdminonly() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BEzsigntemplateAdminonly
}

// GetBEzsigntemplateAdminonlyOk returns a tuple with the BEzsigntemplateAdminonly field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateRequestCompound) GetBEzsigntemplateAdminonlyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BEzsigntemplateAdminonly, true
}

// SetBEzsigntemplateAdminonly sets field value
func (o *EzsigntemplateRequestCompound) SetBEzsigntemplateAdminonly(v bool) {
	o.BEzsigntemplateAdminonly = v
}

func (o EzsigntemplateRequestCompound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigntemplateRequestCompound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PkiEzsigntemplateID) {
		toSerialize["pkiEzsigntemplateID"] = o.PkiEzsigntemplateID
	}
	toSerialize["fkiEzsignfoldertypeID"] = o.FkiEzsignfoldertypeID
	toSerialize["fkiLanguageID"] = o.FkiLanguageID
	toSerialize["sEzsigntemplateDescription"] = o.SEzsigntemplateDescription
	if !IsNil(o.SEzsigntemplateFilenamepattern) {
		toSerialize["sEzsigntemplateFilenamepattern"] = o.SEzsigntemplateFilenamepattern
	}
	toSerialize["bEzsigntemplateAdminonly"] = o.BEzsigntemplateAdminonly
	return toSerialize, nil
}

func (o *EzsigntemplateRequestCompound) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fkiEzsignfoldertypeID",
		"fkiLanguageID",
		"sEzsigntemplateDescription",
		"bEzsigntemplateAdminonly",
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

	varEzsigntemplateRequestCompound := _EzsigntemplateRequestCompound{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsigntemplateRequestCompound)

	if err != nil {
		return err
	}

	*o = EzsigntemplateRequestCompound(varEzsigntemplateRequestCompound)

	return err
}

type NullableEzsigntemplateRequestCompound struct {
	value *EzsigntemplateRequestCompound
	isSet bool
}

func (v NullableEzsigntemplateRequestCompound) Get() *EzsigntemplateRequestCompound {
	return v.value
}

func (v *NullableEzsigntemplateRequestCompound) Set(val *EzsigntemplateRequestCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigntemplateRequestCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigntemplateRequestCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigntemplateRequestCompound(val *EzsigntemplateRequestCompound) *NullableEzsigntemplateRequestCompound {
	return &NullableEzsigntemplateRequestCompound{value: val, isSet: true}
}

func (v NullableEzsigntemplateRequestCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigntemplateRequestCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


