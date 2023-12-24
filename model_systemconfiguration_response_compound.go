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

// checks if the SystemconfigurationResponseCompound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SystemconfigurationResponseCompound{}

// SystemconfigurationResponseCompound A Systemconfiguration Object
type SystemconfigurationResponseCompound struct {
	// The unique ID of the Systemconfiguration
	PkiSystemconfigurationID int32 `json:"pkiSystemconfigurationID"`
	// The unique ID of the Systemconfigurationtype
	FkiSystemconfigurationtypeID int32 `json:"fkiSystemconfigurationtypeID"`
	// The description of the Systemconfigurationtype in the language of the requester
	SSystemconfigurationtypeDescriptionX string `json:"sSystemconfigurationtypeDescriptionX"`
	ESystemconfigurationNewexternaluseraction FieldESystemconfigurationNewexternaluseraction `json:"eSystemconfigurationNewexternaluseraction"`
	ESystemconfigurationLanguage1 FieldESystemconfigurationLanguage1 `json:"eSystemconfigurationLanguage1"`
	ESystemconfigurationLanguage2 FieldESystemconfigurationLanguage2 `json:"eSystemconfigurationLanguage2"`
	ESystemconfigurationEzsign FieldESystemconfigurationEzsign `json:"eSystemconfigurationEzsign"`
	// Whether if we allow the creation of personal files in eZsign
	BSystemconfigurationEzsignpersonnal bool `json:"bSystemconfigurationEzsignpersonnal"`
	// Whether if we allow SSPR
	BSystemconfigurationSspr bool `json:"bSystemconfigurationSspr"`
	// The start date where the system will be in read only
	DtSystemconfigurationReadonlyexpirationstart *string `json:"dtSystemconfigurationReadonlyexpirationstart,omitempty"`
	// The end date where the system will be in read only
	DtSystemconfigurationReadonlyexpirationend *string `json:"dtSystemconfigurationReadonlyexpirationend,omitempty"`
}

type _SystemconfigurationResponseCompound SystemconfigurationResponseCompound

// NewSystemconfigurationResponseCompound instantiates a new SystemconfigurationResponseCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemconfigurationResponseCompound(pkiSystemconfigurationID int32, fkiSystemconfigurationtypeID int32, sSystemconfigurationtypeDescriptionX string, eSystemconfigurationNewexternaluseraction FieldESystemconfigurationNewexternaluseraction, eSystemconfigurationLanguage1 FieldESystemconfigurationLanguage1, eSystemconfigurationLanguage2 FieldESystemconfigurationLanguage2, eSystemconfigurationEzsign FieldESystemconfigurationEzsign, bSystemconfigurationEzsignpersonnal bool, bSystemconfigurationSspr bool) *SystemconfigurationResponseCompound {
	this := SystemconfigurationResponseCompound{}
	this.PkiSystemconfigurationID = pkiSystemconfigurationID
	this.FkiSystemconfigurationtypeID = fkiSystemconfigurationtypeID
	this.SSystemconfigurationtypeDescriptionX = sSystemconfigurationtypeDescriptionX
	this.ESystemconfigurationNewexternaluseraction = eSystemconfigurationNewexternaluseraction
	this.ESystemconfigurationLanguage1 = eSystemconfigurationLanguage1
	this.ESystemconfigurationLanguage2 = eSystemconfigurationLanguage2
	this.ESystemconfigurationEzsign = eSystemconfigurationEzsign
	this.BSystemconfigurationEzsignpersonnal = bSystemconfigurationEzsignpersonnal
	this.BSystemconfigurationSspr = bSystemconfigurationSspr
	return &this
}

// NewSystemconfigurationResponseCompoundWithDefaults instantiates a new SystemconfigurationResponseCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemconfigurationResponseCompoundWithDefaults() *SystemconfigurationResponseCompound {
	this := SystemconfigurationResponseCompound{}
	return &this
}

// GetPkiSystemconfigurationID returns the PkiSystemconfigurationID field value
func (o *SystemconfigurationResponseCompound) GetPkiSystemconfigurationID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiSystemconfigurationID
}

// GetPkiSystemconfigurationIDOk returns a tuple with the PkiSystemconfigurationID field value
// and a boolean to check if the value has been set.
func (o *SystemconfigurationResponseCompound) GetPkiSystemconfigurationIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiSystemconfigurationID, true
}

