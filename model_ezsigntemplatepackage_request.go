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

// checks if the EzsigntemplatepackageRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigntemplatepackageRequest{}

// EzsigntemplatepackageRequest A Ezsigntemplatepackage Object
type EzsigntemplatepackageRequest struct {
	// The unique ID of the Ezsigntemplatepackage
	PkiEzsigntemplatepackageID *int32 `json:"pkiEzsigntemplatepackageID,omitempty"`
	// The unique ID of the Ezsignfoldertype.
	FkiEzsignfoldertypeID int32 `json:"fkiEzsignfoldertypeID"`
	// The unique ID of the Language.  Valid values:  |Value|Description| |-|-| |1|French| |2|English|
	FkiLanguageID int32 `json:"fkiLanguageID"`
	// The description of the Ezsigntemplatepackage
	SEzsigntemplatepackageDescription string `json:"sEzsigntemplatepackageDescription"`
	// Whether the Ezsigntemplatepackage can be accessed by admin users only (eUserType=Normal)
	BEzsigntemplatepackageAdminonly bool `json:"bEzsigntemplatepackageAdminonly"`
	// Whether the Ezsigntemplatepackage is active or not
	BEzsigntemplatepackageIsactive bool `json:"bEzsigntemplatepackageIsactive"`
}

// NewEzsigntemplatepackageRequest instantiates a new EzsigntemplatepackageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigntemplatepackageRequest(fkiEzsignfoldertypeID int32, fkiLanguageID int32, sEzsigntemplatepackageDescription string, bEzsigntemplatepackageAdminonly bool, bEzsigntemplatepackageIsactive bool) *EzsigntemplatepackageRequest {
	this := EzsigntemplatepackageRequest{}
	this.FkiEzsignfoldertypeID = fkiEzsignfoldertypeID
	this.FkiLanguageID = fkiLanguageID
	this.SEzsigntemplatepackageDescription = sEzsigntemplatepackageDescription
	this.BEzsigntemplatepackageAdminonly = bEzsigntemplatepackageAdminonly
	this.BEzsigntemplatepackageIsactive = bEzsigntemplatepackageIsactive
	return &this
}

// NewEzsigntemplatepackageRequestWithDefaults instantiates a new EzsigntemplatepackageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigntemplatepackageRequestWithDefaults() *EzsigntemplatepackageRequest {
	this := EzsigntemplatepackageRequest{}
	return &this
}

// GetPkiEzsigntemplatepackageID returns the PkiEzsigntemplatepackageID field value if set, zero value otherwise.
func (o *EzsigntemplatepackageRequest) GetPkiEzsigntemplatepackageID() int32 {
	if o == nil || IsNil(o.PkiEzsigntemplatepackageID) {
		var ret int32
		return ret
	}
	return *o.PkiEzsigntemplatepackageID
}

// GetPkiEzsigntemplatepackageIDOk returns a tuple with the PkiEzsigntemplatepackageID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepackageRequest) GetPkiEzsigntemplatepackageIDOk() (*int32, bool) {
	if o == nil || IsNil(o.PkiEzsigntemplatepackageID) {
		return nil, false
	}
	return o.PkiEzsigntemplatepackageID, true
}

// HasPkiEzsigntemplatepackageID returns a boolean if a field has been set.
func (o *EzsigntemplatepackageRequest) HasPkiEzsigntemplatepackageID() bool {
	if o != nil && !IsNil(o.PkiEzsigntemplatepackageID) {
		return true
	}

	return false
}

// SetPkiEzsigntemplatepackageID gets a reference to the given int32 and assigns it to the PkiEzsigntemplatepackageID field.
func (o *EzsigntemplatepackageRequest) SetPkiEzsigntemplatepackageID(v int32) {
	o.PkiEzsigntemplatepackageID = &v
}

// GetFkiEzsignfoldertypeID returns the FkiEzsignfoldertypeID field value
func (o *EzsigntemplatepackageRequest) GetFkiEzsignfoldertypeID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzsignfoldertypeID
}

// GetFkiEzsignfoldertypeIDOk returns a tuple with the FkiEzsignfoldertypeID field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepackageRequest) GetFkiEzsignfoldertypeIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiEzsignfoldertypeID, true
}

