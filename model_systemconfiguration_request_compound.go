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

// checks if the SystemconfigurationRequestCompound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SystemconfigurationRequestCompound{}

// SystemconfigurationRequestCompound A Systemconfiguration Object and children
type SystemconfigurationRequestCompound struct {
	// The unique ID of the Systemconfiguration
	PkiSystemconfigurationID *int32 `json:"pkiSystemconfigurationID,omitempty"`
	// The unique ID of the Branding
	FkiBrandingID *int32 `json:"fkiBrandingID,omitempty"`
	ESystemconfigurationNewexternaluseraction FieldESystemconfigurationNewexternaluseraction `json:"eSystemconfigurationNewexternaluseraction"`
	ESystemconfigurationLanguage1 FieldESystemconfigurationLanguage1 `json:"eSystemconfigurationLanguage1"`
	ESystemconfigurationLanguage2 FieldESystemconfigurationLanguage2 `json:"eSystemconfigurationLanguage2"`
	// Deprecated
	ESystemconfigurationEzsign *FieldESystemconfigurationEzsign `json:"eSystemconfigurationEzsign,omitempty"`
	ESystemconfigurationEzsignofficeplan *FieldESystemconfigurationEzsignofficeplan `json:"eSystemconfigurationEzsignofficeplan,omitempty"`
	// Whether if Ezsign is paid by the company or not
	BSystemconfigurationEzsignpaidbyoffice *bool `json:"bSystemconfigurationEzsignpaidbyoffice,omitempty"`
	// Whether if we allow the creation of personal files in eZsign
	BSystemconfigurationEzsignpersonnal bool `json:"bSystemconfigurationEzsignpersonnal"`
	// Whether if we allow SSPR
	BSystemconfigurationSspr bool `json:"bSystemconfigurationSspr"`
	// The start date where the system will be in read only
	DtSystemconfigurationReadonlyexpirationstart *string `json:"dtSystemconfigurationReadonlyexpirationstart,omitempty" validate:"regexp=^[0-9]{4}-(0[1-9]|1[0-2])-(0[1-9]|[1-2][0-9]|3[0-1])$"`
	// The end date where the system will be in read only
	DtSystemconfigurationReadonlyexpirationend *string `json:"dtSystemconfigurationReadonlyexpirationend,omitempty" validate:"regexp=^[0-9]{4}-(0[1-9]|1[0-2])-(0[1-9]|[1-2][0-9]|3[0-1])$"`
}

type _SystemconfigurationRequestCompound SystemconfigurationRequestCompound

// NewSystemconfigurationRequestCompound instantiates a new SystemconfigurationRequestCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemconfigurationRequestCompound(eSystemconfigurationNewexternaluseraction FieldESystemconfigurationNewexternaluseraction, eSystemconfigurationLanguage1 FieldESystemconfigurationLanguage1, eSystemconfigurationLanguage2 FieldESystemconfigurationLanguage2, bSystemconfigurationEzsignpersonnal bool, bSystemconfigurationSspr bool) *SystemconfigurationRequestCompound {
	this := SystemconfigurationRequestCompound{}
	this.ESystemconfigurationNewexternaluseraction = eSystemconfigurationNewexternaluseraction
	this.ESystemconfigurationLanguage1 = eSystemconfigurationLanguage1
	this.ESystemconfigurationLanguage2 = eSystemconfigurationLanguage2
	this.BSystemconfigurationEzsignpersonnal = bSystemconfigurationEzsignpersonnal
	this.BSystemconfigurationSspr = bSystemconfigurationSspr
	return &this
}

// NewSystemconfigurationRequestCompoundWithDefaults instantiates a new SystemconfigurationRequestCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemconfigurationRequestCompoundWithDefaults() *SystemconfigurationRequestCompound {
	this := SystemconfigurationRequestCompound{}
	return &this
}

// GetPkiSystemconfigurationID returns the PkiSystemconfigurationID field value if set, zero value otherwise.
func (o *SystemconfigurationRequestCompound) GetPkiSystemconfigurationID() int32 {
	if o == nil || IsNil(o.PkiSystemconfigurationID) {
		var ret int32
		return ret
	}
	return *o.PkiSystemconfigurationID
}

// GetPkiSystemconfigurationIDOk returns a tuple with the PkiSystemconfigurationID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemconfigurationRequestCompound) GetPkiSystemconfigurationIDOk() (*int32, bool) {
	if o == nil || IsNil(o.PkiSystemconfigurationID) {
		return nil, false
	}
	return o.PkiSystemconfigurationID, true
}