// SetPkiSystemconfigurationID sets field value
func (o *SystemconfigurationResponseCompound) SetPkiSystemconfigurationID(v int32) {
	o.PkiSystemconfigurationID = v
}

// GetFkiSystemconfigurationtypeID returns the FkiSystemconfigurationtypeID field value
func (o *SystemconfigurationResponseCompound) GetFkiSystemconfigurationtypeID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiSystemconfigurationtypeID
}

// GetFkiSystemconfigurationtypeIDOk returns a tuple with the FkiSystemconfigurationtypeID field value
// and a boolean to check if the value has been set.
func (o *SystemconfigurationResponseCompound) GetFkiSystemconfigurationtypeIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiSystemconfigurationtypeID, true
}

// SetFkiSystemconfigurationtypeID sets field value
func (o *SystemconfigurationResponseCompound) SetFkiSystemconfigurationtypeID(v int32) {
	o.FkiSystemconfigurationtypeID = v
}

// GetSSystemconfigurationtypeDescriptionX returns the SSystemconfigurationtypeDescriptionX field value
func (o *SystemconfigurationResponseCompound) GetSSystemconfigurationtypeDescriptionX() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SSystemconfigurationtypeDescriptionX
}

// GetSSystemconfigurationtypeDescriptionXOk returns a tuple with the SSystemconfigurationtypeDescriptionX field value
// and a boolean to check if the value has been set.
func (o *SystemconfigurationResponseCompound) GetSSystemconfigurationtypeDescriptionXOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SSystemconfigurationtypeDescriptionX, true
}

// SetSSystemconfigurationtypeDescriptionX sets field value
func (o *SystemconfigurationResponseCompound) SetSSystemconfigurationtypeDescriptionX(v string) {
	o.SSystemconfigurationtypeDescriptionX = v
}

// GetESystemconfigurationNewexternaluseraction returns the ESystemconfigurationNewexternaluseraction field value
func (o *SystemconfigurationResponseCompound) GetESystemconfigurationNewexternaluseraction() FieldESystemconfigurationNewexternaluseraction {
	if o == nil {
		var ret FieldESystemconfigurationNewexternaluseraction
		return ret
	}

	return o.ESystemconfigurationNewexternaluseraction
}

// GetESystemconfigurationNewexternaluseractionOk returns a tuple with the ESystemconfigurationNewexternaluseraction field value
// and a boolean to check if the value has been set.
func (o *SystemconfigurationResponseCompound) GetESystemconfigurationNewexternaluseractionOk() (*FieldESystemconfigurationNewexternaluseraction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ESystemconfigurationNewexternaluseraction, true
}

// SetESystemconfigurationNewexternaluseraction sets field value
func (o *SystemconfigurationResponseCompound) SetESystemconfigurationNewexternaluseraction(v FieldESystemconfigurationNewexternaluseraction) {
	o.ESystemconfigurationNewexternaluseraction = v
}

// GetESystemconfigurationLanguage1 returns the ESystemconfigurationLanguage1 field value
func (o *SystemconfigurationResponseCompound) GetESystemconfigurationLanguage1() FieldESystemconfigurationLanguage1 {
	if o == nil {
		var ret FieldESystemconfigurationLanguage1
		return ret
	}

	return o.ESystemconfigurationLanguage1
}

// GetESystemconfigurationLanguage1Ok returns a tuple with the ESystemconfigurationLanguage1 field value
// and a boolean to check if the value has been set.
func (o *SystemconfigurationResponseCompound) GetESystemconfigurationLanguage1Ok() (*FieldESystemconfigurationLanguage1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ESystemconfigurationLanguage1, true
}

// SetESystemconfigurationLanguage1 sets field value
func (o *SystemconfigurationResponseCompound) SetESystemconfigurationLanguage1(v FieldESystemconfigurationLanguage1) {
	o.ESystemconfigurationLanguage1 = v
}

// GetESystemconfigurationLanguage2 returns the ESystemconfigurationLanguage2 field value
func (o *SystemconfigurationResponseCompound) GetESystemconfigurationLanguage2() FieldESystemconfigurationLanguage2 {
	if o == nil {
		var ret FieldESystemconfigurationLanguage2
		return ret
	}

	return o.ESystemconfigurationLanguage2
}