// SetFkiEzsignfoldertypeID sets field value
func (o *EzsigntemplatepackageRequest) SetFkiEzsignfoldertypeID(v int32) {
	o.FkiEzsignfoldertypeID = v
}

// GetFkiLanguageID returns the FkiLanguageID field value
func (o *EzsigntemplatepackageRequest) GetFkiLanguageID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiLanguageID
}

// GetFkiLanguageIDOk returns a tuple with the FkiLanguageID field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepackageRequest) GetFkiLanguageIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiLanguageID, true
}

// SetFkiLanguageID sets field value
func (o *EzsigntemplatepackageRequest) SetFkiLanguageID(v int32) {
	o.FkiLanguageID = v
}

// GetSEzsigntemplatepackageDescription returns the SEzsigntemplatepackageDescription field value
func (o *EzsigntemplatepackageRequest) GetSEzsigntemplatepackageDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzsigntemplatepackageDescription
}

// GetSEzsigntemplatepackageDescriptionOk returns a tuple with the SEzsigntemplatepackageDescription field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepackageRequest) GetSEzsigntemplatepackageDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SEzsigntemplatepackageDescription, true
}

// SetSEzsigntemplatepackageDescription sets field value
func (o *EzsigntemplatepackageRequest) SetSEzsigntemplatepackageDescription(v string) {
	o.SEzsigntemplatepackageDescription = v
}

// GetBEzsigntemplatepackageAdminonly returns the BEzsigntemplatepackageAdminonly field value
func (o *EzsigntemplatepackageRequest) GetBEzsigntemplatepackageAdminonly() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BEzsigntemplatepackageAdminonly
}

// GetBEzsigntemplatepackageAdminonlyOk returns a tuple with the BEzsigntemplatepackageAdminonly field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepackageRequest) GetBEzsigntemplatepackageAdminonlyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BEzsigntemplatepackageAdminonly, true
}

// SetBEzsigntemplatepackageAdminonly sets field value
func (o *EzsigntemplatepackageRequest) SetBEzsigntemplatepackageAdminonly(v bool) {
	o.BEzsigntemplatepackageAdminonly = v
}

// GetBEzsigntemplatepackageIsactive returns the BEzsigntemplatepackageIsactive field value
func (o *EzsigntemplatepackageRequest) GetBEzsigntemplatepackageIsactive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BEzsigntemplatepackageIsactive
}

// GetBEzsigntemplatepackageIsactiveOk returns a tuple with the BEzsigntemplatepackageIsactive field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepackageRequest) GetBEzsigntemplatepackageIsactiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BEzsigntemplatepackageIsactive, true
}

// SetBEzsigntemplatepackageIsactive sets field value
func (o *EzsigntemplatepackageRequest) SetBEzsigntemplatepackageIsactive(v bool) {
	o.BEzsigntemplatepackageIsactive = v
}

func (o EzsigntemplatepackageRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigntemplatepackageRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PkiEzsigntemplatepackageID) {
		toSerialize["pkiEzsigntemplatepackageID"] = o.PkiEzsigntemplatepackageID
	}
	toSerialize["fkiEzsignfoldertypeID"] = o.FkiEzsignfoldertypeID
	toSerialize["fkiLanguageID"] = o.FkiLanguageID
	toSerialize["sEzsigntemplatepackageDescription"] = o.SEzsigntemplatepackageDescription
	toSerialize["bEzsigntemplatepackageAdminonly"] = o.BEzsigntemplatepackageAdminonly
	toSerialize["bEzsigntemplatepackageIsactive"] = o.BEzsigntemplatepackageIsactive
	return toSerialize, nil
}

type NullableEzsigntemplatepackageRequest struct {
	value *EzsigntemplatepackageRequest
	isSet bool
}

func (v NullableEzsigntemplatepackageRequest) Get() *EzsigntemplatepackageRequest {
	return v.value
}

func (v *NullableEzsigntemplatepackageRequest) Set(val *EzsigntemplatepackageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigntemplatepackageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigntemplatepackageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigntemplatepackageRequest(val *EzsigntemplatepackageRequest) *NullableEzsigntemplatepackageRequest {
	return &NullableEzsigntemplatepackageRequest{value: val, isSet: true}
}

func (v NullableEzsigntemplatepackageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigntemplatepackageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


