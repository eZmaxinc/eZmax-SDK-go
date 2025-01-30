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

// checks if the EzsigntemplatesignerRequestCompound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigntemplatesignerRequestCompound{}

// EzsigntemplatesignerRequestCompound A Ezsigntemplatesigner Object and children
type EzsigntemplatesignerRequestCompound struct {
	// The unique ID of the Ezsigntemplatesigner
	PkiEzsigntemplatesignerID *int32 `json:"pkiEzsigntemplatesignerID,omitempty"`
	// The unique ID of the Ezsigntemplate
	FkiEzsigntemplateID int32 `json:"fkiEzsigntemplateID"`
	// The unique ID of the User
	FkiUserID *int32 `json:"fkiUserID,omitempty"`
	// The unique ID of the Usergroup
	FkiUsergroupID *int32 `json:"fkiUsergroupID,omitempty"`
	// The unique ID of the Ezdoctemplatedocument
	FkiEzdoctemplatedocumentID *int32 `json:"fkiEzdoctemplatedocumentID,omitempty"`
	// If this flag is true. The signatory will receive a copy of every signed Ezsigndocument even if it ain't required to sign the document.
	BEzsigntemplatesignerReceivecopy *bool `json:"bEzsigntemplatesignerReceivecopy,omitempty"`
	EEzsigntemplatesignerMapping *FieldEEzsigntemplatesignerMapping `json:"eEzsigntemplatesignerMapping,omitempty"`
	// The description of the Ezsigntemplatesigner
	SEzsigntemplatesignerDescription string `json:"sEzsigntemplatesignerDescription" validate:"regexp=^.{1,50}$"`
}

type _EzsigntemplatesignerRequestCompound EzsigntemplatesignerRequestCompound

// NewEzsigntemplatesignerRequestCompound instantiates a new EzsigntemplatesignerRequestCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigntemplatesignerRequestCompound(fkiEzsigntemplateID int32, sEzsigntemplatesignerDescription string) *EzsigntemplatesignerRequestCompound {
	this := EzsigntemplatesignerRequestCompound{}
	this.FkiEzsigntemplateID = fkiEzsigntemplateID
	this.SEzsigntemplatesignerDescription = sEzsigntemplatesignerDescription
	return &this
}

// NewEzsigntemplatesignerRequestCompoundWithDefaults instantiates a new EzsigntemplatesignerRequestCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigntemplatesignerRequestCompoundWithDefaults() *EzsigntemplatesignerRequestCompound {
	this := EzsigntemplatesignerRequestCompound{}
	return &this
}

// GetPkiEzsigntemplatesignerID returns the PkiEzsigntemplatesignerID field value if set, zero value otherwise.
func (o *EzsigntemplatesignerRequestCompound) GetPkiEzsigntemplatesignerID() int32 {
	if o == nil || IsNil(o.PkiEzsigntemplatesignerID) {
		var ret int32
		return ret
	}
	return *o.PkiEzsigntemplatesignerID
}

// GetPkiEzsigntemplatesignerIDOk returns a tuple with the PkiEzsigntemplatesignerID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplatesignerRequestCompound) GetPkiEzsigntemplatesignerIDOk() (*int32, bool) {
	if o == nil || IsNil(o.PkiEzsigntemplatesignerID) {
		return nil, false
	}
	return o.PkiEzsigntemplatesignerID, true
}

// HasPkiEzsigntemplatesignerID returns a boolean if a field has been set.
func (o *EzsigntemplatesignerRequestCompound) HasPkiEzsigntemplatesignerID() bool {
	if o != nil && !IsNil(o.PkiEzsigntemplatesignerID) {
		return true
	}

	return false
}

// SetPkiEzsigntemplatesignerID gets a reference to the given int32 and assigns it to the PkiEzsigntemplatesignerID field.
func (o *EzsigntemplatesignerRequestCompound) SetPkiEzsigntemplatesignerID(v int32) {
	o.PkiEzsigntemplatesignerID = &v
}

// GetFkiEzsigntemplateID returns the FkiEzsigntemplateID field value
func (o *EzsigntemplatesignerRequestCompound) GetFkiEzsigntemplateID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzsigntemplateID
}

// GetFkiEzsigntemplateIDOk returns a tuple with the FkiEzsigntemplateID field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatesignerRequestCompound) GetFkiEzsigntemplateIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiEzsigntemplateID, true
}

// SetFkiEzsigntemplateID sets field value
func (o *EzsigntemplatesignerRequestCompound) SetFkiEzsigntemplateID(v int32) {
	o.FkiEzsigntemplateID = v
}

// GetFkiUserID returns the FkiUserID field value if set, zero value otherwise.
func (o *EzsigntemplatesignerRequestCompound) GetFkiUserID() int32 {
	if o == nil || IsNil(o.FkiUserID) {
		var ret int32
		return ret
	}
	return *o.FkiUserID
}

