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

// checks if the VersionhistoryResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VersionhistoryResponse{}

// VersionhistoryResponse A Versionhistory Object
type VersionhistoryResponse struct {
	// The unique ID of the Versionhistory
	PkiVersionhistoryID int32 `json:"pkiVersionhistoryID"`
	// The unique ID of the Module
	FkiModuleID *int32 `json:"fkiModuleID,omitempty"`
	// The unique ID of the Modulesection
	FkiModulesectionID *int32 `json:"fkiModulesectionID,omitempty"`
	// The Name of the Module in the language of the requester
	SModuleNameX *string `json:"sModuleNameX,omitempty"`
	// The Name of the Modulesection in the language of the requester
	SModulesectionNameX *string `json:"sModulesectionNameX,omitempty"`
	EVersionhistoryUsertype *FieldEVersionhistoryUsertype `json:"eVersionhistoryUsertype,omitempty"`
	ObjVersionhistoryDetail MultilingualVersionhistoryDetail `json:"objVersionhistoryDetail"`
	// The date  at which the Versionhistory was published or should be published
	DtVersionhistoryDate string `json:"dtVersionhistoryDate"`
	// The date  at which the Versionhistory will no longer be visible
	DtVersionhistoryDateend *string `json:"dtVersionhistoryDateend,omitempty"`
	EVersionhistoryType FieldEVersionhistoryType `json:"eVersionhistoryType"`
	// Whether the Versionhistory is published or still a draft
	BVersionhistoryDraft bool `json:"bVersionhistoryDraft"`
}

// NewVersionhistoryResponse instantiates a new VersionhistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVersionhistoryResponse(pkiVersionhistoryID int32, objVersionhistoryDetail MultilingualVersionhistoryDetail, dtVersionhistoryDate string, eVersionhistoryType FieldEVersionhistoryType, bVersionhistoryDraft bool) *VersionhistoryResponse {
	this := VersionhistoryResponse{}
	this.PkiVersionhistoryID = pkiVersionhistoryID
	this.ObjVersionhistoryDetail = objVersionhistoryDetail
	this.DtVersionhistoryDate = dtVersionhistoryDate
	this.EVersionhistoryType = eVersionhistoryType
	this.BVersionhistoryDraft = bVersionhistoryDraft
	return &this
}

// NewVersionhistoryResponseWithDefaults instantiates a new VersionhistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVersionhistoryResponseWithDefaults() *VersionhistoryResponse {
	this := VersionhistoryResponse{}
	return &this
}

// GetPkiVersionhistoryID returns the PkiVersionhistoryID field value
func (o *VersionhistoryResponse) GetPkiVersionhistoryID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiVersionhistoryID
}

// GetPkiVersionhistoryIDOk returns a tuple with the PkiVersionhistoryID field value
// and a boolean to check if the value has been set.
func (o *VersionhistoryResponse) GetPkiVersionhistoryIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiVersionhistoryID, true
}

// SetPkiVersionhistoryID sets field value
func (o *VersionhistoryResponse) SetPkiVersionhistoryID(v int32) {
	o.PkiVersionhistoryID = v
}

// GetFkiModuleID returns the FkiModuleID field value if set, zero value otherwise.
func (o *VersionhistoryResponse) GetFkiModuleID() int32 {
	if o == nil || IsNil(o.FkiModuleID) {
		var ret int32
		return ret
	}
	return *o.FkiModuleID
}

// GetFkiModuleIDOk returns a tuple with the FkiModuleID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionhistoryResponse) GetFkiModuleIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiModuleID) {
		return nil, false
	}
	return o.FkiModuleID, true
}

// HasFkiModuleID returns a boolean if a field has been set.
func (o *VersionhistoryResponse) HasFkiModuleID() bool {
	if o != nil && !IsNil(o.FkiModuleID) {
		return true
	}

	return false
}

// SetFkiModuleID gets a reference to the given int32 and assigns it to the FkiModuleID field.
func (o *VersionhistoryResponse) SetFkiModuleID(v int32) {
	o.FkiModuleID = &v
}

// GetFkiModulesectionID returns the FkiModulesectionID field value if set, zero value otherwise.
func (o *VersionhistoryResponse) GetFkiModulesectionID() int32 {
	if o == nil || IsNil(o.FkiModulesectionID) {
		var ret int32
		return ret
	}
	return *o.FkiModulesectionID
}

// GetFkiModulesectionIDOk returns a tuple with the FkiModulesectionID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionhistoryResponse) GetFkiModulesectionIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiModulesectionID) {
		return nil, false
	}
	return o.FkiModulesectionID, true
}

// HasFkiModulesectionID returns a boolean if a field has been set.
func (o *VersionhistoryResponse) HasFkiModulesectionID() bool {
	if o != nil && !IsNil(o.FkiModulesectionID) {
		return true
	}

	return false
}