// HasPkiSystemconfigurationID returns a boolean if a field has been set.
func (o *SystemconfigurationRequestCompound) HasPkiSystemconfigurationID() bool {
	if o != nil && !IsNil(o.PkiSystemconfigurationID) {
		return true
	}

	return false
}

// SetPkiSystemconfigurationID gets a reference to the given int32 and assigns it to the PkiSystemconfigurationID field.
func (o *SystemconfigurationRequestCompound) SetPkiSystemconfigurationID(v int32) {
	o.PkiSystemconfigurationID = &v
}

// GetFkiBrandingID returns the FkiBrandingID field value if set, zero value otherwise.
func (o *SystemconfigurationRequestCompound) GetFkiBrandingID() int32 {
	if o == nil || IsNil(o.FkiBrandingID) {
		var ret int32
		return ret
	}
	return *o.FkiBrandingID
}

// GetFkiBrandingIDOk returns a tuple with the FkiBrandingID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemconfigurationRequestCompound) GetFkiBrandingIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiBrandingID) {
		return nil, false
	}
	return o.FkiBrandingID, true
}

// HasFkiBrandingID returns a boolean if a field has been set.
func (o *SystemconfigurationRequestCompound) HasFkiBrandingID() bool {
	if o != nil && !IsNil(o.FkiBrandingID) {
		return true
	}

	return false
}

// SetFkiBrandingID gets a reference to the given int32 and assigns it to the FkiBrandingID field.
func (o *SystemconfigurationRequestCompound) SetFkiBrandingID(v int32) {
	o.FkiBrandingID = &v
}

// GetESystemconfigurationNewexternaluseraction returns the ESystemconfigurationNewexternaluseraction field value
func (o *SystemconfigurationRequestCompound) GetESystemconfigurationNewexternaluseraction() FieldESystemconfigurationNewexternaluseraction {
	if o == nil {
		var ret FieldESystemconfigurationNewexternaluseraction
		return ret
	}

	return o.ESystemconfigurationNewexternaluseraction
}

// GetESystemconfigurationNewexternaluseractionOk returns a tuple with the ESystemconfigurationNewexternaluseraction field value
// and a boolean to check if the value has been set.
func (o *SystemconfigurationRequestCompound) GetESystemconfigurationNewexternaluseractionOk() (*FieldESystemconfigurationNewexternaluseraction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ESystemconfigurationNewexternaluseraction, true
}

// SetESystemconfigurationNewexternaluseraction sets field value
func (o *SystemconfigurationRequestCompound) SetESystemconfigurationNewexternaluseraction(v FieldESystemconfigurationNewexternaluseraction) {
	o.ESystemconfigurationNewexternaluseraction = v
}

// GetESystemconfigurationLanguage1 returns the ESystemconfigurationLanguage1 field value
func (o *SystemconfigurationRequestCompound) GetESystemconfigurationLanguage1() FieldESystemconfigurationLanguage1 {
	if o == nil {
		var ret FieldESystemconfigurationLanguage1
		return ret
	}

	return o.ESystemconfigurationLanguage1
}

// GetESystemconfigurationLanguage1Ok returns a tuple with the ESystemconfigurationLanguage1 field value
// and a boolean to check if the value has been set.
func (o *SystemconfigurationRequestCompound) GetESystemconfigurationLanguage1Ok() (*FieldESystemconfigurationLanguage1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ESystemconfigurationLanguage1, true
}

// SetESystemconfigurationLanguage1 sets field value
func (o *SystemconfigurationRequestCompound) SetESystemconfigurationLanguage1(v FieldESystemconfigurationLanguage1) {
	o.ESystemconfigurationLanguage1 = v
}

// GetESystemconfigurationLanguage2 returns the ESystemconfigurationLanguage2 field value
func (o *SystemconfigurationRequestCompound) GetESystemconfigurationLanguage2() FieldESystemconfigurationLanguage2 {
	if o == nil {
		var ret FieldESystemconfigurationLanguage2
		return ret
	}

	return o.ESystemconfigurationLanguage2
}

// GetESystemconfigurationLanguage2Ok returns a tuple with the ESystemconfigurationLanguage2 field value
// and a boolean to check if the value has been set.
func (o *SystemconfigurationRequestCompound) GetESystemconfigurationLanguage2Ok() (*FieldESystemconfigurationLanguage2, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ESystemconfigurationLanguage2, true
}

