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

// checks if the EzsigntemplateglobalResponseCompound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigntemplateglobalResponseCompound{}

// EzsigntemplateglobalResponseCompound A Ezsigntemplateglobal Object
type EzsigntemplateglobalResponseCompound struct {
	EzsigntemplateglobalResponse
	ObjEzsigntemplateglobaldocument *EzsigntemplateglobaldocumentResponse `json:"objEzsigntemplateglobaldocument,omitempty"`
	AObjEzsigntemplateglobalsigner []EzsigntemplateglobalsignerResponseCompound `json:"a_objEzsigntemplateglobalsigner"`
}

type _EzsigntemplateglobalResponseCompound EzsigntemplateglobalResponseCompound

// NewEzsigntemplateglobalResponseCompound instantiates a new EzsigntemplateglobalResponseCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigntemplateglobalResponseCompound(aObjEzsigntemplateglobalsigner []EzsigntemplateglobalsignerResponseCompound, pkiEzsigntemplateglobalID int32, fkiEzsigntemplateglobaldocumentID int32, fkiModuleID int32, fkiLanguageID int32, sLanguageNameX string, eEzsigntemplateglobalModule FieldEEzsigntemplateglobalModule, eEzsigntemplateglobalSupplier FieldEEzsigntemplateglobalSupplier, sEzsigntemplateglobalCode string, sEzsigntemplateglobalDescription string) *EzsigntemplateglobalResponseCompound {
	this := EzsigntemplateglobalResponseCompound{}
	this.PkiEzsigntemplateglobalID = pkiEzsigntemplateglobalID
	this.FkiEzsigntemplateglobaldocumentID = fkiEzsigntemplateglobaldocumentID
	this.FkiModuleID = fkiModuleID
	this.FkiLanguageID = fkiLanguageID
	this.SLanguageNameX = sLanguageNameX
	this.EEzsigntemplateglobalModule = eEzsigntemplateglobalModule
	this.EEzsigntemplateglobalSupplier = eEzsigntemplateglobalSupplier
	this.SEzsigntemplateglobalCode = sEzsigntemplateglobalCode
	this.SEzsigntemplateglobalDescription = sEzsigntemplateglobalDescription
	this.AObjEzsigntemplateglobalsigner = aObjEzsigntemplateglobalsigner
	return &this
}

// NewEzsigntemplateglobalResponseCompoundWithDefaults instantiates a new EzsigntemplateglobalResponseCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigntemplateglobalResponseCompoundWithDefaults() *EzsigntemplateglobalResponseCompound {
	this := EzsigntemplateglobalResponseCompound{}
	return &this
}

// GetObjEzsigntemplateglobaldocument returns the ObjEzsigntemplateglobaldocument field value if set, zero value otherwise.
func (o *EzsigntemplateglobalResponseCompound) GetObjEzsigntemplateglobaldocument() EzsigntemplateglobaldocumentResponse {
	if o == nil || IsNil(o.ObjEzsigntemplateglobaldocument) {
		var ret EzsigntemplateglobaldocumentResponse
		return ret
	}
	return *o.ObjEzsigntemplateglobaldocument
}

// GetObjEzsigntemplateglobaldocumentOk returns a tuple with the ObjEzsigntemplateglobaldocument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateglobalResponseCompound) GetObjEzsigntemplateglobaldocumentOk() (*EzsigntemplateglobaldocumentResponse, bool) {
	if o == nil || IsNil(o.ObjEzsigntemplateglobaldocument) {
		return nil, false
	}
	return o.ObjEzsigntemplateglobaldocument, true
}

// HasObjEzsigntemplateglobaldocument returns a boolean if a field has been set.
func (o *EzsigntemplateglobalResponseCompound) HasObjEzsigntemplateglobaldocument() bool {
	if o != nil && !IsNil(o.ObjEzsigntemplateglobaldocument) {
		return true
	}

	return false
}

// SetObjEzsigntemplateglobaldocument gets a reference to the given EzsigntemplateglobaldocumentResponse and assigns it to the ObjEzsigntemplateglobaldocument field.
func (o *EzsigntemplateglobalResponseCompound) SetObjEzsigntemplateglobaldocument(v EzsigntemplateglobaldocumentResponse) {
	o.ObjEzsigntemplateglobaldocument = &v
}

// GetAObjEzsigntemplateglobalsigner returns the AObjEzsigntemplateglobalsigner field value
func (o *EzsigntemplateglobalResponseCompound) GetAObjEzsigntemplateglobalsigner() []EzsigntemplateglobalsignerResponseCompound {
	if o == nil {
		var ret []EzsigntemplateglobalsignerResponseCompound
		return ret
	}

	return o.AObjEzsigntemplateglobalsigner
}

// GetAObjEzsigntemplateglobalsignerOk returns a tuple with the AObjEzsigntemplateglobalsigner field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateglobalResponseCompound) GetAObjEzsigntemplateglobalsignerOk() ([]EzsigntemplateglobalsignerResponseCompound, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjEzsigntemplateglobalsigner, true
}

// SetAObjEzsigntemplateglobalsigner sets field value
func (o *EzsigntemplateglobalResponseCompound) SetAObjEzsigntemplateglobalsigner(v []EzsigntemplateglobalsignerResponseCompound) {
	o.AObjEzsigntemplateglobalsigner = v
}

func (o EzsigntemplateglobalResponseCompound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigntemplateglobalResponseCompound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ObjEzsigntemplateglobaldocument) {
		toSerialize["objEzsigntemplateglobaldocument"] = o.ObjEzsigntemplateglobaldocument
	}
	toSerialize["a_objEzsigntemplateglobalsigner"] = o.AObjEzsigntemplateglobalsigner
	return toSerialize, nil
}

func (o *EzsigntemplateglobalResponseCompound) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"a_objEzsigntemplateglobalsigner",
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

	varEzsigntemplateglobalResponseCompound := _EzsigntemplateglobalResponseCompound{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsigntemplateglobalResponseCompound)

	if err != nil {
		return err
	}

	*o = EzsigntemplateglobalResponseCompound(varEzsigntemplateglobalResponseCompound)

	return err
}

type NullableEzsigntemplateglobalResponseCompound struct {
	value *EzsigntemplateglobalResponseCompound
	isSet bool
}

func (v NullableEzsigntemplateglobalResponseCompound) Get() *EzsigntemplateglobalResponseCompound {
	return v.value
}

func (v *NullableEzsigntemplateglobalResponseCompound) Set(val *EzsigntemplateglobalResponseCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigntemplateglobalResponseCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigntemplateglobalResponseCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigntemplateglobalResponseCompound(val *EzsigntemplateglobalResponseCompound) *NullableEzsigntemplateglobalResponseCompound {
	return &NullableEzsigntemplateglobalResponseCompound{value: val, isSet: true}
}

func (v NullableEzsigntemplateglobalResponseCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigntemplateglobalResponseCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


