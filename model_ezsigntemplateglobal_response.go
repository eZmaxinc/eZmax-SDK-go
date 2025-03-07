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

// checks if the EzsigntemplateglobalResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigntemplateglobalResponse{}

// EzsigntemplateglobalResponse A Ezsigntemplateglobal Object
type EzsigntemplateglobalResponse struct {
	// The unique ID of the Ezsigntemplateglobal
	PkiEzsigntemplateglobalID int32 `json:"pkiEzsigntemplateglobalID"`
	// The unique ID of the Ezsigntemplateglobaldocument
	FkiEzsigntemplateglobaldocumentID int32 `json:"fkiEzsigntemplateglobaldocumentID"`
	// The unique ID of the Module
	FkiModuleID int32 `json:"fkiModuleID"`
	// The Name of the Module in the language of the requester
	SModuleNameX *string `json:"sModuleNameX,omitempty"`
	// The unique ID of the Language.  Valid values:  |Value|Description| |-|-| |1|French| |2|English|
	FkiLanguageID int32 `json:"fkiLanguageID"`
	// The Name of the Language in the language of the requester
	SLanguageNameX string `json:"sLanguageNameX"`
	EEzsigntemplateglobalModule FieldEEzsigntemplateglobalModule `json:"eEzsigntemplateglobalModule"`
	EEzsigntemplateglobalSupplier FieldEEzsigntemplateglobalSupplier `json:"eEzsigntemplateglobalSupplier"`
	// The Code of the Ezsigntemplateglobal
	SEzsigntemplateglobalCode string `json:"sEzsigntemplateglobalCode" validate:"regexp=^.{0,10}$"`
	// The description of the Ezsigntemplate
	SEzsigntemplateglobalDescription string `json:"sEzsigntemplateglobalDescription"`
}

type _EzsigntemplateglobalResponse EzsigntemplateglobalResponse

// NewEzsigntemplateglobalResponse instantiates a new EzsigntemplateglobalResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigntemplateglobalResponse(pkiEzsigntemplateglobalID int32, fkiEzsigntemplateglobaldocumentID int32, fkiModuleID int32, fkiLanguageID int32, sLanguageNameX string, eEzsigntemplateglobalModule FieldEEzsigntemplateglobalModule, eEzsigntemplateglobalSupplier FieldEEzsigntemplateglobalSupplier, sEzsigntemplateglobalCode string, sEzsigntemplateglobalDescription string) *EzsigntemplateglobalResponse {
	this := EzsigntemplateglobalResponse{}
	this.PkiEzsigntemplateglobalID = pkiEzsigntemplateglobalID
	this.FkiEzsigntemplateglobaldocumentID = fkiEzsigntemplateglobaldocumentID
	this.FkiModuleID = fkiModuleID
	this.FkiLanguageID = fkiLanguageID
	this.SLanguageNameX = sLanguageNameX
	this.EEzsigntemplateglobalModule = eEzsigntemplateglobalModule
	this.EEzsigntemplateglobalSupplier = eEzsigntemplateglobalSupplier
	this.SEzsigntemplateglobalCode = sEzsigntemplateglobalCode
	this.SEzsigntemplateglobalDescription = sEzsigntemplateglobalDescription
	return &this
}

// NewEzsigntemplateglobalResponseWithDefaults instantiates a new EzsigntemplateglobalResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigntemplateglobalResponseWithDefaults() *EzsigntemplateglobalResponse {
	this := EzsigntemplateglobalResponse{}
	return &this
}

// GetPkiEzsigntemplateglobalID returns the PkiEzsigntemplateglobalID field value
func (o *EzsigntemplateglobalResponse) GetPkiEzsigntemplateglobalID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiEzsigntemplateglobalID
}

// GetPkiEzsigntemplateglobalIDOk returns a tuple with the PkiEzsigntemplateglobalID field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateglobalResponse) GetPkiEzsigntemplateglobalIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiEzsigntemplateglobalID, true
}

// SetPkiEzsigntemplateglobalID sets field value
func (o *EzsigntemplateglobalResponse) SetPkiEzsigntemplateglobalID(v int32) {
	o.PkiEzsigntemplateglobalID = v
}

// GetFkiEzsigntemplateglobaldocumentID returns the FkiEzsigntemplateglobaldocumentID field value
func (o *EzsigntemplateglobalResponse) GetFkiEzsigntemplateglobaldocumentID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzsigntemplateglobaldocumentID
}