// GetESystemconfigurationLanguage2Ok returns a tuple with the ESystemconfigurationLanguage2 field value
// and a boolean to check if the value has been set.
func (o *SystemconfigurationResponseCompound) GetESystemconfigurationLanguage2Ok() (*FieldESystemconfigurationLanguage2, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ESystemconfigurationLanguage2, true
}

// SetESystemconfigurationLanguage2 sets field value
func (o *SystemconfigurationResponseCompound) SetESystemconfigurationLanguage2(v FieldESystemconfigurationLanguage2) {
	o.ESystemconfigurationLanguage2 = v
}

// GetESystemconfigurationEzsign returns the ESystemconfigurationEzsign field value
func (o *SystemconfigurationResponseCompound) GetESystemconfigurationEzsign() FieldESystemconfigurationEzsign {
	if o == nil {
		var ret FieldESystemconfigurationEzsign
		return ret
	}

	return o.ESystemconfigurationEzsign
}

// GetESystemconfigurationEzsignOk returns a tuple with the ESystemconfigurationEzsign field value
// and a boolean to check if the value has been set.
func (o *SystemconfigurationResponseCompound) GetESystemconfigurationEzsignOk() (*FieldESystemconfigurationEzsign, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ESystemconfigurationEzsign, true
}

// SetESystemconfigurationEzsign sets field value
func (o *SystemconfigurationResponseCompound) SetESystemconfigurationEzsign(v FieldESystemconfigurationEzsign) {
	o.ESystemconfigurationEzsign = v
}

// GetBSystemconfigurationEzsignpersonnal returns the BSystemconfigurationEzsignpersonnal field value
func (o *SystemconfigurationResponseCompound) GetBSystemconfigurationEzsignpersonnal() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BSystemconfigurationEzsignpersonnal
}

// GetBSystemconfigurationEzsignpersonnalOk returns a tuple with the BSystemconfigurationEzsignpersonnal field value
// and a boolean to check if the value has been set.
func (o *SystemconfigurationResponseCompound) GetBSystemconfigurationEzsignpersonnalOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BSystemconfigurationEzsignpersonnal, true
}

// SetBSystemconfigurationEzsignpersonnal sets field value
func (o *SystemconfigurationResponseCompound) SetBSystemconfigurationEzsignpersonnal(v bool) {
	o.BSystemconfigurationEzsignpersonnal = v
}

// GetBSystemconfigurationSspr returns the BSystemconfigurationSspr field value
func (o *SystemconfigurationResponseCompound) GetBSystemconfigurationSspr() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BSystemconfigurationSspr
}

// GetBSystemconfigurationSsprOk returns a tuple with the BSystemconfigurationSspr field value
// and a boolean to check if the value has been set.
func (o *SystemconfigurationResponseCompound) GetBSystemconfigurationSsprOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BSystemconfigurationSspr, true
}

// SetBSystemconfigurationSspr sets field value
func (o *SystemconfigurationResponseCompound) SetBSystemconfigurationSspr(v bool) {
	o.BSystemconfigurationSspr = v
}

// GetDtSystemconfigurationReadonlyexpirationstart returns the DtSystemconfigurationReadonlyexpirationstart field value if set, zero value otherwise.
func (o *SystemconfigurationResponseCompound) GetDtSystemconfigurationReadonlyexpirationstart() string {
	if o == nil || IsNil(o.DtSystemconfigurationReadonlyexpirationstart) {
		var ret string
		return ret
	}
	return *o.DtSystemconfigurationReadonlyexpirationstart
}

// GetDtSystemconfigurationReadonlyexpirationstartOk returns a tuple with the DtSystemconfigurationReadonlyexpirationstart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemconfigurationResponseCompound) GetDtSystemconfigurationReadonlyexpirationstartOk() (*string, bool) {
	if o == nil || IsNil(o.DtSystemconfigurationReadonlyexpirationstart) {
		return nil, false
	}
	return o.DtSystemconfigurationReadonlyexpirationstart, true
}

// HasDtSystemconfigurationReadonlyexpirationstart returns a boolean if a field has been set.
func (o *SystemconfigurationResponseCompound) HasDtSystemconfigurationReadonlyexpirationstart() bool {
	if o != nil && !IsNil(o.DtSystemconfigurationReadonlyexpirationstart) {
		return true
	}

	return false
}

