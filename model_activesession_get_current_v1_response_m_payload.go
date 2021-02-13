/*
 * eZmax API Definition
 *
 * This API expose all the functionnalities for the eZmax and eZsign application.
 *
 * API version: 1.0.30
 * Contact: support-api@ezmax.ca
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxinc/eZmax-SDK-go

import (
	"encoding/json"
)

// ActivesessionGetCurrentV1ResponseMPayload Payload for the /1/object/activesession/getCurrent API Request
type ActivesessionGetCurrentV1ResponseMPayload struct {
	// The customer code specific to the client in which the API request is being made
	SCustomerCode string `json:"sCustomerCode"`
	// The type of session used for the API request call
	EActivesessionSessiontype string `json:"eActivesessionSessiontype"`
	// The unique ID of the Language.  Valid values:  |Value|Description| |-|-| |1|French| |2|English|
	FkiLanguageID int32 `json:"fkiLanguageID"`
	// The name of the active Company in the current language
	SCompanyNameX string `json:"sCompanyNameX"`
	// The name of the active Department in the current language
	SDepartmentNameX string `json:"sDepartmentNameX"`
	// An Array of Registered modules.  These are the modules that are Licensed to be used by the User or the API Key.
	ARegisteredModules []string `json:"a_RegisteredModules"`
	// An array of permissions granted to the user or api key
	APermissions []int32 `json:"a_Permissions"`
}

// NewActivesessionGetCurrentV1ResponseMPayload instantiates a new ActivesessionGetCurrentV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActivesessionGetCurrentV1ResponseMPayload(sCustomerCode string, eActivesessionSessiontype string, fkiLanguageID int32, sCompanyNameX string, sDepartmentNameX string, aRegisteredModules []string, aPermissions []int32, ) *ActivesessionGetCurrentV1ResponseMPayload {
	this := ActivesessionGetCurrentV1ResponseMPayload{}
	this.SCustomerCode = sCustomerCode
	this.EActivesessionSessiontype = eActivesessionSessiontype
	this.FkiLanguageID = fkiLanguageID
	this.SCompanyNameX = sCompanyNameX
	this.SDepartmentNameX = sDepartmentNameX
	this.ARegisteredModules = aRegisteredModules
	this.APermissions = aPermissions
	return &this
}

// NewActivesessionGetCurrentV1ResponseMPayloadWithDefaults instantiates a new ActivesessionGetCurrentV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActivesessionGetCurrentV1ResponseMPayloadWithDefaults() *ActivesessionGetCurrentV1ResponseMPayload {
	this := ActivesessionGetCurrentV1ResponseMPayload{}
	return &this
}

// GetSCustomerCode returns the SCustomerCode field value
func (o *ActivesessionGetCurrentV1ResponseMPayload) GetSCustomerCode() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.SCustomerCode
}

// GetSCustomerCodeOk returns a tuple with the SCustomerCode field value
// and a boolean to check if the value has been set.
func (o *ActivesessionGetCurrentV1ResponseMPayload) GetSCustomerCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SCustomerCode, true
}

// SetSCustomerCode sets field value
func (o *ActivesessionGetCurrentV1ResponseMPayload) SetSCustomerCode(v string) {
	o.SCustomerCode = v
}

// GetEActivesessionSessiontype returns the EActivesessionSessiontype field value
func (o *ActivesessionGetCurrentV1ResponseMPayload) GetEActivesessionSessiontype() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.EActivesessionSessiontype
}

// GetEActivesessionSessiontypeOk returns a tuple with the EActivesessionSessiontype field value
// and a boolean to check if the value has been set.
func (o *ActivesessionGetCurrentV1ResponseMPayload) GetEActivesessionSessiontypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EActivesessionSessiontype, true
}

// SetEActivesessionSessiontype sets field value
func (o *ActivesessionGetCurrentV1ResponseMPayload) SetEActivesessionSessiontype(v string) {
	o.EActivesessionSessiontype = v
}

// GetFkiLanguageID returns the FkiLanguageID field value
func (o *ActivesessionGetCurrentV1ResponseMPayload) GetFkiLanguageID() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.FkiLanguageID
}

// GetFkiLanguageIDOk returns a tuple with the FkiLanguageID field value
// and a boolean to check if the value has been set.
func (o *ActivesessionGetCurrentV1ResponseMPayload) GetFkiLanguageIDOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FkiLanguageID, true
}

// SetFkiLanguageID sets field value
func (o *ActivesessionGetCurrentV1ResponseMPayload) SetFkiLanguageID(v int32) {
	o.FkiLanguageID = v
}

// GetSCompanyNameX returns the SCompanyNameX field value
func (o *ActivesessionGetCurrentV1ResponseMPayload) GetSCompanyNameX() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.SCompanyNameX
}

// GetSCompanyNameXOk returns a tuple with the SCompanyNameX field value
// and a boolean to check if the value has been set.
func (o *ActivesessionGetCurrentV1ResponseMPayload) GetSCompanyNameXOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SCompanyNameX, true
}

// SetSCompanyNameX sets field value
func (o *ActivesessionGetCurrentV1ResponseMPayload) SetSCompanyNameX(v string) {
	o.SCompanyNameX = v
}

// GetSDepartmentNameX returns the SDepartmentNameX field value
func (o *ActivesessionGetCurrentV1ResponseMPayload) GetSDepartmentNameX() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.SDepartmentNameX
}

// GetSDepartmentNameXOk returns a tuple with the SDepartmentNameX field value
// and a boolean to check if the value has been set.
func (o *ActivesessionGetCurrentV1ResponseMPayload) GetSDepartmentNameXOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SDepartmentNameX, true
}

// SetSDepartmentNameX sets field value
func (o *ActivesessionGetCurrentV1ResponseMPayload) SetSDepartmentNameX(v string) {
	o.SDepartmentNameX = v
}

// GetARegisteredModules returns the ARegisteredModules field value
func (o *ActivesessionGetCurrentV1ResponseMPayload) GetARegisteredModules() []string {
	if o == nil  {
		var ret []string
		return ret
	}

	return o.ARegisteredModules
}

// GetARegisteredModulesOk returns a tuple with the ARegisteredModules field value
// and a boolean to check if the value has been set.
func (o *ActivesessionGetCurrentV1ResponseMPayload) GetARegisteredModulesOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ARegisteredModules, true
}

// SetARegisteredModules sets field value
func (o *ActivesessionGetCurrentV1ResponseMPayload) SetARegisteredModules(v []string) {
	o.ARegisteredModules = v
}

// GetAPermissions returns the APermissions field value
func (o *ActivesessionGetCurrentV1ResponseMPayload) GetAPermissions() []int32 {
	if o == nil  {
		var ret []int32
		return ret
	}

	return o.APermissions
}

// GetAPermissionsOk returns a tuple with the APermissions field value
// and a boolean to check if the value has been set.
func (o *ActivesessionGetCurrentV1ResponseMPayload) GetAPermissionsOk() (*[]int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.APermissions, true
}

// SetAPermissions sets field value
func (o *ActivesessionGetCurrentV1ResponseMPayload) SetAPermissions(v []int32) {
	o.APermissions = v
}

func (o ActivesessionGetCurrentV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sCustomerCode"] = o.SCustomerCode
	}
	if true {
		toSerialize["eActivesessionSessiontype"] = o.EActivesessionSessiontype
	}
	if true {
		toSerialize["fkiLanguageID"] = o.FkiLanguageID
	}
	if true {
		toSerialize["sCompanyNameX"] = o.SCompanyNameX
	}
	if true {
		toSerialize["sDepartmentNameX"] = o.SDepartmentNameX
	}
	if true {
		toSerialize["a_RegisteredModules"] = o.ARegisteredModules
	}
	if true {
		toSerialize["a_Permissions"] = o.APermissions
	}
	return json.Marshal(toSerialize)
}

type NullableActivesessionGetCurrentV1ResponseMPayload struct {
	value *ActivesessionGetCurrentV1ResponseMPayload
	isSet bool
}

func (v NullableActivesessionGetCurrentV1ResponseMPayload) Get() *ActivesessionGetCurrentV1ResponseMPayload {
	return v.value
}

func (v *NullableActivesessionGetCurrentV1ResponseMPayload) Set(val *ActivesessionGetCurrentV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableActivesessionGetCurrentV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableActivesessionGetCurrentV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActivesessionGetCurrentV1ResponseMPayload(val *ActivesessionGetCurrentV1ResponseMPayload) *NullableActivesessionGetCurrentV1ResponseMPayload {
	return &NullableActivesessionGetCurrentV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableActivesessionGetCurrentV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActivesessionGetCurrentV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