// GetFkiEzsigntemplateglobaldocumentIDOk returns a tuple with the FkiEzsigntemplateglobaldocumentID field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateglobalResponse) GetFkiEzsigntemplateglobaldocumentIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiEzsigntemplateglobaldocumentID, true
}

// SetFkiEzsigntemplateglobaldocumentID sets field value
func (o *EzsigntemplateglobalResponse) SetFkiEzsigntemplateglobaldocumentID(v int32) {
	o.FkiEzsigntemplateglobaldocumentID = v
}

// GetFkiModuleID returns the FkiModuleID field value
func (o *EzsigntemplateglobalResponse) GetFkiModuleID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiModuleID
}

// GetFkiModuleIDOk returns a tuple with the FkiModuleID field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateglobalResponse) GetFkiModuleIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiModuleID, true
}

// SetFkiModuleID sets field value
func (o *EzsigntemplateglobalResponse) SetFkiModuleID(v int32) {
	o.FkiModuleID = v
}

// GetSModuleNameX returns the SModuleNameX field value if set, zero value otherwise.
func (o *EzsigntemplateglobalResponse) GetSModuleNameX() string {
	if o == nil || IsNil(o.SModuleNameX) {
		var ret string
		return ret
	}
	return *o.SModuleNameX
}

// GetSModuleNameXOk returns a tuple with the SModuleNameX field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateglobalResponse) GetSModuleNameXOk() (*string, bool) {
	if o == nil || IsNil(o.SModuleNameX) {
		return nil, false
	}
	return o.SModuleNameX, true
}

// HasSModuleNameX returns a boolean if a field has been set.
func (o *EzsigntemplateglobalResponse) HasSModuleNameX() bool {
	if o != nil && !IsNil(o.SModuleNameX) {
		return true
	}

	return false
}

// SetSModuleNameX gets a reference to the given string and assigns it to the SModuleNameX field.
func (o *EzsigntemplateglobalResponse) SetSModuleNameX(v string) {
	o.SModuleNameX = &v
}

// GetFkiLanguageID returns the FkiLanguageID field value
func (o *EzsigntemplateglobalResponse) GetFkiLanguageID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiLanguageID
}

// GetFkiLanguageIDOk returns a tuple with the FkiLanguageID field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateglobalResponse) GetFkiLanguageIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiLanguageID, true
}

// SetFkiLanguageID sets field value
func (o *EzsigntemplateglobalResponse) SetFkiLanguageID(v int32) {
	o.FkiLanguageID = v
}

// GetSLanguageNameX returns the SLanguageNameX field value
func (o *EzsigntemplateglobalResponse) GetSLanguageNameX() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SLanguageNameX
}

// GetSLanguageNameXOk returns a tuple with the SLanguageNameX field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateglobalResponse) GetSLanguageNameXOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SLanguageNameX, true
}

// SetSLanguageNameX sets field value
func (o *EzsigntemplateglobalResponse) SetSLanguageNameX(v string) {
	o.SLanguageNameX = v
}

// GetEEzsigntemplateglobalModule returns the EEzsigntemplateglobalModule field value
func (o *EzsigntemplateglobalResponse) GetEEzsigntemplateglobalModule() FieldEEzsigntemplateglobalModule {
	if o == nil {
		var ret FieldEEzsigntemplateglobalModule
		return ret
	}

	return o.EEzsigntemplateglobalModule
}

// GetEEzsigntemplateglobalModuleOk returns a tuple with the EEzsigntemplateglobalModule field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateglobalResponse) GetEEzsigntemplateglobalModuleOk() (*FieldEEzsigntemplateglobalModule, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EEzsigntemplateglobalModule, true
}

// SetEEzsigntemplateglobalModule sets field value
func (o *EzsigntemplateglobalResponse) SetEEzsigntemplateglobalModule(v FieldEEzsigntemplateglobalModule) {
	o.EEzsigntemplateglobalModule = v
}

// GetEEzsigntemplateglobalSupplier returns the EEzsigntemplateglobalSupplier field value
func (o *EzsigntemplateglobalResponse) GetEEzsigntemplateglobalSupplier() FieldEEzsigntemplateglobalSupplier {
	if o == nil {
		var ret FieldEEzsigntemplateglobalSupplier
		return ret
	}

	return o.EEzsigntemplateglobalSupplier
}

