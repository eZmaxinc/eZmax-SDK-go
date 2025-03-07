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

// checks if the EzsigntemplatepackageRequestCompound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigntemplatepackageRequestCompound{}

// EzsigntemplatepackageRequestCompound A Ezsigntemplatepackage Object and children
type EzsigntemplatepackageRequestCompound struct {
	// The unique ID of the Ezsigntemplatepackage
	PkiEzsigntemplatepackageID *int32 `json:"pkiEzsigntemplatepackageID,omitempty"`
	// The unique ID of the Ezsignfoldertype.
	FkiEzsignfoldertypeID int32 `json:"fkiEzsignfoldertypeID"`
	// The unique ID of the Ezdoctemplatedocument
	FkiEzdoctemplatedocumentID *int32 `json:"fkiEzdoctemplatedocumentID,omitempty"`
	// The unique ID of the Language.  Valid values:  |Value|Description| |-|-| |1|French| |2|English|
	FkiLanguageID int32 `json:"fkiLanguageID"`
	// The description of the Ezsigntemplatepackage
	SEzsigntemplatepackageDescription string `json:"sEzsigntemplatepackageDescription" validate:"regexp=^.{0,80}$"`
	// Whether the Ezsigntemplatepackage can be accessed by admin users only (eUserType=Normal)
	BEzsigntemplatepackageAdminonly bool `json:"bEzsigntemplatepackageAdminonly"`
	// Whether the Ezsigntemplatepackage is active or not
	BEzsigntemplatepackageIsactive bool `json:"bEzsigntemplatepackageIsactive"`
}

type _EzsigntemplatepackageRequestCompound EzsigntemplatepackageRequestCompound

// NewEzsigntemplatepackageRequestCompound instantiates a new EzsigntemplatepackageRequestCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigntemplatepackageRequestCompound(fkiEzsignfoldertypeID int32, fkiLanguageID int32, sEzsigntemplatepackageDescription string, bEzsigntemplatepackageAdminonly bool, bEzsigntemplatepackageIsactive bool) *EzsigntemplatepackageRequestCompound {
	this := EzsigntemplatepackageRequestCompound{}
	this.FkiEzsignfoldertypeID = fkiEzsignfoldertypeID
	this.FkiLanguageID = fkiLanguageID
	this.SEzsigntemplatepackageDescription = sEzsigntemplatepackageDescription
	this.BEzsigntemplatepackageAdminonly = bEzsigntemplatepackageAdminonly
	this.BEzsigntemplatepackageIsactive = bEzsigntemplatepackageIsactive
	return &this
}

// NewEzsigntemplatepackageRequestCompoundWithDefaults instantiates a new EzsigntemplatepackageRequestCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigntemplatepackageRequestCompoundWithDefaults() *EzsigntemplatepackageRequestCompound {
	this := EzsigntemplatepackageRequestCompound{}
	return &this
}

// GetPkiEzsigntemplatepackageID returns the PkiEzsigntemplatepackageID field value if set, zero value otherwise.
func (o *EzsigntemplatepackageRequestCompound) GetPkiEzsigntemplatepackageID() int32 {
	if o == nil || IsNil(o.PkiEzsigntemplatepackageID) {
		var ret int32
		return ret
	}
	return *o.PkiEzsigntemplatepackageID
}

// GetPkiEzsigntemplatepackageIDOk returns a tuple with the PkiEzsigntemplatepackageID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepackageRequestCompound) GetPkiEzsigntemplatepackageIDOk() (*int32, bool) {
	if o == nil || IsNil(o.PkiEzsigntemplatepackageID) {
		return nil, false
	}
	return o.PkiEzsigntemplatepackageID, true
}

// HasPkiEzsigntemplatepackageID returns a boolean if a field has been set.
func (o *EzsigntemplatepackageRequestCompound) HasPkiEzsigntemplatepackageID() bool {
	if o != nil && !IsNil(o.PkiEzsigntemplatepackageID) {
		return true
	}

	return false
}

// SetPkiEzsigntemplatepackageID gets a reference to the given int32 and assigns it to the PkiEzsigntemplatepackageID field.
func (o *EzsigntemplatepackageRequestCompound) SetPkiEzsigntemplatepackageID(v int32) {
	o.PkiEzsigntemplatepackageID = &v
}

// GetFkiEzsignfoldertypeID returns the FkiEzsignfoldertypeID field value
func (o *EzsigntemplatepackageRequestCompound) GetFkiEzsignfoldertypeID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzsignfoldertypeID
}

// GetFkiEzsignfoldertypeIDOk returns a tuple with the FkiEzsignfoldertypeID field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepackageRequestCompound) GetFkiEzsignfoldertypeIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiEzsignfoldertypeID, true
}

// SetFkiEzsignfoldertypeID sets field value
func (o *EzsigntemplatepackageRequestCompound) SetFkiEzsignfoldertypeID(v int32) {
	o.FkiEzsignfoldertypeID = v
}

// GetFkiEzdoctemplatedocumentID returns the FkiEzdoctemplatedocumentID field value if set, zero value otherwise.
func (o *EzsigntemplatepackageRequestCompound) GetFkiEzdoctemplatedocumentID() int32 {
	if o == nil || IsNil(o.FkiEzdoctemplatedocumentID) {
		var ret int32
		return ret
	}
	return *o.FkiEzdoctemplatedocumentID
}

// GetFkiEzdoctemplatedocumentIDOk returns a tuple with the FkiEzdoctemplatedocumentID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepackageRequestCompound) GetFkiEzdoctemplatedocumentIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiEzdoctemplatedocumentID) {
		return nil, false
	}
	return o.FkiEzdoctemplatedocumentID, true
}