// SetESystemconfigurationLanguage2 sets field value
func (o *SystemconfigurationRequestCompound) SetESystemconfigurationLanguage2(v FieldESystemconfigurationLanguage2) {
	o.ESystemconfigurationLanguage2 = v
}

// GetESystemconfigurationEzsign returns the ESystemconfigurationEzsign field value if set, zero value otherwise.
// Deprecated
func (o *SystemconfigurationRequestCompound) GetESystemconfigurationEzsign() FieldESystemconfigurationEzsign {
	if o == nil || IsNil(o.ESystemconfigurationEzsign) {
		var ret FieldESystemconfigurationEzsign
		return ret
	}
	return *o.ESystemconfigurationEzsign
}

// GetESystemconfigurationEzsignOk returns a tuple with the ESystemconfigurationEzsign field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *SystemconfigurationRequestCompound) GetESystemconfigurationEzsignOk() (*FieldESystemconfigurationEzsign, bool) {
	if o == nil || IsNil(o.ESystemconfigurationEzsign) {
		return nil, false
	}
	return o.ESystemconfigurationEzsign, true
}

// HasESystemconfigurationEzsign returns a boolean if a field has been set.
func (o *SystemconfigurationRequestCompound) HasESystemconfigurationEzsign() bool {
	if o != nil && !IsNil(o.ESystemconfigurationEzsign) {
		return true
	}

	return false
}

// SetESystemconfigurationEzsign gets a reference to the given FieldESystemconfigurationEzsign and assigns it to the ESystemconfigurationEzsign field.
// Deprecated
func (o *SystemconfigurationRequestCompound) SetESystemconfigurationEzsign(v FieldESystemconfigurationEzsign) {
	o.ESystemconfigurationEzsign = &v
}

// GetESystemconfigurationEzsignofficeplan returns the ESystemconfigurationEzsignofficeplan field value if set, zero value otherwise.
func (o *SystemconfigurationRequestCompound) GetESystemconfigurationEzsignofficeplan() FieldESystemconfigurationEzsignofficeplan {
	if o == nil || IsNil(o.ESystemconfigurationEzsignofficeplan) {
		var ret FieldESystemconfigurationEzsignofficeplan
		return ret
	}
	return *o.ESystemconfigurationEzsignofficeplan
}

// GetESystemconfigurationEzsignofficeplanOk returns a tuple with the ESystemconfigurationEzsignofficeplan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemconfigurationRequestCompound) GetESystemconfigurationEzsignofficeplanOk() (*FieldESystemconfigurationEzsignofficeplan, bool) {
	if o == nil || IsNil(o.ESystemconfigurationEzsignofficeplan) {
		return nil, false
	}
	return o.ESystemconfigurationEzsignofficeplan, true
}

// HasESystemconfigurationEzsignofficeplan returns a boolean if a field has been set.
func (o *SystemconfigurationRequestCompound) HasESystemconfigurationEzsignofficeplan() bool {
	if o != nil && !IsNil(o.ESystemconfigurationEzsignofficeplan) {
		return true
	}

	return false
}

// SetESystemconfigurationEzsignofficeplan gets a reference to the given FieldESystemconfigurationEzsignofficeplan and assigns it to the ESystemconfigurationEzsignofficeplan field.
func (o *SystemconfigurationRequestCompound) SetESystemconfigurationEzsignofficeplan(v FieldESystemconfigurationEzsignofficeplan) {
	o.ESystemconfigurationEzsignofficeplan = &v
}

// GetBSystemconfigurationEzsignpaidbyoffice returns the BSystemconfigurationEzsignpaidbyoffice field value if set, zero value otherwise.
func (o *SystemconfigurationRequestCompound) GetBSystemconfigurationEzsignpaidbyoffice() bool {
	if o == nil || IsNil(o.BSystemconfigurationEzsignpaidbyoffice) {
		var ret bool
		return ret
	}
	return *o.BSystemconfigurationEzsignpaidbyoffice
}

// GetBSystemconfigurationEzsignpaidbyofficeOk returns a tuple with the BSystemconfigurationEzsignpaidbyoffice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemconfigurationRequestCompound) GetBSystemconfigurationEzsignpaidbyofficeOk() (*bool, bool) {
	if o == nil || IsNil(o.BSystemconfigurationEzsignpaidbyoffice) {
		return nil, false
	}
	return o.BSystemconfigurationEzsignpaidbyoffice, true
}