// GetFkiUserIDOk returns a tuple with the FkiUserID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplatesignerRequestCompound) GetFkiUserIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiUserID) {
		return nil, false
	}
	return o.FkiUserID, true
}

// HasFkiUserID returns a boolean if a field has been set.
func (o *EzsigntemplatesignerRequestCompound) HasFkiUserID() bool {
	if o != nil && !IsNil(o.FkiUserID) {
		return true
	}

	return false
}

// SetFkiUserID gets a reference to the given int32 and assigns it to the FkiUserID field.
func (o *EzsigntemplatesignerRequestCompound) SetFkiUserID(v int32) {
	o.FkiUserID = &v
}

// GetFkiUsergroupID returns the FkiUsergroupID field value if set, zero value otherwise.
func (o *EzsigntemplatesignerRequestCompound) GetFkiUsergroupID() int32 {
	if o == nil || IsNil(o.FkiUsergroupID) {
		var ret int32
		return ret
	}
	return *o.FkiUsergroupID
}

// GetFkiUsergroupIDOk returns a tuple with the FkiUsergroupID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplatesignerRequestCompound) GetFkiUsergroupIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiUsergroupID) {
		return nil, false
	}
	return o.FkiUsergroupID, true
}

// HasFkiUsergroupID returns a boolean if a field has been set.
func (o *EzsigntemplatesignerRequestCompound) HasFkiUsergroupID() bool {
	if o != nil && !IsNil(o.FkiUsergroupID) {
		return true
	}

	return false
}

// SetFkiUsergroupID gets a reference to the given int32 and assigns it to the FkiUsergroupID field.
func (o *EzsigntemplatesignerRequestCompound) SetFkiUsergroupID(v int32) {
	o.FkiUsergroupID = &v
}

// GetFkiEzdoctemplatedocumentID returns the FkiEzdoctemplatedocumentID field value if set, zero value otherwise.
func (o *EzsigntemplatesignerRequestCompound) GetFkiEzdoctemplatedocumentID() int32 {
	if o == nil || IsNil(o.FkiEzdoctemplatedocumentID) {
		var ret int32
		return ret
	}
	return *o.FkiEzdoctemplatedocumentID
}

// GetFkiEzdoctemplatedocumentIDOk returns a tuple with the FkiEzdoctemplatedocumentID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplatesignerRequestCompound) GetFkiEzdoctemplatedocumentIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiEzdoctemplatedocumentID) {
		return nil, false
	}
	return o.FkiEzdoctemplatedocumentID, true
}

// HasFkiEzdoctemplatedocumentID returns a boolean if a field has been set.
func (o *EzsigntemplatesignerRequestCompound) HasFkiEzdoctemplatedocumentID() bool {
	if o != nil && !IsNil(o.FkiEzdoctemplatedocumentID) {
		return true
	}

	return false
}

// SetFkiEzdoctemplatedocumentID gets a reference to the given int32 and assigns it to the FkiEzdoctemplatedocumentID field.
func (o *EzsigntemplatesignerRequestCompound) SetFkiEzdoctemplatedocumentID(v int32) {
	o.FkiEzdoctemplatedocumentID = &v
}

// GetBEzsigntemplatesignerReceivecopy returns the BEzsigntemplatesignerReceivecopy field value if set, zero value otherwise.
func (o *EzsigntemplatesignerRequestCompound) GetBEzsigntemplatesignerReceivecopy() bool {
	if o == nil || IsNil(o.BEzsigntemplatesignerReceivecopy) {
		var ret bool
		return ret
	}
	return *o.BEzsigntemplatesignerReceivecopy
}

// GetBEzsigntemplatesignerReceivecopyOk returns a tuple with the BEzsigntemplatesignerReceivecopy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplatesignerRequestCompound) GetBEzsigntemplatesignerReceivecopyOk() (*bool, bool) {
	if o == nil || IsNil(o.BEzsigntemplatesignerReceivecopy) {
		return nil, false
	}
	return o.BEzsigntemplatesignerReceivecopy, true
}

// HasBEzsigntemplatesignerReceivecopy returns a boolean if a field has been set.
func (o *EzsigntemplatesignerRequestCompound) HasBEzsigntemplatesignerReceivecopy() bool {
	if o != nil && !IsNil(o.BEzsigntemplatesignerReceivecopy) {
		return true
	}

	return false
}

// SetBEzsigntemplatesignerReceivecopy gets a reference to the given bool and assigns it to the BEzsigntemplatesignerReceivecopy field.
func (o *EzsigntemplatesignerRequestCompound) SetBEzsigntemplatesignerReceivecopy(v bool) {
	o.BEzsigntemplatesignerReceivecopy = &v
}