// SetFkiModulesectionID gets a reference to the given int32 and assigns it to the FkiModulesectionID field.
func (o *VersionhistoryResponse) SetFkiModulesectionID(v int32) {
	o.FkiModulesectionID = &v
}

// GetSModuleNameX returns the SModuleNameX field value if set, zero value otherwise.
func (o *VersionhistoryResponse) GetSModuleNameX() string {
	if o == nil || IsNil(o.SModuleNameX) {
		var ret string
		return ret
	}
	return *o.SModuleNameX
}

// GetSModuleNameXOk returns a tuple with the SModuleNameX field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionhistoryResponse) GetSModuleNameXOk() (*string, bool) {
	if o == nil || IsNil(o.SModuleNameX) {
		return nil, false
	}
	return o.SModuleNameX, true
}

// HasSModuleNameX returns a boolean if a field has been set.
func (o *VersionhistoryResponse) HasSModuleNameX() bool {
	if o != nil && !IsNil(o.SModuleNameX) {
		return true
	}

	return false
}

// SetSModuleNameX gets a reference to the given string and assigns it to the SModuleNameX field.
func (o *VersionhistoryResponse) SetSModuleNameX(v string) {
	o.SModuleNameX = &v
}

// GetSModulesectionNameX returns the SModulesectionNameX field value if set, zero value otherwise.
func (o *VersionhistoryResponse) GetSModulesectionNameX() string {
	if o == nil || IsNil(o.SModulesectionNameX) {
		var ret string
		return ret
	}
	return *o.SModulesectionNameX
}

// GetSModulesectionNameXOk returns a tuple with the SModulesectionNameX field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionhistoryResponse) GetSModulesectionNameXOk() (*string, bool) {
	if o == nil || IsNil(o.SModulesectionNameX) {
		return nil, false
	}
	return o.SModulesectionNameX, true
}

// HasSModulesectionNameX returns a boolean if a field has been set.
func (o *VersionhistoryResponse) HasSModulesectionNameX() bool {
	if o != nil && !IsNil(o.SModulesectionNameX) {
		return true
	}

	return false
}

// SetSModulesectionNameX gets a reference to the given string and assigns it to the SModulesectionNameX field.
func (o *VersionhistoryResponse) SetSModulesectionNameX(v string) {
	o.SModulesectionNameX = &v
}

// GetEVersionhistoryUsertype returns the EVersionhistoryUsertype field value if set, zero value otherwise.
func (o *VersionhistoryResponse) GetEVersionhistoryUsertype() FieldEVersionhistoryUsertype {
	if o == nil || IsNil(o.EVersionhistoryUsertype) {
		var ret FieldEVersionhistoryUsertype
		return ret
	}
	return *o.EVersionhistoryUsertype
}

// GetEVersionhistoryUsertypeOk returns a tuple with the EVersionhistoryUsertype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionhistoryResponse) GetEVersionhistoryUsertypeOk() (*FieldEVersionhistoryUsertype, bool) {
	if o == nil || IsNil(o.EVersionhistoryUsertype) {
		return nil, false
	}
	return o.EVersionhistoryUsertype, true
}

// HasEVersionhistoryUsertype returns a boolean if a field has been set.
func (o *VersionhistoryResponse) HasEVersionhistoryUsertype() bool {
	if o != nil && !IsNil(o.EVersionhistoryUsertype) {
		return true
	}

	return false
}

// SetEVersionhistoryUsertype gets a reference to the given FieldEVersionhistoryUsertype and assigns it to the EVersionhistoryUsertype field.
func (o *VersionhistoryResponse) SetEVersionhistoryUsertype(v FieldEVersionhistoryUsertype) {
	o.EVersionhistoryUsertype = &v
}

// GetObjVersionhistoryDetail returns the ObjVersionhistoryDetail field value
func (o *VersionhistoryResponse) GetObjVersionhistoryDetail() MultilingualVersionhistoryDetail {
	if o == nil {
		var ret MultilingualVersionhistoryDetail
		return ret
	}

	return o.ObjVersionhistoryDetail
}

// GetObjVersionhistoryDetailOk returns a tuple with the ObjVersionhistoryDetail field value
// and a boolean to check if the value has been set.
func (o *VersionhistoryResponse) GetObjVersionhistoryDetailOk() (*MultilingualVersionhistoryDetail, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjVersionhistoryDetail, true
}

// SetObjVersionhistoryDetail sets field value
func (o *VersionhistoryResponse) SetObjVersionhistoryDetail(v MultilingualVersionhistoryDetail) {
	o.ObjVersionhistoryDetail = v
}

// GetDtVersionhistoryDate returns the DtVersionhistoryDate field value
func (o *VersionhistoryResponse) GetDtVersionhistoryDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DtVersionhistoryDate
}