// HasBSystemconfigurationEzsignpaidbyoffice returns a boolean if a field has been set.
func (o *SystemconfigurationRequestCompound) HasBSystemconfigurationEzsignpaidbyoffice() bool {
	if o != nil && !IsNil(o.BSystemconfigurationEzsignpaidbyoffice) {
		return true
	}

	return false
}

// SetBSystemconfigurationEzsignpaidbyoffice gets a reference to the given bool and assigns it to the BSystemconfigurationEzsignpaidbyoffice field.
func (o *SystemconfigurationRequestCompound) SetBSystemconfigurationEzsignpaidbyoffice(v bool) {
	o.BSystemconfigurationEzsignpaidbyoffice = &v
}

// GetBSystemconfigurationEzsignpersonnal returns the BSystemconfigurationEzsignpersonnal field value
func (o *SystemconfigurationRequestCompound) GetBSystemconfigurationEzsignpersonnal() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BSystemconfigurationEzsignpersonnal
}

// GetBSystemconfigurationEzsignpersonnalOk returns a tuple with the BSystemconfigurationEzsignpersonnal field value
// and a boolean to check if the value has been set.
func (o *SystemconfigurationRequestCompound) GetBSystemconfigurationEzsignpersonnalOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BSystemconfigurationEzsignpersonnal, true
}

// SetBSystemconfigurationEzsignpersonnal sets field value
func (o *SystemconfigurationRequestCompound) SetBSystemconfigurationEzsignpersonnal(v bool) {
	o.BSystemconfigurationEzsignpersonnal = v
}

// GetBSystemconfigurationSspr returns the BSystemconfigurationSspr field value
func (o *SystemconfigurationRequestCompound) GetBSystemconfigurationSspr() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BSystemconfigurationSspr
}

// GetBSystemconfigurationSsprOk returns a tuple with the BSystemconfigurationSspr field value
// and a boolean to check if the value has been set.
func (o *SystemconfigurationRequestCompound) GetBSystemconfigurationSsprOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BSystemconfigurationSspr, true
}

// SetBSystemconfigurationSspr sets field value
func (o *SystemconfigurationRequestCompound) SetBSystemconfigurationSspr(v bool) {
	o.BSystemconfigurationSspr = v
}

// GetDtSystemconfigurationReadonlyexpirationstart returns the DtSystemconfigurationReadonlyexpirationstart field value if set, zero value otherwise.
func (o *SystemconfigurationRequestCompound) GetDtSystemconfigurationReadonlyexpirationstart() string {
	if o == nil || IsNil(o.DtSystemconfigurationReadonlyexpirationstart) {
		var ret string
		return ret
	}
	return *o.DtSystemconfigurationReadonlyexpirationstart
}

// GetDtSystemconfigurationReadonlyexpirationstartOk returns a tuple with the DtSystemconfigurationReadonlyexpirationstart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemconfigurationRequestCompound) GetDtSystemconfigurationReadonlyexpirationstartOk() (*string, bool) {
	if o == nil || IsNil(o.DtSystemconfigurationReadonlyexpirationstart) {
		return nil, false
	}
	return o.DtSystemconfigurationReadonlyexpirationstart, true
}

// HasDtSystemconfigurationReadonlyexpirationstart returns a boolean if a field has been set.
func (o *SystemconfigurationRequestCompound) HasDtSystemconfigurationReadonlyexpirationstart() bool {
	if o != nil && !IsNil(o.DtSystemconfigurationReadonlyexpirationstart) {
		return true
	}

	return false
}

// SetDtSystemconfigurationReadonlyexpirationstart gets a reference to the given string and assigns it to the DtSystemconfigurationReadonlyexpirationstart field.
func (o *SystemconfigurationRequestCompound) SetDtSystemconfigurationReadonlyexpirationstart(v string) {
	o.DtSystemconfigurationReadonlyexpirationstart = &v
}

// GetDtSystemconfigurationReadonlyexpirationend returns the DtSystemconfigurationReadonlyexpirationend field value if set, zero value otherwise.
func (o *SystemconfigurationRequestCompound) GetDtSystemconfigurationReadonlyexpirationend() string {
	if o == nil || IsNil(o.DtSystemconfigurationReadonlyexpirationend) {
		var ret string
		return ret
	}
	return *o.DtSystemconfigurationReadonlyexpirationend
}

