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

// checks if the EzsigntemplateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigntemplateRequest{}

// EzsigntemplateRequest A Ezsigntemplate Object
type EzsigntemplateRequest struct {
	// The unique ID of the Ezsigntemplate
	PkiEzsigntemplateID *int32 `json:"pkiEzsigntemplateID,omitempty"`
	// The unique ID of the Ezsignfoldertype.
	FkiEzsignfoldertypeID int32 `json:"fkiEzsignfoldertypeID"`
	// The unique ID of the Language.  Valid values:  |Value|Description| |-|-| |1|French| |2|English|
	FkiLanguageID int32 `json:"fkiLanguageID"`
	// The description of the Ezsigntemplate
	SEzsigntemplateDescription string `json:"sEzsigntemplateDescription"`
	// Whether the Ezsigntemplate can be accessed by admin users only (eUserType=Normal)
	BEzsigntemplateAdminonly bool `json:"bEzsigntemplateAdminonly"`
}

// NewEzsigntemplateRequest instantiates a new EzsigntemplateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigntemplateRequest(fkiEzsignfoldertypeID int32, fkiLanguageID int32, sEzsigntemplateDescription string, bEzsigntemplateAdminonly bool) *EzsigntemplateRequest {
	this := EzsigntemplateRequest{}
	this.FkiEzsignfoldertypeID = fkiEzsignfoldertypeID
	this.FkiLanguageID = fkiLanguageID
	this.SEzsigntemplateDescription = sEzsigntemplateDescription
	this.BEzsigntemplateAdminonly = bEzsigntemplateAdminonly
	return &this
}

// NewEzsigntemplateRequestWithDefaults instantiates a new EzsigntemplateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigntemplateRequestWithDefaults() *EzsigntemplateRequest {
	this := EzsigntemplateRequest{}
	return &this
}

// GetPkiEzsigntemplateID returns the PkiEzsigntemplateID field value if set, zero value otherwise.
func (o *EzsigntemplateRequest) GetPkiEzsigntemplateID() int32 {
	if o == nil || IsNil(o.PkiEzsigntemplateID) {
		var ret int32
		return ret
	}
	return *o.PkiEzsigntemplateID
}

// GetPkiEzsigntemplateIDOk returns a tuple with the PkiEzsigntemplateID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateRequest) GetPkiEzsigntemplateIDOk() (*int32, bool) {
	if o == nil || IsNil(o.PkiEzsigntemplateID) {
		return nil, false
	}
	return o.PkiEzsigntemplateID, true
}

// HasPkiEzsigntemplateID returns a boolean if a field has been set.
func (o *EzsigntemplateRequest) HasPkiEzsigntemplateID() bool {
	if o != nil && !IsNil(o.PkiEzsigntemplateID) {
		return true
	}

	return false
}

// SetPkiEzsigntemplateID gets a reference to the given int32 and assigns it to the PkiEzsigntemplateID field.
func (o *EzsigntemplateRequest) SetPkiEzsigntemplateID(v int32) {
	o.PkiEzsigntemplateID = &v
}

// GetFkiEzsignfoldertypeID returns the FkiEzsignfoldertypeID field value
func (o *EzsigntemplateRequest) GetFkiEzsignfoldertypeID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzsignfoldertypeID
}

// GetFkiEzsignfoldertypeIDOk returns a tuple with the FkiEzsignfoldertypeID field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateRequest) GetFkiEzsignfoldertypeIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiEzsignfoldertypeID, true
}

// SetFkiEzsignfoldertypeID sets field value
func (o *EzsigntemplateRequest) SetFkiEzsignfoldertypeID(v int32) {
	o.FkiEzsignfoldertypeID = v
}

// GetFkiLanguageID returns the FkiLanguageID field value
func (o *EzsigntemplateRequest) GetFkiLanguageID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiLanguageID
}

// GetFkiLanguageIDOk returns a tuple with the FkiLanguageID field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateRequest) GetFkiLanguageIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiLanguageID, true
}

// SetFkiLanguageID sets field value
func (o *EzsigntemplateRequest) SetFkiLanguageID(v int32) {
	o.FkiLanguageID = v
}

// GetSEzsigntemplateDescription returns the SEzsigntemplateDescription field value
func (o *EzsigntemplateRequest) GetSEzsigntemplateDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzsigntemplateDescription
}

// GetSEzsigntemplateDescriptionOk returns a tuple with the SEzsigntemplateDescription field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateRequest) GetSEzsigntemplateDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SEzsigntemplateDescription, true
}

// SetSEzsigntemplateDescription sets field value
func (o *EzsigntemplateRequest) SetSEzsigntemplateDescription(v string) {
	o.SEzsigntemplateDescription = v
}

// GetBEzsigntemplateAdminonly returns the BEzsigntemplateAdminonly field value
func (o *EzsigntemplateRequest) GetBEzsigntemplateAdminonly() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BEzsigntemplateAdminonly
}

// GetBEzsigntemplateAdminonlyOk returns a tuple with the BEzsigntemplateAdminonly field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateRequest) GetBEzsigntemplateAdminonlyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BEzsigntemplateAdminonly, true
}

// SetBEzsigntemplateAdminonly sets field value
func (o *EzsigntemplateRequest) SetBEzsigntemplateAdminonly(v bool) {
	o.BEzsigntemplateAdminonly = v
}

func (o EzsigntemplateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigntemplateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PkiEzsigntemplateID) {
		toSerialize["pkiEzsigntemplateID"] = o.PkiEzsigntemplateID
	}
	toSerialize["fkiEzsignfoldertypeID"] = o.FkiEzsignfoldertypeID
	toSerialize["fkiLanguageID"] = o.FkiLanguageID
	toSerialize["sEzsigntemplateDescription"] = o.SEzsigntemplateDescription
	toSerialize["bEzsigntemplateAdminonly"] = o.BEzsigntemplateAdminonly
	return toSerialize, nil
}

type NullableEzsigntemplateRequest struct {
	value *EzsigntemplateRequest
	isSet bool
}

func (v NullableEzsigntemplateRequest) Get() *EzsigntemplateRequest {
	return v.value
}

func (v *NullableEzsigntemplateRequest) Set(val *EzsigntemplateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigntemplateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigntemplateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigntemplateRequest(val *EzsigntemplateRequest) *NullableEzsigntemplateRequest {
	return &NullableEzsigntemplateRequest{value: val, isSet: true}
}

func (v NullableEzsigntemplateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigntemplateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