// SetDtSystemconfigurationReadonlyexpirationstart gets a reference to the given string and assigns it to the DtSystemconfigurationReadonlyexpirationstart field.
func (o *SystemconfigurationResponseCompound) SetDtSystemconfigurationReadonlyexpirationstart(v string) {
	o.DtSystemconfigurationReadonlyexpirationstart = &v
}

// GetDtSystemconfigurationReadonlyexpirationend returns the DtSystemconfigurationReadonlyexpirationend field value if set, zero value otherwise.
func (o *SystemconfigurationResponseCompound) GetDtSystemconfigurationReadonlyexpirationend() string {
	if o == nil || IsNil(o.DtSystemconfigurationReadonlyexpirationend) {
		var ret string
		return ret
	}
	return *o.DtSystemconfigurationReadonlyexpirationend
}

// GetDtSystemconfigurationReadonlyexpirationendOk returns a tuple with the DtSystemconfigurationReadonlyexpirationend field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemconfigurationResponseCompound) GetDtSystemconfigurationReadonlyexpirationendOk() (*string, bool) {
	if o == nil || IsNil(o.DtSystemconfigurationReadonlyexpirationend) {
		return nil, false
	}
	return o.DtSystemconfigurationReadonlyexpirationend, true
}

// HasDtSystemconfigurationReadonlyexpirationend returns a boolean if a field has been set.
func (o *SystemconfigurationResponseCompound) HasDtSystemconfigurationReadonlyexpirationend() bool {
	if o != nil && !IsNil(o.DtSystemconfigurationReadonlyexpirationend) {
		return true
	}

	return false
}

// SetDtSystemconfigurationReadonlyexpirationend gets a reference to the given string and assigns it to the DtSystemconfigurationReadonlyexpirationend field.
func (o *SystemconfigurationResponseCompound) SetDtSystemconfigurationReadonlyexpirationend(v string) {
	o.DtSystemconfigurationReadonlyexpirationend = &v
}

func (o SystemconfigurationResponseCompound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SystemconfigurationResponseCompound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiSystemconfigurationID"] = o.PkiSystemconfigurationID
	toSerialize["fkiSystemconfigurationtypeID"] = o.FkiSystemconfigurationtypeID
	toSerialize["sSystemconfigurationtypeDescriptionX"] = o.SSystemconfigurationtypeDescriptionX
	toSerialize["eSystemconfigurationNewexternaluseraction"] = o.ESystemconfigurationNewexternaluseraction
	toSerialize["eSystemconfigurationLanguage1"] = o.ESystemconfigurationLanguage1
	toSerialize["eSystemconfigurationLanguage2"] = o.ESystemconfigurationLanguage2
	toSerialize["eSystemconfigurationEzsign"] = o.ESystemconfigurationEzsign
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

func (o *SystemconfigurationResponseCompound) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiSystemconfigurationID",
		"fkiSystemconfigurationtypeID",
		"sSystemconfigurationtypeDescriptionX",
		"eSystemconfigurationNewexternaluseraction",
		"eSystemconfigurationLanguage1",
		"eSystemconfigurationLanguage2",
		"eSystemconfigurationEzsign",
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

	varSystemconfigurationResponseCompound := _SystemconfigurationResponseCompound{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSystemconfigurationResponseCompound)

	if err != nil {
		return err
	}

	*o = SystemconfigurationResponseCompound(varSystemconfigurationResponseCompound)

	return err
}

type NullableSystemconfigurationResponseCompound struct {
	value *SystemconfigurationResponseCompound
	isSet bool
}

func (v NullableSystemconfigurationResponseCompound) Get() *SystemconfigurationResponseCompound {
	return v.value
}

func (v *NullableSystemconfigurationResponseCompound) Set(val *SystemconfigurationResponseCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemconfigurationResponseCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemconfigurationResponseCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemconfigurationResponseCompound(val *SystemconfigurationResponseCompound) *NullableSystemconfigurationResponseCompound {
	return &NullableSystemconfigurationResponseCompound{value: val, isSet: true}
}

func (v NullableSystemconfigurationResponseCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemconfigurationResponseCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