// GetEEzsigntemplatesignerMapping returns the EEzsigntemplatesignerMapping field value if set, zero value otherwise.
func (o *EzsigntemplatesignerRequestCompound) GetEEzsigntemplatesignerMapping() FieldEEzsigntemplatesignerMapping {
	if o == nil || IsNil(o.EEzsigntemplatesignerMapping) {
		var ret FieldEEzsigntemplatesignerMapping
		return ret
	}
	return *o.EEzsigntemplatesignerMapping
}

// GetEEzsigntemplatesignerMappingOk returns a tuple with the EEzsigntemplatesignerMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplatesignerRequestCompound) GetEEzsigntemplatesignerMappingOk() (*FieldEEzsigntemplatesignerMapping, bool) {
	if o == nil || IsNil(o.EEzsigntemplatesignerMapping) {
		return nil, false
	}
	return o.EEzsigntemplatesignerMapping, true
}

// HasEEzsigntemplatesignerMapping returns a boolean if a field has been set.
func (o *EzsigntemplatesignerRequestCompound) HasEEzsigntemplatesignerMapping() bool {
	if o != nil && !IsNil(o.EEzsigntemplatesignerMapping) {
		return true
	}

	return false
}

// SetEEzsigntemplatesignerMapping gets a reference to the given FieldEEzsigntemplatesignerMapping and assigns it to the EEzsigntemplatesignerMapping field.
func (o *EzsigntemplatesignerRequestCompound) SetEEzsigntemplatesignerMapping(v FieldEEzsigntemplatesignerMapping) {
	o.EEzsigntemplatesignerMapping = &v
}

// GetSEzsigntemplatesignerDescription returns the SEzsigntemplatesignerDescription field value
func (o *EzsigntemplatesignerRequestCompound) GetSEzsigntemplatesignerDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzsigntemplatesignerDescription
}

// GetSEzsigntemplatesignerDescriptionOk returns a tuple with the SEzsigntemplatesignerDescription field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatesignerRequestCompound) GetSEzsigntemplatesignerDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SEzsigntemplatesignerDescription, true
}

// SetSEzsigntemplatesignerDescription sets field value
func (o *EzsigntemplatesignerRequestCompound) SetSEzsigntemplatesignerDescription(v string) {
	o.SEzsigntemplatesignerDescription = v
}

func (o EzsigntemplatesignerRequestCompound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigntemplatesignerRequestCompound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PkiEzsigntemplatesignerID) {
		toSerialize["pkiEzsigntemplatesignerID"] = o.PkiEzsigntemplatesignerID
	}
	toSerialize["fkiEzsigntemplateID"] = o.FkiEzsigntemplateID
	if !IsNil(o.FkiUserID) {
		toSerialize["fkiUserID"] = o.FkiUserID
	}
	if !IsNil(o.FkiUsergroupID) {
		toSerialize["fkiUsergroupID"] = o.FkiUsergroupID
	}
	if !IsNil(o.FkiEzdoctemplatedocumentID) {
		toSerialize["fkiEzdoctemplatedocumentID"] = o.FkiEzdoctemplatedocumentID
	}
	if !IsNil(o.BEzsigntemplatesignerReceivecopy) {
		toSerialize["bEzsigntemplatesignerReceivecopy"] = o.BEzsigntemplatesignerReceivecopy
	}
	if !IsNil(o.EEzsigntemplatesignerMapping) {
		toSerialize["eEzsigntemplatesignerMapping"] = o.EEzsigntemplatesignerMapping
	}
	toSerialize["sEzsigntemplatesignerDescription"] = o.SEzsigntemplatesignerDescription
	return toSerialize, nil
}

func (o *EzsigntemplatesignerRequestCompound) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fkiEzsigntemplateID",
		"sEzsigntemplatesignerDescription",
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

	varEzsigntemplatesignerRequestCompound := _EzsigntemplatesignerRequestCompound{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsigntemplatesignerRequestCompound)

	if err != nil {
		return err
	}

	*o = EzsigntemplatesignerRequestCompound(varEzsigntemplatesignerRequestCompound)

	return err
}

type NullableEzsigntemplatesignerRequestCompound struct {
	value *EzsigntemplatesignerRequestCompound
	isSet bool
}

func (v NullableEzsigntemplatesignerRequestCompound) Get() *EzsigntemplatesignerRequestCompound {
	return v.value
}

func (v *NullableEzsigntemplatesignerRequestCompound) Set(val *EzsigntemplatesignerRequestCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigntemplatesignerRequestCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigntemplatesignerRequestCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigntemplatesignerRequestCompound(val *EzsigntemplatesignerRequestCompound) *NullableEzsigntemplatesignerRequestCompound {
	return &NullableEzsigntemplatesignerRequestCompound{value: val, isSet: true}
}

func (v NullableEzsigntemplatesignerRequestCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigntemplatesignerRequestCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