// HasFkiEzdoctemplatedocumentID returns a boolean if a field has been set.
func (o *EzsigntemplatepackageRequestCompound) HasFkiEzdoctemplatedocumentID() bool {
	if o != nil && !IsNil(o.FkiEzdoctemplatedocumentID) {
		return true
	}

	return false
}

// SetFkiEzdoctemplatedocumentID gets a reference to the given int32 and assigns it to the FkiEzdoctemplatedocumentID field.
func (o *EzsigntemplatepackageRequestCompound) SetFkiEzdoctemplatedocumentID(v int32) {
	o.FkiEzdoctemplatedocumentID = &v
}

// GetFkiLanguageID returns the FkiLanguageID field value
func (o *EzsigntemplatepackageRequestCompound) GetFkiLanguageID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiLanguageID
}

// GetFkiLanguageIDOk returns a tuple with the FkiLanguageID field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepackageRequestCompound) GetFkiLanguageIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiLanguageID, true
}

// SetFkiLanguageID sets field value
func (o *EzsigntemplatepackageRequestCompound) SetFkiLanguageID(v int32) {
	o.FkiLanguageID = v
}

// GetSEzsigntemplatepackageDescription returns the SEzsigntemplatepackageDescription field value
func (o *EzsigntemplatepackageRequestCompound) GetSEzsigntemplatepackageDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzsigntemplatepackageDescription
}

// GetSEzsigntemplatepackageDescriptionOk returns a tuple with the SEzsigntemplatepackageDescription field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepackageRequestCompound) GetSEzsigntemplatepackageDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SEzsigntemplatepackageDescription, true
}

// SetSEzsigntemplatepackageDescription sets field value
func (o *EzsigntemplatepackageRequestCompound) SetSEzsigntemplatepackageDescription(v string) {
	o.SEzsigntemplatepackageDescription = v
}

// GetBEzsigntemplatepackageAdminonly returns the BEzsigntemplatepackageAdminonly field value
func (o *EzsigntemplatepackageRequestCompound) GetBEzsigntemplatepackageAdminonly() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BEzsigntemplatepackageAdminonly
}

// GetBEzsigntemplatepackageAdminonlyOk returns a tuple with the BEzsigntemplatepackageAdminonly field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepackageRequestCompound) GetBEzsigntemplatepackageAdminonlyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BEzsigntemplatepackageAdminonly, true
}

// SetBEzsigntemplatepackageAdminonly sets field value
func (o *EzsigntemplatepackageRequestCompound) SetBEzsigntemplatepackageAdminonly(v bool) {
	o.BEzsigntemplatepackageAdminonly = v
}

// GetBEzsigntemplatepackageIsactive returns the BEzsigntemplatepackageIsactive field value
func (o *EzsigntemplatepackageRequestCompound) GetBEzsigntemplatepackageIsactive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BEzsigntemplatepackageIsactive
}

// GetBEzsigntemplatepackageIsactiveOk returns a tuple with the BEzsigntemplatepackageIsactive field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepackageRequestCompound) GetBEzsigntemplatepackageIsactiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BEzsigntemplatepackageIsactive, true
}

// SetBEzsigntemplatepackageIsactive sets field value
func (o *EzsigntemplatepackageRequestCompound) SetBEzsigntemplatepackageIsactive(v bool) {
	o.BEzsigntemplatepackageIsactive = v
}

func (o EzsigntemplatepackageRequestCompound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigntemplatepackageRequestCompound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PkiEzsigntemplatepackageID) {
		toSerialize["pkiEzsigntemplatepackageID"] = o.PkiEzsigntemplatepackageID
	}
	toSerialize["fkiEzsignfoldertypeID"] = o.FkiEzsignfoldertypeID
	if !IsNil(o.FkiEzdoctemplatedocumentID) {
		toSerialize["fkiEzdoctemplatedocumentID"] = o.FkiEzdoctemplatedocumentID
	}
	toSerialize["fkiLanguageID"] = o.FkiLanguageID
	toSerialize["sEzsigntemplatepackageDescription"] = o.SEzsigntemplatepackageDescription
	toSerialize["bEzsigntemplatepackageAdminonly"] = o.BEzsigntemplatepackageAdminonly
	toSerialize["bEzsigntemplatepackageIsactive"] = o.BEzsigntemplatepackageIsactive
	return toSerialize, nil
}

func (o *EzsigntemplatepackageRequestCompound) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fkiEzsignfoldertypeID",
		"fkiLanguageID",
		"sEzsigntemplatepackageDescription",
		"bEzsigntemplatepackageAdminonly",
		"bEzsigntemplatepackageIsactive",
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

	varEzsigntemplatepackageRequestCompound := _EzsigntemplatepackageRequestCompound{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsigntemplatepackageRequestCompound)

	if err != nil {
		return err
	}

	*o = EzsigntemplatepackageRequestCompound(varEzsigntemplatepackageRequestCompound)

	return err
}

type NullableEzsigntemplatepackageRequestCompound struct {
	value *EzsigntemplatepackageRequestCompound
	isSet bool
}

func (v NullableEzsigntemplatepackageRequestCompound) Get() *EzsigntemplatepackageRequestCompound {
	return v.value
}

func (v *NullableEzsigntemplatepackageRequestCompound) Set(val *EzsigntemplatepackageRequestCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigntemplatepackageRequestCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigntemplatepackageRequestCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigntemplatepackageRequestCompound(val *EzsigntemplatepackageRequestCompound) *NullableEzsigntemplatepackageRequestCompound {
	return &NullableEzsigntemplatepackageRequestCompound{value: val, isSet: true}
}

func (v NullableEzsigntemplatepackageRequestCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigntemplatepackageRequestCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