// GetDtVersionhistoryDateOk returns a tuple with the DtVersionhistoryDate field value
// and a boolean to check if the value has been set.
func (o *VersionhistoryResponse) GetDtVersionhistoryDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DtVersionhistoryDate, true
}

// SetDtVersionhistoryDate sets field value
func (o *VersionhistoryResponse) SetDtVersionhistoryDate(v string) {
	o.DtVersionhistoryDate = v
}

// GetDtVersionhistoryDateend returns the DtVersionhistoryDateend field value if set, zero value otherwise.
func (o *VersionhistoryResponse) GetDtVersionhistoryDateend() string {
	if o == nil || IsNil(o.DtVersionhistoryDateend) {
		var ret string
		return ret
	}
	return *o.DtVersionhistoryDateend
}

// GetDtVersionhistoryDateendOk returns a tuple with the DtVersionhistoryDateend field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionhistoryResponse) GetDtVersionhistoryDateendOk() (*string, bool) {
	if o == nil || IsNil(o.DtVersionhistoryDateend) {
		return nil, false
	}
	return o.DtVersionhistoryDateend, true
}

// HasDtVersionhistoryDateend returns a boolean if a field has been set.
func (o *VersionhistoryResponse) HasDtVersionhistoryDateend() bool {
	if o != nil && !IsNil(o.DtVersionhistoryDateend) {
		return true
	}

	return false
}

// SetDtVersionhistoryDateend gets a reference to the given string and assigns it to the DtVersionhistoryDateend field.
func (o *VersionhistoryResponse) SetDtVersionhistoryDateend(v string) {
	o.DtVersionhistoryDateend = &v
}

// GetEVersionhistoryType returns the EVersionhistoryType field value
func (o *VersionhistoryResponse) GetEVersionhistoryType() FieldEVersionhistoryType {
	if o == nil {
		var ret FieldEVersionhistoryType
		return ret
	}

	return o.EVersionhistoryType
}

// GetEVersionhistoryTypeOk returns a tuple with the EVersionhistoryType field value
// and a boolean to check if the value has been set.
func (o *VersionhistoryResponse) GetEVersionhistoryTypeOk() (*FieldEVersionhistoryType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EVersionhistoryType, true
}

// SetEVersionhistoryType sets field value
func (o *VersionhistoryResponse) SetEVersionhistoryType(v FieldEVersionhistoryType) {
	o.EVersionhistoryType = v
}

// GetBVersionhistoryDraft returns the BVersionhistoryDraft field value
func (o *VersionhistoryResponse) GetBVersionhistoryDraft() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BVersionhistoryDraft
}

// GetBVersionhistoryDraftOk returns a tuple with the BVersionhistoryDraft field value
// and a boolean to check if the value has been set.
func (o *VersionhistoryResponse) GetBVersionhistoryDraftOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BVersionhistoryDraft, true
}

// SetBVersionhistoryDraft sets field value
func (o *VersionhistoryResponse) SetBVersionhistoryDraft(v bool) {
	o.BVersionhistoryDraft = v
}

func (o VersionhistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VersionhistoryResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiVersionhistoryID"] = o.PkiVersionhistoryID
	if !IsNil(o.FkiModuleID) {
		toSerialize["fkiModuleID"] = o.FkiModuleID
	}
	if !IsNil(o.FkiModulesectionID) {
		toSerialize["fkiModulesectionID"] = o.FkiModulesectionID
	}
	if !IsNil(o.SModuleNameX) {
		toSerialize["sModuleNameX"] = o.SModuleNameX
	}
	if !IsNil(o.SModulesectionNameX) {
		toSerialize["sModulesectionNameX"] = o.SModulesectionNameX
	}
	if !IsNil(o.EVersionhistoryUsertype) {
		toSerialize["eVersionhistoryUsertype"] = o.EVersionhistoryUsertype
	}
	toSerialize["objVersionhistoryDetail"] = o.ObjVersionhistoryDetail
	toSerialize["dtVersionhistoryDate"] = o.DtVersionhistoryDate
	if !IsNil(o.DtVersionhistoryDateend) {
		toSerialize["dtVersionhistoryDateend"] = o.DtVersionhistoryDateend
	}
	toSerialize["eVersionhistoryType"] = o.EVersionhistoryType
	toSerialize["bVersionhistoryDraft"] = o.BVersionhistoryDraft
	return toSerialize, nil
}

type NullableVersionhistoryResponse struct {
	value *VersionhistoryResponse
	isSet bool
}

func (v NullableVersionhistoryResponse) Get() *VersionhistoryResponse {
	return v.value
}

func (v *NullableVersionhistoryResponse) Set(val *VersionhistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableVersionhistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableVersionhistoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVersionhistoryResponse(val *VersionhistoryResponse) *NullableVersionhistoryResponse {
	return &NullableVersionhistoryResponse{value: val, isSet: true}
}

func (v NullableVersionhistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVersionhistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