// GetEEzsigntemplateglobalSupplierOk returns a tuple with the EEzsigntemplateglobalSupplier field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateglobalResponse) GetEEzsigntemplateglobalSupplierOk() (*FieldEEzsigntemplateglobalSupplier, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EEzsigntemplateglobalSupplier, true
}

// SetEEzsigntemplateglobalSupplier sets field value
func (o *EzsigntemplateglobalResponse) SetEEzsigntemplateglobalSupplier(v FieldEEzsigntemplateglobalSupplier) {
	o.EEzsigntemplateglobalSupplier = v
}

// GetSEzsigntemplateglobalCode returns the SEzsigntemplateglobalCode field value
func (o *EzsigntemplateglobalResponse) GetSEzsigntemplateglobalCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzsigntemplateglobalCode
}

// GetSEzsigntemplateglobalCodeOk returns a tuple with the SEzsigntemplateglobalCode field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateglobalResponse) GetSEzsigntemplateglobalCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SEzsigntemplateglobalCode, true
}

// SetSEzsigntemplateglobalCode sets field value
func (o *EzsigntemplateglobalResponse) SetSEzsigntemplateglobalCode(v string) {
	o.SEzsigntemplateglobalCode = v
}

// GetSEzsigntemplateglobalDescription returns the SEzsigntemplateglobalDescription field value
func (o *EzsigntemplateglobalResponse) GetSEzsigntemplateglobalDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzsigntemplateglobalDescription
}

// GetSEzsigntemplateglobalDescriptionOk returns a tuple with the SEzsigntemplateglobalDescription field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateglobalResponse) GetSEzsigntemplateglobalDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SEzsigntemplateglobalDescription, true
}

// SetSEzsigntemplateglobalDescription sets field value
func (o *EzsigntemplateglobalResponse) SetSEzsigntemplateglobalDescription(v string) {
	o.SEzsigntemplateglobalDescription = v
}

func (o EzsigntemplateglobalResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigntemplateglobalResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiEzsigntemplateglobalID"] = o.PkiEzsigntemplateglobalID
	toSerialize["fkiEzsigntemplateglobaldocumentID"] = o.FkiEzsigntemplateglobaldocumentID
	toSerialize["fkiModuleID"] = o.FkiModuleID
	if !IsNil(o.SModuleNameX) {
		toSerialize["sModuleNameX"] = o.SModuleNameX
	}
	toSerialize["fkiLanguageID"] = o.FkiLanguageID
	toSerialize["sLanguageNameX"] = o.SLanguageNameX
	toSerialize["eEzsigntemplateglobalModule"] = o.EEzsigntemplateglobalModule
	toSerialize["eEzsigntemplateglobalSupplier"] = o.EEzsigntemplateglobalSupplier
	toSerialize["sEzsigntemplateglobalCode"] = o.SEzsigntemplateglobalCode
	toSerialize["sEzsigntemplateglobalDescription"] = o.SEzsigntemplateglobalDescription
	return toSerialize, nil
}

func (o *EzsigntemplateglobalResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiEzsigntemplateglobalID",
		"fkiEzsigntemplateglobaldocumentID",
		"fkiModuleID",
		"fkiLanguageID",
		"sLanguageNameX",
		"eEzsigntemplateglobalModule",
		"eEzsigntemplateglobalSupplier",
		"sEzsigntemplateglobalCode",
		"sEzsigntemplateglobalDescription",
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

	varEzsigntemplateglobalResponse := _EzsigntemplateglobalResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsigntemplateglobalResponse)

	if err != nil {
		return err
	}

	*o = EzsigntemplateglobalResponse(varEzsigntemplateglobalResponse)

	return err
}

type NullableEzsigntemplateglobalResponse struct {
	value *EzsigntemplateglobalResponse
	isSet bool
}

func (v NullableEzsigntemplateglobalResponse) Get() *EzsigntemplateglobalResponse {
	return v.value
}

func (v *NullableEzsigntemplateglobalResponse) Set(val *EzsigntemplateglobalResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigntemplateglobalResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigntemplateglobalResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigntemplateglobalResponse(val *EzsigntemplateglobalResponse) *NullableEzsigntemplateglobalResponse {
	return &NullableEzsigntemplateglobalResponse{value: val, isSet: true}
}

func (v NullableEzsigntemplateglobalResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigntemplateglobalResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