// GetDtSystemconfigurationReadonlyexpirationendOk returns a tuple with the DtSystemconfigurationReadonlyexpirationend field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemconfigurationRequestCompound) GetDtSystemconfigurationReadonlyexpirationendOk() (*string, bool) {
	if o == nil || IsNil(o.DtSystemconfigurationReadonlyexpirationend) {
		return nil, false
	}
	return o.DtSystemconfigurationReadonlyexpirationend, true
}

// HasDtSystemconfigurationReadonlyexpirationend returns a boolean if a field has been set.
func (o *SystemconfigurationRequestCompound) HasDtSystemconfigurationReadonlyexpirationend() bool {
	if o != nil && !IsNil(o.DtSystemconfigurationReadonlyexpirationend) {
		return true
	}

	return false
}

// SetDtSystemconfigurationReadonlyexpirationend gets a reference to the given string and assigns it to the DtSystemconfigurationReadonlyexpirationend field.
func (o *SystemconfigurationRequestCompound) SetDtSystemconfigurationReadonlyexpirationend(v string) {
	o.DtSystemconfigurationReadonlyexpirationend = &v
}

func (o SystemconfigurationRequestCompound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SystemconfigurationRequestCompound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PkiSystemconfigurationID) {
		toSerialize["pkiSystemconfigurationID"] = o.PkiSystemconfigurationID
	}
	if !IsNil(o.FkiBrandingID) {
		toSerialize["fkiBrandingID"] = o.FkiBrandingID
	}
	toSerialize["eSystemconfigurationNewexternaluseraction"] = o.ESystemconfigurationNewexternaluseraction
	toSerialize["eSystemconfigurationLanguage1"] = o.ESystemconfigurationLanguage1
	toSerialize["eSystemconfigurationLanguage2"] = o.ESystemconfigurationLanguage2
	if !IsNil(o.ESystemconfigurationEzsign) {
		toSerialize["eSystemconfigurationEzsign"] = o.ESystemconfigurationEzsign
	}
	if !IsNil(o.ESystemconfigurationEzsignofficeplan) {
		toSerialize["eSystemconfigurationEzsignofficeplan"] = o.ESystemconfigurationEzsignofficeplan
	}
	if !IsNil(o.BSystemconfigurationEzsignpaidbyoffice) {
		toSerialize["bSystemconfigurationEzsignpaidbyoffice"] = o.BSystemconfigurationEzsignpaidbyoffice
	}
	toSerialize["bSystemconfigurationEzsignpersonnal"] = o.BSystemconfigurationEzsignpersonnal
	toSerialize["bSystemconfigurationSspr"] = o.BSystemconfigurationSspr
	if !IsNil(o.DtSystemconfigurationReadonlyexpirationstart) {
		toSerialize["dtSystemconfigurationReadonlyexpirationstart"] = o.DtSystemconfigurationReadonlyexpirationstart
	}
	if !IsNil(o.DtSystemconfigurationReadonlyexpirationend) {
		toSerialize["dtSystemconfigurationReadonlyexpirationend"] = o.DtSystemconfigurationReadonlyexpirationend
	}
	return toSerialize, nil
}

func (o *SystemconfigurationRequestCompound) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"eSystemconfigurationNewexternaluseraction",
		"eSystemconfigurationLanguage1",
		"eSystemconfigurationLanguage2",
		"bSystemconfigurationEzsignpersonnal",
		"bSystemconfigurationSspr",
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

	varSystemconfigurationRequestCompound := _SystemconfigurationRequestCompound{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSystemconfigurationRequestCompound)

	if err != nil {
		return err
	}

	*o = SystemconfigurationRequestCompound(varSystemconfigurationRequestCompound)

	return err
}

type NullableSystemconfigurationRequestCompound struct {
	value *SystemconfigurationRequestCompound
	isSet bool
}

func (v NullableSystemconfigurationRequestCompound) Get() *SystemconfigurationRequestCompound {
	return v.value
}

func (v *NullableSystemconfigurationRequestCompound) Set(val *SystemconfigurationRequestCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemconfigurationRequestCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemconfigurationRequestCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemconfigurationRequestCompound(val *SystemconfigurationRequestCompound) *NullableSystemconfigurationRequestCompound {
	return &NullableSystemconfigurationRequestCompound{value: val, isSet: true}
}

func (v NullableSystemconfigurationRequestCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemconfigurationRequestCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


